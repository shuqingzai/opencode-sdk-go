// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"net/url"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
)

// TestV2LocationGetParamsDeepObjectSerialization verifies that V2LocationGetParams
// correctly serializes the nested V2LocationParam as a deepObject query string,
// i.e. location[directory]=... and location[workspace]=..., as required by the
// OpenAPI spec (style: deepObject, explode: true).
//
// This test also serves as a regression guard: V2LocationParam intentionally does
// NOT implement URLQuery() because it is a nested value object, not a top-level
// Params struct. The apiquery encoder recurses into its query: tagged fields via
// the param.Field[T] path, producing the correct bracket-encoded keys without any
// URLQuery() on the inner type.
func TestV2LocationGetParamsDeepObjectSerialization(t *testing.T) {
	t.Parallel()

	params := opencode.V2LocationGetParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/a"),
			Workspace: opencode.F("wrk_1"),
		}),
	}

	got := params.URLQuery()

	want := url.Values{
		"location[directory]": []string{"/a"},
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

// TestV2LocationGetParamsEmptyLocation verifies that omitted (zero-value) location
// fields are not emitted in the serialized query string.
func TestV2LocationGetParamsEmptyLocation(t *testing.T) {
	t.Parallel()

	params := opencode.V2LocationGetParams{}
	got := params.URLQuery()

	if len(got) != 0 {
		t.Errorf("URLQuery() with empty params: expected no keys, got %v", got)
	}
}

// TestV2LocationGetParamsPartialLocation verifies that only present fields are
// serialized when only one of the nested fields is set.
func TestV2LocationGetParamsPartialLocation(t *testing.T) {
	t.Parallel()

	params := opencode.V2LocationGetParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/workspace/myproject"),
		}),
	}

	got := params.URLQuery()

	if _, ok := got["location[directory]"]; !ok {
		t.Errorf("URLQuery() missing location[directory]; got: %v", got)
	}
	if _, ok := got["location[workspace]"]; ok {
		t.Errorf("URLQuery() should not contain location[workspace] when Workspace is not set; got: %v", got)
	}
	if got.Get("location[directory]") != "/workspace/myproject" {
		t.Errorf("URLQuery()[location[directory]] = %q, want %q", got.Get("location[directory]"), "/workspace/myproject")
	}
}
