// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// V2LocationService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2LocationService] method instead.
type V2LocationService struct {
	Options []option.RequestOption
}

// NewV2LocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2LocationService(opts ...option.RequestOption) (r *V2LocationService) {
	r = &V2LocationService{}
	r.Options = opts
	return
}

// Get location
//
// Resolve the requested location or the server default location.
func (r *V2LocationService) Get(ctx context.Context, query V2LocationGetParams, opts ...option.RequestOption) (res *LocationInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/location"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// V2LocationParam contains the query parameters for location.
type V2LocationParam struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// V2LocationGetParams contains the query parameters for getting location.
type V2LocationGetParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2LocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
