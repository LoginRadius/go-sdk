package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

// The demo has two authenticated route kinds and they must not be
// interchangeable: a session cookie proves a completed login, while the MFA
// cookie proves only that a login got as far as a challenge. Letting the
// second stand in for the first would mean a half-authenticated user could
// reach signed-in routes, which is the whole reason they are separate cookies.
//
// These tests walk the GENERATED route table rather than a hand-written list,
// so a route added to the shared SDK manifest is covered the moment it appears.

// newTestServer builds the minimum a route's auth check needs. wrap() only
// touches the session store, so the SDK client is deliberately absent — every
// assertion here is about requests that must be rejected BEFORE any handler,
// and therefore before anything would dereference it.
func newTestServer(t *testing.T) (*server, string) {
	t.Helper()
	s := &server{sessions: newSessionStore()}
	id, err := s.sessions.Create("test-access-token", "test-refresh-token")
	if err != nil {
		t.Fatalf("create session: %v", err)
	}
	return s, id
}

func request(t *testing.T, s *server, route demoRoute, cookies ...*http.Cookie) int {
	t.Helper()
	// Path parameters are irrelevant to the auth check; a literal placeholder
	// keeps the URL well-formed.
	path := strings.NewReplacer("{objectRecordId}", "some-id").Replace(route.Path)
	req := httptest.NewRequest(route.Method, path, strings.NewReader("{}"))
	req.Header.Set("Content-Type", "application/json")
	for _, c := range cookies {
		req.AddCookie(c)
	}
	rec := httptest.NewRecorder()
	s.wrap(route).ServeHTTP(rec, req)
	return rec.Code
}

func TestSessionRoutesRejectAnonymousCallers(t *testing.T) {
	s, _ := newTestServer(t)
	checked := 0
	for _, route := range s.demoRoutes() {
		if !route.RequiresSession {
			continue
		}
		checked++
		if got := request(t, s, route); got != http.StatusUnauthorized {
			t.Errorf("%s %s without a session: got %d, want %d",
				route.Method, route.Path, got, http.StatusUnauthorized)
		}
	}
	if checked == 0 {
		t.Fatal("no session-authenticated routes found — the route table looks wrong")
	}
}

func TestMFARoutesRejectAnonymousCallers(t *testing.T) {
	s, _ := newTestServer(t)
	checked := 0
	for _, route := range s.demoRoutes() {
		if !route.RequiresMFAToken {
			continue
		}
		checked++
		if got := request(t, s, route); got != http.StatusUnauthorized {
			t.Errorf("%s %s without an MFA challenge: got %d, want %d",
				route.Method, route.Path, got, http.StatusUnauthorized)
		}
	}
	if checked == 0 {
		t.Fatal("no MFA-token routes found — the route table looks wrong")
	}
}

// The one that matters: a fully signed-in user must still not be able to
// complete somebody's MFA challenge just by holding a session.
func TestSessionCookieDoesNotSatisfyMFARoutes(t *testing.T) {
	s, id := newTestServer(t)
	session := &http.Cookie{Name: sessionCookieName, Value: id}
	checked := 0
	for _, route := range s.demoRoutes() {
		if !route.RequiresMFAToken {
			continue
		}
		checked++
		if got := request(t, s, route, session); got != http.StatusUnauthorized {
			t.Errorf("%s %s accepted a session cookie: got %d, want %d",
				route.Method, route.Path, got, http.StatusUnauthorized)
		}
	}
	if checked == 0 {
		t.Fatal("no MFA-token routes found — the route table looks wrong")
	}
}

// ...and the reverse: an in-progress challenge must not open signed-in routes.
func TestMFACookieDoesNotSatisfySessionRoutes(t *testing.T) {
	s, _ := newTestServer(t)
	mfa := &http.Cookie{Name: mfaCookieName, Value: "second-factor-token"}
	for _, route := range s.demoRoutes() {
		if !route.RequiresSession {
			continue
		}
		if got := request(t, s, route, mfa); got != http.StatusUnauthorized {
			t.Errorf("%s %s accepted an MFA cookie: got %d, want %d",
				route.Method, route.Path, got, http.StatusUnauthorized)
		}
	}
}

func TestAuthModesAreMutuallyExclusive(t *testing.T) {
	s, _ := newTestServer(t)
	for _, route := range s.demoRoutes() {
		if route.RequiresSession && route.RequiresMFAToken {
			t.Errorf("%s %s requires both a session and an MFA token; "+
				"the two credentials mean different things and no route should need both",
				route.Method, route.Path)
		}
	}
}

