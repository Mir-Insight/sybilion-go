package sybilion

import (
	"os"
	"strings"
)

// EnvSybilionAPIBaseURL is the canonical environment variable that overrides the
// compiled-in default API origin when Options.BaseURL is empty.
const EnvSybilionAPIBaseURL = "SYBILION_API_BASE_URL"

// EnvOperationalAPIBaseURL is the previous (pre-launch) name of the env var; kept
// as a silent fallback for one transition release. New code should set
// SYBILION_API_BASE_URL.
const EnvOperationalAPIBaseURL = "OPERATIONAL_API_BASE_URL"

func resolveAPIBaseURL(explicit string) string {
	s := strings.TrimSuffix(strings.TrimSpace(explicit), "/")
	if s != "" {
		return s
	}
	if v := strings.TrimSpace(os.Getenv(EnvSybilionAPIBaseURL)); v != "" {
		return strings.TrimSuffix(v, "/")
	}
	if v := strings.TrimSpace(os.Getenv(EnvOperationalAPIBaseURL)); v != "" {
		return strings.TrimSuffix(v, "/")
	}
	return DefaultPublicAPIBaseURL
}
