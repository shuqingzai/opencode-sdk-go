// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
)

// ===== GlobalSessionTime.Archived =====
//
// OpenAPI `GlobalSession.time` is byte-identical to `Session.time`:
//
//	{"created":{"type":"integer","minimum":0},
//	 "updated":{"type":"integer","minimum":0},
//	 "compacting":{"type":"integer","minimum":0},
//	 "archived":{"type":"number"},
//	 "required":["created","updated"]}
//
// JS SDK v2 `GlobalSession.time` is `{created: number; updated: number;
// compacting?: number; archived?: number}`. All four members are Unix
// millisecond timestamps, so all four map to Go `int64` (matching
// [opencode.SessionTime], which models the identical schema).

func TestGlobalSessionTimeArchivedIsInt64Timestamp(t *testing.T) {
	// A real Unix millisecond timestamp: 2025-01-01T00:00:00Z == 1735689600000.
	// Stored as float64 this renders as 1.7356896e+12 and cannot be used as a
	// timestamp without a lossy conversion.
	const raw = `{"created":1735689600000,"updated":1735689600001,"compacting":1735689600002,"archived":1735689600003}`

	var got opencode.GlobalSessionTime
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}

	if got.Created != 1735689600000 {
		t.Errorf("Created: want 1735689600000, got %d", got.Created)
	}
	if got.Updated != 1735689600001 {
		t.Errorf("Updated: want 1735689600001, got %d", got.Updated)
	}
	if got.Compacting != 1735689600002 {
		t.Errorf("Compacting: want 1735689600002, got %d", got.Compacting)
	}
	if got.Archived != 1735689600003 {
		t.Errorf("Archived: want 1735689600003, got %d", got.Archived)
	}
	if got.JSON.RawJSON() != raw {
		t.Errorf("RawJSON: want %q, got %q", raw, got.JSON.RawJSON())
	}
}

func TestGlobalSessionTimeArchivedMaxSafeTimestamp(t *testing.T) {
	// 2^53-1 is the largest integer exactly representable as a float64. Any
	// timestamp beyond it silently loses precision when decoded as float64, so
	// this asserts the int64 field survives it byte-for-byte.
	const raw = `{"created":9007199254740993,"updated":9007199254740993,"archived":9007199254740993}`

	var got opencode.GlobalSessionTime
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.Archived != 9007199254740993 {
		t.Errorf("Archived: want 9007199254740993, got %d", got.Archived)
	}
}

