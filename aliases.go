// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"github.com/sst/opencode-sdk-go/internal/apierror"
	"github.com/sst/opencode-sdk-go/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type MessageAbortedError = shared.MessageAbortedError

// This is an alias to an internal type.
type MessageAbortedErrorData = shared.MessageAbortedErrorData

// This is an alias to an internal type.
type MessageAbortedErrorName = shared.MessageAbortedErrorName

// This is an alias to an internal value.
const MessageAbortedErrorNameMessageAbortedError = shared.MessageAbortedErrorNameMessageAbortedError

// This is an alias to an internal type.
type ProviderAuthError = shared.ProviderAuthError

// This is an alias to an internal type.
type ProviderAuthErrorData = shared.ProviderAuthErrorData

// This is an alias to an internal type.
type ProviderAuthErrorName = shared.ProviderAuthErrorName

// This is an alias to an internal value.
const ProviderAuthErrorNameProviderAuthError = shared.ProviderAuthErrorNameProviderAuthError

// This is an alias to an internal type.
type ProviderAuthError1 = shared.ProviderAuthError1

// This is an alias to an internal type.
type ProviderAuthError1Data = shared.ProviderAuthError1Data

// This is an alias to an internal type.
type ProviderAuthError1Name = shared.ProviderAuthError1Name

// This is an alias to an internal value.
const ProviderAuthError1NameBadRequest = shared.ProviderAuthError1NameBadRequest

// This is an alias to an internal value.
const ProviderAuthError1NameProviderAuthOauthMissing = shared.ProviderAuthError1NameProviderAuthOauthMissing

// This is an alias to an internal value.
const ProviderAuthError1NameProviderAuthOauthCodeMissing = shared.ProviderAuthError1NameProviderAuthOauthCodeMissing

// This is an alias to an internal value.
const ProviderAuthError1NameProviderAuthOauthCallbackFailed = shared.ProviderAuthError1NameProviderAuthOauthCallbackFailed

// This is an alias to an internal value.
const ProviderAuthError1NameProviderAuthValidationFailed = shared.ProviderAuthError1NameProviderAuthValidationFailed

// This is an alias to an internal type.
type UnknownError = shared.UnknownError

// This is an alias to an internal type.
type UnknownErrorData = shared.UnknownErrorData

// This is an alias to an internal type.
type UnknownErrorName = shared.UnknownErrorName

// This is an alias to an internal value.
const UnknownErrorNameUnknownError = shared.UnknownErrorNameUnknownError

// This is an alias to an internal type.
type StructuredOutputError = shared.StructuredOutputError

// This is an alias to an internal type.
type StructuredOutputErrorData = shared.StructuredOutputErrorData

// This is an alias to an internal type.
type StructuredOutputErrorName = shared.StructuredOutputErrorName

// This is an alias to an internal value.
const StructuredOutputErrorNameStructuredOutputError = shared.StructuredOutputErrorNameStructuredOutputError

// This is an alias to an internal type.
type ContextOverflowError = shared.ContextOverflowError

// This is an alias to an internal type.
type ContextOverflowErrorData = shared.ContextOverflowErrorData

// This is an alias to an internal type.
type ContextOverflowErrorName = shared.ContextOverflowErrorName

// This is an alias to an internal value.
const ContextOverflowErrorNameContextOverflowError = shared.ContextOverflowErrorNameContextOverflowError

// This is an alias to an internal type.
type MessageOutputLengthError = shared.MessageOutputLengthError

// This is an alias to an internal type.
type MessageOutputLengthErrorName = shared.MessageOutputLengthErrorName

// This is an alias to an internal value.
const MessageOutputLengthErrorNameMessageOutputLengthError = shared.MessageOutputLengthErrorNameMessageOutputLengthError

// This is an alias to an internal type.
type APIError = shared.APIError

// This is an alias to an internal type.
type APIErrorData = shared.APIErrorData

// This is an alias to an internal type.
type APIErrorName = shared.APIErrorName

// This is an alias to an internal value.
const APIErrorNameAPIError = shared.APIErrorNameAPIError

// This is an alias to an internal type.
type VcsApplyError = shared.VcsApplyError

