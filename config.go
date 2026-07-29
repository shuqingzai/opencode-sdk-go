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
	// This field can have the runtime type of [bool], [string].
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
	// Per OpenAPI `Config.formatter` is `anyOf [boolean, object]`, where the object
	// maps formatter names to [ConfigFormatter] overrides.
	// This field can have the runtime type of [bool], [map[string]ConfigFormatter].
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
	//
	// Permission is the OpenAPI `PermissionConfig` anyOf; the decoder selects the
	// concrete variant structurally, so a nil value means the field was absent, null
	// or matched no variant.
	Permission ConfigPermissionUnion `json:"permission"`
	// This field can have the runtime type of [string], [[]any].
	Plugin []any `json:"plugin"`
	// Custom provider configurations and model overrides
	Provider map[string]ConfigProvider `json:"provider"`
	// Reference configuration for external documentation. Keys are reference
	// names, values can be a plain URL/path string or a structured config (git
	// or local).
	// Each value can have the runtime type of [ConfigV2ReferenceString],
	// [ConfigV2ReferenceGit], [ConfigV2ReferenceLocal].
	Reference map[string]any `json:"reference"`
	// References from external sources. Keys are reference names, values can be a
	// plain URL/path string or a structured config (git or local).
	// Each value can have the runtime type of [ConfigV2ReferenceString],
	// [ConfigV2ReferenceGit], [ConfigV2ReferenceLocal].
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
	// `formatter`, `lsp`, `reference` and `references` are all anyOf-typed and land
	// on `any` carrier fields, so the generic interface branch of the decoder
	// (internal/apijson/decoder.go, `case reflect.Interface`) would hand back
	// `map[string]any` for every object payload and the declared runtime types would
	// be unreachable. Route the raw sub-documents explicitly instead.
	//
	// `permission` needs no such help: it is declared as [ConfigPermissionUnion], so
	// the decoder resolves it through the registered union natively.
	r.Formatter = configObjectMapField[ConfigFormatter](data, "formatter", r.Formatter)
	r.Lsp = configObjectMapField[ConfigLsp](data, "lsp", r.Lsp)
	r.Reference = configReferenceField(data, "reference", r.Reference)
	r.References = configReferenceField(data, "references", r.References)
	return nil
}

// configObjectMapField routes the object arm of an OpenAPI
// `anyOf [boolean, object(additionalProperties: <shape>)]` field onto
// `map[string]T`, which is what the runtime comments on [Config.Formatter] and
// [Config.Lsp] declare. The boolean arm needs no help: a JSON scalar already
// decodes to a Go `bool` through the decoder's generic interface branch; only the
// object arm degrades to `map[string]any`.
//
// A payload matching neither arm keeps whatever the generic decoder produced
// rather than failing, so a single unreadable field cannot abort the decode of the
// whole `/config` document. The raw bytes stay available via [Config.RawJSON].
//
// Entries decode through `T`'s own [json.Unmarshaler], so [ConfigLsp]'s nested
// `anyOf [{disabled: true}, {command, ...}]` is shape-routed here as well.
func configObjectMapField[T any](data []byte, key string, fallback any) any {
	node := gjson.GetBytes(data, key)
	if node.Type != gjson.JSON || !node.IsObject() {
		return fallback
	}
	typed := map[string]T{}
	if err := apijson.UnmarshalRoot([]byte(node.Raw), &typed); err != nil {
		return fallback
	}
	return typed
}

// configReferenceField routes every entry of the OpenAPI `Config.reference` /
// `Config.references` object (`additionalProperties: anyOf [string,
// ConfigV2ReferenceGit, ConfigV2ReferenceLocal]`) onto its declared variant.
//
// The map itself stays `map[string]any` -- the OpenAPI value is a union and a
// Response carrier field may not be typed as a union interface -- but the values
// become the concrete variants named by the runtime comment instead of the
// `string` / `map[string]any` the generic decoder produces.
//
// Entries that match no variant keep the generically decoded value, so one drifted
// reference cannot cost the caller the rest of the document.
func configReferenceField(data []byte, key string, fallback map[string]any) map[string]any {
	node := gjson.GetBytes(data, key)
	if node.Type != gjson.JSON || !node.IsObject() {
		return fallback
	}
	routed := make(map[string]any, len(fallback))
	for name, generic := range fallback {
		routed[name] = generic
	}
	node.ForEach(func(key, value gjson.Result) bool {
		if variant, ok := configV2Reference(value); ok {
			routed[key.String()] = variant
		}
		return true
	})
	return routed
}

// configV2Reference picks the variant of the OpenAPI `Config.reference` /
// `Config.references` `additionalProperties` anyOf for one raw entry: a plain
// string, [ConfigV2ReferenceGit] (required `[repository]`) or
// [ConfigV2ReferenceLocal] (required `[path]`).
//
// Do not replace this with `apijson.UnmarshalRoot(&ConfigV2ReferenceUnion)`. The
// two object variants share no discriminator, and the union decoder's exactness
// heuristic penalises unknown extra properties but never a missing `required`
// field -- so `{"path": "./docs", "zz": 1}` ties on both variants and the
// left-to-right tie-break misroutes it to [ConfigV2ReferenceGit].
// [ConfigLsp.UnmarshalJSON] routes by hand for the same reason.
//
// `repository` is tested first: it is the git variant's only required property and
// `additionalProperties: false` forbids it on the local one, and vice versa. An
// explicit `null` counts as absent; unknown properties are ignored so a property
// added by a newer server can never flip the variant.
func configV2Reference(node gjson.Result) (any, bool) {
	switch {
	case node.Type == gjson.String:
		return ConfigV2ReferenceString(node.String()), true
	case node.Type != gjson.JSON || !node.IsObject():
		return nil, false
	}
	if repository := node.Get("repository"); repository.Exists() && repository.Type != gjson.Null {
		var git ConfigV2ReferenceGit
		if err := apijson.UnmarshalRoot([]byte(node.Raw), &git); err != nil {
			return nil, false
		}
		return git, true
	}
	var local ConfigV2ReferenceLocal
	if err := apijson.UnmarshalRoot([]byte(node.Raw), &local); err != nil {
		return nil, false
	}
	return local, true
}

