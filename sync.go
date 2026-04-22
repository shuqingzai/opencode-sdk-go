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

// SyncService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSyncService] method instead.
type SyncService struct {
	Options  []option.RequestOption
	History *SyncHistoryService
}

// NewSyncService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSyncService(opts ...option.RequestOption) (r *SyncService) {
	r = &SyncService{}
	r.Options = opts
	r.History = NewSyncHistoryService(opts...)
	return
}

// Start a sync
func (r *SyncService) Start(ctx context.Context, query SyncStartParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

// Replay sync events
func (r *SyncService) Replay(ctx context.Context, body SyncReplayInput, opts ...option.RequestOption) (res *SyncReplayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/replay"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// SyncHistoryService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSyncHistoryService] method instead.
type SyncHistoryService struct {
	Options []option.RequestOption
}

// NewSyncHistoryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSyncHistoryService(opts ...option.RequestOption) (r *SyncHistoryService) {
	r = &SyncHistoryService{}
	r.Options = opts
	return
}

// List sync history
func (r *SyncHistoryService) List(ctx context.Context, body SyncHistoryListInput, opts ...option.RequestOption) (res *[]SyncHistoryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SyncStartParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r SyncStartParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SyncReplayInput struct {
	Directory param.Field[string]             `json:"directory,required"`
	Events   param.Field[[]SyncReplayEvent] `json:"events,required"`
	JSON     syncReplayInputJSON            `json:"-"`
}

type syncReplayInputJSON struct {
	Directory   apijson.Field
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncReplayInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncReplayInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r syncReplayInputJSON) RawJSON() string {
	return r.raw
}

type SyncReplayEvent struct {
	ID          string                 `json:"id,required"`
	AggregateID string                `json:"aggregateID,required"`
	Seq        float64               `json:"seq,required"`
	Type       string                `json:"type,required"`
	Data       map[string]interface{} `json:"data,required"`
	JSON       syncReplayEventJSON   `json:"-"`
}

type syncReplayEventJSON struct {
	ID          apijson.Field
	AggregateID apijson.Field
	Seq        apijson.Field
	Type       apijson.Field
	Data       apijson.Field
	raw        string
	ExtraFields map[string]apijson.Field
}

func (r *SyncReplayEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncReplayEventJSON) RawJSON() string {
	return r.raw
}

type SyncReplayResponse struct {
	SessionID string `json:"sessionID,omitempty"`
	JSON     syncReplayResponseJSON `json:"-"`
}

type syncReplayResponseJSON struct {
	SessionID  apijson.Field
	raw        string
	ExtraFields map[string]apijson.Field
}

func (r *SyncReplayResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncReplayResponseJSON) RawJSON() string {
	return r.raw
}

type SyncHistoryListInput struct {
	Body param.Field[map[string]float64] `json:"body"`
}

func (r SyncHistoryListInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SyncHistoryListResponse struct {
	ID         string                 `json:"id,required"`
	AggregateID string                `json:"aggregate_id,required"`
	Seq        float64               `json:"seq,required"`
	Type       string                `json:"type,required"`
	Data       map[string]interface{} `json:"data,required"`
	JSON       syncHistoryListResponseJSON `json:"-"`
}

type syncHistoryListResponseJSON struct {
	ID          apijson.Field
	AggregateID apijson.Field
	Seq        apijson.Field
	Type       apijson.Field
	Data       apijson.Field
	raw        string
	ExtraFields map[string]apijson.Field
}

func (r *SyncHistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncHistoryListResponseJSON) RawJSON() string {
	return r.raw
}