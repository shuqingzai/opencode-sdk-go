// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"slices"

	"github.com/tidwall/gjson"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
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
func (r *GlobalService) Dispose(ctx context.Context, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/dispose"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Upgrade global instance
func (r *GlobalService) Upgrade(ctx context.Context, params GlobalUpgradeParams, opts ...option.RequestOption) (res *GlobalUpgradeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/upgrade"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Subscribe to global events via SSE
func (r *GlobalService) Event(ctx context.Context, opts ...option.RequestOption) (stream *ssestream.Stream[GlobalEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "global/event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[GlobalEvent](ssestream.NewDecoder(raw), err)
}

// Get global config
func (r *GlobalService) ConfigGet(ctx context.Context, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update global config
func (r *GlobalService) ConfigUpdate(ctx context.Context, params GlobalConfigUpdateParams, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
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

type GlobalUpgradeParams struct {
	Target param.Field[string] `json:"target"`
}

func (r GlobalUpgradeParams) MarshalJSON() ([]byte, error) {
	return apijson.MarshalRoot(r)
}

// Union satisfied by [GlobalUpgradeResponseSuccess] or [GlobalUpgradeResponseFailed].
type GlobalUpgradeResponseUnion interface {
	implementsGlobalUpgradeResponse()
}

type GlobalUpgradeResponseSuccess struct {
	Success bool                             `json:"success,required"`
	Version string                           `json:"version,required"`
	JSON    globalUpgradeResponseSuccessJSON `json:"-"`
}

// globalUpgradeResponseSuccessJSON contains the JSON metadata for the struct [GlobalUpgradeResponseSuccess]
type globalUpgradeResponseSuccessJSON struct {
	Success     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalUpgradeResponseSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalUpgradeResponseSuccessJSON) RawJSON() string {
	return r.raw
}

func (r GlobalUpgradeResponseSuccess) implementsGlobalUpgradeResponse() {}

type GlobalUpgradeResponseFailed struct {
	Success bool                            `json:"success,required"`
	Error   string                          `json:"error,required"`
	JSON    globalUpgradeResponseFailedJSON `json:"-"`
}

// globalUpgradeResponseFailedJSON contains the JSON metadata for the struct [GlobalUpgradeResponseFailed]
type globalUpgradeResponseFailedJSON struct {
	Success     apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalUpgradeResponseFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalUpgradeResponseFailedJSON) RawJSON() string {
	return r.raw
}

func (r GlobalUpgradeResponseFailed) implementsGlobalUpgradeResponse() {}

type GlobalUpgradeResponse struct {
	Success bool `json:"success,required"`
	// This field can have the runtime type of [string].
	Version any `json:"version"`
	// This field can have the runtime type of [string].
	Error any                       `json:"error"`
	JSON  globalUpgradeResponseJSON `json:"-"`
	union GlobalUpgradeResponseUnion
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
	*r = GlobalUpgradeResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r globalUpgradeResponseJSON) RawJSON() string {
	return r.raw
}

// AsUnion returns a [GlobalUpgradeResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [GlobalUpgradeResponseSuccess],
// [GlobalUpgradeResponseFailed].
func (r GlobalUpgradeResponse) AsUnion() GlobalUpgradeResponseUnion {
	return r.union
}

// GlobalConfigUpdateParams contains the parameters for updating global configuration.
// All fields are optional for PATCH semantics.
type GlobalConfigUpdateParams struct {
	// Body parameters — all Config fields as optional
	Schema     param.Field[string]                `json:"$schema"`
	Agent      param.Field[ConfigAgentParam]      `json:"agent"`
	Attachment param.Field[AttachmentConfigParam] `json:"attachment"`
	Autoshare  param.Field[bool]                  `json:"autoshare"`
	// Automatically update to the latest version. Pass true to auto-update,
	// false to disable, or "notify" to show update notifications.
	// Accepts [bool] or [string] ("notify").
	Autoupdate        param.Field[any]                           `json:"autoupdate"`
	Command           param.Field[map[string]ConfigCommandParam] `json:"command"`
	Compaction        param.Field[ConfigCompactionParam]         `json:"compaction"`
	DisabledProviders param.Field[[]string]                      `json:"disabled_providers"`
	EnabledProviders  param.Field[[]string]                      `json:"enabled_providers"`
	Enterprise        param.Field[EnterpriseConfigParam]         `json:"enterprise"`
	Experimental      param.Field[ConfigExperimentalParam]       `json:"experimental"`
	// Enable or configure formatters. Pass [shared.UnionBool](false) to disable,
	// [shared.UnionBool](true) to enable built-ins, or a [ConfigFormatterMapParam]
	// of formatter-name to [ConfigFormatterParam] to enable with overrides.
	Formatter    param.Field[ConfigFormatterSettingUnionParam] `json:"formatter"`
	Instructions param.Field[[]string]                         `json:"instructions"`
	Layout       param.Field[ConfigLayout]                     `json:"layout"`
	LogLevel     param.Field[ConfigLogLevel]                   `json:"logLevel"`
	// Enable or configure LSP servers. Pass [shared.UnionBool](false) to disable,
	// [shared.UnionBool](true) to enable built-ins, or a [ConfigLspMapParam] of
	// LSP-name to [ConfigLspUnionParam] to enable with overrides.
	Lsp param.Field[ConfigLspSettingUnionParam] `json:"lsp"`
	// Map of MCP server name → configuration. Each value is a
	// [ConfigMcpLocalParam], a [ConfigMcpRemoteParam] or a [ConfigMcpDisabledParam].
	Mcp   param.Field[map[string]ConfigMcpUnionParam] `json:"mcp"`
	Mode  param.Field[ConfigModeParam]                `json:"mode"`
	Model param.Field[string]                         `json:"model"`
	// Permission configuration. A short string ("ask"|"allow"|"deny") or an
	// object with per-action permission rule overrides. Accepts [ConfigPermissionAction]
	// (a string constant) or [ConfigPermissionParam].
	Permission param.Field[ConfigPermissionUnionParam] `json:"permission"`
	// Plugins to load. Each item is either a plugin name (string) or a 2-tuple
	// of [pluginName, configObject] (where configObject is a map[string]any).
	Plugin   param.Field[[]any]                          `json:"plugin"`
	Provider param.Field[map[string]ConfigProviderParam] `json:"provider"`
	// Map of reference name → value. Each value can be a plain [string] (URL/path),
	// a [ConfigV2ReferenceGitParam], or a [ConfigV2ReferenceLocalParam].
	Reference param.Field[map[string]ConfigV2ReferenceUnionParam] `json:"reference"`
	// Map of reference name → value. Each value can be a plain [string] (URL/path),
	// a [ConfigV2ReferenceGitParam], or a [ConfigV2ReferenceLocalParam].
	References    param.Field[map[string]ConfigV2ReferenceUnionParam] `json:"references"`
	Share         param.Field[ConfigShare]                            `json:"share"`
	Shell         param.Field[string]                                 `json:"shell"`
	Server        param.Field[ServerConfigParam]                      `json:"server"`
	Skills        param.Field[ConfigSkillsParam]                      `json:"skills"`
	SmallModel    param.Field[string]                                 `json:"small_model"`
	Snapshot      param.Field[bool]                                   `json:"snapshot"`
	ToolOutput    param.Field[ConfigToolOutputParam]                  `json:"tool_output"`
	Tools         param.Field[map[string]bool]                        `json:"tools"`
	Username      param.Field[string]                                 `json:"username"`
	Watcher       param.Field[ConfigWatcherParam]                     `json:"watcher"`
	DefaultAgent  param.Field[string]                                 `json:"default_agent"`
	SubagentDepth param.Field[int64]                                  `json:"subagent_depth"`
}

func (r GlobalConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GlobalEvent struct {
	Directory string `json:"directory,required"`
	// This field can have the runtime type of
	// [EventListResponseEventXxx] (all 89 OpenAPI `Event` variants):
	// [EventListResponseEventCommandExecuted],
	// [EventListResponseEventFileEdited],
	// [EventListResponseEventFileWatcherUpdated],
	// [EventListResponseEventGlobalDisposed],
	// [EventListResponseEventInstallationUpdateAvailable],
	// [EventListResponseEventInstallationUpdated],
	// [EventListResponseEventLspUpdated],
	// [EventListResponseEventMcpBrowserOpenFailed],
	// [EventListResponseEventMcpToolsChanged],
	// [EventListResponseEventModelsDevRefreshed],
	// [EventListResponseEventMessagePartDelta],
	// [EventListResponseEventMessagePartRemoved],
	// [EventListResponseEventMessagePartUpdated],
	// [EventListResponseEventMessageRemoved],
	// [EventListResponseEventMessageUpdated],
	// [EventListResponseEventPermissionAsked],
	// [EventListResponseEventPermissionReplied],
	// [EventListResponseEventPluginAdded],
	// [EventListResponseEventProjectUpdated],
	// [EventListResponseEventPtyCreated],
	// [EventListResponseEventPtyDeleted],
	// [EventListResponseEventPtyExited],
	// [EventListResponseEventPtyUpdated],
	// [EventListResponseEventQuestionAsked],
	// [EventListResponseEventQuestionRejected],
	// [EventListResponseEventQuestionReplied],
	// [EventListResponseEventServerConnected],
	// [EventListResponseEventServerInstanceDisposed],
	// [EventListResponseEventSessionCompacted],
	// [EventListResponseEventSessionCreated],
	// [EventListResponseEventSessionDeleted],
	// [EventListResponseEventSessionDiff],
	// [EventListResponseEventSessionError],
	// [EventListResponseEventSessionIdle],
	// [EventListResponseEventSessionNextAgentSwitched],
	// [EventListResponseEventSessionNextCompactionDelta],
	// [EventListResponseEventSessionNextCompactionEnded],
	// [EventListResponseEventSessionNextCompactionStarted],
	// [EventListResponseEventSessionNextModelSwitched],
	// [EventListResponseEventSessionNextPrompted],
	// [EventListResponseEventSessionNextReasoningDelta],
	// [EventListResponseEventSessionNextReasoningEnded],
	// [EventListResponseEventSessionNextReasoningStarted],
	// [EventListResponseEventSessionNextRetried],
	// [EventListResponseEventSessionNextShellEnded],
	// [EventListResponseEventSessionNextShellStarted],
	// [EventListResponseEventSessionNextStepEnded],
	// [EventListResponseEventSessionNextStepFailed],
	// [EventListResponseEventSessionNextStepStarted],
	// [EventListResponseEventSessionNextSynthetic],
	// [EventListResponseEventSessionNextTextDelta],
	// [EventListResponseEventSessionNextTextEnded],
	// [EventListResponseEventSessionNextTextStarted],
	// [EventListResponseEventSessionNextToolCalled],
	// [EventListResponseEventSessionNextToolFailed],
	// [EventListResponseEventSessionNextToolInputDelta],
	// [EventListResponseEventSessionNextToolInputEnded],
	// [EventListResponseEventSessionNextToolInputStarted],
	// [EventListResponseEventSessionNextToolProgress],
	// [EventListResponseEventSessionNextToolSuccess],
	// [EventListResponseEventSessionStatus],
	// [EventListResponseEventSessionUpdated],
	// [EventListResponseEventTodoUpdated],
	// [EventListResponseEventTuiCommandExecute],
	// [EventListResponseEventTuiPromptAppend],
	// [EventListResponseEventTuiSessionSelect],
	// [EventListResponseEventTuiToastShow],
	// [EventListResponseEventVcsBranchUpdated],
	// [EventListResponseEventWorkspaceFailed],
	// [EventListResponseEventWorkspaceReady],
	// [EventListResponseEventWorkspaceStatus],
	// [EventListResponseEventWorktreeFailed],
	// [EventListResponseEventWorktreeReady],
	//
	// [EventListResponseEventIntegrationUpdated],
	// [EventListResponseEventIntegrationConnectionUpdated],
	// [EventListResponseEventCatalogUpdated],
	// [EventListResponseEventPermissionV2Asked],
	// [EventListResponseEventPermissionV2Replied],
	// [EventListResponseEventReferenceUpdated],
	// [EventListResponseEventQuestionV2Asked],
	// [EventListResponseEventQuestionV2Replied],
	// [EventListResponseEventQuestionV2Rejected],
	// [EventListResponseEventSessionNextMoved],
	// [EventListResponseEventSessionNextRevertStaged],
	// [EventListResponseEventSessionNextRevertCleared],
	// [EventListResponseEventSessionNextRevertCommitted],
	// [EventListResponseEventSessionNextPromptAdmitted],
	// [EventListResponseEventSessionNextContextUpdated],
	// [EventListResponseEventProjectDirectoriesUpdated],
	//
	// [SyncEventResponse] (V1 SyncEvent).
	Payload   any             `json:"payload,required"`
	Project   string          `json:"project"`
	Workspace string          `json:"workspace"`
	JSON      globalEventJSON `json:"-"`
	union     GlobalEventPayloadUnion
}

type globalEventJSON struct {
	Directory   apijson.Field
	Payload     apijson.Field
	Project     apijson.Field
	Workspace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r globalEventJSON) RawJSON() string {
	return r.raw
}

func (r *GlobalEvent) UnmarshalJSON(data []byte) (err error) {
	*r = GlobalEvent{}

	// Phase 1: parse only the wrapper fields (directory, project, workspace)
	// into an alias struct that intentionally omits Payload, avoiding an
	// unnecessary pass through the any decoder. The full JSON metadata
	// is still captured correctly via globalEventJSON.
	type globalEventAlias struct {
		Directory string          `json:"directory,required"`
		Project   string          `json:"project"`
		Workspace string          `json:"workspace"`
		JSON      globalEventJSON `json:"-"`
	}
	var alias globalEventAlias
	if err := apijson.UnmarshalRoot(data, &alias); err != nil {
		return err
	}
	r.Directory = alias.Directory
	r.Project = alias.Project
	r.Workspace = alias.Workspace
	r.JSON = alias.JSON

	// Phase 2: extract the payload and route to the correct handler.
	//
	// Every variant of the OpenAPI `GlobalEvent.payload` anyOf is an object, so a
	// missing, null, or non-object payload cannot match any of them. Leave
	// [GlobalEvent.Payload] nil in that case instead of failing: this method is the
	// root decoder for the `/global/event` SSE stream, so a returned error is not
	// swallowed by a parent struct decoder the way it is for nested properties —
	// it surfaces through [ssestream.Stream.Err] and tears down the whole stream.
	// The wrapper fields and [GlobalEvent.RawJSON] are still available to callers.
	result := gjson.ParseBytes(data)
	payloadResult := result.Get("payload")
	if !payloadResult.Exists() || payloadResult.Type == gjson.Null || !payloadResult.IsObject() {
		return nil
	}

	// V1 SyncEvents: the server wraps them as {type:"sync", syncEvent:{...}}.
	// SyncEventResponse handles this directly via its own UnmarshalJSON.
	if payloadResult.Get("type").String() == "sync" {
		var resp SyncEventResponse
		if err := json.Unmarshal([]byte(payloadResult.Raw), &resp); err != nil {
			return err
		}
		r.union = resp
		r.Payload = resp
		return nil
	}

	// Regular events: payload has the shape {id, type, properties}, i.e. the
	// OpenAPI `Event` schema. Match directly against the union of
	// EventListResponseEventXxx types.
	if err := apijson.UnmarshalRoot([]byte(payloadResult.Raw), &r.union); err != nil {
		return err
	}
	r.Payload = r.union
	return nil
}

// AsUnion returns a [GlobalEventPayloadUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are all 89 OpenAPI `Event` variants
// (EventListResponseEvent*) and SyncEventResponse (for V1 SyncEvent types).
func (r GlobalEvent) AsUnion() GlobalEventPayloadUnion {
	return r.union
}

// GlobalEventPayloadUnion is a union of all 89 OpenAPI `Event` variants and
// SyncEventResponse that can appear in the GlobalEvent.payload field.
type GlobalEventPayloadUnion interface {
	implementsGlobalEventPayload()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[GlobalEventPayloadUnion](),
		"",
		// OpenAPI `Event` variants
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventInstallationUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventInstallationUpdateAvailable](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventLspUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventMessageUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventMessageRemoved](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventMessagePartUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventMessagePartDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventMessagePartRemoved](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionCompacted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventPermissionAsked](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventPermissionReplied](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventFileEdited](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventFileWatcherUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventTodoUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventTuiPromptAppend](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventTuiCommandExecute](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventTuiToastShow](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventTuiSessionSelect](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventMcpToolsChanged](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventModelsDevRefreshed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventMcpBrowserOpenFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventCommandExecuted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionIdle](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionStatus](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionDiff](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionCreated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionDeleted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventServerConnected](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventServerInstanceDisposed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventGlobalDisposed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventProjectUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventQuestionAsked](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventQuestionRejected](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventQuestionReplied](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventVcsBranchUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventWorkspaceReady](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventWorkspaceFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventWorkspaceStatus](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventPtyCreated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventPtyUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventPtyExited](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventPtyDeleted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventWorktreeReady](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventIntegrationUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventIntegrationConnectionUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventCatalogUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventPermissionV2Asked](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventPermissionV2Replied](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventReferenceUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventQuestionV2Asked](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventQuestionV2Replied](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventQuestionV2Rejected](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextMoved](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextRevertStaged](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextRevertCleared](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextRevertCommitted](),
		},
		// V1 SyncEvent types
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventWorktreeFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextAgentSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextModelSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextPrompted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextSynthetic](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextShellStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextShellEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextStepStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextStepEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextStepFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextTextStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextTextDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextTextEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextReasoningStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextReasoningDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextReasoningEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextToolInputStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextToolInputDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextToolInputEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextToolCalled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextToolProgress](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextToolSuccess](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextToolFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextRetried](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextCompactionStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextCompactionDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextCompactionEnded](),
		},
		// V1 SyncEvent types — wrapped in SyncEventResponse
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventResponse](),
		},
		// V2 Event: catalog.model.updated
		// V2 Event: plugin.added
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventPluginAdded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextPromptAdmitted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventSessionNextContextUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[EventListResponseEventProjectDirectoriesUpdated](),
		},
	)

	// SyncEventResponseSyncEventDataUnion: V1 SyncEvent parent types registered
	// as union variants. Each parent type carries a polymorphic Data field whose
	// concrete type reuses the V2 EventListResponseEventXxxProperties for the
	// same event (except session.updated which uses a custom Data type).
	// The parent type's enum Type field provides the discriminator for matching.
	apijson.RegisterUnion(
		reflect.TypeFor[SyncEventResponseSyncEventDataUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventMessageUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventMessageRemoved](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventMessagePartUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventMessagePartRemoved](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionCreated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionDeleted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextAgentSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextModelSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextPrompted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextPromptAdmitted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextSynthetic](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextShellStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextContextUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextShellEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextStepStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextStepEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextStepFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextTextStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextTextEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextReasoningStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextReasoningEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextToolInputStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextToolInputEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextToolCalled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextToolProgress](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextToolSuccess](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextToolFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextRetried](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextCompactionStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextCompactionEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextMoved](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextRevertStaged](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextRevertCleared](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SyncEventSessionNextRevertCommitted](),
		},
	)

	apijson.RegisterUnion(
		reflect.TypeFor[GlobalUpgradeResponseUnion](),
		// OpenAPI discriminates the two `/global/upgrade` response variants on
		// `success` (`{"type":"boolean","enum":[true]}` vs `enum:[false]`), which JS
		// SDK v2 mirrors as the literal types `success: true` / `success: false`.
		// `Success` is a plain `bool` on both variants, so it carries no `IsKnown`
		// enum for the exactness pass to discriminate on; without this key an
		// unrecognised extra property drops both variants to the same exactness and
		// the left-to-right tie-break always yields the `Success` variant, silently
		// losing `Error`.
		"success",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: true,
			Type:               reflect.TypeFor[GlobalUpgradeResponseSuccess](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: false,
			Type:               reflect.TypeFor[GlobalUpgradeResponseFailed](),
		},
	)
}

