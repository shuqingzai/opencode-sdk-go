package opencode

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/sst/opencode-sdk-go/internal/apijson"
)

// Aligned with OpenAPI /provider/list, /provider/auth, /provider/{providerID}/oauth/authorize,
// /provider/{providerID}/oauth/callback and JS SDK(v2) types.gen.ts.

// ── Request serialization ──────────────────────────────────────────────────────

func TestProviderOauthAuthorizeParamsSerialization(t *testing.T) {
	t.Run("required method only; query fields not in body", func(t *testing.T) {
		p := ProviderOauthAuthorizeParams{
			Method:    F(int64(0)),
			Directory: F("dir"),
			Workspace: F("ws"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		// apijson marshals alphabetically; query fields MUST NOT appear in body
		want := `{"method":0}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
		if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
			t.Errorf("query fields leaked into body: %s", got)
		}
	})

	t.Run("full params including optional inputs", func(t *testing.T) {
		p := ProviderOauthAuthorizeParams{
			Method: F(int64(2)),
			Inputs: F(map[string]string{"api_key": "sk-test", "region": "us-east"}),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"method":2`) {
			t.Errorf("missing method field in: %s", got)
		}
		if !strings.Contains(got, `"inputs"`) {
			t.Errorf("missing inputs field in: %s", got)
		}
		if !strings.Contains(got, `"api_key"`) {
			t.Errorf("missing api_key in inputs: %s", got)
		}
	})

	t.Run("method is int64 (OpenAPI number → integer index → int64)", func(t *testing.T) {
		p := ProviderOauthAuthorizeParams{
			Method: F(int64(9999)),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"method":9999}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestProviderOauthCallbackParamsSerialization(t *testing.T) {
	t.Run("required method only; query fields not in body", func(t *testing.T) {
		p := ProviderOauthCallbackParams{
			Method:    F(int64(1)),
			Directory: F("d"),
			Workspace: F("w"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"method":1}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
		if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
			t.Errorf("query fields leaked into body: %s", got)
		}
	})

	t.Run("full params including optional code", func(t *testing.T) {
		p := ProviderOauthCallbackParams{
			Method: F(int64(0)),
			Code:   F("auth-code-xyz"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"code":"auth-code-xyz","method":0}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

// ── Response deserialization ───────────────────────────────────────────────────

func TestProviderListResponseDeserialization(t *testing.T) {
	t.Run("full response with provider and model", func(t *testing.T) {
		raw := `{
			"all": [
				{
					"id": "anthropic",
					"name": "Anthropic",
					"source": "config",
					"env": ["ANTHROPIC_API_KEY"],
					"key": "anthropic",
					"options": {},
					"models": {
						"claude-3-5-sonnet": {
							"id": "claude-3-5-sonnet",
							"providerID": "anthropic",
							"api": {"id": "anthropic", "url": "https://api.anthropic.com", "npm": "@ai-sdk/anthropic"},
							"name": "Claude 3.5 Sonnet",
							"capabilities": {
								"temperature": true,
								"reasoning": false,
								"attachment": true,
								"toolcall": true,
								"input": {"text": true, "audio": false, "image": true, "video": false, "pdf": false},
								"output": {"text": true, "audio": false, "image": false, "video": false, "pdf": false},
								"interleaved": false
							},
							"cost": {
								"input": 3.0,
								"output": 15.0,
								"cache": {"read": 0.3, "write": 3.75}
							},
							"limit": {"context": 200000, "output": 8096},
							"status": "active",
							"options": {},
							"headers": {},
							"release_date": "2024-10"
						}
					}
				}
			],
			"default": {"anthropic": "claude-3-5-sonnet"},
			"connected": ["anthropic"]
		}`
		var resp ProviderListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatal(err)
		}
		if len(resp.All) != 1 {
			t.Fatalf("expected 1 provider, got %d", len(resp.All))
		}
		p := resp.All[0]
		if p.ID != "anthropic" {
			t.Errorf("ID = %s, want anthropic", p.ID)
		}
		if p.Source != ProviderInfoSourceConfig {
			t.Errorf("Source = %s, want config", p.Source)
		}
		if len(p.Env) != 1 || p.Env[0] != "ANTHROPIC_API_KEY" {
			t.Errorf("Env = %v", p.Env)
		}
		m := p.Models["claude-3-5-sonnet"]
		if m.ID != "claude-3-5-sonnet" {
			t.Errorf("model ID = %s", m.ID)
		}
		if m.API.URL != "https://api.anthropic.com" {
			t.Errorf("API.URL = %s", m.API.URL)
		}
		if m.Cost.Input != 3.0 {
			t.Errorf("Cost.Input = %v, want 3.0", m.Cost.Input)
		}
		if m.Limit.Context != 200000 {
			t.Errorf("Limit.Context = %d, want 200000", m.Limit.Context)
		}
		if resp.Default["anthropic"] != "claude-3-5-sonnet" {
			t.Errorf("Default = %v", resp.Default)
		}
		if len(resp.Connected) != 1 || resp.Connected[0] != "anthropic" {
			t.Errorf("Connected = %v", resp.Connected)
		}
		// RawJSON preserved
		if resp.JSON.RawJSON() == "" {
			t.Error("RawJSON should not be empty")
		}
	})

	t.Run("ProviderInfo source enum values", func(t *testing.T) {
		cases := map[string]ProviderInfoSource{
			"env":    ProviderInfoSourceEnv,
			"config": ProviderInfoSourceConfig,
			"custom": ProviderInfoSourceCustom,
			"api":    ProviderInfoSourceAPI,
		}
		for jsonVal, goConst := range cases {
			raw := `{"id":"p","name":"n","source":"` + jsonVal + `","env":[],"options":{},"models":{}}`
			var p ProviderInfo
			if err := json.Unmarshal([]byte(raw), &p); err != nil {
				t.Fatalf("source=%s: %v", jsonVal, err)
			}
			if p.Source != goConst {
				t.Errorf("source=%s: got %s", jsonVal, p.Source)
			}
			if !p.Source.IsKnown() {
				t.Errorf("source=%s: IsKnown() = false", jsonVal)
			}
		}
	})

	t.Run("ProviderInfoSourceAPI renamed constant has value 'api'", func(t *testing.T) {
		if string(ProviderInfoSourceAPI) != "api" {
			t.Errorf("ProviderInfoSourceAPI = %q, want \"api\"", ProviderInfoSourceAPI)
		}
	})
}

// TestProviderModelVariantsDeserialization pins the static type of
// ProviderModel.Variants.
//
// OpenAPI Model.variants is
//
//	{"type":"object","additionalProperties":{"type":"object"}}
//
// and JS SDK(v2) types it as `variants?: {[key: string]: {[key: string]: unknown}}`.
// The value schema is constrained to `object`, so the Go mapping is
// map[string]map[string]any -- a bare map[string]any would under-specify the
// declared value type.
func TestProviderModelVariantsDeserialization(t *testing.T) {
	if got, want := reflect.TypeOf(ProviderModel{}.Variants).String(), "map[string]map[string]interface {}"; got != want {
		t.Errorf("ProviderModel.Variants static type = %s, want %s", got, want)
	}

	model := func(variants string) string {
		return `{"id":"m","providerID":"p","api":{"id":"a","url":"u","npm":"n"},"name":"M",` +
			`"capabilities":{"temperature":true,"reasoning":true,"attachment":false,"toolcall":true,` +
			`"input":{"text":true,"audio":false,"image":false,"video":false,"pdf":false},` +
			`"output":{"text":true,"audio":false,"image":false,"video":false,"pdf":false},` +
			`"interleaved":false},` +
			`"cost":{"input":1,"output":2,"cache":{"read":0,"write":0}},` +
			`"limit":{"context":100,"output":50},"status":"active","options":{},"headers":{},` +
			`"release_date":"2024-01-01"` + variants + `}`
	}

	t.Run("nested free-form objects survive", func(t *testing.T) {
		var m ProviderModel
		raw := model(`,"variants":{"thinking":{"reasoningEffort":"high","budget":32000,` +
			`"nested":{"a":[1,2]}},"fast":{}}`)
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatal(err)
		}
		if len(m.Variants) != 2 {
			t.Fatalf("len(Variants) = %d, want 2", len(m.Variants))
		}
		// Element type is a concrete map -- no type assertion needed, which is
		// exactly what the tightened static type buys the caller.
		thinking := m.Variants["thinking"]
		if thinking["reasoningEffort"] != "high" {
			t.Errorf("variants.thinking.reasoningEffort = %v", thinking["reasoningEffort"])
		}
		if nested, ok := thinking["nested"].(map[string]any); !ok || len(nested) != 1 {
			t.Errorf("variants.thinking.nested = %#v", thinking["nested"])
		}
		if fast, ok := m.Variants["fast"]; !ok || len(fast) != 0 {
			t.Errorf("variants.fast = %#v (present=%v), want empty object", fast, ok)
		}
	})

	t.Run("variants absent (optional per OpenAPI)", func(t *testing.T) {
		var m ProviderModel
		if err := json.Unmarshal([]byte(model("")), &m); err != nil {
			t.Fatal(err)
		}
		if m.Variants != nil {
			t.Errorf("Variants = %#v, want nil", m.Variants)
		}
	})

	t.Run("variants null does not abort the decode", func(t *testing.T) {
		var m ProviderModel
		if err := json.Unmarshal([]byte(model(`,"variants":null`)), &m); err != nil {
			t.Fatalf("variants:null must not fail the decode, got: %v", err)
		}
		if m.Variants != nil {
			t.Errorf("Variants = %#v, want nil", m.Variants)
		}
		if m.ID != "m" {
			t.Error("sibling fields lost on a null variants")
		}
	})

	t.Run("empty variants object", func(t *testing.T) {
		var m ProviderModel
		if err := json.Unmarshal([]byte(model(`,"variants":{}`)), &m); err != nil {
			t.Fatal(err)
		}
		if len(m.Variants) != 0 {
			t.Errorf("len(Variants) = %d, want 0", len(m.Variants))
		}
	})
}

func TestProviderAuthMethodDeserialization(t *testing.T) {
	t.Run("type oauth no prompts", func(t *testing.T) {
		raw := `{"type":"oauth","label":"Sign in with Google"}`
		var a ProviderAuthMethod
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatal(err)
		}
		if a.Type != ProviderAuthMethodTypeOauth {
			t.Errorf("Type = %s", a.Type)
		}
		if a.Label != "Sign in with Google" {
			t.Errorf("Label = %s", a.Label)
		}
		if a.Prompts != nil {
			t.Errorf("Prompts should be nil, got %v", a.Prompts)
		}
	})

	t.Run("type api IsKnown check", func(t *testing.T) {
		if !ProviderAuthMethodTypeAPI.IsKnown() {
			t.Error("ProviderAuthMethodTypeAPI.IsKnown() = false")
		}
		if string(ProviderAuthMethodTypeAPI) != "api" {
			t.Errorf("ProviderAuthMethodTypeAPI = %q, want \"api\"", ProviderAuthMethodTypeAPI)
		}
	})

	t.Run("prompts mixed array — text then select (anyOf per element)", func(t *testing.T) {
		// OpenAPI ProviderAuthMethod.prompts = array with items.anyOf: [text, select]
		// Each element is independently text OR select — mixed array.
		raw := `{
			"type": "api",
			"label": "API Key",
			"prompts": [
				{"type":"text","key":"api_key","message":"Enter API key","placeholder":"sk-..."},
				{"type":"select","key":"region","message":"Choose region","options":[
					{"label":"US East","value":"us-east"},
					{"label":"EU West","value":"eu-west","hint":"GDPR-compliant"}
				]}
			]
		}`
		var a ProviderAuthMethod
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatal(err)
		}
		if a.Type != ProviderAuthMethodTypeAPI {
			t.Errorf("Type = %s", a.Type)
		}
		if a.Prompts == nil {
			t.Fatal("Prompts should not be nil")
		}
		// Prompts is declared []ProviderAuthMethodPrompt, so the array decoder routes
		// every element through the registered variants of that union: one concrete
		// variant per element.
		prompts := a.Prompts
		if len(prompts) != 2 {
			t.Fatalf("expected 2 prompts, got %d", len(prompts))
		}
		text, ok := prompts[0].(ProviderAuthMethodPromptText)
		if !ok {
			t.Fatalf("prompts[0] = %T, want ProviderAuthMethodPromptText", prompts[0])
		}
		if text.Key != "api_key" || text.Message != "Enter API key" || text.Placeholder != "sk-..." {
			t.Errorf("prompts[0] = %+v", text)
		}
		sel, ok := prompts[1].(ProviderAuthMethodPromptSelect)
		if !ok {
			t.Fatalf("prompts[1] = %T, want ProviderAuthMethodPromptSelect", prompts[1])
		}
		if sel.Key != "region" || len(sel.Options) != 2 {
			t.Fatalf("prompts[1] = %+v", sel)
		}
		if sel.Options[1].Value != "eu-west" || sel.Options[1].Hint != "GDPR-compliant" {
			t.Errorf("prompts[1].Options[1] = %+v", sel.Options[1])
		}
		// AsPromptsUnion is the typed accessor for the same payload.
		if got := a.AsPromptsUnion(); len(got) != 2 {
			t.Fatalf("AsPromptsUnion() len = %d, want 2", len(got))
		}
	})

	t.Run("prompts with only text type", func(t *testing.T) {
		raw := `{
			"type": "api",
			"label": "API Key",
			"prompts": [
				{"type":"text","key":"key1","message":"Enter value"},
				{"type":"text","key":"key2","message":"Enter another"}
			]
		}`
		var a ProviderAuthMethod
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatal(err)
		}
		prompts := a.Prompts
		if len(prompts) != 2 {
			t.Fatalf("expected 2 prompts, got %d", len(prompts))
		}
		for i, want := range []string{"key1", "key2"} {
			text, ok := prompts[i].(ProviderAuthMethodPromptText)
			if !ok {
				t.Fatalf("prompts[%d] = %T, want ProviderAuthMethodPromptText", i, prompts[i])
			}
			if text.Type != ProviderAuthMethodPromptTextTypeText || text.Key != want {
				t.Errorf("prompts[%d] = %+v, want key %q", i, text, want)
			}
		}
	})

	t.Run("prompts with select and when condition", func(t *testing.T) {
		raw := `{
			"type": "api",
			"label": "Custom Provider",
			"prompts": [
				{
					"type": "select",
					"key": "mode",
					"message": "Select mode",
					"options": [{"label":"Auto","value":"auto"},{"label":"Manual","value":"manual"}],
					"when": {"key":"env","op":"eq","value":"prod"}
				}
			]
		}`
		var a ProviderAuthMethod
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatal(err)
		}
		if a.Prompts == nil {
			t.Fatal("Prompts should not be nil")
		}
	})

	t.Run("RawJSON preserved", func(t *testing.T) {
		raw := `{"type":"oauth","label":"OAuth"}`
		var a ProviderAuthMethod
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatal(err)
		}
		if a.JSON.RawJSON() == "" {
			t.Error("RawJSON should not be empty")
		}
	})
}

// Tests the ProviderAuthMethodPrompt union discriminator dispatch.
func TestProviderAuthMethodPromptUnionDispatch(t *testing.T) {
	t.Run("text discriminator dispatches to ProviderAuthMethodPromptText", func(t *testing.T) {
		raw := `{"type":"text","key":"api_key","message":"Enter API key","placeholder":"sk-..."}`
		var p ProviderAuthMethodPrompt
		if err := apijson.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatal(err)
		}
		pt, ok := p.(ProviderAuthMethodPromptText)
		if !ok {
			t.Fatalf("expected ProviderAuthMethodPromptText, got %T", p)
		}
		if pt.Key != "api_key" {
			t.Errorf("Key = %s", pt.Key)
		}
		if pt.Type != ProviderAuthMethodPromptTextTypeText {
			t.Errorf("Type = %s", pt.Type)
		}
		if pt.Placeholder != "sk-..." {
			t.Errorf("Placeholder = %s", pt.Placeholder)
		}
	})

	t.Run("select discriminator dispatches to ProviderAuthMethodPromptSelect", func(t *testing.T) {
		raw := `{"type":"select","key":"region","message":"Choose region","options":[{"label":"US","value":"us"},{"label":"EU","value":"eu","hint":"GDPR"}]}`
		var p ProviderAuthMethodPrompt
		if err := apijson.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatal(err)
		}
		ps, ok := p.(ProviderAuthMethodPromptSelect)
		if !ok {
			t.Fatalf("expected ProviderAuthMethodPromptSelect, got %T", p)
		}
		if ps.Key != "region" {
			t.Errorf("Key = %s", ps.Key)
		}
		if len(ps.Options) != 2 {
			t.Fatalf("expected 2 options, got %d", len(ps.Options))
		}
		if ps.Options[1].Hint != "GDPR" {
			t.Errorf("Options[1].Hint = %s", ps.Options[1].Hint)
		}
	})

	t.Run("text with when condition", func(t *testing.T) {
		raw := `{"type":"text","key":"k","message":"m","when":{"key":"env","op":"neq","value":"prod"}}`
		var p ProviderAuthMethodPrompt
		if err := apijson.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatal(err)
		}
		pt, ok := p.(ProviderAuthMethodPromptText)
		if !ok {
			t.Fatalf("expected ProviderAuthMethodPromptText, got %T", p)
		}
		if pt.When.Op != ProviderAuthMethodPromptWhenOpNeq {
			t.Errorf("When.Op = %s", pt.When.Op)
		}
	})
}

