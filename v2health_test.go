// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// TestV2HealthGet tests the health endpoint.
// Aligned with OpenAPI operationId "v2.health.get", GET /api/health.
// No query or body parameters.
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

// TestHealthV2InfoUnmarshal verifies deserialization of the HealthV2Info response.
// Aligned with OpenAPI schema "HealthInfo": required field "healthy" (boolean).
func TestHealthV2InfoUnmarshal(t *testing.T) {
	t.Run("healthy=true", func(t *testing.T) {
		raw := `{"healthy":true}`
		var h opencode.HealthV2Info
		if err := json.Unmarshal([]byte(raw), &h); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if !h.Healthy {
			t.Error("expected Healthy=true")
		}
		if h.JSON.RawJSON() != raw {
			t.Errorf("RawJSON mismatch: got %q, want %q", h.JSON.RawJSON(), raw)
		}
	})

	t.Run("healthy=false", func(t *testing.T) {
		raw := `{"healthy":false}`
		var h opencode.HealthV2Info
		if err := json.Unmarshal([]byte(raw), &h); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if h.Healthy {
			t.Error("expected Healthy=false")
		}
	})

	t.Run("unknown fields are tolerated via ExtraFields", func(t *testing.T) {
		raw := `{"healthy":true,"version":"1.2.3","extra_flag":42}`
		var h opencode.HealthV2Info
		if err := json.Unmarshal([]byte(raw), &h); err != nil {
			t.Fatalf("unmarshal with unknown fields: %v", err)
		}
		if !h.Healthy {
			t.Error("expected Healthy=true with unknown fields present")
		}
		// RawJSON preserves the full payload including unknown fields
		if h.JSON.RawJSON() != raw {
			t.Errorf("RawJSON mismatch: got %q, want %q", h.JSON.RawJSON(), raw)
		}
	})
}
