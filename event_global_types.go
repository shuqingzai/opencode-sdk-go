// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"github.com/sst/opencode-sdk-go/internal/apijson"
)

type EventListResponseEventMessageUpdated struct {
	Properties EventListResponseEventMessageUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventMessageUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventMessageUpdatedJSON       `json:"-"`
}

type eventListResponseEventMessageUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessageUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessageUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMessageUpdated) implementsEventListResponse() {}

func (r EventListResponseEventMessageUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventMessageUpdatedProperties struct {
	SessionID string `json:"sessionID,required"`
	Info Message `json:"info,required"`
	JSON      eventListResponseEventMessageUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventMessageUpdatedPropertiesJSON struct {
	SessionID   apijson.Field
	Info   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessageUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessageUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMessageUpdatedType string

const (
	EventListResponseEventMessageUpdatedTypeMessageUpdated EventListResponseEventMessageUpdatedType = "message.updated"
)

func (r EventListResponseEventMessageUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMessageUpdatedTypeMessageUpdated:
		return true
	}
	return false
}

type EventListResponseEventMessageRemoved struct {
	Properties EventListResponseEventMessageRemovedProperties `json:"properties,required"`
	Type       EventListResponseEventMessageRemovedType       `json:"type,required"`
	JSON       eventListResponseEventMessageRemovedJSON       `json:"-"`
}

type eventListResponseEventMessageRemovedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessageRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessageRemovedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMessageRemoved) implementsEventListResponse() {}

func (r EventListResponseEventMessageRemoved) implementsGlobalEventPayload() {}

type EventListResponseEventMessageRemovedProperties struct {
	MessageID string `json:"messageID,required"`
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventMessageRemovedPropertiesJSON `json:"-"`
}

type eventListResponseEventMessageRemovedPropertiesJSON struct {
	MessageID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessageRemovedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessageRemovedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMessageRemovedType string

const (
	EventListResponseEventMessageRemovedTypeMessageRemoved EventListResponseEventMessageRemovedType = "message.removed"
)

func (r EventListResponseEventMessageRemovedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMessageRemovedTypeMessageRemoved:
		return true
	}
	return false
}

type EventListResponseEventMessagePartUpdated struct {
	Properties EventListResponseEventMessagePartUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventMessagePartUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventMessagePartUpdatedJSON       `json:"-"`
}

type eventListResponseEventMessagePartUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessagePartUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessagePartUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMessagePartUpdated) implementsEventListResponse() {}

func (r EventListResponseEventMessagePartUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventMessagePartUpdatedProperties struct {
	Part Part `json:"part,required"`
	SessionID string `json:"sessionID,required"`
	Time int64 `json:"time,required"`
	JSON      eventListResponseEventMessagePartUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventMessagePartUpdatedPropertiesJSON struct {
	Part   apijson.Field
	SessionID   apijson.Field
	Time   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessagePartUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessagePartUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMessagePartUpdatedType string

const (
	EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated EventListResponseEventMessagePartUpdatedType = "message.part.updated"
)

func (r EventListResponseEventMessagePartUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMessagePartUpdatedTypeMessagePartUpdated:
		return true
	}
	return false
}

type EventListResponseEventMessagePartRemoved struct {
	Properties EventListResponseEventMessagePartRemovedProperties `json:"properties,required"`
	Type       EventListResponseEventMessagePartRemovedType       `json:"type,required"`
	JSON       eventListResponseEventMessagePartRemovedJSON       `json:"-"`
}

type eventListResponseEventMessagePartRemovedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessagePartRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessagePartRemovedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMessagePartRemoved) implementsEventListResponse() {}

func (r EventListResponseEventMessagePartRemoved) implementsGlobalEventPayload() {}

type EventListResponseEventMessagePartRemovedProperties struct {
	MessageID string `json:"messageID,required"`
	PartID string `json:"partID,required"`
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventMessagePartRemovedPropertiesJSON `json:"-"`
}

type eventListResponseEventMessagePartRemovedPropertiesJSON struct {
	MessageID   apijson.Field
	PartID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessagePartRemovedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessagePartRemovedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMessagePartRemovedType string

const (
	EventListResponseEventMessagePartRemovedTypeMessagePartRemoved EventListResponseEventMessagePartRemovedType = "message.part.removed"
)

func (r EventListResponseEventMessagePartRemovedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMessagePartRemovedTypeMessagePartRemoved:
		return true
	}
	return false
}

type EventListResponseEventSessionCreated struct {
	Properties EventListResponseEventSessionCreatedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionCreatedType       `json:"type,required"`
	JSON       eventListResponseEventSessionCreatedJSON       `json:"-"`
}

type eventListResponseEventSessionCreatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionCreatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionCreated) implementsEventListResponse() {}

