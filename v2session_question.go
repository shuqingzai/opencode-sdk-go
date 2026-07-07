// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// V2SessionQuestionService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2SessionQuestionService] method instead.
type V2SessionQuestionService struct {
	Options []option.RequestOption
}

// NewV2SessionQuestionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2SessionQuestionService(opts ...option.RequestOption) (r *V2SessionQuestionService) {
	r = &V2SessionQuestionService{}
	r.Options = opts
	return
}

// List session question requests
//
// Retrieve pending question requests owned by a session.
func (r *V2SessionQuestionService) List(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *[]QuestionV2Request, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/question", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Reply to pending question request
//
// Answer a pending question request owned by a session.
func (r *V2SessionQuestionService) Reply(ctx context.Context, sessionID string, requestID string, body V2SessionQuestionReplyParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/question/%s/reply", sessionID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reject pending question request
//
// Reject a pending question request owned by a session.
func (r *V2SessionQuestionService) Reject(ctx context.Context, sessionID string, requestID string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/question/%s/reject", sessionID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// ===== Param Types =====

type V2SessionQuestionReplyParams struct {
	Answers param.Field[[]QuestionV2Answer] `json:"answers,required"`
}

func (r V2SessionQuestionReplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// QuestionV2Answer represents the answer type for v2 question replies.
// Each answer is an array of selected label strings.
type QuestionV2Answer []string

func (r QuestionV2Answer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
