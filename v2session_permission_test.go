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

func TestV2SessionPermissionList(t *testing.T) {
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
	_, err := client.V2Session.Permission.List(context.TODO(), "sessionID")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2SessionPermissionNew(t *testing.T) {
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
	_, err := client.V2Session.Permission.New(context.TODO(), "sessionID", opencode.V2SessionPermissionNewParams{
		Action:    opencode.F("bash"),
		Resources: opencode.F([]string{"*.sh", "/usr/bin/*"}),
		Save:      opencode.F([]string{"*.sh"}),
		Metadata:  opencode.F[any](map[string]string{"tool": "bash"}),
		Source: opencode.F(opencode.PermissionV2SourceParam{
			Type:      opencode.F(opencode.PermissionV2SourceTypeTool),
			MessageID: opencode.F("msg_001"),
			CallID:    opencode.F("call_001"),
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

func TestV2SessionPermissionGet(t *testing.T) {
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
	_, err := client.V2Session.Permission.Get(context.TODO(), "sessionID", "requestID")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2SessionPermissionReply(t *testing.T) {
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
	err := client.V2Session.Permission.Reply(context.TODO(), "sessionID", "requestID", opencode.V2SessionPermissionReplyParams{
		Reply:   opencode.F(opencode.PermissionV2ReplyOnce),
		Message: opencode.F("approved"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// ===== Missing path parameter unit tests =====

func TestV2SessionPermissionListEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionPermissionService()
	_, err := svc.List(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionPermissionNewEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionPermissionService()
	_, err := svc.New(context.Background(), "", opencode.V2SessionPermissionNewParams{
		Action:    opencode.F("bash"),
		Resources: opencode.F([]string{"*"}),
	})
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionPermissionGetEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionPermissionService()
	_, err := svc.Get(context.Background(), "", "req_1")
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionPermissionGetEmptyRequestID(t *testing.T) {
	svc := opencode.NewV2SessionPermissionService()
	_, err := svc.Get(context.Background(), "ses_1", "")
	if err == nil {
		t.Fatal("expected error for empty requestID")
	}
	if !strings.Contains(err.Error(), "missing required requestID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionPermissionReplyEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionPermissionService()
	err := svc.Reply(context.Background(), "", "req_1", opencode.V2SessionPermissionReplyParams{
		Reply: opencode.F(opencode.PermissionV2ReplyOnce),
	})
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionPermissionReplyEmptyRequestID(t *testing.T) {
	svc := opencode.NewV2SessionPermissionService()
	err := svc.Reply(context.Background(), "ses_1", "", opencode.V2SessionPermissionReplyParams{
		Reply: opencode.F(opencode.PermissionV2ReplyOnce),
	})
	if err == nil {
		t.Fatal("expected error for empty requestID")
	}
	if !strings.Contains(err.Error(), "missing required requestID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

// ===== Request serialization tests =====

// V2SessionPermissionNewParams: required action + resources, optional save/metadata/source/agent.
func TestV2SessionPermissionNewParamsSerialization(t *testing.T) {
	t.Run("required fields only", func(t *testing.T) {
		p := opencode.V2SessionPermissionNewParams{
			Action:    opencode.F("bash"),
			Resources: opencode.F([]string{"*.sh"}),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(data)
		if !strings.Contains(got, `"action":"bash"`) {
			t.Errorf("action missing: %s", got)
		}
		if !strings.Contains(got, `"resources"`) {
			t.Errorf("resources missing: %s", got)
		}
	})

	t.Run("all fields including optional", func(t *testing.T) {
		p := opencode.V2SessionPermissionNewParams{
			Action:    opencode.F("write"),
			Resources: opencode.F([]string{"*.go", "*.ts"}),
			Save:      opencode.F([]string{"*.go"}),
			Metadata:  opencode.F[any](map[string]interface{}{"context": "ai-edit"}),
			Agent:     opencode.F("coder"),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(data)
		if !strings.Contains(got, `"action":"write"`) {
			t.Errorf("action missing: %s", got)
		}
		if !strings.Contains(got, `"save"`) {
			t.Errorf("save missing: %s", got)
		}
		if !strings.Contains(got, `"metadata"`) {
			t.Errorf("metadata missing: %s", got)
		}
		if !strings.Contains(got, `"agent":"coder"`) {
			t.Errorf("agent missing: %s", got)
		}
	})

	t.Run("metadata is any — accepts nested object", func(t *testing.T) {
		p := opencode.V2SessionPermissionNewParams{
			Action:    opencode.F("read"),
			Resources: opencode.F([]string{"*"}),
			Metadata:  opencode.F[any](map[string]interface{}{"nested": map[string]int{"count": 3}}),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), `"metadata"`) {
			t.Errorf("metadata missing: %s", string(data))
		}
	})
}

// V2SessionPermissionReplyParams: required reply, optional message.
func TestV2SessionPermissionReplyParamsSerialization(t *testing.T) {
	t.Run("once reply with message", func(t *testing.T) {
		p := opencode.V2SessionPermissionReplyParams{
			Reply:   opencode.F(opencode.PermissionV2ReplyOnce),
			Message: opencode.F("approved for this session"),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(data)
		if !strings.Contains(got, `"reply":"once"`) {
			t.Errorf("reply missing: %s", got)
		}
		if !strings.Contains(got, `"message":"approved for this session"`) {
			t.Errorf("message missing: %s", got)
		}
	})

	t.Run("always reply without message", func(t *testing.T) {
		p := opencode.V2SessionPermissionReplyParams{
			Reply: opencode.F(opencode.PermissionV2ReplyAlways),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), `"reply":"always"`) {
			t.Errorf("reply missing: %s", string(data))
		}
	})

	t.Run("reject reply", func(t *testing.T) {
		p := opencode.V2SessionPermissionReplyParams{
			Reply: opencode.F(opencode.PermissionV2ReplyReject),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), `"reply":"reject"`) {
			t.Errorf("reject missing: %s", string(data))
		}
	})
}

// PermissionV2Reply enum IsKnown.
func TestPermissionV2ReplyIsKnown(t *testing.T) {
	for _, v := range []opencode.PermissionV2Reply{
		opencode.PermissionV2ReplyOnce,
		opencode.PermissionV2ReplyAlways,
		opencode.PermissionV2ReplyReject,
	} {
		if !v.IsKnown() {
			t.Errorf("%q should be known", v)
		}
	}
	if opencode.PermissionV2Reply("unknown").IsKnown() {
		t.Error("unknown reply should not be known")
	}
}

// ===== Response deserialization tests =====

// V2SessionPermissionListResponse: required data array.
func TestV2SessionPermissionListResponseUnmarshal(t *testing.T) {
	raw := `{
		"data": [
			{
				"id": "req_1",
				"sessionID": "ses_1",
				"action": "bash",
				"resources": ["*.sh"],
				"source": {"type": "tool", "messageID": "msg_1", "callID": "c1"}
			}
		]
	}`
	var resp opencode.V2SessionPermissionListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Data))
	}
	if resp.Data[0].ID != "req_1" {
		t.Errorf("ID = %q", resp.Data[0].ID)
	}
	if resp.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// V2SessionPermissionCreateResponse: data with effect.
func TestV2SessionPermissionCreateResponseUnmarshal(t *testing.T) {
	raw := `{
		"data": {
			"id": "req_new",
			"effect": "ask"
		}
	}`
	var resp opencode.V2SessionPermissionCreateResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Data.ID != "req_new" {
		t.Errorf("Data.ID = %q", resp.Data.ID)
	}
	if resp.Data.Effect != opencode.PermissionV2EffectAsk {
		t.Errorf("Data.Effect = %q", resp.Data.Effect)
	}
	if resp.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// V2SessionPermissionGetResponse: single PermissionV2Request.
func TestV2SessionPermissionGetResponseUnmarshal(t *testing.T) {
	raw := `{
		"data": {
			"id": "req_get",
			"sessionID": "ses_get",
			"action": "write",
			"resources": ["*.md"],
			"save": [],
			"metadata": null
		}
	}`
	var resp opencode.V2SessionPermissionGetResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Data.ID != "req_get" {
		t.Errorf("Data.ID = %q", resp.Data.ID)
	}
	if resp.Data.Action != "write" {
		t.Errorf("Data.Action = %q", resp.Data.Action)
	}
}
