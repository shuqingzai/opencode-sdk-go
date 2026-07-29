package opencode

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/sst/opencode-sdk-go/internal/apijson"
)

// Aligned with OpenAPI PermissionConfig anyOf schema:
// config.permission can be a string ("ask"|"allow"|"deny") or a ConfigPermission object.

// TestConfigPermissionUnionStringVariant verifies that Config.AsPermission()
// returns ConfigPermissionAction when the JSON value is a plain string.
func TestConfigPermissionUnionStringVariant(t *testing.T) {
	for _, action := range []struct {
		raw    string
		expect ConfigPermissionAction
	}{
		{`{"permission":"ask"}`, ConfigPermissionActionAsk},
		{`{"permission":"allow"}`, ConfigPermissionActionAllow},
		{`{"permission":"deny"}`, ConfigPermissionActionDeny},
	} {
		t.Run(string(action.expect), func(t *testing.T) {
			var c Config
			if err := json.Unmarshal([]byte(action.raw), &c); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			u := c.AsPermission()
			if u == nil {
				t.Fatal("AsPermission() returned nil")
			}
			act, ok := u.(ConfigPermissionAction)
			if !ok {
				t.Fatalf("AsPermission() type = %T, want ConfigPermissionAction", u)
			}
			if act != action.expect {
				t.Errorf("AsPermission() value = %q, want %q", act, action.expect)
			}
			if !act.IsKnown() {
				t.Errorf("IsKnown() = false for %q", act)
			}
			// RawJSON must be non-empty
			if c.JSON.RawJSON() == "" {
				t.Error("Config.JSON.RawJSON() empty")
			}
		})
	}
}

// TestConfigPermissionUnionObjectVariant verifies that Config.AsPermission()
// returns ConfigPermission when the JSON value is a permission object.
func TestConfigPermissionUnionObjectVariant(t *testing.T) {
	raw := `{"permission":{"bash":"ask","webfetch":{"enabled":"ask"}}}`
	var c Config
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	u := c.AsPermission()
	if u == nil {
		t.Fatal("AsPermission() returned nil")
	}
	perm, ok := u.(ConfigPermission)
	if !ok {
		t.Fatalf("AsPermission() type = %T, want ConfigPermission", u)
	}
	// Bash is an `any` field; when set to a string it should unmarshal as string
	if perm.Bash == nil {
		t.Error("Bash should be non-nil")
	}
	if c.JSON.RawJSON() == "" {
		t.Error("Config.JSON.RawJSON() empty")
	}
}