func (r EventListResponseEventSessionCreated) implementsGlobalEventPayload() {}

type EventListResponseEventSessionCreatedProperties struct {
	SessionID string `json:"sessionID,required"`
	Info Session `json:"info,required"`
	JSON      eventListResponseEventSessionCreatedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionCreatedPropertiesJSON struct {
	SessionID   apijson.Field
	Info   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionCreatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionCreatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionCreatedType string

const (
	EventListResponseEventSessionCreatedTypeSessionCreated EventListResponseEventSessionCreatedType = "session.created"
)

func (r EventListResponseEventSessionCreatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionCreatedTypeSessionCreated:
		return true
	}
	return false
}

type EventListResponseEventSessionUpdated struct {
	Properties EventListResponseEventSessionUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventSessionUpdatedJSON       `json:"-"`
}

type eventListResponseEventSessionUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionUpdated) implementsEventListResponse() {}

func (r EventListResponseEventSessionUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventSessionUpdatedProperties struct {
	SessionID string `json:"sessionID,required"`
	Info Session `json:"info,required"`
	JSON      eventListResponseEventSessionUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionUpdatedPropertiesJSON struct {
	SessionID   apijson.Field
	Info   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionUpdatedType string

const (
	EventListResponseEventSessionUpdatedTypeSessionUpdated EventListResponseEventSessionUpdatedType = "session.updated"
)

func (r EventListResponseEventSessionUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionUpdatedTypeSessionUpdated:
		return true
	}
	return false
}

type EventListResponseEventSessionDeleted struct {
	Properties EventListResponseEventSessionDeletedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionDeletedType       `json:"type,required"`
	JSON       eventListResponseEventSessionDeletedJSON       `json:"-"`
}

type eventListResponseEventSessionDeletedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionDeletedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionDeleted) implementsEventListResponse() {}

func (r EventListResponseEventSessionDeleted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionDeletedProperties struct {
	SessionID string `json:"sessionID,required"`
	Info Session `json:"info,required"`
	JSON      eventListResponseEventSessionDeletedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionDeletedPropertiesJSON struct {
	SessionID   apijson.Field
	Info   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionDeletedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionDeletedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionDeletedType string

const (
	EventListResponseEventSessionDeletedTypeSessionDeleted EventListResponseEventSessionDeletedType = "session.deleted"
)

func (r EventListResponseEventSessionDeletedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionDeletedTypeSessionDeleted:
		return true
	}
	return false
}

type EventListResponseEventSessionNextAgentSwitched struct {
	Properties EventListResponseEventSessionNextAgentSwitchedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextAgentSwitchedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextAgentSwitchedJSON       `json:"-"`
}

type eventListResponseEventSessionNextAgentSwitchedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextAgentSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextAgentSwitchedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextAgentSwitched) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextAgentSwitched) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextAgentSwitchedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Agent string `json:"agent,required"`
	JSON      eventListResponseEventSessionNextAgentSwitchedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextAgentSwitchedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Agent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextAgentSwitchedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextAgentSwitchedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextAgentSwitchedType string

const (
	EventListResponseEventSessionNextAgentSwitchedTypeSessionNextAgentSwitched EventListResponseEventSessionNextAgentSwitchedType = "session.next.agent.switched"
)

func (r EventListResponseEventSessionNextAgentSwitchedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextAgentSwitchedTypeSessionNextAgentSwitched:
		return true
	}
	return false
}

type EventListResponseEventSessionNextModelSwitched struct {
	Properties EventListResponseEventSessionNextModelSwitchedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextModelSwitchedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextModelSwitchedJSON       `json:"-"`
}

type eventListResponseEventSessionNextModelSwitchedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextModelSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextModelSwitchedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextModelSwitched) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextModelSwitched) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextModelSwitchedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Model EventListResponseEventSessionNextModelSwitchedModel `json:"model,required"`
	JSON      eventListResponseEventSessionNextModelSwitchedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextModelSwitchedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Model   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextModelSwitchedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextModelSwitchedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextModelSwitchedType string

const (
	EventListResponseEventSessionNextModelSwitchedTypeSessionNextModelSwitched EventListResponseEventSessionNextModelSwitchedType = "session.next.model.switched"
)

func (r EventListResponseEventSessionNextModelSwitchedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextModelSwitchedTypeSessionNextModelSwitched:
		return true
	}
	return false
}

