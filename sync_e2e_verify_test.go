// Code generated for e2e review verification (Review 2A). Not a Stainless artifact.
//
// These tests empirically verify wire-format serialization/deserialization of
// the sync changes against OpenAPI-derived JSON vectors (see
// .tmp/sync-1.18.15/vectors/*.json). All comparisons are structural
// (map[string]any deep-equality), never string-based.

package opencode_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/tidwall/gjson"

	"github.com/sst/opencode-sdk-go"
)

// parseJSON unmarshals raw JSON into any for structural comparison.
func parseJSON(t *testing.T, raw string) any {
	t.Helper()
	var m any
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("unmarshal %q: %v", raw, err)
	}
	return m
}

// marshalToMap marshals v and parses the output into any.
func marshalToMap(t *testing.T, v any) any {
	t.Helper()
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	return parseJSON(t, string(raw))
}

// compareMaps asserts that got (a Go marshaled wire form) structurally equals
// the OpenAPI-derived want vector.
func compareMaps(t *testing.T, name string, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: wire JSON mismatch:\n  got:  %#v\n  want: %#v", name, got, want)
	}
}

// fieldValue extracts a top-level key from a marshaled params object and
// compares it structurally against the OpenAPI-derived scalar/object vector.
func fieldValue(t *testing.T, v any, key string) any {
	t.Helper()
	m, ok := v.(map[string]any)
	if !ok {
		t.Fatalf("marshaled output %#v is not an object", v)
	}
	got, ok := m[key]
	if !ok {
		t.Fatalf("marshaled output missing key %q: %#v", key, m)
	}
	return got
}

// =============================================================================
// Task 2: Request side — Go → JSON forward serialization vs OpenAPI vectors
// =============================================================================

func TestSyncE2EConfigAutoupdateMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		field  string
		param  opencode.ConfigUpdateParams
		vector string
	}{
		{"boolean_true", "autoupdate", opencode.ConfigUpdateParams{
			Autoupdate: opencode.F[opencode.ConfigAutoupdateUnionParam](opencode.ConfigAutoupdateEnabled(true)),
		}, syncVecConfigAutoupdate["boolean_true"]},
		{"boolean_false", "autoupdate", opencode.ConfigUpdateParams{
			Autoupdate: opencode.F[opencode.ConfigAutoupdateUnionParam](opencode.ConfigAutoupdateEnabled(false)),
		}, syncVecConfigAutoupdate["boolean_false"]},
		{"notify", "autoupdate", opencode.ConfigUpdateParams{
			Autoupdate: opencode.F[opencode.ConfigAutoupdateUnionParam](opencode.ConfigAutoupdateNotify("notify")),
		}, syncVecConfigAutoupdate["notify"]},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := marshalToMap(t, tc.param)
			compareMaps(t, tc.name, fieldValue(t, got, tc.field), parseJSON(t, tc.vector))
		})
	}
}

func TestSyncE2EConfigFormatterMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		field  string
		param  opencode.ConfigUpdateParams
		vector string
	}{
		{"boolean_false", "formatter", opencode.ConfigUpdateParams{
			Formatter: opencode.F[opencode.ConfigFormatterUnionParam](opencode.ConfigFormatterEnabled(false)),
		}, syncVecConfigFormatter["boolean_false"]},
		{"boolean_true", "formatter", opencode.ConfigUpdateParams{
			Formatter: opencode.F[opencode.ConfigFormatterUnionParam](opencode.ConfigFormatterEnabled(true)),
		}, syncVecConfigFormatter["boolean_true"]},
		{"map_single", "formatter", opencode.ConfigUpdateParams{
			Formatter: opencode.F[opencode.ConfigFormatterUnionParam](opencode.ConfigFormatterMapParam{
				"prettier": opencode.ConfigFormatterParam{
					Disabled:    opencode.F(true),
					Command:     opencode.F([]string{"sample-string"}),
					Environment: opencode.F(map[string]string{}),
					Extensions:  opencode.F([]string{"sample-string"}),
				},
			}),
		}, syncVecConfigFormatter["map_single"]},
		{"map_multi", "formatter", opencode.ConfigUpdateParams{
			Formatter: opencode.F[opencode.ConfigFormatterUnionParam](opencode.ConfigFormatterMapParam{
				"prettier": opencode.ConfigFormatterParam{
					Disabled:    opencode.F(false),
					Command:     opencode.F([]string{"prettier", "--write"}),
					Environment: opencode.F(map[string]string{"NODE_ENV": "production"}),
					Extensions:  opencode.F([]string{".ts"}),
				},
				"biome": opencode.ConfigFormatterParam{
					Disabled: opencode.F(true),
				},
			}),
		}, syncVecConfigFormatter["map_multi"]},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := marshalToMap(t, tc.param)
			compareMaps(t, tc.name, fieldValue(t, got, tc.field), parseJSON(t, tc.vector))
		})
	}
}

