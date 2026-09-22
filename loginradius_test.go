package loginradius

// Parity tests for the manifest-driven behaviour of the v12 facade.
//
// These assert the CONCRETE values that the shared configuration in the LoginRadius SDK generator
// repository declares: header names, query-parameter names, credential
// precedence, base-URL resolution, error classification, and the full service
// set. They are deliberately literal — if you change the manifest, these tests
// should fail and be updated in the same change, which is what makes a
// behavioural change to the SDKs visible rather than silent.
//
// The Node SDK's __tests__/ suite asserts the same contract on its side. When
// you add a case here, add the mirror case there.

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// --------------------------------------------------------------- base URL --

func TestResolveBaseURLPrecedence(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		want string
	}{
		{
			name: "default when nothing is set",
			opts: nil,
			want: "https://api.loginradius.com",
		},
		{
			name: "domain builds the hosted-page server",
			opts: []Option{WithDomain("acme")},
			want: "https://acme.hub.loginradius.com",
		},
		{
			name: "customDomain wins over domain",
			opts: []Option{WithDomain("acme"), WithCustomDomain("id.acme.com")},
			want: "https://id.acme.com",
		},
		{
			name: "baseURL wins over everything",
			opts: []Option{
				WithDomain("acme"),
				WithCustomDomain("id.acme.com"),
				WithBaseURL("https://staging.internal"),
			},
			want: "https://staging.internal",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := defaultConfig()
			for _, o := range tc.opts {
				o(cfg)
			}
			if got := resolveBaseURL(cfg); got != tc.want {
				t.Errorf("resolveBaseURL() = %q, want %q", got, tc.want)
			}
		})
	}
}

// -------------------------------------------------------------- validation --

func TestValidateRequiresACredential(t *testing.T) {
	if err := defaultConfig().validate(); err == nil {
		t.Fatal("validate() with no credentials = nil, want an error")
	}

	// Every credential the manifest marks `satisfiesValidation: true` must be
	// sufficient on its own.
	sufficient := map[string]Option{
		"apiKey":             WithAPIKey("k"),
		"clientID":           WithClientID("c"),
		"xLoginRadiusAPIKey": WithXLoginRadiusAPIKey("k"),
		"accessToken":        WithAccessToken("t"),
		"bearerToken":        WithBearerToken("t"),
		"m2mBearerToken":     WithM2MBearerToken("t"),
	}
	for name, opt := range sufficient {
		t.Run(name, func(t *testing.T) {
			cfg := defaultConfig()
			opt(cfg)
			if err := cfg.validate(); err != nil {
				t.Errorf("validate() with only %s = %v, want nil", name, err)
			}
		})
	}

	// A credential NOT marked satisfiesValidation must not be sufficient:
	// apiSecret alone is a misconfiguration, not a usable client.
	cfg := defaultConfig()
	WithAPISecret("s")(cfg)
	if err := cfg.validate(); err == nil {
		t.Error("validate() with only apiSecret = nil, want an error")
	}
}

func TestDefaults(t *testing.T) {
	cfg := defaultConfig()
	if want := 30 * time.Second; cfg.timeout != want {
		t.Errorf("default timeout = %v, want %v", cfg.timeout, want)
	}
	if want := "loginradius-go/" + Version; cfg.userAgent != want {
		t.Errorf("default userAgent = %q, want %q", cfg.userAgent, want)
	}
}

// -------------------------------------------------------- auth injection --

// recordingTransport captures the request the authTransport produced.
type recordingTransport struct{ got *http.Request }

func (rt *recordingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.got = req
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       http.NoBody,
		Header:     make(http.Header),
		Request:    req,
	}, nil
}

