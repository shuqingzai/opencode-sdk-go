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

// PartService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPartService] method instead.
type PartService struct {
	Options []option.RequestOption
}

// NewPartService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPartService(opts ...option.RequestOption) (r *PartService) {
	r = &PartService{}
	r.Options = opts
	return
}

// Delete a part from a message
func (r *PartService) Delete(ctx context.Context, sessionID string, messageID string, partID string, body PartDeleteParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	if partID == "" {
		err = errors.New("missing required partID parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message/%s/part/%s", sessionID, messageID, partID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Update a part in a message
func (r *PartService) Update(ctx context.Context, sessionID string, messageID string, partID string, params PartUpdateParams, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	if partID == "" {
		err = errors.New("missing required partID parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message/%s/part/%s", sessionID, messageID, partID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type PartDeleteParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r PartDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// PartUpdateParams is the request payload for [PartService.Update].
//
// The PATCH body itself is a [Part] object (OpenAPI: $ref: Part, no `required`
// key in `requestBody`), so [PartUpdateParams.Part]'s value is serialized as the
// request body root when present. When Part is not set, no body payload is
// emitted — the JS SDK treats `body` as optional and omits it entirely.
type PartUpdateParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	// Part is the PATCH body. Its value is serialized as the body root when
	// present (any of the 12 [PartUpdatePartUnion] variants). When unset, no
	// body is sent. The `json:"-"` tag prevents the field from being
	// serialized as a `{ "Part": ... }` wrapper.
	Part param.Field[PartUpdatePartUnion] `json:"-"`
}

func (r PartUpdateParams) MarshalJSON() (data []byte, err error) {
	if r.Part.Present {
		return apijson.MarshalRoot(r.Part)
	}
	return nil, nil
}

func (r PartUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Satisfied by [PartUpdatePartText], [PartUpdatePartSubtask],
// [PartUpdatePartReasoning], [PartUpdatePartFile], [PartUpdatePartTool],
// [PartUpdatePartStepStart], [PartUpdatePartStepFinish], [PartUpdatePartSnapshot],
// [PartUpdatePartPatch], [PartUpdatePartAgent], [PartUpdatePartRetry] or
// [PartUpdatePartCompaction].
type PartUpdatePartUnion interface {
	implementsPartUpdatePartUnion()
}

type PartUpdatePartText struct {
	ID        param.Field[string]                 `json:"id,required"`
	MessageID param.Field[string]                 `json:"messageID,required"`
	SessionID param.Field[string]                 `json:"sessionID,required"`
	Text      param.Field[string]                 `json:"text,required"`
	Type      param.Field[PartUpdatePartTextType] `json:"type,required"`
	Ignored   param.Field[bool]                   `json:"ignored"`
	Metadata  param.Field[map[string]any]         `json:"metadata"`
	Synthetic param.Field[bool]                   `json:"synthetic"`
	Time      param.Field[PartUpdatePartTextTime] `json:"time"`
}

func (r PartUpdatePartText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartText) implementsPartUpdatePartUnion() {}

type PartUpdatePartTextType string

const (
	PartUpdatePartTextTypeText PartUpdatePartTextType = "text"
)

func (r PartUpdatePartTextType) IsKnown() bool {
	switch r {
	case PartUpdatePartTextTypeText:
		return true
	}
	return false
}

type PartUpdatePartTextTime struct {
	Start param.Field[int64] `json:"start,required"`
	End   param.Field[int64] `json:"end"`
}

func (r PartUpdatePartTextTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdatePartSubtask struct {
	Agent       param.Field[string]                     `json:"agent,required"`
	Description param.Field[string]                     `json:"description,required"`
	ID          param.Field[string]                     `json:"id,required"`
	MessageID   param.Field[string]                     `json:"messageID,required"`
	Prompt      param.Field[string]                     `json:"prompt,required"`
	SessionID   param.Field[string]                     `json:"sessionID,required"`
	Type        param.Field[PartUpdatePartSubtaskType]  `json:"type,required"`
	Command     param.Field[string]                     `json:"command"`
	Model       param.Field[PartUpdatePartSubtaskModel] `json:"model"`
}

func (r PartUpdatePartSubtask) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartSubtask) implementsPartUpdatePartUnion() {}

type PartUpdatePartSubtaskType string

const (
	PartUpdatePartSubtaskTypeSubtask PartUpdatePartSubtaskType = "subtask"
)

func (r PartUpdatePartSubtaskType) IsKnown() bool {
	switch r {
	case PartUpdatePartSubtaskTypeSubtask:
		return true
	}
	return false
}

type PartUpdatePartSubtaskModel struct {
	ModelID    param.Field[string] `json:"modelID,required"`
	ProviderID param.Field[string] `json:"providerID,required"`
}

func (r PartUpdatePartSubtaskModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdatePartReasoning struct {
	ID        param.Field[string]                      `json:"id,required"`
	MessageID param.Field[string]                      `json:"messageID,required"`
	SessionID param.Field[string]                      `json:"sessionID,required"`
	Text      param.Field[string]                      `json:"text,required"`
	Time      param.Field[PartUpdatePartReasoningTime] `json:"time,required"`
	Type      param.Field[PartUpdatePartReasoningType] `json:"type,required"`
	Metadata  param.Field[map[string]any]              `json:"metadata"`
}

func (r PartUpdatePartReasoning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartReasoning) implementsPartUpdatePartUnion() {}

type PartUpdatePartReasoningType string

const (
	PartUpdatePartReasoningTypeReasoning PartUpdatePartReasoningType = "reasoning"
)

func (r PartUpdatePartReasoningType) IsKnown() bool {
	switch r {
	case PartUpdatePartReasoningTypeReasoning:
		return true
	}
	return false
}

type PartUpdatePartReasoningTime struct {
	Start param.Field[int64] `json:"start,required"`
	End   param.Field[int64] `json:"end"`
}

func (r PartUpdatePartReasoningTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdatePartFile struct {
	ID        param.Field[string]                 `json:"id,required"`
	MessageID param.Field[string]                 `json:"messageID,required"`
	Mime      param.Field[string]                 `json:"mime,required"`
	SessionID param.Field[string]                 `json:"sessionID,required"`
	Type      param.Field[PartUpdatePartFileType] `json:"type,required"`
	URL       param.Field[string]                 `json:"url,required"`
	Filename  param.Field[string]                 `json:"filename"`
	// Source of the file (e.g. symbol, resource, or inline file source).
	// Accepts [map[string]any] or an appropriate FilePartSource variant.
	Source param.Field[any] `json:"source"`
}

func (r PartUpdatePartFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartFile) implementsPartUpdatePartUnion() {}

type PartUpdatePartFileType string

const (
	PartUpdatePartFileTypeFile PartUpdatePartFileType = "file"
)

func (r PartUpdatePartFileType) IsKnown() bool {
	switch r {
	case PartUpdatePartFileTypeFile:
		return true
	}
	return false
}

type PartUpdatePartTool struct {
	CallID    param.Field[string]                       `json:"callID,required"`
	ID        param.Field[string]                       `json:"id,required"`
	MessageID param.Field[string]                       `json:"messageID,required"`
	SessionID param.Field[string]                       `json:"sessionID,required"`
	State     param.Field[PartUpdatePartToolStateUnion] `json:"state,required"`
	Tool      param.Field[string]                       `json:"tool,required"`
	Type      param.Field[PartUpdatePartToolType]       `json:"type,required"`
	Metadata  param.Field[map[string]any]               `json:"metadata"`
}

func (r PartUpdatePartTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartTool) implementsPartUpdatePartUnion() {}

// PartUpdatePartToolStatePending represents the pending variant of
// [PartUpdatePartToolStateUnion].
type PartUpdatePartToolStatePending struct {
	Status param.Field[PartUpdatePartToolStatePendingStatus] `json:"status,required"`
}

func (r PartUpdatePartToolStatePending) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartToolStatePending) implementsPartUpdatePartToolStateUnion() {}

type PartUpdatePartToolStatePendingStatus string

const (
	PartUpdatePartToolStatePendingStatusPending PartUpdatePartToolStatePendingStatus = "pending"
)

func (r PartUpdatePartToolStatePendingStatus) IsKnown() bool {
	switch r {
	case PartUpdatePartToolStatePendingStatusPending:
		return true
	}
	return false
}

// PartUpdatePartToolStateRunning represents the running variant of
// [PartUpdatePartToolStateUnion].
type PartUpdatePartToolStateRunning struct {
	Status   param.Field[PartUpdatePartToolStateRunningStatus] `json:"status,required"`
	Time     param.Field[PartUpdatePartToolStateRunningTime]   `json:"time,required"`
	Title    param.Field[string]                               `json:"title"`
	Metadata param.Field[map[string]any]                       `json:"metadata"`
}

func (r PartUpdatePartToolStateRunning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartToolStateRunning) implementsPartUpdatePartToolStateUnion() {}

type PartUpdatePartToolStateRunningStatus string

const (
	PartUpdatePartToolStateRunningStatusRunning PartUpdatePartToolStateRunningStatus = "running"
)

func (r PartUpdatePartToolStateRunningStatus) IsKnown() bool {
	switch r {
	case PartUpdatePartToolStateRunningStatusRunning:
		return true
	}
	return false
}

type PartUpdatePartToolStateRunningTime struct {
	Start param.Field[int64] `json:"start,required"`
}

func (r PartUpdatePartToolStateRunningTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PartUpdatePartToolStateCompleted represents the completed variant of
// [PartUpdatePartToolStateUnion].
type PartUpdatePartToolStateCompleted struct {
	Input    param.Field[map[string]any]                         `json:"input,required"`
	Metadata param.Field[map[string]any]                         `json:"metadata"`
	Output   param.Field[string]                                 `json:"output"`
	Status   param.Field[PartUpdatePartToolStateCompletedStatus] `json:"status,required"`
	Time     param.Field[PartUpdatePartToolStateCompletedTime]   `json:"time,required"`
	Title    param.Field[string]                                 `json:"title"`
}

func (r PartUpdatePartToolStateCompleted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartToolStateCompleted) implementsPartUpdatePartToolStateUnion() {}

type PartUpdatePartToolStateCompletedStatus string

const (
	PartUpdatePartToolStateCompletedStatusCompleted PartUpdatePartToolStateCompletedStatus = "completed"
)

func (r PartUpdatePartToolStateCompletedStatus) IsKnown() bool {
	switch r {
	case PartUpdatePartToolStateCompletedStatusCompleted:
		return true
	}
	return false
}

type PartUpdatePartToolStateCompletedTime struct {
	Start param.Field[int64] `json:"start,required"`
	End   param.Field[int64] `json:"end,required"`
}

func (r PartUpdatePartToolStateCompletedTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PartUpdatePartToolStateError represents the error variant of
// [PartUpdatePartToolStateUnion].
type PartUpdatePartToolStateError struct {
	Error    param.Field[string]                             `json:"error,required"`
	Input    param.Field[map[string]any]                     `json:"input,required"`
	Status   param.Field[PartUpdatePartToolStateErrorStatus] `json:"status,required"`
	Time     param.Field[PartUpdatePartToolStateErrorTime]   `json:"time,required"`
	Metadata param.Field[map[string]any]                     `json:"metadata"`
}

func (r PartUpdatePartToolStateError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartToolStateError) implementsPartUpdatePartToolStateUnion() {}

type PartUpdatePartToolStateErrorStatus string

const (
	PartUpdatePartToolStateErrorStatusError PartUpdatePartToolStateErrorStatus = "error"
)

func (r PartUpdatePartToolStateErrorStatus) IsKnown() bool {
	switch r {
	case PartUpdatePartToolStateErrorStatusError:
		return true
	}
	return false
}

type PartUpdatePartToolStateErrorTime struct {
	Start param.Field[int64] `json:"start,required"`
	End   param.Field[int64] `json:"end,required"`
}

func (r PartUpdatePartToolStateErrorTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PartUpdatePartToolStateUnion represents the OpenAPI ToolState anyOf union
// (per the [PartUpdatePartTool.State] field).
//
// Satisfied by [PartUpdatePartToolStatePending], [PartUpdatePartToolStateRunning],
// [PartUpdatePartToolStateCompleted] or [PartUpdatePartToolStateError].
type PartUpdatePartToolStateUnion interface {
	implementsPartUpdatePartToolStateUnion()
}

type PartUpdatePartToolType string

const (
	PartUpdatePartToolTypeTool PartUpdatePartToolType = "tool"
)

func (r PartUpdatePartToolType) IsKnown() bool {
	switch r {
	case PartUpdatePartToolTypeTool:
		return true
	}
	return false
}

type PartUpdatePartStepStart struct {
	ID        param.Field[string]                      `json:"id,required"`
	MessageID param.Field[string]                      `json:"messageID,required"`
	SessionID param.Field[string]                      `json:"sessionID,required"`
	Type      param.Field[PartUpdatePartStepStartType] `json:"type,required"`
	Snapshot  param.Field[string]                      `json:"snapshot"`
}

func (r PartUpdatePartStepStart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartStepStart) implementsPartUpdatePartUnion() {}

type PartUpdatePartStepStartType string

const (
	PartUpdatePartStepStartTypeStepStart PartUpdatePartStepStartType = "step-start"
)

func (r PartUpdatePartStepStartType) IsKnown() bool {
	switch r {
	case PartUpdatePartStepStartTypeStepStart:
		return true
	}
	return false
}

type PartUpdatePartStepFinish struct {
	Cost      param.Field[float64]                        `json:"cost,required"`
	ID        param.Field[string]                         `json:"id,required"`
	MessageID param.Field[string]                         `json:"messageID,required"`
	Reason    param.Field[string]                         `json:"reason,required"`
	SessionID param.Field[string]                         `json:"sessionID,required"`
	Tokens    param.Field[PartUpdatePartStepFinishTokens] `json:"tokens,required"`
	Type      param.Field[PartUpdatePartStepFinishType]   `json:"type,required"`
	Snapshot  param.Field[string]                         `json:"snapshot"`
}

func (r PartUpdatePartStepFinish) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartStepFinish) implementsPartUpdatePartUnion() {}

type PartUpdatePartStepFinishType string

const (
	PartUpdatePartStepFinishTypeStepFinish PartUpdatePartStepFinishType = "step-finish"
)

func (r PartUpdatePartStepFinishType) IsKnown() bool {
	switch r {
	case PartUpdatePartStepFinishTypeStepFinish:
		return true
	}
	return false
}

type PartUpdatePartStepFinishTokens struct {
	Cache     param.Field[PartUpdatePartStepFinishTokensCache] `json:"cache,required"`
	Input     param.Field[int64]                               `json:"input,required"`
	Output    param.Field[int64]                               `json:"output,required"`
	Reasoning param.Field[int64]                               `json:"reasoning,required"`
	Total     param.Field[int64]                               `json:"total"`
}

func (r PartUpdatePartStepFinishTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdatePartStepFinishTokensCache struct {
	Read  param.Field[int64] `json:"read,required"`
	Write param.Field[int64] `json:"write,required"`
}

func (r PartUpdatePartStepFinishTokensCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdatePartSnapshot struct {
	ID        param.Field[string]                     `json:"id,required"`
	MessageID param.Field[string]                     `json:"messageID,required"`
	SessionID param.Field[string]                     `json:"sessionID,required"`
	Snapshot  param.Field[string]                     `json:"snapshot,required"`
	Type      param.Field[PartUpdatePartSnapshotType] `json:"type,required"`
}

func (r PartUpdatePartSnapshot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartSnapshot) implementsPartUpdatePartUnion() {}

type PartUpdatePartSnapshotType string

const (
	PartUpdatePartSnapshotTypeSnapshot PartUpdatePartSnapshotType = "snapshot"
)

func (r PartUpdatePartSnapshotType) IsKnown() bool {
	switch r {
	case PartUpdatePartSnapshotTypeSnapshot:
		return true
	}
	return false
}

type PartUpdatePartPatch struct {
	Files     param.Field[[]string]                `json:"files,required"`
	Hash      param.Field[string]                  `json:"hash,required"`
	ID        param.Field[string]                  `json:"id,required"`
	MessageID param.Field[string]                  `json:"messageID,required"`
	SessionID param.Field[string]                  `json:"sessionID,required"`
	Type      param.Field[PartUpdatePartPatchType] `json:"type,required"`
}

func (r PartUpdatePartPatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartPatch) implementsPartUpdatePartUnion() {}

type PartUpdatePartPatchType string

const (
	PartUpdatePartPatchTypePatch PartUpdatePartPatchType = "patch"
)

func (r PartUpdatePartPatchType) IsKnown() bool {
	switch r {
	case PartUpdatePartPatchTypePatch:
		return true
	}
	return false
}

type PartUpdatePartAgent struct {
	ID        param.Field[string]                    `json:"id,required"`
	MessageID param.Field[string]                    `json:"messageID,required"`
	Name      param.Field[string]                    `json:"name,required"`
	SessionID param.Field[string]                    `json:"sessionID,required"`
	Type      param.Field[PartUpdatePartAgentType]   `json:"type,required"`
	Source    param.Field[PartUpdatePartAgentSource] `json:"source"`
}

func (r PartUpdatePartAgent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartAgent) implementsPartUpdatePartUnion() {}

type PartUpdatePartAgentType string

const (
	PartUpdatePartAgentTypeAgent PartUpdatePartAgentType = "agent"
)

func (r PartUpdatePartAgentType) IsKnown() bool {
	switch r {
	case PartUpdatePartAgentTypeAgent:
		return true
	}
	return false
}

type PartUpdatePartAgentSource struct {
	End   param.Field[int64]  `json:"end,required"`
	Start param.Field[int64]  `json:"start,required"`
	Value param.Field[string] `json:"value,required"`
}

func (r PartUpdatePartAgentSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdatePartRetry struct {
	Attempt   param.Field[int64]                    `json:"attempt,required"`
	ID        param.Field[string]                   `json:"id,required"`
	MessageID param.Field[string]                   `json:"messageID,required"`
	SessionID param.Field[string]                   `json:"sessionID,required"`
	Time      param.Field[PartUpdatePartRetryTime]  `json:"time,required"`
	Type      param.Field[PartUpdatePartRetryType]  `json:"type,required"`
	Error     param.Field[PartUpdatePartRetryError] `json:"error,required"`
}

func (r PartUpdatePartRetry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartRetry) implementsPartUpdatePartUnion() {}

// PartUpdatePartRetryError mirrors the OpenAPI APIError payload schema for the
// [PartUpdatePartRetry.Error] field — a `{"name":"APIError","data":{...}}` envelope.
type PartUpdatePartRetryError struct {
	Name param.Field[string]                       `json:"name,required"`
	Data param.Field[PartUpdatePartRetryErrorData] `json:"data,required"`
}

func (r PartUpdatePartRetryError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PartUpdatePartRetryErrorData mirrors OpenAPI APIErrorData.
type PartUpdatePartRetryErrorData struct {
	IsRetryable     param.Field[bool]              `json:"isRetryable,required"`
	Message         param.Field[string]            `json:"message,required"`
	Metadata        param.Field[map[string]string] `json:"metadata"`
	ResponseBody    param.Field[string]            `json:"responseBody"`
	ResponseHeaders param.Field[map[string]string] `json:"responseHeaders"`
	StatusCode      param.Field[int64]             `json:"statusCode"`
}

func (r PartUpdatePartRetryErrorData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdatePartRetryType string

const (
	PartUpdatePartRetryTypeRetry PartUpdatePartRetryType = "retry"
)

func (r PartUpdatePartRetryType) IsKnown() bool {
	switch r {
	case PartUpdatePartRetryTypeRetry:
		return true
	}
	return false
}

type PartUpdatePartRetryTime struct {
	Created param.Field[int64] `json:"created,required"`
}

func (r PartUpdatePartRetryTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdatePartCompaction struct {
	Auto        param.Field[bool]                         `json:"auto,required"`
	ID          param.Field[string]                       `json:"id,required"`
	MessageID   param.Field[string]                       `json:"messageID,required"`
	SessionID   param.Field[string]                       `json:"sessionID,required"`
	Type        param.Field[PartUpdatePartCompactionType] `json:"type,required"`
	Overflow    param.Field[bool]                         `json:"overflow"`
	TailStartID param.Field[string]                       `json:"tail_start_id"`
}

func (r PartUpdatePartCompaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdatePartCompaction) implementsPartUpdatePartUnion() {}

type PartUpdatePartCompactionType string

const (
	PartUpdatePartCompactionTypeCompaction PartUpdatePartCompactionType = "compaction"
)

func (r PartUpdatePartCompactionType) IsKnown() bool {
	switch r {
	case PartUpdatePartCompactionTypeCompaction:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[PartUpdatePartUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartText](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartSubtask](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartReasoning](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartFile](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartTool](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartStepStart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartStepFinish](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartSnapshot](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartPatch](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartAgent](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartRetry](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartCompaction](),
		},
	)
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[PartUpdatePartToolStateUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartToolStatePending](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartToolStateRunning](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartToolStateCompleted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdatePartToolStateError](),
		},
	)
}
