package opencode

import (
	"encoding/json"
	"testing"
)

func TestSessionStatusMapUnmarshal(t *testing.T) {
	t.Parallel()

	data := `{
		"ses_idle":  {"type": "idle"},
		"ses_busy":  {"type": "busy"},
		"ses_retry": {"type": "retry", "attempt": 3, "message": "rate limited", "next": 1700000000}
	}`

	var m SessionStatusMap
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(m) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(m))
	}

	tests := []struct {
		id       string
		wantType string
		check    func(t *testing.T, status SessionStatus)
	}{
		{
			id:       "ses_idle",
			wantType: "SessionStatusIdle",
			check: func(t *testing.T, status SessionStatus) {
				_, ok := status.(SessionStatusIdle)
				if !ok {
					t.Errorf("expected SessionStatusIdle, got %T", status)
				}
			},
		},
		{
			id:       "ses_busy",
			wantType: "SessionStatusBusy",
			check: func(t *testing.T, status SessionStatus) {
				_, ok := status.(SessionStatusBusy)
				if !ok {
					t.Errorf("expected SessionStatusBusy, got %T", status)
				}
			},
		},
		{
			id:       "ses_retry",
			wantType: "SessionStatusRetry",
			check: func(t *testing.T, status SessionStatus) {
				s, ok := status.(SessionStatusRetry)
				if !ok {
					t.Errorf("expected SessionStatusRetry, got %T", status)
				}
				if s.Attempt != 3 {
					t.Errorf("expected Attempt=3, got %d", s.Attempt)
				}
				if s.Message != "rate limited" {
					t.Errorf("expected Message='rate limited', got %q", s.Message)
				}
				if s.Next != 1700000000 {
					t.Errorf("expected Next=1700000000, got %d", s.Next)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			status, ok := m[tt.id]
			if !ok {
				t.Fatalf("key %q not found in map", tt.id)
			}
			tt.check(t, status)
		})
	}
}

func TestSessionStatusMapUnmarshalInvalid(t *testing.T) {
	t.Parallel()

	// Not a JSON object
	data := `["not", "an", "object"]`
	var m SessionStatusMap
	err := json.Unmarshal([]byte(data), &m)
	if err == nil {
		t.Error("expected error for non-object JSON, got nil")
	}
}

func TestSessionStatusMapUnmarshalEmpty(t *testing.T) {
	t.Parallel()

	data := `{}`
	var m SessionStatusMap
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if len(m) != 0 {
		t.Errorf("expected empty map, got %d entries", len(m))
	}
}
