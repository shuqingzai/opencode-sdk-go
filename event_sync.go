// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"github.com/sst/opencode-sdk-go/internal/apijson"
)

// SyncEventMessageUpdated is a V1 sync event with name "message.updated.1".
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
	Type        SyncEventMessageUpdatedType `json:"type,required"`
	Name        string                      `json:"name,required"`
	ID          string                      `json:"id,required"`
	Seq         int64                       `json:"seq,required"`
	AggregateID string                      `json:"aggregateID,required"`
	Data        SyncEventMessageUpdatedData `json:"data,required"`
	JSON        syncEventMessageUpdatedJSON `json:"-"`
}

type syncEventMessageUpdatedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessageUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessageUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventMessageUpdated) implementsGlobalEventPayload() {}

type SyncEventMessageUpdatedData struct {
	SessionID string                          `json:"sessionID,required"`
	Info      Message                         `json:"info,required"`
	JSON      syncEventMessageUpdatedDataJSON `json:"-"`
}

type syncEventMessageUpdatedDataJSON struct {
	SessionID   apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessageUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessageUpdatedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventMessageRemoved is a V1 sync event with name "message.removed.1".
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
	Type        SyncEventMessageRemovedType `json:"type,required"`
	Name        string                      `json:"name,required"`
	ID          string                      `json:"id,required"`
	Seq         int64                       `json:"seq,required"`
	AggregateID string                      `json:"aggregateID,required"`
	Data        SyncEventMessageRemovedData `json:"data,required"`
	JSON        syncEventMessageRemovedJSON `json:"-"`
}

type syncEventMessageRemovedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessageRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessageRemovedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventMessageRemoved) implementsGlobalEventPayload() {}

type SyncEventMessageRemovedData struct {
	SessionID string                          `json:"sessionID,required"`
	MessageID string                          `json:"messageID,required"`
	JSON      syncEventMessageRemovedDataJSON `json:"-"`
}

type syncEventMessageRemovedDataJSON struct {
	SessionID   apijson.Field
	MessageID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessageRemovedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessageRemovedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventMessagePartUpdated is a V1 sync event with name "message.part.updated.1".
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
	Type        SyncEventMessagePartUpdatedType `json:"type,required"`
	Name        string                          `json:"name,required"`
	ID          string                          `json:"id,required"`
	Seq         int64                           `json:"seq,required"`
	AggregateID string                          `json:"aggregateID,required"`
	Data        SyncEventMessagePartUpdatedData `json:"data,required"`
	JSON        syncEventMessagePartUpdatedJSON `json:"-"`
}

type syncEventMessagePartUpdatedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessagePartUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessagePartUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventMessagePartUpdated) implementsGlobalEventPayload() {}

type SyncEventMessagePartUpdatedData struct {
	SessionID string                              `json:"sessionID,required"`
	Part      Part                                `json:"part,required"`
	Time      int64                               `json:"time,required"`
	JSON      syncEventMessagePartUpdatedDataJSON `json:"-"`
}

type syncEventMessagePartUpdatedDataJSON struct {
	SessionID   apijson.Field
	Part        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessagePartUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessagePartUpdatedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventMessagePartRemoved is a V1 sync event with name "message.part.removed.1".
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
	Type        SyncEventMessagePartRemovedType `json:"type,required"`
	Name        string                          `json:"name,required"`
	ID          string                          `json:"id,required"`
	Seq         int64                           `json:"seq,required"`
	AggregateID string                          `json:"aggregateID,required"`
	Data        SyncEventMessagePartRemovedData `json:"data,required"`
	JSON        syncEventMessagePartRemovedJSON `json:"-"`
}

type syncEventMessagePartRemovedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessagePartRemoved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessagePartRemovedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventMessagePartRemoved) implementsGlobalEventPayload() {}

type SyncEventMessagePartRemovedData struct {
	SessionID string                              `json:"sessionID,required"`
	MessageID string                              `json:"messageID,required"`
	PartID    string                              `json:"partID,required"`
	JSON      syncEventMessagePartRemovedDataJSON `json:"-"`
}

type syncEventMessagePartRemovedDataJSON struct {
	SessionID   apijson.Field
	MessageID   apijson.Field
	PartID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventMessagePartRemovedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventMessagePartRemovedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionCreated is a V1 sync event with name "session.created.1".
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
	Type        SyncEventSessionCreatedType `json:"type,required"`
	Name        string                      `json:"name,required"`
	ID          string                      `json:"id,required"`
	Seq         int64                       `json:"seq,required"`
	AggregateID string                      `json:"aggregateID,required"`
	Data        SyncEventSessionCreatedData `json:"data,required"`
	JSON        syncEventSessionCreatedJSON `json:"-"`
}

type syncEventSessionCreatedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionCreatedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionCreated) implementsGlobalEventPayload() {}

type SyncEventSessionCreatedData struct {
	SessionID string                          `json:"sessionID,required"`
	Info      Session                         `json:"info,required"`
	JSON      syncEventSessionCreatedDataJSON `json:"-"`
}

type syncEventSessionCreatedDataJSON struct {
	SessionID   apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionCreatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionCreatedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionUpdated is a V1 sync event with name "session.updated.1".
// The data.info field is a partial session update object.
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
	Type        SyncEventSessionUpdatedType `json:"type,required"`
	Name        string                      `json:"name,required"`
	ID          string                      `json:"id,required"`
	Seq         int64                       `json:"seq,required"`
	AggregateID string                      `json:"aggregateID,required"`
	Data        SyncEventSessionUpdatedData `json:"data,required"`
	JSON        syncEventSessionUpdatedJSON `json:"-"`
}

type syncEventSessionUpdatedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionUpdated) implementsGlobalEventPayload() {}

