package opencode

import (
	"context"
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

func TestV2SessionMessageToolProviderResultMetadata(t *testing.T) {
	t.Parallel()
	jsonStr := `{
		"executed": true,
		"metadata": {"key": "value"},
		"resultMetadata": {"result_key": "result_value"}
	}`
	var provider V2SessionMessageToolProvider
	if err := provider.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if !provider.Executed {
		t.Error("expected Executed to be true")
	}
	metadata, ok := provider.Metadata.(map[string]any)
	if !ok {
		t.Fatalf("expected Metadata to be map[string]interface{}, got %T", provider.Metadata)
	}
	if metadata["key"] != "value" {
		t.Errorf("expected metadata key=value, got %v", metadata["key"])
	}
	resultMetadata, ok := provider.ResultMetadata.(map[string]any)
	if !ok {
		t.Fatalf("expected ResultMetadata to be map[string]interface{}, got %T", provider.ResultMetadata)
	}
	if resultMetadata["result_key"] != "result_value" {
		t.Errorf("expected resultMetadata result_key=result_value, got %v", resultMetadata["result_key"])
	}
}

func TestV2SessionMessageToolProviderWithoutResultMetadata(t *testing.T) {
	t.Parallel()
	jsonStr := `{"executed": false}`
	var provider V2SessionMessageToolProvider
	if err := provider.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if provider.Executed {
		t.Error("expected Executed to be false")
	}
	if provider.Metadata != nil {
		t.Errorf("expected Metadata to be nil, got %v", provider.Metadata)
	}
	if provider.ResultMetadata != nil {
		t.Errorf("expected ResultMetadata to be nil, got %v", provider.ResultMetadata)
	}
}

// TestToolStateContentUnionDecoding verifies that the Content field on
// V2SessionMessageToolStateRunning / Completed / Error decodes as
// []LLMToolContent with concrete element types, not []interface{}.
func TestToolStateContentUnionDecoding(t *testing.T) {
	t.Parallel()

	mixedContent := `[{"type":"text","text":"hello"},{"type":"file","uri":"file:///a.txt","mime":"text/plain","name":"a.txt"}]`

	t.Run("Running", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{"status":"running","input":{},"structured":{},"content":` + mixedContent + `}`
		var state V2SessionMessageToolStateRunning
		if err := state.UnmarshalJSON([]byte(jsonStr)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		if len(state.Content) != 2 {
			t.Fatalf("expected 2 content items, got %d", len(state.Content))
		}
		textItem, ok := state.Content[0].AsUnion().(ToolTextContent)
		if !ok {
			t.Errorf("element[0]: expected ToolTextContent, got %T", state.Content[0])
		} else if textItem.Text != "hello" {
			t.Errorf("element[0].Text: expected hello, got %q", textItem.Text)
		}
		fileItem, ok := state.Content[1].AsUnion().(ToolFileContent)
		if !ok {
			t.Errorf("element[1]: expected ToolFileContent, got %T", state.Content[1])
		} else if fileItem.URI != "file:///a.txt" {
			t.Errorf("element[1].URI: expected file:///a.txt, got %q", fileItem.URI)
		}
	})

	t.Run("Completed", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{"status":"completed","input":{},"structured":{},"content":` + mixedContent + `}`
		var state V2SessionMessageToolStateCompleted
		if err := state.UnmarshalJSON([]byte(jsonStr)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		if len(state.Content) != 2 {
			t.Fatalf("expected 2 content items, got %d", len(state.Content))
		}
		if _, ok := state.Content[0].AsUnion().(ToolTextContent); !ok {
			t.Errorf("element[0]: expected ToolTextContent, got %T", state.Content[0])
		}
		if _, ok := state.Content[1].AsUnion().(ToolFileContent); !ok {
			t.Errorf("element[1]: expected ToolFileContent, got %T", state.Content[1])
		}
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{"status":"error","input":{},"structured":{},"content":` + mixedContent + `,"error":{"type":"unknown","message":"boom"}}`
		var state V2SessionMessageToolStateError
		if err := state.UnmarshalJSON([]byte(jsonStr)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		if len(state.Content) != 2 {
			t.Fatalf("expected 2 content items, got %d", len(state.Content))
		}
		if _, ok := state.Content[0].AsUnion().(ToolTextContent); !ok {
			t.Errorf("element[0]: expected ToolTextContent, got %T", state.Content[0])
		}
		if _, ok := state.Content[1].AsUnion().(ToolFileContent); !ok {
			t.Errorf("element[1]: expected ToolFileContent, got %T", state.Content[1])
		}
	})

	t.Run("EmptyContent", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{"status":"completed","input":{},"structured":{},"content":[]}`
		var state V2SessionMessageToolStateCompleted
		if err := state.UnmarshalJSON([]byte(jsonStr)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		if len(state.Content) != 0 {
			t.Errorf("expected empty content, got %d items", len(state.Content))
		}
	})

	t.Run("ViaToolState", func(t *testing.T) {
		// Verify decoding through the parent V2SessionMessageToolState union.
		t.Parallel()
		jsonStr := `{"status":"completed","input":{},"structured":{},"content":[{"type":"text","text":"world"}]}`
		var state V2SessionMessageToolState
		if err := state.UnmarshalJSON([]byte(jsonStr)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		completed, ok := state.AsUnion().(V2SessionMessageToolStateCompleted)
		if !ok {
			t.Fatalf("expected V2SessionMessageToolStateCompleted, got %T", state.AsUnion())
		}
		if len(completed.Content) != 1 {
			t.Fatalf("expected 1 content item, got %d", len(completed.Content))
		}
		textItem, ok := completed.Content[0].AsUnion().(ToolTextContent)
		if !ok {
			t.Errorf("element[0]: expected ToolTextContent, got %T", completed.Content[0])
		} else if textItem.Text != "world" {
			t.Errorf("element[0].Text: expected world, got %q", textItem.Text)
		}
	})
}

// TestLLMToolContentCarrierTypedFields verifies that the LLMToolContent union
// carrier exposes typed string fields (Text/URI/Mime/Name) and an enum-typed
// Type discriminator instead of any.
func TestLLMToolContentCarrierTypedFields(t *testing.T) {
	t.Parallel()

	t.Run("TextVariant", func(t *testing.T) {
		t.Parallel()
		var c LLMToolContent
		if err := c.UnmarshalJSON([]byte(`{"type":"text","text":"hello"}`)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		if c.Type != LLMToolContentTypeText {
			t.Errorf("Type: got %q, want %q", c.Type, LLMToolContentTypeText)
		}
		if c.Text != "hello" {
			t.Errorf("Text: got %q, want hello", c.Text)
		}
		if _, ok := c.AsUnion().(ToolTextContent); !ok {
			t.Errorf("AsUnion: expected ToolTextContent, got %T", c.AsUnion())
		}
	})

	t.Run("FileVariant", func(t *testing.T) {
		t.Parallel()
		var c LLMToolContent
		if err := c.UnmarshalJSON([]byte(`{"type":"file","uri":"file:///a.txt","mime":"text/plain","name":"a.txt"}`)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		if c.Type != LLMToolContentTypeFile {
			t.Errorf("Type: got %q, want %q", c.Type, LLMToolContentTypeFile)
		}
		if c.URI != "file:///a.txt" || c.Mime != "text/plain" || c.Name != "a.txt" {
			t.Errorf("URI/Mime/Name: %q / %q / %q", c.URI, c.Mime, c.Name)
		}
		if _, ok := c.AsUnion().(ToolFileContent); !ok {
			t.Errorf("AsUnion: expected ToolFileContent, got %T", c.AsUnion())
		}
	})
}

// ===== F1/F2 regression tests (v1.18.18 阶段二 · D 任务) =====
//
// These tests guard the F1 fix (38/39 discriminated-union `type` fields that
// were previously a bare `string` are now typed enums with `IsKnown()`) and
// the F2 fix (`SessionErrorUnknown.Type` is now `SessionErrorUnknownType`).
// Discriminator string literals below are copied verbatim from
// opencode/packages/sdk/openapi.json (SessionDurableEvent.oneOf /
// SessionMessage.anyOf / SessionMessageAssistant.content.items.anyOf /
// SessionErrorUnknown.properties.type.enum) — see
// .tmp/sync-v1.18.18/D-phase2/enum_diff.py for the extraction script and its
// bidirectional-diff output (empty in both directions).

// typeFieldString extracts the string value of a struct's exported `Type`
// field via reflection. All discriminator enum types in this file are named
// string types, so Value.String() returns the underlying value directly.
func typeFieldString(v any) string {
	return reflect.ValueOf(v).FieldByName("Type").String()
}

// TestV2SessionMessageAllDiscriminators covers all 8 V2SessionMessage union
// variants (F1: 8 of the 39 bare-string discriminator fields). For each
// OpenAPI SessionMessage.anyOf member, decoding must route AsUnion() to the
// correct concrete Go type and populate the newly-typed Type field with the
// exact OpenAPI enum literal.
func TestV2SessionMessageAllDiscriminators(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		json        string
		wantType    reflect.Type
		wantTypeStr string
	}{
		{
			name:        "agent-switched",
			json:        `{"id":"m1","time":{"created":1},"type":"agent-switched","agent":"build"}`,
			wantType:    reflect.TypeOf(V2SessionMessageAgentSwitched{}),
			wantTypeStr: "agent-switched",
		},
		{
			name:        "model-switched",
			json:        `{"id":"m1","time":{"created":1},"type":"model-switched","model":{"id":"gpt","providerID":"openai"}}`,
			wantType:    reflect.TypeOf(V2SessionMessageModelSwitched{}),
			wantTypeStr: "model-switched",
		},
		{
			name:        "user",
			json:        `{"id":"m1","time":{"created":1},"type":"user","text":"hi"}`,
			wantType:    reflect.TypeOf(V2SessionMessageUser{}),
			wantTypeStr: "user",
		},
		{
			name:        "synthetic",
			json:        `{"id":"m1","time":{"created":1},"type":"synthetic","sessionID":"s1","text":"hi"}`,
			wantType:    reflect.TypeOf(V2SessionMessageSynthetic{}),
			wantTypeStr: "synthetic",
		},
		{
			name:        "system",
			json:        `{"id":"m1","time":{"created":1},"type":"system","text":"hi"}`,
			wantType:    reflect.TypeOf(V2SessionMessageSystem{}),
			wantTypeStr: "system",
		},
		{
			name:        "shell",
			json:        `{"id":"m1","time":{"created":1},"type":"shell","callID":"c1","command":"ls","output":"out"}`,
			wantType:    reflect.TypeOf(V2SessionMessageShell{}),
			wantTypeStr: "shell",
		},
		{
			name:        "assistant",
			json:        `{"id":"m1","time":{"created":1},"type":"assistant","agent":"build","model":{"id":"gpt","providerID":"openai"},"content":[]}`,
			wantType:    reflect.TypeOf(V2SessionMessageAssistant{}),
			wantTypeStr: "assistant",
		},
		{
			name:        "compaction",
			json:        `{"id":"m1","time":{"created":1},"type":"compaction","reason":"auto","summary":"s","recent":"r"}`,
			wantType:    reflect.TypeOf(V2SessionMessageCompaction{}),
			wantTypeStr: "compaction",
		},
	}

	if len(tests) != 8 {
		t.Fatalf("expected 8 V2SessionMessage discriminator cases (OpenAPI SessionMessage.anyOf has 8 members), got %d", len(tests))
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var msg V2SessionMessage
			if err := msg.UnmarshalJSON([]byte(tc.json)); err != nil {
				t.Fatalf("UnmarshalJSON: %v", err)
			}
			got := msg.AsUnion()
			if reflect.TypeOf(got) != tc.wantType {
				t.Fatalf("AsUnion: got %T, want %s", got, tc.wantType)
			}
			if gotStr := typeFieldString(got); gotStr != tc.wantTypeStr {
				t.Errorf("Type field: got %q, want %q", gotStr, tc.wantTypeStr)
			}
		})
	}
}

// TestV2SessionAssistantContentAllDiscriminators covers all 3
// V2SessionMessageAssistantContent union variants (F1: 3 of the 39 bare-string
// discriminator fields), sourced from OpenAPI
// SessionMessageAssistant.properties.content.items.anyOf.
func TestV2SessionAssistantContentAllDiscriminators(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		json        string
		wantType    reflect.Type
		wantTypeStr string
	}{
		{
			name:        "text",
			json:        `{"type":"text","id":"c1","text":"hello"}`,
			wantType:    reflect.TypeOf(V2SessionMessageAssistantTextContent{}),
			wantTypeStr: "text",
		},
		{
			name:        "reasoning",
			json:        `{"type":"reasoning","id":"c1","text":"thinking"}`,
			wantType:    reflect.TypeOf(V2SessionMessageAssistantReasoningContent{}),
			wantTypeStr: "reasoning",
		},
		{
			name:        "tool",
			json:        `{"type":"tool","id":"c1","name":"bash","state":{"status":"pending","input":""},"time":{"created":1}}`,
			wantType:    reflect.TypeOf(V2SessionMessageAssistantToolContent{}),
			wantTypeStr: "tool",
		},
	}

	if len(tests) != 3 {
		t.Fatalf("expected 3 AssistantContent discriminator cases, got %d", len(tests))
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var content V2SessionMessageAssistantContent
			if err := content.UnmarshalJSON([]byte(tc.json)); err != nil {
				t.Fatalf("UnmarshalJSON: %v", err)
			}
			got := content.AsUnion()
			if reflect.TypeOf(got) != tc.wantType {
				t.Fatalf("AsUnion: got %T, want %s", got, tc.wantType)
			}
			if gotStr := typeFieldString(got); gotStr != tc.wantTypeStr {
				t.Errorf("Type field: got %q, want %q", gotStr, tc.wantTypeStr)
			}
		})
	}
}

