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
	Autoupdate interface{} `json:"autoupdate"`
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
	Formatter interface{} `json:"formatter"`
	// Additional instruction files or patterns to include
	Instructions []string `json:"instructions"`
	// @deprecated Always uses stretch layout.
	Layout ConfigLayout `json:"layout"`
	// Log level for the application
	LogLevel ConfigLogLevel `json:"logLevel"`
	// This field can have the runtime type of [bool], [map[string]ConfigLsp].
	Lsp interface{} `json:"lsp"`
	// MCP (Model Context Protocol) server configurations
	Mcp map[string]ConfigMcp `json:"mcp"`
	// @deprecated Use `agent` field instead.
	Mode ConfigMode `json:"mode"`
	// Model to use in the format of provider/model, eg anthropic/claude-2
	Model string `json:"model"`
	// Permission configuration. A short string ("ask"|"allow"|"deny") or an
	// object with per-action permission rule overrides.
	// This field can have the runtime type of [ConfigPermissionAction],
	// [ConfigPermission].
	Permission interface{} `json:"permission"`
	// This field can have the runtime type of [string] or [][2]interface{}{string, object}.
	Plugin []interface{} `json:"plugin"`
	// Custom provider configurations and model overrides
	Provider map[string]ConfigProvider `json:"provider"`
	// Reference configuration for external documentation. Keys are reference
	// names, values can be a plain URL/path string or a structured config (git
	// or local).
	// This field can have the runtime type of [string], [ConfigV2ReferenceGit],
	// [ConfigV2ReferenceLocal].
	Reference map[string]interface{} `json:"reference"`
	// References from external sources. Keys are reference names, values can be a
	// plain URL/path string or a structured config (git or local).
	// This field can have the runtime type of [string], [ConfigV2ReferenceGit],
	// [ConfigV2ReferenceLocal].
	References map[string]interface{} `json:"references"`
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

func (r *ConfigPermissionAction) UnmarshalJSON(data []byte) (err error) {
	var raw string
	if err = apijson.UnmarshalRoot(data, &raw); err != nil {
		return err
	}
	*r = ConfigPermissionAction(raw)
	return nil
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
	Build       ConfigAgentBuild       `json:"build"`
	Compaction  ConfigAgentCompaction  `json:"compaction"`
	Explore     ConfigAgentExplore     `json:"explore"`
	General     ConfigAgentGeneral     `json:"general"`
	Plan        ConfigAgentPlan        `json:"plan"`
	Summary     ConfigAgentSummary     `json:"summary"`
	Title       ConfigAgentTitle       `json:"title"`
	ExtraFields map[string]ConfigAgent `json:"-,extras"`
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

type ConfigAgentBuild struct {
	// Description of when to use the agent
	Description string                     `json:"description"`
	Disable     bool                       `json:"disable"`
	Mode        ConfigAgentBuildMode       `json:"mode"`
	Model       string                     `json:"model"`
	Permission  ConfigAgentBuildPermission `json:"permission"`
	Prompt      string                     `json:"prompt"`
	Temperature float64                    `json:"temperature"`
	Tools       map[string]bool            `json:"tools"`
	TopP        float64                    `json:"top_p"`
	Variant     string                     `json:"variant"`
	Hidden      bool                       `json:"hidden"`
	Options     map[string]interface{}     `json:"options"`
	Color       string                     `json:"color"`
	Steps       int64                      `json:"steps"`
	MaxSteps    int64                      `json:"maxSteps"`
	ExtraFields map[string]interface{}     `json:"-,extras"`
	JSON        configAgentBuildJSON       `json:"-"`
}

// configAgentBuildJSON contains the JSON metadata for the struct
// [ConfigAgentBuild]
type configAgentBuildJSON struct {
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

func (r *ConfigAgentBuild) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentBuildJSON) RawJSON() string {
	return r.raw
}

type ConfigAgentBuildMode string

const (
	ConfigAgentBuildModeSubagent ConfigAgentBuildMode = "subagent"
	ConfigAgentBuildModePrimary  ConfigAgentBuildMode = "primary"
	ConfigAgentBuildModeAll      ConfigAgentBuildMode = "all"
)

func (r ConfigAgentBuildMode) IsKnown() bool {
	switch r {
	case ConfigAgentBuildModeSubagent, ConfigAgentBuildModePrimary, ConfigAgentBuildModeAll:
		return true
	}
	return false
}

type ConfigAgentBuildPermission struct {
	// This field can have the runtime type of [ConfigAgentBuildPermissionBashString], [ConfigAgentBuildPermissionBashMap].
	Bash interface{} `json:"bash"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Edit interface{} `json:"edit"`
	// This field can have the runtime type of [string].
	Webfetch interface{} `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}                    `json:"skill"`
	JSON  configAgentBuildPermissionJSON `json:"-"`
}

// configAgentBuildPermissionJSON contains the JSON metadata for the struct
// [ConfigAgentBuildPermission]
type configAgentBuildPermissionJSON struct {
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

func (r *ConfigAgentBuildPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentBuildPermissionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ConfigAgentBuildPermissionBashString] or
// [ConfigAgentBuildPermissionBashMap].
type ConfigAgentBuildPermissionBashUnion interface {
	implementsConfigAgentBuildPermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigAgentBuildPermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigAgentBuildPermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigAgentBuildPermissionBashMap{}),
		},
	)
}

type ConfigAgentBuildPermissionBashString string

const (
	ConfigAgentBuildPermissionBashStringAsk   ConfigAgentBuildPermissionBashString = "ask"
	ConfigAgentBuildPermissionBashStringAllow ConfigAgentBuildPermissionBashString = "allow"
	ConfigAgentBuildPermissionBashStringDeny  ConfigAgentBuildPermissionBashString = "deny"
)

func (r ConfigAgentBuildPermissionBashString) IsKnown() bool {
	switch r {
	case ConfigAgentBuildPermissionBashStringAsk, ConfigAgentBuildPermissionBashStringAllow, ConfigAgentBuildPermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigAgentBuildPermissionBashString) implementsConfigAgentBuildPermissionBashUnion() {}

type ConfigAgentBuildPermissionBashMap map[string]ConfigAgentBuildPermissionBashMapItem

func (r ConfigAgentBuildPermissionBashMap) implementsConfigAgentBuildPermissionBashUnion() {}

type ConfigAgentBuildPermissionBashMapItem string

const (
	ConfigAgentBuildPermissionBashMapAsk   ConfigAgentBuildPermissionBashMapItem = "ask"
	ConfigAgentBuildPermissionBashMapAllow ConfigAgentBuildPermissionBashMapItem = "allow"
	ConfigAgentBuildPermissionBashMapDeny  ConfigAgentBuildPermissionBashMapItem = "deny"
)

func (r ConfigAgentBuildPermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigAgentBuildPermissionBashMapAsk, ConfigAgentBuildPermissionBashMapAllow, ConfigAgentBuildPermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigAgentBuildPermissionEdit string

const (
	ConfigAgentBuildPermissionEditAsk   ConfigAgentBuildPermissionEdit = "ask"
	ConfigAgentBuildPermissionEditAllow ConfigAgentBuildPermissionEdit = "allow"
	ConfigAgentBuildPermissionEditDeny  ConfigAgentBuildPermissionEdit = "deny"
)

func (r ConfigAgentBuildPermissionEdit) IsKnown() bool {
	switch r {
	case ConfigAgentBuildPermissionEditAsk, ConfigAgentBuildPermissionEditAllow, ConfigAgentBuildPermissionEditDeny:
		return true
	}
	return false
}

type ConfigAgentBuildPermissionWebfetch string

const (
	ConfigAgentBuildPermissionWebfetchAsk   ConfigAgentBuildPermissionWebfetch = "ask"
	ConfigAgentBuildPermissionWebfetchAllow ConfigAgentBuildPermissionWebfetch = "allow"
	ConfigAgentBuildPermissionWebfetchDeny  ConfigAgentBuildPermissionWebfetch = "deny"
)

func (r ConfigAgentBuildPermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigAgentBuildPermissionWebfetchAsk, ConfigAgentBuildPermissionWebfetchAllow, ConfigAgentBuildPermissionWebfetchDeny:
		return true
	}
	return false
}

type ConfigAgentGeneral struct {
	// Description of when to use the agent
	Description string                       `json:"description"`
	Disable     bool                         `json:"disable"`
	Mode        ConfigAgentGeneralMode       `json:"mode"`
	Model       string                       `json:"model"`
	Permission  ConfigAgentGeneralPermission `json:"permission"`
	Prompt      string                       `json:"prompt"`
	Temperature float64                      `json:"temperature"`
	Tools       map[string]bool              `json:"tools"`
	TopP        float64                      `json:"top_p"`
	Variant     string                       `json:"variant"`
	Hidden      bool                         `json:"hidden"`
	Options     map[string]interface{}       `json:"options"`
	Color       string                       `json:"color"`
	Steps       int64                        `json:"steps"`
	MaxSteps    int64                        `json:"maxSteps"`
	ExtraFields map[string]interface{}       `json:"-,extras"`
	JSON        configAgentGeneralJSON       `json:"-"`
}

// configAgentGeneralJSON contains the JSON metadata for the struct
// [ConfigAgentGeneral]
type configAgentGeneralJSON struct {
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

func (r *ConfigAgentGeneral) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentGeneralJSON) RawJSON() string {
	return r.raw
}

type ConfigAgentGeneralMode string

const (
	ConfigAgentGeneralModeSubagent ConfigAgentGeneralMode = "subagent"
	ConfigAgentGeneralModePrimary  ConfigAgentGeneralMode = "primary"
	ConfigAgentGeneralModeAll      ConfigAgentGeneralMode = "all"
)

func (r ConfigAgentGeneralMode) IsKnown() bool {
	switch r {
	case ConfigAgentGeneralModeSubagent, ConfigAgentGeneralModePrimary, ConfigAgentGeneralModeAll:
		return true
	}
	return false
}

type ConfigAgentGeneralPermission struct {
	// This field can have the runtime type of [ConfigAgentGeneralPermissionBashString], [ConfigAgentGeneralPermissionBashMap].
	Bash interface{} `json:"bash"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Edit interface{} `json:"edit"`
	// This field can have the runtime type of [string].
	Webfetch interface{} `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}                      `json:"skill"`
	JSON  configAgentGeneralPermissionJSON `json:"-"`
}

