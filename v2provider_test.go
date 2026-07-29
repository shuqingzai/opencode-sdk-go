// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"reflect"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
)

// TestV2ProviderInfoAPIUnionAisdk verifies the aisdk branch of V2ProviderInfo.API.
// Aligned with OpenAPI ProviderApi anyOf -> ProviderAISDK: required type, package;
// optional url, settings.
//
// V2ProviderInfo is a plain response struct that holds a union-typed `api` field —
// it is not a union root (its UnmarshalJSON decodes into itself, then decodes the
// `api` field separately) — so AsAPIUnion is the correct accessor name.
func TestV2ProviderInfoAPIUnionAisdk(t *testing.T) {
	raw := `{
		"id": "openai",
		"integrationID": "openai-int",
		"name": "OpenAI",
		"disabled": false,
		"api": {
			"type": "aisdk",
			"package": "@ai-sdk/openai",
			"url": "https://api.openai.com/v1",
			"settings": {"compatibility": "strict"}
		},
		"request": {"headers": {"authorization": "Bearer x"}, "body": {"stream": true}}
	}`
	var p opencode.V2ProviderInfo
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if p.ID != "openai" {
		t.Errorf("ID = %q, want openai", p.ID)
	}
	if p.IntegrationID != "openai-int" {
		t.Errorf("IntegrationID = %q, want openai-int", p.IntegrationID)
	}
	if p.Name != "OpenAI" {
		t.Errorf("Name = %q, want OpenAI", p.Name)
	}
	if p.Disabled {
		t.Error("Disabled should be false")
	}

	u := p.AsAPIUnion()
	if u == nil {
		t.Fatal("AsAPIUnion() should not be nil for aisdk api")
	}
	aisdk, ok := u.(opencode.V2ProviderInfoAPIAisdk)
	if !ok {
		t.Fatalf("expected V2ProviderInfoAPIAisdk, got %T", u)
	}
	if aisdk.Type != opencode.V2ProviderInfoAPIAisdkTypeAisdk {
		t.Errorf("aisdk.Type = %q, want aisdk", aisdk.Type)
	}
	if aisdk.Package != "@ai-sdk/openai" {
		t.Errorf("aisdk.Package = %q, want @ai-sdk/openai", aisdk.Package)
	}
	if aisdk.URL != "https://api.openai.com/v1" {
		t.Errorf("aisdk.URL = %q, want https://api.openai.com/v1", aisdk.URL)
	}

	// ProviderRequest.Body is map[string]any per OpenAPI `{"type":"object"}`.
	if p.Request.Headers["authorization"] != "Bearer x" {
		t.Errorf("Request.Headers[authorization] = %q, want Bearer x", p.Request.Headers["authorization"])
	}
	if p.Request.Body["stream"] != true {
		t.Errorf("Request.Body[stream] = %v, want true", p.Request.Body["stream"])
	}
}

// TestV2ProviderInfoAPIUnionNative verifies the native branch of V2ProviderInfo.API.
// Aligned with OpenAPI ProviderApi anyOf -> ProviderNative: required type, settings;
// optional url.
func TestV2ProviderInfoAPIUnionNative(t *testing.T) {
	raw := `{
		"id": "ollama",
		"name": "Ollama",
		"api": {
			"type": "native",
			"url": "http://localhost:11434",
			"settings": {"keep_alive": "5m"}
		},
		"request": {"headers": {}, "body": {}}
	}`
	var p opencode.V2ProviderInfo
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	u := p.AsAPIUnion()
	if u == nil {
		t.Fatal("AsAPIUnion() should not be nil for native api")
	}
	native, ok := u.(opencode.V2ProviderInfoAPINative)
	if !ok {
		t.Fatalf("expected V2ProviderInfoAPINative, got %T", u)
	}
	if native.Type != opencode.V2ProviderInfoAPINativeTypeNative {
		t.Errorf("native.Type = %q, want native", native.Type)
	}
	if native.URL != "http://localhost:11434" {
		t.Errorf("native.URL = %q, want http://localhost:11434", native.URL)
	}
	// Settings is a concrete map[string]any — indexable without a type assertion.
	if native.Settings["keep_alive"] != "5m" {
		t.Errorf("native.Settings[keep_alive] = %v, want 5m", native.Settings["keep_alive"])
	}

	// The API field mirrors the decoded union value.
	if p.API == nil {
		t.Error("API field should be populated alongside the union")
	}
}

