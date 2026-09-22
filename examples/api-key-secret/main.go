// Auth scheme: APIKey + APISecret.
//
// Server-side-only operations (token exchange, account lookup, management)
// additionally require the tenant API secret. The SDK sends the secret as the
// X-LoginRadius-ApiSecret header (preferred) with apisecret= query fallback
// for legacy endpoints. NEVER expose the API secret in a browser or mobile
// context.
//
// Required env: LR_API_KEY, LR_API_SECRET, LR_SOCIAL_TOKEN.
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
	apiSecret := os.Getenv("LR_API_SECRET")
	socialToken := os.Getenv("LR_SOCIAL_TOKEN")
	if apiKey == "" || apiSecret == "" || socialToken == "" {
		fmt.Fprintln(os.Stderr, "Set LR_API_KEY, LR_API_SECRET, and LR_SOCIAL_TOKEN.")
		os.Exit(2)
	}

	client, err := loginradius.NewClient(
		loginradius.WithAPIKey(apiKey),
		loginradius.WithAPISecret(apiSecret),
	)
	if err != nil {
		log.Fatalf("init client: %v", err)
	}

	// Exchange a social-provider token for a LoginRadius access_token. This
	// operation requires both apiKey AND apiSecret per the OpenAPI spec.
	resp, _, err := client.AccountSession.
		GetAccessToken(context.Background()).
		Token(socialToken).
		Execute()
	if err != nil {
		var lrErr *loginradius.Error
		if errors.As(err, &lrErr) {
			fmt.Fprintf(os.Stderr, "LoginRadius error: %d %s %s\n", lrErr.StatusCode, lrErr.Code, lrErr.Description)
			os.Exit(1)
		}
		log.Fatalf("AccountSession.GetAccessToken: %v", err)
	}
	fmt.Printf("access_token = %v\n", resp.GetAccessToken())
}
