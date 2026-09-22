// request-signing demonstrates WithAPIRequestSigning — the `digest` and
// `x-Request-Expires` headers the LoginRadius API accepts on management
// endpoints.
//
// RUNS OFFLINE. Real SDK operations are called; a recording transport captures
// the request the SDK produced instead of sending it, so you can inspect the
// headers without a tenant or a network call:
//
//	go run ./examples/request-signing
//
// What to look for:
//   - EmailTemplates.GetEmailTemplates (/v2/manage/...) carries `digest`
//   - Login.CheckUserNameAvailability (/identity/v2/auth/...) does NOT —
//     only management paths are signed
//   - Accounts.GetImpersonationToken is excluded even though its path is under
//     /manage/, because it is the access-token exchange
//   - `apisecret` is stripped from the signed request: the secret must never be
//     part of the string that is signed, nor travel on a signed request
package main

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

// recorder captures the request the SDK built and returns a canned response.
type recorder struct{ got *http.Request }

func (r *recorder) RoundTrip(req *http.Request) (*http.Response, error) {
	r.got = req
	return &http.Response{StatusCode: 200, Body: http.NoBody, Header: make(http.Header), Request: req}, nil
}

// call runs one real SDK operation and reports what went on the wire.
func call(label string, signing bool, op func(context.Context, *loginradius.Client) error) {
	rec := &recorder{}
	client, err := loginradius.NewClient(
		loginradius.WithAPIKey("demo-api-key"),
		loginradius.WithAPISecret("demo-api-secret"),
		loginradius.WithHTTPClient(&http.Client{Transport: rec}),
		loginradius.WithAPIRequestSigning(signing),
	)
	if err != nil {
		fmt.Println("  client error:", err)
		return
	}
	// The API call itself fails (the recorder returns an empty 200), which is
	// fine — we only care about the request the SDK built.
	_ = op(context.Background(), client)
	if rec.got == nil {
		fmt.Printf("%s\n  (no request reached the transport)\n\n", label)
		return
	}

	fmt.Printf("%s\n  %s\n", label, rec.got.URL.Path)
	if digest := rec.got.Header.Get("digest"); digest == "" {
		fmt.Println("  digest:            (not signed)")
	} else {
		fmt.Println("  digest:           ", digest)
		fmt.Println("  x-Request-Expires:", rec.got.Header.Get("x-Request-Expires"))
	}
	keys := make([]string, 0)
	for k := range rec.got.URL.Query() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Printf("  query keys:        %v\n\n", keys)
}

// The three operations below are chosen for the path each one sits on.
func getEmailTemplates(ctx context.Context, c *loginradius.Client) error {
	_, _, err := c.EmailTemplates.GetEmailTemplates(ctx).Execute()
	return err
}

func checkUsername(ctx context.Context, c *loginradius.Client) error {
	_, _, err := c.Login.CheckUserNameAvailability(ctx).Username("alice").Execute()
	return err
}

func impersonationToken(ctx context.Context, c *loginradius.Client) error {
	_, _, err := c.Accounts.GetImpersonationToken(ctx).Uid("demo-uid").Execute()
	return err
}

func main() {
	fmt.Println("=== signing DISABLED (the default) ===")
	call("EmailTemplates.GetEmailTemplates", false, getEmailTemplates)

	fmt.Println("=== signing ENABLED ===")
	call("EmailTemplates.GetEmailTemplates   -> signed", true, getEmailTemplates)
	call("Login.CheckUserNameAvailability    -> not signed (auth path)", true, checkUsername)
	call("Accounts.GetImpersonationToken     -> excluded (access-token exchange)", true, impersonationToken)

	fmt.Println("Note: `apisecret` is absent from the query keys of the signed request.")
	fmt.Println("The secret is stripped before signing and never sent on a signed call.")
}
