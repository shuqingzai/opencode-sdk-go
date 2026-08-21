package opencode

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/sst/opencode-sdk-go/internal/apierror"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/pagination"
)

// recordingServer captures the query string of every request it serves so that
// tests can assert on how a page advanced itself.
type recordingServer struct {
	*httptest.Server
	queries []url.Values
	paths   []string
	headers []http.Header
}

func newRecordingServer(t *testing.T, handler func(w http.ResponseWriter, r *http.Request, hit int)) *recordingServer {
	t.Helper()
	rec := &recordingServer{}
	rec.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hit := len(rec.queries)
		rec.queries = append(rec.queries, r.URL.Query())
		rec.paths = append(rec.paths, r.URL.Path)
		rec.headers = append(rec.headers, r.Header.Clone())
		w.Header().Set("Content-Type", "application/json")
		handler(w, r, hit)
	}))
	t.Cleanup(rec.Close)
	return rec
}

func (rec *recordingServer) client() *Client {
	return NewClient(option.WithBaseURL(rec.URL + "/"))
}

func (rec *recordingServer) query(t *testing.T, hit int) url.Values {
	t.Helper()
	if hit >= len(rec.queries) {
		t.Fatalf("expected at least %d requests, server only saw %d", hit+1, len(rec.queries))
	}
	return rec.queries[hit]
}

// ===== CursorPage: GET /api/session =====

func TestV2SessionListPaginatesThroughBodyCursor(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			fmt.Fprint(w, `{"data":[{"id":"ses_1"},{"id":"ses_2"}],"cursor":{"previous":"prev_1","next":"next_1"}}`)
		case 1:
			fmt.Fprint(w, `{"data":[{"id":"ses_3"}],"cursor":{"previous":"prev_2","next":"next_2"}}`)
		default:
			fmt.Fprint(w, `{"data":[],"cursor":{}}`)
		}
	})

	client := rec.client()
	page, err := client.V2Session.List(context.Background(), V2SessionListParams{
		Limit: Int(2),
		Order: F(V2SessionOrderDesc),
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}

	if got, want := len(page.Data), 2; got != want {
		t.Fatalf("first page size = %d, want %d", got, want)
	}
	if got, want := page.Data[0].ID, "ses_1"; got != want {
		t.Errorf("first item = %q, want %q", got, want)
	}
	if got, want := page.Cursor.Next, "next_1"; got != want {
		t.Errorf("cursor.next = %q, want %q", got, want)
	}
	if got, want := page.Cursor.Previous, "prev_1"; got != want {
		t.Errorf("cursor.previous = %q, want %q", got, want)
	}
	if page.RawJSON() == "" {
		t.Error("RawJSON is empty, the raw body should be retained")
	}

	// The first request must carry the caller's params verbatim.
	first := rec.query(t, 0)
	if got, want := first.Get("limit"), "2"; got != want {
		t.Errorf("request 0 limit = %q, want %q", got, want)
	}
	if got, want := first.Get("order"), "desc"; got != want {
		t.Errorf("request 0 order = %q, want %q", got, want)
	}
	if first.Has("cursor") {
		t.Errorf("request 0 should not send a cursor, got %q", first.Get("cursor"))
	}

	next, err := page.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	if next == nil {
		t.Fatal("GetNextPage returned nil, want the second page")
	}
	if got, want := len(next.Data), 1; got != want {
		t.Fatalf("second page size = %d, want %d", got, want)
	}
	if got, want := next.Data[0].ID, "ses_3"; got != want {
		t.Errorf("second page item = %q, want %q", got, want)
	}

	// Advancing must replay cursor.next and must drop `order`: the cursor already
	// encodes the ordering and the server rejects requests that send both.
	second := rec.query(t, 1)
	if got, want := second.Get("cursor"), "next_1"; got != want {
		t.Errorf("request 1 cursor = %q, want %q", got, want)
	}
	if second.Has("order") {
		t.Errorf("request 1 must not send order, got %q", second.Get("order"))
	}
	if got, want := second.Get("limit"), "2"; got != want {
		t.Errorf("request 1 limit = %q, want %q (page size must be preserved)", got, want)
	}
	if got, want := rec.paths[1], "/api/session"; got != want {
		t.Errorf("request 1 path = %q, want %q", got, want)
	}
}

func TestV2SessionListGetPreviousPage(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		if hit == 0 {
			fmt.Fprint(w, `{"data":[{"id":"ses_5"}],"cursor":{"previous":"prev_1","next":"next_1"}}`)
			return
		}
		fmt.Fprint(w, `{"data":[{"id":"ses_4"}],"cursor":{"previous":"prev_0","next":"next_0"}}`)
	})

	client := rec.client()
	page, err := client.V2Session.List(context.Background(), V2SessionListParams{Order: F(V2SessionOrderAsc)})
	if err != nil {
		t.Fatalf("List: %v", err)
	}

	prev, err := page.GetPreviousPage()
	if err != nil {
		t.Fatalf("GetPreviousPage: %v", err)
	}
	if prev == nil {
		t.Fatal("GetPreviousPage returned nil, want the previous page")
	}
	if got, want := prev.Data[0].ID, "ses_4"; got != want {
		t.Errorf("previous page item = %q, want %q", got, want)
	}

	second := rec.query(t, 1)
	if got, want := second.Get("cursor"), "prev_1"; got != want {
		t.Errorf("request 1 cursor = %q, want %q", got, want)
	}
	if second.Has("order") {
		t.Errorf("request 1 must not send order, got %q", second.Get("order"))
	}
}

func TestV2SessionListAutoPagingWalksEveryPage(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			fmt.Fprint(w, `{"data":[{"id":"ses_1"},{"id":"ses_2"}],"cursor":{"next":"c1"}}`)
		case 1:
			fmt.Fprint(w, `{"data":[{"id":"ses_3"},{"id":"ses_4"}],"cursor":{"next":"c2"}}`)
		case 2:
			// The server keeps handing out a cursor for the last non-empty page, so
			// the walk terminates on the empty page that follows it.
			fmt.Fprint(w, `{"data":[{"id":"ses_5"}],"cursor":{"next":"c3"}}`)
		default:
			fmt.Fprint(w, `{"data":[],"cursor":{}}`)
		}
	})

	client := rec.client()
	iter := client.V2Session.ListAutoPaging(context.Background(), V2SessionListParams{Limit: Int(2)})

	var ids []string
	for iter.Next() {
		ids = append(ids, iter.Current().ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}

	want := []string{"ses_1", "ses_2", "ses_3", "ses_4", "ses_5"}
	if len(ids) != len(want) {
		t.Fatalf("collected %d sessions (%v), want %d (%v)", len(ids), ids, len(want), want)
	}
	for i := range want {
		if ids[i] != want[i] {
			t.Errorf("item %d = %q, want %q", i, ids[i], want[i])
		}
	}
	if got, want := iter.Index(), 5; got != want {
		t.Errorf("Index() = %d, want %d", got, want)
	}
	if got, want := len(rec.queries), 4; got != want {
		t.Errorf("server saw %d requests, want %d", got, want)
	}
	if got, want := rec.query(t, 3).Get("cursor"), "c3"; got != want {
		t.Errorf("final request cursor = %q, want %q", got, want)
	}
}

func TestV2SessionListAutoPagingStopsOnEmptyFirstPage(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		fmt.Fprint(w, `{"data":[],"cursor":{}}`)
	})

	iter := rec.client().V2Session.ListAutoPaging(context.Background(), V2SessionListParams{})
	if iter.Next() {
		t.Fatal("Next() returned true on an empty first page")
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}
	if got, want := len(rec.queries), 1; got != want {
		t.Errorf("server saw %d requests, want %d", got, want)
	}
}

