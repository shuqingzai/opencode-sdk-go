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
	// present (any of the 12 [PartUpdateParamsPartUnion] variants). When
	// unset, no body is sent. The `json:"-"` tag prevents the field from
	// being serialized as a `{ "Part": ... }` wrapper.
	Part param.Field[PartUpdateParamsPartUnion] `json:"-"`
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

// Satisfied by [PartUpdateParamsPartText], [PartUpdateParamsPartSubtask],
// [PartUpdateParamsPartReasoning], [PartUpdateParamsPartFile],
// [PartUpdateParamsPartTool], [PartUpdateParamsPartStepStart],
// [PartUpdateParamsPartStepFinish], [PartUpdateParamsPartSnapshot],
// [PartUpdateParamsPartPatch], [PartUpdateParamsPartAgent],
// [PartUpdateParamsPartRetry] or [PartUpdateParamsPartCompaction].
type PartUpdateParamsPartUnion interface {
	implementsPartUpdateParamsPartUnion()
}

type PartUpdateParamsPartText struct {
	ID        param.Field[string]                       `json:"id,required"`
	MessageID param.Field[string]                       `json:"messageID,required"`
	SessionID param.Field[string]                       `json:"sessionID,required"`
	Text      param.Field[string]                       `json:"text,required"`
	Type      param.Field[PartUpdateParamsPartTextType] `json:"type,required"`
	Ignored   param.Field[bool]                         `json:"ignored"`
	Metadata  param.Field[map[string]any]               `json:"metadata"`
	Synthetic param.Field[bool]                         `json:"synthetic"`
	Time      param.Field[PartUpdateParamsPartTextTime] `json:"time"`
}

func (r PartUpdateParamsPartText) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartText) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartTextType string

const (
	PartUpdateParamsPartTextTypeText PartUpdateParamsPartTextType = "text"
)

func (r PartUpdateParamsPartTextType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartTextTypeText:
		return true
	}
	return false
}

type PartUpdateParamsPartTextTime struct {
	Start param.Field[int64] `json:"start,required"`
	End   param.Field[int64] `json:"end"`
}

func (r PartUpdateParamsPartTextTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdateParamsPartSubtask struct {
	Agent       param.Field[string]                           `json:"agent,required"`
	Description param.Field[string]                           `json:"description,required"`
	ID          param.Field[string]                           `json:"id,required"`
	MessageID   param.Field[string]                           `json:"messageID,required"`
	Prompt      param.Field[string]                           `json:"prompt,required"`
	SessionID   param.Field[string]                           `json:"sessionID,required"`
	Type        param.Field[PartUpdateParamsPartSubtaskType]  `json:"type,required"`
	Command     param.Field[string]                           `json:"command"`
	Model       param.Field[PartUpdateParamsPartSubtaskModel] `json:"model"`
}

func (r PartUpdateParamsPartSubtask) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartSubtask) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartSubtaskType string

const (
	PartUpdateParamsPartSubtaskTypeSubtask PartUpdateParamsPartSubtaskType = "subtask"
)

func (r PartUpdateParamsPartSubtaskType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartSubtaskTypeSubtask:
		return true
	}
	return false
}

type PartUpdateParamsPartSubtaskModel struct {
	ModelID    param.Field[string] `json:"modelID,required"`
	ProviderID param.Field[string] `json:"providerID,required"`
}

func (r PartUpdateParamsPartSubtaskModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdateParamsPartReasoning struct {
	ID        param.Field[string]                            `json:"id,required"`
	MessageID param.Field[string]                            `json:"messageID,required"`
	SessionID param.Field[string]                            `json:"sessionID,required"`
	Text      param.Field[string]                            `json:"text,required"`
	Time      param.Field[PartUpdateParamsPartReasoningTime] `json:"time,required"`
	Type      param.Field[PartUpdateParamsPartReasoningType] `json:"type,required"`
	Metadata  param.Field[map[string]any]                    `json:"metadata"`
}

func (r PartUpdateParamsPartReasoning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartReasoning) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartReasoningType string

const (
	PartUpdateParamsPartReasoningTypeReasoning PartUpdateParamsPartReasoningType = "reasoning"
)

func (r PartUpdateParamsPartReasoningType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartReasoningTypeReasoning:
		return true
	}
	return false
}

