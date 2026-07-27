package opencode

import (
	"encoding/json"
	"testing"
)

// Aligned with OpenAPI GET /experimental/tool + JS SDK(v2) Tool.list.
// query required: provider, model; optional: directory, workspace
func TestToolListParamsQuery(t *testing.T) {
	t.Run("full query", func(t *testing.T) {
		p := ToolListParams{
			Directory: F("d"),
			Workspace: F("w"),
			Provider:  F("anthropic"),
			Model:     F("claude-3-5-sonnet"),
		}
		got := p.URLQuery().Encode()
		want := "directory=d&model=claude-3-5-sonnet&provider=anthropic&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("required-only", func(t *testing.T) {
		p := ToolListParams{Provider: F("openai"), Model: F("gpt-4")}
		got := p.URLQuery().Encode()
		want := "model=gpt-4&provider=openai"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

// Aligned with OpenAPI GET /experimental/tool/ids + JS SDK(v2) Tool.ids.
// query all optional: directory, workspace
func TestToolIdsParamsQuery(t *testing.T) {
	t.Run("both fields", func(t *testing.T) {
		p := ToolIdsParams{Directory: F("d"), Workspace: F("w")}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("empty", func(t *testing.T) {
		p := ToolIdsParams{}
		if got := p.URLQuery().Encode(); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
}

// Aligned with OpenAPI GET /experimental/tool 200 items (ToolListItem).
// required: id, description, parameters (free-form object)
func TestToolListItemUnmarshal(t *testing.T) {
	raw := `{
		"id":"bash",
		"description":"Run a bash command",
		"parameters":{
			"type":"object",
			"properties":{"command":{"type":"string"}},
			"required":["command"]
		}
	}`
	var item ToolListItem
	if err := json.Unmarshal([]byte(raw), &item); err != nil {
		t.Fatal(err)
	}
	if item.ID != "bash" {
		t.Errorf("id = %q", item.ID)
	}
	if item.Description != "Run a bash command" {
		t.Errorf("description = %q", item.Description)
	}
	// parameters is free-form (OpenAPI {} / JS unknown) -> Go any (map)
	params, ok := item.Parameters.(map[string]any)
	if !ok {
		t.Fatalf("parameters type = %T, want map[string]any", item.Parameters)
	}
	if params["type"] != "object" {
		t.Errorf("parameters.type = %v", params["type"])
	}
	if item.JSON.raw == "" {
		t.Error("RawJSON not preserved")
	}
}

// Array responses for tool.list and tool.ids.
func TestToolArrayResponsesUnmarshal(t *testing.T) {
	t.Run("tool.list []ToolListItem", func(t *testing.T) {
		raw := `[
			{"id":"read","description":"read file","parameters":{}},
			{"id":"write","description":"write file","parameters":{"type":"object"}}
		]`
		var arr []ToolListItem
		if err := json.Unmarshal([]byte(raw), &arr); err != nil {
			t.Fatal(err)
		}
		if len(arr) != 2 || arr[0].ID != "read" || arr[1].ID != "write" {
			t.Errorf("arr = %+v", arr)
		}
	})

	t.Run("tool.ids string array", func(t *testing.T) {
		raw := `["bash","read","write","edit"]`
		var arr []string
		if err := json.Unmarshal([]byte(raw), &arr); err != nil {
			t.Fatal(err)
		}
		if len(arr) != 4 || arr[3] != "edit" {
			t.Errorf("arr = %+v", arr)
		}
	})
}
