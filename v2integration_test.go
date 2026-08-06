// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/sst/opencode-sdk-go"
)

// TestIntegrationOAuthMethodPromptsDecoding verifies that IntegrationOAuthMethod.Prompts
// decodes to []IntegrationPrompt carrier structs whose AsUnion() resolves to the
// concrete variant types, not []interface{}.
func TestIntegrationOAuthMethodPromptsDecoding(t *testing.T) {
	t.Parallel()

	t.Run("MixedPrompts", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{
			"id":"oauth_1",
			"type":"oauth",
			"label":"My OAuth",
			"prompts":[
				{"type":"text","key":"api_key","label":"API Key","placeholder":"Enter key","description":"Your API key","required":true},
				{"type":"select","key":"region","label":"Region","options":["us","eu"]}
			]
		}`
		var m opencode.IntegrationOAuthMethod
		if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if len(m.Prompts) != 2 {
			t.Fatalf("expected 2 prompts, got %d", len(m.Prompts))
		}
		textPrompt, ok := m.Prompts[0].AsUnion().(opencode.IntegrationTextPrompt)
		if !ok {
			t.Errorf("Prompts[0]: expected IntegrationTextPrompt, got %T", m.Prompts[0].AsUnion())
		} else if textPrompt.Key != "api_key" {
			t.Errorf("Prompts[0].Key: expected api_key, got %q", textPrompt.Key)
		}
		if m.Prompts[0].Type != opencode.IntegrationPromptTypeText {
			t.Errorf("Prompts[0].Type: got %q, want %q", m.Prompts[0].Type, opencode.IntegrationPromptTypeText)
		}
		selectPrompt, ok := m.Prompts[1].AsUnion().(opencode.IntegrationSelectPrompt)
		if !ok {
			t.Errorf("Prompts[1]: expected IntegrationSelectPrompt, got %T", m.Prompts[1].AsUnion())
		} else if selectPrompt.Key != "region" {
			t.Errorf("Prompts[1].Key: expected region, got %q", selectPrompt.Key)
		}
		if m.Prompts[1].Type != opencode.IntegrationPromptTypeSelect {
			t.Errorf("Prompts[1].Type: got %q, want %q", m.Prompts[1].Type, opencode.IntegrationPromptTypeSelect)
		}
	})

	t.Run("EmptyPrompts", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{"id":"oauth_2","type":"oauth","label":"No prompts","prompts":[]}`
		var m opencode.IntegrationOAuthMethod
		if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if len(m.Prompts) != 0 {
			t.Errorf("expected 0 prompts, got %d", len(m.Prompts))
		}
	})

	t.Run("NullPrompts", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{"id":"oauth_3","type":"oauth","label":"Null prompts","prompts":null}`
		var m opencode.IntegrationOAuthMethod
		if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if m.Prompts != nil {
			t.Errorf("expected nil prompts, got %v", m.Prompts)
		}
	})

	t.Run("AbsentPrompts", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{"id":"oauth_4","type":"oauth","label":"No prompts field"}`
		var m opencode.IntegrationOAuthMethod
		if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if m.Prompts != nil {
			t.Errorf("expected nil prompts on absent field, got %v", m.Prompts)
		}
	})
}

