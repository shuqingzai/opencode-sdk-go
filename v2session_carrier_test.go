package opencode

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/sst/opencode-sdk-go/internal/apijson"
)

// This file verifies the promoted (hoisted) fields on the four v2 session
// response-union carriers: [V2SessionMessage], [V2SessionMessageAssistantContent],
// [V2SessionMessageToolState] and [V2SessionDurableEvent].
//
// Each carrier decodes via `apijson.UnmarshalRoot(data, &r.union)` followed by
// `apijson.Port(r.union, &r)`. Port only transports variant fields whose `json`
// tag name matches a carrier field, so every promoted field is asserted here for
// every variant of every union. Mismatched carrier field types make Port panic
// (reflect Set/SetString on a non-assignable destination), so these tests double
// as the type-compatibility guard.

// wantField asserts a promoted carrier field equals want.
func wantField(t *testing.T, carrier, variant, field string, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s[%s].%s = %#v (%T), want %#v (%T)", carrier, variant, field, got, got, want, want)
	}
}

// wantZero asserts a promoted carrier field is the zero value, i.e. the field
// belongs to a different variant and Port correctly left it untouched.
func wantZero(t *testing.T, carrier, variant, field string, got any) {
	t.Helper()
	if got == nil {
		return
	}
	if !reflect.ValueOf(got).IsZero() {
		t.Errorf("%s[%s].%s = %#v, want zero value", carrier, variant, field, got)
	}
}

// ===== V2SessionMessage (8 variants, 22 promoted fields) =====