type EventListResponseEventSessionNextPrompted struct {
	Properties EventListResponseEventSessionNextPromptedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextPromptedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextPromptedJSON       `json:"-"`
}

type eventListResponseEventSessionNextPromptedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextPrompted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextPromptedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextPrompted) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextPrompted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextPromptedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Prompt EventListResponseEventSessionNextPromptedPrompt `json:"prompt,required"`
	JSON      eventListResponseEventSessionNextPromptedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextPromptedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Prompt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextPromptedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextPromptedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextPromptedType string

const (
	EventListResponseEventSessionNextPromptedTypeSessionNextPrompted EventListResponseEventSessionNextPromptedType = "session.next.prompted"
)

func (r EventListResponseEventSessionNextPromptedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextPromptedTypeSessionNextPrompted:
		return true
	}
	return false
}

type EventListResponseEventSessionNextSynthetic struct {
	Properties EventListResponseEventSessionNextSyntheticProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextSyntheticType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextSyntheticJSON       `json:"-"`
}

type eventListResponseEventSessionNextSyntheticJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextSynthetic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextSyntheticJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextSynthetic) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextSynthetic) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextSyntheticProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Text string `json:"text,required"`
	JSON      eventListResponseEventSessionNextSyntheticPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextSyntheticPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextSyntheticProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextSyntheticPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextSyntheticType string

const (
	EventListResponseEventSessionNextSyntheticTypeSessionNextSynthetic EventListResponseEventSessionNextSyntheticType = "session.next.synthetic"
)

func (r EventListResponseEventSessionNextSyntheticType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextSyntheticTypeSessionNextSynthetic:
		return true
	}
	return false
}

type EventListResponseEventSessionNextShellStarted struct {
	Properties EventListResponseEventSessionNextShellStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextShellStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextShellStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextShellStartedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextShellStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextShellStartedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextShellStarted) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextShellStarted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextShellStartedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID string `json:"callID,required"`
	Command string `json:"command,required"`
	JSON      eventListResponseEventSessionNextShellStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextShellStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID   apijson.Field
	Command   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextShellStartedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextShellStartedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextShellStartedType string

const (
	EventListResponseEventSessionNextShellStartedTypeSessionNextShellStarted EventListResponseEventSessionNextShellStartedType = "session.next.shell.started"
)

func (r EventListResponseEventSessionNextShellStartedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextShellStartedTypeSessionNextShellStarted:
		return true
	}
	return false
}

type EventListResponseEventSessionNextShellEnded struct {
	Properties EventListResponseEventSessionNextShellEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextShellEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextShellEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextShellEndedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextShellEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextShellEndedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextShellEnded) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextShellEnded) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextShellEndedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID string `json:"callID,required"`
	Output string `json:"output,required"`
	JSON      eventListResponseEventSessionNextShellEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextShellEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID   apijson.Field
	Output   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextShellEndedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextShellEndedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextShellEndedType string

const (
	EventListResponseEventSessionNextShellEndedTypeSessionNextShellEnded EventListResponseEventSessionNextShellEndedType = "session.next.shell.ended"
)

func (r EventListResponseEventSessionNextShellEndedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextShellEndedTypeSessionNextShellEnded:
		return true
	}
	return false
}

type EventListResponseEventSessionNextStepStarted struct {
	Properties EventListResponseEventSessionNextStepStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextStepStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextStepStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextStepStartedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepStartedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextStepStarted) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextStepStarted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextStepStartedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Agent string `json:"agent,required"`
	Model EventListResponseEventSessionNextModelSwitchedModel `json:"model,required"`
	Snapshot string `json:"snapshot"`
	JSON      eventListResponseEventSessionNextStepStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextStepStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Agent   apijson.Field
	Model   apijson.Field
	Snapshot   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepStartedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepStartedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextStepStartedType string

const (
	EventListResponseEventSessionNextStepStartedTypeSessionNextStepStarted EventListResponseEventSessionNextStepStartedType = "session.next.step.started"
)

func (r EventListResponseEventSessionNextStepStartedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextStepStartedTypeSessionNextStepStarted:
		return true
	}
	return false
}

