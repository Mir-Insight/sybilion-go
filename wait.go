package sybilion

import (
	"context"
	"time"

	api "go.sybilion.dev/sybilion/api"
)

// WaitForecast polls GET /api/v1/forecasts/{id} until the job is settled or the context is done.
func WaitForecast(ctx context.Context, defaultAPI *api.DefaultAPIService, jobID string, poll time.Duration) (*api.ApiV1ForecastsIdGet200Response, error) {
	if poll <= 0 {
		poll = 2 * time.Second
	}
	t := time.NewTicker(poll)
	defer t.Stop()
	for {
		j, _, err := defaultAPI.ApiV1ForecastsIdGet(ctx, jobID).Execute()
		if err != nil {
			return nil, err
		}
		if j.GetSettled() {
			return j, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-t.C:
		}
	}
}
