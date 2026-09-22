// Auth scheme: APIKey (X-LoginRadius-ApiKey header + apikey query parameter).
//
// The simplest LoginRadius auth setup — most "public" operations (availability
// checks, login flows, registration) accept just the tenant API key. The SDK
// sends it as a header by preference, keeping it out of access logs and URL
// caches, and additionally as a query parameter for the operations that accept
// nothing else.
//
// Nothing here is server-side-only: an API key on its own identifies your app,
// not a user, and cannot act on a user's behalf. That is what separates it from
// api-key-secret/.
//
// Required env: LR_API_KEY.
//
//	LR_API_KEY=... go run ./examples/api-key
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
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Set LR_API_KEY.")
		os.Exit(2)
	}

	client, err := loginradius.NewClient(loginradius.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("init client: %v", err)
	}

	resp, _, err := client.Login.
		CheckUserNameAvailability(context.Background()).
		Username("alice").
		Execute()
	if err != nil {
		// The typed error carries the HTTP status, the LoginRadius error code,
		// and the raw body — branch on intent rather than on status codes.
		var lrErr *loginradius.Error
		if errors.As(err, &lrErr) {
			if lrErr.IsAuth() {
				log.Fatalf("the API key was rejected: %s", lrErr.Description)
			}
			log.Fatalf("failed: %s (%s)", lrErr.Description, lrErr.Code)
		}
		log.Fatalf("failed: %v", err)
	}

	if resp.IsExist != nil {
		fmt.Println("IsExist =", *resp.IsExist)
	}
}
