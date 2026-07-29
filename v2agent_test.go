// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	opencode "github.com/sst/opencode-sdk-go"
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

// TestV2AgentListWithOptionalParams tests the v2 agent list endpoint.
// Aligned with OpenAPI operationId "v2.agent.list", GET /api/agent.
// Query parameters: location (optional, nested object with directory and workspace).
func TestV2AgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.V2Agent.List(context.TODO(), opencode.V2AgentListParams{
		Location: opencode.F(opencode.V2LocationParam{
			Directory: opencode.F("directory"),
			Workspace: opencode.F("workspace"),
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

// TestV2AgentListParamsURLQuery verifies the query serialization for V2AgentListParams.
// Aligned with OpenAPI GET /api/agent location[directory] and location[workspace] params.
func TestV2AgentListParamsURLQuery(t *testing.T) {
	t.Run("with location directory and workspace", func(t *testing.T) {
		p := opencode.V2AgentListParams{
			Location: opencode.F(opencode.V2LocationParam{
				Directory: opencode.F("mydir"),
				Workspace: opencode.F("mywsp"),
			}),
		}
		got := p.URLQuery().Encode()
		// Expects nested bracket format: location[directory]=...&location[workspace]=...
		if got == "" {
			t.Error("expected non-empty query string")
		}
	})

	t.Run("empty params produces no query", func(t *testing.T) {
		p := opencode.V2AgentListParams{}
		got := p.URLQuery().Encode()
		if got != "" {
			t.Errorf("expected empty query, got %q", got)
		}
	})
}

// TestV2AgentListResponseUnmarshal verifies deserialization of V2AgentListResponse.
// Aligned with OpenAPI schema for GET /api/agent response.
// Required fields: location (LocationInfo), data ([]V2AgentInfo).
func TestV2AgentListResponseUnmarshal(t *testing.T) {
	t.Run("full response with agents", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/home/user/project",
				"workspaceID": "ws_123",
				"project": {"id": "proj_1", "directory": "/home/user/project"}
			},
			"data": [
				{
					"id": "agent_build",
					"model": {"id": "claude-3", "providerID": "anthropic"},
					"request": {"model": "claude-3", "temperature": 0, "headers": {}},
					"system": "You are a build agent.",
					"description": "Handles build tasks",
					"mode": "primary",
					"hidden": false,
					"color": "primary",
					"steps": 10,
					"permissions": []
				}
			]
		}`
		var resp opencode.V2AgentListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if resp.Location.Directory != "/home/user/project" {
			t.Errorf("Location.Directory = %q, want /home/user/project", resp.Location.Directory)
		}
		if resp.Location.WorkspaceID != "ws_123" {
			t.Errorf("Location.WorkspaceID = %q, want ws_123", resp.Location.WorkspaceID)
		}
		if resp.Location.Project.ID != "proj_1" {
			t.Errorf("Location.Project.ID = %q, want proj_1", resp.Location.Project.ID)
		}
		if len(resp.Data) != 1 {
			t.Fatalf("Data len = %d, want 1", len(resp.Data))
		}
		agent := resp.Data[0]
		if agent.ID != "agent_build" {
			t.Errorf("agent.ID = %q, want agent_build", agent.ID)
		}
		if agent.Mode != opencode.V2AgentInfoModePrimary {
			t.Errorf("agent.Mode = %q, want primary", agent.Mode)
		}
		if agent.Hidden {
			t.Error("agent.Hidden should be false")
		}
		if agent.Steps != 10 {
			t.Errorf("agent.Steps = %d, want 10", agent.Steps)
		}
		if resp.JSON.RawJSON() == "" {
			t.Error("RawJSON should be preserved")
		}
	})

	t.Run("empty data array", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/",
				"project": {"id": "p", "directory": "/"}
			},
			"data": []
		}`
		var resp opencode.V2AgentListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if len(resp.Data) != 0 {
			t.Errorf("expected 0 agents, got %d", len(resp.Data))
		}
	})

	t.Run("unknown fields tolerated via ExtraFields", func(t *testing.T) {
		raw := `{
			"location": {
				"directory": "/",
				"project": {"id": "p", "directory": "/"}
			},
			"data": [],
			"future_field": "value"
		}`
		var resp opencode.V2AgentListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal with unknown field: %v", err)
		}
	})
}

