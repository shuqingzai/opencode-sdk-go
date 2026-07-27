package opencode

import (
	"encoding/json"
	"testing"
)

// Aligned with OpenAPI GET /formatter + JS SDK(v2) Formatter.status.
// query all optional: directory, workspace
func TestFormatterStatusParamsQuery(t *testing.T) {
	t.Run("both fields", func(t *testing.T) {
		p := FormatterStatusParams{Directory: F("d"), Workspace: F("w")}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("empty", func(t *testing.T) {
		p := FormatterStatusParams{}
		if got := p.URLQuery().Encode(); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
}

// Aligned with OpenAPI GET /formatter 200 items (FormatterStatus).
// required: name, extensions, enabled
func TestFormatterStatusUnmarshal(t *testing.T) {
	raw := `{"name":"prettier","extensions":[".js",".ts",".json"],"enabled":true}`
	var f FormatterStatus
	if err := json.Unmarshal([]byte(raw), &f); err != nil {
		t.Fatal(err)
	}
	if f.Name != "prettier" {
		t.Errorf("name = %q", f.Name)
	}
	if len(f.Extensions) != 3 || f.Extensions[2] != ".json" {
		t.Errorf("extensions = %+v", f.Extensions)
	}
	if !f.Enabled {
		t.Error("enabled should be true")
	}
	if f.JSON.raw == "" {
		t.Error("RawJSON not preserved")
	}
}

// Array response and empty-extensions boundary.
func TestFormatterStatusArrayUnmarshal(t *testing.T) {
	raw := `[
		{"name":"gofmt","extensions":[".go"],"enabled":true},
		{"name":"disabled-fmt","extensions":[],"enabled":false}
	]`
	var arr []FormatterStatus
	if err := json.Unmarshal([]byte(raw), &arr); err != nil {
		t.Fatal(err)
	}
	if len(arr) != 2 {
		t.Fatalf("arr len = %d", len(arr))
	}
	if arr[0].Name != "gofmt" || !arr[0].Enabled {
		t.Errorf("arr[0] = %+v", arr[0])
	}
	if len(arr[1].Extensions) != 0 || arr[1].Enabled {
		t.Errorf("arr[1] = %+v", arr[1])
	}
}