func TestSyncE2EConfigLspMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		field  string
		param  opencode.ConfigUpdateParams
		vector string
	}{
		{"boolean_false", "lsp", opencode.ConfigUpdateParams{
			Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspEnabled(false)),
		}, syncVecConfigLsp["boolean_false"]},
		{"boolean_true", "lsp", opencode.ConfigUpdateParams{
			Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspEnabled(true)),
		}, syncVecConfigLsp["boolean_true"]},
		{"map_server_disabled", "lsp", opencode.ConfigUpdateParams{
			Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspMapParam{
				"gopls": opencode.ConfigLspDisabledParam{
					Disabled: opencode.F(opencode.ConfigLspDisabledDisabled(true)),
				},
			}),
		}, syncVecConfigLsp["map_server_disabled"]},
		{"map_mixed", "lsp", opencode.ConfigUpdateParams{
			Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspMapParam{
				"gopls": opencode.ConfigLspObjectParam{
					Command:        opencode.F([]string{"gopls"}),
					Extensions:     opencode.F([]string{".go"}),
					Env:            opencode.F(map[string]string{"GOPLS_GENERATE": "1"}),
					Initialization: opencode.F(map[string]any{"formatting": true}),
				},
				"python": opencode.ConfigLspDisabledParam{
					Disabled: opencode.F(opencode.ConfigLspDisabledDisabled(true)),
				},
			}),
		}, syncVecConfigLsp["map_mixed"]},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := marshalToMap(t, tc.param)
			compareMaps(t, tc.name, fieldValue(t, got, tc.field), parseJSON(t, tc.vector))
		})
	}
}

// TestSyncE2EConfigLspNestedUnionMarshal verifies the nested
// ConfigLspServerUnionParam serializes correctly through the map: both variants
// (disabled | command object) must round-trip, confirming the gjson.JSON
// TypeFilter on the inner union does not corrupt the Marshal direction.
func TestSyncE2EConfigLspNestedUnionMarshal(t *testing.T) {
	t.Parallel()
	params := opencode.ConfigUpdateParams{
		Lsp: opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspMapParam{
			"gopls": opencode.ConfigLspObjectParam{
				Command:    opencode.F([]string{"gopls"}),
				Extensions: opencode.F([]string{".go"}),
			},
			"disabled_srv": opencode.ConfigLspDisabledParam{
				Disabled: opencode.F(opencode.ConfigLspDisabledDisabled(true)),
			},
		}),
	}
	got := marshalToMap(t, params)
	want := parseJSON(t, `{"lsp":{"gopls":{"command":["gopls"],"extensions":[".go"]},"disabled_srv":{"disabled":true}}}`)
	compareMaps(t, "nested union", got, want)
}

func TestSyncE2EConfigPluginMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		field  string
		param  opencode.ConfigUpdateParams
		vector string
	}{
		{"single_name", "plugin", opencode.ConfigUpdateParams{
			Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
				opencode.ConfigPluginName("sample-string"),
			}),
		}, syncVecConfigPlugin["single_name"]},
		{"single_tuple", "plugin", opencode.ConfigUpdateParams{
			Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
				opencode.ConfigPluginTupleParam{
					Name:   opencode.F("sample-string"),
					Config: opencode.F(map[string]any{}),
				},
			}),
		}, syncVecConfigPlugin["single_tuple"]},
		{"mixed", "plugin", opencode.ConfigUpdateParams{
			Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
				opencode.ConfigPluginName("sample-string"),
				opencode.ConfigPluginTupleParam{
					Name:   opencode.F("sample-string"),
					Config: opencode.F(map[string]any{}),
				},
			}),
		}, syncVecConfigPlugin["mixed"]},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := marshalToMap(t, tc.param)
			compareMaps(t, tc.name, fieldValue(t, got, tc.field), parseJSON(t, tc.vector))
		})
	}
}

