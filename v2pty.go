// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
	"github.com/tidwall/gjson"
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
// Establish a connection streaming PTY output and accepting terminal input.
//
// Note: OpenAPI marks this endpoint with `x-websocket: true` (WebSocket upgrade),
// and the JSON response schema is `boolean`. This matches the behavior of the
// legacy `PtyService.Connect`. The SDK exposes it as an SSE-style stream; consumers
// should be aware that the underlying transport is WebSocket.
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

// PtyEvent represents an event received from a PTY WebSocket connection.
type PtyEvent struct {
	ID   string       `json:"id,required"`
	Type PtyEventType `json:"type,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{} `json:"metadata"`
	// This field can have the runtime type of [PtyDurable].
	Durable interface{} `json:"durable"`
	// This field can have the runtime type of [LocationRef].
	Location interface{} `json:"location"`
	// This field can have the runtime type of [PtyCreatedEventData], [PtyUpdatedEventData],
	// [PtyExitedEventData], [PtyDeletedEventData].
	Data  interface{}  `json:"data,required"`
	JSON  ptyEventJSON `json:"-"`
	union PtyEventUnion
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
	*r = PtyEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ptyEventJSON) RawJSON() string {
	return r.raw
}

// AsUnion returns a [PtyEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [PtyCreatedEvent], [PtyUpdatedEvent],
// [PtyExitedEvent], [PtyDeletedEvent].
func (r PtyEvent) AsUnion() PtyEventUnion {
	return r.union
}

// PtyEventType enumerates all possible event types for PtyEvent.
type PtyEventType string

const (
	PtyEventTypePtyCreated PtyEventType = "pty.created"
	PtyEventTypePtyUpdated PtyEventType = "pty.updated"
	PtyEventTypePtyExited  PtyEventType = "pty.exited"
	PtyEventTypePtyDeleted PtyEventType = "pty.deleted"
)

func (r PtyEventType) IsKnown() bool {
	switch r {
	case PtyEventTypePtyCreated,
		PtyEventTypePtyUpdated,
		PtyEventTypePtyExited,
		PtyEventTypePtyDeleted:
		return true
	}
	return false
}

// PtyEventUnion is satisfied by all PtyEvent variant types.
type PtyEventUnion interface {
	implementsPtyEvent()
}

// PtyDurable represents the durable identifier metadata for a PTY event.
type PtyDurable struct {
	AggregateID string         `json:"aggregateID,required"`
	Seq         int64          `json:"seq,required"`
	Version     int64          `json:"version,required"`
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
	Directory   string          `json:"directory,required"`
	WorkspaceID string          `json:"workspaceID"`
	JSON        ptyLocationJSON `json:"-"`
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

// PtyCreatedEvent represents a PtyEvent of type "pty.created".
type PtyCreatedEvent struct {
	ID   string              `json:"id,required"`
	Type PtyCreatedEventType `json:"type,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{} `json:"metadata"`
	// This field can have the runtime type of [PtyDurable].
	Durable interface{} `json:"durable"`
	// This field can have the runtime type of [LocationRef].
	Location interface{}         `json:"location"`
	Data     PtyCreatedEventData `json:"data,required"`
	JSON     ptyCreatedEventJSON `json:"-"`
}

// ptyCreatedEventJSON contains the JSON metadata for the struct [PtyCreatedEvent]
type ptyCreatedEventJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Metadata    apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyCreatedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyCreatedEventJSON) RawJSON() string {
	return r.raw
}

func (r PtyCreatedEvent) implementsPtyEvent() {}

// PtyCreatedEventType enumerates all possible event types for PtyCreatedEvent.
type PtyCreatedEventType string

const (
	PtyCreatedEventTypePtyCreated PtyCreatedEventType = "pty.created"
)

func (r PtyCreatedEventType) IsKnown() bool {
	switch r {
	case PtyCreatedEventTypePtyCreated:
		return true
	}
	return false
}

type PtyCreatedEventData struct {
	Info Pty                     `json:"info,required"`
	JSON ptyCreatedEventDataJSON `json:"-"`
}

// ptyCreatedEventDataJSON contains the JSON metadata for the struct [PtyCreatedEventData]
type ptyCreatedEventDataJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyCreatedEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyCreatedEventDataJSON) RawJSON() string {
	return r.raw
}

