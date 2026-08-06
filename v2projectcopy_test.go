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

func TestV2ProjectCopyNew(t *testing.T) {
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
	_, err := client.V2ProjectCopy.New(
		context.TODO(),
		"projectID",
		opencode.V2ProjectCopyNewParams{
			Location:  opencode.F(opencode.V2LocationParam{Directory: opencode.F("directory")}),
			Strategy:  opencode.F("strategy"),
			Directory: opencode.F("directory"),
			Name:      opencode.F("name"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ProjectCopyRefresh(t *testing.T) {
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
	err := client.V2ProjectCopy.Refresh(
		context.TODO(),
		"projectID",
		opencode.V2ProjectCopyRefreshParams{
			Location: opencode.F(opencode.V2LocationParam{Directory: opencode.F("directory")}),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ProjectCopyRemove(t *testing.T) {
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
	err := client.V2ProjectCopy.Remove(
		context.TODO(),
		"projectID",
		opencode.V2ProjectCopyRemoveParams{
			Location:  opencode.F(opencode.V2LocationParam{Directory: opencode.F("directory")}),
			Directory: opencode.F("directory"),
			Force:     opencode.F(true),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ProjectCopyRequiresProjectID(t *testing.T) {
	s := opencode.NewV2ProjectCopyService()
	ctx := context.Background()
	if _, err := s.New(ctx, "", opencode.V2ProjectCopyNewParams{}); err == nil {
		t.Error("expected error for empty projectID")
	}
	if err := s.Refresh(ctx, "", opencode.V2ProjectCopyRefreshParams{}); err == nil {
		t.Error("expected error for empty projectID")
	}
	if err := s.Remove(ctx, "", opencode.V2ProjectCopyRemoveParams{}); err == nil {
		t.Error("expected error for empty projectID")
	}
}
