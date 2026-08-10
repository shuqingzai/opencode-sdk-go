package opencode_test

import (
	"embed"
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/sst/opencode-sdk-go"
)

// review2PayloadRouteTestdata 是第二轮 SSE Reviewer 独立生成的 124 个 payload 路由用例
// （89 个 V2 event + 35 个 V1 SyncEvent），逐条验证 GlobalEvent.Payload / AsUnion() 路由。
//
//go:embed testdata/review2_payload_route.json
var review2PayloadRouteTestdata embed.FS

type review2PayloadRouteCase struct {
	Name        string          `json:"name"`
	GoType      string          `json:"goType"`
	InnerGoType string          `json:"innerGoType"`
	Payload     json.RawMessage `json:"payload"`
}

// TestReview2GlobalEventPayloadRouteCoverage 遍历 124 个 OpenAPI payload 成员，
// 对每个 type 值构造最小合法报文并验证 GlobalEvent 反序列化路由：
//   - ev.Payload != nil
//   - ev.AsUnion() 落到期望的具体 Go 类型（89 个 V2 EventListResponseEventXxx）
//   - sync 类事件路由到 SyncEventResponse，并继续 AsUnion() 落到内层 SyncEventXxx
func TestReview2GlobalEventPayloadRouteCoverage(t *testing.T) {
	data, err := review2PayloadRouteTestdata.ReadFile("testdata/review2_payload_route.json")
	if err != nil {
		t.Fatalf("read testdata/review2_payload_route.json: %v", err)
	}
	var cases []review2PayloadRouteCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("parse testdata: %v", err)
	}
	if len(cases) != 124 {
		t.Fatalf("expected 124 payload members, got %d", len(cases))
	}

	var v2Count, syncCount int
	for _, tc := range cases {
		raw := []byte(`{"directory":"d","payload":` + string(tc.Payload) + `}`)

		var ev opencode.GlobalEvent
		if err := json.Unmarshal(raw, &ev); err != nil {
			t.Errorf("[%s] type=%s unmarshal error: %v", tc.Name, tc.GoType, err)
			continue
		}
		if ev.Payload == nil {
			t.Errorf("[%s] Payload is nil, want non-nil %s", tc.Name, tc.GoType)
			continue
		}
		u := ev.AsUnion()
		if u == nil {
			t.Errorf("[%s] AsUnion() is nil", tc.Name)
			continue
		}

		gotName := reflect.TypeOf(u).Name()
		if gotName != tc.GoType {
			t.Errorf("[%s] AsUnion() resolved to %s, want %s", tc.Name, gotName, tc.GoType)
			continue
		}

		if tc.InnerGoType == "" {
			v2Count++
			continue
		}

		// sync 类事件：外层 SyncEventResponse，内层还需继续 AsUnion()
		resp, ok := u.(opencode.SyncEventResponse)
		if !ok {
			t.Errorf("[%s] expected SyncEventResponse, got %T", tc.Name, u)
			continue
		}
		inner := resp.SyncEvent.AsUnion()
		if inner == nil {
			t.Errorf("[%s] SyncEvent.AsUnion() is nil", tc.Name)
			continue
		}
		innerName := reflect.TypeOf(inner).Name()
		if innerName != tc.InnerGoType {
			t.Errorf("[%s] inner AsUnion() resolved to %s, want %s", tc.Name, innerName, tc.InnerGoType)
			continue
		}
		syncCount++
	}

	t.Logf("route coverage: %d V2 events OK, %d SyncEvents OK (total %d)", v2Count, syncCount, v2Count+syncCount)
	if v2Count != 89 {
		t.Errorf("expected 89 V2 events routed correctly, got %d", v2Count)
	}
	if syncCount != 35 {
		t.Errorf("expected 35 SyncEvents routed correctly, got %d", syncCount)
	}
}

// =============================================================================
// event_sync.go 去重后的 SSE 解码零回归（GlobalEvent → SyncEventResponse →
// 内层 union 的完整三层链路）
// =============================================================================

