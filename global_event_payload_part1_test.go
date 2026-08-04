package opencode_test

import (
	"embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sst/opencode-sdk-go"
)

// globalEventTestdata 嵌入生产环境真实 SSE payload，作为 Global Event 反序列化的黄金样本。
// 每个文件对应一条服务端实际下发的 /global/event 报文，字段不可增删改。
//
// 任何改动后，必须运行本测试文件，使用真实 SSE payload 测试通过才算真正通过测试
//
//go:embed testdata/global_event_*.json
var globalEventTestdata embed.FS

// decodeGlobalEvent 从 testdata 读取真实 SSE payload 并反序列化为 [opencode.GlobalEvent]。
func decodeGlobalEvent(t *testing.T, name string) opencode.GlobalEvent {
	t.Helper()

	data, err := globalEventTestdata.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatalf("read testdata/%s: %v", name, err)
	}

	var ev opencode.GlobalEvent
	if err := json.Unmarshal(data, &ev); err != nil {
		t.Fatalf("unmarshal testdata/%s: %v", name, err)
	}
	return ev
}

// asUnion 断言 [opencode.GlobalEvent] 的 payload 落到期望的 variant 类型。
func asUnion[T any](t *testing.T, ev opencode.GlobalEvent) T {
	t.Helper()

	u := ev.AsUnion()
	got, ok := u.(T)
	if !ok {
		var want T
		t.Fatalf("payload resolved to %T, want %T", u, want)
	}
	return got
}

// asVariant 断言任意联合体的 AsUnion() 结果落到期望的 variant 类型。
func asVariant[T any](t *testing.T, field string, u any) T {
	t.Helper()

	got, ok := u.(T)
	if !ok {
		var want T
		t.Fatalf("%s resolved to %T, want %T", field, u, want)
	}
	return got
}

// eq 逐字段比对实际值与真实 payload 中的期望值。
func eq[T comparable](t *testing.T, field string, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("%s = %#v, want %#v", field, got, want)
	}
}

// fieldMeta 抽象 apijson 生成的 JSON 元数据字段，用于区分「显式零值」与「字段缺失」。
type fieldMeta interface {
	IsMissing() bool
	Raw() string
}

// present 断言字段确实由服务端下发（即便其值为零值）。
func present(t *testing.T, field string, meta fieldMeta) {
	t.Helper()

	if meta.IsMissing() {
		t.Errorf("%s should be present in payload, but is reported missing", field)
	}
}

// missing 断言字段未由服务端下发（Go 零值不可与显式零值混淆）。
func missing(t *testing.T, field string, meta fieldMeta) {
	t.Helper()

	if !meta.IsMissing() {
		t.Errorf("%s should be missing, but is present with raw=%s", field, meta.Raw())
	}
}

const (
	sessionMain     = "ses_03447e959ffec5bK9xrcGkt38N"
	directoryMain   = "/workspace/019dd6d3-9e26-7cb9-99d1-1b9fe82a57b3/019fcbb8-07d6-77ab-9015-3106c9364285"
	projectMain     = "b8d349ffbdefd0afab2f1f1ca03829cd7662268f"
	sessionSync     = "ses_034273f8affeAH2tdzMh9WZMh2"
	directorySync   = "/workspace/019dd6d3-9e26-7cb9-99d1-1b9fe82a57b3/019fcbd8-b6f9-73f2-8211-96a2184ccc67"
	projectSync     = "68e92b396e0e4333aeed4a3df91e524b54d6a173"
	messageAssistID = "msg_fcbc06b27001uNOhch8dPpRW0E"
)

// TestGlobalEventRealPayloadReasoningPart 覆盖 testdata/global_event_reasoning_part.json，
// 逐字段校验 message.part.updated + reasoning part（思考完成：text 有内容，time 含 start/end）。
func TestGlobalEventRealPayloadReasoningPart(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_reasoning_part.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbbb6252001hYg5ePEGcOywWX")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	if !pu.Type.IsKnown() {
		t.Errorf("payload.type %q is not known", pu.Type)
	}
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829483090))

	// Part 承载结构
	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbbb5b90001qV5QXBym6KMFgi")
	eq(t, "properties.part.messageID", carrier.MessageID, "msg_fcbbb4d18001mEwkbP5dEz4rAG")
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeReasoning)
	if !carrier.Type.IsKnown() {
		t.Errorf("properties.part.type %q is not known", carrier.Type)
	}
	eq(t, "properties.part.text", carrier.Text, "I'll generate the visual prototype with the exact requirements summary from Step 5.")

	// ReasoningPart variant（不可被 TextPart 抢占）
	rp := asVariant[opencode.ReasoningPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", rp.ID, "prt_fcbbb5b90001qV5QXBym6KMFgi")
	eq(t, "part.messageID", rp.MessageID, "msg_fcbbb4d18001mEwkbP5dEz4rAG")
	eq(t, "part.sessionID", rp.SessionID, sessionMain)
	eq(t, "part.type", rp.Type, opencode.ReasoningPartTypeReasoning)
	if !rp.Type.IsKnown() {
		t.Errorf("part.type %q is not known", rp.Type)
	}
	eq(t, "part.text", rp.Text, "I'll generate the visual prototype with the exact requirements summary from Step 5.")
	eq(t, "part.time.start", rp.Time.Start, int64(1785829481360))
	eq(t, "part.time.end", rp.Time.End, int64(1785829483090))

	// payload 未下发的可选字段
	eq(t, "part.metadata", len(rp.Metadata), 0)
	missing(t, "part.metadata", rp.JSON.Metadata)
	present(t, "part.text", rp.JSON.Text)
	present(t, "part.time", rp.JSON.Time)
	present(t, "part.time.start", rp.Time.JSON.Start)
	present(t, "part.time.end", rp.Time.JSON.End)
}

// TestGlobalEventRealPayloadStepFinishPart 覆盖 testdata/global_event_step_finish_part.json，
// 逐字段校验 message.part.updated + step-finish part（reason=tool-calls）。
func TestGlobalEventRealPayloadStepFinishPart(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_step_finish_part.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbbb4ceb001GVwrDJLwWljum8")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829477611))

	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbbb4ceb001mD3gM1rbqb5MV7")
	eq(t, "properties.part.messageID", carrier.MessageID, "msg_fcbbb3f480017UYG1SH55nPzS8")
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeStepFinish)
	eq(t, "properties.part.reason", carrier.Reason, "tool-calls")
	eq(t, "properties.part.cost", carrier.Cost, float64(0))

	sf := asVariant[opencode.StepFinishPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", sf.ID, "prt_fcbbb4ceb001mD3gM1rbqb5MV7")
	eq(t, "part.messageID", sf.MessageID, "msg_fcbbb3f480017UYG1SH55nPzS8")
	eq(t, "part.sessionID", sf.SessionID, sessionMain)
	eq(t, "part.type", sf.Type, opencode.StepFinishPartTypeStepFinish)
	if !sf.Type.IsKnown() {
		t.Errorf("part.type %q is not known", sf.Type)
	}
	eq(t, "part.reason", sf.Reason, "tool-calls")
	eq(t, "part.cost", sf.Cost, float64(0))
	eq(t, "part.tokens.total", sf.Tokens.Total, int64(50229))
	eq(t, "part.tokens.input", sf.Tokens.Input, int64(356))
	eq(t, "part.tokens.output", sf.Tokens.Output, int64(45))
	eq(t, "part.tokens.reasoning", sf.Tokens.Reasoning, int64(36))
	eq(t, "part.tokens.cache.write", sf.Tokens.Cache.Write, int64(0))
	eq(t, "part.tokens.cache.read", sf.Tokens.Cache.Read, int64(49792))

	// cost 与 cache.write 均为显式 0，必须判定为「已下发」而非「缺失」
	present(t, "part.cost", sf.JSON.Cost)
	present(t, "part.tokens.cache.write", sf.Tokens.Cache.JSON.Write)

	// payload 未下发的可选字段
	eq(t, "part.snapshot", sf.Snapshot, "")
	missing(t, "part.snapshot", sf.JSON.Snapshot)
}

