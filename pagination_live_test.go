package opencode

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/sst/opencode-sdk-go/option"
)

// These tests run against a real opencode server rather than an httptest stub,
// so that the pagination contract is verified end to end against the actual
// handlers. They are skipped unless a server is pointed at explicitly:
//
//	OPENCODE_LIVE_BASE_URL=http://127.0.0.1:4096 \
//	OPENCODE_LIVE_DIRECTORY=/path/to/a/project \
//	go test -count=1 -run Live ./...
//
// OPENCODE_LIVE_DIRECTORY is required by the v1 endpoints, which scope sessions
// to a project directory.

func liveClient(t *testing.T) (*Client, string) {
	t.Helper()
	base := os.Getenv("OPENCODE_LIVE_BASE_URL")
	if base == "" {
		t.Skip("set OPENCODE_LIVE_BASE_URL to run live pagination tests")
	}
	return NewClient(option.WithBaseURL(base)), os.Getenv("OPENCODE_LIVE_DIRECTORY")
}

// distinct reports the first duplicate it finds, which is the failure mode that
// matters most for cursor paging: a cursor that does not advance replays rows.
func distinct(t *testing.T, label string, ids []string) {
	t.Helper()
	seen := make(map[string]int, len(ids))
	for i, id := range ids {
		if prev, dup := seen[id]; dup {
			t.Errorf("%s: %q appears at both index %d and %d, the cursor replayed a row", label, id, prev, i)
			return
		}
		seen[id] = i
	}
}

