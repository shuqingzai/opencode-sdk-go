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

func (r MessageAbortedError) ImplementsV2EventSessionErrorDataError() {}

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

func (r ProviderAuthError) ImplementsV2EventSessionErrorDataError() {}

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

func (r UnknownError) ImplementsV2EventSessionErrorDataError() {}

func (r UnknownError) ImplementsAssistantMessageError() {}

type UnknownErrorData struct {
	Message string               `json:"message,required"`
	Ref     string               `json:"ref"`
	JSON    unknownErrorDataJSON `json:"-"`
}

// unknownErrorDataJSON contains the JSON metadata for the struct
// [UnknownErrorData]
type unknownErrorDataJSON struct {
	Message     apijson.Field
	Ref         apijson.Field
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

func (r StructuredOutputError) ImplementsV2EventSessionErrorDataError() {}

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

func (r ContextOverflowError) ImplementsV2EventSessionErrorDataError() {}

func (r ContextOverflowError) ImplementsAssistantMessageError() {}

type ContextOverflowErrorData struct {
	Message      string                       `json:"message,required"`
	ResponseBody string                       `json:"responseBody"`
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
	// Per OpenAPI `MessageOutputLengthError.data` is an object with no declared
	// properties, so the server may return any JSON object here.
	// This field can have the runtime type of [map[string]any].
	Data any                          `json:"data,required"`
	Name MessageOutputLengthErrorName `json:"name,required"`
	JSON messageOutputLengthErrorJSON `json:"-"`
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

func (r MessageOutputLengthError) ImplementsV2EventSessionErrorDataError() {}

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

func (r APIError) ImplementsV2EventSessionErrorDataError() {}

func (r APIError) ImplementsAssistantMessageError() {}

type APIErrorData struct {
	IsRetryable     bool              `json:"isRetryable,required"`
	Message         string            `json:"message,required"`
	Metadata        map[string]string `json:"metadata"`
	ResponseBody    string            `json:"responseBody"`
	ResponseHeaders map[string]string `json:"responseHeaders"`
	StatusCode      int64             `json:"statusCode"`
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
	VcsApplyErrorReasonNonGit   VcsApplyErrorReason = "non-git"
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

type WorkspaceCreateError struct {
	Data WorkspaceCreateErrorData `json:"data,required"`
	Name WorkspaceCreateErrorName `json:"name,required"`
	JSON workspaceCreateErrorJSON `json:"-"`
}

// workspaceCreateErrorJSON contains the JSON metadata for the struct [WorkspaceCreateError]
type workspaceCreateErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceCreateError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceCreateErrorJSON) RawJSON() string {
	return r.raw
}

type WorkspaceCreateErrorData struct {
	Message string                       `json:"message,required"`
	JSON    workspaceCreateErrorDataJSON `json:"-"`
}

// workspaceCreateErrorDataJSON contains the JSON metadata for the struct [WorkspaceCreateErrorData]
type workspaceCreateErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceCreateErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceCreateErrorDataJSON) RawJSON() string {
	return r.raw
}

type WorkspaceCreateErrorName string

const (
	WorkspaceCreateErrorNameWorkspaceCreateError WorkspaceCreateErrorName = "WorkspaceCreateError"
)

func (r WorkspaceCreateErrorName) IsKnown() bool {
	switch r {
	case WorkspaceCreateErrorNameWorkspaceCreateError:
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
	Message string                  `json:"message,required"`
	Kind    BadRequestErrorKind     `json:"kind"`
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
	Message string                `json:"message,required"`
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
	WorktreeErrorNameWorktreeNameGenerationFailedError WorktreeErrorName = "WorktreeNameGenerationFailedError"
	WorktreeErrorNameWorktreeCreateFailedError         WorktreeErrorName = "WorktreeCreateFailedError"
	WorktreeErrorNameWorktreeStartCommandFailedError   WorktreeErrorName = "WorktreeStartCommandFailedError"
	WorktreeErrorNameWorktreeRemoveFailedError         WorktreeErrorName = "WorktreeRemoveFailedError"
	WorktreeErrorNameWorktreeResetFailedError          WorktreeErrorName = "WorktreeResetFailedError"
	WorktreeErrorNameWorktreeListFailedError           WorktreeErrorName = "WorktreeListFailedError"
)

func (r WorktreeErrorName) IsKnown() bool {
	switch r {
	case WorktreeErrorNameWorktreeNotGitError, WorktreeErrorNameWorktreeNameGenerationFailedError, WorktreeErrorNameWorktreeCreateFailedError, WorktreeErrorNameWorktreeStartCommandFailedError, WorktreeErrorNameWorktreeRemoveFailedError, WorktreeErrorNameWorktreeResetFailedError, WorktreeErrorNameWorktreeListFailedError:
		return true
	}
	return false
}

// === _tag-based errors (HTTP infrastructure errors) ===

type InvalidCursorError struct {
	Tag     InvalidCursorErrorTag  `json:"_tag,required"`
	Message string                 `json:"message,required"`
	JSON    invalidCursorErrorJSON `json:"-"`
}

type invalidCursorErrorJSON struct {
	Tag         apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvalidCursorError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invalidCursorErrorJSON) RawJSON() string {
	return r.raw
}

type InvalidCursorErrorTag string

const (
	InvalidCursorErrorTagInvalidCursorError InvalidCursorErrorTag = "InvalidCursorError"
)

func (r InvalidCursorErrorTag) IsKnown() bool {
	switch r {
	case InvalidCursorErrorTagInvalidCursorError:
		return true
	}
	return false
}

type InvalidRequestError struct {
	Tag     InvalidRequestErrorTag  `json:"_tag,required"`
	Message string                  `json:"message,required"`
	Kind    string                  `json:"kind"`
	Field   string                  `json:"field"`
	JSON    invalidRequestErrorJSON `json:"-"`
}

type invalidRequestErrorJSON struct {
	Tag         apijson.Field
	Message     apijson.Field
	Kind        apijson.Field
	Field       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvalidRequestError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invalidRequestErrorJSON) RawJSON() string {
	return r.raw
}

type InvalidRequestErrorTag string

const (
	InvalidRequestErrorTagInvalidRequestError InvalidRequestErrorTag = "InvalidRequestError"
)

func (r InvalidRequestErrorTag) IsKnown() bool {
	switch r {
	case InvalidRequestErrorTagInvalidRequestError:
		return true
	}
	return false
}

type McpUnsupportedOAuthError struct {
	Error string                       `json:"error,required"`
	JSON  mcpUnsupportedOAuthErrorJSON `json:"-"`
}

type mcpUnsupportedOAuthErrorJSON struct {
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpUnsupportedOAuthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpUnsupportedOAuthErrorJSON) RawJSON() string {
	return r.raw
}

type ProviderNotFoundError struct {
	Tag        ProviderNotFoundErrorTag  `json:"_tag,required"`
	ProviderID string                    `json:"providerID,required"`
	Message    string                    `json:"message,required"`
	JSON       providerNotFoundErrorJSON `json:"-"`
}

type providerNotFoundErrorJSON struct {
	Tag         apijson.Field
	ProviderID  apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

type ProviderNotFoundErrorTag string

const (
	ProviderNotFoundErrorTagProviderNotFoundError ProviderNotFoundErrorTag = "ProviderNotFoundError"
)

func (r ProviderNotFoundErrorTag) IsKnown() bool {
	switch r {
	case ProviderNotFoundErrorTagProviderNotFoundError:
		return true
	}
	return false
}

type ServiceUnavailableError struct {
	Tag     ServiceUnavailableErrorTag  `json:"_tag,required"`
	Message string                      `json:"message,required"`
	Service string                      `json:"service"`
	JSON    serviceUnavailableErrorJSON `json:"-"`
}

type serviceUnavailableErrorJSON struct {
	Tag         apijson.Field
	Message     apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceUnavailableError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceUnavailableErrorJSON) RawJSON() string {
	return r.raw
}

type ServiceUnavailableErrorTag string

const (
	ServiceUnavailableErrorTagServiceUnavailableError ServiceUnavailableErrorTag = "ServiceUnavailableError"
)

func (r ServiceUnavailableErrorTag) IsKnown() bool {
	switch r {
	case ServiceUnavailableErrorTagServiceUnavailableError:
		return true
	}
	return false
}

type SessionNotFoundError struct {
	Tag       SessionNotFoundErrorTag  `json:"_tag,required"`
	SessionID string                   `json:"sessionID,required"`
	Message   string                   `json:"message,required"`
	JSON      sessionNotFoundErrorJSON `json:"-"`
}

type sessionNotFoundErrorJSON struct {
	Tag         apijson.Field
	SessionID   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

type SessionNotFoundErrorTag string

const (
	SessionNotFoundErrorTagSessionNotFoundError SessionNotFoundErrorTag = "SessionNotFoundError"
)

func (r SessionNotFoundErrorTag) IsKnown() bool {
	switch r {
	case SessionNotFoundErrorTagSessionNotFoundError:
		return true
	}
	return false
}

type UnknownError1 struct {
	Tag     UnknownError1Tag  `json:"_tag,required"`
	Message string            `json:"message,required"`
	Ref     string            `json:"ref"`
	JSON    unknownError1JSON `json:"-"`
}

type unknownError1JSON struct {
	Tag         apijson.Field
	Message     apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnknownError1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unknownError1JSON) RawJSON() string {
	return r.raw
}

type UnknownError1Tag string

const (
	UnknownError1TagUnknownError UnknownError1Tag = "UnknownError"
)

func (r UnknownError1Tag) IsKnown() bool {
	switch r {
	case UnknownError1TagUnknownError:
		return true
	}
	return false
}

type UnauthorizedError struct {
	Tag     UnauthorizedErrorTag  `json:"_tag,required"`
	Message string                `json:"message,required"`
	JSON    unauthorizedErrorJSON `json:"-"`
}

type unauthorizedErrorJSON struct {
	Tag         apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnauthorizedError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unauthorizedErrorJSON) RawJSON() string {
	return r.raw
}

type UnauthorizedErrorTag string

const (
	UnauthorizedErrorTagUnauthorizedError UnauthorizedErrorTag = "UnauthorizedError"
)

func (r UnauthorizedErrorTag) IsKnown() bool {
	switch r {
	case UnauthorizedErrorTagUnauthorizedError:
		return true
	}
	return false
}

type EffectHttpApiErrorBadRequest struct {
	Tag  EffectHttpApiErrorBadRequestTag  `json:"_tag,required"`
	JSON effectHttpApiErrorBadRequestJSON `json:"-"`
}

type effectHttpApiErrorBadRequestJSON struct {
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EffectHttpApiErrorBadRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r effectHttpApiErrorBadRequestJSON) RawJSON() string {
	return r.raw
}

type EffectHttpApiErrorBadRequestTag string

const (
	EffectHttpApiErrorBadRequestTagBadRequest EffectHttpApiErrorBadRequestTag = "BadRequest"
)

func (r EffectHttpApiErrorBadRequestTag) IsKnown() bool {
	switch r {
	case EffectHttpApiErrorBadRequestTagBadRequest:
		return true
	}
	return false
}

// === v1.15.10 — 新的标签式 HTTP 错误类型 ===

type McpServerNotFoundError struct {
	Tag     McpServerNotFoundErrorTag  `json:"_tag,required"`
	Name    string                     `json:"name,required"`
	Message string                     `json:"message,required"`
	JSON    mcpServerNotFoundErrorJSON `json:"-"`
}

type mcpServerNotFoundErrorJSON struct {
	Tag         apijson.Field
	Name        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpServerNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpServerNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

type McpServerNotFoundErrorTag string

const (
	McpServerNotFoundErrorTagMcpServerNotFoundError McpServerNotFoundErrorTag = "McpServerNotFoundError"
)

func (r McpServerNotFoundErrorTag) IsKnown() bool {
	switch r {
	case McpServerNotFoundErrorTagMcpServerNotFoundError:
		return true
	}
	return false
}

type ProjectNotFoundError struct {
	Tag       ProjectNotFoundErrorTag  `json:"_tag,required"`
	ProjectID string                   `json:"projectID,required"`
	Message   string                   `json:"message,required"`
	JSON      projectNotFoundErrorJSON `json:"-"`
}

type projectNotFoundErrorJSON struct {
	Tag         apijson.Field
	ProjectID   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

type ProjectNotFoundErrorTag string

const (
	ProjectNotFoundErrorTagProjectNotFoundError ProjectNotFoundErrorTag = "ProjectNotFoundError"
)

func (r ProjectNotFoundErrorTag) IsKnown() bool {
	switch r {
	case ProjectNotFoundErrorTagProjectNotFoundError:
		return true
	}
	return false
}

type PtyNotFoundError struct {
	Tag     PtyNotFoundErrorTag  `json:"_tag,required"`
	PtyID   string               `json:"ptyID,required"`
	Message string               `json:"message,required"`
	JSON    ptyNotFoundErrorJSON `json:"-"`
}

type ptyNotFoundErrorJSON struct {
	Tag         apijson.Field
	PtyID       apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

type PtyNotFoundErrorTag string

const (
	PtyNotFoundErrorTagPtyNotFoundError PtyNotFoundErrorTag = "PtyNotFoundError"
)

func (r PtyNotFoundErrorTag) IsKnown() bool {
	switch r {
	case PtyNotFoundErrorTagPtyNotFoundError:
		return true
	}
	return false
}

type PtyForbiddenError struct {
	Tag     PtyForbiddenErrorTag  `json:"_tag,required"`
	Message string                `json:"message,required"`
	JSON    ptyForbiddenErrorJSON `json:"-"`
}

type ptyForbiddenErrorJSON struct {
	Tag         apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyForbiddenError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyForbiddenErrorJSON) RawJSON() string {
	return r.raw
}

type PtyForbiddenErrorTag string

const (
	PtyForbiddenErrorTagPtyForbiddenError PtyForbiddenErrorTag = "PtyForbiddenError"
)

func (r PtyForbiddenErrorTag) IsKnown() bool {
	switch r {
	case PtyForbiddenErrorTagPtyForbiddenError:
		return true
	}
	return false
}

type QuestionNotFoundError struct {
	Tag       QuestionNotFoundErrorTag  `json:"_tag,required"`
	RequestID string                    `json:"requestID,required"`
	Message   string                    `json:"message,required"`
	JSON      questionNotFoundErrorJSON `json:"-"`
}

type questionNotFoundErrorJSON struct {
	Tag         apijson.Field
	RequestID   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

type QuestionNotFoundErrorTag string

const (
	QuestionNotFoundErrorTagQuestionNotFoundError QuestionNotFoundErrorTag = "QuestionNotFoundError"
)

func (r QuestionNotFoundErrorTag) IsKnown() bool {
	switch r {
	case QuestionNotFoundErrorTagQuestionNotFoundError:
		return true
	}
	return false
}

type PermissionNotFoundError struct {
	Tag       PermissionNotFoundErrorTag  `json:"_tag,required"`
	RequestID string                      `json:"requestID,required"`
	Message   string                      `json:"message,required"`
	JSON      permissionNotFoundErrorJSON `json:"-"`
}

type permissionNotFoundErrorJSON struct {
	Tag         apijson.Field
	RequestID   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

type PermissionNotFoundErrorTag string

const (
	PermissionNotFoundErrorTagPermissionNotFoundError PermissionNotFoundErrorTag = "PermissionNotFoundError"
)

func (r PermissionNotFoundErrorTag) IsKnown() bool {
	switch r {
	case PermissionNotFoundErrorTagPermissionNotFoundError:
		return true
	}
	return false
}

type SessionBusyError struct {
	Tag       SessionBusyErrorTag  `json:"_tag,required"`
	SessionID string               `json:"sessionID,required"`
	Message   string               `json:"message,required"`
	JSON      sessionBusyErrorJSON `json:"-"`
}

type sessionBusyErrorJSON struct {
	Tag         apijson.Field
	SessionID   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SessionBusyError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionBusyErrorJSON) RawJSON() string {
	return r.raw
}

type SessionBusyErrorTag string

const (
	SessionBusyErrorTagSessionBusyError SessionBusyErrorTag = "SessionBusyError"
)

func (r SessionBusyErrorTag) IsKnown() bool {
	switch r {
	case SessionBusyErrorTagSessionBusyError:
		return true
	}
	return false
}

type EffectHttpApiErrorForbidden struct {
	Tag  EffectHttpApiErrorForbiddenTag  `json:"_tag,required"`
	JSON effectHttpApiErrorForbiddenJSON `json:"-"`
}

type effectHttpApiErrorForbiddenJSON struct {
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EffectHttpApiErrorForbidden) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r effectHttpApiErrorForbiddenJSON) RawJSON() string {
	return r.raw
}

type EffectHttpApiErrorForbiddenTag string

const (
	EffectHttpApiErrorForbiddenTagForbidden EffectHttpApiErrorForbiddenTag = "Forbidden"
)

func (r EffectHttpApiErrorForbiddenTag) IsKnown() bool {
	switch r {
	case EffectHttpApiErrorForbiddenTagForbidden:
		return true
	}
	return false
}

type EffectHttpApiErrorInternalServerError struct {
	Tag  EffectHttpApiErrorInternalServerErrorTag  `json:"_tag,required"`
	JSON effectHttpApiErrorInternalServerErrorJSON `json:"-"`
}

type effectHttpApiErrorInternalServerErrorJSON struct {
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EffectHttpApiErrorInternalServerError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r effectHttpApiErrorInternalServerErrorJSON) RawJSON() string {
	return r.raw
}

type EffectHttpApiErrorInternalServerErrorTag string

const (
	EffectHttpApiErrorInternalServerErrorTagInternalServerError EffectHttpApiErrorInternalServerErrorTag = "InternalServerError"
)

func (r EffectHttpApiErrorInternalServerErrorTag) IsKnown() bool {
	switch r {
	case EffectHttpApiErrorInternalServerErrorTagInternalServerError:
		return true
	}
	return false
}

type ConflictError struct {
	Tag      ConflictErrorTag  `json:"_tag,required"`
	Message  string            `json:"message,required"`
	Resource string            `json:"resource"`
	JSON     conflictErrorJSON `json:"-"`
}

type conflictErrorJSON struct {
	Tag         apijson.Field
	Message     apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConflictError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r conflictErrorJSON) RawJSON() string {
	return r.raw
}

type ConflictErrorTag string

const (
	ConflictErrorTagConflictError ConflictErrorTag = "ConflictError"
)

func (r ConflictErrorTag) IsKnown() bool {
	switch r {
	case ConflictErrorTagConflictError:
		return true
	}
	return false
}

type ContentFilterError struct {
	Data ContentFilterErrorData `json:"data,required"`
	Name ContentFilterErrorName `json:"name,required"`
	JSON contentFilterErrorJSON `json:"-"`
}

type contentFilterErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentFilterError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentFilterErrorJSON) RawJSON() string {
	return r.raw
}

func (r ContentFilterError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r ContentFilterError) ImplementsV2EventSessionErrorDataError() {}

func (r ContentFilterError) ImplementsAssistantMessageError() {}

type ContentFilterErrorData struct {
	Message string                     `json:"message,required"`
	JSON    contentFilterErrorDataJSON `json:"-"`
}

type contentFilterErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentFilterErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentFilterErrorDataJSON) RawJSON() string {
	return r.raw
}

type ContentFilterErrorName string

const (
	ContentFilterErrorNameContentFilterError ContentFilterErrorName = "ContentFilterError"
)

func (r ContentFilterErrorName) IsKnown() bool {
	switch r {
	case ContentFilterErrorNameContentFilterError:
		return true
	}
	return false
}

type ForbiddenError struct {
	Tag     ForbiddenErrorTag  `json:"_tag,required"`
	Message string             `json:"message,required"`
	JSON    forbiddenErrorJSON `json:"-"`
}

type forbiddenErrorJSON struct {
	Tag         apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForbiddenError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r forbiddenErrorJSON) RawJSON() string {
	return r.raw
}

type ForbiddenErrorTag string

const (
	ForbiddenErrorTagForbiddenError ForbiddenErrorTag = "ForbiddenError"
)

func (r ForbiddenErrorTag) IsKnown() bool {
	switch r {
	case ForbiddenErrorTagForbiddenError:
		return true
	}
	return false
}

type MessageNotFoundError struct {
	Tag       MessageNotFoundErrorTag  `json:"_tag,required"`
	SessionID string                   `json:"sessionID,required"`
	MessageID string                   `json:"messageID,required"`
	Message   string                   `json:"message,required"`
	JSON      messageNotFoundErrorJSON `json:"-"`
}

type messageNotFoundErrorJSON struct {
	Tag         apijson.Field
	SessionID   apijson.Field
	MessageID   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageNotFoundError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageNotFoundErrorJSON) RawJSON() string {
	return r.raw
}

type MessageNotFoundErrorTag string

const (
	MessageNotFoundErrorTagMessageNotFoundError MessageNotFoundErrorTag = "MessageNotFoundError"
)

func (r MessageNotFoundErrorTag) IsKnown() bool {
	switch r {
	case MessageNotFoundErrorTagMessageNotFoundError:
		return true
	}
	return false
}

type MoveSessionError struct {
	Data MoveSessionErrorData `json:"data,required"`
	Name MoveSessionErrorName `json:"name,required"`
	JSON moveSessionErrorJSON `json:"-"`
}

type moveSessionErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MoveSessionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r moveSessionErrorJSON) RawJSON() string {
	return r.raw
}

type MoveSessionErrorData struct {
	Message string                   `json:"message,required"`
	JSON    moveSessionErrorDataJSON `json:"-"`
}

type moveSessionErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MoveSessionErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r moveSessionErrorDataJSON) RawJSON() string {
	return r.raw
}

type MoveSessionErrorName string

const (
	MoveSessionErrorNameMoveSessionError MoveSessionErrorName = "MoveSessionError"
)

func (r MoveSessionErrorName) IsKnown() bool {
	switch r {
	case MoveSessionErrorNameMoveSessionError:
		return true
	}
	return false
}

type ProjectCopyError struct {
	Data ProjectCopyErrorData `json:"data,required"`
	Name ProjectCopyErrorName `json:"name,required"`
	JSON projectCopyErrorJSON `json:"-"`
}

type projectCopyErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectCopyError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCopyErrorJSON) RawJSON() string {
	return r.raw
}

type ProjectCopyErrorData struct {
	Message       string                   `json:"message,required"`
	ForceRequired bool                     `json:"forceRequired"`
	JSON          projectCopyErrorDataJSON `json:"-"`
}

type projectCopyErrorDataJSON struct {
	Message       apijson.Field
	ForceRequired apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectCopyErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCopyErrorDataJSON) RawJSON() string {
	return r.raw
}

type ProjectCopyErrorName string

const (
	ProjectCopyErrorNameProjectCopyError ProjectCopyErrorName = "ProjectCopyError"
)

func (r ProjectCopyErrorName) IsKnown() bool {
	switch r {
	case ProjectCopyErrorNameProjectCopyError:
		return true
	}
	return false
}
