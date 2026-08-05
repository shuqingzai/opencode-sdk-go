package opencode

import (
	"encoding/json"
	"reflect"
	"testing"
)

// =============================================================================
// Task 1: IsKnown() fallthrough bug regression tests
// =============================================================================

// TestEventListResponseEventSessionNextPromptedDeliveryIsKnown verifies that
// all enum values for EventListResponseEventSessionNextPromptedDelivery return
// true from IsKnown(), and unknown values return false.
// Previously "steer" incorrectly returned false due to a switch fallthrough bug.
func TestEventListResponseEventSessionNextPromptedDeliveryIsKnown(t *testing.T) {
	t.Parallel()
	knownValues := []EventListResponseEventSessionNextPromptedDelivery{
		EventListResponseEventSessionNextPromptedDeliverySteer,
		EventListResponseEventSessionNextPromptedDeliveryQueue,
	}
	for _, v := range knownValues {
		if !v.IsKnown() {
			t.Errorf("IsKnown() returned false for known value %q", v)
		}
	}
	unknown := EventListResponseEventSessionNextPromptedDelivery("unknown")
	if unknown.IsKnown() {
		t.Errorf("IsKnown() returned true for unknown value %q", unknown)
	}
}

// TestEventListResponseEventSessionNextCompactionEndedReasonIsKnown verifies
// that all enum values for EventListResponseEventSessionNextCompactionEndedReason
// return true from IsKnown(), and unknown values return false.
// Previously "auto" incorrectly returned false due to a switch fallthrough bug.
func TestEventListResponseEventSessionNextCompactionEndedReasonIsKnown(t *testing.T) {
	t.Parallel()
	knownValues := []EventListResponseEventSessionNextCompactionEndedReason{
		EventListResponseEventSessionNextCompactionEndedReasonAuto,
		EventListResponseEventSessionNextCompactionEndedReasonManual,
	}
	for _, v := range knownValues {
		if !v.IsKnown() {
			t.Errorf("IsKnown() returned false for known value %q", v)
		}
	}
	unknown := EventListResponseEventSessionNextCompactionEndedReason("unknown")
	if unknown.IsKnown() {
		t.Errorf("IsKnown() returned true for unknown value %q", unknown)
	}
}

// =============================================================================
// Task 2: EventListResponseEventPermissionRepliedPropertiesReply tests
// =============================================================================

// TestEventListResponseEventPermissionRepliedPropertiesReplyIsKnown verifies
// that all OpenAPI enum values return true from IsKnown().
func TestEventListResponseEventPermissionRepliedPropertiesReplyIsKnown(t *testing.T) {
	t.Parallel()
	knownValues := []EventListResponseEventPermissionRepliedPropertiesReply{
		EventListResponseEventPermissionRepliedPropertiesReplyOnce,
		EventListResponseEventPermissionRepliedPropertiesReplyAlways,
		EventListResponseEventPermissionRepliedPropertiesReplyReject,
	}
	for _, v := range knownValues {
		if !v.IsKnown() {
			t.Errorf("IsKnown() returned false for known value %q", v)
		}
	}
	unknown := EventListResponseEventPermissionRepliedPropertiesReply("unknown")
	if unknown.IsKnown() {
		t.Errorf("IsKnown() returned true for unknown value %q", unknown)
	}
}

