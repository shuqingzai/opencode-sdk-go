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
	"github.com/sst/opencode-sdk-go/packages/ssestream"
)

// PtyService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPtyService] method instead.
type PtyService struct {
	Options []option.RequestOption
}

// NewPtyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPtyService(opts ...option.RequestOption) (r *PtyService) {
	r = &PtyService{}
	r.Options = opts
	return
}

// List all PTYs
func (r *PtyService) List(ctx context.Context, query PtyListParams, opts ...option.RequestOption) (res *[]Pty, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pty"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List available PTY shells
func (r *PtyService) Shells(ctx context.Context, query PtyShellsParams, opts ...option.RequestOption) (res *[]PtyShell, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pty/shells"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// New PTY
//
// Create a new PTY.
func (r *PtyService) New(ctx context.Context, params PtyNewParams, opts ...option.RequestOption) (res *Pty, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pty"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a PTY
func (r *PtyService) Get(ctx context.Context, ptyID string, query PtyGetParams, opts ...option.RequestOption) (res *Pty, err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a PTY
func (r *PtyService) Update(ctx context.Context, ptyID string, params PtyUpdateParams, opts ...option.RequestOption) (res *Pty, err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Remove a PTY
func (r *PtyService) Remove(ctx context.Context, ptyID string, query PtyRemoveParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, &res, opts...)
	return
}

// Connect to a PTY
//
// Establish a WebSocket connection streaming PTY output and accepting terminal input.
func (r *PtyService) Connect(ctx context.Context, ptyID string, query PtyConnectParams, opts ...option.RequestOption) (stream *ssestream.Stream[PtyEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		return ssestream.NewStream[PtyEvent](ssestream.NewDecoder(raw), errors.New("missing required ptyID parameter"))
	}
	path := fmt.Sprintf("pty/%s/connect", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[PtyEvent](ssestream.NewDecoder(raw), err)
}

// Get a connect token for a PTY
func (r *PtyService) ConnectToken(ctx context.Context, ptyID string, query PtyConnectTokenParams, opts ...option.RequestOption) (res *PtyConnectTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("pty/%s/connect-token", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

type PtyNewParams struct {
	Command   param.Field[string]            `json:"command"`
	Args      param.Field[[]string]          `json:"args"`
	Cwd       param.Field[string]            `json:"cwd"`
	Title     param.Field[string]            `json:"title"`
	Env       param.Field[map[string]string] `json:"env"`
	Directory param.Field[string]            `query:"directory"`
	Workspace param.Field[string]            `query:"workspace"`
}

func (r PtyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [PtyNewParams]'s query parameters as `url.Values`.
func (r PtyNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [PtyGetParams]'s query parameters as `url.Values`.
func (r PtyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyUpdateParams struct {
	Title     param.Field[string]  `json:"title"`
	Size      param.Field[PtySize] `json:"size"`
	Directory param.Field[string]  `query:"directory"`
	Workspace param.Field[string]  `query:"workspace"`
}

func (r PtyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [PtyUpdateParams]'s query parameters as `url.Values`.
func (r PtyUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtySize struct {
	Rows int64       `json:"rows,required"`
	Cols int64       `json:"cols,required"`
	JSON ptySizeJSON `json:"-"`
}

// ptySizeJSON contains the JSON metadata for the struct [PtySize]
type ptySizeJSON struct {
	Rows        apijson.Field
	Cols        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtySize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptySizeJSON) RawJSON() string {
	return r.raw
}

type PtyRemoveParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [PtyRemoveParams]'s query parameters as `url.Values`.
func (r PtyRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [PtyListParams]'s query parameters as `url.Values`.
func (r PtyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyConnectParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Cursor    param.Field[string] `query:"cursor"`
	Ticket    param.Field[string] `query:"ticket"`
}

// URLQuery serializes [PtyConnectParams]'s query parameters as `url.Values`.
func (r PtyConnectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyShellsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r PtyShellsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyShell struct {
	Path       string       `json:"path,required"`
	Name       string       `json:"name,required"`
	Acceptable bool         `json:"acceptable,required"`
	JSON       ptyShellJSON `json:"-"`
}

type ptyShellJSON struct {
	Path        apijson.Field
	Name        apijson.Field
	Acceptable  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyShell) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyShellJSON) RawJSON() string {
	return r.raw
}

type PtyConnectTokenParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r PtyConnectTokenParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PtyConnectTokenResponse struct {
	Ticket    string                      `json:"ticket,required"`
	ExpiresIn int64                       `json:"expires_in,required"`
	JSON      ptyConnectTokenResponseJSON `json:"-"`
}

type ptyConnectTokenResponseJSON struct {
	Ticket      apijson.Field
	ExpiresIn   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyConnectTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyConnectTokenResponseJSON) RawJSON() string {
	return r.raw
}
