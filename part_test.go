package opencode

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPartUpdateParamsBodySerialization(t *testing.T) {
	tests := []struct {
		name     string
		params   PartUpdateParams
		contains string
		exact    string
	}{
		{
			name: "Part set — serializes Part fields",
			params: PartUpdateParams{
				Part: F(Part{ID: "1", Type: "text", Text: "hello"}),
			},
			contains: `"id":"1"`,
		},
		{
			name:   "Part not set — empty object",
			params: PartUpdateParams{},
			exact:  `{}`,
		},
		{
			name: "Directory/Workspace set, Part not set — empty object",
			params: PartUpdateParams{
				Directory: F("d"),
				Workspace: F("w"),
			},
			exact: `{}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.params)
			if err != nil {
				t.Fatal(err)
			}
			got := string(b)
			if tt.exact != "" && got != tt.exact {
				t.Errorf("expected %s, got %s", tt.exact, got)
			}
			if tt.contains != "" && !strings.Contains(got, tt.contains) {
				t.Errorf("expected containing %s, got %s", tt.contains, got)
			}
		})
	}
}