// SyncEventResponse wraps the server's V1 SyncEvent format:
// {type: "sync", syncEvent: {type: "message.updated.1", ...}, id: "..."}.
type SyncEventResponse struct {
	Type      SyncEventResponseType      `json:"type,required"`
	SyncEvent SyncEventResponseSyncEvent `json:"syncEvent,required"`
	ID        string                     `json:"id"`
	JSON      syncEventResponseJSON      `json:"-"`
}

type syncEventResponseJSON struct {
	Type        apijson.Field
	SyncEvent   apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventResponse) implementsGlobalEventPayload() {}

type SyncEventResponseType string

const (
	SyncEventResponseTypeSync SyncEventResponseType = "sync"
)

func (r SyncEventResponseType) IsKnown() bool {
	switch r {
	case SyncEventResponseTypeSync:
		return true
	}
	return false
}

// GlobalEventResponseType enumerates all possible event types in the
// [GlobalEvent.Payload] field. Each value corresponds to the Type field
// of the concrete variant stored in Payload.
type GlobalEventResponseType string

const (
	GlobalEventResponseTypeCommandExecuted              GlobalEventResponseType = "command.executed"
	GlobalEventResponseTypeFileEdited                   GlobalEventResponseType = "file.edited"
	GlobalEventResponseTypeFileWatcherUpdated           GlobalEventResponseType = "file.watcher.updated"
	GlobalEventResponseTypeGlobalDisposed               GlobalEventResponseType = "global.disposed"
	GlobalEventResponseTypeInstallationUpdateAvailable  GlobalEventResponseType = "installation.update-available"
	GlobalEventResponseTypeInstallationUpdated          GlobalEventResponseType = "installation.updated"
	GlobalEventResponseTypeLspUpdated                   GlobalEventResponseType = "lsp.updated"
	GlobalEventResponseTypeMcpBrowserOpenFailed         GlobalEventResponseType = "mcp.browser.open.failed"
	GlobalEventResponseTypeMcpToolsChanged              GlobalEventResponseType = "mcp.tools.changed"
	GlobalEventResponseTypeModelsDevRefreshed           GlobalEventResponseType = "models-dev.refreshed"
	GlobalEventResponseTypeMessagePartDelta             GlobalEventResponseType = "message.part.delta"
	GlobalEventResponseTypeMessagePartRemoved           GlobalEventResponseType = "message.part.removed"
	GlobalEventResponseTypeMessagePartUpdated           GlobalEventResponseType = "message.part.updated"
	GlobalEventResponseTypeMessageRemoved               GlobalEventResponseType = "message.removed"
	GlobalEventResponseTypeMessageUpdated               GlobalEventResponseType = "message.updated"
	GlobalEventResponseTypePermissionAsked              GlobalEventResponseType = "permission.asked"
	GlobalEventResponseTypePermissionReplied            GlobalEventResponseType = "permission.replied"
	GlobalEventResponseTypePluginAdded                  GlobalEventResponseType = "plugin.added"
	GlobalEventResponseTypeProjectUpdated               GlobalEventResponseType = "project.updated"
	GlobalEventResponseTypePtyCreated                   GlobalEventResponseType = "pty.created"
	GlobalEventResponseTypePtyDeleted                   GlobalEventResponseType = "pty.deleted"
	GlobalEventResponseTypePtyExited                    GlobalEventResponseType = "pty.exited"
	GlobalEventResponseTypePtyUpdated                   GlobalEventResponseType = "pty.updated"
	GlobalEventResponseTypeQuestionAsked                GlobalEventResponseType = "question.asked"
	GlobalEventResponseTypeQuestionRejected             GlobalEventResponseType = "question.rejected"
	GlobalEventResponseTypeQuestionReplied              GlobalEventResponseType = "question.replied"
	GlobalEventResponseTypeServerConnected              GlobalEventResponseType = "server.connected"
	GlobalEventResponseTypeServerInstanceDisposed       GlobalEventResponseType = "server.instance.disposed"
	GlobalEventResponseTypeSessionCompacted             GlobalEventResponseType = "session.compacted"
	GlobalEventResponseTypeSessionCreated               GlobalEventResponseType = "session.created"
	GlobalEventResponseTypeSessionDeleted               GlobalEventResponseType = "session.deleted"
	GlobalEventResponseTypeSessionDiff                  GlobalEventResponseType = "session.diff"
	GlobalEventResponseTypeSessionError                 GlobalEventResponseType = "session.error"
	GlobalEventResponseTypeSessionIdle                  GlobalEventResponseType = "session.idle"
	GlobalEventResponseTypeSessionNextAgentSwitched     GlobalEventResponseType = "session.next.agent.switched"
	GlobalEventResponseTypeSessionNextCompactionDelta   GlobalEventResponseType = "session.next.compaction.delta"
	GlobalEventResponseTypeSessionNextCompactionEnded   GlobalEventResponseType = "session.next.compaction.ended"
	GlobalEventResponseTypeSessionNextCompactionStarted GlobalEventResponseType = "session.next.compaction.started"
	GlobalEventResponseTypeSessionNextModelSwitched     GlobalEventResponseType = "session.next.model.switched"
	GlobalEventResponseTypeSessionNextPrompted          GlobalEventResponseType = "session.next.prompted"
	GlobalEventResponseTypeSessionNextReasoningDelta    GlobalEventResponseType = "session.next.reasoning.delta"
	GlobalEventResponseTypeSessionNextReasoningEnded    GlobalEventResponseType = "session.next.reasoning.ended"
	GlobalEventResponseTypeSessionNextReasoningStarted  GlobalEventResponseType = "session.next.reasoning.started"
	GlobalEventResponseTypeSessionNextRetried           GlobalEventResponseType = "session.next.retried"
	GlobalEventResponseTypeSessionNextShellEnded        GlobalEventResponseType = "session.next.shell.ended"
	GlobalEventResponseTypeSessionNextShellStarted      GlobalEventResponseType = "session.next.shell.started"
	GlobalEventResponseTypeSessionNextStepEnded         GlobalEventResponseType = "session.next.step.ended"
	GlobalEventResponseTypeSessionNextStepFailed        GlobalEventResponseType = "session.next.step.failed"
	GlobalEventResponseTypeSessionNextStepStarted       GlobalEventResponseType = "session.next.step.started"
	GlobalEventResponseTypeSessionNextSynthetic         GlobalEventResponseType = "session.next.synthetic"
	GlobalEventResponseTypeSessionNextTextDelta         GlobalEventResponseType = "session.next.text.delta"
	GlobalEventResponseTypeSessionNextTextEnded         GlobalEventResponseType = "session.next.text.ended"
	GlobalEventResponseTypeSessionNextTextStarted       GlobalEventResponseType = "session.next.text.started"
	GlobalEventResponseTypeSessionNextToolCalled        GlobalEventResponseType = "session.next.tool.called"
	GlobalEventResponseTypeSessionNextToolFailed        GlobalEventResponseType = "session.next.tool.failed"
	GlobalEventResponseTypeSessionNextToolInputDelta    GlobalEventResponseType = "session.next.tool.input.delta"
	GlobalEventResponseTypeSessionNextToolInputEnded    GlobalEventResponseType = "session.next.tool.input.ended"
	GlobalEventResponseTypeSessionNextToolInputStarted  GlobalEventResponseType = "session.next.tool.input.started"
	GlobalEventResponseTypeSessionNextToolProgress      GlobalEventResponseType = "session.next.tool.progress"
	GlobalEventResponseTypeSessionNextToolSuccess       GlobalEventResponseType = "session.next.tool.success"
	GlobalEventResponseTypeSessionStatus                GlobalEventResponseType = "session.status"
	GlobalEventResponseTypeSessionUpdated               GlobalEventResponseType = "session.updated"
	GlobalEventResponseTypeTodoUpdated                  GlobalEventResponseType = "todo.updated"
	GlobalEventResponseTypeTuiCommandExecute            GlobalEventResponseType = "tui.command.execute"
	GlobalEventResponseTypeTuiPromptAppend              GlobalEventResponseType = "tui.prompt.append"
	GlobalEventResponseTypeTuiSessionSelect             GlobalEventResponseType = "tui.session.select"
	GlobalEventResponseTypeTuiToastShow                 GlobalEventResponseType = "tui.toast.show"
	GlobalEventResponseTypeVcsBranchUpdated             GlobalEventResponseType = "vcs.branch.updated"
	GlobalEventResponseTypeWorkspaceFailed              GlobalEventResponseType = "workspace.failed"
	GlobalEventResponseTypeWorkspaceReady               GlobalEventResponseType = "workspace.ready"
	GlobalEventResponseTypeWorkspaceStatus              GlobalEventResponseType = "workspace.status"
	GlobalEventResponseTypeWorktreeFailed               GlobalEventResponseType = "worktree.failed"
	GlobalEventResponseTypeWorktreeReady                GlobalEventResponseType = "worktree.ready"
	// SyncEventResponse type — all V1 SyncEvents use "sync" as the type value.
	GlobalEventResponseTypeSync                         GlobalEventResponseType = "sync"
	GlobalEventResponseTypeIntegrationUpdated           GlobalEventResponseType = "integration.updated"
	GlobalEventResponseTypeIntegrationConnectionUpdated GlobalEventResponseType = "integration.connection.updated"
	GlobalEventResponseTypeCatalogUpdated               GlobalEventResponseType = "catalog.updated"
	GlobalEventResponseTypePermissionV2Asked            GlobalEventResponseType = "permission.v2.asked"
	GlobalEventResponseTypePermissionV2Replied          GlobalEventResponseType = "permission.v2.replied"
	GlobalEventResponseTypeReferenceUpdated             GlobalEventResponseType = "reference.updated"
	GlobalEventResponseTypeQuestionV2Asked              GlobalEventResponseType = "question.v2.asked"
	GlobalEventResponseTypeQuestionV2Replied            GlobalEventResponseType = "question.v2.replied"
	GlobalEventResponseTypeQuestionV2Rejected           GlobalEventResponseType = "question.v2.rejected"
	GlobalEventResponseTypeSessionNextMoved             GlobalEventResponseType = "session.next.moved"
	GlobalEventResponseTypeSessionNextRevertStaged      GlobalEventResponseType = "session.next.revert.staged"
	GlobalEventResponseTypeSessionNextRevertCleared     GlobalEventResponseType = "session.next.revert.cleared"
	GlobalEventResponseTypeSessionNextRevertCommitted   GlobalEventResponseType = "session.next.revert.committed"
	GlobalEventResponseTypeSessionNextPromptAdmitted    GlobalEventResponseType = "session.next.prompt.admitted"
	GlobalEventResponseTypeSessionNextContextUpdated    GlobalEventResponseType = "session.next.context.updated"
	GlobalEventResponseTypeProjectDirectoriesUpdated    GlobalEventResponseType = "project.directories.updated"
)