func (r configJSON) RawJSON() string {
	return r.raw
}

// AsPermission returns the permission field as a typed union.
//
// Possible runtime types of the union are [ConfigPermissionAction] (a short
// string: "ask"|"allow"|"deny") or [ConfigPermission].
func (r *Config) AsPermission() ConfigPermissionUnion {
	return r.Permission
}

// AsFormatter returns the object arm of the `formatter` field as a typed map of
// formatter name to override.
//
// It is nil when `formatter` is absent or carries the boolean arm of the OpenAPI
// `anyOf [boolean, object]`; use [Config.Formatter] directly to read that arm.
func (r *Config) AsFormatter() map[string]ConfigFormatter {
	formatter, _ := r.Formatter.(map[string]ConfigFormatter)
	return formatter
}

// AsLsp returns the object arm of the `lsp` field as a typed map of LSP server name
// to configuration. Each [ConfigLsp] carries its own variant, reachable through
// [ConfigLsp.AsUnion].
//
// It is nil when `lsp` is absent or carries the boolean arm of the OpenAPI
// `anyOf [boolean, object]`; use [Config.Lsp] directly to read that arm.
func (r *Config) AsLsp() map[string]ConfigLsp {
	lsp, _ := r.Lsp.(map[string]ConfigLsp)
	return lsp
}

// AsReference returns the `reference` field with every entry as a typed
// [ConfigV2ReferenceUnion].
//
// Entries that matched no variant of the OpenAPI anyOf are omitted; read
// [Config.Reference] directly to reach them.
func (r *Config) AsReference() map[string]ConfigV2ReferenceUnion {
	return configReferenceUnions(r.Reference)
}

// AsReferences returns the `references` field with every entry as a typed
// [ConfigV2ReferenceUnion].
//
// Entries that matched no variant of the OpenAPI anyOf are omitted; read
// [Config.References] directly to reach them.
func (r *Config) AsReferences() map[string]ConfigV2ReferenceUnion {
	return configReferenceUnions(r.References)
}

// configReferenceUnions narrows an already-routed reference map (see
// [configReferenceField]) to the typed union. It returns nil for an absent field so
// that callers can distinguish it from a present-but-empty object.
func configReferenceUnions(entries map[string]any) map[string]ConfigV2ReferenceUnion {
	if entries == nil {
		return nil
	}
	unions := make(map[string]ConfigV2ReferenceUnion, len(entries))
	for name, value := range entries {
		if variant, ok := value.(ConfigV2ReferenceUnion); ok {
			unions[name] = variant
		}
	}
	return unions
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
	//
	// Permission is the OpenAPI `PermissionConfig` anyOf; the decoder selects the
	// concrete variant structurally, so a nil value means the field was absent, null
	// or matched no variant.
	Permission  ConfigPermissionUnion `json:"permission"`
	Prompt      string                `json:"prompt"`
	Temperature float64               `json:"temperature"`
	Tools       map[string]bool       `json:"tools"`
	TopP        float64               `json:"top_p"`
	Variant     string                `json:"variant"`
	Hidden      bool                  `json:"hidden"`
	Options     map[string]any        `json:"options"`
	Color       string                `json:"color"`
	Steps       int64                 `json:"steps"`
	MaxSteps    int64                 `json:"maxSteps"`
	ExtraFields map[string]any        `json:"-,extras"`
	JSON        agentConfigJSON       `json:"-"`
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
	return apijson.UnmarshalRoot(data, r)
}

func (r agentConfigJSON) RawJSON() string {
	return r.raw
}

// AsPermission returns the permission field as a typed union.
//
// Possible runtime types of the union are [ConfigPermissionAction] (a short
// string: "ask"|"allow"|"deny") or [ConfigPermission].
func (r *AgentConfig) AsPermission() ConfigPermissionUnion {
	return r.Permission
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

// ConfigLsp is the response-side carrier for one entry of the OpenAPI
// `Config.lsp` object form, i.e. `Config.lsp.additionalProperties`, which is
// `anyOf [{disabled: boolean(enum: true)}, {command, extensions?, disabled?, env?,
// initialization?}]` (JS SDK v2: `{disabled: true} | {command: Array<string>, ...}`).
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
	// Pick the variant from the payload shape instead of delegating to the exactness
	// heuristic behind `apijson.UnmarshalRoot(data, &r.union)` -- see the comment on
	// [ConfigLspUnion] for why that heuristic cannot decide this particular union.
	// Decoding straight into the chosen variant keeps the rest of the contract
	// identical: [apijson.Port] still transfers the typed field values plus
	// `JSON.raw` and `JSON.ExtraFields` onto the carrier.
	if configLspIsDisabledVariant(data) {
		var disabled ConfigLspDisabled
		if err = apijson.UnmarshalRoot(data, &disabled); err != nil {
			return err
		}
		r.union = disabled
	} else {
		var object ConfigLspObject
		if err = apijson.UnmarshalRoot(data, &object); err != nil {
			return err
		}
		r.union = object
	}
	return apijson.Port(r.union, &r)
}

