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

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestProviderList(t *testing.T) {
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
	_, err := client.Provider.List(context.TODO(), opencode.ProviderListParams{
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

func TestProviderAuth(t *testing.T) {
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
	_, err := client.Provider.Auth(context.TODO(), opencode.ProviderAuthParams{
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

func TestProviderOauthAuthorize(t *testing.T) {
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
	_, err := client.Provider.Oauth.Authorize(context.TODO(), "providerID", opencode.ProviderOauthAuthorizeParams{
		Method: opencode.F(int64(0)),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProviderOauthCallback(t *testing.T) {
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
	_, err := client.Provider.Oauth.Callback(context.TODO(), "providerID", opencode.ProviderOauthCallbackParams{
		Method: opencode.F(int64(0)),
		Code:   opencode.F("code"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// ===== ProviderAuthMethod.prompts array-of-union routing =====

// OpenAPI ProviderAuthMethod.properties.prompts is
//
//	{"type":"array","items":{"anyOf":[<text object>,<select object>]}}
//
// and JS SDK v2 types it as `prompts?: Array<{type:"text",...} | {type:"select",...}>`.
//
// The anyOf sits on `items`, so a single array may legally mix both element kinds.
// The carrier is `any` per the Union field rules; [apijson] keys its union registry
// on the interface's [reflect.Type], so a bare `any` carrier never matches an entry
// and the array decoder would yield []any of map[string]any. UnmarshalJSON routes
// each element through ProviderAuthMethodPrompt, which is what the declared
// [[]ProviderAuthMethodPrompt] runtime type denotes.
func TestProviderAuthMethodPromptsRoutingMatrix(t *testing.T) {
	text := `{"type":"text","key":"api_key","message":"Enter API key","placeholder":"sk-..."}`
	sel := `{"type":"select","key":"region","message":"Choose region","options":[{"label":"US","value":"us","hint":"United States"},{"label":"EU","value":"eu"}]}`

	for _, tc := range []struct {
		name    string
		prompts string
		want    []any
	}{
		{"empty array", `[]`, []any{}},
		{"single text", `[` + text + `]`, []any{opencode.ProviderAuthMethodPromptText{}}},
		{"single select", `[` + sel + `]`, []any{opencode.ProviderAuthMethodPromptSelect{}}},
		{
			"mixed text+select",
			`[` + text + `,` + sel + `]`,
			[]any{opencode.ProviderAuthMethodPromptText{}, opencode.ProviderAuthMethodPromptSelect{}},
		},
		{
			"mixed select+text (order preserved)",
			`[` + sel + `,` + text + `]`,
			[]any{opencode.ProviderAuthMethodPromptSelect{}, opencode.ProviderAuthMethodPromptText{}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			raw := `{"type":"oauth","label":"L","prompts":` + tc.prompts + `}`
			var m opencode.ProviderAuthMethod
			if err := json.Unmarshal([]byte(raw), &m); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			prompts := m.Prompts
			if len(prompts) != len(tc.want) {
				t.Fatalf("len(Prompts) = %d, want %d", len(prompts), len(tc.want))
			}
			for i, want := range tc.want {
				if reflect.TypeOf(prompts[i]) != reflect.TypeOf(want) {
					t.Errorf("Prompts[%d] = %s, want %s", i, reflect.TypeOf(prompts[i]), reflect.TypeOf(want))
				}
			}
			if got := m.AsPromptsUnion(); !reflect.DeepEqual(got, prompts) {
				t.Error("AsPromptsUnion() disagrees with Prompts carrier")
			}
			if m.JSON.RawJSON() == "" {
				t.Error("RawJSON not preserved")
			}
		})
	}

	t.Run("prompts absent (optional per OpenAPI)", func(t *testing.T) {
		var m opencode.ProviderAuthMethod
		if err := json.Unmarshal([]byte(`{"type":"api","label":"API Key"}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if m.Prompts != nil {
			t.Errorf("Prompts = %v (%s), want nil", m.Prompts, reflect.TypeOf(m.Prompts))
		}
		if m.AsPromptsUnion() != nil {
			t.Error("AsPromptsUnion() should be nil when prompts absent")
		}
	})

	t.Run("prompts null", func(t *testing.T) {
		var m opencode.ProviderAuthMethod
		if err := json.Unmarshal([]byte(`{"type":"api","label":"API Key","prompts":null}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if m.Prompts != nil {
			t.Errorf("Prompts = %v (%s), want nil", m.Prompts, reflect.TypeOf(m.Prompts))
		}
	})

	t.Run("nested fields survive routing", func(t *testing.T) {
		raw := `{"type":"api","label":"Custom","prompts":[
			{"type":"text","key":"k","message":"M","placeholder":"P","when":{"key":"env","op":"neq","value":"dev"}},
			{"type":"select","key":"mode","message":"Mode","options":[
				{"label":"Auto","value":"auto"},{"label":"Manual","value":"manual","hint":"H"}],
			 "when":{"key":"env","op":"eq","value":"prod"}}
		]}`
		var m opencode.ProviderAuthMethod
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		prompts := m.AsPromptsUnion()
		if len(prompts) != 2 {
			t.Fatalf("len = %d, want 2", len(prompts))
		}
		tp, ok := prompts[0].(opencode.ProviderAuthMethodPromptText)
		if !ok {
			t.Fatalf("prompts[0] = %s", reflect.TypeOf(prompts[0]))
		}
		if tp.Key != "k" || tp.Message != "M" || tp.Placeholder != "P" {
			t.Errorf("prompts[0] = %+v", tp)
		}
		if tp.When.Key != "env" || tp.When.Op != opencode.ProviderAuthMethodPromptWhenOpNeq || tp.When.Value != "dev" {
			t.Errorf("prompts[0].When = %+v", tp.When)
		}
		sp, ok := prompts[1].(opencode.ProviderAuthMethodPromptSelect)
		if !ok {
			t.Fatalf("prompts[1] = %s", reflect.TypeOf(prompts[1]))
		}
		if len(sp.Options) != 2 || sp.Options[1].Value != "manual" || sp.Options[1].Hint != "H" {
			t.Errorf("prompts[1].Options = %+v", sp.Options)
		}
		if sp.When.Op != opencode.ProviderAuthMethodPromptWhenOpEq {
			t.Errorf("prompts[1].When = %+v", sp.When)
		}
		if sp.JSON.RawJSON() == "" {
			t.Error("prompts[1].RawJSON() empty")
		}
	})
}

// ProviderAuthMethod is returned as map[string][]ProviderAuthMethod by
// ProviderService.Auth; routing must survive the map + array layers.
func TestProviderAuthMethodPromptsThroughAuthEnvelope(t *testing.T) {
	raw := `{
		"anthropic": [
			{"type":"oauth","label":"Claude Pro"},
			{"type":"api","label":"API Key","prompts":[{"type":"text","key":"key","message":"Key"}]}
		],
		"openai": [
			{"type":"api","label":"API Key","prompts":[
				{"type":"select","key":"org","message":"Org","options":[{"label":"Default","value":"d"}]},
				{"type":"text","key":"key","message":"Key"}
			]}
		]
	}`
	var got map[string][]opencode.ProviderAuthMethod
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("providers = %d, want 2", len(got))
	}
	anthropic := got["anthropic"]
	if len(anthropic) != 2 {
		t.Fatalf("anthropic methods = %d, want 2", len(anthropic))
	}
	if anthropic[0].Prompts != nil {
		t.Errorf("anthropic[0].Prompts = %v, want nil", anthropic[0].Prompts)
	}
	first := anthropic[1].Prompts
	if len(first) != 1 || reflect.TypeOf(first[0]) != reflect.TypeOf(opencode.ProviderAuthMethodPromptText{}) {
		t.Errorf("anthropic[1].Prompts = %#v", first)
	}
	openai := got["openai"]
	if len(openai) != 1 {
		t.Fatalf("openai methods = %d, want 1", len(openai))
	}
	second := openai[0].AsPromptsUnion()
	if len(second) != 2 {
		t.Fatalf("openai[0] prompts = %d, want 2", len(second))
	}
	if reflect.TypeOf(second[0]) != reflect.TypeOf(opencode.ProviderAuthMethodPromptSelect{}) {
		t.Errorf("openai[0].Prompts[0] = %s, want ProviderAuthMethodPromptSelect", reflect.TypeOf(second[0]))
	}
	if reflect.TypeOf(second[1]) != reflect.TypeOf(opencode.ProviderAuthMethodPromptText{}) {
		t.Errorf("openai[0].Prompts[1] = %s, want ProviderAuthMethodPromptText", reflect.TypeOf(second[1]))
	}
}

// Regression: the ProviderAuthMethodPrompt union resolves by its `type`
// discriminator, independently of the routing above.
func TestProviderAuthMethodPromptUnionDiscriminator(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		want any
	}{
		{`{"type":"text","key":"k","message":"m"}`, opencode.ProviderAuthMethodPromptText{}},
		{`{"type":"text","key":"k","message":"m","placeholder":"p"}`, opencode.ProviderAuthMethodPromptText{}},
		{`{"type":"select","key":"k","message":"m","options":[]}`, opencode.ProviderAuthMethodPromptSelect{}},
		{`{"type":"select","key":"k","message":"m","options":[{"label":"L","value":"V"}]}`, opencode.ProviderAuthMethodPromptSelect{}},
		// Unknown extra properties must not change the resolved variant.
		{`{"type":"text","key":"k","message":"m","zz":1}`, opencode.ProviderAuthMethodPromptText{}},
		{`{"type":"select","key":"k","message":"m","options":[],"zz":1}`, opencode.ProviderAuthMethodPromptSelect{}},
	} {
		var u opencode.ProviderAuthMethodPrompt
		if err := apijson.UnmarshalRoot([]byte(tc.raw), &u); err != nil {
			t.Fatalf("%s: %v", tc.raw, err)
		}
		if reflect.TypeOf(u) != reflect.TypeOf(tc.want) {
			t.Errorf("%s -> %s, want %s", tc.raw, reflect.TypeOf(u), reflect.TypeOf(tc.want))
		}
	}
}

// TestProviderAuthMethodPromptsNullElement pins the illegal-element boundary of
// ProviderAuthMethod.prompts.
//
// OpenAPI ProviderAuthMethod.prompts is `array` whose `items` is an inline anyOf of
// the text and select prompt objects; neither member is `{"type":"null"}`, so a null
// element is not a legal server response. Because `prompts` is declared as
// []ProviderAuthMethodPrompt, internal/apijson decodes it field-by-field: the array
// decoder rejects the illegal element, and the struct decoder records the field as
// invalid and moves on instead of propagating the failure. The observable contract
// is therefore: no error, the whole `prompts` field degrades to nil, every sibling
// field survives, and the payload stays verbatim in RawJSON.
func TestProviderAuthMethodPromptsNullElement(t *testing.T) {
	t.Run("null element degrades the field without failing the decode", func(t *testing.T) {
		raw := `{"type":"api","label":"API Key","prompts":[null,{"type":"text","key":"k","message":"Key"}]}`
		var m opencode.ProviderAuthMethod
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("null prompts element must not fail the decode, got: %v", err)
		}
		if m.Label != "API Key" || m.Type != opencode.ProviderAuthMethodTypeAPI {
			t.Errorf("sibling fields lost: type=%q label=%q", m.Type, m.Label)
		}
		if got := m.AsPromptsUnion(); len(got) != 0 {
			t.Errorf("len(AsPromptsUnion()) = %d, want 0 (illegal element degrades the field)", len(got))
		}
		// The whole payload, including the illegal element, is still recoverable.
		if !strings.Contains(m.JSON.RawJSON(), "null") {
			t.Error("RawJSON() should still carry the rejected null element")
		}
	})

	t.Run("null element inside the auth envelope keeps every provider", func(t *testing.T) {
		raw := `{"anthropic":[{"type":"oauth","label":"Claude","prompts":[null]}],` +
			`"openai":[{"type":"api","label":"Key","prompts":[{"type":"text","key":"k","message":"K"}]}]}`
		var got map[string][]opencode.ProviderAuthMethod
		if err := json.Unmarshal([]byte(raw), &got); err != nil {
			t.Fatalf("unmarshal envelope: %v", err)
		}
		if len(got) != 2 {
			t.Fatalf("providers = %d, want 2 — a null prompt must not drop providers", len(got))
		}
		if len(got["anthropic"]) != 1 || len(got["anthropic"][0].AsPromptsUnion()) != 0 {
			t.Errorf("anthropic = %#v, want 1 method with 0 prompts", got["anthropic"])
		}
		if len(got["openai"][0].AsPromptsUnion()) != 1 {
			t.Errorf("openai prompts = %d, want 1", len(got["openai"][0].AsPromptsUnion()))
		}
	})

	t.Run("wrong-typed element degrades the field, siblings survive", func(t *testing.T) {
		var m opencode.ProviderAuthMethod
		if err := json.Unmarshal([]byte(`{"type":"api","label":"L","prompts":["text"]}`), &m); err != nil {
			t.Fatalf(`prompts:["text"] must not fail the decode, got: %v`, err)
		}
		if m.Type != opencode.ProviderAuthMethodTypeAPI || m.Label != "L" {
			t.Errorf("sibling fields lost: type=%q label=%q", m.Type, m.Label)
		}
		if got := m.AsPromptsUnion(); len(got) != 0 {
			t.Errorf("len(AsPromptsUnion()) = %d, want 0", len(got))
		}
	})

	t.Run("scalar prompts degrades the field, siblings survive", func(t *testing.T) {
		// Previously the `any` carrier leaked the raw scalar; the typed field cannot.
		var m opencode.ProviderAuthMethod
		if err := json.Unmarshal([]byte(`{"type":"api","label":"L","prompts":"nope"}`), &m); err != nil {
			t.Fatalf(`prompts:"nope" must not fail the decode, got: %v`, err)
		}
		if m.Label != "L" {
			t.Errorf("sibling field lost: label=%q", m.Label)
		}
		if m.Prompts != nil {
			t.Errorf("Prompts = %#v, want nil for a non-array payload", m.Prompts)
		}
	})
}

// TestProviderModelCapabilitiesInterleavedUnion verifies that the `interleaved`
// anyOf actually reaches its registered variants.
//
// OpenAPI Model.capabilities.interleaved is
// `anyOf [boolean, {field: enum[reasoning, reasoning_content, reasoning_details]}]`
// and is required. The carrier is `any` per the Union field rules, but internal
// /apijson keys its union registry on the interface's reflect.Type, and a bare `any`
// is `interface {}` — it never matches a registry entry, so without explicit routing
// the generic interface branch wins and yields a plain `bool` / `map[string]any`,
// contradicting the declared runtime types and leaving the RegisterUnion call dead.
func TestProviderModelCapabilitiesInterleavedUnion(t *testing.T) {
	scaffold := func(interleaved string) string {
		return `{"temperature":true,"reasoning":true,"attachment":true,"toolcall":true,` +
			`"input":{"text":true,"audio":false,"image":false,"video":false,"pdf":false},` +
			`"output":{"text":true,"audio":false,"image":false,"video":false,"pdf":false},` +
			`"interleaved":` + interleaved + `}`
	}

	for _, tc := range []struct {
		name    string
		raw     string
		want    any
		wantVal any
	}{
		{"scalar true", `true`, opencode.UnionBool(true), opencode.UnionBool(true)},
		{"scalar false", `false`, opencode.UnionBool(false), opencode.UnionBool(false)},
		{
			"object form", `{"field":"reasoning_content"}`,
			opencode.ProviderModelCapabilitiesInterleavedField{},
			nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var caps opencode.ProviderModelCapabilities
			if err := json.Unmarshal([]byte(scaffold(tc.raw)), &caps); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got, want := reflect.TypeOf(caps.Interleaved), reflect.TypeOf(tc.want); got != want {
				t.Errorf("Interleaved = %s, want %s", got, want)
			}
			if got, want := reflect.TypeOf(caps.AsInterleavedUnion()), reflect.TypeOf(tc.want); got != want {
				t.Errorf("AsInterleavedUnion() = %s, want %s", got, want)
			}
			if tc.wantVal != nil && caps.Interleaved != tc.wantVal {
				t.Errorf("Interleaved value = %#v, want %#v", caps.Interleaved, tc.wantVal)
			}
			if f, ok := caps.Interleaved.(opencode.ProviderModelCapabilitiesInterleavedField); ok {
				if f.Field != opencode.ProviderModelCapabilitiesInterleavedFieldFieldReasoningContent {
					t.Errorf("Field = %q, want reasoning_content", f.Field)
				}
				if !f.Field.IsKnown() {
					t.Error("reasoning_content should be a known field")
				}
			}
			// Sibling fields must survive the extra routing pass.
			if !caps.Temperature || !caps.Input.Text || caps.Input.Audio {
				t.Errorf("sibling fields lost: %+v", caps)
			}
		})
	}

	t.Run("reachable through the full provider list envelope", func(t *testing.T) {
		raw := `{"all":[{"id":"anthropic","name":"Anthropic","source":"env","env":[],"options":{},"models":{"m":{` +
			`"id":"m","providerID":"anthropic","api":{"id":"a","url":"u","npm":"n"},"name":"M",` +
			`"capabilities":{"temperature":true,"reasoning":true,"attachment":false,"toolcall":true,` +
			`"input":{"text":true,"audio":false,"image":false,"video":false,"pdf":false},` +
			`"output":{"text":true,"audio":false,"image":false,"video":false,"pdf":false},` +
			`"interleaved":{"field":"reasoning"}},` +
			`"cost":{"input":1,"output":2,"cache":{"read":0,"write":0}},` +
			`"limit":{"context":100,"output":50},"status":"active","options":{},"headers":{},` +
			`"release_date":"2024-01-01"}}}],"default":{},"connected":[]}`
		var resp opencode.ProviderListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(resp.All) != 1 {
			t.Fatalf("len(All) = %d, want 1", len(resp.All))
		}
		caps := resp.All[0].Models["m"].Capabilities
		f, ok := caps.Interleaved.(opencode.ProviderModelCapabilitiesInterleavedField)
		if !ok {
			t.Fatalf("Interleaved = %s, want ProviderModelCapabilitiesInterleavedField", reflect.TypeOf(caps.Interleaved))
		}
		if f.Field != opencode.ProviderModelCapabilitiesInterleavedFieldFieldReasoning {
			t.Errorf("Field = %q, want reasoning", f.Field)
		}
	})

	t.Run("null interleaved does not abort the decode", func(t *testing.T) {
		// `interleaved` is required with no null member, so a null is not legal — but
		// it must not take the whole document down with it.
		var caps opencode.ProviderModelCapabilities
		if err := json.Unmarshal([]byte(scaffold(`null`)), &caps); err != nil {
			t.Fatalf("interleaved:null must not fail the decode, got: %v", err)
		}
		if caps.Interleaved != nil || caps.AsInterleavedUnion() != nil {
			t.Errorf("Interleaved = %#v / AsInterleavedUnion = %v, want both nil", caps.Interleaved, caps.AsInterleavedUnion())
		}
		if !caps.Temperature {
			t.Error("sibling fields lost on a null interleaved")
		}
	})

	t.Run("scalar interleaved does not abort the decode", func(t *testing.T) {
		// A string matches neither anyOf member. Declaring the field as the union
		// interface lets the struct decoder record it as invalid and keep going, so
		// the surrounding document survives.
		var caps opencode.ProviderModelCapabilities
		if err := json.Unmarshal([]byte(scaffold(`"yes"`)), &caps); err != nil {
			t.Fatalf(`interleaved:"yes" must not fail the decode, got: %v`, err)
		}
		if caps.Interleaved != nil || caps.AsInterleavedUnion() != nil {
			t.Errorf("Interleaved = %#v, want nil", caps.Interleaved)
		}
		if !caps.Temperature || !caps.Input.Text {
			t.Error("sibling fields lost on a scalar interleaved")
		}
	})
}

// TestProviderUnionVariantCoverage asserts that every variant registered for the
// two provider.go unions is actually reachable, counted after de-duplicating by
// concrete [reflect.Type].
//
//   - ProviderModelCapabilitiesInterleavedUnion registers 3 UnionVariant entries
//     but only 2 distinct types: shared.UnionBool is registered twice (once under
//     gjson.True, once under gjson.False, because the decoder filters on the
//     concrete gjson token kind) plus
//     ProviderModelCapabilitiesInterleavedField under gjson.JSON.
//   - ProviderAuthMethodPrompt registers 2 UnionVariant entries / 2 distinct
//     types, dispatched by the `type` discriminator.
//
// Each case is exercised twice: once with the exact OpenAPI-legal payload and
// once with an unknown extra property appended, since an unrecognised field must
// never change which variant wins.
func TestProviderUnionVariantCoverage(t *testing.T) {
	t.Run("ProviderModelCapabilitiesInterleavedUnion", func(t *testing.T) {
		scaffold := func(interleaved string) string {
			return `{"temperature":true,"reasoning":true,"attachment":true,"toolcall":true,` +
				`"input":{"text":true,"audio":false,"image":false,"video":false,"pdf":false},` +
				`"output":{"text":true,"audio":false,"image":false,"video":false,"pdf":false},` +
				`"interleaved":` + interleaved + `}`
		}

		cases := []struct {
			name string
			raw  string
			want any
		}{
			{"gjson.True -> UnionBool", `true`, opencode.UnionBool(true)},
			{"gjson.False -> UnionBool", `false`, opencode.UnionBool(false)},
			{"gjson.JSON -> InterleavedField", `{"field":"reasoning"}`, opencode.ProviderModelCapabilitiesInterleavedField{}},
			{"gjson.JSON + unknown extra field", `{"field":"reasoning_details","zz":1}`, opencode.ProviderModelCapabilitiesInterleavedField{}},
		}

		reached := map[reflect.Type]bool{}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				var caps opencode.ProviderModelCapabilities
				if err := json.Unmarshal([]byte(scaffold(tc.raw)), &caps); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				got := reflect.TypeOf(caps.AsInterleavedUnion())
				if want := reflect.TypeOf(tc.want); got != want {
					t.Fatalf("AsInterleavedUnion() = %s, want %s", got, want)
				}
				if carrier := reflect.TypeOf(caps.Interleaved); carrier != got {
					t.Errorf("Interleaved carrier = %s, disagrees with union %s", carrier, got)
				}
				reached[got] = true
			})
		}
		// 3 registered UnionVariant entries collapse to 2 distinct types.
		if len(reached) != 2 {
			t.Errorf("distinct reached variants = %d, want 2 (%v)", len(reached), reached)
		}
		for _, want := range []any{
			opencode.UnionBool(false),
			opencode.ProviderModelCapabilitiesInterleavedField{},
		} {
			if !reached[reflect.TypeOf(want)] {
				t.Errorf("registered variant %s never reached", reflect.TypeOf(want))
			}
		}
	})

	t.Run("ProviderAuthMethodPrompt", func(t *testing.T) {
		cases := []struct {
			name string
			raw  string
			want any
		}{
			{"text discriminator", `{"type":"text","key":"k","message":"m"}`, opencode.ProviderAuthMethodPromptText{}},
			{"text + unknown extra field", `{"type":"text","key":"k","message":"m","zz":1}`, opencode.ProviderAuthMethodPromptText{}},
			{"select discriminator", `{"type":"select","key":"k","message":"m","options":[]}`, opencode.ProviderAuthMethodPromptSelect{}},
			{"select + unknown extra field", `{"type":"select","key":"k","message":"m","options":[],"zz":1}`, opencode.ProviderAuthMethodPromptSelect{}},
		}

		reached := map[reflect.Type]bool{}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				var u opencode.ProviderAuthMethodPrompt
				if err := apijson.UnmarshalRoot([]byte(tc.raw), &u); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				got := reflect.TypeOf(u)
				if want := reflect.TypeOf(tc.want); got != want {
					t.Fatalf("%s -> %s, want %s", tc.raw, got, want)
				}
				reached[got] = true
			})
		}
		// 2 registered UnionVariant entries, 2 distinct types.
		if len(reached) != 2 {
			t.Errorf("distinct reached variants = %d, want 2 (%v)", len(reached), reached)
		}
		for _, want := range []any{
			opencode.ProviderAuthMethodPromptText{},
			opencode.ProviderAuthMethodPromptSelect{},
		} {
			if !reached[reflect.TypeOf(want)] {
				t.Errorf("registered variant %s never reached", reflect.TypeOf(want))
			}
		}
	})
}
