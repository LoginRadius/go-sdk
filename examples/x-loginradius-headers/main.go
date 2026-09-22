// Auth scheme: XLoginRadiusAPIKey + XLoginRadiusAPISecret (header-only).
//
// Use this when the header credentials must differ from the query-param
// credentials — typically when traffic flows through an internal gateway that
// rewrites one but not the other. WithXLoginRadiusAPIKey /
// WithXLoginRadiusAPISecret override ONLY the header values; the query params
// still come from WithAPIKey / WithAPISecret.
//
// If both header and query credentials are identical, just use WithAPIKey /
// WithAPISecret — the SDK populates both transparently.
//
// Required env: LR_X_API_KEY, LR_X_API_SECRET, LR_SOCIAL_TOKEN.
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
	xAPIKey := os.Getenv("LR_X_API_KEY")
	xAPISecret := os.Getenv("LR_X_API_SECRET")
	socialToken := os.Getenv("LR_SOCIAL_TOKEN")
	if xAPIKey == "" || xAPISecret == "" || socialToken == "" {
		fmt.Fprintln(os.Stderr, "Set LR_X_API_KEY, LR_X_API_SECRET, and LR_SOCIAL_TOKEN.")
		os.Exit(2)
	}

	client, err := loginradius.NewClient(
		loginradius.WithXLoginRadiusAPIKey(xAPIKey),       // header: X-LoginRadius-ApiKey
		loginradius.WithXLoginRadiusAPISecret(xAPISecret), // header: X-LoginRadius-ApiSecret
	)
	if err != nil {
		log.Fatalf("init client: %v", err)
	}

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
