package opencode_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
)

// =============================================================================
// SSE 容错铁律
//
// ssestream.Stream.Next() 中 json.Unmarshal 一旦返回 error 即 return false，
// 整条 SSE 流永久终止。因此 session.status 事件的 status 载体对「无法匹配任何
// variant」的输入（null / 标量 / 数组 / 未知 type / 缺失 type）必须静默降级，
// 绝不可返回 error。三条链路的容错策略必须完全一致：
//   - EventListResponse.Properties
//   - GlobalEvent.Payload
//   - V2Event.Data
// =============================================================================

// sseEventFrame 构造单条 `event:`/`data:` SSE 帧。data 必须为单行，
// 续行会被 SSE 解析器当作新的 name/value 对。
func sseEventFrame(data []byte) string {
	return "event: message\ndata: " + string(data) + "\n\n"
}

// newSSEResponse 把若干条 SSE 帧拼成 text/event-stream 响应体。
func newSSEResponse(frames ...string) *http.Response {
	var b strings.Builder
	for _, f := range frames {
		b.WriteString(f)
	}
	return &http.Response{
		Body:   io.NopCloser(strings.NewReader(b.String())),
		Header: http.Header{"Content-Type": []string{"text/event-stream"}},
	}
}

// readGlobalEventTestdataCompact 读取真实 SSE 报文 testdata 并压缩成单行 JSON，
// 还原生产环境 data: 帧的紧凑形态。
func readGlobalEventTestdataCompact(t *testing.T, name string) []byte {
	t.Helper()

	raw, err := globalEventTestdata.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatalf("read testdata/%s: %v", name, err)
	}
	var buf bytes.Buffer
	if err := json.Compact(&buf, raw); err != nil {
		t.Fatalf("compact testdata/%s: %v", name, err)
	}
	return buf.Bytes()
}

// TestSessionStatusSSEFaultToleranceMatrix 对三条链路 × 各类畸形 status 做全矩阵
// 实测：任何一格返回 error 都会终结整条 SSE 流，即为 🔴 阻塞项。
func TestSessionStatusSSEFaultToleranceMatrix(t *testing.T) {
	chains := []struct {
		label  string
		wrap   func(status string) string
		decode func(raw []byte) (variant string, errText string)
	}{
		{
			label: "EventListResponse.Properties",
			wrap: func(s string) string {
				return `{"id":"evt_1","type":"session.status","properties":{"sessionID":"ses_1","status":` + s + `}}`
			},
			decode: func(raw []byte) (string, string) {
				var ev opencode.EventListResponse
				if err := json.Unmarshal(raw, &ev); err != nil {
					return "", err.Error()
				}
				return describeStatus(t, ev.AsUnion()), ""
			},
		},
		{
			label: "GlobalEvent.Payload",
			wrap: func(s string) string {
				return `{"directory":"d","payload":{"id":"evt_1","type":"session.status","properties":{"sessionID":"ses_1","status":` + s + `}}}`
			},
			decode: func(raw []byte) (string, string) {
				var ev opencode.GlobalEvent
				if err := json.Unmarshal(raw, &ev); err != nil {
					return "", err.Error()
				}
				return describeStatus(t, ev.AsUnion()), ""
			},
		},
		{
			label: "V2Event.Data",
			wrap: func(s string) string {
				return `{"id":"evt_1","type":"session.status","data":{"sessionID":"ses_1","status":` + s + `}}`
			},
			decode: func(raw []byte) (string, string) {
				var ev opencode.V2Event
				if err := json.Unmarshal(raw, &ev); err != nil {
					return "", err.Error()
				}
				return describeStatus(t, ev.AsUnion()), ""
			},
		},
	}

	statuses := []struct {
		name string
		raw  string
	}{
		{"null", `null`},
		{"string", `"not-an-object"`},
		{"number", `123`},
		{"bool", `true`},
		{"array", `[1,2,3]`},
		{"type-null", `{"type":null}`},
		{"unknown-type", `{"type":"compacting"}`},
		{"missing-type", `{"attempt":1,"message":"m","next":2}`},
		{"valid-idle", `{"type":"idle"}`},
		{"valid-retry", `{"type":"retry","attempt":1,"message":"m","next":2}`},
		{"valid-busy", `{"type":"busy"}`},
	}

	for _, ch := range chains {
		for _, st := range statuses {
			variant, errText := ch.decode([]byte(ch.wrap(st.raw)))
			t.Logf("[%s] status=%-13s -> %s", ch.label, st.name, variant)
			if errText != "" {
				t.Errorf("[%s] status=%s: decode error %q would terminate the SSE stream",
					ch.label, st.name, errText)
			}
		}
	}
}