type SyncEventSessionUpdatedData struct {
	SessionID string                          `json:"sessionID,required"`
	Info      SyncEventSessionUpdatedDataInfo `json:"info,required"`
	JSON      syncEventSessionUpdatedDataJSON `json:"-"`
}

type syncEventSessionUpdatedDataJSON struct {
	SessionID   apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionUpdatedDataInfo is a partial session update object.
// All fields are optional.
type SyncEventSessionUpdatedDataInfo struct {
	ID          string                                 `json:"id"`
	Slug        string                                 `json:"slug"`
	ProjectID   string                                 `json:"projectID"`
	WorkspaceID string                                 `json:"workspaceID"`
	Directory   string                                 `json:"directory"`
	Path        string                                 `json:"path"`
	ParentID    string                                 `json:"parentID"`
	Summary     SyncEventSessionUpdatedDataInfoSummary `json:"summary"`
	Cost        float64                                `json:"cost"`
	Tokens      SyncEventSessionUpdatedDataInfoTokens  `json:"tokens"`
	Share       SyncEventSessionUpdatedDataInfoShare   `json:"share"`
	Title       string                                 `json:"title"`
	Agent       string                                 `json:"agent"`
	Model       SyncEventSessionUpdatedDataInfoModel   `json:"model"`
	Version     string                                 `json:"version"`
	Time        SyncEventSessionUpdatedDataInfoTime    `json:"time"`
	Permission  PermissionRuleset                      `json:"permission"`
	Revert      SyncEventSessionUpdatedDataInfoRevert  `json:"revert"`
	JSON        syncEventSessionUpdatedDataInfoJSON    `json:"-"`
}

type syncEventSessionUpdatedDataInfoJSON struct {
	ID          apijson.Field
	Slug        apijson.Field
	ProjectID   apijson.Field
	WorkspaceID apijson.Field
	Directory   apijson.Field
	Path        apijson.Field
	ParentID    apijson.Field
	Summary     apijson.Field
	Cost        apijson.Field
	Tokens      apijson.Field
	Share       apijson.Field
	Title       apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
	Version     apijson.Field
	Time        apijson.Field
	Permission  apijson.Field
	Revert      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionUpdatedDataInfoSummary struct {
	Additions int64                                      `json:"additions,required"`
	Deletions int64                                      `json:"deletions,required"`
	Files     int64                                      `json:"files,required"`
	JSON      syncEventSessionUpdatedDataInfoSummaryJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoSummaryJSON struct {
	Additions   apijson.Field
	Deletions   apijson.Field
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoSummaryJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionUpdatedDataInfoTokens struct {
	Input     int64                                      `json:"input,required"`
	Output    int64                                      `json:"output,required"`
	Reasoning int64                                      `json:"reasoning,required"`
	Cache     SyncEventSessionUpdatedDataInfoTokensCache `json:"cache,required"`
	JSON      syncEventSessionUpdatedDataInfoTokensJSON  `json:"-"`
}

type syncEventSessionUpdatedDataInfoTokensJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoTokensJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionUpdatedDataInfoTokensCache struct {
	Read  int64                                          `json:"read,required"`
	Write int64                                          `json:"write,required"`
	JSON  syncEventSessionUpdatedDataInfoTokensCacheJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoTokensCacheJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionUpdatedDataInfoShare struct {
	URL  string                                   `json:"url"`
	JSON syncEventSessionUpdatedDataInfoShareJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoShareJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoShareJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionUpdatedDataInfoModel struct {
	ID         string                                   `json:"id,required"`
	ProviderID string                                   `json:"providerID,required"`
	Variant    string                                   `json:"variant"`
	JSON       syncEventSessionUpdatedDataInfoModelJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoModelJSON struct {
	ID          apijson.Field
	ProviderID  apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoModelJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionUpdatedDataInfoTime struct {
	Created    int64                                   `json:"created"`
	Updated    int64                                   `json:"updated"`
	Compacting int64                                   `json:"compacting"`
	Archived   int64                                   `json:"archived"`
	JSON       syncEventSessionUpdatedDataInfoTimeJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoTimeJSON struct {
	Created     apijson.Field
	Updated     apijson.Field
	Compacting  apijson.Field
	Archived    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoTimeJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionUpdatedDataInfoRevert struct {
	MessageID string                                    `json:"messageID,required"`
	PartID    string                                    `json:"partID"`
	Snapshot  string                                    `json:"snapshot"`
	Diff      string                                    `json:"diff"`
	JSON      syncEventSessionUpdatedDataInfoRevertJSON `json:"-"`
}

type syncEventSessionUpdatedDataInfoRevertJSON struct {
	MessageID   apijson.Field
	PartID      apijson.Field
	Snapshot    apijson.Field
	Diff        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionUpdatedDataInfoRevert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionUpdatedDataInfoRevertJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionDeleted is a V1 sync event with name "session.deleted.1".
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
	Type        SyncEventSessionDeletedType `json:"type,required"`
	Name        string                      `json:"name,required"`
	ID          string                      `json:"id,required"`
	Seq         int64                       `json:"seq,required"`
	AggregateID string                      `json:"aggregateID,required"`
	Data        SyncEventSessionDeletedData `json:"data,required"`
	JSON        syncEventSessionDeletedJSON `json:"-"`
}

type syncEventSessionDeletedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionDeleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionDeletedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionDeleted) implementsGlobalEventPayload() {}

type SyncEventSessionDeletedData struct {
	SessionID string                          `json:"sessionID,required"`
	Info      Session                         `json:"info,required"`
	JSON      syncEventSessionDeletedDataJSON `json:"-"`
}

type syncEventSessionDeletedDataJSON struct {
	SessionID   apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionDeletedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionDeletedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextAgentSwitched is a V1 sync event with name "session.next.agent.switched.1".
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
	Type        SyncEventSessionNextAgentSwitchedType `json:"type,required"`
	Name        string                                `json:"name,required"`
	ID          string                                `json:"id,required"`
	Seq         int64                                 `json:"seq,required"`
	AggregateID string                                `json:"aggregateID,required"`
	Data        SyncEventSessionNextAgentSwitchedData `json:"data,required"`
	JSON        syncEventSessionNextAgentSwitchedJSON `json:"-"`
}

type syncEventSessionNextAgentSwitchedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextAgentSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextAgentSwitchedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextAgentSwitched) implementsGlobalEventPayload() {}

type SyncEventSessionNextAgentSwitchedData struct {
	Timestamp int64                                     `json:"timestamp,required"`
	SessionID string                                    `json:"sessionID,required"`
	Agent     string                                    `json:"agent,required"`
	JSON      syncEventSessionNextAgentSwitchedDataJSON `json:"-"`
}

type syncEventSessionNextAgentSwitchedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Agent       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextAgentSwitchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextAgentSwitchedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextModelSwitched is a V1 sync event with name "session.next.model.switched.1".
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
	Type        SyncEventSessionNextModelSwitchedType `json:"type,required"`
	Name        string                                `json:"name,required"`
	ID          string                                `json:"id,required"`
	Seq         int64                                 `json:"seq,required"`
	AggregateID string                                `json:"aggregateID,required"`
	Data        SyncEventSessionNextModelSwitchedData `json:"data,required"`
	JSON        syncEventSessionNextModelSwitchedJSON `json:"-"`
}

type syncEventSessionNextModelSwitchedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextModelSwitched) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextModelSwitchedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextModelSwitched) implementsGlobalEventPayload() {}

type SyncEventSessionNextModelSwitchedData struct {
	Timestamp int64                                      `json:"timestamp,required"`
	SessionID string                                     `json:"sessionID,required"`
	Model     SyncEventSessionNextModelSwitchedDataModel `json:"model,required"`
	JSON      syncEventSessionNextModelSwitchedDataJSON  `json:"-"`
}

type syncEventSessionNextModelSwitchedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Model       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextModelSwitchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextModelSwitchedDataJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionNextModelSwitchedDataModel struct {
	ID         string                                         `json:"id,required"`
	ProviderID string                                         `json:"providerID,required"`
	Variant    string                                         `json:"variant,required"`
	JSON       syncEventSessionNextModelSwitchedDataModelJSON `json:"-"`
}

type syncEventSessionNextModelSwitchedDataModelJSON struct {
	ID          apijson.Field
	ProviderID  apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextModelSwitchedDataModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextModelSwitchedDataModelJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextPrompted is a V1 sync event with name "session.next.prompted.1".
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
	Type        SyncEventSessionNextPromptedType `json:"type,required"`
	Name        string                           `json:"name,required"`
	ID          string                           `json:"id,required"`
	Seq         int64                            `json:"seq,required"`
	AggregateID string                           `json:"aggregateID,required"`
	Data        SyncEventSessionNextPromptedData `json:"data,required"`
	JSON        syncEventSessionNextPromptedJSON `json:"-"`
}

type syncEventSessionNextPromptedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextPrompted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextPromptedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextPrompted) implementsGlobalEventPayload() {}

type SyncEventSessionNextPromptedData struct {
	Timestamp int64  `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	// This field can have the runtime type of map[string]interface{}.
	Prompt interface{}                          `json:"prompt,required"`
	JSON   syncEventSessionNextPromptedDataJSON `json:"-"`
}

type syncEventSessionNextPromptedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Prompt      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextPromptedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextPromptedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextSynthetic is a V1 sync event with name "session.next.synthetic.1".
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
	Type        SyncEventSessionNextSyntheticType `json:"type,required"`
	Name        string                            `json:"name,required"`
	ID          string                            `json:"id,required"`
	Seq         int64                             `json:"seq,required"`
	AggregateID string                            `json:"aggregateID,required"`
	Data        SyncEventSessionNextSyntheticData `json:"data,required"`
	JSON        syncEventSessionNextSyntheticJSON `json:"-"`
}

type syncEventSessionNextSyntheticJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextSynthetic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextSyntheticJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextSynthetic) implementsGlobalEventPayload() {}

type SyncEventSessionNextSyntheticData struct {
	Timestamp int64                                 `json:"timestamp,required"`
	SessionID string                                `json:"sessionID,required"`
	Text      string                                `json:"text,required"`
	JSON      syncEventSessionNextSyntheticDataJSON `json:"-"`
}

type syncEventSessionNextSyntheticDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextSyntheticData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextSyntheticDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextShellStarted is a V1 sync event with name "session.next.shell.started.1".
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
	Type        SyncEventSessionNextShellStartedType `json:"type,required"`
	Name        string                               `json:"name,required"`
	ID          string                               `json:"id,required"`
	Seq         int64                                `json:"seq,required"`
	AggregateID string                               `json:"aggregateID,required"`
	Data        SyncEventSessionNextShellStartedData `json:"data,required"`
	JSON        syncEventSessionNextShellStartedJSON `json:"-"`
}

type syncEventSessionNextShellStartedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextShellStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextShellStartedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextShellStarted) implementsGlobalEventPayload() {}

type SyncEventSessionNextShellStartedData struct {
	Timestamp int64                                    `json:"timestamp,required"`
	SessionID string                                   `json:"sessionID,required"`
	CallID    string                                   `json:"callID,required"`
	Command   string                                   `json:"command,required"`
	JSON      syncEventSessionNextShellStartedDataJSON `json:"-"`
}

type syncEventSessionNextShellStartedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Command     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextShellStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextShellStartedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextShellEnded is a V1 sync event with name "session.next.shell.ended.1".
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
	Type        SyncEventSessionNextShellEndedType `json:"type,required"`
	Name        string                             `json:"name,required"`
	ID          string                             `json:"id,required"`
	Seq         int64                              `json:"seq,required"`
	AggregateID string                             `json:"aggregateID,required"`
	Data        SyncEventSessionNextShellEndedData `json:"data,required"`
	JSON        syncEventSessionNextShellEndedJSON `json:"-"`
}

type syncEventSessionNextShellEndedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextShellEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextShellEndedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextShellEnded) implementsGlobalEventPayload() {}

type SyncEventSessionNextShellEndedData struct {
	Timestamp int64                                  `json:"timestamp,required"`
	SessionID string                                 `json:"sessionID,required"`
	CallID    string                                 `json:"callID,required"`
	Output    string                                 `json:"output,required"`
	JSON      syncEventSessionNextShellEndedDataJSON `json:"-"`
}

type syncEventSessionNextShellEndedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextShellEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextShellEndedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextStepStarted is a V1 sync event with name "session.next.step.started.1".
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
	Type        SyncEventSessionNextStepStartedType `json:"type,required"`
	Name        string                              `json:"name,required"`
	ID          string                              `json:"id,required"`
	Seq         int64                               `json:"seq,required"`
	AggregateID string                              `json:"aggregateID,required"`
	Data        SyncEventSessionNextStepStartedData `json:"data,required"`
	JSON        syncEventSessionNextStepStartedJSON `json:"-"`
}

type syncEventSessionNextStepStartedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextStepStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextStepStartedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextStepStarted) implementsGlobalEventPayload() {}

type SyncEventSessionNextStepStartedData struct {
	Timestamp int64                                    `json:"timestamp,required"`
	SessionID string                                   `json:"sessionID,required"`
	Agent     string                                   `json:"agent,required"`
	Model     SyncEventSessionNextStepStartedDataModel `json:"model,required"`
	Snapshot  string                                   `json:"snapshot"`
	JSON      syncEventSessionNextStepStartedDataJSON  `json:"-"`
}

type syncEventSessionNextStepStartedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
	Snapshot    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextStepStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextStepStartedDataJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionNextStepStartedDataModel struct {
	ID         string                                       `json:"id,required"`
	ProviderID string                                       `json:"providerID,required"`
	Variant    string                                       `json:"variant,required"`
	JSON       syncEventSessionNextStepStartedDataModelJSON `json:"-"`
}

type syncEventSessionNextStepStartedDataModelJSON struct {
	ID          apijson.Field
	ProviderID  apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextStepStartedDataModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextStepStartedDataModelJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextStepEnded is a V1 sync event with name "session.next.step.ended.1".
type SyncEventSessionNextStepEndedType string

const (
	SyncEventSessionNextStepEndedTypeSessionNextStepEnded1 SyncEventSessionNextStepEndedType = "session.next.step.ended.1"
)

func (r SyncEventSessionNextStepEndedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextStepEndedTypeSessionNextStepEnded1:
		return true
	}
	return false
}

type SyncEventSessionNextStepEnded struct {
	Type        SyncEventSessionNextStepEndedType `json:"type,required"`
	Name        string                            `json:"name,required"`
	ID          string                            `json:"id,required"`
	Seq         int64                             `json:"seq,required"`
	AggregateID string                            `json:"aggregateID,required"`
	Data        SyncEventSessionNextStepEndedData `json:"data,required"`
	JSON        syncEventSessionNextStepEndedJSON `json:"-"`
}

type syncEventSessionNextStepEndedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextStepEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextStepEndedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextStepEnded) implementsGlobalEventPayload() {}

type SyncEventSessionNextStepEndedData struct {
	Timestamp int64                                   `json:"timestamp,required"`
	SessionID string                                  `json:"sessionID,required"`
	Finish    string                                  `json:"finish,required"`
	Cost      float64                                 `json:"cost,required"`
	Tokens    SyncEventSessionNextStepEndedDataTokens `json:"tokens,required"`
	Snapshot  string                                  `json:"snapshot"`
	JSON      syncEventSessionNextStepEndedDataJSON   `json:"-"`
}

type syncEventSessionNextStepEndedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Finish      apijson.Field
	Cost        apijson.Field
	Tokens      apijson.Field
	Snapshot    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextStepEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextStepEndedDataJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionNextStepEndedDataTokens struct {
	Input     int64                                        `json:"input,required"`
	Output    int64                                        `json:"output,required"`
	Reasoning int64                                        `json:"reasoning,required"`
	Cache     SyncEventSessionNextStepEndedDataTokensCache `json:"cache,required"`
	JSON      syncEventSessionNextStepEndedDataTokensJSON  `json:"-"`
}

type syncEventSessionNextStepEndedDataTokensJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextStepEndedDataTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextStepEndedDataTokensJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionNextStepEndedDataTokensCache struct {
	Read  int64                                            `json:"read,required"`
	Write int64                                            `json:"write,required"`
	JSON  syncEventSessionNextStepEndedDataTokensCacheJSON `json:"-"`
}

type syncEventSessionNextStepEndedDataTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextStepEndedDataTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextStepEndedDataTokensCacheJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextStepFailed is a V1 sync event with name "session.next.step.failed.1".
type SyncEventSessionNextStepFailedType string

const (
	SyncEventSessionNextStepFailedTypeSessionNextStepFailed1 SyncEventSessionNextStepFailedType = "session.next.step.failed.1"
)

func (r SyncEventSessionNextStepFailedType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextStepFailedTypeSessionNextStepFailed1:
		return true
	}
	return false
}

type SyncEventSessionNextStepFailed struct {
	Type        SyncEventSessionNextStepFailedType `json:"type,required"`
	Name        string                             `json:"name,required"`
	ID          string                             `json:"id,required"`
	Seq         int64                              `json:"seq,required"`
	AggregateID string                             `json:"aggregateID,required"`
	Data        SyncEventSessionNextStepFailedData `json:"data,required"`
	JSON        syncEventSessionNextStepFailedJSON `json:"-"`
}

type syncEventSessionNextStepFailedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextStepFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextStepFailedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextStepFailed) implementsGlobalEventPayload() {}

type SyncEventSessionNextStepFailedData struct {
	Timestamp int64  `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	// This field can have the runtime type of [SessionErrorUnknown].
	Error interface{}                            `json:"error,required"`
	JSON  syncEventSessionNextStepFailedDataJSON `json:"-"`
}

type syncEventSessionNextStepFailedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextStepFailedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextStepFailedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextTextStarted is a V1 sync event with name "session.next.text.started.1".
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
	Type        SyncEventSessionNextTextStartedType `json:"type,required"`
	Name        string                              `json:"name,required"`
	ID          string                              `json:"id,required"`
	Seq         int64                               `json:"seq,required"`
	AggregateID string                              `json:"aggregateID,required"`
	Data        SyncEventSessionNextTextStartedData `json:"data,required"`
	JSON        syncEventSessionNextTextStartedJSON `json:"-"`
}

type syncEventSessionNextTextStartedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextTextStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextTextStartedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextTextStarted) implementsGlobalEventPayload() {}

type SyncEventSessionNextTextStartedData struct {
	Timestamp int64                                   `json:"timestamp,required"`
	SessionID string                                  `json:"sessionID,required"`
	JSON      syncEventSessionNextTextStartedDataJSON `json:"-"`
}

type syncEventSessionNextTextStartedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextTextStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextTextStartedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextTextDelta is a V1 sync event with name "session.next.text.delta.1".
type SyncEventSessionNextTextDeltaType string

const (
	SyncEventSessionNextTextDeltaTypeSessionNextTextDelta1 SyncEventSessionNextTextDeltaType = "session.next.text.delta.1"
)

func (r SyncEventSessionNextTextDeltaType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextTextDeltaTypeSessionNextTextDelta1:
		return true
	}
	return false
}

type SyncEventSessionNextTextDelta struct {
	Type        SyncEventSessionNextTextDeltaType `json:"type,required"`
	Name        string                            `json:"name,required"`
	ID          string                            `json:"id,required"`
	Seq         int64                             `json:"seq,required"`
	AggregateID string                            `json:"aggregateID,required"`
	Data        SyncEventSessionNextTextDeltaData `json:"data,required"`
	JSON        syncEventSessionNextTextDeltaJSON `json:"-"`
}

type syncEventSessionNextTextDeltaJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextTextDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextTextDeltaJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextTextDelta) implementsGlobalEventPayload() {}

