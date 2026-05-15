// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
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
	"github.com/tidwall/gjson"
)

// V2SessionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2SessionService] method instead.
type V2SessionService struct {
	Options []option.RequestOption
}

// NewV2SessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2SessionService(opts ...option.RequestOption) (r *V2SessionService) {
	r = &V2SessionService{}
	r.Options = opts
	return
}

// List v2 sessions
//
// Retrieve sessions in the requested order. Items keep that order across pages;
// use cursor.next or cursor.previous to move through the ordered list.
func (r *V2SessionService) List(ctx context.Context, query V2SessionListParams, opts ...option.RequestOption) (res *V2SessionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
func (r *V2SessionService) Compact(ctx context.Context, sessionID string, query V2SessionCompactParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/compact", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, nil, opts...)
	return
}

// Wait for v2 session
//
// Wait for a v2 session agent loop to become idle.
func (r *V2SessionService) Wait(ctx context.Context, sessionID string, query V2SessionWaitParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/wait", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, nil, opts...)
	return
}

// Get v2 session context
//
// Retrieve the active context messages for a v2 session (all messages after the
// last compaction).
func (r *V2SessionService) Context(ctx context.Context, sessionID string, query V2SessionContextParams, opts ...option.RequestOption) (res *[]V2SessionMessage, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/context", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get v2 session messages
//
// Retrieve projected v2 messages for a session. Items keep the requested order
// across pages; use cursor.next or cursor.previous to move through the ordered
// timeline.
func (r *V2SessionService) Messages(ctx context.Context, sessionID string, query V2SessionMessagesParams, opts ...option.RequestOption) (res *V2SessionMessagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/message", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// ===== Response Types =====

type V2SessionsResponse struct {
	Items  []V2SessionInfo         `json:"items,required"`
	Cursor V2Cursor              `json:"cursor,required"`
	JSON   v2SessionsResponseJSON `json:"-"`
}

type v2SessionsResponseJSON struct {
	Items       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionsResponseJSON) RawJSON() string {
	return r.raw
}

type V2SessionInfo struct {
	ID          string               `json:"id,required"`
	ProjectID   string               `json:"projectID,required"`
	Cost        float64              `json:"cost,required"`
	Time        V2SessionInfoTime    `json:"time,required"`
	Title       string               `json:"title,required"`
	Tokens      V2SessionInfoTokens  `json:"tokens,required"`
	ParentID    string               `json:"parentID"`
	WorkspaceID string               `json:"workspaceID"`
	Path        string               `json:"path"`
	Agent       string               `json:"agent"`
	Model       V2SessionInfoModel   `json:"model"`
	JSON        v2SessionInfoJSON    `json:"-"`
}

type v2SessionInfoJSON struct {
	ID          apijson.Field
	ProjectID   apijson.Field
	Cost        apijson.Field
	Time        apijson.Field
	Title       apijson.Field
	Tokens      apijson.Field
	ParentID    apijson.Field
	WorkspaceID apijson.Field
	Path        apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
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
	Created  float64            `json:"created,required"`
	Updated  float64            `json:"updated,required"`
	Archived float64            `json:"archived"`
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

type V2SessionInfoModel struct {
	ID         string              `json:"id,required"`
	ProviderID string              `json:"providerID,required"`
	Variant    string              `json:"variant,required"`
	JSON       v2SessionInfoModelJSON `json:"-"`
}

type v2SessionInfoModelJSON struct {
	ID          apijson.Field
	ProviderID  apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionInfoModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionInfoModelJSON) RawJSON() string {
	return r.raw
}

type V2SessionInfoTokens struct {
	Input     float64                    `json:"input,required"`
	Output    float64                    `json:"output,required"`
	Reasoning float64                    `json:"reasoning,required"`
	Cache     V2SessionInfoTokensCache   `json:"cache,required"`
	JSON      v2SessionInfoTokensJSON    `json:"-"`
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
	Read  float64                        `json:"read,required"`
	Write float64                        `json:"write,required"`
	JSON  v2SessionInfoTokensCacheJSON   `json:"-"`
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

type V2Cursor struct {
	Previous string     `json:"previous"`
	Next     string     `json:"next"`
	JSON     v2CursorJSON `json:"-"`
}

type v2CursorJSON struct {
	Previous    apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2Cursor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2CursorJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessagesResponse struct {
	Items  []V2SessionMessage            `json:"items,required"`
	Cursor V2Cursor                      `json:"cursor,required"`
	JSON   v2SessionMessagesResponseJSON `json:"-"`
}

type v2SessionMessagesResponseJSON struct {
	Items       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessagesResponseJSON) RawJSON() string {
	return r.raw
}

// V2SessionMessage is a union type that can be one of the following concrete
// types: V2SessionMessageAgentSwitched, V2SessionMessageModelSwitched,
// V2SessionMessageUser, V2SessionMessageSynthetic, V2SessionMessageShell,
// V2SessionMessageAssistant, V2SessionMessageCompaction.
type V2SessionMessage interface {
	implementsV2SessionMessage()
}

func (V2SessionMessageAgentSwitched) implementsV2SessionMessage() {}
func (V2SessionMessageModelSwitched) implementsV2SessionMessage() {}
func (V2SessionMessageUser) implementsV2SessionMessage()          {}
func (V2SessionMessageSynthetic) implementsV2SessionMessage()     {}
func (V2SessionMessageShell) implementsV2SessionMessage()         {}
func (V2SessionMessageAssistant) implementsV2SessionMessage()     {}
func (V2SessionMessageCompaction) implementsV2SessionMessage()    {}

type V2SessionMessageAgentSwitched struct {
	ID       string                            `json:"id,required"`
	Time     V2SessionMessageTime              `json:"time,required"`
	Type     string                            `json:"type,required"`
	Agent    string                            `json:"agent,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{}                       `json:"metadata"`
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

type V2SessionMessageModelSwitched struct {
	ID       string                            `json:"id,required"`
	Time     V2SessionMessageTime              `json:"time,required"`
	Type     string                            `json:"type,required"`
	Model    V2SessionMessageModel             `json:"model,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{}                       `json:"metadata"`
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

type V2SessionMessageUser struct {
	ID         string                     `json:"id,required"`
	Time       V2SessionMessageTime       `json:"time,required"`
	Text       string                     `json:"text,required"`
	Type       string                     `json:"type,required"`
	Files      []V2PromptFileAttachment     `json:"files"`
	Agents     []V2PromptAgentAttachment    `json:"agents"`
	References []V2PromptReferenceAttachment `json:"references"`
	Metadata   interface{}                `json:"metadata"`
	JSON       v2SessionMessageUserJSON   `json:"-"`
}

type v2SessionMessageUserJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	Files       apijson.Field
	Agents      apijson.Field
	References  apijson.Field
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

type V2SessionMessageSynthetic struct {
	ID        string                          `json:"id,required"`
	Time      V2SessionMessageTime            `json:"time,required"`
	SessionID string                          `json:"sessionID,required"`
	Text      string                          `json:"text,required"`
	Type      string                          `json:"type,required"`
	Metadata  interface{}                     `json:"metadata"`
	JSON      v2SessionMessageSyntheticJSON   `json:"-"`
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

type V2SessionMessageShell struct {
	ID       string                       `json:"id,required"`
	Time     V2SessionMessageShellTime    `json:"time,required"`
	Type     string                       `json:"type,required"`
	CallID   string                       `json:"callID,required"`
	Command  string                       `json:"command,required"`
	Output   string                       `json:"output,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{}                  `json:"metadata"`
	JSON     v2SessionMessageShellJSON    `json:"-"`
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

type V2SessionMessageAssistant struct {
	ID        string                              `json:"id,required"`
	Time      V2SessionMessageAssistantTime       `json:"time,required"`
	Type      string                              `json:"type,required"`
	Agent     string                              `json:"agent,required"`
	Model     V2SessionMessageModel               `json:"model,required"`
	Content   []V2SessionMessageAssistantContent  `json:"content,required"`
	Snapshot  V2SessionMessageAssistantSnapshot   `json:"snapshot"`
	Finish    string                              `json:"finish"`
	Cost      float64                             `json:"cost"`
	Tokens    V2SessionMessageTokens              `json:"tokens"`
	Error     SessionErrorUnknown                 `json:"error"`
	Metadata  interface{}                         `json:"metadata"`
	JSON      v2SessionMessageAssistantJSON       `json:"-"`
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

type V2SessionMessageCompaction struct {
	ID       string                                    `json:"id,required"`
	Time     V2SessionMessageTime                      `json:"time,required"`
	Type     string                                    `json:"type,required"`
	Reason   V2SessionMessageCompactionReason           `json:"reason,required"`
	Summary  string                                    `json:"summary,required"`
	Include  string                                    `json:"include"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{}                               `json:"metadata"`
	JSON     v2SessionMessageCompactionJSON            `json:"-"`
}

type v2SessionMessageCompactionJSON struct {
	ID          apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	Reason      apijson.Field
	Summary     apijson.Field
	Include     apijson.Field
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
	Type    string                    `json:"type,required"`
	Message string                    `json:"message,required"`
	JSON    sessionErrorUnknownJSON   `json:"-"`
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

// Shared sub-types

type V2SessionMessageTime struct {
	Created float64                  `json:"created,required"`
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
	Created   float64                       `json:"created,required"`
	Completed float64                       `json:"completed"`
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
	Created   float64                            `json:"created,required"`
	Completed float64                            `json:"completed"`
	JSON      v2SessionMessageAssistantTimeJSON  `json:"-"`
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

type V2SessionMessageModel struct {
	ID         string                     `json:"id,required"`
	ProviderID string                     `json:"providerID,required"`
	Variant    string                     `json:"variant,required"`
	JSON       v2SessionMessageModelJSON  `json:"-"`
}

type v2SessionMessageModelJSON struct {
	ID          apijson.Field
	ProviderID  apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageModelJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageTokens struct {
	Input     float64                      `json:"input,required"`
	Output    float64                      `json:"output,required"`
	Reasoning float64                      `json:"reasoning,required"`
	Cache     V2SessionMessageTokensCache  `json:"cache,required"`
	JSON      v2SessionMessageTokensJSON   `json:"-"`
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
	Read  float64                          `json:"read,required"`
	Write float64                          `json:"write,required"`
	JSON  v2SessionMessageTokensCacheJSON  `json:"-"`
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

type V2SessionMessageAssistantContent struct {
	Type     string                                     `json:"type,required"`
	Text     string                                     `json:"text"`
	ID       string                                     `json:"id"`
	Name     string                                     `json:"name"`
	State    V2SessionMessageToolState                  `json:"state"`
	Provider V2SessionMessageToolProvider               `json:"provider"`
	Time     V2SessionMessageToolTime                   `json:"time"`
	JSON     v2SessionMessageAssistantContentJSON       `json:"-"`
}

type v2SessionMessageAssistantContentJSON struct {
	Type        apijson.Field
	Text        apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	State       apijson.Field
	Provider    apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageAssistantContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAssistantContentJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageAssistantSnapshot struct {
	Start string                                `json:"start"`
	End   string                                `json:"end"`
	JSON  v2SessionMessageAssistantSnapshotJSON `json:"-"`
}

type v2SessionMessageAssistantSnapshotJSON struct {
	Start       apijson.Field
	End         apijson.Field
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
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{}                         `json:"metadata"`
	JSON     v2SessionMessageToolProviderJSON    `json:"-"`
}

type v2SessionMessageToolProviderJSON struct {
	Executed    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageToolProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageToolProviderJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageToolTime struct {
	Created   float64                        `json:"created,required"`
	Ran       float64                        `json:"ran"`
	Completed float64                        `json:"completed"`
	Pruned    float64                        `json:"pruned"`
	JSON      v2SessionMessageToolTimeJSON   `json:"-"`
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
// (V2SessionMessageAssistantToolContent).
//
// The [V2SessionMessageAssistantContent] struct is a fat struct containing all
// possible fields across content types; the concrete sub-types provide
// type-safe alternatives.

// V2SessionPromptResponse is returned by the Prompt method. It represents a
// V2SessionMessage union type.
type V2SessionPromptResponse = V2SessionMessage

// ===== Prompt Attachment Types =====

type V2PromptFileAttachment struct {
	URI         string                       `json:"uri,required"`
	Mime        string                       `json:"mime,required"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Source      V2PromptSource               `json:"source"`
	JSON        v2PromptFileAttachmentJSON   `json:"-"`
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
	Name   string                     `json:"name,required"`
	Source V2PromptSource             `json:"source"`
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

type V2PromptReferenceAttachment struct {
	Name       string                              `json:"name,required"`
	Kind       V2PromptReferenceAttachmentKind      `json:"kind,required"`
	URI        string                              `json:"uri"`
	Repository string                           `json:"repository"`
	Branch     string                           `json:"branch"`
	Target     string                           `json:"target"`
	TargetURI  string                           `json:"targetUri"`
	Problem    string                           `json:"problem"`
	Source     V2PromptSource                   `json:"source"`
	JSON       v2PromptReferenceAttachmentJSON  `json:"-"`
}

type v2PromptReferenceAttachmentJSON struct {
	Name        apijson.Field
	Kind        apijson.Field
	URI         apijson.Field
	Repository  apijson.Field
	Branch      apijson.Field
	Target      apijson.Field
	TargetURI   apijson.Field
	Problem     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PromptReferenceAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PromptReferenceAttachmentJSON) RawJSON() string {
	return r.raw
}

type V2PromptReferenceAttachmentKind string

const (
	V2PromptReferenceAttachmentKindLocal   V2PromptReferenceAttachmentKind = "local"
	V2PromptReferenceAttachmentKindGit     V2PromptReferenceAttachmentKind = "git"
	V2PromptReferenceAttachmentKindInvalid V2PromptReferenceAttachmentKind = "invalid"
)

func (r V2PromptReferenceAttachmentKind) IsKnown() bool {
	switch r {
	case V2PromptReferenceAttachmentKindLocal, V2PromptReferenceAttachmentKindGit, V2PromptReferenceAttachmentKindInvalid:
		return true
	}
	return false
}

type V2PromptSource struct {
	Start float64           `json:"start,required"`
	End   float64           `json:"end,required"`
	Text  string            `json:"text,required"`
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
	Type string                                     `json:"type,required"`
	Text string                                     `json:"text,required"`
	JSON v2SessionMessageAssistantTextContentJSON   `json:"-"`
}

type v2SessionMessageAssistantTextContentJSON struct {
	Type        apijson.Field
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

type V2SessionMessageAssistantReasoningContent struct {
	Type string                                          `json:"type,required"`
	ID   string                                          `json:"id,required"`
	Text string                                          `json:"text,required"`
	JSON v2SessionMessageAssistantReasoningContentJSON   `json:"-"`
}

type v2SessionMessageAssistantReasoningContentJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionMessageAssistantReasoningContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionMessageAssistantReasoningContentJSON) RawJSON() string {
	return r.raw
}

type V2SessionMessageAssistantToolContent struct {
	Type     string                                     `json:"type,required"`
	ID       string                                     `json:"id,required"`
	Name     string                                     `json:"name,required"`
	Provider V2SessionMessageToolProvider               `json:"provider"`
	// This field can have the runtime type of
	// [V2SessionMessageToolStatePending], [V2SessionMessageToolStateRunning],
	// [V2SessionMessageToolStateCompleted], [V2SessionMessageToolStateError].
	State    V2SessionMessageToolState                  `json:"state,required"`
	Time     V2SessionMessageToolTime                   `json:"time,required"`
	JSON     v2SessionMessageAssistantToolContentJSON   `json:"-"`
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

// ===== Tool Content Types (shared) =====

type ToolTextContent struct {
	Type string              `json:"type,required"`
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

type ToolFileContent struct {
	Type string              `json:"type,required"`
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
	Input      map[string]interface{}                 `json:"input,required"`
	Structured map[string]interface{}                 `json:"structured,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content interface{}                               `json:"content,required"`
	JSON    v2SessionMessageToolStateRunningJSON      `json:"-"`
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
	Status      V2SessionMessageToolStateCompletedStatus `json:"status,required"`
	Input       map[string]interface{}                   `json:"input,required"`
	Structured  map[string]interface{}                   `json:"structured,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content     interface{}                              `json:"content,required"`
	Attachments []V2PromptFileAttachment                 `json:"attachments"`
	JSON        v2SessionMessageToolStateCompletedJSON   `json:"-"`
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
	Input      map[string]interface{}               `json:"input,required"`
	Structured map[string]interface{}               `json:"structured,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content interface{}                              `json:"content,required"`
	Error   SessionErrorUnknown                      `json:"error,required"`
	JSON    v2SessionMessageToolStateErrorJSON        `json:"-"`
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

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2SessionMessageToolStateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2SessionMessageToolStatePending{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2SessionMessageToolStateRunning{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2SessionMessageToolStateCompleted{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2SessionMessageToolStateError{}),
		},
	)
}

// ===== Param Types =====

type V2SessionListParams struct {
	Directory param.Field[string]  `query:"directory"`
	Workspace param.Field[string]  `query:"workspace"`
	Limit     param.Field[float64] `query:"limit"`
	Order     param.Field[string]  `query:"order"`
	Path      param.Field[string]  `query:"path"`
	Roots     param.Field[bool]    `query:"roots"`
	Start     param.Field[float64] `query:"start"`
	Search    param.Field[string]  `query:"search"`
	Cursor    param.Field[string]  `query:"cursor"`
}

func (r V2SessionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2SessionPromptParams struct {
	Directory param.Field[string]           `query:"directory"`
	Workspace param.Field[string]           `query:"workspace"`
	Body      V2SessionPromptParamsBody     `json:"body,required"`
}

func (r V2SessionPromptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

func (r V2SessionPromptParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2SessionPromptParamsBody struct {
	Prompt   param.Field[V2Prompt]   `json:"prompt,required"`
	Delivery param.Field[SessionDelivery] `json:"delivery"`
}

type SessionDelivery string

const (
	SessionDeliveryImmediate SessionDelivery = "immediate"
	SessionDeliveryDeferred  SessionDelivery = "deferred"
)

func (r SessionDelivery) IsKnown() bool {
	switch r {
	case SessionDeliveryImmediate, SessionDeliveryDeferred:
		return true
	}
	return false
}

func (r V2SessionPromptParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2Prompt struct {
	Text       param.Field[string]                            `json:"text,required"`
	Files      param.Field[[]V2PromptFileAttachmentParam]     `json:"files"`
	Agents     param.Field[[]V2PromptAgentAttachmentParam]    `json:"agents"`
	References param.Field[[]V2PromptReferenceAttachmentParam] `json:"references"`
}

func (r V2Prompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2PromptFileAttachmentParam struct {
	URI         param.Field[string]           `json:"uri,required"`
	Mime        param.Field[string]           `json:"mime,required"`
	Name        param.Field[string]           `json:"name"`
	Description param.Field[string]           `json:"description"`
	Source      param.Field[V2PromptSourceParam] `json:"source"`
}

func (r V2PromptFileAttachmentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2PromptAgentAttachmentParam struct {
	Name   param.Field[string]            `json:"name,required"`
	Source param.Field[V2PromptSourceParam] `json:"source"`
}

func (r V2PromptAgentAttachmentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2PromptReferenceAttachmentParam struct {
	Name       param.Field[string]            `json:"name,required"`
	Kind       param.Field[string]            `json:"kind,required"`
	URI        param.Field[string]            `json:"uri"`
	Repository param.Field[string]            `json:"repository"`
	Branch     param.Field[string]            `json:"branch"`
	Target     param.Field[string]            `json:"target"`
	TargetURI  param.Field[string]            `json:"targetUri"`
	Problem    param.Field[string]            `json:"problem"`
	Source     param.Field[V2PromptSourceParam] `json:"source"`
}

func (r V2PromptReferenceAttachmentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2PromptSourceParam struct {
	Start param.Field[float64] `json:"start,required"`
	End   param.Field[float64] `json:"end,required"`
	Text  param.Field[string]  `json:"text,required"`
}

func (r V2PromptSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2SessionCompactParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r V2SessionCompactParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2SessionWaitParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r V2SessionWaitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2SessionContextParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r V2SessionContextParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2SessionMessagesParams struct {
	Directory param.Field[string]  `query:"directory"`
	Workspace param.Field[string]  `query:"workspace"`
	Limit     param.Field[float64] `query:"limit"`
	Order     param.Field[string]  `query:"order"`
	Cursor    param.Field[string]  `query:"cursor"`
}

func (r V2SessionMessagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
