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
	Location LocationInfo                `json:"location,required"`
	Data     []V2ReferenceInfo           `json:"data,required"`
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
	Name        string              `json:"name,required"`
	Path        string              `json:"path,required"`
	Description string              `json:"description"`
	Hidden      bool                `json:"hidden"`
	Source      ReferenceSource     `json:"source,required"`
	JSON        v2ReferenceInfoJSON `json:"-"`
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
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ReferenceInfoJSON) RawJSON() string {
	return r.raw
}

// AsSourceUnion returns the source field as a typed union.
//
// Deprecated: use [V2ReferenceInfo.Source.AsUnion] instead.
func (r *V2ReferenceInfo) AsSourceUnion() ReferenceSourceUnion {
	return r.Source.AsUnion()
}

// ReferenceSource is the union bearer for the source field of [V2ReferenceInfo].
// It holds the decoded source configuration and provides typed access via [AsUnion].
//
// The runtime union variant can be one of [ReferenceLocalSource] or
// [ReferenceGitSource]; use [AsUnion] to obtain the concrete type.
type ReferenceSource struct {
	Type        ReferenceSourceType `json:"type,required"`
	Path        string              `json:"path"`
	Repository  string              `json:"repository"`
	Branch      string              `json:"branch"`
	Description string              `json:"description"`
	Hidden      bool                `json:"hidden"`
	JSON        referenceSourceJSON `json:"-"`
	union       ReferenceSourceUnion
}

// referenceSourceJSON contains the JSON metadata for the struct [ReferenceSource]
type referenceSourceJSON struct {
	Type        apijson.Field
	Path        apijson.Field
	Repository  apijson.Field
	Branch      apijson.Field
	Description apijson.Field
	Hidden      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r referenceSourceJSON) RawJSON() string {
	return r.raw
}

func (r *ReferenceSource) UnmarshalJSON(data []byte) (err error) {
	*r = ReferenceSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ReferenceSourceUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ReferenceLocalSource],
// [ReferenceGitSource].
func (r ReferenceSource) AsUnion() ReferenceSourceUnion {
	return r.union
}

// ReferenceSourceType is the discriminator shared by every [ReferenceSourceUnion]
// variant.
type ReferenceSourceType string

const (
	ReferenceSourceTypeLocal ReferenceSourceType = "local"
	ReferenceSourceTypeGit   ReferenceSourceType = "git"
)

func (r ReferenceSourceType) IsKnown() bool {
	switch r {
	case ReferenceSourceTypeLocal, ReferenceSourceTypeGit:
		return true
	}
	return false
}

// ReferenceSourceUnion represents the source of a reference.
// Possible runtime types are [ReferenceLocalSource], [ReferenceGitSource].
type ReferenceSourceUnion interface {
	implementsReferenceSourceUnion()
}

// ReferenceLocalSource represents a local reference source.
type ReferenceLocalSource struct {
	Type        ReferenceLocalSourceType `json:"type,required"`
	Path        string                   `json:"path,required"`
	Description string                   `json:"description"`
	Hidden      bool                     `json:"hidden"`
	JSON        referenceLocalSourceJSON `json:"-"`
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
	Type        ReferenceGitSourceType `json:"type,required"`
	Repository  string                 `json:"repository,required"`
	Branch      string                 `json:"branch"`
	Description string                 `json:"description"`
	Hidden      bool                   `json:"hidden"`
	JSON        referenceGitSourceJSON `json:"-"`
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

type ReferenceLocalSourceType string

const (
	ReferenceLocalSourceTypeLocal ReferenceLocalSourceType = "local"
)

func (r ReferenceLocalSourceType) IsKnown() bool {
	switch r {
	case ReferenceLocalSourceTypeLocal:
		return true
	}
	return false
}

type ReferenceGitSourceType string

const (
	ReferenceGitSourceTypeGit ReferenceGitSourceType = "git"
)

func (r ReferenceGitSourceType) IsKnown() bool {
	switch r {
	case ReferenceGitSourceTypeGit:
		return true
	}
	return false
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[ReferenceSourceUnion](),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "local",
			Type:               reflect.TypeFor[ReferenceLocalSource](),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "git",
			Type:               reflect.TypeFor[ReferenceGitSource](),
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