type SyncEventSessionNextTextDeltaData struct {
	Timestamp int64                                 `json:"timestamp,required"`
	SessionID string                                `json:"sessionID,required"`
	Delta     string                                `json:"delta,required"`
	JSON      syncEventSessionNextTextDeltaDataJSON `json:"-"`
}

type syncEventSessionNextTextDeltaDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Delta       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextTextDeltaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextTextDeltaDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextTextEnded is a V1 sync event with name "session.next.text.ended.1".
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
	Type        SyncEventSessionNextTextEndedType `json:"type,required"`
	Name        string                            `json:"name,required"`
	ID          string                            `json:"id,required"`
	Seq         int64                             `json:"seq,required"`
	AggregateID string                            `json:"aggregateID,required"`
	Data        SyncEventSessionNextTextEndedData `json:"data,required"`
	JSON        syncEventSessionNextTextEndedJSON `json:"-"`
}

type syncEventSessionNextTextEndedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextTextEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextTextEndedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextTextEnded) implementsGlobalEventPayload() {}

type SyncEventSessionNextTextEndedData struct {
	Timestamp int64                                 `json:"timestamp,required"`
	SessionID string                                `json:"sessionID,required"`
	Text      string                                `json:"text,required"`
	JSON      syncEventSessionNextTextEndedDataJSON `json:"-"`
}

type syncEventSessionNextTextEndedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextTextEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextTextEndedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextReasoningStarted is a V1 sync event with name "session.next.reasoning.started.1".
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
	Type        SyncEventSessionNextReasoningStartedType `json:"type,required"`
	Name        string                                   `json:"name,required"`
	ID          string                                   `json:"id,required"`
	Seq         int64                                    `json:"seq,required"`
	AggregateID string                                   `json:"aggregateID,required"`
	Data        SyncEventSessionNextReasoningStartedData `json:"data,required"`
	JSON        syncEventSessionNextReasoningStartedJSON `json:"-"`
}

type syncEventSessionNextReasoningStartedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextReasoningStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextReasoningStartedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextReasoningStarted) implementsGlobalEventPayload() {}

type SyncEventSessionNextReasoningStartedData struct {
	Timestamp   int64                                        `json:"timestamp,required"`
	SessionID   string                                       `json:"sessionID,required"`
	ReasoningID string                                       `json:"reasoningID,required"`
	JSON        syncEventSessionNextReasoningStartedDataJSON `json:"-"`
}

type syncEventSessionNextReasoningStartedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	ReasoningID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextReasoningStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextReasoningStartedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextReasoningDelta is a V1 sync event with name "session.next.reasoning.delta.1".
type SyncEventSessionNextReasoningDeltaType string

