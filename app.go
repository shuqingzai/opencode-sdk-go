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

// AppService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	return
}

// Write a log entry to the server logs
func (r *AppService) Log(ctx context.Context, params AppLogParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "log"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List all agents
func (r *AppService) Agents(ctx context.Context, query AgentListParams, opts ...option.RequestOption) (res *[]Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List all skills
func (r *AppService) Skills(ctx context.Context, query AppSkillsParams, opts ...option.RequestOption) (res *[]Skill, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skill"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Model struct {
	ID           string                 `json:"id,required"`
	Attachment   bool                   `json:"attachment,required"`
	Cost         ModelCost              `json:"cost,required"`
	Limit        ModelLimit             `json:"limit,required"`
	Name         string                 `json:"name,required"`
	Options      map[string]interface{} `json:"options,required"`
	Reasoning    bool                   `json:"reasoning,required"`
	ReleaseDate  string                 `json:"release_date,required"`
	Temperature  bool                   `json:"temperature,required"`
	ToolCall     bool                   `json:"tool_call,required"`
	Experimental bool                   `json:"experimental"`
	Modalities   ModelModalities        `json:"modalities"`
	Provider     ModelProvider          `json:"provider"`
	Status       ModelStatus            `json:"status"`
	ProviderID   string                 `json:"providerID"`
	API          ModelAPI               `json:"api"`
	Family       string                 `json:"family"`
	Capabilities ModelCapabilities      `json:"capabilities"`
	Headers      map[string]string      `json:"headers"`
	// This field can have the runtime type of map of model variants.
	Variants interface{} `json:"variants"`
	JSON     modelJSON   `json:"-"`
}

// modelJSON contains the JSON metadata for the struct [Model]
type modelJSON struct {
	ID           apijson.Field
	Attachment   apijson.Field
	Cost         apijson.Field
	Limit        apijson.Field
	Name         apijson.Field
	Options      apijson.Field
	Reasoning    apijson.Field
	ReleaseDate  apijson.Field
	Temperature  apijson.Field
	ToolCall     apijson.Field
	Experimental apijson.Field
	Modalities   apijson.Field
	Provider     apijson.Field
	Status       apijson.Field
	ProviderID   apijson.Field
	API          apijson.Field
	Family       apijson.Field
	Capabilities apijson.Field
	Headers      apijson.Field
	Variants     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Model) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelJSON) RawJSON() string {
	return r.raw
}

type ModelCost struct {
	Input      float64 `json:"input,required"`
	Output     float64 `json:"output,required"`
	CacheRead  float64 `json:"cache_read"`
	CacheWrite float64 `json:"cache_write"`
	// This field can have the runtime type of object with read, write, input, output.
	ExperimentalOver200K interface{}   `json:"experimentalOver200K"`
	JSON                 modelCostJSON `json:"-"`
}

// modelCostJSON contains the JSON metadata for the struct [ModelCost]
type modelCostJSON struct {
	Input                apijson.Field
	Output               apijson.Field
	CacheRead            apijson.Field
	CacheWrite           apijson.Field
	ExperimentalOver200K apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ModelCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelCostJSON) RawJSON() string {
	return r.raw
}

type ModelLimit struct {
	Context float64        `json:"context,required"`
	Input   float64        `json:"input"`
	Output  float64        `json:"output,required"`
	JSON    modelLimitJSON `json:"-"`
}

// modelLimitJSON contains the JSON metadata for the struct [ModelLimit]
type modelLimitJSON struct {
	Context     apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelLimitJSON) RawJSON() string {
	return r.raw
}

type ModelModalities struct {
	Input  []ModelModalitiesInput  `json:"input,required"`
	Output []ModelModalitiesOutput `json:"output,required"`
	JSON   modelModalitiesJSON     `json:"-"`
}

// modelModalitiesJSON contains the JSON metadata for the struct [ModelModalities]
type modelModalitiesJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelModalities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelModalitiesJSON) RawJSON() string {
	return r.raw
}

type ModelModalitiesInput string