func (r GlobalEventResponseType) IsKnown() bool {
	switch r {
	case GlobalEventResponseTypeCommandExecuted,
		GlobalEventResponseTypeFileEdited,
		GlobalEventResponseTypeFileWatcherUpdated,
		GlobalEventResponseTypeGlobalDisposed,
		GlobalEventResponseTypeInstallationUpdateAvailable,
		GlobalEventResponseTypeInstallationUpdated,
		GlobalEventResponseTypeLspUpdated,
		GlobalEventResponseTypeMcpBrowserOpenFailed,
		GlobalEventResponseTypeMcpToolsChanged,
		GlobalEventResponseTypeModelsDevRefreshed,
		GlobalEventResponseTypeMessagePartDelta,
		GlobalEventResponseTypeMessagePartRemoved,
		GlobalEventResponseTypeMessagePartUpdated,
		GlobalEventResponseTypeMessageRemoved,
		GlobalEventResponseTypeMessageUpdated,
		GlobalEventResponseTypePermissionAsked,
		GlobalEventResponseTypePermissionReplied,
		GlobalEventResponseTypePluginAdded,
		GlobalEventResponseTypeProjectUpdated,
		GlobalEventResponseTypePtyCreated,
		GlobalEventResponseTypePtyDeleted,
		GlobalEventResponseTypePtyExited,
		GlobalEventResponseTypePtyUpdated,
		GlobalEventResponseTypeQuestionAsked,
		GlobalEventResponseTypeQuestionRejected,
		GlobalEventResponseTypeQuestionReplied,
		GlobalEventResponseTypeServerConnected,
		GlobalEventResponseTypeServerInstanceDisposed,
		GlobalEventResponseTypeSessionCompacted,
		GlobalEventResponseTypeSessionCreated,
		GlobalEventResponseTypeSessionDeleted,
		GlobalEventResponseTypeSessionDiff,
		GlobalEventResponseTypeSessionError,
		GlobalEventResponseTypeSessionIdle,
		GlobalEventResponseTypeSessionNextAgentSwitched,
		GlobalEventResponseTypeSessionNextCompactionDelta,
		GlobalEventResponseTypeSessionNextCompactionEnded,
		GlobalEventResponseTypeSessionNextCompactionStarted,
		GlobalEventResponseTypeSessionNextModelSwitched,
		GlobalEventResponseTypeSessionNextPrompted,
		GlobalEventResponseTypeSessionNextReasoningDelta,
		GlobalEventResponseTypeSessionNextReasoningEnded,
		GlobalEventResponseTypeSessionNextReasoningStarted,
		GlobalEventResponseTypeSessionNextRetried,
		GlobalEventResponseTypeSessionNextShellEnded,
		GlobalEventResponseTypeSessionNextShellStarted,
		GlobalEventResponseTypeSessionNextStepEnded,
		GlobalEventResponseTypeSessionNextStepFailed,
		GlobalEventResponseTypeSessionNextStepStarted,
		GlobalEventResponseTypeSessionNextSynthetic,
		GlobalEventResponseTypeSessionNextTextDelta,
		GlobalEventResponseTypeSessionNextTextEnded,
		GlobalEventResponseTypeSessionNextTextStarted,
		GlobalEventResponseTypeSessionNextToolCalled,
		GlobalEventResponseTypeSessionNextToolFailed,
		GlobalEventResponseTypeSessionNextToolInputDelta,
		GlobalEventResponseTypeSessionNextToolInputEnded,
		GlobalEventResponseTypeSessionNextToolInputStarted,
		GlobalEventResponseTypeSessionNextToolProgress,
		GlobalEventResponseTypeSessionNextToolSuccess,
		GlobalEventResponseTypeSessionStatus,
		GlobalEventResponseTypeSessionUpdated,
		GlobalEventResponseTypeTodoUpdated,
		GlobalEventResponseTypeTuiCommandExecute,
		GlobalEventResponseTypeTuiPromptAppend,
		GlobalEventResponseTypeTuiSessionSelect,
		GlobalEventResponseTypeTuiToastShow,
		GlobalEventResponseTypeVcsBranchUpdated,
		GlobalEventResponseTypeWorkspaceFailed,
		GlobalEventResponseTypeWorkspaceReady,
		GlobalEventResponseTypeWorkspaceStatus,
		GlobalEventResponseTypeWorktreeFailed,
		GlobalEventResponseTypeWorktreeReady,
		GlobalEventResponseTypeSync,
		GlobalEventResponseTypeIntegrationUpdated,
		GlobalEventResponseTypeIntegrationConnectionUpdated,
		GlobalEventResponseTypeCatalogUpdated,
		GlobalEventResponseTypePermissionV2Asked,
		GlobalEventResponseTypePermissionV2Replied,
		GlobalEventResponseTypeReferenceUpdated,
		GlobalEventResponseTypeQuestionV2Asked,
		GlobalEventResponseTypeQuestionV2Replied,
		GlobalEventResponseTypeQuestionV2Rejected,
		GlobalEventResponseTypeSessionNextMoved,
		GlobalEventResponseTypeSessionNextRevertStaged,
		GlobalEventResponseTypeSessionNextRevertCleared,
		GlobalEventResponseTypeSessionNextRevertCommitted,
		GlobalEventResponseTypeSessionNextPromptAdmitted,
		GlobalEventResponseTypeSessionNextContextUpdated,
		GlobalEventResponseTypeProjectDirectoriesUpdated:
		return true
	}
	return false
}