type EventListResponseEventSessionNextStepEnded struct {
	Properties EventListResponseEventSessionNextStepEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextStepEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextStepEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextStepEndedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepEndedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextStepEnded) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextStepEnded) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextStepEndedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Finish string `json:"finish,required"`
	Cost float64 `json:"cost,required"`
	Tokens EventListResponseEventSessionNextStepEndedTokens `json:"tokens,required"`
	Snapshot string `json:"snapshot"`
	JSON      eventListResponseEventSessionNextStepEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextStepEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Finish   apijson.Field
	Cost   apijson.Field
	Tokens   apijson.Field
	Snapshot   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepEndedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepEndedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextStepEndedType string

const (
	EventListResponseEventSessionNextStepEndedTypeSessionNextStepEnded EventListResponseEventSessionNextStepEndedType = "session.next.step.ended"
)

func (r EventListResponseEventSessionNextStepEndedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextStepEndedTypeSessionNextStepEnded:
		return true
	}
	return false
}

type EventListResponseEventSessionNextStepFailed struct {
	Properties EventListResponseEventSessionNextStepFailedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextStepFailedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextStepFailedJSON       `json:"-"`
}

type eventListResponseEventSessionNextStepFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextStepFailed) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextStepFailed) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextStepFailedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Error EventListResponseEventSessionNextStepFailedError `json:"error,required"`
	JSON      eventListResponseEventSessionNextStepFailedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextStepFailedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Error   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextStepFailedType string

const (
	EventListResponseEventSessionNextStepFailedTypeSessionNextStepFailed EventListResponseEventSessionNextStepFailedType = "session.next.step.failed"
)

func (r EventListResponseEventSessionNextStepFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextStepFailedTypeSessionNextStepFailed:
		return true
	}
	return false
}

type EventListResponseEventSessionNextTextStarted struct {
	Properties EventListResponseEventSessionNextTextStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextTextStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextTextStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextTextStartedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextTextStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextTextStartedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextTextStarted) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextTextStarted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextTextStartedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventSessionNextTextStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextTextStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextTextStartedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextTextStartedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextTextStartedType string

const (
	EventListResponseEventSessionNextTextStartedTypeSessionNextTextStarted EventListResponseEventSessionNextTextStartedType = "session.next.text.started"
)

func (r EventListResponseEventSessionNextTextStartedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextTextStartedTypeSessionNextTextStarted:
		return true
	}
	return false
}

type EventListResponseEventSessionNextTextDelta struct {
	Properties EventListResponseEventSessionNextTextDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextTextDeltaType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextTextDeltaJSON       `json:"-"`
}

type eventListResponseEventSessionNextTextDeltaJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextTextDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextTextDeltaJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextTextDelta) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextTextDelta) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextTextDeltaProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Delta string `json:"delta,required"`
	JSON      eventListResponseEventSessionNextTextDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextTextDeltaPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Delta   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextTextDeltaProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextTextDeltaPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextTextDeltaType string

const (
	EventListResponseEventSessionNextTextDeltaTypeSessionNextTextDelta EventListResponseEventSessionNextTextDeltaType = "session.next.text.delta"
)

func (r EventListResponseEventSessionNextTextDeltaType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextTextDeltaTypeSessionNextTextDelta:
		return true
	}
	return false
}

type EventListResponseEventSessionNextTextEnded struct {
	Properties EventListResponseEventSessionNextTextEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextTextEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextTextEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextTextEndedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextTextEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextTextEndedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextTextEnded) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextTextEnded) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextTextEndedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Text string `json:"text,required"`
	JSON      eventListResponseEventSessionNextTextEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextTextEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextTextEndedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextTextEndedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextTextEndedType string

const (
	EventListResponseEventSessionNextTextEndedTypeSessionNextTextEnded EventListResponseEventSessionNextTextEndedType = "session.next.text.ended"
)

func (r EventListResponseEventSessionNextTextEndedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextTextEndedTypeSessionNextTextEnded:
		return true
	}
	return false
}

type EventListResponseEventSessionNextReasoningStarted struct {
	Properties EventListResponseEventSessionNextReasoningStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextReasoningStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextReasoningStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextReasoningStartedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextReasoningStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextReasoningStartedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextReasoningStarted) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextReasoningStarted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextReasoningStartedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	ReasoningID string `json:"reasoningID,required"`
	JSON      eventListResponseEventSessionNextReasoningStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextReasoningStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	ReasoningID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextReasoningStartedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextReasoningStartedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextReasoningStartedType string

const (
	EventListResponseEventSessionNextReasoningStartedTypeSessionNextReasoningStarted EventListResponseEventSessionNextReasoningStartedType = "session.next.reasoning.started"
)

func (r EventListResponseEventSessionNextReasoningStartedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextReasoningStartedTypeSessionNextReasoningStarted:
		return true
	}
	return false
}