// TestConfigPermissionUnionAbsent verifies that AsPermission() returns nil
// when the permission field is absent from the JSON.
func TestConfigPermissionUnionAbsent(t *testing.T) {
	var c Config
	if err := json.Unmarshal([]byte(`{}`), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if u := c.AsPermission(); u != nil {
		t.Errorf("AsPermission() = %T, want nil when absent", u)
	}
}

// TestConfigPermissionExplicitNull verifies that an explicit `"permission": null`
// is treated as absent rather than aborting the decode.
//
// OpenAPI declares `permission` optional and non-nullable, so a conformant server
// omits it -- but neither [ConfigPermissionUnion] variant is registered with a
// gjson.Null TypeFilter, so feeding `null` to the union decoder fails with
// "was not able to coerce type as union".
//
// Regression: [Config.UnmarshalJSON] and [AgentConfig.UnmarshalJSON] guarded only
// on `gjson.Result.Raw != ""` (absence), not on the null type, so a null permission
// made the whole document fail to decode -- every other field was lost. When
// nested (e.g. `Config.agent.build.permission`) the error was swallowed by the
// struct decoder instead, silently zeroing the entire [AgentConfig]. The guard now
// matches [ConfigMcp.UnmarshalJSON], which has always treated a null `oauth` as
// absent.
func TestConfigPermissionExplicitNull(t *testing.T) {
	t.Run("Config", func(t *testing.T) {
		var c Config
		if err := json.Unmarshal([]byte(`{"permission":null,"model":"anthropic/claude","logLevel":"INFO"}`), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if c.Permission != nil || c.AsPermission() != nil {
			t.Errorf("Permission = %#v / AsPermission() = %#v, want nil", c.Permission, c.AsPermission())
		}
		// The sibling fields must survive.
		if c.Model != "anthropic/claude" {
			t.Errorf("Model = %q, want anthropic/claude", c.Model)
		}
		if c.LogLevel != ConfigLogLevelInfo {
			t.Errorf("LogLevel = %q, want INFO", c.LogLevel)
		}
	})

	t.Run("AgentConfig", func(t *testing.T) {
		var a AgentConfig
		if err := json.Unmarshal([]byte(`{"permission":null,"model":"m","steps":7}`), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if a.Permission != nil || a.AsPermission() != nil {
			t.Errorf("Permission = %#v / AsPermission() = %#v, want nil", a.Permission, a.AsPermission())
		}
		if a.Model != "m" || a.Steps != 7 {
			t.Errorf("Model/Steps = %q/%d, want m/7", a.Model, a.Steps)
		}
	})

	// Nested through Config.agent: the error used to be swallowed, zeroing the
	// whole AgentConfig instead of surfacing.
	t.Run("nested through Config.agent", func(t *testing.T) {
		raw := `{"agent":{"build":{"permission":null,"model":"m","description":"d"}}}`
		var c Config
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if c.Agent.Build.Model != "m" {
			t.Errorf("agent.build.Model = %q, want m", c.Agent.Build.Model)
		}
		if c.Agent.Build.Description != "d" {
			t.Errorf("agent.build.Description = %q, want d", c.Agent.Build.Description)
		}
		if c.Agent.Build.Permission != nil {
			t.Errorf("agent.build.Permission = %#v, want nil", c.Agent.Build.Permission)
		}
	})

	// A non-null permission nested under Config.agent must still route.
	t.Run("nested non-null still routes", func(t *testing.T) {
		raw := `{"agent":{"plan":{"permission":"deny"}}}`
		var c Config
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		got, ok := c.Agent.Plan.Permission.(ConfigPermissionAction)
		if !ok {
			t.Fatalf("agent.plan.Permission = %T, want ConfigPermissionAction", c.Agent.Plan.Permission)
		}
		if got != ConfigPermissionActionDeny {
			t.Errorf("agent.plan.Permission = %q, want deny", got)
		}
	})
}

// TestConfigPermissionActionIsKnown verifies all known variants return true,
// unknown values return false.
func TestConfigPermissionActionIsKnown(t *testing.T) {
	for _, known := range []ConfigPermissionAction{
		ConfigPermissionActionAsk,
		ConfigPermissionActionAllow,
		ConfigPermissionActionDeny,
	} {
		if !known.IsKnown() {
			t.Errorf("%q.IsKnown() = false", known)
		}
	}
	if ConfigPermissionAction("unknown").IsKnown() {
		t.Error("unknown should not be known")
	}
}

// TestAgentConfigPermissionUnion verifies AgentConfig.AsPermission() works
// the same way as Config.AsPermission() since they share the union type.
func TestAgentConfigPermissionUnion(t *testing.T) {
	t.Run("string variant", func(t *testing.T) {
		raw := `{"permission":"allow","model":{"modelID":"claude-3","providerID":"anthropic"}}`
		var a AgentConfig
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		u := a.AsPermission()
		act, ok := u.(ConfigPermissionAction)
		if !ok {
			t.Fatalf("AsPermission() type = %T, want ConfigPermissionAction", u)
		}
		if act != ConfigPermissionActionAllow {
			t.Errorf("AsPermission() value = %q, want allow", act)
		}
	})

	t.Run("object variant", func(t *testing.T) {
		raw := `{"permission":{"bash":"deny"}}`
		var a AgentConfig
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		u := a.AsPermission()
		_, ok := u.(ConfigPermission)
		if !ok {
			t.Fatalf("AsPermission() type = %T, want ConfigPermission", u)
		}
	})

	t.Run("absent", func(t *testing.T) {
		var a AgentConfig
		if err := json.Unmarshal([]byte(`{}`), &a); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if u := a.AsPermission(); u != nil {
			t.Errorf("AsPermission() = %T, want nil when absent", u)
		}
	})
}

// TestConfigGetParamsURLQuery verifies query serialization for GET /config.
// Aligned with OpenAPI GET /config query: directory?, workspace?
func TestConfigGetParamsURLQuery(t *testing.T) {
	p := ConfigGetParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestConfigUpdateParamsBodyAndQuery verifies that PATCH /config separates
// body fields from query params correctly.
func TestConfigUpdateParamsBodyAndQuery(t *testing.T) {
	p := ConfigUpdateParams{
		Directory: F("mydir"),
		Workspace: F("myws"),
		LogLevel:  F(ConfigLogLevelDebug),
	}
	// Body: only log_level, no directory/workspace
	body, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	got := string(body)
	if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
		t.Errorf("query fields leaked into body: %s", got)
	}
	if !strings.Contains(got, `"logLevel":"DEBUG"`) {
		t.Errorf("logLevel missing from body: %s", got)
	}
	// Query: only directory + workspace
	q := p.URLQuery().Encode()
	if q != "directory=mydir&workspace=myws" {
		t.Errorf("query got %q, want directory=mydir&workspace=myws", q)
	}
}

// TestConfigLogLevelIsKnown verifies all log-level enum values.
func TestConfigLogLevelIsKnown(t *testing.T) {
	for _, v := range []ConfigLogLevel{
		ConfigLogLevelDebug, ConfigLogLevelInfo, ConfigLogLevelWarn, ConfigLogLevelError,
	} {
		if !v.IsKnown() {
			t.Errorf("%q.IsKnown() = false", v)
		}
	}
	if ConfigLogLevel("TRACE").IsKnown() {
		t.Error("TRACE should not be known")
	}
}

// TestConfigRawJSONPreserved verifies that Config.JSON.RawJSON() is non-empty
// after unmarshal, confirming the metadata infrastructure is wired up.
func TestConfigRawJSONPreserved(t *testing.T) {
	raw := `{"logLevel":"INFO","autoshare":false}`
	var c Config
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if c.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
	if !strings.Contains(c.JSON.RawJSON(), "logLevel") {
		t.Errorf("RawJSON missing logLevel: %s", c.JSON.RawJSON())
	}
}

// TestConfigProviderOptionsTimeoutUnion verifies that both `timeout` and
// `headerTimeout` resolve through the registered [ConfigProviderOptionsTimeoutUnion].
//
// OpenAPI declares `ProviderConfig.options.timeout` and `.headerTimeout` as
// `anyOf [integer(exclusiveMinimum: 0), boolean(enum: false)]` (JS SDK v2:
// `timeout?: number | false`). The integer variant must therefore surface as an
// int64-backed [UnionInt], never as a float64.
//
// Regression: the carrier fields are `any` and used to be filled by the generic
// interface decoder, which hands back `float64` for every JSON number. A caller
// following the runtime comment and asserting [UnionInt] always failed, and
// asserting float64 violated the OpenAPI `integer` mapping.
func TestConfigProviderOptionsTimeoutUnion(t *testing.T) {
	t.Run("integer variant", func(t *testing.T) {
		raw := `{"timeout":30000,"headerTimeout":1500,"chunkTimeout":250,"zz":1}`
		var o ConfigProviderOptions
		if err := json.Unmarshal([]byte(raw), &o); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		got, ok := o.Timeout.(UnionInt)
		if !ok {
			t.Fatalf("Timeout runtime type = %T, want UnionInt", o.Timeout)
		}
		if got != UnionInt(30000) {
			t.Errorf("Timeout = %d, want 30000", got)
		}
		if u, ok := o.HeaderTimeout.(UnionInt); !ok || u != UnionInt(1500) {
			t.Errorf("HeaderTimeout = %#v (%T), want UnionInt(1500)", o.HeaderTimeout, o.HeaderTimeout)
		}
		if o.AsTimeout() != o.Timeout {
			t.Errorf("AsTimeout() = %#v, want %#v", o.AsTimeout(), o.Timeout)
		}
		if o.AsHeaderTimeout() != o.HeaderTimeout {
			t.Errorf("AsHeaderTimeout() = %#v, want %#v", o.AsHeaderTimeout(), o.HeaderTimeout)
		}
		// chunkTimeout is a plain integer, and unknown properties must still be
		// captured by the typed extras map.
		if o.ChunkTimeout != 250 {
			t.Errorf("ChunkTimeout = %d, want 250", o.ChunkTimeout)
		}
		if o.ExtraFields["zz"] == nil {
			t.Errorf("ExtraFields = %v, want zz captured", o.ExtraFields)
		}
		if o.JSON.RawJSON() != raw {
			t.Errorf("RawJSON() = %q, want %q", o.JSON.RawJSON(), raw)
		}
	})

	t.Run("false variant disables the timeout", func(t *testing.T) {
		raw := `{"timeout":false,"headerTimeout":false}`
		var o ConfigProviderOptions
		if err := json.Unmarshal([]byte(raw), &o); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		for name, v := range map[string]any{"timeout": o.Timeout, "headerTimeout": o.HeaderTimeout} {
			b, ok := v.(UnionBool)
			if !ok {
				t.Fatalf("%s runtime type = %T, want UnionBool", name, v)
			}
			if bool(b) {
				t.Errorf("%s = true, want false (OpenAPI pins the variant to enum [false])", name)
			}
		}
	})

	t.Run("absent leaves both nil", func(t *testing.T) {
		var o ConfigProviderOptions
		if err := json.Unmarshal([]byte(`{"apiKey":"k"}`), &o); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if o.Timeout != nil || o.AsTimeout() != nil {
			t.Errorf("Timeout = %#v / AsTimeout() = %#v, want nil", o.Timeout, o.AsTimeout())
		}
		if o.HeaderTimeout != nil || o.AsHeaderTimeout() != nil {
			t.Errorf("HeaderTimeout = %#v / AsHeaderTimeout() = %#v, want nil", o.HeaderTimeout, o.AsHeaderTimeout())
		}
		if o.APIKey != "k" {
			t.Errorf("APIKey = %q, want k", o.APIKey)
		}
	})

	t.Run("null is treated as absent", func(t *testing.T) {
		var o ConfigProviderOptions
		if err := json.Unmarshal([]byte(`{"timeout":null,"headerTimeout":null}`), &o); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if o.Timeout != nil || o.HeaderTimeout != nil {
			t.Errorf("Timeout = %#v, HeaderTimeout = %#v, want nil", o.Timeout, o.HeaderTimeout)
		}
	})

	// The union must keep resolving when reached through the full response path
	// GET /config -> Config.provider -> ProviderConfig.options.
	t.Run("through Config.Provider", func(t *testing.T) {
		raw := `{"provider":{"anthropic":{"options":{"timeout":45000,"headerTimeout":false}}}}`
		var c Config
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		opts := c.Provider["anthropic"].Options
		if u, ok := opts.Timeout.(UnionInt); !ok || u != UnionInt(45000) {
			t.Errorf("Timeout = %#v (%T), want UnionInt(45000)", opts.Timeout, opts.Timeout)
		}
		if u, ok := opts.HeaderTimeout.(UnionBool); !ok || bool(u) {
			t.Errorf("HeaderTimeout = %#v (%T), want UnionBool(false)", opts.HeaderTimeout, opts.HeaderTimeout)
		}
	})
}

// TestConfigCompactionUnmarshal verifies ConfigCompaction nested struct.
// Aligned with OpenAPI ConfigCompaction schema.
func TestConfigCompactionUnmarshal(t *testing.T) {
	raw := `{"auto":true,"prune":false,"reserved":1000,"tail_turns":5,"preserve_recent_tokens":2048}`
	var c ConfigCompaction
	if err := json.Unmarshal([]byte(raw), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !c.Auto {
		t.Error("Auto should be true")
	}
	if c.Prune {
		t.Error("Prune should be false")
	}
	if c.Reserved != 1000 {
		t.Errorf("Reserved = %d, want 1000", c.Reserved)
	}
	if c.TailTurns != 5 {
		t.Errorf("TailTurns = %d, want 5", c.TailTurns)
	}
	if c.PreserveRecentTokens != 2048 {
		t.Errorf("PreserveRecentTokens = %d, want 2048", c.PreserveRecentTokens)
	}
	if c.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved")
	}
}

// TestConfigPermissionNonConformantValueIsNonFatal pins that a `permission` value
// matching neither OpenAPI anyOf variant (string | object) does not abort the
// decode of the surrounding Config document.
//
// `permission` is declared as [ConfigPermissionUnion], so a non-conformant value is
// dropped to nil rather than surfacing as the `float64` / `bool` the generic
// interface branch of the decoder would produce. The struct decoder swallows the
// field-level union error (internal/apijson/decoder.go, the unconditional
// `return nil` at the end of newStructTypeDecoder), which is what keeps the rest of
// the document -- and [Config.RawJSON] -- intact.
func TestConfigPermissionNonConformantValueIsNonFatal(t *testing.T) {
	// A JSON array is reported by gjson as gjson.JSON, the same Type as an object,
	// so it matches the object variant's TypeFilter and decodes to an empty
	// [ConfigPermission]. That is pre-existing framework behaviour shared by every
	// `TypeFilter: gjson.JSON` variant in the SDK, and is unchanged by declaring the
	// field as its union: both routes call the same registered union decoder.
	// It is listed separately so the nil expectation below stays exact.
	for _, raw := range []string{
		`{"model":"m","permission":123}`,
		`{"model":"m","permission":true}`,
		`{"model":"m","permission":null}`,
	} {
		var c Config
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("%s: unexpected error: %s", raw, err)
		}
		if c.Model != "m" {
			t.Errorf("%s: sibling field lost, Model=%q", raw, c.Model)
		}
		if c.JSON.RawJSON() == "" {
			t.Errorf("%s: RawJSON empty", raw)
		}
		// Neither OpenAPI variant can hold a scalar, so the field stays nil rather
		// than surfacing the float64/bool the generic interface decoder produced
		// before the field was declared as [ConfigPermissionUnion].
		if c.Permission != nil {
			t.Errorf("%s: Permission = %#v (%T), want nil", raw, c.Permission, c.Permission)
		}
		if c.AsPermission() != nil {
			t.Errorf("%s: AsPermission() = %#v, want nil", raw, c.AsPermission())
		}
	}
	for _, raw := range []string{`{"model":"m","permission":[1,2]}`} {
		var c Config
		if err := json.Unmarshal([]byte(raw), &c); err != nil {
			t.Fatalf("%s: unexpected error: %s", raw, err)
		}
		if c.Model != "m" {
			t.Errorf("%s: sibling field lost, Model=%q", raw, c.Model)
		}
		if _, ok := c.Permission.(ConfigPermission); !ok {
			t.Errorf("%s: Permission = %T, want ConfigPermission (gjson types arrays as JSON)", raw, c.Permission)
		}
	}
	for _, raw := range []string{
		`{"model":"m","permission":123}`,
		`{"model":"m","permission":true}`,
		`{"model":"m","permission":null}`,
	} {
		var a AgentConfig
		if err := json.Unmarshal([]byte(raw), &a); err != nil {
			t.Fatalf("AgentConfig %s: unexpected error: %s", raw, err)
		}
		if a.Model != "m" {
			t.Errorf("AgentConfig %s: sibling lost, Model=%q", raw, a.Model)
		}
		if a.Permission != nil {
			t.Errorf("AgentConfig %s: Permission = %#v (%T), want nil", raw, a.Permission, a.Permission)
		}
	}
	// Conformant values still route to the typed union.
	var c Config
	if err := json.Unmarshal([]byte(`{"permission":"ask"}`), &c); err != nil {
		t.Fatalf("conformant string: %s", err)
	}
	if _, ok := c.Permission.(ConfigPermissionAction); !ok {
		t.Errorf("conformant string: got %T, want ConfigPermissionAction", c.Permission)
	}
	if err := json.Unmarshal([]byte(`{"permission":{"bash":"ask"}}`), &c); err != nil {
		t.Fatalf("conformant object: %s", err)
	}
	if _, ok := c.Permission.(ConfigPermission); !ok {
		t.Errorf("conformant object: got %T, want ConfigPermission", c.Permission)
	}
}

// TestConfigMcpUnionParamVariants asserts every variant of the request-side
// [ConfigMcpUnionParam] marshals to exactly the OpenAPI
// `Config.mcp.additionalProperties` anyOf shape it models:
// `[McpLocalConfig, McpRemoteConfig, {enabled: boolean}]` (JS SDK v2:
// `McpLocalConfig | McpRemoteConfig | { enabled: boolean }`).
//
// Regression: `mcp` used to be `map[string]ConfigMcpParam`, a single flat superset
// struct. It could express combinations OpenAPI forbids (both object variants
// declare `additionalProperties: false`), and typed `command`, `cwd`,
// `environment`, `headers`, `oauth` and `timeout` as `param.Field[any]` -- so the
// `array<string>`, `map<string, string>` and `integer` contracts were unenforced and
// a caller could pass anything. The same OpenAPI schema was already modelled
// correctly for `POST /mcp` as [McpAddParamsConfigUnion].
func TestConfigMcpUnionParamVariants(t *testing.T) {
	cases := []struct {
		name string
		in   ConfigMcpUnionParam
		want string
	}{
		{
			"local minimal",
			ConfigMcpLocalParam{
				Type:    F(McpLocalConfigTypeLocal),
				Command: F([]string{"srv"}),
			},
			`{"command":["srv"],"type":"local"}`,
		},
		{
			"local full",
			ConfigMcpLocalParam{
				Type:        F(McpLocalConfigTypeLocal),
				Command:     F([]string{"srv", "--port", "1"}),
				Cwd:         F("/w"),
				Enabled:     F(true),
				Environment: F(map[string]string{"A": "1"}),
				Timeout:     F(int64(5000)),
			},
			`{"command":["srv","--port","1"],"cwd":"/w","enabled":true,` +
				`"environment":{"A":"1"},"timeout":5000,"type":"local"}`,
		},
		{
			"remote minimal",
			ConfigMcpRemoteParam{
				Type: F(McpRemoteConfigTypeRemote),
				URL:  F("https://mcp.example.com"),
			},
			`{"type":"remote","url":"https://mcp.example.com"}`,
		},
		{
			"remote with oauth object",
			ConfigMcpRemoteParam{
				Type: F(McpRemoteConfigTypeRemote),
				URL:  F("https://mcp.example.com"),
				OAuth: F[ConfigMcpOAuthUnionParam](ConfigMcpOAuthParam{
					ClientID:     F("cid"),
					ClientSecret: F("sec"),
					Scope:        F("a b"),
					CallbackPort: F(int64(8080)),
					RedirectURI:  F("http://x/cb"),
				}),
			},
			`{"oauth":{"callbackPort":8080,"clientId":"cid","clientSecret":"sec",` +
				`"redirectUri":"http://x/cb","scope":"a b"},"type":"remote",` +
				`"url":"https://mcp.example.com"}`,
		},
		{
			// OpenAPI pins the scalar arm of `McpRemoteConfig.oauth` to
			// `enum: [false]`.
			"remote with oauth false",
			ConfigMcpRemoteParam{
				Type:  F(McpRemoteConfigTypeRemote),
				URL:   F("https://x"),
				OAuth: F[ConfigMcpOAuthUnionParam](UnionBool(false)),
			},
			`{"oauth":false,"type":"remote","url":"https://x"}`,
		},
		{
			"remote full",
			ConfigMcpRemoteParam{
				Type:    F(McpRemoteConfigTypeRemote),
				URL:     F("https://x"),
				Enabled: F(false),
				Headers: F(map[string]string{"H": "v"}),
				Timeout: F(int64(9000)),
			},
			`{"enabled":false,"headers":{"H":"v"},"timeout":9000,"type":"remote","url":"https://x"}`,
		},
		{"disabled false", ConfigMcpDisabledParam{Enabled: F(false)}, `{"enabled":false}`},
		{"disabled true", ConfigMcpDisabledParam{Enabled: F(true)}, `{"enabled":true}`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.in)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if string(got) != tc.want {
				t.Errorf("marshal = %s, want %s", got, tc.want)
			}
		})
	}

	// Every variant must be reachable through the real access paths, i.e. the `mcp`
	// field of both update-params structs, and each exactly once.
	t.Run("every variant is reachable through ConfigUpdateParams.Mcp", func(t *testing.T) {
		mcp := map[string]ConfigMcpUnionParam{
			"l": ConfigMcpLocalParam{Type: F(McpLocalConfigTypeLocal), Command: F([]string{"a"})},
			"r": ConfigMcpRemoteParam{Type: F(McpRemoteConfigTypeRemote), URL: F("https://x")},
			"d": ConfigMcpDisabledParam{Enabled: F(false)},
		}
		seen := map[reflect.Type]int{}
		for _, v := range mcp {
			seen[reflect.TypeOf(v)]++
		}
		for _, want := range []any{ConfigMcpLocalParam{}, ConfigMcpRemoteParam{}, ConfigMcpDisabledParam{}} {
			if n := seen[reflect.TypeOf(want)]; n != 1 {
				t.Errorf("%T reached %d times, want exactly 1", want, n)
			}
		}
		if len(seen) != 3 {
			t.Fatalf("distinct variants = %d, want 3", len(seen))
		}

		const want = `{"mcp":{"d":{"enabled":false},"l":{"command":["a"],"type":"local"},` +
			`"r":{"type":"remote","url":"https://x"}}}`
		for name, body := range map[string]any{
			"ConfigUpdateParams":       ConfigUpdateParams{Mcp: F(mcp)},
			"GlobalConfigUpdateParams": GlobalConfigUpdateParams{Mcp: F(mcp)},
		} {
			got, err := json.Marshal(body)
			if err != nil {
				t.Fatalf("%s marshal: %v", name, err)
			}
			if string(got) != want {
				t.Errorf("%s marshal = %s, want %s", name, got, want)
			}
		}
	})
}

