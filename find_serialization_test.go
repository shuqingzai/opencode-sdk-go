package opencode

import (
	"encoding/json"
	"testing"
)

// Aligned with OpenAPI GET /find/file + JS SDK(v2) Find.files.
// query required: query; optional: directory, workspace, dirs, type, limit
func TestFindFilesParamsQuery(t *testing.T) {
	t.Run("full query with enums and limit", func(t *testing.T) {
		p := FindFilesParams{
			Query:     F("main"),
			Directory: F("d"),
			Workspace: F("w"),
			Dirs:      F(FindFilesParamsDirsTrue),
			Type:      F(FindFilesParamsTypeFile),
			Limit:     F(int64(50)),
		}
		got := p.URLQuery().Encode()
		want := "directory=d&dirs=true&limit=50&query=main&type=file&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("required-only query", func(t *testing.T) {
		p := FindFilesParams{Query: F("q")}
		got := p.URLQuery().Encode()
		want := "query=q"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("limit boundary values 1 and 200", func(t *testing.T) {
		for _, lim := range []int64{1, 200} {
			p := FindFilesParams{Query: F("q"), Limit: F(lim)}
			got := p.URLQuery().Get("limit")
			if got == "" {
				t.Errorf("limit %d not serialized", lim)
			}
		}
	})

	t.Run("dirs and type variants", func(t *testing.T) {
		p := FindFilesParams{
			Query: F("q"),
			Dirs:  F(FindFilesParamsDirsFalse),
			Type:  F(FindFilesParamsTypeDirectory),
		}
		got := p.URLQuery().Encode()
		want := "dirs=false&query=q&type=directory"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestFindFilesParamsDirsIsKnown(t *testing.T) {
	for _, d := range []FindFilesParamsDirs{FindFilesParamsDirsTrue, FindFilesParamsDirsFalse} {
		if !d.IsKnown() {
			t.Errorf("%q should be known", d)
		}
	}
	if FindFilesParamsDirs("maybe").IsKnown() {
		t.Error("maybe should not be known")
	}
}

func TestFindFilesParamsTypeIsKnown(t *testing.T) {
	for _, ty := range []FindFilesParamsType{FindFilesParamsTypeFile, FindFilesParamsTypeDirectory} {
		if !ty.IsKnown() {
			t.Errorf("%q should be known", ty)
		}
	}
	if FindFilesParamsType("symlink").IsKnown() {
		t.Error("symlink should not be known")
	}
}

// Aligned with OpenAPI GET /find + JS SDK(v2) Find.text.
// query required: pattern; optional: directory, workspace
func TestFindTextParamsQuery(t *testing.T) {
	p := FindTextParams{
		Pattern:   F("TODO"),
		Directory: F("d"),
		Workspace: F("w"),
	}
	got := p.URLQuery().Encode()
	want := "directory=d&pattern=TODO&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI GET /find/symbol + JS SDK(v2) Find.symbols.
// query required: query; optional: directory, workspace
func TestFindSymbolsParamsQuery(t *testing.T) {
	p := FindSymbolsParams{
		Query:     F("Foo"),
		Directory: F("d"),
		Workspace: F("w"),
	}
	got := p.URLQuery().Encode()
	want := "directory=d&query=Foo&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI GET /find 200 response items.
// required: path, lines, line_number, absolute_offset, submatches
func TestFindTextResponseUnmarshal(t *testing.T) {
	raw := `{
		"path": {"text": "src/main.go"},
		"lines": {"text": "func main() {"},
		"line_number": 12,
		"absolute_offset": 340,
		"submatches": [
			{"match": {"text": "main"}, "start": 5, "end": 9}
		]
	}`
	var r FindTextResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatal(err)
	}
	if r.Path.Text != "src/main.go" {
		t.Errorf("Path.Text = %q", r.Path.Text)
	}
	if r.Lines.Text != "func main() {" {
		t.Errorf("Lines.Text = %q", r.Lines.Text)
	}
	if r.LineNumber != 12 {
		t.Errorf("LineNumber = %d", r.LineNumber)
	}
	if r.AbsoluteOffset != 340 {
		t.Errorf("AbsoluteOffset = %d", r.AbsoluteOffset)
	}
	if len(r.Submatches) != 1 {
		t.Fatalf("Submatches len = %d", len(r.Submatches))
	}
	sm := r.Submatches[0]
	if sm.Match.Text != "main" || sm.Start != 5 || sm.End != 9 {
		t.Errorf("submatch = %+v", sm)
	}
	if r.JSON.raw == "" {
		t.Error("RawJSON not preserved")
	}
}

// Aligned with OpenAPI GET /find/symbol 200 response items.
// required: name, kind, location{uri, range{start, end}}
func TestSymbolUnmarshal(t *testing.T) {
	raw := `{
		"name": "MyFunc",
		"kind": 12,
		"location": {
			"uri": "file:///src/main.go",
			"range": {
				"start": {"line": 10, "character": 0},
				"end": {"line": 10, "character": 6}
			}
		}
	}`
	var s Symbol
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatal(err)
	}
	if s.Name != "MyFunc" || s.Kind != 12 {
		t.Errorf("symbol = %+v", s)
	}
	if s.Location.URI != "file:///src/main.go" {
		t.Errorf("URI = %q", s.Location.URI)
	}
	if s.Location.Range.Start.Line != 10 || s.Location.Range.Start.Character != 0 {
		t.Errorf("start = %+v", s.Location.Range.Start)
	}
	if s.Location.Range.End.Line != 10 || s.Location.Range.End.Character != 6 {
		t.Errorf("end = %+v", s.Location.Range.End)
	}
}

// Array responses for find.text and find.symbols.
func TestFindArrayResponsesUnmarshal(t *testing.T) {
	t.Run("find.text array", func(t *testing.T) {
		raw := `[{"path":{"text":"a"},"lines":{"text":"b"},"line_number":1,"absolute_offset":2,"submatches":[]}]`
		var arr []FindTextResponse
		if err := json.Unmarshal([]byte(raw), &arr); err != nil {
			t.Fatal(err)
		}
		if len(arr) != 1 || arr[0].Path.Text != "a" {
			t.Errorf("arr = %+v", arr)
		}
	})

	t.Run("find.files string array", func(t *testing.T) {
		raw := `["a.go","b.go","c.go"]`
		var arr []string
		if err := json.Unmarshal([]byte(raw), &arr); err != nil {
			t.Fatal(err)
		}
		if len(arr) != 3 || arr[2] != "c.go" {
			t.Errorf("arr = %+v", arr)
		}
	})
}
