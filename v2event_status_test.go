package opencode

import (
	"encoding/json"
	"reflect"
	"testing"
)

// OpenAPI evidence (components/schemas/SessionStatus, anyOf):
//
//	{ type: "idle" }                                              required: type
//	{ type: "retry", attempt, message, action?, next }             required: type, attempt, message, next
//	{ type: "busy" }                                              required: type
//
// JS SDK v2 evidence (gen/types.gen.ts:673 `export type SessionStatus`):
//
//	| { type: "idle" }
//	| { type: "retry"; attempt: number; message: string; action?: {...}; next: number }
//	| { type: "busy" }
//
// Carriers of that union:
//   - session.status.data      -> V2EventSessionStatusData             (v2event.go)
//   - EventSessionStatus.properties -> EventListResponseEventSessionStatusProperties (event.go)
//
// Both carry a sibling `sessionID` field, so they follow the Case-B Response Union
// pattern (regular fields decoded from the root, union decoded from the `status`
// sub-JSON) — identical to V2EventSessionErrorData.

// TestV2EventSessionStatusDataUnionContract locks the Response Union 铁律 shape of
// [V2EventSessionStatusData]: the union carrier field must be `any` (never the
// SessionStatus interface), and an unexported `union` field must exist without a
// json tag.
func TestV2EventSessionStatusDataUnionContract(t *testing.T) {
	rt := reflect.TypeFor[V2EventSessionStatusData]()

	status, ok := rt.FieldByName("Status")
	if !ok {
		t.Fatal("V2EventSessionStatusData.Status field missing")
	}
	if status.Type.Kind() != reflect.Interface || status.Type.NumMethod() != 0 {
		t.Errorf("Status must be `any` (Response Union carrier), got %s", status.Type)
	}
	if got := status.Tag.Get("json"); got != "status,required" {
		t.Errorf("Status json tag = %q, want %q", got, "status,required")
	}

	union, ok := rt.FieldByName("union")
	if !ok {
		t.Fatal("V2EventSessionStatusData must declare an unexported `union` field")
	}
	if union.Type != reflect.TypeFor[SessionStatus]() {
		t.Errorf("union field type = %s, want SessionStatus", union.Type)
	}
	if tag, hasTag := union.Tag.Lookup("json"); hasTag {
		t.Errorf("union field must not carry a json tag, got %q", tag)
	}
	if union.IsExported() {
		t.Error("union field must be unexported")
	}
}

