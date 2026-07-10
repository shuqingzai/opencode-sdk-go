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

// FileService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// List files and directories
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *[]FileNode, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Read a file
func (r *FileService) Read(ctx context.Context, query FileReadParams, opts ...option.RequestOption) (res *FileContent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "file/content"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get file status
func (r *FileService) Status(ctx context.Context, query FileStatusParams, opts ...option.RequestOption) (res *[]File, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "file/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type File struct {
	Added   int64      `json:"added,required"`
	Path    string     `json:"path,required"`
	Removed int64      `json:"removed,required"`
	Status  FileStatus `json:"status,required"`
	JSON    fileJSON   `json:"-"`
}

// fileJSON contains the JSON metadata for the struct [File]
type fileJSON struct {
	Added       apijson.Field
	Path        apijson.Field
	Removed     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *File) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileJSON) RawJSON() string {
	return r.raw
}

type FileStatus string

const (
	FileStatusAdded    FileStatus = "added"
	FileStatusDeleted  FileStatus = "deleted"
	FileStatusModified FileStatus = "modified"
)

func (r FileStatus) IsKnown() bool {
	switch r {
	case FileStatusAdded, FileStatusDeleted, FileStatusModified:
		return true
	}
	return false
}

type FileNode struct {
	Absolute string       `json:"absolute,required"`
	Ignored  bool         `json:"ignored,required"`
	Name     string       `json:"name,required"`
	Path     string       `json:"path,required"`
	Type     FileNodeType `json:"type,required"`
	JSON     fileNodeJSON `json:"-"`
}

// fileNodeJSON contains the JSON metadata for the struct [FileNode]
type fileNodeJSON struct {
	Absolute    apijson.Field
	Ignored     apijson.Field
	Name        apijson.Field
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileNode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileNodeJSON) RawJSON() string {
	return r.raw
}

type FileNodeType string

const (
	FileNodeTypeFile      FileNodeType = "file"
	FileNodeTypeDirectory FileNodeType = "directory"
)

func (r FileNodeType) IsKnown() bool {
	switch r {
	case FileNodeTypeFile, FileNodeTypeDirectory:
		return true
	}
	return false
}

type FileContent struct {
	Content  string              `json:"content,required"`
	Type     FileContentType     `json:"type,required"`
	Diff     string              `json:"diff"`
	Encoding FileContentEncoding `json:"encoding"`
	MIMEType string              `json:"mimeType"`
	Patch    FileContentPatch    `json:"patch"`
	JSON     fileContentJSON     `json:"-"`
}

// fileContentJSON contains the JSON metadata for the struct
// [FileContent]
type fileContentJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	Diff        apijson.Field
	Encoding    apijson.Field
	MIMEType    apijson.Field
	Patch       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileContentJSON) RawJSON() string {
	return r.raw
}

type FileContentType string

const (
	FileContentTypeText   FileContentType = "text"
	FileContentTypeBinary FileContentType = "binary"
)

func (r FileContentType) IsKnown() bool {
	switch r {
	case FileContentTypeText, FileContentTypeBinary:
		return true
	}
	return false
}

type FileContentEncoding string

const (
	FileContentEncodingBase64 FileContentEncoding = "base64"
)

func (r FileContentEncoding) IsKnown() bool {
	switch r {
	case FileContentEncodingBase64:
		return true
	}
	return false
}

type FileContentPatch struct {
	Hunks       []FileContentPatchHunk `json:"hunks,required"`
	NewFileName string                 `json:"newFileName,required"`
	OldFileName string                 `json:"oldFileName,required"`
	Index       string                 `json:"index"`
	NewHeader   string                 `json:"newHeader"`
	OldHeader   string                 `json:"oldHeader"`
	JSON        fileContentPatchJSON   `json:"-"`
}

// fileContentPatchJSON contains the JSON metadata for the struct
// [FileContentPatch]
type fileContentPatchJSON struct {
	Hunks       apijson.Field
	NewFileName apijson.Field
	OldFileName apijson.Field
	Index       apijson.Field
	NewHeader   apijson.Field
	OldHeader   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileContentPatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileContentPatchJSON) RawJSON() string {
	return r.raw
}

type FileContentPatchHunk struct {
	Lines    []string                 `json:"lines,required"`
	NewLines int64                    `json:"newLines,required"`
	NewStart int64                    `json:"newStart,required"`
	OldLines int64                    `json:"oldLines,required"`
	OldStart int64                    `json:"oldStart,required"`
	JSON     fileContentPatchHunkJSON `json:"-"`
}

// fileContentPatchHunkJSON contains the JSON metadata for the struct
// [FileContentPatchHunk]
type fileContentPatchHunkJSON struct {
	Lines       apijson.Field
	NewLines    apijson.Field
	NewStart    apijson.Field
	OldLines    apijson.Field
	OldStart    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileContentPatchHunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileContentPatchHunkJSON) RawJSON() string {
	return r.raw
}

type FileListParams struct {
	Path      param.Field[string] `query:"path,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileReadParams struct {
	Path      param.Field[string] `query:"path,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [FileReadParams]'s query parameters as `url.Values`.
func (r FileReadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FileStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [FileStatusParams]'s query parameters as `url.Values`.
func (r FileStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