func TestV2SessionListAutoPagingSurfacesRequestError(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		if hit == 0 {
			fmt.Fprint(w, `{"data":[{"id":"ses_1"}],"cursor":{"next":"c1"}}`)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"_tag":"InvalidCursorError","message":"Cursor cannot be combined with order"}`)
	})

	iter := rec.client().V2Session.ListAutoPaging(context.Background(), V2SessionListParams{})
	if !iter.Next() {
		t.Fatal("Next() returned false on the first item")
	}
	if iter.Next() {
		t.Fatal("Next() returned true after the second page failed")
	}
	if iter.Err() == nil {
		t.Fatal("Err() is nil, the failed page request should surface")
	}
}

// ===== CursorPage: GET /api/session/{sessionID}/message =====

func TestV2SessionMessagesPaginatesThroughBodyCursor(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			fmt.Fprint(w, `{"data":[{"id":"msg_1"}],"cursor":{"previous":"p1","next":"n1"}}`)
		default:
			fmt.Fprint(w, `{"data":[],"cursor":{}}`)
		}
	})

	client := rec.client()
	page, err := client.V2Session.Messages(context.Background(), "ses_abc", V2SessionMessagesParams{
		Limit: Int(1),
		Order: F(V2SessionOrderAsc),
	})
	if err != nil {
		t.Fatalf("Messages: %v", err)
	}
	if got, want := len(page.Data), 1; got != want {
		t.Fatalf("page size = %d, want %d", got, want)
	}
	if got, want := rec.paths[0], "/api/session/ses_abc/message"; got != want {
		t.Errorf("path = %q, want %q", got, want)
	}

	if _, err := page.GetNextPage(); err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	second := rec.query(t, 1)
	if got, want := second.Get("cursor"), "n1"; got != want {
		t.Errorf("request 1 cursor = %q, want %q", got, want)
	}
	// The server answers 400 "Cursor cannot be combined with order" if both are
	// present, so dropping order here is load bearing.
	if second.Has("order") {
		t.Errorf("request 1 must not send order, got %q", second.Get("order"))
	}
}

func TestV2SessionMessagesRequiresSessionID(t *testing.T) {
	client := NewClient(option.WithBaseURL("http://127.0.0.1:1/"))
	page, err := client.V2Session.Messages(context.Background(), "", V2SessionMessagesParams{})
	if err == nil {
		t.Fatal("expected an error for an empty sessionID")
	}
	if page != nil {
		t.Errorf("expected a nil page, got %#v", page)
	}
	if got, want := err.Error(), "missing required sessionID parameter"; got != want {
		t.Errorf("error = %q, want %q", got, want)
	}
}

// ===== SeqPage: GET /api/session/{sessionID}/history =====

func TestV2SessionHistoryPaginatesThroughDurableSeq(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			fmt.Fprint(w, `{"data":[
				{"id":"evt_1","type":"session.next.prompted","durable":{"aggregateID":"ses_1","seq":1,"version":1},"data":{}},
				{"id":"evt_2","type":"session.next.prompted","durable":{"aggregateID":"ses_1","seq":7,"version":1},"data":{}}
			],"hasMore":true}`)
		default:
			fmt.Fprint(w, `{"data":[
				{"id":"evt_3","type":"session.next.prompted","durable":{"aggregateID":"ses_1","seq":9,"version":1},"data":{}}
			],"hasMore":false}`)
		}
	})

	client := rec.client()
	page, err := client.V2Session.History(context.Background(), "ses_1", V2SessionHistoryParams{Limit: Int(2)})
	if err != nil {
		t.Fatalf("History: %v", err)
	}
	if got, want := len(page.Data), 2; got != want {
		t.Fatalf("page size = %d, want %d", got, want)
	}
	if !page.HasMore {
		t.Error("HasMore = false, want true")
	}

	after, ok := page.NextAfter()
	if !ok {
		t.Fatal("NextAfter reported no sequence, want the seq of the last event")
	}
	if got, want := after, int64(7); got != want {
		t.Errorf("NextAfter = %d, want %d (durable.seq of the last event)", got, want)
	}

	next, err := page.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	if next == nil {
		t.Fatal("GetNextPage returned nil, want the second page")
	}
	if got, want := rec.query(t, 1).Get("after"), "7"; got != want {
		t.Errorf("request 1 after = %q, want %q", got, want)
	}
	if got, want := rec.query(t, 1).Get("limit"), "2"; got != want {
		t.Errorf("request 1 limit = %q, want %q", got, want)
	}

	// hasMore is false, so the walk stops without another round trip.
	last, err := next.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage on the final page: %v", err)
	}
	if last != nil {
		t.Errorf("GetNextPage returned %#v after hasMore=false, want nil", last)
	}
	if got, want := len(rec.queries), 2; got != want {
		t.Errorf("server saw %d requests, want %d", got, want)
	}
}

func TestV2SessionHistoryAutoPagingWalksEveryPage(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			fmt.Fprint(w, `{"data":[
				{"id":"evt_1","type":"session.next.prompted","durable":{"aggregateID":"ses_1","seq":1,"version":1},"data":{}}
			],"hasMore":true}`)
		case 1:
			fmt.Fprint(w, `{"data":[
				{"id":"evt_2","type":"session.next.prompted","durable":{"aggregateID":"ses_1","seq":2,"version":1},"data":{}}
			],"hasMore":true}`)
		default:
			fmt.Fprint(w, `{"data":[
				{"id":"evt_3","type":"session.next.prompted","durable":{"aggregateID":"ses_1","seq":3,"version":1},"data":{}}
			],"hasMore":false}`)
		}
	})

	iter := rec.client().V2Session.HistoryAutoPaging(context.Background(), "ses_1", V2SessionHistoryParams{Limit: Int(1)})
	count := 0
	for iter.Next() {
		count++
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}
	if got, want := count, 3; got != want {
		t.Fatalf("collected %d events, want %d", got, want)
	}
	if got, want := len(rec.queries), 3; got != want {
		t.Fatalf("server saw %d requests, want %d", got, want)
	}
	if got, want := rec.query(t, 1).Get("after"), "1"; got != want {
		t.Errorf("request 1 after = %q, want %q", got, want)
	}
	if got, want := rec.query(t, 2).Get("after"), "2"; got != want {
		t.Errorf("request 2 after = %q, want %q", got, want)
	}
}

// ===== HeaderCursorPage: GET /experimental/session =====

func TestExperimentalSessionListPaginatesThroughHeaderCursor(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			w.Header().Set("X-Next-Cursor", "1700000000000")
			fmt.Fprint(w, `[{"id":"ses_1"},{"id":"ses_2"}]`)
		default:
			// No X-Next-Cursor header: this is the last page.
			fmt.Fprint(w, `[{"id":"ses_3"}]`)
		}
	})

	client := rec.client()
	page, err := client.Experimental.Session.List(context.Background(), ExperimentalSessionListParams{Limit: Int(2)})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if got, want := len(page.Data), 2; got != want {
		t.Fatalf("first page size = %d, want %d", got, want)
	}
	if got, want := page.Data[0].ID, "ses_1"; got != want {
		t.Errorf("first item = %q, want %q", got, want)
	}
	if got, want := page.NextCursor(), "1700000000000"; got != want {
		t.Errorf("NextCursor = %q, want %q", got, want)
	}
	if page.RawJSON() == "" {
		t.Error("RawJSON is empty, the raw array body should be retained")
	}

	next, err := page.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	if next == nil {
		t.Fatal("GetNextPage returned nil, want the second page")
	}
	if got, want := next.Data[0].ID, "ses_3"; got != want {
		t.Errorf("second page item = %q, want %q", got, want)
	}
	if got, want := rec.query(t, 1).Get("cursor"), "1700000000000"; got != want {
		t.Errorf("request 1 cursor = %q, want %q", got, want)
	}

	// The header is gone, so the walk is over.
	if got, want := next.NextCursor(), ""; got != want {
		t.Errorf("NextCursor on the last page = %q, want %q", got, want)
	}
	last, err := next.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage on the last page: %v", err)
	}
	if last != nil {
		t.Errorf("GetNextPage returned %#v with no cursor header, want nil", last)
	}
	if got, want := len(rec.queries), 2; got != want {
		t.Errorf("server saw %d requests, want %d", got, want)
	}
}

func TestExperimentalSessionListAutoPagingWalksEveryPage(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			w.Header().Set("X-Next-Cursor", "300")
			fmt.Fprint(w, `[{"id":"ses_1"}]`)
		case 1:
			w.Header().Set("X-Next-Cursor", "200")
			fmt.Fprint(w, `[{"id":"ses_2"}]`)
		default:
			fmt.Fprint(w, `[{"id":"ses_3"}]`)
		}
	})

	iter := rec.client().Experimental.Session.ListAutoPaging(context.Background(), ExperimentalSessionListParams{Limit: Int(1)})
	var ids []string
	for iter.Next() {
		ids = append(ids, iter.Current().ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}

	want := []string{"ses_1", "ses_2", "ses_3"}
	if len(ids) != len(want) {
		t.Fatalf("collected %v, want %v", ids, want)
	}
	for i := range want {
		if ids[i] != want[i] {
			t.Errorf("item %d = %q, want %q", i, ids[i], want[i])
		}
	}
	if got, want := len(rec.queries), 3; got != want {
		t.Errorf("server saw %d requests, want %d", got, want)
	}
	if got, want := rec.query(t, 2).Get("cursor"), "200"; got != want {
		t.Errorf("request 2 cursor = %q, want %q", got, want)
	}
}

// ===== HeaderBeforePage: GET /session/{id}/message =====

func TestSessionMessagesPaginatesThroughBeforeCursor(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			w.Header().Set("X-Next-Cursor", "cursor_page_2")
			fmt.Fprint(w, `[{"info":{"id":"msg_2"},"parts":[]}]`)
		default:
			fmt.Fprint(w, `[{"info":{"id":"msg_1"},"parts":[]}]`)
		}
	})

	client := rec.client()
	page, err := client.Session.Messages(context.Background(), "ses_abc", SessionMessagesParams{Limit: Int(1)})
	if err != nil {
		t.Fatalf("Messages: %v", err)
	}
	if got, want := len(page.Data), 1; got != want {
		t.Fatalf("first page size = %d, want %d", got, want)
	}
	if got, want := page.Data[0].Info.ID, "msg_2"; got != want {
		t.Errorf("first item = %q, want %q", got, want)
	}
	if got, want := rec.paths[0], "/session/ses_abc/message"; got != want {
		t.Errorf("path = %q, want %q", got, want)
	}
	if got, want := page.NextCursor(), "cursor_page_2"; got != want {
		t.Errorf("NextCursor = %q, want %q", got, want)
	}

	next, err := page.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	if next == nil {
		t.Fatal("GetNextPage returned nil, want the second page")
	}
	if got, want := next.Data[0].Info.ID, "msg_1"; got != want {
		t.Errorf("second page item = %q, want %q", got, want)
	}

	second := rec.query(t, 1)
	// This endpoint takes its cursor as `before`, not `cursor`.
	if got, want := second.Get("before"), "cursor_page_2"; got != want {
		t.Errorf("request 1 before = %q, want %q", got, want)
	}
	if second.Has("cursor") {
		t.Errorf("request 1 must not send a `cursor` parameter, got %q", second.Get("cursor"))
	}
	// The server rejects `before` without an explicit `limit`.
	if got, want := second.Get("limit"), "1"; got != want {
		t.Errorf("request 1 limit = %q, want %q", got, want)
	}
}

func TestSessionMessagesAutoPagingWalksEveryPage(t *testing.T) {
	// The server sorts each page oldest first and hands out successively older
	// pages, so a full walk reads as a sawtooth rather than one sorted run.
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			w.Header().Set("X-Next-Cursor", "c2")
			fmt.Fprint(w, `[{"info":{"id":"msg_5"},"parts":[]},{"info":{"id":"msg_6"},"parts":[]}]`)
		case 1:
			w.Header().Set("X-Next-Cursor", "c3")
			fmt.Fprint(w, `[{"info":{"id":"msg_3"},"parts":[]},{"info":{"id":"msg_4"},"parts":[]}]`)
		default:
			fmt.Fprint(w, `[{"info":{"id":"msg_1"},"parts":[]},{"info":{"id":"msg_2"},"parts":[]}]`)
		}
	})

	iter := rec.client().Session.MessagesAutoPaging(context.Background(), "ses_abc", SessionMessagesParams{Limit: Int(2)})
	var ids []string
	for iter.Next() {
		ids = append(ids, iter.Current().Info.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}

	// Each page is yielded in the order the server sent it, and pages are yielded
	// newest batch first.
	want := []string{"msg_5", "msg_6", "msg_3", "msg_4", "msg_1", "msg_2"}
	if len(ids) != len(want) {
		t.Fatalf("collected %v, want %v", ids, want)
	}
	for i := range want {
		if ids[i] != want[i] {
			t.Errorf("item %d = %q, want %q", i, ids[i], want[i])
		}
	}
	if got, want := len(rec.queries), 3; got != want {
		t.Errorf("server saw %d requests, want %d", got, want)
	}
}

func TestSessionMessagesWithoutLimitReturnsSinglePage(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		// With no limit the server returns the whole conversation and never sets
		// X-Next-Cursor.
		fmt.Fprint(w, `[{"info":{"id":"msg_1"},"parts":[]},{"info":{"id":"msg_2"},"parts":[]}]`)
	})

	client := rec.client()
	page, err := client.Session.Messages(context.Background(), "ses_abc", SessionMessagesParams{})
	if err != nil {
		t.Fatalf("Messages: %v", err)
	}
	if got, want := len(page.Data), 2; got != want {
		t.Fatalf("page size = %d, want %d", got, want)
	}
	next, err := page.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	if next != nil {
		t.Errorf("GetNextPage returned %#v with no cursor header, want nil", next)
	}
	if got, want := len(rec.queries), 1; got != want {
		t.Errorf("server saw %d requests, want %d", got, want)
	}
}

func TestSessionMessagesRequiresID(t *testing.T) {
	client := NewClient(option.WithBaseURL("http://127.0.0.1:1/"))
	page, err := client.Session.Messages(context.Background(), "", SessionMessagesParams{})
	if err == nil {
		t.Fatal("expected an error for an empty id")
	}
	if page != nil {
		t.Errorf("expected a nil page, got %#v", page)
	}
	if got, want := err.Error(), "missing required id parameter"; got != want {
		t.Errorf("error = %q, want %q", got, want)
	}
}

// ===== Page types in isolation =====

func TestCursorPageUnmarshalKeepsAliasedShape(t *testing.T) {
	// V2SessionsResponse is an alias for pagination.CursorPage[V2SessionInfo], so
	// the historical field layout must still decode.
	var page V2SessionsResponse
	body := `{"data":[{"id":"ses_1","title":"hello"}],"cursor":{"previous":"p","next":"n"}}`
	if err := json.Unmarshal([]byte(body), &page); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got, want := len(page.Data), 1; got != want {
		t.Fatalf("len(Data) = %d, want %d", got, want)
	}
	if got, want := page.Data[0].Title, "hello"; got != want {
		t.Errorf("Data[0].Title = %q, want %q", got, want)
	}

	var cursor V2Cursor = page.Cursor
	if got, want := cursor.Next, "n"; got != want {
		t.Errorf("Cursor.Next = %q, want %q", got, want)
	}
	if got, want := cursor.Previous, "p"; got != want {
		t.Errorf("Cursor.Previous = %q, want %q", got, want)
	}
	if got, want := page.RawJSON(), body; got != want {
		t.Errorf("RawJSON = %q, want %q", got, want)
	}
}

func TestSeqPageNextAfterWithoutDurableSeq(t *testing.T) {
	var page V2SessionHistoryResponse
	if err := json.Unmarshal([]byte(`{"data":[{"id":"evt_1","type":"x","data":{}}],"hasMore":true}`), &page); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, ok := page.NextAfter(); ok {
		t.Error("NextAfter reported a sequence for an event without durable.seq")
	}
	next, err := page.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	if next != nil {
		t.Errorf("GetNextPage returned %#v without an advanceable sequence, want nil", next)
	}
}

func TestSeqPageNextAfterOnEmptyPage(t *testing.T) {
	var page V2SessionHistoryResponse
	if err := json.Unmarshal([]byte(`{"data":[],"hasMore":false}`), &page); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if _, ok := page.NextAfter(); ok {
		t.Error("NextAfter reported a sequence for an empty page")
	}
}

func TestUnboundPagesCannotAdvance(t *testing.T) {
	t.Run("CursorPage", func(t *testing.T) {
		var page V2SessionsResponse
		if err := json.Unmarshal([]byte(`{"data":[{"id":"ses_1"}],"cursor":{"previous":"p","next":"n"}}`), &page); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if _, err := page.GetNextPage(); !errors.Is(err, pagination.ErrPageUnbound) {
			t.Errorf("GetNextPage error = %v, want %v", err, pagination.ErrPageUnbound)
		}
		if _, err := page.GetPreviousPage(); !errors.Is(err, pagination.ErrPageUnbound) {
			t.Errorf("GetPreviousPage error = %v, want %v", err, pagination.ErrPageUnbound)
		}
	})

	t.Run("SeqPage", func(t *testing.T) {
		var page V2SessionHistoryResponse
		body := `{"data":[{"id":"evt_1","durable":{"seq":3}}],"hasMore":true}`
		if err := json.Unmarshal([]byte(body), &page); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if _, err := page.GetNextPage(); !errors.Is(err, pagination.ErrPageUnbound) {
			t.Errorf("GetNextPage error = %v, want %v", err, pagination.ErrPageUnbound)
		}
	})
}

func TestHeaderPagesWithoutResponseHaveNoCursor(t *testing.T) {
	var cursorPage pagination.HeaderCursorPage[GlobalSession]
	if err := json.Unmarshal([]byte(`[{"id":"ses_1"}]`), &cursorPage); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got, want := len(cursorPage.Data), 1; got != want {
		t.Fatalf("len(Data) = %d, want %d", got, want)
	}
	if got := cursorPage.NextCursor(); got != "" {
		t.Errorf("NextCursor = %q, want an empty string", got)
	}
	next, err := cursorPage.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	if next != nil {
		t.Errorf("GetNextPage returned %#v, want nil", next)
	}

	var beforePage pagination.HeaderBeforePage[SessionMessagesResponse]
	if err := json.Unmarshal([]byte(`[{"info":{"id":"msg_1"},"parts":[]}]`), &beforePage); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if got, want := beforePage.Data[0].Info.ID, "msg_1"; got != want {
		t.Errorf("Data[0].Info.ID = %q, want %q", got, want)
	}
	if got := beforePage.NextCursor(); got != "" {
		t.Errorf("NextCursor = %q, want an empty string", got)
	}
}

func TestNilPagesAdvanceToNothing(t *testing.T) {
	var cursorPage *pagination.CursorPage[V2SessionInfo]
	if next, err := cursorPage.GetNextPage(); next != nil || err != nil {
		t.Errorf("GetNextPage on a nil CursorPage = (%v, %v), want (nil, nil)", next, err)
	}
	var seqPage *pagination.SeqPage[V2SessionDurableEvent]
	if next, err := seqPage.GetNextPage(); next != nil || err != nil {
		t.Errorf("GetNextPage on a nil SeqPage = (%v, %v), want (nil, nil)", next, err)
	}
	var headerPage *pagination.HeaderCursorPage[GlobalSession]
	if next, err := headerPage.GetNextPage(); next != nil || err != nil {
		t.Errorf("GetNextPage on a nil HeaderCursorPage = (%v, %v), want (nil, nil)", next, err)
	}
	var beforePage *pagination.HeaderBeforePage[SessionMessagesResponse]
	if next, err := beforePage.GetNextPage(); next != nil || err != nil {
		t.Errorf("GetNextPage on a nil HeaderBeforePage = (%v, %v), want (nil, nil)", next, err)
	}
}

// ===== Regressions =====

// countingDoer is an option.HTTPClient that is not an *http.Client, which is the
// case that routes through RequestConfig.CustomHTTPDoer.
type countingDoer struct {
	calls int
	inner *http.Client
}

func (d *countingDoer) Do(req *http.Request) (*http.Response, error) {
	d.calls++
	return d.inner.Do(req)
}

// A custom HTTPDoer used to be dropped when a page cloned its request config, so
// every request after the first silently bypassed the caller's transport.
func TestPagingKeepsCustomHTTPDoer(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			fmt.Fprint(w, `{"data":[{"id":"ses_1"}],"cursor":{"next":"c1"}}`)
		case 1:
			fmt.Fprint(w, `{"data":[{"id":"ses_2"}],"cursor":{"next":"c2"}}`)
		default:
			fmt.Fprint(w, `{"data":[],"cursor":{}}`)
		}
	})

	doer := &countingDoer{inner: rec.Client()}
	client := NewClient(option.WithBaseURL(rec.URL+"/"), option.WithHTTPClient(doer))

	iter := client.V2Session.ListAutoPaging(context.Background(), V2SessionListParams{Limit: Int(1)})
	count := 0
	for iter.Next() {
		count++
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}
	if got, want := count, 2; got != want {
		t.Fatalf("collected %d sessions, want %d", got, want)
	}
	if got, want := doer.calls, len(rec.queries); got != want {
		t.Errorf("custom doer saw %d requests but the server saw %d: paging bypassed the caller's HTTP client", got, want)
	}
	if got, want := doer.calls, 3; got != want {
		t.Errorf("custom doer saw %d requests, want %d", got, want)
	}
}

// A caller that supplies its own WithResponseInto takes over the destination the
// raw response lands in; header driven pages must still find it, otherwise they
// silently stop after the first page.
func TestHeaderPagingWithCallerSuppliedResponseInto(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			w.Header().Set("X-Next-Cursor", "300")
			fmt.Fprint(w, `[{"id":"ses_1"}]`)
		default:
			fmt.Fprint(w, `[{"id":"ses_2"}]`)
		}
	})

	var mine *http.Response
	client := rec.client()
	page, err := client.Experimental.Session.List(
		context.Background(),
		ExperimentalSessionListParams{Limit: Int(1)},
		option.WithResponseInto(&mine),
	)
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if mine == nil {
		t.Fatal("the caller's WithResponseInto was not honoured")
	}
	if got, want := page.NextCursor(), "300"; got != want {
		t.Fatalf("NextCursor = %q, want %q", got, want)
	}

	next, err := page.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	if next == nil {
		t.Fatal("GetNextPage returned nil, paging stopped even though a cursor header was present")
	}
	if got, want := next.Data[0].ID, "ses_2"; got != want {
		t.Errorf("second page item = %q, want %q", got, want)
	}
}

// Advancing a page must not mutate it, so the same page can be replayed.
func TestGetNextPageIsRepeatable(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		if hit == 0 {
			fmt.Fprint(w, `{"data":[{"id":"ses_1"}],"cursor":{"next":"c1"}}`)
			return
		}
		fmt.Fprint(w, `{"data":[{"id":"ses_2"}],"cursor":{"next":"c2"}}`)
	})

	page, err := rec.client().V2Session.List(context.Background(), V2SessionListParams{Limit: Int(1)})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if _, err := page.GetNextPage(); err != nil {
		t.Fatalf("first GetNextPage: %v", err)
	}
	if _, err := page.GetNextPage(); err != nil {
		t.Fatalf("second GetNextPage: %v", err)
	}

	first, second := rec.query(t, 1), rec.query(t, 2)
	if first.Encode() != second.Encode() {
		t.Errorf("replaying GetNextPage sent %q then %q, want identical queries", first.Encode(), second.Encode())
	}
	if got, want := first.Get("cursor"), "c1"; got != want {
		t.Errorf("cursor = %q, want %q", got, want)
	}
}

func TestHeaderPageDecodesNullBodyAsEmpty(t *testing.T) {
	var page pagination.HeaderCursorPage[GlobalSession]
	if err := json.Unmarshal([]byte(`null`), &page); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(page.Data) != 0 {
		t.Errorf("len(Data) = %d, want 0", len(page.Data))
	}
}

// A literal JSON `null` body decodes cfg.ResponseBodyInto's pointee (the page
// itself) to nil, so the service method calls (*CursorPage[T])(nil).SetPageConfig
// on a nil receiver. SetPageConfig used to reassign its *local* copy of that nil
// pointer to &T{} on entry, which is a no-op the caller never observes (Go passes
// pointer receivers by value), so the returned page stayed nil either way; the
// guard clause form must preserve exactly that behavior without panicking.
func TestListWithNullBodyReturnsNilPageWithoutPanic(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		fmt.Fprint(w, `null`)
	})

	page, err := rec.client().V2Session.List(context.Background(), V2SessionListParams{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if page != nil {
		t.Errorf("List with a null body = %#v, want nil", page)
	}
}

// ===== Concurrency: Clone() is now shared across goroutines that advance the
// same page, since GetNextPage clones r.cfg (which carries DefaultBaseURL and
// CustomHTTPDoer as of this round's fix) instead of mutating it. Run with
// `-race` to catch data races on the shared RequestConfig or a shared custom
// HTTPDoer.

// atomicCountingDoer is like countingDoer but safe for concurrent use, so a
// race reported against it can only come from RequestConfig.Clone or
// RequestConfig.Execute, not from the test double itself.
type atomicCountingDoer struct {
	calls int64
	inner *http.Client
}

func (d *atomicCountingDoer) Do(req *http.Request) (*http.Response, error) {
	atomic.AddInt64(&d.calls, 1)
	return d.inner.Do(req)
}

// The handler is intentionally stateless (keyed only by the incoming query
// string, never by call count) so that concurrent hits from the goroutines
// below can't race on server-side bookkeeping and any race the detector finds
// is attributable to the client-side pagination code under test.
func TestConcurrentGetNextPageOnSharedPageIsRaceFree(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("cursor") == "" {
			fmt.Fprint(w, `{"data":[{"id":"ses_1"}],"cursor":{"next":"c1"}}`)
			return
		}
		// Every subsequent request (regardless of how many goroutines send one
		// concurrently) advertises the same next cursor, so the walk never
		// terminates on its own; the test bounds the number of hops explicitly.
		fmt.Fprint(w, `{"data":[{"id":"ses_2"}],"cursor":{"next":"c1"}}`)
	}))
	t.Cleanup(srv.Close)

	doer := &atomicCountingDoer{inner: srv.Client()}
	client := NewClient(option.WithBaseURL(srv.URL+"/"), option.WithHTTPClient(doer))

	page, err := client.V2Session.List(context.Background(), V2SessionListParams{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}

	const goroutines = 32
	var wg sync.WaitGroup
	pages := make([]*V2SessionsResponse, goroutines)
	errs := make([]error, goroutines)
	for i := range goroutines {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Every goroutine advances the *same* page value concurrently. This
			// only reads page.cfg/page.Cursor and clones them, so it must not
			// race even though there is no synchronization between goroutines.
			pages[i], errs[i] = page.GetNextPage()
		}(i)
	}
	wg.Wait()

	for i := range goroutines {
		if errs[i] != nil {
			t.Fatalf("goroutine %d: GetNextPage: %v", i, errs[i])
		}
		if pages[i] == nil {
			t.Fatalf("goroutine %d: GetNextPage returned nil, want a page", i)
		}
		if got, want := pages[i].Data[0].ID, "ses_2"; got != want {
			t.Errorf("goroutine %d: item = %q, want %q", i, got, want)
		}
	}

	// One request for the initial List plus one per goroutine; if the shared
	// CustomHTTPDoer had been dropped by Clone (as it was before this round's
	// fix), the follow-up requests would fall through to http.DefaultClient
	// instead and this count would stay at 1.
	if got, want := atomic.LoadInt64(&doer.calls), int64(goroutines+1); got != want {
		t.Errorf("custom doer saw %d requests, want %d", got, want)
	}
}

// A second, independent page bound to a different client/doer must not have
// its RequestConfig fields disturbed by concurrent Clone calls happening on
// the first page's RequestConfig, since Clone must not mutate its receiver.
func TestConcurrentGetNextPageAcrossDistinctPagesIsRaceFree(t *testing.T) {
	newServer := func() *httptest.Server {
		return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if r.URL.Query().Get("cursor") == "" {
				fmt.Fprint(w, `{"data":[{"id":"a1"}],"cursor":{"next":"c1"}}`)
				return
			}
			fmt.Fprint(w, `{"data":[{"id":"a2"}],"cursor":{"next":"c1"}}`)
		}))
	}
	srvA, srvB := newServer(), newServer()
	t.Cleanup(srvA.Close)
	t.Cleanup(srvB.Close)

	clientA := NewClient(option.WithBaseURL(srvA.URL + "/"))
	clientB := NewClient(option.WithBaseURL(srvB.URL + "/"))

	pageA, err := clientA.V2Session.List(context.Background(), V2SessionListParams{})
	if err != nil {
		t.Fatalf("List A: %v", err)
	}
	pageB, err := clientB.V2Session.List(context.Background(), V2SessionListParams{})
	if err != nil {
		t.Fatalf("List B: %v", err)
	}

	var wg sync.WaitGroup
	const goroutines = 16
	errsA := make([]error, goroutines)
	errsB := make([]error, goroutines)
	for i := range goroutines {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, errsA[i] = pageA.GetNextPage()
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, errsB[i] = pageB.GetNextPage()
		}(i)
	}
	wg.Wait()

	for i := range goroutines {
		if errsA[i] != nil {
			t.Errorf("pageA goroutine %d: %v", i, errsA[i])
		}
		if errsB[i] != nil {
			t.Errorf("pageB goroutine %d: %v", i, errsB[i])
		}
	}
}

// ===== Round 2 review additions =====

// GetNextPage always replays the *original* context captured by the very first
// call (CursorPage.cfg.Context), not a fresh one, so cancelling the context the
// caller used to fetch page one must abort every later hop too. There was no
// test asserting this propagation before this round.
func TestGetNextPagePropagatesContextCancellation(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		fmt.Fprint(w, `{"data":[{"id":"ses_1"}],"cursor":{"next":"c1"}}`)
	})

	ctx, cancel := context.WithCancel(context.Background())
	page, err := rec.client().V2Session.List(ctx, V2SessionListParams{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}

	cancel()
	next, err := page.GetNextPage()
	if err == nil {
		t.Fatal("GetNextPage succeeded after its context was canceled, want an error")
	}
	if next != nil {
		t.Errorf("GetNextPage returned %#v after cancellation, want nil", next)
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("GetNextPage error = %v, want context.Canceled", err)
	}
}

// A 4xx/5xx response from a follow-up page must surface as the same
// *apierror.Error the first request would have produced, carrying the real
// status code and body, not a generic error the caller can't type-assert on.
func TestGetNextPageSurfacesTypedAPIError(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		if hit == 0 {
			fmt.Fprint(w, `{"data":[{"id":"ses_1"}],"cursor":{"next":"c1"}}`)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"_tag":"InvalidCursorError","message":"Cursor cannot be combined with order"}`)
	})

	page, err := rec.client().V2Session.List(context.Background(), V2SessionListParams{})
	if err != nil {
		t.Fatalf("List: %v", err)
	}

	next, err := page.GetNextPage()
	if err == nil {
		t.Fatal("GetNextPage succeeded against a 400 response, want an error")
	}
	if next != nil {
		t.Errorf("GetNextPage returned %#v alongside an error, want nil", next)
	}
	var apiErr *apierror.Error
	if !errors.As(err, &apiErr) {
		t.Fatalf("GetNextPage error = %T (%v), want *apierror.Error", err, err)
	}
	if got, want := apiErr.StatusCode, http.StatusBadRequest; got != want {
		t.Errorf("apiErr.StatusCode = %d, want %d", got, want)
	}
	if got := apiErr.JSON.RawJSON(); got == "" {
		t.Error("apiErr.JSON.RawJSON() is empty, the error body should be retained")
	} else if want := "InvalidCursorError"; !strings.Contains(got, want) {
		t.Errorf("apiErr.JSON.RawJSON() = %q, want it to contain %q", got, want)
	}
}

