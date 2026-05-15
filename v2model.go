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

// V2ModelService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2ModelService] method instead.
type V2ModelService struct {
	Options []option.RequestOption
}

// NewV2ModelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2ModelService(opts ...option.RequestOption) (r *V2ModelService) {
	r = &V2ModelService{}
	r.Options = opts
	return
}

// List v2 models
func (r *V2ModelService) List(ctx context.Context, query V2ModelListParams, opts ...option.RequestOption) (res *[]V2ModelInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V2ModelInfo struct {
	ID           string                     `json:"id,required"`
	ApiID        string                     `json:"apiID,required"`
	ProviderID   string                     `json:"providerID,required"`
	Family       string                     `json:"family"`
	Name         string                     `json:"name,required"`
	Endpoint     V2ModelInfoEndpointUnion   `json:"endpoint,required"`
	Capabilities V2ModelInfoCapabilities    `json:"capabilities,required"`
	Options      V2ModelInfoOptions         `json:"options,required"`
	Variants     []V2ModelInfoVariant       `json:"variants,required"`
	Time         V2ModelInfoTime            `json:"time,required"`
	Cost         []V2ModelInfoCostItem      `json:"cost,required"`
	Status       V2ModelInfoStatus          `json:"status,required"`
	Enabled      bool                       `json:"enabled,required"`
	Limit        V2ModelInfoLimit           `json:"limit,required"`
	JSON         v2ModelInfoJSON            `json:"-"`
}

// v2ModelInfoJSON contains the JSON metadata for the struct [V2ModelInfo]
type v2ModelInfoJSON struct {
	ID           apijson.Field
	ApiID        apijson.Field
	ProviderID   apijson.Field
	Family       apijson.Field
	Name         apijson.Field
	Endpoint     apijson.Field
	Capabilities apijson.Field
	Options      apijson.Field
	Variants     apijson.Field
	Time         apijson.Field
	Cost         apijson.Field
	Status       apijson.Field
	Enabled      apijson.Field
	Limit        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ModelInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoStatus string

const (
	V2ModelInfoStatusAlpha      V2ModelInfoStatus = "alpha"
	V2ModelInfoStatusBeta       V2ModelInfoStatus = "beta"
	V2ModelInfoStatusDeprecated V2ModelInfoStatus = "deprecated"
	V2ModelInfoStatusActive     V2ModelInfoStatus = "active"
)

func (r V2ModelInfoStatus) IsKnown() bool {
	switch r {
	case V2ModelInfoStatusAlpha, V2ModelInfoStatusBeta, V2ModelInfoStatusDeprecated, V2ModelInfoStatusActive:
		return true
	}
	return false
}

type V2ModelInfoEndpointUnion interface {
	implementsV2ModelInfoEndpointUnion()
}

type V2ModelInfoEndpointUnknown struct {
	Type V2ModelInfoEndpointType            `json:"type,required"`
	JSON v2ModelInfoEndpointUnknownJSON     `json:"-"`
}

type V2ModelInfoEndpointType string

const (
	V2ModelInfoEndpointTypeUnknown            V2ModelInfoEndpointType = "unknown"
	V2ModelInfoEndpointTypeOpenAIResponses    V2ModelInfoEndpointType = "openai/responses"
	V2ModelInfoEndpointTypeOpenAICompletions  V2ModelInfoEndpointType = "openai/completions"
	V2ModelInfoEndpointTypeAnthropicMessages  V2ModelInfoEndpointType = "anthropic/messages"
	V2ModelInfoEndpointTypeAisdk              V2ModelInfoEndpointType = "aisdk"
)

func (r V2ModelInfoEndpointType) IsKnown() bool {
	switch r {
	case V2ModelInfoEndpointTypeUnknown, V2ModelInfoEndpointTypeOpenAIResponses,
		V2ModelInfoEndpointTypeOpenAICompletions, V2ModelInfoEndpointTypeAnthropicMessages,
		V2ModelInfoEndpointTypeAisdk:
		return true
	}
	return false
}

// v2ModelInfoEndpointUnknownJSON contains the JSON metadata
type v2ModelInfoEndpointUnknownJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoEndpointUnknown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoEndpointUnknownJSON) RawJSON() string {
	return r.raw
}

func (r V2ModelInfoEndpointUnknown) implementsV2ModelInfoEndpointUnion() {}

type V2ModelInfoEndpointOpenAIResponses struct {
	Type      V2ModelInfoEndpointType               `json:"type,required"`
	URL       string                                 `json:"url,required"`
	Websocket bool                                   `json:"websocket"`
	JSON      v2ModelInfoEndpointOpenAIResponsesJSON `json:"-"`
}

type v2ModelInfoEndpointOpenAIResponsesJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	Websocket   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoEndpointOpenAIResponses) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoEndpointOpenAIResponsesJSON) RawJSON() string {
	return r.raw
}