type PartUpdateParamsPartReasoningTime struct {
	Start param.Field[int64] `json:"start,required"`
	End   param.Field[int64] `json:"end"`
}

func (r PartUpdateParamsPartReasoningTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdateParamsPartFile struct {
	ID        param.Field[string]                       `json:"id,required"`
	MessageID param.Field[string]                       `json:"messageID,required"`
	Mime      param.Field[string]                       `json:"mime,required"`
	SessionID param.Field[string]                       `json:"sessionID,required"`
	Type      param.Field[PartUpdateParamsPartFileType] `json:"type,required"`
	URL       param.Field[string]                       `json:"url,required"`
	Filename  param.Field[string]                       `json:"filename"`
	// Source of the file (e.g. symbol, resource, or inline file source).
	// Accepts [map[string]any] or an appropriate FilePartSource variant.
	Source param.Field[any] `json:"source"`
}

func (r PartUpdateParamsPartFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartFile) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartFileType string

const (
	PartUpdateParamsPartFileTypeFile PartUpdateParamsPartFileType = "file"
)

func (r PartUpdateParamsPartFileType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartFileTypeFile:
		return true
	}
	return false
}

type PartUpdateParamsPartTool struct {
	CallID    param.Field[string]                             `json:"callID,required"`
	ID        param.Field[string]                             `json:"id,required"`
	MessageID param.Field[string]                             `json:"messageID,required"`
	SessionID param.Field[string]                             `json:"sessionID,required"`
	State     param.Field[PartUpdateParamsPartToolStateUnion] `json:"state,required"`
	Tool      param.Field[string]                             `json:"tool,required"`
	Type      param.Field[PartUpdateParamsPartToolType]       `json:"type,required"`
	Metadata  param.Field[map[string]any]                     `json:"metadata"`
}

func (r PartUpdateParamsPartTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartTool) implementsPartUpdateParamsPartUnion() {}

// PartUpdateParamsPartToolStatePending represents the pending variant of
// [PartUpdateParamsPartToolStateUnion].
type PartUpdateParamsPartToolStatePending struct {
	Status param.Field[PartUpdateParamsPartToolStatePendingStatus] `json:"status,required"`
	Input  param.Field[map[string]any]                             `json:"input,required"`
	Raw    param.Field[string]                                     `json:"raw,required"`
}

func (r PartUpdateParamsPartToolStatePending) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartToolStatePending) implementsPartUpdateParamsPartToolStateUnion() {}

type PartUpdateParamsPartToolStatePendingStatus string

const (
	PartUpdateParamsPartToolStatePendingStatusPending PartUpdateParamsPartToolStatePendingStatus = "pending"
)

func (r PartUpdateParamsPartToolStatePendingStatus) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartToolStatePendingStatusPending:
		return true
	}
	return false
}

// PartUpdateParamsPartToolStateRunning represents the running variant of
// [PartUpdateParamsPartToolStateUnion].
type PartUpdateParamsPartToolStateRunning struct {
	Input    param.Field[map[string]any]                             `json:"input,required"`
	Status   param.Field[PartUpdateParamsPartToolStateRunningStatus] `json:"status,required"`
	Time     param.Field[PartUpdateParamsPartToolStateRunningTime]   `json:"time,required"`
	Title    param.Field[string]                                     `json:"title"`
	Metadata param.Field[map[string]any]                             `json:"metadata"`
}

func (r PartUpdateParamsPartToolStateRunning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartToolStateRunning) implementsPartUpdateParamsPartToolStateUnion() {}

type PartUpdateParamsPartToolStateRunningStatus string

const (
	PartUpdateParamsPartToolStateRunningStatusRunning PartUpdateParamsPartToolStateRunningStatus = "running"
)

func (r PartUpdateParamsPartToolStateRunningStatus) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartToolStateRunningStatusRunning:
		return true
	}
	return false
}

type PartUpdateParamsPartToolStateRunningTime struct {
	Start param.Field[int64] `json:"start,required"`
}

func (r PartUpdateParamsPartToolStateRunningTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PartUpdateParamsPartToolStateCompleted represents the completed variant of
// [PartUpdateParamsPartToolStateUnion].
type PartUpdateParamsPartToolStateCompleted struct {
	Input    param.Field[map[string]any]                               `json:"input,required"`
	Metadata param.Field[map[string]any]                               `json:"metadata,required"`
	Output   param.Field[string]                                       `json:"output,required"`
	Status   param.Field[PartUpdateParamsPartToolStateCompletedStatus] `json:"status,required"`
	Time     param.Field[PartUpdateParamsPartToolStateCompletedTime]   `json:"time,required"`
	Title    param.Field[string]                                       `json:"title,required"`
	// Optional list of file attachments.
	Attachments param.Field[[]PartUpdateParamsPartFile] `json:"attachments"`
}

func (r PartUpdateParamsPartToolStateCompleted) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartToolStateCompleted) implementsPartUpdateParamsPartToolStateUnion() {}

type PartUpdateParamsPartToolStateCompletedStatus string

const (
	PartUpdateParamsPartToolStateCompletedStatusCompleted PartUpdateParamsPartToolStateCompletedStatus = "completed"
)

func (r PartUpdateParamsPartToolStateCompletedStatus) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartToolStateCompletedStatusCompleted:
		return true
	}
	return false
}

type PartUpdateParamsPartToolStateCompletedTime struct {
	Start     param.Field[int64] `json:"start,required"`
	End       param.Field[int64] `json:"end,required"`
	Compacted param.Field[int64] `json:"compacted"`
}

func (r PartUpdateParamsPartToolStateCompletedTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PartUpdateParamsPartToolStateError represents the error variant of
// [PartUpdateParamsPartToolStateUnion].
type PartUpdateParamsPartToolStateError struct {
	Error    param.Field[string]                                   `json:"error,required"`
	Input    param.Field[map[string]any]                           `json:"input,required"`
	Status   param.Field[PartUpdateParamsPartToolStateErrorStatus] `json:"status,required"`
	Time     param.Field[PartUpdateParamsPartToolStateErrorTime]   `json:"time,required"`
	Metadata param.Field[map[string]any]                           `json:"metadata"`
}

func (r PartUpdateParamsPartToolStateError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartToolStateError) implementsPartUpdateParamsPartToolStateUnion() {}

type PartUpdateParamsPartToolStateErrorStatus string

const (
	PartUpdateParamsPartToolStateErrorStatusError PartUpdateParamsPartToolStateErrorStatus = "error"
)

func (r PartUpdateParamsPartToolStateErrorStatus) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartToolStateErrorStatusError:
		return true
	}
	return false
}

type PartUpdateParamsPartToolStateErrorTime struct {
	Start param.Field[int64] `json:"start,required"`
	End   param.Field[int64] `json:"end,required"`
}

func (r PartUpdateParamsPartToolStateErrorTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PartUpdateParamsPartToolStateUnion represents the OpenAPI ToolState anyOf
// union (per the [PartUpdateParamsPartTool.State] field).
//
// Satisfied by [PartUpdateParamsPartToolStatePending],
// [PartUpdateParamsPartToolStateRunning],
// [PartUpdateParamsPartToolStateCompleted] or
// [PartUpdateParamsPartToolStateError].
type PartUpdateParamsPartToolStateUnion interface {
	implementsPartUpdateParamsPartToolStateUnion()
}

type PartUpdateParamsPartToolType string

const (
	PartUpdateParamsPartToolTypeTool PartUpdateParamsPartToolType = "tool"
)

func (r PartUpdateParamsPartToolType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartToolTypeTool:
		return true
	}
	return false
}

type PartUpdateParamsPartStepStart struct {
	ID        param.Field[string]                            `json:"id,required"`
	MessageID param.Field[string]                            `json:"messageID,required"`
	SessionID param.Field[string]                            `json:"sessionID,required"`
	Type      param.Field[PartUpdateParamsPartStepStartType] `json:"type,required"`
	Snapshot  param.Field[string]                            `json:"snapshot"`
}

