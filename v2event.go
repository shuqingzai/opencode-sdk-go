// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
	"github.com/tidwall/gjson"
)

// V2EventService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2EventService] method instead.
type V2EventService struct {
	Options []option.RequestOption
}

// NewV2EventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2EventService(opts ...option.RequestOption) (r *V2EventService) {
	r = &V2EventService{}
	r.Options = opts
	return
}

// Subscribe to native event payloads for the server.
func (r *V2EventService) Subscribe(ctx context.Context, opts ...option.RequestOption) (stream *ssestream.Stream[V2Event]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "api/event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[V2Event](ssestream.NewDecoder(raw), err)
}

// V2Event represents a native event payload from the V2 /api/event endpoint.
// The V2 format differs from V1: it uses "data" instead of "properties", and
// includes optional "durable", "location", and "metadata" fields.
type V2Event struct {
	ID   string      `json:"id,required"`
	Type V2EventType `json:"type,required"`
	// This field can have the runtime type of [V2EventDurable].
	Durable any `json:"durable"`
	// This field can have the runtime type of [LocationRef].
	Location any `json:"location"`
	// This field can have the runtime type of [map[string]any].
	Metadata any `json:"metadata"`
	// This field can have the runtime type of one of the [V2EventPayloadUnion] variants.
	Data  any         `json:"data,required"`
	JSON  v2EventJSON `json:"-"`
	union V2EventPayloadUnion
}

// V2EventDurable contains the durable sequencing information for a V2Event.
type V2EventDurable struct {
	AggregateID string             `json:"aggregateID,required"`
	Seq         int64              `json:"seq,required"`
	Version     int64              `json:"version,required"`
	JSON        v2EventDurableJSON `json:"-"`
}

type v2EventDurableJSON struct {
	AggregateID apijson.Field
	Seq         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventDurable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventDurableJSON) RawJSON() string {
	return r.raw
}

type v2EventJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2EventJSON) RawJSON() string {
	return r.raw
}

func (r *V2Event) UnmarshalJSON(data []byte) (err error) {
	*r = V2Event{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2EventPayloadUnion] interface which you can cast to the
// specific types for more type safety.
func (r V2Event) AsUnion() V2EventPayloadUnion {
	return r.union
}

// V2EventType enumerates all possible event types for V2Event.
type V2EventType string

const (
	V2EventTypeCatalogUpdated               V2EventType = "catalog.updated"
	V2EventTypeCommandExecuted              V2EventType = "command.executed"
	V2EventTypeFileEdited                   V2EventType = "file.edited"
	V2EventTypeFileWatcherUpdated           V2EventType = "file.watcher.updated"
	V2EventTypeGlobalDisposed               V2EventType = "global.disposed"
	V2EventTypeInstallationUpdateAvailable  V2EventType = "installation.update-available"
	V2EventTypeInstallationUpdated          V2EventType = "installation.updated"
	V2EventTypeIntegrationConnectionUpdated V2EventType = "integration.connection.updated"
	V2EventTypeIntegrationUpdated           V2EventType = "integration.updated"
	V2EventTypeLspUpdated                   V2EventType = "lsp.updated"
	V2EventTypeMcpBrowserOpenFailed         V2EventType = "mcp.browser.open.failed"
	V2EventTypeMcpToolsChanged              V2EventType = "mcp.tools.changed"
	V2EventTypeMessagePartDelta             V2EventType = "message.part.delta"
	V2EventTypeMessagePartRemoved           V2EventType = "message.part.removed"
	V2EventTypeMessagePartUpdated           V2EventType = "message.part.updated"
	V2EventTypeMessageRemoved               V2EventType = "message.removed"
	V2EventTypeMessageUpdated               V2EventType = "message.updated"
	V2EventTypeModelsDevRefreshed           V2EventType = "models-dev.refreshed"
	V2EventTypePermissionAsked              V2EventType = "permission.asked"
	V2EventTypePermissionReplied            V2EventType = "permission.replied"
	V2EventTypePermissionV2Asked            V2EventType = "permission.v2.asked"
	V2EventTypePermissionV2Replied          V2EventType = "permission.v2.replied"
	V2EventTypePluginAdded                  V2EventType = "plugin.added"
	V2EventTypeProjectDirectoriesUpdated    V2EventType = "project.directories.updated"
	V2EventTypeProjectUpdated               V2EventType = "project.updated"
	V2EventTypePtyCreated                   V2EventType = "pty.created"
	V2EventTypePtyDeleted                   V2EventType = "pty.deleted"
	V2EventTypePtyExited                    V2EventType = "pty.exited"
	V2EventTypePtyUpdated                   V2EventType = "pty.updated"
	V2EventTypeQuestionAsked                V2EventType = "question.asked"
	V2EventTypeQuestionRejected             V2EventType = "question.rejected"
	V2EventTypeQuestionReplied              V2EventType = "question.replied"
	V2EventTypeQuestionV2Asked              V2EventType = "question.v2.asked"
	V2EventTypeQuestionV2Rejected           V2EventType = "question.v2.rejected"
	V2EventTypeQuestionV2Replied            V2EventType = "question.v2.replied"
	V2EventTypeReferenceUpdated             V2EventType = "reference.updated"
	V2EventTypeServerConnected              V2EventType = "server.connected"
	V2EventTypeSessionCompacted             V2EventType = "session.compacted"
	V2EventTypeSessionCreated               V2EventType = "session.created"
	V2EventTypeSessionDeleted               V2EventType = "session.deleted"
	V2EventTypeSessionDiff                  V2EventType = "session.diff"
	V2EventTypeSessionError                 V2EventType = "session.error"
	V2EventTypeSessionIdle                  V2EventType = "session.idle"
	V2EventTypeSessionNextAgentSwitched     V2EventType = "session.next.agent.switched"
	V2EventTypeSessionNextCompactionDelta   V2EventType = "session.next.compaction.delta"
	V2EventTypeSessionNextCompactionEnded   V2EventType = "session.next.compaction.ended"
	V2EventTypeSessionNextCompactionStarted V2EventType = "session.next.compaction.started"
	V2EventTypeSessionNextContextUpdated    V2EventType = "session.next.context.updated"
	V2EventTypeSessionNextModelSwitched     V2EventType = "session.next.model.switched"
	V2EventTypeSessionNextMoved             V2EventType = "session.next.moved"
	V2EventTypeSessionNextPromptAdmitted    V2EventType = "session.next.prompt.admitted"
	V2EventTypeSessionNextPrompted          V2EventType = "session.next.prompted"
	V2EventTypeSessionNextReasoningDelta    V2EventType = "session.next.reasoning.delta"
	V2EventTypeSessionNextReasoningEnded    V2EventType = "session.next.reasoning.ended"
	V2EventTypeSessionNextReasoningStarted  V2EventType = "session.next.reasoning.started"
	V2EventTypeSessionNextRetried           V2EventType = "session.next.retried"
	V2EventTypeSessionNextRevertCleared     V2EventType = "session.next.revert.cleared"
	V2EventTypeSessionNextRevertCommitted   V2EventType = "session.next.revert.committed"
	V2EventTypeSessionNextRevertStaged      V2EventType = "session.next.revert.staged"
	V2EventTypeSessionNextShellEnded        V2EventType = "session.next.shell.ended"
	V2EventTypeSessionNextShellStarted      V2EventType = "session.next.shell.started"
	V2EventTypeSessionNextStepEnded         V2EventType = "session.next.step.ended"
	V2EventTypeSessionNextStepFailed        V2EventType = "session.next.step.failed"
	V2EventTypeSessionNextStepStarted       V2EventType = "session.next.step.started"
	V2EventTypeSessionNextSynthetic         V2EventType = "session.next.synthetic"
	V2EventTypeSessionNextTextDelta         V2EventType = "session.next.text.delta"
	V2EventTypeSessionNextTextEnded         V2EventType = "session.next.text.ended"
	V2EventTypeSessionNextTextStarted       V2EventType = "session.next.text.started"
	V2EventTypeSessionNextToolCalled        V2EventType = "session.next.tool.called"
	V2EventTypeSessionNextToolFailed        V2EventType = "session.next.tool.failed"
	V2EventTypeSessionNextToolInputDelta    V2EventType = "session.next.tool.input.delta"
	V2EventTypeSessionNextToolInputEnded    V2EventType = "session.next.tool.input.ended"
	V2EventTypeSessionNextToolInputStarted  V2EventType = "session.next.tool.input.started"
	V2EventTypeSessionNextToolProgress      V2EventType = "session.next.tool.progress"
	V2EventTypeSessionNextToolSuccess       V2EventType = "session.next.tool.success"
	V2EventTypeSessionStatus                V2EventType = "session.status"
	V2EventTypeSessionUpdated               V2EventType = "session.updated"
	V2EventTypeTodoUpdated                  V2EventType = "todo.updated"
	V2EventTypeTuiCommandExecute            V2EventType = "tui.command.execute"
	V2EventTypeTuiPromptAppend              V2EventType = "tui.prompt.append"
	V2EventTypeTuiSessionSelect             V2EventType = "tui.session.select"
	V2EventTypeTuiToastShow                 V2EventType = "tui.toast.show"
	V2EventTypeVcsBranchUpdated             V2EventType = "vcs.branch.updated"
	V2EventTypeWorkspaceFailed              V2EventType = "workspace.failed"
	V2EventTypeWorkspaceReady               V2EventType = "workspace.ready"
	V2EventTypeWorkspaceStatus              V2EventType = "workspace.status"
	V2EventTypeWorktreeFailed               V2EventType = "worktree.failed"
	V2EventTypeWorktreeReady                V2EventType = "worktree.ready"
)

func (r V2EventType) IsKnown() bool {
	switch r {
	case V2EventTypeCatalogUpdated,
		V2EventTypeCommandExecuted,
		V2EventTypeFileEdited,
		V2EventTypeFileWatcherUpdated,
		V2EventTypeGlobalDisposed,
		V2EventTypeInstallationUpdateAvailable,
		V2EventTypeInstallationUpdated,
		V2EventTypeIntegrationConnectionUpdated,
		V2EventTypeIntegrationUpdated,
		V2EventTypeLspUpdated,
		V2EventTypeMcpBrowserOpenFailed,
		V2EventTypeMcpToolsChanged,
		V2EventTypeMessagePartDelta,
		V2EventTypeMessagePartRemoved,
		V2EventTypeMessagePartUpdated,
		V2EventTypeMessageRemoved,
		V2EventTypeMessageUpdated,
		V2EventTypeModelsDevRefreshed,
		V2EventTypePermissionAsked,
		V2EventTypePermissionReplied,
		V2EventTypePermissionV2Asked,
		V2EventTypePermissionV2Replied,
		V2EventTypePluginAdded,
		V2EventTypeProjectDirectoriesUpdated,
		V2EventTypeProjectUpdated,
		V2EventTypePtyCreated,
		V2EventTypePtyDeleted,
		V2EventTypePtyExited,
		V2EventTypePtyUpdated,
		V2EventTypeQuestionAsked,
		V2EventTypeQuestionRejected,
		V2EventTypeQuestionReplied,
		V2EventTypeQuestionV2Asked,
		V2EventTypeQuestionV2Rejected,
		V2EventTypeQuestionV2Replied,
		V2EventTypeReferenceUpdated,
		V2EventTypeServerConnected,
		V2EventTypeSessionCompacted,
		V2EventTypeSessionCreated,
		V2EventTypeSessionDeleted,
		V2EventTypeSessionDiff,
		V2EventTypeSessionError,
		V2EventTypeSessionIdle,
		V2EventTypeSessionNextAgentSwitched,
		V2EventTypeSessionNextCompactionDelta,
		V2EventTypeSessionNextCompactionEnded,
		V2EventTypeSessionNextCompactionStarted,
		V2EventTypeSessionNextContextUpdated,
		V2EventTypeSessionNextModelSwitched,
		V2EventTypeSessionNextMoved,
		V2EventTypeSessionNextPromptAdmitted,
		V2EventTypeSessionNextPrompted,
		V2EventTypeSessionNextReasoningDelta,
		V2EventTypeSessionNextReasoningEnded,
		V2EventTypeSessionNextReasoningStarted,
		V2EventTypeSessionNextRetried,
		V2EventTypeSessionNextRevertCleared,
		V2EventTypeSessionNextRevertCommitted,
		V2EventTypeSessionNextRevertStaged,
		V2EventTypeSessionNextShellEnded,
		V2EventTypeSessionNextShellStarted,
		V2EventTypeSessionNextStepEnded,
		V2EventTypeSessionNextStepFailed,
		V2EventTypeSessionNextStepStarted,
		V2EventTypeSessionNextSynthetic,
		V2EventTypeSessionNextTextDelta,
		V2EventTypeSessionNextTextEnded,
		V2EventTypeSessionNextTextStarted,
		V2EventTypeSessionNextToolCalled,
		V2EventTypeSessionNextToolFailed,
		V2EventTypeSessionNextToolInputDelta,
		V2EventTypeSessionNextToolInputEnded,
		V2EventTypeSessionNextToolInputStarted,
		V2EventTypeSessionNextToolProgress,
		V2EventTypeSessionNextToolSuccess,
		V2EventTypeSessionStatus,
		V2EventTypeSessionUpdated,
		V2EventTypeTodoUpdated,
		V2EventTypeTuiCommandExecute,
		V2EventTypeTuiPromptAppend,
		V2EventTypeTuiSessionSelect,
		V2EventTypeTuiToastShow,
		V2EventTypeVcsBranchUpdated,
		V2EventTypeWorkspaceFailed,
		V2EventTypeWorkspaceReady,
		V2EventTypeWorkspaceStatus,
		V2EventTypeWorktreeFailed,
		V2EventTypeWorktreeReady:
		return true
	}
	return false
}

// V2EventPayloadUnion is satisfied by all V2Event variant types.
type V2EventPayloadUnion interface {
	implementsV2EventPayload()
}

type V2EventCatalogUpdated struct {
	ID       string                    `json:"id,required"`
	Data     V2EventCatalogUpdatedData `json:"data,required"`
	Type     V2EventCatalogUpdatedType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventCatalogUpdatedJSON `json:"-"`
}

type v2EventCatalogUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventCatalogUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventCatalogUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventCatalogUpdated) implementsV2EventPayload() {}

type V2EventCatalogUpdatedType string

const (
	V2EventCatalogUpdatedTypeCatalogUpdated V2EventCatalogUpdatedType = "catalog.updated"
)

func (r V2EventCatalogUpdatedType) IsKnown() bool {
	switch r {
	case V2EventCatalogUpdatedTypeCatalogUpdated:
		return true
	}
	return false
}

type V2EventCatalogUpdatedData struct {
	JSON v2EventCatalogUpdatedDataJSON `json:"-"`
}

type v2EventCatalogUpdatedDataJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventCatalogUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventCatalogUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventCommandExecuted struct {
	ID       string                     `json:"id,required"`
	Data     V2EventCommandExecutedData `json:"data,required"`
	Type     V2EventCommandExecutedType `json:"type,required"`
	Durable  V2EventDurable             `json:"durable"`
	Location LocationRef                `json:"location"`
	Metadata map[string]any             `json:"metadata"`
	JSON     v2EventCommandExecutedJSON `json:"-"`
}

type v2EventCommandExecutedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventCommandExecuted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventCommandExecutedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventCommandExecuted) implementsV2EventPayload() {}

type V2EventCommandExecutedType string

const (
	V2EventCommandExecutedTypeCommandExecuted V2EventCommandExecutedType = "command.executed"
)

func (r V2EventCommandExecutedType) IsKnown() bool {
	switch r {
	case V2EventCommandExecutedTypeCommandExecuted:
		return true
	}
	return false
}

type V2EventCommandExecutedData struct {
	Arguments string                         `json:"arguments,required"`
	MessageID string                         `json:"messageID,required"`
	Name      string                         `json:"name,required"`
	SessionID string                         `json:"sessionID,required"`
	JSON      v2EventCommandExecutedDataJSON `json:"-"`
}

type v2EventCommandExecutedDataJSON struct {
	Arguments   apijson.Field
	MessageID   apijson.Field
	Name        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventCommandExecutedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventCommandExecutedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventFileEdited struct {
	ID       string                `json:"id,required"`
	Data     V2EventFileEditedData `json:"data,required"`
	Type     V2EventFileEditedType `json:"type,required"`
	Durable  V2EventDurable        `json:"durable"`
	Location LocationRef           `json:"location"`
	Metadata map[string]any        `json:"metadata"`
	JSON     v2EventFileEditedJSON `json:"-"`
}

type v2EventFileEditedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventFileEdited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventFileEditedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventFileEdited) implementsV2EventPayload() {}

type V2EventFileEditedType string

const (
	V2EventFileEditedTypeFileEdited V2EventFileEditedType = "file.edited"
)

func (r V2EventFileEditedType) IsKnown() bool {
	switch r {
	case V2EventFileEditedTypeFileEdited:
		return true
	}
	return false
}

type V2EventFileEditedData struct {
	File string                    `json:"file,required"`
	JSON v2EventFileEditedDataJSON `json:"-"`
}

type v2EventFileEditedDataJSON struct {
	File        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventFileEditedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventFileEditedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventFileWatcherUpdatedEvent string

const (
	V2EventFileWatcherUpdatedEventAdd    V2EventFileWatcherUpdatedEvent = "add"
	V2EventFileWatcherUpdatedEventChange V2EventFileWatcherUpdatedEvent = "change"
	V2EventFileWatcherUpdatedEventUnlink V2EventFileWatcherUpdatedEvent = "unlink"
)

func (r V2EventFileWatcherUpdatedEvent) IsKnown() bool {
	switch r {
	case V2EventFileWatcherUpdatedEventAdd,
		V2EventFileWatcherUpdatedEventChange,
		V2EventFileWatcherUpdatedEventUnlink:
		return true
	}
	return false
}

type V2EventFileWatcherUpdated struct {
	ID       string                        `json:"id,required"`
	Data     V2EventFileWatcherUpdatedData `json:"data,required"`
	Type     V2EventFileWatcherUpdatedType `json:"type,required"`
	Durable  V2EventDurable                `json:"durable"`
	Location LocationRef                   `json:"location"`
	Metadata map[string]any                `json:"metadata"`
	JSON     v2EventFileWatcherUpdatedJSON `json:"-"`
}

type v2EventFileWatcherUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventFileWatcherUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventFileWatcherUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventFileWatcherUpdated) implementsV2EventPayload() {}