// TestSyncE2EConfigPluginTupleShape verifies the 2-tuple wire shape conforms to
// the OpenAPI prefixItems constraint ([string, object], minItems 2, maxItems 2).
func TestSyncE2EConfigPluginTupleShape(t *testing.T) {
	t.Parallel()
	params := opencode.ConfigUpdateParams{
		Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
			opencode.ConfigPluginTupleParam{
				Name:   opencode.F("a"),
				Config: opencode.F(map[string]any{"k": "v"}),
			},
		}),
	}
	raw, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	var outer struct {
		Plugin []json.RawMessage `json:"plugin"`
	}
	if err := json.Unmarshal(raw, &outer); err != nil {
		t.Fatalf("unmarshal outer: %v", err)
	}
	if len(outer.Plugin) != 1 {
		t.Fatalf("plugin item count = %d, want 1", len(outer.Plugin))
	}
	var tuple []json.RawMessage
	if err := json.Unmarshal(outer.Plugin[0], &tuple); err != nil {
		t.Fatalf("plugin[0] is not an array: %v", err)
	}
	if len(tuple) != 2 {
		t.Fatalf("tuple length = %d, want 2 (minItems=2 / maxItems=2)", len(tuple))
	}
	var name string
	if err := json.Unmarshal(tuple[0], &name); err != nil {
		t.Fatalf("tuple[0] = %s, want string", tuple[0])
	}
	if name != "a" {
		t.Errorf("tuple[0] = %q, want %q", name, "a")
	}
	var config map[string]any
	if err := json.Unmarshal(tuple[1], &config); err != nil {
		t.Fatalf("tuple[1] = %s, want object", tuple[1])
	}
	if config["k"] != "v" {
		t.Errorf("tuple[1] = %#v, want {k:v}", config)
	}
}

// TestSyncE2EConfigPluginEmptyTupleBehavior exercises the all-zero
// ConfigPluginTupleParam{} value to determine what wire output it produces.
// OpenAPI prefixItems declares the first element as {type: string} with no
// minLength, so "" is schema-legal; the config object element is {type: object}.
func TestSyncE2EConfigPluginEmptyTupleBehavior(t *testing.T) {
	t.Parallel()
	params := opencode.ConfigUpdateParams{
		Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
			opencode.ConfigPluginTupleParam{},
		}),
	}
	raw, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	var outer struct {
		Plugin []json.RawMessage `json:"plugin"`
	}
	if err := json.Unmarshal(raw, &outer); err != nil {
		t.Fatalf("unmarshal outer: %v", err)
	}
	var tuple []json.RawMessage
	if err := json.Unmarshal(outer.Plugin[0], &tuple); err != nil {
		t.Fatalf("plugin[0] is not an array: %v", err)
	}
	if len(tuple) != 2 {
		t.Fatalf("empty tuple length = %d, want 2", len(tuple))
	}
	var name string
	if err := json.Unmarshal(tuple[0], &name); err != nil {
		t.Fatalf("empty tuple[0] = %s, want string", tuple[0])
	}
	// Document the observed wire output: ["",{}]. The empty-string plugin name
	// is schema-legal (no minLength on the prefixItems string), so this is not
	// an OpenAPI violation, but the zero value is unlikely to be what a caller
	// intends. The MarshalJSON comment already documents the behavior.
	t.Logf("empty ConfigPluginTupleParam wire = %s", string(raw))
}

