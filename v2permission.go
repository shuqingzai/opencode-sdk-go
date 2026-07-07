// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// V2PermissionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2PermissionService] method instead.
type V2PermissionService struct {
	Options []option.RequestOption
	Request *V2PermissionRequestService
	Saved   *V2PermissionSavedService
}

// NewV2PermissionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2PermissionService(opts ...option.RequestOption) (r *V2PermissionService) {
	r = &V2PermissionService{}
	r.Options = opts
	r.Request = NewV2PermissionRequestService(opts...)
	r.Saved = NewV2PermissionSavedService(opts...)
	return
}

// V2PermissionRequestService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2PermissionRequestService] method instead.
type V2PermissionRequestService struct {
	Options []option.RequestOption
}

// NewV2PermissionRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2PermissionRequestService(opts ...option.RequestOption) (r *V2PermissionRequestService) {
	r = &V2PermissionRequestService{}
	r.Options = opts
	return
}

// List pending permission requests for a location.
func (r *V2PermissionRequestService) List(ctx context.Context, query V2PermissionRequestListParams, opts ...option.RequestOption) (res *V2PermissionRequestListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/permission/request"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// V2PermissionSavedService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2PermissionSavedService] method instead.
type V2PermissionSavedService struct {
	Options []option.RequestOption
}

// NewV2PermissionSavedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2PermissionSavedService(opts ...option.RequestOption) (r *V2PermissionSavedService) {
	r = &V2PermissionSavedService{}
	r.Options = opts
	return
}

// List saved permissions, optionally filtered by project.
func (r *V2PermissionSavedService) List(ctx context.Context, query V2PermissionSavedListParams, opts ...option.RequestOption) (res *V2PermissionSavedListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/permission/saved"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove a saved permission by ID.
func (r *V2PermissionSavedService) Remove(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/permission/saved/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// ===== Param Types =====

type V2PermissionRequestListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2PermissionRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2PermissionSavedListParams struct {
	ProjectID param.Field[string] `query:"projectID"`
}

func (r V2PermissionSavedListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ===== Response Types =====

// V2PermissionRequestListResponse is returned by the Request.List method.
type V2PermissionRequestListResponse struct {
	Location LocationInfo                            `json:"location,required"`
	Data     []PermissionV2Request                   `json:"data,required"`
	JSON     v2PermissionRequestListResponseJSON     `json:"-"`
}

// v2PermissionRequestListResponseJSON contains the JSON metadata for the struct [V2PermissionRequestListResponse]
type v2PermissionRequestListResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PermissionRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PermissionRequestListResponseJSON) RawJSON() string {
	return r.raw
}

// V2PermissionSavedListResponse is returned by the Saved.List method.
type V2PermissionSavedListResponse struct {
	Data []PermissionSavedInfo                      `json:"data,required"`
	JSON v2PermissionSavedListResponseJSON          `json:"-"`
}

// v2PermissionSavedListResponseJSON contains the JSON metadata for the struct [V2PermissionSavedListResponse]
type v2PermissionSavedListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PermissionSavedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PermissionSavedListResponseJSON) RawJSON() string {
	return r.raw
}

// PermissionV2Request represents a pending v2 permission request.
type PermissionV2Request struct {
	ID        string                   `json:"id,required"`
	SessionID string                   `json:"sessionID,required"`
	Action    string                   `json:"action,required"`
	Resources []string                 `json:"resources,required"`
	Save      []string                 `json:"save"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata  interface{}              `json:"metadata"`
	Source    PermissionV2Source       `json:"source"`
	JSON      permissionV2RequestJSON  `json:"-"`
}

// permissionV2RequestJSON contains the JSON metadata for the struct [PermissionV2Request]
type permissionV2RequestJSON struct {
	ID          apijson.Field
	SessionID   apijson.Field
	Action      apijson.Field
	Resources   apijson.Field
	Save        apijson.Field
	Metadata    apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionV2Request) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionV2RequestJSON) RawJSON() string {
	return r.raw
}

// PermissionV2Source represents the source tool of a permission v2 request.
type PermissionV2Source struct {
	Type      string                    `json:"type,required"`
	MessageID string                    `json:"messageID,required"`
	CallID    string                    `json:"callID,required"`
	JSON      permissionV2SourceJSON    `json:"-"`
}

// permissionV2SourceJSON contains the JSON metadata for the struct [PermissionV2Source]
type permissionV2SourceJSON struct {
	Type        apijson.Field
	MessageID   apijson.Field
	CallID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionV2Source) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionV2SourceJSON) RawJSON() string {
	return r.raw
}

// PermissionSavedInfo represents a saved permission entry.
type PermissionSavedInfo struct {
	ID        string                      `json:"id,required"`
	ProjectID string                      `json:"projectID,required"`
	Action    string                      `json:"action,required"`
	Resource  string                      `json:"resource,required"`
	JSON      permissionSavedInfoJSON     `json:"-"`
}

// permissionSavedInfoJSON contains the JSON metadata for the struct [PermissionSavedInfo]
type permissionSavedInfoJSON struct {
	ID          apijson.Field
	ProjectID   apijson.Field
	Action      apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionSavedInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionSavedInfoJSON) RawJSON() string {
	return r.raw
}
