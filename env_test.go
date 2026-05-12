package sybilion

import "testing"

func TestResolveAPIBaseURL(t *testing.T) {
	t.Run("explicit wins over env", func(t *testing.T) {
		t.Setenv(EnvSybilionAPIBaseURL, "https://env.example")
		got := resolveAPIBaseURL("https://explicit.example/")
		want := "https://explicit.example"
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})
	t.Run("env when explicit empty", func(t *testing.T) {
		t.Setenv(EnvSybilionAPIBaseURL, "https://from-env.example/")
		got := resolveAPIBaseURL("")
		want := "https://from-env.example"
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})
	t.Run("legacy env still honored", func(t *testing.T) {
		t.Setenv(EnvSybilionAPIBaseURL, "")
		t.Setenv(EnvOperationalAPIBaseURL, "https://legacy.example/")
		got := resolveAPIBaseURL("")
		want := "https://legacy.example"
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})
	t.Run("default when env unset", func(t *testing.T) {
		t.Setenv(EnvSybilionAPIBaseURL, "")
		t.Setenv(EnvOperationalAPIBaseURL, "")
		got := resolveAPIBaseURL("")
		if got != DefaultPublicAPIBaseURL {
			t.Fatalf("got %q, want %q", got, DefaultPublicAPIBaseURL)
		}
	})
}
