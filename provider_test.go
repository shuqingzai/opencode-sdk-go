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

// TestProviderAuthMethodPromptsUnmarshal verifies that ProviderAuthMethod.Prompts
// (typed as []ProviderAuthMethodPrompt) correctly deserialises from JSON. Each
// element of the heterogeneous array resolves to the concrete variant type
// ([ProviderAuthMethodPromptText] or [ProviderAuthMethodPromptSelect]) via the
// registered apijson Union discriminator "type".
//
// Run with: go test -run TestProviderAuthMethodPromptsUnmarshal -v ./...
func TestProviderAuthMethodPromptsUnmarshal(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name           string
		json           string
		wantType       opencode.ProviderAuthMethodType
		wantLabel      string
		wantPromptsNil bool
		wantPromptsLen int
	}{
		{
			name:           "no_prompts",
			json:           `{"type":"oauth","label":"OAuth"}`,
			wantType:       opencode.ProviderAuthMethodTypeOauth,
			wantLabel:      "OAuth",
			wantPromptsNil: true,
		},
		{
			name:           "text_prompt",
			json:           `{"type":"api","label":"API Key","prompts":[{"type":"text","key":"api_key","message":"Enter your API key"}]}`,
			wantType:       opencode.ProviderAuthMethodTypeAPI,
			wantLabel:      "API Key",
			wantPromptsLen: 1,
		},
		{
			name:           "select_prompt",
			json:           `{"type":"api","label":"Select Region","prompts":[{"type":"select","key":"region","message":"Select region","options":[{"label":"US","value":"us"},{"label":"EU","value":"eu"}]}]}`,
			wantType:       opencode.ProviderAuthMethodTypeAPI,
			wantLabel:      "Select Region",
			wantPromptsLen: 1,
		},
		{
			name:           "mixed_prompts",
			json:           `{"type":"api","label":"Mixed","prompts":[{"type":"text","key":"api_key","message":"Enter key"},{"type":"select","key":"region","message":"Region","options":[{"label":"US","value":"us"}]}]}`,
			wantType:       opencode.ProviderAuthMethodTypeAPI,
			wantLabel:      "Mixed",
			wantPromptsLen: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var pam opencode.ProviderAuthMethod
			if err := json.Unmarshal([]byte(tc.json), &pam); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			if pam.Type != tc.wantType {
				t.Errorf("Type: got %q, want %q", pam.Type, tc.wantType)
			}
			if pam.Label != tc.wantLabel {
				t.Errorf("Label: got %q, want %q", pam.Label, tc.wantLabel)
			}
			if tc.wantPromptsNil {
				if pam.Prompts != nil {
					t.Errorf("Prompts: expected nil, got %v", pam.Prompts)
				}
				return
			}
			if len(pam.Prompts) != tc.wantPromptsLen {
				t.Errorf("Prompts length: got %d, want %d", len(pam.Prompts), tc.wantPromptsLen)
			}
			// Each element must resolve to a concrete variant, not a bare map.
			for i, p := range pam.Prompts {
				switch p.AsUnion().(type) {
				case opencode.ProviderAuthMethodPromptText, opencode.ProviderAuthMethodPromptSelect:
					// correct typed variant
				default:
					t.Errorf("Prompts[%d]: expected ProviderAuthMethodPromptText or ProviderAuthMethodPromptSelect, got %T", i, p.AsUnion())
				}
			}
		})
	}
}

