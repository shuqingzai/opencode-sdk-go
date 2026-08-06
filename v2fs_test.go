// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/option"
)

// TestV2FsReadPathEmptyPathReturnsError verifies that V2FsService.ReadPath
// returns an error with message "missing required path parameter" when the path
// argument is empty, consistent with the validation pattern used by all other SDK
// methods that accept a required path parameter.
func TestV2FsReadPathEmptyPathReturnsError(t *testing.T) {
	t.Parallel()
	client := opencode.NewClient(
		option.WithBaseURL("http://localhost:4010"),
	)
	_, err := client.V2Fs.ReadPath(context.Background(), "", opencode.V2FsReadParams{})
	if err == nil {
		t.Fatal("expected error for empty path, got nil")
	}
	const want = "missing required path parameter"
	if err.Error() != want {
		t.Errorf("error message: got %q, want %q", err.Error(), want)
	}
}

// TestV2FsReadPathNonEmptyPathNoValidationError verifies that
// V2FsService.ReadPath does not return a validation error for a non-empty path
// (the request will fail for other reasons in a test environment, but the path
// validation must pass).
func TestV2FsReadPathNonEmptyPathNoValidationError(t *testing.T) {
	t.Parallel()
	client := opencode.NewClient(
		option.WithBaseURL("http://localhost:4010"),
	)
	_, err := client.V2Fs.ReadPath(context.Background(), "some/path/file.txt", opencode.V2FsReadParams{})
	// The request will fail due to no server, but not due to path validation.
	if err != nil && err.Error() == "missing required path parameter" {
		t.Errorf("non-empty path should not trigger validation error, but got: %v", err)
	}
}