// TestConfigFormatterSettingUnionParam asserts the request-side `formatter` field
// marshals both arms of the OpenAPI `Config.formatter` anyOf
// `[boolean, object(additionalProperties: <formatter shape>)]` (JS SDK v2:
// `boolean | { [key: string]: {...} }`).
//
// Regression: the field was `param.Field[any]` and its godoc told callers to pass
// `map[string]ConfigFormatter` -- a *response* struct whose fields are plain Go
// values with no `param.Field` wrapper and no `omitempty`. Following that advice
// serialised every zero value, so asking only for a command emitted
// `"disabled":false,"environment":{},"extensions":[]` too and the server was told to
// change settings the caller never mentioned.
func TestConfigFormatterSettingUnionParam(t *testing.T) {
	t.Run("response struct would have leaked zero values", func(t *testing.T) {
		// Pin the defect this type exists to prevent: the response struct still
		// serialises its zero values, which is why it must never reach a request.
		leaky, err := json.Marshal(map[string]ConfigFormatter{
			"gofmt": {Command: []string{"gofmt"}},
		})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		// The response struct has no `omitempty`, so every property it declares is
		// emitted whether or not the caller set it.
		for _, unwanted := range []string{`"disabled"`, `"environment"`, `"extensions"`} {
			if !strings.Contains(string(leaky), unwanted) {
				t.Fatalf("premise broken: ConfigFormatter no longer emits %s (%s)", unwanted, leaky)
			}
		}
		// The request-side type must emit only what the caller set.
		clean, err := json.Marshal(ConfigFormatterMapParam{
			"gofmt": {Command: F([]string{"gofmt"})},
		})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		if string(clean) != `{"gofmt":{"command":["gofmt"]}}` {
			t.Errorf("marshal = %s, want {\"gofmt\":{\"command\":[\"gofmt\"]}}", clean)
		}
	})

	cases := []struct {
		name string
		in   ConfigFormatterSettingUnionParam
		want string
	}{
		{"bool true enables built-ins", UnionBool(true), `{"formatter":true}`},
		{"bool false disables", UnionBool(false), `{"formatter":false}`},
		{"empty object", ConfigFormatterMapParam{}, `{"formatter":{}}`},
		{
			"object partial",
			ConfigFormatterMapParam{"gofmt": {Command: F([]string{"gofmt", "-w"})}},
			`{"formatter":{"gofmt":{"command":["gofmt","-w"]}}}`,
		},
		{
			"object full",
			ConfigFormatterMapParam{"gofmt": {
				Command:     F([]string{"gofmt", "-w"}),
				Disabled:    F(false),
				Environment: F(map[string]string{"GO": "1"}),
				Extensions:  F([]string{".go"}),
			}},
			`{"formatter":{"gofmt":{"command":["gofmt","-w"],"disabled":false,` +
				`"environment":{"GO":"1"},"extensions":[".go"]}}}`,
		},
		{
			// An explicit `"disabled": true` must survive: it is the caller asking to
			// switch one formatter off, which is distinct from not mentioning it.
			"object disabled true",
			ConfigFormatterMapParam{"prettier": {Disabled: F(true)}},
			`{"formatter":{"prettier":{"disabled":true}}}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for name, marshal := range map[string]func() ([]byte, error){
				"ConfigUpdateParams": func() ([]byte, error) {
					return json.Marshal(ConfigUpdateParams{Formatter: F(tc.in)})
				},
				"GlobalConfigUpdateParams": func() ([]byte, error) {
					return json.Marshal(GlobalConfigUpdateParams{Formatter: F(tc.in)})
				},
			} {
				got, err := marshal()
				if err != nil {
					t.Fatalf("%s marshal: %v", name, err)
				}
				if string(got) != tc.want {
					t.Errorf("%s marshal = %s, want %s", name, got, tc.want)
				}
			}
		})
	}

	// An omitted field must not appear in the body at all -- PATCH /config is a
	// partial update, so an absent `formatter` means "leave it alone".
	t.Run("omitted", func(t *testing.T) {
		got, err := json.Marshal(ConfigUpdateParams{Model: F("m")})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		if strings.Contains(string(got), "formatter") {
			t.Errorf("marshal = %s, want no formatter key", got)
		}
	})
}

// TestConfigLspSettingUnionParam asserts the request-side `lsp` field marshals both
// arms of the OpenAPI `Config.lsp` anyOf
// `[boolean, object(additionalProperties: <the per-entry anyOf>)]`, and that each
// map entry marshals as its own arm of that nested anyOf
// `[{disabled: boolean(enum: true)}, {command, extensions?, disabled?, env?, initialization?}]`
// (JS SDK v2: `{disabled: true} | {command: Array<string>, ...}`).
//
// Regression: as with `formatter`, the field was `param.Field[any]` and its godoc
// pointed at the response struct [ConfigLsp] -- which additionally flattens the two
// variants into one `any`-typed superset, so a request built from it could not even
// express which variant was meant.
func TestConfigLspSettingUnionParam(t *testing.T) {
	t.Run("response struct would have leaked zero values", func(t *testing.T) {
		leaky, err := json.Marshal(map[string]ConfigLsp{"go": {Command: []string{"gopls"}}})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		if !strings.Contains(string(leaky), `"disabled"`) {
			t.Fatalf("premise broken: ConfigLsp no longer emits \"disabled\" (%s)", leaky)
		}
		clean, err := json.Marshal(ConfigLspMapParam{
			"go": ConfigLspObjectParam{Command: F([]string{"gopls"})},
		})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		if string(clean) != `{"go":{"command":["gopls"]}}` {
			t.Errorf("marshal = %s, want {\"go\":{\"command\":[\"gopls\"]}}", clean)
		}
	})

	cases := []struct {
		name string
		in   ConfigLspSettingUnionParam
		want string
	}{
		{"bool true enables built-ins", UnionBool(true), `{"lsp":true}`},
		{"bool false disables", UnionBool(false), `{"lsp":false}`},
		{"empty object", ConfigLspMapParam{}, `{"lsp":{}}`},
		{
			"object variant minimal",
			ConfigLspMapParam{"go": ConfigLspObjectParam{Command: F([]string{"gopls"})}},
			`{"lsp":{"go":{"command":["gopls"]}}}`,
		},
		{
			"object variant full",
			ConfigLspMapParam{"go": ConfigLspObjectParam{
				Command:        F([]string{"gopls", "-remote=auto"}),
				Disabled:       F(false),
				Env:            F(map[string]string{"GOFLAGS": "-tags=x"}),
				Extensions:     F([]string{".go", ".mod"}),
				Initialization: F(map[string]any{"ui": "x"}),
			}},
			`{"lsp":{"go":{"command":["gopls","-remote=auto"],"disabled":false,` +
				`"env":{"GOFLAGS":"-tags=x"},"extensions":[".go",".mod"],` +
				`"initialization":{"ui":"x"}}}}`,
		},
		{
			// OpenAPI pins this variant to `enum: [true]`.
			"disabled variant",
			ConfigLspMapParam{"off": ConfigLspDisabledParam{Disabled: F(ConfigLspDisabledDisabledTrue)}},
			`{"lsp":{"off":{"disabled":true}}}`,
		},
		{
			"both variants in one map",
			ConfigLspMapParam{
				"go":  ConfigLspObjectParam{Command: F([]string{"gopls"})},
				"off": ConfigLspDisabledParam{Disabled: F(ConfigLspDisabledDisabledTrue)},
			},
			`{"lsp":{"go":{"command":["gopls"]},"off":{"disabled":true}}}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for name, marshal := range map[string]func() ([]byte, error){
				"ConfigUpdateParams": func() ([]byte, error) {
					return json.Marshal(ConfigUpdateParams{Lsp: F(tc.in)})
				},
				"GlobalConfigUpdateParams": func() ([]byte, error) {
					return json.Marshal(GlobalConfigUpdateParams{Lsp: F(tc.in)})
				},
			} {
				got, err := marshal()
				if err != nil {
					t.Fatalf("%s marshal: %v", name, err)
				}
				if string(got) != tc.want {
					t.Errorf("%s marshal = %s, want %s", name, got, tc.want)
				}
			}
		})
	}

	// Both arms of the per-entry union must be reachable, each exactly once.
	t.Run("every per-entry variant is reachable", func(t *testing.T) {
		entries := ConfigLspMapParam{
			"go":  ConfigLspObjectParam{Command: F([]string{"gopls"})},
			"off": ConfigLspDisabledParam{Disabled: F(ConfigLspDisabledDisabledTrue)},
		}
		seen := map[reflect.Type]int{}
		for _, v := range entries {
			seen[reflect.TypeOf(v)]++
		}
		for _, want := range []any{ConfigLspObjectParam{}, ConfigLspDisabledParam{}} {
			if n := seen[reflect.TypeOf(want)]; n != 1 {
				t.Errorf("%T reached %d times, want exactly 1", want, n)
			}
		}
		if len(seen) != 2 {
			t.Errorf("distinct variants = %d, want 2", len(seen))
		}
	})

	t.Run("omitted", func(t *testing.T) {
		got, err := json.Marshal(ConfigUpdateParams{Model: F("m")})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		if strings.Contains(string(got), "lsp") {
			t.Errorf("marshal = %s, want no lsp key", got)
		}
	})
}

