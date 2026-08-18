// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestExperimentalWorkspaceList(t *testing.T) {
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
	_, err := client.Experimental.WorkspaceList(context.TODO(), opencode.ExperimentalWorkspaceListParams{
		Directory: opencode.F("directory"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentalWorkspaceCreate(t *testing.T) {
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
	_, err := client.Experimental.Workspace.New(context.TODO(), opencode.ExperimentalWorkspaceNewParams{
		Body: opencode.F(opencode.ExperimentalWorkspaceCreateInput{
			Type:   opencode.F("type"),
			Branch: opencode.F("main"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentalWorkspaceRemove(t *testing.T) {
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
	_, err := client.Experimental.WorkspaceRemove(context.TODO(), "workspaceID", opencode.ExperimentalWorkspaceRemoveParams{
		Directory: opencode.F("directory"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestExperimentalWorkspaceNewFlatMethod verifies that the WorkspaceNew flat method exists
// on ExperimentalService and has the correct signature (compile-time check).
func TestExperimentalWorkspaceNewFlatMethod(t *testing.T) {
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
	_, err := client.Experimental.WorkspaceNew(context.TODO(), opencode.ExperimentalWorkspaceNewParams{
		Body: opencode.F(opencode.ExperimentalWorkspaceCreateInput{
			Type: opencode.F("type"),
		}),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestExperimentalWorkspaceAdapterListFlatMethod verifies the WorkspaceAdapterList alias exists.
func TestExperimentalWorkspaceAdapterListFlatMethod(t *testing.T) {
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
	_, err := client.Experimental.WorkspaceAdapterList(context.TODO(), opencode.ExperimentalAdapterListParams{
		Directory: opencode.F("directory"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestExperimentalConsoleSwitchOrgParamsAlias verifies the type alias is in place.
func TestExperimentalConsoleSwitchOrgParamsAlias(t *testing.T) {
	// Compile-time type compatibility check: ConsoleSwitchOrgParams must be an alias for
	// ExperimentalConsoleSwitchOrgParams.
	var _ opencode.ExperimentalConsoleSwitchOrgParams = opencode.ConsoleSwitchOrgParams{}
	var _ opencode.ExperimentalConsoleSwitchOrgInput = opencode.ConsoleSwitchOrgInput{}
}

// TestExperimentalWarpParamsNullIDSerialization verifies that ExperimentalWarpParams with
// ID set to Null[string]() serializes the "id" field as JSON null.
// OpenAPI spec: /experimental/workspace/warp "id" is anyOf[string(^wrk), null] and required,
// allowing explicit null to express "detach" semantics.
// param.Field[string] with Null[string]() correctly serializes to "id": null via
// the Stainless encoder (field.go Null bool handling in encoder.go:newFieldTypeEncoder).
func TestExperimentalWarpParamsNullIDSerialization(t *testing.T) {
	t.Parallel()
	params := opencode.ExperimentalWarpParams{
		ID:        opencode.Null[string](),
		SessionID: opencode.F("ses_test123"),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	// "id" must be present and null (not absent)
	idVal, present := got["id"]
	if !present {
		t.Errorf("id field: expected present with null value, but field was absent; body=%s", string(data))
	}
	if idVal != nil {
		t.Errorf("id field: expected null, got %v (%T)", idVal, idVal)
	}

	// sessionID must be present
	if got["sessionID"] != "ses_test123" {
		t.Errorf("sessionID: got %v, want %q", got["sessionID"], "ses_test123")
	}
}

// TestExperimentalWarpParamsStringIDSerialization verifies that a normal string ID
// is serialized correctly.
func TestExperimentalWarpParamsStringIDSerialization(t *testing.T) {
	t.Parallel()
	params := opencode.ExperimentalWarpParams{
		ID:        opencode.F("wrk_abc123"),
		SessionID: opencode.F("ses_xyz456"),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	if got["id"] != "wrk_abc123" {
		t.Errorf("id: got %v, want %q", got["id"], "wrk_abc123")
	}
	if got["sessionID"] != "ses_xyz456" {
		t.Errorf("sessionID: got %v, want %q", got["sessionID"], "ses_xyz456")
	}
}

// TestWorkspaceTimeUsedDeserialization verifies that Workspace.TimeUsed correctly
// deserializes all four valid forms: float64 number, "NaN", "Infinity", "-Infinity".
// OpenAPI schema: anyOf[number, "NaN", "Infinity", "-Infinity"].
func TestWorkspaceTimeUsedDeserialization(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name     string
		json     string
		wantType string
		wantVal  any
	}{
		{
			name:     "number",
			json:     `{"id":"wrk_1","type":"local","name":"ws","projectID":"proj","timeUsed":12345.6}`,
			wantType: "float64",
			wantVal:  float64(12345.6),
		},
		{
			name:     "NaN",
			json:     `{"id":"wrk_1","type":"local","name":"ws","projectID":"proj","timeUsed":"NaN"}`,
			wantType: "string",
			wantVal:  "NaN",
		},
		{
			name:     "Infinity",
			json:     `{"id":"wrk_1","type":"local","name":"ws","projectID":"proj","timeUsed":"Infinity"}`,
			wantType: "string",
			wantVal:  "Infinity",
		},
		{
			name:     "NegInfinity",
			json:     `{"id":"wrk_1","type":"local","name":"ws","projectID":"proj","timeUsed":"-Infinity"}`,
			wantType: "string",
			wantVal:  "-Infinity",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var ws opencode.Workspace
			if err := json.Unmarshal([]byte(tc.json), &ws); err != nil {
				t.Fatalf("Unmarshal: %v", err)
			}
			if ws.TimeUsed == nil {
				t.Fatal("TimeUsed: expected non-nil, got nil")
			}
			switch v := ws.TimeUsed.(type) {
			case float64:
				if tc.wantType != "float64" {
					t.Errorf("TimeUsed type: got float64, want %s", tc.wantType)
				}
				if v != tc.wantVal.(float64) {
					t.Errorf("TimeUsed value: got %v, want %v", v, tc.wantVal)
				}
			case string:
				if tc.wantType != "string" {
					t.Errorf("TimeUsed type: got string, want %s", tc.wantType)
				}
				if v != tc.wantVal.(string) {
					t.Errorf("TimeUsed value: got %q, want %q", v, tc.wantVal)
				}
			default:
				t.Errorf("TimeUsed: unexpected type %T, value %v", ws.TimeUsed, ws.TimeUsed)
			}
		})
	}
}

// TestWorkspaceEventConnectionStatusAlias verifies backward-compatible alias.
func TestWorkspaceEventConnectionStatusAlias(t *testing.T) {
	// Type alias check: WorkspaceStatusItem must be identical to WorkspaceEventConnectionStatus.
	var wecs opencode.WorkspaceEventConnectionStatus
	var wsi opencode.WorkspaceStatusItem = wecs
	_ = wsi

	// Enum alias check.
	var s opencode.WorkspaceStatusItemStatus = opencode.WorkspaceStatusItemStatusConnected
	_ = s
	var s2 opencode.WorkspaceEventConnectionStatusStatus = opencode.WorkspaceEventConnectionStatusStatusConnected
	_ = s2
}

// TestExperimentalConsoleSwitchOrgParamsBodyPresence verifies that
// ExperimentalConsoleSwitchOrgParams.Body follows the param.Field[T] Present
// contract: unset -> MarshalJSON returns nil (no body sent), matching JS
// SDK(v2) `switchOrg({directory?, workspace?, accountID?, orgID?})` where the
// body-mapped fields are all optional (sdk.gen.ts:770-802). OpenAPI requestBody
// for POST /experimental/console/switch has no `required` key (optional body).
func TestExperimentalConsoleSwitchOrgParamsBodyPresence(t *testing.T) {
	t.Parallel()

	// Unset Body -> no body bytes.
	unset := opencode.ExperimentalConsoleSwitchOrgParams{
		Directory: opencode.F("dir"),
	}
	data, err := unset.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON (unset): %v", err)
	}
	if data != nil {
		t.Errorf("MarshalJSON with unset Body: expected nil, got %q", data)
	}

	// Present Body -> emits accountID/orgID.
	present := opencode.ExperimentalConsoleSwitchOrgParams{
		Body: opencode.F(opencode.ExperimentalConsoleSwitchOrgInput{
			AccountID: opencode.F("acc_1"),
			OrgID:     opencode.F("org_1"),
		}),
	}
	data, err = present.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON (present): %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("re-parse: %v", err)
	}
	if got["accountID"] != "acc_1" {
		t.Errorf("accountID: got %v, want acc_1", got["accountID"])
	}
	if got["orgID"] != "org_1" {
		t.Errorf("orgID: got %v, want org_1", got["orgID"])
	}
}

// TestExperimentalWorkspaceNewParamsBodyPresence verifies that
// ExperimentalWorkspaceNewParams.Body follows the param.Field[T] Present
// contract: unset -> MarshalJSON returns nil (no body sent), matching JS
// SDK(v2) `workspace.create({directory?, workspace?, id?, type?, branch?, extra?})`
// where all body-mapped fields are optional (sdk.gen.ts:1046-1071). OpenAPI
// requestBody for POST /experimental/workspace has no `required` key.
func TestExperimentalWorkspaceNewParamsBodyPresence(t *testing.T) {
	t.Parallel()

	unset := opencode.ExperimentalWorkspaceNewParams{
		Directory: opencode.F("dir"),
	}
	data, err := unset.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON (unset): %v", err)
	}
	if data != nil {
		t.Errorf("MarshalJSON with unset Body: expected nil, got %q", data)
	}

	present := opencode.ExperimentalWorkspaceNewParams{
		Body: opencode.F(opencode.ExperimentalWorkspaceCreateInput{
			Type:   opencode.F("local"),
			Branch: opencode.F("main"),
		}),
	}
	data, err = present.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON (present): %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("re-parse: %v", err)
	}
	if got["type"] != "local" {
		t.Errorf("type: got %v, want local", got["type"])
	}
	if got["branch"] != "main" {
		t.Errorf("branch: got %v, want main", got["branch"])
	}
}

// TestWorkspaceStatusResponseTypeUnmarshal verifies that the Workspace.Status
// (and flat ExperimentalService.WorkspaceStatus) canonical response element
// type WorkspaceEventConnectionStatus deserializes correctly, and that the
// deprecated WorkspaceStatusItem alias remains usable (compile-time + runtime
// equivalence). OpenAPI: `experimental.workspace.status` 200 =
// array[WorkspaceEventConnectionStatus].
func TestWorkspaceStatusResponseTypeUnmarshal(t *testing.T) {
	t.Parallel()

	var items []opencode.WorkspaceEventConnectionStatus
	raw := `[{"workspaceID":"w1","status":"connected"}]`
	if err := json.Unmarshal([]byte(raw), &items); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0].WorkspaceID != "w1" {
		t.Errorf("WorkspaceID: got %q, want %q", items[0].WorkspaceID, "w1")
	}
	if items[0].Status != opencode.WorkspaceEventConnectionStatusStatusConnected {
		t.Errorf("Status: got %q, want %q", items[0].Status, opencode.WorkspaceEventConnectionStatusStatusConnected)
	}

	// The deprecated alias must still deserialize identically (compile-time
	// assertion that WorkspaceStatusItem == WorkspaceEventConnectionStatus).
	var aliasItems []opencode.WorkspaceStatusItem
	if err := json.Unmarshal([]byte(raw), &aliasItems); err != nil {
		t.Fatalf("Unmarshal (alias): %v", err)
	}
	if aliasItems[0].WorkspaceID != "w1" || aliasItems[0].Status != opencode.WorkspaceStatusItemStatusConnected {
		t.Errorf("alias unmarshal mismatch: %+v", aliasItems[0])
	}
}