func TestCookieNamesAreDistinct(t *testing.T) {
	if sessionCookieName == mfaCookieName {
		t.Fatalf("session and MFA cookies share the name %q; a half-authenticated "+
			"user would hold a credential the session middleware accepts", sessionCookieName)
	}
}

func TestWrapRejectsTheWrongMethod(t *testing.T) {
	s, _ := newTestServer(t)
	for _, route := range s.demoRoutes() {
		if route.Method == http.MethodPost {
			continue
		}
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		rec := httptest.NewRecorder()
		s.wrap(route).ServeHTTP(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("%s %s accepted POST: got %d, want %d",
				route.Method, route.Path, rec.Code, http.StatusMethodNotAllowed)
		}
		if allow := rec.Header().Get("Allow"); allow != route.Method {
			t.Errorf("%s %s: Allow header = %q, want %q", route.Method, route.Path, allow, route.Method)
		}
	}
}

// Refreshing an access token INVALIDATES the previous one upstream. A session
// left holding the old token is therefore not merely stale — every subsequent
// authenticated call fails, which is exactly what a caller sees as "validate
// and active-session broke after I refreshed". These tests pin the rotation to
// the session store rather than to a live tenant.

// stubAPI returns a demo server whose SDK client talks to a local stub instead
// of LoginRadius, so the assertions are about our own state handling.
func stubAPI(t *testing.T, body string) (*server, func()) {
	t.Helper()
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, body)
	}))
	client, err := loginradius.NewClient(
		loginradius.WithAPIKey("test-key"),
		loginradius.WithBaseURL(upstream.URL),
	)
	if err != nil {
		upstream.Close()
		t.Fatalf("build client: %v", err)
	}
	return &server{client: client, sessions: newSessionStore()}, upstream.Close
}

func TestRefreshRotatesTheSessionTokens(t *testing.T) {
	s, done := stubAPI(t, `{"access_token":"new-access","refresh_token":"new-refresh","expires_in":"3600"}`)
	defer done()

	id, err := s.sessions.Create("old-access", "old-refresh")
	if err != nil {
		t.Fatalf("create session: %v", err)
	}
	ctx := withSessionID(withRefreshToken(context.Background(), "old-refresh"), id)

	if _, err := s.callRefreshToken(ctx, refreshTokenRequest{}, "old-access"); err != nil {
		t.Fatalf("callRefreshToken: %v", err)
	}

	if got := s.sessions.Lookup(id); got != "new-access" {
		t.Errorf("session access token = %q, want %q — the refreshed token was not stored, "+
			"so every later authenticated call would use an invalidated token", got, "new-access")
	}
	if got := s.sessions.LookupRefresh(id); got != "new-refresh" {
		t.Errorf("session refresh token = %q, want %q", got, "new-refresh")
	}
}

// A tenant that rotates only the access token must not blank the refresh token.
func TestRefreshKeepsTheOldRefreshTokenWhenNoneIsReturned(t *testing.T) {
	s, done := stubAPI(t, `{"access_token":"new-access","expires_in":"3600"}`)
	defer done()

	id, _ := s.sessions.Create("old-access", "old-refresh")
	ctx := withSessionID(withRefreshToken(context.Background(), "old-refresh"), id)

	if _, err := s.callRefreshToken(ctx, refreshTokenRequest{}, "old-access"); err != nil {
		t.Fatalf("callRefreshToken: %v", err)
	}
	if got := s.sessions.LookupRefresh(id); got != "old-refresh" {
		t.Errorf("refresh token = %q, want the original %q kept", got, "old-refresh")
	}
}

// A success-shaped response carrying no token must not wipe a live session.
func TestRefreshWithNoTokenLeavesTheSessionIntact(t *testing.T) {
	s, done := stubAPI(t, `{"expires_in":"3600"}`)
	defer done()

	id, _ := s.sessions.Create("old-access", "old-refresh")
	ctx := withSessionID(withRefreshToken(context.Background(), "old-refresh"), id)

	if _, err := s.callRefreshToken(ctx, refreshTokenRequest{}, "old-access"); err == nil {
		t.Fatal("expected an error when the refresh returns no access_token")
	}
	if got := s.sessions.Lookup(id); got != "old-access" {
		t.Errorf("session access token = %q, want the original %q — a tokenless "+
			"response must not sign the user out", got, "old-access")
	}
}
