// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
	"github.com/tidwall/gjson"
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
func (r *GlobalService) Upgrade(ctx context.Context, body GlobalUpgradeBody, opts ...option.RequestOption) (res *GlobalUpgradeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/upgrade"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// GlobalConfigUpdateParams contains the parameters for updating global configuration.
// All fields are optional for PATCH semantics.
type GlobalConfigUpdateParams struct {
	// Body parameters — all Config fields as optional
	Schema            param.Field[string]                     `json:"$schema"`
	Agent             param.Field[ConfigAgent]                `json:"agent"`
	Attachment        param.Field[AttachmentConfig]           `json:"attachment"`
	Autoshare         param.Field[bool]                       `json:"autoshare"`
	Autoupdate        param.Field[interface{}]                `json:"autoupdate"`
	Command           param.Field[map[string]ConfigCommand]   `json:"command"`
	Compaction        param.Field[ConfigCompaction]           `json:"compaction"`
	DisabledProviders param.Field[[]string]                   `json:"disabled_providers"`
	EnabledProviders  param.Field[[]string]                   `json:"enabled_providers"`
	Enterprise        param.Field[EnterpriseConfig]           `json:"enterprise"`
	Experimental      param.Field[ConfigExperimental]         `json:"experimental"`
	Formatter         param.Field[map[string]ConfigFormatter] `json:"formatter"`
	Instructions      param.Field[[]string]                   `json:"instructions"`
	Layout            param.Field[ConfigLayout]               `json:"layout"`
	LogLevel          param.Field[ConfigLogLevel]             `json:"logLevel"`
	Lsp               param.Field[map[string]ConfigLsp]       `json:"lsp"`
	Mcp               param.Field[map[string]ConfigMcp]       `json:"mcp"`
	Mode              param.Field[ConfigMode]                 `json:"mode"`
	Model             param.Field[string]                     `json:"model"`
	Permission        param.Field[ConfigPermission]           `json:"permission"`
	// Plugins to load. Each item is either a plugin name (string) or a 2-tuple
	// of [pluginName, configObject] (where configObject is a map[string]any).
	Plugin        param.Field[[]interface{}]             `json:"plugin"`
	Provider      param.Field[map[string]ConfigProvider] `json:"provider"`
	// Map of reference name → value. Each value can be a plain [string] (URL/path),
	// a [ConfigV2ReferenceGit], or a [ConfigV2ReferenceLocal].
	Reference     param.Field[map[string]interface{}]    `json:"reference"`
	// Map of reference name → value. Each value can be a plain [string] (URL/path),
	// a [ConfigV2ReferenceGit], or a [ConfigV2ReferenceLocal].
	References    param.Field[map[string]interface{}]    `json:"references"`
	Share         param.Field[ConfigShare]               `json:"share"`
	Shell         param.Field[string]                    `json:"shell"`
	Server        param.Field[ServerConfig]              `json:"server"`
	Skills        param.Field[ConfigSkills]              `json:"skills"`
	SmallModel    param.Field[string]                    `json:"small_model"`
	Snapshot      param.Field[bool]                      `json:"snapshot"`
	ToolOutput    param.Field[ConfigToolOutput]          `json:"tool_output"`
	Tools         param.Field[map[string]bool]           `json:"tools"`
	Username      param.Field[string]                    `json:"username"`
	Watcher       param.Field[ConfigWatcher]             `json:"watcher"`
	DefaultAgent  param.Field[string]                    `json:"default_agent"`
	SubagentDepth param.Field[int64]                     `json:"subagent_depth"`
}

func (r GlobalConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GlobalEvent struct {
	Directory string `json:"directory,required"`
	// This field can have the runtime type of
	// [EventListResponseEventXxx] (88 V2 Event types):
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
	//
	// [SyncEventResponse] (V1 SyncEvent).
	Payload   interface{}     `json:"payload,required"`
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
	// unnecessary pass through the interface{} decoder. The full JSON metadata
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
	result := gjson.ParseBytes(data)
	payloadResult := result.Get("payload")
	if !payloadResult.Exists() {
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

	// V2 events: payload has the shape {id, type, properties}.
	// Match directly against the union of EventListResponseEventXxx types.
	if err := apijson.UnmarshalRoot([]byte(payloadResult.Raw), &r.union); err != nil {
		return err
	}
	r.Payload = r.union
	return nil
}

// AsUnion returns a [GlobalEventPayloadUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are all V2 Event types (EventListResponseEvent*)
// and SyncEventResponse (for V1 SyncEvent types).
func (r GlobalEvent) AsUnion() GlobalEventPayloadUnion {
	return r.union
}

// GlobalEventPayloadUnion is a union of all V2 Event types and SyncEventResponse
// that can appear in the GlobalEvent.payload field.
type GlobalEventPayloadUnion interface {
	implementsGlobalEventPayload()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalEventPayloadUnion)(nil)).Elem(),
		"",
		// V2 Event types
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventInstallationUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventInstallationUpdateAvailable{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventLspUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventMessageUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventMessageRemoved{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventMessagePartUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventMessagePartDelta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventMessagePartRemoved{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionCompacted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPermissionAsked{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPermissionReplied{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventFileEdited{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventFileWatcherUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTodoUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTuiPromptAppend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTuiCommandExecute{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTuiToastShow{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTuiSessionSelect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventMcpToolsChanged{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventModelsDevRefreshed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventMcpBrowserOpenFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventCommandExecuted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionIdle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionDiff{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionCreated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionDeleted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventServerConnected{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventServerInstanceDisposed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventGlobalDisposed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventProjectUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventQuestionAsked{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventQuestionRejected{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventQuestionReplied{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventVcsBranchUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorkspaceReady{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorkspaceFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorkspaceStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPtyCreated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPtyUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPtyExited{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPtyDeleted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorktreeReady{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventIntegrationUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventIntegrationConnectionUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventCatalogUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPermissionV2Asked{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPermissionV2Replied{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventReferenceUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventQuestionV2Asked{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventQuestionV2Replied{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventQuestionV2Rejected{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextMoved{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextRevertStaged{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextRevertCleared{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextRevertCommitted{}),
		},
		// V1 SyncEvent types
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorktreeFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextAgentSwitched{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextModelSwitched{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextPrompted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextSynthetic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextShellStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextShellEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextStepStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextStepEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextStepFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextTextStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextTextDelta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextTextEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextReasoningStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextReasoningDelta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextReasoningEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextToolInputStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextToolInputDelta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextToolInputEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextToolCalled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextToolProgress{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextToolSuccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextToolFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextRetried{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextCompactionStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextCompactionDelta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextCompactionEnded{}),
		},
		// V1 SyncEvent types — wrapped in SyncEventResponse
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventResponse{}),
		},
		// V2 Event: catalog.model.updated
		// V2 Event: plugin.added
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPluginAdded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextPromptAdmitted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionNextContextUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventProjectDirectoriesUpdated{}),
		},
	)

	// SyncEventResponseSyncEventDataUnion: V1 SyncEvent parent types registered
	// as union variants. Each parent type carries a polymorphic Data field whose
	// concrete type reuses the V2 EventListResponseEventXxxProperties for the
	// same event (except session.updated which uses a custom Data type).
	// The parent type's enum Type field provides the discriminator for matching.
	apijson.RegisterUnion(
		reflect.TypeOf((*SyncEventResponseSyncEventDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventMessageUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventMessageRemoved{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventMessagePartUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventMessagePartRemoved{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionCreated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionDeleted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextAgentSwitched{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextModelSwitched{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextPrompted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextPromptAdmitted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextSynthetic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextShellStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextContextUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextShellEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextStepStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextStepEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextStepFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextTextStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextTextEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextReasoningStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextReasoningEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextToolInputStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextToolInputEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextToolCalled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextToolProgress{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextToolSuccess{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextToolFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextRetried{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextCompactionStarted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextCompactionEnded{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextMoved{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextRevertStaged{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextRevertCleared{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextRevertCommitted{}),
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
	// This field's runtime type is determined by the underlying SyncEvent variant
	// (populated via union matching). Possible runtime types include:
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
	Data  interface{}                    `json:"data,required"`
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