func roundTrip(t *testing.T, url string, opts ...Option) *http.Request {
	t.Helper()
	cfg := defaultConfig()
	for _, o := range opts {
		o(cfg)
	}
	rec := &recordingTransport{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if _, err := newAuthTransport(cfg, rec).RoundTrip(req); err != nil {
		t.Fatalf("RoundTrip: %v", err)
	}
	if rec.got == nil {
		t.Fatal("transport was not reached")
	}
	return rec.got
}

func TestAuthInjectsHeadersAndQuery(t *testing.T) {
	got := roundTrip(t, "https://api.loginradius.com/identity/v2/auth/login",
		WithAPIKey("KEY"),
		WithAPISecret("SECRET"),
		WithClientID("CID"),
		WithClientSecret("CSEC"),
		WithAccessToken("ATOK"),
	)

	wantHeaders := map[string]string{
		"X-LoginRadius-ApiKey":    "KEY",
		"X-LoginRadius-ApiSecret": "SECRET",
		"User-Agent":              "loginradius-go/" + Version,
	}
	for k, want := range wantHeaders {
		if v := got.Header.Get(k); v != want {
			t.Errorf("header %s = %q, want %q", k, v, want)
		}
	}

	wantQuery := map[string]string{
		"apikey":        "KEY",
		"apisecret":     "SECRET",
		"client_id":     "CID",
		"client_secret": "CSEC",
		"access_token":  "ATOK",
	}
	q := got.URL.Query()
	for k, want := range wantQuery {
		if v := q.Get(k); v != want {
			t.Errorf("query %s = %q, want %q", k, v, want)
		}
	}
}

func TestAuthHeaderOverrideOptions(t *testing.T) {
	// The X-prefixed options override the HEADER value only; the query
	// parameter must keep carrying the base credential.
	got := roundTrip(t, "https://api.loginradius.com/x",
		WithAPIKey("BASE_KEY"),
		WithAPISecret("BASE_SECRET"),
		WithXLoginRadiusAPIKey("HDR_KEY"),
		WithXLoginRadiusAPISecret("HDR_SECRET"),
	)

	if v := got.Header.Get("X-LoginRadius-ApiKey"); v != "HDR_KEY" {
		t.Errorf("X-LoginRadius-ApiKey = %q, want the override %q", v, "HDR_KEY")
	}
	if v := got.Header.Get("X-LoginRadius-ApiSecret"); v != "HDR_SECRET" {
		t.Errorf("X-LoginRadius-ApiSecret = %q, want the override %q", v, "HDR_SECRET")
	}
	if v := got.URL.Query().Get("apikey"); v != "BASE_KEY" {
		t.Errorf("apikey query = %q, want the base credential %q", v, "BASE_KEY")
	}
	if v := got.URL.Query().Get("apisecret"); v != "BASE_SECRET" {
		t.Errorf("apisecret query = %q, want the base credential %q", v, "BASE_SECRET")
	}
}

func TestAuthBearerPrecedence(t *testing.T) {
	// bearerToken has the lower precedence number in the manifest, so it wins.
	got := roundTrip(t, "https://api.loginradius.com/x",
		WithBearerToken("PLAIN"),
		WithM2MBearerToken("M2M"),
	)
	if v := got.Header.Get("Authorization"); v != "Bearer PLAIN" {
		t.Errorf("Authorization = %q, want %q", v, "Bearer PLAIN")
	}

	// With only the M2M token set, it is used.
	got = roundTrip(t, "https://api.loginradius.com/x", WithM2MBearerToken("M2M"))
	if v := got.Header.Get("Authorization"); v != "Bearer M2M" {
		t.Errorf("Authorization = %q, want %q", v, "Bearer M2M")
	}
}

func TestAuthDoesNotOverridePerCallQueryParams(t *testing.T) {
	// A value already on the URL is a per-call override and must survive.
	got := roundTrip(t,
		"https://api.loginradius.com/x?access_token=PER_CALL",
		WithAPIKey("KEY"), WithAccessToken("CLIENT_WIDE"),
	)
	if v := got.URL.Query().Get("access_token"); v != "PER_CALL" {
		t.Errorf("access_token = %q, want the per-call value %q", v, "PER_CALL")
	}
	if v := got.URL.Query().Get("apikey"); v != "KEY" {
		t.Errorf("apikey = %q, want %q", v, "KEY")
	}
}

func TestAuthOmitsUnsetCredentials(t *testing.T) {
	got := roundTrip(t, "https://api.loginradius.com/x", WithAPIKey("KEY"))
	if _, ok := got.Header["X-Loginradius-Apisecret"]; ok {
		t.Error("X-LoginRadius-ApiSecret was set despite no apiSecret being configured")
	}
	if got.Header.Get("Authorization") != "" {
		t.Error("Authorization was set despite no bearer token being configured")
	}
	if _, present := got.URL.Query()["apisecret"]; present {
		t.Error("apisecret query parameter was set despite no apiSecret being configured")
	}
}

func TestAuthDoesNotMutateCallerRequest(t *testing.T) {
	// The RoundTripper contract: the caller's *http.Request must not be
	// visibly modified. Credentials leaking onto a reused request would be a
	// real disclosure risk.
	cfg := defaultConfig()
	WithAPIKey("KEY")(cfg)
	req, err := http.NewRequest(http.MethodGet, "https://api.loginradius.com/x", nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if _, err := newAuthTransport(cfg, &recordingTransport{}).RoundTrip(req); err != nil {
		t.Fatalf("RoundTrip: %v", err)
	}
	if req.Header.Get("X-LoginRadius-ApiKey") != "" {
		t.Error("caller's request was mutated: credential was written to the original header")
	}
	if req.URL.RawQuery != "" {
		t.Errorf("caller's request URL was mutated: RawQuery = %q, want empty", req.URL.RawQuery)
	}
}

// ------------------------------------------------------- error envelopes --

func TestExtractEnvelopeAcrossWireShapes(t *testing.T) {
	tests := []struct {
		name            string
		body            string
		wantCode        string
		wantMessage     string
		wantDescription string
	}{
		{
			name:            "ApiError / ErrorResponse (PascalCase)",
			body:            `{"ErrorCode":1043,"Message":"Invalid credentials","Description":"The email and password do not match."}`,
			wantCode:        "1043",
			wantMessage:     "Invalid credentials",
			wantDescription: "The email and password do not match.",
		},
		{
			name:            "ErrorResponseNative (camelCase, /api/v2/*)",
			body:            `{"errorCode":1066,"message":"Token expired","description":"The access token has expired."}`,
			wantCode:        "1066",
			wantMessage:     "Token expired",
			wantDescription: "The access token has expired.",
		},
		{
			name:            "OAuthErrorResponse (token endpoints)",
			body:            `{"error":"invalid_grant","error_description":"Refresh token is invalid."}`,
			wantCode:        "invalid_grant",
			wantDescription: "Refresh token is invalid.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env, ok := extractEnvelope([]byte(tc.body))
			if !ok {
				t.Fatalf("extractEnvelope(%s) reported no match", tc.body)
			}
			if env.Code != tc.wantCode {
				t.Errorf("Code = %q, want %q", env.Code, tc.wantCode)
			}
			if env.Message != tc.wantMessage {
				t.Errorf("Message = %q, want %q", env.Message, tc.wantMessage)
			}
			if env.Description != tc.wantDescription {
				t.Errorf("Description = %q, want %q", env.Description, tc.wantDescription)
			}
		})
	}
}

func TestExtractEnvelopeRejectsUnusableBodies(t *testing.T) {
	for _, body := range []string{``, `not json at all`, `{}`, `{"unrelated":"value"}`} {
		if _, ok := extractEnvelope([]byte(body)); ok {
			t.Errorf("extractEnvelope(%q) reported a match, want none", body)
		}
	}
}

func TestErrorPredicates(t *testing.T) {
	tests := []struct {
		status                             int
		auth, forbidden, rateLimit, server bool
	}{
		{status: 401, auth: true},
		{status: 403, forbidden: true},
		{status: 429, rateLimit: true},
		{status: 500, server: true},
		{status: 503, server: true},
		{status: 599, server: true},
		{status: 400},
		{status: 200},
		{status: 0},
	}
	for _, tc := range tests {
		e := &Error{StatusCode: tc.status}
		if e.IsAuth() != tc.auth {
			t.Errorf("status %d: IsAuth() = %v, want %v", tc.status, e.IsAuth(), tc.auth)
		}
		if e.IsForbidden() != tc.forbidden {
			t.Errorf("status %d: IsForbidden() = %v, want %v", tc.status, e.IsForbidden(), tc.forbidden)
		}
		if e.IsRateLimit() != tc.rateLimit {
			t.Errorf("status %d: IsRateLimit() = %v, want %v", tc.status, e.IsRateLimit(), tc.rateLimit)
		}
		if e.IsServer() != tc.server {
			t.Errorf("status %d: IsServer() = %v, want %v", tc.status, e.IsServer(), tc.server)
		}
	}
}

func TestDefaultDescriptionForStatus(t *testing.T) {
	// An empty 403 body is the common case behind an IP/domain restriction —
	// the hint is the only thing the caller gets, so it must be present.
	got := defaultDescriptionForStatus(http.StatusForbidden, nil)
	if !strings.Contains(got, "IP-access restriction") {
		t.Errorf("403 hint = %q, want it to mention the IP-access restriction cause", got)
	}

	// A non-JSON body is appended so the caller can see what really replied.
	got = defaultDescriptionForStatus(http.StatusForbidden, []byte("<html>\n  blocked by\tWAF</html>"))
	if !strings.Contains(got, "body: <html> blocked by WAF</html>") {
		t.Errorf("403 with body = %q, want a whitespace-collapsed body excerpt", got)
	}

	// An unmapped status with no body has nothing useful to say.
	if got := defaultDescriptionForStatus(418, nil); got != "" {
		t.Errorf("418 hint = %q, want empty", got)
	}
}

func TestBodyExcerptTruncates(t *testing.T) {
	got := bodyExcerpt([]byte(strings.Repeat("x", 500)), 240)
	// 240 characters plus the ellipsis rune.
	if want := strings.Repeat("x", 240) + "…"; got != want {
		t.Errorf("bodyExcerpt truncation = %d chars, want %d plus an ellipsis",
			len([]rune(got)), 240)
	}
}

func TestAsErrorIgnoresUnknownErrors(t *testing.T) {
	if _, ok := AsError(nil); ok {
		t.Error("AsError(nil) reported a match, want none")
	}
	if _, ok := AsError(http.ErrHandlerTimeout); ok {
		t.Error("AsError(transport error) reported a match, want none")
	}
	// An *Error already in the chain is returned as-is.
	orig := &Error{StatusCode: 403, Code: "1234"}
	got, ok := AsError(orig)
	if !ok || got != orig {
		t.Errorf("AsError(*Error) = (%v, %v), want the original error", got, ok)
	}
}

// -------------------------------------------------------------- the client --

func TestNewClientRequiresCredentials(t *testing.T) {
	if _, err := NewClient(); err == nil {
		t.Error("NewClient() with no options = nil error, want a misconfiguration error")
	}
	if _, err := NewClient(WithAPIKey("k")); err != nil {
		t.Errorf("NewClient(WithAPIKey) = %v, want nil", err)
	}
}

func TestNewClientWiresEveryService(t *testing.T) {
	// The service list is derived from the generated client, so this guards
	// against a service being generated but never surfaced on the facade —
	// the exact gap that left OAuthM2M unreachable in Go before the SDKs were
	// generated from a shared manifest.
	c, err := NewClient(WithAPIKey("k"))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	for _, tc := range []struct {
		name string
		got  any
	}{
		{"Accounts", c.Accounts},
		{"Login", c.Login},
		{"Registration", c.Registration},
		{"OAuth", c.OAuth},
		{"OAuthM2M", c.OAuthM2M},
		{"OIDC", c.OIDC},
		{"SAML", c.SAML},
		{"JWT", c.JWT},
		{"SOTT", c.SOTT},
		{"IPAccessRestrictions", c.IPAccessRestrictions},
		{"Workflows", c.Workflows},
	} {
		if tc.got == nil {
			t.Errorf("client.%s is nil, want a wired service", tc.name)
		}
	}
}

func TestClientAppliesTimeoutAndPreservesCallerTransport(t *testing.T) {
	// A caller-supplied client keeps its own Transport underneath ours, so
	// their proxy/TLS/dialer configuration still runs.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	base := &recordingTransport{}
	caller := &http.Client{Transport: base, Timeout: 5 * time.Second}

	c, err := NewClient(WithAPIKey("k"), WithHTTPClient(caller))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if c.HTTPClient() != caller {
		t.Error("HTTPClient() did not return the caller's client")
	}
	if caller.Timeout != 5*time.Second {
		t.Errorf("caller timeout = %v, want it preserved at 5s", caller.Timeout)
	}

	// Driving a request through proves our transport wrapped, not replaced.
	req, err := http.NewRequest(http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if _, err := caller.Do(req); err != nil {
		t.Fatalf("Do: %v", err)
	}
	if base.got == nil {
		t.Fatal("caller's original Transport was replaced, not wrapped")
	}
	if base.got.Header.Get("X-LoginRadius-ApiKey") != "k" {
		t.Error("credentials were not injected before reaching the caller's Transport")
	}
}

// ------------------------------------------------------------------- SOTT --

// ---------------------------------------------------------------- signing --

// TestSigningGoldenValue pins the exact digest the signing algorithm must
// produce.
//
// CROSS-LANGUAGE PARITY: the golden values came from running the REFERENCE
// implementation — `creatHashForApplicationSignig` in
// admin-console-backend/lib/services/loginradius-v2-sdk/lr.js — against these
// inputs, and the Node SDK asserts the same values in __tests__/signing.test.ts.
// They are not copied from this SDK's own output, so this proves the port
// matches the known-working algorithm rather than merely matching itself.
//
// If this fails, fix the implementation, not the expectation. The shared
// parameters live in the SDK generator's shared configuration under `signing:`.
func TestSigningGoldenValue(t *testing.T) {
	const (
		secret        = "test-api-secret"
		uri           = "https://api.loginradius.com/identity/v2/manage/account/uid?apikey=test-api-key"
		goldenNoBody  = "SHA-256=WhdnDwiFzLUkrhBOUQYzrec+ZllDrY6X0hdovov8bKY="
		goldenBody    = "SHA-256=cVEPKKM+Dd1fzQePAPDKMKCn+2QollSbzWfVx8YxSzM="
		goldenExpires = "2026-01-02 03:24:05"
	)
	// signRequest adds the 20-minute window, so now is the golden expiry - 20m.
	now := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)

	got, expires := signRequest(secret, uri, nil, now)
	if got != goldenNoBody {
		t.Errorf("no-body digest drifted:\ngot:  %s\nwant: %s", got, goldenNoBody)
	}
	if expires != goldenExpires {
		t.Errorf("expiry = %q, want %q", expires, goldenExpires)
	}

	got, _ = signRequest(secret, uri, []byte(`{"Uid":"abc123"}`), now)
	if got != goldenBody {
		t.Errorf("with-body digest drifted:\ngot:  %s\nwant: %s", got, goldenBody)
	}
}

// TestEncodeURIComponentMatchesJavaScript guards the escaper.
//
// Go's url.QueryEscape is NOT encodeURIComponent: it renders a space as "+"
// and escapes !'()*. Signing a URL containing an email address (which may
// legally contain those characters) with the wrong escaper yields a digest the
// API rejects, so the character set is asserted explicitly.
func TestEncodeURIComponentMatchesJavaScript(t *testing.T) {
	// Verified against Node: encodeURIComponent("a b!'()*~-_.é/?&=")
	const want = "a%20b!'()*~-_.%C3%A9%2F%3F%26%3D"
	if got := encodeURIComponent("a b!'()*~-_.é/?&="); got != want {
		t.Errorf("encodeURIComponent mismatch:\ngot:  %s\nwant: %s", got, want)
	}
}

func TestShouldSignRequest(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"/identity/v2/manage/account/uid", true},
		{"/v2/manage/roles", true},
		{"/identity/v2/auth/login", false},
		{"/identity/v2/auth/register", false},
		// The access-token exchange is explicitly excluded by the reference.
		{"/identity/v2/manage/account/access_token", false},
	}
	for _, tc := range tests {
		if got := shouldSignRequest(tc.path); got != tc.want {
			t.Errorf("shouldSignRequest(%q) = %v, want %v", tc.path, got, tc.want)
		}
	}
}