// configLspIsDisabledVariant reports whether a raw `Config.lsp` entry is the
// OpenAPI `{disabled: boolean(enum: true)}` variant rather than the general LSP
// server object.
//
// The two anyOf variants share no discriminator key, so the decision is derived
// from their `required` sets instead: the disabled form requires `disabled` and
// pins it to `true` via `enum: [true]`, while the object form requires `command`.
// A payload carrying `command` can therefore never be the disabled form, and a
// payload whose `disabled` is absent or not literally `true` can never be it
// either -- selecting [ConfigLspDisabled] there would produce a
// [ConfigLspDisabledDisabled] value that fails its own [IsKnown] contract.
//
// Unknown properties are deliberately ignored rather than treated as a mismatch
// (both variants declare `additionalProperties: false`), so that a property added
// by a newer server can never flip the variant. An explicit `null` counts as absent,
// matching how [ConfigMcp.UnmarshalJSON] treats a null `oauth`.
func configLspIsDisabledVariant(data []byte) bool {
	if command := gjson.GetBytes(data, "command"); command.Exists() && command.Type != gjson.Null {
		return false
	}
	return gjson.GetBytes(data, "disabled").Type == gjson.True
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

// The variants are registered so that [ConfigLspUnion] participates in apijson's
// union machinery like every other union, but [ConfigLsp.UnmarshalJSON] routes the
// variant itself and never relies on the registration order below.
//
// Regression: the exactness heuristic cannot decide this union. It only penalises
// unknown extra properties and never penalises a missing `required` field, so
// `{"command": ["gopls"], "zz": 1}` scored `extras` on both variants and the
// left-to-right tie-break handed a genuine LSP server config to
// [ConfigLspDisabled]. Neither of the two generic escapes works here:
//   - A discriminator is unavailable. OpenAPI gives the variants no shared
//     judgement key: one is keyed on `disabled` (`enum: [true]`), the other on
//     `command`.
//   - Reordering only moves the bug. With [ConfigLspObject] first, `{"disabled":
//     true}` scores `exact` on it (a missing `required` `command` costs nothing)
//     and short-circuits, making [ConfigLspDisabled] entirely unreachable.
//
// Making [ConfigLspDisabled.UnmarshalJSON] reject foreign payloads would not help
// either: apijson skips a registered variant's own [json.Unmarshaler] and always
// uses the struct decoder for it.
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

// ConfigLspDisabled is the `{disabled: boolean(enum: true)}` variant of the
// OpenAPI `Config.lsp.additionalProperties` anyOf, i.e. the form that switches a
// single LSP server off without configuring it.
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

// ConfigLspObject is the LSP server variant of the OpenAPI
// `Config.lsp.additionalProperties` anyOf. OpenAPI marks `command` as its only
// required property; it is also the variant that carries every payload the
// disabled form cannot represent, including `{"disabled": false}`.
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
	// Command and arguments to run the MCP server (for "local" type).
	// This field can have the runtime type of [[]string].
	Command any `json:"command"`
	// Working directory for the MCP server process (for "local" type).
	// This field can have the runtime type of [string].
	Cwd any `json:"cwd"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// Environment variables to set when running the MCP server (for "local" type).
	// This field can have the runtime type of [map[string]string].
	Environment any `json:"environment"`
	// Headers to send with the request (for "remote" type).
	// This field can have the runtime type of [map[string]string].
	Headers any `json:"headers"`
	// OAuth authentication configuration for the MCP server (for "remote" type).
	// Set to false to disable OAuth auto-detection.
	// This field can have the runtime type of [McpOAuthConfig], [McpOAuthDisabled].
	//
	// Unlike [McpRemoteConfig.OAuth] this stays `any`: it is a union carrier filled by
	// [apijson.Port], which panics when the destination field is a union interface.
	// See [ConfigMcp.UnmarshalJSON].
	OAuth any `json:"oauth"`
	// Timeout in milliseconds for MCP server requests.
	// This field can have the runtime type of [int64].
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

// UnmarshalJSON decodes the OpenAPI `Config.mcp.additionalProperties` anyOf.
//
// The nested `oauth` anyOf on [McpRemoteConfig] needs no routing here: that field
// is declared as [McpOAuthUnion], so the variant decode already produced the typed
// value and [apijson.Port] transfers it onto the `any` carrier below. Porting in
// that direction (union value -> `any` field) is the safe one; the reverse
// (`any` value -> union-typed field) panics in [apijson.Port], which is why
// [ConfigMcp.OAuth] must stay `any`.
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

// AsOAuth returns the oauth field as a typed union.
//
// Possible runtime types of the union are [McpOAuthConfig] or [McpOAuthDisabled]
// (the scalar `false` value that disables OAuth auto-detection). It is nil when
// the MCP server is not a remote one or when `oauth` is absent.
func (r ConfigMcp) AsOAuth() McpOAuthUnion {
	oauth, _ := r.OAuth.(McpOAuthUnion)
	return oauth
}

// Union satisfied by [McpLocalConfig], [McpRemoteConfig] or [ConfigMcpDisabled].
type ConfigMcpUnion interface {
	implementsConfigMcp()
}

// The union is discriminated on `type`, which OpenAPI pins to `enum: ["local"]`
// on McpLocalConfig and `enum: ["remote"]` on McpRemoteConfig. The disabled form
// `{enabled: boolean}` declares no `type` at all, so it is matched by the absent
// (nil) discriminator value.
//
// Regression: without a discriminator the exactness heuristic decided the variant,
// and `{"enabled": false}` resolved to [McpLocalConfig] — the struct decoder never
// penalises missing `required` fields, so `enabled` alone scored as an exact match
// on the left-most variant and [ConfigMcpDisabled] was unreachable. Reordering the
// variants cannot fix this: ties are broken left-to-right, so any object carrying
// an unknown property would then score `extras` on [ConfigMcpDisabled] first and
// swallow genuine local/remote configs.
//
// DiscriminatorValue must stay an untyped string constant ("local"/"remote"):
// the decoder compares it against the `any` gjson value with `==`, so a typed
// enum constant such as [McpLocalConfigTypeLocal] would never match.
func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigMcpUnion](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "local",
			Type:               reflect.TypeFor[McpLocalConfig](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "remote",
			Type:               reflect.TypeFor[McpRemoteConfig](),
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

// ConfigPermission is the object arm of the OpenAPI `PermissionConfig` anyOf.
//
// Ten of its properties -- `read`, `edit`, `glob`, `grep`, `list`, `bash`, `task`,
// `external_directory`, `lsp` and `skill` -- are declared `$ref
// PermissionRuleConfig`, i.e. the identical `anyOf [PermissionActionConfig,
// PermissionObjectConfig]`. They are therefore modelled identically, as
// [ConfigPermissionBashUnion] whose two variants [ConfigPermissionBashString] and
// [ConfigPermissionBashMap] are exactly that anyOf. The decoder selects the
// concrete variant structurally, so a nil rule means the property was absent, null
// or matched no variant. Use [ConfigPermission.AsBash] and its siblings for a read
// that mirrors the other `AsXxx` accessors in this package.
//
// The remaining five -- `todowrite`, `question`, `webfetch`, `websearch` and
// `doom_loop` -- are declared `$ref PermissionActionConfig`, a plain string enum,
// and so are typed directly.
type ConfigPermission struct {
	Bash              ConfigPermissionBashUnion `json:"bash"`
	Edit              ConfigPermissionBashUnion `json:"edit"`
	Webfetch          ConfigPermissionWebfetch  `json:"webfetch"`
	Read              ConfigPermissionBashUnion `json:"read"`
	Glob              ConfigPermissionBashUnion `json:"glob"`
	Grep              ConfigPermissionBashUnion `json:"grep"`
	List              ConfigPermissionBashUnion `json:"list"`
	Task              ConfigPermissionBashUnion `json:"task"`
	ExternalDirectory ConfigPermissionBashUnion `json:"external_directory"`
	Todowrite         ConfigPermissionTodowrite `json:"todowrite"`
	Question          ConfigPermissionQuestion  `json:"question"`
	Websearch         ConfigPermissionWebsearch `json:"websearch"`
	Lsp               ConfigPermissionBashUnion `json:"lsp"`
	DoomLoop          ConfigPermissionDoomLoop  `json:"doom_loop"`
	Skill             ConfigPermissionBashUnion `json:"skill"`
	// ExtraFields carries the properties allowed by the OpenAPI
	// `additionalProperties: $ref PermissionRuleConfig` of this object, i.e.
	// permission rules for tools this SDK does not know yet.
	ExtraFields map[string]ConfigPermissionBashUnion `json:"-,extras"`
	JSON        configPermissionJSON                 `json:"-"`
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

// AsBash returns the `bash` permission rule as a typed union.
//
// Possible runtime types of the union are [ConfigPermissionBashString] (a short
// action: "ask"|"allow"|"deny") or [ConfigPermissionBashMap] (per-pattern actions).
// It is nil when the rule is absent.
func (r ConfigPermission) AsBash() ConfigPermissionBashUnion {
	return r.Bash
}

// AsEdit returns the `edit` permission rule as a typed union. See
// [ConfigPermission.AsBash] for the possible runtime types.
func (r ConfigPermission) AsEdit() ConfigPermissionBashUnion {
	return r.Edit
}

// AsRead returns the `read` permission rule as a typed union. See
// [ConfigPermission.AsBash] for the possible runtime types.
func (r ConfigPermission) AsRead() ConfigPermissionBashUnion {
	return r.Read
}

// AsGlob returns the `glob` permission rule as a typed union. See
// [ConfigPermission.AsBash] for the possible runtime types.
func (r ConfigPermission) AsGlob() ConfigPermissionBashUnion {
	return r.Glob
}

// AsGrep returns the `grep` permission rule as a typed union. See
// [ConfigPermission.AsBash] for the possible runtime types.
func (r ConfigPermission) AsGrep() ConfigPermissionBashUnion {
	return r.Grep
}

// AsList returns the `list` permission rule as a typed union. See
// [ConfigPermission.AsBash] for the possible runtime types.
func (r ConfigPermission) AsList() ConfigPermissionBashUnion {
	return r.List
}

// AsTask returns the `task` permission rule as a typed union. See
// [ConfigPermission.AsBash] for the possible runtime types.
func (r ConfigPermission) AsTask() ConfigPermissionBashUnion {
	return r.Task
}

// AsExternalDirectory returns the `external_directory` permission rule as a typed
// union. See [ConfigPermission.AsBash] for the possible runtime types.
func (r ConfigPermission) AsExternalDirectory() ConfigPermissionBashUnion {
	return r.ExternalDirectory
}

// AsLsp returns the `lsp` permission rule as a typed union. See
// [ConfigPermission.AsBash] for the possible runtime types.
func (r ConfigPermission) AsLsp() ConfigPermissionBashUnion {
	return r.Lsp
}

// AsSkill returns the `skill` permission rule as a typed union. See
// [ConfigPermission.AsBash] for the possible runtime types.
func (r ConfigPermission) AsSkill() ConfigPermissionBashUnion {
	return r.Skill
}

// ConfigPermissionBashUnion models the OpenAPI `PermissionRuleConfig` anyOf
// (`anyOf [PermissionActionConfig, PermissionObjectConfig]`, JS SDK v2
// `PermissionRuleConfig = PermissionActionConfig | PermissionObjectConfig`).
//
// Despite the name it is not specific to the `bash` rule: OpenAPI declares all ten
// rule properties of [ConfigPermission] -- and its `additionalProperties` -- with
// the same `$ref PermissionRuleConfig`, so they all route through this union.
//
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
	// This field can have the runtime type of [bool], [map[string]any].
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
	// This field can have the runtime type of [map[string]any].
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
	//
	// Timeout is the OpenAPI `ProviderConfig.options.timeout` anyOf; the decoder
	// selects the concrete variant structurally, so a nil value means the field was
	// absent, null or matched no variant.
	Timeout ConfigProviderOptionsTimeoutUnion `json:"timeout"`
	// Timeout in milliseconds to wait for response headers. Provider integrations
	// may set defaults. Set to false to disable timeout.
	//
	// HeaderTimeout is the OpenAPI `ProviderConfig.options.headerTimeout` anyOf and
	// resolves exactly like [ConfigProviderOptions.Timeout].
	HeaderTimeout ConfigProviderOptionsTimeoutUnion `json:"headerTimeout"`
	ChunkTimeout  int64                             `json:"chunkTimeout"`
	ExtraFields   map[string]any                    `json:"-,extras"`
	JSON          configProviderOptionsJSON         `json:"-"`
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

// UnmarshalJSON decodes [ConfigProviderOptions].
//
// `timeout` and `headerTimeout` are both `anyOf [integer, boolean(enum: false)]`
// and are declared as [ConfigProviderOptionsTimeoutUnion], so the decoder resolves
// them through the registered union: the numeric variant lands on [shared.UnionInt]
// (an int64, as OpenAPI's `integer` requires) and the scalar `false` on
// [shared.UnionBool]. Declaring them `any` instead would send them through the
// generic interface branch, which hands back `float64` for every JSON number.
func (r *ConfigProviderOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configProviderOptionsJSON) RawJSON() string {
	return r.raw
}

// AsTimeout returns the timeout field as a typed union.
//
// Possible runtime types of the union are [shared.UnionInt] (milliseconds, > 0),
// [shared.UnionFloat] (a non-integral number, which OpenAPI does not allow) or
// [shared.UnionBool] (the scalar `false` value that disables the timeout). It is
// nil when `timeout` is absent.
func (r ConfigProviderOptions) AsTimeout() ConfigProviderOptionsTimeoutUnion {
	return r.Timeout
}

// AsHeaderTimeout returns the headerTimeout field as a typed union.
//
// Possible runtime types of the union are [shared.UnionInt] (milliseconds, > 0),
// [shared.UnionFloat] (a non-integral number, which OpenAPI does not allow) or
// [shared.UnionBool] (the scalar `false` value that disables the timeout). It is
// nil when `headerTimeout` is absent.
func (r ConfigProviderOptions) AsHeaderTimeout() ConfigProviderOptionsTimeoutUnion {
	return r.HeaderTimeout
}

// ConfigProviderOptionsTimeoutUnion represents a timeout duration as either an
// integer (milliseconds, > 0) or false (to disable). Used by
// [ConfigProviderOptions.Timeout] and [ConfigProviderOptions.HeaderTimeout].
//
// Union satisfied by [shared.UnionInt], [shared.UnionFloat] or [shared.UnionBool].
//
// OpenAPI constrains the numeric variant to `integer` and the boolean one to
// `enum: [false]`. [shared.UnionFloat] and the `gjson.True` filter are registered
// beyond that so a spec-violating payload still decodes losslessly instead of
// failing the whole document: a fractional number would otherwise be truncated
// into [shared.UnionInt], and `true` would match no variant at all and make the
// union decoder fail. [shared.UnionInt] is registered first, so a whole number
// still short-circuits onto it as an exact match.
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
			TypeFilter: gjson.Number,
			Type:       reflect.TypeFor[shared.UnionFloat](),
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
	// Set to false to disable OAuth auto-detection.
	//
	// OAuth is the OpenAPI `McpRemoteConfig.oauth` anyOf; the decoder selects the
	// concrete variant structurally, so a nil value means the field was absent, null
	// or matched no variant.
	OAuth McpOAuthUnion       `json:"oauth"`
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

// AsOAuth returns the oauth field as a typed union.
//
// Possible runtime types of the union are [McpOAuthConfig] or [McpOAuthDisabled]
// (the scalar `false` value that disables OAuth auto-detection). It is nil when
// `oauth` is absent.
func (r McpRemoteConfig) AsOAuth() McpOAuthUnion {
	return r.OAuth
}

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

// McpOAuthUnion represents the OpenAPI `McpRemoteConfig.oauth` anyOf union:
// `anyOf [McpOAuthConfig, boolean(enum: false)]`.
//
// Union satisfied by [McpOAuthConfig] (a complete OAuth configuration) or
// [McpOAuthDisabled] (the scalar `false` value that disables OAuth
// auto-detection).
type McpOAuthUnion interface {
	implementsMcpOAuthUnion()
}

// McpOAuthDisabled is the scalar variant of the OpenAPI `McpRemoteConfig.oauth`
// anyOf, i.e. `false` to disable OAuth auto-detection for the MCP server.
//
// OpenAPI constrains this variant to `enum: [false]`, which is enforced by
// [McpOAuthDisabled.IsKnown] rather than by rejecting the payload, so a
// spec-violating `"oauth": true` still decodes instead of failing the whole
// document.
type McpOAuthDisabled bool

const (
	McpOAuthDisabledFalse McpOAuthDisabled = false
)

func (r McpOAuthDisabled) IsKnown() bool {
	switch r {
	case McpOAuthDisabledFalse:
		return true
	}
	return false
}

func (r McpOAuthDisabled) implementsMcpOAuthUnion() {}

func (r McpOAuthConfig) implementsMcpOAuthUnion() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[McpOAuthUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[McpOAuthConfig](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[McpOAuthDisabled](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeFor[McpOAuthDisabled](),
		},
	)
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
	// Enable or configure formatters. Pass [shared.UnionBool](false) to disable,
	// [shared.UnionBool](true) to enable built-ins, or a [ConfigFormatterMapParam]
	// of formatter-name to [ConfigFormatterParam] to enable with overrides.
	Formatter    param.Field[ConfigFormatterSettingUnionParam] `json:"formatter"`
	Instructions param.Field[[]string]                         `json:"instructions"`
	Layout       param.Field[ConfigLayout]                     `json:"layout"`
	LogLevel     param.Field[ConfigLogLevel]                   `json:"logLevel"`
	// Enable or configure LSP servers. Pass [shared.UnionBool](false) to disable,
	// [shared.UnionBool](true) to enable built-ins, or a [ConfigLspMapParam] of
	// LSP-name to [ConfigLspUnionParam] to enable with overrides.
	Lsp param.Field[ConfigLspSettingUnionParam] `json:"lsp"`
	// Map of MCP server name → configuration. Each value is a
	// [ConfigMcpLocalParam], a [ConfigMcpRemoteParam] or a [ConfigMcpDisabledParam].
	Mcp   param.Field[map[string]ConfigMcpUnionParam] `json:"mcp"`
	Mode  param.Field[ConfigModeParam]                `json:"mode"`
	Model param.Field[string]                         `json:"model"`
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
	// ExtraFields carries the properties allowed by the OpenAPI
	// `AgentConfig.additionalProperties: {}`, which places no constraint on the value.
	ExtraFields map[string]any `json:"-,extras"`
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

// ConfigMcpUnionParam is the request-side union for one entry of the `mcp` config
// map, i.e. the OpenAPI `Config.mcp.additionalProperties` anyOf
// `[McpLocalConfig, McpRemoteConfig, {enabled: boolean}]` (JS SDK v2:
// `McpLocalConfig | McpRemoteConfig | { enabled: boolean }`).
//
// Satisfied by [ConfigMcpLocalParam], [ConfigMcpRemoteParam],
// [ConfigMcpDisabledParam].
//
// It replaces the former flat superset `ConfigMcpParam`, whose single struct could
// express combinations OpenAPI forbids (both variants declare
// `additionalProperties: false`) and whose `command` / `cwd` / `environment` /
// `headers` / `oauth` / `timeout` fields were typed `param.Field[any]`, losing the
// `integer` and `map<string, string>` contracts. The same OpenAPI schema is modelled
// this way for `POST /mcp` in [McpAddParamsConfigUnion].
type ConfigMcpUnionParam interface {
	implementsConfigMcpUnionParam()
}

// The union is discriminated on `type`, mirroring the response-side
// [ConfigMcpUnion]: OpenAPI pins it to `enum: ["local"]` on McpLocalConfig and
// `enum: ["remote"]` on McpRemoteConfig, while the disabled form `{enabled:
// boolean}` declares no `type` at all and so is matched by the absent (nil)
// discriminator value.
//
// DiscriminatorValue must stay an untyped string constant ("local"/"remote"): the
// decoder compares it against the `any` gjson value with `==`, so a typed enum
// constant such as [McpLocalConfigTypeLocal] would never match.
func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigMcpUnionParam](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "local",
			Type:               reflect.TypeFor[ConfigMcpLocalParam](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "remote",
			Type:               reflect.TypeFor[ConfigMcpRemoteParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigMcpDisabledParam](),
		},
	)
}

// ConfigMcpLocalParam is the request-side representation of the OpenAPI
// [McpLocalConfig] variant of [ConfigMcpUnionParam].
type ConfigMcpLocalParam struct {
	// Type of MCP server connection
	Type param.Field[McpLocalConfigType] `json:"type,required"`
	// Command and arguments to run the MCP server
	Command param.Field[[]string] `json:"command,required"`
	// Working directory for the MCP server process
	Cwd param.Field[string] `json:"cwd"`
	// Enable or disable the MCP server on startup
	Enabled param.Field[bool] `json:"enabled"`
	// Environment variables to set when running the MCP server
	Environment param.Field[map[string]string] `json:"environment"`
	// Timeout in milliseconds for MCP server requests
	Timeout param.Field[int64] `json:"timeout"`
}

func (r ConfigMcpLocalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigMcpLocalParam) implementsConfigMcpUnionParam() {}

// ConfigMcpRemoteParam is the request-side representation of the OpenAPI
// [McpRemoteConfig] variant of [ConfigMcpUnionParam].
type ConfigMcpRemoteParam struct {
	// Type of MCP server connection
	Type param.Field[McpRemoteConfigType] `json:"type,required"`
	// URL of the remote MCP server
	URL param.Field[string] `json:"url,required"`
	// Enable or disable the MCP server on startup
	Enabled param.Field[bool] `json:"enabled"`
	// Headers to send with the request
	Headers param.Field[map[string]string] `json:"headers"`
	// OAuth authentication configuration for the MCP server. Pass
	// [shared.UnionBool](false) to disable OAuth auto-detection.
	OAuth param.Field[ConfigMcpOAuthUnionParam] `json:"oauth"`
	// Timeout in milliseconds for MCP server requests
	Timeout param.Field[int64] `json:"timeout"`
}

func (r ConfigMcpRemoteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigMcpRemoteParam) implementsConfigMcpUnionParam() {}

// ConfigMcpDisabledParam is the request-side representation of the
// `{enabled: boolean}` variant of [ConfigMcpUnionParam], i.e. the form that
// switches an MCP server off without configuring it. OpenAPI marks `enabled` as its
// only (and required) property.
type ConfigMcpDisabledParam struct {
	// Enable or disable the MCP server on startup
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ConfigMcpDisabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigMcpDisabledParam) implementsConfigMcpUnionParam() {}

// ConfigMcpOAuthUnionParam is the request-side union for the `oauth` property of
// [ConfigMcpRemoteParam], i.e. the OpenAPI `McpRemoteConfig.oauth` anyOf
// `[McpOAuthConfig, boolean(enum: false)]` (JS SDK v2: `McpOAuthConfig | false`).
//
// Satisfied by [ConfigMcpOAuthParam] or [shared.UnionBool] (the scalar `false`
// value that disables OAuth auto-detection).
type ConfigMcpOAuthUnionParam interface {
	ImplementsConfigMcpOAuthUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigMcpOAuthUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigMcpOAuthParam](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[shared.UnionBool](),
		},
	)
}