// TestGlobalEventRealPayloadSyncSessionCreated 覆盖 testdata/global_event_sync_session_created.json，
// 逐字段校验 type=sync 的 V1 SyncEvent（session.created.1），含 4 条 permission 规则。
func TestGlobalEventRealPayloadSyncSessionCreated(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_sync_session_created.json")

	eq(t, "directory", ev.Directory, directorySync)
	eq(t, "project", ev.Project, projectSync)

	sync := asUnion[opencode.SyncEventResponse](t, ev)
	eq(t, "payload.type", sync.Type, opencode.SyncEventResponseTypeSync)
	if !sync.Type.IsKnown() {
		t.Errorf("payload.type %q is not known", sync.Type)
	}
	eq(t, "payload.id", sync.ID, "evt_fcbd8c076001YwctXZSulR3dYO")

	eq(t, "payload.syncEvent.id", sync.SyncEvent.ID, "evt_fcbd8c076001YwctXZSulR3dYO")
	eq(t, "payload.syncEvent.type", sync.SyncEvent.Type, opencode.SyncEventResponseSyncEventType("session.created.1"))
	if !sync.SyncEvent.Type.IsKnown() {
		t.Errorf("payload.syncEvent.type %q is not known", sync.SyncEvent.Type)
	}
	eq(t, "payload.syncEvent.seq", sync.SyncEvent.Seq, int64(0))
	eq(t, "payload.syncEvent.aggregateID", sync.SyncEvent.AggregateID, sessionSync)
	// seq 为显式 0，必须判定为「已下发」
	present(t, "payload.syncEvent.seq", sync.SyncEvent.JSON.Seq)

	sc := asVariant[opencode.SyncEventSessionCreated](t, "payload.syncEvent", sync.SyncEvent.AsUnion())
	eq(t, "syncEvent.type", sc.Type, opencode.SyncEventSessionCreatedTypeSessionCreated1)
	eq(t, "syncEvent.id", sc.ID, "evt_fcbd8c076001YwctXZSulR3dYO")
	eq(t, "syncEvent.seq", sc.Seq, int64(0))
	eq(t, "syncEvent.aggregateID", sc.AggregateID, sessionSync)
	eq(t, "syncEvent.data.sessionID", sc.Data.SessionID, sessionSync)

	info := sc.Data.Info
	eq(t, "info.id", info.ID, sessionSync)
	eq(t, "info.slug", info.Slug, "proud-tiger")
	eq(t, "info.version", info.Version, "1.18.9")
	eq(t, "info.projectID", info.ProjectID, projectSync)
	eq(t, "info.directory", info.Directory, directorySync)
	eq(t, "info.path", info.Path, "")
	eq(t, "info.title", info.Title, "Coding Upload App(dev-upload-app-16)")
	eq(t, "info.cost", info.Cost, float64(0))
	// path 为显式空串、cost 为显式 0，必须判定为「已下发」
	present(t, "info.path", info.JSON.Path)
	present(t, "info.cost", info.JSON.Cost)

	eq(t, "info.tokens.input", info.Tokens.Input, int64(0))
	eq(t, "info.tokens.output", info.Tokens.Output, int64(0))
	eq(t, "info.tokens.reasoning", info.Tokens.Reasoning, int64(0))
	eq(t, "info.tokens.cache.read", info.Tokens.Cache.Read, int64(0))
	eq(t, "info.tokens.cache.write", info.Tokens.Cache.Write, int64(0))
	present(t, "info.tokens.input", info.Tokens.JSON.Input)
	present(t, "info.tokens.cache.read", info.Tokens.Cache.JSON.Read)

	eq(t, "info.time.created", info.Time.Created, int64(1785831407733))
	eq(t, "info.time.updated", info.Time.Updated, int64(1785831407733))
	eq(t, "info.time.compacting", info.Time.Compacting, int64(0))
	eq(t, "info.time.archived", info.Time.Archived, int64(0))
	missing(t, "info.time.compacting", info.Time.JSON.Compacting)
	missing(t, "info.time.archived", info.Time.JSON.Archived)

	// 🔴 permission 在 OpenAPI 中是 PermissionRuleset = Array<PermissionRule>（非联合体），
	// 必须反序列化为强类型切片，逐条校验全部 4 条规则。
	eq(t, "len(info.permission)", len(info.Permission), 4)
	wantRules := []struct {
		permission string
		pattern    string
		action     opencode.PermissionAction
	}{
		{"read", directorySync + "/**", opencode.PermissionActionAllow},
		{"write", directorySync + "/**", opencode.PermissionActionAllow},
		{"execute", directorySync + "/**", opencode.PermissionActionAllow},
		{"all", "**", opencode.PermissionActionDeny},
	}
	for i, want := range wantRules {
		if i >= len(info.Permission) {
			break
		}
		got := info.Permission[i]
		field := fmt.Sprintf("info.permission[%d]", i)
		eq(t, field+".permission", got.Permission, want.permission)
		eq(t, field+".pattern", got.Pattern, want.pattern)
		eq(t, field+".action", got.Action, want.action)
		if !got.Action.IsKnown() {
			t.Errorf("%s.action %q is not known", field, got.Action)
		}
	}

	// payload 未下发的可选字段
	eq(t, "info.agent", info.Agent, "")
	eq(t, "info.parentID", info.ParentID, "")
	eq(t, "info.workspaceID", info.WorkspaceID, "")
	missing(t, "info.agent", info.JSON.Agent)
	missing(t, "info.parentID", info.JSON.ParentID)
	missing(t, "info.workspaceID", info.JSON.WorkspaceID)
	missing(t, "info.metadata", info.JSON.Metadata)
	missing(t, "info.share", info.JSON.Share)
	missing(t, "info.revert", info.JSON.Revert)
	missing(t, "info.summary", info.JSON.Summary)
	missing(t, "info.model", info.JSON.Model)
}

// TestGlobalEventRealPayloadMessagePartDelta 覆盖 testdata/global_event_message_part_delta.json，
// 逐字段校验 message.part.delta（文本流式增量，含 UTF-8 多字节字符）。
func TestGlobalEventRealPayloadMessagePartDelta(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_message_part_delta.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	d := asUnion[opencode.EventListResponseEventMessagePartDelta](t, ev)
	eq(t, "payload.id", d.ID, "evt_fcbbb4bc4003DF4SfJVuoUQqJ1")
	eq(t, "payload.type", d.Type, opencode.EventListResponseEventMessagePartDeltaTypeMessagePartDelta)
	if !d.Type.IsKnown() {
		t.Errorf("payload.type %q is not known", d.Type)
	}

	eq(t, "properties.sessionID", d.Properties.SessionID, sessionMain)
	eq(t, "properties.messageID", d.Properties.MessageID, "msg_fcbbb3f480017UYG1SH55nPzS8")
	eq(t, "properties.partID", d.Properties.PartID, "prt_fcbbb48de001AJse4fWcQNSdb4")
	eq(t, "properties.field", d.Properties.Field, "text")
	eq(t, "properties.delta", d.Properties.Delta, "供")

	present(t, "properties.sessionID", d.Properties.JSON.SessionID)
	present(t, "properties.messageID", d.Properties.JSON.MessageID)
	present(t, "properties.partID", d.Properties.JSON.PartID)
	present(t, "properties.field", d.Properties.JSON.Field)
	present(t, "properties.delta", d.Properties.JSON.Delta)
}