// TestV2AgentInfoColorUnion verifies that V2AgentInfo.Color union handles both
// preset names and hex codes, with AsColorUnion()/AsColor() for typed access.
func TestV2AgentInfoColorUnion(t *testing.T) {
	t.Run("preset color name primary", func(t *testing.T) {
		raw := `{
			"id": "a1",
			"model": {"id": "m", "providerID": "p"},
			"request": {"model": "m", "temperature": 0, "headers": {}},
			"mode": "primary",
			"hidden": false,
			"color": "primary",
			"permissions": []
		}`
		var agent opencode.V2AgentInfo
		if err := json.Unmarshal([]byte(raw), &agent); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		color := agent.AsColor()
		if color != opencode.AgentColorPrimary {
			t.Errorf("color = %q, want primary", color)
		}
		if !color.IsKnown() {
			t.Error("primary should be a known color")
		}
	})

	t.Run("hex color code", func(t *testing.T) {
		raw := `{
			"id": "a2",
			"model": {"id": "m", "providerID": "p"},
			"request": {"model": "m", "temperature": 0, "headers": {}},
			"mode": "subagent",
			"hidden": true,
			"color": "#FF5733",
			"permissions": []
		}`
		var agent opencode.V2AgentInfo
		if err := json.Unmarshal([]byte(raw), &agent); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		u := agent.AsColorUnion()
		if u == nil {
			t.Fatal("AsColorUnion() should not be nil for hex color")
		}
		color, ok := u.(opencode.AgentColor)
		if !ok {
			t.Fatalf("expected AgentColor, got %T", u)
		}
		if string(color) != "#FF5733" {
			t.Errorf("color = %q, want #FF5733", color)
		}
		// Hex colors are not in the known presets list
		if color.IsKnown() {
			t.Error("hex color #FF5733 should not be a known preset")
		}
	})

	t.Run("missing color field — no error and union is nil", func(t *testing.T) {
		raw := `{
			"id": "a3",
			"model": {"id": "m", "providerID": "p"},
			"request": {"model": "m", "temperature": 0, "headers": {}},
			"mode": "all",
			"hidden": false,
			"permissions": []
		}`
		var agent opencode.V2AgentInfo
		if err := json.Unmarshal([]byte(raw), &agent); err != nil {
			t.Fatalf("unmarshal without color: %v", err)
		}
		if agent.AsColorUnion() != nil {
			t.Errorf("AsColorUnion() = %v, want nil when color absent", agent.AsColorUnion())
		}
	})
}