const (
	SyncEventSessionNextReasoningDeltaTypeSessionNextReasoningDelta1 SyncEventSessionNextReasoningDeltaType = "session.next.reasoning.delta.1"
)

func (r SyncEventSessionNextReasoningDeltaType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextReasoningDeltaTypeSessionNextReasoningDelta1:
		return true
	}
	return false
}

type SyncEventSessionNextReasoningDelta struct {
	Type        SyncEventSessionNextReasoningDeltaType `json:"type,required"`
	Name        string                                 `json:"name,required"`
	ID          string                                 `json:"id,required"`
	Seq         int64                                  `json:"seq,required"`
	AggregateID string                                 `json:"aggregateID,required"`
	Data        SyncEventSessionNextReasoningDeltaData `json:"data,required"`
	JSON        syncEventSessionNextReasoningDeltaJSON `json:"-"`
}

type syncEventSessionNextReasoningDeltaJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextReasoningDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextReasoningDeltaJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextReasoningDelta) implementsGlobalEventPayload() {}

type SyncEventSessionNextReasoningDeltaData struct {
	Timestamp   int64                                      `json:"timestamp,required"`
	SessionID   string                                     `json:"sessionID,required"`
	ReasoningID string                                     `json:"reasoningID,required"`
	Delta       string                                     `json:"delta,required"`
	JSON        syncEventSessionNextReasoningDeltaDataJSON `json:"-"`
}

type syncEventSessionNextReasoningDeltaDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	ReasoningID apijson.Field
	Delta       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextReasoningDeltaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextReasoningDeltaDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextReasoningEnded is a V1 sync event with name "session.next.reasoning.ended.1".
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
	Type        SyncEventSessionNextReasoningEndedType `json:"type,required"`
	Name        string                                 `json:"name,required"`
	ID          string                                 `json:"id,required"`
	Seq         int64                                  `json:"seq,required"`
	AggregateID string                                 `json:"aggregateID,required"`
	Data        SyncEventSessionNextReasoningEndedData `json:"data,required"`
	JSON        syncEventSessionNextReasoningEndedJSON `json:"-"`
}

type syncEventSessionNextReasoningEndedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextReasoningEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextReasoningEndedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextReasoningEnded) implementsGlobalEventPayload() {}

type SyncEventSessionNextReasoningEndedData struct {
	Timestamp   int64                                      `json:"timestamp,required"`
	SessionID   string                                     `json:"sessionID,required"`
	ReasoningID string                                     `json:"reasoningID,required"`
	Text        string                                     `json:"text,required"`
	JSON        syncEventSessionNextReasoningEndedDataJSON `json:"-"`
}

type syncEventSessionNextReasoningEndedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	ReasoningID apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextReasoningEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextReasoningEndedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextToolInputStarted is a V1 sync event with name "session.next.tool.input.started.1".
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
	Type        SyncEventSessionNextToolInputStartedType `json:"type,required"`
	Name        string                                   `json:"name,required"`
	ID          string                                   `json:"id,required"`
	Seq         int64                                    `json:"seq,required"`
	AggregateID string                                   `json:"aggregateID,required"`
	Data        SyncEventSessionNextToolInputStartedData `json:"data,required"`
	JSON        syncEventSessionNextToolInputStartedJSON `json:"-"`
}

type syncEventSessionNextToolInputStartedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolInputStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolInputStartedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextToolInputStarted) implementsGlobalEventPayload() {}

type SyncEventSessionNextToolInputStartedData struct {
	Timestamp int64                                        `json:"timestamp,required"`
	SessionID string                                       `json:"sessionID,required"`
	CallID    string                                       `json:"callID,required"`
	Name      string                                       `json:"name,required"`
	JSON      syncEventSessionNextToolInputStartedDataJSON `json:"-"`
}

type syncEventSessionNextToolInputStartedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolInputStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolInputStartedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextToolInputDelta is a V1 sync event with name "session.next.tool.input.delta.1".
type SyncEventSessionNextToolInputDeltaType string

const (
	SyncEventSessionNextToolInputDeltaTypeSessionNextToolInputDelta1 SyncEventSessionNextToolInputDeltaType = "session.next.tool.input.delta.1"
)

func (r SyncEventSessionNextToolInputDeltaType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextToolInputDeltaTypeSessionNextToolInputDelta1:
		return true
	}
	return false
}

type SyncEventSessionNextToolInputDelta struct {
	Type        SyncEventSessionNextToolInputDeltaType `json:"type,required"`
	Name        string                                 `json:"name,required"`
	ID          string                                 `json:"id,required"`
	Seq         int64                                  `json:"seq,required"`
	AggregateID string                                 `json:"aggregateID,required"`
	Data        SyncEventSessionNextToolInputDeltaData `json:"data,required"`
	JSON        syncEventSessionNextToolInputDeltaJSON `json:"-"`
}

type syncEventSessionNextToolInputDeltaJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolInputDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolInputDeltaJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextToolInputDelta) implementsGlobalEventPayload() {}

type SyncEventSessionNextToolInputDeltaData struct {
	Timestamp int64                                      `json:"timestamp,required"`
	SessionID string                                     `json:"sessionID,required"`
	CallID    string                                     `json:"callID,required"`
	Delta     string                                     `json:"delta,required"`
	JSON      syncEventSessionNextToolInputDeltaDataJSON `json:"-"`
}

type syncEventSessionNextToolInputDeltaDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Delta       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolInputDeltaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolInputDeltaDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextToolInputEnded is a V1 sync event with name "session.next.tool.input.ended.1".
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
	Type        SyncEventSessionNextToolInputEndedType `json:"type,required"`
	Name        string                                 `json:"name,required"`
	ID          string                                 `json:"id,required"`
	Seq         int64                                  `json:"seq,required"`
	AggregateID string                                 `json:"aggregateID,required"`
	Data        SyncEventSessionNextToolInputEndedData `json:"data,required"`
	JSON        syncEventSessionNextToolInputEndedJSON `json:"-"`
}

type syncEventSessionNextToolInputEndedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolInputEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolInputEndedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextToolInputEnded) implementsGlobalEventPayload() {}

