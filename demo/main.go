// Demo HTTP service for the LoginRadius Go v12 SDK.
//
// The endpoint surface comes from the shared SDK manifest and covers registration and
// login, profile and identifier management, password reset by token or OTP,
// custom objects, the access-token lifecycle, MFA enrolment and challenge, and
// passkey (WebAuthn) — backed by the v12 SDK exactly the way a customer
// integration would use it. A single static HTML page with tabs drives the
// flows in a browser.
//
// DEMO ONLY. Cookies are unsigned. There is no CSRF protection, no rate
// limiting, and the session map lives in process memory. Do not copy this
// code into production. See README.md in this directory for the full caveats.
//
// Run:
//
//	cd demo && go run .
//
// Browse http://localhost:8080/ and walk through the tabs.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

func main() {
	// Populate the process env from demo/.env (or ./.env) before reading it.
	// Anything already exported in the shell takes precedence.
	loadDotEnv()

	apiKey := os.Getenv("LR_API_KEY")
	apiSecret := os.Getenv("LR_API_SECRET")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "LR_API_KEY is required. See .env.example in this directory.")
		os.Exit(2)
	}

	opts := []loginradius.Option{
		loginradius.WithAPIKey(apiKey),
		loginradius.WithUserAgent("loginradius-go-demo/0.1"),
	}
	if apiSecret != "" {
		opts = append(opts, loginradius.WithAPISecret(apiSecret))
	}
	if domain := os.Getenv("LR_DOMAIN"); domain != "" {
		opts = append(opts, loginradius.WithDomain(domain))
	}
	if cd := os.Getenv("LR_CUSTOM_DOMAIN"); cd != "" {
		opts = append(opts, loginradius.WithCustomDomain(cd))
	}
	if base := os.Getenv("LR_BASE_URL"); base != "" {
		opts = append(opts, loginradius.WithBaseURL(base))
	}

	client, err := loginradius.NewClient(opts...)
	if err != nil {
		log.Fatalf("init SDK client: %v", err)
	}

	srv := &server{
		client:    client,
		sessions:  newSessionStore(),
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseOpts:  opts,
	}

	mux := http.NewServeMux()

	// The UI. Not part of the shared contract — each language's demo ships its
	// own look, which is deliberate.
	mux.HandleFunc("/", srv.handleIndex)

	// The API surface comes from demo/routes_gen.go, generated from
	// the shared SDK manifest. Registering from the table (rather than by hand) is
	// what keeps every language's demo on the same endpoints: method checking
	// and session enforcement are applied uniformly here instead of being
	// re-implemented, slightly differently, in each handler.
	//
	// Registered as "METHOD /path" (Go 1.22+ patterns) rather than by path
	// alone: the contract now has paths that carry more than one verb —
	// /api/password/otp is POST to request an OTP and PUT to consume it — and
	// path-only registration would make the second one silently replace the
	// first. Path parameters such as {objectRecordId} come through the same
	// pattern syntax and are read with r.PathValue.
	for _, route := range srv.demoRoutes() {
		mux.Handle(route.Method+" "+route.Path, srv.wrap(route))
	}

	addr := envOr("LR_DEMO_ADDR", ":8080")
	// addr may already carry a host (LR_DEMO_ADDR=127.0.0.1:8099), and prefixing
	// "localhost" then produced http://localhost127.0.0.1:8099/ — the one line a
	// reader copies out of the log.
	browse := addr
	if strings.HasPrefix(browse, ":") {
		browse = "localhost" + browse
	}
	log.Printf("demo listening on %s — browse http://%s/", addr, browse)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// wrap applies the cross-cutting rules the generated route table declares:
// one allowed HTTP method, and session enforcement for authenticated routes.
// Doing it here means each handler is only the interesting part — the SDK
// call — and every language's demo enforces the contract identically.
func (s *server) wrap(route demoRoute) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != route.Method {
			w.Header().Set("Allow", route.Method)
			http.Error(w, errMethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}
		if route.RequiresSession {
			id := getSessionCookie(r)
			token := s.sessions.Lookup(id)
			if token == "" {
				writeJSON(w, http.StatusUnauthorized, map[string]any{"error": "not signed in"})
				return
			}
			// Hand the tokens to the handler without re-reading the cookie.
			// The session id goes too, so a handler that rotates tokens can
			// write the new ones back to the store it came from.
			ctx := withToken(r.Context(), token)
			ctx = withRefreshToken(ctx, s.sessions.LookupRefresh(id))
			ctx = withSessionID(ctx, id)
			r = r.WithContext(ctx)
		}
		if route.RequiresMFAToken {
			// A different credential from a different cookie. This route is
			// reachable only between a login that returned a challenge and the
			// call that completes it, and a session cookie will not open it.
			token := getMFACookie(r)
			if token == "" {
				writeJSON(w, http.StatusUnauthorized, map[string]any{
					"error": "no MFA challenge in progress",
					"hint":  "sign in first; a login that requires a second factor starts the challenge",
				})
				return
			}
			r = r.WithContext(withToken(r.Context(), token))
		}
		route.Handler(w, r)
	})
}

// tokenKey, refreshTokenKey and sessionIDKey type the context keys so they
// cannot collide with another package's values.
type tokenKey struct{}
type refreshTokenKey struct{}
type sessionIDKey struct{}

func withToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey{}, token)
}

// tokenFrom returns the credential wrap() resolved for this route: the
// session's access token for a RequiresSession route, or the second-factor
// token for a RequiresMFAToken one. Generated wrappers pass it straight to the
// matching call<Name>, so a handler never re-reads a cookie and the two token
// kinds can never be confused for one another.
func tokenFrom(ctx context.Context) string {
	token, _ := ctx.Value(tokenKey{}).(string)
	return token
}

func withRefreshToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, refreshTokenKey{}, token)
}

// refreshTokenFrom returns the session's refresh token, or "" when the tenant
// returned none. Only POST /api/token/refresh needs it.
func refreshTokenFrom(ctx context.Context) string {
	token, _ := ctx.Value(refreshTokenKey{}).(string)
	return token
}

func withSessionID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, sessionIDKey{}, id)
}

// sessionIDFrom returns the id of the session backing this request, so a
// handler that rotates the access token can store the replacement.
func sessionIDFrom(ctx context.Context) string {
	id, _ := ctx.Value(sessionIDKey{}).(string)
	return id
}
