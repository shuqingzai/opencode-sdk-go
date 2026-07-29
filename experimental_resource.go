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

// ExperimentalResourceService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalResourceService] method instead.
type ExperimentalResourceService struct {
	Options []option.RequestOption
}

// NewExperimentalResourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExperimentalResourceService(opts ...option.RequestOption) (r *ExperimentalResourceService) {
	r = &ExperimentalResourceService{}
	r.Options = opts
	return
}

// List MCP resources
func (r *ExperimentalResourceService) List(ctx context.Context, query ExperimentalResourceListParams, opts ...option.RequestOption) (res *map[string]McpResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/resource"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type McpResource struct {
	Name        string          `json:"name,required"`
	URI         string          `json:"uri,required"`
	Description string          `json:"description"`
	MIMEType    string          `json:"mimeType"`
	Client      string          `json:"client,required"`
	JSON        mcpResourceJSON `json:"-"`
}

type mcpResourceJSON struct {
	Name        apijson.Field
	URI         apijson.Field
	Description apijson.Field
	MIMEType    apijson.Field
	Client      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpResourceJSON) RawJSON() string {
	return r.raw
}

type ExperimentalResourceListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ExperimentalResourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
