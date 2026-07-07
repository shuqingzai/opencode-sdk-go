// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
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
func (r *V2EventService) Subscribe(ctx context.Context, query EventListParams, opts ...option.RequestOption) (stream *ssestream.Stream[V2Event]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "api/event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[V2Event](ssestream.NewDecoder(raw), err)
}

// V2Event represents a native event payload from the V2 /api/event endpoint.
// The V2 format differs from V1: it uses "data" instead of "properties", and
// includes optional "durable", "location", and "metadata" fields.
type V2Event struct {
	ID       string      `json:"id,required"`
	Type     V2EventType `json:"type,required"`
	Durable  *V2EventDurable `json:"durable"`
	Location *LocationRef   `json:"location"`
	Metadata interface{}     `json:"metadata"`
	Data     interface{}     `json:"data,required"`
	JSON     v2EventJSON     `json:"-"`
}

// V2EventDurable contains the durable sequencing information for a V2Event.
type V2EventDurable struct {
	AggregateID string            `json:"aggregateID,required"`
	Seq         int64             `json:"seq,required"`
	Version     int64             `json:"version,required"`
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
	return apijson.UnmarshalRoot(data, r)
}

// V2EventType enumerates all possible event types for V2Event.
type V2EventType string

const (
	V2EventTypeCatalogUpdated                 V2EventType = "catalog.updated"
	V2EventTypeCommandExecuted                V2EventType = "command.executed"
	V2EventTypeFileEdited                     V2EventType = "file.edited"
	V2EventTypeFileWatcherUpdated             V2EventType = "file.watcher.updated"
	V2EventTypeGlobalDisposed                 V2EventType = "global.disposed"
	V2EventTypeInstallationUpdateAvailable    V2EventType = "installation.update-available"
	V2EventTypeInstallationUpdated            V2EventType = "installation.updated"
	V2EventTypeIntegrationUpdated             V2EventType = "integration.updated"
	V2EventTypeIntegrationConnectionUpdated   V2EventType = "integration.connection.updated"
	V2EventTypeLspUpdated                     V2EventType = "lsp.updated"
	V2EventTypeMcpBrowserOpenFailed           V2EventType = "mcp.browser.open.failed"
	V2EventTypeMcpToolsChanged                V2EventType = "mcp.tools.changed"
	V2EventTypeModelsDevRefreshed             V2EventType = "models-dev.refreshed"
	V2EventTypeMessagePartDelta               V2EventType = "message.part.delta"
	V2EventTypeMessagePartRemoved             V2EventType = "message.part.removed"
	V2EventTypeMessagePartUpdated             V2EventType = "message.part.updated"
	V2EventTypeMessageRemoved                 V2EventType = "message.removed"
	V2EventTypeMessageUpdated                 V2EventType = "message.updated"
	V2EventTypePermissionAsked                V2EventType = "permission.asked"
	V2EventTypePermissionReplied              V2EventType = "permission.replied"
	V2EventTypePermissionV2Asked              V2EventType = "permission.v2.asked"
	V2EventTypePermissionV2Replied            V2EventType = "permission.v2.replied"
	V2EventTypePluginAdded                    V2EventType = "plugin.added"
	V2EventTypeProjectDirectoriesUpdated      V2EventType = "project.directories.updated"
	V2EventTypeProjectUpdated                 V2EventType = "project.updated"
	V2EventTypePtyCreated                     V2EventType = "pty.created"
	V2EventTypePtyDeleted                    V2EventType = "pty.deleted"
	V2EventTypePtyExited                      V2EventType = "pty.exited"
	V2EventTypePtyUpdated                     V2EventType = "pty.updated"
	V2EventTypeQuestionAsked                  V2EventType = "question.asked"
	V2EventTypeQuestionRejected               V2EventType = "question.rejected"
	V2EventTypeQuestionReplied                V2EventType = "question.replied"
	V2EventTypeQuestionV2Asked                V2EventType = "question.v2.asked"
	V2EventTypeQuestionV2Rejected             V2EventType = "question.v2.rejected"
	V2EventTypeQuestionV2Replied              V2EventType = "question.v2.replied"
	V2EventTypeReferenceUpdated               V2EventType = "reference.updated"
	V2EventTypeServerConnected                V2EventType = "server.connected"
	V2EventTypeSessionCompacted               V2EventType = "session.compacted"
	V2EventTypeSessionCreated                 V2EventType = "session.created"
	V2EventTypeSessionDeleted                 V2EventType = "session.deleted"
	V2EventTypeSessionDiff                    V2EventType = "session.diff"
	V2EventTypeSessionError                   V2EventType = "session.error"
	V2EventTypeSessionIdle                    V2EventType = "session.idle"
	V2EventTypeSessionNextAgentSwitched       V2EventType = "session.next.agent.switched"
	V2EventTypeSessionNextCompactionDelta     V2EventType = "session.next.compaction.delta"
	V2EventTypeSessionNextCompactionEnded     V2EventType = "session.next.compaction.ended"
	V2EventTypeSessionNextCompactionStarted   V2EventType = "session.next.compaction.started"
	V2EventTypeSessionNextContextUpdated      V2EventType = "session.next.context.updated"
	V2EventTypeSessionNextModelSwitched       V2EventType = "session.next.model.switched"
	V2EventTypeSessionNextMoved               V2EventType = "session.next.moved"
	V2EventTypeSessionNextPromptAdmitted      V2EventType = "session.next.prompt.admitted"
	V2EventTypeSessionNextPrompted            V2EventType = "session.next.prompted"
	V2EventTypeSessionNextReasoningDelta      V2EventType = "session.next.reasoning.delta"
	V2EventTypeSessionNextReasoningEnded      V2EventType = "session.next.reasoning.ended"
	V2EventTypeSessionNextReasoningStarted    V2EventType = "session.next.reasoning.started"
	V2EventTypeSessionNextRetried             V2EventType = "session.next.retried"
	V2EventTypeSessionNextRevertCleared       V2EventType = "session.next.revert.cleared"
	V2EventTypeSessionNextRevertCommitted     V2EventType = "session.next.revert.committed"
	V2EventTypeSessionNextRevertStaged        V2EventType = "session.next.revert.staged"
	V2EventTypeSessionNextShellEnded          V2EventType = "session.next.shell.ended"
	V2EventTypeSessionNextShellStarted        V2EventType = "session.next.shell.started"
	V2EventTypeSessionNextStepEnded           V2EventType = "session.next.step.ended"
	V2EventTypeSessionNextStepFailed          V2EventType = "session.next.step.failed"
	V2EventTypeSessionNextStepStarted         V2EventType = "session.next.step.started"
	V2EventTypeSessionNextSynthetic           V2EventType = "session.next.synthetic"
	V2EventTypeSessionNextTextDelta           V2EventType = "session.next.text.delta"
	V2EventTypeSessionNextTextEnded           V2EventType = "session.next.text.ended"
	V2EventTypeSessionNextTextStarted         V2EventType = "session.next.text.started"
	V2EventTypeSessionNextToolCalled          V2EventType = "session.next.tool.called"
	V2EventTypeSessionNextToolFailed          V2EventType = "session.next.tool.failed"
	V2EventTypeSessionNextToolInputDelta      V2EventType = "session.next.tool.input.delta"
	V2EventTypeSessionNextToolInputEnded      V2EventType = "session.next.tool.input.ended"
	V2EventTypeSessionNextToolInputStarted    V2EventType = "session.next.tool.input.started"
	V2EventTypeSessionNextToolProgress        V2EventType = "session.next.tool.progress"
	V2EventTypeSessionNextToolSuccess         V2EventType = "session.next.tool.success"
	V2EventTypeSessionStatus                  V2EventType = "session.status"
	V2EventTypeSessionUpdated                 V2EventType = "session.updated"
	V2EventTypeTodoUpdated                    V2EventType = "todo.updated"
	V2EventTypeTuiCommandExecute              V2EventType = "tui.command.execute"
	V2EventTypeTuiPromptAppend                V2EventType = "tui.prompt.append"
	V2EventTypeTuiSessionSelect               V2EventType = "tui.session.select"
	V2EventTypeTuiToastShow                   V2EventType = "tui.toast.show"
	V2EventTypeVcsBranchUpdated               V2EventType = "vcs.branch.updated"
	V2EventTypeWorkspaceFailed                V2EventType = "workspace.failed"
	V2EventTypeWorkspaceReady                 V2EventType = "workspace.ready"
	V2EventTypeWorkspaceStatus                V2EventType = "workspace.status"
	V2EventTypeWorktreeFailed                 V2EventType = "worktree.failed"
	V2EventTypeWorktreeReady                  V2EventType = "worktree.ready"
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
		V2EventTypeIntegrationUpdated,
		V2EventTypeIntegrationConnectionUpdated,
		V2EventTypeLspUpdated,
		V2EventTypeMcpBrowserOpenFailed,
		V2EventTypeMcpToolsChanged,
		V2EventTypeModelsDevRefreshed,
		V2EventTypeMessagePartDelta,
		V2EventTypeMessagePartRemoved,
		V2EventTypeMessagePartUpdated,
		V2EventTypeMessageRemoved,
		V2EventTypeMessageUpdated,
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
