package opencode_test

import (
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
)

// =============================================================================
// SSE 容错铁律：四条事件链路容错策略一致性（横切验证）
//
// ssestream.Stream.Next() 中 json.Unmarshal 一旦出错即 return false，整条流
// 永久终止。因此四条事件载体链路对「不可能匹配任何 variant」的输入（OpenAPI
// union 全为 object schema 时收到 null/标量/数组）必须静默降级 return nil：
//
//	Chain 1  EventListResponse        (/event)                     event.go
//	Chain 2  GlobalEvent.Payload      (/global/event)              global.go
//	Chain 3  V2Event.Data             (/api/event)                 v2event.go
//	Chain 4  V2SessionDurableEvent    (/api/session/{id}/event)    v2session.go
//
// 链路 1/3/4 是整帧级守卫（UnmarshalJSON 入口处 gjson.IsObject 判定）；
// 链路 2 是 wrapper 结构的字段级守卫（payload 非对象时跳过 union 路由）。
// 机制不同，但对外契约必须一致（本文件逐链路实测锁定）：
//   - 全帧非对象输入：四链路均静默降级为零值事件派发（不吞帧、不报错），
//     且 JSON.RawJSON() 保留原始报文 —— 链路 1/3/4 由守卫显式写入 raw，
//     链路 2 由 apijson 结构体解码器在解析 alias 时写入（对任意根类型生效）；
//   - 畸形帧之后流仍存活，后续合法事件照常路由到具体 variant；
//   - 链路 2 的字段级降级语义（对象帧内非对象 payload 由 JSON.Payload 的
//     Raw()/IsNull()/IsMissing() 记录）由 global_event_metadata_test.go 单独覆盖。
// =============================================================================

// sseCollectEvents 把 frames 拼成 text/event-stream 响应体，通过真实
// ssestream.Stream[T] 消费，返回全部派发事件与 stream.Err()。
func sseCollectEvents[T any](frames ...string) ([]T, error) {
	resp := newSSEResponse(frames...)
	stream := ssestream.NewStream[T](ssestream.NewDecoder(resp), nil)
	var out []T
	for stream.Next() {
		out = append(out, stream.Current())
	}
	return out, stream.Err()
}

// sseMalformedCarriers 是「不可能匹配任何 variant」的畸形载体全集：
// null / number / string / bool(true+false) / array（OpenAPI 事件 union 仅含
// object schema；gjson 将 true/false 视为不同类型，故双向都必须覆盖）。
var sseMalformedCarriers = []string{"null", "123", `"hello"`, "true", "false", "[1,2]"}

// TestSSEToleranceConsistencyAllChains 同一组畸形载体喂给全部四条链路，
// 断言行为一致：均不报错、各派发一个零值事件、RawJSON() 保留原始报文。
func TestSSEToleranceConsistencyAllChains(t *testing.T) {
	t.Parallel()

	for _, raw := range sseMalformedCarriers {
		frame := sseEventFrame([]byte(raw))

		// Chain 1: EventListResponse (/event)
		evs, err := sseCollectEvents[opencode.EventListResponse](frame)
		if err != nil {
			t.Errorf("chain1 raw=%s: err=%v, want nil (must not kill the stream)", raw, err)
		}
		if len(evs) != 1 {
			t.Errorf("chain1 raw=%s: delivered %d events, want 1 (degraded zero-value event)", raw, len(evs))
			continue
		}
		if got := evs[0].JSON.RawJSON(); got != raw {
			t.Errorf("chain1 raw=%s: RawJSON()=%q, want original frame preserved", raw, got)
		}
		if evs[0].Type != "" || evs[0].Properties != nil {
			t.Errorf("chain1 raw=%s: expected zero-value event, got type=%q properties=%v", raw, evs[0].Type, evs[0].Properties)
		}

		// Chain 2: GlobalEvent (/global/event) —— 字段级守卫的全帧退化路径
		ges, err := sseCollectEvents[opencode.GlobalEvent](frame)
		if err != nil {
			t.Errorf("chain2 raw=%s: err=%v, want nil (must not kill the stream)", raw, err)
		}
		if len(ges) != 1 {
			t.Errorf("chain2 raw=%s: delivered %d events, want 1 (degraded zero-value event)", raw, len(ges))
			continue
		}
		if got := ges[0].JSON.RawJSON(); got != raw {
			t.Errorf("chain2 raw=%s: RawJSON()=%q, want original frame preserved", raw, got)
		}
		if ges[0].Directory != "" || ges[0].Project != "" || ges[0].Workspace != "" || ges[0].Payload != nil || ges[0].AsUnion() != nil {
			t.Errorf("chain2 raw=%s: expected zero-value event, got dir=%q project=%q workspace=%q payload=%v union=%T",
				raw, ges[0].Directory, ges[0].Project, ges[0].Workspace, ges[0].Payload, ges[0].AsUnion())
		}

		// Chain 3: V2Event (/api/event)
		v2s, err := sseCollectEvents[opencode.V2Event](frame)
		if err != nil {
			t.Errorf("chain3 raw=%s: err=%v, want nil (must not kill the stream)", raw, err)
		}
		if len(v2s) != 1 {
			t.Errorf("chain3 raw=%s: delivered %d events, want 1 (degraded zero-value event)", raw, len(v2s))
			continue
		}
		if got := v2s[0].JSON.RawJSON(); got != raw {
			t.Errorf("chain3 raw=%s: RawJSON()=%q, want original frame preserved", raw, got)
		}
		if v2s[0].ID != "" || v2s[0].Data != nil || v2s[0].AsUnion() != nil {
			t.Errorf("chain3 raw=%s: expected zero-value event, got id=%q data=%v union=%T",
				raw, v2s[0].ID, v2s[0].Data, v2s[0].AsUnion())
		}

		// Chain 4: V2SessionDurableEvent (/api/session/{id}/event)
		des, err := sseCollectEvents[opencode.V2SessionDurableEvent](frame)
		if err != nil {
			t.Errorf("chain4 raw=%s: err=%v, want nil (must not kill the stream)", raw, err)
		}
		if len(des) != 1 {
			t.Errorf("chain4 raw=%s: delivered %d events, want 1 (degraded zero-value event)", raw, len(des))
			continue
		}
		if got := des[0].JSON.RawJSON(); got != raw {
			t.Errorf("chain4 raw=%s: RawJSON()=%q, want original frame preserved", raw, got)
		}
		if des[0].AsUnion() != nil {
			t.Errorf("chain4 raw=%s: expected zero-value event, got union=%T", raw, des[0].AsUnion())
		}
	}
}

