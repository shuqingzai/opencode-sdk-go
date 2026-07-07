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

// V2ProjectCopyService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2ProjectCopyService] method instead.
type V2ProjectCopyService struct {
	Options []option.RequestOption
}

// NewV2ProjectCopyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2ProjectCopyService(opts ...option.RequestOption) (r *V2ProjectCopyService) {
	r = &V2ProjectCopyService{}
	r.Options = opts
	return
}

// Create a project copy
func (r *V2ProjectCopyService) Create(ctx context.Context, projectID string, body V2ProjectCopyCreateParams, opts ...option.RequestOption) (res *ProjectCopyCopy, err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectID parameter")
		return
	}
	path := fmt.Sprintf("experimental/project/%s/copy", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Refresh a project copy
func (r *V2ProjectCopyService) Refresh(ctx context.Context, projectID string, query V2ProjectCopyRefreshParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectID parameter")
		return
	}
	path := fmt.Sprintf("experimental/project/%s/copy/refresh", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, nil, opts...)
	return
}

// Remove a project copy
func (r *V2ProjectCopyService) Remove(ctx context.Context, projectID string, body V2ProjectCopyRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectID parameter")
		return
	}
	path := fmt.Sprintf("experimental/project/%s/copy", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// ProjectCopyCopy represents the response from creating a project copy.
type ProjectCopyCopy struct {
	Directory string              `json:"directory,required"`
	JSON      projectCopyCopyJSON `json:"-"`
}

// projectCopyCopyJSON contains the JSON metadata for the struct [ProjectCopyCopy]
type projectCopyCopyJSON struct {
	Directory   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectCopyCopy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCopyCopyJSON) RawJSON() string {
	return r.raw
}

// V2ProjectCopyCreateParams contains the request parameters for creating a project copy.
type V2ProjectCopyCreateParams struct {
	// Location query params
	Location param.Field[V2LocationParam] `query:"location"`
	// Strategy for the copy operation
	Strategy  param.Field[string] `json:"strategy,required"`
	// Directory for the copy
	Directory param.Field[string] `json:"directory,required"`
	// Name for the copy
	Name param.Field[string] `json:"name"`
}

func (r V2ProjectCopyCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V2ProjectCopyCreateParams]'s query parameters as `url.Values`.
func (r V2ProjectCopyCreateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2ProjectCopyRefreshParams contains the query parameters for refreshing a project copy.
type V2ProjectCopyRefreshParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

// URLQuery serializes [V2ProjectCopyRefreshParams]'s query parameters as `url.Values`.
func (r V2ProjectCopyRefreshParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2ProjectCopyRemoveParams contains the request parameters for removing a project copy.
type V2ProjectCopyRemoveParams struct {
	// Location query params
	Location param.Field[V2LocationParam] `query:"location"`
	// Directory of the copy to remove
	Directory param.Field[string] `json:"directory,required"`
	// Force removal
	Force param.Field[bool] `json:"force,required"`
}

func (r V2ProjectCopyRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V2ProjectCopyRemoveParams]'s query parameters as `url.Values`.
func (r V2ProjectCopyRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
