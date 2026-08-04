// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
)

// TestLocationInfoUnmarshal verifies that LocationInfo correctly deserialises
// a JSON payload that includes all fields (including the optional workspaceID).
func TestLocationInfoUnmarshal(t *testing.T) {
	t.Parallel()

	raw := `{
		"directory": "/workspace/myproject",
		"workspaceID": "wrk_abc123",
		"project": {
			"id": "proj_001",
			"directory": "/workspace/myproject"
		}
	}`

	var got opencode.LocationInfo
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if got.Directory != "/workspace/myproject" {
		t.Errorf("Directory: got %q, want %q", got.Directory, "/workspace/myproject")
	}
	if got.WorkspaceID != "wrk_abc123" {
		t.Errorf("WorkspaceID: got %q, want %q", got.WorkspaceID, "wrk_abc123")
	}
	if got.Project.ID != "proj_001" {
		t.Errorf("Project.ID: got %q, want %q", got.Project.ID, "proj_001")
	}
	if got.Project.Directory != "/workspace/myproject" {
		t.Errorf("Project.Directory: got %q, want %q", got.Project.Directory, "/workspace/myproject")
	}
}

// TestLocationInfoUnmarshalMinimal verifies that LocationInfo correctly deserialises
// a payload that omits the optional workspaceID field.
func TestLocationInfoUnmarshalMinimal(t *testing.T) {
	t.Parallel()

	raw := `{
		"directory": "/tmp/proj",
		"project": {
			"id": "proj_002",
			"directory": "/tmp/proj"
		}
	}`

	var got opencode.LocationInfo
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if got.Directory != "/tmp/proj" {
		t.Errorf("Directory: got %q, want %q", got.Directory, "/tmp/proj")
	}
	if got.WorkspaceID != "" {
		t.Errorf("WorkspaceID: expected empty string for absent field, got %q", got.WorkspaceID)
	}
	if got.Project.ID != "proj_002" {
		t.Errorf("Project.ID: got %q, want %q", got.Project.ID, "proj_002")
	}
}

// TestV2SkillListResponseUnmarshal verifies full round-trip deserialization of
// V2SkillListResponse including its nested SkillV2Info list and LocationInfo.
func TestV2SkillListResponseUnmarshal(t *testing.T) {
	t.Parallel()

	raw := `{
		"location": {
			"directory": "/workspace/proj",
			"project": { "id": "proj_x", "directory": "/workspace/proj" }
		},
		"data": [
			{
				"name": "my-skill",
				"description": "does stuff",
				"slash": true,
				"location": "/path/to/skill",
				"content": "# Skill content"
			},
			{
				"name": "bare-skill",
				"location": "/bare",
				"content": "bare"
			}
		]
	}`

	var got opencode.V2SkillListResponse
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if got.Location.Directory != "/workspace/proj" {
		t.Errorf("Location.Directory: got %q, want /workspace/proj", got.Location.Directory)
	}
	if len(got.Data) != 2 {
		t.Fatalf("Data length: got %d, want 2", len(got.Data))
	}
	if got.Data[0].Name != "my-skill" {
		t.Errorf("Data[0].Name: got %q, want my-skill", got.Data[0].Name)
	}
	if got.Data[0].Description != "does stuff" {
		t.Errorf("Data[0].Description: got %q, want 'does stuff'", got.Data[0].Description)
	}
	if !got.Data[0].Slash {
		t.Errorf("Data[0].Slash: got false, want true")
	}
	if got.Data[1].Name != "bare-skill" {
		t.Errorf("Data[1].Name: got %q, want bare-skill", got.Data[1].Name)
	}
}
