package opencode

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/sst/opencode-sdk-go/option"
)

// TestWorktreeRemoveParamsSplit verifies that query params and body params
// are correctly separated: MarshalJSON returns only body, URLQuery only query.
func TestWorktreeRemoveParamsSplit(t *testing.T) {
	params := WorktreeRemoveParams{
		Directory: F("query-workspace-root"),
		Workspace: F("query-workspace"),
		Body: F(WorktreeRemoveParamsBody{
			Directory: F("/path/to/worktree"),
		}),
	}

	// Test body serialization: should only contain body.directory
	bodyJSON, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	var body map[string]any
	if err := json.Unmarshal(bodyJSON, &body); err != nil {
		t.Fatalf("body JSON invalid: %v", err)
	}
	if body["directory"] != "/path/to/worktree" {
		t.Errorf("body.directory = %v, want /path/to/worktree", body["directory"])
	}
	// Query-only fields must NOT appear in body
	if _, ok := body["workspace"]; ok {
		t.Error("workspace should not appear in body JSON")
	}

	// Test query serialization: should only contain query Directory and Workspace
	qv := params.URLQuery()
	if got := qv.Get("directory"); got != "query-workspace-root" {
		t.Errorf("query directory = %q, want query-workspace-root", got)
	}
	if got := qv.Get("workspace"); got != "query-workspace" {
		t.Errorf("query workspace = %q, want query-workspace", got)
	}
	// Body field must NOT appear in query
	if qv.Get("Body") != "" || qv.Get("body") != "" {
		t.Error("Body should not appear in URL query")
	}
	_ = url.Values{}
}

func TestWorktreeResetParamsSplit(t *testing.T) {
	params := WorktreeResetParams{
		Directory: F("query-workspace-root"),
		Workspace: F("query-workspace"),
		Body: F(WorktreeResetParamsBody{
			Directory: F("/path/to/reset"),
		}),
	}

	bodyJSON, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	var body map[string]any
	if err := json.Unmarshal(bodyJSON, &body); err != nil {
		t.Fatalf("body JSON invalid: %v", err)
	}
	if body["directory"] != "/path/to/reset" {
		t.Errorf("body.directory = %v, want /path/to/reset", body["directory"])
	}
	if _, ok := body["workspace"]; ok {
		t.Error("workspace should not appear in body JSON")
	}

	qv := params.URLQuery()
	if got := qv.Get("directory"); got != "query-workspace-root" {
		t.Errorf("query directory = %q, want query-workspace-root", got)
	}
	if got := qv.Get("workspace"); got != "query-workspace" {
		t.Errorf("query workspace = %q, want query-workspace", got)
	}
}

// The OpenAPI requestBody for worktree.remove / worktree.reset is optional
// (requestBody.required == false), and the JS SDK exposes it as an optional
// `worktreeRemoveInput?` / `worktreeResetInput?` parameter that JSON.stringify
// drops entirely when absent. Wrapping Body in param.Field makes
// apijson.MarshalRoot emit nothing for an unset Body — no manual presence
// check is needed, and an unset Body must not degrade into an empty `{}`.
func TestWorktreeOptionalBodyOmitted(t *testing.T) {
	t.Run("remove without body", func(t *testing.T) {
		p := WorktreeRemoveParams{Directory: F("/proj")}
		b, err := p.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON: %v", err)
		}
		if len(b) != 0 {
			t.Errorf("body = %q, want empty (no body sent)", b)
		}
		if got := p.URLQuery().Get("directory"); got != "/proj" {
			t.Errorf("query directory = %q, want /proj", got)
		}
	})

	t.Run("reset without body", func(t *testing.T) {
		p := WorktreeResetParams{Workspace: F("ws-1")}
		b, err := p.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON: %v", err)
		}
		if len(b) != 0 {
			t.Errorf("body = %q, want empty (no body sent)", b)
		}
	})

	t.Run("explicitly empty body is still sent", func(t *testing.T) {
		// param.Field lets callers distinguish "unset" from "explicitly empty".
		p := WorktreeRemoveParams{Body: F(WorktreeRemoveParamsBody{})}
		b, err := p.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON: %v", err)
		}
		if string(b) != "{}" {
			t.Errorf("body = %q, want {}", b)
		}
	})
}

// End-to-end: what actually reaches the wire.
func TestWorktreeOptionalBodyOnTheWire(t *testing.T) {
	capture := func(call func(*Client) error) (path, body string) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			raw, _ := io.ReadAll(r.Body)
			path, body = r.URL.String(), string(raw)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`true`))
		}))
		defer srv.Close()
		if err := call(NewClient(option.WithBaseURL(srv.URL))); err != nil {
			t.Fatalf("call: %v", err)
		}
		return
	}

	t.Run("remove without body sends no body", func(t *testing.T) {
		path, body := capture(func(c *Client) error {
			_, err := c.Worktree.Remove(context.Background(), WorktreeRemoveParams{Directory: F("/proj")})
			return err
		})
		if body != "" {
			t.Errorf("body = %q, want empty", body)
		}
		if path != "/experimental/worktree?directory=%2Fproj" {
			t.Errorf("path = %q", path)
		}
	})

	t.Run("remove with body sends only body fields", func(t *testing.T) {
		path, body := capture(func(c *Client) error {
			_, err := c.Worktree.Remove(context.Background(), WorktreeRemoveParams{
				Directory: F("/proj"),
				Workspace: F("ws-1"),
				Body:      F(WorktreeRemoveParamsBody{Directory: F("/proj/wt-1")}),
			})
			return err
		})
		if body != `{"directory":"/proj/wt-1"}` {
			t.Errorf("body = %q", body)
		}
		// query directory must stay the routing value, not the worktree path
		if path != "/experimental/worktree?directory=%2Fproj&workspace=ws-1" {
			t.Errorf("path = %q", path)
		}
	})

	t.Run("reset with body", func(t *testing.T) {
		_, body := capture(func(c *Client) error {
			_, err := c.Worktree.Reset(context.Background(), WorktreeResetParams{
				Body: F(WorktreeResetParamsBody{Directory: F("/proj/wt-2")}),
			})
			return err
		})
		if body != `{"directory":"/proj/wt-2"}` {
			t.Errorf("body = %q", body)
		}
	})
}