// TestGlobalEventRealPayloadSessionStatusIdle 覆盖 testdata/global_event_session_status_idle.json，
// 逐字段校验 session.status（status 为 SessionStatus 联合体的 idle variant）。
func TestGlobalEventRealPayloadSessionStatusIdle(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_session_status_idle.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	ss := asUnion[opencode.EventListResponseEventSessionStatus](t, ev)
	eq(t, "payload.id", ss.ID, "evt_fcbc079d0001peox3WNMvFjNGI")
	eq(t, "payload.type", ss.Type, opencode.EventListResponseEventSessionStatusTypeSessionStatus)
	if !ss.Type.IsKnown() {
		t.Errorf("payload.type %q is not known", ss.Type)
	}
	eq(t, "properties.sessionID", ss.Properties.SessionID, sessionMain)
	present(t, "properties.sessionID", ss.Properties.JSON.SessionID)
	present(t, "properties.status", ss.Properties.JSON.Status)

	idle := asVariant[opencode.SessionStatusIdle](t, "properties.status", ss.Properties.AsStatus())
	eq(t, "properties.status.type", idle.Type, "idle")
	present(t, "properties.status.type", idle.JSON.Type)

	// 承载字段 Status 必须持有同一个 variant 实例
	carried := asVariant[opencode.SessionStatusIdle](t, "properties.status(any)", ss.Properties.Status)
	eq(t, "properties.status(any).type", carried.Type, "idle")
}

// TestGlobalEventRealPayloadMessageUpdatedAssistant 覆盖
// testdata/global_event_message_updated_assistant.json，逐字段校验 message.updated
// + AssistantMessage（已完成：time 含 created 与 completed）。
func TestGlobalEventRealPayloadMessageUpdatedAssistant(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_message_updated_assistant.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	mu := asUnion[opencode.EventListResponseEventMessageUpdated](t, ev)
	eq(t, "payload.id", mu.ID, "evt_fcbc079af0019rOzyCx24nb0cU")
	eq(t, "payload.type", mu.Type, opencode.EventListResponseEventMessageUpdatedTypeMessageUpdated)
	if !mu.Type.IsKnown() {
		t.Errorf("payload.type %q is not known", mu.Type)
	}
	eq(t, "properties.sessionID", mu.Properties.SessionID, sessionMain)

	// Message 承载结构
	carrier := mu.Properties.Info
	eq(t, "properties.info.id", carrier.ID, messageAssistID)
	eq(t, "properties.info.role", carrier.Role, opencode.MessageRoleAssistant)
	if !carrier.Role.IsKnown() {
		t.Errorf("properties.info.role %q is not known", carrier.Role)
	}
	eq(t, "properties.info.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.info.agent", carrier.Agent, "ai-coding-assistant")
	eq(t, "properties.info.mode", carrier.Mode, "ai-coding-assistant")
	eq(t, "properties.info.modelID", carrier.ModelID, "grok/grok-4.5")
	eq(t, "properties.info.providerID", carrier.ProviderID, "cf-anthropic")
	eq(t, "properties.info.parentID", carrier.ParentID, "msg_fcbb816b40017LmBKE1UFzBLrx")
	eq(t, "properties.info.cost", carrier.Cost, float64(0))
	eq(t, "properties.info.finish", carrier.Finish, "stop")

	// AssistantMessage variant
	am := asVariant[opencode.AssistantMessage](t, "properties.info", carrier.AsUnion())
	eq(t, "info.id", am.ID, messageAssistID)
	eq(t, "info.parentID", am.ParentID, "msg_fcbb816b40017LmBKE1UFzBLrx")
	eq(t, "info.role", am.Role, opencode.AssistantMessageRoleAssistant)
	if !am.Role.IsKnown() {
		t.Errorf("info.role %q is not known", am.Role)
	}
	eq(t, "info.sessionID", am.SessionID, sessionMain)
	eq(t, "info.mode", am.Mode, "ai-coding-assistant")
	eq(t, "info.agent", am.Agent, "ai-coding-assistant")
	eq(t, "info.modelID", am.ModelID, "grok/grok-4.5")
	eq(t, "info.providerID", am.ProviderID, "cf-anthropic")
	eq(t, "info.cost", am.Cost, float64(0))
	eq(t, "info.finish", am.Finish, "stop")
	eq(t, "info.path.cwd", am.Path.Cwd, directoryMain)
	eq(t, "info.path.root", am.Path.Root, directoryMain)
	eq(t, "info.time.created", am.Time.Created, int64(1785829813031))
	eq(t, "info.time.completed", am.Time.Completed, int64(1785829816751))
	eq(t, "info.tokens.total", am.Tokens.Total, int64(72351))
	eq(t, "info.tokens.input", am.Tokens.Input, int64(498))
	eq(t, "info.tokens.output", am.Tokens.Output, int64(45))
	eq(t, "info.tokens.reasoning", am.Tokens.Reasoning, int64(0))
	eq(t, "info.tokens.cache.write", am.Tokens.Cache.Write, int64(0))
	eq(t, "info.tokens.cache.read", am.Tokens.Cache.Read, int64(71808))

	// cost / reasoning / cache.write 均为显式 0，必须判定为「已下发」
	present(t, "info.cost", am.JSON.Cost)
	present(t, "info.tokens.reasoning", am.Tokens.JSON.Reasoning)
	present(t, "info.tokens.cache.write", am.Tokens.Cache.JSON.Write)
	present(t, "info.time.completed", am.Time.JSON.Completed)

	// payload 未下发的可选字段
	eq(t, "info.variant", am.Variant, "")
	eq(t, "info.summary", am.Summary, false)
	if am.Structured != nil {
		t.Errorf("info.structured = %#v, want nil", am.Structured)
	}
	missing(t, "info.variant", am.JSON.Variant)
	missing(t, "info.summary", am.JSON.Summary)
	missing(t, "info.structured", am.JSON.Structured)
	missing(t, "info.error", am.JSON.Error)
}