// This is an alias to an internal type.
type VcsApplyErrorData = shared.VcsApplyErrorData

// This is an alias to an internal type.
type VcsApplyErrorName = shared.VcsApplyErrorName

// This is an alias to an internal value.
const VcsApplyErrorNameVcsApplyError = shared.VcsApplyErrorNameVcsApplyError

// This is an alias to an internal type.
type VcsApplyErrorReason = shared.VcsApplyErrorReason

// This is an alias to an internal value.
const VcsApplyErrorReasonNonGit = shared.VcsApplyErrorReasonNonGit

// This is an alias to an internal value.
const VcsApplyErrorReasonNotClean = shared.VcsApplyErrorReasonNotClean

// This is an alias to an internal type.
type WorkspaceWarpError = shared.WorkspaceWarpError

// This is an alias to an internal type.
type WorkspaceWarpErrorData = shared.WorkspaceWarpErrorData

// This is an alias to an internal type.
type WorkspaceWarpErrorName = shared.WorkspaceWarpErrorName

// This is an alias to an internal value.
const WorkspaceWarpErrorNameWorkspaceWarpError = shared.WorkspaceWarpErrorNameWorkspaceWarpError

// This is an alias to an internal type.
type NotFoundError = shared.NotFoundError

// This is an alias to an internal type.
type NotFoundErrorData = shared.NotFoundErrorData

// This is an alias to an internal type.
type NotFoundErrorName = shared.NotFoundErrorName

// This is an alias to an internal value.
const NotFoundErrorNameNotFoundError = shared.NotFoundErrorNameNotFoundError

// This is an alias to an internal type.
type BadRequestError = shared.BadRequestError

// This is an alias to an internal type.
type BadRequestErrorData = shared.BadRequestErrorData

// This is an alias to an internal type.
type BadRequestErrorName = shared.BadRequestErrorName

// This is an alias to an internal value.
const BadRequestErrorNameBadRequest = shared.BadRequestErrorNameBadRequest

// This is an alias to an internal type.
type BadRequestErrorKind = shared.BadRequestErrorKind

// This is an alias to an internal value.
const BadRequestErrorKindParams = shared.BadRequestErrorKindParams

// This is an alias to an internal value.
const BadRequestErrorKindHeaders = shared.BadRequestErrorKindHeaders

// This is an alias to an internal value.
const BadRequestErrorKindQuery = shared.BadRequestErrorKindQuery

// This is an alias to an internal value.
const BadRequestErrorKindBody = shared.BadRequestErrorKindBody

// This is an alias to an internal value.
const BadRequestErrorKindPayload = shared.BadRequestErrorKindPayload

// This is an alias to an internal type.
type WorktreeError = shared.WorktreeError

// This is an alias to an internal type.
type WorktreeErrorData = shared.WorktreeErrorData

// This is an alias to an internal type.
type WorktreeErrorName = shared.WorktreeErrorName

// This is an alias to an internal value.
const WorktreeErrorNameWorktreeNotGitError = shared.WorktreeErrorNameWorktreeNotGitError

// This is an alias to an internal value.
const WorktreeErrorNameWorktreeNameGenerationFailedError = shared.WorktreeErrorNameWorktreeNameGenerationFailedError

// This is an alias to an internal value.
const WorktreeErrorNameWorktreeCreateFailedError = shared.WorktreeErrorNameWorktreeCreateFailedError

// This is an alias to an internal value.
const WorktreeErrorNameWorktreeStartCommandFailedError = shared.WorktreeErrorNameWorktreeStartCommandFailedError

// This is an alias to an internal value.
const WorktreeErrorNameWorktreeRemoveFailedError = shared.WorktreeErrorNameWorktreeRemoveFailedError

// This is an alias to an internal value.
const WorktreeErrorNameWorktreeResetFailedError = shared.WorktreeErrorNameWorktreeResetFailedError

// This is an alias to an internal value.
const WorktreeErrorNameWorktreeListFailedError = shared.WorktreeErrorNameWorktreeListFailedError

// === _tag-based error type aliases ===

// This is an alias to an internal type.
type InvalidCursorError = shared.InvalidCursorError

