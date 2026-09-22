// Auth scheme: BearerToken (Authorization: Bearer <token>).
//
// Endpoints protected with the HTTP bearer scheme accept a bearer token in
// the Authorization header. The SDK sets it on every request when
// WithBearerToken is supplied at construction.
//
// Required env: LR_API_KEY, LR_BEARER_TOKEN.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

func main() {
	apiKey := os.Getenv("LR_API_KEY")
	bearerToken := os.Getenv("LR_BEARER_TOKEN")
	if apiKey == "" || bearerToken == "" {
		fmt.Fprintln(os.Stderr, "Set LR_API_KEY and LR_BEARER_TOKEN.")
		os.Exit(2)
	}

	client, err := loginradius.NewClient(
		loginradius.WithAPIKey(apiKey),
		loginradius.WithBearerToken(bearerToken),
	)
	if err != nil {
		log.Fatalf("init client: %v", err)
	}

	// Retrieve the signed-in user's profile via a bearer-token-protected endpoint.
	resp, _, err := client.User.
		GetAccountDetails(context.Background()).
		Execute()
	if err != nil {
		var lrErr *loginradius.Error
		if errors.As(err, &lrErr) {
			fmt.Fprintf(os.Stderr, "LoginRadius error: %d %s %s\n", lrErr.StatusCode, lrErr.Code, lrErr.Description)
			os.Exit(1)
		}
		log.Fatalf("User.GetAccountDetails: %v", err)
	}
	fmt.Printf("user = %+v\n", resp)
}
