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

// ExperimentalService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalService] method instead.
type ExperimentalService struct {
	Options []option.RequestOption
}

// NewExperimentalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExperimentalService(opts ...option.RequestOption) (r *ExperimentalService) {
	r = &ExperimentalService{}
	r.Options = opts
	return
}

// Get tool IDs
func (r *ExperimentalService) ToolIds(ctx context.Context, query ExperimentalToolIdsParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/tool/ids"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List tools
func (r *ExperimentalService) ToolList(ctx context.Context, query ExperimentalToolListParams, opts ...option.RequestOption) (res *[]ToolListItem, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/tool"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List workspaces
func (r *ExperimentalService) WorkspaceList(ctx context.Context, query ExperimentalWorkspaceListParams, opts ...option.RequestOption) (res *[]Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create a workspace
func (r *ExperimentalService) WorkspaceCreate(ctx context.Context, body ExperimentalWorkspaceCreateInput, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove a workspace
func (r *ExperimentalService) WorkspaceRemove(ctx context.Context, id string, query ExperimentalWorkspaceRemoveParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("experimental/workspace/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, &res, opts...)
	return
}

type ToolListItem struct {
	Id          string           `json:"id,required"`
	Description string           `json:"description,required"`
	Parameters  any              `json:"parameters,required"`
	JSON        toolListItemJSON `json:"-"`
}

// toolListItemJSON contains the JSON metadata for the struct [ToolListItem]
type toolListItemJSON struct {
	Id          apijson.Field
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

type Workspace struct {
	Id        string        `json:"id,required"`
	Type      string        `json:"type,required"`
	Branch    string        `json:"branch"`
	Name      string        `json:"name"`
	Directory string        `json:"directory"`
	Extra     any           `json:"extra"`
	ProjectID string        `json:"projectID,required"`
	JSON      workspaceJSON `json:"-"`
}

// workspaceJSON contains the JSON metadata for the struct [Workspace]
type workspaceJSON struct {
	Id          apijson.Field
	Type        apijson.Field
	Branch      apijson.Field
	Name        apijson.Field
	Directory   apijson.Field
	Extra       apijson.Field
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Workspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceJSON) RawJSON() string {
	return r.raw
}

type ExperimentalToolIdsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ExperimentalToolIdsParams]'s query parameters as `url.Values`.
func (r ExperimentalToolIdsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentalToolListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Provider  param.Field[string] `query:"provider,required"`
	Model     param.Field[string] `query:"model,required"`
}

// URLQuery serializes [ExperimentalToolListParams]'s query parameters as `url.Values`.
func (r ExperimentalToolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentalWorkspaceListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ExperimentalWorkspaceListParams]'s query parameters as `url.Values`.
func (r ExperimentalWorkspaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentalWorkspaceCreateInput struct {
	Id     param.Field[string]                  `json:"id"`
	Type   param.Field[string]                  `json:"type"`
	Branch param.Field[string]                  `json:"branch"`
	Extra  param.Field[any]                     `json:"extra"`
	JSON   experimentalWorkspaceCreateInputJSON `json:"-"`
}

// experimentalWorkspaceCreateInputJSON contains the JSON metadata for the struct
// [ExperimentalWorkspaceCreateInput]
type experimentalWorkspaceCreateInputJSON struct {
	Id          apijson.Field
	Type        apijson.Field
	Branch      apijson.Field
	Extra       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExperimentalWorkspaceCreateInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ExperimentalWorkspaceCreateInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r experimentalWorkspaceCreateInputJSON) RawJSON() string {
	return r.raw
}

type ExperimentalWorkspaceCreateParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ExperimentalWorkspaceCreateParams]'s query parameters as `url.Values`.
func (r ExperimentalWorkspaceCreateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentalWorkspaceRemoveParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ExperimentalWorkspaceRemoveParams]'s query parameters as `url.Values`.
func (r ExperimentalWorkspaceRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
