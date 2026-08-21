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

	"github.com/tidwall/gjson"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/pagination"
	"github.com/sst/opencode-sdk-go/shared"
)

// SessionService contains methods and other services that help with interacting
// with the opencode API. This includes session CRUD operations, message management
// (prompt, messages, commands, shell).
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionService] method instead.
type SessionService struct {
	Options []option.RequestOption
}

// NewSessionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSessionService(opts ...option.RequestOption) (r *SessionService) {
	r = &SessionService{}
	r.Options = opts
	return
}

// Create a new session
func (r *SessionService) New(ctx context.Context, params SessionNewParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update session properties
func (r *SessionService) Update(ctx context.Context, id string, params SessionUpdateParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all sessions
//
// This endpoint is not paginated and so has no auto-paging variant. The server
// returns the newest Limit sessions (100 when Limit is unset) ordered by their
// updated timestamp, and sends neither a body cursor nor an X-Next-Cursor
// header. Start is a lower bound on the updated timestamp rather than a cursor,
// so it cannot be used to walk backwards into older sessions.
//
// To reach beyond the first window, raise Limit and re-request, which is what
// the opencode web client does; a response holding exactly Limit sessions means
// more may exist.
//
// When real pagination is required, use [ExperimentalSessionService.List]
// instead. Given the same Directory it returns the same sessions in the same
// order, but backed by a cursor, and its GlobalSession carries every field of
// Session plus more. [ExperimentalSessionService.ListAutoPaging] then walks the
// whole list for you.
func (r *SessionService) List(ctx context.Context, query SessionListParams, opts ...option.RequestOption) (res *[]Session, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a session and all its data
func (r *SessionService) Delete(ctx context.Context, id string, body SessionDeleteParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Abort a session
func (r *SessionService) Abort(ctx context.Context, id string, body SessionAbortParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/abort", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get session todos
func (r *SessionService) Todo(ctx context.Context, id string, query SessionTodoParams, opts ...option.RequestOption) (res *[]Todo, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/todo", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a session's children
func (r *SessionService) Children(ctx context.Context, id string, query SessionChildrenParams, opts ...option.RequestOption) (res *[]Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/children", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Send a new command to a session
func (r *SessionService) Command(ctx context.Context, id string, params SessionCommandParams, opts ...option.RequestOption) (res *SessionCommandResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/command", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get session
func (r *SessionService) Get(ctx context.Context, id string, query SessionGetParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Analyze the app and create an AGENTS.md file
func (r *SessionService) Init(ctx context.Context, id string, params SessionInitParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/init", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a message from a session
func (r *SessionService) Message(ctx context.Context, id string, messageID string, query SessionMessageParams, opts ...option.RequestOption) (res *SessionMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message/%s", id, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List messages for a session
//
// The response body is a bare array of messages, and the server advertises the
// next page through the X-Next-Cursor response header, which it sends only while
// further messages remain. Read the items from Data, and call GetNextPage to
// replay the header value as the `before` query parameter.
//
// Paging only happens when Limit is set: with no Limit the server returns the
// whole conversation in a single response and never advertises a cursor. Each
// page is ordered oldest first, and every following page holds strictly older
// messages than the one before it.
func (r *SessionService) Messages(ctx context.Context, id string, query SessionMessagesParams, opts ...option.RequestOption) (res *pagination.HeaderBeforePage[SessionMessagesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat([]option.RequestOption{option.WithResponseInto(&raw)}, r.Options, opts)
	if id == "" {
		return nil, errors.New("missing required id parameter")
	}
	path := fmt.Sprintf("session/%s/message", id)
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

// List messages for a session
//
// MessagesAutoPaging walks every page for you, replaying the X-Next-Cursor
// response header as the `before` query parameter until the server stops sending
// it. Set Limit on the params to control the page size; without it the first
// response already carries the whole conversation.
func (r *SessionService) MessagesAutoPaging(ctx context.Context, id string, query SessionMessagesParams, opts ...option.RequestOption) *pagination.HeaderBeforePageAutoPager[SessionMessagesResponse] {
	return pagination.NewHeaderBeforePageAutoPager(r.Messages(ctx, id, query, opts...))
}

// Create and send a new message to a session
func (r *SessionService) Prompt(ctx context.Context, id string, params SessionPromptParams, opts ...option.RequestOption) (res *SessionPromptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Revert a message
func (r *SessionService) Revert(ctx context.Context, id string, params SessionRevertParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/revert", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Share a session
func (r *SessionService) Share(ctx context.Context, id string, body SessionShareParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/share", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run a shell command
func (r *SessionService) Shell(ctx context.Context, id string, params SessionShellParams, opts ...option.RequestOption) (res *SessionMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/shell", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Summarize the session
func (r *SessionService) Summarize(ctx context.Context, id string, params SessionSummarizeParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/summarize", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Restore all reverted messages
func (r *SessionService) Unrevert(ctx context.Context, id string, body SessionUnrevertParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/unrevert", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unshare the session
func (r *SessionService) Unshare(ctx context.Context, id string, body SessionUnshareParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/share", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Get session status
//
// Retrieve the current status of all sessions, including active, idle, and completed states.
func (r *SessionService) Status(ctx context.Context, query SessionStatusParams, opts ...option.RequestOption) (res *SessionStatusMap, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "session/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fork a session
func (r *SessionService) Fork(ctx context.Context, id string, params SessionForkParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/fork", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get session diff
func (r *SessionService) Diff(ctx context.Context, id string, query SessionDiffParams, opts ...option.RequestOption) (res *[]SnapshotFileDiff, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/diff", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a message from a session
func (r *SessionService) DeleteMessage(ctx context.Context, id string, messageID string, body SessionDeleteMessageParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message/%s", id, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Send a new message to a session asynchronously
func (r *SessionService) PromptAsync(ctx context.Context, id string, params SessionPromptParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("session/%s/prompt_async", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type Todo struct {
	Content  string   `json:"content,required"`
	Status   string   `json:"status,required"`
	Priority string   `json:"priority,required"`
	JSON     todoJSON `json:"-"`
}

// todoJSON contains the JSON metadata for the struct [Todo]
type todoJSON struct {
	Content     apijson.Field
	Status      apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Todo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r todoJSON) RawJSON() string {
	return r.raw
}

type PermissionAction string

const (
	PermissionActionAllow PermissionAction = "allow"
	PermissionActionDeny  PermissionAction = "deny"
	PermissionActionAsk   PermissionAction = "ask"
)

func (r PermissionAction) IsKnown() bool {
	switch r {
	case PermissionActionAllow, PermissionActionDeny, PermissionActionAsk:
		return true
	}
	return false
}

// PermissionRule is used for request parameters.
type PermissionRule struct {
	Permission param.Field[string]           `json:"permission,required"`
	Pattern    param.Field[string]           `json:"pattern,required"`
	Action     param.Field[PermissionAction] `json:"action,required"`
}

func (r PermissionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PermissionRuleResponse is used for response deserialization.
type PermissionRuleResponse struct {
	Permission string                     `json:"permission,required"`
	Pattern    string                     `json:"pattern,required"`
	Action     PermissionAction           `json:"action,required"`
	JSON       permissionRuleResponseJSON `json:"-"`
}

// permissionRuleResponseJSON contains the JSON metadata for the struct [PermissionRuleResponse]
type permissionRuleResponseJSON struct {
	Permission  apijson.Field
	Pattern     apijson.Field
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionRuleResponseJSON) RawJSON() string {
	return r.raw
}

// PermissionRuleset is used for response deserialization.
type PermissionRuleset []PermissionRuleResponse

type AgentPart struct {
	ID        string          `json:"id,required"`
	MessageID string          `json:"messageID,required"`
	Name      string          `json:"name,required"`
	SessionID string          `json:"sessionID,required"`
	Type      AgentPartType   `json:"type,required"`
	Source    AgentPartSource `json:"source"`
	JSON      agentPartJSON   `json:"-"`
}

// agentPartJSON contains the JSON metadata for the struct [AgentPart]
type agentPartJSON struct {
	ID          apijson.Field
	MessageID   apijson.Field
	Name        apijson.Field
	SessionID   apijson.Field
	Type        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentPartJSON) RawJSON() string {
	return r.raw
}

func (r AgentPart) implementsPart() {}

type AgentPartType string

const (
	AgentPartTypeAgent AgentPartType = "agent"
)

func (r AgentPartType) IsKnown() bool {
	switch r {
	case AgentPartTypeAgent:
		return true
	}
	return false
}

type AgentPartSource struct {
	End   int64               `json:"end,required"`
	Start int64               `json:"start,required"`
	Value string              `json:"value,required"`
	JSON  agentPartSourceJSON `json:"-"`
}

// agentPartSourceJSON contains the JSON metadata for the struct [AgentPartSource]
type agentPartSourceJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentPartSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentPartSourceJSON) RawJSON() string {
	return r.raw
}

type AgentPartInputParam struct {
	Name   param.Field[string]                    `json:"name,required"`
	Type   param.Field[AgentPartInputType]        `json:"type,required"`
	ID     param.Field[string]                    `json:"id"`
	Source param.Field[AgentPartInputSourceParam] `json:"source"`
}

func (r AgentPartInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AgentPartInputParam) implementsSessionPromptParamsPartUnion() {}

type AgentPartInputType string

const (
	AgentPartInputTypeAgent AgentPartInputType = "agent"
)

func (r AgentPartInputType) IsKnown() bool {
	switch r {
	case AgentPartInputTypeAgent:
		return true
	}
	return false
}

type SubtaskPartInputParam struct {
	ID          param.Field[string]                `json:"id"`
	Type        param.Field[SubtaskPartInputType]  `json:"type,required"`
	Prompt      param.Field[string]                `json:"prompt,required"`
	Description param.Field[string]                `json:"description,required"`
	Agent       param.Field[string]                `json:"agent,required"`
	Model       param.Field[SubtaskPartModelParam] `json:"model"`
	Command     param.Field[string]                `json:"command"`
}

func (r SubtaskPartInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubtaskPartInputParam) implementsSessionPromptParamsPartUnion() {}

// SubtaskPartModelParam is used for request parameters.
type SubtaskPartModelParam struct {
	ProviderID param.Field[string] `json:"providerID,required"`
	ModelID    param.Field[string] `json:"modelID,required"`
}

func (r SubtaskPartModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubtaskPartInputType string

const (
	SubtaskPartInputTypeSubtask SubtaskPartInputType = "subtask"
)

func (r SubtaskPartInputType) IsKnown() bool {
	switch r {
	case SubtaskPartInputTypeSubtask:
		return true
	}
	return false
}

type AgentPartInputSourceParam struct {
	End   param.Field[int64]  `json:"end,required"`
	Start param.Field[int64]  `json:"start,required"`
	Value param.Field[string] `json:"value,required"`
}

func (r AgentPartInputSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OutputFormatTextType string

const (
	OutputFormatTextTypeText OutputFormatTextType = "text"
)

func (r OutputFormatTextType) IsKnown() bool {
	switch r {
	case OutputFormatTextTypeText:
		return true
	}
	return false
}

type OutputFormatJsonSchemaType string

const (
	OutputFormatJsonSchemaTypeJsonSchema OutputFormatJsonSchemaType = "json_schema"
)

func (r OutputFormatJsonSchemaType) IsKnown() bool {
	switch r {
	case OutputFormatJsonSchemaTypeJsonSchema:
		return true
	}
	return false
}

// OutputFormat is the OpenAPI `OutputFormat` anyOf union, carrying the
// flattened superset of every variant's fields. `Type` discriminates which
// variant the payload actually is; use [OutputFormat.AsUnion] to recover the
// concrete variant.
type OutputFormat struct {
	Type OutputFormatType `json:"type,required"`
	// This field is only present on the [OutputFormatJsonSchema] variant.
	Schema map[string]any `json:"schema,omitzero"`
	// This field is only present on the [OutputFormatJsonSchema] variant.
	RetryCount int64            `json:"retryCount,omitzero"`
	JSON       outputFormatJSON `json:"-"`
	union      OutputFormatUnion
}

// outputFormatJSON contains the JSON metadata for the struct [OutputFormat]
type outputFormatJSON struct {
	Type        apijson.Field
	Schema      apijson.Field
	RetryCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r outputFormatJSON) RawJSON() string {
	return r.raw
}

func (r *OutputFormat) UnmarshalJSON(data []byte) (err error) {
	*r = OutputFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns an [OutputFormatUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [OutputFormatText],
// [OutputFormatJsonSchema].
func (r OutputFormat) AsUnion() OutputFormatUnion {
	return r.union
}

// OutputFormatType is the discriminator shared by every [OutputFormat] variant.
type OutputFormatType string

const (
	OutputFormatTypeText       OutputFormatType = "text"
	OutputFormatTypeJsonSchema OutputFormatType = "json_schema"
)

func (r OutputFormatType) IsKnown() bool {
	switch r {
	case OutputFormatTypeText, OutputFormatTypeJsonSchema:
		return true
	}
	return false
}

// OutputFormatUnion is the OpenAPI `OutputFormat` anyOf union, used by
// [UserMessage.Format].
//
// Satisfied by [OutputFormatText], [OutputFormatJsonSchema].
type OutputFormatUnion interface {
	implementsOutputFormatUnion()
}

// OutputFormatText is the Response-side representation of the OpenAPI
// `OutputFormatText` schema (`type: "text"`).
type OutputFormatText struct {
	Type OutputFormatTextType `json:"type,required"`
	JSON outputFormatTextJSON `json:"-"`
}

// outputFormatTextJSON contains the JSON metadata for the struct
// [OutputFormatText]
type outputFormatTextJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputFormatText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outputFormatTextJSON) RawJSON() string {
	return r.raw
}

func (r OutputFormatText) implementsOutputFormatUnion() {}

// OutputFormatJsonSchema is the Response-side representation of the OpenAPI
// `OutputFormatJsonSchema` schema (`type: "json_schema"`).
type OutputFormatJsonSchema struct {
	Schema     map[string]any             `json:"schema,required"`
	Type       OutputFormatJsonSchemaType `json:"type,required"`
	RetryCount int64                      `json:"retryCount"`
	JSON       outputFormatJsonSchemaJSON `json:"-"`
}

// outputFormatJsonSchemaJSON contains the JSON metadata for the struct
// [OutputFormatJsonSchema]
type outputFormatJsonSchemaJSON struct {
	Schema      apijson.Field
	Type        apijson.Field
	RetryCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutputFormatJsonSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outputFormatJsonSchemaJSON) RawJSON() string {
	return r.raw
}

func (r OutputFormatJsonSchema) implementsOutputFormatUnion() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[OutputFormatUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[OutputFormatText](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[OutputFormatJsonSchema](),
		},
	)
}

type AssistantMessage struct {
	ID         string                 `json:"id,required"`
	Agent      string                 `json:"agent,required"`
	Cost       float64                `json:"cost,required"`
	Mode       string                 `json:"mode,required"`
	ModelID    string                 `json:"modelID,required"`
	ParentID   string                 `json:"parentID,required"`
	Path       AssistantMessagePath   `json:"path,required"`
	ProviderID string                 `json:"providerID,required"`
	Role       AssistantMessageRole   `json:"role,required"`
	SessionID  string                 `json:"sessionID,required"`
	Time       AssistantMessageTime   `json:"time,required"`
	Tokens     AssistantMessageTokens `json:"tokens,required"`
	// This field is an untyped arbitrary value. The OpenAPI schema declares it as
	// an empty schema (`{}`), meaning it may hold any JSON value. Use a
	// type-switch or json.Unmarshal to inspect the runtime value.
	Structured any                   `json:"structured,omitzero"`
	Variant    string                `json:"variant,omitzero"`
	Finish     string                `json:"finish,omitzero"`
	Error      AssistantMessageError `json:"error"`
	Summary    bool                  `json:"summary"`
	JSON       assistantMessageJSON  `json:"-"`
}

// assistantMessageJSON contains the JSON metadata for the struct
// [AssistantMessage]
type assistantMessageJSON struct {
	ID          apijson.Field
	Agent       apijson.Field
	Cost        apijson.Field
	Mode        apijson.Field
	ModelID     apijson.Field
	ParentID    apijson.Field
	Path        apijson.Field
	ProviderID  apijson.Field
	Role        apijson.Field
	SessionID   apijson.Field
	Time        apijson.Field
	Tokens      apijson.Field
	Structured  apijson.Field
	Variant     apijson.Field
	Finish      apijson.Field
	Error       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageJSON) RawJSON() string {
	return r.raw
}

func (r AssistantMessage) implementsMessage() {}

type AssistantMessagePath struct {
	Cwd  string                   `json:"cwd,required"`
	Root string                   `json:"root,required"`
	JSON assistantMessagePathJSON `json:"-"`
}

// assistantMessagePathJSON contains the JSON metadata for the struct
// [AssistantMessagePath]
type assistantMessagePathJSON struct {
	Cwd         apijson.Field
	Root        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessagePath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessagePathJSON) RawJSON() string {
	return r.raw
}

type AssistantMessageRole string

const (
	AssistantMessageRoleAssistant AssistantMessageRole = "assistant"
)

func (r AssistantMessageRole) IsKnown() bool {
	switch r {
	case AssistantMessageRoleAssistant:
		return true
	}
	return false
}

type AssistantMessageTime struct {
	Created   int64                    `json:"created,required"`
	Completed int64                    `json:"completed"`
	JSON      assistantMessageTimeJSON `json:"-"`
}

// assistantMessageTimeJSON contains the JSON metadata for the struct
// [AssistantMessageTime]
type assistantMessageTimeJSON struct {
	Created     apijson.Field
	Completed   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessageTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageTimeJSON) RawJSON() string {
	return r.raw
}

type AssistantMessageTokens struct {
	Cache     AssistantMessageTokensCache `json:"cache,required"`
	Input     int64                       `json:"input,required"`
	Output    int64                       `json:"output,required"`
	Reasoning int64                       `json:"reasoning,required"`
	Total     int64                       `json:"total"`
	JSON      assistantMessageTokensJSON  `json:"-"`
}

// assistantMessageTokensJSON contains the JSON metadata for the struct
// [AssistantMessageTokens]
type assistantMessageTokensJSON struct {
	Cache       apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessageTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageTokensJSON) RawJSON() string {
	return r.raw
}

type AssistantMessageTokensCache struct {
	Read  int64                           `json:"read,required"`
	Write int64                           `json:"write,required"`
	JSON  assistantMessageTokensCacheJSON `json:"-"`
}

// assistantMessageTokensCacheJSON contains the JSON metadata for the struct
// [AssistantMessageTokensCache]
type assistantMessageTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessageTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageTokensCacheJSON) RawJSON() string {
	return r.raw
}

type AssistantMessageError struct {
	// This field can have the runtime type of [shared.ProviderAuthErrorData],
	// [shared.UnknownErrorData], [any], [shared.MessageAbortedErrorData],
	// [shared.StructuredOutputErrorData], [shared.ContextOverflowErrorData],
	// [shared.ContentFilterErrorData], [shared.APIErrorData].
	Data  any                       `json:"data,required"`
	Name  AssistantMessageErrorName `json:"name,required"`
	JSON  assistantMessageErrorJSON `json:"-"`
	union AssistantMessageErrorUnion
}

// assistantMessageErrorJSON contains the JSON metadata for the struct
// [AssistantMessageError]
type assistantMessageErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r assistantMessageErrorJSON) RawJSON() string {
	return r.raw
}

func (r *AssistantMessageError) UnmarshalJSON(data []byte) (err error) {
	*r = AssistantMessageError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AssistantMessageErrorUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [shared.ProviderAuthError],
// [shared.UnknownError], [shared.MessageOutputLengthError],
// [shared.MessageAbortedError], [shared.StructuredOutputError],
// [shared.ContentFilterError], [shared.ContextOverflowError],
// [shared.APIError].
func (r AssistantMessageError) AsUnion() AssistantMessageErrorUnion {
	return r.union
}

// Union satisfied by [shared.ProviderAuthError], [shared.UnknownError],
// [shared.MessageOutputLengthError], [shared.MessageAbortedError],
// [shared.StructuredOutputError], [shared.ContentFilterError],
// [shared.ContextOverflowError] or [shared.APIError].
type AssistantMessageErrorUnion interface {
	ImplementsAssistantMessageError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[AssistantMessageErrorUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[shared.ProviderAuthError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[shared.UnknownError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[shared.MessageOutputLengthError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[shared.MessageAbortedError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[shared.StructuredOutputError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[shared.ContentFilterError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[shared.ContextOverflowError](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[shared.APIError](),
		},
	)
}

type AssistantMessageErrorName string

const (
	AssistantMessageErrorNameProviderAuthError        AssistantMessageErrorName = "ProviderAuthError"
	AssistantMessageErrorNameUnknownError             AssistantMessageErrorName = "UnknownError"
	AssistantMessageErrorNameMessageOutputLengthError AssistantMessageErrorName = "MessageOutputLengthError"
	AssistantMessageErrorNameMessageAbortedError      AssistantMessageErrorName = "MessageAbortedError"
	AssistantMessageErrorNameStructuredOutputError    AssistantMessageErrorName = "StructuredOutputError"
	AssistantMessageErrorNameContextOverflowError     AssistantMessageErrorName = "ContextOverflowError"
	AssistantMessageErrorNameAPIError                 AssistantMessageErrorName = "APIError"
	AssistantMessageErrorNameContentFilterError       AssistantMessageErrorName = "ContentFilterError"
)

func (r AssistantMessageErrorName) IsKnown() bool {
	switch r {
	case AssistantMessageErrorNameProviderAuthError, AssistantMessageErrorNameUnknownError, AssistantMessageErrorNameMessageOutputLengthError, AssistantMessageErrorNameMessageAbortedError, AssistantMessageErrorNameStructuredOutputError, AssistantMessageErrorNameContextOverflowError, AssistantMessageErrorNameAPIError, AssistantMessageErrorNameContentFilterError:
		return true
	}
	return false
}

type FilePart struct {
	ID        string         `json:"id,required"`
	MessageID string         `json:"messageID,required"`
	Mime      string         `json:"mime,required"`
	SessionID string         `json:"sessionID,required"`
	Type      FilePartType   `json:"type,required"`
	URL       string         `json:"url,required"`
	Filename  string         `json:"filename"`
	Source    FilePartSource `json:"source"`
	JSON      filePartJSON   `json:"-"`
}

// filePartJSON contains the JSON metadata for the struct [FilePart]
type filePartJSON struct {
	ID          apijson.Field
	MessageID   apijson.Field
	Mime        apijson.Field
	SessionID   apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	Filename    apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilePart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filePartJSON) RawJSON() string {
	return r.raw
}

func (r FilePart) implementsPart() {}

type FilePartType string

const (
	FilePartTypeFile FilePartType = "file"
)

func (r FilePartType) IsKnown() bool {
	switch r {
	case FilePartTypeFile:
		return true
	}
	return false
}

type FilePartInputParam struct {
	Mime     param.Field[string]                   `json:"mime,required"`
	Type     param.Field[FilePartInputType]        `json:"type,required"`
	URL      param.Field[string]                   `json:"url,required"`
	ID       param.Field[string]                   `json:"id"`
	Filename param.Field[string]                   `json:"filename"`
	Source   param.Field[FilePartSourceUnionParam] `json:"source"`
}

func (r FilePartInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FilePartInputParam) implementsSessionPromptParamsPartUnion() {}

type FilePartInputType string

const (
	FilePartInputTypeFile FilePartInputType = "file"
)

func (r FilePartInputType) IsKnown() bool {
	switch r {
	case FilePartInputTypeFile:
		return true
	}
	return false
}

type FilePartSource struct {
	Path string             `json:"path,required"`
	Text FilePartSourceText `json:"text,required"`
	Type FilePartSourceType `json:"type,required"`
	Kind int64              `json:"kind"`
	Name string             `json:"name"`
	// This field can have the runtime type of [Range].
	Range any                `json:"range"`
	JSON  filePartSourceJSON `json:"-"`
	union FilePartSourceUnion
}

// filePartSourceJSON contains the JSON metadata for the struct [FilePartSource]
type filePartSourceJSON struct {
	Path        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Range       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r filePartSourceJSON) RawJSON() string {
	return r.raw
}

func (r *FilePartSource) UnmarshalJSON(data []byte) (err error) {
	*r = FilePartSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FilePartSourceUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [FileSource], [SymbolSource],
// [ResourceSource].
func (r FilePartSource) AsUnion() FilePartSourceUnion {
	return r.union
}

// Union satisfied by [FileSource], [SymbolSource] or [ResourceSource].
type FilePartSourceUnion interface {
	implementsFilePartSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[FilePartSourceUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[FileSource](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SymbolSource](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ResourceSource](),
		},
	)
}

type FilePartSourceType string

const (
	FilePartSourceTypeFile     FilePartSourceType = "file"
	FilePartSourceTypeSymbol   FilePartSourceType = "symbol"
	FilePartSourceTypeResource FilePartSourceType = "resource"
)

func (r FilePartSourceType) IsKnown() bool {
	switch r {
	case FilePartSourceTypeFile, FilePartSourceTypeSymbol, FilePartSourceTypeResource:
		return true
	}
	return false
}

type FilePartSourceParam struct {
	Path  param.Field[string]                  `json:"path,required"`
	Text  param.Field[FilePartSourceTextParam] `json:"text,required"`
	Type  param.Field[FilePartSourceType]      `json:"type,required"`
	Kind  param.Field[int64]                   `json:"kind"`
	Name  param.Field[string]                  `json:"name"`
	Range param.Field[any]                     `json:"range"`
}

func (r FilePartSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FilePartSourceParam) implementsFilePartSourceUnionParam() {}

// Satisfied by [FileSourceParam], [SymbolSourceParam], [FilePartSourceParam].
type FilePartSourceUnionParam interface {
	implementsFilePartSourceUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[FilePartSourceUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[FilePartSourceParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[FileSourceParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SymbolSourceParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ResourceSourceParam](),
		},
	)
}

type FilePartSourceText struct {
	End   int64                  `json:"end,required"`
	Start int64                  `json:"start,required"`
	Value string                 `json:"value,required"`
	JSON  filePartSourceTextJSON `json:"-"`
}

// filePartSourceTextJSON contains the JSON metadata for the struct
// [FilePartSourceText]
type filePartSourceTextJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilePartSourceText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filePartSourceTextJSON) RawJSON() string {
	return r.raw
}

type FilePartSourceTextParam struct {
	End   param.Field[int64]  `json:"end,required"`
	Start param.Field[int64]  `json:"start,required"`
	Value param.Field[string] `json:"value,required"`
}

func (r FilePartSourceTextParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileSource struct {
	Path string             `json:"path,required"`
	Text FilePartSourceText `json:"text,required"`
	Type FileSourceType     `json:"type,required"`
	JSON fileSourceJSON     `json:"-"`
}

// fileSourceJSON contains the JSON metadata for the struct [FileSource]
type fileSourceJSON struct {
	Path        apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileSourceJSON) RawJSON() string {
	return r.raw
}

func (r FileSource) implementsFilePartSource() {}

type FileSourceType string

const (
	FileSourceTypeFile FileSourceType = "file"
)

func (r FileSourceType) IsKnown() bool {
	switch r {
	case FileSourceTypeFile:
		return true
	}
	return false
}

type FileSourceParam struct {
	Path param.Field[string]                  `json:"path,required"`
	Text param.Field[FilePartSourceTextParam] `json:"text,required"`
	Type param.Field[FileSourceType]          `json:"type,required"`
}

func (r FileSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FileSourceParam) implementsFilePartSourceUnionParam() {}

type Message struct {
	ID        string      `json:"id,required"`
	Agent     string      `json:"agent,required"`
	Role      MessageRole `json:"role,required"`
	SessionID string      `json:"sessionID,required"`
	// This field can have the runtime type of [UserMessageTime],
	// [AssistantMessageTime].
	Time any     `json:"time,required"`
	Cost float64 `json:"cost"`
	// This field can have the runtime type of [AssistantMessageError].
	Error    any    `json:"error"`
	Mode     string `json:"mode"`
	ModelID  string `json:"modelID"`
	ParentID string `json:"parentID"`
	// This field can have the runtime type of [AssistantMessagePath].
	Path       any    `json:"path"`
	ProviderID string `json:"providerID"`
	// This field can have the runtime type of [UserMessageSummary], [bool].
	Summary any    `json:"summary"`
	Finish  string `json:"finish"`
	System  string `json:"system"`
	// This field can have the runtime type of [AssistantMessageTokens].
	Tokens any         `json:"tokens"`
	JSON   messageJSON `json:"-"`
	union  MessageUnion
}

// messageJSON contains the JSON metadata for the struct [Message]
type messageJSON struct {
	ID          apijson.Field
	Agent       apijson.Field
	Role        apijson.Field
	SessionID   apijson.Field
	Time        apijson.Field
	Cost        apijson.Field
	Error       apijson.Field
	Mode        apijson.Field
	ModelID     apijson.Field
	ParentID    apijson.Field
	Path        apijson.Field
	ProviderID  apijson.Field
	Summary     apijson.Field
	Finish      apijson.Field
	System      apijson.Field
	Tokens      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r messageJSON) RawJSON() string {
	return r.raw
}

func (r *Message) UnmarshalJSON(data []byte) (err error) {
	*r = Message{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MessageUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [UserMessage], [AssistantMessage].
func (r Message) AsUnion() MessageUnion {
	return r.union
}

// Union satisfied by [UserMessage] or [AssistantMessage].
type MessageUnion interface {
	implementsMessage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[MessageUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[UserMessage](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[AssistantMessage](),
		},
	)
}

type MessageRole string

const (
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
)

func (r MessageRole) IsKnown() bool {
	switch r {
	case MessageRoleUser, MessageRoleAssistant:
		return true
	}
	return false
}

type Part struct {
	ID          string             `json:"id,required"`
	MessageID   string             `json:"messageID,required"`
	SessionID   string             `json:"sessionID,required"`
	Type        PartType           `json:"type,required"`
	Agent       string             `json:"agent"`
	Auto        bool               `json:"auto"`
	Attempt     int64              `json:"attempt"`
	CallID      string             `json:"callID"`
	Command     string             `json:"command"`
	Cost        float64            `json:"cost"`
	Description string             `json:"description"`
	Error       PartRetryPartError `json:"error"`
	Filename    string             `json:"filename"`
	Files       []string           `json:"files"`
	Hash        string             `json:"hash"`
	Ignored     bool               `json:"ignored"`
	// This field can have the runtime type of [map[string]any].
	Metadata map[string]any   `json:"metadata"`
	Mime     string           `json:"mime"`
	Model    SubtaskPartModel `json:"model"`
	Name     string           `json:"name"`
	Overflow bool             `json:"overflow"`
	Prompt   string           `json:"prompt"`
	Reason   string           `json:"reason"`
	Snapshot string           `json:"snapshot"`
	// This field can have the runtime type of [FilePartSource], [AgentPartSource].
	Source any `json:"source"`
	// This field can have the runtime type of [ToolPartState].
	State       any    `json:"state"`
	Synthetic   bool   `json:"synthetic"`
	TailStartID string `json:"tail_start_id"`
	Text        string `json:"text"`
	// This field can have the runtime type of [TextPartTime], [ReasoningPartTime],
	// [PartRetryPartTime].
	Time any `json:"time"`
	// This field can have the runtime type of [StepFinishPartTokens].
	Tokens any      `json:"tokens"`
	Tool   string   `json:"tool"`
	URL    string   `json:"url"`
	JSON   partJSON `json:"-"`
	union  PartUnion
}

// partJSON contains the JSON metadata for the struct [Part]
type partJSON struct {
	ID          apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Type        apijson.Field
	Agent       apijson.Field
	Auto        apijson.Field
	Attempt     apijson.Field
	CallID      apijson.Field
	Command     apijson.Field
	Cost        apijson.Field
	Description apijson.Field
	Error       apijson.Field
	Filename    apijson.Field
	Files       apijson.Field
	Hash        apijson.Field
	Ignored     apijson.Field
	Metadata    apijson.Field
	Mime        apijson.Field
	Model       apijson.Field
	Name        apijson.Field
	Overflow    apijson.Field
	Prompt      apijson.Field
	Reason      apijson.Field
	Snapshot    apijson.Field
	Source      apijson.Field
	State       apijson.Field
	Synthetic   apijson.Field
	TailStartID apijson.Field
	Text        apijson.Field
	Time        apijson.Field
	Tokens      apijson.Field
	Tool        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r partJSON) RawJSON() string {
	return r.raw
}

func (r *Part) UnmarshalJSON(data []byte) (err error) {
	*r = Part{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PartUnion] interface which you can cast to the specific types
// for more type safety.
//
// Possible runtime types of the union are [TextPart], [SubtaskPart], [ReasoningPart],
// [FilePart], [ToolPart], [StepStartPart], [StepFinishPart], [SnapshotPart],
// [PartPatchPart], [AgentPart], [PartRetryPart], [CompactionPart].
func (r Part) AsUnion() PartUnion {
	return r.union
}

// Union satisfied by [TextPart], [SubtaskPart], [ReasoningPart], [FilePart],
// [ToolPart], [StepStartPart], [StepFinishPart], [SnapshotPart], [PartPatchPart],
// [AgentPart], [PartRetryPart] or [CompactionPart].
type PartUnion interface {
	implementsPart()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[PartUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[TextPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SubtaskPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ReasoningPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[FilePart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ToolPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[StepStartPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[StepFinishPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SnapshotPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartPatchPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[AgentPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartRetryPart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[CompactionPart](),
		},
	)
}

type PartPatchPart struct {
	ID        string            `json:"id,required"`
	Files     []string          `json:"files,required"`
	Hash      string            `json:"hash,required"`
	MessageID string            `json:"messageID,required"`
	SessionID string            `json:"sessionID,required"`
	Type      PartPatchPartType `json:"type,required"`
	JSON      partPatchPartJSON `json:"-"`
}

// partPatchPartJSON contains the JSON metadata for the struct [PartPatchPart]
type partPatchPartJSON struct {
	ID          apijson.Field
	Files       apijson.Field
	Hash        apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PartPatchPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r partPatchPartJSON) RawJSON() string {
	return r.raw
}

func (r PartPatchPart) implementsPart() {}

type PartPatchPartType string

const (
	PartPatchPartTypePatch PartPatchPartType = "patch"
)

func (r PartPatchPartType) IsKnown() bool {
	switch r {
	case PartPatchPartTypePatch:
		return true
	}
	return false
}

type PartRetryPart struct {
	ID        string             `json:"id,required"`
	Attempt   int64              `json:"attempt,required"`
	Error     PartRetryPartError `json:"error,required"`
	MessageID string             `json:"messageID,required"`
	SessionID string             `json:"sessionID,required"`
	Time      PartRetryPartTime  `json:"time,required"`
	Type      PartRetryPartType  `json:"type,required"`
	JSON      partRetryPartJSON  `json:"-"`
}

// partRetryPartJSON contains the JSON metadata for the struct [PartRetryPart]
type partRetryPartJSON struct {
	ID          apijson.Field
	Attempt     apijson.Field
	Error       apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PartRetryPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r partRetryPartJSON) RawJSON() string {
	return r.raw
}

func (r PartRetryPart) implementsPart() {}

type PartRetryPartError struct {
	Data PartRetryPartErrorData `json:"data,required"`
	Name PartRetryPartErrorName `json:"name,required"`
	JSON partRetryPartErrorJSON `json:"-"`
}

// partRetryPartErrorJSON contains the JSON metadata for the struct
// [PartRetryPartError]
type partRetryPartErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PartRetryPartError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r partRetryPartErrorJSON) RawJSON() string {
	return r.raw
}

type PartRetryPartErrorData struct {
	IsRetryable     bool                       `json:"isRetryable,required"`
	Message         string                     `json:"message,required"`
	Metadata        map[string]string          `json:"metadata"`
	ResponseBody    string                     `json:"responseBody"`
	ResponseHeaders map[string]string          `json:"responseHeaders"`
	StatusCode      int64                      `json:"statusCode"`
	JSON            partRetryPartErrorDataJSON `json:"-"`
}

// partRetryPartErrorDataJSON contains the JSON metadata for the struct
// [PartRetryPartErrorData]
type partRetryPartErrorDataJSON struct {
	IsRetryable     apijson.Field
	Message         apijson.Field
	Metadata        apijson.Field
	ResponseBody    apijson.Field
	ResponseHeaders apijson.Field
	StatusCode      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PartRetryPartErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r partRetryPartErrorDataJSON) RawJSON() string {
	return r.raw
}

type PartRetryPartErrorName string

const (
	PartRetryPartErrorNameAPIError PartRetryPartErrorName = "APIError"
)

func (r PartRetryPartErrorName) IsKnown() bool {
	switch r {
	case PartRetryPartErrorNameAPIError:
		return true
	}
	return false
}

type PartRetryPartTime struct {
	Created int64                 `json:"created,required"`
	JSON    partRetryPartTimeJSON `json:"-"`
}

// partRetryPartTimeJSON contains the JSON metadata for the struct
// [PartRetryPartTime]
type partRetryPartTimeJSON struct {
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PartRetryPartTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r partRetryPartTimeJSON) RawJSON() string {
	return r.raw
}

type PartRetryPartType string

const (
	PartRetryPartTypeRetry PartRetryPartType = "retry"
)

func (r PartRetryPartType) IsKnown() bool {
	switch r {
	case PartRetryPartTypeRetry:
		return true
	}
	return false
}

type PartType string

const (
	PartTypeText       PartType = "text"
	PartTypeReasoning  PartType = "reasoning"
	PartTypeFile       PartType = "file"
	PartTypeTool       PartType = "tool"
	PartTypeStepStart  PartType = "step-start"
	PartTypeStepFinish PartType = "step-finish"
	PartTypeSnapshot   PartType = "snapshot"
	PartTypePatch      PartType = "patch"
	PartTypeAgent      PartType = "agent"
	PartTypeRetry      PartType = "retry"
	PartTypeSubtask    PartType = "subtask"
	PartTypeCompaction PartType = "compaction"
)

func (r PartType) IsKnown() bool {
	switch r {
	case PartTypeText, PartTypeReasoning, PartTypeFile, PartTypeTool, PartTypeStepStart, PartTypeStepFinish, PartTypeSnapshot, PartTypePatch, PartTypeAgent, PartTypeRetry, PartTypeSubtask, PartTypeCompaction:
		return true
	}
	return false
}

type ReasoningPart struct {
	ID        string            `json:"id,required"`
	MessageID string            `json:"messageID,required"`
	SessionID string            `json:"sessionID,required"`
	Text      string            `json:"text,required"`
	Time      ReasoningPartTime `json:"time,required"`
	Type      ReasoningPartType `json:"type,required"`
	Metadata  map[string]any    `json:"metadata"`
	JSON      reasoningPartJSON `json:"-"`
}

// reasoningPartJSON contains the JSON metadata for the struct [ReasoningPart]
type reasoningPartJSON struct {
	ID          apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReasoningPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reasoningPartJSON) RawJSON() string {
	return r.raw
}

func (r ReasoningPart) implementsPart() {}

type ReasoningPartTime struct {
	Start int64                 `json:"start,required"`
	End   int64                 `json:"end"`
	JSON  reasoningPartTimeJSON `json:"-"`
}

// reasoningPartTimeJSON contains the JSON metadata for the struct
// [ReasoningPartTime]
type reasoningPartTimeJSON struct {
	Start       apijson.Field
	End         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReasoningPartTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reasoningPartTimeJSON) RawJSON() string {
	return r.raw
}

type ReasoningPartType string

const (
	ReasoningPartTypeReasoning ReasoningPartType = "reasoning"
)

func (r ReasoningPartType) IsKnown() bool {
	switch r {
	case ReasoningPartTypeReasoning:
		return true
	}
	return false
}

type Session struct {
	ID          string            `json:"id,required"`
	Directory   string            `json:"directory,required"`
	ProjectID   string            `json:"projectID,required"`
	Time        SessionTime       `json:"time,required"`
	Title       string            `json:"title,required"`
	Version     string            `json:"version,required"`
	Agent       string            `json:"agent"`
	Cost        float64           `json:"cost"`
	Model       SessionModel      `json:"model"`
	ParentID    string            `json:"parentID"`
	Path        string            `json:"path"`
	Revert      SessionRevert     `json:"revert"`
	Share       SessionShare      `json:"share"`
	Slug        string            `json:"slug,required"`
	Summary     SessionSummary    `json:"summary"`
	Tokens      SessionTokens     `json:"tokens"`
	WorkspaceID string            `json:"workspaceID"`
	Permission  PermissionRuleset `json:"permission"`
	// This field can have the runtime type of [map[string]any].
	Metadata map[string]any `json:"metadata"`
	JSON     sessionJSON    `json:"-"`
}

// sessionJSON contains the JSON metadata for the struct [Session]
type sessionJSON struct {
	ID          apijson.Field
	Directory   apijson.Field
	ProjectID   apijson.Field
	Time        apijson.Field
	Title       apijson.Field
	Version     apijson.Field
	Agent       apijson.Field
	Cost        apijson.Field
	Model       apijson.Field
	ParentID    apijson.Field
	Path        apijson.Field
	Revert      apijson.Field
	Share       apijson.Field
	Slug        apijson.Field
	Summary     apijson.Field
	Tokens      apijson.Field
	WorkspaceID apijson.Field
	Permission  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Session) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionJSON) RawJSON() string {
	return r.raw
}

type SessionTime struct {
	Created    int64           `json:"created,required"`
	Updated    int64           `json:"updated,required"`
	Compacting int64           `json:"compacting"`
	Archived   int64           `json:"archived"`
	JSON       sessionTimeJSON `json:"-"`
}

// sessionTimeJSON contains the JSON metadata for the struct [SessionTime]
type sessionTimeJSON struct {
	Created     apijson.Field
	Updated     apijson.Field
	Compacting  apijson.Field
	Archived    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionTimeJSON) RawJSON() string {
	return r.raw
}

type SessionRevert struct {
	MessageID string            `json:"messageID,required"`
	Diff      string            `json:"diff"`
	PartID    string            `json:"partID"`
	Snapshot  string            `json:"snapshot"`
	JSON      sessionRevertJSON `json:"-"`
}

// sessionRevertJSON contains the JSON metadata for the struct [SessionRevert]
type sessionRevertJSON struct {
	MessageID   apijson.Field
	Diff        apijson.Field
	PartID      apijson.Field
	Snapshot    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionRevert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionRevertJSON) RawJSON() string {
	return r.raw
}

type SessionShare struct {
	URL  string           `json:"url,required"`
	JSON sessionShareJSON `json:"-"`
}

// sessionShareJSON contains the JSON metadata for the struct [SessionShare]
type sessionShareJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionShareJSON) RawJSON() string {
	return r.raw
}

type SessionSummary struct {
	Additions int64              `json:"additions,required"`
	Deletions int64              `json:"deletions,required"`
	Files     int64              `json:"files,required"`
	Diffs     []SnapshotFileDiff `json:"diffs"`
	JSON      sessionSummaryJSON `json:"-"`
}

// sessionSummaryJSON contains the JSON metadata for the struct [SessionSummary]
type sessionSummaryJSON struct {
	Additions   apijson.Field
	Deletions   apijson.Field
	Files       apijson.Field
	Diffs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionSummaryJSON) RawJSON() string {
	return r.raw
}

type SessionModel struct {
	ID         string           `json:"id,required"`
	ProviderID string           `json:"providerID,required"`
	Variant    string           `json:"variant"`
	JSON       sessionModelJSON `json:"-"`
}

// sessionModelJSON contains the JSON metadata for the struct [SessionModel]
type sessionModelJSON struct {
	ID          apijson.Field
	ProviderID  apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionModelJSON) RawJSON() string {
	return r.raw
}

type SessionTokens struct {
	Input     int64              `json:"input,required"`
	Output    int64              `json:"output,required"`
	Reasoning int64              `json:"reasoning,required"`
	Cache     SessionTokensCache `json:"cache,required"`
	JSON      sessionTokensJSON  `json:"-"`
}

// sessionTokensJSON contains the JSON metadata for the struct [SessionTokens]
type sessionTokensJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionTokensJSON) RawJSON() string {
	return r.raw
}

type SessionTokensCache struct {
	Read  int64                  `json:"read,required"`
	Write int64                  `json:"write,required"`
	JSON  sessionTokensCacheJSON `json:"-"`
}

// sessionTokensCacheJSON contains the JSON metadata for the struct [SessionTokensCache]
type sessionTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionTokensCacheJSON) RawJSON() string {
	return r.raw
}

type SnapshotPart struct {
	ID        string           `json:"id,required"`
	MessageID string           `json:"messageID,required"`
	SessionID string           `json:"sessionID,required"`
	Snapshot  string           `json:"snapshot,required"`
	Type      SnapshotPartType `json:"type,required"`
	JSON      snapshotPartJSON `json:"-"`
}

// snapshotPartJSON contains the JSON metadata for the struct [SnapshotPart]
type snapshotPartJSON struct {
	ID          apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Snapshot    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnapshotPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotPartJSON) RawJSON() string {
	return r.raw
}

func (r SnapshotPart) implementsPart() {}

type SnapshotPartType string

const (
	SnapshotPartTypeSnapshot SnapshotPartType = "snapshot"
)

func (r SnapshotPartType) IsKnown() bool {
	switch r {
	case SnapshotPartTypeSnapshot:
		return true
	}
	return false
}

type StepFinishPart struct {
	ID        string               `json:"id,required"`
	Cost      float64              `json:"cost,required"`
	MessageID string               `json:"messageID,required"`
	Reason    string               `json:"reason,required"`
	SessionID string               `json:"sessionID,required"`
	Tokens    StepFinishPartTokens `json:"tokens,required"`
	Type      StepFinishPartType   `json:"type,required"`
	Snapshot  string               `json:"snapshot"`
	JSON      stepFinishPartJSON   `json:"-"`
}

// stepFinishPartJSON contains the JSON metadata for the struct [StepFinishPart]
type stepFinishPartJSON struct {
	ID          apijson.Field
	Cost        apijson.Field
	MessageID   apijson.Field
	Reason      apijson.Field
	SessionID   apijson.Field
	Tokens      apijson.Field
	Type        apijson.Field
	Snapshot    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StepFinishPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stepFinishPartJSON) RawJSON() string {
	return r.raw
}

func (r StepFinishPart) implementsPart() {}

type StepFinishPartTokens struct {
	Cache     StepFinishPartTokensCache `json:"cache,required"`
	Input     int64                     `json:"input,required"`
	Output    int64                     `json:"output,required"`
	Reasoning int64                     `json:"reasoning,required"`
	Total     int64                     `json:"total"`
	JSON      stepFinishPartTokensJSON  `json:"-"`
}

type stepFinishPartTokensJSON struct {
	Cache       apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StepFinishPartTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stepFinishPartTokensJSON) RawJSON() string {
	return r.raw
}

type StepFinishPartTokensCache struct {
	Read  int64                         `json:"read,required"`
	Write int64                         `json:"write,required"`
	JSON  stepFinishPartTokensCacheJSON `json:"-"`
}

// stepFinishPartTokensCacheJSON contains the JSON metadata for the struct
// [StepFinishPartTokensCache]
type stepFinishPartTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StepFinishPartTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stepFinishPartTokensCacheJSON) RawJSON() string {
	return r.raw
}

type StepFinishPartType string

const (
	StepFinishPartTypeStepFinish StepFinishPartType = "step-finish"
)

func (r StepFinishPartType) IsKnown() bool {
	switch r {
	case StepFinishPartTypeStepFinish:
		return true
	}
	return false
}

type StepStartPart struct {
	ID        string            `json:"id,required"`
	MessageID string            `json:"messageID,required"`
	SessionID string            `json:"sessionID,required"`
	Type      StepStartPartType `json:"type,required"`
	Snapshot  string            `json:"snapshot"`
	JSON      stepStartPartJSON `json:"-"`
}

// stepStartPartJSON contains the JSON metadata for the struct [StepStartPart]
type stepStartPartJSON struct {
	ID          apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Type        apijson.Field
	Snapshot    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StepStartPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stepStartPartJSON) RawJSON() string {
	return r.raw
}

func (r StepStartPart) implementsPart() {}

type StepStartPartType string

const (
	StepStartPartTypeStepStart StepStartPartType = "step-start"
)

func (r StepStartPartType) IsKnown() bool {
	switch r {
	case StepStartPartTypeStepStart:
		return true
	}
	return false
}

type SymbolSource struct {
	Kind  int64              `json:"kind,required"`
	Name  string             `json:"name,required"`
	Path  string             `json:"path,required"`
	Range Range              `json:"range,required"`
	Text  FilePartSourceText `json:"text,required"`
	Type  SymbolSourceType   `json:"type,required"`
	JSON  symbolSourceJSON   `json:"-"`
}

// symbolSourceJSON contains the JSON metadata for the struct [SymbolSource]
type symbolSourceJSON struct {
	Kind        apijson.Field
	Name        apijson.Field
	Path        apijson.Field
	Range       apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SymbolSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r symbolSourceJSON) RawJSON() string {
	return r.raw
}

func (r SymbolSource) implementsFilePartSource() {}

// Deprecated: use [Range] instead.
type SymbolSourceRange = Range

// Deprecated: use [Position] instead.
type SymbolSourceRangeEnd = Position

// Deprecated: use [Position] instead.
type SymbolSourceRangeStart = Position

type SymbolSourceType string

const (
	SymbolSourceTypeSymbol SymbolSourceType = "symbol"
)

func (r SymbolSourceType) IsKnown() bool {
	switch r {
	case SymbolSourceTypeSymbol:
		return true
	}
	return false
}

type SymbolSourceParam struct {
	Kind  param.Field[int64]                   `json:"kind,required"`
	Name  param.Field[string]                  `json:"name,required"`
	Path  param.Field[string]                  `json:"path,required"`
	Range param.Field[SymbolSourceRangeParam]  `json:"range,required"`
	Text  param.Field[FilePartSourceTextParam] `json:"text,required"`
	Type  param.Field[SymbolSourceType]        `json:"type,required"`
}

func (r SymbolSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SymbolSourceParam) implementsFilePartSourceUnionParam() {}

type SymbolSourceRangeParam struct {
	End   param.Field[SymbolSourceRangeEndParam]   `json:"end,required"`
	Start param.Field[SymbolSourceRangeStartParam] `json:"start,required"`
}

func (r SymbolSourceRangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SymbolSourceRangeEndParam struct {
	Character param.Field[int64] `json:"character,required"`
	Line      param.Field[int64] `json:"line,required"`
}

func (r SymbolSourceRangeEndParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SymbolSourceRangeStartParam struct {
	Character param.Field[int64] `json:"character,required"`
	Line      param.Field[int64] `json:"line,required"`
}

func (r SymbolSourceRangeStartParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceSource struct {
	ClientName string             `json:"clientName,required"`
	Text       FilePartSourceText `json:"text,required"`
	Type       ResourceSourceType `json:"type,required"`
	URI        string             `json:"uri,required"`
	JSON       resourceSourceJSON `json:"-"`
}

// resourceSourceJSON contains the JSON metadata for the struct [ResourceSource]
type resourceSourceJSON struct {
	ClientName  apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSourceJSON) RawJSON() string {
	return r.raw
}

func (r ResourceSource) implementsFilePartSource() {}

type ResourceSourceType string

const (
	ResourceSourceTypeResource ResourceSourceType = "resource"
)

func (r ResourceSourceType) IsKnown() bool {
	switch r {
	case ResourceSourceTypeResource:
		return true
	}
	return false
}

type ResourceSourceParam struct {
	ClientName param.Field[string]                  `json:"clientName,required"`
	Text       param.Field[FilePartSourceTextParam] `json:"text,required"`
	Type       param.Field[ResourceSourceType]      `json:"type,required"`
	URI        param.Field[string]                  `json:"uri,required"`
}

func (r ResourceSourceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ResourceSourceParam) implementsFilePartSourceUnionParam() {}

type TextPart struct {
	ID        string         `json:"id,required"`
	MessageID string         `json:"messageID,required"`
	SessionID string         `json:"sessionID,required"`
	Text      string         `json:"text,required"`
	Type      TextPartType   `json:"type,required"`
	Metadata  map[string]any `json:"metadata"`
	Synthetic bool           `json:"synthetic"`
	Ignored   bool           `json:"ignored"`
	Time      TextPartTime   `json:"time"`
	JSON      textPartJSON   `json:"-"`
}

// textPartJSON contains the JSON metadata for the struct [TextPart]
type textPartJSON struct {
	ID          apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	Metadata    apijson.Field
	Synthetic   apijson.Field
	Ignored     apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TextPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r textPartJSON) RawJSON() string {
	return r.raw
}

func (r TextPart) implementsPart() {}

type SubtaskPart struct {
	ID          string           `json:"id,required"`
	MessageID   string           `json:"messageID,required"`
	SessionID   string           `json:"sessionID,required"`
	Type        SubtaskPartType  `json:"type,required"`
	Prompt      string           `json:"prompt,required"`
	Description string           `json:"description,required"`
	Agent       string           `json:"agent,required"`
	Model       SubtaskPartModel `json:"model"`
	Command     string           `json:"command"`
	JSON        subtaskPartJSON  `json:"-"`
}

type subtaskPartJSON struct {
	ID          apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Type        apijson.Field
	Prompt      apijson.Field
	Description apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubtaskPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subtaskPartJSON) RawJSON() string {
	return r.raw
}

func (r SubtaskPart) implementsPart() {}

type SubtaskPartType string

const (
	SubtaskPartTypeSubtask SubtaskPartType = "subtask"
)

func (r SubtaskPartType) IsKnown() bool {
	switch r {
	case SubtaskPartTypeSubtask:
		return true
	}
	return false
}

type SubtaskPartModel struct {
	ProviderID string               `json:"providerID,required"`
	ModelID    string               `json:"modelID,required"`
	JSON       subtaskPartModelJSON `json:"-"`
}

type subtaskPartModelJSON struct {
	ProviderID  apijson.Field
	ModelID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubtaskPartModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subtaskPartModelJSON) RawJSON() string {
	return r.raw
}

type CompactionPart struct {
	ID          string             `json:"id,required"`
	MessageID   string             `json:"messageID,required"`
	SessionID   string             `json:"sessionID,required"`
	Type        CompactionPartType `json:"type,required"`
	Auto        bool               `json:"auto,required"`
	Overflow    bool               `json:"overflow"`
	TailStartID string             `json:"tail_start_id"`
	JSON        compactionPartJSON `json:"-"`
}

type compactionPartJSON struct {
	ID          apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Type        apijson.Field
	Auto        apijson.Field
	Overflow    apijson.Field
	TailStartID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompactionPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compactionPartJSON) RawJSON() string {
	return r.raw
}

func (r CompactionPart) implementsPart() {}

type CompactionPartType string

const (
	CompactionPartTypeCompaction CompactionPartType = "compaction"
)

func (r CompactionPartType) IsKnown() bool {
	switch r {
	case CompactionPartTypeCompaction:
		return true
	}
	return false
}

type TextPartType string

const (
	TextPartTypeText TextPartType = "text"
)

func (r TextPartType) IsKnown() bool {
	switch r {
	case TextPartTypeText:
		return true
	}
	return false
}

type TextPartTime struct {
	Start int64            `json:"start,required"`
	End   int64            `json:"end"`
	JSON  textPartTimeJSON `json:"-"`
}

// textPartTimeJSON contains the JSON metadata for the struct [TextPartTime]
type textPartTimeJSON struct {
	Start       apijson.Field
	End         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TextPartTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r textPartTimeJSON) RawJSON() string {
	return r.raw
}

type TextPartInputParam struct {
	Text      param.Field[string]                 `json:"text,required"`
	Type      param.Field[TextPartInputType]      `json:"type,required"`
	ID        param.Field[string]                 `json:"id"`
	Metadata  param.Field[map[string]any]         `json:"metadata"`
	Synthetic param.Field[bool]                   `json:"synthetic"`
	Ignored   param.Field[bool]                   `json:"ignored"`
	Time      param.Field[TextPartInputTimeParam] `json:"time"`
}

func (r TextPartInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TextPartInputParam) implementsSessionPromptParamsPartUnion() {}

type TextPartInputType string

const (
	TextPartInputTypeText TextPartInputType = "text"
)

func (r TextPartInputType) IsKnown() bool {
	switch r {
	case TextPartInputTypeText:
		return true
	}
	return false
}

type TextPartInputTimeParam struct {
	Start param.Field[int64] `json:"start,required"`
	End   param.Field[int64] `json:"end"`
}

func (r TextPartInputTimeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolPart struct {
	ID        string         `json:"id,required"`
	CallID    string         `json:"callID,required"`
	MessageID string         `json:"messageID,required"`
	SessionID string         `json:"sessionID,required"`
	State     ToolPartState  `json:"state,required"`
	Tool      string         `json:"tool,required"`
	Type      ToolPartType   `json:"type,required"`
	Metadata  map[string]any `json:"metadata"`
	JSON      toolPartJSON   `json:"-"`
}

// toolPartJSON contains the JSON metadata for the struct [ToolPart]
type toolPartJSON struct {
	ID          apijson.Field
	CallID      apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	State       apijson.Field
	Tool        apijson.Field
	Type        apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolPart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolPartJSON) RawJSON() string {
	return r.raw
}

func (r ToolPart) implementsPart() {}

type ToolPartState struct {
	Status ToolPartStateStatus `json:"status,required"`
	// This field can have the runtime type of [[]FilePart].
	Attachments any            `json:"attachments"`
	Error       string         `json:"error"`
	Input       map[string]any `json:"input"`
	Metadata    map[string]any `json:"metadata"`
	Output      string         `json:"output"`
	Raw         string         `json:"raw"`
	// This field can have the runtime type of [ToolStateRunningTime],
	// [ToolStateCompletedTime], [ToolStateErrorTime].
	Time  any               `json:"time"`
	Title string            `json:"title"`
	JSON  toolPartStateJSON `json:"-"`
	union ToolPartStateUnion
}

// toolPartStateJSON contains the JSON metadata for the struct [ToolPartState]
type toolPartStateJSON struct {
	Status      apijson.Field
	Attachments apijson.Field
	Error       apijson.Field
	Input       apijson.Field
	Metadata    apijson.Field
	Output      apijson.Field
	Raw         apijson.Field
	Time        apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r toolPartStateJSON) RawJSON() string {
	return r.raw
}

func (r *ToolPartState) UnmarshalJSON(data []byte) (err error) {
	*r = ToolPartState{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ToolPartStateUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ToolStatePending], [ToolStateRunning],
// [ToolStateCompleted], [ToolStateError].
func (r ToolPartState) AsUnion() ToolPartStateUnion {
	return r.union
}

// Union satisfied by [ToolStatePending], [ToolStateRunning], [ToolStateCompleted]
// or [ToolStateError].
type ToolPartStateUnion interface {
	implementsToolPartState()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ToolPartStateUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ToolStatePending](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ToolStateRunning](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ToolStateCompleted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ToolStateError](),
		},
	)
}

type ToolPartStateStatus string

const (
	ToolPartStateStatusPending   ToolPartStateStatus = "pending"
	ToolPartStateStatusRunning   ToolPartStateStatus = "running"
	ToolPartStateStatusCompleted ToolPartStateStatus = "completed"
	ToolPartStateStatusError     ToolPartStateStatus = "error"
)

func (r ToolPartStateStatus) IsKnown() bool {
	switch r {
	case ToolPartStateStatusPending, ToolPartStateStatusRunning, ToolPartStateStatusCompleted, ToolPartStateStatusError:
		return true
	}
	return false
}

type ToolPartType string

const (
	ToolPartTypeTool ToolPartType = "tool"
)

func (r ToolPartType) IsKnown() bool {
	switch r {
	case ToolPartTypeTool:
		return true
	}
	return false
}

type ToolStateCompleted struct {
	Input       map[string]any           `json:"input,required"`
	Metadata    map[string]any           `json:"metadata,required"`
	Output      string                   `json:"output,required"`
	Status      ToolStateCompletedStatus `json:"status,required"`
	Time        ToolStateCompletedTime   `json:"time,required"`
	Title       string                   `json:"title,required"`
	Attachments []FilePart               `json:"attachments"`
	JSON        toolStateCompletedJSON   `json:"-"`
}

// toolStateCompletedJSON contains the JSON metadata for the struct
// [ToolStateCompleted]
type toolStateCompletedJSON struct {
	Input       apijson.Field
	Metadata    apijson.Field
	Output      apijson.Field
	Status      apijson.Field
	Time        apijson.Field
	Title       apijson.Field
	Attachments apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolStateCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolStateCompletedJSON) RawJSON() string {
	return r.raw
}

func (r ToolStateCompleted) implementsToolPartState() {}

type ToolStateCompletedStatus string

const (
	ToolStateCompletedStatusCompleted ToolStateCompletedStatus = "completed"
)

func (r ToolStateCompletedStatus) IsKnown() bool {
	switch r {
	case ToolStateCompletedStatusCompleted:
		return true
	}
	return false
}

type ToolStateCompletedTime struct {
	End       int64                      `json:"end,required"`
	Start     int64                      `json:"start,required"`
	Compacted int64                      `json:"compacted"`
	JSON      toolStateCompletedTimeJSON `json:"-"`
}

// toolStateCompletedTimeJSON contains the JSON metadata for the struct
// [ToolStateCompletedTime]
type toolStateCompletedTimeJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Compacted   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolStateCompletedTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolStateCompletedTimeJSON) RawJSON() string {
	return r.raw
}

type ToolStateError struct {
	Error    string               `json:"error,required"`
	Input    map[string]any       `json:"input,required"`
	Status   ToolStateErrorStatus `json:"status,required"`
	Time     ToolStateErrorTime   `json:"time,required"`
	Metadata map[string]any       `json:"metadata"`
	JSON     toolStateErrorJSON   `json:"-"`
}

// toolStateErrorJSON contains the JSON metadata for the struct [ToolStateError]
type toolStateErrorJSON struct {
	Error       apijson.Field
	Input       apijson.Field
	Status      apijson.Field
	Time        apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolStateError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolStateErrorJSON) RawJSON() string {
	return r.raw
}

func (r ToolStateError) implementsToolPartState() {}

type ToolStateErrorStatus string

const (
	ToolStateErrorStatusError ToolStateErrorStatus = "error"
)

func (r ToolStateErrorStatus) IsKnown() bool {
	switch r {
	case ToolStateErrorStatusError:
		return true
	}
	return false
}

type ToolStateErrorTime struct {
	End   int64                  `json:"end,required"`
	Start int64                  `json:"start,required"`
	JSON  toolStateErrorTimeJSON `json:"-"`
}

// toolStateErrorTimeJSON contains the JSON metadata for the struct
// [ToolStateErrorTime]
type toolStateErrorTimeJSON struct {
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolStateErrorTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolStateErrorTimeJSON) RawJSON() string {
	return r.raw
}

type ToolStatePending struct {
	Status ToolStatePendingStatus `json:"status,required"`
	Input  map[string]any         `json:"input,required"`
	Raw    string                 `json:"raw,required"`
	JSON   toolStatePendingJSON   `json:"-"`
}

// toolStatePendingJSON contains the JSON metadata for the struct
// [ToolStatePending]
type toolStatePendingJSON struct {
	Status      apijson.Field
	Input       apijson.Field
	Raw         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolStatePending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolStatePendingJSON) RawJSON() string {
	return r.raw
}

func (r ToolStatePending) implementsToolPartState() {}

type ToolStatePendingStatus string

const (
	ToolStatePendingStatusPending ToolStatePendingStatus = "pending"
)

func (r ToolStatePendingStatus) IsKnown() bool {
	switch r {
	case ToolStatePendingStatusPending:
		return true
	}
	return false
}

type ToolStateRunning struct {
	Input    map[string]any         `json:"input,required"`
	Status   ToolStateRunningStatus `json:"status,required"`
	Time     ToolStateRunningTime   `json:"time,required"`
	Metadata map[string]any         `json:"metadata"`
	Title    string                 `json:"title"`
	JSON     toolStateRunningJSON   `json:"-"`
}

// toolStateRunningJSON contains the JSON metadata for the struct
// [ToolStateRunning]
type toolStateRunningJSON struct {
	Input       apijson.Field
	Status      apijson.Field
	Time        apijson.Field
	Metadata    apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolStateRunning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolStateRunningJSON) RawJSON() string {
	return r.raw
}

func (r ToolStateRunning) implementsToolPartState() {}

type ToolStateRunningStatus string

const (
	ToolStateRunningStatusRunning ToolStateRunningStatus = "running"
)

func (r ToolStateRunningStatus) IsKnown() bool {
	switch r {
	case ToolStateRunningStatusRunning:
		return true
	}
	return false
}

type ToolStateRunningTime struct {
	Start int64                    `json:"start,required"`
	JSON  toolStateRunningTimeJSON `json:"-"`
}

// toolStateRunningTimeJSON contains the JSON metadata for the struct
// [ToolStateRunningTime]
type toolStateRunningTimeJSON struct {
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolStateRunningTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolStateRunningTimeJSON) RawJSON() string {
	return r.raw
}

type UserMessage struct {
	ID        string             `json:"id,required"`
	Agent     string             `json:"agent,required"`
	Model     UserMessageModel   `json:"model,required"`
	Role      UserMessageRole    `json:"role,required"`
	SessionID string             `json:"sessionID,required"`
	Time      UserMessageTime    `json:"time,required"`
	Format    OutputFormat       `json:"format,omitzero"`
	System    string             `json:"system,omitzero"`
	Tools     map[string]bool    `json:"tools,omitzero"`
	Summary   UserMessageSummary `json:"summary"`
	JSON      userMessageJSON    `json:"-"`
}

// userMessageJSON contains the JSON metadata for the struct [UserMessage]
type userMessageJSON struct {
	ID          apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
	Role        apijson.Field
	SessionID   apijson.Field
	Time        apijson.Field
	Format      apijson.Field
	System      apijson.Field
	Tools       apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMessageJSON) RawJSON() string {
	return r.raw
}

func (r UserMessage) implementsMessage() {}

type UserMessageRole string

const (
	UserMessageRoleUser UserMessageRole = "user"
)

func (r UserMessageRole) IsKnown() bool {
	switch r {
	case UserMessageRoleUser:
		return true
	}
	return false
}

type UserMessageModel struct {
	ProviderID string               `json:"providerID,required"`
	ModelID    string               `json:"modelID,required"`
	Variant    string               `json:"variant"`
	JSON       userMessageModelJSON `json:"-"`
}

// userMessageModelJSON contains the JSON metadata for the struct [UserMessageModel]
type userMessageModelJSON struct {
	ProviderID  apijson.Field
	ModelID     apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMessageModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMessageModelJSON) RawJSON() string {
	return r.raw
}

type UserMessageTime struct {
	Created int64               `json:"created,required"`
	JSON    userMessageTimeJSON `json:"-"`
}

// userMessageTimeJSON contains the JSON metadata for the struct [UserMessageTime]
type userMessageTimeJSON struct {
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMessageTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMessageTimeJSON) RawJSON() string {
	return r.raw
}

type UserMessageSummary struct {
	Diffs []SnapshotFileDiff     `json:"diffs,required"`
	Body  string                 `json:"body"`
	Title string                 `json:"title"`
	JSON  userMessageSummaryJSON `json:"-"`
}

// userMessageSummaryJSON contains the JSON metadata for the struct
// [UserMessageSummary]
type userMessageSummaryJSON struct {
	Diffs       apijson.Field
	Body        apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMessageSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMessageSummaryJSON) RawJSON() string {
	return r.raw
}

type SessionCommandResponse struct {
	Info  AssistantMessage           `json:"info,required"`
	Parts []Part                     `json:"parts,required"`
	JSON  sessionCommandResponseJSON `json:"-"`
}

// sessionCommandResponseJSON contains the JSON metadata for the struct
// [SessionCommandResponse]
type sessionCommandResponseJSON struct {
	Info        apijson.Field
	Parts       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionCommandResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionCommandResponseJSON) RawJSON() string {
	return r.raw
}

type SessionMessageResponse struct {
	Info  Message                    `json:"info,required"`
	Parts []Part                     `json:"parts,required"`
	JSON  sessionMessageResponseJSON `json:"-"`
}

// sessionMessageResponseJSON contains the JSON metadata for the struct
// [SessionMessageResponse]
type sessionMessageResponseJSON struct {
	Info        apijson.Field
	Parts       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionMessageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionMessageResponseJSON) RawJSON() string {
	return r.raw
}

type SessionMessagesResponse struct {
	Info  Message                     `json:"info,required"`
	Parts []Part                      `json:"parts,required"`
	JSON  sessionMessagesResponseJSON `json:"-"`
}

// sessionMessagesResponseJSON contains the JSON metadata for the struct
// [SessionMessagesResponse]
type sessionMessagesResponseJSON struct {
	Info        apijson.Field
	Parts       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionMessagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionMessagesResponseJSON) RawJSON() string {
	return r.raw
}

type SessionPromptResponse struct {
	Info  AssistantMessage          `json:"info,required"`
	Parts []Part                    `json:"parts,required"`
	JSON  sessionPromptResponseJSON `json:"-"`
}

// sessionPromptResponseJSON contains the JSON metadata for the struct
// [SessionPromptResponse]
type sessionPromptResponseJSON struct {
	Info        apijson.Field
	Parts       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionPromptResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionPromptResponseJSON) RawJSON() string {
	return r.raw
}

type SessionNewParams struct {
	Directory   param.Field[string]                `query:"directory"`
	Workspace   param.Field[string]                `query:"workspace"`
	ParentID    param.Field[string]                `json:"parentID"`
	Title       param.Field[string]                `json:"title"`
	Agent       param.Field[string]                `json:"agent"`
	Model       param.Field[SessionNewParamsModel] `json:"model"`
	Permission  param.Field[[]PermissionRule]      `json:"permission"`
	WorkspaceID param.Field[string]                `json:"workspaceID"`
	Metadata    param.Field[any]                   `json:"metadata"`
}

type SessionNewParamsModel struct {
	ID         param.Field[string] `json:"id,required"`
	ProviderID param.Field[string] `json:"providerID,required"`
	Variant    param.Field[string] `json:"variant"`
}

func (r SessionNewParamsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SessionNewParams]'s query parameters as `url.Values`.
func (r SessionNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionUpdateParams struct {
	Directory  param.Field[string]                  `query:"directory"`
	Workspace  param.Field[string]                  `query:"workspace"`
	Title      param.Field[string]                  `json:"title"`
	Permission param.Field[[]PermissionRule]        `json:"permission"`
	Time       param.Field[SessionUpdateParamsTime] `json:"time"`
	Metadata   param.Field[any]                     `json:"metadata"`
}

type SessionUpdateParamsTime struct {
	Archived param.Field[int64] `json:"archived"`
}

func (r SessionUpdateParamsTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SessionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SessionUpdateParams]'s query parameters as `url.Values`.
func (r SessionUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionListParams struct {
	Directory param.Field[string]                 `query:"directory"`
	Workspace param.Field[string]                 `query:"workspace"`
	Scope     param.Field[SessionListParamsScope] `query:"scope"`
	Path      param.Field[string]                 `query:"path"`
	Roots     param.Field[bool]                   `query:"roots"`
	Start     param.Field[int64]                  `query:"start"`
	Search    param.Field[string]                 `query:"search"`
	Limit     param.Field[int64]                  `query:"limit"`
}

// URLQuery serializes [SessionListParams]'s query parameters as `url.Values`.
func (r SessionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionListParamsScope string

const (
	SessionListParamsScopeProject SessionListParamsScope = "project"
)

func (r SessionListParamsScope) IsKnown() bool {
	switch r {
	case SessionListParamsScopeProject:
		return true
	}
	return false
}

type SessionDeleteParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SessionDeleteParams]'s query parameters as `url.Values`.
func (r SessionDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionAbortParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SessionAbortParams]'s query parameters as `url.Values`.
func (r SessionAbortParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionChildrenParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SessionChildrenParams]'s query parameters as `url.Values`.
func (r SessionChildrenParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionTodoParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SessionTodoParams]'s query parameters as `url.Values`.
func (r SessionTodoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionCommandParams struct {
	Directory param.Field[string]                     `query:"directory"`
	Workspace param.Field[string]                     `query:"workspace"`
	Agent     param.Field[string]                     `json:"agent"`
	Arguments param.Field[string]                     `json:"arguments,required"`
	Command   param.Field[string]                     `json:"command,required"`
	MessageID param.Field[string]                     `json:"messageID"`
	Model     param.Field[string]                     `json:"model"`
	Variant   param.Field[string]                     `json:"variant"`
	Parts     param.Field[[]SessionCommandParamsPart] `json:"parts"`
}

type SessionCommandParamsPart struct {
	ID       param.Field[string]                        `json:"id"`
	Type     param.Field[SessionCommandParamsPartsType] `json:"type,required"`
	Mime     param.Field[string]                        `json:"mime,required"`
	Filename param.Field[string]                        `json:"filename"`
	URL      param.Field[string]                        `json:"url,required"`
	Source   param.Field[FilePartSourceUnionParam]      `json:"source"`
}

func (r SessionCommandParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionCommandParamsPartsType string

const (
	SessionCommandParamsPartsTypeFile SessionCommandParamsPartsType = "file"
)

func (r SessionCommandParamsPartsType) IsKnown() bool {
	switch r {
	case SessionCommandParamsPartsTypeFile:
		return true
	}
	return false
}

func (r SessionCommandParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SessionCommandParams]'s query parameters as `url.Values`.
func (r SessionCommandParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SessionGetParams]'s query parameters as `url.Values`.
func (r SessionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionInitParams struct {
	Directory  param.Field[string] `query:"directory"`
	Workspace  param.Field[string] `query:"workspace"`
	MessageID  param.Field[string] `json:"messageID,required"`
	ModelID    param.Field[string] `json:"modelID,required"`
	ProviderID param.Field[string] `json:"providerID,required"`
}

func (r SessionInitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SessionInitParams]'s query parameters as `url.Values`.
func (r SessionInitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionMessageParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SessionMessageParams]'s query parameters as `url.Values`.
func (r SessionMessageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionMessagesParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Limit     param.Field[int64]  `query:"limit"`
	Before    param.Field[string] `query:"before"`
}

// URLQuery serializes [SessionMessagesParams]'s query parameters as `url.Values`.
func (r SessionMessagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionPromptParams struct {
	Parts     param.Field[[]SessionPromptParamsPartUnion] `json:"parts,required"`
	Directory param.Field[string]                         `query:"directory"`
	Workspace param.Field[string]                         `query:"workspace"`
	Agent     param.Field[string]                         `json:"agent"`
	MessageID param.Field[string]                         `json:"messageID"`
	Model     param.Field[SessionPromptParamsModel]       `json:"model"`
	NoReply   param.Field[bool]                           `json:"noReply"`
	System    param.Field[string]                         `json:"system"`
	Tools     param.Field[map[string]bool]                `json:"tools"`
	Format    param.Field[SessionPromptParamsFormatUnion] `json:"format"`
	Variant   param.Field[string]                         `json:"variant"`
}

func (r SessionPromptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SessionPromptParams]'s query parameters as `url.Values`.
func (r SessionPromptParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionPromptParamsPart struct {
	Type      param.Field[SessionPromptParamsPartsType] `json:"type,required"`
	ID        param.Field[string]                       `json:"id"`
	Filename  param.Field[string]                       `json:"filename"`
	Metadata  param.Field[any]                          `json:"metadata"`
	Mime      param.Field[string]                       `json:"mime"`
	Name      param.Field[string]                       `json:"name"`
	Source    param.Field[any]                          `json:"source"`
	Synthetic param.Field[bool]                         `json:"synthetic"`
	Text      param.Field[string]                       `json:"text"`
	Time      param.Field[any]                          `json:"time"`
	URL       param.Field[string]                       `json:"url"`
}

func (r SessionPromptParamsPart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SessionPromptParamsPart) implementsSessionPromptParamsPartUnion() {}

// Satisfied by [TextPartInputParam], [FilePartInputParam], [AgentPartInputParam],
// [SubtaskPartInputParam], [SessionPromptParamsPart].
type SessionPromptParamsPartUnion interface {
	implementsSessionPromptParamsPartUnion()
}

type SessionPromptParamsPartsType string

const (
	SessionPromptParamsPartsTypeText    SessionPromptParamsPartsType = "text"
	SessionPromptParamsPartsTypeFile    SessionPromptParamsPartsType = "file"
	SessionPromptParamsPartsTypeAgent   SessionPromptParamsPartsType = "agent"
	SessionPromptParamsPartsTypeSubtask SessionPromptParamsPartsType = "subtask"
)

func (r SessionPromptParamsPartsType) IsKnown() bool {
	switch r {
	case SessionPromptParamsPartsTypeText, SessionPromptParamsPartsTypeFile, SessionPromptParamsPartsTypeAgent, SessionPromptParamsPartsTypeSubtask:
		return true
	}
	return false
}

// Satisfied by [SessionPromptParamsFormatText],
// [SessionPromptParamsFormatJsonSchema].
type SessionPromptParamsFormatUnion interface {
	implementsSessionPromptParamsFormatUnion()
}

type SessionPromptParamsFormatText struct {
	Type param.Field[SessionPromptParamsFormatTextType] `json:"type,required"`
}

func (r SessionPromptParamsFormatText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SessionPromptParamsFormatText) implementsSessionPromptParamsFormatUnion() {}

type SessionPromptParamsFormatTextType string

const (
	SessionPromptParamsFormatTextTypeText SessionPromptParamsFormatTextType = "text"
)

func (r SessionPromptParamsFormatTextType) IsKnown() bool {
	switch r {
	case SessionPromptParamsFormatTextTypeText:
		return true
	}
	return false
}

type SessionPromptParamsFormatJsonSchema struct {
	Type       param.Field[SessionPromptParamsFormatJsonSchemaType] `json:"type,required"`
	Schema     param.Field[any]                                     `json:"schema,required"`
	RetryCount param.Field[int64]                                   `json:"retryCount"`
}

func (r SessionPromptParamsFormatJsonSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SessionPromptParamsFormatJsonSchema) implementsSessionPromptParamsFormatUnion() {}

type SessionPromptParamsFormatJsonSchemaType string

const (
	SessionPromptParamsFormatJsonSchemaTypeJsonSchema SessionPromptParamsFormatJsonSchemaType = "json_schema"
)

func (r SessionPromptParamsFormatJsonSchemaType) IsKnown() bool {
	switch r {
	case SessionPromptParamsFormatJsonSchemaTypeJsonSchema:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[SessionPromptParamsPartUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[TextPartInputParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[FilePartInputParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[AgentPartInputParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SubtaskPartInputParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SessionPromptParamsPart](),
		},
	)
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[SessionPromptParamsFormatUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SessionPromptParamsFormatText](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[SessionPromptParamsFormatJsonSchema](),
		},
	)
}

type SessionPromptParamsModel struct {
	ModelID    param.Field[string] `json:"modelID,required"`
	ProviderID param.Field[string] `json:"providerID,required"`
}

func (r SessionPromptParamsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SessionRevertParams struct {
	MessageID param.Field[string] `json:"messageID,required"`
	PartID    param.Field[string] `json:"partID"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r SessionRevertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SessionRevertParams]'s query parameters as `url.Values`.
func (r SessionRevertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionShareParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SessionShareParams]'s query parameters as `url.Values`.
func (r SessionShareParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionShellParams struct {
	Directory param.Field[string]                   `query:"directory"`
	Workspace param.Field[string]                   `query:"workspace"`
	Agent     param.Field[string]                   `json:"agent,required"`
	Command   param.Field[string]                   `json:"command,required"`
	MessageID param.Field[string]                   `json:"messageID"`
	Model     param.Field[SessionPromptParamsModel] `json:"model"`
}

func (r SessionShellParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SessionShellParams]'s query parameters as `url.Values`.
func (r SessionShellParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionSummarizeParams struct {
	Directory  param.Field[string] `query:"directory"`
	Workspace  param.Field[string] `query:"workspace"`
	Auto       param.Field[bool]   `json:"auto"`
	ModelID    param.Field[string] `json:"modelID,required"`
	ProviderID param.Field[string] `json:"providerID,required"`
}

func (r SessionSummarizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [SessionSummarizeParams]'s query parameters as `url.Values`.
func (r SessionSummarizeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionUnrevertParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SessionUnrevertParams]'s query parameters as `url.Values`.
func (r SessionUnrevertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionUnshareParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SessionUnshareParams]'s query parameters as `url.Values`.
func (r SessionUnshareParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// SessionStatusIdle represents an idle session status
type SessionStatusIdle struct {
	Type SessionStatusIdleType `json:"type,required"`
	JSON sessionStatusIdleJSON `json:"-"`
}

type sessionStatusIdleJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r sessionStatusIdleJSON) RawJSON() string {
	return r.raw
}

func (r *SessionStatusIdle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SessionStatusIdle) implementsSessionStatus() {}

type SessionStatusIdleType string

const (
	SessionStatusIdleTypeIdle SessionStatusIdleType = "idle"
)

func (r SessionStatusIdleType) IsKnown() bool {
	switch r {
	case SessionStatusIdleTypeIdle:
		return true
	}
	return false
}

// SessionStatusRetry represents a retry session status
type SessionStatusRetry struct {
	Type    SessionStatusRetryType   `json:"type,required"`
	Attempt int64                    `json:"attempt,required"`
	Message string                   `json:"message,required"`
	Action  SessionStatusRetryAction `json:"action,omitzero"`
	Next    int64                    `json:"next,required"`
	JSON    sessionStatusRetryJSON   `json:"-"`
}

type sessionStatusRetryJSON struct {
	Type        apijson.Field
	Attempt     apijson.Field
	Message     apijson.Field
	Action      apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

// SessionStatusRetryAction represents the action to retry a session
type SessionStatusRetryAction struct {
	Reason   string                       `json:"reason,required"`
	Provider string                       `json:"provider,required"`
	Title    string                       `json:"title,required"`
	Message  string                       `json:"message,required"`
	Label    string                       `json:"label,required"`
	Link     string                       `json:"link,omitzero"`
	JSON     sessionStatusRetryActionJSON `json:"-"`
}

type sessionStatusRetryActionJSON struct {
	Reason      apijson.Field
	Provider    apijson.Field
	Title       apijson.Field
	Message     apijson.Field
	Label       apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r sessionStatusRetryActionJSON) RawJSON() string {
	return r.raw
}

func (r *SessionStatusRetryAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionStatusRetryJSON) RawJSON() string {
	return r.raw
}

func (r *SessionStatusRetry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SessionStatusRetry) implementsSessionStatus() {}

type SessionStatusRetryType string

const (
	SessionStatusRetryTypeRetry SessionStatusRetryType = "retry"
)

func (r SessionStatusRetryType) IsKnown() bool {
	switch r {
	case SessionStatusRetryTypeRetry:
		return true
	}
	return false
}

// SessionStatusBusy represents a busy session status
type SessionStatusBusy struct {
	Type SessionStatusBusyType `json:"type,required"`
	JSON sessionStatusBusyJSON `json:"-"`
}

type sessionStatusBusyJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r sessionStatusBusyJSON) RawJSON() string {
	return r.raw
}

func (r *SessionStatusBusy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SessionStatusBusy) implementsSessionStatus() {}

type SessionStatusBusyType string

const (
	SessionStatusBusyTypeBusy SessionStatusBusyType = "busy"
)

func (r SessionStatusBusyType) IsKnown() bool {
	switch r {
	case SessionStatusBusyTypeBusy:
		return true
	}
	return false
}

// SessionStatus is the OpenAPI `SessionStatus` anyOf union, carrying the
// flattened superset of every variant's fields. `Type` discriminates which
// variant the payload actually is; use [SessionStatus.AsUnion] to recover the
// concrete variant.
type SessionStatus struct {
	Type    SessionStatusType `json:"type,required"`
	Attempt int64             `json:"attempt"`
	Message string            `json:"message"`
	// This field is only present on the [SessionStatusRetry] variant.
	Action SessionStatusRetryAction `json:"action,omitzero"`
	Next   int64                    `json:"next"`
	JSON   sessionStatusJSON        `json:"-"`
	union  SessionStatusUnion
}

// sessionStatusJSON contains the JSON metadata for the struct [SessionStatus]
type sessionStatusJSON struct {
	Type        apijson.Field
	Attempt     apijson.Field
	Message     apijson.Field
	Action      apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r sessionStatusJSON) RawJSON() string {
	return r.raw
}

func (r *SessionStatus) UnmarshalJSON(data []byte) (err error) {
	*r = SessionStatus{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SessionStatusUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [SessionStatusIdle],
// [SessionStatusRetry], [SessionStatusBusy].
func (r SessionStatus) AsUnion() SessionStatusUnion {
	return r.union
}

// Union satisfied by [SessionStatusIdle], [SessionStatusRetry] or
// [SessionStatusBusy].
type SessionStatusUnion interface {
	implementsSessionStatus()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[SessionStatusUnion](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "idle",
			Type:               reflect.TypeFor[SessionStatusIdle](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "retry",
			Type:               reflect.TypeFor[SessionStatusRetry](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "busy",
			Type:               reflect.TypeFor[SessionStatusBusy](),
		},
	)
}

type SessionStatusType string

const (
	SessionStatusTypeIdle  SessionStatusType = "idle"
	SessionStatusTypeRetry SessionStatusType = "retry"
	SessionStatusTypeBusy  SessionStatusType = "busy"
)

func (r SessionStatusType) IsKnown() bool {
	switch r {
	case SessionStatusTypeIdle, SessionStatusTypeRetry, SessionStatusTypeBusy:
		return true
	}
	return false
}

// SessionStatusMap is a map of session IDs to their status, returned by
// [SessionService.Status].
type SessionStatusMap map[string]SessionStatus

// SnapshotFileDiff represents a diff for a file in a snapshot
type SnapshotFileDiff struct {
	File      string                 `json:"file"`
	Patch     string                 `json:"patch"`
	Additions int64                  `json:"additions,required"`
	Deletions int64                  `json:"deletions,required"`
	Status    SnapshotFileDiffStatus `json:"status"`
	JSON      snapshotFileDiffJSON   `json:"-"`
}

type snapshotFileDiffJSON struct {
	File        apijson.Field
	Patch       apijson.Field
	Additions   apijson.Field
	Deletions   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnapshotFileDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snapshotFileDiffJSON) RawJSON() string {
	return r.raw
}

// SnapshotFileDiffStatus represents the status of a file diff
type SnapshotFileDiffStatus string

const (
	SnapshotFileDiffStatusAdded    SnapshotFileDiffStatus = "added"
	SnapshotFileDiffStatusDeleted  SnapshotFileDiffStatus = "deleted"
	SnapshotFileDiffStatusModified SnapshotFileDiffStatus = "modified"
)

func (r SnapshotFileDiffStatus) IsKnown() bool {
	switch r {
	case SnapshotFileDiffStatusAdded, SnapshotFileDiffStatusDeleted, SnapshotFileDiffStatusModified:
		return true
	}
	return false
}

type SessionForkParams struct {
	MessageID param.Field[string] `json:"messageID"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r SessionForkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SessionForkParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionDiffParams struct {
	MessageID param.Field[string] `query:"messageID"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r SessionDiffParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionDeleteMessageParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r SessionDeleteMessageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r SessionStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
