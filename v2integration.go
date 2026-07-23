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

// V2IntegrationService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2IntegrationService] method instead.
type V2IntegrationService struct {
	Options []option.RequestOption
	Connect *V2IntegrationConnectService
	Attempt *V2IntegrationAttemptService
}

// NewV2IntegrationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2IntegrationService(opts ...option.RequestOption) (r *V2IntegrationService) {
	r = &V2IntegrationService{}
	r.Options = opts
	r.Connect = NewV2IntegrationConnectService(opts...)
	r.Attempt = NewV2IntegrationAttemptService(opts...)
	return
}

// V2IntegrationConnectService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2IntegrationConnectService] method instead.
type V2IntegrationConnectService struct {
	Options []option.RequestOption
}

// NewV2IntegrationConnectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2IntegrationConnectService(opts ...option.RequestOption) (r *V2IntegrationConnectService) {
	r = &V2IntegrationConnectService{}
	r.Options = opts
	return
}

// V2IntegrationAttemptService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2IntegrationAttemptService] method instead.
type V2IntegrationAttemptService struct {
	Options []option.RequestOption
}

// NewV2IntegrationAttemptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2IntegrationAttemptService(opts ...option.RequestOption) (r *V2IntegrationAttemptService) {
	r = &V2IntegrationAttemptService{}
	r.Options = opts
	return
}

