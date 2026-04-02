// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// GlobalService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalService] method instead.
type GlobalService struct {
	Options []option.RequestOption
}

// NewGlobalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGlobalService(opts ...option.RequestOption) (r *GlobalService) {
	r = &GlobalService{}
	r.Options = opts
	return
}

// Get global health status
func (r *GlobalService) Health(ctx context.Context, opts ...option.RequestOption) (res *GlobalHealthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Dispose global instance
func (r *GlobalService) Dispose(ctx context.Context, query GlobalDisposeParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/dispose"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

// Upgrade global instance
func (r *GlobalService) Upgrade(ctx context.Context, body GlobalUpgradeBody, opts ...option.RequestOption) (res *GlobalUpgradeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/upgrade"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GlobalHealthResponse struct {
	Healthy bool                     `json:"healthy,required"`
	Version string                   `json:"version,required"`
	JSON    globalHealthResponseJSON `json:"-"`
}

// globalHealthResponseJSON contains the JSON metadata for the struct [GlobalHealthResponse]
type globalHealthResponseJSON struct {
	Healthy     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalHealthResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalHealthResponseJSON) RawJSON() string {
	return r.raw
}

type GlobalDisposeParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [GlobalDisposeParams]'s query parameters as `url.Values`.
func (r GlobalDisposeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalUpgradeBody struct {
	Target param.Field[string] `json:"target,omitempty"`
}

func (r GlobalUpgradeBody) MarshalJSON() ([]byte, error) {
	return apijson.MarshalRoot(r)
}

type GlobalUpgradeResponse struct {
	Success bool                      `json:"success,required"`
	Version string                    `json:"version,omitempty"`
	Error   string                    `json:"error,omitempty"`
	JSON    globalUpgradeResponseJSON `json:"-"`
}

// globalUpgradeResponseJSON contains the JSON metadata for the struct [GlobalUpgradeResponse]
type globalUpgradeResponseJSON struct {
	Success     apijson.Field
	Version     apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalUpgradeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalUpgradeResponseJSON) RawJSON() string {
	return r.raw
}
