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

// TuiService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTuiService] method instead.
type TuiService struct {
	Options  []option.RequestOption
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
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Body      param.Field[TuiPublishBodyUnion] `json:"body"`
}

func (r TuiPublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodyUnion interface {
	implementsTuiPublishBody()
}

type TuiPublishBodyPromptAppend struct {
	Type       param.Field[TuiPublishBodyPromptAppendType]       `json:"type,required"`
	Properties param.Field[TuiPublishBodyPromptAppendProperties] `json:"properties,required"`
}

func (r TuiPublishBodyPromptAppend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodyPromptAppend) implementsTuiPublishBody() {}

type TuiPublishBodyPromptAppendType string

const (
	TuiPublishBodyPromptAppendTypeTuiPromptAppend TuiPublishBodyPromptAppendType = "tui.prompt.append"
)

func (r TuiPublishBodyPromptAppendType) IsKnown() bool {
	switch r {
	case TuiPublishBodyPromptAppendTypeTuiPromptAppend:
		return true
	}
	return false
}

type TuiPublishBodyPromptAppendProperties struct {
	Text param.Field[string] `json:"text,required"`
}

func (r TuiPublishBodyPromptAppendProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodyCommandExecute struct {
	Type       param.Field[TuiPublishBodyCommandExecuteType]       `json:"type,required"`
	Properties param.Field[TuiPublishBodyCommandExecuteProperties] `json:"properties,required"`
}

func (r TuiPublishBodyCommandExecute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodyCommandExecute) implementsTuiPublishBody() {}

type TuiPublishBodyCommandExecuteType string

const (
	TuiPublishBodyCommandExecuteTypeTuiCommandExecute TuiPublishBodyCommandExecuteType = "tui.command.execute"
)

func (r TuiPublishBodyCommandExecuteType) IsKnown() bool {
	switch r {
	case TuiPublishBodyCommandExecuteTypeTuiCommandExecute:
		return true
	}
	return false
}

type TuiPublishBodyCommandExecuteProperties struct {
	Command param.Field[string] `json:"command,required"`
}

func (r TuiPublishBodyCommandExecuteProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodyToastShow struct {
	Type       param.Field[TuiPublishBodyToastShowType]       `json:"type,required"`
	Properties param.Field[TuiPublishBodyToastShowProperties] `json:"properties,required"`
}

func (r TuiPublishBodyToastShow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodyToastShow) implementsTuiPublishBody() {}

type TuiPublishBodyToastShowType string

const (
	TuiPublishBodyToastShowTypeTuiToastShow TuiPublishBodyToastShowType = "tui.toast.show"
)

func (r TuiPublishBodyToastShowType) IsKnown() bool {
	switch r {
	case TuiPublishBodyToastShowTypeTuiToastShow:
		return true
	}
	return false
}

type TuiPublishBodyToastShowProperties struct {
	Message  param.Field[string]                              `json:"message,required"`
	Variant  param.Field[TuiPublishBodyToastShowPropertiesVariant] `json:"variant,required"`
	Title    param.Field[string]                              `json:"title"`
	Duration param.Field[int64]                               `json:"duration"`
}

func (r TuiPublishBodyToastShowProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TuiPublishBodyToastShowPropertiesVariant string

const (
	TuiPublishBodyToastShowPropertiesVariantInfo    TuiPublishBodyToastShowPropertiesVariant = "info"
	TuiPublishBodyToastShowPropertiesVariantSuccess TuiPublishBodyToastShowPropertiesVariant = "success"
	TuiPublishBodyToastShowPropertiesVariantWarning TuiPublishBodyToastShowPropertiesVariant = "warning"
	TuiPublishBodyToastShowPropertiesVariantError   TuiPublishBodyToastShowPropertiesVariant = "error"
)

func (r TuiPublishBodyToastShowPropertiesVariant) IsKnown() bool {
	switch r {
	case TuiPublishBodyToastShowPropertiesVariantInfo, TuiPublishBodyToastShowPropertiesVariantSuccess, TuiPublishBodyToastShowPropertiesVariantWarning, TuiPublishBodyToastShowPropertiesVariantError:
		return true
	}
	return false
}

type TuiPublishBodySessionSelect struct {
	Type       param.Field[TuiPublishBodySessionSelectType]       `json:"type,required"`
	Properties param.Field[TuiPublishBodySessionSelectProperties] `json:"properties,required"`
}

func (r TuiPublishBodySessionSelect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TuiPublishBodySessionSelect) implementsTuiPublishBody() {}

type TuiPublishBodySessionSelectType string

const (
	TuiPublishBodySessionSelectTypeTuiSessionSelect TuiPublishBodySessionSelectType = "tui.session.select"
)

func (r TuiPublishBodySessionSelectType) IsKnown() bool {
	switch r {
	case TuiPublishBodySessionSelectTypeTuiSessionSelect:
		return true
	}
	return false
}

type TuiPublishBodySessionSelectProperties struct {
	SessionID param.Field[string] `json:"sessionID,required"`
}

func (r TuiPublishBodySessionSelectProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r *TuiPublishParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TuiSelectSessionParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	SessionID param.Field[string] `json:"sessionID,required"`
}

func (r TuiSelectSessionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Path string       `json:"path,required"`
	Body interface{} `json:"body,required"`
	JSON tuicontrolNextResponseJSON `json:"-"`
}

type tuicontrolNextResponseJSON struct {
	Path        apijson.Field
	Body        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TuiControlNextResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tuicontrolNextResponseJSON) RawJSON() string {
	return r.raw
}

type TuiControlResponseParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Body      param.Field[interface{}] `json:"body"`
}

func (r TuiControlResponseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
