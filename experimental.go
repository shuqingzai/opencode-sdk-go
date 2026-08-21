// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"

	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/pagination"
)

// ExperimentalService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalService] method instead.
type ExperimentalService struct {
	Options      []option.RequestOption
	Capabilities *ExperimentalCapabilitiesService
	Console      *ExperimentalConsoleService
	ControlPlane *ExperimentalControlPlaneService
	ProjectCopy  *ExperimentalProjectCopyService
	Resource     *ExperimentalResourceService
	Session      *ExperimentalSessionService
	Workspace    *ExperimentalWorkspaceService
}

// NewExperimentalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExperimentalService(opts ...option.RequestOption) (r *ExperimentalService) {
	r = &ExperimentalService{}
	r.Options = opts
	r.Capabilities = NewExperimentalCapabilitiesService(opts...)
	r.Console = NewExperimentalConsoleService(opts...)
	r.ControlPlane = NewExperimentalControlPlaneService(opts...)
	r.ProjectCopy = NewExperimentalProjectCopyService(opts...)
	r.Resource = NewExperimentalResourceService(opts...)
	r.Session = NewExperimentalSessionService(opts...)
	r.Workspace = NewExperimentalWorkspaceService(opts...)
	return
}

// WorkspaceList lists workspaces in the experimental API.
//
// Deprecated: Use [ExperimentalService.Workspace.List] instead.
func (r *ExperimentalService) WorkspaceList(ctx context.Context, query ExperimentalWorkspaceListParams, opts ...option.RequestOption) (res *[]Workspace, err error) {
	return r.Workspace.List(ctx, query, opts...)
}

// WorkspaceNew creates a workspace in the experimental API.
//
// Deprecated: Use [ExperimentalService.Workspace.New] instead.
func (r *ExperimentalService) WorkspaceNew(ctx context.Context, body ExperimentalWorkspaceNewParams, opts ...option.RequestOption) (res *Workspace, err error) {
	return r.Workspace.New(ctx, body, opts...)
}

// WorkspaceRemove removes a workspace in the experimental API.
//
// Deprecated: Use [ExperimentalService.Workspace.Remove] instead.
func (r *ExperimentalService) WorkspaceRemove(ctx context.Context, id string, query ExperimentalWorkspaceRemoveParams, opts ...option.RequestOption) (res *Workspace, err error) {
	return r.Workspace.Remove(ctx, id, query, opts...)
}

// AdapterList lists workspace adapters in the experimental API.
//
// Deprecated: Use [ExperimentalService.Workspace.Adapter.List] instead.
func (r *ExperimentalService) AdapterList(ctx context.Context, query ExperimentalAdapterListParams, opts ...option.RequestOption) (res *[]WorkspaceAdapter, err error) {
	return r.Workspace.Adapter.List(ctx, query, opts...)
}

// WorkspaceAdapterList lists workspace adapters in the experimental API.
//
// Deprecated: Use [ExperimentalService.Workspace.Adapter.List] instead.
func (r *ExperimentalService) WorkspaceAdapterList(ctx context.Context, query ExperimentalAdapterListParams, opts ...option.RequestOption) (res *[]WorkspaceAdapter, err error) {
	return r.Workspace.Adapter.List(ctx, query, opts...)
}

// WorkspaceStatus gets the workspace status in the experimental API.
//
// Deprecated: Use [ExperimentalService.Workspace.Status] instead.
func (r *ExperimentalService) WorkspaceStatus(ctx context.Context, query ExperimentalWorkspaceStatusParams, opts ...option.RequestOption) (res *[]WorkspaceEventConnectionStatus, err error) {
	return r.Workspace.Status(ctx, query, opts...)
}

// Warp warps a workspace in the experimental API.
//
// Deprecated: Use [ExperimentalService.Workspace.Warp] instead.
func (r *ExperimentalService) Warp(ctx context.Context, params ExperimentalWarpParams, opts ...option.RequestOption) (err error) {
	return r.Workspace.Warp(ctx, params, opts...)
}

// WorkspaceWarp warps a workspace in the experimental API.
//
// Deprecated: Use [ExperimentalService.Workspace.Warp] instead.
func (r *ExperimentalService) WorkspaceWarp(ctx context.Context, params ExperimentalWarpParams, opts ...option.RequestOption) (err error) {
	return r.Workspace.Warp(ctx, params, opts...)
}

