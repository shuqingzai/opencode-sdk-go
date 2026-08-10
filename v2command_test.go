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

func TestV2CommandList(t *testing.T) {
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
	_, err := client.V2Command.List(context.TODO(), opencode.V2CommandListParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestV2CommandListParamsURLQuery verifies V2CommandListParams.URLQuery serializes correctly.
func TestV2CommandListParamsURLQuery(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		params   opencode.V2CommandListParams
		wantKeys []string
	}{
		{
			name:     "empty_params",
			params:   opencode.V2CommandListParams{},
			wantKeys: nil,
		},
		{
			name: "with_location",
			params: opencode.V2CommandListParams{
				Location: opencode.F(opencode.V2LocationParam{
					Workspace: opencode.F("ws1"),
				}),
			},
			wantKeys: []string{"location[workspace]"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			v := tc.params.URLQuery()
			if _, ok := any(v).(url.Values); !ok {
				t.Fatal("URLQuery() did not return url.Values")
			}
			for _, key := range tc.wantKeys {
				if v.Get(key) == "" {
					t.Errorf("expected query key %q to be set, got empty", key)
				}
			}
		})
	}
}

// TestV2CommandInfoUnmarshal verifies V2CommandInfo unmarshals all fields correctly.
func TestV2CommandInfoUnmarshal(t *testing.T) {
	t.Parallel()
	raw := `{"name":"deploy","template":"Deploy {{env}}","description":"Deploy to environment","agent":"deployer","model":{"id":"gpt-4","providerID":"openai"},"subtask":true}`
	var info opencode.V2CommandInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if info.Name != "deploy" {
		t.Errorf("Name: got %q, want %q", info.Name, "deploy")
	}
	if info.Template != "Deploy {{env}}" {
		t.Errorf("Template: got %q, want %q", info.Template, "Deploy {{env}}")
	}
	if info.Description != "Deploy to environment" {
		t.Errorf("Description: got %q, want %q", info.Description, "Deploy to environment")
	}
	if info.Agent != "deployer" {
		t.Errorf("Agent: got %q, want %q", info.Agent, "deployer")
	}
	if !info.Subtask {
		t.Error("Subtask: got false, want true")
	}
}
