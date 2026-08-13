package opencode

import (
	"encoding/json"
	"testing"

	"github.com/sst/opencode-sdk-go/internal/apijson"
)

// sessionStatusMapPayload 覆盖 OpenAPI `/session/status` 的 200 响应
// （additionalProperties -> $ref SessionStatus）的全部三个 anyOf variant，
// 外加一个服务端将来可能新增的未知 variant，用于验证前向兼容。
const sessionStatusMapPayload = `{
		"ses_idle":    {"type": "idle"},
		"ses_busy":    {"type": "busy"},
		"ses_retry":   {"type": "retry", "attempt": 3, "message": "rate limited", "next": 1700000000},
		"ses_unknown": {"type": "compacting"}
	}`

// assertSessionStatusMap 逐条校验载体结构体字段、合并枚举与 AsUnion() 路由结果。
func assertSessionStatusMap(t *testing.T, m SessionStatusMap) {
	t.Helper()

	if len(m) != 4 {
		t.Fatalf("expected 4 entries, got %d", len(m))
	}

	t.Run("idle", func(t *testing.T) {
		s, ok := m["ses_idle"]
		if !ok {
			t.Fatal(`key "ses_idle" not found in map`)
		}
		if s.Type != SessionStatusTypeIdle {
			t.Errorf("Type = %q, want %q", s.Type, SessionStatusTypeIdle)
		}
		if !s.Type.IsKnown() {
			t.Errorf("Type %q is not known", s.Type)
		}
		v, ok := s.AsUnion().(SessionStatusIdle)
		if !ok {
			t.Fatalf("AsUnion() = %T, want SessionStatusIdle", s.AsUnion())
		}
		if v.Type != SessionStatusIdleTypeIdle {
			t.Errorf("variant Type = %q, want %q", v.Type, SessionStatusIdleTypeIdle)
		}
		if !v.Type.IsKnown() {
			t.Errorf("variant Type %q is not known", v.Type)
		}
	})

	t.Run("busy", func(t *testing.T) {
		s, ok := m["ses_busy"]
		if !ok {
			t.Fatal(`key "ses_busy" not found in map`)
		}
		if s.Type != SessionStatusTypeBusy {
			t.Errorf("Type = %q, want %q", s.Type, SessionStatusTypeBusy)
		}
		if !s.Type.IsKnown() {
			t.Errorf("Type %q is not known", s.Type)
		}
		v, ok := s.AsUnion().(SessionStatusBusy)
		if !ok {
			t.Fatalf("AsUnion() = %T, want SessionStatusBusy", s.AsUnion())
		}
		if v.Type != SessionStatusBusyTypeBusy {
			t.Errorf("variant Type = %q, want %q", v.Type, SessionStatusBusyTypeBusy)
		}
		if !v.Type.IsKnown() {
			t.Errorf("variant Type %q is not known", v.Type)
		}
	})

	t.Run("retry", func(t *testing.T) {
		s, ok := m["ses_retry"]
		if !ok {
			t.Fatal(`key "ses_retry" not found in map`)
		}
		if s.Type != SessionStatusTypeRetry {
			t.Errorf("Type = %q, want %q", s.Type, SessionStatusTypeRetry)
		}
		// retry 专属字段必须 port 到载体结构体上
		if s.Attempt != 3 {
			t.Errorf("Attempt = %d, want 3", s.Attempt)
		}
		if s.Message != "rate limited" {
			t.Errorf("Message = %q, want %q", s.Message, "rate limited")
		}
		if s.Next != 1700000000 {
			t.Errorf("Next = %d, want 1700000000", s.Next)
		}
		// action 未下发，JSON 元数据必须如实反映缺失
		if !s.JSON.Action.IsMissing() {
			t.Errorf("JSON.Action should be missing, got raw=%s", s.JSON.Action.Raw())
		}

		v, ok := s.AsUnion().(SessionStatusRetry)
		if !ok {
			t.Fatalf("AsUnion() = %T, want SessionStatusRetry", s.AsUnion())
		}
		if v.Type != SessionStatusRetryTypeRetry {
			t.Errorf("variant Type = %q, want %q", v.Type, SessionStatusRetryTypeRetry)
		}
		if !v.Type.IsKnown() {
			t.Errorf("variant Type %q is not known", v.Type)
		}
		if v.Attempt != 3 {
			t.Errorf("variant Attempt = %d, want 3", v.Attempt)
		}
		if v.Message != "rate limited" {
			t.Errorf("variant Message = %q, want %q", v.Message, "rate limited")
		}
		if v.Next != 1700000000 {
			t.Errorf("variant Next = %d, want 1700000000", v.Next)
		}
	})

	// 服务端新增 variant 时不得导致解码失败：type 原样保留且 IsKnown() 为 false
	t.Run("unknown", func(t *testing.T) {
		s, ok := m["ses_unknown"]
		if !ok {
			t.Fatal(`key "ses_unknown" not found in map`)
		}
		if s.Type != "compacting" {
			t.Errorf("Type = %q, want %q", s.Type, "compacting")
		}
		if s.Type.IsKnown() {
			t.Errorf("Type %q should not be known", s.Type)
		}
	})
}

// TestSessionStatusMapUnmarshal 走 encoding/json 路径（SDK 使用方直接反序列化）。
func TestSessionStatusMapUnmarshal(t *testing.T) {
	t.Parallel()

	var m SessionStatusMap
	if err := json.Unmarshal([]byte(sessionStatusMapPayload), &m); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	assertSessionStatusMap(t, m)
}

