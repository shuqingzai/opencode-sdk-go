// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestConfigGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Config.Get(context.TODO(), opencode.ConfigGetParams{
		Directory: opencode.F("directory"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigUpdate(t *testing.T) {
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
	_, err := client.Config.Update(
		context.TODO(),
		opencode.ConfigUpdateParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestConfigProviderModelInterleavedUnmarshal verifies that
// ConfigProviderModel.Interleaved correctly deserialises all four OpenAPI
// anyOf variants of ProviderConfig.models.*.interleaved:
//
//  1. boolean  → runtime type bool
//  2. enum string ("reasoning"|"reasoning_content"|"reasoning_text") → string
//  3. open string (any vendor-defined value)  → string
//  4. object { "field": string } → map[string]any
//
// The field is typed as any with no apijson.RegisterUnion registration
// because it is a leaf anyOf carried directly by apijson's reflect.Interface
// branch (decoder.go:183-191), which hands the gjson native value straight to
// the field. No t.Skip — these are pure unit tests with no HTTP dependency.
//
// Run with: go test -run TestConfigProviderModelInterleavedUnmarshal -v ./...
func TestConfigProviderModelInterleavedUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name         string
		json         string
		wantRuntime  reflect.Type
		wantValue    any
		wantMapField any // non-nil only for map variant
	}{
		// Variant 1a: boolean true → bool
		{
			name:        "bool_true",
			json:        `{"interleaved":true}`,
			wantRuntime: reflect.TypeFor[bool](),
			wantValue:   true,
		},
		// Variant 1b: boolean false → bool
		{
			name:        "bool_false",
			json:        `{"interleaved":false}`,
			wantRuntime: reflect.TypeFor[bool](),
			wantValue:   false,
		},
		// Variant 2a: enum string "reasoning" → string
		{
			name:        "enum_reasoning",
			json:        `{"interleaved":"reasoning"}`,
			wantRuntime: reflect.TypeFor[string](),
			wantValue:   "reasoning",
		},
		// Variant 2b: enum string "reasoning_content" → string
		{
			name:        "enum_reasoning_content",
			json:        `{"interleaved":"reasoning_content"}`,
			wantRuntime: reflect.TypeFor[string](),
			wantValue:   "reasoning_content",
		},
		// Variant 2c: enum string "reasoning_text" → string
		{
			name:        "enum_reasoning_text",
			json:        `{"interleaved":"reasoning_text"}`,
			wantRuntime: reflect.TypeFor[string](),
			wantValue:   "reasoning_text",
		},
		// Variant 3: open string (vendor-defined) → string
		{
			name:        "open_string",
			json:        `{"interleaved":"vendor_custom_field"}`,
			wantRuntime: reflect.TypeFor[string](),
			wantValue:   "vendor_custom_field",
		},
		// Variant 4a: object with known enum field → map[string]any
		// OpenAPI: { "field": "reasoning"|"reasoning_content"|"reasoning_text"|string }
		{
			name:         "object_field_reasoning_text",
			json:         `{"interleaved":{"field":"reasoning_text"}}`,
			wantRuntime:  reflect.TypeFor[map[string]any](),
			wantMapField: "reasoning_text",
		},
		// Variant 4b: object with open (vendor-custom) field value → map[string]any
		{
			name:         "object_field_vendor_custom",
			json:         `{"interleaved":{"field":"vendor_custom"}}`,
			wantRuntime:  reflect.TypeFor[map[string]any](),
			wantMapField: "vendor_custom",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var m opencode.ConfigProviderModel
			if err := json.Unmarshal([]byte(tc.json), &m); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			gotRuntime := reflect.TypeOf(m.Interleaved)
			if gotRuntime != tc.wantRuntime {
				t.Errorf("runtime type: got %v, want %v (value: %#v)", gotRuntime, tc.wantRuntime, m.Interleaved)
			}
			if tc.wantMapField != nil {
				// object variant: runtime type must be map[string]any
				gotMap, ok := m.Interleaved.(map[string]any)
				if !ok {
					t.Fatalf("expected map[string]any, got %T", m.Interleaved)
				}
				gotField, exists := gotMap["field"]
				if !exists {
					t.Fatal("map missing key \"field\"")
				}
				if gotField != tc.wantMapField {
					t.Errorf("map[\"field\"]: got %q, want %q", gotField, tc.wantMapField)
				}
			} else {
				// scalar variant: value must match exactly
				if m.Interleaved != tc.wantValue {
					t.Errorf("value: got %#v, want %#v", m.Interleaved, tc.wantValue)
				}
			}
		})
	}
}

// TestPermissionConfigObjectRuleUnmarshal verifies that every OpenAPI
// `PermissionConfig` object property typed as `PermissionRuleConfig` accepts both
// wire shapes:
//
//  1. string form ("ask"|"allow"|"deny")
//  2. map form ({"src/**": "allow"})
//
// OpenAPI: components.schemas.PermissionConfig.anyOf[1].properties.
//
// The rule fields are statically typed as the carrier [opencode.PermissionRuleConfig];
// apijson resolves the registered union inside the carrier and callers get the
// concrete variant type via [opencode.PermissionRuleConfig.AsUnion] — no `any`
// type-switching on Go's generic JSON types.
//
// Run with: go test -run TestPermissionConfigObjectRuleUnmarshal -v ./...
func TestPermissionConfigObjectRuleUnmarshal(t *testing.T) {
	t.Parallel()
	ruleFields := []struct {
		key string
		get func(opencode.PermissionConfigObject) opencode.PermissionRuleConfig
	}{
		{"read", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.Read }},
		{"edit", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.Edit }},
		{"glob", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.Glob }},
		{"grep", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.Grep }},
		{"list", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.List }},
		{"bash", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.Bash }},
		{"task", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.Task }},
		{"external_directory", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.ExternalDirectory }},
		{"lsp", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.Lsp }},
		{"skill", func(p opencode.PermissionConfigObject) opencode.PermissionRuleConfig { return p.Skill }},
	}
	if len(ruleFields) != 10 {
		t.Fatalf("OpenAPI defines 10 PermissionRuleConfig properties, test covers %d", len(ruleFields))
	}

	for _, rf := range ruleFields {
		t.Run(rf.key+"_string", func(t *testing.T) {
			t.Parallel()
			var p opencode.PermissionConfigObject
			if err := json.Unmarshal([]byte(`{"`+rf.key+`":"allow"}`), &p); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			got, ok := rf.get(p).AsUnion().(opencode.PermissionActionConfig)
			if !ok {
				t.Fatalf("runtime type = %T, want opencode.PermissionActionConfig", rf.get(p).AsUnion())
			}
			if got != opencode.PermissionActionConfigAllow {
				t.Errorf("value = %q, want %q", got, opencode.PermissionActionConfigAllow)
			}
			if !got.IsKnown() {
				t.Errorf("PermissionActionConfig(%q).IsKnown() = false", got)
			}
		})

		t.Run(rf.key+"_map", func(t *testing.T) {
			t.Parallel()
			var p opencode.PermissionConfigObject
			if err := json.Unmarshal([]byte(`{"`+rf.key+`":{"src/**":"allow","**":"ask"}}`), &p); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			obj, ok := rf.get(p).AsUnion().(opencode.PermissionObjectConfig)
			if !ok {
				t.Fatalf("runtime type = %T, want opencode.PermissionObjectConfig", rf.get(p).AsUnion())
			}
			if obj["src/**"] != opencode.PermissionActionConfigAllow {
				t.Errorf(`["src/**"] = %q, want "allow"`, obj["src/**"])
			}
			if obj["**"] != opencode.PermissionActionConfigAsk {
				t.Errorf(`["**"] = %q, want "ask"`, obj["**"])
			}
		})

		t.Run(rf.key+"_null_and_absent", func(t *testing.T) {
			t.Parallel()
			for _, body := range []string{`{"` + rf.key + `":null}`, `{}`} {
				var p opencode.PermissionConfigObject
				if err := json.Unmarshal([]byte(body), &p); err != nil {
					t.Fatalf("%s: json.Unmarshal: %v", body, err)
				}
				if rf.get(p).AsUnion() != nil {
					t.Errorf("%s: got %#v, want nil", body, rf.get(p).AsUnion())
				}
			}
		})
	}
}

