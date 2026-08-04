// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/shared"
)

// TestAPIErrorDataUnmarshalAllFields verifies that APIErrorData correctly
// deserialises a JSON payload containing all 6 fields, including the 4 optional
// ones (statusCode, responseHeaders, responseBody, metadata).
// The optional fields must NOT use omitempty per the Response struct rule.
func TestAPIErrorDataUnmarshalAllFields(t *testing.T) {
	t.Parallel()

	raw := `{
		"message": "request failed",
		"isRetryable": true,
		"statusCode": 503,
		"responseHeaders": {"content-type": "application/json", "x-request-id": "req_123"},
		"responseBody": "{\"error\":\"service unavailable\"}",
		"metadata": {"region": "us-east-1", "trace": "t_abc"}
	}`

	var got shared.APIErrorData
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if got.Message != "request failed" {
		t.Errorf("Message: got %q, want %q", got.Message, "request failed")
	}
	if !got.IsRetryable {
		t.Error("IsRetryable: got false, want true")
	}
	if got.StatusCode != 503 {
		t.Errorf("StatusCode: got %d, want 503", got.StatusCode)
	}
	if got.ResponseHeaders["content-type"] != "application/json" {
		t.Errorf("ResponseHeaders[content-type]: got %q, want application/json", got.ResponseHeaders["content-type"])
	}
	if got.ResponseHeaders["x-request-id"] != "req_123" {
		t.Errorf("ResponseHeaders[x-request-id]: got %q, want req_123", got.ResponseHeaders["x-request-id"])
	}
	if got.ResponseBody != `{"error":"service unavailable"}` {
		t.Errorf("ResponseBody: got %q", got.ResponseBody)
	}
	if got.Metadata["region"] != "us-east-1" {
		t.Errorf("Metadata[region]: got %q, want us-east-1", got.Metadata["region"])
	}
	if got.Metadata["trace"] != "t_abc" {
		t.Errorf("Metadata[trace]: got %q, want t_abc", got.Metadata["trace"])
	}
}

// TestAPIErrorDataUnmarshalRequiredOnly verifies that APIErrorData correctly
// deserialises a minimal payload with only the required fields (message, isRetryable).
// The optional fields must unmarshal to their zero values without error.
func TestAPIErrorDataUnmarshalRequiredOnly(t *testing.T) {
	t.Parallel()

	raw := `{"message": "not retryable", "isRetryable": false}`

	var got shared.APIErrorData
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if got.Message != "not retryable" {
		t.Errorf("Message: got %q, want %q", got.Message, "not retryable")
	}
	if got.IsRetryable {
		t.Error("IsRetryable: got true, want false")
	}
	if got.StatusCode != 0 {
		t.Errorf("StatusCode: got %d, want 0 for absent field", got.StatusCode)
	}
	if got.ResponseBody != "" {
		t.Errorf("ResponseBody: got %q, want empty string for absent field", got.ResponseBody)
	}
	if got.ResponseHeaders != nil {
		t.Errorf("ResponseHeaders: got non-nil %v, want nil for absent field", got.ResponseHeaders)
	}
	if got.Metadata != nil {
		t.Errorf("Metadata: got non-nil %v, want nil for absent field", got.Metadata)
	}
}

// TestV2AgentListResponseUnmarshal verifies full round-trip deserialization of
// V2AgentListResponse including nested types.
func TestV2AgentListResponseUnmarshal(t *testing.T) {
	t.Parallel()

	raw := `{
		"location": {
			"directory": "/workspace/proj",
			"project": { "id": "proj_a", "directory": "/workspace/proj" }
		},
		"data": [
			{
				"id": "agent_001",
				"request": { "headers": {"x-key": "val"}, "body": {} },
				"mode": "primary",
				"hidden": false,
				"permissions": [
					{"action": "read", "resource": "file", "effect": "allow"}
				],
				"color": "primary",
				"steps": 10
			}
		]
	}`

	var got opencode.V2AgentListResponse
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if got.Location.Directory != "/workspace/proj" {
		t.Errorf("Location.Directory: got %q, want /workspace/proj", got.Location.Directory)
	}
	if len(got.Data) != 1 {
		t.Fatalf("Data length: got %d, want 1", len(got.Data))
	}
	agent := got.Data[0]
	if agent.ID != "agent_001" {
		t.Errorf("agent.ID: got %q, want agent_001", agent.ID)
	}
	if agent.Mode != opencode.V2AgentInfoModePrimary {
		t.Errorf("agent.Mode: got %q, want primary", agent.Mode)
	}
	if agent.Hidden {
		t.Error("agent.Hidden: got true, want false")
	}
	if agent.Steps != 10 {
		t.Errorf("agent.Steps: got %d, want 10", agent.Steps)
	}
	if len(agent.Permissions) != 1 {
		t.Fatalf("agent.Permissions length: got %d, want 1", len(agent.Permissions))
	}
	perm := agent.Permissions[0]
	if perm.Action != "read" {
		t.Errorf("perm.Action: got %q, want read", perm.Action)
	}
	if perm.Resource != "file" {
		t.Errorf("perm.Resource: got %q, want file", perm.Resource)
	}
	if perm.Effect != opencode.PermissionV2EffectAllow {
		t.Errorf("perm.Effect: got %q, want allow", perm.Effect)
	}
	color := agent.AsColor()
	if color != opencode.AgentColorPrimary {
		t.Errorf("agent.AsColor(): got %q, want primary", color)
	}
}

// TestUnionBoolUnionIntAliases verifies that the package-level aliases for
// shared.UnionBool and shared.UnionInt are correctly exposed.
func TestUnionBoolUnionIntAliases(t *testing.T) {
	t.Parallel()

	// Ensure type aliases compile and are usable without importing shared sub-package.
	var b opencode.UnionBool = true
	var i opencode.UnionInt = 42

	_ = b
	_ = i

	// The aliases must implement the expected interface method.
	var _ interface{ ImplementsConfigProviderOptionsTimeoutUnion() } = b
	var _ interface{ ImplementsConfigProviderOptionsTimeoutUnion() } = i
}