// This is an alias to an internal type.
type InvalidCursorErrorTag = shared.InvalidCursorErrorTag

// This is an alias to an internal value.
const InvalidCursorErrorTagInvalidCursorError = shared.InvalidCursorErrorTagInvalidCursorError

// This is an alias to an internal type.
type InvalidRequestError = shared.InvalidRequestError

// This is an alias to an internal type.
type InvalidRequestErrorTag = shared.InvalidRequestErrorTag

// This is an alias to an internal value.
const InvalidRequestErrorTagInvalidRequestError = shared.InvalidRequestErrorTagInvalidRequestError

// This is an alias to an internal type.
type McpUnsupportedOAuthError = shared.McpUnsupportedOAuthError

// This is an alias to an internal type.
type ProviderNotFoundError = shared.ProviderNotFoundError

// This is an alias to an internal type.
type ProviderNotFoundErrorTag = shared.ProviderNotFoundErrorTag

// This is an alias to an internal value.
const ProviderNotFoundErrorTagProviderNotFoundError = shared.ProviderNotFoundErrorTagProviderNotFoundError

// This is an alias to an internal type.
type ServiceUnavailableError = shared.ServiceUnavailableError

// This is an alias to an internal type.
type ServiceUnavailableErrorTag = shared.ServiceUnavailableErrorTag

// This is an alias to an internal value.
const ServiceUnavailableErrorTagServiceUnavailableError = shared.ServiceUnavailableErrorTagServiceUnavailableError

// This is an alias to an internal type.
type SessionNotFoundError = shared.SessionNotFoundError

// This is an alias to an internal type.
type SessionNotFoundErrorTag = shared.SessionNotFoundErrorTag

// This is an alias to an internal value.
const SessionNotFoundErrorTagSessionNotFoundError = shared.SessionNotFoundErrorTagSessionNotFoundError

// This is an alias to an internal type.
type UnknownError1 = shared.UnknownError1

// This is an alias to an internal type.
type UnknownError1Tag = shared.UnknownError1Tag

// This is an alias to an internal value.
const UnknownError1TagUnknownError = shared.UnknownError1TagUnknownError

// This is an alias to an internal type.
type AccountV2OAuthCredential = shared.AccountV2OAuthCredential

// This is an alias to an internal type.
type AccountV2ApiKeyCredential = shared.AccountV2ApiKeyCredential

// This is an alias to an internal type.
type AccountV2CredentialUnion = shared.AccountV2CredentialUnion

// This is an alias to an internal type.
type AccountV2Info = shared.AccountV2Info

// This is an alias to an internal type.
type UnauthorizedError = shared.UnauthorizedError

// This is an alias to an internal type.
type UnauthorizedErrorTag = shared.UnauthorizedErrorTag

// This is an alias to an internal value.
const UnauthorizedErrorTagUnauthorizedError = shared.UnauthorizedErrorTagUnauthorizedError

// This is an alias to an internal type.
type EffectHttpApiErrorBadRequest = shared.EffectHttpApiErrorBadRequest

// This is an alias to an internal type.
type EffectHttpApiErrorBadRequestTag = shared.EffectHttpApiErrorBadRequestTag

// This is an alias to an internal value.
const EffectHttpApiErrorBadRequestTagBadRequest = shared.EffectHttpApiErrorBadRequestTagBadRequest

// This is an alias to an internal type.
type EffectHttpApiErrorForbidden = shared.EffectHttpApiErrorForbidden

// This is an alias to an internal type.
type EffectHttpApiErrorForbiddenTag = shared.EffectHttpApiErrorForbiddenTag

// This is an alias to an internal value.
const EffectHttpApiErrorForbiddenTagForbidden = shared.EffectHttpApiErrorForbiddenTagForbidden

// This is an alias to an internal type.
type EffectHttpApiErrorInternalServerError = shared.EffectHttpApiErrorInternalServerError

// This is an alias to an internal type.
type EffectHttpApiErrorInternalServerErrorTag = shared.EffectHttpApiErrorInternalServerErrorTag

// This is an alias to an internal value.
const EffectHttpApiErrorInternalServerErrorTagInternalServerError = shared.EffectHttpApiErrorInternalServerErrorTagInternalServerError
