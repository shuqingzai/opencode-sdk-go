// Package opencode — empirical round-trip / serialization tests for the SDK
// changes introduced in this sync cycle.
//
// # Verification matrix
//
//  1. Request serialization (MarshalJSON)
//  2. Query serialization (URLQuery)
//  3. Response deserialization (UnmarshalJSON + RawJSON)
//  4. Round-trip consistency (marshal → unmarshal → re-marshal)
//  5. Required-tag semantics investigation in the Stainless framework

package opencode

import (
	"encoding/json"
	"strings"
	"testing"
)

// ────────────────────────────────────────────────────────────────────────────
// VERIFICATION 1 — Request serialization (MarshalJSON)
// ────────────────────────────────────────────────────────────────────────────

// TestV1_RequestSerialization_RequiredTagSemantics probes what the Stainless
// `param.Field[T]` + `required` JSON tag actually does at runtime.
//
// The critical question: does the `,required` tag affect whether a field is
// emitted in the serialized JSON when the param.Field is NOT set (Present=false)?
func TestV1_RequestSerialization_RequiredTagSemantics(t *testing.T) {
	// Sub-test A: required field unset (Present=false) — must NOT appear in JSON
	t.Run("required field unset → NOT emitted (Present wins over required tag)", func(t *testing.T) {
		// PartUpdatePartToolStatePending has Status, Input, Raw all required
		s := PartUpdatePartToolStatePending{}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		// All three fields are required by OpenAPI but Present=false → must be omitted
		for _, absent := range []string{`"status"`, `"input"`, `"raw"`} {
			if strings.Contains(got, absent) {
				t.Errorf("FAIL: required-but-unset field %q appeared in JSON: %s\n"+
					"→ This proves ,required tag does NOT force emission when Present=false", absent, got)
			}
		}
		// Serialized output should be an empty object
		if got != "{}" {
			t.Errorf("expected {}, got %s", got)
		}
		t.Logf("Confirmed: empty struct serializes to %s", got)
	})

	// Sub-test B: required field with zero value (Present=true, Value=zero) — MUST appear
	t.Run("required field set to zero value → IS emitted (Present=true)", func(t *testing.T) {
		s := PartUpdatePartToolStatePending{
			Status: F(PartUpdatePartToolStatePendingStatus("")), // zero string
			Input:  F(map[string]any{}),                         // zero map (empty)
			Raw:    F(""),                                       // zero string
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{`"status":`, `"input":`, `"raw":`} {
			if !strings.Contains(got, want) {
				t.Errorf("zero-value required field %q missing from JSON: %s", want, got)
			}
		}
		t.Logf("Confirmed: zero-value required fields serialize to %s", got)
	})

	// Sub-test C: optional field unset → must NOT appear
	t.Run("optional field unset → NOT emitted", func(t *testing.T) {
		s := PartUpdatePartToolStateCompleted{
			Input:    F(map[string]any{"k": "v"}),
			Metadata: F(map[string]any{"m": 1}),
			Output:   F("out"),
			Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
			Time: F(PartUpdatePartToolStateCompletedTime{
				Start: F(int64(1000)),
				End:   F(int64(2000)),
			}),
			Title: F("T"),
			// Attachments not set (optional)
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if strings.Contains(got, `"attachments"`) {
			t.Errorf("optional unset 'attachments' appeared in JSON: %s", got)
		}
		t.Logf("Confirmed: unset optional attachments absent from %s", got)
	})

	// Sub-test D: optional compacted unset → not in time object
	t.Run("optional compacted unset → NOT in time object", func(t *testing.T) {
		s := PartUpdatePartToolStateCompletedTime{
			Start: F(int64(100)),
			End:   F(int64(200)),
			// Compacted not set
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if strings.Contains(got, `"compacted"`) {
			t.Errorf("unset optional 'compacted' appeared in JSON: %s", got)
		}
		t.Logf("Confirmed: %s", got)
	})

	// Sub-test E: optional compacted SET → appears
	t.Run("optional compacted set → IS in time object", func(t *testing.T) {
		s := PartUpdatePartToolStateCompletedTime{
			Start:     F(int64(100)),
			End:       F(int64(200)),
			Compacted: F(int64(150)),
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"compacted":150`) {
			t.Errorf("set 'compacted' not in JSON: %s", got)
		}
		t.Logf("Confirmed: %s", got)
	})
}

// TestV1_RequestSerialization_ToolStatePending verifies Pending variant serialization.
func TestV1_RequestSerialization_ToolStatePending(t *testing.T) {
	t.Run("all fields set — full JSON", func(t *testing.T) {
		s := PartUpdatePartToolStatePending{
			Status: F(PartUpdatePartToolStatePendingStatusPending),
			Input:  F(map[string]any{"cmd": "ls", "dir": "/tmp"}),
			Raw:    F(`{"cmd":"ls","dir":"/tmp"}`),
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"status":"pending"`,
			`"input":`,
			`"raw":`,
			`"cmd"`,
			`"dir"`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
		t.Logf("Pending serialized: %s", got)
	})

	t.Run("whole tool part with pending state via union", func(t *testing.T) {
		tool := PartUpdatePartTool{
			CallID:    F("call_abc"),
			ID:        F("prt_tool"),
			MessageID: F("msg_1"),
			SessionID: F("ses_1"),
			Tool:      F("bash"),
			Type:      F(PartUpdatePartToolTypeTool),
			State: F(PartUpdatePartToolStateUnion(PartUpdatePartToolStatePending{
				Status: F(PartUpdatePartToolStatePendingStatusPending),
				Input:  F(map[string]any{"cmd": "echo hello"}),
				Raw:    F(`{"cmd":"echo hello"}`),
			})),
		}
		b, err := json.Marshal(tool)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"type":"tool"`,
			`"callID":"call_abc"`,
			`"tool":"bash"`,
			`"state":`,
			`"status":"pending"`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
		t.Logf("Tool with pending state: %s", got)
	})
}

// TestV1_RequestSerialization_ToolStateCompleted_Attachments verifies the new
// Attachments field on PartUpdatePartToolStateCompleted serializes correctly and
// that each FilePart element matches OpenAPI FilePart field names.
func TestV1_RequestSerialization_ToolStateCompleted_Attachments(t *testing.T) {
	s := PartUpdatePartToolStateCompleted{
		Input:    F(map[string]any{"x": 1}),
		Metadata: F(map[string]any{"provider": "test", "model": "gpt-4"}),
		Output:   F("analysis complete"),
		Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
		Time: F(PartUpdatePartToolStateCompletedTime{
			Start:     F(int64(1700000000)),
			End:       F(int64(1700000100)),
			Compacted: F(int64(1700000050)),
		}),
		Title: F("code_analysis"),
		Attachments: F([]PartUpdatePartFile{
			{
				ID:        F("fprt_1"),
				SessionID: F("ses_1"),
				MessageID: F("msg_1"),
				Mime:      F("image/png"),
				URL:       F("file:///chart.png"),
				Type:      F(PartUpdatePartFileTypeFile),
				Filename:  F("chart.png"),
			},
			{
				ID:        F("fprt_2"),
				SessionID: F("ses_1"),
				MessageID: F("msg_1"),
				Mime:      F("application/pdf"),
				URL:       F("file:///report.pdf"),
				Type:      F(PartUpdatePartFileTypeFile),
			},
		}),
	}

	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)

	// Verify attachments array present
	if !strings.Contains(got, `"attachments":[`) {
		t.Errorf("missing 'attachments' array: %s", got)
	}
	// Verify FilePart field names match OpenAPI (id, messageID, mime, sessionID, type, url)
	for _, want := range []string{
		`"id":"fprt_1"`,
		`"mime":"image/png"`,
		`"type":"file"`,
		`"url":"file:///chart.png"`,
		`"filename":"chart.png"`,
		`"id":"fprt_2"`,
		`"mime":"application/pdf"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in attachments: %s", want, got)
		}
	}
	// Verify compacted in time
	if !strings.Contains(got, `"compacted":1700000050`) {
		t.Errorf("missing 'compacted' in time: %s", got)
	}
	t.Logf("Completed with attachments: %s", got)
}

// TestV1_RequestSerialization_UnionInPartUpdateParams verifies putting ToolState
// variants into PartUpdatePartTool and then into PartUpdateParams produces correct
// top-level JSON.
func TestV1_RequestSerialization_UnionInPartUpdateParams(t *testing.T) {
	params := PartUpdateParams{
		Directory: F("d"),
		Workspace: F("w"),
		Part: F(PartUpdatePartUnion(PartUpdatePartTool{
			CallID:    F("cid"),
			ID:        F("pid"),
			MessageID: F("mid"),
			SessionID: F("sid"),
			Tool:      F("read_file"),
			Type:      F(PartUpdatePartToolTypeTool),
			State: F(PartUpdatePartToolStateUnion(PartUpdatePartToolStateCompleted{
				Input:    F(map[string]any{"path": "/etc/hosts"}),
				Metadata: F(map[string]any{}),
				Output:   F("127.0.0.1 localhost"),
				Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
				Time: F(PartUpdatePartToolStateCompletedTime{
					Start: F(int64(1000)),
					End:   F(int64(1010)),
				}),
				Title: F("read_file"),
			})),
		})),
	}

	b, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)

	// The Part is set, so body is the union variant (PartUpdatePartTool) at root
	for _, want := range []string{
		`"type":"tool"`,
		`"callID":"cid"`,
		`"state":`,
		`"status":"completed"`,
		`"output":"127.0.0.1 localhost"`,
		`"title":"read_file"`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in full params marshal: %s", want, got)
		}
	}
	// Directory and workspace are query params only — must NOT appear in body JSON
	if strings.Contains(got, `"directory"`) || strings.Contains(got, `"workspace"`) {
		t.Errorf("query params leaked into body JSON: %s", got)
	}
	t.Logf("Full PartUpdateParams body: %s", got)
}

// ────────────────────────────────────────────────────────────────────────────
// VERIFICATION 2 — Query serialization (URLQuery)
// ────────────────────────────────────────────────────────────────────────────

// TestV2_QuerySerialization_SessionListParams verifies the new Scope enum
// and all other SessionListParams fields serialize to correct query strings.
func TestV2_QuerySerialization_SessionListParams(t *testing.T) {
	t.Run("all fields set — correct query string", func(t *testing.T) {
		params := SessionListParams{
			Directory: F("d"),
			Workspace: F("w"),
			Scope:     F(SessionListParamsScopeProject),
		}
		v := params.URLQuery()
		if got := v.Get("scope"); got != "project" {
			t.Errorf("scope: got %q, want \"project\"", got)
		}
		if got := v.Get("directory"); got != "d" {
			t.Errorf("directory: got %q, want \"d\"", got)
		}
		if got := v.Get("workspace"); got != "w" {
			t.Errorf("workspace: got %q, want \"w\"", got)
		}
		t.Logf("Full query: %s", v.Encode())
	})

	t.Run("only scope set", func(t *testing.T) {
		params := SessionListParams{
			Scope: F(SessionListParamsScopeProject),
		}
		v := params.URLQuery()
		if got := v.Get("scope"); got != "project" {
			t.Errorf("scope: got %q, want \"project\"", got)
		}
		if v.Has("directory") || v.Has("workspace") {
			t.Errorf("unset fields appeared: %s", v.Encode())
		}
	})

	t.Run("scope unset — absent from query", func(t *testing.T) {
		params := SessionListParams{Directory: F("d")}
		v := params.URLQuery()
		if v.Has("scope") {
			t.Errorf("unset scope appeared: %s", v.Encode())
		}
	})

	t.Run("scope IsKnown round-trip", func(t *testing.T) {
		if !SessionListParamsScopeProject.IsKnown() {
			t.Error("SessionListParamsScopeProject.IsKnown() == false")
		}
		if SessionListParamsScope("garbage").IsKnown() {
			t.Error("unknown scope should not be known")
		}
	})
}

// TestV2_QuerySerialization_PermissionListParams spot-checks PermissionListParams.
func TestV2_QuerySerialization_PermissionListParams(t *testing.T) {
	t.Run("both fields set", func(t *testing.T) {
		params := PermissionListParams{
			Directory: F("mydir"),
			Workspace: F("myws"),
		}
		v := params.URLQuery()
		if got := v.Get("directory"); got != "mydir" {
			t.Errorf("got %q", got)
		}
		if got := v.Get("workspace"); got != "myws" {
			t.Errorf("got %q", got)
		}
	})

	t.Run("no fields set — empty query", func(t *testing.T) {
		v := PermissionListParams{}.URLQuery()
		if v.Encode() != "" {
			t.Errorf("expected empty query, got %s", v.Encode())
		}
	})
}

// ────────────────────────────────────────────────────────────────────────────
// VERIFICATION 3 — Response deserialization (UnmarshalJSON + RawJSON)
// ────────────────────────────────────────────────────────────────────────────

// TestV3_Response_PermissionRequest_MetadataIsAny verifies the Metadata=any
// change on PermissionRequest accepts all JSON shapes.
func TestV3_Response_PermissionRequest_MetadataIsAny(t *testing.T) {
	cases := []struct {
		name     string
		raw      string
		wantMeta any // nil means we just verify no error
	}{
		{
			name:     "metadata is empty object",
			raw:      `{"id":"r1","sessionID":"s1","permission":"bash","patterns":[],"always":[],"metadata":{}}`,
			wantMeta: map[string]any{},
		},
		{
			name:     "metadata is nested object",
			raw:      `{"id":"r1","sessionID":"s1","permission":"bash","patterns":[],"always":[],"metadata":{"tool":"bash","count":3}}`,
			wantMeta: map[string]any{"tool": "bash", "count": float64(3)},
		},
		{
			name:     "metadata is array",
			raw:      `{"id":"r1","sessionID":"s1","permission":"bash","patterns":[],"always":[],"metadata":["a","b"]}`,
			wantMeta: nil, // just verify no parse error
		},
		{
			name:     "metadata is null",
			raw:      `{"id":"r1","sessionID":"s1","permission":"bash","patterns":[],"always":[],"metadata":null}`,
			wantMeta: nil,
		},
		{
			name:     "metadata absent",
			raw:      `{"id":"r1","sessionID":"s1","permission":"bash","patterns":[],"always":[]}`,
			wantMeta: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var pr PermissionRequest
			if err := json.Unmarshal([]byte(tc.raw), &pr); err != nil {
				t.Fatalf("UnmarshalJSON failed: %v", err)
			}
			// Verify ID was parsed (basic sanity)
			if pr.ID != "r1" {
				t.Errorf("ID: got %q, want r1", pr.ID)
			}
			// Verify RawJSON preserves full original
			raw := pr.JSON.RawJSON()
			if raw != tc.raw {
				t.Errorf("RawJSON not preserved.\ngot:  %s\nwant: %s", raw, tc.raw)
			}
			t.Logf("%s → Metadata type: %T, value: %v", tc.name, pr.Metadata, pr.Metadata)
		})
	}
}

// TestV3_Response_PermissionRequest_ExtraFields verifies unknown fields don't
// cause errors and RawJSON preserves the full payload (forward compatibility).
func TestV3_Response_PermissionRequest_ExtraFields(t *testing.T) {
	raw := `{"id":"r99","sessionID":"s1","permission":"fs_read","patterns":["/tmp"],"always":["bash"],"metadata":{"k":"v"},"unknownFutureField":"some_value","anotherNew":{"nested":true}}`
	var pr PermissionRequest
	if err := json.Unmarshal([]byte(raw), &pr); err != nil {
		t.Fatalf("UnmarshalJSON with extra fields failed: %v", err)
	}
	if pr.ID != "r99" {
		t.Errorf("ID: got %q", pr.ID)
	}
	// RawJSON must be preserved exactly
	if got := pr.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON mismatch.\ngot:  %s\nwant: %s", got, raw)
	}
	t.Logf("Extra fields preserved in RawJSON: %s", pr.JSON.RawJSON())
}

// TestV3_Response_ToolStatePending verifies ToolStatePending deserializes
// correctly with Input and Raw fields (newly added per sync).
func TestV3_Response_ToolStatePending(t *testing.T) {
	raw := `{"status":"pending","input":{"cmd":"ls","dir":"/tmp"},"raw":"{\"cmd\":\"ls\",\"dir\":\"/tmp\"}"}`
	var ts ToolStatePending
	if err := json.Unmarshal([]byte(raw), &ts); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if ts.Status != ToolStatePendingStatusPending {
		t.Errorf("Status: got %q, want %q", ts.Status, ToolStatePendingStatusPending)
	}
	if ts.Input == nil {
		t.Error("Input should not be nil")
	} else if ts.Input["cmd"] != "ls" {
		t.Errorf("Input.cmd: got %v", ts.Input["cmd"])
	}
	if !strings.Contains(ts.Raw, "cmd") {
		t.Errorf("Raw field not populated: %q", ts.Raw)
	}
	// RawJSON preserves original
	if got := ts.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON mismatch: got %s", got)
	}
	t.Logf("ToolStatePending deserialized: status=%s input=%v raw=%s", ts.Status, ts.Input, ts.Raw)
}

