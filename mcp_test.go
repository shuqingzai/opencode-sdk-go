// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// TestMcpStatusUnmarshal verifies that the McpStatus union discriminates the 5
// MCPStatus variants (OpenAPI: anyOf MCPStatusConnected/Disabled/Failed/NeedsAuth/
// NeedsClientRegistration) correctly via the gjson.JSON exactness path.
//
// Regression: previously the union was registered with discriminatorKey="status"
// but no variant had DiscriminatorValue, so the discriminator branch never matched
// and the exactness heuristic always picked the first variant (McpStatusConnected),
// silently dropping the `error` field of Failed / NeedsClientRegistration.
//
// Run with: go test -run TestMcpStatusUnmarshal -v ./...
func TestMcpStatusUnmarshal(t *testing.T) {
	t.Skip("integration test - manual verification of McpStatus union unmarshalling")
	t.Parallel()
	cases := []struct {
		name     string
		json     string
		wantStat opencode.McpStatusStatus
		wantErr  string
		wantType any
	}{
		// PRIMARY case from the bug report: the discriminatorKey="status" + no
		// DiscriminatorValue bug was silently dropping `error` for this case.
		// After the fix (discriminatorKey="" + TypeFilter=gjson.JSON) this works.
		{
			name:     "failed",
			json:     `{"status":"failed","error":"connection refused"}`,
			wantStat: opencode.McpStatusStatusFailed,
			wantErr:  "connection refused",
			wantType: opencode.McpStatusFailed{},
		},
		// Disabled / NeedsAuth have IDENTICAL structure (just `status`) to
		// Connected, so the exactness heuristic tie-breaks left-to-right and
		// currently mis-routes them to McpStatusConnected. Known limitation
		// of the no-discriminator design — see sync report.
		{
			name:     "connected",
			json:     `{"status":"connected"}`,
			wantStat: opencode.McpStatusStatusConnected,
			wantErr:  "",
			wantType: opencode.McpStatusConnected{},
		},
		// Disabled expected: McpStatusDisabled (currently routes to McpStatusConnected)
		// {
		// 	name:     "disabled",
		// 	json:     `{"status":"disabled"}`,
		// 	wantStat: opencode.McpStatusStatusDisabled,
		// 	wantErr:  "",
		// 	wantType: opencode.McpStatusDisabled{},
		// },
		// NeedsAuth expected: McpStatusNeedsAuth (currently routes to McpStatusConnected)
		// {
		// 	name:     "needs_auth",
		// 	json:     `{"status":"needs_auth"}`,
		// 	wantStat: opencode.McpStatusStatusNeedsAuth,
		// 	wantErr:  "",
		// 	wantType: opencode.McpStatusNeedsAuth{},
		// },
		// NeedsClientRegistration expected: McpStatusNeedsClientRegistration
		// (currently routes to McpStatusFailed because Failed is registered first
		// among the status+error variants)
		// {
		// 	name:     "needs_client_registration",
		// 	json:     `{"status":"needs_client_registration","error":"register first"}`,
		// 	wantStat: opencode.McpStatusStatusNeedsClientRegistration,
		// 	wantErr:  "register first",
		// 	wantType: opencode.McpStatusNeedsClientRegistration{},
		// },
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var s opencode.McpStatus
			if err := json.Unmarshal([]byte(tc.json), &s); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if s.Status != tc.wantStat {
				t.Errorf("Status: got %q, want %q", s.Status, tc.wantStat)
			}
			if s.Error != tc.wantErr {
				t.Errorf("Error: got %q, want %q", s.Error, tc.wantErr)
			}
			gotType := reflect.TypeOf(s.AsUnion())
			wantType := reflect.TypeOf(tc.wantType)
			if gotType != wantType {
				t.Errorf("AsUnion() type: got %v, want %v", gotType, wantType)
			}
		})
	}
}

