package sessions

// Warm starts the session warm-up process.
func Warm() {
	// Previously this would panic due to nil channel close in statsIn/publish.
	// With the nil-channel guard, this no longer panics.
	// Use NewSessionManager instead of newSessionsManager
	_ = NewSessionManager()
}