// TestV3_Response_ToolStateCompleted_WithAttachmentsAndCompacted verifies
// ToolStateCompleted with the new Attachments and Compacted time field.
func TestV3_Response_ToolStateCompleted_WithAttachmentsAndCompacted(t *testing.T) {
	raw := `{
		"status":"completed",
		"input":{"path":"/etc/hosts"},
		"metadata":{"provider":"openai","model":"gpt-4"},
		"output":"127.0.0.1 localhost",
		"title":"read_file",
		"time":{"start":1700000000,"end":1700000010,"compacted":1700000005},
		"attachments":[
			{"id":"fp1","messageID":"m1","mime":"text/plain","sessionID":"s1","type":"file","url":"file:///out.txt","filename":"out.txt"}
		]
	}`
	var ts ToolStateCompleted
	if err := json.Unmarshal([]byte(raw), &ts); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if ts.Status != ToolStateCompletedStatusCompleted {
		t.Errorf("Status: got %q", ts.Status)
	}
	if ts.Output != "127.0.0.1 localhost" {
		t.Errorf("Output: got %q", ts.Output)
	}
	if ts.Title != "read_file" {
		t.Errorf("Title: got %q", ts.Title)
	}
	if ts.Time.Start != 1700000000 {
		t.Errorf("Time.Start: got %d", ts.Time.Start)
	}
	if ts.Time.End != 1700000010 {
		t.Errorf("Time.End: got %d", ts.Time.End)
	}
	// Verify new Compacted field
	if ts.Time.Compacted != 1700000005 {
		t.Errorf("Time.Compacted: got %d, want 1700000005", ts.Time.Compacted)
	}
	// Verify Attachments (new field)
	if len(ts.Attachments) != 1 {
		t.Errorf("Attachments length: got %d, want 1", len(ts.Attachments))
	} else {
		att := ts.Attachments[0]
		if att.ID != "fp1" {
			t.Errorf("Attachment.ID: got %q, want fp1", att.ID)
		}
		if att.Mime != "text/plain" {
			t.Errorf("Attachment.Mime: got %q", att.Mime)
		}
		if att.Filename != "out.txt" {
			t.Errorf("Attachment.Filename: got %q", att.Filename)
		}
		if att.Type != FilePartTypeFile {
			t.Errorf("Attachment.Type: got %q", att.Type)
		}
	}
	t.Logf("ToolStateCompleted deserialized: status=%s output=%s compacted=%d attachments=%d",
		ts.Status, ts.Output, ts.Time.Compacted, len(ts.Attachments))
}

