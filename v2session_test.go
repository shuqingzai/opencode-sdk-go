package opencode

import (
	"context"
	"testing"
)

func TestV2SessionService(t *testing.T) {
	s := NewV2SessionService()
	_ = s
}

func TestV2SessionListParams(t *testing.T) {
	params := V2SessionListParams{
		Limit: Float(10),
		Order: String("desc"),
	}
	v := params.URLQuery()
	if v.Get("limit") != "10" {
		t.Errorf("expected limit=10, got %q", v.Get("limit"))
	}
	if v.Get("order") != "desc" {
		t.Errorf("expected order=desc, got %q", v.Get("order"))
	}
}

func TestV2SessionCompactParams(t *testing.T) {
	params := V2SessionCompactParams{
		Directory: String("/tmp"),
	}
	v := params.URLQuery()
	if v.Get("directory") != "/tmp" {
		t.Errorf("expected directory=/tmp, got %q", v.Get("directory"))
	}
}

func TestV2SessionPromptParams(t *testing.T) {
	params := V2SessionPromptParams{
		Directory: String("/tmp"),
		Body: V2SessionPromptParamsBody{
			Prompt: F(V2Prompt{
				Text: String("hello"),
			}),
			Delivery: String("immediate"),
		},
	}
	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty JSON body")
	}
	v := params.URLQuery()
	if v.Get("directory") != "/tmp" {
		t.Errorf("expected directory=/tmp, got %q", v.Get("directory"))
	}
}

func TestV2SessionsResponseUnmarshal(t *testing.T) {
	jsonStr := `{
		"items": [
			{
				"id": "ses_123",
				"projectID": "proj_1",
				"time": {"created": 1, "updated": 2},
				"title": "test"
			}
		],
		"cursor": {
			"previous": "abc",
			"next": "def"
		}
	}`
	var resp V2SessionsResponse
	if err := resp.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if len(resp.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(resp.Items))
	}
	if resp.Items[0].ID != "ses_123" {
		t.Errorf("expected ses_123, got %s", resp.Items[0].ID)
	}
	if resp.Cursor.Previous != "abc" {
		t.Errorf("expected abc, got %s", resp.Cursor.Previous)
	}
}

func TestV2SessionMessagesResponseUnmarshal(t *testing.T) {
	jsonStr := `{
		"items": [],
		"cursor": {
			"previous": "",
			"next": ""
		}
	}`
	var resp V2SessionMessagesResponse
	if err := resp.UnmarshalJSON([]byte(jsonStr)); err != nil {
		t.Fatal(err)
	}
	if len(resp.Items) != 0 {
		t.Errorf("expected 0 items, got %d", len(resp.Items))
	}
}

func TestV2SessionServiceRequiresSessionID(t *testing.T) {
	s := NewV2SessionService()
	ctx := context.Background()

	_, err := s.Prompt(ctx, "", V2SessionPromptParams{})
	if err == nil {
		t.Error("expected error for empty sessionID")
	}

	err = s.Compact(ctx, "", V2SessionCompactParams{})
	if err == nil {
		t.Error("expected error for empty sessionID")
	}

	err = s.Wait(ctx, "", V2SessionWaitParams{})
	if err == nil {
		t.Error("expected error for empty sessionID")
	}

	_, err = s.Context(ctx, "", V2SessionContextParams{})
	if err == nil {
		t.Error("expected error for empty sessionID")
	}

	_, err = s.Messages(ctx, "", V2SessionMessagesParams{})
	if err == nil {
		t.Error("expected error for empty sessionID")
	}
}
