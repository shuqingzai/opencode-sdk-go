package opencode

import (
	"encoding/json"
	"testing"
)

// =============================================================================
// SSE 容错守卫回归测试（V2Event.UnmarshalJSON）
//
// ssestream.Stream.Next() 中 json.Unmarshal 一旦出错即 return false，整条 SSE
// 流永久终止。OpenAPI 的 /api/event 响应 union（88 个 type）全部为 object
// schema，任何非对象载体（null/bool/number/string/array）都不可能命中
// variant。守卫在 union 路由前检测非对象载体并静默降级（保留 RawJSON 原始
// 值），与 global.go 中 GlobalEvent.UnmarshalJSON 及 event.go 中
// EventListResponse.UnmarshalJSON 的容错策略保持一致。
// =============================================================================

// TestV2EventToleranceMalformedPayloads verifies that malformed non-object SSE
// payloads (null, number, string, bool, array) do not fail unmarshaling — an
// error here would permanently terminate the whole SSE stream. Each degraded
// frame must also preserve the original payload via RawJSON() and leave the
// routed fields at their zero values, so callers can distinguish a malformed
// frame from a routed event body.
func TestV2EventToleranceMalformedPayloads(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		raw  string
	}{
		{name: "null", raw: "null"},
		{name: "number", raw: "123"},
		{name: "string", raw: `"hello"`},
		{name: "bool-true", raw: "true"},
		{name: "bool-false", raw: "false"},
		{name: "array", raw: "[1,2]"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var evt V2Event
			if err := json.Unmarshal([]byte(tt.raw), &evt); err != nil {
				t.Fatalf("malformed payload %s must not fail unmarshaling (would kill the whole SSE stream), got: %v", tt.raw, err)
			}
			if got := evt.JSON.RawJSON(); got != tt.raw {
				t.Errorf("RawJSON() = %q, want original payload preserved %q", got, tt.raw)
			}
			if evt.ID != "" {
				t.Errorf("ID = %q, want empty on degraded path", evt.ID)
			}
			if evt.Type != "" {
				t.Errorf("Type = %q, want empty on degraded path", evt.Type)
			}
			if evt.Data != nil {
				t.Errorf("Data = %#v, want nil on degraded path", evt.Data)
			}
		})
	}
}

// TestV2EventToleranceKnownEventUnaffected verifies that a known event object
// still routes through the union normally after the guard was added: no error,
// Type populated with the canonical enum constant, ID decoded, AsUnion()
// returning the typed variant, and RawJSON() returning the exact input bytes.
func TestV2EventToleranceKnownEventUnaffected(t *testing.T) {
	t.Parallel()
	raw := `{"id":"evt_1","type":"models-dev.refreshed","data":{}}`
	var evt V2Event
	if err := json.Unmarshal([]byte(raw), &evt); err != nil {
		t.Fatalf("known event must still unmarshal without error: %v", err)
	}
	if evt.Type != V2EventTypeModelsDevRefreshed {
		t.Errorf("Type = %q, want %q (models-dev.refreshed)", evt.Type, V2EventTypeModelsDevRefreshed)
	}
	if evt.ID != "evt_1" {
		t.Errorf("ID = %q, want evt_1", evt.ID)
	}
	if got := evt.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %q, want exact input %q", got, raw)
	}
	if _, ok := evt.AsUnion().(V2EventModelsDevRefreshed); !ok {
		t.Fatalf("AsUnion() = %T, want V2EventModelsDevRefreshed", evt.AsUnion())
	}
}

// TestV2EventToleranceUnknownTypeObject verifies that an object carrying an
// unknown event type does not error out (forward compatibility: servers may
// add new event types before clients learn about them).
func TestV2EventToleranceUnknownTypeObject(t *testing.T) {
	t.Parallel()
	raw := `{"id":"evt_1","type":"brand.new.v2","data":{}}`
	var evt V2Event
	if err := json.Unmarshal([]byte(raw), &evt); err != nil {
		t.Fatalf("unknown event type object must not fail unmarshaling, got: %v", err)
	}
	if evt.ID != "evt_1" {
		t.Errorf("ID = %q, want evt_1", evt.ID)
	}
	if got := evt.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %q, want exact input %q", got, raw)
	}
}

// TestV2EventToleranceEmptyObject verifies that an empty JSON object is
// accepted without error. Unlike the malformed payloads above, an object does
// enter union routing (it may bind to the first registered variant), so only
// the error-free contract, RawJSON preservation and the scalar zero values are
// asserted here — Data is intentionally NOT asserted to be nil.
func TestV2EventToleranceEmptyObject(t *testing.T) {
	t.Parallel()
	raw := `{}`
	var evt V2Event
	if err := json.Unmarshal([]byte(raw), &evt); err != nil {
		t.Fatalf("empty object must not fail unmarshaling, got: %v", err)
	}
	if got := evt.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %q, want exact input %q", got, raw)
	}
	if evt.ID != "" {
		t.Errorf("ID = %q, want empty (no id in payload)", evt.ID)
	}
	if evt.Type != "" {
		t.Errorf("Type = %q, want empty (no type in payload)", evt.Type)
	}
}