type V2EventFileWatcherUpdatedType string

const (
	V2EventFileWatcherUpdatedTypeFileWatcherUpdated V2EventFileWatcherUpdatedType = "file.watcher.updated"
)

func (r V2EventFileWatcherUpdatedType) IsKnown() bool {
	switch r {
	case V2EventFileWatcherUpdatedTypeFileWatcherUpdated:
		return true
	}
	return false
}

type V2EventFileWatcherUpdatedData struct {
	Event V2EventFileWatcherUpdatedEvent    `json:"event,required"`
	File  string                            `json:"file,required"`
	JSON  v2EventFileWatcherUpdatedDataJSON `json:"-"`
}

type v2EventFileWatcherUpdatedDataJSON struct {
	Event       apijson.Field
	File        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventFileWatcherUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventFileWatcherUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventGlobalDisposed struct {
	ID       string                    `json:"id,required"`
	Data     V2EventGlobalDisposedData `json:"data,required"`
	Type     V2EventGlobalDisposedType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventGlobalDisposedJSON `json:"-"`
}

type v2EventGlobalDisposedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventGlobalDisposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventGlobalDisposedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventGlobalDisposed) implementsV2EventPayload() {}

type V2EventGlobalDisposedType string

const (
	V2EventGlobalDisposedTypeGlobalDisposed V2EventGlobalDisposedType = "global.disposed"
)

func (r V2EventGlobalDisposedType) IsKnown() bool {
	switch r {
	case V2EventGlobalDisposedTypeGlobalDisposed:
		return true
	}
	return false
}

type V2EventGlobalDisposedData struct {
	JSON v2EventGlobalDisposedDataJSON `json:"-"`
}

type v2EventGlobalDisposedDataJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventGlobalDisposedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventGlobalDisposedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventInstallationUpdateAvailable struct {
	ID       string                                 `json:"id,required"`
	Data     V2EventInstallationUpdateAvailableData `json:"data,required"`
	Type     V2EventInstallationUpdateAvailableType `json:"type,required"`
	Durable  V2EventDurable                         `json:"durable"`
	Location LocationRef                            `json:"location"`
	Metadata map[string]any                         `json:"metadata"`
	JSON     v2EventInstallationUpdateAvailableJSON `json:"-"`
}

type v2EventInstallationUpdateAvailableJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventInstallationUpdateAvailable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventInstallationUpdateAvailableJSON) RawJSON() string {
	return r.raw
}

func (r V2EventInstallationUpdateAvailable) implementsV2EventPayload() {}

type V2EventInstallationUpdateAvailableType string

const (
	V2EventInstallationUpdateAvailableTypeInstallationUpdateAvailable V2EventInstallationUpdateAvailableType = "installation.update-available"
)

func (r V2EventInstallationUpdateAvailableType) IsKnown() bool {
	switch r {
	case V2EventInstallationUpdateAvailableTypeInstallationUpdateAvailable:
		return true
	}
	return false
}

type V2EventInstallationUpdateAvailableData struct {
	Version string                                     `json:"version,required"`
	JSON    v2EventInstallationUpdateAvailableDataJSON `json:"-"`
}

type v2EventInstallationUpdateAvailableDataJSON struct {
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventInstallationUpdateAvailableData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventInstallationUpdateAvailableDataJSON) RawJSON() string {
	return r.raw
}

type V2EventInstallationUpdated struct {
	ID       string                         `json:"id,required"`
	Data     V2EventInstallationUpdatedData `json:"data,required"`
	Type     V2EventInstallationUpdatedType `json:"type,required"`
	Durable  V2EventDurable                 `json:"durable"`
	Location LocationRef                    `json:"location"`
	Metadata map[string]any                 `json:"metadata"`
	JSON     v2EventInstallationUpdatedJSON `json:"-"`
}

type v2EventInstallationUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventInstallationUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventInstallationUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventInstallationUpdated) implementsV2EventPayload() {}

type V2EventInstallationUpdatedType string

const (
	V2EventInstallationUpdatedTypeInstallationUpdated V2EventInstallationUpdatedType = "installation.updated"
)

func (r V2EventInstallationUpdatedType) IsKnown() bool {
	switch r {
	case V2EventInstallationUpdatedTypeInstallationUpdated:
		return true
	}
	return false
}

type V2EventInstallationUpdatedData struct {
	Version string                             `json:"version,required"`
	JSON    v2EventInstallationUpdatedDataJSON `json:"-"`
}

type v2EventInstallationUpdatedDataJSON struct {
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventInstallationUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventInstallationUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventIntegrationConnectionUpdated struct {
	ID       string                                  `json:"id,required"`
	Data     V2EventIntegrationConnectionUpdatedData `json:"data,required"`
	Type     V2EventIntegrationConnectionUpdatedType `json:"type,required"`
	Durable  V2EventDurable                          `json:"durable"`
	Location LocationRef                             `json:"location"`
	Metadata map[string]any                          `json:"metadata"`
	JSON     v2EventIntegrationConnectionUpdatedJSON `json:"-"`
}

type v2EventIntegrationConnectionUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventIntegrationConnectionUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventIntegrationConnectionUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventIntegrationConnectionUpdated) implementsV2EventPayload() {}

type V2EventIntegrationConnectionUpdatedType string

const (
	V2EventIntegrationConnectionUpdatedTypeIntegrationConnectionUpdated V2EventIntegrationConnectionUpdatedType = "integration.connection.updated"
)

func (r V2EventIntegrationConnectionUpdatedType) IsKnown() bool {
	switch r {
	case V2EventIntegrationConnectionUpdatedTypeIntegrationConnectionUpdated:
		return true
	}
	return false
}

type V2EventIntegrationConnectionUpdatedData struct {
	IntegrationID string                                      `json:"integrationID,required"`
	JSON          v2EventIntegrationConnectionUpdatedDataJSON `json:"-"`
}

type v2EventIntegrationConnectionUpdatedDataJSON struct {
	IntegrationID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2EventIntegrationConnectionUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventIntegrationConnectionUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventIntegrationUpdated struct {
	ID       string                        `json:"id,required"`
	Data     V2EventIntegrationUpdatedData `json:"data,required"`
	Type     V2EventIntegrationUpdatedType `json:"type,required"`
	Durable  V2EventDurable                `json:"durable"`
	Location LocationRef                   `json:"location"`
	Metadata map[string]any                `json:"metadata"`
	JSON     v2EventIntegrationUpdatedJSON `json:"-"`
}

type v2EventIntegrationUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventIntegrationUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventIntegrationUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventIntegrationUpdated) implementsV2EventPayload() {}

type V2EventIntegrationUpdatedType string

const (
	V2EventIntegrationUpdatedTypeIntegrationUpdated V2EventIntegrationUpdatedType = "integration.updated"
)

func (r V2EventIntegrationUpdatedType) IsKnown() bool {
	switch r {
	case V2EventIntegrationUpdatedTypeIntegrationUpdated:
		return true
	}
	return false
}

type V2EventIntegrationUpdatedData struct {
	JSON v2EventIntegrationUpdatedDataJSON `json:"-"`
}

type v2EventIntegrationUpdatedDataJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventIntegrationUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventIntegrationUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventLspUpdated struct {
	ID       string                `json:"id,required"`
	Data     V2EventLspUpdatedData `json:"data,required"`
	Type     V2EventLspUpdatedType `json:"type,required"`
	Durable  V2EventDurable        `json:"durable"`
	Location LocationRef           `json:"location"`
	Metadata map[string]any        `json:"metadata"`
	JSON     v2EventLspUpdatedJSON `json:"-"`
}

type v2EventLspUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventLspUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventLspUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventLspUpdated) implementsV2EventPayload() {}

type V2EventLspUpdatedType string

const (
	V2EventLspUpdatedTypeLspUpdated V2EventLspUpdatedType = "lsp.updated"
)

func (r V2EventLspUpdatedType) IsKnown() bool {
	switch r {
	case V2EventLspUpdatedTypeLspUpdated:
		return true
	}
	return false
}

type V2EventLspUpdatedData struct {
	JSON v2EventLspUpdatedDataJSON `json:"-"`
}

type v2EventLspUpdatedDataJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventLspUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventLspUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventMcpBrowserOpenFailed struct {
	ID       string                          `json:"id,required"`
	Data     V2EventMcpBrowserOpenFailedData `json:"data,required"`
	Type     V2EventMcpBrowserOpenFailedType `json:"type,required"`
	Durable  V2EventDurable                  `json:"durable"`
	Location LocationRef                     `json:"location"`
	Metadata map[string]any                  `json:"metadata"`
	JSON     v2EventMcpBrowserOpenFailedJSON `json:"-"`
}

type v2EventMcpBrowserOpenFailedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMcpBrowserOpenFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMcpBrowserOpenFailedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventMcpBrowserOpenFailed) implementsV2EventPayload() {}

type V2EventMcpBrowserOpenFailedType string

const (
	V2EventMcpBrowserOpenFailedTypeMcpBrowserOpenFailed V2EventMcpBrowserOpenFailedType = "mcp.browser.open.failed"
)

func (r V2EventMcpBrowserOpenFailedType) IsKnown() bool {
	switch r {
	case V2EventMcpBrowserOpenFailedTypeMcpBrowserOpenFailed:
		return true
	}
	return false
}

type V2EventMcpBrowserOpenFailedData struct {
	McpName string                              `json:"mcpName,required"`
	URL     string                              `json:"url,required"`
	JSON    v2EventMcpBrowserOpenFailedDataJSON `json:"-"`
}

type v2EventMcpBrowserOpenFailedDataJSON struct {
	McpName     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMcpBrowserOpenFailedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMcpBrowserOpenFailedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventMcpToolsChanged struct {
	ID       string                     `json:"id,required"`
	Data     V2EventMcpToolsChangedData `json:"data,required"`
	Type     V2EventMcpToolsChangedType `json:"type,required"`
	Durable  V2EventDurable             `json:"durable"`
	Location LocationRef                `json:"location"`
	Metadata map[string]any             `json:"metadata"`
	JSON     v2EventMcpToolsChangedJSON `json:"-"`
}

type v2EventMcpToolsChangedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMcpToolsChanged) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMcpToolsChangedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventMcpToolsChanged) implementsV2EventPayload() {}

type V2EventMcpToolsChangedType string

const (
	V2EventMcpToolsChangedTypeMcpToolsChanged V2EventMcpToolsChangedType = "mcp.tools.changed"
)

func (r V2EventMcpToolsChangedType) IsKnown() bool {
	switch r {
	case V2EventMcpToolsChangedTypeMcpToolsChanged:
		return true
	}
	return false
}

type V2EventMcpToolsChangedData struct {
	Server string                         `json:"server,required"`
	JSON   v2EventMcpToolsChangedDataJSON `json:"-"`
}

type v2EventMcpToolsChangedDataJSON struct {
	Server      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMcpToolsChangedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMcpToolsChangedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventMessagePartDelta struct {
	ID       string                      `json:"id,required"`
	Data     V2EventMessagePartDeltaData `json:"data,required"`
	Type     V2EventMessagePartDeltaType `json:"type,required"`
	Durable  V2EventDurable              `json:"durable"`
	Location LocationRef                 `json:"location"`
	Metadata map[string]any              `json:"metadata"`
	JSON     v2EventMessagePartDeltaJSON `json:"-"`
}

type v2EventMessagePartDeltaJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessagePartDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessagePartDeltaJSON) RawJSON() string {
	return r.raw
}

func (r V2EventMessagePartDelta) implementsV2EventPayload() {}

type V2EventMessagePartDeltaType string

const (
	V2EventMessagePartDeltaTypeMessagePartDelta V2EventMessagePartDeltaType = "message.part.delta"
)

func (r V2EventMessagePartDeltaType) IsKnown() bool {
	switch r {
	case V2EventMessagePartDeltaTypeMessagePartDelta:
		return true
	}
	return false
}

type V2EventMessagePartDeltaData struct {
	Delta     string                          `json:"delta,required"`
	Field     string                          `json:"field,required"`
	MessageID string                          `json:"messageID,required"`
	PartID    string                          `json:"partID,required"`
	SessionID string                          `json:"sessionID,required"`
	JSON      v2EventMessagePartDeltaDataJSON `json:"-"`
}

type v2EventMessagePartDeltaDataJSON struct {
	Delta       apijson.Field
	Field       apijson.Field
	MessageID   apijson.Field
	PartID      apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessagePartDeltaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessagePartDeltaDataJSON) RawJSON() string {
	return r.raw
}

type V2EventMessagePartRemoved struct {
	ID       string                        `json:"id,required"`
	Data     V2EventMessagePartRemovedData `json:"data,required"`
	Type     V2EventMessagePartRemovedType `json:"type,required"`
	Durable  V2EventDurable                `json:"durable"`
	Location LocationRef                   `json:"location"`
	Metadata map[string]any                `json:"metadata"`
	JSON     v2EventMessagePartRemovedJSON `json:"-"`
}

type v2EventMessagePartRemovedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessagePartRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessagePartRemovedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventMessagePartRemoved) implementsV2EventPayload() {}

type V2EventMessagePartRemovedType string

const (
	V2EventMessagePartRemovedTypeMessagePartRemoved V2EventMessagePartRemovedType = "message.part.removed"
)

func (r V2EventMessagePartRemovedType) IsKnown() bool {
	switch r {
	case V2EventMessagePartRemovedTypeMessagePartRemoved:
		return true
	}
	return false
}

type V2EventMessagePartRemovedData struct {
	MessageID string                            `json:"messageID,required"`
	PartID    string                            `json:"partID,required"`
	SessionID string                            `json:"sessionID,required"`
	JSON      v2EventMessagePartRemovedDataJSON `json:"-"`
}

type v2EventMessagePartRemovedDataJSON struct {
	MessageID   apijson.Field
	PartID      apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessagePartRemovedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessagePartRemovedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventMessagePartUpdated struct {
	ID       string                        `json:"id,required"`
	Data     V2EventMessagePartUpdatedData `json:"data,required"`
	Type     V2EventMessagePartUpdatedType `json:"type,required"`
	Durable  V2EventDurable                `json:"durable"`
	Location LocationRef                   `json:"location"`
	Metadata map[string]any                `json:"metadata"`
	JSON     v2EventMessagePartUpdatedJSON `json:"-"`
}

type v2EventMessagePartUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessagePartUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessagePartUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventMessagePartUpdated) implementsV2EventPayload() {}

type V2EventMessagePartUpdatedType string

const (
	V2EventMessagePartUpdatedTypeMessagePartUpdated V2EventMessagePartUpdatedType = "message.part.updated"
)

func (r V2EventMessagePartUpdatedType) IsKnown() bool {
	switch r {
	case V2EventMessagePartUpdatedTypeMessagePartUpdated:
		return true
	}
	return false
}

type V2EventMessagePartUpdatedData struct {
	Part      Part                              `json:"part,required"`
	SessionID string                            `json:"sessionID,required"`
	Time      int64                             `json:"time,required"`
	JSON      v2EventMessagePartUpdatedDataJSON `json:"-"`
}

type v2EventMessagePartUpdatedDataJSON struct {
	Part        apijson.Field
	SessionID   apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessagePartUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessagePartUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventMessageRemoved struct {
	ID       string                    `json:"id,required"`
	Data     V2EventMessageRemovedData `json:"data,required"`
	Type     V2EventMessageRemovedType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventMessageRemovedJSON `json:"-"`
}

type v2EventMessageRemovedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessageRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessageRemovedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventMessageRemoved) implementsV2EventPayload() {}

type V2EventMessageRemovedType string

const (
	V2EventMessageRemovedTypeMessageRemoved V2EventMessageRemovedType = "message.removed"
)

func (r V2EventMessageRemovedType) IsKnown() bool {
	switch r {
	case V2EventMessageRemovedTypeMessageRemoved:
		return true
	}
	return false
}

type V2EventMessageRemovedData struct {
	MessageID string                        `json:"messageID,required"`
	SessionID string                        `json:"sessionID,required"`
	JSON      v2EventMessageRemovedDataJSON `json:"-"`
}

type v2EventMessageRemovedDataJSON struct {
	MessageID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessageRemovedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessageRemovedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventMessageUpdated struct {
	ID       string                    `json:"id,required"`
	Data     V2EventMessageUpdatedData `json:"data,required"`
	Type     V2EventMessageUpdatedType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventMessageUpdatedJSON `json:"-"`
}

type v2EventMessageUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessageUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessageUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventMessageUpdated) implementsV2EventPayload() {}

type V2EventMessageUpdatedType string

const (
	V2EventMessageUpdatedTypeMessageUpdated V2EventMessageUpdatedType = "message.updated"
)

func (r V2EventMessageUpdatedType) IsKnown() bool {
	switch r {
	case V2EventMessageUpdatedTypeMessageUpdated:
		return true
	}
	return false
}

type V2EventMessageUpdatedData struct {
	Info      Message                       `json:"info,required"`
	SessionID string                        `json:"sessionID,required"`
	JSON      v2EventMessageUpdatedDataJSON `json:"-"`
}

type v2EventMessageUpdatedDataJSON struct {
	Info        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventMessageUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventMessageUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventModelsDevRefreshed struct {
	ID       string                        `json:"id,required"`
	Data     V2EventModelsDevRefreshedData `json:"data,required"`
	Type     V2EventModelsDevRefreshedType `json:"type,required"`
	Durable  V2EventDurable                `json:"durable"`
	Location LocationRef                   `json:"location"`
	Metadata map[string]any                `json:"metadata"`
	JSON     v2EventModelsDevRefreshedJSON `json:"-"`
}

type v2EventModelsDevRefreshedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventModelsDevRefreshed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventModelsDevRefreshedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventModelsDevRefreshed) implementsV2EventPayload() {}

type V2EventModelsDevRefreshedType string

const (
	V2EventModelsDevRefreshedTypeModelsDevRefreshed V2EventModelsDevRefreshedType = "models-dev.refreshed"
)

