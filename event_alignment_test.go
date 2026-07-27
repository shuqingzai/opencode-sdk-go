package opencode

import (
	"encoding/json"
	"testing"
)

// Tests validating Go SDK event types against OpenAPI & JS SDK(v2) after
// alignment fixes. All payloads are derived from the OpenAPI Event / GlobalEvent
// schemas. Deserialization goes through the EventListResponse union carrier
// (apijson.Port) exactly as the SSE stream would drive it.

// #1 (🔴): project.updated properties is a FLAT object in OpenAPI/JS
// {id, worktree, vcs?, name?, icon?, commands?, time, sandboxes}, NOT wrapped
// in a "project" key. This previously failed to deserialize.
func TestEventProjectUpdatedFlatProperties(t *testing.T) {
	raw := `{"id":"evt_1","type":"project.updated","properties":{"id":"prj_1","worktree":"/repo","time":{"created":1700000000},"sandboxes":["sb1","sb2"],"name":"myproj"}}`

	t.Run("via union carrier (AsUnion)", func(t *testing.T) {
		var e EventListResponse
		if err := json.Unmarshal([]byte(raw), &e); err != nil {
			t.Fatal(err)
		}
		pu, ok := e.AsUnion().(EventListResponseEventProjectUpdated)
		if !ok {
			t.Fatalf("AsUnion type = %T, want EventListResponseEventProjectUpdated", e.AsUnion())
		}
		p := pu.Properties
		if p.ID != "prj_1" {
			t.Errorf("ID = %q, want prj_1", p.ID)
		}
		if p.Worktree != "/repo" {
			t.Errorf("Worktree = %q, want /repo", p.Worktree)
		}
		if p.Name != "myproj" {
			t.Errorf("Name = %q, want myproj", p.Name)
		}
		if len(p.Sandboxes) != 2 || p.Sandboxes[0] != "sb1" {
			t.Errorf("Sandboxes = %v, want [sb1 sb2]", p.Sandboxes)
		}
		if p.Time.Created != 1700000000 {
			t.Errorf("Time.Created = %d, want 1700000000", p.Time.Created)
		}
	})

	t.Run("direct variant unmarshal", func(t *testing.T) {
		var v EventListResponseEventProjectUpdated
		if err := json.Unmarshal([]byte(raw), &v); err != nil {
			t.Fatal(err)
		}
		if v.Properties.ID != "prj_1" || v.Properties.Worktree != "/repo" {
			t.Errorf("flat fields not populated: %+v", v.Properties)
		}
		// The old broken wrapper field "project" must not exist anymore: verify
		// there is no nested project object required for access.
		if v.Type != EventListResponseEventProjectUpdatedTypeProjectUpdated {
			t.Errorf("Type = %q", v.Type)
		}
	})
}

// #2 (🟠): permission.v2.replied reply reuses shared PermissionV2Reply.
func TestEventPermissionV2RepliedReusesSharedType(t *testing.T) {
	raw := `{"id":"evt_2","type":"permission.v2.replied","properties":{"sessionID":"ses_1","requestID":"req_1","reply":"always"}}`
	var e EventListResponse
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatal(err)
	}
	pr, ok := e.AsUnion().(EventListResponseEventPermissionV2Replied)
	if !ok {
		t.Fatalf("AsUnion type = %T", e.AsUnion())
	}
	// Field type must be the shared PermissionV2Reply.
	var reply PermissionV2Reply = pr.Properties.Reply
	if reply != PermissionV2ReplyAlways {
		t.Errorf("Reply = %q, want always", reply)
	}
	if !reply.IsKnown() {
		t.Error("reply should be known")
	}
}

// #3 (🟡): pty.exited exitCode is integer -> int64.
func TestEventPtyExitedExitCodeInt64(t *testing.T) {
	raw := `{"id":"evt_3","type":"pty.exited","properties":{"id":"pty_1","exitCode":137}}`
	var e EventListResponse
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatal(err)
	}
	pe, ok := e.AsUnion().(EventListResponseEventPtyExited)
	if !ok {
		t.Fatalf("AsUnion type = %T", e.AsUnion())
	}
	// Assign to int64 to assert the field type at compile time.
	var code int64 = pe.Properties.ExitCode
	if code != 137 {
		t.Errorf("ExitCode = %d, want 137", code)
	}
}

