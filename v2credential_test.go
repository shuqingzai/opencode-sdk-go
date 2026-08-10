// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

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
	err := client.V2Credential.Update(context.TODO(), "credentialID", opencode.V2CredentialUpdateParams{
		Label: opencode.F("label"),
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
	err := client.V2Credential.Remove(context.TODO(), "credentialID", opencode.V2CredentialRemoveParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestV2CredentialUpdateParamsBodyQuerySeparation verifies that V2CredentialUpdateParams
// correctly separates body fields (json tag) from query fields (query tag).
// MarshalJSON should only include body fields; URLQuery should only include query fields.
func TestV2CredentialUpdateParamsBodyQuerySeparation(t *testing.T) {
	t.Parallel()

	params := opencode.V2CredentialUpdateParams{
		Label: opencode.F("my-label"),
		Location: opencode.F(opencode.V2LocationParam{
			Workspace: opencode.F("ws1"),
		}),
	}

	// Test MarshalJSON: should include "label", should NOT include "location"
	t.Run("marshal_json_body_only", func(t *testing.T) {
		t.Parallel()
		data, err := json.Marshal(params)
		if err != nil {
			t.Fatalf("json.Marshal: %v", err)
		}
		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			t.Fatalf("json.Unmarshal result: %v", err)
		}
		if _, ok := m["label"]; !ok {
			t.Error("MarshalJSON should contain 'label' field")
		}
		if _, ok := m["location"]; ok {
			t.Error("MarshalJSON must NOT contain 'location' (query-only) field")
		}
	})

	// Test URLQuery: should include "location", should NOT include "label"
	t.Run("url_query_only", func(t *testing.T) {
		t.Parallel()
		v := params.URLQuery()
		if _, ok := any(v).(url.Values); !ok {
			t.Fatal("URLQuery() did not return url.Values")
		}
		// location should appear in query
		if v.Get("location[workspace]") == "" {
			t.Error("URLQuery should contain 'location[workspace]'")
		}
		// label should NOT appear in query
		if v.Get("label") != "" {
			t.Error("URLQuery must NOT contain 'label' (body-only) field")
		}
	})
}

// TestV2CredentialRemoveParamsURLQuery verifies V2CredentialRemoveParams.URLQuery.
func TestV2CredentialRemoveParamsURLQuery(t *testing.T) {
	t.Parallel()
	params := opencode.V2CredentialRemoveParams{
		Location: opencode.F(opencode.V2LocationParam{
			Workspace: opencode.F("ws2"),
		}),
	}
	v := params.URLQuery()
	if v.Get("location[workspace]") == "" {
		t.Error("URLQuery should contain 'location[workspace]'")
	}
}