// WorkspaceSyncList registers missing workspaces returned by workspace adapters in the experimental API.
//
// Deprecated: Use [ExperimentalService.Workspace.SyncList] instead.
func (r *ExperimentalService) WorkspaceSyncList(ctx context.Context, query ExperimentalWorkspaceSyncListParams, opts ...option.RequestOption) (err error) {
	return r.Workspace.SyncList(ctx, query, opts...)
}

// ConsoleGet gets the active Console provider metadata in the experimental API.
//
// Deprecated: Use [ExperimentalService.Console.Get] instead.
func (r *ExperimentalService) ConsoleGet(ctx context.Context, query ExperimentalConsoleGetParams, opts ...option.RequestOption) (res *ConsoleState, err error) {
	return r.Console.Get(ctx, query, opts...)
}

// ConsoleListOrgs lists switchable Console orgs in the experimental API.
//
// Deprecated: Use [ExperimentalService.Console.ListOrgs] instead.
func (r *ExperimentalService) ConsoleListOrgs(ctx context.Context, query ExperimentalConsoleListOrgsParams, opts ...option.RequestOption) (res *ConsoleListOrgsResponse, err error) {
	return r.Console.ListOrgs(ctx, query, opts...)
}

// ConsoleSwitchOrg switches the active Console org in the experimental API.
//
// Deprecated: Use [ExperimentalService.Console.SwitchOrg] instead.
func (r *ExperimentalService) ConsoleSwitchOrg(ctx context.Context, body ExperimentalConsoleSwitchOrgParams, opts ...option.RequestOption) (res *bool, err error) {
	return r.Console.SwitchOrg(ctx, body, opts...)
}

// SessionList lists sessions across projects in the experimental API.
//
// Deprecated: Use [ExperimentalService.Session.List] instead.
func (r *ExperimentalService) SessionList(ctx context.Context, query ExperimentalSessionListParams, opts ...option.RequestOption) (res *pagination.HeaderCursorPage[GlobalSession], err error) {
	return r.Session.List(ctx, query, opts...)
}

// CapabilitiesGet gets the experimental capabilities in the experimental API.
//
// Deprecated: Use [ExperimentalService.Capabilities.Get] instead.
func (r *ExperimentalService) CapabilitiesGet(ctx context.Context, query ExperimentalCapabilitiesGetParams, opts ...option.RequestOption) (res *ExperimentalCapabilities, err error) {
	return r.Capabilities.Get(ctx, query, opts...)
}

// ControlPlaneMoveSession moves a session in the experimental API.
//
// Deprecated: Use [ExperimentalService.ControlPlane.MoveSession] instead.
func (r *ExperimentalService) ControlPlaneMoveSession(ctx context.Context, body ExperimentalControlPlaneMoveSessionParams, opts ...option.RequestOption) (err error) {
	return r.ControlPlane.MoveSession(ctx, body, opts...)
}

// ProjectCopyGenerateName generates a project copy name in the experimental API.
//
// Deprecated: Use [ExperimentalService.ProjectCopy.GenerateName] instead.
func (r *ExperimentalService) ProjectCopyGenerateName(ctx context.Context, projectID string, params ExperimentalProjectCopyGenerateNameParams, opts ...option.RequestOption) (res *ProjectCopyGenerateNameResponse, err error) {
	return r.ProjectCopy.GenerateName(ctx, projectID, params, opts...)
}

// SessionBackground detaches any synchronous subagents currently blocking the session and continues them in the background.
//
// Deprecated: Use [ExperimentalService.Session.Background] instead.
func (r *ExperimentalService) SessionBackground(ctx context.Context, sessionID string, query ExperimentalSessionBackgroundParams, opts ...option.RequestOption) (res *bool, err error) {
	return r.Session.Background(ctx, sessionID, query, opts...)
}

// ResourceList lists MCP resources in the experimental API.
//
// Deprecated: Use [ExperimentalService.Resource.List] instead.
func (r *ExperimentalService) ResourceList(ctx context.Context, query ExperimentalResourceListParams, opts ...option.RequestOption) (res *map[string]McpResource, err error) {
	return r.Resource.List(ctx, query, opts...)
}
