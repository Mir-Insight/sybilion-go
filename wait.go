package devportalclient

import (
	"context"
	"time"

	devportal "github.com/Mir-Insight/developers-portal-api-sdk-go/gen"
)

// WaitForecast polls GET /api/v1/forecasts/{id} until the job is settled or the context is done.
func WaitForecast(ctx context.Context, api *devportal.DefaultAPIService, jobID string, poll time.Duration) (*devportal.ForecastJobResponse, error) {
	if poll <= 0 {
		poll = 2 * time.Second
	}
	t := time.NewTicker(poll)
	defer t.Stop()
	for {
		j, _, err := api.ApiV1ForecastsIdGet(ctx, jobID).Execute()
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
