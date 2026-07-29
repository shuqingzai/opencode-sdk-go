package opencode

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/sst/opencode-sdk-go/shared"
)

// ─────────────────────────────────────────────────────────────────────────────
// FilePartSource Union (FileSource / SymbolSource / ResourceSource)
// Aligned with OpenAPI FilePart.source anyOf schema.
// ─────────────────────────────────────────────────────────────────────────────

// TestFilePartSourceUnionFileVariant verifies FilePartSource.AsUnion() returns
// FileSource when type="file".
func TestFilePartSourceUnionFileVariant(t *testing.T) {
	raw := `{
		"type": "file",
		"path": "/src/main.go",
		"text": {"value": "package main", "start": 0, "end": 12}
	}`
	var src FilePartSource
	if err := json.Unmarshal([]byte(raw), &src); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if src.Type != FilePartSourceTypeFile {
		t.Errorf("Type = %q, want file", src.Type)
	}
	if src.Path != "/src/main.go" {
		t.Errorf("Path = %q, want /src/main.go", src.Path)
	}
	u := src.AsUnion()
	if u == nil {
		t.Fatal("AsUnion() returned nil")
	}
	fs, ok := u.(FileSource)
	if !ok {
		t.Fatalf("AsUnion() = %T, want FileSource", u)
	}
	if fs.Path != "/src/main.go" {
		t.Errorf("FileSource.Path = %q, want /src/main.go", fs.Path)
	}
	if fs.JSON.RawJSON() == "" {
		t.Error("FileSource.JSON.RawJSON() empty")
	}
	if src.JSON.RawJSON() == "" {
		t.Error("FilePartSource.JSON.RawJSON() empty")
	}
}

// TestFilePartSourceUnionSymbolVariant verifies FilePartSource.AsUnion() returns
// SymbolSource when type="symbol".
func TestFilePartSourceUnionSymbolVariant(t *testing.T) {
	raw := `{
		"type": "symbol",
		"path": "/src/main.go",
		"name": "main",
		"kind": 12,
		"text": {"value": "func main()", "start": 0, "end": 11},
		"range": {
			"start": {"line": 1, "character": 0},
			"end":   {"line": 5, "character": 1}
		}
	}`
	var src FilePartSource
	if err := json.Unmarshal([]byte(raw), &src); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if src.Type != FilePartSourceTypeSymbol {
		t.Errorf("Type = %q, want symbol", src.Type)
	}
	if src.Name != "main" {
		t.Errorf("Name = %q, want main", src.Name)
	}
	if src.Kind != 12 {
		t.Errorf("Kind = %d, want 12", src.Kind)
	}
	u := src.AsUnion()
	sym, ok := u.(SymbolSource)
	if !ok {
		t.Fatalf("AsUnion() = %T, want SymbolSource", u)
	}
	if sym.Name != "main" || sym.Kind != 12 {
		t.Errorf("SymbolSource = %+v", sym)
	}
	if sym.JSON.RawJSON() == "" {
		t.Error("SymbolSource.JSON.RawJSON() empty")
	}
}

// TestFilePartSourceUnionResourceVariant verifies FilePartSource.AsUnion() returns
// ResourceSource when type="resource".
func TestFilePartSourceUnionResourceVariant(t *testing.T) {
	raw := `{
		"type": "resource",
		"clientName": "my-mcp-server",
		"uri": "file:///docs/readme.md",
		"text": {"value": "# Docs", "start": 0, "end": 6}
	}`
	var src FilePartSource
	if err := json.Unmarshal([]byte(raw), &src); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if src.Type != FilePartSourceTypeResource {
		t.Errorf("Type = %q, want resource", src.Type)
	}
	u := src.AsUnion()
	rs, ok := u.(ResourceSource)
	if !ok {
		t.Fatalf("AsUnion() = %T, want ResourceSource", u)
	}
	if rs.ClientName != "my-mcp-server" {
		t.Errorf("ResourceSource.ClientName = %q", rs.ClientName)
	}
	if rs.URI != "file:///docs/readme.md" {
		t.Errorf("ResourceSource.URI = %q", rs.URI)
	}
	if rs.JSON.RawJSON() == "" {
		t.Error("ResourceSource.JSON.RawJSON() empty")
	}
}

