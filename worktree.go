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

// WorktreeService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorktreeService] method instead.
type WorktreeService struct {
	Options []option.RequestOption
}

// NewWorktreeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorktreeService(opts ...option.RequestOption) (r *WorktreeService) {
	r = &WorktreeService{}
	r.Options = opts
	return
}

// List worktrees
func (r *WorktreeService) List(ctx context.Context, query WorktreeListParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/worktree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create a worktree
func (r *WorktreeService) Create(ctx context.Context, params WorktreeCreateParams, opts ...option.RequestOption) (res *Worktree, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/worktree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Remove a worktree
func (r *WorktreeService) Remove(ctx context.Context, params WorktreeRemoveParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/worktree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Reset a worktree
func (r *WorktreeService) Reset(ctx context.Context, params WorktreeResetParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/worktree/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type Worktree struct {
	Name      string       `json:"name,required"`
	Branch    string       `json:"branch"`
	Directory string       `json:"directory,required"`
	JSON      worktreeJSON `json:"-"`
}

// worktreeJSON contains the JSON metadata for the struct [Worktree]
type worktreeJSON struct {
	Name        apijson.Field
	Branch      apijson.Field
	Directory   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Worktree) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r worktreeJSON) RawJSON() string {
	return r.raw
}

type WorktreeCreateParams struct {
	Name         param.Field[string] `json:"name"`
	StartCommand param.Field[string] `json:"startCommand"`
	Directory    param.Field[string] `query:"directory"`
	Workspace    param.Field[string] `query:"workspace"`
}

func (r WorktreeCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [WorktreeCreateParams]'s query parameters as `url.Values`.
func (r WorktreeCreateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorktreeRemoveParams struct {
	Directory param.Field[string] `json:"directory,required" query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorktreeRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [WorktreeRemoveParams]'s query parameters as `url.Values`.
func (r WorktreeRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorktreeResetParams struct {
	Directory param.Field[string] `json:"directory,required" query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r WorktreeResetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [WorktreeResetParams]'s query parameters as `url.Values`.
func (r WorktreeResetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorktreeListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [WorktreeListParams]'s query parameters as `url.Values`.
func (r WorktreeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
