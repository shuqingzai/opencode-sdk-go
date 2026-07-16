// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/tidwall/gjson"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
	"github.com/sst/opencode-sdk-go/shared"
)

// EventService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// Get events
func (r *EventService) ListStreaming(ctx context.Context, query EventListParams, opts ...option.RequestOption) (stream *ssestream.Stream[EventListResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[EventListResponse](ssestream.NewDecoder(raw), err)
}

type EventListResponse struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of
	// [EventListResponseEventCommandExecutedProperties],
	// [EventListResponseEventFileEditedProperties],
	// [EventListResponseEventFileWatcherUpdatedProperties],
	// [EventListResponseEventGlobalDisposedProperties],
	// [EventListResponseEventInstallationUpdateAvailableProperties],
	// [EventListResponseEventInstallationUpdatedProperties],
	// [EventListResponseEventLspUpdatedProperties],
	// [EventListResponseEventMcpBrowserOpenFailedProperties],
	// [EventListResponseEventMcpToolsChangedProperties],
	// [EventListResponseEventModelsDevRefreshedProperties],
	// [EventListResponseEventMessagePartDeltaProperties],
	// [EventListResponseEventMessagePartRemovedProperties],
	// [EventListResponseEventMessagePartUpdatedProperties],
	// [EventListResponseEventMessageRemovedProperties],
	// [EventListResponseEventMessageUpdatedProperties],
	// [EventListResponseEventPermissionAskedProperties],
	// [EventListResponseEventPermissionRepliedProperties],
	// [EventListResponseEventPluginAddedProperties],
	// [EventListResponseEventProjectUpdatedProperties],
	// [EventListResponseEventPtyCreatedProperties],
	// [EventListResponseEventPtyDeletedProperties],
	// [EventListResponseEventPtyExitedProperties],
	// [EventListResponseEventPtyUpdatedProperties],
	// [EventListResponseEventQuestionAskedProperties],
	// [EventListResponseEventQuestionRejectedProperties],
	// [EventListResponseEventQuestionRepliedProperties],
	// [EventListResponseEventServerConnectedProperties],
	// [EventListResponseEventServerInstanceDisposedProperties],
	// [EventListResponseEventSessionCompactedProperties],
	// [EventListResponseEventSessionCreatedProperties],
	// [EventListResponseEventSessionDeletedProperties],
	// [EventListResponseEventSessionDiffProperties],
	// [EventListResponseEventSessionErrorProperties],
	// [EventListResponseEventSessionIdleProperties],
	// [EventListResponseEventSessionNextAgentSwitchedProperties],
	// [EventListResponseEventSessionNextCompactionDeltaProperties],
	// [EventListResponseEventSessionNextCompactionEndedProperties],
	// [EventListResponseEventSessionNextCompactionStartedProperties],
	// [EventListResponseEventSessionNextModelSwitchedProperties],
	// [EventListResponseEventSessionNextPromptedProperties],
	// [EventListResponseEventSessionNextReasoningDeltaProperties],
	// [EventListResponseEventSessionNextReasoningEndedProperties],
	// [EventListResponseEventSessionNextReasoningStartedProperties],
	// [EventListResponseEventSessionNextRetriedProperties],
	// [EventListResponseEventSessionNextShellEndedProperties],
	// [EventListResponseEventSessionNextShellStartedProperties],
	// [EventListResponseEventSessionNextStepEndedProperties],
	// [EventListResponseEventSessionNextStepFailedProperties],
	// [EventListResponseEventSessionNextStepStartedProperties],
	// [EventListResponseEventSessionNextSyntheticProperties],
	// [EventListResponseEventSessionNextTextDeltaProperties],
	// [EventListResponseEventSessionNextTextEndedProperties],
	// [EventListResponseEventSessionNextTextStartedProperties],
	// [EventListResponseEventSessionNextToolCalledProperties],
	// [EventListResponseEventSessionNextToolFailedProperties],
	// [EventListResponseEventSessionNextToolInputDeltaProperties],
	// [EventListResponseEventSessionNextToolInputEndedProperties],
	// [EventListResponseEventSessionNextToolInputStartedProperties],
	// [EventListResponseEventSessionNextToolProgressProperties],
	// [EventListResponseEventSessionNextToolSuccessProperties],
	// [EventListResponseEventSessionStatusProperties],
	// [EventListResponseEventSessionUpdatedProperties],
	// [EventListResponseEventTodoUpdatedProperties],
	// [EventListResponseEventTuiCommandExecuteProperties],
	// [EventListResponseEventTuiPromptAppendProperties],
	// [EventListResponseEventTuiSessionSelectProperties],
	// [EventListResponseEventTuiToastShowProperties],
	// [EventListResponseEventVcsBranchUpdatedProperties],
	// [EventListResponseEventWorkspaceFailedProperties],
	// [EventListResponseEventWorkspaceReadyProperties],
	// [EventListResponseEventWorkspaceStatusProperties],
	// [EventListResponseEventWorktreeFailedProperties],
	// [EventListResponseEventWorktreeReadyProperties],
	// [EventListResponseEventIntegrationUpdatedProperties],
	// [EventListResponseEventIntegrationConnectionUpdatedProperties],
	// [EventListResponseEventCatalogUpdatedProperties],
	// [EventListResponseEventPermissionV2AskedProperties],
	// [EventListResponseEventPermissionV2RepliedProperties],
	// [EventListResponseEventReferenceUpdatedProperties],
	// [EventListResponseEventQuestionV2AskedProperties],
	// [EventListResponseEventQuestionV2RepliedProperties],
	// [EventListResponseEventQuestionV2RejectedProperties],
	// [EventListResponseEventSessionNextMovedProperties],
	// [EventListResponseEventSessionNextRevertStagedProperties],
	// [EventListResponseEventSessionNextRevertClearedProperties],
	// [EventListResponseEventSessionNextRevertCommittedProperties],
	// [EventListResponseEventSessionNextPromptAdmittedProperties],
	// [EventListResponseEventSessionNextContextUpdatedProperties],
	// [EventListResponseEventProjectDirectoriesUpdatedProperties].
	Properties interface{}           `json:"properties,required"`
	Type       EventListResponseType `json:"type,required"`
	JSON       eventListResponseJSON `json:"-"`
	union      EventListResponseUnion
}

// eventListResponseJSON contains the JSON metadata for the struct
// [EventListResponse]
type eventListResponseJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r eventListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *EventListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = EventListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EventListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [EventListResponseEventCommandExecuted],
// [EventListResponseEventFileEdited],
// [EventListResponseEventFileWatcherUpdated],
// [EventListResponseEventGlobalDisposed],
// [EventListResponseEventInstallationUpdateAvailable],
// [EventListResponseEventInstallationUpdated],
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
// [EventListResponseEventProjectDirectoriesUpdated].
func (r EventListResponse) AsUnion() EventListResponseUnion {
	return r.union
}

