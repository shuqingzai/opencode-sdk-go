// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

// Package pagination implements the page types used by the opencode API.
//
// The opencode server exposes three distinct pagination styles, and this package
// provides one page type (plus a matching auto-pager) for each of them:
//
//   - [CursorPage] wraps `{ "data": [...], "cursor": { "previous": ..., "next": ... } }`
//     bodies. Paging is bidirectional and driven by the opaque `cursor` query
//     parameter. Used by `GET /api/session` and `GET /api/session/{sessionID}/message`.
//   - [SeqPage] wraps `{ "data": [...], "hasMore": bool }` bodies. Paging is forward
//     only and driven by the `after` query parameter, whose value is the durable
//     sequence number of the last event on the page. Used by
//     `GET /api/session/{sessionID}/history`.
//   - [HeaderCursorPage] and [HeaderBeforePage] wrap bare JSON array bodies whose
//     next-page cursor is advertised through the `X-Next-Cursor` response header.
//     They differ only in the query parameter used to send the cursor back:
//     `cursor` for `GET /experimental/session`, `before` for
//     `GET /session/{sessionID}/message`.
package pagination

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// HeaderNextCursor is the response header the opencode server uses to advertise
// the cursor of the next page for endpoints that respond with a bare JSON array.
// The header is omitted entirely once the last page has been reached.
const HeaderNextCursor = "X-Next-Cursor"

// ErrPageUnbound is returned by GetNextPage and GetPreviousPage when the page was
// not produced by an API call (for example when it was built by unmarshalling a
// JSON fixture) and therefore cannot be advanced.
var ErrPageUnbound = errors.New("opencode: page is not bound to a request and cannot be advanced")

// PageCursor holds the opaque cursors returned alongside a [CursorPage]. Both
// cursors encode the query that produced the page, so they must be replayed
// as-is and must not be combined with the `order` query parameter.
type PageCursor struct {
	Previous string         `json:"previous"`
	Next     string         `json:"next"`
	JSON     pageCursorJSON `json:"-"`
}

// pageCursorJSON contains the JSON metadata for the struct [PageCursor]
type pageCursorJSON struct {
	Previous    apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageCursor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageCursorJSON) RawJSON() string {
	return r.raw
}

type CursorPage[T any] struct {
	Data   []T        `json:"data,required"`
	Cursor PageCursor `json:"cursor,required"`
	// JSON contains metadata for fields, check presence with [apijson.Field.IsMissing].
	JSON cursorPageJSON `json:"-"`
	cfg  *requestconfig.RequestConfig
	res  *http.Response
}

