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

// PermissionService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPermissionService] method instead.
type PermissionService struct {
	Options []option.RequestOption
}

// NewPermissionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPermissionService(opts ...option.RequestOption) (r *PermissionService) {
	r = &PermissionService{}
	r.Options = opts
	return
}

// List pending permission requests
func (r *PermissionService) List(ctx context.Context, query PermissionListParams, opts ...option.RequestOption) (res *[]PermissionRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "permission"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Reply to a permission request
func (r *PermissionService) Reply(ctx context.Context, requestID string, params PermissionReplyParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return
	}
	path := fmt.Sprintf("permission/%s/reply", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Deprecated: Use [PermissionService.Reply] instead.
//
// Respond to a permission request (deprecated)
func (r *PermissionService) Respond(ctx context.Context, sessionID string, permissionID string, params PermissionRespondParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if permissionID == "" {
		err = errors.New("missing required permissionID parameter")
		return
	}
	path := fmt.Sprintf("session/%s/permissions/%s", sessionID, permissionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PermissionRequest struct {
	ID         string   `json:"id,required"`
	SessionID  string   `json:"sessionID,required"`
	Permission string   `json:"permission,required"`
	Patterns   []string `json:"patterns,required"`
	// This field can have the runtime type of [map[string]any].
	Metadata any                   `json:"metadata,required"`
	Always   []string              `json:"always,required"`
	Tool     PermissionRequestTool `json:"tool"`
	JSON     permissionRequestJSON `json:"-"`
}

// permissionRequestJSON contains the JSON metadata for the struct
// [PermissionRequest]
type permissionRequestJSON struct {
	ID          apijson.Field
	SessionID   apijson.Field
	Permission  apijson.Field
	Patterns    apijson.Field
	Metadata    apijson.Field
	Always      apijson.Field
	Tool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionRequestJSON) RawJSON() string {
	return r.raw
}

type PermissionRequestTool struct {
	MessageID string                    `json:"messageID,required"`
	CallID    string                    `json:"callID,required"`
	JSON      permissionRequestToolJSON `json:"-"`
}

// permissionRequestToolJSON contains the JSON metadata for the struct
// [PermissionRequestTool]
type permissionRequestToolJSON struct {
	MessageID   apijson.Field
	CallID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionRequestTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionRequestToolJSON) RawJSON() string {
	return r.raw
}

// Reply to a permission request
type PermissionReplyParams struct {
	// Reply type: "once", "always", or "reject"
	Reply param.Field[PermissionReplyParamsReply] `json:"reply,required"`
	// Optional message to include with the reply
	Message   param.Field[string] `json:"message"`
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r PermissionReplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [PermissionReplyParams]'s query parameters as `url.Values`.
func (r PermissionReplyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Deprecated: Use [PermissionReplyParams] instead.
//
// Respond to a permission request (deprecated)
type PermissionRespondParams struct {
	// Response type: "once", "always", or "reject"
	Response  param.Field[PermissionRespondParamsResponse] `json:"response,required"`
	Directory param.Field[string]                          `query:"directory"`
	Workspace param.Field[string]                          `query:"workspace"`
}

func (r PermissionRespondParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [PermissionRespondParams]'s query parameters as `url.Values`.
func (r PermissionRespondParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Deprecated: Use [PermissionReplyParamsReply] instead.
//
// Response type for permission respond
type PermissionRespondParamsResponse string

const (
	PermissionRespondParamsResponseOnce   PermissionRespondParamsResponse = "once"
	PermissionRespondParamsResponseAlways PermissionRespondParamsResponse = "always"
	PermissionRespondParamsResponseReject PermissionRespondParamsResponse = "reject"
)

func (r PermissionRespondParamsResponse) IsKnown() bool {
	switch r {
	case PermissionRespondParamsResponseOnce, PermissionRespondParamsResponseAlways, PermissionRespondParamsResponseReject:
		return true
	}
	return false
}

type PermissionListParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [PermissionListParams]'s query parameters as `url.Values`.
func (r PermissionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Reply to a permission request
type PermissionReplyParamsReply string

const (
	PermissionReplyParamsReplyOnce   PermissionReplyParamsReply = "once"
	PermissionReplyParamsReplyAlways PermissionReplyParamsReply = "always"
	PermissionReplyParamsReplyReject PermissionReplyParamsReply = "reject"
)

func (r PermissionReplyParamsReply) IsKnown() bool {
	switch r {
	case PermissionReplyParamsReplyOnce, PermissionReplyParamsReplyAlways, PermissionReplyParamsReplyReject:
		return true
	}
	return false
}