func TestV2SessionMessageCarrierPromotedFields(t *testing.T) {
	t.Run("agent-switched", func(t *testing.T) {
		var v V2SessionMessage
		raw := `{"id":"msg_1","type":"agent-switched","time":{"created":11},"agent":"build","metadata":{"k":"v"}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.AsUnion().(V2SessionMessageAgentSwitched); !ok {
			t.Fatalf("union = %T, want V2SessionMessageAgentSwitched", v.AsUnion())
		}
		wantField(t, "V2SessionMessage", "agent-switched", "ID", v.ID, "msg_1")
		wantField(t, "V2SessionMessage", "agent-switched", "Type", v.Type, "agent-switched")
		wantField(t, "V2SessionMessage", "agent-switched", "Agent", v.Agent, "build")
		wantField(t, "V2SessionMessage", "agent-switched", "Metadata", v.Metadata, map[string]any{"k": "v"})
		tm, ok := v.Time.(V2SessionMessageTime)
		if !ok {
			t.Fatalf("Time runtime type = %T, want V2SessionMessageTime", v.Time)
		}
		wantField(t, "V2SessionMessage", "agent-switched", "Time.Created", tm.Created, int64(11))
		// fields owned by other variants stay zero
		wantZero(t, "V2SessionMessage", "agent-switched", "Text", v.Text)
		wantZero(t, "V2SessionMessage", "agent-switched", "Model", v.Model)
		wantZero(t, "V2SessionMessage", "agent-switched", "Cost", v.Cost)
		wantZero(t, "V2SessionMessage", "agent-switched", "Reason", v.Reason)
		wantZero(t, "V2SessionMessage", "agent-switched", "SessionID", v.SessionID)
	})

	t.Run("model-switched", func(t *testing.T) {
		var v V2SessionMessage
		raw := `{"id":"msg_2","type":"model-switched","time":{"created":12},"model":{"id":"gpt","providerID":"openai","variant":"mini"},"metadata":{}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "V2SessionMessage", "model-switched", "ID", v.ID, "msg_2")
		wantField(t, "V2SessionMessage", "model-switched", "Type", v.Type, "model-switched")
		m, ok := v.Model.(V2SessionMessageModel)
		if !ok {
			t.Fatalf("Model runtime type = %T, want V2SessionMessageModel", v.Model)
		}
		wantField(t, "V2SessionMessage", "model-switched", "Model.ID", m.ID, "gpt")
		wantField(t, "V2SessionMessage", "model-switched", "Model.ProviderID", m.ProviderID, "openai")
		wantField(t, "V2SessionMessage", "model-switched", "Model.Variant", m.Variant, "mini")
		wantZero(t, "V2SessionMessage", "model-switched", "Agent", v.Agent)
		wantZero(t, "V2SessionMessage", "model-switched", "Text", v.Text)
	})

	t.Run("user", func(t *testing.T) {
		var v V2SessionMessage
		raw := `{"id":"msg_3","type":"user","time":{"created":13},"text":"hello",
			"files":[{"uri":"file:///a.txt","mime":"text/plain","name":"a.txt"}],
			"agents":[{"name":"reviewer"}],"metadata":{"n":1}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "V2SessionMessage", "user", "ID", v.ID, "msg_3")
		wantField(t, "V2SessionMessage", "user", "Type", v.Type, "user")
		wantField(t, "V2SessionMessage", "user", "Text", v.Text, "hello")
		files, ok := v.Files.([]V2PromptFileAttachment)
		if !ok {
			t.Fatalf("Files runtime type = %T, want []V2PromptFileAttachment", v.Files)
		}
		if len(files) != 1 || files[0].Mime != "text/plain" || files[0].URI != "file:///a.txt" || files[0].Name != "a.txt" {
			t.Errorf("Files = %#v", files)
		}
		agents, ok := v.Agents.([]V2PromptAgentAttachment)
		if !ok {
			t.Fatalf("Agents runtime type = %T, want []V2PromptAgentAttachment", v.Agents)
		}
		if len(agents) != 1 || agents[0].Name != "reviewer" {
			t.Errorf("Agents = %#v", agents)
		}
		wantField(t, "V2SessionMessage", "user", "Metadata", v.Metadata, map[string]any{"n": 1.0})
		wantZero(t, "V2SessionMessage", "user", "SessionID", v.SessionID)
		wantZero(t, "V2SessionMessage", "user", "Output", v.Output)
	})

	t.Run("synthetic", func(t *testing.T) {
		var v V2SessionMessage
		raw := `{"id":"msg_4","type":"synthetic","time":{"created":14},"sessionID":"ses_1","text":"syn","metadata":{}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "V2SessionMessage", "synthetic", "ID", v.ID, "msg_4")
		wantField(t, "V2SessionMessage", "synthetic", "Type", v.Type, "synthetic")
		wantField(t, "V2SessionMessage", "synthetic", "SessionID", v.SessionID, "ses_1")
		wantField(t, "V2SessionMessage", "synthetic", "Text", v.Text, "syn")
		wantZero(t, "V2SessionMessage", "synthetic", "Agent", v.Agent)
	})

	t.Run("system", func(t *testing.T) {
		var v V2SessionMessage
		raw := `{"id":"msg_5","type":"system","time":{"created":15},"text":"sys","metadata":{}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.AsUnion().(V2SessionMessageSystem); !ok {
			t.Fatalf("union = %T, want V2SessionMessageSystem", v.AsUnion())
		}
		wantField(t, "V2SessionMessage", "system", "ID", v.ID, "msg_5")
		wantField(t, "V2SessionMessage", "system", "Type", v.Type, "system")
		wantField(t, "V2SessionMessage", "system", "Text", v.Text, "sys")
		wantZero(t, "V2SessionMessage", "system", "SessionID", v.SessionID)
	})

	t.Run("shell", func(t *testing.T) {
		var v V2SessionMessage
		raw := `{"id":"msg_6","type":"shell","time":{"created":16,"completed":17},"callID":"call_1","command":"ls -la","output":"total 0","metadata":{}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "V2SessionMessage", "shell", "ID", v.ID, "msg_6")
		wantField(t, "V2SessionMessage", "shell", "Type", v.Type, "shell")
		wantField(t, "V2SessionMessage", "shell", "CallID", v.CallID, "call_1")
		wantField(t, "V2SessionMessage", "shell", "Command", v.Command, "ls -la")
		wantField(t, "V2SessionMessage", "shell", "Output", v.Output, "total 0")
		st, ok := v.Time.(V2SessionMessageShellTime)
		if !ok {
			t.Fatalf("Time runtime type = %T, want V2SessionMessageShellTime", v.Time)
		}
		wantField(t, "V2SessionMessage", "shell", "Time.Created", st.Created, int64(16))
		wantField(t, "V2SessionMessage", "shell", "Time.Completed", st.Completed, int64(17))
	})

	t.Run("assistant", func(t *testing.T) {
		var v V2SessionMessage
		raw := `{"id":"msg_7","type":"assistant","time":{"created":18,"completed":19},"agent":"build",
			"model":{"id":"gpt","providerID":"openai"},
			"content":[{"type":"text","id":"c1","text":"answer"}],
			"snapshot":{"start":"s","end":"e","files":["f1"]},
			"finish":"stop","cost":1.25,
			"tokens":{"input":1,"output":2,"reasoning":3,"cache":{"read":4,"write":5}},
			"error":{"type":"unknown","message":"boom"},"metadata":{}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "V2SessionMessage", "assistant", "ID", v.ID, "msg_7")
		wantField(t, "V2SessionMessage", "assistant", "Type", v.Type, "assistant")
		wantField(t, "V2SessionMessage", "assistant", "Agent", v.Agent, "build")
		wantField(t, "V2SessionMessage", "assistant", "Finish", v.Finish, "stop")
		wantField(t, "V2SessionMessage", "assistant", "Cost", v.Cost, 1.25)
		at, ok := v.Time.(V2SessionMessageAssistantTime)
		if !ok {
			t.Fatalf("Time runtime type = %T, want V2SessionMessageAssistantTime", v.Time)
		}
		wantField(t, "V2SessionMessage", "assistant", "Time.Completed", at.Completed, int64(19))
		content, ok := v.Content.([]V2SessionMessageAssistantContent)
		if !ok {
			t.Fatalf("Content runtime type = %T, want []V2SessionMessageAssistantContent", v.Content)
		}
		if len(content) != 1 {
			t.Fatalf("Content len = %d, want 1", len(content))
		}
		// nested carrier promotion must work too
		wantField(t, "V2SessionMessage", "assistant", "Content[0].ID", content[0].ID, "c1")
		wantField(t, "V2SessionMessage", "assistant", "Content[0].Type", content[0].Type, "text")
		wantField(t, "V2SessionMessage", "assistant", "Content[0].Text", content[0].Text, "answer")
		snap, ok := v.Snapshot.(V2SessionMessageAssistantSnapshot)
		if !ok {
			t.Fatalf("Snapshot runtime type = %T, want V2SessionMessageAssistantSnapshot", v.Snapshot)
		}
		wantField(t, "V2SessionMessage", "assistant", "Snapshot.Start", snap.Start, "s")
		wantField(t, "V2SessionMessage", "assistant", "Snapshot.Files", snap.Files, []string{"f1"})
		tok, ok := v.Tokens.(V2SessionMessageTokens)
		if !ok {
			t.Fatalf("Tokens runtime type = %T, want V2SessionMessageTokens", v.Tokens)
		}
		wantField(t, "V2SessionMessage", "assistant", "Tokens.Input", tok.Input, int64(1))
		wantField(t, "V2SessionMessage", "assistant", "Tokens.Cache.Write", tok.Cache.Write, int64(5))
		errv, ok := v.Error.(SessionErrorUnknown)
		if !ok {
			t.Fatalf("Error runtime type = %T, want SessionErrorUnknown", v.Error)
		}
		wantField(t, "V2SessionMessage", "assistant", "Error.Message", errv.Message, "boom")
		wantField(t, "V2SessionMessage", "assistant", "Error.Type", errv.Type, SessionErrorUnknownTypeUnknown)
		wantZero(t, "V2SessionMessage", "assistant", "Reason", v.Reason)
		wantZero(t, "V2SessionMessage", "assistant", "Summary", v.Summary)
	})

	t.Run("compaction", func(t *testing.T) {
		var v V2SessionMessage
		raw := `{"id":"msg_8","type":"compaction","time":{"created":20},"reason":"auto","summary":"sum","recent":"rec","metadata":{}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "V2SessionMessage", "compaction", "ID", v.ID, "msg_8")
		wantField(t, "V2SessionMessage", "compaction", "Type", v.Type, "compaction")
		wantField(t, "V2SessionMessage", "compaction", "Reason", v.Reason, V2SessionMessageCompactionReasonAuto)
		if !v.Reason.IsKnown() {
			t.Errorf("Reason %q should be known", v.Reason)
		}
		wantField(t, "V2SessionMessage", "compaction", "Summary", v.Summary, "sum")
		wantField(t, "V2SessionMessage", "compaction", "Recent", v.Recent, "rec")
		wantZero(t, "V2SessionMessage", "compaction", "Cost", v.Cost)
		wantZero(t, "V2SessionMessage", "compaction", "Content", v.Content)
	})
}

