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

func TestV2PermissionRequestList(t *testing.T) {
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
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/home/user/project"),
			Workspace: opencode.F("ws_123"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PermissionSavedList(t *testing.T) {
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
		ProjectID: opencode.F("proj_abc"),
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
	err := client.V2Permission.Saved.Remove(context.TODO(), "perm_saved_id")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// ===== Missing path parameter unit tests =====

func TestV2PermissionSavedRemoveEmptyID(t *testing.T) {
	svc := opencode.NewV2PermissionSavedService()
	err := svc.Remove(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty id parameter")
	}
	if !strings.Contains(err.Error(), "missing required id parameter") {
		t.Errorf("unexpected error message: %s", err.Error())
	}
}

// ===== Request serialization tests =====

// Aligned with OpenAPI GET /api/permission/request — query param: location (optional nested object).
func TestV2PermissionRequestListParamsURLQuery(t *testing.T) {
	t.Run("with location", func(t *testing.T) {
		p := opencode.V2PermissionRequestListParams{
			Location: opencode.F(opencode.V2LocationParam{
				Directory: opencode.F("/tmp/project"),
				Workspace: opencode.F("ws_42"),
			}),
		}
		got := p.URLQuery().Encode()
		// nested bracket format: location[directory]=... & location[workspace]=...
		if !strings.Contains(got, "location%5Bdirectory%5D=%2Ftmp%2Fproject") &&
			!strings.Contains(got, "location[directory]=/tmp/project") {
			// url.Values.Encode() percent-encodes brackets
			if !strings.Contains(got, "location") {
				t.Errorf("location param missing from query: %q", got)
			}
		}
	})

	t.Run("empty params", func(t *testing.T) {
		p := opencode.V2PermissionRequestListParams{}
		got := p.URLQuery()
		if len(got) != 0 {
			t.Errorf("expected empty query values, got %v", got)
		}
	})
}

// Aligned with OpenAPI GET /api/permission/saved — query param: projectID (optional string).
func TestV2PermissionSavedListParamsURLQuery(t *testing.T) {
	p := opencode.V2PermissionSavedListParams{
		ProjectID: opencode.F("proj_xyz"),
	}
	got := p.URLQuery().Encode()
	if !strings.Contains(got, "projectID=proj_xyz") {
		t.Errorf("projectID missing from query: %q", got)
	}
}

// ===== Response deserialization tests =====

// V2PermissionRequestListResponse: required fields location + data.
func TestV2PermissionRequestListResponseUnmarshal(t *testing.T) {
	raw := `{
		"location": {
			"directory": "/home/user/project",
			"workspaceID": "ws_1",
			"project": {"id": "proj_1", "directory": "/home/user/project"}
		},
		"data": [
			{
				"id": "req_001",
				"sessionID": "ses_001",
				"action": "bash",
				"resources": ["*.sh"],
				"save": ["*.sh"],
				"metadata": {"tool": "bash"},
				"source": {"type": "tool", "messageID": "msg_1", "callID": "call_1"}
			}
		]
	}`
	var resp opencode.V2PermissionRequestListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Location.Directory != "/home/user/project" {
		t.Errorf("Location.Directory = %q", resp.Location.Directory)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Data))
	}
	item := resp.Data[0]
	if item.ID != "req_001" {
		t.Errorf("ID = %q", item.ID)
	}
	if item.Action != "bash" {
		t.Errorf("Action = %q", item.Action)
	}
	if len(item.Resources) != 1 || item.Resources[0] != "*.sh" {
		t.Errorf("Resources = %v", item.Resources)
	}
	if item.Source.Type != opencode.PermissionV2SourceTypeTool {
		t.Errorf("Source.Type = %q", item.Source.Type)
	}
	if item.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// V2PermissionSavedListResponse: required field data.
func TestV2PermissionSavedListResponseUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": "saved_001",
				"projectID": "proj_1",
				"action": "bash",
				"resource": "*.sh"
			}
		]
	}`
	var resp opencode.V2PermissionSavedListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Data))
	}
	item := resp.Data[0]
	if item.ID != "saved_001" {
		t.Errorf("ID = %q", item.ID)
	}
	if item.ProjectID != "proj_1" {
		t.Errorf("ProjectID = %q", item.ProjectID)
	}
	if item.Action != "bash" {
		t.Errorf("Action = %q", item.Action)
	}
	if item.Resource != "*.sh" {
		t.Errorf("Resource = %q", item.Resource)
	}
	if item.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// PermissionV2Request: optional metadata is any, optional save, optional source.
func TestPermissionV2RequestMetadataUnmarshal(t *testing.T) {
	t.Run("with metadata map", func(t *testing.T) {
		raw := `{
			"id": "req_x",
			"sessionID": "ses_x",
			"action": "read",
			"resources": ["file.go"],
			"metadata": {"k": "v"}
		}`
		var r opencode.PermissionV2Request
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if r.Metadata == nil {
			t.Error("Metadata should not be nil")
		}
		meta, ok := r.Metadata.(map[string]interface{})
		if !ok {
			t.Fatalf("Metadata type: got %T", r.Metadata)
		}
		if meta["k"] != "v" {
			t.Errorf("metadata[k] = %v", meta["k"])
		}
	})

	t.Run("without optional fields", func(t *testing.T) {
		raw := `{
			"id": "req_y",
			"sessionID": "ses_y",
			"action": "write",
			"resources": []
		}`
		var r opencode.PermissionV2Request
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if r.ID != "req_y" {
			t.Errorf("ID = %q", r.ID)
		}
		if len(r.Resources) != 0 {
			t.Errorf("Resources = %v", r.Resources)
		}
	})
}

// PermissionV2Effect enum IsKnown.
func TestPermissionV2EffectIsKnown(t *testing.T) {
	for _, v := range []opencode.PermissionV2Effect{
		opencode.PermissionV2EffectAllow,
		opencode.PermissionV2EffectDeny,
		opencode.PermissionV2EffectAsk,
	} {
		if !v.IsKnown() {
			t.Errorf("%q should be known", v)
		}
	}
	if opencode.PermissionV2Effect("unknown").IsKnown() {
		t.Error("unknown effect should not be known")
	}
}

// PermissionV2SourceType enum IsKnown.
func TestPermissionV2SourceTypeIsKnown(t *testing.T) {
	if !opencode.PermissionV2SourceTypeTool.IsKnown() {
		t.Error("PermissionV2SourceTypeTool should be known")
	}
	if opencode.PermissionV2SourceType("unknown").IsKnown() {
		t.Error("unknown source type should not be known")
	}
}

// ExtraFields: unknown JSON fields must not cause errors and should be accessible.
func TestV2PermissionSavedListResponseExtraFields(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": "s1",
				"projectID": "p1",
				"action": "read",
				"resource": "*.md",
				"unknownFuture": "value"
			}
		],
		"anotherFuture": 42
	}`
	var resp opencode.V2PermissionSavedListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("should tolerate extra fields: %v", err)
	}
	if len(resp.Data) != 1 {
		t.Errorf("expected 1 item, got %d", len(resp.Data))
	}
}