func (r V2EventModelsDevRefreshedType) IsKnown() bool {
	switch r {
	case V2EventModelsDevRefreshedTypeModelsDevRefreshed:
		return true
	}
	return false
}

type V2EventModelsDevRefreshedData struct {
	JSON v2EventModelsDevRefreshedDataJSON `json:"-"`
}

type v2EventModelsDevRefreshedDataJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventModelsDevRefreshedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventModelsDevRefreshedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventPermissionAsked struct {
	ID       string                     `json:"id,required"`
	Data     V2EventPermissionAskedData `json:"data,required"`
	Type     V2EventPermissionAskedType `json:"type,required"`
	Durable  V2EventDurable             `json:"durable"`
	Location LocationRef                `json:"location"`
	Metadata map[string]any             `json:"metadata"`
	JSON     v2EventPermissionAskedJSON `json:"-"`
}

type v2EventPermissionAskedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPermissionAsked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPermissionAskedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventPermissionAsked) implementsV2EventPayload() {}

type V2EventPermissionAskedType string

const (
	V2EventPermissionAskedTypePermissionAsked V2EventPermissionAskedType = "permission.asked"
)

func (r V2EventPermissionAskedType) IsKnown() bool {
	switch r {
	case V2EventPermissionAskedTypePermissionAsked:
		return true
	}
	return false
}

type V2EventPermissionAskedData struct {
	Always []string `json:"always,required"`
	ID     string   `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata   any                            `json:"metadata,required"`
	Patterns   []string                       `json:"patterns,required"`
	Permission string                         `json:"permission,required"`
	SessionID  string                         `json:"sessionID,required"`
	Tool       V2EventPermissionAskedDataTool `json:"tool"`
	JSON       v2EventPermissionAskedDataJSON `json:"-"`
}

type v2EventPermissionAskedDataJSON struct {
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

func (r *V2EventPermissionAskedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPermissionAskedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventPermissionAskedDataTool struct {
	CallID    string                             `json:"callID,required"`
	MessageID string                             `json:"messageID,required"`
	JSON      V2EventPermissionAskedDataToolJSON `json:"-"`
}

type V2EventPermissionAskedDataToolJSON struct {
	CallID      apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPermissionAskedDataTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r V2EventPermissionAskedDataToolJSON) RawJSON() string {
	return r.raw
}

type V2EventPermissionReplied struct {
	ID       string                       `json:"id,required"`
	Data     V2EventPermissionRepliedData `json:"data,required"`
	Type     V2EventPermissionRepliedType `json:"type,required"`
	Durable  V2EventDurable               `json:"durable"`
	Location LocationRef                  `json:"location"`
	Metadata map[string]any               `json:"metadata"`
	JSON     v2EventPermissionRepliedJSON `json:"-"`
}

type v2EventPermissionRepliedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPermissionReplied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPermissionRepliedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventPermissionReplied) implementsV2EventPayload() {}

type V2EventPermissionRepliedType string

const (
	V2EventPermissionRepliedTypePermissionReplied V2EventPermissionRepliedType = "permission.replied"
)

func (r V2EventPermissionRepliedType) IsKnown() bool {
	switch r {
	case V2EventPermissionRepliedTypePermissionReplied:
		return true
	}
	return false
}

type V2EventPermissionRepliedData struct {
	Reply     PermissionV2Reply                `json:"reply,required"`
	RequestID string                           `json:"requestID,required"`
	SessionID string                           `json:"sessionID,required"`
	JSON      v2EventPermissionRepliedDataJSON `json:"-"`
}

type v2EventPermissionRepliedDataJSON struct {
	Reply       apijson.Field
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPermissionRepliedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPermissionRepliedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventPermissionV2Asked struct {
	ID       string                       `json:"id,required"`
	Data     V2EventPermissionV2AskedData `json:"data,required"`
	Type     V2EventPermissionV2AskedType `json:"type,required"`
	Durable  V2EventDurable               `json:"durable"`
	Location LocationRef                  `json:"location"`
	Metadata map[string]any               `json:"metadata"`
	JSON     v2EventPermissionV2AskedJSON `json:"-"`
}

type v2EventPermissionV2AskedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPermissionV2Asked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPermissionV2AskedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventPermissionV2Asked) implementsV2EventPayload() {}

type V2EventPermissionV2AskedType string

const (
	V2EventPermissionV2AskedTypePermissionV2Asked V2EventPermissionV2AskedType = "permission.v2.asked"
)

func (r V2EventPermissionV2AskedType) IsKnown() bool {
	switch r {
	case V2EventPermissionV2AskedTypePermissionV2Asked:
		return true
	}
	return false
}

type V2EventPermissionV2AskedData struct {
	Action string `json:"action,required"`
	ID     string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata  any                              `json:"metadata"`
	Resources []string                         `json:"resources,required"`
	Save      []string                         `json:"save"`
	SessionID string                           `json:"sessionID,required"`
	Source    PermissionV2Source               `json:"source"`
	JSON      v2EventPermissionV2AskedDataJSON `json:"-"`
}

type v2EventPermissionV2AskedDataJSON struct {
	Action      apijson.Field
	ID          apijson.Field
	Metadata    apijson.Field
	Resources   apijson.Field
	Save        apijson.Field
	SessionID   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPermissionV2AskedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPermissionV2AskedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventPermissionV2Replied struct {
	ID       string                         `json:"id,required"`
	Data     V2EventPermissionV2RepliedData `json:"data,required"`
	Type     V2EventPermissionV2RepliedType `json:"type,required"`
	Durable  V2EventDurable                 `json:"durable"`
	Location LocationRef                    `json:"location"`
	Metadata map[string]any                 `json:"metadata"`
	JSON     v2EventPermissionV2RepliedJSON `json:"-"`
}

type v2EventPermissionV2RepliedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPermissionV2Replied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPermissionV2RepliedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventPermissionV2Replied) implementsV2EventPayload() {}

type V2EventPermissionV2RepliedType string

const (
	V2EventPermissionV2RepliedTypePermissionV2Replied V2EventPermissionV2RepliedType = "permission.v2.replied"
)

func (r V2EventPermissionV2RepliedType) IsKnown() bool {
	switch r {
	case V2EventPermissionV2RepliedTypePermissionV2Replied:
		return true
	}
	return false
}

type V2EventPermissionV2RepliedData struct {
	Reply     PermissionV2Reply                  `json:"reply,required"`
	RequestID string                             `json:"requestID,required"`
	SessionID string                             `json:"sessionID,required"`
	JSON      v2EventPermissionV2RepliedDataJSON `json:"-"`
}

type v2EventPermissionV2RepliedDataJSON struct {
	Reply       apijson.Field
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPermissionV2RepliedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPermissionV2RepliedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventPluginAdded struct {
	ID       string                 `json:"id,required"`
	Data     V2EventPluginAddedData `json:"data,required"`
	Type     V2EventPluginAddedType `json:"type,required"`
	Durable  V2EventDurable         `json:"durable"`
	Location LocationRef            `json:"location"`
	Metadata map[string]any         `json:"metadata"`
	JSON     v2EventPluginAddedJSON `json:"-"`
}

type v2EventPluginAddedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPluginAdded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPluginAddedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventPluginAdded) implementsV2EventPayload() {}

type V2EventPluginAddedType string

const (
	V2EventPluginAddedTypePluginAdded V2EventPluginAddedType = "plugin.added"
)

func (r V2EventPluginAddedType) IsKnown() bool {
	switch r {
	case V2EventPluginAddedTypePluginAdded:
		return true
	}
	return false
}

type V2EventPluginAddedData struct {
	ID   string                     `json:"id,required"`
	JSON v2EventPluginAddedDataJSON `json:"-"`
}

type v2EventPluginAddedDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPluginAddedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPluginAddedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventProjectDirectoriesUpdated struct {
	ID       string                               `json:"id,required"`
	Data     V2EventProjectDirectoriesUpdatedData `json:"data,required"`
	Type     V2EventProjectDirectoriesUpdatedType `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Metadata map[string]any                       `json:"metadata"`
	JSON     v2EventProjectDirectoriesUpdatedJSON `json:"-"`
}

type v2EventProjectDirectoriesUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventProjectDirectoriesUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventProjectDirectoriesUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventProjectDirectoriesUpdated) implementsV2EventPayload() {}

type V2EventProjectDirectoriesUpdatedType string

const (
	V2EventProjectDirectoriesUpdatedTypeProjectDirectoriesUpdated V2EventProjectDirectoriesUpdatedType = "project.directories.updated"
)

func (r V2EventProjectDirectoriesUpdatedType) IsKnown() bool {
	switch r {
	case V2EventProjectDirectoriesUpdatedTypeProjectDirectoriesUpdated:
		return true
	}
	return false
}

type V2EventProjectDirectoriesUpdatedData struct {
	ProjectID string                                   `json:"projectID,required"`
	JSON      v2EventProjectDirectoriesUpdatedDataJSON `json:"-"`
}

type v2EventProjectDirectoriesUpdatedDataJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventProjectDirectoriesUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventProjectDirectoriesUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventProjectUpdated struct {
	ID       string                    `json:"id,required"`
	Data     V2EventProjectUpdatedData `json:"data,required"`
	Type     V2EventProjectUpdatedType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventProjectUpdatedJSON `json:"-"`
}

type v2EventProjectUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventProjectUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventProjectUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventProjectUpdated) implementsV2EventPayload() {}

type V2EventProjectUpdatedType string

const (
	V2EventProjectUpdatedTypeProjectUpdated V2EventProjectUpdatedType = "project.updated"
)

func (r V2EventProjectUpdatedType) IsKnown() bool {
	switch r {
	case V2EventProjectUpdatedTypeProjectUpdated:
		return true
	}
	return false
}

type V2EventProjectUpdatedData struct {
	Commands  ProjectCommands               `json:"commands"`
	Icon      ProjectIcon                   `json:"icon"`
	ID        string                        `json:"id,required"`
	Name      string                        `json:"name"`
	Sandboxes []string                      `json:"sandboxes,required"`
	Time      ProjectTime                   `json:"time,required"`
	Vcs       ProjectVcs                    `json:"vcs"`
	Worktree  string                        `json:"worktree,required"`
	JSON      v2EventProjectUpdatedDataJSON `json:"-"`
}

type v2EventProjectUpdatedDataJSON struct {
	Commands    apijson.Field
	Icon        apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Sandboxes   apijson.Field
	Time        apijson.Field
	Vcs         apijson.Field
	Worktree    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventProjectUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventProjectUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventPtyCreated struct {
	ID       string                `json:"id,required"`
	Data     V2EventPtyCreatedData `json:"data,required"`
	Type     V2EventPtyCreatedType `json:"type,required"`
	Durable  V2EventDurable        `json:"durable"`
	Location LocationRef           `json:"location"`
	Metadata map[string]any        `json:"metadata"`
	JSON     v2EventPtyCreatedJSON `json:"-"`
}

type v2EventPtyCreatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPtyCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPtyCreatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventPtyCreated) implementsV2EventPayload() {}

type V2EventPtyCreatedType string

const (
	V2EventPtyCreatedTypePtyCreated V2EventPtyCreatedType = "pty.created"
)

func (r V2EventPtyCreatedType) IsKnown() bool {
	switch r {
	case V2EventPtyCreatedTypePtyCreated:
		return true
	}
	return false
}

type V2EventPtyCreatedData struct {
	Info Pty                       `json:"info,required"`
	JSON v2EventPtyCreatedDataJSON `json:"-"`
}

type v2EventPtyCreatedDataJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPtyCreatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPtyCreatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventPtyDeleted struct {
	ID       string                `json:"id,required"`
	Data     V2EventPtyDeletedData `json:"data,required"`
	Type     V2EventPtyDeletedType `json:"type,required"`
	Durable  V2EventDurable        `json:"durable"`
	Location LocationRef           `json:"location"`
	Metadata map[string]any        `json:"metadata"`
	JSON     v2EventPtyDeletedJSON `json:"-"`
}

type v2EventPtyDeletedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPtyDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPtyDeletedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventPtyDeleted) implementsV2EventPayload() {}

type V2EventPtyDeletedType string

const (
	V2EventPtyDeletedTypePtyDeleted V2EventPtyDeletedType = "pty.deleted"
)

func (r V2EventPtyDeletedType) IsKnown() bool {
	switch r {
	case V2EventPtyDeletedTypePtyDeleted:
		return true
	}
	return false
}

type V2EventPtyDeletedData struct {
	ID   string                    `json:"id,required"`
	JSON v2EventPtyDeletedDataJSON `json:"-"`
}

type v2EventPtyDeletedDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPtyDeletedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPtyDeletedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventPtyExited struct {
	ID       string               `json:"id,required"`
	Data     V2EventPtyExitedData `json:"data,required"`
	Type     V2EventPtyExitedType `json:"type,required"`
	Durable  V2EventDurable       `json:"durable"`
	Location LocationRef          `json:"location"`
	Metadata map[string]any       `json:"metadata"`
	JSON     v2EventPtyExitedJSON `json:"-"`
}

type v2EventPtyExitedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPtyExited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPtyExitedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventPtyExited) implementsV2EventPayload() {}

type V2EventPtyExitedType string

const (
	V2EventPtyExitedTypePtyExited V2EventPtyExitedType = "pty.exited"
)

func (r V2EventPtyExitedType) IsKnown() bool {
	switch r {
	case V2EventPtyExitedTypePtyExited:
		return true
	}
	return false
}

type V2EventPtyExitedData struct {
	ExitCode int64                    `json:"exitCode,required"`
	ID       string                   `json:"id,required"`
	JSON     v2EventPtyExitedDataJSON `json:"-"`
}

type v2EventPtyExitedDataJSON struct {
	ExitCode    apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPtyExitedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPtyExitedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventPtyUpdated struct {
	ID       string                `json:"id,required"`
	Data     V2EventPtyUpdatedData `json:"data,required"`
	Type     V2EventPtyUpdatedType `json:"type,required"`
	Durable  V2EventDurable        `json:"durable"`
	Location LocationRef           `json:"location"`
	Metadata map[string]any        `json:"metadata"`
	JSON     v2EventPtyUpdatedJSON `json:"-"`
}

type v2EventPtyUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPtyUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPtyUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventPtyUpdated) implementsV2EventPayload() {}

type V2EventPtyUpdatedType string

const (
	V2EventPtyUpdatedTypePtyUpdated V2EventPtyUpdatedType = "pty.updated"
)

func (r V2EventPtyUpdatedType) IsKnown() bool {
	switch r {
	case V2EventPtyUpdatedTypePtyUpdated:
		return true
	}
	return false
}

type V2EventPtyUpdatedData struct {
	Info Pty                       `json:"info,required"`
	JSON v2EventPtyUpdatedDataJSON `json:"-"`
}

type v2EventPtyUpdatedDataJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventPtyUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventPtyUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventQuestionAsked struct {
	ID       string                   `json:"id,required"`
	Data     V2EventQuestionAskedData `json:"data,required"`
	Type     V2EventQuestionAskedType `json:"type,required"`
	Durable  V2EventDurable           `json:"durable"`
	Location LocationRef              `json:"location"`
	Metadata map[string]any           `json:"metadata"`
	JSON     v2EventQuestionAskedJSON `json:"-"`
}

type v2EventQuestionAskedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionAsked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionAskedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventQuestionAsked) implementsV2EventPayload() {}

type V2EventQuestionAskedType string

const (
	V2EventQuestionAskedTypeQuestionAsked V2EventQuestionAskedType = "question.asked"
)

func (r V2EventQuestionAskedType) IsKnown() bool {
	switch r {
	case V2EventQuestionAskedTypeQuestionAsked:
		return true
	}
	return false
}

type V2EventQuestionAskedData struct {
	ID        string                       `json:"id,required"`
	Questions []QuestionInfo               `json:"questions,required"`
	SessionID string                       `json:"sessionID,required"`
	Tool      QuestionTool                 `json:"tool"`
	JSON      v2EventQuestionAskedDataJSON `json:"-"`
}

type v2EventQuestionAskedDataJSON struct {
	ID          apijson.Field
	Questions   apijson.Field
	SessionID   apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionAskedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionAskedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventQuestionRejected struct {
	ID       string                      `json:"id,required"`
	Data     V2EventQuestionRejectedData `json:"data,required"`
	Type     V2EventQuestionRejectedType `json:"type,required"`
	Durable  V2EventDurable              `json:"durable"`
	Location LocationRef                 `json:"location"`
	Metadata map[string]any              `json:"metadata"`
	JSON     v2EventQuestionRejectedJSON `json:"-"`
}

type v2EventQuestionRejectedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionRejected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionRejectedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventQuestionRejected) implementsV2EventPayload() {}

type V2EventQuestionRejectedType string

const (
	V2EventQuestionRejectedTypeQuestionRejected V2EventQuestionRejectedType = "question.rejected"
)

func (r V2EventQuestionRejectedType) IsKnown() bool {
	switch r {
	case V2EventQuestionRejectedTypeQuestionRejected:
		return true
	}
	return false
}

type V2EventQuestionRejectedData struct {
	RequestID string                          `json:"requestID,required"`
	SessionID string                          `json:"sessionID,required"`
	JSON      v2EventQuestionRejectedDataJSON `json:"-"`
}

type v2EventQuestionRejectedDataJSON struct {
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionRejectedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionRejectedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventQuestionReplied struct {
	ID       string                     `json:"id,required"`
	Data     V2EventQuestionRepliedData `json:"data,required"`
	Type     V2EventQuestionRepliedType `json:"type,required"`
	Durable  V2EventDurable             `json:"durable"`
	Location LocationRef                `json:"location"`
	Metadata map[string]any             `json:"metadata"`
	JSON     v2EventQuestionRepliedJSON `json:"-"`
}

type v2EventQuestionRepliedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionReplied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionRepliedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventQuestionReplied) implementsV2EventPayload() {}

type V2EventQuestionRepliedType string

const (
	V2EventQuestionRepliedTypeQuestionReplied V2EventQuestionRepliedType = "question.replied"
)

func (r V2EventQuestionRepliedType) IsKnown() bool {
	switch r {
	case V2EventQuestionRepliedTypeQuestionReplied:
		return true
	}
	return false
}