// [EventListResponseEventCommandExecuted],
// [EventListResponseEventFileEdited],
// [EventListResponseEventFileWatcherUpdated],
// [EventListResponseEventGlobalDisposed],
// [EventListResponseEventInstallationUpdateAvailable],
// [EventListResponseEventInstallationUpdated],
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
// [EventListResponseEventProjectDirectoriesUpdated].
type EventListResponseUnion interface {
	implementsEventListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EventListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventCommandExecuted{}),
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
			Type:       reflect.TypeOf(EventListResponseEventGlobalDisposed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventInstallationUpdateAvailable{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventInstallationUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventLspUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventMcpBrowserOpenFailed{}),
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
			Type:       reflect.TypeOf(EventListResponseEventMessagePartDelta{}),
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
			Type:       reflect.TypeOf(EventListResponseEventProjectUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPtyCreated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPtyDeleted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPtyExited{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPtyUpdated{}),
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
			Type:       reflect.TypeOf(EventListResponseEventServerConnected{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventServerInstanceDisposed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionCompacted{}),
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
			Type:       reflect.TypeOf(EventListResponseEventSessionIdle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventSessionStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTodoUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTuiCommandExecute{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTuiPromptAppend{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTuiSessionSelect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventTuiToastShow{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventVcsBranchUpdated{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorkspaceFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorkspaceReady{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorkspaceStatus{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorktreeFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventWorktreeReady{}),
		},
		// V2 Event types from event_global_types.go
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
			Type:       reflect.TypeOf(EventListResponseEventMessagePartRemoved{}),
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
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventListResponseEventPluginAdded{}),
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
}

type EventListResponseEventInstallationUpdated struct {
	Properties EventListResponseEventInstallationUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventInstallationUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventInstallationUpdatedJSON       `json:"-"`
}

// eventListResponseEventInstallationUpdatedJSON contains the JSON metadata for the
// struct [EventListResponseEventInstallationUpdated]
type eventListResponseEventInstallationUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventInstallationUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventInstallationUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventInstallationUpdated) implementsEventListResponse() {}

func (r EventListResponseEventInstallationUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventInstallationUpdatedProperties struct {
	Version string                                                  `json:"version,required"`
	JSON    eventListResponseEventInstallationUpdatedPropertiesJSON `json:"-"`
}

// eventListResponseEventInstallationUpdatedPropertiesJSON contains the JSON
// metadata for the struct [EventListResponseEventInstallationUpdatedProperties]
type eventListResponseEventInstallationUpdatedPropertiesJSON struct {
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventInstallationUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventInstallationUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventInstallationUpdatedType string

const (
	EventListResponseEventInstallationUpdatedTypeInstallationUpdated EventListResponseEventInstallationUpdatedType = "installation.updated"
)

func (r EventListResponseEventInstallationUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventInstallationUpdatedTypeInstallationUpdated:
		return true
	}
	return false
}

type EventListResponseEventSessionCompacted struct {
	Properties EventListResponseEventSessionCompactedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionCompactedType       `json:"type,required"`
	JSON       eventListResponseEventSessionCompactedJSON       `json:"-"`
}

// eventListResponseEventSessionCompactedJSON contains the JSON metadata for the
// struct [EventListResponseEventSessionCompacted]
type eventListResponseEventSessionCompactedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionCompacted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionCompactedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionCompacted) implementsEventListResponse() {}

func (r EventListResponseEventSessionCompacted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionCompactedProperties struct {
	SessionID string                                               `json:"sessionID,required"`
	JSON      eventListResponseEventSessionCompactedPropertiesJSON `json:"-"`
}

// eventListResponseEventSessionCompactedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventSessionCompactedProperties]
type eventListResponseEventSessionCompactedPropertiesJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionCompactedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionCompactedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionCompactedType string

const (
	EventListResponseEventSessionCompactedTypeSessionCompacted EventListResponseEventSessionCompactedType = "session.compacted"
)

func (r EventListResponseEventSessionCompactedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionCompactedTypeSessionCompacted:
		return true
	}
	return false
}

type EventListResponseEventPermissionReplied struct {
	Properties EventListResponseEventPermissionRepliedProperties `json:"properties,required"`
	Type       EventListResponseEventPermissionRepliedType       `json:"type,required"`
	JSON       eventListResponseEventPermissionRepliedJSON       `json:"-"`
}

// eventListResponseEventPermissionRepliedJSON contains the JSON metadata for the
// struct [EventListResponseEventPermissionReplied]
type eventListResponseEventPermissionRepliedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionReplied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionRepliedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPermissionReplied) implementsEventListResponse() {}

func (r EventListResponseEventPermissionReplied) implementsGlobalEventPayload() {}

type EventListResponseEventPermissionRepliedProperties struct {
	RequestID string                                                `json:"requestID,required"`
	Reply     PermissionReplyParamsReply                            `json:"reply,required"`
	SessionID string                                                `json:"sessionID,required"`
	JSON      eventListResponseEventPermissionRepliedPropertiesJSON `json:"-"`
}

// eventListResponseEventPermissionRepliedPropertiesJSON contains the JSON metadata
// for the struct [EventListResponseEventPermissionRepliedProperties]
type eventListResponseEventPermissionRepliedPropertiesJSON struct {
	RequestID   apijson.Field
	Reply       apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionRepliedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionRepliedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPermissionRepliedType string

const (
	EventListResponseEventPermissionRepliedTypePermissionReplied EventListResponseEventPermissionRepliedType = "permission.replied"
)

func (r EventListResponseEventPermissionRepliedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPermissionRepliedTypePermissionReplied:
		return true
	}
	return false
}

type EventListResponseEventFileEdited struct {
	Properties EventListResponseEventFileEditedProperties `json:"properties,required"`
	Type       EventListResponseEventFileEditedType       `json:"type,required"`
	JSON       eventListResponseEventFileEditedJSON       `json:"-"`
}

// eventListResponseEventFileEditedJSON contains the JSON metadata for the struct
// [EventListResponseEventFileEdited]
type eventListResponseEventFileEditedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventFileEdited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventFileEditedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventFileEdited) implementsEventListResponse() {}

func (r EventListResponseEventFileEdited) implementsGlobalEventPayload() {}

type EventListResponseEventFileEditedProperties struct {
	File string                                         `json:"file,required"`
	JSON eventListResponseEventFileEditedPropertiesJSON `json:"-"`
}

// eventListResponseEventFileEditedPropertiesJSON contains the JSON metadata for
// the struct [EventListResponseEventFileEditedProperties]
type eventListResponseEventFileEditedPropertiesJSON struct {
	File        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventFileEditedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventFileEditedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventFileEditedType string

const (
	EventListResponseEventFileEditedTypeFileEdited EventListResponseEventFileEditedType = "file.edited"
)

func (r EventListResponseEventFileEditedType) IsKnown() bool {
	switch r {
	case EventListResponseEventFileEditedTypeFileEdited:
		return true
	}
	return false
}

type EventListResponseEventFileWatcherUpdated struct {
	Properties EventListResponseEventFileWatcherUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventFileWatcherUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventFileWatcherUpdatedJSON       `json:"-"`
}

// eventListResponseEventFileWatcherUpdatedJSON contains the JSON metadata for the
// struct [EventListResponseEventFileWatcherUpdated]
type eventListResponseEventFileWatcherUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventFileWatcherUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventFileWatcherUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventFileWatcherUpdated) implementsEventListResponse() {}

func (r EventListResponseEventFileWatcherUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventFileWatcherUpdatedProperties struct {
	Event EventListResponseEventFileWatcherUpdatedPropertiesEvent `json:"event,required"`
	File  string                                                  `json:"file,required"`
	JSON  eventListResponseEventFileWatcherUpdatedPropertiesJSON  `json:"-"`
}

// eventListResponseEventFileWatcherUpdatedPropertiesJSON contains the JSON
// metadata for the struct [EventListResponseEventFileWatcherUpdatedProperties]
type eventListResponseEventFileWatcherUpdatedPropertiesJSON struct {
	Event       apijson.Field
	File        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventFileWatcherUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventFileWatcherUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventFileWatcherUpdatedPropertiesEvent string

const (
	EventListResponseEventFileWatcherUpdatedPropertiesEventAdd    EventListResponseEventFileWatcherUpdatedPropertiesEvent = "add"
	EventListResponseEventFileWatcherUpdatedPropertiesEventChange EventListResponseEventFileWatcherUpdatedPropertiesEvent = "change"
	EventListResponseEventFileWatcherUpdatedPropertiesEventUnlink EventListResponseEventFileWatcherUpdatedPropertiesEvent = "unlink"
)

func (r EventListResponseEventFileWatcherUpdatedPropertiesEvent) IsKnown() bool {
	switch r {
	case EventListResponseEventFileWatcherUpdatedPropertiesEventAdd, EventListResponseEventFileWatcherUpdatedPropertiesEventChange, EventListResponseEventFileWatcherUpdatedPropertiesEventUnlink:
		return true
	}
	return false
}

type EventListResponseEventFileWatcherUpdatedType string

const (
	EventListResponseEventFileWatcherUpdatedTypeFileWatcherUpdated EventListResponseEventFileWatcherUpdatedType = "file.watcher.updated"
)

func (r EventListResponseEventFileWatcherUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventFileWatcherUpdatedTypeFileWatcherUpdated:
		return true
	}
	return false
}

type EventListResponseEventTodoUpdated struct {
	Properties EventListResponseEventTodoUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventTodoUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventTodoUpdatedJSON       `json:"-"`
}

// eventListResponseEventTodoUpdatedJSON contains the JSON metadata for the struct
// [EventListResponseEventTodoUpdated]
type eventListResponseEventTodoUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTodoUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTodoUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTodoUpdated) implementsEventListResponse() {}

func (r EventListResponseEventTodoUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventTodoUpdatedProperties struct {
	SessionID string                                            `json:"sessionID,required"`
	Todos     []EventListResponseEventTodoUpdatedPropertiesTodo `json:"todos,required"`
	JSON      eventListResponseEventTodoUpdatedPropertiesJSON   `json:"-"`
}

// eventListResponseEventTodoUpdatedPropertiesJSON contains the JSON metadata for
// the struct [EventListResponseEventTodoUpdatedProperties]
type eventListResponseEventTodoUpdatedPropertiesJSON struct {
	SessionID   apijson.Field
	Todos       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTodoUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTodoUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTodoUpdatedPropertiesTodo struct {
	// Unique identifier for the todo item
	ID string `json:"id,required"`
	// Brief description of the task
	Content string `json:"content,required"`
	// Priority level of the task: high, medium, low
	Priority string `json:"priority,required"`
	// Current status of the task: pending, in_progress, completed, cancelled
	Status string                                              `json:"status,required"`
	JSON   eventListResponseEventTodoUpdatedPropertiesTodoJSON `json:"-"`
}

// eventListResponseEventTodoUpdatedPropertiesTodoJSON contains the JSON metadata
// for the struct [EventListResponseEventTodoUpdatedPropertiesTodo]
type eventListResponseEventTodoUpdatedPropertiesTodoJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTodoUpdatedPropertiesTodo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTodoUpdatedPropertiesTodoJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTodoUpdatedType string

const (
	EventListResponseEventTodoUpdatedTypeTodoUpdated EventListResponseEventTodoUpdatedType = "todo.updated"
)

func (r EventListResponseEventTodoUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventTodoUpdatedTypeTodoUpdated:
		return true
	}
	return false
}

type EventListResponseEventSessionIdle struct {
	Properties EventListResponseEventSessionIdleProperties `json:"properties,required"`
	Type       EventListResponseEventSessionIdleType       `json:"type,required"`
	JSON       eventListResponseEventSessionIdleJSON       `json:"-"`
}

// eventListResponseEventSessionIdleJSON contains the JSON metadata for the struct
// [EventListResponseEventSessionIdle]
type eventListResponseEventSessionIdleJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionIdle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionIdleJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionIdle) implementsEventListResponse() {}

func (r EventListResponseEventSessionIdle) implementsGlobalEventPayload() {}

type EventListResponseEventSessionIdleProperties struct {
	SessionID string                                          `json:"sessionID,required"`
	JSON      eventListResponseEventSessionIdlePropertiesJSON `json:"-"`
}

// eventListResponseEventSessionIdlePropertiesJSON contains the JSON metadata for
// the struct [EventListResponseEventSessionIdleProperties]
type eventListResponseEventSessionIdlePropertiesJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionIdleProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionIdlePropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionIdleType string

const (
	EventListResponseEventSessionIdleTypeSessionIdle EventListResponseEventSessionIdleType = "session.idle"
)

func (r EventListResponseEventSessionIdleType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionIdleTypeSessionIdle:
		return true
	}
	return false
}

type EventListResponseEventSessionError struct {
	Properties EventListResponseEventSessionErrorProperties `json:"properties,required"`
	Type       EventListResponseEventSessionErrorType       `json:"type,required"`
	JSON       eventListResponseEventSessionErrorJSON       `json:"-"`
}

// eventListResponseEventSessionErrorJSON contains the JSON metadata for the struct
// [EventListResponseEventSessionError]
type eventListResponseEventSessionErrorJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionErrorJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionError) implementsEventListResponse() {}

func (r EventListResponseEventSessionError) implementsGlobalEventPayload() {}

type EventListResponseEventSessionErrorProperties struct {
	Error     EventListResponseEventSessionErrorPropertiesError `json:"error,required"`
	SessionID string                                            `json:"sessionID"`
	JSON      eventListResponseEventSessionErrorPropertiesJSON  `json:"-"`
}

// eventListResponseEventSessionErrorPropertiesJSON contains the JSON metadata for
// the struct [EventListResponseEventSessionErrorProperties]
type eventListResponseEventSessionErrorPropertiesJSON struct {
	Error       apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionErrorProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionErrorPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionErrorPropertiesError struct {
	// This field can have the runtime type of [shared.ProviderAuthErrorData],
	// [shared.UnknownErrorData], [interface{}], [shared.MessageAbortedErrorData],
	// [shared.StructuredOutputErrorData], [shared.ContextOverflowErrorData],
	// [shared.APIErrorData], [shared.ContentFilterErrorData].
	Data  interface{}                                           `json:"data,required"`
	Name  EventListResponseEventSessionErrorPropertiesErrorName `json:"name,required"`
	JSON  eventListResponseEventSessionErrorPropertiesErrorJSON `json:"-"`
	union EventListResponseEventSessionErrorPropertiesErrorUnion
}

// eventListResponseEventSessionErrorPropertiesErrorJSON contains the JSON metadata
// for the struct [EventListResponseEventSessionErrorPropertiesError]
type eventListResponseEventSessionErrorPropertiesErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r eventListResponseEventSessionErrorPropertiesErrorJSON) RawJSON() string {
	return r.raw
}

func (r *EventListResponseEventSessionErrorPropertiesError) UnmarshalJSON(data []byte) (err error) {
	*r = EventListResponseEventSessionErrorPropertiesError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EventListResponseEventSessionErrorPropertiesErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.ProviderAuthError],
// [shared.UnknownError],
// [shared.MessageOutputLengthError],
// [shared.MessageAbortedError], [shared.StructuredOutputError],
// [shared.ContentFilterError],
// [shared.ContextOverflowError],
// [shared.APIError].
func (r EventListResponseEventSessionErrorPropertiesError) AsUnion() EventListResponseEventSessionErrorPropertiesErrorUnion {
	return r.union
}

// Union satisfied by [shared.ProviderAuthError], [shared.UnknownError],
// [shared.MessageOutputLengthError],
// [shared.MessageAbortedError], [shared.StructuredOutputError],
// [shared.ContextOverflowError] or
// [shared.APIError].
type EventListResponseEventSessionErrorPropertiesErrorUnion interface {
	ImplementsEventListResponseEventSessionErrorPropertiesError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EventListResponseEventSessionErrorPropertiesErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ProviderAuthError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.UnknownError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.MessageOutputLengthError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.MessageAbortedError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.StructuredOutputError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ContentFilterError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.ContextOverflowError{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.APIError{}),
		},
	)
}