// ===== V2SessionMessageAssistantContent (3 variants, 8 promoted fields) =====

func TestV2SessionMessageAssistantContentCarrierPromotedFields(t *testing.T) {
	t.Run("text", func(t *testing.T) {
		var v V2SessionMessageAssistantContent
		raw := `{"type":"text","id":"c1","text":"hello"}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.AsUnion().(V2SessionMessageAssistantTextContent); !ok {
			t.Fatalf("union = %T", v.AsUnion())
		}
		wantField(t, "AssistantContent", "text", "ID", v.ID, "c1")
		wantField(t, "AssistantContent", "text", "Type", v.Type, "text")
		wantField(t, "AssistantContent", "text", "Text", v.Text, "hello")
		wantZero(t, "AssistantContent", "text", "Name", v.Name)
		wantZero(t, "AssistantContent", "text", "State", v.State)
		wantZero(t, "AssistantContent", "text", "Time", v.Time)
		wantZero(t, "AssistantContent", "text", "Provider", v.Provider)
		wantZero(t, "AssistantContent", "text", "ProviderMetadata", v.ProviderMetadata)
	})

	t.Run("reasoning", func(t *testing.T) {
		var v V2SessionMessageAssistantContent
		raw := `{"type":"reasoning","id":"c2","text":"think","providerMetadata":{"openai":{"effort":"high"}},"time":{"created":1,"completed":2}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.AsUnion().(V2SessionMessageAssistantReasoningContent); !ok {
			t.Fatalf("union = %T", v.AsUnion())
		}
		wantField(t, "AssistantContent", "reasoning", "ID", v.ID, "c2")
		wantField(t, "AssistantContent", "reasoning", "Type", v.Type, "reasoning")
		wantField(t, "AssistantContent", "reasoning", "Text", v.Text, "think")
		wantField(t, "AssistantContent", "reasoning", "ProviderMetadata", v.ProviderMetadata,
			map[string]any{"openai": map[string]any{"effort": "high"}})
		// The reasoning variant declares `time` as `any`, so Port transports the
		// generic decode result rather than a typed struct.
		wantField(t, "AssistantContent", "reasoning", "Time", v.Time,
			map[string]any{"created": 1.0, "completed": 2.0})
		wantZero(t, "AssistantContent", "reasoning", "Name", v.Name)
		wantZero(t, "AssistantContent", "reasoning", "State", v.State)
	})

	t.Run("tool", func(t *testing.T) {
		var v V2SessionMessageAssistantContent
		raw := `{"type":"tool","id":"c3","name":"bash","provider":{"executed":true,"metadata":{"m":1}},
			"state":{"status":"pending","input":"echo hi"},
			"time":{"created":1,"completed":2,"ran":3,"pruned":4}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.AsUnion().(V2SessionMessageAssistantToolContent); !ok {
			t.Fatalf("union = %T", v.AsUnion())
		}
		wantField(t, "AssistantContent", "tool", "ID", v.ID, "c3")
		wantField(t, "AssistantContent", "tool", "Type", v.Type, "tool")
		wantField(t, "AssistantContent", "tool", "Name", v.Name, "bash")
		prov, ok := v.Provider.(V2SessionMessageToolProvider)
		if !ok {
			t.Fatalf("Provider runtime type = %T, want V2SessionMessageToolProvider", v.Provider)
		}
		wantField(t, "AssistantContent", "tool", "Provider.Executed", prov.Executed, true)
		state, ok := v.State.(V2SessionMessageToolState)
		if !ok {
			t.Fatalf("State runtime type = %T, want V2SessionMessageToolState", v.State)
		}
		wantField(t, "AssistantContent", "tool", "State.Status", state.Status, V2SessionMessageToolStateStatusPending)
		wantField(t, "AssistantContent", "tool", "State.Input", state.Input, "echo hi")
		tt, ok := v.Time.(V2SessionMessageToolTime)
		if !ok {
			t.Fatalf("Time runtime type = %T, want V2SessionMessageToolTime", v.Time)
		}
		wantField(t, "AssistantContent", "tool", "Time.Created", tt.Created, int64(1))
		wantZero(t, "AssistantContent", "tool", "Text", v.Text)
		wantZero(t, "AssistantContent", "tool", "ProviderMetadata", v.ProviderMetadata)
	})
}

// ===== V2SessionMessageToolState (4 variants, 8 promoted fields) =====

func TestV2SessionMessageToolStateCarrierPromotedFields(t *testing.T) {
	t.Run("pending", func(t *testing.T) {
		var v V2SessionMessageToolState
		raw := `{"status":"pending","input":"raw-input"}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.AsUnion().(V2SessionMessageToolStatePending); !ok {
			t.Fatalf("union = %T", v.AsUnion())
		}
		wantField(t, "ToolState", "pending", "Status", v.Status, V2SessionMessageToolStateStatusPending)
		// pending declares `input` as a string; every other variant uses an object
		wantField(t, "ToolState", "pending", "Input", v.Input, "raw-input")
		wantZero(t, "ToolState", "pending", "Content", v.Content)
		wantZero(t, "ToolState", "pending", "Structured", v.Structured)
		wantZero(t, "ToolState", "pending", "Attachments", v.Attachments)
		wantZero(t, "ToolState", "pending", "OutputPaths", v.OutputPaths)
		wantZero(t, "ToolState", "pending", "Result", v.Result)
		wantZero(t, "ToolState", "pending", "Error", v.Error)
	})

	t.Run("running", func(t *testing.T) {
		var v V2SessionMessageToolState
		raw := `{"status":"running","input":{"cmd":"ls"},"structured":{"pid":7},"content":[{"type":"text","text":"partial"}]}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "ToolState", "running", "Status", v.Status, V2SessionMessageToolStateStatusRunning)
		wantField(t, "ToolState", "running", "Input", v.Input, map[string]any{"cmd": "ls"})
		wantField(t, "ToolState", "running", "Structured", v.Structured, map[string]any{"pid": 7.0})
		wantField(t, "ToolState", "running", "Content", v.Content,
			[]any{map[string]any{"type": "text", "text": "partial"}})
		wantZero(t, "ToolState", "running", "Attachments", v.Attachments)
		wantZero(t, "ToolState", "running", "Error", v.Error)
	})

	t.Run("completed", func(t *testing.T) {
		var v V2SessionMessageToolState
		raw := `{"status":"completed","input":{"cmd":"ls"},"structured":{"code":0},
			"content":[{"type":"file","uri":"file:///o.txt","mime":"text/plain"}],
			"attachments":[{"uri":"file:///a.txt","mime":"text/plain"}],
			"outputPaths":["/o.txt"],"result":{"ok":true}}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "ToolState", "completed", "Status", v.Status, V2SessionMessageToolStateStatusCompleted)
		wantField(t, "ToolState", "completed", "Input", v.Input, map[string]any{"cmd": "ls"})
		wantField(t, "ToolState", "completed", "Structured", v.Structured, map[string]any{"code": 0.0})
		attach, ok := v.Attachments.([]V2PromptFileAttachment)
		if !ok {
			t.Fatalf("Attachments runtime type = %T, want []V2PromptFileAttachment", v.Attachments)
		}
		if len(attach) != 1 || attach[0].Mime != "text/plain" || attach[0].URI != "file:///a.txt" {
			t.Errorf("Attachments = %#v", attach)
		}
		wantField(t, "ToolState", "completed", "OutputPaths", v.OutputPaths, []any{"/o.txt"})
		wantField(t, "ToolState", "completed", "Result", v.Result, map[string]any{"ok": true})
		wantZero(t, "ToolState", "completed", "Error", v.Error)
	})

	t.Run("error", func(t *testing.T) {
		var v V2SessionMessageToolState
		raw := `{"status":"error","input":{"cmd":"ls"},"structured":{},
			"content":[{"type":"text","text":"stderr"}],
			"error":{"type":"unknown","message":"exit 1"},"result":42}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "ToolState", "error", "Status", v.Status, V2SessionMessageToolStateStatusError)
		wantField(t, "ToolState", "error", "Input", v.Input, map[string]any{"cmd": "ls"})
		errv, ok := v.Error.(SessionErrorUnknown)
		if !ok {
			t.Fatalf("Error runtime type = %T, want SessionErrorUnknown", v.Error)
		}
		wantField(t, "ToolState", "error", "Error.Message", errv.Message, "exit 1")
		// `result` is an unconstrained OpenAPI schema, so a bare number is valid
		wantField(t, "ToolState", "error", "Result", v.Result, 42.0)
		wantZero(t, "ToolState", "error", "Attachments", v.Attachments)
		wantZero(t, "ToolState", "error", "OutputPaths", v.OutputPaths)
	})

	t.Run("all statuses known", func(t *testing.T) {
		for _, s := range []V2SessionMessageToolStateStatus{
			V2SessionMessageToolStateStatusPending,
			V2SessionMessageToolStateStatusRunning,
			V2SessionMessageToolStateStatusCompleted,
			V2SessionMessageToolStateStatusError,
		} {
			if !s.IsKnown() {
				t.Errorf("status %q should be known", s)
			}
		}
	})
}

// ===== V2SessionDurableEvent (28 variants, 6 promoted fields) =====

// v2DurableEventTypeToVariant maps every OpenAPI SessionDurableEvent
// discriminator value to the Go variant it must route to.
var v2DurableEventTypeToVariant = []struct {
	typ     string
	variant any
	data    any
}{
	{"session.next.agent.switched", V2SessionDurableEventAgentSwitched{}, V2EventSessionNextAgentSwitchedData{}},
	{"session.next.model.switched", V2SessionDurableEventModelSwitched{}, V2EventSessionNextModelSwitchedData{}},
	{"session.next.moved", V2SessionDurableEventMoved{}, V2EventSessionNextMovedData{}},
	{"session.next.prompted", V2SessionDurableEventPrompted{}, V2EventSessionNextPromptedData{}},
	{"session.next.prompt.admitted", V2SessionDurableEventPromptAdmitted{}, V2EventSessionNextPromptAdmittedData{}},
	{"session.next.context.updated", V2SessionDurableEventContextUpdated{}, V2EventSessionNextContextUpdatedData{}},
	{"session.next.synthetic", V2SessionDurableEventSynthetic{}, V2EventSessionNextSyntheticData{}},
	{"session.next.shell.started", V2SessionDurableEventShellStarted{}, V2EventSessionNextShellStartedData{}},
	{"session.next.shell.ended", V2SessionDurableEventShellEnded{}, V2EventSessionNextShellEndedData{}},
	{"session.next.step.started", V2SessionDurableEventStepStarted{}, V2EventSessionNextStepStartedData{}},
	{"session.next.step.ended", V2SessionDurableEventStepEnded{}, V2EventSessionNextStepEndedData{}},
	{"session.next.step.failed", V2SessionDurableEventStepFailed{}, V2EventSessionNextStepFailedData{}},
	{"session.next.text.started", V2SessionDurableEventTextStarted{}, V2EventSessionNextTextStartedData{}},
	{"session.next.text.ended", V2SessionDurableEventTextEnded{}, V2EventSessionNextTextEndedData{}},
	{"session.next.tool.input.started", V2SessionDurableEventToolInputStarted{}, V2EventSessionNextToolInputStartedData{}},
	{"session.next.tool.input.ended", V2SessionDurableEventToolInputEnded{}, V2EventSessionNextToolInputEndedData{}},
	{"session.next.tool.called", V2SessionDurableEventToolCalled{}, V2EventSessionNextToolCalledData{}},
	{"session.next.tool.progress", V2SessionDurableEventToolProgress{}, V2EventSessionNextToolProgressData{}},
	{"session.next.tool.success", V2SessionDurableEventToolSuccess{}, V2EventSessionNextToolSuccessData{}},
	{"session.next.tool.failed", V2SessionDurableEventToolFailed{}, V2EventSessionNextToolFailedData{}},
	{"session.next.reasoning.started", V2SessionDurableEventReasoningStarted{}, V2EventSessionNextReasoningStartedData{}},
	{"session.next.reasoning.ended", V2SessionDurableEventReasoningEnded{}, V2EventSessionNextReasoningEndedData{}},
	{"session.next.retried", V2SessionDurableEventRetried{}, V2EventSessionNextRetriedData{}},
	{"session.next.compaction.started", V2SessionDurableEventCompactionStarted{}, V2EventSessionNextCompactionStartedData{}},
	{"session.next.compaction.ended", V2SessionDurableEventCompactionEnded{}, V2EventSessionNextCompactionEndedData{}},
	{"session.next.revert.staged", V2SessionDurableEventRevertStaged{}, V2EventSessionNextRevertStagedData{}},
	{"session.next.revert.cleared", V2SessionDurableEventRevertCleared{}, V2EventSessionNextRevertClearedData{}},
	{"session.next.revert.committed", V2SessionDurableEventRevertCommitted{}, V2EventSessionNextRevertCommittedData{}},
}

func TestV2SessionDurableEventCarrierPromotedFields(t *testing.T) {
	if len(v2DurableEventTypeToVariant) != 28 {
		t.Fatalf("expected 28 durable event variants, got %d", len(v2DurableEventTypeToVariant))
	}
	for i, tc := range v2DurableEventTypeToVariant {
		t.Run(tc.typ, func(t *testing.T) {
			raw := fmt.Sprintf(`{"id":"evt_%d","type":%q,
				"durable":{"aggregateID":"agg_1","seq":%d,"version":3},
				"location":{"directory":"/work","workspaceID":"wrk_1"},
				"metadata":{"trace":"t1"},
				"data":{"sessionID":"ses_1","timestamp":%d}}`, i, tc.typ, i, 1000+i)
			var v V2SessionDurableEvent
			if err := json.Unmarshal([]byte(raw), &v); err != nil {
				t.Fatal(err)
			}
			if got, want := reflect.TypeOf(v.AsUnion()), reflect.TypeOf(tc.variant); got != want {
				t.Fatalf("union = %v, want %v", got, want)
			}
			wantField(t, "V2SessionDurableEvent", tc.typ, "ID", v.ID, fmt.Sprintf("evt_%d", i))
			wantField(t, "V2SessionDurableEvent", tc.typ, "Type", v.Type, tc.typ)
			d, ok := v.Durable.(V2EventDurable)
			if !ok {
				t.Fatalf("Durable runtime type = %T, want V2EventDurable", v.Durable)
			}
			wantField(t, "V2SessionDurableEvent", tc.typ, "Durable.AggregateID", d.AggregateID, "agg_1")
			wantField(t, "V2SessionDurableEvent", tc.typ, "Durable.Seq", d.Seq, int64(i))
			wantField(t, "V2SessionDurableEvent", tc.typ, "Durable.Version", d.Version, int64(3))
			loc, ok := v.Location.(LocationRef)
			if !ok {
				t.Fatalf("Location runtime type = %T, want LocationRef", v.Location)
			}
			wantField(t, "V2SessionDurableEvent", tc.typ, "Location.Directory", loc.Directory, "/work")
			wantField(t, "V2SessionDurableEvent", tc.typ, "Location.WorkspaceID", loc.WorkspaceID, "wrk_1")
			wantField(t, "V2SessionDurableEvent", tc.typ, "Metadata", v.Metadata, map[string]any{"trace": "t1"})
			if got, want := reflect.TypeOf(v.Data), reflect.TypeOf(tc.data); got != want {
				t.Errorf("Data runtime type = %v, want %v", got, want)
			}
		})
	}
}

// ===== reachability regression: all 43 variants across the 4 carriers =====

func TestV2SessionCarrierVariantReachability(t *testing.T) {
	seen := map[string]bool{}

	messages := []string{
		`{"id":"a","type":"agent-switched","time":{"created":1},"agent":"x"}`,
		`{"id":"a","type":"model-switched","time":{"created":1},"model":{"id":"m","providerID":"p"}}`,
		`{"id":"a","type":"user","time":{"created":1},"text":"t"}`,
		`{"id":"a","type":"synthetic","time":{"created":1},"sessionID":"s","text":"t"}`,
		`{"id":"a","type":"system","time":{"created":1},"text":"t"}`,
		`{"id":"a","type":"shell","time":{"created":1},"callID":"c","command":"c","output":"o"}`,
		`{"id":"a","type":"assistant","time":{"created":1},"agent":"x","model":{"id":"m","providerID":"p"},"content":[]}`,
		`{"id":"a","type":"compaction","time":{"created":1},"reason":"manual","summary":"s","recent":"r"}`,
	}
	for _, raw := range messages {
		var v V2SessionMessage
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		seen[fmt.Sprintf("%T", v.AsUnion())] = true
	}
	if got := len(seen); got != 8 {
		t.Errorf("V2SessionMessage reachable variants = %d, want 8", got)
	}

	contents := []string{
		`{"type":"text","id":"c","text":"t"}`,
		`{"type":"reasoning","id":"c","text":"t"}`,
		`{"type":"tool","id":"c","name":"n","state":{"status":"pending","input":"i"},"time":{"created":1}}`,
	}
	before := len(seen)
	for _, raw := range contents {
		var v V2SessionMessageAssistantContent
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		seen[fmt.Sprintf("%T", v.AsUnion())] = true
	}
	if got := len(seen) - before; got != 3 {
		t.Errorf("AssistantContent reachable variants = %d, want 3", got)
	}

	states := []string{
		`{"status":"pending","input":"i"}`,
		`{"status":"running","input":{},"structured":{},"content":[]}`,
		`{"status":"completed","input":{},"structured":{},"content":[]}`,
		`{"status":"error","input":{},"structured":{},"content":[],"error":{"type":"unknown","message":"m"}}`,
	}
	before = len(seen)
	for _, raw := range states {
		var v V2SessionMessageToolState
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		seen[fmt.Sprintf("%T", v.AsUnion())] = true
	}
	if got := len(seen) - before; got != 4 {
		t.Errorf("ToolState reachable variants = %d, want 4", got)
	}

	before = len(seen)
	for _, tc := range v2DurableEventTypeToVariant {
		raw := fmt.Sprintf(`{"id":"e","type":%q,"data":{"sessionID":"s","timestamp":1}}`, tc.typ)
		var v V2SessionDurableEvent
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		seen[fmt.Sprintf("%T", v.AsUnion())] = true
	}
	if got := len(seen) - before; got != 28 {
		t.Errorf("V2SessionDurableEvent reachable variants = %d, want 28", got)
	}

	if got := len(seen); got != 43 {
		t.Errorf("total distinct reachable variants = %d, want 43", got)
	}
}

// ===== unknown extra fields must not break routing or promotion =====

func TestV2SessionCarrierUnknownFieldsTolerated(t *testing.T) {
	t.Run("V2SessionMessage", func(t *testing.T) {
		var v V2SessionMessage
		raw := `{"id":"msg_x","type":"shell","time":{"created":5,"completed":6},"callID":"c","command":"cmd","output":"o","futureField":{"deep":[1,2]},"anotherNew":"x"}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.AsUnion().(V2SessionMessageShell); !ok {
			t.Fatalf("union = %T, want V2SessionMessageShell", v.AsUnion())
		}
		wantField(t, "V2SessionMessage", "unknown-fields", "ID", v.ID, "msg_x")
		wantField(t, "V2SessionMessage", "unknown-fields", "Command", v.Command, "cmd")
		wantField(t, "V2SessionMessage", "unknown-fields", "Output", v.Output, "o")
		if _, ok := v.JSON.ExtraFields["futureField"]; !ok {
			t.Errorf("ExtraFields missing futureField, got %v", v.JSON.ExtraFields)
		}
		if _, ok := v.JSON.ExtraFields["anotherNew"]; !ok {
			t.Errorf("ExtraFields missing anotherNew, got %v", v.JSON.ExtraFields)
		}
	})

	t.Run("V2SessionDurableEvent", func(t *testing.T) {
		var v V2SessionDurableEvent
		raw := `{"id":"evt_x","type":"session.next.tool.success","durable":{"aggregateID":"a","seq":9,"version":1},"location":{"directory":"/d"},"data":{"sessionID":"s","timestamp":1},"brandNew":true}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.AsUnion().(V2SessionDurableEventToolSuccess); !ok {
			t.Fatalf("union = %T", v.AsUnion())
		}
		wantField(t, "V2SessionDurableEvent", "unknown-fields", "ID", v.ID, "evt_x")
		wantField(t, "V2SessionDurableEvent", "unknown-fields", "Type", v.Type, "session.next.tool.success")
		if d, ok := v.Durable.(V2EventDurable); !ok || d.Seq != 9 {
			t.Errorf("Durable = %#v", v.Durable)
		}
		if _, ok := v.JSON.ExtraFields["brandNew"]; !ok {
			t.Errorf("ExtraFields missing brandNew, got %v", v.JSON.ExtraFields)
		}
	})

	t.Run("V2SessionMessageToolState", func(t *testing.T) {
		var v V2SessionMessageToolState
		raw := `{"status":"completed","input":{"a":1},"structured":{},"content":[],"unexpected":"keep"}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "ToolState", "unknown-fields", "Status", v.Status, V2SessionMessageToolStateStatusCompleted)
		wantField(t, "ToolState", "unknown-fields", "Input", v.Input, map[string]any{"a": 1.0})
		if _, ok := v.JSON.ExtraFields["unexpected"]; !ok {
			t.Errorf("ExtraFields missing unexpected, got %v", v.JSON.ExtraFields)
		}
	})

	t.Run("V2SessionMessageAssistantContent", func(t *testing.T) {
		var v V2SessionMessageAssistantContent
		raw := `{"type":"text","id":"c9","text":"t","surprise":[1]}`
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		wantField(t, "AssistantContent", "unknown-fields", "ID", v.ID, "c9")
		wantField(t, "AssistantContent", "unknown-fields", "Text", v.Text, "t")
		if _, ok := v.JSON.ExtraFields["surprise"]; !ok {
			t.Errorf("ExtraFields missing surprise, got %v", v.JSON.ExtraFields)
		}
	})
}