// TestV2SessionToolStateAllDiscriminators covers all 4
// V2SessionMessageToolState union variants (status discriminator; these were
// already typed pre-F1, kept here for completeness of the ToolState routing
// regression guard requested for the D task).
func TestV2SessionToolStateAllDiscriminators(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		json       string
		wantType   reflect.Type
		wantStatus V2SessionMessageToolStateStatus
	}{
		{
			name:       "pending",
			json:       `{"status":"pending","input":""}`,
			wantType:   reflect.TypeOf(V2SessionMessageToolStatePending{}),
			wantStatus: V2SessionMessageToolStateStatusPending,
		},
		{
			name:       "running",
			json:       `{"status":"running","input":{},"structured":{},"content":[]}`,
			wantType:   reflect.TypeOf(V2SessionMessageToolStateRunning{}),
			wantStatus: V2SessionMessageToolStateStatusRunning,
		},
		{
			name:       "completed",
			json:       `{"status":"completed","input":{},"structured":{},"content":[]}`,
			wantType:   reflect.TypeOf(V2SessionMessageToolStateCompleted{}),
			wantStatus: V2SessionMessageToolStateStatusCompleted,
		},
		{
			name:       "error",
			json:       `{"status":"error","input":{},"structured":{},"content":[],"error":{"type":"unknown","message":"boom"}}`,
			wantType:   reflect.TypeOf(V2SessionMessageToolStateError{}),
			wantStatus: V2SessionMessageToolStateStatusError,
		},
	}

	if len(tests) != 4 {
		t.Fatalf("expected 4 ToolState discriminator cases, got %d", len(tests))
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var state V2SessionMessageToolState
			if err := state.UnmarshalJSON([]byte(tc.json)); err != nil {
				t.Fatalf("UnmarshalJSON: %v", err)
			}
			got := state.AsUnion()
			if reflect.TypeOf(got) != tc.wantType {
				t.Fatalf("AsUnion: got %T, want %s", got, tc.wantType)
			}
			if state.Status != tc.wantStatus {
				t.Errorf("Status: got %q, want %q", state.Status, tc.wantStatus)
			}
		})
	}
}

