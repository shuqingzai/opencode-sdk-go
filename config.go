// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/shared"
	"github.com/tidwall/gjson"
)

// ConfigService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// Get config info
func (r *ConfigService) Get(ctx context.Context, query ConfigGetParams, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update config
func (r *ConfigService) Update(ctx context.Context, params ConfigUpdateParams, opts ...option.RequestOption) (res *Config, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List config providers
func (r *ConfigService) Providers(ctx context.Context, query ConfigProvidersParams, opts ...option.RequestOption) (res *ConfigProvidersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "config/providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Config struct {
	// JSON schema reference for configuration validation
	Schema string `json:"$schema"`
	// Agent configuration, see https://opencode.ai/docs/agent
	Agent ConfigAgent `json:"agent"`
	// Attachment configuration for image handling
	Attachment AttachmentConfig `json:"attachment"`
	// @deprecated Use 'share' field instead. Share newly created sessions
	// automatically
	Autoshare bool `json:"autoshare"`
	// Automatically update to the latest version. Set to true to auto-update, false to disable, or 'notify' to show update notifications
	// This field can have the runtime type of [bool] or "notify".
	Autoupdate any `json:"autoupdate"`
	// Command configuration, see https://opencode.ai/docs/commands
	Command map[string]ConfigCommand `json:"command"`
	// Compaction settings for session history
	Compaction ConfigCompaction `json:"compaction"`
	// Disable providers that are loaded automatically
	DisabledProviders []string `json:"disabled_providers"`
	// Enable specific providers
	EnabledProviders []string `json:"enabled_providers"`
	// Enterprise configuration
	Enterprise   EnterpriseConfig   `json:"enterprise"`
	Experimental ConfigExperimental `json:"experimental"`
	// Enable or configure formatters. Omit or set to false to disable, true to
	// enable built-ins, or an object to enable built-ins with overrides.
	// Per OpenAPI `Config.formatter` is `boolean | object | map[string]ConfigFormatter`.
	// This field can have the runtime type of [bool], [map[string]ConfigFormatter],
	// [map[string]bool].
	Formatter any `json:"formatter"`
	// Additional instruction files or patterns to include
	Instructions []string `json:"instructions"`
	// @deprecated Always uses stretch layout.
	Layout ConfigLayout `json:"layout"`
	// Log level for the application
	LogLevel ConfigLogLevel `json:"logLevel"`
	// This field can have the runtime type of [bool], [map[string]ConfigLsp].
	Lsp any `json:"lsp"`
	// MCP (Model Context Protocol) server configurations
	Mcp map[string]ConfigMcp `json:"mcp"`
	// @deprecated Use `agent` field instead.
	Mode ConfigMode `json:"mode"`
	// Model to use in the format of provider/model, eg anthropic/claude-2
	Model string `json:"model"`
	// Permission configuration. A short string ("ask"|"allow"|"deny") or an
	// object with per-action permission rule overrides.
	// This field can have the runtime type of [PermissionActionConfig],
	// [PermissionConfigObject].
	Permission any `json:"permission"`
	// This field can have the runtime type of [string] or [][2]any{string, object}.
	Plugin []any `json:"plugin"`
	// Custom provider configurations and model overrides
	Provider map[string]ConfigProvider `json:"provider"`
	// Reference configuration for external documentation. Keys are reference
	// names, values can be a plain URL/path string or a structured config (git
	// or local).
	// Each value decodes to [ConfigV2Reference]; use [ConfigV2Reference.AsUnion]
	// to get the [ConfigV2ReferenceString], [ConfigV2ReferenceGit] or
	// [ConfigV2ReferenceLocal] variant.
	Reference map[string]ConfigV2Reference `json:"reference"`
	// References from external sources. Keys are reference names, values can be a
	// plain URL/path string or a structured config (git or local).
	// Each value decodes to [ConfigV2Reference]; use [ConfigV2Reference.AsUnion]
	// to get the [ConfigV2ReferenceString], [ConfigV2ReferenceGit] or
	// [ConfigV2ReferenceLocal] variant.
	References map[string]ConfigV2Reference `json:"references"`
	// Control sharing behavior:'manual' allows manual sharing via commands, 'auto'
	// enables automatic sharing, 'disabled' disables all sharing
	Share ConfigShare `json:"share"`
	// Shell command to use for terminal operations
	Shell string `json:"shell"`
	// Server configuration
	Server ServerConfig `json:"server"`
	// Skills configuration for paths and URLs
	Skills ConfigSkills `json:"skills"`
	// Small model to use for tasks like title generation in the format of
	// provider/model
	SmallModel string `json:"small_model"`
	Snapshot   bool   `json:"snapshot"`
	// Tool output configuration
	ToolOutput ConfigToolOutput `json:"tool_output"`
	Tools      map[string]bool  `json:"tools"`
	// Custom username to display in conversations instead of system username
	Username string        `json:"username"`
	Watcher  ConfigWatcher `json:"watcher"`
	// Default agent ID to use
	DefaultAgent string `json:"default_agent"`
	// Maximum depth for nested subagents
	SubagentDepth int64      `json:"subagent_depth"`
	JSON          configJSON `json:"-"`
	// permissionUnion holds the typed permission payload after [UnmarshalJSON]
	// routes the raw data through [PermissionConfigUnion] registered variants.
	permissionUnion PermissionConfigUnion
}

// configJSON contains the JSON metadata for the struct [Config]
type configJSON struct {
	Schema            apijson.Field
	Agent             apijson.Field
	Attachment        apijson.Field
	Autoshare         apijson.Field
	Autoupdate        apijson.Field
	Command           apijson.Field
	Compaction        apijson.Field
	DisabledProviders apijson.Field
	EnabledProviders  apijson.Field
	Enterprise        apijson.Field
	Experimental      apijson.Field
	Formatter         apijson.Field
	Instructions      apijson.Field
	Layout            apijson.Field
	LogLevel          apijson.Field
	Lsp               apijson.Field
	Mcp               apijson.Field
	Mode              apijson.Field
	Model             apijson.Field
	Permission        apijson.Field
	Plugin            apijson.Field
	Provider          apijson.Field
	Reference         apijson.Field
	References        apijson.Field
	Share             apijson.Field
	Shell             apijson.Field
	Server            apijson.Field
	Skills            apijson.Field
	SmallModel        apijson.Field
	Snapshot          apijson.Field
	ToolOutput        apijson.Field
	Tools             apijson.Field
	Username          apijson.Field
	Watcher           apijson.Field
	DefaultAgent      apijson.Field
	SubagentDepth     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Config) UnmarshalJSON(data []byte) (err error) {
	*r = Config{}
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	permissionData := gjson.GetBytes(data, "permission").Raw
	if permissionData != "" && permissionData != "null" {
		if err = apijson.UnmarshalRoot([]byte(permissionData), &r.permissionUnion); err != nil {
			return err
		}
		r.Permission = r.permissionUnion
	}
	return nil
}

func (r configJSON) RawJSON() string {
	return r.raw
}

// AsPermission returns the permission field as a typed union.
//
// Possible runtime types of the union are [PermissionActionConfig] (a short
// string "ask"|"allow"|"deny") or [PermissionConfigObject].
func (r *Config) AsPermission() PermissionConfigUnion {
	return r.permissionUnion
}

// PermissionActionConfig is the OpenAPI `PermissionActionConfig` schema: a short
// string permission, one of "ask" / "allow" / "deny".
type PermissionActionConfig string

const (
	PermissionActionConfigAsk   PermissionActionConfig = "ask"
	PermissionActionConfigAllow PermissionActionConfig = "allow"
	PermissionActionConfigDeny  PermissionActionConfig = "deny"
)

func (r PermissionActionConfig) IsKnown() bool {
	switch r {
	case PermissionActionConfigAsk, PermissionActionConfigAllow, PermissionActionConfigDeny:
		return true
	}
	return false
}

func (r *PermissionActionConfig) UnmarshalJSON(data []byte) (err error) {
	var raw string
	if err = apijson.UnmarshalRoot(data, &raw); err != nil {
		return err
	}
	*r = PermissionActionConfig(raw)
	return nil
}

func (r PermissionActionConfig) implementsPermissionConfigUnion()          {}
func (r PermissionActionConfig) implementsPermissionRuleConfigUnion()      {}
func (r PermissionActionConfig) implementsPermissionConfigUnionParam()     {}
func (r PermissionActionConfig) implementsPermissionRuleConfigUnionParam() {}

// PermissionObjectConfig is the OpenAPI `PermissionObjectConfig` schema: a map of
// pattern to [PermissionActionConfig].
type PermissionObjectConfig map[string]PermissionActionConfig

func (r PermissionObjectConfig) implementsPermissionRuleConfigUnion()      {}
func (r PermissionObjectConfig) implementsPermissionRuleConfigUnionParam() {}

// PermissionRuleConfigUnion is the OpenAPI `PermissionRuleConfig` anyOf union,
// used by the per-tool rule properties of [PermissionConfigObject].
//
// Union satisfied by [PermissionActionConfig] (a short string
// "ask"|"allow"|"deny") or [PermissionObjectConfig] (a map of pattern to action).
type PermissionRuleConfigUnion interface {
	implementsPermissionRuleConfigUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[PermissionRuleConfigUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[PermissionActionConfig](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PermissionObjectConfig](),
		},
	)
}

// PermissionRuleConfig is the carrier for the OpenAPI `PermissionRuleConfig`
// anyOf union, used by the per-tool rule properties of [PermissionConfigObject].
// The rule resolves at decode time to either [PermissionActionConfig] (a short
// string "ask"|"allow"|"deny") or [PermissionObjectConfig] (a map of pattern to
// action).
//
// The decoded union is available via [PermissionRuleConfig.AsUnion].
type PermissionRuleConfig struct {
	JSON  permissionRuleConfigJSON `json:"-"`
	union PermissionRuleConfigUnion
}

// permissionRuleConfigJSON contains the JSON metadata for the struct
// [PermissionRuleConfig]
type permissionRuleConfigJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r permissionRuleConfigJSON) RawJSON() string {
	return r.raw
}

func (r *PermissionRuleConfig) UnmarshalJSON(data []byte) (err error) {
	*r = PermissionRuleConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	r.JSON.raw = string(data)
	return nil
}

// AsUnion returns a [PermissionRuleConfigUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [PermissionActionConfig],
// [PermissionObjectConfig].
func (r PermissionRuleConfig) AsUnion() PermissionRuleConfigUnion {
	return r.union
}

// PermissionConfigUnion is the OpenAPI `PermissionConfig` anyOf union.
//
// Satisfied by [PermissionActionConfig] (a short string "ask"|"allow"|"deny") or
// [PermissionConfigObject] (per-tool permission rule overrides).
type PermissionConfigUnion interface {
	implementsPermissionConfigUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[PermissionConfigUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[PermissionActionConfig](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PermissionConfigObject](),
		},
	)
}