// TestReview2SyncEventThreeLayerChainContextUpdated 验证去重后的
// SyncEventSessionNextContextUpdated 经 GlobalEvent → SyncEventResponse →
// SyncEvent.AsUnion() 三层链路解码后，最终拿到正确的具体类型与字段值。
func TestReview2SyncEventThreeLayerChainContextUpdated(t *testing.T) {
	raw := []byte(`{"directory":"d","payload":{"type":"sync","id":"evt_outer","syncEvent":{"type":"session.next.context.updated.1","id":"evt_inner","seq":42,"aggregateID":"ses_9","data":{"timestamp":1234,"sessionID":"ses_9","messageID":"msg_7","text":"three layer chain","extra":{"a":1}}}}}`)

	var ev opencode.GlobalEvent
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("GlobalEvent unmarshal: %v", err)
	}
	if ev.Payload == nil {
		t.Fatal("Payload is nil")
	}
	resp, ok := ev.AsUnion().(opencode.SyncEventResponse)
	if !ok {
		t.Fatalf("AsUnion() = %T, want SyncEventResponse", ev.AsUnion())
	}
	inner, ok := resp.SyncEvent.AsUnion().(opencode.SyncEventSessionNextContextUpdated)
	if !ok {
		t.Fatalf("SyncEvent.AsUnion() = %T, want SyncEventSessionNextContextUpdated", resp.SyncEvent.AsUnion())
	}
	if inner.Type != opencode.SyncEventSessionNextContextUpdatedTypeSessionNextContextUpdated1 {
		t.Errorf("inner.Type = %q", inner.Type)
	}
	if inner.Seq != 42 || inner.AggregateID != "ses_9" {
		t.Errorf("inner Seq/AggregateID = %d/%q", inner.Seq, inner.AggregateID)
	}
	if inner.Data.Timestamp != 1234 || inner.Data.SessionID != "ses_9" || inner.Data.MessageID != "msg_7" {
		t.Errorf("inner.Data = %+v", inner.Data)
	}
	if inner.Data.Text != "three layer chain" {
		t.Errorf("inner.Data.Text = %q", inner.Data.Text)
	}
	if inner.Data.JSON.Timestamp.IsMissing() || inner.Data.JSON.Text.IsMissing() {
		t.Error("inner.Data.JSON metadata reported missing")
	}
	if _, ok := inner.Data.JSON.ExtraFields["extra"]; !ok {
		t.Error("inner.Data.JSON.ExtraFields should capture the unknown 'extra' field")
	}
	if !strings.Contains(inner.Data.JSON.RawJSON(), "three layer chain") {
		t.Errorf("inner.Data.RawJSON() = %q, want it to contain text", inner.Data.JSON.RawJSON())
	}
	if !strings.Contains(ev.JSON.RawJSON(), "three layer chain") {
		t.Error("GlobalEvent.RawJSON() should preserve the full original payload")
	}
}

// TestReview2SyncEventThreeLayerChainPromptAdmitted 验证去重后的
// SyncEventSessionNextPromptAdmitted 三层链路解码，含 Delivery 枚举与 Prompt。
func TestReview2SyncEventThreeLayerChainPromptAdmitted(t *testing.T) {
	raw := []byte(`{"directory":"d","payload":{"type":"sync","id":"evt_outer","syncEvent":{"type":"session.next.prompt.admitted.1","id":"evt_inner","seq":3,"aggregateID":"ses_9","data":{"timestamp":5678,"sessionID":"ses_9","messageID":"msg_8","prompt":{"text":"please admit","agents":[{"name":"reader"}]},"delivery":"steer"}}}}`)

	var ev opencode.GlobalEvent
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("GlobalEvent unmarshal: %v", err)
	}
	resp, ok := ev.AsUnion().(opencode.SyncEventResponse)
	if !ok {
		t.Fatalf("AsUnion() = %T, want SyncEventResponse", ev.AsUnion())
	}
	inner, ok := resp.SyncEvent.AsUnion().(opencode.SyncEventSessionNextPromptAdmitted)
	if !ok {
		t.Fatalf("SyncEvent.AsUnion() = %T, want SyncEventSessionNextPromptAdmitted", resp.SyncEvent.AsUnion())
	}
	if inner.Data.Timestamp != 5678 {
		t.Errorf("inner.Data.Timestamp = %d", inner.Data.Timestamp)
	}
	if inner.Data.Delivery != opencode.EventListResponseEventSessionNextPromptAdmittedPropertiesDeliverySteer {
		t.Errorf("inner.Data.Delivery = %q, want steer", inner.Data.Delivery)
	}
	if !inner.Data.Delivery.IsKnown() {
		t.Error("Delivery.IsKnown() = false for steer")
	}
	if inner.Data.Prompt.Text != "please admit" || len(inner.Data.Prompt.Agents) != 1 {
		t.Errorf("inner.Data.Prompt = %+v", inner.Data.Prompt)
	}
	if inner.Data.JSON.Delivery.IsMissing() || inner.Data.JSON.Prompt.IsMissing() {
		t.Error("inner.Data.JSON metadata reported missing")
	}
	if !strings.Contains(inner.Data.JSON.RawJSON(), "please admit") {
		t.Errorf("inner.Data.RawJSON() = %q, want prompt text", inner.Data.JSON.RawJSON())
	}
}

