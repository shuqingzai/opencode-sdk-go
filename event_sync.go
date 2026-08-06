// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"github.com/sst/opencode-sdk-go/internal/apijson"
)

type SyncEventMessagePartRemovedType string

const (
	SyncEventMessagePartRemovedTypeMessagePartRemoved1 SyncEventMessagePartRemovedType = "message.part.removed.1"
)

func (r SyncEventMessagePartRemovedType) IsKnown() bool {
	switch r {
	case SyncEventMessagePartRemovedTypeMessagePartRemoved1:
		return true
	}
	return false
}

type SyncEventMessagePartRemoved struct {
	Type        SyncEventMessagePartRemovedType                    `json:"type,required"`
	ID          string                                             `json:"id,required"`
	Seq         int64                                              `json:"seq,required"`
	AggregateID string                                             `json:"aggregateID,required"`
	Data        EventListResponseEventMessagePartRemovedProperties `json:"data,required"`
	JSON        syncEventMessagePartRemovedJSON                    `json:"-"`
}

type syncEventMessagePartRemovedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventMessagePartRemovedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventMessagePartRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventMessagePartRemoved) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventMessagePartUpdatedType string

const (
	SyncEventMessagePartUpdatedTypeMessagePartUpdated1 SyncEventMessagePartUpdatedType = "message.part.updated.1"
)

func (r SyncEventMessagePartUpdatedType) IsKnown() bool {
	switch r {
	case SyncEventMessagePartUpdatedTypeMessagePartUpdated1:
		return true
	}
	return false
}

type SyncEventMessagePartUpdated struct {
	Type        SyncEventMessagePartUpdatedType                    `json:"type,required"`
	ID          string                                             `json:"id,required"`
	Seq         int64                                              `json:"seq,required"`
	AggregateID string                                             `json:"aggregateID,required"`
	Data        EventListResponseEventMessagePartUpdatedProperties `json:"data,required"`
	JSON        syncEventMessagePartUpdatedJSON                    `json:"-"`
}

type syncEventMessagePartUpdatedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventMessagePartUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventMessagePartUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventMessagePartUpdated) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventMessageRemovedType string

const (
	SyncEventMessageRemovedTypeMessageRemoved1 SyncEventMessageRemovedType = "message.removed.1"
)

func (r SyncEventMessageRemovedType) IsKnown() bool {
	switch r {
	case SyncEventMessageRemovedTypeMessageRemoved1:
		return true
	}
	return false
}

type SyncEventMessageRemoved struct {
	Type        SyncEventMessageRemovedType                    `json:"type,required"`
	ID          string                                         `json:"id,required"`
	Seq         int64                                          `json:"seq,required"`
	AggregateID string                                         `json:"aggregateID,required"`
	Data        EventListResponseEventMessageRemovedProperties `json:"data,required"`
	JSON        syncEventMessageRemovedJSON                    `json:"-"`
}

type syncEventMessageRemovedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventMessageRemovedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventMessageRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventMessageRemoved) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventMessageUpdatedType string

const (
	SyncEventMessageUpdatedTypeMessageUpdated1 SyncEventMessageUpdatedType = "message.updated.1"
)

func (r SyncEventMessageUpdatedType) IsKnown() bool {
	switch r {
	case SyncEventMessageUpdatedTypeMessageUpdated1:
		return true
	}
	return false
}

type SyncEventMessageUpdated struct {
	Type        SyncEventMessageUpdatedType                    `json:"type,required"`
	ID          string                                         `json:"id,required"`
	Seq         int64                                          `json:"seq,required"`
	AggregateID string                                         `json:"aggregateID,required"`
	Data        EventListResponseEventMessageUpdatedProperties `json:"data,required"`
	JSON        syncEventMessageUpdatedJSON                    `json:"-"`
}

type syncEventMessageUpdatedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventMessageUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventMessageUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventMessageUpdated) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionCreatedType string

const (
	SyncEventSessionCreatedTypeSessionCreated1 SyncEventSessionCreatedType = "session.created.1"
)

func (r SyncEventSessionCreatedType) IsKnown() bool {
	switch r {
	case SyncEventSessionCreatedTypeSessionCreated1:
		return true
	}
	return false
}

