package opencode

import (
	"encoding/json"
	"testing"
)

// =============================================================================
// SSE 容错守卫回归测试（V2SessionDurableEvent.UnmarshalJSON）
//
// ssestream.Stream.Next() 中 json.Unmarshal 一旦出错即 return false，整条 SSE
// 流永久终止。OpenAPI 的 SessionDurableEvent union（用于 /api/session/{id}/event
// SSE 与 history 分页）全部为 object schema（oneOf），任何非对象载体
// （null/bool/number/string/array）都不可能命中 variant。守卫在 union 路由前
// 检测非对象载体并静默降级（保留 RawJSON 原始值），与 global.go 中
// GlobalEvent.UnmarshalJSON 的容错策略保持一致。
// =============================================================================

// TestV2SessionDurableEventToleranceMalformedPayloads verifies that malformed
// non-object SSE payloads (null, number, string, bool, array) do not fail
// unmarshaling — an error here would permanently terminate the whole SSE
// stream. Each degraded frame must also preserve the original payload via
// RawJSON() and leave the union unrouted (AsUnion() == nil), so callers can
// distinguish a malformed frame from a routed event body.
func TestV2SessionDurableEventToleranceMalformedPayloads(t *testing.T) {
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
			var evt V2SessionDurableEvent
			if err := json.Unmarshal([]byte(tt.raw), &evt); err != nil {
				t.Fatalf("malformed payload %s must not fail unmarshaling (would kill the whole SSE stream), got: %v", tt.raw, err)
			}
			if got := evt.JSON.RawJSON(); got != tt.raw {
				t.Errorf("RawJSON() = %q, want original payload preserved %q", got, tt.raw)
			}
			if routed := evt.AsUnion(); routed != nil {
				t.Errorf("AsUnion() = %#v, want nil on degraded path (non-object payload cannot match any variant)", routed)
			}
		})
	}
}

// TestV2SessionDurableEventToleranceKnownEventUnaffected verifies that a known
// durable event object still routes through the union normally after the guard
// was added: no error, AsUnion() returning the typed variant struct with its
// fields decoded, and RawJSON() returning the exact input bytes.
func TestV2SessionDurableEventToleranceKnownEventUnaffected(t *testing.T) {
	t.Parallel()
	raw := `{"id":"evt_1","type":"session.next.agent.switched","durable":{"aggregateID":"x","seq":1,"version":1},"location":{"directory":"/w"},"data":{"agent":"build","messageID":"msg_1","sessionID":"ses_1","timestamp":123}}`
	var evt V2SessionDurableEvent
	if err := json.Unmarshal([]byte(raw), &evt); err != nil {
		t.Fatalf("known durable event must still unmarshal without error: %v", err)
	}
	if got := evt.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %q, want exact input %q", got, raw)
	}
	switched, ok := evt.AsUnion().(V2SessionDurableEventAgentSwitched)
	if !ok {
		t.Fatalf("AsUnion() = %T, want V2SessionDurableEventAgentSwitched", evt.AsUnion())
	}
	if switched.ID != "evt_1" {
		t.Errorf("ID = %q, want evt_1", switched.ID)
	}
	if switched.Type != V2EventSessionNextAgentSwitchedTypeSessionNextAgentSwitched {
		t.Errorf("Type = %q, want %q (session.next.agent.switched)", switched.Type, V2EventSessionNextAgentSwitchedTypeSessionNextAgentSwitched)
	}
	if switched.Durable.AggregateID != "x" || switched.Durable.Seq != 1 || switched.Durable.Version != 1 {
		t.Errorf("Durable = {aggregateID:%q seq:%d version:%d}, want {x 1 1}", switched.Durable.AggregateID, switched.Durable.Seq, switched.Durable.Version)
	}
	if switched.Data.Agent != "build" || switched.Data.SessionID != "ses_1" {
		t.Errorf("Data = {agent:%q sessionID:%q}, want {build ses_1}", switched.Data.Agent, switched.Data.SessionID)
	}
}

// TestV2SessionDurableEventToleranceUnknownTypeObject verifies that an object
// carrying an unknown durable-event type does not error out (forward
// compatibility: servers may add new event types before clients learn about
// them), and that RawJSON() still returns the exact input bytes — matching
// the assertion caliber of the EventListResponse / V2Event sibling tests.
func TestV2SessionDurableEventToleranceUnknownTypeObject(t *testing.T) {
	t.Parallel()
	raw := `{"id":"evt_1","type":"session.next.brand.new","data":{}}`
	var evt V2SessionDurableEvent
	if err := json.Unmarshal([]byte(raw), &evt); err != nil {
		t.Fatalf("unknown durable-event type object must not fail unmarshaling, got: %v", err)
	}
	if got := evt.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %q, want exact input %q", got, raw)
	}
}

// TestV2SessionDurableEventToleranceEmptyObject verifies that an empty JSON
// object is accepted without error and RawJSON() still returns the exact
// input bytes. The union carrier struct has no scalar fields of its own, so
// no zero-value field assertions are possible here (unlike the
// EventListResponse / V2Event siblings) — whether {} binds to a variant is
// deliberately left unasserted to stay independent of variant registration
// order.
func TestV2SessionDurableEventToleranceEmptyObject(t *testing.T) {
	t.Parallel()
	raw := `{}`
	var evt V2SessionDurableEvent
	if err := json.Unmarshal([]byte(raw), &evt); err != nil {
		t.Fatalf("empty object must not fail unmarshaling, got: %v", err)
	}
	if got := evt.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON() = %q, want exact input %q", got, raw)
	}
}