// ===== runtime-comment honesty for generically decoded fields =====

// A handful of variant structs declare their own field as `any`, so
// [apijson.Port] transports the generic decode result into the carrier rather
// than a typed struct. Their runtime comments name the *re-decodable* typed form,
// exactly like [Config.Lsp] (`[bool], [map[string]ConfigLsp]`) whose object form
// also arrives as `map[string]any`. This test pins both halves of that promise:
// the generic runtime type, and that the raw sub-document re-decodes into the
// documented type.
func TestV2SessionCarrierGenericFieldsAreRedecodable(t *testing.T) {
	t.Run("AssistantContent.Time reasoning", func(t *testing.T) {
		raw := `{"type":"reasoning","id":"c1","text":"t","time":{"created":7,"completed":9}}`
		var v V2SessionMessageAssistantContent
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.Time.(map[string]any); !ok {
			t.Fatalf("Time runtime type = %T, want map[string]any", v.Time)
		}
		sub, err := json.Marshal(v.Time)
		if err != nil {
			t.Fatal(err)
		}
		var typed V2SessionMessageAssistantReasoningContentTime
		if err := json.Unmarshal(sub, &typed); err != nil {
			t.Fatalf("re-decode: %v", err)
		}
		wantField(t, "AssistantContent", "reasoning", "Time.Created", typed.Created, int64(7))
		wantField(t, "AssistantContent", "reasoning", "Time.Completed", typed.Completed, int64(9))
	})

	t.Run("ToolState.Content text", func(t *testing.T) {
		raw := `{"status":"running","input":{},"structured":{},"content":[{"type":"text","text":"hello"}]}`
		var v V2SessionMessageToolState
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.Content.([]any); !ok {
			t.Fatalf("Content runtime type = %T, want []any", v.Content)
		}
		sub, err := json.Marshal(v.Content)
		if err != nil {
			t.Fatal(err)
		}
		var typed []ToolTextContent
		if err := json.Unmarshal(sub, &typed); err != nil {
			t.Fatalf("re-decode: %v", err)
		}
		if len(typed) != 1 || typed[0].Type != "text" || typed[0].Text != "hello" {
			t.Errorf("re-decoded content = %#v", typed)
		}
	})

	t.Run("ToolState.Content file", func(t *testing.T) {
		raw := `{"status":"running","input":{},"structured":{},"content":[{"type":"file","uri":"file:///o.txt","mime":"text/plain","name":"o.txt"}]}`
		var v V2SessionMessageToolState
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		sub, err := json.Marshal(v.Content)
		if err != nil {
			t.Fatal(err)
		}
		var typed []ToolFileContent
		if err := json.Unmarshal(sub, &typed); err != nil {
			t.Fatalf("re-decode: %v", err)
		}
		if len(typed) != 1 || typed[0].URI != "file:///o.txt" || typed[0].Mime != "text/plain" {
			t.Errorf("re-decoded content = %#v", typed)
		}
	})

	t.Run("ToolState.OutputPaths", func(t *testing.T) {
		raw := `{"status":"completed","input":{},"structured":{},"content":[],"outputPaths":["/a.txt","/b.txt"]}`
		var v V2SessionMessageToolState
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if _, ok := v.OutputPaths.([]any); !ok {
			t.Fatalf("OutputPaths runtime type = %T, want []any", v.OutputPaths)
		}
		sub, err := json.Marshal(v.OutputPaths)
		if err != nil {
			t.Fatal(err)
		}
		var typed []string
		if err := json.Unmarshal(sub, &typed); err != nil {
			t.Fatalf("re-decode: %v", err)
		}
		wantField(t, "ToolState", "completed", "OutputPaths", typed, []string{"/a.txt", "/b.txt"})
	})
}

