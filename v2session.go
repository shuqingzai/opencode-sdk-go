// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/pagination"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
	"github.com/tidwall/gjson"
)

// V2SessionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2SessionService] method instead.
type V2SessionService struct {
	Options    []option.RequestOption
	Revert     *V2SessionRevertService
	Permission *V2SessionPermissionService
	Question   *V2SessionQuestionService
}

// NewV2SessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2SessionService(opts ...option.RequestOption) (r *V2SessionService) {
	r = &V2SessionService{}
	r.Options = opts
	r.Revert = NewV2SessionRevertService(opts...)
	r.Permission = NewV2SessionPermissionService(opts...)
	r.Question = NewV2SessionQuestionService(opts...)
	return
}

// List v2 sessions
//
// Retrieve sessions in the requested order. Items keep that order across pages;
// use cursor.next or cursor.previous to move through the ordered list.
func (r *V2SessionService) List(ctx context.Context, query V2SessionListParams, opts ...option.RequestOption) (res *V2SessionsResponse, err error) {
	var raw *http.Response
	opts = slices.Concat([]option.RequestOption{option.WithResponseInto(&raw)}, r.Options, opts)
	path := "api/session"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List v2 sessions
//
// Retrieve sessions in the requested order. Items keep that order across pages;
// use cursor.next or cursor.previous to move through the ordered list.
//
// ListAutoPaging walks every page for you, transparently replaying cursor.next
// until the server runs out of sessions.
func (r *V2SessionService) ListAutoPaging(ctx context.Context, query V2SessionListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V2SessionInfo] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Send v2 message
//
// Create a v2 session message and queue it for the agent loop.
func (r *V2SessionService) Prompt(ctx context.Context, sessionID string, params V2SessionPromptParams, opts ...option.RequestOption) (res *V2SessionPromptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/prompt", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Compact v2 session
//
// Compact a v2 session conversation.
func (r *V2SessionService) Compact(ctx context.Context, sessionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/compact", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Wait for v2 session
//
// Wait for a v2 session agent loop to become idle.
func (r *V2SessionService) Wait(ctx context.Context, sessionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/wait", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Get v2 session context
//
// Retrieve the active context messages for a v2 session (all messages after the
// last compaction).
func (r *V2SessionService) Context(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *V2SessionContextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/context", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get v2 session messages
//
// Retrieve projected v2 messages for a session. Items keep the requested order
// across pages; use cursor.next or cursor.previous to move through the ordered
// timeline.
func (r *V2SessionService) Messages(ctx context.Context, sessionID string, query V2SessionMessagesParams, opts ...option.RequestOption) (res *V2SessionMessagesResponse, err error) {
	var raw *http.Response
	opts = slices.Concat([]option.RequestOption{option.WithResponseInto(&raw)}, r.Options, opts)
	if sessionID == "" {
		return nil, errors.New("missing required sessionID parameter")
	}
	path := fmt.Sprintf("api/session/%s/message", sessionID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get v2 session messages
//
// Retrieve projected v2 messages for a session. Items keep the requested order
// across pages; use cursor.next or cursor.previous to move through the ordered
// timeline.
//
// MessagesAutoPaging walks every page for you, transparently replaying
// cursor.next until the session timeline is exhausted. The order requested on
// the first call is preserved across pages because it is baked into the cursor.
func (r *V2SessionService) MessagesAutoPaging(ctx context.Context, sessionID string, query V2SessionMessagesParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V2SessionMessage] {
	return pagination.NewCursorPageAutoPager(r.Messages(ctx, sessionID, query, opts...))
}

// New v2 session
//
// Create a session at the requested location.
func (r *V2SessionService) New(ctx context.Context, params V2SessionNewParams, opts ...option.RequestOption) (res *V2SessionCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get active sessions
//
// Retrieve foreground Session drains currently owned by this OpenCode process.
// Sessions absent from the result are inactive.
func (r *V2SessionService) Active(ctx context.Context, opts ...option.RequestOption) (res *V2SessionActiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/session/active"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get v2 session
//
// Retrieve a session by ID.
func (r *V2SessionService) Get(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *V2SessionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Switch session agent
//
// Switch the agent used by subsequent provider turns.
func (r *V2SessionService) SwitchAgent(ctx context.Context, sessionID string, params V2SessionSwitchAgentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/agent", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Switch session model
//
// Switch the model used by subsequent provider turns.
func (r *V2SessionService) SwitchModel(ctx context.Context, sessionID string, params V2SessionSwitchModelParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/model", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Subscribe to session events
//
// Replay durable events after an aggregate sequence, then continue with new
// durable events.
func (r *V2SessionService) Events(ctx context.Context, sessionID string, query V2SessionEventsParams, opts ...option.RequestOption) (stream *ssestream.Stream[V2SessionEventResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if sessionID == "" {
		return ssestream.NewStream[V2SessionEventResponse](ssestream.NewDecoder(nil), errors.New("missing required sessionID parameter"))
	}
	path := fmt.Sprintf("api/session/%s/event", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[V2SessionEventResponse](ssestream.NewDecoder(raw), err)
}

// Get session history
//
// Read one finite page of public durable Session events after an exclusive
// aggregate sequence. Newly committed events may appear on later pages.
func (r *V2SessionService) History(ctx context.Context, sessionID string, query V2SessionHistoryParams, opts ...option.RequestOption) (res *V2SessionHistoryResponse, err error) {
	var raw *http.Response
	opts = slices.Concat([]option.RequestOption{option.WithResponseInto(&raw)}, r.Options, opts)
	if sessionID == "" {
		return nil, errors.New("missing required sessionID parameter")
	}
	path := fmt.Sprintf("api/session/%s/history", sessionID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get session history
//
// Read one finite page of public durable Session events after an exclusive
// aggregate sequence. Newly committed events may appear on later pages.
//
// HistoryAutoPaging walks every page for you, advancing the `after` parameter to
// the durable sequence number of the last event it saw until hasMore turns
// false.
func (r *V2SessionService) HistoryAutoPaging(ctx context.Context, sessionID string, query V2SessionHistoryParams, opts ...option.RequestOption) *pagination.SeqPageAutoPager[V2SessionDurableEvent] {
	return pagination.NewSeqPageAutoPager(r.History(ctx, sessionID, query, opts...))
}

// Interrupt session execution
//
// Interrupt active execution owned by this OpenCode process. Idle interruption is
// a no-op.
func (r *V2SessionService) Interrupt(ctx context.Context, sessionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/interrupt", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Get session message
//
// Retrieve one projected message owned by the Session.
func (r *V2SessionService) Message(ctx context.Context, sessionID string, messageID string, opts ...option.RequestOption) (res *V2SessionMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/message/%s", sessionID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// ===== Response Types =====

// V2SessionsResponse is returned by the List method. It is a page of sessions:
// use its GetNextPage / GetPreviousPage methods, or [V2SessionService.ListAutoPaging],
// to walk the ordered list through Cursor.Next and Cursor.Previous.
type V2SessionsResponse = pagination.CursorPage[V2SessionInfo]

type V2SessionInfo struct {
	ID        string              `json:"id,required"`
	ProjectID string              `json:"projectID,required"`
	Cost      float64             `json:"cost,required"`
	Time      V2SessionInfoTime   `json:"time,required"`
	Title     string              `json:"title,required"`
	Tokens    V2SessionInfoTokens `json:"tokens,required"`
	ParentID  string              `json:"parentID"`
	Location  LocationRef         `json:"location,required"`
	Subpath   string              `json:"subpath"`
	Agent     string              `json:"agent"`
	Model     ModelRef            `json:"model"`
	Revert    RevertState         `json:"revert"`
	JSON      v2SessionInfoJSON   `json:"-"`
}

type v2SessionInfoJSON struct {
	ID          apijson.Field
	ProjectID   apijson.Field
	Cost        apijson.Field
	Time        apijson.Field
	Title       apijson.Field
	Tokens      apijson.Field
	ParentID    apijson.Field
	Location    apijson.Field
	Subpath     apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
	Revert      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionInfoJSON) RawJSON() string {
	return r.raw
}

type V2SessionInfoTime struct {
	Created  int64                 `json:"created,required"`
	Updated  int64                 `json:"updated,required"`
	Archived int64                 `json:"archived"`
	JSON     v2SessionInfoTimeJSON `json:"-"`
}

type v2SessionInfoTimeJSON struct {
	Created     apijson.Field
	Updated     apijson.Field
	Archived    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionInfoTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionInfoTimeJSON) RawJSON() string {
	return r.raw
}

type V2SessionInfoTokens struct {
	Input     int64                    `json:"input,required"`
	Output    int64                    `json:"output,required"`
	Reasoning int64                    `json:"reasoning,required"`
	Cache     V2SessionInfoTokensCache `json:"cache,required"`
	JSON      v2SessionInfoTokensJSON  `json:"-"`
}

type v2SessionInfoTokensJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionInfoTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionInfoTokensJSON) RawJSON() string {
	return r.raw
}

type V2SessionInfoTokensCache struct {
	Read  int64                        `json:"read,required"`
	Write int64                        `json:"write,required"`
	JSON  v2SessionInfoTokensCacheJSON `json:"-"`
}

type v2SessionInfoTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionInfoTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionInfoTokensCacheJSON) RawJSON() string {
	return r.raw
}

// V2Cursor carries the opaque pagination cursors returned by the v2 session
// endpoints. Both values already encode the ordering of the sequence they walk,
// so they must be replayed as-is and must not be combined with the `order` query
// parameter.
type V2Cursor = pagination.PageCursor

// V2SessionMessagesResponse is returned by the Messages method. It is a page of
// messages: use its GetNextPage / GetPreviousPage methods, or
// [V2SessionService.MessagesAutoPaging], to walk the ordered timeline through
// Cursor.Next and Cursor.Previous.
type V2SessionMessagesResponse = pagination.CursorPage[V2SessionMessage]

// V2SessionContextResponse is returned by the Context method. It wraps messages
// in a data field.
type V2SessionContextResponse struct {
	Data []V2SessionMessage           `json:"data,required"`
	JSON v2SessionContextResponseJSON `json:"-"`
}

type v2SessionContextResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionContextResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionContextResponseJSON) RawJSON() string {
	return r.raw
}

// V2SessionMessage represents the discriminated union of v2 session message
// types. Possible runtime types of the union are
// [V2SessionMessageAgentSwitched], [V2SessionMessageModelSwitched],
// [V2SessionMessageUser], [V2SessionMessageSynthetic], [V2SessionMessageSystem],
// [V2SessionMessageShell], [V2SessionMessageAssistant],
// [V2SessionMessageCompaction].
type V2SessionMessage struct {
	JSON  v2SessionMessageJSON `json:"-"`
	union V2SessionMessageUnion
}

// v2SessionMessageJSON contains the JSON metadata for the struct
// [V2SessionMessage]
type v2SessionMessageJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2SessionMessageJSON) RawJSON() string {
	return r.raw
}

func (r *V2SessionMessage) UnmarshalJSON(data []byte) (err error) {
	*r = V2SessionMessage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2SessionMessageUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [V2SessionMessageAgentSwitched],
// [V2SessionMessageModelSwitched], [V2SessionMessageUser],
// [V2SessionMessageSynthetic], [V2SessionMessageSystem],
// [V2SessionMessageShell], [V2SessionMessageAssistant],
// [V2SessionMessageCompaction].
func (r V2SessionMessage) AsUnion() V2SessionMessageUnion {
	return r.union
}

// V2SessionMessageResponse is returned by the Message method. The OpenAPI
// response wraps the projected [V2SessionMessage] in a `data` envelope.
type V2SessionMessageResponse struct {
	Data V2SessionMessage             `json:"data,required"`
	JSON v2SessionMessageResponseJSON `json:"-"`
}

type v2SessionMessageResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [V2SessionMessageAgentSwitched],
// [V2SessionMessageModelSwitched], [V2SessionMessageUser],
// [V2SessionMessageSynthetic], [V2SessionMessageSystem],
// [V2SessionMessageShell], [V2SessionMessageAssistant] or
// [V2SessionMessageCompaction].
type V2SessionMessageUnion interface {
	implementsV2SessionMessageUnion()
}

func (V2SessionMessageAgentSwitched) implementsV2SessionMessageUnion() {}
func (V2SessionMessageModelSwitched) implementsV2SessionMessageUnion() {}
func (V2SessionMessageUser) implementsV2SessionMessageUnion()          {}
func (V2SessionMessageSynthetic) implementsV2SessionMessageUnion()     {}
func (V2SessionMessageSystem) implementsV2SessionMessageUnion()        {}
func (V2SessionMessageShell) implementsV2SessionMessageUnion()         {}
func (V2SessionMessageAssistant) implementsV2SessionMessageUnion()     {}
func (V2SessionMessageCompaction) implementsV2SessionMessageUnion()    {}

type V2SessionMessageAgentSwitched struct {
	ID    string                            `json:"id,required"`
	Time  V2SessionMessageTime              `json:"time,required"`
	Type  V2SessionMessageAgentSwitchedType `json:"type,required"`
	Agent string                            `json:"agent,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                               `json:"metadata"`
	JSON     v2SessionMessageAgentSwitchedJSON `json:"-"`
}

type v2SessionMessageAgentSwitchedJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	Agent       apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageAgentSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAgentSwitchedJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageAgentSwitchedType string

const (
	V2SessionMessageAgentSwitchedTypeAgentSwitched V2SessionMessageAgentSwitchedType = "agent-switched"
)

func (r V2SessionMessageAgentSwitchedType) IsKnown() bool {
	switch r {
	case V2SessionMessageAgentSwitchedTypeAgentSwitched:
		return true
	}
	return false
}

type V2SessionMessageModelSwitched struct {
	ID    string                            `json:"id,required"`
	Time  V2SessionMessageTime              `json:"time,required"`
	Type  V2SessionMessageModelSwitchedType `json:"type,required"`
	Model ModelRef                          `json:"model,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                               `json:"metadata"`
	JSON     v2SessionMessageModelSwitchedJSON `json:"-"`
}

type v2SessionMessageModelSwitchedJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	Model       apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageModelSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageModelSwitchedJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageModelSwitchedType string

const (
	V2SessionMessageModelSwitchedTypeModelSwitched V2SessionMessageModelSwitchedType = "model-switched"
)

func (r V2SessionMessageModelSwitchedType) IsKnown() bool {
	switch r {
	case V2SessionMessageModelSwitchedTypeModelSwitched:
		return true
	}
	return false
}

type V2SessionMessageUser struct {
	ID     string                    `json:"id,required"`
	Time   V2SessionMessageTime      `json:"time,required"`
	Text   string                    `json:"text,required"`
	Type   V2SessionMessageUserType  `json:"type,required"`
	Files  []V2PromptFileAttachment  `json:"files"`
	Agents []V2PromptAgentAttachment `json:"agents"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                      `json:"metadata"`
	JSON     v2SessionMessageUserJSON `json:"-"`
}

type v2SessionMessageUserJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	Files       apijson.Field
	Agents      apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageUserJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageUserType string

const (
	V2SessionMessageUserTypeUser V2SessionMessageUserType = "user"
)

func (r V2SessionMessageUserType) IsKnown() bool {
	switch r {
	case V2SessionMessageUserTypeUser:
		return true
	}
	return false
}

type V2SessionMessageSynthetic struct {
	ID        string                        `json:"id,required"`
	Time      V2SessionMessageTime          `json:"time,required"`
	SessionID string                        `json:"sessionID,required"`
	Text      string                        `json:"text,required"`
	Type      V2SessionMessageSyntheticType `json:"type,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                           `json:"metadata"`
	JSON     v2SessionMessageSyntheticJSON `json:"-"`
}

type v2SessionMessageSyntheticJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageSynthetic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageSyntheticJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageSyntheticType string

const (
	V2SessionMessageSyntheticTypeSynthetic V2SessionMessageSyntheticType = "synthetic"
)

func (r V2SessionMessageSyntheticType) IsKnown() bool {
	switch r {
	case V2SessionMessageSyntheticTypeSynthetic:
		return true
	}
	return false
}

type V2SessionMessageShell struct {
	ID      string                    `json:"id,required"`
	Time    V2SessionMessageShellTime `json:"time,required"`
	Type    V2SessionMessageShellType `json:"type,required"`
	CallID  string                    `json:"callID,required"`
	Command string                    `json:"command,required"`
	Output  string                    `json:"output,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                       `json:"metadata"`
	JSON     v2SessionMessageShellJSON `json:"-"`
}

type v2SessionMessageShellJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	CallID      apijson.Field
	Command     apijson.Field
	Output      apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageShell) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageShellJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageShellType string

const (
	V2SessionMessageShellTypeShell V2SessionMessageShellType = "shell"
)

func (r V2SessionMessageShellType) IsKnown() bool {
	switch r {
	case V2SessionMessageShellTypeShell:
		return true
	}
	return false
}

type V2SessionMessageAssistant struct {
	ID    string                        `json:"id,required"`
	Time  V2SessionMessageAssistantTime `json:"time,required"`
	Type  V2SessionMessageAssistantType `json:"type,required"`
	Agent string                        `json:"agent,required"`
	Model ModelRef                      `json:"model,required"`
	// This field can have the runtime type of
	// [V2SessionMessageAssistantTextContent],
	// [V2SessionMessageAssistantReasoningContent],
	// [V2SessionMessageAssistantToolContent].
	Content  []V2SessionMessageAssistantContent `json:"content,required"`
	Snapshot V2SessionMessageAssistantSnapshot  `json:"snapshot"`
	Finish   string                             `json:"finish"`
	Cost     float64                            `json:"cost"`
	Tokens   V2SessionMessageTokens             `json:"tokens"`
	Error    SessionErrorUnknown                `json:"error"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                           `json:"metadata"`
	JSON     v2SessionMessageAssistantJSON `json:"-"`
}

type v2SessionMessageAssistantJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
	Content     apijson.Field
	Snapshot    apijson.Field
	Finish      apijson.Field
	Cost        apijson.Field
	Tokens      apijson.Field
	Error       apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageAssistant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAssistantJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageAssistantType string

const (
	V2SessionMessageAssistantTypeAssistant V2SessionMessageAssistantType = "assistant"
)

func (r V2SessionMessageAssistantType) IsKnown() bool {
	switch r {
	case V2SessionMessageAssistantTypeAssistant:
		return true
	}
	return false
}

type V2SessionMessageCompaction struct {
	ID      string                           `json:"id,required"`
	Time    V2SessionMessageTime             `json:"time,required"`
	Type    V2SessionMessageCompactionType   `json:"type,required"`
	Reason  V2SessionMessageCompactionReason `json:"reason,required"`
	Summary string                           `json:"summary,required"`
	Recent  string                           `json:"recent,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                            `json:"metadata"`
	JSON     v2SessionMessageCompactionJSON `json:"-"`
}

type v2SessionMessageCompactionJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	Reason      apijson.Field
	Summary     apijson.Field
	Recent      apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageCompactionJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageCompactionType string

const (
	V2SessionMessageCompactionTypeCompaction V2SessionMessageCompactionType = "compaction"
)

func (r V2SessionMessageCompactionType) IsKnown() bool {
	switch r {
	case V2SessionMessageCompactionTypeCompaction:
		return true
	}
	return false
}

type V2SessionMessageCompactionReason string

const (
	V2SessionMessageCompactionReasonAuto   V2SessionMessageCompactionReason = "auto"
	V2SessionMessageCompactionReasonManual V2SessionMessageCompactionReason = "manual"
)

func (r V2SessionMessageCompactionReason) IsKnown() bool {
	switch r {
	case V2SessionMessageCompactionReasonAuto, V2SessionMessageCompactionReasonManual:
		return true
	}
	return false
}

type SessionErrorUnknown struct {
	Type    SessionErrorUnknownType `json:"type,required"`
	Message string                  `json:"message,required"`
	JSON    sessionErrorUnknownJSON `json:"-"`
}

// sessionErrorUnknownJSON contains the JSON metadata for the struct [SessionErrorUnknown]
type sessionErrorUnknownJSON struct {
	Type        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionErrorUnknown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionErrorUnknownJSON) RawJSON() string {
	return r.raw
}

type SessionErrorUnknownType string

const (
	SessionErrorUnknownTypeUnknown SessionErrorUnknownType = "unknown"
)

func (r SessionErrorUnknownType) IsKnown() bool {
	switch r {
	case SessionErrorUnknownTypeUnknown:
		return true
	}
	return false
}

// Shared sub-types

type V2SessionMessageTime struct {
	Created int64                    `json:"created,required"`
	JSON    v2SessionMessageTimeJSON `json:"-"`
}

type v2SessionMessageTimeJSON struct {
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageTimeJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageShellTime struct {
	Created   int64                         `json:"created,required"`
	Completed int64                         `json:"completed"`
	JSON      v2SessionMessageShellTimeJSON `json:"-"`
}

type v2SessionMessageShellTimeJSON struct {
	Created     apijson.Field
	Completed   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageShellTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageShellTimeJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageAssistantTime struct {
	Created   int64                             `json:"created,required"`
	Completed int64                             `json:"completed"`
	JSON      v2SessionMessageAssistantTimeJSON `json:"-"`
}

type v2SessionMessageAssistantTimeJSON struct {
	Created     apijson.Field
	Completed   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageAssistantTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAssistantTimeJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageTokens struct {
	Input     int64                       `json:"input,required"`
	Output    int64                       `json:"output,required"`
	Reasoning int64                       `json:"reasoning,required"`
	Cache     V2SessionMessageTokensCache `json:"cache,required"`
	JSON      v2SessionMessageTokensJSON  `json:"-"`
}

type v2SessionMessageTokensJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageTokensJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageTokensCache struct {
	Read  int64                           `json:"read,required"`
	Write int64                           `json:"write,required"`
	JSON  v2SessionMessageTokensCacheJSON `json:"-"`
}

type v2SessionMessageTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageTokensCacheJSON) RawJSON() string {
	return r.raw
}

// V2SessionMessageAssistantContent represents a discriminated union of v2
// assistant content types. Possible runtime types of the union are
// [V2SessionMessageAssistantTextContent], [V2SessionMessageAssistantReasoningContent]
// or [V2SessionMessageAssistantToolContent].
type V2SessionMessageAssistantContent struct {
	JSON  v2SessionMessageAssistantContentJSON `json:"-"`
	union V2SessionMessageAssistantContentUnion
}

// v2SessionMessageAssistantContentJSON contains the JSON metadata for the
// struct [V2SessionMessageAssistantContent]
type v2SessionMessageAssistantContentJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2SessionMessageAssistantContentJSON) RawJSON() string {
	return r.raw
}

func (r *V2SessionMessageAssistantContent) UnmarshalJSON(data []byte) (err error) {
	*r = V2SessionMessageAssistantContent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2SessionMessageAssistantContentUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2SessionMessageAssistantTextContent],
// [V2SessionMessageAssistantReasoningContent] or
// [V2SessionMessageAssistantToolContent].
func (r V2SessionMessageAssistantContent) AsUnion() V2SessionMessageAssistantContentUnion {
	return r.union
}

// V2SessionMessageAssistantContentUnion is satisfied by
// [V2SessionMessageAssistantTextContent],
// [V2SessionMessageAssistantReasoningContent] or
// [V2SessionMessageAssistantToolContent].
type V2SessionMessageAssistantContentUnion interface {
	implementsV2SessionMessageAssistantContent()
}

func (V2SessionMessageAssistantTextContent) implementsV2SessionMessageAssistantContent()      {}
func (V2SessionMessageAssistantReasoningContent) implementsV2SessionMessageAssistantContent() {}
func (V2SessionMessageAssistantToolContent) implementsV2SessionMessageAssistantContent()      {}

type V2SessionMessageAssistantSnapshot struct {
	Start string                                `json:"start"`
	End   string                                `json:"end"`
	Files []string                              `json:"files"`
	JSON  v2SessionMessageAssistantSnapshotJSON `json:"-"`
}

type v2SessionMessageAssistantSnapshotJSON struct {
	Start       apijson.Field
	End         apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageAssistantSnapshot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAssistantSnapshotJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageToolProvider struct {
	Executed bool `json:"executed,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any `json:"metadata"`
	// This field can have the runtime type of [map[string]any].
	ResultMetadata any                              `json:"resultMetadata"`
	JSON           v2SessionMessageToolProviderJSON `json:"-"`
}

type v2SessionMessageToolProviderJSON struct {
	Executed       apijson.Field
	Metadata       apijson.Field
	ResultMetadata apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *V2SessionMessageToolProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageToolProviderJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageToolTime struct {
	Created   int64                        `json:"created,required"`
	Ran       int64                        `json:"ran"`
	Completed int64                        `json:"completed"`
	Pruned    int64                        `json:"pruned"`
	JSON      v2SessionMessageToolTimeJSON `json:"-"`
}

type v2SessionMessageToolTimeJSON struct {
	Created     apijson.Field
	Ran         apijson.Field
	Completed   apijson.Field
	Pruned      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageToolTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageToolTimeJSON) RawJSON() string {
	return r.raw
}

// V2SessionMessage assistant content sub-types are discriminated by the "type"
// field: "text" (V2SessionMessageAssistantTextContent), "reasoning"
// (V2SessionMessageAssistantReasoningContent), and "tool"
// (V2SessionMessageAssistantToolContent). They are exposed through the
// [V2SessionMessageAssistantContent] union, accessible via [V2SessionMessageAssistantContent.AsUnion].

// V2SessionInputAdmitted is returned by the Prompt method. It represents the
// server's admission of a prompt input.
type V2SessionInputAdmitted struct {
	AdmittedSeq int64                      `json:"admittedSeq,required"`
	ID          string                     `json:"id,required"`
	SessionID   string                     `json:"sessionID,required"`
	Prompt      V2SessionInputPrompt       `json:"prompt,required"`
	Delivery    SessionDelivery            `json:"delivery,required"`
	TimeCreated int64                      `json:"timeCreated,required"`
	PromotedSeq int64                      `json:"promotedSeq"`
	JSON        v2SessionInputAdmittedJSON `json:"-"`
}

type v2SessionInputAdmittedJSON struct {
	AdmittedSeq apijson.Field
	ID          apijson.Field
	SessionID   apijson.Field
	Prompt      apijson.Field
	Delivery    apijson.Field
	TimeCreated apijson.Field
	PromotedSeq apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionInputAdmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionInputAdmittedJSON) RawJSON() string {
	return r.raw
}

// V2SessionPromptResponse wraps the SessionInputAdmitted returned by the v2
// Prompt endpoint. The OpenAPI response is {data: SessionInputAdmitted}.
type V2SessionPromptResponse struct {
	Data V2SessionInputAdmitted      `json:"data,required"`
	JSON v2SessionPromptResponseJSON `json:"-"`
}

type v2SessionPromptResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionPromptResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionPromptResponseJSON) RawJSON() string {
	return r.raw
}

type V2SessionInputPrompt struct {
	Text   string                    `json:"text,required"`
	Files  []V2PromptFileAttachment  `json:"files"`
	Agents []V2PromptAgentAttachment `json:"agents"`
	JSON   v2SessionInputPromptJSON  `json:"-"`
}

type v2SessionInputPromptJSON struct {
	Text        apijson.Field
	Files       apijson.Field
	Agents      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionInputPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionInputPromptJSON) RawJSON() string {
	return r.raw
}

// ===== Prompt Attachment Types =====

type V2PromptFileAttachment struct {
	URI         string                     `json:"uri,required"`
	Mime        string                     `json:"mime,required"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Source      V2PromptSource             `json:"source"`
	JSON        v2PromptFileAttachmentJSON `json:"-"`
}

type v2PromptFileAttachmentJSON struct {
	URI         apijson.Field
	Mime        apijson.Field
	Name        apijson.Field
	Description apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PromptFileAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PromptFileAttachmentJSON) RawJSON() string {
	return r.raw
}

type V2PromptAgentAttachment struct {
	Name   string                      `json:"name,required"`
	Source V2PromptSource              `json:"source"`
	JSON   v2PromptAgentAttachmentJSON `json:"-"`
}

type v2PromptAgentAttachmentJSON struct {
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PromptAgentAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PromptAgentAttachmentJSON) RawJSON() string {
	return r.raw
}

type V2PromptSource struct {
	Start int64              `json:"start,required"`
	End   int64              `json:"end,required"`
	Text  string             `json:"text,required"`
	JSON  v2PromptSourceJSON `json:"-"`
}

type v2PromptSourceJSON struct {
	Start       apijson.Field
	End         apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PromptSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PromptSourceJSON) RawJSON() string {
	return r.raw
}

// ===== V2 Content Sub-Types =====

type V2SessionMessageAssistantTextContent struct {
	Type V2SessionMessageAssistantTextContentType `json:"type,required"`
	ID   string                                   `json:"id,required"`
	Text string                                   `json:"text,required"`
	JSON v2SessionMessageAssistantTextContentJSON `json:"-"`
}

type v2SessionMessageAssistantTextContentJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageAssistantTextContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAssistantTextContentJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageAssistantTextContentType string

const (
	V2SessionMessageAssistantTextContentTypeText V2SessionMessageAssistantTextContentType = "text"
)

func (r V2SessionMessageAssistantTextContentType) IsKnown() bool {
	switch r {
	case V2SessionMessageAssistantTextContentTypeText:
		return true
	}
	return false
}

type V2SessionMessageAssistantReasoningContent struct {
	Type             V2SessionMessageAssistantReasoningContentType `json:"type,required"`
	ID               string                                        `json:"id,required"`
	Text             string                                        `json:"text,required"`
	ProviderMetadata map[string]any                                `json:"providerMetadata"`
	Time             V2SessionMessageAssistantReasoningContentTime `json:"time"`
	JSON             v2SessionMessageAssistantReasoningContentJSON `json:"-"`
}

type v2SessionMessageAssistantReasoningContentJSON struct {
	Type             apijson.Field
	ID               apijson.Field
	Text             apijson.Field
	ProviderMetadata apijson.Field
	Time             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V2SessionMessageAssistantReasoningContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAssistantReasoningContentJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageAssistantReasoningContentType string

const (
	V2SessionMessageAssistantReasoningContentTypeReasoning V2SessionMessageAssistantReasoningContentType = "reasoning"
)

func (r V2SessionMessageAssistantReasoningContentType) IsKnown() bool {
	switch r {
	case V2SessionMessageAssistantReasoningContentTypeReasoning:
		return true
	}
	return false
}

type V2SessionMessageAssistantToolContent struct {
	Type     V2SessionMessageAssistantToolContentType `json:"type,required"`
	ID       string                                   `json:"id,required"`
	Name     string                                   `json:"name,required"`
	Provider V2SessionMessageToolProvider             `json:"provider"`
	// This field can have the runtime type of
	// [V2SessionMessageToolStatePending], [V2SessionMessageToolStateRunning],
	// [V2SessionMessageToolStateCompleted], [V2SessionMessageToolStateError].
	State V2SessionMessageToolState                `json:"state,required"`
	Time  V2SessionMessageToolTime                 `json:"time,required"`
	JSON  v2SessionMessageAssistantToolContentJSON `json:"-"`
}

type v2SessionMessageAssistantToolContentJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	Provider    apijson.Field
	State       apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageAssistantToolContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAssistantToolContentJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageAssistantToolContentType string

const (
	V2SessionMessageAssistantToolContentTypeTool V2SessionMessageAssistantToolContentType = "tool"
)

func (r V2SessionMessageAssistantToolContentType) IsKnown() bool {
	switch r {
	case V2SessionMessageAssistantToolContentTypeTool:
		return true
	}
	return false
}

// ===== Tool Content Types (shared) =====

type ToolTextContent struct {
	Type ToolTextContentType `json:"type,required"`
	Text string              `json:"text,required"`
	JSON toolTextContentJSON `json:"-"`
}

type toolTextContentJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolTextContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolTextContentJSON) RawJSON() string {
	return r.raw
}

// ToolTextContentType is the single-value discriminator for [ToolTextContent].
//
// OpenAPI: ToolTextContent.type (enum ["text"]); JS SDK: ToolTextContent.type
// (literal "text"). This is intentionally its own single-value enum rather
// than a reuse of [LLMToolContentType] (which additionally allows "file" for
// the merged [LLMToolContent] carrier) — matching the established pattern of
// [AgentPartType]/[TextPartType]/[FilePartType] in session.go, where each
// individual union variant gets its own single-value enum distinct from the
// merged carrier's wider enum (e.g. [PartType]).
type ToolTextContentType string

const (
	ToolTextContentTypeText ToolTextContentType = "text"
)

func (r ToolTextContentType) IsKnown() bool {
	switch r {
	case ToolTextContentTypeText:
		return true
	}
	return false
}

type ToolFileContent struct {
	Type ToolFileContentType `json:"type,required"`
	URI  string              `json:"uri,required"`
	Mime string              `json:"mime,required"`
	Name string              `json:"name"`
	JSON toolFileContentJSON `json:"-"`
}

type toolFileContentJSON struct {
	Type        apijson.Field
	URI         apijson.Field
	Mime        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolFileContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolFileContentJSON) RawJSON() string {
	return r.raw
}

// ToolFileContentType is the single-value discriminator for [ToolFileContent].
//
// OpenAPI: ToolFileContent.type (enum ["file"]); JS SDK: ToolFileContent.type
// (literal "file"). See [ToolTextContentType] for why this is a dedicated
// single-value enum rather than a reuse of [LLMToolContentType].
type ToolFileContentType string

const (
	ToolFileContentTypeFile ToolFileContentType = "file"
)

func (r ToolFileContentType) IsKnown() bool {
	switch r {
	case ToolFileContentTypeFile:
		return true
	}
	return false
}

func (r ToolTextContent) implementsLLMToolContentUnion() {}

func (r ToolFileContent) implementsLLMToolContentUnion() {}

// LLMToolContentUnion is satisfied by [ToolTextContent] or [ToolFileContent].
// It corresponds to the OpenAPI schema LLMToolContent (anyOf ToolTextContent,
// ToolFileContent) and the JS SDK type LlmToolContent.
type LLMToolContentUnion interface {
	implementsLLMToolContentUnion()
}

// LLMToolContent is the response-side carrier for the LLMToolContent union. It
// merges the fields of [ToolTextContent] and [ToolFileContent] and preserves
// access to the decoded union variant via [LLMToolContent.AsUnion].
//
// OpenAPI: LLMToolContent (anyOf ToolTextContent, ToolFileContent); JS SDK:
// LlmToolContent.
type LLMToolContent struct {
	Type  LLMToolContentType `json:"type,required"`
	Text  string             `json:"text"`
	URI   string             `json:"uri"`
	Mime  string             `json:"mime"`
	Name  string             `json:"name"`
	JSON  lLMToolContentJSON `json:"-"`
	union LLMToolContentUnion
}

// lLMToolContentJSON contains the JSON metadata for the struct [LLMToolContent]
type lLMToolContentJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	URI         apijson.Field
	Mime        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r lLMToolContentJSON) RawJSON() string {
	return r.raw
}

func (r *LLMToolContent) UnmarshalJSON(data []byte) (err error) {
	*r = LLMToolContent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [LLMToolContentUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ToolTextContent], [ToolFileContent].
func (r LLMToolContent) AsUnion() LLMToolContentUnion {
	return r.union
}

type LLMToolContentType string

const (
	LLMToolContentTypeText LLMToolContentType = "text"
	LLMToolContentTypeFile LLMToolContentType = "file"
)

func (r LLMToolContentType) IsKnown() bool {
	switch r {
	case LLMToolContentTypeText, LLMToolContentTypeFile:
		return true
	}
	return false
}

// ===== V2 Session Message Tool State Types =====

// V2SessionMessageToolState represents the state of a tool execution within a V2
// session assistant message content item.
//
// It is a discriminated union of the following types:
// [V2SessionMessageToolStatePending], [V2SessionMessageToolStateRunning],
// [V2SessionMessageToolStateCompleted], [V2SessionMessageToolStateError].
type V2SessionMessageToolState struct {
	Status V2SessionMessageToolStateStatus `json:"status,required"`
	JSON   v2SessionMessageToolStateJSON   `json:"-"`
	union  V2SessionMessageToolStateUnion
}

type v2SessionMessageToolStateJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2SessionMessageToolStateJSON) RawJSON() string {
	return r.raw
}

func (r *V2SessionMessageToolState) UnmarshalJSON(data []byte) (err error) {
	*r = V2SessionMessageToolState{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2SessionMessageToolStateUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [V2SessionMessageToolStatePending],
// [V2SessionMessageToolStateRunning], [V2SessionMessageToolStateCompleted],
// [V2SessionMessageToolStateError].
func (r V2SessionMessageToolState) AsUnion() V2SessionMessageToolStateUnion {
	return r.union
}

// Union satisfied by [V2SessionMessageToolStatePending],
// [V2SessionMessageToolStateRunning], [V2SessionMessageToolStateCompleted] or
// [V2SessionMessageToolStateError].
type V2SessionMessageToolStateUnion interface {
	implementsV2SessionMessageToolStateUnion()
}

type V2SessionMessageToolStateStatus string

const (
	V2SessionMessageToolStateStatusPending   V2SessionMessageToolStateStatus = "pending"
	V2SessionMessageToolStateStatusRunning   V2SessionMessageToolStateStatus = "running"
	V2SessionMessageToolStateStatusCompleted V2SessionMessageToolStateStatus = "completed"
	V2SessionMessageToolStateStatusError     V2SessionMessageToolStateStatus = "error"
)

func (r V2SessionMessageToolStateStatus) IsKnown() bool {
	switch r {
	case V2SessionMessageToolStateStatusPending,
		V2SessionMessageToolStateStatusRunning,
		V2SessionMessageToolStateStatusCompleted,
		V2SessionMessageToolStateStatusError:
		return true
	}
	return false
}

type V2SessionMessageToolStatePending struct {
	Status V2SessionMessageToolStatePendingStatus `json:"status,required"`
	Input  string                                 `json:"input,required"`
	JSON   v2SessionMessageToolStatePendingJSON   `json:"-"`
}

type V2SessionMessageToolStatePendingStatus string

const (
	V2SessionMessageToolStatePendingStatusPending V2SessionMessageToolStatePendingStatus = "pending"
)

func (r V2SessionMessageToolStatePendingStatus) IsKnown() bool {
	switch r {
	case V2SessionMessageToolStatePendingStatusPending:
		return true
	}
	return false
}

type v2SessionMessageToolStatePendingJSON struct {
	Status      apijson.Field
	Input       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageToolStatePending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageToolStatePendingJSON) RawJSON() string {
	return r.raw
}

func (r V2SessionMessageToolStatePending) implementsV2SessionMessageToolStateUnion() {}

type V2SessionMessageToolStateRunning struct {
	Status     V2SessionMessageToolStateRunningStatus `json:"status,required"`
	Input      map[string]any                         `json:"input,required"`
	Structured map[string]any                         `json:"structured,required"`
	// Each element can be [ToolTextContent] or [ToolFileContent].
	Content []LLMToolContent                     `json:"content,required"`
	JSON    v2SessionMessageToolStateRunningJSON `json:"-"`
}

type V2SessionMessageToolStateRunningStatus string

const (
	V2SessionMessageToolStateRunningStatusRunning V2SessionMessageToolStateRunningStatus = "running"
)

func (r V2SessionMessageToolStateRunningStatus) IsKnown() bool {
	switch r {
	case V2SessionMessageToolStateRunningStatusRunning:
		return true
	}
	return false
}

type v2SessionMessageToolStateRunningJSON struct {
	Status      apijson.Field
	Input       apijson.Field
	Structured  apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageToolStateRunning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageToolStateRunningJSON) RawJSON() string {
	return r.raw
}

func (r V2SessionMessageToolStateRunning) implementsV2SessionMessageToolStateUnion() {}

type V2SessionMessageToolStateCompleted struct {
	Status     V2SessionMessageToolStateCompletedStatus `json:"status,required"`
	Input      map[string]any                           `json:"input,required"`
	Structured map[string]any                           `json:"structured,required"`
	// Each element can be [ToolTextContent] or [ToolFileContent].
	Content     []LLMToolContent         `json:"content,required"`
	Attachments []V2PromptFileAttachment `json:"attachments"`
	OutputPaths []string                 `json:"outputPaths"`
	// This field is an untyped arbitrary value. The OpenAPI schema declares it as
	// an empty schema (`{}`), meaning it may hold any JSON value. Use a
	// type-switch or json.Unmarshal to inspect the runtime value.
	Result any                                    `json:"result"`
	JSON   v2SessionMessageToolStateCompletedJSON `json:"-"`
}

type V2SessionMessageToolStateCompletedStatus string

const (
	V2SessionMessageToolStateCompletedStatusCompleted V2SessionMessageToolStateCompletedStatus = "completed"
)

func (r V2SessionMessageToolStateCompletedStatus) IsKnown() bool {
	switch r {
	case V2SessionMessageToolStateCompletedStatusCompleted:
		return true
	}
	return false
}

type v2SessionMessageToolStateCompletedJSON struct {
	Status      apijson.Field
	Input       apijson.Field
	Structured  apijson.Field
	Content     apijson.Field
	Attachments apijson.Field
	OutputPaths apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageToolStateCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageToolStateCompletedJSON) RawJSON() string {
	return r.raw
}

func (r V2SessionMessageToolStateCompleted) implementsV2SessionMessageToolStateUnion() {}

type V2SessionMessageToolStateError struct {
	Status     V2SessionMessageToolStateErrorStatus `json:"status,required"`
	Input      map[string]any                       `json:"input,required"`
	Structured map[string]any                       `json:"structured,required"`
	// Each element can be [ToolTextContent] or [ToolFileContent].
	Content []LLMToolContent    `json:"content,required"`
	Error   SessionErrorUnknown `json:"error,required"`
	// This field is an untyped arbitrary value. The OpenAPI schema declares it as
	// an empty schema (`{}`), meaning it may hold any JSON value. Use a
	// type-switch or json.Unmarshal to inspect the runtime value.
	Result any                                `json:"result"`
	JSON   v2SessionMessageToolStateErrorJSON `json:"-"`
}

type V2SessionMessageToolStateErrorStatus string

const (
	V2SessionMessageToolStateErrorStatusError V2SessionMessageToolStateErrorStatus = "error"
)

func (r V2SessionMessageToolStateErrorStatus) IsKnown() bool {
	switch r {
	case V2SessionMessageToolStateErrorStatusError:
		return true
	}
	return false
}

type v2SessionMessageToolStateErrorJSON struct {
	Status      apijson.Field
	Input       apijson.Field
	Structured  apijson.Field
	Content     apijson.Field
	Error       apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageToolStateError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageToolStateErrorJSON) RawJSON() string {
	return r.raw
}

func (r V2SessionMessageToolStateError) implementsV2SessionMessageToolStateUnion() {}

// ===== V2SessionDurableEvent union =====

// V2SessionDurableEvent represents a discriminated union of durable session
// events. Possible runtime types of the union are
// [V2SessionDurableEventAgentSwitched], [V2SessionDurableEventModelSwitched],
// [V2SessionDurableEventMoved], [V2SessionDurableEventPrompted],
// [V2SessionDurableEventPromptAdmitted], [V2SessionDurableEventContextUpdated],
// [V2SessionDurableEventSynthetic], [V2SessionDurableEventShellStarted],
// [V2SessionDurableEventShellEnded], [V2SessionDurableEventStepStarted],
// [V2SessionDurableEventStepEnded], [V2SessionDurableEventStepFailed],
// [V2SessionDurableEventTextStarted], [V2SessionDurableEventTextEnded],
// [V2SessionDurableEventToolInputStarted], [V2SessionDurableEventToolInputEnded],
// [V2SessionDurableEventToolCalled], [V2SessionDurableEventToolProgress],
// [V2SessionDurableEventToolSuccess], [V2SessionDurableEventToolFailed],
// [V2SessionDurableEventReasoningStarted], [V2SessionDurableEventReasoningEnded],
// [V2SessionDurableEventRetried], [V2SessionDurableEventCompactionStarted],
// [V2SessionDurableEventCompactionEnded], [V2SessionDurableEventRevertStaged],
// [V2SessionDurableEventRevertCleared] or [V2SessionDurableEventRevertCommitted].
type V2SessionDurableEvent struct {
	JSON  v2SessionDurableEventJSON `json:"-"`
	union V2SessionDurableEventUnion
}

// v2SessionDurableEventJSON contains the JSON metadata for the struct
// [V2SessionDurableEvent].
type v2SessionDurableEventJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2SessionDurableEventJSON) RawJSON() string {
	return r.raw
}

func (r *V2SessionDurableEvent) UnmarshalJSON(data []byte) (err error) {
	*r = V2SessionDurableEvent{}
	// The OpenAPI durable-event union is a oneOf of object schemas only, so
	// any non-object payload (null, bool, number, string, array) cannot match
	// a variant. Skip union routing instead of failing, which would terminate
	// the whole SSE stream. RawJSON() still exposes the original value so
	// callers can distinguish a malformed frame from a routed event body.
	if !gjson.ParseBytes(data).IsObject() {
		r.JSON.raw = string(data)
		return nil
	}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2SessionDurableEventUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [V2SessionDurableEventAgentSwitched], [V2SessionDurableEventModelSwitched],
// [V2SessionDurableEventMoved], [V2SessionDurableEventPrompted],
// [V2SessionDurableEventPromptAdmitted], [V2SessionDurableEventContextUpdated],
// [V2SessionDurableEventSynthetic], [V2SessionDurableEventShellStarted],
// [V2SessionDurableEventShellEnded], [V2SessionDurableEventStepStarted],
// [V2SessionDurableEventStepEnded], [V2SessionDurableEventStepFailed],
// [V2SessionDurableEventTextStarted], [V2SessionDurableEventTextEnded],
// [V2SessionDurableEventToolInputStarted], [V2SessionDurableEventToolInputEnded],
// [V2SessionDurableEventToolCalled], [V2SessionDurableEventToolProgress],
// [V2SessionDurableEventToolSuccess], [V2SessionDurableEventToolFailed],
// [V2SessionDurableEventReasoningStarted], [V2SessionDurableEventReasoningEnded],
// [V2SessionDurableEventRetried], [V2SessionDurableEventCompactionStarted],
// [V2SessionDurableEventCompactionEnded], [V2SessionDurableEventRevertStaged],
// [V2SessionDurableEventRevertCleared] or [V2SessionDurableEventRevertCommitted].
func (r V2SessionDurableEvent) AsUnion() V2SessionDurableEventUnion {
	return r.union
}

// V2SessionDurableEventUnion is satisfied by
// [V2SessionDurableEventAgentSwitched], [V2SessionDurableEventModelSwitched],
// [V2SessionDurableEventMoved], [V2SessionDurableEventPrompted],
// [V2SessionDurableEventPromptAdmitted], [V2SessionDurableEventContextUpdated],
// [V2SessionDurableEventSynthetic], [V2SessionDurableEventShellStarted],
// [V2SessionDurableEventShellEnded], [V2SessionDurableEventStepStarted],
// [V2SessionDurableEventStepEnded], [V2SessionDurableEventStepFailed],
// [V2SessionDurableEventTextStarted], [V2SessionDurableEventTextEnded],
// [V2SessionDurableEventToolInputStarted], [V2SessionDurableEventToolInputEnded],
// [V2SessionDurableEventToolCalled], [V2SessionDurableEventToolProgress],
// [V2SessionDurableEventToolSuccess], [V2SessionDurableEventToolFailed],
// [V2SessionDurableEventReasoningStarted], [V2SessionDurableEventReasoningEnded],
// [V2SessionDurableEventRetried], [V2SessionDurableEventCompactionStarted],
// [V2SessionDurableEventCompactionEnded], [V2SessionDurableEventRevertStaged],
// [V2SessionDurableEventRevertCleared] or [V2SessionDurableEventRevertCommitted].
type V2SessionDurableEventUnion interface {
	implementsV2SessionDurableEvent()
}

// ===== V2SessionDurableEvent variants =====

type V2SessionDurableEventAgentSwitched struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                    `json:"metadata"`
	Type     V2EventSessionNextAgentSwitchedType    `json:"type,required"`
	Durable  V2EventDurable                         `json:"durable"`
	Location LocationRef                            `json:"location"`
	Data     V2EventSessionNextAgentSwitchedData    `json:"data,required"`
	JSON     v2SessionDurableEventAgentSwitchedJSON `json:"-"`
}

type v2SessionDurableEventAgentSwitchedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventAgentSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventAgentSwitchedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventAgentSwitched) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventModelSwitched struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                    `json:"metadata"`
	Type     V2EventSessionNextModelSwitchedType    `json:"type,required"`
	Durable  V2EventDurable                         `json:"durable"`
	Location LocationRef                            `json:"location"`
	Data     V2EventSessionNextModelSwitchedData    `json:"data,required"`
	JSON     v2SessionDurableEventModelSwitchedJSON `json:"-"`
}

type v2SessionDurableEventModelSwitchedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventModelSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventModelSwitchedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventModelSwitched) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventMoved struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                            `json:"metadata"`
	Type     V2EventSessionNextMovedType    `json:"type,required"`
	Durable  V2EventDurable                 `json:"durable"`
	Location LocationRef                    `json:"location"`
	Data     V2EventSessionNextMovedData    `json:"data,required"`
	JSON     v2SessionDurableEventMovedJSON `json:"-"`
}

type v2SessionDurableEventMovedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventMoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventMovedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventMoved) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventPrompted struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                               `json:"metadata"`
	Type     V2EventSessionNextPromptedType    `json:"type,required"`
	Durable  V2EventDurable                    `json:"durable"`
	Location LocationRef                       `json:"location"`
	Data     V2EventSessionNextPromptedData    `json:"data,required"`
	JSON     v2SessionDurableEventPromptedJSON `json:"-"`
}

type v2SessionDurableEventPromptedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventPrompted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventPromptedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventPrompted) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventPromptAdmitted struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                     `json:"metadata"`
	Type     V2EventSessionNextPromptAdmittedType    `json:"type,required"`
	Durable  V2EventDurable                          `json:"durable"`
	Location LocationRef                             `json:"location"`
	Data     V2EventSessionNextPromptAdmittedData    `json:"data,required"`
	JSON     v2SessionDurableEventPromptAdmittedJSON `json:"-"`
}

type v2SessionDurableEventPromptAdmittedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventPromptAdmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventPromptAdmittedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventPromptAdmitted) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventContextUpdated struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                     `json:"metadata"`
	Type     V2EventSessionNextContextUpdatedType    `json:"type,required"`
	Durable  V2EventDurable                          `json:"durable"`
	Location LocationRef                             `json:"location"`
	Data     V2EventSessionNextContextUpdatedData    `json:"data,required"`
	JSON     v2SessionDurableEventContextUpdatedJSON `json:"-"`
}

type v2SessionDurableEventContextUpdatedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventContextUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventContextUpdatedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventContextUpdated) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventSynthetic struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                `json:"metadata"`
	Type     V2EventSessionNextSyntheticType    `json:"type,required"`
	Durable  V2EventDurable                     `json:"durable"`
	Location LocationRef                        `json:"location"`
	Data     V2EventSessionNextSyntheticData    `json:"data,required"`
	JSON     v2SessionDurableEventSyntheticJSON `json:"-"`
}

type v2SessionDurableEventSyntheticJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventSynthetic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventSyntheticJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventSynthetic) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventShellStarted struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                   `json:"metadata"`
	Type     V2EventSessionNextShellStartedType    `json:"type,required"`
	Durable  V2EventDurable                        `json:"durable"`
	Location LocationRef                           `json:"location"`
	Data     V2EventSessionNextShellStartedData    `json:"data,required"`
	JSON     v2SessionDurableEventShellStartedJSON `json:"-"`
}

type v2SessionDurableEventShellStartedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventShellStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventShellStartedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventShellStarted) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventShellEnded struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                 `json:"metadata"`
	Type     V2EventSessionNextShellEndedType    `json:"type,required"`
	Durable  V2EventDurable                      `json:"durable"`
	Location LocationRef                         `json:"location"`
	Data     V2EventSessionNextShellEndedData    `json:"data,required"`
	JSON     v2SessionDurableEventShellEndedJSON `json:"-"`
}

type v2SessionDurableEventShellEndedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventShellEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventShellEndedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventShellEnded) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventStepStarted struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                  `json:"metadata"`
	Type     V2EventSessionNextStepStartedType    `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Data     V2EventSessionNextStepStartedData    `json:"data,required"`
	JSON     v2SessionDurableEventStepStartedJSON `json:"-"`
}

type v2SessionDurableEventStepStartedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventStepStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventStepStartedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventStepStarted) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventStepEnded struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                `json:"metadata"`
	Type     V2EventSessionNextStepEndedType    `json:"type,required"`
	Durable  V2EventDurable                     `json:"durable"`
	Location LocationRef                        `json:"location"`
	Data     V2EventSessionNextStepEndedData    `json:"data,required"`
	JSON     v2SessionDurableEventStepEndedJSON `json:"-"`
}

type v2SessionDurableEventStepEndedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventStepEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventStepEndedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventStepEnded) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventStepFailed struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                 `json:"metadata"`
	Type     V2EventSessionNextStepFailedType    `json:"type,required"`
	Durable  V2EventDurable                      `json:"durable"`
	Location LocationRef                         `json:"location"`
	Data     V2EventSessionNextStepFailedData    `json:"data,required"`
	JSON     v2SessionDurableEventStepFailedJSON `json:"-"`
}

type v2SessionDurableEventStepFailedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventStepFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventStepFailedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventStepFailed) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventTextStarted struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                  `json:"metadata"`
	Type     V2EventSessionNextTextStartedType    `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Data     V2EventSessionNextTextStartedData    `json:"data,required"`
	JSON     v2SessionDurableEventTextStartedJSON `json:"-"`
}

type v2SessionDurableEventTextStartedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventTextStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventTextStartedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventTextStarted) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventTextEnded struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                `json:"metadata"`
	Type     V2EventSessionNextTextEndedType    `json:"type,required"`
	Durable  V2EventDurable                     `json:"durable"`
	Location LocationRef                        `json:"location"`
	Data     V2EventSessionNextTextEndedData    `json:"data,required"`
	JSON     v2SessionDurableEventTextEndedJSON `json:"-"`
}

type v2SessionDurableEventTextEndedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventTextEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventTextEndedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventTextEnded) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventToolInputStarted struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                       `json:"metadata"`
	Type     V2EventSessionNextToolInputStartedType    `json:"type,required"`
	Durable  V2EventDurable                            `json:"durable"`
	Location LocationRef                               `json:"location"`
	Data     V2EventSessionNextToolInputStartedData    `json:"data,required"`
	JSON     v2SessionDurableEventToolInputStartedJSON `json:"-"`
}

type v2SessionDurableEventToolInputStartedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventToolInputStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventToolInputStartedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventToolInputStarted) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventToolInputEnded struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                     `json:"metadata"`
	Type     V2EventSessionNextToolInputEndedType    `json:"type,required"`
	Durable  V2EventDurable                          `json:"durable"`
	Location LocationRef                             `json:"location"`
	Data     V2EventSessionNextToolInputEndedData    `json:"data,required"`
	JSON     v2SessionDurableEventToolInputEndedJSON `json:"-"`
}

type v2SessionDurableEventToolInputEndedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventToolInputEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventToolInputEndedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventToolInputEnded) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventToolCalled struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                 `json:"metadata"`
	Type     V2EventSessionNextToolCalledType    `json:"type,required"`
	Durable  V2EventDurable                      `json:"durable"`
	Location LocationRef                         `json:"location"`
	Data     V2EventSessionNextToolCalledData    `json:"data,required"`
	JSON     v2SessionDurableEventToolCalledJSON `json:"-"`
}

type v2SessionDurableEventToolCalledJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventToolCalled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventToolCalledJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventToolCalled) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventToolProgress struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                   `json:"metadata"`
	Type     V2EventSessionNextToolProgressType    `json:"type,required"`
	Durable  V2EventDurable                        `json:"durable"`
	Location LocationRef                           `json:"location"`
	Data     V2EventSessionNextToolProgressData    `json:"data,required"`
	JSON     v2SessionDurableEventToolProgressJSON `json:"-"`
}

type v2SessionDurableEventToolProgressJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventToolProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventToolProgressJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventToolProgress) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventToolSuccess struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                  `json:"metadata"`
	Type     V2EventSessionNextToolSuccessType    `json:"type,required"`
	Durable  V2EventDurable                       `json:"durable"`
	Location LocationRef                          `json:"location"`
	Data     V2EventSessionNextToolSuccessData    `json:"data,required"`
	JSON     v2SessionDurableEventToolSuccessJSON `json:"-"`
}

type v2SessionDurableEventToolSuccessJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventToolSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventToolSuccessJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventToolSuccess) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventToolFailed struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                 `json:"metadata"`
	Type     V2EventSessionNextToolFailedType    `json:"type,required"`
	Durable  V2EventDurable                      `json:"durable"`
	Location LocationRef                         `json:"location"`
	Data     V2EventSessionNextToolFailedData    `json:"data,required"`
	JSON     v2SessionDurableEventToolFailedJSON `json:"-"`
}

type v2SessionDurableEventToolFailedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventToolFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventToolFailedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventToolFailed) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventReasoningStarted struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                       `json:"metadata"`
	Type     V2EventSessionNextReasoningStartedType    `json:"type,required"`
	Durable  V2EventDurable                            `json:"durable"`
	Location LocationRef                               `json:"location"`
	Data     V2EventSessionNextReasoningStartedData    `json:"data,required"`
	JSON     v2SessionDurableEventReasoningStartedJSON `json:"-"`
}

type v2SessionDurableEventReasoningStartedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventReasoningStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventReasoningStartedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventReasoningStarted) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventReasoningEnded struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                     `json:"metadata"`
	Type     V2EventSessionNextReasoningEndedType    `json:"type,required"`
	Durable  V2EventDurable                          `json:"durable"`
	Location LocationRef                             `json:"location"`
	Data     V2EventSessionNextReasoningEndedData    `json:"data,required"`
	JSON     v2SessionDurableEventReasoningEndedJSON `json:"-"`
}

type v2SessionDurableEventReasoningEndedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventReasoningEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventReasoningEndedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventReasoningEnded) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventRetried struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                              `json:"metadata"`
	Type     V2EventSessionNextRetriedType    `json:"type,required"`
	Durable  V2EventDurable                   `json:"durable"`
	Location LocationRef                      `json:"location"`
	Data     V2EventSessionNextRetriedData    `json:"data,required"`
	JSON     v2SessionDurableEventRetriedJSON `json:"-"`
}

type v2SessionDurableEventRetriedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventRetried) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventRetriedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventRetried) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventCompactionStarted struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                        `json:"metadata"`
	Type     V2EventSessionNextCompactionStartedType    `json:"type,required"`
	Durable  V2EventDurable                             `json:"durable"`
	Location LocationRef                                `json:"location"`
	Data     V2EventSessionNextCompactionStartedData    `json:"data,required"`
	JSON     v2SessionDurableEventCompactionStartedJSON `json:"-"`
}

type v2SessionDurableEventCompactionStartedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventCompactionStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventCompactionStartedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventCompactionStarted) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventCompactionEnded struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                      `json:"metadata"`
	Type     V2EventSessionNextCompactionEndedType    `json:"type,required"`
	Durable  V2EventDurable                           `json:"durable"`
	Location LocationRef                              `json:"location"`
	Data     V2EventSessionNextCompactionEndedData    `json:"data,required"`
	JSON     v2SessionDurableEventCompactionEndedJSON `json:"-"`
}

type v2SessionDurableEventCompactionEndedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventCompactionEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventCompactionEndedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventCompactionEnded) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventRevertStaged struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                   `json:"metadata"`
	Type     V2EventSessionNextRevertStagedType    `json:"type,required"`
	Durable  V2EventDurable                        `json:"durable"`
	Location LocationRef                           `json:"location"`
	Data     V2EventSessionNextRevertStagedData    `json:"data,required"`
	JSON     v2SessionDurableEventRevertStagedJSON `json:"-"`
}

type v2SessionDurableEventRevertStagedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventRevertStaged) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventRevertStagedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventRevertStaged) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventRevertCleared struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                    `json:"metadata"`
	Type     V2EventSessionNextRevertClearedType    `json:"type,required"`
	Durable  V2EventDurable                         `json:"durable"`
	Location LocationRef                            `json:"location"`
	Data     V2EventSessionNextRevertClearedData    `json:"data,required"`
	JSON     v2SessionDurableEventRevertClearedJSON `json:"-"`
}

type v2SessionDurableEventRevertClearedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventRevertCleared) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventRevertClearedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventRevertCleared) implementsV2SessionDurableEvent() {}

type V2SessionDurableEventRevertCommitted struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                      `json:"metadata"`
	Type     V2EventSessionNextRevertCommittedType    `json:"type,required"`
	Durable  V2EventDurable                           `json:"durable"`
	Location LocationRef                              `json:"location"`
	Data     V2EventSessionNextRevertCommittedData    `json:"data,required"`
	JSON     v2SessionDurableEventRevertCommittedJSON `json:"-"`
}

type v2SessionDurableEventRevertCommittedJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Type        apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionDurableEventRevertCommitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionDurableEventRevertCommittedJSON) RawJSON() string {
	return r.raw
}

func (V2SessionDurableEventRevertCommitted) implementsV2SessionDurableEvent() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[LLMToolContentUnion](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "text",
			Type:               reflect.TypeFor[ToolTextContent](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "file",
			Type:               reflect.TypeFor[ToolFileContent](),
		},
	)
	apijson.RegisterUnion(
		reflect.TypeFor[V2SessionMessageToolStateUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2SessionMessageToolStatePending](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2SessionMessageToolStateRunning](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2SessionMessageToolStateCompleted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[V2SessionMessageToolStateError](),
		},
	)
	apijson.RegisterUnion(
		reflect.TypeFor[V2SessionMessageAssistantContentUnion](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "text",
			Type:               reflect.TypeFor[V2SessionMessageAssistantTextContent](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "reasoning",
			Type:               reflect.TypeFor[V2SessionMessageAssistantReasoningContent](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "tool",
			Type:               reflect.TypeFor[V2SessionMessageAssistantToolContent](),
		},
	)
	apijson.RegisterUnion(
		reflect.TypeFor[V2SessionMessageUnion](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "agent-switched",
			Type:               reflect.TypeFor[V2SessionMessageAgentSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "model-switched",
			Type:               reflect.TypeFor[V2SessionMessageModelSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "user",
			Type:               reflect.TypeFor[V2SessionMessageUser](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "synthetic",
			Type:               reflect.TypeFor[V2SessionMessageSynthetic](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "shell",
			Type:               reflect.TypeFor[V2SessionMessageShell](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "assistant",
			Type:               reflect.TypeFor[V2SessionMessageAssistant](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "compaction",
			Type:               reflect.TypeFor[V2SessionMessageCompaction](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "system",
			Type:               reflect.TypeFor[V2SessionMessageSystem](),
		},
	)
	apijson.RegisterUnion(
		reflect.TypeFor[V2SessionDurableEventUnion](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.agent.switched",
			Type:               reflect.TypeFor[V2SessionDurableEventAgentSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.model.switched",
			Type:               reflect.TypeFor[V2SessionDurableEventModelSwitched](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.moved",
			Type:               reflect.TypeFor[V2SessionDurableEventMoved](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.prompted",
			Type:               reflect.TypeFor[V2SessionDurableEventPrompted](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.prompt.admitted",
			Type:               reflect.TypeFor[V2SessionDurableEventPromptAdmitted](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.context.updated",
			Type:               reflect.TypeFor[V2SessionDurableEventContextUpdated](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.synthetic",
			Type:               reflect.TypeFor[V2SessionDurableEventSynthetic](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.shell.started",
			Type:               reflect.TypeFor[V2SessionDurableEventShellStarted](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.shell.ended",
			Type:               reflect.TypeFor[V2SessionDurableEventShellEnded](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.step.started",
			Type:               reflect.TypeFor[V2SessionDurableEventStepStarted](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.step.ended",
			Type:               reflect.TypeFor[V2SessionDurableEventStepEnded](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.step.failed",
			Type:               reflect.TypeFor[V2SessionDurableEventStepFailed](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.text.started",
			Type:               reflect.TypeFor[V2SessionDurableEventTextStarted](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.text.ended",
			Type:               reflect.TypeFor[V2SessionDurableEventTextEnded](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.tool.input.started",
			Type:               reflect.TypeFor[V2SessionDurableEventToolInputStarted](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.tool.input.ended",
			Type:               reflect.TypeFor[V2SessionDurableEventToolInputEnded](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.tool.called",
			Type:               reflect.TypeFor[V2SessionDurableEventToolCalled](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.tool.progress",
			Type:               reflect.TypeFor[V2SessionDurableEventToolProgress](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.tool.success",
			Type:               reflect.TypeFor[V2SessionDurableEventToolSuccess](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.tool.failed",
			Type:               reflect.TypeFor[V2SessionDurableEventToolFailed](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.reasoning.started",
			Type:               reflect.TypeFor[V2SessionDurableEventReasoningStarted](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.reasoning.ended",
			Type:               reflect.TypeFor[V2SessionDurableEventReasoningEnded](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.retried",
			Type:               reflect.TypeFor[V2SessionDurableEventRetried](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.compaction.started",
			Type:               reflect.TypeFor[V2SessionDurableEventCompactionStarted](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.compaction.ended",
			Type:               reflect.TypeFor[V2SessionDurableEventCompactionEnded](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.revert.staged",
			Type:               reflect.TypeFor[V2SessionDurableEventRevertStaged](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.revert.cleared",
			Type:               reflect.TypeFor[V2SessionDurableEventRevertCleared](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "session.next.revert.committed",
			Type:               reflect.TypeFor[V2SessionDurableEventRevertCommitted](),
		},
	)
}

// ===== Param Types =====

type V2SessionOrder string

const (
	V2SessionOrderAsc  V2SessionOrder = "asc"
	V2SessionOrderDesc V2SessionOrder = "desc"
)

func (r V2SessionOrder) IsKnown() bool {
	switch r {
	case V2SessionOrderAsc, V2SessionOrderDesc:
		return true
	}
	return false
}

type V2SessionListParams struct {
	Directory param.Field[string]         `query:"directory"`
	Workspace param.Field[string]         `query:"workspace"`
	Limit     param.Field[int64]          `query:"limit"`
	Order     param.Field[V2SessionOrder] `query:"order"`
	Project   param.Field[string]         `query:"project"`
	Subpath   param.Field[string]         `query:"subpath"`
	Search    param.Field[string]         `query:"search"`
	Cursor    param.Field[string]         `query:"cursor"`
}

func (r V2SessionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2SessionPromptParams struct {
	ID       param.Field[string]             `json:"id"`
	Prompt   param.Field[V2PromptInputParam] `json:"prompt,required"`
	Delivery param.Field[SessionDelivery]    `json:"delivery"`
	Resume   param.Field[bool]               `json:"resume"`
}

func (r V2SessionPromptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionDelivery string

const (
	SessionDeliverySteer SessionDelivery = "steer"
	SessionDeliveryQueue SessionDelivery = "queue"
)

func (r SessionDelivery) IsKnown() bool {
	switch r {
	case SessionDeliverySteer, SessionDeliveryQueue:
		return true
	}
	return false
}

// V2PromptInputParam represents the prompt input for the v2.session.prompt
// request body.
type V2PromptInputParam struct {
	Text   param.Field[string]                             `json:"text,required"`
	Files  param.Field[[]V2PromptInputFileAttachmentParam] `json:"files"`
	Agents param.Field[[]V2PromptAgentAttachmentParam]     `json:"agents"`
}

func (r V2PromptInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// V2PromptInputFileAttachmentParam represents a file attachment on a
// [V2PromptInputParam].
type V2PromptInputFileAttachmentParam struct {
	URI         param.Field[string]              `json:"uri,required"`
	Name        param.Field[string]              `json:"name"`
	Description param.Field[string]              `json:"description"`
	Source      param.Field[V2PromptSourceParam] `json:"source"`
}

func (r V2PromptInputFileAttachmentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2PromptAgentAttachmentParam struct {
	Name   param.Field[string]              `json:"name,required"`
	Source param.Field[V2PromptSourceParam] `json:"source"`
}

func (r V2PromptAgentAttachmentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2PromptSourceParam struct {
	Start param.Field[int64]  `json:"start,required"`
	End   param.Field[int64]  `json:"end,required"`
	Text  param.Field[string] `json:"text,required"`
}

func (r V2PromptSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2SessionMessagesParams struct {
	Limit  param.Field[int64]          `query:"limit"`
	Order  param.Field[V2SessionOrder] `query:"order"`
	Cursor param.Field[string]         `query:"cursor"`
}

func (r V2SessionMessagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ===== LocationRef & RevertState =====

// V2SessionMessageAssistantReasoningContentTime contains the start and end
// timestamps of an assistant reasoning content block.
type V2SessionMessageAssistantReasoningContentTime struct {
	Created   int64                                             `json:"created,required"`
	Completed int64                                             `json:"completed"`
	JSON      v2SessionMessageAssistantReasoningContentTimeJSON `json:"-"`
}

// v2SessionMessageAssistantReasoningContentTimeJSON contains the JSON
// metadata for the struct [V2SessionMessageAssistantReasoningContentTime]
type v2SessionMessageAssistantReasoningContentTimeJSON struct {
	Created     apijson.Field
	Completed   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageAssistantReasoningContentTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAssistantReasoningContentTimeJSON) RawJSON() string {
	return r.raw
}

// LocationRef represents a reference to a location in a workspace.
type LocationRef struct {
	Directory   string          `json:"directory,required"`
	WorkspaceID string          `json:"workspaceID"`
	JSON        locationRefJSON `json:"-"`
}

// locationRefJSON contains the JSON metadata for the struct [LocationRef]
type locationRefJSON struct {
	Directory   apijson.Field
	WorkspaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationRefJSON) RawJSON() string {
	return r.raw
}

// LocationRefParam is the request-side equivalent of [LocationRef].
// All fields use param.Field[T] so unset optional fields are omitted from the
// serialised JSON body rather than being sent as zero values.
//
// Note: this is distinct from [V2LocationParam], which uses query-string tags
// and a different field name ("workspace" vs "workspaceID") for URL parameters.
//
// OpenAPI: LocationRef — required: [directory]; optional: [workspaceID]
type LocationRefParam struct {
	Directory   param.Field[string] `json:"directory,required"`
	WorkspaceID param.Field[string] `json:"workspaceID"`
}

func (r LocationRefParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// RevertState represents a revert state.
type RevertState struct {
	MessageID string          `json:"messageID,required"`
	PartID    string          `json:"partID"`
	Snapshot  string          `json:"snapshot"`
	Diff      string          `json:"diff"`
	Files     []FileDiff      `json:"files"`
	JSON      revertStateJSON `json:"-"`
}

// revertStateJSON contains the JSON metadata for the struct [RevertState]
type revertStateJSON struct {
	MessageID   apijson.Field
	PartID      apijson.Field
	Snapshot    apijson.Field
	Diff        apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RevertState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r revertStateJSON) RawJSON() string {
	return r.raw
}

// ===== FileDiff =====

// FileDiff represents a diff for a file.
type FileDiff struct {
	Path      string         `json:"path,required"`
	Status    FileDiffStatus `json:"status,required"`
	Additions int64          `json:"additions,required"`
	Deletions int64          `json:"deletions,required"`
	Patch     string         `json:"patch,required"`
	JSON      fileDiffJSON   `json:"-"`
}

// fileDiffJSON contains the JSON metadata for the struct [FileDiff]
type fileDiffJSON struct {
	Path        apijson.Field
	Status      apijson.Field
	Additions   apijson.Field
	Deletions   apijson.Field
	Patch       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileDiffJSON) RawJSON() string {
	return r.raw
}

// FileDiffStatus represents the status of a file diff.
type FileDiffStatus string

const (
	FileDiffStatusAdded    FileDiffStatus = "added"
	FileDiffStatusModified FileDiffStatus = "modified"
	FileDiffStatusDeleted  FileDiffStatus = "deleted"
)

func (r FileDiffStatus) IsKnown() bool {
	switch r {
	case FileDiffStatusAdded, FileDiffStatusModified, FileDiffStatusDeleted:
		return true
	}
	return false
}

// ===== V2SessionNewParams =====

type V2SessionNewParams struct {
	ID       param.Field[string]           `json:"id"`
	Agent    param.Field[string]           `json:"agent"`
	Model    param.Field[ModelRefParam]    `json:"model"`
	Location param.Field[LocationRefParam] `json:"location"`
}

func (r V2SessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ===== V2SessionSwitchAgentParams =====

type V2SessionSwitchAgentParams struct {
	Agent param.Field[string] `json:"agent,required"`
}

func (r V2SessionSwitchAgentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ===== V2SessionSwitchModelParams =====

type V2SessionSwitchModelParams struct {
	Model param.Field[ModelRefParam] `json:"model,required"`
}

func (r V2SessionSwitchModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ===== V2SessionEventsParams =====

type V2SessionEventsParams struct {
	After param.Field[string] `query:"after"`
}

func (r V2SessionEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ===== V2SessionHistoryParams =====

type V2SessionHistoryParams struct {
	Limit param.Field[int64] `query:"limit"`
	After param.Field[int64] `query:"after"`
}

func (r V2SessionHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ===== V2SessionCreateResponse =====

type V2SessionCreateResponse struct {
	Data V2SessionInfo               `json:"data,required"`
	JSON v2SessionCreateResponseJSON `json:"-"`
}

type v2SessionCreateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionCreateResponseJSON) RawJSON() string {
	return r.raw
}

// ===== V2SessionGetResponse =====

// V2SessionGetResponse is returned by the Get method.
type V2SessionGetResponse struct {
	Data V2SessionInfo            `json:"data,required"`
	JSON v2SessionGetResponseJSON `json:"-"`
}

type v2SessionGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionGetResponseJSON) RawJSON() string {
	return r.raw
}

// ===== V2SessionEventResponse (SSE) =====

// V2SessionEventResponse represents an SSE event from the session event stream.
type V2SessionEventResponse struct {
	ID    string                     `json:"id,required"`
	Event string                     `json:"event,required"`
	Data  string                     `json:"data,required"`
	JSON  v2SessionEventResponseJSON `json:"-"`
}

type v2SessionEventResponseJSON struct {
	ID          apijson.Field
	Event       apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionEventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionEventResponseJSON) RawJSON() string {
	return r.raw
}

// ParseData deserializes the [V2SessionEventResponse.Data] field (which is
// encoded as a JSON string in SSE events) into a [V2SessionDurableEvent]
// value.
//
// Returns the deserialized event union, or an error if [Data] is empty or
// contains malformed JSON.
func (r V2SessionEventResponse) ParseData() (V2SessionDurableEvent, error) {
	var result V2SessionDurableEvent
	if r.Data == "" {
		return result, errors.New("V2SessionEventResponse.Data is empty")
	}
	err := json.Unmarshal([]byte(r.Data), &result)
	if err != nil {
		return result, fmt.Errorf("V2SessionEventResponse.ParseData: %w", err)
	}
	return result, nil
}

// ===== V2SessionHistoryResponse =====

// V2SessionHistoryResponse is returned by the History method. It is a finite page
// of durable session events: use its GetNextPage method, or
// [V2SessionService.HistoryAutoPaging], to keep reading while HasMore is true.
type V2SessionHistoryResponse = pagination.SeqPage[V2SessionDurableEvent]

// ===== V2SessionMessageSystem =====

// V2SessionMessageSystem represents a system message in the V2 session.
type V2SessionMessageSystem struct {
	ID   string                     `json:"id,required"`
	Time V2SessionMessageTime       `json:"time,required"`
	Type V2SessionMessageSystemType `json:"type,required"`
	Text string                     `json:"text,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                        `json:"metadata"`
	JSON     v2SessionMessageSystemJSON `json:"-"`
}

type v2SessionMessageSystemJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	Text        apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageSystem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageSystemJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageSystemType string

const (
	V2SessionMessageSystemTypeSystem V2SessionMessageSystemType = "system"
)

func (r V2SessionMessageSystemType) IsKnown() bool {
	switch r {
	case V2SessionMessageSystemTypeSystem:
		return true
	}
	return false
}

// ===== V2SessionActiveResponse =====

// V2SessionActiveResponse is returned by the Active method. It contains a map of
// session IDs to their active state.
type V2SessionActiveResponse struct {
	Data map[string]any              `json:"data,required"`
	JSON v2SessionActiveResponseJSON `json:"-"`
}

type v2SessionActiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionActiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionActiveResponseJSON) RawJSON() string {
	return r.raw
}
