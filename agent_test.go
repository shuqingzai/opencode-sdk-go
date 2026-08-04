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

func TestAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.App.Agents(context.TODO(), opencode.AgentListParams{
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

// TestAgentUnmarshal verifies that Agent can be deserialized from JSON
// (OpenAPI: components.schemas.Agent).
//
// Run with: go test -run TestAgentUnmarshal -v ./...
func TestAgentUnmarshal(t *testing.T) {
	t.Parallel()
	const raw = `{
		"name": "coder",
		"mode": "primary",
		"permission": {},
		"options": {},
		"description": "A coding agent",
		"steps": 10,
		"temperature": 0.7,
		"topP": 0.9,
		"model": {"modelID": "gpt-4o", "providerID": "openai"}
	}`
	var a opencode.Agent
	if err := json.Unmarshal([]byte(raw), &a); err != nil {
		t.Fatalf("json.Unmarshal Agent: %v", err)
	}
	if a.Name != "coder" {
		t.Errorf("Name: got %q, want %q", a.Name, "coder")
	}
	if a.Mode != opencode.AgentModePrimary {
		t.Errorf("Mode: got %q, want %q", a.Mode, opencode.AgentModePrimary)
	}
	if a.Steps != 10 {
		t.Errorf("Steps: got %d, want 10", a.Steps)
	}
	if a.Temperature != 0.7 {
		t.Errorf("Temperature: got %v, want 0.7", a.Temperature)
	}
	if a.Model.ModelID != "gpt-4o" {
		t.Errorf("Model.ModelID: got %q, want %q", a.Model.ModelID, "gpt-4o")
	}
	if a.Model.ProviderID != "openai" {
		t.Errorf("Model.ProviderID: got %q, want %q", a.Model.ProviderID, "openai")
	}
	if a.JSON.Name.Raw() == "" {
		t.Error("JSON.Name.Raw() should be non-empty")
	}
}

// TestSkillUnmarshal verifies that Skill can be deserialized from JSON
// (OpenAPI: app.skills response item schema).
//
// Run with: go test -run TestSkillUnmarshal -v ./...
func TestSkillUnmarshal(t *testing.T) {
	t.Parallel()
	const raw = `{
		"name": "brainstorming",
		"location": "/path/to/SKILL.md",
		"content": "## Instructions\n...",
		"description": "Explore user intent before implementing"
	}`
	var s opencode.Skill
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatalf("json.Unmarshal Skill: %v", err)
	}
	if s.Name != "brainstorming" {
		t.Errorf("Name: got %q, want %q", s.Name, "brainstorming")
	}
	if s.Location != "/path/to/SKILL.md" {
		t.Errorf("Location: got %q, want %q", s.Location, "/path/to/SKILL.md")
	}
	if s.Content != "## Instructions\n..." {
		t.Errorf("Content: got %q, want %q", s.Content, "## Instructions\n...")
	}
	if s.Description != "Explore user intent before implementing" {
		t.Errorf("Description: got %q, want %q", s.Description, "Explore user intent before implementing")
	}
	if s.JSON.Name.Raw() == "" {
		t.Error("JSON.Name.Raw() should be non-empty")
	}
	// Variant: description absent (optional per OpenAPI)
	const rawNoDesc = `{"name":"x","location":"/x","content":"y"}`
	var s2 opencode.Skill
	if err := json.Unmarshal([]byte(rawNoDesc), &s2); err != nil {
		t.Fatalf("json.Unmarshal Skill (no description): %v", err)
	}
	if s2.Description != "" {
		t.Errorf("Description should be empty string when absent, got %q", s2.Description)
	}
}
