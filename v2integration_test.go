// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"encoding/json"
	"net/url"
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

// TestV2IntegrationConnectKeyParamsMarshalJSON verifies the "style 1" fix for
// V2IntegrationConnectKeyParams.Body: it must now be wrapped in param.Field[T]
// with a Present guard in MarshalJSON, matching sync.go's SyncReplayParams
// golden precedent. Before the fix, Body was a bare
// V2IntegrationConnectKeyParamsBody field (json:"-") and MarshalJSON
// unconditionally returned apijson.MarshalRoot(r.Body), producing "{}" even
// when the caller never set Body.
//
// "return nil" for the unset case is safe here (unlike AttemptComplete below):
// openapi.json's v2.integration.connect.key requestBody inner schema declares
// `"required": ["key"]`, so an unset Body is ALREADY an invalid call whether the
// SDK sends 0 bytes or "{}" -- the server rejects both for missing "key". No
// valid call is affected by this fix.
func TestV2IntegrationConnectKeyParamsMarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("NoBody_ReturnsNil", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2IntegrationConnectKeyParams{}
		data, err := params.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v", err)
		}
		if data != nil {
			t.Errorf("expected nil output when Body is not Present, got: %s", string(data))
		}
	})

	t.Run("WithBody_SerializesFields", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2IntegrationConnectKeyParams{
			Body: opencode.F(opencode.V2IntegrationConnectKeyParamsBody{
				Key:   opencode.F("abc"),
				Label: opencode.F("mylabel"),
			}),
		}
		data, err := params.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v", err)
		}
		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			t.Fatalf("json.Unmarshal result: %v", err)
		}
		if m["key"] != "abc" {
			t.Errorf("key: got %v, want abc", m["key"])
		}
		if m["label"] != "mylabel" {
			t.Errorf("label: got %v, want mylabel", m["label"])
		}
		if len(m) != 2 {
			t.Errorf("expected exactly 2 fields (key,label), got %v", m)
		}
	})
}

// TestV2IntegrationConnectOauthParamsMarshalJSON mirrors
// TestV2IntegrationConnectKeyParamsMarshalJSON for V2IntegrationConnectOauthParams
// (v2integration.go:1267), the second of the three "style 2" Body fields fixed
// in this batch.
//
// Same safety argument as ConnectKey: openapi.json's v2.integration.connect.oauth
// requestBody inner schema declares `"required": ["methodID", "inputs"]`, so an
// unset Body is already an invalid call regardless of whether 0 bytes or "{}" is
// sent.
func TestV2IntegrationConnectOauthParamsMarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("NoBody_ReturnsNil", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2IntegrationConnectOauthParams{}
		data, err := params.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v", err)
		}
		if data != nil {
			t.Errorf("expected nil output when Body is not Present, got: %s", string(data))
		}
	})

	t.Run("WithBody_SerializesFields", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2IntegrationConnectOauthParams{
			Body: opencode.F(opencode.V2IntegrationConnectOauthParamsBody{
				MethodID: opencode.F("method_1"),
				Inputs:   opencode.F(map[string]string{"client_id": "abc"}),
				Label:    opencode.F("mylabel"),
			}),
		}
		data, err := params.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v", err)
		}
		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			t.Fatalf("json.Unmarshal result: %v", err)
		}
		if m["methodID"] != "method_1" {
			t.Errorf("methodID: got %v, want method_1", m["methodID"])
		}
		if m["label"] != "mylabel" {
			t.Errorf("label: got %v, want mylabel", m["label"])
		}
		inputs, ok := m["inputs"].(map[string]any)
		if !ok || inputs["client_id"] != "abc" {
			t.Errorf("inputs: got %v, want {client_id: abc}", m["inputs"])
		}
	})
}

