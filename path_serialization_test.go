package opencode

import (
	"encoding/json"
	"testing"
)

// Aligned with OpenAPI GET /path + JS SDK(v2) Path.get.
// query all optional: directory, workspace
func TestPathGetParamsQuery(t *testing.T) {
	t.Run("both fields", func(t *testing.T) {
		p := PathGetParams{Directory: F("d"), Workspace: F("w")}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("empty", func(t *testing.T) {
		p := PathGetParams{}
		if got := p.URLQuery().Encode(); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
}

// Aligned with OpenAPI GET /path 200 (Path).
// required: home, config, directory, state, worktree
func TestPathUnmarshal(t *testing.T) {
	raw := `{
		"home":"/home/user",
		"config":"/home/user/.config/opencode",
		"directory":"/repo",
		"state":"/home/user/.local/state/opencode",
		"worktree":"/repo"
	}`
	var p Path
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		t.Fatal(err)
	}
	if p.Home != "/home/user" {
		t.Errorf("home = %q", p.Home)
	}
	if p.Config != "/home/user/.config/opencode" {
		t.Errorf("config = %q", p.Config)
	}
	if p.Directory != "/repo" {
		t.Errorf("directory = %q", p.Directory)
	}
	if p.State != "/home/user/.local/state/opencode" {
		t.Errorf("state = %q", p.State)
	}
	if p.Worktree != "/repo" {
		t.Errorf("worktree = %q", p.Worktree)
	}
	if p.JSON.raw == "" {
		t.Error("RawJSON not preserved")
	}
}
