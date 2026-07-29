// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/tidwall/gjson"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// TuiService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTuiService] method instead.
type TuiService struct {
	Options []option.RequestOption
	Control *TuiControlService
}

// NewTuiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTuiService(opts ...option.RequestOption) (r *TuiService) {
	r = &TuiService{}
	r.Options = opts
	r.Control = NewTuiControlService(opts...)
	return
}

// Append prompt to the TUI
func (r *TuiService) AppendPrompt(ctx context.Context, params TuiAppendPromptParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/append-prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Clear the prompt
func (r *TuiService) ClearPrompt(ctx context.Context, body TuiClearPromptParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/clear-prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute a TUI command (e.g. agent_cycle)
func (r *TuiService) ExecuteCommand(ctx context.Context, params TuiExecuteCommandParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/execute-command"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Open the help dialog
func (r *TuiService) OpenHelp(ctx context.Context, body TuiOpenHelpParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/open-help"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Open the model dialog
func (r *TuiService) OpenModels(ctx context.Context, body TuiOpenModelsParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/open-models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Open the session dialog
func (r *TuiService) OpenSessions(ctx context.Context, body TuiOpenSessionsParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/open-sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Open the theme dialog
func (r *TuiService) OpenThemes(ctx context.Context, body TuiOpenThemesParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/open-themes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Show a toast notification in the TUI
func (r *TuiService) ShowToast(ctx context.Context, params TuiShowToastParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/show-toast"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Submit the prompt
func (r *TuiService) SubmitPrompt(ctx context.Context, body TuiSubmitPromptParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/submit-prompt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Publish a TUI event
func (r *TuiService) Publish(ctx context.Context, body TuiPublishParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/publish"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Select session in TUI
func (r *TuiService) SelectSession(ctx context.Context, body TuiSelectSessionParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/select-session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// TuiControlService contains methods for TUI control operations.
type TuiControlService struct {
	Options []option.RequestOption
}

// NewTuiControlService creates a new TuiControlService.
func NewTuiControlService(opts ...option.RequestOption) (r *TuiControlService) {
	r = &TuiControlService{}
	r.Options = opts
	return
}

// Get next TUI request
func (r *TuiControlService) Next(ctx context.Context, query TuiControlNextParams, opts ...option.RequestOption) (res *TuiControlNextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/control/next"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Submit TUI response
func (r *TuiControlService) Response(ctx context.Context, body TuiControlResponseParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tui/control/response"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TuiAppendPromptParams struct {
	Text      param.Field[string] `json:"text,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r TuiAppendPromptParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [TuiAppendPromptParams]'s query parameters as `url.Values`.
func (r TuiAppendPromptParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiClearPromptParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiClearPromptParams]'s query parameters as `url.Values`.
func (r TuiClearPromptParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiExecuteCommandParams struct {
	Command   param.Field[string] `json:"command,required"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r TuiExecuteCommandParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [TuiExecuteCommandParams]'s query parameters as
// `url.Values`.
func (r TuiExecuteCommandParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiOpenHelpParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiOpenHelpParams]'s query parameters as `url.Values`.
func (r TuiOpenHelpParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiOpenModelsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiOpenModelsParams]'s query parameters as `url.Values`.
func (r TuiOpenModelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiOpenSessionsParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiOpenSessionsParams]'s query parameters as `url.Values`.
func (r TuiOpenSessionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiOpenThemesParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiOpenThemesParams]'s query parameters as `url.Values`.
func (r TuiOpenThemesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiShowToastParams struct {
	Message   param.Field[string]                    `json:"message,required"`
	Variant   param.Field[TuiShowToastParamsVariant] `json:"variant,required"`
	Directory param.Field[string]                    `query:"directory"`
	Workspace param.Field[string]                    `query:"workspace"`
	Title     param.Field[string]                    `json:"title"`
	Duration  param.Field[int64]                     `json:"duration"`
}

func (r TuiShowToastParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [TuiShowToastParams]'s query parameters as `url.Values`.
func (r TuiShowToastParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiShowToastParamsVariant string

const (
	TuiShowToastParamsVariantInfo    TuiShowToastParamsVariant = "info"
	TuiShowToastParamsVariantSuccess TuiShowToastParamsVariant = "success"
	TuiShowToastParamsVariantWarning TuiShowToastParamsVariant = "warning"
	TuiShowToastParamsVariantError   TuiShowToastParamsVariant = "error"
)

func (r TuiShowToastParamsVariant) IsKnown() bool {
	switch r {
	case TuiShowToastParamsVariantInfo, TuiShowToastParamsVariantSuccess, TuiShowToastParamsVariantWarning, TuiShowToastParamsVariantError:
		return true
	}
	return false
}

type TuiSubmitPromptParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [TuiSubmitPromptParams]'s query parameters as `url.Values`.
func (r TuiSubmitPromptParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiPublishParams struct {
	Directory param.Field[string]                    `query:"directory"`
	Workspace param.Field[string]                    `query:"workspace"`
	Body      param.Field[TuiPublishParamsBodyUnion] `json:"-"`
}

func (r TuiPublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [TuiPublishParams]'s query parameters as `url.Values`.
func (r TuiPublishParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiPublishParamsBodyUnion interface {
	implementsTuiPublishParamsBodyUnion()
}

type TuiPublishParamsBodyPromptAppend struct {
	Type       param.Field[TuiPublishParamsBodyPromptAppendType]       `json:"type,required"`
	Properties param.Field[TuiPublishParamsBodyPromptAppendProperties] `json:"properties,required"`
}

func (r TuiPublishParamsBodyPromptAppend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishParamsBodyPromptAppend) implementsTuiPublishParamsBodyUnion() {}

type TuiPublishParamsBodyPromptAppendType string

const (
	TuiPublishParamsBodyPromptAppendTypeTuiPromptAppend TuiPublishParamsBodyPromptAppendType = "tui.prompt.append"
)

func (r TuiPublishParamsBodyPromptAppendType) IsKnown() bool {
	switch r {
	case TuiPublishParamsBodyPromptAppendTypeTuiPromptAppend:
		return true
	}
	return false
}

type TuiPublishParamsBodyPromptAppendProperties struct {
	Text param.Field[string] `json:"text,required"`
}

func (r TuiPublishParamsBodyPromptAppendProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishParamsBodyCommandExecute struct {
	Type       param.Field[TuiPublishParamsBodyCommandExecuteType]       `json:"type,required"`
	Properties param.Field[TuiPublishParamsBodyCommandExecuteProperties] `json:"properties,required"`
}

func (r TuiPublishParamsBodyCommandExecute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishParamsBodyCommandExecute) implementsTuiPublishParamsBodyUnion() {}

type TuiPublishParamsBodyCommandExecuteType string

const (
	TuiPublishParamsBodyCommandExecuteTypeTuiCommandExecute TuiPublishParamsBodyCommandExecuteType = "tui.command.execute"
)

func (r TuiPublishParamsBodyCommandExecuteType) IsKnown() bool {
	switch r {
	case TuiPublishParamsBodyCommandExecuteTypeTuiCommandExecute:
		return true
	}
	return false
}

type TuiPublishParamsBodyCommandExecuteProperties struct {
	Command param.Field[string] `json:"command,required"`
}

func (r TuiPublishParamsBodyCommandExecuteProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TuiPublishParamsBodyCommandExecutePropertiesCommand enumerates the well-known command
// identifiers for the tui.command.execute event. Because the OpenAPI schema uses
// anyOf[enum, string], arbitrary string values are also accepted; use a plain
// string literal when a command is not listed here.
type TuiPublishParamsBodyCommandExecutePropertiesCommand string

const (
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionList         TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.list"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionNew          TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.new"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionShare        TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.share"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionInterrupt    TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.interrupt"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionCompact      TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.compact"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionPageUp       TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.page.up"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionPageDown     TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.page.down"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionLineUp       TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.line.up"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionLineDown     TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.line.down"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionHalfPageUp   TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.half.page.up"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionHalfPageDown TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.half.page.down"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionFirst        TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.first"
	TuiPublishParamsBodyCommandExecutePropertiesCommandSessionLast         TuiPublishParamsBodyCommandExecutePropertiesCommand = "session.last"
	TuiPublishParamsBodyCommandExecutePropertiesCommandPromptClear         TuiPublishParamsBodyCommandExecutePropertiesCommand = "prompt.clear"
	TuiPublishParamsBodyCommandExecutePropertiesCommandPromptSubmit        TuiPublishParamsBodyCommandExecutePropertiesCommand = "prompt.submit"
	TuiPublishParamsBodyCommandExecutePropertiesCommandAgentCycle          TuiPublishParamsBodyCommandExecutePropertiesCommand = "agent.cycle"
)

func (r TuiPublishParamsBodyCommandExecutePropertiesCommand) IsKnown() bool {
	switch r {
	case TuiPublishParamsBodyCommandExecutePropertiesCommandSessionList,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionNew,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionShare,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionInterrupt,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionCompact,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionPageUp,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionPageDown,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionLineUp,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionLineDown,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionHalfPageUp,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionHalfPageDown,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionFirst,
		TuiPublishParamsBodyCommandExecutePropertiesCommandSessionLast,
		TuiPublishParamsBodyCommandExecutePropertiesCommandPromptClear,
		TuiPublishParamsBodyCommandExecutePropertiesCommandPromptSubmit,
		TuiPublishParamsBodyCommandExecutePropertiesCommandAgentCycle:
		return true
	}
	return false
}

type TuiPublishParamsBodyToastShow struct {
	Type       param.Field[TuiPublishParamsBodyToastShowType]       `json:"type,required"`
	Properties param.Field[TuiPublishParamsBodyToastShowProperties] `json:"properties,required"`
}

func (r TuiPublishParamsBodyToastShow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishParamsBodyToastShow) implementsTuiPublishParamsBodyUnion() {}

type TuiPublishParamsBodyToastShowType string

const (
	TuiPublishParamsBodyToastShowTypeTuiToastShow TuiPublishParamsBodyToastShowType = "tui.toast.show"
)

func (r TuiPublishParamsBodyToastShowType) IsKnown() bool {
	switch r {
	case TuiPublishParamsBodyToastShowTypeTuiToastShow:
		return true
	}
	return false
}

type TuiPublishParamsBodyToastShowProperties struct {
	Message  param.Field[string]                                         `json:"message,required"`
	Variant  param.Field[TuiPublishParamsBodyToastShowPropertiesVariant] `json:"variant,required"`
	Title    param.Field[string]                                         `json:"title"`
	Duration param.Field[int64]                                          `json:"duration"`
}

func (r TuiPublishParamsBodyToastShowProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishParamsBodyToastShowPropertiesVariant string

const (
	TuiPublishParamsBodyToastShowPropertiesVariantInfo    TuiPublishParamsBodyToastShowPropertiesVariant = "info"
	TuiPublishParamsBodyToastShowPropertiesVariantSuccess TuiPublishParamsBodyToastShowPropertiesVariant = "success"
	TuiPublishParamsBodyToastShowPropertiesVariantWarning TuiPublishParamsBodyToastShowPropertiesVariant = "warning"
	TuiPublishParamsBodyToastShowPropertiesVariantError   TuiPublishParamsBodyToastShowPropertiesVariant = "error"
)

func (r TuiPublishParamsBodyToastShowPropertiesVariant) IsKnown() bool {
	switch r {
	case TuiPublishParamsBodyToastShowPropertiesVariantInfo, TuiPublishParamsBodyToastShowPropertiesVariantSuccess, TuiPublishParamsBodyToastShowPropertiesVariantWarning, TuiPublishParamsBodyToastShowPropertiesVariantError:
		return true
	}
	return false
}

type TuiPublishParamsBodySessionSelect struct {
	Type       param.Field[TuiPublishParamsBodySessionSelectType]       `json:"type,required"`
	Properties param.Field[TuiPublishParamsBodySessionSelectProperties] `json:"properties,required"`
}

func (r TuiPublishParamsBodySessionSelect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishParamsBodySessionSelect) implementsTuiPublishParamsBodyUnion() {}

type TuiPublishParamsBodySessionSelectType string

const (
	TuiPublishParamsBodySessionSelectTypeTuiSessionSelect TuiPublishParamsBodySessionSelectType = "tui.session.select"
)

func (r TuiPublishParamsBodySessionSelectType) IsKnown() bool {
	switch r {
	case TuiPublishParamsBodySessionSelectTypeTuiSessionSelect:
		return true
	}
	return false
}

type TuiPublishParamsBodySessionSelectProperties struct {
	SessionID param.Field[string] `json:"sessionID,required"`
}

func (r TuiPublishParamsBodySessionSelectProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiSelectSessionParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	SessionID param.Field[string] `json:"sessionID,required"`
}

func (r TuiSelectSessionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [TuiSelectSessionParams]'s query parameters as
// `url.Values`.
func (r TuiSelectSessionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiControlNextParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r TuiControlNextParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TuiControlNextResponse struct {
	Path string `json:"path,required"`
	// Arbitrary JSON value holding the TUI request payload. Per OpenAPI
	// `tui.control.next` response `body` is an unconstrained schema (`{}`), so no
	// fixed set of runtime types applies; callers should type-assert as needed.
	Body any                        `json:"body,required"`
	JSON tuiControlNextResponseJSON `json:"-"`
}

// tuiControlNextResponseJSON contains the JSON metadata for the struct
// [TuiControlNextResponse]
type tuiControlNextResponseJSON struct {
	Path        apijson.Field
	Body        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TuiControlNextResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tuiControlNextResponseJSON) RawJSON() string {
	return r.raw
}

type TuiControlResponseParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Body      param.Field[any]    `json:"-"`
}

func (r TuiControlResponseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [TuiControlResponseParams]'s query parameters as
// `url.Values`.
func (r TuiControlResponseParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeFor[TuiPublishParamsBodyUnion](),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[TuiPublishParamsBodyPromptAppend](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[TuiPublishParamsBodyCommandExecute](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[TuiPublishParamsBodyToastShow](),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeFor[TuiPublishParamsBodySessionSelect](),
		},
	)
}