type V2EventQuestionRepliedData struct {
	Answers   []QuestionAnswer               `json:"answers,required"`
	RequestID string                         `json:"requestID,required"`
	SessionID string                         `json:"sessionID,required"`
	JSON      v2EventQuestionRepliedDataJSON `json:"-"`
}

type v2EventQuestionRepliedDataJSON struct {
	Answers     apijson.Field
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionRepliedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionRepliedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventQuestionV2Asked struct {
	ID       string                     `json:"id,required"`
	Data     V2EventQuestionV2AskedData `json:"data,required"`
	Type     V2EventQuestionV2AskedType `json:"type,required"`
	Durable  V2EventDurable             `json:"durable"`
	Location LocationRef                `json:"location"`
	Metadata map[string]any             `json:"metadata"`
	JSON     v2EventQuestionV2AskedJSON `json:"-"`
}

type v2EventQuestionV2AskedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionV2Asked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionV2AskedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventQuestionV2Asked) implementsV2EventPayload() {}

type V2EventQuestionV2AskedType string

const (
	V2EventQuestionV2AskedTypeQuestionV2Asked V2EventQuestionV2AskedType = "question.v2.asked"
)

func (r V2EventQuestionV2AskedType) IsKnown() bool {
	switch r {
	case V2EventQuestionV2AskedTypeQuestionV2Asked:
		return true
	}
	return false
}

type V2EventQuestionV2AskedData struct {
	ID        string                         `json:"id,required"`
	Questions []QuestionV2Info               `json:"questions,required"`
	SessionID string                         `json:"sessionID,required"`
	Tool      QuestionV2Tool                 `json:"tool"`
	JSON      v2EventQuestionV2AskedDataJSON `json:"-"`
}

type v2EventQuestionV2AskedDataJSON struct {
	ID          apijson.Field
	Questions   apijson.Field
	SessionID   apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionV2AskedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionV2AskedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventQuestionV2Rejected struct {
	ID       string                        `json:"id,required"`
	Data     V2EventQuestionV2RejectedData `json:"data,required"`
	Type     V2EventQuestionV2RejectedType `json:"type,required"`
	Durable  V2EventDurable                `json:"durable"`
	Location LocationRef                   `json:"location"`
	Metadata map[string]any                `json:"metadata"`
	JSON     v2EventQuestionV2RejectedJSON `json:"-"`
}

type v2EventQuestionV2RejectedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionV2Rejected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionV2RejectedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventQuestionV2Rejected) implementsV2EventPayload() {}

type V2EventQuestionV2RejectedType string

const (
	V2EventQuestionV2RejectedTypeQuestionV2Rejected V2EventQuestionV2RejectedType = "question.v2.rejected"
)

func (r V2EventQuestionV2RejectedType) IsKnown() bool {
	switch r {
	case V2EventQuestionV2RejectedTypeQuestionV2Rejected:
		return true
	}
	return false
}

type V2EventQuestionV2RejectedData struct {
	RequestID string                            `json:"requestID,required"`
	SessionID string                            `json:"sessionID,required"`
	JSON      v2EventQuestionV2RejectedDataJSON `json:"-"`
}

type v2EventQuestionV2RejectedDataJSON struct {
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionV2RejectedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionV2RejectedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventQuestionV2Replied struct {
	ID       string                       `json:"id,required"`
	Data     V2EventQuestionV2RepliedData `json:"data,required"`
	Type     V2EventQuestionV2RepliedType `json:"type,required"`
	Durable  V2EventDurable               `json:"durable"`
	Location LocationRef                  `json:"location"`
	Metadata map[string]any               `json:"metadata"`
	JSON     v2EventQuestionV2RepliedJSON `json:"-"`
}

type v2EventQuestionV2RepliedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionV2Replied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionV2RepliedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventQuestionV2Replied) implementsV2EventPayload() {}

type V2EventQuestionV2RepliedType string

const (
	V2EventQuestionV2RepliedTypeQuestionV2Replied V2EventQuestionV2RepliedType = "question.v2.replied"
)

func (r V2EventQuestionV2RepliedType) IsKnown() bool {
	switch r {
	case V2EventQuestionV2RepliedTypeQuestionV2Replied:
		return true
	}
	return false
}

type V2EventQuestionV2RepliedData struct {
	Answers   []QuestionV2Answer               `json:"answers,required"`
	RequestID string                           `json:"requestID,required"`
	SessionID string                           `json:"sessionID,required"`
	JSON      v2EventQuestionV2RepliedDataJSON `json:"-"`
}

type v2EventQuestionV2RepliedDataJSON struct {
	Answers     apijson.Field
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventQuestionV2RepliedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventQuestionV2RepliedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventReferenceUpdated struct {
	ID       string                      `json:"id,required"`
	Data     V2EventReferenceUpdatedData `json:"data,required"`
	Type     V2EventReferenceUpdatedType `json:"type,required"`
	Durable  V2EventDurable              `json:"durable"`
	Location LocationRef                 `json:"location"`
	Metadata map[string]any              `json:"metadata"`
	JSON     v2EventReferenceUpdatedJSON `json:"-"`
}

type v2EventReferenceUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventReferenceUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventReferenceUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventReferenceUpdated) implementsV2EventPayload() {}

type V2EventReferenceUpdatedType string

const (
	V2EventReferenceUpdatedTypeReferenceUpdated V2EventReferenceUpdatedType = "reference.updated"
)

func (r V2EventReferenceUpdatedType) IsKnown() bool {
	switch r {
	case V2EventReferenceUpdatedTypeReferenceUpdated:
		return true
	}
	return false
}

type V2EventReferenceUpdatedData struct {
	JSON v2EventReferenceUpdatedDataJSON `json:"-"`
}

type v2EventReferenceUpdatedDataJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventReferenceUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventReferenceUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventServerConnected struct {
	ID       string                     `json:"id,required"`
	Data     V2EventServerConnectedData `json:"data,required"`
	Type     V2EventServerConnectedType `json:"type,required"`
	Durable  V2EventDurable             `json:"durable"`
	Location LocationRef                `json:"location"`
	Metadata map[string]any             `json:"metadata"`
	JSON     v2EventServerConnectedJSON `json:"-"`
}

type v2EventServerConnectedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventServerConnected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventServerConnectedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventServerConnected) implementsV2EventPayload() {}

type V2EventServerConnectedType string

const (
	V2EventServerConnectedTypeServerConnected V2EventServerConnectedType = "server.connected"
)

func (r V2EventServerConnectedType) IsKnown() bool {
	switch r {
	case V2EventServerConnectedTypeServerConnected:
		return true
	}
	return false
}

type V2EventServerConnectedData struct {
	JSON v2EventServerConnectedDataJSON `json:"-"`
}

type v2EventServerConnectedDataJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventServerConnectedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventServerConnectedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionCompacted struct {
	ID       string                      `json:"id,required"`
	Data     V2EventSessionCompactedData `json:"data,required"`
	Type     V2EventSessionCompactedType `json:"type,required"`
	Durable  V2EventDurable              `json:"durable"`
	Location LocationRef                 `json:"location"`
	Metadata map[string]any              `json:"metadata"`
	JSON     v2EventSessionCompactedJSON `json:"-"`
}

type v2EventSessionCompactedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionCompacted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionCompactedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionCompacted) implementsV2EventPayload() {}

type V2EventSessionCompactedType string

const (
	V2EventSessionCompactedTypeSessionCompacted V2EventSessionCompactedType = "session.compacted"
)

func (r V2EventSessionCompactedType) IsKnown() bool {
	switch r {
	case V2EventSessionCompactedTypeSessionCompacted:
		return true
	}
	return false
}

type V2EventSessionCompactedData struct {
	SessionID string                          `json:"sessionID,required"`
	JSON      v2EventSessionCompactedDataJSON `json:"-"`
}

type v2EventSessionCompactedDataJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionCompactedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionCompactedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionCreated struct {
	ID       string                    `json:"id,required"`
	Data     V2EventSessionCreatedData `json:"data,required"`
	Type     V2EventSessionCreatedType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventSessionCreatedJSON `json:"-"`
}

type v2EventSessionCreatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionCreatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionCreated) implementsV2EventPayload() {}

type V2EventSessionCreatedType string

const (
	V2EventSessionCreatedTypeSessionCreated V2EventSessionCreatedType = "session.created"
)

func (r V2EventSessionCreatedType) IsKnown() bool {
	switch r {
	case V2EventSessionCreatedTypeSessionCreated:
		return true
	}
	return false
}

type V2EventSessionCreatedData struct {
	Info      Session                       `json:"info,required"`
	SessionID string                        `json:"sessionID,required"`
	JSON      v2EventSessionCreatedDataJSON `json:"-"`
}

type v2EventSessionCreatedDataJSON struct {
	Info        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionCreatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionCreatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionDeleted struct {
	ID       string                    `json:"id,required"`
	Data     V2EventSessionDeletedData `json:"data,required"`
	Type     V2EventSessionDeletedType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventSessionDeletedJSON `json:"-"`
}

type v2EventSessionDeletedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionDeletedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionDeleted) implementsV2EventPayload() {}

type V2EventSessionDeletedType string

const (
	V2EventSessionDeletedTypeSessionDeleted V2EventSessionDeletedType = "session.deleted"
)

func (r V2EventSessionDeletedType) IsKnown() bool {
	switch r {
	case V2EventSessionDeletedTypeSessionDeleted:
		return true
	}
	return false
}

type V2EventSessionDeletedData struct {
	Info      Session                       `json:"info,required"`
	SessionID string                        `json:"sessionID,required"`
	JSON      v2EventSessionDeletedDataJSON `json:"-"`
}

type v2EventSessionDeletedDataJSON struct {
	Info        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionDeletedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionDeletedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionDiff struct {
	ID       string                 `json:"id,required"`
	Data     V2EventSessionDiffData `json:"data,required"`
	Type     V2EventSessionDiffType `json:"type,required"`
	Durable  V2EventDurable         `json:"durable"`
	Location LocationRef            `json:"location"`
	Metadata map[string]any         `json:"metadata"`
	JSON     v2EventSessionDiffJSON `json:"-"`
}

type v2EventSessionDiffJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionDiffJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionDiff) implementsV2EventPayload() {}

type V2EventSessionDiffType string

const (
	V2EventSessionDiffTypeSessionDiff V2EventSessionDiffType = "session.diff"
)

func (r V2EventSessionDiffType) IsKnown() bool {
	switch r {
	case V2EventSessionDiffTypeSessionDiff:
		return true
	}
	return false
}

type V2EventSessionDiffData struct {
	Diff      []SnapshotFileDiff         `json:"diff,required"`
	SessionID string                     `json:"sessionID,required"`
	JSON      v2EventSessionDiffDataJSON `json:"-"`
}

type v2EventSessionDiffDataJSON struct {
	Diff        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionDiffData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionDiffDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionError struct {
	ID       string                  `json:"id,required"`
	Data     V2EventSessionErrorData `json:"data,required"`
	Type     V2EventSessionErrorType `json:"type,required"`
	Durable  V2EventDurable          `json:"durable"`
	Location LocationRef             `json:"location"`
	Metadata map[string]any          `json:"metadata"`
	JSON     v2EventSessionErrorJSON `json:"-"`
}

type v2EventSessionErrorJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionErrorJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionError) implementsV2EventPayload() {}

type V2EventSessionErrorType string

const (
	V2EventSessionErrorTypeSessionError V2EventSessionErrorType = "session.error"
)

func (r V2EventSessionErrorType) IsKnown() bool {
	switch r {
	case V2EventSessionErrorTypeSessionError:
		return true
	}
	return false
}

type V2EventSessionErrorData struct {
	// This field can have the runtime type of [ProviderAuthError], [UnknownError],
	// [MessageOutputLengthError], [MessageAbortedError], [StructuredOutputError],
	// [ContextOverflowError], [ContentFilterError], [APIError].
	Error     any                         `json:"error"`
	SessionID string                      `json:"sessionID"`
	JSON      v2EventSessionErrorDataJSON `json:"-"`
	// errorUnion holds the typed payload after [UnmarshalJSON] routes the raw
	// data through the registered [V2EventSessionErrorDataErrorUnion] variants.
	errorUnion V2EventSessionErrorDataErrorUnion
}

type v2EventSessionErrorDataJSON struct {
	Error       apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionErrorData) UnmarshalJSON(data []byte) (err error) {
	*r = V2EventSessionErrorData{}
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	errorData := gjson.GetBytes(data, "error").Raw
	if errorData != "" && errorData != "null" {
		if err = apijson.UnmarshalRoot([]byte(errorData), &r.errorUnion); err != nil {
			return err
		}
		r.Error = r.errorUnion
	}
	return nil
}

func (r v2EventSessionErrorDataJSON) RawJSON() string {
	return r.raw
}

// AsError returns the error field as a typed union.
//
// Possible runtime types of the union are [ProviderAuthError], [UnknownError],
// [MessageOutputLengthError], [MessageAbortedError], [StructuredOutputError],
// [ContextOverflowError], [ContentFilterError], [APIError].
func (r *V2EventSessionErrorData) AsError() V2EventSessionErrorDataErrorUnion {
	return r.errorUnion
}

// Union satisfied by [ProviderAuthError], [UnknownError], [MessageOutputLengthError],
// [MessageAbortedError], [StructuredOutputError], [ContextOverflowError],
// [ContentFilterError] or [APIError].
type V2EventSessionErrorDataErrorUnion interface {
	ImplementsEventListResponseEventSessionErrorPropertiesError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[V2EventSessionErrorDataErrorUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ProviderAuthError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[UnknownError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[MessageOutputLengthError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[MessageAbortedError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[StructuredOutputError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ContextOverflowError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ContentFilterError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[APIError](),
		},
	)
}

type V2EventSessionIdle struct {
	ID       string                 `json:"id,required"`
	Data     V2EventSessionIdleData `json:"data,required"`
	Type     V2EventSessionIdleType `json:"type,required"`
	Durable  V2EventDurable         `json:"durable"`
	Location LocationRef            `json:"location"`
	Metadata map[string]any         `json:"metadata"`
	JSON     v2EventSessionIdleJSON `json:"-"`
}

type v2EventSessionIdleJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionIdle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionIdleJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionIdle) implementsV2EventPayload() {}

type V2EventSessionIdleType string

const (
	V2EventSessionIdleTypeSessionIdle V2EventSessionIdleType = "session.idle"
)

func (r V2EventSessionIdleType) IsKnown() bool {
	switch r {
	case V2EventSessionIdleTypeSessionIdle:
		return true
	}
	return false
}

type V2EventSessionIdleData struct {
	SessionID string                     `json:"sessionID,required"`
	JSON      v2EventSessionIdleDataJSON `json:"-"`
}

type v2EventSessionIdleDataJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionIdleData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionIdleDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextAgentSwitched struct {
	ID       string                              `json:"id,required"`
	Data     V2EventSessionNextAgentSwitchedData `json:"data,required"`
	Type     V2EventSessionNextAgentSwitchedType `json:"type,required"`
	Durable  V2EventDurable                      `json:"durable"`
	Location LocationRef                         `json:"location"`
	Metadata map[string]any                      `json:"metadata"`
	JSON     v2EventSessionNextAgentSwitchedJSON `json:"-"`
}

type v2EventSessionNextAgentSwitchedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextAgentSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextAgentSwitchedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextAgentSwitched) implementsV2EventPayload() {}

type V2EventSessionNextAgentSwitchedType string

const (
	V2EventSessionNextAgentSwitchedTypeSessionNextAgentSwitched V2EventSessionNextAgentSwitchedType = "session.next.agent.switched"
)

func (r V2EventSessionNextAgentSwitchedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextAgentSwitchedTypeSessionNextAgentSwitched:
		return true
	}
	return false
}

type V2EventSessionNextAgentSwitchedData struct {
	Agent     string                                  `json:"agent,required"`
	MessageID string                                  `json:"messageID,required"`
	SessionID string                                  `json:"sessionID,required"`
	Timestamp int64                                   `json:"timestamp,required"`
	JSON      v2EventSessionNextAgentSwitchedDataJSON `json:"-"`
}

type v2EventSessionNextAgentSwitchedDataJSON struct {
	Agent       apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextAgentSwitchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextAgentSwitchedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextCompactionDelta struct {
	ID       string                                `json:"id,required"`
	Data     V2EventSessionNextCompactionDeltaData `json:"data,required"`
	Type     V2EventSessionNextCompactionDeltaType `json:"type,required"`
	Durable  V2EventDurable                        `json:"durable"`
	Location LocationRef                           `json:"location"`
	Metadata map[string]any                        `json:"metadata"`
	JSON     v2EventSessionNextCompactionDeltaJSON `json:"-"`
}

type v2EventSessionNextCompactionDeltaJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextCompactionDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextCompactionDeltaJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextCompactionDelta) implementsV2EventPayload() {}

type V2EventSessionNextCompactionDeltaType string

const (
	V2EventSessionNextCompactionDeltaTypeSessionNextCompactionDelta V2EventSessionNextCompactionDeltaType = "session.next.compaction.delta"
)

func (r V2EventSessionNextCompactionDeltaType) IsKnown() bool {
	switch r {
	case V2EventSessionNextCompactionDeltaTypeSessionNextCompactionDelta:
		return true
	}
	return false
}

type V2EventSessionNextCompactionDeltaData struct {
	MessageID string                                    `json:"messageID,required"`
	SessionID string                                    `json:"sessionID,required"`
	Text      string                                    `json:"text,required"`
	Timestamp int64                                     `json:"timestamp,required"`
	JSON      v2EventSessionNextCompactionDeltaDataJSON `json:"-"`
}