// TestV2SessionDurableEventAllDiscriminators covers all 28
// V2SessionDurableEvent union variants (F1: 28 of the 39 bare-string
// discriminator fields; these now reuse the pre-existing
// V2EventSessionNext*Type enums from v2event.go per the SKILL's "reuse before
// creating" rule, since the OpenAPI event-type literal, JS SDK(v2) literal,
// and v2event.go const value are all identical for every one of the 28
// members). This is the core regression guard for the F1 fix.
func TestV2SessionDurableEventAllDiscriminators(t *testing.T) {
	t.Parallel()

	tests := []struct {
		discriminator string
		wantType      reflect.Type
	}{
		{"session.next.agent.switched", reflect.TypeOf(V2SessionDurableEventAgentSwitched{})},
		{"session.next.model.switched", reflect.TypeOf(V2SessionDurableEventModelSwitched{})},
		{"session.next.moved", reflect.TypeOf(V2SessionDurableEventMoved{})},
		{"session.next.prompted", reflect.TypeOf(V2SessionDurableEventPrompted{})},
		{"session.next.prompt.admitted", reflect.TypeOf(V2SessionDurableEventPromptAdmitted{})},
		{"session.next.context.updated", reflect.TypeOf(V2SessionDurableEventContextUpdated{})},
		{"session.next.synthetic", reflect.TypeOf(V2SessionDurableEventSynthetic{})},
		{"session.next.shell.started", reflect.TypeOf(V2SessionDurableEventShellStarted{})},
		{"session.next.shell.ended", reflect.TypeOf(V2SessionDurableEventShellEnded{})},
		{"session.next.step.started", reflect.TypeOf(V2SessionDurableEventStepStarted{})},
		{"session.next.step.ended", reflect.TypeOf(V2SessionDurableEventStepEnded{})},
		{"session.next.step.failed", reflect.TypeOf(V2SessionDurableEventStepFailed{})},
		{"session.next.text.started", reflect.TypeOf(V2SessionDurableEventTextStarted{})},
		{"session.next.text.ended", reflect.TypeOf(V2SessionDurableEventTextEnded{})},
		{"session.next.tool.input.started", reflect.TypeOf(V2SessionDurableEventToolInputStarted{})},
		{"session.next.tool.input.ended", reflect.TypeOf(V2SessionDurableEventToolInputEnded{})},
		{"session.next.tool.called", reflect.TypeOf(V2SessionDurableEventToolCalled{})},
		{"session.next.tool.progress", reflect.TypeOf(V2SessionDurableEventToolProgress{})},
		{"session.next.tool.success", reflect.TypeOf(V2SessionDurableEventToolSuccess{})},
		{"session.next.tool.failed", reflect.TypeOf(V2SessionDurableEventToolFailed{})},
		{"session.next.reasoning.started", reflect.TypeOf(V2SessionDurableEventReasoningStarted{})},
		{"session.next.reasoning.ended", reflect.TypeOf(V2SessionDurableEventReasoningEnded{})},
		{"session.next.retried", reflect.TypeOf(V2SessionDurableEventRetried{})},
		{"session.next.compaction.started", reflect.TypeOf(V2SessionDurableEventCompactionStarted{})},
		{"session.next.compaction.ended", reflect.TypeOf(V2SessionDurableEventCompactionEnded{})},
		{"session.next.revert.staged", reflect.TypeOf(V2SessionDurableEventRevertStaged{})},
		{"session.next.revert.cleared", reflect.TypeOf(V2SessionDurableEventRevertCleared{})},
		{"session.next.revert.committed", reflect.TypeOf(V2SessionDurableEventRevertCommitted{})},
	}

	if len(tests) != 28 {
		t.Fatalf("expected 28 V2SessionDurableEvent discriminator cases (OpenAPI SessionDurableEvent.oneOf has 28 members), got %d", len(tests))
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.discriminator, func(t *testing.T) {
			t.Parallel()
			jsonStr := `{"id":"evt_1","type":"` + tc.discriminator + `","data":{}}`
			var evt V2SessionDurableEvent
			if err := evt.UnmarshalJSON([]byte(jsonStr)); err != nil {
				t.Fatalf("UnmarshalJSON: %v", err)
			}
			got := evt.AsUnion()
			if reflect.TypeOf(got) != tc.wantType {
				t.Fatalf("AsUnion: got %T, want %s", got, tc.wantType)
			}
			if gotStr := typeFieldString(got); gotStr != tc.discriminator {
				t.Errorf("Type field: got %q, want %q", gotStr, tc.discriminator)
			}
		})
	}
}