// TestPermissionConfigObjectActionFields verifies that the five OpenAPI
// `PermissionConfig` properties typed as plain `PermissionActionConfig`
// (todowrite/question/webfetch/websearch/doom_loop) decode as typed enums.
// These tools take no matchable argument, so the spec forbids the map form.
//
// Run with: go test -run TestPermissionConfigObjectActionFields -v ./...
func TestPermissionConfigObjectActionFields(t *testing.T) {
	t.Parallel()
	const raw = `{"todowrite":"allow","question":"ask","webfetch":"deny","websearch":"allow","doom_loop":"ask"}`
	var p opencode.PermissionConfigObject
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	for name, got := range map[string]opencode.PermissionActionConfig{
		"todowrite": p.Todowrite,
		"question":  p.Question,
		"webfetch":  p.Webfetch,
		"websearch": p.Websearch,
		"doom_loop": p.DoomLoop,
	} {
		if !got.IsKnown() {
			t.Errorf("%s: IsKnown() = false for %q", name, got)
		}
	}
	if p.Todowrite != opencode.PermissionActionConfigAllow {
		t.Errorf("todowrite = %q, want allow", p.Todowrite)
	}
	if p.Question != opencode.PermissionActionConfigAsk {
		t.Errorf("question = %q, want ask", p.Question)
	}
	if p.Webfetch != opencode.PermissionActionConfigDeny {
		t.Errorf("webfetch = %q, want deny", p.Webfetch)
	}
	if p.JSON.RawJSON() != raw {
		t.Errorf("RawJSON() = %q, want %q", p.JSON.RawJSON(), raw)
	}
}

// TestPermissionConfigObjectExtraFields verifies the OpenAPI
// `PermissionConfig.additionalProperties -> PermissionRuleConfig` mapping.
//
// Run with: go test -run TestPermissionConfigObjectExtraFields -v ./...
func TestPermissionConfigObjectExtraFields(t *testing.T) {
	t.Parallel()
	var p opencode.PermissionConfigObject
	if err := json.Unmarshal([]byte(`{"edit":"ask","custom_tool":"deny","other_tool":{"a/**":"allow"}}`), &p); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if len(p.ExtraFields) != 2 {
		t.Fatalf("len(ExtraFields) = %d, want 2 (got %#v)", len(p.ExtraFields), p.ExtraFields)
	}
	action, ok := p.ExtraFields["custom_tool"].AsUnion().(opencode.PermissionActionConfig)
	if !ok {
		t.Fatalf(`ExtraFields["custom_tool"] type = %T, want opencode.PermissionActionConfig`, p.ExtraFields["custom_tool"].AsUnion())
	}
	if action != opencode.PermissionActionConfigDeny {
		t.Errorf(`ExtraFields["custom_tool"] = %q, want "deny"`, action)
	}
	obj, ok := p.ExtraFields["other_tool"].AsUnion().(opencode.PermissionObjectConfig)
	if !ok {
		t.Fatalf(`ExtraFields["other_tool"] type = %T, want opencode.PermissionObjectConfig`, p.ExtraFields["other_tool"].AsUnion())
	}
	if obj["a/**"] != opencode.PermissionActionConfigAllow {
		t.Errorf(`ExtraFields["other_tool"]["a/**"] = %q, want "allow"`, obj["a/**"])
	}
	// Named properties must never leak into ExtraFields.
	if _, bad := p.ExtraFields["edit"]; bad {
		t.Error(`named property "edit" leaked into ExtraFields`)
	}
}

// TestAgentConfigPermissionUnmarshal verifies that [opencode.AgentConfig].Permission
// deserialises both OpenAPI `PermissionConfig` anyOf variants:
//
//  1. string form ("ask"|"allow"|"deny") -> [opencode.PermissionActionConfig]
//  2. object form ({...per-tool rules...}) -> [opencode.PermissionConfigObject]
//
// Run with: go test -run TestAgentConfigPermissionUnmarshal -v ./...
func TestAgentConfigPermissionUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		json        string
		wantRuntime reflect.Type
	}{
		{"string_ask", `{"permission":"ask"}`, reflect.TypeFor[opencode.PermissionActionConfig]()},
		{"string_allow", `{"permission":"allow"}`, reflect.TypeFor[opencode.PermissionActionConfig]()},
		{"string_deny", `{"permission":"deny"}`, reflect.TypeFor[opencode.PermissionActionConfig]()},
		{"object_form", `{"permission":{"edit":"allow","bash":{"git *":"allow"}}}`, reflect.TypeFor[opencode.PermissionConfigObject]()},
		{"empty_object", `{"permission":{}}`, reflect.TypeFor[opencode.PermissionConfigObject]()},
		{"null", `{"permission":null}`, nil},
		{"absent", `{}`, nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var a opencode.AgentConfig
			if err := json.Unmarshal([]byte(tc.json), &a); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			if got := reflect.TypeOf(a.Permission); got != tc.wantRuntime {
				t.Fatalf("Permission runtime type = %v, want %v", got, tc.wantRuntime)
			}
			if got := reflect.TypeOf(a.AsPermission()); got != tc.wantRuntime {
				t.Errorf("AsPermission() runtime type = %v, want %v", got, tc.wantRuntime)
			}
			if action, ok := a.Permission.(opencode.PermissionActionConfig); ok && !action.IsKnown() {
				t.Errorf("PermissionActionConfig.IsKnown() = false for %q", action)
			}
		})
	}
}

// TestConfigTopLevelPermissionUnmarshal mirrors the AgentConfig test for the
// top-level `Config.permission` field, including the `null` case the live server
// actually returns.
//
// Run with: go test -run TestConfigTopLevelPermissionUnmarshal -v ./...
func TestConfigTopLevelPermissionUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		json        string
		wantRuntime reflect.Type
	}{
		{"string_ask", `{"permission":"ask"}`, reflect.TypeFor[opencode.PermissionActionConfig]()},
		{"string_deny", `{"permission":"deny"}`, reflect.TypeFor[opencode.PermissionActionConfig]()},
		{"object_form", `{"permission":{"edit":"allow"}}`, reflect.TypeFor[opencode.PermissionConfigObject]()},
		{"empty_object", `{"permission":{}}`, reflect.TypeFor[opencode.PermissionConfigObject]()},
		// The live OpenCode server returns `"permission": null` for GET /config
		// without a directory; this must not fail to decode.
		{"null", `{"permission":null}`, nil},
		{"absent", `{"model":"a/b"}`, nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var c opencode.Config
			if err := json.Unmarshal([]byte(tc.json), &c); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			if got := reflect.TypeOf(c.Permission); got != tc.wantRuntime {
				t.Errorf("Config.Permission runtime type = %v, want %v", got, tc.wantRuntime)
			}
			if got := reflect.TypeOf(c.AsPermission()); got != tc.wantRuntime {
				t.Errorf("Config.AsPermission() runtime type = %v, want %v", got, tc.wantRuntime)
			}
		})
	}
}

// TestConfigAgentSlotsShareAgentConfig asserts that every agent slot of
// [opencode.ConfigAgent] and [opencode.ConfigMode] uses the single
// [opencode.AgentConfig] type, matching the OpenAPI spec where all of
// `Config.agent.{plan,build,general,explore,title,summary,compaction}` and
// `Config.mode.{build,plan}` are `$ref: AgentConfig`.
//
// Run with: go test -run TestConfigAgentSlotsShareAgentConfig -v ./...
func TestConfigAgentSlotsShareAgentConfig(t *testing.T) {
	t.Parallel()
	want := reflect.TypeFor[opencode.AgentConfig]()
	wantMap := reflect.TypeFor[map[string]opencode.AgentConfig]()

	agentT := reflect.TypeFor[opencode.ConfigAgent]()
	for _, name := range []string{"Build", "Compaction", "Explore", "General", "Plan", "Summary", "Title"} {
		f, ok := agentT.FieldByName(name)
		if !ok {
			t.Fatalf("ConfigAgent.%s: field missing", name)
		}
		if f.Type != want {
			t.Errorf("ConfigAgent.%s type = %v, want %v", name, f.Type, want)
		}
	}
	if f, _ := agentT.FieldByName("ExtraFields"); f.Type != wantMap {
		t.Errorf("ConfigAgent.ExtraFields type = %v, want %v", f.Type, wantMap)
	}

	modeT := reflect.TypeFor[opencode.ConfigMode]()
	for _, name := range []string{"Build", "Plan"} {
		f, ok := modeT.FieldByName(name)
		if !ok {
			t.Fatalf("ConfigMode.%s: field missing", name)
		}
		if f.Type != want {
			t.Errorf("ConfigMode.%s type = %v, want %v", name, f.Type, want)
		}
	}
	if f, _ := modeT.FieldByName("ExtraFields"); f.Type != wantMap {
		t.Errorf("ConfigMode.ExtraFields type = %v, want %v", f.Type, wantMap)
	}
}