// describeStatus 提取路由结果用于日志（不得让测试失败）。
func describeStatus(t *testing.T, union any) string {
	t.Helper()

	switch ev := union.(type) {
	case opencode.EventListResponseEventSessionStatus:
		return describeCarrier(ev.Properties.Status)
	case opencode.V2EventSessionStatus:
		return describeCarrier(ev.Data.Status)
	default:
		return "(" + typeName(union) + ")"
	}
}

func describeCarrier(s opencode.SessionStatus) string {
	rawInfo := "no-raw"
	if raw := s.JSON.RawJSON(); raw != "" {
		rawInfo = "raw=" + raw
	}
	return "carrier.Type=" + string(s.Type) + " union=" + typeName(s.AsUnion()) +
		" TypeMeta=" + fieldStatus(s.JSON.Type) + " " + rawInfo
}

func fieldStatus(f apijson.Field) string {
	switch {
	case f.IsMissing():
		return "missing"
	case f.IsNull():
		return "null"
	case f.IsInvalid():
		return "invalid"
	default:
		return "present"
	}
}

func typeName(v any) string {
	if v == nil {
		return "nil"
	}
	return reflect.TypeOf(v).Name()
}

// TestSessionStatusSSEParentMetaPreservesRaw 验证 status 路由降级时，父级结构体的
// JSON.Status 元数据仍完整保留原始值，调用方据此可分辨降级原因。
func TestSessionStatusSSEParentMetaPreservesRaw(t *testing.T) {
	for _, st := range []string{`null`, `"not-an-object"`, `123`, `[1,2,3]`} {
		raw := `{"id":"evt_1","type":"session.status","properties":{"sessionID":"ses_1","status":` + st + `}}`

		var ev opencode.EventListResponse
		if err := json.Unmarshal([]byte(raw), &ev); err != nil {
			t.Fatalf("status=%s: decode error %v", st, err)
		}
		ss, ok := ev.AsUnion().(opencode.EventListResponseEventSessionStatus)
		if !ok {
			t.Fatalf("status=%s: AsUnion() = %T", st, ev.AsUnion())
		}

		meta := ss.Properties.JSON.Status
		if meta.IsMissing() {
			t.Errorf("status=%s: Properties.JSON.Status is missing, want raw preserved", st)
			continue
		}
		if meta.Raw() != st {
			t.Errorf("status=%s: Properties.JSON.Status.Raw() = %q, want %q", st, meta.Raw(), st)
		}
	}
}

// TestSessionStatusSSEStreamSurvivesMalformedEvent 验证畸形 session.status 事件
// 不会终结整条流：后续的其他类型事件（file.edited）仍须正常送达并路由。
func TestSessionStatusSSEStreamSurvivesMalformedEvent(t *testing.T) {
	resp := newSSEResponse(
		sseEventFrame([]byte(`{"id":"evt_1","type":"session.status","properties":{"sessionID":"ses_1","status":"not-an-object"}}`)),
		sseEventFrame([]byte(`{"id":"evt_2","type":"file.edited","properties":{"file":"/tmp/a.txt"}}`)),
	)
	stream := ssestream.NewStream[opencode.EventListResponse](ssestream.NewDecoder(resp), nil)

	var events []opencode.EventListResponse
	for stream.Next() {
		events = append(events, stream.Current())
	}
	if err := stream.Err(); err != nil {
		t.Fatalf("stream terminated with error: %v", err)
	}
	if len(events) != 2 {
		t.Fatalf("got %d events, want 2 (malformed status event must not kill the stream)", len(events))
	}
	if ev := events[1].AsUnion(); ev != nil {
		if _, ok := ev.(opencode.EventListResponseEventFileEdited); !ok {
			t.Errorf("second event AsUnion() = %T, want EventListResponseEventFileEdited", ev)
		}
	}
}

