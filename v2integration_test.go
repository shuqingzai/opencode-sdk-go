// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// ===== Prism / live-server tests =====

func TestV2IntegrationList(t *testing.T) {
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
	_, err := client.V2Integration.List(context.TODO(), opencode.V2IntegrationListParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/home/user/project"),
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

func TestV2IntegrationGet(t *testing.T) {
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
	_, err := client.V2Integration.Get(context.TODO(), "github", opencode.V2IntegrationGetParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("/home/user/project"),
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

func TestV2IntegrationConnectKey(t *testing.T) {
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
	err := client.V2Integration.Connect.Key(context.TODO(), "github", opencode.V2IntegrationConnectKeyParams{
		Body: opencode.F(opencode.V2IntegrationConnectKeyParamsBody{
			Key:   opencode.F("ghp_secret_key"),
			Label: opencode.F("My GitHub Token"),
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

func TestV2IntegrationConnectOauth(t *testing.T) {
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
	_, err := client.V2Integration.Connect.Oauth(context.TODO(), "github", opencode.V2IntegrationConnectOauthParams{
		Body: opencode.F(opencode.V2IntegrationConnectOauthParamsBody{
			MethodID: opencode.F("method_oauth"),
			Inputs:   opencode.F(map[string]string{"enterpriseUrl": "https://github.example.com"}),
			Label:    opencode.F("Optional Label"),
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

func TestV2IntegrationAttemptStatus(t *testing.T) {
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
	_, err := client.V2Integration.Attempt.Status(context.TODO(), "attempt_xyz", opencode.V2IntegrationAttemptStatusParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2IntegrationAttemptComplete(t *testing.T) {
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
	err := client.V2Integration.Attempt.Complete(context.TODO(), "attempt_xyz", opencode.V2IntegrationAttemptCompleteParams{
		Body: opencode.F(opencode.V2IntegrationAttemptCompleteParamsBody{
			Code: opencode.F("auth-code-from-redirect"),
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

func TestV2IntegrationAttemptCancel(t *testing.T) {
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
	err := client.V2Integration.Attempt.Cancel(context.TODO(), "attempt_xyz", opencode.V2IntegrationAttemptCancelParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// ===== Missing path parameter unit tests =====

func TestV2IntegrationGetEmptyID(t *testing.T) {
	svc := opencode.NewV2IntegrationService()
	_, err := svc.Get(context.Background(), "", opencode.V2IntegrationGetParams{})
	if err == nil {
		t.Fatal("expected error for empty integrationID")
	}
	if !strings.Contains(err.Error(), "missing required integrationID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2IntegrationConnectKeyEmptyID(t *testing.T) {
	svc := opencode.NewV2IntegrationConnectService()
	err := svc.Key(context.Background(), "", opencode.V2IntegrationConnectKeyParams{
		Body: opencode.F(opencode.V2IntegrationConnectKeyParamsBody{Key: opencode.F("k")}),
	})
	if err == nil {
		t.Fatal("expected error for empty integrationID")
	}
	if !strings.Contains(err.Error(), "missing required integrationID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2IntegrationConnectOauthEmptyID(t *testing.T) {
	svc := opencode.NewV2IntegrationConnectService()
	_, err := svc.Oauth(context.Background(), "", opencode.V2IntegrationConnectOauthParams{
		Body: opencode.F(opencode.V2IntegrationConnectOauthParamsBody{
			MethodID: opencode.F("m"),
			Inputs:   opencode.F(map[string]string{}),
		}),
	})
	if err == nil {
		t.Fatal("expected error for empty integrationID")
	}
	if !strings.Contains(err.Error(), "missing required integrationID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2IntegrationAttemptStatusEmptyID(t *testing.T) {
	svc := opencode.NewV2IntegrationAttemptService()
	_, err := svc.Status(context.Background(), "", opencode.V2IntegrationAttemptStatusParams{})
	if err == nil {
		t.Fatal("expected error for empty attemptID")
	}
	if !strings.Contains(err.Error(), "missing required attemptID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2IntegrationAttemptCompleteEmptyID(t *testing.T) {
	svc := opencode.NewV2IntegrationAttemptService()
	err := svc.Complete(context.Background(), "", opencode.V2IntegrationAttemptCompleteParams{})
	if err == nil {
		t.Fatal("expected error for empty attemptID")
	}
	if !strings.Contains(err.Error(), "missing required attemptID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2IntegrationAttemptCancelEmptyID(t *testing.T) {
	svc := opencode.NewV2IntegrationAttemptService()
	err := svc.Cancel(context.Background(), "", opencode.V2IntegrationAttemptCancelParams{})
	if err == nil {
		t.Fatal("expected error for empty attemptID")
	}
	if !strings.Contains(err.Error(), "missing required attemptID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

// ===== Request serialization tests =====

// V2IntegrationConnectKeyParamsBody: required key, optional label.
func TestV2IntegrationConnectKeyParamsBodySerialization(t *testing.T) {
	t.Run("key required present", func(t *testing.T) {
		b := opencode.V2IntegrationConnectKeyParamsBody{
			Key:   opencode.F("my-secret-key"),
			Label: opencode.F("My Label"),
		}
		data, err := json.Marshal(b)
		if err != nil {
			t.Fatal(err)
		}
		got := string(data)
		if !strings.Contains(got, `"key":"my-secret-key"`) {
			t.Errorf("key missing: %s", got)
		}
		if !strings.Contains(got, `"label":"My Label"`) {
			t.Errorf("label missing: %s", got)
		}
	})

	t.Run("key required without optional label", func(t *testing.T) {
		b := opencode.V2IntegrationConnectKeyParamsBody{
			Key: opencode.F("sk"),
		}
		data, err := json.Marshal(b)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), `"key":"sk"`) {
			t.Errorf("key missing: %s", string(data))
		}
	})
}

// V2IntegrationConnectOauthParamsBody: required methodID + inputs, optional label.
func TestV2IntegrationConnectOauthParamsBodySerialization(t *testing.T) {
	b := opencode.V2IntegrationConnectOauthParamsBody{
		MethodID: opencode.F("oauth_method_id"),
		Inputs:   opencode.F(map[string]string{"key1": "val1"}),
		Label:    opencode.F("GitHub Connection"),
	}
	data, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	got := string(data)
	if !strings.Contains(got, `"methodID":"oauth_method_id"`) {
		t.Errorf("methodID missing: %s", got)
	}
	if !strings.Contains(got, `"inputs"`) {
		t.Errorf("inputs missing: %s", got)
	}
	if !strings.Contains(got, `"label":"GitHub Connection"`) {
		t.Errorf("label missing: %s", got)
	}
}

// V2IntegrationConnectKeyParams: body marshaled via inner Body, query via URLQuery.
func TestV2IntegrationConnectKeyParamsMarshal(t *testing.T) {
	p := opencode.V2IntegrationConnectKeyParams{
		Location: opencode.F(opencode.V2LocationParam{Directory: opencode.F("/proj")}),
		Body: opencode.F(opencode.V2IntegrationConnectKeyParamsBody{
			Key:   opencode.F("secret"),
			Label: opencode.F("label"),
		}),
	}
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	// The outer MarshalJSON delegates to Body
	if !strings.Contains(string(data), `"key":"secret"`) {
		t.Errorf("body.key missing: %s", string(data))
	}
	// location should not appear in body
	if strings.Contains(string(data), "location") {
		t.Errorf("location leaked into body: %s", string(data))
	}
}

// ===== Response deserialization tests =====

// IntegrationInfo: Methods union — oauth/key/env variants.
func TestIntegrationInfoMethodsUnion(t *testing.T) {
	raw := `{
		"id": "github",
		"name": "GitHub",
		"methods": [
			{"id": "oauth_gh", "type": "oauth", "label": "OAuth", "prompts": []},
			{"type": "key", "label": "API Key"},
			{"type": "env", "names": ["GITHUB_TOKEN"]}
		],
		"connections": []
	}`
	var info opencode.IntegrationInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if info.ID != "github" {
		t.Errorf("ID = %q", info.ID)
	}
	if info.Name != "GitHub" {
		t.Errorf("Name = %q", info.Name)
	}
	if info.Methods == nil {
		t.Error("Methods should not be nil")
	}
	methods := info.Methods
	wantTypes := []any{
		opencode.IntegrationOAuthMethod{},
		opencode.IntegrationKeyMethod{},
		opencode.IntegrationEnvMethod{},
	}
	if len(methods) != len(wantTypes) {
		t.Fatalf("len(Methods) = %d, want %d", len(methods), len(wantTypes))
	}
	for i, want := range wantTypes {
		if reflect.TypeOf(methods[i]) != reflect.TypeOf(want) {
			t.Errorf("Methods[%d] = %s, want %s", i, reflect.TypeOf(methods[i]), reflect.TypeOf(want))
		}
	}
	if got := info.AsMethodsUnion(); len(got) != len(wantTypes) {
		t.Errorf("AsMethodsUnion() len = %d, want %d", len(got), len(wantTypes))
	}
	if info.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// IntegrationInfo: Connections union — credential/env variants.
func TestIntegrationInfoConnectionsUnion(t *testing.T) {
	raw := `{
		"id": "linear",
		"name": "Linear",
		"methods": [],
		"connections": [
			{"type": "credential", "id": "cred_1", "label": "My Token"},
			{"type": "env", "name": "LINEAR_API_KEY"}
		]
	}`
	var info opencode.IntegrationInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if info.Connections == nil {
		t.Error("Connections should not be nil")
	}
	conns := info.Connections
	if len(conns) != 2 {
		t.Fatalf("len(Connections) = %d, want 2", len(conns))
	}
	cred, ok := conns[0].(opencode.ConnectionCredentialInfo)
	if !ok {
		t.Fatalf("Connections[0] = %s, want ConnectionCredentialInfo", reflect.TypeOf(conns[0]))
	}
	if cred.ID != "cred_1" || cred.Label != "My Token" {
		t.Errorf("Connections[0] = %+v", cred)
	}
	env, ok := conns[1].(opencode.ConnectionEnvInfo)
	if !ok {
		t.Fatalf("Connections[1] = %s, want ConnectionEnvInfo", reflect.TypeOf(conns[1]))
	}
	if env.Name != "LINEAR_API_KEY" {
		t.Errorf("Connections[1] = %+v", env)
	}
	if got := info.AsConnectionsUnion(); len(got) != 2 {
		t.Errorf("AsConnectionsUnion() len = %d, want 2", len(got))
	}
}

// IntegrationAttemptStatus union: pending variant.
func TestIntegrationAttemptStatusPendingUnion(t *testing.T) {
	raw := `{
		"status": "pending",
		"time": {"created": 1700000000, "expires": 1700003600}
	}`
	var s opencode.IntegrationAttemptStatus
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if s.Status != opencode.IntegrationAttemptStatusTypePending {
		t.Errorf("Status = %q", s.Status)
	}
	u := s.AsUnion()
	if u == nil {
		t.Fatal("AsUnion() returned nil")
	}
	pending, ok := u.(opencode.IntegrationAttemptStatusPending)
	if !ok {
		t.Fatalf("expected IntegrationAttemptStatusPending, got %T", u)
	}
	if pending.Status != opencode.IntegrationAttemptStatusTypePending {
		t.Errorf("pending.Status = %q", pending.Status)
	}
}

// IntegrationAttemptStatus union: complete variant.
func TestIntegrationAttemptStatusCompleteUnion(t *testing.T) {
	raw := `{
		"status": "complete",
		"time": {"created": 1700000000, "expires": 1700003600}
	}`
	var s opencode.IntegrationAttemptStatus
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if s.Status != opencode.IntegrationAttemptStatusTypeComplete {
		t.Errorf("Status = %q", s.Status)
	}
	u := s.AsUnion()
	if u == nil {
		t.Fatal("AsUnion() returned nil")
	}
	complete, ok := u.(opencode.IntegrationAttemptStatusComplete)
	if !ok {
		t.Fatalf("expected IntegrationAttemptStatusComplete, got %T", u)
	}
	if complete.Status != opencode.IntegrationAttemptStatusTypeComplete {
		t.Errorf("complete.Status = %q", complete.Status)
	}
}

// IntegrationAttemptStatus union: failed variant with message.
func TestIntegrationAttemptStatusFailedUnion(t *testing.T) {
	raw := `{
		"status": "failed",
		"message": "User denied access",
		"time": {"created": 1700000000, "expires": 1700003600}
	}`
	var s opencode.IntegrationAttemptStatus
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if s.Status != opencode.IntegrationAttemptStatusTypeFailed {
		t.Errorf("Status = %q", s.Status)
	}
	u := s.AsUnion()
	failed, ok := u.(opencode.IntegrationAttemptStatusFailed)
	if !ok {
		t.Fatalf("expected IntegrationAttemptStatusFailed, got %T", u)
	}
	if failed.Message != "User denied access" {
		t.Errorf("Message = %q", failed.Message)
	}
}

// IntegrationAttemptStatus union: expired variant.
func TestIntegrationAttemptStatusExpiredUnion(t *testing.T) {
	raw := `{
		"status": "expired",
		"time": {"created": 1700000000, "expires": 1700003600}
	}`
	var s opencode.IntegrationAttemptStatus
	if err := json.Unmarshal([]byte(raw), &s); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if s.Status != opencode.IntegrationAttemptStatusTypeExpired {
		t.Errorf("Status = %q", s.Status)
	}
	u := s.AsUnion()
	if u == nil {
		t.Fatal("AsUnion() returned nil")
	}
	expired, ok := u.(opencode.IntegrationAttemptStatusExpired)
	if !ok {
		t.Fatalf("expected IntegrationAttemptStatusExpired, got %T", u)
	}
	if expired.Status != opencode.IntegrationAttemptStatusTypeExpired {
		t.Errorf("expired.Status = %q", expired.Status)
	}
}

// IntegrationAttemptStatusType enum IsKnown.
func TestIntegrationAttemptStatusTypeIsKnown(t *testing.T) {
	for _, v := range []opencode.IntegrationAttemptStatusType{
		opencode.IntegrationAttemptStatusTypePending,
		opencode.IntegrationAttemptStatusTypeComplete,
		opencode.IntegrationAttemptStatusTypeFailed,
		opencode.IntegrationAttemptStatusTypeExpired,
	} {
		if !v.IsKnown() {
			t.Errorf("%q should be known", v)
		}
	}
	if opencode.IntegrationAttemptStatusType("unknown").IsKnown() {
		t.Error("unknown status should not be known")
	}
}

// IntegrationAttemptMode enum IsKnown.
func TestIntegrationAttemptModeIsKnown(t *testing.T) {
	for _, v := range []opencode.IntegrationAttemptMode{
		opencode.IntegrationAttemptModeAuto,
		opencode.IntegrationAttemptModeCode,
	} {
		if !v.IsKnown() {
			t.Errorf("%q should be known", v)
		}
	}
}

// IntegrationOAuthMethod prompt union: text + select variants.
func TestIntegrationOAuthMethodPromptsUnion(t *testing.T) {
	raw := `{
		"id": "oauth_gh",
		"type": "oauth",
		"label": "OAuth",
		"prompts": [
			{
				"type": "text",
				"key": "enterpriseUrl",
				"message": "GitHub Enterprise URL",
				"placeholder": "https://github.example.com"
			},
			{
				"type": "select",
				"key": "region",
				"message": "Select region",
				"options": [
					{"label": "US", "value": "us", "hint": "United States"},
					{"label": "EU", "value": "eu"}
				]
			}
		]
	}`
	var m opencode.IntegrationOAuthMethod
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.ID != "oauth_gh" {
		t.Errorf("ID = %q", m.ID)
	}
	if m.Type != opencode.IntegrationOAuthMethodTypeOAuth {
		t.Errorf("Type = %q", m.Type)
	}
	if m.Prompts == nil {
		t.Error("Prompts should not be nil")
	}
	prompts := m.Prompts
	if len(prompts) != 2 {
		t.Fatalf("len(Prompts) = %d, want 2", len(prompts))
	}
	text, ok := prompts[0].(opencode.IntegrationTextPrompt)
	if !ok {
		t.Fatalf("Prompts[0] = %s, want IntegrationTextPrompt", reflect.TypeOf(prompts[0]))
	}
	if text.Key != "enterpriseUrl" || text.Placeholder != "https://github.example.com" {
		t.Errorf("Prompts[0] = %+v", text)
	}
	sel, ok := prompts[1].(opencode.IntegrationSelectPrompt)
	if !ok {
		t.Fatalf("Prompts[1] = %s, want IntegrationSelectPrompt", reflect.TypeOf(prompts[1]))
	}
	if sel.Key != "region" || len(sel.Options) != 2 || sel.Options[0].Hint != "United States" {
		t.Errorf("Prompts[1] = %+v", sel)
	}
	if m.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// ConnectionCredentialInfo deserialization.
func TestConnectionCredentialInfoUnmarshal(t *testing.T) {
	raw := `{"type": "credential", "id": "cred_1", "label": "My GitHub Token"}`
	var c opencode.ConnectionCredentialInfo
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if c.Type != opencode.ConnectionCredentialInfoTypeCredential {
		t.Errorf("Type = %q", c.Type)
	}
	if c.ID != "cred_1" {
		t.Errorf("ID = %q", c.ID)
	}
	if c.Label != "My GitHub Token" {
		t.Errorf("Label = %q", c.Label)
	}
	if c.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// ConnectionEnvInfo deserialization.
func TestConnectionEnvInfoUnmarshal(t *testing.T) {
	raw := `{"type": "env", "name": "GITHUB_TOKEN"}`
	var c opencode.ConnectionEnvInfo
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if c.Type != opencode.ConnectionEnvInfoTypeEnv {
		t.Errorf("Type = %q", c.Type)
	}
	if c.Name != "GITHUB_TOKEN" {
		t.Errorf("Name = %q", c.Name)
	}
}

// V2IntegrationConnectOauthResponse: IntegrationAttempt deserialization.
func TestV2IntegrationConnectOauthResponseUnmarshal(t *testing.T) {
	raw := `{
		"location": {
			"directory": "/home/user",
			"project": {"id": "p1", "directory": "/home/user"}
		},
		"data": {
			"attemptID": "att_abc",
			"url": "https://github.com/login/oauth/authorize?client_id=xxx",
			"instructions": "Click the URL to authenticate",
			"mode": "auto",
			"time": {"created": 1700000000, "expires": 1700003600}
		}
	}`
	var resp opencode.V2IntegrationConnectOauthResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Data.AttemptID != "att_abc" {
		t.Errorf("AttemptID = %q", resp.Data.AttemptID)
	}
	if resp.Data.Mode != opencode.IntegrationAttemptModeAuto {
		t.Errorf("Mode = %q", resp.Data.Mode)
	}
	if resp.Data.Instructions == "" {
		t.Error("Instructions should not be empty")
	}
	if resp.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// IntegrationWhen: conditional visibility rule.
func TestIntegrationWhenUnmarshal(t *testing.T) {
	raw := `{"key": "authType", "op": "eq", "value": "oauth"}`
	var w opencode.IntegrationWhen
	if err := json.Unmarshal([]byte(raw), &w); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if w.Key != "authType" {
		t.Errorf("Key = %q", w.Key)
	}
	if w.Op != opencode.IntegrationWhenOpEq {
		t.Errorf("Op = %q", w.Op)
	}
	if w.Value != "oauth" {
		t.Errorf("Value = %q", w.Value)
	}
}

func TestIntegrationWhenOpIsKnown(t *testing.T) {
	if !opencode.IntegrationWhenOpEq.IsKnown() {
		t.Error("eq should be known")
	}
	if !opencode.IntegrationWhenOpNeq.IsKnown() {
		t.Error("neq should be known")
	}
	if opencode.IntegrationWhenOp("gt").IsKnown() {
		t.Error("gt should not be known")
	}
}

// ===== Union discriminator routing tests =====
//
// The three unions below are registered with discriminatorKey="type" plus a
// DiscriminatorValue on every variant, mirroring IntegrationAttemptStatusUnion
// (discriminatorKey="status"). OpenAPI defines each variant with a required
// `type` property whose enum holds exactly one value, and those values are
// unique inside each union:
//
//	IntegrationMethod  (anyOf) -> IntegrationOAuthMethod   type=oauth
//	                              IntegrationKeyMethod     type=key
//	                              IntegrationEnvMethod     type=env
//	ConnectionInfo     (anyOf) -> ConnectionCredentialInfo type=credential
//	                              ConnectionEnvInfo        type=env
//	IntegrationPrompt  (anyOf) -> IntegrationTextPrompt    type=text
//	                              IntegrationSelectPrompt  type=select
//
// Regression guard: a union must never be registered with a discriminatorKey
// while leaving DiscriminatorValue unset on its variants. internal/apijson
// decoder.go compares `n.Get(key).Value() == variant.DiscriminatorValue`; an
// unset DiscriminatorValue is a nil `any`, so payloads that omit the
// discriminator key (Value() == nil) would match the first such variant and be
// silently mis-routed, while payloads that carry it would never match at all.
// This is the bug previously found in McpStatusUnion (see mcp_test.go
// TestMcpStatusUnmarshal). Both halves must always be set together.
//
// DiscriminatorValue must be an untyped string constant ("oauth"), never a
// typed enum constant such as IntegrationOAuthMethodTypeOAuth: the comparison
// is an interface equality check, so a typed constant would carry a different
// dynamic type than the plain string returned by gjson and never compare equal.

// IntegrationMethodUnion: every variant routes via discriminatorKey="type".
func TestIntegrationMethodUnionDiscriminator(t *testing.T) {
	t.Parallel()

	t.Run("oauth routes to IntegrationOAuthMethod", func(t *testing.T) {
		raw := `{"id":"oauth_gh","type":"oauth","label":"GitHub OAuth"}`
		var u opencode.IntegrationMethodUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		m, ok := u.(opencode.IntegrationOAuthMethod)
		if !ok {
			t.Fatalf("expected IntegrationOAuthMethod, got %T", u)
		}
		if m.Type != opencode.IntegrationOAuthMethodTypeOAuth {
			t.Errorf("Type = %q, want %q", m.Type, opencode.IntegrationOAuthMethodTypeOAuth)
		}
		if m.ID != "oauth_gh" {
			t.Errorf("ID = %q, want %q", m.ID, "oauth_gh")
		}
		if m.Label != "GitHub OAuth" {
			t.Errorf("Label = %q, want %q", m.Label, "GitHub OAuth")
		}
	})

	t.Run("key routes to IntegrationKeyMethod", func(t *testing.T) {
		raw := `{"type":"key","label":"API Key"}`
		var u opencode.IntegrationMethodUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		m, ok := u.(opencode.IntegrationKeyMethod)
		if !ok {
			t.Fatalf("expected IntegrationKeyMethod, got %T", u)
		}
		if m.Type != opencode.IntegrationKeyMethodTypeKey {
			t.Errorf("Type = %q, want %q", m.Type, opencode.IntegrationKeyMethodTypeKey)
		}
		if m.Label != "API Key" {
			t.Errorf("Label = %q, want %q", m.Label, "API Key")
		}
	})

	// OpenAPI marks `label` optional on IntegrationKeyMethod, so the bare form
	// must still route on the discriminator alone.
	t.Run("key without optional label still routes", func(t *testing.T) {
		raw := `{"type":"key"}`
		var u opencode.IntegrationMethodUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		m, ok := u.(opencode.IntegrationKeyMethod)
		if !ok {
			t.Fatalf("expected IntegrationKeyMethod, got %T", u)
		}
		if m.Label != "" {
			t.Errorf("Label = %q, want empty", m.Label)
		}
	})

	t.Run("env routes to IntegrationEnvMethod", func(t *testing.T) {
		raw := `{"type":"env","names":["GITHUB_TOKEN","GH_TOKEN"]}`
		var u opencode.IntegrationMethodUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		m, ok := u.(opencode.IntegrationEnvMethod)
		if !ok {
			t.Fatalf("expected IntegrationEnvMethod, got %T", u)
		}
		if m.Type != opencode.IntegrationEnvMethodTypeEnv {
			t.Errorf("Type = %q, want %q", m.Type, opencode.IntegrationEnvMethodTypeEnv)
		}
		if len(m.Names) != 2 || m.Names[0] != "GITHUB_TOKEN" || m.Names[1] != "GH_TOKEN" {
			t.Errorf("Names = %#v, want [GITHUB_TOKEN GH_TOKEN]", m.Names)
		}
	})

	// The discriminator branch must win over the exactness heuristic: `names`
	// belongs to IntegrationEnvMethod, yet type=oauth pins IntegrationOAuthMethod.
	t.Run("discriminator wins over field-shape heuristic", func(t *testing.T) {
		raw := `{"id":"x","type":"oauth","label":"L","names":["IGNORED"]}`
		var u opencode.IntegrationMethodUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := u.(opencode.IntegrationOAuthMethod); !ok {
			t.Fatalf("expected IntegrationOAuthMethod, got %T", u)
		}
	})
}

// ConnectionInfoUnion: every variant routes via discriminatorKey="type".
func TestConnectionInfoUnionDiscriminator(t *testing.T) {
	t.Parallel()

	t.Run("credential routes to ConnectionCredentialInfo", func(t *testing.T) {
		raw := `{"type":"credential","id":"cred_1","label":"My Token"}`
		var u opencode.ConnectionInfoUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		c, ok := u.(opencode.ConnectionCredentialInfo)
		if !ok {
			t.Fatalf("expected ConnectionCredentialInfo, got %T", u)
		}
		if c.Type != opencode.ConnectionCredentialInfoTypeCredential {
			t.Errorf("Type = %q, want %q", c.Type, opencode.ConnectionCredentialInfoTypeCredential)
		}
		if c.ID != "cred_1" {
			t.Errorf("ID = %q, want %q", c.ID, "cred_1")
		}
		if c.Label != "My Token" {
			t.Errorf("Label = %q, want %q", c.Label, "My Token")
		}
	})

	t.Run("env routes to ConnectionEnvInfo", func(t *testing.T) {
		raw := `{"type":"env","name":"LINEAR_API_KEY"}`
		var u opencode.ConnectionInfoUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		c, ok := u.(opencode.ConnectionEnvInfo)
		if !ok {
			t.Fatalf("expected ConnectionEnvInfo, got %T", u)
		}
		if c.Type != opencode.ConnectionEnvInfoTypeEnv {
			t.Errorf("Type = %q, want %q", c.Type, opencode.ConnectionEnvInfoTypeEnv)
		}
		if c.Name != "LINEAR_API_KEY" {
			t.Errorf("Name = %q, want %q", c.Name, "LINEAR_API_KEY")
		}
	})
}

// IntegrationPromptUnion: every variant routes via discriminatorKey="type".
func TestIntegrationPromptUnionDiscriminator(t *testing.T) {
	t.Parallel()

	t.Run("text routes to IntegrationTextPrompt", func(t *testing.T) {
		raw := `{"type":"text","key":"enterpriseUrl","message":"GitHub Enterprise URL","placeholder":"https://github.example.com"}`
		var u opencode.IntegrationPromptUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		p, ok := u.(opencode.IntegrationTextPrompt)
		if !ok {
			t.Fatalf("expected IntegrationTextPrompt, got %T", u)
		}
		if p.Type != opencode.IntegrationTextPromptTypeText {
			t.Errorf("Type = %q, want %q", p.Type, opencode.IntegrationTextPromptTypeText)
		}
		if p.Key != "enterpriseUrl" {
			t.Errorf("Key = %q, want %q", p.Key, "enterpriseUrl")
		}
		if p.Placeholder != "https://github.example.com" {
			t.Errorf("Placeholder = %q", p.Placeholder)
		}
	})

	t.Run("select routes to IntegrationSelectPrompt", func(t *testing.T) {
		raw := `{"type":"select","key":"region","message":"Select region","options":[{"label":"US","value":"us","hint":"United States"},{"label":"EU","value":"eu"}]}`
		var u opencode.IntegrationPromptUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		p, ok := u.(opencode.IntegrationSelectPrompt)
		if !ok {
			t.Fatalf("expected IntegrationSelectPrompt, got %T", u)
		}
		if p.Type != opencode.IntegrationSelectPromptTypeSelect {
			t.Errorf("Type = %q, want %q", p.Type, opencode.IntegrationSelectPromptTypeSelect)
		}
		if len(p.Options) != 2 {
			t.Fatalf("Options length = %d, want 2", len(p.Options))
		}
		if p.Options[0].Hint != "United States" {
			t.Errorf("Options[0].Hint = %q", p.Options[0].Hint)
		}
		if p.Options[1].Hint != "" {
			t.Errorf("Options[1].Hint = %q, want empty", p.Options[1].Hint)
		}
	})

	t.Run("text with when condition", func(t *testing.T) {
		raw := `{"type":"text","key":"k","message":"m","when":{"key":"env","op":"neq","value":"prod"}}`
		var u opencode.IntegrationPromptUnion
		if err := apijson.Unmarshal([]byte(raw), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		p, ok := u.(opencode.IntegrationTextPrompt)
		if !ok {
			t.Fatalf("expected IntegrationTextPrompt, got %T", u)
		}
		if p.When.Op != opencode.IntegrationWhenOpNeq {
			t.Errorf("When.Op = %q, want %q", p.When.Op, opencode.IntegrationWhenOpNeq)
		}
		if p.When.Key != "env" || p.When.Value != "prod" {
			t.Errorf("When = %+v", p.When)
		}
	})
}

// The literal "env" is a discriminator value in both IntegrationMethodUnion
// (IntegrationEnvMethod) and ConnectionInfoUnion (ConnectionEnvInfo). Registration
// is keyed per union type, so the two must never bleed into each other.
func TestIntegrationEnvDiscriminatorIsolationAcrossUnions(t *testing.T) {
	t.Parallel()

	var method opencode.IntegrationMethodUnion
	if err := apijson.Unmarshal([]byte(`{"type":"env","names":["A"]}`), &method); err != nil {
		t.Fatalf("unmarshal method: %v", err)
	}
	if _, ok := method.(opencode.IntegrationEnvMethod); !ok {
		t.Fatalf("IntegrationMethodUnion: expected IntegrationEnvMethod, got %T", method)
	}

	var conn opencode.ConnectionInfoUnion
	if err := apijson.Unmarshal([]byte(`{"type":"env","name":"A"}`), &conn); err != nil {
		t.Fatalf("unmarshal connection: %v", err)
	}
	if _, ok := conn.(opencode.ConnectionEnvInfo); !ok {
		t.Fatalf("ConnectionInfoUnion: expected ConnectionEnvInfo, got %T", conn)
	}
}

// TestIntegrationOAuthMethodPromptsMixedArray pins down the shape of
// IntegrationOAuthMethod.Prompts across every boundary the spec allows.
//
// OpenAPI IntegrationOAuthMethod.properties.prompts is:
//
//	{"type":"array","items":{"anyOf":[
//	   {"$ref":"IntegrationTextPrompt"},{"$ref":"IntegrationSelectPrompt"}]}}
//
// and JS SDK v2 types it as
// `prompts?: Array<IntegrationTextPrompt | IntegrationSelectPrompt>`.
//
// The anyOf sits on `items`, not on `prompts`, so a single array may legally mix
// text and select elements. The declared runtime-type comment is therefore
// [[]IntegrationPromptUnion] (one array whose elements are the prompt union) —
// matching the [[]IntegrationMethodUnion] / [[]ConnectionInfoUnion] convention
// used by IntegrationInfo in this same file — and NOT two homogeneous arrays.
//
// UnmarshalJSON routes each element through IntegrationPromptUnion, so the
// carrier's runtime type is []IntegrationPromptUnion and each element is the
// concrete variant selected by its `type` discriminator.
func TestIntegrationOAuthMethodPromptsMixedArray(t *testing.T) {
	// decodePrompts asserts the routed decode shape and returns the typed elements.
	decodePrompts := func(t *testing.T, raw string) []opencode.IntegrationPromptUnion {
		t.Helper()
		var m opencode.IntegrationOAuthMethod
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if m.Type != opencode.IntegrationOAuthMethodTypeOAuth {
			t.Errorf("Type = %q, want oauth", m.Type)
		}
		if m.Prompts == nil {
			return nil
		}
		elems := m.Prompts
		// AsPromptsUnion must agree with the Prompts field.
		if got := m.AsPromptsUnion(); len(got) != len(elems) {
			t.Errorf("AsPromptsUnion() len = %d, want %d", len(got), len(elems))
		}
		return elems
	}

	t.Run("prompts absent (optional per OpenAPI)", func(t *testing.T) {
		var m opencode.IntegrationOAuthMethod
		raw := `{"id":"o","type":"oauth","label":"OAuth"}`
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if m.Prompts != nil {
			t.Errorf("Prompts should be nil when absent, got %v", m.Prompts)
		}
		if got := m.AsPromptsUnion(); got != nil {
			t.Errorf("AsPromptsUnion() = %v, want nil when absent", got)
		}
		if m.ID != "o" || m.Label != "OAuth" {
			t.Errorf("ID/Label = %q/%q, want o/OAuth", m.ID, m.Label)
		}
	})

	t.Run("prompts null", func(t *testing.T) {
		var m opencode.IntegrationOAuthMethod
		raw := `{"id":"o","type":"oauth","label":"OAuth","prompts":null}`
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if m.Prompts != nil {
			t.Errorf("Prompts should be nil for null, got %v (%s)", m.Prompts, reflect.TypeOf(m.Prompts))
		}
	})

	t.Run("empty array", func(t *testing.T) {
		elems := decodePrompts(t, `{"id":"o","type":"oauth","label":"OAuth","prompts":[]}`)
		if elems == nil {
			t.Fatal("Prompts should be a non-nil empty slice for []")
		}
		if len(elems) != 0 {
			t.Errorf("len(Prompts) = %d, want 0", len(elems))
		}
	})

	t.Run("all text", func(t *testing.T) {
		raw := `{"id":"o","type":"oauth","label":"OAuth","prompts":[
			{"type":"text","key":"a","message":"A"},
			{"type":"text","key":"b","message":"B","placeholder":"bb"}
		]}`
		elems := decodePrompts(t, raw)
		if len(elems) != 2 {
			t.Fatalf("len(Prompts) = %d, want 2", len(elems))
		}
		wantKeys := []string{"a", "b"}
		for i, e := range elems {
			tp, ok := e.(opencode.IntegrationTextPrompt)
			if !ok {
				t.Fatalf("Prompts[%d] = %s, want IntegrationTextPrompt", i, reflect.TypeOf(e))
			}
			if tp.Type != opencode.IntegrationTextPromptTypeText {
				t.Errorf("Prompts[%d].Type = %q, want text", i, tp.Type)
			}
			if tp.Key != wantKeys[i] {
				t.Errorf("Prompts[%d].Key = %q, want %q", i, tp.Key, wantKeys[i])
			}
		}
		if elems[1].(opencode.IntegrationTextPrompt).Placeholder != "bb" {
			t.Errorf("Prompts[1].Placeholder = %q, want bb", elems[1].(opencode.IntegrationTextPrompt).Placeholder)
		}
	})

	t.Run("all select", func(t *testing.T) {
		raw := `{"id":"o","type":"oauth","label":"OAuth","prompts":[
			{"type":"select","key":"r","message":"R","options":[{"label":"US","value":"us"}]},
			{"type":"select","key":"s","message":"S","options":[{"label":"EU","value":"eu","hint":"Europe"}]}
		]}`
		elems := decodePrompts(t, raw)
		if len(elems) != 2 {
			t.Fatalf("len(Prompts) = %d, want 2", len(elems))
		}
		for i, e := range elems {
			sp, ok := e.(opencode.IntegrationSelectPrompt)
			if !ok {
				t.Fatalf("Prompts[%d] = %s, want IntegrationSelectPrompt", i, reflect.TypeOf(e))
			}
			if sp.Type != opencode.IntegrationSelectPromptTypeSelect {
				t.Errorf("Prompts[%d].Type = %q, want select", i, sp.Type)
			}
			if len(sp.Options) != 1 {
				t.Fatalf("Prompts[%d].Options len = %d, want 1", i, len(sp.Options))
			}
		}
		if hint := elems[1].(opencode.IntegrationSelectPrompt).Options[0].Hint; hint != "Europe" {
			t.Errorf("Prompts[1].Options[0].Hint = %q, want Europe", hint)
		}
	})

	// The decisive case: one array, both variants, routed per element.
	t.Run("mixed text and select in one array", func(t *testing.T) {
		raw := `{"id":"oauth_gh","type":"oauth","label":"OAuth","prompts":[
			{"type":"text","key":"enterpriseUrl","message":"GitHub Enterprise URL","placeholder":"https://github.example.com"},
			{"type":"select","key":"region","message":"Select region","options":[
				{"label":"US","value":"us","hint":"United States"},
				{"label":"EU","value":"eu"}
			]},
			{"type":"text","key":"note","message":"Note","when":{"key":"region","op":"eq","value":"eu"}}
		]}`
		elems := decodePrompts(t, raw)
		if len(elems) != 3 {
			t.Fatalf("len(Prompts) = %d, want 3", len(elems))
		}
		wantTypes := []any{
			opencode.IntegrationTextPrompt{},
			opencode.IntegrationSelectPrompt{},
			opencode.IntegrationTextPrompt{},
		}
		for i, want := range wantTypes {
			if reflect.TypeOf(elems[i]) != reflect.TypeOf(want) {
				t.Fatalf("Prompts[%d] = %s, want %s", i, reflect.TypeOf(elems[i]), reflect.TypeOf(want))
			}
		}
		// A heterogeneous array proves Prompts can never be a homogeneous
		// []IntegrationTextPrompt or []IntegrationSelectPrompt.
		if reflect.TypeOf(elems[0]) == reflect.TypeOf(elems[1]) {
			t.Error("mixed fixture should hold two different prompt types")
		}
		// Nested `when` must survive routing.
		third := elems[2].(opencode.IntegrationTextPrompt)
		if third.When.Key != "region" || third.When.Op != opencode.IntegrationWhenOpEq || third.When.Value != "eu" {
			t.Errorf("Prompts[2].When = %+v", third.When)
		}
		// Nested select options must survive routing.
		second := elems[1].(opencode.IntegrationSelectPrompt)
		if len(second.Options) != 2 || second.Options[0].Hint != "United States" || second.Options[1].Value != "eu" {
			t.Errorf("Prompts[1].Options = %+v", second.Options)
		}
		// Each element keeps its own RawJSON.
		if second.JSON.RawJSON() == "" {
			t.Error("Prompts[1].RawJSON() empty")
		}
	})

	// Each element is individually decodable into its concrete union variant, and
	// the registered IntegrationPromptUnion resolves the same elements by
	// discriminator — this is what [[]IntegrationPromptUnion] denotes.
	t.Run("elements resolve through IntegrationPromptUnion", func(t *testing.T) {
		text := `{"type":"text","key":"enterpriseUrl","message":"URL","placeholder":"https://x"}`
		sel := `{"type":"select","key":"region","message":"Region","options":[{"label":"US","value":"us"}]}`

		var tp opencode.IntegrationTextPrompt
		if err := json.Unmarshal([]byte(text), &tp); err != nil {
			t.Fatalf("text unmarshal: %v", err)
		}
		if tp.Type != opencode.IntegrationTextPromptTypeText || tp.Key != "enterpriseUrl" {
			t.Errorf("text prompt = %+v", tp)
		}

		var sp opencode.IntegrationSelectPrompt
		if err := json.Unmarshal([]byte(sel), &sp); err != nil {
			t.Fatalf("select unmarshal: %v", err)
		}
		if sp.Type != opencode.IntegrationSelectPromptTypeSelect || len(sp.Options) != 1 {
			t.Errorf("select prompt = %+v", sp)
		}

		for _, tc := range []struct {
			raw  string
			want any
		}{
			{text, opencode.IntegrationTextPrompt{}},
			{sel, opencode.IntegrationSelectPrompt{}},
		} {
			var u opencode.IntegrationPromptUnion
			if err := apijson.UnmarshalRoot([]byte(tc.raw), &u); err != nil {
				t.Fatalf("union unmarshal: %v", err)
			}
			if reflect.TypeOf(u) != reflect.TypeOf(tc.want) {
				t.Errorf("union resolved to %s, want %s", reflect.TypeOf(u), reflect.TypeOf(tc.want))
			}
		}
	})
}

// ===== Array-of-union routing: full boundary matrix =====

// IntegrationInfo.methods is `{"type":"array","items":{"$ref":"IntegrationMethod"}}`
// and IntegrationMethod is `anyOf: [IntegrationOAuthMethod, IntegrationKeyMethod,
// IntegrationEnvMethod]`. JS SDK v2 types it `methods: Array<IntegrationMethod>`.
//
// The carrier is `any` per the Union field rules, so UnmarshalJSON must route every
// element for the declared [[]IntegrationMethodUnion] runtime type to be reachable.
func TestIntegrationInfoMethodsRoutingMatrix(t *testing.T) {
	oauth := `{"id":"m_oauth","type":"oauth","label":"OAuth","prompts":[{"type":"text","key":"k","message":"m"}]}`
	key := `{"type":"key","label":"API Key"}`
	env := `{"type":"env","names":["A","B"]}`

	for _, tc := range []struct {
		name    string
		methods string
		want    []any
	}{
		{"empty array", `[]`, []any{}},
		{"single oauth", `[` + oauth + `]`, []any{opencode.IntegrationOAuthMethod{}}},
		{"single key", `[` + key + `]`, []any{opencode.IntegrationKeyMethod{}}},
		{"single env", `[` + env + `]`, []any{opencode.IntegrationEnvMethod{}}},
		{
			"mixed oauth+key+env",
			`[` + oauth + `,` + key + `,` + env + `]`,
			[]any{opencode.IntegrationOAuthMethod{}, opencode.IntegrationKeyMethod{}, opencode.IntegrationEnvMethod{}},
		},
		{
			"mixed env+oauth+key (order preserved)",
			`[` + env + `,` + oauth + `,` + key + `]`,
			[]any{opencode.IntegrationEnvMethod{}, opencode.IntegrationOAuthMethod{}, opencode.IntegrationKeyMethod{}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			raw := `{"id":"i","name":"N","methods":` + tc.methods + `,"connections":[]}`
			var info opencode.IntegrationInfo
			if err := json.Unmarshal([]byte(raw), &info); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			methods := info.Methods
			if len(methods) != len(tc.want) {
				t.Fatalf("len(Methods) = %d, want %d", len(methods), len(tc.want))
			}
			for i, want := range tc.want {
				if reflect.TypeOf(methods[i]) != reflect.TypeOf(want) {
					t.Errorf("Methods[%d] = %s, want %s", i, reflect.TypeOf(methods[i]), reflect.TypeOf(want))
				}
			}
			// AsMethodsUnion must be the same typed slice.
			if got := info.AsMethodsUnion(); !reflect.DeepEqual(got, methods) {
				t.Errorf("AsMethodsUnion() disagrees with Methods field")
			}
			// Nested prompts on the oauth variant are routed by the field-level array
			// decoder, which internal/apijson reaches even though the variant's own
			// UnmarshalJSON is bypassed for registered union variants.
			for i, m := range methods {
				oa, ok := m.(opencode.IntegrationOAuthMethod)
				if !ok {
					continue
				}
				prompts := oa.Prompts
				if len(prompts) != 1 {
					t.Fatalf("Methods[%d].Prompts len = %d, want 1", i, len(prompts))
				}
				tp, ok := prompts[0].(opencode.IntegrationTextPrompt)
				if !ok {
					t.Fatalf("Methods[%d].Prompts[0] = %s, want IntegrationTextPrompt", i, reflect.TypeOf(prompts[0]))
				}
				if tp.Key != "k" || tp.Message != "m" {
					t.Errorf("Methods[%d].Prompts[0] = %+v", i, tp)
				}
			}
		})
	}

	t.Run("methods absent", func(t *testing.T) {
		var info opencode.IntegrationInfo
		if err := json.Unmarshal([]byte(`{"id":"i","name":"N","connections":[]}`), &info); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if info.Methods != nil {
			t.Errorf("Methods = %v (%s), want nil when absent", info.Methods, reflect.TypeOf(info.Methods))
		}
		if info.AsMethodsUnion() != nil {
			t.Error("AsMethodsUnion() should be nil when methods absent")
		}
	})

	t.Run("oauth without prompts inside methods", func(t *testing.T) {
		raw := `{"id":"i","name":"N","methods":[{"id":"m","type":"oauth","label":"L"}],"connections":[]}`
		var info opencode.IntegrationInfo
		if err := json.Unmarshal([]byte(raw), &info); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		methods := info.AsMethodsUnion()
		if len(methods) != 1 {
			t.Fatalf("len = %d, want 1", len(methods))
		}
		oa, ok := methods[0].(opencode.IntegrationOAuthMethod)
		if !ok {
			t.Fatalf("Methods[0] = %s", reflect.TypeOf(methods[0]))
		}
		if oa.Prompts != nil {
			t.Errorf("Prompts = %v, want nil when absent", oa.Prompts)
		}
		if oa.ID != "m" || oa.Label != "L" {
			t.Errorf("oauth = %+v", oa)
		}
	})
}

// IntegrationInfo.connections is `array` of `ConnectionInfo` =
// `anyOf: [ConnectionCredentialInfo, ConnectionEnvInfo]`.
func TestIntegrationInfoConnectionsRoutingMatrix(t *testing.T) {
	cred := `{"type":"credential","id":"c1","label":"Cred"}`
	env := `{"type":"env","name":"TOKEN"}`

	for _, tc := range []struct {
		name        string
		connections string
		want        []any
	}{
		{"empty array", `[]`, []any{}},
		{"single credential", `[` + cred + `]`, []any{opencode.ConnectionCredentialInfo{}}},
		{"single env", `[` + env + `]`, []any{opencode.ConnectionEnvInfo{}}},
		{"mixed credential+env", `[` + cred + `,` + env + `]`, []any{opencode.ConnectionCredentialInfo{}, opencode.ConnectionEnvInfo{}}},
		{"mixed env+credential (order preserved)", `[` + env + `,` + cred + `]`, []any{opencode.ConnectionEnvInfo{}, opencode.ConnectionCredentialInfo{}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			raw := `{"id":"i","name":"N","methods":[],"connections":` + tc.connections + `}`
			var info opencode.IntegrationInfo
			if err := json.Unmarshal([]byte(raw), &info); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			conns := info.Connections
			if len(conns) != len(tc.want) {
				t.Fatalf("len(Connections) = %d, want %d", len(conns), len(tc.want))
			}
			for i, want := range tc.want {
				if reflect.TypeOf(conns[i]) != reflect.TypeOf(want) {
					t.Errorf("Connections[%d] = %s, want %s", i, reflect.TypeOf(conns[i]), reflect.TypeOf(want))
				}
			}
			if got := info.AsConnectionsUnion(); !reflect.DeepEqual(got, conns) {
				t.Error("AsConnectionsUnion() disagrees with Connections field")
			}
		})
	}

	t.Run("connections absent", func(t *testing.T) {
		var info opencode.IntegrationInfo
		if err := json.Unmarshal([]byte(`{"id":"i","name":"N","methods":[]}`), &info); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if info.Connections != nil {
			t.Errorf("Connections = %v (%s), want nil", info.Connections, reflect.TypeOf(info.Connections))
		}
		if info.AsConnectionsUnion() != nil {
			t.Error("AsConnectionsUnion() should be nil when connections absent")
		}
	})
}

// IntegrationOAuthMethod is a registered variant of IntegrationMethodUnion, so
// internal/apijson skips its own UnmarshalJSON on that path (decoder.go gates the
// indirect unmarshaler on the unionVariants table). Because `prompts` is declared
// as []IntegrationPromptUnion, the struct decoder's field-level array decoder
// routes it anyway — no lazy recovery from RawJSON is needed.
func TestIntegrationOAuthMethodPromptsViaMethodUnion(t *testing.T) {
	raw := `{"id":"m","type":"oauth","label":"L","prompts":[
		{"type":"text","key":"k","message":"m"},
		{"type":"select","key":"s","message":"s","options":[{"label":"L","value":"v"}]}]}`

	var u opencode.IntegrationMethodUnion
	if err := apijson.UnmarshalRoot([]byte(raw), &u); err != nil {
		t.Fatalf("union unmarshal: %v", err)
	}
	oa, ok := u.(opencode.IntegrationOAuthMethod)
	if !ok {
		t.Fatalf("union resolved to %s, want IntegrationOAuthMethod", reflect.TypeOf(u))
	}
	if oa.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
	// The Prompts field itself is already typed on this path.
	prompts := oa.Prompts
	if len(prompts) != 2 {
		t.Fatalf("Prompts len = %d, want 2", len(prompts))
	}
	if reflect.TypeOf(prompts[0]) != reflect.TypeOf(opencode.IntegrationTextPrompt{}) {
		t.Errorf("prompts[0] = %s, want IntegrationTextPrompt", reflect.TypeOf(prompts[0]))
	}
	if reflect.TypeOf(prompts[1]) != reflect.TypeOf(opencode.IntegrationSelectPrompt{}) {
		t.Errorf("prompts[1] = %s, want IntegrationSelectPrompt", reflect.TypeOf(prompts[1]))
	}
	if tp := prompts[0].(opencode.IntegrationTextPrompt); tp.Key != "k" || tp.Message != "m" {
		t.Errorf("prompts[0] = %+v", tp)
	}
	// AsPromptsUnion is the same slice, and stable across calls.
	if got := oa.AsPromptsUnion(); !reflect.DeepEqual(got, prompts) {
		t.Error("AsPromptsUnion() disagrees with Prompts field")
	}
	if again := oa.AsPromptsUnion(); !reflect.DeepEqual(again, prompts) {
		t.Error("AsPromptsUnion() is not stable across calls")
	}
}

// ===== End-to-end envelope decoding =====

// V2IntegrationListResponse wraps `data: Array<IntegrationInfo>`; the routing must
// survive the extra array + envelope layers.
func TestV2IntegrationListResponseUnionRouting(t *testing.T) {
	raw := `{
		"location": {"directory": "/tmp/p"},
		"data": [
			{"id":"github","name":"GitHub","methods":[
				{"id":"m","type":"oauth","label":"OAuth","prompts":[
					{"type":"text","key":"url","message":"URL"},
					{"type":"select","key":"r","message":"R","options":[{"label":"US","value":"us"}]}]},
				{"type":"key","label":"Key"},
				{"type":"env","names":["GH_TOKEN"]}
			],"connections":[
				{"type":"credential","id":"c","label":"C"},
				{"type":"env","name":"GH_TOKEN"}
			]},
			{"id":"linear","name":"Linear","methods":[{"type":"key"}],"connections":[]}
		]
	}`
	var resp opencode.V2IntegrationListResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(resp.Data) != 2 {
		t.Fatalf("len(Data) = %d, want 2", len(resp.Data))
	}
	if resp.JSON.RawJSON() == "" {
		t.Error("envelope RawJSON not preserved")
	}

	first := resp.Data[0]
	methods := first.Methods
	wantMethods := []any{
		opencode.IntegrationOAuthMethod{},
		opencode.IntegrationKeyMethod{},
		opencode.IntegrationEnvMethod{},
	}
	if len(methods) != len(wantMethods) {
		t.Fatalf("len(Data[0].Methods) = %d, want %d", len(methods), len(wantMethods))
	}
	for i, want := range wantMethods {
		if reflect.TypeOf(methods[i]) != reflect.TypeOf(want) {
			t.Errorf("Data[0].Methods[%d] = %s, want %s", i, reflect.TypeOf(methods[i]), reflect.TypeOf(want))
		}
	}
	// Deepest level: prompts inside the oauth method inside data[0].
	oa := methods[0].(opencode.IntegrationOAuthMethod)
	prompts := oa.Prompts
	if len(prompts) != 2 {
		t.Fatalf("prompts len = %d, want 2", len(prompts))
	}
	if reflect.TypeOf(prompts[0]) != reflect.TypeOf(opencode.IntegrationTextPrompt{}) {
		t.Errorf("prompts[0] = %s", reflect.TypeOf(prompts[0]))
	}
	if reflect.TypeOf(prompts[1]) != reflect.TypeOf(opencode.IntegrationSelectPrompt{}) {
		t.Errorf("prompts[1] = %s", reflect.TypeOf(prompts[1]))
	}
	env := methods[2].(opencode.IntegrationEnvMethod)
	if len(env.Names) != 1 || env.Names[0] != "GH_TOKEN" {
		t.Errorf("env method = %+v", env)
	}

	conns := first.Connections
	if len(conns) != 2 {
		t.Fatalf("len(Data[0].Connections) = %d, want 2", len(conns))
	}
	if reflect.TypeOf(conns[0]) != reflect.TypeOf(opencode.ConnectionCredentialInfo{}) {
		t.Errorf("Connections[0] = %s", reflect.TypeOf(conns[0]))
	}
	if reflect.TypeOf(conns[1]) != reflect.TypeOf(opencode.ConnectionEnvInfo{}) {
		t.Errorf("Connections[1] = %s", reflect.TypeOf(conns[1]))
	}

	// Second element: empty connections array plus a key method with no label.
	second := resp.Data[1]
	if got := second.Connections; got == nil || len(got) != 0 {
		t.Errorf("Data[1].Connections = %#v, want empty []ConnectionInfoUnion", second.Connections)
	}
	secondMethods := second.AsMethodsUnion()
	if len(secondMethods) != 1 {
		t.Fatalf("Data[1] methods len = %d, want 1", len(secondMethods))
	}
	if km, ok := secondMethods[0].(opencode.IntegrationKeyMethod); !ok || km.Label != "" {
		t.Errorf("Data[1].Methods[0] = %#v", secondMethods[0])
	}
}

// V2IntegrationGetResponse wraps a single `data: IntegrationInfo`.
func TestV2IntegrationGetResponseUnionRouting(t *testing.T) {
	raw := `{
		"location": {"directory": "/tmp/p", "workspaceID": "w1"},
		"data": {"id":"github","name":"GitHub","methods":[
			{"type":"env","names":["A"]},
			{"id":"m","type":"oauth","label":"OAuth","prompts":[{"type":"select","key":"s","message":"S","options":[]}]}
		],"connections":[{"type":"env","name":"A"}]}
	}`
	var resp opencode.V2IntegrationGetResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	methods := resp.Data.AsMethodsUnion()
	if len(methods) != 2 {
		t.Fatalf("methods len = %d, want 2", len(methods))
	}
	if reflect.TypeOf(methods[0]) != reflect.TypeOf(opencode.IntegrationEnvMethod{}) {
		t.Errorf("methods[0] = %s, want IntegrationEnvMethod", reflect.TypeOf(methods[0]))
	}
	oa, ok := methods[1].(opencode.IntegrationOAuthMethod)
	if !ok {
		t.Fatalf("methods[1] = %s, want IntegrationOAuthMethod", reflect.TypeOf(methods[1]))
	}
	prompts := oa.Prompts
	if len(prompts) != 1 {
		t.Fatalf("prompts len = %d, want 1", len(prompts))
	}
	sp, ok := prompts[0].(opencode.IntegrationSelectPrompt)
	if !ok {
		t.Fatalf("prompts[0] = %s, want IntegrationSelectPrompt", reflect.TypeOf(prompts[0]))
	}
	if sp.Key != "s" || len(sp.Options) != 0 {
		t.Errorf("prompts[0] = %+v", sp)
	}
	conns := resp.Data.AsConnectionsUnion()
	if len(conns) != 1 {
		t.Fatalf("connections len = %d, want 1", len(conns))
	}
	if ce, ok := conns[0].(opencode.ConnectionEnvInfo); !ok || ce.Name != "A" {
		t.Errorf("connections[0] = %#v", conns[0])
	}
}

// ===== IntegrationAttemptStatus carrier types =====

// OpenAPI IntegrationAttemptStatus is an anyOf of four objects. `message` appears
// in exactly one variant ("failed") and is always `string`; `time` appears in all
// four with the identical shape. Both carriers are therefore concrete Go types —
// only genuinely polymorphic carriers use `any` + a runtime-type comment (compare
// ToolPartState.Error string vs ToolPartState.Time any in session.go).
func TestIntegrationAttemptStatusCarrierPort(t *testing.T) {
	for _, tc := range []struct {
		name        string
		raw         string
		wantStatus  opencode.IntegrationAttemptStatusType
		wantUnion   any
		wantMessage string
	}{
		{
			"pending", `{"status":"pending","time":{"created":1700000000,"expires":1700003600}}`,
			opencode.IntegrationAttemptStatusTypePending, opencode.IntegrationAttemptStatusPending{}, "",
		},
		{
			"complete", `{"status":"complete","time":{"created":1700000000,"expires":1700003600}}`,
			opencode.IntegrationAttemptStatusTypeComplete, opencode.IntegrationAttemptStatusComplete{}, "",
		},
		{
			"failed", `{"status":"failed","message":"User denied access","time":{"created":1700000000,"expires":1700003600}}`,
			opencode.IntegrationAttemptStatusTypeFailed, opencode.IntegrationAttemptStatusFailed{}, "User denied access",
		},
		{
			"expired", `{"status":"expired","time":{"created":1700000000,"expires":1700003600}}`,
			opencode.IntegrationAttemptStatusTypeExpired, opencode.IntegrationAttemptStatusExpired{}, "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var s opencode.IntegrationAttemptStatus
			if err := json.Unmarshal([]byte(tc.raw), &s); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if s.Status != tc.wantStatus {
				t.Errorf("Status = %q, want %q", s.Status, tc.wantStatus)
			}
			// Discriminator routing.
			if reflect.TypeOf(s.AsUnion()) != reflect.TypeOf(tc.wantUnion) {
				t.Fatalf("AsUnion() = %s, want %s", reflect.TypeOf(s.AsUnion()), reflect.TypeOf(tc.wantUnion))
			}
			// apijson.Port must transfer the concrete carrier types.
			if s.Message != tc.wantMessage {
				t.Errorf("Message = %q, want %q", s.Message, tc.wantMessage)
			}
			if reflect.TypeOf(s.Time) != reflect.TypeOf(opencode.IntegrationAttemptTime{}) {
				t.Fatalf("Time = %s, want IntegrationAttemptTime", reflect.TypeOf(s.Time))
			}
			// The time sub-object's scalar anyOf carriers decode as float64 for numbers.
			created, ok := s.Time.Created.(float64)
			if !ok {
				t.Fatalf("Time.Created = %s, want float64", reflect.TypeOf(s.Time.Created))
			}
			if created != 1700000000 {
				t.Errorf("Time.Created = %v, want 1700000000", created)
			}
			expires, ok := s.Time.Expires.(float64)
			if !ok {
				t.Fatalf("Time.Expires = %s, want float64", reflect.TypeOf(s.Time.Expires))
			}
			if expires != 1700003600 {
				t.Errorf("Time.Expires = %v, want 1700003600", expires)
			}
			if s.Time.JSON.RawJSON() == "" {
				t.Error("Time.RawJSON() empty")
			}
			if s.JSON.RawJSON() == "" {
				t.Error("status RawJSON() empty")
			}
		})
	}
}

// The scalar anyOf on time.created/time.expires also admits the non-finite string
// sentinels "NaN", "Infinity" and "-Infinity" per OpenAPI.
func TestIntegrationAttemptStatusTimeNonFiniteSentinels(t *testing.T) {
	for _, sentinel := range []string{"NaN", "Infinity", "-Infinity"} {
		t.Run(sentinel, func(t *testing.T) {
			raw := `{"status":"pending","time":{"created":"` + sentinel + `","expires":"` + sentinel + `"}}`
			var s opencode.IntegrationAttemptStatus
			if err := json.Unmarshal([]byte(raw), &s); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if reflect.TypeOf(s.AsUnion()) != reflect.TypeOf(opencode.IntegrationAttemptStatusPending{}) {
				t.Errorf("AsUnion() = %s", reflect.TypeOf(s.AsUnion()))
			}
			got, ok := s.Time.Created.(string)
			if !ok {
				t.Fatalf("Time.Created = %s, want string", reflect.TypeOf(s.Time.Created))
			}
			if got != sentinel {
				t.Errorf("Time.Created = %q, want %q", got, sentinel)
			}
		})
	}
}

// TestIntegrationUnionSliceNullElements pins the illegal-element boundary of the
// arrays of anyOf shared by IntegrationInfo.methods, IntegrationInfo.connections
// and IntegrationOAuthMethod.prompts.
//
// OpenAPI types these arrays as `array` whose `items` is an anyOf of object
// variants; no variant is `{"type":"null"}`, so a null element is not a legal server
// response. The fields are declared as typed slices of the union interface, so
// internal/apijson decodes them field-by-field: the array decoder rejects the
// illegal element and the struct decoder records the field as invalid instead of
// propagating the failure. The observable contract is therefore: no error, the whole
// array field degrades to nil, sibling fields survive, the payload stays verbatim in
// RawJSON, and an enclosing envelope keeps all of its entries.
func TestIntegrationUnionSliceNullElements(t *testing.T) {
	t.Run("null methods and connections elements degrade those fields", func(t *testing.T) {
		raw := `{"id":"gh","name":"GitHub","methods":[null,{"type":"key","label":"Key"}],"connections":[null,{"type":"env","name":"GH_TOKEN"}]}`
		var info opencode.IntegrationInfo
		if err := json.Unmarshal([]byte(raw), &info); err != nil {
			t.Fatalf("null array elements must not fail the decode, got: %v", err)
		}
		if info.ID != "gh" || info.Name != "GitHub" {
			t.Errorf("sibling fields lost: id=%q name=%q", info.ID, info.Name)
		}
		if got := info.AsMethodsUnion(); len(got) != 0 {
			t.Errorf("len(AsMethodsUnion()) = %d, want 0 (illegal element degrades the field)", len(got))
		}
		if got := info.AsConnectionsUnion(); len(got) != 0 {
			t.Errorf("len(AsConnectionsUnion()) = %d, want 0 (illegal element degrades the field)", len(got))
		}
		// The rejected element is still available verbatim.
		if !strings.Contains(info.JSON.RawJSON(), "null") {
			t.Error("RawJSON() should still carry the rejected null element")
		}
	})

	t.Run("null prompts element degrades only the nested prompts field", func(t *testing.T) {
		raw := `{"id":"gh","name":"GitHub","connections":[],"methods":[{"id":"m1","type":"oauth","label":"OAuth","prompts":[null,{"type":"text","key":"k","message":"m"}]}]}`
		var info opencode.IntegrationInfo
		if err := json.Unmarshal([]byte(raw), &info); err != nil {
			t.Fatalf("null prompts element must not fail the decode, got: %v", err)
		}
		methods := info.AsMethodsUnion()
		if len(methods) != 1 {
			t.Fatalf("len(AsMethodsUnion()) = %d, want 1 — the outer array must survive", len(methods))
		}
		oauth, ok := methods[0].(opencode.IntegrationOAuthMethod)
		if !ok {
			t.Fatalf("methods[0] = %T, want opencode.IntegrationOAuthMethod", methods[0])
		}
		// The sibling fields of the variant itself are intact.
		if oauth.ID != "m1" || oauth.Label != "OAuth" {
			t.Errorf("variant sibling fields lost: id=%q label=%q", oauth.ID, oauth.Label)
		}
		if got := oauth.AsPromptsUnion(); len(got) != 0 {
			t.Errorf("len(AsPromptsUnion()) = %d, want 0", len(got))
		}
	})

	t.Run("null prompts element via the standalone oauth method decode", func(t *testing.T) {
		raw := `{"id":"m1","type":"oauth","label":"OAuth","prompts":[{"type":"select","key":"s","message":"m","options":[]},null]}`
		var oauth opencode.IntegrationOAuthMethod
		if err := json.Unmarshal([]byte(raw), &oauth); err != nil {
			t.Fatalf("null prompts element must not fail the decode, got: %v", err)
		}
		if oauth.ID != "m1" || oauth.Label != "OAuth" {
			t.Errorf("sibling fields lost: id=%q label=%q", oauth.ID, oauth.Label)
		}
		if got := oauth.AsPromptsUnion(); len(got) != 0 {
			t.Errorf("len(AsPromptsUnion()) = %d, want 0", len(got))
		}
	})

	t.Run("null element inside the list envelope keeps every integration", func(t *testing.T) {
		raw := `{"location":{"directory":"/d","project":{"id":"pr","directory":"/d"}},"data":[` +
			`{"id":"gh","name":"GitHub","methods":[null],"connections":[]},` +
			`{"id":"gl","name":"GitLab","methods":[{"type":"env","names":["GL"]}],"connections":[]}]}`
		var resp opencode.V2IntegrationListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal envelope: %v", err)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("len(Data) = %d, want 2 — a null element must not drop list entries", len(resp.Data))
		}
		if resp.Data[0].ID != "gh" || len(resp.Data[0].AsMethodsUnion()) != 0 {
			t.Errorf("Data[0] = {id:%q methods:%d}, want {gh 0}", resp.Data[0].ID, len(resp.Data[0].AsMethodsUnion()))
		}
		if resp.Data[1].ID != "gl" || len(resp.Data[1].AsMethodsUnion()) != 1 {
			t.Errorf("Data[1] = {id:%q methods:%d}, want {gl 1}", resp.Data[1].ID, len(resp.Data[1].AsMethodsUnion()))
		}
	})

	t.Run("wrong-typed element degrades the field, siblings survive", func(t *testing.T) {
		var info opencode.IntegrationInfo
		if err := json.Unmarshal([]byte(`{"id":"gh","name":"n","methods":["oauth"],"connections":[]}`), &info); err != nil {
			t.Fatalf(`methods:["oauth"] must not fail the decode, got: %v`, err)
		}
		if info.ID != "gh" || info.Name != "n" {
			t.Errorf("sibling fields lost: id=%q name=%q", info.ID, info.Name)
		}
		if info.Methods != nil {
			t.Errorf("Methods = %#v, want nil", info.Methods)
		}
		if info.Connections == nil || len(info.Connections) != 0 {
			t.Errorf("Connections = %#v, want empty slice — a sibling array must be unaffected", info.Connections)
		}
	})

	t.Run("non-array payload degrades the field, siblings survive", func(t *testing.T) {
		// Previously the `any` carrier leaked the raw scalar; the typed field cannot.
		var info opencode.IntegrationInfo
		if err := json.Unmarshal([]byte(`{"id":"gh","name":"n","methods":"nope","connections":[]}`), &info); err != nil {
			t.Fatalf(`methods:"nope" must not fail the decode, got: %v`, err)
		}
		if info.ID != "gh" {
			t.Errorf("sibling field lost: id=%q", info.ID)
		}
		if info.Methods != nil {
			t.Errorf("Methods = %#v, want nil for a non-array payload", info.Methods)
		}
	})

	t.Run("null and absent arrays leave the carriers nil", func(t *testing.T) {
		for name, raw := range map[string]string{
			"null":   `{"id":"gh","name":"n","methods":null,"connections":null}`,
			"absent": `{"id":"gh","name":"n"}`,
		} {
			var info opencode.IntegrationInfo
			if err := json.Unmarshal([]byte(raw), &info); err != nil {
				t.Fatalf("%s arrays: %v", name, err)
			}
			if info.Methods != nil || info.Connections != nil {
				t.Errorf("%s arrays: Methods = %#v, Connections = %#v, want both nil", name, info.Methods, info.Connections)
			}
			if info.AsMethodsUnion() != nil || info.AsConnectionsUnion() != nil {
				t.Errorf("%s arrays: accessors should return nil", name)
			}
		}
	})

	t.Run("empty arrays yield empty typed slices", func(t *testing.T) {
		var info opencode.IntegrationInfo
		if err := json.Unmarshal([]byte(`{"id":"gh","name":"n","methods":[],"connections":[]}`), &info); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if got := reflect.TypeOf(info.Methods).String(); got != "[]opencode.IntegrationMethodUnion" {
			t.Errorf("Methods type = %s, want []opencode.IntegrationMethodUnion", got)
		}
		if got := reflect.TypeOf(info.Connections).String(); got != "[]opencode.ConnectionInfoUnion" {
			t.Errorf("Connections type = %s, want []opencode.ConnectionInfoUnion", got)
		}
		if len(info.AsMethodsUnion()) != 0 || len(info.AsConnectionsUnion()) != 0 {
			t.Error("empty arrays should produce zero-length typed slices")
		}
	})
}
