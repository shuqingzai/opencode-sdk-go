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

// ToolService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolService] method instead.
type ToolService struct {
	Options []option.RequestOption
}

// NewToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolService(opts ...option.RequestOption) (r *ToolService) {
	r = &ToolService{}
	r.Options = opts
	return
}

// List tool IDs
func (r *ToolService) IDs(ctx context.Context, query ToolIDsParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/tool/ids"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List tool IDs
//
// Deprecated: Use [ToolService.IDs] instead.
func (r *ToolService) Ids(ctx context.Context, query ToolIDsParams, opts ...option.RequestOption) (res *[]string, err error) {
	return r.IDs(ctx, query, opts...)
}

// List tools
func (r *ToolService) List(ctx context.Context, query ToolListParams, opts ...option.RequestOption) (res *[]ToolListItem, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/tool"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ToolListItem struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required"`
	// This field is an untyped arbitrary value. The OpenAPI schema declares it as
	// an empty schema (`{}`), meaning it may hold any JSON value. Use a
	// type-switch or json.Unmarshal to inspect the runtime value.
	Parameters any              `json:"parameters,required"`
	JSON       toolListItemJSON `json:"-"`
}

// toolListItemJSON contains the JSON metadata for the struct [ToolListItem]
type toolListItemJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Parameters  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolListItemJSON) RawJSON() string {
	return r.raw
}

type ToolIDsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ToolIDsParams]'s query parameters as `url.Values`.
func (r ToolIDsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Deprecated: Use [ToolIDsParams] instead.
type ToolIdsParams = ToolIDsParams

type ToolListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Provider  param.Field[string] `query:"provider,required"`
	Model     param.Field[string] `query:"model,required"`
}

// URLQuery serializes [ToolListParams]'s query parameters as `url.Values`.
func (r ToolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
