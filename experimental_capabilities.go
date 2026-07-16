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

// ExperimentalCapabilitiesService contains methods and other services that help
// with interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalCapabilitiesService] method instead.
type ExperimentalCapabilitiesService struct {
	Options []option.RequestOption
}

// NewExperimentalCapabilitiesService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExperimentalCapabilitiesService(opts ...option.RequestOption) (r *ExperimentalCapabilitiesService) {
	r = &ExperimentalCapabilitiesService{}
	r.Options = opts
	return
}

// Get experimental capabilities
//
// Get experimental features enabled on the OpenCode server.
func (r *ExperimentalCapabilitiesService) Get(ctx context.Context, query ExperimentalCapabilitiesGetParams, opts ...option.RequestOption) (res *ExperimentalCapabilities, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/capabilities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// ExperimentalCapabilities represents the experimental capabilities response.
type ExperimentalCapabilities struct {
	BackgroundSubagents bool                         `json:"backgroundSubagents,required"`
	JSON                experimentalCapabilitiesJSON `json:"-"`
}

// experimentalCapabilitiesJSON contains the JSON metadata for the struct [ExperimentalCapabilities]
type experimentalCapabilitiesJSON struct {
	BackgroundSubagents apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ExperimentalCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r experimentalCapabilitiesJSON) RawJSON() string {
	return r.raw
}

type ExperimentalCapabilitiesGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ExperimentalCapabilitiesGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
