package opencode

import (
	"encoding/json"
	"testing"
)

// TestPromptAdmittedDeliveryIsEnumTyped verifies that the `delivery` field of the
// session.next.prompt.admitted event is modelled as a Go enum type rather than a
// bare string.
//
// OpenAPI fact: SessionNextPromptAdmitted.data.delivery is
// {"type":"string","enum":["steer","queue"]} and is listed in data.required.
// JS SDK v2 fact: types.gen.ts EventSessionNextPromptAdmitted.properties.delivery
// is `"steer" | "queue"`.
//
// The field was previously `Delivery string`, which lost the IsKnown() contract
// that the sibling session.next.prompted event, the /api/event surface and the
// sync surface all already provided. Regression test for that gap.
func TestPromptAdmittedDeliveryIsEnumTyped(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		want EventListResponseEventSessionNextPromptAdmittedDelivery
	}{
		{"steer", EventListResponseEventSessionNextPromptAdmittedDeliverySteer},
		{"queue", EventListResponseEventSessionNextPromptAdmittedDeliveryQueue},
	} {
		t.Run(tc.raw, func(t *testing.T) {
			data := []byte(`{
				"timestamp": 1730000000,
				"sessionID": "ses_abc",
				"messageID": "msg_abc",
				"prompt": {"parts": []},
				"delivery": "` + tc.raw + `"
			}`)

			var props EventListResponseEventSessionNextPromptAdmittedProperties
			if err := json.Unmarshal(data, &props); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if props.Delivery != tc.want {
				t.Errorf("Delivery = %q, want %q", props.Delivery, tc.want)
			}
			if !props.Delivery.IsKnown() {
				t.Errorf("IsKnown() = false for spec-declared value %q", tc.raw)
			}
			if props.Timestamp != 1730000000 {
				t.Errorf("Timestamp = %d, want 1730000000", props.Timestamp)
			}
			if props.SessionID != "ses_abc" || props.MessageID != "msg_abc" {
				t.Errorf("ID fields lost: sessionID=%q messageID=%q", props.SessionID, props.MessageID)
			}
		})
	}
}

// TestPromptAdmittedDeliveryUnknownValueIsForwardCompatible verifies an
// unrecognised server value still decodes (no error, value preserved) but reports
// IsKnown() == false, so a future spec addition never breaks existing clients.
func TestPromptAdmittedDeliveryUnknownValueIsForwardCompatible(t *testing.T) {
	data := []byte(`{
		"timestamp": 1,
		"sessionID": "ses_a",
		"messageID": "msg_a",
		"prompt": {"parts": []},
		"delivery": "some-future-mode"
	}`)

	var props EventListResponseEventSessionNextPromptAdmittedProperties
	if err := json.Unmarshal(data, &props); err != nil {
		t.Fatalf("unmarshal must not fail on unknown enum value: %v", err)
	}
	if got := string(props.Delivery); got != "some-future-mode" {
		t.Errorf("Delivery = %q, want raw value preserved", got)
	}
	if props.Delivery.IsKnown() {
		t.Error("IsKnown() = true for a value absent from the OpenAPI enum")
	}
	if props.JSON.RawJSON() == "" {
		t.Error("RawJSON() empty — raw payload not retained")
	}
}

// TestCompactionStartedReasonIsEnumTyped verifies that the `reason` field of the
// session.next.compaction.started event uses the
// EventListResponseEventSessionNextCompactionStartedReason enum.
//
// OpenAPI fact: SessionNextCompactionStarted.data.reason is
// {"type":"string","enum":["auto","manual"]} and is listed in data.required.
// JS SDK v2 fact: types.gen.ts EventSessionNextCompactionStarted.properties.reason
// is `"auto" | "manual"`.
//
// The enum type already existed but the struct field was a bare string, leaving
// the enum unreferenced (dead code) and the field unvalidated.
func TestCompactionStartedReasonIsEnumTyped(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		want EventListResponseEventSessionNextCompactionStartedReason
	}{
		{"auto", EventListResponseEventSessionNextCompactionStartedReasonAuto},
		{"manual", EventListResponseEventSessionNextCompactionStartedReasonManual},
	} {
		t.Run(tc.raw, func(t *testing.T) {
			data := []byte(`{
				"timestamp": 1730000001,
				"messageID": "msg_c",
				"sessionID": "ses_c",
				"reason": "` + tc.raw + `"
			}`)

			var props EventListResponseEventSessionNextCompactionStartedProperties
			if err := json.Unmarshal(data, &props); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if props.Reason != tc.want {
				t.Errorf("Reason = %q, want %q", props.Reason, tc.want)
			}
			if !props.Reason.IsKnown() {
				t.Errorf("IsKnown() = false for spec-declared value %q", tc.raw)
			}
		})
	}
}

