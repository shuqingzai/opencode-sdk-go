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
func (r *V2ProviderService) List(ctx context.Context, query V2ProviderListParams, opts ...option.RequestOption) (res *V2ProviderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/provider"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get v2 provider
func (r *V2ProviderService) Get(ctx context.Context, providerID string, query V2ProviderGetParams, opts ...option.RequestOption) (res *V2ProviderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if providerID == "" {
		err = errors.New("missing required providerID parameter")
		return
	}
	path := fmt.Sprintf("api/provider/%s", providerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// V2ProviderListResponse is returned by the Provider.List method. It wraps a list of
// providers alongside the active location metadata.
type V2ProviderListResponse struct {
	Location LocationInfo               `json:"location,required"`
	Data     []V2ProviderInfo           `json:"data,required"`
	JSON     v2ProviderListResponseJSON `json:"-"`
}

// v2ProviderListResponseJSON contains the JSON metadata for the struct [V2ProviderListResponse]
type v2ProviderListResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderListResponseJSON) RawJSON() string {
	return r.raw
}

// V2ProviderGetResponse is returned by the Provider.Get method. It wraps the requested
// provider alongside the active location metadata.
type V2ProviderGetResponse struct {
	Location LocationInfo              `json:"location,required"`
	Data     V2ProviderInfo            `json:"data,required"`
	JSON     v2ProviderGetResponseJSON `json:"-"`
}

// v2ProviderGetResponseJSON contains the JSON metadata for the struct [V2ProviderGetResponse]
type v2ProviderGetResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderGetResponseJSON) RawJSON() string {
	return r.raw
}

type V2ProviderInfo struct {
	ID            string `json:"id,required"`
	IntegrationID string `json:"integrationID"`
	Name          string `json:"name,required"`
	Disabled      bool   `json:"disabled"`
	// This field can have the runtime type of [V2ProviderInfoApiAisdk], [V2ProviderInfoApiNative].
	Api      interface{}        `json:"api,required"`
	Request  ProviderRequest    `json:"request,required"`
	JSON     v2ProviderInfoJSON `json:"-"`
	apiUnion V2ProviderInfoApiUnion
}

// v2ProviderInfoJSON contains the JSON metadata for the struct [V2ProviderInfo]
type v2ProviderInfoJSON struct {
	ID            apijson.Field
	IntegrationID apijson.Field
	Name          apijson.Field
	Disabled      apijson.Field
	Api           apijson.Field
	Request       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2ProviderInfo) UnmarshalJSON(data []byte) (err error) {
	*r = V2ProviderInfo{}
	err = apijson.UnmarshalRoot(data, r)
	if err != nil {
		return err
	}
	apiData := gjson.GetBytes(data, "api").Raw
	if apiData != "" {
		err = apijson.UnmarshalRoot([]byte(apiData), &r.apiUnion)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r v2ProviderInfoJSON) RawJSON() string {
	return r.raw
}

// AsAPIUnion returns the api field as a typed union.
func (r *V2ProviderInfo) AsAPIUnion() V2ProviderInfoApiUnion {
	return r.apiUnion
}

// V2ProviderInfoApiUnion represents the api configuration of a provider.
// Possible runtime types are [V2ProviderInfoApiAisdk], [V2ProviderInfoApiNative].
type V2ProviderInfoApiUnion interface {
	implementsV2ProviderInfoApiUnion()
}

type V2ProviderInfoApiAisdk struct {
	Type    V2ProviderInfoApiAisdkType `json:"type,required"`
	Package string                     `json:"package,required"`
	URL     string                     `json:"url"`
	// This field can have the runtime type of [map[string]interface{}].
	Settings interface{}                `json:"settings"`
	JSON     v2ProviderInfoApiAisdkJSON `json:"-"`
}

type v2ProviderInfoApiAisdkJSON struct {
	Type        apijson.Field
	Package     apijson.Field
	URL         apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderInfoApiAisdk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderInfoApiAisdkJSON) RawJSON() string {
	return r.raw
}

func (r V2ProviderInfoApiAisdk) implementsV2ProviderInfoApiUnion() {}

type V2ProviderInfoApiNative struct {
	Type V2ProviderInfoApiNativeType `json:"type,required"`
	URL  string                      `json:"url"`
	// This field can have the runtime type of [map[string]interface{}].
	Settings interface{}                 `json:"settings,required"`
	JSON     v2ProviderInfoApiNativeJSON `json:"-"`
}

type v2ProviderInfoApiNativeJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderInfoApiNative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderInfoApiNativeJSON) RawJSON() string {
	return r.raw
}

func (r V2ProviderInfoApiNative) implementsV2ProviderInfoApiUnion() {}

type V2ProviderInfoApiAisdkType string

const (
	V2ProviderInfoApiAisdkTypeAisdk V2ProviderInfoApiAisdkType = "aisdk"
)

func (r V2ProviderInfoApiAisdkType) IsKnown() bool {
	switch r {
	case V2ProviderInfoApiAisdkTypeAisdk:
		return true
	}
	return false
}

type V2ProviderInfoApiNativeType string

const (
	V2ProviderInfoApiNativeTypeNative V2ProviderInfoApiNativeType = "native"
)

func (r V2ProviderInfoApiNativeType) IsKnown() bool {
	switch r {
	case V2ProviderInfoApiNativeTypeNative:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ProviderInfoApiUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ProviderInfoApiAisdk{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ProviderInfoApiNative{}),
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
