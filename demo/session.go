package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
	"time"
)

// sessionStore maps an opaque session id (from the lr_session cookie) to the
// LoginRadius access_token returned by /api/login. In-memory and unsigned —
// adequate for a single-process demo, NOT for production.
type sessionStore struct {
	mu      sync.RWMutex
	entries map[string]sessionEntry
}

type sessionEntry struct {
	AccessToken string
	// RefreshToken is what POST /api/token/refresh exchanges for a new access
	// token. Stored alongside it because the refresh operation needs the
	// refresh token, not the access token, and the demo has nowhere else to
	// keep it. Empty when the tenant did not return one.
	RefreshToken string
	CreatedAt    time.Time
}

func newSessionStore() *sessionStore {
	return &sessionStore{entries: make(map[string]sessionEntry)}
}

// Create issues a new session id, stores the tokens under it, and returns the
// id so the caller can set it on the response cookie. refreshToken may be
// empty — not every tenant configuration returns one.
func (s *sessionStore) Create(accessToken, refreshToken string) (string, error) {
	id, err := randomID()
	if err != nil {
		return "", err
	}
	s.mu.Lock()
	s.entries[id] = sessionEntry{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now(),
	}
	s.mu.Unlock()
	return id, nil
}

// Lookup returns the access_token bound to the session id, or "" if unknown.
func (s *sessionStore) Lookup(id string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.entries[id].AccessToken
}

// LookupRefresh returns the refresh_token bound to the session id, or "" when
// the session is unknown or the tenant returned no refresh token.
func (s *sessionStore) LookupRefresh(id string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.entries[id].RefreshToken
}

// Replace swaps the tokens held under an existing session id, used after a
// successful refresh so the rotated tokens take effect without forcing the
// user to sign in again.
func (s *sessionStore) Replace(id, accessToken, refreshToken string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	entry, ok := s.entries[id]
	if !ok {
		return
	}
	entry.AccessToken = accessToken
	if refreshToken != "" {
		entry.RefreshToken = refreshToken
	}
	s.entries[id] = entry
}

// Delete drops the session, used on logout.
func (s *sessionStore) Delete(id string) {
	s.mu.Lock()
	delete(s.entries, id)
	s.mu.Unlock()
}

func setSessionCookie(w http.ResponseWriter, id string) {
	http.SetCookie(w, &http.Cookie{
		Name:     sessionCookieName,
		Value:    id,
		Path:     "/",
		HttpOnly: true,
		// Demo only: Secure/SameSite/MaxAge omitted. A production cookie would
		// set Secure=true, SameSite=http.SameSiteLaxMode, and a sensible
		// MaxAge plus a signed value.
	})
}

func clearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    sessionCookieName,
		Value:   "",
		Path:    "/",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})
}

func getSessionCookie(r *http.Request) string {
	return cookieValue(r, sessionCookieName)
}

// The MFA cookie carries the second-factor token between a login that returned
// a challenge and the call that completes it. It is NOT a session: nothing that
// reads the session cookie will accept it, which is the point of keeping the
// two apart. Cleared as soon as the challenge is completed or abandoned.
func setMFACookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     mfaCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		// Demo only, exactly as for the session cookie: a production cookie
		// would set Secure, SameSite, and a short MaxAge matching the
		// challenge's own expiry.
	})
}

func clearMFACookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    mfaCookieName,
		Value:   "",
		Path:    "/",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})
}

func getMFACookie(r *http.Request) string {
	return cookieValue(r, mfaCookieName)
}

func cookieValue(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return c.Value
}

func randomID() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