// TestV2ProviderInfoAPIUnionNilWhenAbsent verifies AsAPIUnion returns nil when the
// api field is absent, so callers can nil-check instead of panicking.
func TestV2ProviderInfoAPIUnionNilWhenAbsent(t *testing.T) {
	raw := `{"id":"p","name":"P","request":{"headers":{},"body":{}}}`
	var p opencode.V2ProviderInfo
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		t.Fatalf("unmarshal without api: %v", err)
	}
	if p.AsAPIUnion() != nil {
		t.Errorf("AsAPIUnion() = %v, want nil when api absent", p.AsAPIUnion())
	}
}

// TestV2ProviderInfoAPINullDoesNotAbortDecode covers the four boundary shapes of
// the union-carrying `api` field: a legal value, an explicit null, an absent field,
// and a wrong-typed scalar.
//
// OpenAPI ProviderV2Info lists `api` in required and resolves it to ProviderApi,
// whose anyOf holds only ProviderAISDK and ProviderNative — there is no
// `{"type":"null"}` member, so a null is not a legal server response. It must still
// not abort the decode: neither registered variant filters on gjson.Null, so routing
// a null fails with "was not able to coerce type as union", and that error
// propagating out of UnmarshalJSON kills the surrounding document — inside
// V2ProviderListResponse it silently empties the whole `data` array.
func TestV2ProviderInfoAPINullDoesNotAbortDecode(t *testing.T) {
	t.Run("explicit null is treated as absent", func(t *testing.T) {
		var p opencode.V2ProviderInfo
		if err := json.Unmarshal([]byte(`{"id":"p1","name":"P One","api":null}`), &p); err != nil {
			t.Fatalf("api:null must not fail the decode, got: %v", err)
		}
		if p.API != nil || p.AsAPIUnion() != nil {
			t.Errorf("API = %#v / AsAPIUnion = %v, want both nil", p.API, p.AsAPIUnion())
		}
		if p.ID != "p1" || p.Name != "P One" {
			t.Errorf("sibling fields lost: id=%q name=%q", p.ID, p.Name)
		}
	})

	t.Run("null api inside the list envelope keeps every element", func(t *testing.T) {
		raw := `{"location":{"directory":"/d","project":{"id":"pr","directory":"/d"}},"data":[` +
			`{"id":"p1","name":"One","api":null},` +
			`{"id":"p2","name":"Two","api":{"type":"native","settings":{}}}]}`
		var resp opencode.V2ProviderListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal envelope: %v", err)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("len(Data) = %d, want 2 — a null `api` must not drop list elements", len(resp.Data))
		}
		if resp.Data[0].ID != "p1" || resp.Data[0].Name != "One" {
			t.Errorf("Data[0] = {id:%q name:%q}, want {p1 One}", resp.Data[0].ID, resp.Data[0].Name)
		}
		if _, ok := resp.Data[1].API.(opencode.V2ProviderInfoAPINative); !ok {
			t.Errorf("Data[1].API = %T, want opencode.V2ProviderInfoAPINative", resp.Data[1].API)
		}
	})

	t.Run("wrong-typed api is reported through field metadata, not a fatal error", func(t *testing.T) {
		// OpenAPI admits no scalar variant for ProviderApi, so `api:3` is an illegal
		// server response. Because `api` is declared as the registered
		// [V2ProviderInfoAPIUnion], the framework's struct decoder records the
		// coercion failure on the field's metadata and keeps decoding the rest of
		// the document. The "null == absent" vs "wrong type" distinction survives on
		// JSON.API rather than as a document-killing error.
		var p opencode.V2ProviderInfo
		if err := json.Unmarshal([]byte(`{"id":"p1","name":"P","api":3}`), &p); err != nil {
			t.Fatalf("a wrong-typed api must not fail the whole decode, got: %v", err)
		}
		if p.API != nil {
			t.Errorf("API = %#v, want nil for a non-object api", p.API)
		}
		if !p.JSON.API.IsInvalid() {
			t.Error("JSON.API.IsInvalid() = false, want true for a non-object api")
		}
		if p.JSON.API.IsNull() || p.JSON.API.IsMissing() {
			t.Error("a wrong-typed api must stay distinguishable from null/absent")
		}
		if got := p.JSON.API.Raw(); got != "3" {
			t.Errorf("JSON.API.Raw() = %s, want 3 — the raw value must stay available", got)
		}
		if p.ID != "p1" || p.Name != "P" {
			t.Errorf("sibling fields lost: id=%q name=%q", p.ID, p.Name)
		}
	})

	t.Run("wrong-typed api inside the list envelope keeps every element", func(t *testing.T) {
		// The regression this locks down: routing the sub-document by hand inside
		// UnmarshalJSON returned the coercion error out of the element decoder, the
		// slice decoder propagated it, and the outer struct decoder swallowed it —
		// so `data` came back empty with a nil error, losing the valid element too.
		raw := `{"location":{"directory":"/d","project":{"id":"pr","directory":"/d"}},"data":[` +
			`{"id":"p1","name":"One","api":3},` +
			`{"id":"p2","name":"Two","api":{"type":"native","settings":{}}}]}`
		var resp opencode.V2ProviderListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal envelope: %v", err)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("len(Data) = %d, want 2 — a wrong-typed `api` must not drop list elements", len(resp.Data))
		}
		if resp.Data[0].ID != "p1" || resp.Data[0].Name != "One" || resp.Data[0].API != nil {
			t.Errorf("Data[0] = {id:%q name:%q api:%#v}, want {p1 One <nil>}", resp.Data[0].ID, resp.Data[0].Name, resp.Data[0].API)
		}
		if _, ok := resp.Data[1].API.(opencode.V2ProviderInfoAPINative); !ok {
			t.Errorf("Data[1].API = %T, want opencode.V2ProviderInfoAPINative", resp.Data[1].API)
		}
	})

	t.Run("unknown future variant falls back without an error", func(t *testing.T) {
		// Forward compatibility: a `type` the SDK has never heard of must not fail
		// the decode. It coerces loosely onto a registered variant whose IsKnown()
		// reports false, which is the documented way to detect it.
		var p opencode.V2ProviderInfo
		if err := json.Unmarshal([]byte(`{"id":"p1","name":"P","api":{"type":"future","package":"pkg"}}`), &p); err != nil {
			t.Fatalf("an unknown api variant must not fail the decode, got: %v", err)
		}
		aisdk, ok := p.API.(opencode.V2ProviderInfoAPIAisdk)
		if !ok {
			t.Fatalf("API = %T, want opencode.V2ProviderInfoAPIAisdk as the loose fallback", p.API)
		}
		if aisdk.Type.IsKnown() {
			t.Errorf("Type.IsKnown() = true for %q, want false so callers can detect it", aisdk.Type)
		}
		if p.ID != "p1" || p.Name != "P" {
			t.Errorf("sibling fields lost: id=%q name=%q", p.ID, p.Name)
		}
	})
}

