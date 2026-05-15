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
