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
	"github.com/sst/opencode-sdk-go/internal/testutil"
	"github.com/sst/opencode-sdk-go/option"
)

func TestSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.New(context.TODO(), opencode.SessionNewParams{
		Directory: opencode.F("directory"),
		ParentID:  opencode.F("sesJ!"),
		Title:     opencode.F("title"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Update(
		context.TODO(),
		"id",
		opencode.SessionUpdateParams{
			Directory: opencode.F("directory"),
			Title:     opencode.F("title"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.List(context.TODO(), opencode.SessionListParams{
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

func TestSessionDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Delete(
		context.TODO(),
		"sesJ!",
		opencode.SessionDeleteParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionAbortWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Abort(
		context.TODO(),
		"id",
		opencode.SessionAbortParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionChildrenWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Children(
		context.TODO(),
		"sesJ!",
		opencode.SessionChildrenParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionCommandWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Command(
		context.TODO(),
		"id",
		opencode.SessionCommandParams{
			Arguments: opencode.F("arguments"),
			Command:   opencode.F("command"),
			Directory: opencode.F("directory"),
			Agent:     opencode.F("agent"),
			MessageID: opencode.F("msgJ!"),
			Model:     opencode.F("model"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Get(
		context.TODO(),
		"sesJ!",
		opencode.SessionGetParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionInitWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Init(
		context.TODO(),
		"id",
		opencode.SessionInitParams{
			MessageID:  opencode.F("msgJ!"),
			ModelID:    opencode.F("modelID"),
			ProviderID: opencode.F("providerID"),
			Directory:  opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Message(
		context.TODO(),
		"id",
		"messageID",
		opencode.SessionMessageParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessagesWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Messages(
		context.TODO(),
		"id",
		opencode.SessionMessagesParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Prompt(
		context.TODO(),
		"id",
		opencode.SessionPromptParams{
			Parts: opencode.F([]opencode.SessionPromptParamsPartUnion{opencode.TextPartInputParam{
				Text: opencode.F("text"),
				Type: opencode.F(opencode.TextPartInputTypeText),
				ID:   opencode.F("id"),
				Metadata: opencode.F(map[string]any{
					"foo": "bar",
				}),
				Synthetic: opencode.F(true),
				Time: opencode.F(opencode.TextPartInputTimeParam{
					Start: opencode.F(int64(0)),
					End:   opencode.F(int64(0)),
				}),
			}}),
			Directory: opencode.F("directory"),
			Agent:     opencode.F("agent"),
			MessageID: opencode.F("msgJ!"),
			Model: opencode.F(opencode.SessionPromptParamsModel{
				ModelID:    opencode.F("modelID"),
				ProviderID: opencode.F("providerID"),
			}),
			NoReply: opencode.F(true),
			System:  opencode.F("system"),
			Tools: opencode.F(map[string]bool{
				"foo": true,
			}),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionRevertWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Revert(
		context.TODO(),
		"id",
		opencode.SessionRevertParams{
			MessageID: opencode.F("msgJ!"),
			Directory: opencode.F("directory"),
			PartID:    opencode.F("prtJ!"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionShareWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Share(
		context.TODO(),
		"id",
		opencode.SessionShareParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionShellWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Shell(
		context.TODO(),
		"id",
		opencode.SessionShellParams{
			Agent:     opencode.F("agent"),
			Command:   opencode.F("command"),
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionSummarizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Summarize(
		context.TODO(),
		"id",
		opencode.SessionSummarizeParams{
			ModelID:    opencode.F("modelID"),
			ProviderID: opencode.F("providerID"),
			Directory:  opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUnrevertWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Unrevert(
		context.TODO(),
		"id",
		opencode.SessionUnrevertParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionUnshareWithOptionalParams(t *testing.T) {
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
	_, err := client.Session.Unshare(
		context.TODO(),
		"sesJ!",
		opencode.SessionUnshareParams{
			Directory: opencode.F("directory"),
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

// TestPartRetryPartErrorDataMetadataUnmarshal verifies that PartRetryPartErrorData
// can deserialize a JSON payload that includes the `metadata` field
// (OpenAPI: components.schemas.APIError.properties.data.properties.metadata).
//
// Run with: go test -run TestPartRetryPartErrorDataMetadataUnmarshal -v ./...
func TestPartRetryPartErrorDataMetadataUnmarshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		jsonInput  string
		wantMeta   map[string]string
		wantMsg    string
		wantRetry  bool
		wantStatus int64
	}{
		// Variant 1: metadata present with multiple entries
		{
			name:       "metadata_present",
			jsonInput:  `{"isRetryable":true,"message":"rate limit exceeded","statusCode":429,"metadata":{"x-request-id":"abc123","x-ratelimit-reset":"60"}}`,
			wantMeta:   map[string]string{"x-request-id": "abc123", "x-ratelimit-reset": "60"},
			wantMsg:    "rate limit exceeded",
			wantRetry:  true,
			wantStatus: 429,
		},
		// Variant 2: metadata absent (optional field)
		{
			name:       "metadata_absent",
			jsonInput:  `{"isRetryable":false,"message":"not found","statusCode":404}`,
			wantMeta:   nil,
			wantMsg:    "not found",
			wantRetry:  false,
			wantStatus: 404,
		},
		// Variant 3: metadata empty object
		{
			name:       "metadata_empty",
			jsonInput:  `{"isRetryable":true,"message":"server error","metadata":{}}`,
			wantMeta:   map[string]string{},
			wantMsg:    "server error",
			wantRetry:  true,
			wantStatus: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var d opencode.PartRetryPartErrorData
			if err := json.Unmarshal([]byte(tc.jsonInput), &d); err != nil {
				t.Fatalf("json.Unmarshal: %v", err)
			}
			if d.Message != tc.wantMsg {
				t.Errorf("Message: got %q, want %q", d.Message, tc.wantMsg)
			}
			if d.IsRetryable != tc.wantRetry {
				t.Errorf("IsRetryable: got %v, want %v", d.IsRetryable, tc.wantRetry)
			}
			if d.StatusCode != tc.wantStatus {
				t.Errorf("StatusCode: got %v, want %v", d.StatusCode, tc.wantStatus)
			}
			if !reflect.DeepEqual(d.Metadata, tc.wantMeta) {
				t.Errorf("Metadata: got %v, want %v", d.Metadata, tc.wantMeta)
			}
			// Verify JSON metadata apijson.Field is populated when metadata is present
			if tc.wantMeta != nil && len(tc.wantMeta) > 0 {
				if d.JSON.Metadata.Raw() == "" {
					t.Error("JSON.Metadata.Raw() should be non-empty when metadata is present")
				}
			}
		})
	}
}

// TestSessionListParamsScopeEnum verifies that SessionListParamsScope enum
// contains the correct values as defined in OpenAPI
// (paths./session.get.parameters.scope.schema.enum).
//
// Run with: go test -run TestSessionListParamsScopeEnum -v ./...
func TestSessionListParamsScopeEnum(t *testing.T) {
	t.Parallel()
	// OpenAPI defines scope as string enum["project"]
	scope := opencode.SessionListParamsScopeProject
	if scope != "project" {
		t.Errorf("SessionListParamsScopeProject: got %q, want %q", scope, "project")
	}
	if !scope.IsKnown() {
		t.Error("SessionListParamsScopeProject.IsKnown() should return true")
	}
	// An unknown value should return false
	unknown := opencode.SessionListParamsScope("unknown")
	if unknown.IsKnown() {
		t.Error("unknown scope.IsKnown() should return false")
	}
}

// TestUserMessageFormatUnionDecoding verifies that [UserMessage.Format] is
// correctly decoded into the typed Response union variants
// ([OutputFormatText] and [OutputFormatJsonSchema]) rather than
// falling back to map[string]any.
//
// Root cause fixed: apijson routes by the static type of the field; an `any`
// field always resolves to interface{} → generic map. The custom UnmarshalJSON
// on UserMessage now explicitly decodes the "format" key through the registered
// OutputFormatUnion, so the typed assertion succeeds.
//
// Run with: go test -run TestUserMessageFormatUnionDecoding -v ./...
func TestUserMessageFormatUnionDecoding(t *testing.T) {
	t.Parallel()

	baseMsg := func(formatJSON string) string {
		return `{
			"id": "msg_001",
			"sessionID": "ses_001",
			"role": "user",
			"time": {"created": 1234567890},
			"agent": "test",
			"model": {"providerID": "openai", "modelID": "gpt-4"},
			"format": ` + formatJSON + `
		}`
	}

	t.Run("text variant decodes to OutputFormatText", func(t *testing.T) {
		t.Parallel()
		var msg opencode.UserMessage
		if err := json.Unmarshal([]byte(baseMsg(`{"type":"text"}`)), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		// runtime type must be OutputFormatText
		ft, ok := msg.Format.(opencode.OutputFormatText)
		if !ok {
			t.Fatalf("Format: got %T, want opencode.OutputFormatText", msg.Format)
		}
		if ft.Type != "text" {
			t.Errorf("Type: got %q, want %q", ft.Type, "text")
		}
		// AsFormat() must return the same typed value
		af, ok2 := msg.AsFormat().(opencode.OutputFormatText)
		if !ok2 {
			t.Fatalf("AsFormat(): got %T, want opencode.OutputFormatText", msg.AsFormat())
		}
		if af.Type != "text" {
			t.Errorf("AsFormat().Type: got %q, want %q", af.Type, "text")
		}
	})

	t.Run("json_schema variant decodes to OutputFormatJsonSchema", func(t *testing.T) {
		t.Parallel()
		payload := `{"type":"json_schema","schema":{"type":"object","properties":{"name":{"type":"string"}}},"retryCount":3}`
		var msg opencode.UserMessage
		if err := json.Unmarshal([]byte(baseMsg(payload)), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		fj, ok := msg.Format.(opencode.OutputFormatJsonSchema)
		if !ok {
			t.Fatalf("Format: got %T, want opencode.OutputFormatJsonSchema", msg.Format)
		}
		if fj.Type != "json_schema" {
			t.Errorf("Type: got %q, want %q", fj.Type, "json_schema")
		}
		if fj.RetryCount != 3 {
			t.Errorf("RetryCount: got %d, want 3", fj.RetryCount)
		}
		if _, ok := fj.Schema["type"]; !ok {
			t.Error("Schema should contain 'type' key")
		}
		// AsFormat() must return the same typed value
		if _, ok2 := msg.AsFormat().(opencode.OutputFormatJsonSchema); !ok2 {
			t.Fatalf("AsFormat(): got %T, want opencode.OutputFormatJsonSchema", msg.AsFormat())
		}
	})

	t.Run("null format field leaves Format nil", func(t *testing.T) {
		t.Parallel()
		raw := `{
			"id": "msg_002",
			"sessionID": "ses_001",
			"role": "user",
			"time": {"created": 1234567890},
			"agent": "test",
			"model": {"providerID": "openai", "modelID": "gpt-4"},
			"format": null
		}`
		var msg opencode.UserMessage
		if err := json.Unmarshal([]byte(raw), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if msg.Format != nil {
			t.Errorf("Format: got %T, want nil", msg.Format)
		}
		if msg.AsFormat() != nil {
			t.Errorf("AsFormat(): got %T, want nil", msg.AsFormat())
		}
	})

	t.Run("missing format field leaves Format nil", func(t *testing.T) {
		t.Parallel()
		raw := `{
			"id": "msg_003",
			"sessionID": "ses_001",
			"role": "user",
			"time": {"created": 1234567890},
			"agent": "test",
			"model": {"providerID": "openai", "modelID": "gpt-4"}
		}`
		var msg opencode.UserMessage
		if err := json.Unmarshal([]byte(raw), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if msg.Format != nil {
			t.Errorf("Format: got %T, want nil", msg.Format)
		}
	})

	t.Run("other required fields unaffected by custom UnmarshalJSON", func(t *testing.T) {
		t.Parallel()
		var msg opencode.UserMessage
		if err := json.Unmarshal([]byte(baseMsg(`{"type":"text"}`)), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if msg.ID != "msg_001" {
			t.Errorf("ID: got %q, want %q", msg.ID, "msg_001")
		}
		if msg.Agent != "test" {
			t.Errorf("Agent: got %q, want %q", msg.Agent, "test")
		}
		if msg.Model.ProviderID != "openai" {
			t.Errorf("Model.ProviderID: got %q, want %q", msg.Model.ProviderID, "openai")
		}
		if msg.Model.ModelID != "gpt-4" {
			t.Errorf("Model.ModelID: got %q, want %q", msg.Model.ModelID, "gpt-4")
		}
	})
}

// TestAssistantMessageStructuredIsUnknown verifies that [AssistantMessage.Structured]
// accepts arbitrary JSON values, consistent with the OpenAPI schema `{}` (empty =
// unknown). The field must NOT be typed as a fixed union — it may hold any shape.
//
// Run with: go test -run TestAssistantMessageStructuredIsUnknown -v ./...
func TestAssistantMessageStructuredIsUnknown(t *testing.T) {
	t.Parallel()

	baseMsg := func(structuredJSON string) string {
		return `{
			"id": "msg_002",
			"sessionID": "ses_001",
			"role": "assistant",
			"time": {"created": 1234567890},
			"parentID": "msg_001",
			"modelID": "gpt-4",
			"providerID": "openai",
			"mode": "auto",
			"agent": "test",
			"path": {"cwd": "/work", "root": "/work"},
			"cost": 0.01,
			"tokens": {"input": 10, "output": 20, "reasoning": 0, "cache": {"read": 0, "write": 0}},
			"structured": ` + structuredJSON + `
		}`
	}

	t.Run("object value decodes to map[string]any", func(t *testing.T) {
		t.Parallel()
		var msg opencode.AssistantMessage
		if err := json.Unmarshal([]byte(baseMsg(`{"key":"value","num":42}`)), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		m, ok := msg.Structured.(map[string]any)
		if !ok {
			t.Fatalf("Structured: got %T, want map[string]any", msg.Structured)
		}
		if m["key"] != "value" {
			t.Errorf("Structured[key]: got %v, want %q", m["key"], "value")
		}
	})

	t.Run("string value decodes to string", func(t *testing.T) {
		t.Parallel()
		var msg opencode.AssistantMessage
		if err := json.Unmarshal([]byte(baseMsg(`"hello"`)), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if s, ok := msg.Structured.(string); !ok || s != "hello" {
			t.Errorf("Structured: got %T(%v), want string(hello)", msg.Structured, msg.Structured)
		}
	})

	t.Run("array value decodes to []any", func(t *testing.T) {
		t.Parallel()
		var msg opencode.AssistantMessage
		if err := json.Unmarshal([]byte(baseMsg(`[1,2,3]`)), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if _, ok := msg.Structured.([]any); !ok {
			t.Fatalf("Structured: got %T, want []any", msg.Structured)
		}
	})

	t.Run("null value leaves Structured nil", func(t *testing.T) {
		t.Parallel()
		var msg opencode.AssistantMessage
		if err := json.Unmarshal([]byte(baseMsg(`null`)), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if msg.Structured != nil {
			t.Errorf("Structured: got %T, want nil", msg.Structured)
		}
	})

	t.Run("missing structured field leaves Structured nil", func(t *testing.T) {
		t.Parallel()
		raw := `{
			"id": "msg_002",
			"sessionID": "ses_001",
			"role": "assistant",
			"time": {"created": 1234567890},
			"parentID": "msg_001",
			"modelID": "gpt-4",
			"providerID": "openai",
			"mode": "auto",
			"agent": "test",
			"path": {"cwd": "/work", "root": "/work"},
			"cost": 0.01,
			"tokens": {"input": 10, "output": 20, "reasoning": 0, "cache": {"read": 0, "write": 0}}
		}`
		var msg opencode.AssistantMessage
		if err := json.Unmarshal([]byte(raw), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if msg.Structured != nil {
			t.Errorf("Structured: got %T, want nil", msg.Structured)
		}
	})
}
