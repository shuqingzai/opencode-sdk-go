package opencode_test

import (
	"embed"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
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
}

// =============================================================================
// null / 未知 / 畸形 payload 的鲁棒性对照表
// =============================================================================

// TestReview2SSERobustnessMatrix 对三条 SSE 链路（GlobalEvent / EventListResponse /
// V2Event）分别喂入 ①载体为 null ②未知 type ③缺失必填字段 ④类型错误 ⑤畸形原始值，
// 以及整帧级非对象输入（null/array/string），并断言两个口径：
//   - 字段级载体任何畸形输入（含 string/number/bool/array）都不得报错；
//   - 整帧级非对象载体同样不得报错（IsObject 守卫静默降级，与
//     TestReview2SSEBareScalarWholeEventTolerated /
//     TestReview2SSEBareArrayWholeEventTolerated 锁定的行为一致）。
//
// 任何解码错误都会经 ssestream.Next 终结整个 SSE 流，故一律视为缺陷。
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
			if errText != "" {
				// 整帧级非对象输入在三条链路上都必须容错（IsObject 守卫 /
				// gjson 路径取值静默降级），任何解码错误都会终结整条流。
				t.Errorf("[%s] %s: decode error %q would terminate the SSE stream", ch.label, in.name, errText)
			}
		}
	}
}

// =============================================================================
// Q2 独立复核角度 3：真实字节流层面的 SSE 容错（不止单个 JSON 反序列化）。
// newSSEResponse / sseEventFrame 复用 session_status_sse_test.go 中已有的
// package-level helper（同 opencode_test 包，未修改该文件）。
// =============================================================================

// sseCountEvents 把 frames 拼成一个 text/event-stream 响应体，通过真实
// ssestream.Stream[T] 消费，返回收到的事件数与 stream.Err()。
func sseCountEventListResponse(frames ...string) (int, error) {
	resp := newSSEResponse(frames...)
	stream := ssestream.NewStream[opencode.EventListResponse](ssestream.NewDecoder(resp), nil)
	var n int
	for stream.Next() {
		n++
	}
	return n, stream.Err()
}

func sseCountGlobalEvent(frames ...string) (int, error) {
	resp := newSSEResponse(frames...)
	stream := ssestream.NewStream[opencode.GlobalEvent](ssestream.NewDecoder(resp), nil)
	var n int
	for stream.Next() {
		n++
	}
	return n, stream.Err()
}

func sseCountV2Event(frames ...string) (int, error) {
	resp := newSSEResponse(frames...)
	stream := ssestream.NewStream[opencode.V2Event](ssestream.NewDecoder(resp), nil)
	var n int
	for stream.Next() {
		n++
	}
	return n, stream.Err()
}

// TestReview2SSEMultiLineDataFold 验证 SSE 规范的多行 data: 折叠
// （连续多条 "data:" 行按 \n 拼接成一个 JSON 文档）在三条链路上都能正确
// 解码——拼接产生的裸换行落在 JSON token 之间（合法空白），不破坏 JSON。
func TestReview2SSEMultiLineDataFold(t *testing.T) {
	t.Parallel()

	elrFrame := "event: message\n" +
		`data: {"id":"evt_1","type":"file.edited",` + "\n" +
		`data: "properties":{"file":"/multi.go"}}` + "\n\n"
	n, err := sseCountEventListResponse(elrFrame)
	if err != nil {
		t.Fatalf("EventListResponse multi-line data fold: unexpected error %v", err)
	}
	if n != 1 {
		t.Fatalf("EventListResponse multi-line data fold: got %d events, want 1", n)
	}

	geFrame := "event: message\n" +
		`data: {"directory":"d","payload":{"id":"evt_1","type":"file.edited",` + "\n" +
		`data: "data":{"file":"/multi.go"}}}` + "\n\n"
	n, err = sseCountGlobalEvent(geFrame)
	if err != nil {
		t.Fatalf("GlobalEvent multi-line data fold: unexpected error %v", err)
	}
	if n != 1 {
		t.Fatalf("GlobalEvent multi-line data fold: got %d events, want 1", n)
	}

	v2Frame := "event: message\n" +
		`data: {"id":"evt_1","type":"file.edited",` + "\n" +
		`data: "data":{"file":"/multi.go"}}` + "\n\n"
	n, err = sseCountV2Event(v2Frame)
	if err != nil {
		t.Fatalf("V2Event multi-line data fold: unexpected error %v", err)
	}
	if n != 1 {
		t.Fatalf("V2Event multi-line data fold: got %d events, want 1", n)
	}
}