// TestFilePartSourceTypeIsKnown verifies the type enum values.
func TestFilePartSourceTypeIsKnown(t *testing.T) {
	for _, v := range []FilePartSourceType{
		FilePartSourceTypeFile, FilePartSourceTypeSymbol, FilePartSourceTypeResource,
	} {
		if !v.IsKnown() {
			t.Errorf("%q.IsKnown() = false", v)
		}
	}
	if FilePartSourceType("unknown").IsKnown() {
		t.Error("unknown should not be known")
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// AssistantMessageError Union
// Aligned with OpenAPI AssistantMessage.error anyOf schema (8 error variants).
// ─────────────────────────────────────────────────────────────────────────────

// TestAssistantMessageErrorUnionProviderAuth verifies ProviderAuthError routing.
func TestAssistantMessageErrorUnionProviderAuth(t *testing.T) {
	raw := `{
		"name": "ProviderAuthError",
		"data": {"message": "Invalid API key", "providerID": "anthropic"}
	}`
	var e AssistantMessageError
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if e.Name != AssistantMessageErrorNameProviderAuthError {
		t.Errorf("Name = %q, want ProviderAuthError", e.Name)
	}
	u := e.AsUnion()
	if u == nil {
		t.Fatal("AsUnion() returned nil")
	}
	pae, ok := u.(shared.ProviderAuthError)
	if !ok {
		t.Fatalf("AsUnion() = %T, want shared.ProviderAuthError", u)
	}
	if pae.Data.ProviderID != "anthropic" {
		t.Errorf("ProviderAuthError.Data.ProviderID = %q", pae.Data.ProviderID)
	}
	if pae.Data.Message != "Invalid API key" {
		t.Errorf("ProviderAuthError.Data.Message = %q", pae.Data.Message)
	}
	if pae.JSON.RawJSON() == "" {
		t.Error("ProviderAuthError.JSON.RawJSON() empty")
	}
	if e.JSON.RawJSON() == "" {
		t.Error("AssistantMessageError.JSON.RawJSON() empty")
	}
}

// TestAssistantMessageErrorUnionContextOverflow verifies ContextOverflowError routing.
func TestAssistantMessageErrorUnionContextOverflow(t *testing.T) {
	raw := `{
		"name": "ContextOverflowError",
		"data": {}
	}`
	var e AssistantMessageError
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if e.Name != AssistantMessageErrorNameContextOverflowError {
		t.Errorf("Name = %q, want ContextOverflowError", e.Name)
	}
	u := e.AsUnion()
	if u == nil {
		t.Fatal("AsUnion() returned nil")
	}
	_, ok := u.(shared.ContextOverflowError)
	if !ok {
		t.Fatalf("AsUnion() = %T, want shared.ContextOverflowError", u)
	}
}

// TestAssistantMessageErrorUnionMessageAborted verifies MessageAbortedError routing.
func TestAssistantMessageErrorUnionMessageAborted(t *testing.T) {
	raw := `{
		"name": "MessageAbortedError",
		"data": {}
	}`
	var e AssistantMessageError
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	u := e.AsUnion()
	_, ok := u.(shared.MessageAbortedError)
	if !ok {
		t.Fatalf("AsUnion() = %T, want shared.MessageAbortedError", u)
	}
}

// TestAssistantMessageErrorNameIsKnown verifies all error name enum values.
func TestAssistantMessageErrorNameIsKnown(t *testing.T) {
	known := []AssistantMessageErrorName{
		AssistantMessageErrorNameProviderAuthError,
		AssistantMessageErrorNameUnknownError,
		AssistantMessageErrorNameMessageOutputLengthError,
		AssistantMessageErrorNameMessageAbortedError,
		AssistantMessageErrorNameStructuredOutputError,
		AssistantMessageErrorNameContextOverflowError,
		AssistantMessageErrorNameAPIError,
		AssistantMessageErrorNameContentFilterError,
	}
	for _, v := range known {
		if !v.IsKnown() {
			t.Errorf("%q.IsKnown() = false", v)
		}
	}
	if AssistantMessageErrorName("FakeError").IsKnown() {
		t.Error("FakeError should not be known")
	}
}

// TestAssistantMessageInContext verifies AssistantMessage.Error is deserialized
// correctly when embedded in a full AssistantMessage.
func TestAssistantMessageInContext(t *testing.T) {
	raw := `{
		"id": "msg_1",
		"sessionID": "ses_1",
		"role": "assistant",
		"parentID": "msg_0",
		"modelID": "claude-3-5-sonnet",
		"providerID": "anthropic",
		"mode": "default",
		"agent": "default",
		"path": {"cwd": "/tmp", "root": "/"},
		"cost": 0.001,
		"time": {"created": 1700000000},
		"tokens": {"input": 10, "output": 5, "reasoning": 0, "cache": {"read": 0, "write": 0}},
		"error": {
			"name": "ProviderAuthError",
			"data": {"message": "Unauthorized", "providerID": "anthropic"}
		}
	}`
	var m AssistantMessage
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.Error.Name != AssistantMessageErrorNameProviderAuthError {
		t.Errorf("Error.Name = %q, want ProviderAuthError", m.Error.Name)
	}
	u := m.Error.AsUnion()
	if u == nil {
		t.Fatal("Error.AsUnion() returned nil")
	}
	pae, ok := u.(shared.ProviderAuthError)
	if !ok {
		t.Fatalf("Error.AsUnion() = %T, want shared.ProviderAuthError", u)
	}
	if pae.Data.Message != "Unauthorized" {
		t.Errorf("ProviderAuthError.Data.Message = %q", pae.Data.Message)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Full union-routing regression matrix for session.go
// ─────────────────────────────────────────────────────────────────────────────

// TestSessionUnionRoutingMatrix asserts every union registered in session.go still
// routes every one of its OpenAPI `anyOf` variants to its own Go type. It is the
// regression net for changes to any single union's registration: adding a
// discriminator to one union (or converting a request-side variant to a
// response-side one) must not make any other variant unreachable.
//
// Variant lists are taken from the OpenAPI `anyOf` members of, respectively,
// `Message`, `Part`, `ToolState`, `FilePart.source`, `SessionStatus`,
// `AssistantMessage.error` and `UserMessage.format`.
func TestSessionUnionRoutingMatrix(t *testing.T) {
	const userBase = `"id":"msg_u","sessionID":"ses_1","role":"user","time":{"created":1},` +
		`"agent":"general","model":{"modelID":"m","providerID":"p"}`
	const asstBase = `"id":"msg_a","sessionID":"ses_1","role":"assistant","time":{"created":1},` +
		`"parentID":"","modelID":"m","providerID":"p","mode":"build","agent":"general",` +
		`"path":{"cwd":"/a","root":"/b"},"cost":0,` +
		`"tokens":{"input":1,"output":2,"reasoning":0,"cache":{"read":0,"write":0}}`
	const partBase = `"id":"prt_1","messageID":"msg_1","sessionID":"ses_1"`

	// MessageUnion — OpenAPI Message anyOf [UserMessage, AssistantMessage].
	t.Run("MessageUnion", func(t *testing.T) {
		cases := map[string]any{
			`{` + userBase + `}`: UserMessage{},
			`{` + asstBase + `}`: AssistantMessage{},
		}
		for raw, want := range cases {
			var m Message
			if err := json.Unmarshal([]byte(raw), &m); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			assertRuntimeType(t, m.AsUnion(), want)
			if m.JSON.RawJSON() != raw {
				t.Errorf("RawJSON() = %q, want %q", m.JSON.RawJSON(), raw)
			}
		}
	})

	// PartUnion — OpenAPI Part anyOf, 12 members. Every payload carries exactly the
	// properties its OpenAPI member declares as `required`; a payload that omits a
	// required property, or carries one the member does not declare, would be routed
	// by apijson's exactness heuristic rather than by the member's `type` enum and so
	// would not prove the variant is reachable for real server output.
	t.Run("PartUnion", func(t *testing.T) {
		cases := []struct {
			raw  string
			want any
		}{
			{`{` + partBase + `,"type":"text","text":"hi"}`, TextPart{}},
			{`{` + partBase + `,"type":"subtask","prompt":"p","description":"d","agent":"a"}`, SubtaskPart{}},
			{`{` + partBase + `,"type":"reasoning","text":"why","time":{"start":1}}`, ReasoningPart{}},
			{`{` + partBase + `,"type":"file","mime":"text/plain","url":"http://x"}`, FilePart{}},
			{`{` + partBase + `,"type":"tool","callID":"c1","tool":"bash","state":{"status":"pending","input":{},"raw":"{}"}}`, ToolPart{}},
			{`{` + partBase + `,"type":"step-start"}`, StepStartPart{}},
			{`{` + partBase + `,"type":"step-finish","reason":"stop","cost":0,"tokens":{"input":1,"output":1,"reasoning":0,"cache":{"read":0,"write":0}}}`, StepFinishPart{}},
			{`{` + partBase + `,"type":"snapshot","snapshot":"snap"}`, SnapshotPart{}},
			{`{` + partBase + `,"type":"patch","hash":"h","files":["a"]}`, PartPatchPart{}},
			{`{` + partBase + `,"type":"agent","name":"n"}`, AgentPart{}},
			{`{` + partBase + `,"type":"retry","attempt":1,"error":{"name":"APIError","data":{"message":"m","isRetryable":true}},"time":{"created":1}}`, PartRetryPart{}},
			{`{` + partBase + `,"type":"compaction","auto":true}`, CompactionPart{}},
		}
		seen := map[string]bool{}
		for _, tc := range cases {
			var p Part
			if err := json.Unmarshal([]byte(tc.raw), &p); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.raw, err)
			}
			assertRuntimeType(t, p.AsUnion(), tc.want)
			seen[fmt.Sprintf("%T", p.AsUnion())] = true

			// An unknown property drops every candidate to the same "extras" exactness
			// level, so routing must not depend on it.
			withExtra := tc.raw[:len(tc.raw)-1] + `,"unknownFutureField":"x"}`
			var q Part
			if err := json.Unmarshal([]byte(withExtra), &q); err != nil {
				t.Fatalf("unmarshal %s: %v", withExtra, err)
			}
			assertRuntimeType(t, q.AsUnion(), tc.want)
		}
		if len(seen) != 12 {
			t.Errorf("distinct PartUnion variants reached = %d, want 12 (%v)", len(seen), seen)
		}
	})

	// ToolPartStateUnion — OpenAPI ToolState anyOf, 4 members, keyed on `status`.
	t.Run("ToolPartStateUnion", func(t *testing.T) {
		cases := []struct {
			raw  string
			want any
		}{
			{`{"status":"pending","input":{},"raw":"{}"}`, ToolStatePending{}},
			{`{"status":"running","input":{},"time":{"start":1}}`, ToolStateRunning{}},
			{`{"status":"completed","input":{},"metadata":{},"output":"o","title":"t","time":{"start":1,"end":2}}`, ToolStateCompleted{}},
			{`{"status":"error","error":"boom","input":{},"time":{"start":1,"end":2}}`, ToolStateError{}},
		}
		seen := map[string]bool{}
		for _, tc := range cases {
			var s ToolPartState
			if err := json.Unmarshal([]byte(tc.raw), &s); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.raw, err)
			}
			assertRuntimeType(t, s.AsUnion(), tc.want)
			seen[fmt.Sprintf("%T", s.AsUnion())] = true
		}
		if len(seen) != 4 {
			t.Errorf("distinct ToolPartStateUnion variants reached = %d, want 4 (%v)", len(seen), seen)
		}
		// ToolStatePending.raw is required by OpenAPI and must survive onto the carrier.
		var pending ToolPartState
		if err := json.Unmarshal([]byte(`{"status":"pending","input":{"a":1},"raw":"{\"a\":1}"}`), &pending); err != nil {
			t.Fatalf("unmarshal pending: %v", err)
		}
		if pending.Raw != `{"a":1}` {
			t.Errorf("ToolPartState.Raw = %q, want %q", pending.Raw, `{"a":1}`)
		}
	})

	// FilePartSourceUnion — OpenAPI FilePart.source anyOf, 3 members. Note that
	// `path` belongs to FileSource/SymbolSource only, and `clientName`/`uri` to
	// ResourceSource only; `text` and `type` are the sole properties all three share.
	t.Run("FilePartSourceUnion", func(t *testing.T) {
		cases := []struct {
			raw  string
			want any
		}{
			{`{"type":"file","path":"/a","text":{"value":"v","start":0,"end":1}}`, FileSource{}},
			{`{"type":"symbol","path":"/a","name":"N","kind":1,"range":{"start":{"line":1,"character":0},"end":{"line":1,"character":2}},"text":{"value":"v","start":0,"end":1}}`, SymbolSource{}},
			{`{"type":"resource","clientName":"cn","uri":"res://x","text":{"value":"v","start":0,"end":1}}`, ResourceSource{}},
		}
		seen := map[string]bool{}
		for _, tc := range cases {
			var s FilePartSource
			if err := json.Unmarshal([]byte(tc.raw), &s); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.raw, err)
			}
			assertRuntimeType(t, s.AsUnion(), tc.want)
			seen[fmt.Sprintf("%T", s.AsUnion())] = true
		}
		if len(seen) != 3 {
			t.Errorf("distinct FilePartSourceUnion variants reached = %d, want 3 (%v)", len(seen), seen)
		}
		// ResourceSource's own properties must be ported onto the carrier too.
		var rs FilePartSource
		if err := json.Unmarshal([]byte(`{"type":"resource","clientName":"cn","uri":"res://x","text":{"value":"v","start":0,"end":1}}`), &rs); err != nil {
			t.Fatalf("unmarshal resource: %v", err)
		}
		if rs.ClientName != "cn" || rs.URI != "res://x" {
			t.Errorf("FilePartSource carrier ClientName/URI = %q/%q, want cn/res://x", rs.ClientName, rs.URI)
		}
	})

	// SessionStatus — discriminated on `type`; 3 members.
	t.Run("SessionStatus", func(t *testing.T) {
		cases := []struct {
			raw  string
			want any
		}{
			{`{"type":"idle"}`, SessionStatusIdle{}},
			{`{"type":"retry","attempt":1,"message":"m","next":1}`, SessionStatusRetry{}},
			{`{"type":"busy"}`, SessionStatusBusy{}},
		}
		for _, tc := range cases {
			var m SessionStatusMap
			if err := json.Unmarshal([]byte(`{"ses_1":`+tc.raw+`}`), &m); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.raw, err)
			}
			assertRuntimeType(t, m["ses_1"], tc.want)
		}
	})

	// AssistantMessageErrorUnion — OpenAPI AssistantMessage.error anyOf, 8 members,
	// all declared in package shared.
	t.Run("AssistantMessageErrorUnion", func(t *testing.T) {
		cases := []struct {
			raw  string
			want any
		}{
			{`{"name":"ProviderAuthError","data":{"providerID":"p","message":"m"}}`, shared.ProviderAuthError{}},
			{`{"name":"UnknownError","data":{"message":"m"}}`, shared.UnknownError{}},
			{`{"name":"MessageOutputLengthError","data":{}}`, shared.MessageOutputLengthError{}},
			{`{"name":"MessageAbortedError","data":{"message":"m"}}`, shared.MessageAbortedError{}},
			{`{"name":"StructuredOutputError","data":{"message":"m","retries":1}}`, shared.StructuredOutputError{}},
			{`{"name":"ContextOverflowError","data":{"message":"m"}}`, shared.ContextOverflowError{}},
			{`{"name":"ContentFilterError","data":{"message":"m"}}`, shared.ContentFilterError{}},
			{`{"name":"APIError","data":{"message":"m"}}`, shared.APIError{}},
		}
		seen := map[string]bool{}
		for _, tc := range cases {
			var e AssistantMessageError
			if err := json.Unmarshal([]byte(tc.raw), &e); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.raw, err)
			}
			assertRuntimeType(t, e.AsUnion(), tc.want)
			seen[fmt.Sprintf("%T", e.AsUnion())] = true
		}
		if len(seen) != 8 {
			t.Errorf("distinct AssistantMessageErrorUnion variants reached = %d, want 8 (%v)", len(seen), seen)
		}
	})

	// OutputFormatUnion — OpenAPI UserMessage.format anyOf, 2 members, discriminated
	// on `type`. Reached through the UserMessage carrier.
	t.Run("OutputFormatUnion", func(t *testing.T) {
		cases := []struct {
			raw  string
			want any
		}{
			{`{"type":"text"}`, OutputFormatText{}},
			{`{"type":"json_schema","schema":{"type":"object"}}`, OutputFormatJsonSchema{}},
		}
		for _, tc := range cases {
			var m UserMessage
			if err := json.Unmarshal([]byte(`{`+userBase+`,"format":`+tc.raw+`}`), &m); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.raw, err)
			}
			assertRuntimeType(t, m.AsFormat(), tc.want)
			assertRuntimeType(t, m.Format, tc.want)
		}
	})
}

