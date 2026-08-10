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

func TestTuiAppendPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.AppendPrompt(context.TODO(), opencode.TuiAppendPromptParams{
		Text:      opencode.F("text"),
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

func TestTuiClearPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.ClearPrompt(context.TODO(), opencode.TuiClearPromptParams{
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

func TestTuiExecuteCommandWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.ExecuteCommand(context.TODO(), opencode.TuiExecuteCommandParams{
		Command:   opencode.F("command"),
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

func TestTuiOpenHelpWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenHelp(context.TODO(), opencode.TuiOpenHelpParams{
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

func TestTuiOpenModelsWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenModels(context.TODO(), opencode.TuiOpenModelsParams{
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

func TestTuiOpenSessionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenSessions(context.TODO(), opencode.TuiOpenSessionsParams{
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

func TestTuiOpenThemesWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.OpenThemes(context.TODO(), opencode.TuiOpenThemesParams{
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

func TestTuiShowToastWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.ShowToast(context.TODO(), opencode.TuiShowToastParams{
		Message:   opencode.F("message"),
		Variant:   opencode.F(opencode.TuiShowToastParamsVariantInfo),
		Directory: opencode.F("directory"),
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

func TestTuiSubmitPromptWithOptionalParams(t *testing.T) {
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
	_, err := client.Tui.SubmitPrompt(context.TODO(), opencode.TuiSubmitPromptParams{
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

// TestTuiPublishBodyPromptAppendSerialization verifies that the PromptAppend variant
// serializes correctly with type discriminator "tui.prompt.append" and required fields.
func TestTuiPublishBodyPromptAppendSerialization(t *testing.T) {
	t.Parallel()
	params := opencode.TuiPublishParams{
		Body: opencode.F[opencode.TuiPublishBodyUnion](opencode.TuiPublishBodyPromptAppend{
			Type: opencode.F(opencode.TuiPublishBodyPromptAppendTypeTuiPromptAppend),
			Properties: opencode.F(opencode.TuiPublishBodyPromptAppendProperties{
				Text: opencode.F("hello world"),
			}),
		}),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if got["type"] != "tui.prompt.append" {
		t.Errorf("type: got %v, want %q", got["type"], "tui.prompt.append")
	}
	props, ok := got["properties"].(map[string]any)
	if !ok {
		t.Fatalf("properties: expected object, got %T", got["properties"])
	}
	if props["text"] != "hello world" {
		t.Errorf("properties.text: got %v, want %q", props["text"], "hello world")
	}
}

// TestTuiPublishBodyCommandExecuteSerialization verifies the CommandExecute variant.
func TestTuiPublishBodyCommandExecuteSerialization(t *testing.T) {
	t.Parallel()
	params := opencode.TuiPublishParams{
		Body: opencode.F[opencode.TuiPublishBodyUnion](opencode.TuiPublishBodyCommandExecute{
			Type: opencode.F(opencode.TuiPublishBodyCommandExecuteTypeTuiCommandExecute),
			Properties: opencode.F(opencode.TuiPublishBodyCommandExecuteProperties{
				Command: opencode.F("agent.cycle"),
			}),
		}),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if got["type"] != "tui.command.execute" {
		t.Errorf("type: got %v, want %q", got["type"], "tui.command.execute")
	}
	props, ok := got["properties"].(map[string]any)
	if !ok {
		t.Fatalf("properties: expected object, got %T", got["properties"])
	}
	if props["command"] != "agent.cycle" {
		t.Errorf("properties.command: got %v, want %q", props["command"], "agent.cycle")
	}
}

// TestTuiPublishBodyToastShowSerialization verifies the ToastShow variant.
func TestTuiPublishBodyToastShowSerialization(t *testing.T) {
	t.Parallel()
	params := opencode.TuiPublishParams{
		Body: opencode.F[opencode.TuiPublishBodyUnion](opencode.TuiPublishBodyToastShow{
			Type: opencode.F(opencode.TuiPublishBodyToastShowTypeTuiToastShow),
			Properties: opencode.F(opencode.TuiPublishBodyToastShowProperties{
				Message: opencode.F("Operation complete"),
				Variant: opencode.F(opencode.TuiPublishBodyToastShowPropertiesVariantSuccess),
				Title:   opencode.F("Done"),
			}),
		}),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if got["type"] != "tui.toast.show" {
		t.Errorf("type: got %v, want %q", got["type"], "tui.toast.show")
	}
	props, ok := got["properties"].(map[string]any)
	if !ok {
		t.Fatalf("properties: expected object, got %T", got["properties"])
	}
	if props["message"] != "Operation complete" {
		t.Errorf("properties.message: got %v, want %q", props["message"], "Operation complete")
	}
	if props["variant"] != "success" {
		t.Errorf("properties.variant: got %v, want %q", props["variant"], "success")
	}
	if props["title"] != "Done" {
		t.Errorf("properties.title: got %v, want %q", props["title"], "Done")
	}
}

// TestTuiPublishBodySessionSelectSerialization verifies the SessionSelect variant.
func TestTuiPublishBodySessionSelectSerialization(t *testing.T) {
	t.Parallel()
	params := opencode.TuiPublishParams{
		Body: opencode.F[opencode.TuiPublishBodyUnion](opencode.TuiPublishBodySessionSelect{
			Type: opencode.F(opencode.TuiPublishBodySessionSelectTypeTuiSessionSelect),
			Properties: opencode.F(opencode.TuiPublishBodySessionSelectProperties{
				SessionID: opencode.F("ses_abc123"),
			}),
		}),
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if got["type"] != "tui.session.select" {
		t.Errorf("type: got %v, want %q", got["type"], "tui.session.select")
	}
	props, ok := got["properties"].(map[string]any)
	if !ok {
		t.Fatalf("properties: expected object, got %T", got["properties"])
	}
	if props["sessionID"] != "ses_abc123" {
		t.Errorf("properties.sessionID: got %v, want %q", props["sessionID"], "ses_abc123")
	}
}