// TestV3_Response_ToolPartState_UnionDispatch verifies ToolPartState correctly
// dispatches to the right concrete type for all 4 status variants.
func TestV3_Response_ToolPartState_UnionDispatch(t *testing.T) {
	cases := []struct {
		name       string
		raw        string
		wantStatus ToolPartStateStatus
		wantType   string // type assertion target name
	}{
		{
			name:       "pending",
			raw:        `{"status":"pending","input":{"cmd":"ls"},"raw":"{\"cmd\":\"ls\"}"}`,
			wantStatus: ToolPartStateStatusPending,
			wantType:   "ToolStatePending",
		},
		{
			name:       "running",
			raw:        `{"status":"running","input":{"cmd":"ls"},"time":{"start":1000}}`,
			wantStatus: ToolPartStateStatusRunning,
			wantType:   "ToolStateRunning",
		},
		{
			name:       "completed",
			raw:        `{"status":"completed","input":{},"metadata":{},"output":"ok","title":"T","time":{"start":1000,"end":2000}}`,
			wantStatus: ToolPartStateStatusCompleted,
			wantType:   "ToolStateCompleted",
		},
		{
			name:       "error",
			raw:        `{"status":"error","input":{},"error":"something failed","time":{"start":1000,"end":1001}}`,
			wantStatus: ToolPartStateStatusError,
			wantType:   "ToolStateError",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var state ToolPartState
			if err := json.Unmarshal([]byte(tc.raw), &state); err != nil {
				t.Fatalf("UnmarshalJSON failed: %v", err)
			}
			if state.Status != tc.wantStatus {
				t.Errorf("Status: got %q, want %q", state.Status, tc.wantStatus)
			}
			// Verify AsUnion() returns the right concrete type
			union := state.AsUnion()
			if union == nil {
				t.Fatal("AsUnion() returned nil")
			}
			switch tc.wantType {
			case "ToolStatePending":
				if _, ok := union.(ToolStatePending); !ok {
					t.Errorf("AsUnion() type: got %T, want ToolStatePending", union)
				}
			case "ToolStateRunning":
				if _, ok := union.(ToolStateRunning); !ok {
					t.Errorf("AsUnion() type: got %T, want ToolStateRunning", union)
				}
			case "ToolStateCompleted":
				if _, ok := union.(ToolStateCompleted); !ok {
					t.Errorf("AsUnion() type: got %T, want ToolStateCompleted", union)
				}
			case "ToolStateError":
				if _, ok := union.(ToolStateError); !ok {
					t.Errorf("AsUnion() type: got %T, want ToolStateError", union)
				}
			}
			// RawJSON preserves full payload
			rawGot := state.JSON.RawJSON()
			if rawGot != tc.raw {
				t.Errorf("RawJSON mismatch.\ngot:  %s\nwant: %s", rawGot, tc.raw)
			}
			t.Logf("%s: union type=%T status=%s", tc.name, union, state.Status)
		})
	}
}