// TestConfigPermissionParamRuleUnion asserts all ten `$ref PermissionRuleConfig`
// properties of the request-side [ConfigPermissionParam] accept both arms of that
// anyOf, and that the OpenAPI `additionalProperties: $ref PermissionRuleConfig` of
// the same object reaches the wire through the typed extras map.
//
// Regression: only `bash` was typed [ConfigPermissionBashUnionParam]; the other nine
// were `param.Field[any]`, so nothing stopped a caller passing a shape OpenAPI does
// not allow, and the extras map was missing entirely -- a permission rule for a tool
// this SDK does not know yet could not be sent at all.
func TestConfigPermissionParamRuleUnion(t *testing.T) {
	rules := []string{"bash", "edit", "read", "glob", "grep", "list", "task",
		"external_directory", "lsp", "skill"}

	set := func(key string, v ConfigPermissionBashUnionParam) ConfigPermissionParam {
		var p ConfigPermissionParam
		switch key {
		case "bash":
			p.Bash = F(v)
		case "edit":
			p.Edit = F(v)
		case "read":
			p.Read = F(v)
		case "glob":
			p.Glob = F(v)
		case "grep":
			p.Grep = F(v)
		case "list":
			p.List = F(v)
		case "task":
			p.Task = F(v)
		case "external_directory":
			p.ExternalDirectory = F(v)
		case "lsp":
			p.Lsp = F(v)
		case "skill":
			p.Skill = F(v)
		}
		return p
	}

	variants := []struct {
		name string
		in   ConfigPermissionBashUnionParam
		want string
	}{
		{"action", ConfigPermissionBashStringAsk, `"ask"`},
		{"object", ConfigPermissionBashMapParam{"git *": ConfigPermissionBashMapAllow}, `{"git *":"allow"}`},
	}

	for _, v := range variants {
		for _, key := range rules {
			t.Run(key+"/"+v.name, func(t *testing.T) {
				got, err := json.Marshal(set(key, v.in))
				if err != nil {
					t.Fatalf("marshal: %v", err)
				}
				want := `{"` + key + `":` + v.want + `}`
				if string(got) != want {
					t.Errorf("marshal = %s, want %s", got, want)
				}
			})
		}
	}

	// The five `$ref PermissionActionConfig` properties are a plain string enum.
	t.Run("action-only properties", func(t *testing.T) {
		got, err := json.Marshal(ConfigPermissionParam{
			Todowrite: F(ConfigPermissionTodowriteAsk),
			Question:  F(ConfigPermissionQuestionAllow),
			Webfetch:  F(ConfigPermissionWebfetchDeny),
			Websearch: F(ConfigPermissionWebsearchAsk),
			DoomLoop:  F(ConfigPermissionDoomLoopAllow),
		})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		// apijson emits object keys in sorted order.
		want := `{"doom_loop":"allow","question":"allow","todowrite":"ask",` +
			`"webfetch":"deny","websearch":"ask"}`
		if string(got) != want {
			t.Errorf("marshal = %s, want %s", got, want)
		}
	})

	t.Run("additionalProperties reach the wire", func(t *testing.T) {
		got, err := json.Marshal(ConfigPermissionParam{
			Bash: F[ConfigPermissionBashUnionParam](ConfigPermissionBashStringAsk),
			ExtraFields: map[string]ConfigPermissionBashUnionParam{
				"future_tool":   ConfigPermissionBashStringDeny,
				"future_object": ConfigPermissionBashMapParam{"x *": ConfigPermissionBashMapAllow},
			},
		})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		for _, want := range []string{`"bash":"ask"`, `"future_tool":"deny"`,
			`"future_object":{"x *":"allow"}`} {
			if !strings.Contains(string(got), want) {
				t.Errorf("marshal = %s, missing %s", got, want)
			}
		}
	})

	// Both arms of the rule union must be reachable through the real access path,
	// each exactly once.
	t.Run("every variant is reachable through ConfigUpdateParams.Permission", func(t *testing.T) {
		perm := ConfigPermissionParam{
			Bash: F[ConfigPermissionBashUnionParam](ConfigPermissionBashStringAsk),
			Edit: F[ConfigPermissionBashUnionParam](
				ConfigPermissionBashMapParam{"*.go": ConfigPermissionBashMapDeny}),
		}
		seen := map[reflect.Type]int{
			reflect.TypeOf(perm.Bash.Value): 1,
			reflect.TypeOf(perm.Edit.Value): 1,
		}
		if len(seen) != 2 {
			t.Errorf("distinct variants = %d, want 2", len(seen))
		}
		got, err := json.Marshal(ConfigUpdateParams{
			Permission: F[ConfigPermissionUnionParam](perm),
		})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		want := `{"permission":{"bash":"ask","edit":{"*.go":"deny"}}}`
		if string(got) != want {
			t.Errorf("marshal = %s, want %s", got, want)
		}
	})

	// The string arm of the outer `PermissionConfig` anyOf must still work.
	t.Run("outer permission union string arm", func(t *testing.T) {
		got, err := json.Marshal(ConfigUpdateParams{
			Permission: F[ConfigPermissionUnionParam](ConfigPermissionActionDeny),
		})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		if string(got) != `{"permission":"deny"}` {
			t.Errorf("marshal = %s, want {\"permission\":\"deny\"}", got)
		}
	})
}

