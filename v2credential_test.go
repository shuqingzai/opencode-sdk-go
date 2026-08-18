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

// TestV2CredentialUpdateParamsMarshalJSON documents and locks in the final
// ruling on V2CredentialUpdateParams from the Phase 2 brief: unlike the three
// v2integration.go Params fixed in this batch (bare `Body XxxParamsBody
// json:"-"` + unconditional apijson.MarshalRoot(r.Body)), V2CredentialUpdateParams
// already inlines its one body field (Label) directly on the Params struct as
// param.Field[string] `json:"label,required"`, with MarshalJSON delegating to
// apijson.MarshalRoot(r) — the exact "inline golden standard" pattern of
// session.go's SessionPromptParams (and, per the D4 report, app.go's
// AppLogParams). That pattern has no bare/unwrapped Body field, so there is no
// "style 2" defect to fix here, and no code change was made to
// V2CredentialUpdateParams in this batch.
//
// The nil-vs-"{}" distinction that motivated the sync.go "style 1" fix for the
// three v2integration.go Params does NOT apply to the inline pattern: an
// inline Params struct always serializes to at least "{}" (never nil) via
// apijson.MarshalRoot(r), exactly like SessionPromptParams and AppLogParams
// when none of their optional fields are set. This is asserted directly below.
func TestV2CredentialUpdateParamsMarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("LabelSet_SerializesLabel", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2CredentialUpdateParams{
			Label: opencode.F("mylabel"),
		}
		data, err := params.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v", err)
		}
		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			t.Fatalf("json.Unmarshal result: %v", err)
		}
		if m["label"] != "mylabel" {
			t.Errorf("label: got %v, want mylabel", m["label"])
		}
		if len(m) != 1 {
			t.Errorf("expected exactly 1 field (label), got %v", m)
		}
	})

	// LabelUnset_ReturnsEmptyObject documents actual behavior: because Label is
	// inlined (not wrapped in a separate Body struct), MarshalJSON always calls
	// apijson.MarshalRoot(r) unconditionally and returns "{}" — never nil — when
	// Label is unset. Per OpenAPI, v2.credential.update's requestBody is
	// `required: true` with `label` itself `required` inside the schema, so a
	// valid caller must always set Label; "{}" here is the same "all-optional-
	// fields-unset" degenerate case documented for SessionPromptParams/
	// AppLogParams, not a behavioral regression to fix.
	t.Run("LabelUnset_ReturnsEmptyObject", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2CredentialUpdateParams{}
		data, err := params.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v", err)
		}
		if data == nil {
			t.Fatal("expected non-nil output (inline pattern always calls apijson.MarshalRoot(r))")
		}
		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			t.Fatalf("json.Unmarshal result: %v", err)
		}
		if len(m) != 0 {
			t.Errorf("expected empty object {} when Label is unset, got %v", m)
		}
	})
}