// TestSessionRequestUnionMarshalling asserts the request-side unions in session.go
// still serialize correctly. They are param.Field-based and must produce the exact
// body OpenAPI declares for `POST /session/{sessionID}/message` and
// `POST /session/{sessionID}/prompt_async` (identical request schemas).
func TestSessionRequestUnionMarshalling(t *testing.T) {
	// SessionPromptParamsPartUnion — OpenAPI `parts` items anyOf, JS SDK v2
	// `Array<TextPartInput | FilePartInput | AgentPartInput | SubtaskPartInput>`.
	t.Run("SessionPromptParamsPartUnion", func(t *testing.T) {
		params := SessionPromptParams{
			Parts: F([]SessionPromptParamsPartUnion{
				TextPartInputParam{
					Type: F(TextPartInputTypeText),
					Text: F("hello"),
				},
				FilePartInputParam{
					Type: F(FilePartInputTypeFile),
					Mime: F("text/plain"),
					URL:  F("http://x/a.txt"),
				},
			}),
		}
		got, err := params.MarshalJSON()
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		for _, want := range []string{`"type":"text"`, `"text":"hello"`, `"type":"file"`, `"mime":"text/plain"`} {
			if !strings.Contains(string(got), want) {
				t.Errorf("body %s missing %s", got, want)
			}
		}
	})

	// SessionPromptParamsFormatUnion — OpenAPI `format` anyOf. This is the
	// request-side spelling of the same schema the response side models with
	// [OutputFormatUnion]; both must keep working independently.
	t.Run("SessionPromptParamsFormatUnion", func(t *testing.T) {
		cases := []struct {
			name  string
			value SessionPromptParamsFormatUnion
			want  []string
		}{
			{
				"text",
				SessionPromptParamsFormatText{Type: F(SessionPromptParamsFormatTextTypeText)},
				[]string{`"format":{"type":"text"}`},
			},
			{
				"json_schema",
				SessionPromptParamsFormatJsonSchema{
					Type:       F(SessionPromptParamsFormatJsonSchemaTypeJsonSchema),
					Schema:     F[any](map[string]any{"type": "object"}),
					RetryCount: F(int64(2)),
				},
				[]string{`"type":"json_schema"`, `"schema":{"type":"object"}`, `"retryCount":2`},
			},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				params := SessionPromptParams{
					Parts: F([]SessionPromptParamsPartUnion{
						TextPartInputParam{Type: F(TextPartInputTypeText), Text: F("x")},
					}),
					Format: F(tc.value),
				}
				got, err := params.MarshalJSON()
				if err != nil {
					t.Fatalf("marshal: %v", err)
				}
				for _, want := range tc.want {
					if !strings.Contains(string(got), want) {
						t.Errorf("body %s missing %s", got, want)
					}
				}
			})
		}
	})
}

// assertRuntimeType fails the test when got's dynamic type differs from want's.
func assertRuntimeType(t *testing.T, got, want any) {
	t.Helper()
	if g, w := reflect.TypeOf(got), reflect.TypeOf(want); g != w {
		t.Errorf("runtime type = %v, want %v", g, w)
	}
}