// SyncEventResponseSyncEventType values correspond to sync event versioned type names.
type SyncEventResponseSyncEventType string

const (
	SyncEventResponseSyncEventTypeMessageUpdated1               SyncEventResponseSyncEventType = "message.updated.1"
	SyncEventResponseSyncEventTypeMessageRemoved1               SyncEventResponseSyncEventType = "message.removed.1"
	SyncEventResponseSyncEventTypeMessagePartUpdated1           SyncEventResponseSyncEventType = "message.part.updated.1"
	SyncEventResponseSyncEventTypeMessagePartRemoved1           SyncEventResponseSyncEventType = "message.part.removed.1"
	SyncEventResponseSyncEventTypeSessionCreated1               SyncEventResponseSyncEventType = "session.created.1"
	SyncEventResponseSyncEventTypeSessionUpdated1               SyncEventResponseSyncEventType = "session.updated.1"
	SyncEventResponseSyncEventTypeSessionDeleted1               SyncEventResponseSyncEventType = "session.deleted.1"
	SyncEventResponseSyncEventTypeSessionNextAgentSwitched1     SyncEventResponseSyncEventType = "session.next.agent.switched.1"
	SyncEventResponseSyncEventTypeSessionNextModelSwitched1     SyncEventResponseSyncEventType = "session.next.model.switched.1"
	SyncEventResponseSyncEventTypeSessionNextPrompted1          SyncEventResponseSyncEventType = "session.next.prompted.1"
	SyncEventResponseSyncEventTypeSessionNextPromptAdmitted1    SyncEventResponseSyncEventType = "session.next.prompt.admitted.1"
	SyncEventResponseSyncEventTypeSessionNextSynthetic1         SyncEventResponseSyncEventType = "session.next.synthetic.1"
	SyncEventResponseSyncEventTypeSessionNextShellStarted1      SyncEventResponseSyncEventType = "session.next.shell.started.1"
	SyncEventResponseSyncEventTypeSessionNextContextUpdated1    SyncEventResponseSyncEventType = "session.next.context.updated.1"
	SyncEventResponseSyncEventTypeSessionNextShellEnded1        SyncEventResponseSyncEventType = "session.next.shell.ended.1"
	SyncEventResponseSyncEventTypeSessionNextStepStarted1       SyncEventResponseSyncEventType = "session.next.step.started.1"
	SyncEventResponseSyncEventTypeSessionNextStepEnded2         SyncEventResponseSyncEventType = "session.next.step.ended.2"
	SyncEventResponseSyncEventTypeSessionNextStepFailed2        SyncEventResponseSyncEventType = "session.next.step.failed.2"
	SyncEventResponseSyncEventTypeSessionNextTextStarted1       SyncEventResponseSyncEventType = "session.next.text.started.1"
	SyncEventResponseSyncEventTypeSessionNextTextEnded1         SyncEventResponseSyncEventType = "session.next.text.ended.1"
	SyncEventResponseSyncEventTypeSessionNextReasoningStarted1  SyncEventResponseSyncEventType = "session.next.reasoning.started.1"
	SyncEventResponseSyncEventTypeSessionNextReasoningEnded1    SyncEventResponseSyncEventType = "session.next.reasoning.ended.1"
	SyncEventResponseSyncEventTypeSessionNextToolInputStarted1  SyncEventResponseSyncEventType = "session.next.tool.input.started.1"
	SyncEventResponseSyncEventTypeSessionNextToolInputEnded1    SyncEventResponseSyncEventType = "session.next.tool.input.ended.1"
	SyncEventResponseSyncEventTypeSessionNextToolCalled1        SyncEventResponseSyncEventType = "session.next.tool.called.1"
	SyncEventResponseSyncEventTypeSessionNextToolProgress1      SyncEventResponseSyncEventType = "session.next.tool.progress.1"
	SyncEventResponseSyncEventTypeSessionNextToolSuccess1       SyncEventResponseSyncEventType = "session.next.tool.success.1"
	SyncEventResponseSyncEventTypeSessionNextToolFailed1        SyncEventResponseSyncEventType = "session.next.tool.failed.1"
	SyncEventResponseSyncEventTypeSessionNextRetried1           SyncEventResponseSyncEventType = "session.next.retried.1"
	SyncEventResponseSyncEventTypeSessionNextCompactionStarted1 SyncEventResponseSyncEventType = "session.next.compaction.started.1"
	SyncEventResponseSyncEventTypeSessionNextCompactionEnded1   SyncEventResponseSyncEventType = "session.next.compaction.ended.1"
	SyncEventResponseSyncEventTypeSessionNextMoved1             SyncEventResponseSyncEventType = "session.next.moved.1"
	SyncEventResponseSyncEventTypeSessionNextRevertStaged1      SyncEventResponseSyncEventType = "session.next.revert.staged.1"
	SyncEventResponseSyncEventTypeSessionNextRevertCleared1     SyncEventResponseSyncEventType = "session.next.revert.cleared.1"
	SyncEventResponseSyncEventTypeSessionNextRevertCommitted1   SyncEventResponseSyncEventType = "session.next.revert.committed.1"
)

