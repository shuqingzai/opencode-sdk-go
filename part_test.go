package opencode

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPartUpdateParamsBodySerialization(t *testing.T) {
	t.Run("Part set — Text variant serializes at root", func(t *testing.T) {
		params := PartUpdateParams{
			Part: F(PartUpdatePartUnion(PartUpdatePartText{
				ID:        F("prt_1"),
				SessionID: F("ses_1"),
				MessageID: F("msg_1"),
				Text:      F("hello"),
				Type:      F(PartUpdatePartTextTypeText),
			})),
		}
		b, err := json.Marshal(params)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"id":"prt_1","messageID":"msg_1","sessionID":"ses_1","text":"hello","type":"text"}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Part not set — MarshalJSON returns nil (no body sent)", func(t *testing.T) {
		params := PartUpdateParams{}
		b, err := params.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		if b != nil {
			t.Errorf("expected nil bytes when Part is unset, got %q", string(b))
		}
	})

	t.Run("Directory/Workspace set, Part not set — MarshalJSON still nil body", func(t *testing.T) {
		params := PartUpdateParams{
			Directory: F("d"),
			Workspace: F("w"),
		}
		b, err := params.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		if b != nil {
			t.Errorf("expected nil bytes when Part is unset, got %q", string(b))
		}
	})

	t.Run("Part set — File variant serializes at root", func(t *testing.T) {
		params := PartUpdateParams{
			Part: F(PartUpdatePartUnion(PartUpdatePartFile{
				ID:        F("prt_2"),
				SessionID: F("ses_1"),
				MessageID: F("msg_1"),
				Mime:      F("text/plain"),
				URL:       F("file:///tmp/a.txt"),
				Type:      F(PartUpdatePartFileTypeFile),
				Filename:  F("a.txt"),
			})),
		}
		b, err := json.Marshal(params)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"id":"prt_2"`,
			`"mime":"text/plain"`,
			`"type":"file"`,
			`"filename":"a.txt"`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	t.Run("Part set — Compaction variant with tail_start_id", func(t *testing.T) {
		params := PartUpdateParams{
			Part: F(PartUpdatePartUnion(PartUpdatePartCompaction{
				ID:          F("prt_3"),
				SessionID:   F("ses_1"),
				MessageID:   F("msg_1"),
				Auto:        F(true),
				Type:        F(PartUpdatePartCompactionTypeCompaction),
				TailStartID: F("msg_99"),
			})),
		}
		b, err := json.Marshal(params)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"auto":true`,
			`"tail_start_id":"msg_99"`,
			`"type":"compaction"`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})
}