// PermissionConfigObject is the object variant of the OpenAPI `PermissionConfig`
// anyOf union: per-tool permission rule overrides.
//
// The ten tools that accept a matchable argument (paths, commands, skill names)
// are typed as [PermissionRuleConfig] and resolve at decode time to either
// [PermissionActionConfig] or [PermissionObjectConfig]. The remaining five tools
// take no argument and are therefore plain [PermissionActionConfig], exactly as
// the OpenAPI `PermissionConfig.anyOf[1]` schema declares.
type PermissionConfigObject struct {
	Read              PermissionRuleConfig   `json:"read"`
	Edit              PermissionRuleConfig   `json:"edit"`
	Glob              PermissionRuleConfig   `json:"glob"`
	Grep              PermissionRuleConfig   `json:"grep"`
	List              PermissionRuleConfig   `json:"list"`
	Bash              PermissionRuleConfig   `json:"bash"`
	Task              PermissionRuleConfig   `json:"task"`
	ExternalDirectory PermissionRuleConfig   `json:"external_directory"`
	Lsp               PermissionRuleConfig   `json:"lsp"`
	Skill             PermissionRuleConfig   `json:"skill"`
	Todowrite         PermissionActionConfig `json:"todowrite"`
	Question          PermissionActionConfig `json:"question"`
	Webfetch          PermissionActionConfig `json:"webfetch"`
	Websearch         PermissionActionConfig `json:"websearch"`
	DoomLoop          PermissionActionConfig `json:"doom_loop"`
	// Additional per-tool rules beyond the properties listed above, per the OpenAPI
	// `PermissionConfig.additionalProperties` -> `PermissionRuleConfig` mapping.
	ExtraFields map[string]PermissionRuleConfig `json:"-,extras"`
	JSON        permissionConfigObjectJSON      `json:"-"`
}

// permissionConfigObjectJSON contains the JSON metadata for the struct
// [PermissionConfigObject]
type permissionConfigObjectJSON struct {
	Read              apijson.Field
	Edit              apijson.Field
	Glob              apijson.Field
	Grep              apijson.Field
	List              apijson.Field
	Bash              apijson.Field
	Task              apijson.Field
	ExternalDirectory apijson.Field
	Lsp               apijson.Field
	Skill             apijson.Field
	Todowrite         apijson.Field
	Question          apijson.Field
	Webfetch          apijson.Field
	Websearch         apijson.Field
	DoomLoop          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PermissionConfigObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionConfigObjectJSON) RawJSON() string {
	return r.raw
}

func (r PermissionConfigObject) implementsPermissionConfigUnion() {}

type ConfigLogLevel string

const (
	ConfigLogLevelDebug ConfigLogLevel = "DEBUG"
	ConfigLogLevelInfo  ConfigLogLevel = "INFO"
	ConfigLogLevelWarn  ConfigLogLevel = "WARN"
	ConfigLogLevelError ConfigLogLevel = "ERROR"
)

func (r ConfigLogLevel) IsKnown() bool {
	switch r {
	case ConfigLogLevelDebug, ConfigLogLevelInfo, ConfigLogLevelWarn, ConfigLogLevelError:
		return true
	}
	return false
}

type ServerConfig struct {
	Port       int64            `json:"port"`
	Hostname   string           `json:"hostname"`
	Mdns       bool             `json:"mdns"`
	MdnsDomain string           `json:"mdnsDomain"`
	Cors       []string         `json:"cors"`
	JSON       serverConfigJSON `json:"-"`
}

type serverConfigJSON struct {
	Port        apijson.Field
	Hostname    apijson.Field
	Mdns        apijson.Field
	MdnsDomain  apijson.Field
	Cors        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r serverConfigJSON) RawJSON() string {
	return r.raw
}

type ConfigSkills struct {
	Paths []string         `json:"paths"`
	Urls  []string         `json:"urls"`
	JSON  configSkillsJSON `json:"-"`
}

type configSkillsJSON struct {
	Paths       apijson.Field
	Urls        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigSkills) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r configSkillsJSON) RawJSON() string {
	return r.raw
}

type ConfigCompaction struct {
	Auto                 bool                 `json:"auto"`
	Prune                bool                 `json:"prune"`
	Reserved             int64                `json:"reserved"`
	TailTurns            int64                `json:"tail_turns"`
	PreserveRecentTokens int64                `json:"preserve_recent_tokens"`
	JSON                 configCompactionJSON `json:"-"`
}

type configCompactionJSON struct {
	Auto                 apijson.Field
	Prune                apijson.Field
	Reserved             apijson.Field
	TailTurns            apijson.Field
	PreserveRecentTokens apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ConfigCompaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r configCompactionJSON) RawJSON() string {
	return r.raw
}

type EnterpriseConfig struct {
	URL  string               `json:"url"`
	JSON enterpriseConfigJSON `json:"-"`
}

type enterpriseConfigJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnterpriseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r enterpriseConfigJSON) RawJSON() string {
	return r.raw
}

// AgentConfig is the OpenAPI `AgentConfig` schema. It is shared by every agent
// slot of [ConfigAgent] and [ConfigMode], and by their user-defined entries.
type AgentConfig struct {
	Model       string          `json:"model"`
	Variant     string          `json:"variant"`
	Temperature float64         `json:"temperature"`
	TopP        float64         `json:"top_p"`
	Prompt      string          `json:"prompt"`
	Tools       map[string]bool `json:"tools"`
	Disable     bool            `json:"disable"`
	Description string          `json:"description"`
	Mode        AgentConfigMode `json:"mode"`
	Hidden      bool            `json:"hidden"`
	Options     map[string]any  `json:"options"`
	// Hex color code (e.g., #FF5733) or theme color (e.g., primary)
	Color    string `json:"color"`
	Steps    int64  `json:"steps"`
	MaxSteps int64  `json:"maxSteps"`
	// This field can have the runtime type of [PermissionActionConfig],
	// [PermissionConfigObject].
	Permission any `json:"permission"`
	// Additional agent properties not listed above. The OpenAPI `AgentConfig`
	// schema allows arbitrary extra properties.
	ExtraFields map[string]any  `json:"-,extras"`
	JSON        agentConfigJSON `json:"-"`
	// permissionUnion holds the typed permission payload after UnmarshalJSON
	// routes the raw data through the registered [PermissionConfigUnion] variants.
	permissionUnion PermissionConfigUnion
}

// agentConfigJSON contains the JSON metadata for the struct [AgentConfig]
type agentConfigJSON struct {
	Model       apijson.Field
	Variant     apijson.Field
	Temperature apijson.Field
	TopP        apijson.Field
	Prompt      apijson.Field
	Tools       apijson.Field
	Disable     apijson.Field
	Description apijson.Field
	Mode        apijson.Field
	Hidden      apijson.Field
	Options     apijson.Field
	Color       apijson.Field
	Steps       apijson.Field
	MaxSteps    apijson.Field
	Permission  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentConfig) UnmarshalJSON(data []byte) (err error) {
	*r = AgentConfig{}
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	permissionData := gjson.GetBytes(data, "permission").Raw
	if permissionData != "" && permissionData != "null" {
		if err = apijson.UnmarshalRoot([]byte(permissionData), &r.permissionUnion); err != nil {
			return err
		}
		r.Permission = r.permissionUnion
	}
	return nil
}

func (r agentConfigJSON) RawJSON() string {
	return r.raw
}

// AsPermission returns the permission field as a typed union.
//
// Possible runtime types of the union are [PermissionActionConfig] (a short
// string "ask"|"allow"|"deny") or [PermissionConfigObject].
func (r *AgentConfig) AsPermission() PermissionConfigUnion {
	return r.permissionUnion
}

type AgentConfigMode string

const (
	AgentConfigModeSubagent AgentConfigMode = "subagent"
	AgentConfigModePrimary  AgentConfigMode = "primary"
	AgentConfigModeAll      AgentConfigMode = "all"
)

func (r AgentConfigMode) IsKnown() bool {
	switch r {
	case AgentConfigModeSubagent, AgentConfigModePrimary, AgentConfigModeAll:
		return true
	}
	return false
}

// Agent configuration, see https://opencode.ai/docs/agent
type ConfigAgent struct {
	Build      AgentConfig `json:"build"`
	Compaction AgentConfig `json:"compaction"`
	Explore    AgentConfig `json:"explore"`
	General    AgentConfig `json:"general"`
	Plan       AgentConfig `json:"plan"`
	Summary    AgentConfig `json:"summary"`
	Title      AgentConfig `json:"title"`
	// User-defined agents keyed by agent name, per the OpenAPI
	// `Config.agent.additionalProperties` -> `AgentConfig` mapping.
	ExtraFields map[string]AgentConfig `json:"-,extras"`
	JSON        configAgentJSON        `json:"-"`
}

// configAgentJSON contains the JSON metadata for the struct [ConfigAgent]
type configAgentJSON struct {
	Build       apijson.Field
	Compaction  apijson.Field
	Explore     apijson.Field
	General     apijson.Field
	Plan        apijson.Field
	Summary     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentJSON) RawJSON() string {
	return r.raw
}

type ConfigCommand struct {
	Template    string            `json:"template,required"`
	Agent       string            `json:"agent"`
	Description string            `json:"description"`
	Model       string            `json:"model"`
	Variant     string            `json:"variant"`
	Subtask     bool              `json:"subtask"`
	JSON        configCommandJSON `json:"-"`
}

// configCommandJSON contains the JSON metadata for the struct [ConfigCommand]
type configCommandJSON struct {
	Template    apijson.Field
	Agent       apijson.Field
	Description apijson.Field
	Model       apijson.Field
	Variant     apijson.Field
	Subtask     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigCommand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configCommandJSON) RawJSON() string {
	return r.raw
}

type ConfigExperimental struct {
	BatchTool           bool                         `json:"batch_tool"`
	ContinueLoopOnDeny  bool                         `json:"continue_loop_on_deny"`
	DisablePasteSummary bool                         `json:"disable_paste_summary"`
	McpTimeout          int64                        `json:"mcp_timeout"`
	OpenTelemetry       bool                         `json:"openTelemetry"`
	Policies            []ConfigV2ExperimentalPolicy `json:"policies"`
	PrimaryTools        []string                     `json:"primary_tools"`
	JSON                configExperimentalJSON       `json:"-"`
}

// configExperimentalJSON contains the JSON metadata for the struct
// [ConfigExperimental]
type configExperimentalJSON struct {
	BatchTool           apijson.Field
	ContinueLoopOnDeny  apijson.Field
	DisablePasteSummary apijson.Field
	McpTimeout          apijson.Field
	OpenTelemetry       apijson.Field
	Policies            apijson.Field
	PrimaryTools        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ConfigExperimental) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configExperimentalJSON) RawJSON() string {
	return r.raw
}

type PolicyEffect string

const (
	PolicyEffectAllow PolicyEffect = "allow"
	PolicyEffectDeny  PolicyEffect = "deny"
)

func (r PolicyEffect) IsKnown() bool {
	switch r {
	case PolicyEffectAllow, PolicyEffectDeny:
		return true
	}
	return false
}

type ConfigV2ExperimentalPolicyAction string

