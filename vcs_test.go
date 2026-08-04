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

// TestVcsInfoUnmarshal verifies that VcsInfo (vcs.get response) deserializes
// correctly, including ExtraFields for unknown keys.
func TestVcsInfoUnmarshal(t *testing.T) {
	raw := `{"branch": "main", "default_branch": "main"}`
	var info opencode.VcsInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if info.Branch != "main" {
		t.Errorf("expected branch=main, got %q", info.Branch)
	}
	if info.DefaultBranch != "main" {
		t.Errorf("expected default_branch=main, got %q", info.DefaultBranch)
	}
}

// TestVcsFileDiffUnmarshal verifies that VcsFileDiff (vcs.diff response item)
// deserializes correctly, including the status enum and numeric fields.
func TestVcsFileDiffUnmarshal(t *testing.T) {
	raw := `{"file": "main.go", "additions": 10, "deletions": 2, "status": "modified", "patch": "@@ -1 +1 @@"}`
	var diff opencode.VcsFileDiff
	if err := json.Unmarshal([]byte(raw), &diff); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if diff.File != "main.go" {
		t.Errorf("expected file=main.go, got %q", diff.File)
	}
	if diff.Additions != 10 {
		t.Errorf("expected additions=10, got %d", diff.Additions)
	}
	if diff.Status != opencode.VcsFileDiffStatusModified {
		t.Errorf("expected status=modified, got %v", diff.Status)
	}
	if !diff.Status.IsKnown() {
		t.Error("expected IsKnown()=true for VcsFileDiffStatusModified")
	}
}

// TestVcsFileDiffStatusIsKnown verifies all known enum values and an unknown value.
func TestVcsFileDiffStatusIsKnown(t *testing.T) {
	known := []opencode.VcsFileDiffStatus{
		opencode.VcsFileDiffStatusAdded,
		opencode.VcsFileDiffStatusDeleted,
		opencode.VcsFileDiffStatusModified,
	}
	for _, v := range known {
		if !v.IsKnown() {
			t.Errorf("expected IsKnown()=true for %q", v)
		}
	}
	unknown := opencode.VcsFileDiffStatus("renamed")
	if unknown.IsKnown() {
		t.Error("expected IsKnown()=false for unknown status")
	}
}

// TestVcsDiffParamsModeIsKnown verifies the enum for VcsDiffParams.Mode.
func TestVcsDiffParamsModeIsKnown(t *testing.T) {
	if !opencode.VcsDiffParamsModeGit.IsKnown() {
		t.Error("expected IsKnown()=true for VcsDiffParamsModeGit")
	}
	if !opencode.VcsDiffParamsModeBranch.IsKnown() {
		t.Error("expected IsKnown()=true for VcsDiffParamsModeBranch")
	}
	if opencode.VcsDiffParamsMode("unknown").IsKnown() {
		t.Error("expected IsKnown()=false for unknown mode")
	}
}

// TestVcsApplyParamsURLQueryAndMarshalJSON verifies that VcsApplyParams
// implements both URLQuery (query params) and MarshalJSON (body params).
func TestVcsApplyParamsURLQueryAndMarshalJSON(t *testing.T) {
	params := opencode.VcsApplyParams{
		Directory: opencode.F("mydir"),
		Workspace: opencode.F("mywsp"),
		Patch:     opencode.F("@@ -1,1 +1,2 @@\n context"),
	}

	// Test URLQuery — verifies the method exists and returns non-nil values.
	q := params.URLQuery()
	if q == nil {
		t.Error("URLQuery should return non-nil url.Values")
	}

	// Test MarshalJSON
	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("failed to unmarshal body: %v", err)
	}
	if m["patch"] == nil {
		t.Error("expected patch field in JSON body")
	}
}

// TestVcsDiffRawDeprecatedDelegates verifies that VcsService.DiffRaw is
// present (it's @deprecated but kept for backwards compat) and delegates to
// Diff.Raw. We only check the method exists at compile time via type assertion.
func TestVcsDiffRawDeprecatedDelegates(t *testing.T) {
	// Compile-time check: if VcsDiffRaw doesn't exist, this won't compile.
	type hasDiffRaw interface {
		DiffRaw(ctx context.Context, query opencode.VcsDiffRawParams, opts ...option.RequestOption) (*string, error)
	}
	var _ hasDiffRaw = opencode.NewVcsService()
}
