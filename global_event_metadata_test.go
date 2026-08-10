package opencode_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/sst/opencode-sdk-go"
)

// TestGlobalEventPayloadMissingMetadata 验证 payload 键完全缺失时：
// JSON.Payload 元数据报告 missing，解码不报错，AsUnion() 为 nil。
func TestGlobalEventPayloadMissingMetadata(t *testing.T) {
	raw := []byte(`{"directory":"d"}`)
	var ev opencode.GlobalEvent
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("unmarshal GlobalEvent: %v", err)
	}
	if !ev.JSON.Payload.IsMissing() {
		t.Errorf("JSON.Payload.IsMissing() = false, want true for absent payload")
	}
	if ev.Payload != nil {
		t.Errorf("Payload = %#v, want nil", ev.Payload)
	}
	if ev.AsUnion() != nil {
		t.Errorf("AsUnion() = %T, want nil", ev.AsUnion())
	}
}

// TestGlobalEventPayloadNullMetadata 验证 payload 显式为 null 时：
// 解码不报错，IsMissing()==false 且 IsNull()==true（区别于缺失），Raw()=="null"。
func TestGlobalEventPayloadNullMetadata(t *testing.T) {
	raw := []byte(`{"directory":"d","payload":null}`)
	var ev opencode.GlobalEvent
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("unmarshal GlobalEvent: %v", err)
	}
	if ev.JSON.Payload.IsMissing() {
		t.Error("JSON.Payload.IsMissing() = true, want false for explicit null")
	}
	if !ev.JSON.Payload.IsNull() {
		t.Error("JSON.Payload.IsNull() = false, want true for explicit null")
	}
	if ev.JSON.Payload.Raw() != "null" {
		t.Errorf("JSON.Payload.Raw() = %q, want \"null\"", ev.JSON.Payload.Raw())
	}
	if ev.Payload != nil {
		t.Errorf("Payload = %#v, want nil", ev.Payload)
	}
	if ev.AsUnion() != nil {
		t.Errorf("AsUnion() = %T, want nil", ev.AsUnion())
	}
	if !strings.Contains(ev.JSON.RawJSON(), `"payload":null`) {
		t.Errorf("RawJSON() = %q, want it to preserve the null payload", ev.JSON.RawJSON())
	}
}

// TestGlobalEventPayloadPresentMetadata 验证 payload 为合法对象时：
// IsMissing()==false、IsNull()==false，Raw() 保留原始子对象，RawJSON() 完整。
func TestGlobalEventPayloadPresentMetadata(t *testing.T) {
	raw := []byte(`{"directory":"d","payload":{"id":"x","type":"file.edited","properties":{"filePath":"/a.go"}}}`)
	var ev opencode.GlobalEvent
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("unmarshal GlobalEvent: %v", err)
	}
	if ev.JSON.Payload.IsMissing() {
		t.Error("JSON.Payload.IsMissing() = true, want false for present payload")
	}
	if ev.JSON.Payload.IsNull() {
		t.Error("JSON.Payload.IsNull() = true, want false for present object payload")
	}
	if ev.JSON.Payload.Raw() != `{"id":"x","type":"file.edited","properties":{"filePath":"/a.go"}}` {
		t.Errorf("JSON.Payload.Raw() = %q, want the raw payload object", ev.JSON.Payload.Raw())
	}
	if ev.Payload == nil {
		t.Error("Payload is nil, want routed event")
	}
	if ev.AsUnion() == nil {
		t.Error("AsUnion() is nil, want routed event")
	}
	if !strings.Contains(ev.JSON.RawJSON(), `"filePath":"/a.go"`) {
		t.Errorf("RawJSON() = %q, want full original payload", ev.JSON.RawJSON())
	}
}

