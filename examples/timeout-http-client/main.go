// timeout-http-client demonstrates how WithTimeout interacts with an injected
// *http.Client.
//
// Previously WithTimeout was silently dropped whenever WithHTTPClient was
// supplied, so a caller who injected a client to configure a proxy also had to
// remember to set their own timeout, with nothing warning them.
//
// The rule now: an explicit WithTimeout is applied to your instance, because
// you asked for it. Without one, your instance is left exactly as you
// configured it — the SDK does not impose its 30s default on an object it
// does not own.
//
// RUNS OFFLINE:
//
//	go run ./examples/timeout-http-client
package main

import (
	"fmt"
	"net/http"
	"time"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

func main() {
	// 1. Your client, no WithTimeout -> untouched.
	yours := &http.Client{Timeout: 5 * time.Second}
	if _, err := loginradius.NewClient(
		loginradius.WithAPIKey("demo-api-key"),
		loginradius.WithHTTPClient(yours),
	); err != nil {
		panic(err)
	}
	fmt.Printf("injected client, no WithTimeout:   %v  (your value preserved)\n", yours.Timeout)

	// 2. Your client + an explicit WithTimeout -> applied.
	both := &http.Client{Timeout: 5 * time.Second}
	if _, err := loginradius.NewClient(
		loginradius.WithAPIKey("demo-api-key"),
		loginradius.WithHTTPClient(both),
		loginradius.WithTimeout(90*time.Second),
	); err != nil {
		panic(err)
	}
	fmt.Printf("injected client + WithTimeout(90s): %v  (explicitly applied)\n", both.Timeout)

	// 3. No client at all -> the SDK builds one with its default.
	client, err := loginradius.NewClient(loginradius.WithAPIKey("demo-api-key"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("no injected client:                %v  (SDK default)\n", client.HTTPClient().Timeout)
}