// TestV2SessionMessageDiscriminatorEnumsIsKnown asserts IsKnown() for every
// newly-typed V2SessionMessage* discriminator enum (F1): known value -> true,
// unknown value -> false.
func TestV2SessionMessageDiscriminatorEnumsIsKnown(t *testing.T) {
	t.Parallel()

	if !V2SessionMessageAgentSwitchedTypeAgentSwitched.IsKnown() {
		t.Error("V2SessionMessageAgentSwitchedTypeAgentSwitched.IsKnown() = false, want true")
	}
	if V2SessionMessageAgentSwitchedType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageAgentSwitchedType.IsKnown() = true, want false")
	}

	if !V2SessionMessageModelSwitchedTypeModelSwitched.IsKnown() {
		t.Error("V2SessionMessageModelSwitchedTypeModelSwitched.IsKnown() = false, want true")
	}
	if V2SessionMessageModelSwitchedType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageModelSwitchedType.IsKnown() = true, want false")
	}

	if !V2SessionMessageUserTypeUser.IsKnown() {
		t.Error("V2SessionMessageUserTypeUser.IsKnown() = false, want true")
	}
	if V2SessionMessageUserType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageUserType.IsKnown() = true, want false")
	}

	if !V2SessionMessageSyntheticTypeSynthetic.IsKnown() {
		t.Error("V2SessionMessageSyntheticTypeSynthetic.IsKnown() = false, want true")
	}
	if V2SessionMessageSyntheticType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageSyntheticType.IsKnown() = true, want false")
	}

	if !V2SessionMessageSystemTypeSystem.IsKnown() {
		t.Error("V2SessionMessageSystemTypeSystem.IsKnown() = false, want true")
	}
	if V2SessionMessageSystemType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageSystemType.IsKnown() = true, want false")
	}

	if !V2SessionMessageShellTypeShell.IsKnown() {
		t.Error("V2SessionMessageShellTypeShell.IsKnown() = false, want true")
	}
	if V2SessionMessageShellType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageShellType.IsKnown() = true, want false")
	}

	if !V2SessionMessageAssistantTypeAssistant.IsKnown() {
		t.Error("V2SessionMessageAssistantTypeAssistant.IsKnown() = false, want true")
	}
	if V2SessionMessageAssistantType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageAssistantType.IsKnown() = true, want false")
	}

	if !V2SessionMessageCompactionTypeCompaction.IsKnown() {
		t.Error("V2SessionMessageCompactionTypeCompaction.IsKnown() = false, want true")
	}
	if V2SessionMessageCompactionType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageCompactionType.IsKnown() = true, want false")
	}
}

