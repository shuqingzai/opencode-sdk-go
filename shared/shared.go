// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/sst/opencode-sdk-go/internal/apijson"
)

type MessageAbortedError struct {
	Data MessageAbortedErrorData `json:"data,required"`
	Name MessageAbortedErrorName `json:"name,required"`
	JSON messageAbortedErrorJSON `json:"-"`
}

// messageAbortedErrorJSON contains the JSON metadata for the struct
// [MessageAbortedError]
type messageAbortedErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAbortedError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAbortedErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageAbortedError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r MessageAbortedError) ImplementsAssistantMessageError() {}

type MessageAbortedErrorData struct {
	Message string                      `json:"message,required"`
	JSON    messageAbortedErrorDataJSON `json:"-"`
}

// messageAbortedErrorDataJSON contains the JSON metadata for the struct
// [MessageAbortedErrorData]
type messageAbortedErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAbortedErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAbortedErrorDataJSON) RawJSON() string {
	return r.raw
}

type MessageAbortedErrorName string

const (
	MessageAbortedErrorNameMessageAbortedError MessageAbortedErrorName = "MessageAbortedError"
)

func (r MessageAbortedErrorName) IsKnown() bool {
	switch r {
	case MessageAbortedErrorNameMessageAbortedError:
		return true
	}
	return false
}

type ProviderAuthError struct {
	Data ProviderAuthErrorData `json:"data,required"`
	Name ProviderAuthErrorName `json:"name,required"`
	JSON providerAuthErrorJSON `json:"-"`
}

// providerAuthErrorJSON contains the JSON metadata for the struct
// [ProviderAuthError]
type providerAuthErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthErrorJSON) RawJSON() string {
	return r.raw
}

func (r ProviderAuthError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r ProviderAuthError) ImplementsAssistantMessageError() {}

type ProviderAuthErrorData struct {
	Message    string                    `json:"message,required"`
	ProviderID string                    `json:"providerID,required"`
	JSON       providerAuthErrorDataJSON `json:"-"`
}

// providerAuthErrorDataJSON contains the JSON metadata for the struct
// [ProviderAuthErrorData]
type providerAuthErrorDataJSON struct {
	Message     apijson.Field
	ProviderID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthErrorDataJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthErrorName string

const (
	ProviderAuthErrorNameProviderAuthError ProviderAuthErrorName = "ProviderAuthError"
)

func (r ProviderAuthErrorName) IsKnown() bool {
	switch r {
	case ProviderAuthErrorNameProviderAuthError:
		return true
	}
	return false
}

type ProviderAuthError1 struct {
	Data ProviderAuthError1Data `json:"data,required"`
	Name ProviderAuthError1Name `json:"name,required"`
	JSON providerAuthError1JSON `json:"-"`
}

// providerAuthError1JSON contains the JSON metadata for the struct
// [ProviderAuthError1]
type providerAuthError1JSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthError1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthError1JSON) RawJSON() string {
	return r.raw
}

type ProviderAuthError1Data struct {
	ProviderID string                     `json:"providerID"`
	Field      string                     `json:"field"`
	Message    string                     `json:"message"`
	Kind       string                     `json:"kind"`
	JSON       providerAuthError1DataJSON `json:"-"`
}

// providerAuthError1DataJSON contains the JSON metadata for the struct
// [ProviderAuthError1Data]
type providerAuthError1DataJSON struct {
	ProviderID  apijson.Field
	Field       apijson.Field
	Message     apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthError1Data) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthError1DataJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthError1Name string

const (
	ProviderAuthError1NameBadRequest                      ProviderAuthError1Name = "BadRequest"
	ProviderAuthError1NameProviderAuthOauthMissing        ProviderAuthError1Name = "ProviderAuthOauthMissing"
	ProviderAuthError1NameProviderAuthOauthCodeMissing    ProviderAuthError1Name = "ProviderAuthOauthCodeMissing"
	ProviderAuthError1NameProviderAuthOauthCallbackFailed ProviderAuthError1Name = "ProviderAuthOauthCallbackFailed"
	ProviderAuthError1NameProviderAuthValidationFailed    ProviderAuthError1Name = "ProviderAuthValidationFailed"
)

func (r ProviderAuthError1Name) IsKnown() bool {
	switch r {
	case ProviderAuthError1NameBadRequest, ProviderAuthError1NameProviderAuthOauthMissing, ProviderAuthError1NameProviderAuthOauthCodeMissing, ProviderAuthError1NameProviderAuthOauthCallbackFailed, ProviderAuthError1NameProviderAuthValidationFailed:
		return true
	}
	return false
}

