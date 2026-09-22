// Auth scheme: M2MBearerToken (Authorization: Bearer <JWT>).
//
// Machine-to-machine endpoints require an M2M JWT obtained via
// client.OAuthM2M.GenerateM2MToken(...). The SDK sends it on the Authorization
// header. WithBearerToken and WithM2MBearerToken share the same header; use
// WithM2MBearerToken here so the intent is explicit.
//
// Required env: LR_API_KEY, LR_M2M_BEARER_TOKEN.
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
	m2mBearerToken := os.Getenv("LR_M2M_BEARER_TOKEN")
	if apiKey == "" || m2mBearerToken == "" {
		fmt.Fprintln(os.Stderr, "Set LR_API_KEY and LR_M2M_BEARER_TOKEN.")
		os.Exit(2)
	}

	client, err := loginradius.NewClient(
		loginradius.WithAPIKey(apiKey),
		loginradius.WithM2MBearerToken(m2mBearerToken),
	)
	if err != nil {
		log.Fatalf("init client: %v", err)
	}

	// List all SOTT (Secure One-Time Tokens) — an M2M-protected management endpoint.
	resp, _, err := client.SOTT.
		GetAllSOTT(context.Background()).
		Execute()
	if err != nil {
		var lrErr *loginradius.Error
		if errors.As(err, &lrErr) {
			fmt.Fprintf(os.Stderr, "LoginRadius error: %d %s %s\n", lrErr.StatusCode, lrErr.Code, lrErr.Description)
			os.Exit(1)
		}
		log.Fatalf("SOTT.GetAllSOTT: %v", err)
	}
	fmt.Printf("SOTT entries = %+v\n", resp)
}