// TestGlobalEventRealPayloadStepFinishPartStop 覆盖
// testdata/global_event_step_finish_part_stop.json，逐字段校验 message.part.updated
// + step-finish part（reason=stop）。
func TestGlobalEventRealPayloadStepFinishPartStop(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_step_finish_part_stop.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbc079a8001HfBylEtfwUuNna")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829816744))

	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbc079a80013YUZubd5fqYkr0")
	eq(t, "properties.part.messageID", carrier.MessageID, messageAssistID)
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeStepFinish)
	eq(t, "properties.part.reason", carrier.Reason, "stop")
	eq(t, "properties.part.cost", carrier.Cost, float64(0))

	sf := asVariant[opencode.StepFinishPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", sf.ID, "prt_fcbc079a80013YUZubd5fqYkr0")
	eq(t, "part.messageID", sf.MessageID, messageAssistID)
	eq(t, "part.sessionID", sf.SessionID, sessionMain)
	eq(t, "part.type", sf.Type, opencode.StepFinishPartTypeStepFinish)
	eq(t, "part.reason", sf.Reason, "stop")
	eq(t, "part.cost", sf.Cost, float64(0))
	eq(t, "part.tokens.total", sf.Tokens.Total, int64(72351))
	eq(t, "part.tokens.input", sf.Tokens.Input, int64(498))
	eq(t, "part.tokens.output", sf.Tokens.Output, int64(45))
	eq(t, "part.tokens.reasoning", sf.Tokens.Reasoning, int64(0))
	eq(t, "part.tokens.cache.write", sf.Tokens.Cache.Write, int64(0))
	eq(t, "part.tokens.cache.read", sf.Tokens.Cache.Read, int64(71808))

	present(t, "part.cost", sf.JSON.Cost)
	present(t, "part.tokens.reasoning", sf.Tokens.JSON.Reasoning)
	present(t, "part.tokens.cache.write", sf.Tokens.Cache.JSON.Write)

	eq(t, "part.snapshot", sf.Snapshot, "")
	missing(t, "part.snapshot", sf.JSON.Snapshot)
}

// TestGlobalEventRealPayloadToolPartPending 覆盖 testdata/global_event_tool_part_pending.json，
// 逐字段校验 message.part.updated + tool part（ToolState 联合体的 pending variant）。
func TestGlobalEventRealPayloadToolPartPending(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_tool_part_pending.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbbfe12d001BAyLkdZ4YUrpkq")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829777709))

	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbbfe12d001q7wKT0ndNdnUSR")
	eq(t, "properties.part.messageID", carrier.MessageID, "msg_fcbbfbbf2001sWFh4dT6ZnT5t1")
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeTool)
	eq(t, "properties.part.tool", carrier.Tool, "edit")
	eq(t, "properties.part.callID", carrier.CallID, "call-16bd3101-9ce2-44ea-936b-160841721a0e-34")

	tp := asVariant[opencode.ToolPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", tp.ID, "prt_fcbbfe12d001q7wKT0ndNdnUSR")
	eq(t, "part.messageID", tp.MessageID, "msg_fcbbfbbf2001sWFh4dT6ZnT5t1")
	eq(t, "part.sessionID", tp.SessionID, sessionMain)
	eq(t, "part.type", tp.Type, opencode.ToolPartTypeTool)
	if !tp.Type.IsKnown() {
		t.Errorf("part.type %q is not known", tp.Type)
	}
	eq(t, "part.tool", tp.Tool, "edit")
	eq(t, "part.callID", tp.CallID, "call-16bd3101-9ce2-44ea-936b-160841721a0e-34")
	eq(t, "part.metadata", len(tp.Metadata), 0)
	missing(t, "part.metadata", tp.JSON.Metadata)

	// ToolPartState 承载结构必须完整聚合 pending variant 的全部字段
	eq(t, "part.state.status", tp.State.Status, opencode.ToolPartStateStatusPending)
	if !tp.State.Status.IsKnown() {
		t.Errorf("part.state.status %q is not known", tp.State.Status)
	}
	eq(t, "part.state.raw", tp.State.Raw, "")
	stateInput := asVariant[map[string]any](t, "part.state.input", tp.State.Input)
	eq(t, "len(part.state.input)", len(stateInput), 0)
	present(t, "part.state.raw", tp.State.JSON.Raw)
	present(t, "part.state.input", tp.State.JSON.Input)

	// ToolStatePending variant
	pending := asVariant[opencode.ToolStatePending](t, "part.state", tp.State.AsUnion())
	eq(t, "state.status", pending.Status, opencode.ToolStatePendingStatusPending)
	if !pending.Status.IsKnown() {
		t.Errorf("state.status %q is not known", pending.Status)
	}
	eq(t, "state.raw", pending.Raw, "")
	eq(t, "len(state.input)", len(pending.Input), 0)
	// raw 为显式空串、input 为显式空对象，必须判定为「已下发」
	present(t, "state.raw", pending.JSON.Raw)
	present(t, "state.input", pending.JSON.Input)
}

// TestGlobalEventRealPayloadMessageUpdatedAssistantInFlight 覆盖
// testdata/global_event_message_updated_assistant_inflight.json，逐字段校验 message.updated
// + AssistantMessage（进行中：time 仅含 created，completed 缺失）。
func TestGlobalEventRealPayloadMessageUpdatedAssistantInFlight(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_message_updated_assistant_inflight.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	mu := asUnion[opencode.EventListResponseEventMessageUpdated](t, ev)
	eq(t, "payload.id", mu.ID, "evt_fcbc079ac001fkx1tXuUdImIAC")
	eq(t, "payload.type", mu.Type, opencode.EventListResponseEventMessageUpdatedTypeMessageUpdated)
	eq(t, "properties.sessionID", mu.Properties.SessionID, sessionMain)

	carrier := mu.Properties.Info
	eq(t, "properties.info.id", carrier.ID, messageAssistID)
	eq(t, "properties.info.role", carrier.Role, opencode.MessageRoleAssistant)
	eq(t, "properties.info.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.info.agent", carrier.Agent, "ai-coding-assistant")
	eq(t, "properties.info.mode", carrier.Mode, "ai-coding-assistant")
	eq(t, "properties.info.modelID", carrier.ModelID, "grok/grok-4.5")
	eq(t, "properties.info.providerID", carrier.ProviderID, "cf-anthropic")
	eq(t, "properties.info.parentID", carrier.ParentID, "msg_fcbb816b40017LmBKE1UFzBLrx")
	eq(t, "properties.info.cost", carrier.Cost, float64(0))
	eq(t, "properties.info.finish", carrier.Finish, "stop")

	am := asVariant[opencode.AssistantMessage](t, "properties.info", carrier.AsUnion())
	eq(t, "info.id", am.ID, messageAssistID)
	eq(t, "info.parentID", am.ParentID, "msg_fcbb816b40017LmBKE1UFzBLrx")
	eq(t, "info.role", am.Role, opencode.AssistantMessageRoleAssistant)
	eq(t, "info.sessionID", am.SessionID, sessionMain)
	eq(t, "info.mode", am.Mode, "ai-coding-assistant")
	eq(t, "info.agent", am.Agent, "ai-coding-assistant")
	eq(t, "info.modelID", am.ModelID, "grok/grok-4.5")
	eq(t, "info.providerID", am.ProviderID, "cf-anthropic")
	eq(t, "info.cost", am.Cost, float64(0))
	eq(t, "info.finish", am.Finish, "stop")
	eq(t, "info.path.cwd", am.Path.Cwd, directoryMain)
	eq(t, "info.path.root", am.Path.Root, directoryMain)
	eq(t, "info.tokens.total", am.Tokens.Total, int64(72351))
	eq(t, "info.tokens.input", am.Tokens.Input, int64(498))
	eq(t, "info.tokens.output", am.Tokens.Output, int64(45))
	eq(t, "info.tokens.reasoning", am.Tokens.Reasoning, int64(0))
	eq(t, "info.tokens.cache.write", am.Tokens.Cache.Write, int64(0))
	eq(t, "info.tokens.cache.read", am.Tokens.Cache.Read, int64(71808))

	// 🔴 边界：time.completed 在 OpenAPI/JS SDK(v2) 中可选，本 payload 未下发。
	// Go SDK 不使用指针，缺失即零值，必须靠 JSON 元数据区分「缺失」与「显式 0」。
	eq(t, "info.time.created", am.Time.Created, int64(1785829813031))
	eq(t, "info.time.completed", am.Time.Completed, int64(0))
	present(t, "info.time.created", am.Time.JSON.Created)
	missing(t, "info.time.completed", am.Time.JSON.Completed)

	// 对照：同为零值但已下发的字段
	present(t, "info.cost", am.JSON.Cost)
	present(t, "info.tokens.reasoning", am.Tokens.JSON.Reasoning)
	present(t, "info.tokens.cache.write", am.Tokens.Cache.JSON.Write)

	// payload 未下发的可选字段
	eq(t, "info.variant", am.Variant, "")
	eq(t, "info.summary", am.Summary, false)
	if am.Structured != nil {
		t.Errorf("info.structured = %#v, want nil", am.Structured)
	}
	missing(t, "info.variant", am.JSON.Variant)
	missing(t, "info.summary", am.JSON.Summary)
	missing(t, "info.structured", am.JSON.Structured)
	missing(t, "info.error", am.JSON.Error)
}

