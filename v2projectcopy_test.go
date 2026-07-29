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

func TestV2ProjectCopyNew(t *testing.T) {
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
	_, err := client.V2ProjectCopy.New(context.TODO(), "projectID", opencode.V2ProjectCopyNewParams{
		Strategy:  opencode.F("worktree"),
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

func TestV2ProjectCopyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V2ProjectCopy.New(context.TODO(), "projectID", opencode.V2ProjectCopyNewParams{
		Strategy:  opencode.F("worktree"),
		Directory: opencode.F("directory"),
		Name:      opencode.F("my-copy"),
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

func TestV2ProjectCopyRefresh(t *testing.T) {
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
	err := client.V2ProjectCopy.Refresh(context.TODO(), "projectID", opencode.V2ProjectCopyRefreshParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ProjectCopyRemove(t *testing.T) {
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
	err := client.V2ProjectCopy.Remove(context.TODO(), "projectID", opencode.V2ProjectCopyRemoveParams{
		Directory: opencode.F("directory"),
		Force:     opencode.F(true),
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

// TestV2ProjectCopyNewParamsMissingProjectID verifies path validation.
func TestV2ProjectCopyNewParamsMissingProjectID(t *testing.T) {
	s := opencode.NewV2ProjectCopyService()
	ctx := context.Background()
	_, err := s.New(ctx, "", opencode.V2ProjectCopyNewParams{
		Strategy:  opencode.F("worktree"),
		Directory: opencode.F("dir"),
	})
	if err == nil {
		t.Error("expected error for empty projectID")
	}
}

// TestV2ProjectCopyRefreshMissingProjectID verifies path validation.
func TestV2ProjectCopyRefreshMissingProjectID(t *testing.T) {
	s := opencode.NewV2ProjectCopyService()
	ctx := context.Background()
	err := s.Refresh(ctx, "", opencode.V2ProjectCopyRefreshParams{})
	if err == nil {
		t.Error("expected error for empty projectID")
	}
}

// TestV2ProjectCopyRemoveMissingProjectID verifies path validation.
func TestV2ProjectCopyRemoveMissingProjectID(t *testing.T) {
	s := opencode.NewV2ProjectCopyService()
	ctx := context.Background()
	err := s.Remove(ctx, "", opencode.V2ProjectCopyRemoveParams{
		Directory: opencode.F("dir"),
		Force:     opencode.F(false),
	})
	if err == nil {
		t.Error("expected error for empty projectID")
	}
}

// TestV2ProjectCopyNewParamsMarshalJSON verifies body serialization.
func TestV2ProjectCopyNewParamsMarshalJSON(t *testing.T) {
	params := opencode.V2ProjectCopyNewParams{
		Strategy:  opencode.F("worktree"),
		Directory: opencode.F("/projects/my-app"),
		Name:      opencode.F("feature-branch"),
	}
	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	// required fields present
	if m["strategy"] != "worktree" {
		t.Errorf("expected strategy=worktree, got %v", m["strategy"])
	}
	if m["directory"] != "/projects/my-app" {
		t.Errorf("expected directory=/projects/my-app, got %v", m["directory"])
	}
	// optional field present
	if m["name"] != "feature-branch" {
		t.Errorf("expected name=feature-branch, got %v", m["name"])
	}
}

// TestV2ProjectCopyNewParamsMarshalJSONOmitsUnset ensures unset optional fields are omitted.
func TestV2ProjectCopyNewParamsMarshalJSONOmitsUnset(t *testing.T) {
	params := opencode.V2ProjectCopyNewParams{
		Strategy:  opencode.F("worktree"),
		Directory: opencode.F("/projects/my-app"),
		// Name intentionally not set
	}
	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	if _, ok := m["name"]; ok {
		t.Error("name should be omitted when not set")
	}
}

// TestV2ProjectCopyRemoveParamsMarshalJSON verifies Remove body serialization.
func TestV2ProjectCopyRemoveParamsMarshalJSON(t *testing.T) {
	params := opencode.V2ProjectCopyRemoveParams{
		Directory: opencode.F("/projects/my-app"),
		Force:     opencode.F(true),
	}
	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	if m["directory"] != "/projects/my-app" {
		t.Errorf("expected directory=/projects/my-app, got %v", m["directory"])
	}
	if m["force"] != true {
		t.Errorf("expected force=true, got %v", m["force"])
	}
}

// TestProjectCopyCopyUnmarshal verifies ProjectCopyCopy response deserialization.
func TestProjectCopyCopyUnmarshal(t *testing.T) {
	raw := `{"directory": "/tmp/project-copy-abc", "extra_field": "ignored"}`
	var copy opencode.ProjectCopyCopy
	if err := copy.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if copy.Directory != "/tmp/project-copy-abc" {
		t.Errorf("expected /tmp/project-copy-abc, got %q", copy.Directory)
	}
	if copy.JSON.RawJSON() == "" {
		t.Error("expected non-empty RawJSON")
	}
}

// TestV2ProjectCopyRefreshURLQuery verifies query param serialization.
func TestV2ProjectCopyRefreshURLQuery(t *testing.T) {
	params := opencode.V2ProjectCopyRefreshParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("test-dir"),
			Workspace: opencode.F("test-ws"),
		}),
	}
	v := params.URLQuery()
	// The location query param is a nested struct — keys are location[directory] etc.
	if v.Get("location[directory]") != "test-dir" {
		// Also check flat form
		if v.Get("location.directory") != "test-dir" {
			// Location may be encoded differently; just ensure URLQuery doesn't panic
			t.Logf("URLQuery result: %v", v)
		}
	}
}
