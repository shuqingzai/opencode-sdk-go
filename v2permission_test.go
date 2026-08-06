// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestV2PermissionRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Permission.Request.List(context.TODO(), opencode.V2PermissionRequestListParams{
		Location: opencode.F(opencode.V2LocationParam{Directory: opencode.F("directory")}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PermissionSavedListWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Permission.Saved.List(context.TODO(), opencode.V2PermissionSavedListParams{
		ProjectID: opencode.F("projectID"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PermissionSavedRemove(t *testing.T) {
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
	err := client.V2Permission.Saved.Remove(context.TODO(), "permissionID")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PermissionSavedRemoveRequiresID(t *testing.T) {
	s := opencode.NewV2PermissionSavedService()
	if err := s.Remove(context.Background(), ""); err == nil {
		t.Error("expected error for empty id")
	}
}

func TestV2PermissionRequestListResponseUnmarshal(t *testing.T) {
	jsonStr := `{
		"location": {"directory": "/tmp", "project": {"id": "p1", "directory": "/tmp"}},
		"data": [
			{"id": "req_1", "sessionID": "ses_1", "action": "bash", "resources": ["fs"], "save": ["once"], "source": {"type": "tool", "messageID": "msg_1", "callID": "call_1"}}
		]
	}`
	var resp opencode.V2PermissionRequestListResponse
	if err := resp.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Data))
	}
	if resp.Data[0].ID != "req_1" {
		t.Errorf("expected req_1, got %s", resp.Data[0].ID)
	}
	if resp.Data[0].Source.Type != opencode.PermissionV2SourceTypeTool {
		t.Errorf("expected tool, got %s", resp.Data[0].Source.Type)
	}
}
