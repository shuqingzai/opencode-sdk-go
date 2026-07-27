package opencode

import (
	"encoding/json"
	"testing"
)

// Aligned with OpenAPI GET /file + JS SDK(v2) File.list.
// query required: path; optional: directory, workspace
func TestFileListParamsQuery(t *testing.T) {
	t.Run("full query", func(t *testing.T) {
		p := FileListParams{Path: F("src"), Directory: F("d"), Workspace: F("w")}
		got := p.URLQuery().Encode()
		want := "directory=d&path=src&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("required-only", func(t *testing.T) {
		p := FileListParams{Path: F("src")}
		got := p.URLQuery().Encode()
		want := "path=src"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

// Aligned with OpenAPI GET /file/content + JS SDK(v2) File.read.
func TestFileReadParamsQuery(t *testing.T) {
	p := FileReadParams{Path: F("a.go"), Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&path=a.go&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI GET /file/status + JS SDK(v2) File.status.
// query all optional: directory, workspace
func TestFileStatusParamsQuery(t *testing.T) {
	t.Run("both fields", func(t *testing.T) {
		p := FileStatusParams{Directory: F("d"), Workspace: F("w")}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("empty", func(t *testing.T) {
		p := FileStatusParams{}
		if got := p.URLQuery().Encode(); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
}

// Aligned with OpenAPI GET /file 200 items (FileNode).
// required: name, path, absolute, type, ignored
func TestFileNodeUnmarshal(t *testing.T) {
	raw := `{"name":"main.go","path":"src/main.go","absolute":"/repo/src/main.go","type":"file","ignored":false}`
	var n FileNode
	if err := json.Unmarshal([]byte(raw), &n); err != nil {
		t.Fatal(err)
	}
	if n.Name != "main.go" || n.Path != "src/main.go" || n.Absolute != "/repo/src/main.go" {
		t.Errorf("node = %+v", n)
	}
	if n.Type != FileNodeTypeFile {
		t.Errorf("type = %q", n.Type)
	}
	if n.Ignored {
		t.Error("ignored should be false")
	}
}

// Aligned with OpenAPI GET /file/status 200 items (File).
// required: path, added, removed, status
func TestFileStatusResponseUnmarshal(t *testing.T) {
	raw := `{"path":"src/main.go","added":10,"removed":3,"status":"modified"}`
	var f File
	if err := json.Unmarshal([]byte(raw), &f); err != nil {
		t.Fatal(err)
	}
	if f.Path != "src/main.go" || f.Added != 10 || f.Removed != 3 {
		t.Errorf("file = %+v", f)
	}
	if f.Status != FileStatusModified {
		t.Errorf("status = %q", f.Status)
	}
}

// Aligned with OpenAPI GET /file/content 200 (FileContent).
// required: type, content; optional: diff, patch, encoding, mimeType
func TestFileContentUnmarshal(t *testing.T) {
	t.Run("text minimal (required only)", func(t *testing.T) {
		raw := `{"type":"text","content":"hello"}`
		var c FileContent
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatal(err)
		}
		if c.Type != FileContentTypeText || c.Content != "hello" {
			t.Errorf("content = %+v", c)
		}
	})

	t.Run("binary with base64 encoding and mimeType", func(t *testing.T) {
		raw := `{"type":"binary","content":"aGVsbG8=","encoding":"base64","mimeType":"image/png"}`
		var c FileContent
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatal(err)
		}
		if c.Type != FileContentTypeBinary {
			t.Errorf("type = %q", c.Type)
		}
		if c.Encoding != FileContentEncodingBase64 {
			t.Errorf("encoding = %q", c.Encoding)
		}
		if c.MIMEType != "image/png" {
			t.Errorf("mimeType = %q", c.MIMEType)
		}
	})

	t.Run("full with diff and patch hunks", func(t *testing.T) {
		raw := `{
			"type":"text",
			"content":"new content",
			"diff":"@@ -1 +1 @@",
			"patch":{
				"oldFileName":"a.go",
				"newFileName":"b.go",
				"oldHeader":"old",
				"newHeader":"new",
				"index":"idx",
				"hunks":[
					{"oldStart":1,"oldLines":2,"newStart":1,"newLines":3,"lines":["-a","+b","+c"]}
				]
			}
		}`
		var c FileContent
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatal(err)
		}
		if c.Diff != "@@ -1 +1 @@" {
			t.Errorf("diff = %q", c.Diff)
		}
		if c.Patch.OldFileName != "a.go" || c.Patch.NewFileName != "b.go" {
			t.Errorf("patch names = %+v", c.Patch)
		}
		if c.Patch.OldHeader != "old" || c.Patch.NewHeader != "new" || c.Patch.Index != "idx" {
			t.Errorf("patch headers = %+v", c.Patch)
		}
		if len(c.Patch.Hunks) != 1 {
			t.Fatalf("hunks len = %d", len(c.Patch.Hunks))
		}
		h := c.Patch.Hunks[0]
		if h.OldStart != 1 || h.OldLines != 2 || h.NewStart != 1 || h.NewLines != 3 {
			t.Errorf("hunk = %+v", h)
		}
		if len(h.Lines) != 3 || h.Lines[2] != "+c" {
			t.Errorf("hunk lines = %+v", h.Lines)
		}
	})
}

// Array responses for file.list and file.status.
func TestFileArrayResponsesUnmarshal(t *testing.T) {
	t.Run("file.list []FileNode", func(t *testing.T) {
		raw := `[{"name":"a","path":"a","absolute":"/a","type":"file","ignored":false},
				{"name":"d","path":"d","absolute":"/d","type":"directory","ignored":true}]`
		var arr []FileNode
		if err := json.Unmarshal([]byte(raw), &arr); err != nil {
			t.Fatal(err)
		}
		if len(arr) != 2 || arr[1].Type != FileNodeTypeDirectory || !arr[1].Ignored {
			t.Errorf("arr = %+v", arr)
		}
	})

	t.Run("file.status []File", func(t *testing.T) {
		raw := `[{"path":"a","added":1,"removed":0,"status":"added"},
				{"path":"b","added":0,"removed":5,"status":"deleted"}]`
		var arr []File
		if err := json.Unmarshal([]byte(raw), &arr); err != nil {
			t.Fatal(err)
		}
		if len(arr) != 2 || arr[0].Status != FileStatusAdded || arr[1].Status != FileStatusDeleted {
			t.Errorf("arr = %+v", arr)
		}
	})
}

func TestFileStatusIsKnown(t *testing.T) {
	for _, s := range []FileStatus{FileStatusAdded, FileStatusDeleted, FileStatusModified} {
		if !s.IsKnown() {
			t.Errorf("%q should be known", s)
		}
	}
	if FileStatus("renamed").IsKnown() {
		t.Error("renamed should not be known")
	}
}

func TestFileNodeTypeIsKnown(t *testing.T) {
	for _, ty := range []FileNodeType{FileNodeTypeFile, FileNodeTypeDirectory} {
		if !ty.IsKnown() {
			t.Errorf("%q should be known", ty)
		}
	}
	if FileNodeType("symlink").IsKnown() {
		t.Error("symlink should not be known")
	}
}

func TestFileContentTypeIsKnown(t *testing.T) {
	for _, ty := range []FileContentType{FileContentTypeText, FileContentTypeBinary} {
		if !ty.IsKnown() {
			t.Errorf("%q should be known", ty)
		}
	}
	if FileContentType("json").IsKnown() {
		t.Error("json should not be known")
	}
}

func TestFileContentEncodingIsKnown(t *testing.T) {
	if !FileContentEncodingBase64.IsKnown() {
		t.Error("base64 should be known")
	}
	if FileContentEncoding("hex").IsKnown() {
		t.Error("hex should not be known")
	}
}