// TestV2EventSessionStatusDataUnmarshal verifies every SessionStatus variant
// decodes while the sibling `sessionID` field is preserved — the regression the
// old `Status SessionStatus` interface-typed field could not guarantee.
func TestV2EventSessionStatusDataUnmarshal(t *testing.T) {
	tests := []struct {
		name   string
		raw    string
		verify func(t *testing.T, d V2EventSessionStatusData)
	}{
		{
			name: "idle",
			raw:  `{"sessionID":"ses_idle","status":{"type":"idle"}}`,
			verify: func(t *testing.T, d V2EventSessionStatusData) {
				idle, ok := d.AsStatus().(SessionStatusIdle)
				if !ok {
					t.Fatalf("AsStatus() = %T, want SessionStatusIdle", d.AsStatus())
				}
				if idle.Type != "idle" {
					t.Errorf("Type = %q, want %q", idle.Type, "idle")
				}
				if idle.JSON.RawJSON() != `{"type":"idle"}` {
					t.Errorf("variant RawJSON = %q", idle.JSON.RawJSON())
				}
			},
		},
		{
			name: "busy",
			raw:  `{"sessionID":"ses_busy","status":{"type":"busy"}}`,
			verify: func(t *testing.T, d V2EventSessionStatusData) {
				busy, ok := d.AsStatus().(SessionStatusBusy)
				if !ok {
					t.Fatalf("AsStatus() = %T, want SessionStatusBusy", d.AsStatus())
				}
				if busy.Type != "busy" {
					t.Errorf("Type = %q, want %q", busy.Type, "busy")
				}
			},
		},
		{
			name: "retry_with_action",
			raw: `{"sessionID":"ses_retry","status":{"type":"retry","attempt":3,"message":"rate limited",` +
				`"next":1700000000,"action":{"reason":"quota","provider":"anthropic","title":"Upgrade",` +
				`"message":"Out of credits","label":"Buy","link":"https://example.com"}}}`,
			verify: func(t *testing.T, d V2EventSessionStatusData) {
				retry, ok := d.AsStatus().(SessionStatusRetry)
				if !ok {
					t.Fatalf("AsStatus() = %T, want SessionStatusRetry", d.AsStatus())
				}
				if retry.Type != "retry" {
					t.Errorf("Type = %q, want %q", retry.Type, "retry")
				}
				if retry.Attempt != 3 {
					t.Errorf("Attempt = %d, want 3", retry.Attempt)
				}
				if retry.Message != "rate limited" {
					t.Errorf("Message = %q", retry.Message)
				}
				if retry.Next != 1700000000 {
					t.Errorf("Next = %d, want 1700000000", retry.Next)
				}
				if retry.Action.Reason != "quota" || retry.Action.Provider != "anthropic" ||
					retry.Action.Title != "Upgrade" || retry.Action.Message != "Out of credits" ||
					retry.Action.Label != "Buy" || retry.Action.Link != "https://example.com" {
					t.Errorf("Action = %+v", retry.Action)
				}
			},
		},
		{
			// `action` is optional per OpenAPI (not in the retry variant's required list).
			name: "retry_without_action",
			raw:  `{"sessionID":"ses_retry2","status":{"type":"retry","attempt":1,"message":"m","next":42}}`,
			verify: func(t *testing.T, d V2EventSessionStatusData) {
				retry, ok := d.AsStatus().(SessionStatusRetry)
				if !ok {
					t.Fatalf("AsStatus() = %T, want SessionStatusRetry", d.AsStatus())
				}
				if retry.Attempt != 1 || retry.Next != 42 {
					t.Errorf("Attempt/Next = %d/%d, want 1/42", retry.Attempt, retry.Next)
				}
				if retry.Action.Reason != "" || retry.Action.Provider != "" || retry.Action.Title != "" ||
					retry.Action.Message != "" || retry.Action.Label != "" || retry.Action.Link != "" {
					t.Errorf("Action should be zero when omitted, got %+v", retry.Action)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var d V2EventSessionStatusData
			if err := json.Unmarshal([]byte(tc.raw), &d); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}

			wantSessionID := map[string]string{
				"idle":                 "ses_idle",
				"busy":                 "ses_busy",
				"retry_with_action":    "ses_retry",
				"retry_without_action": "ses_retry2",
			}[tc.name]
			if d.SessionID != wantSessionID {
				t.Errorf("SessionID = %q, want %q (sibling field must survive union decode)", d.SessionID, wantSessionID)
			}

			// Status (the `any` carrier) and AsStatus() must agree.
			if !reflect.DeepEqual(d.Status, d.AsStatus()) {
				t.Errorf("Status (%T) and AsStatus() (%T) diverge", d.Status, d.AsStatus())
			}
			if d.JSON.RawJSON() != tc.raw {
				t.Errorf("RawJSON not preserved:\n got %s\nwant %s", d.JSON.RawJSON(), tc.raw)
			}
			tc.verify(t, d)
		})
	}
}

