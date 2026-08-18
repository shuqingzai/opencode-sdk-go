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
	ID            string             `json:"id,required"`
	IntegrationID string             `json:"integrationID"`
	Name          string             `json:"name,required"`
	Disabled      bool               `json:"disabled"`
	API           V2ProviderInfoAPI  `json:"api,required"`
	Request       ProviderRequest    `json:"request,required"`
	JSON          v2ProviderInfoJSON `json:"-"`
}

// v2ProviderInfoJSON contains the JSON metadata for the struct [V2ProviderInfo]
type v2ProviderInfoJSON struct {
	ID            apijson.Field
	IntegrationID apijson.Field
	Name          apijson.Field
	Disabled      apijson.Field
	API           apijson.Field
	Request       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *V2ProviderInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderInfoJSON) RawJSON() string {
	return r.raw
}

// AsAPIUnion returns the api field as a typed union.
//
// Deprecated: use [V2ProviderInfo.API.AsUnion] instead.
func (r *V2ProviderInfo) AsAPIUnion() V2ProviderInfoAPIUnion {
	return r.API.AsUnion()
}

// V2ProviderInfoAPI is the union bearer for the api field of [V2ProviderInfo].
// It holds the decoded API configuration and provides typed access via [AsUnion].
//
// The runtime union variant can be one of [V2ProviderInfoAPIAisdk] or
// [V2ProviderInfoAPINative]; use [AsUnion] to obtain the concrete type.
type V2ProviderInfoAPI struct {
	Type    V2ProviderInfoAPIType `json:"type,required"`
	Package string                `json:"package"`
	URL     string                `json:"url"`
	// This field can have the runtime type of [map[string]any].
	Settings map[string]any        `json:"settings"`
	JSON     v2ProviderInfoAPIJSON `json:"-"`
	union    V2ProviderInfoAPIUnion
}

// v2ProviderInfoAPIJSON contains the JSON metadata for the struct [V2ProviderInfoAPI]
type v2ProviderInfoAPIJSON struct {
	Type        apijson.Field
	Package     apijson.Field
	URL         apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r v2ProviderInfoAPIJSON) RawJSON() string {
	return r.raw
}

func (r *V2ProviderInfoAPI) UnmarshalJSON(data []byte) (err error) {
	*r = V2ProviderInfoAPI{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V2ProviderInfoAPIUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [V2ProviderInfoAPIAisdk],
// [V2ProviderInfoAPINative].
func (r V2ProviderInfoAPI) AsUnion() V2ProviderInfoAPIUnion {
	return r.union
}

// Deprecated: use [V2ProviderInfoAPI] instead.
type V2ProviderInfoApi = V2ProviderInfoAPI

// V2ProviderInfoAPIType is the discriminator shared by every
// [V2ProviderInfoAPIUnion] variant.
type V2ProviderInfoAPIType string

const (
	V2ProviderInfoAPITypeAisdk  V2ProviderInfoAPIType = "aisdk"
	V2ProviderInfoAPITypeNative V2ProviderInfoAPIType = "native"
)

// Deprecated: use [V2ProviderInfoAPIType] instead.
type V2ProviderInfoApiType = V2ProviderInfoAPIType

// Deprecated: use [V2ProviderInfoAPITypeAisdk] instead.
const V2ProviderInfoApiTypeAisdk = V2ProviderInfoAPITypeAisdk

// Deprecated: use [V2ProviderInfoAPITypeNative] instead.
const V2ProviderInfoApiTypeNative = V2ProviderInfoAPITypeNative

func (r V2ProviderInfoAPIType) IsKnown() bool {
	switch r {
	case V2ProviderInfoAPITypeAisdk, V2ProviderInfoAPITypeNative:
		return true
	}
	return false
}

// V2ProviderInfoAPIUnion represents the api configuration of a provider.
// Possible runtime types are [V2ProviderInfoAPIAisdk], [V2ProviderInfoAPINative].
type V2ProviderInfoAPIUnion interface {
	implementsV2ProviderInfoAPIUnion()
}

// Deprecated: use [V2ProviderInfoAPIUnion] instead.
type V2ProviderInfoApiUnion = V2ProviderInfoAPIUnion

type V2ProviderInfoAPIAisdk struct {
	Type    V2ProviderInfoAPIAisdkType `json:"type,required"`
	Package string                     `json:"package,required"`
	URL     string                     `json:"url"`
	// This field can have the runtime type of [map[string]any].
	Settings map[string]any             `json:"settings"`
	JSON     v2ProviderInfoAPIAisdkJSON `json:"-"`
}

type v2ProviderInfoAPIAisdkJSON struct {
	Type        apijson.Field
	Package     apijson.Field
	URL         apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderInfoAPIAisdk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderInfoAPIAisdkJSON) RawJSON() string {
	return r.raw
}

func (r V2ProviderInfoAPIAisdk) implementsV2ProviderInfoAPIUnion() {}

// Deprecated: use [V2ProviderInfoAPIAisdk] instead.
type V2ProviderInfoApiAisdk = V2ProviderInfoAPIAisdk

type V2ProviderInfoAPINative struct {
	Type V2ProviderInfoAPINativeType `json:"type,required"`
	URL  string                      `json:"url"`
	// This field can have the runtime type of [map[string]any].
	Settings map[string]any              `json:"settings,required"`
	JSON     v2ProviderInfoAPINativeJSON `json:"-"`
}

type v2ProviderInfoAPINativeJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProviderInfoAPINative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProviderInfoAPINativeJSON) RawJSON() string {
	return r.raw
}

func (r V2ProviderInfoAPINative) implementsV2ProviderInfoAPIUnion() {}

// Deprecated: use [V2ProviderInfoAPINative] instead.
type V2ProviderInfoApiNative = V2ProviderInfoAPINative

type V2ProviderInfoAPIAisdkType string

const (
	V2ProviderInfoAPIAisdkTypeAisdk V2ProviderInfoAPIAisdkType = "aisdk"
)

// Deprecated: use [V2ProviderInfoAPIAisdkType] instead.
type V2ProviderInfoApiAisdkType = V2ProviderInfoAPIAisdkType

// Deprecated: use [V2ProviderInfoAPIAisdkTypeAisdk] instead.
const V2ProviderInfoApiAisdkTypeAisdk = V2ProviderInfoAPIAisdkTypeAisdk

func (r V2ProviderInfoAPIAisdkType) IsKnown() bool {
	switch r {
	case V2ProviderInfoAPIAisdkTypeAisdk:
		return true
	}
	return false
}

type V2ProviderInfoAPINativeType string

const (
	V2ProviderInfoAPINativeTypeNative V2ProviderInfoAPINativeType = "native"
)

// Deprecated: use [V2ProviderInfoAPINativeType] instead.
type V2ProviderInfoApiNativeType = V2ProviderInfoAPINativeType

// Deprecated: use [V2ProviderInfoAPINativeTypeNative] instead.
const V2ProviderInfoApiNativeTypeNative = V2ProviderInfoAPINativeTypeNative

func (r V2ProviderInfoAPINativeType) IsKnown() bool {
	switch r {
	case V2ProviderInfoAPINativeTypeNative:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[V2ProviderInfoAPIUnion](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "aisdk",
			Type:               reflect.TypeFor[V2ProviderInfoAPIAisdk](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "native",
			Type:               reflect.TypeFor[V2ProviderInfoAPINative](),
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