// PtyUpdatedEvent represents a PtyEvent of type "pty.updated".
type PtyUpdatedEvent struct {
	ID   string              `json:"id,required"`
	Type PtyUpdatedEventType `json:"type,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{} `json:"metadata"`
	// This field can have the runtime type of [PtyDurable].
	Durable interface{} `json:"durable"`
	// This field can have the runtime type of [LocationRef].
	Location interface{}         `json:"location"`
	Data     PtyUpdatedEventData `json:"data,required"`
	JSON     ptyUpdatedEventJSON `json:"-"`
}

// ptyUpdatedEventJSON contains the JSON metadata for the struct [PtyUpdatedEvent]
type ptyUpdatedEventJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Metadata    apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyUpdatedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyUpdatedEventJSON) RawJSON() string {
	return r.raw
}

func (r PtyUpdatedEvent) implementsPtyEvent() {}

// PtyUpdatedEventType enumerates all possible event types for PtyUpdatedEvent.
type PtyUpdatedEventType string

const (
	PtyUpdatedEventTypePtyUpdated PtyUpdatedEventType = "pty.updated"
)

func (r PtyUpdatedEventType) IsKnown() bool {
	switch r {
	case PtyUpdatedEventTypePtyUpdated:
		return true
	}
	return false
}

type PtyUpdatedEventData struct {
	Info Pty                     `json:"info,required"`
	JSON ptyUpdatedEventDataJSON `json:"-"`
}

// ptyUpdatedEventDataJSON contains the JSON metadata for the struct [PtyUpdatedEventData]
type ptyUpdatedEventDataJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyUpdatedEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyUpdatedEventDataJSON) RawJSON() string {
	return r.raw
}

// PtyExitedEvent represents a PtyEvent of type "pty.exited".
type PtyExitedEvent struct {
	ID   string             `json:"id,required"`
	Type PtyExitedEventType `json:"type,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{} `json:"metadata"`
	// This field can have the runtime type of [PtyDurable].
	Durable interface{} `json:"durable"`
	// This field can have the runtime type of [LocationRef].
	Location interface{}        `json:"location"`
	Data     PtyExitedEventData `json:"data,required"`
	JSON     ptyExitedEventJSON `json:"-"`
}

// ptyExitedEventJSON contains the JSON metadata for the struct [PtyExitedEvent]
type ptyExitedEventJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Metadata    apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyExitedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyExitedEventJSON) RawJSON() string {
	return r.raw
}

func (r PtyExitedEvent) implementsPtyEvent() {}

// PtyExitedEventType enumerates all possible event types for PtyExitedEvent.
type PtyExitedEventType string

const (
	PtyExitedEventTypePtyExited PtyExitedEventType = "pty.exited"
)

func (r PtyExitedEventType) IsKnown() bool {
	switch r {
	case PtyExitedEventTypePtyExited:
		return true
	}
	return false
}

type PtyExitedEventData struct {
	ExitCode int64                  `json:"exitCode,required"`
	ID       string                 `json:"id,required"`
	JSON     ptyExitedEventDataJSON `json:"-"`
}

// ptyExitedEventDataJSON contains the JSON metadata for the struct [PtyExitedEventData]
type ptyExitedEventDataJSON struct {
	ExitCode    apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyExitedEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyExitedEventDataJSON) RawJSON() string {
	return r.raw
}

// PtyDeletedEvent represents a PtyEvent of type "pty.deleted".
type PtyDeletedEvent struct {
	ID   string              `json:"id,required"`
	Type PtyDeletedEventType `json:"type,required"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata interface{} `json:"metadata"`
	// This field can have the runtime type of [PtyDurable].
	Durable interface{} `json:"durable"`
	// This field can have the runtime type of [LocationRef].
	Location interface{}         `json:"location"`
	Data     PtyDeletedEventData `json:"data,required"`
	JSON     ptyDeletedEventJSON `json:"-"`
}

// ptyDeletedEventJSON contains the JSON metadata for the struct [PtyDeletedEvent]
type ptyDeletedEventJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	Metadata    apijson.Field
	Durable     apijson.Field
	Location    apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyDeletedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyDeletedEventJSON) RawJSON() string {
	return r.raw
}

func (r PtyDeletedEvent) implementsPtyEvent() {}

// PtyDeletedEventType enumerates all possible event types for PtyDeletedEvent.
type PtyDeletedEventType string

const (
	PtyDeletedEventTypePtyDeleted PtyDeletedEventType = "pty.deleted"
)

func (r PtyDeletedEventType) IsKnown() bool {
	switch r {
	case PtyDeletedEventTypePtyDeleted:
		return true
	}
	return false
}

type PtyDeletedEventData struct {
	ID   string                  `json:"id,required"`
	JSON ptyDeletedEventDataJSON `json:"-"`
}

// ptyDeletedEventDataJSON contains the JSON metadata for the struct [PtyDeletedEventData]
type ptyDeletedEventDataJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PtyDeletedEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyDeletedEventDataJSON) RawJSON() string {
	return r.raw
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PtyEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PtyCreatedEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PtyUpdatedEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PtyExitedEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PtyDeletedEvent{}),
		},
	)
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
	Rows int64         `json:"rows,required"`
	Cols int64         `json:"cols,required"`
	JSON v2PtySizeJSON `json:"-"`
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