// ------------------------------------------- cross-cutting request options --

func TestRequestOptionsAppliedToEveryRequest(t *testing.T) {
	got := roundTrip(t, "https://api.loginradius.com/identity/v2/auth/login",
		WithAPIKey("KEY"),
		WithOriginIP("203.0.113.7"),
		WithServerRegion("eu"),
		WithFields("Email,Uid"),
		WithPreventWebhook(true),
	)

	if v := got.Header.Get("X-Origin-IP"); v != "203.0.113.7" {
		t.Errorf("X-Origin-IP = %q, want %q", v, "203.0.113.7")
	}
	if v := got.Header.Get("X-PreventWebhook"); v != "true" {
		t.Errorf("X-PreventWebhook = %q, want %q", v, "true")
	}
	q := got.URL.Query()
	if v := q.Get("region"); v != "eu" {
		t.Errorf("region = %q, want %q", v, "eu")
	}
	if v := q.Get("fields"); v != "Email,Uid" {
		t.Errorf("fields = %q, want %q", v, "Email,Uid")
	}
}

func TestRequestOptionsOmittedWhenUnset(t *testing.T) {
	got := roundTrip(t, "https://api.loginradius.com/identity/v2/auth/login", WithAPIKey("KEY"))

	for _, h := range []string{"X-Origin-IP", "X-PreventWebhook"} {
		if v := got.Header.Get(h); v != "" {
			t.Errorf("%s = %q, want it unset", h, v)
		}
	}
	for _, p := range []string{"region", "fields"} {
		if _, present := got.URL.Query()[p]; present {
			t.Errorf("%s was set despite no option being configured", p)
		}
	}
}

