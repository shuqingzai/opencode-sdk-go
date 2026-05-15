// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// AuthService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r *AuthService) {
	r = &AuthService{}
	r.Options = opts
	return
}

// Set authentication credentials for a provider
func (r *AuthService) Set(ctx context.Context, providerID string, body Auth, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if providerID == "" {
		err = fmt.Errorf("missing required providerID parameter")
		return
	}
	path := fmt.Sprintf("auth/%s", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove authentication credentials for a provider
func (r *AuthService) Remove(ctx context.Context, providerID string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if providerID == "" {
		err = fmt.Errorf("missing required providerID parameter")
		return
	}
	path := fmt.Sprintf("auth/%s", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Union satisfied by [OAuth], [ApiAuth], or [WellKnownAuth].
type Auth interface {
	implementsAuth()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Auth)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			DiscriminatorValue: "oauth",
			Type:               reflect.TypeOf(OAuth{}),
		},
		apijson.UnionVariant{
			DiscriminatorValue: "api",
			Type:               reflect.TypeOf(ApiAuth{}),
		},
		apijson.UnionVariant{
			DiscriminatorValue: "wellknown",
			Type:               reflect.TypeOf(WellKnownAuth{}),
		},
	)
}

type OAuth struct {
	Type          string    `json:"type,required"`
	Refresh       string    `json:"refresh,required"`
	Access        string    `json:"access,required"`
	Expires       int64     `json:"expires,required"`
	AccountID     string    `json:"accountId"`
	EnterpriseURL string    `json:"enterpriseUrl"`
	JSON          oauthJSON `json:"-"`
}

// oauthJSON contains the JSON metadata for the struct [OAuth]
type oauthJSON struct {
	Type          apijson.Field
	Refresh       apijson.Field
	Access        apijson.Field
	Expires       apijson.Field
	AccountID     apijson.Field
	EnterpriseURL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oauthJSON) RawJSON() string {
	return r.raw
}

func (r OAuth) implementsAuth() {}

type ApiAuth struct {
	Type     string            `json:"type,required"`
	Key      string            `json:"key,required"`
	Metadata map[string]string `json:"metadata,required"`
	JSON     apiAuthJSON       `json:"-"`
}

// apiAuthJSON contains the JSON metadata for the struct [ApiAuth]
type apiAuthJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApiAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiAuthJSON) RawJSON() string {
	return r.raw
}

func (r ApiAuth) implementsAuth() {}

type WellKnownAuth struct {
	Type  string            `json:"type,required"`
	Key   string            `json:"key,required"`
	Token string            `json:"token,required"`
	JSON  wellKnownAuthJSON `json:"-"`
}

// wellKnownAuthJSON contains the JSON metadata for the struct [WellKnownAuth]
type wellKnownAuthJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WellKnownAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wellKnownAuthJSON) RawJSON() string {
	return r.raw
}

func (r WellKnownAuth) implementsAuth() {}
