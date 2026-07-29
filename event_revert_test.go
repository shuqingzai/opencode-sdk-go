package opencode

import (
	"encoding/json"
	"testing"
)

// TestRevertStagedPropertiesRevertUsesFileDiff verifies that the `files` array of
// the session.next.revert.staged event decodes into []FileDiff.
//
// OpenAPI fact: components.schemas.RevertState.properties.files.items is
// $ref -> #/components/schemas/FileDiff (NOT VcsFileDiff). FileDiff carries the
// key `path`; VcsFileDiff carries `file`. The struct previously referenced
// []VcsFileDiff, so a server payload of {"path": "..."} fell silently into
// ExtraFields and the path was lost.
//
// Regression test for that blocker.
func TestRevertStagedPropertiesRevertUsesFileDiff(t *testing.T) {
	raw := []byte(`{
		"messageID": "msg_x",
		"files": [
			{"path": "a.go", "status": "modified", "additions": 3, "deletions": 1, "patch": "@@ -1 +1 @@"}
		]
	}`)

	var revert RevertState
	if err := json.Unmarshal(raw, &revert); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if revert.MessageID != "msg_x" {
		t.Errorf("MessageID: expected %q, got %q", "msg_x", revert.MessageID)
	}
	if len(revert.Files) != 1 {
		t.Fatalf("Files: expected 1 element, got %d", len(revert.Files))
	}

	f := revert.Files[0]
	if f.Path != "a.go" {
		t.Errorf("Files[0].Path: expected %q, got %q (the `path` key was dropped)", "a.go", f.Path)
	}
	if f.Status != FileDiffStatusModified {
		t.Errorf("Files[0].Status: expected %q, got %q", FileDiffStatusModified, f.Status)
	}
	if !f.Status.IsKnown() {
		t.Errorf("Files[0].Status: %q should be a known enum value", f.Status)
	}
	if f.Additions != 3 {
		t.Errorf("Files[0].Additions: expected 3, got %d", f.Additions)
	}
	if f.Deletions != 1 {
		t.Errorf("Files[0].Deletions: expected 1, got %d", f.Deletions)
	}
	if f.Patch != "@@ -1 +1 @@" {
		t.Errorf("Files[0].Patch: expected %q, got %q", "@@ -1 +1 @@", f.Patch)
	}
}

// TestRevertStagedPropertiesRevertAllFields covers every field of the OpenAPI
// RevertState schema: messageID (required), partID, snapshot, diff, files.
func TestRevertStagedPropertiesRevertAllFields(t *testing.T) {
	raw := []byte(`{
		"messageID": "msg_full",
		"partID": "prt_1",
		"snapshot": "snap_abc",
		"diff": "diff --git a/a.go b/a.go",
		"files": [
			{"path": "added.go", "status": "added", "additions": 10, "deletions": 0, "patch": "+new"},
			{"path": "gone.go", "status": "deleted", "additions": 0, "deletions": 7, "patch": "-old"}
		]
	}`)

	var revert RevertState
	if err := json.Unmarshal(raw, &revert); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if revert.MessageID != "msg_full" {
		t.Errorf("MessageID: expected %q, got %q", "msg_full", revert.MessageID)
	}
	if revert.PartID != "prt_1" {
		t.Errorf("PartID: expected %q, got %q", "prt_1", revert.PartID)
	}
	if revert.Snapshot != "snap_abc" {
		t.Errorf("Snapshot: expected %q, got %q", "snap_abc", revert.Snapshot)
	}
	if revert.Diff != "diff --git a/a.go b/a.go" {
		t.Errorf("Diff: expected %q, got %q", "diff --git a/a.go b/a.go", revert.Diff)
	}
	if len(revert.Files) != 2 {
		t.Fatalf("Files: expected 2 elements, got %d", len(revert.Files))
	}

	wantPaths := []string{"added.go", "gone.go"}
	wantStatus := []FileDiffStatus{FileDiffStatusAdded, FileDiffStatusDeleted}
	for i, f := range revert.Files {
		if f.Path != wantPaths[i] {
			t.Errorf("Files[%d].Path: expected %q, got %q", i, wantPaths[i], f.Path)
		}
		if f.Status != wantStatus[i] {
			t.Errorf("Files[%d].Status: expected %q, got %q", i, wantStatus[i], f.Status)
		}
	}

	if revert.JSON.RawJSON() == "" {
		t.Error("RawJSON() should expose the original payload")
	}
}