// TestProviderAuthMethodPromptCarryingStructUnmarshal verifies that the
// carrying struct [ProviderAuthMethodPrompt] (the element type of
// [ProviderAuthMethod.Prompts]) correctly decodes the concrete field values
// from both the text and select JSON payloads, including the typed Type enum.
func TestProviderAuthMethodPromptCarryingStructUnmarshal(t *testing.T) {
	t.Parallel()

	t.Run("text", func(t *testing.T) {
		t.Parallel()
		const raw = `{"type":"text","key":"api_key","message":"Enter your API key","placeholder":"sk-...","when":{"key":"needs_key","op":"eq","value":"true"}}`
		var p opencode.ProviderAuthMethodPrompt
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if p.Type != opencode.ProviderAuthMethodPromptTypeText {
			t.Errorf("Type: got %q, want %q", p.Type, opencode.ProviderAuthMethodPromptTypeText)
		}
		if !p.Type.IsKnown() {
			t.Error("IsKnown: expected true for text")
		}
		if p.Key != "api_key" {
			t.Errorf("Key: got %q, want api_key", p.Key)
		}
		if p.Message != "Enter your API key" {
			t.Errorf("Message: got %q", p.Message)
		}
		if p.Placeholder != "sk-..." {
			t.Errorf("Placeholder: got %q, want sk-...", p.Placeholder)
		}
		if p.When.Key != "needs_key" || p.When.Op != opencode.ProviderAuthMethodPromptWhenOpEq || p.When.Value != "true" {
			t.Errorf("When: got %+v", p.When)
		}
		if len(p.Options) != 0 {
			t.Errorf("Options: got %d, want 0", len(p.Options))
		}
		// RawJSON must be preserved so callers can access the full payload.
		if p.JSON.RawJSON() == "" {
			t.Error("RawJSON: expected non-empty raw payload")
		}
	})

	t.Run("select", func(t *testing.T) {
		t.Parallel()
		const raw = `{"type":"select","key":"region","message":"Select region","options":[{"label":"US East","value":"us-east-1","hint":"Lower latency"},{"label":"EU West","value":"eu-west-1"}]}`
		var p opencode.ProviderAuthMethodPrompt
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if p.Type != opencode.ProviderAuthMethodPromptTypeSelect {
			t.Errorf("Type: got %q, want %q", p.Type, opencode.ProviderAuthMethodPromptTypeSelect)
		}
		if !p.Type.IsKnown() {
			t.Error("IsKnown: expected true for select")
		}
		if p.Key != "region" {
			t.Errorf("Key: got %q, want region", p.Key)
		}
		if p.Message != "Select region" {
			t.Errorf("Message: got %q", p.Message)
		}
		if len(p.Options) != 2 {
			t.Fatalf("Options length: got %d, want 2", len(p.Options))
		}
		if p.Options[0].Label != "US East" || p.Options[0].Value != "us-east-1" || p.Options[0].Hint != "Lower latency" {
			t.Errorf("Options[0]: %+v", p.Options[0])
		}
		if p.Options[1].Label != "EU West" || p.Options[1].Value != "eu-west-1" {
			t.Errorf("Options[1]: %+v", p.Options[1])
		}
		if p.Placeholder != "" {
			t.Errorf("Placeholder: got %q, want empty", p.Placeholder)
		}
	})

	t.Run("unknown_type", func(t *testing.T) {
		t.Parallel()
		const raw = `{"type":"custom","key":"k","message":"m"}`
		var p opencode.ProviderAuthMethodPrompt
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if p.Type != opencode.ProviderAuthMethodPromptType("custom") {
			t.Errorf("Type: got %q, want custom", p.Type)
		}
		if p.Type.IsKnown() {
			t.Error("IsKnown: expected false for unknown type")
		}
	})
}

// TestProviderAuthMethodPromptTextUnmarshal verifies that ProviderAuthMethodPromptText
// still deserialises correctly on its own (regression guard for Task 1 change).
func TestProviderAuthMethodPromptTextUnmarshal(t *testing.T) {
	t.Parallel()
	const raw = `{"type":"text","key":"api_key","message":"Enter your API key","placeholder":"sk-..."}`
	var p opencode.ProviderAuthMethodPromptText
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if p.Type != opencode.ProviderAuthMethodPromptTextTypeText {
		t.Errorf("Type: got %q, want %q", p.Type, opencode.ProviderAuthMethodPromptTextTypeText)
	}
	if p.Key != "api_key" {
		t.Errorf("Key: got %q, want api_key", p.Key)
	}
	if p.Message != "Enter your API key" {
		t.Errorf("Message: got %q", p.Message)
	}
	if p.Placeholder != "sk-..." {
		t.Errorf("Placeholder: got %q, want sk-...", p.Placeholder)
	}
}

