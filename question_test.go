// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestQuestionList(t *testing.T) {
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
	_, err := client.Question.List(context.TODO(), opencode.QuestionListParams{
		Directory: opencode.F("directory"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuestionReply(t *testing.T) {
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
	_, err := client.Question.Reply(context.TODO(), "requestID", opencode.QuestionReplyParams{
		Answers: opencode.F([][]string{{"answer1"}}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuestionReject(t *testing.T) {
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
	_, err := client.Question.Reject(context.TODO(), "requestID", opencode.QuestionRejectParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestQuestionToolUnmarshal verifies that QuestionTool (renamed from
// QuestionRequestTool) deserializes correctly from JSON.
func TestQuestionToolUnmarshal(t *testing.T) {
	raw := `{"messageID": "msg_001", "callID": "call_abc"}`
	var tool opencode.QuestionTool
	if err := json.Unmarshal([]byte(raw), &tool); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if tool.MessageID != "msg_001" {
		t.Errorf("expected messageID=msg_001, got %q", tool.MessageID)
	}
	if tool.CallID != "call_abc" {
		t.Errorf("expected callID=call_abc, got %q", tool.CallID)
	}
}

// TestQuestionRequestToolAliasUsable verifies that the backwards-compatibility
// alias QuestionRequestTool still refers to QuestionTool.
func TestQuestionRequestToolAliasUsable(t *testing.T) {
	raw := `{"messageID": "msg_002", "callID": "call_xyz"}`
	var tool opencode.QuestionRequestTool // alias = QuestionTool
	if err := json.Unmarshal([]byte(raw), &tool); err != nil {
		t.Fatalf("Unmarshal error using alias: %v", err)
	}
	if tool.MessageID != "msg_002" {
		t.Errorf("expected messageID=msg_002, got %q", tool.MessageID)
	}
}

// TestQuestionRequestUnmarshal verifies that QuestionRequest (which embeds
// QuestionTool as the tool field) deserializes correctly.
func TestQuestionRequestUnmarshal(t *testing.T) {
	raw := `{
		"id": "que_001",
		"sessionID": "ses_001",
		"questions": [
			{
				"question": "Choose",
				"header": "Hdr",
				"options": [{"label": "Yes", "description": "Affirm"}]
			}
		],
		"tool": {"messageID": "msg_003", "callID": "call_003"}
	}`
	var req opencode.QuestionRequest
	if err := json.Unmarshal([]byte(raw), &req); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if req.ID != "que_001" {
		t.Errorf("expected id=que_001, got %q", req.ID)
	}
	if req.Tool.MessageID != "msg_003" {
		t.Errorf("expected tool.messageID=msg_003, got %q", req.Tool.MessageID)
	}
}
