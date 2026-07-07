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

// V2ReferenceService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2ReferenceService] method instead.
type V2ReferenceService struct {
	Options []option.RequestOption
}

// NewV2ReferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2ReferenceService(opts ...option.RequestOption) (r *V2ReferenceService) {
	r = &V2ReferenceService{}
	r.Options = opts
	return
}

// List references
//
// List references available in the requested location.
func (r *V2ReferenceService) List(ctx context.Context, query V2ReferenceListParams, opts ...option.RequestOption) (res *V2ReferenceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/reference"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// V2ReferenceListResponse contains the response from the reference list endpoint.
type V2ReferenceListResponse struct {
	Location LocationInfo               `json:"location,required"`
	Data     []V2ReferenceInfo          `json:"data,required"`
	JSON     v2ReferenceListResponseJSON `json:"-"`
}

// v2ReferenceListResponseJSON contains the JSON metadata for the struct [V2ReferenceListResponse]
type v2ReferenceListResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ReferenceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ReferenceListResponseJSON) RawJSON() string {
	return r.raw
}

// V2ReferenceInfo represents a reference entry.
type V2ReferenceInfo struct {
	Name        string `json:"name,required"`
	Path        string `json:"path,required"`
	Description string `json:"description"`
	Hidden      bool   `json:"hidden"`
	// This field can have the runtime type of [ReferenceLocalSource],
	// [ReferenceGitSource].
	Source      interface{}          `json:"source,required"`
	JSON        v2ReferenceInfoJSON  `json:"-"`
	sourceUnion ReferenceSourceUnion
}

// v2ReferenceInfoJSON contains the JSON metadata for the struct [V2ReferenceInfo]
type v2ReferenceInfoJSON struct {
	Name        apijson.Field
	Path        apijson.Field
	Description apijson.Field
	Hidden      apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ReferenceInfo) UnmarshalJSON(data []byte) (err error) {
	*r = V2ReferenceInfo{}
	err = apijson.UnmarshalRoot(data, &r.sourceUnion)
	if err != nil {
		return err
	}
	return apijson.Port(r.sourceUnion, r)
}

func (r v2ReferenceInfoJSON) RawJSON() string {
	return r.raw
}

// AsSourceUnion returns the source field as a typed union.
func (r *V2ReferenceInfo) AsSourceUnion() ReferenceSourceUnion {
	return r.sourceUnion
}

// ReferenceSourceUnion represents the source of a reference.
// Possible runtime types are [ReferenceLocalSource], [ReferenceGitSource].
type ReferenceSourceUnion interface {
	implementsReferenceSourceUnion()
}

// ReferenceLocalSource represents a local reference source.
type ReferenceLocalSource struct {
	Type        string                      `json:"type,required"`
	Path        string                      `json:"path,required"`
	Description string                      `json:"description"`
	Hidden      bool                        `json:"hidden"`
	JSON        referenceLocalSourceJSON    `json:"-"`
}

// referenceLocalSourceJSON contains the JSON metadata for the struct [ReferenceLocalSource]
type referenceLocalSourceJSON struct {
	Type        apijson.Field
	Path        apijson.Field
	Description apijson.Field
	Hidden      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferenceLocalSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referenceLocalSourceJSON) RawJSON() string {
	return r.raw
}

func (r ReferenceLocalSource) implementsReferenceSourceUnion() {}

// ReferenceGitSource represents a git reference source.
type ReferenceGitSource struct {
	Type        string                    `json:"type,required"`
	Repository  string                    `json:"repository,required"`
	Branch      string                    `json:"branch"`
	Description string                    `json:"description"`
	Hidden      bool                      `json:"hidden"`
	JSON        referenceGitSourceJSON    `json:"-"`
}

// referenceGitSourceJSON contains the JSON metadata for the struct [ReferenceGitSource]
type referenceGitSourceJSON struct {
	Type        apijson.Field
	Repository  apijson.Field
	Branch      apijson.Field
	Description apijson.Field
	Hidden      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferenceGitSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referenceGitSourceJSON) RawJSON() string {
	return r.raw
}

func (r ReferenceGitSource) implementsReferenceSourceUnion() {}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ReferenceSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReferenceLocalSource{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ReferenceGitSource{}),
		},
	)
}

// V2ReferenceListParams contains the query parameters for listing references.
type V2ReferenceListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2ReferenceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
