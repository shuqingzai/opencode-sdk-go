// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
)

// TestPtyEventCarrierFields verifies that PtyEvent (union carrier) correctly
// populates the common public fields (ID, Metadata, Durable, Location) from a
// variant during JSON unmarshaling. Each variant struct must contain the
// same shared fields so that apijson.Port can copy them from variant to carrier.
func TestPtyEventCarrierFields(t *testing.T) {
	jsonData := `{"id":"evt_123","type":"pty.created","metadata":{"k":"v"},"durable":{"aggregateID":"pty","seq":1,"version":1},"location":{"directory":"/tmp","workspaceID":"ws1"},"data":{"info":{"id":"pty_1","title":"test","command":"bash","args":[],"cwd":"/","status":"running","pid":1234}}}`
	var event opencode.PtyEvent
	if err := json.Unmarshal([]byte(jsonData), &event); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if event.ID != "evt_123" {
		t.Errorf("event.ID = %q, want %q", event.ID, "evt_123")
	}
	if event.Metadata == nil {
		t.Errorf("event.Metadata is nil, want non-nil")
	}
	if event.Durable == nil {
		t.Errorf("event.Durable is nil, want non-nil")
	}
	if event.Location == nil {
		t.Errorf("event.Location is nil, want non-nil")
	}
	if event.Type != opencode.PtyEventTypePtyCreated {
		t.Errorf("event.Type = %q, want %q", event.Type, opencode.PtyEventTypePtyCreated)
	}
	if event.Data == nil {
		t.Errorf("event.Data is nil, want non-nil")
	}
}

// TestPtyEventVariantDirectUnmarshal verifies that direct unmarshal into a
// specific variant struct (e.g. PtyCreatedEvent) populates the common fields.
func TestPtyEventVariantDirectUnmarshal(t *testing.T) {
	jsonData := `{"id":"evt_456","type":"pty.created","metadata":{"k":"v"},"durable":{"aggregateID":"pty","seq":2,"version":2},"location":{"directory":"/var","workspaceID":"ws2"},"data":{"info":{"id":"pty_2","title":"test2","command":"sh","args":[],"cwd":"/","status":"running","pid":5678}}}`
	var event opencode.PtyCreatedEvent
	if err := json.Unmarshal([]byte(jsonData), &event); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if event.ID != "evt_456" {
		t.Errorf("event.ID = %q, want %q", event.ID, "evt_456")
	}
	if event.Metadata == nil {
		t.Errorf("event.Metadata is nil, want non-nil")
	}
	if event.Durable == nil {
		t.Errorf("event.Durable is nil, want non-nil")
	}
	if event.Location == nil {
		t.Errorf("event.Location is nil, want non-nil")
	}
}

// TestPtyExitedEventVariantDirectUnmarshal verifies the ID field name fix
// (Id -> ID) on PtyExitedEventData and carrier field propagation for the
// PtyExitedEvent variant.
func TestPtyExitedEventVariantDirectUnmarshal(t *testing.T) {
	jsonData := `{"id":"evt_exited","type":"pty.exited","metadata":{"k":"v"},"durable":{"aggregateID":"pty","seq":3,"version":3},"location":{"directory":"/etc","workspaceID":"ws3"},"data":{"id":"pty_3","exitCode":0}}`
	var event opencode.PtyExitedEvent
	if err := json.Unmarshal([]byte(jsonData), &event); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if event.ID != "evt_exited" {
		t.Errorf("event.ID = %q, want %q", event.ID, "evt_exited")
	}
	if event.Data.ID != "pty_3" {
		t.Errorf("event.Data.ID = %q, want %q", event.Data.ID, "pty_3")
	}
	if event.Data.ExitCode != 0 {
		t.Errorf("event.Data.ExitCode = %d, want 0", event.Data.ExitCode)
	}
}

// TestPtyDeletedEventVariantDirectUnmarshal verifies the ID field name fix
// (Id -> ID) on PtyDeletedEventData and carrier field propagation for the
// PtyDeletedEvent variant.
func TestPtyDeletedEventVariantDirectUnmarshal(t *testing.T) {
	jsonData := `{"id":"evt_deleted","type":"pty.deleted","metadata":{"k":"v"},"durable":{"aggregateID":"pty","seq":4,"version":4},"location":{"directory":"/","workspaceID":"ws4"},"data":{"id":"pty_4"}}`
	var event opencode.PtyDeletedEvent
	if err := json.Unmarshal([]byte(jsonData), &event); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if event.ID != "evt_deleted" {
		t.Errorf("event.ID = %q, want %q", event.ID, "evt_deleted")
	}
	if event.Data.ID != "pty_4" {
		t.Errorf("event.Data.ID = %q, want %q", event.Data.ID, "pty_4")
	}
}

// TestPtyUpdatedEventVariantDirectUnmarshal verifies carrier field propagation
// for the PtyUpdatedEvent variant.
func TestPtyUpdatedEventVariantDirectUnmarshal(t *testing.T) {
	jsonData := `{"id":"evt_updated","type":"pty.updated","metadata":{"k":"v"},"durable":{"aggregateID":"pty","seq":5,"version":5},"location":{"directory":"/","workspaceID":"ws5"},"data":{"info":{"id":"pty_5","title":"updated","command":"bash","args":[],"cwd":"/","status":"running","pid":42}}}`
	var event opencode.PtyUpdatedEvent
	if err := json.Unmarshal([]byte(jsonData), &event); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if event.ID != "evt_updated" {
		t.Errorf("event.ID = %q, want %q", event.ID, "evt_updated")
	}
}
