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
func (r *V2PtyService) List(ctx context.Context, query V2PtyListParams, opts ...option.RequestOption) (res *[]V2PtyInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/pty"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create a v2 PTY
func (r *V2PtyService) Create(ctx context.Context, params V2PtyCreateParams, opts ...option.RequestOption) (res *V2PtyInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/pty"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a v2 PTY
func (r *V2PtyService) Get(ctx context.Context, ptyID string, query V2PtyGetParams, opts ...option.RequestOption) (res *V2PtyInfo, err error) {
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
func (r *V2PtyService) Update(ctx context.Context, ptyID string, params V2PtyUpdateParams, opts ...option.RequestOption) (res *V2PtyInfo, err error) {
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
func (r *V2PtyService) Remove(ctx context.Context, ptyID string, query V2PtyRemoveParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		err = errors.New("missing required ptyID parameter")
		return
	}
	path := fmt.Sprintf("api/pty/%s", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, &res, opts...)
	return
}

// Connect to a v2 PTY session
//
// Establish a WebSocket connection streaming PTY output and accepting terminal input.
func (r *V2PtyService) Connect(ctx context.Context, ptyID string, query V2PtyConnectParams, opts ...option.RequestOption) (stream *ssestream.Stream[PtyEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	if ptyID == "" {
		return ssestream.NewStream[PtyEvent](ssestream.NewDecoder(raw), errors.New("missing required ptyID parameter"))
	}
	path := fmt.Sprintf("api/pty/%s/connect", ptyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[PtyEvent](ssestream.NewDecoder(raw), err)
}

// Create a short-lived ticket for opening a PTY WebSocket connection.
func (r *V2PtyService) ConnectToken(ctx context.Context, ptyID string, query V2PtyConnectTokenParams, opts ...option.RequestOption) (res *PtyConnectToken, err error) {
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

// V2PtyInfo represents a PTY session.
type V2PtyInfo struct {
	ID       string        `json:"id,required"`
	Title    string        `json:"title,required"`
	Command  string        `json:"command,required"`
	Args     []string      `json:"args,required"`
	Cwd      string        `json:"cwd,required"`
	Status   V2PtyStatus   `json:"status,required"`
	Pid      int64         `json:"pid,required"`
	ExitCode int64         `json:"exitCode"`
	JSON     v2PtyInfoJSON `json:"-"`
}

// v2PtyInfoJSON contains the JSON metadata for the struct [V2PtyInfo]
type v2PtyInfoJSON struct {
	ID          apijson.Field
	Title       apijson.Field
	Command     apijson.Field
	Args        apijson.Field
	Cwd         apijson.Field
	Status      apijson.Field
	Pid         apijson.Field
	ExitCode    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PtyInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PtyInfoJSON) RawJSON() string {
	return r.raw
}

// V2PtyStatus represents the status of a PTY session.
type V2PtyStatus string

const (
	V2PtyStatusRunning V2PtyStatus = "running"
	V2PtyStatusExited  V2PtyStatus = "exited"
)

func (r V2PtyStatus) IsKnown() bool {
	switch r {
	case V2PtyStatusRunning, V2PtyStatusExited:
		return true
	}
	return false
}

// PtyConnectToken represents a short-lived ticket for opening a PTY WebSocket connection.
type PtyConnectToken struct {
	Ticket    string                `json:"ticket,required"`
	ExpiresIn int64                 `json:"expires_in,required"`
	JSON      ptyConnectTokenJSON   `json:"-"`
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

// PtyEvent represents an event received from a PTY WebSocket connection.
//
// This field can have the runtime type of [PtyCreatedEvent], [PtyUpdatedEvent],
// [PtyExitedEvent], [PtyDeletedEvent].
type PtyEvent struct {
	ID       string      `json:"id,required"`
	Type     string      `json:"type,required"`
	Metadata interface{} `json:"metadata"`
	Durable  PtyDurable  `json:"durable"`
	Location PtyLocation `json:"location"`
	Data     interface{} `json:"data,required"`
	JSON     ptyEventJSON `json:"-"`
}

// ptyEventJSON contains the JSON metadata for the struct [PtyEvent]
type ptyEventJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Metadata    apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyEventJSON) RawJSON() string {
	return r.raw
}

// PtyDurable represents the durable identifier metadata for a PTY event.
type PtyDurable struct {
	AggregateID string        `json:"aggregateID,required"`
	Seq         int64         `json:"seq,required"`
	Version     int64         `json:"version,required"`
	JSON        ptyDurableJSON `json:"-"`
}

// ptyDurableJSON contains the JSON metadata for the struct [PtyDurable]
type ptyDurableJSON struct {
	AggregateID apijson.Field
	Seq         apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyDurable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyDurableJSON) RawJSON() string {
	return r.raw
}

// PtyLocation represents the location context for a PTY event.
type PtyLocation struct {
	Directory   string           `json:"directory,required"`
	WorkspaceID string           `json:"workspaceID"`
	JSON        ptyLocationJSON  `json:"-"`
}

// ptyLocationJSON contains the JSON metadata for the struct [PtyLocation]
type ptyLocationJSON struct {
	Directory   apijson.Field
	WorkspaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyLocationJSON) RawJSON() string {
	return r.raw
}

// ===== Param Types =====

// V2PtyListParams contains the query parameters for listing v2 PTYs.
type V2PtyListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2PtyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2PtyCreateParams contains the body and query parameters for creating a v2 PTY.
type V2PtyCreateParams struct {
	Command  param.Field[string]            `json:"command"`
	Args     param.Field[[]string]          `json:"args"`
	Cwd      param.Field[string]            `json:"cwd"`
	Title    param.Field[string]            `json:"title"`
	Env      param.Field[map[string]string] `json:"env"`
	Location param.Field[V2LocationParam]   `query:"location"`
}

func (r V2PtyCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V2PtyCreateParams]'s query parameters as `url.Values`.
func (r V2PtyCreateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// V2PtyGetParams contains the query parameters for getting a v2 PTY.
type V2PtyGetParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

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

// V2PtySize represents the terminal size for a PTY.
type V2PtySize struct {
	Rows int64          `json:"rows,required"`
	Cols int64          `json:"cols,required"`
	JSON v2PtySizeJSON  `json:"-"`
}

// v2PtySizeJSON contains the JSON metadata for the struct [V2PtySize]
type v2PtySizeJSON struct {
	Rows        apijson.Field
	Cols        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2PtySize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2PtySizeJSON) RawJSON() string {
	return r.raw
}

// V2PtyRemoveParams contains the query parameters for removing a v2 PTY.
type V2PtyRemoveParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

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

func (r V2PtyConnectTokenParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