type SyncEventSessionCreated struct {
	Type        SyncEventSessionCreatedType                    `json:"type,required"`
	ID          string                                         `json:"id,required"`
	Seq         int64                                          `json:"seq,required"`
	AggregateID string                                         `json:"aggregateID,required"`
	Data        EventListResponseEventSessionCreatedProperties `json:"data,required"`
	JSON        syncEventSessionCreatedJSON                    `json:"-"`
}

type syncEventSessionCreatedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionCreatedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionCreated) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionDeletedType string

const (
	SyncEventSessionDeletedTypeSessionDeleted1 SyncEventSessionDeletedType = "session.deleted.1"
)

func (r SyncEventSessionDeletedType) IsKnown() bool {
	switch r {
	case SyncEventSessionDeletedTypeSessionDeleted1:
		return true
	}
	return false
}

type SyncEventSessionDeleted struct {
	Type        SyncEventSessionDeletedType                    `json:"type,required"`
	ID          string                                         `json:"id,required"`
	Seq         int64                                          `json:"seq,required"`
	AggregateID string                                         `json:"aggregateID,required"`
	Data        EventListResponseEventSessionDeletedProperties `json:"data,required"`
	JSON        syncEventSessionDeletedJSON                    `json:"-"`
}

type syncEventSessionDeletedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionDeletedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionDeleted) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextAgentSwitchedType string

const (
	SyncEventSessionNextAgentSwitchedTypeSessionNextAgentSwitched1 SyncEventSessionNextAgentSwitchedType = "session.next.agent.switched.1"
)

func (r SyncEventSessionNextAgentSwitchedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextAgentSwitchedTypeSessionNextAgentSwitched1:
		return true
	}
	return false
}

type SyncEventSessionNextAgentSwitched struct {
	Type        SyncEventSessionNextAgentSwitchedType                    `json:"type,required"`
	ID          string                                                   `json:"id,required"`
	Seq         int64                                                    `json:"seq,required"`
	AggregateID string                                                   `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextAgentSwitchedProperties `json:"data,required"`
	JSON        syncEventSessionNextAgentSwitchedJSON                    `json:"-"`
}

type syncEventSessionNextAgentSwitchedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextAgentSwitchedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextAgentSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextAgentSwitched) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextCompactionEndedType string

const (
	SyncEventSessionNextCompactionEndedTypeSessionNextCompactionEnded1 SyncEventSessionNextCompactionEndedType = "session.next.compaction.ended.1"
)

func (r SyncEventSessionNextCompactionEndedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextCompactionEndedTypeSessionNextCompactionEnded1:
		return true
	}
	return false
}

type SyncEventSessionNextCompactionEnded struct {
	Type        SyncEventSessionNextCompactionEndedType                    `json:"type,required"`
	ID          string                                                     `json:"id,required"`
	Seq         int64                                                      `json:"seq,required"`
	AggregateID string                                                     `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextCompactionEndedProperties `json:"data,required"`
	JSON        syncEventSessionNextCompactionEndedJSON                    `json:"-"`
}

type syncEventSessionNextCompactionEndedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextCompactionEndedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextCompactionEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextCompactionEnded) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextCompactionStartedType string

const (
	SyncEventSessionNextCompactionStartedTypeSessionNextCompactionStarted1 SyncEventSessionNextCompactionStartedType = "session.next.compaction.started.1"
)

func (r SyncEventSessionNextCompactionStartedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextCompactionStartedTypeSessionNextCompactionStarted1:
		return true
	}
	return false
}

type SyncEventSessionNextCompactionStarted struct {
	Type        SyncEventSessionNextCompactionStartedType                    `json:"type,required"`
	ID          string                                                       `json:"id,required"`
	Seq         int64                                                        `json:"seq,required"`
	AggregateID string                                                       `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextCompactionStartedProperties `json:"data,required"`
	JSON        syncEventSessionNextCompactionStartedJSON                    `json:"-"`
}

type syncEventSessionNextCompactionStartedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextCompactionStartedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextCompactionStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextCompactionStarted) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextModelSwitchedType string

const (
	SyncEventSessionNextModelSwitchedTypeSessionNextModelSwitched1 SyncEventSessionNextModelSwitchedType = "session.next.model.switched.1"
)

func (r SyncEventSessionNextModelSwitchedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextModelSwitchedTypeSessionNextModelSwitched1:
		return true
	}
	return false
}

type SyncEventSessionNextModelSwitched struct {
	Type        SyncEventSessionNextModelSwitchedType                    `json:"type,required"`
	ID          string                                                   `json:"id,required"`
	Seq         int64                                                    `json:"seq,required"`
	AggregateID string                                                   `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextModelSwitchedProperties `json:"data,required"`
	JSON        syncEventSessionNextModelSwitchedJSON                    `json:"-"`
}

type syncEventSessionNextModelSwitchedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextModelSwitchedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextModelSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextModelSwitched) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextPromptedType string

const (
	SyncEventSessionNextPromptedTypeSessionNextPrompted1 SyncEventSessionNextPromptedType = "session.next.prompted.1"
)

func (r SyncEventSessionNextPromptedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextPromptedTypeSessionNextPrompted1:
		return true
	}
	return false
}

type SyncEventSessionNextPrompted struct {
	Type        SyncEventSessionNextPromptedType                    `json:"type,required"`
	ID          string                                              `json:"id,required"`
	Seq         int64                                               `json:"seq,required"`
	AggregateID string                                              `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextPromptedProperties `json:"data,required"`
	JSON        syncEventSessionNextPromptedJSON                    `json:"-"`
}

type syncEventSessionNextPromptedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextPromptedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextPrompted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextPrompted) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextReasoningEndedType string

const (
	SyncEventSessionNextReasoningEndedTypeSessionNextReasoningEnded1 SyncEventSessionNextReasoningEndedType = "session.next.reasoning.ended.1"
)

func (r SyncEventSessionNextReasoningEndedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextReasoningEndedTypeSessionNextReasoningEnded1:
		return true
	}
	return false
}

type SyncEventSessionNextReasoningEnded struct {
	Type        SyncEventSessionNextReasoningEndedType                    `json:"type,required"`
	ID          string                                                    `json:"id,required"`
	Seq         int64                                                     `json:"seq,required"`
	AggregateID string                                                    `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextReasoningEndedProperties `json:"data,required"`
	JSON        syncEventSessionNextReasoningEndedJSON                    `json:"-"`
}

type syncEventSessionNextReasoningEndedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextReasoningEndedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextReasoningEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextReasoningEnded) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextReasoningStartedType string

const (
	SyncEventSessionNextReasoningStartedTypeSessionNextReasoningStarted1 SyncEventSessionNextReasoningStartedType = "session.next.reasoning.started.1"
)

func (r SyncEventSessionNextReasoningStartedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextReasoningStartedTypeSessionNextReasoningStarted1:
		return true
	}
	return false
}

type SyncEventSessionNextReasoningStarted struct {
	Type        SyncEventSessionNextReasoningStartedType                    `json:"type,required"`
	ID          string                                                      `json:"id,required"`
	Seq         int64                                                       `json:"seq,required"`
	AggregateID string                                                      `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextReasoningStartedProperties `json:"data,required"`
	JSON        syncEventSessionNextReasoningStartedJSON                    `json:"-"`
}

type syncEventSessionNextReasoningStartedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextReasoningStartedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextReasoningStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextReasoningStarted) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextRetriedType string

const (
	SyncEventSessionNextRetriedTypeSessionNextRetried1 SyncEventSessionNextRetriedType = "session.next.retried.1"
)

func (r SyncEventSessionNextRetriedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextRetriedTypeSessionNextRetried1:
		return true
	}
	return false
}

type SyncEventSessionNextRetried struct {
	Type        SyncEventSessionNextRetriedType                    `json:"type,required"`
	ID          string                                             `json:"id,required"`
	Seq         int64                                              `json:"seq,required"`
	AggregateID string                                             `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextRetriedProperties `json:"data,required"`
	JSON        syncEventSessionNextRetriedJSON                    `json:"-"`
}

type syncEventSessionNextRetriedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextRetriedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextRetried) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextRetried) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextShellEndedType string

const (
	SyncEventSessionNextShellEndedTypeSessionNextShellEnded1 SyncEventSessionNextShellEndedType = "session.next.shell.ended.1"
)

func (r SyncEventSessionNextShellEndedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextShellEndedTypeSessionNextShellEnded1:
		return true
	}
	return false
}

type SyncEventSessionNextShellEnded struct {
	Type        SyncEventSessionNextShellEndedType                    `json:"type,required"`
	ID          string                                                `json:"id,required"`
	Seq         int64                                                 `json:"seq,required"`
	AggregateID string                                                `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextShellEndedProperties `json:"data,required"`
	JSON        syncEventSessionNextShellEndedJSON                    `json:"-"`
}

type syncEventSessionNextShellEndedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextShellEndedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextShellEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextShellEnded) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextShellStartedType string

const (
	SyncEventSessionNextShellStartedTypeSessionNextShellStarted1 SyncEventSessionNextShellStartedType = "session.next.shell.started.1"
)

func (r SyncEventSessionNextShellStartedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextShellStartedTypeSessionNextShellStarted1:
		return true
	}
	return false
}

type SyncEventSessionNextShellStarted struct {
	Type        SyncEventSessionNextShellStartedType                    `json:"type,required"`
	ID          string                                                  `json:"id,required"`
	Seq         int64                                                   `json:"seq,required"`
	AggregateID string                                                  `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextShellStartedProperties `json:"data,required"`
	JSON        syncEventSessionNextShellStartedJSON                    `json:"-"`
}

type syncEventSessionNextShellStartedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextShellStartedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextShellStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextShellStarted) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextStepEndedType string

const (
	SyncEventSessionNextStepEndedTypeSessionNextStepEnded2 SyncEventSessionNextStepEndedType = "session.next.step.ended.2"
)

func (r SyncEventSessionNextStepEndedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextStepEndedTypeSessionNextStepEnded2:
		return true
	}
	return false
}

type SyncEventSessionNextStepEnded struct {
	Type        SyncEventSessionNextStepEndedType                    `json:"type,required"`
	ID          string                                               `json:"id,required"`
	Seq         int64                                                `json:"seq,required"`
	AggregateID string                                               `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextStepEndedProperties `json:"data,required"`
	JSON        syncEventSessionNextStepEndedJSON                    `json:"-"`
}

type syncEventSessionNextStepEndedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextStepEndedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextStepEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextStepEnded) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextStepFailedType string

const (
	SyncEventSessionNextStepFailedTypeSessionNextStepFailed2 SyncEventSessionNextStepFailedType = "session.next.step.failed.2"
)

func (r SyncEventSessionNextStepFailedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextStepFailedTypeSessionNextStepFailed2:
		return true
	}
	return false
}

type SyncEventSessionNextStepFailed struct {
	Type        SyncEventSessionNextStepFailedType                    `json:"type,required"`
	ID          string                                                `json:"id,required"`
	Seq         int64                                                 `json:"seq,required"`
	AggregateID string                                                `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextStepFailedProperties `json:"data,required"`
	JSON        syncEventSessionNextStepFailedJSON                    `json:"-"`
}

type syncEventSessionNextStepFailedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextStepFailedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextStepFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextStepFailed) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextStepStartedType string

const (
	SyncEventSessionNextStepStartedTypeSessionNextStepStarted1 SyncEventSessionNextStepStartedType = "session.next.step.started.1"
)

func (r SyncEventSessionNextStepStartedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextStepStartedTypeSessionNextStepStarted1:
		return true
	}
	return false
}

type SyncEventSessionNextStepStarted struct {
	Type        SyncEventSessionNextStepStartedType                    `json:"type,required"`
	ID          string                                                 `json:"id,required"`
	Seq         int64                                                  `json:"seq,required"`
	AggregateID string                                                 `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextStepStartedProperties `json:"data,required"`
	JSON        syncEventSessionNextStepStartedJSON                    `json:"-"`
}

type syncEventSessionNextStepStartedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextStepStartedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextStepStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextStepStarted) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextSyntheticType string

const (
	SyncEventSessionNextSyntheticTypeSessionNextSynthetic1 SyncEventSessionNextSyntheticType = "session.next.synthetic.1"
)

func (r SyncEventSessionNextSyntheticType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextSyntheticTypeSessionNextSynthetic1:
		return true
	}
	return false
}

type SyncEventSessionNextSynthetic struct {
	Type        SyncEventSessionNextSyntheticType                    `json:"type,required"`
	ID          string                                               `json:"id,required"`
	Seq         int64                                                `json:"seq,required"`
	AggregateID string                                               `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextSyntheticProperties `json:"data,required"`
	JSON        syncEventSessionNextSyntheticJSON                    `json:"-"`
}

type syncEventSessionNextSyntheticJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextSyntheticJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextSynthetic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextSynthetic) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextTextEndedType string

const (
	SyncEventSessionNextTextEndedTypeSessionNextTextEnded1 SyncEventSessionNextTextEndedType = "session.next.text.ended.1"
)

func (r SyncEventSessionNextTextEndedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextTextEndedTypeSessionNextTextEnded1:
		return true
	}
	return false
}

type SyncEventSessionNextTextEnded struct {
	Type        SyncEventSessionNextTextEndedType                    `json:"type,required"`
	ID          string                                               `json:"id,required"`
	Seq         int64                                                `json:"seq,required"`
	AggregateID string                                               `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextTextEndedProperties `json:"data,required"`
	JSON        syncEventSessionNextTextEndedJSON                    `json:"-"`
}

type syncEventSessionNextTextEndedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextTextEndedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextTextEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextTextEnded) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextTextStartedType string

const (
	SyncEventSessionNextTextStartedTypeSessionNextTextStarted1 SyncEventSessionNextTextStartedType = "session.next.text.started.1"
)

func (r SyncEventSessionNextTextStartedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextTextStartedTypeSessionNextTextStarted1:
		return true
	}
	return false
}

type SyncEventSessionNextTextStarted struct {
	Type        SyncEventSessionNextTextStartedType                    `json:"type,required"`
	ID          string                                                 `json:"id,required"`
	Seq         int64                                                  `json:"seq,required"`
	AggregateID string                                                 `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextTextStartedProperties `json:"data,required"`
	JSON        syncEventSessionNextTextStartedJSON                    `json:"-"`
}

type syncEventSessionNextTextStartedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextTextStartedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextTextStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextTextStarted) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextToolCalledType string

const (
	SyncEventSessionNextToolCalledTypeSessionNextToolCalled1 SyncEventSessionNextToolCalledType = "session.next.tool.called.1"
)

func (r SyncEventSessionNextToolCalledType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextToolCalledTypeSessionNextToolCalled1:
		return true
	}
	return false
}

type SyncEventSessionNextToolCalled struct {
	Type        SyncEventSessionNextToolCalledType                    `json:"type,required"`
	ID          string                                                `json:"id,required"`
	Seq         int64                                                 `json:"seq,required"`
	AggregateID string                                                `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextToolCalledProperties `json:"data,required"`
	JSON        syncEventSessionNextToolCalledJSON                    `json:"-"`
}

type syncEventSessionNextToolCalledJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextToolCalledJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextToolCalled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextToolCalled) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextToolFailedType string

const (
	SyncEventSessionNextToolFailedTypeSessionNextToolFailed1 SyncEventSessionNextToolFailedType = "session.next.tool.failed.1"
)

func (r SyncEventSessionNextToolFailedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextToolFailedTypeSessionNextToolFailed1:
		return true
	}
	return false
}

type SyncEventSessionNextToolFailed struct {
	Type        SyncEventSessionNextToolFailedType                    `json:"type,required"`
	ID          string                                                `json:"id,required"`
	Seq         int64                                                 `json:"seq,required"`
	AggregateID string                                                `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextToolFailedProperties `json:"data,required"`
	JSON        syncEventSessionNextToolFailedJSON                    `json:"-"`
}

type syncEventSessionNextToolFailedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextToolFailedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextToolFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextToolFailed) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextToolInputEndedType string

const (
	SyncEventSessionNextToolInputEndedTypeSessionNextToolInputEnded1 SyncEventSessionNextToolInputEndedType = "session.next.tool.input.ended.1"
)

func (r SyncEventSessionNextToolInputEndedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextToolInputEndedTypeSessionNextToolInputEnded1:
		return true
	}
	return false
}