type EventListResponseEventSessionNextReasoningDelta struct {
	Properties EventListResponseEventSessionNextReasoningDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextReasoningDeltaType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextReasoningDeltaJSON       `json:"-"`
}

type eventListResponseEventSessionNextReasoningDeltaJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextReasoningDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextReasoningDeltaJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextReasoningDelta) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextReasoningDelta) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextReasoningDeltaProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	ReasoningID string `json:"reasoningID,required"`
	Delta string `json:"delta,required"`
	JSON      eventListResponseEventSessionNextReasoningDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextReasoningDeltaPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	ReasoningID   apijson.Field
	Delta   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextReasoningDeltaProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextReasoningDeltaPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextReasoningDeltaType string

const (
	EventListResponseEventSessionNextReasoningDeltaTypeSessionNextReasoningDelta EventListResponseEventSessionNextReasoningDeltaType = "session.next.reasoning.delta"
)

func (r EventListResponseEventSessionNextReasoningDeltaType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextReasoningDeltaTypeSessionNextReasoningDelta:
		return true
	}
	return false
}

type EventListResponseEventSessionNextReasoningEnded struct {
	Properties EventListResponseEventSessionNextReasoningEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextReasoningEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextReasoningEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextReasoningEndedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextReasoningEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextReasoningEndedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextReasoningEnded) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextReasoningEnded) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextReasoningEndedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	ReasoningID string `json:"reasoningID,required"`
	Text string `json:"text,required"`
	JSON      eventListResponseEventSessionNextReasoningEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextReasoningEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	ReasoningID   apijson.Field
	Text   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextReasoningEndedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextReasoningEndedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextReasoningEndedType string

const (
	EventListResponseEventSessionNextReasoningEndedTypeSessionNextReasoningEnded EventListResponseEventSessionNextReasoningEndedType = "session.next.reasoning.ended"
)

func (r EventListResponseEventSessionNextReasoningEndedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextReasoningEndedTypeSessionNextReasoningEnded:
		return true
	}
	return false
}

type EventListResponseEventSessionNextToolInputStarted struct {
	Properties EventListResponseEventSessionNextToolInputStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolInputStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolInputStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolInputStartedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolInputStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolInputStartedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextToolInputStarted) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextToolInputStarted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextToolInputStartedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID string `json:"callID,required"`
	Name string `json:"name,required"`
	JSON      eventListResponseEventSessionNextToolInputStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolInputStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID   apijson.Field
	Name   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolInputStartedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolInputStartedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextToolInputStartedType string

const (
	EventListResponseEventSessionNextToolInputStartedTypeSessionNextToolInputStarted EventListResponseEventSessionNextToolInputStartedType = "session.next.tool.input.started"
)

func (r EventListResponseEventSessionNextToolInputStartedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextToolInputStartedTypeSessionNextToolInputStarted:
		return true
	}
	return false
}

type EventListResponseEventSessionNextToolInputDelta struct {
	Properties EventListResponseEventSessionNextToolInputDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolInputDeltaType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolInputDeltaJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolInputDeltaJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolInputDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolInputDeltaJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextToolInputDelta) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextToolInputDelta) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextToolInputDeltaProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID string `json:"callID,required"`
	Delta string `json:"delta,required"`
	JSON      eventListResponseEventSessionNextToolInputDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolInputDeltaPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID   apijson.Field
	Delta   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolInputDeltaProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolInputDeltaPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextToolInputDeltaType string

const (
	EventListResponseEventSessionNextToolInputDeltaTypeSessionNextToolInputDelta EventListResponseEventSessionNextToolInputDeltaType = "session.next.tool.input.delta"
)

func (r EventListResponseEventSessionNextToolInputDeltaType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextToolInputDeltaTypeSessionNextToolInputDelta:
		return true
	}
	return false
}

type EventListResponseEventSessionNextToolInputEnded struct {
	Properties EventListResponseEventSessionNextToolInputEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolInputEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolInputEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolInputEndedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolInputEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolInputEndedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextToolInputEnded) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextToolInputEnded) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextToolInputEndedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID string `json:"callID,required"`
	Text string `json:"text,required"`
	JSON      eventListResponseEventSessionNextToolInputEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolInputEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID   apijson.Field
	Text   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolInputEndedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolInputEndedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextToolInputEndedType string

const (
	EventListResponseEventSessionNextToolInputEndedTypeSessionNextToolInputEnded EventListResponseEventSessionNextToolInputEndedType = "session.next.tool.input.ended"
)