func TestDefaultHeadersNeverMaskACredential(t *testing.T) {
	// A default header is applied first so the SDK's own headers overwrite it.
	// Letting a caller override X-LoginRadius-ApiKey here would silently send
	// the wrong credential.
	got := roundTrip(t, "https://api.loginradius.com/identity/v2/auth/login",
		WithAPIKey("REAL_KEY"),
		WithDefaultHeaders(map[string]string{
			"X-Tenant-Trace":       "abc123",
			"X-LoginRadius-ApiKey": "HIJACKED",
			"User-Agent":           "HIJACKED",
		}),
	)

	if v := got.Header.Get("X-Tenant-Trace"); v != "abc123" {
		t.Errorf("X-Tenant-Trace = %q, want it merged", v)
	}
	if v := got.Header.Get("X-LoginRadius-ApiKey"); v != "REAL_KEY" {
		t.Errorf("credential header = %q, want the SDK's value to win", v)
	}
	if v := got.Header.Get("User-Agent"); v == "HIJACKED" {
		t.Error("User-Agent was overridden by a default header")
	}
}

func TestSigningAppliedOnlyToManagementPaths(t *testing.T) {
	opts := []Option{WithAPIKey("KEY"), WithAPISecret("SECRET"), WithAPIRequestSigning(true)}

	signed := roundTrip(t, "https://api.loginradius.com/identity/v2/manage/account/uid", opts...)
	if signed.Header.Get(signingDigestHeader) == "" {
		t.Error("management request was not signed")
	}
	if signed.Header.Get(signingExpiresHeader) == "" {
		t.Error("management request has no expiry stamp")
	}

	unsigned := roundTrip(t, "https://api.loginradius.com/identity/v2/auth/login", opts...)
	if v := unsigned.Header.Get(signingDigestHeader); v != "" {
		t.Errorf("auth request was signed (%q); only /manage/ paths should be", v)
	}

	excluded := roundTrip(t, "https://api.loginradius.com/identity/v2/manage/account/access_token", opts...)
	if v := excluded.Header.Get(signingDigestHeader); v != "" {
		t.Errorf("access_token exchange was signed (%q); it is explicitly excluded", v)
	}
}

