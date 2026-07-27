package opencode

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/sst/opencode-sdk-go/shared"
)

// Aligned with OpenAPI GET /mcp + JS SDK(v2) Mcp.status.
func TestMcpStatusParamsQuery(t *testing.T) {
	p := McpStatusParams{Directory: F("d"), Workspace: F("w")}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// Aligned with OpenAPI POST /mcp + JS SDK(v2) Mcp.add.
// body required: name, config; query: directory, workspace
func TestMcpAddParamsBodyAndQuery(t *testing.T) {
	t.Run("local config body excludes query", func(t *testing.T) {
		p := McpAddParams{
			Directory: F("d"),
			Workspace: F("w"),
			Name:      F("my-server"),
			Config: F[McpAddParamsConfigUnion](McpAddParamsConfigLocal{
				Type:    F(McpLocalConfigTypeLocal),
				Command: F([]string{"node", "server.js"}),
				Cwd:     F("/srv"),
				Enabled: F(true),
				Timeout: F(int64(5000)),
			}),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
			t.Errorf("query fields leaked into body: %s", got)
		}
		if !strings.Contains(got, `"name":"my-server"`) {
			t.Errorf("name missing: %s", got)
		}
		if !strings.Contains(got, `"type":"local"`) || !strings.Contains(got, `"command":["node","server.js"]`) {
			t.Errorf("local config missing: %s", got)
		}
		if !strings.Contains(got, `"cwd":"/srv"`) || !strings.Contains(got, `"timeout":5000`) {
			t.Errorf("optional local fields missing: %s", got)
		}
	})

	t.Run("remote config with oauth object", func(t *testing.T) {
		p := McpAddParams{
			Name: F("remote-server"),
			Config: F[McpAddParamsConfigUnion](McpAddParamsConfigRemote{
				Type:    F(McpRemoteConfigTypeRemote),
				URL:     F("https://mcp.example.com"),
				Enabled: F(true),
				Headers: F(map[string]string{"Authorization": "Bearer x"}),
				OAuth: F[McpAddParamsConfigRemoteOAuthUnion](McpAddParamsConfigRemoteOAuth{
					ClientID:     F("cid"),
					Scope:        F("read"),
					CallbackPort: F(int64(8080)),
					RedirectURI:  F("http://localhost:8080/cb"),
				}),
			}),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"type":"remote"`) || !strings.Contains(got, `"url":"https://mcp.example.com"`) {
			t.Errorf("remote config missing: %s", got)
		}
		if !strings.Contains(got, `"clientId":"cid"`) || !strings.Contains(got, `"callbackPort":8080`) {
			t.Errorf("oauth object missing: %s", got)
		}
		if !strings.Contains(got, `"redirectUri":"http://localhost:8080/cb"`) {
			t.Errorf("redirectUri missing: %s", got)
		}
	})

	t.Run("remote config with oauth disabled (false)", func(t *testing.T) {
		p := McpAddParams{
			Name: F("remote-server"),
			Config: F[McpAddParamsConfigUnion](McpAddParamsConfigRemote{
				Type:  F(McpRemoteConfigTypeRemote),
				URL:   F("https://mcp.example.com"),
				OAuth: F[McpAddParamsConfigRemoteOAuthUnion](shared.UnionBool(false)),
			}),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		if !strings.Contains(got, `"oauth":false`) {
			t.Errorf("oauth should serialize to false, got: %s", got)
		}
	})

	t.Run("query serialization", func(t *testing.T) {
		p := McpAddParams{Directory: F("d"), Workspace: F("w"), Name: F("n")}
		got := p.URLQuery().Encode()
		want := "directory=d&workspace=w"
		if got != want {
			t.Errorf("query got %q, want %q", got, want)
		}
	})
}

// Aligned with OpenAPI POST /mcp/{name}/auth/callback + JS SDK(v2) Auth2.callback.
// body required: code; query: directory, workspace
func TestMcpAuthCallbackParamsBodyAndQuery(t *testing.T) {
	p := McpAuthCallbackParams{
		Directory: F("d"),
		Workspace: F("w"),
		Code:      F("auth-code-123"),
	}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)
	want := `{"code":"auth-code-123"}`
	if got != want {
		t.Errorf("body got %s, want %s", got, want)
	}
	if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
		t.Errorf("query fields leaked into body: %s", got)
	}
	if q := p.URLQuery().Encode(); q != "directory=d&workspace=w" {
		t.Errorf("query got %q", q)
	}
}

// Query-only params for connect/disconnect/auth start/authenticate/remove.
func TestMcpQueryOnlyParams(t *testing.T) {
	if q := (McpConnectParams{Directory: F("d")}).URLQuery().Encode(); q != "directory=d" {
		t.Errorf("connect: %q", q)
	}
	if q := (McpDisconnectParams{Workspace: F("w")}).URLQuery().Encode(); q != "workspace=w" {
		t.Errorf("disconnect: %q", q)
	}
	if q := (McpAuthStartParams{Directory: F("d"), Workspace: F("w")}).URLQuery().Encode(); q != "directory=d&workspace=w" {
		t.Errorf("authStart: %q", q)
	}
	if q := (McpAuthAuthenticateParams{Directory: F("d")}).URLQuery().Encode(); q != "directory=d" {
		t.Errorf("authenticate: %q", q)
	}
	if q := (McpAuthRemoveParams{Workspace: F("w")}).URLQuery().Encode(); q != "workspace=w" {
		t.Errorf("authRemove: %q", q)
	}
}

// Aligned with OpenAPI MCPStatus anyOf (connected/disabled/failed/needs_auth/
// needs_client_registration). The top-level McpStatus.Status/.Error fields are
// populated by apijson.Port for every variant regardless of union routing.
func TestMcpStatusFieldsUnmarshal(t *testing.T) {
	cases := []struct {
		name     string
		json     string
		wantStat McpStatusStatus
		wantErr  string
	}{
		{"connected", `{"status":"connected"}`, McpStatusStatusConnected, ""},
		{"disabled", `{"status":"disabled"}`, McpStatusStatusDisabled, ""},
		{"failed", `{"status":"failed","error":"connection refused"}`, McpStatusStatusFailed, "connection refused"},
		{"needs_auth", `{"status":"needs_auth"}`, McpStatusStatusNeedsAuth, ""},
		{"needs_client_registration", `{"status":"needs_client_registration","error":"register first"}`, McpStatusStatusNeedsClientRegistration, "register first"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s McpStatus
			if err := json.Unmarshal([]byte(tc.json), &s); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if s.Status != tc.wantStat {
				t.Errorf("Status: got %q, want %q", s.Status, tc.wantStat)
			}
			if s.Error != tc.wantErr {
				t.Errorf("Error: got %q, want %q", s.Error, tc.wantErr)
			}
			if s.JSON.raw == "" {
				t.Error("RawJSON not preserved")
			}
		})
	}
}

// Aligned with OpenAPI MCPStatus anyOf, discriminated by the `status` const.
// McpStatus.AsUnion() must return the exact variant type for all 5 states.
// Regression guard: an empty discriminatorKey lets the exactness heuristic route
// disabled/needs_auth -> Connected and needs_client_registration -> Failed, because
// all variants share the McpStatusStatus enum (IsKnown() is true for every value),
// so only discriminatorKey="status" + per-variant DiscriminatorValue can distinguish.
func TestMcpStatusAsUnionDiscriminates(t *testing.T) {
	cases := []struct {
		name     string
		json     string
		wantType any
	}{
		{"connected", `{"status":"connected"}`, McpStatusConnected{}},
		{"disabled", `{"status":"disabled"}`, McpStatusDisabled{}},
		{"failed", `{"status":"failed","error":"connection refused"}`, McpStatusFailed{}},
		{"needs_auth", `{"status":"needs_auth"}`, McpStatusNeedsAuth{}},
		{"needs_client_registration", `{"status":"needs_client_registration","error":"register first"}`, McpStatusNeedsClientRegistration{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s McpStatus
			if err := json.Unmarshal([]byte(tc.json), &s); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			got := reflect.TypeOf(s.AsUnion())
			want := reflect.TypeOf(tc.wantType)
			if got != want {
				t.Errorf("AsUnion() type: got %v, want %v", got, want)
			}
		})
	}
}

// Each MCPStatus variant struct unmarshals correctly when decoded directly.
func TestMcpStatusVariantStructsUnmarshal(t *testing.T) {
	var connected McpStatusConnected
	if err := json.Unmarshal([]byte(`{"status":"connected"}`), &connected); err != nil {
		t.Fatal(err)
	}
	if connected.Status != McpStatusStatusConnected {
		t.Errorf("connected.Status = %q", connected.Status)
	}

	var failed McpStatusFailed
	if err := json.Unmarshal([]byte(`{"status":"failed","error":"boom"}`), &failed); err != nil {
		t.Fatal(err)
	}
	if failed.Status != McpStatusStatusFailed || failed.Error != "boom" {
		t.Errorf("failed = %+v", failed)
	}

	var reg McpStatusNeedsClientRegistration
	if err := json.Unmarshal([]byte(`{"status":"needs_client_registration","error":"reg"}`), &reg); err != nil {
		t.Fatal(err)
	}
	if reg.Status != McpStatusStatusNeedsClientRegistration || reg.Error != "reg" {
		t.Errorf("reg = %+v", reg)
	}
}

// Aligned with OpenAPI map<string, MCPStatus> for GET /mcp and POST /mcp response.
func TestMcpStatusMapUnmarshal(t *testing.T) {
	raw := `{
		"server-a":{"status":"connected"},
		"server-b":{"status":"failed","error":"boom"}
	}`
	var m map[string]McpStatus
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		t.Fatal(err)
	}
	if len(m) != 2 {
		t.Fatalf("map len = %d", len(m))
	}
	if m["server-a"].Status != McpStatusStatusConnected {
		t.Errorf("server-a = %+v", m["server-a"])
	}
	if m["server-b"].Status != McpStatusStatusFailed || m["server-b"].Error != "boom" {
		t.Errorf("server-b = %+v", m["server-b"])
	}
}

// Aligned with OpenAPI POST /mcp/{name}/auth 200 (McpAuthStartResponse).
// required: authorizationUrl, oauthState
func TestMcpAuthStartResponseUnmarshal(t *testing.T) {
	raw := `{"authorizationUrl":"https://auth.example.com/authorize","oauthState":"state-xyz"}`
	var r McpAuthStartResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatal(err)
	}
	if r.AuthorizationURL != "https://auth.example.com/authorize" {
		t.Errorf("authorizationUrl = %q", r.AuthorizationURL)
	}
	if r.OAuthState != "state-xyz" {
		t.Errorf("oauthState = %q", r.OAuthState)
	}
}

// Aligned with OpenAPI DELETE /mcp/{name}/auth 200 (McpAuthRemoveResponse).
// required: success
func TestMcpAuthRemoveResponseUnmarshal(t *testing.T) {
	raw := `{"success":true}`
	var r McpAuthRemoveResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatal(err)
	}
	if !r.Success {
		t.Error("success should be true")
	}
}

func TestMcpStatusStatusIsKnown(t *testing.T) {
	for _, s := range []McpStatusStatus{
		McpStatusStatusConnected, McpStatusStatusDisabled, McpStatusStatusFailed,
		McpStatusStatusNeedsAuth, McpStatusStatusNeedsClientRegistration,
	} {
		if !s.IsKnown() {
			t.Errorf("%q should be known", s)
		}
	}
	if McpStatusStatus("unknown").IsKnown() {
		t.Error("unknown should not be known")
	}
}

func TestMcpConfigTypeIsKnown(t *testing.T) {
	if !McpLocalConfigTypeLocal.IsKnown() {
		t.Error("local should be known")
	}
	if !McpRemoteConfigTypeRemote.IsKnown() {
		t.Error("remote should be known")
	}
}
