// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/sst/opencode-sdk-go"
)

func TestV2FsListResponseUnmarshal(t *testing.T) {
	t.Parallel()
	t.Run("parses location and data entries", func(t *testing.T) {
		t.Parallel()
		const body = `{
			"location": {
				"directory": "/repo",
				"workspaceID": "ws_1",
				"project": {"id": "prj_1", "directory": "/repo"}
			},
			"data": [
				{"path": "a.txt", "type": "file"},
				{"path": "src", "type": "directory"}
			]
		}`
		var res opencode.V2FsListResponse
		if err := json.Unmarshal([]byte(body), &res); err != nil {
			t.Fatal(err)
		}
		if res.Location.Directory != "/repo" {
			t.Errorf("Location.Directory = %q, want %q", res.Location.Directory, "/repo")
		}
		if res.Location.Project.ID != "prj_1" {
			t.Errorf("Location.Project.ID = %q, want %q", res.Location.Project.ID, "prj_1")
		}
		if len(res.Data) != 2 {
			t.Fatalf("len(Data) = %d, want 2", len(res.Data))
		}
		if res.Data[0].Path != "a.txt" || res.Data[0].Type != opencode.FileSystemEntryTypeFile {
			t.Errorf("Data[0] = %+v, want {a.txt file}", res.Data[0])
		}
		if res.Data[1].Type != opencode.FileSystemEntryTypeDirectory {
			t.Errorf("Data[1].Type = %q, want %q", res.Data[1].Type, opencode.FileSystemEntryTypeDirectory)
		}
	})

	t.Run("raw JSON preserved for unknown fields", func(t *testing.T) {
		t.Parallel()
		const body = `{"location":{"directory":"/x","project":{"id":"p","directory":"/x"}},"data":[],"extra":"kept"}`
		var res opencode.V2FsListResponse
		if err := json.Unmarshal([]byte(body), &res); err != nil {
			t.Fatal(err)
		}
		if res.JSON.RawJSON() == "" {
			t.Error("RawJSON() returned empty string")
		}
	})
}

func TestV2FsFindResponseUnmarshal(t *testing.T) {
	t.Parallel()
	const body = `{
		"location": {"directory": "/repo", "project": {"id": "prj_1", "directory": "/repo"}},
		"data": [{"path": "go.mod", "type": "file"}]
	}`
	var res opencode.V2FsFindResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		t.Fatal(err)
	}
	if len(res.Data) != 1 {
		t.Fatalf("len(Data) = %d, want 1", len(res.Data))
	}
	if res.Data[0].Path != "go.mod" || res.Data[0].Type != opencode.FileSystemEntryTypeFile {
		t.Errorf("Data[0] = %+v, want {go.mod file}", res.Data[0])
	}
	if res.JSON.Data.IsMissing() {
		t.Error("Data presence metadata missing after unmarshal")
	}
}

func TestFileSystemEntryTypeIsKnown(t *testing.T) {
	t.Parallel()
	for _, v := range []opencode.FileSystemEntryType{
		opencode.FileSystemEntryTypeFile,
		opencode.FileSystemEntryTypeDirectory,
	} {
		if !v.IsKnown() {
			t.Errorf("IsKnown() = false for legal value %q", v)
		}
	}
	for _, v := range []opencode.FileSystemEntryType{"", "unknown", "link", "file "} {
		if v.IsKnown() {
			t.Errorf("IsKnown() = true for illegal value %q", v)
		}
	}
}

// TestV2FsReadParamsDeepObjectSerialization verifies that V2FsReadParams
// serializes the nested location as a deepObject query string
// (location[directory]=..., location[workspace]=...) matching the OpenAPI
// style: deepObject, explode: true.
func TestV2FsReadParamsDeepObjectSerialization(t *testing.T) {
	t.Parallel()

	params := opencode.V2FsReadParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/repo"),
			Workspace: opencode.F("wrk_1"),
		}),
	}
	got := params.URLQuery()

	want := url.Values{
		"location[directory]": []string{"/repo"},
		"location[workspace]": []string{"wrk_1"},
	}

	if len(got) != len(want) {
		t.Fatalf("URLQuery() length: got %d keys %v, want %d keys %v", len(got), got, len(want), want)
	}
	for k, wantVals := range want {
		gotVals, ok := got[k]
		if !ok {
			t.Errorf("URLQuery() missing key %q; got keys: %v", k, got)
			continue
		}
		if len(gotVals) != len(wantVals) || gotVals[0] != wantVals[0] {
			t.Errorf("URLQuery()[%q]: got %v, want %v", k, gotVals, wantVals)
		}
	}
}

// TestV2FsReadParamsEmptyLocation verifies that an omitted location is not
// emitted in the serialized query string.
func TestV2FsReadParamsEmptyLocation(t *testing.T) {
	t.Parallel()

	params := opencode.V2FsReadParams{}
	got := params.URLQuery()
	if len(got) != 0 {
		t.Errorf("URLQuery() with empty params: expected no keys, got %v", got)
	}
}