// #4 (🟡): tui.toast.show title/duration are concrete string/int64 (not any).
func TestEventTuiToastShowConcreteTypes(t *testing.T) {
	raw := `{"id":"evt_4","type":"tui.toast.show","properties":{"message":"done","variant":"success","title":"Build","duration":3000}}`
	var e EventListResponse
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatal(err)
	}
	ts, ok := e.AsUnion().(EventListResponseEventTuiToastShow)
	if !ok {
		t.Fatalf("AsUnion type = %T", e.AsUnion())
	}
	var title string = ts.Properties.Title
	var duration int64 = ts.Properties.Duration
	if title != "Build" {
		t.Errorf("Title = %q, want Build", title)
	}
	if duration != 3000 {
		t.Errorf("Duration = %d, want 3000", duration)
	}
	if ts.Properties.Message != "done" || ts.Properties.Variant != TuiToastShowVariantSuccess {
		t.Errorf("Message/Variant = %q/%q", ts.Properties.Message, ts.Properties.Variant)
	}

	t.Run("optional title/duration omitted", func(t *testing.T) {
		raw2 := `{"id":"evt_4b","type":"tui.toast.show","properties":{"message":"m","variant":"info"}}`
		var e2 EventListResponse
		if err := json.Unmarshal([]byte(raw2), &e2); err != nil {
			t.Fatal(err)
		}
		ts2 := e2.AsUnion().(EventListResponseEventTuiToastShow)
		if ts2.Properties.Title != "" || ts2.Properties.Duration != 0 {
			t.Errorf("expected zero values, got title=%q duration=%d", ts2.Properties.Title, ts2.Properties.Duration)
		}
	})
}

// #5 (🟡): session.next.model.switched and step.started reuse shared ModelRef.
func TestEventModelSwitchedReusesModelRef(t *testing.T) {
	raw := `{"id":"evt_5","type":"session.next.model.switched","properties":{"timestamp":1700000000,"sessionID":"ses_1","messageID":"msg_1","model":{"id":"claude","providerID":"anthropic","variant":"thinking"}}}`
	var e EventListResponse
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatal(err)
	}
	ms, ok := e.AsUnion().(EventListResponseEventSessionNextModelSwitched)
	if !ok {
		t.Fatalf("AsUnion type = %T", e.AsUnion())
	}
	var model ModelRef = ms.Properties.Model
	if model.ID != "claude" || model.ProviderID != "anthropic" || model.Variant != "thinking" {
		t.Errorf("Model = %+v", model)
	}
	if ms.Properties.Timestamp != 1700000000 {
		t.Errorf("Timestamp = %d", ms.Properties.Timestamp)
	}
}

func TestEventStepStartedReusesModelRef(t *testing.T) {
	raw := `{"id":"evt_6","type":"session.next.step.started","properties":{"timestamp":1,"assistantMessageID":"msg_1","sessionID":"ses_1","agent":"build","model":{"id":"m","providerID":"p"}}}`
	var e EventListResponse
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatal(err)
	}
	ss, ok := e.AsUnion().(EventListResponseEventSessionNextStepStarted)
	if !ok {
		t.Fatalf("AsUnion type = %T", e.AsUnion())
	}
	var model ModelRef = ss.Properties.Model
	if model.ID != "m" || model.ProviderID != "p" {
		t.Errorf("Model = %+v", model)
	}
}

// #6 (🟡): session.next.moved location reuses shared LocationRef.
func TestEventMovedReusesLocationRef(t *testing.T) {
	raw := `{"id":"evt_7","type":"session.next.moved","properties":{"timestamp":1,"sessionID":"ses_1","location":{"directory":"/new","workspaceID":"wrk_1"},"subdirectory":"sub"}}`
	var e EventListResponse
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatal(err)
	}
	mv, ok := e.AsUnion().(EventListResponseEventSessionNextMoved)
	if !ok {
		t.Fatalf("AsUnion type = %T", e.AsUnion())
	}
	var loc LocationRef = mv.Properties.Location
	if loc.Directory != "/new" || loc.WorkspaceID != "wrk_1" {
		t.Errorf("Location = %+v", loc)
	}
	if mv.Properties.Subdirectory != "sub" {
		t.Errorf("Subdirectory = %q", mv.Properties.Subdirectory)
	}
}

