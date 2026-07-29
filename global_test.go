// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// TestEventNestedObjectFieldsUseConcreteTypes pins the optional nested object
// fields of the V1 `/event` payloads to the concrete Go type of the schema they
// reference. None of these schemas is an `anyOf`/`oneOf` union nor a free-form
// object, so `any` would be wrong: it discards type safety and any
// "runtime type of [X]" comment would be factually false, because apijson's
// interface decoder yields `map[string]any` for a bare `any` field and never the
// documented struct.
//
// Fact sources:
//   - OpenAPI EventPermissionAsked.properties.tool     -> inline object {messageID, callID}
//   - OpenAPI EventQuestionAsked.properties.tool       -> $ref QuestionTool
//   - OpenAPI EventPermissionV2Asked.properties.source -> $ref PermissionV2Source
//   - OpenAPI EventQuestionV2Asked.properties.tool     -> $ref QuestionV2Tool
//   - JS SDK v2 types.gen.ts: `tool?: { messageID: string; callID: string }`,
//     `tool?: QuestionTool`, `source?: PermissionV2Source`, `tool?: QuestionV2Tool`
func TestEventNestedObjectFieldsUseConcreteTypes(t *testing.T) {
	tests := []struct {
		name  string
		typ   reflect.Type
		field string
		want  reflect.Type
	}{
		{
			name:  "permission asked tool",
			typ:   reflect.TypeFor[opencode.EventListResponseEventPermissionAskedProperties](),
			field: "Tool",
			want:  reflect.TypeFor[opencode.EventListResponseEventPermissionAskedPropertiesTool](),
		},
		{
			name:  "question asked tool",
			typ:   reflect.TypeFor[opencode.EventListResponseEventQuestionAskedProperties](),
			field: "Tool",
			want:  reflect.TypeFor[opencode.QuestionTool](),
		},
		{
			name:  "permission v2 asked source",
			typ:   reflect.TypeFor[opencode.EventListResponseEventPermissionV2AskedProperties](),
			field: "Source",
			want:  reflect.TypeFor[opencode.PermissionV2Source](),
		},
		{
			name:  "question v2 asked tool",
			typ:   reflect.TypeFor[opencode.EventListResponseEventQuestionV2AskedProperties](),
			field: "Tool",
			want:  reflect.TypeFor[opencode.QuestionV2Tool](),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field, ok := tt.typ.FieldByName(tt.field)
			if !ok {
				t.Fatalf("missing %s field", tt.field)
			}
			if field.Type != tt.want {
				t.Fatalf("%s.%s type = %s, want %s", tt.typ.Name(), tt.field, field.Type, tt.want)
			}
			if field.Type.Kind() == reflect.Pointer {
				t.Fatalf("%s.%s must not be a pointer", tt.typ.Name(), tt.field)
			}

			jsonField, ok := tt.typ.FieldByName("JSON")
			if !ok {
				t.Fatal("missing JSON metadata field")
			}
			// The metadata must still include the field so the apijson
			// framework can track it.
			jsonFieldEntry, ok := jsonField.Type.FieldByName(tt.field)
			if !ok {
				t.Fatalf("JSON metadata must contain %s field", tt.field)
			}
			if jsonFieldEntry.Type != reflect.TypeFor[apijson.Field]() {
				t.Fatalf("JSON metadata field %s type = %s, want apijson.Field", tt.field, jsonFieldEntry.Type)
			}
		})
	}
}

