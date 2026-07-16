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
			typ:   reflect.TypeOf(opencode.EventListResponseEventPermissionAskedProperties{}),
			field: "Tool",
		},
		{
			name:  "question asked tool",
			typ:   reflect.TypeOf(opencode.EventListResponseEventQuestionAskedProperties{}),
			field: "Tool",
		},
		{
			name:  "permission v2 asked source",
			typ:   reflect.TypeOf(opencode.EventListResponseEventPermissionV2AskedProperties{}),
			field: "Source",
		},
		{
			name:  "question v2 asked tool",
			typ:   reflect.TypeOf(opencode.EventListResponseEventQuestionV2AskedProperties{}),
			field: "Tool",
		},
	}

	interfaceType := reflect.TypeOf((*interface{})(nil)).Elem()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field, ok := tt.typ.FieldByName(tt.field)
			if !ok {
				t.Fatalf("missing %s field", tt.field)
			}
			if field.Type != interfaceType {
				t.Fatalf("%s field type = %s, want interface{}", tt.field, field.Type)
			}

			jsonField, ok := tt.typ.FieldByName("JSON")
			if !ok {
				t.Fatal("missing JSON metadata field")
			}
			// The metadata must still include the field so the apijson
			// framework can track it; the field is just typed as
			// interface{} at runtime instead of a concrete pointer.
			jsonFieldEntry, ok := jsonField.Type.FieldByName(tt.field)
			if !ok {
				t.Fatalf("JSON metadata must contain %s field", tt.field)
			}
			if jsonFieldEntry.Type != reflect.TypeOf(apijson.Field{}) {
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