func (r SyncEventResponseSyncEventType) IsKnown() bool {
	switch r {
	case SyncEventResponseSyncEventTypeMessageUpdated1,
		SyncEventResponseSyncEventTypeMessageRemoved1,
		SyncEventResponseSyncEventTypeMessagePartUpdated1,
		SyncEventResponseSyncEventTypeMessagePartRemoved1,
		SyncEventResponseSyncEventTypeSessionCreated1,
		SyncEventResponseSyncEventTypeSessionUpdated1,
		SyncEventResponseSyncEventTypeSessionDeleted1,
		SyncEventResponseSyncEventTypeSessionNextAgentSwitched1,
		SyncEventResponseSyncEventTypeSessionNextModelSwitched1,
		SyncEventResponseSyncEventTypeSessionNextPrompted1,
		SyncEventResponseSyncEventTypeSessionNextPromptAdmitted1,
		SyncEventResponseSyncEventTypeSessionNextSynthetic1,
		SyncEventResponseSyncEventTypeSessionNextShellStarted1,
		SyncEventResponseSyncEventTypeSessionNextContextUpdated1,
		SyncEventResponseSyncEventTypeSessionNextShellEnded1,
		SyncEventResponseSyncEventTypeSessionNextStepStarted1,
		SyncEventResponseSyncEventTypeSessionNextStepEnded2,
		SyncEventResponseSyncEventTypeSessionNextStepFailed2,
		SyncEventResponseSyncEventTypeSessionNextTextStarted1,
		SyncEventResponseSyncEventTypeSessionNextTextEnded1,
		SyncEventResponseSyncEventTypeSessionNextReasoningStarted1,
		SyncEventResponseSyncEventTypeSessionNextReasoningEnded1,
		SyncEventResponseSyncEventTypeSessionNextToolInputStarted1,
		SyncEventResponseSyncEventTypeSessionNextToolInputEnded1,
		SyncEventResponseSyncEventTypeSessionNextToolCalled1,
		SyncEventResponseSyncEventTypeSessionNextToolProgress1,
		SyncEventResponseSyncEventTypeSessionNextToolSuccess1,
		SyncEventResponseSyncEventTypeSessionNextToolFailed1,
		SyncEventResponseSyncEventTypeSessionNextRetried1,
		SyncEventResponseSyncEventTypeSessionNextCompactionStarted1,
		SyncEventResponseSyncEventTypeSessionNextCompactionEnded1,
		SyncEventResponseSyncEventTypeSessionNextMoved1,
		SyncEventResponseSyncEventTypeSessionNextRevertStaged1,
		SyncEventResponseSyncEventTypeSessionNextRevertCleared1,
		SyncEventResponseSyncEventTypeSessionNextRevertCommitted1:
		return true
	}
	return false
}

