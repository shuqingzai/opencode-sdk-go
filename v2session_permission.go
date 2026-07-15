// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// V2SessionPermissionService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2SessionPermissionService] method instead.
type V2SessionPermissionService struct {
	Options []option.RequestOption
}

// NewV2SessionPermissionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2SessionPermissionService(opts ...option.RequestOption) (r *V2SessionPermissionService) {
	r = &V2SessionPermissionService{}
	r.Options = opts
	return
}

// List session permission requests
//
// Retrieve pending permission requests owned by a session.
func (r *V2SessionPermissionService) List(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *V2SessionPermissionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/permission", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create permission request
//
// Evaluate and, when approval is required, create a permission request for a
// session.
func (r *V2SessionPermissionService) Create(ctx context.Context, sessionID string, body V2SessionPermissionCreateParams, opts ...option.RequestOption) (res *V2SessionPermissionCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/permission", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get permission request
//
// Retrieve a pending permission request owned by a session.
func (r *V2SessionPermissionService) Get(ctx context.Context, sessionID string, requestID string, opts ...option.RequestOption) (res *V2SessionPermissionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/permission/%s", sessionID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Reply to pending permission request
//
// Respond to a pending permission request owned by a session.
func (r *V2SessionPermissionService) Reply(ctx context.Context, sessionID string, requestID string, body V2SessionPermissionReplyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required requestID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/permission/%s/reply", sessionID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// ===== Param Types =====

type V2SessionPermissionCreateParams struct {
	ID        param.Field[string]   `json:"id"`
	Action    param.Field[string]   `json:"action,required"`
	Resources param.Field[[]string] `json:"resources,required"`
	Save      param.Field[[]string] `json:"save"`
	// This field can have the runtime type of [map[string]interface{}].
	Metadata param.Field[interface{}]        `json:"metadata"`
	Source   param.Field[PermissionV2Source] `json:"source"`
	Agent    param.Field[string]             `json:"agent"`
}

func (r V2SessionPermissionCreateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V2SessionPermissionReplyParams struct {
	Reply   param.Field[PermissionV2Reply] `json:"reply,required"`
	Message param.Field[string]            `json:"message"`
}

func (r V2SessionPermissionReplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// PermissionV2Reply represents the reply type for permission v2 requests.
type PermissionV2Reply string

const (
	PermissionV2ReplyOnce   PermissionV2Reply = "once"
	PermissionV2ReplyAlways PermissionV2Reply = "always"
	PermissionV2ReplyReject PermissionV2Reply = "reject"
)

func (r PermissionV2Reply) IsKnown() bool {
	switch r {
	case PermissionV2ReplyOnce, PermissionV2ReplyAlways, PermissionV2ReplyReject:
		return true
	}
	return false
}

// ===== Response Types =====

// V2SessionPermissionCreateResponse is returned by the Permission.Create method.
type V2SessionPermissionCreateResponse struct {
	Data V2SessionPermissionCreateData         `json:"data,required"`
	JSON v2SessionPermissionCreateResponseJSON `json:"-"`
}

type v2SessionPermissionCreateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionPermissionCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionPermissionCreateResponseJSON) RawJSON() string {
	return r.raw
}

type V2SessionPermissionCreateData struct {
	ID     string                            `json:"id,required"`
	Effect PermissionV2Effect                `json:"effect,required"`
	JSON   v2SessionPermissionCreateDataJSON `json:"-"`
}

type v2SessionPermissionCreateDataJSON struct {
	ID          apijson.Field
	Effect      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionPermissionCreateData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionPermissionCreateDataJSON) RawJSON() string {
	return r.raw
}

// V2SessionPermissionListResponse is returned by the Permission.List method.
type V2SessionPermissionListResponse struct {
	Data []PermissionV2Request               `json:"data,required"`
	JSON v2SessionPermissionListResponseJSON `json:"-"`
}

type v2SessionPermissionListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionPermissionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionPermissionListResponseJSON) RawJSON() string {
	return r.raw
}

// V2SessionPermissionGetResponse is returned by the Permission.Get method.
type V2SessionPermissionGetResponse struct {
	Data PermissionV2Request                `json:"data,required"`
	JSON v2SessionPermissionGetResponseJSON `json:"-"`
}

type v2SessionPermissionGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionPermissionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionPermissionGetResponseJSON) RawJSON() string {
	return r.raw
}