// assertRealSessionStatusEvent 逐字段校验一条真实 session.status 事件。
// wantVariantType 为期望 variant 的单值枚举字面量（idle/busy）。
func assertRealSessionStatusEvent(t *testing.T, ev opencode.GlobalEvent, wantID, wantStatusRaw string, wantType opencode.SessionStatusType) {
	t.Helper()

	ss := asUnion[opencode.EventListResponseEventSessionStatus](t, ev)
	eq(t, "payload.id", ss.ID, wantID)
	eq(t, "payload.type", ss.Type, opencode.EventListResponseEventSessionStatusTypeSessionStatus)
	eq(t, "properties.sessionID", ss.Properties.SessionID, sessionLocal)
	eq(t, "properties.status.type", ss.Properties.Status.Type, wantType)
	if !ss.Properties.Status.Type.IsKnown() {
		t.Errorf("properties.status.type %q is not known", ss.Properties.Status.Type)
	}
	present(t, "properties.status.type", ss.Properties.Status.JSON.Type)
	rawJSONEq(t, "properties.status.raw", ss.Properties.Status.JSON.RawJSON(), wantStatusRaw)

	// idle/busy variant 不携带 retry 专属字段
	missing(t, "properties.status.attempt", ss.Properties.Status.JSON.Attempt)
	missing(t, "properties.status.message", ss.Properties.Status.JSON.Message)
	missing(t, "properties.status.action", ss.Properties.Status.JSON.Action)
	missing(t, "properties.status.next", ss.Properties.Status.JSON.Next)

	switch wantType {
	case opencode.SessionStatusTypeIdle:
		v := asVariant[opencode.SessionStatusIdle](t, "properties.status", ss.Properties.Status.AsUnion())
		eq(t, "properties.status.variant.type", v.Type, opencode.SessionStatusIdleTypeIdle)
		present(t, "properties.status.variant.type", v.JSON.Type)
	case opencode.SessionStatusTypeBusy:
		v := asVariant[opencode.SessionStatusBusy](t, "properties.status", ss.Properties.Status.AsUnion())
		eq(t, "properties.status.variant.type", v.Type, opencode.SessionStatusBusyTypeBusy)
		present(t, "properties.status.variant.type", v.JSON.Type)
	default:
		t.Fatalf("unsupported expected carrier type %q", wantType)
	}
}

// TestSessionStatusSSERealPayloadConsecutive 把 testdata 里两条真实生产报文
// （busy -> idle 状态转换）作为连续 SSE 事件送入流，逐条校验。
func TestSessionStatusSSERealPayloadConsecutive(t *testing.T) {
	busy := readGlobalEventTestdataCompact(t, "global_event_session_status_busy.json")
	idle := readGlobalEventTestdataCompact(t, "global_event_session_status_idle_local.json")

	resp := newSSEResponse(sseEventFrame(busy), sseEventFrame(idle))
	stream := ssestream.NewStream[opencode.GlobalEvent](ssestream.NewDecoder(resp), nil)

	var events []opencode.GlobalEvent
	for stream.Next() {
		events = append(events, stream.Current())
	}
	if err := stream.Err(); err != nil {
		t.Fatalf("stream terminated with error: %v", err)
	}
	if len(events) != 2 {
		t.Fatalf("got %d events, want 2", len(events))
	}

	eq(t, "directory", events[0].Directory, directoryLocal)
	eq(t, "project", events[0].Project, projectLocal)
	assertRealSessionStatusEvent(t, events[0],
		"evt_ff578e105001VFQHCtTTTpB5Dv", `{"type":"busy"}`, opencode.SessionStatusTypeBusy)

	eq(t, "directory", events[1].Directory, directoryLocal)
	eq(t, "project", events[1].Project, projectLocal)
	assertRealSessionStatusEvent(t, events[1],
		"evt_ff578e6930017xOOazETWIrHMc", `{"type":"idle"}`, opencode.SessionStatusTypeIdle)
}

