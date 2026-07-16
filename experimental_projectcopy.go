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

// ExperimentalProjectCopyService contains methods and other services that help
// with interacting with the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentalProjectCopyService] method instead.
type ExperimentalProjectCopyService struct {
	Options []option.RequestOption
}

// NewExperimentalProjectCopyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExperimentalProjectCopyService(opts ...option.RequestOption) (r *ExperimentalProjectCopyService) {
	r = &ExperimentalProjectCopyService{}
	r.Options = opts
	return
}

// Generate project copy name
//
// Generate a short name for a project copy from task context.
func (r *ExperimentalProjectCopyService) GenerateName(ctx context.Context, projectID string, params ExperimentalProjectCopyGenerateNameParams, opts ...option.RequestOption) (res *ProjectCopyGenerateNameResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required projectID parameter")
		return
	}
	path := fmt.Sprintf("experimental/project/%s/copy/generate-name", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ProjectCopyGenerateNameResponse represents the response from generating a project copy name.
type ProjectCopyGenerateNameResponse struct {
	Name string                              `json:"name,required"`
	JSON projectCopyGenerateNameResponseJSON `json:"-"`
}

// projectCopyGenerateNameResponseJSON contains the JSON metadata for the struct
// [ProjectCopyGenerateNameResponse]
type projectCopyGenerateNameResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectCopyGenerateNameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectCopyGenerateNameResponseJSON) RawJSON() string {
	return r.raw
}

type ExperimentalProjectCopyGenerateNameParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	Context   param.Field[string] `json:"context"`
}

func (r ExperimentalProjectCopyGenerateNameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExperimentalProjectCopyGenerateNameParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
