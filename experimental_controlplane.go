// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// ExperimentalControlPlaneService contains methods and other services that help
// with interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalControlPlaneService] method instead.
type ExperimentalControlPlaneService struct {
	Options []option.RequestOption
}

// NewExperimentalControlPlaneService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExperimentalControlPlaneService(opts ...option.RequestOption) (r *ExperimentalControlPlaneService) {
	r = &ExperimentalControlPlaneService{}
	r.Options = opts
	return
}

// Move session
//
// Move a session to another project directory, optionally transferring local changes.
func (r *ExperimentalControlPlaneService) MoveSession(ctx context.Context, body ExperimentalControlPlaneMoveSessionParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	path := "experimental/control-plane/move-session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// MoveSessionDestination represents the destination for a session move.
type MoveSessionDestination struct {
	Directory param.Field[string] `json:"directory,required"`
}

func (r MoveSessionDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExperimentalControlPlaneMoveSessionParams struct {
	SessionID   param.Field[string]                 `json:"sessionID,required"`
	Destination param.Field[MoveSessionDestination] `json:"destination,required"`
	MoveChanges param.Field[bool]                   `json:"moveChanges"`
}

func (r ExperimentalControlPlaneMoveSessionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