// TestSSEToleranceStreamSurvivesMalformedFrame 畸形帧之后跟随合法事件，
// 断言四条链路的流均存活：畸形帧被静默降级但仍作为零值事件派发（n=2 =
// 降级零值事件 + 合法事件），且合法事件路由到正确的 variant 与字段值。
func TestSSEToleranceStreamSurvivesMalformedFrame(t *testing.T) {
	t.Parallel()

	t.Run("chain1 EventListResponse", func(t *testing.T) {
		t.Parallel()
		valid := `{"id":"evt_1","type":"session.idle","properties":{"sessionID":"ses_1"}}`
		evs, err := sseCollectEvents[opencode.EventListResponse](
			sseEventFrame([]byte(`null`)),
			sseEventFrame([]byte(valid)),
		)
		if err != nil {
			t.Fatalf("err=%v, want nil (stream must survive a malformed frame)", err)
		}
		if len(evs) != 2 {
			t.Fatalf("delivered %d events, want 2 (degraded zero-value event + the valid event)", len(evs))
		}
		if evs[0].JSON.RawJSON() != "null" || evs[0].Type != "" {
			t.Errorf("degraded event: RawJSON()=%q type=%q, want \"null\" and zero value", evs[0].JSON.RawJSON(), evs[0].Type)
		}
		idle, ok := evs[1].AsUnion().(opencode.EventListResponseEventSessionIdle)
		if !ok {
			t.Fatalf("valid event routed to %T, want EventListResponseEventSessionIdle", evs[1].AsUnion())
		}
		if idle.ID != "evt_1" || idle.Properties.SessionID != "ses_1" ||
			idle.Type != opencode.EventListResponseEventSessionIdleTypeSessionIdle {
			t.Errorf("routed variant mismatch: id=%q sessionID=%q type=%q", idle.ID, idle.Properties.SessionID, idle.Type)
		}
	})

	t.Run("chain2 GlobalEvent", func(t *testing.T) {
		t.Parallel()
		valid := `{"directory":"/","project":"p","workspace":"w","payload":{"id":"evt_1","type":"models-dev.refreshed","properties":{}}}`
		ges, err := sseCollectEvents[opencode.GlobalEvent](
			sseEventFrame([]byte(`null`)),
			sseEventFrame([]byte(valid)),
		)
		if err != nil {
			t.Fatalf("err=%v, want nil (stream must survive a malformed frame)", err)
		}
		if len(ges) != 2 {
			t.Fatalf("delivered %d events, want 2 (degraded zero-value event + the valid event)", len(ges))
		}
		if ges[0].JSON.RawJSON() != "null" || ges[0].Payload != nil {
			t.Errorf("degraded event: RawJSON()=%q payload=%v, want \"null\" and nil", ges[0].JSON.RawJSON(), ges[0].Payload)
		}
		if ges[1].Directory != "/" || ges[1].Project != "p" || ges[1].Workspace != "w" {
			t.Errorf("wrapper fields mismatch: dir=%q project=%q workspace=%q", ges[1].Directory, ges[1].Project, ges[1].Workspace)
		}
		mdr, ok := ges[1].AsUnion().(opencode.EventListResponseEventModelsDevRefreshed)
		if !ok {
			t.Fatalf("valid payload routed to %T, want EventListResponseEventModelsDevRefreshed", ges[1].AsUnion())
		}
		if mdr.ID != "evt_1" || mdr.Type != opencode.EventListResponseEventModelsDevRefreshedTypeModelsDevRefreshed {
			t.Errorf("routed variant mismatch: id=%q type=%q", mdr.ID, mdr.Type)
		}
	})

	t.Run("chain3 V2Event", func(t *testing.T) {
		t.Parallel()
		valid := `{"id":"evt_1","type":"models-dev.refreshed","data":{}}`
		v2s, err := sseCollectEvents[opencode.V2Event](
			sseEventFrame([]byte(`null`)),
			sseEventFrame([]byte(valid)),
		)
		if err != nil {
			t.Fatalf("err=%v, want nil (stream must survive a malformed frame)", err)
		}
		if len(v2s) != 2 {
			t.Fatalf("delivered %d events, want 2 (degraded zero-value event + the valid event)", len(v2s))
		}
		if v2s[0].JSON.RawJSON() != "null" || v2s[0].ID != "" {
			t.Errorf("degraded event: RawJSON()=%q id=%q, want \"null\" and empty id", v2s[0].JSON.RawJSON(), v2s[0].ID)
		}
		mdr, ok := v2s[1].AsUnion().(opencode.V2EventModelsDevRefreshed)
		if !ok {
			t.Fatalf("valid event routed to %T, want V2EventModelsDevRefreshed", v2s[1].AsUnion())
		}
		if mdr.ID != "evt_1" || mdr.Type != opencode.V2EventModelsDevRefreshedTypeModelsDevRefreshed {
			t.Errorf("routed variant mismatch: id=%q type=%q", mdr.ID, mdr.Type)
		}
	})

	t.Run("chain4 V2SessionDurableEvent", func(t *testing.T) {
		t.Parallel()
		valid := `{"id":"evt_1","type":"session.next.agent.switched","durable":{"aggregateID":"x","seq":1,"version":1},"location":{"directory":"/w"},"data":{"agent":"build","messageID":"msg_1","sessionID":"ses_1","timestamp":123}}`
		des, err := sseCollectEvents[opencode.V2SessionDurableEvent](
			sseEventFrame([]byte(`null`)),
			sseEventFrame([]byte(valid)),
		)
		if err != nil {
			t.Fatalf("err=%v, want nil (stream must survive a malformed frame)", err)
		}
		if len(des) != 2 {
			t.Fatalf("delivered %d events, want 2 (degraded zero-value event + the valid event)", len(des))
		}
		if des[0].JSON.RawJSON() != "null" || des[0].AsUnion() != nil {
			t.Errorf("degraded event: RawJSON()=%q union=%T, want \"null\" and nil", des[0].JSON.RawJSON(), des[0].AsUnion())
		}
		ag, ok := des[1].AsUnion().(opencode.V2SessionDurableEventAgentSwitched)
		if !ok {
			t.Fatalf("valid event routed to %T, want V2SessionDurableEventAgentSwitched", des[1].AsUnion())
		}
		if ag.ID != "evt_1" || ag.Data.Agent != "build" || ag.Data.SessionID != "ses_1" ||
			ag.Data.Timestamp != 123 || ag.Durable.Seq != 1 || ag.Location.Directory != "/w" {
			t.Errorf("routed variant mismatch: id=%q agent=%q sessionID=%q ts=%d seq=%d dir=%q",
				ag.ID, ag.Data.Agent, ag.Data.SessionID, ag.Data.Timestamp, ag.Durable.Seq, ag.Location.Directory)
		}
	})
}

