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

// Generate a short name for a project copy from task context.
func (r *V2ProjectCopyService) GenerateName(ctx context.Context, projectID string, params V2ProjectCopyGenerateNameParams, opts ...option.RequestOption) (res *V2ProjectCopyGenerateNameResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectID parameter")
		return
	}
	path := fmt.Sprintf("experimental/project/%s/copy/generate-name", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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

// V2ProjectCopyGenerateNameResponse contains the response from generating a project copy name.
type V2ProjectCopyGenerateNameResponse struct {
	Name string                                `json:"name,required"`
	JSON v2ProjectCopyGenerateNameResponseJSON `json:"-"`
}

// v2ProjectCopyGenerateNameResponseJSON contains the JSON metadata for the struct [V2ProjectCopyGenerateNameResponse].
type v2ProjectCopyGenerateNameResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ProjectCopyGenerateNameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ProjectCopyGenerateNameResponseJSON) RawJSON() string {
	return r.raw
}

// V2ProjectCopyCreateParams contains the request parameters for creating a project copy.
type V2ProjectCopyCreateParams struct {
	// Location query params
	Location param.Field[V2LocationParam] `query:"location"`
	// Strategy for the copy operation
	Strategy param.Field[string] `json:"strategy"`
	// Directory for the copy
	Directory param.Field[string] `json:"directory"`
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

// V2ProjectCopyGenerateNameParams contains the request parameters for generating a project copy name.
type V2ProjectCopyGenerateNameParams struct {
	// Directory for the copy
	Directory param.Field[string] `query:"directory"`
	// Workspace for the copy
	Workspace param.Field[string] `query:"workspace"`
	// Context for generating the name
	Context param.Field[string] `json:"context"`
}

func (r V2ProjectCopyGenerateNameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V2ProjectCopyGenerateNameParams]'s query parameters as `url.Values`.
func (r V2ProjectCopyGenerateNameParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