// TestV3_Response_ToolPartState_ExtraFields verifies forward compatibility:
// unknown fields in ToolPartState payloads are accepted without error and
// RawJSON preserves the full payload.
func TestV3_Response_ToolPartState_ExtraFields(t *testing.T) {
	raw := `{"status":"completed","input":{},"metadata":{},"output":"ok","title":"T","time":{"start":1,"end":2},"futureField":"foo","anotherFuture":{"x":1}}`
	var state ToolPartState
	if err := json.Unmarshal([]byte(raw), &state); err != nil {
		t.Fatalf("UnmarshalJSON with extra fields failed: %v", err)
	}
	if state.Status != ToolPartStateStatusCompleted {
		t.Errorf("Status: got %q", state.Status)
	}
	if got := state.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON not preserved.\ngot:  %s\nwant: %s", got, raw)
	}
	t.Logf("Extra fields handled, RawJSON preserved")
}

// TestV3_Response_ToolStatePending_FieldValuesFromAsUnion verifies that after
// ToolPartState.UnmarshalJSON and AsUnion(), the concrete fields are accessible.
func TestV3_Response_ToolStatePending_FieldValuesFromAsUnion(t *testing.T) {
	raw := `{"status":"pending","input":{"file":"/etc/hosts","mode":"read"},"raw":"{\"file\":\"/etc/hosts\"}"}`
	var state ToolPartState
	if err := json.Unmarshal([]byte(raw), &state); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	pending, ok := state.AsUnion().(ToolStatePending)
	if !ok {
		t.Fatalf("AsUnion() did not return ToolStatePending, got %T", state.AsUnion())
	}
	if pending.Status != ToolStatePendingStatusPending {
		t.Errorf("pending.Status: got %q", pending.Status)
	}
	if pending.Input["file"] != "/etc/hosts" {
		t.Errorf("pending.Input[file]: got %v", pending.Input["file"])
	}
	if !strings.Contains(pending.Raw, "etc/hosts") {
		t.Errorf("pending.Raw: got %q", pending.Raw)
	}
	t.Logf("ToolStatePending via AsUnion: status=%s input=%v raw=%s",
		pending.Status, pending.Input, pending.Raw)
}

// TestV3_Response_ToolStateCompleted_FieldValuesFromAsUnion verifies completed
// variant fields are accessible after AsUnion().
func TestV3_Response_ToolStateCompleted_FieldValuesFromAsUnion(t *testing.T) {
	raw := `{"status":"completed","input":{"path":"/etc/passwd"},"metadata":{"model":"gpt-4"},"output":"root:x:0:0","title":"read_file","time":{"start":1700000000,"end":1700000010,"compacted":1700000005},"attachments":[]}`
	var state ToolPartState
	if err := json.Unmarshal([]byte(raw), &state); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	completed, ok := state.AsUnion().(ToolStateCompleted)
	if !ok {
		t.Fatalf("AsUnion() did not return ToolStateCompleted, got %T", state.AsUnion())
	}
	if completed.Output != "root:x:0:0" {
		t.Errorf("completed.Output: got %q", completed.Output)
	}
	if completed.Title != "read_file" {
		t.Errorf("completed.Title: got %q", completed.Title)
	}
	if completed.Time.Compacted != 1700000005 {
		t.Errorf("completed.Time.Compacted: got %d", completed.Time.Compacted)
	}
	t.Logf("ToolStateCompleted via AsUnion: output=%s compacted=%d", completed.Output, completed.Time.Compacted)
}

// ────────────────────────────────────────────────────────────────────────────
// VERIFICATION 4 — Round-trip consistency (marshal → unmarshal → re-marshal)
// ────────────────────────────────────────────────────────────────────────────

// TestV4_RoundTrip_ToolStatePending performs request-side marshal, then
// response-side unmarshal, then re-marshal and verifies data stability.
//
// Note: Request structs (param.Field) produce sorted JSON keys and omit
// unset fields. Response structs (standard encoding/json) preserve declaration
// order. These are intentionally different — the test compares semantic data
// rather than byte-for-byte strings.
func TestV4_RoundTrip_ToolStatePending(t *testing.T) {
	// Step 1: marshal the request-side param struct
	req := PartUpdatePartToolStatePending{
		Status: F(PartUpdatePartToolStatePendingStatusPending),
		Input:  F(map[string]any{"cmd": "cat /etc/hosts"}),
		Raw:    F(`{"cmd":"cat /etc/hosts"}`),
	}
	b1, err := json.Marshal(req)
	if err != nil {
		t.Fatal("marshal step 1:", err)
	}
	t.Logf("Step 1 (request marshal): %s", b1)

	// Step 2: unmarshal into response-side struct
	var resp ToolStatePending
	if err := json.Unmarshal(b1, &resp); err != nil {
		t.Fatal("unmarshal step 2:", err)
	}

	// Step 3: verify semantic data is preserved (not byte-for-byte order)
	if resp.Status != ToolStatePendingStatusPending {
		t.Errorf("Status not preserved: got %q", resp.Status)
	}
	if resp.Input["cmd"] != "cat /etc/hosts" {
		t.Errorf("Input.cmd not preserved: got %v", resp.Input["cmd"])
	}
	if !strings.Contains(resp.Raw, "cat /etc/hosts") {
		t.Errorf("Raw not preserved: got %q", resp.Raw)
	}

	// Step 4: re-marshal the response-side struct into same response-side struct
	// and verify stability (response → unmarshal → re-marshal is idempotent)
	b2, err := json.Marshal(resp)
	if err != nil {
		t.Fatal("marshal step 4:", err)
	}
	t.Logf("Step 4 (response marshal): %s", b2)

	var resp2 ToolStatePending
	if err := json.Unmarshal(b2, &resp2); err != nil {
		t.Fatal("unmarshal step 4b:", err)
	}
	b3, err := json.Marshal(resp2)
	if err != nil {
		t.Fatal("marshal step 4c:", err)
	}
	// Response → unmarshal → re-marshal must be byte-for-byte identical (idempotent)
	if string(b2) != string(b3) {
		t.Errorf("response re-marshal not idempotent:\nb2: %s\nb3: %s", b2, b3)
	}
	t.Logf("Response round-trip idempotent: %s", b2)
}