// TestConfigAgentCustomAgentUnmarshal is the regression test for the
// `Config.agent.additionalProperties -> AgentConfig` mapping. Previously
// ConfigAgent.ExtraFields was typed map[string]ConfigAgent, which corrupted every
// user-defined agent: each scalar property was decoded into a whole empty
// ConfigAgent instead of the agent's own field.
//
// Run with: go test -run TestConfigAgentCustomAgentUnmarshal -v ./...
func TestConfigAgentCustomAgentUnmarshal(t *testing.T) {
	t.Parallel()
	const raw = `{
		"build":{"model":"anthropic/claude-sonnet-4-6"},
		"reviewer":{
			"model":"openai/gpt-5",
			"temperature":0.3,
			"top_p":0.9,
			"mode":"subagent",
			"description":"Reviews code",
			"prompt":"You are a reviewer",
			"disable":false,
			"hidden":true,
			"color":"#FF5733",
			"steps":12,
			"maxSteps":34,
			"tools":{"bash":true,"edit":false},
			"options":{"k":"v"},
			"permission":{"edit":"deny"},
			"name":"reviewer"
		}
	}`

	var a opencode.ConfigAgent
	if err := json.Unmarshal([]byte(raw), &a); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	if a.Build.Model != "anthropic/claude-sonnet-4-6" {
		t.Errorf("Build.Model = %q, want %q", a.Build.Model, "anthropic/claude-sonnet-4-6")
	}
	// Named slots absent from the payload must stay zero-valued.
	if a.Plan.Model != "" {
		t.Errorf("Plan.Model = %q, want empty", a.Plan.Model)
	}
	// Named slots must never leak into ExtraFields.
	if _, bad := a.ExtraFields["build"]; bad {
		t.Error(`named slot "build" leaked into ExtraFields`)
	}

	rev, ok := a.ExtraFields["reviewer"]
	if !ok {
		t.Fatalf(`ExtraFields["reviewer"] missing; got %v`, sortedTestKeys(a.ExtraFields))
	}
	if len(a.ExtraFields) != 1 {
		t.Errorf("len(ExtraFields) = %d, want 1", len(a.ExtraFields))
	}

	if rev.Model != "openai/gpt-5" {
		t.Errorf("reviewer.Model = %q, want %q", rev.Model, "openai/gpt-5")
	}
	if rev.Temperature != 0.3 {
		t.Errorf("reviewer.Temperature = %v, want 0.3", rev.Temperature)
	}
	if rev.TopP != 0.9 {
		t.Errorf("reviewer.TopP = %v, want 0.9", rev.TopP)
	}
	if rev.Mode != opencode.AgentConfigModeSubagent {
		t.Errorf("reviewer.Mode = %q, want %q", rev.Mode, opencode.AgentConfigModeSubagent)
	}
	if !rev.Mode.IsKnown() {
		t.Errorf("reviewer.Mode.IsKnown() = false for %q", rev.Mode)
	}
	if rev.Description != "Reviews code" {
		t.Errorf("reviewer.Description = %q", rev.Description)
	}
	if rev.Prompt != "You are a reviewer" {
		t.Errorf("reviewer.Prompt = %q", rev.Prompt)
	}
	if rev.Disable {
		t.Error("reviewer.Disable = true, want false")
	}
	if !rev.Hidden {
		t.Error("reviewer.Hidden = false, want true")
	}
	if rev.Color != "#FF5733" {
		t.Errorf("reviewer.Color = %q", rev.Color)
	}
	if rev.Steps != 12 {
		t.Errorf("reviewer.Steps = %d, want 12", rev.Steps)
	}
	if rev.MaxSteps != 34 {
		t.Errorf("reviewer.MaxSteps = %d, want 34", rev.MaxSteps)
	}
	if !rev.Tools["bash"] || rev.Tools["edit"] {
		t.Errorf("reviewer.Tools = %#v, want {bash:true, edit:false}", rev.Tools)
	}
	if rev.Options["k"] != "v" {
		t.Errorf("reviewer.Options = %#v", rev.Options)
	}
	perm, ok := rev.Permission.(opencode.PermissionConfigObject)
	if !ok {
		t.Fatalf("reviewer.Permission type = %T, want opencode.PermissionConfigObject", rev.Permission)
	}
	if edit, _ := perm.Edit.AsUnion().(opencode.PermissionActionConfig); edit != opencode.PermissionActionConfigDeny {
		t.Errorf("reviewer.Permission.Edit = %#v, want PermissionActionConfigDeny", perm.Edit.AsUnion())
	}
	// `name` is not part of the OpenAPI AgentConfig schema; AgentConfig allows
	// arbitrary additional properties, so it must land in AgentConfig.ExtraFields.
	if rev.ExtraFields["name"] != "reviewer" {
		t.Errorf(`reviewer.ExtraFields["name"] = %#v, want "reviewer"`, rev.ExtraFields["name"])
	}
	if rev.JSON.RawJSON() == "" {
		t.Error("reviewer.JSON.RawJSON() is empty")
	}
}

// TestConfigModeCustomEntryUnmarshal covers the deprecated `Config.mode` map,
// which shares the same `additionalProperties -> AgentConfig` mapping.
//
// Run with: go test -run TestConfigModeCustomEntryUnmarshal -v ./...
func TestConfigModeCustomEntryUnmarshal(t *testing.T) {
	t.Parallel()
	var m opencode.ConfigMode
	if err := json.Unmarshal([]byte(`{"build":{"model":"a/b"},"custom":{"model":"c/d","temperature":0.7}}`), &m); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if m.Build.Model != "a/b" {
		t.Errorf("Build.Model = %q, want a/b", m.Build.Model)
	}
	custom, ok := m.ExtraFields["custom"]
	if !ok {
		t.Fatal(`ExtraFields["custom"] missing`)
	}
	if custom.Model != "c/d" {
		t.Errorf("custom.Model = %q, want c/d", custom.Model)
	}
	if custom.Temperature != 0.7 {
		t.Errorf("custom.Temperature = %v, want 0.7", custom.Temperature)
	}
}

// TestAgentConfigEnumsAndEdgeValues covers AgentConfigMode / PermissionActionConfig
// enum coverage plus numeric, unicode and RawJSON edge cases.
//
// Run with: go test -run TestAgentConfigEnumsAndEdgeValues -v ./...
func TestAgentConfigEnumsAndEdgeValues(t *testing.T) {
	t.Parallel()

	t.Run("AgentConfigMode_IsKnown", func(t *testing.T) {
		t.Parallel()
		for _, m := range []opencode.AgentConfigMode{
			opencode.AgentConfigModeSubagent, opencode.AgentConfigModePrimary, opencode.AgentConfigModeAll,
		} {
			if !m.IsKnown() {
				t.Errorf("IsKnown() = false for known %q", m)
			}
		}
		if opencode.AgentConfigMode("bogus").IsKnown() {
			t.Error(`IsKnown() = true for unknown "bogus"`)
		}
	})

	t.Run("PermissionActionConfig_IsKnown", func(t *testing.T) {
		t.Parallel()
		for _, a := range []opencode.PermissionActionConfig{
			opencode.PermissionActionConfigAsk, opencode.PermissionActionConfigAllow, opencode.PermissionActionConfigDeny,
		} {
			if !a.IsKnown() {
				t.Errorf("IsKnown() = false for known %q", a)
			}
		}
		if opencode.PermissionActionConfig("bogus").IsKnown() {
			t.Error(`IsKnown() = true for unknown "bogus"`)
		}
	})

	t.Run("numeric_and_unicode_edges", func(t *testing.T) {
		t.Parallel()
		const raw = `{"steps":9223372036854775807,"maxSteps":-1,"temperature":1,"top_p":0,` +
			`"model":"提供商/模型-名称","color":"#FF5733","prompt":"你好世界"}`
		var a opencode.AgentConfig
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if a.Steps != 9223372036854775807 {
			t.Errorf("Steps = %d, want max int64", a.Steps)
		}
		if a.MaxSteps != -1 {
			t.Errorf("MaxSteps = %d, want -1", a.MaxSteps)
		}
		// integer 1 on the wire must widen to float64 1.0
		if a.Temperature != 1 {
			t.Errorf("Temperature = %v, want 1", a.Temperature)
		}
		if a.Model != "提供商/模型-名称" {
			t.Errorf("Model = %q", a.Model)
		}
		if a.Prompt != "你好世界" {
			t.Errorf("Prompt = %q", a.Prompt)
		}
		if a.JSON.RawJSON() != raw {
			t.Errorf("RawJSON() mismatch\n got: %s\nwant: %s", a.JSON.RawJSON(), raw)
		}
	})

	t.Run("unknown_properties_go_to_ExtraFields", func(t *testing.T) {
		t.Parallel()
		var a opencode.AgentConfig
		if err := json.Unmarshal([]byte(`{"model":"a/b","name":"x","future_field":42}`), &a); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if a.Model != "a/b" {
			t.Errorf("Model = %q", a.Model)
		}
		if a.ExtraFields["name"] != "x" {
			t.Errorf(`ExtraFields["name"] = %#v`, a.ExtraFields["name"])
		}
		if a.ExtraFields["future_field"] != float64(42) {
			t.Errorf(`ExtraFields["future_field"] = %#v`, a.ExtraFields["future_field"])
		}
		if _, bad := a.ExtraFields["model"]; bad {
			t.Error(`named property "model" leaked into ExtraFields`)
		}
	})
}