func TestGlobalSessionTimeOptionalMembersAbsent(t *testing.T) {
	// OpenAPI requires only created + updated. compacting and archived must
	// decode to the zero value when omitted.
	var got opencode.GlobalSessionTime
	if err := json.Unmarshal([]byte(`{"created":1735689600000,"updated":1735689600001}`), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.Archived != 0 {
		t.Errorf("Archived: want 0 when absent, got %d", got.Archived)
	}
	if got.Compacting != 0 {
		t.Errorf("Compacting: want 0 when absent, got %d", got.Compacting)
	}
	if !got.JSON.Archived.IsMissing() {
		t.Error("JSON.Archived: want IsMissing() for an omitted field")
	}
	if !got.JSON.Compacting.IsMissing() {
		t.Error("JSON.Compacting: want IsMissing() for an omitted field")
	}
}

func TestGlobalSessionTimeArchivedZeroIsDistinguishable(t *testing.T) {
	// An explicit 0 must be distinguishable from an omitted field via the
	// JSON metadata, since the Go zero value is identical in both cases.
	var got opencode.GlobalSessionTime
	if err := json.Unmarshal([]byte(`{"created":1,"updated":2,"archived":0}`), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.Archived != 0 {
		t.Errorf("Archived: want 0, got %d", got.Archived)
	}
	if got.JSON.Archived.IsMissing() {
		t.Error("JSON.Archived: want present for an explicit 0")
	}
}

func TestGlobalSessionTimeMatchesSessionTimeShape(t *testing.T) {
	// GlobalSession.time and Session.time are the same OpenAPI schema, so the
	// two Go structs must decode identical payloads to identical values.
	const raw = `{"created":1735689600000,"updated":1735689600001,"compacting":1735689600002,"archived":1735689600003}`

	var global opencode.GlobalSessionTime
	if err := json.Unmarshal([]byte(raw), &global); err != nil {
		t.Fatalf("unmarshal GlobalSessionTime: %s", err)
	}
	var session opencode.SessionTime
	if err := json.Unmarshal([]byte(raw), &session); err != nil {
		t.Fatalf("unmarshal SessionTime: %s", err)
	}

	if global.Created != session.Created || global.Updated != session.Updated ||
		global.Compacting != session.Compacting || global.Archived != session.Archived {
		t.Errorf("GlobalSessionTime %+v diverges from SessionTime %+v for an identical schema",
			struct{ C, U, Cp, A int64 }{global.Created, global.Updated, global.Compacting, global.Archived},
			struct{ C, U, Cp, A int64 }{session.Created, session.Updated, session.Compacting, session.Archived})
	}
}

// ===== GlobalSession.Project =====
//
// OpenAPI: `"project": {"anyOf":[{"$ref":"ProjectSummary"},{"type":"null"}]}`
// and `project` IS listed in GlobalSession.required.
// JS SDK v2: `project: ProjectSummary | null`.

func TestGlobalSessionProjectObject(t *testing.T) {
	const raw = `{"id":"ses_1","slug":"slug","projectID":"prj_1","directory":"/d","title":"t","version":"1",` +
		`"time":{"created":1735689600000,"updated":1735689600001},` +
		`"project":{"id":"prj_1","worktree":"/w","name":"n"}}`

	var got opencode.GlobalSession
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.Project == nil {
		t.Fatal("Project: want non-nil for a JSON object")
	}
	if got.JSON.Project.IsMissing() {
		t.Error("JSON.Project: want present")
	}

	// Round-trip the carried value into ProjectSummary, which is the schema
	// type documented by the field's runtime-type comment.
	summaryRaw, err := json.Marshal(got.Project)
	if err != nil {
		t.Fatalf("marshal Project: %s", err)
	}
	var summary opencode.ProjectSummary
	if err := json.Unmarshal(summaryRaw, &summary); err != nil {
		t.Fatalf("unmarshal ProjectSummary: %s", err)
	}
	if summary.ID != "prj_1" || summary.Worktree != "/w" || summary.Name != "n" {
		t.Errorf("ProjectSummary: got %+v", summary)
	}
}

func TestGlobalSessionProjectNull(t *testing.T) {
	// The null variant of the anyOf must leave the any field nil rather than
	// panicking or producing a typed zero value.
	const raw = `{"id":"ses_1","slug":"slug","projectID":"prj_1","directory":"/d","title":"t","version":"1",` +
		`"time":{"created":1,"updated":2},"project":null}`

	var got opencode.GlobalSession
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.Project != nil {
		t.Errorf("Project: want nil for JSON null, got %#v", got.Project)
	}
	if got.ID != "ses_1" {
		t.Errorf("ID: want ses_1, got %q", got.ID)
	}
}

func TestGlobalSessionProjectAbsent(t *testing.T) {
	const raw = `{"id":"ses_1","slug":"slug","projectID":"prj_1","directory":"/d","title":"t","version":"1",` +
		`"time":{"created":1,"updated":2}}`

	var got opencode.GlobalSession
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.Project != nil {
		t.Errorf("Project: want nil when absent, got %#v", got.Project)
	}
	if !got.JSON.Project.IsMissing() {
		t.Error("JSON.Project: want IsMissing() when absent")
	}
}

// ===== Workspace.Branch / Directory / Extra / TimeUsed =====
//
// OpenAPI Workspace:
//
//	"branch":    {"anyOf":[{"type":"string"},{"type":"null"}]}
//	"directory": {"anyOf":[{"type":"string"},{"type":"null"}]}
//	"extra":     {"anyOf":[{},{"type":"null"}]}
//	"timeUsed":  {"anyOf":[{"type":"number"},{"type":"string","enum":["NaN"]},
//	                       {"type":"string","enum":["Infinity"]},
//	                       {"type":"string","enum":["-Infinity"]},
//	                       {"type":"string","enum":["Infinity","-Infinity","NaN"]}]}
//	"required":  ["id","type","name","projectID","timeUsed"]
//
// JS SDK v2:
//
//	branch?: string | null
//	directory?: string | null
//	extra?: unknown | null
//	timeUsed: number | "NaN" | "Infinity" | "-Infinity"

func TestWorkspaceNullableFieldsPresent(t *testing.T) {
	const raw = `{"id":"wrk_1","type":"git","name":"n","projectID":"prj_1",` +
		`"branch":"main","directory":"/d","extra":{"a":1},"timeUsed":1735689600000}`

	var got opencode.Workspace
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if branch, ok := got.Branch.(string); !ok || branch != "main" {
		t.Errorf("Branch: want string %q, got %T %#v", "main", got.Branch, got.Branch)
	}
	if dir, ok := got.Directory.(string); !ok || dir != "/d" {
		t.Errorf("Directory: want string %q, got %T %#v", "/d", got.Directory, got.Directory)
	}
	if got.Extra == nil {
		t.Error("Extra: want non-nil for a JSON object")
	}
	if got.ID != "wrk_1" || got.Type != "git" || got.Name != "n" || got.ProjectID != "prj_1" {
		t.Errorf("scalar fields: got %+v", struct{ ID, Type, Name, ProjectID string }{got.ID, got.Type, got.Name, got.ProjectID})
	}
}

func TestWorkspaceNullableFieldsNull(t *testing.T) {
	// The null variant of every nullable anyOf must decode to nil.
	const raw = `{"id":"wrk_1","type":"git","name":"n","projectID":"prj_1",` +
		`"branch":null,"directory":null,"extra":null,"timeUsed":0}`

	var got opencode.Workspace
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.Branch != nil {
		t.Errorf("Branch: want nil for JSON null, got %#v", got.Branch)
	}
	if got.Directory != nil {
		t.Errorf("Directory: want nil for JSON null, got %#v", got.Directory)
	}
	if got.Extra != nil {
		t.Errorf("Extra: want nil for JSON null, got %#v", got.Extra)
	}
}

func TestWorkspaceNullableFieldsAbsent(t *testing.T) {
	// branch, directory and extra are all absent from OpenAPI's required list.
	const raw = `{"id":"wrk_1","type":"git","name":"n","projectID":"prj_1","timeUsed":0}`

	var got opencode.Workspace
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.Branch != nil || got.Directory != nil || got.Extra != nil {
		t.Errorf("want all-nil when absent, got Branch=%#v Directory=%#v Extra=%#v",
			got.Branch, got.Directory, got.Extra)
	}
	if !got.JSON.Branch.IsMissing() || !got.JSON.Directory.IsMissing() || !got.JSON.Extra.IsMissing() {
		t.Error("JSON metadata: want IsMissing() for every omitted nullable field")
	}
}

func TestWorkspaceExtraUnconstrainedShapes(t *testing.T) {
	// OpenAPI models extra as anyOf[{}, null] — the non-null variant is the
	// unconstrained empty schema, so every JSON type must decode without error.
	for _, tc := range []struct {
		name string
		body string
		want any
	}{
		{"object", `{"k":"v"}`, map[string]any{"k": "v"}},
		{"array", `[1,2]`, []any{float64(1), float64(2)}},
		{"string", `"s"`, "s"},
		{"number", `1.5`, 1.5},
		{"bool", `true`, true},
		{"null", `null`, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			raw := `{"id":"wrk_1","type":"git","name":"n","projectID":"prj_1","timeUsed":0,"extra":` + tc.body + `}`
			var got opencode.Workspace
			if err := json.Unmarshal([]byte(raw), &got); err != nil {
				t.Fatalf("unmarshal: %s", err)
			}
			gotJSON, err := json.Marshal(got.Extra)
			if err != nil {
				t.Fatalf("marshal Extra: %s", err)
			}
			wantJSON, err := json.Marshal(tc.want)
			if err != nil {
				t.Fatalf("marshal want: %s", err)
			}
			if string(gotJSON) != string(wantJSON) {
				t.Errorf("Extra: want %s, got %s (%T)", wantJSON, gotJSON, got.Extra)
			}
		})
	}
}

func TestWorkspaceTimeUsedVariants(t *testing.T) {
	// timeUsed is required and its anyOf admits a number plus the three
	// non-finite sentinel strings.
	t.Run("number", func(t *testing.T) {
		var got opencode.Workspace
		raw := `{"id":"wrk_1","type":"git","name":"n","projectID":"prj_1","timeUsed":1735689600000}`
		if err := json.Unmarshal([]byte(raw), &got); err != nil {
			t.Fatalf("unmarshal: %s", err)
		}
		f, ok := got.TimeUsed.(float64)
		if !ok {
			t.Fatalf("TimeUsed: want float64, got %T %#v", got.TimeUsed, got.TimeUsed)
		}
		if f != 1735689600000 {
			t.Errorf("TimeUsed: want 1735689600000, got %v", f)
		}
	})

	t.Run("fractional number", func(t *testing.T) {
		var got opencode.Workspace
		raw := `{"id":"wrk_1","type":"git","name":"n","projectID":"prj_1","timeUsed":12.5}`
		if err := json.Unmarshal([]byte(raw), &got); err != nil {
			t.Fatalf("unmarshal: %s", err)
		}
		if f, ok := got.TimeUsed.(float64); !ok || f != 12.5 {
			t.Errorf("TimeUsed: want float64 12.5, got %T %#v", got.TimeUsed, got.TimeUsed)
		}
	})

	for _, sentinel := range []string{"NaN", "Infinity", "-Infinity"} {
		t.Run(sentinel, func(t *testing.T) {
			var got opencode.Workspace
			raw := `{"id":"wrk_1","type":"git","name":"n","projectID":"prj_1","timeUsed":"` + sentinel + `"}`
			if err := json.Unmarshal([]byte(raw), &got); err != nil {
				t.Fatalf("unmarshal: %s", err)
			}
			s, ok := got.TimeUsed.(string)
			if !ok {
				t.Fatalf("TimeUsed: want string, got %T %#v", got.TimeUsed, got.TimeUsed)
			}
			if s != sentinel {
				t.Errorf("TimeUsed: want %q, got %q", sentinel, s)
			}
		})
	}
}

func TestWorkspaceRawJSONPreservesUnknownFields(t *testing.T) {
	// Unknown fields must survive in RawJSON + ExtraFields so callers are not
	// blocked when the server ships a field ahead of the SDK.
	const raw = `{"id":"wrk_1","type":"git","name":"n","projectID":"prj_1","timeUsed":0,"futureField":"x"}`

	var got opencode.Workspace
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.JSON.RawJSON() != raw {
		t.Errorf("RawJSON: want %q, got %q", raw, got.JSON.RawJSON())
	}
	if _, ok := got.JSON.ExtraFields["futureField"]; !ok {
		t.Errorf("ExtraFields: want futureField to be captured, got %v", got.JSON.ExtraFields)
	}
}

func TestGlobalSessionRawJSONPreservesUnknownFields(t *testing.T) {
	const raw = `{"id":"ses_1","slug":"slug","projectID":"prj_1","directory":"/d","title":"t","version":"1",` +
		`"time":{"created":1,"updated":2},"project":null,"futureField":"x"}`

	var got opencode.GlobalSession
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.JSON.RawJSON() != raw {
		t.Errorf("RawJSON: want %q, got %q", raw, got.JSON.RawJSON())
	}
	if _, ok := got.JSON.ExtraFields["futureField"]; !ok {
		t.Errorf("ExtraFields: want futureField to be captured, got %v", got.JSON.ExtraFields)
	}
}

// ===== GlobalSession.Permission / Metadata =====
//
// OpenAPI: `"permission": {"$ref":"PermissionRuleset"}` where
// `PermissionRuleset` is `{"type":"array","items":{"$ref":"PermissionRule"}}`,
// and `"metadata": {"type":"object"}`. Neither is required. This mirrors
// [opencode.Session], whose Permission/Metadata carry the same schema.

func TestGlobalSessionPermissionAndMetadata(t *testing.T) {
	const raw = `{"id":"ses_1","slug":"slug","projectID":"prj_1","directory":"/d","title":"t","version":"1",` +
		`"time":{"created":1,"updated":2},"project":null,` +
		`"permission":[{"permission":"edit","pattern":"**","action":"allow"}],` +
		`"metadata":{"k":"v"}}`

	var got opencode.GlobalSession
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}

	// Permission is a concrete []PermissionRuleResponse: `PermissionRuleset` is a
	// plain array in OpenAPI, not an anyOf, so it needs no `any` carrier and is
	// readable without a marshal/re-parse round trip.
	if len(got.Permission) != 1 || got.Permission[0].Permission != "edit" ||
		got.Permission[0].Pattern != "**" ||
		got.Permission[0].Action != opencode.PermissionActionAllow {
		t.Errorf("Permission: got %+v", got.Permission)
	}

	metaRaw, err := json.Marshal(got.Metadata)
	if err != nil {
		t.Fatalf("marshal Metadata: %s", err)
	}
	var meta map[string]any
	if err := json.Unmarshal(metaRaw, &meta); err != nil {
		t.Fatalf("unmarshal map[string]any: %s", err)
	}
	if meta["k"] != "v" {
		t.Errorf("Metadata: want k=v, got %v", meta)
	}
}

func TestGlobalSessionPermissionAndMetadataAbsent(t *testing.T) {
	const raw = `{"id":"ses_1","slug":"slug","projectID":"prj_1","directory":"/d","title":"t","version":"1",` +
		`"time":{"created":1,"updated":2},"project":null}`

	var got opencode.GlobalSession
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %s", err)
	}
	if got.Permission != nil {
		t.Errorf("Permission: want nil when absent, got %#v", got.Permission)
	}
	if got.Metadata != nil {
		t.Errorf("Metadata: want nil when absent, got %#v", got.Metadata)
	}
}