type SyncEventSessionNextToolInputEndedData struct {
	Timestamp int64                                      `json:"timestamp,required"`
	SessionID string                                     `json:"sessionID,required"`
	CallID    string                                     `json:"callID,required"`
	Text      string                                     `json:"text,required"`
	JSON      syncEventSessionNextToolInputEndedDataJSON `json:"-"`
}

type syncEventSessionNextToolInputEndedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolInputEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolInputEndedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextToolCalled is a V1 sync event with name "session.next.tool.called.1".
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
	Type        SyncEventSessionNextToolCalledType `json:"type,required"`
	Name        string                             `json:"name,required"`
	ID          string                             `json:"id,required"`
	Seq         int64                              `json:"seq,required"`
	AggregateID string                             `json:"aggregateID,required"`
	Data        SyncEventSessionNextToolCalledData `json:"data,required"`
	JSON        syncEventSessionNextToolCalledJSON `json:"-"`
}

type syncEventSessionNextToolCalledJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolCalled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolCalledJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextToolCalled) implementsGlobalEventPayload() {}

type SyncEventSessionNextToolCalledData struct {
	Timestamp int64  `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID    string `json:"callID,required"`
	Tool      string `json:"tool,required"`
	// This field can have the runtime type of map[string]interface{}.
	Input    interface{}                                `json:"input,required"`
	Provider SyncEventSessionNextToolCalledDataProvider `json:"provider,required"`
	JSON     syncEventSessionNextToolCalledDataJSON     `json:"-"`
}

type syncEventSessionNextToolCalledDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Tool        apijson.Field
	Input       apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolCalledData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolCalledDataJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionNextToolCalledDataProvider struct {
	Executed bool                                           `json:"executed,required"`
	Metadata interface{}                                    `json:"metadata"`
	JSON     syncEventSessionNextToolCalledDataProviderJSON `json:"-"`
}

type syncEventSessionNextToolCalledDataProviderJSON struct {
	Executed    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolCalledDataProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolCalledDataProviderJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextToolProgress is a V1 sync event with name "session.next.tool.progress.1".
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
	Type        SyncEventSessionNextToolProgressType `json:"type,required"`
	Name        string                               `json:"name,required"`
	ID          string                               `json:"id,required"`
	Seq         int64                                `json:"seq,required"`
	AggregateID string                               `json:"aggregateID,required"`
	Data        SyncEventSessionNextToolProgressData `json:"data,required"`
	JSON        syncEventSessionNextToolProgressJSON `json:"-"`
}

type syncEventSessionNextToolProgressJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolProgressJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextToolProgress) implementsGlobalEventPayload() {}

type SyncEventSessionNextToolProgressData struct {
	Timestamp int64  `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID    string `json:"callID,required"`
	// This field can have the runtime type of map[string]interface{}.
	Structured interface{} `json:"structured,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content []interface{}                            `json:"content,required"`
	JSON    syncEventSessionNextToolProgressDataJSON `json:"-"`
}

type syncEventSessionNextToolProgressDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Structured  apijson.Field
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolProgressData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolProgressDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextToolSuccess is a V1 sync event with name "session.next.tool.success.1".
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
	Type        SyncEventSessionNextToolSuccessType `json:"type,required"`
	Name        string                              `json:"name,required"`
	ID          string                              `json:"id,required"`
	Seq         int64                               `json:"seq,required"`
	AggregateID string                              `json:"aggregateID,required"`
	Data        SyncEventSessionNextToolSuccessData `json:"data,required"`
	JSON        syncEventSessionNextToolSuccessJSON `json:"-"`
}

type syncEventSessionNextToolSuccessJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolSuccessJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextToolSuccess) implementsGlobalEventPayload() {}

type SyncEventSessionNextToolSuccessData struct {
	Timestamp int64  `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID    string `json:"callID,required"`
	// This field can have the runtime type of map[string]interface{}.
	Structured interface{} `json:"structured,required"`
	// This field can have the runtime type of [[]ToolTextContent], [[]ToolFileContent].
	Content  []interface{}                               `json:"content,required"`
	Provider SyncEventSessionNextToolSuccessDataProvider `json:"provider,required"`
	JSON     syncEventSessionNextToolSuccessDataJSON     `json:"-"`
}

type syncEventSessionNextToolSuccessDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Structured  apijson.Field
	Content     apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolSuccessData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolSuccessDataJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionNextToolSuccessDataProvider struct {
	Executed bool                                            `json:"executed,required"`
	Metadata interface{}                                     `json:"metadata"`
	JSON     syncEventSessionNextToolSuccessDataProviderJSON `json:"-"`
}

type syncEventSessionNextToolSuccessDataProviderJSON struct {
	Executed    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolSuccessDataProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolSuccessDataProviderJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextToolFailed is a V1 sync event with name "session.next.tool.failed.1".
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
	Type        SyncEventSessionNextToolFailedType `json:"type,required"`
	Name        string                             `json:"name,required"`
	ID          string                             `json:"id,required"`
	Seq         int64                              `json:"seq,required"`
	AggregateID string                             `json:"aggregateID,required"`
	Data        SyncEventSessionNextToolFailedData `json:"data,required"`
	JSON        syncEventSessionNextToolFailedJSON `json:"-"`
}

type syncEventSessionNextToolFailedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolFailedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextToolFailed) implementsGlobalEventPayload() {}

