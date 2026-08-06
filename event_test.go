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
}