// TestAgentConfigParamMarshal verifies the request-side [opencode.AgentConfigParam]
// serialises exactly the OpenAPI `AgentConfig` property names, and that unset
// fields are omitted (PATCH semantics).
//
// Run with: go test -run TestAgentConfigParamMarshal -v ./...
func TestAgentConfigParamMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		param    opencode.AgentConfigParam
		wantJSON string
	}{
		{
			name:     "empty_omits_everything",
			param:    opencode.AgentConfigParam{},
			wantJSON: `{}`,
		},
		{
			name: "scalar_fields",
			param: opencode.AgentConfigParam{
				Model:       opencode.F("openai/gpt-5"),
				Temperature: opencode.F(0.25),
				TopP:        opencode.F(0.9),
				Mode:        opencode.F(opencode.AgentConfigModeSubagent),
				Steps:       opencode.F(int64(7)),
				MaxSteps:    opencode.F(int64(9)),
			},
			wantJSON: `{"maxSteps":9,"mode":"subagent","model":"openai/gpt-5","steps":7,"temperature":0.25,"top_p":0.9}`,
		},
		{
			name: "permission_string_variant",
			param: opencode.AgentConfigParam{
				Permission: opencode.F[opencode.PermissionConfigUnionParam](opencode.PermissionActionConfigAsk),
			},
			wantJSON: `{"permission":"ask"}`,
		},
		{
			name: "permission_object_variant",
			param: opencode.AgentConfigParam{
				Permission: opencode.F[opencode.PermissionConfigUnionParam](opencode.PermissionConfigObjectParam{
					Edit:      opencode.F[opencode.PermissionRuleConfigUnionParam](opencode.PermissionActionConfigDeny),
					Bash:      opencode.F[opencode.PermissionRuleConfigUnionParam](opencode.PermissionObjectConfig{"git *": opencode.PermissionActionConfigAllow}),
					Todowrite: opencode.F(opencode.PermissionActionConfigAllow),
				}),
			},
			wantJSON: `{"permission":{"bash":{"git *":"allow"},"edit":"deny","todowrite":"allow"}}`,
		},
		{
			name: "extra_properties",
			param: opencode.AgentConfigParam{
				Model:       opencode.F("a/b"),
				ExtraFields: map[string]any{"name": "reviewer"},
			},
			wantJSON: `{"model":"a/b","name":"reviewer"}`,
		},
		{
			name: "tools_and_options",
			param: opencode.AgentConfigParam{
				Tools:   opencode.F(map[string]bool{"bash": true}),
				Options: opencode.F(map[string]any{"k": "v"}),
			},
			wantJSON: `{"options":{"k":"v"},"tools":{"bash":true}}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			if string(got) != tc.wantJSON {
				t.Errorf("got  %s\nwant %s", got, tc.wantJSON)
			}
		})
	}
}

// TestPermissionConfigObjectParamMarshal verifies the request-side per-tool
// permission object, including both PermissionRuleConfig variants and extras.
//
// Run with: go test -run TestPermissionConfigObjectParamMarshal -v ./...
func TestPermissionConfigObjectParamMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		param    opencode.PermissionConfigObjectParam
		wantJSON string
	}{
		{"empty", opencode.PermissionConfigObjectParam{}, `{}`},
		{
			"rule_string_variant",
			opencode.PermissionConfigObjectParam{
				Read: opencode.F[opencode.PermissionRuleConfigUnionParam](opencode.PermissionActionConfigAsk),
				Bash: opencode.F[opencode.PermissionRuleConfigUnionParam](opencode.PermissionActionConfigAllow),
			},
			`{"bash":"allow","read":"ask"}`,
		},
		{
			"rule_object_variant",
			opencode.PermissionConfigObjectParam{
				Bash: opencode.F[opencode.PermissionRuleConfigUnionParam](
					opencode.PermissionObjectConfig{"git *": opencode.PermissionActionConfigAllow, "rm *": opencode.PermissionActionConfigDeny}),
			},
			`{"bash":{"git *":"allow","rm *":"deny"}}`,
		},
		{
			"action_only_tools",
			opencode.PermissionConfigObjectParam{
				Todowrite: opencode.F(opencode.PermissionActionConfigAllow),
				DoomLoop:  opencode.F(opencode.PermissionActionConfigDeny),
			},
			`{"doom_loop":"deny","todowrite":"allow"}`,
		},
		{
			"extra_tool",
			opencode.PermissionConfigObjectParam{
				Read:        opencode.F[opencode.PermissionRuleConfigUnionParam](opencode.PermissionActionConfigAllow),
				ExtraFields: map[string]opencode.PermissionRuleConfigUnionParam{"my_tool": opencode.PermissionActionConfigAsk},
			},
			`{"read":"allow","my_tool":"ask"}`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			if string(got) != tc.wantJSON {
				t.Errorf("got  %s\nwant %s", got, tc.wantJSON)
			}
		})
	}
}

// TestConfigAgentParamMarshal verifies that [opencode.ConfigAgentParam] emits
// named agent slots plus user-defined agents (additionalProperties).
//
// Run with: go test -run TestConfigAgentParamMarshal -v ./...
func TestConfigAgentParamMarshal(t *testing.T) {
	t.Parallel()
	p := opencode.ConfigAgentParam{
		Build: opencode.F(opencode.AgentConfigParam{
			Model: opencode.F("anthropic/claude-sonnet-4-6"),
		}),
		ExtraFields: map[string]opencode.AgentConfigParam{
			"reviewer": {
				Model:       opencode.F("openai/gpt-5"),
				Temperature: opencode.F(0.3),
			},
		},
	}
	got, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	const want = `{"build":{"model":"anthropic/claude-sonnet-4-6"},"reviewer":{"model":"openai/gpt-5","temperature":0.3}}`
	if string(got) != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}

	m := opencode.ConfigModeParam{
		Build:       opencode.F(opencode.AgentConfigParam{Model: opencode.F("a/b")}),
		ExtraFields: map[string]opencode.AgentConfigParam{"custom": {Prompt: opencode.F("p")}},
	}
	gotMode, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("json.Marshal(mode): %v", err)
	}
	const wantMode = `{"build":{"model":"a/b"},"custom":{"prompt":"p"}}`
	if string(gotMode) != wantMode {
		t.Errorf("mode: got  %s\nwant %s", gotMode, wantMode)
	}
}

// TestConfigRoundTripRawJSON asserts Config and every nested AgentConfig preserve
// the exact bytes received from the server.
//
// Run with: go test -run TestConfigRoundTripRawJSON -v ./...
func TestConfigRoundTripRawJSON(t *testing.T) {
	t.Parallel()
	const raw = `{"model":"a/b","agent":{"build":{"model":"c/d"},"custom":{"model":"e/f","temperature":0.5}},"permission":null}`
	var c opencode.Config
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if c.JSON.RawJSON() != raw {
		t.Errorf("Config.RawJSON() mismatch\n got: %s\nwant: %s", c.JSON.RawJSON(), raw)
	}
	if got, want := c.Agent.Build.JSON.RawJSON(), `{"model":"c/d"}`; got != want {
		t.Errorf("Agent.Build.RawJSON() = %s, want %s", got, want)
	}
	custom := c.Agent.ExtraFields["custom"]
	if got, want := custom.JSON.RawJSON(), `{"model":"e/f","temperature":0.5}`; got != want {
		t.Errorf("custom agent RawJSON() = %s, want %s", got, want)
	}
	if c.Permission != nil {
		t.Errorf("Config.Permission = %#v, want nil", c.Permission)
	}
}

// TestConfigProviderOptionsTimeoutUnion verifies the OpenAPI
// `ProviderConfig.options.timeout` / `headerTimeout` / `chunkTimeout` anyOf:
//
//	anyOf: [ {type: integer, exclusiveMinimum: 0}, {type: boolean, enum: [false]} ]
//
// All three fields are statically typed as the carrier
// [opencode.ConfigProviderOptionsTimeout]; the registered union resolves to
// [shared.UnionInt] (int64 milliseconds) or [shared.UnionBool] via
// [opencode.ConfigProviderOptionsTimeout.AsUnion].
// Declaring them `any` would silently yield float64 / bool, violating the
// OpenAPI `integer` contract and losing precision above 2^53.
//
// Run with: go test -run TestConfigProviderOptionsTimeoutUnion -v ./...
func TestConfigProviderOptionsTimeoutUnion(t *testing.T) {
	t.Parallel()

	t.Run("integer_variant", func(t *testing.T) {
		t.Parallel()
		var o opencode.ConfigProviderOptions
		if err := json.Unmarshal([]byte(`{"timeout":600000,"headerTimeout":30000,"chunkTimeout":120000}`), &o); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		ms, ok := o.Timeout.AsUnion().(opencode.UnionInt)
		if !ok {
			t.Fatalf("Timeout runtime type = %T, want shared.UnionInt", o.Timeout.AsUnion())
		}
		if int64(ms) != 600000 {
			t.Errorf("Timeout = %d, want 600000", int64(ms))
		}
		hms, ok := o.HeaderTimeout.AsUnion().(opencode.UnionInt)
		if !ok {
			t.Fatalf("HeaderTimeout runtime type = %T, want shared.UnionInt", o.HeaderTimeout.AsUnion())
		}
		if int64(hms) != 30000 {
			t.Errorf("HeaderTimeout = %d, want 30000", int64(hms))
		}
		cms, ok := o.ChunkTimeout.AsUnion().(opencode.UnionInt)
		if !ok {
			t.Fatalf("ChunkTimeout runtime type = %T, want shared.UnionInt", o.ChunkTimeout.AsUnion())
		}
		if int64(cms) != 120000 {
			t.Errorf("ChunkTimeout = %d, want 120000", int64(cms))
		}
	})

	t.Run("false_variant", func(t *testing.T) {
		t.Parallel()
		var o opencode.ConfigProviderOptions
		if err := json.Unmarshal([]byte(`{"timeout":false,"headerTimeout":false,"chunkTimeout":false}`), &o); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		for name, got := range map[string]any{
			"timeout":       o.Timeout.AsUnion(),
			"headerTimeout": o.HeaderTimeout.AsUnion(),
			"chunkTimeout":  o.ChunkTimeout.AsUnion(),
		} {
			v, ok := got.(opencode.UnionBool)
			if !ok {
				t.Fatalf("%s runtime type = %T, want shared.UnionBool", name, got)
			}
			if bool(v) {
				t.Errorf("%s = true, want false", name)
			}
		}
	})

	t.Run("int64_precision_above_2pow53", func(t *testing.T) {
		t.Parallel()
		// float64 would round 9007199254740993 (2^53+1) down to ...992.
		var o opencode.ConfigProviderOptions
		if err := json.Unmarshal([]byte(`{"timeout":9007199254740993}`), &o); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		got, ok := o.Timeout.AsUnion().(opencode.UnionInt)
		if !ok {
			t.Fatalf("runtime type = %T, want shared.UnionInt", o.Timeout.AsUnion())
		}
		if int64(got) != 9007199254740993 {
			t.Errorf("Timeout = %d, want 9007199254740993 (int64 precision lost)", int64(got))
		}
	})

	t.Run("null_and_absent", func(t *testing.T) {
		t.Parallel()
		for _, body := range []string{`{"timeout":null}`, `{}`} {
			var o opencode.ConfigProviderOptions
			if err := json.Unmarshal([]byte(body), &o); err != nil {
				t.Fatalf("%s: json.Unmarshal: %v", body, err)
			}
			if o.Timeout.AsUnion() != nil {
				t.Errorf("%s: Timeout = %#v, want nil", body, o.Timeout.AsUnion())
			}
			if o.HeaderTimeout.AsUnion() != nil {
				t.Errorf("%s: HeaderTimeout = %#v, want nil", body, o.HeaderTimeout.AsUnion())
			}
			if o.ChunkTimeout.AsUnion() != nil {
				t.Errorf("%s: ChunkTimeout = %#v, want nil", body, o.ChunkTimeout.AsUnion())
			}
		}
	})

	t.Run("unexpected_shapes_degrade_to_nil", func(t *testing.T) {
		t.Parallel()
		for _, body := range []string{`{"timeout":"weird"}`, `{"timeout":[1]}`, `{"timeout":{"a":1}}`} {
			var o opencode.ConfigProviderOptions
			if err := json.Unmarshal([]byte(body), &o); err != nil {
				t.Fatalf("%s: must not fail the whole decode: %v", body, err)
			}
			if o.Timeout.AsUnion() != nil {
				t.Errorf("%s: Timeout = %#v, want nil", body, o.Timeout.AsUnion())
			}
		}
	})

	t.Run("param_marshal", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			param    opencode.ConfigProviderOptionsParam
			wantJSON string
		}{
			{opencode.ConfigProviderOptionsParam{}, `{}`},
			{opencode.ConfigProviderOptionsParam{
				Timeout: opencode.F[opencode.ConfigProviderOptionsTimeoutUnion](opencode.UnionInt(600000)),
			}, `{"timeout":600000}`},
			{opencode.ConfigProviderOptionsParam{
				HeaderTimeout: opencode.F[opencode.ConfigProviderOptionsTimeoutUnion](opencode.UnionBool(false)),
			}, `{"headerTimeout":false}`},
			{opencode.ConfigProviderOptionsParam{
				ChunkTimeout: opencode.F[opencode.ConfigProviderOptionsTimeoutUnion](opencode.UnionInt(300000)),
			}, `{"chunkTimeout":300000}`},
			{opencode.ConfigProviderOptionsParam{
				ChunkTimeout: opencode.F[opencode.ConfigProviderOptionsTimeoutUnion](opencode.UnionBool(false)),
			}, `{"chunkTimeout":false}`},
		}
		for _, tc := range cases {
			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			if string(got) != tc.wantJSON {
				t.Errorf("got %s, want %s", got, tc.wantJSON)
			}
		}
	})

	t.Run("response_value_reusable_as_request", func(t *testing.T) {
		t.Parallel()
		var src opencode.ConfigProviderOptions
		if err := json.Unmarshal([]byte(`{"timeout":5000,"headerTimeout":false,"chunkTimeout":false}`), &src); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		got, err := json.Marshal(opencode.ConfigProviderOptionsParam{
			Timeout:       opencode.F(src.Timeout.AsUnion()),
			HeaderTimeout: opencode.F(src.HeaderTimeout.AsUnion()),
			ChunkTimeout:  opencode.F(src.ChunkTimeout.AsUnion()),
		})
		if err != nil {
			t.Fatalf("json.Marshal: %v", err)
		}
		const want = `{"chunkTimeout":false,"headerTimeout":false,"timeout":5000}`
		if string(got) != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

// TestMcpRemoteConfigOAuthUnion verifies the OpenAPI
// `McpRemoteConfig.oauth` anyOf:
//
//	anyOf: [ {$ref: McpOAuthConfig}, {type: boolean, enum: [false]} ]
//
// The field is statically typed as the carrier [opencode.McpRemoteConfigOAuth];
// apijson resolves the registered union inside the carrier and the concrete
// variant is reached via [opencode.McpRemoteConfigOAuth.AsUnion].
//
// Run with: go test -run TestMcpRemoteConfigOAuthUnion -v ./...
func TestMcpRemoteConfigOAuthUnion(t *testing.T) {
	t.Parallel()

	t.Run("object_variant", func(t *testing.T) {
		t.Parallel()
		var v opencode.McpRemoteConfig
		if err := json.Unmarshal([]byte(`{"type":"remote","url":"u","oauth":{"clientId":"cid","callbackPort":4096}}`), &v); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		c, ok := v.OAuth.AsUnion().(opencode.McpOAuthConfig)
		if !ok {
			t.Fatalf("OAuth runtime type = %T, want opencode.McpOAuthConfig", v.OAuth.AsUnion())
		}
		if c.ClientID != "cid" {
			t.Errorf("ClientID = %q, want cid", c.ClientID)
		}
		if c.CallbackPort != 4096 {
			t.Errorf("CallbackPort = %d, want 4096", c.CallbackPort)
		}
	})

	t.Run("false_variant", func(t *testing.T) {
		t.Parallel()
		var v opencode.McpRemoteConfig
		if err := json.Unmarshal([]byte(`{"type":"remote","url":"u","oauth":false}`), &v); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		b, ok := v.OAuth.AsUnion().(opencode.UnionBool)
		if !ok {
			t.Fatalf("OAuth runtime type = %T, want shared.UnionBool", v.OAuth.AsUnion())
		}
		if bool(b) {
			t.Error("OAuth = true, want false")
		}
	})

	t.Run("null_and_absent", func(t *testing.T) {
		t.Parallel()
		for _, body := range []string{`{"type":"remote","url":"u","oauth":null}`, `{"type":"remote","url":"u"}`} {
			var v opencode.McpRemoteConfig
			if err := json.Unmarshal([]byte(body), &v); err != nil {
				t.Fatalf("%s: json.Unmarshal: %v", body, err)
			}
			if v.OAuth.AsUnion() != nil {
				t.Errorf("%s: OAuth = %#v, want nil", body, v.OAuth.AsUnion())
			}
		}
	})

	// ConfigMcp is the union carrier for Config.mcp values; apijson.Port must
	// carry the typed OAuth carrier value through to the carrier's any field.
	t.Run("through_ConfigMcp_carrier", func(t *testing.T) {
		t.Parallel()
		var cm opencode.ConfigMcp
		if err := json.Unmarshal([]byte(`{"type":"remote","url":"u","oauth":{"clientId":"c"}}`), &cm); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		o, ok := cm.OAuth.(opencode.McpRemoteConfigOAuth)
		if !ok {
			t.Fatalf("ConfigMcp.OAuth runtime type = %T, want opencode.McpRemoteConfigOAuth", cm.OAuth)
		}
		if _, ok := o.AsUnion().(opencode.McpOAuthConfig); !ok {
			t.Errorf("ConfigMcp.OAuth union runtime type = %T, want opencode.McpOAuthConfig", o.AsUnion())
		}
		rc, ok := cm.AsUnion().(opencode.McpRemoteConfig)
		if !ok {
			t.Fatalf("AsUnion() = %T, want opencode.McpRemoteConfig", cm.AsUnion())
		}
		if _, ok := rc.OAuth.AsUnion().(opencode.McpOAuthConfig); !ok {
			t.Errorf("McpRemoteConfig.OAuth union runtime type = %T", rc.OAuth.AsUnion())
		}
	})
}

func sortedTestKeys[T any](m map[string]T) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}

// Run with: go test -run TestConfigProviderModelParamInterleavedMarshal -v ./...
func TestConfigProviderModelParamInterleavedMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		param    opencode.ConfigProviderModelParam
		wantJSON string
	}{
		// Variant 1: bool
		{
			name: "bool_true",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedEnabled(true)),
			},
			wantJSON: `{"interleaved":true}`,
		},
		// Variant 1b: bool false
		{
			name: "bool_false",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedEnabled(false)),
			},
			wantJSON: `{"interleaved":false}`,
		},
		// Variant 2: enum string
		{
			name: "string_reasoning_text",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedString("reasoning_text")),
			},
			wantJSON: `{"interleaved":"reasoning_text"}`,
		},
		// Variant 3: open string
		{
			name: "string_vendor",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedString("vendor_custom")),
			},
			wantJSON: `{"interleaved":"vendor_custom"}`,
		},
		// Variant 4: object { "field": "reasoning_text" }
		// Uses the ConfigProviderModelsInterleavedFieldParam request-side type.
		{
			name: "object_field_reasoning_text",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedFieldParam{
					Field: opencode.F(opencode.ProviderModelCapabilitiesInterleavedFieldField("reasoning_text")),
				}),
			},
			wantJSON: `{"interleaved":{"field":"reasoning_text"}}`,
		},
		// Variant 4b: object with vendor-custom field value
		{
			name: "object_field_vendor_custom",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedFieldParam{
					Field: opencode.F(opencode.ProviderModelCapabilitiesInterleavedFieldField("vendor_custom")),
				}),
			},
			wantJSON: `{"interleaved":{"field":"vendor_custom"}}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			// Normalise to map for order-independent comparison
			var gotMap, wantMap map[string]any
			if err := json.Unmarshal(got, &gotMap); err != nil {
				t.Fatalf("re-parse got: %v", err)
			}
			if err := json.Unmarshal([]byte(tc.wantJSON), &wantMap); err != nil {
				t.Fatalf("re-parse want: %v", err)
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("wire JSON mismatch:\n  got:  %s\n  want: %s", got, tc.wantJSON)
			}
		})
	}
}