// configAgentGeneralPermissionJSON contains the JSON metadata for the struct
// [ConfigAgentGeneralPermission]
type configAgentGeneralPermissionJSON struct {
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

func (r *ConfigAgentGeneralPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentGeneralPermissionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ConfigAgentGeneralPermissionBashString] or
// [ConfigAgentGeneralPermissionBashMap].
type ConfigAgentGeneralPermissionBashUnion interface {
	implementsConfigAgentGeneralPermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigAgentGeneralPermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigAgentGeneralPermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigAgentGeneralPermissionBashMap{}),
		},
	)
}

type ConfigAgentGeneralPermissionBashString string

const (
	ConfigAgentGeneralPermissionBashStringAsk   ConfigAgentGeneralPermissionBashString = "ask"
	ConfigAgentGeneralPermissionBashStringAllow ConfigAgentGeneralPermissionBashString = "allow"
	ConfigAgentGeneralPermissionBashStringDeny  ConfigAgentGeneralPermissionBashString = "deny"
)

func (r ConfigAgentGeneralPermissionBashString) IsKnown() bool {
	switch r {
	case ConfigAgentGeneralPermissionBashStringAsk, ConfigAgentGeneralPermissionBashStringAllow, ConfigAgentGeneralPermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigAgentGeneralPermissionBashString) implementsConfigAgentGeneralPermissionBashUnion() {}

type ConfigAgentGeneralPermissionBashMap map[string]ConfigAgentGeneralPermissionBashMapItem

func (r ConfigAgentGeneralPermissionBashMap) implementsConfigAgentGeneralPermissionBashUnion() {}

type ConfigAgentGeneralPermissionBashMapItem string

const (
	ConfigAgentGeneralPermissionBashMapAsk   ConfigAgentGeneralPermissionBashMapItem = "ask"
	ConfigAgentGeneralPermissionBashMapAllow ConfigAgentGeneralPermissionBashMapItem = "allow"
	ConfigAgentGeneralPermissionBashMapDeny  ConfigAgentGeneralPermissionBashMapItem = "deny"
)

func (r ConfigAgentGeneralPermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigAgentGeneralPermissionBashMapAsk, ConfigAgentGeneralPermissionBashMapAllow, ConfigAgentGeneralPermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigAgentGeneralPermissionEdit string

const (
	ConfigAgentGeneralPermissionEditAsk   ConfigAgentGeneralPermissionEdit = "ask"
	ConfigAgentGeneralPermissionEditAllow ConfigAgentGeneralPermissionEdit = "allow"
	ConfigAgentGeneralPermissionEditDeny  ConfigAgentGeneralPermissionEdit = "deny"
)

func (r ConfigAgentGeneralPermissionEdit) IsKnown() bool {
	switch r {
	case ConfigAgentGeneralPermissionEditAsk, ConfigAgentGeneralPermissionEditAllow, ConfigAgentGeneralPermissionEditDeny:
		return true
	}
	return false
}

type ConfigAgentGeneralPermissionWebfetch string

const (
	ConfigAgentGeneralPermissionWebfetchAsk   ConfigAgentGeneralPermissionWebfetch = "ask"
	ConfigAgentGeneralPermissionWebfetchAllow ConfigAgentGeneralPermissionWebfetch = "allow"
	ConfigAgentGeneralPermissionWebfetchDeny  ConfigAgentGeneralPermissionWebfetch = "deny"
)

func (r ConfigAgentGeneralPermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigAgentGeneralPermissionWebfetchAsk, ConfigAgentGeneralPermissionWebfetchAllow, ConfigAgentGeneralPermissionWebfetchDeny:
		return true
	}
	return false
}

type ConfigAgentPlan struct {
	// Description of when to use the agent
	Description string                    `json:"description"`
	Disable     bool                      `json:"disable"`
	Mode        ConfigAgentPlanMode       `json:"mode"`
	Model       string                    `json:"model"`
	Permission  ConfigAgentPlanPermission `json:"permission"`
	Prompt      string                    `json:"prompt"`
	Temperature float64                   `json:"temperature"`
	Tools       map[string]bool           `json:"tools"`
	TopP        float64                   `json:"top_p"`
	Variant     string                    `json:"variant"`
	Hidden      bool                      `json:"hidden"`
	Options     map[string]interface{}    `json:"options"`
	Color       string                    `json:"color"`
	Steps       int64                     `json:"steps"`
	MaxSteps    int64                     `json:"maxSteps"`
	ExtraFields map[string]interface{}    `json:"-,extras"`
	JSON        configAgentPlanJSON       `json:"-"`
}

// configAgentPlanJSON contains the JSON metadata for the struct [ConfigAgentPlan]
type configAgentPlanJSON struct {
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

func (r *ConfigAgentPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentPlanJSON) RawJSON() string {
	return r.raw
}

type ConfigAgentPlanMode string

const (
	ConfigAgentPlanModeSubagent ConfigAgentPlanMode = "subagent"
	ConfigAgentPlanModePrimary  ConfigAgentPlanMode = "primary"
	ConfigAgentPlanModeAll      ConfigAgentPlanMode = "all"
)

func (r ConfigAgentPlanMode) IsKnown() bool {
	switch r {
	case ConfigAgentPlanModeSubagent, ConfigAgentPlanModePrimary, ConfigAgentPlanModeAll:
		return true
	}
	return false
}

type ConfigAgentPlanPermission struct {
	// This field can have the runtime type of [ConfigAgentPlanPermissionBashString], [ConfigAgentPlanPermissionBashMap].
	Bash interface{} `json:"bash"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Edit interface{} `json:"edit"`
	// This field can have the runtime type of [string].
	Webfetch interface{} `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}                   `json:"skill"`
	JSON  configAgentPlanPermissionJSON `json:"-"`
}

// configAgentPlanPermissionJSON contains the JSON metadata for the struct
// [ConfigAgentPlanPermission]
type configAgentPlanPermissionJSON struct {
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

func (r *ConfigAgentPlanPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentPlanPermissionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ConfigAgentPlanPermissionBashString] or
// [ConfigAgentPlanPermissionBashMap].
type ConfigAgentPlanPermissionBashUnion interface {
	implementsConfigAgentPlanPermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigAgentPlanPermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigAgentPlanPermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigAgentPlanPermissionBashMap{}),
		},
	)
}

type ConfigAgentPlanPermissionBashString string

const (
	ConfigAgentPlanPermissionBashStringAsk   ConfigAgentPlanPermissionBashString = "ask"
	ConfigAgentPlanPermissionBashStringAllow ConfigAgentPlanPermissionBashString = "allow"
	ConfigAgentPlanPermissionBashStringDeny  ConfigAgentPlanPermissionBashString = "deny"
)

