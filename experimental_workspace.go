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

// ExperimentalWorkspaceService contains methods and other services that help
// with interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalWorkspaceService] method instead.
type ExperimentalWorkspaceService struct {
	Options []option.RequestOption
	Adapter *ExperimentalWorkspaceAdapterService
}

// NewExperimentalWorkspaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExperimentalWorkspaceService(opts ...option.RequestOption) (r *ExperimentalWorkspaceService) {
	r = &ExperimentalWorkspaceService{}
	r.Options = opts
	r.Adapter = NewExperimentalWorkspaceAdapterService(opts...)
	return
}

// List workspaces
func (r *ExperimentalWorkspaceService) List(ctx context.Context, query ExperimentalWorkspaceListParams, opts ...option.RequestOption) (res *[]Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// New workspace
//
// Create a workspace.
func (r *ExperimentalWorkspaceService) New(ctx context.Context, body ExperimentalWorkspaceNewParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove a workspace
func (r *ExperimentalWorkspaceService) Remove(ctx context.Context, id string, query ExperimentalWorkspaceRemoveParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("experimental/workspace/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, &res, opts...)
	return
}

// Get workspace status
func (r *ExperimentalWorkspaceService) Status(ctx context.Context, query ExperimentalWorkspaceStatusParams, opts ...option.RequestOption) (res *[]WorkspaceStatusItem, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Warp a workspace
func (r *ExperimentalWorkspaceService) Warp(ctx context.Context, params ExperimentalWarpParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/warp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Register missing workspaces returned by workspace adapters.
func (r *ExperimentalWorkspaceService) SyncList(ctx context.Context, query ExperimentalWorkspaceSyncListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/sync-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, nil, opts...)
	return
}

// ExperimentalWorkspaceAdapterService contains methods and other services that
// help with interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the
// environment automatically. You should not instantiate this service directly,
// and instead use the [NewExperimentalWorkspaceAdapterService] method instead.
type ExperimentalWorkspaceAdapterService struct {
	Options []option.RequestOption
}

// NewExperimentalWorkspaceAdapterService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExperimentalWorkspaceAdapterService(opts ...option.RequestOption) (r *ExperimentalWorkspaceAdapterService) {
	r = &ExperimentalWorkspaceAdapterService{}
	r.Options = opts
	return
}

// List workspace adapters
func (r *ExperimentalWorkspaceAdapterService) List(ctx context.Context, query ExperimentalAdapterListParams, opts ...option.RequestOption) (res *[]WorkspaceAdapter, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/adapter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Workspace struct {
	ID        string `json:"id,required"`
	Type      string `json:"type,required"`
	Name      string `json:"name,required"`
	Branch    string `json:"branch"`
	Directory string `json:"directory"`
	Extra     any    `json:"extra"`
	ProjectID string `json:"projectID,required"`
	// The amount of time in milliseconds that this workspace has been used.
	// This field can have the runtime type of [float64], [string] (one of "NaN",
	// "Infinity", "-Infinity").
	TimeUsed interface{}   `json:"timeUsed,required"`
	JSON     workspaceJSON `json:"-"`
}

// workspaceJSON contains the JSON metadata for the struct [Workspace]
type workspaceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Name        apijson.Field
	Branch      apijson.Field
	Directory   apijson.Field
	Extra       apijson.Field
	ProjectID   apijson.Field
	TimeUsed    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Workspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceJSON) RawJSON() string {
	return r.raw
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

// ExperimentalWorkspaceNewParams contains the request parameters for creating a workspace.
type ExperimentalWorkspaceNewParams struct {
	Directory param.Field[string]                `query:"directory"`
	Workspace param.Field[string]                `query:"workspace"`
	Body      ExperimentalWorkspaceCreateInput   `json:"-"`
	JSON      experimentalWorkspaceNewParamsJSON `json:"-"`
}

func (r ExperimentalWorkspaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ExperimentalWorkspaceNewParams]'s query parameters as `url.Values`.
func (r ExperimentalWorkspaceNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// experimentalWorkspaceNewParamsJSON contains the JSON metadata for the struct
// [ExperimentalWorkspaceNewParams]
type experimentalWorkspaceNewParamsJSON struct {
	Directory   apijson.Field
	Workspace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r experimentalWorkspaceNewParamsJSON) RawJSON() string {
	return r.raw
}

type ExperimentalWorkspaceCreateInput struct {
	ID     param.Field[string]                  `json:"id"`
	Type   param.Field[string]                  `json:"type,required"`
	Branch param.Field[string]                  `json:"branch"`
	Extra  param.Field[any]                     `json:"extra"`
	JSON   experimentalWorkspaceCreateInputJSON `json:"-"`
}

// experimentalWorkspaceCreateInputJSON contains the JSON metadata for the struct
// [ExperimentalWorkspaceCreateInput]
type experimentalWorkspaceCreateInputJSON struct {
	ID          apijson.Field
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

type WorkspaceAdapter struct {
	Type        string               `json:"type,required"`
	Name        string               `json:"name,required"`
	Description string               `json:"description,required"`
	JSON        workspaceAdapterJSON `json:"-"`
}

type workspaceAdapterJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceAdapter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceAdapterJSON) RawJSON() string {
	return r.raw
}

type ExperimentalAdapterListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ExperimentalAdapterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceStatusItem struct {
	WorkspaceID string                    `json:"workspaceID,required"`
	Status      WorkspaceStatusItemStatus `json:"status,required"`
	JSON        workspaceStatusItemJSON   `json:"-"`
}

type workspaceStatusItemJSON struct {
	WorkspaceID apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceStatusItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceStatusItemJSON) RawJSON() string {
	return r.raw
}

type WorkspaceStatusItemStatus string

const (
	WorkspaceStatusItemStatusConnected    WorkspaceStatusItemStatus = "connected"
	WorkspaceStatusItemStatusConnecting   WorkspaceStatusItemStatus = "connecting"
	WorkspaceStatusItemStatusDisconnected WorkspaceStatusItemStatus = "disconnected"
	WorkspaceStatusItemStatusError        WorkspaceStatusItemStatus = "error"
)

func (r WorkspaceStatusItemStatus) IsKnown() bool {
	switch r {
	case WorkspaceStatusItemStatusConnected, WorkspaceStatusItemStatusConnecting, WorkspaceStatusItemStatusDisconnected, WorkspaceStatusItemStatusError:
		return true
	}
	return false
}

type ExperimentalWorkspaceStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ExperimentalWorkspaceStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentalWarpParams struct {
	Directory   param.Field[string] `query:"directory"`
	Workspace   param.Field[string] `query:"workspace"`
	ID          param.Field[string] `json:"id,required"`
	SessionID   param.Field[string] `json:"sessionID,required"`
	CopyChanges param.Field[bool]   `json:"copyChanges"`
}

func (r ExperimentalWarpParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExperimentalWarpParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentalWorkspaceSyncListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ExperimentalWorkspaceSyncListParams]'s query parameters as `url.Values`.
func (r ExperimentalWorkspaceSyncListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