// TestEventNestedObjectFieldsDecode verifies the concrete nested object fields
// decode fully (not just non-nil) from a realistic SSE payload.
func TestEventNestedObjectFieldsDecode(t *testing.T) {
	t.Run("permission asked tool", func(t *testing.T) {
		var p opencode.EventListResponseEventPermissionAskedProperties
		raw := `{"always":[],"id":"per_1","metadata":{},"patterns":[],"permission":"edit",` +
			`"sessionID":"ses_1","tool":{"messageID":"msg_1","callID":"call_1"}}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if p.Tool.MessageID != "msg_1" || p.Tool.CallID != "call_1" {
			t.Errorf("Tool = %+v", p.Tool)
		}
	})

	t.Run("question asked tool", func(t *testing.T) {
		var p opencode.EventListResponseEventQuestionAskedProperties
		raw := `{"id":"que_1","sessionID":"ses_1","questions":[{"question":"q","header":"h",` +
			`"options":[{"label":"l","description":"d"}]}],"tool":{"messageID":"msg_1","callID":"call_1"}}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if p.Tool.MessageID != "msg_1" || p.Tool.CallID != "call_1" {
			t.Errorf("Tool = %+v", p.Tool)
		}
		if len(p.Questions) != 1 || p.Questions[0].Options[0].Label != "l" {
			t.Errorf("Questions = %+v", p.Questions)
		}
	})

	t.Run("permission v2 asked source", func(t *testing.T) {
		var p opencode.EventListResponseEventPermissionV2AskedProperties
		raw := `{"id":"per_1","sessionID":"ses_1","action":"edit","resources":["a"],` +
			`"source":{"type":"tool","messageID":"msg_1","callID":"call_1"}}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if p.Source.Type != opencode.PermissionV2SourceTypeTool {
			t.Errorf("Source.Type = %q", p.Source.Type)
		}
		if p.Source.MessageID != "msg_1" || p.Source.CallID != "call_1" {
			t.Errorf("Source = %+v", p.Source)
		}
	})

	t.Run("question v2 asked tool", func(t *testing.T) {
		var p opencode.EventListResponseEventQuestionV2AskedProperties
		raw := `{"id":"que_1","sessionID":"ses_1","questions":[{"question":"q","header":"h",` +
			`"options":[{"label":"l","description":"d"}]}],"tool":{"messageID":"msg_1","callID":"call_1"}}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if p.Tool.MessageID != "msg_1" || p.Tool.CallID != "call_1" {
			t.Errorf("Tool = %+v", p.Tool)
		}
	})
}

