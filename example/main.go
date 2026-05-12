// Command example is a minimal consumer of the Go SDK (same pattern as an external app
// using a go.mod replace to this repo). It calls GET /api/v1/me.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	devportalclient "github.com/Mir-Insight/developers-portal-api-sdk-go"
)

func main() {
	token := os.Getenv("OPERATIONAL_API_TOKEN")
	if token == "" {
		log.Fatal("OPERATIONAL_API_TOKEN is required (see .env.example)")
	}

	c := devportalclient.New(devportalclient.Options{Token: token})
	me, _, err := c.DefaultAPI().ApiV1MeGet(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok user_id=%s balance_eur_cents=%d available_eur_cents=%d\n",
		me.GetUserId(), me.GetBalanceEurCents(), me.GetAvailableEurCents())
}
