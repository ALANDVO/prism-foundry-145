package sessions

import (
	"testing"
)

func TestRefreshNoDuplicatePaths(t *testing.T) {
	sm := NewSessionManager()
	// Refresh with no duplicates should work fine.
	sm.refresh([]string{"file1.txt", "file2.txt", "file3.txt"})
	// Publish should not panic with nil channels.
	sm.publish(map[string][]byte{
		"file1.txt": []byte("result1"),
		"file2.txt": []byte("result2"),
	})
	// After publish, parsing map should be cleaned up.
}

func TestPublishSkipsNilChannel(t *testing.T) {
	sm := NewSessionManager()
	// Set up a parsing entry with nil channel (simulating the bug scenario).
	sm.parsing["test.txt"] = &parsedFile{path: "test.txt"} // ch is nil zero value
	// Publish should skip nil channel without panicking.
	sm.publish(map[string][]byte{
		"test.txt": []byte("result"),
	})
	// After publish, entry should be cleaned up.
	// Verify the entry was removed.
	if _, ok := sm.parsing["test.txt"]; ok {
		t.Fatal("expected entry to be removed after publish")
	}
}

func TestStatsInNoPanic(t *testing.T) {
	sm := NewSessionManager()
	// Add an entry with a channel - but don't call publish first so the entry persists.
	sm.parsing["test.txt"] = &parsedFile{path: "test.txt", ch: make(chan []byte, 1)}
	// Send result directly to channel (simulating what publish would do).
	select {
	case sm.parsing["test.txt"].ch <- []byte("result"):
	default:
	}
	// StatsAndNative should not panic.
	count, err := sm.StatsAndNative()
	if err != nil {
		t.Fatalf("StatsAndNative returned error: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected count 1, got %d", count)
	}
}

func TestVerifySessionID(t *testing.T) {
	sm := NewSessionManager()
	// Unknown path should return false.
	if sm.VerifySessionID("nonexistent.txt") {
		t.Fatal("expected false for nonexistent path")
	}
	// Known path with nil channel should return false.
	sm.parsing["test.txt"] = &parsedFile{path: "test.txt"} // ch is nil
	if sm.VerifySessionID("test.txt") {
		t.Fatal("expected false for path with nil channel")
	}
	// Known path with non-nil channel should return true.
	sm.parsing["test.txt"] = &parsedFile{path: "test.txt", ch: make(chan []byte, 1)}
	if !sm.VerifySessionID("test.txt") {
		t.Fatal("expected true for path with non-nil channel")
	}
}
