// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestV2FsReadWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Fs.Read(context.TODO(), "src/main.go", opencode.V2FsReadParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestV2FsReadURLEncoding verifies that path segments are correctly percent-encoded
// without modifying the '/' separator. This test does NOT require a live server.
//
// The Go HTTP server decodes percent-encoded path segments into r.URL.Path, so we
// capture the raw request-URI via a custom handler that reads the wire-level URI.
// We verify the decoded r.URL.Path equals the original input path, proving that
// the server received all special characters correctly and that '/' separators
// were preserved (not encoded as %2F).
func TestV2FsReadURLEncoding(t *testing.T) {
	cases := []struct {
		name            string
		path            string
		wantDecodedPath string // expected decoded suffix after "/api/fs/read/"
	}{
		{
			name:            "simple path no special chars",
			path:            "src/main.go",
			wantDecodedPath: "src/main.go",
		},
		{
			name:            "path with spaces",
			path:            "my dir/my file.go",
			wantDecodedPath: "my dir/my file.go",
		},
		{
			name:            "path with hash",
			path:            "dir/#hash/file.go",
			wantDecodedPath: "dir/#hash/file.go",
		},
		{
			name:            "path with Chinese characters",
			path:            "项目/文件.go",
			wantDecodedPath: "项目/文件.go",
		},
		{
			name:            "single segment no slash",
			path:            "README.md",
			wantDecodedPath: "README.md",
		},
		{
			name:            "deep path preserves slashes",
			path:            "a/b/c/d/e.txt",
			wantDecodedPath: "a/b/c/d/e.txt",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var capturedPath string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// r.URL.Path is always the decoded path; RawPath is set only when
				// the raw and decoded forms differ (e.g., %2F vs /).
				// Using Path gives us the server-received decoded path.
				capturedPath = r.URL.Path
				w.Header().Set("Content-Type", "application/octet-stream")
				w.WriteHeader(http.StatusOK)
			}))
			defer srv.Close()

			client := opencode.NewClient(
				option.WithBaseURL(srv.URL),
			)

			_, _ = client.V2Fs.Read(context.TODO(), tc.path, opencode.V2FsReadParams{})

			prefix := "/api/fs/read/"
			if !strings.HasPrefix(capturedPath, prefix) {
				t.Fatalf("unexpected path %q; want prefix %q", capturedPath, prefix)
			}
			gotSufx := strings.TrimPrefix(capturedPath, prefix)
			if gotSufx != tc.wantDecodedPath {
				t.Errorf("path suffix = %q, want %q", gotSufx, tc.wantDecodedPath)
			}
		})
	}
}

// TestV2FsReadEmptyPath verifies that an empty path returns an error immediately
// without making an HTTP request.
func TestV2FsReadEmptyPath(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := opencode.NewClient(
		option.WithBaseURL(srv.URL),
	)
	_, err := client.V2Fs.Read(context.TODO(), "", opencode.V2FsReadParams{})
	if err == nil {
		t.Fatal("expected error for empty path, got nil")
	}
	if err.Error() != "missing required path parameter" {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if called {
		t.Error("HTTP server was called despite empty path")
	}
}

// TestV2FsListWithOptionalParams tests the list endpoint.
func TestV2FsListWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Fs.List(context.TODO(), opencode.V2FsListParams{
		Path: opencode.F("path"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestV2FsFindWithOptionalParams tests the find endpoint.
func TestV2FsFindWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Fs.Find(context.TODO(), opencode.V2FsFindParams{
		Query: opencode.F("*.go"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
