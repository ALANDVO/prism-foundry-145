// Package router selects an available provider from a routing group.
//
// It addresses issue #1418: when a group's first provider reports
// "Selected model is at capacity", routing must fall through to the next
// provider that can still serve the request instead of surfacing the
// capacity error to the user.
package router

// Provider is a single routable backend within a group.
type Provider struct {
	// ID uniquely identifies the provider within its group.
	ID string
	// AtCapacity reports whether the provider's selected model is currently
	// unable to accept new requests (e.g. "Selected model is at capacity").
	AtCapacity bool
}

// SelectAvailable returns the ID of the first provider in group that is not
// at capacity, preserving the caller-supplied order (priority order).
//
// It returns ok=false when every provider is at capacity or the group is
// empty, so callers can surface a real "all providers at capacity" error
// rather than silently mis-routing.
func SelectAvailable(group []Provider) (id string, ok bool) {
	for _, p := range group {
		if !p.AtCapacity {
			return p.ID, true
		}
	}
	return "", false
}
