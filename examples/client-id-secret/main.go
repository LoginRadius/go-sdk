// Auth scheme: ClientId + ClientSecret (query parameters).
//
// OAuth-style endpoints — typically multipurpose token operations and account
// management calls that act on behalf of a registered application rather than
// a tenant — require the application's client_id and client_secret. The SDK
// sends both as query parameters on every request.
//
// Required env: LR_CLIENT_ID, LR_CLIENT_SECRET, LR_TARGET_UID.
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
	clientID := os.Getenv("LR_CLIENT_ID")
	clientSecret := os.Getenv("LR_CLIENT_SECRET")
	targetUID := os.Getenv("LR_TARGET_UID")
	if clientID == "" || clientSecret == "" || targetUID == "" {
		fmt.Fprintln(os.Stderr, "Set LR_CLIENT_ID, LR_CLIENT_SECRET, and LR_TARGET_UID.")
		os.Exit(2)
	}

	client, err := loginradius.NewClient(
		loginradius.WithClientID(clientID),
		loginradius.WithClientSecret(clientSecret),
	)
	if err != nil {
		log.Fatalf("init client: %v", err)
	}

	// Generate a delete-user token for a target user — uses ClientId + ClientSecret.
	// The request body is the polymorphic MultipurposeEmailTokenAPIRequest wrapper;
	// set exactly one branch (DeleteUserModel here, matching tokentype=deleteuser).
	body := loginradius.MultipurposeEmailTokenAPIRequest{
		DeleteUserModel: &loginradius.DeleteUserModel{Uid: targetUID},
	}
	resp, _, err := client.MultipurposeTokens.
		MultipurposeEmailTokenAPI(context.Background(), "deleteuser").
		MultipurposeEmailTokenAPIRequest(body).
		Execute()
	if err != nil {
		var lrErr *loginradius.Error
		if errors.As(err, &lrErr) {
			fmt.Fprintf(os.Stderr, "LoginRadius error: %d %s %s\n", lrErr.StatusCode, lrErr.Code, lrErr.Description)
			os.Exit(1)
		}
		log.Fatalf("MultipurposeTokens.MultipurposeEmailTokenAPI: %v", err)
	}
	fmt.Printf("verification token = %+v\n", resp)
}
