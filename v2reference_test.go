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

// TestV2ReferenceListWithOptionalParams tests the v2 reference list endpoint.
// Aligned with OpenAPI operationId "v2.reference.list", GET /api/reference.
// Query parameters: location (optional, nested with directory and workspace).
func TestV2ReferenceListWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Reference.List(context.TODO(), opencode.V2ReferenceListParams{
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

// TestV2ReferenceListParamsURLQuery verifies query serialization for V2ReferenceListParams.
func TestV2ReferenceListParamsURLQuery(t *testing.T) {
	t.Run("with full location", func(t *testing.T) {
		p := opencode.V2ReferenceListParams{
			Location: opencode.F(opencode.V2LocationParam{
				Directory: opencode.F("mydir"),
				Workspace: opencode.F("mywsp"),
			}),
		}
		got := p.URLQuery().Encode()
		if got == "" {
			t.Error("expected non-empty query string")
		}
	})

	t.Run("empty params produces no query", func(t *testing.T) {
		p := opencode.V2ReferenceListParams{}
		got := p.URLQuery().Encode()
		if got != "" {
			t.Errorf("expected empty query, got %q", got)
		}
	})
}

// TestV2ReferenceListResponseUnmarshal verifies deserialization of V2ReferenceListResponse.
// Aligned with OpenAPI GET /api/reference response schema.
// Required fields: location (LocationInfo), data ([]V2ReferenceInfo).
func TestV2ReferenceListResponseUnmarshal(t *testing.T) {
	t.Run("full response with local and git references", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/home/user/project",
				"workspaceID": "ws_1",
				"project": {"id": "proj_1", "directory": "/home/user/project"}
			},
			"data": [
				{
					"name": "local-docs",
					"path": "/docs/reference",
					"description": "Local documentation reference",
					"hidden": false,
					"source": {
						"type": "local",
						"path": "/docs/reference",
						"description": "Local docs"
					}
				},
				{
					"name": "git-stdlib",
					"path": "/tmp/stdlib",
					"description": "Standard library reference from git",
					"hidden": false,
					"source": {
						"type": "git",
						"repository": "https://github.com/example/stdlib",
						"branch": "main",
						"description": "Standard library"
					}
				}
			]
		}`
		var resp opencode.V2ReferenceListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if resp.Location.Directory != "/home/user/project" {
			t.Errorf("Location.Directory = %q, want /home/user/project", resp.Location.Directory)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("Data len = %d, want 2", len(resp.Data))
		}
		ref0 := resp.Data[0]
		if ref0.Name != "local-docs" {
			t.Errorf("ref0.Name = %q, want local-docs", ref0.Name)
		}
		if ref0.Path != "/docs/reference" {
			t.Errorf("ref0.Path = %q, want /docs/reference", ref0.Path)
		}
		if ref0.Description != "Local documentation reference" {
			t.Errorf("ref0.Description = %q", ref0.Description)
		}
		if ref0.Hidden {
			t.Error("ref0.Hidden should be false")
		}
		ref1 := resp.Data[1]
		if ref1.Name != "git-stdlib" {
			t.Errorf("ref1.Name = %q, want git-stdlib", ref1.Name)
		}
		if resp.JSON.RawJSON() == "" {
			t.Error("RawJSON should be preserved")
		}
	})

	t.Run("empty data array", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/",
				"project": {"id": "p", "directory": "/"}
			},
			"data": []
		}`
		var resp opencode.V2ReferenceListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(resp.Data) != 0 {
			t.Errorf("expected 0 references, got %d", len(resp.Data))
		}
	})
}

// TestV2ReferenceInfoSourceUnionLocal verifies the local branch of ReferenceSource union.
// Aligned with OpenAPI ReferenceLocalSource schema: type="local", path required.
func TestV2ReferenceInfoSourceUnionLocal(t *testing.T) {
	raw := `{
		"name": "my-docs",
		"path": "/project/docs",
		"description": "Project docs",
		"hidden": false,
		"source": {
			"type": "local",
			"path": "/project/docs",
			"description": "Local docs reference",
			"hidden": false
		}
	}`
	var ref opencode.V2ReferenceInfo
	if err := json.Unmarshal([]byte(raw), &ref); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	u := ref.Source.AsUnion()
	if u == nil {
		t.Fatal("AsUnion() should not be nil for local source")
	}
	local, ok := u.(opencode.ReferenceLocalSource)
	if !ok {
		t.Fatalf("expected ReferenceLocalSource, got %T", u)
	}
	if local.Type != opencode.ReferenceLocalSourceTypeLocal {
		t.Errorf("local.Type = %q, want local", local.Type)
	}
	if local.Path != "/project/docs" {
		t.Errorf("local.Path = %q, want /project/docs", local.Path)
	}
	if local.Description != "Local docs reference" {
		t.Errorf("local.Description = %q, want Local docs reference", local.Description)
	}
	if local.Hidden {
		t.Error("local.Hidden should be false")
	}
}