// TestV2EventSessionStatusEndToEnd decodes a complete `session.status` V2 SSE
// event through the V2EventPayloadUnion, mirroring OpenAPI schema `session.status`.
func TestV2EventSessionStatusEndToEnd(t *testing.T) {
	raw := `{"id":"evt_1","type":"session.status",` +
		`"durable":{"aggregateID":"ses_e2e","seq":7,"version":1},` +
		`"metadata":{"k":"v"},` +
		`"data":{"sessionID":"ses_e2e","status":{"type":"retry","attempt":2,"message":"m","next":9}}}`

	var ev V2Event
	if err := json.Unmarshal([]byte(raw), &ev); err != nil {
		t.Fatalf("unmarshal V2Event: %v", err)
	}
	if ev.ID != "evt_1" {
		t.Errorf("ID = %q, want evt_1", ev.ID)
	}
	if ev.Type != V2EventTypeSessionStatus {
		t.Errorf("Type = %q, want %q", ev.Type, V2EventTypeSessionStatus)
	}

	payload, ok := ev.AsUnion().(V2EventSessionStatus)
	if !ok {
		t.Fatalf("AsUnion() = %T, want V2EventSessionStatus", ev.AsUnion())
	}
	if payload.Data.SessionID != "ses_e2e" {
		t.Errorf("Data.SessionID = %q, want ses_e2e", payload.Data.SessionID)
	}
	retry, ok := payload.Data.AsStatus().(SessionStatusRetry)
	if !ok {
		t.Fatalf("Data.AsStatus() = %T, want SessionStatusRetry", payload.Data.AsStatus())
	}
	if retry.Attempt != 2 || retry.Message != "m" || retry.Next != 9 {
		t.Errorf("retry = %+v", retry)
	}
	if payload.Durable.AggregateID != "ses_e2e" || payload.Durable.Seq != 7 || payload.Durable.Version != 1 {
		t.Errorf("Durable = %+v", payload.Durable)
	}
}

// TestEventListResponseEventSessionStatusPropertiesUnmarshal is the V1 `/event`
// counterpart (OpenAPI schema EventSessionStatus.properties). It guards the same
// Case-B contract: `sessionID` preserved and `status` resolved to the right variant.
func TestEventListResponseEventSessionStatusPropertiesUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		wantType string
		assert   func(t *testing.T, s SessionStatus)
	}{
		{
			name:     "idle",
			raw:      `{"sessionID":"ses_a","status":{"type":"idle"}}`,
			wantType: "idle",
			assert: func(t *testing.T, s SessionStatus) {
				if v, ok := s.(SessionStatusIdle); !ok || v.Type != "idle" {
					t.Errorf("got %#v, want SessionStatusIdle{Type:\"idle\"}", s)
				}
			},
		},
		{
			name:     "busy",
			raw:      `{"sessionID":"ses_b","status":{"type":"busy"}}`,
			wantType: "busy",
			assert: func(t *testing.T, s SessionStatus) {
				if v, ok := s.(SessionStatusBusy); !ok || v.Type != "busy" {
					t.Errorf("got %#v, want SessionStatusBusy{Type:\"busy\"}", s)
				}
			},
		},
		{
			name:     "retry",
			raw:      `{"sessionID":"ses_c","status":{"type":"retry","attempt":5,"message":"boom","next":11}}`,
			wantType: "retry",
			assert: func(t *testing.T, s SessionStatus) {
				v, ok := s.(SessionStatusRetry)
				if !ok {
					t.Fatalf("got %T, want SessionStatusRetry", s)
				}
				if v.Type != "retry" || v.Attempt != 5 || v.Message != "boom" || v.Next != 11 {
					t.Errorf("retry = %+v", v)
				}
			},
		},
	}

	wantSessionIDs := map[string]string{"idle": "ses_a", "busy": "ses_b", "retry": "ses_c"}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var p EventListResponseEventSessionStatusProperties
			if err := json.Unmarshal([]byte(tc.raw), &p); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if p.SessionID != wantSessionIDs[tc.name] {
				t.Errorf("SessionID = %q, want %q (sibling field must survive union decode)", p.SessionID, wantSessionIDs[tc.name])
			}
			if !reflect.DeepEqual(p.Status, p.AsStatus()) {
				t.Errorf("Status (%T) and AsStatus() (%T) diverge", p.Status, p.AsStatus())
			}
			if p.JSON.RawJSON() != tc.raw {
				t.Errorf("RawJSON not preserved:\n got %s\nwant %s", p.JSON.RawJSON(), tc.raw)
			}
			tc.assert(t, p.AsStatus())
		})
	}
}

