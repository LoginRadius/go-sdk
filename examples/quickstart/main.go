// quickstart shows the minimum needed to call a LoginRadius endpoint with the
// v12 SDK: construct a client with an API key, then call an operation through
// the typed service. Run with:
//
//	LR_API_KEY=... go run ./examples/quickstart
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

func main() {
	apiKey := os.Getenv("LR_API_KEY")
	if apiKey == "" {
		log.Fatal("LR_API_KEY is required")
	}

	client, err := loginradius.NewClient(
		loginradius.WithAPIKey(apiKey),
	)
	if err != nil {
		log.Fatalf("init client: %v", err)
	}

	// Check whether a username is taken. Simple GET, no body.
	resp, _, err := client.Login.
		CheckUserNameAvailability(context.Background()).
		Username("alice").
		Execute()
	if err != nil {
		log.Fatalf("CheckUserNameAvailability: %v", err)
	}

	fmt.Printf("response: %+v\n", resp)
}