func (r ConfigAgentPlanPermissionBashString) IsKnown() bool {
	switch r {
	case ConfigAgentPlanPermissionBashStringAsk, ConfigAgentPlanPermissionBashStringAllow, ConfigAgentPlanPermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigAgentPlanPermissionBashString) implementsConfigAgentPlanPermissionBashUnion() {}

type ConfigAgentPlanPermissionBashMap map[string]ConfigAgentPlanPermissionBashMapItem

func (r ConfigAgentPlanPermissionBashMap) implementsConfigAgentPlanPermissionBashUnion() {}

type ConfigAgentPlanPermissionBashMapItem string

const (
	ConfigAgentPlanPermissionBashMapAsk   ConfigAgentPlanPermissionBashMapItem = "ask"
	ConfigAgentPlanPermissionBashMapAllow ConfigAgentPlanPermissionBashMapItem = "allow"
	ConfigAgentPlanPermissionBashMapDeny  ConfigAgentPlanPermissionBashMapItem = "deny"
)

func (r ConfigAgentPlanPermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigAgentPlanPermissionBashMapAsk, ConfigAgentPlanPermissionBashMapAllow, ConfigAgentPlanPermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigAgentPlanPermissionEdit string

const (
	ConfigAgentPlanPermissionEditAsk   ConfigAgentPlanPermissionEdit = "ask"
	ConfigAgentPlanPermissionEditAllow ConfigAgentPlanPermissionEdit = "allow"
	ConfigAgentPlanPermissionEditDeny  ConfigAgentPlanPermissionEdit = "deny"
)

func (r ConfigAgentPlanPermissionEdit) IsKnown() bool {
	switch r {
	case ConfigAgentPlanPermissionEditAsk, ConfigAgentPlanPermissionEditAllow, ConfigAgentPlanPermissionEditDeny:
		return true
	}
	return false
}

type ConfigAgentPlanPermissionWebfetch string

const (
	ConfigAgentPlanPermissionWebfetchAsk   ConfigAgentPlanPermissionWebfetch = "ask"
	ConfigAgentPlanPermissionWebfetchAllow ConfigAgentPlanPermissionWebfetch = "allow"
	ConfigAgentPlanPermissionWebfetchDeny  ConfigAgentPlanPermissionWebfetch = "deny"
)

func (r ConfigAgentPlanPermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigAgentPlanPermissionWebfetchAsk, ConfigAgentPlanPermissionWebfetchAllow, ConfigAgentPlanPermissionWebfetchDeny:
		return true
	}
	return false
}

type ConfigAgentExplore struct {
	Description string                       `json:"description"`
	Disable     bool                         `json:"disable"`
	Mode        ConfigAgentExploreMode       `json:"mode"`
	Model       string                       `json:"model"`
	Permission  ConfigAgentExplorePermission `json:"permission"`
	Prompt      string                       `json:"prompt"`
	Temperature float64                      `json:"temperature"`
	Tools       map[string]bool              `json:"tools"`
	TopP        float64                      `json:"top_p"`
	Variant     string                       `json:"variant"`
	Hidden      bool                         `json:"hidden"`
	Options     map[string]interface{}       `json:"options"`
	Color       string                       `json:"color"`
	Steps       int64                        `json:"steps"`
	MaxSteps    int64                        `json:"maxSteps"`
	ExtraFields map[string]interface{}       `json:"-,extras"`
	JSON        configAgentExploreJSON       `json:"-"`
}

// configAgentExploreJSON contains the JSON metadata for the struct
// [ConfigAgentExplore]
type configAgentExploreJSON struct {
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

func (r *ConfigAgentExplore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentExploreJSON) RawJSON() string {
	return r.raw
}

type ConfigAgentExploreMode string

const (
	ConfigAgentExploreModeSubagent ConfigAgentExploreMode = "subagent"
	ConfigAgentExploreModePrimary  ConfigAgentExploreMode = "primary"
	ConfigAgentExploreModeAll      ConfigAgentExploreMode = "all"
)

func (r ConfigAgentExploreMode) IsKnown() bool {
	switch r {
	case ConfigAgentExploreModeSubagent, ConfigAgentExploreModePrimary, ConfigAgentExploreModeAll:
		return true
	}
	return false
}

type ConfigAgentExplorePermission struct {
	// This field can have the runtime type of [ConfigAgentExplorePermissionBashString], [ConfigAgentExplorePermissionBashMap].
	Bash interface{} `json:"bash"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Edit interface{} `json:"edit"`
	// This field can have the runtime type of [string].
	Webfetch interface{} `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}                      `json:"skill"`
	JSON  configAgentExplorePermissionJSON `json:"-"`
}

// configAgentExplorePermissionJSON contains the JSON metadata for the struct
// [ConfigAgentExplorePermission]
type configAgentExplorePermissionJSON struct {
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

func (r *ConfigAgentExplorePermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentExplorePermissionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ConfigAgentExplorePermissionBashString] or
// [ConfigAgentExplorePermissionBashMap].
type ConfigAgentExplorePermissionBashUnion interface {
	implementsConfigAgentExplorePermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigAgentExplorePermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigAgentExplorePermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigAgentExplorePermissionBashMap{}),
		},
	)
}

type ConfigAgentExplorePermissionBashString string

const (
	ConfigAgentExplorePermissionBashStringAsk   ConfigAgentExplorePermissionBashString = "ask"
	ConfigAgentExplorePermissionBashStringAllow ConfigAgentExplorePermissionBashString = "allow"
	ConfigAgentExplorePermissionBashStringDeny  ConfigAgentExplorePermissionBashString = "deny"
)

func (r ConfigAgentExplorePermissionBashString) IsKnown() bool {
	switch r {
	case ConfigAgentExplorePermissionBashStringAsk, ConfigAgentExplorePermissionBashStringAllow, ConfigAgentExplorePermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigAgentExplorePermissionBashString) implementsConfigAgentExplorePermissionBashUnion() {}

type ConfigAgentExplorePermissionBashMap map[string]ConfigAgentExplorePermissionBashMapItem

func (r ConfigAgentExplorePermissionBashMap) implementsConfigAgentExplorePermissionBashUnion() {}

type ConfigAgentExplorePermissionBashMapItem string

const (
	ConfigAgentExplorePermissionBashMapAsk   ConfigAgentExplorePermissionBashMapItem = "ask"
	ConfigAgentExplorePermissionBashMapAllow ConfigAgentExplorePermissionBashMapItem = "allow"
	ConfigAgentExplorePermissionBashMapDeny  ConfigAgentExplorePermissionBashMapItem = "deny"
)

func (r ConfigAgentExplorePermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigAgentExplorePermissionBashMapAsk, ConfigAgentExplorePermissionBashMapAllow, ConfigAgentExplorePermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigAgentExplorePermissionEdit string

const (
	ConfigAgentExplorePermissionEditAsk   ConfigAgentExplorePermissionEdit = "ask"
	ConfigAgentExplorePermissionEditAllow ConfigAgentExplorePermissionEdit = "allow"
	ConfigAgentExplorePermissionEditDeny  ConfigAgentExplorePermissionEdit = "deny"
)

func (r ConfigAgentExplorePermissionEdit) IsKnown() bool {
	switch r {
	case ConfigAgentExplorePermissionEditAsk, ConfigAgentExplorePermissionEditAllow, ConfigAgentExplorePermissionEditDeny:
		return true
	}
	return false
}

type ConfigAgentExplorePermissionWebfetch string

const (
	ConfigAgentExplorePermissionWebfetchAsk   ConfigAgentExplorePermissionWebfetch = "ask"
	ConfigAgentExplorePermissionWebfetchAllow ConfigAgentExplorePermissionWebfetch = "allow"
	ConfigAgentExplorePermissionWebfetchDeny  ConfigAgentExplorePermissionWebfetch = "deny"
)

func (r ConfigAgentExplorePermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigAgentExplorePermissionWebfetchAsk, ConfigAgentExplorePermissionWebfetchAllow, ConfigAgentExplorePermissionWebfetchDeny:
		return true
	}
	return false
}

type ConfigAgentTitle struct {
	Description string                     `json:"description"`
	Disable     bool                       `json:"disable"`
	Mode        ConfigAgentTitleMode       `json:"mode"`
	Model       string                     `json:"model"`
	Permission  ConfigAgentTitlePermission `json:"permission"`
	Prompt      string                     `json:"prompt"`
	Temperature float64                    `json:"temperature"`
	Tools       map[string]bool            `json:"tools"`
	TopP        float64                    `json:"top_p"`
	Variant     string                     `json:"variant"`
	Hidden      bool                       `json:"hidden"`
	Options     map[string]interface{}     `json:"options"`
	Color       string                     `json:"color"`
	Steps       int64                      `json:"steps"`
	MaxSteps    int64                      `json:"maxSteps"`
	ExtraFields map[string]interface{}     `json:"-,extras"`
	JSON        configAgentTitleJSON       `json:"-"`
}

// configAgentTitleJSON contains the JSON metadata for the struct [ConfigAgentTitle]
type configAgentTitleJSON struct {
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

func (r *ConfigAgentTitle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentTitleJSON) RawJSON() string {
	return r.raw
}

type ConfigAgentTitleMode string

const (
	ConfigAgentTitleModeSubagent ConfigAgentTitleMode = "subagent"
	ConfigAgentTitleModePrimary  ConfigAgentTitleMode = "primary"
	ConfigAgentTitleModeAll      ConfigAgentTitleMode = "all"
)

func (r ConfigAgentTitleMode) IsKnown() bool {
	switch r {
	case ConfigAgentTitleModeSubagent, ConfigAgentTitleModePrimary, ConfigAgentTitleModeAll:
		return true
	}
	return false
}

type ConfigAgentTitlePermission struct {
	// This field can have the runtime type of [ConfigAgentTitlePermissionBashString], [ConfigAgentTitlePermissionBashMap].
	Bash interface{} `json:"bash"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Edit interface{} `json:"edit"`
	// This field can have the runtime type of [string].
	Webfetch interface{} `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}                    `json:"skill"`
	JSON  configAgentTitlePermissionJSON `json:"-"`
}

// configAgentTitlePermissionJSON contains the JSON metadata for the struct
// [ConfigAgentTitlePermission]
type configAgentTitlePermissionJSON struct {
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

func (r *ConfigAgentTitlePermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentTitlePermissionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ConfigAgentTitlePermissionBashString] or
// [ConfigAgentTitlePermissionBashMap].
type ConfigAgentTitlePermissionBashUnion interface {
	implementsConfigAgentTitlePermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigAgentTitlePermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigAgentTitlePermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigAgentTitlePermissionBashMap{}),
		},
	)
}

