package opencode

import (
	"encoding/json"
	"testing"
)

// Aligned with OpenAPI GET /lsp + JS SDK(v2) Lsp.status.
// query all optional: directory, workspace
func TestLspStatusParamsQuery(t *testing.T) {
	t.Run("both fields", func(t *testing.T) {
		p := LspStatusParams{Directory: F("d"), Workspace: F("w")}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("empty", func(t *testing.T) {
		p := LspStatusParams{}
		if got := p.URLQuery().Encode(); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
}

// Aligned with OpenAPI GET /lsp 200 items (LspStatus).
// required: id, name, root, status
func TestLspStatusUnmarshal(t *testing.T) {
	raw := `{"id":"gopls","name":"Go Language Server","root":"/repo","status":"connected"}`
	var s LspStatus
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatal(err)
	}
	if s.ID != "gopls" || s.Name != "Go Language Server" || s.Root != "/repo" {
		t.Errorf("lsp = %+v", s)
	}
	if s.Status != LspStatusStatusConnected {
		t.Errorf("status = %q", s.Status)
	}
	if s.JSON.raw == "" {
		t.Error("RawJSON not preserved")
	}
}

// Array response + error status boundary.
func TestLspStatusArrayUnmarshal(t *testing.T) {
	raw := `[
		{"id":"gopls","name":"gopls","root":"/a","status":"connected"},
		{"id":"tsserver","name":"tsserver","root":"/b","status":"error"}
	]`
	var arr []LspStatus
	if err := json.Unmarshal([]byte(raw), &arr); err != nil {
		t.Fatal(err)
	}
	if len(arr) != 2 {
		t.Fatalf("arr len = %d", len(arr))
	}
	if arr[0].Status != LspStatusStatusConnected || arr[1].Status != LspStatusStatusError {
		t.Errorf("statuses = %q, %q", arr[0].Status, arr[1].Status)
	}
}

func TestLspStatusStatusIsKnown(t *testing.T) {
	for _, s := range []LspStatusStatus{LspStatusStatusConnected, LspStatusStatusError} {
		if !s.IsKnown() {
			t.Errorf("%q should be known", s)
		}
	}
	if LspStatusStatus("stopped").IsKnown() {
		t.Error("stopped should not be known")
	}
}
