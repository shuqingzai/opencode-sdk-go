// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
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