// cursorPageJSON contains the JSON metadata for the struct [CursorPage]
type cursorPageJSON struct {
	Data        apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

// Returns the unmodified JSON received from the API
func (r CursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
//
// The server keeps returning `cursor.next` for the final non-empty page, so the
// last call made while walking a list resolves to an empty page rather than to
// 'nil'. [CursorPageAutoPager] absorbs that extra round trip for you.
func (r *CursorPage[T]) GetNextPage() (res *CursorPage[T], err error) {
	if r == nil || len(r.Data) == 0 {
		return nil, nil
	}

	next := r.Cursor.Next
	if next == "" {
		return nil, nil
	}
	if r.cfg == nil {
		return nil, ErrPageUnbound
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	if cfg == nil {
		return nil, ErrPageUnbound
	}
	// The cursor already encodes the ordering of the sequence it walks, so order
	// is redundant here. Dropping it is also required rather than merely tidy:
	// GET /api/session/{sessionID}/message answers 400 InvalidCursorError when a
	// request carries both.
	err = cfg.Apply(option.WithQuery("cursor", next), option.WithQueryDel("order"))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// GetPreviousPage returns the previous page as defined by this pagination style.
// When there is no previous page, this function will return a 'nil' for the page
// value, but will not return an error
func (r *CursorPage[T]) GetPreviousPage() (res *CursorPage[T], err error) {
	if r == nil || len(r.Data) == 0 {
		return nil, nil
	}

	previous := r.Cursor.Previous
	if previous == "" {
		return nil, nil
	}
	if r.cfg == nil {
		return nil, ErrPageUnbound
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	if cfg == nil {
		return nil, ErrPageUnbound
	}
	// See GetNextPage for why order is dropped.
	err = cfg.Apply(option.WithQuery("cursor", previous), option.WithQueryDel("order"))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		return
	}
	// A caller that supplied its own option.WithResponseInto takes over the
	// destination the raw response is written to, so recover it from the config
	// rather than leaving the page without one.
	if res == nil && cfg != nil && cfg.ResponseInto != nil {
		res = *cfg.ResponseInto
	}
	r.cfg = cfg
	r.res = res
}

// CursorPageAutoPager iterates every item of every page. Like the other
// auto-pagers it is a single cursor over a sequence of requests and is not safe
// for concurrent use: drive Next from one goroutine.
type CursorPageAutoPager[T any] struct {
	page *CursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewCursorPageAutoPager[T any](page *CursorPage[T], err error) *CursorPageAutoPager[T] {
	return &CursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPageAutoPager[T]) Index() int {
	return r.run
}

type SeqPage[T any] struct {
	Data    []T  `json:"data,required"`
	HasMore bool `json:"hasMore,required"`
	// JSON contains metadata for fields, check presence with [apijson.Field.IsMissing].
	JSON seqPageJSON `json:"-"`
	cfg  *requestconfig.RequestConfig
	res  *http.Response
}

// seqPageJSON contains the JSON metadata for the struct [SeqPage]
type seqPageJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

// Returns the unmodified JSON received from the API
func (r SeqPage[T]) RawJSON() string { return r.JSON.raw }
func (r *SeqPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NextAfter returns the exclusive aggregate sequence number to pass as the
// `after` query parameter in order to read the page that follows this one. It
// reads `durable.seq` off the last event of the page and reports false when the
// page carries no advanceable sequence.
func (r SeqPage[T]) NextAfter() (int64, bool) {
	events := gjson.Get(r.JSON.raw, "data").Array()
	if len(events) == 0 {
		return 0, false
	}
	seq := events[len(events)-1].Get("durable.seq")
	if !seq.Exists() {
		return 0, false
	}
	return seq.Int(), true
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
//
// A session history is an append-only log, so a page that reports HasMore as
// false today may report true on a later read once new events are committed.
func (r *SeqPage[T]) GetNextPage() (res *SeqPage[T], err error) {
	if r == nil || len(r.Data) == 0 {
		return nil, nil
	}

	if !r.HasMore {
		return nil, nil
	}
	after, ok := r.NextAfter()
	if !ok {
		return nil, nil
	}
	if r.cfg == nil {
		return nil, ErrPageUnbound
	}
	// Sequence numbers must strictly increase for the read to make progress. If
	// the page did not advance past the `after` that produced it, stop instead of
	// requesting the same window forever.
	if sent := r.cfg.Request.URL.Query().Get("after"); sent != "" {
		if previous, convErr := strconv.ParseInt(sent, 10, 64); convErr == nil && after <= previous {
			return nil, nil
		}
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	if cfg == nil {
		return nil, ErrPageUnbound
	}
	err = cfg.Apply(option.WithQuery("after", strconv.FormatInt(after, 10)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SeqPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		return
	}
	// A caller that supplied its own option.WithResponseInto takes over the
	// destination the raw response is written to, so recover it from the config
	// rather than leaving the page without one.
	if res == nil && cfg != nil && cfg.ResponseInto != nil {
		res = *cfg.ResponseInto
	}
	r.cfg = cfg
	r.res = res
}

// SeqPageAutoPager iterates every item of every page. Like the other
// auto-pagers it is a single cursor over a sequence of requests and is not safe
// for concurrent use: drive Next from one goroutine.
type SeqPageAutoPager[T any] struct {
	page *SeqPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSeqPageAutoPager[T any](page *SeqPage[T], err error) *SeqPageAutoPager[T] {
	return &SeqPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SeqPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SeqPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SeqPageAutoPager[T]) Err() error {
	return r.err
}

func (r *SeqPageAutoPager[T]) Index() int {
	return r.run
}

// HeaderCursorPage pages a bare JSON array body. Data is tagged `json:"-"` because
// the payload is the array itself rather than an object with a data key, so it
// is filled in by UnmarshalJSON alone: re-marshalling a page does not round trip.
// A body that is neither an array nor null fails to decode.
type HeaderCursorPage[T any] struct {
	Data []T `json:"-"`
	// JSON contains metadata for fields, check presence with [apijson.Field.IsMissing].
	JSON headerCursorPageJSON `json:"-"`
	cfg  *requestconfig.RequestConfig
	res  *http.Response
}

// headerCursorPageJSON contains the JSON metadata for the struct
// [HeaderCursorPage]
type headerCursorPageJSON struct {
	raw string
}

// Returns the unmodified JSON received from the API
func (r HeaderCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *HeaderCursorPage[T]) UnmarshalJSON(data []byte) error {
	r.JSON.raw = string(data)
	return json.Unmarshal(data, &r.Data)
}

// NextCursor returns the value of the X-Next-Cursor response header, which the
// server only sends while further pages remain.
func (r HeaderCursorPage[T]) NextCursor() string {
	if r.res == nil {
		return ""
	}
	return r.res.Header.Get(HeaderNextCursor)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *HeaderCursorPage[T]) GetNextPage() (res *HeaderCursorPage[T], err error) {
	if r == nil || len(r.Data) == 0 {
		return nil, nil
	}

	next := r.NextCursor()
	if next == "" {
		return nil, nil
	}
	if r.cfg == nil {
		return nil, ErrPageUnbound
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	if cfg == nil {
		return nil, ErrPageUnbound
	}
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *HeaderCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		return
	}
	// A caller that supplied its own option.WithResponseInto takes over the
	// destination the raw response is written to, so recover it from the config
	// rather than leaving the page without one.
	if res == nil && cfg != nil && cfg.ResponseInto != nil {
		res = *cfg.ResponseInto
	}
	r.cfg = cfg
	r.res = res
}

// HeaderCursorPageAutoPager iterates every item of every page. Like the other
// auto-pagers it is a single cursor over a sequence of requests and is not safe
// for concurrent use: drive Next from one goroutine.
type HeaderCursorPageAutoPager[T any] struct {
	page *HeaderCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewHeaderCursorPageAutoPager[T any](page *HeaderCursorPage[T], err error) *HeaderCursorPageAutoPager[T] {
	return &HeaderCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *HeaderCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *HeaderCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *HeaderCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *HeaderCursorPageAutoPager[T]) Index() int {
	return r.run
}

// HeaderBeforePage pages a bare JSON array body. Data is tagged `json:"-"` because
// the payload is the array itself rather than an object with a data key, so it
// is filled in by UnmarshalJSON alone: re-marshalling a page does not round trip.
// A body that is neither an array nor null fails to decode.
type HeaderBeforePage[T any] struct {
	Data []T `json:"-"`
	// JSON contains metadata for fields, check presence with [apijson.Field.IsMissing].
	JSON headerBeforePageJSON `json:"-"`
	cfg  *requestconfig.RequestConfig
	res  *http.Response
}

// headerBeforePageJSON contains the JSON metadata for the struct
// [HeaderBeforePage]
type headerBeforePageJSON struct {
	raw string
}

// Returns the unmodified JSON received from the API
func (r HeaderBeforePage[T]) RawJSON() string { return r.JSON.raw }
func (r *HeaderBeforePage[T]) UnmarshalJSON(data []byte) error {
	r.JSON.raw = string(data)
	return json.Unmarshal(data, &r.Data)
}

// NextCursor returns the value of the X-Next-Cursor response header, which the
// server only sends while further pages remain.
func (r HeaderBeforePage[T]) NextCursor() string {
	if r.res == nil {
		return ""
	}
	return r.res.Header.Get(HeaderNextCursor)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *HeaderBeforePage[T]) GetNextPage() (res *HeaderBeforePage[T], err error) {
	if r == nil || len(r.Data) == 0 {
		return nil, nil
	}

	next := r.NextCursor()
	if next == "" {
		return nil, nil
	}
	if r.cfg == nil {
		return nil, ErrPageUnbound
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	if cfg == nil {
		return nil, ErrPageUnbound
	}
	err = cfg.Apply(option.WithQuery("before", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *HeaderBeforePage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		return
	}
	// A caller that supplied its own option.WithResponseInto takes over the
	// destination the raw response is written to, so recover it from the config
	// rather than leaving the page without one.
	if res == nil && cfg != nil && cfg.ResponseInto != nil {
		res = *cfg.ResponseInto
	}
	r.cfg = cfg
	r.res = res
}

// HeaderBeforePageAutoPager iterates every item of every page. Like the other
// auto-pagers it is a single cursor over a sequence of requests and is not safe
// for concurrent use: drive Next from one goroutine.
type HeaderBeforePageAutoPager[T any] struct {
	page *HeaderBeforePage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewHeaderBeforePageAutoPager[T any](page *HeaderBeforePage[T], err error) *HeaderBeforePageAutoPager[T] {
	return &HeaderBeforePageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *HeaderBeforePageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *HeaderBeforePageAutoPager[T]) Current() T {
	return r.cur
}

func (r *HeaderBeforePageAutoPager[T]) Err() error {
	return r.err
}

func (r *HeaderBeforePageAutoPager[T]) Index() int {
	return r.run
}
