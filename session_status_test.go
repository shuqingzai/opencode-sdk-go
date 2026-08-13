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

// =============================================================================
// omitzero 契约
//
// OpenAPI 中 `SessionStatus.retry.action` 与 `action.link` 均为可选字段。本 SDK
// 的 Response 一律用值类型（全仓库 Response 指针数为 0），可选性由两部分共同承担：
//   - 解码侧：xxxJSON 元数据的 IsMissing()/IsNull()/Raw() 分辨「缺失」与「显式零值」；
//   - 编码侧：json tag 的 `omitzero`（Go 1.24+ stdlib 支持；Response 结构体未定义
//     MarshalJSON，json.Marshal 走 stdlib）保证字段缺失时不产出 `"action":{}`。
//
// 两者合起来等价于指针语义，且不会破坏 apijson.Port（指针会因
// `reflect.Set: value of type X is not assignable to type *X` 而 panic）。
// 以下测试锁死该契约，防止 `omitzero` 被误当作无效 tag 删除。
// =============================================================================

// marshalKeys 序列化后返回顶层键集合，避免依赖字段顺序。
func marshalKeys(t *testing.T, v any) map[string]json.RawMessage {
	t.Helper()

	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	var keys map[string]json.RawMessage
	if err := json.Unmarshal(data, &keys); err != nil {
		t.Fatalf("re-decode marshaled output %s failed: %v", data, err)
	}
	return keys
}

func hasKey(keys map[string]json.RawMessage, name string) bool {
	_, ok := keys[name]
	return ok
}

// TestSessionStatusOmitzeroAction 锁定：可选 action 缺失时不得出现在序列化结果中，
// 存在时必须原样保留 —— 载体结构体与 retry variant 行为一致。
func TestSessionStatusOmitzeroAction(t *testing.T) {
	t.Parallel()

	const (
		withoutAction = `{"type":"retry","attempt":1,"message":"m","next":2}`
		withAction    = `{"type":"retry","attempt":1,"message":"m","next":2,` +
			`"action":{"reason":"r","provider":"p","title":"t","message":"m2","label":"l"}}`
	)

	t.Run("carrier/absent", func(t *testing.T) {
		var s SessionStatus
		if err := json.Unmarshal([]byte(withoutAction), &s); err != nil {
			t.Fatalf("decode failed: %v", err)
		}
		keys := marshalKeys(t, s)
		if hasKey(keys, "action") {
			t.Errorf(`action 缺失时序列化结果不应含 "action" 键，实际含 %s`, keys["action"])
		}
		// 其余必填字段必须仍在
		for _, k := range []string{"type", "attempt", "message", "next"} {
			if !hasKey(keys, k) {
				t.Errorf("required key %q missing from marshaled output", k)
			}
		}
	})

	t.Run("carrier/present", func(t *testing.T) {
		var s SessionStatus
		if err := json.Unmarshal([]byte(withAction), &s); err != nil {
			t.Fatalf("decode failed: %v", err)
		}
		keys := marshalKeys(t, s)
		if !hasKey(keys, "action") {
			t.Fatal(`action 有值时序列化结果必须含 "action" 键`)
		}
		var got SessionStatusRetryAction
		if err := json.Unmarshal(keys["action"], &got); err != nil {
			t.Fatalf("decode marshaled action failed: %v", err)
		}
		if got.Provider != "p" || got.Label != "l" {
			t.Errorf("round-tripped action = %+v, want Provider=p Label=l", got)
		}
	})

	t.Run("variant/absent", func(t *testing.T) {
		var v SessionStatusRetry
		if err := json.Unmarshal([]byte(withoutAction), &v); err != nil {
			t.Fatalf("decode failed: %v", err)
		}
		if keys := marshalKeys(t, v); hasKey(keys, "action") {
			t.Errorf(`variant: action 缺失时不应含 "action" 键，实际含 %s`, keys["action"])
		}
	})

	t.Run("variant/present", func(t *testing.T) {
		var v SessionStatusRetry
		if err := json.Unmarshal([]byte(withAction), &v); err != nil {
			t.Fatalf("decode failed: %v", err)
		}
		if keys := marshalKeys(t, v); !hasKey(keys, "action") {
			t.Error(`variant: action 有值时必须含 "action" 键`)
		}
	})
}

