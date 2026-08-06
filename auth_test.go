// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestAuthSet(t *testing.T) {
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
	_, err := client.Auth.Set(context.TODO(), "providerID", opencode.OAuthParam{
		Type:    opencode.F(opencode.OAuthParamTypeOAuth),
		Refresh: opencode.F("refresh_token"),
		Access:  opencode.F("access_token"),
		Expires: opencode.F[int64](3600),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuthRemove(t *testing.T) {
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
	_, err := client.Auth.Remove(context.TODO(), "providerID")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestOAuthParamRequiredOnlyNoOptionalFields verifies that OAuthParam only
// serialises required fields when optional fields are not set, and does not
// emit zero-value optional fields (accountId / enterpriseUrl).
//
// This is the core value of the Request/Response separation fix:
// optional fields must be absent from the wire when not provided.
//
// Run with: go test -run TestOAuthParamRequiredOnlyNoOptionalFields -v ./...
func TestOAuthParamRequiredOnlyNoOptionalFields(t *testing.T) {
	t.Parallel()

	p := opencode.OAuthParam{
		Type:    opencode.F(opencode.OAuthParamTypeOAuth),
		Refresh: opencode.F("refresh_tok"),
		Access:  opencode.F("access_tok"),
		Expires: opencode.F[int64](9999),
		// AccountID and EnterpriseURL deliberately omitted
	}

	got, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}

	var m map[string]any
	if err := json.Unmarshal(got, &m); err != nil {
		t.Fatalf("re-parse: %v", err)
	}

	// Required fields must be present
	for _, key := range []string{"type", "refresh", "access", "expires"} {
		if _, ok := m[key]; !ok {
			t.Errorf("expected required field %q to be present in JSON, got: %s", key, string(got))
		}
	}

	// Optional fields must be absent when not set
	for _, key := range []string{"accountId", "enterpriseUrl"} {
		if _, ok := m[key]; ok {
			t.Errorf("optional field %q must NOT appear in JSON when not set, got: %s", key, string(got))
		}
	}

	// Discriminator value must match OpenAPI
	if m["type"] != "oauth" {
		t.Errorf("type discriminator: got %q, want %q", m["type"], "oauth")
	}
}

// TestAuthParamVariantsSerialization verifies all three AuthParam variants
// serialise correctly with the right type discriminator and required fields.
//
// Run with: go test -run TestAuthParamVariantsSerialization -v ./...
func TestAuthParamVariantsSerialization(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		param    opencode.AuthParam
		wantType string
		wantKeys []string
		noKeys   []string
	}{
		{
			name: "OAuthParam_required_only",
			param: opencode.OAuthParam{
				Type:    opencode.F(opencode.OAuthParamTypeOAuth),
				Refresh: opencode.F("r"),
				Access:  opencode.F("a"),
				Expires: opencode.F[int64](1),
			},
			wantType: "oauth",
			wantKeys: []string{"type", "refresh", "access", "expires"},
			noKeys:   []string{"accountId", "enterpriseUrl"},
		},
		{
			name: "OAuthParam_with_optional",
			param: opencode.OAuthParam{
				Type:          opencode.F(opencode.OAuthParamTypeOAuth),
				Refresh:       opencode.F("r"),
				Access:        opencode.F("a"),
				Expires:       opencode.F[int64](1),
				AccountID:     opencode.F("acct-123"),
				EnterpriseURL: opencode.F("https://github.example.com"),
			},
			wantType: "oauth",
			wantKeys: []string{"type", "refresh", "access", "expires", "accountId", "enterpriseUrl"},
			noKeys:   nil,
		},
		{
			name: "ApiAuthParam_required_only",
			param: opencode.ApiAuthParam{
				Type: opencode.F(opencode.ApiAuthParamTypeAPI),
				Key:  opencode.F("sk-secret"),
			},
			wantType: "api",
			wantKeys: []string{"type", "key"},
			noKeys:   []string{"metadata"},
		},
		{
			name: "ApiAuthParam_with_metadata",
			param: opencode.ApiAuthParam{
				Type:     opencode.F(opencode.ApiAuthParamTypeAPI),
				Key:      opencode.F("sk-secret"),
				Metadata: opencode.F(map[string]string{"env": "prod"}),
			},
			wantType: "api",
			wantKeys: []string{"type", "key", "metadata"},
			noKeys:   nil,
		},
		{
			name: "WellKnownAuthParam_all_required",
			param: opencode.WellKnownAuthParam{
				Type:  opencode.F(opencode.WellKnownAuthParamTypeWellKnown),
				Key:   opencode.F("github"),
				Token: opencode.F("tok_xyz"),
			},
			wantType: "wellknown",
			wantKeys: []string{"type", "key", "token"},
			noKeys:   nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(tc.param)
			if err != nil {
				t.Fatalf("json.Marshal: %v", err)
			}

			var m map[string]any
			if err := json.Unmarshal(got, &m); err != nil {
				t.Fatalf("re-parse: %v", err)
			}

			// Check discriminator
			if m["type"] != tc.wantType {
				t.Errorf("type discriminator: got %q, want %q (json=%s)", m["type"], tc.wantType, string(got))
			}

			// Check required/expected keys present
			for _, key := range tc.wantKeys {
				if _, ok := m[key]; !ok {
					t.Errorf("expected field %q absent from JSON: %s", key, string(got))
				}
			}

			// Check optional keys absent when not set
			for _, key := range tc.noKeys {
				if _, ok := m[key]; ok {
					t.Errorf("optional field %q must NOT appear when not set: %s", key, string(got))
				}
			}
		})
	}
}

// TestOAuthParamImplementsAuthParam verifies the interface satisfaction at compile time.
func TestOAuthParamImplementsAuthParam(t *testing.T) {
	t.Parallel()

	// This test validates interface satisfaction at compile time — if any of
	// these assignments fail, the build fails before the test even runs.
	var _ opencode.AuthParam = opencode.OAuthParam{}
	var _ opencode.AuthParam = opencode.ApiAuthParam{}
	var _ opencode.AuthParam = opencode.WellKnownAuthParam{}

	// Also verify JSON output is non-empty
	p := opencode.ApiAuthParam{
		Type: opencode.F(opencode.ApiAuthParamTypeAPI),
		Key:  opencode.F("k"),
	}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	if !strings.Contains(string(b), `"type"`) {
		t.Errorf("expected type field in JSON output, got: %s", string(b))
	}
}