const (
	ConfigV2ExperimentalPolicyActionProviderUse ConfigV2ExperimentalPolicyAction = "provider.use"
)

func (r ConfigV2ExperimentalPolicyAction) IsKnown() bool {
	switch r {
	case ConfigV2ExperimentalPolicyActionProviderUse:
		return true
	}
	return false
}

type ConfigV2ExperimentalPolicy struct {
	Action   ConfigV2ExperimentalPolicyAction `json:"action,required"`
	Effect   PolicyEffect                     `json:"effect,required"`
	Resource string                           `json:"resource,required"`
	JSON     configV2ExperimentalPolicyJSON   `json:"-"`
}

// configV2ExperimentalPolicyJSON contains the JSON metadata for the struct
// [ConfigV2ExperimentalPolicy]
type configV2ExperimentalPolicyJSON struct {
	Action      apijson.Field
	Effect      apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigV2ExperimentalPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configV2ExperimentalPolicyJSON) RawJSON() string {
	return r.raw
}

type ConfigFormatter struct {
	Command     []string            `json:"command"`
	Disabled    bool                `json:"disabled"`
	Environment map[string]string   `json:"environment"`
	Extensions  []string            `json:"extensions"`
	JSON        configFormatterJSON `json:"-"`
}

// configFormatterJSON contains the JSON metadata for the struct [ConfigFormatter]
type configFormatterJSON struct {
	Command     apijson.Field
	Disabled    apijson.Field
	Environment apijson.Field
	Extensions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigFormatter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFormatterJSON) RawJSON() string {
	return r.raw
}

// @deprecated Always uses stretch layout.
type ConfigLayout string

const (
	ConfigLayoutAuto    ConfigLayout = "auto"
	ConfigLayoutStretch ConfigLayout = "stretch"
)

func (r ConfigLayout) IsKnown() bool {
	switch r {
	case ConfigLayoutAuto, ConfigLayoutStretch:
		return true
	}
	return false
}

type ConfigLsp struct {
	// This field can have the runtime type of [[]string].
	Command  any  `json:"command"`
	Disabled bool `json:"disabled"`
	// This field can have the runtime type of [map[string]string].
	Env any `json:"env"`
	// This field can have the runtime type of [[]string].
	Extensions any `json:"extensions"`
	// This field can have the runtime type of [map[string]any].
	Initialization any           `json:"initialization"`
	JSON           configLspJSON `json:"-"`
	union          ConfigLspUnion
}

// configLspJSON contains the JSON metadata for the struct [ConfigLsp]
type configLspJSON struct {
	Command        apijson.Field
	Disabled       apijson.Field
	Env            apijson.Field
	Extensions     apijson.Field
	Initialization apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r configLspJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigLsp) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigLsp{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigLspUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ConfigLspDisabled], [ConfigLspObject].
func (r ConfigLsp) AsUnion() ConfigLspUnion {
	return r.union
}

// Union satisfied by [ConfigLspDisabled] or [ConfigLspObject].
type ConfigLspUnion interface {
	implementsConfigLsp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigLspUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigLspDisabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigLspObject](),
		},
	)
}

type ConfigLspDisabled struct {
	Disabled ConfigLspDisabledDisabled `json:"disabled,required"`
	JSON     configLspDisabledJSON     `json:"-"`
}

// configLspDisabledJSON contains the JSON metadata for the struct
// [ConfigLspDisabled]
type configLspDisabledJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigLspDisabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configLspDisabledJSON) RawJSON() string {
	return r.raw
}

func (r ConfigLspDisabled) implementsConfigLsp() {}

type ConfigLspDisabledDisabled bool

const (
	ConfigLspDisabledDisabledTrue ConfigLspDisabledDisabled = true
)

func (r ConfigLspDisabledDisabled) IsKnown() bool {
	switch r {
	case ConfigLspDisabledDisabledTrue:
		return true
	}
	return false
}

type ConfigLspObject struct {
	Command        []string            `json:"command,required"`
	Disabled       bool                `json:"disabled"`
	Env            map[string]string   `json:"env"`
	Extensions     []string            `json:"extensions"`
	Initialization map[string]any      `json:"initialization"`
	JSON           configLspObjectJSON `json:"-"`
}

// configLspObjectJSON contains the JSON metadata for the struct [ConfigLspObject]
type configLspObjectJSON struct {
	Command        apijson.Field
	Disabled       apijson.Field
	Env            apijson.Field
	Extensions     apijson.Field
	Initialization apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ConfigLspObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configLspObjectJSON) RawJSON() string {
	return r.raw
}

func (r ConfigLspObject) implementsConfigLsp() {}

