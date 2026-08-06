package opencode

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPartToolStateSerialization(t *testing.T) {
	t.Parallel()
	// Task 1: PartUpdatePartToolStatePending must include status, input, raw (all required).
	t.Run("ToolStatePending serializes status+input+raw", func(t *testing.T) {
		t.Parallel()
		state := PartUpdatePartToolStatePending{
			Status: F(PartUpdatePartToolStatePendingStatusPending),
			Input:  F(map[string]any{"tool_call_id": "tc_1"}),
			Raw:    F(`{"tool_call_id":"tc_1"}`),
		}
		b, err := json.Marshal(state)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{`"status":"pending"`, `"input"`, `"raw"`} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	// Task 2: PartUpdatePartToolStateRunning must include input (required alongside status and time).
	t.Run("ToolStateRunning serializes input", func(t *testing.T) {
		t.Parallel()
		state := PartUpdatePartToolStateRunning{
			Status: F(PartUpdatePartToolStateRunningStatusRunning),
			Input:  F(map[string]any{"cmd": "ls"}),
			Time:   F(PartUpdatePartToolStateRunningTime{Start: F(int64(1000))}),
		}
		b, err := json.Marshal(state)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{`"status":"running"`, `"input"`, `"time"`} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	// Task 3: PartUpdatePartToolStateCompleted must include all 6 required keys;
	// optional attachments are serialized when set.
	t.Run("ToolStateCompleted serializes all required keys", func(t *testing.T) {
		t.Parallel()
		state := PartUpdatePartToolStateCompleted{
			Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
			Input:    F(map[string]any{"cmd": "ls"}),
			Output:   F("file.txt\n"),
			Title:    F("List files"),
			Metadata: F(map[string]any{"duration": 42}),
			Time: F(PartUpdatePartToolStateCompletedTime{
				Start: F(int64(1000)),
				End:   F(int64(2000)),
			}),
		}
		b, err := json.Marshal(state)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"status":"completed"`,
			`"input"`,
			`"output":"file.txt\n"`,
			`"title":"List files"`,
			`"metadata"`,
			`"time"`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	t.Run("ToolStateCompleted serializes optional attachments", func(t *testing.T) {
		t.Parallel()
		state := PartUpdatePartToolStateCompleted{
			Status:   F(PartUpdatePartToolStateCompletedStatusCompleted),
			Input:    F(map[string]any{}),
			Output:   F("result"),
			Title:    F("run"),
			Metadata: F(map[string]any{}),
			Time: F(PartUpdatePartToolStateCompletedTime{
				Start: F(int64(1000)),
				End:   F(int64(2000)),
			}),
			Attachments: F([]FilePartInputParam{
				{
					Mime: F("image/png"),
					Type: F(FilePartInputTypeFile),
					URL:  F("file:///out.png"),
				},
			}),
		}
		b, err := json.Marshal(state)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{`"attachments"`, `"mime":"image/png"`, `"type":"file"`} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	// Task 4: PartUpdatePartToolStateCompletedTime must serialize optional compacted field.
	t.Run("ToolStateCompletedTime serializes compacted", func(t *testing.T) {
		t.Parallel()
		timeVal := PartUpdatePartToolStateCompletedTime{
			Start:     F(int64(1000)),
			End:       F(int64(2000)),
			Compacted: F(int64(500)),
		}
		b, err := json.Marshal(timeVal)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{`"start":1000`, `"end":2000`, `"compacted":500`} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	t.Run("ToolStateCompletedTime omits compacted when not set", func(t *testing.T) {
		t.Parallel()
		timeVal := PartUpdatePartToolStateCompletedTime{
			Start: F(int64(1000)),
			End:   F(int64(2000)),
		}
		b, err := json.Marshal(timeVal)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if strings.Contains(got, `"compacted"`) {
			t.Errorf("unexpected \"compacted\" key in %s", got)
		}
	})
}

func TestPartUpdateParamsBodySerialization(t *testing.T) {
	t.Parallel()
	t.Run("Part set — Text variant serializes at root", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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
		t.Parallel()
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
		t.Parallel()
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

	t.Run("Part set — File variant with source union serializes", func(t *testing.T) {
		t.Parallel()
		params := PartUpdateParams{
			Part: F(PartUpdatePartUnion(PartUpdatePartFile{
				ID:        F("prt_4"),
				SessionID: F("ses_1"),
				MessageID: F("msg_1"),
				Mime:      F("text/plain"),
				URL:       F("file:///tmp/b.ts"),
				Type:      F(PartUpdatePartFileTypeFile),
				Filename:  F("b.ts"),
				Source: F(FilePartSourceUnionParam(FileSourceParam{
					Path: F("/tmp/b.ts"),
					Text: F(FilePartSourceTextParam{
						End:   F(int64(0)),
						Start: F(int64(10)),
						Value: F("const x = 1"),
					}),
					Type: F(FileSourceTypeFile),
				})),
			})),
		}
		b, err := json.Marshal(params)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		for _, want := range []string{
			`"id":"prt_4"`,
			`"type":"file"`,
			`"source":{"path":"/tmp/b.ts"`,
			`"text":{"end":0,"start":10,"value":"const x = 1"}`,
		} {
			if !strings.Contains(got, want) {
				t.Errorf("missing %q in %s", want, got)
			}
		}
	})

	t.Run("Part set — Compaction variant with tail_start_id", func(t *testing.T) {
		t.Parallel()
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