// TestRevertStagedPropertiesRevertOnlyRequired verifies the boundary where only
// the required `messageID` is present: files stays nil, no error.
func TestRevertStagedPropertiesRevertOnlyRequired(t *testing.T) {
	raw := []byte(`{"messageID":"msg_min"}`)

	var revert RevertState
	if err := json.Unmarshal(raw, &revert); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if revert.MessageID != "msg_min" {
		t.Errorf("MessageID: expected %q, got %q", "msg_min", revert.MessageID)
	}
	if revert.Files != nil {
		t.Errorf("Files: expected nil for an absent array, got %#v", revert.Files)
	}
	if revert.PartID != "" || revert.Snapshot != "" || revert.Diff != "" {
		t.Errorf("optional string fields should be empty, got PartID=%q Snapshot=%q Diff=%q",
			revert.PartID, revert.Snapshot, revert.Diff)
	}
}

// TestRevertStagedEventEndToEnd decodes a full session.next.revert.staged event
// off the /event stream and reaches through to the nested file path.
func TestRevertStagedEventEndToEnd(t *testing.T) {
	raw := []byte(`{
		"id": "evt_revert1",
		"type": "session.next.revert.staged",
		"properties": {
			"timestamp": 1730000000000,
			"sessionID": "ses_abc",
			"revert": {
				"messageID": "msg_x",
				"files": [
					{"path": "a.go", "status": "modified", "additions": 3, "deletions": 1, "patch": "@@"}
				]
			}
		}
	}`)

	var resp EventListResponse
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if resp.ID != "evt_revert1" {
		t.Errorf("ID: expected %q, got %q", "evt_revert1", resp.ID)
	}
	if resp.Type != EventListResponseTypeSessionNextRevertStaged {
		t.Errorf("Type: expected %q, got %q", EventListResponseTypeSessionNextRevertStaged, resp.Type)
	}

	variant, ok := resp.AsUnion().(EventListResponseEventSessionNextRevertStaged)
	if !ok {
		t.Fatalf("AsUnion(): expected EventListResponseEventSessionNextRevertStaged, got %T", resp.AsUnion())
	}
	if variant.Properties.SessionID != "ses_abc" {
		t.Errorf("SessionID: expected %q, got %q", "ses_abc", variant.Properties.SessionID)
	}
	if variant.Properties.Timestamp != 1730000000000 {
		t.Errorf("Timestamp: expected 1730000000000, got %d", variant.Properties.Timestamp)
	}
	if len(variant.Properties.Revert.Files) != 1 {
		t.Fatalf("Revert.Files: expected 1 element, got %d", len(variant.Properties.Revert.Files))
	}
	if got := variant.Properties.Revert.Files[0].Path; got != "a.go" {
		t.Errorf("Revert.Files[0].Path: expected %q, got %q", "a.go", got)
	}
}

// TestSyncEventRevertStagedUsesFileDiff verifies the sync event surface, which
// reuses EventListResponseEventSessionNextRevertStagedProperties for its `data`.
// OpenAPI: SyncEventSessionNextRevertStaged.syncEvent.data.revert -> RevertState.
func TestSyncEventRevertStagedUsesFileDiff(t *testing.T) {
	raw := []byte(`{
		"type": "session.next.revert.staged.1",
		"id": "evt_sync1",
		"seq": 42,
		"aggregateID": "ses_abc",
		"data": {
			"timestamp": 1730000000000,
			"sessionID": "ses_abc",
			"revert": {
				"messageID": "msg_x",
				"files": [
					{"path": "a.go", "status": "modified", "additions": 3, "deletions": 1, "patch": "@@"}
				]
			}
		}
	}`)

	var ev SyncEventSessionNextRevertStaged
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if ev.Seq != 42 {
		t.Errorf("Seq: expected 42, got %d", ev.Seq)
	}
	if !ev.Type.IsKnown() {
		t.Errorf("Type %q should be known", ev.Type)
	}
	if len(ev.Data.Revert.Files) != 1 {
		t.Fatalf("Data.Revert.Files: expected 1 element, got %d", len(ev.Data.Revert.Files))
	}
	if got := ev.Data.Revert.Files[0].Path; got != "a.go" {
		t.Errorf("Data.Revert.Files[0].Path: expected %q, got %q", "a.go", got)
	}
}

