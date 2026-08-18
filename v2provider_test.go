// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestV2ProviderList(t *testing.T) {
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
	_, err := client.V2Provider.List(context.TODO(), opencode.V2ProviderListParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ProviderGet(t *testing.T) {
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
	_, err := client.V2Provider.Get(context.TODO(), "providerID", opencode.V2ProviderGetParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestV2ProviderInfoApiUnionDiscriminator verifies that V2ProviderInfoApiUnion
// correctly dispatches to V2ProviderInfoApiAisdk or V2ProviderInfoApiNative
// using the "type" discriminator field ("aisdk" / "native"), fixed in Task 3.
//
// Before the fix, init() used an empty discriminator "" which relied on exactness
// matching. With the fix, discriminator="type" + DiscriminatorValue ensures
// deterministic dispatch even when both variants share common fields.
//
// Run with: go test -run TestV2ProviderInfoApiUnionDiscriminator -v ./...
func TestV2ProviderInfoApiUnionDiscriminator(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name        string
		json        string
		wantAisdk   bool
		wantNative  bool
		wantPackage string // for aisdk variant
	}{
		{
			name:        "aisdk_variant",
			json:        `{"id":"openai","name":"OpenAI","api":{"type":"aisdk","package":"@ai-sdk/openai"},"request":{"headers":{},"body":{},"variant":""}}`,
			wantAisdk:   true,
			wantNative:  false,
			wantPackage: "@ai-sdk/openai",
		},
		{
			name:       "native_variant",
			json:       `{"id":"anthropic","name":"Anthropic","api":{"type":"native","settings":{}},"request":{"headers":{},"body":{},"variant":""}}`,
			wantAisdk:  false,
			wantNative: true,
		},
		{
			name:        "aisdk_with_url",
			json:        `{"id":"custom","name":"Custom","api":{"type":"aisdk","package":"@ai-sdk/custom","url":"https://api.custom.com"},"request":{"headers":{},"body":{},"variant":""}}`,
			wantAisdk:   true,
			wantNative:  false,
			wantPackage: "@ai-sdk/custom",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var info opencode.V2ProviderInfo
			if err := json.Unmarshal([]byte(tc.json), &info); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			union := info.AsAPIUnion()
			if union == nil {
				t.Fatal("AsAPIUnion() returned nil")
			}

			_, isAisdk := union.(opencode.V2ProviderInfoApiAisdk)
			_, isNative := union.(opencode.V2ProviderInfoApiNative)

			if isAisdk != tc.wantAisdk {
				t.Errorf("V2ProviderInfoApiAisdk assertion: got %v, want %v", isAisdk, tc.wantAisdk)
			}
			if isNative != tc.wantNative {
				t.Errorf("V2ProviderInfoApiNative assertion: got %v, want %v", isNative, tc.wantNative)
			}

			if tc.wantAisdk && tc.wantPackage != "" {
				aisdk, _ := union.(opencode.V2ProviderInfoApiAisdk)
				if aisdk.Package != tc.wantPackage {
					t.Errorf("Package: got %q, want %q", aisdk.Package, tc.wantPackage)
				}
			}
		})
	}
}

// TestV2ProviderInfoApiAisdkType verifies V2ProviderInfoApiAisdkType.IsKnown.
func TestV2ProviderInfoApiAisdkType(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2ProviderInfoApiAisdkType
		want  bool
	}{
		{name: "aisdk", value: opencode.V2ProviderInfoApiAisdkTypeAisdk, want: true},
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
}

// TestV2ProviderInfoApiNativeType verifies V2ProviderInfoApiNativeType.IsKnown.
func TestV2ProviderInfoApiNativeType(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2ProviderInfoApiNativeType
		want  bool
	}{
		{name: "native", value: opencode.V2ProviderInfoApiNativeTypeNative, want: true},
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
}

// TestV2ProviderInfoApiUnionNoCrossContamination verifies that aisdk JSON is never
// decoded as V2ProviderInfoApiNative, and vice versa (discriminator correctness).
func TestV2ProviderInfoApiUnionNoCrossContamination(t *testing.T) {
	t.Parallel()

	aisdkJSON := `{"id":"openai","name":"OpenAI","api":{"type":"aisdk","package":"@ai-sdk/openai"},"request":{"headers":{},"body":{},"variant":""}}`
	nativeJSON := `{"id":"anthropic","name":"Anthropic","api":{"type":"native","settings":{}},"request":{"headers":{},"body":{},"variant":""}}`

	var aisdk opencode.V2ProviderInfo
	if err := json.Unmarshal([]byte(aisdkJSON), &aisdk); err != nil {
		t.Fatalf("aisdk unmarshal: %v", err)
	}
	aisdkUnion := aisdk.AsAPIUnion()
	if _, bad := aisdkUnion.(opencode.V2ProviderInfoApiNative); bad {
		t.Error("CROSS-CONTAMINATION: aisdk JSON decoded as V2ProviderInfoApiNative")
	}
	if _, ok := aisdkUnion.(opencode.V2ProviderInfoApiAisdk); !ok {
		t.Errorf("aisdk JSON: AsAPIUnion() type assertion to V2ProviderInfoApiAisdk failed, got %T", aisdkUnion)
	}

	var native opencode.V2ProviderInfo
	if err := json.Unmarshal([]byte(nativeJSON), &native); err != nil {
		t.Fatalf("native unmarshal: %v", err)
	}
	nativeUnion := native.AsAPIUnion()
	if _, bad := nativeUnion.(opencode.V2ProviderInfoApiAisdk); bad {
		t.Error("CROSS-CONTAMINATION: native JSON decoded as V2ProviderInfoApiAisdk")
	}
	if _, ok := nativeUnion.(opencode.V2ProviderInfoApiNative); !ok {
		t.Errorf("native JSON: AsAPIUnion() type assertion to V2ProviderInfoApiNative failed, got %T", nativeUnion)
	}
}

// TestV2ProviderInfoAPIUnionRouting verifies (using the post-rename
// [V2ProviderInfoAPI] symbols, see Phase 2 "Api"->"API" initialism rename) that
// [V2ProviderInfoAPIUnion] correctly routes an "aisdk" payload to
// [V2ProviderInfoAPIAisdk] and a "native" payload to [V2ProviderInfoAPINative].
// Unlike ModelApi (see TestV2ModelInfoAPIUnionRouting), the OpenAPI ProviderApi
// variants carry no "id" property, so only Type/Package/URL/Settings are asserted.
func TestV2ProviderInfoAPIUnionRouting(t *testing.T) {
	t.Parallel()

	t.Run("aisdk", func(t *testing.T) {
		t.Parallel()
		raw := `{"id":"openai","name":"OpenAI","api":{"type":"aisdk","package":"@ai-sdk/openai","url":"https://ai-sdk.dev"},"request":{"headers":{},"body":{},"variant":""}}`
		var info opencode.V2ProviderInfo
		if err := json.Unmarshal([]byte(raw), &info); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		aisdk, ok := info.API.AsUnion().(opencode.V2ProviderInfoAPIAisdk)
		if !ok {
			t.Fatalf("AsUnion() type=%T, want V2ProviderInfoAPIAisdk", info.API.AsUnion())
		}
		if aisdk.Type != opencode.V2ProviderInfoAPIAisdkTypeAisdk {
			t.Errorf("Type: got %q, want %q", aisdk.Type, opencode.V2ProviderInfoAPIAisdkTypeAisdk)
		}
		if aisdk.Package != "@ai-sdk/openai" {
			t.Errorf("Package: got %q, want %q", aisdk.Package, "@ai-sdk/openai")
		}
		if aisdk.URL != "https://ai-sdk.dev" {
			t.Errorf("URL: got %q, want %q", aisdk.URL, "https://ai-sdk.dev")
		}
	})

	t.Run("native", func(t *testing.T) {
		t.Parallel()
		raw := `{"id":"anthropic","name":"Anthropic","api":{"type":"native","url":"https://api.anthropic.com","settings":{"apiKey":"sk-456"}},"request":{"headers":{},"body":{},"variant":""}}`
		var info opencode.V2ProviderInfo
		if err := json.Unmarshal([]byte(raw), &info); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		native, ok := info.API.AsUnion().(opencode.V2ProviderInfoAPINative)
		if !ok {
			t.Fatalf("AsUnion() type=%T, want V2ProviderInfoAPINative", info.API.AsUnion())
		}
		if native.Type != opencode.V2ProviderInfoAPINativeTypeNative {
			t.Errorf("Type: got %q, want %q", native.Type, opencode.V2ProviderInfoAPINativeTypeNative)
		}
		if native.URL != "https://api.anthropic.com" {
			t.Errorf("URL: got %q, want %q", native.URL, "https://api.anthropic.com")
		}
		if got := native.Settings["apiKey"]; got != "sk-456" {
			t.Errorf("Settings[apiKey]: got %v, want %q", got, "sk-456")
		}
	})
}

// TestV2ProviderInfoAPIWireCompatRegression is the wire-format regression required
// by Phase 2: after the "Api"->"API" Go identifier rename, the `json:"api"` wire
// key must still populate the (renamed) [V2ProviderInfo.API] field, and RawJSON()
// must still return the exact original bytes for the "api" sub-object.
func TestV2ProviderInfoAPIWireCompatRegression(t *testing.T) {
	t.Parallel()

	apiJSON := `{"type":"aisdk","package":"@ai-sdk/openai","url":"https://raw.dev"}`
	raw := `{"id":"openai","name":"OpenAI","api":` + apiJSON + `,"request":{"headers":{},"body":{},"variant":""}}`

	var info opencode.V2ProviderInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	// json:"api" wire key must still populate the renamed API field.
	if info.API.AsUnion() == nil {
		t.Fatal("API.AsUnion() returned nil (wire key \"api\" not routed to renamed API field)")
	}
	aisdk, ok := info.API.AsUnion().(opencode.V2ProviderInfoAPIAisdk)
	if !ok {
		t.Fatalf("API.AsUnion() type=%T, want V2ProviderInfoAPIAisdk", info.API.AsUnion())
	}
	if aisdk.Package != "@ai-sdk/openai" {
		t.Errorf("Package: got %q, want %q", aisdk.Package, "@ai-sdk/openai")
	}

	got := info.API.JSON.RawJSON()
	if got != apiJSON {
		t.Errorf("API.JSON.RawJSON() = %q, want %q", got, apiJSON)
	}
	if !json.Valid([]byte(got)) {
		t.Errorf("API.JSON.RawJSON() = %q, not valid JSON", got)
	}

	// The outer V2ProviderInfo.JSON.RawJSON() must also still hold the complete
	// original payload untouched by the identifier rename.
	if outerRaw := info.JSON.RawJSON(); outerRaw != raw {
		t.Errorf("V2ProviderInfo.JSON.RawJSON() = %q, want %q", outerRaw, raw)
	}
}

// TestV2ProviderInfoAPIDeprecatedAliasCompileTime is a compile-time-checked
// assertion that every "Api"-named symbol removed from v2provider.go in Phase 2
// is still usable as a [Deprecated] alias of its "API"-named replacement. If any
// alias were dropped, this file would fail to compile.
func TestV2ProviderInfoAPIDeprecatedAliasCompileTime(t *testing.T) {
	t.Parallel()

	// Deprecated type aliases must be identical (assignable both ways) to the new types.
	var _ opencode.V2ProviderInfoApi = opencode.V2ProviderInfoAPI{}
	var _ opencode.V2ProviderInfoAPI = opencode.V2ProviderInfoApi{}
	var _ opencode.V2ProviderInfoApiType = opencode.V2ProviderInfoAPIType("aisdk")
	var _ opencode.V2ProviderInfoApiUnion = opencode.V2ProviderInfoAPIAisdk{}
	var _ opencode.V2ProviderInfoApiAisdk = opencode.V2ProviderInfoAPIAisdk{}
	var _ opencode.V2ProviderInfoApiNative = opencode.V2ProviderInfoAPINative{}
	var _ opencode.V2ProviderInfoApiAisdkType = opencode.V2ProviderInfoAPIAisdkType("aisdk")
	var _ opencode.V2ProviderInfoApiNativeType = opencode.V2ProviderInfoAPINativeType("native")

	// Deprecated enum constants must equal their new-named counterparts.
	if opencode.V2ProviderInfoApiTypeAisdk != opencode.V2ProviderInfoAPITypeAisdk {
		t.Errorf("V2ProviderInfoApiTypeAisdk != V2ProviderInfoAPITypeAisdk")
	}
	if opencode.V2ProviderInfoApiTypeNative != opencode.V2ProviderInfoAPITypeNative {
		t.Errorf("V2ProviderInfoApiTypeNative != V2ProviderInfoAPITypeNative")
	}
	if opencode.V2ProviderInfoApiAisdkTypeAisdk != opencode.V2ProviderInfoAPIAisdkTypeAisdk {
		t.Errorf("V2ProviderInfoApiAisdkTypeAisdk != V2ProviderInfoAPIAisdkTypeAisdk")
	}
	if opencode.V2ProviderInfoApiNativeTypeNative != opencode.V2ProviderInfoAPINativeTypeNative {
		t.Errorf("V2ProviderInfoApiNativeTypeNative != V2ProviderInfoAPINativeTypeNative")
	}

	// Deprecated constants must still satisfy IsKnown() via the shared underlying type.
	if !opencode.V2ProviderInfoApiTypeAisdk.IsKnown() {
		t.Error("V2ProviderInfoApiTypeAisdk.IsKnown() = false, want true")
	}
	if !opencode.V2ProviderInfoApiAisdkTypeAisdk.IsKnown() {
		t.Error("V2ProviderInfoApiAisdkTypeAisdk.IsKnown() = false, want true")
	}
	if !opencode.V2ProviderInfoApiNativeTypeNative.IsKnown() {
		t.Error("V2ProviderInfoApiNativeTypeNative.IsKnown() = false, want true")
	}
}

// TestV2ProviderInfoAPITypeIsKnown verifies the post-rename [V2ProviderInfoAPIType],
// [V2ProviderInfoAPIAisdkType] and [V2ProviderInfoAPINativeType] IsKnown()
// implementations report true for known enum values and false for unknown ones.
func TestV2ProviderInfoAPITypeIsKnown(t *testing.T) {
	t.Parallel()

	typeCases := []struct {
		name  string
		value opencode.V2ProviderInfoAPIType
		want  bool
	}{
		{name: "aisdk", value: opencode.V2ProviderInfoAPITypeAisdk, want: true},
		{name: "native", value: opencode.V2ProviderInfoAPITypeNative, want: true},
		{name: "unknown", value: "unknown", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range typeCases {
		t.Run("V2ProviderInfoAPIType/"+tc.name, func(t *testing.T) {
			t.Parallel()
			if got := tc.value.IsKnown(); got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}

	aisdkCases := []struct {
		name  string
		value opencode.V2ProviderInfoAPIAisdkType
		want  bool
	}{
		{name: "aisdk", value: opencode.V2ProviderInfoAPIAisdkTypeAisdk, want: true},
		{name: "unknown", value: "unknown", want: false},
	}
	for _, tc := range aisdkCases {
		t.Run("V2ProviderInfoAPIAisdkType/"+tc.name, func(t *testing.T) {
			t.Parallel()
			if got := tc.value.IsKnown(); got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}

	nativeCases := []struct {
		name  string
		value opencode.V2ProviderInfoAPINativeType
		want  bool
	}{
		{name: "native", value: opencode.V2ProviderInfoAPINativeTypeNative, want: true},
		{name: "unknown", value: "unknown", want: false},
	}
	for _, tc := range nativeCases {
		t.Run("V2ProviderInfoAPINativeType/"+tc.name, func(t *testing.T) {
			t.Parallel()
			if got := tc.value.IsKnown(); got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}