type ConfigMcp struct {
	// Type of MCP server connection
	Type ConfigMcpType `json:"type,required"`
	// This field can have the runtime type of [[]string]. Command and arguments to run the MCP server (for "local" type).
	Command any `json:"command"`
	// This field can have the runtime type of [string, nil]. Working directory for the MCP server process (for "local" type).
	Cwd any `json:"cwd"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [map[string]string]. Environment variables to set when running the MCP server (for "local" type).
	Environment any `json:"environment"`
	// This field can have the runtime type of [map[string]string]. Headers to send with the request (for "remote" type).
	Headers any `json:"headers"`
	// This field can have the runtime type of [McpRemoteConfigOAuth, nil]. OAuth authentication configuration for the MCP server (for "remote" type).
	OAuth any `json:"oauth"`
	// This field can have the runtime type of [int64, nil]. Timeout in milliseconds for MCP server requests.
	Timeout any `json:"timeout"`
	// URL of the remote MCP server (for "remote" type).
	URL   string        `json:"url"`
	JSON  configMcpJSON `json:"-"`
	union ConfigMcpUnion
}

// configMcpJSON contains the JSON metadata for the struct [ConfigMcp]
type configMcpJSON struct {
	Type        apijson.Field
	Command     apijson.Field
	Cwd         apijson.Field
	Enabled     apijson.Field
	Environment apijson.Field
	Headers     apijson.Field
	OAuth       apijson.Field
	Timeout     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configMcpJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigMcp) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigMcp{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigMcpUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [McpLocalConfig], [McpRemoteConfig],
// [ConfigMcpDisabled].
func (r ConfigMcp) AsUnion() ConfigMcpUnion {
	return r.union
}

// Union satisfied by [McpLocalConfig], [McpRemoteConfig] or [ConfigMcpDisabled].
type ConfigMcpUnion interface {
	implementsConfigMcp()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigMcpUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[McpLocalConfig](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[McpRemoteConfig](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigMcpDisabled](),
		},
	)
}

// ConfigMcpDisabled represents a disabled MCP server configuration.
//
// This struct maps to the variant `{type: "object", properties: {enabled: boolean}}`
// in the OpenAPI `Config.mcp.additionalProperties.anyOf`. Use this to mark an MCP
// server as disabled without providing local or remote configuration.
type ConfigMcpDisabled struct {
	Enabled bool                  `json:"enabled,required"`
	JSON    configMcpDisabledJSON `json:"-"`
}

// configMcpDisabledJSON contains the JSON metadata for [ConfigMcpDisabled].
type configMcpDisabledJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigMcpDisabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configMcpDisabledJSON) RawJSON() string {
	return r.raw
}

func (r ConfigMcpDisabled) implementsConfigMcp() {}

// Type of MCP server connection
type ConfigMcpType string

const (
	ConfigMcpTypeLocal  ConfigMcpType = "local"
	ConfigMcpTypeRemote ConfigMcpType = "remote"
)

func (r ConfigMcpType) IsKnown() bool {
	switch r {
	case ConfigMcpTypeLocal, ConfigMcpTypeRemote:
		return true
	}
	return false
}

// @deprecated Use `agent` field instead.
type ConfigMode struct {
	Build AgentConfig `json:"build"`
	Plan  AgentConfig `json:"plan"`
	// User-defined modes keyed by mode name, per the OpenAPI
	// `Config.mode.additionalProperties` -> `AgentConfig` mapping.
	ExtraFields map[string]AgentConfig `json:"-,extras"`
	JSON        configModeJSON         `json:"-"`
}

// configModeJSON contains the JSON metadata for the struct [ConfigMode]
type configModeJSON struct {
	Build       apijson.Field
	Plan        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configModeJSON) RawJSON() string {
	return r.raw
}

type ConfigProvider struct {
	ID        string                         `json:"id"`
	API       string                         `json:"api"`
	Env       []string                       `json:"env"`
	Models    map[string]ConfigProviderModel `json:"models"`
	Name      string                         `json:"name"`
	NPM       string                         `json:"npm"`
	Whitelist []string                       `json:"whitelist"`
	Blacklist []string                       `json:"blacklist"`
	Options   ConfigProviderOptions          `json:"options"`
	JSON      configProviderJSON             `json:"-"`
}

// configProviderJSON contains the JSON metadata for the struct [ConfigProvider]
type configProviderJSON struct {
	ID          apijson.Field
	API         apijson.Field
	Env         apijson.Field
	Models      apijson.Field
	Name        apijson.Field
	NPM         apijson.Field
	Whitelist   apijson.Field
	Blacklist   apijson.Field
	Options     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModel struct {
	ID           string                   `json:"id"`
	Attachment   bool                     `json:"attachment"`
	Cost         ConfigProviderModelsCost `json:"cost"`
	Experimental bool                     `json:"experimental"`
	Family       string                   `json:"family"`
	Headers      map[string]string        `json:"headers"`
	// Per OpenAPI `ProviderConfig.models.*.interleaved` is an anyOf of four
	// variants: boolean, the enum "reasoning"|"reasoning_content"|"reasoning_text",
	// any arbitrary string, or the object `{ "field": string }` where the field
	// name identifies where interleaved reasoning content is located (known values:
	// "reasoning", "reasoning_content", "reasoning_text"; open string union).
	// This field can have the runtime type of [bool], [string], [map[string]any].
	Interleaved any                            `json:"interleaved"`
	Limit       ConfigProviderModelsLimit      `json:"limit"`
	Modalities  ConfigProviderModelsModalities `json:"modalities"`
	Name        string                         `json:"name"`
	Options     map[string]any                 `json:"options"`
	Provider    ConfigProviderModelsProvider   `json:"provider"`
	Reasoning   bool                           `json:"reasoning"`
	ReleaseDate string                         `json:"release_date"`
	Status      ConfigProviderModelsStatus     `json:"status"`
	Temperature bool                           `json:"temperature"`
	ToolCall    bool                           `json:"tool_call"`
	// This field can have the runtime type of object.
	Variants any                     `json:"variants"`
	JSON     configProviderModelJSON `json:"-"`
}

// configProviderModelJSON contains the JSON metadata for the struct
// [ConfigProviderModel]
type configProviderModelJSON struct {
	ID           apijson.Field
	Attachment   apijson.Field
	Cost         apijson.Field
	Experimental apijson.Field
	Family       apijson.Field
	Headers      apijson.Field
	Interleaved  apijson.Field
	Limit        apijson.Field
	Modalities   apijson.Field
	Name         apijson.Field
	Options      apijson.Field
	Provider     apijson.Field
	Reasoning    apijson.Field
	ReleaseDate  apijson.Field
	Status       apijson.Field
	Temperature  apijson.Field
	ToolCall     apijson.Field
	Variants     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConfigProviderModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsCost struct {
	Input           float64                                 `json:"input,required"`
	Output          float64                                 `json:"output,required"`
	CacheRead       float64                                 `json:"cache_read"`
	CacheWrite      float64                                 `json:"cache_write"`
	ContextOver200k ConfigProviderModelsCostContextOver200k `json:"context_over_200k"`
	JSON            configProviderModelsCostJSON            `json:"-"`
}

type ConfigProviderModelsCostContextOver200k struct {
	Input      float64                                     `json:"input,required"`
	Output     float64                                     `json:"output,required"`
	CacheRead  float64                                     `json:"cache_read"`
	CacheWrite float64                                     `json:"cache_write"`
	JSON       configProviderModelsCostContextOver200kJSON `json:"-"`
}

// configProviderModelsCostContextOver200kJSON contains the JSON metadata for the struct
// [ConfigProviderModelsCostContextOver200k]
type configProviderModelsCostContextOver200kJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	CacheRead   apijson.Field
	CacheWrite  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProviderModelsCostContextOver200k) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsCostContextOver200kJSON) RawJSON() string {
	return r.raw
}

// configProviderModelsCostJSON contains the JSON metadata for the struct
// [ConfigProviderModelsCost]
type configProviderModelsCostJSON struct {
	Input           apijson.Field
	Output          apijson.Field
	CacheRead       apijson.Field
	CacheWrite      apijson.Field
	ContextOver200k apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigProviderModelsCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsCostJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsLimit struct {
	Context int64                         `json:"context,required"`
	Input   int64                         `json:"input"`
	Output  int64                         `json:"output,required"`
	JSON    configProviderModelsLimitJSON `json:"-"`
}

// configProviderModelsLimitJSON contains the JSON metadata for the struct
// [ConfigProviderModelsLimit]
type configProviderModelsLimitJSON struct {
	Context     apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProviderModelsLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsLimitJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsModalities struct {
	Input  []ConfigProviderModelsModalitiesInput  `json:"input"`
	Output []ConfigProviderModelsModalitiesOutput `json:"output"`
	JSON   configProviderModelsModalitiesJSON     `json:"-"`
}

// configProviderModelsModalitiesJSON contains the JSON metadata for the struct
// [ConfigProviderModelsModalities]
type configProviderModelsModalitiesJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProviderModelsModalities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsModalitiesJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsModalitiesInput string

const (
	ConfigProviderModelsModalitiesInputText  ConfigProviderModelsModalitiesInput = "text"
	ConfigProviderModelsModalitiesInputAudio ConfigProviderModelsModalitiesInput = "audio"
	ConfigProviderModelsModalitiesInputImage ConfigProviderModelsModalitiesInput = "image"
	ConfigProviderModelsModalitiesInputVideo ConfigProviderModelsModalitiesInput = "video"
	ConfigProviderModelsModalitiesInputPdf   ConfigProviderModelsModalitiesInput = "pdf"
)

func (r ConfigProviderModelsModalitiesInput) IsKnown() bool {
	switch r {
	case ConfigProviderModelsModalitiesInputText, ConfigProviderModelsModalitiesInputAudio, ConfigProviderModelsModalitiesInputImage, ConfigProviderModelsModalitiesInputVideo, ConfigProviderModelsModalitiesInputPdf:
		return true
	}
	return false
}

type ConfigProviderModelsModalitiesOutput string

const (
	ConfigProviderModelsModalitiesOutputText  ConfigProviderModelsModalitiesOutput = "text"
	ConfigProviderModelsModalitiesOutputAudio ConfigProviderModelsModalitiesOutput = "audio"
	ConfigProviderModelsModalitiesOutputImage ConfigProviderModelsModalitiesOutput = "image"
	ConfigProviderModelsModalitiesOutputVideo ConfigProviderModelsModalitiesOutput = "video"
	ConfigProviderModelsModalitiesOutputPdf   ConfigProviderModelsModalitiesOutput = "pdf"
)

func (r ConfigProviderModelsModalitiesOutput) IsKnown() bool {
	switch r {
	case ConfigProviderModelsModalitiesOutputText, ConfigProviderModelsModalitiesOutputAudio, ConfigProviderModelsModalitiesOutputImage, ConfigProviderModelsModalitiesOutputVideo, ConfigProviderModelsModalitiesOutputPdf:
		return true
	}
	return false
}

type ConfigProviderModelsProvider struct {
	NPM  string                           `json:"npm"`
	API  string                           `json:"api"`
	JSON configProviderModelsProviderJSON `json:"-"`
}

// configProviderModelsProviderJSON contains the JSON metadata for the struct
// [ConfigProviderModelsProvider]
type configProviderModelsProviderJSON struct {
	NPM         apijson.Field
	API         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProviderModelsProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderModelsProviderJSON) RawJSON() string {
	return r.raw
}

type ConfigProviderModelsStatus string

const (
	ConfigProviderModelsStatusAlpha      ConfigProviderModelsStatus = "alpha"
	ConfigProviderModelsStatusBeta       ConfigProviderModelsStatus = "beta"
	ConfigProviderModelsStatusDeprecated ConfigProviderModelsStatus = "deprecated"
	ConfigProviderModelsStatusActive     ConfigProviderModelsStatus = "active"
)

func (r ConfigProviderModelsStatus) IsKnown() bool {
	switch r {
	case ConfigProviderModelsStatusAlpha, ConfigProviderModelsStatusBeta, ConfigProviderModelsStatusDeprecated, ConfigProviderModelsStatusActive:
		return true
	}
	return false
}

type ConfigProviderOptions struct {
	APIKey        string `json:"apiKey"`
	BaseURL       string `json:"baseURL"`
	EnterpriseURL string `json:"enterpriseUrl"`
	SetCacheKey   bool   `json:"setCacheKey"`
	// Timeout in milliseconds for full requests to this provider. Set to false to
	// disable timeout.
	//
	// Decodes to [ConfigProviderOptionsTimeout]; use
	// [ConfigProviderOptionsTimeout.AsUnion] to get [shared.UnionInt] (an int64
	// millisecond duration) or [shared.UnionBool] (always false).
	Timeout ConfigProviderOptionsTimeout `json:"timeout"`
	// Timeout in milliseconds to wait for response headers. Provider integrations
	// may set defaults. Set to false to disable timeout.
	//
	// Decodes to [ConfigProviderOptionsTimeout]; use
	// [ConfigProviderOptionsTimeout.AsUnion] to get [shared.UnionInt] (an int64
	// millisecond duration) or [shared.UnionBool] (always false).
	HeaderTimeout ConfigProviderOptionsTimeout `json:"headerTimeout"`
	ChunkTimeout  int64                        `json:"chunkTimeout"`
	ExtraFields   map[string]any               `json:"-,extras"`
	JSON          configProviderOptionsJSON    `json:"-"`
}

// configProviderOptionsJSON contains the JSON metadata for the struct
// [ConfigProviderOptions]
type configProviderOptionsJSON struct {
	APIKey        apijson.Field
	BaseURL       apijson.Field
	EnterpriseURL apijson.Field
	SetCacheKey   apijson.Field
	Timeout       apijson.Field
	HeaderTimeout apijson.Field
	ChunkTimeout  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConfigProviderOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderOptionsJSON) RawJSON() string {
	return r.raw
}

// ConfigProviderOptionsTimeoutUnion represents a timeout duration as either an
// integer (milliseconds, > 0) or false (to disable). Used by
// [ConfigProviderOptions.Timeout] and [ConfigProviderOptions.HeaderTimeout].
//
// Union satisfied by [shared.UnionInt] or [shared.UnionBool].
type ConfigProviderOptionsTimeoutUnion interface {
	ImplementsConfigProviderOptionsTimeoutUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigProviderOptionsTimeoutUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeFor[shared.UnionInt](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[shared.UnionBool](),
		},
	)
}

// ConfigProviderOptionsTimeout is the carrier for the OpenAPI
// `ProviderConfig.options.timeout` / `headerTimeout` anyOf union: an integer
// millisecond duration or false (to disable).
//
// The decoded union is available via [ConfigProviderOptionsTimeout.AsUnion].
type ConfigProviderOptionsTimeout struct {
	JSON  configProviderOptionsTimeoutJSON `json:"-"`
	union ConfigProviderOptionsTimeoutUnion
}

// configProviderOptionsTimeoutJSON contains the JSON metadata for the struct
// [ConfigProviderOptionsTimeout]
type configProviderOptionsTimeoutJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configProviderOptionsTimeoutJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigProviderOptionsTimeout) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigProviderOptionsTimeout{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	r.JSON.raw = string(data)
	return nil
}

// AsUnion returns a [ConfigProviderOptionsTimeoutUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.UnionInt] or
// [shared.UnionBool].
func (r ConfigProviderOptionsTimeout) AsUnion() ConfigProviderOptionsTimeoutUnion {
	return r.union
}

// Control sharing behavior:'manual' allows manual sharing via commands, 'auto'
// enables automatic sharing, 'disabled' disables all sharing
type ConfigShare string

const (
	ConfigShareManual   ConfigShare = "manual"
	ConfigShareAuto     ConfigShare = "auto"
	ConfigShareDisabled ConfigShare = "disabled"
)

func (r ConfigShare) IsKnown() bool {
	switch r {
	case ConfigShareManual, ConfigShareAuto, ConfigShareDisabled:
		return true
	}
	return false
}

type ConfigWatcher struct {
	Ignore []string          `json:"ignore"`
	JSON   configWatcherJSON `json:"-"`
}

// configWatcherJSON contains the JSON metadata for the struct [ConfigWatcher]
type configWatcherJSON struct {
	Ignore      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigWatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configWatcherJSON) RawJSON() string {
	return r.raw
}

type McpLocalConfig struct {
	// Command and arguments to run the MCP server
	Command []string `json:"command,required"`
	Cwd     string   `json:"cwd"`
	// Type of MCP server connection
	Type McpLocalConfigType `json:"type,required"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// Timeout in milliseconds
	Timeout int64 `json:"timeout"`
	// Environment variables to set when running the MCP server
	Environment map[string]string  `json:"environment"`
	JSON        mcpLocalConfigJSON `json:"-"`
}

// mcpLocalConfigJSON contains the JSON metadata for the struct [McpLocalConfig]
type mcpLocalConfigJSON struct {
	Command     apijson.Field
	Cwd         apijson.Field
	Type        apijson.Field
	Enabled     apijson.Field
	Timeout     apijson.Field
	Environment apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpLocalConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpLocalConfigJSON) RawJSON() string {
	return r.raw
}

func (r McpLocalConfig) implementsConfigMcp() {}

// Type of MCP server connection
type McpLocalConfigType string

const (
	McpLocalConfigTypeLocal McpLocalConfigType = "local"
)