// TestEventListResponseEventPermissionRepliedDeserialization verifies that
// EventListResponseEventPermissionReplied deserializes correctly from a
// real-world JSON payload matching the OpenAPI schema.
func TestEventListResponseEventPermissionRepliedDeserialization(t *testing.T) {
	t.Parallel()
	raw := []byte(`{
		"id": "evt_perm001",
		"type": "permission.replied",
		"properties": {
			"sessionID": "ses_abc123",
			"requestID": "per_req456",
			"reply": "once"
		}
	}`)

	var props EventListResponseEventPermissionRepliedProperties
	if err := json.Unmarshal(raw, &props); err == nil {
		// Direct unmarshal into properties shouldn't work since the raw has outer keys
	}

	// Unmarshal the full event struct
	var event EventListResponseEventPermissionReplied
	// Extract properties object
	propsRaw := []byte(`{
		"sessionID": "ses_abc123",
		"requestID": "per_req456",
		"reply": "once"
	}`)
	if err := json.Unmarshal(propsRaw, &props); err != nil {
		t.Fatalf("failed to unmarshal properties: %v", err)
	}
	if props.SessionID != "ses_abc123" {
		t.Errorf("expected SessionID=ses_abc123, got %q", props.SessionID)
	}
	if props.RequestID != "per_req456" {
		t.Errorf("expected RequestID=per_req456, got %q", props.RequestID)
	}
	if props.Reply != EventListResponseEventPermissionRepliedPropertiesReplyOnce {
		t.Errorf("expected Reply=once, got %q", props.Reply)
	}
	if !props.Reply.IsKnown() {
		t.Errorf("Reply.IsKnown() returned false for 'once'")
	}

	// Unmarshal the full event
	fullRaw := []byte(`{"id":"evt_perm001","type":"permission.replied","properties":{"sessionID":"ses_abc123","requestID":"per_req456","reply":"always"}}`)
	if err := json.Unmarshal(fullRaw, &event); err != nil {
		t.Fatalf("failed to unmarshal EventListResponseEventPermissionReplied: %v", err)
	}
	if event.ID != "evt_perm001" {
		t.Errorf("expected ID=evt_perm001, got %q", event.ID)
	}
	if event.Properties.Reply != EventListResponseEventPermissionRepliedPropertiesReplyAlways {
		t.Errorf("expected Reply=always, got %q", event.Properties.Reply)
	}

	// Test "reject" value
	rejectRaw := []byte(`{"reply":"reject"}`)
	var rejectProps EventListResponseEventPermissionRepliedProperties
	if err := json.Unmarshal([]byte(`{"sessionID":"ses_x","requestID":"per_x","reply":"reject"}`), &rejectProps); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	_ = rejectRaw
	if rejectProps.Reply != EventListResponseEventPermissionRepliedPropertiesReplyReject {
		t.Errorf("expected Reply=reject, got %q", rejectProps.Reply)
	}
	if !rejectProps.Reply.IsKnown() {
		t.Errorf("Reply.IsKnown() returned false for 'reject'")
	}
}

// =============================================================================
// Task 3.1: EventListResponseEventSessionNextPromptAdmittedPropertiesDelivery
// =============================================================================

// TestEventListResponseEventSessionNextPromptAdmittedPropertiesDeliveryIsKnown
// verifies that all OpenAPI enum values return true from IsKnown().
func TestEventListResponseEventSessionNextPromptAdmittedPropertiesDeliveryIsKnown(t *testing.T) {
	t.Parallel()
	knownValues := []EventListResponseEventSessionNextPromptAdmittedPropertiesDelivery{
		EventListResponseEventSessionNextPromptAdmittedPropertiesDeliverySteer,
		EventListResponseEventSessionNextPromptAdmittedPropertiesDeliveryQueue,
	}
	for _, v := range knownValues {
		if !v.IsKnown() {
			t.Errorf("IsKnown() returned false for known value %q", v)
		}
	}
	unknown := EventListResponseEventSessionNextPromptAdmittedPropertiesDelivery("unknown")
	if unknown.IsKnown() {
		t.Errorf("IsKnown() returned true for unknown value %q", unknown)
	}
}

// =============================================================================
// Task 3.2 & 3.3: EventListResponseEventTuiToastShowProperties type checks
// =============================================================================

// TestEventListResponseEventTuiToastShowPropertiesVariantIsKnown verifies
// all OpenAPI enum values for the toast variant return true from IsKnown().
func TestEventListResponseEventTuiToastShowPropertiesVariantIsKnown(t *testing.T) {
	t.Parallel()
	knownValues := []EventListResponseEventTuiToastShowPropertiesVariant{
		EventListResponseEventTuiToastShowPropertiesVariantInfo,
		EventListResponseEventTuiToastShowPropertiesVariantSuccess,
		EventListResponseEventTuiToastShowPropertiesVariantWarning,
		EventListResponseEventTuiToastShowPropertiesVariantError,
	}
	for _, v := range knownValues {
		if !v.IsKnown() {
			t.Errorf("IsKnown() returned false for known value %q", v)
		}
	}
	unknown := EventListResponseEventTuiToastShowPropertiesVariant("unknown")
	if unknown.IsKnown() {
		t.Errorf("IsKnown() returned true for unknown value %q", unknown)
	}
}