func TestSigningStripsTheSecretFromTheURL(t *testing.T) {
	// The tenant secret must not travel on a signed request, nor be part of
	// the string that is signed.
	got := roundTrip(t, "https://api.loginradius.com/identity/v2/manage/account/uid",
		WithAPIKey("KEY"), WithAPISecret("SECRET"), WithAPIRequestSigning(true))

	if v := got.URL.Query().Get(signingStripParam); v != "" {
		t.Errorf("%s survived on a signed request: %q", signingStripParam, v)
	}
	if strings.Contains(got.URL.String(), "SECRET") {
		t.Errorf("secret value present in signed URL: %s", got.URL.String())
	}
}

func TestSigningRequiresBothOptInAndSecret(t *testing.T) {
	// Enabled but no secret: nothing to sign with, so no header rather than a
	// signature over an empty key.
	got := roundTrip(t, "https://api.loginradius.com/identity/v2/manage/account/uid",
		WithAPIKey("KEY"), WithAPIRequestSigning(true))
	if v := got.Header.Get(signingDigestHeader); v != "" {
		t.Errorf("signed without an apiSecret: %q", v)
	}

	// Secret present but signing not enabled: off by default.
	got = roundTrip(t, "https://api.loginradius.com/identity/v2/manage/account/uid",
		WithAPIKey("KEY"), WithAPISecret("SECRET"))
	if v := got.Header.Get(signingDigestHeader); v != "" {
		t.Errorf("signed without opting in: %q", v)
	}
}

