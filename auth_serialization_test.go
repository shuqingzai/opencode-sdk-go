package opencode

import (
	"encoding/json"
	"testing"

	"github.com/sst/opencode-sdk-go/internal/apijson"
)

// Aligned with OpenAPI schema "Auth" (anyOf: OAuth | ApiAuth | WellKnownAuth)
// used as the request body for auth.set (PUT /auth/{providerID}), and JS SDK(v2).

func TestAuthBodySerialization(t *testing.T) {
	t.Run("OAuth variant — full fields", func(t *testing.T) {
		// OpenAPI OAuth required: type, refresh, access, expires; optional accountId, enterpriseUrl
		a := OAuth{
			Type:          "oauth",
			Refresh:       "r",
			Access:        "a",
			Expires:       3600,
			AccountID:     "acc",
			EnterpriseURL: "https://ghe.example.com",
		}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"type":"oauth","refresh":"r","access":"a","expires":3600,"accountId":"acc","enterpriseUrl":"https://ghe.example.com"}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("OAuth variant — required only (optional fields emitted empty)", func(t *testing.T) {
		// Auth is a shared discriminated-union struct serialized via encoding/json,
		// so optional string fields are emitted as empty strings (no omitempty).
		a := OAuth{Type: "oauth", Refresh: "r", Access: "a", Expires: 0}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"type":"oauth","refresh":"r","access":"a","expires":0,"accountId":"","enterpriseUrl":""}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("ApiAuth variant with metadata", func(t *testing.T) {
		a := ApiAuth{Type: "api", Key: "sk-123", Metadata: map[string]string{"region": "us"}}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"type":"api","key":"sk-123","metadata":{"region":"us"}}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("ApiAuth variant without optional metadata (emitted null)", func(t *testing.T) {
		// metadata is an optional map without omitempty -> serialized as null.
		a := ApiAuth{Type: "api", Key: "sk-123"}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"type":"api","key":"sk-123","metadata":null}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("WellKnownAuth variant", func(t *testing.T) {
		a := WellKnownAuth{Type: "wellknown", Key: "k", Token: "t"}
		b, err := json.Marshal(a)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"type":"wellknown","key":"k","token":"t"}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Auth interface accepts all three variants", func(t *testing.T) {
		var _ Auth = OAuth{Type: "oauth"}
		var _ Auth = ApiAuth{Type: "api"}
		var _ Auth = WellKnownAuth{Type: "wellknown"}
	})
}

// Verifies the discriminated union registration ("type") dispatches to the
// correct concrete variant on unmarshal, per apijson.RegisterUnion in auth.go.
func TestAuthUnionDiscriminatorDispatch(t *testing.T) {
	t.Run("oauth discriminator", func(t *testing.T) {
		var a Auth
		raw := []byte(`{"type":"oauth","refresh":"r","access":"a","expires":100}`)
		if err := apijson.Unmarshal(raw, &a); err != nil {
			t.Fatal(err)
		}
		o, ok := a.(OAuth)
		if !ok {
			t.Fatalf("expected OAuth, got %T", a)
		}
		if o.Refresh != "r" || o.Access != "a" || o.Expires != 100 {
			t.Errorf("OAuth = %+v", o)
		}
	})

	t.Run("api discriminator", func(t *testing.T) {
		var a Auth
		raw := []byte(`{"type":"api","key":"k","metadata":{"x":"y"}}`)
		if err := apijson.Unmarshal(raw, &a); err != nil {
			t.Fatal(err)
		}
		ap, ok := a.(ApiAuth)
		if !ok {
			t.Fatalf("expected ApiAuth, got %T", a)
		}
		if ap.Key != "k" || ap.Metadata["x"] != "y" {
			t.Errorf("ApiAuth = %+v", ap)
		}
	})

	t.Run("wellknown discriminator", func(t *testing.T) {
		var a Auth
		raw := []byte(`{"type":"wellknown","key":"k","token":"t"}`)
		if err := apijson.Unmarshal(raw, &a); err != nil {
			t.Fatal(err)
		}
		w, ok := a.(WellKnownAuth)
		if !ok {
			t.Fatalf("expected WellKnownAuth, got %T", a)
		}
		if w.Key != "k" || w.Token != "t" {
			t.Errorf("WellKnownAuth = %+v", w)
		}
	})
}

// Variant Response-side deserialization with JSON metadata / RawJSON.
func TestAuthVariantDeserialization(t *testing.T) {
	t.Run("OAuth expires is int64 (integer semantics)", func(t *testing.T) {
		raw := `{"type":"oauth","refresh":"r","access":"a","expires":9999999999}`
		var o OAuth
		if err := json.Unmarshal([]byte(raw), &o); err != nil {
			t.Fatal(err)
		}
		if o.Expires != 9999999999 {
			t.Errorf("Expires = %d, want 9999999999", o.Expires)
		}
		if o.JSON.RawJSON() != raw {
			t.Errorf("RawJSON mismatch: %s", o.JSON.RawJSON())
		}
	})

	t.Run("ApiAuth metadata map[string]string", func(t *testing.T) {
		raw := `{"type":"api","key":"k","metadata":{"a":"1","b":"2"}}`
		var ap ApiAuth
		if err := json.Unmarshal([]byte(raw), &ap); err != nil {
			t.Fatal(err)
		}
		if ap.Metadata["a"] != "1" || ap.Metadata["b"] != "2" {
			t.Errorf("Metadata = %+v", ap.Metadata)
		}
	})
}
