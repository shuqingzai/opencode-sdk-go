// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestV2PtyList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.V2Pty.List(context.TODO(), opencode.V2PtyListParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PtyCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.V2Pty.New(context.TODO(), opencode.V2PtyNewParams{
		Command: opencode.F("cmd"),
		Args:    opencode.F([]string{"arg1"}),
		Cwd:     opencode.F("/tmp"),
		Title:   opencode.F("title"),
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PtyGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.V2Pty.Get(context.TODO(), "ptyID", opencode.V2PtyGetParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PtyUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.V2Pty.Update(context.TODO(), "ptyID", opencode.V2PtyUpdateParams{
		Title: opencode.F("newTitle"),
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PtyRemove(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.V2Pty.Remove(context.TODO(), "ptyID", opencode.V2PtyRemoveParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PtyConnect(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.V2Pty.Connect(context.TODO(), "ptyID", opencode.V2PtyConnectParams{
		Cursor: opencode.F("cursor"),
		Ticket: opencode.F("ticket"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestV2PtySizeAlias verifies that V2PtySize is a type alias for PtySize so
// that callers using either name interchangeably are not broken by the renaming
// introduced in fix #16.
func TestV2PtySizeAlias(t *testing.T) {
	t.Parallel()
	// Compile-time alias check: V2PtySize must be assignable to/from PtySize.
	// PtySize is a Request type so fields are param.Field[int64].
	psize := opencode.PtySize{
		Rows: opencode.F[int64](24),
		Cols: opencode.F[int64](80),
	}
	// V2PtySize is a type alias (=) so the following assignment is valid at compile time.
	var v2size opencode.V2PtySize = psize
	if v2size.Rows.Value != 24 || v2size.Cols.Value != 80 {
		t.Errorf("V2PtySize alias: got Rows=%d Cols=%d, want 24 80", v2size.Rows.Value, v2size.Cols.Value)
	}
}

func TestV2PtyConnectToken(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.V2Pty.ConnectToken(context.TODO(), "ptyID", opencode.V2PtyConnectTokenParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