// SyncEventResponseSyncEventDataUnion is satisfied by all V1 SyncEvent types.
type SyncEventResponseSyncEventDataUnion interface {
	implementsSyncEventResponseSyncEventDataUnion()
}

// SyncEventResponseSyncEvent holds the actual SyncEvent fields.
// The Data field is populated via union matching from the underlying SyncEvent type.
type SyncEventResponseSyncEvent struct {
	Type        SyncEventResponseSyncEventType `json:"type,required"`
	ID          string                         `json:"id,required"`
	Seq         int64                          `json:"seq,required"`
	AggregateID string                         `json:"aggregateID,required"`
	// This field can have the runtime type of
	// [EventListResponseEventMessageUpdatedProperties],
	// [EventListResponseEventMessageRemovedProperties],
	// [EventListResponseEventMessagePartUpdatedProperties],
	// [EventListResponseEventMessagePartRemovedProperties],
	// [EventListResponseEventSessionCreatedProperties],
	// [EventListResponseEventSessionUpdatedProperties],
	// [EventListResponseEventSessionDeletedProperties],
	// [EventListResponseEventSessionNextAgentSwitchedProperties],
	// [EventListResponseEventSessionNextModelSwitchedProperties],
	// [EventListResponseEventSessionNextPromptedProperties],
	// [EventListResponseEventSessionNextPromptAdmittedProperties],
	// [EventListResponseEventSessionNextSyntheticProperties],
	// [EventListResponseEventSessionNextShellStartedProperties],
	// [EventListResponseEventSessionNextContextUpdatedProperties],
	// [EventListResponseEventSessionNextShellEndedProperties],
	// [EventListResponseEventSessionNextStepStartedProperties],
	// [EventListResponseEventSessionNextStepEndedProperties],
	// [EventListResponseEventSessionNextStepFailedProperties],
	// [EventListResponseEventSessionNextTextStartedProperties],
	// [EventListResponseEventSessionNextTextEndedProperties],
	// [EventListResponseEventSessionNextReasoningStartedProperties],
	// [EventListResponseEventSessionNextReasoningEndedProperties],
	// [EventListResponseEventSessionNextToolInputStartedProperties],
	// [EventListResponseEventSessionNextToolInputEndedProperties],
	// [EventListResponseEventSessionNextToolCalledProperties],
	// [EventListResponseEventSessionNextToolProgressProperties],
	// [EventListResponseEventSessionNextToolSuccessProperties],
	// [EventListResponseEventSessionNextToolFailedProperties],
	// [EventListResponseEventSessionNextRetriedProperties],
	// [EventListResponseEventSessionNextCompactionStartedProperties],
	// [EventListResponseEventSessionNextCompactionEndedProperties],
	// [EventListResponseEventSessionNextMovedProperties],
	// [EventListResponseEventSessionNextRevertStagedProperties],
	// [EventListResponseEventSessionNextRevertClearedProperties],
	// [EventListResponseEventSessionNextRevertCommittedProperties].
	Data  any                            `json:"data,required"`
	JSON  syncEventResponseSyncEventJSON `json:"-"`
	union SyncEventResponseSyncEventDataUnion
}

type syncEventResponseSyncEventJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventResponseSyncEventJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventResponseSyncEvent) UnmarshalJSON(data []byte) (err error) {
	*r = SyncEventResponseSyncEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SyncEventResponseSyncEventDataUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are all V1 SyncEvent types (SyncEvent*), such
// as [SyncEventMessageUpdated], [SyncEventSessionCreated] or
// [SyncEventSessionNextToolCalled].
func (r SyncEventResponseSyncEvent) AsUnion() SyncEventResponseSyncEventDataUnion {
	return r.union
}