// TestTuiToastShowVariantAliasIsCompatible verifies backward-compat alias works.
func TestTuiToastShowVariantAliasIsCompatible(t *testing.T) {
	t.Parallel()
	// TuiToastShowVariant is a type alias; ensure constants are identical
	var v TuiToastShowVariant = TuiToastShowVariantInfo
	var v2 EventListResponseEventTuiToastShowPropertiesVariant = EventListResponseEventTuiToastShowPropertiesVariantInfo
	if v != v2 {
		t.Errorf("alias mismatch: TuiToastShowVariantInfo=%q vs EventListResponseEventTuiToastShowPropertiesVariantInfo=%q", v, v2)
	}
}

// TestEventListResponseEventTuiToastShowPropertiesFieldTypes verifies that
// Duration and Title fields are concrete types (int64, string) not any.
func TestEventListResponseEventTuiToastShowPropertiesFieldTypes(t *testing.T) {
	t.Parallel()
	props := EventListResponseEventTuiToastShowProperties{}
	rt := reflect.TypeOf(props)

	durationField, ok := rt.FieldByName("Duration")
	if !ok {
		t.Fatal("Duration field not found")
	}
	if durationField.Type.Kind() != reflect.Int64 {
		t.Errorf("Duration field type = %v, want int64", durationField.Type)
	}

	titleField, ok := rt.FieldByName("Title")
	if !ok {
		t.Fatal("Title field not found")
	}
	if titleField.Type.Kind() != reflect.String {
		t.Errorf("Title field type = %v, want string", titleField.Type)
	}
}

// TestEventListResponseEventTuiToastShowPropertiesDeserialization verifies
// that the toast properties struct correctly deserializes from JSON including
// optional Duration and Title fields.
func TestEventListResponseEventTuiToastShowPropertiesDeserialization(t *testing.T) {
	t.Parallel()
	// With optional fields
	raw := []byte(`{"message":"hello","variant":"warning","duration":3000,"title":"My Title"}`)
	var props EventListResponseEventTuiToastShowProperties
	if err := json.Unmarshal(raw, &props); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if props.Message != "hello" {
		t.Errorf("expected Message=hello, got %q", props.Message)
	}
	if props.Variant != EventListResponseEventTuiToastShowPropertiesVariantWarning {
		t.Errorf("expected Variant=warning, got %q", props.Variant)
	}
	if props.Duration != 3000 {
		t.Errorf("expected Duration=3000, got %d", props.Duration)
	}
	if props.Title != "My Title" {
		t.Errorf("expected Title=My Title, got %q", props.Title)
	}

	// Without optional fields
	raw2 := []byte(`{"message":"test","variant":"error"}`)
	var props2 EventListResponseEventTuiToastShowProperties
	if err := json.Unmarshal(raw2, &props2); err != nil {
		t.Fatalf("failed to unmarshal without optional fields: %v", err)
	}
	if props2.Duration != 0 {
		t.Errorf("expected Duration=0 when absent, got %d", props2.Duration)
	}
	if props2.Title != "" {
		t.Errorf("expected Title= when absent, got %q", props2.Title)
	}
}

