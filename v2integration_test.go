// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/sst/opencode-sdk-go"
)

// TestIntegrationAttemptTimeUnmarshal verifies that IntegrationAttemptTime
// correctly deserializes all four value variants defined by the OpenAPI schema:
//
//	anyOf[number, "NaN", "Infinity", "-Infinity"]
//
// At runtime, a JSON number is decoded to float64, and a JSON string is decoded
// to string. This matches the Workspace.TimeUsed pattern and the provider_test.go
// TestProviderModelCapabilitiesInterleavedUnmarshal template.
func TestIntegrationAttemptTimeUnmarshal(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		jsonStr      string
		wantCreatedT reflect.Type
		wantCreatedF float64 // checked only when wantCreatedT == float64
		wantCreatedS string  // checked only when wantCreatedT == string
		wantExpiresT reflect.Type
		wantExpiresF float64
		wantExpiresS string
	}{
		// number variant: regular Unix timestamp
		{
			name:         "number",
			jsonStr:      `{"created":1700000000.5,"expires":1700003600.0}`,
			wantCreatedT: reflect.TypeFor[float64](),
			wantCreatedF: 1700000000.5,
			wantExpiresT: reflect.TypeFor[float64](),
			wantExpiresF: 1700003600.0,
		},
		// string variant: "NaN"
		{
			name:         "NaN",
			jsonStr:      `{"created":"NaN","expires":"NaN"}`,
			wantCreatedT: reflect.TypeFor[string](),
			wantCreatedS: "NaN",
			wantExpiresT: reflect.TypeFor[string](),
			wantExpiresS: "NaN",
		},
		// string variant: "Infinity"
		{
			name:         "Infinity",
			jsonStr:      `{"created":"Infinity","expires":"Infinity"}`,
			wantCreatedT: reflect.TypeFor[string](),
			wantCreatedS: "Infinity",
			wantExpiresT: reflect.TypeFor[string](),
			wantExpiresS: "Infinity",
		},
		// string variant: "-Infinity"
		{
			name:         "NegInfinity",
			jsonStr:      `{"created":"-Infinity","expires":"-Infinity"}`,
			wantCreatedT: reflect.TypeFor[string](),
			wantCreatedS: "-Infinity",
			wantExpiresT: reflect.TypeFor[string](),
			wantExpiresS: "-Infinity",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var at opencode.IntegrationAttemptTime
			if err := json.Unmarshal([]byte(tc.jsonStr), &at); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}

			// Check Created runtime type
			gotCreatedT := reflect.TypeOf(at.Created)
			if gotCreatedT != tc.wantCreatedT {
				t.Errorf("Created runtime type: got %v, want %v", gotCreatedT, tc.wantCreatedT)
			}
			switch tc.wantCreatedT {
			case reflect.TypeFor[float64]():
				v, ok := at.Created.(float64)
				if !ok {
					t.Fatalf("expected Created to be float64, got %T", at.Created)
				}
				if v != tc.wantCreatedF {
					t.Errorf("Created float64 value: got %v, want %v", v, tc.wantCreatedF)
				}
			case reflect.TypeFor[string]():
				s, ok := at.Created.(string)
				if !ok {
					t.Fatalf("expected Created to be string, got %T", at.Created)
				}
				if s != tc.wantCreatedS {
					t.Errorf("Created string value: got %q, want %q", s, tc.wantCreatedS)
				}
			}

			// Check Expires runtime type
			gotExpiresT := reflect.TypeOf(at.Expires)
			if gotExpiresT != tc.wantExpiresT {
				t.Errorf("Expires runtime type: got %v, want %v", gotExpiresT, tc.wantExpiresT)
			}
			switch tc.wantExpiresT {
			case reflect.TypeFor[float64]():
				v, ok := at.Expires.(float64)
				if !ok {
					t.Fatalf("expected Expires to be float64, got %T", at.Expires)
				}
				if v != tc.wantExpiresF {
					t.Errorf("Expires float64 value: got %v, want %v", v, tc.wantExpiresF)
				}
			case reflect.TypeFor[string]():
				s, ok := at.Expires.(string)
				if !ok {
					t.Fatalf("expected Expires to be string, got %T", at.Expires)
				}
				if s != tc.wantExpiresS {
					t.Errorf("Expires string value: got %q, want %q", s, tc.wantExpiresS)
				}
			}
		})
	}
}