// TestReview2SyncEventDedupExtraFields 直接解码两个去重事件，验证未知字段进入
// ExtraFields 且不影响已知字段。
func TestReview2SyncEventDedupExtraFields(t *testing.T) {
	t.Parallel()

	ctxRaw := []byte(`{"type":"session.next.context.updated.1","id":"evt_1","seq":1,"aggregateID":"ses_1","data":{"timestamp":1,"sessionID":"ses_1","messageID":"msg_1","text":"t","futureField":"kept"}}`)
	var ctx opencode.SyncEventSessionNextContextUpdated
	if err := json.Unmarshal(ctxRaw, &ctx); err != nil {
		t.Fatalf("context.updated unmarshal: %v", err)
	}
	if ctx.Data.Text != "t" {
		t.Errorf("ctx.Data.Text = %q", ctx.Data.Text)
	}
	if _, ok := ctx.Data.JSON.ExtraFields["futureField"]; !ok {
		t.Error("ctx.Data.JSON.ExtraFields missing futureField")
	}
	if !strings.Contains(ctx.Data.JSON.RawJSON(), "kept") {
		t.Error("ctx.Data.RawJSON() should preserve the future field")
	}

	promptRaw := []byte(`{"type":"session.next.prompt.admitted.1","id":"evt_2","seq":2,"aggregateID":"ses_1","data":{"timestamp":2,"sessionID":"ses_1","messageID":"msg_2","prompt":{"text":"p"},"delivery":"queue","futureFlag":true}}`)
	var pa opencode.SyncEventSessionNextPromptAdmitted
	if err := json.Unmarshal(promptRaw, &pa); err != nil {
		t.Fatalf("prompt.admitted unmarshal: %v", err)
	}
	if pa.Data.Delivery != opencode.EventListResponseEventSessionNextPromptAdmittedPropertiesDeliveryQueue {
		t.Errorf("pa.Data.Delivery = %q", pa.Data.Delivery)
	}
	if _, ok := pa.Data.JSON.ExtraFields["futureFlag"]; !ok {
		t.Error("pa.Data.JSON.ExtraFields missing futureFlag")
	}
}

// TestReview2AliasTypeIdentity 独立实证：Go 类型别名（type A = B）不产生新类型，
// reflect.TypeFor / reflect.TypeOf 均与目标类型完全相等，因此
// apijson.RegisterUnion 的注册不受别名影响。
func TestReview2AliasTypeIdentity(t *testing.T) {
	t.Parallel()

	cases := []struct {
		alias  reflect.Type
		target reflect.Type
		name   string
	}{
		{reflect.TypeFor[opencode.SyncEventSessionNextContextUpdatedProperties](), reflect.TypeFor[opencode.EventListResponseEventSessionNextContextUpdatedProperties](), "ContextUpdatedProperties"},
		{reflect.TypeFor[opencode.SyncEventSessionNextPromptAdmittedProperties](), reflect.TypeFor[opencode.EventListResponseEventSessionNextPromptAdmittedProperties](), "PromptAdmittedProperties"},
		{reflect.TypeFor[opencode.SyncEventSessionNextPromptAdmittedDelivery](), reflect.TypeFor[opencode.EventListResponseEventSessionNextPromptAdmittedPropertiesDelivery](), "Delivery"},
	}
	for _, c := range cases {
		if c.alias != c.target {
			t.Errorf("%s: TypeFor[alias] != TypeFor[target]", c.name)
		}
	}
	if reflect.TypeFor[opencode.SyncEventSessionNextContextUpdatedProperties]() != reflect.TypeFor[opencode.EventListResponseEventSessionNextContextUpdatedProperties]() {
		t.Error("TypeOf(alias{}) != TypeFor[target]")
	}
	if reflect.TypeFor[opencode.SyncEventSessionNextPromptAdmittedProperties]() != reflect.TypeFor[opencode.EventListResponseEventSessionNextPromptAdmittedProperties]() {
		t.Error("TypeOf(alias{}) != TypeFor[target]")
	}
}

