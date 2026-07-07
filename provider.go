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
)

// ProviderService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderService] method instead.
type ProviderService struct {
	Options []option.RequestOption
}

// NewProviderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProviderService(opts ...option.RequestOption) (r *ProviderService) {
	r = &ProviderService{}
	r.Options = opts
	return
}

// List all providers
func (r *ProviderService) List(ctx context.Context, query ProviderListParams, opts ...option.RequestOption) (res *ProviderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "provider"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get provider authentication methods
func (r *ProviderService) Auth(ctx context.Context, query ProviderAuthParams, opts ...option.RequestOption) (res *map[string][]ProviderAuthMethod, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "provider/auth"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Start OAuth authorization flow
func (r *ProviderService) OauthAuthorize(ctx context.Context, providerID string, params ProviderOauthAuthorizeParams, opts ...option.RequestOption) (res *ProviderOauthAuthorizeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if providerID == "" {
		err = errors.New("missing required providerID parameter")
		return
	}
	path := fmt.Sprintf("provider/%s/oauth/authorize", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// OAuth callback
func (r *ProviderService) OauthCallback(ctx context.Context, providerID string, params ProviderOauthCallbackParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if providerID == "" {
		err = errors.New("missing required providerID parameter")
		return
	}
	path := fmt.Sprintf("provider/%s/oauth/callback", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ProviderListResponse represents the response from listing providers.
type ProviderListResponse struct {
	All       []ProviderInfo           `json:"all,required"`
	Default   map[string]string        `json:"default,required"`
	Connected []string                 `json:"connected,required"`
	JSON      providerListResponseJSON `json:"-"`
}

// providerListResponseJSON contains the JSON metadata for the struct
// [ProviderListResponse]
type providerListResponseJSON struct {
	All         apijson.Field
	Default     apijson.Field
	Connected   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListResponseJSON) RawJSON() string {
	return r.raw
}

// ProviderInfo represents information about a provider.
type ProviderInfo struct {
	ID      string                       `json:"id,required"`
	Name    string                       `json:"name,required"`
	Source  string                       `json:"source,required"`
	Env     []string                     `json:"env,required"`
	Key     string                       `json:"key"`
	Options map[string]interface{}       `json:"options,required"`
	Models  map[string]ProviderModelInfo `json:"models,required"`
	JSON    providerInfoJSON             `json:"-"`
}

// providerInfoJSON contains the JSON metadata for the struct [ProviderInfo]
type providerInfoJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Env         apijson.Field
	Key         apijson.Field
	Options     apijson.Field
	Models      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerInfoJSON) RawJSON() string {
	return r.raw
}

// ProviderModelInfo represents information about a model provided by a provider.
type ProviderModelInfo struct {
	ID           string                          `json:"id,required"`
	Name         string                          `json:"name,required"`
	Family       string                          `json:"family"`
	ReleaseDate  string                          `json:"release_date,required"`
	Attachment   bool                            `json:"attachment,required"`
	Reasoning    bool                            `json:"reasoning,required"`
	Temperature  bool                            `json:"temperature,required"`
	ToolCall     bool                            `json:"toolcall,required"`
	Cost         ProviderModelCost               `json:"cost"`
	Limit        ProviderModelLimit              `json:"limit,required"`
	Modalities   ProviderModelModalities         `json:"modalities"`
	Experimental bool                            `json:"experimental"`
	Status       string                          `json:"status"`
	Options      map[string]interface{}          `json:"options,required"`
	Headers      map[string]string               `json:"headers"`
	API          ProviderModelInfoAPI            `json:"api"`
	Provider     ProviderModelProvider           `json:"provider"`
	Variants     map[string]ProviderModelVariant `json:"variants"`
	JSON         providerModelInfoJSON           `json:"-"`
}

// providerModelInfoJSON contains the JSON metadata for the struct [ProviderModelInfo]
type providerModelInfoJSON struct {
	ID           apijson.Field
	Name         apijson.Field
	Family       apijson.Field
	ReleaseDate  apijson.Field
	Attachment   apijson.Field
	Reasoning    apijson.Field
	Temperature  apijson.Field
	ToolCall     apijson.Field
	Cost         apijson.Field
	Limit        apijson.Field
	Modalities   apijson.Field
	Experimental apijson.Field
	Status       apijson.Field
	Options      apijson.Field
	Headers      apijson.Field
	API          apijson.Field
	Provider     apijson.Field
	Variants     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProviderModelInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelInfoJSON) RawJSON() string {
	return r.raw
}

// ProviderModelInfoAPI represents API information for a model.
type ProviderModelInfoAPI struct {
	ID   string                   `json:"id"`
	URL  string                   `json:"url"`
	NPM  string                   `json:"npm"`
	JSON providerModelInfoAPIJSON `json:"-"`
}

// providerModelInfoAPIJSON contains the JSON metadata for the struct [ProviderModelInfoAPI]
type providerModelInfoAPIJSON struct {
	ID          apijson.Field
	URL         apijson.Field
	NPM         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelInfoAPI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelInfoAPIJSON) RawJSON() string {
	return r.raw
}

// ProviderModelCost represents the cost structure for a model.
type ProviderModelCost struct {
	Input           float64                          `json:"input,required"`
	Output          float64                          `json:"output,required"`
	CacheRead       float64                          `json:"cache_read"`
	CacheWrite      float64                          `json:"cache_write"`
	ContextOver200k ProviderModelCostContextOver200k `json:"experimentalOver200K"`
	JSON            providerModelCostJSON            `json:"-"`
}

// providerModelCostJSON contains the JSON metadata for the struct [ProviderModelCost]
type providerModelCostJSON struct {
	Input           apijson.Field
	Output          apijson.Field
	CacheRead       apijson.Field
	CacheWrite      apijson.Field
	ContextOver200k apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ProviderModelCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCostJSON) RawJSON() string {
	return r.raw
}

// ProviderModelCostContextOver200k represents cost structure for context over 200k tokens.
type ProviderModelCostContextOver200k struct {
	Input      float64                              `json:"input,required"`
	Output     float64                              `json:"output,required"`
	CacheRead  float64                              `json:"cache_read"`
	CacheWrite float64                              `json:"cache_write"`
	JSON       providerModelCostContextOver200kJSON `json:"-"`
}

// providerModelCostContextOver200kJSON contains the JSON metadata for the struct
// [ProviderModelCostContextOver200k]
type providerModelCostContextOver200kJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	CacheRead   apijson.Field
	CacheWrite  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelCostContextOver200k) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCostContextOver200kJSON) RawJSON() string {
	return r.raw
}

// ProviderModelLimit represents limits for a model.
type ProviderModelLimit struct {
	Context int64                  `json:"context,required"`
	Input   int64                  `json:"input"`
	Output  int64                  `json:"output,required"`
	JSON    providerModelLimitJSON `json:"-"`
}

// providerModelLimitJSON contains the JSON metadata for the struct [ProviderModelLimit]
type providerModelLimitJSON struct {
	Context     apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelLimitJSON) RawJSON() string {
	return r.raw
}

// ProviderModelModalities represents input and output modalities for a model.
type ProviderModelModalities struct {
	Input  []string                    `json:"input"`
	Output []string                    `json:"output"`
	JSON   providerModelModalitiesJSON `json:"-"`
}

// providerModelModalitiesJSON contains the JSON metadata for the struct
// [ProviderModelModalities]
type providerModelModalitiesJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelModalities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelModalitiesJSON) RawJSON() string {
	return r.raw
}

// ProviderModelProvider represents provider information for a model.
type ProviderModelProvider struct {
	NPM  string                    `json:"npm"`
	API  string                    `json:"api"`
	JSON providerModelProviderJSON `json:"-"`
}

// providerModelProviderJSON contains the JSON metadata for the struct
// [ProviderModelProvider]
type providerModelProviderJSON struct {
	NPM         apijson.Field
	API         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelProviderJSON) RawJSON() string {
	return r.raw
}

// ProviderModelVariant represents a variant of a model.
type ProviderModelVariant struct {
	Disabled bool                     `json:"disabled"`
	JSON     providerModelVariantJSON `json:"-"`
}

// providerModelVariantJSON contains the JSON metadata for the struct
// [ProviderModelVariant]
type providerModelVariantJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelVariantJSON) RawJSON() string {
	return r.raw
}

// ProviderAuthMethod represents an authentication method for a provider.
type ProviderAuthMethod struct {
	Type    string                     `json:"type,required"`
	Label   string                     `json:"label,required"`
	Prompts []ProviderAuthMethodPrompt `json:"prompts"`
	JSON    providerAuthMethodJSON     `json:"-"`
}

// providerAuthMethodJSON contains the JSON metadata for the struct [ProviderAuthMethod]
type providerAuthMethodJSON struct {
	Type        apijson.Field
	Label       apijson.Field
	Prompts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodJSON) RawJSON() string {
	return r.raw
}

// ProviderAuthMethodPrompt represents a prompt in an authentication method.
// It can be either a text prompt or a select prompt.
type ProviderAuthMethodPrompt interface {
	implementsProviderAuthMethodPrompt()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProviderAuthMethodPrompt)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			Type: reflect.TypeOf(ProviderAuthMethodPromptText{}),
		},
		apijson.UnionVariant{
			Type: reflect.TypeOf(ProviderAuthMethodPromptSelect{}),
		},
	)
}