// TestGlobalEventRealPayloadTextPart 覆盖 testdata/global_event_text_part.json，
// 逐字段校验 message.part.updated + text part（与 reasoning part 结构高度相似，验证联合体判别）。
func TestGlobalEventRealPayloadTextPart(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_text_part.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbbe7e6d001vJ53BF2apLxGA0")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829686893))

	const wantText = "正在编写设置与完整上传界面组件。"

	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbbda8c50014VSF6PT7PFpqi3")
	eq(t, "properties.part.messageID", carrier.MessageID, "msg_fcbbd8c04001Ng0bz0ST0DLHFs")
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeText)
	eq(t, "properties.part.text", carrier.Text, wantText)

	// 🔴 TextPart 与 ReasoningPart 字段几乎一致，PartUnion 注册时无判别器，
	// 必须依赖 type 枚举严格性落到 TextPart 而非 ReasoningPart。
	if _, isReasoning := carrier.AsUnion().(opencode.ReasoningPart); isReasoning {
		t.Fatalf("text part was mis-resolved to ReasoningPart")
	}
	tp := asVariant[opencode.TextPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", tp.ID, "prt_fcbbda8c50014VSF6PT7PFpqi3")
	eq(t, "part.messageID", tp.MessageID, "msg_fcbbd8c04001Ng0bz0ST0DLHFs")
	eq(t, "part.sessionID", tp.SessionID, sessionMain)
	eq(t, "part.type", tp.Type, opencode.TextPartTypeText)
	if !tp.Type.IsKnown() {
		t.Errorf("part.type %q is not known", tp.Type)
	}
	eq(t, "part.text", tp.Text, wantText)
	eq(t, "part.time.start", tp.Time.Start, int64(1785829632197))
	eq(t, "part.time.end", tp.Time.End, int64(1785829686892))
	present(t, "part.time", tp.JSON.Time)
	present(t, "part.time.start", tp.Time.JSON.Start)
	present(t, "part.time.end", tp.Time.JSON.End)

	// payload 未下发的可选字段
	eq(t, "part.synthetic", tp.Synthetic, false)
	eq(t, "part.ignored", tp.Ignored, false)
	eq(t, "part.metadata", len(tp.Metadata), 0)
	missing(t, "part.synthetic", tp.JSON.Synthetic)
	missing(t, "part.ignored", tp.JSON.Ignored)
	missing(t, "part.metadata", tp.JSON.Metadata)
}

// TestGlobalEventRealPayloadReasoningPartStarting 覆盖
// testdata/global_event_reasoning_part_starting.json，逐字段校验 message.part.updated
// + reasoning part（刚开始：text 为显式空串，time 仅含 start）。
func TestGlobalEventRealPayloadReasoningPartStarting(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_reasoning_part_starting.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbbe8a2b001oH7JuCHbbWqB0V")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829689899))

	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbbe8a2b001wPMT2wBiEwoAGd")
	eq(t, "properties.part.messageID", carrier.MessageID, "msg_fcbbe7ec40012yp6E0P3CPk0hs")
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeReasoning)
	eq(t, "properties.part.text", carrier.Text, "")

	// 🔴 text 为空串时仍必须落到 ReasoningPart，不可被 TextPart 抢占
	if _, isText := carrier.AsUnion().(opencode.TextPart); isText {
		t.Fatalf("reasoning part was mis-resolved to TextPart")
	}
	rp := asVariant[opencode.ReasoningPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", rp.ID, "prt_fcbbe8a2b001wPMT2wBiEwoAGd")
	eq(t, "part.messageID", rp.MessageID, "msg_fcbbe7ec40012yp6E0P3CPk0hs")
	eq(t, "part.sessionID", rp.SessionID, sessionMain)
	eq(t, "part.type", rp.Type, opencode.ReasoningPartTypeReasoning)
	if !rp.Type.IsKnown() {
		t.Errorf("part.type %q is not known", rp.Type)
	}

	// 🔴 边界：text 是 required 字段但取空串，必须判定为「已下发的空值」而非「缺失」
	eq(t, "part.text", rp.Text, "")
	present(t, "part.text", rp.JSON.Text)
	eq(t, "part.text raw", rp.JSON.Text.Raw(), `""`)

	// 边界：time.start 已下发；time.end 可选未下发 → 零值且标记缺失
	eq(t, "part.time.start", rp.Time.Start, int64(1785829689899))
	eq(t, "part.time.end", rp.Time.End, int64(0))
	present(t, "part.time", rp.JSON.Time)
	present(t, "part.time.start", rp.Time.JSON.Start)
	missing(t, "part.time.end", rp.Time.JSON.End)

	// payload 未下发的可选字段
	eq(t, "part.metadata", len(rp.Metadata), 0)
	missing(t, "part.metadata", rp.JSON.Metadata)
}