func TestDebugRedactsCredentialValues(t *testing.T) {
	var buf bytes.Buffer
	roundTrip(t, "https://api.loginradius.com/identity/v2/auth/login",
		WithAPIKey("SUPER_SECRET_KEY"), WithAPISecret("SUPER_SECRET_VALUE"),
		WithBearerToken("SUPER_SECRET_TOKEN"), WithDebug(&buf))

	out := buf.String()
	if out == "" {
		t.Fatal("debug writer received nothing")
	}
	for _, secret := range []string{"SUPER_SECRET_KEY", "SUPER_SECRET_VALUE", "SUPER_SECRET_TOKEN"} {
		if strings.Contains(out, secret) {
			t.Errorf("debug log leaked a credential value: %s", out)
		}
	}
	// Go canonicalises header names, so compare case-insensitively.
	if !strings.Contains(strings.ToLower(out), "x-loginradius-apikey") {
		t.Errorf("debug log should name the header: %s", out)
	}
	if !strings.Contains(out, "[REDACTED]") {
		t.Errorf("debug log should mark redacted values: %s", out)
	}
}

func TestTimeoutOnlyAppliedToInjectedClientWhenExplicit(t *testing.T) {
	// Caller supplies a client and an explicit timeout: honour it.
	caller := &http.Client{Timeout: 5 * time.Second}
	if _, err := NewClient(WithAPIKey("k"), WithHTTPClient(caller), WithTimeout(90*time.Second)); err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if caller.Timeout != 90*time.Second {
		t.Errorf("explicit timeout = %v, want it applied to the injected client", caller.Timeout)
	}

	// Caller supplies only a client: leave their configuration alone rather
	// than imposing the SDK default.
	untouched := &http.Client{Timeout: 5 * time.Second}
	if _, err := NewClient(WithAPIKey("k"), WithHTTPClient(untouched)); err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if untouched.Timeout != 5*time.Second {
		t.Errorf("timeout = %v, want the caller's 5s preserved", untouched.Timeout)
	}
}

