// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/tidwall/gjson"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/shared"
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
	// Enable or configure LSP servers. Omit or set to false to disable, true to
	// enable built-ins, or an object to enable built-ins with overrides.
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
	// This field can have the runtime type of [string], [ConfigPermission].
	Permission any `json:"permission"`
	// This field can have the runtime type of [string] or [][2]any{string, object}.
	Plugin []any `json:"plugin"`
	// Custom provider configurations and model overrides
	Provider map[string]ConfigProvider `json:"provider"`
	// Reference configuration for external documentation. Keys are reference
	// names, values can be a plain URL/path string or a structured config (git
	// or local).
	// This field can have the runtime type of [string], [ConfigV2ReferenceGit],
	// [ConfigV2ReferenceLocal].
	Reference map[string]any `json:"reference"`
	// References from external sources. Keys are reference names, values can be a
	// plain URL/path string or a structured config (git or local).
	// This field can have the runtime type of [string], [ConfigV2ReferenceGit],
	// [ConfigV2ReferenceLocal].
	References map[string]any `json:"references"`
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
	// routes the raw data through [ConfigPermissionUnion] registered variants.
	permissionUnion ConfigPermissionUnion
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
	if permissionData != "" {
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
// Possible runtime types of the union are [string] (PermissionActionConfig:
// "ask"|"allow"|"deny") or [ConfigPermission].
func (r *Config) AsPermission() ConfigPermissionUnion {
	return r.permissionUnion
}

// ConfigPermissionUnion represents the OpenAPI PermissionConfig anyOf union.
//
// Satisfied by [ConfigPermission], [ConfigPermissionAction] (a short string
// permission: "ask"|"allow"|"deny").
type ConfigPermissionUnion interface {
	implementsConfigPermissionUnion()
}

// ConfigPermissionAction is a short string permission, e.g. "ask" / "allow" /
// "deny", corresponding to OpenAPI [PermissionActionConfig].
type ConfigPermissionAction string

const (
	ConfigPermissionActionAsk   ConfigPermissionAction = "ask"
	ConfigPermissionActionAllow ConfigPermissionAction = "allow"
	ConfigPermissionActionDeny  ConfigPermissionAction = "deny"
)

func (r ConfigPermissionAction) IsKnown() bool {
	switch r {
	case ConfigPermissionActionAsk, ConfigPermissionActionAllow, ConfigPermissionActionDeny:
		return true
	}
	return false
}

func (r ConfigPermissionAction) implementsConfigPermissionUnion() {}

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

// Agent configuration, see https://opencode.ai/docs/agent
type ConfigAgent struct {
	Build       AgentConfig            `json:"build"`
	Compaction  AgentConfig            `json:"compaction"`
	Explore     AgentConfig            `json:"explore"`
	General     AgentConfig            `json:"general"`
	Plan        AgentConfig            `json:"plan"`
	Summary     AgentConfig            `json:"summary"`
	Title       AgentConfig            `json:"title"`
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

type AgentConfig struct {
	// Description of when to use the agent
	Description string          `json:"description"`
	Disable     bool            `json:"disable"`
	Mode        AgentConfigMode `json:"mode"`
	Model       string          `json:"model"`
	// Permission configuration. Can be a short string ("ask"|"allow"|"deny") or
	// an object with per-action permission rule overrides.
	// This field can have the runtime type of [ConfigPermissionAction] or [ConfigPermission].
	Permission  any             `json:"permission"`
	Prompt      string          `json:"prompt"`
	Temperature float64         `json:"temperature"`
	Tools       map[string]bool `json:"tools"`
	TopP        float64         `json:"top_p"`
	Variant     string          `json:"variant"`
	Hidden      bool            `json:"hidden"`
	Options     map[string]any  `json:"options"`
	Color       string          `json:"color"`
	Steps       int64           `json:"steps"`
	MaxSteps    int64           `json:"maxSteps"`
	ExtraFields map[string]any  `json:"-,extras"`
	JSON        agentConfigJSON `json:"-"`
	// permissionUnion holds the typed permission payload after [UnmarshalJSON]
	// routes the raw data through [ConfigPermissionUnion] registered variants.
	permissionUnion ConfigPermissionUnion
}

// agentConfigJSON contains the JSON metadata for the struct
// [AgentConfig]
type agentConfigJSON struct {
	Description apijson.Field
	Disable     apijson.Field
	Mode        apijson.Field
	Model       apijson.Field
	Permission  apijson.Field
	Prompt      apijson.Field
	Temperature apijson.Field
	Tools       apijson.Field
	TopP        apijson.Field
	Variant     apijson.Field
	Hidden      apijson.Field
	Options     apijson.Field
	Color       apijson.Field
	Steps       apijson.Field
	MaxSteps    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentConfig) UnmarshalJSON(data []byte) (err error) {
	*r = AgentConfig{}
	if err = apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}
	permissionData := gjson.GetBytes(data, "permission").Raw
	if permissionData != "" {
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
// Possible runtime types of the union are [ConfigPermissionAction] (a short
// string: "ask"|"allow"|"deny") or [ConfigPermission].
func (r *AgentConfig) AsPermission() ConfigPermissionUnion {
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
	// This field can have the runtime type of [McpOAuthConfig, nil]. OAuth authentication configuration for the MCP server (for "remote" type).
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
	Build       AgentConfig            `json:"build"`
	Plan        AgentConfig            `json:"plan"`
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

type ConfigPermission struct {
	// This field can have the runtime type of [ConfigPermissionBashString], [ConfigPermissionBashMap].
	Bash any `json:"bash"`
	// This field can have the runtime type of [string], [map[string]any].
	Edit     any                      `json:"edit"`
	Webfetch ConfigPermissionWebfetch `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]any].
	Read any `json:"read"`
	// This field can have the runtime type of [string] or [map[string]any].
	Glob any `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]any].
	Grep any `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]any].
	List any `json:"list"`
	// This field can have the runtime type of [string] or [map[string]any].
	Task any `json:"task"`
	// This field can have the runtime type of [string] or [map[string]any].
	ExternalDirectory any                       `json:"external_directory"`
	Todowrite         ConfigPermissionTodowrite `json:"todowrite"`
	Question          ConfigPermissionQuestion  `json:"question"`
	Websearch         ConfigPermissionWebsearch `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]any].
	Lsp      any                      `json:"lsp"`
	DoomLoop ConfigPermissionDoomLoop `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]any].
	Skill any                  `json:"skill"`
	JSON  configPermissionJSON `json:"-"`
}

// configPermissionJSON contains the JSON metadata for the struct
// [ConfigPermission]
type configPermissionJSON struct {
	Bash              apijson.Field
	Edit              apijson.Field
	Webfetch          apijson.Field
	Read              apijson.Field
	Glob              apijson.Field
	Grep              apijson.Field
	List              apijson.Field
	Task              apijson.Field
	ExternalDirectory apijson.Field
	Todowrite         apijson.Field
	Question          apijson.Field
	Websearch         apijson.Field
	Lsp               apijson.Field
	DoomLoop          apijson.Field
	Skill             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ConfigPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configPermissionJSON) RawJSON() string {
	return r.raw
}

func (r ConfigPermission) implementsConfigPermissionUnion() {}

// Union satisfied by [ConfigPermissionBashString] or [ConfigPermissionBashMap].
type ConfigPermissionBashUnion interface {
	implementsConfigPermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigPermissionUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[ConfigPermissionAction](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigPermission](),
		},
	)
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigPermissionBashUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[ConfigPermissionBashString](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigPermissionBashMap](),
		},
	)
}

type ConfigPermissionBashString string

const (
	ConfigPermissionBashStringAsk   ConfigPermissionBashString = "ask"
	ConfigPermissionBashStringAllow ConfigPermissionBashString = "allow"
	ConfigPermissionBashStringDeny  ConfigPermissionBashString = "deny"
)

func (r ConfigPermissionBashString) IsKnown() bool {
	switch r {
	case ConfigPermissionBashStringAsk, ConfigPermissionBashStringAllow, ConfigPermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigPermissionBashString) implementsConfigPermissionBashUnion() {}

type ConfigPermissionBashMap map[string]ConfigPermissionBashMapItem

func (r ConfigPermissionBashMap) implementsConfigPermissionBashUnion() {}

type ConfigPermissionBashMapItem string

const (
	ConfigPermissionBashMapAsk   ConfigPermissionBashMapItem = "ask"
	ConfigPermissionBashMapAllow ConfigPermissionBashMapItem = "allow"
	ConfigPermissionBashMapDeny  ConfigPermissionBashMapItem = "deny"
)

func (r ConfigPermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigPermissionBashMapAsk, ConfigPermissionBashMapAllow, ConfigPermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigPermissionWebfetch string

const (
	ConfigPermissionWebfetchAsk   ConfigPermissionWebfetch = "ask"
	ConfigPermissionWebfetchAllow ConfigPermissionWebfetch = "allow"
	ConfigPermissionWebfetchDeny  ConfigPermissionWebfetch = "deny"
)

func (r ConfigPermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigPermissionWebfetchAsk, ConfigPermissionWebfetchAllow, ConfigPermissionWebfetchDeny:
		return true
	}
	return false
}

type ConfigPermissionTodowrite string

const (
	ConfigPermissionTodowriteAsk   ConfigPermissionTodowrite = "ask"
	ConfigPermissionTodowriteAllow ConfigPermissionTodowrite = "allow"
	ConfigPermissionTodowriteDeny  ConfigPermissionTodowrite = "deny"
)

func (r ConfigPermissionTodowrite) IsKnown() bool {
	switch r {
	case ConfigPermissionTodowriteAsk, ConfigPermissionTodowriteAllow, ConfigPermissionTodowriteDeny:
		return true
	}
	return false
}

type ConfigPermissionQuestion string

const (
	ConfigPermissionQuestionAsk   ConfigPermissionQuestion = "ask"
	ConfigPermissionQuestionAllow ConfigPermissionQuestion = "allow"
	ConfigPermissionQuestionDeny  ConfigPermissionQuestion = "deny"
)

func (r ConfigPermissionQuestion) IsKnown() bool {
	switch r {
	case ConfigPermissionQuestionAsk, ConfigPermissionQuestionAllow, ConfigPermissionQuestionDeny:
		return true
	}
	return false
}

type ConfigPermissionWebsearch string

const (
	ConfigPermissionWebsearchAsk   ConfigPermissionWebsearch = "ask"
	ConfigPermissionWebsearchAllow ConfigPermissionWebsearch = "allow"
	ConfigPermissionWebsearchDeny  ConfigPermissionWebsearch = "deny"
)

func (r ConfigPermissionWebsearch) IsKnown() bool {
	switch r {
	case ConfigPermissionWebsearchAsk, ConfigPermissionWebsearchAllow, ConfigPermissionWebsearchDeny:
		return true
	}
	return false
}

type ConfigPermissionDoomLoop string

const (
	ConfigPermissionDoomLoopAsk   ConfigPermissionDoomLoop = "ask"
	ConfigPermissionDoomLoopAllow ConfigPermissionDoomLoop = "allow"
	ConfigPermissionDoomLoopDeny  ConfigPermissionDoomLoop = "deny"
)

func (r ConfigPermissionDoomLoop) IsKnown() bool {
	switch r {
	case ConfigPermissionDoomLoopAsk, ConfigPermissionDoomLoopAllow, ConfigPermissionDoomLoopDeny:
		return true
	}
	return false
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
	// This field can have the runtime type of [bool] or object.
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
	Context float64                       `json:"context,required"`
	Input   float64                       `json:"input"`
	Output  float64                       `json:"output,required"`
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
	// This field can have the runtime type of [shared.UnionInt], [shared.UnionBool].
	Timeout any `json:"timeout"`
	// Timeout in milliseconds to wait for response headers. Provider integrations
	// may set defaults. Set to false to disable timeout.
	// This field can have the runtime type of [shared.UnionInt], [shared.UnionBool].
	HeaderTimeout any                       `json:"headerTimeout"`
	ChunkTimeout  int64                     `json:"chunkTimeout"`
	ExtraFields   map[string]any            `json:"-,extras"`
	JSON          configProviderOptionsJSON `json:"-"`
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
			TypeFilter: gjson.True,
			Type:       reflect.TypeFor[shared.UnionBool](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[shared.UnionBool](),
		},
	)
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
	// OAuth authentication configuration for this MCP server.
	// This field can have the runtime type of [McpOAuthConfig] or bool (false).
	OAuth any                 `json:"oauth"`
	JSON  mcpRemoteConfigJSON `json:"-"`
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
	// Accepts [bool] or [string] ("notify").
	Autoupdate        param.Field[any]                           `json:"autoupdate"`
	Command           param.Field[map[string]ConfigCommandParam] `json:"command"`
	Compaction        param.Field[ConfigCompactionParam]         `json:"compaction"`
	DisabledProviders param.Field[[]string]                      `json:"disabled_providers"`
	EnabledProviders  param.Field[[]string]                      `json:"enabled_providers"`
	Enterprise        param.Field[EnterpriseConfigParam]         `json:"enterprise"`
	Experimental      param.Field[ConfigExperimentalParam]       `json:"experimental"`
	// Enable or configure formatters. Pass false to disable, true to enable
	// built-ins, or a map of formatter-name to config to enable with overrides.
	// Accepts [bool] or [map[string]ConfigFormatter].
	Formatter    param.Field[any]            `json:"formatter"`
	Instructions param.Field[[]string]       `json:"instructions"`
	Layout       param.Field[ConfigLayout]   `json:"layout"`
	LogLevel     param.Field[ConfigLogLevel] `json:"logLevel"`
	// Enable or configure LSP servers. Pass false to disable, true to enable
	// built-ins, or a map of LSP-name to config to enable with overrides.
	// Accepts [bool] or [map[string]ConfigLsp].
	Lsp   param.Field[any]                       `json:"lsp"`
	Mcp   param.Field[map[string]ConfigMcpParam] `json:"mcp"`
	Mode  param.Field[ConfigModeParam]           `json:"mode"`
	Model param.Field[string]                    `json:"model"`
	// Permission configuration. A short string ("ask"|"allow"|"deny") or an
	// object with per-action permission rule overrides. Accepts [ConfigPermissionAction]
	// (a string constant) or [ConfigPermissionParam].
	Permission param.Field[ConfigPermissionUnionParam] `json:"permission"`
	// Plugins to load. Each item is either a plugin name (string) or a 2-tuple
	// of [pluginName, configObject] (where configObject is a map[string]any).
	Plugin   param.Field[[]any]                          `json:"plugin"`
	Provider param.Field[map[string]ConfigProviderParam] `json:"provider"`
	// Map of reference name → value. Each value can be a plain [string] (URL/path),
	// a [ConfigV2ReferenceGitParam], or a [ConfigV2ReferenceLocalParam].
	Reference param.Field[map[string]ConfigV2ReferenceUnionParam] `json:"reference"`
	// Map of reference name → value. Each value can be a plain [string] (URL/path),
	// a [ConfigV2ReferenceGitParam], or a [ConfigV2ReferenceLocalParam].
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

// ConfigV2ReferenceString is a plain string (URL or path) variant of
// [ConfigV2ReferenceUnion].
type ConfigV2ReferenceString string

func (r ConfigV2ReferenceString) implementsConfigV2ReferenceUnion() {}

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
// Request param types for ConfigUpdateParams (Request/Response separation).
// Every field is wrapped in param.Field[T]; each struct implements MarshalJSON.
// ---------------------------------------------------------------------------

// ConfigAgentParam is the request-side container for the `agent` config map.
type ConfigAgentParam struct {
	Build       param.Field[AgentConfigParam] `json:"build"`
	Compaction  param.Field[AgentConfigParam] `json:"compaction"`
	Explore     param.Field[AgentConfigParam] `json:"explore"`
	General     param.Field[AgentConfigParam] `json:"general"`
	Plan        param.Field[AgentConfigParam] `json:"plan"`
	Summary     param.Field[AgentConfigParam] `json:"summary"`
	Title       param.Field[AgentConfigParam] `json:"title"`
	ExtraFields map[string]AgentConfigParam   `json:"-,extras"`
}

func (r ConfigAgentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigModeParam is the request-side container for the deprecated `mode` config map.
//
// @deprecated Use `agent` field instead.
type ConfigModeParam struct {
	Build       param.Field[AgentConfigParam] `json:"build"`
	Plan        param.Field[AgentConfigParam] `json:"plan"`
	ExtraFields map[string]AgentConfigParam   `json:"-,extras"`
}

func (r ConfigModeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AgentConfigParam is the request-side representation of AgentConfig.
type AgentConfigParam struct {
	// Description of when to use the agent
	Description param.Field[string]          `json:"description"`
	Disable     param.Field[bool]            `json:"disable"`
	Mode        param.Field[AgentConfigMode] `json:"mode"`
	Model       param.Field[string]          `json:"model"`
	// Permission accepts a short string ("ask"|"allow"|"deny") via
	// [ConfigPermissionAction], or a detailed object via [ConfigPermissionParam].
	Permission  param.Field[ConfigPermissionUnionParam] `json:"permission"`
	Prompt      param.Field[string]                     `json:"prompt"`
	Temperature param.Field[float64]                    `json:"temperature"`
	Tools       param.Field[map[string]bool]            `json:"tools"`
	TopP        param.Field[float64]                    `json:"top_p"`
	Variant     param.Field[string]                     `json:"variant"`
	Hidden      param.Field[bool]                       `json:"hidden"`
	Options     param.Field[map[string]any]             `json:"options"`
	Color       param.Field[string]                     `json:"color"`
	Steps       param.Field[int64]                      `json:"steps"`
	MaxSteps    param.Field[int64]                      `json:"maxSteps"`
}

func (r AgentConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AttachmentConfigParam is the request-side representation of AttachmentConfig.
type AttachmentConfigParam struct {
	Image param.Field[ImageAttachmentConfigParam] `json:"image"`
}

func (r AttachmentConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ImageAttachmentConfigParam is the request-side representation of ImageAttachmentConfig.
type ImageAttachmentConfigParam struct {
	AutoResize     param.Field[bool]  `json:"auto_resize"`
	MaxWidth       param.Field[int64] `json:"max_width"`
	MaxHeight      param.Field[int64] `json:"max_height"`
	MaxBase64Bytes param.Field[int64] `json:"max_base64_bytes"`
}

func (r ImageAttachmentConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigCommandParam is the request-side representation of ConfigCommand.
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

// ConfigCompactionParam is the request-side representation of ConfigCompaction.
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

// EnterpriseConfigParam is the request-side representation of EnterpriseConfig.
type EnterpriseConfigParam struct {
	URL param.Field[string] `json:"url"`
}

func (r EnterpriseConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigExperimentalParam is the request-side representation of ConfigExperimental.
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

// ConfigV2ExperimentalPolicyParam is the request-side representation of ConfigV2ExperimentalPolicy.
type ConfigV2ExperimentalPolicyParam struct {
	Action   param.Field[ConfigV2ExperimentalPolicyAction] `json:"action,required"`
	Effect   param.Field[PolicyEffect]                     `json:"effect,required"`
	Resource param.Field[string]                           `json:"resource,required"`
}

func (r ConfigV2ExperimentalPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigMcpParam is the request-side representation of ConfigMcp.
//
// ConfigMcp is a union of three variants: McpLocalConfig (type="local"),
// McpRemoteConfig (type="remote"), and the disabled form ({enabled: false}).
// The disabled form has no `type` field, so Type is not required here.
type ConfigMcpParam struct {
	Type        param.Field[ConfigMcpType] `json:"type"`
	Command     param.Field[any]           `json:"command"`
	Cwd         param.Field[any]           `json:"cwd"`
	Enabled     param.Field[bool]          `json:"enabled"`
	Environment param.Field[any]           `json:"environment"`
	Headers     param.Field[any]           `json:"headers"`
	OAuth       param.Field[any]           `json:"oauth"`
	Timeout     param.Field[any]           `json:"timeout"`
	URL         param.Field[string]        `json:"url"`
}

func (r ConfigMcpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ServerConfigParam is the request-side representation of ServerConfig.
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

// ConfigSkillsParam is the request-side representation of ConfigSkills.
type ConfigSkillsParam struct {
	Paths param.Field[[]string] `json:"paths"`
	Urls  param.Field[[]string] `json:"urls"`
}

func (r ConfigSkillsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigToolOutputParam is the request-side representation of ConfigToolOutput.
type ConfigToolOutputParam struct {
	MaxLines param.Field[int64] `json:"max_lines"`
	MaxBytes param.Field[int64] `json:"max_bytes"`
}

func (r ConfigToolOutputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigWatcherParam is the request-side representation of ConfigWatcher.
type ConfigWatcherParam struct {
	Ignore param.Field[[]string] `json:"ignore"`
}

func (r ConfigWatcherParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderParam is the request-side representation of ConfigProvider.
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

// ConfigProviderModelParam is the request-side representation of ConfigProviderModel.
type ConfigProviderModelParam struct {
	ID           param.Field[string]                              `json:"id"`
	Attachment   param.Field[bool]                                `json:"attachment"`
	Cost         param.Field[ConfigProviderModelsCostParam]       `json:"cost"`
	Experimental param.Field[bool]                                `json:"experimental"`
	Family       param.Field[string]                              `json:"family"`
	Headers      param.Field[map[string]string]                   `json:"headers"`
	Interleaved  param.Field[any]                                 `json:"interleaved"`
	Limit        param.Field[ConfigProviderModelsLimitParam]      `json:"limit"`
	Modalities   param.Field[ConfigProviderModelsModalitiesParam] `json:"modalities"`
	Name         param.Field[string]                              `json:"name"`
	Options      param.Field[map[string]any]                      `json:"options"`
	Provider     param.Field[ConfigProviderModelsProviderParam]   `json:"provider"`
	Reasoning    param.Field[bool]                                `json:"reasoning"`
	ReleaseDate  param.Field[string]                              `json:"release_date"`
	Status       param.Field[ConfigProviderModelsStatus]          `json:"status"`
	Temperature  param.Field[bool]                                `json:"temperature"`
	ToolCall     param.Field[bool]                                `json:"tool_call"`
	Variants     param.Field[any]                                 `json:"variants"`
}

func (r ConfigProviderModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsCostParam is the request-side representation of ConfigProviderModelsCost.
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

// ConfigProviderModelsCostContextOver200kParam is the request-side representation
// of ConfigProviderModelsCostContextOver200k.
type ConfigProviderModelsCostContextOver200kParam struct {
	Input      param.Field[float64] `json:"input,required"`
	Output     param.Field[float64] `json:"output,required"`
	CacheRead  param.Field[float64] `json:"cache_read"`
	CacheWrite param.Field[float64] `json:"cache_write"`
}

func (r ConfigProviderModelsCostContextOver200kParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsLimitParam is the request-side representation of ConfigProviderModelsLimit.
type ConfigProviderModelsLimitParam struct {
	Context param.Field[float64] `json:"context,required"`
	Input   param.Field[float64] `json:"input"`
	Output  param.Field[float64] `json:"output,required"`
}

func (r ConfigProviderModelsLimitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsModalitiesParam is the request-side representation of ConfigProviderModelsModalities.
type ConfigProviderModelsModalitiesParam struct {
	Input  param.Field[[]ConfigProviderModelsModalitiesInput]  `json:"input"`
	Output param.Field[[]ConfigProviderModelsModalitiesOutput] `json:"output"`
}

func (r ConfigProviderModelsModalitiesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderModelsProviderParam is the request-side representation of ConfigProviderModelsProvider.
type ConfigProviderModelsProviderParam struct {
	NPM param.Field[string] `json:"npm"`
	API param.Field[string] `json:"api"`
}

func (r ConfigProviderModelsProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigProviderOptionsParam is the request-side representation of ConfigProviderOptions.
type ConfigProviderOptionsParam struct {
	APIKey        param.Field[string] `json:"apiKey"`
	BaseURL       param.Field[string] `json:"baseURL"`
	EnterpriseURL param.Field[string] `json:"enterpriseUrl"`
	SetCacheKey   param.Field[bool]   `json:"setCacheKey"`
	Timeout       param.Field[any]    `json:"timeout"`
	HeaderTimeout param.Field[any]    `json:"headerTimeout"`
	ChunkTimeout  param.Field[int64]  `json:"chunkTimeout"`
}

func (r ConfigProviderOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigV2ReferenceGitParam is the request-side representation of ConfigV2ReferenceGit.
type ConfigV2ReferenceGitParam struct {
	Repository  param.Field[string] `json:"repository,required"`
	Branch      param.Field[string] `json:"branch"`
	Description param.Field[string] `json:"description"`
	Hidden      param.Field[bool]   `json:"hidden"`
}

func (r ConfigV2ReferenceGitParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigV2ReferenceGitParam) ImplementsConfigV2ReferenceUnionParam() {}

// ConfigV2ReferenceLocalParam is the request-side representation of ConfigV2ReferenceLocal.
type ConfigV2ReferenceLocalParam struct {
	Path        param.Field[string] `json:"path,required"`
	Description param.Field[string] `json:"description"`
	Hidden      param.Field[bool]   `json:"hidden"`
}

func (r ConfigV2ReferenceLocalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigV2ReferenceLocalParam) ImplementsConfigV2ReferenceUnionParam() {}

// ConfigV2ReferenceUnionParam is the request-side union for a reference entry.
// Satisfied by [shared.UnionString], [ConfigV2ReferenceGitParam],
// [ConfigV2ReferenceLocalParam].
type ConfigV2ReferenceUnionParam interface {
	ImplementsConfigV2ReferenceUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigV2ReferenceUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[shared.UnionString](),
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

// ConfigPermissionBashUnionParam is the request-side union for the `bash`
// permission rule in [ConfigPermissionParam]. Satisfied by
// [ConfigPermissionBashString] or [ConfigPermissionBashMapParam].
type ConfigPermissionBashUnionParam interface {
	implementsConfigPermissionBashUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigPermissionBashUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[ConfigPermissionBashString](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigPermissionBashMapParam](),
		},
	)
}

func (r ConfigPermissionBashString) implementsConfigPermissionBashUnionParam() {}

// ConfigPermissionBashMapParam is the request-side map variant of the bash
// permission union in [ConfigPermissionParam].
type ConfigPermissionBashMapParam map[string]ConfigPermissionBashMapItem

func (r ConfigPermissionBashMapParam) implementsConfigPermissionBashUnionParam() {}

// ConfigPermissionParam is the request-side representation of ConfigPermission.
type ConfigPermissionParam struct {
	Bash              param.Field[ConfigPermissionBashUnionParam] `json:"bash"`
	Edit              param.Field[any]                            `json:"edit"`
	Webfetch          param.Field[ConfigPermissionWebfetch]       `json:"webfetch"`
	Read              param.Field[any]                            `json:"read"`
	Glob              param.Field[any]                            `json:"glob"`
	Grep              param.Field[any]                            `json:"grep"`
	List              param.Field[any]                            `json:"list"`
	Task              param.Field[any]                            `json:"task"`
	ExternalDirectory param.Field[any]                            `json:"external_directory"`
	Todowrite         param.Field[ConfigPermissionTodowrite]      `json:"todowrite"`
	Question          param.Field[ConfigPermissionQuestion]       `json:"question"`
	Websearch         param.Field[ConfigPermissionWebsearch]      `json:"websearch"`
	Lsp               param.Field[any]                            `json:"lsp"`
	DoomLoop          param.Field[ConfigPermissionDoomLoop]       `json:"doom_loop"`
	Skill             param.Field[any]                            `json:"skill"`
}

func (r ConfigPermissionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigPermissionParam) implementsConfigPermissionUnionParam() {}

// ConfigPermissionUnionParam is the request-side union for the top-level
// `permission` config. Satisfied by [ConfigPermissionParam],
// [ConfigPermissionAction] (a short string permission: "ask"|"allow"|"deny").
type ConfigPermissionUnionParam interface {
	implementsConfigPermissionUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigPermissionUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeFor[ConfigPermissionAction](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigPermissionParam](),
		},
	)
}

func (r ConfigPermissionAction) implementsConfigPermissionUnionParam() {}
