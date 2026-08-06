// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// ExperimentalSessionService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalSessionService] method instead.
type ExperimentalSessionService struct {
	Options []option.RequestOption
}

// NewExperimentalSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExperimentalSessionService(opts ...option.RequestOption) (r *ExperimentalSessionService) {
	r = &ExperimentalSessionService{}
	r.Options = opts
	return
}

// List sessions across projects
func (r *ExperimentalSessionService) List(ctx context.Context, query ExperimentalSessionListParams, opts ...option.RequestOption) (res *[]GlobalSession, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Background subagents
//
// Detach any synchronous subagents currently blocking the session and continue them in the background.
func (r *ExperimentalSessionService) Background(ctx context.Context, sessionID string, query ExperimentalSessionBackgroundParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("experimental/session/%s/background", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, query, &res, opts...)
	return
}

type ExperimentalSessionListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Roots     param.Field[bool]   `query:"roots"`
	Start     param.Field[int64]  `query:"start"`
	Cursor    param.Field[int64]  `query:"cursor"`
	Search    param.Field[string] `query:"search"`
	Limit     param.Field[int64]  `query:"limit"`
	Archived  param.Field[bool]   `query:"archived"`
}

func (r ExperimentalSessionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentalSessionBackgroundParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ExperimentalSessionBackgroundParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalSession struct {
	ID        string            `json:"id,required"`
	Slug      string            `json:"slug,required"`
	ProjectID string            `json:"projectID,required"`
	Directory string            `json:"directory,required"`
	Title     string            `json:"title,required"`
	Version   string            `json:"version,required"`
	Time      GlobalSessionTime `json:"time,required"`
	// This field can have the runtime type of [ProjectSummary, nil].
	Project     any                  `json:"project,required"`
	WorkspaceID string               `json:"workspaceID"`
	Path        string               `json:"path"`
	ParentID    string               `json:"parentID"`
	Summary     GlobalSessionSummary `json:"summary"`
	Cost        float64              `json:"cost"`
	Tokens      GlobalSessionTokens  `json:"tokens"`
	Share       GlobalSessionShare   `json:"share"`
	Agent       string               `json:"agent"`
	Model       GlobalSessionModel   `json:"model"`
	Permission  PermissionRuleset    `json:"permission"`
	Revert      GlobalSessionRevert  `json:"revert"`
	Metadata    map[string]any       `json:"metadata"`
	JSON        globalSessionJSON    `json:"-"`
}

// globalSessionJSON contains the JSON metadata for the struct [GlobalSession]
type globalSessionJSON struct {
	ID          apijson.Field
	Slug        apijson.Field
	ProjectID   apijson.Field
	Directory   apijson.Field
	Title       apijson.Field
	Version     apijson.Field
	Time        apijson.Field
	Project     apijson.Field
	WorkspaceID apijson.Field
	Path        apijson.Field
	ParentID    apijson.Field
	Summary     apijson.Field
	Cost        apijson.Field
	Tokens      apijson.Field
	Share       apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
	Permission  apijson.Field
	Revert      apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionJSON) RawJSON() string {
	return r.raw
}

type GlobalSessionTime struct {
	Created    int64                 `json:"created,required"`
	Updated    int64                 `json:"updated,required"`
	Compacting int64                 `json:"compacting"`
	Archived   int64                 `json:"archived"`
	JSON       globalSessionTimeJSON `json:"-"`
}

// globalSessionTimeJSON contains the JSON metadata for the struct [GlobalSessionTime]
type globalSessionTimeJSON struct {
	Created     apijson.Field
	Updated     apijson.Field
	Compacting  apijson.Field
	Archived    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSessionTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionTimeJSON) RawJSON() string {
	return r.raw
}

type GlobalSessionTokens struct {
	Input     int64                    `json:"input,required"`
	Output    int64                    `json:"output,required"`
	Reasoning int64                    `json:"reasoning,required"`
	Cache     GlobalSessionTokensCache `json:"cache,required"`
	JSON      globalSessionTokensJSON  `json:"-"`
}

// globalSessionTokensJSON contains the JSON metadata for the struct [GlobalSessionTokens]
type globalSessionTokensJSON struct {
	Input       apijson.Field
	Output      apijson.Field
	Reasoning   apijson.Field
	Cache       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSessionTokens) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionTokensJSON) RawJSON() string {
	return r.raw
}

type GlobalSessionTokensCache struct {
	Read  int64                        `json:"read,required"`
	Write int64                        `json:"write,required"`
	JSON  globalSessionTokensCacheJSON `json:"-"`
}

// globalSessionTokensCacheJSON contains the JSON metadata for the struct
// [GlobalSessionTokensCache]
type globalSessionTokensCacheJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSessionTokensCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionTokensCacheJSON) RawJSON() string {
	return r.raw
}

type GlobalSessionShare struct {
	URL  string                 `json:"url,required"`
	JSON globalSessionShareJSON `json:"-"`
}

// globalSessionShareJSON contains the JSON metadata for the struct [GlobalSessionShare]
type globalSessionShareJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSessionShare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionShareJSON) RawJSON() string {
	return r.raw
}

type GlobalSessionModel struct {
	ID         string                 `json:"id,required"`
	ProviderID string                 `json:"providerID,required"`
	Variant    string                 `json:"variant"`
	JSON       globalSessionModelJSON `json:"-"`
}

// globalSessionModelJSON contains the JSON metadata for the struct [GlobalSessionModel]
type globalSessionModelJSON struct {
	ID          apijson.Field
	ProviderID  apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSessionModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionModelJSON) RawJSON() string {
	return r.raw
}

type GlobalSessionRevert struct {
	MessageID string                  `json:"messageID,required"`
	PartID    string                  `json:"partID"`
	Snapshot  string                  `json:"snapshot"`
	Diff      string                  `json:"diff"`
	JSON      globalSessionRevertJSON `json:"-"`
}

// globalSessionRevertJSON contains the JSON metadata for the struct [GlobalSessionRevert]
type globalSessionRevertJSON struct {
	MessageID   apijson.Field
	PartID      apijson.Field
	Snapshot    apijson.Field
	Diff        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSessionRevert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionRevertJSON) RawJSON() string {
	return r.raw
}

type GlobalSessionSummary struct {
	Additions int64                    `json:"additions,required"`
	Deletions int64                    `json:"deletions,required"`
	Files     int64                    `json:"files,required"`
	Diffs     []SnapshotFileDiff       `json:"diffs"`
	JSON      globalSessionSummaryJSON `json:"-"`
}

type globalSessionSummaryJSON struct {
	Additions   apijson.Field
	Deletions   apijson.Field
	Files       apijson.Field
	Diffs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GlobalSessionSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSessionSummaryJSON) RawJSON() string {
	return r.raw
}

type ProjectSummary struct {
	ID       string             `json:"id,required"`
	Worktree string             `json:"worktree,required"`
	Name     string             `json:"name"`
	JSON     projectSummaryJSON `json:"-"`
}

// projectSummaryJSON contains the JSON metadata for the struct [ProjectSummary]
type projectSummaryJSON struct {
	ID          apijson.Field
	Worktree    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectSummaryJSON) RawJSON() string {
	return r.raw
}
