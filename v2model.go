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
	"github.com/tidwall/gjson"
)

// V2ModelService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2ModelService] method instead.
type V2ModelService struct {
	Options []option.RequestOption
}

// NewV2ModelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2ModelService(opts ...option.RequestOption) (r *V2ModelService) {
	r = &V2ModelService{}
	r.Options = opts
	return
}

// List v2 models
func (r *V2ModelService) List(ctx context.Context, query V2ModelListParams, opts ...option.RequestOption) (res *V2ModelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// V2ModelListResponse is returned by the Model.List method. It wraps a list of models
// alongside the active location metadata.
type V2ModelListResponse struct {
	Location LocationInfo            `json:"location,required"`
	Data     []V2ModelInfo           `json:"data,required"`
	JSON     v2ModelListResponseJSON `json:"-"`
}

// v2ModelListResponseJSON contains the JSON metadata for the struct [V2ModelListResponse]
type v2ModelListResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelListResponseJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfo struct {
	ID         string `json:"id,required"`
	ProviderID string `json:"providerID,required"`
	Family     string `json:"family"`
	Name       string `json:"name,required"`
	// This field can have the runtime type of [V2ModelInfoApiAisdk], [V2ModelInfoApiNative].
	Api          interface{}             `json:"api,required"`
	Capabilities V2ModelInfoCapabilities `json:"capabilities,required"`
	Request      V2ModelInfoRequest      `json:"request,required"`
	Variants     []V2ModelInfoVariant    `json:"variants,required"`
	Time         V2ModelInfoTime         `json:"time,required"`
	Cost         []V2ModelInfoCostItem   `json:"cost,required"`
	Status       V2ModelInfoStatus       `json:"status,required"`
	Enabled      bool                    `json:"enabled,required"`
	Limit        V2ModelInfoLimit        `json:"limit,required"`
	JSON         v2ModelInfoJSON         `json:"-"`
	apiUnion     V2ModelInfoApiUnion
}

// v2ModelInfoJSON contains the JSON metadata for the struct [V2ModelInfo]
type v2ModelInfoJSON struct {
	ID           apijson.Field
	ProviderID   apijson.Field
	Family       apijson.Field
	Name         apijson.Field
	Api          apijson.Field
	Capabilities apijson.Field
	Request      apijson.Field
	Variants     apijson.Field
	Time         apijson.Field
	Cost         apijson.Field
	Status       apijson.Field
	Enabled      apijson.Field
	Limit        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2ModelInfo) UnmarshalJSON(data []byte) (err error) {
	*r = V2ModelInfo{}
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

func (r v2ModelInfoJSON) RawJSON() string {
	return r.raw
}

// AsAPIUnion returns the api field as a typed union.
func (r *V2ModelInfo) AsAPIUnion() V2ModelInfoApiUnion {
	return r.apiUnion
}

// V2ModelInfoApiUnion represents the api configuration of a model.
// OpenAPI anyOf has no discriminator; variants are selected by structural matching
// on the type enum field.
// Possible runtime types are [V2ModelInfoApiAisdk], [V2ModelInfoApiNative].
type V2ModelInfoApiUnion interface {
	implementsV2ModelInfoApiUnion()
}

type V2ModelInfoApiAisdk struct {
	ID      string                  `json:"id,required"`
	Type    V2ModelInfoApiAisdkType `json:"type,required"`
	Package string                  `json:"package,required"`
	URL     string                  `json:"url"`
	// This field can have the runtime type of [map[string]interface{}].
	Settings interface{}             `json:"settings"`
	JSON     v2ModelInfoApiAisdkJSON `json:"-"`
}

type v2ModelInfoApiAisdkJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Package     apijson.Field
	URL         apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoApiAisdk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoApiAisdkJSON) RawJSON() string {
	return r.raw
}

func (r V2ModelInfoApiAisdk) implementsV2ModelInfoApiUnion() {}

type V2ModelInfoApiNative struct {
	ID   string                   `json:"id,required"`
	Type V2ModelInfoApiNativeType `json:"type,required"`
	URL  string                   `json:"url"`
	// This field can have the runtime type of [map[string]interface{}].
	Settings interface{}              `json:"settings,required"`
	JSON     v2ModelInfoApiNativeJSON `json:"-"`
}

type v2ModelInfoApiNativeJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoApiNative) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoApiNativeJSON) RawJSON() string {
	return r.raw
}

func (r V2ModelInfoApiNative) implementsV2ModelInfoApiUnion() {}

type V2ModelInfoApiAisdkType string

