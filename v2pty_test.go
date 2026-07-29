// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

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
	_, err := client.V2Pty.List(context.TODO(), opencode.V2PtyListParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PtyNewWithOptionalParams(t *testing.T) {
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
		Command: opencode.F("bash"),
		Args:    opencode.F([]string{"-l"}),
		Cwd:     opencode.F("/home/user"),
		Title:   opencode.F("My Terminal"),
		Env: opencode.F(map[string]string{
			"TERM": "xterm-256color",
		}),
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
			Workspace: opencode.F("workspace"),
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
	_, err := client.V2Pty.Get(context.TODO(), "ptyID", opencode.V2PtyGetParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2PtyUpdateWithOptionalParams(t *testing.T) {
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
		Title: opencode.F("Updated Title"),
		Size: opencode.F(opencode.V2PtySizeParam{
			Rows: opencode.F(int64(24)),
			Cols: opencode.F(int64(80)),
		}),
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
			Workspace: opencode.F("workspace"),
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
	err := client.V2Pty.Remove(context.TODO(), "ptyID", opencode.V2PtyRemoveParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
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
	_, err := client.V2Pty.ConnectToken(context.TODO(), "ptyID", opencode.V2PtyConnectTokenParams{})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// --- Unit tests (no server required) ---

// TestV2PtyMissingPtyID verifies empty ptyID path validation for all methods.
func TestV2PtyMissingPtyID(t *testing.T) {
	s := opencode.NewV2PtyService()
	ctx := context.Background()

	_, err := s.Get(ctx, "", opencode.V2PtyGetParams{})
	if err == nil {
		t.Error("Get: expected error for empty ptyID")
	}

	_, err = s.Update(ctx, "", opencode.V2PtyUpdateParams{})
	if err == nil {
		t.Error("Update: expected error for empty ptyID")
	}

	err = s.Remove(ctx, "", opencode.V2PtyRemoveParams{})
	if err == nil {
		t.Error("Remove: expected error for empty ptyID")
	}

	_, err = s.Connect(ctx, "", opencode.V2PtyConnectParams{})
	if err == nil {
		t.Error("Connect: expected error for empty ptyID")
	}

	_, err = s.ConnectToken(ctx, "", opencode.V2PtyConnectTokenParams{})
	if err == nil {
		t.Error("ConnectToken: expected error for empty ptyID")
	}
}

// TestV2PtyNewParamsMarshalJSON verifies V2PtyNewParams body serialization.
func TestV2PtyNewParamsMarshalJSON(t *testing.T) {
	params := opencode.V2PtyNewParams{
		Command: opencode.F("bash"),
		Args:    opencode.F([]string{"-l", "--norc"}),
		Cwd:     opencode.F("/home/user"),
		Title:   opencode.F("My Terminal"),
		Env: opencode.F(map[string]string{
			"TERM":      "xterm-256color",
			"COLORTERM": "truecolor",
		}),
	}
	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	if m["command"] != "bash" {
		t.Errorf("expected command=bash, got %v", m["command"])
	}
	if m["title"] != "My Terminal" {
		t.Errorf("expected title=My Terminal, got %v", m["title"])
	}
	// location is a query param, should NOT appear in body JSON
	if _, ok := m["location"]; ok {
		t.Error("location should not appear in body JSON (it's a query param)")
	}
}

// TestV2PtyUpdateParamsMarshalJSON verifies V2PtyUpdateParams body serialization.
func TestV2PtyUpdateParamsMarshalJSON(t *testing.T) {
	params := opencode.V2PtyUpdateParams{
		Title: opencode.F("New Title"),
		Size: opencode.F(opencode.V2PtySizeParam{
			Rows: opencode.F(int64(40)),
			Cols: opencode.F(int64(120)),
		}),
	}
	data, err := params.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	if m["title"] != "New Title" {
		t.Errorf("expected title=New Title, got %v", m["title"])
	}
	sizeVal, ok := m["size"].(map[string]any)
	if !ok {
		t.Fatalf("expected size to be a map, got %T", m["size"])
	}
	// JSON numbers unmarshal as float64
	if sizeVal["rows"] != float64(40) {
		t.Errorf("expected rows=40, got %v", sizeVal["rows"])
	}
	if sizeVal["cols"] != float64(120) {
		t.Errorf("expected cols=120, got %v", sizeVal["cols"])
	}
}

// TestV2PtySizeParamMarshalJSON verifies V2PtySizeParam serialization.
func TestV2PtySizeParamMarshalJSON(t *testing.T) {
	size := opencode.V2PtySizeParam{
		Rows: opencode.F(int64(24)),
		Cols: opencode.F(int64(80)),
	}
	data, err := size.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("JSON invalid: %v", err)
	}
	if m["rows"] != float64(24) {
		t.Errorf("expected rows=24, got %v", m["rows"])
	}
	if m["cols"] != float64(80) {
		t.Errorf("expected cols=80, got %v", m["cols"])
	}
}

// TestPtyUnmarshal verifies Pty response struct deserialization.
func TestPtyUnmarshal(t *testing.T) {
	raw := `{
		"args": ["bash", "-l"],
		"command": "bash",
		"cwd": "/home/user",
		"exitCode": 0,
		"id": "pty-123",
		"pid": 12345,
		"status": "running",
		"title": "My Terminal"
	}`
	var pty opencode.Pty
	if err := pty.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if pty.ID != "pty-123" {
		t.Errorf("expected ID=pty-123, got %q", pty.ID)
	}
	if pty.Command != "bash" {
		t.Errorf("expected Command=bash, got %q", pty.Command)
	}
	if pty.Pid != 12345 {
		t.Errorf("expected Pid=12345, got %d", pty.Pid)
	}
	if pty.Status != opencode.PtyStatusRunning {
		t.Errorf("expected Status=running, got %q", pty.Status)
	}
	if len(pty.Args) != 2 {
		t.Errorf("expected 2 args, got %d", len(pty.Args))
	}
	if pty.JSON.RawJSON() == "" {
		t.Error("expected non-empty RawJSON")
	}
}

// TestPtyTicketConnectTokenUnmarshal verifies PtyTicketConnectToken deserialization.
func TestPtyTicketConnectTokenUnmarshal(t *testing.T) {
	raw := `{"ticket": "tok-abc123", "expires_in": 300}`
	var tok opencode.PtyTicketConnectToken
	if err := tok.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if tok.Ticket != "tok-abc123" {
		t.Errorf("expected Ticket=tok-abc123, got %q", tok.Ticket)
	}
	if tok.ExpiresIn != 300 {
		t.Errorf("expected ExpiresIn=300, got %d", tok.ExpiresIn)
	}
	if tok.JSON.RawJSON() == "" {
		t.Error("expected non-empty RawJSON")
	}
}

// TestV2PtyListResponseUnmarshal verifies V2PtyListResponse deserialization.
func TestV2PtyListResponseUnmarshal(t *testing.T) {
	raw := `{
		"location": {
			"directory": "/home/user",
			"workspaceID": "ws-1"
		},
		"data": [
			{
				"args": [],
				"command": "bash",
				"cwd": "/home/user",
				"id": "pty-1",
				"pid": 1001,
				"status": "running",
				"title": "Shell"
			}
		]
	}`
	var resp opencode.V2PtyListResponse
	if err := resp.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 pty, got %d", len(resp.Data))
	}
	if resp.Data[0].ID != "pty-1" {
		t.Errorf("expected ID=pty-1, got %q", resp.Data[0].ID)
	}
}

// TestV2PtyConnectTokenResponseUnmarshal verifies the envelope response.
func TestV2PtyConnectTokenResponseUnmarshal(t *testing.T) {
	raw := `{
		"location": {
			"directory": "/home/user",
			"workspaceID": "ws-1"
		},
		"data": {
			"ticket": "ticket-xyz",
			"expires_in": 60
		}
	}`
	var resp opencode.V2PtyConnectTokenResponse
	if err := resp.UnmarshalJSON([]byte(raw)); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}
	if resp.Data.Ticket != "ticket-xyz" {
		t.Errorf("expected Ticket=ticket-xyz, got %q", resp.Data.Ticket)
	}
	if resp.Data.ExpiresIn != 60 {
		t.Errorf("expected ExpiresIn=60, got %d", resp.Data.ExpiresIn)
	}
}