// List integrations
//
// Retrieve available integrations and their authentication methods.
func (r *V2IntegrationService) List(ctx context.Context, query V2IntegrationListParams, opts ...option.RequestOption) (res *V2IntegrationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/integration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get integration
//
// Retrieve one integration and its authentication methods.
func (r *V2IntegrationService) Get(ctx context.Context, integrationID string, query V2IntegrationGetParams, opts ...option.RequestOption) (res *V2IntegrationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if integrationID == "" {
		err = errors.New("missing required integrationID parameter")
		return
	}
	path := fmt.Sprintf("api/integration/%s", integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Connect with key
//
// Run a key authentication method and store the resulting credential.
func (r *V2IntegrationConnectService) Key(ctx context.Context, integrationID string, params V2IntegrationConnectKeyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if integrationID == "" {
		err = errors.New("missing required integrationID parameter")
		return
	}
	path := fmt.Sprintf("api/integration/%s/connect/key", integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Begin OAuth connection
//
// Start an OAuth attempt and return the authorization details.
func (r *V2IntegrationConnectService) Oauth(ctx context.Context, integrationID string, params V2IntegrationConnectOauthParams, opts ...option.RequestOption) (res *V2IntegrationConnectOauthResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if integrationID == "" {
		err = errors.New("missing required integrationID parameter")
		return
	}
	path := fmt.Sprintf("api/integration/%s/connect/oauth", integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Cancel OAuth connection
//
// Cancel an OAuth attempt and release its resources.
func (r *V2IntegrationAttemptService) Cancel(ctx context.Context, attemptID string, query V2IntegrationAttemptCancelParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if attemptID == "" {
		err = errors.New("missing required attemptID parameter")
		return
	}
	path := fmt.Sprintf("api/integration/attempt/%s", attemptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, nil, opts...)
	return
}

// Get OAuth attempt status
//
// Poll the current status of an OAuth attempt.
func (r *V2IntegrationAttemptService) Status(ctx context.Context, attemptID string, query V2IntegrationAttemptStatusParams, opts ...option.RequestOption) (res *V2IntegrationAttemptStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if attemptID == "" {
		err = errors.New("missing required attemptID parameter")
		return
	}
	path := fmt.Sprintf("api/integration/attempt/%s", attemptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Complete OAuth connection
//
// Complete a code-based OAuth attempt and store the resulting credential.
func (r *V2IntegrationAttemptService) Complete(ctx context.Context, attemptID string, params V2IntegrationAttemptCompleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if attemptID == "" {
		err = errors.New("missing required attemptID parameter")
		return
	}
	path := fmt.Sprintf("api/integration/attempt/%s/complete", attemptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// ===== Response Types =====

// V2IntegrationListResponse contains the response from the integration list endpoint.
type V2IntegrationListResponse struct {
	Location LocationInfo                  `json:"location,required"`
	Data     []IntegrationInfo             `json:"data,required"`
	JSON     v2IntegrationListResponseJSON `json:"-"`
}

// v2IntegrationListResponseJSON contains the JSON metadata for the struct [V2IntegrationListResponse]
type v2IntegrationListResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2IntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2IntegrationListResponseJSON) RawJSON() string {
	return r.raw
}

// V2IntegrationGetResponse contains the response from the integration get endpoint.
type V2IntegrationGetResponse struct {
	Location LocationInfo                 `json:"location,required"`
	Data     IntegrationInfo              `json:"data,required"`
	JSON     v2IntegrationGetResponseJSON `json:"-"`
}

// v2IntegrationGetResponseJSON contains the JSON metadata for the struct [V2IntegrationGetResponse]
type v2IntegrationGetResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2IntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2IntegrationGetResponseJSON) RawJSON() string {
	return r.raw
}

// V2IntegrationConnectOauthResponse contains the response from the oauth connect endpoint.
type V2IntegrationConnectOauthResponse struct {
	Location LocationInfo                          `json:"location,required"`
	Data     IntegrationAttempt                    `json:"data,required"`
	JSON     v2IntegrationConnectOauthResponseJSON `json:"-"`
}

// v2IntegrationConnectOauthResponseJSON contains the JSON metadata for the struct [V2IntegrationConnectOauthResponse]
type v2IntegrationConnectOauthResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2IntegrationConnectOauthResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2IntegrationConnectOauthResponseJSON) RawJSON() string {
	return r.raw
}

// V2IntegrationAttemptStatusResponse contains the response from the attempt status endpoint.
type V2IntegrationAttemptStatusResponse struct {
	Location LocationInfo                           `json:"location,required"`
	Data     IntegrationAttemptStatus               `json:"data,required"`
	JSON     v2IntegrationAttemptStatusResponseJSON `json:"-"`
}

// v2IntegrationAttemptStatusResponseJSON contains the JSON metadata for the struct [V2IntegrationAttemptStatusResponse]
type v2IntegrationAttemptStatusResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2IntegrationAttemptStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2IntegrationAttemptStatusResponseJSON) RawJSON() string {
	return r.raw
}

// ===== IntegrationInfo =====

// IntegrationInfo represents an integration with its authentication methods and connections.
type IntegrationInfo struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// This field can have the runtime type of []IntegrationOAuthMethod,
	// []IntegrationKeyMethod, []IntegrationEnvMethod.
	Methods interface{} `json:"methods,required"`
	// This field can have the runtime type of []ConnectionCredentialInfo,
	// []ConnectionEnvInfo.
	Connections interface{}         `json:"connections,required"`
	JSON        integrationInfoJSON `json:"-"`
}

// integrationInfoJSON contains the JSON metadata for the struct [IntegrationInfo]
type integrationInfoJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Methods     apijson.Field
	Connections apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationInfoJSON) RawJSON() string {
	return r.raw
}

// ===== IntegrationMethod Union =====

// IntegrationMethodUnion represents the union of integration method types.
// Possible runtime types are [IntegrationOAuthMethod], [IntegrationKeyMethod],
// [IntegrationEnvMethod].
type IntegrationMethodUnion interface {
	implementsIntegrationMethodUnion()
}

// IntegrationOAuthMethod represents an OAuth-based integration method.
type IntegrationOAuthMethod struct {
	ID    string                     `json:"id,required"`
	Type  IntegrationOAuthMethodType `json:"type,required"`
	Label string                     `json:"label,required"`
	// This field can have the runtime type of [[]IntegrationTextPrompt],
	// [[]IntegrationSelectPrompt].
	Prompts interface{}                `json:"prompts"`
	JSON    integrationOAuthMethodJSON `json:"-"`
}

// integrationOAuthMethodJSON contains the JSON metadata for the struct [IntegrationOAuthMethod]
type integrationOAuthMethodJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Label       apijson.Field
	Prompts     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationOAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationOAuthMethodJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationOAuthMethod) implementsIntegrationMethodUnion() {}

type IntegrationOAuthMethodType string

const (
	IntegrationOAuthMethodTypeOAuth IntegrationOAuthMethodType = "oauth"
)

func (r IntegrationOAuthMethodType) IsKnown() bool {
	switch r {
	case IntegrationOAuthMethodTypeOAuth:
		return true
	}
	return false
}

// IntegrationKeyMethod represents a key-based integration method.
type IntegrationKeyMethod struct {
	Type  IntegrationKeyMethodType `json:"type,required"`
	Label string                   `json:"label"`
	JSON  integrationKeyMethodJSON `json:"-"`
}

// integrationKeyMethodJSON contains the JSON metadata for the struct [IntegrationKeyMethod]
type integrationKeyMethodJSON struct {
	Type        apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationKeyMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationKeyMethodJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationKeyMethod) implementsIntegrationMethodUnion() {}

type IntegrationKeyMethodType string

const (
	IntegrationKeyMethodTypeKey IntegrationKeyMethodType = "key"
)

func (r IntegrationKeyMethodType) IsKnown() bool {
	switch r {
	case IntegrationKeyMethodTypeKey:
		return true
	}
	return false
}

// IntegrationEnvMethod represents an environment-variable-based integration method.
type IntegrationEnvMethod struct {
	Type  IntegrationEnvMethodType `json:"type,required"`
	Names []string                 `json:"names,required"`
	JSON  integrationEnvMethodJSON `json:"-"`
}

// integrationEnvMethodJSON contains the JSON metadata for the struct [IntegrationEnvMethod]
type integrationEnvMethodJSON struct {
	Type        apijson.Field
	Names       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationEnvMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationEnvMethodJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationEnvMethod) implementsIntegrationMethodUnion() {}

type IntegrationEnvMethodType string

const (
	IntegrationEnvMethodTypeEnv IntegrationEnvMethodType = "env"
)

func (r IntegrationEnvMethodType) IsKnown() bool {
	switch r {
	case IntegrationEnvMethodTypeEnv:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IntegrationMethodUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationOAuthMethod{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationKeyMethod{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationEnvMethod{}),
		},
	)
}

// ===== ConnectionInfo Union =====

// ConnectionInfoUnion represents the union of connection types.
// Possible runtime types are [ConnectionCredentialInfo], [ConnectionEnvInfo].
type ConnectionInfoUnion interface {
	implementsConnectionInfoUnion()
}

// ConnectionCredentialInfo represents a credential-based connection.
type ConnectionCredentialInfo struct {
	Type  ConnectionCredentialInfoType `json:"type,required"`
	ID    string                       `json:"id,required"`
	Label string                       `json:"label,required"`
	JSON  connectionCredentialInfoJSON `json:"-"`
}

// connectionCredentialInfoJSON contains the JSON metadata for the struct [ConnectionCredentialInfo]
type connectionCredentialInfoJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionCredentialInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionCredentialInfoJSON) RawJSON() string {
	return r.raw
}

func (r ConnectionCredentialInfo) implementsConnectionInfoUnion() {}

type ConnectionCredentialInfoType string

const (
	ConnectionCredentialInfoTypeCredential ConnectionCredentialInfoType = "credential"
)

func (r ConnectionCredentialInfoType) IsKnown() bool {
	switch r {
	case ConnectionCredentialInfoTypeCredential:
		return true
	}
	return false
}

// ConnectionEnvInfo represents an environment-variable-based connection.
type ConnectionEnvInfo struct {
	Type ConnectionEnvInfoType `json:"type,required"`
	Name string                `json:"name,required"`
	JSON connectionEnvInfoJSON `json:"-"`
}

// connectionEnvInfoJSON contains the JSON metadata for the struct [ConnectionEnvInfo]
type connectionEnvInfoJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionEnvInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionEnvInfoJSON) RawJSON() string {
	return r.raw
}

func (r ConnectionEnvInfo) implementsConnectionInfoUnion() {}

type ConnectionEnvInfoType string

const (
	ConnectionEnvInfoTypeEnv ConnectionEnvInfoType = "env"
)

func (r ConnectionEnvInfoType) IsKnown() bool {
	switch r {
	case ConnectionEnvInfoTypeEnv:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectionInfoUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectionCredentialInfo{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectionEnvInfo{}),
		},
	)
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IntegrationPromptUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationTextPrompt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationSelectPrompt{}),
		},
	)
}