// TestEventListResponseIDPopulated verifies that EventListResponse.ID is populated
// after JSON unmarshaling. Previously, all 89 Event variant structs were missing
// the `ID` field, so the SSE-decoded EventListResponse.ID was always empty.
//
// This is a regression test for the critical blocker fixed in this batch.
func TestEventListResponseIDPopulated(t *testing.T) {
	t.Parallel()

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
func TestEventVariantIDFieldsExist(t *testing.T) {
	t.Parallel()
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

// TestSSEPropertiesTypedFields verifies the four SSE `properties` sub-fields
// that were previously `any` with a "runtime type of" comment but decoded to
// map[string]any (apijson cannot resolve registered unions through `any`).
//
// OpenAPI declares these as single `$ref`s (not anyOf), so the standard fix is
// to declare the field with the concrete type — no union routing needed.
//
// Run with: go test -run TestSSEPropertiesTypedFields -v ./...
func TestSSEPropertiesTypedFields(t *testing.T) {
	t.Parallel()

	t.Run("permission_asked_tool", func(t *testing.T) {
		t.Parallel()
		var e EventListResponse
		if err := json.Unmarshal([]byte(`{"type":"permission.asked","properties":{"id":"i","sessionID":"s","messageID":"m","callID":"c","title":"t","always":[],"patterns":[],"permission":"ask","tool":{"messageID":"mm","callID":"cc"}}}`), &e); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		p, ok := e.Properties.(EventListResponseEventPermissionAskedProperties)
		if !ok {
			t.Fatalf("Properties = %T, want EventListResponseEventPermissionAskedProperties", e.Properties)
		}
		if p.Tool == nil {
			t.Fatalf("Tool = nil, want non-nil")
		}
		if p.Tool.MessageID != "mm" {
			t.Errorf("Tool.MessageID = %q, want mm", p.Tool.MessageID)
		}
		if p.Tool.CallID != "cc" {
			t.Errorf("Tool.CallID = %q, want cc", p.Tool.CallID)
		}
	})

	t.Run("question_asked_tool", func(t *testing.T) {
		t.Parallel()
		var e EventListResponse
		if err := json.Unmarshal([]byte(`{"type":"question.asked","properties":{"id":"i","sessionID":"s","questions":[],"tool":{"messageID":"mm","callID":"cc"}}}`), &e); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		p, ok := e.Properties.(EventListResponseEventQuestionAskedProperties)
		if !ok {
			t.Fatalf("Properties = %T, want EventListResponseEventQuestionAskedProperties", e.Properties)
		}
		if p.Tool == nil {
			t.Fatalf("Tool = nil, want non-nil")
		}
		if p.Tool.MessageID != "mm" || p.Tool.CallID != "cc" {
			t.Errorf("Tool = %#v, want MessageID=mm CallID=cc", p.Tool)
		}
	})

	t.Run("permission_v2_asked_source", func(t *testing.T) {
		t.Parallel()
		var g GlobalEvent
		if err := json.Unmarshal([]byte(`{"type":"permission.v2.asked","payload":{"type":"permission.v2.asked","properties":{"id":"i","sessionID":"s","action":"open","resources":["/a"],"source":{"type":"tool","messageID":"m","callID":"c"}}}}`), &g); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		p, ok := g.Payload.(EventListResponseEventPermissionV2Asked)
		if !ok {
			t.Fatalf("Payload = %T, want EventListResponseEventPermissionV2Asked", g.Payload)
		}
		if p.Properties.Source == nil {
			t.Fatalf("Source = nil, want non-nil")
		}
		if p.Properties.Source.MessageID != "m" {
			t.Errorf("Source.MessageID = %q, want m", p.Properties.Source.MessageID)
		}
		if p.Properties.Source.CallID != "c" {
			t.Errorf("Source.CallID = %q, want c", p.Properties.Source.CallID)
		}
	})

	t.Run("question_v2_asked_tool", func(t *testing.T) {
		t.Parallel()
		var g GlobalEvent
		if err := json.Unmarshal([]byte(`{"type":"question.v2.asked","payload":{"type":"question.v2.asked","properties":{"id":"i","sessionID":"s","questions":[],"tool":{"messageID":"m","callID":"c"}}}}`), &g); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		p, ok := g.Payload.(EventListResponseEventQuestionV2Asked)
		if !ok {
			t.Fatalf("Payload = %T, want EventListResponseEventQuestionV2Asked", g.Payload)
		}
		if p.Properties.Tool == nil {
			t.Fatalf("Tool = nil, want non-nil")
		}
		if p.Properties.Tool.MessageID != "m" || p.Properties.Tool.CallID != "c" {
			t.Errorf("Tool = %#v, want MessageID=m CallID=c", p.Properties.Tool)
		}
	})

	t.Run("optional_field_absent", func(t *testing.T) {
		t.Parallel()
		var e EventListResponse
		if err := json.Unmarshal([]byte(`{"type":"permission.asked","properties":{"id":"i","sessionID":"s","messageID":"m","callID":"c","title":"t","always":[],"patterns":[],"permission":"ask"}}`), &e); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		p, ok := e.Properties.(EventListResponseEventPermissionAskedProperties)
		if !ok {
			t.Fatalf("Properties = %T", e.Properties)
		}
		if p.Tool != nil {
			t.Errorf("Tool = %#v, want nil (absent)", p.Tool)
		}
	})
}