// TestV4_RoundTrip_ToolStateCompleted performs a full round-trip for the
// completed variant including Attachments and Compacted fields.
// Tests semantic round-trip: marshal request params → unmarshal response → verify data intact.
// Also tests response-side idempotency: unmarshal → re-marshal → unmarshal gives same data.
func TestV4_RoundTrip_ToolStateCompleted(t *testing.T) {
	req := PartUpdatePartToolStateCompleted{
		Input:    F(map[string]any{"p": "/etc/hosts"}),
		Metadata: F(map[string]any{"m": "v"}),
		Output:   F("127.0.0.1 localhost"),
		Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
		Time: F(PartUpdatePartToolStateCompletedTime{
			Start:     F(int64(1700000000)),
			End:       F(int64(1700000010)),
			Compacted: F(int64(1700000005)),
		}),
		Title: F("read_file"),
		Attachments: F([]PartUpdatePartFile{
			{
				ID:        F("fp1"),
				SessionID: F("s1"),
				MessageID: F("m1"),
				Mime:      F("text/plain"),
				URL:       F("file:///out.txt"),
				Type:      F(PartUpdatePartFileTypeFile),
			},
		}),
	}

	b1, err := json.Marshal(req)
	if err != nil {
		t.Fatal("marshal step 1:", err)
	}
	t.Logf("Step 1 (request marshal): %s", b1)

	var resp ToolStateCompleted
	if err := json.Unmarshal(b1, &resp); err != nil {
		t.Fatal("unmarshal step 2:", err)
	}

	// Verify key semantic data is preserved across the request→response boundary
	if resp.Status != ToolStateCompletedStatusCompleted {
		t.Errorf("Status not preserved: got %q", resp.Status)
	}
	if resp.Output != "127.0.0.1 localhost" {
		t.Errorf("Output not preserved: got %q", resp.Output)
	}
	if resp.Title != "read_file" {
		t.Errorf("Title not preserved: got %q", resp.Title)
	}
	if resp.Time.Start != 1700000000 {
		t.Errorf("Time.Start not preserved: got %d", resp.Time.Start)
	}
	if resp.Time.End != 1700000010 {
		t.Errorf("Time.End not preserved: got %d", resp.Time.End)
	}
	if resp.Time.Compacted != 1700000005 {
		t.Errorf("Time.Compacted not preserved: got %d", resp.Time.Compacted)
	}
	if len(resp.Attachments) != 1 {
		t.Errorf("Attachments length not preserved: got %d", len(resp.Attachments))
	} else if resp.Attachments[0].ID != "fp1" {
		t.Errorf("Attachment ID not preserved: got %q", resp.Attachments[0].ID)
	}

	// Response-side idempotency: re-marshal and unmarshal must give same data
	b2, err := json.Marshal(resp)
	if err != nil {
		t.Fatal("marshal step 3:", err)
	}
	t.Logf("Step 3 (response marshal): %s", b2)

	var resp2 ToolStateCompleted
	if err := json.Unmarshal(b2, &resp2); err != nil {
		t.Fatal("unmarshal step 4:", err)
	}
	b3, err := json.Marshal(resp2)
	if err != nil {
		t.Fatal("marshal step 5:", err)
	}
	// Response → unmarshal → re-marshal is idempotent
	if string(b2) != string(b3) {
		t.Errorf("response re-marshal not idempotent:\nb2: %s\nb3: %s", b2, b3)
	}
	t.Logf("Response round-trip idempotent: %s", b2)
}

// TestV4_RoundTrip_PermissionRequest_Metadata verifies that PermissionRequest
// semantic data is preserved through unmarshal → re-marshal cycle.
// Note: re-marshal emits zero-value Tool struct field (Go standard encoding/json behavior)
// which is expected for non-pointer struct fields not in the original JSON.
func TestV4_RoundTrip_PermissionRequest_Metadata(t *testing.T) {
	cases := []struct {
		name string
		raw  string
	}{
		{
			name: "metadata is map",
			raw:  `{"id":"r1","sessionID":"s1","permission":"bash","patterns":["*"],"always":[],"metadata":{"k":"v","n":1}}`,
		},
		{
			name: "metadata is empty object",
			raw:  `{"id":"r1","sessionID":"s1","permission":"bash","patterns":[],"always":[],"metadata":{}}`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var pr PermissionRequest
			if err := json.Unmarshal([]byte(tc.raw), &pr); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			b2, err := json.Marshal(pr)
			if err != nil {
				t.Fatalf("re-marshal: %v", err)
			}
			// Verify key fields are preserved (not byte-for-byte — re-marshal adds zero Tool)
			var remarshal map[string]any
			if err := json.Unmarshal(b2, &remarshal); err != nil {
				t.Fatalf("re-parse: %v", err)
			}
			if remarshal["id"] != pr.ID {
				t.Errorf("id mismatch: got %v", remarshal["id"])
			}
			if remarshal["permission"] != pr.Permission {
				t.Errorf("permission mismatch: got %v", remarshal["permission"])
			}
			// Idempotency: re-marshal the already-re-marshalled struct gives same bytes
			var pr2 PermissionRequest
			if err := json.Unmarshal(b2, &pr2); err != nil {
				t.Fatalf("re-unmarshal: %v", err)
			}
			b3, err := json.Marshal(pr2)
			if err != nil {
				t.Fatalf("re-re-marshal: %v", err)
			}
			if string(b2) != string(b3) {
				t.Errorf("response not idempotent:\nb2: %s\nb3: %s", b2, b3)
			}
			t.Logf("%s: semantic fields preserved; response marshal idempotent", tc.name)
		})
	}
}

// TestV4_RoundTrip_ToolPartState_AllVariants verifies idempotent re-marshaling
// within the same Response struct family for all 4 ToolPartState union variants.
//
// Design note: ToolPartState is an intermediary union-bearing struct. Its
// standard encoding/json re-marshal emits zero-value fields (unlike Request
// param.Field which omits absent fields). This is expected SDK behavior.
// The meaningful round-trip test is: unmarshal JSON → AsUnion() → concrete type
// → re-marshal concrete type → unchanged.
func TestV4_RoundTrip_ToolPartState_AllVariants(t *testing.T) {
	t.Run("pending — concrete type round-trip", func(t *testing.T) {
		raw := `{"status":"pending","input":{"x":1},"raw":"{\"x\":1}"}`
		var state ToolPartState
		if err := json.Unmarshal([]byte(raw), &state); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		pending := state.AsUnion().(ToolStatePending)

		// Re-marshal the concrete type
		b2, err := json.Marshal(pending)
		if err != nil {
			t.Fatalf("re-marshal: %v", err)
		}
		var pending2 ToolStatePending
		if err := json.Unmarshal(b2, &pending2); err != nil {
			t.Fatalf("re-unmarshal: %v", err)
		}
		b3, err := json.Marshal(pending2)
		if err != nil {
			t.Fatalf("re-re-marshal: %v", err)
		}
		if string(b2) != string(b3) {
			t.Errorf("pending concrete not idempotent:\nb2: %s\nb3: %s", b2, b3)
		}
		// Verify data
		if pending2.Status != ToolStatePendingStatusPending {
			t.Errorf("Status: got %q", pending2.Status)
		}
		t.Logf("pending round-trip: %s", b2)
	})

	t.Run("running — concrete type round-trip", func(t *testing.T) {
		raw := `{"status":"running","input":{"x":1},"time":{"start":1000}}`
		var state ToolPartState
		if err := json.Unmarshal([]byte(raw), &state); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		running := state.AsUnion().(ToolStateRunning)
		b2, _ := json.Marshal(running)
		var running2 ToolStateRunning
		json.Unmarshal(b2, &running2)
		b3, _ := json.Marshal(running2)
		if string(b2) != string(b3) {
			t.Errorf("running concrete not idempotent:\nb2: %s\nb3: %s", b2, b3)
		}
		if running2.Time.Start != 1000 {
			t.Errorf("Time.Start: got %d", running2.Time.Start)
		}
		t.Logf("running round-trip: %s", b2)
	})

	t.Run("completed — concrete type round-trip with compacted", func(t *testing.T) {
		raw := `{"status":"completed","input":{"x":1},"metadata":{"m":"v"},"output":"ok","title":"T","time":{"start":1000,"end":2000,"compacted":1500}}`
		var state ToolPartState
		if err := json.Unmarshal([]byte(raw), &state); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		completed := state.AsUnion().(ToolStateCompleted)
		b2, _ := json.Marshal(completed)
		var completed2 ToolStateCompleted
		json.Unmarshal(b2, &completed2)
		b3, _ := json.Marshal(completed2)
		if string(b2) != string(b3) {
			t.Errorf("completed concrete not idempotent:\nb2: %s\nb3: %s", b2, b3)
		}
		if completed2.Time.Compacted != 1500 {
			t.Errorf("Time.Compacted: got %d", completed2.Time.Compacted)
		}
		t.Logf("completed round-trip: %s", b2)
	})

	t.Run("error — concrete type round-trip", func(t *testing.T) {
		raw := `{"status":"error","input":{"x":1},"error":"fail","time":{"start":1000,"end":1001}}`
		var state ToolPartState
		if err := json.Unmarshal([]byte(raw), &state); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		errVariant := state.AsUnion().(ToolStateError)
		b2, _ := json.Marshal(errVariant)
		var errVariant2 ToolStateError
		json.Unmarshal(b2, &errVariant2)
		b3, _ := json.Marshal(errVariant2)
		if string(b2) != string(b3) {
			t.Errorf("error concrete not idempotent:\nb2: %s\nb3: %s", b2, b3)
		}
		if errVariant2.Error != "fail" {
			t.Errorf("Error: got %q", errVariant2.Error)
		}
		t.Logf("error round-trip: %s", b2)
	})
}

