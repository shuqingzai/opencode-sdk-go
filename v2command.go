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

// V2CommandService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2CommandService] method instead.
type V2CommandService struct {
	Options []option.RequestOption
}

// NewV2CommandService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2CommandService(opts ...option.RequestOption) (r *V2CommandService) {
	r = &V2CommandService{}
	r.Options = opts
	return
}

// List v2 commands
func (r *V2CommandService) List(ctx context.Context, query V2CommandListParams, opts ...option.RequestOption) (res *[]V2CommandInfo, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/command"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// V2CommandInfo represents a registered command.
type V2CommandInfo struct {
	Name        string              `json:"name,required"`
	Template    string              `json:"template,required"`
	Description string              `json:"description"`
	Agent       string              `json:"agent"`
	Model       ModelRef            `json:"model"`
	Subtask     bool                `json:"subtask"`
	JSON        v2CommandInfoJSON   `json:"-"`
}

// v2CommandInfoJSON contains the JSON metadata for the struct [V2CommandInfo]
type v2CommandInfoJSON struct {
	Name        apijson.Field
	Template    apijson.Field
	Description apijson.Field
	Agent       apijson.Field
	Model       apijson.Field
	Subtask     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2CommandInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2CommandInfoJSON) RawJSON() string {
	return r.raw
}

type V2CommandListParams struct {
	Location param.Field[V2LocationParam] `query:"location"`
}

func (r V2CommandListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