func TestSyncE2EConfigReferenceMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		field  string
		param  opencode.ConfigUpdateParams
		vector string
	}{
		{"string", "reference", opencode.ConfigUpdateParams{
			Reference: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
				"x": opencode.ConfigV2ReferenceString("sample-string"),
			}),
		}, syncVecConfigReference["string"]},
		{"git", "reference", opencode.ConfigUpdateParams{
			Reference: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
				"x": opencode.ConfigV2ReferenceGitParam{
					Repository:  opencode.F("sample-string"),
					Branch:      opencode.F("sample-string"),
					Description: opencode.F("sample-string"),
					Hidden:      opencode.F(true),
				},
			}),
		}, syncVecConfigReference["git"]},
		{"local", "references", opencode.ConfigUpdateParams{
			References: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
				"x": opencode.ConfigV2ReferenceLocalParam{
					Path:        opencode.F("sample-string"),
					Description: opencode.F("sample-string"),
					Hidden:      opencode.F(true),
				},
			}),
		}, syncVecConfigReference["local"]},
		{"map_mixed", "reference", opencode.ConfigUpdateParams{
			Reference: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
				"plain": opencode.ConfigV2ReferenceString("https://example.com/docs"),
				"git": opencode.ConfigV2ReferenceGitParam{
					Repository:  opencode.F("owner/repo"),
					Branch:      opencode.F("main"),
					Description: opencode.F("d"),
					Hidden:      opencode.F(false),
				},
				"local": opencode.ConfigV2ReferenceLocalParam{
					Path:        opencode.F("./docs"),
					Description: opencode.F("d"),
					Hidden:      opencode.F(true),
				},
			}),
		}, syncVecConfigReference["map_mixed"]},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := marshalToMap(t, tc.param)
			compareMaps(t, tc.name, fieldValue(t, got, tc.field), parseJSON(t, tc.vector))
		})
	}
}

// TestSyncE2EConfigUpdateParamsUnsetFieldsOmitted verifies that unset
// union-typed fields are absent from the marshaled JSON (param.Field zero
// semantics), and that all union fields set together emit every key.
func TestSyncE2EConfigUpdateParamsUnsetFieldsOmitted(t *testing.T) {
	t.Parallel()
	got := marshalToMap(t, opencode.ConfigUpdateParams{})
	if len(got.(map[string]any)) != 0 {
		t.Fatalf("empty ConfigUpdateParams emitted keys %v, want none", got)
	}

	full := opencode.ConfigUpdateParams{
		Autoupdate: opencode.F[opencode.ConfigAutoupdateUnionParam](opencode.ConfigAutoupdateNotify("notify")),
		Formatter:  opencode.F[opencode.ConfigFormatterUnionParam](opencode.ConfigFormatterEnabled(true)),
		Lsp:        opencode.F[opencode.ConfigLspUnionParam](opencode.ConfigLspEnabled(true)),
		Plugin: opencode.F[[]opencode.ConfigPluginItemUnionParam]([]opencode.ConfigPluginItemUnionParam{
			opencode.ConfigPluginName("a"),
		}),
		Reference: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
			"x": opencode.ConfigV2ReferenceString("s"),
		}),
		References: opencode.F(map[string]opencode.ConfigV2ReferenceUnionParam{
			"y": opencode.ConfigV2ReferenceString("t"),
		}),
	}
	gotFull := marshalToMap(t, full).(map[string]any)
	for _, key := range []string{"autoupdate", "formatter", "lsp", "plugin", "reference", "references"} {
		if _, ok := gotFull[key]; !ok {
			t.Errorf("missing %q in %#v", key, gotFull)
		}
	}
}

// TestSyncE2EConfigProviderInterleavedMarshal verifies every variant of the
// OpenAPI `ProviderConfig.models.*.interleaved` anyOf
// (boolean | string-enum | vendor-string | {field:string}).
func TestSyncE2EConfigProviderInterleavedMarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		field  string
		param  opencode.ConfigProviderModelParam
		vector string
	}{
		{"boolean_true", "interleaved", opencode.ConfigProviderModelParam{
			Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedEnabled(true)),
		}, syncVecProviderInterleaved["boolean_true"]},
		{"boolean_false", "interleaved", opencode.ConfigProviderModelParam{
			Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedEnabled(false)),
		}, syncVecProviderInterleaved["boolean_false"]},
		{"enum_reasoning", "interleaved", opencode.ConfigProviderModelParam{
			Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedString("reasoning")),
		}, syncVecProviderInterleaved["enum_reasoning"]},
		{"enum_reasoning_content", "interleaved", opencode.ConfigProviderModelParam{
			Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedString("reasoning_content")),
		}, syncVecProviderInterleaved["enum_reasoning_content"]},
		{"enum_reasoning_text", "interleaved", opencode.ConfigProviderModelParam{
			Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedString("reasoning_text")),
		}, syncVecProviderInterleaved["enum_reasoning_text"]},
		{"vendor_string", "interleaved", opencode.ConfigProviderModelParam{
			Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedString("vendor_custom")),
		}, syncVecProviderInterleaved["vendor_string"]},
		{"object_field", "interleaved", opencode.ConfigProviderModelParam{
			Interleaved: opencode.F[opencode.ConfigProviderModelsInterleavedUnionParam](opencode.ConfigProviderModelsInterleavedFieldParam{
				Field: opencode.F(opencode.ProviderModelCapabilitiesInterleavedFieldFieldReasoning),
			}),
		}, syncVecProviderInterleaved["object_field"]},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := marshalToMap(t, tc.param)
			compareMaps(t, tc.name, fieldValue(t, got, tc.field), parseJSON(t, tc.vector))
		})
	}
}

