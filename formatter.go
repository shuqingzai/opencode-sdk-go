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

// FormatterService contains methods and other services that help with interacting
// with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFormatterService] method instead.
type FormatterService struct {
	Options []option.RequestOption
}

// NewFormatterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFormatterService(opts ...option.RequestOption) (r *FormatterService) {
	r = &FormatterService{}
	r.Options = opts
	return
}

func (r *FormatterService) Status(ctx context.Context, query FormatterStatusParams, opts ...option.RequestOption) (res *[]FormatterStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "formatter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FormatterStatus struct {
	Name       string              `json:"name,required"`
	Extensions []string            `json:"extensions,required"`
	Enabled    bool                `json:"enabled,required"`
	JSON       formatterStatusJSON `json:"-"`
}

type formatterStatusJSON struct {
	Name        apijson.Field
	Extensions  apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FormatterStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r formatterStatusJSON) RawJSON() string {
	return r.raw
}

type FormatterStatusParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

// URLQuery serializes [FormatterStatusParams]'s query parameters as `url.Values`.
func (r FormatterStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