func (r EventListResponseEventSessionNextToolInputEndedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextToolInputEndedTypeSessionNextToolInputEnded:
		return true
	}
	return false
}

type EventListResponseEventSessionNextToolCalled struct {
	Properties EventListResponseEventSessionNextToolCalledProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolCalledType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolCalledJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolCalledJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolCalled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolCalledJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextToolCalled) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextToolCalled) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextToolCalledProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID string `json:"callID,required"`
	Tool string `json:"tool,required"`
	Input interface{} `json:"input,required"`
	Provider EventListResponseEventSessionNextToolCalledProvider `json:"provider,required"`
	JSON      eventListResponseEventSessionNextToolCalledPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolCalledPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID   apijson.Field
	Tool   apijson.Field
	Input   apijson.Field
	Provider   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolCalledProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolCalledPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextToolCalledType string

const (
	EventListResponseEventSessionNextToolCalledTypeSessionNextToolCalled EventListResponseEventSessionNextToolCalledType = "session.next.tool.called"
)

func (r EventListResponseEventSessionNextToolCalledType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextToolCalledTypeSessionNextToolCalled:
		return true
	}
	return false
}

type EventListResponseEventSessionNextToolProgress struct {
	Properties EventListResponseEventSessionNextToolProgressProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolProgressType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolProgressJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolProgressJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolProgressJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextToolProgress) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextToolProgress) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextToolProgressProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID string `json:"callID,required"`
	Structured interface{} `json:"structured,required"`
	Content []interface{} `json:"content,required"`
	JSON      eventListResponseEventSessionNextToolProgressPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolProgressPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID   apijson.Field
	Structured   apijson.Field
	Content   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolProgressProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolProgressPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextToolProgressType string

const (
	EventListResponseEventSessionNextToolProgressTypeSessionNextToolProgress EventListResponseEventSessionNextToolProgressType = "session.next.tool.progress"
)

func (r EventListResponseEventSessionNextToolProgressType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextToolProgressTypeSessionNextToolProgress:
		return true
	}
	return false
}

type EventListResponseEventSessionNextToolSuccess struct {
	Properties EventListResponseEventSessionNextToolSuccessProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolSuccessType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolSuccessJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolSuccessJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolSuccessJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextToolSuccess) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextToolSuccess) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextToolSuccessProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID string `json:"callID,required"`
	Structured interface{} `json:"structured,required"`
	Content []interface{} `json:"content,required"`
	Provider EventListResponseEventSessionNextToolCalledProvider `json:"provider,required"`
	JSON      eventListResponseEventSessionNextToolSuccessPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolSuccessPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID   apijson.Field
	Structured   apijson.Field
	Content   apijson.Field
	Provider   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolSuccessProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolSuccessPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextToolSuccessType string

const (
	EventListResponseEventSessionNextToolSuccessTypeSessionNextToolSuccess EventListResponseEventSessionNextToolSuccessType = "session.next.tool.success"
)

func (r EventListResponseEventSessionNextToolSuccessType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextToolSuccessTypeSessionNextToolSuccess:
		return true
	}
	return false
}

type EventListResponseEventSessionNextToolFailed struct {
	Properties EventListResponseEventSessionNextToolFailedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolFailedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolFailedJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextToolFailed) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextToolFailed) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextToolFailedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID string `json:"callID,required"`
	Error EventListResponseEventSessionNextStepFailedError `json:"error,required"`
	Provider EventListResponseEventSessionNextToolCalledProvider `json:"provider,required"`
	JSON      eventListResponseEventSessionNextToolFailedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolFailedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID   apijson.Field
	Error   apijson.Field
	Provider   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextToolFailedType string

const (
	EventListResponseEventSessionNextToolFailedTypeSessionNextToolFailed EventListResponseEventSessionNextToolFailedType = "session.next.tool.failed"
)

func (r EventListResponseEventSessionNextToolFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextToolFailedTypeSessionNextToolFailed:
		return true
	}
	return false
}

type EventListResponseEventSessionNextRetried struct {
	Properties EventListResponseEventSessionNextRetriedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextRetriedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextRetriedJSON       `json:"-"`
}

type eventListResponseEventSessionNextRetriedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextRetried) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextRetriedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextRetried) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextRetried) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextRetriedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Attempt float64 `json:"attempt,required"`
	Error EventListResponseEventSessionNextRetriedError `json:"error,required"`
	JSON      eventListResponseEventSessionNextRetriedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextRetriedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Attempt   apijson.Field
	Error   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextRetriedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextRetriedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextRetriedType string

