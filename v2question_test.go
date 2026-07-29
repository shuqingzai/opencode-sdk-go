// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// ===== Prism / live-server tests =====

func TestV2QuestionRequestList(t *testing.T) {
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
	_, err := client.V2Question.Request.List(context.TODO(), opencode.V2QuestionRequestListParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/home/user/project"),
			Workspace: opencode.F("ws_123"),
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

// ===== Request serialization tests =====

// Aligned with OpenAPI GET /api/question/request — query param: location (optional nested object).
func TestV2QuestionRequestListParamsURLQuery(t *testing.T) {
	p := opencode.V2QuestionRequestListParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/work"),
		}),
	}
	got := p.URLQuery()
	// The location param uses bracket notation: location[directory]=...
	encoded := got.Encode()
	if len(encoded) == 0 {
		t.Error("expected non-empty encoded query")
	}
}

func TestV2QuestionRequestListParamsEmptyQuery(t *testing.T) {
	p := opencode.V2QuestionRequestListParams{}
	got := p.URLQuery()
	if len(got) != 0 {
		t.Errorf("expected empty query values, got %v", got)
	}
}

// ===== Response deserialization tests =====

// V2QuestionRequestListResponse: required fields location + data.
func TestV2QuestionRequestListResponseUnmarshal(t *testing.T) {
	raw := `{
		"location": {
			"directory": "/home/user",
			"workspaceID": "ws_1",
			"project": {"id": "proj_1", "directory": "/home/user"}
		},
		"data": [
			{
				"id": "qreq_001",
				"sessionID": "ses_001",
				"questions": [
					{
						"question": "Which model should I use?",
						"header": "Model Choice",
						"options": [
							{"label": "GPT-4", "description": "OpenAI GPT-4"},
							{"label": "Claude", "description": "Anthropic Claude"}
						],
						"multiple": false,
						"custom": false
					}
				],
				"tool": {"messageID": "msg_1", "callID": "call_1"}
			}
		]
	}`
	var resp opencode.V2QuestionRequestListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Location.Directory != "/home/user" {
		t.Errorf("Location.Directory = %q", resp.Location.Directory)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Data))
	}
	item := resp.Data[0]
	if item.ID != "qreq_001" {
		t.Errorf("ID = %q", item.ID)
	}
	if item.SessionID != "ses_001" {
		t.Errorf("SessionID = %q", item.SessionID)
	}
	if len(item.Questions) != 1 {
		t.Fatalf("expected 1 question, got %d", len(item.Questions))
	}
	q := item.Questions[0]
	if q.Question != "Which model should I use?" {
		t.Errorf("Question = %q", q.Question)
	}
	if q.Header != "Model Choice" {
		t.Errorf("Header = %q", q.Header)
	}
	if len(q.Options) != 2 {
		t.Fatalf("expected 2 options, got %d", len(q.Options))
	}
	if q.Options[0].Label != "GPT-4" {
		t.Errorf("Options[0].Label = %q", q.Options[0].Label)
	}
	if q.Options[1].Description != "Anthropic Claude" {
		t.Errorf("Options[1].Description = %q", q.Options[1].Description)
	}
	if item.Tool.MessageID != "msg_1" {
		t.Errorf("Tool.MessageID = %q", item.Tool.MessageID)
	}
	if item.Tool.CallID != "call_1" {
		t.Errorf("Tool.CallID = %q", item.Tool.CallID)
	}
	if item.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved on QuestionV2Request")
	}
	if q.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved on QuestionV2Info")
	}
}

// QuestionV2Request without optional tool field.
func TestQuestionV2RequestWithoutTool(t *testing.T) {
	raw := `{
		"id": "qreq_002",
		"sessionID": "ses_002",
		"questions": [
			{
				"question": "Continue?",
				"header": "Confirm",
				"options": [{"label": "Yes", "description": "Yes please"}],
				"multiple": true,
				"custom": true
			}
		]
	}`
	var r opencode.QuestionV2Request
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.ID != "qreq_002" {
		t.Errorf("ID = %q", r.ID)
	}
	// Tool should be zero-value when absent
	if r.Tool.MessageID != "" || r.Tool.CallID != "" {
		t.Errorf("Tool should be zero-value, got %+v", r.Tool)
	}
	q := r.Questions[0]
	if !q.Multiple {
		t.Error("expected Multiple=true")
	}
	if !q.Custom {
		t.Error("expected Custom=true")
	}
}

// QuestionV2Option: required label + description.
func TestQuestionV2OptionUnmarshal(t *testing.T) {
	raw := `{"label": "Option A", "description": "First option"}`
	var o opencode.QuestionV2Option
	if err := json.Unmarshal([]byte(raw), &o); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if o.Label != "Option A" {
		t.Errorf("Label = %q", o.Label)
	}
	if o.Description != "First option" {
		t.Errorf("Description = %q", o.Description)
	}
	if o.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// Empty data array.
func TestV2QuestionRequestListResponseEmptyData(t *testing.T) {
	raw := `{
		"location": {
			"directory": "/empty",
			"project": {"id": "p0", "directory": "/empty"}
		},
		"data": []
	}`
	var resp opencode.V2QuestionRequestListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp.Data) != 0 {
		t.Errorf("expected 0 items, got %d", len(resp.Data))
	}
}
