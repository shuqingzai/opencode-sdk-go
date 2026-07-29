package opencode

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/sst/opencode-sdk-go/internal/apijson"
)

// TestV2EventNestedObjectFieldsUseConcreteTypes pins nested object fields of
// V2Event data structs to the concrete Go type of the OpenAPI schema they
// reference. None of these OpenAPI schemas is an `anyOf`/`oneOf` union nor a
// free-form object, so `any` + a runtime-type comment would be wrong twice over:
// it discards type safety and the comment would be factually false (apijson's
// interface decoder yields `map[string]any` for a bare `any` field, never the
// documented struct).
//
// OpenAPI evidence (components/schemas):
//   - QuestionAsked.data.tool                -> $ref QuestionTool          (plain object)
//   - QuestionV2Asked.data.tool              -> $ref QuestionV2Tool        (plain object)
//   - PermissionV2Asked.data.source          -> $ref PermissionV2Source    (plain object)
//   - SessionNextPrompted.data.prompt        -> $ref Prompt                (plain object, required)
//   - SessionNextPromptAdmitted.data.prompt  -> $ref Prompt                (plain object, required)
//   - SessionNextRetried.data.error          -> $ref SessionNextRetry_error (plain object, required)
//
// This matches how the very same schemas are typed everywhere else in the SDK:
// [QuestionRequest.Tool], [QuestionV2Request.Tool], [PermissionV2Request.Source],
// [V2EventSessionNextRevertStagedData.Revert] and
// [EventListResponseEventSessionNextRetriedProperties.Error].
func TestV2EventNestedObjectFieldsUseConcreteTypes(t *testing.T) {
	tests := []struct {
		name  string
		typ   reflect.Type
		field string
		want  reflect.Type
	}{
		{
			name:  "question asked data tool",
			typ:   reflect.TypeFor[V2EventQuestionAskedData](),
			field: "Tool",
			want:  reflect.TypeFor[QuestionTool](),
		},
		{
			name:  "question v2 asked data tool",
			typ:   reflect.TypeFor[V2EventQuestionV2AskedData](),
			field: "Tool",
			want:  reflect.TypeFor[QuestionV2Tool](),
		},
		{
			name:  "permission v2 asked data source",
			typ:   reflect.TypeFor[V2EventPermissionV2AskedData](),
			field: "Source",
			want:  reflect.TypeFor[PermissionV2Source](),
		},
		{
			name:  "session next prompted data prompt",
			typ:   reflect.TypeFor[V2EventSessionNextPromptedData](),
			field: "Prompt",
			want:  reflect.TypeFor[V2SessionInputPrompt](),
		},
		{
			name:  "session next prompt admitted data prompt",
			typ:   reflect.TypeFor[V2EventSessionNextPromptAdmittedData](),
			field: "Prompt",
			want:  reflect.TypeFor[V2SessionInputPrompt](),
		},
		{
			name:  "session next retried data error",
			typ:   reflect.TypeFor[V2EventSessionNextRetriedData](),
			field: "Error",
			want:  reflect.TypeFor[EventListResponseEventSessionNextRetriedError](),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field, ok := tt.typ.FieldByName(tt.field)
			if !ok {
				t.Fatalf("missing %s field", tt.field)
			}
			if field.Type != tt.want {
				t.Fatalf("%s.%s type = %s, want %s", tt.typ.Name(), tt.field, field.Type, tt.want)
			}
			if field.Type.Kind() == reflect.Pointer {
				t.Fatalf("%s.%s must not be a pointer", tt.typ.Name(), tt.field)
			}

			jsonField, ok := tt.typ.FieldByName("JSON")
			if !ok {
				t.Fatal("missing JSON metadata field")
			}
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

// TestV2EventOptionalObjectFieldsRoundtrip verifies that omitting optional
// Tool/Source fields does not cause unmarshal errors.
func TestV2EventOptionalObjectFieldsRoundtrip(t *testing.T) {
	cases := []struct {
		name string
		data string
		into func() any
	}{
		{
			name: "question asked data without tool",
			data: `{"id":"que_1","sessionID":"ses_1","questions":[]}`,
			into: func() any { return new(V2EventQuestionAskedData) },
		},
		{
			name: "question v2 asked data without tool",
			data: `{"id":"que_1","sessionID":"ses_1","questions":[]}`,
			into: func() any { return new(V2EventQuestionV2AskedData) },
		},
		{
			name: "permission v2 asked data without source",
			data: `{"id":"per_1","sessionID":"ses_1","action":"bash","resources":[]}`,
			into: func() any { return new(V2EventPermissionV2AskedData) },
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

// TestV2EventQuestionAskedToolPresent verifies that Tool is fully decoded when
// provided in JSON (OpenAPI QuestionTool: messageID + callID, both required).
func TestV2EventQuestionAskedToolPresent(t *testing.T) {
	raw := `{"id":"que_1","sessionID":"ses_1","questions":[],"tool":{"messageID":"msg_1","callID":"call_1"}}`
	var d V2EventQuestionAskedData
	if err := json.Unmarshal([]byte(raw), &d); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if d.Tool.MessageID != "msg_1" || d.Tool.CallID != "call_1" {
		t.Fatalf("Tool = %+v, want messageID=msg_1 callID=call_1", d.Tool)
	}
}

// TestV2EventQuestionV2AskedToolPresent mirrors the above for
// OpenAPI QuestionV2Tool.
func TestV2EventQuestionV2AskedToolPresent(t *testing.T) {
	raw := `{"id":"que_1","sessionID":"ses_1","questions":[],"tool":{"messageID":"msg_1","callID":"call_1"}}`
	var d V2EventQuestionV2AskedData
	if err := json.Unmarshal([]byte(raw), &d); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if d.Tool.MessageID != "msg_1" || d.Tool.CallID != "call_1" {
		t.Fatalf("Tool = %+v, want messageID=msg_1 callID=call_1", d.Tool)
	}
}

// TestV2EventPermissionV2AskedSourcePresent verifies that Source is fully
// decoded when provided in JSON (OpenAPI PermissionV2Source: type enum "tool",
// messageID, callID, all required).
func TestV2EventPermissionV2AskedSourcePresent(t *testing.T) {
	raw := `{"id":"per_1","sessionID":"ses_1","action":"bash","resources":[],"source":{"type":"tool","messageID":"msg_1","callID":"call_1"}}`
	var d V2EventPermissionV2AskedData
	if err := json.Unmarshal([]byte(raw), &d); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if d.Source.Type != PermissionV2SourceTypeTool {
		t.Errorf("Source.Type = %q, want %q", d.Source.Type, PermissionV2SourceTypeTool)
	}
	if d.Source.MessageID != "msg_1" || d.Source.CallID != "call_1" {
		t.Errorf("Source = %+v", d.Source)
	}
}

// TestV2EventPromptAndRetryErrorDecode exercises the two required nested-object
// fields that previously used `any` (OpenAPI `$ref Prompt` and
// `$ref SessionNextRetry_error`) end-to-end through the V2Event envelope.
func TestV2EventPromptAndRetryErrorDecode(t *testing.T) {
	t.Run("session.next.prompted", func(t *testing.T) {
		raw := `{"id":"evt_1","type":"session.next.prompted","data":{"timestamp":1,"messageID":"msg_1",` +
			`"sessionID":"ses_1","delivery":"immediate","prompt":{"text":"hi",` +
			`"files":[{"uri":"file:///a.txt","mime":"text/plain","name":"a.txt"}],` +
			`"agents":[{"name":"gopher"}]}}}`
		var e V2Event
		if err := json.Unmarshal([]byte(raw), &e); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		d, ok := e.Data.(V2EventSessionNextPromptedData)
		if !ok {
			t.Fatalf("Data runtime type = %T, want V2EventSessionNextPromptedData", e.Data)
		}
		if d.Prompt.Text != "hi" {
			t.Errorf("Prompt.Text = %q", d.Prompt.Text)
		}
		if len(d.Prompt.Files) != 1 || d.Prompt.Files[0].URI != "file:///a.txt" || d.Prompt.Files[0].Mime != "text/plain" {
			t.Errorf("Prompt.Files = %+v", d.Prompt.Files)
		}
		if len(d.Prompt.Agents) != 1 || d.Prompt.Agents[0].Name != "gopher" {
			t.Errorf("Prompt.Agents = %+v", d.Prompt.Agents)
		}
	})

	t.Run("session.next.prompt.admitted", func(t *testing.T) {
		raw := `{"id":"evt_2","type":"session.next.prompt.admitted","data":{"timestamp":2,"messageID":"msg_2",` +
			`"sessionID":"ses_1","delivery":"queued","prompt":{"text":"yo"}}}`
		var e V2Event
		if err := json.Unmarshal([]byte(raw), &e); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		d, ok := e.Data.(V2EventSessionNextPromptAdmittedData)
		if !ok {
			t.Fatalf("Data runtime type = %T, want V2EventSessionNextPromptAdmittedData", e.Data)
		}
		if d.Prompt.Text != "yo" {
			t.Errorf("Prompt.Text = %q", d.Prompt.Text)
		}
	})

	t.Run("session.next.retried", func(t *testing.T) {
		raw := `{"id":"evt_3","type":"session.next.retried","data":{"timestamp":3,"sessionID":"ses_1",` +
			`"attempt":2,"error":{"message":"429","isRetryable":true,"statusCode":429,` +
			`"responseHeaders":{"retry-after":"5"},"responseBody":"slow down",` +
			`"metadata":{"provider":"anthropic"}}}}`
		var e V2Event
		if err := json.Unmarshal([]byte(raw), &e); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		d, ok := e.Data.(V2EventSessionNextRetriedData)
		if !ok {
			t.Fatalf("Data runtime type = %T, want V2EventSessionNextRetriedData", e.Data)
		}
		if d.Error.Message != "429" || !d.Error.IsRetryable || d.Error.StatusCode != 429 {
			t.Errorf("Error = %+v", d.Error)
		}
		if d.Error.ResponseHeaders["retry-after"] != "5" || d.Error.ResponseBody != "slow down" {
			t.Errorf("Error = %+v", d.Error)
		}
		if d.Error.Metadata["provider"] != "anthropic" {
			t.Errorf("Error.Metadata = %+v", d.Error.Metadata)
		}
	})
}
