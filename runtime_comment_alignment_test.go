// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/sst/opencode-sdk-go/internal/apijson"
)

// TestAssistantMessageStructuredIsArbitraryJSON asserts that
// `AssistantMessage.structured` accepts any JSON value, matching OpenAPI where the
// property is declared as an unconstrained schema (`{}`) and JS SDK v2 where it is
// typed `structured?: unknown`. It must NOT be constrained to the OutputFormat
// union (that union belongs to `UserMessage.format`).
func TestAssistantMessageStructuredIsArbitraryJSON(t *testing.T) {
	base := `"id":"msg_1","sessionID":"ses_1","role":"assistant","time":{"created":1},` +
		`"parentID":"","modelID":"m","providerID":"p","mode":"build","agent":"general",` +
		`"path":{"cwd":"/a","root":"/b"},"cost":0,` +
		`"tokens":{"input":1,"output":2,"reasoning":0,"cache":{"read":0,"write":0}}`

	cases := []struct {
		name string
		raw  string
		want string
	}{
		{"object", `{"a":1}`, "map[string]interface {}"},
		{"array", `[1,2,3]`, "[]interface {}"},
		{"string", `"hello"`, "string"},
		{"number", `42`, "float64"},
		{"bool", `true`, "bool"},
		{"null", `null`, "<nil>"},
		{"nested", `{"items":[{"k":"v"}],"n":1.5}`, "map[string]interface {}"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var m AssistantMessage
			payload := `{` + base + `,"structured":` + tc.raw + `}`
			if err := json.Unmarshal([]byte(payload), &m); err != nil {
				t.Fatalf("unmarshal %s: %v", tc.raw, err)
			}
			got := fmt.Sprintf("%T", m.Structured)
			if got != tc.want {
				t.Errorf("structured=%s: runtime type = %s, want %s", tc.raw, got, tc.want)
			}
			// Whatever the shape, the raw JSON must always be recoverable.
			if m.JSON.RawJSON() == "" {
				t.Error("RawJSON() is empty")
			}
		})
	}

	// Absent `structured` must stay nil (it is optional in OpenAPI).
	t.Run("absent", func(t *testing.T) {
		var m AssistantMessage
		if err := json.Unmarshal([]byte(`{`+base+`}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if m.Structured != nil {
			t.Errorf("absent structured = %#v, want nil", m.Structured)
		}
	})
}

// TestUserMessageFormatRuntimeTypes asserts `UserMessage.format` resolves to the
// OpenAPI `OutputFormat` union (`anyOf [OutputFormatText, OutputFormatJsonSchema]`,
// JS SDK v2 `OutputFormat = OutputFormatText | OutputFormatJsonSchema`), which is
// what the runtime comment on [UserMessage.Format] documents.
//
// OpenAPI references `OutputFormat` from three places, and the Go SDK splits them
// along the request/response boundary:
//   - request bodies of `POST /session/{sessionID}/message` and
//     `POST /session/{sessionID}/prompt_async` -> [SessionPromptParamsFormatUnion],
//     whose variants wrap every field in param.Field;
//   - response schema `UserMessage.format` -> [OutputFormatUnion], whose variants
//     [OutputFormatText] / [OutputFormatJsonSchema] use plain Go field types and
//     carry JSON metadata.
//
// Regression guard: [OutputFormatText] / [OutputFormatJsonSchema] used to be
// param.Field-based request types that no Params field could ever accept, so no
// response-side carrier existed at all and an object payload decoded generically to
// map[string]any -- callers lost every typed field. They are now the response-side
// variants, and the request side keeps its own [SessionPromptParamsFormatText] /
// [SessionPromptParamsFormatJsonSchema] spelling.
func TestUserMessageFormatRuntimeTypes(t *testing.T) {
	base := `"id":"msg_1","sessionID":"ses_1","role":"user","time":{"created":1},` +
		`"agent":"general","model":{"modelID":"m","providerID":"p"}`

	t.Run("text", func(t *testing.T) {
		var m UserMessage
		raw := `{` + base + `,"format":{"type":"text"}}`
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		got, ok := m.Format.(OutputFormatText)
		if !ok {
			t.Fatalf("format runtime type = %T, want OutputFormatText", m.Format)
		}
		if got.Type != OutputFormatTextTypeText || !got.Type.IsKnown() {
			t.Errorf("format.Type = %q (known=%v), want text", got.Type, got.Type.IsKnown())
		}
		if got.JSON.RawJSON() != `{"type":"text"}` {
			t.Errorf("variant RawJSON() = %q, want {\"type\":\"text\"}", got.JSON.RawJSON())
		}
		// AsFormat must agree with the mirrored `any` field.
		if _, ok := m.AsFormat().(OutputFormatText); !ok {
			t.Errorf("AsFormat() runtime type = %T, want OutputFormatText", m.AsFormat())
		}
	})

	t.Run("json_schema", func(t *testing.T) {
		var m UserMessage
		raw := `{` + base + `,"format":{"type":"json_schema","schema":{"type":"object"},"retryCount":3}}`
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		got, ok := m.Format.(OutputFormatJsonSchema)
		if !ok {
			t.Fatalf("format runtime type = %T, want OutputFormatJsonSchema", m.Format)
		}
		if got.Type != OutputFormatJsonSchemaTypeJsonSchema || !got.Type.IsKnown() {
			t.Errorf("format.Type = %q (known=%v), want json_schema", got.Type, got.Type.IsKnown())
		}
		// OpenAPI `$ref JSONSchema` is the open object `{"type": "object"}`, so the
		// payload must survive as a generic map rather than being dropped.
		if got.Schema["type"] != "object" {
			t.Errorf("format.Schema = %v, want map with type=object", got.Schema)
		}
		// OpenAPI declares retryCount as `integer` (minimum 0), so it must be int64.
		if got.RetryCount != 3 {
			t.Errorf("format.RetryCount = %d, want 3", got.RetryCount)
		}
		if _, ok := m.AsFormat().(OutputFormatJsonSchema); !ok {
			t.Errorf("AsFormat() runtime type = %T, want OutputFormatJsonSchema", m.AsFormat())
		}
	})

	// `retryCount` is optional in OpenAPI (not in OutputFormatJsonSchema.required),
	// so an absent value must decode to the int64 zero value, not fail the document.
	t.Run("json_schema without retryCount", func(t *testing.T) {
		var m UserMessage
		raw := `{` + base + `,"format":{"type":"json_schema","schema":{"type":"object"}}}`
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		got, ok := m.Format.(OutputFormatJsonSchema)
		if !ok {
			t.Fatalf("format runtime type = %T, want OutputFormatJsonSchema", m.Format)
		}
		if got.RetryCount != 0 {
			t.Errorf("format.RetryCount = %d, want 0", got.RetryCount)
		}
	})

	// `format` is not in `UserMessage.required`, so absent and null must both leave
	// the field nil without panicking.
	t.Run("absent and null", func(t *testing.T) {
		for name, raw := range map[string]string{
			"absent": `{` + base + `}`,
			"null":   `{` + base + `,"format":null}`,
		} {
			t.Run(name, func(t *testing.T) {
				var m UserMessage
				if err := json.Unmarshal([]byte(raw), &m); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				if m.Format != nil {
					t.Errorf("format = %#v, want nil", m.Format)
				}
				if m.AsFormat() != nil {
					t.Errorf("AsFormat() = %#v, want nil", m.AsFormat())
				}
			})
		}
	})

	// Forward compatibility: an unrecognised property must not shift the document
	// onto the wrong variant, and must stay reachable via RawJSON/ExtraFields.
	t.Run("unknown property", func(t *testing.T) {
		var m UserMessage
		inner := `{"type":"json_schema","schema":{"type":"object"},"zz":1}`
		if err := json.Unmarshal([]byte(`{`+base+`,"format":`+inner+`}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		got, ok := m.Format.(OutputFormatJsonSchema)
		if !ok {
			t.Fatalf("format runtime type = %T, want OutputFormatJsonSchema", m.Format)
		}
		if got.JSON.RawJSON() != inner {
			t.Errorf("variant RawJSON() = %q, want %q", got.JSON.RawJSON(), inner)
		}
		if _, ok := got.JSON.ExtraFields["zz"]; !ok {
			t.Errorf("variant ExtraFields = %v, want key \"zz\"", got.JSON.ExtraFields)
		}
	})

	// Guard the request/response split in the other direction: the param.Field-based
	// [SessionPromptParamsFormatText] / [SessionPromptParamsFormatJsonSchema] are
	// request types, so they must never appear as the runtime type of a response.
	t.Run("request-side param types never appear in a response", func(t *testing.T) {
		payloads := []string{
			`{"type":"text"}`,
			`{"type":"json_schema","schema":{"type":"object"}}`,
		}
		for _, p := range payloads {
			var m UserMessage
			if err := json.Unmarshal([]byte(`{`+base+`,"format":`+p+`}`), &m); err != nil {
				t.Fatalf("unmarshal %s: %v", p, err)
			}
			// Asserted through an `any` view: the request-side param types and the
			// generic map do not implement the union `Format` is now declared as, so a
			// direct assertion would not compile. Keeping the runtime guard means it
			// still fires if the field ever regresses to an untyped carrier.
			if _, ok := any(m.Format).(SessionPromptParamsFormatText); ok {
				t.Errorf("%s: response decoded into request type SessionPromptParamsFormatText", p)
			}
			if _, ok := any(m.Format).(SessionPromptParamsFormatJsonSchema); ok {
				t.Errorf("%s: response decoded into request type SessionPromptParamsFormatJsonSchema", p)
			}
			// The generic map is what the pre-carrier behaviour produced; it must be
			// gone now that the union is routed.
			if _, ok := any(m.Format).(map[string]any); ok {
				t.Errorf("%s: format decoded generically to map[string]any, union not routed", p)
			}
		}
	})

	// The `format` union must survive routing through [MessageUnion]. internal/apijson
	// skips the indirect unmarshaler for registered variants
	// (internal/apijson/decoder.go:145-149), so [UserMessage.UnmarshalJSON] never runs
	// on this path. It does not need to: `Format` is declared as [OutputFormatUnion],
	// so the field is routed by the struct decoder itself, and [Message.UnmarshalJSON]
	// is a plain carrier that ports the already-typed value onto [Message.Format].
	t.Run("through MessageUnion", func(t *testing.T) {
		cases := map[string]any{
			`{"type":"text"}`: OutputFormatText{},
			`{"type":"json_schema","schema":{"type":"object"}}`: OutputFormatJsonSchema{},
			`{"type":"json_schema","schema":{},"retryCount":9}`: OutputFormatJsonSchema{},
		}
		for p, want := range cases {
			var msg Message
			if err := json.Unmarshal([]byte(`{`+base+`,"format":`+p+`}`), &msg); err != nil {
				t.Fatalf("unmarshal %s: %v", p, err)
			}
			um, ok := msg.AsUnion().(UserMessage)
			if !ok {
				t.Fatalf("%s: union runtime type = %T, want UserMessage", p, msg.AsUnion())
			}
			if got, wantT := reflect.TypeOf(um.Format), reflect.TypeOf(want); got != wantT {
				t.Errorf("%s: Format runtime type = %v, want %v", p, got, wantT)
			}
			if got, wantT := reflect.TypeOf(um.AsFormat()), reflect.TypeOf(want); got != wantT {
				t.Errorf("%s: AsFormat() runtime type = %v, want %v", p, got, wantT)
			}
		}
	})

	// Forward compatibility, part two: OpenAPI types `format` as an anyOf of two
	// objects, so a payload that is not an object cannot be routed. It must not fail
	// the whole document -- messages arrive over SSE streams, where dropping a message
	// because one optional property drifted would be worse than losing the property.
	//
	// `Format` is declared as [OutputFormatUnion], so it can only ever hold a
	// registered variant; a payload no variant accepts leaves it nil. The field-level
	// error is absorbed by apijson's struct decoder, which reports per-field status
	// through the JSON metadata and returns nil for the struct
	// (internal/apijson/decoder.go, newStructTypeDecoder), so sibling fields survive
	// and the payload stays reachable through RawJSON.
	//
	// The two gjson scalar types map to nil. A JSON array does not: gjson reports both
	// objects and arrays as gjson.JSON, so an array passes the `TypeFilter: gjson.JSON`
	// of both variants and decodes as an empty left-most variant. That is a property of
	// the framework's union decoder shared by every union in this SDK -- an array
	// decodes into ToolPartState as an empty ToolStatePending, into Part as an empty
	// TextPart, and into FilePartSource as an empty FileSource in exactly the same way
	// -- and the drift is surfaced by IsKnown() returning false on the discriminator,
	// the same signal the unknown-discriminator case below relies on.
	t.Run("non-object format is non-fatal", func(t *testing.T) {
		cases := map[string]string{
			`"text"`: "<nil>",
			`123`:    "<nil>",
			`[]`:     "opencode.OutputFormatText",
		}
		for p, want := range cases {
			var m UserMessage
			if err := json.Unmarshal([]byte(`{`+base+`,"format":`+p+`}`), &m); err != nil {
				t.Fatalf("format=%s must not fail the document: %v", p, err)
			}
			if got := fmt.Sprintf("%T", m.Format); got != want {
				t.Errorf("format=%s runtime type = %s, want %s", p, got, want)
			}
			// Whatever happened to the field, the raw payload must stay reachable.
			if got := m.JSON.Format.Raw(); got != p {
				t.Errorf("format=%s: JSON.Format.Raw() = %q, want %q", p, got, p)
			}
			// An array must not masquerade as a *known* format.
			if f, ok := m.Format.(OutputFormatText); ok && f.Type.IsKnown() {
				t.Errorf("format=%s: decoded as a known OutputFormatText (%q)", p, f.Type)
			}
			// The whole message must still be intact.
			if m.ID != "msg_1" || m.Agent != "general" {
				t.Errorf("format=%s: sibling fields lost (%+v)", p, m)
			}
			// And the same must hold through the Message carrier.
			var msg Message
			if err := json.Unmarshal([]byte(`{`+base+`,"format":`+p+`}`), &msg); err != nil {
				t.Errorf("format=%s through Message: %v", p, err)
			}
		}
	})

	// An unrecognised discriminator value must not fail the document either. It falls
	// back to apijson's exactness heuristic (left-most variant) and the drift is
	// surfaced through IsKnown(), exactly as McpOAuthDisabled(true) is in
	// [TestConfigMcpOAuthRuntimeTypes].
	t.Run("unknown discriminator value is decoded but not known", func(t *testing.T) {
		var m UserMessage
		inner := `{"type":"future_format","x":1}`
		if err := json.Unmarshal([]byte(`{`+base+`,"format":`+inner+`}`), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		got, ok := m.Format.(OutputFormatText)
		if !ok {
			t.Fatalf("format runtime type = %T, want OutputFormatText", m.Format)
		}
		if string(got.Type) != "future_format" {
			t.Errorf("format.Type = %q, want future_format (value must survive)", got.Type)
		}
		if got.Type.IsKnown() {
			t.Error("format.Type.IsKnown() = true, want false (OpenAPI enum is [text])")
		}
		if got.JSON.RawJSON() != inner {
			t.Errorf("variant RawJSON() = %q, want %q", got.JSON.RawJSON(), inner)
		}
	})

	// A raw [MessageUnion] decode bypasses [UserMessage.UnmarshalJSON] *and*
	// [Message.UnmarshalJSON]: there is no carrier anywhere in the chain. It is the one
	// path no hand-written routing hook can reach, and it is why `Format` is declared
	// as [OutputFormatUnion] rather than `any` -- apijson installs newUnionDecoder for
	// the field itself (internal/apijson/decoder.go, newTypeDecoder consults the union
	// registry), so the value arrives typed with no UnmarshalJSON, no carrier and no
	// recovery step involved.
	//
	// Regression guard: while `Format` was `any` this path left it as map[string]any,
	// and only a lazy backfill inside [UserMessage.AsFormat] could recover it. That
	// backfill mutated the receiver, which made a bare read of `Format` in the same
	// expression as an `AsFormat()` call order-dependent. Both are now gone.
	t.Run("raw MessageUnion decode needs no recovery", func(t *testing.T) {
		var u MessageUnion
		raw := `{` + base + `,"format":{"type":"json_schema","schema":{"type":"object"},"retryCount":4}}`
		if err := apijson.UnmarshalRoot([]byte(raw), &u); err != nil {
			t.Fatalf("union unmarshal: %v", err)
		}
		um, ok := u.(UserMessage)
		if !ok {
			t.Fatalf("union runtime type = %T, want UserMessage", u)
		}
		// The field must already be typed, before AsFormat is ever called.
		got, ok := um.Format.(OutputFormatJsonSchema)
		if !ok {
			t.Fatalf("Format runtime type = %T, want OutputFormatJsonSchema", um.Format)
		}
		if got.RetryCount != 4 {
			t.Errorf("Format.RetryCount = %d, want 4", got.RetryCount)
		}
		if got.Schema["type"] != "object" {
			t.Errorf("Format.Schema = %v, want map with type=object", got.Schema)
		}
		// AsFormat is now a pure read of that field. It must agree exactly, and being a
		// value receiver it must also work on a non-addressable value such as the
		// result of a type assertion. (DeepEqual because OutputFormatJsonSchema holds a
		// map and so is not comparable with ==.)
		if asf := u.(UserMessage).AsFormat(); !reflect.DeepEqual(asf, OutputFormatUnion(got)) {
			t.Errorf("AsFormat() = %#v, want the identical field value %#v", asf, got)
		}
	})
}

// TestConfigFormatterRuntimeTypes asserts the runtime types of `Config.formatter`
// match OpenAPI `anyOf [boolean, object(additionalProperties: <formatter shape>)]`
// and JS SDK v2 `boolean | { [key: string]: {...} }`. There is no `map[string]bool`
// variant in either source.
func TestConfigFormatterRuntimeTypes(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		var c Config
		if err := json.Unmarshal([]byte(`{"formatter":true}`), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		v, ok := c.Formatter.(bool)
		if !ok {
			t.Fatalf("formatter runtime type = %T, want bool", c.Formatter)
		}
		if !v {
			t.Error("formatter = false, want true")
		}
	})

	t.Run("bool false", func(t *testing.T) {
		var c Config
		if err := json.Unmarshal([]byte(`{"formatter":false}`), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := c.Formatter.(bool); !ok {
			t.Fatalf("formatter runtime type = %T, want bool", c.Formatter)
		}
	})

	// The object arm is routed by [Config.UnmarshalJSON], so it lands on the typed
	// `map[string]ConfigFormatter` the runtime comment declares -- not on the
	// `map[string]any` the generic interface branch of the decoder would produce.
	// Crucially the entry values are objects, never bare booleans, so
	// `map[string]bool` is not reachable.
	//
	// Regression: the comment used to name a type the field could never hold. It was
	// read as "the payload is re-decodable into this", which no other runtime comment
	// in the SDK means and which left [Config.Formatter] callers with untyped maps.
	t.Run("object of overrides", func(t *testing.T) {
		raw := `{"formatter":{"gofmt":{"disabled":false,"command":["gofmt","-w"],` +
			`"environment":{"GO":"1"},"extensions":[".go"]}}}`
		var c Config
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		m, ok := c.Formatter.(map[string]ConfigFormatter)
		if !ok {
			t.Fatalf("formatter runtime type = %T, want map[string]ConfigFormatter", c.Formatter)
		}
		if _, generic := c.Formatter.(map[string]any); generic {
			t.Error("formatter decoded generically to map[string]any, object arm not routed")
		}
		entry := m["gofmt"]
		if got := entry.Command; len(got) != 2 || got[0] != "gofmt" || got[1] != "-w" {
			t.Errorf("ConfigFormatter.Command = %v, want [gofmt -w]", got)
		}
		if got := entry.Extensions; len(got) != 1 || got[0] != ".go" {
			t.Errorf("ConfigFormatter.Extensions = %v, want [.go]", got)
		}
		if got := entry.Environment["GO"]; got != "1" {
			t.Errorf("ConfigFormatter.Environment[GO] = %q, want \"1\"", got)
		}
		if entry.Disabled {
			t.Error("ConfigFormatter.Disabled = true, want false")
		}
		// Routing must never cost the caller the original entry.
		if entry.JSON.RawJSON() == "" {
			t.Error("ConfigFormatter.JSON.RawJSON() is empty")
		}
		// AsFormatter must agree with the routed field.
		if got := c.AsFormatter(); len(got) != 1 || got["gofmt"].Environment["GO"] != "1" {
			t.Errorf("AsFormatter() = %v, want the routed map", got)
		}
	})

	// An unrecognised property must reach the caller through ExtraFields rather than
	// pushing the whole field back to a generic map.
	t.Run("object of overrides with unknown property", func(t *testing.T) {
		var c Config
		if err := json.Unmarshal([]byte(`{"formatter":{"gofmt":{"command":["gofmt"],"zz":1}}}`), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		m, ok := c.Formatter.(map[string]ConfigFormatter)
		if !ok {
			t.Fatalf("formatter runtime type = %T, want map[string]ConfigFormatter", c.Formatter)
		}
		if _, present := m["gofmt"].JSON.ExtraFields["zz"]; !present {
			t.Errorf("formatter[gofmt].JSON.ExtraFields = %v, want key \"zz\"", m["gofmt"].JSON.ExtraFields)
		}
	})

	// The boolean arm must leave AsFormatter nil so callers can tell the two arms
	// apart, and a non-conformant payload must not fail the document.
	t.Run("AsFormatter is nil for the non-object arms", func(t *testing.T) {
		for _, raw := range []string{`{"formatter":true}`, `{"formatter":false}`,
			`{}`, `{"formatter":null}`, `{"formatter":"yes"}`, `{"formatter":[1]}`} {
			var c Config
			if err := json.Unmarshal([]byte(raw), &c); err != nil {
				t.Fatalf("%s must not fail the document: %v", raw, err)
			}
			if got := c.AsFormatter(); got != nil {
				t.Errorf("%s: AsFormatter() = %v, want nil", raw, got)
			}
		}
	})
}

// TestConfigMcpOAuthRuntimeTypes covers both OpenAPI `McpRemoteConfig.oauth`
// variants: `anyOf [McpOAuthConfig, boolean(enum: false)]` (JS SDK v2:
// `oauth?: McpOAuthConfig | false`).
//
// Both variants are routed through the registered [McpOAuthUnion], so the object
// variant resolves to the typed [McpOAuthConfig] and the scalar variant to
// [McpOAuthDisabled] — exactly what the `[McpOAuthConfig], [McpOAuthDisabled]`
// runtime comments on ConfigMcp.OAuth and McpRemoteConfig.OAuth document.
//
// Regression guard: before the union existed, `OAuth` was an `any` field with no
// union routing at all, so apijson decoded JSON objects generically to
// map[string]any and callers lost every typed field.
func TestConfigMcpOAuthRuntimeTypes(t *testing.T) {
	const oauthObj = `{"clientId":"cid","clientSecret":"sec","scope":"a b",` +
		`"callbackPort":8080,"redirectUri":"http://x/cb"}`

	assertOAuthObject := func(t *testing.T, v any) {
		t.Helper()
		cfg, ok := v.(McpOAuthConfig)
		if !ok {
			t.Fatalf("oauth runtime type = %T, want McpOAuthConfig", v)
		}
		if cfg.ClientID != "cid" || cfg.ClientSecret != "sec" || cfg.Scope != "a b" {
			t.Errorf("McpOAuthConfig = %+v, unexpected values", cfg)
		}
		if cfg.CallbackPort != 8080 {
			t.Errorf("CallbackPort = %d, want 8080", cfg.CallbackPort)
		}
		if cfg.RedirectURI != "http://x/cb" {
			t.Errorf("RedirectURI = %q, want http://x/cb", cfg.RedirectURI)
		}
		// The raw payload must stay reachable for forward compatibility.
		if cfg.JSON.RawJSON() == "" {
			t.Error("McpOAuthConfig.JSON.RawJSON() is empty, want the raw oauth object")
		}
	}

	t.Run("ConfigMcp with oauth object", func(t *testing.T) {
		raw := `{"type":"remote","url":"https://mcp.example.com","oauth":` + oauthObj + `}`
		var m ConfigMcp
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		remote, ok := m.AsUnion().(McpRemoteConfig)
		if !ok {
			t.Fatalf("union runtime type = %T, want McpRemoteConfig", m.AsUnion())
		}
		assertOAuthObject(t, m.OAuth)
		assertOAuthObject(t, m.AsOAuth())
		// The typed value must also be mirrored back onto the outer union variant.
		assertOAuthObject(t, remote.OAuth)
		assertOAuthObject(t, remote.AsOAuth())
	})

	t.Run("McpRemoteConfig with oauth object", func(t *testing.T) {
		raw := `{"type":"remote","url":"https://mcp.example.com","oauth":` + oauthObj + `}`
		var c McpRemoteConfig
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		assertOAuthObject(t, c.OAuth)
		assertOAuthObject(t, c.AsOAuth())
	})

	// OpenAPI declares the second variant as `boolean` with `enum: [false]`.
	assertOAuthDisabled := func(t *testing.T, v any) {
		t.Helper()
		got, ok := v.(McpOAuthDisabled)
		if !ok {
			t.Fatalf("oauth runtime type = %T, want McpOAuthDisabled", v)
		}
		if bool(got) {
			t.Error("oauth = true, want false")
		}
		if !got.IsKnown() {
			t.Error("McpOAuthDisabled(false).IsKnown() = false, want true")
		}
	}

	t.Run("ConfigMcp with oauth false", func(t *testing.T) {
		raw := `{"type":"remote","url":"https://mcp.example.com","oauth":false}`
		var m ConfigMcp
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		assertOAuthDisabled(t, m.OAuth)
		assertOAuthDisabled(t, m.AsOAuth())
		remote, ok := m.AsUnion().(McpRemoteConfig)
		if !ok {
			t.Fatalf("union runtime type = %T, want McpRemoteConfig", m.AsUnion())
		}
		assertOAuthDisabled(t, remote.OAuth)
	})

	t.Run("McpRemoteConfig with oauth false", func(t *testing.T) {
		raw := `{"type":"remote","url":"https://mcp.example.com","oauth":false}`
		var c McpRemoteConfig
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		assertOAuthDisabled(t, c.OAuth)
		assertOAuthDisabled(t, c.AsOAuth())
	})

	// `oauth: true` violates `enum: [false]`. It must not fail the whole document;
	// the constraint is surfaced through IsKnown() instead.
	t.Run("oauth true is decoded but not known", func(t *testing.T) {
		raw := `{"type":"remote","url":"https://mcp.example.com","oauth":true}`
		var c McpRemoteConfig
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		got, ok := c.OAuth.(McpOAuthDisabled)
		if !ok {
			t.Fatalf("oauth runtime type = %T, want McpOAuthDisabled", c.OAuth)
		}
		if !bool(got) {
			t.Error("oauth = false, want true")
		}
		if got.IsKnown() {
			t.Error("McpOAuthDisabled(true).IsKnown() = true, want false (enum is [false])")
		}
	})

	// oauth absent must leave the field nil (it is optional in OpenAPI).
	t.Run("oauth absent", func(t *testing.T) {
		for name, raw := range map[string]string{
			"absent": `{"type":"remote","url":"https://mcp.example.com"}`,
			"null":   `{"type":"remote","url":"https://mcp.example.com","oauth":null}`,
		} {
			t.Run(name, func(t *testing.T) {
				var c McpRemoteConfig
				if err := json.Unmarshal([]byte(raw), &c); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				if c.OAuth != nil {
					t.Errorf("%s oauth = %#v, want nil", name, c.OAuth)
				}
				if c.AsOAuth() != nil {
					t.Errorf("%s AsOAuth() = %#v, want nil", name, c.AsOAuth())
				}

				var m ConfigMcp
				if err := json.Unmarshal([]byte(raw), &m); err != nil {
					t.Fatalf("ConfigMcp unmarshal: %v", err)
				}
				if m.OAuth != nil {
					t.Errorf("%s ConfigMcp.OAuth = %#v, want nil", name, m.OAuth)
				}
			})
		}
	})

	// The outer ConfigMcpUnion routing must keep working while the nested oauth
	// union is resolved.
	t.Run("outer ConfigMcpUnion field mapping", func(t *testing.T) {
		var local ConfigMcp
		if err := json.Unmarshal([]byte(`{"type":"local","command":["srv","--port","1"],`+
			`"cwd":"/w","environment":{"A":"1"},"enabled":true,"timeout":5000}`), &local); err != nil {
			t.Fatalf("local unmarshal: %v", err)
		}
		lv, ok := local.AsUnion().(McpLocalConfig)
		if !ok {
			t.Fatalf("local union runtime type = %T, want McpLocalConfig", local.AsUnion())
		}
		if len(lv.Command) != 3 || lv.Command[0] != "srv" {
			t.Errorf("McpLocalConfig.Command = %v, want [srv --port 1]", lv.Command)
		}
		if lv.Cwd != "/w" || lv.Timeout != 5000 || !lv.Enabled {
			t.Errorf("McpLocalConfig = %+v, unexpected values", lv)
		}
		if lv.Environment["A"] != "1" {
			t.Errorf("McpLocalConfig.Environment = %v, want map[A:1]", lv.Environment)
		}
		if local.AsOAuth() != nil {
			t.Errorf("local AsOAuth() = %#v, want nil", local.AsOAuth())
		}

		var remote ConfigMcp
		if err := json.Unmarshal([]byte(`{"type":"remote","url":"https://mcp.example.com",`+
			`"headers":{"H":"v"},"enabled":true,"timeout":9000,"oauth":`+oauthObj+`}`), &remote); err != nil {
			t.Fatalf("remote unmarshal: %v", err)
		}
		rv, ok := remote.AsUnion().(McpRemoteConfig)
		if !ok {
			t.Fatalf("remote union runtime type = %T, want McpRemoteConfig", remote.AsUnion())
		}
		if rv.URL != "https://mcp.example.com" || rv.Timeout != 9000 || rv.Headers["H"] != "v" {
			t.Errorf("McpRemoteConfig = %+v, unexpected values", rv)
		}
	})
}

// TestConfigMcpUnionDiscriminator asserts every variant of the OpenAPI
// `Config.mcp.additionalProperties.anyOf` routes to its own Go type:
// McpLocalConfig (`type: "local"`), McpRemoteConfig (`type: "remote"`) and the
// disabled form `{enabled: boolean}` (no `type` at all).
//
// Regression: the union used to be registered without a discriminator, so the
// exactness heuristic picked the variant. Because the struct decoder never
// penalises missing `required` fields, `{"enabled": false}` scored as an exact
// match on the left-most variant and always resolved to [McpLocalConfig], leaving
// [ConfigMcpDisabled] unreachable. `type` is now the discriminator key.
//
// The `*-unknown` rows pin forward compatibility: an unrecognised property must
// not shift a document onto a different variant. That is what rules out simply
// reordering the variants — ties in exactness are broken left-to-right, so a
// leading [ConfigMcpDisabled] would swallow any object with extra properties.
func TestConfigMcpUnionDiscriminator(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		want any
	}{
		{"local", `{"type":"local","command":["a","b"],"cwd":"/tmp"}`, McpLocalConfig{}},
		{"local minimal", `{"type":"local","command":["a"]}`, McpLocalConfig{}},
		{"local unknown property", `{"type":"local","command":["a"],"zz":1}`, McpLocalConfig{}},
		{"remote", `{"type":"remote","url":"https://x","oauth":{"clientId":"a"}}`, McpRemoteConfig{}},
		{"remote minimal", `{"type":"remote","url":"https://x"}`, McpRemoteConfig{}},
		{"remote oauth false", `{"type":"remote","url":"https://x","oauth":false}`, McpRemoteConfig{}},
		{"remote unknown property", `{"type":"remote","url":"https://x","zz":1}`, McpRemoteConfig{}},
		{"disabled false", `{"enabled":false}`, ConfigMcpDisabled{}},
		{"disabled true", `{"enabled":true}`, ConfigMcpDisabled{}},
		{"disabled unknown property", `{"enabled":false,"zz":1}`, ConfigMcpDisabled{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var m ConfigMcp
			if err := json.Unmarshal([]byte(tc.raw), &m); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got, want := reflect.TypeOf(m.AsUnion()), reflect.TypeOf(tc.want); got != want {
				t.Fatalf("union runtime type = %v, want %v", got, want)
			}
			if m.JSON.RawJSON() != tc.raw {
				t.Errorf("RawJSON() = %q, want %q", m.JSON.RawJSON(), tc.raw)
			}
		})
	}

	// The same routing must hold through the real access path, Config.mcp, which
	// OpenAPI types as an object map.
	t.Run("through Config.mcp map", func(t *testing.T) {
		var cfg Config
		raw := `{"mcp":{"l":{"type":"local","command":["a"]},` +
			`"r":{"type":"remote","url":"https://x","oauth":{"clientId":"cid"}},` +
			`"d":{"enabled":false}}}`
		if err := json.Unmarshal([]byte(raw), &cfg); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(cfg.Mcp) != 3 {
			t.Fatalf("len(Config.Mcp) = %d, want 3", len(cfg.Mcp))
		}
		if _, ok := cfg.Mcp["l"].AsUnion().(McpLocalConfig); !ok {
			t.Errorf("mcp[l] = %T, want McpLocalConfig", cfg.Mcp["l"].AsUnion())
		}
		if _, ok := cfg.Mcp["d"].AsUnion().(ConfigMcpDisabled); !ok {
			t.Errorf("mcp[d] = %T, want ConfigMcpDisabled", cfg.Mcp["d"].AsUnion())
		}
		if cfg.Mcp["d"].Enabled {
			t.Error("mcp[d].Enabled = true, want false")
		}
		r, ok := cfg.Mcp["r"].AsUnion().(McpRemoteConfig)
		if !ok {
			t.Fatalf("mcp[r] = %T, want McpRemoteConfig", cfg.Mcp["r"].AsUnion())
		}
		if r.URL != "https://x" {
			t.Errorf("mcp[r].URL = %q, want https://x", r.URL)
		}
		oauth, ok := cfg.Mcp["r"].OAuth.(McpOAuthConfig)
		if !ok {
			t.Fatalf("mcp[r].OAuth = %T, want McpOAuthConfig", cfg.Mcp["r"].OAuth)
		}
		if oauth.ClientID != "cid" {
			t.Errorf("mcp[r].OAuth.ClientID = %q, want cid", oauth.ClientID)
		}
	})
}

// TestReferenceSourceRuntimeTypes asserts `ReferenceInfo.source` resolves to the
// OpenAPI `ReferenceSource` union (`anyOf [ReferenceLocalSource, ReferenceGitSource]`),
// which is what the runtime comment on [V2ReferenceInfo.Source] documents.
func TestReferenceSourceRuntimeTypes(t *testing.T) {
	t.Run("local", func(t *testing.T) {
		raw := `{"name":"docs","path":"./docs","source":{"type":"local","path":"./docs",` +
			`"description":"d","hidden":true}}`
		var info V2ReferenceInfo
		if err := json.Unmarshal([]byte(raw), &info); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		local, ok := info.Source.AsUnion().(ReferenceLocalSource)
		if !ok {
			t.Fatalf("source union runtime type = %T, want ReferenceLocalSource", info.Source.AsUnion())
		}
		if local.Type != ReferenceLocalSourceTypeLocal || !local.Type.IsKnown() {
			t.Errorf("local.Type = %q, want %q", local.Type, ReferenceLocalSourceTypeLocal)
		}
		if local.Path != "./docs" || local.Description != "d" || !local.Hidden {
			t.Errorf("ReferenceLocalSource = %+v, unexpected values", local)
		}
		// Flattened carrier fields must be ported from the variant.
		if info.Source.Type != "local" || info.Source.Path != "./docs" {
			t.Errorf("carrier = %+v, want Type=local Path=./docs", info.Source)
		}
		if info.Source.Repository != "" || info.Source.Branch != "" {
			t.Errorf("git-only carrier fields populated for local variant: %+v", info.Source)
		}
	})

	t.Run("git", func(t *testing.T) {
		raw := `{"name":"up","path":"./up","source":{"type":"git",` +
			`"repository":"https://github.com/x/y","branch":"main"}}`
		var info V2ReferenceInfo
		if err := json.Unmarshal([]byte(raw), &info); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		git, ok := info.Source.AsUnion().(ReferenceGitSource)
		if !ok {
			t.Fatalf("source union runtime type = %T, want ReferenceGitSource", info.Source.AsUnion())
		}
		if git.Type != ReferenceGitSourceTypeGit || !git.Type.IsKnown() {
			t.Errorf("git.Type = %q, want %q", git.Type, ReferenceGitSourceTypeGit)
		}
		if git.Repository != "https://github.com/x/y" || git.Branch != "main" {
			t.Errorf("ReferenceGitSource = %+v, unexpected values", git)
		}
		if info.Source.Type != "git" || info.Source.Repository != "https://github.com/x/y" {
			t.Errorf("carrier = %+v, want Type=git Repository=...", info.Source)
		}
		if info.Source.Path != "" {
			t.Errorf("local-only carrier field populated for git variant: %+v", info.Source)
		}
	})

	// ReferenceInfo is an alias of V2ReferenceInfo, so it must behave identically.
	t.Run("ReferenceInfo alias", func(t *testing.T) {
		var info ReferenceInfo
		raw := `{"name":"n","path":"p","source":{"type":"git","repository":"r"}}`
		if err := json.Unmarshal([]byte(raw), &info); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := info.Source.AsUnion().(ReferenceGitSource); !ok {
			t.Fatalf("source union runtime type = %T, want ReferenceGitSource", info.Source.AsUnion())
		}
		if info.Source.JSON.RawJSON() == "" {
			t.Error("ReferenceSource.RawJSON() is empty")
		}
	})
}

// TestConfigPermissionRuntimeTypes asserts `Config.permission` and
// `AgentConfig.permission` resolve to the registered [ConfigPermissionUnion]
// variants. OpenAPI `PermissionConfig` is
// `anyOf [PermissionActionConfig, object]`, and the short-string variant is
// modelled by [ConfigPermissionAction] -- NOT by a bare `string`.
func TestConfigPermissionRuntimeTypes(t *testing.T) {
	for _, action := range []string{"ask", "allow", "deny"} {
		t.Run("Config short string "+action, func(t *testing.T) {
			var c Config
			if err := json.Unmarshal([]byte(`{"permission":"`+action+`"}`), &c); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			got, ok := c.Permission.(ConfigPermissionAction)
			if !ok {
				t.Fatalf("permission runtime type = %T, want ConfigPermissionAction", c.Permission)
			}
			if string(got) != action || !got.IsKnown() {
				t.Errorf("permission = %q (known=%v), want %q", got, got.IsKnown(), action)
			}
			// The bare `string` the generic interface decoder used to produce is now
			// impossible by construction: the field is declared
			// [ConfigPermissionUnion], which `string` does not implement. Assert
			// through an `any` view so this guard still fires if the field ever
			// regresses to an untyped carrier.
			if _, isString := any(c.Permission).(string); isString {
				t.Error("permission asserted as bare string; runtime comment must not claim [string]")
			}
			if _, ok := c.AsPermission().(ConfigPermissionAction); !ok {
				t.Errorf("AsPermission() runtime type = %T, want ConfigPermissionAction", c.AsPermission())
			}
		})
	}

	t.Run("Config object", func(t *testing.T) {
		var c Config
		raw := `{"permission":{"bash":"ask","edit":"allow","webfetch":"deny"}}`
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		perm, ok := c.Permission.(ConfigPermission)
		if !ok {
			t.Fatalf("permission runtime type = %T, want ConfigPermission", c.Permission)
		}
		if perm.Webfetch != ConfigPermissionWebfetchDeny {
			t.Errorf("Webfetch = %q, want deny", perm.Webfetch)
		}
		if _, ok := c.AsPermission().(ConfigPermission); !ok {
			t.Errorf("AsPermission() runtime type = %T, want ConfigPermission", c.AsPermission())
		}
	})

	// AgentConfig routes through the identical union and already documented
	// [ConfigPermissionAction]; both must agree.
	t.Run("AgentConfig short string", func(t *testing.T) {
		var a AgentConfig
		if err := json.Unmarshal([]byte(`{"permission":"allow"}`), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := a.Permission.(ConfigPermissionAction); !ok {
			t.Fatalf("permission runtime type = %T, want ConfigPermissionAction", a.Permission)
		}
		if _, ok := a.AsPermission().(ConfigPermissionAction); !ok {
			t.Errorf("AsPermission() runtime type = %T, want ConfigPermissionAction", a.AsPermission())
		}
	})

	t.Run("AgentConfig object", func(t *testing.T) {
		var a AgentConfig
		if err := json.Unmarshal([]byte(`{"permission":{"bash":"ask"}}`), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := a.Permission.(ConfigPermission); !ok {
			t.Fatalf("permission runtime type = %T, want ConfigPermission", a.Permission)
		}
	})
}

// configPermissionRuleFields names the ten properties OpenAPI declares as
// `$ref PermissionRuleConfig` on the object arm of `PermissionConfig`, paired with
// the carrier and the typed accessor each one must expose.
func configPermissionRuleFields(p ConfigPermission) map[string]struct {
	carrier  any
	accessor ConfigPermissionBashUnion
} {
	return map[string]struct {
		carrier  any
		accessor ConfigPermissionBashUnion
	}{
		"bash":               {p.Bash, p.AsBash()},
		"edit":               {p.Edit, p.AsEdit()},
		"read":               {p.Read, p.AsRead()},
		"glob":               {p.Glob, p.AsGlob()},
		"grep":               {p.Grep, p.AsGrep()},
		"list":               {p.List, p.AsList()},
		"task":               {p.Task, p.AsTask()},
		"external_directory": {p.ExternalDirectory, p.AsExternalDirectory()},
		"lsp":                {p.Lsp, p.AsLsp()},
		"skill":              {p.Skill, p.AsSkill()},
	}
}

// TestConfigPermissionRuleRuntimeTypes asserts every one of the ten
// `$ref PermissionRuleConfig` properties of the OpenAPI `PermissionConfig` object arm
// routes through [ConfigPermissionBashUnion].
//
// OpenAPI declares `read`, `edit`, `glob`, `grep`, `list`, `bash`, `task`,
// `external_directory`, `lsp` and `skill` with the *identical* `$ref
// PermissionRuleConfig` (JS SDK v2: `PermissionRuleConfig = PermissionActionConfig |
// PermissionObjectConfig`), so all ten must behave identically.
//
// Regression, two defects at once:
//   - only `bash` carried the `[ConfigPermissionBashString], [ConfigPermissionBashMap]`
//     comment while the other nine claimed `[string], [map[string]any]`, an asymmetry
//     OpenAPI does not justify;
//   - no field routed at all, so every comment was wrong: the generic interface branch
//     of the decoder produced a bare `string` or `map[string]any` and
//     [ConfigPermissionBashUnion] was a registration nothing decoded through.
func TestConfigPermissionRuleRuntimeTypes(t *testing.T) {
	variants := []struct {
		name string
		raw  string
		want any
	}{
		{"action", `"ask"`, ConfigPermissionBashString("")},
		{"action allow", `"allow"`, ConfigPermissionBashString("")},
		{"action deny", `"deny"`, ConfigPermissionBashString("")},
		{"object", `{"git *":"allow"}`, ConfigPermissionBashMap(nil)},
		{"object empty", `{}`, ConfigPermissionBashMap(nil)},
		{"object multi", `{"git *":"allow","rm *":"deny"}`, ConfigPermissionBashMap(nil)},
	}

	// Every rule must accept every variant, through both the carrier and AsXxx().
	for _, v := range variants {
		for _, key := range []string{"bash", "edit", "read", "glob", "grep", "list",
			"task", "external_directory", "lsp", "skill"} {
			t.Run(key+"/"+v.name, func(t *testing.T) {
				var p ConfigPermission
				if err := json.Unmarshal([]byte(`{"`+key+`":`+v.raw+`}`), &p); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				field := configPermissionRuleFields(p)[key]
				want := reflect.TypeOf(v.want)
				if got := reflect.TypeOf(field.carrier); got != want {
					t.Fatalf("%s runtime type = %v, want %v", key, got, want)
				}
				if got := reflect.TypeOf(field.accessor); got != want {
					t.Errorf("%s accessor runtime type = %v, want %v", key, got, want)
				}
				// The pre-fix generic types must be gone.
				if _, bare := field.carrier.(string); bare {
					t.Errorf("%s decoded to a bare string; the union is not routed", key)
				}
				if _, generic := field.carrier.(map[string]any); generic {
					t.Errorf("%s decoded to map[string]any; the union is not routed", key)
				}
			})
		}
	}

	// The routed values must carry the actual data, and the enum contract must be
	// surfaced through IsKnown() on both arms.
	t.Run("values and IsKnown", func(t *testing.T) {
		var p ConfigPermission
		raw := `{"bash":{"git *":"allow","rm *":"deny"},"read":"ask","skill":"nope"}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		bash, ok := p.AsBash().(ConfigPermissionBashMap)
		if !ok {
			t.Fatalf("AsBash() = %T, want ConfigPermissionBashMap", p.AsBash())
		}
		if bash["git *"] != ConfigPermissionBashMapAllow || bash["rm *"] != ConfigPermissionBashMapDeny {
			t.Errorf("ConfigPermissionBashMap = %v, unexpected values", bash)
		}
		for pattern, action := range bash {
			if !action.IsKnown() {
				t.Errorf("bash[%q] = %q, IsKnown() = false", pattern, action)
			}
		}
		read, ok := p.AsRead().(ConfigPermissionBashString)
		if !ok {
			t.Fatalf("AsRead() = %T, want ConfigPermissionBashString", p.AsRead())
		}
		if read != ConfigPermissionBashStringAsk || !read.IsKnown() {
			t.Errorf("AsRead() = %q (known=%v), want ask", read, read.IsKnown())
		}
		// A value outside `enum: [ask, allow, deny]` must still decode, with the
		// drift surfaced by IsKnown() rather than failing the document.
		skill, ok := p.AsSkill().(ConfigPermissionBashString)
		if !ok {
			t.Fatalf("AsSkill() = %T, want ConfigPermissionBashString", p.AsSkill())
		}
		if string(skill) != "nope" {
			t.Errorf("AsSkill() = %q, want nope (the value must survive)", skill)
		}
		if skill.IsKnown() {
			t.Error("AsSkill().IsKnown() = true, want false")
		}
	})

	// The five `$ref PermissionActionConfig` properties are a plain string enum and
	// must stay directly typed, not routed through the rule union.
	t.Run("action-only properties stay typed", func(t *testing.T) {
		var p ConfigPermission
		raw := `{"todowrite":"ask","question":"allow","webfetch":"deny",` +
			`"websearch":"ask","doom_loop":"allow"}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if p.Todowrite != ConfigPermissionTodowriteAsk || !p.Todowrite.IsKnown() {
			t.Errorf("Todowrite = %q, want ask", p.Todowrite)
		}
		if p.Question != ConfigPermissionQuestionAllow || !p.Question.IsKnown() {
			t.Errorf("Question = %q, want allow", p.Question)
		}
		if p.Webfetch != ConfigPermissionWebfetchDeny || !p.Webfetch.IsKnown() {
			t.Errorf("Webfetch = %q, want deny", p.Webfetch)
		}
		if p.Websearch != ConfigPermissionWebsearchAsk || !p.Websearch.IsKnown() {
			t.Errorf("Websearch = %q, want ask", p.Websearch)
		}
		if p.DoomLoop != ConfigPermissionDoomLoopAllow || !p.DoomLoop.IsKnown() {
			t.Errorf("DoomLoop = %q, want allow", p.DoomLoop)
		}
	})

	// OpenAPI declares `additionalProperties: $ref PermissionRuleConfig` on this
	// object, so a rule for a tool this SDK does not know yet must reach the caller
	// through the typed extras map.
	t.Run("additionalProperties reach ExtraFields", func(t *testing.T) {
		var p ConfigPermission
		raw := `{"bash":"ask","future_tool":"deny","future_object":{"x *":"allow"}}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		action, ok := p.ExtraFields["future_tool"].(ConfigPermissionBashString)
		if !ok {
			t.Fatalf("ExtraFields[future_tool] = %T, want ConfigPermissionBashString",
				p.ExtraFields["future_tool"])
		}
		if action != ConfigPermissionBashStringDeny {
			t.Errorf("ExtraFields[future_tool] = %q, want deny", action)
		}
		object, ok := p.ExtraFields["future_object"].(ConfigPermissionBashMap)
		if !ok {
			t.Fatalf("ExtraFields[future_object] = %T, want ConfigPermissionBashMap",
				p.ExtraFields["future_object"])
		}
		if object["x *"] != ConfigPermissionBashMapAllow {
			t.Errorf("ExtraFields[future_object] = %v, want map[x *:allow]", object)
		}
		// A known property must not leak into the extras map.
		if _, leaked := p.ExtraFields["bash"]; leaked {
			t.Error("ExtraFields contains the known property \"bash\"")
		}
		// An extra value matching neither arm must not fail the document.
		var q ConfigPermission
		if err := json.Unmarshal([]byte(`{"bash":"ask","weird":123}`), &q); err != nil {
			t.Fatalf("non-conformant extra must not fail the document: %v", err)
		}
		if _, ok := q.AsBash().(ConfigPermissionBashString); !ok {
			t.Errorf("sibling rule lost: AsBash() = %T", q.AsBash())
		}
	})

	// Absent rules must leave both the carrier and the accessor nil, and an explicit
	// null must be treated as absent rather than aborting the decode.
	t.Run("absent and null", func(t *testing.T) {
		for name, raw := range map[string]string{
			"absent": `{"webfetch":"ask"}`,
			"null":   `{"webfetch":"ask","bash":null,"edit":null,"skill":null}`,
		} {
			t.Run(name, func(t *testing.T) {
				var p ConfigPermission
				if err := json.Unmarshal([]byte(raw), &p); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				for key, field := range configPermissionRuleFields(p) {
					if field.carrier != nil {
						t.Errorf("%s = %#v, want nil", key, field.carrier)
					}
					if field.accessor != nil {
						t.Errorf("As%s() = %#v, want nil", key, field.accessor)
					}
				}
				// The sibling must survive.
				if p.Webfetch != ConfigPermissionWebfetchAsk {
					t.Errorf("Webfetch = %q, want ask", p.Webfetch)
				}
			})
		}
	})

	// The routing must also hold on the real access paths, where [ConfigPermission]
	// is decoded as a variant of [ConfigPermissionUnion] and so never gets its own
	// UnmarshalJSON invoked (internal/apijson/decoder.go:145-149).
	t.Run("through Config.permission", func(t *testing.T) {
		var c Config
		raw := `{"permission":{"bash":{"git *":"allow"},"read":"ask","future":"deny"}}`
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		perm, ok := c.Permission.(ConfigPermission)
		if !ok {
			t.Fatalf("permission runtime type = %T, want ConfigPermission", c.Permission)
		}
		if _, ok := perm.AsBash().(ConfigPermissionBashMap); !ok {
			t.Errorf("AsBash() = %T, want ConfigPermissionBashMap", perm.AsBash())
		}
		if _, ok := perm.AsRead().(ConfigPermissionBashString); !ok {
			t.Errorf("AsRead() = %T, want ConfigPermissionBashString", perm.AsRead())
		}
		if _, ok := perm.ExtraFields["future"].(ConfigPermissionBashString); !ok {
			t.Errorf("ExtraFields[future] = %T, want ConfigPermissionBashString",
				perm.ExtraFields["future"])
		}
		// AsPermission must see the same routed variant as the `any` carrier.
		routed, ok := c.AsPermission().(ConfigPermission)
		if !ok {
			t.Fatalf("AsPermission() = %T, want ConfigPermission", c.AsPermission())
		}
		if _, ok := routed.AsBash().(ConfigPermissionBashMap); !ok {
			t.Errorf("AsPermission().AsBash() = %T, want ConfigPermissionBashMap", routed.AsBash())
		}
	})

	t.Run("through AgentConfig.permission", func(t *testing.T) {
		var a AgentConfig
		if err := json.Unmarshal([]byte(`{"permission":{"edit":{"*.go":"deny"}}}`), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		perm, ok := a.Permission.(ConfigPermission)
		if !ok {
			t.Fatalf("permission runtime type = %T, want ConfigPermission", a.Permission)
		}
		edit, ok := perm.AsEdit().(ConfigPermissionBashMap)
		if !ok {
			t.Fatalf("AsEdit() = %T, want ConfigPermissionBashMap", perm.AsEdit())
		}
		if edit["*.go"] != ConfigPermissionBashMapDeny {
			t.Errorf("AsEdit() = %v, want map[*.go:deny]", edit)
		}
	})

	// Nested one level deeper, through Config.agent.build.permission.
	t.Run("through Config.agent.build.permission", func(t *testing.T) {
		var c Config
		raw := `{"agent":{"build":{"permission":{"task":"allow"}}}}`
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		perm, ok := c.Agent.Build.Permission.(ConfigPermission)
		if !ok {
			t.Fatalf("permission runtime type = %T, want ConfigPermission", c.Agent.Build.Permission)
		}
		if got, ok := perm.AsTask().(ConfigPermissionBashString); !ok || got != ConfigPermissionBashStringAllow {
			t.Errorf("AsTask() = %#v, want ConfigPermissionBashString(allow)", perm.AsTask())
		}
	})

	// Both arms of the rule union must be reachable, each exactly once.
	t.Run("every variant is reachable", func(t *testing.T) {
		var p ConfigPermission
		if err := json.Unmarshal([]byte(`{"bash":"ask","edit":{"*.go":"deny"}}`), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		seen := map[reflect.Type]int{}
		for _, field := range configPermissionRuleFields(p) {
			if field.accessor != nil {
				seen[reflect.TypeOf(field.accessor)]++
			}
		}
		for _, want := range []any{ConfigPermissionBashString(""), ConfigPermissionBashMap(nil)} {
			if n := seen[reflect.TypeOf(want)]; n != 1 {
				t.Errorf("%T reached %d times, want exactly 1", want, n)
			}
		}
		if len(seen) != 2 {
			t.Errorf("distinct variants = %d, want 2", len(seen))
		}
	})
}

// TestConfigLspUnionVariantRouting asserts every variant of the OpenAPI
// `Config.lsp.additionalProperties.anyOf` routes to its own Go type:
// [ConfigLspDisabled] for `{disabled: boolean(enum: true)}` (required
// `[disabled]`) and [ConfigLspObject] for the LSP server form (required
// `[command]`). JS SDK v2 spells the same union
// `{disabled: true} | {command: Array<string>, extensions?, disabled?, env?, initialization?}`.
//
// Regression: the union was decided by apijson's exactness heuristic, which only
// penalises unknown extra properties and never penalises a missing `required`
// field. `{"command": ["gopls"], "zz": 1}` therefore scored `extras` on *both*
// variants and the left-to-right tie-break handed a genuine LSP server config to
// [ConfigLspDisabled]. [ConfigLsp.UnmarshalJSON] now routes on payload shape.
//
// The `*-unknown` rows pin forward compatibility in both directions: an
// unrecognised property must never shift a document onto the other variant.
// That is also what rules out the two generic escapes -- there is no shared
// discriminator key to register, and putting [ConfigLspObject] first would let
// `{"disabled": true}` score `exact` on it (a missing required `command` costs
// nothing) and make [ConfigLspDisabled] unreachable.
func TestConfigLspUnionVariantRouting(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		want any
	}{
		// `disabled: true` with no `command` is the only shape the disabled
		// variant can legally represent.
		{"disabled", `{"disabled":true}`, ConfigLspDisabled{}},
		{"disabled unknown property", `{"disabled":true,"zz":1}`, ConfigLspDisabled{}},
		// An explicit null counts as absent, so a null `command` does not drag the
		// document onto the object variant.
		{"disabled with null command", `{"disabled":true,"command":null}`, ConfigLspDisabled{}},

		// `disabled: false` violates the disabled variant's `enum: [true]`, so it
		// belongs to the object variant.
		{"disabled false", `{"disabled":false}`, ConfigLspObject{}},
		{"object", `{"command":["gopls"]}`, ConfigLspObject{}},
		{"object with disabled false", `{"command":["gopls"],"disabled":false}`, ConfigLspObject{}},
		// `command` is forbidden on the disabled variant
		// (`additionalProperties: false`), so `command` always wins.
		{"object with disabled true", `{"command":["gopls"],"disabled":true}`, ConfigLspObject{}},
		{"object unknown property", `{"command":["gopls"],"zz":1}`, ConfigLspObject{}},
		{"object full", `{"command":["gopls"],"extensions":[".go"],"disabled":false,` +
			`"env":{"GOFLAGS":"-tags=x"},"initialization":{"ui":{"semanticTokens":true}}}`,
			ConfigLspObject{}},
		// Neither variant matches strictly (one requires `disabled: true`, the other
		// `command`), but only the object variant can hold them without producing a
		// [ConfigLspDisabledDisabled] that fails its own IsKnown contract.
		{"empty object", `{}`, ConfigLspObject{}},
		{"partial object", `{"extensions":[".go"]}`, ConfigLspObject{}},
		{"null command without disabled", `{"command":null}`, ConfigLspObject{}},
		{"null disabled", `{"disabled":null,"command":["gopls"]}`, ConfigLspObject{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var l ConfigLsp
			if err := json.Unmarshal([]byte(tc.raw), &l); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got, want := reflect.TypeOf(l.AsUnion()), reflect.TypeOf(tc.want); got != want {
				t.Fatalf("union runtime type = %v, want %v", got, want)
			}
			// Routing must never cost the caller the original document.
			if l.JSON.RawJSON() != tc.raw {
				t.Errorf("RawJSON() = %q, want %q", l.JSON.RawJSON(), tc.raw)
			}
		})
	}
}

// TestConfigLspVariantFieldValues asserts the routed variant carries every OpenAPI
// property, and that apijson.Port mirrors them onto the [ConfigLsp] carrier.
func TestConfigLspVariantFieldValues(t *testing.T) {
	t.Run("object", func(t *testing.T) {
		raw := `{"command":["gopls","-remote=auto"],"extensions":[".go",".mod"],` +
			`"disabled":false,"env":{"GOFLAGS":"-tags=x"},` +
			`"initialization":{"ui":{"semanticTokens":true}}}`
		var l ConfigLsp
		if err := json.Unmarshal([]byte(raw), &l); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		obj, ok := l.AsUnion().(ConfigLspObject)
		if !ok {
			t.Fatalf("union = %T, want ConfigLspObject", l.AsUnion())
		}
		if got := obj.Command; len(got) != 2 || got[0] != "gopls" || got[1] != "-remote=auto" {
			t.Errorf("Command = %v, want [gopls -remote=auto]", got)
		}
		if got := obj.Extensions; len(got) != 2 || got[0] != ".go" || got[1] != ".mod" {
			t.Errorf("Extensions = %v, want [.go .mod]", got)
		}
		if obj.Disabled {
			t.Error("Disabled = true, want false")
		}
		if got := obj.Env["GOFLAGS"]; got != "-tags=x" {
			t.Errorf("Env[GOFLAGS] = %q, want -tags=x", got)
		}
		if _, ok := obj.Initialization["ui"]; !ok {
			t.Errorf("Initialization = %v, want key \"ui\"", obj.Initialization)
		}
		if obj.JSON.RawJSON() != raw {
			t.Errorf("variant RawJSON() = %q, want %q", obj.JSON.RawJSON(), raw)
		}

		// apijson.Port mirrors the typed values onto the `any` carrier fields, which
		// is what the runtime comments on [ConfigLsp] document.
		if got, ok := l.Command.([]string); !ok || len(got) != 2 {
			t.Errorf("carrier Command = %#v, want []string of len 2", l.Command)
		}
		if got, ok := l.Extensions.([]string); !ok || len(got) != 2 {
			t.Errorf("carrier Extensions = %#v, want []string of len 2", l.Extensions)
		}
		if got, ok := l.Env.(map[string]string); !ok || got["GOFLAGS"] != "-tags=x" {
			t.Errorf("carrier Env = %#v, want map[string]string", l.Env)
		}
		if _, ok := l.Initialization.(map[string]any); !ok {
			t.Errorf("carrier Initialization = %#v, want map[string]any", l.Initialization)
		}
		if l.Disabled {
			t.Error("carrier Disabled = true, want false")
		}
	})

	t.Run("disabled", func(t *testing.T) {
		raw := `{"disabled":true}`
		var l ConfigLsp
		if err := json.Unmarshal([]byte(raw), &l); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		d, ok := l.AsUnion().(ConfigLspDisabled)
		if !ok {
			t.Fatalf("union = %T, want ConfigLspDisabled", l.AsUnion())
		}
		if d.Disabled != ConfigLspDisabledDisabledTrue {
			t.Errorf("Disabled = %v, want true", d.Disabled)
		}
		if !d.Disabled.IsKnown() {
			t.Error("Disabled.IsKnown() = false; OpenAPI pins this variant to enum [true]")
		}
		if !l.Disabled {
			t.Error("carrier Disabled = false, want true")
		}
		// The disabled variant declares no other property.
		if l.Command != nil {
			t.Errorf("carrier Command = %#v, want nil", l.Command)
		}
	})

	// ConfigLspDisabled must never be selected with a value that fails its own
	// enum: [true] contract -- that was the practical damage of the old routing.
	t.Run("disabled variant is never selected with a false value", func(t *testing.T) {
		for _, raw := range []string{`{}`, `{"disabled":false}`, `{"command":["gopls"],"zz":1}`} {
			var l ConfigLsp
			if err := json.Unmarshal([]byte(raw), &l); err != nil {
				t.Fatalf("unmarshal %s: %v", raw, err)
			}
			if d, ok := l.AsUnion().(ConfigLspDisabled); ok && !d.Disabled.IsKnown() {
				t.Errorf("%s routed to ConfigLspDisabled{Disabled:%v}, which violates enum [true]",
					raw, d.Disabled)
			}
		}
	})
}

// TestConfigLspRuntimeTypes documents the decode behaviour of `Config.lsp`, whose
// OpenAPI type is `anyOf [boolean, object(additionalProperties: <the ConfigLsp
// union>)]` (JS SDK v2: `lsp?: boolean | { [key: string]: ... }`).
//
// As with [TestConfigFormatterRuntimeTypes], the object arm is routed by
// [Config.UnmarshalJSON] onto the `map[string]ConfigLsp` its runtime comment
// declares, and each entry decodes through [ConfigLsp.UnmarshalJSON] so the nested
// per-entry union is routed by payload shape as well.
func TestConfigLspRuntimeTypes(t *testing.T) {
	t.Run("bool true", func(t *testing.T) {
		var c Config
		if err := json.Unmarshal([]byte(`{"lsp":true}`), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		v, ok := c.Lsp.(bool)
		if !ok {
			t.Fatalf("lsp runtime type = %T, want bool", c.Lsp)
		}
		if !v {
			t.Error("lsp = false, want true")
		}
	})

	t.Run("bool false", func(t *testing.T) {
		var c Config
		if err := json.Unmarshal([]byte(`{"lsp":false}`), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := c.Lsp.(bool); !ok {
			t.Fatalf("lsp runtime type = %T, want bool", c.Lsp)
		}
	})

	t.Run("object of overrides", func(t *testing.T) {
		raw := `{"lsp":{"go":{"command":["gopls"],"extensions":[".go"]},"off":{"disabled":true}}}`
		var c Config
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		m, ok := c.Lsp.(map[string]ConfigLsp)
		if !ok {
			t.Fatalf("lsp runtime type = %T, want map[string]ConfigLsp", c.Lsp)
		}
		if _, generic := c.Lsp.(map[string]any); generic {
			t.Error("lsp decoded generically to map[string]any, object arm not routed")
		}
		if len(m) != 2 {
			t.Fatalf("len(lsp) = %d, want 2", len(m))
		}
		// Each entry must land on its own variant of the nested per-entry union.
		if _, ok := m["go"].AsUnion().(ConfigLspObject); !ok {
			t.Errorf("lsp[go] = %T, want ConfigLspObject", m["go"].AsUnion())
		}
		if _, ok := m["off"].AsUnion().(ConfigLspDisabled); !ok {
			t.Errorf("lsp[off] = %T, want ConfigLspDisabled", m["off"].AsUnion())
		}
		// AsLsp must agree with the routed field.
		if got := c.AsLsp(); len(got) != 2 {
			t.Errorf("AsLsp() = %v, want the routed map", got)
		}
	})

	// The boolean arm must leave AsLsp nil so callers can tell the two arms apart,
	// and a non-conformant payload must not fail the document.
	t.Run("AsLsp is nil for the non-object arms", func(t *testing.T) {
		for _, raw := range []string{`{"lsp":true}`, `{"lsp":false}`,
			`{}`, `{"lsp":null}`, `{"lsp":"yes"}`, `{"lsp":[1]}`} {
			var c Config
			if err := json.Unmarshal([]byte(raw), &c); err != nil {
				t.Fatalf("%s must not fail the document: %v", raw, err)
			}
			if got := c.AsLsp(); got != nil {
				t.Errorf("%s: AsLsp() = %v, want nil", raw, got)
			}
		}
	})

	// The same routing must hold on a standalone `map[string]ConfigLsp` decode, which
	// is the shape [Config.AsLsp] returns.
	t.Run("re-decode into map[string]ConfigLsp", func(t *testing.T) {
		raw := `{"go":{"command":["gopls"],"extensions":[".go"]},` +
			`"off":{"disabled":true},` +
			`"future":{"command":["x"],"zz":1}}`
		var typed map[string]ConfigLsp
		if err := json.Unmarshal([]byte(raw), &typed); err != nil {
			t.Fatalf("typed unmarshal: %v", err)
		}
		if len(typed) != 3 {
			t.Fatalf("len = %d, want 3", len(typed))
		}
		gopls, ok := typed["go"].AsUnion().(ConfigLspObject)
		if !ok {
			t.Fatalf("lsp[go] = %T, want ConfigLspObject", typed["go"].AsUnion())
		}
		if len(gopls.Command) != 1 || gopls.Command[0] != "gopls" {
			t.Errorf("lsp[go].Command = %v, want [gopls]", gopls.Command)
		}
		off, ok := typed["off"].AsUnion().(ConfigLspDisabled)
		if !ok {
			t.Fatalf("lsp[off] = %T, want ConfigLspDisabled", typed["off"].AsUnion())
		}
		if off.Disabled != ConfigLspDisabledDisabledTrue {
			t.Errorf("lsp[off].Disabled = %v, want true", off.Disabled)
		}
		// An entry carrying a property this SDK does not know yet must still land on
		// the object variant instead of being mistaken for the disabled form.
		future, ok := typed["future"].AsUnion().(ConfigLspObject)
		if !ok {
			t.Fatalf("lsp[future] = %T, want ConfigLspObject", typed["future"].AsUnion())
		}
		if len(future.Command) != 1 || future.Command[0] != "x" {
			t.Errorf("lsp[future].Command = %v, want [x]", future.Command)
		}
		if _, ok := future.JSON.ExtraFields["zz"]; !ok {
			t.Errorf("lsp[future].JSON.ExtraFields = %v, want key \"zz\"", future.JSON.ExtraFields)
		}
	})
}

// TestConfigV2ReferenceUnionVariantRouting asserts all three variants of the
// OpenAPI `Config.reference`/`Config.references` `additionalProperties` anyOf are
// reachable: a plain string, [ConfigV2ReferenceGit] (required `[repository]`) and
// [ConfigV2ReferenceLocal] (required `[path]`). JS SDK v2 spells the same union
// `string | ConfigV2ReferenceGit | ConfigV2ReferenceLocal`.
//
// Neither object variant declares a `type` property, so like [ConfigLspUnion] this
// union has no discriminator. Registration therefore falls back to apijson's
// exactness heuristic, which cannot decide it -- see the subtest at the bottom. The
// public access paths `Config.reference` / `Config.references` do not rely on that
// heuristic: [Config.UnmarshalJSON] routes each entry by payload shape through
// [configV2Reference], which is asserted by
// [TestConfigReferenceRuntimeTypes].
func TestConfigV2ReferenceUnionVariantRouting(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		want any
	}{
		{"string", `"https://github.com/o/r"`, ConfigV2ReferenceString("")},
		{"git", `{"repository":"o/r"}`, ConfigV2ReferenceGit{}},
		{"git full", `{"repository":"o/r","branch":"main","description":"d","hidden":true}`,
			ConfigV2ReferenceGit{}},
		{"local", `{"path":"./docs"}`, ConfigV2ReferenceLocal{}},
		{"local full", `{"path":"./docs","description":"d","hidden":false}`,
			ConfigV2ReferenceLocal{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var u ConfigV2ReferenceUnion
			if err := apijson.UnmarshalRoot([]byte(tc.raw), &u); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if got, want := reflect.TypeOf(u), reflect.TypeOf(tc.want); got != want {
				t.Fatalf("union runtime type = %v, want %v", got, want)
			}
		})
	}

	// Known limitation of the raw apijson path only: with an unrecognised property
	// present, both object variants score `extras` (the struct decoder never
	// penalises the missing `required` field) and the left-to-right tie-break picks
	// the registered-first [ConfigV2ReferenceGit], so a local reference is
	// mis-routed. Fixing the heuristic itself would mean changing internal/apijson;
	// [Config.UnmarshalJSON] instead routes by payload shape, so no public access
	// path is affected -- see [TestConfigReferenceRuntimeTypes], which pins the very
	// same payload to [ConfigV2ReferenceLocal].
	t.Run("unknown property on local is a known mis-route", func(t *testing.T) {
		var u ConfigV2ReferenceUnion
		if err := apijson.UnmarshalRoot([]byte(`{"path":"./docs","zz":1}`), &u); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if _, ok := u.(ConfigV2ReferenceLocal); ok {
			// Routing improved; accept it rather than failing the suite, and drop
			// this subtest along with the note above.
			return
		}
		if _, ok := u.(ConfigV2ReferenceGit); !ok {
			t.Errorf("runtime type = %T, want ConfigV2ReferenceGit or ConfigV2ReferenceLocal", u)
		}
	})
}

// TestConfigReferenceRuntimeTypes asserts `Config.reference` and
// `Config.references` route every entry of their OpenAPI
// `additionalProperties: anyOf [string, ConfigV2ReferenceGit, ConfigV2ReferenceLocal]`
// (JS SDK v2: `{ [key: string]: string | ConfigV2ReferenceGit | ConfigV2ReferenceLocal }`)
// onto the variant its runtime comment declares.
//
// Regression, two defects at once:
//   - the entries used to be filled by the generic interface branch of the decoder,
//     so an object arrived as `map[string]any` and callers lost every typed field --
//     the `[ConfigV2ReferenceGit], [ConfigV2ReferenceLocal]` comment was unreachable
//     and [ConfigV2ReferenceString] was never produced at all;
//   - routing through the registered union would not have been enough: the two object
//     variants share no discriminator and the exactness heuristic mis-routes
//     `{"path": "./docs", "zz": 1}` to [ConfigV2ReferenceGit] (see
//     [TestConfigV2ReferenceUnionVariantRouting]). [configV2Reference] decides on the
//     `required` keys instead, exactly as [ConfigLsp.UnmarshalJSON] does.
func TestConfigReferenceRuntimeTypes(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		want any
	}{
		{"string", `"https://github.com/o/r"`, ConfigV2ReferenceString("")},
		{"git", `{"repository":"o/r"}`, ConfigV2ReferenceGit{}},
		{"git full", `{"repository":"o/r","branch":"main","description":"d","hidden":true}`,
			ConfigV2ReferenceGit{}},
		{"local", `{"path":"./docs"}`, ConfigV2ReferenceLocal{}},
		{"local full", `{"path":"./docs","description":"d","hidden":false}`,
			ConfigV2ReferenceLocal{}},
		// Forward compatibility: an unrecognised property must never shift an entry
		// onto the other variant. These are the rows the exactness heuristic fails.
		{"git unknown property", `{"repository":"o/r","zz":1}`, ConfigV2ReferenceGit{}},
		{"local unknown property", `{"path":"./docs","zz":1}`, ConfigV2ReferenceLocal{}},
		// `repository` is forbidden on the local variant
		// (`additionalProperties: false`), so it always wins.
		{"both keys", `{"repository":"o/r","path":"./docs"}`, ConfigV2ReferenceGit{}},
		// An explicit null counts as absent, so a null `repository` does not drag the
		// entry onto the git variant.
		{"null repository", `{"repository":null,"path":"./docs"}`, ConfigV2ReferenceLocal{}},
		// Neither variant matches strictly; the local one is the only form that can
		// hold the payload without inventing a `repository`.
		{"empty object", `{}`, ConfigV2ReferenceLocal{}},
	}

	for _, field := range []string{"reference", "references"} {
		for _, tc := range cases {
			t.Run(field+"/"+tc.name, func(t *testing.T) {
				var c Config
				raw := `{"` + field + `":{"k":` + tc.raw + `}}`
				if err := json.Unmarshal([]byte(raw), &c); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				entries, unions := c.Reference, c.AsReference()
				if field == "references" {
					entries, unions = c.References, c.AsReferences()
				}
				got, want := reflect.TypeOf(entries["k"]), reflect.TypeOf(tc.want)
				if got != want {
					t.Fatalf("%s[k] runtime type = %v, want %v", field, got, want)
				}
				if _, generic := entries["k"].(map[string]any); generic {
					t.Errorf("%s[k] decoded generically to map[string]any", field)
				}
				// The typed accessor must agree with the routed field.
				if got := reflect.TypeOf(unions["k"]); got != want {
					t.Errorf("As%s()[k] runtime type = %v, want %v", field, got, want)
				}
				// Routing must never cost the caller the original document.
				if c.JSON.RawJSON() != raw {
					t.Errorf("RawJSON() = %q, want %q", c.JSON.RawJSON(), raw)
				}
			})
		}
	}

	// The routed variants must carry every OpenAPI property.
	t.Run("variant field values", func(t *testing.T) {
		raw := `{"reference":{"s":"https://github.com/o/r",` +
			`"g":{"repository":"o/r","branch":"main","description":"gd","hidden":true},` +
			`"l":{"path":"./docs","description":"ld","hidden":false}}}`
		var c Config
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if got := c.Reference["s"].(ConfigV2ReferenceString); got != "https://github.com/o/r" {
			t.Errorf("reference[s] = %q, want https://github.com/o/r", got)
		}
		git := c.Reference["g"].(ConfigV2ReferenceGit)
		if git.Repository != "o/r" || git.Branch != "main" || git.Description != "gd" || !git.Hidden {
			t.Errorf("ConfigV2ReferenceGit = %+v, unexpected values", git)
		}
		if git.JSON.RawJSON() == "" {
			t.Error("ConfigV2ReferenceGit.JSON.RawJSON() is empty")
		}
		local := c.Reference["l"].(ConfigV2ReferenceLocal)
		if local.Path != "./docs" || local.Description != "ld" || local.Hidden {
			t.Errorf("ConfigV2ReferenceLocal = %+v, unexpected values", local)
		}
		if local.JSON.RawJSON() == "" {
			t.Error("ConfigV2ReferenceLocal.JSON.RawJSON() is empty")
		}
	})

	// All three variants must be reachable in one document, and each exactly once.
	t.Run("every variant is reachable", func(t *testing.T) {
		raw := `{"reference":{"s":"u","g":{"repository":"o/r"},"l":{"path":"./d"}}}`
		var c Config
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		seen := map[reflect.Type]int{}
		for _, v := range c.AsReference() {
			seen[reflect.TypeOf(v)]++
		}
		for _, want := range []any{ConfigV2ReferenceString(""), ConfigV2ReferenceGit{}, ConfigV2ReferenceLocal{}} {
			if n := seen[reflect.TypeOf(want)]; n != 1 {
				t.Errorf("%T reached %d times, want exactly 1", want, n)
			}
		}
		if len(seen) != 3 {
			t.Errorf("distinct variants = %d, want 3", len(seen))
		}
	})

	// An absent field must stay nil so callers can tell it from an empty object, and
	// a non-conformant payload must not fail the document.
	t.Run("absent and non-conformant", func(t *testing.T) {
		for _, raw := range []string{`{}`, `{"reference":null}`, `{"reference":true}`,
			`{"reference":"x"}`, `{"reference":[1]}`} {
			var c Config
			if err := json.Unmarshal([]byte(raw), &c); err != nil {
				t.Fatalf("%s must not fail the document: %v", raw, err)
			}
			if c.Reference != nil {
				t.Errorf("%s: Reference = %#v, want nil", raw, c.Reference)
			}
			if c.AsReference() != nil {
				t.Errorf("%s: AsReference() = %#v, want nil", raw, c.AsReference())
			}
		}
		// A present but empty object must be non-nil and empty.
		var c Config
		if err := json.Unmarshal([]byte(`{"reference":{}}`), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if c.AsReference() == nil || len(c.AsReference()) != 0 {
			t.Errorf("AsReference() = %#v, want an empty non-nil map", c.AsReference())
		}
	})

	// An entry matching no variant must not cost the caller its siblings.
	t.Run("non-conformant entry is non-fatal", func(t *testing.T) {
		var c Config
		raw := `{"reference":{"bad":123,"good":{"path":"./d"}},"model":"m"}`
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if c.Model != "m" {
			t.Errorf("sibling field lost, Model = %q", c.Model)
		}
		if _, ok := c.Reference["good"].(ConfigV2ReferenceLocal); !ok {
			t.Errorf("reference[good] = %T, want ConfigV2ReferenceLocal", c.Reference["good"])
		}
		// The unroutable entry must still be present, just untyped.
		if _, present := c.Reference["bad"]; !present {
			t.Error("reference[bad] dropped; the raw value must survive")
		}
	})
}
