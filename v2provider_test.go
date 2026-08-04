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
