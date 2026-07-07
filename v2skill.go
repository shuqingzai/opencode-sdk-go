// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// V2SkillService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2SkillService] method instead.
type V2SkillService struct {
	Options []option.RequestOption
}

// NewV2SkillService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2SkillService(opts ...option.RequestOption) (r *V2SkillService) {
	r = &V2SkillService{}
	r.Options = opts
	return
}

// List skills
func (r *V2SkillService) List(ctx context.Context, query V2SkillListParams, opts ...option.RequestOption) (res *V2SkillListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/skill"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// V2SkillListResponse contains the response from the skill list endpoint.
type V2SkillListResponse struct {
	Location LocationInfo    `json:"location,required"`
	Data     []SkillV2Info   `json:"data,required"`
	JSON     v2SkillListResponseJSON `json:"-"`
}

// v2SkillListResponseJSON contains the JSON metadata for the struct [V2SkillListResponse]
type v2SkillListResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SkillListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SkillListResponseJSON) RawJSON() string {
	return r.raw
}

// SkillV2Info represents a v2 skill entry.
type SkillV2Info struct {
	Name        string `json:"name,required"`
	Description string `json:"description"`
	Slash       bool   `json:"slash"`
	Location    string `json:"location,required"`
	Content     string `json:"content,required"`
	JSON        skillV2InfoJSON `json:"-"`
}

// skillV2InfoJSON contains the JSON metadata for the struct [SkillV2Info]
type skillV2InfoJSON struct {
	Name        apijson.Field
	Description apijson.Field
	Slash       apijson.Field
	Location    apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkillV2Info) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skillV2InfoJSON) RawJSON() string {
	return r.raw
}

// LocationInfo represents location metadata returned in list responses.
type LocationInfo struct {
	Directory   string                `json:"directory,required"`
	WorkspaceID string                `json:"workspaceID"`
	Project     LocationInfoProject   `json:"project,required"`
	JSON        locationInfoJSON      `json:"-"`
}

// locationInfoJSON contains the JSON metadata for the struct [LocationInfo]
type locationInfoJSON struct {
	Directory    apijson.Field
	WorkspaceID  apijson.Field
	Project      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LocationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationInfoJSON) RawJSON() string {
	return r.raw
}

// LocationInfoProject represents project information within a location.
type LocationInfoProject struct {
	ID        string                    `json:"id,required"`
	Directory string                    `json:"directory,required"`
	JSON      locationInfoProjectJSON   `json:"-"`
}

// locationInfoProjectJSON contains the JSON metadata for the struct [LocationInfoProject]
type locationInfoProjectJSON struct {
	ID          apijson.Field
	Directory   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationInfoProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationInfoProjectJSON) RawJSON() string {
	return r.raw
}

// V2SkillListParams contains the query parameters for listing skills.
type V2SkillListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2SkillListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