const (
	V2ModelInfoApiAisdkTypeAisdk V2ModelInfoApiAisdkType = "aisdk"
)

func (r V2ModelInfoApiAisdkType) IsKnown() bool {
	switch r {
	case V2ModelInfoApiAisdkTypeAisdk:
		return true
	}
	return false
}

type V2ModelInfoApiNativeType string

const (
	V2ModelInfoApiNativeTypeNative V2ModelInfoApiNativeType = "native"
)

func (r V2ModelInfoApiNativeType) IsKnown() bool {
	switch r {
	case V2ModelInfoApiNativeTypeNative:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ModelInfoApiUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ModelInfoApiAisdk{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V2ModelInfoApiNative{}),
		},
	)
}

type V2ModelInfoRequest struct {
	Headers map[string]string `json:"headers,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Body    interface{}            `json:"body,required"`
	Variant string                 `json:"variant"`
	JSON    v2ModelInfoRequestJSON `json:"-"`
}

type v2ModelInfoRequestJSON struct {
	Headers     apijson.Field
	Body        apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoRequestJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoVariant struct {
	ID      string            `json:"id,required"`
	Headers map[string]string `json:"headers,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Body interface{}            `json:"body,required"`
	JSON v2ModelInfoVariantJSON `json:"-"`
}

type v2ModelInfoVariantJSON struct {
	ID          apijson.Field
	Headers     apijson.Field
	Body        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoVariantJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoStatus string

const (
	V2ModelInfoStatusAlpha      V2ModelInfoStatus = "alpha"
	V2ModelInfoStatusBeta       V2ModelInfoStatus = "beta"
	V2ModelInfoStatusDeprecated V2ModelInfoStatus = "deprecated"
	V2ModelInfoStatusActive     V2ModelInfoStatus = "active"
)

func (r V2ModelInfoStatus) IsKnown() bool {
	switch r {
	case V2ModelInfoStatusAlpha, V2ModelInfoStatusBeta, V2ModelInfoStatusDeprecated, V2ModelInfoStatusActive:
		return true
	}
	return false
}

type V2ModelInfoCapabilities struct {
	Tools  bool                        `json:"tools,required"`
	Input  []string                    `json:"input,required"`
	Output []string                    `json:"output,required"`
	JSON   v2ModelInfoCapabilitiesJSON `json:"-"`
}

type v2ModelInfoCapabilitiesJSON struct {
	Tools       apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoCapabilities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoCapabilitiesJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoTime struct {
	// The release date as a Unix timestamp.
	Released int64               `json:"released,required"`
	JSON     v2ModelInfoTimeJSON `json:"-"`
}

type v2ModelInfoTimeJSON struct {
	Released    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoTimeJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoCostItem struct {
	Tier   V2ModelInfoCostTier     `json:"tier"`
	Input  float64                 `json:"input,required"`
	Output float64                 `json:"output,required"`
	Cache  V2ModelInfoCostCache    `json:"cache,required"`
	JSON   v2ModelInfoCostItemJSON `json:"-"`
}

type v2ModelInfoCostItemJSON struct {
	Tier        apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoCostItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoCostItemJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoCostTierType string

const (
	V2ModelInfoCostTierTypeContext V2ModelInfoCostTierType = "context"
)

func (r V2ModelInfoCostTierType) IsKnown() bool {
	switch r {
	case V2ModelInfoCostTierTypeContext:
		return true
	}
	return false
}

type V2ModelInfoCostTier struct {
	Type V2ModelInfoCostTierType `json:"type,required"`
	Size int64                   `json:"size,required"`
	JSON v2ModelInfoCostTierJSON `json:"-"`
}

type v2ModelInfoCostTierJSON struct {
	Type        apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoCostTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoCostTierJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoCostCache struct {
	Read  float64                  `json:"read,required"`
	Write float64                  `json:"write,required"`
	JSON  v2ModelInfoCostCacheJSON `json:"-"`
}

type v2ModelInfoCostCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ModelInfoCostCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoCostCacheJSON) RawJSON() string {
	return r.raw
}

type V2ModelInfoLimit struct {
	Context int64                `json:"context,required"`
	Input   int64                `json:"input"`
	Output  int64                `json:"output,required"`
	JSON    v2ModelInfoLimitJSON `json:"-"`
}

type v2ModelInfoLimitJSON struct {
	Context     apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	ExtraFields map[string]apijson.Field
	raw         string
}

func (r *V2ModelInfoLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ModelInfoLimitJSON) RawJSON() string {
	return r.raw
}

type V2ModelListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2ModelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