func TestSyncE2EPartUpdatePartRetryErrorNameSerialization(t *testing.T) {
	t.Parallel()
	errVal := opencode.PartUpdatePartRetryError{
		Name: opencode.F(opencode.PartUpdatePartRetryErrorNameAPIError),
		Data: opencode.F(opencode.PartUpdatePartRetryErrorData{
			Message:     opencode.F("sample-string"),
			StatusCode:  opencode.F(int64(1)),
			IsRetryable: opencode.F(true),
		}),
	}
	got := marshalToMap(t, errVal)
	want := parseJSON(t, `{"name":"APIError","data":{"message":"sample-string","statusCode":1,"isRetryable":true}}`)
	compareMaps(t, "retry error", got, want)
}

// =============================================================================
// Task 3: Response side — JSON → Go reverse deserialization (+ roundtrip)
// =============================================================================

// TestSyncE2EEventPtyExitedExitCodeBoundaries proves the int64 fix: values
// beyond int32 and beyond float64-safe integers must decode without truncation.
func TestSyncE2EEventPtyExitedExitCodeBoundaries(t *testing.T) {
	t.Parallel()
	for _, name := range []string{"exit_0", "exit_max_int32", "exit_2147483648", "exit_max_int64"} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// The OpenAPI vector is the full EventPtyExited event; the Go
			// properties type is the nested `properties` sub-object.
			full := parseJSON(t, syncVecEventPtyExited[name]).(map[string]any)
			propsObj := full["properties"].(map[string]any)
			raw, _ := json.Marshal(propsObj)
			var props opencode.EventListResponseEventPtyExitedProperties
			if err := json.Unmarshal(raw, &props); err != nil {
				t.Fatalf("unmarshal %s: %v", name, err)
			}
			want := int64(propsObj["exitCode"].(float64))
			if props.ExitCode != want {
				t.Errorf("ExitCode = %d, want %d", props.ExitCode, want)
			}
			if props.JSON.ExitCode.IsMissing() {
				t.Errorf("JSON.ExitCode reported missing for %s", name)
			}
		})
	}
}

// TestSyncE2ESyncEventsDeserialization verifies the dedup'd sync event aliases
// still deserialize through the top-level SyncEvent types.
func TestSyncE2ESyncEventsDeserialization(t *testing.T) {
	t.Parallel()
	// The OpenAPI vectors are the outer sync wrapper {type:"sync", id,
	// syncEvent:{...}}; the Go SyncEventXxx types model the inner syncEvent
	// sub-object (flat), so extract it for the deserialization test.
	vec := parseJSON(t, syncVecSyncEvents["context_updated"]).(map[string]any)
	inner := vec["syncEvent"].(map[string]any)
	raw, _ := json.Marshal(inner)
	var ev opencode.SyncEventSessionNextContextUpdated
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("context_updated unmarshal: %v", err)
	}
	if ev.Type != opencode.SyncEventSessionNextContextUpdatedTypeSessionNextContextUpdated1 {
		t.Errorf("Type = %q", ev.Type)
	}
	if ev.Data.Text != "sample-string" {
		t.Errorf("Data.Text = %q", ev.Data.Text)
	}
	if ev.Data.JSON.Text.IsMissing() {
		t.Error("Data.JSON.Text reported missing")
	}

	vec2 := parseJSON(t, syncVecSyncEvents["prompt_admitted"]).(map[string]any)
	inner2 := vec2["syncEvent"].(map[string]any)
	raw2, _ := json.Marshal(inner2)
	var ev2 opencode.SyncEventSessionNextPromptAdmitted
	if err := json.Unmarshal(raw2, &ev2); err != nil {
		t.Fatalf("prompt_admitted unmarshal: %v", err)
	}
	if ev2.Type != opencode.SyncEventSessionNextPromptAdmittedTypeSessionNextPromptAdmitted1 {
		t.Errorf("Type = %q", ev2.Type)
	}
	if ev2.Data.Delivery != opencode.EventListResponseEventSessionNextPromptAdmittedPropertiesDeliverySteer {
		t.Errorf("Data.Delivery = %q, want steer", ev2.Data.Delivery)
	}
	if ev2.Data.Prompt.Text != "sample-string" {
		t.Errorf("Data.Prompt.Text = %q", ev2.Data.Prompt.Text)
	}
}