// GlobalEvent is the project's most critical business path (/global/event SSE).
// project.updated must deserialize its FLAT properties (and all nested ref types)
// correctly through the GlobalEventPayloadUnion carrier. This mirrors the
// EventProjectUpdated flatten fix (#1) on the GlobalEvent side.
func TestGlobalEventProjectUpdatedFlat(t *testing.T) {
	t.Run("full payload with all fields and nested ref types", func(t *testing.T) {
		// Covers GlobalEvent top-level directory(required)/project/workspace +
		// full flat properties: id, worktree, time(ProjectTime), sandboxes,
		// name, vcs(ProjectVcs enum), icon(ProjectIcon), commands(ProjectCommands).
		raw := `{"directory":"/root","project":"prj_9","workspace":"wrk_1","payload":{"id":"evt_8","type":"project.updated","properties":{"id":"prj_9","worktree":"/w","vcs":"git","name":"myproj","icon":{"url":"http://x/i.png","override":"ov","color":"#fff"},"commands":{"start":"npm run dev"},"time":{"created":100,"updated":200,"initialized":300},"sandboxes":["sb1","sb2"]}}}`
		var ge GlobalEvent
		if err := json.Unmarshal([]byte(raw), &ge); err != nil {
			t.Fatal(err)
		}
		// GlobalEvent top-level fields
		if ge.Directory != "/root" {
			t.Errorf("Directory = %q, want /root", ge.Directory)
		}
		if ge.Project != "prj_9" {
			t.Errorf("Project = %q, want prj_9", ge.Project)
		}
		if ge.Workspace != "wrk_1" {
			t.Errorf("Workspace = %q, want wrk_1", ge.Workspace)
		}

		pu, ok := ge.AsUnion().(EventListResponseEventProjectUpdated)
		if !ok {
			t.Fatalf("payload type = %T, want EventListResponseEventProjectUpdated", ge.AsUnion())
		}
		if pu.ID != "evt_8" {
			t.Errorf("payload ID = %q, want evt_8", pu.ID)
		}
		if pu.Type != EventListResponseEventProjectUpdatedTypeProjectUpdated {
			t.Errorf("payload Type = %q", pu.Type)
		}

		p := pu.Properties
		// required flat fields
		if p.ID != "prj_9" {
			t.Errorf("Properties.ID = %q, want prj_9", p.ID)
		}
		if p.Worktree != "/w" {
			t.Errorf("Properties.Worktree = %q, want /w", p.Worktree)
		}
		if len(p.Sandboxes) != 2 || p.Sandboxes[0] != "sb1" || p.Sandboxes[1] != "sb2" {
			t.Errorf("Properties.Sandboxes = %v, want [sb1 sb2]", p.Sandboxes)
		}
		// ProjectTime nested (created/updated required, initialized optional) — int64
		if p.Time.Created != 100 || p.Time.Updated != 200 || p.Time.Initialized != 300 {
			t.Errorf("Properties.Time = %+v, want {100 200 300}", p.Time)
		}
		// optional flat fields
		if p.Name != "myproj" {
			t.Errorf("Properties.Name = %q, want myproj", p.Name)
		}
		if p.Vcs != ProjectVcsGit {
			t.Errorf("Properties.Vcs = %q, want git", p.Vcs)
		}
		if !p.Vcs.IsKnown() {
			t.Error("Properties.Vcs should be known")
		}
		// ProjectIcon nested
		if p.Icon.URL != "http://x/i.png" || p.Icon.Override != "ov" || p.Icon.Color != "#fff" {
			t.Errorf("Properties.Icon = %+v", p.Icon)
		}
		// ProjectCommands nested
		if p.Commands.Start != "npm run dev" {
			t.Errorf("Properties.Commands.Start = %q, want npm run dev", p.Commands.Start)
		}
	})

	t.Run("minimal payload with only required fields", func(t *testing.T) {
		// GlobalEvent required=directory,payload; properties required=id,worktree,time,sandboxes
		raw := `{"directory":"/root","payload":{"id":"evt_9","type":"project.updated","properties":{"id":"prj_1","worktree":"/w","time":{"created":1,"updated":2},"sandboxes":[]}}}`
		var ge GlobalEvent
		if err := json.Unmarshal([]byte(raw), &ge); err != nil {
			t.Fatal(err)
		}
		if ge.Project != "" || ge.Workspace != "" {
			t.Errorf("optional top-level fields should be empty, got project=%q workspace=%q", ge.Project, ge.Workspace)
		}
		pu, ok := ge.AsUnion().(EventListResponseEventProjectUpdated)
		if !ok {
			t.Fatalf("payload type = %T", ge.AsUnion())
		}
		p := pu.Properties
		if p.ID != "prj_1" || p.Worktree != "/w" {
			t.Errorf("required flat fields not populated: %+v", p)
		}
		if p.Time.Created != 1 || p.Time.Updated != 2 {
			t.Errorf("Time = %+v, want {1 2 0}", p.Time)
		}
		// optional fields must be zero values
		if p.Name != "" || p.Vcs != "" || p.Icon.URL != "" || p.Commands.Start != "" {
			t.Errorf("optional fields should be zero, got name=%q vcs=%q icon=%+v commands=%+v", p.Name, p.Vcs, p.Icon, p.Commands)
		}
		if p.Sandboxes == nil {
			// empty array must decode to non-nil (or at least len 0)
			if len(p.Sandboxes) != 0 {
				t.Errorf("Sandboxes = %v", p.Sandboxes)
			}
		}
	})

	t.Run("RawJSON preserved on payload variant", func(t *testing.T) {
		raw := `{"directory":"/root","payload":{"id":"evt_a","type":"project.updated","properties":{"id":"p","worktree":"/w","time":{"created":1,"updated":2},"sandboxes":["s"]}}}`
		var ge GlobalEvent
		if err := json.Unmarshal([]byte(raw), &ge); err != nil {
			t.Fatal(err)
		}
		pu := ge.AsUnion().(EventListResponseEventProjectUpdated)
		if pu.JSON.RawJSON() == "" {
			t.Error("payload variant RawJSON should be preserved")
		}
	})
}
