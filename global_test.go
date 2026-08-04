// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestEventOptionalObjectFieldsUseInterfaces(t *testing.T) {
	tests := []struct {
		name  string
		typ   reflect.Type
		field string
	}{
		{
			name:  "permission asked tool",
			typ:   reflect.TypeFor[opencode.EventListResponseEventPermissionAskedProperties](),
			field: "Tool",
		},
		{
			name:  "question asked tool",
			typ:   reflect.TypeFor[opencode.EventListResponseEventQuestionAskedProperties](),
			field: "Tool",
		},
		{
			name:  "permission v2 asked source",
			typ:   reflect.TypeFor[opencode.EventListResponseEventPermissionV2AskedProperties](),
			field: "Source",
		},
		{
			name:  "question v2 asked tool",
			typ:   reflect.TypeFor[opencode.EventListResponseEventQuestionV2AskedProperties](),
			field: "Tool",
		},
	}

	interfaceType := reflect.TypeFor[any]()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field, ok := tt.typ.FieldByName(tt.field)
			if !ok {
				t.Fatalf("missing %s field", tt.field)
			}
			if field.Type != interfaceType {
				t.Fatalf("%s field type = %s, want any", tt.field, field.Type)
			}

			jsonField, ok := tt.typ.FieldByName("JSON")
			if !ok {
				t.Fatal("missing JSON metadata field")
			}
			// The metadata must still include the field so the apijson
			// framework can track it; the field is just typed as
			// any at runtime instead of a concrete pointer.
			jsonFieldEntry, ok := jsonField.Type.FieldByName(tt.field)
			if !ok {
				t.Fatalf("JSON metadata must contain %s field", tt.field)
			}
			if jsonFieldEntry.Type != reflect.TypeFor[apijson.Field]() {
				t.Fatalf("JSON metadata field %s type = %s, want apijson.Field", tt.field, jsonFieldEntry.Type)
			}
		})
	}
}