// TestProviderAuthMethodPromptSelectUnmarshal verifies that ProviderAuthMethodPromptSelect
// still deserialises correctly on its own (regression guard for Task 1 change).
func TestProviderAuthMethodPromptSelectUnmarshal(t *testing.T) {
	t.Parallel()
	const raw = `{"type":"select","key":"region","message":"Select region","options":[{"label":"US East","value":"us-east-1","hint":"Lower latency"},{"label":"EU West","value":"eu-west-1"}]}`
	var p opencode.ProviderAuthMethodPromptSelect
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if p.Type != opencode.ProviderAuthMethodPromptSelectTypeSelect {
		t.Errorf("Type: got %q, want %q", p.Type, opencode.ProviderAuthMethodPromptSelectTypeSelect)
	}
	if p.Key != "region" {
		t.Errorf("Key: got %q, want region", p.Key)
	}
	if len(p.Options) != 2 {
		t.Fatalf("Options length: got %d, want 2", len(p.Options))
	}
	if p.Options[0].Label != "US East" || p.Options[0].Value != "us-east-1" || p.Options[0].Hint != "Lower latency" {
		t.Errorf("Options[0]: %+v", p.Options[0])
	}
	if p.Options[1].Label != "EU West" || p.Options[1].Value != "eu-west-1" {
		t.Errorf("Options[1]: %+v", p.Options[1])
	}
}

// TestProviderInfoSourceAPIIsKnown verifies the renamed constant ProviderInfoSourceAPI
// (formerly ProviderInfoSourceApi) and its deprecated alias both resolve to "api".
func TestProviderInfoSourceAPIIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.ProviderInfoSource
		want  bool
	}{
		{name: "env", value: opencode.ProviderInfoSourceEnv, want: true},
		{name: "config", value: opencode.ProviderInfoSourceConfig, want: true},
		{name: "custom", value: opencode.ProviderInfoSourceCustom, want: true},
		// new canonical name
		{name: "api_new", value: opencode.ProviderInfoSourceAPI, want: true},
		// deprecated alias — must equal the new constant
		{name: "api_deprecated", value: opencode.ProviderInfoSourceApi, want: true},
		// unknown
		{name: "unknown", value: "unknown", want: false},
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
	// Alias equality guard
	if opencode.ProviderInfoSourceApi != opencode.ProviderInfoSourceAPI {
		t.Errorf("ProviderInfoSourceApi alias != ProviderInfoSourceAPI: %q vs %q",
			opencode.ProviderInfoSourceApi, opencode.ProviderInfoSourceAPI)
	}
}

// TestProviderAuthMethodTypeAPIIsKnown verifies the renamed constant ProviderAuthMethodTypeAPI
// (formerly ProviderAuthMethodTypeApi) and its deprecated alias both resolve to "api".
func TestProviderAuthMethodTypeAPIIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.ProviderAuthMethodType
		want  bool
	}{
		{name: "oauth", value: opencode.ProviderAuthMethodTypeOauth, want: true},
		// new canonical name
		{name: "api_new", value: opencode.ProviderAuthMethodTypeAPI, want: true},
		// deprecated alias
		{name: "api_deprecated", value: opencode.ProviderAuthMethodTypeApi, want: true},
		// unknown
		{name: "unknown", value: "unknown", want: false},
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
	// Alias equality guard
	if opencode.ProviderAuthMethodTypeApi != opencode.ProviderAuthMethodTypeAPI {
		t.Errorf("ProviderAuthMethodTypeApi alias != ProviderAuthMethodTypeAPI: %q vs %q",
			opencode.ProviderAuthMethodTypeApi, opencode.ProviderAuthMethodTypeAPI)
	}
}

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
