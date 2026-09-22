// default-headers demonstrates WithDefaultHeaders — headers merged into every
// outgoing request.
//
// The important property this example proves: default headers are applied at
// the LOWEST precedence. The SDK's own credential, User-Agent, and signing
// headers always win, so a default header can never mask a credential and
// silently send the wrong one.
//
// RUNS OFFLINE — a real SDK operation is called and the request captured
// rather than sent:
//
//	go run ./examples/default-headers
package main

import (
	"context"
	"fmt"
	"net/http"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

type recorder struct{ got *http.Request }

func (r *recorder) RoundTrip(req *http.Request) (*http.Response, error) {
	r.got = req
	return &http.Response{StatusCode: 200, Body: http.NoBody, Header: make(http.Header), Request: req}, nil
}

func main() {
	rec := &recorder{}
	client, err := loginradius.NewClient(
		loginradius.WithAPIKey("REAL-API-KEY"),
		loginradius.WithHTTPClient(&http.Client{Transport: rec}),
		loginradius.WithDefaultHeaders(map[string]string{
			// A genuinely useful default: a correlation ID for your logs.
			"X-Tenant-Trace": "trace-abc123",
			// These two ATTEMPT to override values the SDK owns. Both are
			// expected to lose.
			"X-LoginRadius-ApiKey": "HIJACKED",
			"User-Agent":           "HIJACKED",
		}),
	)
	if err != nil {
		panic(err)
	}

	_, _, _ = client.Login.CheckUserNameAvailability(context.Background()).Username("alice").Execute()

	fmt.Println("Login.CheckUserNameAvailability ->", rec.got.URL.Path)
	fmt.Println()
	fmt.Println("X-Tenant-Trace:      ", rec.got.Header.Get("X-Tenant-Trace"), "  <- merged")
	fmt.Println("X-LoginRadius-ApiKey:", rec.got.Header.Get("X-LoginRadius-ApiKey"), "<- SDK wins, not HIJACKED")
	fmt.Println("User-Agent:          ", rec.got.Header.Get("User-Agent"), "<- SDK wins, not HIJACKED")
}
