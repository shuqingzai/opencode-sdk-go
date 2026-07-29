// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"github.com/sst/opencode-sdk-go/internal/apijson"
)

type EventListResponseEventMessageUpdated struct {
	ID         string                                         `json:"id,required"`
	Properties EventListResponseEventMessageUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventMessageUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventMessageUpdatedJSON       `json:"-"`
}

type eventListResponseEventMessageUpdatedJSON struct {
	ID          apijson.Field
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
	SessionID string                                             `json:"sessionID,required"`
	Info      Message                                            `json:"info,required"`
	JSON      eventListResponseEventMessageUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventMessageUpdatedPropertiesJSON struct {
	SessionID   apijson.Field
	Info        apijson.Field
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
	ID         string                                         `json:"id,required"`
	Properties EventListResponseEventMessageRemovedProperties `json:"properties,required"`
	Type       EventListResponseEventMessageRemovedType       `json:"type,required"`
	JSON       eventListResponseEventMessageRemovedJSON       `json:"-"`
}

type eventListResponseEventMessageRemovedJSON struct {
	ID          apijson.Field
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
	MessageID string                                             `json:"messageID,required"`
	SessionID string                                             `json:"sessionID,required"`
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
	ID         string                                             `json:"id,required"`
	Properties EventListResponseEventMessagePartUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventMessagePartUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventMessagePartUpdatedJSON       `json:"-"`
}

type eventListResponseEventMessagePartUpdatedJSON struct {
	ID          apijson.Field
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
	Part      Part                                                   `json:"part,required"`
	SessionID string                                                 `json:"sessionID,required"`
	Time      int64                                                  `json:"time,required"`
	JSON      eventListResponseEventMessagePartUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventMessagePartUpdatedPropertiesJSON struct {
	Part        apijson.Field
	SessionID   apijson.Field
	Time        apijson.Field
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
	ID         string                                             `json:"id,required"`
	Properties EventListResponseEventMessagePartRemovedProperties `json:"properties,required"`
	Type       EventListResponseEventMessagePartRemovedType       `json:"type,required"`
	JSON       eventListResponseEventMessagePartRemovedJSON       `json:"-"`
}

type eventListResponseEventMessagePartRemovedJSON struct {
	ID          apijson.Field
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
	MessageID string                                                 `json:"messageID,required"`
	PartID    string                                                 `json:"partID,required"`
	SessionID string                                                 `json:"sessionID,required"`
	JSON      eventListResponseEventMessagePartRemovedPropertiesJSON `json:"-"`
}

type eventListResponseEventMessagePartRemovedPropertiesJSON struct {
	MessageID   apijson.Field
	PartID      apijson.Field
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
	ID         string                                         `json:"id,required"`
	Properties EventListResponseEventSessionCreatedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionCreatedType       `json:"type,required"`
	JSON       eventListResponseEventSessionCreatedJSON       `json:"-"`
}

type eventListResponseEventSessionCreatedJSON struct {
	ID          apijson.Field
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
	SessionID string                                             `json:"sessionID,required"`
	Info      Session                                            `json:"info,required"`
	JSON      eventListResponseEventSessionCreatedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionCreatedPropertiesJSON struct {
	SessionID   apijson.Field
	Info        apijson.Field
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
	ID         string                                         `json:"id,required"`
	Properties EventListResponseEventSessionUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventSessionUpdatedJSON       `json:"-"`
}

type eventListResponseEventSessionUpdatedJSON struct {
	ID          apijson.Field
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
	SessionID string                                             `json:"sessionID,required"`
	Info      Session                                            `json:"info,required"`
	JSON      eventListResponseEventSessionUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionUpdatedPropertiesJSON struct {
	SessionID   apijson.Field
	Info        apijson.Field
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
	ID         string                                         `json:"id,required"`
	Properties EventListResponseEventSessionDeletedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionDeletedType       `json:"type,required"`
	JSON       eventListResponseEventSessionDeletedJSON       `json:"-"`
}

type eventListResponseEventSessionDeletedJSON struct {
	ID          apijson.Field
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
	SessionID string                                             `json:"sessionID,required"`
	Info      Session                                            `json:"info,required"`
	JSON      eventListResponseEventSessionDeletedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionDeletedPropertiesJSON struct {
	SessionID   apijson.Field
	Info        apijson.Field
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
	ID         string                                                   `json:"id,required"`
	Properties EventListResponseEventSessionNextAgentSwitchedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextAgentSwitchedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextAgentSwitchedJSON       `json:"-"`
}

type eventListResponseEventSessionNextAgentSwitchedJSON struct {
	ID          apijson.Field
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
	Timestamp int64                                                        `json:"timestamp,required"`
	MessageID string                                                       `json:"messageID,required"`
	SessionID string                                                       `json:"sessionID,required"`
	Agent     string                                                       `json:"agent,required"`
	JSON      eventListResponseEventSessionNextAgentSwitchedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextAgentSwitchedPropertiesJSON struct {
	Timestamp   apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Agent       apijson.Field
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
	ID         string                                                   `json:"id,required"`
	Properties EventListResponseEventSessionNextModelSwitchedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextModelSwitchedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextModelSwitchedJSON       `json:"-"`
}

type eventListResponseEventSessionNextModelSwitchedJSON struct {
	ID          apijson.Field
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
	Timestamp int64                                                        `json:"timestamp,required"`
	MessageID string                                                       `json:"messageID,required"`
	SessionID string                                                       `json:"sessionID,required"`
	Model     ModelRef                                                     `json:"model,required"`
	JSON      eventListResponseEventSessionNextModelSwitchedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextModelSwitchedPropertiesJSON struct {
	Timestamp   apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Model       apijson.Field
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
	ID         string                                              `json:"id,required"`
	Properties EventListResponseEventSessionNextPromptedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextPromptedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextPromptedJSON       `json:"-"`
}

type eventListResponseEventSessionNextPromptedJSON struct {
	ID          apijson.Field
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

type EventListResponseEventSessionNextPromptedDelivery string

const (
	EventListResponseEventSessionNextPromptedDeliverySteer EventListResponseEventSessionNextPromptedDelivery = "steer"
	EventListResponseEventSessionNextPromptedDeliveryQueue EventListResponseEventSessionNextPromptedDelivery = "queue"
)

func (r EventListResponseEventSessionNextPromptedDelivery) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextPromptedDeliverySteer:
		return true
	case EventListResponseEventSessionNextPromptedDeliveryQueue:
		return true
	}
	return false
}

type EventListResponseEventSessionNextPromptedProperties struct {
	Timestamp int64                                                   `json:"timestamp,required"`
	MessageID string                                                  `json:"messageID,required"`
	SessionID string                                                  `json:"sessionID,required"`
	Prompt    V2SessionInputPrompt                                    `json:"prompt,required"`
	Delivery  EventListResponseEventSessionNextPromptedDelivery       `json:"delivery,required"`
	JSON      eventListResponseEventSessionNextPromptedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextPromptedPropertiesJSON struct {
	Timestamp   apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Prompt      apijson.Field
	Delivery    apijson.Field
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
	ID         string                                               `json:"id,required"`
	Properties EventListResponseEventSessionNextSyntheticProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextSyntheticType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextSyntheticJSON       `json:"-"`
}

type eventListResponseEventSessionNextSyntheticJSON struct {
	ID          apijson.Field
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
	Timestamp int64                                                    `json:"timestamp,required"`
	MessageID string                                                   `json:"messageID,required"`
	SessionID string                                                   `json:"sessionID,required"`
	Text      string                                                   `json:"text,required"`
	JSON      eventListResponseEventSessionNextSyntheticPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextSyntheticPropertiesJSON struct {
	Timestamp   apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
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
	ID         string                                                  `json:"id,required"`
	Properties EventListResponseEventSessionNextShellStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextShellStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextShellStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextShellStartedJSON struct {
	ID          apijson.Field
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
	Timestamp int64                                                       `json:"timestamp,required"`
	MessageID string                                                      `json:"messageID,required"`
	SessionID string                                                      `json:"sessionID,required"`
	CallID    string                                                      `json:"callID,required"`
	Command   string                                                      `json:"command,required"`
	JSON      eventListResponseEventSessionNextShellStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextShellStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Command     apijson.Field
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
	ID         string                                                `json:"id,required"`
	Properties EventListResponseEventSessionNextShellEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextShellEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextShellEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextShellEndedJSON struct {
	ID          apijson.Field
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
	Timestamp int64                                                     `json:"timestamp,required"`
	SessionID string                                                    `json:"sessionID,required"`
	CallID    string                                                    `json:"callID,required"`
	Output    string                                                    `json:"output,required"`
	JSON      eventListResponseEventSessionNextShellEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextShellEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Output      apijson.Field
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
	ID         string                                                 `json:"id,required"`
	Properties EventListResponseEventSessionNextStepStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextStepStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextStepStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextStepStartedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                      `json:"timestamp,required"`
	AssistantMessageID string                                                     `json:"assistantMessageID,required"`
	SessionID          string                                                     `json:"sessionID,required"`
	Agent              string                                                     `json:"agent,required"`
	Model              ModelRef                                                   `json:"model,required"`
	Snapshot           string                                                     `json:"snapshot"`
	JSON               eventListResponseEventSessionNextStepStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextStepStartedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	Agent              apijson.Field
	Model              apijson.Field
	Snapshot           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                               `json:"id,required"`
	Properties EventListResponseEventSessionNextStepEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextStepEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextStepEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextStepEndedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                    `json:"timestamp,required"`
	AssistantMessageID string                                                   `json:"assistantMessageID,required"`
	SessionID          string                                                   `json:"sessionID,required"`
	Finish             string                                                   `json:"finish,required"`
	Cost               float64                                                  `json:"cost,required"`
	Tokens             EventListResponseEventSessionNextStepEndedTokens         `json:"tokens,required"`
	Snapshot           string                                                   `json:"snapshot"`
	Files              []string                                                 `json:"files"`
	JSON               eventListResponseEventSessionNextStepEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextStepEndedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	Finish             apijson.Field
	Cost               apijson.Field
	Tokens             apijson.Field
	Snapshot           apijson.Field
	Files              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                `json:"id,required"`
	Properties EventListResponseEventSessionNextStepFailedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextStepFailedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextStepFailedJSON       `json:"-"`
}

type eventListResponseEventSessionNextStepFailedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                     `json:"timestamp,required"`
	AssistantMessageID string                                                    `json:"assistantMessageID,required"`
	SessionID          string                                                    `json:"sessionID,required"`
	Error              SessionErrorUnknown                                       `json:"error,required"`
	JSON               eventListResponseEventSessionNextStepFailedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextStepFailedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	Error              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                 `json:"id,required"`
	Properties EventListResponseEventSessionNextTextStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextTextStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextTextStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextTextStartedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                      `json:"timestamp,required"`
	AssistantMessageID string                                                     `json:"assistantMessageID,required"`
	SessionID          string                                                     `json:"sessionID,required"`
	TextID             string                                                     `json:"textID,required"`
	JSON               eventListResponseEventSessionNextTextStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextTextStartedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	TextID             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                               `json:"id,required"`
	Properties EventListResponseEventSessionNextTextDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextTextDeltaType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextTextDeltaJSON       `json:"-"`
}

type eventListResponseEventSessionNextTextDeltaJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                    `json:"timestamp,required"`
	AssistantMessageID string                                                   `json:"assistantMessageID,required"`
	SessionID          string                                                   `json:"sessionID,required"`
	TextID             string                                                   `json:"textID,required"`
	Delta              string                                                   `json:"delta,required"`
	JSON               eventListResponseEventSessionNextTextDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextTextDeltaPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	TextID             apijson.Field
	Delta              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                               `json:"id,required"`
	Properties EventListResponseEventSessionNextTextEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextTextEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextTextEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextTextEndedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                    `json:"timestamp,required"`
	AssistantMessageID string                                                   `json:"assistantMessageID,required"`
	SessionID          string                                                   `json:"sessionID,required"`
	TextID             string                                                   `json:"textID,required"`
	Text               string                                                   `json:"text,required"`
	JSON               eventListResponseEventSessionNextTextEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextTextEndedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	TextID             apijson.Field
	Text               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                      `json:"id,required"`
	Properties EventListResponseEventSessionNextReasoningStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextReasoningStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextReasoningStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextReasoningStartedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64  `json:"timestamp,required"`
	AssistantMessageID string `json:"assistantMessageID,required"`
	SessionID          string `json:"sessionID,required"`
	ReasoningID        string `json:"reasoningID,required"`
	// This field can have the runtime type of [map[string]any].
	ProviderMetadata any                                                             `json:"providerMetadata"`
	JSON             eventListResponseEventSessionNextReasoningStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextReasoningStartedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	ReasoningID        apijson.Field
	ProviderMetadata   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                    `json:"id,required"`
	Properties EventListResponseEventSessionNextReasoningDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextReasoningDeltaType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextReasoningDeltaJSON       `json:"-"`
}

type eventListResponseEventSessionNextReasoningDeltaJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                         `json:"timestamp,required"`
	AssistantMessageID string                                                        `json:"assistantMessageID,required"`
	SessionID          string                                                        `json:"sessionID,required"`
	ReasoningID        string                                                        `json:"reasoningID,required"`
	Delta              string                                                        `json:"delta,required"`
	JSON               eventListResponseEventSessionNextReasoningDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextReasoningDeltaPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	ReasoningID        apijson.Field
	Delta              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                    `json:"id,required"`
	Properties EventListResponseEventSessionNextReasoningEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextReasoningEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextReasoningEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextReasoningEndedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64  `json:"timestamp,required"`
	AssistantMessageID string `json:"assistantMessageID,required"`
	SessionID          string `json:"sessionID,required"`
	ReasoningID        string `json:"reasoningID,required"`
	Text               string `json:"text,required"`
	// This field can have the runtime type of [map[string]any].
	ProviderMetadata any                                                           `json:"providerMetadata"`
	JSON             eventListResponseEventSessionNextReasoningEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextReasoningEndedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	ReasoningID        apijson.Field
	Text               apijson.Field
	ProviderMetadata   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                      `json:"id,required"`
	Properties EventListResponseEventSessionNextToolInputStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolInputStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolInputStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolInputStartedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                           `json:"timestamp,required"`
	AssistantMessageID string                                                          `json:"assistantMessageID,required"`
	SessionID          string                                                          `json:"sessionID,required"`
	CallID             string                                                          `json:"callID,required"`
	Name               string                                                          `json:"name,required"`
	JSON               eventListResponseEventSessionNextToolInputStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolInputStartedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	CallID             apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                    `json:"id,required"`
	Properties EventListResponseEventSessionNextToolInputDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolInputDeltaType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolInputDeltaJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolInputDeltaJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                         `json:"timestamp,required"`
	AssistantMessageID string                                                        `json:"assistantMessageID,required"`
	SessionID          string                                                        `json:"sessionID,required"`
	CallID             string                                                        `json:"callID,required"`
	Delta              string                                                        `json:"delta,required"`
	JSON               eventListResponseEventSessionNextToolInputDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolInputDeltaPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	CallID             apijson.Field
	Delta              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                    `json:"id,required"`
	Properties EventListResponseEventSessionNextToolInputEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolInputEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolInputEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolInputEndedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                                         `json:"timestamp,required"`
	AssistantMessageID string                                                        `json:"assistantMessageID,required"`
	SessionID          string                                                        `json:"sessionID,required"`
	CallID             string                                                        `json:"callID,required"`
	Text               string                                                        `json:"text,required"`
	JSON               eventListResponseEventSessionNextToolInputEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolInputEndedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	CallID             apijson.Field
	Text               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                `json:"id,required"`
	Properties EventListResponseEventSessionNextToolCalledProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolCalledType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolCalledJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolCalledJSON struct {
	ID          apijson.Field
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
	Timestamp          int64  `json:"timestamp,required"`
	AssistantMessageID string `json:"assistantMessageID,required"`
	SessionID          string `json:"sessionID,required"`
	CallID             string `json:"callID,required"`
	Tool               string `json:"tool,required"`
	// This field can have the runtime type of [map[string]any].
	Input    any                                                       `json:"input,required"`
	Provider EventListResponseEventSessionNextToolCalledProvider       `json:"provider,required"`
	JSON     eventListResponseEventSessionNextToolCalledPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolCalledPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	CallID             apijson.Field
	Tool               apijson.Field
	Input              apijson.Field
	Provider           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                  `json:"id,required"`
	Properties EventListResponseEventSessionNextToolProgressProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolProgressType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolProgressJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolProgressJSON struct {
	ID          apijson.Field
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
	Timestamp          int64  `json:"timestamp,required"`
	AssistantMessageID string `json:"assistantMessageID,required"`
	SessionID          string `json:"sessionID,required"`
	CallID             string `json:"callID,required"`
	// This field can have the runtime type of [map[string]any].
	Structured any `json:"structured,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content []any                                                       `json:"content,required"`
	JSON    eventListResponseEventSessionNextToolProgressPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolProgressPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	CallID             apijson.Field
	Structured         apijson.Field
	Content            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                 `json:"id,required"`
	Properties EventListResponseEventSessionNextToolSuccessProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolSuccessType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolSuccessJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolSuccessJSON struct {
	ID          apijson.Field
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
	Timestamp          int64  `json:"timestamp,required"`
	AssistantMessageID string `json:"assistantMessageID,required"`
	SessionID          string `json:"sessionID,required"`
	CallID             string `json:"callID,required"`
	// This field can have the runtime type of [map[string]any].
	Structured any `json:"structured,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content     []any                                               `json:"content,required"`
	OutputPaths []string                                            `json:"outputPaths"`
	Provider    EventListResponseEventSessionNextToolCalledProvider `json:"provider,required"`
	// Arbitrary JSON value holding the tool's result. Per OpenAPI
	// `EventSessionNextToolSuccess.properties.result` is an unconstrained schema
	// (`{}`), so no fixed set of runtime types applies.
	Result any                                                        `json:"result"`
	JSON   eventListResponseEventSessionNextToolSuccessPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolSuccessPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	CallID             apijson.Field
	Structured         apijson.Field
	Content            apijson.Field
	OutputPaths        apijson.Field
	Provider           apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                                `json:"id,required"`
	Properties EventListResponseEventSessionNextToolFailedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextToolFailedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextToolFailedJSON       `json:"-"`
}

type eventListResponseEventSessionNextToolFailedJSON struct {
	ID          apijson.Field
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
	Timestamp          int64                                               `json:"timestamp,required"`
	AssistantMessageID string                                              `json:"assistantMessageID,required"`
	SessionID          string                                              `json:"sessionID,required"`
	CallID             string                                              `json:"callID,required"`
	Error              SessionErrorUnknown                                 `json:"error,required"`
	Provider           EventListResponseEventSessionNextToolCalledProvider `json:"provider,required"`
	// Arbitrary JSON value holding the tool's result. Per OpenAPI
	// `EventSessionNextToolFailed.properties.result` is an unconstrained schema
	// (`{}`), so no fixed set of runtime types applies.
	Result any                                                       `json:"result"`
	JSON   eventListResponseEventSessionNextToolFailedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolFailedPropertiesJSON struct {
	Timestamp          apijson.Field
	AssistantMessageID apijson.Field
	SessionID          apijson.Field
	CallID             apijson.Field
	Error              apijson.Field
	Provider           apijson.Field
	Result             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	ID         string                                             `json:"id,required"`
	Properties EventListResponseEventSessionNextRetriedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextRetriedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextRetriedJSON       `json:"-"`
}

type eventListResponseEventSessionNextRetriedJSON struct {
	ID          apijson.Field
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
	Timestamp int64                                                  `json:"timestamp,required"`
	SessionID string                                                 `json:"sessionID,required"`
	Attempt   int64                                                  `json:"attempt,required"`
	Error     EventListResponseEventSessionNextRetriedError          `json:"error,required"`
	JSON      eventListResponseEventSessionNextRetriedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextRetriedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Attempt     apijson.Field
	Error       apijson.Field
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
	ID         string                                                       `json:"id,required"`
	Properties EventListResponseEventSessionNextCompactionStartedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextCompactionStartedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextCompactionStartedJSON       `json:"-"`
}

type eventListResponseEventSessionNextCompactionStartedJSON struct {
	ID          apijson.Field
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
	Timestamp int64                                                            `json:"timestamp,required"`
	MessageID string                                                           `json:"messageID,required"`
	SessionID string                                                           `json:"sessionID,required"`
	Reason    EventListResponseEventSessionNextCompactionStartedReason         `json:"reason,required"`
	JSON      eventListResponseEventSessionNextCompactionStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextCompactionStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Reason      apijson.Field
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
	ID         string                                                     `json:"id,required"`
	Properties EventListResponseEventSessionNextCompactionDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextCompactionDeltaType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextCompactionDeltaJSON       `json:"-"`
}

type eventListResponseEventSessionNextCompactionDeltaJSON struct {
	ID          apijson.Field
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
	Timestamp int64                                                          `json:"timestamp,required"`
	MessageID string                                                         `json:"messageID,required"`
	SessionID string                                                         `json:"sessionID,required"`
	Text      string                                                         `json:"text,required"`
	JSON      eventListResponseEventSessionNextCompactionDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextCompactionDeltaPropertiesJSON struct {
	Timestamp   apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
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
	ID         string                                                     `json:"id,required"`
	Properties EventListResponseEventSessionNextCompactionEndedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextCompactionEndedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextCompactionEndedJSON       `json:"-"`
}

type eventListResponseEventSessionNextCompactionEndedJSON struct {
	ID          apijson.Field
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

type EventListResponseEventSessionNextCompactionEndedReason string

const (
	EventListResponseEventSessionNextCompactionEndedReasonAuto   EventListResponseEventSessionNextCompactionEndedReason = "auto"
	EventListResponseEventSessionNextCompactionEndedReasonManual EventListResponseEventSessionNextCompactionEndedReason = "manual"
)

func (r EventListResponseEventSessionNextCompactionEndedReason) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextCompactionEndedReasonAuto, EventListResponseEventSessionNextCompactionEndedReasonManual:
		return true
	}
	return false
}

type EventListResponseEventSessionNextCompactionEndedProperties struct {
	Timestamp int64                                                          `json:"timestamp,required"`
	MessageID string                                                         `json:"messageID,required"`
	SessionID string                                                         `json:"sessionID,required"`
	Reason    EventListResponseEventSessionNextCompactionEndedReason         `json:"reason,required"`
	Text      string                                                         `json:"text,required"`
	Recent    string                                                         `json:"recent,required"`
	JSON      eventListResponseEventSessionNextCompactionEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextCompactionEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	MessageID   apijson.Field
	SessionID   apijson.Field
	Reason      apijson.Field
	Text        apijson.Field
	Recent      apijson.Field
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

// EventListResponseEventSessionNextToolCalledProvider
type EventListResponseEventSessionNextToolCalledProvider struct {
	Executed bool `json:"executed,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                                                     `json:"metadata"`
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

// EventListResponseEventSessionNextRetriedError
type EventListResponseEventSessionNextRetriedError struct {
	Message         string                                            `json:"message,required"`
	IsRetryable     bool                                              `json:"isRetryable,required"`
	StatusCode      int64                                             `json:"statusCode"`
	ResponseHeaders map[string]string                                 `json:"responseHeaders"`
	ResponseBody    string                                            `json:"responseBody"`
	Metadata        map[string]string                                 `json:"metadata"`
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

// EventListResponseEventSessionNextStepEndedTokens
type EventListResponseEventSessionNextStepEndedTokens struct {
	Input     int64                                                 `json:"input,required"`
	Output    int64                                                 `json:"output,required"`
	Reasoning int64                                                 `json:"reasoning,required"`
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
	Read  int64                                                     `json:"read,required"`
	Write int64                                                     `json:"write,required"`
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

type EventListResponseEventPluginAdded struct {
	ID         string                                      `json:"id,required"`
	Properties EventListResponseEventPluginAddedProperties `json:"properties,required"`
	Type       EventListResponseEventPluginAddedType       `json:"type,required"`
	JSON       eventListResponseEventPluginAddedJSON       `json:"-"`
}

type eventListResponseEventPluginAddedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPluginAdded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPluginAddedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPluginAdded) implementsEventListResponse() {}

func (r EventListResponseEventPluginAdded) implementsGlobalEventPayload() {}

type EventListResponseEventPluginAddedProperties struct {
	ID   string                                          `json:"id,required"`
	JSON eventListResponseEventPluginAddedPropertiesJSON `json:"-"`
}

type eventListResponseEventPluginAddedPropertiesJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPluginAddedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPluginAddedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPluginAddedType string

const (
	EventListResponseEventPluginAddedTypePluginAdded EventListResponseEventPluginAddedType = "plugin.added"
)

func (r EventListResponseEventPluginAddedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPluginAddedTypePluginAdded:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventIntegrationUpdated
// =============================================================================

type EventListResponseEventIntegrationUpdated struct {
	ID         string                                             `json:"id,required"`
	Properties EventListResponseEventIntegrationUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventIntegrationUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventIntegrationUpdatedJSON       `json:"-"`
}

type eventListResponseEventIntegrationUpdatedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventIntegrationUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventIntegrationUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventIntegrationUpdated) implementsEventListResponse()  {}
func (r EventListResponseEventIntegrationUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventIntegrationUpdatedProperties struct {
	JSON eventListResponseEventIntegrationUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventIntegrationUpdatedPropertiesJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventIntegrationUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventIntegrationUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventIntegrationUpdatedType string

const (
	EventListResponseEventIntegrationUpdatedTypeIntegrationUpdated EventListResponseEventIntegrationUpdatedType = "integration.updated"
)

func (r EventListResponseEventIntegrationUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventIntegrationUpdatedTypeIntegrationUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventIntegrationConnectionUpdated
// =============================================================================

type EventListResponseEventIntegrationConnectionUpdated struct {
	ID         string                                                       `json:"id,required"`
	Properties EventListResponseEventIntegrationConnectionUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventIntegrationConnectionUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventIntegrationConnectionUpdatedJSON       `json:"-"`
}

type eventListResponseEventIntegrationConnectionUpdatedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventIntegrationConnectionUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventIntegrationConnectionUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventIntegrationConnectionUpdated) implementsEventListResponse()  {}
func (r EventListResponseEventIntegrationConnectionUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventIntegrationConnectionUpdatedProperties struct {
	IntegrationID string                                                           `json:"integrationID,required"`
	JSON          eventListResponseEventIntegrationConnectionUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventIntegrationConnectionUpdatedPropertiesJSON struct {
	IntegrationID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EventListResponseEventIntegrationConnectionUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventIntegrationConnectionUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventIntegrationConnectionUpdatedType string

const (
	EventListResponseEventIntegrationConnectionUpdatedTypeIntegrationConnectionUpdated EventListResponseEventIntegrationConnectionUpdatedType = "integration.connection.updated"
)

func (r EventListResponseEventIntegrationConnectionUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventIntegrationConnectionUpdatedTypeIntegrationConnectionUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventCatalogUpdated
// =============================================================================

type EventListResponseEventCatalogUpdated struct {
	ID         string                                         `json:"id,required"`
	Properties EventListResponseEventCatalogUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventCatalogUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventCatalogUpdatedJSON       `json:"-"`
}

type eventListResponseEventCatalogUpdatedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCatalogUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCatalogUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventCatalogUpdated) implementsEventListResponse()  {}
func (r EventListResponseEventCatalogUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventCatalogUpdatedProperties struct {
	JSON eventListResponseEventCatalogUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventCatalogUpdatedPropertiesJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCatalogUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCatalogUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventCatalogUpdatedType string

const (
	EventListResponseEventCatalogUpdatedTypeCatalogUpdated EventListResponseEventCatalogUpdatedType = "catalog.updated"
)

func (r EventListResponseEventCatalogUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventCatalogUpdatedTypeCatalogUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPermissionV2Asked
// =============================================================================

type EventListResponseEventPermissionV2Asked struct {
	ID         string                                            `json:"id,required"`
	Properties EventListResponseEventPermissionV2AskedProperties `json:"properties,required"`
	Type       EventListResponseEventPermissionV2AskedType       `json:"type,required"`
	JSON       eventListResponseEventPermissionV2AskedJSON       `json:"-"`
}

type eventListResponseEventPermissionV2AskedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionV2Asked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionV2AskedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPermissionV2Asked) implementsEventListResponse()  {}
func (r EventListResponseEventPermissionV2Asked) implementsGlobalEventPayload() {}

type EventListResponseEventPermissionV2AskedProperties struct {
	ID        string                                                `json:"id,required"`
	SessionID string                                                `json:"sessionID,required"`
	Action    string                                                `json:"action,required"`
	Resources []string                                              `json:"resources,required"`
	Save      []string                                              `json:"save"`
	Metadata  map[string]any                                        `json:"metadata"`
	Source    PermissionV2Source                                    `json:"source"`
	JSON      eventListResponseEventPermissionV2AskedPropertiesJSON `json:"-"`
}

type eventListResponseEventPermissionV2AskedPropertiesJSON struct {
	ID          apijson.Field
	SessionID   apijson.Field
	Action      apijson.Field
	Resources   apijson.Field
	Save        apijson.Field
	Metadata    apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionV2AskedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionV2AskedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPermissionV2AskedType string

const (
	EventListResponseEventPermissionV2AskedTypePermissionV2Asked EventListResponseEventPermissionV2AskedType = "permission.v2.asked"
)

func (r EventListResponseEventPermissionV2AskedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPermissionV2AskedTypePermissionV2Asked:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPermissionV2Replied
// =============================================================================

type EventListResponseEventPermissionV2Replied struct {
	ID         string                                              `json:"id,required"`
	Properties EventListResponseEventPermissionV2RepliedProperties `json:"properties,required"`
	Type       EventListResponseEventPermissionV2RepliedType       `json:"type,required"`
	JSON       eventListResponseEventPermissionV2RepliedJSON       `json:"-"`
}

type eventListResponseEventPermissionV2RepliedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionV2Replied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionV2RepliedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPermissionV2Replied) implementsEventListResponse()  {}
func (r EventListResponseEventPermissionV2Replied) implementsGlobalEventPayload() {}

type EventListResponseEventPermissionV2RepliedProperties struct {
	SessionID string                                                  `json:"sessionID,required"`
	RequestID string                                                  `json:"requestID,required"`
	Reply     PermissionV2Reply                                       `json:"reply,required"`
	JSON      eventListResponseEventPermissionV2RepliedPropertiesJSON `json:"-"`
}

type eventListResponseEventPermissionV2RepliedPropertiesJSON struct {
	SessionID   apijson.Field
	RequestID   apijson.Field
	Reply       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionV2RepliedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionV2RepliedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPermissionV2RepliedType string

const (
	EventListResponseEventPermissionV2RepliedTypePermissionV2Replied EventListResponseEventPermissionV2RepliedType = "permission.v2.replied"
)

func (r EventListResponseEventPermissionV2RepliedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPermissionV2RepliedTypePermissionV2Replied:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventReferenceUpdated
// =============================================================================

type EventListResponseEventReferenceUpdated struct {
	ID         string                                           `json:"id,required"`
	Properties EventListResponseEventReferenceUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventReferenceUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventReferenceUpdatedJSON       `json:"-"`
}

type eventListResponseEventReferenceUpdatedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventReferenceUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventReferenceUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventReferenceUpdated) implementsEventListResponse()  {}
func (r EventListResponseEventReferenceUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventReferenceUpdatedProperties struct {
	JSON eventListResponseEventReferenceUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventReferenceUpdatedPropertiesJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventReferenceUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventReferenceUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventReferenceUpdatedType string

const (
	EventListResponseEventReferenceUpdatedTypeReferenceUpdated EventListResponseEventReferenceUpdatedType = "reference.updated"
)

func (r EventListResponseEventReferenceUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventReferenceUpdatedTypeReferenceUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventQuestionV2Asked
// =============================================================================

type EventListResponseEventQuestionV2Asked struct {
	ID         string                                          `json:"id,required"`
	Properties EventListResponseEventQuestionV2AskedProperties `json:"properties,required"`
	Type       EventListResponseEventQuestionV2AskedType       `json:"type,required"`
	JSON       eventListResponseEventQuestionV2AskedJSON       `json:"-"`
}

type eventListResponseEventQuestionV2AskedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionV2Asked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionV2AskedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionV2Asked) implementsEventListResponse()  {}
func (r EventListResponseEventQuestionV2Asked) implementsGlobalEventPayload() {}

type EventListResponseEventQuestionV2AskedProperties struct {
	ID        string                                              `json:"id,required"`
	SessionID string                                              `json:"sessionID,required"`
	Questions []QuestionV2Info                                    `json:"questions,required"`
	Tool      QuestionV2Tool                                      `json:"tool"`
	JSON      eventListResponseEventQuestionV2AskedPropertiesJSON `json:"-"`
}

type eventListResponseEventQuestionV2AskedPropertiesJSON struct {
	ID          apijson.Field
	SessionID   apijson.Field
	Questions   apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionV2AskedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionV2AskedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionV2AskedType string

const (
	EventListResponseEventQuestionV2AskedTypeQuestionV2Asked EventListResponseEventQuestionV2AskedType = "question.v2.asked"
)

func (r EventListResponseEventQuestionV2AskedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionV2AskedTypeQuestionV2Asked:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventQuestionV2Replied
// =============================================================================

type EventListResponseEventQuestionV2Replied struct {
	ID         string                                            `json:"id,required"`
	Properties EventListResponseEventQuestionV2RepliedProperties `json:"properties,required"`
	Type       EventListResponseEventQuestionV2RepliedType       `json:"type,required"`
	JSON       eventListResponseEventQuestionV2RepliedJSON       `json:"-"`
}

type eventListResponseEventQuestionV2RepliedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionV2Replied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionV2RepliedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionV2Replied) implementsEventListResponse()  {}
func (r EventListResponseEventQuestionV2Replied) implementsGlobalEventPayload() {}

type EventListResponseEventQuestionV2RepliedProperties struct {
	SessionID string                                                `json:"sessionID,required"`
	RequestID string                                                `json:"requestID,required"`
	Answers   [][]string                                            `json:"answers,required"`
	JSON      eventListResponseEventQuestionV2RepliedPropertiesJSON `json:"-"`
}

type eventListResponseEventQuestionV2RepliedPropertiesJSON struct {
	SessionID   apijson.Field
	RequestID   apijson.Field
	Answers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionV2RepliedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionV2RepliedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionV2RepliedType string

const (
	EventListResponseEventQuestionV2RepliedTypeQuestionV2Replied EventListResponseEventQuestionV2RepliedType = "question.v2.replied"
)

func (r EventListResponseEventQuestionV2RepliedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionV2RepliedTypeQuestionV2Replied:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventQuestionV2Rejected
// =============================================================================

type EventListResponseEventQuestionV2Rejected struct {
	ID         string                                             `json:"id,required"`
	Properties EventListResponseEventQuestionV2RejectedProperties `json:"properties,required"`
	Type       EventListResponseEventQuestionV2RejectedType       `json:"type,required"`
	JSON       eventListResponseEventQuestionV2RejectedJSON       `json:"-"`
}

type eventListResponseEventQuestionV2RejectedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionV2Rejected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionV2RejectedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionV2Rejected) implementsEventListResponse()  {}
func (r EventListResponseEventQuestionV2Rejected) implementsGlobalEventPayload() {}

type EventListResponseEventQuestionV2RejectedProperties struct {
	SessionID string                                                 `json:"sessionID,required"`
	RequestID string                                                 `json:"requestID,required"`
	JSON      eventListResponseEventQuestionV2RejectedPropertiesJSON `json:"-"`
}

type eventListResponseEventQuestionV2RejectedPropertiesJSON struct {
	SessionID   apijson.Field
	RequestID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionV2RejectedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionV2RejectedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionV2RejectedType string

const (
	EventListResponseEventQuestionV2RejectedTypeQuestionV2Rejected EventListResponseEventQuestionV2RejectedType = "question.v2.rejected"
)

func (r EventListResponseEventQuestionV2RejectedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionV2RejectedTypeQuestionV2Rejected:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionNextMoved
// =============================================================================

type EventListResponseEventSessionNextMoved struct {
	ID         string                                           `json:"id,required"`
	Properties EventListResponseEventSessionNextMovedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextMovedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextMovedJSON       `json:"-"`
}

type eventListResponseEventSessionNextMovedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextMoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextMovedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextMoved) implementsEventListResponse()  {}
func (r EventListResponseEventSessionNextMoved) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextMovedProperties struct {
	Timestamp    int64                                                `json:"timestamp,required"`
	SessionID    string                                               `json:"sessionID,required"`
	Location     LocationRef                                          `json:"location,required"`
	Subdirectory string                                               `json:"subdirectory"`
	JSON         eventListResponseEventSessionNextMovedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextMovedPropertiesJSON struct {
	Timestamp    apijson.Field
	SessionID    apijson.Field
	Location     apijson.Field
	Subdirectory apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextMovedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextMovedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextMovedType string

const (
	EventListResponseEventSessionNextMovedTypeSessionNextMoved EventListResponseEventSessionNextMovedType = "session.next.moved"
)

func (r EventListResponseEventSessionNextMovedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextMovedTypeSessionNextMoved:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionNextRevertStaged
// =============================================================================

type EventListResponseEventSessionNextRevertStaged struct {
	ID         string                                                  `json:"id,required"`
	Properties EventListResponseEventSessionNextRevertStagedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextRevertStagedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextRevertStagedJSON       `json:"-"`
}

type eventListResponseEventSessionNextRevertStagedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextRevertStaged) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextRevertStagedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextRevertStaged) implementsEventListResponse()  {}
func (r EventListResponseEventSessionNextRevertStaged) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextRevertStagedProperties struct {
	Timestamp int64                                                       `json:"timestamp,required"`
	SessionID string                                                      `json:"sessionID,required"`
	Revert    RevertState                                                 `json:"revert,required"`
	JSON      eventListResponseEventSessionNextRevertStagedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextRevertStagedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Revert      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextRevertStagedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextRevertStagedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextRevertStagedType string

const (
	EventListResponseEventSessionNextRevertStagedTypeSessionNextRevertStaged EventListResponseEventSessionNextRevertStagedType = "session.next.revert.staged"
)

func (r EventListResponseEventSessionNextRevertStagedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextRevertStagedTypeSessionNextRevertStaged:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionNextRevertCleared
// =============================================================================

type EventListResponseEventSessionNextRevertCleared struct {
	ID         string                                                   `json:"id,required"`
	Properties EventListResponseEventSessionNextRevertClearedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextRevertClearedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextRevertClearedJSON       `json:"-"`
}

type eventListResponseEventSessionNextRevertClearedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextRevertCleared) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextRevertClearedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextRevertCleared) implementsEventListResponse()  {}
func (r EventListResponseEventSessionNextRevertCleared) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextRevertClearedProperties struct {
	Timestamp int64                                                        `json:"timestamp,required"`
	SessionID string                                                       `json:"sessionID,required"`
	JSON      eventListResponseEventSessionNextRevertClearedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextRevertClearedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextRevertClearedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextRevertClearedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextRevertClearedType string

const (
	EventListResponseEventSessionNextRevertClearedTypeSessionNextRevertCleared EventListResponseEventSessionNextRevertClearedType = "session.next.revert.cleared"
)

func (r EventListResponseEventSessionNextRevertClearedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextRevertClearedTypeSessionNextRevertCleared:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionNextRevertCommitted
// =============================================================================

type EventListResponseEventSessionNextRevertCommitted struct {
	ID         string                                                     `json:"id,required"`
	Properties EventListResponseEventSessionNextRevertCommittedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextRevertCommittedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextRevertCommittedJSON       `json:"-"`
}

type eventListResponseEventSessionNextRevertCommittedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextRevertCommitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextRevertCommittedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextRevertCommitted) implementsEventListResponse()  {}
func (r EventListResponseEventSessionNextRevertCommitted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextRevertCommittedProperties struct {
	Timestamp int64                                                          `json:"timestamp,required"`
	SessionID string                                                         `json:"sessionID,required"`
	MessageID string                                                         `json:"messageID,required"`
	JSON      eventListResponseEventSessionNextRevertCommittedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextRevertCommittedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextRevertCommittedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextRevertCommittedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextRevertCommittedType string

const (
	EventListResponseEventSessionNextRevertCommittedTypeSessionNextRevertCommitted EventListResponseEventSessionNextRevertCommittedType = "session.next.revert.committed"
)

func (r EventListResponseEventSessionNextRevertCommittedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextRevertCommittedTypeSessionNextRevertCommitted:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionNextPromptAdmitted
// =============================================================================

type EventListResponseEventSessionNextPromptAdmitted struct {
	ID         string                                                    `json:"id,required"`
	Properties EventListResponseEventSessionNextPromptAdmittedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextPromptAdmittedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextPromptAdmittedJSON       `json:"-"`
}

type eventListResponseEventSessionNextPromptAdmittedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextPromptAdmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextPromptAdmittedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextPromptAdmitted) implementsEventListResponse()  {}
func (r EventListResponseEventSessionNextPromptAdmitted) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextPromptAdmittedDelivery string

const (
	EventListResponseEventSessionNextPromptAdmittedDeliverySteer EventListResponseEventSessionNextPromptAdmittedDelivery = "steer"
	EventListResponseEventSessionNextPromptAdmittedDeliveryQueue EventListResponseEventSessionNextPromptAdmittedDelivery = "queue"
)

func (r EventListResponseEventSessionNextPromptAdmittedDelivery) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextPromptAdmittedDeliverySteer:
		return true
	case EventListResponseEventSessionNextPromptAdmittedDeliveryQueue:
		return true
	}
	return false
}

type EventListResponseEventSessionNextPromptAdmittedProperties struct {
	Timestamp int64                                                         `json:"timestamp,required"`
	SessionID string                                                        `json:"sessionID,required"`
	MessageID string                                                        `json:"messageID,required"`
	Prompt    V2SessionInputPrompt                                          `json:"prompt,required"`
	Delivery  EventListResponseEventSessionNextPromptAdmittedDelivery       `json:"delivery,required"`
	JSON      eventListResponseEventSessionNextPromptAdmittedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextPromptAdmittedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	MessageID   apijson.Field
	Prompt      apijson.Field
	Delivery    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextPromptAdmittedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextPromptAdmittedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextPromptAdmittedType string

const (
	EventListResponseEventSessionNextPromptAdmittedTypeSessionNextPromptAdmitted EventListResponseEventSessionNextPromptAdmittedType = "session.next.prompt.admitted"
)

func (r EventListResponseEventSessionNextPromptAdmittedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextPromptAdmittedTypeSessionNextPromptAdmitted:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionNextContextUpdated
// =============================================================================

type EventListResponseEventSessionNextContextUpdated struct {
	ID         string                                                    `json:"id,required"`
	Properties EventListResponseEventSessionNextContextUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventSessionNextContextUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventSessionNextContextUpdatedJSON       `json:"-"`
}

type eventListResponseEventSessionNextContextUpdatedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextContextUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextContextUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionNextContextUpdated) implementsEventListResponse()  {}
func (r EventListResponseEventSessionNextContextUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventSessionNextContextUpdatedProperties struct {
	Timestamp int64                                                         `json:"timestamp,required"`
	SessionID string                                                        `json:"sessionID,required"`
	MessageID string                                                        `json:"messageID,required"`
	Text      string                                                        `json:"text,required"`
	JSON      eventListResponseEventSessionNextContextUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextContextUpdatedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	MessageID   apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextContextUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextContextUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextContextUpdatedType string

const (
	EventListResponseEventSessionNextContextUpdatedTypeSessionNextContextUpdated EventListResponseEventSessionNextContextUpdatedType = "session.next.context.updated"
)

func (r EventListResponseEventSessionNextContextUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionNextContextUpdatedTypeSessionNextContextUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventProjectDirectoriesUpdated
// =============================================================================

type EventListResponseEventProjectDirectoriesUpdated struct {
	ID         string                                                    `json:"id,required"`
	Properties EventListResponseEventProjectDirectoriesUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventProjectDirectoriesUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventProjectDirectoriesUpdatedJSON       `json:"-"`
}

type eventListResponseEventProjectDirectoriesUpdatedJSON struct {
	ID          apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventProjectDirectoriesUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventProjectDirectoriesUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventProjectDirectoriesUpdated) implementsEventListResponse()  {}
func (r EventListResponseEventProjectDirectoriesUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventProjectDirectoriesUpdatedProperties struct {
	ProjectID string                                                        `json:"projectID,required"`
	JSON      eventListResponseEventProjectDirectoriesUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventProjectDirectoriesUpdatedPropertiesJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventProjectDirectoriesUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventProjectDirectoriesUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventProjectDirectoriesUpdatedType string

const (
	EventListResponseEventProjectDirectoriesUpdatedTypeProjectDirectoriesUpdated EventListResponseEventProjectDirectoriesUpdatedType = "project.directories.updated"
)

func (r EventListResponseEventProjectDirectoriesUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventProjectDirectoriesUpdatedTypeProjectDirectoriesUpdated:
		return true
	}
	return false
}