// SeqPage.GetNextPage only stops when the server reports hasMore=false; it does
// not detect a durable.seq that fails to advance. A server bug that keeps
// reporting hasMore=true for the same seq would make an unguarded caller loop
// forever. The only thing that can break the cycle today is a context
// deadline supplied by the caller, which this test exercises directly: it
// bounds its own wait so a regression that removes context checking from the
// request path would hang this test instead of silently passing.
func TestSeqPageAutoPagingStopsQuicklyWhenSeqStalls(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		// Always the same seq and hasMore=true: a stalled durable log.
		fmt.Fprint(w, `{"data":[{"id":"evt_1","type":"x","durable":{"aggregateID":"ses_1","seq":1,"version":1},"data":{}}],"hasMore":true}`)
	})

	// No deadline: the walk has to terminate on its own because the sequence
	// stopped advancing, not because the caller cut it short.
	iter := rec.client().V2Session.HistoryAutoPaging(context.Background(), "ses_1", V2SessionHistoryParams{Limit: Int(1)})
	done := make(chan struct{})
	go func() {
		for iter.Next() {
		}
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("auto pager did not stop within 5s: a stalled durable.seq loops forever")
	}
	if err := iter.Err(); err != nil {
		t.Errorf("Err() = %v, want nil: a stalled sequence is a clean end of walk, not a failure", err)
	}
	if got, want := len(rec.queries), 2; got != want {
		t.Errorf("server saw %d requests, want %d (one initial, one that proved the sequence stalled)", got, want)
	}
}

