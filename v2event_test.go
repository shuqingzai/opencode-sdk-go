// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
)

func TestV2EventSessionNextCompactionEndedReasonIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2EventSessionNextCompactionEndedReason
		want  bool
	}{
		// known values (OpenAPI enum)
		{name: "auto", value: opencode.V2EventSessionNextCompactionEndedReasonAuto, want: true},
		{name: "manual", value: opencode.V2EventSessionNextCompactionEndedReasonManual, want: true},
		// unknown values
		{name: "unknown", value: "__unknown__", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}

func TestV2EventSessionNextCompactionStartedReasonIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2EventSessionNextCompactionStartedReason
		want  bool
	}{
		// known values (OpenAPI enum)
		{name: "auto", value: opencode.V2EventSessionNextCompactionStartedReasonAuto, want: true},
		{name: "manual", value: opencode.V2EventSessionNextCompactionStartedReasonManual, want: true},
		// unknown values
		{name: "unknown", value: "__unknown__", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}

func TestV2EventSessionNextPromptAdmittedDeliveryIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2EventSessionNextPromptAdmittedDelivery
		want  bool
	}{
		// known values (OpenAPI enum)
		{name: "steer", value: opencode.V2EventSessionNextPromptAdmittedDeliverySteer, want: true},
		{name: "queue", value: opencode.V2EventSessionNextPromptAdmittedDeliveryQueue, want: true},
		// unknown values
		{name: "unknown", value: "__unknown__", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}

func TestV2EventSessionNextPromptedDeliveryIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2EventSessionNextPromptedDelivery
		want  bool
	}{
		// known values (OpenAPI enum)
		{name: "steer", value: opencode.V2EventSessionNextPromptedDeliverySteer, want: true},
		{name: "queue", value: opencode.V2EventSessionNextPromptedDeliveryQueue, want: true},
		// unknown values
		{name: "unknown", value: "__unknown__", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}

func TestV2EventTuiToastShowVariantIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2EventTuiToastShowVariant
		want  bool
	}{
		// known values (OpenAPI enum)
		{name: "info", value: opencode.V2EventTuiToastShowVariantInfo, want: true},
		{name: "success", value: opencode.V2EventTuiToastShowVariantSuccess, want: true},
		{name: "warning", value: opencode.V2EventTuiToastShowVariantWarning, want: true},
		{name: "error", value: opencode.V2EventTuiToastShowVariantError, want: true},
		// unknown values
		{name: "unknown", value: "__unknown__", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}

func TestV2EventWorkspaceStatusStatusIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2EventWorkspaceStatusStatus
		want  bool
	}{
		// known values (OpenAPI enum)
		{name: "connected", value: opencode.V2EventWorkspaceStatusStatusConnected, want: true},
		{name: "connecting", value: opencode.V2EventWorkspaceStatusStatusConnecting, want: true},
		{name: "disconnected", value: opencode.V2EventWorkspaceStatusStatusDisconnected, want: true},
		{name: "error", value: opencode.V2EventWorkspaceStatusStatusError, want: true},
		// unknown values
		{name: "unknown", value: "__unknown__", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}

func TestV2EventTuiCommandExecuteCommandIsKnown(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name  string
		value opencode.V2EventTuiCommandExecuteCommand
		want  bool
	}{
		// known values (OpenAPI enum)
		{name: "session.list", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionList, want: true},
		{name: "session.new", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionNew, want: true},
		{name: "session.share", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionShare, want: true},
		{name: "session.interrupt", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionInterrupt, want: true},
		{name: "session.compact", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionCompact, want: true},
		{name: "session.page.up", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionPageUp, want: true},
		{name: "session.page.down", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionPageDown, want: true},
		{name: "session.line.up", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionLineUp, want: true},
		{name: "session.line.down", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionLineDown, want: true},
		{name: "session.half.page.up", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionHalfPageUp, want: true},
		{name: "session.half.page.down", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionHalfPageDown, want: true},
		{name: "session.first", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionFirst, want: true},
		{name: "session.last", value: opencode.V2EventTuiCommandExecuteCommandCommandSessionLast, want: true},
		{name: "prompt.clear", value: opencode.V2EventTuiCommandExecuteCommandCommandPromptClear, want: true},
		{name: "prompt.submit", value: opencode.V2EventTuiCommandExecuteCommandCommandPromptSubmit, want: true},
		{name: "agent.cycle", value: opencode.V2EventTuiCommandExecuteCommandCommandAgentCycle, want: true},
		// unknown values
		{name: "unknown", value: "__unknown__", want: false},
		{name: "empty", value: "", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.value.IsKnown()
			if got != tc.want {
				t.Errorf("IsKnown(%q): got %v, want %v", tc.value, got, tc.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Union field correctness tests (fixes for "lying comment" defects)
// ---------------------------------------------------------------------------

// TestV2EventSessionErrorDataErrorUnion verifies that V2EventSessionErrorData.Error
// is decoded into the correct concrete type for each anyOf variant defined in
// the OpenAPI spec (SessionError.data.error anyOf).
func TestV2EventSessionErrorDataErrorUnion(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		rawJSON  string
		assertFn func(t *testing.T, d opencode.V2EventSessionErrorData)
	}{
		{
			name:    "ProviderAuthError",
			rawJSON: `{"sessionID":"ses_1","error":{"name":"ProviderAuthError","data":{"providerID":"openai","message":"bad key"}}}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionErrorData) {
				t.Helper()
				e, ok := d.Error.(opencode.ProviderAuthError)
				if !ok {
					t.Fatalf("Error: got %T, want opencode.ProviderAuthError", d.Error)
				}
				if e.Data.ProviderID != "openai" {
					t.Errorf("ProviderID: got %q, want %q", e.Data.ProviderID, "openai")
				}
			},
		},
		{
			name:    "UnknownError",
			rawJSON: `{"sessionID":"ses_1","error":{"name":"UnknownError","data":{"message":"oops","ref":"r1"}}}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionErrorData) {
				t.Helper()
				e, ok := d.Error.(opencode.UnknownError)
				if !ok {
					t.Fatalf("Error: got %T, want opencode.UnknownError", d.Error)
				}
				if e.Data.Message != "oops" {
					t.Errorf("Message: got %q, want %q", e.Data.Message, "oops")
				}
			},
		},
		{
			name:    "APIError",
			rawJSON: `{"sessionID":"ses_1","error":{"name":"APIError","data":{"message":"rate limited","isRetryable":true,"statusCode":429}}}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionErrorData) {
				t.Helper()
				e, ok := d.Error.(opencode.APIError)
				if !ok {
					t.Fatalf("Error: got %T, want opencode.APIError", d.Error)
				}
				if !e.Data.IsRetryable {
					t.Error("IsRetryable: got false, want true")
				}
			},
		},
		{
			name:    "MessageOutputLengthError",
			rawJSON: `{"sessionID":"ses_1","error":{"name":"MessageOutputLengthError","data":{}}}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionErrorData) {
				t.Helper()
				_, ok := d.Error.(opencode.MessageOutputLengthError)
				if !ok {
					t.Fatalf("Error: got %T, want opencode.MessageOutputLengthError", d.Error)
				}
			},
		},
		{
			name:    "null_error_field",
			rawJSON: `{"sessionID":"ses_1","error":null}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionErrorData) {
				t.Helper()
				if d.Error != nil {
					t.Errorf("Error: got %v (%T), want nil", d.Error, d.Error)
				}
			},
		},
		{
			name:    "missing_error_field",
			rawJSON: `{"sessionID":"ses_1"}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionErrorData) {
				t.Helper()
				if d.Error != nil {
					t.Errorf("Error: got %v (%T), want nil", d.Error, d.Error)
				}
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var d opencode.V2EventSessionErrorData
			if err := json.Unmarshal([]byte(tc.rawJSON), &d); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			tc.assertFn(t, d)
		})
	}
}

// TestV2EventQuestionAskedDataTool verifies that the Tool field decodes into a
// concrete QuestionTool (not map[string]any) per OpenAPI QuestionAsked.data.tool.
func TestV2EventQuestionAskedDataTool(t *testing.T) {
	t.Parallel()
	t.Run("present", func(t *testing.T) {
		t.Parallel()
		raw := `{"id":"que_1","sessionID":"ses_1","questions":[],"tool":{"messageID":"msg_1","callID":"call_abc"}}`
		var d opencode.V2EventQuestionAskedData
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if d.Tool.MessageID != "msg_1" {
			t.Errorf("Tool.MessageID: got %q, want %q", d.Tool.MessageID, "msg_1")
		}
		if d.Tool.CallID != "call_abc" {
			t.Errorf("Tool.CallID: got %q, want %q", d.Tool.CallID, "call_abc")
		}
	})
	t.Run("missing_tool", func(t *testing.T) {
		t.Parallel()
		raw := `{"id":"que_1","sessionID":"ses_1","questions":[]}`
		var d opencode.V2EventQuestionAskedData
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if d.Tool.MessageID != "" || d.Tool.CallID != "" {
			t.Errorf("expected zero-value Tool, got %+v", d.Tool)
		}
	})
}

// TestV2EventSessionNextPromptAdmittedDataPrompt verifies that the Prompt field
// decodes into a concrete V2SessionInputPrompt per OpenAPI SessionNextPromptAdmitted.data.prompt.
func TestV2EventSessionNextPromptAdmittedDataPrompt(t *testing.T) {
	t.Parallel()
	t.Run("with_text", func(t *testing.T) {
		t.Parallel()
		raw := `{"delivery":"queue","messageID":"msg_1","prompt":{"text":"hello world","files":[]},"sessionID":"ses_1","timestamp":1234567890}`
		var d opencode.V2EventSessionNextPromptAdmittedData
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if d.Prompt.Text != "hello world" {
			t.Errorf("Prompt.Text: got %q, want %q", d.Prompt.Text, "hello world")
		}
	})
	t.Run("with_agents", func(t *testing.T) {
		t.Parallel()
		raw := `{"delivery":"steer","messageID":"msg_2","prompt":{"text":"fix it","agents":[{"agentID":"ag_1"}]},"sessionID":"ses_2","timestamp":9876543210}`
		var d opencode.V2EventSessionNextPromptAdmittedData
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if d.Prompt.Text != "fix it" {
			t.Errorf("Prompt.Text: got %q, want %q", d.Prompt.Text, "fix it")
		}
		if len(d.Prompt.Agents) != 1 {
			t.Errorf("Prompt.Agents: got len=%d, want 1", len(d.Prompt.Agents))
		}
	})
}

// TestV2EventSessionNextPromptedDataPrompt verifies that the Prompt field
// decodes into a concrete V2SessionInputPrompt per OpenAPI SessionNextPrompted.data.prompt.
func TestV2EventSessionNextPromptedDataPrompt(t *testing.T) {
	t.Parallel()
	raw := `{"delivery":"steer","messageID":"msg_2","prompt":{"text":"do the thing"},"sessionID":"ses_3","timestamp":5555}`
	var d opencode.V2EventSessionNextPromptedData
	if err := json.Unmarshal([]byte(raw), &d); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if d.Prompt.Text != "do the thing" {
		t.Errorf("Prompt.Text: got %q, want %q", d.Prompt.Text, "do the thing")
	}
}

// TestV2EventSessionNextRetriedDataError verifies that the Error field decodes
// into a concrete EventListResponseEventSessionNextRetriedError per OpenAPI
// SessionNextRetried.data.error ($ref: SessionNextRetry_error).
func TestV2EventSessionNextRetriedDataError(t *testing.T) {
	t.Parallel()
	t.Run("retryable", func(t *testing.T) {
		t.Parallel()
		raw := `{"attempt":2,"error":{"message":"timeout","isRetryable":true,"statusCode":503},"sessionID":"ses_3","timestamp":111222333}`
		var d opencode.V2EventSessionNextRetriedData
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if d.Error.Message != "timeout" {
			t.Errorf("Error.Message: got %q, want %q", d.Error.Message, "timeout")
		}
		if !d.Error.IsRetryable {
			t.Error("Error.IsRetryable: got false, want true")
		}
	})
	t.Run("not_retryable", func(t *testing.T) {
		t.Parallel()
		raw := `{"attempt":1,"error":{"message":"bad request","isRetryable":false},"sessionID":"ses_4","timestamp":999}`
		var d opencode.V2EventSessionNextRetriedData
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if d.Error.IsRetryable {
			t.Error("Error.IsRetryable: got true, want false")
		}
	})
}

// TestV2EventSessionErrorDataAsError verifies that AsError() returns the typed
// union value.
func TestV2EventSessionErrorDataAsError(t *testing.T) {
	t.Parallel()
	raw := `{"sessionID":"ses_1","error":{"name":"ContextOverflowError","data":{"message":"context too long"}}}`
	var d opencode.V2EventSessionErrorData
	if err := json.Unmarshal([]byte(raw), &d); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	u := d.AsError()
	if u == nil {
		t.Fatal("AsError(): got nil, want non-nil")
	}
	e, ok := u.(opencode.ContextOverflowError)
	if !ok {
		t.Fatalf("AsError() type: got %T, want opencode.ContextOverflowError", u)
	}
	if e.Data.Message != "context too long" {
		t.Errorf("ContextOverflowError.Data.Message: got %q, want %q", e.Data.Message, "context too long")
	}
}

// TestV2EventSessionStatusDataStatusUnion verifies that
// V2EventSessionStatusData.Status decodes into the correct concrete variant for
// each SessionStatus anyOf member (idle/retry/busy), both through the public
// `any` field and through the AsStatus() accessor.
func TestV2EventSessionStatusDataStatusUnion(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		rawJSON  string
		assertFn func(t *testing.T, d opencode.V2EventSessionStatusData)
	}{
		{
			name:    "idle",
			rawJSON: `{"sessionID":"ses_1","status":{"type":"idle"}}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionStatusData) {
				t.Helper()
				if _, ok := d.Status.(opencode.SessionStatusIdle); !ok {
					t.Fatalf("Status: got %T, want opencode.SessionStatusIdle", d.Status)
				}
				if _, ok := d.AsStatus().(opencode.SessionStatusIdle); !ok {
					t.Fatalf("AsStatus(): got %T, want opencode.SessionStatusIdle", d.AsStatus())
				}
			},
		},
		{
			name:    "retry",
			rawJSON: `{"sessionID":"ses_1","status":{"type":"retry","attempt":2,"message":"busy","next":3,"action":{"reason":"rate limited","provider":"openai","title":"retry","message":"try again","label":"Retry"}}}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionStatusData) {
				t.Helper()
				s, ok := d.Status.(opencode.SessionStatusRetry)
				if !ok {
					t.Fatalf("Status: got %T, want opencode.SessionStatusRetry", d.Status)
				}
				if s.Attempt != 2 {
					t.Errorf("Status.Attempt: got %d, want 2", s.Attempt)
				}
				if s.Next != 3 {
					t.Errorf("Status.Next: got %d, want 3", s.Next)
				}
				if s.Action.Provider != "openai" {
					t.Errorf("Status.Action.Provider: got %q, want %q", s.Action.Provider, "openai")
				}
				if _, ok := d.AsStatus().(opencode.SessionStatusRetry); !ok {
					t.Fatalf("AsStatus(): got %T, want opencode.SessionStatusRetry", d.AsStatus())
				}
			},
		},
		{
			name:    "busy",
			rawJSON: `{"sessionID":"ses_1","status":{"type":"busy"}}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionStatusData) {
				t.Helper()
				if _, ok := d.Status.(opencode.SessionStatusBusy); !ok {
					t.Fatalf("Status: got %T, want opencode.SessionStatusBusy", d.Status)
				}
				if _, ok := d.AsStatus().(opencode.SessionStatusBusy); !ok {
					t.Fatalf("AsStatus(): got %T, want opencode.SessionStatusBusy", d.AsStatus())
				}
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var d opencode.V2EventSessionStatusData
			if err := json.Unmarshal([]byte(tc.rawJSON), &d); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			if d.SessionID != "ses_1" {
				t.Errorf("SessionID: got %q, want %q", d.SessionID, "ses_1")
			}
			tc.assertFn(t, d)
		})
	}
}

// TestV2EventSessionStatusDataStatusRawJSON verifies that the raw JSON metadata
// is preserved after the shadow-routed decode.
func TestV2EventSessionStatusDataStatusRawJSON(t *testing.T) {
	t.Parallel()
	raw := `{"sessionID":"ses_1","status":{"type":"busy"}}`
	var d opencode.V2EventSessionStatusData
	if err := json.Unmarshal([]byte(raw), &d); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if got := d.JSON.RawJSON(); got != raw {
		t.Errorf("RawJSON(): got %s, want %s", got, raw)
	}
}

// TestV2EventTuiCommandExecuteDataCommand verifies that Command decodes into
// the concrete string enum type (known values and arbitrary strings).
func TestV2EventTuiCommandExecuteDataCommand(t *testing.T) {
	t.Parallel()
	t.Run("known_value", func(t *testing.T) {
		t.Parallel()
		raw := `{"command":"prompt.submit"}`
		var d opencode.V2EventTuiCommandExecuteData
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if d.Command != opencode.V2EventTuiCommandExecuteCommandCommandPromptSubmit {
			t.Errorf("Command: got %q, want %q", d.Command, opencode.V2EventTuiCommandExecuteCommandCommandPromptSubmit)
		}
		if !d.Command.IsKnown() {
			t.Errorf("Command.IsKnown(): got false, want true")
		}
	})
	t.Run("arbitrary_string", func(t *testing.T) {
		t.Parallel()
		raw := `{"command":"custom.command"}`
		var d opencode.V2EventTuiCommandExecuteData
		if err := json.Unmarshal([]byte(raw), &d); err != nil {
			t.Fatalf("json.Unmarshal: %v", err)
		}
		if d.Command != "custom.command" {
			t.Errorf("Command: got %q, want %q", d.Command, "custom.command")
		}
		if d.Command.IsKnown() {
			t.Errorf("Command.IsKnown(): got true, want false")
		}
	})
}
