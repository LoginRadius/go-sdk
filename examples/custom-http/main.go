// custom-http demonstrates how to plug a custom *http.Client into the SDK.
// Useful for: adding a proxy, custom TLS config, request logging, metrics, or
// a shared transport pool across SDKs.
//
// The SDK wraps the supplied client's Transport with its own RoundTripper to
// inject auth and the User-Agent. The original Transport remains underneath,
// so any dialer/proxy/TLS config you set up is preserved.
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

func main() {
	apiKey := os.Getenv("LR_API_KEY")
	if apiKey == "" {
		log.Fatal("LR_API_KEY is required")
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        50,
			MaxIdleConnsPerHost: 25,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	client, err := loginradius.NewClient(
		loginradius.WithAPIKey(apiKey),
		loginradius.WithHTTPClient(httpClient),
		loginradius.WithUserAgent("acme-corp/1.0 loginradius-go"),
	)
	if err != nil {
		log.Fatalf("init client: %v", err)
	}

	if _, _, err := client.Login.
		CheckUserNameAvailability(context.Background()).
		Username("alice").
		Execute(); err != nil {
		log.Fatalf("CheckUserNameAvailability: %v", err)
	}
}