// TestReview2SSEOversizedDataDoesNotOverflowScanner 验证 >64KB 的单个 data
// 帧（bufio.Scanner 默认 token 上限是 64KB）不会导致
// bufio.Scanner: token too long——ssestream.NewDecoder 已把 buffer 上限设为
// bufio.MaxScanTokenSize<<9（32MB），三条链路均需覆盖。
func TestReview2SSEOversizedDataDoesNotOverflowScanner(t *testing.T) {
	t.Parallel()
	bigValue := strings.Repeat("a", 200*1024) // 200KB，远超 64KB 默认上限

	elrPayload := fmt.Sprintf(`{"id":"evt_1","type":"file.edited","properties":{"file":"/%s.go"}}`, bigValue)
	if n, err := sseCountEventListResponse(sseEventFrame([]byte(elrPayload))); err != nil || n != 1 {
		t.Fatalf("EventListResponse oversized data: n=%d err=%v, want n=1 err=nil", n, err)
	}

	gePayload := fmt.Sprintf(`{"directory":"d","payload":{"id":"evt_1","type":"file.edited","data":{"file":"/%s.go"}}}`, bigValue)
	if n, err := sseCountGlobalEvent(sseEventFrame([]byte(gePayload))); err != nil || n != 1 {
		t.Fatalf("GlobalEvent oversized data: n=%d err=%v, want n=1 err=nil", n, err)
	}

	v2Payload := fmt.Sprintf(`{"id":"evt_1","type":"file.edited","data":{"file":"/%s.go"}}`, bigValue)
	if n, err := sseCountV2Event(sseEventFrame([]byte(v2Payload))); err != nil || n != 1 {
		t.Fatalf("V2Event oversized data: n=%d err=%v, want n=1 err=nil", n, err)
	}
}