type SyncEventSessionNextToolInputEnded struct {
	Type        SyncEventSessionNextToolInputEndedType                    `json:"type,required"`
	ID          string                                                    `json:"id,required"`
	Seq         int64                                                     `json:"seq,required"`
	AggregateID string                                                    `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextToolInputEndedProperties `json:"data,required"`
	JSON        syncEventSessionNextToolInputEndedJSON                    `json:"-"`
}

type syncEventSessionNextToolInputEndedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextToolInputEndedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextToolInputEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextToolInputEnded) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextToolInputStartedType string

const (
	SyncEventSessionNextToolInputStartedTypeSessionNextToolInputStarted1 SyncEventSessionNextToolInputStartedType = "session.next.tool.input.started.1"
)

func (r SyncEventSessionNextToolInputStartedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextToolInputStartedTypeSessionNextToolInputStarted1:
		return true
	}
	return false
}

type SyncEventSessionNextToolInputStarted struct {
	Type        SyncEventSessionNextToolInputStartedType                    `json:"type,required"`
	ID          string                                                      `json:"id,required"`
	Seq         int64                                                       `json:"seq,required"`
	AggregateID string                                                      `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextToolInputStartedProperties `json:"data,required"`
	JSON        syncEventSessionNextToolInputStartedJSON                    `json:"-"`
}

type syncEventSessionNextToolInputStartedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextToolInputStartedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextToolInputStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextToolInputStarted) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextToolProgressType string

const (
	SyncEventSessionNextToolProgressTypeSessionNextToolProgress1 SyncEventSessionNextToolProgressType = "session.next.tool.progress.1"
)

func (r SyncEventSessionNextToolProgressType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextToolProgressTypeSessionNextToolProgress1:
		return true
	}
	return false
}

type SyncEventSessionNextToolProgress struct {
	Type        SyncEventSessionNextToolProgressType                    `json:"type,required"`
	ID          string                                                  `json:"id,required"`
	Seq         int64                                                   `json:"seq,required"`
	AggregateID string                                                  `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextToolProgressProperties `json:"data,required"`
	JSON        syncEventSessionNextToolProgressJSON                    `json:"-"`
}

type syncEventSessionNextToolProgressJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextToolProgressJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextToolProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextToolProgress) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextToolSuccessType string

const (
	SyncEventSessionNextToolSuccessTypeSessionNextToolSuccess1 SyncEventSessionNextToolSuccessType = "session.next.tool.success.1"
)

func (r SyncEventSessionNextToolSuccessType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextToolSuccessTypeSessionNextToolSuccess1:
		return true
	}
	return false
}

type SyncEventSessionNextToolSuccess struct {
	Type        SyncEventSessionNextToolSuccessType                    `json:"type,required"`
	ID          string                                                 `json:"id,required"`
	Seq         int64                                                  `json:"seq,required"`
	AggregateID string                                                 `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextToolSuccessProperties `json:"data,required"`
	JSON        syncEventSessionNextToolSuccessJSON                    `json:"-"`
}

type syncEventSessionNextToolSuccessJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextToolSuccessJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextToolSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextToolSuccess) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionUpdatedType string

const (
	SyncEventSessionUpdatedTypeSessionUpdated1 SyncEventSessionUpdatedType = "session.updated.1"
)

func (r SyncEventSessionUpdatedType) IsKnown() bool {
	switch r {
	case SyncEventSessionUpdatedTypeSessionUpdated1:
		return true
	}
	return false
}

type SyncEventSessionUpdated struct {
	Type        SyncEventSessionUpdatedType                    `json:"type,required"`
	ID          string                                         `json:"id,required"`
	Seq         int64                                          `json:"seq,required"`
	AggregateID string                                         `json:"aggregateID,required"`
	Data        EventListResponseEventSessionUpdatedProperties `json:"data,required"`
	JSON        syncEventSessionUpdatedJSON                    `json:"-"`
}

type syncEventSessionUpdatedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionUpdated) implementsSyncEventResponseSyncEventDataUnion() {}

// =============================================================================
// SyncEventSessionNextMoved
// =============================================================================

type SyncEventSessionNextMovedType string

