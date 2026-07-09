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
	ID      string                   `json:"id,required"`
	Name    string                   `json:"name,required"`
	Source  ProviderInfoSource       `json:"source,required"`
	Env     []string                 `json:"env,required"`
	Key     string                   `json:"key"`
	Options map[string]interface{}   `json:"options,required"`
	Models  map[string]ProviderModel `json:"models,required"`
	JSON    providerInfoJSON         `json:"-"`
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

// ProviderInfoSource represents the source of a provider configuration.
type ProviderInfoSource string

const (
	ProviderInfoSourceEnv    ProviderInfoSource = "env"
	ProviderInfoSourceConfig ProviderInfoSource = "config"
	ProviderInfoSourceCustom ProviderInfoSource = "custom"
	ProviderInfoSourceApi    ProviderInfoSource = "api"
)

func (r ProviderInfoSource) IsKnown() bool {
	switch r {
	case ProviderInfoSourceEnv, ProviderInfoSourceConfig, ProviderInfoSourceCustom, ProviderInfoSourceApi:
		return true
	}
	return false
}

// ProviderModel represents information about a model provided by a provider.
type ProviderModel struct {
	ID           string                    `json:"id,required"`
	ProviderID   string                    `json:"providerID,required"`
	API          ProviderModelAPI          `json:"api,required"`
	Name         string                    `json:"name,required"`
	Family       string                    `json:"family"`
	Capabilities ProviderModelCapabilities `json:"capabilities,required"`
	Cost         ProviderModelCost         `json:"cost,required"`
	Limit        ProviderModelLimit        `json:"limit,required"`
	Status       ProviderModelStatus       `json:"status,required"`
	Options      map[string]interface{}    `json:"options,required"`
	Headers      map[string]string         `json:"headers,required"`
	ReleaseDate  string                    `json:"release_date,required"`
	Variants     map[string]interface{}    `json:"variants"`
	JSON         providerModelJSON         `json:"-"`
}

// providerModelJSON contains the JSON metadata for the struct [ProviderModel]
type providerModelJSON struct {
	ID           apijson.Field
	ProviderID   apijson.Field
	API          apijson.Field
	Name         apijson.Field
	Family       apijson.Field
	Capabilities apijson.Field
	Cost         apijson.Field
	Limit        apijson.Field
	Status       apijson.Field
	Options      apijson.Field
	Headers      apijson.Field
	ReleaseDate  apijson.Field
	Variants     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProviderModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelJSON) RawJSON() string {
	return r.raw
}

// ProviderModelAPI represents API information for a model.
type ProviderModelAPI struct {
	ID   string               `json:"id,required"`
	URL  string               `json:"url,required"`
	NPM  string               `json:"npm,required"`
	JSON providerModelAPIJSON `json:"-"`
}

// providerModelAPIJSON contains the JSON metadata for the struct [ProviderModelAPI]
type providerModelAPIJSON struct {
	ID          apijson.Field
	URL         apijson.Field
	NPM         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelAPI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelAPIJSON) RawJSON() string {
	return r.raw
}

// ProviderModelStatus represents the status of a model.
type ProviderModelStatus string

const (
	ProviderModelStatusAlpha      ProviderModelStatus = "alpha"
	ProviderModelStatusBeta       ProviderModelStatus = "beta"
	ProviderModelStatusDeprecated ProviderModelStatus = "deprecated"
	ProviderModelStatusActive     ProviderModelStatus = "active"
)

func (r ProviderModelStatus) IsKnown() bool {
	switch r {
	case ProviderModelStatusAlpha, ProviderModelStatusBeta, ProviderModelStatusDeprecated, ProviderModelStatusActive:
		return true
	}
	return false
}

// ProviderModelCapabilities represents the capabilities of a model.
type ProviderModelCapabilities struct {
	Temperature bool                                `json:"temperature,required"`
	Reasoning   bool                                `json:"reasoning,required"`
	Attachment  bool                                `json:"attachment,required"`
	ToolCall    bool                                `json:"toolcall,required"`
	Input       ProviderModelCapabilitiesModalities `json:"input,required"`
	Output      ProviderModelCapabilitiesModalities `json:"output,required"`
	// This field can have the runtime type of [bool], [ProviderModelCapabilitiesInterleavedField].
	Interleaved interface{}                   `json:"interleaved,required"`
	JSON        providerModelCapabilitiesJSON `json:"-"`
}