// TestEventOptionalObjectFieldsRoundtrip verifies that the apijson.Field
// metadata for Tool/Source is wired up correctly: omitting these optional
// fields from JSON must not cause an error (the apijson framework should
// tolerate their absence because the metadata is still tracked).
func TestEventOptionalObjectFieldsRoundtrip(t *testing.T) {
	// Each payload omits the optional Tool/Source field, exercising the
	// path where the apijson framework looks up the metadata entry and
	// simply finds it absent rather than treating absence as a structural
	// violation of a now-removed field.
	cases := []struct {
		name string
		data string
		into func() any
	}{
		{
			name: "permission asked without tool",
			data: `{"always":[],"id":"x","metadata":{},"patterns":[],"permission":"p","sessionID":"s"}`,
			into: func() any { return new(opencode.EventListResponseEventPermissionAskedProperties) },
		},
		{
			name: "question asked without tool",
			data: `{"id":"x","questions":[],"sessionID":"s"}`,
			into: func() any { return new(opencode.EventListResponseEventQuestionAskedProperties) },
		},
		{
			name: "permission v2 asked without source",
			data: `{"id":"x","sessionID":"s","action":"a","resources":[]}`,
			into: func() any { return new(opencode.EventListResponseEventPermissionV2AskedProperties) },
		},
		{
			name: "question v2 asked without tool",
			data: `{"id":"x","sessionID":"s","questions":[]}`,
			into: func() any { return new(opencode.EventListResponseEventQuestionV2AskedProperties) },
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tc.data), tc.into()); err != nil {
				t.Fatalf("unexpected unmarshal error: %v", err)
			}
		})
	}
}
func TestGlobalHealth(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Global.Health(context.TODO())
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalDispose(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Global.Dispose(context.TODO())
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalUpgrade(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Global.Upgrade(context.TODO(), opencode.GlobalUpgradeParams{
		Target: opencode.F("version"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalConfigGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Global.ConfigGet(context.TODO())
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalConfigUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Global.ConfigUpdate(context.TODO(), opencode.GlobalConfigUpdateParams{
		Model: opencode.F("anthropic/claude-3-5-sonnet"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalEvent(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	stream := client.Global.Event(context.TODO())
	defer stream.Close()
	for stream.Next() {
		_ = stream.Current()
	}
	if err := stream.Err(); err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// --- Unit tests (no server required) ---

// TestGlobalConfigUpdateParamsMarshalJSON verifies that GlobalConfigUpdateParams
// serializes to JSON with the correct field names matching the OpenAPI Config schema.
// This is a regression test for the fix that changed Response types to Param types.
func TestGlobalConfigUpdateParamsMarshalJSON(t *testing.T) {
	params := opencode.GlobalConfigUpdateParams{
		Model:      opencode.F("anthropic/claude-3-5-sonnet"),
		SmallModel: opencode.F("anthropic/claude-3-haiku"),
		Autoshare:  opencode.F(false),
		Instructions: opencode.F([]string{
			"CLAUDE.md",
		}),
		Agent: opencode.F(opencode.ConfigAgentParam{
			Build: opencode.F(opencode.AgentConfigParam{
				Model:  opencode.F("anthropic/claude-3-5-sonnet"),
				Prompt: opencode.F("You are a build agent."),
			}),
		}),
		Mcp: opencode.F(map[string]opencode.ConfigMcpUnionParam{
			"my-mcp": opencode.ConfigMcpLocalParam{
				Type:    opencode.F(opencode.McpLocalConfigTypeLocal),
				Command: opencode.F([]string{"npx", "my-mcp-server"}),
			},
		}),
		Attachment: opencode.F(opencode.AttachmentConfigParam{
			Image: opencode.F(opencode.ImageAttachmentConfigParam{
				AutoResize: opencode.F(true),
				MaxWidth:   opencode.F(int64(1920)),
			}),
		}),
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON is invalid: %v", err)
	}

	// Verify top-level fields
	if m["model"] != "anthropic/claude-3-5-sonnet" {
		t.Errorf("expected model=anthropic/claude-3-5-sonnet, got %v", m["model"])
	}
	if m["small_model"] != "anthropic/claude-3-haiku" {
		t.Errorf("expected small_model=anthropic/claude-3-haiku, got %v", m["small_model"])
	}

	// Verify nested agent field structure
	agentVal, ok := m["agent"].(map[string]any)
	if !ok {
		t.Fatalf("expected agent to be a map, got %T", m["agent"])
	}
	buildVal, ok := agentVal["build"].(map[string]any)
	if !ok {
		t.Fatalf("expected agent.build to be a map, got %T", agentVal["build"])
	}
	if buildVal["model"] != "anthropic/claude-3-5-sonnet" {
		t.Errorf("expected agent.build.model=anthropic/claude-3-5-sonnet, got %v", buildVal["model"])
	}

	// Verify mcp field
	mcpVal, ok := m["mcp"].(map[string]any)
	if !ok {
		t.Fatalf("expected mcp to be a map, got %T", m["mcp"])
	}
	if _, ok := mcpVal["my-mcp"]; !ok {
		t.Error("expected mcp to contain my-mcp key")
	}

	// Verify attachment field
	attachVal, ok := m["attachment"].(map[string]any)
	if !ok {
		t.Fatalf("expected attachment to be a map, got %T", m["attachment"])
	}
	imageVal, ok := attachVal["image"].(map[string]any)
	if !ok {
		t.Fatalf("expected attachment.image to be a map, got %T", attachVal["image"])
	}
	if imageVal["auto_resize"] != true {
		t.Errorf("expected attachment.image.auto_resize=true, got %v", imageVal["auto_resize"])
	}
}

// TestGlobalConfigUpdateParamsOmitsUnset verifies that unset optional fields
// are omitted from the serialized JSON (PATCH semantics).
func TestGlobalConfigUpdateParamsOmitsUnset(t *testing.T) {
	// Only set model; all other fields should be absent
	params := opencode.GlobalConfigUpdateParams{
		Model: opencode.F("anthropic/claude-3-5-sonnet"),
	}
	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	if m["model"] != "anthropic/claude-3-5-sonnet" {
		t.Errorf("expected model field, got %v", m["model"])
	}
	// unset fields must not appear
	for _, field := range []string{"agent", "mcp", "attachment", "enterprise", "experimental"} {
		if _, ok := m[field]; ok {
			t.Errorf("field %q should be omitted when not set", field)
		}
	}
}

// TestGlobalEventTypeStructFields checks the GlobalEvent struct has required fields
// and that Payload is typed as any (Union承载字段规范).
func TestGlobalEventTypeStructFields(t *testing.T) {
	typ := reflect.TypeFor[opencode.GlobalEvent]()

	// Directory (required string)
	dirField, ok := typ.FieldByName("Directory")
	if !ok {
		t.Fatal("missing Directory field on GlobalEvent")
	}
	if dirField.Type.Kind() != reflect.String {
		t.Errorf("Directory should be string, got %s", dirField.Type)
	}

	// Payload must be any (interface{})
	payloadField, ok := typ.FieldByName("Payload")
	if !ok {
		t.Fatal("missing Payload field on GlobalEvent")
	}
	if payloadField.Type != reflect.TypeFor[any]() {
		t.Errorf("Payload must be typed as any, got %s", payloadField.Type)
	}

	// JSON metadata must be present
	jsonField, ok := typ.FieldByName("JSON")
	if !ok {
		t.Fatal("missing JSON metadata field on GlobalEvent")
	}
	// Verify JSON field has apijson.Field for Payload
	payloadMeta, ok := jsonField.Type.FieldByName("Payload")
	if !ok {
		t.Fatal("JSON metadata missing Payload field")
	}
	if payloadMeta.Type != reflect.TypeFor[apijson.Field]() {
		t.Errorf("JSON.Payload should be apijson.Field, got %s", payloadMeta.Type)
	}
}

// TestGlobalEventUnmarshalV2Event verifies GlobalEvent can unmarshal a V2 event payload.
func TestGlobalEventUnmarshalV2Event(t *testing.T) {
	raw := `{
		"directory": "/home/user/project",
		"project": "proj-1",
		"workspace": "ws-1",
		"payload": {
			"id": "evt_123",
			"type": "session.created",
			"properties": {"id": "ses_abc"}
		}
	}`
	var evt opencode.GlobalEvent
	if err := evt.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if evt.Directory != "/home/user/project" {
		t.Errorf("expected directory=/home/user/project, got %q", evt.Directory)
	}
	if evt.Project != "proj-1" {
		t.Errorf("expected project=proj-1, got %q", evt.Project)
	}
	if evt.Payload == nil {
		t.Error("expected non-nil Payload")
	}
}

// TestGlobalConfigUpdateParamsExtendedCoverage verifies that additional Param types
// (Provider, Reference, Permission, Server, Skills, Compaction, Watcher, ToolOutput)
// serialize correctly as part of GlobalConfigUpdateParams.
// This extends 🔴-2 regression coverage to all nested Param types.
func TestGlobalConfigUpdateParamsExtendedCoverage(t *testing.T) {
	params := opencode.GlobalConfigUpdateParams{
		// Permission as string constant (ConfigPermissionUnionParam)
		Permission: opencode.F[opencode.ConfigPermissionUnionParam](opencode.ConfigPermissionActionAsk),
		// Provider map
		Provider: opencode.F(map[string]opencode.ConfigProviderParam{
			"anthropic": {
				ID:   opencode.F("anthropic"),
				Name: opencode.F("Anthropic"),
			},
		}),
		// Reference map with git variant
		Reference: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
			"myref": opencode.ConfigV2ReferenceGitParam{
				Repository: opencode.F("https://github.com/example/repo"),
			},
		}),
		// Compaction
		Compaction: opencode.F(opencode.ConfigCompactionParam{
			Auto: opencode.F(true),
		}),
		// Server
		Server: opencode.F(opencode.ServerConfigParam{
			Port: opencode.F(int64(8080)),
		}),
		// Skills
		Skills: opencode.F(opencode.ConfigSkillsParam{
			Paths: opencode.F([]string{"~/.claude/skills"}),
		}),
		// Watcher
		Watcher: opencode.F(opencode.ConfigWatcherParam{
			Ignore: opencode.F([]string{"node_modules"}),
		}),
		// ToolOutput
		ToolOutput: opencode.F(opencode.ConfigToolOutputParam{
			MaxLines: opencode.F(int64(2000)),
		}),
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON is invalid: %v", err)
	}

	// permission should be a string "ask"
	if m["permission"] != "ask" {
		t.Errorf("expected permission=\"ask\", got %v", m["permission"])
	}

	// provider should be a map
	if _, ok := m["provider"].(map[string]any); !ok {
		t.Errorf("expected provider to be a map, got %T", m["provider"])
	}

	// reference should be a map
	refVal, ok := m["reference"].(map[string]any)
	if !ok {
		t.Fatalf("expected reference to be a map, got %T", m["reference"])
	}
	if _, ok := refVal["myref"]; !ok {
		t.Error("expected reference to contain myref key")
	}

	// compaction should be a map
	if _, ok := m["compaction"].(map[string]any); !ok {
		t.Errorf("expected compaction to be a map, got %T", m["compaction"])
	}

	// server should be a map
	if _, ok := m["server"].(map[string]any); !ok {
		t.Errorf("expected server to be a map, got %T", m["server"])
	}

	// skills should be a map
	if _, ok := m["skills"].(map[string]any); !ok {
		t.Errorf("expected skills to be a map, got %T", m["skills"])
	}

	// watcher should be a map
	if _, ok := m["watcher"].(map[string]any); !ok {
		t.Errorf("expected watcher to be a map, got %T", m["watcher"])
	}

	// tool_output should be a map
	if _, ok := m["tool_output"].(map[string]any); !ok {
		t.Errorf("expected tool_output to be a map, got %T", m["tool_output"])
	}
}

// TestGlobalConfigUpdateParamsHasNoQueryParameters asserts that
// GlobalConfigUpdateParams carries body parameters only.
//
// OpenAPI `/global/config` PATCH declares `"parameters": []` and JS SDK v2
// `GlobalConfigUpdateData` declares `query?: never`, so the params struct must
// not implement [apiquery.Queryer] — a URLQuery method there could only ever
// encode an empty query string.
func TestGlobalConfigUpdateParamsHasNoQueryParameters(t *testing.T) {
	if _, ok := any(opencode.GlobalConfigUpdateParams{}).(apiquery.Queryer); ok {
		t.Error("GlobalConfigUpdateParams implements apiquery.Queryer, but /global/config PATCH declares no query parameters")
	}

	typ := reflect.TypeFor[opencode.GlobalConfigUpdateParams]()
	for i := range typ.NumField() {
		field := typ.Field(i)
		if tag, ok := field.Tag.Lookup("query"); ok {
			t.Errorf("field %s has query tag %q, but /global/config PATCH declares no query parameters", field.Name, tag)
		}
		if _, ok := field.Tag.Lookup("json"); !ok {
			t.Errorf("field %s is missing a json tag; all GlobalConfigUpdateParams fields are body parameters", field.Name)
		}
	}

	// Positive control: `/config` PATCH does declare `directory` and `workspace`
	// query parameters, so ConfigUpdateParams must remain a Queryer.
	if _, ok := any(opencode.ConfigUpdateParams{}).(apiquery.Queryer); !ok {
		t.Error("ConfigUpdateParams does not implement apiquery.Queryer, but /config PATCH declares directory and workspace query parameters")
	}
}

// TestSyncEventResponseSyncEventAsUnion verifies that AsUnion returns the
// concrete V1 SyncEvent variant registered for the decoded `type` discriminator,
// and that the ported Data field agrees with the variant's own Data.
func TestSyncEventResponseSyncEventAsUnion(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		wantType opencode.SyncEventResponseSyncEventType
		assert   func(t *testing.T, union opencode.SyncEventResponseSyncEventDataUnion) any
	}{
		{
			name: "message.part.removed.1",
			raw: `{
				"type": "message.part.removed.1",
				"id": "evt_part_removed",
				"seq": 7,
				"aggregateID": "ses_agg_1",
				"data": {"messageID": "msg_1", "partID": "prt_1", "sessionID": "ses_1"}
			}`,
			wantType: opencode.SyncEventResponseSyncEventTypeMessagePartRemoved1,
			assert: func(t *testing.T, union opencode.SyncEventResponseSyncEventDataUnion) any {
				variant, ok := union.(opencode.SyncEventMessagePartRemoved)
				if !ok {
					t.Fatalf("AsUnion() = %T, want opencode.SyncEventMessagePartRemoved", union)
				}
				if variant.Data.PartID != "prt_1" {
					t.Errorf("Data.PartID = %q, want %q", variant.Data.PartID, "prt_1")
				}
				return variant.Data
			},
		},
		{
			name: "session.next.text.started.1",
			raw: `{
				"type": "session.next.text.started.1",
				"id": "evt_text_started",
				"seq": 12,
				"aggregateID": "ses_agg_2",
				"data": {"timestamp": 1700000000, "assistantMessageID": "msg_2", "sessionID": "ses_2", "textID": "txt_1"}
			}`,
			wantType: opencode.SyncEventResponseSyncEventTypeSessionNextTextStarted1,
			assert: func(t *testing.T, union opencode.SyncEventResponseSyncEventDataUnion) any {
				variant, ok := union.(opencode.SyncEventSessionNextTextStarted)
				if !ok {
					t.Fatalf("AsUnion() = %T, want opencode.SyncEventSessionNextTextStarted", union)
				}
				if variant.Data.TextID != "txt_1" {
					t.Errorf("Data.TextID = %q, want %q", variant.Data.TextID, "txt_1")
				}
				if variant.Data.Timestamp != 1700000000 {
					t.Errorf("Data.Timestamp = %d, want %d", variant.Data.Timestamp, 1700000000)
				}
				return variant.Data
			},
		},
		{
			name: "session.next.tool.called.1",
			raw: `{
				"type": "session.next.tool.called.1",
				"id": "evt_tool_called",
				"seq": 33,
				"aggregateID": "ses_agg_3",
				"data": {
					"timestamp": 1700000001,
					"assistantMessageID": "msg_3",
					"sessionID": "ses_3",
					"callID": "call_1",
					"tool": "bash",
					"input": {"command": "ls"},
					"provider": {"executed": true}
				}
			}`,
			wantType: opencode.SyncEventResponseSyncEventTypeSessionNextToolCalled1,
			assert: func(t *testing.T, union opencode.SyncEventResponseSyncEventDataUnion) any {
				variant, ok := union.(opencode.SyncEventSessionNextToolCalled)
				if !ok {
					t.Fatalf("AsUnion() = %T, want opencode.SyncEventSessionNextToolCalled", union)
				}
				if variant.Data.Tool != "bash" {
					t.Errorf("Data.Tool = %q, want %q", variant.Data.Tool, "bash")
				}
				if variant.Data.CallID != "call_1" {
					t.Errorf("Data.CallID = %q, want %q", variant.Data.CallID, "call_1")
				}
				return variant.Data
			},
		},
		{
			name: "session.created.1",
			raw: `{
				"type": "session.created.1",
				"id": "evt_session_created",
				"seq": 1,
				"aggregateID": "ses_agg_4",
				"data": {"sessionID": "ses_4", "info": {"id": "ses_4", "title": "hello"}}
			}`,
			wantType: opencode.SyncEventResponseSyncEventTypeSessionCreated1,
			assert: func(t *testing.T, union opencode.SyncEventResponseSyncEventDataUnion) any {
				variant, ok := union.(opencode.SyncEventSessionCreated)
				if !ok {
					t.Fatalf("AsUnion() = %T, want opencode.SyncEventSessionCreated", union)
				}
				if variant.Data.Info.ID != "ses_4" {
					t.Errorf("Data.Info.ID = %q, want %q", variant.Data.Info.ID, "ses_4")
				}
				return variant.Data
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var event opencode.SyncEventResponseSyncEvent
			if err := json.Unmarshal([]byte(tt.raw), &event); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}

			union := event.AsUnion()
			if union == nil {
				t.Fatal("AsUnion() = nil, want the registered SyncEvent variant")
			}

			variantData := tt.assert(t, union)

			// Scalar fields must be ported from the union variant onto the root.
			if event.Type != tt.wantType {
				t.Errorf("Type = %q, want %q", event.Type, tt.wantType)
			}
			if !event.Type.IsKnown() {
				t.Errorf("Type %q is not covered by IsKnown()", event.Type)
			}
			// The polymorphic Data carrier must agree with the variant's own Data.
			if !reflect.DeepEqual(event.Data, variantData) {
				t.Errorf("Data = %#v, want %#v", event.Data, variantData)
			}
			if event.JSON.RawJSON() == "" {
				t.Error("JSON.RawJSON() is empty, want the original payload")
			}
		})
	}
}

// TestSyncEventResponseSyncEventAsUnionPortsScalars verifies the scalar fields
// ported onto the union root, and that AsUnion is stable across calls.
func TestSyncEventResponseSyncEventAsUnionPortsScalars(t *testing.T) {
	const raw = `{
		"type": "message.part.removed.1",
		"id": "evt_scalars",
		"seq": 42,
		"aggregateID": "ses_scalars",
		"data": {"messageID": "msg_s", "partID": "prt_s", "sessionID": "ses_s"}
	}`

	var event opencode.SyncEventResponseSyncEvent
	if err := json.Unmarshal([]byte(raw), &event); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	if event.ID != "evt_scalars" {
		t.Errorf("ID = %q, want %q", event.ID, "evt_scalars")
	}
	if event.Seq != 42 {
		t.Errorf("Seq = %d, want 42", event.Seq)
	}
	if event.AggregateID != "ses_scalars" {
		t.Errorf("AggregateID = %q, want %q", event.AggregateID, "ses_scalars")
	}

	// AsUnion has a value receiver and must be side-effect free.
	if !reflect.DeepEqual(event.AsUnion(), event.AsUnion()) {
		t.Error("AsUnion() is not stable across calls")
	}
}

// TestSyncEventResponseSyncEventAsUnionZeroValue verifies that an undecoded
// value reports a nil union rather than panicking.
func TestSyncEventResponseSyncEventAsUnionZeroValue(t *testing.T) {
	var event opencode.SyncEventResponseSyncEvent
	if union := event.AsUnion(); union != nil {
		t.Errorf("AsUnion() = %#v, want nil for the zero value", union)
	}
}

// TestSyncEventResponseAsUnionThroughGlobalEventPayload walks the full SSE path
// for a V1 SyncEvent: GlobalEvent.payload -> SyncEventResponse -> SyncEvent.AsUnion.
func TestSyncEventResponseAsUnionThroughGlobalEventPayload(t *testing.T) {
	const raw = `{
		"directory": "/tmp/project",
		"payload": {
			"type": "sync",
			"id": "evt_outer",
			"syncEvent": {
				"type": "session.next.text.started.1",
				"id": "evt_inner",
				"seq": 9,
				"aggregateID": "ses_agg",
				"data": {"timestamp": 1700000002, "assistantMessageID": "msg_x", "sessionID": "ses_x", "textID": "txt_x"}
			}
		}
	}`

	var event opencode.GlobalEvent
	if err := json.Unmarshal([]byte(raw), &event); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if event.Directory != "/tmp/project" {
		t.Errorf("Directory = %q, want %q", event.Directory, "/tmp/project")
	}

	payload, ok := event.AsUnion().(opencode.SyncEventResponse)
	if !ok {
		t.Fatalf("GlobalEvent.AsUnion() = %T, want opencode.SyncEventResponse", event.AsUnion())
	}
	if payload.Type != opencode.SyncEventResponseTypeSync {
		t.Errorf("payload.Type = %q, want %q", payload.Type, opencode.SyncEventResponseTypeSync)
	}
	if payload.ID != "evt_outer" {
		t.Errorf("payload.ID = %q, want %q", payload.ID, "evt_outer")
	}

	variant, ok := payload.SyncEvent.AsUnion().(opencode.SyncEventSessionNextTextStarted)
	if !ok {
		t.Fatalf("SyncEvent.AsUnion() = %T, want opencode.SyncEventSessionNextTextStarted", payload.SyncEvent.AsUnion())
	}
	if variant.Data.TextID != "txt_x" {
		t.Errorf("Data.TextID = %q, want %q", variant.Data.TextID, "txt_x")
	}
	if payload.SyncEvent.ID != "evt_inner" {
		t.Errorf("SyncEvent.ID = %q, want %q", payload.SyncEvent.ID, "evt_inner")
	}
	if payload.SyncEvent.Seq != 9 {
		t.Errorf("SyncEvent.Seq = %d, want 9", payload.SyncEvent.Seq)
	}
}

// TestGlobalEventResponseTypeRetainsNonPayloadVariants guards the two enum
// values that are not V2 `EventListResponseEvent*` payload variants:
// `server.instance.disposed` (declared by OpenAPI EventServerInstanceDisposed)
// and `sync` (declared by the 35 OpenAPI SyncEvent* schemas). Both are part of
// the public surface and must never be dropped.
func TestGlobalEventResponseTypeRetainsNonPayloadVariants(t *testing.T) {
	tests := []struct {
		value opencode.GlobalEventResponseType
		want  string
	}{
		{opencode.GlobalEventResponseTypeServerInstanceDisposed, "server.instance.disposed"},
		{opencode.GlobalEventResponseTypeSync, "sync"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if string(tt.value) != tt.want {
				t.Errorf("value = %q, want %q", string(tt.value), tt.want)
			}
			if !tt.value.IsKnown() {
				t.Errorf("IsKnown() = false for %q, want true", tt.want)
			}
		})
	}
}