// ────────────────────────────────────────────────────────────────────────────
// VERIFICATION 5 — Required-tag semantics summary test
//
// This test explicitly documents the Stainless framework's behavior with
// respect to the ,required JSON tag, providing evidence for the review.
// ────────────────────────────────────────────────────────────────────────────

// TestV5_RequiredTagSemantics_Summary summarizes the empirically observed
// behavior of the `,required` JSON struct tag in this Stainless-generated SDK.
func TestV5_RequiredTagSemantics_Summary(t *testing.T) {
	t.Run("FINDINGS: required tag has NO effect on serialization output", func(t *testing.T) {
		// Evidence 1: The tag.go parser stores `required bool` but encoder.go
		// never reads `tag.required` in newFieldTypeEncoder or newStructTypeEncoder.
		// The only check is `present.Bool()` on the param.Field.Present field.

		// Evidence 2: A required param.Field that is NOT set (Present=false)
		// produces NO JSON output — just like an optional unset field.
		pending := PartUpdatePartToolStatePending{
			// All three fields (status, input, raw) are tagged required
			// but none are set — Present=false for all
		}
		b, _ := json.Marshal(pending)
		if string(b) != "{}" {
			t.Errorf("expected {}, got %s — breaks assumption", string(b))
		}

		// Evidence 3: An OPTIONAL param.Field that IS set (Present=true)
		// DOES appear in JSON — just like a required set field.
		running := PartUpdatePartToolStateRunning{
			Input:  F(map[string]any{}),
			Status: F(PartUpdatePartToolStateRunningStatusRunning),
			Time:   F(PartUpdatePartToolStateRunningTime{Start: F(int64(1))}),
			Title:  F("optional_title"), // Title is OPTIONAL but set
		}
		b2, _ := json.Marshal(running)
		if !strings.Contains(string(b2), `"title":"optional_title"`) {
			t.Errorf("optional set field missing from JSON: %s", b2)
		}

		t.Logf("CONFIRMED: ,required tag in param.Field structs is purely documentation.")
		t.Logf("Serialization behavior is governed SOLELY by param.Field.Present (bool).")
		t.Logf("The `,required` tag correctly marks fields that OpenAPI requires clients to send,")
		t.Logf("but it does NOT enforce this at runtime — callers are responsible for setting them.")
	})

	t.Run("FINDINGS: required tag in Response structs also has no decoder effect", func(t *testing.T) {
		// The decoder's newStructTypeDecoder does not check tag.required.
		// A JSON payload missing a `required` field simply leaves the Go field
		// at its zero value — no error is returned.
		raw := `{"status":"completed","time":{"start":1,"end":2}}`
		// ToolStateCompleted has input/metadata/output/title all required
		// but they're absent from this payload
		var ts ToolStateCompleted
		if err := json.Unmarshal([]byte(raw), &ts); err != nil {
			t.Fatalf("UnmarshalJSON unexpectedly failed: %v", err)
		}
		// Fields not present in JSON get zero values
		if ts.Output != "" {
			t.Errorf("expected empty output, got %q", ts.Output)
		}
		t.Logf("CONFIRMED: Missing required fields in response JSON yield zero values, no error.")
		t.Logf("The ,required tag on Response structs is documentation only.")
	})
}

// ────────────────────────────────────────────────────────────────────────────
// VERIFICATION 6 — Deep audit: Union dispatch, JSON metadata, ExtraFields,
//                  UserMessage.Format deserialization
//
// Covers audit requirements from the second-round review of session.go:
//   - Scope enum (SessionListParamsScope) serialization
//   - Union discriminator dispatch (SessionStatus, MessageUnion, PartUnion)
//   - JSON metadata ExtraFields forward-compat (unknown fields preserved)
//   - UserMessage.Format any-field deserialization
//   - AssistantMessage.Structured any-field deserialization
// ────────────────────────────────────────────────────────────────────────────

// TestV6_UnionDispatch_MessageRole verifies MessageUnion correctly dispatches
// to UserMessage vs AssistantMessage based on the role discriminator.
func TestV6_UnionDispatch_MessageRole(t *testing.T) {
	t.Run("user message dispatches to UserMessage", func(t *testing.T) {
		raw := `{
			"id": "msg_1",
			"sessionID": "ses_1",
			"role": "user",
			"time": {"created": 1700000000},
			"agent": "default",
			"model": {"providerID": "anthropic", "modelID": "claude-3-5"}
		}`
		var m Message
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if m.Role != MessageRoleUser {
			t.Errorf("Role: got %q, want %q", m.Role, MessageRoleUser)
		}
		um, ok := m.AsUnion().(UserMessage)
		if !ok {
			t.Fatalf("expected UserMessage, got %T", m.AsUnion())
		}
		if um.ID != "msg_1" {
			t.Errorf("UserMessage.ID: got %q", um.ID)
		}
		if um.Agent != "default" {
			t.Errorf("UserMessage.Agent: got %q", um.Agent)
		}
		t.Logf("user message dispatched: %+v", um.Role)
	})

	t.Run("assistant message dispatches to AssistantMessage", func(t *testing.T) {
		raw := `{
			"id": "msg_2",
			"sessionID": "ses_1",
			"role": "assistant",
			"parentID": "msg_0",
			"modelID": "claude-3-5",
			"providerID": "anthropic",
			"mode": "default",
			"agent": "default",
			"path": {"cwd": "/tmp", "root": "/"},
			"cost": 0.002,
			"time": {"created": 1700000001},
			"tokens": {"input": 100, "output": 50, "reasoning": 0, "cache": {"read": 0, "write": 0}}
		}`
		var m Message
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		am, ok := m.AsUnion().(AssistantMessage)
		if !ok {
			t.Fatalf("expected AssistantMessage, got %T", m.AsUnion())
		}
		if am.ModelID != "claude-3-5" {
			t.Errorf("AssistantMessage.ModelID: got %q", am.ModelID)
		}
		if am.Cost != 0.002 {
			t.Errorf("AssistantMessage.Cost: got %v", am.Cost)
		}
		t.Logf("assistant message dispatched: %+v", am.Role)
	})
}