// TestPromptTypeSharedBetweenPromptedAndAdmitted verifies that the merged
// `Prompt` type decodes identically for both session.next.prompted and
// session.next.prompt.admitted.
//
// OpenAPI fact: EventSessionNextPrompted.properties.prompt and
// EventSessionNextPromptAdmitted.properties.prompt both $ref the very same
// #/components/schemas/Prompt, so a single Go type is correct.
func TestPromptTypeSharedBetweenPromptedAndAdmitted(t *testing.T) {
	promptJSON := `{
		"text": "hello world",
		"files": [{"uri": "file:///a.go", "mime": "text/x-go", "name": "a.go"}],
		"agents": [{"name": "build"}]
	}`

	check := func(t *testing.T, p V2SessionInputPrompt) {
		t.Helper()
		if p.Text != "hello world" {
			t.Errorf("Text: expected %q, got %q", "hello world", p.Text)
		}
		if len(p.Files) != 1 {
			t.Fatalf("Files: expected 1 element, got %d", len(p.Files))
		}
		if p.Files[0].URI != "file:///a.go" {
			t.Errorf("Files[0].URI: expected %q, got %q", "file:///a.go", p.Files[0].URI)
		}
		if p.Files[0].Mime != "text/x-go" {
			t.Errorf("Files[0].Mime: expected %q, got %q", "text/x-go", p.Files[0].Mime)
		}
		if len(p.Agents) != 1 {
			t.Fatalf("Agents: expected 1 element, got %d", len(p.Agents))
		}
		if p.Agents[0].Name != "build" {
			t.Errorf("Agents[0].Name: expected %q, got %q", "build", p.Agents[0].Name)
		}
	}

	t.Run("session.next.prompted", func(t *testing.T) {
		raw := []byte(`{
			"timestamp": 1730000000000,
			"sessionID": "ses_abc",
			"messageID": "msg_x",
			"prompt": ` + promptJSON + `,
			"delivery": "steer"
		}`)
		var props EventListResponseEventSessionNextPromptedProperties
		if err := json.Unmarshal(raw, &props); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		check(t, props.Prompt)
		if props.Delivery != EventListResponseEventSessionNextPromptedDeliverySteer {
			t.Errorf("Delivery: expected steer, got %q", props.Delivery)
		}
	})

	t.Run("session.next.prompt.admitted", func(t *testing.T) {
		raw := []byte(`{
			"timestamp": 1730000000000,
			"sessionID": "ses_abc",
			"messageID": "msg_x",
			"prompt": ` + promptJSON + `,
			"delivery": "queue"
		}`)
		var props EventListResponseEventSessionNextPromptAdmittedProperties
		if err := json.Unmarshal(raw, &props); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		check(t, props.Prompt)
	})

	t.Run("sync.session.next.prompt.admitted", func(t *testing.T) {
		raw := []byte(`{
			"timestamp": 1730000000000,
			"sessionID": "ses_abc",
			"messageID": "msg_x",
			"prompt": ` + promptJSON + `,
			"delivery": "queue"
		}`)
		var props SyncEventSessionNextPromptAdmittedProperties
		if err := json.Unmarshal(raw, &props); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		check(t, props.Prompt)
		if props.Delivery != SyncEventSessionNextPromptAdmittedDeliveryQueue {
			t.Errorf("Delivery: expected queue, got %q", props.Delivery)
		}
	})
}

// TestPromptOnlyRequiredText verifies the boundary where only the required
// `text` field is present (OpenAPI Prompt.required == ["text"]).
func TestPromptOnlyRequiredText(t *testing.T) {
	raw := []byte(`{"text":"just text"}`)

	var p V2SessionInputPrompt
	if err := json.Unmarshal(raw, &p); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if p.Text != "just text" {
		t.Errorf("Text: expected %q, got %q", "just text", p.Text)
	}
	if p.Files != nil {
		t.Errorf("Files: expected nil, got %#v", p.Files)
	}
	if p.Agents != nil {
		t.Errorf("Agents: expected nil, got %#v", p.Agents)
	}
}

// TestFileDiffVersusVcsFileDiffKeys pins the distinction that caused the bug:
// FileDiff reads `path`, VcsFileDiff reads `file`. They are separate OpenAPI
// schemas and must not be substituted for one another.
func TestFileDiffVersusVcsFileDiffKeys(t *testing.T) {
	var fd FileDiff
	if err := json.Unmarshal([]byte(`{"path":"p.go","status":"added","additions":1,"deletions":0,"patch":"+x"}`), &fd); err != nil {
		t.Fatalf("FileDiff unmarshal failed: %v", err)
	}
	if fd.Path != "p.go" {
		t.Errorf("FileDiff.Path: expected %q, got %q", "p.go", fd.Path)
	}

	var vfd VcsFileDiff
	if err := json.Unmarshal([]byte(`{"file":"f.go","additions":1,"deletions":0}`), &vfd); err != nil {
		t.Fatalf("VcsFileDiff unmarshal failed: %v", err)
	}
	if vfd.File != "f.go" {
		t.Errorf("VcsFileDiff.File: expected %q, got %q", "f.go", vfd.File)
	}

	// A FileDiff payload fed to VcsFileDiff loses the path — this is exactly the
	// silent data loss the RevertState.files fix eliminates.
	var wrong VcsFileDiff
	if err := json.Unmarshal([]byte(`{"path":"p.go","status":"added","additions":1,"deletions":0,"patch":"+x"}`), &wrong); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if wrong.File != "" {
		t.Errorf("VcsFileDiff.File should stay empty for a FileDiff payload, got %q", wrong.File)
	}
}