// TestV2ProviderInfoAPITypeIsKnown verifies both ProviderApi discriminator enums.
func TestV2ProviderInfoAPITypeIsKnown(t *testing.T) {
	if !opencode.V2ProviderInfoAPIAisdkTypeAisdk.IsKnown() {
		t.Error("aisdk type should be known")
	}
	if opencode.V2ProviderInfoAPIAisdkType("unknown").IsKnown() {
		t.Error("unknown aisdk type should not be known")
	}
	if !opencode.V2ProviderInfoAPINativeTypeNative.IsKnown() {
		t.Error("native type should be known")
	}
	if opencode.V2ProviderInfoAPINativeType("unknown").IsKnown() {
		t.Error("unknown native type should not be known")
	}
}

// TestV2ProviderInfoUnknownFieldTolerated verifies forward compatibility: unknown
// fields must not break deserialization and are retained in RawJSON.
func TestV2ProviderInfoUnknownFieldTolerated(t *testing.T) {
	raw := `{
		"id": "p",
		"name": "P",
		"api": {"type": "native", "settings": {}},
		"request": {"headers": {}, "body": {}},
		"futureField": "ignored"
	}`
	var p opencode.V2ProviderInfo
	if err := json.Unmarshal([]byte(raw), &p); err != nil {
		t.Fatalf("unmarshal with unknown field: %v", err)
	}
	if p.JSON.RawJSON() == "" {
		t.Error("RawJSON should retain the original payload")
	}
}