type ConfigAgentTitlePermissionBashString string

const (
	ConfigAgentTitlePermissionBashStringAsk   ConfigAgentTitlePermissionBashString = "ask"
	ConfigAgentTitlePermissionBashStringAllow ConfigAgentTitlePermissionBashString = "allow"
	ConfigAgentTitlePermissionBashStringDeny  ConfigAgentTitlePermissionBashString = "deny"
)

func (r ConfigAgentTitlePermissionBashString) IsKnown() bool {
	switch r {
	case ConfigAgentTitlePermissionBashStringAsk, ConfigAgentTitlePermissionBashStringAllow, ConfigAgentTitlePermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigAgentTitlePermissionBashString) implementsConfigAgentTitlePermissionBashUnion() {}

type ConfigAgentTitlePermissionBashMap map[string]ConfigAgentTitlePermissionBashMapItem

func (r ConfigAgentTitlePermissionBashMap) implementsConfigAgentTitlePermissionBashUnion() {}

type ConfigAgentTitlePermissionBashMapItem string

const (
	ConfigAgentTitlePermissionBashMapAsk   ConfigAgentTitlePermissionBashMapItem = "ask"
	ConfigAgentTitlePermissionBashMapAllow ConfigAgentTitlePermissionBashMapItem = "allow"
	ConfigAgentTitlePermissionBashMapDeny  ConfigAgentTitlePermissionBashMapItem = "deny"
)

func (r ConfigAgentTitlePermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigAgentTitlePermissionBashMapAsk, ConfigAgentTitlePermissionBashMapAllow, ConfigAgentTitlePermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigAgentTitlePermissionEdit string

const (
	ConfigAgentTitlePermissionEditAsk   ConfigAgentTitlePermissionEdit = "ask"
	ConfigAgentTitlePermissionEditAllow ConfigAgentTitlePermissionEdit = "allow"
	ConfigAgentTitlePermissionEditDeny  ConfigAgentTitlePermissionEdit = "deny"
)

func (r ConfigAgentTitlePermissionEdit) IsKnown() bool {
	switch r {
	case ConfigAgentTitlePermissionEditAsk, ConfigAgentTitlePermissionEditAllow, ConfigAgentTitlePermissionEditDeny:
		return true
	}
	return false
}

type ConfigAgentTitlePermissionWebfetch string

const (
	ConfigAgentTitlePermissionWebfetchAsk   ConfigAgentTitlePermissionWebfetch = "ask"
	ConfigAgentTitlePermissionWebfetchAllow ConfigAgentTitlePermissionWebfetch = "allow"
	ConfigAgentTitlePermissionWebfetchDeny  ConfigAgentTitlePermissionWebfetch = "deny"
)

func (r ConfigAgentTitlePermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigAgentTitlePermissionWebfetchAsk, ConfigAgentTitlePermissionWebfetchAllow, ConfigAgentTitlePermissionWebfetchDeny:
		return true
	}
	return false
}

type ConfigAgentSummary struct {
	Description string                       `json:"description"`
	Disable     bool                         `json:"disable"`
	Mode        ConfigAgentSummaryMode       `json:"mode"`
	Model       string                       `json:"model"`
	Permission  ConfigAgentSummaryPermission `json:"permission"`
	Prompt      string                       `json:"prompt"`
	Temperature float64                      `json:"temperature"`
	Tools       map[string]bool              `json:"tools"`
	TopP        float64                      `json:"top_p"`
	Variant     string                       `json:"variant"`
	Hidden      bool                         `json:"hidden"`
	Options     map[string]interface{}       `json:"options"`
	Color       string                       `json:"color"`
	Steps       int64                        `json:"steps"`
	MaxSteps    int64                        `json:"maxSteps"`
	ExtraFields map[string]interface{}       `json:"-,extras"`
	JSON        configAgentSummaryJSON       `json:"-"`
}

// configAgentSummaryJSON contains the JSON metadata for the struct [ConfigAgentSummary]
type configAgentSummaryJSON struct {
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

func (r *ConfigAgentSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentSummaryJSON) RawJSON() string {
	return r.raw
}

type ConfigAgentSummaryMode string

const (
	ConfigAgentSummaryModeSubagent ConfigAgentSummaryMode = "subagent"
	ConfigAgentSummaryModePrimary  ConfigAgentSummaryMode = "primary"
	ConfigAgentSummaryModeAll      ConfigAgentSummaryMode = "all"
)

func (r ConfigAgentSummaryMode) IsKnown() bool {
	switch r {
	case ConfigAgentSummaryModeSubagent, ConfigAgentSummaryModePrimary, ConfigAgentSummaryModeAll:
		return true
	}
	return false
}

type ConfigAgentSummaryPermission struct {
	// This field can have the runtime type of [ConfigAgentSummaryPermissionBashString], [ConfigAgentSummaryPermissionBashMap].
	Bash interface{} `json:"bash"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Edit interface{} `json:"edit"`
	// This field can have the runtime type of [string].
	Webfetch interface{} `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}                      `json:"skill"`
	JSON  configAgentSummaryPermissionJSON `json:"-"`
}

// configAgentSummaryPermissionJSON contains the JSON metadata for the struct
// [ConfigAgentSummaryPermission]
type configAgentSummaryPermissionJSON struct {
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

func (r *ConfigAgentSummaryPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentSummaryPermissionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ConfigAgentSummaryPermissionBashString] or
// [ConfigAgentSummaryPermissionBashMap].
type ConfigAgentSummaryPermissionBashUnion interface {
	implementsConfigAgentSummaryPermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigAgentSummaryPermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigAgentSummaryPermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigAgentSummaryPermissionBashMap{}),
		},
	)
}

type ConfigAgentSummaryPermissionBashString string

const (
	ConfigAgentSummaryPermissionBashStringAsk   ConfigAgentSummaryPermissionBashString = "ask"
	ConfigAgentSummaryPermissionBashStringAllow ConfigAgentSummaryPermissionBashString = "allow"
	ConfigAgentSummaryPermissionBashStringDeny  ConfigAgentSummaryPermissionBashString = "deny"
)