// ===== RawJSON fidelity =====

func TestV2SessionCarrierRawJSONFidelity(t *testing.T) {
	t.Run("V2SessionMessage", func(t *testing.T) {
		raw := `{"id":"msg_r","type":"user","time":{"created":1},"text":"hi"}`
		var v V2SessionMessage
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if v.JSON.RawJSON() != raw {
			t.Errorf("RawJSON()\n got = %s\nwant = %s", v.JSON.RawJSON(), raw)
		}
	})

	t.Run("V2SessionMessageAssistantContent", func(t *testing.T) {
		raw := `{"type":"text","id":"c_r","text":"x"}`
		var v V2SessionMessageAssistantContent
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if v.JSON.RawJSON() != raw {
			t.Errorf("RawJSON()\n got = %s\nwant = %s", v.JSON.RawJSON(), raw)
		}
	})

	t.Run("V2SessionMessageToolState", func(t *testing.T) {
		raw := `{"status":"pending","input":"i"}`
		var v V2SessionMessageToolState
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if v.JSON.RawJSON() != raw {
			t.Errorf("RawJSON()\n got = %s\nwant = %s", v.JSON.RawJSON(), raw)
		}
	})

	t.Run("V2SessionDurableEvent", func(t *testing.T) {
		raw := `{"id":"e_r","type":"session.next.moved","data":{"sessionID":"s","timestamp":1}}`
		var v V2SessionDurableEvent
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if v.JSON.RawJSON() != raw {
			t.Errorf("RawJSON()\n got = %s\nwant = %s", v.JSON.RawJSON(), raw)
		}
	})
}