// ProviderAuthMethodPromptText represents a text input prompt for authentication.
type ProviderAuthMethodPromptText struct {
	Type        string                           `json:"type,required"`
	Key         string                           `json:"key,required"`
	Message     string                           `json:"message,required"`
	Placeholder string                           `json:"placeholder"`
	When        ProviderAuthMethodPromptWhen     `json:"when"`
	JSON        providerAuthMethodPromptTextJSON `json:"-"`
}

// providerAuthMethodPromptTextJSON contains the JSON metadata for the struct
// [ProviderAuthMethodPromptText]
type providerAuthMethodPromptTextJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Placeholder apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptText) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptTextJSON) RawJSON() string {
	return r.raw
}

func (r ProviderAuthMethodPromptText) implementsProviderAuthMethodPrompt() {}

// ProviderAuthMethodPromptSelect represents a select prompt for authentication.
type ProviderAuthMethodPromptSelect struct {
	Type    string                                 `json:"type,required"`
	Key     string                                 `json:"key,required"`
	Message string                                 `json:"message,required"`
	Options []ProviderAuthMethodPromptSelectOption `json:"options,required"`
	When    ProviderAuthMethodPromptWhen           `json:"when"`
	JSON    providerAuthMethodPromptSelectJSON     `json:"-"`
}