// TestV2SessionAssistantContentDiscriminatorEnumsIsKnown asserts IsKnown() for
// the 3 newly-typed V2SessionMessageAssistant*Content discriminator enums (F1).
func TestV2SessionAssistantContentDiscriminatorEnumsIsKnown(t *testing.T) {
	t.Parallel()

	if !V2SessionMessageAssistantTextContentTypeText.IsKnown() {
		t.Error("V2SessionMessageAssistantTextContentTypeText.IsKnown() = false, want true")
	}
	if V2SessionMessageAssistantTextContentType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageAssistantTextContentType.IsKnown() = true, want false")
	}

	if !V2SessionMessageAssistantReasoningContentTypeReasoning.IsKnown() {
		t.Error("V2SessionMessageAssistantReasoningContentTypeReasoning.IsKnown() = false, want true")
	}
	if V2SessionMessageAssistantReasoningContentType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageAssistantReasoningContentType.IsKnown() = true, want false")
	}

	if !V2SessionMessageAssistantToolContentTypeTool.IsKnown() {
		t.Error("V2SessionMessageAssistantToolContentTypeTool.IsKnown() = false, want true")
	}
	if V2SessionMessageAssistantToolContentType("__unknown__").IsKnown() {
		t.Error("unknown V2SessionMessageAssistantToolContentType.IsKnown() = true, want false")
	}
}