func (r McpLocalConfigType) IsKnown() bool {
	switch r {
	case McpLocalConfigTypeLocal:
		return true
	}
	return false
}

type McpRemoteConfig struct {
	// Type of MCP server connection
	Type McpRemoteConfigType `json:"type,required"`
	// URL of the remote MCP server
	URL string `json:"url,required"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// Timeout in milliseconds
	Timeout int64 `json:"timeout"`
	// Headers to send with the request
	Headers map[string]string `json:"headers"`
	// OAuth authentication configuration for the MCP server. Set to false to
	// disable OAuth auto-detection.
	//
	// Decodes to [McpRemoteConfigOAuth]; use [McpRemoteConfigOAuth.AsUnion] to get
	// [McpOAuthConfig] or [shared.UnionBool] (always false).
	OAuth McpRemoteConfigOAuth `json:"oauth"`
	JSON  mcpRemoteConfigJSON  `json:"-"`
}

// mcpRemoteConfigJSON contains the JSON metadata for the struct [McpRemoteConfig]
type mcpRemoteConfigJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	Enabled     apijson.Field
	Timeout     apijson.Field
	Headers     apijson.Field
	OAuth       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpRemoteConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpRemoteConfigJSON) RawJSON() string {
	return r.raw
}

func (r McpRemoteConfig) implementsConfigMcp() {}

// Type of MCP server connection
type McpRemoteConfigType string

const (
	McpRemoteConfigTypeRemote McpRemoteConfigType = "remote"
)

func (r McpRemoteConfigType) IsKnown() bool {
	switch r {
	case McpRemoteConfigTypeRemote:
		return true
	}
	return false
}

// McpRemoteConfigOAuth is the carrier for the OpenAPI `McpRemoteConfig.oauth`
// anyOf union: a complete OAuth config ([McpOAuthConfig]) or false (to disable
// OAuth auto-detection).
//
// The decoded union is available via [McpRemoteConfigOAuth.AsUnion].
type McpRemoteConfigOAuth struct {
	// OAuth client ID. If not provided, dynamic client registration (RFC 7591) will be attempted.
	ClientID string `json:"clientId"`
	// OAuth client secret (if required by the authorization server)
	ClientSecret string `json:"clientSecret"`
	// OAuth scopes to request during authorization
	Scope string `json:"scope"`
	// OAuth callback port for the local HTTP server
	CallbackPort int64 `json:"callbackPort"`
	// OAuth redirect URI
	RedirectURI string                   `json:"redirectUri"`
	JSON        mcpRemoteConfigOAuthJSON `json:"-"`
	union       McpOAuthConfigUnion
}

// mcpRemoteConfigOAuthJSON contains the JSON metadata for the struct
// [McpRemoteConfigOAuth]
type mcpRemoteConfigOAuthJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	Scope        apijson.Field
	CallbackPort apijson.Field
	RedirectURI  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r mcpRemoteConfigOAuthJSON) RawJSON() string {
	return r.raw
}

func (r *McpRemoteConfigOAuth) UnmarshalJSON(data []byte) (err error) {
	*r = McpRemoteConfigOAuth{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	if _, ok := r.union.(shared.UnionBool); ok {
		r.JSON.raw = string(data)
		return nil
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [McpOAuthConfigUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [McpOAuthConfig], [shared.UnionBool].
func (r McpRemoteConfigOAuth) AsUnion() McpOAuthConfigUnion {
	return r.union
}

type ConfigGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ConfigGetParams]'s query parameters as `url.Values`.
func (r ConfigGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigUpdateParams struct {
	// Query parameters
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	// Body parameters — all Config fields as optional
	Schema     param.Field[string]                `json:"$schema"`
	Agent      param.Field[ConfigAgentParam]      `json:"agent"`
	Attachment param.Field[AttachmentConfigParam] `json:"attachment"`
	Autoshare  param.Field[bool]                  `json:"autoshare"`
	// Automatically update to the latest version. Pass true to auto-update,
	// false to disable, or "notify" to show update notifications.
	// Accepts [ConfigAutoupdateEnabled] or [ConfigAutoupdateNotify].
	Autoupdate        param.Field[ConfigAutoupdateUnionParam]    `json:"autoupdate"`
	Command           param.Field[map[string]ConfigCommandParam] `json:"command"`
	Compaction        param.Field[ConfigCompactionParam]         `json:"compaction"`
	DisabledProviders param.Field[[]string]                      `json:"disabled_providers"`
	EnabledProviders  param.Field[[]string]                      `json:"enabled_providers"`
	Enterprise        param.Field[EnterpriseConfigParam]         `json:"enterprise"`
	Experimental      param.Field[ConfigExperimentalParam]       `json:"experimental"`
	// Enable or configure formatters. Pass false to disable, true to enable
	// built-ins, or a map of formatter-name to config to enable with overrides.
	// Accepts [ConfigFormatterEnabled] or [ConfigFormatterMapParam].
	Formatter    param.Field[ConfigFormatterUnionParam] `json:"formatter"`
	Instructions param.Field[[]string]                  `json:"instructions"`
	Layout       param.Field[ConfigLayout]              `json:"layout"`
	LogLevel     param.Field[ConfigLogLevel]            `json:"logLevel"`
	// Enable or configure LSP servers. Pass false to disable, true to enable
	// built-ins, or a map of lsp-name to config to enable with overrides.
	// Accepts [ConfigLspEnabled] or [ConfigLspMapParam].
	Lsp   param.Field[ConfigLspUnionParam]            `json:"lsp"`
	Mcp   param.Field[map[string]ConfigMcpUnionParam] `json:"mcp"`
	Mode  param.Field[ConfigModeParam]                `json:"mode"`
	Model param.Field[string]                         `json:"model"`
	// Permission configuration. A short string ("ask"|"allow"|"deny") or an
	// object with per-action permission rule overrides. Accepts
	// [PermissionActionConfig] (a string constant) or [PermissionConfigObjectParam].
	Permission param.Field[PermissionConfigUnionParam] `json:"permission"`
	// Plugins to load. Each item is a [ConfigPluginName] (a plugin name) or a
	// [ConfigPluginTupleParam] 2-tuple of [pluginName, configObject].
	Plugin   param.Field[[]ConfigPluginItemUnionParam]   `json:"plugin"`
	Provider param.Field[map[string]ConfigProviderParam] `json:"provider"`
	// Map of reference name → value. Each value can be a [ConfigV2ReferenceString]
	// (URL/path), a [ConfigV2ReferenceGitParam], or a [ConfigV2ReferenceLocalParam].
	Reference param.Field[map[string]ConfigV2ReferenceUnionParam] `json:"reference"`
	// Map of reference name → value. Each value can be a [ConfigV2ReferenceString]
	// (URL/path), a [ConfigV2ReferenceGitParam], or a [ConfigV2ReferenceLocalParam].
	References    param.Field[map[string]ConfigV2ReferenceUnionParam] `json:"references"`
	Share         param.Field[ConfigShare]                            `json:"share"`
	Shell         param.Field[string]                                 `json:"shell"`
	Server        param.Field[ServerConfigParam]                      `json:"server"`
	Skills        param.Field[ConfigSkillsParam]                      `json:"skills"`
	SmallModel    param.Field[string]                                 `json:"small_model"`
	Snapshot      param.Field[bool]                                   `json:"snapshot"`
	ToolOutput    param.Field[ConfigToolOutputParam]                  `json:"tool_output"`
	Tools         param.Field[map[string]bool]                        `json:"tools"`
	Username      param.Field[string]                                 `json:"username"`
	Watcher       param.Field[ConfigWatcherParam]                     `json:"watcher"`
	DefaultAgent  param.Field[string]                                 `json:"default_agent"`
	SubagentDepth param.Field[int64]                                  `json:"subagent_depth"`
}

func (r ConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ConfigUpdateParams]'s query parameters as `url.Values`.
func (r ConfigUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ConfigAutoupdateEnabled is the boolean variant of the OpenAPI
// `Config.autoupdate` anyOf union.
type ConfigAutoupdateEnabled bool

func (r ConfigAutoupdateEnabled) implementsConfigAutoupdateUnionParam() {}

// ConfigAutoupdateNotify is the string "notify" variant of the OpenAPI
// `Config.autoupdate` anyOf union.
type ConfigAutoupdateNotify string

func (r ConfigAutoupdateNotify) implementsConfigAutoupdateUnionParam() {}

// ConfigAutoupdateUnionParam is the request-side union for the OpenAPI
// `Config.autoupdate` anyOf (boolean | "notify").
//
// Satisfied by [ConfigAutoupdateEnabled], [ConfigAutoupdateNotify].
type ConfigAutoupdateUnionParam interface {
	implementsConfigAutoupdateUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigAutoupdateUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeFor[ConfigAutoupdateEnabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[ConfigAutoupdateEnabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[ConfigAutoupdateNotify](),
		},
	)
}

// ConfigFormatterEnabled is the boolean variant of the OpenAPI
// `Config.formatter` anyOf union.
type ConfigFormatterEnabled bool

func (r ConfigFormatterEnabled) implementsConfigFormatterUnionParam() {}

// ConfigFormatterParam is the request-side counterpart of [ConfigFormatter].
type ConfigFormatterParam struct {
	Disabled    param.Field[bool]              `json:"disabled"`
	Command     param.Field[[]string]          `json:"command"`
	Environment param.Field[map[string]string] `json:"environment"`
	Extensions  param.Field[[]string]          `json:"extensions"`
}

func (r ConfigFormatterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigFormatterMapParam is the object variant of the OpenAPI
// `Config.formatter` anyOf union: a map of formatter-name to per-formatter
// config.
type ConfigFormatterMapParam map[string]ConfigFormatterParam

func (r ConfigFormatterMapParam) implementsConfigFormatterUnionParam() {}

// ConfigFormatterUnionParam is the request-side union for the OpenAPI
// `Config.formatter` anyOf (boolean | map[string]ConfigFormatter).
//
// Satisfied by [ConfigFormatterEnabled], [ConfigFormatterMapParam].
type ConfigFormatterUnionParam interface {
	implementsConfigFormatterUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigFormatterUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeFor[ConfigFormatterEnabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[ConfigFormatterEnabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigFormatterMapParam](),
		},
	)
}

// ConfigLspEnabled is the boolean variant of the OpenAPI `Config.lsp` anyOf
// union.
type ConfigLspEnabled bool

func (r ConfigLspEnabled) implementsConfigLspUnionParam() {}

// ConfigLspDisabledParam is the request-side counterpart of [ConfigLspDisabled].
type ConfigLspDisabledParam struct {
	Disabled param.Field[ConfigLspDisabledDisabled] `json:"disabled,required"`
}

func (r ConfigLspDisabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigLspDisabledParam) implementsConfigLspServerUnionParam() {}

// ConfigLspObjectParam is the request-side counterpart of [ConfigLspObject].
type ConfigLspObjectParam struct {
	Command        param.Field[[]string]          `json:"command,required"`
	Disabled       param.Field[bool]              `json:"disabled"`
	Env            param.Field[map[string]string] `json:"env"`
	Extensions     param.Field[[]string]          `json:"extensions"`
	Initialization param.Field[map[string]any]    `json:"initialization"`
}

func (r ConfigLspObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigLspObjectParam) implementsConfigLspServerUnionParam() {}

// ConfigLspServerUnionParam is the per-server union of the OpenAPI
// `Config.lsp` object variant (disabled | command-based config).
//
// Satisfied by [ConfigLspDisabledParam], [ConfigLspObjectParam].
type ConfigLspServerUnionParam interface {
	implementsConfigLspServerUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigLspServerUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigLspDisabledParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigLspObjectParam](),
		},
	)
}

// ConfigLspMapParam is the object variant of the OpenAPI `Config.lsp` anyOf
// union: a map of lsp-name to per-server config.
type ConfigLspMapParam map[string]ConfigLspServerUnionParam

func (r ConfigLspMapParam) implementsConfigLspUnionParam() {}

// ConfigLspUnionParam is the request-side union for the OpenAPI `Config.lsp`
// anyOf (boolean | map[string]ConfigLsp).
//
// Satisfied by [ConfigLspEnabled], [ConfigLspMapParam].
type ConfigLspUnionParam interface {
	implementsConfigLspUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigLspUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeFor[ConfigLspEnabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[ConfigLspEnabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigLspMapParam](),
		},
	)
}

// ConfigPluginName is the string variant of the OpenAPI `Config.plugin` item
// anyOf union: a bare plugin name.
type ConfigPluginName string

func (r ConfigPluginName) implementsConfigPluginItemUnionParam() {}

// ConfigPluginTupleParam is the 2-tuple variant of the OpenAPI `Config.plugin`
// item anyOf union: [pluginName, configObject].
//
// It serializes as the OpenAPI `prefixItems` tuple `[pluginName, configObject]`
// (a JSON array, not an object), so [MarshalJSON] emits `[Name.Value,
// Config.Value]` and the struct json tags are not used. Both [Name] and
// [Config] must be set; leaving either unset emits its zero value.
type ConfigPluginTupleParam struct {
	Name   param.Field[string]
	Config param.Field[map[string]any]
}

func (r ConfigPluginTupleParam) MarshalJSON() (data []byte, err error) {
	return apijson.Marshal([]any{r.Name.Value, r.Config.Value})
}

func (r ConfigPluginTupleParam) implementsConfigPluginItemUnionParam() {}

// ConfigPluginItemUnionParam is the request-side union for each item of the
// OpenAPI `Config.plugin` array (string | [string, object]).
//
// Satisfied by [ConfigPluginName], [ConfigPluginTupleParam].
type ConfigPluginItemUnionParam interface {
	implementsConfigPluginItemUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigPluginItemUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[ConfigPluginName](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigPluginTupleParam](),
		},
	)
}

