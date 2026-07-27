package opencode

import (
	"encoding/json"
	"testing"
)

// Aligned with OpenAPI POST /instance/dispose + JS SDK(v2) Instance.dispose.
// query all optional: directory, workspace
func TestInstanceDisposeParamsQuery(t *testing.T) {
	t.Run("both fields", func(t *testing.T) {
		p := InstanceDisposeParams{Directory: F("d"), Workspace: F("w")}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("empty", func(t *testing.T) {
		p := InstanceDisposeParams{}
		if got := p.URLQuery().Encode(); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
	t.Run("directory-only", func(t *testing.T) {
		p := InstanceDisposeParams{Directory: F("d")}
		got := p.URLQuery().Encode()
		want := "directory=d"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

// Aligned with OpenAPI POST /instance/dispose 200 response (boolean).
func TestInstanceDisposeResponseUnmarshal(t *testing.T) {
	for _, want := range []bool{true, false} {
		raw := "false"
		if want {
			raw = "true"
		}
		var res bool
		if err := json.Unmarshal([]byte(raw), &res); err != nil {
			t.Fatal(err)
		}
		if res != want {
			t.Errorf("got %v, want %v", res, want)
		}
	}
}
