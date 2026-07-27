package opencode

import (
	"encoding/json"
	"strings"
	"testing"
)

// Aligned with OpenAPI GET /question + JS SDK(v2) Question.list.
// query-only params: directory, workspace
func TestQuestionListParamsQuery(t *testing.T) {
	p := QuestionListParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI POST /question/{requestID}/reply + JS SDK(v2) Question.reply.
// body required: answers; query: directory, workspace
func TestQuestionReplyParamsBodyAndQuery(t *testing.T) {
	t.Run("required answers field serialized to body", func(t *testing.T) {
		p := QuestionReplyParams{
			Answers:   F([]QuestionAnswer{{"answer1", "answer2"}, {"yes"}}),
			Directory: F("d"),
			Workspace: F("w"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"answers"`) {
			t.Errorf("answers missing from body: %s", got)
		}
		if !strings.Contains(got, `"answer1"`) {
			t.Errorf("answer1 missing from body: %s", got)
		}
		if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
			t.Errorf("query fields leaked into body: %s", got)
		}
	})

	t.Run("empty answers array", func(t *testing.T) {
		p := QuestionReplyParams{
			Answers: F([]QuestionAnswer{}),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"answers":[]`) {
			t.Errorf("empty answers missing from body: %s", got)
		}
	})

	t.Run("query serialization", func(t *testing.T) {
		p := QuestionReplyParams{
			Answers:   F([]QuestionAnswer{{"yes"}}),
			Directory: F("mydir"),
			Workspace: F("mywsp"),
		}
		got := p.URLQuery().Encode()
		want := "directory=mydir&workspace=mywsp"
		if got != want {
			t.Errorf("query got %q, want %q", got, want)
		}
	})
}

// Aligned with OpenAPI POST /question/{requestID}/reject + JS SDK(v2) Question.reject.
// query-only params: directory, workspace
func TestQuestionRejectParamsQuery(t *testing.T) {
	p := QuestionRejectParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI QuestionRequest schema.
// required: id, sessionID, questions; optional: tool
func TestQuestionRequestUnmarshal(t *testing.T) {
	t.Run("full object with tool", func(t *testing.T) {
		raw := `{
			"id": "que_abc123",
			"sessionID": "ses_xyz",
			"questions": [
				{
					"question": "Which approach do you prefer?",
					"header": "Approach",
					"options": [
						{"label": "Option A", "description": "First choice"},
						{"label": "Option B", "description": "Second choice"}
					],
					"multiple": false,
					"custom": true
				}
			],
			"tool": {"messageID": "msg_001", "callID": "call_001"}
		}`
		var r QuestionRequest
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if r.ID != "que_abc123" {
			t.Errorf("ID = %q", r.ID)
		}
		if r.SessionID != "ses_xyz" {
			t.Errorf("SessionID = %q", r.SessionID)
		}
		if len(r.Questions) != 1 {
			t.Fatalf("Questions len = %d", len(r.Questions))
		}
		q := r.Questions[0]
		if q.Question != "Which approach do you prefer?" {
			t.Errorf("Question = %q", q.Question)
		}
		if q.Header != "Approach" {
			t.Errorf("Header = %q", q.Header)
		}
		if len(q.Options) != 2 {
			t.Errorf("Options len = %d", len(q.Options))
		}
		if q.Options[0].Label != "Option A" || q.Options[0].Description != "First choice" {
			t.Errorf("Option[0] = %+v", q.Options[0])
		}
		if q.Multiple {
			t.Error("Multiple should be false")
		}
		if !q.Custom {
			t.Error("Custom should be true")
		}
		if r.Tool.MessageID != "msg_001" || r.Tool.CallID != "call_001" {
			t.Errorf("Tool = %+v", r.Tool)
		}
		if r.JSON.raw == "" {
			t.Error("RawJSON not preserved")
		}
	})

	t.Run("without optional tool field", func(t *testing.T) {
		raw := `{
			"id": "que_000",
			"sessionID": "ses_000",
			"questions": [
				{
					"question": "Simple question?",
					"header": "Q",
					"options": [{"label": "Yes", "description": "Affirmative"}]
				}
			]
		}`
		var r QuestionRequest
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if r.Tool.MessageID != "" || r.Tool.CallID != "" {
			t.Errorf("Tool should be zero-value when absent, got %+v", r.Tool)
		}
		// multiple/custom should be false by default
		if r.Questions[0].Multiple {
			t.Error("Multiple should default to false")
		}
		if r.Questions[0].Custom {
			t.Error("Custom should default to false")
		}
	})

	t.Run("multiple questions", func(t *testing.T) {
		raw := `{
			"id": "que_multi",
			"sessionID": "ses_multi",
			"questions": [
				{"question": "Q1", "header": "H1", "options": [{"label": "A", "description": "First"}]},
				{"question": "Q2", "header": "H2", "options": [{"label": "B", "description": "Second"}, {"label": "C", "description": "Third"}], "multiple": true}
			]
		}`
		var r QuestionRequest
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(r.Questions) != 2 {
			t.Fatalf("Questions len = %d", len(r.Questions))
		}
		if !r.Questions[1].Multiple {
			t.Error("Q2.Multiple should be true")
		}
		if len(r.Questions[1].Options) != 2 {
			t.Errorf("Q2 options len = %d", len(r.Questions[1].Options))
		}
	})
}

// Aligned with OpenAPI QuestionInfo schema.
// required: question, header, options; optional: multiple, custom
func TestQuestionInfoUnmarshal(t *testing.T) {
	raw := `{"question":"What now?","header":"Next","options":[{"label":"Go","description":"Proceed"}],"multiple":true,"custom":false}`
	var qi QuestionInfo
	if err := json.Unmarshal([]byte(raw), &qi); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if qi.Question != "What now?" {
		t.Errorf("Question = %q", qi.Question)
	}
	if qi.Header != "Next" {
		t.Errorf("Header = %q", qi.Header)
	}
	if len(qi.Options) != 1 || qi.Options[0].Label != "Go" {
		t.Errorf("Options = %v", qi.Options)
	}
	if !qi.Multiple {
		t.Error("Multiple should be true")
	}
	if qi.Custom {
		t.Error("Custom should be false")
	}
	if qi.JSON.raw == "" {
		t.Error("RawJSON not preserved")
	}
}

// Aligned with OpenAPI QuestionOption schema.
// required: label, description
func TestQuestionOptionUnmarshal(t *testing.T) {
	raw := `{"label":"Confirm","description":"Yes, proceed with the operation"}`
	var o QuestionOption
	if err := json.Unmarshal([]byte(raw), &o); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if o.Label != "Confirm" {
		t.Errorf("Label = %q", o.Label)
	}
	if o.Description != "Yes, proceed with the operation" {
		t.Errorf("Description = %q", o.Description)
	}
	if o.JSON.raw == "" {
		t.Error("RawJSON not preserved")
	}
}

// Aligned with OpenAPI QuestionTool schema (Go type: QuestionRequestTool).
// required: messageID, callID
func TestQuestionRequestToolUnmarshal(t *testing.T) {
	raw := `{"messageID": "msg_q1", "callID": "call_q1"}`
	var t2 QuestionRequestTool
	if err := json.Unmarshal([]byte(raw), &t2); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if t2.MessageID != "msg_q1" {
		t.Errorf("MessageID = %q", t2.MessageID)
	}
	if t2.CallID != "call_q1" {
		t.Errorf("CallID = %q", t2.CallID)
	}
	if t2.JSON.raw == "" {
		t.Error("RawJSON not preserved")
	}
}

// Aligned with OpenAPI QuestionAnswer = Array<string>.
func TestQuestionAnswerType(t *testing.T) {
	// QuestionAnswer is a type alias for []string
	var a QuestionAnswer = []string{"choice1", "choice2"}
	b, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)
	want := `["choice1","choice2"]`
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	var a2 QuestionAnswer
	if err := json.Unmarshal([]byte(`["yes","no"]`), &a2); err != nil {
		t.Fatal(err)
	}
	if len(a2) != 2 || a2[0] != "yes" || a2[1] != "no" {
		t.Errorf("QuestionAnswer = %v", a2)
	}
}