// ConfigV2ReferenceGitParam is the request-side counterpart of
// [ConfigV2ReferenceGit].
type ConfigV2ReferenceGitParam struct {
	Repository  param.Field[string] `json:"repository,required"`
	Branch      param.Field[string] `json:"branch"`
	Description param.Field[string] `json:"description"`
	Hidden      param.Field[bool]   `json:"hidden"`
}

func (r ConfigV2ReferenceGitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigV2ReferenceGitParam) implementsConfigV2ReferenceUnionParam() {}

// ConfigV2ReferenceLocalParam is the request-side counterpart of
// [ConfigV2ReferenceLocal].
type ConfigV2ReferenceLocalParam struct {
	Path        param.Field[string] `json:"path,required"`
	Description param.Field[string] `json:"description"`
	Hidden      param.Field[bool]   `json:"hidden"`
}

func (r ConfigV2ReferenceLocalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigV2ReferenceLocalParam) implementsConfigV2ReferenceUnionParam() {}

// ConfigV2ReferenceUnionParam is the request-side union for each value of the
// OpenAPI `Config.reference` / `Config.references` object
// (string | ConfigV2ReferenceGit | ConfigV2ReferenceLocal).
//
// Satisfied by [ConfigV2ReferenceString], [ConfigV2ReferenceGitParam],
// [ConfigV2ReferenceLocalParam].
type ConfigV2ReferenceUnionParam interface {
	implementsConfigV2ReferenceUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigV2ReferenceUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[ConfigV2ReferenceString](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigV2ReferenceGitParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigV2ReferenceLocalParam](),
		},
	)
}

// ConfigProviderModelsInterleavedEnabled is the boolean variant of the OpenAPI
// `ProviderConfig.models.*.interleaved` anyOf union.
type ConfigProviderModelsInterleavedEnabled bool

func (r ConfigProviderModelsInterleavedEnabled) implementsConfigProviderModelsInterleavedUnionParam() {
}

// ConfigProviderModelsInterleavedString is the string variant of the OpenAPI
// `ProviderConfig.models.*.interleaved` anyOf union (covers both the known enum
// values and arbitrary vendor strings).
type ConfigProviderModelsInterleavedString string

func (r ConfigProviderModelsInterleavedString) implementsConfigProviderModelsInterleavedUnionParam() {
}

// ConfigProviderModelsInterleavedUnionParam is the request-side union for the
// OpenAPI `ProviderConfig.models.*.interleaved` anyOf
// (boolean | string | { "field": string }).
//
// Satisfied by [ConfigProviderModelsInterleavedEnabled],
// [ConfigProviderModelsInterleavedString], [ConfigProviderModelsInterleavedFieldParam].
type ConfigProviderModelsInterleavedUnionParam interface {
	implementsConfigProviderModelsInterleavedUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigProviderModelsInterleavedUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeFor[ConfigProviderModelsInterleavedEnabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[ConfigProviderModelsInterleavedEnabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[ConfigProviderModelsInterleavedString](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigProviderModelsInterleavedFieldParam](),
		},
	)
}

type ConfigProvidersParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [ConfigProvidersParams]'s query parameters as `url.Values`.
func (r ConfigProvidersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigProvidersResponse struct {
	Default   map[string]string           `json:"default,required"`
	Providers []ProviderInfo              `json:"providers,required"`
	JSON      configProvidersResponseJSON `json:"-"`
}

// configProvidersResponseJSON contains the JSON metadata for the struct [ConfigProvidersResponse]
type configProvidersResponseJSON struct {
	Default     apijson.Field
	Providers   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigProvidersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProvidersResponseJSON) RawJSON() string {
	return r.raw
}

// Attachment configuration for image handling
type AttachmentConfig struct {
	// Image attachment configuration
	Image ImageAttachmentConfig `json:"image"`
	JSON  attachmentConfigJSON  `json:"-"`
}

// attachmentConfigJSON contains the JSON metadata for the struct [AttachmentConfig]
type attachmentConfigJSON struct {
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttachmentConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attachmentConfigJSON) RawJSON() string {
	return r.raw
}

// Image attachment configuration
type ImageAttachmentConfig struct {
	// Automatically resize images before sending
	AutoResize bool `json:"auto_resize"`
	// Maximum image width in pixels
	MaxWidth int64 `json:"max_width"`
	// Maximum image height in pixels
	MaxHeight int64 `json:"max_height"`
	// Maximum base64 encoded image size in bytes
	MaxBase64Bytes int64                     `json:"max_base64_bytes"`
	JSON           imageAttachmentConfigJSON `json:"-"`
}

// imageAttachmentConfigJSON contains the JSON metadata for the struct [ImageAttachmentConfig]
type imageAttachmentConfigJSON struct {
	AutoResize     apijson.Field
	MaxWidth       apijson.Field
	MaxHeight      apijson.Field
	MaxBase64Bytes apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ImageAttachmentConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageAttachmentConfigJSON) RawJSON() string {
	return r.raw
}

type ConfigV2ReferenceUnion interface {
	implementsConfigV2ReferenceUnion()
}

type ConfigV2ReferenceGit struct {
	// Git repository URL, host/path reference, or GitHub owner/repo shorthand
	Repository string `json:"repository,required"`
	// Branch to reference
	Branch string `json:"branch"`
	// Human-readable description of the reference
	Description string `json:"description"`
	// Whether to hide this reference from listings
	Hidden bool                     `json:"hidden"`
	JSON   configV2ReferenceGitJSON `json:"-"`
}

// configV2ReferenceGitJSON contains the JSON metadata for the struct [ConfigV2ReferenceGit]
type configV2ReferenceGitJSON struct {
	Repository  apijson.Field
	Branch      apijson.Field
	Description apijson.Field
	Hidden      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigV2ReferenceGit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configV2ReferenceGitJSON) RawJSON() string {
	return r.raw
}

func (r ConfigV2ReferenceGit) implementsConfigV2ReferenceUnion() {}

type ConfigV2ReferenceLocal struct {
	// Absolute path, ~/ path, or workspace-relative path to a local reference directory
	Path string `json:"path,required"`
	// Human-readable description of the reference
	Description string `json:"description"`
	// Whether to hide this reference from listings
	Hidden bool                       `json:"hidden"`
	JSON   configV2ReferenceLocalJSON `json:"-"`
}

// configV2ReferenceLocalJSON contains the JSON metadata for the struct [ConfigV2ReferenceLocal]
type configV2ReferenceLocalJSON struct {
	Path        apijson.Field
	Description apijson.Field
	Hidden      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigV2ReferenceLocal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configV2ReferenceLocalJSON) RawJSON() string {
	return r.raw
}

func (r ConfigV2ReferenceLocal) implementsConfigV2ReferenceUnion() {}

// ConfigV2ReferenceString is the bare string variant of
// [ConfigV2ReferenceUnion], e.g. a GitHub "owner/repo" shorthand or a local
// path.
type ConfigV2ReferenceString string

func (r ConfigV2ReferenceString) implementsConfigV2ReferenceUnion()      {}
func (r ConfigV2ReferenceString) implementsConfigV2ReferenceUnionParam() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigV2ReferenceUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[ConfigV2ReferenceString](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigV2ReferenceGit](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigV2ReferenceLocal](),
		},
	)
}

// ConfigV2Reference is the carrier for the OpenAPI `Config.reference` /
// `Config.references` anyOf union: a bare string (URL/path shorthand), a git
// reference ([ConfigV2ReferenceGit]) or a local reference
// ([ConfigV2ReferenceLocal]).
//
// The decoded union is available via [ConfigV2Reference.AsUnion].
type ConfigV2Reference struct {
	// Git repository URL, host/path reference, or GitHub owner/repo shorthand
	Repository string `json:"repository"`
	// Branch to reference
	Branch string `json:"branch"`
	// Absolute path, ~/ path, or workspace-relative path to a local reference directory
	Path string `json:"path"`
	// Human-readable description of the reference
	Description string `json:"description"`
	// Whether to hide this reference from listings
	Hidden bool                  `json:"hidden"`
	JSON   configV2ReferenceJSON `json:"-"`
	union  ConfigV2ReferenceUnion
}