// ------------------------------------------------------ operation servers --

// pinnedOperationHost drives one of the operations the spec pins to its own
// server list and reports the host the request actually went to.
func pinnedOperationHost(t *testing.T, opts ...Option) string {
	t.Helper()
	rec := &recordingTransport{}
	all := append([]Option{WithAPIKey("k"), WithHTTPClient(&http.Client{Transport: rec})}, opts...)
	c, err := NewClient(all...)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	// GetBigCommerceLoginUrl is pinned to https://{domain}.hub.loginradius.com.
	_, _, err = c.BigCommerceSSO.GetBigCommerceLoginUrl(context.Background()).
		AccessToken("tok").Store("mystore").Execute()
	if rec.got == nil {
		t.Fatalf("no request reached the transport: %v", err)
	}
	return rec.got.URL.Host
}

// TestOperationServersHonourClientConfiguration covers the defect where the
// generator pins 42 operations to their own hosts and the facade passed only a
// basePath — so an explicitly configured baseURL was silently ignored for all
// of them, and tenant-templated hosts always resolved to the spec's
// placeholder tenant.
func TestOperationServersHonourClientConfiguration(t *testing.T) {
	// No server options: the spec's own placeholder stands.
	if got := pinnedOperationHost(t); got != "example.hub.loginradius.com" {
		t.Errorf("default host = %q, want the spec placeholder", got)
	}

	// WithDomain fills the {domain} template variable instead of leaving the
	// placeholder in place.
	if got := pinnedOperationHost(t, WithDomain("acme")); got != "acme.hub.loginradius.com" {
		t.Errorf("WithDomain host = %q, want acme.hub.loginradius.com", got)
	}

	// An explicit baseURL means "send everything here" and wins over the pin.
	if got := pinnedOperationHost(t, WithBaseURL("https://staging.internal")); got != "staging.internal" {
		t.Errorf("WithBaseURL host = %q, want staging.internal", got)
	}
}