// TestV2ReferenceInfoSourceUnionGit verifies the git branch of ReferenceSource union.
// Aligned with OpenAPI ReferenceGitSource schema: type="git", repository required, branch optional.
func TestV2ReferenceInfoSourceUnionGit(t *testing.T) {
	raw := `{
		"name": "stdlib-ref",
		"path": "/tmp/stdlib",
		"source": {
			"type": "git",
			"repository": "https://github.com/example/stdlib",
			"branch": "v2",
			"description": "Stdlib v2 docs",
			"hidden": true
		}
	}`
	var ref opencode.V2ReferenceInfo
	if err := json.Unmarshal([]byte(raw), &ref); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	u := ref.Source.AsUnion()
	if u == nil {
		t.Fatal("AsUnion() should not be nil for git source")
	}
	git, ok := u.(opencode.ReferenceGitSource)
	if !ok {
		t.Fatalf("expected ReferenceGitSource, got %T", u)
	}
	if git.Type != opencode.ReferenceGitSourceTypeGit {
		t.Errorf("git.Type = %q, want git", git.Type)
	}
	if git.Repository != "https://github.com/example/stdlib" {
		t.Errorf("git.Repository = %q, want https://github.com/example/stdlib", git.Repository)
	}
	if git.Branch != "v2" {
		t.Errorf("git.Branch = %q, want v2", git.Branch)
	}
	if git.Description != "Stdlib v2 docs" {
		t.Errorf("git.Description = %q, want Stdlib v2 docs", git.Description)
	}
	if !git.Hidden {
		t.Error("git.Hidden should be true")
	}
}

// TestReferenceGitSourceWithoutBranch verifies that branch is optional in ReferenceGitSource.
// Aligned with OpenAPI ReferenceGitSource: branch not required.
func TestReferenceGitSourceWithoutBranch(t *testing.T) {
	raw := `{"type":"git","repository":"https://github.com/example/repo"}`
	var gs opencode.ReferenceGitSource
	if err := json.Unmarshal([]byte(raw), &gs); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if gs.Branch != "" {
		t.Errorf("Branch should be empty when absent, got %q", gs.Branch)
	}
	if gs.Repository != "https://github.com/example/repo" {
		t.Errorf("Repository = %q", gs.Repository)
	}
}

// TestReferenceLocalSourceTypeIsKnown verifies ReferenceLocalSourceType enum.
func TestReferenceLocalSourceTypeIsKnown(t *testing.T) {
	if !opencode.ReferenceLocalSourceTypeLocal.IsKnown() {
		t.Error("local should be known")
	}
	if opencode.ReferenceLocalSourceType("unknown").IsKnown() {
		t.Error("unknown should not be known")
	}
}

// TestReferenceGitSourceTypeIsKnown verifies ReferenceGitSourceType enum.
func TestReferenceGitSourceTypeIsKnown(t *testing.T) {
	if !opencode.ReferenceGitSourceTypeGit.IsKnown() {
		t.Error("git should be known")
	}
	if opencode.ReferenceGitSourceType("unknown").IsKnown() {
		t.Error("unknown should not be known")
	}
}

// TestV2ReferenceInfoUnknownFields verifies that unknown fields are tolerated.
func TestV2ReferenceInfoUnknownFields(t *testing.T) {
	raw := `{
		"name": "x",
		"path": "/x",
		"source": {"type": "local", "path": "/x"},
		"future_field": "value"
	}`
	var ref opencode.V2ReferenceInfo
	if err := json.Unmarshal([]byte(raw), &ref); err != nil {
		t.Fatalf("unmarshal with unknown field: %v", err)
	}
	if ref.Name != "x" {
		t.Errorf("Name = %q, want x", ref.Name)
	}
}