// TestV2SessionDurableEventReusedEnumsIsKnown spot-checks IsKnown() on a
// sample of the reused V2EventSessionNext*Type enums as decoded through the
// V2SessionDurableEvent carrier, confirming the F1 fix's reuse strategy
// produces working IsKnown() semantics end-to-end (not just in isolation in
// v2event.go).
func TestV2SessionDurableEventReusedEnumsIsKnown(t *testing.T) {
	t.Parallel()

	var evt V2SessionDurableEvent
	if err := evt.UnmarshalJSON([]byte(`{"id":"evt_1","type":"session.next.retried","data":{}}`)); err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}
	retried, ok := evt.AsUnion().(V2SessionDurableEventRetried)
	if !ok {
		t.Fatalf("AsUnion: got %T, want V2SessionDurableEventRetried", evt.AsUnion())
	}
	if !retried.Type.IsKnown() {
		t.Errorf("V2SessionDurableEventRetried.Type.IsKnown() = false, want true (value %q)", retried.Type)
	}
	if V2EventSessionNextRetriedType("__unknown__").IsKnown() {
		t.Error("unknown V2EventSessionNextRetriedType.IsKnown() = true, want false")
	}

	var evt2 V2SessionDurableEvent
	if err := evt2.UnmarshalJSON([]byte(`{"id":"evt_2","type":"session.next.revert.committed","data":{}}`)); err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}
	committed, ok := evt2.AsUnion().(V2SessionDurableEventRevertCommitted)
	if !ok {
		t.Fatalf("AsUnion: got %T, want V2SessionDurableEventRevertCommitted", evt2.AsUnion())
	}
	if !committed.Type.IsKnown() {
		t.Errorf("V2SessionDurableEventRevertCommitted.Type.IsKnown() = false, want true (value %q)", committed.Type)
	}
}

