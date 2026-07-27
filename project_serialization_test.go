package opencode

import (
	"encoding/json"
	"strings"
	"testing"
)

// Aligned with OpenAPI GET /project + JS SDK(v2) Project.list.
func TestProjectListParamsQuery(t *testing.T) {
	p := ProjectListParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI GET /project/current + JS SDK(v2) Project.current.
func TestProjectCurrentParamsQuery(t *testing.T) {
	p := ProjectCurrentParams{Directory: F("d")}
	got := p.URLQuery().Encode()
	want := "directory=d"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI POST /project/git/init (no body) + Project.initGit.
func TestProjectInitGitParamsQuery(t *testing.T) {
	p := ProjectInitGitParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI GET /project/{projectID}/directories + Project.directories.
func TestProjectDirectoriesParamsQuery(t *testing.T) {
	p := ProjectDirectoriesParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI PATCH /project/{projectID} + JS SDK(v2) Project.update.
// path: projectID; query: directory, workspace; body: name, icon, commands
func TestProjectUpdateParamsBodyAndQuery(t *testing.T) {
	t.Run("body excludes query fields", func(t *testing.T) {
		p := ProjectUpdateParams{
			Directory: F("d"),
			Workspace: F("w"),
			Name:      F("My Project"),
			Icon: F(ProjectUpdateParamsIcon{
				URL:      F("https://example.com/icon.png"),
				Override: F("custom"),
				Color:    F("#ff0000"),
			}),
			Commands: F(ProjectUpdateParamsCommands{Start: F("npm start")}),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"commands":{"start":"npm start"},"icon":{"color":"#ff0000","override":"custom","url":"https://example.com/icon.png"},"name":"My Project"}`
		if got != want {
			t.Errorf("body got %s, want %s", got, want)
		}
		if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
			t.Errorf("query fields leaked into body: %s", got)
		}
	})

	t.Run("query serialization", func(t *testing.T) {
		p := ProjectUpdateParams{Directory: F("d"), Workspace: F("w"), Name: F("x")}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("query got %q, want %q", got, want)
		}
	})

	t.Run("name-only body", func(t *testing.T) {
		p := ProjectUpdateParams{Name: F("only")}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		want := `{"name":"only"}`
		if string(b) != want {
			t.Errorf("got %s, want %s", string(b), want)
		}
	})
}

// Aligned with OpenAPI Project schema (returned by list/current/update/initGit).
// required: id, worktree, time, sandboxes; optional: vcs, name, icon, commands
func TestProjectUnmarshal(t *testing.T) {
	t.Run("full project", func(t *testing.T) {
		raw := `{
			"id":"proj_123",
			"worktree":"/repo",
			"time":{"created":1700000000,"updated":1700000100,"initialized":1700000050},
			"sandboxes":["sb1","sb2"],
			"vcs":"git",
			"name":"My Project",
			"icon":{"url":"https://x/i.png","override":"o","color":"#fff"},
			"commands":{"start":"npm start"}
		}`
		var p Project
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatal(err)
		}
		if p.ID != "proj_123" || p.Worktree != "/repo" {
			t.Errorf("project = %+v", p)
		}
		if p.Time.Created != 1700000000 || p.Time.Updated != 1700000100 || p.Time.Initialized != 1700000050 {
			t.Errorf("time = %+v", p.Time)
		}
		if len(p.Sandboxes) != 2 || p.Sandboxes[1] != "sb2" {
			t.Errorf("sandboxes = %+v", p.Sandboxes)
		}
		if p.Vcs != ProjectVcsGit {
			t.Errorf("vcs = %q", p.Vcs)
		}
		if p.Name != "My Project" {
			t.Errorf("name = %q", p.Name)
		}
		if p.Icon.URL != "https://x/i.png" || p.Icon.Override != "o" || p.Icon.Color != "#fff" {
			t.Errorf("icon = %+v", p.Icon)
		}
		if p.Commands.Start != "npm start" {
			t.Errorf("commands = %+v", p.Commands)
		}
	})

	t.Run("required-only (optional fields absent)", func(t *testing.T) {
		raw := `{"id":"p","worktree":"/w","time":{"created":1,"updated":2},"sandboxes":[]}`
		var p Project
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatal(err)
		}
		if p.ID != "p" || p.Worktree != "/w" {
			t.Errorf("project = %+v", p)
		}
		// initialized absent -> zero value int64
		if p.Time.Initialized != 0 {
			t.Errorf("initialized should be 0, got %d", p.Time.Initialized)
		}
		// vcs absent -> empty string
		if p.Vcs != "" {
			t.Errorf("vcs should be empty, got %q", p.Vcs)
		}
		if len(p.Sandboxes) != 0 {
			t.Errorf("sandboxes should be empty, got %+v", p.Sandboxes)
		}
	})
}

// Aligned with OpenAPI GET /project/{projectID}/directories 200 items.
// required: directory; optional: strategy
func TestProjectDirectoryEntryUnmarshal(t *testing.T) {
	raw := `[
		{"directory":"/repo/a","strategy":"explicit"},
		{"directory":"/repo/b"}
	]`
	var arr []ProjectDirectoryEntry
	if err := json.Unmarshal([]byte(raw), &arr); err != nil {
		t.Fatal(err)
	}
	if len(arr) != 2 {
		t.Fatalf("arr len = %d", len(arr))
	}
	if arr[0].Directory != "/repo/a" || arr[0].Strategy != "explicit" {
		t.Errorf("arr[0] = %+v", arr[0])
	}
	// strategy absent -> empty string
	if arr[1].Directory != "/repo/b" || arr[1].Strategy != "" {
		t.Errorf("arr[1] = %+v", arr[1])
	}
}

func TestProjectVcsIsKnown(t *testing.T) {
	if !ProjectVcsGit.IsKnown() {
		t.Error("git should be known")
	}
	if ProjectVcs("svn").IsKnown() {
		t.Error("svn should not be known")
	}
}