func (r V2ModelInfoEndpointOpenAIResponses) implementsV2ModelInfoEndpointUnion() {}

type V2ModelInfoEndpointOpenAICompletions struct {
	Type      V2ModelInfoEndpointType                 `json:"type,required"`
	URL       string                                   `json:"url,required"`
	Reasoning V2ModelInfoEndpointReasoningUnion        `json:"reasoning"`
	JSON      v2ModelInfoEndpointOpenAICompletionsJSON `json:"-"`
}

type v2ModelInfoEndpointOpenAICompletionsJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	Reasoning   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoEndpointOpenAICompletions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoEndpointOpenAICompletionsJSON) RawJSON() string {
	return r.raw
}

func (r V2ModelInfoEndpointOpenAICompletions) implementsV2ModelInfoEndpointUnion() {}

type V2ModelInfoEndpointReasoningUnion interface {
	implementsV2ModelInfoEndpointReasoningUnion()
}

type V2ModelInfoEndpointReasoningContent struct {
	Type V2ModelInfoEndpointReasoningType         `json:"type,required"`
	JSON v2ModelInfoEndpointReasoningContentJSON  `json:"-"`
}

type V2ModelInfoEndpointReasoningType string

const (
	V2ModelInfoEndpointReasoningTypeContent V2ModelInfoEndpointReasoningType = "reasoning_content"
	V2ModelInfoEndpointReasoningTypeDetails V2ModelInfoEndpointReasoningType = "reasoning_details"
)

func (r V2ModelInfoEndpointReasoningType) IsKnown() bool {
	switch r {
	case V2ModelInfoEndpointReasoningTypeContent, V2ModelInfoEndpointReasoningTypeDetails:
		return true
	}
	return false
}

type v2ModelInfoEndpointReasoningContentJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoEndpointReasoningContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoEndpointReasoningContentJSON) RawJSON() string {
	return r.raw
}

func (r V2ModelInfoEndpointReasoningContent) implementsV2ModelInfoEndpointReasoningUnion() {}

type V2ModelInfoEndpointReasoningDetails struct {
	Type V2ModelInfoEndpointReasoningType          `json:"type,required"`
	JSON v2ModelInfoEndpointReasoningDetailsJSON   `json:"-"`
}

type v2ModelInfoEndpointReasoningDetailsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoEndpointReasoningDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoEndpointReasoningDetailsJSON) RawJSON() string {
	return r.raw
}

func (r V2ModelInfoEndpointReasoningDetails) implementsV2ModelInfoEndpointReasoningUnion() {}

type V2ModelInfoEndpointAnthropicMessages struct {
	Type V2ModelInfoEndpointType                  `json:"type,required"`
	URL  string                                    `json:"url,required"`
	JSON v2ModelInfoEndpointAnthropicMessagesJSON  `json:"-"`
}

type v2ModelInfoEndpointAnthropicMessagesJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoEndpointAnthropicMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoEndpointAnthropicMessagesJSON) RawJSON() string {
	return r.raw
}

func (r V2ModelInfoEndpointAnthropicMessages) implementsV2ModelInfoEndpointUnion() {}