// TestSessionErrorUnknownTypeUnmarshal covers F2: SessionErrorUnknown.Type is
// now SessionErrorUnknownType instead of a bare string.
func TestSessionErrorUnknownTypeUnmarshal(t *testing.T) {
	t.Parallel()

	var e SessionErrorUnknown
	if err := e.UnmarshalJSON([]byte(`{"type":"unknown","message":"boom"}`)); err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}
	if e.Type != SessionErrorUnknownTypeUnknown {
		t.Errorf("Type: got %q, want %q", e.Type, SessionErrorUnknownTypeUnknown)
	}
	if e.Message != "boom" {
		t.Errorf("Message: got %q, want boom", e.Message)
	}
	if !e.Type.IsKnown() {
		t.Error("SessionErrorUnknownTypeUnknown.IsKnown() = false, want true")
	}
	if SessionErrorUnknownType("__unknown__").IsKnown() {
		t.Error("unknown SessionErrorUnknownType.IsKnown() = true, want false")
	}
}

// TestSessionErrorUnknownSharedAcrossFiles verifies the F2 fix does not break
// the 3 files that embed SessionErrorUnknown as a field type
// (v2session.go: V2SessionMessageAssistant.Error, V2SessionMessageToolStateError.Error;
// v2event.go and event_global_types.go embed it too) by round-tripping it as
// a nested field inside V2SessionMessageAssistant.
func TestSessionErrorUnknownSharedAcrossFiles(t *testing.T) {
	t.Parallel()

	jsonStr := `{
		"id":"m1","time":{"created":1},"type":"assistant","agent":"build",
		"model":{"id":"gpt","providerID":"openai"},"content":[],
		"error":{"type":"unknown","message":"provider timeout"}
	}`
	var msg V2SessionMessage
	if err := msg.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatalf("UnmarshalJSON: %v", err)
	}
	assistant, ok := msg.AsUnion().(V2SessionMessageAssistant)
	if !ok {
		t.Fatalf("AsUnion: got %T, want V2SessionMessageAssistant", msg.AsUnion())
	}
	if assistant.Error.Type != SessionErrorUnknownTypeUnknown {
		t.Errorf("Error.Type: got %q, want %q", assistant.Error.Type, SessionErrorUnknownTypeUnknown)
	}
	if assistant.Error.Message != "provider timeout" {
		t.Errorf("Error.Message: got %q, want %q", assistant.Error.Message, "provider timeout")
	}
}