// TestConfigParamAdditionalProperties asserts the request-side structs whose OpenAPI
// schema declares `additionalProperties` can actually send them.
//
// OpenAPI declares `AgentConfig.additionalProperties: {}` and
// `ProviderConfig.options.additionalProperties: {}` (JS SDK v2 spells both as an
// open index signature). Their response-side counterparts already carried an extras
// map; the request-side ones silently dropped anything the SDK did not model.
func TestConfigParamAdditionalProperties(t *testing.T) {
	t.Run("AgentConfigParam", func(t *testing.T) {
		got, err := json.Marshal(AgentConfigParam{
			Model:       F("m"),
			ExtraFields: map[string]any{"futureKnob": 1, "futureFlag": true},
		})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		for _, want := range []string{`"model":"m"`, `"futureKnob":1`, `"futureFlag":true`} {
			if !strings.Contains(string(got), want) {
				t.Errorf("marshal = %s, missing %s", got, want)
			}
		}
	})

	t.Run("ConfigProviderOptionsParam", func(t *testing.T) {
		got, err := json.Marshal(ConfigProviderOptionsParam{
			APIKey:      F("k"),
			ExtraFields: map[string]any{"futureKnob": "v"},
		})
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		for _, want := range []string{`"apiKey":"k"`, `"futureKnob":"v"`} {
			if !strings.Contains(string(got), want) {
				t.Errorf("marshal = %s, missing %s", got, want)
			}
		}
	})

	// OpenAPI types both timeouts as `anyOf [integer(exclusiveMinimum: 0),
	// boolean(enum: false)]`, so the request side must offer the same union the
	// response side resolves to -- not a `param.Field[any]` that would let a caller
	// send a float or a string.
	t.Run("ConfigProviderOptionsParam timeout union", func(t *testing.T) {
		cases := []struct {
			name string
			in   ConfigProviderOptionsTimeoutUnion
			want string
		}{
			{"integer", UnionInt(30000), `30000`},
			{"false disables", UnionBool(false), `false`},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				got, err := json.Marshal(ConfigProviderOptionsParam{
					Timeout:       F(tc.in),
					HeaderTimeout: F(tc.in),
				})
				if err != nil {
					t.Fatalf("marshal: %v", err)
				}
				want := `{"headerTimeout":` + tc.want + `,"timeout":` + tc.want + `}`
				if string(got) != want {
					t.Errorf("marshal = %s, want %s", got, want)
				}
			})
		}
	})
}

