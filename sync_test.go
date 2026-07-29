// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestSyncStartWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Sync.Start(context.TODO(), opencode.SyncStartParams{
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSyncReplay(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Sync.Replay(context.TODO(), opencode.SyncReplayParams{
		Body: opencode.F(opencode.SyncReplayParamsBody{
			Directory: opencode.F("directory"),
			Events: opencode.F([]opencode.SyncReplayParamsBodyEvent{
				{
					ID:          opencode.F("id"),
					AggregateID: opencode.F("aggregateID"),
					Seq:         opencode.F(int64(0)),
					Type:        opencode.F("type"),
					Data: opencode.F(map[string]any{
						"key": "value",
					}),
				},
			}),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSyncSteal(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Sync.Steal(context.TODO(), opencode.SyncStealParams{
		Body: opencode.F(opencode.SyncStealParamsBody{
			SessionID: opencode.F("sessionID"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSyncHistoryList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Sync.History.List(context.TODO(), opencode.SyncHistoryListParams{
		Body: opencode.F(opencode.SyncHistoryListParamsBody{
			"aggregate_id_1": 0,
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// --- Unit tests (no server required) ---

// TestSyncStartParamsURLQuery checks that query params serialize correctly.
func TestSyncStartParamsURLQuery(t *testing.T) {
	params := opencode.SyncStartParams{
		Directory: opencode.F("my-dir"),
		Workspace: opencode.F("my-ws"),
	}
	v := params.URLQuery()
	if v.Get("directory") != "my-dir" {
		t.Errorf("expected directory=my-dir, got %q", v.Get("directory"))
	}
	if v.Get("workspace") != "my-ws" {
		t.Errorf("expected workspace=my-ws, got %q", v.Get("workspace"))
	}
}

// TestSyncReplayParamsBodyMarshal verifies request body serialization.
func TestSyncReplayParamsBodyMarshal(t *testing.T) {
	body := opencode.SyncReplayParamsBody{
		Directory: opencode.F("dir"),
		Events: opencode.F([]opencode.SyncReplayParamsBodyEvent{
			{
				ID:          opencode.F("evt-1"),
				AggregateID: opencode.F("agg-1"),
				Seq:         opencode.F(int64(5)),
				Type:        opencode.F("session.created"),
				Data: opencode.F(map[string]any{
					"foo": "bar",
				}),
			},
		}),
	}
	data, err := body.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	// required fields must be present
	if _, ok := m["directory"]; !ok {
		t.Error("missing directory field")
	}
	if _, ok := m["events"]; !ok {
		t.Error("missing events field")
	}
}

// TestSyncStealParamsBodyMarshal verifies SyncStealParamsBody serialization.
func TestSyncStealParamsBodyMarshal(t *testing.T) {
	body := opencode.SyncStealParamsBody{
		SessionID: opencode.F("ses-123"),
	}
	data, err := body.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	if m["sessionID"] != "ses-123" {
		t.Errorf("expected sessionID=ses-123, got %v", m["sessionID"])
	}
}

// TestSyncHistoryListParamsBodyMarshal verifies map body serialization.
func TestSyncHistoryListParamsBodyMarshal(t *testing.T) {
	body := opencode.SyncHistoryListParamsBody{
		"agg-1": 10,
		"agg-2": 20,
	}
	data, err := body.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]int64
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	if m["agg-1"] != 10 {
		t.Errorf("expected agg-1=10, got %d", m["agg-1"])
	}
	if m["agg-2"] != 20 {
		t.Errorf("expected agg-2=20, got %d", m["agg-2"])
	}
}

// TestSyncHistoryEventUnmarshal verifies SyncHistoryEvent deserialization.
func TestSyncHistoryEventUnmarshal(t *testing.T) {
	raw := `{
		"id": "evt-abc",
		"aggregate_id": "agg-xyz",
		"seq": 42,
		"type": "session.created",
		"data": {"sessionID": "ses-1"},
		"extra_unknown_field": "ignored"
	}`
	var evt opencode.SyncHistoryEvent
	if err := evt.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if evt.ID != "evt-abc" {
		t.Errorf("expected ID=evt-abc, got %q", evt.ID)
	}
	if evt.AggregateID != "agg-xyz" {
		t.Errorf("expected AggregateID=agg-xyz, got %q", evt.AggregateID)
	}
	if evt.Seq != 42 {
		t.Errorf("expected Seq=42, got %d", evt.Seq)
	}
	if evt.Type != "session.created" {
		t.Errorf("expected Type=session.created, got %q", evt.Type)
	}
	if evt.Data == nil {
		t.Error("expected non-nil Data map")
	}
	// RawJSON must be populated
	if evt.JSON.RawJSON() == "" {
		t.Error("expected non-empty RawJSON")
	}
}

// TestSyncReplayResponseUnmarshal verifies SyncReplayResponse deserialization.
func TestSyncReplayResponseUnmarshal(t *testing.T) {
	raw := `{"sessionID": "ses-replay-1"}`
	var resp opencode.SyncReplayResponse
	if err := resp.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if resp.SessionID != "ses-replay-1" {
		t.Errorf("expected ses-replay-1, got %q", resp.SessionID)
	}
	if resp.JSON.RawJSON() == "" {
		t.Error("expected non-empty RawJSON")
	}
}

// TestSyncStealResponseUnmarshal verifies SyncStealResponse deserialization.
func TestSyncStealResponseUnmarshal(t *testing.T) {
	raw := `{"sessionID": "ses-steal-2"}`
	var resp opencode.SyncStealResponse
	if err := resp.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if resp.SessionID != "ses-steal-2" {
		t.Errorf("expected ses-steal-2, got %q", resp.SessionID)
	}
}

// TestSyncHistoryEventMissingOptionalFieldsOK ensures missing optional fields
// don't cause unmarshal errors.
func TestSyncHistoryEventMissingOptionalFieldsOK(t *testing.T) {
	// Only required fields
	raw := `{
		"id": "evt-1",
		"aggregate_id": "agg-1",
		"seq": 1,
		"type": "test.event",
		"data": {}
	}`
	var evt opencode.SyncHistoryEvent
	if err := evt.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if evt.ID != "evt-1" {
		t.Errorf("expected evt-1, got %q", evt.ID)
	}
}