// TestToolTextContentAndToolFileContentTypedType covers the second-round (Q3)
// review correction of an R5 finding (阶段二 · D 任务遗漏项). R5 correctly
// identified that ToolTextContent.Type / ToolFileContent.Type were bare
// `string` and fixed it by reusing the pre-existing LLMToolContentType enum
// (which allows both "text" and "file", since it backs the *merged*
// LLMToolContent carrier one struct above). Q3 review reverts that specific
// choice: per OpenAPI, ToolTextContent.type.enum == ["text"] only and
// ToolFileContent.type.enum == ["file"] only — each is a single-value
// discriminator. Reusing the merged carrier's wider enum on the individual
// variants weakens IsKnown() (it would wrongly accept "file" on
// ToolTextContent and "text" on ToolFileContent). This mirrors the
// established golden pattern in session.go: the merged carrier [Part] uses
// the wide [PartType] enum, but each individual variant (e.g. [AgentPart],
// [TextPart], [FilePart], [ReasoningPart]) has its own dedicated
// single-value enum ([AgentPartType], [TextPartType], [FilePartType],
// [ReasoningPartType]) rather than reusing [PartType]. This fix mints
// [ToolTextContentType] and [ToolFileContentType] analogously.
func TestToolTextContentAndToolFileContentTypedType(t *testing.T) {
	t.Parallel()

	t.Run("ToolTextContent", func(t *testing.T) {
		t.Parallel()
		var c ToolTextContent
		if err := c.UnmarshalJSON([]byte(`{"type":"text","text":"hello"}`)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		if c.Type != ToolTextContentTypeText {
			t.Errorf("Type: got %q, want %q", c.Type, ToolTextContentTypeText)
		}
		if !c.Type.IsKnown() {
			t.Error("ToolTextContent.Type.IsKnown() = false, want true")
		}
	})

	t.Run("ToolFileContent", func(t *testing.T) {
		t.Parallel()
		var c ToolFileContent
		if err := c.UnmarshalJSON([]byte(`{"type":"file","uri":"file:///a.txt","mime":"text/plain"}`)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		if c.Type != ToolFileContentTypeFile {
			t.Errorf("Type: got %q, want %q", c.Type, ToolFileContentTypeFile)
		}
		if !c.Type.IsKnown() {
			t.Error("ToolFileContent.Type.IsKnown() = false, want true")
		}
	})

	t.Run("ToolTextContentTypeRejectsFileValue", func(t *testing.T) {
		// This is the concrete regression the Q3 fix guards against: with
		// R5's LLMToolContentType reuse, ToolTextContentType("file").IsKnown()
		// would have returned true (LLMToolContentType allows both "text" and
		// "file"). The dedicated ToolTextContentType must only recognize
		// "text" — "file" is a foreign value on this single-value variant.
		t.Parallel()
		if ToolTextContentType("file").IsKnown() {
			t.Error(`ToolTextContentType("file").IsKnown() = true, want false (single-value variant must reject the sibling variant's value)`)
		}
	})

	t.Run("ToolFileContentTypeRejectsTextValue", func(t *testing.T) {
		t.Parallel()
		if ToolFileContentType("text").IsKnown() {
			t.Error(`ToolFileContentType("text").IsKnown() = true, want false (single-value variant must reject the sibling variant's value)`)
		}
	})

	t.Run("ViaLLMToolContentUnionRouting", func(t *testing.T) {
		// Regression guard: routing through the LLMToolContentUnion still
		// resolves to the correctly-typed variants after the field type
		// change (i.e. this fix did not break the union's init() routing,
		// which keys off the raw JSON string, not the Go field type).
		t.Parallel()
		var text LLMToolContent
		if err := text.UnmarshalJSON([]byte(`{"type":"text","text":"hi"}`)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		textVariant, ok := text.AsUnion().(ToolTextContent)
		if !ok {
			t.Fatalf("AsUnion: got %T, want ToolTextContent", text.AsUnion())
		}
		if textVariant.Type != ToolTextContentTypeText {
			t.Errorf("ToolTextContent.Type: got %q, want %q", textVariant.Type, ToolTextContentTypeText)
		}
		// The merged LLMToolContent carrier's own .Type field still uses the
		// wide LLMToolContentType enum, unaffected by this fix.
		if text.Type != LLMToolContentTypeText {
			t.Errorf("LLMToolContent.Type: got %q, want %q", text.Type, LLMToolContentTypeText)
		}

		var file LLMToolContent
		if err := file.UnmarshalJSON([]byte(`{"type":"file","uri":"file:///a.txt","mime":"text/plain"}`)); err != nil {
			t.Fatalf("UnmarshalJSON: %v", err)
		}
		fileVariant, ok := file.AsUnion().(ToolFileContent)
		if !ok {
			t.Fatalf("AsUnion: got %T, want ToolFileContent", file.AsUnion())
		}
		if fileVariant.Type != ToolFileContentTypeFile {
			t.Errorf("ToolFileContent.Type: got %q, want %q", fileVariant.Type, ToolFileContentTypeFile)
		}
		if file.Type != LLMToolContentTypeFile {
			t.Errorf("LLMToolContent.Type: got %q, want %q", file.Type, LLMToolContentTypeFile)
		}
	})
}
