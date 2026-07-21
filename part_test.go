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