// TestGlobalEventRealPayloadToolPartCompleted 覆盖 testdata/global_event_tool_part_completed.json，
// 逐字段校验 message.part.updated + tool part（ToolState 联合体的 completed variant）。
// 注意：OpenAPI 将 ToolStateCompleted.metadata 标为 required，但生产端未下发该字段，
// 联合体判别必须仍能落到 completed variant。
func TestGlobalEventRealPayloadToolPartCompleted(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_tool_part_completed.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbbeae7d001rxnEG4oLFSsXuq")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829699197))

	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbbeae51001R17TlXhvmkBBzX")
	eq(t, "properties.part.messageID", carrier.MessageID, "msg_fcbbe7ec40012yp6E0P3CPk0hs")
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeTool)
	eq(t, "properties.part.tool", carrier.Tool, "write")
	eq(t, "properties.part.callID", carrier.CallID, "call-fd43e9bd-11b6-4f6c-9458-213c74886821-30")

	tp := asVariant[opencode.ToolPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", tp.ID, "prt_fcbbeae51001R17TlXhvmkBBzX")
	eq(t, "part.messageID", tp.MessageID, "msg_fcbbe7ec40012yp6E0P3CPk0hs")
	eq(t, "part.sessionID", tp.SessionID, sessionMain)
	eq(t, "part.type", tp.Type, opencode.ToolPartTypeTool)
	eq(t, "part.tool", tp.Tool, "write")
	eq(t, "part.callID", tp.CallID, "call-fd43e9bd-11b6-4f6c-9458-213c74886821-30")
	eq(t, "len(part.metadata)", len(tp.Metadata), 0)
	missing(t, "part.metadata", tp.JSON.Metadata)

	// ToolPartState 承载结构必须完整聚合 completed variant 的全部字段
	eq(t, "part.state.status", tp.State.Status, opencode.ToolPartStateStatusCompleted)
	if !tp.State.Status.IsKnown() {
		t.Errorf("part.state.status %q is not known", tp.State.Status)
	}
	eq(t, "part.state.output", tp.State.Output, "Wrote file successfully.")
	eq(t, "part.state.title", tp.State.Title, "custom-photo-upload-preview.tsx")
	carrierInput := asVariant[map[string]any](t, "part.state.input", tp.State.Input)
	eq(t, "len(part.state.input)", len(carrierInput), 2)
	eq(t, "part.state.input.filePath", carrierInput["filePath"], any("custom-photo-upload-preview.tsx"))
	eq(t, "part.state.input.content", carrierInput["content"], any(""))
	carrierTime := asVariant[opencode.ToolStateCompletedTime](t, "part.state.time", tp.State.Time)
	eq(t, "part.state.time.start", carrierTime.Start, int64(1785829699178))
	eq(t, "part.state.time.end", carrierTime.End, int64(1785829699196))

	// ToolStateCompleted variant
	done := asVariant[opencode.ToolStateCompleted](t, "part.state", tp.State.AsUnion())
	eq(t, "state.status", done.Status, opencode.ToolStateCompletedStatusCompleted)
	if !done.Status.IsKnown() {
		t.Errorf("state.status %q is not known", done.Status)
	}
	eq(t, "state.output", done.Output, "Wrote file successfully.")
	eq(t, "state.title", done.Title, "custom-photo-upload-preview.tsx")
	eq(t, "len(state.input)", len(done.Input), 2)
	eq(t, "state.input.filePath", done.Input["filePath"], any("custom-photo-upload-preview.tsx"))
	eq(t, "state.input.content", done.Input["content"], any(""))
	eq(t, "state.time.start", done.Time.Start, int64(1785829699178))
	eq(t, "state.time.end", done.Time.End, int64(1785829699196))
	present(t, "state.output", done.JSON.Output)
	present(t, "state.title", done.JSON.Title)
	present(t, "state.input", done.JSON.Input)
	present(t, "state.time", done.JSON.Time)
	present(t, "state.time.start", done.Time.JSON.Start)
	present(t, "state.time.end", done.Time.JSON.End)
	// input.content 为显式空串，必须判定为「已下发」
	if _, ok := done.Input["content"]; !ok {
		t.Errorf("state.input.content key should exist even when empty")
	}

	// payload 未下发的可选/未填字段
	eq(t, "len(state.metadata)", len(done.Metadata), 0)
	eq(t, "len(state.attachments)", len(done.Attachments), 0)
	eq(t, "state.time.compacted", done.Time.Compacted, int64(0))
	missing(t, "state.metadata", done.JSON.Metadata)
	missing(t, "state.attachments", done.JSON.Attachments)
	missing(t, "state.time.compacted", done.Time.JSON.Compacted)
	// completed variant 不含 raw / error，承载结构对应字段必须为零值
	eq(t, "part.state.raw", tp.State.Raw, "")
	eq(t, "part.state.error", tp.State.Error, "")
}

// TestGlobalEventRealPayloadToolPartRunning 覆盖 testdata/global_event_tool_part_running.json，
// 逐字段校验 message.part.updated + tool part（ToolState 联合体的 running variant，
// input 内嵌 6 条 todo 数组，time 仅含 start）。
func TestGlobalEventRealPayloadToolPartRunning(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_tool_part_running.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbc06af5001jE1PdsQzXbr9y9")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829812981))

	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbc06aeb001ssEf7L6Lw30DG9")
	eq(t, "properties.part.messageID", carrier.MessageID, "msg_fcbc0574a001k6zg8RxJLvHHlX")
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeTool)
	eq(t, "properties.part.tool", carrier.Tool, "todowrite")
	eq(t, "properties.part.callID", carrier.CallID, "call-e38f2040-27ad-40eb-8e34-8f2c8f6364e7-38")

	tp := asVariant[opencode.ToolPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", tp.ID, "prt_fcbc06aeb001ssEf7L6Lw30DG9")
	eq(t, "part.messageID", tp.MessageID, "msg_fcbc0574a001k6zg8RxJLvHHlX")
	eq(t, "part.sessionID", tp.SessionID, sessionMain)
	eq(t, "part.type", tp.Type, opencode.ToolPartTypeTool)
	eq(t, "part.tool", tp.Tool, "todowrite")
	eq(t, "part.callID", tp.CallID, "call-e38f2040-27ad-40eb-8e34-8f2c8f6364e7-38")
	eq(t, "len(part.metadata)", len(tp.Metadata), 0)
	missing(t, "part.metadata", tp.JSON.Metadata)

	// ToolPartState 承载结构
	eq(t, "part.state.status", tp.State.Status, opencode.ToolPartStateStatusRunning)
	if !tp.State.Status.IsKnown() {
		t.Errorf("part.state.status %q is not known", tp.State.Status)
	}
	carrierTime := asVariant[opencode.ToolStateRunningTime](t, "part.state.time", tp.State.Time)
	eq(t, "part.state.time.start", carrierTime.Start, int64(1785829812981))

	// ToolStateRunning variant（不可被 pending/completed 抢占）
	running := asVariant[opencode.ToolStateRunning](t, "part.state", tp.State.AsUnion())
	eq(t, "state.status", running.Status, opencode.ToolStateRunningStatusRunning)
	if !running.Status.IsKnown() {
		t.Errorf("state.status %q is not known", running.Status)
	}
	eq(t, "state.time.start", running.Time.Start, int64(1785829812981))
	present(t, "state.input", running.JSON.Input)
	present(t, "state.time", running.JSON.Time)
	present(t, "state.time.start", running.Time.JSON.Start)

	// state.input.todos — 逐条校验全部 6 个元素的 3 个字段
	eq(t, "len(state.input)", len(running.Input), 1)
	todos := asVariant[[]any](t, "state.input.todos", running.Input["todos"])
	eq(t, "len(state.input.todos)", len(todos), 6)
	wantTodos := []string{
		"Configure upload section settings and merchant options",
		"Build main custom photo upload interface",
		"Build image grid and product mock preview pieces",
		"Connect page with editable options",
		"Generate platform code and compile",
		"Mark codebase version",
	}
	for i, wantContent := range wantTodos {
		if i >= len(todos) {
			break
		}
		field := fmt.Sprintf("state.input.todos[%d]", i)
		todo := asVariant[map[string]any](t, field, todos[i])
		eq(t, field+" len", len(todo), 3)
		eq(t, field+".content", todo["content"], any(wantContent))
		eq(t, field+".status", todo["status"], any("completed"))
		eq(t, field+".priority", todo["priority"], any("high"))
	}

	// payload 未下发的可选字段
	eq(t, "state.title", running.Title, "")
	eq(t, "len(state.metadata)", len(running.Metadata), 0)
	missing(t, "state.title", running.JSON.Title)
	missing(t, "state.metadata", running.JSON.Metadata)
	// running variant 不含 raw / output / error，承载结构对应字段必须为零值
	eq(t, "part.state.raw", tp.State.Raw, "")
	eq(t, "part.state.output", tp.State.Output, "")
	eq(t, "part.state.title", tp.State.Title, "")
	eq(t, "part.state.error", tp.State.Error, "")
}

