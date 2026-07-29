// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// TestV2ModelListWithOptionalParams tests the v2 model list endpoint.
// Aligned with OpenAPI operationId "v2.model.list", GET /api/model.
// Query parameters: location (optional, nested object with directory and workspace).
func TestV2ModelListWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Model.List(context.TODO(), opencode.V2ModelListParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
			Workspace: opencode.F("workspace"),
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

// TestV2ModelListResponseUnmarshal verifies deserialization of V2ModelListResponse.
// Aligned with OpenAPI GET /api/model response schema.
// Required fields: location (LocationInfo), data ([]V2ModelInfo).
func TestV2ModelListResponseUnmarshal(t *testing.T) {
	t.Run("full response with one model", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/home/user/proj",
				"workspaceID": "ws_42",
				"project": {"id": "proj_abc", "directory": "/home/user/proj"}
			},
			"data": [
				{
					"id": "claude-3-sonnet",
					"providerID": "anthropic",
					"family": "claude-3",
					"name": "Claude 3 Sonnet",
					"api": {
						"type": "aisdk",
						"id": "anthropic/claude-3-sonnet",
						"package": "@ai-sdk/anthropic"
					},
					"capabilities": {"tools": true, "input": ["text", "image"], "output": ["text"]},
					"request": {"headers": {}, "body": null, "variant": ""},
					"variants": [],
					"time": {"released": 1700000000},
					"cost": [{"input": 0.003, "output": 0.015, "cache": {"read": 0.0003, "write": 0.00375}}],
					"status": "active",
					"enabled": true,
					"limit": {"context": 200000, "output": 4096}
				}
			]
		}`
		var resp opencode.V2ModelListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if resp.Location.Directory != "/home/user/proj" {
			t.Errorf("Location.Directory = %q, want /home/user/proj", resp.Location.Directory)
		}
		if len(resp.Data) != 1 {
			t.Fatalf("Data len = %d, want 1", len(resp.Data))
		}
		m := resp.Data[0]
		if m.ID != "claude-3-sonnet" {
			t.Errorf("model.ID = %q, want claude-3-sonnet", m.ID)
		}
		if m.ProviderID != "anthropic" {
			t.Errorf("model.ProviderID = %q, want anthropic", m.ProviderID)
		}
		if m.Name != "Claude 3 Sonnet" {
			t.Errorf("model.Name = %q, want Claude 3 Sonnet", m.Name)
		}
		if m.Status != opencode.V2ModelInfoStatusActive {
			t.Errorf("model.Status = %q, want active", m.Status)
		}
		if !m.Enabled {
			t.Error("model.Enabled should be true")
		}
		if m.Time.Released != 1700000000 {
			t.Errorf("model.Time.Released = %d, want 1700000000", m.Time.Released)
		}
		if m.Limit.Context != 200000 {
			t.Errorf("model.Limit.Context = %d, want 200000", m.Limit.Context)
		}
		if m.Limit.Output != 4096 {
			t.Errorf("model.Limit.Output = %d, want 4096", m.Limit.Output)
		}
		if resp.JSON.RawJSON() == "" {
			t.Error("RawJSON should be preserved")
		}
	})

	t.Run("empty data array", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/",
				"project": {"id": "p", "directory": "/"}
			},
			"data": []
		}`
		var resp opencode.V2ModelListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(resp.Data) != 0 {
			t.Errorf("expected 0 models, got %d", len(resp.Data))
		}
	})
}