// TestIntegrationInfoMethodsConnectionsDecoding verifies that IntegrationInfo
// (typed as []IntegrationMethod / []ConnectionInfo carrier structs) deserialises
// heterogeneous arrays via the registered "type" discriminators, resolves each
// element's AsUnion() to the concrete variant, and populates the carrier fields.
func TestIntegrationInfoMethodsConnectionsDecoding(t *testing.T) {
	t.Parallel()

	jsonStr := `{
		"id":"github",
		"name":"GitHub",
		"methods":[
			{"id":"oauth_main","type":"oauth","label":"OAuth"},
			{"type":"key","label":"Personal token"},
			{"type":"env","names":["GITHUB_TOKEN"]}
		],
		"connections":[
			{"type":"credential","id":"cred_1","label":"user cred"},
			{"type":"env","name":"GITHUB_TOKEN"}
		]
	}`
	var info opencode.IntegrationInfo
	if err := json.Unmarshal([]byte(jsonStr), &info); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if info.ID != "github" || info.Name != "GitHub" {
		t.Errorf("ID/Name: %q / %q", info.ID, info.Name)
	}
	if len(info.Methods) != 3 {
		t.Fatalf("expected 3 methods, got %d", len(info.Methods))
	}
	if _, ok := info.Methods[0].AsUnion().(opencode.IntegrationOAuthMethod); !ok {
		t.Errorf("Methods[0]: expected IntegrationOAuthMethod, got %T", info.Methods[0].AsUnion())
	}
	if info.Methods[0].ID != "oauth_main" || info.Methods[0].Type != opencode.IntegrationMethodTypeOAuth || info.Methods[0].Label != "OAuth" {
		t.Errorf("Methods[0] carrier fields: %+v", info.Methods[0])
	}
	if _, ok := info.Methods[1].AsUnion().(opencode.IntegrationKeyMethod); !ok {
		t.Errorf("Methods[1]: expected IntegrationKeyMethod, got %T", info.Methods[1].AsUnion())
	}
	if info.Methods[1].Type != opencode.IntegrationMethodTypeKey || info.Methods[1].Label != "Personal token" {
		t.Errorf("Methods[1] carrier fields: %+v", info.Methods[1])
	}
	if _, ok := info.Methods[2].AsUnion().(opencode.IntegrationEnvMethod); !ok {
		t.Errorf("Methods[2]: expected IntegrationEnvMethod, got %T", info.Methods[2].AsUnion())
	}
	if info.Methods[2].Type != opencode.IntegrationMethodTypeEnv ||
		len(info.Methods[2].Names) != 1 || info.Methods[2].Names[0] != "GITHUB_TOKEN" {
		t.Errorf("Methods[2] carrier fields: %+v", info.Methods[2])
	}
	if len(info.Connections) != 2 {
		t.Fatalf("expected 2 connections, got %d", len(info.Connections))
	}
	if _, ok := info.Connections[0].AsUnion().(opencode.ConnectionCredentialInfo); !ok {
		t.Errorf("Connections[0]: expected ConnectionCredentialInfo, got %T", info.Connections[0].AsUnion())
	}
	if info.Connections[0].Type != opencode.ConnectionInfoTypeCredential ||
		info.Connections[0].ID != "cred_1" || info.Connections[0].Label != "user cred" {
		t.Errorf("Connections[0] carrier fields: %+v", info.Connections[0])
	}
	if _, ok := info.Connections[1].AsUnion().(opencode.ConnectionEnvInfo); !ok {
		t.Errorf("Connections[1]: expected ConnectionEnvInfo, got %T", info.Connections[1].AsUnion())
	}
	if info.Connections[1].Type != opencode.ConnectionInfoTypeEnv || info.Connections[1].Name != "GITHUB_TOKEN" {
		t.Errorf("Connections[1] carrier fields: %+v", info.Connections[1])
	}
}

// TestIntegrationAttemptStatusTypedFields verifies that IntegrationAttemptStatus
// exposes typed Time (IntegrationAttemptTime) and Message (string) fields after
// Task 3 tightened them from any.
func TestIntegrationAttemptStatusTypedFields(t *testing.T) {
	t.Parallel()

	t.Run("failed", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{"status":"failed","message":"bad code","time":{"created":1700000000.0,"expires":1700003600.0}}`
		var s opencode.IntegrationAttemptStatus
		if err := json.Unmarshal([]byte(jsonStr), &s); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if s.Status != opencode.IntegrationAttemptStatusTypeFailed {
			t.Errorf("Status: got %q, want %q", s.Status, opencode.IntegrationAttemptStatusTypeFailed)
		}
		if s.Message != "bad code" {
			t.Errorf("Message: got %q, want bad code", s.Message)
		}
		if s.Time.Created != 1700000000.0 {
			t.Errorf("Time.Created: got %v, want 1700000000.0", s.Time.Created)
		}
		if s.Time.Expires != 1700003600.0 {
			t.Errorf("Time.Expires: got %v, want 1700003600.0", s.Time.Expires)
		}
		if _, ok := s.AsUnion().(opencode.IntegrationAttemptStatusFailed); !ok {
			t.Errorf("AsUnion: expected IntegrationAttemptStatusFailed, got %T", s.AsUnion())
		}
	})

	t.Run("pending", func(t *testing.T) {
		t.Parallel()
		jsonStr := `{"status":"pending","time":{"created":1700000000.0,"expires":1700003600.0}}`
		var s opencode.IntegrationAttemptStatus
		if err := json.Unmarshal([]byte(jsonStr), &s); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if s.Status != opencode.IntegrationAttemptStatusTypePending {
			t.Errorf("Status: got %q, want %q", s.Status, opencode.IntegrationAttemptStatusTypePending)
		}
		if s.Message != "" {
			t.Errorf("Message: got %q, want empty", s.Message)
		}
		if s.Time.Created != 1700000000.0 {
			t.Errorf("Time.Created: got %v, want 1700000000.0", s.Time.Created)
		}
		if _, ok := s.AsUnion().(opencode.IntegrationAttemptStatusPending); !ok {
			t.Errorf("AsUnion: expected IntegrationAttemptStatusPending, got %T", s.AsUnion())
		}
	})
}

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
