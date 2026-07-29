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

// TestMcpRemoteConfigOAuthUnmarshal verifies that McpRemoteConfig.OAuth (an any
// field) correctly deserializes both the McpOAuthConfig object variant and the
// boolean false variant from the OpenAPI anyOf schema
// (`anyOf [McpOAuthConfig, boolean(enum: false)]`; JS SDK v2:
// `oauth?: McpOAuthConfig | false`).
//
// Both variants are routed through the registered [opencode.McpOAuthUnion] (see
// config.go), so the object variant resolves to [opencode.McpOAuthConfig] and the
// scalar variant to [opencode.McpOAuthDisabled].
//
// Regression: previously `OAuth` was a plain `any` field with no union routing at
// all, so internal/apijson decoded it generically via node.Value() —
// map[string]interface{} for the object variant — and every typed OAuth field was
// lost to callers. The union is registered with TypeFilter gjson.JSON for the
// object variant and gjson.False/gjson.True for the scalar variant, and
// McpRemoteConfig.UnmarshalJSON extracts the `oauth` sub-document with gjson
// before decoding it, so the outer ConfigMcpUnion routing is untouched.
func TestMcpRemoteConfigOAuthUnmarshal(t *testing.T) {
	t.Parallel()

	t.Run("oauth_false", func(t *testing.T) {
		data := `{"type":"remote","url":"https://example.com","oauth":false}`
		var cfg opencode.McpRemoteConfig
		if err := json.Unmarshal([]byte(data), &cfg); err != nil {
			t.Fatalf("unmarshal with oauth=false: %v", err)
		}
		if cfg.OAuth == nil {
			t.Fatal("OAuth field should not be nil when oauth=false")
		}
		oauthDisabled, ok := cfg.OAuth.(opencode.McpOAuthDisabled)
		if !ok {
			t.Fatalf("expected OAuth to be opencode.McpOAuthDisabled, got %T", cfg.OAuth)
		}
		if bool(oauthDisabled) != false {
			t.Errorf("expected OAuth=false, got %v", bool(oauthDisabled))
		}
		if !oauthDisabled.IsKnown() {
			t.Error("expected McpOAuthDisabled(false).IsKnown()=true")
		}
		if got := cfg.AsOAuth(); got != oauthDisabled {
			t.Errorf("AsOAuth() = %#v, want %#v", got, oauthDisabled)
		}
	})

	t.Run("oauth_config_object", func(t *testing.T) {
		data := `{"type":"remote","url":"https://example.com","oauth":{"clientId":"myid","scope":"openid","callbackPort":8080}}`
		var cfg opencode.McpRemoteConfig
		if err := json.Unmarshal([]byte(data), &cfg); err != nil {
			t.Fatalf("unmarshal with oauth object: %v", err)
		}
		if cfg.OAuth == nil {
			t.Fatal("OAuth field should not be nil when oauth is an object")
		}
		oauthConfig, ok := cfg.OAuth.(opencode.McpOAuthConfig)
		if !ok {
			t.Fatalf("expected OAuth to be opencode.McpOAuthConfig, got %T", cfg.OAuth)
		}
		if oauthConfig.ClientID != "myid" {
			t.Errorf("expected clientId=myid, got %v", oauthConfig.ClientID)
		}
		if oauthConfig.Scope != "openid" {
			t.Errorf("expected scope=openid, got %v", oauthConfig.Scope)
		}
		if oauthConfig.CallbackPort != 8080 {
			t.Errorf("expected callbackPort=8080, got %v", oauthConfig.CallbackPort)
		}
		if _, ok := cfg.AsOAuth().(opencode.McpOAuthConfig); !ok {
			t.Errorf("AsOAuth() = %T, want opencode.McpOAuthConfig", cfg.AsOAuth())
		}
	})

	t.Run("oauth_absent", func(t *testing.T) {
		data := `{"type":"remote","url":"https://example.com"}`
		var cfg opencode.McpRemoteConfig
		if err := json.Unmarshal([]byte(data), &cfg); err != nil {
			t.Fatalf("unmarshal without oauth: %v", err)
		}
		if cfg.OAuth != nil {
			t.Errorf("expected OAuth=nil when absent, got %v", cfg.OAuth)
		}
	})
}

// TestMcpStatusUnmarshal verifies that the McpStatus union discriminates all 5
// MCPStatus variants (OpenAPI: anyOf MCPStatusConnected/Disabled/Failed/NeedsAuth/
// NeedsClientRegistration) correctly via discriminatorKey="status"+DiscriminatorValue.
//
// Regression: previously the union was registered with discriminatorKey="status"
// but no variant had DiscriminatorValue, so the discriminator branch never matched
// and the exactness heuristic always picked the first variant (McpStatusConnected),
// silently dropping the `error` field of Failed / NeedsClientRegistration.
//
// Fix: RegisterUnion now uses discriminatorKey="status" with DiscriminatorValue set
// per variant ("connected", "disabled", "failed", "needs_auth",
// "needs_client_registration"). All 5 variants now route correctly.
//
// Run with: go test -run TestMcpStatusUnmarshal -v ./...
func TestMcpStatusUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		json     string
		wantStat opencode.McpStatusStatus
		wantErr  string
		wantType any
	}{
		{
			name:     "connected",
			json:     `{"status":"connected"}`,
			wantStat: opencode.McpStatusStatusConnected,
			wantErr:  "",
			wantType: opencode.McpStatusConnected{},
		},
		{
			name:     "disabled",
			json:     `{"status":"disabled"}`,
			wantStat: opencode.McpStatusStatusDisabled,
			wantErr:  "",
			wantType: opencode.McpStatusDisabled{},
		},
		{
			name:     "failed",
			json:     `{"status":"failed","error":"connection refused"}`,
			wantStat: opencode.McpStatusStatusFailed,
			wantErr:  "connection refused",
			wantType: opencode.McpStatusFailed{},
		},
		{
			name:     "needs_auth",
			json:     `{"status":"needs_auth"}`,
			wantStat: opencode.McpStatusStatusNeedsAuth,
			wantErr:  "",
			wantType: opencode.McpStatusNeedsAuth{},
		},
		{
			name:     "needs_client_registration",
			json:     `{"status":"needs_client_registration","error":"register first"}`,
			wantStat: opencode.McpStatusStatusNeedsClientRegistration,
			wantErr:  "register first",
			wantType: opencode.McpStatusNeedsClientRegistration{},
		},
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