type UnknownError struct {
	Data UnknownErrorData `json:"data,required"`
	Name UnknownErrorName `json:"name,required"`
	JSON unknownErrorJSON `json:"-"`
}

// unknownErrorJSON contains the JSON metadata for the struct [UnknownError]
type unknownErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnknownError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unknownErrorJSON) RawJSON() string {
	return r.raw
}

func (r UnknownError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r UnknownError) ImplementsAssistantMessageError() {}

type UnknownErrorData struct {
	Message string               `json:"message,required"`
	JSON    unknownErrorDataJSON `json:"-"`
}

// unknownErrorDataJSON contains the JSON metadata for the struct
// [UnknownErrorData]
type unknownErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnknownErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unknownErrorDataJSON) RawJSON() string {
	return r.raw
}

type UnknownErrorName string

const (
	UnknownErrorNameUnknownError UnknownErrorName = "UnknownError"
)

func (r UnknownErrorName) IsKnown() bool {
	switch r {
	case UnknownErrorNameUnknownError:
		return true
	}
	return false
}

type StructuredOutputError struct {
	Data StructuredOutputErrorData `json:"data,required"`
	Name StructuredOutputErrorName `json:"name,required"`
	JSON structuredOutputErrorJSON `json:"-"`
}

// structuredOutputErrorJSON contains the JSON metadata for the struct
// [StructuredOutputError]
type structuredOutputErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StructuredOutputError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r structuredOutputErrorJSON) RawJSON() string {
	return r.raw
}

func (r StructuredOutputError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r StructuredOutputError) ImplementsAssistantMessageError() {}

type StructuredOutputErrorData struct {
	Message string                        `json:"message,required"`
	Retries int64                         `json:"retries,required"`
	JSON    structuredOutputErrorDataJSON `json:"-"`
}

// structuredOutputErrorDataJSON contains the JSON metadata for the struct
// [StructuredOutputErrorData]
type structuredOutputErrorDataJSON struct {
	Message     apijson.Field
	Retries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StructuredOutputErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r structuredOutputErrorDataJSON) RawJSON() string {
	return r.raw
}

type StructuredOutputErrorName string

const (
	StructuredOutputErrorNameStructuredOutputError StructuredOutputErrorName = "StructuredOutputError"
)

func (r StructuredOutputErrorName) IsKnown() bool {
	switch r {
	case StructuredOutputErrorNameStructuredOutputError:
		return true
	}
	return false
}

type ContextOverflowError struct {
	Data ContextOverflowErrorData `json:"data,required"`
	Name ContextOverflowErrorName `json:"name,required"`
	JSON contextOverflowErrorJSON `json:"-"`
}

// contextOverflowErrorJSON contains the JSON metadata for the struct
// [ContextOverflowError]
type contextOverflowErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextOverflowError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextOverflowErrorJSON) RawJSON() string {
	return r.raw
}

func (r ContextOverflowError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r ContextOverflowError) ImplementsAssistantMessageError() {}

type ContextOverflowErrorData struct {
	Message      string                      `json:"message,required"`
	ResponseBody string                      `json:"responseBody"`
	JSON         contextOverflowErrorDataJSON `json:"-"`
}

// contextOverflowErrorDataJSON contains the JSON metadata for the struct
// [ContextOverflowErrorData]
type contextOverflowErrorDataJSON struct {
	Message      apijson.Field
	ResponseBody apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContextOverflowErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextOverflowErrorDataJSON) RawJSON() string {
	return r.raw
}

type ContextOverflowErrorName string

const (
	ContextOverflowErrorNameContextOverflowError ContextOverflowErrorName = "ContextOverflowError"
)

func (r ContextOverflowErrorName) IsKnown() bool {
	switch r {
	case ContextOverflowErrorNameContextOverflowError:
		return true
	}
	return false
}

type MessageOutputLengthError struct {
	Data interface{}                    `json:"data,required"`
	Name MessageOutputLengthErrorName   `json:"name,required"`
	JSON messageOutputLengthErrorJSON   `json:"-"`
}

// messageOutputLengthErrorJSON contains the JSON metadata for the struct
// [MessageOutputLengthError]
type messageOutputLengthErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageOutputLengthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageOutputLengthErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageOutputLengthError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r MessageOutputLengthError) ImplementsAssistantMessageError() {}

type MessageOutputLengthErrorName string

const (
	MessageOutputLengthErrorNameMessageOutputLengthError MessageOutputLengthErrorName = "MessageOutputLengthError"
)

func (r MessageOutputLengthErrorName) IsKnown() bool {
	switch r {
	case MessageOutputLengthErrorNameMessageOutputLengthError:
		return true
	}
	return false
}

