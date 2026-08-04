// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// ExperimentalConsoleService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalConsoleService] method instead.
type ExperimentalConsoleService struct {
	Options []option.RequestOption
}

// NewExperimentalConsoleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExperimentalConsoleService(opts ...option.RequestOption) (r *ExperimentalConsoleService) {
	r = &ExperimentalConsoleService{}
	r.Options = opts
	return
}

// Get active Console provider metadata
func (r *ExperimentalConsoleService) Get(ctx context.Context, query ExperimentalConsoleGetParams, opts ...option.RequestOption) (res *ConsoleState, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/console"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List switchable Console orgs
func (r *ExperimentalConsoleService) ListOrgs(ctx context.Context, query ExperimentalConsoleListOrgsParams, opts ...option.RequestOption) (res *ConsoleListOrgsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/console/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Switch active Console org
func (r *ExperimentalConsoleService) SwitchOrg(ctx context.Context, body ExperimentalConsoleSwitchOrgParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/console/switch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ConsoleState struct {
	ConsoleManagedProviders []string         `json:"consoleManagedProviders,required"`
	ActiveOrgName           string           `json:"activeOrgName"`
	SwitchableOrgCount      int64            `json:"switchableOrgCount,required"`
	JSON                    consoleStateJSON `json:"-"`
}

type consoleStateJSON struct {
	ConsoleManagedProviders apijson.Field
	ActiveOrgName           apijson.Field
	SwitchableOrgCount      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ConsoleState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consoleStateJSON) RawJSON() string {
	return r.raw
}

type ExperimentalConsoleGetParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ExperimentalConsoleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConsoleOrg struct {
	AccountID    string         `json:"accountID,required"`
	AccountEmail string         `json:"accountEmail,required"`
	AccountURL   string         `json:"accountUrl,required"`
	OrgID        string         `json:"orgID,required"`
	OrgName      string         `json:"orgName,required"`
	Active       bool           `json:"active,required"`
	JSON         consoleOrgJSON `json:"-"`
}

type consoleOrgJSON struct {
	AccountID    apijson.Field
	AccountEmail apijson.Field
	AccountURL   apijson.Field
	OrgID        apijson.Field
	OrgName      apijson.Field
	Active       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConsoleOrg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consoleOrgJSON) RawJSON() string {
	return r.raw
}

type ConsoleListOrgsResponse struct {
	Orgs []ConsoleOrg                `json:"orgs,required"`
	JSON consoleListOrgsResponseJSON `json:"-"`
}

type consoleListOrgsResponseJSON struct {
	Orgs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConsoleListOrgsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r consoleListOrgsResponseJSON) RawJSON() string {
	return r.raw
}

type ExperimentalConsoleListOrgsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r ExperimentalConsoleListOrgsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ExperimentalConsoleSwitchOrgParams contains the request parameters for switching the active Console org.
type ExperimentalConsoleSwitchOrgParams struct {
	Directory param.Field[string]               `query:"directory"`
	Workspace param.Field[string]               `query:"workspace"`
	Body      ExperimentalConsoleSwitchOrgInput `json:"-"`
}

func (r ExperimentalConsoleSwitchOrgParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [ExperimentalConsoleSwitchOrgParams]'s query parameters as `url.Values`.
func (r ExperimentalConsoleSwitchOrgParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ConsoleSwitchOrgParams is an alias for [ExperimentalConsoleSwitchOrgParams] for backward compatibility.
type ConsoleSwitchOrgParams = ExperimentalConsoleSwitchOrgParams

type ExperimentalConsoleSwitchOrgInput struct {
	AccountID param.Field[string] `json:"accountID,required"`
	OrgID     param.Field[string] `json:"orgID,required"`
}

func (r ExperimentalConsoleSwitchOrgInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ConsoleSwitchOrgInput is an alias for [ExperimentalConsoleSwitchOrgInput] for backward compatibility.
type ConsoleSwitchOrgInput = ExperimentalConsoleSwitchOrgInput
