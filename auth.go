// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/tidwall/gjson"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/param"
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

// Union satisfied by [AuthParamOAuth], [AuthParamAPIAuth], or [AuthParamWellKnownAuth].
type AuthParam interface {
	implementsAuthParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[AuthParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[AuthParamOAuth](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[AuthParamAPIAuth](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[AuthParamWellKnownAuth](),
		},
	)
}

type AuthParamOAuth struct {
	Type          param.Field[AuthParamOAuthType] `json:"type,required"`
	Refresh       param.Field[string]             `json:"refresh,required"`
	Access        param.Field[string]             `json:"access,required"`
	Expires       param.Field[int64]              `json:"expires,required"`
	AccountID     param.Field[string]             `json:"accountId"`
	EnterpriseURL param.Field[string]             `json:"enterpriseUrl"`
}

func (r AuthParamOAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthParamOAuth) implementsAuthParam() {}

type AuthParamOAuthType string

const (
	AuthParamOAuthTypeOAuth AuthParamOAuthType = "oauth"
)

func (r AuthParamOAuthType) IsKnown() bool {
	switch r {
	case AuthParamOAuthTypeOAuth:
		return true
	}
	return false
}

type AuthParamAPIAuth struct {
	Type     param.Field[AuthParamAPIAuthType] `json:"type,required"`
	Key      param.Field[string]               `json:"key,required"`
	Metadata param.Field[map[string]string]    `json:"metadata"`
}

func (r AuthParamAPIAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthParamAPIAuth) implementsAuthParam() {}

type AuthParamAPIAuthType string

const (
	AuthParamAPIAuthTypeAPI AuthParamAPIAuthType = "api"
)

func (r AuthParamAPIAuthType) IsKnown() bool {
	switch r {
	case AuthParamAPIAuthTypeAPI:
		return true
	}
	return false
}

type AuthParamWellKnownAuth struct {
	Type  param.Field[AuthParamWellKnownAuthType] `json:"type,required"`
	Key   param.Field[string]                     `json:"key,required"`
	Token param.Field[string]                     `json:"token,required"`
}

func (r AuthParamWellKnownAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AuthParamWellKnownAuth) implementsAuthParam() {}

type AuthParamWellKnownAuthType string

const (
	AuthParamWellKnownAuthTypeWellKnown AuthParamWellKnownAuthType = "wellknown"
)

func (r AuthParamWellKnownAuthType) IsKnown() bool {
	switch r {
	case AuthParamWellKnownAuthTypeWellKnown:
		return true
	}
	return false
}
