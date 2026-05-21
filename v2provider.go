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

// V2ProviderService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2ProviderService] method instead.
type V2ProviderService struct {
	Options []option.RequestOption
}

// NewV2ProviderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2ProviderService(opts ...option.RequestOption) (r *V2ProviderService) {
	r = &V2ProviderService{}
	r.Options = opts
	return
}

// List v2 providers
func (r *V2ProviderService) List(ctx context.Context, query V2ProviderListParams, opts ...option.RequestOption) (res *[]V2ProviderInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/provider"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get v2 provider
func (r *V2ProviderService) Get(ctx context.Context, providerID string, query V2ProviderGetParams, opts ...option.RequestOption) (res *V2ProviderInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	if providerID == "" {
		err = errors.New("missing required providerID parameter")
		return
	}
	path := fmt.Sprintf("api/provider/%s", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V2ProviderInfo struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// This field can have the runtime type of [V2ProviderInfoEnabledEnv],
	// [V2ProviderInfoEnabledAuth], [V2ProviderInfoEnabledCustom].
	// When the provider is disabled, this field is `false` (a JSON boolean).
	Enabled  interface{}            `json:"enabled,required"`
	Env      []string               `json:"env,required"`
	// This field can have the runtime type of [V2ModelInfoEndpointUnknown],
	// [V2ModelInfoEndpointOpenAIResponses], [V2ModelInfoEndpointOpenAICompletions],
	// [V2ModelInfoEndpointAnthropicMessages], [V2ModelInfoEndpointAisdk].
	Endpoint interface{}            `json:"endpoint,required"`
	Options  V2ModelInfoOptions     `json:"options,required"`
	JSON     v2ProviderInfoJSON     `json:"-"`
	enabledUnion  V2ProviderInfoEnabledUnion
	endpointUnion V2ModelInfoEndpointUnion
}

// v2ProviderInfoJSON contains the JSON metadata for the struct [V2ProviderInfo]
type v2ProviderInfoJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Enabled     apijson.Field
	Env         apijson.Field
	Endpoint    apijson.Field
	Options     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderInfo) UnmarshalJSON(data []byte) (err error) {
	*r = V2ProviderInfo{}
	err = apijson.UnmarshalRoot(data, &r.enabledUnion)
	if err != nil {
		return err
	}
	err = apijson.UnmarshalRoot(data, &r.endpointUnion)
	if err != nil {
		return err
	}
	err = apijson.Port(r.enabledUnion, r)
	if err != nil {
		return err
	}
	return apijson.Port(r.endpointUnion, r)
}

func (r v2ProviderInfoJSON) RawJSON() string {
	return r.raw
}

// AsEnabledUnion returns the enabled field as a typed union.
func (r *V2ProviderInfo) AsEnabledUnion() V2ProviderInfoEnabledUnion {
	return r.enabledUnion
}

// AsEndpointUnion returns the endpoint field as a typed union.
func (r *V2ProviderInfo) AsEndpointUnion() V2ModelInfoEndpointUnion {
	return r.endpointUnion
}

// V2ProviderInfoEnabledUnion represents the enabled state of a provider.
// Possible runtime types are [V2ProviderInfoEnabledEnv], [V2ProviderInfoEnabledAuth],
// [V2ProviderInfoEnabledCustom]. When disabled, the value is `false` (a JSON boolean).
type V2ProviderInfoEnabledUnion interface {
	implementsV2ProviderInfoEnabledUnion()
}

type V2ProviderInfoEnabledEnv struct {
	Via  string                          `json:"via,required"`
	Name string                          `json:"name,required"`
	JSON v2ProviderInfoEnabledEnvJSON    `json:"-"`
}

type v2ProviderInfoEnabledEnvJSON struct {
	Via         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderInfoEnabledEnv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderInfoEnabledEnvJSON) RawJSON() string {
	return r.raw
}

func (r V2ProviderInfoEnabledEnv) implementsV2ProviderInfoEnabledUnion() {}

type V2ProviderInfoEnabledAuth struct {
	Via     string                           `json:"via,required"`
	Service string                           `json:"service,required"`
	JSON    v2ProviderInfoEnabledAuthJSON    `json:"-"`
}

type v2ProviderInfoEnabledAuthJSON struct {
	Via         apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderInfoEnabledAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderInfoEnabledAuthJSON) RawJSON() string {
	return r.raw
}

func (r V2ProviderInfoEnabledAuth) implementsV2ProviderInfoEnabledUnion() {}

type V2ProviderInfoEnabledCustom struct {
	Via  string                            `json:"via,required"`
	Data map[string]interface{}            `json:"data,required"`
	JSON v2ProviderInfoEnabledCustomJSON   `json:"-"`
}

type v2ProviderInfoEnabledCustomJSON struct {
	Via         apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderInfoEnabledCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderInfoEnabledCustomJSON) RawJSON() string {
	return r.raw
}

func (r V2ProviderInfoEnabledCustom) implementsV2ProviderInfoEnabledUnion() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ProviderInfoEnabledUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ProviderInfoEnabledEnv{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ProviderInfoEnabledAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ProviderInfoEnabledCustom{}),
		},
	)
}

type V2ProviderListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2ProviderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2ProviderGetParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2ProviderGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