// Sessions: GET /session cannot paginate, so the paginating equivalent is
// GET /experimental/session. This walks the whole list in small pages and
// requires it to match a single large unpaginated read exactly.
func TestLiveExperimentalSessionListPaginates(t *testing.T) {
	client, directory := liveClient(t)
	ctx := context.Background()

	baseline, err := client.Experimental.Session.List(ctx, ExperimentalSessionListParams{
		Directory: F(directory),
		Limit:     Int(500),
	})
	if err != nil {
		t.Fatalf("baseline List: %v", err)
	}
	var want []string
	for _, s := range baseline.Data {
		want = append(want, s.ID)
	}
	if len(want) < 2 {
		t.Skipf("server only has %d sessions, need at least 2 to exercise paging", len(want))
	}
	t.Logf("baseline: %d sessions in one unpaginated read", len(want))

	iter := client.Experimental.Session.ListAutoPaging(ctx, ExperimentalSessionListParams{
		Directory: F(directory),
		Limit:     Int(3),
	})
	var got []string
	for iter.Next() {
		got = append(got, iter.Current().ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}
	t.Logf("paged: %d sessions collected in pages of 3", len(got))

	distinct(t, "sessions", got)
	if len(got) != len(want) {
		t.Fatalf("paged walk saw %d sessions, unpaginated read saw %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("session %d = %q, want %q (paging changed the order)", i, got[i], want[i])
		}
	}
}

// Sessions, v2: GET /api/session pages on an opaque body cursor. There is no
// unpaginated read to compare against here, because every response is a page.
// Instead this asserts the property that actually matters: the page size must
// not change what a full walk yields.
func TestLiveV2SessionListPaginates(t *testing.T) {
	client, _ := liveClient(t)
	ctx := context.Background()

	walk := func(limit int64) []string {
		iter := client.V2Session.ListAutoPaging(ctx, V2SessionListParams{Limit: Int(limit)})
		var ids []string
		for iter.Next() {
			ids = append(ids, iter.Current().ID)
		}
		if err := iter.Err(); err != nil {
			t.Fatalf("auto pager (limit=%d): %v", limit, err)
		}
		return ids
	}

	probe, err := client.V2Session.List(ctx, V2SessionListParams{Limit: Int(2)})
	if err != nil {
		t.Skipf("v2 session list unavailable on this server: %v", err)
	}
	if len(probe.Data) < 2 {
		t.Skipf("server only has %d v2 sessions, need at least 2", len(probe.Data))
	}

	wide := walk(1000)
	t.Logf("wide walk: %d v2 sessions in pages of 1000", len(wide))
	narrow := walk(50)
	t.Logf("narrow walk: %d v2 sessions in pages of 50", len(narrow))

	distinct(t, "v2 sessions (wide)", wide)
	distinct(t, "v2 sessions (narrow)", narrow)

	if len(wide) != len(narrow) {
		t.Fatalf("page size changed the result: %d sessions with limit=1000 but %d with limit=50", len(wide), len(narrow))
	}
	for i := range wide {
		if wide[i] != narrow[i] {
			t.Fatalf("v2 session %d differs by page size: %q with limit=1000, %q with limit=50", i, wide[i], narrow[i])
		}
	}

	// The walk must also agree with the first page it started from.
	for i, s := range probe.Data {
		if wide[i] != s.ID {
			t.Fatalf("walk item %d = %q, but a direct read returned %q", i, wide[i], s.ID)
		}
	}
}

// Messages: GET /session/{id}/message pages on the X-Next-Cursor header, sent
// back as `before`. Omitting Limit returns the whole conversation, which gives
// an independent baseline to compare the paged walk against.
func TestLiveSessionMessagesPaginate(t *testing.T) {
	client, directory := liveClient(t)
	ctx := context.Background()

	sessionID := liveBusiestSession(t, client, directory)

	baseline, err := client.Session.Messages(ctx, sessionID, SessionMessagesParams{
		Directory: F(directory),
	})
	if err != nil {
		t.Fatalf("baseline Messages: %v", err)
	}
	var want []string
	for _, m := range baseline.Data {
		want = append(want, m.Info.ID)
	}
	if len(want) < 2 {
		t.Skipf("session %s only has %d messages", sessionID, len(want))
	}
	if got := baseline.NextCursor(); got != "" {
		t.Errorf("an unlimited read advertised cursor %q, want none", got)
	}
	t.Logf("baseline: session %s has %d messages in one unpaginated read", sessionID, len(want))

	iter := client.Session.MessagesAutoPaging(ctx, sessionID, SessionMessagesParams{
		Directory: F(directory),
		Limit:     Int(3),
	})
	var got []string
	for iter.Next() {
		got = append(got, iter.Current().Info.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}
	t.Logf("paged: %d messages collected in pages of 3", len(got))

	distinct(t, "messages", got)
	if len(got) != len(want) {
		t.Fatalf("paged walk saw %d messages, unpaginated read saw %d", len(got), len(want))
	}
	// Each page is ordered oldest first while pages walk backwards in time, so a
	// full walk is a permutation of the conversation rather than the same order.
	wantSet := make(map[string]bool, len(want))
	for _, id := range want {
		wantSet[id] = true
	}
	for _, id := range got {
		if !wantSet[id] {
			t.Errorf("paged walk produced %q, which the unpaginated read never returned", id)
		}
	}
}

// Messages, v2: GET /api/session/{id}/message pages on the body cursor.
func TestLiveV2SessionMessagesPaginate(t *testing.T) {
	client, _ := liveClient(t)
	ctx := context.Background()

	page, err := client.V2Session.List(ctx, V2SessionListParams{Limit: Int(200)})
	if err != nil {
		t.Skipf("v2 session list unavailable on this server: %v", err)
	}

	scanned, broken := 0, 0
	for page != nil {
		for _, session := range page.Data {
			scanned++
			baseline, err := client.V2Session.Messages(ctx, session.ID, V2SessionMessagesParams{Limit: Int(200)})
			if err != nil {
				// Some historical sessions cannot be projected into v2 messages and
				// make the server answer 500. That is a server side problem, not a
				// paging one, so keep looking for a session that does project.
				broken++
				continue
			}
			if len(baseline.Data) < 4 {
				continue
			}
			t.Logf("baseline: v2 session %s has %d messages", session.ID, len(baseline.Data))

			iter := client.V2Session.MessagesAutoPaging(ctx, session.ID, V2SessionMessagesParams{Limit: Int(2)})
			count := 0
			for iter.Next() {
				count++
			}
			if err := iter.Err(); err != nil {
				t.Fatalf("auto pager: %v", err)
			}
			t.Logf("paged: %d v2 messages collected in pages of 2", count)
			if count != len(baseline.Data) {
				t.Fatalf("paged walk saw %d messages, unpaginated read saw %d", count, len(baseline.Data))
			}
			return
		}
		if scanned >= 200 {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			t.Fatalf("GetNextPage while scanning for a populated session: %v", err)
		}
	}
	t.Skipf("scanned %d v2 sessions (%d returned a server error), none had the 4+ projected messages needed to exercise paging", scanned, broken)
}

// GET /session really does not paginate: this pins the live behaviour that
// justifies SessionService.List having no auto-paging variant.
func TestLiveSessionListHasNoPagination(t *testing.T) {
	client, directory := liveClient(t)
	ctx := context.Background()

	var raw *http.Response
	first, err := client.Session.List(ctx, SessionListParams{
		Directory: F(directory),
		Limit:     Int(3),
	}, option.WithResponseInto(&raw))
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(*first) == 0 {
		t.Skip("no sessions on this server")
	}
	if got := raw.Header.Get("X-Next-Cursor"); got != "" {
		t.Fatalf("GET /session advertised X-Next-Cursor %q: it now paginates and the SDK must expose that", got)
	}
	if got := raw.Header.Get("Link"); got != "" {
		t.Fatalf("GET /session advertised Link %q: it now paginates and the SDK must expose that", got)
	}
	t.Logf("GET /session?limit=3 returned %d sessions with no cursor header", len(*first))

	// Growing the limit is the only way to see more, which is what the web client
	// does. Verify it really is a superset that starts with the same rows.
	grown, err := client.Session.List(ctx, SessionListParams{
		Directory: F(directory),
		Limit:     Int(6),
	})
	if err != nil {
		t.Fatalf("List(limit=6): %v", err)
	}
	if len(*grown) < len(*first) {
		t.Fatalf("raising the limit returned fewer sessions: %d then %d", len(*first), len(*grown))
	}
	for i := range *first {
		if (*first)[i].ID != (*grown)[i].ID {
			t.Fatalf("session %d changed from %q to %q when the limit grew", i, (*first)[i].ID, (*grown)[i].ID)
		}
	}
	t.Logf("raising limit to 6 returned %d sessions, a stable superset of the first %d", len(*grown), len(*first))
}

// liveBusiestSession picks the session with the most messages so that the paging
// test has something real to walk. Set OPENCODE_LIVE_SESSION_ID to pin a
// specific session instead, which is useful when reproducing a request captured
// from the browser.
func liveBusiestSession(t *testing.T, client *Client, directory string) string {
	t.Helper()
	if pinned := os.Getenv("OPENCODE_LIVE_SESSION_ID"); pinned != "" {
		t.Logf("using pinned session %s", pinned)
		return pinned
	}
	ctx := context.Background()
	sessions, err := client.Session.List(ctx, SessionListParams{
		Directory: F(directory),
		Limit:     Int(10),
	})
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	best, bestCount := "", 0
	for _, session := range *sessions {
		page, err := client.Session.Messages(ctx, session.ID, SessionMessagesParams{
			Directory: F(directory),
		})
		if err != nil {
			continue
		}
		if len(page.Data) > bestCount {
			best, bestCount = session.ID, len(page.Data)
		}
	}
	if best == "" {
		t.Skip("no session with messages on this server")
	}
	return best
}

// The practical answer to "I need sessions to paginate": GET /experimental/session
// returns the same rows in the same order as GET /session, but with a working
// cursor. This proves the substitution is exact rather than merely plausible, so
// callers who need paging can switch endpoints without changing what they see.
func TestLiveExperimentalSessionListMatchesSessionList(t *testing.T) {
	client, directory := liveClient(t)
	ctx := context.Background()

	// /session cannot page, so read it with a limit large enough to hold
	// everything and use it as the reference ordering.
	plain, err := client.Session.List(ctx, SessionListParams{
		Directory: F(directory),
		Limit:     Int(1000),
	})
	if err != nil {
		t.Fatalf("Session.List: %v", err)
	}
	if len(*plain) < 2 {
		t.Skipf("only %d sessions in %s", len(*plain), directory)
	}

	iter := client.Experimental.Session.ListAutoPaging(ctx, ExperimentalSessionListParams{
		Directory: F(directory),
		Limit:     Int(3),
	})
	var paged []GlobalSession
	for iter.Next() {
		paged = append(paged, iter.Current())
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("auto pager: %v", err)
	}

	t.Logf("GET /session returned %d sessions; paged GET /experimental/session returned %d", len(*plain), len(paged))
	if len(paged) != len(*plain) {
		t.Fatalf("paged experimental walk saw %d sessions, /session saw %d", len(paged), len(*plain))
	}
	for i := range *plain {
		want, got := (*plain)[i], paged[i]
		if got.ID != want.ID {
			t.Fatalf("session %d = %q, want %q: the two endpoints disagree on ordering", i, got.ID, want.ID)
		}
		if got.Title != want.Title {
			t.Errorf("session %s title = %q, want %q", want.ID, got.Title, want.Title)
		}
		if got.ParentID != want.ParentID {
			t.Errorf("session %s parentID = %q, want %q", want.ID, got.ParentID, want.ParentID)
		}
		if got.Time.Updated != want.Time.Updated {
			t.Errorf("session %s time.updated = %d, want %d", want.ID, got.Time.Updated, want.Time.Updated)
		}
	}
	t.Logf("all %d sessions match on id, title, parentID and time.updated", len(paged))
}