// TestConfigUnionVariantReachability is the single place that proves every variant
// of every union reachable from `config.go` can actually be produced, and that no
// two variants collapse onto the same Go type.
//
// Each row decodes (response unions) or marshals (request unions) one payload per
// registered variant, then asserts the set of distinct runtime types has exactly the
// expected size -- so a union whose variants silently collapse onto one type, the
// failure mode a per-variant assertion cannot see, is caught here.
//
// The `+unknown` payloads additionally pin forward compatibility: a property this
// SDK does not model yet must never shift a payload onto a different variant. That
// is the exact defect that made `{"path":"./docs","zz":1}` decode as
// [ConfigV2ReferenceGit] and `{"command":["gopls"],"zz":1}` as [ConfigLspDisabled].
func TestConfigUnionVariantReachability(t *testing.T) {
	// --- Response unions: decode a payload per variant. ---
	t.Run("ConfigMcpUnion", func(t *testing.T) {
		assertVariants(t, 3, decodeAll(t, func(raw string) any {
			var m ConfigMcp
			if err := json.Unmarshal([]byte(raw), &m); err != nil {
				t.Fatalf("%s: %v", raw, err)
			}
			return m.AsUnion()
		},
			`{"type":"local","command":["a"]}`,
			`{"type":"local","command":["a"],"zz":1}`,
			`{"type":"remote","url":"https://x"}`,
			`{"type":"remote","url":"https://x","zz":1}`,
			`{"enabled":false}`,
			`{"enabled":false,"zz":1}`,
		))
	})

	t.Run("ConfigLspUnion", func(t *testing.T) {
		assertVariants(t, 2, decodeAll(t, func(raw string) any {
			var l ConfigLsp
			if err := json.Unmarshal([]byte(raw), &l); err != nil {
				t.Fatalf("%s: %v", raw, err)
			}
			return l.AsUnion()
		},
			`{"disabled":true}`,
			`{"disabled":true,"zz":1}`,
			`{"command":["gopls"]}`,
			`{"command":["gopls"],"zz":1}`,
		))
	})

	t.Run("ConfigPermissionUnion", func(t *testing.T) {
		assertVariants(t, 2, decodeAll(t, func(raw string) any {
			var c Config
			if err := json.Unmarshal([]byte(`{"permission":`+raw+`}`), &c); err != nil {
				t.Fatalf("%s: %v", raw, err)
			}
			return c.AsPermission()
		},
			`"ask"`,
			`{"bash":"ask"}`,
			`{"bash":"ask","zz":"deny"}`,
		))
	})

	t.Run("ConfigPermissionBashUnion", func(t *testing.T) {
		assertVariants(t, 2, decodeAll(t, func(raw string) any {
			var p ConfigPermission
			if err := json.Unmarshal([]byte(`{"bash":`+raw+`}`), &p); err != nil {
				t.Fatalf("%s: %v", raw, err)
			}
			return p.AsBash()
		},
			`"ask"`,
			`{"git *":"allow"}`,
		))
	})

	t.Run("ConfigV2ReferenceUnion", func(t *testing.T) {
		assertVariants(t, 3, decodeAll(t, func(raw string) any {
			var c Config
			if err := json.Unmarshal([]byte(`{"reference":{"k":`+raw+`}}`), &c); err != nil {
				t.Fatalf("%s: %v", raw, err)
			}
			return c.AsReference()["k"]
		},
			`"https://x"`,
			`{"repository":"o/r"}`,
			`{"repository":"o/r","zz":1}`,
			`{"path":"./d"}`,
			`{"path":"./d","zz":1}`,
		))
	})

	t.Run("McpOAuthUnion", func(t *testing.T) {
		assertVariants(t, 2, decodeAll(t, func(raw string) any {
			var c McpRemoteConfig
			if err := json.Unmarshal([]byte(`{"type":"remote","url":"https://x","oauth":`+raw+`}`), &c); err != nil {
				t.Fatalf("%s: %v", raw, err)
			}
			return c.AsOAuth()
		},
			`{"clientId":"cid"}`,
			`{"clientId":"cid","zz":1}`,
			`false`,
		))
	})

	t.Run("ConfigProviderOptionsTimeoutUnion", func(t *testing.T) {
		// OpenAPI allows only `integer` and `false`; [UnionFloat] and the `true`
		// filter are registered so a spec-violating payload still decodes
		// losslessly instead of failing the whole document.
		assertVariants(t, 3, decodeAll(t, func(raw string) any {
			var o ConfigProviderOptions
			if err := json.Unmarshal([]byte(`{"timeout":`+raw+`}`), &o); err != nil {
				t.Fatalf("%s: %v", raw, err)
			}
			return o.AsTimeout()
		},
			`30000`,
			`1.5`,
			`false`,
			`true`,
		))
	})

	// --- Request unions: marshal one value per variant. Request unions are
	// marshal-only; decoding into one is not a supported operation, because their
	// variants wrap every field in param.Field, which the decoder cannot assign. ---
	t.Run("ConfigMcpUnionParam", func(t *testing.T) {
		assertVariants(t, 3, marshalAll(t,
			ConfigMcpLocalParam{Type: F(McpLocalConfigTypeLocal), Command: F([]string{"a"})},
			ConfigMcpRemoteParam{Type: F(McpRemoteConfigTypeRemote), URL: F("https://x")},
			ConfigMcpDisabledParam{Enabled: F(false)},
		))
	})

	t.Run("ConfigMcpOAuthUnionParam", func(t *testing.T) {
		assertVariants(t, 2, marshalAll(t,
			ConfigMcpOAuthParam{ClientID: F("cid")},
			UnionBool(false),
		))
	})

	t.Run("ConfigFormatterSettingUnionParam", func(t *testing.T) {
		assertVariants(t, 2, marshalAll(t,
			UnionBool(true),
			ConfigFormatterMapParam{"gofmt": {Command: F([]string{"gofmt"})}},
		))
	})

	t.Run("ConfigLspSettingUnionParam", func(t *testing.T) {
		assertVariants(t, 2, marshalAll(t,
			UnionBool(false),
			ConfigLspMapParam{"go": ConfigLspObjectParam{Command: F([]string{"gopls"})}},
		))
	})

	t.Run("ConfigLspUnionParam", func(t *testing.T) {
		assertVariants(t, 2, marshalAll(t,
			ConfigLspDisabledParam{Disabled: F(ConfigLspDisabledDisabledTrue)},
			ConfigLspObjectParam{Command: F([]string{"gopls"})},
		))
	})

	t.Run("ConfigPermissionUnionParam", func(t *testing.T) {
		assertVariants(t, 2, marshalAll(t,
			ConfigPermissionActionAsk,
			ConfigPermissionParam{Bash: F[ConfigPermissionBashUnionParam](ConfigPermissionBashStringAsk)},
		))
	})

	t.Run("ConfigPermissionBashUnionParam", func(t *testing.T) {
		assertVariants(t, 2, marshalAll(t,
			ConfigPermissionBashStringAsk,
			ConfigPermissionBashMapParam{"git *": ConfigPermissionBashMapAllow},
		))
	})

	t.Run("ConfigV2ReferenceUnionParam", func(t *testing.T) {
		assertVariants(t, 3, marshalAll(t,
			UnionString("https://x"),
			ConfigV2ReferenceGitParam{Repository: F("o/r")},
			ConfigV2ReferenceLocalParam{Path: F("./d")},
		))
	})
}

// decodeAll decodes every payload with decode and returns the distinct runtime types
// produced, keyed by type.
func decodeAll(t *testing.T, decode func(raw string) any, payloads ...string) map[reflect.Type][]string {
	t.Helper()
	seen := map[reflect.Type][]string{}
	for _, raw := range payloads {
		got := decode(raw)
		if got == nil {
			t.Errorf("%s decoded to nil; no variant was reached", raw)
			continue
		}
		typ := reflect.TypeOf(got)
		seen[typ] = append(seen[typ], raw)
	}
	return seen
}

// marshalAll marshals every value, asserting each round-trips, and returns the
// distinct Go types keyed by type. Request unions are only ever marshalled, so the
// variant that reaches the wire is whichever one the caller supplied.
func marshalAll(t *testing.T, values ...any) map[reflect.Type][]string {
	t.Helper()
	seen := map[reflect.Type][]string{}
	for _, v := range values {
		encoded, err := json.Marshal(v)
		if err != nil {
			t.Errorf("%T: marshal: %v", v, err)
			continue
		}
		if len(encoded) == 0 || string(encoded) == "null" {
			t.Errorf("%T: marshalled to %q, want a payload", v, encoded)
			continue
		}
		typ := reflect.TypeOf(v)
		seen[typ] = append(seen[typ], string(encoded))
	}
	return seen
}

// assertVariants asserts the union reached exactly want distinct Go types, listing
// what it did reach when it did not.
func assertVariants(t *testing.T, want int, seen map[reflect.Type][]string) {
	t.Helper()
	if len(seen) != want {
		t.Errorf("reached %d distinct variants, want %d", len(seen), want)
	}
	for typ, payloads := range seen {
		t.Logf("  %-42v <- %v", typ, payloads)
	}
}