const (
	SyncEventSessionNextMovedTypeSessionNextMoved1 SyncEventSessionNextMovedType = "session.next.moved.1"
)

func (r SyncEventSessionNextMovedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextMovedTypeSessionNextMoved1:
		return true
	}
	return false
}

type SyncEventSessionNextMoved struct {
	Type        SyncEventSessionNextMovedType                    `json:"type,required"`
	ID          string                                           `json:"id,required"`
	Seq         int64                                            `json:"seq,required"`
	AggregateID string                                           `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextMovedProperties `json:"data,required"`
	JSON        syncEventSessionNextMovedJSON                    `json:"-"`
}

type syncEventSessionNextMovedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextMovedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextMoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextMoved) implementsSyncEventResponseSyncEventDataUnion() {}

// =============================================================================
// SyncEventSessionNextRevertStaged
// =============================================================================

type SyncEventSessionNextRevertStagedType string

const (
	SyncEventSessionNextRevertStagedTypeSessionNextRevertStaged1 SyncEventSessionNextRevertStagedType = "session.next.revert.staged.1"
)

func (r SyncEventSessionNextRevertStagedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextRevertStagedTypeSessionNextRevertStaged1:
		return true
	}
	return false
}

type SyncEventSessionNextRevertStaged struct {
	Type        SyncEventSessionNextRevertStagedType                    `json:"type,required"`
	ID          string                                                  `json:"id,required"`
	Seq         int64                                                   `json:"seq,required"`
	AggregateID string                                                  `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextRevertStagedProperties `json:"data,required"`
	JSON        syncEventSessionNextRevertStagedJSON                    `json:"-"`
}

type syncEventSessionNextRevertStagedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextRevertStagedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextRevertStaged) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextRevertStaged) implementsSyncEventResponseSyncEventDataUnion() {}

// =============================================================================
// SyncEventSessionNextRevertCleared
// =============================================================================

type SyncEventSessionNextRevertClearedType string

const (
	SyncEventSessionNextRevertClearedTypeSessionNextRevertCleared1 SyncEventSessionNextRevertClearedType = "session.next.revert.cleared.1"
)

func (r SyncEventSessionNextRevertClearedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextRevertClearedTypeSessionNextRevertCleared1:
		return true
	}
	return false
}

