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

// V2QuestionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2QuestionService] method instead.
type V2QuestionService struct {
	Options []option.RequestOption
	Request *V2QuestionRequestService
}

// NewV2QuestionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2QuestionService(opts ...option.RequestOption) (r *V2QuestionService) {
	r = &V2QuestionService{}
	r.Options = opts
	r.Request = NewV2QuestionRequestService(opts...)
	return
}

// V2QuestionRequestService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2QuestionRequestService] method instead.
type V2QuestionRequestService struct {
	Options []option.RequestOption
}

// NewV2QuestionRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2QuestionRequestService(opts ...option.RequestOption) (r *V2QuestionRequestService) {
	r = &V2QuestionRequestService{}
	r.Options = opts
	return
}

// List pending question requests for a location.
func (r *V2QuestionRequestService) List(ctx context.Context, query V2QuestionRequestListParams, opts ...option.RequestOption) (res *V2QuestionRequestListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/question/request"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// V2QuestionRequestListResponse contains the response from the question request list endpoint.
type V2QuestionRequestListResponse struct {
	Location LocationInfo                     `json:"location,required"`
	Data     []QuestionV2Request              `json:"data,required"`
	JSON     v2QuestionRequestListResponseJSON `json:"-"`
}

// v2QuestionRequestListResponseJSON contains the JSON metadata for the struct [V2QuestionRequestListResponse]
type v2QuestionRequestListResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2QuestionRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2QuestionRequestListResponseJSON) RawJSON() string {
	return r.raw
}

type QuestionV2Request struct {
	ID        string                 `json:"id,required"`
	SessionID string                 `json:"sessionID,required"`
	// Questions to ask
	Questions []QuestionV2Info       `json:"questions,required"`
	Tool      QuestionV2Tool         `json:"tool"`
	JSON      questionV2RequestJSON  `json:"-"`
}

// questionV2RequestJSON contains the JSON metadata for the struct [QuestionV2Request]
type questionV2RequestJSON struct {
	ID          apijson.Field
	SessionID   apijson.Field
	Questions   apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionV2Request) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionV2RequestJSON) RawJSON() string {
	return r.raw
}

type QuestionV2Info struct {
	// Complete question
	Question string              `json:"question,required"`
	// Very short label (max 30 chars)
	Header  string              `json:"header,required"`
	// Available choices
	Options []QuestionV2Option  `json:"options,required"`
	Multiple bool               `json:"multiple"`
	Custom   bool               `json:"custom"`
	JSON     questionV2InfoJSON `json:"-"`
}

// questionV2InfoJSON contains the JSON metadata for the struct [QuestionV2Info]
type questionV2InfoJSON struct {
	Question    apijson.Field
	Header      apijson.Field
	Options     apijson.Field
	Multiple    apijson.Field
	Custom      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionV2Info) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionV2InfoJSON) RawJSON() string {
	return r.raw
}

type QuestionV2Option struct {
	// Display text (1-5 words, concise)
	Label string `json:"label,required"`
	// Explanation of choice
	Description string               `json:"description,required"`
	JSON        questionV2OptionJSON `json:"-"`
}

// questionV2OptionJSON contains the JSON metadata for the struct [QuestionV2Option]
type questionV2OptionJSON struct {
	Label       apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionV2Option) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionV2OptionJSON) RawJSON() string {
	return r.raw
}

type QuestionV2Tool struct {
	MessageID string             `json:"messageID,required"`
	CallID    string             `json:"callID,required"`
	JSON      questionV2ToolJSON `json:"-"`
}

// questionV2ToolJSON contains the JSON metadata for the struct [QuestionV2Tool]
type questionV2ToolJSON struct {
	MessageID   apijson.Field
	CallID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuestionV2Tool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r questionV2ToolJSON) RawJSON() string {
	return r.raw
}

type V2QuestionRequestListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2QuestionRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
