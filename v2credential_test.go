// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// ===== Prism / live-server tests =====

func TestV2CredentialUpdate(t *testing.T) {
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
	err := client.V2Credential.Update(context.TODO(), "cred_abc", opencode.V2CredentialUpdateParams{
		Label:    opencode.F("My GitHub Credential"),
		Location: opencode.F(opencode.V2LocationParam{Directory: opencode.F("/home/user")}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2CredentialRemove(t *testing.T) {
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
	err := client.V2Credential.Remove(context.TODO(), "cred_abc", opencode.V2CredentialRemoveParams{
		Location: opencode.F(opencode.V2LocationParam{Directory: opencode.F("/home/user")}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// ===== Missing path parameter unit tests =====

func TestV2CredentialUpdateEmptyCredentialID(t *testing.T) {
	svc := opencode.NewV2CredentialService()
	err := svc.Update(context.Background(), "", opencode.V2CredentialUpdateParams{
		Label: opencode.F("label"),
	})
	if err == nil {
		t.Fatal("expected error for empty credentialID")
	}
	if !strings.Contains(err.Error(), "missing required credentialID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2CredentialRemoveEmptyCredentialID(t *testing.T) {
	svc := opencode.NewV2CredentialService()
	err := svc.Remove(context.Background(), "", opencode.V2CredentialRemoveParams{})
	if err == nil {
		t.Fatal("expected error for empty credentialID")
	}
	if !strings.Contains(err.Error(), "missing required credentialID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

// ===== Request serialization tests =====

// Aligned with OpenAPI PATCH /api/credential/{credentialID}:
// body required: label; query: location (optional nested object).
func TestV2CredentialUpdateParamsBody(t *testing.T) {
	t.Run("required label in body", func(t *testing.T) {
		p := opencode.V2CredentialUpdateParams{
			Label: opencode.F("My Credential"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"label":"My Credential"`) {
			t.Errorf("label missing from body: %s", got)
		}
		// location is a query param — must not appear in body
		if strings.Contains(got, "location") {
			t.Errorf("location leaked into body: %s", got)
		}
	})

	t.Run("label is required", func(t *testing.T) {
		p := opencode.V2CredentialUpdateParams{
			Label: opencode.F("new-label"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(b), `"label"`) {
			t.Errorf("label not found in JSON: %s", string(b))
		}
	})
}

// V2CredentialUpdateParams.URLQuery: location nested bracket format.
func TestV2CredentialUpdateParamsURLQuery(t *testing.T) {
	p := opencode.V2CredentialUpdateParams{
		Label: opencode.F("x"),
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/my/dir"),
			Workspace: opencode.F("ws_99"),
		}),
	}
	encoded := p.URLQuery().Encode()
	if !strings.Contains(encoded, "location") {
		t.Errorf("location missing from query: %q", encoded)
	}
}

// V2CredentialRemoveParams.URLQuery: location is optional.
func TestV2CredentialRemoveParamsURLQuery(t *testing.T) {
	t.Run("with location", func(t *testing.T) {
		p := opencode.V2CredentialRemoveParams{
			Location: opencode.F(opencode.V2LocationParam{
				Directory: opencode.F("/proj"),
			}),
		}
		encoded := p.URLQuery().Encode()
		if !strings.Contains(encoded, "location") {
			t.Errorf("location missing: %q", encoded)
		}
	})
	t.Run("without location", func(t *testing.T) {
		p := opencode.V2CredentialRemoveParams{}
		got := p.URLQuery()
		if len(got) != 0 {
			t.Errorf("expected empty query, got %v", got)
		}
	})
}