// TestCompactionStartedReasonUnknownValueIsForwardCompatible mirrors the delivery
// forward-compatibility contract for `reason`.
func TestCompactionStartedReasonUnknownValueIsForwardCompatible(t *testing.T) {
	data := []byte(`{"timestamp":1,"messageID":"m","sessionID":"s","reason":"scheduled"}`)

	var props EventListResponseEventSessionNextCompactionStartedProperties
	if err := json.Unmarshal(data, &props); err != nil {
		t.Fatalf("unmarshal must not fail on unknown enum value: %v", err)
	}
	if got := string(props.Reason); got != "scheduled" {
		t.Errorf("Reason = %q, want raw value preserved", got)
	}
	if props.Reason.IsKnown() {
		t.Error("IsKnown() = true for a value absent from the OpenAPI enum")
	}
}

// TestEnumParityAcrossEventSurfaces asserts the three SSE surfaces agree on the
// enum value sets for the two fields fixed above. OpenAPI defines these events
// once and reuses them across /event, /global/event and /api/event, so the Go
// enums must not drift apart.
func TestEnumParityAcrossEventSurfaces(t *testing.T) {
	t.Run("delivery", func(t *testing.T) {
		for _, v := range []string{"steer", "queue"} {
			eventSurface := EventListResponseEventSessionNextPromptAdmittedDelivery(v)
			v2Surface := V2EventSessionNextPromptAdmittedDelivery(v)
			syncSurface := SyncEventSessionNextPromptAdmittedDelivery(v)
			promptedSurface := EventListResponseEventSessionNextPromptedDelivery(v)

			if !eventSurface.IsKnown() || !v2Surface.IsKnown() ||
				!syncSurface.IsKnown() || !promptedSurface.IsKnown() {
				t.Errorf("value %q not known on all four surfaces: "+
					"event=%v v2=%v sync=%v prompted=%v", v,
					eventSurface.IsKnown(), v2Surface.IsKnown(),
					syncSurface.IsKnown(), promptedSurface.IsKnown())
			}
		}
		// A value outside the spec enum must be rejected uniformly.
		if EventListResponseEventSessionNextPromptAdmittedDelivery("nope").IsKnown() ||
			V2EventSessionNextPromptAdmittedDelivery("nope").IsKnown() {
			t.Error("out-of-spec delivery value reported as known")
		}
	})

	t.Run("reason", func(t *testing.T) {
		for _, v := range []string{"auto", "manual"} {
			if !EventListResponseEventSessionNextCompactionStartedReason(v).IsKnown() {
				t.Errorf("/event surface: %q not known", v)
			}
			if !V2EventSessionNextCompactionStartedReason(v).IsKnown() {
				t.Errorf("/api/event surface: %q not known", v)
			}
		}
		if EventListResponseEventSessionNextCompactionStartedReason("nope").IsKnown() ||
			V2EventSessionNextCompactionStartedReason("nope").IsKnown() {
			t.Error("out-of-spec reason value reported as known")
		}
	})
}

// TestPermissionRulesetNamedTypeDecodes verifies that Session.Permission and
// GlobalSession.Permission — both $ref PermissionRuleset in OpenAPI — decode
// through the shared PermissionRuleset named type.
func TestPermissionRulesetNamedTypeDecodes(t *testing.T) {
	rules := `[{"type":"bash","pattern":["rm -rf *"],"effect":"deny"}]`

	t.Run("Session", func(t *testing.T) {
		var s Session
		raw := []byte(`{"id":"ses_1","slug":"s","projectID":"p","directory":"/d",
			"title":"t","version":"1","time":{"created":1,"updated":2},
			"permission":` + rules + `}`)
		if err := json.Unmarshal(raw, &s); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		var want PermissionRuleset = s.Permission // must compile as the named type
		if len(want) != 1 {
			t.Fatalf("len(Permission) = %d, want 1", len(want))
		}
	})

	t.Run("GlobalSession", func(t *testing.T) {
		var gs GlobalSession
		raw := []byte(`{"id":"ses_2","permission":` + rules + `}`)
		if err := json.Unmarshal(raw, &gs); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		var want PermissionRuleset = gs.Permission // must compile as the named type
		if len(want) != 1 {
			t.Fatalf("len(Permission) = %d, want 1", len(want))
		}
	})
}