// ------------------------------------------ per-call vs client-wide params --

// TestPerCallParamsWinOverClientWide guards the precedence rule for every
// query-based option and credential.
//
// A duplicated key is not merely untidy: the server chooses between the two
// values non-deterministically, so a duplicated access_token can authenticate
// the request as the wrong principal. The Node SDK asserts the same cases in
// __tests__/request-options.test.ts.
func TestPerCallParamsWinOverClientWide(t *testing.T) {
	// fields: client-wide only.
	rec := &recordingTransport{}
	c, err := NewClient(WithAPIKey("k"), WithAPISecret("s"), WithFields("Email,Uid"),
		WithHTTPClient(&http.Client{Transport: rec}))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	_, _, _ = c.Accounts.GetIdentities(context.Background()).Email("a@b.c").Execute()
	if got := rec.got.URL.Query()["fields"]; len(got) != 1 || got[0] != "Email,Uid" {
		t.Errorf("client-wide fields = %v, want exactly [Email,Uid]", got)
	}

	// fields: both set — the per-call value wins and appears once.
	rec = &recordingTransport{}
	c, _ = NewClient(WithAPIKey("k"), WithAPISecret("s"), WithFields("Email,Uid"),
		WithHTTPClient(&http.Client{Transport: rec}))
	_, _, _ = c.Accounts.GetIdentities(context.Background()).
		Email("a@b.c").Fields("FirstName").Execute()
	if got := rec.got.URL.Query()["fields"]; len(got) != 1 || got[0] != "FirstName" {
		t.Errorf("fields = %v, want exactly [FirstName] (per-call wins, no duplicate)", got)
	}

	// access_token: both set. A duplicate here is a correctness problem, not a
	// cosmetic one.
	rec = &recordingTransport{}
	c, _ = NewClient(WithAPIKey("k"), WithAccessToken("CLIENT_TOKEN"),
		WithHTTPClient(&http.Client{Transport: rec}))
	_, _, _ = c.User.GetAccountDetails(context.Background()).AccessToken("PER_CALL_TOKEN").Execute()
	if got := rec.got.URL.Query()["access_token"]; len(got) != 1 || got[0] != "PER_CALL_TOKEN" {
		t.Errorf("access_token = %v, want exactly [PER_CALL_TOKEN]", got)
	}
}