// TestGlobalEventRealPayloadToolPartRunningEdit 覆盖
// testdata/global_event_tool_part_running_edit.json，逐字段校验 message.part.updated
// + tool part（edit 工具，ToolState 联合体的 running variant，input 为 3 个字符串键的扁平对象，
// time 仅含 start）。
func TestGlobalEventRealPayloadToolPartRunningEdit(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_tool_part_running_edit.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbbfd879001uMa2hBgrOMvZqB")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	if !pu.Type.IsKnown() {
		t.Errorf("payload.type %q is not known", pu.Type)
	}
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829775481))

	const (
		wantFilePath  = "custom-photo-upload.tsx"
		wantOldString = "import { useCallback, useId, useRef, useState } from 'react';"
		wantNewString = "import { useCallback, useId, useRef, useState, type DragEvent } from 'react';"
	)

	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbbfd853001k0fG4geBV7a8Va")
	eq(t, "properties.part.messageID", carrier.MessageID, "msg_fcbbfbbf2001sWFh4dT6ZnT5t1")
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeTool)
	eq(t, "properties.part.tool", carrier.Tool, "edit")
	eq(t, "properties.part.callID", carrier.CallID, "call-16bd3101-9ce2-44ea-936b-160841721a0e-33")

	tp := asVariant[opencode.ToolPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", tp.ID, "prt_fcbbfd853001k0fG4geBV7a8Va")
	eq(t, "part.messageID", tp.MessageID, "msg_fcbbfbbf2001sWFh4dT6ZnT5t1")
	eq(t, "part.sessionID", tp.SessionID, sessionMain)
	eq(t, "part.type", tp.Type, opencode.ToolPartTypeTool)
	if !tp.Type.IsKnown() {
		t.Errorf("part.type %q is not known", tp.Type)
	}
	eq(t, "part.tool", tp.Tool, "edit")
	eq(t, "part.callID", tp.CallID, "call-16bd3101-9ce2-44ea-936b-160841721a0e-33")
	eq(t, "len(part.metadata)", len(tp.Metadata), 0)
	missing(t, "part.metadata", tp.JSON.Metadata)

	// ToolPartState 承载结构必须完整聚合 running variant 的全部字段
	eq(t, "part.state.status", tp.State.Status, opencode.ToolPartStateStatusRunning)
	if !tp.State.Status.IsKnown() {
		t.Errorf("part.state.status %q is not known", tp.State.Status)
	}
	carrierInput := asVariant[map[string]any](t, "part.state.input", tp.State.Input)
	eq(t, "len(part.state.input)", len(carrierInput), 3)
	eq(t, "part.state.input.filePath", carrierInput["filePath"], any(wantFilePath))
	eq(t, "part.state.input.oldString", carrierInput["oldString"], any(wantOldString))
	eq(t, "part.state.input.newString", carrierInput["newString"], any(wantNewString))
	carrierTime := asVariant[opencode.ToolStateRunningTime](t, "part.state.time", tp.State.Time)
	eq(t, "part.state.time.start", carrierTime.Start, int64(1785829775481))
	present(t, "part.state.input", tp.State.JSON.Input)
	present(t, "part.state.time", tp.State.JSON.Time)

	// 🔴 running 与 pending 均含 input，且本 payload 未下发 raw，
	// 联合体判别必须依赖 status 枚举严格性落到 ToolStateRunning。
	if _, isPending := tp.State.AsUnion().(opencode.ToolStatePending); isPending {
		t.Fatalf("running state was mis-resolved to ToolStatePending")
	}
	if _, isCompleted := tp.State.AsUnion().(opencode.ToolStateCompleted); isCompleted {
		t.Fatalf("running state was mis-resolved to ToolStateCompleted")
	}
	running := asVariant[opencode.ToolStateRunning](t, "part.state", tp.State.AsUnion())
	eq(t, "state.status", running.Status, opencode.ToolStateRunningStatusRunning)
	if !running.Status.IsKnown() {
		t.Errorf("state.status %q is not known", running.Status)
	}
	eq(t, "len(state.input)", len(running.Input), 3)
	eq(t, "state.input.filePath", running.Input["filePath"], any(wantFilePath))
	eq(t, "state.input.oldString", running.Input["oldString"], any(wantOldString))
	eq(t, "state.input.newString", running.Input["newString"], any(wantNewString))
	eq(t, "state.time.start", running.Time.Start, int64(1785829775481))
	present(t, "state.input", running.JSON.Input)
	present(t, "state.time", running.JSON.Time)
	present(t, "state.time.start", running.Time.JSON.Start)

	// payload 未下发的可选字段
	eq(t, "state.title", running.Title, "")
	eq(t, "len(state.metadata)", len(running.Metadata), 0)
	missing(t, "state.title", running.JSON.Title)
	missing(t, "state.metadata", running.JSON.Metadata)
	// running variant 不含 raw / output / error / attachments，承载结构对应字段必须为零值
	eq(t, "part.state.raw", tp.State.Raw, "")
	eq(t, "part.state.output", tp.State.Output, "")
	eq(t, "part.state.title", tp.State.Title, "")
	eq(t, "part.state.error", tp.State.Error, "")
	missing(t, "part.state.raw", tp.State.JSON.Raw)
	missing(t, "part.state.output", tp.State.JSON.Output)
	missing(t, "part.state.title", tp.State.JSON.Title)
	missing(t, "part.state.error", tp.State.JSON.Error)
	missing(t, "part.state.metadata", tp.State.JSON.Metadata)
	missing(t, "part.state.attachments", tp.State.JSON.Attachments)
}