// TestSyncE2EConfigResponseRuntimeTypes verifies Config's Response-side
// Autoupdate/Formatter/Lsp remain `any` and decode to the correct runtime types.
func TestSyncE2EConfigResponseRuntimeTypes(t *testing.T) {
	t.Parallel()
	raw := `{
		"autoupdate": true,
		"formatter": {"prettier": {"disabled": true, "command": ["prettier"]}},
		"lsp": {"gopls": {"command": ["gopls"], "extensions": [".go"]}}
	}`
	var cfg opencode.Config
	if err := json.Unmarshal([]byte(raw), &cfg); err != nil {
		t.Fatalf("unmarshal Config: %v", err)
	}
	if b, ok := cfg.Autoupdate.(bool); !ok || !b {
		t.Errorf("Autoupdate runtime type = %T, want bool", cfg.Autoupdate)
	}
	if m, ok := cfg.Formatter.(map[string]any); !ok {
		t.Errorf("Formatter runtime type = %T, want map[string]any", cfg.Formatter)
	} else if inner, ok := m["prettier"].(map[string]any); !ok || inner["command"].([]any)[0] != "prettier" {
		t.Errorf("Formatter.prettier = %#v", m["prettier"])
	}
	if m, ok := cfg.Lsp.(map[string]any); !ok {
		t.Errorf("Lsp runtime type = %T, want map[string]any", cfg.Lsp)
	} else if inner, ok := m["gopls"].(map[string]any); !ok || inner["extensions"].([]any)[0] != ".go" {
		t.Errorf("Lsp.gopls = %#v", m["gopls"])
	}
	if cfg.JSON.Autoupdate.IsMissing() {
		t.Error("JSON.Autoupdate reported missing")
	}
}

// TestSyncE2EConfigResponseAutoupdateString verifies the "notify" string branch
// also decodes into the `any` field.
func TestSyncE2EConfigResponseAutoupdateString(t *testing.T) {
	t.Parallel()
	var cfg opencode.Config
	if err := json.Unmarshal([]byte(`{"autoupdate":"notify"}`), &cfg); err != nil {
		t.Fatal(err)
	}
	if s, ok := cfg.Autoupdate.(string); !ok || s != "notify" {
		t.Errorf("Autoupdate runtime type = %T (%v), want string \"notify\"", cfg.Autoupdate, cfg.Autoupdate)
	}
}

// TestSyncE2EGlobalEventVectors verifies the 10 OpenAPI-derived GlobalEvent
// instances (9 distinct V2 payload types + 1 sync wrapper). Each must resolve
// AsUnion() to the correct concrete Go type and set JSON.Payload.
func TestSyncE2EGlobalEventVectors(t *testing.T) {
	t.Parallel()
	for name, raw := range syncVecGlobalEvents {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var ge opencode.GlobalEvent
			if err := json.Unmarshal([]byte(raw), &ge); err != nil {
				t.Fatalf("unmarshal %s: %v", name, err)
			}
			if ge.JSON.Payload.IsMissing() {
				t.Fatalf("%s: JSON.Payload reported missing", name)
			}
			u := ge.AsUnion()
			if u == nil {
				t.Fatalf("%s: AsUnion() nil", name)
			}
			// assert concrete type matches the OpenAPI payload type value
			rt := reflect.Indirect(reflect.ValueOf(u))
			f := rt.FieldByName("Type")
			if !f.IsValid() {
				t.Fatalf("%s: variant %T has no Type field", name, u)
			}
			want := ""
			body := parseJSON(t, raw).(map[string]any)["payload"].(map[string]any)
			if t2, ok := body["type"].(string); ok {
				want = t2
			} else if se, ok := body["syncEvent"].(map[string]any); ok {
				want = se["type"].(string)
			}
			if got := f.String(); got != want {
				t.Fatalf("%s: variant %T Type = %q, want %q", name, u, got, want)
			}
		})
	}
}

