// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// V2CredentialService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2CredentialService] method instead.
type V2CredentialService struct {
	Options []option.RequestOption
}

// NewV2CredentialService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2CredentialService(opts ...option.RequestOption) (r *V2CredentialService) {
	r = &V2CredentialService{}
	r.Options = opts
	return
}

// Update credential label
func (r *V2CredentialService) Update(ctx context.Context, credentialID string, body V2CredentialUpdateParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if credentialID == "" {
		err = errors.New("missing required credentialID parameter")
		return
	}
	path := fmt.Sprintf("api/credential/%s", credentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Remove credential
func (r *V2CredentialService) Remove(ctx context.Context, credentialID string, query V2CredentialRemoveParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if credentialID == "" {
		err = errors.New("missing required credentialID parameter")
		return
	}
	path := fmt.Sprintf("api/credential/%s", credentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, &res, opts...)
	return
}

type V2CredentialUpdateParams struct {
	Label    param.Field[string]          `json:"label,required"`
	Location param.Field[V2LocationParam] `query:"location"`
}

// MarshalJSON serializes [V2CredentialUpdateParams] omitting query parameters.
func (r V2CredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V2CredentialUpdateParams]'s query parameters as `url.Values`.
func (r V2CredentialUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2CredentialRemoveParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

// URLQuery serializes [V2CredentialRemoveParams]'s query parameters as `url.Values`.
func (r V2CredentialRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