// GetNextPage clones the previous page's RequestConfig, so `search`/`workspace`
// (and any other query the caller originally sent) must still be present on
// every later hop, not just `limit`. Only `limit` is exercised by the existing
// tests; this pins down the rest of the caller's query.
func TestV2SessionListPreservesAdditionalQueryParamsAcrossPages(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			fmt.Fprint(w, `{"data":[{"id":"ses_1"}],"cursor":{"next":"c1"}}`)
		default:
			fmt.Fprint(w, `{"data":[{"id":"ses_2"}],"cursor":{}}`)
		}
	})

	page, err := rec.client().V2Session.List(context.Background(), V2SessionListParams{
		Limit:     Int(5),
		Search:    F("hello"),
		Workspace: F("ws_1"),
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if _, err := page.GetNextPage(); err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}

	second := rec.query(t, 1)
	if got, want := second.Get("search"), "hello"; got != want {
		t.Errorf("request 1 search = %q, want %q", got, want)
	}
	if got, want := second.Get("workspace"), "ws_1"; got != want {
		t.Errorf("request 1 workspace = %q, want %q", got, want)
	}
	if got, want := second.Get("limit"), "5"; got != want {
		t.Errorf("request 1 limit = %q, want %q", got, want)
	}
}

// Auto-paging through many pages must terminate on the server's own signal
// (an eventual empty page) rather than relying on some hidden internal cap;
// this guards against an off-by-one or accidental infinite loop at scale.
func TestV2SessionListAutoPagingHandlesManyPagesWithoutHanging(t *testing.T) {
	const totalPages = 50
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		if hit >= totalPages {
			fmt.Fprint(w, `{"data":[],"cursor":{}}`)
			return
		}
		fmt.Fprintf(w, `{"data":[{"id":"ses_%d"}],"cursor":{"next":"c%d"}}`, hit, hit)
	})

	iter := rec.client().V2Session.ListAutoPaging(context.Background(), V2SessionListParams{Limit: Int(1)})
	done := make(chan struct{})
	var count int
	go func() {
		for iter.Next() {
			count++
		}
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("auto pager did not stop within 5s while walking 50 pages")
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}
	if got, want := count, totalPages; got != want {
		t.Fatalf("collected %d items, want %d", got, want)
	}
	if got, want := len(rec.queries), totalPages+1; got != want {
		t.Errorf("server saw %d requests, want %d (one empty terminator page)", got, want)
	}
}

