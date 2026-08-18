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
	rt := reflect.TypeFor[EventListResponseEventTuiToastShowProperties]()

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
		if p.JSON.Tool.IsMissing() {
			t.Fatalf("Tool is missing, want present")
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
		if p.JSON.Tool.IsMissing() {
			t.Fatalf("Tool is missing, want present")
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
		if p.Properties.JSON.Source.IsMissing() {
			t.Fatalf("Source is missing, want present")
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
		if p.Properties.JSON.Tool.IsMissing() {
			t.Fatalf("Tool is missing, want present")
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
		if !p.JSON.Tool.IsMissing() {
			t.Errorf("Tool = %#v, want missing (absent)", p.Tool)
		}
	})
}

// =============================================================================
// Deprecated type-alias SSE decoding regression tests
//
// These events previously declared their own dedicated struct types
// (EventListResponseEventSessionNextModelSwitchedModel,
// EventListResponseEventSessionNextPromptedPrompt,
// EventListResponseEventSessionNextPromptAdmittedPropertiesPrompt,
// SyncEventPrompt, EventListResponseEventPermissionV2RepliedPropertiesReply).
// They were converted to `// Deprecated:` type aliases of the canonical
// cross-service types (ModelRef, V2SessionInputPrompt, PermissionV2Reply).
// These tests prove the SSE deserialization path still decodes non-zero values.
// =============================================================================

// TestSSEAliasedModelSwitchedDecodes verifies session.next.model.switched
// decodes the model field (aliased to ModelRef) to concrete values.
func TestSSEAliasedModelSwitchedDecodes(t *testing.T) {
	t.Parallel()
	var e EventListResponse
	data := `{"id":"evt_1","type":"session.next.model.switched","properties":{"timestamp":123,"sessionID":"ses_1","messageID":"msg_1","model":{"id":"m-1","providerID":"p-1","variant":"v1"}}}`
	if err := json.Unmarshal([]byte(data), &e); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if e.Type != EventListResponseTypeSessionNextModelSwitched {
		t.Fatalf("Type = %q", e.Type)
	}
	p, ok := e.Properties.(EventListResponseEventSessionNextModelSwitchedProperties)
	if !ok {
		t.Fatalf("Properties = %T, want EventListResponseEventSessionNextModelSwitchedProperties", e.Properties)
	}
	if p.Model.ID != "m-1" || p.Model.ProviderID != "p-1" || p.Model.Variant != "v1" {
		t.Fatalf("Model = %+v, want non-zero decoded values", p.Model)
	}
}

// TestSSEAliasedPromptedPromptDecodes verifies session.next.prompted decodes
// the prompt field (aliased to V2SessionInputPrompt) with nested attachments.
func TestSSEAliasedPromptedPromptDecodes(t *testing.T) {
	t.Parallel()
	var e EventListResponse
	data := `{"id":"evt_2","type":"session.next.prompted","properties":{"timestamp":1,"sessionID":"ses_1","messageID":"msg_2","prompt":{"text":"hello","files":[{"uri":"file:///a.txt","mime":"text/plain","name":"a.txt"}],"agents":[{"name":"writer"}]},"delivery":"steer"}}`
	if err := json.Unmarshal([]byte(data), &e); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	p, ok := e.Properties.(EventListResponseEventSessionNextPromptedProperties)
	if !ok {
		t.Fatalf("Properties = %T, want EventListResponseEventSessionNextPromptedProperties", e.Properties)
	}
	if p.Prompt.Text != "hello" {
		t.Fatalf("Prompt.Text = %q, want hello", p.Prompt.Text)
	}
	if len(p.Prompt.Files) != 1 || p.Prompt.Files[0].URI != "file:///a.txt" || p.Prompt.Files[0].Mime != "text/plain" {
		t.Fatalf("Prompt.Files = %+v", p.Prompt.Files)
	}
	if len(p.Prompt.Agents) != 1 || p.Prompt.Agents[0].Name != "writer" {
		t.Fatalf("Prompt.Agents = %+v", p.Prompt.Agents)
	}
}

// TestSSEAliasedPromptAdmittedDecodes verifies session.next.prompt.admitted
// decodes the prompt field (aliased to V2SessionInputPrompt).
func TestSSEAliasedPromptAdmittedDecodes(t *testing.T) {
	t.Parallel()
	var e EventListResponse
	data := `{"id":"evt_3","type":"session.next.prompt.admitted","properties":{"timestamp":1,"sessionID":"ses_1","messageID":"msg_3","prompt":{"text":"admitted","agents":[{"name":"reader"}]},"delivery":"queue"}}`
	if err := json.Unmarshal([]byte(data), &e); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	p, ok := e.Properties.(EventListResponseEventSessionNextPromptAdmittedProperties)
	if !ok {
		t.Fatalf("Properties = %T, want EventListResponseEventSessionNextPromptAdmittedProperties", e.Properties)
	}
	if p.Prompt.Text != "admitted" || len(p.Prompt.Agents) != 1 {
		t.Fatalf("Prompt = %+v, want non-zero decoded values", p.Prompt)
	}
}

// TestSSEAliasedPermissionV2RepliedDecodes verifies permission.v2.replied
// decodes the reply field (aliased to PermissionV2Reply) and that both the new
// and deprecated constant names resolve to the same value.
func TestSSEAliasedPermissionV2RepliedDecodes(t *testing.T) {
	t.Parallel()
	var e EventListResponse
	data := `{"id":"evt_4","type":"permission.v2.replied","properties":{"sessionID":"ses_1","requestID":"per_1","reply":"once"}}`
	if err := json.Unmarshal([]byte(data), &e); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	p, ok := e.Properties.(EventListResponseEventPermissionV2RepliedProperties)
	if !ok {
		t.Fatalf("Properties = %T, want EventListResponseEventPermissionV2RepliedProperties", e.Properties)
	}
	if p.Reply != PermissionV2ReplyOnce {
		t.Fatalf("Reply = %q, want once", p.Reply)
	}
	if !p.Reply.IsKnown() {
		t.Fatalf("Reply.IsKnown() = false for %q", p.Reply)
	}
	// Deprecated constant aliases must still resolve to the canonical values.
	if EventListResponseEventPermissionV2RepliedPropertiesReplyOnce != PermissionV2ReplyOnce {
		t.Fatalf("deprecated const Once mismatch")
	}
	if EventListResponseEventPermissionV2RepliedPropertiesReplyAlways != PermissionV2ReplyAlways {
		t.Fatalf("deprecated const Always mismatch")
	}
	if EventListResponseEventPermissionV2RepliedPropertiesReplyReject != PermissionV2ReplyReject {
		t.Fatalf("deprecated const Reject mismatch")
	}
}

// TestSSEAliasedSyncEventPromptDecodes verifies the V1 sync path decodes a
// prompt.admitted sync event whose prompt field uses V2SessionInputPrompt.
func TestSSEAliasedSyncEventPromptDecodes(t *testing.T) {
	t.Parallel()
	var sync1 SyncEventResponse
	data := `{"type":"sync","id":"evt_5","syncEvent":{"type":"session.next.prompt.admitted.1","id":"evt_5","seq":1,"aggregateID":"ses_1","data":{"timestamp":1,"sessionID":"ses_1","messageID":"msg_5","prompt":{"text":"sync-prompt","files":[{"uri":"file:///b.txt","mime":"text/plain"}]},"delivery":"steer"}}}`
	if err := json.Unmarshal([]byte(data), &sync1); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if sync1.Type != SyncEventResponseTypeSync {
		t.Fatalf("Type = %q", sync1.Type)
	}
	inner, ok := sync1.SyncEvent.AsUnion().(SyncEventSessionNextPromptAdmitted)
	if !ok {
		t.Fatalf("SyncEvent.AsUnion() = %T, want SyncEventSessionNextPromptAdmitted", sync1.SyncEvent.AsUnion())
	}
	if inner.Data.Prompt.Text != "sync-prompt" {
		t.Fatalf("Prompt.Text = %q, want sync-prompt", inner.Data.Prompt.Text)
	}
	if len(inner.Data.Prompt.Files) != 1 || inner.Data.Prompt.Files[0].URI != "file:///b.txt" {
		t.Fatalf("Prompt.Files = %+v", inner.Data.Prompt.Files)
	}
}

// TestDeprecatedAliasTypeIdentity verifies the deprecated aliases are true type
// aliases of the canonical cross-service types (identical, not distinct types).
func TestDeprecatedAliasTypeIdentity(t *testing.T) {
	t.Parallel()
	if reflect.TypeFor[EventListResponseEventSessionNextModelSwitchedModel]() != reflect.TypeFor[ModelRef]() {
		t.Fatalf("EventListResponseEventSessionNextModelSwitchedModel is not an alias of ModelRef")
	}
	if reflect.TypeFor[EventListResponseEventSessionNextPromptedPrompt]() != reflect.TypeFor[V2SessionInputPrompt]() {
		t.Fatalf("EventListResponseEventSessionNextPromptedPrompt is not an alias of V2SessionInputPrompt")
	}
	if reflect.TypeFor[EventListResponseEventSessionNextPromptAdmittedPropertiesPrompt]() != reflect.TypeFor[V2SessionInputPrompt]() {
		t.Fatalf("EventListResponseEventSessionNextPromptAdmittedPropertiesPrompt is not an alias of V2SessionInputPrompt")
	}
	if reflect.TypeFor[SyncEventPrompt]() != reflect.TypeFor[V2SessionInputPrompt]() {
		t.Fatalf("SyncEventPrompt is not an alias of V2SessionInputPrompt")
	}
	if reflect.TypeFor[EventListResponseEventPermissionV2RepliedPropertiesReply]() != reflect.TypeFor[PermissionV2Reply]() {
		t.Fatalf("EventListResponseEventPermissionV2RepliedPropertiesReply is not an alias of PermissionV2Reply")
	}
	if reflect.TypeFor[EventListResponseEventQuestionAskedPropertiesQuestions]() != reflect.TypeFor[QuestionInfo]() {
		t.Fatalf("EventListResponseEventQuestionAskedPropertiesQuestions is not an alias of QuestionInfo")
	}
	if reflect.TypeFor[EventListResponseEventQuestionAskedPropertiesTool]() != reflect.TypeFor[QuestionTool]() {
		t.Fatalf("EventListResponseEventQuestionAskedPropertiesTool is not an alias of QuestionTool")
	}
	if reflect.TypeFor[EventListResponseEventPermissionV2AskedPropertiesSource]() != reflect.TypeFor[PermissionV2Source]() {
		t.Fatalf("EventListResponseEventPermissionV2AskedPropertiesSource is not an alias of PermissionV2Source")
	}
	if reflect.TypeFor[EventListResponseEventSessionNextMovedPropertiesLocation]() != reflect.TypeFor[LocationRef]() {
		t.Fatalf("EventListResponseEventSessionNextMovedPropertiesLocation is not an alias of LocationRef")
	}
	if reflect.TypeFor[EventListResponseEventSessionNextRevertStagedPropertiesRevert]() != reflect.TypeFor[RevertState]() {
		t.Fatalf("EventListResponseEventSessionNextRevertStagedPropertiesRevert is not an alias of RevertState")
	}
}

// =============================================================================
// Task 4 (A1 阻塞项 1/高 2 修复回归): session.next.revert.staged 的 revert 字段与
// session.next.moved 的 location 字段必须直接引用共享 RevertState/LocationRef，
// 且 FileDiff.path 字段不得再被静默丢弃（此前 []VcsFileDiff 用 "file" key 承接
// "path" 数据导致丢失，实测 Files[0].File==""）。
// =============================================================================

// TestEventListResponseSessionNextRevertStagedFilesPathRegression 是 A1 🔴 阻塞项
// 1 的回归护栏：修复前 revert.files[0].path 会被静默丢弃
// (EventListResponseEventSessionNextRevertStagedPropertiesRevert.Files 曾是
// []VcsFileDiff，其 json key 是 "file" 而非 OpenAPI FileDiff 的 "path")。
// 使用 OpenAPI 真实 wire 格式反序列化 EventListResponseEventSessionNextRevertStagedProperties
// (对应 /event 链路)，断言 Files[0].Path 保留原始路径。
func TestEventListResponseSessionNextRevertStagedFilesPathRegression(t *testing.T) {
	t.Parallel()
	raw := `{"timestamp":1700000000,"sessionID":"ses_1","revert":{"messageID":"msg_1","partID":"prt_1","snapshot":"snap_1","diff":"diff text","files":[{"path":"/a/b.txt","status":"added","additions":2,"deletions":0,"patch":"@@"}]}}`
	var props EventListResponseEventSessionNextRevertStagedProperties
	if err := json.Unmarshal([]byte(raw), &props); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if props.Revert.MessageID != "msg_1" {
		t.Errorf("Revert.MessageID = %q, want msg_1", props.Revert.MessageID)
	}
	if len(props.Revert.Files) != 1 {
		t.Fatalf("Revert.Files len = %d, want 1", len(props.Revert.Files))
	}
	// 🔴 回归护栏：修复前该字段为 ""（path 被静默丢弃）。
	if got := props.Revert.Files[0].Path; got != "/a/b.txt" {
		t.Errorf("Revert.Files[0].Path = %q, want \"/a/b.txt\" (pre-fix regression: silently dropped to \"\")", got)
	}
	if props.Revert.Files[0].Status != FileDiffStatusAdded {
		t.Errorf("Revert.Files[0].Status = %q, want added", props.Revert.Files[0].Status)
	}
	if props.Revert.Files[0].Additions != 2 {
		t.Errorf("Revert.Files[0].Additions = %d, want 2", props.Revert.Files[0].Additions)
	}
	if props.Revert.Files[0].JSON.Path.IsMissing() {
		t.Error("Revert.Files[0].JSON.Path reported missing, want present")
	}
	// Revert 字段本身是共享 RevertState 的直接引用（非重复克隆类型）。
	var _ RevertState = props.Revert
}

// TestEventListResponseSessionNextMovedLocationFields 是 A1 🟠 高 2 的回归护栏：
// location 字段必须直接引用共享 LocationRef，Directory/WorkspaceID 正确解析。
func TestEventListResponseSessionNextMovedLocationFields(t *testing.T) {
	t.Parallel()
	raw := `{"timestamp":1700000001,"sessionID":"ses_2","location":{"directory":"/repo/workspace","workspaceID":"wrk_1"},"subdirectory":"sub/dir"}`
	var props EventListResponseEventSessionNextMovedProperties
	if err := json.Unmarshal([]byte(raw), &props); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if props.Location.Directory != "/repo/workspace" {
		t.Errorf("Location.Directory = %q, want /repo/workspace", props.Location.Directory)
	}
	if props.Location.WorkspaceID != "wrk_1" {
		t.Errorf("Location.WorkspaceID = %q, want wrk_1", props.Location.WorkspaceID)
	}
	if props.Subdirectory != "sub/dir" {
		t.Errorf("Subdirectory = %q, want sub/dir", props.Subdirectory)
	}
	// Location 字段本身是共享 LocationRef 的直接引用（非重复克隆类型）。
	var _ LocationRef = props.Location
}

// TestEventListResponseSessionNextRevertStagedNoDuplicateFileDiffType 确认
// event_global_types.go 不再重复定义 Files 字段类型（[]VcsFileDiff 的错配已被
// 移除），Files 元素类型必须与共享 FileDiff 完全一致（reflect 恒等）。
func TestEventListResponseSessionNextRevertStagedNoDuplicateFileDiffType(t *testing.T) {
	t.Parallel()
	var props EventListResponseEventSessionNextRevertStagedProperties
	filesField := reflect.TypeOf(props.Revert.Files)
	wantElem := reflect.TypeFor[FileDiff]()
	if filesField.Elem() != wantElem {
		t.Fatalf("Revert.Files element type = %s, want %s", filesField.Elem(), wantElem)
	}
}

// =============================================================================
// Q2 独立复核发现项 1（🔴）：todo.updated 的 todos[] 字段重复定义 Todo 且
// 私造了 OpenAPI/JS SDK(v2) 均不存在的 "id" 字段。
//
// OpenAPI `EventTodoUpdated.properties.properties.todos.items` 是
// `$ref: #/components/schemas/Todo`，且 `Todo` schema 只有
// {content, status, priority} 三个字段（无 id，additionalProperties:false）；
// JS SDK(v2) `types.gen.ts` 的 `Todo` 类型同样只有这三个字段。修复前 Go 侧
// 造了一个本地类型 EventListResponseEventTodoUpdatedPropertiesTodo，凭空加了
// 一个 `ID string json:"id,required"` 字段——服务端 wire 上从不会有这个 key，
// 该字段永远是空字符串，属于无中生有的伪造字段（比 RevertState/LocationRef
// 的“重复定义但字段一致”更严重：这里字段集合本身就与三源不符）。
// 修复：改为直接引用 session.go 已有的共享 Todo 类型（type alias 保留旧名，
// 遵循本轮 RevertState/LocationRef 同款 Deprecated 迁移模式）。
// =============================================================================

// TestEventListResponseTodoUpdatedReusesSharedTodoType 验证
// EventListResponseEventTodoUpdatedPropertiesTodo 是共享 [Todo] 类型的真别名
// （reflect.TypeFor 恒等），而不是重新造的本地类型。
func TestEventListResponseTodoUpdatedReusesSharedTodoType(t *testing.T) {
	t.Parallel()
	if reflect.TypeFor[EventListResponseEventTodoUpdatedPropertiesTodo]() != reflect.TypeFor[Todo]() {
		t.Fatalf("EventListResponseEventTodoUpdatedPropertiesTodo is not an alias of Todo")
	}
	var props EventListResponseEventTodoUpdatedProperties
	todosField := reflect.TypeOf(props.Todos)
	wantElem := reflect.TypeFor[Todo]()
	if todosField.Elem() != wantElem {
		t.Fatalf("Todos element type = %s, want %s", todosField.Elem(), wantElem)
	}
}

// TestEventListResponseTodoUpdatedDeserialization 用真实 wire 格式
// （content/status/priority，无 id）反序列化，断言三个字段都被正确解析且
// JSON 元数据完整；同时确认没有任何 "id" 字段被 apijson 期待（多余的
// "id" key 若出现在 wire 上会被静默放进 ExtraFields，而不是覆盖一个不存在
// 于 OpenAPI/JS 的 struct 字段）。
func TestEventListResponseTodoUpdatedDeserialization(t *testing.T) {
	t.Parallel()
	raw := `{"sessionID":"ses_1","todos":[{"content":"write tests","status":"in_progress","priority":"high"}]}`
	var props EventListResponseEventTodoUpdatedProperties
	if err := json.Unmarshal([]byte(raw), &props); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if props.SessionID != "ses_1" {
		t.Errorf("SessionID = %q, want ses_1", props.SessionID)
	}
	if len(props.Todos) != 1 {
		t.Fatalf("Todos len = %d, want 1", len(props.Todos))
	}
	got := props.Todos[0]
	if got.Content != "write tests" || got.Status != "in_progress" || got.Priority != "high" {
		t.Errorf("Todos[0] = %+v, want {write tests, in_progress, high}", got)
	}
	if got.JSON.Content.IsMissing() || got.JSON.Status.IsMissing() || got.JSON.Priority.IsMissing() {
		t.Error("Todos[0].JSON metadata reports fields missing, want present")
	}
	if got.JSON.RawJSON() == "" {
		t.Error("Todos[0].JSON.RawJSON() is empty, want original wire JSON preserved")
	}
}

// TestEventListResponseTodoUpdatedExtraIDGoesToExtraFields 若服务端 wire 上
// 出现一个多余的 "id" key（例如未来 API 扩展），修复后的共享 Todo 类型必须
// 把它放进 ExtraFields（因为 Todo 没有导出的 ID 字段），而不是像修复前那样
// 悄悄把 wire 上从来不存在的 "id" 映射成一个永远为空的字段。
func TestEventListResponseTodoUpdatedExtraIDGoesToExtraFields(t *testing.T) {
	t.Parallel()
	raw := `{"sessionID":"ses_2","todos":[{"id":"todo_999","content":"c","status":"pending","priority":"low"}]}`
	var props EventListResponseEventTodoUpdatedProperties
	if err := json.Unmarshal([]byte(raw), &props); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if len(props.Todos) != 1 {
		t.Fatalf("Todos len = %d, want 1", len(props.Todos))
	}
	if _, ok := props.Todos[0].JSON.ExtraFields["id"]; !ok {
		t.Error(`Todos[0].JSON.ExtraFields["id"] missing, want the unknown "id" key captured there`)
	}
}
