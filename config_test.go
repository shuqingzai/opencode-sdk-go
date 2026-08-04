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

// TestConfigPermissionEditUnmarshal verifies that ConfigPermission.Edit correctly
// deserialises both OpenAPI PermissionRuleConfig variants:
//
//  1. string form ("ask"|"allow"|"deny") → runtime type string
//  2. map form ({"src/**": "allow", "**": "ask"}) → runtime type map[string]any
//
// This guards against the regression where Edit was typed as ConfigPermissionEdit
// (a concrete string type) which caused silent data loss when a map was received.
//
// Run with: go test -run TestConfigPermissionEditUnmarshal -v ./...
func TestConfigPermissionEditUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		json        string
		wantRuntime reflect.Type
		wantValue   any
	}{
		// String form: "ask"
		{
			name:        "string_ask",
			json:        `{"edit":"ask"}`,
			wantRuntime: reflect.TypeFor[string](),
			wantValue:   "ask",
		},
		// String form: "allow"
		{
			name:        "string_allow",
			json:        `{"edit":"allow"}`,
			wantRuntime: reflect.TypeFor[string](),
			wantValue:   "allow",
		},
		// String form: "deny"
		{
			name:        "string_deny",
			json:        `{"edit":"deny"}`,
			wantRuntime: reflect.TypeFor[string](),
			wantValue:   "deny",
		},
		// Map form: {"src/**": "allow", "**": "ask"}
		{
			name:        "map_form",
			json:        `{"edit":{"src/**":"allow","**":"ask"}}`,
			wantRuntime: reflect.TypeFor[map[string]any](),
			wantValue:   nil, // checked separately below
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var p opencode.ConfigPermission
			if err := json.Unmarshal([]byte(tc.json), &p); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			gotRuntime := reflect.TypeOf(p.Edit)
			if gotRuntime != tc.wantRuntime {
				t.Errorf("runtime type: got %v, want %v (value: %#v)", gotRuntime, tc.wantRuntime, p.Edit)
			}
			if tc.wantValue != nil {
				if p.Edit != tc.wantValue {
					t.Errorf("value: got %#v, want %#v", p.Edit, tc.wantValue)
				}
			} else {
				// map form: verify key/value
				gotMap, ok := p.Edit.(map[string]any)
				if !ok {
					t.Fatalf("expected map[string]any, got %T", p.Edit)
				}
				if gotMap["src/**"] != "allow" {
					t.Errorf("map[\"src/**\"]: got %q, want \"allow\"", gotMap["src/**"])
				}
				if gotMap["**"] != "ask" {
					t.Errorf("map[\"**\"]: got %q, want \"ask\"", gotMap["**"])
				}
			}
		})
	}
}

// TestConfigAgentBuildPermissionUnmarshal verifies that ConfigAgentBuild.Permission
// correctly deserialises both OpenAPI PermissionConfig anyOf variants:
//
//  1. string form ("ask"|"allow"|"deny") → runtime type ConfigPermissionAction
//  2. object form ({...per-tool rules...}) → runtime type ConfigAgentBuildPermission
//
// This guards against the regression where Permission was typed as
// ConfigAgentBuildPermission (a concrete struct) which caused silent data loss
// when the server returned the short string form.
//
// Run with: go test -run TestConfigAgentBuildPermissionUnmarshal -v ./...
func TestConfigAgentBuildPermissionUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		json        string
		wantRuntime reflect.Type
	}{
		// String form
		{
			name:        "string_ask",
			json:        `{"permission":"ask"}`,
			wantRuntime: reflect.TypeFor[opencode.ConfigPermissionAction](),
		},
		{
			name:        "string_allow",
			json:        `{"permission":"allow"}`,
			wantRuntime: reflect.TypeFor[opencode.ConfigPermissionAction](),
		},
		{
			name:        "string_deny",
			json:        `{"permission":"deny"}`,
			wantRuntime: reflect.TypeFor[opencode.ConfigPermissionAction](),
		},
		// Object form: per-tool rules
		{
			name:        "object_form",
			json:        `{"permission":{"edit":"ask","bash":"allow"}}`,
			wantRuntime: reflect.TypeFor[opencode.ConfigAgentBuildPermission](),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var a opencode.ConfigAgentBuild
			if err := json.Unmarshal([]byte(tc.json), &a); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			gotRuntime := reflect.TypeOf(a.Permission)
			if gotRuntime != tc.wantRuntime {
				t.Errorf("runtime type: got %v, want %v (value: %#v)", gotRuntime, tc.wantRuntime, a.Permission)
			}
			// For string form: verify value
			if tc.wantRuntime == reflect.TypeFor[opencode.ConfigPermissionAction]() {
				action, ok := a.Permission.(opencode.ConfigPermissionAction)
				if !ok {
					t.Fatalf("expected ConfigPermissionAction, got %T", a.Permission)
				}
				if !action.IsKnown() {
					t.Errorf("ConfigPermissionAction.IsKnown() returned false for %q", action)
				}
			}
			// For object form: verify it's the right struct
			if tc.wantRuntime == reflect.TypeFor[opencode.ConfigAgentBuildPermission]() {
				_, ok := a.Permission.(opencode.ConfigAgentBuildPermission)
				if !ok {
					t.Fatalf("expected ConfigAgentBuildPermission, got %T", a.Permission)
				}
			}
		})
	}
}