// A session history only makes progress while durable sequence numbers increase.
// If the server keeps claiming hasMore while replaying the same sequence, paging
// must stop instead of spinning on the same window forever.
func TestSeqPageStopsWhenSequenceDoesNotAdvance(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		fmt.Fprint(w, `{"data":[
			{"id":"evt_1","type":"session.next.prompted","durable":{"aggregateID":"ses_1","seq":5,"version":1},"data":{}}
		],"hasMore":true}`)
	})

	client := rec.client()
	page, err := client.V2Session.History(context.Background(), "ses_1", V2SessionHistoryParams{Limit: Int(1)})
	if err != nil {
		t.Fatalf("History: %v", err)
	}

	// The first hop is legitimate: the initial request sent no `after` at all.
	second, err := page.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage: %v", err)
	}
	if second == nil {
		t.Fatal("GetNextPage returned nil on the first hop, want a page")
	}
	if got, want := rec.query(t, 1).Get("after"), "5"; got != want {
		t.Fatalf("request 1 after = %q, want %q", got, want)
	}

	// The server replayed seq 5, so there is no progress to be made.
	third, err := second.GetNextPage()
	if err != nil {
		t.Fatalf("GetNextPage on the stalled page: %v", err)
	}
	if third != nil {
		t.Errorf("GetNextPage returned %#v for a stalled sequence, want nil", third)
	}
	if got, want := len(rec.queries), 2; got != want {
		t.Errorf("server saw %d requests, want %d", got, want)
	}
}

