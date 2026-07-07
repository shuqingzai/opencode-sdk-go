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

// V2AgentService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AgentService] method instead.
type V2AgentService struct {
	Options []option.RequestOption
}

// NewV2AgentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2AgentService(opts ...option.RequestOption) (r *V2AgentService) {
	r = &V2AgentService{}
	r.Options = opts
	return
}

// List agents
func (r *V2AgentService) List(ctx context.Context, query V2AgentListParams, opts ...option.RequestOption) (res *[]V2AgentInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V2AgentInfo struct {
	ID          string             `json:"id,required"`
	Model       ModelRef           `json:"model"`
	Request     ProviderRequest    `json:"request,required"`
	System      string             `json:"system"`
	Description string             `json:"description"`
	Mode        V2AgentInfoMode    `json:"mode,required"`
	Hidden      bool               `json:"hidden,required"`
	Color       AgentColor         `json:"color"`
	Steps       int64              `json:"steps"`
	Permissions []PermissionV2Rule `json:"permissions,required"`
	JSON        v2AgentInfoJSON    `json:"-"`
}

// v2AgentInfoJSON contains the JSON metadata for the struct [V2AgentInfo]
type v2AgentInfoJSON struct {
	ID          apijson.Field
	Model       apijson.Field
	Request     apijson.Field
	System      apijson.Field
	Description apijson.Field
	Mode        apijson.Field
	Hidden      apijson.Field
	Color       apijson.Field
	Steps       apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2AgentInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2AgentInfoJSON) RawJSON() string {
	return r.raw
}

type V2AgentInfoMode string

const (
	V2AgentInfoModeSubagent V2AgentInfoMode = "subagent"
	V2AgentInfoModePrimary  V2AgentInfoMode = "primary"
	V2AgentInfoModeAll      V2AgentInfoMode = "all"
)

func (r V2AgentInfoMode) IsKnown() bool {
	switch r {
	case V2AgentInfoModeSubagent, V2AgentInfoModePrimary, V2AgentInfoModeAll:
		return true
	}
	return false
}

type AgentColor string

const (
	AgentColorPrimary   AgentColor = "primary"
	AgentColorSecondary AgentColor = "secondary"
	AgentColorAccent    AgentColor = "accent"
	AgentColorSuccess   AgentColor = "success"
	AgentColorWarning   AgentColor = "warning"
	AgentColorError     AgentColor = "error"
	AgentColorInfo      AgentColor = "info"
)

func (r AgentColor) IsKnown() bool {
	switch r {
	case AgentColorPrimary, AgentColorSecondary, AgentColorAccent,
		AgentColorSuccess, AgentColorWarning, AgentColorError, AgentColorInfo:
		return true
	}
	return false
}

type ModelRef struct {
	ID         string       `json:"id,required"`
	ProviderID string       `json:"providerID,required"`
	Variant    string       `json:"variant"`
	JSON       modelRefJSON `json:"-"`
}

// modelRefJSON contains the JSON metadata for the struct [ModelRef]
type modelRefJSON struct {
	ID          apijson.Field
	ProviderID  apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ModelRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelRefJSON) RawJSON() string {
	return r.raw
}

type ProviderRequest struct {
	Headers map[string]string      `json:"headers,required"`
	Body    map[string]interface{} `json:"body,required"`
	JSON    providerRequestJSON    `json:"-"`
}

type providerRequestJSON struct {
	Headers     apijson.Field
	Body        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerRequestJSON) RawJSON() string {
	return r.raw
}

type PermissionV2Rule struct {
	Action   string               `json:"action,required"`
	Resource string               `json:"resource,required"`
	Effect   PermissionV2Effect   `json:"effect,required"`
	JSON     permissionV2RuleJSON `json:"-"`
}

type permissionV2RuleJSON struct {
	Action      apijson.Field
	Resource    apijson.Field
	Effect      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionV2Rule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionV2RuleJSON) RawJSON() string {
	return r.raw
}

type PermissionV2Effect string

const (
	PermissionV2EffectAllow PermissionV2Effect = "allow"
	PermissionV2EffectDeny  PermissionV2Effect = "deny"
	PermissionV2EffectAsk   PermissionV2Effect = "ask"
)

func (r PermissionV2Effect) IsKnown() bool {
	switch r {
	case PermissionV2EffectAllow, PermissionV2EffectDeny, PermissionV2EffectAsk:
		return true
	}
	return false
}

type V2AgentListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2AgentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
