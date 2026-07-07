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

// V2SessionRevertService contains methods and other services that help with
// interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2SessionRevertService] method instead.
type V2SessionRevertService struct {
	Options []option.RequestOption
}

// NewV2SessionRevertService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2SessionRevertService(opts ...option.RequestOption) (r *V2SessionRevertService) {
	r = &V2SessionRevertService{}
	r.Options = opts
	return
}

// Stage session revert
//
// Stage or move a reversible session boundary and optionally apply its file
// changes.
func (r *V2SessionRevertService) Stage(ctx context.Context, sessionID string, body V2SessionRevertStageParams, opts ...option.RequestOption) (res *V2SessionRevertStageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/revert/stage", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Clear staged revert
func (r *V2SessionRevertService) Clear(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/revert/clear", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Commit staged revert
func (r *V2SessionRevertService) Commit(ctx context.Context, sessionID string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	path := fmt.Sprintf("api/session/%s/revert/commit", sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// ===== Param Types =====

type V2SessionRevertStageParams struct {
	MessageID param.Field[string] `json:"messageID,required"`
	Files     param.Field[bool]   `json:"files"`
}

func (r V2SessionRevertStageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ===== Response Types =====

// V2SessionRevertStageResponse is returned by the Revert.Stage method.
type V2SessionRevertStageResponse struct {
	Data RevertState                          `json:"data,required"`
	JSON v2SessionRevertStageResponseJSON     `json:"-"`
}

type v2SessionRevertStageResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2SessionRevertStageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2SessionRevertStageResponseJSON) RawJSON() string {
	return r.raw
}