// TestEventListResponseEventSessionStatusEndToEnd decodes a full V1 `session.status`
// event through the EventListResponse union.
func TestEventListResponseEventSessionStatusEndToEnd(t *testing.T) {
	raw := `{"id":"evt_2","type":"session.status","properties":{"sessionID":"ses_v1","status":{"type":"busy"}}}`

	var ev EventListResponse
	if err := json.Unmarshal([]byte(raw), &ev); err != nil {
		t.Fatalf("unmarshal EventListResponse: %v", err)
	}
	if ev.Type != EventListResponseTypeSessionStatus {
		t.Errorf("Type = %q, want %q", ev.Type, EventListResponseTypeSessionStatus)
	}
	payload, ok := ev.AsUnion().(EventListResponseEventSessionStatus)
	if !ok {
		t.Fatalf("AsUnion() = %T, want EventListResponseEventSessionStatus", ev.AsUnion())
	}
	if payload.Properties.SessionID != "ses_v1" {
		t.Errorf("Properties.SessionID = %q, want ses_v1", payload.Properties.SessionID)
	}
	if _, ok := payload.Properties.AsStatus().(SessionStatusBusy); !ok {
		t.Fatalf("AsStatus() = %T, want SessionStatusBusy", payload.Properties.AsStatus())
	}
}

// TestSessionStatusCarriersToleratePartialPayloads makes sure the carriers never
// panic and never lose `sessionID` when `status` is absent or null — the SDK must
// stay forward compatible with server payloads it does not fully model.
func TestSessionStatusCarriersToleratePartialPayloads(t *testing.T) {
	for _, raw := range []string{
		`{"sessionID":"ses_x"}`,
		`{"sessionID":"ses_x","status":null}`,
	} {
		var d V2EventSessionStatusData
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Fatalf("V2EventSessionStatusData unmarshal %s: %v", raw, err)
		}
		if d.SessionID != "ses_x" {
			t.Errorf("V2EventSessionStatusData.SessionID = %q for %s", d.SessionID, raw)
		}
		if d.AsStatus() != nil {
			t.Errorf("AsStatus() = %#v, want nil for %s", d.AsStatus(), raw)
		}

		var p EventListResponseEventSessionStatusProperties
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("EventListResponseEventSessionStatusProperties unmarshal %s: %v", raw, err)
		}
		if p.SessionID != "ses_x" {
			t.Errorf("EventListResponseEventSessionStatusProperties.SessionID = %q for %s", p.SessionID, raw)
		}
		if p.AsStatus() != nil {
			t.Errorf("AsStatus() = %#v, want nil for %s", p.AsStatus(), raw)
		}
	}
}

// TestQuestionToolNaming pins the Go type name to the OpenAPI/JS SDK v2 fact
// standard: OpenAPI declares `components/schemas/QuestionTool` and JS SDK v2
// exports `QuestionTool` (types.gen.ts:723). `QuestionRequestTool` never existed
// upstream.
func TestQuestionToolNaming(t *testing.T) {
	rt := reflect.TypeFor[QuestionTool]()
	if rt.Name() != "QuestionTool" {
		t.Fatalf("type name = %q, want QuestionTool", rt.Name())
	}
	if rt.NumField() != 3 {
		t.Errorf("QuestionTool has %d fields, want 3 (messageID, callID, JSON)", rt.NumField())
	}
	for name, wantTag := range map[string]string{
		"MessageID": "messageID,required",
		"CallID":    "callID,required",
	} {
		f, ok := rt.FieldByName(name)
		if !ok {
			t.Errorf("QuestionTool.%s missing", name)
			continue
		}
		if f.Type.Kind() != reflect.String {
			t.Errorf("QuestionTool.%s type = %s, want string", name, f.Type)
		}
		if got := f.Tag.Get("json"); got != wantTag {
			t.Errorf("QuestionTool.%s json tag = %q, want %q", name, got, wantTag)
		}
	}

	// QuestionRequest.tool -> $ref QuestionTool, optional (not in required list).
	tool, ok := reflect.TypeFor[QuestionRequest]().FieldByName("Tool")
	if !ok {
		t.Fatal("QuestionRequest.Tool missing")
	}
	if tool.Type != rt {
		t.Errorf("QuestionRequest.Tool type = %s, want QuestionTool", tool.Type)
	}
	if got := tool.Tag.Get("json"); got != "tool" {
		t.Errorf("QuestionRequest.Tool json tag = %q, want %q (optional)", got, "tool")
	}
}