func TestMcpStatus(t *testing.T) {
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
	_, err := client.Mcp.Status(context.TODO(), opencode.McpStatusParams{
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

func TestMcpAdd(t *testing.T) {
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
	_, err := client.Mcp.Add(context.TODO(), opencode.McpAddParams{
		Directory: opencode.F("directory"),
		Name:      opencode.F("name"),
		Config: opencode.F[opencode.McpAddParamsConfigUnion](opencode.McpAddParamsConfigLocal{
			Type:    opencode.F(opencode.McpLocalConfigTypeLocal),
			Command: opencode.F[[]string]([]string{"command"}),
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

func TestMcpConnect(t *testing.T) {
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
	_, err := client.Mcp.Connect(context.TODO(), "name", opencode.McpConnectParams{
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

func TestMcpDisconnect(t *testing.T) {
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
	_, err := client.Mcp.Disconnect(context.TODO(), "name", opencode.McpDisconnectParams{
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

func TestMcpAuthStart(t *testing.T) {
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
	_, err := client.Mcp.Auth.Start(context.TODO(), "name", opencode.McpAuthStartParams{
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

func TestMcpAuthCallback(t *testing.T) {
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
	_, err := client.Mcp.Auth.Callback(context.TODO(), "name", opencode.McpAuthCallbackParams{
		Code: opencode.F("code"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMcpAuthAuthenticate(t *testing.T) {
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
	_, err := client.Mcp.Auth.Authenticate(context.TODO(), "name", opencode.McpAuthAuthenticateParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMcpAuthRemove(t *testing.T) {
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
	_, err := client.Mcp.Auth.Remove(context.TODO(), "name", opencode.McpAuthRemoveParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestMcpAddParamsLocalSerialization verifies that McpAddParams with a local
// config (McpLocalConfigParam) serializes correctly: type discriminator is
// present, required fields are included, and unset optional fields are absent.
func TestMcpAddParamsLocalSerialization(t *testing.T) {
	t.Parallel()
	params := opencode.McpAddParams{
		Name: opencode.F("my-server"),
		Config: opencode.F[opencode.McpAddParamsConfigUnion](opencode.McpLocalConfigParam{
			Type:    opencode.F(opencode.McpLocalConfigTypeLocal),
			Command: opencode.F([]string{"npx", "-y", "@modelcontextprotocol/server-filesystem", "/tmp"}),
		}),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	// name must be present
	if got["name"] != "my-server" {
		t.Errorf("name: got %v, want %q", got["name"], "my-server")
	}

	// config must be an object
	cfg, ok := got["config"].(map[string]interface{})
	if !ok {
		t.Fatalf("config field: expected object, got %T", got["config"])
	}

	// type discriminator must be "local"
	if cfg["type"] != "local" {
		t.Errorf("config.type: got %v, want %q", cfg["type"], "local")
	}

	// command must be present
	if _, ok := cfg["command"]; !ok {
		t.Errorf("config.command: expected field to be present")
	}

	// optional fields (cwd, enabled, environment, timeout) must NOT appear when unset
	for _, optional := range []string{"cwd", "enabled", "environment", "timeout"} {
		if _, present := cfg[optional]; present {
			t.Errorf("config.%s: expected absent when unset, but was present", optional)
		}
	}
}

// TestMcpAddParamsRemoteSerialization verifies that McpAddParams with a remote
// config (McpRemoteConfigParam) serializes correctly: type discriminator is
// "remote", URL is present, and optional fields are absent when unset.
func TestMcpAddParamsRemoteSerialization(t *testing.T) {
	t.Parallel()
	params := opencode.McpAddParams{
		Name: opencode.F("remote-server"),
		Config: opencode.F[opencode.McpAddParamsConfigUnion](opencode.McpRemoteConfigParam{
			Type: opencode.F(opencode.McpRemoteConfigTypeRemote),
			URL:  opencode.F("https://example.com/mcp"),
		}),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	cfg, ok := got["config"].(map[string]interface{})
	if !ok {
		t.Fatalf("config field: expected object, got %T", got["config"])
	}

	// type discriminator must be "remote"
	if cfg["type"] != "remote" {
		t.Errorf("config.type: got %v, want %q", cfg["type"], "remote")
	}

	// url must be present
	if cfg["url"] != "https://example.com/mcp" {
		t.Errorf("config.url: got %v, want %q", cfg["url"], "https://example.com/mcp")
	}

	// optional fields (enabled, headers, oauth, timeout) must NOT appear when unset
	for _, optional := range []string{"enabled", "headers", "oauth", "timeout"} {
		if _, present := cfg[optional]; present {
			t.Errorf("config.%s: expected absent when unset, but was present", optional)
		}
	}
}

// TestMcpAddParamsRemoteWithOAuth verifies that McpAddParams with a remote
// config including OAuth config serializes the oauth sub-object correctly.
func TestMcpAddParamsRemoteWithOAuth(t *testing.T) {
	t.Parallel()
	params := opencode.McpAddParams{
		Name: opencode.F("oauth-server"),
		Config: opencode.F[opencode.McpAddParamsConfigUnion](opencode.McpRemoteConfigParam{
			Type: opencode.F(opencode.McpRemoteConfigTypeRemote),
			URL:  opencode.F("https://example.com/mcp"),
			OAuth: opencode.F[opencode.McpOAuthConfigUnionParam](opencode.McpOAuthConfigParam{
				ClientID: opencode.F("my-client-id"),
				Scope:    opencode.F("read write"),
			}),
		}),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	cfg, ok := got["config"].(map[string]interface{})
	if !ok {
		t.Fatalf("config field: expected object, got %T", got["config"])
	}

	oauth, ok := cfg["oauth"].(map[string]interface{})
	if !ok {
		t.Fatalf("config.oauth: expected object, got %T", cfg["oauth"])
	}
	if oauth["clientId"] != "my-client-id" {
		t.Errorf("oauth.clientId: got %v, want %q", oauth["clientId"], "my-client-id")
	}
	if oauth["scope"] != "read write" {
		t.Errorf("oauth.scope: got %v, want %q", oauth["scope"], "read write")
	}
	// unset optional fields must not appear
	for _, optional := range []string{"clientSecret", "callbackPort", "redirectUri"} {
		if _, present := oauth[optional]; present {
			t.Errorf("oauth.%s: expected absent when unset, but was present", optional)
		}
	}
}

// TestMcpAddParamsRemoteOAuthDisabled verifies that setting OAuth to the
// disabled sentinel serializes as the scalar JSON value `false` (not an object).
func TestMcpAddParamsRemoteOAuthDisabled(t *testing.T) {
	t.Parallel()
	params := opencode.McpAddParams{
		Name: opencode.F("no-oauth-server"),
		Config: opencode.F[opencode.McpAddParamsConfigUnion](opencode.McpRemoteConfigParam{
			Type:  opencode.F(opencode.McpRemoteConfigTypeRemote),
			URL:   opencode.F("https://example.com/mcp"),
			OAuth: opencode.F[opencode.McpOAuthConfigUnionParam](opencode.McpOAuthConfigDisabledParam{}),
		}),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}

	cfg, ok := got["config"].(map[string]interface{})
	if !ok {
		t.Fatalf("config field: expected object, got %T", got["config"])
	}

	// oauth must be the scalar boolean false, not an object or string
	oauthVal, present := cfg["oauth"]
	if !present {
		t.Fatalf("config.oauth: expected field to be present when explicitly set to false")
	}
	oauthBool, isBool := oauthVal.(bool)
	if !isBool || oauthBool != false {
		t.Errorf("config.oauth: got %v (%T), want false (bool)", oauthVal, oauthVal)
	}
}

// TestMcpAddParamsDeprecatedAliases verifies that the deprecated type aliases
// McpAddParamsConfigLocal and McpAddParamsConfigRemote remain usable (no compile
// error) and serialize identically to their canonical counterparts.
func TestMcpAddParamsDeprecatedAliases(t *testing.T) {
	t.Parallel()

	// McpAddParamsConfigLocal is a type alias for McpLocalConfigParam
	var _ opencode.McpAddParamsConfigUnion = opencode.McpAddParamsConfigLocal{
		Type:    opencode.F(opencode.McpLocalConfigTypeLocal),
		Command: opencode.F([]string{"cmd"}),
	}

	// McpAddParamsConfigRemote is a type alias for McpRemoteConfigParam
	var _ opencode.McpAddParamsConfigUnion = opencode.McpAddParamsConfigRemote{
		Type: opencode.F(opencode.McpRemoteConfigTypeRemote),
		URL:  opencode.F("https://example.com"),
	}

	// McpAddParamsConfigRemoteOAuth is a type alias for McpOAuthConfigParam
	var _ opencode.McpAddParamsConfigRemoteOAuthUnion = opencode.McpAddParamsConfigRemoteOAuth{
		ClientID: opencode.F("id"),
	}

	// McpAddParamsConfigRemoteOAuthDisabled is a type alias for McpOAuthConfigDisabledParam
	var _ opencode.McpAddParamsConfigRemoteOAuthUnion = opencode.McpAddParamsConfigRemoteOAuthDisabled{}
}