// ===== IntegrationAttempt =====

// IntegrationAttempt represents an OAuth integration attempt.
type IntegrationAttempt struct {
	AttemptID    string                 `json:"attemptID,required"`
	URL          string                 `json:"url,required"`
	Instructions string                 `json:"instructions,required"`
	Mode         IntegrationAttemptMode `json:"mode,required"`
	Time         IntegrationAttemptTime `json:"time,required"`
	JSON         integrationAttemptJSON `json:"-"`
}

// integrationAttemptJSON contains the JSON metadata for the struct [IntegrationAttempt]
type integrationAttemptJSON struct {
	AttemptID    apijson.Field
	URL          apijson.Field
	Instructions apijson.Field
	Mode         apijson.Field
	Time         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IntegrationAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationAttemptJSON) RawJSON() string {
	return r.raw
}

type IntegrationAttemptMode string

const (
	IntegrationAttemptModeAuto IntegrationAttemptMode = "auto"
	IntegrationAttemptModeCode IntegrationAttemptMode = "code"
)

func (r IntegrationAttemptMode) IsKnown() bool {
	switch r {
	case IntegrationAttemptModeAuto, IntegrationAttemptModeCode:
		return true
	}
	return false
}

// IntegrationAttemptTime represents the timing information for an integration attempt.
type IntegrationAttemptTime struct {
	// This field can have the runtime type of [float64], [string].
	Created interface{} `json:"created,required"`
	// This field can have the runtime type of [float64], [string].
	Expires interface{}                `json:"expires,required"`
	JSON    integrationAttemptTimeJSON `json:"-"`
}

