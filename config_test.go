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