type SyncEventSessionNextToolFailedData struct {
	Timestamp int64  `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	CallID    string `json:"callID,required"`
	// This field can have the runtime type of [SessionErrorUnknown].
	Error    interface{}                                `json:"error,required"`
	Provider SyncEventSessionNextToolFailedDataProvider `json:"provider,required"`
	JSON     syncEventSessionNextToolFailedDataJSON     `json:"-"`
}

type syncEventSessionNextToolFailedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	CallID      apijson.Field
	Error       apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolFailedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolFailedDataJSON) RawJSON() string {
	return r.raw
}

type SyncEventSessionNextToolFailedDataProvider struct {
	Executed bool                                           `json:"executed,required"`
	Metadata interface{}                                    `json:"metadata"`
	JSON     syncEventSessionNextToolFailedDataProviderJSON `json:"-"`
}

type syncEventSessionNextToolFailedDataProviderJSON struct {
	Executed    apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextToolFailedDataProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextToolFailedDataProviderJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextRetried is a V1 sync event with name "session.next.retried.1".
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
	Type        SyncEventSessionNextRetriedType `json:"type,required"`
	Name        string                          `json:"name,required"`
	ID          string                          `json:"id,required"`
	Seq         int64                           `json:"seq,required"`
	AggregateID string                          `json:"aggregateID,required"`
	Data        SyncEventSessionNextRetriedData `json:"data,required"`
	JSON        syncEventSessionNextRetriedJSON `json:"-"`
}

type syncEventSessionNextRetriedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextRetried) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextRetriedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextRetried) implementsGlobalEventPayload() {}

type SyncEventSessionNextRetriedData struct {
	Timestamp int64  `json:"timestamp,required"`
	SessionID string `json:"sessionID,required"`
	Attempt   int64  `json:"attempt,required"`
	// This field can have the runtime type of map[string]interface{}.
	Error interface{}                         `json:"error,required"`
	JSON  syncEventSessionNextRetriedDataJSON `json:"-"`
}

type syncEventSessionNextRetriedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Attempt     apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextRetriedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextRetriedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextCompactionStarted is a V1 sync event with name "session.next.compaction.started.1".
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
	Type        SyncEventSessionNextCompactionStartedType `json:"type,required"`
	Name        string                                    `json:"name,required"`
	ID          string                                    `json:"id,required"`
	Seq         int64                                     `json:"seq,required"`
	AggregateID string                                    `json:"aggregateID,required"`
	Data        SyncEventSessionNextCompactionStartedData `json:"data,required"`
	JSON        syncEventSessionNextCompactionStartedJSON `json:"-"`
}

type syncEventSessionNextCompactionStartedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextCompactionStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextCompactionStartedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextCompactionStarted) implementsGlobalEventPayload() {}

type SyncEventSessionNextCompactionStartedData struct {
	Timestamp int64                                         `json:"timestamp,required"`
	SessionID string                                        `json:"sessionID,required"`
	Reason    string                                        `json:"reason,required"`
	JSON      syncEventSessionNextCompactionStartedDataJSON `json:"-"`
}

type syncEventSessionNextCompactionStartedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextCompactionStartedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextCompactionStartedDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextCompactionDelta is a V1 sync event with name "session.next.compaction.delta.1".
type SyncEventSessionNextCompactionDeltaType string

const (
	SyncEventSessionNextCompactionDeltaTypeSessionNextCompactionDelta1 SyncEventSessionNextCompactionDeltaType = "session.next.compaction.delta.1"
)

func (r SyncEventSessionNextCompactionDeltaType) IsKnown() bool {
	switch r {
	case SyncEventSessionNextCompactionDeltaTypeSessionNextCompactionDelta1:
		return true
	}
	return false
}

type SyncEventSessionNextCompactionDelta struct {
	Type        SyncEventSessionNextCompactionDeltaType `json:"type,required"`
	Name        string                                  `json:"name,required"`
	ID          string                                  `json:"id,required"`
	Seq         int64                                   `json:"seq,required"`
	AggregateID string                                  `json:"aggregateID,required"`
	Data        SyncEventSessionNextCompactionDeltaData `json:"data,required"`
	JSON        syncEventSessionNextCompactionDeltaJSON `json:"-"`
}

type syncEventSessionNextCompactionDeltaJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextCompactionDelta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextCompactionDeltaJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextCompactionDelta) implementsGlobalEventPayload() {}

type SyncEventSessionNextCompactionDeltaData struct {
	Timestamp int64                                       `json:"timestamp,required"`
	SessionID string                                      `json:"sessionID,required"`
	Text      string                                      `json:"text,required"`
	JSON      syncEventSessionNextCompactionDeltaDataJSON `json:"-"`
}

type syncEventSessionNextCompactionDeltaDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextCompactionDeltaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextCompactionDeltaDataJSON) RawJSON() string {
	return r.raw
}

// SyncEventSessionNextCompactionEnded is a V1 sync event with name "session.next.compaction.ended.1".
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
	Type        SyncEventSessionNextCompactionEndedType `json:"type,required"`
	Name        string                                  `json:"name,required"`
	ID          string                                  `json:"id,required"`
	Seq         int64                                   `json:"seq,required"`
	AggregateID string                                  `json:"aggregateID,required"`
	Data        SyncEventSessionNextCompactionEndedData `json:"data,required"`
	JSON        syncEventSessionNextCompactionEndedJSON `json:"-"`
}

type syncEventSessionNextCompactionEndedJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Seq         apijson.Field
	AggregateID apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextCompactionEnded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextCompactionEndedJSON) RawJSON() string {
	return r.raw
}

func (r SyncEventSessionNextCompactionEnded) implementsGlobalEventPayload() {}

type SyncEventSessionNextCompactionEndedData struct {
	Timestamp int64                                       `json:"timestamp,required"`
	SessionID string                                      `json:"sessionID,required"`
	Text      string                                      `json:"text,required"`
	Include   string                                      `json:"include"`
	JSON      syncEventSessionNextCompactionEndedDataJSON `json:"-"`
}

type syncEventSessionNextCompactionEndedDataJSON struct {
	Timestamp   apijson.Field
	SessionID   apijson.Field
	Text        apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncEventSessionNextCompactionEndedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncEventSessionNextCompactionEndedDataJSON) RawJSON() string {
	return r.raw
}
