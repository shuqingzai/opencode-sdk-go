package opencode

import (
	"encoding/json"
	"testing"
)

// =============================================================================
// SSE 容错守卫回归测试（EventListResponse.UnmarshalJSON）
//
// ssestream.Stream.Next() 中 json.Unmarshal 一旦出错即 return false，整条 SSE
// 流永久终止。OpenAPI 的 /event 响应 union 全部为 object schema，任何非对象
// 载体（null/bool/number/string/array）都不可能命中 variant。守卫在 union
// 路由前检测非对象载体并静默降级（保留 RawJSON 原始值），与 global.go 中
// GlobalEvent.UnmarshalJSON 的容错策略保持一致。
// =============================================================================

// TestEventListResponseToleranceMalformedPayloads verifies that malformed
// non-object SSE payloads (null, number, string, bool, array) do not fail
// unmarshaling — an error here would permanently terminate the whole SSE
// stream. Each degraded frame must also preserve the original payload via
// RawJSON() and leave the routed fields at their zero values, so callers can
// distinguish a malformed frame from a routed event body.
func TestEventListResponseToleranceMalformedPayloads(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		raw  string
	}{
		{name: "null", raw: "null"},
		{name: "number", raw: "123"},
		{name: "string", raw: `"hello"`},
		{name: "bool true", raw: "true"},
		{name: "bool false", raw: "false"},
		{name: "array", raw: "[1,2]"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var resp EventListResponse
			if err := json.Unmarshal([]byte(tt.raw), &resp); err != nil {
				t.Fatalf("malformed payload %s must not fail unmarshaling (would kill the whole SSE stream), got: %v", tt.raw, err)
			}
			if got := resp.JSON.RawJSON(); got != tt.raw {
				t.Errorf("RawJSON() = %q, want original payload preserved %q", got, tt.raw)
			}
			if resp.Type != "" {
				t.Errorf("Type = %q, want empty on degraded path", resp.Type)
			}
			if resp.Properties != nil {
				t.Errorf("Properties = %#v, want nil on degraded path", resp.Properties)
			}
			if resp.ID != "" {
				t.Errorf("ID = %q, want empty on degraded path", resp.ID)
			}
		})
	}
}

// TestEventListResponseToleranceKnownEventUnaffected verifies that a known
// event object still routes through the union normally after the guard was
// added: no error, Type populated with the canonical enum constant,
// Properties decoded into the typed variant struct, and RawJSON() returning
// the exact input bytes.
func TestEventListResponseToleranceKnownEventUnaffected(t *testing.T) {
	t.Parallel()
	raw := `{"id":"evt_1","type":"session.idle","properties":{"sessionID":"ses_1"}}`
	var resp EventListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("known event must still unmarshal without error: %v", err)
	}
	if resp.Type != EventListResponseTypeSessionIdle {
		t.Errorf("Type = %q, want %q (session.idle)", resp.Type, EventListResponseTypeSessionIdle)
	}
	if got := resp.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %q, want exact input %q", got, raw)
	}
	props, ok := resp.Properties.(EventListResponseEventSessionIdleProperties)
	if !ok {
		t.Fatalf("Properties = %T, want EventListResponseEventSessionIdleProperties", resp.Properties)
	}
	if props.SessionID != "ses_1" {
		t.Errorf("Properties.SessionID = %q, want ses_1", props.SessionID)
	}
}

// TestEventListResponseToleranceUnknownTypeObject verifies that an object
// carrying an unknown event type does not error out (forward compatibility:
// servers may add new event types before clients learn about them).
func TestEventListResponseToleranceUnknownTypeObject(t *testing.T) {
	t.Parallel()
	raw := `{"id":"evt_1","type":"brand.new.event","properties":{}}`
	var resp EventListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unknown event type object must not fail unmarshaling, got: %v", err)
	}
	if resp.ID != "evt_1" {
		t.Errorf("ID = %q, want evt_1", resp.ID)
	}
}

// TestEventListResponseToleranceEmptyObject verifies that an empty JSON
// object is accepted without error. Unlike the malformed payloads above, an
// empty object passes the IsObject guard and enters union routing, where the
// framework falls back to a generic variant; ID/Type stay at their zero
// values and RawJSON() still returns the exact input bytes.
func TestEventListResponseToleranceEmptyObject(t *testing.T) {
	t.Parallel()
	raw := `{}`
	var resp EventListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("empty object must not fail unmarshaling, got: %v", err)
	}
	if got := resp.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %q, want exact input %q", got, raw)
	}
	if resp.ID != "" {
		t.Errorf("ID = %q, want empty (no id field in input)", resp.ID)
	}
	if resp.Type != "" {
		t.Errorf("Type = %q, want empty (no type field in input)", resp.Type)
	}
}