const (
	EventListResponseEventSessionNextRetriedTypeSessionNextRetried EventListResponseEventSessionNextRetriedType = "session.next.retried"
)

func (r EventListResponseEventSessionNextRetriedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextRetriedTypeSessionNextRetried:
		return true
	}
	return false
}

type EventListResponseEventSessionNextCompactionStarted struct {
	Properties EventListResponseEventSessionNextCompactionStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextCompactionStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextCompactionStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextCompactionStartedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextCompactionStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextCompactionStartedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextCompactionStarted) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextCompactionStarted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextCompactionStartedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Reason string `json:"reason,required"`
	JSON      eventListResponseEventSessionNextCompactionStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextCompactionStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Reason   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextCompactionStartedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextCompactionStartedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextCompactionStartedType string

const (
	EventListResponseEventSessionNextCompactionStartedTypeSessionNextCompactionStarted EventListResponseEventSessionNextCompactionStartedType = "session.next.compaction.started"
)

func (r EventListResponseEventSessionNextCompactionStartedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextCompactionStartedTypeSessionNextCompactionStarted:
		return true
	}
	return false
}

type EventListResponseEventSessionNextCompactionDelta struct {
	Properties EventListResponseEventSessionNextCompactionDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextCompactionDeltaType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextCompactionDeltaJSON       `json:"-"`
}

type eventListResponseEventSessionNextCompactionDeltaJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextCompactionDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextCompactionDeltaJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextCompactionDelta) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextCompactionDelta) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextCompactionDeltaProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Text string `json:"text,required"`
	JSON      eventListResponseEventSessionNextCompactionDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextCompactionDeltaPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextCompactionDeltaProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextCompactionDeltaPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextCompactionDeltaType string

const (
	EventListResponseEventSessionNextCompactionDeltaTypeSessionNextCompactionDelta EventListResponseEventSessionNextCompactionDeltaType = "session.next.compaction.delta"
)

func (r EventListResponseEventSessionNextCompactionDeltaType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextCompactionDeltaTypeSessionNextCompactionDelta:
		return true
	}
	return false
}

type EventListResponseEventSessionNextCompactionEnded struct {
	Properties EventListResponseEventSessionNextCompactionEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextCompactionEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextCompactionEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextCompactionEndedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextCompactionEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextCompactionEndedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextCompactionEnded) implementsEventListResponse() {}