// TestV2ModelInfoAPIUnionAisdk verifies the aisdk branch of V2ModelInfo.API union.
// Aligned with OpenAPI anyOf V2ModelInfoAPI: aisdk variant with "type":"aisdk".
func TestV2ModelInfoAPIUnionAisdk(t *testing.T) {
	raw := `{
		"id": "gpt-4",
		"providerID": "openai",
		"name": "GPT-4",
		"api": {
			"type": "aisdk",
			"id": "openai/gpt-4",
			"package": "@ai-sdk/openai",
			"url": "https://api.openai.com"
		},
		"capabilities": {"tools": true, "input": ["text"], "output": ["text"]},
		"request": {"headers": {}, "body": null},
		"variants": [],
		"time": {"released": 1680000000},
		"cost": [{"input": 0.03, "output": 0.06, "cache": {"read": 0.003, "write": 0.006}}],
		"status": "active",
		"enabled": true,
		"limit": {"context": 128000, "output": 8192}
	}`
	var m opencode.V2ModelInfo
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	u := m.AsAPIUnion()
	if u == nil {
		t.Fatal("AsAPIUnion() should not be nil for aisdk api")
	}
	aisdk, ok := u.(opencode.V2ModelInfoAPIAisdk)
	if !ok {
		t.Fatalf("expected V2ModelInfoAPIAisdk, got %T", u)
	}
	if aisdk.ID != "openai/gpt-4" {
		t.Errorf("aisdk.ID = %q, want openai/gpt-4", aisdk.ID)
	}
	if aisdk.Type != opencode.V2ModelInfoAPIAisdkTypeAisdk {
		t.Errorf("aisdk.Type = %q, want aisdk", aisdk.Type)
	}
	if aisdk.Package != "@ai-sdk/openai" {
		t.Errorf("aisdk.Package = %q, want @ai-sdk/openai", aisdk.Package)
	}
	if aisdk.URL != "https://api.openai.com" {
		t.Errorf("aisdk.URL = %q, want https://api.openai.com", aisdk.URL)
	}
}

// TestV2ModelInfoAPIUnionNative verifies the native branch of V2ModelInfo.API union.
// Aligned with OpenAPI anyOf V2ModelInfoAPI: native variant with "type":"native".
func TestV2ModelInfoAPIUnionNative(t *testing.T) {
	raw := `{
		"id": "local-model",
		"providerID": "ollama",
		"name": "Local Model",
		"api": {
			"type": "native",
			"id": "ollama/local-model",
			"url": "http://localhost:11434",
			"settings": {"timeout": 30}
		},
		"capabilities": {"tools": false, "input": ["text"], "output": ["text"]},
		"request": {"headers": {}, "body": null},
		"variants": [],
		"time": {"released": 1690000000},
		"cost": [{"input": 0.0, "output": 0.0, "cache": {"read": 0.0, "write": 0.0}}],
		"status": "beta",
		"enabled": true,
		"limit": {"context": 4096, "output": 1024}
	}`
	var m opencode.V2ModelInfo
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	u := m.AsAPIUnion()
	if u == nil {
		t.Fatal("AsAPIUnion() should not be nil for native api")
	}
	native, ok := u.(opencode.V2ModelInfoAPINative)
	if !ok {
		t.Fatalf("expected V2ModelInfoAPINative, got %T", u)
	}
	if native.ID != "ollama/local-model" {
		t.Errorf("native.ID = %q, want ollama/local-model", native.ID)
	}
	if native.Type != opencode.V2ModelInfoAPINativeTypeNative {
		t.Errorf("native.Type = %q, want native", native.Type)
	}
	if native.URL != "http://localhost:11434" {
		t.Errorf("native.URL = %q, want http://localhost:11434", native.URL)
	}
}

// TestV2ModelInfoStatusIsKnown verifies all known V2ModelInfoStatus values.
func TestV2ModelInfoStatusIsKnown(t *testing.T) {
	known := []opencode.V2ModelInfoStatus{
		opencode.V2ModelInfoStatusAlpha,
		opencode.V2ModelInfoStatusBeta,
		opencode.V2ModelInfoStatusDeprecated,
		opencode.V2ModelInfoStatusActive,
	}
	for _, s := range known {
		if !s.IsKnown() {
			t.Errorf("%q should be known", s)
		}
	}
	if opencode.V2ModelInfoStatus("unknown").IsKnown() {
		t.Error("unknown status should not be known")
	}
}

// TestV2ModelInfoAPIAisdkTypeIsKnown verifies V2ModelInfoAPIAisdkType enum.
func TestV2ModelInfoAPIAisdkTypeIsKnown(t *testing.T) {
	if !opencode.V2ModelInfoAPIAisdkTypeAisdk.IsKnown() {
		t.Error("aisdk type should be known")
	}
	if opencode.V2ModelInfoAPIAisdkType("unknown").IsKnown() {
		t.Error("unknown type should not be known")
	}
}