// ConfigMcpOAuthParam is the request-side representation of the OpenAPI
// [McpOAuthConfig] schema, the object arm of [ConfigMcpOAuthUnionParam].
type ConfigMcpOAuthParam struct {
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

func (r ConfigMcpOAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigMcpOAuthParam) ImplementsConfigMcpOAuthUnionParam() {}

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
	// Timeout in milliseconds for full requests to this provider. Pass
	// [shared.UnionInt] for a duration or [shared.UnionBool](false) to disable it.
	Timeout param.Field[ConfigProviderOptionsTimeoutUnion] `json:"timeout"`
	// Timeout in milliseconds to wait for response headers. Pass [shared.UnionInt]
	// for a duration or [shared.UnionBool](false) to disable it.
	HeaderTimeout param.Field[ConfigProviderOptionsTimeoutUnion] `json:"headerTimeout"`
	ChunkTimeout  param.Field[int64]                             `json:"chunkTimeout"`
	// ExtraFields carries the properties allowed by the OpenAPI
	// `ProviderConfig.options.additionalProperties: {}`, which places no constraint
	// on the value.
	ExtraFields map[string]any `json:"-,extras"`
}

func (r ConfigProviderOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigFormatterSettingUnionParam is the request-side union for the whole
// `formatter` config field, i.e. the OpenAPI `Config.formatter` anyOf
// `[boolean, object(additionalProperties: <formatter shape>)]` (JS SDK v2:
// `boolean | { [key: string]: {...} }`). Per its OpenAPI description: omit or pass
// false to disable, true to enable built-ins, or an object to enable built-ins with
// overrides.
//
// Satisfied by [shared.UnionBool] or [ConfigFormatterMapParam].
type ConfigFormatterSettingUnionParam interface {
	ImplementsConfigFormatterSettingUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigFormatterSettingUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeFor[shared.UnionBool](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[shared.UnionBool](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigFormatterMapParam](),
		},
	)
}