type APIError struct {
	Data APIErrorData `json:"data,required"`
	Name APIErrorName `json:"name,required"`
	JSON apiErrorJSON `json:"-"`
}

// apiErrorJSON contains the JSON metadata for the struct [APIError]
type apiErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiErrorJSON) RawJSON() string {
	return r.raw
}

func (r APIError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r APIError) ImplementsAssistantMessageError() {}

type APIErrorData struct {
	IsRetryable     bool              `json:"isRetryable,required"`
	Message         string            `json:"message,required"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	ResponseBody    string            `json:"responseBody,omitempty"`
	ResponseHeaders map[string]string `json:"responseHeaders,omitempty"`
	StatusCode      float64           `json:"statusCode,omitempty"`
	JSON            apiErrorDataJSON  `json:"-"`
}

// apiErrorDataJSON contains the JSON metadata for the struct [APIErrorData]
type apiErrorDataJSON struct {
	IsRetryable     apijson.Field
	Message         apijson.Field
	Metadata        apijson.Field
	ResponseBody    apijson.Field
	ResponseHeaders apijson.Field
	StatusCode      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *APIErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiErrorDataJSON) RawJSON() string {
	return r.raw
}

type APIErrorName string

const (
	APIErrorNameAPIError APIErrorName = "APIError"
)

func (r APIErrorName) IsKnown() bool {
	switch r {
	case APIErrorNameAPIError:
		return true
	}
	return false
}

type VcsApplyError struct {
	Data VcsApplyErrorData `json:"data,required"`
	Name VcsApplyErrorName `json:"name,required"`
	JSON vcsApplyErrorJSON `json:"-"`
}

// vcsApplyErrorJSON contains the JSON metadata for the struct [VcsApplyError]
type vcsApplyErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VcsApplyError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vcsApplyErrorJSON) RawJSON() string {
	return r.raw
}

type VcsApplyErrorData struct {
	Message string                `json:"message,required"`
	Reason  VcsApplyErrorReason   `json:"reason,required"`
	JSON    vcsApplyErrorDataJSON `json:"-"`
}

// vcsApplyErrorDataJSON contains the JSON metadata for the struct [VcsApplyErrorData]
type vcsApplyErrorDataJSON struct {
	Message     apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VcsApplyErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vcsApplyErrorDataJSON) RawJSON() string {
	return r.raw
}

type VcsApplyErrorName string

const (
	VcsApplyErrorNameVcsApplyError VcsApplyErrorName = "VcsApplyError"
)

func (r VcsApplyErrorName) IsKnown() bool {
	switch r {
	case VcsApplyErrorNameVcsApplyError:
		return true
	}
	return false
}

type VcsApplyErrorReason string

const (
	VcsApplyErrorReasonNonGit  VcsApplyErrorReason = "non-git"
	VcsApplyErrorReasonNotClean VcsApplyErrorReason = "not-clean"
)

func (r VcsApplyErrorReason) IsKnown() bool {
	switch r {
	case VcsApplyErrorReasonNonGit, VcsApplyErrorReasonNotClean:
		return true
	}
	return false
}

type WorkspaceWarpError struct {
	Data WorkspaceWarpErrorData `json:"data,required"`
	Name WorkspaceWarpErrorName `json:"name,required"`
	JSON workspaceWarpErrorJSON `json:"-"`
}

// workspaceWarpErrorJSON contains the JSON metadata for the struct [WorkspaceWarpError]
type workspaceWarpErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceWarpError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceWarpErrorJSON) RawJSON() string {
	return r.raw
}

type WorkspaceWarpErrorData struct {
	Message string                     `json:"message,required"`
	JSON    workspaceWarpErrorDataJSON `json:"-"`
}

// workspaceWarpErrorDataJSON contains the JSON metadata for the struct [WorkspaceWarpErrorData]
type workspaceWarpErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceWarpErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceWarpErrorDataJSON) RawJSON() string {
	return r.raw
}

type WorkspaceWarpErrorName string

const (
	WorkspaceWarpErrorNameWorkspaceWarpError WorkspaceWarpErrorName = "WorkspaceWarpError"
)

func (r WorkspaceWarpErrorName) IsKnown() bool {
	switch r {
	case WorkspaceWarpErrorNameWorkspaceWarpError:
		return true
	}
	return false
}

type NotFoundError struct {
	Data NotFoundErrorData `json:"data,required"`
	Name NotFoundErrorName `json:"name,required"`
	JSON notFoundErrorJSON `json:"-"`
}

// notFoundErrorJSON contains the JSON metadata for the struct [NotFoundError]
type notFoundErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notFoundErrorJSON) RawJSON() string {
	return r.raw
}

type NotFoundErrorData struct {
	Message string                `json:"message,required"`
	JSON    notFoundErrorDataJSON `json:"-"`
}

// notFoundErrorDataJSON contains the JSON metadata for the struct [NotFoundErrorData]
type notFoundErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotFoundErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notFoundErrorDataJSON) RawJSON() string {
	return r.raw
}

type NotFoundErrorName string

const (
	NotFoundErrorNameNotFoundError NotFoundErrorName = "NotFoundError"
)

func (r NotFoundErrorName) IsKnown() bool {
	switch r {
	case NotFoundErrorNameNotFoundError:
		return true
	}
	return false
}

type BadRequestError struct {
	Data BadRequestErrorData `json:"data,required"`
	Name BadRequestErrorName `json:"name,required"`
	JSON badRequestErrorJSON `json:"-"`
}

// badRequestErrorJSON contains the JSON metadata for the struct [BadRequestError]
type badRequestErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BadRequestError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r badRequestErrorJSON) RawJSON() string {
	return r.raw
}

type BadRequestErrorData struct {
	Message string                `json:"message,required"`
	Kind    BadRequestErrorKind   `json:"kind"`
	JSON    badRequestErrorDataJSON `json:"-"`
}

// badRequestErrorDataJSON contains the JSON metadata for the struct [BadRequestErrorData]
type badRequestErrorDataJSON struct {
	Message     apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BadRequestErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r badRequestErrorDataJSON) RawJSON() string {
	return r.raw
}

type BadRequestErrorName string

const (
	BadRequestErrorNameBadRequest BadRequestErrorName = "BadRequest"
)

func (r BadRequestErrorName) IsKnown() bool {
	switch r {
	case BadRequestErrorNameBadRequest:
		return true
	}
	return false
}

type BadRequestErrorKind string

const (
	BadRequestErrorKindParams  BadRequestErrorKind = "Params"
	BadRequestErrorKindHeaders BadRequestErrorKind = "Headers"
	BadRequestErrorKindQuery   BadRequestErrorKind = "Query"
	BadRequestErrorKindBody    BadRequestErrorKind = "Body"
	BadRequestErrorKindPayload BadRequestErrorKind = "Payload"
)

func (r BadRequestErrorKind) IsKnown() bool {
	switch r {
	case BadRequestErrorKindParams, BadRequestErrorKindHeaders, BadRequestErrorKindQuery, BadRequestErrorKindBody, BadRequestErrorKindPayload:
		return true
	}
	return false
}

type WorktreeError struct {
	Data WorktreeErrorData `json:"data,required"`
	Name WorktreeErrorName `json:"name,required"`
	JSON worktreeErrorJSON `json:"-"`
}

// worktreeErrorJSON contains the JSON metadata for the struct [WorktreeError]
type worktreeErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorktreeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r worktreeErrorJSON) RawJSON() string {
	return r.raw
}

type WorktreeErrorData struct {
	Message string               `json:"message,required"`
	JSON    worktreeErrorDataJSON `json:"-"`
}

// worktreeErrorDataJSON contains the JSON metadata for the struct [WorktreeErrorData]
type worktreeErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorktreeErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r worktreeErrorDataJSON) RawJSON() string {
	return r.raw
}

type WorktreeErrorName string

const (
	WorktreeErrorNameWorktreeNotGitError               WorktreeErrorName = "WorktreeNotGitError"
	WorktreeErrorNameWorktreeNameGenerationFailedError  WorktreeErrorName = "WorktreeNameGenerationFailedError"
	WorktreeErrorNameWorktreeCreateFailedError          WorktreeErrorName = "WorktreeCreateFailedError"
	WorktreeErrorNameWorktreeStartCommandFailedError    WorktreeErrorName = "WorktreeStartCommandFailedError"
	WorktreeErrorNameWorktreeRemoveFailedError          WorktreeErrorName = "WorktreeRemoveFailedError"
	WorktreeErrorNameWorktreeResetFailedError           WorktreeErrorName = "WorktreeResetFailedError"
	WorktreeErrorNameWorktreeListFailedError            WorktreeErrorName = "WorktreeListFailedError"
)

func (r WorktreeErrorName) IsKnown() bool {
	switch r {
	case WorktreeErrorNameWorktreeNotGitError, WorktreeErrorNameWorktreeNameGenerationFailedError, WorktreeErrorNameWorktreeCreateFailedError, WorktreeErrorNameWorktreeStartCommandFailedError, WorktreeErrorNameWorktreeRemoveFailedError, WorktreeErrorNameWorktreeResetFailedError, WorktreeErrorNameWorktreeListFailedError:
		return true
	}
	return false
}