// TestV2ProviderInfoAPISettingsIsMap verifies that the `settings` field on both
// ProviderApi anyOf variants decodes as a concrete map[string]any.
//
// OpenAPI ProviderAISDK.properties.settings and ProviderNative.properties.settings
// are both `{"type":"object"}` with no anyOf/oneOf, and JS SDK v2 types them as
// `settings?: { [key: string]: unknown }` (ProviderAisdk) / `settings: { [key:
// string]: unknown }` (ProviderNative). They are plain objects — not unions — so
// the declared Go type is map[string]any with no runtime-type comment.
//
// Required-ness per OpenAPI: ProviderAISDK required=[type,package] => settings
// optional; ProviderNative required=[type,settings] => settings required.
func TestV2ProviderInfoAPISettingsIsMap(t *testing.T) {
	t.Run("aisdk populated settings", func(t *testing.T) {
		raw := `{"type":"aisdk","package":"@ai-sdk/anthropic","url":"https://api.anthropic.com","settings":{"apiVersion":"2023-06-01","beta":true}}`
		var a opencode.V2ProviderInfoAPIAisdk
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if a.Settings["apiVersion"] != "2023-06-01" {
			t.Errorf("Settings[apiVersion] = %v, want 2023-06-01", a.Settings["apiVersion"])
		}
		if a.Settings["beta"] != true {
			t.Errorf("Settings[beta] = %v, want true", a.Settings["beta"])
		}
		if got := reflect.TypeOf(a.Settings); got.String() != "map[string]interface {}" {
			t.Errorf("reflect.TypeOf(Settings) = %s, want map[string]interface {}", got)
		}
	})

	t.Run("aisdk empty settings object", func(t *testing.T) {
		raw := `{"type":"aisdk","package":"p","settings":{}}`
		var a opencode.V2ProviderInfoAPIAisdk
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if a.Settings == nil {
			t.Error("Settings should be a non-nil empty map for {}")
		}
		if len(a.Settings) != 0 {
			t.Errorf("Settings should be empty, got %d keys", len(a.Settings))
		}
	})

	t.Run("aisdk settings absent (optional per OpenAPI)", func(t *testing.T) {
		raw := `{"type":"aisdk","package":"p"}`
		var a opencode.V2ProviderInfoAPIAisdk
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if a.Settings != nil {
			t.Errorf("Settings should be nil when absent, got %v", a.Settings)
		}
	})

	t.Run("aisdk settings null", func(t *testing.T) {
		raw := `{"type":"aisdk","package":"p","settings":null}`
		var a opencode.V2ProviderInfoAPIAisdk
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if a.Settings != nil {
			t.Errorf("Settings should be nil for null, got %v", a.Settings)
		}
	})

	t.Run("native nested settings", func(t *testing.T) {
		raw := `{"type":"native","url":"http://localhost:11434","settings":{"headers":{"x-trace":"on"},"retries":2}}`
		var n opencode.V2ProviderInfoAPINative
		if err := json.Unmarshal([]byte(raw), &n); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if n.Settings["retries"] != float64(2) {
			t.Errorf("Settings[retries] = %v, want 2", n.Settings["retries"])
		}
		hdrs, ok := n.Settings["headers"].(map[string]any)
		if !ok {
			t.Fatalf("Settings[headers] should be a nested object, got %T", n.Settings["headers"])
		}
		if hdrs["x-trace"] != "on" {
			t.Errorf("Settings[headers][x-trace] = %v, want on", hdrs["x-trace"])
		}
		if got := reflect.TypeOf(n.Settings); got.String() != "map[string]interface {}" {
			t.Errorf("reflect.TypeOf(Settings) = %s, want map[string]interface {}", got)
		}
	})

	t.Run("native empty settings object", func(t *testing.T) {
		raw := `{"type":"native","settings":{}}`
		var n opencode.V2ProviderInfoAPINative
		if err := json.Unmarshal([]byte(raw), &n); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if n.Settings == nil {
			t.Error("Settings should be a non-nil empty map for {}")
		}
		if len(n.Settings) != 0 {
			t.Errorf("Settings should be empty, got %d keys", len(n.Settings))
		}
	})

	t.Run("settings survives the union round trip", func(t *testing.T) {
		raw := `{"id":"ollama","name":"Ollama","api":{"type":"native","settings":{"keep_alive":"5m"}},"request":{"headers":{},"body":{}}}`
		var p opencode.V2ProviderInfo
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		native, ok := p.AsAPIUnion().(opencode.V2ProviderInfoAPINative)
		if !ok {
			t.Fatalf("expected V2ProviderInfoAPINative, got %T", p.AsAPIUnion())
		}
		if native.Settings["keep_alive"] != "5m" {
			t.Errorf("Settings[keep_alive] = %v, want 5m", native.Settings["keep_alive"])
		}
	})
}