// ConfigFormatterMapParam is the object arm of [ConfigFormatterSettingUnionParam]:
// a map of formatter name to override.
type ConfigFormatterMapParam map[string]ConfigFormatterParam

func (r ConfigFormatterMapParam) ImplementsConfigFormatterSettingUnionParam() {}

// ConfigFormatterParam is the request-side representation of [ConfigFormatter], one
// entry of the object arm of the OpenAPI `Config.formatter` anyOf.
//
// It exists because [ConfigFormatter] is a response type: its fields are plain Go
// values with no `param.Field` wrapper, so marshalling one into a request body emits
// every zero value (`"disabled": false`, `"environment": {}`, `"extensions": []`)
// as if the caller had asked for it.
type ConfigFormatterParam struct {
	Command     param.Field[[]string]          `json:"command"`
	Disabled    param.Field[bool]              `json:"disabled"`
	Environment param.Field[map[string]string] `json:"environment"`
	Extensions  param.Field[[]string]          `json:"extensions"`
}

func (r ConfigFormatterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConfigLspSettingUnionParam is the request-side union for the whole `lsp` config
// field, i.e. the OpenAPI `Config.lsp` anyOf
// `[boolean, object(additionalProperties: <the ConfigLsp union>)]` (JS SDK v2:
// `boolean | { [key: string]: ... }`). Per its OpenAPI description: omit or pass
// false to disable, true to enable built-ins, or an object to enable built-ins with
// overrides.
//
// Satisfied by [shared.UnionBool] or [ConfigLspMapParam].
type ConfigLspSettingUnionParam interface {
	ImplementsConfigLspSettingUnionParam()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigLspSettingUnionParam](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeFor[shared.UnionBool](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeFor[shared.UnionBool](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[ConfigLspMapParam](),
		},
	)
}

// ConfigLspMapParam is the object arm of [ConfigLspSettingUnionParam]: a map of LSP
// server name to configuration, each entry being a [ConfigLspUnionParam].
type ConfigLspMapParam map[string]ConfigLspUnionParam

func (r ConfigLspMapParam) ImplementsConfigLspSettingUnionParam() {}

// ConfigLspUnionParam is the request-side union for one entry of the `lsp` config
// map, i.e. the OpenAPI `Config.lsp.additionalProperties` anyOf
// `[{disabled: boolean(enum: true)}, {command, extensions?, disabled?, env?,
// initialization?}]` (JS SDK v2:
// `{disabled: true} | {command: Array<string>, ...}`). It mirrors the response-side
// [ConfigLspUnion].
//
// Satisfied by [ConfigLspDisabledParam] or [ConfigLspObjectParam].
type ConfigLspUnionParam interface {
	implementsConfigLspUnionParam()
}

// The variants share no discriminator key -- OpenAPI keys one on `disabled`
// (`enum: [true]`) and the other on `command` -- so no discriminator is registered,
// exactly as for the response-side [ConfigLspUnion]. Requests are only ever
// marshalled, so the caller's choice of variant is what reaches the wire and the
// decoder's exactness heuristic is never consulted here.
func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ConfigLspUnionParam](),
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

