package devportalclient

import (
	"os"
	"strings"
)

// EnvOperationalAPIBaseURL is the environment variable that overrides the compiled-in
// default API origin when Options.BaseURL is empty.
const EnvOperationalAPIBaseURL = "OPERATIONAL_API_BASE_URL"

func resolveAPIBaseURL(explicit string) string {
	s := strings.TrimSuffix(strings.TrimSpace(explicit), "/")
	if s != "" {
		return s
	}
	if v := strings.TrimSpace(os.Getenv(EnvOperationalAPIBaseURL)); v != "" {
		return strings.TrimSuffix(v, "/")
	}
	return DefaultPublicAPIBaseURL
}
