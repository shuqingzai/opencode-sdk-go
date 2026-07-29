// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// ===== Prism / live-server tests =====

func TestV2SessionQuestionList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.V2Session.Question.List(context.TODO(), "sessionID")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2SessionQuestionReply(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.V2Session.Question.Reply(context.TODO(), "sessionID", "requestID", opencode.V2SessionQuestionReplyParams{
		Answers: opencode.F([]opencode.QuestionV2Answer{
			{"Yes"},
			{"Option A", "Option B"},
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2SessionQuestionReject(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.V2Session.Question.Reject(context.TODO(), "sessionID", "requestID")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// ===== Missing path parameter unit tests =====

func TestV2SessionQuestionListEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionQuestionService()
	_, err := svc.List(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionQuestionReplyEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionQuestionService()
	err := svc.Reply(context.Background(), "", "req_1", opencode.V2SessionQuestionReplyParams{
		Answers: opencode.F([]opencode.QuestionV2Answer{{"yes"}}),
	})
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionQuestionReplyEmptyRequestID(t *testing.T) {
	svc := opencode.NewV2SessionQuestionService()
	err := svc.Reply(context.Background(), "ses_1", "", opencode.V2SessionQuestionReplyParams{
		Answers: opencode.F([]opencode.QuestionV2Answer{{"yes"}}),
	})
	if err == nil {
		t.Fatal("expected error for empty requestID")
	}
	if !strings.Contains(err.Error(), "missing required requestID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionQuestionRejectEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionQuestionService()
	err := svc.Reject(context.Background(), "", "req_1")
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionQuestionRejectEmptyRequestID(t *testing.T) {
	svc := opencode.NewV2SessionQuestionService()
	err := svc.Reject(context.Background(), "ses_1", "")
	if err == nil {
		t.Fatal("expected error for empty requestID")
	}
	if !strings.Contains(err.Error(), "missing required requestID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

// ===== Request serialization tests =====

// V2SessionQuestionReplyParams: required answers (array of arrays of strings).
func TestV2SessionQuestionReplyParamsSerialization(t *testing.T) {
	t.Run("single question single answer", func(t *testing.T) {
		p := opencode.V2SessionQuestionReplyParams{
			Answers: opencode.F([]opencode.QuestionV2Answer{
				{"Yes"},
			}),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(data)
		if !strings.Contains(got, `"answers"`) {
			t.Errorf("answers missing: %s", got)
		}
	})

	t.Run("multiple questions with multiple answers", func(t *testing.T) {
		p := opencode.V2SessionQuestionReplyParams{
			Answers: opencode.F([]opencode.QuestionV2Answer{
				{"Option A", "Option B"},
				{"Yes"},
			}),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(data)
		if !strings.Contains(got, `"answers"`) {
			t.Errorf("answers missing: %s", got)
		}
	})

	t.Run("answers is required", func(t *testing.T) {
		p := opencode.V2SessionQuestionReplyParams{
			Answers: opencode.F([]opencode.QuestionV2Answer{}),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), `"answers"`) {
			t.Errorf("answers required field missing: %s", string(data))
		}
	})
}

// QuestionV2Answer is a type alias for []string; it should marshal correctly.
func TestQuestionV2AnswerMarshal(t *testing.T) {
	a := opencode.QuestionV2Answer{"Option A", "Option B"}
	data, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	got := string(data)
	if !strings.Contains(got, "Option A") || !strings.Contains(got, "Option B") {
		t.Errorf("unexpected marshal: %s", got)
	}
}

// ===== Response deserialization tests =====

// V2SessionQuestionListResponse: required data array.
func TestV2SessionQuestionListResponseUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": "qreq_1",
				"sessionID": "ses_1",
				"questions": [
					{
						"question": "Which approach?",
						"header": "Approach",
						"options": [
							{"label": "Refactor", "description": "Refactor the code"},
							{"label": "Rewrite", "description": "Rewrite from scratch"}
						],
						"multiple": false,
						"custom": false
					}
				],
				"tool": {"messageID": "m1", "callID": "c1"}
			}
		]
	}`
	var resp opencode.V2SessionQuestionListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Data))
	}
	item := resp.Data[0]
	if item.ID != "qreq_1" {
		t.Errorf("ID = %q", item.ID)
	}
	if len(item.Questions) != 1 {
		t.Fatalf("expected 1 question, got %d", len(item.Questions))
	}
	q := item.Questions[0]
	if q.Question != "Which approach?" {
		t.Errorf("Question = %q", q.Question)
	}
	if len(q.Options) != 2 {
		t.Errorf("expected 2 options, got %d", len(q.Options))
	}
	if resp.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// V2SessionQuestionListResponse: empty data is valid.
func TestV2SessionQuestionListResponseEmptyData(t *testing.T) {
	raw := `{"data": []}`
	var resp opencode.V2SessionQuestionListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp.Data) != 0 {
		t.Errorf("expected 0 items, got %d", len(resp.Data))
	}
}

// QuestionV2Request without optional tool.
func TestV2SessionQuestionRequestNoTool(t *testing.T) {
	raw := `{
		"id": "qreq_notool",
		"sessionID": "ses_notool",
		"questions": [
			{
				"question": "Continue?",
				"header": "Confirm",
				"options": [{"label": "Yes", "description": "Yes"}],
				"multiple": true,
				"custom": true
			}
		]
	}`
	var r opencode.QuestionV2Request
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.ID != "qreq_notool" {
		t.Errorf("ID = %q", r.ID)
	}
	if r.Tool.MessageID != "" {
		t.Errorf("Tool.MessageID should be empty, got %q", r.Tool.MessageID)
	}
	if !r.Questions[0].Multiple {
		t.Error("expected Multiple=true")
	}
	if !r.Questions[0].Custom {
		t.Error("expected Custom=true")
	}
}
