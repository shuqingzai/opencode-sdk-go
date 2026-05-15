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

// List workspace adapters
func (r *ExperimentalService) AdapterList(ctx context.Context, query ExperimentalAdapterListParams, opts ...option.RequestOption) (res *[]WorkspaceAdapter, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/adapter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get workspace status
func (r *ExperimentalService) WorkspaceStatus(ctx context.Context, query ExperimentalWorkspaceStatusParams, opts ...option.RequestOption) (res *[]WorkspaceStatusItem, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Warp a workspace
func (r *ExperimentalService) Warp(ctx context.Context, params ExperimentalWarpParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/warp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Register missing workspaces returned by workspace adapters.
func (r *ExperimentalService) WorkspaceSyncList(ctx context.Context, query ExperimentalWorkspaceSyncListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/workspace/sync-list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, nil, opts...)
	return
}

// Get active Console provider metadata
func (r *ExperimentalService) ConsoleGet(ctx context.Context, query ExperimentalConsoleGetParams, opts ...option.RequestOption) (res *ConsoleState, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/console"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List switchable Console orgs
func (r *ExperimentalService) ConsoleListOrgs(ctx context.Context, query ExperimentalConsoleListOrgsParams, opts ...option.RequestOption) (res *ConsoleListOrgsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/console/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Switch active Console org
func (r *ExperimentalService) ConsoleSwitchOrg(ctx context.Context, body ConsoleSwitchOrgInput, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/console/switch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List sessions across projects
func (r *ExperimentalService) SessionList(ctx context.Context, query ExperimentalSessionListParams, opts ...option.RequestOption) (res *[]GlobalSession, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List MCP resources
func (r *ExperimentalService) ResourceList(ctx context.Context, query ExperimentalResourceListParams, opts ...option.RequestOption) (res *map[string]McpResource, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/resource"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	ID        param.Field[string] `json:"id"`
	SessionID param.Field[string] `json:"sessionID,required"`
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

type ConsoleState struct {
	ConsoleManagedProviders []string         `json:"consoleManagedProviders,required"`
	ActiveOrgName           string           `json:"activeOrgName,omitempty"`
	SwitchableOrgCount      float64          `json:"switchableOrgCount,required"`
	JSON                    consoleStateJSON `json:"-"`
}

type consoleStateJSON struct {
	ConsoleManagedProviders apijson.Field
	ActiveOrgName           apijson.Field
	SwitchableOrgCount      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ConsoleState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consoleStateJSON) RawJSON() string {
	return r.raw
}

type ExperimentalConsoleGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ExperimentalConsoleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConsoleOrg struct {
	AccountID    string         `json:"accountID,required"`
	AccountEmail string         `json:"accountEmail,required"`
	AccountUrl   string         `json:"accountUrl,required"`
	OrgID        string         `json:"orgID,required"`
	OrgName      string         `json:"orgName,required"`
	Active       bool           `json:"active,required"`
	JSON         consoleOrgJSON `json:"-"`
}

type consoleOrgJSON struct {
	AccountID    apijson.Field
	AccountEmail apijson.Field
	AccountUrl   apijson.Field
	OrgID        apijson.Field
	OrgName      apijson.Field
	Active       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConsoleOrg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consoleOrgJSON) RawJSON() string {
	return r.raw
}

type ConsoleListOrgsResponse struct {
	Orgs []ConsoleOrg                `json:"orgs,required"`
	JSON consoleListOrgsResponseJSON `json:"-"`
}

type consoleListOrgsResponseJSON struct {
	Orgs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsoleListOrgsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consoleListOrgsResponseJSON) RawJSON() string {
	return r.raw
}

type ExperimentalConsoleListOrgsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ExperimentalConsoleListOrgsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConsoleSwitchOrgInput struct {
	AccountID param.Field[string]       `json:"accountID,required"`
	OrgID     param.Field[string]       `json:"orgID,required"`
	JSON      consoleSwitchOrgInputJSON `json:"-"`
}

type consoleSwitchOrgInputJSON struct {
	AccountID   apijson.Field
	OrgID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsoleSwitchOrgInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ConsoleSwitchOrgInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r consoleSwitchOrgInputJSON) RawJSON() string {
	return r.raw
}

type ExperimentalSessionListParams struct {
	Directory param.Field[string]  `query:"directory"`
	Workspace param.Field[string]  `query:"workspace"`
	Roots     param.Field[bool]    `query:"roots"`
	Start     param.Field[float64] `query:"start"`
	Cursor    param.Field[float64] `query:"cursor"`
	Search    param.Field[string]  `query:"search"`
	Limit     param.Field[float64] `query:"limit"`
	Archived  param.Field[bool]    `query:"archived"`
}

func (r ExperimentalSessionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalSession struct {
	Id          string                `json:"id,required"`
	Slug        string                `json:"slug,omitempty"`
	ProjectID   string                `json:"projectID,omitempty"`
	WorkspaceID string                `json:"workspaceID,omitempty"`
	Directory   string                `json:"directory,omitempty"`
	ParentID    string                `json:"parentID,omitempty"`
	Summary     *GlobalSessionSummary `json:"summary,omitempty"`
	JSON        globalSessionJSON     `json:"-"`
}

type globalSessionJSON struct {
	Id          apijson.Field
	Slug        apijson.Field
	ProjectID   apijson.Field
	WorkspaceID apijson.Field
	Directory   apijson.Field
	ParentID    apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionJSON) RawJSON() string {
	return r.raw
}

type GlobalSessionSummary struct {
	Additions float64                  `json:"additions,required"`
	Deletions float64                  `json:"deletions,required"`
	Files     float64                  `json:"files,required"`
	JSON      globalSessionSummaryJSON `json:"-"`
}

type globalSessionSummaryJSON struct {
	Additions   apijson.Field
	Deletions   apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSessionSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionSummaryJSON) RawJSON() string {
	return r.raw
}

type McpResource struct {
	Name        string          `json:"name,required"`
	Uri         string          `json:"uri,required"`
	Description string          `json:"description,omitempty"`
	MimeType    string          `json:"mimeType,omitempty"`
	Client      string          `json:"client,required"`
	JSON        mcpResourceJSON `json:"-"`
}

type mcpResourceJSON struct {
	Name        apijson.Field
	Uri         apijson.Field
	Description apijson.Field
	MimeType    apijson.Field
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
