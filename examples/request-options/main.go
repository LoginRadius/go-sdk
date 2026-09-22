// request-options demonstrates the four client-wide request options that the
// legacy v11 SDK had and v12 was missing: WithOriginIP, WithServerRegion,
// WithFields, and WithPreventWebhook.
//
// Each is applied to EVERY outgoing request by the same interceptor that
// injects credentials, so there is one place to audit rather than 210
// hand-written call sites.
//
// RUNS OFFLINE — a real SDK operation is called and the request is captured
// rather than sent:
//
//	go run ./examples/request-options
package main

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

type recorder struct{ got *http.Request }

func (r *recorder) RoundTrip(req *http.Request) (*http.Response, error) {
	r.got = req
	return &http.Response{StatusCode: 200, Body: http.NoBody, Header: make(http.Header), Request: req}, nil
}

func capture(opts ...loginradius.Option) *http.Request {
	rec := &recorder{}
	all := append([]loginradius.Option{
		loginradius.WithAPIKey("demo-api-key"),
		loginradius.WithHTTPClient(&http.Client{Transport: rec}),
	}, opts...)
	client, err := loginradius.NewClient(all...)
	if err != nil {
		panic(err)
	}
	// Any operation would do — the options below apply to every request.
	_, _, _ = client.Login.CheckUserNameAvailability(context.Background()).Username("alice").Execute()
	return rec.got
}

func report(label string, r *http.Request) {
	fmt.Println(label)
	for _, h := range []string{"X-Origin-IP", "X-PreventWebhook"} {
		v := r.Header.Get(h)
		if v == "" {
			v = "(unset)"
		}
		fmt.Printf("  %-18s %s\n", h+":", v)
	}
	q := r.URL.Query()
	for _, p := range []string{"region", "fields"} {
		v := q.Get(p)
		if v == "" {
			v = "(unset)"
		}
		fmt.Printf("  %-18s %s\n", p+"=", v)
	}
	keys := make([]string, 0, len(q))
	for k := range q {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Printf("  all query keys:    %v\n\n", keys)
}

func main() {
	fmt.Println("=== nothing set: the SDK adds only what it must ===")
	report("Login.CheckUserNameAvailability", capture())

	fmt.Println("=== all four options set ===")
	report("Login.CheckUserNameAvailability", capture(
		// Forwards the end user's IP for risk-based auth and audit trails.
		loginradius.WithOriginIP("203.0.113.7"),
		// Routes to a regional host. Distinct from WithDomain, which selects
		// a tenant rather than a region.
		loginradius.WithServerRegion("eu"),
		// Trims every response to these fields.
		loginradius.WithFields("Email,Uid"),
		// Suppresses webhooks for every call from this client.
		loginradius.WithPreventWebhook(true),
	))
}