const (
	ModelModalitiesInputText  ModelModalitiesInput = "text"
	ModelModalitiesInputAudio ModelModalitiesInput = "audio"
	ModelModalitiesInputImage ModelModalitiesInput = "image"
	ModelModalitiesInputVideo ModelModalitiesInput = "video"
	ModelModalitiesInputPdf   ModelModalitiesInput = "pdf"
)

func (r ModelModalitiesInput) IsKnown() bool {
	switch r {
	case ModelModalitiesInputText, ModelModalitiesInputAudio, ModelModalitiesInputImage, ModelModalitiesInputVideo, ModelModalitiesInputPdf:
		return true
	}
	return false
}

type ModelModalitiesOutput string

const (
	ModelModalitiesOutputText  ModelModalitiesOutput = "text"
	ModelModalitiesOutputAudio ModelModalitiesOutput = "audio"
	ModelModalitiesOutputImage ModelModalitiesOutput = "image"
	ModelModalitiesOutputVideo ModelModalitiesOutput = "video"
	ModelModalitiesOutputPdf   ModelModalitiesOutput = "pdf"
)

func (r ModelModalitiesOutput) IsKnown() bool {
	switch r {
	case ModelModalitiesOutputText, ModelModalitiesOutputAudio, ModelModalitiesOutputImage, ModelModalitiesOutputVideo, ModelModalitiesOutputPdf:
		return true
	}
	return false
}

type ModelProvider struct {
	NPM  string            `json:"npm,required"`
	JSON modelProviderJSON `json:"-"`
}

// modelProviderJSON contains the JSON metadata for the struct [ModelProvider]
type modelProviderJSON struct {
	NPM         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelProviderJSON) RawJSON() string {
	return r.raw
}

type ModelAPI struct {
	ID   string       `json:"id"`
	URL  string       `json:"url"`
	NPM  string       `json:"npm"`
	JSON modelAPIJSON `json:"-"`
}

type modelAPIJSON struct {
	ID          apijson.Field
	URL         apijson.Field
	NPM         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelAPI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelAPIJSON) RawJSON() string {
	return r.raw
}

type ModelCapabilities struct {
	// This field can have the runtime type of [bool] or object.
	Interleaved interface{}           `json:"interleaved"`
	JSON        modelCapabilitiesJSON `json:"-"`
}

type modelCapabilitiesJSON struct {
	Interleaved apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelCapabilitiesJSON) RawJSON() string {
	return r.raw
}

type ModelStatus string

const (
	ModelStatusAlpha      ModelStatus = "alpha"
	ModelStatusBeta       ModelStatus = "beta"
	ModelStatusDeprecated ModelStatus = "deprecated"
	ModelStatusActive     ModelStatus = "active"
)

func (r ModelStatus) IsKnown() bool {
	switch r {
	case ModelStatusAlpha, ModelStatusBeta, ModelStatusDeprecated, ModelStatusActive:
		return true
	}
	return false
}

type AppLogParams struct {
	// Log level
	Level param.Field[AppLogParamsLevel] `json:"level,required"`
	// Log message
	Message param.Field[string] `json:"message,required"`
	// Service name for the log entry
	Service   param.Field[string] `json:"service,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	// Additional metadata for the log entry
	Extra param.Field[map[string]interface{}] `json:"extra"`
}

func (r AppLogParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AppLogParams]'s query parameters as `url.Values`.
func (r AppLogParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Log level
type AppLogParamsLevel string

const (
	AppLogParamsLevelDebug AppLogParamsLevel = "debug"
	AppLogParamsLevelInfo  AppLogParamsLevel = "info"
	AppLogParamsLevelError AppLogParamsLevel = "error"
	AppLogParamsLevelWarn  AppLogParamsLevel = "warn"
)

func (r AppLogParamsLevel) IsKnown() bool {
	switch r {
	case AppLogParamsLevelDebug, AppLogParamsLevelInfo, AppLogParamsLevelError, AppLogParamsLevelWarn:
		return true
	}
	return false
}