// integrationAttemptTimeJSON contains the JSON metadata for the struct [IntegrationAttemptTime]
type integrationAttemptTimeJSON struct {
	Created     apijson.Field
	Expires     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationAttemptTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationAttemptTimeJSON) RawJSON() string {
	return r.raw
}

// ===== IntegrationAttemptStatus =====

// IntegrationAttemptStatus represents the status of an integration attempt.
// The status field determines which variant is returned: pending, complete,
// failed, or expired.
type IntegrationAttemptStatus struct {
	Status IntegrationAttemptStatusType `json:"status,required"`
	// Message is present only when status is "failed". This field can have the
	// runtime type of [string].
	Message interface{} `json:"message"`
	// This field can have the runtime type of [IntegrationAttemptTime].
	Time  interface{}                  `json:"time,required"`
	JSON  integrationAttemptStatusJSON `json:"-"`
	union IntegrationAttemptStatusUnion
}

// integrationAttemptStatusJSON contains the JSON metadata for the struct [IntegrationAttemptStatus]
type integrationAttemptStatusJSON struct {
	Status      apijson.Field
	Message     apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r integrationAttemptStatusJSON) RawJSON() string {
	return r.raw
}

func (r *IntegrationAttemptStatus) UnmarshalJSON(data []byte) (err error) {
	*r = IntegrationAttemptStatus{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IntegrationAttemptStatusUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [IntegrationAttemptStatusPending],
// [IntegrationAttemptStatusComplete], [IntegrationAttemptStatusFailed],
// [IntegrationAttemptStatusExpired].
func (r IntegrationAttemptStatus) AsUnion() IntegrationAttemptStatusUnion {
	return r.union
}

// IntegrationAttemptStatusUnion represents the union of integration attempt status
// types.
//
// Possible runtime types are [IntegrationAttemptStatusPending],
// [IntegrationAttemptStatusComplete], [IntegrationAttemptStatusFailed],
// [IntegrationAttemptStatusExpired].
type IntegrationAttemptStatusUnion interface {
	implementsIntegrationAttemptStatus()
}

type IntegrationAttemptStatusType string

const (
	IntegrationAttemptStatusTypePending  IntegrationAttemptStatusType = "pending"
	IntegrationAttemptStatusTypeComplete IntegrationAttemptStatusType = "complete"
	IntegrationAttemptStatusTypeFailed   IntegrationAttemptStatusType = "failed"
	IntegrationAttemptStatusTypeExpired  IntegrationAttemptStatusType = "expired"
)

func (r IntegrationAttemptStatusType) IsKnown() bool {
	switch r {
	case IntegrationAttemptStatusTypePending, IntegrationAttemptStatusTypeComplete,
		IntegrationAttemptStatusTypeFailed, IntegrationAttemptStatusTypeExpired:
		return true
	}
	return false
}

// IntegrationAttemptStatusPending represents the status of an integration attempt
// that is pending.
type IntegrationAttemptStatusPending struct {
	Status IntegrationAttemptStatusType        `json:"status,required"`
	Time   IntegrationAttemptTime              `json:"time,required"`
	JSON   integrationAttemptStatusPendingJSON `json:"-"`
}

// integrationAttemptStatusPendingJSON contains the JSON metadata for the struct
// [IntegrationAttemptStatusPending]
type integrationAttemptStatusPendingJSON struct {
	Status      apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationAttemptStatusPending) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationAttemptStatusPendingJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationAttemptStatusPending) implementsIntegrationAttemptStatus() {}

// IntegrationAttemptStatusComplete represents the status of an integration attempt
// that has completed.
type IntegrationAttemptStatusComplete struct {
	Status IntegrationAttemptStatusType         `json:"status,required"`
	Time   IntegrationAttemptTime               `json:"time,required"`
	JSON   integrationAttemptStatusCompleteJSON `json:"-"`
}

// integrationAttemptStatusCompleteJSON contains the JSON metadata for the struct
// [IntegrationAttemptStatusComplete]
type integrationAttemptStatusCompleteJSON struct {
	Status      apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationAttemptStatusComplete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationAttemptStatusCompleteJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationAttemptStatusComplete) implementsIntegrationAttemptStatus() {}

// IntegrationAttemptStatusFailed represents the status of an integration attempt
// that has failed.
type IntegrationAttemptStatusFailed struct {
	Status  IntegrationAttemptStatusType       `json:"status,required"`
	Message string                             `json:"message,required"`
	Time    IntegrationAttemptTime             `json:"time,required"`
	JSON    integrationAttemptStatusFailedJSON `json:"-"`
}

// integrationAttemptStatusFailedJSON contains the JSON metadata for the struct
// [IntegrationAttemptStatusFailed]
type integrationAttemptStatusFailedJSON struct {
	Status      apijson.Field
	Message     apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationAttemptStatusFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationAttemptStatusFailedJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationAttemptStatusFailed) implementsIntegrationAttemptStatus() {}

// IntegrationAttemptStatusExpired represents the status of an integration attempt
// that has expired.
type IntegrationAttemptStatusExpired struct {
	Status IntegrationAttemptStatusType        `json:"status,required"`
	Time   IntegrationAttemptTime              `json:"time,required"`
	JSON   integrationAttemptStatusExpiredJSON `json:"-"`
}

// integrationAttemptStatusExpiredJSON contains the JSON metadata for the struct
// [IntegrationAttemptStatusExpired]
type integrationAttemptStatusExpiredJSON struct {
	Status      apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationAttemptStatusExpired) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationAttemptStatusExpiredJSON) RawJSON() string {
	return r.raw
}