// TestGlobalEventRealPayloadToolPartCompletedQuestion 覆盖
// testdata/global_event_tool_part_completed_question.json，逐字段校验 message.part.updated
// + tool part（question 工具，ToolState 联合体的 completed variant）。
// 该 payload 是 completed variant 字段最完整的真实样本：input 为三层嵌套结构
// （questions[] → options[]），metadata 含二维数组 answers 与显式 false 的 truncated，
// output/title/time.start/time.end 全部下发。
func TestGlobalEventRealPayloadToolPartCompletedQuestion(t *testing.T) {
	ev := decodeGlobalEvent(t, "global_event_tool_part_completed_question.json")

	eq(t, "directory", ev.Directory, directoryMain)
	eq(t, "project", ev.Project, projectMain)

	pu := asUnion[opencode.EventListResponseEventMessagePartUpdated](t, ev)
	eq(t, "payload.id", pu.ID, "evt_fcbbc9952001PKetOYc1pMRcKt")
	eq(t, "payload.type", pu.Type, opencode.EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated)
	if !pu.Type.IsKnown() {
		t.Errorf("payload.type %q is not known", pu.Type)
	}
	eq(t, "properties.sessionID", pu.Properties.SessionID, sessionMain)
	eq(t, "properties.time", pu.Properties.Time, int64(1785829562706))

	const (
		wantQuestion = "图片需要保存到云端，请选择后端方式（选一项即开始制作）："
		wantOutput   = `User has answered your questions: "图片需要保存到云端，请选择后端方式（选一项即开始制作）："="暂不启用云服务". You can now continue with the user's answers in mind...`
		wantTitle    = "Asked 1 question"
	)

	carrier := pu.Properties.Part
	eq(t, "properties.part.id", carrier.ID, "prt_fcbbc8a14001uwcOBwsQRX1Caf")
	eq(t, "properties.part.messageID", carrier.MessageID, "msg_fcbbc6781001K8g8Xrn8R3uAQH")
	eq(t, "properties.part.sessionID", carrier.SessionID, sessionMain)
	eq(t, "properties.part.type", carrier.Type, opencode.PartTypeTool)
	eq(t, "properties.part.tool", carrier.Tool, "question")
	eq(t, "properties.part.callID", carrier.CallID, "call-7e85e941-ef6b-48f9-8cad-acddba23b29e-13")

	tp := asVariant[opencode.ToolPart](t, "properties.part", carrier.AsUnion())
	eq(t, "part.id", tp.ID, "prt_fcbbc8a14001uwcOBwsQRX1Caf")
	eq(t, "part.messageID", tp.MessageID, "msg_fcbbc6781001K8g8Xrn8R3uAQH")
	eq(t, "part.sessionID", tp.SessionID, sessionMain)
	eq(t, "part.type", tp.Type, opencode.ToolPartTypeTool)
	eq(t, "part.tool", tp.Tool, "question")
	eq(t, "part.callID", tp.CallID, "call-7e85e941-ef6b-48f9-8cad-acddba23b29e-13")
	eq(t, "len(part.metadata)", len(tp.Metadata), 0)
	missing(t, "part.metadata", tp.JSON.Metadata)

	// ToolPartState 承载结构必须完整聚合 completed variant 的全部字段
	eq(t, "part.state.status", tp.State.Status, opencode.ToolPartStateStatusCompleted)
	if !tp.State.Status.IsKnown() {
		t.Errorf("part.state.status %q is not known", tp.State.Status)
	}
	eq(t, "part.state.output", tp.State.Output, wantOutput)
	eq(t, "part.state.title", tp.State.Title, wantTitle)
	carrierTime := asVariant[opencode.ToolStateCompletedTime](t, "part.state.time", tp.State.Time)
	eq(t, "part.state.time.start", carrierTime.Start, int64(1785829558816))
	eq(t, "part.state.time.end", carrierTime.End, int64(1785829562706))
	carrierMetadata := asVariant[map[string]any](t, "part.state.metadata", tp.State.Metadata)
	eq(t, "len(part.state.metadata)", len(carrierMetadata), 2)
	eq(t, "part.state.metadata.truncated", carrierMetadata["truncated"], any(false))
	present(t, "part.state.input", tp.State.JSON.Input)
	present(t, "part.state.metadata", tp.State.JSON.Metadata)
	present(t, "part.state.output", tp.State.JSON.Output)
	present(t, "part.state.title", tp.State.JSON.Title)
	present(t, "part.state.time", tp.State.JSON.Time)

	// 🔴 completed 与 running/error 均含 input + time，
	// 联合体判别必须依赖 status 枚举严格性落到 ToolStateCompleted。
	if _, isRunning := tp.State.AsUnion().(opencode.ToolStateRunning); isRunning {
		t.Fatalf("completed state was mis-resolved to ToolStateRunning")
	}
	if _, isError := tp.State.AsUnion().(opencode.ToolStateError); isError {
		t.Fatalf("completed state was mis-resolved to ToolStateError")
	}
	done := asVariant[opencode.ToolStateCompleted](t, "part.state", tp.State.AsUnion())
	eq(t, "state.status", done.Status, opencode.ToolStateCompletedStatusCompleted)
	if !done.Status.IsKnown() {
		t.Errorf("state.status %q is not known", done.Status)
	}
	eq(t, "state.output", done.Output, wantOutput)
	eq(t, "state.title", done.Title, wantTitle)
	eq(t, "state.time.start", done.Time.Start, int64(1785829558816))
	eq(t, "state.time.end", done.Time.End, int64(1785829562706))
	present(t, "state.output", done.JSON.Output)
	present(t, "state.title", done.JSON.Title)
	present(t, "state.input", done.JSON.Input)
	present(t, "state.metadata", done.JSON.Metadata)
	present(t, "state.time", done.JSON.Time)
	present(t, "state.time.start", done.Time.JSON.Start)
	present(t, "state.time.end", done.Time.JSON.End)

	// state.input.questions — 三层嵌套：questions[0].options[0..2]
	eq(t, "len(state.input)", len(done.Input), 1)
	questions := asVariant[[]any](t, "state.input.questions", done.Input["questions"])
	eq(t, "len(state.input.questions)", len(questions), 1)
	question := asVariant[map[string]any](t, "state.input.questions[0]", questions[0])
	eq(t, "len(state.input.questions[0])", len(question), 3)
	eq(t, "state.input.questions[0].header", question["header"], any("cloud-setup"))
	eq(t, "state.input.questions[0].question", question["question"], any(wantQuestion))

	options := asVariant[[]any](t, "state.input.questions[0].options", question["options"])
	eq(t, "len(state.input.questions[0].options)", len(options), 3)
	wantOptions := []struct {
		label       string
		description string
	}{
		{"启用 Meowgic Cloud（推荐）", "一键获得数据库、文件存储等能力，数据性能更好。启用会产生费用——预估 $10/月起，按计算与存储用量计费。"},
		{"连接我的 Supabase 项目", "把应用快速连到已有的 Supabase 项目。Supabase 有免费档，适合测试和验证。"},
		{"暂不启用云服务", "先只做前端界面，以后再开云服务。注意：不启用则图片无法真正保存。"},
	}
	for i, want := range wantOptions {
		if i >= len(options) {
			break
		}
		field := fmt.Sprintf("state.input.questions[0].options[%d]", i)
		option := asVariant[map[string]any](t, field, options[i])
		eq(t, field+" len", len(option), 2)
		eq(t, field+".label", option["label"], any(want.label))
		eq(t, field+".description", option["description"], any(want.description))
	}

	// state.metadata — answers 为二维数组，truncated 为显式 false
	eq(t, "len(state.metadata)", len(done.Metadata), 2)
	answers := asVariant[[]any](t, "state.metadata.answers", done.Metadata["answers"])
	eq(t, "len(state.metadata.answers)", len(answers), 1)
	answer := asVariant[[]any](t, "state.metadata.answers[0]", answers[0])
	eq(t, "len(state.metadata.answers[0])", len(answer), 1)
	eq(t, "state.metadata.answers[0][0]", answer[0], any("暂不启用云服务"))
	// 🔴 边界：truncated 为显式 false，键必须存在，不可与「字段缺失」混淆
	truncated, ok := done.Metadata["truncated"]
	if !ok {
		t.Errorf("state.metadata.truncated key should exist even when false")
	}
	eq(t, "state.metadata.truncated", truncated, any(false))

	// payload 未下发的可选/未填字段
	eq(t, "len(state.attachments)", len(done.Attachments), 0)
	eq(t, "state.time.compacted", done.Time.Compacted, int64(0))
	missing(t, "state.attachments", done.JSON.Attachments)
	missing(t, "state.time.compacted", done.Time.JSON.Compacted)
	// completed variant 不含 raw / error，承载结构对应字段必须为零值
	eq(t, "part.state.raw", tp.State.Raw, "")
	eq(t, "part.state.error", tp.State.Error, "")
	missing(t, "part.state.raw", tp.State.JSON.Raw)
	missing(t, "part.state.error", tp.State.JSON.Error)
	missing(t, "part.state.attachments", tp.State.JSON.Attachments)
}
