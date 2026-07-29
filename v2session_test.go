package opencode

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestV2SessionService(t *testing.T) {
	s := NewV2SessionService()
	_ = s
}

func TestV2SessionListParams(t *testing.T) {
	params := V2SessionListParams{
		Limit: Int(10),
		Order: F(V2SessionOrderDesc),
	}
	v := params.URLQuery()
	if v.Get("limit") != "10" {
		t.Errorf("expected limit=10, got %q", v.Get("limit"))
	}
	if v.Get("order") != "desc" {
		t.Errorf("expected order=desc, got %q", v.Get("order"))
	}
}

func TestV2SessionPromptParams(t *testing.T) {
	params := V2SessionPromptParams{
		Prompt: F(V2PromptInputParam{
			Text: String("hello"),
		}),
		Delivery: F(SessionDeliverySteer),
	}
	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty JSON body")
	}
}

func TestV2SessionsResponseUnmarshal(t *testing.T) {
	jsonStr := `{
		"data": [
			{
				"id": "ses_123",
				"projectID": "proj_1",
				"time": {"created": 1, "updated": 2},
				"title": "test"
			}
		],
		"cursor": {
			"previous": "abc",
			"next": "def"
		}
	}`
	var resp V2SessionsResponse
	if err := resp.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("expected 1 item, got %d", len(resp.Data))
	}
	if resp.Data[0].ID != "ses_123" {
		t.Errorf("expected ses_123, got %s", resp.Data[0].ID)
	}
	if resp.Cursor.Previous != "abc" {
		t.Errorf("expected abc, got %s", resp.Cursor.Previous)
	}
}

func TestV2SessionMessagesResponseUnmarshal(t *testing.T) {
	jsonStr := `{
		"data": [],
		"cursor": {
			"previous": "",
			"next": ""
		}
	}`
	var resp V2SessionMessagesResponse
	if err := resp.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 0 {
		t.Errorf("expected 0 items, got %d", len(resp.Data))
	}
}

func TestV2SessionServiceRequiresSessionID(t *testing.T) {
	s := NewV2SessionService()
	ctx := context.Background()

	_, err := s.Prompt(ctx, "", V2SessionPromptParams{})
	if err == nil {
		t.Error("expected error for empty sessionID")
	}

	err = s.Compact(ctx, "")
	if err == nil {
		t.Error("expected error for empty sessionID")
	}

	err = s.Wait(ctx, "")
	if err == nil {
		t.Error("expected error for empty sessionID")
	}

	_, err = s.Context(ctx, "")
	if err == nil {
		t.Error("expected error for empty sessionID")
	}

	_, err = s.Messages(ctx, "", V2SessionMessagesParams{})
	if err == nil {
		t.Error("expected error for empty sessionID")
	}
}

func TestV2SessionMessageResponseUnmarshal(t *testing.T) {
	jsonStr := `{
		"data": {
			"id": "msg_123",
			"type": "user",
			"time": {"created": 123},
			"text": "hello"
		}
	}`
	var resp V2SessionMessageResponse
	if err := resp.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	msg := resp.Data.AsUnion()
	if msg == nil {
		t.Fatal("expected union to be non-nil")
	}
	userMsg, ok := msg.(V2SessionMessageUser)
	if !ok {
		t.Fatalf("expected V2SessionMessageUser, got %T", msg)
	}
	if userMsg.ID != "msg_123" {
		t.Errorf("expected msg_123, got %s", userMsg.ID)
	}
	if userMsg.Text != "hello" {
		t.Errorf("expected hello, got %s", userMsg.Text)
	}
}

func TestV2SessionAssistantContentUnionText(t *testing.T) {
	jsonStr := `{
		"type": "text",
		"id": "txt_1",
		"text": "hello world"
	}`
	var content V2SessionMessageAssistantContent
	if err := content.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	text, ok := content.AsUnion().(V2SessionMessageAssistantTextContent)
	if !ok {
		t.Fatalf("expected V2SessionMessageAssistantTextContent, got %T", content.AsUnion())
	}
	if text.ID != "txt_1" {
		t.Errorf("expected txt_1, got %s", text.ID)
	}
	if text.Text != "hello world" {
		t.Errorf("expected hello world, got %s", text.Text)
	}
}