// Carrier structs with no promoted data fields (疑点 1): the empty-carrier
// PermissionRuleConfig / ConfigProviderOptionsTimeout must still populate the
// decoded union and preserve the full raw JSON of the value.
func TestPermissionRuleConfigEmptyCarrierUnionAndRaw(t *testing.T) {
	t.Parallel()
	var p opencode.PermissionConfigObject
	if err := json.Unmarshal([]byte(`{"read":"allow","edit":{"src/**":"ask"}}`), &p); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	act, ok := p.Read.AsUnion().(opencode.PermissionActionConfig)
	if !ok {
		t.Fatalf("Read union runtime type = %T, want opencode.PermissionActionConfig", p.Read.AsUnion())
	}
	if act != opencode.PermissionActionConfigAllow {
		t.Errorf("Read = %q, want %q", act, opencode.PermissionActionConfigAllow)
	}
	if got := p.Read.JSON.RawJSON(); got != `"allow"` {
		t.Errorf("Read.JSON.RawJSON() = %q, want %q", got, `"allow"`)
	}
	obj, ok := p.Edit.AsUnion().(opencode.PermissionObjectConfig)
	if !ok {
		t.Fatalf("Edit union runtime type = %T, want opencode.PermissionObjectConfig", p.Edit.AsUnion())
	}
	if obj["src/**"] != opencode.PermissionActionConfigAsk {
		t.Errorf(`Edit["src/**"] = %q, want %q`, obj["src/**"], opencode.PermissionActionConfigAsk)
	}
	if got := p.Edit.JSON.RawJSON(); got != `{"src/**":"ask"}` {
		t.Errorf("Edit.JSON.RawJSON() = %q, want %q", got, `{"src/**":"ask"}`)
	}
}