func (r ConfigAgentSummaryPermissionBashString) IsKnown() bool {
	switch r {
	case ConfigAgentSummaryPermissionBashStringAsk, ConfigAgentSummaryPermissionBashStringAllow, ConfigAgentSummaryPermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigAgentSummaryPermissionBashString) implementsConfigAgentSummaryPermissionBashUnion() {}

type ConfigAgentSummaryPermissionBashMap map[string]ConfigAgentSummaryPermissionBashMapItem

func (r ConfigAgentSummaryPermissionBashMap) implementsConfigAgentSummaryPermissionBashUnion() {}

type ConfigAgentSummaryPermissionBashMapItem string

const (
	ConfigAgentSummaryPermissionBashMapAsk   ConfigAgentSummaryPermissionBashMapItem = "ask"
	ConfigAgentSummaryPermissionBashMapAllow ConfigAgentSummaryPermissionBashMapItem = "allow"
	ConfigAgentSummaryPermissionBashMapDeny  ConfigAgentSummaryPermissionBashMapItem = "deny"
)

func (r ConfigAgentSummaryPermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigAgentSummaryPermissionBashMapAsk, ConfigAgentSummaryPermissionBashMapAllow, ConfigAgentSummaryPermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigAgentSummaryPermissionEdit string

const (
	ConfigAgentSummaryPermissionEditAsk   ConfigAgentSummaryPermissionEdit = "ask"
	ConfigAgentSummaryPermissionEditAllow ConfigAgentSummaryPermissionEdit = "allow"
	ConfigAgentSummaryPermissionEditDeny  ConfigAgentSummaryPermissionEdit = "deny"
)

func (r ConfigAgentSummaryPermissionEdit) IsKnown() bool {
	switch r {
	case ConfigAgentSummaryPermissionEditAsk, ConfigAgentSummaryPermissionEditAllow, ConfigAgentSummaryPermissionEditDeny:
		return true
	}
	return false
}

type ConfigAgentSummaryPermissionWebfetch string

const (
	ConfigAgentSummaryPermissionWebfetchAsk   ConfigAgentSummaryPermissionWebfetch = "ask"
	ConfigAgentSummaryPermissionWebfetchAllow ConfigAgentSummaryPermissionWebfetch = "allow"
	ConfigAgentSummaryPermissionWebfetchDeny  ConfigAgentSummaryPermissionWebfetch = "deny"
)

func (r ConfigAgentSummaryPermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigAgentSummaryPermissionWebfetchAsk, ConfigAgentSummaryPermissionWebfetchAllow, ConfigAgentSummaryPermissionWebfetchDeny:
		return true
	}
	return false
}

type ConfigAgentCompaction struct {
	Description string                          `json:"description"`
	Disable     bool                            `json:"disable"`
	Mode        ConfigAgentCompactionMode       `json:"mode"`
	Model       string                          `json:"model"`
	Permission  ConfigAgentCompactionPermission `json:"permission"`
	Prompt      string                          `json:"prompt"`
	Temperature float64                         `json:"temperature"`
	Tools       map[string]bool                 `json:"tools"`
	TopP        float64                         `json:"top_p"`
	Variant     string                          `json:"variant"`
	Hidden      bool                            `json:"hidden"`
	Options     map[string]interface{}          `json:"options"`
	Color       string                          `json:"color"`
	Steps       int64                           `json:"steps"`
	MaxSteps    int64                           `json:"maxSteps"`
	ExtraFields map[string]interface{}          `json:"-,extras"`
	JSON        configAgentCompactionJSON       `json:"-"`
}

// configAgentCompactionJSON contains the JSON metadata for the struct
// [ConfigAgentCompaction]
type configAgentCompactionJSON struct {
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

func (r *ConfigAgentCompaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentCompactionJSON) RawJSON() string {
	return r.raw
}

type ConfigAgentCompactionMode string

const (
	ConfigAgentCompactionModeSubagent ConfigAgentCompactionMode = "subagent"
	ConfigAgentCompactionModePrimary  ConfigAgentCompactionMode = "primary"
	ConfigAgentCompactionModeAll      ConfigAgentCompactionMode = "all"
)

func (r ConfigAgentCompactionMode) IsKnown() bool {
	switch r {
	case ConfigAgentCompactionModeSubagent, ConfigAgentCompactionModePrimary, ConfigAgentCompactionModeAll:
		return true
	}
	return false
}

type ConfigAgentCompactionPermission struct {
	// This field can have the runtime type of [ConfigAgentCompactionPermissionBashString], [ConfigAgentCompactionPermissionBashMap].
	Bash interface{} `json:"bash"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Edit interface{} `json:"edit"`
	// This field can have the runtime type of [string].
	Webfetch interface{} `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}                         `json:"skill"`
	JSON  configAgentCompactionPermissionJSON `json:"-"`
}

// configAgentCompactionPermissionJSON contains the JSON metadata for the struct
// [ConfigAgentCompactionPermission]
type configAgentCompactionPermissionJSON struct {
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

func (r *ConfigAgentCompactionPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configAgentCompactionPermissionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ConfigAgentCompactionPermissionBashString] or
// [ConfigAgentCompactionPermissionBashMap].
type ConfigAgentCompactionPermissionBashUnion interface {
	implementsConfigAgentCompactionPermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigAgentCompactionPermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigAgentCompactionPermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigAgentCompactionPermissionBashMap{}),
		},
	)
}

type ConfigAgentCompactionPermissionBashString string

const (
	ConfigAgentCompactionPermissionBashStringAsk   ConfigAgentCompactionPermissionBashString = "ask"
	ConfigAgentCompactionPermissionBashStringAllow ConfigAgentCompactionPermissionBashString = "allow"
	ConfigAgentCompactionPermissionBashStringDeny  ConfigAgentCompactionPermissionBashString = "deny"
)

