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

func TestWorktreeList(t *testing.T) {
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
	_, err := client.Worktree.List(context.TODO(), opencode.WorktreeListParams{
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

func TestWorktreeCreate(t *testing.T) {
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
	_, err := client.Worktree.New(context.TODO(), opencode.WorktreeNewParams{
		Name:         opencode.F("name"),
		StartCommand: opencode.F("cmd"),
		Directory:    opencode.F("directory"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorktreeRemove(t *testing.T) {
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
	_, err := client.Worktree.Remove(context.TODO(), opencode.WorktreeRemoveParams{
		Directory: opencode.F("routing-directory"),
		Workspace: opencode.F("workspace"),
		Body: opencode.F(opencode.WorktreeRemoveParamsBody{
			Directory: opencode.F("/path/to/worktree"),
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

func TestWorktreeReset(t *testing.T) {
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
	_, err := client.Worktree.Reset(context.TODO(), opencode.WorktreeResetParams{
		Directory: opencode.F("routing-directory"),
		Workspace: opencode.F("workspace"),
		Body: opencode.F(opencode.WorktreeResetParamsBody{
			Directory: opencode.F("/path/to/worktree"),
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

// TestWorktreeRemoveParamsSeparation verifies that WorktreeRemoveParams correctly
// separates the routing query directory from the worktree body directory.
// The body must only contain the worktree path; the query must only contain
// the routing context — they must not pollute each other.
func TestWorktreeRemoveParamsSeparation(t *testing.T) {
	t.Parallel()

	params := opencode.WorktreeRemoveParams{
		Directory: opencode.F("routing-dir"),
		Workspace: opencode.F("ws-1"),
		Body: opencode.F(opencode.WorktreeRemoveParamsBody{
			Directory: opencode.F("/actual/worktree/path"),
		}),
	}

	// 1. Body serialisation must only contain the worktree directory
	bodyBytes, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal body: %v", err)
	}
	var bodyMap map[string]any
	if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
		t.Fatalf("re-parse body JSON: %v", err)
	}
	if got, ok := bodyMap["directory"]; !ok || got != "/actual/worktree/path" {
		t.Errorf("body: expected directory=/actual/worktree/path, got %v", bodyMap)
	}
	if _, ok := bodyMap["workspace"]; ok {
		t.Errorf("body must not contain workspace, got %v", bodyMap)
	}

	// 2. Query serialisation must only contain routing context fields
	query := params.URLQuery()
	if got := query.Get("directory"); got != "routing-dir" {
		t.Errorf("query: expected directory=routing-dir, got %q", got)
	}
	if got := query.Get("workspace"); got != "ws-1" {
		t.Errorf("query: expected workspace=ws-1, got %q", got)
	}

	// 3. Query must not contain worktree-path value in directory
	if query.Get("directory") == "/actual/worktree/path" {
		t.Errorf("query directory must be routing context, not worktree path")
	}
}

// TestWorktreeResetParamsSeparation verifies that WorktreeResetParams correctly
// separates the routing query directory from the worktree body directory.
func TestWorktreeResetParamsSeparation(t *testing.T) {
	t.Parallel()

	params := opencode.WorktreeResetParams{
		Directory: opencode.F("routing-dir"),
		Workspace: opencode.F("ws-2"),
		Body: opencode.F(opencode.WorktreeResetParamsBody{
			Directory: opencode.F("/actual/worktree/reset/path"),
		}),
	}

	// 1. Body serialisation must only contain the worktree directory
	bodyBytes, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal body: %v", err)
	}
	var bodyMap map[string]any
	if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
		t.Fatalf("re-parse body JSON: %v", err)
	}
	if got, ok := bodyMap["directory"]; !ok || got != "/actual/worktree/reset/path" {
		t.Errorf("body: expected directory=/actual/worktree/reset/path, got %v", bodyMap)
	}
	if _, ok := bodyMap["workspace"]; ok {
		t.Errorf("body must not contain workspace, got %v", bodyMap)
	}

	// 2. Query serialisation must only contain routing context fields
	query := params.URLQuery()
	if got := query.Get("directory"); got != "routing-dir" {
		t.Errorf("query: expected directory=routing-dir, got %q", got)
	}
	if got := query.Get("workspace"); got != "ws-2" {
		t.Errorf("query: expected workspace=ws-2, got %q", got)
	}

	// 3. Query must not contain worktree-path value in directory
	if query.Get("directory") == "/actual/worktree/reset/path" {
		t.Errorf("query directory must be routing context, not worktree path")
	}
}

// TestWorktreeRemoveParamsBodyOnlyInBody verifies that when query directory is
// absent but body directory is set, the body still serialises the worktree path.
func TestWorktreeRemoveParamsBodyOnlyInBody(t *testing.T) {
	t.Parallel()

	params := opencode.WorktreeRemoveParams{
		// No routing directory set
		Body: opencode.F(opencode.WorktreeRemoveParamsBody{
			Directory: opencode.F("/standalone/worktree"),
		}),
	}

	bodyBytes, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	var bodyMap map[string]any
	if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
		t.Fatalf("re-parse: %v", err)
	}
	if got, ok := bodyMap["directory"]; !ok || got != "/standalone/worktree" {
		t.Errorf("body: expected directory=/standalone/worktree, got %v", bodyMap)
	}

	query := params.URLQuery()
	if got := query.Get("directory"); got != "" {
		t.Errorf("query directory should be empty when not set, got %q", got)
	}
}

// TestWorktreeRemoveParamsBodyAbsentNoBody verifies that when Body is not set at
// all (zero value, Present == false), MarshalJSON returns nil (no bytes), matching
// JS SDK(v2) behavior where an unset `worktreeRemoveInput` sends no request body
// (sdk.gen.ts:1588-1594: `worktreeRemoveInput?: WorktreeRemoveInput`). Before the
// fix, WorktreeRemoveParams.Body was a bare struct (not param.Field[T] wrapped)
// and MarshalJSON unconditionally emitted `{}`.
func TestWorktreeRemoveParamsBodyAbsentNoBody(t *testing.T) {
	t.Parallel()

	params := opencode.WorktreeRemoveParams{}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	if data != nil {
		t.Errorf("MarshalJSON with unset Body: expected nil, got %q (len=%d)", data, len(data))
	}
	if len(data) != 0 {
		t.Errorf("MarshalJSON with unset Body: expected 0 bytes, got %d", len(data))
	}
}

// TestWorktreeRemoveParamsBodyPresentEmitsBody verifies that once Body is set via
// opencode.F(...), MarshalJSON emits exactly {"directory":"/path"}.
func TestWorktreeRemoveParamsBodyPresentEmitsBody(t *testing.T) {
	t.Parallel()

	params := opencode.WorktreeRemoveParams{
		Body: opencode.F(opencode.WorktreeRemoveParamsBody{
			Directory: opencode.F("/path"),
		}),
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	if string(data) != `{"directory":"/path"}` {
		t.Errorf("MarshalJSON with Body set: got %s, want %s", data, `{"directory":"/path"}`)
	}
}

// TestWorktreeResetParamsBodyAbsentNoBody mirrors
// TestWorktreeRemoveParamsBodyAbsentNoBody for WorktreeResetParams.
func TestWorktreeResetParamsBodyAbsentNoBody(t *testing.T) {
	t.Parallel()

	params := opencode.WorktreeResetParams{}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	if data != nil {
		t.Errorf("MarshalJSON with unset Body: expected nil, got %q (len=%d)", data, len(data))
	}
	if len(data) != 0 {
		t.Errorf("MarshalJSON with unset Body: expected 0 bytes, got %d", len(data))
	}
}

// TestWorktreeResetParamsBodyPresentEmitsBody mirrors
// TestWorktreeRemoveParamsBodyPresentEmitsBody for WorktreeResetParams.
func TestWorktreeResetParamsBodyPresentEmitsBody(t *testing.T) {
	t.Parallel()

	params := opencode.WorktreeResetParams{
		Body: opencode.F(opencode.WorktreeResetParamsBody{
			Directory: opencode.F("/path"),
		}),
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	if string(data) != `{"directory":"/path"}` {
		t.Errorf("MarshalJSON with Body set: got %s, want %s", data, `{"directory":"/path"}`)
	}
}

// TestWorktreeRemoveParamsURLQueryEncoding verifies the exact encoded query string
// for WorktreeRemoveParams.URLQuery(), covering the query:"directory"/"workspace"
// tags independently of the body.
func TestWorktreeRemoveParamsURLQueryEncoding(t *testing.T) {
	t.Parallel()

	params := opencode.WorktreeRemoveParams{
		Directory: opencode.F("/routing/dir"),
		Workspace: opencode.F("ws-1"),
	}
	got := params.URLQuery().Encode()
	want := "directory=%2Frouting%2Fdir&workspace=ws-1"
	if got != want {
		t.Errorf("URLQuery().Encode(): got %q, want %q", got, want)
	}
}

// TestWorktreeResetParamsURLQueryEncoding mirrors
// TestWorktreeRemoveParamsURLQueryEncoding for WorktreeResetParams.
func TestWorktreeResetParamsURLQueryEncoding(t *testing.T) {
	t.Parallel()

	params := opencode.WorktreeResetParams{
		Directory: opencode.F("/routing/dir"),
		Workspace: opencode.F("ws-2"),
	}
	got := params.URLQuery().Encode()
	want := "directory=%2Frouting%2Fdir&workspace=ws-2"
	if got != want {
		t.Errorf("URLQuery().Encode(): got %q, want %q", got, want)
	}
}
