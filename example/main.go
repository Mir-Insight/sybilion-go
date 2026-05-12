// Command example is a minimal consumer of the Sybilion Go SDK. It calls GET /api/v1/me.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.sybilion.dev/sybilion"
)

func main() {
	token := os.Getenv("SYBILION_API_TOKEN")
	if token == "" {
		log.Fatal("SYBILION_API_TOKEN is required (see .env.example)")
	}

	c := sybilion.New(sybilion.Options{Token: token})
	me, _, err := c.DefaultAPI().ApiV1MeGet(context.Background()).Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok user_id=%s balance_eur_cents=%d available_eur_cents=%d\n",
		me.GetUserId(), me.GetBalanceEurCents(), me.GetAvailableEurCents())
}