func (r ConfigAgentCompactionPermissionBashString) IsKnown() bool {
	switch r {
	case ConfigAgentCompactionPermissionBashStringAsk, ConfigAgentCompactionPermissionBashStringAllow, ConfigAgentCompactionPermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigAgentCompactionPermissionBashString) implementsConfigAgentCompactionPermissionBashUnion() {
}

type ConfigAgentCompactionPermissionBashMap map[string]ConfigAgentCompactionPermissionBashMapItem

func (r ConfigAgentCompactionPermissionBashMap) implementsConfigAgentCompactionPermissionBashUnion() {
}

type ConfigAgentCompactionPermissionBashMapItem string

const (
	ConfigAgentCompactionPermissionBashMapAsk   ConfigAgentCompactionPermissionBashMapItem = "ask"
	ConfigAgentCompactionPermissionBashMapAllow ConfigAgentCompactionPermissionBashMapItem = "allow"
	ConfigAgentCompactionPermissionBashMapDeny  ConfigAgentCompactionPermissionBashMapItem = "deny"
)

func (r ConfigAgentCompactionPermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigAgentCompactionPermissionBashMapAsk, ConfigAgentCompactionPermissionBashMapAllow, ConfigAgentCompactionPermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigAgentCompactionPermissionEdit string

const (
	ConfigAgentCompactionPermissionEditAsk   ConfigAgentCompactionPermissionEdit = "ask"
	ConfigAgentCompactionPermissionEditAllow ConfigAgentCompactionPermissionEdit = "allow"
	ConfigAgentCompactionPermissionEditDeny  ConfigAgentCompactionPermissionEdit = "deny"
)

func (r ConfigAgentCompactionPermissionEdit) IsKnown() bool {
	switch r {
	case ConfigAgentCompactionPermissionEditAsk, ConfigAgentCompactionPermissionEditAllow, ConfigAgentCompactionPermissionEditDeny:
		return true
	}
	return false
}

type ConfigAgentCompactionPermissionWebfetch string

const (
	ConfigAgentCompactionPermissionWebfetchAsk   ConfigAgentCompactionPermissionWebfetch = "ask"
	ConfigAgentCompactionPermissionWebfetchAllow ConfigAgentCompactionPermissionWebfetch = "allow"
	ConfigAgentCompactionPermissionWebfetchDeny  ConfigAgentCompactionPermissionWebfetch = "deny"
)

func (r ConfigAgentCompactionPermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigAgentCompactionPermissionWebfetchAsk, ConfigAgentCompactionPermissionWebfetchAllow, ConfigAgentCompactionPermissionWebfetchDeny:
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
	Command  interface{} `json:"command"`
	Disabled bool        `json:"disabled"`
	// This field can have the runtime type of [map[string]string].
	Env interface{} `json:"env"`
	// This field can have the runtime type of [[]string].
	Extensions interface{} `json:"extensions"`
	// This field can have the runtime type of [map[string]interface{}].
	Initialization interface{}   `json:"initialization"`
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
		reflect.TypeOf((*ConfigLspUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigLspDisabled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigLspObject{}),
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
	Command        []string               `json:"command,required"`
	Disabled       bool                   `json:"disabled"`
	Env            map[string]string      `json:"env"`
	Extensions     []string               `json:"extensions"`
	Initialization map[string]interface{} `json:"initialization"`
	JSON           configLspObjectJSON    `json:"-"`
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
	Command interface{} `json:"command"`
	// This field can have the runtime type of [string, nil]. Working directory for the MCP server process (for "local" type).
	Cwd interface{} `json:"cwd"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [map[string]string]. Environment variables to set when running the MCP server (for "local" type).
	Environment interface{} `json:"environment"`
	// This field can have the runtime type of [map[string]string]. Headers to send with the request (for "remote" type).
	Headers interface{} `json:"headers"`
	// This field can have the runtime type of [McpOAuthConfig, nil]. OAuth authentication configuration for the MCP server (for "remote" type).
	OAuth interface{} `json:"oauth"`
	// This field can have the runtime type of [int64, nil]. Timeout in milliseconds for MCP server requests.
	Timeout interface{} `json:"timeout"`
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
		reflect.TypeOf((*ConfigMcpUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McpLocalConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McpRemoteConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigMcpDisabled{}),
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
	Build       ConfigModeBuild        `json:"build"`
	Plan        ConfigModePlan         `json:"plan"`
	ExtraFields map[string]ConfigAgent `json:"-,extras"`
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

type ConfigModeBuild struct {
	// Description of when to use the agent
	Description string                    `json:"description"`
	Disable     bool                      `json:"disable"`
	Mode        ConfigModeBuildMode       `json:"mode"`
	Model       string                    `json:"model"`
	Permission  ConfigModeBuildPermission `json:"permission"`
	Prompt      string                    `json:"prompt"`
	Temperature float64                   `json:"temperature"`
	Tools       map[string]bool           `json:"tools"`
	TopP        float64                   `json:"top_p"`
	Variant     string                    `json:"variant"`
	Hidden      bool                      `json:"hidden"`
	Options     map[string]interface{}    `json:"options"`
	Color       string                    `json:"color"`
	Steps       int64                     `json:"steps"`
	MaxSteps    int64                     `json:"maxSteps"`
	ExtraFields map[string]interface{}    `json:"-,extras"`
	JSON        configModeBuildJSON       `json:"-"`
}

// configModeBuildJSON contains the JSON metadata for the struct [ConfigModeBuild]
type configModeBuildJSON struct {
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

func (r *ConfigModeBuild) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configModeBuildJSON) RawJSON() string {
	return r.raw
}

type ConfigModeBuildMode string

const (
	ConfigModeBuildModeSubagent ConfigModeBuildMode = "subagent"
	ConfigModeBuildModePrimary  ConfigModeBuildMode = "primary"
	ConfigModeBuildModeAll      ConfigModeBuildMode = "all"
)

func (r ConfigModeBuildMode) IsKnown() bool {
	switch r {
	case ConfigModeBuildModeSubagent, ConfigModeBuildModePrimary, ConfigModeBuildModeAll:
		return true
	}
	return false
}

type ConfigModeBuildPermission struct {
	// This field can have the runtime type of [ConfigModeBuildPermissionBashString], [ConfigModeBuildPermissionBashMap].
	Bash interface{} `json:"bash"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Edit interface{} `json:"edit"`
	// This field can have the runtime type of [string].
	Webfetch interface{} `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}                   `json:"skill"`
	JSON  configModeBuildPermissionJSON `json:"-"`
}

// configModeBuildPermissionJSON contains the JSON metadata for the struct
// [ConfigModeBuildPermission]
type configModeBuildPermissionJSON struct {
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

func (r *ConfigModeBuildPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configModeBuildPermissionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ConfigModeBuildPermissionBashString] or
// [ConfigModeBuildPermissionBashMap].
type ConfigModeBuildPermissionBashUnion interface {
	implementsConfigModeBuildPermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigModeBuildPermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigModeBuildPermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigModeBuildPermissionBashMap{}),
		},
	)
}

type ConfigModeBuildPermissionBashString string

const (
	ConfigModeBuildPermissionBashStringAsk   ConfigModeBuildPermissionBashString = "ask"
	ConfigModeBuildPermissionBashStringAllow ConfigModeBuildPermissionBashString = "allow"
	ConfigModeBuildPermissionBashStringDeny  ConfigModeBuildPermissionBashString = "deny"
)

func (r ConfigModeBuildPermissionBashString) IsKnown() bool {
	switch r {
	case ConfigModeBuildPermissionBashStringAsk, ConfigModeBuildPermissionBashStringAllow, ConfigModeBuildPermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigModeBuildPermissionBashString) implementsConfigModeBuildPermissionBashUnion() {}

type ConfigModeBuildPermissionBashMap map[string]ConfigModeBuildPermissionBashMapItem

func (r ConfigModeBuildPermissionBashMap) implementsConfigModeBuildPermissionBashUnion() {}

type ConfigModeBuildPermissionBashMapItem string

const (
	ConfigModeBuildPermissionBashMapAsk   ConfigModeBuildPermissionBashMapItem = "ask"
	ConfigModeBuildPermissionBashMapAllow ConfigModeBuildPermissionBashMapItem = "allow"
	ConfigModeBuildPermissionBashMapDeny  ConfigModeBuildPermissionBashMapItem = "deny"
)

func (r ConfigModeBuildPermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigModeBuildPermissionBashMapAsk, ConfigModeBuildPermissionBashMapAllow, ConfigModeBuildPermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigModeBuildPermissionEdit string

const (
	ConfigModeBuildPermissionEditAsk   ConfigModeBuildPermissionEdit = "ask"
	ConfigModeBuildPermissionEditAllow ConfigModeBuildPermissionEdit = "allow"
	ConfigModeBuildPermissionEditDeny  ConfigModeBuildPermissionEdit = "deny"
)

func (r ConfigModeBuildPermissionEdit) IsKnown() bool {
	switch r {
	case ConfigModeBuildPermissionEditAsk, ConfigModeBuildPermissionEditAllow, ConfigModeBuildPermissionEditDeny:
		return true
	}
	return false
}

type ConfigModeBuildPermissionWebfetch string

const (
	ConfigModeBuildPermissionWebfetchAsk   ConfigModeBuildPermissionWebfetch = "ask"
	ConfigModeBuildPermissionWebfetchAllow ConfigModeBuildPermissionWebfetch = "allow"
	ConfigModeBuildPermissionWebfetchDeny  ConfigModeBuildPermissionWebfetch = "deny"
)

func (r ConfigModeBuildPermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigModeBuildPermissionWebfetchAsk, ConfigModeBuildPermissionWebfetchAllow, ConfigModeBuildPermissionWebfetchDeny:
		return true
	}
	return false
}

type ConfigModePlan struct {
	// Description of when to use the agent
	Description string                   `json:"description"`
	Disable     bool                     `json:"disable"`
	Mode        ConfigModePlanMode       `json:"mode"`
	Model       string                   `json:"model"`
	Permission  ConfigModePlanPermission `json:"permission"`
	Prompt      string                   `json:"prompt"`
	Temperature float64                  `json:"temperature"`
	Tools       map[string]bool          `json:"tools"`
	TopP        float64                  `json:"top_p"`
	Variant     string                   `json:"variant"`
	Hidden      bool                     `json:"hidden"`
	Options     map[string]interface{}   `json:"options"`
	Color       string                   `json:"color"`
	Steps       int64                    `json:"steps"`
	MaxSteps    int64                    `json:"maxSteps"`
	ExtraFields map[string]interface{}   `json:"-,extras"`
	JSON        configModePlanJSON       `json:"-"`
}

// configModePlanJSON contains the JSON metadata for the struct [ConfigModePlan]
type configModePlanJSON struct {
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

func (r *ConfigModePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configModePlanJSON) RawJSON() string {
	return r.raw
}

type ConfigModePlanMode string

const (
	ConfigModePlanModeSubagent ConfigModePlanMode = "subagent"
	ConfigModePlanModePrimary  ConfigModePlanMode = "primary"
	ConfigModePlanModeAll      ConfigModePlanMode = "all"
)

func (r ConfigModePlanMode) IsKnown() bool {
	switch r {
	case ConfigModePlanModeSubagent, ConfigModePlanModePrimary, ConfigModePlanModeAll:
		return true
	}
	return false
}

type ConfigModePlanPermission struct {
	// This field can have the runtime type of [ConfigModePlanPermissionBashString], [ConfigModePlanPermissionBashMap].
	Bash interface{} `json:"bash"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Edit interface{} `json:"edit"`
	// This field can have the runtime type of [string].
	Webfetch interface{} `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}                  `json:"skill"`
	JSON  configModePlanPermissionJSON `json:"-"`
}

// configModePlanPermissionJSON contains the JSON metadata for the struct
// [ConfigModePlanPermission]
type configModePlanPermissionJSON struct {
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

func (r *ConfigModePlanPermission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configModePlanPermissionJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ConfigModePlanPermissionBashString] or
// [ConfigModePlanPermissionBashMap].
type ConfigModePlanPermissionBashUnion interface {
	implementsConfigModePlanPermissionBashUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigModePlanPermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigModePlanPermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigModePlanPermissionBashMap{}),
		},
	)
}

type ConfigModePlanPermissionBashString string

const (
	ConfigModePlanPermissionBashStringAsk   ConfigModePlanPermissionBashString = "ask"
	ConfigModePlanPermissionBashStringAllow ConfigModePlanPermissionBashString = "allow"
	ConfigModePlanPermissionBashStringDeny  ConfigModePlanPermissionBashString = "deny"
)

func (r ConfigModePlanPermissionBashString) IsKnown() bool {
	switch r {
	case ConfigModePlanPermissionBashStringAsk, ConfigModePlanPermissionBashStringAllow, ConfigModePlanPermissionBashStringDeny:
		return true
	}
	return false
}

func (r ConfigModePlanPermissionBashString) implementsConfigModePlanPermissionBashUnion() {}

type ConfigModePlanPermissionBashMap map[string]ConfigModePlanPermissionBashMapItem

func (r ConfigModePlanPermissionBashMap) implementsConfigModePlanPermissionBashUnion() {}

type ConfigModePlanPermissionBashMapItem string

const (
	ConfigModePlanPermissionBashMapAsk   ConfigModePlanPermissionBashMapItem = "ask"
	ConfigModePlanPermissionBashMapAllow ConfigModePlanPermissionBashMapItem = "allow"
	ConfigModePlanPermissionBashMapDeny  ConfigModePlanPermissionBashMapItem = "deny"
)

func (r ConfigModePlanPermissionBashMapItem) IsKnown() bool {
	switch r {
	case ConfigModePlanPermissionBashMapAsk, ConfigModePlanPermissionBashMapAllow, ConfigModePlanPermissionBashMapDeny:
		return true
	}
	return false
}

type ConfigModePlanPermissionEdit string

const (
	ConfigModePlanPermissionEditAsk   ConfigModePlanPermissionEdit = "ask"
	ConfigModePlanPermissionEditAllow ConfigModePlanPermissionEdit = "allow"
	ConfigModePlanPermissionEditDeny  ConfigModePlanPermissionEdit = "deny"
)

func (r ConfigModePlanPermissionEdit) IsKnown() bool {
	switch r {
	case ConfigModePlanPermissionEditAsk, ConfigModePlanPermissionEditAllow, ConfigModePlanPermissionEditDeny:
		return true
	}
	return false
}

type ConfigModePlanPermissionWebfetch string

const (
	ConfigModePlanPermissionWebfetchAsk   ConfigModePlanPermissionWebfetch = "ask"
	ConfigModePlanPermissionWebfetchAllow ConfigModePlanPermissionWebfetch = "allow"
	ConfigModePlanPermissionWebfetchDeny  ConfigModePlanPermissionWebfetch = "deny"
)

func (r ConfigModePlanPermissionWebfetch) IsKnown() bool {
	switch r {
	case ConfigModePlanPermissionWebfetchAsk, ConfigModePlanPermissionWebfetchAllow, ConfigModePlanPermissionWebfetchDeny:
		return true
	}
	return false
}

type ConfigPermission struct {
	// This field can have the runtime type of [ConfigPermissionBashString], [ConfigPermissionBashMap].
	Bash     interface{}              `json:"bash"`
	Edit     ConfigPermissionEdit     `json:"edit"`
	Webfetch ConfigPermissionWebfetch `json:"webfetch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Read interface{} `json:"read"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Glob interface{} `json:"glob"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Grep interface{} `json:"grep"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	List interface{} `json:"list"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Task interface{} `json:"task"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	ExternalDirectory interface{} `json:"external_directory"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Todowrite interface{} `json:"todowrite"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Question interface{} `json:"question"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Websearch interface{} `json:"websearch"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Lsp interface{} `json:"lsp"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	DoomLoop interface{} `json:"doom_loop"`
	// This field can have the runtime type of [string] or [map[string]interface{}].
	Skill interface{}          `json:"skill"`
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
		reflect.TypeOf((*ConfigPermissionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigPermissionAction("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigPermission{}),
		},
	)
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigPermissionBashUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigPermissionBashString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigPermissionBashMap{}),
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

type ConfigPermissionEdit string

const (
	ConfigPermissionEditAsk   ConfigPermissionEdit = "ask"
	ConfigPermissionEditAllow ConfigPermissionEdit = "allow"
	ConfigPermissionEditDeny  ConfigPermissionEdit = "deny"
)

func (r ConfigPermissionEdit) IsKnown() bool {
	switch r {
	case ConfigPermissionEditAsk, ConfigPermissionEditAllow, ConfigPermissionEditDeny:
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
	Interleaved interface{}                    `json:"interleaved"`
	Limit       ConfigProviderModelsLimit      `json:"limit"`
	Modalities  ConfigProviderModelsModalities `json:"modalities"`
	Name        string                         `json:"name"`
	Options     map[string]interface{}         `json:"options"`
	Provider    ConfigProviderModelsProvider   `json:"provider"`
	Reasoning   bool                           `json:"reasoning"`
	ReleaseDate string                         `json:"release_date"`
	Status      ConfigProviderModelsStatus     `json:"status"`
	Temperature bool                           `json:"temperature"`
	ToolCall    bool                           `json:"tool_call"`
	// This field can have the runtime type of object.
	Variants interface{}             `json:"variants"`
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
	Timeout interface{} `json:"timeout"`
	// Timeout in milliseconds to wait for response headers. Provider integrations
	// may set defaults. Set to false to disable timeout.
	// This field can have the runtime type of [shared.UnionInt], [shared.UnionBool].
	HeaderTimeout interface{}               `json:"headerTimeout"`
	ChunkTimeout  int64                     `json:"chunkTimeout"`
	ExtraFields   map[string]interface{}    `json:"-,extras"`
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
		reflect.TypeOf((*ConfigProviderOptionsTimeoutUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
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
	OAuth interface{}         `json:"oauth"`
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
	Autoupdate        param.Field[interface{}]                   `json:"autoupdate"`
	Command           param.Field[map[string]ConfigCommandParam] `json:"command"`
	Compaction        param.Field[ConfigCompactionParam]         `json:"compaction"`
	DisabledProviders param.Field[[]string]                      `json:"disabled_providers"`
	EnabledProviders  param.Field[[]string]                      `json:"enabled_providers"`
	Enterprise        param.Field[EnterpriseConfigParam]         `json:"enterprise"`
	Experimental      param.Field[ConfigExperimentalParam]       `json:"experimental"`
	// Enable or configure formatters. Pass false to disable, true to enable
	// built-ins, or a map of formatter-name to config to enable with overrides.
	// Accepts [bool] or [map[string]ConfigFormatter].
	Formatter    param.Field[interface{}]    `json:"formatter"`
	Instructions param.Field[[]string]       `json:"instructions"`
	Layout       param.Field[ConfigLayout]   `json:"layout"`
	LogLevel     param.Field[ConfigLogLevel] `json:"logLevel"`
	// Enable or configure LSP servers. Pass false to disable, true to enable
	// built-ins, or a map of lsp-name to config to enable with overrides.
	// Accepts [bool] or [map[string]ConfigLsp].
	Lsp   param.Field[interface{}]                    `json:"lsp"`
	Mcp   param.Field[map[string]ConfigMcpUnionParam] `json:"mcp"`
	Mode  param.Field[ConfigModeParam]                `json:"mode"`
	Model param.Field[string]                         `json:"model"`
	// Permission configuration. A short string ("ask"|"allow"|"deny") or an
	// object with per-action permission rule overrides. Accepts [ConfigPermissionAction]
	// (a string constant) or [ConfigPermissionParam].
	Permission param.Field[ConfigPermissionUnionParam] `json:"permission"`
	// Plugins to load. Each item is either a plugin name (string) or a 2-tuple
	// of [pluginName, configObject] (where configObject is a map[string]any).
	Plugin   param.Field[[]interface{}]                  `json:"plugin"`
	Provider param.Field[map[string]ConfigProviderParam] `json:"provider"`
	// Map of reference name → value. Each value can be a plain [string] (URL/path),
	// a [ConfigV2ReferenceGit], or a [ConfigV2ReferenceLocal].
	Reference param.Field[map[string]interface{}] `json:"reference"`
	// Map of reference name → value. Each value can be a plain [string] (URL/path),
	// a [ConfigV2ReferenceGit], or a [ConfigV2ReferenceLocal].
	References    param.Field[map[string]interface{}] `json:"references"`
	Share         param.Field[ConfigShare]            `json:"share"`
	Shell         param.Field[string]                 `json:"shell"`
	Server        param.Field[ServerConfigParam]      `json:"server"`
	Skills        param.Field[ConfigSkillsParam]      `json:"skills"`
	SmallModel    param.Field[string]                 `json:"small_model"`
	Snapshot      param.Field[bool]                   `json:"snapshot"`
	ToolOutput    param.Field[ConfigToolOutputParam]  `json:"tool_output"`
	Tools         param.Field[map[string]bool]        `json:"tools"`
	Username      param.Field[string]                 `json:"username"`
	Watcher       param.Field[ConfigWatcherParam]     `json:"watcher"`
	DefaultAgent  param.Field[string]                 `json:"default_agent"`
	SubagentDepth param.Field[int64]                  `json:"subagent_depth"`
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
		reflect.TypeOf((*ConfigV2ReferenceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigV2ReferenceGit{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigV2ReferenceLocal{}),
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

// ConfigPermissionParam is the request-side counterpart of [ConfigPermission].
// It is also the object variant of the [ConfigPermissionUnionParam] request union.
type ConfigPermissionParam struct {
	// Accepts [ConfigPermissionBashString] or [map[string]ConfigPermissionBashMapItem].
	Bash              param.Field[interface{}] `json:"bash"`
	Edit              param.Field[interface{}] `json:"edit"`
	Webfetch          param.Field[interface{}] `json:"webfetch"`
	Read              param.Field[interface{}] `json:"read"`
	Glob              param.Field[interface{}] `json:"glob"`
	Grep              param.Field[interface{}] `json:"grep"`
	List              param.Field[interface{}] `json:"list"`
	Task              param.Field[interface{}] `json:"task"`
	ExternalDirectory param.Field[interface{}] `json:"external_directory"`
	Todowrite         param.Field[interface{}] `json:"todowrite"`
	Question          param.Field[interface{}] `json:"question"`
	Websearch         param.Field[interface{}] `json:"websearch"`
	Lsp               param.Field[interface{}] `json:"lsp"`
	DoomLoop          param.Field[interface{}] `json:"doom_loop"`
	Skill             param.Field[interface{}] `json:"skill"`
}

func (r ConfigPermissionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigPermissionParam) implementsConfigPermissionUnionParam() {}

// ConfigPermissionUnionParam is the request-side union for the permission field.
//
// Satisfied by [ConfigPermissionAction] (a short string "ask"|"allow"|"deny")
// or [ConfigPermissionParam] (per-action permission rule overrides).
type ConfigPermissionUnionParam interface {
	implementsConfigPermissionUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigPermissionUnionParam)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ConfigPermissionAction("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigPermissionParam{}),
		},
	)
}

// ConfigPermissionAction already implements ConfigPermissionUnionParam.
func (r ConfigPermissionAction) implementsConfigPermissionUnionParam() {}

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
		reflect.TypeOf((*ConfigMcpUnionParam)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McpLocalConfigParam{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McpRemoteConfigParam{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigMcpDisabledParam{}),
		},
	)
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*McpOAuthConfigUnionParam)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McpOAuthConfigParam{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(McpOAuthConfigDisabledParam{}),
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
	Timeout param.Field[interface{}] `json:"timeout"`
	// Timeout in milliseconds to wait for response headers. Set to false to disable.
	// Accepts [shared.UnionInt] or [shared.UnionBool].
	HeaderTimeout param.Field[interface{}] `json:"headerTimeout"`
	ChunkTimeout  param.Field[int64]       `json:"chunkTimeout"`
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
	// Accepts bool or object.
	Interleaved param.Field[interface{}]                         `json:"interleaved"`
	Name        param.Field[string]                              `json:"name"`
	Options     param.Field[map[string]interface{}]              `json:"options"`
	Reasoning   param.Field[bool]                                `json:"reasoning"`
	ReleaseDate param.Field[string]                              `json:"release_date"`
	Temperature param.Field[bool]                                `json:"temperature"`
	ToolCall    param.Field[bool]                                `json:"tool_call"`
	Cost        param.Field[ConfigProviderModelsCostParam]       `json:"cost"`
	Limit       param.Field[ConfigProviderModelsLimitParam]      `json:"limit"`
	Modalities  param.Field[ConfigProviderModelsModalitiesParam] `json:"modalities"`
	Provider    param.Field[ConfigProviderModelsProviderParam]   `json:"provider"`
	Status      param.Field[ConfigProviderModelsStatus]          `json:"status"`
	// Accepts object.
	Variants param.Field[interface{}] `json:"variants"`
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

// configAgentSubParam is a shared param shape for all ConfigAgent sub-types
// (Build, General, Plan, Explore, Title, Summary, Compaction).
// All sub-agent config structs share the same fields; use interface{} for the
// permission field since it accepts a complex per-tool bash/string/map union.
type configAgentSubParam struct {
	Description param.Field[string] `json:"description"`
	Disable     param.Field[bool]   `json:"disable"`
	Mode        param.Field[string] `json:"mode"`
	Model       param.Field[string] `json:"model"`
	// Accepts string ("ask"|"allow"|"deny") or a per-tool permission map object.
	Permission  param.Field[interface{}]            `json:"permission"`
	Prompt      param.Field[string]                 `json:"prompt"`
	Temperature param.Field[float64]                `json:"temperature"`
	Tools       param.Field[map[string]bool]        `json:"tools"`
	TopP        param.Field[float64]                `json:"top_p"`
	Variant     param.Field[string]                 `json:"variant"`
	Hidden      param.Field[bool]                   `json:"hidden"`
	Options     param.Field[map[string]interface{}] `json:"options"`
	Color       param.Field[string]                 `json:"color"`
	Steps       param.Field[int64]                  `json:"steps"`
	MaxSteps    param.Field[int64]                  `json:"maxSteps"`
}

func (r configAgentSubParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigAgentBuildParam is the request-side counterpart of [ConfigAgentBuild].
type ConfigAgentBuildParam = configAgentSubParam

// ConfigAgentGeneralParam is the request-side counterpart of [ConfigAgentGeneral].
type ConfigAgentGeneralParam = configAgentSubParam

// ConfigAgentPlanParam is the request-side counterpart of [ConfigAgentPlan].
type ConfigAgentPlanParam = configAgentSubParam

// ConfigAgentExploreParam is the request-side counterpart of [ConfigAgentExplore].
type ConfigAgentExploreParam = configAgentSubParam

// ConfigAgentTitleParam is the request-side counterpart of [ConfigAgentTitle].
type ConfigAgentTitleParam = configAgentSubParam

// ConfigAgentSummaryParam is the request-side counterpart of [ConfigAgentSummary].
type ConfigAgentSummaryParam = configAgentSubParam

// ConfigAgentCompactionParam is the request-side counterpart of [ConfigAgentCompaction].
type ConfigAgentCompactionParam = configAgentSubParam

// ConfigAgentParam is the request-side counterpart of [ConfigAgent].
// ExtraFields allows passing arbitrary named agent configs beyond the named sub-agents.
type ConfigAgentParam struct {
	Build      param.Field[ConfigAgentBuildParam]      `json:"build"`
	Compaction param.Field[ConfigAgentCompactionParam] `json:"compaction"`
	Explore    param.Field[ConfigAgentExploreParam]    `json:"explore"`
	General    param.Field[ConfigAgentGeneralParam]    `json:"general"`
	Plan       param.Field[ConfigAgentPlanParam]       `json:"plan"`
	Summary    param.Field[ConfigAgentSummaryParam]    `json:"summary"`
	Title      param.Field[ConfigAgentTitleParam]      `json:"title"`
}

func (r ConfigAgentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigModeParam is the request-side counterpart of [ConfigMode].
//
// @deprecated Use ConfigAgentParam instead.
type ConfigModeParam struct {
	Build param.Field[ConfigModeBuildParam] `json:"build"`
	Plan  param.Field[ConfigModePlanParam]  `json:"plan"`
}

func (r ConfigModeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigModeBuildParam is the request-side counterpart of [ConfigModeBuild].
type ConfigModeBuildParam struct {
	Description param.Field[string] `json:"description"`
	Disable     param.Field[bool]   `json:"disable"`
	Mode        param.Field[string] `json:"mode"`
	Model       param.Field[string] `json:"model"`
	// Accepts string ("ask"|"allow"|"deny") or a per-tool permission map object.
	Permission  param.Field[interface{}]            `json:"permission"`
	Prompt      param.Field[string]                 `json:"prompt"`
	Temperature param.Field[float64]                `json:"temperature"`
	Tools       param.Field[map[string]bool]        `json:"tools"`
	TopP        param.Field[float64]                `json:"top_p"`
	Variant     param.Field[string]                 `json:"variant"`
	Hidden      param.Field[bool]                   `json:"hidden"`
	Options     param.Field[map[string]interface{}] `json:"options"`
	Color       param.Field[string]                 `json:"color"`
	Steps       param.Field[int64]                  `json:"steps"`
	MaxSteps    param.Field[int64]                  `json:"maxSteps"`
}

func (r ConfigModeBuildParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigModePlanParam is the request-side counterpart of [ConfigModePlan].
type ConfigModePlanParam struct {
	Description param.Field[string] `json:"description"`
	Disable     param.Field[bool]   `json:"disable"`
	Mode        param.Field[string] `json:"mode"`
	Model       param.Field[string] `json:"model"`
	// Accepts string ("ask"|"allow"|"deny") or a per-tool permission map object.
	Permission  param.Field[interface{}]            `json:"permission"`
	Prompt      param.Field[string]                 `json:"prompt"`
	Temperature param.Field[float64]                `json:"temperature"`
	Tools       param.Field[map[string]bool]        `json:"tools"`
	TopP        param.Field[float64]                `json:"top_p"`
	Variant     param.Field[string]                 `json:"variant"`
	Hidden      param.Field[bool]                   `json:"hidden"`
	Options     param.Field[map[string]interface{}] `json:"options"`
	Color       param.Field[string]                 `json:"color"`
	Steps       param.Field[int64]                  `json:"steps"`
	MaxSteps    param.Field[int64]                  `json:"maxSteps"`
}

func (r ConfigModePlanParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