// The same stall must terminate the auto-pager rather than hang it.
func TestSeqPageAutoPagingTerminatesOnStalledSequence(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		fmt.Fprint(w, `{"data":[
			{"id":"evt_1","type":"session.next.prompted","durable":{"aggregateID":"ses_1","seq":5,"version":1},"data":{}}
		],"hasMore":true}`)
	})

	done := make(chan int, 1)
	go func() {
		iter := rec.client().V2Session.HistoryAutoPaging(context.Background(), "ses_1", V2SessionHistoryParams{Limit: Int(1)})
		count := 0
		for iter.Next() {
			count++
			if count > 100 {
				break
			}
		}
		done <- count
	}()

	select {
	case count := <-done:
		if got, want := count, 2; got != want {
			t.Errorf("auto pager yielded %d events, want %d (one per non-stalled request)", got, want)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("auto pager did not terminate on a stalled sequence")
	}
}

// ===== Parity with the opencode web client =====
//
// The tests below pin the SDK against the three distinct strategies the web app
// actually uses, so that a change here shows up as a failing test rather than as
// a page of sessions that silently goes missing in a consumer.

// The web app loads the sidebar with
//
//	GET /session?directory=<dir>&roots=true&limit=<n>
//
// (app/src/context/global-sync/session-load.ts loadRootSessionsV1). This pins the
// exact wire format the SDK produces for that call.
func TestSessionListMatchesWebClientRequest(t *testing.T) {
	const directory = "/Users/maiyuan/Code/go-project/maiyuan/customeow-goframe"

	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		fmt.Fprint(w, `[{"id":"ses_1","title":"first"},{"id":"ses_2","title":"second"}]`)
	})

	sessions, err := rec.client().Session.List(context.Background(), SessionListParams{
		Directory: F(directory),
		Roots:     F(true),
		Limit:     Int(55),
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if got, want := len(*sessions), 2; got != want {
		t.Fatalf("len(sessions) = %d, want %d", got, want)
	}
	if got, want := (*sessions)[0].ID, "ses_1"; got != want {
		t.Errorf("sessions[0].ID = %q, want %q", got, want)
	}

	if got, want := rec.paths[0], "/session"; got != want {
		t.Errorf("path = %q, want %q", got, want)
	}
	q := rec.query(t, 0)
	if got, want := q.Get("directory"), directory; got != want {
		t.Errorf("directory = %q, want %q", got, want)
	}
	if got, want := q.Get("roots"), "true"; got != want {
		t.Errorf("roots = %q, want %q", got, want)
	}
	if got, want := q.Get("limit"), "55"; got != want {
		t.Errorf("limit = %q, want %q", got, want)
	}
	// Nothing else may leak onto the wire: unset params must stay absent rather
	// than being sent as empty strings.
	if got, want := len(q), 3; got != want {
		t.Errorf("query carried %d params (%v), want exactly %d", got, q.Encode(), want)
	}
	if got, want := q.Encode(), "directory="+url.QueryEscape(directory)+"&limit=55&roots=true"; got != want {
		t.Errorf("query = %q, want %q", got, want)
	}
}

// GET /session is the one list endpoint with no pagination at all: the handler
// returns a bare array and, unlike GET /session/{id}/message, never sets
// X-Next-Cursor. This test pins that so nobody later assumes a cursor exists.
func TestSessionListCarriesNoCursorSignal(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		if got := r.URL.Query().Get("cursor"); got != "" {
			t.Errorf("SDK sent cursor=%q to an endpoint that has none", got)
		}
		if got := r.URL.Query().Get("before"); got != "" {
			t.Errorf("SDK sent before=%q to an endpoint that has none", got)
		}
		fmt.Fprint(w, `[{"id":"ses_1"},{"id":"ses_2"}]`)
	})

	var raw *http.Response
	if _, err := rec.client().Session.List(
		context.Background(),
		SessionListParams{Limit: Int(2)},
		option.WithResponseInto(&raw),
	); err != nil {
		t.Fatalf("List: %v", err)
	}
	if got := raw.Header.Get("X-Next-Cursor"); got != "" {
		t.Errorf("X-Next-Cursor = %q, want empty: /session does not paginate", got)
	}
	if got := raw.Header.Get("Link"); got != "" {
		t.Errorf("Link = %q, want empty: /session does not paginate", got)
	}
	if got, want := len(rec.queries), 1; got != want {
		t.Errorf("server saw %d requests, want %d", got, want)
	}
}