// TestV6_UnionDispatch_PartType verifies PartUnion dispatches to the correct
// Part variant for a representative selection of the 12 registered types.
func TestV6_UnionDispatch_PartType(t *testing.T) {
	cases := []struct {
		name     string
		raw      string
		wantType string
		verify   func(t *testing.T, p Part)
	}{
		{
			name:     "text part",
			raw:      `{"id":"p1","sessionID":"s1","messageID":"m1","type":"text","text":"hello"}`,
			wantType: "TextPart",
			verify: func(t *testing.T, p Part) {
				tp, ok := p.AsUnion().(TextPart)
				if !ok {
					t.Fatalf("expected TextPart, got %T", p.AsUnion())
				}
				if tp.Text != "hello" {
					t.Errorf("TextPart.Text: got %q", tp.Text)
				}
			},
		},
		{
			name:     "tool part",
			raw:      `{"id":"p2","sessionID":"s1","messageID":"m1","type":"tool","callID":"c1","tool":"bash","state":{"status":"pending","input":{},"raw":""}}`,
			wantType: "ToolPart",
			verify: func(t *testing.T, p Part) {
				tp, ok := p.AsUnion().(ToolPart)
				if !ok {
					t.Fatalf("expected ToolPart, got %T", p.AsUnion())
				}
				if tp.Tool != "bash" {
					t.Errorf("ToolPart.Tool: got %q", tp.Tool)
				}
				if tp.State.Status != ToolPartStateStatusPending {
					t.Errorf("ToolPart.State.Status: got %q", tp.State.Status)
				}
			},
		},
		{
			name:     "compaction part",
			raw:      `{"id":"p3","sessionID":"s1","messageID":"m1","type":"compaction","auto":true,"overflow":false}`,
			wantType: "CompactionPart",
			verify: func(t *testing.T, p Part) {
				cp, ok := p.AsUnion().(CompactionPart)
				if !ok {
					t.Fatalf("expected CompactionPart, got %T", p.AsUnion())
				}
				if !cp.Auto {
					t.Errorf("CompactionPart.Auto: expected true")
				}
			},
		},
		{
			name:     "step-finish part",
			raw:      `{"id":"p4","sessionID":"s1","messageID":"m1","type":"step-finish","reason":"stop","cost":0.01,"tokens":{"input":100,"output":50,"reasoning":0,"cache":{"read":0,"write":0}}}`,
			wantType: "StepFinishPart",
			verify: func(t *testing.T, p Part) {
				sfp, ok := p.AsUnion().(StepFinishPart)
				if !ok {
					t.Fatalf("expected StepFinishPart, got %T", p.AsUnion())
				}
				if sfp.Reason != "stop" {
					t.Errorf("StepFinishPart.Reason: got %q", sfp.Reason)
				}
			},
		},
		{
			name:     "subtask part",
			raw:      `{"id":"p5","sessionID":"s1","messageID":"m1","type":"subtask","prompt":"do work","description":"task","agent":"worker"}`,
			wantType: "SubtaskPart",
			verify: func(t *testing.T, p Part) {
				sp, ok := p.AsUnion().(SubtaskPart)
				if !ok {
					t.Fatalf("expected SubtaskPart, got %T", p.AsUnion())
				}
				if sp.Prompt != "do work" {
					t.Errorf("SubtaskPart.Prompt: got %q", sp.Prompt)
				}
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var p Part
			if err := json.Unmarshal([]byte(tc.raw), &p); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			tc.verify(t, p)
			t.Logf("%s dispatched to correct variant", tc.wantType)
		})
	}
}

// TestV6_ExtraFields_ForwardCompat verifies that unknown JSON fields are
// preserved in ExtraFields and retrievable via RawJSON(), ensuring the SDK
// is forward-compatible when the server adds new fields.
func TestV6_ExtraFields_ForwardCompat(t *testing.T) {
	t.Run("Session preserves unknown fields in ExtraFields", func(t *testing.T) {
		raw := `{
			"id": "ses_1",
			"slug": "my-session",
			"projectID": "proj_1",
			"directory": "/tmp",
			"title": "Test",
			"version": "1.0",
			"time": {"created": 1, "updated": 2},
			"newFieldFromFuture": "preserve_me",
			"anotherNewField": 42
		}`
		var s Session
		if err := json.Unmarshal([]byte(raw), &s); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if s.ID != "ses_1" {
			t.Errorf("Session.ID: got %q", s.ID)
		}
		if len(s.JSON.ExtraFields) != 2 {
			t.Errorf("ExtraFields: expected 2 unknown fields, got %d: %v", len(s.JSON.ExtraFields), s.JSON.ExtraFields)
		}
		if _, ok := s.JSON.ExtraFields["newFieldFromFuture"]; !ok {
			t.Error("ExtraFields missing 'newFieldFromFuture'")
		}
		// RawJSON should contain the full original payload
		rawJSON := s.JSON.RawJSON()
		if !strings.Contains(rawJSON, "newFieldFromFuture") {
			t.Errorf("RawJSON missing unknown field: %s", rawJSON)
		}
		t.Logf("ExtraFields preserved: %v", s.JSON.ExtraFields)
	})

	t.Run("ToolStateCompleted preserves unknown fields", func(t *testing.T) {
		raw := `{
			"status": "completed",
			"input": {"cmd": "ls"},
			"output": "file.txt",
			"title": "list",
			"metadata": {},
			"time": {"start": 1, "end": 2},
			"futureField": "new_in_server_v2"
		}`
		var ts ToolStateCompleted
		if err := json.Unmarshal([]byte(raw), &ts); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(ts.JSON.ExtraFields) != 1 {
			t.Errorf("ExtraFields: expected 1 unknown field, got %d", len(ts.JSON.ExtraFields))
		}
		if _, ok := ts.JSON.ExtraFields["futureField"]; !ok {
			t.Error("ExtraFields missing 'futureField'")
		}
		t.Logf("ToolStateCompleted ExtraFields preserved: %v", ts.JSON.ExtraFields)
	})

	t.Run("Message ExtraFields forward-compat", func(t *testing.T) {
		raw := `{
			"id": "msg_1",
			"sessionID": "ses_1",
			"role": "user",
			"time": {"created": 1},
			"agent": "default",
			"model": {"providerID": "p", "modelID": "m"},
			"futureUserField": "x"
		}`
		var m Message
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		rawJSON := m.JSON.RawJSON()
		if !strings.Contains(rawJSON, "futureUserField") {
			t.Errorf("Message.RawJSON missing future field: %s", rawJSON)
		}
		t.Logf("Message RawJSON captures future field")
	})
}

