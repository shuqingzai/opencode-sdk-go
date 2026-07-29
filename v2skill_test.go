// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// TestV2SkillListWithOptionalParams tests the v2 skill list endpoint.
// Aligned with OpenAPI operationId "v2.skill.list", GET /api/skill.
// Query parameters: location (optional, nested with directory and workspace).
func TestV2SkillListWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Skill.List(context.TODO(), opencode.V2SkillListParams{
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

// TestV2SkillListParamsURLQuery verifies query serialization for V2SkillListParams.
func TestV2SkillListParamsURLQuery(t *testing.T) {
	t.Run("with full location", func(t *testing.T) {
		p := opencode.V2SkillListParams{
			Location: opencode.F(opencode.V2LocationParam{
				Directory: opencode.F("mydir"),
				Workspace: opencode.F("mywsp"),
			}),
		}
		got := p.URLQuery().Encode()
		if got == "" {
			t.Error("expected non-empty query string")
		}
	})

	t.Run("empty params produces no query", func(t *testing.T) {
		p := opencode.V2SkillListParams{}
		got := p.URLQuery().Encode()
		if got != "" {
			t.Errorf("expected empty query, got %q", got)
		}
	})
}

// TestV2SkillListResponseUnmarshal verifies deserialization of V2SkillListResponse.
// Aligned with OpenAPI GET /api/skill response schema.
// Required fields: location (LocationInfo), data ([]SkillV2Info).
func TestV2SkillListResponseUnmarshal(t *testing.T) {
	t.Run("full response with skills", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/home/user/project",
				"workspaceID": "ws_1",
				"project": {"id": "proj_1", "directory": "/home/user/project"}
			},
			"data": [
				{
					"name": "brainstorming",
					"description": "Explore ideas before implementing",
					"slash": true,
					"location": "/path/to/SKILL.md",
					"content": "# Brainstorming\nExplore user intent..."
				},
				{
					"name": "graphify",
					"location": "/path/to/graphify/SKILL.md",
					"content": "# Graphify\nTurn any folder..."
				}
			]
		}`
		var resp opencode.V2SkillListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if resp.Location.Directory != "/home/user/project" {
			t.Errorf("Location.Directory = %q, want /home/user/project", resp.Location.Directory)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("Data len = %d, want 2", len(resp.Data))
		}
		s0 := resp.Data[0]
		if s0.Name != "brainstorming" {
			t.Errorf("s0.Name = %q, want brainstorming", s0.Name)
		}
		if s0.Description != "Explore ideas before implementing" {
			t.Errorf("s0.Description = %q", s0.Description)
		}
		if !s0.Slash {
			t.Error("s0.Slash should be true")
		}
		if s0.Location != "/path/to/SKILL.md" {
			t.Errorf("s0.Location = %q, want /path/to/SKILL.md", s0.Location)
		}
		if s0.Content != "# Brainstorming\nExplore user intent..." {
			t.Errorf("s0.Content = %q", s0.Content)
		}
		s1 := resp.Data[1]
		if s1.Name != "graphify" {
			t.Errorf("s1.Name = %q, want graphify", s1.Name)
		}
		if s1.Description != "" {
			t.Errorf("s1.Description should be empty when absent, got %q", s1.Description)
		}
		if resp.JSON.RawJSON() == "" {
			t.Error("RawJSON should be preserved")
		}
	})

	t.Run("empty data array", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/",
				"project": {"id": "p", "directory": "/"}
			},
			"data": []
		}`
		var resp opencode.V2SkillListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(resp.Data) != 0 {
			t.Errorf("expected 0 skills, got %d", len(resp.Data))
		}
	})
}

// TestSkillV2InfoUnmarshal verifies deserialization of SkillV2Info.
// Aligned with OpenAPI SkillInfo schema.
// Required: name, location, content. Optional: description, slash.
func TestSkillV2InfoUnmarshal(t *testing.T) {
	t.Run("full skill with all fields", func(t *testing.T) {
		raw := `{
			"name": "systematic-debugging",
			"description": "Debug systematically before proposing fixes",
			"slash": false,
			"location": "/path/to/skill/SKILL.md",
			"content": "# Systematic Debugging\nUse when encountering any bug..."
		}`
		var skill opencode.SkillV2Info
		if err := json.Unmarshal([]byte(raw), &skill); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if skill.Name != "systematic-debugging" {
			t.Errorf("Name = %q, want systematic-debugging", skill.Name)
		}
		if skill.Description != "Debug systematically before proposing fixes" {
			t.Errorf("Description = %q", skill.Description)
		}
		if skill.Slash {
			t.Error("Slash should be false")
		}
		if skill.Location != "/path/to/skill/SKILL.md" {
			t.Errorf("Location = %q, want /path/to/skill/SKILL.md", skill.Location)
		}
		if skill.Content != "# Systematic Debugging\nUse when encountering any bug..." {
			t.Errorf("Content = %q", skill.Content)
		}
		if skill.JSON.RawJSON() == "" {
			t.Error("RawJSON should be preserved")
		}
	})

	t.Run("skill with required fields only", func(t *testing.T) {
		raw := `{"name":"quick","location":"/q/SKILL.md","content":"quick skill"}`
		var skill opencode.SkillV2Info
		if err := json.Unmarshal([]byte(raw), &skill); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if skill.Description != "" {
			t.Errorf("Description should be empty, got %q", skill.Description)
		}
		if skill.Slash {
			t.Error("Slash should default to false")
		}
	})

	t.Run("unknown fields tolerated via ExtraFields", func(t *testing.T) {
		raw := `{"name":"x","location":"/x","content":"c","new_v2_field":"value"}`
		var skill opencode.SkillV2Info
		if err := json.Unmarshal([]byte(raw), &skill); err != nil {
			t.Fatalf("unmarshal with unknown field: %v", err)
		}
		if skill.Name != "x" {
			t.Errorf("Name = %q, want x", skill.Name)
		}
	})
}