func TestConfigProviderOptionsTimeoutEmptyCarrierUnionAndRaw(t *testing.T) {
	t.Parallel()
	var o opencode.ConfigProviderOptions
	if err := json.Unmarshal([]byte(`{"timeout":600000,"headerTimeout":false}`), &o); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	ms, ok := o.Timeout.AsUnion().(opencode.UnionInt)
	if !ok {
		t.Fatalf("Timeout union runtime type = %T, want opencode.UnionInt", o.Timeout.AsUnion())
	}
	if int64(ms) != 600000 {
		t.Errorf("Timeout = %d, want 600000", int64(ms))
	}
	if got := o.Timeout.JSON.RawJSON(); got != `600000` {
		t.Errorf("Timeout.JSON.RawJSON() = %q, want %q", got, `600000`)
	}
	b, ok := o.HeaderTimeout.AsUnion().(opencode.UnionBool)
	if !ok {
		t.Fatalf("HeaderTimeout union runtime type = %T, want opencode.UnionBool", o.HeaderTimeout.AsUnion())
	}
	if bool(b) {
		t.Error("HeaderTimeout = true, want false")
	}
	if got := o.HeaderTimeout.JSON.RawJSON(); got != `false` {
		t.Errorf("HeaderTimeout.JSON.RawJSON() = %q, want %q", got, `false`)
	}
	// OpenAPI declares timeout as anyOf[integer>0, enum[false]]: true must NOT
	// silently decode to UnionBool.
	var bad opencode.ConfigProviderOptions
	if err := json.Unmarshal([]byte(`{"timeout":true}`), &bad); err == nil {
		if _, ok := bad.Timeout.AsUnion().(opencode.UnionBool); ok {
			t.Error("timeout:true silently decoded as UnionBool, want decode error")
		}
	}
}