func TestV2SessionAssistantContentUnionTool(t *testing.T) {
	jsonStr := `{
		"type": "tool",
		"id": "call_1",
		"name": "bash",
		"state": {
			"status": "completed",
			"input": {"cmd": "ls"},
			"structured": {},
			"content": [],
			"outputPaths": ["/tmp/foo"],
			"result": {"ok": true}
		},
		"time": {"created": 1}
	}`
	var content V2SessionMessageAssistantContent
	if err := content.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	tool, ok := content.AsUnion().(V2SessionMessageAssistantToolContent)
	if !ok {
		t.Fatalf("expected V2SessionMessageAssistantToolContent, got %T", content.AsUnion())
	}
	if tool.ID != "call_1" {
		t.Errorf("expected call_1, got %s", tool.ID)
	}
	if tool.Name != "bash" {
		t.Errorf("expected bash, got %s", tool.Name)
	}
}

func TestV2SessionToolStateCompletedExtraFields(t *testing.T) {
	jsonStr := `{
		"status": "completed",
		"input": {},
		"structured": {},
		"content": [],
		"outputPaths": ["/a", "/b"],
		"result": {"key": "value"}
	}`
	var state V2SessionMessageToolStateCompleted
	if err := state.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if state.Status != V2SessionMessageToolStateCompletedStatusCompleted {
		t.Errorf("expected completed, got %s", state.Status)
	}
}

func TestV2SessionToolStateErrorExtraFields(t *testing.T) {
	jsonStr := `{
		"status": "error",
		"input": {},
		"structured": {},
		"content": [],
		"error": {"type": "unknown", "message": "boom"},
		"result": {"err": "x"}
	}`
	var state V2SessionMessageToolStateError
	if err := state.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if state.Status != V2SessionMessageToolStateErrorStatusError {
		t.Errorf("expected error, got %s", state.Status)
	}
}

func TestV2SessionHistoryResponseUnmarshal(t *testing.T) {
	jsonStr := `{
		"data": [
			{
				"id": "evt_1",
				"type": "session.next.agent.switched",
				"data": {
					"timestamp": 1,
					"sessionID": "ses_1",
					"messageID": "msg_1",
					"agent": "build"
				}
			}
		],
		"hasMore": false
	}`
	var resp V2SessionHistoryResponse
	if err := resp.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Data))
	}
	if resp.HasMore {
		t.Errorf("expected hasMore=false")
	}
}

func TestV2SessionDurableEventAgentSwitchedUnmarshal(t *testing.T) {
	jsonStr := `{
		"id": "evt_1",
		"type": "session.next.agent.switched",
		"durable": {
			"aggregateID": "agg_1",
			"seq": 1,
			"version": 1
		},
		"location": {
			"directory": "/tmp",
			"workspaceID": "ws_1"
		},
		"data": {
			"timestamp": 1,
			"sessionID": "ses_1",
			"messageID": "msg_1",
			"agent": "build"
		}
	}`
	var evt V2SessionDurableEvent
	if err := evt.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	switched, ok := evt.AsUnion().(V2SessionDurableEventAgentSwitched)
	if !ok {
		t.Fatalf("expected V2SessionDurableEventAgentSwitched, got %T", evt.AsUnion())
	}
	if switched.ID != "evt_1" {
		t.Errorf("expected evt_1, got %s", switched.ID)
	}
	if switched.Durable.AggregateID != "agg_1" {
		t.Errorf("expected agg_1, got %s", switched.Durable.AggregateID)
	}
	if switched.Data.Agent != "build" {
		t.Errorf("expected build, got %s", switched.Data.Agent)
	}
}

func TestV2SessionAssistantTextContentHasID(t *testing.T) {
	jsonStr := `{
		"type": "text",
		"id": "txt_42",
		"text": "content"
	}`
	var content V2SessionMessageAssistantTextContent
	if err := content.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if content.ID != "txt_42" {
		t.Errorf("expected txt_42, got %s", content.ID)
	}
	if content.Text != "content" {
		t.Errorf("expected content, got %s", content.Text)
	}
}

func TestRevertStateFilesType(t *testing.T) {
	jsonStr := `{
		"messageID": "msg_1",
		"files": [
			{"path": "/a", "status": "added", "additions": 1, "deletions": 0, "patch": "p"}
		]
	}`
	var rs RevertState
	if err := rs.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if rs.MessageID != "msg_1" {
		t.Errorf("expected msg_1, got %s", rs.MessageID)
	}
	if len(rs.Files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(rs.Files))
	}
	if rs.Files[0].Path != "/a" {
		t.Errorf("expected /a, got %s", rs.Files[0].Path)
	}
}

