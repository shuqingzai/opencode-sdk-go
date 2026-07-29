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

// TestV2LocationGetWithOptionalParams tests the v2 location get endpoint.
// Aligned with OpenAPI operationId "v2.location.get", GET /api/location.
// Query parameters: location (optional, nested with directory and workspace).
func TestV2LocationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Location.Get(context.TODO(), opencode.V2LocationGetParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
			Workspace: opencode.F("workspace"),
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

// TestV2LocationGetParamsURLQuery verifies query serialization for V2LocationGetParams.
// Aligned with OpenAPI GET /api/location location[directory] and location[workspace] params.
func TestV2LocationGetParamsURLQuery(t *testing.T) {
	t.Run("with full location", func(t *testing.T) {
		p := opencode.V2LocationGetParams{
			Location: opencode.F(opencode.V2LocationParam{
				Directory: opencode.F("/home/user/project"),
				Workspace: opencode.F("ws_123"),
			}),
		}
		got := p.URLQuery().Encode()
		if got == "" {
			t.Error("expected non-empty query string")
		}
	})

	t.Run("empty params produces no query", func(t *testing.T) {
		p := opencode.V2LocationGetParams{}
		got := p.URLQuery().Encode()
		if got != "" {
			t.Errorf("expected empty query, got %q", got)
		}
	})

	t.Run("directory only", func(t *testing.T) {
		p := opencode.V2LocationGetParams{
			Location: opencode.F(opencode.V2LocationParam{
				Directory: opencode.F("/mydir"),
			}),
		}
		got := p.URLQuery().Encode()
		if got == "" {
			t.Error("expected non-empty query for directory-only location")
		}
	})
}

// TestLocationInfoUnmarshal verifies deserialization of LocationInfo response.
// Aligned with OpenAPI schema "LocationInfo".
// Required: directory, project (LocationInfoProject).
// Optional: workspaceID.
func TestLocationInfoUnmarshal(t *testing.T) {
	t.Run("full location with workspaceID", func(t *testing.T) {
		raw := `{
			"directory": "/home/user/project",
			"workspaceID": "ws_abc123",
			"project": {
				"id": "proj_xyz",
				"directory": "/home/user/project"
			}
		}`
		var loc opencode.LocationInfo
		if err := json.Unmarshal([]byte(raw), &loc); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if loc.Directory != "/home/user/project" {
			t.Errorf("Directory = %q, want /home/user/project", loc.Directory)
		}
		if loc.WorkspaceID != "ws_abc123" {
			t.Errorf("WorkspaceID = %q, want ws_abc123", loc.WorkspaceID)
		}
		if loc.Project.ID != "proj_xyz" {
			t.Errorf("Project.ID = %q, want proj_xyz", loc.Project.ID)
		}
		if loc.Project.Directory != "/home/user/project" {
			t.Errorf("Project.Directory = %q, want /home/user/project", loc.Project.Directory)
		}
		if loc.JSON.RawJSON() == "" {
			t.Error("RawJSON should be preserved")
		}
	})

	t.Run("location without workspaceID", func(t *testing.T) {
		raw := `{
			"directory": "/tmp/myproj",
			"project": {"id": "proj_1", "directory": "/tmp/myproj"}
		}`
		var loc opencode.LocationInfo
		if err := json.Unmarshal([]byte(raw), &loc); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if loc.WorkspaceID != "" {
			t.Errorf("WorkspaceID should be empty when absent, got %q", loc.WorkspaceID)
		}
	})

	t.Run("unknown fields tolerated via ExtraFields", func(t *testing.T) {
		raw := `{
			"directory": "/",
			"project": {"id": "p", "directory": "/"},
			"future_field": "value"
		}`
		var loc opencode.LocationInfo
		if err := json.Unmarshal([]byte(raw), &loc); err != nil {
			t.Fatalf("unmarshal with unknown field: %v", err)
		}
		if loc.JSON.RawJSON() == "" {
			t.Error("RawJSON should be preserved with unknown fields")
		}
	})
}

// TestLocationInfoProjectUnmarshal verifies deserialization of LocationInfoProject.
// Aligned with OpenAPI LocationInfoProject schema: required id and directory.
func TestLocationInfoProjectUnmarshal(t *testing.T) {
	raw := `{"id":"proj_001","directory":"/workspace/project"}`
	var proj opencode.LocationInfoProject
	if err := json.Unmarshal([]byte(raw), &proj); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if proj.ID != "proj_001" {
		t.Errorf("ID = %q, want proj_001", proj.ID)
	}
	if proj.Directory != "/workspace/project" {
		t.Errorf("Directory = %q, want /workspace/project", proj.Directory)
	}
	if proj.JSON.RawJSON() == "" {
		t.Error("RawJSON should be preserved")
	}
}
