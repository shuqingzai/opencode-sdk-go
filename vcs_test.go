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

func TestVcsGet(t *testing.T) {
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
	_, err := client.Vcs.Get(context.TODO(), opencode.VcsGetParams{
		Directory: opencode.F("directory"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// --- Unit tests (no server required) ---

// TestVcsDiffSubServiceWired verifies that the `vcs.diff.raw` operationId is
// reachable only through the nested [opencode.VcsDiffService] (mirrors the JS SDK
// `Vcs.diff2.raw()` sub-client), and that no flat shim is required.
func TestVcsDiffSubServiceWired(t *testing.T) {
	s := opencode.NewVcsService()
	if s.Diff == nil {
		t.Fatal("VcsService.Diff sub-service should be initialized by NewVcsService")
	}
	client := opencode.NewClient(option.WithBaseURL("http://localhost:4010"))
	if client.Vcs.Diff == nil {
		t.Fatal("Client.Vcs.Diff sub-service should be initialized by NewClient")
	}
}

// TestVcsFileStatusStatusIsKnown covers the OpenAPI `VcsFileStatus.status` enum
// ["added","deleted","modified"].
func TestVcsFileStatusStatusIsKnown(t *testing.T) {
	known := []opencode.VcsFileStatusStatus{
		opencode.VcsFileStatusStatusAdded,
		opencode.VcsFileStatusStatusDeleted,
		opencode.VcsFileStatusStatusModified,
	}
	for _, v := range known {
		if !v.IsKnown() {
			t.Errorf("VcsFileStatusStatus(%q).IsKnown() = false, want true", v)
		}
	}
	for _, v := range []opencode.VcsFileStatusStatus{"", "renamed", "Added", "untracked"} {
		if v.IsKnown() {
			t.Errorf("VcsFileStatusStatus(%q).IsKnown() = true, want false", v)
		}
	}
	// Enum literal values must match the OpenAPI spec exactly.
	if got := string(opencode.VcsFileStatusStatusAdded); got != "added" {
		t.Errorf("VcsFileStatusStatusAdded = %q, want %q", got, "added")
	}
	if got := string(opencode.VcsFileStatusStatusDeleted); got != "deleted" {
		t.Errorf("VcsFileStatusStatusDeleted = %q, want %q", got, "deleted")
	}
	if got := string(opencode.VcsFileStatusStatusModified); got != "modified" {
		t.Errorf("VcsFileStatusStatusModified = %q, want %q", got, "modified")
	}
}

// TestVcsFileDiffStatusIsKnown covers the OpenAPI `VcsFileDiff.status` enum, which
// is a schema-independent inline enum from `VcsFileStatus.status`.
func TestVcsFileDiffStatusIsKnown(t *testing.T) {
	known := []opencode.VcsFileDiffStatus{
		opencode.VcsFileDiffStatusAdded,
		opencode.VcsFileDiffStatusDeleted,
		opencode.VcsFileDiffStatusModified,
	}
	for _, v := range known {
		if !v.IsKnown() {
			t.Errorf("VcsFileDiffStatus(%q).IsKnown() = false, want true", v)
		}
	}
	for _, v := range []opencode.VcsFileDiffStatus{"", "renamed", "MODIFIED"} {
		if v.IsKnown() {
			t.Errorf("VcsFileDiffStatus(%q).IsKnown() = true, want false", v)
		}
	}
}

// TestVcsDiffParamsModeIsKnown covers the OpenAPI `vcs.diff` `mode` query enum
// ["git","branch"].
func TestVcsDiffParamsModeIsKnown(t *testing.T) {
	for _, v := range []opencode.VcsDiffParamsMode{
		opencode.VcsDiffParamsModeGit,
		opencode.VcsDiffParamsModeBranch,
	} {
		if !v.IsKnown() {
			t.Errorf("VcsDiffParamsMode(%q).IsKnown() = false, want true", v)
		}
	}
	for _, v := range []opencode.VcsDiffParamsMode{"", "worktree", "Git"} {
		if v.IsKnown() {
			t.Errorf("VcsDiffParamsMode(%q).IsKnown() = true, want false", v)
		}
	}
}

// TestVcsFileStatusUnmarshal verifies `VcsFileStatus` deserialization. Per OpenAPI
// all four of file/additions/deletions/status are required.
func TestVcsFileStatusUnmarshal(t *testing.T) {
	raw := `{"file":"src/main.go","additions":12,"deletions":3,"status":"modified"}`
	var fs opencode.VcsFileStatus
	if err := fs.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if fs.File != "src/main.go" {
		t.Errorf("File = %q, want %q", fs.File, "src/main.go")
	}
	if fs.Additions != 12 {
		t.Errorf("Additions = %d, want 12", fs.Additions)
	}
	if fs.Deletions != 3 {
		t.Errorf("Deletions = %d, want 3", fs.Deletions)
	}
	if fs.Status != opencode.VcsFileStatusStatusModified {
		t.Errorf("Status = %q, want %q", fs.Status, opencode.VcsFileStatusStatusModified)
	}
	if !fs.Status.IsKnown() {
		t.Error("Status.IsKnown() = false, want true")
	}
	if fs.JSON.RawJSON() == "" {
		t.Error("expected non-empty RawJSON")
	}
}

// TestVcsFileStatusUnmarshalUnknownStatus verifies forward compatibility: an
// unrecognized enum value must still round-trip through the raw JSON.
func TestVcsFileStatusUnmarshalUnknownStatus(t *testing.T) {
	raw := `{"file":"a.txt","additions":0,"deletions":0,"status":"renamed","future":1}`
	var fs opencode.VcsFileStatus
	if err := fs.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if fs.Status != "renamed" {
		t.Errorf("Status = %q, want %q", fs.Status, "renamed")
	}
	if fs.Status.IsKnown() {
		t.Error("Status.IsKnown() = true for \"renamed\", want false")
	}
	if fs.JSON.RawJSON() != raw {
		t.Errorf("RawJSON should preserve the upstream payload verbatim, got %q", fs.JSON.RawJSON())
	}
}

// TestVcsFileDiffUnmarshal verifies `VcsFileDiff` deserialization. Per OpenAPI
// `patch` and `status` are optional while file/additions/deletions are required.
func TestVcsFileDiffUnmarshal(t *testing.T) {
	raw := `{"file":"a.go","patch":"@@ -1 +1 @@\n-a\n+b\n","additions":1,"deletions":1,"status":"added"}`
	var fd opencode.VcsFileDiff
	if err := fd.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if fd.File != "a.go" {
		t.Errorf("File = %q, want %q", fd.File, "a.go")
	}
	if fd.Patch == "" {
		t.Error("Patch should be populated")
	}
	if fd.Additions != 1 || fd.Deletions != 1 {
		t.Errorf("Additions/Deletions = %d/%d, want 1/1", fd.Additions, fd.Deletions)
	}
	if fd.Status != opencode.VcsFileDiffStatusAdded {
		t.Errorf("Status = %q, want %q", fd.Status, opencode.VcsFileDiffStatusAdded)
	}

	// `patch` and `status` omitted - both are optional in the OpenAPI schema.
	var bare opencode.VcsFileDiff
	if err := bare.UnmarshalJSON([]byte(`{"file":"b.go","additions":0,"deletions":0}`)); err != nil {
		t.Fatalf("UnmarshalJSON (bare) failed: %v", err)
	}
	if bare.Patch != "" {
		t.Errorf("Patch = %q, want empty", bare.Patch)
	}
	if bare.Status != "" {
		t.Errorf("Status = %q, want empty", bare.Status)
	}
}

// TestVcsInfoUnmarshal verifies the `VcsInfo` schema (branch, default_branch).
func TestVcsInfoUnmarshal(t *testing.T) {
	raw := `{"branch":"feature/x","default_branch":"main"}`
	var info opencode.VcsInfo
	if err := info.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if info.Branch != "feature/x" {
		t.Errorf("Branch = %q, want %q", info.Branch, "feature/x")
	}
	if info.DefaultBranch != "main" {
		t.Errorf("DefaultBranch = %q, want %q", info.DefaultBranch, "main")
	}
	if info.JSON.RawJSON() == "" {
		t.Error("expected non-empty RawJSON")
	}
}

// TestVcsApplyResponseUnmarshal verifies the `vcs.apply` 200 body {applied:boolean}.
func TestVcsApplyResponseUnmarshal(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		want bool
	}{
		{`{"applied":true}`, true},
		{`{"applied":false}`, false},
	} {
		var resp opencode.VcsApplyResponse
		if err := resp.UnmarshalJSON([]byte(tc.raw)); err != nil {
			t.Fatalf("UnmarshalJSON(%s) failed: %v", tc.raw, err)
		}
		if resp.Applied != tc.want {
			t.Errorf("Applied = %v, want %v", resp.Applied, tc.want)
		}
		if resp.JSON.RawJSON() != tc.raw {
			t.Errorf("RawJSON = %q, want %q", resp.JSON.RawJSON(), tc.raw)
		}
	}
}

// TestVcsApplyParamsSerialization verifies the OpenAPI parameter split for
// `vcs.apply`: directory/workspace are query params, patch is the request body.
func TestVcsApplyParamsSerialization(t *testing.T) {
	params := opencode.VcsApplyParams{
		Directory: opencode.F("/repo"),
		Workspace: opencode.F("ws-1"),
		Patch:     opencode.F("@@ -1 +1 @@\n-a\n+b\n"),
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var body map[string]any
	if err := json.Unmarshal(data, &body); err != nil {
		t.Fatalf("body JSON invalid: %v", err)
	}
	if body["patch"] != "@@ -1 +1 @@\n-a\n+b\n" {
		t.Errorf("body patch = %v, want the raw patch", body["patch"])
	}
	if _, ok := body["directory"]; ok {
		t.Error("directory is a query param and must not appear in the body")
	}
	if _, ok := body["workspace"]; ok {
		t.Error("workspace is a query param and must not appear in the body")
	}

	q := params.URLQuery()
	if got := q.Get("directory"); got != "/repo" {
		t.Errorf("query directory = %q, want %q", got, "/repo")
	}
	if got := q.Get("workspace"); got != "ws-1" {
		t.Errorf("query workspace = %q, want %q", got, "ws-1")
	}
	if _, ok := q["patch"]; ok {
		t.Error("patch is a body param and must not appear in the query string")
	}
}

// TestVcsDiffParamsURLQuery verifies all four `vcs.diff` query parameters.
func TestVcsDiffParamsURLQuery(t *testing.T) {
	q := opencode.VcsDiffParams{
		Directory: opencode.F("/repo"),
		Workspace: opencode.F("ws-1"),
		Mode:      opencode.F(opencode.VcsDiffParamsModeBranch),
		Context:   opencode.F(int64(5)),
	}.URLQuery()

	for k, want := range map[string]string{
		"directory": "/repo",
		"workspace": "ws-1",
		"mode":      "branch",
		"context":   "5",
	} {
		if got := q.Get(k); got != want {
			t.Errorf("query %q = %q, want %q", k, got, want)
		}
	}

	// Only `mode` is required; the optional params must be omitted when unset.
	sparse := opencode.VcsDiffParams{Mode: opencode.F(opencode.VcsDiffParamsModeGit)}.URLQuery()
	if got := sparse.Get("mode"); got != "git" {
		t.Errorf("sparse mode = %q, want %q", got, "git")
	}
	for _, k := range []string{"directory", "workspace", "context"} {
		if _, ok := sparse[k]; ok {
			t.Errorf("unset optional param %q must be omitted from the query string", k)
		}
	}
}

// TestVcsQueryOnlyParamsURLQuery covers the three params structs whose OpenAPI
// endpoints expose exactly directory+workspace: vcs.get, vcs.status, vcs.diff.raw.
func TestVcsQueryOnlyParamsURLQuery(t *testing.T) {
	queries := map[string]func() (dir string, ws string, omitted bool){
		"VcsGetParams": func() (string, string, bool) {
			q := opencode.VcsGetParams{Directory: opencode.F("/d"), Workspace: opencode.F("w")}.URLQuery()
			empty := opencode.VcsGetParams{}.URLQuery()
			return q.Get("directory"), q.Get("workspace"), len(empty) == 0
		},
		"VcsStatusParams": func() (string, string, bool) {
			q := opencode.VcsStatusParams{Directory: opencode.F("/d"), Workspace: opencode.F("w")}.URLQuery()
			empty := opencode.VcsStatusParams{}.URLQuery()
			return q.Get("directory"), q.Get("workspace"), len(empty) == 0
		},
		"VcsDiffRawParams": func() (string, string, bool) {
			q := opencode.VcsDiffRawParams{Directory: opencode.F("/d"), Workspace: opencode.F("w")}.URLQuery()
			empty := opencode.VcsDiffRawParams{}.URLQuery()
			return q.Get("directory"), q.Get("workspace"), len(empty) == 0
		},
	}
	for name, fn := range queries {
		dir, ws, omitted := fn()
		if dir != "/d" {
			t.Errorf("%s: directory = %q, want %q", name, dir, "/d")
		}
		if ws != "w" {
			t.Errorf("%s: workspace = %q, want %q", name, ws, "w")
		}
		if !omitted {
			t.Errorf("%s: zero-value params must produce an empty query string", name)
		}
	}
}