type EventListResponseEventSessionErrorPropertiesErrorName string

const (
	EventListResponseEventSessionErrorPropertiesErrorNameProviderAuthError        EventListResponseEventSessionErrorPropertiesErrorName = "ProviderAuthError"
	EventListResponseEventSessionErrorPropertiesErrorNameUnknownError             EventListResponseEventSessionErrorPropertiesErrorName = "UnknownError"
	EventListResponseEventSessionErrorPropertiesErrorNameMessageOutputLengthError EventListResponseEventSessionErrorPropertiesErrorName = "MessageOutputLengthError"
	EventListResponseEventSessionErrorPropertiesErrorNameMessageAbortedError      EventListResponseEventSessionErrorPropertiesErrorName = "MessageAbortedError"
	EventListResponseEventSessionErrorPropertiesErrorNameStructuredOutputError    EventListResponseEventSessionErrorPropertiesErrorName = "StructuredOutputError"
	EventListResponseEventSessionErrorPropertiesErrorNameContextOverflowError     EventListResponseEventSessionErrorPropertiesErrorName = "ContextOverflowError"
	EventListResponseEventSessionErrorPropertiesErrorNameAPIError                 EventListResponseEventSessionErrorPropertiesErrorName = "APIError"
	EventListResponseEventSessionErrorPropertiesErrorNameContentFilterError       EventListResponseEventSessionErrorPropertiesErrorName = "ContentFilterError"
)

func (r EventListResponseEventSessionErrorPropertiesErrorName) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionErrorPropertiesErrorNameProviderAuthError, EventListResponseEventSessionErrorPropertiesErrorNameUnknownError, EventListResponseEventSessionErrorPropertiesErrorNameMessageOutputLengthError, EventListResponseEventSessionErrorPropertiesErrorNameMessageAbortedError, EventListResponseEventSessionErrorPropertiesErrorNameStructuredOutputError, EventListResponseEventSessionErrorPropertiesErrorNameContextOverflowError, EventListResponseEventSessionErrorPropertiesErrorNameAPIError, EventListResponseEventSessionErrorPropertiesErrorNameContentFilterError:
		return true
	}
	return false
}

type EventListResponseEventSessionErrorType string

const (
	EventListResponseEventSessionErrorTypeSessionError EventListResponseEventSessionErrorType = "session.error"
)

func (r EventListResponseEventSessionErrorType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionErrorTypeSessionError:
		return true
	}
	return false
}

type EventListResponseEventServerConnected struct {
	Properties EventListResponseEventServerConnectedProperties `json:"properties,required"`
	Type       EventListResponseEventServerConnectedType       `json:"type,required"`
	JSON       eventListResponseEventServerConnectedJSON       `json:"-"`
}

// eventListResponseEventServerConnectedJSON contains the JSON metadata for the
// struct [EventListResponseEventServerConnected]
type eventListResponseEventServerConnectedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventServerConnected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventServerConnectedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventServerConnected) implementsEventListResponse() {}

func (r EventListResponseEventServerConnected) implementsGlobalEventPayload() {}

type EventListResponseEventServerConnectedProperties struct {
	JSON eventListResponseEventServerConnectedPropertiesJSON `json:"-"`
}

type eventListResponseEventServerConnectedPropertiesJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventServerConnectedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventServerConnectedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventServerConnectedType string

const (
	EventListResponseEventServerConnectedTypeServerConnected EventListResponseEventServerConnectedType = "server.connected"
)

func (r EventListResponseEventServerConnectedType) IsKnown() bool {
	switch r {
	case EventListResponseEventServerConnectedTypeServerConnected:
		return true
	}
	return false
}

type EventListResponseEventInstallationUpdateAvailable struct {
	Properties EventListResponseEventInstallationUpdateAvailableProperties `json:"properties,required"`
	Type       EventListResponseEventInstallationUpdateAvailableType       `json:"type,required"`
	JSON       eventListResponseEventInstallationUpdateAvailableJSON       `json:"-"`
}

type eventListResponseEventInstallationUpdateAvailableJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventInstallationUpdateAvailable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventInstallationUpdateAvailableJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventInstallationUpdateAvailable) implementsEventListResponse() {}

func (r EventListResponseEventInstallationUpdateAvailable) implementsGlobalEventPayload() {}

type EventListResponseEventInstallationUpdateAvailableProperties struct {
	Version string                                                          `json:"version,required"`
	JSON    eventListResponseEventInstallationUpdateAvailablePropertiesJSON `json:"-"`
}

