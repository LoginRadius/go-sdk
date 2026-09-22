// Auth scheme: AccessToken (query access_token).
//
// User-context endpoints operate on the signed-in user's own profile/session
// and require a LoginRadius access_token obtained from a prior login flow.
// The SDK sends it as the access_token= query parameter.
//
// Required env: LR_API_KEY, LR_ACCESS_TOKEN, LR_OIDC_APP_NAME.
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
	accessToken := os.Getenv("LR_ACCESS_TOKEN")
	oidcAppName := os.Getenv("LR_OIDC_APP_NAME")
	if apiKey == "" || accessToken == "" || oidcAppName == "" {
		fmt.Fprintln(os.Stderr, "Set LR_API_KEY, LR_ACCESS_TOKEN, and LR_OIDC_APP_NAME.")
		os.Exit(2)
	}

	client, err := loginradius.NewClient(
		loginradius.WithAPIKey(apiKey),
		loginradius.WithAccessToken(accessToken),
	)
	if err != nil {
		log.Fatalf("init client: %v", err)
	}

	// OIDC userinfo endpoint — requires access_token in the query string.
	resp, _, err := client.OIDC.
		GetOIDCUserinfo(context.Background(), oidcAppName).
		Execute()
	if err != nil {
		var lrErr *loginradius.Error
		if errors.As(err, &lrErr) {
			fmt.Fprintf(os.Stderr, "LoginRadius error: %d %s %s\n", lrErr.StatusCode, lrErr.Code, lrErr.Description)
			os.Exit(1)
		}
		log.Fatalf("OIDC.GetOIDCUserinfo: %v", err)
	}
	fmt.Printf("userinfo = %+v\n", resp)
}