func (r IntegrationAttemptStatusExpired) implementsIntegrationAttemptStatus() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IntegrationAttemptStatusUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationAttemptStatusPending{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationAttemptStatusComplete{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationAttemptStatusFailed{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IntegrationAttemptStatusExpired{}),
		},
	)
}

// ===== Prompt Types =====

// IntegrationTextPrompt represents a text input prompt for integration configuration.
type IntegrationTextPrompt struct {
	Type        IntegrationTextPromptType `json:"type,required"`
	Key         string                    `json:"key,required"`
	Message     string                    `json:"message,required"`
	Placeholder string                    `json:"placeholder"`
	When        IntegrationWhen           `json:"when"`
	JSON        integrationTextPromptJSON `json:"-"`
}

// integrationTextPromptJSON contains the JSON metadata for the struct [IntegrationTextPrompt]
type integrationTextPromptJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Placeholder apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationTextPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationTextPromptJSON) RawJSON() string {
	return r.raw
}

type IntegrationTextPromptType string

const (
	IntegrationTextPromptTypeText IntegrationTextPromptType = "text"
)

func (r IntegrationTextPromptType) IsKnown() bool {
	switch r {
	case IntegrationTextPromptTypeText:
		return true
	}
	return false
}

// IntegrationSelectPrompt represents a select/dropdown prompt for integration configuration.
type IntegrationSelectPrompt struct {
	Type    IntegrationSelectPromptType `json:"type,required"`
	Key     string                      `json:"key,required"`
	Message string                      `json:"message,required"`
	Options []IntegrationSelectOption   `json:"options,required"`
	When    IntegrationWhen             `json:"when"`
	JSON    integrationSelectPromptJSON `json:"-"`
}

// integrationSelectPromptJSON contains the JSON metadata for the struct [IntegrationSelectPrompt]
type integrationSelectPromptJSON struct {
	Type        apijson.Field
	Key         apijson.Field
	Message     apijson.Field
	Options     apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationSelectPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationSelectPromptJSON) RawJSON() string {
	return r.raw
}

type IntegrationSelectPromptType string

const (
	IntegrationSelectPromptTypeSelect IntegrationSelectPromptType = "select"
)

func (r IntegrationSelectPromptType) IsKnown() bool {
	switch r {
	case IntegrationSelectPromptTypeSelect:
		return true
	}
	return false
}

// ===== IntegrationPrompt Union =====

