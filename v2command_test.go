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

// TestV2CommandListWithOptionalParams tests the v2 command list endpoint.
// Aligned with OpenAPI operationId "v2.command.list", GET /api/command.
// Query parameters: location (optional, nested with directory and workspace).
func TestV2CommandListWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Command.List(context.TODO(), opencode.V2CommandListParams{
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

// TestV2CommandListParamsURLQuery verifies query serialization for V2CommandListParams.
func TestV2CommandListParamsURLQuery(t *testing.T) {
	t.Run("with full location", func(t *testing.T) {
		p := opencode.V2CommandListParams{
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
		p := opencode.V2CommandListParams{}
		got := p.URLQuery().Encode()
		if got != "" {
			t.Errorf("expected empty query, got %q", got)
		}
	})
}

// TestV2CommandListResponseUnmarshal verifies deserialization of V2CommandListResponse.
// Aligned with OpenAPI GET /api/command response schema.
// Required fields: location (LocationInfo), data ([]V2CommandInfo).
func TestV2CommandListResponseUnmarshal(t *testing.T) {
	t.Run("full response with commands", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/home/user/project",
				"workspaceID": "ws_1",
				"project": {"id": "proj_1", "directory": "/home/user/project"}
			},
			"data": [
				{
					"name": "build",
					"template": "npm run build",
					"description": "Build the project",
					"agent": "build",
					"subtask": false
				},
				{
					"name": "test",
					"template": "npm test",
					"description": "Run tests",
					"agent": "test",
					"model": {"id": "claude-3", "providerID": "anthropic"},
					"subtask": true
				}
			]
		}`
		var resp opencode.V2CommandListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if resp.Location.Directory != "/home/user/project" {
			t.Errorf("Location.Directory = %q, want /home/user/project", resp.Location.Directory)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("Data len = %d, want 2", len(resp.Data))
		}
		cmd0 := resp.Data[0]
		if cmd0.Name != "build" {
			t.Errorf("cmd0.Name = %q, want build", cmd0.Name)
		}
		if cmd0.Template != "npm run build" {
			t.Errorf("cmd0.Template = %q, want npm run build", cmd0.Template)
		}
		if cmd0.Description != "Build the project" {
			t.Errorf("cmd0.Description = %q, want Build the project", cmd0.Description)
		}
		if cmd0.Subtask {
			t.Error("cmd0.Subtask should be false")
		}
		cmd1 := resp.Data[1]
		if cmd1.Name != "test" {
			t.Errorf("cmd1.Name = %q, want test", cmd1.Name)
		}
		if !cmd1.Subtask {
			t.Error("cmd1.Subtask should be true")
		}
		if cmd1.Model.ID != "claude-3" {
			t.Errorf("cmd1.Model.ID = %q, want claude-3", cmd1.Model.ID)
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
		var resp opencode.V2CommandListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(resp.Data) != 0 {
			t.Errorf("expected 0 commands, got %d", len(resp.Data))
		}
	})
}

// TestV2CommandInfoUnmarshal verifies deserialization of V2CommandInfo.
// Aligned with OpenAPI CommandInfo schema.
// Required: name, template. Optional: description, agent, model, subtask.
func TestV2CommandInfoUnmarshal(t *testing.T) {
	t.Run("full command with all optional fields", func(t *testing.T) {
		raw := `{
			"name": "deploy",
			"template": "npm run deploy -- {{env}}",
			"description": "Deploy to environment",
			"agent": "deploy-agent",
			"model": {"id": "gpt-4", "providerID": "openai", "variant": "fast"},
			"subtask": true
		}`
		var cmd opencode.V2CommandInfo
		if err := json.Unmarshal([]byte(raw), &cmd); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if cmd.Name != "deploy" {
			t.Errorf("Name = %q, want deploy", cmd.Name)
		}
		if cmd.Template != "npm run deploy -- {{env}}" {
			t.Errorf("Template = %q", cmd.Template)
		}
		if cmd.Description != "Deploy to environment" {
			t.Errorf("Description = %q, want Deploy to environment", cmd.Description)
		}
		if cmd.Agent != "deploy-agent" {
			t.Errorf("Agent = %q, want deploy-agent", cmd.Agent)
		}
		if cmd.Model.ID != "gpt-4" {
			t.Errorf("Model.ID = %q, want gpt-4", cmd.Model.ID)
		}
		if cmd.Model.ProviderID != "openai" {
			t.Errorf("Model.ProviderID = %q, want openai", cmd.Model.ProviderID)
		}
		if cmd.Model.Variant != "fast" {
			t.Errorf("Model.Variant = %q, want fast", cmd.Model.Variant)
		}
		if !cmd.Subtask {
			t.Error("Subtask should be true")
		}
		if cmd.JSON.RawJSON() == "" {
			t.Error("RawJSON should be preserved")
		}
	})

	t.Run("minimal command with required fields only", func(t *testing.T) {
		raw := `{"name":"lint","template":"eslint ."}`
		var cmd opencode.V2CommandInfo
		if err := json.Unmarshal([]byte(raw), &cmd); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if cmd.Name != "lint" {
			t.Errorf("Name = %q, want lint", cmd.Name)
		}
		if cmd.Template != "eslint ." {
			t.Errorf("Template = %q, want eslint .", cmd.Template)
		}
		if cmd.Description != "" {
			t.Errorf("Description should be empty, got %q", cmd.Description)
		}
		if cmd.Agent != "" {
			t.Errorf("Agent should be empty, got %q", cmd.Agent)
		}
		if cmd.Subtask {
			t.Error("Subtask should be false when absent")
		}
	})

	t.Run("unknown fields tolerated via ExtraFields", func(t *testing.T) {
		raw := `{"name":"x","template":"echo x","future_field":"value"}`
		var cmd opencode.V2CommandInfo
		if err := json.Unmarshal([]byte(raw), &cmd); err != nil {
			t.Fatalf("unmarshal with unknown field: %v", err)
		}
		if cmd.Name != "x" {
			t.Errorf("Name = %q, want x", cmd.Name)
		}
	})
}
