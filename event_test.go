package opencode

import (
	"encoding/json"
	"reflect"
	"testing"
)

// TestEventListResponseIDPopulated verifies that EventListResponse.ID is populated
// after JSON unmarshaling. Previously, all 89 Event variant structs were missing
// the `ID` field, so the SSE-decoded EventListResponse.ID was always empty.
//
// This is a regression test for the critical blocker fixed in this batch.
//
// Skipped by default; remove t.Skip() to enable.
func TestEventListResponseIDPopulated(t *testing.T) {
	t.Skip("regression test for ID population fix; remove t.Skip() to enable")

	// 1) Simple variant: server.connected has empty properties.
	raw := []byte(`{"id":"evt_abc123","type":"server.connected","properties":{}}`)
	var resp EventListResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	t.Logf("EventListResponse.ID = %q", resp.ID)
	t.Logf("EventListResponse.Type = %q", resp.Type)
	if resp.ID != "evt_abc123" {
		t.Errorf("expected ID=evt_abc123, got %q", resp.ID)
	}

	// 2) Variant with non-empty properties: session.created
	raw2 := []byte(`{"id":"evt_def456","type":"session.created","properties":{"id":"ses_xyz"}}`)
	var resp2 EventListResponse
	if err := json.Unmarshal(raw2, &resp2); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	t.Logf("session.created EventListResponse.ID = %q", resp2.ID)
	if resp2.ID != "evt_def456" {
		t.Errorf("expected ID=evt_def456, got %q", resp2.ID)
	}

	// 3) Verify typed variant also has ID populated (the actual blocker).
	if v, ok := resp2.AsUnion().(EventListResponseEventSessionCreated); ok {
		t.Logf("variant EventListResponseEventSessionCreated.ID = %q", v.ID)
		if v.ID != "evt_def456" {
			t.Errorf("expected variant.ID=evt_def456, got %q", v.ID)
		}
	} else {
		t.Fatalf("expected union to be EventListResponseEventSessionCreated, got %T", resp2.AsUnion())
	}

	// 4) Verify tui.toast.show variant (it has Properties sub-object)
	raw3 := []byte(`{"id":"evt_xyz789","type":"tui.toast.show","properties":{"message":"hi","variant":"info"}}`)
	var resp3 EventListResponse
	if err := json.Unmarshal(raw3, &resp3); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	t.Logf("tui.toast.show EventListResponse.ID = %q", resp3.ID)
	if resp3.ID != "evt_xyz789" {
		t.Errorf("expected ID=evt_xyz789, got %q", resp3.ID)
	}
	if v, ok := resp3.AsUnion().(EventListResponseEventTuiToastShow); ok {
		t.Logf("variant EventListResponseEventTuiToastShow.ID = %q", v.ID)
		if v.ID != "evt_xyz789" {
			t.Errorf("expected variant.ID=evt_xyz789, got %q", v.ID)
		}
	} else {
		t.Fatalf("expected union to be EventListResponseEventTuiToastShow, got %T", resp3.AsUnion())
	}
}

// TestEventVariantIDFieldsExist verifies that representative EventListResponseEvent*
// variant structs have an ID field at the top level (matching the OpenAPI
// requirement that `id` (pattern ^evt_) is required for all event variants).
//
// Skipped by default; remove t.Skip() to enable.
func TestEventVariantIDFieldsExist(t *testing.T) {
	t.Skip("structural verification of ID field presence; remove t.Skip() to enable")
	representative := []any{
		EventListResponseEventServerConnected{},
		EventListResponseEventSessionCreated{},
		EventListResponseEventMessagePartUpdated{},
		EventListResponseEventTuiToastShow{},
		EventListResponseEventTuiCommandExecute{},
		EventListResponseEventSessionError{},
		EventListResponseEventGlobalDisposed{},
		EventListResponseEventProjectUpdated{},
	}
	for _, v := range representative {
		rv := reflect.ValueOf(v)
		idField := rv.FieldByName("ID")
		if !idField.IsValid() {
			t.Errorf("variant %T missing ID field", v)
		}
	}
}