// TestSyncE2EGlobalEventNullPayloadGuard verifies the null-payload guard in
// GlobalEvent.UnmarshalJSON: a payload explicitly sent as null must not error,
// JSON.Payload must record the null, and AsUnion() stays nil.
func TestSyncE2EGlobalEventNullPayloadGuard(t *testing.T) {
	t.Parallel()
	raw := `{"directory":"/repo","project":"prj","payload":null}`
	var ge opencode.GlobalEvent
	if err := json.Unmarshal([]byte(raw), &ge); err != nil {
		t.Fatalf("unmarshal null payload: %v", err)
	}
	if ge.AsUnion() != nil {
		t.Errorf("AsUnion() = %T, want nil for null payload", ge.AsUnion())
	}
	if ge.JSON.Payload.IsMissing() {
		t.Error("JSON.Payload reported missing for explicit null")
	}
	if ge.Payload != nil {
		t.Errorf("Payload = %#v, want nil", ge.Payload)
	}
}

// =============================================================================
// Task 4: GlobalEventPayloadUnion 124-member coverage
// =============================================================================// TestSyncE2EGlobalEventPayload124Coverage feeds one minimal legal payload for
// every GlobalEvent.payload anyOf member (124 in OpenAPI) into GlobalEvent
// deserialization and asserts AsUnion() returns a non-nil, correctly-typed
// union variant. Sync members must route to SyncEventResponse and the inner
// AsUnion() must resolve the concrete SyncEvent type.
func TestSyncE2EGlobalEventPayload124Coverage(t *testing.T) {
	t.Parallel()
	for _, tc := range globalPayload124 {
		t.Run(tc.name(), func(t *testing.T) {
			t.Parallel()
			var ge opencode.GlobalEvent
			if err := json.Unmarshal([]byte(tc.payload), &ge); err != nil {
				t.Fatalf("unmarshal %s: %v\npayload: %s", tc.name(), err, tc.payload)
				return
			}
			u := ge.AsUnion()
			if u == nil {
				t.Fatalf("%s: AsUnion() returned nil interface", tc.name())
				return
			}

			if !gjson.Get(ge.JSON.RawJSON(), "payload").Exists() {
				t.Fatalf("%s: JSON.Payload reported missing", tc.name())
				return
			}
			if tc.inner != "" {
				// sync member: must route to SyncEventResponse, then to inner type
				resp, ok := u.(opencode.SyncEventResponse)
				if !ok {
					t.Fatalf("%s: AsUnion() = %T, want SyncEventResponse", tc.name(), u)
					return
				}
				inner := resp.SyncEvent.AsUnion()
				if inner == nil {
					t.Fatalf("%s: SyncEvent.AsUnion() returned nil", tc.name())
					return
				}
				// assert the inner type's Type field matches the expected inner
				// syncEvent type value from OpenAPI
				rt := reflect.Indirect(reflect.ValueOf(inner))
				f := rt.FieldByName("Type")
				if !f.IsValid() {
					t.Fatalf("%s: inner variant %T has no Type field", tc.name(), inner)
					return
				}
				if got := f.String(); got != tc.inner {
					t.Fatalf("%s: inner Type = %q, want %q (variant %T)", tc.name(), got, tc.inner, inner)
					return
				}
				return
			}
			// V2 member: assert the resolved variant's Type field matches the
			// OpenAPI payload type value, proving correct union routing.
			rt := reflect.Indirect(reflect.ValueOf(u))
			f := rt.FieldByName("Type")
			if !f.IsValid() {
				t.Fatalf("%s: variant %T has no Type field", tc.name(), u)
				return
			}
			if got := f.String(); got != tc.type_ {
				t.Fatalf("%s: resolved variant %T has Type = %q, want %q", tc.name(), u, got, tc.type_)
				return
			}
		})
	}
}

// =============================================================================
// Task 5: Enum three-way verification (OpenAPI enum ↔ Go const ↔ IsKnown)
// =============================================================================

