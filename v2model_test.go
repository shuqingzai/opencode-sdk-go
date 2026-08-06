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

func TestV2ModelList(t *testing.T) {
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
	_, err := client.V2Model.List(context.TODO(), opencode.V2ModelListParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestV2ModelInfoApiUnionDiscriminator verifies that V2ModelInfoApiUnion
// correctly dispatches to V2ModelInfoApiAisdk or V2ModelInfoApiNative
// using the "type" discriminator field ("aisdk" / "native").
//
// Before the fix, init() used an empty discriminator "" which relied on exactness
// matching. With the fix, discriminator="type" + DiscriminatorValue ensures
// deterministic dispatch even when both variants share common fields.
//
// Run with: go test -run TestV2ModelInfoApiUnionDiscriminator -v ./...
func TestV2ModelInfoApiUnionDiscriminator(t *testing.T) {
	t.Parallel()

	baseModel := `"id":"gpt-4","providerID":"openai","name":"GPT-4","capabilities":{"tools":true,"input":["text"],"output":["text"]},"request":{"headers":{},"body":{}},"variants":[],"time":{"released":1680000000},"cost":[{"input":0.01,"output":0.03,"cache":{"read":0.001,"write":0.002}}],"status":"active","enabled":true,"limit":{"context":128000,"output":4096}`

	cases := []struct {
		name       string
		apiJSON    string
		wantType   reflect.Type
		wantAisdk  bool
		wantNative bool
		wantPkg    string // only for aisdk
	}{
		{
			name:       "aisdk_variant",
			apiJSON:    `{"type":"aisdk","id":"gpt-4-aisdk","package":"@ai-sdk/openai"}`,
			wantType:   reflect.TypeOf(opencode.V2ModelInfoApiAisdk{}),
			wantAisdk:  true,
			wantNative: false,
			wantPkg:    "@ai-sdk/openai",
		},
		{
			name:       "native_variant",
			apiJSON:    `{"type":"native","id":"gpt-4-native","settings":{}}`,
			wantType:   reflect.TypeOf(opencode.V2ModelInfoApiNative{}),
			wantAisdk:  false,
			wantNative: true,
		},
		{
			name:       "aisdk_with_url",
			apiJSON:    `{"type":"aisdk","id":"custom","package":"@ai-sdk/custom","url":"https://api.custom.com"}`,
			wantType:   reflect.TypeOf(opencode.V2ModelInfoApiAisdk{}),
			wantAisdk:  true,
			wantNative: false,
			wantPkg:    "@ai-sdk/custom",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			raw := `{` + baseModel + `,"api":` + tc.apiJSON + `}`
			var info opencode.V2ModelInfo
			if err := json.Unmarshal([]byte(raw), &info); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			union := info.AsAPIUnion()
			if union == nil {
				t.Fatal("AsAPIUnion() returned nil")
			}

			gotType := reflect.TypeOf(union)
			if gotType != tc.wantType {
				t.Errorf("reflect.TypeOf(union): got %v, want %v", gotType, tc.wantType)
			}

			_, isAisdk := union.(opencode.V2ModelInfoApiAisdk)
			_, isNative := union.(opencode.V2ModelInfoApiNative)

			if isAisdk != tc.wantAisdk {
				t.Errorf("V2ModelInfoApiAisdk assertion: got %v, want %v", isAisdk, tc.wantAisdk)
			}
			if isNative != tc.wantNative {
				t.Errorf("V2ModelInfoApiNative assertion: got %v, want %v", isNative, tc.wantNative)
			}

			if tc.wantAisdk && tc.wantPkg != "" {
				aisdk, _ := union.(opencode.V2ModelInfoApiAisdk)
				if aisdk.Package != tc.wantPkg {
					t.Errorf("Package: got %q, want %q", aisdk.Package, tc.wantPkg)
				}
			}
		})
	}
}

// TestV2ModelInfoApiUnionNoCrossContamination verifies that aisdk JSON is never
// decoded as V2ModelInfoApiNative, and vice versa (discriminator correctness).
func TestV2ModelInfoApiUnionNoCrossContamination(t *testing.T) {
	t.Parallel()

	baseModel := `"id":"m1","providerID":"p1","name":"M1","capabilities":{"tools":true,"input":["text"],"output":["text"]},"request":{"headers":{},"body":{}},"variants":[],"time":{"released":1680000000},"cost":[{"input":0.01,"output":0.03,"cache":{"read":0.001,"write":0.002}}],"status":"active","enabled":true,"limit":{"context":128000,"output":4096}`

	aisdkRaw := `{` + baseModel + `,"api":{"type":"aisdk","id":"x","package":"@ai-sdk/openai"}}`
	nativeRaw := `{` + baseModel + `,"api":{"type":"native","id":"x","settings":{}}}`

	var aisdk opencode.V2ModelInfo
	if err := json.Unmarshal([]byte(aisdkRaw), &aisdk); err != nil {
		t.Fatalf("aisdk unmarshal: %v", err)
	}
	aisdkUnion := aisdk.AsAPIUnion()
	if _, bad := aisdkUnion.(opencode.V2ModelInfoApiNative); bad {
		t.Error("CROSS-CONTAMINATION: aisdk JSON decoded as V2ModelInfoApiNative")
	}
	if _, ok := aisdkUnion.(opencode.V2ModelInfoApiAisdk); !ok {
		t.Errorf("aisdk JSON: AsAPIUnion() type=%T, want V2ModelInfoApiAisdk", aisdkUnion)
	}

	var native opencode.V2ModelInfo
	if err := json.Unmarshal([]byte(nativeRaw), &native); err != nil {
		t.Fatalf("native unmarshal: %v", err)
	}
	nativeUnion := native.AsAPIUnion()
	if _, bad := nativeUnion.(opencode.V2ModelInfoApiAisdk); bad {
		t.Error("CROSS-CONTAMINATION: native JSON decoded as V2ModelInfoApiAisdk")
	}
	if _, ok := nativeUnion.(opencode.V2ModelInfoApiNative); !ok {
		t.Errorf("native JSON: AsAPIUnion() type=%T, want V2ModelInfoApiNative", nativeUnion)
	}
}

// TestV2ModelInfoApiAisdkType verifies V2ModelInfoApiAisdkType.IsKnown.
func TestV2ModelInfoApiAisdkType(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2ModelInfoApiAisdkType
		want  bool
	}{
		{name: "aisdk", value: opencode.V2ModelInfoApiAisdkTypeAisdk, want: true},
		{name: "unknown", value: "unknown", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}

// TestV2ModelInfoApiNativeType verifies V2ModelInfoApiNativeType.IsKnown.
func TestV2ModelInfoApiNativeType(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2ModelInfoApiNativeType
		want  bool
	}{
		{name: "native", value: opencode.V2ModelInfoApiNativeTypeNative, want: true},
		{name: "unknown", value: "unknown", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}

// TestV2ModelInfoApiBearerAisdk verifies the bearer struct V2ModelInfoApi decodes
// an "aisdk" api payload into its fields and exposes the V2ModelInfoApiAisdk
// union variant.
func TestV2ModelInfoApiBearerAisdk(t *testing.T) {
	t.Parallel()

	baseModel := `"id":"gpt-4","providerID":"openai","name":"GPT-4","capabilities":{"tools":true,"input":["text"],"output":["text"]},"request":{"headers":{},"body":{}},"variants":[],"time":{"released":1680000000},"cost":[{"input":0.01,"output":0.03,"cache":{"read":0.001,"write":0.002}}],"status":"active","enabled":true,"limit":{"context":128000,"output":4096}`
	raw := `{` + baseModel + `,"api":{"type":"aisdk","id":"gpt-4-aisdk","package":"@ai-sdk/openai","url":"https://ai-sdk.dev"}}`

	var info opencode.V2ModelInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	if info.Api.Type != opencode.V2ModelInfoApiTypeAisdk {
		t.Errorf("Api.Type: got %q, want %q", info.Api.Type, opencode.V2ModelInfoApiTypeAisdk)
	}
	if info.Api.ID != "gpt-4-aisdk" {
		t.Errorf("Api.ID: got %q, want %q", info.Api.ID, "gpt-4-aisdk")
	}
	if info.Api.Package != "@ai-sdk/openai" {
		t.Errorf("Api.Package: got %q, want %q", info.Api.Package, "@ai-sdk/openai")
	}

	aisdk, ok := info.Api.AsUnion().(opencode.V2ModelInfoApiAisdk)
	if !ok {
		t.Fatalf("Api.AsUnion() type=%T, want V2ModelInfoApiAisdk", info.Api.AsUnion())
	}
	if aisdk.ID != "gpt-4-aisdk" || aisdk.Package != "@ai-sdk/openai" {
		t.Errorf("aisdk variant: got ID=%q Package=%q, want ID=%q Package=%q", aisdk.ID, aisdk.Package, "gpt-4-aisdk", "@ai-sdk/openai")
	}
}

// TestV2ModelInfoApiBearerNative verifies the bearer struct V2ModelInfoApi decodes
// a "native" api payload and resolves its Settings.
func TestV2ModelInfoApiBearerNative(t *testing.T) {
	t.Parallel()

	baseModel := `"id":"gpt-4","providerID":"openai","name":"GPT-4","capabilities":{"tools":true,"input":["text"],"output":["text"]},"request":{"headers":{},"body":{}},"variants":[],"time":{"released":1680000000},"cost":[{"input":0.01,"output":0.03,"cache":{"read":0.001,"write":0.002}}],"status":"active","enabled":true,"limit":{"context":128000,"output":4096}`
	raw := `{` + baseModel + `,"api":{"type":"native","id":"gpt-4-native","url":"https://api.openai.com","settings":{"apiKey":"sk-123","baseURL":"https://api.openai.com/v1"}}}`

	var info opencode.V2ModelInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	if info.Api.Type != opencode.V2ModelInfoApiTypeNative {
		t.Errorf("Api.Type: got %q, want %q", info.Api.Type, opencode.V2ModelInfoApiTypeNative)
	}
	if info.Api.ID != "gpt-4-native" {
		t.Errorf("Api.ID: got %q, want %q", info.Api.ID, "gpt-4-native")
	}
	if got := info.Api.Settings["apiKey"]; got != "sk-123" {
		t.Errorf("Api.Settings[apiKey]: got %v, want %q", got, "sk-123")
	}

	native, ok := info.Api.AsUnion().(opencode.V2ModelInfoApiNative)
	if !ok {
		t.Fatalf("Api.AsUnion() type=%T, want V2ModelInfoApiNative", info.Api.AsUnion())
	}
	if got := native.Settings["baseURL"]; got != "https://api.openai.com/v1" {
		t.Errorf("native.Settings[baseURL]: got %v, want %q", got, "https://api.openai.com/v1")
	}
}

// TestV2ModelInfoApiRawJSON verifies that Api.JSON.RawJSON() returns the complete
// raw api JSON payload.
func TestV2ModelInfoApiRawJSON(t *testing.T) {
	t.Parallel()

	apiJSON := `{"type":"aisdk","id":"raw-test","package":"@ai-sdk/openai","url":"https://raw.dev"}`
	baseModel := `"id":"gpt-4","providerID":"openai","name":"GPT-4","capabilities":{"tools":true,"input":["text"],"output":["text"]},"request":{"headers":{},"body":{}},"variants":[],"time":{"released":1680000000},"cost":[{"input":0.01,"output":0.03,"cache":{"read":0.001,"write":0.002}}],"status":"active","enabled":true,"limit":{"context":128000,"output":4096}`
	raw := `{` + baseModel + `,"api":` + apiJSON + `}`

	var info opencode.V2ModelInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	got := info.Api.JSON.RawJSON()
	if got != apiJSON {
		t.Errorf("Api.JSON.RawJSON() = %q, want %q", got, apiJSON)
	}
	if !json.Valid([]byte(got)) {
		t.Errorf("Api.JSON.RawJSON() = %q, not valid JSON", got)
	}
}

// TestV2ModelInfoAsAPIUnionBackwardCompat verifies that the deprecated
// AsAPIUnion() still returns the correct union after the bearer-struct refactor.
func TestV2ModelInfoAsAPIUnionBackwardCompat(t *testing.T) {
	t.Parallel()

	baseModel := `"id":"gpt-4","providerID":"openai","name":"GPT-4","capabilities":{"tools":true,"input":["text"],"output":["text"]},"request":{"headers":{},"body":{}},"variants":[],"time":{"released":1680000000},"cost":[{"input":0.01,"output":0.03,"cache":{"read":0.001,"write":0.002}}],"status":"active","enabled":true,"limit":{"context":128000,"output":4096}`

	cases := []struct {
		name     string
		apiJSON  string
		wantType reflect.Type
	}{
		{
			name:     "aisdk",
			apiJSON:  `{"type":"aisdk","id":"gpt-4-aisdk","package":"@ai-sdk/openai"}`,
			wantType: reflect.TypeOf(opencode.V2ModelInfoApiAisdk{}),
		},
		{
			name:     "native",
			apiJSON:  `{"type":"native","id":"gpt-4-native","settings":{}}`,
			wantType: reflect.TypeOf(opencode.V2ModelInfoApiNative{}),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			raw := `{` + baseModel + `,"api":` + tc.apiJSON + `}`
			var info opencode.V2ModelInfo
			if err := json.Unmarshal([]byte(raw), &info); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			union := info.AsAPIUnion()
			if union == nil {
				t.Fatal("AsAPIUnion() returned nil")
			}
			if got := reflect.TypeOf(union); got != tc.wantType {
				t.Errorf("AsAPIUnion() type=%v, want %v", got, tc.wantType)
			}
			if !reflect.DeepEqual(info.AsAPIUnion(), info.Api.AsUnion()) {
				t.Error("AsAPIUnion() and Api.AsUnion() returned different unions")
			}
		})
	}
}

// TestV2ModelInfoStatus verifies V2ModelInfoStatus.IsKnown.
func TestV2ModelInfoStatus(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2ModelInfoStatus
		want  bool
	}{
		{name: "alpha", value: opencode.V2ModelInfoStatusAlpha, want: true},
		{name: "beta", value: opencode.V2ModelInfoStatusBeta, want: true},
		{name: "deprecated", value: opencode.V2ModelInfoStatusDeprecated, want: true},
		{name: "active", value: opencode.V2ModelInfoStatusActive, want: true},
		{name: "unknown", value: "unknown", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}
