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

// PartService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPartService] method instead.
type PartService struct {
	Options []option.RequestOption
}

// NewPartService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPartService(opts ...option.RequestOption) (r *PartService) {
	r = &PartService{}
	r.Options = opts
	return
}

// Delete a part from a message
func (r *PartService) Delete(ctx context.Context, sessionID string, messageID string, partID string, body PartDeleteParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	if partID == "" {
		err = errors.New("missing required partID parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message/%s/part/%s", sessionID, messageID, partID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Update a part in a message
func (r *PartService) Update(ctx context.Context, sessionID string, messageID string, partID string, params PartUpdateParams, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.Options, opts)
	if sessionID == "" {
		err = errors.New("missing required sessionID parameter")
		return
	}
	if messageID == "" {
		err = errors.New("missing required messageID parameter")
		return
	}
	if partID == "" {
		err = errors.New("missing required partID parameter")
		return
	}
	path := fmt.Sprintf("session/%s/message/%s/part/%s", sessionID, messageID, partID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type PartDeleteParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
}

func (r PartDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PartUpdateParams struct {
	Directory param.Field[string] `query:"directory"`
	Workspace param.Field[string] `query:"workspace"`
	// Part is the request body — the patch body itself is a Part object
	// (OpenAPI: $ref: Part), so this field's value is serialized as the body root.
	// When Part is not set, an empty object is sent (JS SDK treats `part` as optional).
	Part param.Field[Part] `json:"part"`
}

func (r PartUpdateParams) MarshalJSON() (data []byte, err error) {
	if r.Part.Present {
		return apijson.MarshalRoot(r.Part)
	}
	return []byte("{}"), nil
}

func (r PartUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