// TestGlobalEventPayloadNonObjectValues 验证 payload 为非对象原始值
// （string/number/bool/array）时：不报错、Payload==nil、AsUnion()==nil，
// 但 JSON.Payload 元数据仍保留原始值（IsMissing()==false），RawJSON() 完整。
func TestGlobalEventPayloadNonObjectValues(t *testing.T) {
	inputs := []struct {
		name string
		raw  string // the raw JSON value of payload
	}{
		{"string", `"a string"`},
		{"number", `123`},
		{"bool", `true`},
		{"empty-array", `[]`},
		{"array-of-nums", `[1,2]`},
	}
	for _, in := range inputs {
		t.Run(in.name, func(t *testing.T) {
			msg := `{"directory":"d","payload":` + in.raw + `}`
			var ev opencode.GlobalEvent
			if err := json.Unmarshal([]byte(msg), &ev); err != nil {
				t.Fatalf("unmarshal GlobalEvent: %v", err)
			}
			if ev.Payload != nil {
				t.Errorf("Payload = %#v, want nil for non-object payload", ev.Payload)
			}
			if ev.AsUnion() != nil {
				t.Errorf("AsUnion() = %T, want nil for non-object payload", ev.AsUnion())
			}
			if ev.JSON.Payload.IsMissing() {
				t.Error("JSON.Payload.IsMissing() = true, want false (raw value still recorded)")
			}
			if ev.JSON.Payload.Raw() != in.raw {
				t.Errorf("JSON.Payload.Raw() = %q, want %q", ev.JSON.Payload.Raw(), in.raw)
			}
			if !strings.Contains(ev.JSON.RawJSON(), in.raw) {
				t.Errorf("RawJSON() = %q, want it to contain the raw payload %q", ev.JSON.RawJSON(), in.raw)
			}
		})
	}
}

// TestGlobalEventPayloadUnknownTypeFallback 验证 payload 为合法对象但 type
// 未知时，仍能兜底路由到具体 variant（不被非对象守卫误伤）。
func TestGlobalEventPayloadUnknownTypeFallback(t *testing.T) {
	raw := []byte(`{"directory":"d","payload":{"id":"x","type":"totally.unknown","properties":{"filePath":"/a.go"}}}`)
	var ev opencode.GlobalEvent
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("unmarshal GlobalEvent: %v", err)
	}
	if ev.Payload == nil {
		t.Fatal("Payload is nil, want fallback-routed event for unknown type")
	}
	if ev.AsUnion() == nil {
		t.Fatal("AsUnion() is nil, want fallback variant")
	}
	if ev.JSON.Payload.IsMissing() {
		t.Error("JSON.Payload.IsMissing() = true, want false")
	}
	if !strings.Contains(ev.JSON.RawJSON(), "totally.unknown") {
		t.Errorf("RawJSON() = %q, want full original payload", ev.JSON.RawJSON())
	}
}

// TestThreeChainMalformedCarrierConsistency 三链一致性测试：同样的 4 类畸形
// 载体输入（string/number/bool/array）分别喂给 GlobalEvent.payload、
// EventListResponse.properties、V2Event.data，断言三者行为一致（均不报错，
// 不杀流）。
func TestThreeChainMalformedCarrierConsistency(t *testing.T) {
	chains := []struct {
		label  string
		wrap   func(payload string) string
		decode func(raw []byte) error
	}{
		{
			label: "GlobalEvent.payload",
			wrap:  func(p string) string { return `{"directory":"d","payload":` + p + `}` },
			decode: func(raw []byte) error {
				var ev opencode.GlobalEvent
				return json.Unmarshal(raw, &ev)
			},
		},
		{
			label: "EventListResponse.properties",
			wrap:  func(p string) string { return `{"id":"x","type":"file.edited","properties":` + p + `}` },
			decode: func(raw []byte) error {
				var ev opencode.EventListResponse
				return json.Unmarshal(raw, &ev)
			},
		},
		{
			label: "V2Event.data",
			wrap:  func(p string) string { return `{"id":"x","type":"file.edited","data":` + p + `}` },
			decode: func(raw []byte) error {
				var ev opencode.V2Event
				return json.Unmarshal(raw, &ev)
			},
		},
	}

	inputs := []struct {
		name string
		raw  string
	}{
		{"string", `"a string"`},
		{"number", `123`},
		{"bool", `true`},
		{"array", `[1,2]`},
	}

	for _, ch := range chains {
		for _, in := range inputs {
			raw := []byte(ch.wrap(in.raw))
			if err := ch.decode(raw); err != nil {
				t.Errorf("[%s] carrier=%s -> unexpected error: %v (would terminate SSE stream)", ch.label, in.name, err)
			}
		}
	}
}