type v2EventSessionNextCompactionDeltaDataJSON struct {
	MessageID   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextCompactionDeltaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextCompactionDeltaDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextCompactionEndedReason string

const (
	V2EventSessionNextCompactionEndedReasonAuto   V2EventSessionNextCompactionEndedReason = "auto"
	V2EventSessionNextCompactionEndedReasonManual V2EventSessionNextCompactionEndedReason = "manual"
)

func (r V2EventSessionNextCompactionEndedReason) IsKnown() bool {
	switch r {
	case V2EventSessionNextCompactionEndedReasonAuto,
		V2EventSessionNextCompactionEndedReasonManual:
		return true
	}
	return false
}

type V2EventSessionNextCompactionEnded struct {
	ID       string                                `json:"id,required"`
	Data     V2EventSessionNextCompactionEndedData `json:"data,required"`
	Type     V2EventSessionNextCompactionEndedType `json:"type,required"`
	Durable  V2EventDurable                        `json:"durable"`
	Location LocationRef                           `json:"location"`
	Metadata map[string]any                        `json:"metadata"`
	JSON     v2EventSessionNextCompactionEndedJSON `json:"-"`
}

type v2EventSessionNextCompactionEndedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextCompactionEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextCompactionEndedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextCompactionEnded) implementsV2EventPayload() {}

type V2EventSessionNextCompactionEndedType string

const (
	V2EventSessionNextCompactionEndedTypeSessionNextCompactionEnded V2EventSessionNextCompactionEndedType = "session.next.compaction.ended"
)

func (r V2EventSessionNextCompactionEndedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextCompactionEndedTypeSessionNextCompactionEnded:
		return true
	}
	return false
}

type V2EventSessionNextCompactionEndedData struct {
	MessageID string                                    `json:"messageID,required"`
	Reason    V2EventSessionNextCompactionEndedReason   `json:"reason,required"`
	Recent    string                                    `json:"recent,required"`
	SessionID string                                    `json:"sessionID,required"`
	Text      string                                    `json:"text,required"`
	Timestamp int64                                     `json:"timestamp,required"`
	JSON      v2EventSessionNextCompactionEndedDataJSON `json:"-"`
}

type v2EventSessionNextCompactionEndedDataJSON struct {
	MessageID   apijson.Field
	Reason      apijson.Field
	Recent      apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextCompactionEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextCompactionEndedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextCompactionStartedReason string

const (
	V2EventSessionNextCompactionStartedReasonAuto   V2EventSessionNextCompactionStartedReason = "auto"
	V2EventSessionNextCompactionStartedReasonManual V2EventSessionNextCompactionStartedReason = "manual"
)

func (r V2EventSessionNextCompactionStartedReason) IsKnown() bool {
	switch r {
	case V2EventSessionNextCompactionStartedReasonAuto,
		V2EventSessionNextCompactionStartedReasonManual:
		return true
	}
	return false
}

type V2EventSessionNextCompactionStarted struct {
	ID       string                                  `json:"id,required"`
	Data     V2EventSessionNextCompactionStartedData `json:"data,required"`
	Type     V2EventSessionNextCompactionStartedType `json:"type,required"`
	Durable  V2EventDurable                          `json:"durable"`
	Location LocationRef                             `json:"location"`
	Metadata map[string]any                          `json:"metadata"`
	JSON     v2EventSessionNextCompactionStartedJSON `json:"-"`
}

type v2EventSessionNextCompactionStartedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextCompactionStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextCompactionStartedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextCompactionStarted) implementsV2EventPayload() {}

type V2EventSessionNextCompactionStartedType string

const (
	V2EventSessionNextCompactionStartedTypeSessionNextCompactionStarted V2EventSessionNextCompactionStartedType = "session.next.compaction.started"
)

func (r V2EventSessionNextCompactionStartedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextCompactionStartedTypeSessionNextCompactionStarted:
		return true
	}
	return false
}

type V2EventSessionNextCompactionStartedData struct {
	MessageID string                                      `json:"messageID,required"`
	Reason    V2EventSessionNextCompactionStartedReason   `json:"reason,required"`
	SessionID string                                      `json:"sessionID,required"`
	Timestamp int64                                       `json:"timestamp,required"`
	JSON      v2EventSessionNextCompactionStartedDataJSON `json:"-"`
}

type v2EventSessionNextCompactionStartedDataJSON struct {
	MessageID   apijson.Field
	Reason      apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextCompactionStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextCompactionStartedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextContextUpdated struct {
	ID       string                               `json:"id,required"`
	Data     V2EventSessionNextContextUpdatedData `json:"data,required"`
	Type     V2EventSessionNextContextUpdatedType `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Metadata map[string]any                       `json:"metadata"`
	JSON     v2EventSessionNextContextUpdatedJSON `json:"-"`
}

type v2EventSessionNextContextUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextContextUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextContextUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextContextUpdated) implementsV2EventPayload() {}

type V2EventSessionNextContextUpdatedType string

const (
	V2EventSessionNextContextUpdatedTypeSessionNextContextUpdated V2EventSessionNextContextUpdatedType = "session.next.context.updated"
)

func (r V2EventSessionNextContextUpdatedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextContextUpdatedTypeSessionNextContextUpdated:
		return true
	}
	return false
}

type V2EventSessionNextContextUpdatedData struct {
	MessageID string                                   `json:"messageID,required"`
	SessionID string                                   `json:"sessionID,required"`
	Text      string                                   `json:"text,required"`
	Timestamp int64                                    `json:"timestamp,required"`
	JSON      v2EventSessionNextContextUpdatedDataJSON `json:"-"`
}

type v2EventSessionNextContextUpdatedDataJSON struct {
	MessageID   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextContextUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextContextUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextModelSwitched struct {
	ID       string                              `json:"id,required"`
	Data     V2EventSessionNextModelSwitchedData `json:"data,required"`
	Type     V2EventSessionNextModelSwitchedType `json:"type,required"`
	Durable  V2EventDurable                      `json:"durable"`
	Location LocationRef                         `json:"location"`
	Metadata map[string]any                      `json:"metadata"`
	JSON     v2EventSessionNextModelSwitchedJSON `json:"-"`
}

type v2EventSessionNextModelSwitchedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextModelSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextModelSwitchedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextModelSwitched) implementsV2EventPayload() {}

type V2EventSessionNextModelSwitchedType string

const (
	V2EventSessionNextModelSwitchedTypeSessionNextModelSwitched V2EventSessionNextModelSwitchedType = "session.next.model.switched"
)

func (r V2EventSessionNextModelSwitchedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextModelSwitchedTypeSessionNextModelSwitched:
		return true
	}
	return false
}

type V2EventSessionNextModelSwitchedData struct {
	MessageID string                                  `json:"messageID,required"`
	Model     ModelRef                                `json:"model,required"`
	SessionID string                                  `json:"sessionID,required"`
	Timestamp int64                                   `json:"timestamp,required"`
	JSON      v2EventSessionNextModelSwitchedDataJSON `json:"-"`
}

type v2EventSessionNextModelSwitchedDataJSON struct {
	MessageID   apijson.Field
	Model       apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextModelSwitchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextModelSwitchedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextMoved struct {
	ID       string                      `json:"id,required"`
	Data     V2EventSessionNextMovedData `json:"data,required"`
	Type     V2EventSessionNextMovedType `json:"type,required"`
	Durable  V2EventDurable              `json:"durable"`
	Location LocationRef                 `json:"location"`
	Metadata map[string]any              `json:"metadata"`
	JSON     v2EventSessionNextMovedJSON `json:"-"`
}

type v2EventSessionNextMovedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextMoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextMovedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextMoved) implementsV2EventPayload() {}

type V2EventSessionNextMovedType string

const (
	V2EventSessionNextMovedTypeSessionNextMoved V2EventSessionNextMovedType = "session.next.moved"
)

func (r V2EventSessionNextMovedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextMovedTypeSessionNextMoved:
		return true
	}
	return false
}

type V2EventSessionNextMovedData struct {
	Location     LocationRef                     `json:"location,required"`
	SessionID    string                          `json:"sessionID,required"`
	Subdirectory string                          `json:"subdirectory"`
	Timestamp    int64                           `json:"timestamp,required"`
	JSON         v2EventSessionNextMovedDataJSON `json:"-"`
}

type v2EventSessionNextMovedDataJSON struct {
	Location     apijson.Field
	SessionID    apijson.Field
	Subdirectory apijson.Field
	Timestamp    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2EventSessionNextMovedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextMovedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextPromptAdmittedDelivery string

const (
	V2EventSessionNextPromptAdmittedDeliverySteer V2EventSessionNextPromptAdmittedDelivery = "steer"
	V2EventSessionNextPromptAdmittedDeliveryQueue V2EventSessionNextPromptAdmittedDelivery = "queue"
)

func (r V2EventSessionNextPromptAdmittedDelivery) IsKnown() bool {
	switch r {
	case V2EventSessionNextPromptAdmittedDeliverySteer,
		V2EventSessionNextPromptAdmittedDeliveryQueue:
		return true
	}
	return false
}

type V2EventSessionNextPromptAdmitted struct {
	ID       string                               `json:"id,required"`
	Data     V2EventSessionNextPromptAdmittedData `json:"data,required"`
	Type     V2EventSessionNextPromptAdmittedType `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Metadata map[string]any                       `json:"metadata"`
	JSON     v2EventSessionNextPromptAdmittedJSON `json:"-"`
}

type v2EventSessionNextPromptAdmittedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextPromptAdmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextPromptAdmittedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextPromptAdmitted) implementsV2EventPayload() {}

type V2EventSessionNextPromptAdmittedType string

const (
	V2EventSessionNextPromptAdmittedTypeSessionNextPromptAdmitted V2EventSessionNextPromptAdmittedType = "session.next.prompt.admitted"
)

func (r V2EventSessionNextPromptAdmittedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextPromptAdmittedTypeSessionNextPromptAdmitted:
		return true
	}
	return false
}

type V2EventSessionNextPromptAdmittedData struct {
	Delivery  V2EventSessionNextPromptAdmittedDelivery `json:"delivery,required"`
	MessageID string                                   `json:"messageID,required"`
	Prompt    V2SessionInputPrompt                     `json:"prompt,required"`
	SessionID string                                   `json:"sessionID,required"`
	Timestamp int64                                    `json:"timestamp,required"`
	JSON      v2EventSessionNextPromptAdmittedDataJSON `json:"-"`
}

type v2EventSessionNextPromptAdmittedDataJSON struct {
	Delivery    apijson.Field
	MessageID   apijson.Field
	Prompt      apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextPromptAdmittedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextPromptAdmittedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextPromptedDelivery string

const (
	V2EventSessionNextPromptedDeliverySteer V2EventSessionNextPromptedDelivery = "steer"
	V2EventSessionNextPromptedDeliveryQueue V2EventSessionNextPromptedDelivery = "queue"
)

func (r V2EventSessionNextPromptedDelivery) IsKnown() bool {
	switch r {
	case V2EventSessionNextPromptedDeliverySteer,
		V2EventSessionNextPromptedDeliveryQueue:
		return true
	}
	return false
}

type V2EventSessionNextPrompted struct {
	ID       string                         `json:"id,required"`
	Data     V2EventSessionNextPromptedData `json:"data,required"`
	Type     V2EventSessionNextPromptedType `json:"type,required"`
	Durable  V2EventDurable                 `json:"durable"`
	Location LocationRef                    `json:"location"`
	Metadata map[string]any                 `json:"metadata"`
	JSON     v2EventSessionNextPromptedJSON `json:"-"`
}

type v2EventSessionNextPromptedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextPrompted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextPromptedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextPrompted) implementsV2EventPayload() {}

type V2EventSessionNextPromptedType string

const (
	V2EventSessionNextPromptedTypeSessionNextPrompted V2EventSessionNextPromptedType = "session.next.prompted"
)

func (r V2EventSessionNextPromptedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextPromptedTypeSessionNextPrompted:
		return true
	}
	return false
}

type V2EventSessionNextPromptedData struct {
	Delivery  V2EventSessionNextPromptedDelivery `json:"delivery,required"`
	MessageID string                             `json:"messageID,required"`
	Prompt    V2SessionInputPrompt               `json:"prompt,required"`
	SessionID string                             `json:"sessionID,required"`
	Timestamp int64                              `json:"timestamp,required"`
	JSON      v2EventSessionNextPromptedDataJSON `json:"-"`
}

type v2EventSessionNextPromptedDataJSON struct {
	Delivery    apijson.Field
	MessageID   apijson.Field
	Prompt      apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextPromptedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextPromptedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextReasoningDelta struct {
	ID       string                               `json:"id,required"`
	Data     V2EventSessionNextReasoningDeltaData `json:"data,required"`
	Type     V2EventSessionNextReasoningDeltaType `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Metadata map[string]any                       `json:"metadata"`
	JSON     v2EventSessionNextReasoningDeltaJSON `json:"-"`
}

type v2EventSessionNextReasoningDeltaJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextReasoningDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextReasoningDeltaJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextReasoningDelta) implementsV2EventPayload() {}

type V2EventSessionNextReasoningDeltaType string

const (
	V2EventSessionNextReasoningDeltaTypeSessionNextReasoningDelta V2EventSessionNextReasoningDeltaType = "session.next.reasoning.delta"
)

func (r V2EventSessionNextReasoningDeltaType) IsKnown() bool {
	switch r {
	case V2EventSessionNextReasoningDeltaTypeSessionNextReasoningDelta:
		return true
	}
	return false
}

type V2EventSessionNextReasoningDeltaData struct {
	AssistantMessageID string                                   `json:"assistantMessageID,required"`
	Delta              string                                   `json:"delta,required"`
	ReasoningID        string                                   `json:"reasoningID,required"`
	SessionID          string                                   `json:"sessionID,required"`
	Timestamp          int64                                    `json:"timestamp,required"`
	JSON               v2EventSessionNextReasoningDeltaDataJSON `json:"-"`
}

type v2EventSessionNextReasoningDeltaDataJSON struct {
	AssistantMessageID apijson.Field
	Delta              apijson.Field
	ReasoningID        apijson.Field
	SessionID          apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextReasoningDeltaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextReasoningDeltaDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextReasoningEnded struct {
	ID       string                               `json:"id,required"`
	Data     V2EventSessionNextReasoningEndedData `json:"data,required"`
	Type     V2EventSessionNextReasoningEndedType `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Metadata map[string]any                       `json:"metadata"`
	JSON     v2EventSessionNextReasoningEndedJSON `json:"-"`
}

type v2EventSessionNextReasoningEndedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextReasoningEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextReasoningEndedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextReasoningEnded) implementsV2EventPayload() {}

type V2EventSessionNextReasoningEndedType string

const (
	V2EventSessionNextReasoningEndedTypeSessionNextReasoningEnded V2EventSessionNextReasoningEndedType = "session.next.reasoning.ended"
)

func (r V2EventSessionNextReasoningEndedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextReasoningEndedTypeSessionNextReasoningEnded:
		return true
	}
	return false
}

type V2EventSessionNextReasoningEndedData struct {
	AssistantMessageID string `json:"assistantMessageID,required"`
	// This field can have the runtime type of [map[string]any].
	ProviderMetadata any                                      `json:"providerMetadata"`
	ReasoningID      string                                   `json:"reasoningID,required"`
	SessionID        string                                   `json:"sessionID,required"`
	Text             string                                   `json:"text,required"`
	Timestamp        int64                                    `json:"timestamp,required"`
	JSON             v2EventSessionNextReasoningEndedDataJSON `json:"-"`
}

type v2EventSessionNextReasoningEndedDataJSON struct {
	AssistantMessageID apijson.Field
	ProviderMetadata   apijson.Field
	ReasoningID        apijson.Field
	SessionID          apijson.Field
	Text               apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextReasoningEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextReasoningEndedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextReasoningStarted struct {
	ID       string                                 `json:"id,required"`
	Data     V2EventSessionNextReasoningStartedData `json:"data,required"`
	Type     V2EventSessionNextReasoningStartedType `json:"type,required"`
	Durable  V2EventDurable                         `json:"durable"`
	Location LocationRef                            `json:"location"`
	Metadata map[string]any                         `json:"metadata"`
	JSON     v2EventSessionNextReasoningStartedJSON `json:"-"`
}

type v2EventSessionNextReasoningStartedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextReasoningStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextReasoningStartedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextReasoningStarted) implementsV2EventPayload() {}

type V2EventSessionNextReasoningStartedType string

const (
	V2EventSessionNextReasoningStartedTypeSessionNextReasoningStarted V2EventSessionNextReasoningStartedType = "session.next.reasoning.started"
)

func (r V2EventSessionNextReasoningStartedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextReasoningStartedTypeSessionNextReasoningStarted:
		return true
	}
	return false
}

type V2EventSessionNextReasoningStartedData struct {
	AssistantMessageID string `json:"assistantMessageID,required"`
	// This field can have the runtime type of [map[string]any].
	ProviderMetadata any                                        `json:"providerMetadata"`
	ReasoningID      string                                     `json:"reasoningID,required"`
	SessionID        string                                     `json:"sessionID,required"`
	Timestamp        int64                                      `json:"timestamp,required"`
	JSON             v2EventSessionNextReasoningStartedDataJSON `json:"-"`
}

type v2EventSessionNextReasoningStartedDataJSON struct {
	AssistantMessageID apijson.Field
	ProviderMetadata   apijson.Field
	ReasoningID        apijson.Field
	SessionID          apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextReasoningStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextReasoningStartedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextRetried struct {
	ID       string                        `json:"id,required"`
	Data     V2EventSessionNextRetriedData `json:"data,required"`
	Type     V2EventSessionNextRetriedType `json:"type,required"`
	Durable  V2EventDurable                `json:"durable"`
	Location LocationRef                   `json:"location"`
	Metadata map[string]any                `json:"metadata"`
	JSON     v2EventSessionNextRetriedJSON `json:"-"`
}

type v2EventSessionNextRetriedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextRetried) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextRetriedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextRetried) implementsV2EventPayload() {}

type V2EventSessionNextRetriedType string

const (
	V2EventSessionNextRetriedTypeSessionNextRetried V2EventSessionNextRetriedType = "session.next.retried"
)

func (r V2EventSessionNextRetriedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextRetriedTypeSessionNextRetried:
		return true
	}
	return false
}

type V2EventSessionNextRetriedData struct {
	Attempt   int64                                         `json:"attempt,required"`
	Error     EventListResponseEventSessionNextRetriedError `json:"error,required"`
	SessionID string                                        `json:"sessionID,required"`
	Timestamp int64                                         `json:"timestamp,required"`
	JSON      v2EventSessionNextRetriedDataJSON             `json:"-"`
}