// =============================================================================
// null / 未知 / 畸形 payload 的鲁棒性对照表
// =============================================================================

// TestReview2SSERobustnessMatrix 对三条 SSE 链路（GlobalEvent / EventListResponse /
// V2Event）分别喂入 ①载体为 null ②未知 type ③缺失必填字段 ④类型错误 ⑤畸形原始值，
// 记录各自行为并断言：字段级载体任何畸形输入（含 string/number/bool/array）
// 在三条链路上都不得报错（否则 ssestream.Next 会终结整个 SSE 流）。
func TestReview2SSERobustnessMatrix(t *testing.T) {
	chains := []struct {
		label  string
		wrap   func(payload string) string
		decode func(raw []byte) (asUnion string, errText string)
	}{
		{
			label: "GlobalEvent.payload",
			wrap:  func(p string) string { return `{"directory":"d","payload":` + p + `}` },
			decode: func(raw []byte) (string, string) {
				var ev opencode.GlobalEvent
				err := json.Unmarshal(raw, &ev)
				if err != nil {
					return "", err.Error()
				}
				if ev.Payload == nil {
					return "(nil payload)", ""
				}
				return reflect.TypeOf(ev.AsUnion()).Name(), ""
			},
		},
		{
			label: "EventListResponse.properties",
			wrap:  func(p string) string { return `{"id":"x","type":"file.edited","properties":` + p + `}` },
			decode: func(raw []byte) (string, string) {
				var ev opencode.EventListResponse
				err := json.Unmarshal(raw, &ev)
				if err != nil {
					return "", err.Error()
				}
				if ev.Properties == nil {
					return "(nil properties)", ""
				}
				return reflect.TypeOf(ev.AsUnion()).Name(), ""
			},
		},
		{
			label: "V2Event.data",
			wrap:  func(p string) string { return `{"id":"x","type":"file.edited","data":` + p + `}` },
			decode: func(raw []byte) (string, string) {
				var ev opencode.V2Event
				err := json.Unmarshal(raw, &ev)
				if err != nil {
					return "", err.Error()
				}
				if ev.Data == nil {
					return "(nil data)", ""
				}
				return reflect.TypeOf(ev.AsUnion()).Name(), ""
			},
		},
	}

	inputs := []struct {
		name string
		raw  string
	}{
		{"carrier=null", `null`},
		{"unknown-type", `{"id":"x","type":"totally.unknown","properties":{}}`},
		{"missing-required", `{"type":"file.edited"}`},
		{"wrong-type-value", `{"id":"x","type":"file.edited","properties":{"filePath":123}}`},
		{"carrier=array", `[]`},
		{"carrier=string", `"not-an-object"`},
		{"carrier=number", `123`},
	}

	// whole-object 级输入：直接把整个 SSE 报文替换成非对象 JSON
	wholeInputs := []struct {
		name string
		raw  string
	}{
		{"whole=null", `null`},
		{"whole=array", `[]`},
		{"whole=string", `"not-an-object"`},
	}

	for _, ch := range chains {
		for _, in := range inputs {
			raw := []byte(ch.wrap(in.raw))
			asUnion, errText := ch.decode(raw)
			status := "OK"
			if errText != "" {
				status = "ERROR"
			}
			t.Logf("[%s] %-18s -> %-6s asUnion=%s err=%s", ch.label, in.name, status, asUnion, errText)
			if errText != "" {
				// 字段级畸形输入在三条链路上都必须容错：任何解码错误都会经
				// ssestream.Stream.Next() 终结整个 SSE 流。
				t.Errorf("[%s] %s: decode error %q would terminate the SSE stream", ch.label, in.name, errText)
			}
		}
		for _, in := range wholeInputs {
			raw := []byte(in.raw)
			asUnion, errText := ch.decode(raw)
			status := "OK"
			if errText != "" {
				status = "ERROR"
			}
			t.Logf("[%s] %-18s -> %-6s asUnion=%s err=%s", ch.label, in.name, status, asUnion, errText)
		}
	}
}
