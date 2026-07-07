// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// V2HealthService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2HealthService] method instead.
type V2HealthService struct {
	Options []option.RequestOption
}

// NewV2HealthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2HealthService(opts ...option.RequestOption) (r *V2HealthService) {
	r = &V2HealthService{}
	r.Options = opts
	return
}

// Check server health
//
// Check whether the API server is ready to accept requests.
func (r *V2HealthService) Get(ctx context.Context, opts ...option.RequestOption) (res *HealthV2Info, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthV2Info struct {
	Healthy bool             `json:"healthy,required"`
	JSON    healthV2InfoJSON `json:"-"`
}

// healthV2InfoJSON contains the JSON metadata for the struct [HealthV2Info]
type healthV2InfoJSON struct {
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthV2Info) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthV2InfoJSON) RawJSON() string {
	return r.raw
}