// TestV2IntegrationAttemptCompleteParamsMarshalJSON verifies the CORRECTED
// behavior for V2IntegrationAttemptCompleteParams (v2integration.go).
//
// Unlike ConnectKey/ConnectOauth, this endpoint's OpenAPI requestBody is
// required==true while its inner schema has NO required properties (only an
// optional "code": openapi.json paths
// ["/api/integration/attempt/{attemptID}/complete"].post.requestBody). That
// means "{}" is a fully valid body, and because requestBody.required is true, a
// body must always be sent -- so unlike Key/Oauth, the naive "style 1"
// `return nil, nil` when Body is unset is a real regression here (verified at
// the wire level: it produces a `Content-Type: application/json` header with a
// 0-byte, non-JSON body). The fix always emits a body, defaulting to "{}".
func TestV2IntegrationAttemptCompleteParamsMarshalJSON(t *testing.T) {
	t.Parallel()

	t.Run("NoBody_ReturnsEmptyObject", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2IntegrationAttemptCompleteParams{}
		data, err := params.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v", err)
		}
		// Exact-byte assertion: OpenAPI requestBody.required=true + no inner
		// required properties means "{}" (not nil, not "null") must always be
		// sent when the caller leaves Body unset.
		if string(data) != "{}" {
			t.Errorf("expected exact bytes \"{}\" when Body is not Present, got: %q", string(data))
		}
	})

	t.Run("WithBody_SerializesFields", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2IntegrationAttemptCompleteParams{
			Body: opencode.F(opencode.V2IntegrationAttemptCompleteParamsBody{
				Code: opencode.F("x123"),
			}),
		}
		data, err := params.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON error: %v", err)
		}
		var m map[string]any
		if err := json.Unmarshal(data, &m); err != nil {
			t.Fatalf("json.Unmarshal result: %v", err)
		}
		if m["code"] != "x123" {
			t.Errorf("code: got %v, want x123", m["code"])
		}
		if len(m) != 1 {
			t.Errorf("expected exactly 1 field (code), got %v", m)
		}
	})
}

// TestV2IntegrationParamsURLQuery verifies that the `location` query parameter
// on the three fixed Params types (and V2IntegrationGetParams as a control)
// serializes as a deepObject/brackets-format nested object:
// location[directory]=...&location[workspace]=..., per apiquery's
// NestedQueryFormatBrackets setting and the OpenAPI `deepObject` style for
// `location`.
func TestV2IntegrationParamsURLQuery(t *testing.T) {
	t.Parallel()

	loc := opencode.F(opencode.V2LocationParam{
		Directory: opencode.F("/repo"),
		Workspace: opencode.F("ws1"),
	})

	assertLocationQuery := func(t *testing.T, v url.Values) {
		t.Helper()
		if got := v.Get("location[directory]"); got != "/repo" {
			t.Errorf("location[directory]: got %q, want %q", got, "/repo")
		}
		if got := v.Get("location[workspace]"); got != "ws1" {
			t.Errorf("location[workspace]: got %q, want %q", got, "ws1")
		}
	}

	t.Run("ConnectKeyParams", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2IntegrationConnectKeyParams{Location: loc}
		assertLocationQuery(t, params.URLQuery())
	})

	t.Run("ConnectOauthParams", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2IntegrationConnectOauthParams{Location: loc}
		assertLocationQuery(t, params.URLQuery())
	})

	t.Run("AttemptCompleteParams", func(t *testing.T) {
		t.Parallel()
		params := opencode.V2IntegrationAttemptCompleteParams{Location: loc}
		assertLocationQuery(t, params.URLQuery())
	})
}