// ConfigLspDisabledParam is the request-side representation of [ConfigLspDisabled],
// the `{disabled: boolean(enum: true)}` variant of [ConfigLspUnionParam].
type ConfigLspDisabledParam struct {
	Disabled param.Field[ConfigLspDisabledDisabled] `json:"disabled,required"`
}

func (r ConfigLspDisabledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigLspDisabledParam) implementsConfigLspUnionParam() {}

// ConfigLspObjectParam is the request-side representation of [ConfigLspObject], the
// LSP server variant of [ConfigLspUnionParam]. OpenAPI marks `command` as its only
// required property.
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

func (r ConfigLspObjectParam) implementsConfigLspUnionParam() {}

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

// ConfigPermissionParam is the request-side representation of [ConfigPermission].
//
// As on the response side, the ten properties OpenAPI declares `$ref
// PermissionRuleConfig` -- `read`, `edit`, `glob`, `grep`, `list`, `bash`, `task`,
// `external_directory`, `lsp` and `skill` -- are all typed
// [ConfigPermissionBashUnionParam], and the five declared `$ref
// PermissionActionConfig` are plain string enums.
type ConfigPermissionParam struct {
	Bash              param.Field[ConfigPermissionBashUnionParam] `json:"bash"`
	Edit              param.Field[ConfigPermissionBashUnionParam] `json:"edit"`
	Webfetch          param.Field[ConfigPermissionWebfetch]       `json:"webfetch"`
	Read              param.Field[ConfigPermissionBashUnionParam] `json:"read"`
	Glob              param.Field[ConfigPermissionBashUnionParam] `json:"glob"`
	Grep              param.Field[ConfigPermissionBashUnionParam] `json:"grep"`
	List              param.Field[ConfigPermissionBashUnionParam] `json:"list"`
	Task              param.Field[ConfigPermissionBashUnionParam] `json:"task"`
	ExternalDirectory param.Field[ConfigPermissionBashUnionParam] `json:"external_directory"`
	Todowrite         param.Field[ConfigPermissionTodowrite]      `json:"todowrite"`
	Question          param.Field[ConfigPermissionQuestion]       `json:"question"`
	Websearch         param.Field[ConfigPermissionWebsearch]      `json:"websearch"`
	Lsp               param.Field[ConfigPermissionBashUnionParam] `json:"lsp"`
	DoomLoop          param.Field[ConfigPermissionDoomLoop]       `json:"doom_loop"`
	Skill             param.Field[ConfigPermissionBashUnionParam] `json:"skill"`
	// ExtraFields carries the properties allowed by the OpenAPI
	// `additionalProperties: $ref PermissionRuleConfig` of this object, i.e.
	// permission rules for tools this SDK does not know yet.
	ExtraFields map[string]ConfigPermissionBashUnionParam `json:"-,extras"`
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