// providerModelCapabilitiesJSON contains the JSON metadata for the struct
// [ProviderModelCapabilities]
type providerModelCapabilitiesJSON struct {
	Temperature apijson.Field
	Reasoning   apijson.Field
	Attachment  apijson.Field
	ToolCall    apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	Interleaved apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCapabilitiesJSON) RawJSON() string {
	return r.raw
}

// ProviderModelCapabilitiesModalities represents input/output modality support for a model.
type ProviderModelCapabilitiesModalities struct {
	Text  bool                                    `json:"text,required"`
	Audio bool                                    `json:"audio,required"`
	Image bool                                    `json:"image,required"`
	Video bool                                    `json:"video,required"`
	PDF   bool                                    `json:"pdf,required"`
	JSON  providerModelCapabilitiesModalitiesJSON `json:"-"`
}

// providerModelCapabilitiesModalitiesJSON contains the JSON metadata for the struct
// [ProviderModelCapabilitiesModalities]
type providerModelCapabilitiesModalitiesJSON struct {
	Text        apijson.Field
	Audio       apijson.Field
	Image       apijson.Field
	Video       apijson.Field
	PDF         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelCapabilitiesModalities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCapabilitiesModalitiesJSON) RawJSON() string {
	return r.raw
}

// ProviderModelCapabilitiesInterleavedField represents the field name used for interleaved
// reasoning content.
type ProviderModelCapabilitiesInterleavedField struct {
	Field ProviderModelCapabilitiesInterleavedFieldField `json:"field,required"`
	JSON  providerModelCapabilitiesInterleavedFieldJSON  `json:"-"`
}

// providerModelCapabilitiesInterleavedFieldJSON contains the JSON metadata for the struct
// [ProviderModelCapabilitiesInterleavedField]
type providerModelCapabilitiesInterleavedFieldJSON struct {
	Field       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelCapabilitiesInterleavedField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCapabilitiesInterleavedFieldJSON) RawJSON() string {
	return r.raw
}

// ProviderModelCapabilitiesInterleavedFieldField represents the field name where interleaved
// reasoning content is located.
type ProviderModelCapabilitiesInterleavedFieldField string

const (
	ProviderModelCapabilitiesInterleavedFieldFieldReasoning        ProviderModelCapabilitiesInterleavedFieldField = "reasoning"
	ProviderModelCapabilitiesInterleavedFieldFieldReasoningContent ProviderModelCapabilitiesInterleavedFieldField = "reasoning_content"
	ProviderModelCapabilitiesInterleavedFieldFieldReasoningDetails ProviderModelCapabilitiesInterleavedFieldField = "reasoning_details"
)

func (r ProviderModelCapabilitiesInterleavedFieldField) IsKnown() bool {
	switch r {
	case ProviderModelCapabilitiesInterleavedFieldFieldReasoning,
		ProviderModelCapabilitiesInterleavedFieldFieldReasoningContent,
		ProviderModelCapabilitiesInterleavedFieldFieldReasoningDetails:
		return true
	}
	return false
}

// ProviderModelCost represents the cost structure for a model.
type ProviderModelCost struct {
	Input                float64                               `json:"input,required"`
	Output               float64                               `json:"output,required"`
	Cache                ProviderModelCostCache                `json:"cache,required"`
	Tiers                []ProviderModelCostTier               `json:"tiers"`
	ExperimentalOver200K ProviderModelCostExperimentalOver200K `json:"experimentalOver200K"`
	JSON                 providerModelCostJSON                 `json:"-"`
}

// providerModelCostJSON contains the JSON metadata for the struct [ProviderModelCost]
type providerModelCostJSON struct {
	Input                apijson.Field
	Output               apijson.Field
	Cache                apijson.Field
	Tiers                apijson.Field
	ExperimentalOver200K apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProviderModelCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCostJSON) RawJSON() string {
	return r.raw
}

