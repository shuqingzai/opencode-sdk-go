// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
)

// TestV2ReferenceInfoUnmarshalGitSource verifies that V2ReferenceInfo correctly
// deserialises a reference entry whose source field is a git source, and that
// AsSourceUnion() returns the concrete *ReferenceGitSource variant.
func TestV2ReferenceInfoUnmarshalGitSource(t *testing.T) {
	t.Parallel()

	raw := `{
		"name": "my-git-ref",
		"path": "refs/git/my-repo",
		"description": "A git reference",
		"hidden": false,
		"source": {
			"type": "git",
			"repository": "https://github.com/example/repo.git",
			"branch": "main",
			"description": "upstream repo"
		}
	}`

	var info opencode.V2ReferenceInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	// Verify scalar fields
	if info.Name != "my-git-ref" {
		t.Errorf("Name: got %q, want %q", info.Name, "my-git-ref")
	}
	if info.Path != "refs/git/my-repo" {
		t.Errorf("Path: got %q, want %q", info.Path, "refs/git/my-repo")
	}

	// Verify Source bearer is populated
	if info.JSON.Source.IsMissing() {
		t.Fatal("Source field is missing after UnmarshalJSON")
	}

	// Verify AsSourceUnion returns the concrete git type
	union := info.AsSourceUnion()
	if union == nil {
		t.Fatal("AsSourceUnion() returned nil")
	}
	gitSrc, ok := union.(opencode.ReferenceGitSource)
	if !ok {
		t.Fatalf("AsSourceUnion() runtime type: got %T, want ReferenceGitSource", union)
	}
	if gitSrc.Type != "git" {
		t.Errorf("ReferenceGitSource.Type: got %q, want %q", gitSrc.Type, "git")
	}
	if gitSrc.Repository != "https://github.com/example/repo.git" {
		t.Errorf("ReferenceGitSource.Repository: got %q, want %q", gitSrc.Repository, "https://github.com/example/repo.git")
	}
	if gitSrc.Branch != "main" {
		t.Errorf("ReferenceGitSource.Branch: got %q, want %q", gitSrc.Branch, "main")
	}
}

// TestV2ReferenceInfoUnmarshalLocalSource verifies that V2ReferenceInfo correctly
// deserialises a reference entry whose source field is a local source, and that
// AsSourceUnion() returns the concrete *ReferenceLocalSource variant.
func TestV2ReferenceInfoUnmarshalLocalSource(t *testing.T) {
	t.Parallel()

	raw := `{
		"name": "local-ref",
		"path": "refs/local/mypath",
		"source": {
			"type": "local",
			"path": "/home/user/project",
			"description": "local project",
			"hidden": true
		}
	}`

	var info opencode.V2ReferenceInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	if info.Name != "local-ref" {
		t.Errorf("Name: got %q, want %q", info.Name, "local-ref")
	}

	// Verify Source bearer is populated
	if info.JSON.Source.IsMissing() {
		t.Fatal("Source field is missing after UnmarshalJSON")
	}

	// Verify AsSourceUnion returns the concrete local type
	union := info.AsSourceUnion()
	if union == nil {
		t.Fatal("AsSourceUnion() returned nil")
	}
	localSrc, ok := union.(opencode.ReferenceLocalSource)
	if !ok {
		t.Fatalf("AsSourceUnion() runtime type: got %T, want ReferenceLocalSource", union)
	}
	if localSrc.Type != "local" {
		t.Errorf("ReferenceLocalSource.Type: got %q, want %q", localSrc.Type, "local")
	}
	if localSrc.Path != "/home/user/project" {
		t.Errorf("ReferenceLocalSource.Path: got %q, want %q", localSrc.Path, "/home/user/project")
	}
}

// TestV2ReferenceListResponseUnmarshal verifies that V2ReferenceListResponse can
// deserialise a full API response containing mixed source types.
func TestV2ReferenceListResponseUnmarshal(t *testing.T) {
	t.Parallel()

	raw := `{
		"location": {
			"directory": "/workspace/proj",
			"workspaceID": "wrk_123",
			"project": {
				"id": "proj_abc",
				"directory": "/workspace/proj"
			}
		},
		"data": [
			{
				"name": "git-ref",
				"path": "refs/git/something",
				"source": {
					"type": "git",
					"repository": "https://github.com/org/repo.git"
				}
			},
			{
				"name": "local-ref",
				"path": "refs/local/something",
				"source": {
					"type": "local",
					"path": "/local/path"
				}
			}
		]
	}`

	var resp opencode.V2ReferenceListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("UnmarshalJSON V2ReferenceListResponse failed: %v", err)
	}

	if resp.Location.Directory != "/workspace/proj" {
		t.Errorf("Location.Directory: got %q, want %q", resp.Location.Directory, "/workspace/proj")
	}
	if len(resp.Data) != 2 {
		t.Fatalf("Data length: got %d, want 2", len(resp.Data))
	}

	// First item: git
	gitInfo := resp.Data[0]
	gitUnion := gitInfo.AsSourceUnion()
	if _, ok := gitUnion.(opencode.ReferenceGitSource); !ok {
		t.Errorf("Data[0].AsSourceUnion() type: got %T, want ReferenceGitSource", gitUnion)
	}

	// Second item: local
	localInfo := resp.Data[1]
	localUnion := localInfo.AsSourceUnion()
	if _, ok := localUnion.(opencode.ReferenceLocalSource); !ok {
		t.Errorf("Data[1].AsSourceUnion() type: got %T, want ReferenceLocalSource", localUnion)
	}
}

// TestV2ReferenceListParamsURLQuery verifies that V2ReferenceListParams produces
// correct deepObject-encoded query parameters.
func TestV2ReferenceListParamsURLQuery(t *testing.T) {
	t.Parallel()

	params := opencode.V2ReferenceListParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/my/dir"),
			Workspace: opencode.F("wrk_xyz"),
		}),
	}

	got := params.URLQuery()

	cases := map[string]string{
		"location[directory]": "/my/dir",
		"location[workspace]": "wrk_xyz",
	}
	for k, want := range cases {
		if got.Get(k) != want {
			t.Errorf("URLQuery()[%q]: got %q, want %q", k, got.Get(k), want)
		}
	}
}
