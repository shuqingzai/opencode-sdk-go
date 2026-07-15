// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"net/url"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
)

// Agent represents an agent definition.

type Agent struct {
	Mode        AgentMode              `json:"mode,required"`
	Name        string                 `json:"name,required"`
	Options     map[string]interface{} `json:"options,required"`
	Permission  PermissionRuleset      `json:"permission,required"`
	Description string                 `json:"description"`
	Model       AgentModel             `json:"model"`
	Prompt      string                 `json:"prompt"`
	Temperature float64                `json:"temperature"`
	TopP        float64                `json:"topP"`
	Native      bool                   `json:"native"`
	Hidden      bool                   `json:"hidden"`
	Color       string                 `json:"color"`
	Variant     string                 `json:"variant"`
	Steps       int64                  `json:"steps"`
	JSON        agentJSON              `json:"-"`
}

// agentJSON contains the JSON metadata for the struct [Agent]
type agentJSON struct {
	Mode        apijson.Field
	Name        apijson.Field
	Options     apijson.Field
	Permission  apijson.Field
	Description apijson.Field
	Model       apijson.Field
	Prompt      apijson.Field
	Temperature apijson.Field
	TopP        apijson.Field
	Native      apijson.Field
	Hidden      apijson.Field
	Color       apijson.Field
	Variant     apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Agent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentJSON) RawJSON() string {
	return r.raw
}

type AgentMode string

const (
	AgentModeSubagent AgentMode = "subagent"
	AgentModePrimary  AgentMode = "primary"
	AgentModeAll      AgentMode = "all"
)

func (r AgentMode) IsKnown() bool {
	switch r {
	case AgentModeSubagent, AgentModePrimary, AgentModeAll:
		return true
	}
	return false
}

type AgentModel struct {
	ModelID    string         `json:"modelID,required"`
	ProviderID string         `json:"providerID,required"`
	JSON       agentModelJSON `json:"-"`
}

// agentModelJSON contains the JSON metadata for the struct [AgentModel]
type agentModelJSON struct {
	ModelID     apijson.Field
	ProviderID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentModelJSON) RawJSON() string {
	return r.raw
}

type AgentListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [AgentListParams]'s query parameters as `url.Values`.
func (r AgentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type Skill struct {
	Name        string    `json:"name,required"`
	Description string    `json:"description"`
	Location    string    `json:"location,required"`
	Content     string    `json:"content,required"`
	JSON        skillJSON `json:"-"`
}

// skillJSON contains the JSON metadata for the struct [Skill]
type skillJSON struct {
	Name        apijson.Field
	Description apijson.Field
	Location    apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Skill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillJSON) RawJSON() string {
	return r.raw
}

type AppSkillsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [AppSkillsParams]'s query parameters as `url.Values`.
func (r AppSkillsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