type v2EventSessionNextRetriedDataJSON struct {
	Attempt     apijson.Field
	Error       apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextRetriedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextRetriedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextRevertCleared struct {
	ID       string                              `json:"id,required"`
	Data     V2EventSessionNextRevertClearedData `json:"data,required"`
	Type     V2EventSessionNextRevertClearedType `json:"type,required"`
	Durable  V2EventDurable                      `json:"durable"`
	Location LocationRef                         `json:"location"`
	Metadata map[string]any                      `json:"metadata"`
	JSON     v2EventSessionNextRevertClearedJSON `json:"-"`
}

type v2EventSessionNextRevertClearedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextRevertCleared) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextRevertClearedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextRevertCleared) implementsV2EventPayload() {}

type V2EventSessionNextRevertClearedType string

const (
	V2EventSessionNextRevertClearedTypeSessionNextRevertCleared V2EventSessionNextRevertClearedType = "session.next.revert.cleared"
)

func (r V2EventSessionNextRevertClearedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextRevertClearedTypeSessionNextRevertCleared:
		return true
	}
	return false
}

type V2EventSessionNextRevertClearedData struct {
	SessionID string                                  `json:"sessionID,required"`
	Timestamp int64                                   `json:"timestamp,required"`
	JSON      v2EventSessionNextRevertClearedDataJSON `json:"-"`
}

type v2EventSessionNextRevertClearedDataJSON struct {
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextRevertClearedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextRevertClearedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextRevertCommitted struct {
	ID       string                                `json:"id,required"`
	Data     V2EventSessionNextRevertCommittedData `json:"data,required"`
	Type     V2EventSessionNextRevertCommittedType `json:"type,required"`
	Durable  V2EventDurable                        `json:"durable"`
	Location LocationRef                           `json:"location"`
	Metadata map[string]any                        `json:"metadata"`
	JSON     v2EventSessionNextRevertCommittedJSON `json:"-"`
}

type v2EventSessionNextRevertCommittedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextRevertCommitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextRevertCommittedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextRevertCommitted) implementsV2EventPayload() {}

type V2EventSessionNextRevertCommittedType string

const (
	V2EventSessionNextRevertCommittedTypeSessionNextRevertCommitted V2EventSessionNextRevertCommittedType = "session.next.revert.committed"
)

func (r V2EventSessionNextRevertCommittedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextRevertCommittedTypeSessionNextRevertCommitted:
		return true
	}
	return false
}

type V2EventSessionNextRevertCommittedData struct {
	MessageID string                                    `json:"messageID,required"`
	SessionID string                                    `json:"sessionID,required"`
	Timestamp int64                                     `json:"timestamp,required"`
	JSON      v2EventSessionNextRevertCommittedDataJSON `json:"-"`
}

type v2EventSessionNextRevertCommittedDataJSON struct {
	MessageID   apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextRevertCommittedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextRevertCommittedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextRevertStaged struct {
	ID       string                             `json:"id,required"`
	Data     V2EventSessionNextRevertStagedData `json:"data,required"`
	Type     V2EventSessionNextRevertStagedType `json:"type,required"`
	Durable  V2EventDurable                     `json:"durable"`
	Location LocationRef                        `json:"location"`
	Metadata map[string]any                     `json:"metadata"`
	JSON     v2EventSessionNextRevertStagedJSON `json:"-"`
}

type v2EventSessionNextRevertStagedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextRevertStaged) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextRevertStagedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextRevertStaged) implementsV2EventPayload() {}

type V2EventSessionNextRevertStagedType string

const (
	V2EventSessionNextRevertStagedTypeSessionNextRevertStaged V2EventSessionNextRevertStagedType = "session.next.revert.staged"
)

func (r V2EventSessionNextRevertStagedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextRevertStagedTypeSessionNextRevertStaged:
		return true
	}
	return false
}

type V2EventSessionNextRevertStagedData struct {
	Revert    RevertState                            `json:"revert,required"`
	SessionID string                                 `json:"sessionID,required"`
	Timestamp int64                                  `json:"timestamp,required"`
	JSON      v2EventSessionNextRevertStagedDataJSON `json:"-"`
}

type v2EventSessionNextRevertStagedDataJSON struct {
	Revert      apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextRevertStagedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextRevertStagedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextShellEnded struct {
	ID       string                           `json:"id,required"`
	Data     V2EventSessionNextShellEndedData `json:"data,required"`
	Type     V2EventSessionNextShellEndedType `json:"type,required"`
	Durable  V2EventDurable                   `json:"durable"`
	Location LocationRef                      `json:"location"`
	Metadata map[string]any                   `json:"metadata"`
	JSON     v2EventSessionNextShellEndedJSON `json:"-"`
}

type v2EventSessionNextShellEndedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextShellEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextShellEndedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextShellEnded) implementsV2EventPayload() {}

type V2EventSessionNextShellEndedType string

const (
	V2EventSessionNextShellEndedTypeSessionNextShellEnded V2EventSessionNextShellEndedType = "session.next.shell.ended"
)

func (r V2EventSessionNextShellEndedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextShellEndedTypeSessionNextShellEnded:
		return true
	}
	return false
}

type V2EventSessionNextShellEndedData struct {
	CallID    string                               `json:"callID,required"`
	Output    string                               `json:"output,required"`
	SessionID string                               `json:"sessionID,required"`
	Timestamp int64                                `json:"timestamp,required"`
	JSON      v2EventSessionNextShellEndedDataJSON `json:"-"`
}

type v2EventSessionNextShellEndedDataJSON struct {
	CallID      apijson.Field
	Output      apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextShellEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextShellEndedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextShellStarted struct {
	ID       string                             `json:"id,required"`
	Data     V2EventSessionNextShellStartedData `json:"data,required"`
	Type     V2EventSessionNextShellStartedType `json:"type,required"`
	Durable  V2EventDurable                     `json:"durable"`
	Location LocationRef                        `json:"location"`
	Metadata map[string]any                     `json:"metadata"`
	JSON     v2EventSessionNextShellStartedJSON `json:"-"`
}

type v2EventSessionNextShellStartedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextShellStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextShellStartedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextShellStarted) implementsV2EventPayload() {}

type V2EventSessionNextShellStartedType string

const (
	V2EventSessionNextShellStartedTypeSessionNextShellStarted V2EventSessionNextShellStartedType = "session.next.shell.started"
)

func (r V2EventSessionNextShellStartedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextShellStartedTypeSessionNextShellStarted:
		return true
	}
	return false
}

type V2EventSessionNextShellStartedData struct {
	CallID    string                                 `json:"callID,required"`
	Command   string                                 `json:"command,required"`
	MessageID string                                 `json:"messageID,required"`
	SessionID string                                 `json:"sessionID,required"`
	Timestamp int64                                  `json:"timestamp,required"`
	JSON      v2EventSessionNextShellStartedDataJSON `json:"-"`
}

type v2EventSessionNextShellStartedDataJSON struct {
	CallID      apijson.Field
	Command     apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextShellStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextShellStartedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextStepEnded struct {
	ID       string                          `json:"id,required"`
	Data     V2EventSessionNextStepEndedData `json:"data,required"`
	Type     V2EventSessionNextStepEndedType `json:"type,required"`
	Durable  V2EventDurable                  `json:"durable"`
	Location LocationRef                     `json:"location"`
	Metadata map[string]any                  `json:"metadata"`
	JSON     v2EventSessionNextStepEndedJSON `json:"-"`
}

type v2EventSessionNextStepEndedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextStepEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextStepEndedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextStepEnded) implementsV2EventPayload() {}

type V2EventSessionNextStepEndedType string

const (
	V2EventSessionNextStepEndedTypeSessionNextStepEnded V2EventSessionNextStepEndedType = "session.next.step.ended"
)

func (r V2EventSessionNextStepEndedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextStepEndedTypeSessionNextStepEnded:
		return true
	}
	return false
}

type V2EventSessionNextStepEndedData struct {
	AssistantMessageID string                                `json:"assistantMessageID,required"`
	Cost               int64                                 `json:"cost,required"`
	Files              []string                              `json:"files"`
	Finish             string                                `json:"finish,required"`
	SessionID          string                                `json:"sessionID,required"`
	Snapshot           string                                `json:"snapshot"`
	Timestamp          int64                                 `json:"timestamp,required"`
	Tokens             V2EventSessionNextStepEndedDataTokens `json:"tokens,required"`
	JSON               v2EventSessionNextStepEndedDataJSON   `json:"-"`
}

type v2EventSessionNextStepEndedDataJSON struct {
	AssistantMessageID apijson.Field
	Cost               apijson.Field
	Files              apijson.Field
	Finish             apijson.Field
	SessionID          apijson.Field
	Snapshot           apijson.Field
	Timestamp          apijson.Field
	Tokens             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextStepEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextStepEndedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextStepEndedDataTokens struct {
	Cache     V2EventSessionNextStepEndedDataTokensCache `json:"cache,required"`
	Input     int64                                      `json:"input,required"`
	Output    int64                                      `json:"output,required"`
	Reasoning int64                                      `json:"reasoning,required"`
	JSON      V2EventSessionNextStepEndedDataTokensJSON  `json:"-"`
}

type V2EventSessionNextStepEndedDataTokensJSON struct {
	Cache       apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextStepEndedDataTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r V2EventSessionNextStepEndedDataTokensJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextStepEndedDataTokensCache struct {
	Read  int64                                          `json:"read,required"`
	Write int64                                          `json:"write,required"`
	JSON  V2EventSessionNextStepEndedDataTokensCacheJSON `json:"-"`
}

type V2EventSessionNextStepEndedDataTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextStepEndedDataTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r V2EventSessionNextStepEndedDataTokensCacheJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextStepFailed struct {
	ID       string                           `json:"id,required"`
	Data     V2EventSessionNextStepFailedData `json:"data,required"`
	Type     V2EventSessionNextStepFailedType `json:"type,required"`
	Durable  V2EventDurable                   `json:"durable"`
	Location LocationRef                      `json:"location"`
	Metadata map[string]any                   `json:"metadata"`
	JSON     v2EventSessionNextStepFailedJSON `json:"-"`
}

type v2EventSessionNextStepFailedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextStepFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextStepFailedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextStepFailed) implementsV2EventPayload() {}

type V2EventSessionNextStepFailedType string

const (
	V2EventSessionNextStepFailedTypeSessionNextStepFailed V2EventSessionNextStepFailedType = "session.next.step.failed"
)

func (r V2EventSessionNextStepFailedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextStepFailedTypeSessionNextStepFailed:
		return true
	}
	return false
}

type V2EventSessionNextStepFailedData struct {
	AssistantMessageID string                               `json:"assistantMessageID,required"`
	Error              SessionErrorUnknown                  `json:"error,required"`
	SessionID          string                               `json:"sessionID,required"`
	Timestamp          int64                                `json:"timestamp,required"`
	JSON               v2EventSessionNextStepFailedDataJSON `json:"-"`
}

type v2EventSessionNextStepFailedDataJSON struct {
	AssistantMessageID apijson.Field
	Error              apijson.Field
	SessionID          apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextStepFailedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextStepFailedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextStepStarted struct {
	ID       string                            `json:"id,required"`
	Data     V2EventSessionNextStepStartedData `json:"data,required"`
	Type     V2EventSessionNextStepStartedType `json:"type,required"`
	Durable  V2EventDurable                    `json:"durable"`
	Location LocationRef                       `json:"location"`
	Metadata map[string]any                    `json:"metadata"`
	JSON     v2EventSessionNextStepStartedJSON `json:"-"`
}

type v2EventSessionNextStepStartedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextStepStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextStepStartedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextStepStarted) implementsV2EventPayload() {}

type V2EventSessionNextStepStartedType string

const (
	V2EventSessionNextStepStartedTypeSessionNextStepStarted V2EventSessionNextStepStartedType = "session.next.step.started"
)

func (r V2EventSessionNextStepStartedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextStepStartedTypeSessionNextStepStarted:
		return true
	}
	return false
}

type V2EventSessionNextStepStartedData struct {
	Agent              string                                `json:"agent,required"`
	AssistantMessageID string                                `json:"assistantMessageID,required"`
	Model              ModelRef                              `json:"model,required"`
	SessionID          string                                `json:"sessionID,required"`
	Snapshot           string                                `json:"snapshot"`
	Timestamp          int64                                 `json:"timestamp,required"`
	JSON               v2EventSessionNextStepStartedDataJSON `json:"-"`
}

type v2EventSessionNextStepStartedDataJSON struct {
	Agent              apijson.Field
	AssistantMessageID apijson.Field
	Model              apijson.Field
	SessionID          apijson.Field
	Snapshot           apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextStepStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextStepStartedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextSynthetic struct {
	ID       string                          `json:"id,required"`
	Data     V2EventSessionNextSyntheticData `json:"data,required"`
	Type     V2EventSessionNextSyntheticType `json:"type,required"`
	Durable  V2EventDurable                  `json:"durable"`
	Location LocationRef                     `json:"location"`
	Metadata map[string]any                  `json:"metadata"`
	JSON     v2EventSessionNextSyntheticJSON `json:"-"`
}

type v2EventSessionNextSyntheticJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextSynthetic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextSyntheticJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextSynthetic) implementsV2EventPayload() {}

type V2EventSessionNextSyntheticType string

const (
	V2EventSessionNextSyntheticTypeSessionNextSynthetic V2EventSessionNextSyntheticType = "session.next.synthetic"
)

func (r V2EventSessionNextSyntheticType) IsKnown() bool {
	switch r {
	case V2EventSessionNextSyntheticTypeSessionNextSynthetic:
		return true
	}
	return false
}

type V2EventSessionNextSyntheticData struct {
	MessageID string                              `json:"messageID,required"`
	SessionID string                              `json:"sessionID,required"`
	Text      string                              `json:"text,required"`
	Timestamp int64                               `json:"timestamp,required"`
	JSON      v2EventSessionNextSyntheticDataJSON `json:"-"`
}

type v2EventSessionNextSyntheticDataJSON struct {
	MessageID   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextSyntheticData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextSyntheticDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextTextDelta struct {
	ID       string                          `json:"id,required"`
	Data     V2EventSessionNextTextDeltaData `json:"data,required"`
	Type     V2EventSessionNextTextDeltaType `json:"type,required"`
	Durable  V2EventDurable                  `json:"durable"`
	Location LocationRef                     `json:"location"`
	Metadata map[string]any                  `json:"metadata"`
	JSON     v2EventSessionNextTextDeltaJSON `json:"-"`
}

type v2EventSessionNextTextDeltaJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextTextDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextTextDeltaJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextTextDelta) implementsV2EventPayload() {}

type V2EventSessionNextTextDeltaType string

const (
	V2EventSessionNextTextDeltaTypeSessionNextTextDelta V2EventSessionNextTextDeltaType = "session.next.text.delta"
)

func (r V2EventSessionNextTextDeltaType) IsKnown() bool {
	switch r {
	case V2EventSessionNextTextDeltaTypeSessionNextTextDelta:
		return true
	}
	return false
}

type V2EventSessionNextTextDeltaData struct {
	AssistantMessageID string                              `json:"assistantMessageID,required"`
	Delta              string                              `json:"delta,required"`
	SessionID          string                              `json:"sessionID,required"`
	TextID             string                              `json:"textID,required"`
	Timestamp          int64                               `json:"timestamp,required"`
	JSON               v2EventSessionNextTextDeltaDataJSON `json:"-"`
}

type v2EventSessionNextTextDeltaDataJSON struct {
	AssistantMessageID apijson.Field
	Delta              apijson.Field
	SessionID          apijson.Field
	TextID             apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextTextDeltaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextTextDeltaDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextTextEnded struct {
	ID       string                          `json:"id,required"`
	Data     V2EventSessionNextTextEndedData `json:"data,required"`
	Type     V2EventSessionNextTextEndedType `json:"type,required"`
	Durable  V2EventDurable                  `json:"durable"`
	Location LocationRef                     `json:"location"`
	Metadata map[string]any                  `json:"metadata"`
	JSON     v2EventSessionNextTextEndedJSON `json:"-"`
}

type v2EventSessionNextTextEndedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextTextEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextTextEndedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextTextEnded) implementsV2EventPayload() {}

type V2EventSessionNextTextEndedType string

const (
	V2EventSessionNextTextEndedTypeSessionNextTextEnded V2EventSessionNextTextEndedType = "session.next.text.ended"
)

func (r V2EventSessionNextTextEndedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextTextEndedTypeSessionNextTextEnded:
		return true
	}
	return false
}

type V2EventSessionNextTextEndedData struct {
	AssistantMessageID string                              `json:"assistantMessageID,required"`
	SessionID          string                              `json:"sessionID,required"`
	Text               string                              `json:"text,required"`
	TextID             string                              `json:"textID,required"`
	Timestamp          int64                               `json:"timestamp,required"`
	JSON               v2EventSessionNextTextEndedDataJSON `json:"-"`
}

type v2EventSessionNextTextEndedDataJSON struct {
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	Text               apijson.Field
	TextID             apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextTextEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextTextEndedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextTextStarted struct {
	ID       string                            `json:"id,required"`
	Data     V2EventSessionNextTextStartedData `json:"data,required"`
	Type     V2EventSessionNextTextStartedType `json:"type,required"`
	Durable  V2EventDurable                    `json:"durable"`
	Location LocationRef                       `json:"location"`
	Metadata map[string]any                    `json:"metadata"`
	JSON     v2EventSessionNextTextStartedJSON `json:"-"`
}

type v2EventSessionNextTextStartedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextTextStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextTextStartedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextTextStarted) implementsV2EventPayload() {}

type V2EventSessionNextTextStartedType string

const (
	V2EventSessionNextTextStartedTypeSessionNextTextStarted V2EventSessionNextTextStartedType = "session.next.text.started"
)

func (r V2EventSessionNextTextStartedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextTextStartedTypeSessionNextTextStarted:
		return true
	}
	return false
}

type V2EventSessionNextTextStartedData struct {
	AssistantMessageID string                                `json:"assistantMessageID,required"`
	SessionID          string                                `json:"sessionID,required"`
	TextID             string                                `json:"textID,required"`
	Timestamp          int64                                 `json:"timestamp,required"`
	JSON               v2EventSessionNextTextStartedDataJSON `json:"-"`
}

