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
	Options []option.RequestOption
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

// Start sync loops for workspaces in the current project that have active
// sessions.
func (r *SyncService) Start(ctx context.Context, query SyncStartParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

// Validate and replay a complete sync event history.
func (r *SyncService) Replay(ctx context.Context, params SyncReplayParams, opts ...option.RequestOption) (res *SyncReplayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/replay"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update a session to belong to the current workspace through the sync event
// system.
func (r *SyncService) Steal(ctx context.Context, params SyncStealParams, opts ...option.RequestOption) (res *SyncStealResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/steal"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// SyncHistoryService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSyncHistoryService] method instead.
type SyncHistoryService struct {
	Options []option.RequestOption
}

// NewSyncHistoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSyncHistoryService(opts ...option.RequestOption) (r *SyncHistoryService) {
	r = &SyncHistoryService{}
	r.Options = opts
	return
}

// List sync events for all aggregates. Keys are aggregate IDs the client already
// knows about, values are the last known sequence ID. Events with seq > value are
// returned for those aggregates. Aggregates not listed in the input get their full
// history.
func (r *SyncHistoryService) List(ctx context.Context, params SyncHistoryListParams, opts ...option.RequestOption) (res *[]SyncHistoryEvent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sync/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SyncReplayResponse struct {
	SessionID string                 `json:"sessionID,required"`
	JSON      syncReplayResponseJSON `json:"-"`
}

type syncReplayResponseJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncReplayResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncReplayResponseJSON) RawJSON() string {
	return r.raw
}

type SyncStealResponse struct {
	SessionID string                `json:"sessionID,required"`
	JSON      syncStealResponseJSON `json:"-"`
}

type syncStealResponseJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncStealResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncStealResponseJSON) RawJSON() string {
	return r.raw
}

type SyncHistoryEvent struct {
	ID          string               `json:"id,required"`
	AggregateID string               `json:"aggregate_id,required"`
	Seq         int64                `json:"seq,required"`
	Type        string               `json:"type,required"`
	Data        map[string]any       `json:"data,required"`
	JSON        syncHistoryEventJSON `json:"-"`
}

type syncHistoryEventJSON struct {
	ID          apijson.Field
	AggregateID apijson.Field
	Seq         apijson.Field
	Type        apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncHistoryEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncHistoryEventJSON) RawJSON() string {
	return r.raw
}

type SyncStartParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [SyncStartParams]'s query parameters as `url.Values`.
func (r SyncStartParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SyncReplayParams struct {
	Directory param.Field[string]               `query:"directory"`
	Workspace param.Field[string]               `query:"workspace"`
	Body      param.Field[SyncReplayParamsBody] `json:"-"`
}

func (r SyncReplayParams) MarshalJSON() (data []byte, err error) {
	if r.Body.Present {
		return apijson.MarshalRoot(r.Body)
	}
	return nil, nil
}

// URLQuery serializes [SyncReplayParams]'s query parameters as `url.Values`.
func (r SyncReplayParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SyncReplayParamsBody struct {
	Directory param.Field[string]                      `json:"directory,required"`
	Events    param.Field[[]SyncReplayParamsBodyEvent] `json:"events,required"`
}

func (r SyncReplayParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SyncReplayParamsBodyEvent struct {
	ID          param.Field[string]         `json:"id,required"`
	AggregateID param.Field[string]         `json:"aggregateID,required"`
	Seq         param.Field[int64]          `json:"seq,required"`
	Type        param.Field[string]         `json:"type,required"`
	Data        param.Field[map[string]any] `json:"data,required"`
}

func (r SyncReplayParamsBodyEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SyncStealParams struct {
	Directory param.Field[string]              `query:"directory"`
	Workspace param.Field[string]              `query:"workspace"`
	Body      param.Field[SyncStealParamsBody] `json:"-"`
}

func (r SyncStealParams) MarshalJSON() (data []byte, err error) {
	if r.Body.Present {
		return apijson.MarshalRoot(r.Body)
	}
	return nil, nil
}

// URLQuery serializes [SyncStealParams]'s query parameters as `url.Values`.
func (r SyncStealParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SyncStealParamsBody struct {
	SessionID param.Field[string] `json:"sessionID,required"`
}

func (r SyncStealParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SyncHistoryListParams struct {
	Directory param.Field[string]                    `query:"directory"`
	Workspace param.Field[string]                    `query:"workspace"`
	Body      param.Field[SyncHistoryListParamsBody] `json:"-"`
}

func (r SyncHistoryListParams) MarshalJSON() (data []byte, err error) {
	if r.Body.Present {
		return apijson.MarshalRoot(r.Body)
	}
	return nil, nil
}

// URLQuery serializes [SyncHistoryListParams]'s query parameters as `url.Values`.
func (r SyncHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SyncHistoryListParamsBody map[string]int64

func (r SyncHistoryListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