type V2ModelInfoEndpointAisdk struct {
	Type    V2ModelInfoEndpointType        `json:"type,required"`
	Package string                          `json:"package,required"`
	URL     string                          `json:"url"`
	JSON    v2ModelInfoEndpointAisdkJSON   `json:"-"`
}

type v2ModelInfoEndpointAisdkJSON struct {
	Type        apijson.Field
	Package     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoEndpointAisdk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoEndpointAisdkJSON) RawJSON() string {
	return r.raw
}

func (r V2ModelInfoEndpointAisdk) implementsV2ModelInfoEndpointUnion() {}

type V2ModelInfoCapabilities struct {
	Tools  bool                          `json:"tools,required"`
	Input  []string                      `json:"input,required"`
	Output []string                      `json:"output,required"`
	JSON   v2ModelInfoCapabilitiesJSON   `json:"-"`
}

type v2ModelInfoCapabilitiesJSON struct {
	Tools       apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoCapabilitiesJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoOptions struct {
	Headers map[string]string           `json:"headers,required"`
	Body    map[string]interface{}      `json:"body,required"`
	Aisdk   V2ModelInfoOptionsAisdk     `json:"aisdk,required"`
	Variant string                      `json:"variant"`
	JSON    v2ModelInfoOptionsJSON      `json:"-"`
}

type v2ModelInfoOptionsJSON struct {
	Headers     apijson.Field
	Body        apijson.Field
	Aisdk       apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoOptionsJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoOptionsAisdk struct {
	Provider map[string]interface{}      `json:"provider,required"`
	Request  map[string]interface{}      `json:"request,required"`
	JSON     v2ModelInfoOptionsAisdkJSON `json:"-"`
}

type v2ModelInfoOptionsAisdkJSON struct {
	Provider    apijson.Field
	Request     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoOptionsAisdk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoOptionsAisdkJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoVariant struct {
	ID      string                 `json:"id,required"`
	Headers map[string]string      `json:"headers,required"`
	Body    map[string]interface{} `json:"body,required"`
	Aisdk   V2ModelInfoOptionsAisdk `json:"aisdk,required"`
	JSON    v2ModelInfoVariantJSON `json:"-"`
}

type v2ModelInfoVariantJSON struct {
	ID          apijson.Field
	Headers     apijson.Field
	Body        apijson.Field
	Aisdk       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoVariantJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoTime struct {
	Released float64            `json:"released,required"`
	JSON     v2ModelInfoTimeJSON `json:"-"`
}

type v2ModelInfoTimeJSON struct {
	Released    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoTimeJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoCostItem struct {
	Tier   V2ModelInfoCostTier      `json:"tier"`
	Input  float64                  `json:"input,required"`
	Output float64                  `json:"output,required"`
	Cache  V2ModelInfoCostCache     `json:"cache,required"`
	JSON   v2ModelInfoCostItemJSON  `json:"-"`
}

type v2ModelInfoCostItemJSON struct {
	Tier        apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoCostItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoCostItemJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoCostTier struct {
	Type string                 `json:"type,required"`
	Size float64                `json:"size,required"`
	JSON v2ModelInfoCostTierJSON `json:"-"`
}

type v2ModelInfoCostTierJSON struct {
	Type        apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoCostTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoCostTierJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoCostCache struct {
	Read  float64                  `json:"read,required"`
	Write float64                  `json:"write,required"`
	JSON  v2ModelInfoCostCacheJSON `json:"-"`
}

type v2ModelInfoCostCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoCostCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoCostCacheJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoLimit struct {
	JSON v2ModelInfoLimitJSON `json:"-"`
	// ExtraFields contains additional fields that may be present.
	ExtraFields map[string]interface{} `json:"-"`
}

type v2ModelInfoLimitJSON struct {
	ExtraFields map[string]apijson.Field
	raw         string
}

func (r *V2ModelInfoLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoLimitJSON) RawJSON() string {
	return r.raw
}

type V2ModelListParams struct {
	Instance param.Field[V2InstanceParam] `query:"instance"`
}

func (r V2ModelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2InstanceParam struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}