type v2EventSessionNextTextStartedDataJSON struct {
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	TextID             apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextTextStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextTextStartedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolCalled struct {
	ID       string                           `json:"id,required"`
	Data     V2EventSessionNextToolCalledData `json:"data,required"`
	Type     V2EventSessionNextToolCalledType `json:"type,required"`
	Durable  V2EventDurable                   `json:"durable"`
	Location LocationRef                      `json:"location"`
	Metadata map[string]any                   `json:"metadata"`
	JSON     v2EventSessionNextToolCalledJSON `json:"-"`
}

type v2EventSessionNextToolCalledJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolCalled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolCalledJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextToolCalled) implementsV2EventPayload() {}

type V2EventSessionNextToolCalledType string

const (
	V2EventSessionNextToolCalledTypeSessionNextToolCalled V2EventSessionNextToolCalledType = "session.next.tool.called"
)

func (r V2EventSessionNextToolCalledType) IsKnown() bool {
	switch r {
	case V2EventSessionNextToolCalledTypeSessionNextToolCalled:
		return true
	}
	return false
}

type V2EventSessionNextToolCalledData struct {
	AssistantMessageID string `json:"assistantMessageID,required"`
	CallID             string `json:"callID,required"`
	// This field can have the runtime type of [map[string]any].
	Input     any                                      `json:"input,required"`
	Provider  V2EventSessionNextToolCalledDataProvider `json:"provider,required"`
	SessionID string                                   `json:"sessionID,required"`
	Timestamp int64                                    `json:"timestamp,required"`
	Tool      string                                   `json:"tool,required"`
	JSON      v2EventSessionNextToolCalledDataJSON     `json:"-"`
}

type v2EventSessionNextToolCalledDataJSON struct {
	AssistantMessageID apijson.Field
	CallID             apijson.Field
	Input              apijson.Field
	Provider           apijson.Field
	SessionID          apijson.Field
	Timestamp          apijson.Field
	Tool               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextToolCalledData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolCalledDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolCalledDataProvider struct {
	Executed bool `json:"executed,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                          `json:"metadata"`
	JSON     V2EventSessionNextToolCalledDataProviderJSON `json:"-"`
}

type V2EventSessionNextToolCalledDataProviderJSON struct {
	Executed    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolCalledDataProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r V2EventSessionNextToolCalledDataProviderJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolFailed struct {
	ID       string                           `json:"id,required"`
	Data     V2EventSessionNextToolFailedData `json:"data,required"`
	Type     V2EventSessionNextToolFailedType `json:"type,required"`
	Durable  V2EventDurable                   `json:"durable"`
	Location LocationRef                      `json:"location"`
	Metadata map[string]any                   `json:"metadata"`
	JSON     v2EventSessionNextToolFailedJSON `json:"-"`
}

type v2EventSessionNextToolFailedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolFailedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextToolFailed) implementsV2EventPayload() {}

type V2EventSessionNextToolFailedType string

const (
	V2EventSessionNextToolFailedTypeSessionNextToolFailed V2EventSessionNextToolFailedType = "session.next.tool.failed"
)

func (r V2EventSessionNextToolFailedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextToolFailedTypeSessionNextToolFailed:
		return true
	}
	return false
}

type V2EventSessionNextToolFailedData struct {
	AssistantMessageID string                                   `json:"assistantMessageID,required"`
	CallID             string                                   `json:"callID,required"`
	Error              SessionErrorUnknown                      `json:"error,required"`
	Provider           V2EventSessionNextToolFailedDataProvider `json:"provider,required"`
	// This field can have the runtime type of [any].
	Result    any                                  `json:"result"`
	SessionID string                               `json:"sessionID,required"`
	Timestamp int64                                `json:"timestamp,required"`
	JSON      v2EventSessionNextToolFailedDataJSON `json:"-"`
}

type v2EventSessionNextToolFailedDataJSON struct {
	AssistantMessageID apijson.Field
	CallID             apijson.Field
	Error              apijson.Field
	Provider           apijson.Field
	Result             apijson.Field
	SessionID          apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextToolFailedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolFailedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolFailedDataProvider struct {
	Executed bool `json:"executed,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                          `json:"metadata"`
	JSON     V2EventSessionNextToolFailedDataProviderJSON `json:"-"`
}

type V2EventSessionNextToolFailedDataProviderJSON struct {
	Executed    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolFailedDataProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r V2EventSessionNextToolFailedDataProviderJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolInputDelta struct {
	ID       string                               `json:"id,required"`
	Data     V2EventSessionNextToolInputDeltaData `json:"data,required"`
	Type     V2EventSessionNextToolInputDeltaType `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Metadata map[string]any                       `json:"metadata"`
	JSON     v2EventSessionNextToolInputDeltaJSON `json:"-"`
}

type v2EventSessionNextToolInputDeltaJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolInputDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolInputDeltaJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextToolInputDelta) implementsV2EventPayload() {}

type V2EventSessionNextToolInputDeltaType string

const (
	V2EventSessionNextToolInputDeltaTypeSessionNextToolInputDelta V2EventSessionNextToolInputDeltaType = "session.next.tool.input.delta"
)

func (r V2EventSessionNextToolInputDeltaType) IsKnown() bool {
	switch r {
	case V2EventSessionNextToolInputDeltaTypeSessionNextToolInputDelta:
		return true
	}
	return false
}

type V2EventSessionNextToolInputDeltaData struct {
	AssistantMessageID string                                   `json:"assistantMessageID,required"`
	CallID             string                                   `json:"callID,required"`
	Delta              string                                   `json:"delta,required"`
	SessionID          string                                   `json:"sessionID,required"`
	Timestamp          int64                                    `json:"timestamp,required"`
	JSON               v2EventSessionNextToolInputDeltaDataJSON `json:"-"`
}

type v2EventSessionNextToolInputDeltaDataJSON struct {
	AssistantMessageID apijson.Field
	CallID             apijson.Field
	Delta              apijson.Field
	SessionID          apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextToolInputDeltaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolInputDeltaDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolInputEnded struct {
	ID       string                               `json:"id,required"`
	Data     V2EventSessionNextToolInputEndedData `json:"data,required"`
	Type     V2EventSessionNextToolInputEndedType `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Metadata map[string]any                       `json:"metadata"`
	JSON     v2EventSessionNextToolInputEndedJSON `json:"-"`
}

type v2EventSessionNextToolInputEndedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolInputEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolInputEndedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextToolInputEnded) implementsV2EventPayload() {}

type V2EventSessionNextToolInputEndedType string

const (
	V2EventSessionNextToolInputEndedTypeSessionNextToolInputEnded V2EventSessionNextToolInputEndedType = "session.next.tool.input.ended"
)

func (r V2EventSessionNextToolInputEndedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextToolInputEndedTypeSessionNextToolInputEnded:
		return true
	}
	return false
}

type V2EventSessionNextToolInputEndedData struct {
	AssistantMessageID string                                   `json:"assistantMessageID,required"`
	CallID             string                                   `json:"callID,required"`
	SessionID          string                                   `json:"sessionID,required"`
	Text               string                                   `json:"text,required"`
	Timestamp          int64                                    `json:"timestamp,required"`
	JSON               v2EventSessionNextToolInputEndedDataJSON `json:"-"`
}

type v2EventSessionNextToolInputEndedDataJSON struct {
	AssistantMessageID apijson.Field
	CallID             apijson.Field
	SessionID          apijson.Field
	Text               apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextToolInputEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolInputEndedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolInputStarted struct {
	ID       string                                 `json:"id,required"`
	Data     V2EventSessionNextToolInputStartedData `json:"data,required"`
	Type     V2EventSessionNextToolInputStartedType `json:"type,required"`
	Durable  V2EventDurable                         `json:"durable"`
	Location LocationRef                            `json:"location"`
	Metadata map[string]any                         `json:"metadata"`
	JSON     v2EventSessionNextToolInputStartedJSON `json:"-"`
}

type v2EventSessionNextToolInputStartedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolInputStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolInputStartedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextToolInputStarted) implementsV2EventPayload() {}

type V2EventSessionNextToolInputStartedType string

const (
	V2EventSessionNextToolInputStartedTypeSessionNextToolInputStarted V2EventSessionNextToolInputStartedType = "session.next.tool.input.started"
)

func (r V2EventSessionNextToolInputStartedType) IsKnown() bool {
	switch r {
	case V2EventSessionNextToolInputStartedTypeSessionNextToolInputStarted:
		return true
	}
	return false
}

type V2EventSessionNextToolInputStartedData struct {
	AssistantMessageID string                                     `json:"assistantMessageID,required"`
	CallID             string                                     `json:"callID,required"`
	Name               string                                     `json:"name,required"`
	SessionID          string                                     `json:"sessionID,required"`
	Timestamp          int64                                      `json:"timestamp,required"`
	JSON               v2EventSessionNextToolInputStartedDataJSON `json:"-"`
}

type v2EventSessionNextToolInputStartedDataJSON struct {
	AssistantMessageID apijson.Field
	CallID             apijson.Field
	Name               apijson.Field
	SessionID          apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextToolInputStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolInputStartedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolProgress struct {
	ID       string                             `json:"id,required"`
	Data     V2EventSessionNextToolProgressData `json:"data,required"`
	Type     V2EventSessionNextToolProgressType `json:"type,required"`
	Durable  V2EventDurable                     `json:"durable"`
	Location LocationRef                        `json:"location"`
	Metadata map[string]any                     `json:"metadata"`
	JSON     v2EventSessionNextToolProgressJSON `json:"-"`
}

type v2EventSessionNextToolProgressJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolProgressJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextToolProgress) implementsV2EventPayload() {}

type V2EventSessionNextToolProgressType string

const (
	V2EventSessionNextToolProgressTypeSessionNextToolProgress V2EventSessionNextToolProgressType = "session.next.tool.progress"
)

func (r V2EventSessionNextToolProgressType) IsKnown() bool {
	switch r {
	case V2EventSessionNextToolProgressTypeSessionNextToolProgress:
		return true
	}
	return false
}

type V2EventSessionNextToolProgressData struct {
	AssistantMessageID string `json:"assistantMessageID,required"`
	CallID             string `json:"callID,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content   []any  `json:"content,required"`
	SessionID string `json:"sessionID,required"`
	// This field can have the runtime type of [map[string]any].
	Structured any                                    `json:"structured,required"`
	Timestamp  int64                                  `json:"timestamp,required"`
	JSON       v2EventSessionNextToolProgressDataJSON `json:"-"`
}

type v2EventSessionNextToolProgressDataJSON struct {
	AssistantMessageID apijson.Field
	CallID             apijson.Field
	Content            apijson.Field
	SessionID          apijson.Field
	Structured         apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextToolProgressData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolProgressDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolSuccess struct {
	ID       string                            `json:"id,required"`
	Data     V2EventSessionNextToolSuccessData `json:"data,required"`
	Type     V2EventSessionNextToolSuccessType `json:"type,required"`
	Durable  V2EventDurable                    `json:"durable"`
	Location LocationRef                       `json:"location"`
	Metadata map[string]any                    `json:"metadata"`
	JSON     v2EventSessionNextToolSuccessJSON `json:"-"`
}

type v2EventSessionNextToolSuccessJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolSuccessJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionNextToolSuccess) implementsV2EventPayload() {}

type V2EventSessionNextToolSuccessType string

const (
	V2EventSessionNextToolSuccessTypeSessionNextToolSuccess V2EventSessionNextToolSuccessType = "session.next.tool.success"
)

func (r V2EventSessionNextToolSuccessType) IsKnown() bool {
	switch r {
	case V2EventSessionNextToolSuccessTypeSessionNextToolSuccess:
		return true
	}
	return false
}

type V2EventSessionNextToolSuccessData struct {
	AssistantMessageID string `json:"assistantMessageID,required"`
	CallID             string `json:"callID,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content     []any                                     `json:"content,required"`
	OutputPaths []string                                  `json:"outputPaths"`
	Provider    V2EventSessionNextToolSuccessDataProvider `json:"provider,required"`
	// This field can have the runtime type of [any].
	Result    any    `json:"result"`
	SessionID string `json:"sessionID,required"`
	// This field can have the runtime type of [map[string]any].
	Structured any                                   `json:"structured,required"`
	Timestamp  int64                                 `json:"timestamp,required"`
	JSON       v2EventSessionNextToolSuccessDataJSON `json:"-"`
}