// configV2ReferenceJSON contains the JSON metadata for the struct
// [ConfigV2Reference]
type configV2ReferenceJSON struct {
	Repository  apijson.Field
	Branch      apijson.Field
	Path        apijson.Field
	Description apijson.Field
	Hidden      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configV2ReferenceJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigV2Reference) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigV2Reference{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	if _, ok := r.union.(ConfigV2ReferenceString); ok {
		r.JSON.raw = string(data)
		return nil
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigV2ReferenceUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ConfigV2ReferenceString],
// [ConfigV2ReferenceGit], [ConfigV2ReferenceLocal].
func (r ConfigV2Reference) AsUnion() ConfigV2ReferenceUnion {
	return r.union
}

// Tool output configuration
type ConfigToolOutput struct {
	// Maximum number of lines to display in tool output
	MaxLines int64 `json:"max_lines"`
	// Maximum number of bytes to display in tool output
	MaxBytes int64                `json:"max_bytes"`
	JSON     configToolOutputJSON `json:"-"`
}

// configToolOutputJSON contains the JSON metadata for the struct [ConfigToolOutput]
type configToolOutputJSON struct {
	MaxLines    apijson.Field
	MaxBytes    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigToolOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configToolOutputJSON) RawJSON() string {
	return r.raw
}

// ---------------------------------------------------------------------------
// Request (Param) types — one per Response struct that appears inside
// param.Field[T] in ConfigUpdateParams / GlobalConfigUpdateParams.
// All fields are param.Field[T] so unset optionals are omitted from PATCH bodies.
// ---------------------------------------------------------------------------

// ServerConfigParam is the request-side counterpart of [ServerConfig].
type ServerConfigParam struct {
	Port       param.Field[int64]    `json:"port"`
	Hostname   param.Field[string]   `json:"hostname"`
	Mdns       param.Field[bool]     `json:"mdns"`
	MdnsDomain param.Field[string]   `json:"mdnsDomain"`
	Cors       param.Field[[]string] `json:"cors"`
}

func (r ServerConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigToolOutputParam is the request-side counterpart of [ConfigToolOutput].
type ConfigToolOutputParam struct {
	// Maximum number of lines to display in tool output
	MaxLines param.Field[int64] `json:"max_lines"`
	// Maximum number of bytes to display in tool output
	MaxBytes param.Field[int64] `json:"max_bytes"`
}

func (r ConfigToolOutputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigSkillsParam is the request-side counterpart of [ConfigSkills].
type ConfigSkillsParam struct {
	Paths param.Field[[]string] `json:"paths"`
	Urls  param.Field[[]string] `json:"urls"`
}

func (r ConfigSkillsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigCompactionParam is the request-side counterpart of [ConfigCompaction].
type ConfigCompactionParam struct {
	Auto                 param.Field[bool]  `json:"auto"`
	Prune                param.Field[bool]  `json:"prune"`
	Reserved             param.Field[int64] `json:"reserved"`
	TailTurns            param.Field[int64] `json:"tail_turns"`
	PreserveRecentTokens param.Field[int64] `json:"preserve_recent_tokens"`
}

func (r ConfigCompactionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EnterpriseConfigParam is the request-side counterpart of [EnterpriseConfig].
type EnterpriseConfigParam struct {
	URL param.Field[string] `json:"url"`
}

func (r EnterpriseConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigWatcherParam is the request-side counterpart of [ConfigWatcher].
type ConfigWatcherParam struct {
	Ignore param.Field[[]string] `json:"ignore"`
}

func (r ConfigWatcherParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigCommandParam is the request-side counterpart of [ConfigCommand].
type ConfigCommandParam struct {
	Template    param.Field[string] `json:"template,required"`
	Agent       param.Field[string] `json:"agent"`
	Description param.Field[string] `json:"description"`
	Model       param.Field[string] `json:"model"`
	Variant     param.Field[string] `json:"variant"`
	Subtask     param.Field[bool]   `json:"subtask"`
}

func (r ConfigCommandParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigExperimentalParam is the request-side counterpart of [ConfigExperimental].
type ConfigExperimentalParam struct {
	BatchTool           param.Field[bool]                              `json:"batch_tool"`
	ContinueLoopOnDeny  param.Field[bool]                              `json:"continue_loop_on_deny"`
	DisablePasteSummary param.Field[bool]                              `json:"disable_paste_summary"`
	McpTimeout          param.Field[int64]                             `json:"mcp_timeout"`
	OpenTelemetry       param.Field[bool]                              `json:"openTelemetry"`
	Policies            param.Field[[]ConfigV2ExperimentalPolicyParam] `json:"policies"`
	PrimaryTools        param.Field[[]string]                          `json:"primary_tools"`
}

func (r ConfigExperimentalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigV2ExperimentalPolicyParam is the request-side counterpart of [ConfigV2ExperimentalPolicy].
type ConfigV2ExperimentalPolicyParam struct {
	Action   param.Field[ConfigV2ExperimentalPolicyAction] `json:"action,required"`
	Effect   param.Field[PolicyEffect]                     `json:"effect,required"`
	Resource param.Field[string]                           `json:"resource,required"`
}

func (r ConfigV2ExperimentalPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ImageAttachmentConfigParam is the request-side counterpart of [ImageAttachmentConfig].
type ImageAttachmentConfigParam struct {
	// Automatically resize images before sending
	AutoResize param.Field[bool] `json:"auto_resize"`
	// Maximum image width in pixels
	MaxWidth param.Field[int64] `json:"max_width"`
	// Maximum image height in pixels
	MaxHeight param.Field[int64] `json:"max_height"`
	// Maximum base64 encoded image size in bytes
	MaxBase64Bytes param.Field[int64] `json:"max_base64_bytes"`
}

func (r ImageAttachmentConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AttachmentConfigParam is the request-side counterpart of [AttachmentConfig].
type AttachmentConfigParam struct {
	// Image attachment configuration
	Image param.Field[ImageAttachmentConfigParam] `json:"image"`
}

func (r AttachmentConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PermissionConfigObjectParam is the request-side counterpart of
// [PermissionConfigObject]. It is the object variant of the
// [PermissionConfigUnionParam] request union.
type PermissionConfigObjectParam struct {
	Read              param.Field[PermissionRuleConfigUnionParam] `json:"read"`
	Edit              param.Field[PermissionRuleConfigUnionParam] `json:"edit"`
	Glob              param.Field[PermissionRuleConfigUnionParam] `json:"glob"`
	Grep              param.Field[PermissionRuleConfigUnionParam] `json:"grep"`
	List              param.Field[PermissionRuleConfigUnionParam] `json:"list"`
	Bash              param.Field[PermissionRuleConfigUnionParam] `json:"bash"`
	Task              param.Field[PermissionRuleConfigUnionParam] `json:"task"`
	ExternalDirectory param.Field[PermissionRuleConfigUnionParam] `json:"external_directory"`
	Lsp               param.Field[PermissionRuleConfigUnionParam] `json:"lsp"`
	Skill             param.Field[PermissionRuleConfigUnionParam] `json:"skill"`
	Todowrite         param.Field[PermissionActionConfig]         `json:"todowrite"`
	Question          param.Field[PermissionActionConfig]         `json:"question"`
	Webfetch          param.Field[PermissionActionConfig]         `json:"webfetch"`
	Websearch         param.Field[PermissionActionConfig]         `json:"websearch"`
	DoomLoop          param.Field[PermissionActionConfig]         `json:"doom_loop"`
	// Additional per-tool rules beyond the properties listed above, per the
	// OpenAPI `PermissionConfig.additionalProperties` -> `PermissionRuleConfig`
	// mapping.
	ExtraFields map[string]PermissionRuleConfigUnionParam `json:"-,extras"`
}

func (r PermissionConfigObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PermissionConfigObjectParam) implementsPermissionConfigUnionParam() {}

// PermissionRuleConfigUnionParam is the request-side union for the OpenAPI
// `PermissionRuleConfig` anyOf.
//
// Satisfied by [PermissionActionConfig] (a short string "ask"|"allow"|"deny") or
// [PermissionObjectConfig] (a map of pattern to action).
type PermissionRuleConfigUnionParam interface {
	implementsPermissionRuleConfigUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[PermissionRuleConfigUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[PermissionActionConfig](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PermissionObjectConfig](),
		},
	)
}

// PermissionConfigUnionParam is the request-side union for the OpenAPI
// `PermissionConfig` anyOf.
//
// Satisfied by [PermissionActionConfig] (a short string "ask"|"allow"|"deny") or
// [PermissionConfigObjectParam] (per-tool permission rule overrides).
type PermissionConfigUnionParam interface {
	implementsPermissionConfigUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[PermissionConfigUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[PermissionActionConfig](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[PermissionConfigObjectParam](),
		},
	)
}

// McpLocalConfigParam is the request-side counterpart of [McpLocalConfig].
type McpLocalConfigParam struct {
	// Command and arguments to run the MCP server
	Command param.Field[[]string] `json:"command,required"`
	// Type of MCP server connection
	Type param.Field[McpLocalConfigType] `json:"type,required"`
	Cwd  param.Field[string]             `json:"cwd"`
	// Enable or disable the MCP server on startup
	Enabled param.Field[bool] `json:"enabled"`
	// Timeout in milliseconds
	Timeout param.Field[int64] `json:"timeout"`
	// Environment variables to set when running the MCP server
	Environment param.Field[map[string]string] `json:"environment"`
}

func (r McpLocalConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r McpLocalConfigParam) implementsConfigMcpUnionParam() {}

// McpRemoteConfigParam is the request-side counterpart of [McpRemoteConfig].
type McpRemoteConfigParam struct {
	// Type of MCP server connection
	Type param.Field[McpRemoteConfigType] `json:"type,required"`
	// URL of the remote MCP server
	URL param.Field[string] `json:"url,required"`
	// Enable or disable the MCP server on startup
	Enabled param.Field[bool] `json:"enabled"`
	// Timeout in milliseconds
	Timeout param.Field[int64] `json:"timeout"`
	// Headers to send with the request
	Headers param.Field[map[string]string] `json:"headers"`
	// OAuth authentication configuration for the MCP server. Set to false to
	// disable OAuth auto-detection.
	//
	// Per the OpenAPI schema, this field can be either [McpOAuthConfigParam] (a
	// complete OAuth config) or [McpOAuthConfigDisabledParam] (a scalar `false`).
	OAuth param.Field[McpOAuthConfigUnionParam] `json:"oauth"`
}

func (r McpRemoteConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r McpRemoteConfigParam) implementsConfigMcpUnionParam() {}

// McpOAuthConfigParam is the request-side counterpart of [McpOAuthConfig]
// (OpenAPI McpOAuthConfig schema).
type McpOAuthConfigParam struct {
	// OAuth client ID. If not provided, dynamic client registration (RFC 7591) will be attempted.
	ClientID param.Field[string] `json:"clientId"`
	// OAuth client secret (if required by the authorization server)
	ClientSecret param.Field[string] `json:"clientSecret"`
	// OAuth scopes to request during authorization
	Scope param.Field[string] `json:"scope"`
	// OAuth callback port for the local HTTP server
	CallbackPort param.Field[int64] `json:"callbackPort"`
	// OAuth redirect URI
	RedirectURI param.Field[string] `json:"redirectUri"`
}

func (r McpOAuthConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r McpOAuthConfigParam) implementsMcpOAuthConfigUnionParam() {}

// McpOAuthConfigDisabledParam represents the explicit `false` value for OAuth,
// disabling OAuth auto-detection (OpenAPI boolean enum: [false]).
type McpOAuthConfigDisabledParam struct {
}

func (r McpOAuthConfigDisabledParam) MarshalJSON() (data []byte, err error) {
	return []byte("false"), nil
}

func (r McpOAuthConfigDisabledParam) implementsMcpOAuthConfigUnionParam() {}

// McpOAuthConfigUnion is the OpenAPI `McpRemoteConfig.oauth` anyOf union.
//
// Satisfied by [McpOAuthConfig] (a complete OAuth config) or [shared.UnionBool]
// (a scalar false that disables OAuth auto-detection).
type McpOAuthConfigUnion interface {
	ImplementsMcpOAuthConfigUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[McpOAuthConfigUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[McpOAuthConfig](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[shared.UnionBool](),
		},
	)
}

// Satisfied by [McpOAuthConfigParam], [McpOAuthConfigDisabledParam].
type McpOAuthConfigUnionParam interface {
	implementsMcpOAuthConfigUnionParam()
}

// ConfigMcpDisabledParam is the request-side counterpart of [ConfigMcpDisabled].
type ConfigMcpDisabledParam struct {
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ConfigMcpDisabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigMcpDisabledParam) implementsConfigMcpUnionParam() {}

// ConfigMcpUnionParam is the request-side union for per-server MCP config.
//
// Satisfied by [McpLocalConfigParam], [McpRemoteConfigParam], or [ConfigMcpDisabledParam].
type ConfigMcpUnionParam interface {
	implementsConfigMcpUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigMcpUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[McpLocalConfigParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[McpRemoteConfigParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigMcpDisabledParam](),
		},
	)
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[McpOAuthConfigUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[McpOAuthConfigParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[McpOAuthConfigDisabledParam](),
		},
	)
}

// ConfigProviderOptionsParam is the request-side counterpart of [ConfigProviderOptions].
type ConfigProviderOptionsParam struct {
	APIKey        param.Field[string] `json:"apiKey"`
	BaseURL       param.Field[string] `json:"baseURL"`
	EnterpriseURL param.Field[string] `json:"enterpriseUrl"`
	SetCacheKey   param.Field[bool]   `json:"setCacheKey"`
	// Timeout in milliseconds for full requests to this provider. Set to false to
	// disable timeout. Accepts [shared.UnionInt] or [shared.UnionBool].
	Timeout param.Field[ConfigProviderOptionsTimeoutUnion] `json:"timeout"`
	// Timeout in milliseconds to wait for response headers. Set to false to disable.
	// Accepts [shared.UnionInt] or [shared.UnionBool].
	HeaderTimeout param.Field[ConfigProviderOptionsTimeoutUnion] `json:"headerTimeout"`
	ChunkTimeout  param.Field[int64]                             `json:"chunkTimeout"`
}

func (r ConfigProviderOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelParam is the request-side counterpart of [ConfigProviderModel].
type ConfigProviderModelParam struct {
	ID           param.Field[string]            `json:"id"`
	Attachment   param.Field[bool]              `json:"attachment"`
	Experimental param.Field[bool]              `json:"experimental"`
	Family       param.Field[string]            `json:"family"`
	Headers      param.Field[map[string]string] `json:"headers"`
	// Per OpenAPI `ProviderConfig.models.*.interleaved` is an anyOf of four
	// variants: boolean, the enum "reasoning"|"reasoning_content"|"reasoning_text",
	// any arbitrary string, or the object `{ "field": string }` (use
	// [ConfigProviderModelsInterleavedFieldParam] for the object variant).
	// Accepts [ConfigProviderModelsInterleavedEnabled],
	// [ConfigProviderModelsInterleavedString] or
	// [ConfigProviderModelsInterleavedFieldParam].
	Interleaved param.Field[ConfigProviderModelsInterleavedUnionParam] `json:"interleaved"`
	Name        param.Field[string]                                    `json:"name"`
	Options     param.Field[map[string]any]                            `json:"options"`
	Reasoning   param.Field[bool]                                      `json:"reasoning"`
	ReleaseDate param.Field[string]                                    `json:"release_date"`
	Temperature param.Field[bool]                                      `json:"temperature"`
	ToolCall    param.Field[bool]                                      `json:"tool_call"`
	Cost        param.Field[ConfigProviderModelsCostParam]             `json:"cost"`
	Limit       param.Field[ConfigProviderModelsLimitParam]            `json:"limit"`
	Modalities  param.Field[ConfigProviderModelsModalitiesParam]       `json:"modalities"`
	Provider    param.Field[ConfigProviderModelsProviderParam]         `json:"provider"`
	Status      param.Field[ConfigProviderModelsStatus]                `json:"status"`
	// Accepts object.
	Variants param.Field[any] `json:"variants"`
}

func (r ConfigProviderModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsCostParam is the request-side counterpart of
// [ConfigProviderModelsCost].
type ConfigProviderModelsCostParam struct {
	Input           param.Field[float64]                                      `json:"input,required"`
	Output          param.Field[float64]                                      `json:"output,required"`
	CacheRead       param.Field[float64]                                      `json:"cache_read"`
	CacheWrite      param.Field[float64]                                      `json:"cache_write"`
	ContextOver200k param.Field[ConfigProviderModelsCostContextOver200kParam] `json:"context_over_200k"`
}

func (r ConfigProviderModelsCostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsCostContextOver200kParam is the request-side counterpart
// of [ConfigProviderModelsCostContextOver200k].
type ConfigProviderModelsCostContextOver200kParam struct {
	Input      param.Field[float64] `json:"input,required"`
	Output     param.Field[float64] `json:"output,required"`
	CacheRead  param.Field[float64] `json:"cache_read"`
	CacheWrite param.Field[float64] `json:"cache_write"`
}

func (r ConfigProviderModelsCostContextOver200kParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsLimitParam is the request-side counterpart of
// [ConfigProviderModelsLimit].
type ConfigProviderModelsLimitParam struct {
	Context param.Field[int64] `json:"context,required"`
	Input   param.Field[int64] `json:"input"`
	Output  param.Field[int64] `json:"output,required"`
}

func (r ConfigProviderModelsLimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsModalitiesParam is the request-side counterpart of
// [ConfigProviderModelsModalities].
type ConfigProviderModelsModalitiesParam struct {
	Input  param.Field[[]ConfigProviderModelsModalitiesInput]  `json:"input"`
	Output param.Field[[]ConfigProviderModelsModalitiesOutput] `json:"output"`
}

func (r ConfigProviderModelsModalitiesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsProviderParam is the request-side counterpart of
// [ConfigProviderModelsProvider].
type ConfigProviderModelsProviderParam struct {
	NPM param.Field[string] `json:"npm"`
	API param.Field[string] `json:"api"`
}

func (r ConfigProviderModelsProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsInterleavedFieldParam is the request-side counterpart of the
// `{ "field": ... }` variant of ProviderConfig.models.*.interleaved.
type ConfigProviderModelsInterleavedFieldParam struct {
	Field param.Field[ProviderModelCapabilitiesInterleavedFieldField] `json:"field,required"`
}

func (r ConfigProviderModelsInterleavedFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigProviderModelsInterleavedFieldParam) implementsConfigProviderModelsInterleavedUnionParam() {
}

// ConfigProviderParam is the request-side counterpart of [ConfigProvider].
type ConfigProviderParam struct {
	ID        param.Field[string]                              `json:"id"`
	API       param.Field[string]                              `json:"api"`
	Env       param.Field[[]string]                            `json:"env"`
	Models    param.Field[map[string]ConfigProviderModelParam] `json:"models"`
	Name      param.Field[string]                              `json:"name"`
	NPM       param.Field[string]                              `json:"npm"`
	Whitelist param.Field[[]string]                            `json:"whitelist"`
	Blacklist param.Field[[]string]                            `json:"blacklist"`
	Options   param.Field[ConfigProviderOptionsParam]          `json:"options"`
}

func (r ConfigProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AgentConfigParam is the request-side counterpart of [AgentConfig] (the OpenAPI
// `AgentConfig` schema). It is shared by every agent slot of [ConfigAgentParam]
// and [ConfigModeParam], and by their user-defined entries.
type AgentConfigParam struct {
	Model       param.Field[string]          `json:"model"`
	Variant     param.Field[string]          `json:"variant"`
	Temperature param.Field[float64]         `json:"temperature"`
	TopP        param.Field[float64]         `json:"top_p"`
	Prompt      param.Field[string]          `json:"prompt"`
	Tools       param.Field[map[string]bool] `json:"tools"`
	Disable     param.Field[bool]            `json:"disable"`
	Description param.Field[string]          `json:"description"`
	Mode        param.Field[AgentConfigMode] `json:"mode"`
	Hidden      param.Field[bool]            `json:"hidden"`
	Options     param.Field[map[string]any]  `json:"options"`
	// Hex color code (e.g., #FF5733) or theme color (e.g., primary)
	Color    param.Field[string] `json:"color"`
	Steps    param.Field[int64]  `json:"steps"`
	MaxSteps param.Field[int64]  `json:"maxSteps"`
	// Permission configuration. Accepts [PermissionActionConfig] (a short string
	// "ask"|"allow"|"deny") or [PermissionConfigObjectParam].
	Permission param.Field[PermissionConfigUnionParam] `json:"permission"`
	// Additional agent properties not listed above. The OpenAPI `AgentConfig`
	// schema allows arbitrary extra properties.
	ExtraFields map[string]any `json:"-,extras"`
}

func (r AgentConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigAgentParam is the request-side counterpart of [ConfigAgent].
// ExtraFields allows passing arbitrary named agent configs beyond the named
// sub-agents, per the OpenAPI `Config.agent.additionalProperties` mapping.
type ConfigAgentParam struct {
	Build      param.Field[AgentConfigParam] `json:"build"`
	Compaction param.Field[AgentConfigParam] `json:"compaction"`
	Explore    param.Field[AgentConfigParam] `json:"explore"`
	General    param.Field[AgentConfigParam] `json:"general"`
	Plan       param.Field[AgentConfigParam] `json:"plan"`
	Summary    param.Field[AgentConfigParam] `json:"summary"`
	Title      param.Field[AgentConfigParam] `json:"title"`
	// User-defined agents keyed by agent name.
	ExtraFields map[string]AgentConfigParam `json:"-,extras"`
}

func (r ConfigAgentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigModeParam is the request-side counterpart of [ConfigMode].
//
// @deprecated Use ConfigAgentParam instead.
type ConfigModeParam struct {
	Build param.Field[AgentConfigParam] `json:"build"`
	Plan  param.Field[AgentConfigParam] `json:"plan"`
	// User-defined modes keyed by mode name.
	ExtraFields map[string]AgentConfigParam `json:"-,extras"`
}

func (r ConfigModeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