// providerAuthMethodPromptSelectJSON contains the JSON metadata for the struct
// [ProviderAuthMethodPromptSelect]
type providerAuthMethodPromptSelectJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Options     apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptSelect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptSelectJSON) RawJSON() string {
	return r.raw
}

func (r ProviderAuthMethodPromptSelect) implementsProviderAuthMethodPrompt() {}

// ProviderAuthMethodPromptWhen represents condition for showing a prompt.
type ProviderAuthMethodPromptWhen struct {
	Key   string                           `json:"key,required"`
	Op    ProviderAuthMethodPromptWhenOp   `json:"op,required"`
	Value string                           `json:"value,required"`
	JSON  providerAuthMethodPromptWhenJSON `json:"-"`
}

// ProviderAuthMethodPromptWhenOp represents the comparison operator for the when condition.
type ProviderAuthMethodPromptWhenOp string

const (
	ProviderAuthMethodPromptWhenOpEq  ProviderAuthMethodPromptWhenOp = "eq"
	ProviderAuthMethodPromptWhenOpNeq ProviderAuthMethodPromptWhenOp = "neq"
)

func (r ProviderAuthMethodPromptWhenOp) IsKnown() bool {
	switch r {
	case ProviderAuthMethodPromptWhenOpEq, ProviderAuthMethodPromptWhenOpNeq:
		return true
	}
	return false
}

// providerAuthMethodPromptWhenJSON contains the JSON metadata for the struct
// [ProviderAuthMethodPromptWhen]
type providerAuthMethodPromptWhenJSON struct {
	Key         apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptWhen) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptWhenJSON) RawJSON() string {
	return r.raw
}

// ProviderAuthMethodPromptSelectOption represents an option in a select prompt.
type ProviderAuthMethodPromptSelectOption struct {
	Label string                                   `json:"label,required"`
	Value string                                   `json:"value,required"`
	Hint  string                                   `json:"hint"`
	JSON  providerAuthMethodPromptSelectOptionJSON `json:"-"`
}