func (r PartUpdateParamsPartStepStart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartStepStart) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartStepStartType string

const (
	PartUpdateParamsPartStepStartTypeStepStart PartUpdateParamsPartStepStartType = "step-start"
)

func (r PartUpdateParamsPartStepStartType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartStepStartTypeStepStart:
		return true
	}
	return false
}

type PartUpdateParamsPartStepFinish struct {
	Cost      param.Field[float64]                              `json:"cost,required"`
	ID        param.Field[string]                               `json:"id,required"`
	MessageID param.Field[string]                               `json:"messageID,required"`
	Reason    param.Field[string]                               `json:"reason,required"`
	SessionID param.Field[string]                               `json:"sessionID,required"`
	Tokens    param.Field[PartUpdateParamsPartStepFinishTokens] `json:"tokens,required"`
	Type      param.Field[PartUpdateParamsPartStepFinishType]   `json:"type,required"`
	Snapshot  param.Field[string]                               `json:"snapshot"`
}

func (r PartUpdateParamsPartStepFinish) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartStepFinish) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartStepFinishType string

const (
	PartUpdateParamsPartStepFinishTypeStepFinish PartUpdateParamsPartStepFinishType = "step-finish"
)

func (r PartUpdateParamsPartStepFinishType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartStepFinishTypeStepFinish:
		return true
	}
	return false
}

type PartUpdateParamsPartStepFinishTokens struct {
	Cache     param.Field[PartUpdateParamsPartStepFinishTokensCache] `json:"cache,required"`
	Input     param.Field[int64]                                     `json:"input,required"`
	Output    param.Field[int64]                                     `json:"output,required"`
	Reasoning param.Field[int64]                                     `json:"reasoning,required"`
	Total     param.Field[int64]                                     `json:"total"`
}

func (r PartUpdateParamsPartStepFinishTokens) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdateParamsPartStepFinishTokensCache struct {
	Read  param.Field[int64] `json:"read,required"`
	Write param.Field[int64] `json:"write,required"`
}

func (r PartUpdateParamsPartStepFinishTokensCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdateParamsPartSnapshot struct {
	ID        param.Field[string]                           `json:"id,required"`
	MessageID param.Field[string]                           `json:"messageID,required"`
	SessionID param.Field[string]                           `json:"sessionID,required"`
	Snapshot  param.Field[string]                           `json:"snapshot,required"`
	Type      param.Field[PartUpdateParamsPartSnapshotType] `json:"type,required"`
}

func (r PartUpdateParamsPartSnapshot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartSnapshot) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartSnapshotType string

const (
	PartUpdateParamsPartSnapshotTypeSnapshot PartUpdateParamsPartSnapshotType = "snapshot"
)

func (r PartUpdateParamsPartSnapshotType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartSnapshotTypeSnapshot:
		return true
	}
	return false
}

type PartUpdateParamsPartPatch struct {
	Files     param.Field[[]string]                      `json:"files,required"`
	Hash      param.Field[string]                        `json:"hash,required"`
	ID        param.Field[string]                        `json:"id,required"`
	MessageID param.Field[string]                        `json:"messageID,required"`
	SessionID param.Field[string]                        `json:"sessionID,required"`
	Type      param.Field[PartUpdateParamsPartPatchType] `json:"type,required"`
}

func (r PartUpdateParamsPartPatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartPatch) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartPatchType string

const (
	PartUpdateParamsPartPatchTypePatch PartUpdateParamsPartPatchType = "patch"
)

func (r PartUpdateParamsPartPatchType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartPatchTypePatch:
		return true
	}
	return false
}

type PartUpdateParamsPartAgent struct {
	ID        param.Field[string]                          `json:"id,required"`
	MessageID param.Field[string]                          `json:"messageID,required"`
	Name      param.Field[string]                          `json:"name,required"`
	SessionID param.Field[string]                          `json:"sessionID,required"`
	Type      param.Field[PartUpdateParamsPartAgentType]   `json:"type,required"`
	Source    param.Field[PartUpdateParamsPartAgentSource] `json:"source"`
}

func (r PartUpdateParamsPartAgent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartAgent) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartAgentType string