// TestSSEToleranceV2EventUnknownFieldStillRoutes 锁定链路 3 的路由鲁棒性：
// 帧内含 V2Event variant 不认识的字段（如历史报文形态 "properties"）时，
// union smart-mode 下目标 variant 因 type 枚举匹配仅降到 extras（未知字段），
// 其余 variant 因 type 枚举不匹配降到 loose，故仍确定性地路由到
// V2EventModelsDevRefreshed —— 不报错、不误路由到注册序靠前的 variant。
func TestSSEToleranceV2EventUnknownFieldStillRoutes(t *testing.T) {
	t.Parallel()
	valid := `{"id":"evt_1","type":"models-dev.refreshed","properties":{}}`
	v2s, err := sseCollectEvents[opencode.V2Event](sseEventFrame([]byte(valid)))
	if err != nil {
		t.Fatalf("err=%v, want nil", err)
	}
	if len(v2s) != 1 {
		t.Fatalf("delivered %d events, want 1", len(v2s))
	}
	mdr, ok := v2s[0].AsUnion().(opencode.V2EventModelsDevRefreshed)
	if !ok {
		t.Fatalf("event with unknown field routed to %T, want V2EventModelsDevRefreshed", v2s[0].AsUnion())
	}
	if mdr.ID != "evt_1" {
		t.Errorf("routed variant id=%q, want evt_1", mdr.ID)
	}
}