func TestProviderOauthAuthorizeResponseDeserialization(t *testing.T) {
	t.Run("auto method", func(t *testing.T) {
		raw := `{"url":"https://oauth.example.com/authorize?client_id=abc","method":"auto","instructions":"Click the link below"}`
		var r ProviderOauthAuthorizeResponse
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatal(err)
		}
		if r.URL != "https://oauth.example.com/authorize?client_id=abc" {
			t.Errorf("URL = %s", r.URL)
		}
		if r.Method != ProviderOauthAuthorizeResponseMethodAuto {
			t.Errorf("Method = %s", r.Method)
		}
		if r.Instructions != "Click the link below" {
			t.Errorf("Instructions = %s", r.Instructions)
		}
		if r.JSON.RawJSON() == "" {
			t.Error("RawJSON should not be empty")
		}
	})

	t.Run("code method", func(t *testing.T) {
		raw := `{"url":"https://example.com/auth","method":"code","instructions":"Enter the code"}`
		var r ProviderOauthAuthorizeResponse
		if err := json.Unmarshal([]byte(raw), &r); err != nil {
			t.Fatal(err)
		}
		if r.Method != ProviderOauthAuthorizeResponseMethodCode {
			t.Errorf("Method = %s, want code", r.Method)
		}
		if !r.Method.IsKnown() {
			t.Error("Method.IsKnown() = false")
		}
	})

	t.Run("unknown method is not IsKnown", func(t *testing.T) {
		unknown := ProviderOauthAuthorizeResponseMethod("manual")
		if unknown.IsKnown() {
			t.Error("unknown method should not be IsKnown()")
		}
	})
}

