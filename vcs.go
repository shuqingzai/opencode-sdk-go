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

// VcsService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVcsService] method instead.
type VcsService struct {
	Options []option.RequestOption
	Diff    *VcsDiffService
}

// NewVcsService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVcsService(opts ...option.RequestOption) (r *VcsService) {
	r = &VcsService{}
	r.Options = opts
	r.Diff = NewVcsDiffService(opts...)
	return
}

// VcsDiffService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVcsDiffService] method instead.
type VcsDiffService struct {
	Options []option.RequestOption
}

// NewVcsDiffService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVcsDiffService(opts ...option.RequestOption) (r *VcsDiffService) {
	r = &VcsDiffService{}
	r.Options = opts
	return
}

// Get diff between branches
func (r *VcsDiffService) Get(ctx context.Context, query VcsDiffParams, opts ...option.RequestOption) (res *[]VcsFileDiff, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vcs/diff"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get raw diff
func (r *VcsDiffService) Raw(ctx context.Context, query VcsDiffRawParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vcs/diff/raw"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *VcsService) Get(ctx context.Context, query VcsGetParams, opts ...option.RequestOption) (res *VcsInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type VcsInfo struct {
	Branch        string      `json:"branch"`
	DefaultBranch string      `json:"default_branch"`
	JSON          vcsInfoJSON `json:"-"`
}

type vcsInfoJSON struct {
	Branch        apijson.Field
	DefaultBranch apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VcsInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vcsInfoJSON) RawJSON() string {
	return r.raw
}

type VcsGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r VcsGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VcsDiffParams struct {
	Directory param.Field[string]            `query:"directory"`
	Workspace param.Field[string]            `query:"workspace"`
	Mode      param.Field[VcsDiffParamsMode] `query:"mode,required"`
	Context   param.Field[int64]             `query:"context"`
}

func (r VcsDiffParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VcsDiffParamsMode string

const (
	VcsDiffParamsModeGit    VcsDiffParamsMode = "git"
	VcsDiffParamsModeBranch VcsDiffParamsMode = "branch"
)

func (r VcsDiffParamsMode) IsKnown() bool {
	switch r {
	case VcsDiffParamsModeGit, VcsDiffParamsModeBranch:
		return true
	}
	return false
}

type VcsFileDiff struct {
	File      string            `json:"file,required"`
	Patch     string            `json:"patch"`
	Additions int64             `json:"additions,required"`
	Deletions int64             `json:"deletions,required"`
	Status    VcsFileDiffStatus `json:"status"`
	JSON      vcsFileDiffJSON   `json:"-"`
}

type vcsFileDiffJSON struct {
	File        apijson.Field
	Patch       apijson.Field
	Additions   apijson.Field
	Deletions   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VcsFileDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vcsFileDiffJSON) RawJSON() string {
	return r.raw
}

type VcsFileDiffStatus string

const (
	VcsFileDiffStatusAdded    VcsFileDiffStatus = "added"
	VcsFileDiffStatusDeleted  VcsFileDiffStatus = "deleted"
	VcsFileDiffStatusModified VcsFileDiffStatus = "modified"
)

func (r VcsFileDiffStatus) IsKnown() bool {
	switch r {
	case VcsFileDiffStatusAdded, VcsFileDiffStatusDeleted, VcsFileDiffStatusModified:
		return true
	}
	return false
}

func (r *VcsService) Status(ctx context.Context, query VcsStatusParams, opts ...option.RequestOption) (res *[]VcsFileStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vcs/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type VcsFileStatus struct {
	File      string              `json:"file,required"`
	Additions int64               `json:"additions,required"`
	Deletions int64               `json:"deletions,required"`
	Status    VcsFileStatusStatus `json:"status,required"`
	JSON      vcsFileStatusJSON   `json:"-"`
}

type vcsFileStatusJSON struct {
	File        apijson.Field
	Additions   apijson.Field
	Deletions   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VcsFileStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vcsFileStatusJSON) RawJSON() string {
	return r.raw
}

type VcsFileStatusStatus string

const (
	VcsFileStatusStatusAdded    VcsFileStatusStatus = "added"
	VcsFileStatusStatusDeleted  VcsFileStatusStatus = "deleted"
	VcsFileStatusStatusModified VcsFileStatusStatus = "modified"
)

func (r VcsFileStatusStatus) IsKnown() bool {
	switch r {
	case VcsFileStatusStatusAdded, VcsFileStatusStatusDeleted, VcsFileStatusStatusModified:
		return true
	}
	return false
}

type VcsStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r VcsStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VcsDiffRawParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r VcsDiffRawParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (r *VcsService) Apply(ctx context.Context, body VcsApplyParams, opts ...option.RequestOption) (res *VcsApplyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vcs/apply"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VcsApplyParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Patch     param.Field[string] `json:"patch,required"`
}

func (r VcsApplyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func (r VcsApplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VcsApplyResponse struct {
	Applied bool                 `json:"applied,required"`
	JSON    vcsApplyResponseJSON `json:"-"`
}

type vcsApplyResponseJSON struct {
	Applied     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VcsApplyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vcsApplyResponseJSON) RawJSON() string {
	return r.raw
}
