// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
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
// V2EventSessionStatusData.Status decodes into the SessionStatus carrier struct
// and routes to the correct concrete variant for each SessionStatus anyOf
// member (idle/retry/busy), both through the flattened carrier fields and
// through the AsUnion() accessor.
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
				if d.Status.Type != opencode.SessionStatusTypeIdle {
					t.Errorf("Status.Type: got %q, want %q", d.Status.Type, opencode.SessionStatusTypeIdle)
				}
				if !d.Status.Type.IsKnown() {
					t.Errorf("Status.Type %q is not known", d.Status.Type)
				}
				v, ok := d.Status.AsUnion().(opencode.SessionStatusIdle)
				if !ok {
					t.Fatalf("Status.AsUnion(): got %T, want opencode.SessionStatusIdle", d.Status.AsUnion())
				}
				if v.Type != opencode.SessionStatusIdleTypeIdle {
					t.Errorf("variant Type: got %q, want %q", v.Type, opencode.SessionStatusIdleTypeIdle)
				}
			},
		},
		{
			name:    "retry",
			rawJSON: `{"sessionID":"ses_1","status":{"type":"retry","attempt":2,"message":"busy","next":3,"action":{"reason":"rate limited","provider":"openai","title":"retry","message":"try again","label":"Retry"}}}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionStatusData) {
				t.Helper()
				if d.Status.Type != opencode.SessionStatusTypeRetry {
					t.Errorf("Status.Type: got %q, want %q", d.Status.Type, opencode.SessionStatusTypeRetry)
				}
				// retry 专属字段必须 port 到载体结构体上
				if d.Status.Attempt != 2 {
					t.Errorf("Status.Attempt: got %d, want 2", d.Status.Attempt)
				}
				if d.Status.Next != 3 {
					t.Errorf("Status.Next: got %d, want 3", d.Status.Next)
				}
				if d.Status.Action.Provider != "openai" {
					t.Errorf("Status.Action.Provider: got %q, want %q", d.Status.Action.Provider, "openai")
				}

				s, ok := d.Status.AsUnion().(opencode.SessionStatusRetry)
				if !ok {
					t.Fatalf("Status.AsUnion(): got %T, want opencode.SessionStatusRetry", d.Status.AsUnion())
				}
				if s.Type != opencode.SessionStatusRetryTypeRetry {
					t.Errorf("variant Type: got %q, want %q", s.Type, opencode.SessionStatusRetryTypeRetry)
				}
				if s.Attempt != 2 {
					t.Errorf("variant Attempt: got %d, want 2", s.Attempt)
				}
				if s.Next != 3 {
					t.Errorf("variant Next: got %d, want 3", s.Next)
				}
				if s.Action.Provider != "openai" {
					t.Errorf("variant Action.Provider: got %q, want %q", s.Action.Provider, "openai")
				}
			},
		},
		{
			name:    "busy",
			rawJSON: `{"sessionID":"ses_1","status":{"type":"busy"}}`,
			assertFn: func(t *testing.T, d opencode.V2EventSessionStatusData) {
				t.Helper()
				if d.Status.Type != opencode.SessionStatusTypeBusy {
					t.Errorf("Status.Type: got %q, want %q", d.Status.Type, opencode.SessionStatusTypeBusy)
				}
				if !d.Status.Type.IsKnown() {
					t.Errorf("Status.Type %q is not known", d.Status.Type)
				}
				v, ok := d.Status.AsUnion().(opencode.SessionStatusBusy)
				if !ok {
					t.Fatalf("Status.AsUnion(): got %T, want opencode.SessionStatusBusy", d.Status.AsUnion())
				}
				if v.Type != opencode.SessionStatusBusyTypeBusy {
					t.Errorf("variant Type: got %q, want %q", v.Type, opencode.SessionStatusBusyTypeBusy)
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
// is preserved on both the data struct and the nested SessionStatus carrier.
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
	if got, want := d.Status.JSON.RawJSON(), `{"type":"busy"}`; got != want {
		t.Errorf("Status.JSON.RawJSON(): got %s, want %s", got, want)
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

// =============================================================================
// A1/A2 SSE 容错回归护栏：V2Event.Data 字段级畸形载体（string/number/bool/array/
// null）必须静默降级不报错（ssestream.Stream.Next() 中 json.Unmarshal 出错即
// return false，整个流永久终止），且 JSON.Data 元数据/RawJSON() 完整保留原始值。
// =============================================================================

// TestV2EventDataMalformedCarrierRawJSONPreserved 对 V2Event.Data 喂入
// string/number/bool/array/null 五类畸形载体（顶层 type 仍合法，故 union 仍按
// type 路由到 V2EventFileEdited；嵌套的 Data 子结构自身无法从标量解出字段，
// 静默降级为零值，绝不报错）。断言：
//   - json.Unmarshal 不报错（err == nil，否则会终结整条 SSE 流）
//   - 顶层 JSON.RawJSON() 完整保留原始报文（不丢失任何字节）
//   - 路由后的 variant 的 Data 子结构自身 JSON 元数据（RawJSON()）也保留了
//     原始畸形标量，供调用方分辨降级原因
func TestV2EventDataMalformedCarrierRawJSONPreserved(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		data string
	}{
		{"string", `"not-an-object"`},
		{"number", `42`},
		{"bool", `true`},
		{"array", `[1,2,3]`},
		{"null", `null`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			raw := `{"id":"evt_1","type":"file.edited","data":` + tc.data + `}`
			var ev opencode.V2Event
			if err := json.Unmarshal([]byte(raw), &ev); err != nil {
				t.Fatalf("json.Unmarshal: %v (would terminate the SSE stream via ssestream.Stream.Next())", err)
			}
			if got := ev.JSON.RawJSON(); got != raw {
				t.Errorf("JSON.RawJSON() = %q, want %q (raw payload must be fully preserved)", got, raw)
			}
			fe, ok := ev.AsUnion().(opencode.V2EventFileEdited)
			if !ok {
				t.Fatalf("AsUnion() = %T, want V2EventFileEdited (type-based routing must still succeed)", ev.AsUnion())
			}
			if fe.Data.File != "" {
				t.Errorf("Data.File = %q, want zero-value for malformed scalar data carrier", fe.Data.File)
			}
			// "null" is treated like an absent Data sub-object (nothing to
			// port), so its nested JSON metadata stays empty/missing; every
			// other scalar carrier's raw bytes are preserved verbatim.
			wantNestedRaw := tc.data
			if tc.name == "null" {
				wantNestedRaw = ""
			}
			if got := fe.Data.JSON.RawJSON(); got != wantNestedRaw {
				t.Errorf("Data.JSON.RawJSON() = %q, want %q (nested raw scalar must be preserved)", got, wantNestedRaw)
			}
		})
	}
}

// TestEventListResponseMalformedCarrierRawJSONPreserved 是 V2Event 侧的对照
// 实验：EventListResponse.Properties 对同样的畸形载体必须表现一致（三条 SSE 链路
// 容错策略必须一致，参见 A1 报告「SSE 容错铁律实测」）：不报错、顶层与嵌套的
// RawJSON() 均完整保留原始值。
func TestEventListResponseMalformedCarrierRawJSONPreserved(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		data string
	}{
		{"string", `"not-an-object"`},
		{"number", `42`},
		{"bool", `true`},
		{"array", `[1,2,3]`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			raw := `{"id":"evt_1","type":"file.edited","properties":` + tc.data + `}`
			var ev opencode.EventListResponse
			if err := json.Unmarshal([]byte(raw), &ev); err != nil {
				t.Fatalf("json.Unmarshal: %v (would terminate the SSE stream via ssestream.Stream.Next())", err)
			}
			if got := ev.JSON.RawJSON(); got != raw {
				t.Errorf("JSON.RawJSON() = %q, want %q (raw payload must be fully preserved)", got, raw)
			}
			fe, ok := ev.AsUnion().(opencode.EventListResponseEventFileEdited)
			if !ok {
				t.Fatalf("AsUnion() = %T, want EventListResponseEventFileEdited (type-based routing must still succeed)", ev.AsUnion())
			}
			if fe.Properties.File != "" {
				t.Errorf("Properties.File = %q, want zero-value for malformed scalar properties carrier", fe.Properties.File)
			}
			if got := fe.Properties.JSON.RawJSON(); got != tc.data {
				t.Errorf("Properties.JSON.RawJSON() = %q, want %q (nested raw scalar must be preserved)", got, tc.data)
			}
		})
	}
}

// =============================================================================
// 端到端 SSE 流：伪造 text/event-stream，中间夹一条字段级畸形事件，断言
// stream.Err() == nil 且全部事件均送达（畸形事件不得杀死整条流）。
// newSSEResponse / sseEventFrame 复用 session_status_sse_test.go 中已有的
// package-level 测试 helper（同 opencode_test 包）。
// =============================================================================

// TestV2EventE2EStreamSurvivesMalformedDataCarrier 验证 V2EventService.ListStreaming
// 返回的 *ssestream.Stream[V2Event] 在中间夹了一条 data 字段为畸形标量的事件时，
// 流不会提前终止，前后的合法事件都能正确送达并路由。
func TestV2EventE2EStreamSurvivesMalformedDataCarrier(t *testing.T) {
	t.Parallel()
	resp := newSSEResponse(
		sseEventFrame([]byte(`{"id":"evt_1","type":"file.edited","data":{"file":"/a.go"}}`)),
		sseEventFrame([]byte(`{"id":"evt_2","type":"file.edited","data":42}`)),
		sseEventFrame([]byte(`{"id":"evt_3","type":"file.edited","data":{"file":"/b.go"}}`)),
	)
	stream := ssestream.NewStream[opencode.V2Event](ssestream.NewDecoder(resp), nil)

	var events []opencode.V2Event
	for stream.Next() {
		events = append(events, stream.Current())
	}
	if err := stream.Err(); err != nil {
		t.Fatalf("stream terminated with error: %v", err)
	}
	if len(events) != 3 {
		t.Fatalf("got %d events, want 3 (malformed data carrier must not kill the stream)", len(events))
	}
	first, ok := events[0].AsUnion().(opencode.V2EventFileEdited)
	if !ok {
		t.Fatalf("events[0].AsUnion() = %T, want V2EventFileEdited", events[0].AsUnion())
	}
	if first.Data.File != "/a.go" {
		t.Errorf("events[0].Data.File = %q, want /a.go", first.Data.File)
	}
	// 畸形事件本身仍按顶层 type 路由到 V2EventFileEdited（字段级降级，非整体
	// 失败）：Data.File 静默降级为零值，但 Data.JSON.RawJSON() 保留原始标量 42。
	mid, ok := events[1].AsUnion().(opencode.V2EventFileEdited)
	if !ok {
		t.Fatalf("events[1].AsUnion() = %T, want V2EventFileEdited", events[1].AsUnion())
	}
	if mid.Data.File != "" {
		t.Errorf("events[1].Data.File = %q, want zero-value (malformed data carrier)", mid.Data.File)
	}
	if got := mid.Data.JSON.RawJSON(); got != "42" {
		t.Errorf("events[1].Data.JSON.RawJSON() = %q, want %q", got, "42")
	}
	// 🔴 关键断言：畸形事件之后的真实事件必须仍能送达并正确解码
	last, ok := events[2].AsUnion().(opencode.V2EventFileEdited)
	if !ok {
		t.Fatalf("events[2].AsUnion() = %T, want V2EventFileEdited", events[2].AsUnion())
	}
	if last.Data.File != "/b.go" {
		t.Errorf("events[2].Data.File = %q, want /b.go", last.Data.File)
	}
}

// TestEventListResponseE2EStreamSurvivesMalformedPropertiesCarrier 同上，验证
// EventListResponse（/event 链路）在畸形 properties 载体前后的事件均能送达，
// 与 V2Event 链路的容错策略保持一致。
func TestEventListResponseE2EStreamSurvivesMalformedPropertiesCarrier(t *testing.T) {
	t.Parallel()
	resp := newSSEResponse(
		sseEventFrame([]byte(`{"id":"evt_1","type":"file.edited","properties":{"file":"/a.go"}}`)),
		sseEventFrame([]byte(`{"id":"evt_2","type":"file.edited","properties":"not-an-object"}`)),
		sseEventFrame([]byte(`{"id":"evt_3","type":"file.edited","properties":{"file":"/b.go"}}`)),
	)
	stream := ssestream.NewStream[opencode.EventListResponse](ssestream.NewDecoder(resp), nil)

	var events []opencode.EventListResponse
	for stream.Next() {
		events = append(events, stream.Current())
	}
	if err := stream.Err(); err != nil {
		t.Fatalf("stream terminated with error: %v", err)
	}
	if len(events) != 3 {
		t.Fatalf("got %d events, want 3 (malformed properties carrier must not kill the stream)", len(events))
	}
	first, ok := events[0].AsUnion().(opencode.EventListResponseEventFileEdited)
	if !ok {
		t.Fatalf("events[0].AsUnion() = %T, want EventListResponseEventFileEdited", events[0].AsUnion())
	}
	if first.Properties.File != "/a.go" {
		t.Errorf("events[0].Properties.File = %q, want /a.go", first.Properties.File)
	}
	// 畸形事件本身仍按顶层 type 路由（字段级降级，非整体失败）：Properties.File
	// 静默降级为零值，但 Properties.JSON.RawJSON() 保留原始标量。
	mid, ok := events[1].AsUnion().(opencode.EventListResponseEventFileEdited)
	if !ok {
		t.Fatalf("events[1].AsUnion() = %T, want EventListResponseEventFileEdited", events[1].AsUnion())
	}
	if mid.Properties.File != "" {
		t.Errorf("events[1].Properties.File = %q, want zero-value (malformed properties carrier)", mid.Properties.File)
	}
	if got := mid.Properties.JSON.RawJSON(); got != `"not-an-object"` {
		t.Errorf("events[1].Properties.JSON.RawJSON() = %q, want %q", got, `"not-an-object"`)
	}
	// 🔴 关键断言：畸形事件之后的真实事件必须仍能送达并正确解码
	last, ok := events[2].AsUnion().(opencode.EventListResponseEventFileEdited)
	if !ok {
		t.Fatalf("events[2].AsUnion() = %T, want EventListResponseEventFileEdited", events[2].AsUnion())
	}
	if last.Properties.File != "/b.go" {
		t.Errorf("events[2].Properties.File = %q, want /b.go", last.Properties.File)
	}
}
