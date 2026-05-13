// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/sst/opencode-sdk-go/internal/apijson"
)

type MessageAbortedError struct {
	Data MessageAbortedErrorData `json:"data,required"`
	Name MessageAbortedErrorName `json:"name,required"`
	JSON messageAbortedErrorJSON `json:"-"`
}

// messageAbortedErrorJSON contains the JSON metadata for the struct
// [MessageAbortedError]
type messageAbortedErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAbortedError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAbortedErrorJSON) RawJSON() string {
	return r.raw
}

func (r MessageAbortedError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r MessageAbortedError) ImplementsAssistantMessageError() {}

type MessageAbortedErrorData struct {
	Message string                      `json:"message,required"`
	JSON    messageAbortedErrorDataJSON `json:"-"`
}

// messageAbortedErrorDataJSON contains the JSON metadata for the struct
// [MessageAbortedErrorData]
type messageAbortedErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageAbortedErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageAbortedErrorDataJSON) RawJSON() string {
	return r.raw
}

type MessageAbortedErrorName string

const (
	MessageAbortedErrorNameMessageAbortedError MessageAbortedErrorName = "MessageAbortedError"
)

func (r MessageAbortedErrorName) IsKnown() bool {
	switch r {
	case MessageAbortedErrorNameMessageAbortedError:
		return true
	}
	return false
}

type ProviderAuthError struct {
	Data ProviderAuthErrorData `json:"data,required"`
	Name ProviderAuthErrorName `json:"name,required"`
	JSON providerAuthErrorJSON `json:"-"`
}

// providerAuthErrorJSON contains the JSON metadata for the struct
// [ProviderAuthError]
type providerAuthErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthErrorJSON) RawJSON() string {
	return r.raw
}

func (r ProviderAuthError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r ProviderAuthError) ImplementsAssistantMessageError() {}

type ProviderAuthErrorData struct {
	Message    string                    `json:"message,required"`
	ProviderID string                    `json:"providerID,required"`
	JSON       providerAuthErrorDataJSON `json:"-"`
}

// providerAuthErrorDataJSON contains the JSON metadata for the struct
// [ProviderAuthErrorData]
type providerAuthErrorDataJSON struct {
	Message     apijson.Field
	ProviderID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderAuthErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerAuthErrorDataJSON) RawJSON() string {
	return r.raw
}

type ProviderAuthErrorName string

const (
	ProviderAuthErrorNameProviderAuthError ProviderAuthErrorName = "ProviderAuthError"
)

func (r ProviderAuthErrorName) IsKnown() bool {
	switch r {
	case ProviderAuthErrorNameProviderAuthError:
		return true
	}
	return false
}

type UnknownError struct {
	Data UnknownErrorData `json:"data,required"`
	Name UnknownErrorName `json:"name,required"`
	JSON unknownErrorJSON `json:"-"`
}

// unknownErrorJSON contains the JSON metadata for the struct [UnknownError]
type unknownErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnknownError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unknownErrorJSON) RawJSON() string {
	return r.raw
}

func (r UnknownError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r UnknownError) ImplementsAssistantMessageError() {}

type UnknownErrorData struct {
	Message string               `json:"message,required"`
	JSON    unknownErrorDataJSON `json:"-"`
}

// unknownErrorDataJSON contains the JSON metadata for the struct
// [UnknownErrorData]
type unknownErrorDataJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnknownErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unknownErrorDataJSON) RawJSON() string {
	return r.raw
}

type UnknownErrorName string

const (
	UnknownErrorNameUnknownError UnknownErrorName = "UnknownError"
)

func (r UnknownErrorName) IsKnown() bool {
	switch r {
	case UnknownErrorNameUnknownError:
		return true
	}
	return false
}

type StructuredOutputError struct {
	Data StructuredOutputErrorData `json:"data,required"`
	Name StructuredOutputErrorName `json:"name,required"`
	JSON structuredOutputErrorJSON `json:"-"`
}

// structuredOutputErrorJSON contains the JSON metadata for the struct
// [StructuredOutputError]
type structuredOutputErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StructuredOutputError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r structuredOutputErrorJSON) RawJSON() string {
	return r.raw
}

func (r StructuredOutputError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r StructuredOutputError) ImplementsAssistantMessageError() {}

type StructuredOutputErrorData struct {
	Message string                        `json:"message,required"`
	Retries int64                         `json:"retries,required"`
	JSON    structuredOutputErrorDataJSON `json:"-"`
}

// structuredOutputErrorDataJSON contains the JSON metadata for the struct
// [StructuredOutputErrorData]
type structuredOutputErrorDataJSON struct {
	Message     apijson.Field
	Retries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StructuredOutputErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r structuredOutputErrorDataJSON) RawJSON() string {
	return r.raw
}

type StructuredOutputErrorName string

const (
	StructuredOutputErrorNameStructuredOutputError StructuredOutputErrorName = "StructuredOutputError"
)

func (r StructuredOutputErrorName) IsKnown() bool {
	switch r {
	case StructuredOutputErrorNameStructuredOutputError:
		return true
	}
	return false
}

type ContextOverflowError struct {
	Data ContextOverflowErrorData `json:"data,required"`
	Name ContextOverflowErrorName `json:"name,required"`
	JSON contextOverflowErrorJSON `json:"-"`
}

// contextOverflowErrorJSON contains the JSON metadata for the struct
// [ContextOverflowError]
type contextOverflowErrorJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextOverflowError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextOverflowErrorJSON) RawJSON() string {
	return r.raw
}

func (r ContextOverflowError) ImplementsEventListResponseEventSessionErrorPropertiesError() {}

func (r ContextOverflowError) ImplementsAssistantMessageError() {}

type ContextOverflowErrorData struct {
	Message      string                      `json:"message,required"`
	ResponseBody string                      `json:"responseBody"`
	JSON         contextOverflowErrorDataJSON `json:"-"`
}

// contextOverflowErrorDataJSON contains the JSON metadata for the struct
// [ContextOverflowErrorData]
type contextOverflowErrorDataJSON struct {
	Message      apijson.Field
	ResponseBody apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContextOverflowErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextOverflowErrorDataJSON) RawJSON() string {
	return r.raw
}

type ContextOverflowErrorName string

const (
	ContextOverflowErrorNameContextOverflowError ContextOverflowErrorName = "ContextOverflowError"
)

func (r ContextOverflowErrorName) IsKnown() bool {
	switch r {
	case ContextOverflowErrorNameContextOverflowError:
		return true
	}
	return false
}
