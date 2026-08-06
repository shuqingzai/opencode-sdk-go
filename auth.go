// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/tidwall/gjson"
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
func (r *AuthService) Set(ctx context.Context, providerID string, body AuthParam, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if providerID == "" {
		err = errors.New("missing required providerID parameter")
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
		err = errors.New("missing required providerID parameter")
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
		reflect.TypeFor[Auth](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "oauth",
			Type:               reflect.TypeFor[OAuth](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "api",
			Type:               reflect.TypeFor[ApiAuth](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "wellknown",
			Type:               reflect.TypeFor[WellKnownAuth](),
		},
	)
	apijson.RegisterUnion(
		reflect.TypeFor[AuthParam](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "oauth",
			Type:               reflect.TypeFor[OAuthParam](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "api",
			Type:               reflect.TypeFor[ApiAuthParam](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "wellknown",
			Type:               reflect.TypeFor[WellKnownAuthParam](),
		},
	)
}

type OAuth struct {
	Type          OAuthType `json:"type,required"`
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

type OAuthType string

const (
	OAuthTypeOAuth OAuthType = "oauth"
)

func (r OAuthType) IsKnown() bool {
	switch r {
	case OAuthTypeOAuth:
		return true
	}
	return false
}

type ApiAuth struct {
	Type     ApiAuthType       `json:"type,required"`
	Key      string            `json:"key,required"`
	Metadata map[string]string `json:"metadata"`
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

type ApiAuthType string

const (
	ApiAuthTypeAPI ApiAuthType = "api"
)

func (r ApiAuthType) IsKnown() bool {
	switch r {
	case ApiAuthTypeAPI:
		return true
	}
	return false
}

type WellKnownAuth struct {
	Type  WellKnownAuthType `json:"type,required"`
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

type WellKnownAuthType string

const (
	WellKnownAuthTypeWellKnown WellKnownAuthType = "wellknown"
)

func (r WellKnownAuthType) IsKnown() bool {
	switch r {
	case WellKnownAuthTypeWellKnown:
		return true
	}
	return false
}

// Satisfied by [OAuthParam], [ApiAuthParam], or [WellKnownAuthParam].
type AuthParam interface {
	implementsAuthParam()
}

// OAuthParam is the request-side counterpart of [OAuth] (OpenAPI OAuth schema).
type OAuthParam struct {
	Type          param.Field[OAuthParamType] `json:"type,required"`
	Refresh       param.Field[string]         `json:"refresh,required"`
	Access        param.Field[string]         `json:"access,required"`
	Expires       param.Field[int64]          `json:"expires,required"`
	AccountID     param.Field[string]         `json:"accountId"`
	EnterpriseURL param.Field[string]         `json:"enterpriseUrl"`
}

func (r OAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OAuthParam) implementsAuthParam() {}

type OAuthParamType string

const (
	OAuthParamTypeOAuth OAuthParamType = "oauth"
)

func (r OAuthParamType) IsKnown() bool {
	switch r {
	case OAuthParamTypeOAuth:
		return true
	}
	return false
}

// ApiAuthParam is the request-side counterpart of [ApiAuth] (OpenAPI ApiAuth schema).
type ApiAuthParam struct {
	Type     param.Field[ApiAuthParamType]  `json:"type,required"`
	Key      param.Field[string]            `json:"key,required"`
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r ApiAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ApiAuthParam) implementsAuthParam() {}

type ApiAuthParamType string

const (
	ApiAuthParamTypeAPI ApiAuthParamType = "api"
)

func (r ApiAuthParamType) IsKnown() bool {
	switch r {
	case ApiAuthParamTypeAPI:
		return true
	}
	return false
}

// WellKnownAuthParam is the request-side counterpart of [WellKnownAuth] (OpenAPI WellKnownAuth schema).
type WellKnownAuthParam struct {
	Type  param.Field[WellKnownAuthParamType] `json:"type,required"`
	Key   param.Field[string]                 `json:"key,required"`
	Token param.Field[string]                 `json:"token,required"`
}

func (r WellKnownAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WellKnownAuthParam) implementsAuthParam() {}

type WellKnownAuthParamType string

const (
	WellKnownAuthParamTypeWellKnown WellKnownAuthParamType = "wellknown"
)

func (r WellKnownAuthParamType) IsKnown() bool {
	switch r {
	case WellKnownAuthParamTypeWellKnown:
		return true
	}
	return false
}