// TestQuestionRequestWithQuestionToolUnmarshal exercises the renamed type through
// a full QuestionRequest payload (OpenAPI schema QuestionRequest).
func TestQuestionRequestWithQuestionToolUnmarshal(t *testing.T) {
	raw := `{"id":"que_1","sessionID":"ses_1","questions":[{"question":"Pick one","header":"h",` +
		`"options":[{"label":"A","description":"first"}],"multiple":false,"custom":false}],` +
		`"tool":{"messageID":"msg_1","callID":"call_1"}}`

	var req QuestionRequest
	if err := json.Unmarshal([]byte(raw), &req); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if req.ID != "que_1" || req.SessionID != "ses_1" {
		t.Errorf("ID/SessionID = %q/%q", req.ID, req.SessionID)
	}
	if len(req.Questions) != 1 || req.Questions[0].Question != "Pick one" {
		t.Fatalf("Questions = %+v", req.Questions)
	}
	if req.Tool.MessageID != "msg_1" || req.Tool.CallID != "call_1" {
		t.Errorf("Tool = %+v", req.Tool)
	}
	if req.Tool.JSON.RawJSON() != `{"messageID":"msg_1","callID":"call_1"}` {
		t.Errorf("Tool.RawJSON = %q", req.Tool.JSON.RawJSON())
	}

	// `tool` is optional — absence must not error.
	var noTool QuestionRequest
	if err := json.Unmarshal([]byte(`{"id":"que_2","sessionID":"ses_2","questions":[]}`), &noTool); err != nil {
		t.Fatalf("unmarshal without tool: %v", err)
	}
	if noTool.Tool.MessageID != "" || noTool.Tool.CallID != "" {
		t.Errorf("Tool should be zero when omitted, got %+v", noTool.Tool)
	}
}

// TestSessionStatusAsStatusUsesValueReceiver pins both status carriers to a
// *value* receiver for AsStatus().
//
// A pointer receiver keeps the method out of the value's method set, so the
// idiomatic SSE one-liner
//
//	ev.Properties.(EventListResponseEventSessionStatusProperties).AsStatus()
//
// fails to compile ("cannot call pointer method AsStatus on ..."), because a
// type-assertion result is not addressable. Every other union accessor in the SDK
// ([EventListResponse.AsUnion], [V2Event.AsUnion], [GlobalEvent.AsUnion],
// [V2EventSessionErrorData.AsError], ...) uses a value receiver; these two must
// not diverge.
func TestSessionStatusAsStatusUsesValueReceiver(t *testing.T) {
	for _, rt := range []reflect.Type{
		reflect.TypeFor[EventListResponseEventSessionStatusProperties](),
		reflect.TypeFor[V2EventSessionStatusData](),
	} {
		if _, ok := rt.MethodByName("AsStatus"); !ok {
			t.Errorf("%s: AsStatus must be in the value method set (value receiver)", rt.Name())
		}
	}

	// Exercise the one-liner that a pointer receiver would reject.
	var ev EventListResponse
	if err := json.Unmarshal([]byte(`{"id":"evt_1","type":"session.status","properties":{"sessionID":"ses_1","status":{"type":"busy"}}}`), &ev); err != nil {
		t.Fatalf("unmarshal EventListResponse: %v", err)
	}
	if _, ok := ev.Properties.(EventListResponseEventSessionStatusProperties).AsStatus().(SessionStatusBusy); !ok {
		t.Errorf("inline AsStatus() = %T, want SessionStatusBusy",
			ev.Properties.(EventListResponseEventSessionStatusProperties).AsStatus())
	}

	var v2 V2Event
	if err := json.Unmarshal([]byte(`{"id":"evt_2","type":"session.status","data":{"sessionID":"ses_1","status":{"type":"idle"}}}`), &v2); err != nil {
		t.Fatalf("unmarshal V2Event: %v", err)
	}
	if _, ok := v2.Data.(V2EventSessionStatusData).AsStatus().(SessionStatusIdle); !ok {
		t.Errorf("inline AsStatus() = %T, want SessionStatusIdle",
			v2.Data.(V2EventSessionStatusData).AsStatus())
	}
}