// TestSessionStatusMapUnmarshalAPIJSON 走 apijson 路径（SDK 实际响应解码路径）。
func TestSessionStatusMapUnmarshalAPIJSON(t *testing.T) {
	t.Parallel()

	var m SessionStatusMap
	if err := apijson.UnmarshalRoot([]byte(sessionStatusMapPayload), &m); err != nil {
		t.Fatalf("apijson.UnmarshalRoot failed: %v", err)
	}
	assertSessionStatusMap(t, m)
}

// TestSessionStatusRetryAction 校验 retry variant 的可选 action 对象（OpenAPI 中
// link 为可选，其余五个字段必填）。
func TestSessionStatusRetryAction(t *testing.T) {
	t.Parallel()

	data := `{"type":"retry","attempt":2,"message":"overloaded","next":1700000001,` +
		`"action":{"reason":"rate limited","provider":"anthropic","title":"Retry",` +
		`"message":"try again later","label":"Retry now"}}`

	var s SessionStatus
	if err := json.Unmarshal([]byte(data), &s); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if s.Type != SessionStatusTypeRetry {
		t.Fatalf("Type = %q, want %q", s.Type, SessionStatusTypeRetry)
	}
	if s.Action.Reason != "rate limited" {
		t.Errorf("Action.Reason = %q, want %q", s.Action.Reason, "rate limited")
	}
	if s.Action.Provider != "anthropic" {
		t.Errorf("Action.Provider = %q, want %q", s.Action.Provider, "anthropic")
	}
	if s.Action.Title != "Retry" {
		t.Errorf("Action.Title = %q, want %q", s.Action.Title, "Retry")
	}
	if s.Action.Message != "try again later" {
		t.Errorf("Action.Message = %q, want %q", s.Action.Message, "try again later")
	}
	if s.Action.Label != "Retry now" {
		t.Errorf("Action.Label = %q, want %q", s.Action.Label, "Retry now")
	}
	// link 为 OpenAPI 可选字段，未下发时必须是零值且元数据标记缺失
	if s.Action.Link != "" {
		t.Errorf("Action.Link = %q, want empty", s.Action.Link)
	}
	if !s.Action.JSON.Link.IsMissing() {
		t.Errorf("Action.JSON.Link should be missing, got raw=%s", s.Action.JSON.Link.Raw())
	}

	v, ok := s.AsUnion().(SessionStatusRetry)
	if !ok {
		t.Fatalf("AsUnion() = %T, want SessionStatusRetry", s.AsUnion())
	}
	if v.Action.Provider != "anthropic" {
		t.Errorf("variant Action.Provider = %q, want %q", v.Action.Provider, "anthropic")
	}
}

// TestSessionStatusRetryActionLink 校验 action.link 有值时正确解码到载体与 variant。
func TestSessionStatusRetryActionLink(t *testing.T) {
	t.Parallel()

	data := `{"type":"retry","attempt":1,"message":"overloaded","next":1700000002,` +
		`"action":{"reason":"quota","provider":"openai","title":"Upgrade",` +
		`"message":"quota exceeded","label":"Upgrade plan","link":"https://example.com/billing"}}`

	var s SessionStatus
	if err := json.Unmarshal([]byte(data), &s); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	const wantLink = "https://example.com/billing"
	if s.Action.Link != wantLink {
		t.Errorf("Action.Link = %q, want %q", s.Action.Link, wantLink)
	}
	if s.Action.JSON.Link.IsMissing() {
		t.Error("Action.JSON.Link should be present")
	}

	v, ok := s.AsUnion().(SessionStatusRetry)
	if !ok {
		t.Fatalf("AsUnion() = %T, want SessionStatusRetry", s.AsUnion())
	}
	if v.Action.Link != wantLink {
		t.Errorf("variant Action.Link = %q, want %q", v.Action.Link, wantLink)
	}
}

// TestSessionStatusRawJSON 校验载体结构体保留服务端原始报文。
func TestSessionStatusRawJSON(t *testing.T) {
	t.Parallel()

	raw := `{"type":"busy"}`
	var s SessionStatus
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if got := s.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %s, want %s", got, raw)
	}
}

// TestSessionStatusUnknownType 校验服务端新增 variant 时不会打断解码，
// 未知 type 原样保留在载体字段与原始报文中。
func TestSessionStatusUnknownType(t *testing.T) {
	t.Parallel()

	raw := `{"type":"compacting"}`
	var s SessionStatus
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if s.Type != "compacting" {
		t.Errorf("Type = %q, want %q", s.Type, "compacting")
	}
	if s.Type.IsKnown() {
		t.Errorf("Type %q should not be known", s.Type)
	}
	if got := s.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %s, want %s", got, raw)
	}
}

func TestSessionStatusMapUnmarshalInvalid(t *testing.T) {
	t.Parallel()

	// Not a JSON object
	data := `["not", "an", "object"]`
	var m SessionStatusMap
	err := json.Unmarshal([]byte(data), &m)
	if err == nil {
		t.Error("expected error for non-object JSON, got nil")
	}
}

func TestSessionStatusMapUnmarshalEmpty(t *testing.T) {
	t.Parallel()

	data := `{}`
	var m SessionStatusMap
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(m) != 0 {
		t.Errorf("expected empty map, got %d entries", len(m))
	}
}