// Carrier structs with promoted data fields (疑点 1, golden ToolPartState
// pattern): McpRemoteConfigOAuth promotes the McpOAuthConfig object variant
// fields while keeping raw JSON intact for both variants.
func TestMcpRemoteConfigOAuthCarrierPromotedFields(t *testing.T) {
	t.Parallel()
	var v opencode.McpRemoteConfig
	if err := json.Unmarshal([]byte(`{"type":"remote","url":"u","oauth":{"clientId":"cid","clientSecret":"sec","scope":"openid","callbackPort":4096,"redirectUri":"http://x"}}`), &v); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if got := v.OAuth.ClientID; got != "cid" {
		t.Errorf("OAuth.ClientID = %q, want cid", got)
	}
	if got := v.OAuth.ClientSecret; got != "sec" {
		t.Errorf("OAuth.ClientSecret = %q, want sec", got)
	}
	if got := v.OAuth.Scope; got != "openid" {
		t.Errorf("OAuth.Scope = %q, want openid", got)
	}
	if got := v.OAuth.CallbackPort; got != 4096 {
		t.Errorf("OAuth.CallbackPort = %d, want 4096", got)
	}
	if got := v.OAuth.RedirectURI; got != "http://x" {
		t.Errorf("OAuth.RedirectURI = %q, want http://x", got)
	}
	if got := v.OAuth.JSON.RawJSON(); got != `{"clientId":"cid","clientSecret":"sec","scope":"openid","callbackPort":4096,"redirectUri":"http://x"}` {
		t.Errorf("OAuth.JSON.RawJSON() = %q", got)
	}

	var v2 opencode.McpRemoteConfig
	if err := json.Unmarshal([]byte(`{"type":"remote","url":"u","oauth":false}`), &v2); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	b, ok := v2.OAuth.AsUnion().(opencode.UnionBool)
	if !ok {
		t.Fatalf("OAuth union runtime type = %T, want opencode.UnionBool", v2.OAuth.AsUnion())
	}
	if bool(b) {
		t.Error("OAuth = true, want false")
	}
	if got := v2.OAuth.JSON.RawJSON(); got != `false` {
		t.Errorf("OAuth.JSON.RawJSON() = %q, want %q", got, `false`)
	}
}

func TestConfigV2ReferenceCarrierVariants(t *testing.T) {
	t.Parallel()
	var c opencode.Config
	if err := json.Unmarshal([]byte(`{
		"reference": {
			"plain": "owner/repo",
			"git": {"repository": "https://github.com/a/b", "branch": "dev", "description": "g", "hidden": true},
			"local": {"path": "./docs", "description": "d"}
		}
	}`), &c); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	plain, ok := c.Reference["plain"].AsUnion().(opencode.ConfigV2ReferenceString)
	if !ok {
		t.Fatalf("plain union runtime type = %T, want opencode.ConfigV2ReferenceString", c.Reference["plain"].AsUnion())
	}
	if string(plain) != "owner/repo" {
		t.Errorf("plain = %q, want owner/repo", string(plain))
	}
	if got := c.Reference["plain"].JSON.RawJSON(); got != `"owner/repo"` {
		t.Errorf("plain.JSON.RawJSON() = %q, want %q", got, `"owner/repo"`)
	}

	git, ok := c.Reference["git"].AsUnion().(opencode.ConfigV2ReferenceGit)
	if !ok {
		t.Fatalf("git union runtime type = %T, want opencode.ConfigV2ReferenceGit", c.Reference["git"].AsUnion())
	}
	if git.Repository != "https://github.com/a/b" || git.Branch != "dev" || !git.Hidden || git.Description != "g" {
		t.Errorf("git variant = %#v", git)
	}
	if got := c.Reference["git"].Repository; got != "https://github.com/a/b" {
		t.Errorf("promoted Repository = %q, want https://github.com/a/b", got)
	}
	if got := c.Reference["git"].Branch; got != "dev" {
		t.Errorf("promoted Branch = %q, want dev", got)
	}
	if got := c.Reference["git"].JSON.RawJSON(); got != `{"repository": "https://github.com/a/b", "branch": "dev", "description": "g", "hidden": true}` {
		t.Errorf("git.JSON.RawJSON() = %q", got)
	}

	local, ok := c.Reference["local"].AsUnion().(opencode.ConfigV2ReferenceLocal)
	if !ok {
		t.Fatalf("local union runtime type = %T, want opencode.ConfigV2ReferenceLocal", c.Reference["local"].AsUnion())
	}
	if local.Path != "./docs" || local.Description != "d" {
		t.Errorf("local variant = %#v", local)
	}
	if got := c.Reference["local"].Path; got != "./docs" {
		t.Errorf("promoted Path = %q, want ./docs", got)
	}
	if got := c.Reference["local"].JSON.RawJSON(); got != `{"path": "./docs", "description": "d"}` {
		t.Errorf("local.JSON.RawJSON() = %q", got)
	}
}

// 疑点 2: apijson.Port panics on non-struct union variants, which is why the
// scalar-variant carriers skip Port and only set raw.
func TestPortPanicsOnNonStructVariant(t *testing.T) {
	t.Parallel()
	for _, scalar := range []any{
		opencode.PermissionActionConfigAllow,
		opencode.PermissionObjectConfig{"a/**": opencode.PermissionActionConfigAllow},
		opencode.ConfigV2ReferenceString("owner/repo"),
		opencode.UnionBool(false),
		opencode.UnionInt(600000),
	} {
		t.Run(fmt.Sprintf("%T", scalar), func(t *testing.T) {
			t.Parallel()
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("apijson.Port(%T, ...) did not panic", scalar)
				}
			}()
			_ = apijson.Port(scalar, &struct{}{})
		})
	}
}

// TestConfigAutoupdateUnionParamMarshal verifies serialization of every variant of
// the OpenAPI `Config.autoupdate` anyOf (boolean | "notify").
func TestConfigAutoupdateUnionParamMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		param    opencode.ConfigUpdateParams
		wantJSON string
	}{
		// Variant 1a: boolean true
		{
			name: "bool_true",
			param: opencode.ConfigUpdateParams{
				Autoupdate: opencode.F[opencode.ConfigAutoupdateUnionParam](opencode.ConfigAutoupdateEnabled(true)),
			},
			wantJSON: `{"autoupdate":true}`,
		},
		// Variant 1b: boolean false
		{
			name: "bool_false",
			param: opencode.ConfigUpdateParams{
				Autoupdate: opencode.F[opencode.ConfigAutoupdateUnionParam](opencode.ConfigAutoupdateEnabled(false)),
			},
			wantJSON: `{"autoupdate":false}`,
		},
		// Variant 2: string enum "notify"
		{
			name: "string_notify",
			param: opencode.ConfigUpdateParams{
				Autoupdate: opencode.F[opencode.ConfigAutoupdateUnionParam](opencode.ConfigAutoupdateNotify("notify")),
			},
			wantJSON: `{"autoupdate":"notify"}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			var gotMap, wantMap map[string]any
			if err := json.Unmarshal(got, &gotMap); err != nil {
				t.Fatalf("re-parse got: %v", err)
			}
			if err := json.Unmarshal([]byte(tc.wantJSON), &wantMap); err != nil {
				t.Fatalf("re-parse want: %v", err)
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("wire JSON mismatch:\n  got:  %s\n  want: %s", got, tc.wantJSON)
			}
		})
	}
}

// TestConfigFormatterUnionParamMarshal verifies serialization of every variant of
// the OpenAPI `Config.formatter` anyOf (boolean | map[string]ConfigFormatter).
func TestConfigFormatterUnionParamMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		param    opencode.ConfigUpdateParams
		wantJSON string
	}{
		// Variant 1: boolean false (disable)
		{
			name: "bool_false",
			param: opencode.ConfigUpdateParams{
				Formatter: opencode.F[opencode.ConfigFormatterUnionParam](opencode.ConfigFormatterEnabled(false)),
			},
			wantJSON: `{"formatter":false}`,
		},
		// Variant 2: object map
		{
			name: "map_with_command_and_extensions",
			param: opencode.ConfigUpdateParams{
				Formatter: opencode.F[opencode.ConfigFormatterUnionParam](opencode.ConfigFormatterMapParam{
					"prettier": opencode.ConfigFormatterParam{
						Command:    opencode.F([]string{"prettier", "--write"}),
						Extensions: opencode.F([]string{".ts"}),
					},
				}),
			},
			wantJSON: `{"formatter":{"prettier":{"command":["prettier","--write"],"extensions":[".ts"]}}}`,
		},
		// Variant 2b: map with environment and disabled
		{
			name: "map_environment_disabled",
			param: opencode.ConfigUpdateParams{
				Formatter: opencode.F[opencode.ConfigFormatterUnionParam](opencode.ConfigFormatterMapParam{
					"biome": opencode.ConfigFormatterParam{
						Disabled:    opencode.F(true),
						Environment: opencode.F(map[string]string{"NODE_ENV": "production"}),
					},
				}),
			},
			wantJSON: `{"formatter":{"biome":{"disabled":true,"environment":{"NODE_ENV":"production"}}}}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			var gotMap, wantMap map[string]any
			if err := json.Unmarshal(got, &gotMap); err != nil {
				t.Fatalf("re-parse got: %v", err)
			}
			if err := json.Unmarshal([]byte(tc.wantJSON), &wantMap); err != nil {
				t.Fatalf("re-parse want: %v", err)
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("wire JSON mismatch:\n  got:  %s\n  want: %s", got, tc.wantJSON)
			}
		})
	}
}

