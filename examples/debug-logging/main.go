// debug-logging demonstrates WithDebug — a one-line summary of every request.
//
// The property worth verifying here is REDACTION: credential header values are
// replaced with [REDACTED] before anything reaches your writer, and query
// parameters are logged by key only. Writing a tenant secret into application
// logs would be a disclosure, so the writer never sees one.
//
// RUNS OFFLINE — a real SDK operation is called and the request captured
// rather than sent:
//
//	go run ./examples/debug-logging
//
// Exits non-zero if any credential value appears in the log.
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

type recorder struct{}

func (recorder) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200, Body: http.NoBody, Header: make(http.Header), Request: req}, nil
}

func main() {
	var buf strings.Builder

	client, err := loginradius.NewClient(
		loginradius.WithAPIKey("SUPER-SECRET-KEY"),
		loginradius.WithAPISecret("SUPER-SECRET-SECRET"),
		loginradius.WithBearerToken("SUPER-SECRET-TOKEN"),
		loginradius.WithHTTPClient(&http.Client{Transport: recorder{}}),
		// Any io.Writer: os.Stderr, a file, a log pipe. Buffered here so the
		// example can assert on the output.
		loginradius.WithDebug(&buf),
	)
	if err != nil {
		panic(err)
	}

	_, _, _ = client.Login.CheckUserNameAvailability(context.Background()).Username("alice").Execute()

	out := buf.String()
	fmt.Println("--- what the debug writer received ---")
	fmt.Print(out)
	fmt.Println("--------------------------------------")

	leaked := []string{}
	for _, secret := range []string{"SUPER-SECRET-KEY", "SUPER-SECRET-SECRET", "SUPER-SECRET-TOKEN"} {
		if strings.Contains(out, secret) {
			leaked = append(leaked, secret)
		}
	}
	if len(leaked) > 0 {
		fmt.Println("FAIL: credential values leaked:", leaked)
		os.Exit(1)
	}
	fmt.Println("OK: no credential value appears in the log — only header names.")
}