// TestV2SessionUnionRoutingMatrix asserts that every member of the four
// discriminated unions declared in v2session.go routes to its own Go type.
//
// The variant lists come straight from OpenAPI: `SessionMessage` (anyOf, 8
// members), `SessionMessageAssistant.content` items (anyOf, 3), the
// `SessionMessageAssistantTool.state` anyOf (4) and `SessionDurableEvent` (oneOf,
// 28). Every member pins its tag property (`type`, or `status` for the tool
// states) to a single-value enum, which is why all four are registered with that
// property as apijson's discriminator key.
//
// This is a regression net for a real defect: while these unions were registered
// without a discriminator key, apijson's exactness heuristic decided the variant.
// Because it never penalises a missing `required` field and the v2 variants type
// their tag as a plain `string` (so guardUnknown cannot help either), all 28
// `SessionDurableEvent` members tied and the left-most one -- agent.switched --
// won for every payload; `SessionMessage` reached only 3 of 8 and the assistant
// content union only 1 of 3. Each sub-test therefore also asserts the *count* of
// distinct types reached, so re-introducing the collapse cannot pass unnoticed.
//
// Each payload is additionally re-run with an unknown extra property, because a
// single unknown field drops every variant to the same "extras" exactness level
// and used to collapse the routing even for the members that did work.
func TestV2SessionUnionRoutingMatrix(t *testing.T) {
	type routeCase struct {
		raw  string
		want any
	}
	run := func(t *testing.T, name string, decode func(raw string) (any, error), cases []routeCase) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			for _, mode := range []string{"exact", "withUnknownField"} {
				seen := map[string]bool{}
				for _, tc := range cases {
					raw := tc.raw
					if mode == "withUnknownField" {
						raw = raw[:len(raw)-1] + `,"unknownFutureField":"x"}`
					}
					got, err := decode(raw)
					if err != nil {
						t.Fatalf("%s: unmarshal %s: %v", mode, raw, err)
					}
					if reflect.TypeOf(got) != reflect.TypeOf(tc.want) {
						t.Errorf("%s: %s\n  routed to %v, want %v", mode, raw, reflect.TypeOf(got), reflect.TypeOf(tc.want))
					}
					seen[fmt.Sprintf("%T", got)] = true
				}
				if len(seen) != len(cases) {
					t.Errorf("%s: reached %d distinct variants, want %d (%v)", mode, len(seen), len(cases), seen)
				}
			}
		})
	}

	// OpenAPI SessionMessage anyOf — 8 members.
	run(t, "V2SessionMessageUnion", func(raw string) (any, error) {
		var m V2SessionMessage
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			return nil, err
		}
		return m.AsUnion(), nil
	}, []routeCase{
		{`{"id":"msg_1","type":"agent-switched","time":{"created":1},"agent":"build"}`, V2SessionMessageAgentSwitched{}},
		{`{"id":"msg_2","type":"model-switched","time":{"created":1},"model":{"providerID":"p","id":"m"}}`, V2SessionMessageModelSwitched{}},
		{`{"id":"msg_3","type":"user","time":{"created":1},"text":"hi"}`, V2SessionMessageUser{}},
		{`{"id":"msg_4","type":"synthetic","time":{"created":1},"sessionID":"ses_1","text":"hi"}`, V2SessionMessageSynthetic{}},
		{`{"id":"msg_5","type":"system","time":{"created":1},"text":"hi"}`, V2SessionMessageSystem{}},
		{`{"id":"msg_6","type":"shell","time":{"created":1},"callID":"c","command":"ls","output":"a"}`, V2SessionMessageShell{}},
		{`{"id":"msg_7","type":"assistant","time":{"created":1},"agent":"build","model":{"providerID":"p","id":"m"},"content":[]}`, V2SessionMessageAssistant{}},
		{`{"id":"msg_8","type":"compaction","time":{"created":1},"reason":"overflow","summary":"s","recent":[]}`, V2SessionMessageCompaction{}},
	})

	// OpenAPI SessionMessageAssistant.content items anyOf — 3 members.
	run(t, "V2SessionMessageAssistantContentUnion", func(raw string) (any, error) {
		var c V2SessionMessageAssistantContent
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			return nil, err
		}
		return c.AsUnion(), nil
	}, []routeCase{
		{`{"type":"text","id":"c1","text":"hi"}`, V2SessionMessageAssistantTextContent{}},
		{`{"type":"reasoning","id":"c2","text":"why"}`, V2SessionMessageAssistantReasoningContent{}},
		{`{"type":"tool","id":"c3","name":"bash","state":{"status":"pending","input":"{}"},"time":{"created":1}}`, V2SessionMessageAssistantToolContent{}},
	})

	// OpenAPI SessionMessageAssistantTool.state anyOf — 4 members, keyed on `status`.
	run(t, "V2SessionMessageToolStateUnion", func(raw string) (any, error) {
		var s V2SessionMessageToolState
		if err := json.Unmarshal([]byte(raw), &s); err != nil {
			return nil, err
		}
		return s.AsUnion(), nil
	}, []routeCase{
		{`{"status":"pending","input":"{}"}`, V2SessionMessageToolStatePending{}},
		{`{"status":"running","input":{},"structured":{},"content":[]}`, V2SessionMessageToolStateRunning{}},
		{`{"status":"completed","input":{},"structured":{},"content":[]}`, V2SessionMessageToolStateCompleted{}},
		{`{"status":"error","input":{},"structured":{},"content":[],"error":{"message":"boom"}}`, V2SessionMessageToolStateError{}},
	})

	// OpenAPI SessionDurableEvent oneOf — 28 members. `id`, `type` and `data` are
	// the only required properties, and every member declares the identical
	// property set, so `type` is the only thing that can tell them apart.
	run(t, "V2SessionDurableEventUnion", func(raw string) (any, error) {
		var e V2SessionDurableEvent
		if err := json.Unmarshal([]byte(raw), &e); err != nil {
			return nil, err
		}
		return e.AsUnion(), nil
	}, []routeCase{
		{`{"id":"evt_1","type":"session.next.agent.switched","data":{}}`, V2SessionDurableEventAgentSwitched{}},
		{`{"id":"evt_1","type":"session.next.model.switched","data":{}}`, V2SessionDurableEventModelSwitched{}},
		{`{"id":"evt_1","type":"session.next.moved","data":{}}`, V2SessionDurableEventMoved{}},
		{`{"id":"evt_1","type":"session.next.prompted","data":{}}`, V2SessionDurableEventPrompted{}},
		{`{"id":"evt_1","type":"session.next.prompt.admitted","data":{}}`, V2SessionDurableEventPromptAdmitted{}},
		{`{"id":"evt_1","type":"session.next.context.updated","data":{}}`, V2SessionDurableEventContextUpdated{}},
		{`{"id":"evt_1","type":"session.next.synthetic","data":{}}`, V2SessionDurableEventSynthetic{}},
		{`{"id":"evt_1","type":"session.next.shell.started","data":{}}`, V2SessionDurableEventShellStarted{}},
		{`{"id":"evt_1","type":"session.next.shell.ended","data":{}}`, V2SessionDurableEventShellEnded{}},
		{`{"id":"evt_1","type":"session.next.step.started","data":{}}`, V2SessionDurableEventStepStarted{}},
		{`{"id":"evt_1","type":"session.next.step.ended","data":{}}`, V2SessionDurableEventStepEnded{}},
		{`{"id":"evt_1","type":"session.next.step.failed","data":{}}`, V2SessionDurableEventStepFailed{}},
		{`{"id":"evt_1","type":"session.next.text.started","data":{}}`, V2SessionDurableEventTextStarted{}},
		{`{"id":"evt_1","type":"session.next.text.ended","data":{}}`, V2SessionDurableEventTextEnded{}},
		{`{"id":"evt_1","type":"session.next.tool.input.started","data":{}}`, V2SessionDurableEventToolInputStarted{}},
		{`{"id":"evt_1","type":"session.next.tool.input.ended","data":{}}`, V2SessionDurableEventToolInputEnded{}},
		{`{"id":"evt_1","type":"session.next.tool.called","data":{}}`, V2SessionDurableEventToolCalled{}},
		{`{"id":"evt_1","type":"session.next.tool.progress","data":{}}`, V2SessionDurableEventToolProgress{}},
		{`{"id":"evt_1","type":"session.next.tool.success","data":{}}`, V2SessionDurableEventToolSuccess{}},
		{`{"id":"evt_1","type":"session.next.tool.failed","data":{}}`, V2SessionDurableEventToolFailed{}},
		{`{"id":"evt_1","type":"session.next.reasoning.started","data":{}}`, V2SessionDurableEventReasoningStarted{}},
		{`{"id":"evt_1","type":"session.next.reasoning.ended","data":{}}`, V2SessionDurableEventReasoningEnded{}},
		{`{"id":"evt_1","type":"session.next.retried","data":{}}`, V2SessionDurableEventRetried{}},
		{`{"id":"evt_1","type":"session.next.compaction.started","data":{}}`, V2SessionDurableEventCompactionStarted{}},
		{`{"id":"evt_1","type":"session.next.compaction.ended","data":{}}`, V2SessionDurableEventCompactionEnded{}},
		{`{"id":"evt_1","type":"session.next.revert.staged","data":{}}`, V2SessionDurableEventRevertStaged{}},
		{`{"id":"evt_1","type":"session.next.revert.cleared","data":{}}`, V2SessionDurableEventRevertCleared{}},
		{`{"id":"evt_1","type":"session.next.revert.committed","data":{}}`, V2SessionDurableEventRevertCommitted{}}})
}
