// Package sybilion is the official Go client for the Sybilion API.
//
// The hand-written wrapper provides ergonomic helpers (Bearer auth, forecast
// polling, pagination iterators) on top of the OpenAPI-generated client at
// go.sybilion.dev/sybilion/api. Most users only need:
//
//	c := sybilion.New(sybilion.Options{Token: os.Getenv("SYBILION_API_TOKEN")})
//	me, _, err := c.DefaultAPI().ApiV1MeGet(ctx).Execute()
//
// See https://sybilion.dev/docs/ for full guides and feature walkthroughs.
package sybilion

import (
	"context"
	"net/http"
	"time"

	api "go.sybilion.dev/sybilion/api"
)

// Options configures the HTTP client and authentication.
type Options struct {
	// BaseURL is the API origin, e.g. https://api.example.com (no trailing slash required).
	// Empty BaseURL uses SYBILION_API_BASE_URL (if set), else DefaultPublicAPIBaseURL (defaults_gen.go).
	BaseURL string
	// Token is sent as Authorization: Bearer <token> (an API key sk_ops_... or a dashboard session token).
	Token string
	// HTTPClient overrides the default net/http client (timeouts, transport, etc.).
	HTTPClient *http.Client
	// UserAgent overrides the default User-Agent header.
	UserAgent string
}

// Client is a thin wrapper around the openapi-generator DefaultAPI client.
type Client struct {
	raw *api.APIClient
}

// New constructs a Client. Token is required for authenticated calls.
// Base URL resolution: Options.BaseURL, else SYBILION_API_BASE_URL, else DefaultPublicAPIBaseURL.
func New(opts Options) *Client {
	cfg := api.NewConfiguration()
	base := resolveAPIBaseURL(opts.BaseURL)
	cfg.Servers = api.ServerConfigurations{{URL: base}}
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
	return &Client{raw: api.NewAPIClient(cfg)}
}

// Raw exposes the underlying openapi-generated API client.
func (c *Client) Raw() *api.APIClient {
	return c.raw
}

// DefaultAPI returns the DefaultAPIService (all REST operations).
func (c *Client) DefaultAPI() *api.DefaultAPIService {
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
func (f *Forecasts) Wait(ctx context.Context, jobID string, poll time.Duration) (*api.ApiV1ForecastsIdGet200Response, error) {
	return WaitForecast(ctx, f.client.DefaultAPI(), jobID, poll)
}
