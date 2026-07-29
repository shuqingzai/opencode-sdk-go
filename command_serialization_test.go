package opencode

import (
	"encoding/json"
	"testing"
)

// Aligned with OpenAPI GET /command response (Command schema).

// TestCommandUnmarshal verifies Command deserialization covers required and
// optional fields. OpenAPI: name+template+hints are required; source+agent+
// description+model+subtask are optional.
func TestCommandUnmarshal(t *testing.T) {
	t.Run("required fields only", func(t *testing.T) {
		raw := `{
			"name": "build",
			"template": "npm run build",
			"hints": ["production", "ci"]
		}`
		var c Command
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if c.Name != "build" {
			t.Errorf("Name = %q, want build", c.Name)
		}
		if c.Template != "npm run build" {
			t.Errorf("Template = %q, want 'npm run build'", c.Template)
		}
		if len(c.Hints) != 2 || c.Hints[0] != "production" || c.Hints[1] != "ci" {
			t.Errorf("Hints = %v, want [production ci]", c.Hints)
		}
		// Optional fields should be zero
		if c.Source != "" {
			t.Errorf("Source should be empty when absent, got %q", c.Source)
		}
		if c.Agent != "" || c.Description != "" || c.Model != "" || c.Subtask {
			t.Errorf("optional fields should be zero: %+v", c)
		}
		if c.JSON.RawJSON() == "" {
			t.Error("RawJSON not preserved")
		}
	})

	t.Run("all fields present", func(t *testing.T) {
		raw := `{
			"name": "test",
			"template": "go test ./...",
			"hints": [],
			"source": "command",
			"agent": "build",
			"description": "Run all tests",
			"model": "claude-3-5",
			"subtask": true
		}`
		var c Command
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if c.Source != CommandSourceCommand {
			t.Errorf("Source = %q, want command", c.Source)
		}
		if !c.Source.IsKnown() {
			t.Error("Source.IsKnown() = false")
		}
		if c.Agent != "build" {
			t.Errorf("Agent = %q, want build", c.Agent)
		}
		if c.Description != "Run all tests" {
			t.Errorf("Description = %q, want 'Run all tests'", c.Description)
		}
		if c.Model != "claude-3-5" {
			t.Errorf("Model = %q, want claude-3-5", c.Model)
		}
		if !c.Subtask {
			t.Error("Subtask should be true")
		}
		if len(c.Hints) != 0 {
			t.Errorf("Hints = %v, want []", c.Hints)
		}
	})

	t.Run("mcp source", func(t *testing.T) {
		raw := `{"name": "mcp_cmd", "template": "call {{tool}}", "hints": [], "source": "mcp"}`
		var c Command
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if c.Source != CommandSourceMcp {
			t.Errorf("Source = %q, want mcp", c.Source)
		}
	})

	t.Run("skill source", func(t *testing.T) {
		raw := `{"name": "skill_cmd", "template": "run {{skill}}", "hints": [], "source": "skill"}`
		var c Command
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if c.Source != CommandSourceSkill {
			t.Errorf("Source = %q, want skill", c.Source)
		}
	})
}

// TestCommandSourceIsKnown verifies CommandSource enum values.
func TestCommandSourceIsKnown(t *testing.T) {
	for _, v := range []CommandSource{CommandSourceCommand, CommandSourceMcp, CommandSourceSkill} {
		if !v.IsKnown() {
			t.Errorf("%q.IsKnown() = false", v)
		}
	}
	if CommandSource("unknown").IsKnown() {
		t.Error("unknown should not be known")
	}
}

// TestCommandListParamsURLQuery verifies query serialization for GET /command.
// Aligned with OpenAPI: directory?, workspace?
func TestCommandListParamsURLQuery(t *testing.T) {
	p := CommandListParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestCommandListParamsURLQueryEmpty verifies query serialization with no params.
func TestCommandListParamsURLQueryEmpty(t *testing.T) {
	p := CommandListParams{}
	got := p.URLQuery().Encode()
	if got != "" {
		t.Errorf("expected empty query, got %q", got)
	}
}