// TestV6_UserMessage_Format_Deserialization verifies that UserMessage.Format
// correctly captures OutputFormat values when present in JSON, and is absent
// when the Format field is not included in the response.
func TestV6_UserMessage_Format_Deserialization(t *testing.T) {
	t.Run("Format field absent", func(t *testing.T) {
		raw := `{
			"id": "msg_1",
			"sessionID": "ses_1",
			"role": "user",
			"time": {"created": 1},
			"agent": "default",
			"model": {"providerID": "p", "modelID": "m"}
		}`
		var um UserMessage
		if err := json.Unmarshal([]byte(raw), &um); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if um.Format != nil {
			t.Errorf("Format should be nil when absent, got: %v", um.Format)
		}
		t.Logf("Format absent: nil ✓")
	})

	t.Run("Format field present with text type", func(t *testing.T) {
		raw := `{
			"id": "msg_2",
			"sessionID": "ses_1",
			"role": "user",
			"time": {"created": 1},
			"agent": "default",
			"model": {"providerID": "p", "modelID": "m"},
			"format": {"type": "text"}
		}`
		var um UserMessage
		if err := json.Unmarshal([]byte(raw), &um); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if um.Format == nil {
			t.Fatal("Format should be non-nil when present")
		}
		// Format will be deserialized as map[string]any since no Response-side
		// OutputFormat union registration exists — this is expected behavior.
		formatMap, ok := um.Format.(map[string]any)
		if !ok {
			t.Fatalf("Format: expected map[string]any, got %T", um.Format)
		}
		if formatMap["type"] != "text" {
			t.Errorf("Format.type: got %v", formatMap["type"])
		}
		t.Logf("Format[text]: %v ✓", um.Format)
	})

	t.Run("Format field present with json_schema type", func(t *testing.T) {
		raw := `{
			"id": "msg_3",
			"sessionID": "ses_1",
			"role": "user",
			"time": {"created": 1},
			"agent": "default",
			"model": {"providerID": "p", "modelID": "m"},
			"format": {"type": "json_schema", "schema": {"type": "object"}}
		}`
		var um UserMessage
		if err := json.Unmarshal([]byte(raw), &um); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if um.Format == nil {
			t.Fatal("Format should be non-nil when present")
		}
		formatMap, ok := um.Format.(map[string]any)
		if !ok {
			t.Fatalf("Format: expected map[string]any, got %T", um.Format)
		}
		if formatMap["type"] != "json_schema" {
			t.Errorf("Format.type: got %v", formatMap["type"])
		}
		t.Logf("Format[json_schema]: %v ✓", um.Format)
	})
}

// TestV6_AssistantMessage_Structured_Deserialization verifies that
// AssistantMessage.Structured captures the format when returned by the server.
func TestV6_AssistantMessage_Structured_Deserialization(t *testing.T) {
	t.Run("Structured absent", func(t *testing.T) {
		raw := `{
			"id": "msg_1",
			"sessionID": "ses_1",
			"role": "assistant",
			"parentID": "msg_0",
			"modelID": "claude-3-5",
			"providerID": "anthropic",
			"mode": "default",
			"agent": "default",
			"path": {"cwd": "/tmp", "root": "/"},
			"cost": 0.001,
			"time": {"created": 1},
			"tokens": {"input": 10, "output": 5, "reasoning": 0, "cache": {"read": 0, "write": 0}}
		}`
		var am AssistantMessage
		if err := json.Unmarshal([]byte(raw), &am); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if am.Structured != nil {
			t.Errorf("Structured should be nil when absent, got: %v", am.Structured)
		}
		t.Logf("Structured absent: nil ✓")
	})

	t.Run("Structured present with text type", func(t *testing.T) {
		raw := `{
			"id": "msg_2",
			"sessionID": "ses_1",
			"role": "assistant",
			"parentID": "msg_0",
			"modelID": "claude-3-5",
			"providerID": "anthropic",
			"mode": "default",
			"agent": "default",
			"path": {"cwd": "/tmp", "root": "/"},
			"cost": 0.001,
			"time": {"created": 1},
			"tokens": {"input": 10, "output": 5, "reasoning": 0, "cache": {"read": 0, "write": 0}},
			"structured": {"type": "text"}
		}`
		var am AssistantMessage
		if err := json.Unmarshal([]byte(raw), &am); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if am.Structured == nil {
			t.Fatal("Structured should be non-nil when present")
		}
		structMap, ok := am.Structured.(map[string]any)
		if !ok {
			t.Fatalf("Structured: expected map[string]any, got %T", am.Structured)
		}
		if structMap["type"] != "text" {
			t.Errorf("Structured.type: got %v", structMap["type"])
		}
		t.Logf("Structured[text]: %v ✓", am.Structured)
	})
}

// TestV6_SessionStatus_Discriminator verifies the SessionStatus union uses
// the "type" discriminator to correctly dispatch idle/retry/busy variants.
func TestV6_SessionStatus_Discriminator(t *testing.T) {
	t.Run("idle status via discriminator", func(t *testing.T) {
		raw := `{"type": "idle"}`
		var m SessionStatusMap
		if err := json.Unmarshal([]byte(`{"ses_1": `+raw+`}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		s := m["ses_1"]
		idle, ok := s.(SessionStatusIdle)
		if !ok {
			t.Fatalf("expected SessionStatusIdle, got %T", s)
		}
		if idle.Type != "idle" {
			t.Errorf("Type: got %q", idle.Type)
		}
		t.Logf("idle discriminator: ✓")
	})

	t.Run("retry status with action field", func(t *testing.T) {
		raw := `{
			"type": "retry",
			"attempt": 2,
			"message": "rate_limited",
			"next": 1700000000,
			"action": {
				"reason": "quota",
				"provider": "anthropic",
				"title": "Rate Limit",
				"message": "Please upgrade",
				"label": "Upgrade",
				"link": "https://example.com"
			}
		}`
		var m SessionStatusMap
		if err := json.Unmarshal([]byte(`{"ses_2": `+raw+`}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		s := m["ses_2"]
		retry, ok := s.(SessionStatusRetry)
		if !ok {
			t.Fatalf("expected SessionStatusRetry, got %T", s)
		}
		if retry.Attempt != 2 {
			t.Errorf("Attempt: got %d", retry.Attempt)
		}
		if retry.Action.Link != "https://example.com" {
			t.Errorf("Action.Link: got %q", retry.Action.Link)
		}
		t.Logf("retry discriminator with action: ✓")
	})

	t.Run("retry status without action field (omitzero)", func(t *testing.T) {
		raw := `{"type": "retry", "attempt": 1, "message": "timeout", "next": 1000}`
		var m SessionStatusMap
		if err := json.Unmarshal([]byte(`{"ses_3": `+raw+`}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		s := m["ses_3"]
		retry, ok := s.(SessionStatusRetry)
		if !ok {
			t.Fatalf("expected SessionStatusRetry, got %T", s)
		}
		// Action field is omitzero — when absent, it's zero-value
		if retry.Action.Reason != "" {
			t.Errorf("empty Action.Reason expected, got %q", retry.Action.Reason)
		}
		t.Logf("retry without action: ✓")
	})

	t.Run("busy status via discriminator", func(t *testing.T) {
		raw := `{"type": "busy"}`
		var m SessionStatusMap
		if err := json.Unmarshal([]byte(`{"ses_4": `+raw+`}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		s := m["ses_4"]
		_, ok := s.(SessionStatusBusy)
		if !ok {
			t.Fatalf("expected SessionStatusBusy, got %T", s)
		}
		t.Logf("busy discriminator: ✓")
	})
}
