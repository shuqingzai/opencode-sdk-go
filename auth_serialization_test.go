package opencode

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sst/opencode-sdk-go/option"
)

// Aligned with OpenAPI schema "Auth" (anyOf: OAuth | ApiAuth | WellKnownAuth),
// used exclusively as the request body for auth.set (PUT /auth/{providerID}).
//
// Auth never appears in any response, so the three variants are pure Request
// types: every field is wrapped in param.Field[T] and serialization goes through
// apijson.MarshalRoot, which omits fields that were never set. This matches the
// JS SDK(v2), where optional properties are absent from the object and therefore
// dropped by JSON.stringify.
//
// apijson.MarshalRoot sorts keys alphabetically by their JSON name.

func TestAuthBodySerialization(t *testing.T) {
	t.Run("OAuth — all fields set", func(t *testing.T) {
		a := AuthParamOAuth{
			Type:          F(AuthParamOAuthTypeOAuth),
			Refresh:       F("r"),
			Access:        F("a"),
			Expires:       F(int64(3600)),
			AccountID:     F("acc"),
			EnterpriseURL: F("https://ghe.example.com"),
		}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"access":"a","accountId":"acc","enterpriseUrl":"https://ghe.example.com","expires":3600,"refresh":"r","type":"oauth"}`
		if got != want {
			t.Errorf("got  %s\nwant %s", got, want)
		}
	})

	t.Run("OAuth — required only, optional fields omitted", func(t *testing.T) {
		// JS SDK equivalent: JSON.stringify({type:"oauth",refresh:"r",access:"a",expires:0})
		// accountId / enterpriseUrl must NOT appear on the wire.
		a := AuthParamOAuth{
			Type:    F(AuthParamOAuthTypeOAuth),
			Refresh: F("r"),
			Access:  F("a"),
			Expires: F(int64(0)),
		}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"access":"a","expires":0,"refresh":"r","type":"oauth"}`
		if got != want {
			t.Errorf("got  %s\nwant %s", got, want)
		}
	})

	t.Run("ApiAuth — with metadata", func(t *testing.T) {
		a := AuthParamAPIAuth{
			Type:     F(AuthParamAPIAuthTypeAPI),
			Key:      F("sk-123"),
			Metadata: F(map[string]string{"region": "us"}),
		}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"key":"sk-123","metadata":{"region":"us"},"type":"api"}`
		if got != want {
			t.Errorf("got  %s\nwant %s", got, want)
		}
	})

	t.Run("ApiAuth — metadata omitted when unset", func(t *testing.T) {
		// OpenAPI declares metadata as a non-nullable optional object.
		// Emitting null would violate the schema; emitting {} would assert an
		// empty metadata map. The correct wire form is to omit the key.
		a := AuthParamAPIAuth{Type: F(AuthParamAPIAuthTypeAPI), Key: F("sk-123")}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"key":"sk-123","type":"api"}`
		if got != want {
			t.Errorf("got  %s\nwant %s", got, want)
		}
	})

	t.Run("ApiAuth — explicit empty metadata is preserved", func(t *testing.T) {
		// Distinguishing "unset" from "explicitly empty" is exactly what
		// param.Field[T] buys us.
		a := AuthParamAPIAuth{
			Type:     F(AuthParamAPIAuthTypeAPI),
			Key:      F("sk-123"),
			Metadata: F(map[string]string{}),
		}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"key":"sk-123","metadata":{},"type":"api"}`
		if got != want {
			t.Errorf("got  %s\nwant %s", got, want)
		}
	})

	t.Run("WellKnownAuth — all required", func(t *testing.T) {
		a := AuthParamWellKnownAuth{
			Type:  F(AuthParamWellKnownAuthTypeWellKnown),
			Key:   F("k"),
			Token: F("t"),
		}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"key":"k","token":"t","type":"wellknown"}`
		if got != want {
			t.Errorf("got  %s\nwant %s", got, want)
		}
	})

	t.Run("Auth interface accepts all three variants", func(t *testing.T) {
		var _ AuthParam = AuthParamOAuth{Type: F(AuthParamOAuthTypeOAuth)}
		var _ AuthParam = AuthParamAPIAuth{Type: F(AuthParamAPIAuthTypeAPI)}
		var _ AuthParam = AuthParamWellKnownAuth{Type: F(AuthParamWellKnownAuthTypeWellKnown)}
	})
}

// Discriminator enum values must match the OpenAPI `type` enums exactly.
func TestAuthTypeEnumsIsKnown(t *testing.T) {
	if !AuthParamOAuthTypeOAuth.IsKnown() || string(AuthParamOAuthTypeOAuth) != "oauth" {
		t.Errorf("AuthParamOAuthTypeOAuth = %q", AuthParamOAuthTypeOAuth)
	}
	if !AuthParamAPIAuthTypeAPI.IsKnown() || string(AuthParamAPIAuthTypeAPI) != "api" {
		t.Errorf("AuthParamAPIAuthTypeAPI = %q", AuthParamAPIAuthTypeAPI)
	}
	if !AuthParamWellKnownAuthTypeWellKnown.IsKnown() || string(AuthParamWellKnownAuthTypeWellKnown) != "wellknown" {
		t.Errorf("AuthParamWellKnownAuthTypeWellKnown = %q", AuthParamWellKnownAuthTypeWellKnown)
	}
	if AuthParamOAuthType("bogus").IsKnown() || AuthParamAPIAuthType("bogus").IsKnown() || AuthParamWellKnownAuthType("bogus").IsKnown() {
		t.Error("IsKnown() must reject unknown values")
	}
}

// End-to-end guarantee: the bytes actually placed on the wire by
// AuthService.Set must contain only the fields the caller set.
func TestAuthSetRequestBodyOnTheWire(t *testing.T) {
	cases := []struct {
		name string
		body AuthParam
		want string
	}{
		{
			name: "oauth required only",
			body: AuthParamOAuth{
				Type:    F(AuthParamOAuthTypeOAuth),
				Refresh: F("r"),
				Access:  F("a"),
				Expires: F(int64(3600)),
			},
			want: `{"access":"a","expires":3600,"refresh":"r","type":"oauth"}`,
		},
		{
			name: "api without metadata",
			body: AuthParamAPIAuth{Type: F(AuthParamAPIAuthTypeAPI), Key: F("sk-123")},
			want: `{"key":"sk-123","type":"api"}`,
		},
		{
			name: "wellknown",
			body: AuthParamWellKnownAuth{Type: F(AuthParamWellKnownAuthTypeWellKnown), Key: F("k"), Token: F("t")},
			want: `{"key":"k","token":"t","type":"wellknown"}`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotBody, gotPath, gotMethod string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				raw, _ := io.ReadAll(r.Body)
				gotBody, gotPath, gotMethod = string(raw), r.URL.Path, r.Method
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`true`))
			}))
			defer srv.Close()

			client := NewClient(option.WithBaseURL(srv.URL))
			res, err := client.Auth.Set(context.Background(), "anthropic", c.body)
			if err != nil {
				t.Fatalf("Set: %v", err)
			}
			if res == nil || !*res {
				t.Errorf("res = %v, want true", res)
			}
			if gotMethod != http.MethodPut {
				t.Errorf("method = %s, want PUT", gotMethod)
			}
			if gotPath != "/auth/anthropic" {
				t.Errorf("path = %s, want /auth/anthropic", gotPath)
			}
			if gotBody != c.want {
				t.Errorf("body\n got  %s\n want %s", gotBody, c.want)
			}
		})
	}
}
