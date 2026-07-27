package opencode

import (
	"encoding/json"
	"strings"
	"testing"
)

// Aligned with OpenAPI GET /permission + JS SDK(v2) Permission.list.
// query-only params: directory, workspace
func TestPermissionListParamsQuery(t *testing.T) {
	p := PermissionListParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI POST /permission/{requestID}/reply + JS SDK(v2) Permission.reply.
// body required: reply; optional: message; query: directory, workspace
func TestPermissionReplyParamsBodyAndQuery(t *testing.T) {
	t.Run("required reply field serialized to body", func(t *testing.T) {
		p := PermissionReplyParams{
			Reply:     F(PermissionReplyParamsReplyOnce),
			Directory: F("d"),
			Workspace: F("w"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"reply":"once"`) {
			t.Errorf("reply missing from body: %s", got)
		}
		if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
			t.Errorf("query fields leaked into body: %s", got)
		}
	})

	t.Run("optional message field included", func(t *testing.T) {
		p := PermissionReplyParams{
			Reply:   F(PermissionReplyParamsReplyAlways),
			Message: F("please allow"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"reply":"always"`) {
			t.Errorf("reply missing: %s", got)
		}
		if !strings.Contains(got, `"message":"please allow"`) {
			t.Errorf("message missing: %s", got)
		}
	})

	t.Run("reject reply variant", func(t *testing.T) {
		p := PermissionReplyParams{
			Reply: F(PermissionReplyParamsReplyReject),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(b), `"reply":"reject"`) {
			t.Errorf("reject missing: %s", string(b))
		}
	})

	t.Run("query serialization", func(t *testing.T) {
		p := PermissionReplyParams{
			Reply:     F(PermissionReplyParamsReplyOnce),
			Directory: F("mydir"),
			Workspace: F("mywsp"),
		}
		got := p.URLQuery().Encode()
		want := "directory=mydir&workspace=mywsp"
		if got != want {
			t.Errorf("query got %q, want %q", got, want)
		}
	})
}

// Aligned with OpenAPI POST /session/{sessionID}/permissions/{permissionID} (deprecated).
// body required: response; query: directory, workspace
func TestPermissionRespondParamsBodyAndQuery(t *testing.T) {
	t.Run("required response field serialized to body", func(t *testing.T) {
		p := PermissionRespondParams{
			Response:  F(PermissionRespondParamsResponseOnce),
			Directory: F("d"),
			Workspace: F("w"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"response":"once"`) {
			t.Errorf("response missing from body: %s", got)
		}
		if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
			t.Errorf("query fields leaked into body: %s", got)
		}
	})

	t.Run("query serialization", func(t *testing.T) {
		p := PermissionRespondParams{
			Response:  F(PermissionRespondParamsResponseAlways),
			Directory: F("d"),
			Workspace: F("w"),
		}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("query got %q, want %q", got, want)
		}
	})
}

// Aligned with OpenAPI PermissionRequest schema.
// required: id, sessionID, permission, patterns, metadata, always; optional: tool
func TestPermissionRequestUnmarshal(t *testing.T) {
	t.Run("full object with nested metadata", func(t *testing.T) {
		raw := `{
			"id": "per_abc123",
			"sessionID": "ses_xyz",
			"permission": "bash",
			"patterns": ["*.go", "*.ts"],
			"metadata": {"key": "value", "nested": {"foo": 1}},
			"always": ["always_tool"],
			"tool": {"messageID": "msg_001", "callID": "call_001"}
		}`
		var r PermissionRequest
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if r.ID != "per_abc123" {
			t.Errorf("ID = %q", r.ID)
		}
		if r.SessionID != "ses_xyz" {
			t.Errorf("SessionID = %q", r.SessionID)
		}
		if r.Permission != "bash" {
			t.Errorf("Permission = %q", r.Permission)
		}
		if len(r.Patterns) != 2 || r.Patterns[0] != "*.go" {
			t.Errorf("Patterns = %v", r.Patterns)
		}
		if r.Metadata == nil {
			t.Error("Metadata should not be nil")
		}
		// metadata is any — should decode as map[string]interface{}
		meta, ok := r.Metadata.(map[string]interface{})
		if !ok {
			t.Errorf("Metadata type: got %T, want map[string]interface{}", r.Metadata)
		}
		if meta["key"] != "value" {
			t.Errorf("metadata[key] = %v", meta["key"])
		}
		if len(r.Always) != 1 || r.Always[0] != "always_tool" {
			t.Errorf("Always = %v", r.Always)
		}
		if r.Tool.MessageID != "msg_001" {
			t.Errorf("Tool.MessageID = %q", r.Tool.MessageID)
		}
		if r.JSON.raw == "" {
			t.Error("RawJSON not preserved")
		}
	})

	t.Run("empty metadata object", func(t *testing.T) {
		raw := `{
			"id": "per_000",
			"sessionID": "ses_000",
			"permission": "read",
			"patterns": [],
			"metadata": {},
			"always": []
		}`
		var r PermissionRequest
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if r.Metadata == nil {
			t.Error("Metadata should not be nil for empty object")
		}
		meta, ok := r.Metadata.(map[string]interface{})
		if !ok {
			t.Errorf("Metadata type: got %T, want map[string]interface{}", r.Metadata)
		}
		if len(meta) != 0 {
			t.Errorf("Metadata should be empty map, got %v", meta)
		}
	})

	t.Run("deeply nested metadata", func(t *testing.T) {
		raw := `{
			"id": "per_111",
			"sessionID": "ses_111",
			"permission": "glob",
			"patterns": ["**/*"],
			"metadata": {"level1": {"level2": {"level3": "deep"}}},
			"always": []
		}`
		var r PermissionRequest
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		meta, ok := r.Metadata.(map[string]interface{})
		if !ok {
			t.Errorf("Metadata type: got %T, want map[string]interface{}", r.Metadata)
		}
		l1, ok := meta["level1"].(map[string]interface{})
		if !ok {
			t.Errorf("level1 type: got %T", meta["level1"])
		}
		l2, ok := l1["level2"].(map[string]interface{})
		if !ok {
			t.Errorf("level2 type: got %T", l1["level2"])
		}
		if l2["level3"] != "deep" {
			t.Errorf("level3 = %v", l2["level3"])
		}
	})

	t.Run("without optional tool field", func(t *testing.T) {
		raw := `{
			"id": "per_222",
			"sessionID": "ses_222",
			"permission": "edit",
			"patterns": ["*.md"],
			"metadata": {"source": "ai"},
			"always": []
		}`
		var r PermissionRequest
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if r.Tool.MessageID != "" || r.Tool.CallID != "" {
			t.Errorf("Tool should be zero-value when absent, got %+v", r.Tool)
		}
	})
}

// Aligned with OpenAPI PermissionRequest.tool embedded object.
// required: messageID, callID
func TestPermissionRequestToolUnmarshal(t *testing.T) {
	raw := `{"messageID": "msg_xyz", "callID": "call_xyz"}`
	var t2 PermissionRequestTool
	if err := json.Unmarshal([]byte(raw), &t2); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if t2.MessageID != "msg_xyz" {
		t.Errorf("MessageID = %q", t2.MessageID)
	}
	if t2.CallID != "call_xyz" {
		t.Errorf("CallID = %q", t2.CallID)
	}
	if t2.JSON.raw == "" {
		t.Error("RawJSON not preserved")
	}
}

// Enum coverage for PermissionReplyParamsReply and PermissionRespondParamsResponse.
func TestPermissionEnumIsKnown(t *testing.T) {
	for _, v := range []PermissionReplyParamsReply{
		PermissionReplyParamsReplyOnce,
		PermissionReplyParamsReplyAlways,
		PermissionReplyParamsReplyReject,
	} {
		if !v.IsKnown() {
			t.Errorf("%q should be known", v)
		}
	}
	if PermissionReplyParamsReply("unknown").IsKnown() {
		t.Error("unknown reply should not be known")
	}

	for _, v := range []PermissionRespondParamsResponse{
		PermissionRespondParamsResponseOnce,
		PermissionRespondParamsResponseAlways,
		PermissionRespondParamsResponseReject,
	} {
		if !v.IsKnown() {
			t.Errorf("%q should be known", v)
		}
	}
	if PermissionRespondParamsResponse("unknown").IsKnown() {
		t.Error("unknown respond response should not be known")
	}
}