// TestConfigProviderModelParamInterleavedMarshal verifies that
// ConfigProviderModelParam.Interleaved correctly serialises all four OpenAPI
// anyOf variants to the expected wire-format JSON:
//
//  1. bool       → {"interleaved":true}
//  2. string     → {"interleaved":"reasoning_text"}
//  3. ConfigProviderModelsInterleavedFieldParam → {"interleaved":{"field":"reasoning_text"}}
//
// This validates that the new ConfigProviderModelsInterleavedFieldParam type
// (F3) produces the correct wire format and that the Request/Response
// separation is sound.
//
// Run with: go test -run TestConfigAgentAllPermissionUnionUnmarshal -v ./...
func TestConfigAgentAllPermissionUnionUnmarshal(t *testing.T) {
	t.Parallel()
	// Verifies that all 9 ConfigAgent*/ConfigMode* types correctly deserialise
	// the PermissionConfig anyOf union (string or object form).
	// OpenAPI: components.schemas.PermissionConfig anyOf[string, object].
	cases := []struct {
		name        string
		stringJSON  string
		objectJSON  string
		into        func([]byte) (any, error)
		wantObjType reflect.Type
	}{
		{
			name:       "ConfigAgentGeneral",
			stringJSON: `{"permission":"ask"}`,
			objectJSON: `{"permission":{"edit":"allow","bash":"ask"}}`,
			into: func(data []byte) (any, error) {
				var v opencode.ConfigAgentGeneral
				return &v, json.Unmarshal(data, &v)
			},
			wantObjType: reflect.TypeFor[opencode.ConfigAgentGeneralPermission](),
		},
		{
			name:       "ConfigAgentPlan",
			stringJSON: `{"permission":"allow"}`,
			objectJSON: `{"permission":{"edit":"deny"}}`,
			into: func(data []byte) (any, error) {
				var v opencode.ConfigAgentPlan
				return &v, json.Unmarshal(data, &v)
			},
			wantObjType: reflect.TypeFor[opencode.ConfigAgentPlanPermission](),
		},
		{
			name:       "ConfigAgentExplore",
			stringJSON: `{"permission":"deny"}`,
			objectJSON: `{"permission":{"edit":"ask"}}`,
			into: func(data []byte) (any, error) {
				var v opencode.ConfigAgentExplore
				return &v, json.Unmarshal(data, &v)
			},
			wantObjType: reflect.TypeFor[opencode.ConfigAgentExplorePermission](),
		},
		{
			name:       "ConfigAgentTitle",
			stringJSON: `{"permission":"ask"}`,
			objectJSON: `{"permission":{"edit":"allow"}}`,
			into: func(data []byte) (any, error) {
				var v opencode.ConfigAgentTitle
				return &v, json.Unmarshal(data, &v)
			},
			wantObjType: reflect.TypeFor[opencode.ConfigAgentTitlePermission](),
		},
		{
			name:       "ConfigAgentSummary",
			stringJSON: `{"permission":"allow"}`,
			objectJSON: `{"permission":{"edit":"deny"}}`,
			into: func(data []byte) (any, error) {
				var v opencode.ConfigAgentSummary
				return &v, json.Unmarshal(data, &v)
			},
			wantObjType: reflect.TypeFor[opencode.ConfigAgentSummaryPermission](),
		},
		{
			name:       "ConfigAgentCompaction",
			stringJSON: `{"permission":"deny"}`,
			objectJSON: `{"permission":{"edit":"ask"}}`,
			into: func(data []byte) (any, error) {
				var v opencode.ConfigAgentCompaction
				return &v, json.Unmarshal(data, &v)
			},
			wantObjType: reflect.TypeFor[opencode.ConfigAgentCompactionPermission](),
		},
		{
			name:       "ConfigModeBuild",
			stringJSON: `{"permission":"ask"}`,
			objectJSON: `{"permission":{"edit":"allow"}}`,
			into: func(data []byte) (any, error) {
				var v opencode.ConfigModeBuild
				return &v, json.Unmarshal(data, &v)
			},
			wantObjType: reflect.TypeFor[opencode.ConfigModeBuildPermission](),
		},
		{
			name:       "ConfigModePlan",
			stringJSON: `{"permission":"allow"}`,
			objectJSON: `{"permission":{"edit":"deny"}}`,
			into: func(data []byte) (any, error) {
				var v opencode.ConfigModePlan
				return &v, json.Unmarshal(data, &v)
			},
			wantObjType: reflect.TypeFor[opencode.ConfigModePlanPermission](),
		},
	}

	permField := func(v any) any {
		// All ConfigAgent*/ConfigMode* types embed Permission as an any field.
		// Use reflect to access it generically.
		rv := reflect.ValueOf(v).Elem()
		f := rv.FieldByName("Permission")
		if !f.IsValid() {
			return nil
		}
		return f.Interface()
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// String form → ConfigPermissionAction
			v, err := tc.into([]byte(tc.stringJSON))
			if err != nil {
				t.Fatalf("string form unmarshal: %v", err)
			}
			perm := permField(v)
			if _, ok := perm.(opencode.ConfigPermissionAction); !ok {
				t.Errorf("string form: Permission runtime type = %T, want ConfigPermissionAction", perm)
			}
			action, _ := perm.(opencode.ConfigPermissionAction)
			if !action.IsKnown() {
				t.Errorf("string form: ConfigPermissionAction.IsKnown() = false for %q", action)
			}

			// Object form → specific Permission struct
			v2, err := tc.into([]byte(tc.objectJSON))
			if err != nil {
				t.Fatalf("object form unmarshal: %v", err)
			}
			perm2 := permField(v2)
			gotType := reflect.TypeOf(perm2)
			if gotType != tc.wantObjType {
				t.Errorf("object form: Permission runtime type = %v, want %v", gotType, tc.wantObjType)
			}
		})
	}
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
				Interleaved: opencode.F[any](true),
			},
			wantJSON: `{"interleaved":true}`,
		},
		// Variant 2: enum string
		{
			name: "string_reasoning_text",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[any]("reasoning_text"),
			},
			wantJSON: `{"interleaved":"reasoning_text"}`,
		},
		// Variant 3: open string
		{
			name: "string_vendor",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[any]("vendor_custom"),
			},
			wantJSON: `{"interleaved":"vendor_custom"}`,
		},
		// Variant 4: object { "field": "reasoning_text" }
		// Uses the new ConfigProviderModelsInterleavedFieldParam request-side type.
		{
			name: "object_field_reasoning_text",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[any](opencode.ConfigProviderModelsInterleavedFieldParam{
					Field: opencode.F(opencode.ProviderModelCapabilitiesInterleavedFieldField("reasoning_text")),
				}),
			},
			wantJSON: `{"interleaved":{"field":"reasoning_text"}}`,
		},
		// Variant 4b: object with vendor-custom field value
		{
			name: "object_field_vendor_custom",
			param: opencode.ConfigProviderModelParam{
				Interleaved: opencode.F[any](opencode.ConfigProviderModelsInterleavedFieldParam{
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
