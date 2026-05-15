// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// QuestionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuestionService] method instead.
type QuestionService struct {
	Options []option.RequestOption
}

// NewQuestionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuestionService(opts ...option.RequestOption) (r *QuestionService) {
	r = &QuestionService{}
	r.Options = opts
	return
}

// List pending question requests
func (r *QuestionService) List(ctx context.Context, query QuestionListParams, opts ...option.RequestOption) (res *[]QuestionRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "question"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Reply to a question request
func (r *QuestionService) Reply(ctx context.Context, requestID string, params QuestionReplyParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "question/" + requestID + "/reply"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Reject a question request
func (r *QuestionService) Reject(ctx context.Context, requestID string, query QuestionRejectParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "question/" + requestID + "/reject"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

type QuestionRequest struct {
	ID        string              `json:"id,required"`
	SessionID string              `json:"sessionID,required"`
	Questions []QuestionInfo      `json:"questions,required"`
	Tool      QuestionRequestTool `json:"tool"`
	JSON      questionRequestJSON `json:"-"`
}

// questionRequestJSON contains the JSON metadata for the struct
// [QuestionRequest]
type questionRequestJSON struct {
	ID          apijson.Field
	SessionID   apijson.Field
	Questions   apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionRequestJSON) RawJSON() string {
	return r.raw
}

type QuestionInfo struct {
	Question string           `json:"question,required"`
	Header   string           `json:"header,required"`
	Options  []QuestionOption `json:"options,required"`
	Multiple bool             `json:"multiple"`
	Custom   bool             `json:"custom"`
	JSON     questionInfoJSON `json:"-"`
}

// questionInfoJSON contains the JSON metadata for the struct
// [QuestionInfo]
type questionInfoJSON struct {
	Question    apijson.Field
	Header      apijson.Field
	Options     apijson.Field
	Multiple    apijson.Field
	Custom      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionInfoJSON) RawJSON() string {
	return r.raw
}

type QuestionOption struct {
	Label       string             `json:"label,required"`
	Description string             `json:"description,required"`
	JSON        questionOptionJSON `json:"-"`
}

// questionOptionJSON contains the JSON metadata for the struct
// [QuestionOption]
type questionOptionJSON struct {
	Label       apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionOptionJSON) RawJSON() string {
	return r.raw
}

type QuestionRequestTool struct {
	MessageID string                  `json:"messageID,required"`
	CallID    string                  `json:"callID,required"`
	JSON      questionRequestToolJSON `json:"-"`
}

// questionRequestToolJSON contains the JSON metadata for the struct
// [QuestionRequestTool]
type questionRequestToolJSON struct {
	MessageID   apijson.Field
	CallID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionRequestTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionRequestToolJSON) RawJSON() string {
	return r.raw
}

type QuestionAnswer = []string

type QuestionListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [QuestionListParams]'s query parameters as `url.Values`.
func (r QuestionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QuestionReplyParams struct {
	// Answers to the questions (array of string arrays, one for each question)
	Answers   param.Field[QuestionAnswer] `json:"answers,required"`
	Directory param.Field[string]          `query:"directory"`
	Workspace param.Field[string]          `query:"workspace"`
}

func (r QuestionReplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [QuestionReplyParams]'s query parameters as `url.Values`.
func (r QuestionReplyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QuestionRejectParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [QuestionRejectParams]'s query parameters as `url.Values`.
func (r QuestionRejectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