// TestReview2SSEConsecutiveMalformedEvents 验证*连续多条*（不止中间夹一个）
// 畸形事件不会累积杀死流：5 个事件里第 2/3/4 连续三个都是字段级畸形载体，
// 第 1 和第 5 个合法事件都必须完整送达。
func TestReview2SSEConsecutiveMalformedEvents(t *testing.T) {
	t.Parallel()

	t.Run("EventListResponse.properties", func(t *testing.T) {
		n, err := sseCountEventListResponse(
			sseEventFrame([]byte(`{"id":"evt_1","type":"file.edited","properties":{"file":"/a.go"}}`)),
			sseEventFrame([]byte(`{"id":"evt_2","type":"file.edited","properties":42}`)),
			sseEventFrame([]byte(`{"id":"evt_3","type":"file.edited","properties":"str"}`)),
			sseEventFrame([]byte(`{"id":"evt_4","type":"file.edited","properties":true}`)),
			sseEventFrame([]byte(`{"id":"evt_5","type":"file.edited","properties":{"file":"/b.go"}}`)),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if n != 5 {
			t.Fatalf("got %d events, want 5 (3 consecutive malformed events must not kill the stream)", n)
		}
	})

	t.Run("GlobalEvent.payload", func(t *testing.T) {
		n, err := sseCountGlobalEvent(
			sseEventFrame([]byte(`{"directory":"d","payload":{"id":"evt_1","type":"file.edited","data":{"file":"/a.go"}}}`)),
			sseEventFrame([]byte(`{"directory":"d","payload":42}`)),
			sseEventFrame([]byte(`{"directory":"d","payload":"str"}`)),
			sseEventFrame([]byte(`{"directory":"d","payload":true}`)),
			sseEventFrame([]byte(`{"directory":"d","payload":{"id":"evt_5","type":"file.edited","data":{"file":"/b.go"}}}`)),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if n != 5 {
			t.Fatalf("got %d events, want 5 (3 consecutive malformed payloads must not kill the stream)", n)
		}
	})

	t.Run("V2Event.data", func(t *testing.T) {
		n, err := sseCountV2Event(
			sseEventFrame([]byte(`{"id":"evt_1","type":"file.edited","data":{"file":"/a.go"}}`)),
			sseEventFrame([]byte(`{"id":"evt_2","type":"file.edited","data":42}`)),
			sseEventFrame([]byte(`{"id":"evt_3","type":"file.edited","data":"str"}`)),
			sseEventFrame([]byte(`{"id":"evt_4","type":"file.edited","data":true}`)),
			sseEventFrame([]byte(`{"id":"evt_5","type":"file.edited","data":{"file":"/b.go"}}`)),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if n != 5 {
			t.Fatalf("got %d events, want 5 (3 consecutive malformed data fields must not kill the stream)", n)
		}
	})
}

// TestReview2SSEBareArrayWholeEventTolerated 记录一个实测事实：整条
// SSE 事件的 data: 若是裸数组 `[1,2,3]`（而不是 object），三条链路都不会报错。
// 历史上这归因于 gjson 把 JSON array 和 object 统一归为 gjson.JSON 类型、union
// 路由的 TypeFilter（全部注册为 gjson.JSON）恰好也能匹配数组；现整条事件级的
// 非对象载体（含数组）统一在各自 UnmarshalJSON 入口被 IsObject 守卫拦截并静默
// 降级（跳过 union 路由），可观察行为不变：不报错，帧仍作为零值事件派发（n=1）。
// 裸 null/number/string/bool 同样静默降级（见
// TestReview2SSEBareScalarWholeEventTolerated），非对象载体行为已全面统一，
// 供后续回归对比。
func TestReview2SSEBareArrayWholeEventTolerated(t *testing.T) {
	t.Parallel()
	if n, err := sseCountEventListResponse(sseEventFrame([]byte(`[1,2,3]`))); err != nil || n != 1 {
		t.Errorf("EventListResponse bare array whole event: n=%d err=%v, want n=1 err=nil", n, err)
	}
	if n, err := sseCountV2Event(sseEventFrame([]byte(`[1,2,3]`))); err != nil || n != 1 {
		t.Errorf("V2Event bare array whole event: n=%d err=%v, want n=1 err=nil", n, err)
	}
	if n, err := sseCountGlobalEvent(sseEventFrame([]byte(`[1,2,3]`))); err != nil || n != 1 {
		t.Errorf("GlobalEvent bare array whole event: n=%d err=%v, want n=1 err=nil", n, err)
	}
}

// TestReview2SSEBareScalarWholeEventTolerated 三条链路对"整条事件本身就是裸
// 标量"（不是 object）的实测行为：
//   - GlobalEvent：自定义 UnmarshalJSON 走 apijson.UnmarshalRoot（gjson 按路径
//     取字段），根节点不是 object 时具名字段全部取不到值，静默产出零值，
//     不报错、不杀流。
//   - EventListResponse / V2Event：顶层就是 Union carrier。历史上对
//     null/number/string/bool 会返回 "was not able to coerce type as union"
//     并经 ssestream.Stream.Next() 杀死整条流；现已在各自 UnmarshalJSON 入口
//     增加 IsObject 守卫，非对象载体跳过 union 路由、静默降级 return nil，
//     RawJSON() 仍保留原始报文——与 GlobalEvent 链路容错策略完全一致。
//
// 本用例锁定修复后的统一行为：裸标量帧被静默降级——仍作为零值事件派发（n=1）、
// 流继续存活，其后的合法事件照常送达（n=2）；三条链路的降级帧都必须经 RawJSON()
// 保留原始报文。
func TestReview2SSEBareScalarWholeEventTolerated(t *testing.T) {
	t.Parallel()
	// gjson 将 true/false 视为不同类型（gjson.True / gjson.False），bool 必须
	// 双向覆盖，与 sse_tolerance_consistency_test.go 的载体全集口径一致。
	scalars := []string{"42", `"hello"`, "null", "true", "false"}

	for _, raw := range scalars {
		if n, err := sseCountGlobalEvent(sseEventFrame([]byte(raw))); err != nil || n != 1 {
			t.Errorf("GlobalEvent whole=%s: n=%d err=%v, want n=1 err=nil (gjson path lookups degrade silently)", raw, n, err)
		}
	}

	for _, raw := range scalars {
		if n, err := sseCountEventListResponse(sseEventFrame([]byte(raw))); err != nil || n != 1 {
			t.Errorf("EventListResponse whole=%s: n=%d err=%v, want n=1 err=nil (IsObject guard degrades silently)", raw, n, err)
		}
	}

	for _, raw := range scalars {
		if n, err := sseCountV2Event(sseEventFrame([]byte(raw))); err != nil || n != 1 {
			t.Errorf("V2Event whole=%s: n=%d err=%v, want n=1 err=nil (IsObject guard degrades silently)", raw, n, err)
		}
	}

	// 流存活证明：裸标量帧之后紧跟一条合法事件，合法事件必须照常送达
	// （共 2 个事件且 Err() 为 nil）——降级帧绝不能像修复前那样杀死整条流。
	for _, chain := range []struct {
		name  string
		count func(frames ...string) (int, error)
		valid string
	}{
		{
			name:  "EventListResponse",
			count: sseCountEventListResponse,
			valid: sseEventFrame([]byte(`{"id":"evt_ok","type":"file.edited","properties":{"file":"/ok.go"}}`)),
		},
		{
			name:  "GlobalEvent",
			count: sseCountGlobalEvent,
			valid: sseEventFrame([]byte(`{"directory":"d","payload":{"id":"evt_ok","type":"file.edited","data":{"file":"/ok.go"}}}`)),
		},
		{
			name:  "V2Event",
			count: sseCountV2Event,
			valid: sseEventFrame([]byte(`{"id":"evt_ok","type":"file.edited","data":{"file":"/ok.go"}}`)),
		},
	} {
		for _, raw := range scalars {
			n, err := chain.count(sseEventFrame([]byte(raw)), chain.valid)
			if err != nil {
				t.Errorf("%s whole=%s then valid event: stream.Err() = %v, want nil (stream must survive degraded frames)", chain.name, raw, err)
			}
			if n != 2 {
				t.Errorf("%s whole=%s then valid event: got %d events, want 2 (degraded frame + subsequent valid event)", chain.name, raw, n)
			}
		}
	}

	// RawJSON 保留证明：三条链路的降级帧都必须把原始报文暴露在 RawJSON()
	// 上，供调用方区分畸形帧与正常路由的事件体。链路 1/3 由 IsObject 守卫
	// 显式写入 raw；链路 2 的 raw 由 Phase 1 的 apijson alias 元数据填充，
	// 对任意根类型生效（与 sse_tolerance_consistency_test.go 的断言口径一致）。
	for _, tc := range []struct {
		name    string
		consume func(raw []byte) (string, bool)
	}{
		{
			name: "EventListResponse",
			consume: func(raw []byte) (string, bool) {
				resp := newSSEResponse(sseEventFrame(raw))
				stream := ssestream.NewStream[opencode.EventListResponse](ssestream.NewDecoder(resp), nil)
				defer stream.Close()
				if !stream.Next() {
					return "", false
				}
				return stream.Current().JSON.RawJSON(), true
			},
		},
		{
			name: "GlobalEvent",
			consume: func(raw []byte) (string, bool) {
				resp := newSSEResponse(sseEventFrame(raw))
				stream := ssestream.NewStream[opencode.GlobalEvent](ssestream.NewDecoder(resp), nil)
				defer stream.Close()
				if !stream.Next() {
					return "", false
				}
				return stream.Current().JSON.RawJSON(), true
			},
		},
		{
			name: "V2Event",
			consume: func(raw []byte) (string, bool) {
				resp := newSSEResponse(sseEventFrame(raw))
				stream := ssestream.NewStream[opencode.V2Event](ssestream.NewDecoder(resp), nil)
				defer stream.Close()
				if !stream.Next() {
					return "", false
				}
				return stream.Current().JSON.RawJSON(), true
			},
		},
	} {
		for _, raw := range scalars {
			got, ok := tc.consume([]byte(raw))
			if !ok {
				t.Errorf("%s whole=%s: stream produced no event, want one degraded zero-value event", tc.name, raw)
				continue
			}
			if got != raw {
				t.Errorf("%s whole=%s: RawJSON() = %q, want original payload preserved", tc.name, raw, got)
			}
		}
	}
}

// =============================================================================
// SSE keep-alive 容错（已修复项的回归护栏）
//
// 背景：eventStreamDecoder.Next() 此前对"data 缓冲为空"的块仍无条件 dispatch
// 一个 Event{Data: []byte{}}（或仅含一个 "\n"），随后 json.Unmarshal 必报
// "unexpected end of JSON input"，经 ssestream.Stream.Next() 杀死整条 SSE 流，
// 三条链路无一例外。
//
// 这不是理论隐患：opencode v2 服务端（packages/server/src/handlers/event.ts）
// 每 15 秒向 /api/event 推送裸注释帧 ": heartbeat\n\n"，因此
// V2EventService.ListStreaming 在任何空闲 15 秒后必然断流。
//
// 修复已落在 packages/ssestream/ssestream.go 的 eventStreamDecoder.Next()：
// data 缓冲经 bytes.TrimSpace 后为空时跳过该块并重置累加器，与
//   - W3C SSE 规范 dispatch 算法（data 缓冲为空则 return，不派发）
//   - JS SDK(v2) 生成器（`if (data !== "") { yield JSON.parse(data) }`）
// 完全一致。以下测试锁定修复后的行为。
// =============================================================================

// TestReview2SSEKeepAliveDoesNotKillStream 验证携带空 data 缓冲的 SSE 帧
// （keep-alive / 心跳）被静默跳过，既不报错也不丢失其后的合法事件。
//
// 这是真实场景而非理论加固：opencode v2 服务端
// （packages/server/src/handlers/event.ts）每 15 秒向 /api/event 推送一个
// 裸注释帧 ": heartbeat\n\n"。注释行本身会被解码器当作 comment 跳过，但其后
// 的空行此前会派发一个 Data 为空的事件，交给 Stream.Next 的 json.Unmarshal
// 后必然失败并永久终止整条流——即 V2EventService.ListStreaming 在任何空闲
// 15 秒后必然断流。
//
// W3C SSE 规范的 dispatch 算法要求 data 缓冲为空时直接 return（不派发）；
// JS SDK(v2) 生成器同样以 `if (data !== "")` 跳过。三条链路行为必须一致。
func TestReview2SSEKeepAliveDoesNotKillStream(t *testing.T) {
	t.Parallel()

	keepAlives := []struct {
		name  string
		frame string
	}{
		// opencode v2 服务端的真实心跳形态（handlers/event.ts:37）
		{"bare comment heartbeat", ": heartbeat\n\n"},
		// 只有 event: 字段、完全没有 data: 行
		{"event without data line", "event: ping\n\n"},
		// data: 行存在但内容为空（解码器会留下一个裸 "\n"）
		{"empty data line", "data:\n\n"},
		// data: 行内容仅为空白
		{"whitespace-only data line", "data:   \n\n"},
	}

	for _, ka := range keepAlives {
		t.Run(ka.name, func(t *testing.T) {
			t.Run("EventListResponse.properties", func(t *testing.T) {
				n, err := sseCountEventListResponse(
					sseEventFrame([]byte(`{"id":"evt_1","type":"file.edited","properties":{"file":"/a.go"}}`)),
					ka.frame,
					sseEventFrame([]byte(`{"id":"evt_2","type":"file.edited","properties":{"file":"/b.go"}}`)),
				)
				if err != nil {
					t.Errorf("stream.Err() = %v, want nil (keep-alive frame must not kill the stream)", err)
				}
				if n != 2 {
					t.Errorf("got %d events, want 2 (no event may be dropped around a keep-alive frame)", n)
				}
			})

			t.Run("GlobalEvent.payload", func(t *testing.T) {
				n, err := sseCountGlobalEvent(
					sseEventFrame([]byte(`{"directory":"d","payload":{"id":"evt_1","type":"file.edited","data":{"file":"/a.go"}}}`)),
					ka.frame,
					sseEventFrame([]byte(`{"directory":"d","payload":{"id":"evt_2","type":"file.edited","data":{"file":"/b.go"}}}`)),
				)
				if err != nil {
					t.Errorf("stream.Err() = %v, want nil (keep-alive frame must not kill the stream)", err)
				}
				if n != 2 {
					t.Errorf("got %d events, want 2 (no event may be dropped around a keep-alive frame)", n)
				}
			})

			t.Run("V2Event.data", func(t *testing.T) {
				n, err := sseCountV2Event(
					sseEventFrame([]byte(`{"id":"evt_1","type":"file.edited","data":{"file":"/a.go"}}`)),
					ka.frame,
					sseEventFrame([]byte(`{"id":"evt_2","type":"file.edited","data":{"file":"/b.go"}}`)),
				)
				if err != nil {
					t.Errorf("stream.Err() = %v, want nil (keep-alive frame must not kill the stream)", err)
				}
				if n != 2 {
					t.Errorf("got %d events, want 2 (no event may be dropped around a keep-alive frame)", n)
				}
			})
		})
	}
}

// TestReview2SSEKeepAliveLeadingAndConsecutive 验证心跳出现在流最前、以及连续
// 多个心跳时，后续事件依然完整送达（覆盖 data/event 累加器在跳过时被正确重置）。
func TestReview2SSEKeepAliveLeadingAndConsecutive(t *testing.T) {
	t.Parallel()
	const hb = ": heartbeat\n\n"
	evt := sseEventFrame([]byte(`{"id":"evt_1","type":"file.edited","data":{"file":"/a.go"}}`))

	t.Run("leading keep-alive", func(t *testing.T) {
		n, err := sseCountV2Event(hb, evt)
		if err != nil {
			t.Errorf("stream.Err() = %v, want nil", err)
		}
		if n != 1 {
			t.Errorf("got %d events, want 1", n)
		}
	})

	t.Run("consecutive keep-alives", func(t *testing.T) {
		n, err := sseCountV2Event(hb, hb, hb, evt)
		if err != nil {
			t.Errorf("stream.Err() = %v, want nil", err)
		}
		if n != 1 {
			t.Errorf("got %d events, want 1", n)
		}
	})

	// 一条只有心跳、从未推送真实事件的流必须干净结束：0 个事件且 Err() 为 nil，
	// 调用方不应看到任何错误。
	t.Run("keep-alive only stream ends cleanly", func(t *testing.T) {
		n, err := sseCountV2Event(hb, hb)
		if err != nil {
			t.Errorf("stream.Err() = %v, want nil (a heartbeat-only stream must end cleanly)", err)
		}
		if n != 0 {
			t.Errorf("got %d events, want 0", n)
		}
	})
}

// TestReview2SSENonEmptyPayloadsStillDispatched 是跳过逻辑的反向护栏：确认
// 修复只跳过"空 data 缓冲"，不会误伤任何携带真实内容的帧。
func TestReview2SSENonEmptyPayloadsStillDispatched(t *testing.T) {
	t.Parallel()

	// "null" 是合法 JSON 文档（非空 data 缓冲），绝不能被跳过逻辑吞掉。
	// GlobalEvent 的 gjson 路径取值对非 object 根节点静默降级（见
	// TestReview2SSEBareScalarWholeEventTolerated），因此它能正向证明该帧确实
	// 被派发并计数。
	t.Run("data: null is dispatched on GlobalEvent chain", func(t *testing.T) {
		n, err := sseCountGlobalEvent(
			"data: null\n\n",
			sseEventFrame([]byte(`{"directory":"d","payload":{"id":"evt_1","type":"file.edited","data":{"file":"/a.go"}}}`)),
		)
		if err != nil {
			t.Errorf("stream.Err() = %v, want nil", err)
		}
		if n != 2 {
			t.Errorf("got %d events, want 2 (`null` is a valid JSON document and must not be skipped)", n)
		}
	})

	// V2Event 顶层是 Union carrier。历史上对裸标量返回 "was not able to coerce
	// type as union"；现已在 UnmarshalJSON 入口增加 IsObject 守卫，非对象载体
	// 静默降级（见 TestReview2SSEBareScalarWholeEventTolerated）。这里改用计数
	// 正向证明该帧确实被派发并消费——若被空缓冲跳过逻辑误吞，n 会是 0。
	t.Run("data: null is dispatched on V2Event chain", func(t *testing.T) {
		n, err := sseCountV2Event("data: null\n\n")
		if err != nil {
			t.Errorf("stream.Err() = %v, want nil", err)
		}
		if n != 1 {
			t.Errorf("got %d events, want 1 (`null` is a valid JSON document and must not be skipped)", n)
		}
	})

	// 真正的畸形 JSON 仍然必须报错——跳过逻辑不得把真实错误一并吞掉。
	t.Run("malformed JSON still surfaces an error", func(t *testing.T) {
		n, err := sseCountV2Event(
			sseEventFrame([]byte(`{"id":"evt_1","type":"file.edited","data":{"file":"/a.go"}}`)),
			"data: {bad\n\n",
			sseEventFrame([]byte(`{"id":"evt_2","type":"file.edited","data":{"file":"/b.go"}}`)),
		)
		if err == nil {
			t.Error("stream.Err() = nil, want a JSON error (genuinely malformed payloads must not be silently skipped)")
		}
		if n != 1 {
			t.Errorf("got %d events before the error, want 1", n)
		}
	})
}
