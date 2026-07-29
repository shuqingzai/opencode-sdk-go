// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// ===== Prism / live-server tests =====

func TestV2SessionRevertStage(t *testing.T) {
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
	_, err := client.V2Session.Revert.Stage(context.TODO(), "sessionID", opencode.V2SessionRevertStageParams{
		MessageID: opencode.F("msg_abc"),
		Files:     opencode.F(true),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2SessionRevertClear(t *testing.T) {
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
	err := client.V2Session.Revert.Clear(context.TODO(), "sessionID")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2SessionRevertCommit(t *testing.T) {
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
	err := client.V2Session.Revert.Commit(context.TODO(), "sessionID")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// ===== Missing path parameter unit tests =====

func TestV2SessionRevertStageEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionRevertService()
	_, err := svc.Stage(context.Background(), "", opencode.V2SessionRevertStageParams{
		MessageID: opencode.F("msg_1"),
	})
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionRevertClearEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionRevertService()
	err := svc.Clear(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func TestV2SessionRevertCommitEmptySessionID(t *testing.T) {
	svc := opencode.NewV2SessionRevertService()
	err := svc.Commit(context.Background(), "")
	if err == nil {
		t.Fatal("expected error for empty sessionID")
	}
	if !strings.Contains(err.Error(), "missing required sessionID parameter") {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

// ===== Request serialization tests =====

// V2SessionRevertStageParams: required messageID, optional files (bool).
func TestV2SessionRevertStageParamsSerialization(t *testing.T) {
	t.Run("required messageID only", func(t *testing.T) {
		p := opencode.V2SessionRevertStageParams{
			MessageID: opencode.F("msg_required"),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(data)
		if !strings.Contains(got, `"messageID":"msg_required"`) {
			t.Errorf("messageID missing: %s", got)
		}
	})

	t.Run("with optional files=true", func(t *testing.T) {
		p := opencode.V2SessionRevertStageParams{
			MessageID: opencode.F("msg_with_files"),
			Files:     opencode.F(true),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(data)
		if !strings.Contains(got, `"messageID":"msg_with_files"`) {
			t.Errorf("messageID missing: %s", got)
		}
		if !strings.Contains(got, `"files":true`) {
			t.Errorf("files missing: %s", got)
		}
	})

	t.Run("with optional files=false", func(t *testing.T) {
		p := opencode.V2SessionRevertStageParams{
			MessageID: opencode.F("msg_no_files"),
			Files:     opencode.F(false),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(data)
		if !strings.Contains(got, `"messageID":"msg_no_files"`) {
			t.Errorf("messageID missing: %s", got)
		}
	})
}

// ===== Response deserialization tests =====

// V2SessionRevertStageResponse: data is RevertState with messageID + files.
func TestV2SessionRevertStageResponseUnmarshal(t *testing.T) {
	raw := `{
		"data": {
			"messageID": "msg_stage_1",
			"files": [
				{
					"path": "/src/main.go",
					"status": "modified",
					"additions": 5,
					"deletions": 2,
					"patch": "@@ -1,3 +1,6 @@\n+line"
				},
				{
					"path": "/src/new.go",
					"status": "added",
					"additions": 10,
					"deletions": 0,
					"patch": "+++ /dev/null"
				}
			]
		}
	}`
	var resp opencode.V2SessionRevertStageResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Data.MessageID != "msg_stage_1" {
		t.Errorf("Data.MessageID = %q", resp.Data.MessageID)
	}
	if len(resp.Data.Files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(resp.Data.Files))
	}
	f0 := resp.Data.Files[0]
	if f0.Path != "/src/main.go" {
		t.Errorf("Files[0].Path = %q", f0.Path)
	}
	if f0.Status != "modified" {
		t.Errorf("Files[0].Status = %q", f0.Status)
	}
	if f0.Additions != 5 {
		t.Errorf("Files[0].Additions = %d", f0.Additions)
	}
	if f0.Deletions != 2 {
		t.Errorf("Files[0].Deletions = %d", f0.Deletions)
	}
	f1 := resp.Data.Files[1]
	if f1.Status != "added" {
		t.Errorf("Files[1].Status = %q", f1.Status)
	}
	if resp.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved on response")
	}
	if resp.Data.JSON.RawJSON() == "" {
		t.Error("RawJSON not preserved on RevertState")
	}
}

// RevertState: empty files array is valid.
func TestV2SessionRevertStageResponseEmptyFiles(t *testing.T) {
	raw := `{
		"data": {
			"messageID": "msg_empty",
			"files": []
		}
	}`
	var resp opencode.V2SessionRevertStageResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Data.MessageID != "msg_empty" {
		t.Errorf("Data.MessageID = %q", resp.Data.MessageID)
	}
	if len(resp.Data.Files) != 0 {
		t.Errorf("expected 0 files, got %d", len(resp.Data.Files))
	}
}

// RevertState: no files field present (null / omitted).
func TestV2SessionRevertStageResponseNilFiles(t *testing.T) {
	raw := `{
		"data": {
			"messageID": "msg_no_files"
		}
	}`
	var resp opencode.V2SessionRevertStageResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Data.MessageID != "msg_no_files" {
		t.Errorf("Data.MessageID = %q", resp.Data.MessageID)
	}
}

// RevertState: extra/unknown fields should not cause errors.
func TestV2SessionRevertStageResponseExtraFields(t *testing.T) {
	raw := `{
		"data": {
			"messageID": "msg_extra",
			"files": [],
			"futureField": "some value"
		},
		"meta": "extra"
	}`
	var resp opencode.V2SessionRevertStageResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("should tolerate extra fields: %v", err)
	}
	if resp.Data.MessageID != "msg_extra" {
		t.Errorf("Data.MessageID = %q", resp.Data.MessageID)
	}
}