// TestToolStateParamsSerialization verifies that Request-side ToolState variants
// serialize all required fields correctly per OpenAPI schema.
func TestToolStateParamsSerialization(t *testing.T) {
	t.Run("ToolStatePending — status+input+raw required", func(t *testing.T) {
		s := PartUpdatePartToolStatePending{
			Status: F(PartUpdatePartToolStatePendingStatusPending),
			Input:  F(map[string]any{"key": "val"}),
			Raw:    F(`{"key":"val"}`),
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
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	t.Run("ToolStateRunning — status+input+time required", func(t *testing.T) {
		s := PartUpdatePartToolStateRunning{
			Input:  F(map[string]any{"arg": 1}),
			Status: F(PartUpdatePartToolStateRunningStatusRunning),
			Time:   F(PartUpdatePartToolStateRunningTime{Start: F(int64(1700000000))}),
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"status":"running"`,
			`"input":`,
			`"time":`,
			`"start":1700000000`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	t.Run("ToolStateRunning — optional title+metadata omitted when unset", func(t *testing.T) {
		s := PartUpdatePartToolStateRunning{
			Input:  F(map[string]any{}),
			Status: F(PartUpdatePartToolStateRunningStatusRunning),
			Time:   F(PartUpdatePartToolStateRunningTime{Start: F(int64(1700000001))}),
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if strings.Contains(got, `"title"`) {
			t.Errorf("title should be absent when unset, got %s", got)
		}
		if strings.Contains(got, `"metadata"`) {
			t.Errorf("metadata should be absent when unset, got %s", got)
		}
	})

	t.Run("ToolStateCompleted — all required fields present", func(t *testing.T) {
		s := PartUpdatePartToolStateCompleted{
			Input:    F(map[string]any{"x": 1}),
			Metadata: F(map[string]any{"provider": "test"}),
			Output:   F("result text"),
			Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
			Time: F(PartUpdatePartToolStateCompletedTime{
				Start: F(int64(1700000000)),
				End:   F(int64(1700000010)),
			}),
			Title: F("My Tool"),
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"status":"completed"`,
			`"input":`,
			`"output":"result text"`,
			`"title":"My Tool"`,
			`"metadata":`,
			`"time":`,
			`"start":1700000000`,
			`"end":1700000010`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	t.Run("ToolStateCompleted — optional attachments serialized when set", func(t *testing.T) {
		s := PartUpdatePartToolStateCompleted{
			Input:    F(map[string]any{}),
			Metadata: F(map[string]any{}),
			Output:   F("ok"),
			Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
			Time: F(PartUpdatePartToolStateCompletedTime{
				Start: F(int64(1700000000)),
				End:   F(int64(1700000010)),
			}),
			Title: F("T"),
			Attachments: F([]PartUpdatePartFile{
				{
					ID:        F("prt_99"),
					SessionID: F("ses_1"),
					MessageID: F("msg_1"),
					Mime:      F("image/png"),
					URL:       F("file:///img.png"),
					Type:      F(PartUpdatePartFileTypeFile),
				},
			}),
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"attachments":[`,
			`"prt_99"`,
			`"image/png"`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	t.Run("ToolStateCompleted — optional compacted in time", func(t *testing.T) {
		s := PartUpdatePartToolStateCompleted{
			Input:    F(map[string]any{}),
			Metadata: F(map[string]any{}),
			Output:   F("ok"),
			Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
			Time: F(PartUpdatePartToolStateCompletedTime{
				Start:     F(int64(1700000000)),
				End:       F(int64(1700000010)),
				Compacted: F(int64(1700000005)),
			}),
			Title: F("T"),
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"compacted":1700000005`) {
			t.Errorf("missing compacted field in %s", got)
		}
	})

	t.Run("ToolStateCompleted — compacted absent when unset", func(t *testing.T) {
		s := PartUpdatePartToolStateCompleted{
			Input:    F(map[string]any{}),
			Metadata: F(map[string]any{}),
			Output:   F("ok"),
			Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
			Time: F(PartUpdatePartToolStateCompletedTime{
				Start: F(int64(1700000000)),
				End:   F(int64(1700000010)),
			}),
			Title: F("T"),
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if strings.Contains(got, `"compacted"`) {
			t.Errorf("compacted should be absent when unset, got %s", got)
		}
	})

	t.Run("ToolStateError — status+input+error+time required", func(t *testing.T) {
		s := PartUpdatePartToolStateError{
			Error:  F("something failed"),
			Input:  F(map[string]any{"cmd": "ls"}),
			Status: F(PartUpdatePartToolStateErrorStatusError),
			Time: F(PartUpdatePartToolStateErrorTime{
				Start: F(int64(1700000000)),
				End:   F(int64(1700000002)),
			}),
		}
		b, err := json.Marshal(s)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"status":"error"`,
			`"error":"something failed"`,
			`"input":`,
			`"time":`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})
}

// TestSessionListParamsScopeSerialization verifies the Scope enum serializes correctly.
func TestSessionListParamsScopeSerialization(t *testing.T) {
	t.Run("Scope=project serializes as query string", func(t *testing.T) {
		params := SessionListParams{
			Scope: F(SessionListParamsScopeProject),
		}
		v := params.URLQuery()
		if got := v.Get("scope"); got != "project" {
			t.Errorf("expected scope=project, got %q", got)
		}
	})

	t.Run("Scope unset — not in query string", func(t *testing.T) {
		params := SessionListParams{}
		v := params.URLQuery()
		if v.Has("scope") {
			t.Errorf("scope should be absent when unset, got %s", v.Encode())
		}
	})

	t.Run("SessionListParamsScope.IsKnown", func(t *testing.T) {
		if !SessionListParamsScopeProject.IsKnown() {
			t.Error("SessionListParamsScopeProject.IsKnown() should return true")
		}
		unknown := SessionListParamsScope("unknown_value")
		if unknown.IsKnown() {
			t.Error("unknown scope should return false from IsKnown()")
		}
	})
}