// TestSessionStatusSSERealPayloadMalformedMiddle 在两条真实事件之间插入一条畸形
// status 事件，验证流不提前终止且后续真实事件仍能完整解码。
func TestSessionStatusSSERealPayloadMalformedMiddle(t *testing.T) {
	busy := readGlobalEventTestdataCompact(t, "global_event_session_status_busy.json")
	idle := readGlobalEventTestdataCompact(t, "global_event_session_status_idle_local.json")

	for _, mc := range []struct{ name, raw string }{
		{"null", `null`},
		{"string", `"not-an-object"`},
		{"number", `123`},
		{"array", `[1,2,3]`},
		{"unknown-type", `{"type":"compacting"}`},
		{"missing-type", `{"attempt":1,"message":"m","next":2}`},
	} {
		t.Run(mc.name, func(t *testing.T) {
			mid := []byte(`{"directory":"` + directoryLocal + `","payload":{"id":"evt_malformed_1","type":"session.status","properties":{"sessionID":"` + sessionLocal + `","status":` + mc.raw + `}}}`)

			resp := newSSEResponse(sseEventFrame(busy), sseEventFrame(mid), sseEventFrame(idle))
			stream := ssestream.NewStream[opencode.GlobalEvent](ssestream.NewDecoder(resp), nil)

			var events []opencode.GlobalEvent
			for stream.Next() {
				events = append(events, stream.Current())
			}
			if err := stream.Err(); err != nil {
				t.Fatalf("stream terminated with error %v — malformed status %s killed the stream", err, mc.raw)
			}
			if len(events) != 3 {
				t.Fatalf("got %d events, want 3 (malformed status %s must not kill the stream)", len(events), mc.raw)
			}

			assertRealSessionStatusEvent(t, events[0],
				"evt_ff578e105001VFQHCtTTTpB5Dv", `{"type":"busy"}`, opencode.SessionStatusTypeBusy)

			// 畸形事件本身仍路由到 session.status 载体（字段级降级，非整体失败）
			midSS := asUnion[opencode.EventListResponseEventSessionStatus](t, events[1])
			eq(t, "payload.id", midSS.ID, "evt_malformed_1")
			meta := midSS.Properties.JSON.Status
			if meta.IsMissing() {
				t.Errorf("status %s: Properties.JSON.Status is missing, want raw preserved", mc.raw)
			}
			if mc.raw == "null" {
				if !meta.IsNull() {
					t.Errorf("status null: JSON.Status.IsNull() = false, want true (raw=%s)", meta.Raw())
				}
			} else if meta.Raw() != mc.raw {
				t.Errorf("status %s: JSON.Status.Raw() = %q, want %q", mc.raw, meta.Raw(), mc.raw)
			}
			if mc.name == "unknown-type" {
				// 未知 type 原样保留且不得声称已知，保证服务端新增 variant 时前向兼容
				eq(t, "carrier.Type", midSS.Properties.Status.Type, opencode.SessionStatusType("compacting"))
				if midSS.Properties.Status.Type.IsKnown() {
					t.Errorf("unknown-type: carrier Type %q should not be known", midSS.Properties.Status.Type)
				}
			}

			// 🔴 关键断言：畸形事件之后的真实事件必须仍能送达并正确解码
			assertRealSessionStatusEvent(t, events[2],
				"evt_ff578e6930017xOOazETWIrHMc", `{"type":"idle"}`, opencode.SessionStatusTypeIdle)
		})
	}
}

// TestSessionStatusSSEV2EventMalformedMiddle 对 V2Event.Data 链路做同样验证，
// 确保三条 SSE 链路的容错策略一致。
func TestSessionStatusSSEV2EventMalformedMiddle(t *testing.T) {
	busy := []byte(`{"id":"evt_ff578e105001VFQHCtTTTpB5Dv","type":"session.status","data":{"sessionID":"` + sessionLocal + `","status":{"type":"busy"}}}`)
	idle := []byte(`{"id":"evt_ff578e6930017xOOazETWIrHMc","type":"session.status","data":{"sessionID":"` + sessionLocal + `","status":{"type":"idle"}}}`)

	for _, mc := range []struct{ name, raw string }{
		{"null", `null`},
		{"string", `"oops"`},
		{"unknown-type", `{"type":"compacting"}`},
	} {
		t.Run(mc.name, func(t *testing.T) {
			mid := []byte(`{"id":"evt_malformed_1","type":"session.status","data":{"sessionID":"` + sessionLocal + `","status":` + mc.raw + `}}`)

			resp := newSSEResponse(sseEventFrame(busy), sseEventFrame(mid), sseEventFrame(idle))
			stream := ssestream.NewStream[opencode.V2Event](ssestream.NewDecoder(resp), nil)

			var events []opencode.V2Event
			for stream.Next() {
				events = append(events, stream.Current())
			}
			if err := stream.Err(); err != nil {
				t.Fatalf("V2Event stream terminated with error %v — malformed status %s killed the stream", err, mc.raw)
			}
			if len(events) != 3 {
				t.Fatalf("V2Event got %d events, want 3", len(events))
			}

			first, ok := events[0].AsUnion().(opencode.V2EventSessionStatus)
			if !ok {
				t.Fatalf("first event AsUnion() = %T, want V2EventSessionStatus", events[0].AsUnion())
			}
			eq(t, "data.status.type", first.Data.Status.Type, opencode.SessionStatusTypeBusy)
			asVariant[opencode.SessionStatusBusy](t, "data.status", first.Data.Status.AsUnion())

			last, ok := events[2].AsUnion().(opencode.V2EventSessionStatus)
			if !ok {
				t.Fatalf("last event AsUnion() = %T, want V2EventSessionStatus", events[2].AsUnion())
			}
			eq(t, "data.status.type", last.Data.Status.Type, opencode.SessionStatusTypeIdle)
			asVariant[opencode.SessionStatusIdle](t, "data.status", last.Data.Status.AsUnion())
			rawJSONEq(t, "data.status.raw", last.Data.Status.JSON.RawJSON(), `{"type":"idle"}`)
		})
	}
}