type v2EventSessionNextToolSuccessDataJSON struct {
	AssistantMessageID apijson.Field
	CallID             apijson.Field
	Content            apijson.Field
	OutputPaths        apijson.Field
	Provider           apijson.Field
	Result             apijson.Field
	SessionID          apijson.Field
	Structured         apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V2EventSessionNextToolSuccessData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionNextToolSuccessDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionNextToolSuccessDataProvider struct {
	Executed bool `json:"executed,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                           `json:"metadata"`
	JSON     V2EventSessionNextToolSuccessDataProviderJSON `json:"-"`
}

type V2EventSessionNextToolSuccessDataProviderJSON struct {
	Executed    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionNextToolSuccessDataProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r V2EventSessionNextToolSuccessDataProviderJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionStatus struct {
	ID       string                   `json:"id,required"`
	Data     V2EventSessionStatusData `json:"data,required"`
	Type     V2EventSessionStatusType `json:"type,required"`
	Durable  V2EventDurable           `json:"durable"`
	Location LocationRef              `json:"location"`
	Metadata map[string]any           `json:"metadata"`
	JSON     v2EventSessionStatusJSON `json:"-"`
}

type v2EventSessionStatusJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionStatusJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionStatus) implementsV2EventPayload() {}

type V2EventSessionStatusType string

const (
	V2EventSessionStatusTypeSessionStatus V2EventSessionStatusType = "session.status"
)

func (r V2EventSessionStatusType) IsKnown() bool {
	switch r {
	case V2EventSessionStatusTypeSessionStatus:
		return true
	}
	return false
}

type V2EventSessionStatusData struct {
	SessionID string                       `json:"sessionID,required"`
	Status    SessionStatus                `json:"status,required"`
	JSON      v2EventSessionStatusDataJSON `json:"-"`
}

type v2EventSessionStatusDataJSON struct {
	SessionID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionStatusData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionStatusDataJSON) RawJSON() string {
	return r.raw
}

type V2EventSessionUpdated struct {
	ID       string                    `json:"id,required"`
	Data     V2EventSessionUpdatedData `json:"data,required"`
	Type     V2EventSessionUpdatedType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventSessionUpdatedJSON `json:"-"`
}

type v2EventSessionUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventSessionUpdated) implementsV2EventPayload() {}

type V2EventSessionUpdatedType string

const (
	V2EventSessionUpdatedTypeSessionUpdated V2EventSessionUpdatedType = "session.updated"
)

func (r V2EventSessionUpdatedType) IsKnown() bool {
	switch r {
	case V2EventSessionUpdatedTypeSessionUpdated:
		return true
	}
	return false
}

type V2EventSessionUpdatedData struct {
	Info      Session                       `json:"info,required"`
	SessionID string                        `json:"sessionID,required"`
	JSON      v2EventSessionUpdatedDataJSON `json:"-"`
}

type v2EventSessionUpdatedDataJSON struct {
	Info        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventSessionUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventSessionUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventTodoUpdated struct {
	ID       string                 `json:"id,required"`
	Data     V2EventTodoUpdatedData `json:"data,required"`
	Type     V2EventTodoUpdatedType `json:"type,required"`
	Durable  V2EventDurable         `json:"durable"`
	Location LocationRef            `json:"location"`
	Metadata map[string]any         `json:"metadata"`
	JSON     v2EventTodoUpdatedJSON `json:"-"`
}

type v2EventTodoUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTodoUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTodoUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventTodoUpdated) implementsV2EventPayload() {}

type V2EventTodoUpdatedType string

const (
	V2EventTodoUpdatedTypeTodoUpdated V2EventTodoUpdatedType = "todo.updated"
)

func (r V2EventTodoUpdatedType) IsKnown() bool {
	switch r {
	case V2EventTodoUpdatedTypeTodoUpdated:
		return true
	}
	return false
}

type V2EventTodoUpdatedData struct {
	SessionID string                     `json:"sessionID,required"`
	Todos     []Todo                     `json:"todos,required"`
	JSON      v2EventTodoUpdatedDataJSON `json:"-"`
}

type v2EventTodoUpdatedDataJSON struct {
	SessionID   apijson.Field
	Todos       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTodoUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTodoUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventTuiCommandExecute struct {
	ID       string                       `json:"id,required"`
	Data     V2EventTuiCommandExecuteData `json:"data,required"`
	Type     V2EventTuiCommandExecuteType `json:"type,required"`
	Durable  V2EventDurable               `json:"durable"`
	Location LocationRef                  `json:"location"`
	Metadata map[string]any               `json:"metadata"`
	JSON     v2EventTuiCommandExecuteJSON `json:"-"`
}

type v2EventTuiCommandExecuteJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTuiCommandExecute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTuiCommandExecuteJSON) RawJSON() string {
	return r.raw
}

func (r V2EventTuiCommandExecute) implementsV2EventPayload() {}

type V2EventTuiCommandExecuteType string

const (
	V2EventTuiCommandExecuteTypeTuiCommandExecute V2EventTuiCommandExecuteType = "tui.command.execute"
)

func (r V2EventTuiCommandExecuteType) IsKnown() bool {
	switch r {
	case V2EventTuiCommandExecuteTypeTuiCommandExecute:
		return true
	}
	return false
}

type V2EventTuiCommandExecuteData struct {
	// This field can have the runtime type of [string].
	Command any                              `json:"command,required"`
	JSON    v2EventTuiCommandExecuteDataJSON `json:"-"`
}

type v2EventTuiCommandExecuteDataJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTuiCommandExecuteData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTuiCommandExecuteDataJSON) RawJSON() string {
	return r.raw
}

type V2EventTuiPromptAppend struct {
	ID       string                     `json:"id,required"`
	Data     V2EventTuiPromptAppendData `json:"data,required"`
	Type     V2EventTuiPromptAppendType `json:"type,required"`
	Durable  V2EventDurable             `json:"durable"`
	Location LocationRef                `json:"location"`
	Metadata map[string]any             `json:"metadata"`
	JSON     v2EventTuiPromptAppendJSON `json:"-"`
}

type v2EventTuiPromptAppendJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTuiPromptAppend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTuiPromptAppendJSON) RawJSON() string {
	return r.raw
}

func (r V2EventTuiPromptAppend) implementsV2EventPayload() {}

type V2EventTuiPromptAppendType string

const (
	V2EventTuiPromptAppendTypeTuiPromptAppend V2EventTuiPromptAppendType = "tui.prompt.append"
)

func (r V2EventTuiPromptAppendType) IsKnown() bool {
	switch r {
	case V2EventTuiPromptAppendTypeTuiPromptAppend:
		return true
	}
	return false
}

type V2EventTuiPromptAppendData struct {
	Text string                         `json:"text,required"`
	JSON v2EventTuiPromptAppendDataJSON `json:"-"`
}

type v2EventTuiPromptAppendDataJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTuiPromptAppendData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTuiPromptAppendDataJSON) RawJSON() string {
	return r.raw
}

type V2EventTuiSessionSelect struct {
	ID       string                      `json:"id,required"`
	Data     V2EventTuiSessionSelectData `json:"data,required"`
	Type     V2EventTuiSessionSelectType `json:"type,required"`
	Durable  V2EventDurable              `json:"durable"`
	Location LocationRef                 `json:"location"`
	Metadata map[string]any              `json:"metadata"`
	JSON     v2EventTuiSessionSelectJSON `json:"-"`
}

type v2EventTuiSessionSelectJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTuiSessionSelect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTuiSessionSelectJSON) RawJSON() string {
	return r.raw
}

func (r V2EventTuiSessionSelect) implementsV2EventPayload() {}

type V2EventTuiSessionSelectType string

const (
	V2EventTuiSessionSelectTypeTuiSessionSelect V2EventTuiSessionSelectType = "tui.session.select"
)

func (r V2EventTuiSessionSelectType) IsKnown() bool {
	switch r {
	case V2EventTuiSessionSelectTypeTuiSessionSelect:
		return true
	}
	return false
}

type V2EventTuiSessionSelectData struct {
	SessionID string                          `json:"sessionID,required"`
	JSON      v2EventTuiSessionSelectDataJSON `json:"-"`
}

type v2EventTuiSessionSelectDataJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTuiSessionSelectData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTuiSessionSelectDataJSON) RawJSON() string {
	return r.raw
}

type V2EventTuiToastShowVariant string

const (
	V2EventTuiToastShowVariantInfo    V2EventTuiToastShowVariant = "info"
	V2EventTuiToastShowVariantSuccess V2EventTuiToastShowVariant = "success"
	V2EventTuiToastShowVariantWarning V2EventTuiToastShowVariant = "warning"
	V2EventTuiToastShowVariantError   V2EventTuiToastShowVariant = "error"
)

func (r V2EventTuiToastShowVariant) IsKnown() bool {
	switch r {
	case V2EventTuiToastShowVariantInfo,
		V2EventTuiToastShowVariantSuccess,
		V2EventTuiToastShowVariantWarning,
		V2EventTuiToastShowVariantError:
		return true
	}
	return false
}

type V2EventTuiToastShow struct {
	ID       string                  `json:"id,required"`
	Data     V2EventTuiToastShowData `json:"data,required"`
	Type     V2EventTuiToastShowType `json:"type,required"`
	Durable  V2EventDurable          `json:"durable"`
	Location LocationRef             `json:"location"`
	Metadata map[string]any          `json:"metadata"`
	JSON     v2EventTuiToastShowJSON `json:"-"`
}

type v2EventTuiToastShowJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTuiToastShow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTuiToastShowJSON) RawJSON() string {
	return r.raw
}

func (r V2EventTuiToastShow) implementsV2EventPayload() {}

type V2EventTuiToastShowType string

const (
	V2EventTuiToastShowTypeTuiToastShow V2EventTuiToastShowType = "tui.toast.show"
)

func (r V2EventTuiToastShowType) IsKnown() bool {
	switch r {
	case V2EventTuiToastShowTypeTuiToastShow:
		return true
	}
	return false
}

type V2EventTuiToastShowData struct {
	Duration int64                       `json:"duration"`
	Message  string                      `json:"message,required"`
	Title    string                      `json:"title"`
	Variant  V2EventTuiToastShowVariant  `json:"variant,required"`
	JSON     v2EventTuiToastShowDataJSON `json:"-"`
}

type v2EventTuiToastShowDataJSON struct {
	Duration    apijson.Field
	Message     apijson.Field
	Title       apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventTuiToastShowData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventTuiToastShowDataJSON) RawJSON() string {
	return r.raw
}

type V2EventVcsBranchUpdated struct {
	ID       string                      `json:"id,required"`
	Data     V2EventVcsBranchUpdatedData `json:"data,required"`
	Type     V2EventVcsBranchUpdatedType `json:"type,required"`
	Durable  V2EventDurable              `json:"durable"`
	Location LocationRef                 `json:"location"`
	Metadata map[string]any              `json:"metadata"`
	JSON     v2EventVcsBranchUpdatedJSON `json:"-"`
}

type v2EventVcsBranchUpdatedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventVcsBranchUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventVcsBranchUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventVcsBranchUpdated) implementsV2EventPayload() {}

type V2EventVcsBranchUpdatedType string

const (
	V2EventVcsBranchUpdatedTypeVcsBranchUpdated V2EventVcsBranchUpdatedType = "vcs.branch.updated"
)

func (r V2EventVcsBranchUpdatedType) IsKnown() bool {
	switch r {
	case V2EventVcsBranchUpdatedTypeVcsBranchUpdated:
		return true
	}
	return false
}

type V2EventVcsBranchUpdatedData struct {
	Branch string                          `json:"branch"`
	JSON   v2EventVcsBranchUpdatedDataJSON `json:"-"`
}

type v2EventVcsBranchUpdatedDataJSON struct {
	Branch      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventVcsBranchUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventVcsBranchUpdatedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventWorkspaceFailed struct {
	ID       string                     `json:"id,required"`
	Data     V2EventWorkspaceFailedData `json:"data,required"`
	Type     V2EventWorkspaceFailedType `json:"type,required"`
	Durable  V2EventDurable             `json:"durable"`
	Location LocationRef                `json:"location"`
	Metadata map[string]any             `json:"metadata"`
	JSON     v2EventWorkspaceFailedJSON `json:"-"`
}

type v2EventWorkspaceFailedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorkspaceFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorkspaceFailedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventWorkspaceFailed) implementsV2EventPayload() {}

type V2EventWorkspaceFailedType string

const (
	V2EventWorkspaceFailedTypeWorkspaceFailed V2EventWorkspaceFailedType = "workspace.failed"
)

func (r V2EventWorkspaceFailedType) IsKnown() bool {
	switch r {
	case V2EventWorkspaceFailedTypeWorkspaceFailed:
		return true
	}
	return false
}

type V2EventWorkspaceFailedData struct {
	Message string                         `json:"message,required"`
	JSON    v2EventWorkspaceFailedDataJSON `json:"-"`
}

type v2EventWorkspaceFailedDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorkspaceFailedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorkspaceFailedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventWorkspaceReady struct {
	ID       string                    `json:"id,required"`
	Data     V2EventWorkspaceReadyData `json:"data,required"`
	Type     V2EventWorkspaceReadyType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventWorkspaceReadyJSON `json:"-"`
}

type v2EventWorkspaceReadyJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorkspaceReady) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorkspaceReadyJSON) RawJSON() string {
	return r.raw
}

func (r V2EventWorkspaceReady) implementsV2EventPayload() {}

type V2EventWorkspaceReadyType string

const (
	V2EventWorkspaceReadyTypeWorkspaceReady V2EventWorkspaceReadyType = "workspace.ready"
)

func (r V2EventWorkspaceReadyType) IsKnown() bool {
	switch r {
	case V2EventWorkspaceReadyTypeWorkspaceReady:
		return true
	}
	return false
}

type V2EventWorkspaceReadyData struct {
	Name string                        `json:"name,required"`
	JSON v2EventWorkspaceReadyDataJSON `json:"-"`
}

type v2EventWorkspaceReadyDataJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorkspaceReadyData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorkspaceReadyDataJSON) RawJSON() string {
	return r.raw
}

type V2EventWorkspaceStatusStatus string

const (
	V2EventWorkspaceStatusStatusConnected    V2EventWorkspaceStatusStatus = "connected"
	V2EventWorkspaceStatusStatusConnecting   V2EventWorkspaceStatusStatus = "connecting"
	V2EventWorkspaceStatusStatusDisconnected V2EventWorkspaceStatusStatus = "disconnected"
	V2EventWorkspaceStatusStatusError        V2EventWorkspaceStatusStatus = "error"
)

func (r V2EventWorkspaceStatusStatus) IsKnown() bool {
	switch r {
	case V2EventWorkspaceStatusStatusConnected,
		V2EventWorkspaceStatusStatusConnecting,
		V2EventWorkspaceStatusStatusDisconnected,
		V2EventWorkspaceStatusStatusError:
		return true
	}
	return false
}

type V2EventWorkspaceStatus struct {
	ID       string                     `json:"id,required"`
	Data     V2EventWorkspaceStatusData `json:"data,required"`
	Type     V2EventWorkspaceStatusType `json:"type,required"`
	Durable  V2EventDurable             `json:"durable"`
	Location LocationRef                `json:"location"`
	Metadata map[string]any             `json:"metadata"`
	JSON     v2EventWorkspaceStatusJSON `json:"-"`
}

type v2EventWorkspaceStatusJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorkspaceStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorkspaceStatusJSON) RawJSON() string {
	return r.raw
}

func (r V2EventWorkspaceStatus) implementsV2EventPayload() {}

type V2EventWorkspaceStatusType string

const (
	V2EventWorkspaceStatusTypeWorkspaceStatus V2EventWorkspaceStatusType = "workspace.status"
)

func (r V2EventWorkspaceStatusType) IsKnown() bool {
	switch r {
	case V2EventWorkspaceStatusTypeWorkspaceStatus:
		return true
	}
	return false
}

type V2EventWorkspaceStatusData struct {
	Status      V2EventWorkspaceStatusStatus   `json:"status,required"`
	WorkspaceID string                         `json:"workspaceID,required"`
	JSON        v2EventWorkspaceStatusDataJSON `json:"-"`
}

type v2EventWorkspaceStatusDataJSON struct {
	Status      apijson.Field
	WorkspaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorkspaceStatusData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorkspaceStatusDataJSON) RawJSON() string {
	return r.raw
}

type V2EventWorktreeFailed struct {
	ID       string                    `json:"id,required"`
	Data     V2EventWorktreeFailedData `json:"data,required"`
	Type     V2EventWorktreeFailedType `json:"type,required"`
	Durable  V2EventDurable            `json:"durable"`
	Location LocationRef               `json:"location"`
	Metadata map[string]any            `json:"metadata"`
	JSON     v2EventWorktreeFailedJSON `json:"-"`
}

type v2EventWorktreeFailedJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorktreeFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorktreeFailedJSON) RawJSON() string {
	return r.raw
}

func (r V2EventWorktreeFailed) implementsV2EventPayload() {}

type V2EventWorktreeFailedType string

const (
	V2EventWorktreeFailedTypeWorktreeFailed V2EventWorktreeFailedType = "worktree.failed"
)

func (r V2EventWorktreeFailedType) IsKnown() bool {
	switch r {
	case V2EventWorktreeFailedTypeWorktreeFailed:
		return true
	}
	return false
}

type V2EventWorktreeFailedData struct {
	Message string                        `json:"message,required"`
	JSON    v2EventWorktreeFailedDataJSON `json:"-"`
}

type v2EventWorktreeFailedDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorktreeFailedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorktreeFailedDataJSON) RawJSON() string {
	return r.raw
}

type V2EventWorktreeReady struct {
	ID       string                   `json:"id,required"`
	Data     V2EventWorktreeReadyData `json:"data,required"`
	Type     V2EventWorktreeReadyType `json:"type,required"`
	Durable  V2EventDurable           `json:"durable"`
	Location LocationRef              `json:"location"`
	Metadata map[string]any           `json:"metadata"`
	JSON     v2EventWorktreeReadyJSON `json:"-"`
}

type v2EventWorktreeReadyJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorktreeReady) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorktreeReadyJSON) RawJSON() string {
	return r.raw
}

func (r V2EventWorktreeReady) implementsV2EventPayload() {}

type V2EventWorktreeReadyType string

const (
	V2EventWorktreeReadyTypeWorktreeReady V2EventWorktreeReadyType = "worktree.ready"
)

func (r V2EventWorktreeReadyType) IsKnown() bool {
	switch r {
	case V2EventWorktreeReadyTypeWorktreeReady:
		return true
	}
	return false
}

type V2EventWorktreeReadyData struct {
	Branch string                       `json:"branch"`
	Name   string                       `json:"name,required"`
	JSON   v2EventWorktreeReadyDataJSON `json:"-"`
}

type v2EventWorktreeReadyDataJSON struct {
	Branch      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2EventWorktreeReadyData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2EventWorktreeReadyDataJSON) RawJSON() string {
	return r.raw
}

type V2EventTuiCommandExecuteCommand string

const (
	V2EventTuiCommandExecuteCommandCommandSessionList         V2EventTuiCommandExecuteCommand = "session.list"
	V2EventTuiCommandExecuteCommandCommandSessionNew          V2EventTuiCommandExecuteCommand = "session.new"
	V2EventTuiCommandExecuteCommandCommandSessionShare        V2EventTuiCommandExecuteCommand = "session.share"
	V2EventTuiCommandExecuteCommandCommandSessionInterrupt    V2EventTuiCommandExecuteCommand = "session.interrupt"
	V2EventTuiCommandExecuteCommandCommandSessionCompact      V2EventTuiCommandExecuteCommand = "session.compact"
	V2EventTuiCommandExecuteCommandCommandSessionPageUp       V2EventTuiCommandExecuteCommand = "session.page.up"
	V2EventTuiCommandExecuteCommandCommandSessionPageDown     V2EventTuiCommandExecuteCommand = "session.page.down"
	V2EventTuiCommandExecuteCommandCommandSessionLineUp       V2EventTuiCommandExecuteCommand = "session.line.up"
	V2EventTuiCommandExecuteCommandCommandSessionLineDown     V2EventTuiCommandExecuteCommand = "session.line.down"
	V2EventTuiCommandExecuteCommandCommandSessionHalfPageUp   V2EventTuiCommandExecuteCommand = "session.half.page.up"
	V2EventTuiCommandExecuteCommandCommandSessionHalfPageDown V2EventTuiCommandExecuteCommand = "session.half.page.down"
	V2EventTuiCommandExecuteCommandCommandSessionFirst        V2EventTuiCommandExecuteCommand = "session.first"
	V2EventTuiCommandExecuteCommandCommandSessionLast         V2EventTuiCommandExecuteCommand = "session.last"
	V2EventTuiCommandExecuteCommandCommandPromptClear         V2EventTuiCommandExecuteCommand = "prompt.clear"
	V2EventTuiCommandExecuteCommandCommandPromptSubmit        V2EventTuiCommandExecuteCommand = "prompt.submit"
	V2EventTuiCommandExecuteCommandCommandAgentCycle          V2EventTuiCommandExecuteCommand = "agent.cycle"
)

func (r V2EventTuiCommandExecuteCommand) IsKnown() bool {
	switch r {
	case V2EventTuiCommandExecuteCommandCommandSessionList,
		V2EventTuiCommandExecuteCommandCommandSessionNew,
		V2EventTuiCommandExecuteCommandCommandSessionShare,
		V2EventTuiCommandExecuteCommandCommandSessionInterrupt,
		V2EventTuiCommandExecuteCommandCommandSessionCompact,
		V2EventTuiCommandExecuteCommandCommandSessionPageUp,
		V2EventTuiCommandExecuteCommandCommandSessionPageDown,
		V2EventTuiCommandExecuteCommandCommandSessionLineUp,
		V2EventTuiCommandExecuteCommandCommandSessionLineDown,
		V2EventTuiCommandExecuteCommandCommandSessionHalfPageUp,
		V2EventTuiCommandExecuteCommandCommandSessionHalfPageDown,
		V2EventTuiCommandExecuteCommandCommandSessionFirst,
		V2EventTuiCommandExecuteCommandCommandSessionLast,
		V2EventTuiCommandExecuteCommandCommandPromptClear,
		V2EventTuiCommandExecuteCommandCommandPromptSubmit,
		V2EventTuiCommandExecuteCommandCommandAgentCycle:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[V2EventPayloadUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventCatalogUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventCommandExecuted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventFileEdited](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventFileWatcherUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventGlobalDisposed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventInstallationUpdateAvailable](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventInstallationUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventIntegrationConnectionUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventIntegrationUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventLspUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventMcpBrowserOpenFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventMcpToolsChanged](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventMessagePartDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventMessagePartRemoved](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventMessagePartUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventMessageRemoved](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventMessageUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventModelsDevRefreshed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventPermissionAsked](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventPermissionReplied](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventPermissionV2Asked](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventPermissionV2Replied](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventPluginAdded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventProjectDirectoriesUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventProjectUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventPtyCreated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventPtyDeleted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventPtyExited](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventPtyUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventQuestionAsked](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventQuestionRejected](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventQuestionReplied](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventQuestionV2Asked](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventQuestionV2Rejected](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventQuestionV2Replied](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventReferenceUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventServerConnected](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionCompacted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionCreated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionDeleted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionDiff](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionIdle](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextAgentSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextCompactionDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextCompactionEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextCompactionStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextContextUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextModelSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextMoved](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextPromptAdmitted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextPrompted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextReasoningDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextReasoningEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextReasoningStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextRetried](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextRevertCleared](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextRevertCommitted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextRevertStaged](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextShellEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextShellStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextStepEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextStepFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextStepStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextSynthetic](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextTextDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextTextEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextTextStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextToolCalled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextToolFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextToolInputDelta](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextToolInputEnded](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextToolInputStarted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextToolProgress](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionNextToolSuccess](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionStatus](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventSessionUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventTodoUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventTuiCommandExecute](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventTuiPromptAppend](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventTuiSessionSelect](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventTuiToastShow](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventVcsBranchUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventWorkspaceFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventWorkspaceReady](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventWorkspaceStatus](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventWorktreeFailed](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2EventWorktreeReady](),
		},
	)
}