// ===== carrier JSON metadata must cover every promoted field =====

func TestV2SessionCarrierJSONMetadataCoverage(t *testing.T) {
	cases := []struct {
		name     string
		carrier  any
		metadata any
	}{
		{"V2SessionMessage", V2SessionMessage{}, v2SessionMessageJSON{}},
		{"V2SessionMessageAssistantContent", V2SessionMessageAssistantContent{}, v2SessionMessageAssistantContentJSON{}},
		{"V2SessionMessageToolState", V2SessionMessageToolState{}, v2SessionMessageToolStateJSON{}},
		{"V2SessionDurableEvent", V2SessionDurableEvent{}, v2SessionDurableEventJSON{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ct := reflect.TypeOf(tc.carrier)
			mt := reflect.TypeOf(tc.metadata)
			promoted := 0
			for i := range ct.NumField() {
				f := ct.Field(i)
				if f.Name == "JSON" || !f.IsExported() {
					continue
				}
				promoted++
				mf, ok := mt.FieldByName(f.Name)
				if !ok {
					t.Errorf("%s: promoted field %s has no apijson.Field in %s", tc.name, f.Name, mt.Name())
					continue
				}
				if mf.Type != reflect.TypeFor[apijson.Field]() {
					t.Errorf("%s: %s.%s type = %v, want apijson.Field", tc.name, mt.Name(), f.Name, mf.Type)
				}
			}
			if promoted == 0 {
				t.Errorf("%s: no promoted fields found - Port would be a no-op", tc.name)
			}
			if _, ok := mt.FieldByName("ExtraFields"); !ok {
				t.Errorf("%s: %s missing ExtraFields", tc.name, mt.Name())
			}
			if _, ok := mt.FieldByName("raw"); !ok {
				t.Errorf("%s: %s missing raw", tc.name, mt.Name())
			}
		})
	}
}