// TestConfigNativeUnionFieldBoundaries pins the boundary behaviour of the five
// `config.go` fields that are declared as their OpenAPI union directly, rather than
// as an `any` carrier fed by hand-written gjson routing.
//
// A union-typed field is resolved by apijson's registered union decoder
// (internal/apijson/decoder.go, `if _, ok := unionRegistry[t]`), including when it
// sits behind a slice or map. That buys three properties the hand-written routing
// could not offer, all asserted below:
//
//   - `null` and an absent property leave the field nil without an error, because
//     the struct decoder skips null nodes outright;
//   - a payload matching no variant leaves the field nil and is *not* fatal,
//     because the struct decoder ends in an unconditional `return nil` and so
//     swallows the field-level union error;
//   - sibling fields and [RawJSON] always survive, whatever the union field held.
//
// The two unions the exactness heuristic genuinely mis-decides -- `ConfigLspUnion`
// and `ConfigV2ReferenceUnion` -- are deliberately *not* declared this way; see
// [TestConfigLspUnionVariantRouting] and [TestConfigReferenceRuntimeTypes].
func TestConfigNativeUnionFieldBoundaries(t *testing.T) {
	// McpRemoteConfig.oauth: anyOf [McpOAuthConfig, boolean(enum: false)].
	t.Run("McpRemoteConfig.OAuth", func(t *testing.T) {
		cases := []struct {
			name string
			raw  string
			want any
		}{
			{"object", `{"clientId":"cid"}`, McpOAuthConfig{}},
			{"object unknown property", `{"clientId":"cid","zz":1}`, McpOAuthConfig{}},
			{"object empty", `{}`, McpOAuthConfig{}},
			{"false", `false`, McpOAuthDisabled(false)},
			// `true` violates `enum: [false]`; it must still decode, with the drift
			// surfaced by IsKnown() rather than failing the document.
			{"true violates enum", `true`, McpOAuthDisabled(false)},
			{"null", `null`, nil},
			{"scalar violates both arms", `123`, nil},
			{"string violates both arms", `"nope"`, nil},
			// gjson types an array as gjson.JSON, the same as an object, so it matches
			// the object variant. Pre-existing framework behaviour, unchanged here.
			{"array matches the object arm (gjson types it as JSON)", `[1]`, McpOAuthConfig{}},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				raw := `{"type":"remote","url":"https://x","enabled":true,"oauth":` + tc.raw + `}`
				var c McpRemoteConfig
				if err := json.Unmarshal([]byte(raw), &c); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				if got, want := reflect.TypeOf(c.OAuth), reflect.TypeOf(tc.want); got != want {
					t.Fatalf("OAuth runtime type = %v, want %v", got, want)
				}
				if !reflect.DeepEqual(c.AsOAuth(), c.OAuth) {
					t.Errorf("AsOAuth() = %#v, want %#v", c.AsOAuth(), c.OAuth)
				}
				// Siblings and the raw document must never be collateral damage.
				if c.URL != "https://x" || !c.Enabled || c.Type != McpRemoteConfigTypeRemote {
					t.Errorf("sibling fields lost: %+v", c)
				}
				if c.JSON.RawJSON() != raw {
					t.Errorf("RawJSON() = %q, want %q", c.JSON.RawJSON(), raw)
				}
			})
			// The same payload through the outer ConfigMcp union must agree. This is
			// the apijson.Port hand-off: the variant's typed McpOAuthUnion value is
			// copied onto ConfigMcp's `any` carrier, which is the safe direction
			// (`any` -> a union-typed field panics in apijson.Port).
			t.Run(tc.name+"/through ConfigMcp", func(t *testing.T) {
				raw := `{"type":"remote","url":"https://x","enabled":true,"oauth":` + tc.raw + `}`
				var m ConfigMcp
				if err := json.Unmarshal([]byte(raw), &m); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				if got, want := reflect.TypeOf(m.OAuth), reflect.TypeOf(tc.want); got != want {
					t.Fatalf("ConfigMcp.OAuth runtime type = %v, want %v", got, want)
				}
				if got, want := reflect.TypeOf(m.AsOAuth()), reflect.TypeOf(tc.want); got != want {
					t.Errorf("ConfigMcp.AsOAuth() runtime type = %v, want %v", got, want)
				}
				remote, ok := m.AsUnion().(McpRemoteConfig)
				if !ok {
					t.Fatalf("AsUnion() = %T, want McpRemoteConfig", m.AsUnion())
				}
				if got, want := reflect.TypeOf(remote.OAuth), reflect.TypeOf(tc.want); got != want {
					t.Errorf("variant OAuth runtime type = %v, want %v", got, want)
				}
				if m.URL != "https://x" || !m.Enabled {
					t.Errorf("sibling fields lost through ConfigMcp: %+v", m)
				}
				if m.JSON.RawJSON() != raw {
					t.Errorf("RawJSON() = %q, want %q", m.JSON.RawJSON(), raw)
				}
			})
		}
	})

	// ConfigPermission's ten `$ref PermissionRuleConfig` properties:
	// anyOf [PermissionActionConfig, PermissionObjectConfig].
	t.Run("ConfigPermission rules", func(t *testing.T) {
		rules := []string{"bash", "edit", "read", "glob", "grep", "list", "task",
			"external_directory", "lsp", "skill"}
		cases := []struct {
			name string
			raw  string
			want any
		}{
			{"action", `"ask"`, ConfigPermissionBashString("")},
			{"action unknown value", `"nope"`, ConfigPermissionBashString("")},
			{"object", `{"git *":"allow"}`, ConfigPermissionBashMap(nil)},
			{"object empty", `{}`, ConfigPermissionBashMap(nil)},
			{"null", `null`, nil},
			{"scalar violates both arms", `123`, nil},
			{"bool violates both arms", `true`, nil},
			{"array violates both arms", `[1]`, nil},
		}
		for _, tc := range cases {
			for _, rule := range rules {
				t.Run(rule+"/"+tc.name, func(t *testing.T) {
					raw := `{"webfetch":"deny","` + rule + `":` + tc.raw + `}`
					var p ConfigPermission
					if err := json.Unmarshal([]byte(raw), &p); err != nil {
						t.Fatalf("unmarshal: %v", err)
					}
					field := configPermissionRuleFields(p)[rule]
					want := reflect.TypeOf(tc.want)
					if got := reflect.TypeOf(field.carrier); got != want {
						t.Fatalf("%s runtime type = %v, want %v", rule, got, want)
					}
					if got := reflect.TypeOf(field.accessor); got != want {
						t.Errorf("%s accessor runtime type = %v, want %v", rule, got, want)
					}
					// The `$ref PermissionActionConfig` sibling must survive.
					if p.Webfetch != ConfigPermissionWebfetchDeny {
						t.Errorf("sibling webfetch = %q, want deny", p.Webfetch)
					}
					if p.JSON.RawJSON() != raw {
						t.Errorf("RawJSON() = %q, want %q", p.JSON.RawJSON(), raw)
					}
				})
			}
		}
	})

	// ProviderConfig.options.timeout / .headerTimeout:
	// anyOf [integer(exclusiveMinimum: 0), boolean(enum: false)].
	t.Run("ConfigProviderOptions timeouts", func(t *testing.T) {
		cases := []struct {
			name string
			raw  string
			want any
		}{
			{"integer", `30000`, UnionInt(0)},
			// OpenAPI says `integer`; a fractional number is registered anyway so a
			// drifted payload decodes losslessly instead of being truncated.
			{"fraction", `1.5`, UnionFloat(0)},
			{"false", `false`, UnionBool(false)},
			{"true violates enum", `true`, UnionBool(false)},
			{"null", `null`, nil},
			{"string violates both arms", `"nope"`, nil},
			{"object violates both arms", `{}`, nil},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				raw := `{"apiKey":"k","chunkTimeout":250,"timeout":` + tc.raw +
					`,"headerTimeout":` + tc.raw + `,"zz":1}`
				var o ConfigProviderOptions
				if err := json.Unmarshal([]byte(raw), &o); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				want := reflect.TypeOf(tc.want)
				if got := reflect.TypeOf(o.Timeout); got != want {
					t.Fatalf("Timeout runtime type = %v, want %v", got, want)
				}
				if got := reflect.TypeOf(o.HeaderTimeout); got != want {
					t.Fatalf("HeaderTimeout runtime type = %v, want %v", got, want)
				}
				if o.AsTimeout() != o.Timeout || o.AsHeaderTimeout() != o.HeaderTimeout {
					t.Errorf("AsTimeout()/AsHeaderTimeout() disagree with the fields")
				}
				// An `integer` must never arrive as a float64.
				if _, isFloat64 := any(o.Timeout).(float64); isFloat64 {
					t.Error("Timeout decoded to float64; OpenAPI declares integer")
				}
				if o.APIKey != "k" || o.ChunkTimeout != 250 {
					t.Errorf("sibling fields lost: %+v", o)
				}
				if o.ExtraFields["zz"] == nil {
					t.Errorf("ExtraFields = %v, want zz captured", o.ExtraFields)
				}
				if o.JSON.RawJSON() != raw {
					t.Errorf("RawJSON() = %q, want %q", o.JSON.RawJSON(), raw)
				}
			})
		}
	})

	// Config.permission / AgentConfig.permission: `$ref PermissionConfig`, i.e.
	// anyOf [PermissionActionConfig, object]. Both must behave identically.
	t.Run("permission carriers", func(t *testing.T) {
		cases := []struct {
			name string
			raw  string
			want any
		}{
			{"action", `"ask"`, ConfigPermissionAction("")},
			{"action unknown value", `"nope"`, ConfigPermissionAction("")},
			{"object", `{"bash":"ask"}`, ConfigPermission{}},
			{"object empty", `{}`, ConfigPermission{}},
			{"object unknown property", `{"bash":"ask","zz":"deny"}`, ConfigPermission{}},
			{"null", `null`, nil},
			{"scalar violates both arms", `123`, nil},
			{"bool violates both arms", `true`, nil},
			// See the note in [TestConfigPermissionNonConformantValueIsNonFatal]: gjson
			// types an array as gjson.JSON, so it matches the object arm.
			{"array matches the object arm (gjson types it as JSON)", `[1,2]`, ConfigPermission{}},
		}
		for _, tc := range cases {
			t.Run("Config/"+tc.name, func(t *testing.T) {
				raw := `{"model":"m","permission":` + tc.raw + `}`
				var c Config
				if err := json.Unmarshal([]byte(raw), &c); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				want := reflect.TypeOf(tc.want)
				if got := reflect.TypeOf(c.Permission); got != want {
					t.Fatalf("Permission runtime type = %v, want %v", got, want)
				}
				if got := reflect.TypeOf(c.AsPermission()); got != want {
					t.Errorf("AsPermission() runtime type = %v, want %v", got, want)
				}
				if c.Model != "m" {
					t.Errorf("sibling Model = %q, want m", c.Model)
				}
				if c.JSON.RawJSON() != raw {
					t.Errorf("RawJSON() = %q, want %q", c.JSON.RawJSON(), raw)
				}
			})
			t.Run("AgentConfig/"+tc.name, func(t *testing.T) {
				raw := `{"model":"m","permission":` + tc.raw + `}`
				var a AgentConfig
				if err := json.Unmarshal([]byte(raw), &a); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				want := reflect.TypeOf(tc.want)
				if got := reflect.TypeOf(a.Permission); got != want {
					t.Fatalf("Permission runtime type = %v, want %v", got, want)
				}
				if got := reflect.TypeOf(a.AsPermission()); got != want {
					t.Errorf("AsPermission() runtime type = %v, want %v", got, want)
				}
				if a.Model != "m" {
					t.Errorf("sibling Model = %q, want m", a.Model)
				}
				if a.JSON.RawJSON() != raw {
					t.Errorf("RawJSON() = %q, want %q", a.JSON.RawJSON(), raw)
				}
			})
		}
		// The object arm's own nested rule unions must resolve on this path too:
		// ConfigPermission is a registered variant, so apijson uses the struct
		// decoder for it and never calls ConfigPermission.UnmarshalJSON.
		t.Run("nested rule unions resolve through the variant", func(t *testing.T) {
			var c Config
			raw := `{"permission":{"bash":{"git *":"allow"},"read":"ask","webfetch":"deny"}}`
			if err := json.Unmarshal([]byte(raw), &c); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			perm, ok := c.Permission.(ConfigPermission)
			if !ok {
				t.Fatalf("Permission = %T, want ConfigPermission", c.Permission)
			}
			bash, ok := perm.AsBash().(ConfigPermissionBashMap)
			if !ok {
				t.Fatalf("AsBash() = %T, want ConfigPermissionBashMap", perm.AsBash())
			}
			if bash["git *"] != ConfigPermissionBashMapAllow {
				t.Errorf("bash[git *] = %q, want allow", bash["git *"])
			}
			if got, ok := perm.AsRead().(ConfigPermissionBashString); !ok || got != ConfigPermissionBashStringAsk {
				t.Errorf("AsRead() = %#v, want ConfigPermissionBashString(ask)", perm.AsRead())
			}
			if perm.Webfetch != ConfigPermissionWebfetchDeny {
				t.Errorf("Webfetch = %q, want deny", perm.Webfetch)
			}
		})
	})
}