type SyncEventSessionNextRevertCleared struct {
	Type        SyncEventSessionNextRevertClearedType                    `json:"type,required"`
	ID          string                                                   `json:"id,required"`
	Seq         int64                                                    `json:"seq,required"`
	AggregateID string                                                   `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextRevertClearedProperties `json:"data,required"`
	JSON        syncEventSessionNextRevertClearedJSON                    `json:"-"`
}

type syncEventSessionNextRevertClearedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextRevertClearedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextRevertCleared) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextRevertCleared) implementsSyncEventResponseSyncEventDataUnion() {}

// =============================================================================
// SyncEventSessionNextRevertCommitted
// =============================================================================

type SyncEventSessionNextRevertCommittedType string

const (
	SyncEventSessionNextRevertCommittedTypeSessionNextRevertCommitted1 SyncEventSessionNextRevertCommittedType = "session.next.revert.committed.1"
)

func (r SyncEventSessionNextRevertCommittedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextRevertCommittedTypeSessionNextRevertCommitted1:
		return true
	}
	return false
}

type SyncEventSessionNextRevertCommitted struct {
	Type        SyncEventSessionNextRevertCommittedType                    `json:"type,required"`
	ID          string                                                     `json:"id,required"`
	Seq         int64                                                      `json:"seq,required"`
	AggregateID string                                                     `json:"aggregateID,required"`
	Data        EventListResponseEventSessionNextRevertCommittedProperties `json:"data,required"`
	JSON        syncEventSessionNextRevertCommittedJSON                    `json:"-"`
}

type syncEventSessionNextRevertCommittedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextRevertCommittedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextRevertCommitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextRevertCommitted) implementsSyncEventResponseSyncEventDataUnion() {}

// Deprecated: use [V2SessionInputPrompt] instead.
// SyncEventPrompt is a type alias for [V2SessionInputPrompt]. OpenAPI declares
// this as a $ref to Prompt.
type SyncEventPrompt = V2SessionInputPrompt

type SyncEventSessionNextContextUpdatedType string

const (
	SyncEventSessionNextContextUpdatedTypeSessionNextContextUpdated1 SyncEventSessionNextContextUpdatedType = "session.next.context.updated.1"
)

func (r SyncEventSessionNextContextUpdatedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextContextUpdatedTypeSessionNextContextUpdated1:
		return true
	}
	return false
}

type SyncEventSessionNextContextUpdatedProperties struct {
	Timestamp int64                                            `json:"timestamp,required"`
	SessionID string                                           `json:"sessionID,required"`
	MessageID string                                           `json:"messageID,required"`
	Text      string                                           `json:"text,required"`
	JSON      syncEventSessionNextContextUpdatedPropertiesJSON `json:"-"`
}

type syncEventSessionNextContextUpdatedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	MessageID   apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextContextUpdatedPropertiesJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextContextUpdatedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SyncEventSessionNextContextUpdated struct {
	Type        SyncEventSessionNextContextUpdatedType       `json:"type,required"`
	ID          string                                       `json:"id,required"`
	Seq         int64                                        `json:"seq,required"`
	AggregateID string                                       `json:"aggregateID,required"`
	Data        SyncEventSessionNextContextUpdatedProperties `json:"data,required"`
	JSON        syncEventSessionNextContextUpdatedJSON       `json:"-"`
}

type syncEventSessionNextContextUpdatedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextContextUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextContextUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextContextUpdated) implementsSyncEventResponseSyncEventDataUnion() {}

type SyncEventSessionNextPromptAdmittedType string

const (
	SyncEventSessionNextPromptAdmittedTypeSessionNextPromptAdmitted1 SyncEventSessionNextPromptAdmittedType = "session.next.prompt.admitted.1"
)

func (r SyncEventSessionNextPromptAdmittedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextPromptAdmittedTypeSessionNextPromptAdmitted1:
		return true
	}
	return false
}

// SyncEventSessionNextPromptAdmittedDelivery is the delivery mode for an admitted prompt.
type SyncEventSessionNextPromptAdmittedDelivery string

const (
	SyncEventSessionNextPromptAdmittedDeliverySteer SyncEventSessionNextPromptAdmittedDelivery = "steer"
	SyncEventSessionNextPromptAdmittedDeliveryQueue SyncEventSessionNextPromptAdmittedDelivery = "queue"
)

func (r SyncEventSessionNextPromptAdmittedDelivery) IsKnown() bool {
	switch r {
	case SyncEventSessionNextPromptAdmittedDeliverySteer, SyncEventSessionNextPromptAdmittedDeliveryQueue:
		return true
	}
	return false
}

type SyncEventSessionNextPromptAdmittedProperties struct {
	Timestamp int64                                            `json:"timestamp,required"`
	SessionID string                                           `json:"sessionID,required"`
	MessageID string                                           `json:"messageID,required"`
	Prompt    V2SessionInputPrompt                             `json:"prompt,required"`
	Delivery  SyncEventSessionNextPromptAdmittedDelivery       `json:"delivery,required"`
	JSON      syncEventSessionNextPromptAdmittedPropertiesJSON `json:"-"`
}

type syncEventSessionNextPromptAdmittedPropertiesJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	MessageID   apijson.Field
	Prompt      apijson.Field
	Delivery    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextPromptAdmittedPropertiesJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextPromptAdmittedProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SyncEventSessionNextPromptAdmitted struct {
	Type        SyncEventSessionNextPromptAdmittedType       `json:"type,required"`
	ID          string                                       `json:"id,required"`
	Seq         int64                                        `json:"seq,required"`
	AggregateID string                                       `json:"aggregateID,required"`
	Data        SyncEventSessionNextPromptAdmittedProperties `json:"data,required"`
	JSON        syncEventSessionNextPromptAdmittedJSON       `json:"-"`
}

type syncEventSessionNextPromptAdmittedJSON struct {
	Type        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r syncEventSessionNextPromptAdmittedJSON) RawJSON() string {
	return r.raw
}

func (r *SyncEventSessionNextPromptAdmitted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SyncEventSessionNextPromptAdmitted) implementsSyncEventResponseSyncEventDataUnion() {}
