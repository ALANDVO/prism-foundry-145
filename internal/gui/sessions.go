package gui

import (
	"prism-foundry-145/internal/sessions"
)

// combinedSessionStats aggregates session stats for UI display.
func combinedSessionStats() (int, error) {
	sm := sessions.NewSessionManager()
	return sm.StatsAndNative()
}
