package opencode

import (
	"github.com/sst/opencode-sdk-go/internal/apijson"
)

// =============================================================================
// EventListResponseEventInstallationUpdateAvailable
// =============================================================================

type EventListResponseEventInstallationUpdateAvailable struct {
	Properties EventListResponseEventInstallationUpdateAvailableProperties `json:"properties,required"`
	Type       EventListResponseEventInstallationUpdateAvailableType       `json:"type,required"`
	JSON       eventListResponseEventInstallationUpdateAvailableJSON       `json:"-"`
}

type eventListResponseEventInstallationUpdateAvailableJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventInstallationUpdateAvailable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventInstallationUpdateAvailableJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventInstallationUpdateAvailable) implementsEventListResponse() {}

func (r EventListResponseEventInstallationUpdateAvailable) implementsGlobalEventPayload() {}

type EventListResponseEventInstallationUpdateAvailableProperties struct {
	Version string                                                          `json:"version,required"`
	JSON    eventListResponseEventInstallationUpdateAvailablePropertiesJSON `json:"-"`
}

type eventListResponseEventInstallationUpdateAvailablePropertiesJSON struct {
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventInstallationUpdateAvailableProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventInstallationUpdateAvailablePropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventInstallationUpdateAvailableType string

const (
	EventListResponseEventInstallationUpdateAvailableTypeInstallationUpdateAvailable EventListResponseEventInstallationUpdateAvailableType = "installation.update-available"
)

func (r EventListResponseEventInstallationUpdateAvailableType) IsKnown() bool {
	switch r {
	case EventListResponseEventInstallationUpdateAvailableTypeInstallationUpdateAvailable:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventProjectUpdated
// =============================================================================

type EventListResponseEventProjectUpdated struct {
	Properties EventListResponseEventProjectUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventProjectUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventProjectUpdatedJSON       `json:"-"`
}

type eventListResponseEventProjectUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventProjectUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventProjectUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventProjectUpdated) implementsEventListResponse() {}

func (r EventListResponseEventProjectUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventProjectUpdatedProperties struct {
	Project Project `json:"project,required"`
	JSON    eventListResponseEventProjectUpdatedPropertiesJSON
}

type eventListResponseEventProjectUpdatedPropertiesJSON struct {
	Project     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventProjectUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventProjectUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventProjectUpdatedType string

const (
	EventListResponseEventProjectUpdatedTypeProjectUpdated EventListResponseEventProjectUpdatedType = "project.updated"
)

func (r EventListResponseEventProjectUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventProjectUpdatedTypeProjectUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventServerInstanceDisposed
// =============================================================================

type EventListResponseEventServerInstanceDisposed struct {
	Properties EventListResponseEventServerInstanceDisposedProperties `json:"properties,required"`
	Type       EventListResponseEventServerInstanceDisposedType       `json:"type,required"`
	JSON       eventListResponseEventServerInstanceDisposedJSON       `json:"-"`
}

type eventListResponseEventServerInstanceDisposedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventServerInstanceDisposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventServerInstanceDisposedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventServerInstanceDisposed) implementsEventListResponse() {}

func (r EventListResponseEventServerInstanceDisposed) implementsGlobalEventPayload() {}

type EventListResponseEventServerInstanceDisposedProperties struct {
	Directory string                                                     `json:"directory,required"`
	JSON      eventListResponseEventServerInstanceDisposedPropertiesJSON `json:"-"`
}

type eventListResponseEventServerInstanceDisposedPropertiesJSON struct {
	Directory   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventServerInstanceDisposedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventServerInstanceDisposedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventServerInstanceDisposedType string

const (
	EventListResponseEventServerInstanceDisposedTypeServerInstanceDisposed EventListResponseEventServerInstanceDisposedType = "server.instance.disposed"
)

func (r EventListResponseEventServerInstanceDisposedType) IsKnown() bool {
	switch r {
	case EventListResponseEventServerInstanceDisposedTypeServerInstanceDisposed:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventGlobalDisposed
// =============================================================================

type EventListResponseEventGlobalDisposed struct {
	Properties EventListResponseEventGlobalDisposedProperties `json:"properties,required"`
	Type       EventListResponseEventGlobalDisposedType       `json:"type,required"`
	JSON       eventListResponseEventGlobalDisposedJSON       `json:"-"`
}

type eventListResponseEventGlobalDisposedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventGlobalDisposed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventGlobalDisposedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventGlobalDisposed) implementsEventListResponse() {}

func (r EventListResponseEventGlobalDisposed) implementsGlobalEventPayload() {}

type EventListResponseEventGlobalDisposedProperties struct {
	JSON eventListResponseEventGlobalDisposedPropertiesJSON `json:"-"`
}

type eventListResponseEventGlobalDisposedPropertiesJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventGlobalDisposedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventGlobalDisposedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventGlobalDisposedType string

const (
	EventListResponseEventGlobalDisposedTypeGlobalDisposed EventListResponseEventGlobalDisposedType = "global.disposed"
)

func (r EventListResponseEventGlobalDisposedType) IsKnown() bool {
	switch r {
	case EventListResponseEventGlobalDisposedTypeGlobalDisposed:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventLspUpdated
// =============================================================================

type EventListResponseEventLspUpdated struct {
	Properties EventListResponseEventLspUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventLspUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventLspUpdatedJSON       `json:"-"`
}

type eventListResponseEventLspUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventLspUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventLspUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventLspUpdated) implementsEventListResponse() {}

func (r EventListResponseEventLspUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventLspUpdatedProperties struct {
	JSON eventListResponseEventLspUpdatedPropertiesJSON `json:"-"`
}

type eventListResponseEventLspUpdatedPropertiesJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventLspUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventLspUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventLspUpdatedType string

const (
	EventListResponseEventLspUpdatedTypeLspUpdated EventListResponseEventLspUpdatedType = "lsp.updated"
)

func (r EventListResponseEventLspUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventLspUpdatedTypeLspUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventMessagePartDelta
// =============================================================================

type EventListResponseEventMessagePartDelta struct {
	Properties EventListResponseEventMessagePartDeltaProperties `json:"properties,required"`
	Type       EventListResponseEventMessagePartDeltaType       `json:"type,required"`
	JSON       eventListResponseEventMessagePartDeltaJSON       `json:"-"`
}

type eventListResponseEventMessagePartDeltaJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessagePartDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessagePartDeltaJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMessagePartDelta) implementsEventListResponse() {}

func (r EventListResponseEventMessagePartDelta) implementsGlobalEventPayload() {}

type EventListResponseEventMessagePartDeltaProperties struct {
	Delta     string `json:"delta,required"`
	Field     string `json:"field,required"`
	MessageID string `json:"messageID,required"`
	PartID    string `json:"partID,required"`
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventMessagePartDeltaPropertiesJSON
}

type eventListResponseEventMessagePartDeltaPropertiesJSON struct {
	Delta       apijson.Field
	Field       apijson.Field
	MessageID   apijson.Field
	PartID      apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMessagePartDeltaProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMessagePartDeltaPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMessagePartDeltaType string

const (
	EventListResponseEventMessagePartDeltaTypeMessagePartDelta EventListResponseEventMessagePartDeltaType = "message.part.delta"
)

func (r EventListResponseEventMessagePartDeltaType) IsKnown() bool {
	switch r {
	case EventListResponseEventMessagePartDeltaTypeMessagePartDelta:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPermissionAsked
// =============================================================================

type EventListResponseEventPermissionAsked struct {
	Properties EventListResponseEventPermissionAskedProperties `json:"properties,required"`
	Type       EventListResponseEventPermissionAskedType       `json:"type,required"`
	JSON       eventListResponseEventPermissionAskedJSON       `json:"-"`
}

type eventListResponseEventPermissionAskedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionAsked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionAskedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPermissionAsked) implementsEventListResponse() {}

func (r EventListResponseEventPermissionAsked) implementsGlobalEventPayload() {}

type EventListResponseEventPermissionAskedProperties struct {
	Always     []string                                             `json:"always,required"`
	ID         string                                               `json:"id,required"`
	Metadata   map[string]interface{}                               `json:"metadata,required"`
	Patterns   []string                                             `json:"patterns,required"`
	Permission string                                               `json:"permission,required"`
	SessionID  string                                               `json:"sessionID,required"`
	Tool       *EventListResponseEventPermissionAskedPropertiesTool `json:"tool"`
	JSON       eventListResponseEventPermissionAskedPropertiesJSON  `json:"-"`
}

type eventListResponseEventPermissionAskedPropertiesJSON struct {
	Always      apijson.Field
	ID          apijson.Field
	Metadata    apijson.Field
	Patterns    apijson.Field
	Permission  apijson.Field
	SessionID   apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionAskedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionAskedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPermissionAskedPropertiesTool struct {
	CallID    string `json:"callID,required"`
	MessageID string `json:"messageID,required"`
	JSON      eventListResponseEventPermissionAskedPropertiesToolJSON
}

type eventListResponseEventPermissionAskedPropertiesToolJSON struct {
	CallID      apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPermissionAskedPropertiesTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPermissionAskedPropertiesToolJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPermissionAskedType string

const (
	EventListResponseEventPermissionAskedTypePermissionAsked EventListResponseEventPermissionAskedType = "permission.asked"
)

func (r EventListResponseEventPermissionAskedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPermissionAskedTypePermissionAsked:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionStatus
// =============================================================================

type EventListResponseEventSessionStatus struct {
	Properties EventListResponseEventSessionStatusProperties `json:"properties,required"`
	Type       EventListResponseEventSessionStatusType       `json:"type,required"`
	JSON       eventListResponseEventSessionStatusJSON       `json:"-"`
}

type eventListResponseEventSessionStatusJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionStatusJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionStatus) implementsEventListResponse() {}

func (r EventListResponseEventSessionStatus) implementsGlobalEventPayload() {}

type EventListResponseEventSessionStatusProperties struct {
	SessionID string                                    `json:"sessionID,required"`
	Status    EventListResponseEventSessionStatusStatus `json:"status,required"`
	JSON      eventListResponseEventSessionStatusPropertiesJSON
}

type eventListResponseEventSessionStatusPropertiesJSON struct {
	SessionID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionStatusProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionStatusPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionStatusStatus struct {
	Attempt *int    `json:"attempt,omitempty"`
	Message *string `json:"message,omitempty"`
	Next    *int    `json:"next,omitempty"`
	Type    string  `json:"type,required"`
	JSON    eventListResponseEventSessionStatusStatusJSON
}

type eventListResponseEventSessionStatusStatusJSON struct {
	Attempt     apijson.Field
	Message     apijson.Field
	Next        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionStatusStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionStatusStatusJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionStatusType string

const (
	EventListResponseEventSessionStatusTypeSessionStatus EventListResponseEventSessionStatusType = "session.status"
)

func (r EventListResponseEventSessionStatusType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionStatusTypeSessionStatus:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventQuestionAsked
// =============================================================================

type EventListResponseEventQuestionAsked struct {
	Properties EventListResponseEventQuestionAskedProperties `json:"properties,required"`
	Type       EventListResponseEventQuestionAskedType       `json:"type,required"`
	JSON       eventListResponseEventQuestionAskedJSON       `json:"-"`
}

type eventListResponseEventQuestionAskedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAsked) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionAsked) implementsEventListResponse() {}

func (r EventListResponseEventQuestionAsked) implementsGlobalEventPayload() {}

type EventListResponseEventQuestionAskedProperties struct {
	ID        string                                                   `json:"id,required"`
	Questions []EventListResponseEventQuestionAskedPropertiesQuestions `json:"questions,required"`
	SessionID string                                                   `json:"sessionID,required"`
	Tool      *EventListResponseEventQuestionAskedPropertiesTool       `json:"tool"`
	JSON      eventListResponseEventQuestionAskedPropertiesJSON        `json:"-"`
}

type eventListResponseEventQuestionAskedPropertiesJSON struct {
	ID          apijson.Field
	Questions   apijson.Field
	SessionID   apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAskedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionAskedPropertiesQuestions struct {
	Custom   *bool                                                           `json:"custom,omitempty"`
	Header   string                                                          `json:"header,required"`
	Multiple *bool                                                           `json:"multiple,omitempty"`
	Options  []EventListResponseEventQuestionAskedPropertiesQuestionsOptions `json:"options,required"`
	Question string                                                          `json:"question,required"`
	JSON     eventListResponseEventQuestionAskedPropertiesQuestionsJSON
}

type eventListResponseEventQuestionAskedPropertiesQuestionsJSON struct {
	Custom      apijson.Field
	Header      apijson.Field
	Multiple    apijson.Field
	Options     apijson.Field
	Question    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAskedPropertiesQuestions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedPropertiesQuestionsJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionAskedPropertiesQuestionsOptions struct {
	Description string `json:"description,required"`
	Label       string `json:"label,required"`
	JSON        eventListResponseEventQuestionAskedPropertiesQuestionsOptionsJSON
}

type eventListResponseEventQuestionAskedPropertiesQuestionsOptionsJSON struct {
	Description apijson.Field
	Label       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAskedPropertiesQuestionsOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedPropertiesQuestionsOptionsJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionAskedPropertiesTool struct {
	CallID    string `json:"callID,required"`
	MessageID string `json:"messageID,required"`
	JSON      eventListResponseEventQuestionAskedPropertiesToolJSON
}

type eventListResponseEventQuestionAskedPropertiesToolJSON struct {
	CallID      apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionAskedPropertiesTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionAskedPropertiesToolJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionAskedType string

const (
	EventListResponseEventQuestionAskedTypeQuestionAsked EventListResponseEventQuestionAskedType = "question.asked"
)

func (r EventListResponseEventQuestionAskedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionAskedTypeQuestionAsked:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventQuestionRejected
// =============================================================================

type EventListResponseEventQuestionRejected struct {
	Properties EventListResponseEventQuestionRejectedProperties `json:"properties,required"`
	Type       EventListResponseEventQuestionRejectedType       `json:"type,required"`
	JSON       eventListResponseEventQuestionRejectedJSON       `json:"-"`
}

type eventListResponseEventQuestionRejectedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionRejected) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRejectedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionRejected) implementsEventListResponse() {}

func (r EventListResponseEventQuestionRejected) implementsGlobalEventPayload() {}

type EventListResponseEventQuestionRejectedProperties struct {
	RequestID string `json:"requestID,required"`
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventQuestionRejectedPropertiesJSON
}

type eventListResponseEventQuestionRejectedPropertiesJSON struct {
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionRejectedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRejectedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionRejectedType string

const (
	EventListResponseEventQuestionRejectedTypeQuestionRejected EventListResponseEventQuestionRejectedType = "question.rejected"
)

func (r EventListResponseEventQuestionRejectedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionRejectedTypeQuestionRejected:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventQuestionReplied
// =============================================================================

type EventListResponseEventQuestionReplied struct {
	Properties EventListResponseEventQuestionRepliedProperties `json:"properties,required"`
	Type       EventListResponseEventQuestionRepliedType       `json:"type,required"`
	JSON       eventListResponseEventQuestionRepliedJSON       `json:"-"`
}

type eventListResponseEventQuestionRepliedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionReplied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRepliedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventQuestionReplied) implementsEventListResponse() {}

func (r EventListResponseEventQuestionReplied) implementsGlobalEventPayload() {}

type EventListResponseEventQuestionRepliedProperties struct {
	Answers   [][]string `json:"answers,required"`
	RequestID string     `json:"requestID,required"`
	SessionID string     `json:"sessionID,required"`
	JSON      eventListResponseEventQuestionRepliedPropertiesJSON
}

type eventListResponseEventQuestionRepliedPropertiesJSON struct {
	Answers     apijson.Field
	RequestID   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventQuestionRepliedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventQuestionRepliedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventQuestionRepliedType string

const (
	EventListResponseEventQuestionRepliedTypeQuestionReplied EventListResponseEventQuestionRepliedType = "question.replied"
)

func (r EventListResponseEventQuestionRepliedType) IsKnown() bool {
	switch r {
	case EventListResponseEventQuestionRepliedTypeQuestionReplied:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventTuiPromptAppend
// =============================================================================

type EventListResponseEventTuiPromptAppend struct {
	Properties EventListResponseEventTuiPromptAppendProperties `json:"properties,required"`
	Type       EventListResponseEventTuiPromptAppendType       `json:"type,required"`
	JSON       eventListResponseEventTuiPromptAppendJSON       `json:"-"`
}

type eventListResponseEventTuiPromptAppendJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiPromptAppend) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiPromptAppendJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiPromptAppend) implementsEventListResponse() {}

func (r EventListResponseEventTuiPromptAppend) implementsGlobalEventPayload() {}

type EventListResponseEventTuiPromptAppendProperties struct {
	Text string `json:"text,required"`
	JSON eventListResponseEventTuiPromptAppendPropertiesJSON
}

type eventListResponseEventTuiPromptAppendPropertiesJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiPromptAppendProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiPromptAppendPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiPromptAppendType string

const (
	EventListResponseEventTuiPromptAppendTypeTuiPromptAppend EventListResponseEventTuiPromptAppendType = "tui.prompt.append"
)

func (r EventListResponseEventTuiPromptAppendType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiPromptAppendTypeTuiPromptAppend:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventTuiCommandExecute
// =============================================================================

type EventListResponseEventTuiCommandExecute struct {
	Properties EventListResponseEventTuiCommandExecuteProperties `json:"properties,required"`
	Type       EventListResponseEventTuiCommandExecuteType       `json:"type,required"`
	JSON       eventListResponseEventTuiCommandExecuteJSON       `json:"-"`
}

type eventListResponseEventTuiCommandExecuteJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiCommandExecute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiCommandExecuteJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiCommandExecute) implementsEventListResponse() {}

func (r EventListResponseEventTuiCommandExecute) implementsGlobalEventPayload() {}

type EventListResponseEventTuiCommandExecuteProperties struct {
	Command string `json:"command,required"`
	JSON    eventListResponseEventTuiCommandExecutePropertiesJSON
}

type eventListResponseEventTuiCommandExecutePropertiesJSON struct {
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiCommandExecuteProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiCommandExecutePropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiCommandExecuteType string

const (
	EventListResponseEventTuiCommandExecuteTypeTuiCommandExecute EventListResponseEventTuiCommandExecuteType = "tui.command.execute"
)

func (r EventListResponseEventTuiCommandExecuteType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiCommandExecuteTypeTuiCommandExecute:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventTuiToastShow
// =============================================================================

type EventListResponseEventTuiToastShow struct {
	Properties EventListResponseEventTuiToastShowProperties `json:"properties,required"`
	Type       EventListResponseEventTuiToastShowType       `json:"type,required"`
	JSON       eventListResponseEventTuiToastShowJSON       `json:"-"`
}

type eventListResponseEventTuiToastShowJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiToastShow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiToastShowJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiToastShow) implementsEventListResponse() {}

func (r EventListResponseEventTuiToastShow) implementsGlobalEventPayload() {}

type EventListResponseEventTuiToastShowProperties struct {
	Duration *int    `json:"duration,omitempty"`
	Message  string  `json:"message,required"`
	Title    *string `json:"title,omitempty"`
	Variant  string  `json:"variant,required"`
	JSON     eventListResponseEventTuiToastShowPropertiesJSON
}

type eventListResponseEventTuiToastShowPropertiesJSON struct {
	Duration    apijson.Field
	Message     apijson.Field
	Title       apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiToastShowProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiToastShowPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiToastShowType string

const (
	EventListResponseEventTuiToastShowTypeTuiToastShow EventListResponseEventTuiToastShowType = "tui.toast.show"
)

func (r EventListResponseEventTuiToastShowType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiToastShowTypeTuiToastShow:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventTuiSessionSelect
// =============================================================================

type EventListResponseEventTuiSessionSelect struct {
	Properties EventListResponseEventTuiSessionSelectProperties `json:"properties,required"`
	Type       EventListResponseEventTuiSessionSelectType       `json:"type,required"`
	JSON       eventListResponseEventTuiSessionSelectJSON       `json:"-"`
}

type eventListResponseEventTuiSessionSelectJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiSessionSelect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiSessionSelectJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventTuiSessionSelect) implementsEventListResponse() {}

func (r EventListResponseEventTuiSessionSelect) implementsGlobalEventPayload() {}

type EventListResponseEventTuiSessionSelectProperties struct {
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventTuiSessionSelectPropertiesJSON
}

type eventListResponseEventTuiSessionSelectPropertiesJSON struct {
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventTuiSessionSelectProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventTuiSessionSelectPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventTuiSessionSelectType string

const (
	EventListResponseEventTuiSessionSelectTypeTuiSessionSelect EventListResponseEventTuiSessionSelectType = "tui.session.select"
)

func (r EventListResponseEventTuiSessionSelectType) IsKnown() bool {
	switch r {
	case EventListResponseEventTuiSessionSelectTypeTuiSessionSelect:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventMcpToolsChanged
// =============================================================================

type EventListResponseEventMcpToolsChanged struct {
	Properties EventListResponseEventMcpToolsChangedProperties `json:"properties,required"`
	Type       EventListResponseEventMcpToolsChangedType       `json:"type,required"`
	JSON       eventListResponseEventMcpToolsChangedJSON       `json:"-"`
}

type eventListResponseEventMcpToolsChangedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpToolsChanged) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpToolsChangedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMcpToolsChanged) implementsEventListResponse() {}

func (r EventListResponseEventMcpToolsChanged) implementsGlobalEventPayload() {}

type EventListResponseEventMcpToolsChangedProperties struct {
	Server string `json:"server,required"`
	JSON   eventListResponseEventMcpToolsChangedPropertiesJSON
}

type eventListResponseEventMcpToolsChangedPropertiesJSON struct {
	Server      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpToolsChangedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpToolsChangedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMcpToolsChangedType string

const (
	EventListResponseEventMcpToolsChangedTypeMcpToolsChanged EventListResponseEventMcpToolsChangedType = "mcp.tools.changed"
)

func (r EventListResponseEventMcpToolsChangedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMcpToolsChangedTypeMcpToolsChanged:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventMcpBrowserOpenFailed
// =============================================================================

type EventListResponseEventMcpBrowserOpenFailed struct {
	Properties EventListResponseEventMcpBrowserOpenFailedProperties `json:"properties,required"`
	Type       EventListResponseEventMcpBrowserOpenFailedType       `json:"type,required"`
	JSON       eventListResponseEventMcpBrowserOpenFailedJSON       `json:"-"`
}

type eventListResponseEventMcpBrowserOpenFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpBrowserOpenFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpBrowserOpenFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventMcpBrowserOpenFailed) implementsEventListResponse() {}

func (r EventListResponseEventMcpBrowserOpenFailed) implementsGlobalEventPayload() {}

type EventListResponseEventMcpBrowserOpenFailedProperties struct {
	McpName string `json:"mcpName,required"`
	URL     string `json:"url,required"`
	JSON    eventListResponseEventMcpBrowserOpenFailedPropertiesJSON
}

type eventListResponseEventMcpBrowserOpenFailedPropertiesJSON struct {
	McpName     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventMcpBrowserOpenFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventMcpBrowserOpenFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventMcpBrowserOpenFailedType string

const (
	EventListResponseEventMcpBrowserOpenFailedTypeMcpBrowserOpenFailed EventListResponseEventMcpBrowserOpenFailedType = "mcp.browser.open.failed"
)

func (r EventListResponseEventMcpBrowserOpenFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventMcpBrowserOpenFailedTypeMcpBrowserOpenFailed:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventCommandExecuted
// =============================================================================

type EventListResponseEventCommandExecuted struct {
	Properties EventListResponseEventCommandExecutedProperties `json:"properties,required"`
	Type       EventListResponseEventCommandExecutedType       `json:"type,required"`
	JSON       eventListResponseEventCommandExecutedJSON       `json:"-"`
}

type eventListResponseEventCommandExecutedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCommandExecuted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCommandExecutedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventCommandExecuted) implementsEventListResponse() {}

func (r EventListResponseEventCommandExecuted) implementsGlobalEventPayload() {}

type EventListResponseEventCommandExecutedProperties struct {
	Arguments string `json:"arguments,required"`
	MessageID string `json:"messageID,required"`
	Name      string `json:"name,required"`
	SessionID string `json:"sessionID,required"`
	JSON      eventListResponseEventCommandExecutedPropertiesJSON
}

type eventListResponseEventCommandExecutedPropertiesJSON struct {
	Arguments   apijson.Field
	MessageID   apijson.Field
	Name        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventCommandExecutedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventCommandExecutedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventCommandExecutedType string

const (
	EventListResponseEventCommandExecutedTypeCommandExecuted EventListResponseEventCommandExecutedType = "command.executed"
)

func (r EventListResponseEventCommandExecutedType) IsKnown() bool {
	switch r {
	case EventListResponseEventCommandExecutedTypeCommandExecuted:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionDiff
// =============================================================================

type EventListResponseEventSessionDiff struct {
	Properties EventListResponseEventSessionDiffProperties `json:"properties,required"`
	Type       EventListResponseEventSessionDiffType       `json:"type,required"`
	JSON       eventListResponseEventSessionDiffJSON       `json:"-"`
}

type eventListResponseEventSessionDiffJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionDiffJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventSessionDiff) implementsEventListResponse() {}

func (r EventListResponseEventSessionDiff) implementsGlobalEventPayload() {}

type EventListResponseEventSessionDiffProperties struct {
	Diff      []EventListResponseEventSessionDiffPropertiesDiff `json:"diff,required"`
	SessionID string                                            `json:"sessionID,required"`
	JSON      eventListResponseEventSessionDiffPropertiesJSON
}

type eventListResponseEventSessionDiffPropertiesJSON struct {
	Diff        apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionDiffProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionDiffPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionDiffPropertiesDiff struct {
	Additions int    `json:"additions,required"`
	Deletions int    `json:"deletions,required"`
	File      string `json:"file"`
	Patch     string `json:"patch"`
	Status    string `json:"status,omitempty"`
	JSON      eventListResponseEventSessionDiffPropertiesDiffJSON
}

type eventListResponseEventSessionDiffPropertiesDiffJSON struct {
	Additions   apijson.Field
	Deletions   apijson.Field
	File        apijson.Field
	Patch       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionDiffPropertiesDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionDiffPropertiesDiffJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionDiffType string

const (
	EventListResponseEventSessionDiffTypeSessionDiff EventListResponseEventSessionDiffType = "session.diff"
)

func (r EventListResponseEventSessionDiffType) IsKnown() bool {
	switch r {
	case EventListResponseEventSessionDiffTypeSessionDiff:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventVcsBranchUpdated
// =============================================================================

type EventListResponseEventVcsBranchUpdated struct {
	Properties EventListResponseEventVcsBranchUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventVcsBranchUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventVcsBranchUpdatedJSON       `json:"-"`
}

type eventListResponseEventVcsBranchUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventVcsBranchUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventVcsBranchUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventVcsBranchUpdated) implementsEventListResponse() {}

func (r EventListResponseEventVcsBranchUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventVcsBranchUpdatedProperties struct {
	Branch string                                            `json:"branch"`
	JSON   eventListResponseEventVcsBranchUpdatedPropertiesJSON
}

type eventListResponseEventVcsBranchUpdatedPropertiesJSON struct {
	Branch      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventVcsBranchUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventVcsBranchUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventVcsBranchUpdatedType string

const (
	EventListResponseEventVcsBranchUpdatedTypeVcsBranchUpdated EventListResponseEventVcsBranchUpdatedType = "vcs.branch.updated"
)

func (r EventListResponseEventVcsBranchUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventVcsBranchUpdatedTypeVcsBranchUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventWorkspaceReady
// =============================================================================

type EventListResponseEventWorkspaceReady struct {
	Properties EventListResponseEventWorkspaceReadyProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceReadyType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceReadyJSON       `json:"-"`
}

type eventListResponseEventWorkspaceReadyJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceReady) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceReadyJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceReady) implementsEventListResponse() {}

func (r EventListResponseEventWorkspaceReady) implementsGlobalEventPayload() {}

type EventListResponseEventWorkspaceReadyProperties struct {
	Name string `json:"name,required"`
	JSON eventListResponseEventWorkspaceReadyPropertiesJSON
}

type eventListResponseEventWorkspaceReadyPropertiesJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceReadyProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceReadyPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceReadyType string

const (
	EventListResponseEventWorkspaceReadyTypeWorkspaceReady EventListResponseEventWorkspaceReadyType = "workspace.ready"
)

func (r EventListResponseEventWorkspaceReadyType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceReadyTypeWorkspaceReady:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventWorkspaceFailed
// =============================================================================

type EventListResponseEventWorkspaceFailed struct {
	Properties EventListResponseEventWorkspaceFailedProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceFailedType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceFailedJSON       `json:"-"`
}

type eventListResponseEventWorkspaceFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceFailed) implementsEventListResponse() {}

func (r EventListResponseEventWorkspaceFailed) implementsGlobalEventPayload() {}

type EventListResponseEventWorkspaceFailedProperties struct {
	Message string `json:"message,required"`
	JSON    eventListResponseEventWorkspaceFailedPropertiesJSON
}

type eventListResponseEventWorkspaceFailedPropertiesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceFailedType string

const (
	EventListResponseEventWorkspaceFailedTypeWorkspaceFailed EventListResponseEventWorkspaceFailedType = "workspace.failed"
)

func (r EventListResponseEventWorkspaceFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceFailedTypeWorkspaceFailed:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPtyCreated
// =============================================================================

type EventListResponseEventPtyCreated struct {
	Properties EventListResponseEventPtyCreatedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyCreatedType       `json:"type,required"`
	JSON       eventListResponseEventPtyCreatedJSON       `json:"-"`
}

type eventListResponseEventPtyCreatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyCreatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyCreated) implementsEventListResponse() {}

func (r EventListResponseEventPtyCreated) implementsGlobalEventPayload() {}

type EventListResponseEventPtyCreatedProperties struct {
	Info Pty `json:"info,required"`
	JSON eventListResponseEventPtyCreatedPropertiesJSON
}

type eventListResponseEventPtyCreatedPropertiesJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyCreatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyCreatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyCreatedType string

const (
	EventListResponseEventPtyCreatedTypePtyCreated EventListResponseEventPtyCreatedType = "pty.created"
)

func (r EventListResponseEventPtyCreatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyCreatedTypePtyCreated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPtyUpdated
// =============================================================================

type EventListResponseEventPtyUpdated struct {
	Properties EventListResponseEventPtyUpdatedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyUpdatedType       `json:"type,required"`
	JSON       eventListResponseEventPtyUpdatedJSON       `json:"-"`
}

type eventListResponseEventPtyUpdatedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyUpdated) implementsEventListResponse() {}

func (r EventListResponseEventPtyUpdated) implementsGlobalEventPayload() {}

type EventListResponseEventPtyUpdatedProperties struct {
	Info Pty `json:"info,required"`
	JSON eventListResponseEventPtyUpdatedPropertiesJSON
}

type eventListResponseEventPtyUpdatedPropertiesJSON struct {
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyUpdatedType string

const (
	EventListResponseEventPtyUpdatedTypePtyUpdated EventListResponseEventPtyUpdatedType = "pty.updated"
)

func (r EventListResponseEventPtyUpdatedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyUpdatedTypePtyUpdated:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPtyExited
// =============================================================================

type EventListResponseEventPtyExited struct {
	Properties EventListResponseEventPtyExitedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyExitedType       `json:"type,required"`
	JSON       eventListResponseEventPtyExitedJSON       `json:"-"`
}

type eventListResponseEventPtyExitedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyExited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyExitedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyExited) implementsEventListResponse() {}

func (r EventListResponseEventPtyExited) implementsGlobalEventPayload() {}

type EventListResponseEventPtyExitedProperties struct {
	ExitCode int    `json:"exitCode,required"`
	ID       string `json:"id,required"`
	JSON     eventListResponseEventPtyExitedPropertiesJSON
}

type eventListResponseEventPtyExitedPropertiesJSON struct {
	ExitCode    apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyExitedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyExitedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyExitedType string

const (
	EventListResponseEventPtyExitedTypePtyExited EventListResponseEventPtyExitedType = "pty.exited"
)

func (r EventListResponseEventPtyExitedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyExitedTypePtyExited:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventPtyDeleted
// =============================================================================

type EventListResponseEventPtyDeleted struct {
	Properties EventListResponseEventPtyDeletedProperties `json:"properties,required"`
	Type       EventListResponseEventPtyDeletedType       `json:"type,required"`
	JSON       eventListResponseEventPtyDeletedJSON       `json:"-"`
}

type eventListResponseEventPtyDeletedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyDeletedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventPtyDeleted) implementsEventListResponse() {}

func (r EventListResponseEventPtyDeleted) implementsGlobalEventPayload() {}

type EventListResponseEventPtyDeletedProperties struct {
	ID   string `json:"id,required"`
	JSON eventListResponseEventPtyDeletedPropertiesJSON
}

type eventListResponseEventPtyDeletedPropertiesJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventPtyDeletedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventPtyDeletedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventPtyDeletedType string

const (
	EventListResponseEventPtyDeletedTypePtyDeleted EventListResponseEventPtyDeletedType = "pty.deleted"
)

func (r EventListResponseEventPtyDeletedType) IsKnown() bool {
	switch r {
	case EventListResponseEventPtyDeletedTypePtyDeleted:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventWorktreeReady
// =============================================================================

type EventListResponseEventWorktreeReady struct {
	Properties EventListResponseEventWorktreeReadyProperties `json:"properties,required"`
	Type       EventListResponseEventWorktreeReadyType       `json:"type,required"`
	JSON       eventListResponseEventWorktreeReadyJSON       `json:"-"`
}

type eventListResponseEventWorktreeReadyJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeReady) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeReadyJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorktreeReady) implementsEventListResponse() {}

func (r EventListResponseEventWorktreeReady) implementsGlobalEventPayload() {}

type EventListResponseEventWorktreeReadyProperties struct {
	Branch string                                          `json:"branch"`
	Name   string                                          `json:"name,required"`
	JSON   eventListResponseEventWorktreeReadyPropertiesJSON
}

type eventListResponseEventWorktreeReadyPropertiesJSON struct {
	Branch      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeReadyProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeReadyPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorktreeReadyType string

const (
	EventListResponseEventWorktreeReadyTypeWorktreeReady EventListResponseEventWorktreeReadyType = "worktree.ready"
)

func (r EventListResponseEventWorktreeReadyType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorktreeReadyTypeWorktreeReady:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventWorktreeFailed
// =============================================================================

type EventListResponseEventWorktreeFailed struct {
	Properties EventListResponseEventWorktreeFailedProperties `json:"properties,required"`
	Type       EventListResponseEventWorktreeFailedType       `json:"type,required"`
	JSON       eventListResponseEventWorktreeFailedJSON       `json:"-"`
}

type eventListResponseEventWorktreeFailedJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeFailedJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorktreeFailed) implementsEventListResponse() {}

func (r EventListResponseEventWorktreeFailed) implementsGlobalEventPayload() {}

type EventListResponseEventWorktreeFailedProperties struct {
	Message string `json:"message,required"`
	JSON    eventListResponseEventWorktreeFailedPropertiesJSON
}

type eventListResponseEventWorktreeFailedPropertiesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorktreeFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorktreeFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorktreeFailedType string

const (
	EventListResponseEventWorktreeFailedTypeWorktreeFailed EventListResponseEventWorktreeFailedType = "worktree.failed"
)

func (r EventListResponseEventWorktreeFailedType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorktreeFailedTypeWorktreeFailed:
		return true
	}
	return false
}

type Pty struct {
	Args    []string  `json:"args,required"`
	Command string    `json:"command,required"`
	Cwd     string    `json:"cwd,required"`
	ID      string    `json:"id,required"`
	Pid     int       `json:"pid,required"`
	Status  PtyStatus `json:"status,required"`
	Title   string    `json:"title,required"`
	JSON    ptyJSON   `json:"-"`
}

type ptyJSON struct {
	Args        apijson.Field
	Command     apijson.Field
	Cwd         apijson.Field
	ID          apijson.Field
	Pid         apijson.Field
	Status      apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptyJSON) RawJSON() string {
	return r.raw
}

type PtyStatus string

const (
	PtyStatusRunning PtyStatus = "running"
	PtyStatusExited  PtyStatus = "exited"
)

func (r PtyStatus) IsKnown() bool {
	switch r {
	case PtyStatusRunning, PtyStatusExited:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventWorkspaceRestore
// =============================================================================
// EventListResponseEventWorkspaceStatus
// =============================================================================

type EventListResponseEventWorkspaceStatus struct {
	Properties EventListResponseEventWorkspaceStatusProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceStatusType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceStatusJSON       `json:"-"`
}

type eventListResponseEventWorkspaceStatusJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceStatusJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceStatus) implementsEventListResponse() {}

func (r EventListResponseEventWorkspaceStatus) implementsGlobalEventPayload() {}

type EventListResponseEventWorkspaceStatusProperties struct {
	WorkspaceID string                                      `json:"workspaceID,required"`
	Status      EventListResponseEventWorkspaceStatusStatus `json:"status,required"`
	JSON        eventListResponseEventWorkspaceStatusPropertiesJSON
}

type eventListResponseEventWorkspaceStatusPropertiesJSON struct {
	WorkspaceID apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceStatusProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceStatusPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceStatusType string

const (
	EventListResponseEventWorkspaceStatusTypeWorkspaceStatus EventListResponseEventWorkspaceStatusType = "workspace.status"
)

func (r EventListResponseEventWorkspaceStatusType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceStatusTypeWorkspaceStatus:
		return true
	}
	return false
}

type EventListResponseEventWorkspaceStatusStatus string

const (
	EventListResponseEventWorkspaceStatusStatusConnected    EventListResponseEventWorkspaceStatusStatus = "connected"
	EventListResponseEventWorkspaceStatusStatusConnecting   EventListResponseEventWorkspaceStatusStatus = "connecting"
	EventListResponseEventWorkspaceStatusStatusDisconnected EventListResponseEventWorkspaceStatusStatus = "disconnected"
	EventListResponseEventWorkspaceStatusStatusError        EventListResponseEventWorkspaceStatusStatus = "error"
)

func (r EventListResponseEventWorkspaceStatusStatus) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceStatusStatusConnected, EventListResponseEventWorkspaceStatusStatusConnecting, EventListResponseEventWorkspaceStatusStatusDisconnected, EventListResponseEventWorkspaceStatusStatusError:
		return true
	}
	return false
}

// =============================================================================
// EventListResponseEventSessionNextAgentSwitched
// =============================================================================

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
	Timestamp float64                                                      `json:"timestamp,required"`
	SessionID string                                                       `json:"sessionID,required"`
	Agent     string                                                       `json:"agent,required"`
	JSON      eventListResponseEventSessionNextAgentSwitchedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextAgentSwitchedPropertiesJSON struct {
	Timestamp   apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextModelSwitched
// =============================================================================

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
	Timestamp float64                                                      `json:"timestamp,required"`
	SessionID string                                                       `json:"sessionID,required"`
	Model     EventListResponseEventSessionNextModelSwitchedModel          `json:"model,required"`
	JSON      eventListResponseEventSessionNextModelSwitchedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextModelSwitchedPropertiesJSON struct {
	Timestamp   apijson.Field
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

type EventListResponseEventSessionNextModelSwitchedModel struct {
	ID         string                                                 `json:"id,required"`
	ProviderID string                                                 `json:"providerID,required"`
	Variant    string                                                 `json:"variant,required"`
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

// =============================================================================
// EventListResponseEventSessionNextPrompted
// =============================================================================

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
	Timestamp float64                                                 `json:"timestamp,required"`
	SessionID string                                                  `json:"sessionID,required"`
	Prompt    EventListResponseEventSessionNextPromptedPrompt         `json:"prompt,required"`
	JSON      eventListResponseEventSessionNextPromptedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextPromptedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextPromptedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextPromptedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextPromptedPrompt struct {
	Text       string                                                      `json:"text,required"`
	Files      []V2PromptFileAttachment                                    `json:"files"`
	Agents     []V2PromptAgentAttachment                                   `json:"agents"`
	References []V2PromptReferenceAttachment                               `json:"references"`
	JSON       eventListResponseEventSessionNextPromptedPromptJSON         `json:"-"`
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

// =============================================================================
// EventListResponseEventSessionNextSynthetic
// =============================================================================

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
	Timestamp float64                                                  `json:"timestamp,required"`
	SessionID string                                                   `json:"sessionID,required"`
	Text      string                                                   `json:"text,required"`
	JSON      eventListResponseEventSessionNextSyntheticPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextSyntheticPropertiesJSON struct {
	Timestamp   apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextShellStarted
// =============================================================================

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
	Timestamp float64                                                  `json:"timestamp,required"`
	SessionID string                                                   `json:"sessionID,required"`
	CallID    string                                                   `json:"callID,required"`
	Command   string                                                   `json:"command,required"`
	JSON      eventListResponseEventSessionNextShellStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextShellStartedPropertiesJSON struct {
	Timestamp   apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextShellEnded
// =============================================================================

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
	Timestamp float64                                                `json:"timestamp,required"`
	SessionID string                                                 `json:"sessionID,required"`
	CallID    string                                                 `json:"callID,required"`
	Output    string                                                 `json:"output,required"`
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

// =============================================================================
// EventListResponseEventSessionNextStepStarted
// =============================================================================

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
	Timestamp float64                                                  `json:"timestamp,required"`
	SessionID string                                                   `json:"sessionID,required"`
	Agent     string                                                   `json:"agent,required"`
	Model     EventListResponseEventSessionNextModelSwitchedModel      `json:"model,required"`
	Snapshot  string                                                   `json:"snapshot"`
	JSON      eventListResponseEventSessionNextStepStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextStepStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
	Snapshot    apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextStepEnded
// =============================================================================

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
	Timestamp float64                                                `json:"timestamp,required"`
	SessionID string                                                 `json:"sessionID,required"`
	Finish    string                                                 `json:"finish,required"`
	Cost      float64                                                `json:"cost,required"`
	Tokens    EventListResponseEventSessionNextStepEndedTokens        `json:"tokens,required"`
	Snapshot  string                                                 `json:"snapshot"`
	JSON      eventListResponseEventSessionNextStepEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextStepEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Finish      apijson.Field
	Cost        apijson.Field
	Tokens      apijson.Field
	Snapshot    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepEndedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepEndedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextStepEndedTokens struct {
	Input     float64                                              `json:"input,required"`
	Output    float64                                              `json:"output,required"`
	Reasoning float64                                              `json:"reasoning,required"`
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

type EventListResponseEventSessionNextStepEndedTokensCache struct {
	Read  float64                                                    `json:"read,required"`
	Write float64                                                    `json:"write,required"`
	JSON  eventListResponseEventSessionNextStepEndedTokensCacheJSON  `json:"-"`
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

// =============================================================================
// EventListResponseEventSessionNextStepFailed
// =============================================================================

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
	Timestamp float64                                                  `json:"timestamp,required"`
	SessionID string                                                   `json:"sessionID,required"`
	Error     EventListResponseEventSessionNextStepFailedError         `json:"error,required"`
	JSON      eventListResponseEventSessionNextStepFailedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextStepFailedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextStepFailedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextStepFailedPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventSessionNextStepFailedError struct {
	Type    string                                                       `json:"type,required"`
	Message string                                                       `json:"message,required"`
	JSON    eventListResponseEventSessionNextStepFailedErrorJSON         `json:"-"`
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

// =============================================================================
// EventListResponseEventSessionNextTextStarted
// =============================================================================

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
	Timestamp float64                                                  `json:"timestamp,required"`
	SessionID string                                                   `json:"sessionID,required"`
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

// =============================================================================
// EventListResponseEventSessionNextTextDelta
// =============================================================================

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
	Timestamp float64                                                `json:"timestamp,required"`
	SessionID string                                                 `json:"sessionID,required"`
	Delta     string                                                 `json:"delta,required"`
	JSON      eventListResponseEventSessionNextTextDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextTextDeltaPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Delta       apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextTextEnded
// =============================================================================

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
	Timestamp float64                                                `json:"timestamp,required"`
	SessionID string                                                 `json:"sessionID,required"`
	Text      string                                                 `json:"text,required"`
	JSON      eventListResponseEventSessionNextTextEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextTextEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextReasoningStarted
// =============================================================================

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
	Timestamp   float64                                                        `json:"timestamp,required"`
	SessionID   string                                                         `json:"sessionID,required"`
	ReasoningID string                                                         `json:"reasoningID,required"`
	JSON        eventListResponseEventSessionNextReasoningStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextReasoningStartedPropertiesJSON struct {
	Timestamp     apijson.Field
	SessionID     apijson.Field
	ReasoningID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextReasoningDelta
// =============================================================================

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
	Timestamp   float64                                                       `json:"timestamp,required"`
	SessionID   string                                                        `json:"sessionID,required"`
	ReasoningID string                                                        `json:"reasoningID,required"`
	Delta       string                                                        `json:"delta,required"`
	JSON        eventListResponseEventSessionNextReasoningDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextReasoningDeltaPropertiesJSON struct {
	Timestamp     apijson.Field
	SessionID     apijson.Field
	ReasoningID   apijson.Field
	Delta         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextReasoningEnded
// =============================================================================

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
	Timestamp   float64                                                       `json:"timestamp,required"`
	SessionID   string                                                        `json:"sessionID,required"`
	ReasoningID string                                                        `json:"reasoningID,required"`
	Text        string                                                        `json:"text,required"`
	JSON        eventListResponseEventSessionNextReasoningEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextReasoningEndedPropertiesJSON struct {
	Timestamp     apijson.Field
	SessionID     apijson.Field
	ReasoningID   apijson.Field
	Text          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextToolInputStarted
// =============================================================================

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
	Timestamp float64                                                          `json:"timestamp,required"`
	SessionID string                                                           `json:"sessionID,required"`
	CallID    string                                                           `json:"callID,required"`
	Name      string                                                           `json:"name,required"`
	JSON      eventListResponseEventSessionNextToolInputStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolInputStartedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Name        apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextToolInputDelta
// =============================================================================

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
	Timestamp float64                                                         `json:"timestamp,required"`
	SessionID string                                                          `json:"sessionID,required"`
	CallID    string                                                          `json:"callID,required"`
	Delta     string                                                          `json:"delta,required"`
	JSON      eventListResponseEventSessionNextToolInputDeltaPropertiesJSON   `json:"-"`
}

type eventListResponseEventSessionNextToolInputDeltaPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Delta       apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextToolInputEnded
// =============================================================================

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
	Timestamp float64                                                         `json:"timestamp,required"`
	SessionID string                                                          `json:"sessionID,required"`
	CallID    string                                                          `json:"callID,required"`
	Text      string                                                          `json:"text,required"`
	JSON      eventListResponseEventSessionNextToolInputEndedPropertiesJSON   `json:"-"`
}

type eventListResponseEventSessionNextToolInputEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Text        apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextToolCalled
// =============================================================================

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
	Timestamp float64                                                   `json:"timestamp,required"`
	SessionID string                                                    `json:"sessionID,required"`
	CallID    string                                                    `json:"callID,required"`
	Tool      string                                                    `json:"tool,required"`
	// This field can have the runtime type of map[string]interface{}.
	Input    interface{}                                               `json:"input,required"`
	Provider  EventListResponseEventSessionNextToolCalledProvider       `json:"provider,required"`
	JSON      eventListResponseEventSessionNextToolCalledPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolCalledPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Tool        apijson.Field
	Input       apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventSessionNextToolCalledProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventSessionNextToolCalledPropertiesJSON) RawJSON() string {
	return r.raw
}

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

// =============================================================================
// EventListResponseEventSessionNextToolProgress
// =============================================================================

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
	Timestamp  float64                                                    `json:"timestamp,required"`
	SessionID  string                                                     `json:"sessionID,required"`
	CallID     string                                                     `json:"callID,required"`
	// This field can have the runtime type of map[string]interface{}.
	Structured interface{} `json:"structured,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content []interface{} `json:"content,required"`
	JSON       eventListResponseEventSessionNextToolProgressPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolProgressPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Structured  apijson.Field
	Content     apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextToolSuccess
// =============================================================================

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
	Timestamp  float64                                                    `json:"timestamp,required"`
	SessionID  string                                                     `json:"sessionID,required"`
	CallID     string                                                     `json:"callID,required"`
	// This field can have the runtime type of map[string]interface{}.
	Structured interface{} `json:"structured,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content   []interface{}                                              `json:"content,required"`
	Provider   EventListResponseEventSessionNextToolCalledProvider        `json:"provider,required"`
	JSON       eventListResponseEventSessionNextToolSuccessPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolSuccessPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Structured  apijson.Field
	Content     apijson.Field
	Provider    apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextToolFailed
// =============================================================================

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
	Timestamp float64                                                  `json:"timestamp,required"`
	SessionID string                                                   `json:"sessionID,required"`
	CallID    string                                                   `json:"callID,required"`
	Error     EventListResponseEventSessionNextStepFailedError         `json:"error,required"`
	Provider  EventListResponseEventSessionNextToolCalledProvider      `json:"provider,required"`
	JSON      eventListResponseEventSessionNextToolFailedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextToolFailedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Error       apijson.Field
	Provider    apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextRetried
// =============================================================================

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
	Timestamp float64                                               `json:"timestamp,required"`
	SessionID string                                                `json:"sessionID,required"`
	Attempt   float64                                               `json:"attempt,required"`
	Error     EventListResponseEventSessionNextRetriedError         `json:"error,required"`
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

// =============================================================================
// EventListResponseEventSessionNextCompactionStarted
// =============================================================================

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
	Timestamp float64                                                          `json:"timestamp,required"`
	SessionID string                                                           `json:"sessionID,required"`
	Reason    string                                                           `json:"reason,required"`
	JSON      eventListResponseEventSessionNextCompactionStartedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextCompactionStartedPropertiesJSON struct {
	Timestamp   apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextCompactionDelta
// =============================================================================

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
	Timestamp float64                                                         `json:"timestamp,required"`
	SessionID string                                                          `json:"sessionID,required"`
	Text      string                                                          `json:"text,required"`
	JSON      eventListResponseEventSessionNextCompactionDeltaPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextCompactionDeltaPropertiesJSON struct {
	Timestamp   apijson.Field
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

// =============================================================================
// EventListResponseEventSessionNextCompactionEnded
// =============================================================================

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
	Timestamp float64                                                         `json:"timestamp,required"`
	SessionID string                                                          `json:"sessionID,required"`
	Text      string                                                          `json:"text,required"`
	Include   string                                                          `json:"include"`
	JSON      eventListResponseEventSessionNextCompactionEndedPropertiesJSON `json:"-"`
}

type eventListResponseEventSessionNextCompactionEndedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	Include     apijson.Field
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
