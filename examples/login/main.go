// login demonstrates a passwordless login flow: the user receives a one-time
// email link, which on click delivers an access token. This example covers
// the first half (initiating the email); completing the flow happens in the
// browser via the verification endpoint.
//
//	LR_API_KEY=...  go run ./examples/login -- user@example.com
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("usage: login <email>")
	}
	email := flag.Arg(0)

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

	resp, _, err := client.Login.
		PasswordlessLoginByEmail(context.Background()).
		Email(email).
		Execute()
	if err != nil {
		// Typed error gives access to HTTP status, LoginRadius error code, and raw body.
		var lrErr *loginradius.Error
		if errors.As(err, &lrErr) && lrErr.IsAuth() {
			log.Fatalf("authentication rejected: %s", lrErr)
		}
		log.Fatalf("PasswordlessLoginByEmail: %v", err)
	}
	fmt.Printf("posted: %v\n", resp.GetIsPosted())
}
