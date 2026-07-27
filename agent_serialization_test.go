package opencode

import (
	"encoding/json"
	"testing"
)

// Aligned with OpenAPI schema "Agent" and JS SDK(v2) `type Agent`.
// required: name, mode, permission, options
func TestAgentResponseDeserialization(t *testing.T) {
	t.Run("full agent with all fields", func(t *testing.T) {
		raw := `{"name":"build","description":"build agent","mode":"primary","native":true,"hidden":false,"topP":0.9,"temperature":0.7,"color":"blue","permission":[],"model":{"modelID":"claude","providerID":"anthropic"},"variant":"default","prompt":"do build","options":{"foo":"bar","n":1},"steps":5}`
		var a Agent
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatal(err)
		}
		if a.Name != "build" {
			t.Errorf("Name = %q, want build", a.Name)
		}
		if a.Description != "build agent" {
			t.Errorf("Description = %q, want build agent", a.Description)
		}
		if a.Mode != AgentModePrimary {
			t.Errorf("Mode = %q, want primary", a.Mode)
		}
		if a.Native != true {
			t.Errorf("Native = %v, want true", a.Native)
		}
		if a.Hidden != false {
			t.Errorf("Hidden = %v, want false", a.Hidden)
		}
		if a.TopP != 0.9 {
			t.Errorf("TopP = %v, want 0.9", a.TopP)
		}
		if a.Temperature != 0.7 {
			t.Errorf("Temperature = %v, want 0.7", a.Temperature)
		}
		if a.Color != "blue" {
			t.Errorf("Color = %q, want blue", a.Color)
		}
		if a.Model.ModelID != "claude" || a.Model.ProviderID != "anthropic" {
			t.Errorf("Model = %+v, want {claude anthropic}", a.Model)
		}
		if a.Variant != "default" {
			t.Errorf("Variant = %q, want default", a.Variant)
		}
		if a.Prompt != "do build" {
			t.Errorf("Prompt = %q, want do build", a.Prompt)
		}
		if a.Options["foo"] != "bar" {
			t.Errorf("Options[foo] = %v, want bar", a.Options["foo"])
		}
		// OpenAPI `steps` is number(count) -> Go int64
		if a.Steps != 5 {
			t.Errorf("Steps = %d, want 5", a.Steps)
		}
		if a.JSON.RawJSON() != raw {
			t.Errorf("RawJSON mismatch: %s", a.JSON.RawJSON())
		}
	})

	t.Run("minimal agent with only required fields", func(t *testing.T) {
		raw := `{"name":"minimal","mode":"subagent","permission":[],"options":{}}`
		var a Agent
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatal(err)
		}
		if a.Name != "minimal" {
			t.Errorf("Name = %q, want minimal", a.Name)
		}
		if a.Mode != AgentModeSubagent {
			t.Errorf("Mode = %q, want subagent", a.Mode)
		}
		if a.Options == nil {
			t.Error("Options should be non-nil empty map")
		}
		// optional numeric fields default to zero
		if a.Steps != 0 || a.Temperature != 0 || a.TopP != 0 {
			t.Errorf("optional numeric fields should be zero, got steps=%d temp=%v topP=%v", a.Steps, a.Temperature, a.TopP)
		}
	})

	t.Run("mode all variant", func(t *testing.T) {
		raw := `{"name":"a","mode":"all","permission":[],"options":{}}`
		var a Agent
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatal(err)
		}
		if a.Mode != AgentModeAll {
			t.Errorf("Mode = %q, want all", a.Mode)
		}
	})

	t.Run("array of agents (app.agents response)", func(t *testing.T) {
		raw := `[{"name":"a","mode":"primary","permission":[],"options":{}},{"name":"b","mode":"subagent","permission":[],"options":{}}]`
		var agents []Agent
		if err := json.Unmarshal([]byte(raw), &agents); err != nil {
			t.Fatal(err)
		}
		if len(agents) != 2 {
			t.Fatalf("len = %d, want 2", len(agents))
		}
		if agents[0].Name != "a" || agents[1].Name != "b" {
			t.Errorf("names = %q,%q want a,b", agents[0].Name, agents[1].Name)
		}
	})
}

func TestAgentModeIsKnown(t *testing.T) {
	for _, m := range []AgentMode{AgentModeSubagent, AgentModePrimary, AgentModeAll} {
		if !m.IsKnown() {
			t.Errorf("%q should be known", m)
		}
	}
	if AgentMode("unknown").IsKnown() {
		t.Error("unknown should not be known")
	}
}

// Aligned with OpenAPI /skill response items and JS SDK(v2).
// required: name, location, content (description optional)
func TestSkillResponseDeserialization(t *testing.T) {
	t.Run("full skill", func(t *testing.T) {
		raw := `{"name":"graphify","description":"knowledge graph","location":"/path/SKILL.md","content":"# body"}`
		var s Skill
		if err := json.Unmarshal([]byte(raw), &s); err != nil {
			t.Fatal(err)
		}
		if s.Name != "graphify" || s.Description != "knowledge graph" || s.Location != "/path/SKILL.md" || s.Content != "# body" {
			t.Errorf("Skill = %+v", s)
		}
		if s.JSON.RawJSON() != raw {
			t.Errorf("RawJSON mismatch: %s", s.JSON.RawJSON())
		}
	})

	t.Run("skill without optional description", func(t *testing.T) {
		raw := `{"name":"n","location":"/l","content":"c"}`
		var s Skill
		if err := json.Unmarshal([]byte(raw), &s); err != nil {
			t.Fatal(err)
		}
		if s.Name != "n" || s.Location != "/l" || s.Content != "c" {
			t.Errorf("Skill = %+v", s)
		}
		if s.Description != "" {
			t.Errorf("Description = %q, want empty", s.Description)
		}
	})

	t.Run("array of skills (app.skills response)", func(t *testing.T) {
		raw := `[{"name":"a","location":"/a","content":"ca"},{"name":"b","location":"/b","content":"cb"}]`
		var skills []Skill
		if err := json.Unmarshal([]byte(raw), &skills); err != nil {
			t.Fatal(err)
		}
		if len(skills) != 2 {
			t.Fatalf("len = %d, want 2", len(skills))
		}
	})
}

// Aligned with OpenAPI /agent query params (directory, workspace).
func TestAgentListParamsQuery(t *testing.T) {
	t.Run("both params", func(t *testing.T) {
		p := AgentListParams{Directory: F("d"), Workspace: F("w")}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty params", func(t *testing.T) {
		p := AgentListParams{}
		if got := p.URLQuery().Encode(); got != "" {
			t.Errorf("got %q, want empty", got)
		}
	})

	t.Run("only directory", func(t *testing.T) {
		p := AgentListParams{Directory: F("d")}
		if got := p.URLQuery().Encode(); got != "directory=d" {
			t.Errorf("got %q, want directory=d", got)
		}
	})
}

// AppSkillsParams shares the same query shape.
func TestAppSkillsParamsQuery(t *testing.T) {
	p := AppSkillsParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