// ── Enum boundary tests ────────────────────────────────────────────────────────

func TestProviderInfoSourceIsKnown(t *testing.T) {
	known := []ProviderInfoSource{
		ProviderInfoSourceEnv,
		ProviderInfoSourceConfig,
		ProviderInfoSourceCustom,
		ProviderInfoSourceAPI,
	}
	for _, v := range known {
		if !v.IsKnown() {
			t.Errorf("IsKnown() = false for %q", v)
		}
	}
	unknown := ProviderInfoSource("unknown-source")
	if unknown.IsKnown() {
		t.Errorf("IsKnown() = true for unknown value %q", unknown)
	}
}

func TestProviderAuthMethodTypeIsKnown(t *testing.T) {
	known := []ProviderAuthMethodType{
		ProviderAuthMethodTypeOauth,
		ProviderAuthMethodTypeAPI,
	}
	for _, v := range known {
		if !v.IsKnown() {
			t.Errorf("IsKnown() = false for %q", v)
		}
	}
	unknown := ProviderAuthMethodType("unknown-type")
	if unknown.IsKnown() {
		t.Errorf("IsKnown() = true for unknown value %q", unknown)
	}
}

func TestProviderModelStatusIsKnown(t *testing.T) {
	known := []ProviderModelStatus{
		ProviderModelStatusAlpha,
		ProviderModelStatusBeta,
		ProviderModelStatusDeprecated,
		ProviderModelStatusActive,
	}
	for _, v := range known {
		if !v.IsKnown() {
			t.Errorf("IsKnown() = false for %q", v)
		}
	}
	unknown := ProviderModelStatus("unknown-status")
	if unknown.IsKnown() {
		t.Errorf("IsKnown() = true for unknown value %q", unknown)
	}
}

// ── Provider type alias ────────────────────────────────────────────────────────

func TestProviderTypeAlias(t *testing.T) {
	// Provider = ProviderInfo: both types are interchangeable
	var p Provider
	raw := `{"id":"test","name":"Test","source":"env","env":[],"options":{},"models":{}}`
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		t.Fatal(err)
	}
	if p.ID != "test" {
		t.Errorf("Provider.ID = %s, want test", p.ID)
	}
	// Ensure assignment compatibility
	var pi ProviderInfo = p
	_ = pi
	var p2 Provider = pi
	_ = p2
}