// Because /session has no cursor, the web app grows `limit` instead of paging:
// limit = max(retained + SESSION_RECENT_LIMIT, SESSION_RECENT_LIMIT), and it
// infers "there may be more" from len(data) == limit
// (app/src/context/global-sync/session-load.ts estimateRootSessionTotal).
// This verifies the SDK supports that loop.
func TestSessionListSupportsWebClientLimitGrowth(t *testing.T) {
	const recentLimit = 50 // SESSION_RECENT_LIMIT

	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			t.Errorf("request %d sent a non numeric limit: %v", hit, err)
			limit = 0
		}
		// The store holds 60 root sessions, so a request is saturated until the
		// limit exceeds that.
		const total = 60
		n := min(limit, total)
		items := make([]string, 0, n)
		for i := range n {
			items = append(items, fmt.Sprintf(`{"id":"ses_%d"}`, i))
		}
		fmt.Fprintf(w, "[%s]", strings.Join(items, ","))
	})

	client := rec.client()
	load := func(retained int) (int, int, bool) {
		limit := max(retained+recentLimit, recentLimit)
		sessions, err := client.Session.List(context.Background(), SessionListParams{
			Directory: F("/repo"),
			Roots:     F(true),
			Limit:     Int(int64(limit)),
		})
		if err != nil {
			t.Fatalf("List(limit=%d): %v", limit, err)
		}
		count := len(*sessions)
		// estimateRootSessionTotal: a saturated page means there is probably more.
		return limit, count, count >= limit
	}

	limit, count, more := load(5)
	if got, want := limit, 55; got != want {
		t.Fatalf("first limit = %d, want %d (retained 5 + recent 50, the value the web app sends)", got, want)
	}
	if got, want := count, 55; got != want {
		t.Fatalf("first page returned %d sessions, want %d", got, want)
	}
	if !more {
		t.Fatal("a saturated page should be reported as possibly having more")
	}

	// Growing the window is what actually reaches the rest of the sessions.
	limit, count, more = load(50)
	if got, want := limit, 100; got != want {
		t.Fatalf("grown limit = %d, want %d", got, want)
	}
	if got, want := count, 60; got != want {
		t.Fatalf("grown page returned %d sessions, want %d (the whole store)", got, want)
	}
	if more {
		t.Error("an unsaturated page should be reported as complete")
	}

	if got, want := len(rec.queries), 2; got != want {
		t.Fatalf("server saw %d requests, want %d", got, want)
	}
	if got, want := rec.query(t, 0).Get("limit"), "55"; got != want {
		t.Errorf("request 0 limit = %q, want %q", got, want)
	}
	if got, want := rec.query(t, 1).Get("limit"), "100"; got != want {
		t.Errorf("request 1 limit = %q, want %q", got, want)
	}
}

