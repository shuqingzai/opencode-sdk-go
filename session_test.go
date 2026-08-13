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
		// 载体合并枚举
		if msg.Format.Type != opencode.OutputFormatTypeText {
			t.Errorf("Format.Type: got %q, want %q", msg.Format.Type, opencode.OutputFormatTypeText)
		}
		if !msg.Format.Type.IsKnown() {
			t.Errorf("Format.Type %q is not known", msg.Format.Type)
		}
		// AsUnion() must return the concrete variant
		af, ok := msg.Format.AsUnion().(opencode.OutputFormatText)
		if !ok {
			t.Fatalf("Format.AsUnion(): got %T, want opencode.OutputFormatText", msg.Format.AsUnion())
		}
		if af.Type != opencode.OutputFormatTextTypeText {
			t.Errorf("AsUnion().Type: got %q, want %q", af.Type, opencode.OutputFormatTextTypeText)
		}
	})

	t.Run("json_schema variant decodes to OutputFormatJsonSchema", func(t *testing.T) {
		t.Parallel()
		payload := `{"type":"json_schema","schema":{"type":"object","properties":{"name":{"type":"string"}}},"retryCount":3}`
		var msg opencode.UserMessage
		if err := json.Unmarshal([]byte(baseMsg(payload)), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		// 载体扁平字段必须由 apijson.Port 填充
		if msg.Format.Type != opencode.OutputFormatTypeJsonSchema {
			t.Errorf("Format.Type: got %q, want %q", msg.Format.Type, opencode.OutputFormatTypeJsonSchema)
		}
		if msg.Format.RetryCount != 3 {
			t.Errorf("Format.RetryCount: got %d, want 3", msg.Format.RetryCount)
		}
		if _, ok := msg.Format.Schema["type"]; !ok {
			t.Error("Format.Schema should contain 'type' key")
		}
		fj, ok := msg.Format.AsUnion().(opencode.OutputFormatJsonSchema)
		if !ok {
			t.Fatalf("Format.AsUnion(): got %T, want opencode.OutputFormatJsonSchema", msg.Format.AsUnion())
		}
		if fj.RetryCount != 3 {
			t.Errorf("AsUnion().RetryCount: got %d, want 3", fj.RetryCount)
		}
		if _, ok := fj.Schema["type"]; !ok {
			t.Error("AsUnion().Schema should contain 'type' key")
		}
	})

	t.Run("null format field degrades to zero carrier", func(t *testing.T) {
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
		if msg.Format.AsUnion() != nil {
			t.Errorf("Format.AsUnion(): got %T, want nil", msg.Format.AsUnion())
		}
		if msg.Format.Type != "" {
			t.Errorf("Format.Type: got %q, want empty", msg.Format.Type)
		}
		if !msg.JSON.Format.IsNull() {
			t.Errorf("JSON.Format.IsNull() = false, want true (raw=%q)", msg.JSON.Format.Raw())
		}
	})

	t.Run("missing format field degrades to zero carrier", func(t *testing.T) {
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
		if msg.Format.AsUnion() != nil {
			t.Errorf("Format.AsUnion(): got %T, want nil", msg.Format.AsUnion())
		}
		if !msg.JSON.Format.IsMissing() {
			t.Errorf("JSON.Format.IsMissing() = false, want true")
		}
	})

	t.Run("other required fields unaffected by carrier decode", func(t *testing.T) {
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

// TestSessionPromptParamsFormatSerialization verifies that the `format` field of
// [SessionPromptParams] (used by both [SessionService.Prompt] and
// [SessionService.PromptAsync]) serializes correctly for both union variants,
// matching the OpenAPI `OutputFormat` anyOf = OutputFormatText | OutputFormatJsonSchema.
//
// Run with: go test -run TestSessionPromptParamsFormatSerialization -v ./...
func TestSessionPromptParamsFormatSerialization(t *testing.T) {
	t.Parallel()

	t.Run("text variant", func(t *testing.T) {
		t.Parallel()
		p := opencode.SessionPromptParams{
			Parts: opencode.F([]opencode.SessionPromptParamsPartUnion{opencode.TextPartInputParam{
				Type: opencode.F(opencode.TextPartInputTypeText),
				Text: opencode.F("hi"),
			}}),
			Format: opencode.F[opencode.SessionPromptParamsFormatUnion](
				opencode.SessionPromptParamsFormatText{
					Type: opencode.F(opencode.SessionPromptParamsFormatTextTypeText),
				},
			),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		if got, want := string(data), `{"format":{"type":"text"},"parts":[{"text":"hi","type":"text"}]}`; got != want {
			t.Errorf("JSON:\n got %s\nwant %s", got, want)
		}
	})

	t.Run("json_schema variant", func(t *testing.T) {
		t.Parallel()
		p := opencode.SessionPromptParams{
			Parts: opencode.F([]opencode.SessionPromptParamsPartUnion{opencode.TextPartInputParam{
				Type: opencode.F(opencode.TextPartInputTypeText),
				Text: opencode.F("hi"),
			}}),
			Format: opencode.F[opencode.SessionPromptParamsFormatUnion](
				opencode.SessionPromptParamsFormatJsonSchema{
					Type:       opencode.F(opencode.SessionPromptParamsFormatJsonSchemaTypeJsonSchema),
					Schema:     opencode.F[any](map[string]any{"type": "object"}),
					RetryCount: opencode.F(int64(3)),
				},
			),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		var got map[string]any
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal roundtrip: %v", err)
		}
		f, ok := got["format"].(map[string]any)
		if !ok {
			t.Fatalf("format: got %T, want object", got["format"])
		}
		if f["type"] != "json_schema" {
			t.Errorf("format.type: got %v, want %q", f["type"], "json_schema")
		}
		if f["retryCount"] != float64(3) {
			t.Errorf("format.retryCount: got %v, want 3", f["retryCount"])
		}
		schema, ok := f["schema"].(map[string]any)
		if !ok {
			t.Fatalf("format.schema: got %T, want object", f["schema"])
		}
		if schema["type"] != "object" {
			t.Errorf("format.schema.type: got %v, want %q", schema["type"], "object")
		}
	})

	t.Run("format omitted when unset", func(t *testing.T) {
		t.Parallel()
		p := opencode.SessionPromptParams{
			Parts: opencode.F([]opencode.SessionPromptParamsPartUnion{opencode.TextPartInputParam{
				Type: opencode.F(opencode.TextPartInputTypeText),
				Text: opencode.F("hi"),
			}}),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		if len(data) == 0 {
			t.Fatal("empty JSON output")
		}
		var got map[string]any
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("Unmarshal roundtrip: %v", err)
		}
		if _, present := got["format"]; present {
			t.Errorf("format should be omitted, got %v", got["format"])
		}
	})

	t.Run("Prompt and PromptAsync use the same SessionPromptParams body", func(t *testing.T) {
		t.Parallel()
		// Both methods accept SessionPromptParams; the body wire format must be
		// identical regardless of which endpoint consumes it.
		p := opencode.SessionPromptParams{
			Parts: opencode.F([]opencode.SessionPromptParamsPartUnion{opencode.TextPartInputParam{
				Type: opencode.F(opencode.TextPartInputTypeText),
				Text: opencode.F("hi"),
			}}),
			Format: opencode.F[opencode.SessionPromptParamsFormatUnion](
				opencode.SessionPromptParamsFormatText{
					Type: opencode.F(opencode.SessionPromptParamsFormatTextTypeText),
				},
			),
		}
		data, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		want := `{"format":{"type":"text"},"parts":[{"text":"hi","type":"text"}]}`
		if got := string(data); got != want {
			t.Errorf("JSON:\n got %s\nwant %s", got, want)
		}
	})
}

// TestMessageMergedStructFieldTypes verifies that the merged [Message] struct
// decodes the variant-consistent fields with their concrete types (System is
// always `string` per OpenAPI `UserMessage.system`) while keeping the
// variant-conditional union fields as `any` with runtime comments.
//
// Run with: go test -run TestMessageMergedStructFieldTypes -v ./...
func TestMessageMergedStructFieldTypes(t *testing.T) {
	t.Parallel()

	t.Run("UserMessage system decodes into string", func(t *testing.T) {
		t.Parallel()
		raw := `{
			"id": "msg_001",
			"sessionID": "ses_001",
			"role": "user",
			"time": {"created": 123},
			"agent": "test",
			"model": {"providerID": "openai", "modelID": "gpt-4"},
			"system": "you are a helpful assistant"
		}`
		var msg opencode.Message
		if err := json.Unmarshal([]byte(raw), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if msg.System != "you are a helpful assistant" {
			t.Errorf("System: got %q, want %q", msg.System, "you are a helpful assistant")
		}
		if msg.Role != opencode.MessageRoleUser {
			t.Errorf("Role: got %q, want %q", msg.Role, opencode.MessageRoleUser)
		}
	})

	t.Run("AssistantMessage without system leaves System empty string", func(t *testing.T) {
		t.Parallel()
		raw := `{
			"id": "msg_002",
			"sessionID": "ses_001",
			"role": "assistant",
			"time": {"created": 123},
			"agent": "test",
			"modelID": "gpt-4",
			"providerID": "openai",
			"mode": "auto",
			"path": {"cwd": "/work", "root": "/work"},
			"cost": 0.01,
			"tokens": {"input": 10, "output": 20}
		}`
		var msg opencode.Message
		if err := json.Unmarshal([]byte(raw), &msg); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if msg.System != "" {
			t.Errorf("System: got %q, want empty string", msg.System)
		}
		if msg.Tokens == nil {
			t.Error("Tokens: should be populated for AssistantMessage")
		}
	})
}

// TestToolPartStateCarrierFieldEdgeCases (疑点 3) verifies that the merged
// [ToolPartState] carrier preserves `!= nil` semantics after `Input`/`Metadata`
// were tightened from `any` to `map[string]any`: a variant that omits the field,
// or a JSON `null`, must leave the carrier field nil (not an empty map), and a
// present value must decode intact.
//
// Run with: go test -run TestToolPartStateCarrierFieldEdgeCases -v ./...
func TestToolPartStateCarrierFieldEdgeCases(t *testing.T) {
	t.Parallel()

	decode := func(t *testing.T, raw string) opencode.ToolPart {
		t.Helper()
		var p opencode.ToolPart
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		return p
	}

	t.Run("pending omits metadata/output/time/title", func(t *testing.T) {
		t.Parallel()
		p := decode(t, `{
			"id": "prt_1", "sessionID": "ses_1", "messageID": "msg_1",
			"type": "tool", "callID": "call_1", "tool": "bash",
			"state": {"status": "pending", "input": {"command": "ls"}, "raw": "x"}
		}`)
		st := p.State
		if _, ok := st.AsUnion().(opencode.ToolStatePending); !ok {
			t.Fatalf("state union runtime type = %T, want ToolStatePending", st.AsUnion())
		}
		if st.Input == nil {
			t.Error("Input = nil, want non-nil map for pending input")
		}
		if st.Input["command"] != "ls" {
			t.Errorf(`Input["command"] = %v, want "ls"`, st.Input["command"])
		}
		if st.Metadata != nil {
			t.Errorf("Metadata = %#v, want nil (pending variant has no metadata)", st.Metadata)
		}
		if st.Output != "" {
			t.Errorf("Output = %q, want empty (pending variant has no output)", st.Output)
		}
		if st.Title != "" {
			t.Errorf("Title = %q, want empty (pending variant has no title)", st.Title)
		}
		if st.Raw != "x" {
			t.Errorf("Raw = %q, want x", st.Raw)
		}
		if st.Time != nil {
			t.Errorf("Time = %#v, want nil (pending variant has no time)", st.Time)
		}
	})

	t.Run("null input and metadata stay nil", func(t *testing.T) {
		t.Parallel()
		p := decode(t, `{
			"id": "prt_1", "sessionID": "ses_1", "messageID": "msg_1",
			"type": "tool", "callID": "call_1", "tool": "bash",
			"state": {"status": "running", "input": null, "metadata": null,
				"title": "t", "time": {"start": 1}}
		}`)
		st := p.State
		if _, ok := st.AsUnion().(opencode.ToolStateRunning); !ok {
			t.Fatalf("state union runtime type = %T, want ToolStateRunning", st.AsUnion())
		}
		if st.Input != nil {
			t.Errorf("Input = %#v, want nil for JSON null", st.Input)
		}
		if st.Metadata != nil {
			t.Errorf("Metadata = %#v, want nil for JSON null", st.Metadata)
		}
		if !st.JSON.Input.IsNull() {
			t.Error("JSON.Input.IsNull() = false, want true")
		}
		if !st.JSON.Metadata.IsNull() {
			t.Error("JSON.Metadata.IsNull() = false, want true")
		}
		if st.Title != "t" {
			t.Errorf("Title = %q, want t", st.Title)
		}
	})

	t.Run("completed carries all fields", func(t *testing.T) {
		t.Parallel()
		p := decode(t, `{
			"id": "prt_1", "sessionID": "ses_1", "messageID": "msg_1",
			"type": "tool", "callID": "call_1", "tool": "bash",
			"state": {"status": "completed",
				"input": {"command": "ls"}, "metadata": {"cwd": "/w"},
				"output": "ok", "title": "run", "time": {"start": 1, "end": 2}}
		}`)
		st := p.State
		if _, ok := st.AsUnion().(opencode.ToolStateCompleted); !ok {
			t.Fatalf("state union runtime type = %T, want ToolStateCompleted", st.AsUnion())
		}
		if st.Input == nil || st.Input["command"] != "ls" {
			t.Errorf("Input = %#v, want {command: ls}", st.Input)
		}
		if st.Metadata == nil || st.Metadata["cwd"] != "/w" {
			t.Errorf("Metadata = %#v, want {cwd: /w}", st.Metadata)
		}
		if st.Output != "ok" || st.Title != "run" {
			t.Errorf("Output/Title = %q/%q, want ok/run", st.Output, st.Title)
		}
		if st.Raw != "" {
			t.Errorf("Raw = %q, want empty (completed variant has no raw)", st.Raw)
		}
	})
}

// TestMessageSystemFieldEdgeCases (疑点 3) verifies the merged [Message]
// carrier's `System` (tightened from `any` to `string`) across the user/assistant
// variants, including the JSON `null` boundary.
func TestMessageSystemFieldEdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("user with system", func(t *testing.T) {
		t.Parallel()
		var m opencode.Message
		raw := `{"id":"msg_1","sessionID":"ses_1","role":"user","time":{"created":1},
			"agent":"a","model":{"providerID":"p","modelID":"m"},"system":"sys"}`
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if m.System != "sys" {
			t.Errorf("System = %q, want sys", m.System)
		}
		if m.JSON.System.IsMissing() || m.JSON.System.IsNull() {
			t.Error("JSON.System should be present and valid, got missing/null")
		}
	})

	t.Run("user with system null", func(t *testing.T) {
		t.Parallel()
		var m opencode.Message
		raw := `{"id":"msg_1","sessionID":"ses_1","role":"user","time":{"created":1},
			"agent":"a","model":{"providerID":"p","modelID":"m"},"system":null}`
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if m.System != "" {
			t.Errorf("System = %q, want empty for JSON null", m.System)
		}
		if !m.JSON.System.IsNull() {
			t.Error("JSON.System.IsNull() = false, want true")
		}
	})

	t.Run("assistant has no system", func(t *testing.T) {
		t.Parallel()
		var m opencode.Message
		raw := `{"id":"msg_2","sessionID":"ses_1","role":"assistant","time":{"created":1},
			"agent":"a","modelID":"m","providerID":"p","mode":"auto",
			"path":{"cwd":"/w","root":"/w"},"tokens":{"input":1,"output":2}}`
		if err := json.Unmarshal([]byte(raw), &m); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if m.System != "" {
			t.Errorf("System = %q, want empty for assistant", m.System)
		}
		if m.JSON.System.IsMissing() == false {
			t.Error("JSON.System.IsMissing() = false, want true (assistant has no system field)")
		}
	})
}

// TestPartMetadataFieldEdgeCases (疑点 3) verifies the merged [Part] carrier's
// `Metadata` (tightened from `any` to `map[string]any`) stays nil when the
// variant has no metadata and decodes when present.
func TestPartMetadataFieldEdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("text part with metadata", func(t *testing.T) {
		t.Parallel()
		var p opencode.Part
		raw := `{"id":"prt_1","sessionID":"ses_1","messageID":"msg_1","type":"text",
			"text":"hi","metadata":{"mode":"auto"}}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if p.Metadata == nil || p.Metadata["mode"] != "auto" {
			t.Errorf("Metadata = %#v, want {mode: auto}", p.Metadata)
		}
	})

	t.Run("subtask part without metadata", func(t *testing.T) {
		t.Parallel()
		var p opencode.Part
		raw := `{"id":"prt_2","sessionID":"ses_1","messageID":"msg_1","type":"subtask",
			"prompt":"p","description":"d","agent":"a"}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if p.Metadata != nil {
			t.Errorf("Metadata = %#v, want nil", p.Metadata)
		}
	})

	t.Run("metadata null", func(t *testing.T) {
		t.Parallel()
		var p opencode.Part
		raw := `{"id":"prt_3","sessionID":"ses_1","messageID":"msg_1","type":"text",
			"text":"hi","metadata":null}`
		if err := json.Unmarshal([]byte(raw), &p); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if p.Metadata != nil {
			t.Errorf("Metadata = %#v, want nil for JSON null", p.Metadata)
		}
		if !p.JSON.Metadata.IsNull() {
			t.Error("JSON.Metadata.IsNull() = false, want true")
		}
	})
}

// Second-pass independent verification: ToolPartState / Message / Part carrier
// fields tightened from any to concrete types must preserve `!= nil` semantics
// across union variants (missing / null / present).
func TestReview2_MessageSystemPortSemantics(t *testing.T) {
	t.Parallel()
	// UserMessage variant: system present
	var m opencode.Message
	if err := json.Unmarshal([]byte(`{"id":"msg_1","sessionID":"ses_1","role":"user","time":{"created":1},"agent":"a","model":{"providerID":"p","modelID":"m"},"system":"sys"}`), &m); err != nil {
		t.Fatalf("user w/ system: %v", err)
	}
	if m.System != "sys" || m.JSON.System.IsMissing() || m.JSON.System.IsNull() {
		t.Errorf("user w/ system: System=%q meta missing=%v null=%v", m.System, m.JSON.System.IsMissing(), m.JSON.System.IsNull())
	}

	// AssistantMessage variant: no system field at all -> must stay ""
	var a opencode.Message
	if err := json.Unmarshal([]byte(`{"id":"msg_2","sessionID":"ses_1","role":"assistant","time":{"created":1},"agent":"a","modelID":"m","providerID":"p","mode":"auto","path":{"cwd":"/w","root":"/w"},"cost":1,"tokens":{"input":1,"output":2,"reasoning":0,"cache":{"read":0,"write":0}},"parentID":"msg_x"}`), &a); err != nil {
		t.Fatalf("assistant: %v", err)
	}
	if a.System != "" {
		t.Errorf("assistant: System=%q want empty", a.System)
	}
	if !a.JSON.System.IsMissing() {
		t.Errorf("assistant: JSON.System.IsMissing()=false want true (got null=%v)", a.JSON.System.IsNull())
	}

	// UserMessage variant: system explicitly null -> "", IsNull
	var n opencode.Message
	if err := json.Unmarshal([]byte(`{"id":"msg_3","sessionID":"ses_1","role":"user","time":{"created":1},"agent":"a","model":{"providerID":"p","modelID":"m"},"system":null}`), &n); err != nil {
		t.Fatalf("user w/ null system: %v", err)
	}
	if n.System != "" || !n.JSON.System.IsNull() || n.JSON.System.IsMissing() {
		t.Errorf("user w/ null: System=%q null=%v missing=%v", n.System, n.JSON.System.IsNull(), n.JSON.System.IsMissing())
	}
}

func TestReview2_ToolPartStateInputMetadataPortSemantics(t *testing.T) {
	t.Parallel()
	// pending variant: has input, NO metadata/output/time/title/error
	var p opencode.ToolPart
	if err := json.Unmarshal([]byte(`{"id":"p","sessionID":"s","messageID":"m","type":"tool","callID":"c","tool":"bash","state":{"status":"pending","input":{"command":"ls"},"raw":"r"}}`), &p); err != nil {
		t.Fatalf("pending: %v", err)
	}
	st := p.State
	if _, ok := st.AsUnion().(opencode.ToolStatePending); !ok {
		t.Fatalf("union type=%T want pending", st.AsUnion())
	}
	if st.Input == nil {
		t.Error("pending Input nil, want map (present)")
	}
	if st.Metadata != nil {
		t.Errorf("pending Metadata=%#v want nil (variant lacks field)", st.Metadata)
	}
	if st.Output != "" || st.Title != "" || st.Time != nil || st.Error != "" || st.Attachments != nil {
		t.Errorf("pending stray fields: out=%q title=%q time=%v err=%q att=%v", st.Output, st.Title, st.Time, st.Error, st.Attachments)
	}

	// running variant with null input/metadata -> nil + IsNull
	var r opencode.ToolPart
	if err := json.Unmarshal([]byte(`{"id":"p","sessionID":"s","messageID":"m","type":"tool","callID":"c","tool":"bash","state":{"status":"running","input":null,"metadata":null,"title":"t","time":{"start":1}}}`), &r); err != nil {
		t.Fatalf("running null: %v", err)
	}
	rs := r.State
	if rs.Input != nil || rs.Metadata != nil {
		t.Errorf("running null: Input=%#v Metadata=%#v want nil", rs.Input, rs.Metadata)
	}
	if !rs.JSON.Input.IsNull() || !rs.JSON.Metadata.IsNull() {
		t.Errorf("running null: JSON.Input null=%v JSON.Metadata null=%v want both true", rs.JSON.Input.IsNull(), rs.JSON.Metadata.IsNull())
	}
	if rs.Title != "t" {
		t.Errorf("running null: Title=%q want t", rs.Title)
	}

	// completed variant: all present
	var c opencode.ToolPart
	if err := json.Unmarshal([]byte(`{"id":"p","sessionID":"s","messageID":"m","type":"tool","callID":"c","tool":"bash","state":{"status":"completed","input":{"a":1},"metadata":{"b":"c"},"output":"o","title":"tt","time":{"start":1,"end":2}}}`), &c); err != nil {
		t.Fatalf("completed: %v", err)
	}
	cs := c.State
	if cs.Input == nil || cs.Input["a"] == nil || cs.Metadata == nil || cs.Metadata["b"] != "c" || cs.Output != "o" || cs.Title != "tt" {
		t.Errorf("completed: Input=%#v Metadata=%#v Output=%q Title=%q", cs.Input, cs.Metadata, cs.Output, cs.Title)
	}
}

func TestReview2_PartMetadataPortSemantics(t *testing.T) {
	t.Parallel()
	// subtask part: variant lacks metadata -> nil (not empty map)
	var p opencode.Part
	if err := json.Unmarshal([]byte(`{"id":"p","sessionID":"s","messageID":"m","type":"subtask","prompt":"x","description":"d","agent":"a"}`), &p); err != nil {
		t.Fatalf("subtask: %v", err)
	}
	if p.Metadata != nil {
		t.Errorf("subtask Metadata=%#v want nil", p.Metadata)
	}
	// text part: metadata present
	var tp opencode.Part
	if err := json.Unmarshal([]byte(`{"id":"p","sessionID":"s","messageID":"m","type":"text","text":"hi","metadata":{"k":"v"}}`), &tp); err != nil {
		t.Fatalf("text: %v", err)
	}
	if tp.Metadata == nil || tp.Metadata["k"] != "v" {
		t.Errorf("text Metadata=%#v want {k:v}", tp.Metadata)
	}
}

func TestReview2_SessionMetadataDecode(t *testing.T) {
	t.Parallel()
	var s opencode.Session
	if err := json.Unmarshal([]byte(`{"id":"ses_1","workspaceID":"wrk_1","metadata":{"a":"b"}}`), &s); err != nil {
		t.Fatalf("session: %v", err)
	}
	if s.Metadata == nil || s.Metadata["a"] != "b" {
		t.Errorf("Metadata=%#v want {a:b}", s.Metadata)
	}
}
