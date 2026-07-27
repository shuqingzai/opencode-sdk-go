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

// This is an alias to an internal type.
type Model = ProviderModel

// This is an alias to an internal type.
type Provider = ProviderInfo

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
	Extra param.Field[map[string]any] `json:"extra"`
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
