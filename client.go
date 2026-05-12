// Package devportalclient wraps the OpenAPI-generated Go client with ergonomic helpers
// (Bearer auth, forecast polling, pagination). The generated low-level client lives in
// the gen/ subdirectory — regenerate with scripts/sdk-generate.sh.
package devportalclient

import (
	"context"
	"net/http"
	"time"

	devportal "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

// Options configures the HTTP client and authentication.
type Options struct {
	// BaseURL is the API origin, e.g. https://api.example.com (no trailing slash required).
	// Empty BaseURL uses OPERATIONAL_API_BASE_URL (if set), else DefaultPublicAPIBaseURL (defaults_gen.go).
	BaseURL string
	// Token is sent as Authorization: Bearer <token> (API key sk_ops_... or Auth0 access token).
	Token string
	// HTTPClient overrides the default net/http client (timeouts, transport, etc.).
	HTTPClient *http.Client
	// UserAgent overrides the default User-Agent header.
	UserAgent string
}

// Client is a thin wrapper around the openapi-generator DefaultAPI client.
type Client struct {
	raw *devportal.APIClient
}

// New constructs a Client. Token is required for authenticated calls.
// Base URL resolution: Options.BaseURL, else OPERATIONAL_API_BASE_URL, else DefaultPublicAPIBaseURL.
func New(opts Options) *Client {
	cfg := devportal.NewConfiguration()
	base := resolveAPIBaseURL(opts.BaseURL)
	cfg.Servers = devportal.ServerConfigurations{{URL: base}}
	if opts.HTTPClient != nil {
		cfg.HTTPClient = opts.HTTPClient
	} else {
		cfg.HTTPClient = &http.Client{Timeout: 60 * time.Second}
	}
	if opts.Token != "" {
		cfg.AddDefaultHeader("Authorization", "Bearer "+opts.Token)
	}
	if opts.UserAgent != "" {
		cfg.UserAgent = opts.UserAgent
	}
	return &Client{raw: devportal.NewAPIClient(cfg)}
}

// Raw exposes the underlying openapi-generated API client.
func (c *Client) Raw() *devportal.APIClient {
	return c.raw
}

// DefaultAPI returns the DefaultAPIService (all REST operations).
func (c *Client) DefaultAPI() *devportal.DefaultAPIService {
	return c.raw.DefaultAPI
}

// Forecasts groups forecast-related helpers.
type Forecasts struct {
	client *Client
}

// Forecasts returns forecast helpers including Wait.
func (c *Client) Forecasts() *Forecasts {
	return &Forecasts{client: c}
}

// Wait polls GET /api/v1/forecasts/{id} until settled or context cancellation.
func (f *Forecasts) Wait(ctx context.Context, jobID string, poll time.Duration) (*devportal.ForecastJobResponse, error) {
	return WaitForecast(ctx, f.client.DefaultAPI(), jobID, poll)
}
