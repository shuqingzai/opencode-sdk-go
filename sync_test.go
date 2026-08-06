package opencode

import (
	"encoding/json"
	"testing"
)

// TestSyncReplayParamsMarshalJSON_WithBody verifies that when Body is present,
// MarshalJSON serializes the body correctly.
func TestSyncReplayParamsMarshalJSON_WithBody(t *testing.T) {
	t.Parallel()
	params := SyncReplayParams{
		Directory: String("dir"),
		Workspace: String("ws"),
		Body: F(SyncReplayParamsBody{
			Directory: String("replay-dir"),
			Events: F([]SyncReplayParamsBodyEvent{
				{
					ID:          String("evt_001"),
					AggregateID: String("agg_001"),
					Seq:         Int(1),
					Type:        String("some.event"),
					Data:        F(map[string]any{"key": "value"}),
				},
			}),
		}),
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}
	if data == nil {
		t.Fatal("expected non-nil JSON output when Body.Present=true")
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	if m["directory"] != "replay-dir" {
		t.Errorf("expected directory=replay-dir, got %v", m["directory"])
	}
	events, ok := m["events"].([]any)
	if !ok || len(events) != 1 {
		t.Errorf("expected 1 event, got %v", m["events"])
	}
}

// TestSyncReplayParamsMarshalJSON_Empty verifies that when Body is absent,
// MarshalJSON returns nil (no body sent to the server).
func TestSyncReplayParamsMarshalJSON_Empty(t *testing.T) {
	t.Parallel()
	params := SyncReplayParams{
		Directory: String("dir"),
		Workspace: String("ws"),
		// Body not set — Present=false
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}
	if data != nil {
		t.Errorf("expected nil JSON output when Body.Present=false, got: %s", string(data))
	}
}

// TestSyncStealParamsMarshalJSON_WithBody verifies that when Body is present,
// MarshalJSON serializes the body correctly.
func TestSyncStealParamsMarshalJSON_WithBody(t *testing.T) {
	t.Parallel()
	params := SyncStealParams{
		Directory: String("dir"),
		Body: F(SyncStealParamsBody{
			SessionID: String("ses_abc123"),
		}),
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}
	if data == nil {
		t.Fatal("expected non-nil JSON output when Body.Present=true")
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	if m["sessionID"] != "ses_abc123" {
		t.Errorf("expected sessionID=ses_abc123, got %v", m["sessionID"])
	}
}

// TestSyncStealParamsMarshalJSON_Empty verifies that when Body is absent,
// MarshalJSON returns nil.
func TestSyncStealParamsMarshalJSON_Empty(t *testing.T) {
	t.Parallel()
	params := SyncStealParams{
		Directory: String("dir"),
		// Body not set — Present=false
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}
	if data != nil {
		t.Errorf("expected nil JSON output when Body.Present=false, got: %s", string(data))
	}
}

// TestSyncHistoryListParamsMarshalJSON_NonEmpty verifies that a non-empty map
// body is serialized correctly.
func TestSyncHistoryListParamsMarshalJSON_NonEmpty(t *testing.T) {
	t.Parallel()
	params := SyncHistoryListParams{
		Body: F(SyncHistoryListParamsBody{"agg_001": 10, "agg_002": 5}),
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}
	if data == nil {
		t.Fatal("expected non-nil JSON output for non-empty map body")
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}
	if len(m) != 2 {
		t.Errorf("expected 2 keys in body, got %d", len(m))
	}
}

// TestSyncHistoryListParamsMarshalJSON_Empty verifies that an empty map body
// returns nil (no body sent to the server).
func TestSyncHistoryListParamsMarshalJSON_Empty(t *testing.T) {
	t.Parallel()
	params := SyncHistoryListParams{
		// Body not set — nil map
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}
	if data != nil {
		t.Errorf("expected nil JSON output for empty map body, got: %s", string(data))
	}
}