func (r EventListResponseEventSessionNextCompactionEnded) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextCompactionEndedProperties struct {
	Timestamp float64 `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Text string `json:"text,required"`
	Include string `json:"include"`
	JSON      eventListResponseEventSessionNextCompactionEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextCompactionEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text   apijson.Field
	Include   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextCompactionEndedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextCompactionEndedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextCompactionEndedType string

const (
	EventListResponseEventSessionNextCompactionEndedTypeSessionNextCompactionEnded EventListResponseEventSessionNextCompactionEndedType = "session.next.compaction.ended"
)

func (r EventListResponseEventSessionNextCompactionEndedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextCompactionEndedTypeSessionNextCompactionEnded:
		return true
	}
	return false
}


// EventListResponseEventSessionNextModelSwitchedModel
type EventListResponseEventSessionNextModelSwitchedModel struct {
	ID         string `json:"id,required"`
	ProviderID string `json:"providerID,required"`
	Variant    string `json:"variant,required"`
	JSON       eventListResponseEventSessionNextModelSwitchedModelJSON `json:"-"`
}

type eventListResponseEventSessionNextModelSwitchedModelJSON struct {
	ID          apijson.Field
	ProviderID  apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextModelSwitchedModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextModelSwitchedModelJSON) RawJSON() string {
	return r.raw
}

// EventListResponseEventSessionNextToolCalledProvider
type EventListResponseEventSessionNextToolCalledProvider struct {
	Executed bool        `json:"executed,required"`
	Metadata interface{} `json:"metadata"`
	JSON     eventListResponseEventSessionNextToolCalledProviderJSON `json:"-"`
}

type eventListResponseEventSessionNextToolCalledProviderJSON struct {
	Executed    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolCalledProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolCalledProviderJSON) RawJSON() string {
	return r.raw
}

// EventListResponseEventSessionNextStepFailedError
type EventListResponseEventSessionNextStepFailedError struct {
	Type    string `json:"type,required"`
	Message string `json:"message,required"`
	JSON    eventListResponseEventSessionNextStepFailedErrorJSON `json:"-"`
}

type eventListResponseEventSessionNextStepFailedErrorJSON struct {
	Type        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepFailedError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepFailedErrorJSON) RawJSON() string {
	return r.raw
}

// EventListResponseEventSessionNextRetriedError
type EventListResponseEventSessionNextRetriedError struct {
	Message         string            `json:"message,required"`
	IsRetryable     bool              `json:"isRetryable,required"`
	StatusCode      float64           `json:"statusCode"`
	ResponseHeaders map[string]string `json:"responseHeaders"`
	ResponseBody    string            `json:"responseBody"`
	Metadata        map[string]string `json:"metadata"`
	JSON            eventListResponseEventSessionNextRetriedErrorJSON `json:"-"`
}

type eventListResponseEventSessionNextRetriedErrorJSON struct {
	Message         apijson.Field
	IsRetryable     apijson.Field
	StatusCode      apijson.Field
	ResponseHeaders apijson.Field
	ResponseBody    apijson.Field
	Metadata        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextRetriedError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextRetriedErrorJSON) RawJSON() string {
	return r.raw
}

// EventListResponseEventSessionNextPromptedPrompt
type EventListResponseEventSessionNextPromptedPrompt struct {
	Text       string                        `json:"text,required"`
	Files      []V2PromptFileAttachment      `json:"files"`
	Agents     []V2PromptAgentAttachment     `json:"agents"`
	References []V2PromptReferenceAttachment `json:"references"`
	JSON       eventListResponseEventSessionNextPromptedPromptJSON `json:"-"`
}

type eventListResponseEventSessionNextPromptedPromptJSON struct {
	Text        apijson.Field
	Files       apijson.Field
	Agents      apijson.Field
	References  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextPromptedPrompt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextPromptedPromptJSON) RawJSON() string {
	return r.raw
}

// EventListResponseEventSessionNextStepEndedTokens
type EventListResponseEventSessionNextStepEndedTokens struct {
	Input     float64 `json:"input,required"`
	Output    float64 `json:"output,required"`
	Reasoning float64 `json:"reasoning,required"`
	Cache     EventListResponseEventSessionNextStepEndedTokensCache `json:"cache,required"`
	JSON      eventListResponseEventSessionNextStepEndedTokensJSON  `json:"-"`
}

type eventListResponseEventSessionNextStepEndedTokensJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepEndedTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepEndedTokensJSON) RawJSON() string {
	return r.raw
}

// EventListResponseEventSessionNextStepEndedTokensCache
type EventListResponseEventSessionNextStepEndedTokensCache struct {
	Read  float64 `json:"read,required"`
	Write float64 `json:"write,required"`
	JSON  eventListResponseEventSessionNextStepEndedTokensCacheJSON `json:"-"`
}

type eventListResponseEventSessionNextStepEndedTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepEndedTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepEndedTokensCacheJSON) RawJSON() string {
	return r.raw
}

// EventListResponseEventSessionNextCompactionStartedReason
type EventListResponseEventSessionNextCompactionStartedReason string

const (
	EventListResponseEventSessionNextCompactionStartedReasonAuto   EventListResponseEventSessionNextCompactionStartedReason = "auto"
	EventListResponseEventSessionNextCompactionStartedReasonManual EventListResponseEventSessionNextCompactionStartedReason = "manual"
)

func (r EventListResponseEventSessionNextCompactionStartedReason) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextCompactionStartedReasonAuto, EventListResponseEventSessionNextCompactionStartedReasonManual:
		return true
	}
	return false
}

type EventListResponseEventCatalogModelUpdated struct {
	Properties EventListResponseEventCatalogModelUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventCatalogModelUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventCatalogModelUpdatedJSON       `json:"-"`
}

type eventListResponseEventCatalogModelUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCatalogModelUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCatalogModelUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventCatalogModelUpdated) implementsEventListResponse() {}

func (r EventListResponseEventCatalogModelUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventCatalogModelUpdatedProperties struct {
	Model V2ModelInfo                                            `json:"model,required"`
	JSON  eventListResponseEventCatalogModelUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventCatalogModelUpdatedPropertiesJSON struct {
	Model       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCatalogModelUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCatalogModelUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventCatalogModelUpdatedType string

const (
	EventListResponseEventCatalogModelUpdatedTypeCatalogModelUpdated EventListResponseEventCatalogModelUpdatedType = "catalog.model.updated"
)

func (r EventListResponseEventCatalogModelUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventCatalogModelUpdatedTypeCatalogModelUpdated:
		return true
	}
	return false
}