type eventListResponseEventInstallationUpdateAvailablePropertiesJSON struct {
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventInstallationUpdateAvailableProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventInstallationUpdateAvailablePropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventInstallationUpdateAvailableType string

const (
	EventListResponseEventInstallationUpdateAvailableTypeInstallationUpdateAvailable EventListResponseEventInstallationUpdateAvailableType = "installation.update-available"
)

func (r EventListResponseEventInstallationUpdateAvailableType) IsKnown() bool {
	switch r {
	case EventListResponseEventInstallationUpdateAvailableTypeInstallationUpdateAvailable:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventProjectUpdated
// =============================================================================

type EventListResponseEventProjectUpdated struct {
	Properties EventListResponseEventProjectUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventProjectUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventProjectUpdatedJSON       `json:"-"`
}

type eventListResponseEventProjectUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventProjectUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventProjectUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventProjectUpdated) implementsEventListResponse() {}

func (r EventListResponseEventProjectUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventProjectUpdatedProperties struct {
	Project Project `json:"project,required"`
	JSON    eventListResponseEventProjectUpdatedPropertiesJSON
}

type eventListResponseEventProjectUpdatedPropertiesJSON struct {
	Project     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventProjectUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventProjectUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventProjectUpdatedType string

const (
	EventListResponseEventProjectUpdatedTypeProjectUpdated EventListResponseEventProjectUpdatedType = "project.updated"
)

func (r EventListResponseEventProjectUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventProjectUpdatedTypeProjectUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventServerInstanceDisposed
// =============================================================================

type EventListResponseEventServerInstanceDisposed struct {
	Properties EventListResponseEventServerInstanceDisposedProperties `json:"properties,required"`
	Type       EventListResponseEventServerInstanceDisposedType       `json:"type,required"`
	JSON       eventListResponseEventServerInstanceDisposedJSON       `json:"-"`
}

type eventListResponseEventServerInstanceDisposedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventServerInstanceDisposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventServerInstanceDisposedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventServerInstanceDisposed) implementsEventListResponse() {}

func (r EventListResponseEventServerInstanceDisposed) implementsGlobalEventPayload() {}

type EventListResponseEventServerInstanceDisposedProperties struct {
	Directory string                                                     `json:"directory,required"`
	JSON      eventListResponseEventServerInstanceDisposedPropertiesJSON `json:"-"`
}

type eventListResponseEventServerInstanceDisposedPropertiesJSON struct {
	Directory   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventServerInstanceDisposedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventServerInstanceDisposedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventServerInstanceDisposedType string

const (
	EventListResponseEventServerInstanceDisposedTypeServerInstanceDisposed EventListResponseEventServerInstanceDisposedType = "server.instance.disposed"
)

func (r EventListResponseEventServerInstanceDisposedType) IsKnown() bool {
	switch r {
	case EventListResponseEventServerInstanceDisposedTypeServerInstanceDisposed:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventGlobalDisposed
// =============================================================================

type EventListResponseEventGlobalDisposed struct {
	Properties EventListResponseEventGlobalDisposedProperties `json:"properties,required"`
	Type       EventListResponseEventGlobalDisposedType       `json:"type,required"`
	JSON       eventListResponseEventGlobalDisposedJSON       `json:"-"`
}

type eventListResponseEventGlobalDisposedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventGlobalDisposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventGlobalDisposedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventGlobalDisposed) implementsEventListResponse() {}

func (r EventListResponseEventGlobalDisposed) implementsGlobalEventPayload() {}

type EventListResponseEventGlobalDisposedProperties struct {
	JSON eventListResponseEventGlobalDisposedPropertiesJSON `json:"-"`
}

type eventListResponseEventGlobalDisposedPropertiesJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventGlobalDisposedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventGlobalDisposedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventGlobalDisposedType string

const (
	EventListResponseEventGlobalDisposedTypeGlobalDisposed EventListResponseEventGlobalDisposedType = "global.disposed"
)

func (r EventListResponseEventGlobalDisposedType) IsKnown() bool {
	switch r {
	case EventListResponseEventGlobalDisposedTypeGlobalDisposed:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventLspUpdated
// =============================================================================

type EventListResponseEventLspUpdated struct {
	Properties EventListResponseEventLspUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventLspUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventLspUpdatedJSON       `json:"-"`
}

type eventListResponseEventLspUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventLspUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventLspUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventLspUpdated) implementsEventListResponse() {}

func (r EventListResponseEventLspUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventLspUpdatedProperties struct {
	JSON eventListResponseEventLspUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventLspUpdatedPropertiesJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventLspUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventLspUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventLspUpdatedType string

const (
	EventListResponseEventLspUpdatedTypeLspUpdated EventListResponseEventLspUpdatedType = "lsp.updated"
)

func (r EventListResponseEventLspUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventLspUpdatedTypeLspUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventMessagePartDelta
// =============================================================================

type EventListResponseEventMessagePartDelta struct {
	Properties EventListResponseEventMessagePartDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventMessagePartDeltaType       `json:"type,required"`
	JSON       eventListResponseEventMessagePartDeltaJSON       `json:"-"`
}

type eventListResponseEventMessagePartDeltaJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessagePartDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessagePartDeltaJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMessagePartDelta) implementsEventListResponse() {}

func (r EventListResponseEventMessagePartDelta) implementsGlobalEventPayload() {}

type EventListResponseEventMessagePartDeltaProperties struct {
	Delta     string `json:"delta,required"`
	Field     string `json:"field,required"`
	MessageID string `json:"messageID,required"`
	PartID    string `json:"partID,required"`
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventMessagePartDeltaPropertiesJSON
}

type eventListResponseEventMessagePartDeltaPropertiesJSON struct {
	Delta       apijson.Field
	Field       apijson.Field
	MessageID   apijson.Field
	PartID      apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessagePartDeltaProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessagePartDeltaPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMessagePartDeltaType string

const (
	EventListResponseEventMessagePartDeltaTypeMessagePartDelta EventListResponseEventMessagePartDeltaType = "message.part.delta"
)

func (r EventListResponseEventMessagePartDeltaType) IsKnown() bool {
	switch r {
	case EventListResponseEventMessagePartDeltaTypeMessagePartDelta:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPermissionAsked
// =============================================================================

type EventListResponseEventPermissionAsked struct {
	Properties EventListResponseEventPermissionAskedProperties `json:"properties,required"`
	Type       EventListResponseEventPermissionAskedType       `json:"type,required"`
	JSON       eventListResponseEventPermissionAskedJSON       `json:"-"`
}

type eventListResponseEventPermissionAskedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionAsked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionAskedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPermissionAsked) implementsEventListResponse() {}

func (r EventListResponseEventPermissionAsked) implementsGlobalEventPayload() {}

type EventListResponseEventPermissionAskedProperties struct {
	Always     []string               `json:"always,required"`
	ID         string                 `json:"id,required"`
	Metadata   map[string]interface{} `json:"metadata,required"`
	Patterns   []string               `json:"patterns,required"`
	Permission string                 `json:"permission,required"`
	SessionID  string                 `json:"sessionID,required"`
	// This field can have the runtime type of [EventListResponseEventPermissionAskedPropertiesTool].
	Tool interface{}                                         `json:"tool"`
	JSON eventListResponseEventPermissionAskedPropertiesJSON `json:"-"`
}

type eventListResponseEventPermissionAskedPropertiesJSON struct {
	Always      apijson.Field
	ID          apijson.Field
	Metadata    apijson.Field
	Patterns    apijson.Field
	Permission  apijson.Field
	SessionID   apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionAskedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionAskedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPermissionAskedPropertiesTool struct {
	CallID    string `json:"callID,required"`
	MessageID string `json:"messageID,required"`
	JSON      eventListResponseEventPermissionAskedPropertiesToolJSON
}

type eventListResponseEventPermissionAskedPropertiesToolJSON struct {
	CallID      apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionAskedPropertiesTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionAskedPropertiesToolJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPermissionAskedType string

const (
	EventListResponseEventPermissionAskedTypePermissionAsked EventListResponseEventPermissionAskedType = "permission.asked"
)

func (r EventListResponseEventPermissionAskedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPermissionAskedTypePermissionAsked:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionStatus
// =============================================================================

type EventListResponseEventSessionStatus struct {
	Properties EventListResponseEventSessionStatusProperties `json:"properties,required"`
	Type       EventListResponseEventSessionStatusType       `json:"type,required"`
	JSON       eventListResponseEventSessionStatusJSON       `json:"-"`
}

type eventListResponseEventSessionStatusJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionStatusJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionStatus) implementsEventListResponse() {}

func (r EventListResponseEventSessionStatus) implementsGlobalEventPayload() {}

type EventListResponseEventSessionStatusProperties struct {
	SessionID string                                    `json:"sessionID,required"`
	Status    EventListResponseEventSessionStatusStatus `json:"status,required"`
	JSON      eventListResponseEventSessionStatusPropertiesJSON
}

type eventListResponseEventSessionStatusPropertiesJSON struct {
	SessionID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionStatusProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionStatusPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionStatusStatus struct {
	Attempt *int    `json:"attempt,omitempty"`
	Message *string `json:"message,omitempty"`
	Next    *int    `json:"next,omitempty"`
	Type    string  `json:"type,required"`
	JSON    eventListResponseEventSessionStatusStatusJSON
}

type eventListResponseEventSessionStatusStatusJSON struct {
	Attempt     apijson.Field
	Message     apijson.Field
	Next        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionStatusStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionStatusStatusJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionStatusType string

const (
	EventListResponseEventSessionStatusTypeSessionStatus EventListResponseEventSessionStatusType = "session.status"
)

func (r EventListResponseEventSessionStatusType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionStatusTypeSessionStatus:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventQuestionAsked
// =============================================================================

type EventListResponseEventQuestionAsked struct {
	Properties EventListResponseEventQuestionAskedProperties `json:"properties,required"`
	Type       EventListResponseEventQuestionAskedType       `json:"type,required"`
	JSON       eventListResponseEventQuestionAskedJSON       `json:"-"`
}

type eventListResponseEventQuestionAskedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAsked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionAsked) implementsEventListResponse() {}

func (r EventListResponseEventQuestionAsked) implementsGlobalEventPayload() {}

type EventListResponseEventQuestionAskedProperties struct {
	ID        string                                                   `json:"id,required"`
	Questions []EventListResponseEventQuestionAskedPropertiesQuestions `json:"questions,required"`
	SessionID string                                                   `json:"sessionID,required"`
	// This field can have the runtime type of [EventListResponseEventQuestionAskedPropertiesTool].
	Tool interface{}                                       `json:"tool"`
	JSON eventListResponseEventQuestionAskedPropertiesJSON `json:"-"`
}

type eventListResponseEventQuestionAskedPropertiesJSON struct {
	ID          apijson.Field
	Questions   apijson.Field
	SessionID   apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAskedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionAskedPropertiesQuestions struct {
	Custom   *bool                                                           `json:"custom,omitempty"`
	Header   string                                                          `json:"header,required"`
	Multiple *bool                                                           `json:"multiple,omitempty"`
	Options  []EventListResponseEventQuestionAskedPropertiesQuestionsOptions `json:"options,required"`
	Question string                                                          `json:"question,required"`
	JSON     eventListResponseEventQuestionAskedPropertiesQuestionsJSON
}

type eventListResponseEventQuestionAskedPropertiesQuestionsJSON struct {
	Custom      apijson.Field
	Header      apijson.Field
	Multiple    apijson.Field
	Options     apijson.Field
	Question    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAskedPropertiesQuestions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedPropertiesQuestionsJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionAskedPropertiesQuestionsOptions struct {
	Description string `json:"description,required"`
	Label       string `json:"label,required"`
	JSON        eventListResponseEventQuestionAskedPropertiesQuestionsOptionsJSON
}

type eventListResponseEventQuestionAskedPropertiesQuestionsOptionsJSON struct {
	Description apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAskedPropertiesQuestionsOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedPropertiesQuestionsOptionsJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionAskedPropertiesTool struct {
	CallID    string `json:"callID,required"`
	MessageID string `json:"messageID,required"`
	JSON      eventListResponseEventQuestionAskedPropertiesToolJSON
}

type eventListResponseEventQuestionAskedPropertiesToolJSON struct {
	CallID      apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAskedPropertiesTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedPropertiesToolJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionAskedType string

const (
	EventListResponseEventQuestionAskedTypeQuestionAsked EventListResponseEventQuestionAskedType = "question.asked"
)

func (r EventListResponseEventQuestionAskedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionAskedTypeQuestionAsked:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventQuestionRejected
// =============================================================================

type EventListResponseEventQuestionRejected struct {
	Properties EventListResponseEventQuestionRejectedProperties `json:"properties,required"`
	Type       EventListResponseEventQuestionRejectedType       `json:"type,required"`
	JSON       eventListResponseEventQuestionRejectedJSON       `json:"-"`
}

type eventListResponseEventQuestionRejectedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionRejected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRejectedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionRejected) implementsEventListResponse() {}

func (r EventListResponseEventQuestionRejected) implementsGlobalEventPayload() {}

type EventListResponseEventQuestionRejectedProperties struct {
	RequestID string `json:"requestID,required"`
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventQuestionRejectedPropertiesJSON
}

type eventListResponseEventQuestionRejectedPropertiesJSON struct {
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionRejectedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRejectedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionRejectedType string

const (
	EventListResponseEventQuestionRejectedTypeQuestionRejected EventListResponseEventQuestionRejectedType = "question.rejected"
)

func (r EventListResponseEventQuestionRejectedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionRejectedTypeQuestionRejected:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventQuestionReplied
// =============================================================================

type EventListResponseEventQuestionReplied struct {
	Properties EventListResponseEventQuestionRepliedProperties `json:"properties,required"`
	Type       EventListResponseEventQuestionRepliedType       `json:"type,required"`
	JSON       eventListResponseEventQuestionRepliedJSON       `json:"-"`
}

type eventListResponseEventQuestionRepliedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionReplied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRepliedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionReplied) implementsEventListResponse() {}

func (r EventListResponseEventQuestionReplied) implementsGlobalEventPayload() {}

type EventListResponseEventQuestionRepliedProperties struct {
	Answers   [][]string `json:"answers,required"`
	RequestID string     `json:"requestID,required"`
	SessionID string     `json:"sessionID,required"`
	JSON      eventListResponseEventQuestionRepliedPropertiesJSON
}

type eventListResponseEventQuestionRepliedPropertiesJSON struct {
	Answers     apijson.Field
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionRepliedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRepliedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionRepliedType string

const (
	EventListResponseEventQuestionRepliedTypeQuestionReplied EventListResponseEventQuestionRepliedType = "question.replied"
)

func (r EventListResponseEventQuestionRepliedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionRepliedTypeQuestionReplied:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventTuiPromptAppend
// =============================================================================

type EventListResponseEventTuiPromptAppend struct {
	Properties EventListResponseEventTuiPromptAppendProperties `json:"properties,required"`
	Type       EventListResponseEventTuiPromptAppendType       `json:"type,required"`
	JSON       eventListResponseEventTuiPromptAppendJSON       `json:"-"`
}

type eventListResponseEventTuiPromptAppendJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiPromptAppend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiPromptAppendJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiPromptAppend) implementsEventListResponse() {}

func (r EventListResponseEventTuiPromptAppend) implementsGlobalEventPayload() {}

type EventListResponseEventTuiPromptAppendProperties struct {
	Text string `json:"text,required"`
	JSON eventListResponseEventTuiPromptAppendPropertiesJSON
}

type eventListResponseEventTuiPromptAppendPropertiesJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiPromptAppendProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiPromptAppendPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiPromptAppendType string

const (
	EventListResponseEventTuiPromptAppendTypeTuiPromptAppend EventListResponseEventTuiPromptAppendType = "tui.prompt.append"
)

func (r EventListResponseEventTuiPromptAppendType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiPromptAppendTypeTuiPromptAppend:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventTuiCommandExecute
// =============================================================================

type EventListResponseEventTuiCommandExecute struct {
	Properties EventListResponseEventTuiCommandExecuteProperties `json:"properties,required"`
	Type       EventListResponseEventTuiCommandExecuteType       `json:"type,required"`
	JSON       eventListResponseEventTuiCommandExecuteJSON       `json:"-"`
}

type eventListResponseEventTuiCommandExecuteJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiCommandExecute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiCommandExecuteJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiCommandExecute) implementsEventListResponse() {}

func (r EventListResponseEventTuiCommandExecute) implementsGlobalEventPayload() {}

type EventListResponseEventTuiCommandExecuteProperties struct {
	Command string `json:"command,required"`
	JSON    eventListResponseEventTuiCommandExecutePropertiesJSON
}

type eventListResponseEventTuiCommandExecutePropertiesJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiCommandExecuteProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiCommandExecutePropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiCommandExecuteType string

const (
	EventListResponseEventTuiCommandExecuteTypeTuiCommandExecute EventListResponseEventTuiCommandExecuteType = "tui.command.execute"
)

func (r EventListResponseEventTuiCommandExecuteType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiCommandExecuteTypeTuiCommandExecute:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventTuiToastShow
// =============================================================================

type EventListResponseEventTuiToastShow struct {
	Properties EventListResponseEventTuiToastShowProperties `json:"properties,required"`
	Type       EventListResponseEventTuiToastShowType       `json:"type,required"`
	JSON       eventListResponseEventTuiToastShowJSON       `json:"-"`
}

type eventListResponseEventTuiToastShowJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiToastShow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiToastShowJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiToastShow) implementsEventListResponse() {}

func (r EventListResponseEventTuiToastShow) implementsGlobalEventPayload() {}

type EventListResponseEventTuiToastShowProperties struct {
	Duration *int    `json:"duration,omitempty"`
	Message  string  `json:"message,required"`
	Title    *string `json:"title,omitempty"`
	Variant  string  `json:"variant,required"`
	JSON     eventListResponseEventTuiToastShowPropertiesJSON
}

type eventListResponseEventTuiToastShowPropertiesJSON struct {
	Duration    apijson.Field
	Message     apijson.Field
	Title       apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiToastShowProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiToastShowPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiToastShowType string

const (
	EventListResponseEventTuiToastShowTypeTuiToastShow EventListResponseEventTuiToastShowType = "tui.toast.show"
)

func (r EventListResponseEventTuiToastShowType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiToastShowTypeTuiToastShow:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventTuiSessionSelect
// =============================================================================

type EventListResponseEventTuiSessionSelect struct {
	Properties EventListResponseEventTuiSessionSelectProperties `json:"properties,required"`
	Type       EventListResponseEventTuiSessionSelectType       `json:"type,required"`
	JSON       eventListResponseEventTuiSessionSelectJSON       `json:"-"`
}

type eventListResponseEventTuiSessionSelectJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiSessionSelect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiSessionSelectJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiSessionSelect) implementsEventListResponse() {}

func (r EventListResponseEventTuiSessionSelect) implementsGlobalEventPayload() {}

type EventListResponseEventTuiSessionSelectProperties struct {
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventTuiSessionSelectPropertiesJSON
}

type eventListResponseEventTuiSessionSelectPropertiesJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiSessionSelectProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiSessionSelectPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiSessionSelectType string

const (
	EventListResponseEventTuiSessionSelectTypeTuiSessionSelect EventListResponseEventTuiSessionSelectType = "tui.session.select"
)

func (r EventListResponseEventTuiSessionSelectType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiSessionSelectTypeTuiSessionSelect:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventMcpToolsChanged
// =============================================================================

type EventListResponseEventMcpToolsChanged struct {
	Properties EventListResponseEventMcpToolsChangedProperties `json:"properties,required"`
	Type       EventListResponseEventMcpToolsChangedType       `json:"type,required"`
	JSON       eventListResponseEventMcpToolsChangedJSON       `json:"-"`
}

type eventListResponseEventMcpToolsChangedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpToolsChanged) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpToolsChangedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMcpToolsChanged) implementsEventListResponse() {}

func (r EventListResponseEventMcpToolsChanged) implementsGlobalEventPayload() {}

type EventListResponseEventMcpToolsChangedProperties struct {
	Server string `json:"server,required"`
	JSON   eventListResponseEventMcpToolsChangedPropertiesJSON
}

type eventListResponseEventMcpToolsChangedPropertiesJSON struct {
	Server      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpToolsChangedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpToolsChangedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMcpToolsChangedType string

const (
	EventListResponseEventMcpToolsChangedTypeMcpToolsChanged EventListResponseEventMcpToolsChangedType = "mcp.tools.changed"
)

func (r EventListResponseEventMcpToolsChangedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMcpToolsChangedTypeMcpToolsChanged:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventMcpBrowserOpenFailed
// =============================================================================

type EventListResponseEventMcpBrowserOpenFailed struct {
	Properties EventListResponseEventMcpBrowserOpenFailedProperties `json:"properties,required"`
	Type       EventListResponseEventMcpBrowserOpenFailedType       `json:"type,required"`
	JSON       eventListResponseEventMcpBrowserOpenFailedJSON       `json:"-"`
}

type eventListResponseEventMcpBrowserOpenFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpBrowserOpenFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpBrowserOpenFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMcpBrowserOpenFailed) implementsEventListResponse() {}

func (r EventListResponseEventMcpBrowserOpenFailed) implementsGlobalEventPayload() {}

type EventListResponseEventMcpBrowserOpenFailedProperties struct {
	McpName string `json:"mcpName,required"`
	URL     string `json:"url,required"`
	JSON    eventListResponseEventMcpBrowserOpenFailedPropertiesJSON
}

type eventListResponseEventMcpBrowserOpenFailedPropertiesJSON struct {
	McpName     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpBrowserOpenFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpBrowserOpenFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMcpBrowserOpenFailedType string

const (
	EventListResponseEventMcpBrowserOpenFailedTypeMcpBrowserOpenFailed EventListResponseEventMcpBrowserOpenFailedType = "mcp.browser.open.failed"
)

func (r EventListResponseEventMcpBrowserOpenFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMcpBrowserOpenFailedTypeMcpBrowserOpenFailed:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventCommandExecuted
// =============================================================================

type EventListResponseEventCommandExecuted struct {
	Properties EventListResponseEventCommandExecutedProperties `json:"properties,required"`
	Type       EventListResponseEventCommandExecutedType       `json:"type,required"`
	JSON       eventListResponseEventCommandExecutedJSON       `json:"-"`
}

type eventListResponseEventCommandExecutedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCommandExecuted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCommandExecutedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventCommandExecuted) implementsEventListResponse() {}

func (r EventListResponseEventCommandExecuted) implementsGlobalEventPayload() {}

type EventListResponseEventCommandExecutedProperties struct {
	Arguments string `json:"arguments,required"`
	MessageID string `json:"messageID,required"`
	Name      string `json:"name,required"`
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventCommandExecutedPropertiesJSON
}

type eventListResponseEventCommandExecutedPropertiesJSON struct {
	Arguments   apijson.Field
	MessageID   apijson.Field
	Name        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCommandExecutedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCommandExecutedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventCommandExecutedType string

const (
	EventListResponseEventCommandExecutedTypeCommandExecuted EventListResponseEventCommandExecutedType = "command.executed"
)

func (r EventListResponseEventCommandExecutedType) IsKnown() bool {
	switch r {
	case EventListResponseEventCommandExecutedTypeCommandExecuted:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionDiff
// =============================================================================

type EventListResponseEventSessionDiff struct {
	Properties EventListResponseEventSessionDiffProperties `json:"properties,required"`
	Type       EventListResponseEventSessionDiffType       `json:"type,required"`
	JSON       eventListResponseEventSessionDiffJSON       `json:"-"`
}

type eventListResponseEventSessionDiffJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionDiffJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionDiff) implementsEventListResponse() {}

func (r EventListResponseEventSessionDiff) implementsGlobalEventPayload() {}

type EventListResponseEventSessionDiffProperties struct {
	Diff      []EventListResponseEventSessionDiffPropertiesDiff `json:"diff,required"`
	SessionID string                                            `json:"sessionID,required"`
	JSON      eventListResponseEventSessionDiffPropertiesJSON
}

type eventListResponseEventSessionDiffPropertiesJSON struct {
	Diff        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionDiffProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionDiffPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionDiffPropertiesDiff struct {
	Additions int    `json:"additions,required"`
	Deletions int    `json:"deletions,required"`
	File      string `json:"file"`
	Patch     string `json:"patch"`
	Status    string `json:"status,omitempty"`
	JSON      eventListResponseEventSessionDiffPropertiesDiffJSON
}

type eventListResponseEventSessionDiffPropertiesDiffJSON struct {
	Additions   apijson.Field
	Deletions   apijson.Field
	File        apijson.Field
	Patch       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionDiffPropertiesDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionDiffPropertiesDiffJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionDiffType string

const (
	EventListResponseEventSessionDiffTypeSessionDiff EventListResponseEventSessionDiffType = "session.diff"
)

func (r EventListResponseEventSessionDiffType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionDiffTypeSessionDiff:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventVcsBranchUpdated
// =============================================================================

type EventListResponseEventVcsBranchUpdated struct {
	Properties EventListResponseEventVcsBranchUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventVcsBranchUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventVcsBranchUpdatedJSON       `json:"-"`
}

type eventListResponseEventVcsBranchUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventVcsBranchUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventVcsBranchUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventVcsBranchUpdated) implementsEventListResponse() {}

func (r EventListResponseEventVcsBranchUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventVcsBranchUpdatedProperties struct {
	Branch string `json:"branch"`
	JSON   eventListResponseEventVcsBranchUpdatedPropertiesJSON
}

type eventListResponseEventVcsBranchUpdatedPropertiesJSON struct {
	Branch      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventVcsBranchUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventVcsBranchUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventVcsBranchUpdatedType string

const (
	EventListResponseEventVcsBranchUpdatedTypeVcsBranchUpdated EventListResponseEventVcsBranchUpdatedType = "vcs.branch.updated"
)

func (r EventListResponseEventVcsBranchUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventVcsBranchUpdatedTypeVcsBranchUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventWorkspaceReady
// =============================================================================

type EventListResponseEventWorkspaceReady struct {
	Properties EventListResponseEventWorkspaceReadyProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceReadyType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceReadyJSON       `json:"-"`
}

type eventListResponseEventWorkspaceReadyJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceReady) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceReadyJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceReady) implementsEventListResponse() {}

func (r EventListResponseEventWorkspaceReady) implementsGlobalEventPayload() {}

type EventListResponseEventWorkspaceReadyProperties struct {
	Name string `json:"name,required"`
	JSON eventListResponseEventWorkspaceReadyPropertiesJSON
}

type eventListResponseEventWorkspaceReadyPropertiesJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceReadyProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceReadyPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceReadyType string

const (
	EventListResponseEventWorkspaceReadyTypeWorkspaceReady EventListResponseEventWorkspaceReadyType = "workspace.ready"
)

func (r EventListResponseEventWorkspaceReadyType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceReadyTypeWorkspaceReady:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventWorkspaceFailed
// =============================================================================

type EventListResponseEventWorkspaceFailed struct {
	Properties EventListResponseEventWorkspaceFailedProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceFailedType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceFailedJSON       `json:"-"`
}

type eventListResponseEventWorkspaceFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceFailed) implementsEventListResponse() {}

func (r EventListResponseEventWorkspaceFailed) implementsGlobalEventPayload() {}

type EventListResponseEventWorkspaceFailedProperties struct {
	Message string `json:"message,required"`
	JSON    eventListResponseEventWorkspaceFailedPropertiesJSON
}

type eventListResponseEventWorkspaceFailedPropertiesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceFailedType string

const (
	EventListResponseEventWorkspaceFailedTypeWorkspaceFailed EventListResponseEventWorkspaceFailedType = "workspace.failed"
)

func (r EventListResponseEventWorkspaceFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceFailedTypeWorkspaceFailed:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPtyCreated
// =============================================================================

type EventListResponseEventPtyCreated struct {
	Properties EventListResponseEventPtyCreatedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyCreatedType       `json:"type,required"`
	JSON       eventListResponseEventPtyCreatedJSON       `json:"-"`
}

type eventListResponseEventPtyCreatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyCreatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyCreated) implementsEventListResponse() {}

func (r EventListResponseEventPtyCreated) implementsGlobalEventPayload() {}

type EventListResponseEventPtyCreatedProperties struct {
	Info Pty `json:"info,required"`
	JSON eventListResponseEventPtyCreatedPropertiesJSON
}

type eventListResponseEventPtyCreatedPropertiesJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyCreatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyCreatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyCreatedType string

const (
	EventListResponseEventPtyCreatedTypePtyCreated EventListResponseEventPtyCreatedType = "pty.created"
)

func (r EventListResponseEventPtyCreatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyCreatedTypePtyCreated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPtyUpdated
// =============================================================================

type EventListResponseEventPtyUpdated struct {
	Properties EventListResponseEventPtyUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventPtyUpdatedJSON       `json:"-"`
}

type eventListResponseEventPtyUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyUpdated) implementsEventListResponse() {}

func (r EventListResponseEventPtyUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventPtyUpdatedProperties struct {
	Info Pty `json:"info,required"`
	JSON eventListResponseEventPtyUpdatedPropertiesJSON
}

type eventListResponseEventPtyUpdatedPropertiesJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyUpdatedType string

const (
	EventListResponseEventPtyUpdatedTypePtyUpdated EventListResponseEventPtyUpdatedType = "pty.updated"
)

func (r EventListResponseEventPtyUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyUpdatedTypePtyUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPtyExited
// =============================================================================

type EventListResponseEventPtyExited struct {
	Properties EventListResponseEventPtyExitedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyExitedType       `json:"type,required"`
	JSON       eventListResponseEventPtyExitedJSON       `json:"-"`
}

type eventListResponseEventPtyExitedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyExited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyExitedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyExited) implementsEventListResponse() {}

func (r EventListResponseEventPtyExited) implementsGlobalEventPayload() {}

type EventListResponseEventPtyExitedProperties struct {
	ExitCode int    `json:"exitCode,required"`
	ID       string `json:"id,required"`
	JSON     eventListResponseEventPtyExitedPropertiesJSON
}

type eventListResponseEventPtyExitedPropertiesJSON struct {
	ExitCode    apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyExitedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyExitedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyExitedType string

const (
	EventListResponseEventPtyExitedTypePtyExited EventListResponseEventPtyExitedType = "pty.exited"
)

func (r EventListResponseEventPtyExitedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyExitedTypePtyExited:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPtyDeleted
// =============================================================================

type EventListResponseEventPtyDeleted struct {
	Properties EventListResponseEventPtyDeletedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyDeletedType       `json:"type,required"`
	JSON       eventListResponseEventPtyDeletedJSON       `json:"-"`
}

type eventListResponseEventPtyDeletedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyDeletedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyDeleted) implementsEventListResponse() {}

func (r EventListResponseEventPtyDeleted) implementsGlobalEventPayload() {}

type EventListResponseEventPtyDeletedProperties struct {
	ID   string `json:"id,required"`
	JSON eventListResponseEventPtyDeletedPropertiesJSON
}

type eventListResponseEventPtyDeletedPropertiesJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyDeletedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyDeletedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyDeletedType string

const (
	EventListResponseEventPtyDeletedTypePtyDeleted EventListResponseEventPtyDeletedType = "pty.deleted"
)

func (r EventListResponseEventPtyDeletedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyDeletedTypePtyDeleted:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventWorktreeReady
// =============================================================================

type EventListResponseEventWorktreeReady struct {
	Properties EventListResponseEventWorktreeReadyProperties `json:"properties,required"`
	Type       EventListResponseEventWorktreeReadyType       `json:"type,required"`
	JSON       eventListResponseEventWorktreeReadyJSON       `json:"-"`
}

type eventListResponseEventWorktreeReadyJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeReady) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeReadyJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorktreeReady) implementsEventListResponse() {}

func (r EventListResponseEventWorktreeReady) implementsGlobalEventPayload() {}

type EventListResponseEventWorktreeReadyProperties struct {
	Branch string `json:"branch"`
	Name   string `json:"name,required"`
	JSON   eventListResponseEventWorktreeReadyPropertiesJSON
}

type eventListResponseEventWorktreeReadyPropertiesJSON struct {
	Branch      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeReadyProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeReadyPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorktreeReadyType string

const (
	EventListResponseEventWorktreeReadyTypeWorktreeReady EventListResponseEventWorktreeReadyType = "worktree.ready"
)

func (r EventListResponseEventWorktreeReadyType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorktreeReadyTypeWorktreeReady:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventWorktreeFailed
// =============================================================================

type EventListResponseEventWorktreeFailed struct {
	Properties EventListResponseEventWorktreeFailedProperties `json:"properties,required"`
	Type       EventListResponseEventWorktreeFailedType       `json:"type,required"`
	JSON       eventListResponseEventWorktreeFailedJSON       `json:"-"`
}

type eventListResponseEventWorktreeFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorktreeFailed) implementsEventListResponse() {}

func (r EventListResponseEventWorktreeFailed) implementsGlobalEventPayload() {}

type EventListResponseEventWorktreeFailedProperties struct {
	Message string `json:"message,required"`
	JSON    eventListResponseEventWorktreeFailedPropertiesJSON
}

type eventListResponseEventWorktreeFailedPropertiesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorktreeFailedType string

const (
	EventListResponseEventWorktreeFailedTypeWorktreeFailed EventListResponseEventWorktreeFailedType = "worktree.failed"
)

func (r EventListResponseEventWorktreeFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorktreeFailedTypeWorktreeFailed:
		return true
	}
	return false
}

type Pty struct {
	Args    []string `json:"args,required"`
	Command string   `json:"command,required"`
	Cwd     string   `json:"cwd,required"`
	// ExitCode is the exit code of the PTY process if it has exited.
	ExitCode int64     `json:"exitCode"`
	ID       string    `json:"id,required"`
	Pid      int64     `json:"pid,required"`
	Status   PtyStatus `json:"status,required"`
	Title    string    `json:"title,required"`
	JSON     ptyJSON   `json:"-"`
}

type ptyJSON struct {
	Args        apijson.Field
	Command     apijson.Field
	Cwd         apijson.Field
	ExitCode    apijson.Field
	ID          apijson.Field
	Pid         apijson.Field
	Status      apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyJSON) RawJSON() string {
	return r.raw
}

type PtyStatus string

const (
	PtyStatusRunning PtyStatus = "running"
	PtyStatusExited  PtyStatus = "exited"
)

func (r PtyStatus) IsKnown() bool {
	switch r {
	case PtyStatusRunning, PtyStatusExited:
		return true
	}
	return false
}

// EventListResponseEventWorkspaceStatus
// =============================================================================

type EventListResponseEventWorkspaceStatus struct {
	Properties EventListResponseEventWorkspaceStatusProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceStatusType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceStatusJSON       `json:"-"`
}

type eventListResponseEventWorkspaceStatusJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceStatusJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceStatus) implementsEventListResponse() {}

func (r EventListResponseEventWorkspaceStatus) implementsGlobalEventPayload() {}

type EventListResponseEventWorkspaceStatusProperties struct {
	WorkspaceID string                                      `json:"workspaceID,required"`
	Status      EventListResponseEventWorkspaceStatusStatus `json:"status,required"`
	JSON        eventListResponseEventWorkspaceStatusPropertiesJSON
}

type eventListResponseEventWorkspaceStatusPropertiesJSON struct {
	WorkspaceID apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceStatusProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceStatusPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceStatusType string

const (
	EventListResponseEventWorkspaceStatusTypeWorkspaceStatus EventListResponseEventWorkspaceStatusType = "workspace.status"
)

func (r EventListResponseEventWorkspaceStatusType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceStatusTypeWorkspaceStatus:
		return true
	}
	return false
}

type EventListResponseEventWorkspaceStatusStatus string

const (
	EventListResponseEventWorkspaceStatusStatusConnected    EventListResponseEventWorkspaceStatusStatus = "connected"
	EventListResponseEventWorkspaceStatusStatusConnecting   EventListResponseEventWorkspaceStatusStatus = "connecting"
	EventListResponseEventWorkspaceStatusStatusDisconnected EventListResponseEventWorkspaceStatusStatus = "disconnected"
	EventListResponseEventWorkspaceStatusStatusError        EventListResponseEventWorkspaceStatusStatus = "error"
)

func (r EventListResponseEventWorkspaceStatusStatus) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceStatusStatusConnected, EventListResponseEventWorkspaceStatusStatusConnecting, EventListResponseEventWorkspaceStatusStatusDisconnected, EventListResponseEventWorkspaceStatusStatusError:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventModelsDevRefreshed
// =============================================================================

type EventListResponseEventModelsDevRefreshed struct {
	Properties EventListResponseEventModelsDevRefreshedProperties `json:"properties,required"`
	Type       EventListResponseEventModelsDevRefreshedType       `json:"type,required"`
	JSON       eventListResponseEventModelsDevRefreshedJSON       `json:"-"`
}

type eventListResponseEventModelsDevRefreshedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventModelsDevRefreshed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventModelsDevRefreshedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventModelsDevRefreshed) implementsEventListResponse() {}

func (r EventListResponseEventModelsDevRefreshed) implementsGlobalEventPayload() {}

type EventListResponseEventModelsDevRefreshedProperties struct {
	JSON eventListResponseEventModelsDevRefreshedPropertiesJSON `json:"-"`
}

type eventListResponseEventModelsDevRefreshedPropertiesJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventModelsDevRefreshedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventModelsDevRefreshedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventModelsDevRefreshedType string

const (
	EventListResponseEventModelsDevRefreshedTypeModelsDevRefreshed EventListResponseEventModelsDevRefreshedType = "models-dev.refreshed"
)

func (r EventListResponseEventModelsDevRefreshedType) IsKnown() bool {
	switch r {
	case EventListResponseEventModelsDevRefreshedTypeModelsDevRefreshed:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionNextAgentSwitched
// =============================================================================

type EventListResponseType string

const (
	EventListResponseTypeCommandExecuted              EventListResponseType = "command.executed"
	EventListResponseTypeFileEdited                   EventListResponseType = "file.edited"
	EventListResponseTypeFileWatcherUpdated           EventListResponseType = "file.watcher.updated"
	EventListResponseTypeGlobalDisposed               EventListResponseType = "global.disposed"
	EventListResponseTypeInstallationUpdateAvailable  EventListResponseType = "installation.update-available"
	EventListResponseTypeInstallationUpdated          EventListResponseType = "installation.updated"
	EventListResponseTypeLspUpdated                   EventListResponseType = "lsp.updated"
	EventListResponseTypeMcpBrowserOpenFailed         EventListResponseType = "mcp.browser.open.failed"
	EventListResponseTypeMcpToolsChanged              EventListResponseType = "mcp.tools.changed"
	EventListResponseTypeModelsDevRefreshed           EventListResponseType = "models-dev.refreshed"
	EventListResponseTypeMessagePartDelta             EventListResponseType = "message.part.delta"
	EventListResponseTypeMessagePartRemoved           EventListResponseType = "message.part.removed"
	EventListResponseTypeMessagePartUpdated           EventListResponseType = "message.part.updated"
	EventListResponseTypeMessageRemoved               EventListResponseType = "message.removed"
	EventListResponseTypeMessageUpdated               EventListResponseType = "message.updated"
	EventListResponseTypePermissionAsked              EventListResponseType = "permission.asked"
	EventListResponseTypePermissionReplied            EventListResponseType = "permission.replied"
	EventListResponseTypePluginAdded                  EventListResponseType = "plugin.added"
	EventListResponseTypeProjectUpdated               EventListResponseType = "project.updated"
	EventListResponseTypePtyCreated                   EventListResponseType = "pty.created"
	EventListResponseTypePtyDeleted                   EventListResponseType = "pty.deleted"
	EventListResponseTypePtyExited                    EventListResponseType = "pty.exited"
	EventListResponseTypePtyUpdated                   EventListResponseType = "pty.updated"
	EventListResponseTypeQuestionAsked                EventListResponseType = "question.asked"
	EventListResponseTypeQuestionRejected             EventListResponseType = "question.rejected"
	EventListResponseTypeQuestionReplied              EventListResponseType = "question.replied"
	EventListResponseTypeServerConnected              EventListResponseType = "server.connected"
	EventListResponseTypeServerInstanceDisposed       EventListResponseType = "server.instance.disposed"
	EventListResponseTypeSessionCompacted             EventListResponseType = "session.compacted"
	EventListResponseTypeSessionCreated               EventListResponseType = "session.created"
	EventListResponseTypeSessionDeleted               EventListResponseType = "session.deleted"
	EventListResponseTypeSessionDiff                  EventListResponseType = "session.diff"
	EventListResponseTypeSessionError                 EventListResponseType = "session.error"
	EventListResponseTypeSessionIdle                  EventListResponseType = "session.idle"
	EventListResponseTypeSessionNextAgentSwitched     EventListResponseType = "session.next.agent.switched"
	EventListResponseTypeSessionNextCompactionDelta   EventListResponseType = "session.next.compaction.delta"
	EventListResponseTypeSessionNextCompactionEnded   EventListResponseType = "session.next.compaction.ended"
	EventListResponseTypeSessionNextCompactionStarted EventListResponseType = "session.next.compaction.started"
	EventListResponseTypeSessionNextModelSwitched     EventListResponseType = "session.next.model.switched"
	EventListResponseTypeSessionNextPrompted          EventListResponseType = "session.next.prompted"
	EventListResponseTypeSessionNextReasoningDelta    EventListResponseType = "session.next.reasoning.delta"
	EventListResponseTypeSessionNextReasoningEnded    EventListResponseType = "session.next.reasoning.ended"
	EventListResponseTypeSessionNextReasoningStarted  EventListResponseType = "session.next.reasoning.started"
	EventListResponseTypeSessionNextRetried           EventListResponseType = "session.next.retried"
	EventListResponseTypeSessionNextShellEnded        EventListResponseType = "session.next.shell.ended"
	EventListResponseTypeSessionNextShellStarted      EventListResponseType = "session.next.shell.started"
	EventListResponseTypeSessionNextStepEnded         EventListResponseType = "session.next.step.ended"
	EventListResponseTypeSessionNextStepFailed        EventListResponseType = "session.next.step.failed"
	EventListResponseTypeSessionNextStepStarted       EventListResponseType = "session.next.step.started"
	EventListResponseTypeSessionNextSynthetic         EventListResponseType = "session.next.synthetic"
	EventListResponseTypeSessionNextTextDelta         EventListResponseType = "session.next.text.delta"
	EventListResponseTypeSessionNextTextEnded         EventListResponseType = "session.next.text.ended"
	EventListResponseTypeSessionNextTextStarted       EventListResponseType = "session.next.text.started"
	EventListResponseTypeSessionNextToolCalled        EventListResponseType = "session.next.tool.called"
	EventListResponseTypeSessionNextToolFailed        EventListResponseType = "session.next.tool.failed"
	EventListResponseTypeSessionNextToolInputDelta    EventListResponseType = "session.next.tool.input.delta"
	EventListResponseTypeSessionNextToolInputEnded    EventListResponseType = "session.next.tool.input.ended"
	EventListResponseTypeSessionNextToolInputStarted  EventListResponseType = "session.next.tool.input.started"
	EventListResponseTypeSessionNextToolProgress      EventListResponseType = "session.next.tool.progress"
	EventListResponseTypeSessionNextToolSuccess       EventListResponseType = "session.next.tool.success"
	EventListResponseTypeSessionStatus                EventListResponseType = "session.status"
	EventListResponseTypeSessionUpdated               EventListResponseType = "session.updated"
	EventListResponseTypeTodoUpdated                  EventListResponseType = "todo.updated"
	EventListResponseTypeTuiCommandExecute            EventListResponseType = "tui.command.execute"
	EventListResponseTypeTuiPromptAppend              EventListResponseType = "tui.prompt.append"
	EventListResponseTypeTuiSessionSelect             EventListResponseType = "tui.session.select"
	EventListResponseTypeTuiToastShow                 EventListResponseType = "tui.toast.show"
	EventListResponseTypeVcsBranchUpdated             EventListResponseType = "vcs.branch.updated"
	EventListResponseTypeWorkspaceFailed              EventListResponseType = "workspace.failed"
	EventListResponseTypeWorkspaceReady               EventListResponseType = "workspace.ready"
	EventListResponseTypeWorkspaceStatus              EventListResponseType = "workspace.status"
	EventListResponseTypeWorktreeFailed               EventListResponseType = "worktree.failed"
	EventListResponseTypeWorktreeReady                EventListResponseType = "worktree.ready"
	EventListResponseTypeIntegrationUpdated           EventListResponseType = "integration.updated"
	EventListResponseTypeIntegrationConnectionUpdated EventListResponseType = "integration.connection.updated"
	EventListResponseTypeCatalogUpdated               EventListResponseType = "catalog.updated"
	EventListResponseTypePermissionV2Asked            EventListResponseType = "permission.v2.asked"
	EventListResponseTypePermissionV2Replied          EventListResponseType = "permission.v2.replied"
	EventListResponseTypeReferenceUpdated             EventListResponseType = "reference.updated"
	EventListResponseTypeQuestionV2Asked              EventListResponseType = "question.v2.asked"
	EventListResponseTypeQuestionV2Replied            EventListResponseType = "question.v2.replied"
	EventListResponseTypeQuestionV2Rejected           EventListResponseType = "question.v2.rejected"
	EventListResponseTypeSessionNextMoved             EventListResponseType = "session.next.moved"
	EventListResponseTypeSessionNextRevertStaged      EventListResponseType = "session.next.revert.staged"
	EventListResponseTypeSessionNextRevertCleared     EventListResponseType = "session.next.revert.cleared"
	EventListResponseTypeSessionNextRevertCommitted   EventListResponseType = "session.next.revert.committed"
	EventListResponseTypeSessionNextPromptAdmitted    EventListResponseType = "session.next.prompt.admitted"
	EventListResponseTypeSessionNextContextUpdated    EventListResponseType = "session.next.context.updated"
	EventListResponseTypeProjectDirectoriesUpdated    EventListResponseType = "project.directories.updated"
)

func (r EventListResponseType) IsKnown() bool {
	switch r {
	case EventListResponseTypeCommandExecuted,
		EventListResponseTypeFileEdited,
		EventListResponseTypeFileWatcherUpdated,
		EventListResponseTypeGlobalDisposed,
		EventListResponseTypeInstallationUpdateAvailable,
		EventListResponseTypeInstallationUpdated,
		EventListResponseTypeLspUpdated,
		EventListResponseTypeMcpBrowserOpenFailed,
		EventListResponseTypeMcpToolsChanged,
		EventListResponseTypeModelsDevRefreshed,
		EventListResponseTypeMessagePartDelta,
		EventListResponseTypeMessagePartRemoved,
		EventListResponseTypeMessagePartUpdated,
		EventListResponseTypeMessageRemoved,
		EventListResponseTypeMessageUpdated,
		EventListResponseTypePermissionAsked,
		EventListResponseTypePermissionReplied,
		EventListResponseTypePluginAdded,
		EventListResponseTypeProjectUpdated,
		EventListResponseTypePtyCreated,
		EventListResponseTypePtyDeleted,
		EventListResponseTypePtyExited,
		EventListResponseTypePtyUpdated,
		EventListResponseTypeQuestionAsked,
		EventListResponseTypeQuestionRejected,
		EventListResponseTypeQuestionReplied,
		EventListResponseTypeServerConnected,
		EventListResponseTypeServerInstanceDisposed,
		EventListResponseTypeSessionCompacted,
		EventListResponseTypeSessionCreated,
		EventListResponseTypeSessionDeleted,
		EventListResponseTypeSessionDiff,
		EventListResponseTypeSessionError,
		EventListResponseTypeSessionIdle,
		EventListResponseTypeSessionNextAgentSwitched,
		EventListResponseTypeSessionNextCompactionDelta,
		EventListResponseTypeSessionNextCompactionEnded,
		EventListResponseTypeSessionNextCompactionStarted,
		EventListResponseTypeSessionNextModelSwitched,
		EventListResponseTypeSessionNextPrompted,
		EventListResponseTypeSessionNextReasoningDelta,
		EventListResponseTypeSessionNextReasoningEnded,
		EventListResponseTypeSessionNextReasoningStarted,
		EventListResponseTypeSessionNextRetried,
		EventListResponseTypeSessionNextShellEnded,
		EventListResponseTypeSessionNextShellStarted,
		EventListResponseTypeSessionNextStepEnded,
		EventListResponseTypeSessionNextStepFailed,
		EventListResponseTypeSessionNextStepStarted,
		EventListResponseTypeSessionNextSynthetic,
		EventListResponseTypeSessionNextTextDelta,
		EventListResponseTypeSessionNextTextEnded,
		EventListResponseTypeSessionNextTextStarted,
		EventListResponseTypeSessionNextToolCalled,
		EventListResponseTypeSessionNextToolFailed,
		EventListResponseTypeSessionNextToolInputDelta,
		EventListResponseTypeSessionNextToolInputEnded,
		EventListResponseTypeSessionNextToolInputStarted,
		EventListResponseTypeSessionNextToolProgress,
		EventListResponseTypeSessionNextToolSuccess,
		EventListResponseTypeSessionStatus,
		EventListResponseTypeSessionUpdated,
		EventListResponseTypeTodoUpdated,
		EventListResponseTypeTuiCommandExecute,
		EventListResponseTypeTuiPromptAppend,
		EventListResponseTypeTuiSessionSelect,
		EventListResponseTypeTuiToastShow,
		EventListResponseTypeVcsBranchUpdated,
		EventListResponseTypeWorkspaceFailed,
		EventListResponseTypeWorkspaceReady,
		EventListResponseTypeWorkspaceStatus,
		EventListResponseTypeWorktreeFailed,
		EventListResponseTypeWorktreeReady,
		EventListResponseTypeIntegrationUpdated,
		EventListResponseTypeIntegrationConnectionUpdated,
		EventListResponseTypeCatalogUpdated,
		EventListResponseTypePermissionV2Asked,
		EventListResponseTypePermissionV2Replied,
		EventListResponseTypeReferenceUpdated,
		EventListResponseTypeQuestionV2Asked,
		EventListResponseTypeQuestionV2Replied,
		EventListResponseTypeQuestionV2Rejected,
		EventListResponseTypeSessionNextPromptAdmitted,
		EventListResponseTypeSessionNextContextUpdated,
		EventListResponseTypeProjectDirectoriesUpdated,
		EventListResponseTypeSessionNextMoved,
		EventListResponseTypeSessionNextRevertStaged,
		EventListResponseTypeSessionNextRevertCleared,
		EventListResponseTypeSessionNextRevertCommitted:
		return true
	}
	return false
}

type EventListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
