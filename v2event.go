// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"slices"

	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
	"github.com/sst/opencode-sdk-go/packages/ssestream"
)

// V2EventService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2EventService] method instead.
type V2EventService struct {
	Options []option.RequestOption
}

// NewV2EventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2EventService(opts ...option.RequestOption) (r *V2EventService) {
	r = &V2EventService{}
	r.Options = opts
	return
}

// Subscribe to native event payloads for the server.
func (r *V2EventService) Subscribe(ctx context.Context, query EventListParams, opts ...option.RequestOption) (stream *ssestream.Stream[EventListResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "api/event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &raw, opts...)
	return ssestream.NewStream[EventListResponse](ssestream.NewDecoder(raw), err)
}
