// operation-servers demonstrates the fix for a defect that silently ignored
// your server configuration.
//
// The OpenAPI spec pins 42 operations to their own hosts (cloud-api,
// migration.loginradius.com, and tenant-hub / custom-domain templates). The
// generated layer resolves those independently of the client's base URL, so
// before this fix an explicitly configured WithBaseURL was ignored for all 42,
// and tenant-templated hosts always resolved to the spec's placeholder tenant
// ("example").
//
// RUNS OFFLINE:
//
//	go run ./examples/operation-servers
//
// What to look for:
//   - default            -> example.hub.loginradius.com  (spec placeholder)
//   - WithDomain(acme)   -> acme.hub.loginradius.com     (variable filled)
//   - WithBaseURL(...)   -> your host                    (explicit wins)
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

func hostFor(label string, opts ...loginradius.Option) {
	rec := &recorder{}
	all := append([]loginradius.Option{
		loginradius.WithAPIKey("demo-api-key"),
		loginradius.WithHTTPClient(&http.Client{Transport: rec}),
	}, opts...)
	client, err := loginradius.NewClient(all...)
	if err != nil {
		panic(err)
	}

	// GetBigCommerceLoginUrl is one of the 42: the spec pins it to
	// https://{domain}.hub.loginradius.com rather than the tenant API host.
	_, _, _ = client.BigCommerceSSO.GetBigCommerceLoginUrl(context.Background()).
		AccessToken("demo-token").Store("demo-store").Execute()

	if rec.got == nil {
		fmt.Printf("  %-28s (no request)\n", label)
		return
	}
	fmt.Printf("  %-28s %s\n", label, rec.got.URL.Host)
}

func main() {
	fmt.Println("Host used by a pinned operation (BigCommerceSSO.GetBigCommerceLoginUrl):")
	fmt.Println()
	hostFor("no server options")
	hostFor("WithDomain(\"acme\")", loginradius.WithDomain("acme"))
	hostFor("WithBaseURL(staging)", loginradius.WithBaseURL("https://staging.internal"))
	fmt.Println()
	fmt.Println("Before the fix every line above showed example.hub.loginradius.com.")
}