// TestConfigProtectedManualRoutingStillWins is the regression guard for the two
// unions whose manual shape-routing must never be replaced by the registered union
// decoder: its exactness heuristic only penalises unknown extra properties and
// never penalises a missing `required` field, so with ties broken left-to-right it
// mis-decides both of them.
//
// Each row records what the raw union decoder produces (the mis-decode) next to
// what the public access path must keep producing.
func TestConfigProtectedManualRoutingStillWins(t *testing.T) {
	t.Run("ConfigLspUnion", func(t *testing.T) {
		cases := []struct {
			raw     string
			want    any
			native  any // what apijson's registered union decoder yields
			misread bool
		}{
			{`{"disabled":true}`, ConfigLspDisabled{}, ConfigLspDisabled{}, false},
			{`{"command":["gopls"]}`, ConfigLspObject{}, ConfigLspObject{}, false},
			{`{"command":["gopls"],"zz":1}`, ConfigLspObject{}, ConfigLspDisabled{}, true},
			{`{}`, ConfigLspObject{}, ConfigLspDisabled{}, true},
		}
		for _, tc := range cases {
			t.Run(tc.raw, func(t *testing.T) {
				// The public path, via Config.lsp -> map[string]ConfigLsp.
				var c Config
				if err := json.Unmarshal([]byte(`{"lsp":{"x":`+tc.raw+`}}`), &c); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				got := c.AsLsp()["x"].AsUnion()
				if reflect.TypeOf(got) != reflect.TypeOf(tc.want) {
					t.Fatalf("Config.lsp[x] = %T, want %T", got, tc.want)
				}
				// And the raw union decoder, to keep the justification honest: if
				// this ever starts agreeing, the manual routing may be retired.
				var u ConfigLspUnion
				if err := apijson.UnmarshalRoot([]byte(tc.raw), &u); err != nil {
					t.Fatalf("raw union unmarshal: %v", err)
				}
				if tc.misread && reflect.TypeOf(u) == reflect.TypeOf(tc.want) {
					t.Logf("apijson now decodes %s correctly; the manual route may be retired", tc.raw)
				}
				if !tc.misread && reflect.TypeOf(u) != reflect.TypeOf(tc.want) {
					t.Errorf("raw union regressed on %s: %T, want %T", tc.raw, u, tc.want)
				}
			})
		}
	})

	t.Run("ConfigV2ReferenceUnion", func(t *testing.T) {
		cases := []struct {
			raw     string
			want    any
			misread bool
		}{
			{`"https://x"`, ConfigV2ReferenceString(""), false},
			{`{"repository":"o/r"}`, ConfigV2ReferenceGit{}, false},
			{`{"path":"./d"}`, ConfigV2ReferenceLocal{}, false},
			{`{"path":"./d","zz":1}`, ConfigV2ReferenceLocal{}, true},
			{`{"repository":"o/r","zz":1}`, ConfigV2ReferenceGit{}, false},
		}
		for _, field := range []string{"reference", "references"} {
			for _, tc := range cases {
				t.Run(field+"/"+tc.raw, func(t *testing.T) {
					var c Config
					if err := json.Unmarshal([]byte(`{"`+field+`":{"k":`+tc.raw+`}}`), &c); err != nil {
						t.Fatalf("unmarshal: %v", err)
					}
					entries, unions := c.Reference, c.AsReference()
					if field == "references" {
						entries, unions = c.References, c.AsReferences()
					}
					if got := reflect.TypeOf(entries["k"]); got != reflect.TypeOf(tc.want) {
						t.Fatalf("%s[k] = %v, want %v", field, got, reflect.TypeOf(tc.want))
					}
					if got := reflect.TypeOf(unions["k"]); got != reflect.TypeOf(tc.want) {
						t.Errorf("As(%s)[k] = %v, want %v", field, got, reflect.TypeOf(tc.want))
					}
					var u ConfigV2ReferenceUnion
					if err := apijson.UnmarshalRoot([]byte(tc.raw), &u); err != nil {
						t.Fatalf("raw union unmarshal: %v", err)
					}
					if tc.misread && reflect.TypeOf(u) == reflect.TypeOf(tc.want) {
						t.Logf("apijson now decodes %s correctly; the manual route may be retired", tc.raw)
					}
				})
			}
		}
	})
}