// providerAuthMethodPromptSelectOptionJSON contains the JSON metadata for the struct
// [ProviderAuthMethodPromptSelectOption]
type providerAuthMethodPromptSelectOptionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	Hint        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthMethodPromptSelectOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthMethodPromptSelectOptionJSON) RawJSON() string {
	return r.raw
}

// ProviderOauthAuthorizeBody represents the body for OAuth authorization request.
type ProviderOauthAuthorizeBody struct {
	Method int64                          `json:"method,required"`
	Inputs map[string]string              `json:"inputs"`
	JSON   providerOauthAuthorizeBodyJSON `json:"-"`
}

// providerOauthAuthorizeBodyJSON contains the JSON metadata for the struct
// [ProviderOauthAuthorizeBody]
type providerOauthAuthorizeBodyJSON struct {
	Method      apijson.Field
	Inputs      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderOauthAuthorizeBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerOauthAuthorizeBodyJSON) RawJSON() string {
	return r.raw
}

// ProviderOauthAuthorizeResponse represents the response from OAuth authorization.
type ProviderOauthAuthorizeResponse struct {
	URL          string                               `json:"url,required"`
	Method       ProviderOauthAuthorizeResponseMethod `json:"method,required"`
	Instructions string                               `json:"instructions,required"`
	JSON         providerOauthAuthorizeResponseJSON   `json:"-"`
}

type ProviderOauthAuthorizeResponseMethod string

const (
	ProviderOauthAuthorizeResponseMethodAuto ProviderOauthAuthorizeResponseMethod = "auto"
	ProviderOauthAuthorizeResponseMethodCode ProviderOauthAuthorizeResponseMethod = "code"
)

func (r ProviderOauthAuthorizeResponseMethod) IsKnown() bool {
	switch r {
	case ProviderOauthAuthorizeResponseMethodAuto, ProviderOauthAuthorizeResponseMethodCode:
		return true
	}
	return false
}

// providerOauthAuthorizeResponseJSON contains the JSON metadata for the struct
// [ProviderOauthAuthorizeResponse]
type providerOauthAuthorizeResponseJSON struct {
	URL          apijson.Field
	Method       apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProviderOauthAuthorizeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerOauthAuthorizeResponseJSON) RawJSON() string {
	return r.raw
}

// ProviderOauthCallbackBody represents the body for OAuth callback.
type ProviderOauthCallbackBody struct {
	Method int64                         `json:"method,required"`
	Code   string                        `json:"code"`
	JSON   providerOauthCallbackBodyJSON `json:"-"`
}

// providerOauthCallbackBodyJSON contains the JSON metadata for the struct
// [ProviderOauthCallbackBody]
type providerOauthCallbackBodyJSON struct {
	Method      apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderOauthCallbackBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerOauthCallbackBodyJSON) RawJSON() string {
	return r.raw
}

// ProviderListParams contains the query parameters for listing providers.
type ProviderListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ProviderListParams]'s query parameters as `url.Values`.
func (r ProviderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ProviderAuthParams contains the query parameters for getting provider auth methods.
type ProviderAuthParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ProviderAuthParams]'s query parameters as `url.Values`.
func (r ProviderAuthParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ProviderOauthAuthorizeParams contains the parameters for OAuth authorization.
type ProviderOauthAuthorizeParams struct {
	Directory param.Field[string]            `query:"directory"`
	Workspace param.Field[string]            `query:"workspace"`
	Method    param.Field[int64]             `json:"method,required"`
	Inputs    param.Field[map[string]string] `json:"inputs"`
}

// MarshalJSON serializes [ProviderOauthAuthorizeParams] omitting query parameters.
func (r ProviderOauthAuthorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ProviderOauthAuthorizeParams]'s query parameters as `url.Values`.
func (r ProviderOauthAuthorizeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ProviderOauthCallbackParams contains the parameters for OAuth callback.
type ProviderOauthCallbackParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Method    param.Field[int64]  `json:"method,required"`
	Code      param.Field[string] `json:"code"`
}

// MarshalJSON serializes [ProviderOauthCallbackParams] omitting query parameters.
func (r ProviderOauthCallbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ProviderOauthCallbackParams]'s query parameters as `url.Values`.
func (r ProviderOauthCallbackParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
