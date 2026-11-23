package sessions

import "sync"

// parsedFile represents a file being parsed, with its channel for result delivery.
type parsedFile struct {
	path  string
	ch    chan []byte // result channel; may be nil if not yet set
}

// sessionsManager manages session refresh/publish lifecycle.
type sessionsManager struct {
	// parsing maps file paths to their parsing channels.
	parsing map[string]*parsedFile
	mu      sync.Mutex
}

// NewSessionManager creates a new sessions manager.
func NewSessionManager() *sessionsManager {
	return &sessionsManager{
		parsing: make(map[string]*parsedFile),
	}
}

// refresh collects files and sets up the parsing map and channels so publish can deliver results.
func (sm *sessionsManager) refresh(want []string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Reset parsing map for this refresh cycle.
	// Delete entries not in want; keep entries in want.
	for path := range sm.parsing {
		found := false
		for _, w := range want {
			if w == path {
				found = true
				break
			}
		}
		if !found {
			delete(sm.parsing, path)
		}
	}

	// Set up parsing entries for wanted files.
	// Guard against duplicate paths in want (issue #2932).
	seen := make(map[string]bool)
	for _, path := range want {
		if seen[path] {
			// Duplicate path in want - skip to avoid nil channel close.
			continue
		}
		seen[path] = true

		// Ensure entry exists; create if new.
		if _, exists := sm.parsing[path]; !exists {
			sm.parsing[path] = &parsedFile{path: path, ch: make(chan []byte, 1)}
		}
		// If entry already exists (from prior refresh not yet published), reuse its channel.
	}
}

// publish sends parsed results to channels and cleans up the parsing map.
// It is safe to call even if some channels are nil (will skip them).
// This is the function that previously panicked with "close of nil channel" at v0.1.1291:886.
func (sm *sessionsManager) publish(results map[string][]byte) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	for path, result := range results {
		pf, ok := sm.parsing[path]
		if !ok {
			// No entry for this path; skip.
			continue
		}
		// Guard: skip nil channel instead of closing it (fixes #2932).
		if pf.ch == nil {
			delete(sm.parsing, path)
			continue
		}
		select {
		case pf.ch <- result:
			// Result delivered.
		default:
			// Channel full; drop result.
		}
		// Clean up: remove entry after publish.
		delete(sm.parsing, path)
	}
}

// StatsAndNative returns native session stats; safe to call when no refresh is in progress.
func (sm *sessionsManager) StatsAndNative() (int, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Compute total from current parsing entries.
	total := 0
	for _, pf := range sm.parsing {
		if pf.ch != nil {
			// Try to read from channel without blocking.
			select {
			case <-pf.ch:
				total++
			default:
			}
		}
	}
	return total, nil
}

// VerifySessionID checks if a session path has a valid parsing channel.
func (sm *sessionsManager) VerifySessionID(path string) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	pf, ok := sm.parsing[path]
	if !ok {
		return false
	}
	return pf.ch != nil
}