// TestSessionStatusOmitzeroActionExplicitEmpty 锁定往返保真度：服务端显式下发了
// 一个字段全为空串的 action 时，它是「存在」而非「缺失」，序列化必须保留该键。
// 这是值类型 + omitzero 相对裸值类型的关键差异，也是指针方案的唯一卖点。
func TestSessionStatusOmitzeroActionExplicitEmpty(t *testing.T) {
	t.Parallel()

	const explicitEmpty = `{"type":"retry","attempt":1,"message":"m","next":2,` +
		`"action":{"reason":"","provider":"","title":"","message":"","label":""}}`

	var s SessionStatus
	if err := json.Unmarshal([]byte(explicitEmpty), &s); err != nil {
		t.Fatalf("decode failed: %v", err)
	}

	// 解码侧：元数据必须把「显式下发的全零 action」判为存在
	if s.JSON.Action.IsMissing() {
		t.Error("显式下发的 action 不应被判为 missing")
	}
	// 编码侧：必须保留该键，否则往返丢失「服务端确实发了 action」这一事实
	if keys := marshalKeys(t, s); !hasKey(keys, "action") {
		t.Error(`显式下发的全零 action 必须保留 "action" 键，否则往返丢失信息`)
	}
}

// TestSessionStatusOmitzeroLink 锁定嵌套可选字段 action.link 的同一契约。
func TestSessionStatusOmitzeroLink(t *testing.T) {
	t.Parallel()

	const base = `{"reason":"r","provider":"p","title":"t","message":"m","label":"l"`

	t.Run("absent", func(t *testing.T) {
		var a SessionStatusRetryAction
		if err := json.Unmarshal([]byte(base+`}`), &a); err != nil {
			t.Fatalf("decode failed: %v", err)
		}
		if !a.JSON.Link.IsMissing() {
			t.Error("link 未下发时 JSON.Link 应为 missing")
		}
		if keys := marshalKeys(t, a); hasKey(keys, "link") {
			t.Errorf(`link 缺失时不应含 "link" 键，实际含 %s`, keys["link"])
		}
	})

	t.Run("present", func(t *testing.T) {
		var a SessionStatusRetryAction
		if err := json.Unmarshal([]byte(base+`,"link":"https://example.com/billing"}`), &a); err != nil {
			t.Fatalf("decode failed: %v", err)
		}
		keys := marshalKeys(t, a)
		if !hasKey(keys, "link") {
			t.Fatal(`link 有值时必须含 "link" 键`)
		}
		if string(keys["link"]) != `"https://example.com/billing"` {
			t.Errorf("link = %s, want %q", keys["link"], "https://example.com/billing")
		}
	})
}

// TestSessionStatusOptionalPresenceContract 锁定「不用指针也能分辨可选字段状态」这一
// 架构前提：缺失 / 显式 null / 显式零值三态必须互相可分辨。
func TestSessionStatusOptionalPresenceContract(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		raw           string
		wantMissing   bool
		wantNull      bool
		wantMarshaled bool
	}{
		{
			name:          "absent",
			raw:           `{"type":"retry","attempt":1,"message":"m","next":2}`,
			wantMissing:   true,
			wantMarshaled: false,
		},
		{
			name:          "explicit-null",
			raw:           `{"type":"retry","attempt":1,"message":"m","next":2,"action":null}`,
			wantNull:      true,
			wantMarshaled: false,
		},
		{
			name: "explicit-value",
			raw: `{"type":"retry","attempt":1,"message":"m","next":2,` +
				`"action":{"reason":"r","provider":"p","title":"t","message":"m2","label":"l"}}`,
			wantMarshaled: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s SessionStatus
			if err := json.Unmarshal([]byte(tc.raw), &s); err != nil {
				t.Fatalf("decode failed: %v", err)
			}
			if got := s.JSON.Action.IsMissing(); got != tc.wantMissing {
				t.Errorf("JSON.Action.IsMissing() = %v, want %v (raw=%q)",
					got, tc.wantMissing, s.JSON.Action.Raw())
			}
			if tc.wantNull && !s.JSON.Action.IsNull() {
				t.Errorf("JSON.Action.IsNull() = false, want true (raw=%q)", s.JSON.Action.Raw())
			}
			if got := hasKey(marshalKeys(t, s), "action"); got != tc.wantMarshaled {
				t.Errorf(`marshal 是否含 "action" = %v, want %v`, got, tc.wantMarshaled)
			}
		})
	}
}