// TestV2ModelInfoAPINativeTypeIsKnown verifies V2ModelInfoAPINativeType enum.
func TestV2ModelInfoAPINativeTypeIsKnown(t *testing.T) {
	if !opencode.V2ModelInfoAPINativeTypeNative.IsKnown() {
		t.Error("native type should be known")
	}
	if opencode.V2ModelInfoAPINativeType("unknown").IsKnown() {
		t.Error("unknown type should not be known")
	}
}

// TestV2ModelInfoCostItemUnmarshal verifies deserialization of V2ModelInfoCostItem.
// Aligned with OpenAPI ModelCostItem schema.
// Required: input (float64), output (float64), cache (V2ModelInfoCostCache).
// Optional: tier (V2ModelInfoCostTier with type and size).
func TestV2ModelInfoCostItemUnmarshal(t *testing.T) {
	t.Run("with tier", func(t *testing.T) {
		raw := `{"tier":{"type":"context","size":128000},"input":0.003,"output":0.015,"cache":{"read":0.0003,"write":0.00375}}`
		var item opencode.V2ModelInfoCostItem
		if err := json.Unmarshal([]byte(raw), &item); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if item.Tier.Type != opencode.V2ModelInfoCostTierTypeContext {
			t.Errorf("Tier.Type = %q, want context", item.Tier.Type)
		}
		if item.Tier.Size != 128000 {
			t.Errorf("Tier.Size = %d, want 128000", item.Tier.Size)
		}
		if item.Input != 0.003 {
			t.Errorf("Input = %v, want 0.003", item.Input)
		}
		if item.Output != 0.015 {
			t.Errorf("Output = %v, want 0.015", item.Output)
		}
		if item.Cache.Read != 0.0003 {
			t.Errorf("Cache.Read = %v, want 0.0003", item.Cache.Read)
		}
		if item.Cache.Write != 0.00375 {
			t.Errorf("Cache.Write = %v, want 0.00375", item.Cache.Write)
		}
	})

	t.Run("without tier", func(t *testing.T) {
		raw := `{"input":0.0,"output":0.0,"cache":{"read":0.0,"write":0.0}}`
		var item opencode.V2ModelInfoCostItem
		if err := json.Unmarshal([]byte(raw), &item); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if item.Tier.Size != 0 {
			t.Errorf("Tier.Size should be zero when absent, got %d", item.Tier.Size)
		}
	})
}

// TestV2ModelInfoLimitUnmarshal verifies deserialization of V2ModelInfoLimit.
// Aligned with OpenAPI ModelLimit schema: required context and output; optional input.
func TestV2ModelInfoLimitUnmarshal(t *testing.T) {
	t.Run("full limit with input", func(t *testing.T) {
		raw := `{"context":200000,"input":100000,"output":4096}`
		var limit opencode.V2ModelInfoLimit
		if err := json.Unmarshal([]byte(raw), &limit); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if limit.Context != 200000 {
			t.Errorf("Context = %d, want 200000", limit.Context)
		}
		if limit.Input != 100000 {
			t.Errorf("Input = %d, want 100000", limit.Input)
		}
		if limit.Output != 4096 {
			t.Errorf("Output = %d, want 4096", limit.Output)
		}
	})

	t.Run("limit without optional input", func(t *testing.T) {
		raw := `{"context":128000,"output":8192}`
		var limit opencode.V2ModelInfoLimit
		if err := json.Unmarshal([]byte(raw), &limit); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if limit.Input != 0 {
			t.Errorf("Input should be 0 when absent, got %d", limit.Input)
		}
	})
}

