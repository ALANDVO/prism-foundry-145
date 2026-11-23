package sessions

// Stats contains session statistics.
type Stats struct {
	Total       int
	Native      int
	Parsing     int
	ParsingFile map[string]bool
}

// combinedSessionStats aggregates stats from all sessions for UI display.
func combinedSessionStats() Stats {
	sm := NewSessionManager()
	total, _ := sm.StatsAndNative()
	return Stats{
		Total:   total,
		Native:  total,
	}
}