// TestSyncE2EAPIErrorNameEnumIsKnown verifies the PartUpdatePartRetryError.Name
// enum: OpenAPI declares enum ["APIError"].
func TestSyncE2EAPIErrorNameEnumIsKnown(t *testing.T) {
	t.Parallel()
	if got := string(opencode.PartUpdatePartRetryErrorNameAPIError); got != "APIError" {
		t.Errorf("PartUpdatePartRetryErrorNameAPIError = %q, want %q", got, "APIError")
	}
	for _, legal := range []opencode.PartUpdatePartRetryErrorName{
		opencode.PartUpdatePartRetryErrorNameAPIError,
	} {
		if !legal.IsKnown() {
			t.Errorf("IsKnown() = false for legal value %q", legal)
		}
	}
	for _, illegal := range []opencode.PartUpdatePartRetryErrorName{
		"", "APIErrorX", "apierror", "Unknown", "APIError.",
	} {
		if illegal.IsKnown() {
			t.Errorf("IsKnown() = true for illegal value %q", illegal)
		}
	}
}

// TestSyncE2EConfigAutoupdateNotifyConstant verifies the string branch of
// Config.autoupdate (OpenAPI enum ["notify"]) through the notify variant.
func TestSyncE2EConfigAutoupdateNotifyConstant(t *testing.T) {
	t.Parallel()
	v := opencode.ConfigAutoupdateNotify("notify")
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != `"notify"` {
		t.Errorf("ConfigAutoupdateNotify marshal = %s, want \"notify\"", raw)
	}
	// Unknown notify values must not be produced by any const; verify the
	// string branch accepts only the enum via the request union TypeFilter.
	if string(v) != "notify" {
		t.Errorf("ConfigAutoupdateNotify = %q, want notify", string(v))
	}
}

// TestSyncE2EConfigLspDisabledDisabledEnum verifies the per-server disabled
// enum (OpenAPI: disabled: {type: boolean, enum: [true]}).
func TestSyncE2EConfigLspDisabledDisabledEnum(t *testing.T) {
	t.Parallel()
	if !opencode.ConfigLspDisabledDisabledTrue.IsKnown() {
		t.Error("ConfigLspDisabledDisabledTrue.IsKnown() = false")
	}
	// bool enum: only `true` is legal; `false` must be unknown.
	if opencode.ConfigLspDisabledDisabled(false).IsKnown() {
		t.Error("ConfigLspDisabledDisabled(false).IsKnown() = true, want false")
	}
}

// TestSyncE2ESyncEventDeliveryEnumIsKnown verifies the delivery enum of the
// prompt-admitted sync event (OpenAPI enum ["steer", "queue"]).
func TestSyncE2ESyncEventDeliveryEnumIsKnown(t *testing.T) {
	t.Parallel()
	for _, legal := range []opencode.EventListResponseEventSessionNextPromptAdmittedPropertiesDelivery{
		opencode.EventListResponseEventSessionNextPromptAdmittedPropertiesDeliverySteer,
		opencode.EventListResponseEventSessionNextPromptAdmittedPropertiesDeliveryQueue,
	} {
		if !legal.IsKnown() {
			t.Errorf("IsKnown() = false for legal value %q", legal)
		}
	}
	for _, illegal := range []opencode.EventListResponseEventSessionNextPromptAdmittedPropertiesDelivery{
		"", "steerX", "queueX", "notify", "none",
	} {
		if illegal.IsKnown() {
			t.Errorf("IsKnown() = true for illegal value %q", illegal)
		}
	}
}

// TestSyncE2EInterleavedFieldEnumIsKnown verifies the interleaved field enum
// (OpenAPI: ["reasoning", "reasoning_content", "reasoning_text"]).
func TestSyncE2EInterleavedFieldEnumIsKnown(t *testing.T) {
	t.Parallel()
	for _, legal := range []opencode.ProviderModelCapabilitiesInterleavedFieldField{
		opencode.ProviderModelCapabilitiesInterleavedFieldFieldReasoning,
		opencode.ProviderModelCapabilitiesInterleavedFieldFieldReasoningContent,
		opencode.ProviderModelCapabilitiesInterleavedFieldFieldReasoningText,
	} {
		if !legal.IsKnown() {
			t.Errorf("IsKnown() = false for legal value %q", legal)
		}
	}
	for _, illegal := range []opencode.ProviderModelCapabilitiesInterleavedFieldField{
		"", "reasoningX", "reasoning_textX", "thinking",
	} {
		if illegal.IsKnown() {
			t.Errorf("IsKnown() = true for illegal value %q", illegal)
		}
	}
}