// The web app pages v1 messages with `limit` + `before`, reading the next cursor
// from the x-next-cursor header and stopping when the header is absent
// (app/src/context/server-session.ts). This walks that exact loop through the SDK.
func TestSessionMessagesMatchesWebClientPagination(t *testing.T) {
	const pageSize = 50

	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		// The real handler writes the header in lower case; Go canonicalises
		// header keys on read, and this pins that assumption end to end.
		switch hit {
		case 0:
			w.Header()["x-next-cursor"] = []string{"cursor_older_1"}
			fmt.Fprint(w, buildV1MessagePage(1, pageSize))
		case 1:
			w.Header()["x-next-cursor"] = []string{"cursor_older_2"}
			fmt.Fprint(w, buildV1MessagePage(101, pageSize))
		default:
			// A short page: the handler omits the header entirely.
			fmt.Fprint(w, buildV1MessagePage(201, 12))
		}
	})

	client := rec.client()
	iter := client.Session.MessagesAutoPaging(context.Background(), "ses_web", SessionMessagesParams{
		Directory: F("/repo"),
		Limit:     Int(pageSize),
	})

	seen := 0
	for iter.Next() {
		if iter.Current().Info.ID == "" {
			t.Fatalf("message %d has no id", seen)
		}
		seen++
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}
	if got, want := seen, pageSize+pageSize+12; got != want {
		t.Fatalf("collected %d messages, want %d", got, want)
	}
	if got, want := len(rec.queries), 3; got != want {
		t.Fatalf("server saw %d requests, want %d", got, want)
	}

	// The first request must not carry a cursor, and every later one must carry
	// the previous response's header value plus the original limit, because the
	// handler rejects `before` without an explicit `limit`.
	first := rec.query(t, 0)
	if first.Has("before") {
		t.Errorf("request 0 sent before=%q, want none", first.Get("before"))
	}
	if got, want := first.Get("directory"), "/repo"; got != want {
		t.Errorf("request 0 directory = %q, want %q", got, want)
	}
	for i, wantCursor := range []string{"cursor_older_1", "cursor_older_2"} {
		q := rec.query(t, i+1)
		if got := q.Get("before"); got != wantCursor {
			t.Errorf("request %d before = %q, want %q", i+1, got, wantCursor)
		}
		if got, want := q.Get("limit"), strconv.Itoa(pageSize); got != want {
			t.Errorf("request %d limit = %q, want %q", i+1, got, want)
		}
		if got, want := q.Get("directory"), "/repo"; got != want {
			t.Errorf("request %d directory = %q, want %q", i+1, got, want)
		}
	}
}

// The v2 timeline pages on the body cursor and, per the web app, is finished
// when a page comes back with no data at all.
func TestV2SessionMessagesMatchesWebClientPagination(t *testing.T) {
	userMessage := func(id string) string {
		return fmt.Sprintf(`{"id":%q,"type":"user","text":"hello %s","time":{"created":1}}`, id, id)
	}

	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		switch hit {
		case 0:
			fmt.Fprintf(w, `{"data":[%s,%s],"cursor":{"previous":"p1","next":"n1"}}`, userMessage("msg_1"), userMessage("msg_2"))
		case 1:
			fmt.Fprintf(w, `{"data":[%s],"cursor":{"previous":"p2","next":"n2"}}`, userMessage("msg_3"))
		default:
			// The server hands out a cursor even for the final non empty page, so
			// the walk ends on an empty page.
			fmt.Fprint(w, `{"data":[],"cursor":{}}`)
		}
	})

	iter := rec.client().V2Session.MessagesAutoPaging(context.Background(), "ses_web", V2SessionMessagesParams{
		Limit: Int(2),
		Order: F(V2SessionOrderAsc),
	})

	var ids []string
	for iter.Next() {
		user, ok := iter.Current().AsUnion().(V2SessionMessageUser)
		if !ok {
			t.Fatalf("message %d resolved to %T, want V2SessionMessageUser", len(ids), iter.Current().AsUnion())
		}
		if got, want := user.Type, V2SessionMessageUserTypeUser; got != want {
			t.Errorf("message %d type = %q, want %q", len(ids), got, want)
		}
		ids = append(ids, user.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}
	if want := []string{"msg_1", "msg_2", "msg_3"}; !reflect.DeepEqual(ids, want) {
		t.Fatalf("collected %v, want %v", ids, want)
	}
	if got, want := len(rec.queries), 3; got != want {
		t.Fatalf("server saw %d requests, want %d", got, want)
	}

	// order belongs to the first request only; afterwards the cursor carries it
	// and sending both is a 400.
	if got, want := rec.query(t, 0).Get("order"), "asc"; got != want {
		t.Errorf("request 0 order = %q, want %q", got, want)
	}
	for i, wantCursor := range []string{"n1", "n2"} {
		q := rec.query(t, i+1)
		if got := q.Get("cursor"); got != wantCursor {
			t.Errorf("request %d cursor = %q, want %q", i+1, got, wantCursor)
		}
		if q.Has("order") {
			t.Errorf("request %d sent order=%q alongside a cursor", i+1, q.Get("order"))
		}
	}
}

func buildV1MessagePage(start, count int) string {
	items := make([]string, 0, count)
	for i := range count {
		items = append(items, fmt.Sprintf(`{"info":{"id":"msg_%d","sessionID":"ses_web"},"parts":[]}`, start+i))
	}
	return "[" + strings.Join(items, ",") + "]"
}

// The option prologue on a paginated method folds three sources into one slice:
// the SDK's own WithResponseInto first, then the client's options, then the
// caller's. Order is load bearing, because options are applied in sequence and
// the last writer wins. This pins the precedence directly rather than leaving it
// implied by the other tests.
func TestPaginatedMethodOptionPrecedence(t *testing.T) {
	rec := newRecordingServer(t, func(w http.ResponseWriter, r *http.Request, hit int) {
		fmt.Fprint(w, `{"data":[{"id":"ses_1"}],"cursor":{}}`)
	})

	client := NewClient(
		option.WithBaseURL(rec.URL+"/"),
		option.WithHeader("X-Origin", "client"),
		option.WithHeader("X-Client-Only", "yes"),
	)

	var raw *http.Response
	if _, err := client.V2Session.List(
		context.Background(),
		V2SessionListParams{},
		option.WithHeader("X-Origin", "call"),
	); err != nil {
		t.Fatalf("List: %v", err)
	}

	got := rec.headers[0]
	// Caller options are applied after client options, so they win.
	if v := got.Get("X-Origin"); v != "call" {
		t.Errorf("X-Origin = %q, want %q: a per-call option must override the client option", v, "call")
	}
	// Client options the caller did not touch must survive.
	if v := got.Get("X-Client-Only"); v != "yes" {
		t.Errorf("X-Client-Only = %q, want %q: client options must still apply", v, "yes")
	}
	// And the SDK's own WithResponseInto is applied first, so it is still in
	// effect when the caller supplies none.
	if _, err := client.V2Session.List(context.Background(), V2SessionListParams{}, option.WithResponseInto(&raw)); err != nil {
		t.Fatalf("List with caller WithResponseInto: %v", err)
	}
	if raw == nil {
		t.Error("a caller supplied WithResponseInto was not honoured, so it is not being applied after the SDK's")
	}
}
