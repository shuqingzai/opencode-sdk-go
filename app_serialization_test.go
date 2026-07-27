package opencode

import (
	"encoding/json"
	"strings"
	"testing"
)

// Aligned with OpenAPI /log requestBody + JS SDK(v2) App.log.
// body required: service, level, message; optional: extra
// query: directory, workspace (must NOT appear in body)
func TestAppLogParamsBodySerialization(t *testing.T) {
	t.Run("full body with extra and query fields excluded from body", func(t *testing.T) {
		p := AppLogParams{
			Level:     F(AppLogParamsLevelDebug),
			Message:   F("hello"),
			Service:   F("svc"),
			Directory: F("d"),
			Workspace: F("w"),
			Extra:     F(map[string]any{"k": "v"}),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		// apijson marshals keys alphabetically; query fields excluded
		want := `{"extra":{"k":"v"},"level":"debug","message":"hello","service":"svc"}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
		if strings.Contains(got, "directory") || strings.Contains(got, "workspace") {
			t.Errorf("query fields leaked into body: %s", got)
		}
	})

	t.Run("required-only body (no extra)", func(t *testing.T) {
		p := AppLogParams{
			Level:   F(AppLogParamsLevelError),
			Message: F("m"),
			Service: F("s"),
		}
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatal(err)
		}
		got := string(b)
		want := `{"level":"error","message":"m","service":"s"}`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("all log levels serialize", func(t *testing.T) {
		cases := map[AppLogParamsLevel]string{
			AppLogParamsLevelDebug: "debug",
			AppLogParamsLevelInfo:  "info",
			AppLogParamsLevelError: "error",
			AppLogParamsLevelWarn:  "warn",
		}
		for lvl, want := range cases {
			p := AppLogParams{Level: F(lvl), Message: F("m"), Service: F("s")}
			b, err := json.Marshal(p)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(b), `"level":"`+want+`"`) {
				t.Errorf("level %q not serialized in %s", want, string(b))
			}
		}
	})
}

// Query serialization for AppLogParams (directory, workspace).
func TestAppLogParamsQuery(t *testing.T) {
	p := AppLogParams{
		Level:     F(AppLogParamsLevelInfo),
		Message:   F("m"),
		Service:   F("s"),
		Directory: F("d"),
		Workspace: F("w"),
	}
	got := p.URLQuery().Encode()
	want := "directory=d&workspace=w"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAppLogParamsLevelIsKnown(t *testing.T) {
	for _, l := range []AppLogParamsLevel{AppLogParamsLevelDebug, AppLogParamsLevelInfo, AppLogParamsLevelError, AppLogParamsLevelWarn} {
		if !l.IsKnown() {
			t.Errorf("%q should be known", l)
		}
	}
	if AppLogParamsLevel("trace").IsKnown() {
		t.Error("trace should not be known")
	}
}