// IntegrationPromptUnion represents the union of prompt types.
// Possible runtime types are [IntegrationTextPrompt], [IntegrationSelectPrompt].
type IntegrationPromptUnion interface {
	implementsIntegrationPromptUnion()
}

func (r IntegrationTextPrompt) implementsIntegrationPromptUnion() {}

func (r IntegrationSelectPrompt) implementsIntegrationPromptUnion() {}

// IntegrationSelectOption represents an option within a select prompt.
type IntegrationSelectOption struct {
	Label string                      `json:"label,required"`
	Value string                      `json:"value,required"`
	Hint  string                      `json:"hint"`
	JSON  integrationSelectOptionJSON `json:"-"`
}

// integrationSelectOptionJSON contains the JSON metadata for the struct [IntegrationSelectOption]
type integrationSelectOptionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	Hint        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationSelectOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationSelectOptionJSON) RawJSON() string {
	return r.raw
}

// ===== IntegrationWhen =====

// IntegrationWhen represents a conditional visibility rule for integration prompts.
type IntegrationWhen struct {
	Key   string              `json:"key,required"`
	Op    IntegrationWhenOp   `json:"op,required"`
	Value string              `json:"value,required"`
	JSON  integrationWhenJSON `json:"-"`
}

// integrationWhenJSON contains the JSON metadata for the struct [IntegrationWhen]
type integrationWhenJSON struct {
	Key         apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntegrationWhen) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationWhenJSON) RawJSON() string {
	return r.raw
}

type IntegrationWhenOp string

const (
	IntegrationWhenOpEq  IntegrationWhenOp = "eq"
	IntegrationWhenOpNeq IntegrationWhenOp = "neq"
)

func (r IntegrationWhenOp) IsKnown() bool {
	switch r {
	case IntegrationWhenOpEq, IntegrationWhenOpNeq:
		return true
	}
	return false
}

// ===== Param Types =====

// V2IntegrationListParams contains the query parameters for listing integrations.
type V2IntegrationListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2IntegrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2IntegrationGetParams contains the path and query parameters for getting an integration.
type V2IntegrationGetParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2IntegrationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2IntegrationConnectKeyParams contains the body and query parameters for connecting
// with a key.
type V2IntegrationConnectKeyParams struct {
	Location param.Field[V2LocationParam]      `query:"location"`
	Body     V2IntegrationConnectKeyParamsBody `json:"-"`
}

func (r V2IntegrationConnectKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

func (r V2IntegrationConnectKeyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2IntegrationConnectKeyParamsBody contains the body fields for the connect key request.
type V2IntegrationConnectKeyParamsBody struct {
	Key   param.Field[string] `json:"key,required"`
	Label param.Field[string] `json:"label"`
}

func (r V2IntegrationConnectKeyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// V2IntegrationConnectOauthParams contains the body and query parameters for
// beginning an OAuth connection.
type V2IntegrationConnectOauthParams struct {
	Location param.Field[V2LocationParam]        `query:"location"`
	Body     V2IntegrationConnectOauthParamsBody `json:"-"`
}

func (r V2IntegrationConnectOauthParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

func (r V2IntegrationConnectOauthParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2IntegrationConnectOauthParamsBody contains the body fields for the oauth connect request.
type V2IntegrationConnectOauthParamsBody struct {
	MethodID param.Field[string]            `json:"methodID,required"`
	Inputs   param.Field[map[string]string] `json:"inputs,required"`
	Label    param.Field[string]            `json:"label"`
}

func (r V2IntegrationConnectOauthParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// V2IntegrationAttemptStatusParams contains the query parameters for checking attempt status.
type V2IntegrationAttemptStatusParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2IntegrationAttemptStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2IntegrationAttemptCompleteParams contains the body and query parameters for
// completing an OAuth attempt.
type V2IntegrationAttemptCompleteParams struct {
	Location param.Field[V2LocationParam]           `query:"location"`
	Body     V2IntegrationAttemptCompleteParamsBody `json:"-"`
}

func (r V2IntegrationAttemptCompleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

func (r V2IntegrationAttemptCompleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2IntegrationAttemptCompleteParamsBody contains the body fields for completing an attempt.
type V2IntegrationAttemptCompleteParamsBody struct {
	Code param.Field[string] `json:"code"`
}

func (r V2IntegrationAttemptCompleteParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// V2IntegrationAttemptCancelParams contains the query parameters for cancelling an attempt.
type V2IntegrationAttemptCancelParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2IntegrationAttemptCancelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
