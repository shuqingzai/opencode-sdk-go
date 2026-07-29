// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// V2FsService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2FsService] method instead.
type V2FsService struct {
	Options []option.RequestOption
}

// NewV2FsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2FsService(opts ...option.RequestOption) (r *V2FsService) {
	r = &V2FsService{}
	r.Options = opts
	return
}

// List files and directories in a given path
func (r *V2FsService) List(ctx context.Context, query V2FsListParams, opts ...option.RequestOption) (res *V2FsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/fs/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Find files matching a query
func (r *V2FsService) Find(ctx context.Context, query V2FsFindParams, opts ...option.RequestOption) (res *V2FsFindResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/fs/find"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Read a file at the given path.
//
// The path argument is the file path to read, which will be appended to
// the `/api/fs/read/` URL. Each path segment is percent-encoded with
// [url.PathEscape] so that spaces, '#', '?', and other special characters
// are transmitted correctly; '/' segment separators are preserved so that
// multi-level paths (e.g. "src/main.go") match the wildcard route
// `/api/fs/read/*` on the server side.
func (r *V2FsService) Read(ctx context.Context, path string, query V2FsReadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	segs := strings.Split(path, "/")
	for i, s := range segs {
		segs[i] = url.PathEscape(s)
	}
	urlPath := fmt.Sprintf("api/fs/read/%s", strings.Join(segs, "/"))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, urlPath, query, &res, opts...)
	return
}

// V2FsListParams contains the query parameters for listing files.
type V2FsListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
	Path     param.Field[string]          `query:"path"`
}

func (r V2FsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2FsFindParams contains the query parameters for finding files.
type V2FsFindParams struct {
	Location param.Field[V2LocationParam]     `query:"location"`
	Query    param.Field[string]              `query:"query,required"`
	Type     param.Field[FileSystemEntryType] `query:"type"`
	Limit    param.Field[string]              `query:"limit"`
}

func (r V2FsFindParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2FsReadParams contains the query parameters for reading a file.
type V2FsReadParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2FsReadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2FsListResponse represents the response from the v2.fs.list endpoint.
type V2FsListResponse struct {
	Location LocationInfo         `json:"location,required"`
	Data     []FileSystemEntry    `json:"data,required"`
	JSON     v2FsListResponseJSON `json:"-"`
}

// v2FsListResponseJSON contains the JSON metadata for the struct [V2FsListResponse]
type v2FsListResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2FsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2FsListResponseJSON) RawJSON() string {
	return r.raw
}

// V2FsFindResponse represents the response from the v2.fs.find endpoint.
type V2FsFindResponse struct {
	Location LocationInfo         `json:"location,required"`
	Data     []FileSystemEntry    `json:"data,required"`
	JSON     v2FsFindResponseJSON `json:"-"`
}

// v2FsFindResponseJSON contains the JSON metadata for the struct [V2FsFindResponse]
type v2FsFindResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2FsFindResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2FsFindResponseJSON) RawJSON() string {
	return r.raw
}

// FileSystemEntry represents a single file or directory entry in a filesystem
// listing.
type FileSystemEntry struct {
	Path string              `json:"path,required"`
	Type FileSystemEntryType `json:"type,required"`
	JSON fileSystemEntryJSON `json:"-"`
}

// fileSystemEntryJSON contains the JSON metadata for the struct [FileSystemEntry]
type fileSystemEntryJSON struct {
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileSystemEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileSystemEntryJSON) RawJSON() string {
	return r.raw
}

// FileSystemEntryType represents the type of a filesystem entry.
type FileSystemEntryType string

const (
	FileSystemEntryTypeFile      FileSystemEntryType = "file"
	FileSystemEntryTypeDirectory FileSystemEntryType = "directory"
)

func (r FileSystemEntryType) IsKnown() bool {
	switch r {
	case FileSystemEntryTypeFile, FileSystemEntryTypeDirectory:
		return true
	}
	return false
}