// ProviderModelCostCache represents the cache cost structure for a model.
type ProviderModelCostCache struct {
	Read  float64                    `json:"read,required"`
	Write float64                    `json:"write,required"`
	JSON  providerModelCostCacheJSON `json:"-"`
}

// providerModelCostCacheJSON contains the JSON metadata for the struct
// [ProviderModelCostCache]
type providerModelCostCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelCostCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCostCacheJSON) RawJSON() string {
	return r.raw
}

// ProviderModelCostTier represents a single cost tier for a model.
type ProviderModelCostTier struct {
	Input  float64                      `json:"input,required"`
	Output float64                      `json:"output,required"`
	Cache  ProviderModelCostCache       `json:"cache,required"`
	Tier   ProviderModelCostTierContext `json:"tier,required"`
	JSON   providerModelCostTierJSON    `json:"-"`
}

// providerModelCostTierJSON contains the JSON metadata for the struct
// [ProviderModelCostTier]
type providerModelCostTierJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Cache       apijson.Field
	Tier        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelCostTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCostTierJSON) RawJSON() string {
	return r.raw
}

// ProviderModelCostTierContext represents the tier context threshold for a cost tier.
type ProviderModelCostTierContext struct {
	Type ProviderModelCostTierContextType `json:"type,required"`
	Size float64                          `json:"size,required"`
	JSON providerModelCostTierContextJSON `json:"-"`
}

// providerModelCostTierContextJSON contains the JSON metadata for the struct
// [ProviderModelCostTierContext]
type providerModelCostTierContextJSON struct {
	Type        apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelCostTierContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCostTierContextJSON) RawJSON() string {
	return r.raw
}

// ProviderModelCostTierContextType represents the type of cost tier.
type ProviderModelCostTierContextType string

const (
	ProviderModelCostTierContextTypeContext ProviderModelCostTierContextType = "context"
)

func (r ProviderModelCostTierContextType) IsKnown() bool {
	switch r {
	case ProviderModelCostTierContextTypeContext:
		return true
	}
	return false
}

// ProviderModelCostExperimentalOver200K represents cost structure for context over 200k tokens.
type ProviderModelCostExperimentalOver200K struct {
	Input  float64                                   `json:"input,required"`
	Output float64                                   `json:"output,required"`
	Cache  ProviderModelCostCache                    `json:"cache,required"`
	JSON   providerModelCostExperimentalOver200KJSON `json:"-"`
}

// providerModelCostExperimentalOver200KJSON contains the JSON metadata for the struct
// [ProviderModelCostExperimentalOver200K]
type providerModelCostExperimentalOver200KJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderModelCostExperimentalOver200K) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerModelCostExperimentalOver200KJSON) RawJSON() string {
	return r.raw
}

// ProviderModelLimit represents limits for a model.
type ProviderModelLimit struct {
	Context float64                `json:"context,required"`
	Input   float64                `json:"input"`
	Output  float64                `json:"output,required"`
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

// ProviderAuthMethod represents an authentication method for a provider.
type ProviderAuthMethod struct {
	Type    ProviderAuthMethodType     `json:"type,required"`
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

// ProviderAuthMethodType represents the type of an authentication method.
type ProviderAuthMethodType string

const (
	ProviderAuthMethodTypeOauth ProviderAuthMethodType = "oauth"
	ProviderAuthMethodTypeApi   ProviderAuthMethodType = "api"
)

func (r ProviderAuthMethodType) IsKnown() bool {
	switch r {
	case ProviderAuthMethodTypeOauth, ProviderAuthMethodTypeApi:
		return true
	}
	return false
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
			DiscriminatorValue: "text",
			Type:               reflect.TypeOf(ProviderAuthMethodPromptText{}),
		},
		apijson.UnionVariant{
			DiscriminatorValue: "select",
			Type:               reflect.TypeOf(ProviderAuthMethodPromptSelect{}),
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
