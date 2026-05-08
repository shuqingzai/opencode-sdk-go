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
	After     string `json:"after,required"`
	Before    string `json:"before,required"`
	Deletions int    `json:"deletions,required"`
	File      string `json:"file,required"`
	Status    string `json:"status,omitempty"`
	JSON      eventListResponseEventSessionDiffPropertiesDiffJSON
}

type eventListResponseEventSessionDiffPropertiesDiffJSON struct {
	Additions   apijson.Field
	After       apijson.Field
	Before      apijson.Field
	Deletions   apijson.Field
	File        apijson.Field
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

type EventListResponseEventVcsBranchUpdatedProperties struct {
	Branch string `json:"branch,required"`
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

type EventListResponseEventWorktreeReadyProperties struct {
	Branch string `json:"branch,required"`
	Name   string `json:"name,required"`
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

type EventListResponseEventWorkspaceRestore struct {
	Properties EventListResponseEventWorkspaceRestoreProperties `json:"properties,required"`
	Type       EventListResponseEventWorkspaceRestoreType       `json:"type,required"`
	JSON       eventListResponseEventWorkspaceRestoreJSON       `json:"-"`
}

type eventListResponseEventWorkspaceRestoreJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceRestore) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceRestoreJSON) RawJSON() string {
	return r.raw
}

func (r EventListResponseEventWorkspaceRestore) implementsEventListResponse() {}

type EventListResponseEventWorkspaceRestoreProperties struct {
	WorkspaceID string `json:"workspaceID,required"`
	SessionID   string `json:"sessionID,required"`
	Total       int64  `json:"total,required"`
	Step        int64  `json:"step,required"`
	JSON        eventListResponseEventWorkspaceRestorePropertiesJSON
}

type eventListResponseEventWorkspaceRestorePropertiesJSON struct {
	WorkspaceID apijson.Field
	SessionID   apijson.Field
	Total       apijson.Field
	Step        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventListResponseEventWorkspaceRestoreProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventListResponseEventWorkspaceRestorePropertiesJSON) RawJSON() string {
	return r.raw
}

type EventListResponseEventWorkspaceRestoreType string

const (
	EventListResponseEventWorkspaceRestoreTypeWorkspaceRestore EventListResponseEventWorkspaceRestoreType = "workspace.restore"
)

func (r EventListResponseEventWorkspaceRestoreType) IsKnown() bool {
	switch r {
	case EventListResponseEventWorkspaceRestoreTypeWorkspaceRestore:
		return true
	}
	return false
}

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
