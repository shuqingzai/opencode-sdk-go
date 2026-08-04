package opencode

import (
	"context"
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
	metadata, ok := provider.Metadata.(map[string]interface{})
	if !ok {
		t.Fatalf("expected Metadata to be map[string]interface{}, got %T", provider.Metadata)
	}
	if metadata["key"] != "value" {
		t.Errorf("expected metadata key=value, got %v", metadata["key"])
	}
	resultMetadata, ok := provider.ResultMetadata.(map[string]interface{})
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
