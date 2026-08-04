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

// V2PtyService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2PtyService] method instead.
type V2PtyService struct {
	Options []option.RequestOption
}

// NewV2PtyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2PtyService(opts ...option.RequestOption) (r *V2PtyService) {
	r = &V2PtyService{}
	r.Options = opts
	return
}

// List v2 PTYs
func (r *V2PtyService) List(ctx context.Context, query V2PtyListParams, opts ...option.RequestOption) (res *V2PtyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/pty"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// New v2 PTY
//
// Create a v2 PTY.
func (r *V2PtyService) New(ctx context.Context, params V2PtyNewParams, opts ...option.RequestOption) (res *V2PtyCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/pty"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a v2 PTY
func (r *V2PtyService) Get(ctx context.Context, ptyID string, query V2PtyGetParams, opts ...option.RequestOption) (res *V2PtyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("api/pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a v2 PTY
func (r *V2PtyService) Update(ctx context.Context, ptyID string, params V2PtyUpdateParams, opts ...option.RequestOption) (res *V2PtyUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("api/pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Remove a v2 PTY
func (r *V2PtyService) Remove(ctx context.Context, ptyID string, query V2PtyRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("api/pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, nil, opts...)
	return
}

// Connect to a v2 PTY session
//
// Establish a WebSocket connection streaming PTY output and accepting terminal input.
func (r *V2PtyService) Connect(ctx context.Context, ptyID string, query V2PtyConnectParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("api/pty/%s/connect", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create a short-lived ticket for opening a PTY WebSocket connection.
func (r *V2PtyService) ConnectToken(ctx context.Context, ptyID string, query V2PtyConnectTokenParams, opts ...option.RequestOption) (res *V2PtyConnectTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("api/pty/%s/connect-token", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

// ===== Response Types =====

// PtyConnectToken represents a short-lived ticket for opening a PTY WebSocket connection.
type PtyConnectToken struct {
	Ticket    string              `json:"ticket,required"`
	ExpiresIn int64               `json:"expires_in,required"`
	JSON      ptyConnectTokenJSON `json:"-"`
}

// ptyConnectTokenJSON contains the JSON metadata for the struct [PtyConnectToken]
type ptyConnectTokenJSON struct {
	Ticket      apijson.Field
	ExpiresIn   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyConnectToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyConnectTokenJSON) RawJSON() string {
	return r.raw
}

// ===== Response Wrappers =====

// V2PtyListResponse is returned by the List method. It wraps PTY sessions in the
// OpenAPI {location, data} envelope.
type V2PtyListResponse struct {
	Location LocationInfo          `json:"location,required"`
	Data     []Pty                 `json:"data,required"`
	JSON     v2PtyListResponseJSON `json:"-"`
}

// v2PtyListResponseJSON contains the JSON metadata for the struct [V2PtyListResponse]
type v2PtyListResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PtyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PtyListResponseJSON) RawJSON() string {
	return r.raw
}

// V2PtyCreateResponse is returned by the Create method. It wraps a PTY session in
// the OpenAPI {location, data} envelope.
type V2PtyCreateResponse struct {
	Location LocationInfo            `json:"location,required"`
	Data     Pty                     `json:"data,required"`
	JSON     v2PtyCreateResponseJSON `json:"-"`
}

// v2PtyCreateResponseJSON contains the JSON metadata for the struct [V2PtyCreateResponse]
type v2PtyCreateResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PtyCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PtyCreateResponseJSON) RawJSON() string {
	return r.raw
}

// V2PtyGetResponse is returned by the Get method. It wraps a PTY session in the
// OpenAPI {location, data} envelope.
type V2PtyGetResponse struct {
	Location LocationInfo         `json:"location,required"`
	Data     Pty                  `json:"data,required"`
	JSON     v2PtyGetResponseJSON `json:"-"`
}

// v2PtyGetResponseJSON contains the JSON metadata for the struct [V2PtyGetResponse]
type v2PtyGetResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PtyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PtyGetResponseJSON) RawJSON() string {
	return r.raw
}

// V2PtyUpdateResponse is returned by the Update method. It wraps a PTY session in
// the OpenAPI {location, data} envelope.
type V2PtyUpdateResponse struct {
	Location LocationInfo            `json:"location,required"`
	Data     Pty                     `json:"data,required"`
	JSON     v2PtyUpdateResponseJSON `json:"-"`
}

// v2PtyUpdateResponseJSON contains the JSON metadata for the struct [V2PtyUpdateResponse]
type v2PtyUpdateResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PtyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PtyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// V2PtyConnectTokenResponse is returned by the ConnectToken method. It wraps the
// connect token in the OpenAPI {location, data} envelope.
type V2PtyConnectTokenResponse struct {
	Location LocationInfo                  `json:"location,required"`
	Data     PtyConnectToken               `json:"data,required"`
	JSON     v2PtyConnectTokenResponseJSON `json:"-"`
}

// v2PtyConnectTokenResponseJSON contains the JSON metadata for the struct [V2PtyConnectTokenResponse]
type v2PtyConnectTokenResponseJSON struct {
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PtyConnectTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PtyConnectTokenResponseJSON) RawJSON() string {
	return r.raw
}

// ===== Param Types =====

// V2PtyListParams contains the query parameters for listing v2 PTYs.
type V2PtyListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

// URLQuery serializes [V2PtyListParams]'s query parameters as `url.Values`.
func (r V2PtyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2PtyNewParams contains the body and query parameters for creating a v2 PTY.
type V2PtyNewParams struct {
	Command  param.Field[string]            `json:"command"`
	Args     param.Field[[]string]          `json:"args"`
	Cwd      param.Field[string]            `json:"cwd"`
	Title    param.Field[string]            `json:"title"`
	Env      param.Field[map[string]string] `json:"env"`
	Location param.Field[V2LocationParam]   `query:"location"`
}

// MarshalJSON serializes [V2PtyNewParams] into JSON.
func (r V2PtyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V2PtyNewParams]'s query parameters as `url.Values`.
func (r V2PtyNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2PtyGetParams contains the query parameters for getting a v2 PTY.
type V2PtyGetParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

// URLQuery serializes [V2PtyGetParams]'s query parameters as `url.Values`.
func (r V2PtyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2PtyUpdateParams contains the body and query parameters for updating a v2 PTY.
type V2PtyUpdateParams struct {
	Title    param.Field[string]          `json:"title"`
	Size     param.Field[V2PtySize]       `json:"size"`
	Location param.Field[V2LocationParam] `query:"location"`
}

// MarshalJSON serializes [V2PtyUpdateParams] into JSON.
func (r V2PtyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V2PtyUpdateParams]'s query parameters as `url.Values`.
func (r V2PtyUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2PtySize is an alias for [PtySize]. Both v1 and v2 endpoints share the same
// inline size schema (rows/cols integer), so a single definition is canonical.
type V2PtySize = PtySize

// V2PtyRemoveParams contains the query parameters for removing a v2 PTY.
type V2PtyRemoveParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

// URLQuery serializes [V2PtyRemoveParams]'s query parameters as `url.Values`.
func (r V2PtyRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2PtyConnectParams contains the query parameters for connecting to a v2 PTY.
type V2PtyConnectParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
	Cursor   param.Field[string]          `query:"cursor"`
	Ticket   param.Field[string]          `query:"ticket"`
}

// URLQuery serializes [V2PtyConnectParams]'s query parameters as `url.Values`.
func (r V2PtyConnectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2PtyConnectTokenParams contains the query parameters for getting a v2 PTY connect token.
type V2PtyConnectTokenParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

// URLQuery serializes [V2PtyConnectTokenParams]'s query parameters as `url.Values`.
func (r V2PtyConnectTokenParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
