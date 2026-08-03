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

func TestProviderList(t *testing.T) {
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
	_, err := client.Provider.List(context.TODO(), opencode.ProviderListParams{
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

func TestProviderAuth(t *testing.T) {
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
	_, err := client.Provider.Auth(context.TODO(), opencode.ProviderAuthParams{
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

func TestProviderOauthAuthorize(t *testing.T) {
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
	_, err := client.Provider.OauthAuthorize(context.TODO(), "providerID", opencode.ProviderOauthAuthorizeParams{
		Method: opencode.F(int64(0)),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProviderOauthCallback(t *testing.T) {
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
	_, err := client.Provider.OauthCallback(context.TODO(), "providerID", opencode.ProviderOauthCallbackParams{
		Method: opencode.F(int64(0)),
		Code:   opencode.F("code"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestProviderModelCapabilitiesInterleavedUnmarshal verifies that
// ProviderModelCapabilities.Interleaved correctly deserialises all OpenAPI
// v1.18.11 variants of the `interleaved` schema:
//
//	anyOf[boolean, { field: "reasoning"|"reasoning_content"|"reasoning_text"|string }]
//
// Relative to v1.18.4 this schema changed in exactly one way: the object variant's
// `field` became an open union anyOf[enum("reasoning"|"reasoning_content"|
// "reasoning_text"), string] — the enum value "reasoning_details" was replaced by
// "reasoning_text" and arbitrary strings are now admitted. The boolean variant was
// already a plain `{"type":"boolean"}` in v1.18.4 (unlike
// ProviderConfig.models.*.interleaved, which was narrowed to `enum:[true]` back then).
//
// A second, non-schema property this test pins down: at runtime the object variant
// deserialises to map[string]any, NOT to
// ProviderModelCapabilitiesInterleavedField, because apijson hands bare any
// fields the raw gjson value (internal/apijson/decoder.go reflect.Interface branch).
//
// Run with: go test -run TestProviderModelCapabilitiesInterleavedUnmarshal -v ./...
func TestProviderModelCapabilitiesInterleavedUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name         string
		json         string
		wantRuntimeT reflect.Type
		wantBool     bool   // checked only when wantRuntimeT == bool
		wantFieldVal string // checked only when wantRuntimeT == map[string]any
	}{
		// boolean variant — `{"type":"boolean"}`, unchanged since v1.18.4
		{
			name:         "bool_true",
			json:         `{"temperature":true,"reasoning":false,"attachment":false,"toolcall":false,"input":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"output":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"interleaved":true}`,
			wantRuntimeT: reflect.TypeFor[bool](),
			wantBool:     true,
		},
		// boolean variant, false — also admitted by `{"type":"boolean"}`
		{
			name:         "bool_false",
			json:         `{"temperature":true,"reasoning":false,"attachment":false,"toolcall":false,"input":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"output":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"interleaved":false}`,
			wantRuntimeT: reflect.TypeFor[bool](),
			wantBool:     false,
		},
		// object variant — known field "reasoning"
		{
			name:         "object_reasoning",
			json:         `{"temperature":true,"reasoning":false,"attachment":false,"toolcall":false,"input":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"output":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"interleaved":{"field":"reasoning"}}`,
			wantRuntimeT: reflect.TypeFor[map[string]any](),
			wantFieldVal: "reasoning",
		},
		// object variant — known field "reasoning_content"
		{
			name:         "object_reasoning_content",
			json:         `{"temperature":true,"reasoning":false,"attachment":false,"toolcall":false,"input":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"output":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"interleaved":{"field":"reasoning_content"}}`,
			wantRuntimeT: reflect.TypeFor[map[string]any](),
			wantFieldVal: "reasoning_content",
		},
		// object variant — known field "reasoning_text"
		{
			name:         "object_reasoning_text",
			json:         `{"temperature":true,"reasoning":false,"attachment":false,"toolcall":false,"input":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"output":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"interleaved":{"field":"reasoning_text"}}`,
			wantRuntimeT: reflect.TypeFor[map[string]any](),
			wantFieldVal: "reasoning_text",
		},
		// object variant — open-union: vendor-specific field name (OpenAPI allows any string)
		{
			name:         "object_custom_vendor_field",
			json:         `{"temperature":true,"reasoning":false,"attachment":false,"toolcall":false,"input":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"output":{"text":false,"audio":false,"image":false,"video":false,"pdf":false},"interleaved":{"field":"custom_vendor_field"}}`,
			wantRuntimeT: reflect.TypeFor[map[string]any](),
			wantFieldVal: "custom_vendor_field",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var caps opencode.ProviderModelCapabilities
			if err := json.Unmarshal([]byte(tc.json), &caps); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			gotT := reflect.TypeOf(caps.Interleaved)
			if gotT != tc.wantRuntimeT {
				t.Errorf("Interleaved runtime type: got %v, want %v", gotT, tc.wantRuntimeT)
			}
			switch tc.wantRuntimeT {
			case reflect.TypeFor[bool]():
				gotBool, ok := caps.Interleaved.(bool)
				if !ok {
					t.Fatalf("expected bool, got %T", caps.Interleaved)
				}
				if gotBool != tc.wantBool {
					t.Errorf("Interleaved bool value: got %v, want %v", gotBool, tc.wantBool)
				}
			case reflect.TypeFor[map[string]any]():
				m, ok := caps.Interleaved.(map[string]any)
				if !ok {
					t.Fatalf("expected map[string]any, got %T", caps.Interleaved)
				}
				gotField, ok := m["field"].(string)
				if !ok {
					t.Fatalf(`interleaved["field"] is not a string: %T`, m["field"])
				}
				if gotField != tc.wantFieldVal {
					t.Errorf(`interleaved["field"]: got %q, want %q`, gotField, tc.wantFieldVal)
				}
			}
		})
	}
}

// TestProviderModelCapabilitiesInterleavedFieldUnmarshal verifies that
// ProviderModelCapabilitiesInterleavedField can be decoded explicitly via
// json.Unmarshal when callers opt to decode the object variant directly.
func TestProviderModelCapabilitiesInterleavedFieldUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name      string
		json      string
		wantField string
	}{
		{name: "reasoning", json: `{"field":"reasoning"}`, wantField: "reasoning"},
		{name: "reasoning_content", json: `{"field":"reasoning_content"}`, wantField: "reasoning_content"},
		{name: "reasoning_text", json: `{"field":"reasoning_text"}`, wantField: "reasoning_text"},
		{name: "custom", json: `{"field":"custom_vendor_field"}`, wantField: "custom_vendor_field"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var f opencode.ProviderModelCapabilitiesInterleavedField
			if err := json.Unmarshal([]byte(tc.json), &f); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			if string(f.Field) != tc.wantField {
				t.Errorf("Field: got %q, want %q", f.Field, tc.wantField)
			}
		})
	}
}

// TestProviderModelCapabilitiesInterleavedFieldFieldIsKnown verifies that
// ProviderModelCapabilitiesInterleavedFieldField.IsKnown returns true only for
// the three canonical values defined by the OpenAPI schema, and false for any
// other string — including the removed "reasoning_details" value (regression
// guard against re-introducing that constant).
func TestProviderModelCapabilitiesInterleavedFieldFieldIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.ProviderModelCapabilitiesInterleavedFieldField
		want  bool
	}{
		// known values (OpenAPI enum)
		{name: "reasoning", value: opencode.ProviderModelCapabilitiesInterleavedFieldFieldReasoning, want: true},
		{name: "reasoning_content", value: opencode.ProviderModelCapabilitiesInterleavedFieldFieldReasoningContent, want: true},
		{name: "reasoning_text", value: opencode.ProviderModelCapabilitiesInterleavedFieldFieldReasoningText, want: true},
		// unknown values — open-union allows them, but IsKnown must return false
		{name: "reasoning_details_removed", value: "reasoning_details", want: false},
		{name: "custom", value: "custom", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}
