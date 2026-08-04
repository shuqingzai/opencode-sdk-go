package opencode

import (
	"encoding/json"
	"testing"
)

// TestQuestionV2AnswerIsAlias verifies that QuestionV2Answer is a type alias
// for []string and serializes identically to a plain []string.
func TestQuestionV2AnswerIsAlias(t *testing.T) {
	// QuestionV2Answer = []string; this should compile and marshal the same way.
	var ans QuestionV2Answer = []string{"option1", "option2"}

	data, err := json.Marshal(ans)
	if err != nil {
		t.Fatalf("json.Marshal error: %v", err)
	}

	var roundtrip []string
	if err := json.Unmarshal(data, &roundtrip); err != nil {
		t.Fatalf("json.Unmarshal error: %v", err)
	}
	if len(roundtrip) != 2 || roundtrip[0] != "option1" || roundtrip[1] != "option2" {
		t.Errorf("roundtrip mismatch: got %v", roundtrip)
	}
}

// TestV2SessionQuestionReplyParamsMarshalJSON verifies the request body is
// serialized correctly with multiple answers.
func TestV2SessionQuestionReplyParamsMarshalJSON(t *testing.T) {
	params := V2SessionQuestionReplyParams{
		Answers: F([]QuestionV2Answer{
			{"yes", "maybe"},
			{"no"},
		}),
	}

	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}
	if data == nil {
		t.Fatal("expected non-nil JSON output")
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	answers, ok := m["answers"].([]any)
	if !ok {
		t.Fatalf("expected answers to be []any, got %T", m["answers"])
	}
	if len(answers) != 2 {
		t.Errorf("expected 2 answers, got %d", len(answers))
	}

	first, ok := answers[0].([]any)
	if !ok || len(first) != 2 {
		t.Errorf("expected first answer to be [yes, maybe], got %v", answers[0])
	}
}

// TestV2SessionQuestionListResponseUnmarshal verifies that
// V2SessionQuestionListResponse deserializes correctly.
func TestV2SessionQuestionListResponseUnmarshal(t *testing.T) {
	raw := `{"data": [
		{
			"id": "que_001",
			"sessionID": "ses_001",
			"questions": [
				{
					"question": "Pick one",
					"header": "Choice",
					"options": [
						{"label": "A", "description": "Option A"},
						{"label": "B", "description": "Option B"}
					]
				}
			]
		}
	]}`

	var resp V2SessionQuestionListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Data))
	}
	if resp.Data[0].ID != "que_001" {
		t.Errorf("expected id=que_001, got %q", resp.Data[0].ID)
	}
	if len(resp.Data[0].Questions) != 1 {
		t.Errorf("expected 1 question, got %d", len(resp.Data[0].Questions))
	}
	if len(resp.Data[0].Questions[0].Options) != 2 {
		t.Errorf("expected 2 options, got %d", len(resp.Data[0].Questions[0].Options))
	}
}
