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

func TestV2HealthGet(t *testing.T) {
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
	_, err := client.V2Health.Get(context.TODO())
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestHealthV2InfoUnmarshal verifies HealthV2Info unmarshal from JSON.
func TestHealthV2InfoUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name    string
		json    string
		healthy bool
	}{
		{name: "healthy_true", json: `{"healthy":true}`, healthy: true},
		{name: "healthy_false", json: `{"healthy":false}`, healthy: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var h opencode.HealthV2Info
			if err := json.Unmarshal([]byte(tc.json), &h); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			if h.Healthy != tc.healthy {
				t.Errorf("Healthy: got %v, want %v", h.Healthy, tc.healthy)
			}
		})
	}
}
