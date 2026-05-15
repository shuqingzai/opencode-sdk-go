// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
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
func (r *GlobalService) ConfigUpdate(ctx context.Context, body Config, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
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

type GlobalEvent struct {
	Directory string                 `json:"directory,required"`
	// This field can have the runtime type of
	// [EventListResponseEventXxx] (73 V2 Event types):
	// [EventListResponseEventCommandExecuted],
	// [EventListResponseEventFileEdited],
	// [EventListResponseEventFileWatcherUpdated],
	// [EventListResponseEventGlobalDisposed],
	// [EventListResponseEventInstallationUpdateAvailable],
	// [EventListResponseEventInstallationUpdated],
	// [EventListResponseEventLspClientDiagnostics],
	// [EventListResponseEventLspUpdated],
	// [EventListResponseEventMcpBrowserOpenFailed],
	// [EventListResponseEventMcpToolsChanged],
	// [EventListResponseEventMessagePartDelta],
	// [EventListResponseEventMessagePartRemoved],
	// [EventListResponseEventMessagePartUpdated],
	// [EventListResponseEventMessageRemoved],
	// [EventListResponseEventMessageUpdated],
	// [EventListResponseEventPermissionAsked],
	// [EventListResponseEventPermissionReplied],
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
	// [EventListResponseEventCatalogModelUpdated],
	//
	// [SyncEventXxx] (33 V1 SyncEvent types):
	// [SyncEventMessagePartRemoved],
	// [SyncEventMessagePartUpdated],
	// [SyncEventMessageRemoved],
	// [SyncEventMessageUpdated],
	// [SyncEventSessionCreated],
	// [SyncEventSessionDeleted],
	// [SyncEventSessionNextAgentSwitched],
	// [SyncEventSessionNextCompactionDelta],
	// [SyncEventSessionNextCompactionEnded],
	// [SyncEventSessionNextCompactionStarted],
	// [SyncEventSessionNextModelSwitched],
	// [SyncEventSessionNextPrompted],
	// [SyncEventSessionNextReasoningDelta],
	// [SyncEventSessionNextReasoningEnded],
	// [SyncEventSessionNextReasoningStarted],
	// [SyncEventSessionNextRetried],
	// [SyncEventSessionNextShellEnded],
	// [SyncEventSessionNextShellStarted],
	// [SyncEventSessionNextStepEnded],
	// [SyncEventSessionNextStepFailed],
	// [SyncEventSessionNextStepStarted],
	// [SyncEventSessionNextSynthetic],
	// [SyncEventSessionNextTextDelta],
	// [SyncEventSessionNextTextEnded],
	// [SyncEventSessionNextTextStarted],
	// [SyncEventSessionNextToolCalled],
	// [SyncEventSessionNextToolFailed],
	// [SyncEventSessionNextToolInputDelta],
	// [SyncEventSessionNextToolInputEnded],
	// [SyncEventSessionNextToolInputStarted],
	// [SyncEventSessionNextToolProgress],
	// [SyncEventSessionNextToolSuccess],
	// [SyncEventSessionUpdated].
	Payload   interface{}            `json:"payload,required"`
	Project   string                 `json:"project"`
	Workspace string                 `json:"workspace"`
	JSON      globalEventJSON        `json:"-"`
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
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [GlobalEventPayloadUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are all V2 Event types (EventListResponseEvent*)
// and all V1 SyncEvent types (SyncEvent*).
func (r GlobalEvent) AsUnion() GlobalEventPayloadUnion {
	return r.union
}

// GlobalEventPayloadUnion is a union of all V2 Event types and V1 SyncEvent types
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
			Type:       reflect.TypeOf(EventListResponseEventLspClientDiagnostics{}),
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
		// V1 SyncEvent types
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
			Type:       reflect.TypeOf(SyncEventSessionNextSynthetic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextShellStarted{}),
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
			Type:       reflect.TypeOf(SyncEventSessionNextTextDelta{}),
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
			Type:       reflect.TypeOf(SyncEventSessionNextReasoningDelta{}),
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
			Type:       reflect.TypeOf(SyncEventSessionNextToolInputDelta{}),
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
			Type:       reflect.TypeOf(SyncEventSessionNextCompactionDelta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SyncEventSessionNextCompactionEnded{}),
		},
		// V2 Event: catalog.model.updated
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventCatalogModelUpdated{}),
		},
	)
}