// TestConfigLspUnionParamMarshal verifies serialization of every variant of the
// OpenAPI `Config.lsp` anyOf (boolean | map[string] per-server union).
func TestConfigLspUnionParamMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		param    opencode.ConfigUpdateParams
		wantJSON string
	}{
		// Variant 1: boolean false (disable)
		{
			name: "bool_false",
			param: opencode.ConfigUpdateParams{
				Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspEnabled(false)),
			},
			wantJSON: `{"lsp":false}`,
		},
		// Variant 2a: per-server disabled config
		{
			name: "map_server_disabled",
			param: opencode.ConfigUpdateParams{
				Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspMapParam{
					"gopls": opencode.ConfigLspDisabledParam{
						Disabled: opencode.F(opencode.ConfigLspDisabledDisabled(true)),
					},
				}),
			},
			wantJSON: `{"lsp":{"gopls":{"disabled":true}}}`,
		},
		// Variant 2b: per-server command-based config
		{
			name: "map_server_command",
			param: opencode.ConfigUpdateParams{
				Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspMapParam{
					"gopls": opencode.ConfigLspObjectParam{
						Command:    opencode.F([]string{"gopls"}),
						Extensions: opencode.F([]string{".go"}),
					},
				}),
			},
			wantJSON: `{"lsp":{"gopls":{"command":["gopls"],"extensions":[".go"]}}}`,
		},
		// Variant 2c: per-server with env and initialization
		{
			name: "map_server_env_initialization",
			param: opencode.ConfigUpdateParams{
				Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspMapParam{
					"gopls": opencode.ConfigLspObjectParam{
						Command:        opencode.F([]string{"gopls"}),
						Env:            opencode.F(map[string]string{"GOPLS_GENERATE": "1"}),
						Initialization: opencode.F(map[string]any{"formatting": true}),
					},
				}),
			},
			wantJSON: `{"lsp":{"gopls":{"command":["gopls"],"env":{"GOPLS_GENERATE":"1"},"initialization":{"formatting":true}}}}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			var gotMap, wantMap map[string]any
			if err := json.Unmarshal(got, &gotMap); err != nil {
				t.Fatalf("re-parse got: %v", err)
			}
			if err := json.Unmarshal([]byte(tc.wantJSON), &wantMap); err != nil {
				t.Fatalf("re-parse want: %v", err)
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("wire JSON mismatch:\n  got:  %s\n  want: %s", got, tc.wantJSON)
			}
		})
	}
}

// TestConfigPluginItemUnionParamMarshal verifies serialization of every variant of
// the OpenAPI `Config.plugin` array items (string | [string, object]).
func TestConfigPluginItemUnionParamMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		param    opencode.ConfigUpdateParams
		wantJSON string
	}{
		// Variant 1: single bare string
		{
			name: "single_name",
			param: opencode.ConfigUpdateParams{
				Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
					opencode.ConfigPluginName("a"),
				}),
			},
			wantJSON: `{"plugin":["a"]}`,
		},
		// Variant 2: single 2-tuple [name, config]
		{
			name: "single_tuple",
			param: opencode.ConfigUpdateParams{
				Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
					opencode.ConfigPluginTupleParam{
						Name:   opencode.F("a"),
						Config: opencode.F(map[string]any{"k": "v"}),
					},
				}),
			},
			wantJSON: `{"plugin":[["a",{"k":"v"}]]}`,
		},
		// Variant 3: mixed string and tuple
		{
			name: "mixed",
			param: opencode.ConfigUpdateParams{
				Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
					opencode.ConfigPluginName("b"),
					opencode.ConfigPluginTupleParam{
						Name:   opencode.F("a"),
						Config: opencode.F(map[string]any{"k": "v"}),
					},
				}),
			},
			wantJSON: `{"plugin":["b",["a",{"k":"v"}]]}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			var gotMap, wantMap map[string]any
			if err := json.Unmarshal(got, &gotMap); err != nil {
				t.Fatalf("re-parse got: %v", err)
			}
			if err := json.Unmarshal([]byte(tc.wantJSON), &wantMap); err != nil {
				t.Fatalf("re-parse want: %v", err)
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("wire JSON mismatch:\n  got:  %s\n  want: %s", got, tc.wantJSON)
			}
		})
	}
}

// TestConfigV2ReferenceUnionParamMarshal verifies serialization of every
// variant of the OpenAPI `Config.reference` / `Config.references` map values
// (string | ConfigV2ReferenceGit | ConfigV2ReferenceLocal).
func TestConfigV2ReferenceUnionParamMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		param    opencode.ConfigUpdateParams
		wantJSON string
	}{
		// Variant 1: bare string
		{
			name: "reference_string",
			param: opencode.ConfigUpdateParams{
				Reference: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
					"x": opencode.ConfigV2ReferenceString("https://example.com/docs"),
				}),
			},
			wantJSON: `{"reference":{"x":"https://example.com/docs"}}`,
		},
		// Variant 2: git reference
		{
			name: "reference_git",
			param: opencode.ConfigUpdateParams{
				Reference: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
					"docs": opencode.ConfigV2ReferenceGitParam{
						Repository: opencode.F("owner/repo"),
						Branch:     opencode.F("main"),
					},
				}),
			},
			wantJSON: `{"reference":{"docs":{"repository":"owner/repo","branch":"main"}}}`,
		},
		// Variant 3: local reference
		{
			name: "references_local",
			param: opencode.ConfigUpdateParams{
				References: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
					"local": opencode.ConfigV2ReferenceLocalParam{
						Path: opencode.F("./docs"),
					},
				}),
			},
			wantJSON: `{"references":{"local":{"path":"./docs"}}}`,
		},
		// Variant 4: mixed string and git in the same map
		{
			name: "mixed",
			param: opencode.ConfigUpdateParams{
				Reference: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
					"plain": opencode.ConfigV2ReferenceString("owner/repo"),
					"git": opencode.ConfigV2ReferenceGitParam{
						Repository: opencode.F("owner/repo"),
					},
				}),
			},
			wantJSON: `{"reference":{"git":{"repository":"owner/repo"},"plain":"owner/repo"}}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}
			var gotMap, wantMap map[string]any
			if err := json.Unmarshal(got, &gotMap); err != nil {
				t.Fatalf("re-parse got: %v", err)
			}
			if err := json.Unmarshal([]byte(tc.wantJSON), &wantMap); err != nil {
				t.Fatalf("re-parse want: %v", err)
			}
			if !reflect.DeepEqual(gotMap, wantMap) {
				t.Errorf("wire JSON mismatch:\n  got:  %s\n  want: %s", got, tc.wantJSON)
			}
		})
	}
}

// TestConfigUpdateParamsUnionFieldsOmitted verifies that union-typed fields
// that are not set do not appear in the marshaled JSON (param.Field zero-value
// semantics), and that a fully-populated ConfigUpdateParams marshals all the
// union fields together.
func TestConfigUpdateParamsUnionFieldsOmitted(t *testing.T) {
	t.Parallel()
	// No union fields set → nothing emitted.
	empty := opencode.ConfigUpdateParams{}
	gotEmpty, err := json.Marshal(empty)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	if string(gotEmpty) != "{}" {
		t.Errorf("empty ConfigUpdateParams = %s, want {}", gotEmpty)
	}

	// All union fields set together.
	full := opencode.ConfigUpdateParams{
		Autoupdate: opencode.F[opencode.ConfigAutoupdateUnionParam](opencode.ConfigAutoupdateNotify("notify")),
		Formatter:  opencode.F[opencode.ConfigFormatterUnionParam](opencode.ConfigFormatterEnabled(true)),
		Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspMapParam{
			"gopls": opencode.ConfigLspDisabledParam{
				Disabled: opencode.F(opencode.ConfigLspDisabledDisabled(true)),
			},
		}),
		Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
			opencode.ConfigPluginName("a"),
			opencode.ConfigPluginTupleParam{
				Name:   opencode.F("b"),
				Config: opencode.F(map[string]any{"n": 1}),
			},
		}),
		Reference: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
			"x": opencode.ConfigV2ReferenceString("str"),
		}),
		References: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
			"git": opencode.ConfigV2ReferenceGitParam{
				Repository: opencode.F("owner/repo"),
			},
		}),
	}
	gotFull, err := json.Marshal(full)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	var gotFullMap map[string]any
	if err := json.Unmarshal(gotFull, &gotFullMap); err != nil {
		t.Fatalf("re-parse: %v", err)
	}
	for _, key := range []string{"autoupdate", "formatter", "lsp", "plugin", "reference", "references"} {
		if _, ok := gotFullMap[key]; !ok {
			t.Errorf("marshaled JSON missing %q: %s", key, gotFull)
		}
	}
}