// TestV2ModelInfoRequestBodyIsMap verifies V2ModelInfoRequest.Body decodes as a
// concrete map[string]any. OpenAPI ModelV2Info.request.body is `{"type":"object"}`
// with no anyOf/oneOf, and JS SDK v2 types it as `{ [key: string]: unknown }`,
// so it is a plain object — not a union — and needs no runtime-type comment.
func TestV2ModelInfoRequestBodyIsMap(t *testing.T) {
	t.Run("populated body", func(t *testing.T) {
		raw := `{"headers":{"x-api-key":"secret"},"body":{"stream":true,"max_tokens":4096},"variant":"thinking"}`
		var req opencode.V2ModelInfoRequest
		if err := json.Unmarshal([]byte(raw), &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if req.Headers["x-api-key"] != "secret" {
			t.Errorf("Headers[x-api-key] = %q, want secret", req.Headers["x-api-key"])
		}
		// Body is map[string]any — indexable without a type assertion.
		if req.Body["stream"] != true {
			t.Errorf("Body[stream] = %v, want true", req.Body["stream"])
		}
		if req.Body["max_tokens"] != float64(4096) {
			t.Errorf("Body[max_tokens] = %v, want 4096", req.Body["max_tokens"])
		}
		if req.Variant != "thinking" {
			t.Errorf("Variant = %q, want thinking", req.Variant)
		}
	})

	t.Run("empty body object", func(t *testing.T) {
		raw := `{"headers":{},"body":{}}`
		var req opencode.V2ModelInfoRequest
		if err := json.Unmarshal([]byte(raw), &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if req.Body == nil {
			t.Error("Body should be a non-nil empty map for {}")
		}
		if len(req.Body) != 0 {
			t.Errorf("Body should be empty, got %d keys", len(req.Body))
		}
	})

	t.Run("null body", func(t *testing.T) {
		raw := `{"headers":{},"body":null}`
		var req opencode.V2ModelInfoRequest
		if err := json.Unmarshal([]byte(raw), &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if req.Body != nil {
			t.Errorf("Body should be nil for null, got %v", req.Body)
		}
	})
}

// TestV2ModelInfoVariantBodyIsMap verifies V2ModelInfoVariant.Body decodes as a
// concrete map[string]any. OpenAPI ModelV2Info.variants[].body is
// `{"type":"object"}` with no anyOf/oneOf, matching JS SDK v2
// `body: { [key: string]: unknown }`.
func TestV2ModelInfoVariantBodyIsMap(t *testing.T) {
	raw := `{"id":"thinking","headers":{"anthropic-beta":"output-128k"},"body":{"thinking":{"budget_tokens":10000}}}`
	var v opencode.V2ModelInfoVariant
	if err := json.Unmarshal([]byte(raw), &v); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if v.ID != "thinking" {
		t.Errorf("ID = %q, want thinking", v.ID)
	}
	if v.Headers["anthropic-beta"] != "output-128k" {
		t.Errorf("Headers[anthropic-beta] = %q, want output-128k", v.Headers["anthropic-beta"])
	}
	thinking, ok := v.Body["thinking"].(map[string]any)
	if !ok {
		t.Fatalf("Body[thinking] should be a nested object, got %T", v.Body["thinking"])
	}
	if thinking["budget_tokens"] != float64(10000) {
		t.Errorf("Body[thinking][budget_tokens] = %v, want 10000", thinking["budget_tokens"])
	}
}

// TestV2ModelInfoAPIUnionNilWhenAbsent verifies AsAPIUnion returns nil when the
// api field is absent, mirroring the V2AgentInfo.AsColorUnion nil-safety contract.
// V2ModelInfo is a plain response struct holding a union-typed `api` field — it is
// not a union root — so AsAPIUnion (not AsUnion) is the correct accessor name.
func TestV2ModelInfoAPIUnionNilWhenAbsent(t *testing.T) {
	raw := `{
		"id": "no-api",
		"providerID": "p",
		"name": "No API",
		"capabilities": {"tools": false, "input": [], "output": []},
		"request": {"headers": {}, "body": {}},
		"variants": [],
		"time": {"released": 0},
		"cost": [],
		"status": "alpha",
		"enabled": false,
		"limit": {"context": 0, "output": 0}
	}`
	var m opencode.V2ModelInfo
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatalf("unmarshal without api: %v", err)
	}
	if m.AsAPIUnion() != nil {
		t.Errorf("AsAPIUnion() = %v, want nil when api absent", m.AsAPIUnion())
	}
}

// TestV2ModelInfoAPINullDoesNotAbortDecode covers the four boundary shapes of the
// union-carrying `api` field: a legal value, an explicit null, an absent field, and
// a wrong-typed scalar.
//
// OpenAPI ModelV2Info lists `api` in required and resolves it to ModelApi, whose
// anyOf holds only the two object variants — there is no `{"type":"null"}` member,
// so a null is not a legal server response. It must still not abort the decode:
// neither registered variant filters on gjson.Null, so routing a null through the
// union decoder fails with "was not able to coerce type as union", and because the
// error propagates out of V2ModelInfo.UnmarshalJSON the surrounding document dies
// with it — inside V2ModelListResponse that silently empties the whole `data` array.
// Same forward-compatibility contract as ConfigMcp.UnmarshalJSON's null `oauth`.
func TestV2ModelInfoAPINullDoesNotAbortDecode(t *testing.T) {
	const scaffold = `"id":"m1","providerID":"p","name":"Model One"`

	t.Run("legal api value still routes to its variant", func(t *testing.T) {
		var m opencode.V2ModelInfo
		if err := json.Unmarshal([]byte(`{`+scaffold+`,"api":{"id":"a","type":"native","settings":{}}}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := m.API.(opencode.V2ModelInfoAPINative); !ok {
			t.Errorf("API = %T, want opencode.V2ModelInfoAPINative", m.API)
		}
		if m.ID != "m1" || m.Name != "Model One" {
			t.Errorf("sibling fields lost: id=%q name=%q", m.ID, m.Name)
		}
	})

	t.Run("explicit null is treated as absent", func(t *testing.T) {
		var m opencode.V2ModelInfo
		if err := json.Unmarshal([]byte(`{`+scaffold+`,"api":null}`), &m); err != nil {
			t.Fatalf("api:null must not fail the decode, got: %v", err)
		}
		if m.API != nil {
			t.Errorf("API = %#v, want nil for a null api", m.API)
		}
		if m.AsAPIUnion() != nil {
			t.Errorf("AsAPIUnion() = %v, want nil for a null api", m.AsAPIUnion())
		}
		// The whole point: the siblings must survive.
		if m.ID != "m1" || m.ProviderID != "p" || m.Name != "Model One" {
			t.Errorf("sibling fields lost: id=%q providerID=%q name=%q", m.ID, m.ProviderID, m.Name)
		}
	})

	t.Run("absent api leaves the carrier nil", func(t *testing.T) {
		var m opencode.V2ModelInfo
		if err := json.Unmarshal([]byte(`{`+scaffold+`}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if m.API != nil || m.AsAPIUnion() != nil {
			t.Errorf("API = %#v / AsAPIUnion = %v, want both nil", m.API, m.AsAPIUnion())
		}
	})

	t.Run("null api inside the list envelope keeps every element", func(t *testing.T) {
		raw := `{"location":{"directory":"/d","project":{"id":"pr","directory":"/d"}},"data":[` +
			`{"id":"m1","providerID":"p","name":"One","api":null},` +
			`{"id":"m2","providerID":"p","name":"Two","api":{"id":"a","type":"aisdk","package":"pkg"}}]}`
		var resp opencode.V2ModelListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal envelope: %v", err)
		}
		// Before the null guard this was len 0: the element error aborted the slice
		// decode and the outer struct decoder swallowed it, so `data` came back empty
		// with a nil error.
		if len(resp.Data) != 2 {
			t.Fatalf("len(Data) = %d, want 2 — a null `api` must not drop list elements", len(resp.Data))
		}
		if resp.Data[0].ID != "m1" || resp.Data[0].Name != "One" {
			t.Errorf("Data[0] = {id:%q name:%q}, want {m1 One}", resp.Data[0].ID, resp.Data[0].Name)
		}
		if resp.Data[0].API != nil {
			t.Errorf("Data[0].API = %#v, want nil", resp.Data[0].API)
		}
		if _, ok := resp.Data[1].API.(opencode.V2ModelInfoAPIAisdk); !ok {
			t.Errorf("Data[1].API = %T, want opencode.V2ModelInfoAPIAisdk", resp.Data[1].API)
		}
	})

	t.Run("wrong-typed api is reported through field metadata, not a fatal error", func(t *testing.T) {
		// OpenAPI admits no scalar variant for ModelApi, so `api:"text"` is an
		// illegal server response. Because `api` is declared as the registered
		// [V2ModelInfoAPIUnion], the framework's struct decoder records the coercion
		// failure on the field's metadata and keeps decoding the rest of the
		// document, instead of aborting it. The "null == absent" vs "wrong type"
		// distinction survives — it just moved from the returned error onto
		// JSON.API, which is strictly more informative and cannot destroy siblings.
		var m opencode.V2ModelInfo
		if err := json.Unmarshal([]byte(`{`+scaffold+`,"api":"text"}`), &m); err != nil {
			t.Fatalf("a wrong-typed api must not fail the whole decode, got: %v", err)
		}
		if m.API != nil {
			t.Errorf("API = %#v, want nil for a non-object api", m.API)
		}
		if !m.JSON.API.IsInvalid() {
			t.Error("JSON.API.IsInvalid() = false, want true for a non-object api")
		}
		if m.JSON.API.IsNull() || m.JSON.API.IsMissing() {
			t.Error("a wrong-typed api must stay distinguishable from null/absent")
		}
		if got := m.JSON.API.Raw(); got != `"text"` {
			t.Errorf("JSON.API.Raw() = %s, want %q — the raw value must stay available", got, `"text"`)
		}
		if m.ID != "m1" || m.ProviderID != "p" || m.Name != "Model One" {
			t.Errorf("sibling fields lost: id=%q providerID=%q name=%q", m.ID, m.ProviderID, m.Name)
		}
	})

	t.Run("wrong-typed api inside the list envelope keeps every element", func(t *testing.T) {
		// The regression this locks down: routing the sub-document by hand inside
		// UnmarshalJSON returned the coercion error out of the element decoder, the
		// slice decoder propagated it, and the outer struct decoder swallowed it —
		// so `data` came back empty with a nil error, losing the valid element too.
		raw := `{"location":{"directory":"/d","project":{"id":"pr","directory":"/d"}},"data":[` +
			`{"id":"m1","providerID":"p","name":"One","api":"text"},` +
			`{"id":"m2","providerID":"p","name":"Two","api":{"id":"a","type":"aisdk","package":"pkg"}}]}`
		var resp opencode.V2ModelListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal envelope: %v", err)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("len(Data) = %d, want 2 — a wrong-typed `api` must not drop list elements", len(resp.Data))
		}
		if resp.Data[0].ID != "m1" || resp.Data[0].Name != "One" || resp.Data[0].API != nil {
			t.Errorf("Data[0] = {id:%q name:%q api:%#v}, want {m1 One <nil>}", resp.Data[0].ID, resp.Data[0].Name, resp.Data[0].API)
		}
		if _, ok := resp.Data[1].API.(opencode.V2ModelInfoAPIAisdk); !ok {
			t.Errorf("Data[1].API = %T, want opencode.V2ModelInfoAPIAisdk", resp.Data[1].API)
		}
	})

	t.Run("unknown future variant falls back without an error", func(t *testing.T) {
		// Forward compatibility: a `type` the SDK has never heard of must not fail
		// the decode. It coerces loosely onto a registered variant whose IsKnown()
		// reports false, which is the documented way to detect it.
		var m opencode.V2ModelInfo
		if err := json.Unmarshal([]byte(`{`+scaffold+`,"api":{"id":"a","type":"future","package":"pkg"}}`), &m); err != nil {
			t.Fatalf("an unknown api variant must not fail the decode, got: %v", err)
		}
		if m.API == nil {
			t.Fatal("API = nil, want a loose fallback variant for an unknown type")
		}
		aisdk, ok := m.API.(opencode.V2ModelInfoAPIAisdk)
		if !ok {
			t.Fatalf("API = %T, want opencode.V2ModelInfoAPIAisdk as the loose fallback", m.API)
		}
		if aisdk.Type.IsKnown() {
			t.Errorf("Type.IsKnown() = true for %q, want false so callers can detect it", aisdk.Type)
		}
		if m.ID != "m1" || m.Name != "Model One" {
			t.Errorf("sibling fields lost: id=%q name=%q", m.ID, m.Name)
		}
	})
}

// TestV2ModelInfoAPISettingsIsMap verifies that the `settings` field on both
// ModelApi anyOf variants decodes as a concrete map[string]any.
//
// OpenAPI ModelApi anyOf[0].properties.settings and anyOf[1].properties.settings
// are both `{"type":"object"}` with no anyOf/oneOf, and JS SDK v2 types them as
// `settings?: { [key: string]: unknown }` (aisdk) / `settings: { [key: string]:
// unknown }` (native). They are therefore plain objects — not unions — so the
// declared Go type is map[string]any with no runtime-type comment.
//
// Required-ness per OpenAPI: aisdk required=[id,type,package] => settings optional;
// native required=[id,type,settings] => settings required.
func TestV2ModelInfoAPISettingsIsMap(t *testing.T) {
	t.Run("aisdk populated settings", func(t *testing.T) {
		raw := `{"id":"openai/gpt-4","type":"aisdk","package":"@ai-sdk/openai","url":"https://api.openai.com","settings":{"organization":"org_1","maxRetries":3}}`
		var a opencode.V2ModelInfoAPIAisdk
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		// Concrete map — indexable without a type assertion.
		if a.Settings["organization"] != "org_1" {
			t.Errorf("Settings[organization] = %v, want org_1", a.Settings["organization"])
		}
		if a.Settings["maxRetries"] != float64(3) {
			t.Errorf("Settings[maxRetries] = %v, want 3", a.Settings["maxRetries"])
		}
		if got := reflect.TypeOf(a.Settings); got.String() != "map[string]interface {}" {
			t.Errorf("reflect.TypeOf(Settings) = %s, want map[string]interface {}", got)
		}
	})

	t.Run("aisdk empty settings object", func(t *testing.T) {
		raw := `{"id":"i","type":"aisdk","package":"p","settings":{}}`
		var a opencode.V2ModelInfoAPIAisdk
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
		raw := `{"id":"i","type":"aisdk","package":"p"}`
		var a opencode.V2ModelInfoAPIAisdk
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if a.Settings != nil {
			t.Errorf("Settings should be nil when absent, got %v", a.Settings)
		}
		if a.Package != "p" {
			t.Errorf("Package = %q, want p", a.Package)
		}
	})

	t.Run("aisdk settings null", func(t *testing.T) {
		raw := `{"id":"i","type":"aisdk","package":"p","settings":null}`
		var a opencode.V2ModelInfoAPIAisdk
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if a.Settings != nil {
			t.Errorf("Settings should be nil for null, got %v", a.Settings)
		}
	})

	t.Run("native nested settings", func(t *testing.T) {
		raw := `{"id":"ollama/local","type":"native","url":"http://localhost:11434","settings":{"options":{"num_ctx":8192,"stop":["</s>"]},"keep_alive":"5m"}}`
		var n opencode.V2ModelInfoAPINative
		if err := json.Unmarshal([]byte(raw), &n); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if n.Settings["keep_alive"] != "5m" {
			t.Errorf("Settings[keep_alive] = %v, want 5m", n.Settings["keep_alive"])
		}
		opts, ok := n.Settings["options"].(map[string]any)
		if !ok {
			t.Fatalf("Settings[options] should be a nested object, got %T", n.Settings["options"])
		}
		if opts["num_ctx"] != float64(8192) {
			t.Errorf("Settings[options][num_ctx] = %v, want 8192", opts["num_ctx"])
		}
		stop, ok := opts["stop"].([]any)
		if !ok || len(stop) != 1 || stop[0] != "</s>" {
			t.Errorf("Settings[options][stop] = %v, want [</s>]", opts["stop"])
		}
		if got := reflect.TypeOf(n.Settings); got.String() != "map[string]interface {}" {
			t.Errorf("reflect.TypeOf(Settings) = %s, want map[string]interface {}", got)
		}
	})

	t.Run("native empty settings object", func(t *testing.T) {
		raw := `{"id":"i","type":"native","settings":{}}`
		var n opencode.V2ModelInfoAPINative
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
		raw := `{
			"id":"m","providerID":"p","name":"N",
			"api":{"id":"ollama/local","type":"native","settings":{"keep_alive":"10m"}},
			"capabilities":{"tools":false,"input":["text"],"output":["text"]},
			"request":{"headers":{},"body":{}},
			"variants":[],"time":{"released":1},"cost":[],
			"status":"active","enabled":true,"limit":{"context":1,"output":1}
		}`
		var m opencode.V2ModelInfo
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		native, ok := m.AsAPIUnion().(opencode.V2ModelInfoAPINative)
		if !ok {
			t.Fatalf("expected V2ModelInfoAPINative, got %T", m.AsAPIUnion())
		}
		if native.Settings["keep_alive"] != "10m" {
			t.Errorf("Settings[keep_alive] = %v, want 10m", native.Settings["keep_alive"])
		}
	})
}