// TestEventOptionalObjectFieldsRoundtrip verifies that the apijson.Field
// metadata for Tool/Source is wired up correctly: omitting these optional
// fields from JSON must not cause an error (the apijson framework should
// tolerate their absence because the metadata is still tracked).
func TestEventOptionalObjectFieldsRoundtrip(t *testing.T) {
	// Each payload omits the optional Tool/Source field, exercising the
	// path where the apijson framework looks up the metadata entry and
	// simply finds it absent rather than treating absence as a structural
	// violation of a now-removed field.
	cases := []struct {
		name string
		data string
		into func() any
	}{
		{
			name: "permission asked without tool",
			data: `{"always":[],"id":"x","metadata":{},"patterns":[],"permission":"p","sessionID":"s"}`,
			into: func() any { return new(opencode.EventListResponseEventPermissionAskedProperties) },
		},
		{
			name: "question asked without tool",
			data: `{"id":"x","questions":[],"sessionID":"s"}`,
			into: func() any { return new(opencode.EventListResponseEventQuestionAskedProperties) },
		},
		{
			name: "permission v2 asked without source",
			data: `{"id":"x","sessionID":"s","action":"a","resources":[]}`,
			into: func() any { return new(opencode.EventListResponseEventPermissionV2AskedProperties) },
		},
		{
			name: "question v2 asked without tool",
			data: `{"id":"x","sessionID":"s","questions":[]}`,
			into: func() any { return new(opencode.EventListResponseEventQuestionV2AskedProperties) },
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tc.data), tc.into()); err != nil {
				t.Fatalf("unexpected unmarshal error: %v", err)
			}
		})
	}
}
func TestGlobalHealth(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Global.Health(context.TODO())
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalDispose(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Global.Dispose(context.TODO())
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalUpgrade(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Global.Upgrade(context.TODO(), opencode.GlobalUpgradeBody{
		Target: opencode.F("version"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestGlobalUpgradeResponseUnmarshal verifies that GlobalUpgradeResponse
// correctly deserialises both OpenAPI anyOf variants:
//
//	anyOf[{success:true, version:string}, {success:false, error:string}]
//
// and that AsUnion() returns the correct concrete variant type, while
// RawJSON() preserves the original JSON bytes.
//
// Run with: go test -run TestGlobalUpgradeResponse -v ./...
func TestGlobalUpgradeResponseUnmarshal(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		json          string
		wantSuccess   bool
		wantVersion   string
		wantError     string
		wantUnionType reflect.Type
	}{
		{
			name:          "success variant — success=true with version",
			json:          `{"success":true,"version":"1.2.3"}`,
			wantSuccess:   true,
			wantVersion:   "1.2.3",
			wantUnionType: reflect.TypeFor[opencode.GlobalUpgradeSuccess](),
		},
		{
			name:          "error variant — success=false with error",
			json:          `{"success":false,"error":"upgrade failed: network timeout"}`,
			wantSuccess:   false,
			wantError:     "upgrade failed: network timeout",
			wantUnionType: reflect.TypeFor[opencode.GlobalUpgradeError](),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var resp opencode.GlobalUpgradeResponse
			if err := json.Unmarshal([]byte(tc.json), &resp); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}

			// Verify top-level fields are ported correctly.
			if resp.Success != tc.wantSuccess {
				t.Errorf("Success = %v, want %v", resp.Success, tc.wantSuccess)
			}
			if resp.Version != tc.wantVersion {
				t.Errorf("Version = %q, want %q", resp.Version, tc.wantVersion)
			}
			if resp.Error != tc.wantError {
				t.Errorf("Error = %q, want %q", resp.Error, tc.wantError)
			}

			// Verify JSON metadata (apijson.Field) is populated (field present, not missing).
			if resp.JSON.Success.IsMissing() {
				t.Error("JSON.Success.IsMissing() = true, want false (field must be present)")
			}

			// Verify RawJSON() returns the original JSON.
			if got := resp.JSON.RawJSON(); got != tc.json {
				t.Errorf("RawJSON() = %q, want %q", got, tc.json)
			}

			// Verify AsUnion() returns the correct concrete variant type.
			union := resp.AsUnion()
			if union == nil {
				t.Fatal("AsUnion() = nil, want non-nil")
			}
			gotType := reflect.TypeOf(union)
			if gotType != tc.wantUnionType {
				t.Errorf("AsUnion() runtime type = %v, want %v", gotType, tc.wantUnionType)
			}

			// Verify the concrete variant fields.
			switch tc.wantUnionType {
			case reflect.TypeFor[opencode.GlobalUpgradeSuccess]():
				s, ok := union.(opencode.GlobalUpgradeSuccess)
				if !ok {
					t.Fatalf("AsUnion() cannot be asserted to GlobalUpgradeSuccess, got %T", union)
				}
				if s.Success != true {
					t.Errorf("GlobalUpgradeSuccess.Success = %v, want true", s.Success)
				}
				if s.Version != tc.wantVersion {
					t.Errorf("GlobalUpgradeSuccess.Version = %q, want %q", s.Version, tc.wantVersion)
				}
				if s.JSON.RawJSON() != tc.json {
					t.Errorf("GlobalUpgradeSuccess.JSON.RawJSON() = %q, want %q", s.JSON.RawJSON(), tc.json)
				}
			case reflect.TypeFor[opencode.GlobalUpgradeError]():
				e, ok := union.(opencode.GlobalUpgradeError)
				if !ok {
					t.Fatalf("AsUnion() cannot be asserted to GlobalUpgradeError, got %T", union)
				}
				if e.Success != false {
					t.Errorf("GlobalUpgradeError.Success = %v, want false", e.Success)
				}
				if e.Error != tc.wantError {
					t.Errorf("GlobalUpgradeError.Error = %q, want %q", e.Error, tc.wantError)
				}
				if e.JSON.RawJSON() != tc.json {
					t.Errorf("GlobalUpgradeError.JSON.RawJSON() = %q, want %q", e.JSON.RawJSON(), tc.json)
				}
			}
		})
	}
}