// TestV2AgentInfoColorNullDoesNotAbortDecode covers the boundary shapes of the
// union-carrying `color` field: an explicit null, an absent field, and a wrong-typed
// object.
//
// OpenAPI AgentV2Info does NOT list `color` in required (required = id, request,
// mode, hidden, permissions), and AgentColor's anyOf holds only two string members
// (hex pattern | preset enum) with no `{"type":"null"}`. An optional non-nullable
// property is exactly the shape a server is most likely to emit as an explicit null,
// yet the single registered variant filters on gjson.String, so routing a null fails
// with "was not able to coerce type as union" and the error propagating out of
// UnmarshalJSON kills the surrounding document — inside V2AgentListResponse it
// silently empties the whole `data` array.
func TestV2AgentInfoColorNullDoesNotAbortDecode(t *testing.T) {
	const scaffold = `"id":"a1","request":{"headers":{},"body":{}},"mode":"primary","hidden":false,"permissions":[]`

	t.Run("explicit null is treated as absent", func(t *testing.T) {
		var agent opencode.V2AgentInfo
		if err := json.Unmarshal([]byte(`{`+scaffold+`,"color":null}`), &agent); err != nil {
			t.Fatalf("color:null must not fail the decode, got: %v", err)
		}
		if agent.Color != nil {
			t.Errorf("Color = %#v, want nil for a null color", agent.Color)
		}
		if agent.AsColorUnion() != nil {
			t.Errorf("AsColorUnion() = %v, want nil for a null color", agent.AsColorUnion())
		}
		if agent.AsColor() != "" {
			t.Errorf("AsColor() = %q, want empty for a null color", agent.AsColor())
		}
		if agent.ID != "a1" || agent.Mode != opencode.V2AgentInfoModePrimary {
			t.Errorf("sibling fields lost: id=%q mode=%q", agent.ID, agent.Mode)
		}
	})

	t.Run("null color inside the list envelope keeps every element", func(t *testing.T) {
		raw := `{"location":{"directory":"/d","project":{"id":"pr","directory":"/d"}},"data":[` +
			`{"id":"a1","request":{"headers":{},"body":{}},"mode":"primary","hidden":false,"permissions":[],"color":null},` +
			`{"id":"a2","request":{"headers":{},"body":{}},"mode":"all","hidden":true,"permissions":[],"color":"accent"}]}`
		var resp opencode.V2AgentListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal envelope: %v", err)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("len(Data) = %d, want 2 — a null `color` must not drop list elements", len(resp.Data))
		}
		if resp.Data[0].ID != "a1" || resp.Data[0].Color != nil {
			t.Errorf("Data[0] = {id:%q color:%#v}, want {a1 <nil>}", resp.Data[0].ID, resp.Data[0].Color)
		}
		if resp.Data[1].AsColor() != opencode.AgentColorAccent {
			t.Errorf("Data[1].AsColor() = %q, want accent", resp.Data[1].AsColor())
		}
	})

	t.Run("wrong-typed color is reported through field metadata, not a fatal error", func(t *testing.T) {
		// AgentColor admits only strings, so an object is an illegal server
		// response. Because `color` is declared as the registered [AgentColorUnion],
		// the framework's struct decoder records the coercion failure on the field's
		// metadata and keeps decoding the rest of the document. The "null == absent"
		// vs "wrong type" distinction survives on JSON.Color rather than as a
		// document-killing error.
		var agent opencode.V2AgentInfo
		if err := json.Unmarshal([]byte(`{`+scaffold+`,"color":{"x":1}}`), &agent); err != nil {
			t.Fatalf("a wrong-typed color must not fail the whole decode, got: %v", err)
		}
		if agent.Color != nil {
			t.Errorf("Color = %#v, want nil for a non-string color", agent.Color)
		}
		if agent.AsColorUnion() != nil || agent.AsColor() != "" {
			t.Errorf("AsColorUnion() = %v / AsColor() = %q, want nil and empty", agent.AsColorUnion(), agent.AsColor())
		}
		if !agent.JSON.Color.IsInvalid() {
			t.Error("JSON.Color.IsInvalid() = false, want true for a non-string color")
		}
		if agent.JSON.Color.IsNull() || agent.JSON.Color.IsMissing() {
			t.Error("a wrong-typed color must stay distinguishable from null/absent")
		}
		if got := agent.JSON.Color.Raw(); got != `{"x":1}` {
			t.Errorf("JSON.Color.Raw() = %s, want %s — the raw value must stay available", got, `{"x":1}`)
		}
		if agent.ID != "a1" || agent.Mode != opencode.V2AgentInfoModePrimary {
			t.Errorf("sibling fields lost: id=%q mode=%q", agent.ID, agent.Mode)
		}
	})

	t.Run("wrong-typed color inside the list envelope keeps every element", func(t *testing.T) {
		// The regression this locks down: routing the sub-document by hand inside
		// UnmarshalJSON returned the coercion error out of the element decoder, the
		// slice decoder propagated it, and the outer struct decoder swallowed it —
		// so `data` came back empty with a nil error, losing the valid element too.
		raw := `{"location":{"directory":"/d","project":{"id":"pr","directory":"/d"}},"data":[` +
			`{"id":"a1","request":{"headers":{},"body":{}},"mode":"primary","hidden":false,"permissions":[],"color":{"x":1}},` +
			`{"id":"a2","request":{"headers":{},"body":{}},"mode":"all","hidden":true,"permissions":[],"color":"accent"}]}`
		var resp opencode.V2AgentListResponse
		if err := json.Unmarshal([]byte(raw), &resp); err != nil {
			t.Fatalf("unmarshal envelope: %v", err)
		}
		if len(resp.Data) != 2 {
			t.Fatalf("len(Data) = %d, want 2 — a wrong-typed `color` must not drop list elements", len(resp.Data))
		}
		if resp.Data[0].ID != "a1" || resp.Data[0].Color != nil {
			t.Errorf("Data[0] = {id:%q color:%#v}, want {a1 <nil>}", resp.Data[0].ID, resp.Data[0].Color)
		}
		if resp.Data[1].AsColor() != opencode.AgentColorAccent {
			t.Errorf("Data[1].AsColor() = %q, want accent", resp.Data[1].AsColor())
		}
	})

	t.Run("unrecognised colour string still lands on AgentColor", func(t *testing.T) {
		// Forward compatibility: AgentColor's anyOf is `string | preset enum`, so any
		// JSON string is legal. A value outside the preset set must decode, with
		// IsKnown() reporting false as the way to detect it.
		var agent opencode.V2AgentInfo
		if err := json.Unmarshal([]byte(`{`+scaffold+`,"color":"chartreuse"}`), &agent); err != nil {
			t.Fatalf("an unrecognised colour must not fail the decode, got: %v", err)
		}
		if got := agent.AsColor(); got != "chartreuse" {
			t.Errorf("AsColor() = %q, want chartreuse", got)
		}
		if agent.AsColor().IsKnown() {
			t.Error("IsKnown() = true for chartreuse, want false so callers can detect it")
		}
		if _, ok := agent.Color.(opencode.AgentColor); !ok {
			t.Errorf("Color = %T, want opencode.AgentColor", agent.Color)
		}
	})
}

// TestV2AgentInfoModeIsKnown verifies all known V2AgentInfoMode values.
func TestV2AgentInfoModeIsKnown(t *testing.T) {
	known := []opencode.V2AgentInfoMode{
		opencode.V2AgentInfoModeSubagent,
		opencode.V2AgentInfoModePrimary,
		opencode.V2AgentInfoModeAll,
	}
	for _, m := range known {
		if !m.IsKnown() {
			t.Errorf("%q should be known", m)
		}
	}
	if opencode.V2AgentInfoMode("unknown").IsKnown() {
		t.Error("unknown mode should not be known")
	}
}

// TestAgentColorIsKnown verifies all known AgentColor preset values.
func TestAgentColorIsKnown(t *testing.T) {
	known := []opencode.AgentColor{
		opencode.AgentColorPrimary,
		opencode.AgentColorSecondary,
		opencode.AgentColorAccent,
		opencode.AgentColorSuccess,
		opencode.AgentColorWarning,
		opencode.AgentColorError,
		opencode.AgentColorInfo,
	}
	for _, c := range known {
		if !c.IsKnown() {
			t.Errorf("%q should be known", c)
		}
	}
	if opencode.AgentColor("#FF0000").IsKnown() {
		t.Error("hex color should not be a known preset")
	}
}
