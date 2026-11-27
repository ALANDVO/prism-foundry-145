package router

import "testing"

func TestSelectAvailableSkipsAtCapacity(t *testing.T) {
	group := []Provider{
		{ID: "codex-a", AtCapacity: true},
		{ID: "codex-b", AtCapacity: false},
		{ID: "custom-relay", AtCapacity: false},
	}
	id, ok := SelectAvailable(group)
	if !ok {
		t.Fatal("expected an available provider")
	}
	if id != "codex-b" {
		t.Fatalf("expected first non-capacity provider codex-b, got %q", id)
	}
}

func TestSelectAvailableAllAtCapacity(t *testing.T) {
	group := []Provider{
		{ID: "codex-a", AtCapacity: true},
		{ID: "codex-b", AtCapacity: true},
	}
	if id, ok := SelectAvailable(group); ok {
		t.Fatalf("expected no available provider, got %q", id)
	}
}

func TestSelectAvailableEmptyGroup(t *testing.T) {
	if id, ok := SelectAvailable(nil); ok {
		t.Fatalf("expected no available provider for empty group, got %q", id)
	}
}

func TestSelectAvailablePreservesPriorityOrder(t *testing.T) {
	group := []Provider{
		{ID: "primary", AtCapacity: false},
		{ID: "secondary", AtCapacity: false},
	}
	id, ok := SelectAvailable(group)
	if !ok || id != "primary" {
		t.Fatalf("expected priority order to pick primary, got %q ok=%v", id, ok)
	}
}
