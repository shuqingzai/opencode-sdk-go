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
	"github.com/tidwall/gjson"
)

// McpService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMcpService] method instead.
type McpService struct {
	Options []option.RequestOption
	Auth    *McpAuthService
}

// NewMcpService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMcpService(opts ...option.RequestOption) (r *McpService) {
	r = &McpService{}
	r.Options = opts
	r.Auth = NewMcpAuthService(opts...)
	return
}

// McpAuthService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMcpAuthService] method instead.
type McpAuthService struct {
	Options []option.RequestOption
}

// NewMcpAuthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMcpAuthService(opts ...option.RequestOption) (r *McpAuthService) {
	r = &McpAuthService{}
	r.Options = opts
	return
}

// Get MCP server status
func (r *McpService) Status(ctx context.Context, query McpStatusParams, opts ...option.RequestOption) (res *map[string]McpStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "mcp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Add an MCP server
func (r *McpService) Add(ctx context.Context, params McpAddParams, opts ...option.RequestOption) (res *map[string]McpStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "mcp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Connect to an MCP server
func (r *McpService) Connect(ctx context.Context, name string, query McpConnectParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("mcp/%s/connect", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

// Disconnect from an MCP server
func (r *McpService) Disconnect(ctx context.Context, name string, query McpDisconnectParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("mcp/%s/disconnect", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

// Start MCP OAuth authentication flow
func (r *McpAuthService) Start(ctx context.Context, name string, query McpAuthStartParams, opts ...option.RequestOption) (res *McpAuthStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("mcp/%s/auth", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

// Deprecated: Use Auth.Start instead.
func (r *McpService) AuthStart(ctx context.Context, name string, query McpAuthStartParams, opts ...option.RequestOption) (res *McpAuthStartResponse, err error) {
	return r.Auth.Start(ctx, name, query, opts...)
}

// OAuth callback
func (r *McpAuthService) Callback(ctx context.Context, name string, params McpAuthCallbackParams, opts ...option.RequestOption) (res *McpStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("mcp/%s/auth/callback", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Deprecated: Use Auth.Callback instead.
func (r *McpService) AuthCallback(ctx context.Context, name string, params McpAuthCallbackParams, opts ...option.RequestOption) (res *McpStatus, err error) {
	return r.Auth.Callback(ctx, name, params, opts...)
}

// Authenticate with MCP server
func (r *McpAuthService) Authenticate(ctx context.Context, name string, query McpAuthAuthenticateParams, opts ...option.RequestOption) (res *McpStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("mcp/%s/auth/authenticate", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

// Deprecated: Use Auth.Authenticate instead.
func (r *McpService) AuthAuthenticate(ctx context.Context, name string, query McpAuthAuthenticateParams, opts ...option.RequestOption) (res *McpStatus, err error) {
	return r.Auth.Authenticate(ctx, name, query, opts...)
}

// Remove MCP server authentication
func (r *McpAuthService) Remove(ctx context.Context, name string, query McpAuthRemoveParams, opts ...option.RequestOption) (res *McpAuthRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("mcp/%s/auth", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, &res, opts...)
	return
}

// Deprecated: Use Auth.Remove instead.
func (r *McpService) AuthRemove(ctx context.Context, name string, query McpAuthRemoveParams, opts ...option.RequestOption) (res *McpAuthRemoveResponse, err error) {
	return r.Auth.Remove(ctx, name, query, opts...)
}

// McpStatus represents the status of an MCP server connection.
type McpStatus struct {
	Status McpStatusStatus `json:"status,required"`
	Error  string          `json:"error"`
	JSON   mcpStatusJSON   `json:"-"`
	union  McpStatusUnion
}

// mcpStatusJSON contains the JSON metadata for the struct [McpStatus]
type mcpStatusJSON struct {
	Status      apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r mcpStatusJSON) RawJSON() string {
	return r.raw
}

func (r *McpStatus) UnmarshalJSON(data []byte) (err error) {
	*r = McpStatus{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [McpStatusUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [McpStatusConnected], [McpStatusDisabled],
// [McpStatusFailed], [McpStatusNeedsAuth], [McpStatusNeedsClientRegistration].
func (r McpStatus) AsUnion() McpStatusUnion {
	return r.union
}

// Union satisfied by [McpStatusConnected], [McpStatusDisabled], [McpStatusFailed],
// [McpStatusNeedsAuth], or [McpStatusNeedsClientRegistration].
type McpStatusUnion interface {
	implementsMcpStatus()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*McpStatusUnion)(nil)).Elem(),
		"status",
		apijson.UnionVariant{
			Type: reflect.TypeOf(McpStatusConnected{}),
		},
		apijson.UnionVariant{
			Type: reflect.TypeOf(McpStatusDisabled{}),
		},
		apijson.UnionVariant{
			Type: reflect.TypeOf(McpStatusFailed{}),
		},
		apijson.UnionVariant{
			Type: reflect.TypeOf(McpStatusNeedsAuth{}),
		},
		apijson.UnionVariant{
			Type: reflect.TypeOf(McpStatusNeedsClientRegistration{}),
		},
	)
}

// McpStatusConnected indicates the MCP server is connected.
type McpStatusConnected struct {
	Status McpStatusStatus        `json:"status,required"`
	JSON   mcpStatusConnectedJSON `json:"-"`
}

// mcpStatusConnectedJSON contains the JSON metadata for the struct [McpStatusConnected]
type mcpStatusConnectedJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusConnected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpStatusConnectedJSON) RawJSON() string {
	return r.raw
}

func (r McpStatusConnected) implementsMcpStatus() {}

// McpStatusDisabled indicates the MCP server is disabled.
type McpStatusDisabled struct {
	Status McpStatusStatus       `json:"status,required"`
	JSON   mcpStatusDisabledJSON `json:"-"`
}

// mcpStatusDisabledJSON contains the JSON metadata for the struct [McpStatusDisabled]
type mcpStatusDisabledJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusDisabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpStatusDisabledJSON) RawJSON() string {
	return r.raw
}

func (r McpStatusDisabled) implementsMcpStatus() {}

// McpStatusFailed indicates the MCP server connection failed.
type McpStatusFailed struct {
	Error  string              `json:"error,required"`
	Status McpStatusStatus     `json:"status,required"`
	JSON   mcpStatusFailedJSON `json:"-"`
}

// mcpStatusFailedJSON contains the JSON metadata for the struct [McpStatusFailed]
type mcpStatusFailedJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpStatusFailedJSON) RawJSON() string {
	return r.raw
}

func (r McpStatusFailed) implementsMcpStatus() {}

// McpStatusNeedsAuth indicates the MCP server requires authentication.
type McpStatusNeedsAuth struct {
	Status McpStatusStatus        `json:"status,required"`
	JSON   mcpStatusNeedsAuthJSON `json:"-"`
}

// mcpStatusNeedsAuthJSON contains the JSON metadata for the struct [McpStatusNeedsAuth]
type mcpStatusNeedsAuthJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusNeedsAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpStatusNeedsAuthJSON) RawJSON() string {
	return r.raw
}

func (r McpStatusNeedsAuth) implementsMcpStatus() {}

// McpStatusNeedsClientRegistration indicates the MCP server requires client registration.
type McpStatusNeedsClientRegistration struct {
	Error  string                               `json:"error,required"`
	Status McpStatusStatus                      `json:"status,required"`
	JSON   mcpStatusNeedsClientRegistrationJSON `json:"-"`
}

// mcpStatusNeedsClientRegistrationJSON contains the JSON metadata for the struct [McpStatusNeedsClientRegistration]
type mcpStatusNeedsClientRegistrationJSON struct {
	Error       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpStatusNeedsClientRegistration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpStatusNeedsClientRegistrationJSON) RawJSON() string {
	return r.raw
}

func (r McpStatusNeedsClientRegistration) implementsMcpStatus() {}

// McpStatusStatus represents the status of an MCP server connection.
type McpStatusStatus string

const (
	McpStatusStatusConnected               McpStatusStatus = "connected"
	McpStatusStatusDisabled                McpStatusStatus = "disabled"
	McpStatusStatusFailed                  McpStatusStatus = "failed"
	McpStatusStatusNeedsAuth               McpStatusStatus = "needs_auth"
	McpStatusStatusNeedsClientRegistration McpStatusStatus = "needs_client_registration"
)

func (r McpStatusStatus) IsKnown() bool {
	switch r {
	case McpStatusStatusConnected, McpStatusStatusDisabled, McpStatusStatusFailed, McpStatusStatusNeedsAuth, McpStatusStatusNeedsClientRegistration:
		return true
	}
	return false
}

// McpOAuthConfig contains OAuth authentication configuration for an MCP server.
type McpOAuthConfig struct {
	// OAuth client ID. If not provided, dynamic client registration (RFC 7591) will be attempted.
	ClientID string `json:"clientId"`
	// OAuth client secret (if required by the authorization server)
	ClientSecret string `json:"clientSecret"`
	// OAuth scopes to request during authorization
	Scope string `json:"scope"`
	// OAuth callback port for the local HTTP server
	CallbackPort int64 `json:"callbackPort"`
	// OAuth redirect URI
	RedirectURI string             `json:"redirectUri"`
	JSON        mcpOAuthConfigJSON `json:"-"`
}

// mcpOAuthConfigJSON contains the JSON metadata for the struct [McpOAuthConfig]
type mcpOAuthConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	Scope        apijson.Field
	CallbackPort apijson.Field
	RedirectURI  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *McpOAuthConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpOAuthConfigJSON) RawJSON() string {
	return r.raw
}

// McpAddBody represents the body for adding an MCP server.
type McpAddBody struct {
	// Name of the MCP server
	Name string `json:"name,required"`
	// Configuration for the MCP server
	// This field can have the runtime type of [McpAddBodyConfigLocal], [McpAddBodyConfigRemote].
	Config interface{}    `json:"config,required"`
	JSON   mcpAddBodyJSON `json:"-"`
}

// mcpAddBodyJSON contains the JSON metadata for the struct [McpAddBody]
type mcpAddBodyJSON struct {
	Name        apijson.Field
	Config      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpAddBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpAddBodyJSON) RawJSON() string {
	return r.raw
}

// McpAddBodyConfigUnion is a union type for MCP server configuration.
type McpAddBodyConfigUnion interface {
	implementsMcpAddBodyConfigUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*McpAddBodyConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McpAddBodyConfigLocal{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McpAddBodyConfigRemote{}),
		},
	)
}

// McpAddBodyConfigLocal represents local MCP server configuration.
type McpAddBodyConfigLocal struct {
	// Type of MCP server connection
	Type McpLocalConfigType `json:"type,required"`
	// Command and arguments to run the MCP server
	Command []string `json:"command,required"`
	// Cwd is the working directory for the MCP server process.
	Cwd string `json:"cwd"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// Environment variables to set when running the MCP server
	Environment map[string]string `json:"environment"`
	// Timeout in ms for MCP server requests. Defaults to 5000 (5 seconds) if not specified.
	Timeout int64                     `json:"timeout"`
	JSON    mcpAddBodyConfigLocalJSON `json:"-"`
}

// mcpAddBodyConfigLocalJSON contains the JSON metadata for the struct [McpAddBodyConfigLocal]
type mcpAddBodyConfigLocalJSON struct {
	Type        apijson.Field
	Command     apijson.Field
	Cwd         apijson.Field
	Enabled     apijson.Field
	Environment apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpAddBodyConfigLocal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpAddBodyConfigLocalJSON) RawJSON() string {
	return r.raw
}

func (r McpAddBodyConfigLocal) implementsMcpAddBodyConfigUnion() {}

// McpAddBodyConfigRemote represents remote MCP server configuration.
type McpAddBodyConfigRemote struct {
	// Type of MCP server connection
	Type McpRemoteConfigType `json:"type,required"`
	// URL of the remote MCP server
	URL string `json:"url,required"`
	// Enable or disable the MCP server on startup
	Enabled bool `json:"enabled"`
	// Headers to send with the request
	Headers map[string]string `json:"headers"`
	// OAuth authentication configuration for the MCP server. Set to false to disable OAuth auto-detection.
	// This field can have the runtime type of [McpOAuthConfig].
	// When OAuth is disabled, this field is `false` (a JSON boolean).
	OAuth interface{} `json:"oauth"`
	// Timeout in ms for MCP server requests. Defaults to 5000 (5 seconds) if not specified.
	Timeout int64                      `json:"timeout"`
	JSON    mcpAddBodyConfigRemoteJSON `json:"-"`
}

// mcpAddBodyConfigRemoteJSON contains the JSON metadata for the struct [McpAddBodyConfigRemote]
type mcpAddBodyConfigRemoteJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	Enabled     apijson.Field
	Headers     apijson.Field
	OAuth       apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpAddBodyConfigRemote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpAddBodyConfigRemoteJSON) RawJSON() string {
	return r.raw
}

func (r McpAddBodyConfigRemote) implementsMcpAddBodyConfigUnion() {}

// McpAuthCallbackBody represents the body for the OAuth callback.
type McpAuthCallbackBody struct {
	Code string                  `json:"code,required"`
	JSON mcpAuthCallbackBodyJSON `json:"-"`
}

// mcpAuthCallbackBodyJSON contains the JSON metadata for the struct [McpAuthCallbackBody]
type mcpAuthCallbackBodyJSON struct {
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpAuthCallbackBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpAuthCallbackBodyJSON) RawJSON() string {
	return r.raw
}

// McpAuthStartResponse represents the response from starting OAuth authentication.
type McpAuthStartResponse struct {
	AuthorizationURL string                   `json:"authorizationUrl,required"`
	OAuthState       string                   `json:"oauthState,required"`
	JSON             mcpAuthStartResponseJSON `json:"-"`
}

// mcpAuthStartResponseJSON contains the JSON metadata for the struct [McpAuthStartResponse]
type mcpAuthStartResponseJSON struct {
	AuthorizationURL apijson.Field
	OAuthState       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *McpAuthStartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpAuthStartResponseJSON) RawJSON() string {
	return r.raw
}

// McpAuthRemoveResponse represents the response from removing authentication.
type McpAuthRemoveResponse struct {
	Success bool                      `json:"success,required"`
	JSON    mcpAuthRemoveResponseJSON `json:"-"`
}

// mcpAuthRemoveResponseJSON contains the JSON metadata for the struct [McpAuthRemoveResponse]
type mcpAuthRemoveResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpAuthRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpAuthRemoveResponseJSON) RawJSON() string {
	return r.raw
}

// McpStatusParams contains the query parameters for getting MCP status.
type McpStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [McpStatusParams]'s query parameters as `url.Values`.
func (r McpStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// McpAddParams contains the parameters for adding an MCP server.
type McpAddParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Name      param.Field[string] `json:"name,required"`
	// This field can have the runtime type of [McpAddBodyConfigLocal], [McpAddBodyConfigRemote].
	Config interface{} `json:"config,required"`
}

// MarshalJSON serializes [McpAddParams] omitting query parameters.
func (r McpAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [McpAddParams]'s query parameters as `url.Values`.
func (r McpAddParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// McpConnectParams contains the query parameters for connecting to an MCP server.
type McpConnectParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [McpConnectParams]'s query parameters as `url.Values`.
func (r McpConnectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// McpDisconnectParams contains the query parameters for disconnecting from an MCP server.
type McpDisconnectParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [McpDisconnectParams]'s query parameters as `url.Values`.
func (r McpDisconnectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// McpAuthStartParams contains the query parameters for starting OAuth authentication.
type McpAuthStartParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [McpAuthStartParams]'s query parameters as `url.Values`.
func (r McpAuthStartParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// McpAuthCallbackParams contains the parameters for OAuth callback.
type McpAuthCallbackParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Code      param.Field[string] `json:"code,required"`
}

// MarshalJSON serializes [McpAuthCallbackParams] omitting query parameters.
func (r McpAuthCallbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [McpAuthCallbackParams]'s query parameters as `url.Values`.
func (r McpAuthCallbackParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// McpAuthAuthenticateParams contains the query parameters for authenticating with MCP server.
type McpAuthAuthenticateParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [McpAuthAuthenticateParams]'s query parameters as `url.Values`.
func (r McpAuthAuthenticateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// McpAuthRemoveParams contains the query parameters for removing MCP server authentication.
type McpAuthRemoveParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [McpAuthRemoveParams]'s query parameters as `url.Values`.
func (r McpAuthRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