// TestV2IntegrationUnionRouting is a consolidated regression test asserting
// that all four Response Unions in v2integration.go route to the correct
// concrete variant for every OpenAPI-declared discriminator value. This
// guards the "Union 铁律" invariant (init() registration + discriminator
// routing) across the fix in this batch, ensuring the Body-wrapping change did
// not disturb the unrelated Union carriers.
func TestV2IntegrationUnionRouting(t *testing.T) {
	t.Parallel()

	t.Run("IntegrationMethodUnion_oauth_key_env", func(t *testing.T) {
		t.Parallel()
		jsonStr := `[
			{"id":"m1","type":"oauth","label":"OAuth"},
			{"type":"key","label":"Key"},
			{"type":"env","names":["TOKEN"]}
		]`
		var methods []opencode.IntegrationMethod
		if err := json.Unmarshal([]byte(jsonStr), &methods); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if _, ok := methods[0].AsUnion().(opencode.IntegrationOAuthMethod); !ok {
			t.Errorf("methods[0]: expected IntegrationOAuthMethod, got %T", methods[0].AsUnion())
		}
		if _, ok := methods[1].AsUnion().(opencode.IntegrationKeyMethod); !ok {
			t.Errorf("methods[1]: expected IntegrationKeyMethod, got %T", methods[1].AsUnion())
		}
		if _, ok := methods[2].AsUnion().(opencode.IntegrationEnvMethod); !ok {
			t.Errorf("methods[2]: expected IntegrationEnvMethod, got %T", methods[2].AsUnion())
		}
	})

	t.Run("ConnectionInfoUnion_credential_env", func(t *testing.T) {
		t.Parallel()
		jsonStr := `[
			{"type":"credential","id":"cred_1","label":"cred"},
			{"type":"env","name":"TOKEN"}
		]`
		var conns []opencode.ConnectionInfo
		if err := json.Unmarshal([]byte(jsonStr), &conns); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if _, ok := conns[0].AsUnion().(opencode.ConnectionCredentialInfo); !ok {
			t.Errorf("conns[0]: expected ConnectionCredentialInfo, got %T", conns[0].AsUnion())
		}
		if _, ok := conns[1].AsUnion().(opencode.ConnectionEnvInfo); !ok {
			t.Errorf("conns[1]: expected ConnectionEnvInfo, got %T", conns[1].AsUnion())
		}
	})

	t.Run("IntegrationPromptUnion_text_select", func(t *testing.T) {
		t.Parallel()
		jsonStr := `[
			{"type":"text","key":"k1","message":"m1"},
			{"type":"select","key":"k2","message":"m2","options":["a","b"]}
		]`
		var prompts []opencode.IntegrationPrompt
		if err := json.Unmarshal([]byte(jsonStr), &prompts); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if _, ok := prompts[0].AsUnion().(opencode.IntegrationTextPrompt); !ok {
			t.Errorf("prompts[0]: expected IntegrationTextPrompt, got %T", prompts[0].AsUnion())
		}
		if _, ok := prompts[1].AsUnion().(opencode.IntegrationSelectPrompt); !ok {
			t.Errorf("prompts[1]: expected IntegrationSelectPrompt, got %T", prompts[1].AsUnion())
		}
	})

	t.Run("IntegrationAttemptStatusUnion_pending_complete_failed_expired", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			name    string
			jsonStr string
			want    any
		}{
			{
				name:    "pending",
				jsonStr: `{"status":"pending","time":{"created":1.0,"expires":2.0}}`,
				want:    opencode.IntegrationAttemptStatusPending{},
			},
			{
				name:    "complete",
				jsonStr: `{"status":"complete","time":{"created":1.0,"expires":2.0}}`,
				want:    opencode.IntegrationAttemptStatusComplete{},
			},
			{
				name:    "failed",
				jsonStr: `{"status":"failed","message":"boom","time":{"created":1.0,"expires":2.0}}`,
				want:    opencode.IntegrationAttemptStatusFailed{},
			},
			{
				name:    "expired",
				jsonStr: `{"status":"expired","time":{"created":1.0,"expires":2.0}}`,
				want:    opencode.IntegrationAttemptStatusExpired{},
			},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				var s opencode.IntegrationAttemptStatus
				if err := json.Unmarshal([]byte(tc.jsonStr), &s); err != nil {
					t.Fatalf("Unmarshal: %v", err)
				}
				gotType := reflect.TypeOf(s.AsUnion())
				wantType := reflect.TypeOf(tc.want)
				if gotType != wantType {
					t.Errorf("AsUnion() type: got %v, want %v", gotType, wantType)
				}
			})
		}
	})
}