const (
	PartUpdateParamsPartAgentTypeAgent PartUpdateParamsPartAgentType = "agent"
)

func (r PartUpdateParamsPartAgentType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartAgentTypeAgent:
		return true
	}
	return false
}

type PartUpdateParamsPartAgentSource struct {
	End   param.Field[int64]  `json:"end,required"`
	Start param.Field[int64]  `json:"start,required"`
	Value param.Field[string] `json:"value,required"`
}

func (r PartUpdateParamsPartAgentSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdateParamsPartRetry struct {
	Attempt   param.Field[int64]                          `json:"attempt,required"`
	ID        param.Field[string]                         `json:"id,required"`
	MessageID param.Field[string]                         `json:"messageID,required"`
	SessionID param.Field[string]                         `json:"sessionID,required"`
	Time      param.Field[PartUpdateParamsPartRetryTime]  `json:"time,required"`
	Type      param.Field[PartUpdateParamsPartRetryType]  `json:"type,required"`
	Error     param.Field[PartUpdateParamsPartRetryError] `json:"error,required"`
}

func (r PartUpdateParamsPartRetry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartRetry) implementsPartUpdateParamsPartUnion() {}

// PartUpdateParamsPartRetryError mirrors the OpenAPI APIError payload schema
// for the [PartUpdateParamsPartRetry.Error] field — a
// `{"name":"APIError","data":{...}}` envelope.
type PartUpdateParamsPartRetryError struct {
	Name param.Field[string]                             `json:"name,required"`
	Data param.Field[PartUpdateParamsPartRetryErrorData] `json:"data,required"`
}

func (r PartUpdateParamsPartRetryError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PartUpdateParamsPartRetryErrorData mirrors OpenAPI APIErrorData.
type PartUpdateParamsPartRetryErrorData struct {
	IsRetryable     param.Field[bool]              `json:"isRetryable,required"`
	Message         param.Field[string]            `json:"message,required"`
	Metadata        param.Field[map[string]string] `json:"metadata"`
	ResponseBody    param.Field[string]            `json:"responseBody"`
	ResponseHeaders param.Field[map[string]string] `json:"responseHeaders"`
	StatusCode      param.Field[int64]             `json:"statusCode"`
}

func (r PartUpdateParamsPartRetryErrorData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdateParamsPartRetryType string

const (
	PartUpdateParamsPartRetryTypeRetry PartUpdateParamsPartRetryType = "retry"
)

func (r PartUpdateParamsPartRetryType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartRetryTypeRetry:
		return true
	}
	return false
}

type PartUpdateParamsPartRetryTime struct {
	Created param.Field[int64] `json:"created,required"`
}

func (r PartUpdateParamsPartRetryTime) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PartUpdateParamsPartCompaction struct {
	Auto        param.Field[bool]                               `json:"auto,required"`
	ID          param.Field[string]                             `json:"id,required"`
	MessageID   param.Field[string]                             `json:"messageID,required"`
	SessionID   param.Field[string]                             `json:"sessionID,required"`
	Type        param.Field[PartUpdateParamsPartCompactionType] `json:"type,required"`
	Overflow    param.Field[bool]                               `json:"overflow"`
	TailStartID param.Field[string]                             `json:"tail_start_id"`
}

func (r PartUpdateParamsPartCompaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PartUpdateParamsPartCompaction) implementsPartUpdateParamsPartUnion() {}

type PartUpdateParamsPartCompactionType string

const (
	PartUpdateParamsPartCompactionTypeCompaction PartUpdateParamsPartCompactionType = "compaction"
)

func (r PartUpdateParamsPartCompactionType) IsKnown() bool {
	switch r {
	case PartUpdateParamsPartCompactionTypeCompaction:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[PartUpdateParamsPartUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartText](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartSubtask](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartReasoning](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartFile](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartTool](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartStepStart](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartStepFinish](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartSnapshot](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartPatch](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartAgent](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartRetry](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartCompaction](),
		},
	)
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[PartUpdateParamsPartToolStateUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartToolStatePending](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartToolStateRunning](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartToolStateCompleted](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PartUpdateParamsPartToolStateError](),
		},
	)
}
