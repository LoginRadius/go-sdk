package main

import (
	"context"
	"embed"
	"encoding/json"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

//go:embed static/*
var staticFS embed.FS

const errMethodNotAllowed = "method not allowed"

type server struct {
	client   *loginradius.Client
	sessions *sessionStore
	// apiKey + apiSecret are needed by handleRegister to mint a SOTT —
	// the SDK doesn't expose its config back to callers, so we keep our
	// own copy here. Read-only after NewClient.
	apiKey    string
	apiSecret string
	// baseOpts are the tenant options (credentials + server selection) that
	// built `client`. userClient replays them with an access token added.
	baseOpts []loginradius.Option
}

// userClient builds a per-request client carrying the signed-in user's access
// token, mirroring the Node demo's userClient(). Most user-context operations
// also accept the token as a per-call builder argument, but some — notably
// Session.InvalidateAccessToken and User.Deleteemailbyaccesstoken — read it
// only from the client configuration, so a user-context client is the one
// approach that works for all of them.
func (s *server) userClient(accessToken string) (*loginradius.Client, error) {
	opts := make([]loginradius.Option, 0, len(s.baseOpts)+1)
	opts = append(opts, s.baseOpts...)
	opts = append(opts, loginradius.WithAccessToken(accessToken))
	return loginradius.NewClient(opts...)
}

// ---------------------------------------------------------------------------
// Helpers the generated wrappers in routes_gen.go call.
//
// They live here rather than in the generated file so that file imports only
// net/http and its import block cannot go stale as routes are added to or
// removed from the manifest.
// ---------------------------------------------------------------------------

// decodeJSON reads a JSON request body into dst.
func decodeJSON(r *http.Request, dst any) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

// demoEnv reads a demo configuration value with no fallback. Used by generated
// wrappers for values the route cannot work without.
func demoEnv(key string) string { return os.Getenv(key) }

func badRequest(w http.ResponseWriter, msg string) {
	writeJSON(w, http.StatusBadRequest, map[string]any{"error": msg})
}

func missingField(w http.ResponseWriter, field string) {
	writeJSON(w, http.StatusBadRequest, map[string]any{
		"error": "missing required field",
		"field": field,
	})
}

// configMissing reports a demo that is not configured for this route, rather
// than letting the call through to fail upstream with an opaque message.
func configMissing(w http.ResponseWriter, envVar string) {
	writeJSON(w, http.StatusInternalServerError, map[string]any{
		"error": envVar + " is not set",
		"hint":  "this route needs " + envVar + " in demo/.env — see demo/.env.example",
	})
}

// optStr turns an empty string into a nil *string, so optional fields are
// omitted from the request body rather than sent as "".
func optStr(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

// verificationURL is where LoginRadius sends the user from a verification
// email. It must point at the demo's own verify route (with `?vtoken=…`
// appended by LoginRadius) for the loop to close.
func verificationURL() string {
	return envOr("LR_VERIFICATION_URL", "http://localhost:8080/api/auth/verify")
}

// resetPasswordURL is where LoginRadius sends the user from a reset email.
func resetPasswordURL() string {
	return envOr("LR_RESET_PASSWORD_URL", "http://localhost:8080/reset")
}

// totpBody builds the body both TOTP endpoints take.
//
// LoginRadius has two authenticator generations and they do not share a field
// name: a tenant on Google Authenticator requires `googleauthenticatorcode`,
// while the newer generic authenticator uses `authenticatorcode`. Sending only
// the latter returns ErrorCode 908, "The googleauthenticatorcode is a required
// parameter." The spec declared only `authenticatorcode`, which is why this was
// broken until the schema gained both.
//
// Only ONE field is populated, deliberately. The sibling reauth schema
// (ReAuthTwoFAModelCore) declares its equivalent code fields under `oneOf` with
// each one `required`, so the API treats them as mutually exclusive there —
// sending both risks a validation rejection rather than a helpful fallback.
// `googleauthenticatorcode` is the one the live endpoint demands.
//
// A tenant on the newer generic authenticator needs `Authenticatorcode` here
// instead; the field exists on the model, so that is a one-line change.
func totpBody(code string) loginradius.AuthenticatorCodeRequest {
	return loginradius.AuthenticatorCodeRequest{
		Googleauthenticatorcode: &code,
	}
}

// remapJSON converts a decoded JSON object into a typed SDK model by
// round-tripping it through JSON. The passkey routes receive whatever
// navigator.credentials produced; re-encoding is the least surprising way to
// hand that to a generated struct without transcribing every WebAuthn field.
func remapJSON(src any, dst any) error {
	buf, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, dst)
}

// handleIndex serves the single-page demo UI.
// handleIndex serves the demo's static assets: index.html at "/", and anything
// else under static/ by name — demo.css among them, which is generated from a
// template shared by every LoginRadius SDK so the demos look the same.
//
// Serving the whole embedded directory rather than just index.html: a page that
// links a stylesheet the server will not hand out renders unstyled, and it does
// so silently, with a 404 nobody sees unless they open devtools.
func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	name := "static/index.html"
	if r.URL.Path != "/" {
		// path.Clean + the static/ prefix keeps this inside the embedded FS;
		// embed.FS also rejects "..", so traversal cannot escape it.
		name = "static/" + strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	}
	page, err := staticFS.ReadFile(name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if ct := mime.TypeByExtension(filepath.Ext(name)); ct != "" {
		w.Header().Set("Content-Type", ct)
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
	_, _ = w.Write(page)
}

// ---------------------------------------------------------------------------
// Hand-written handlers.
//
// These are the routes the manifest marks as needing bespoke logic: minting a
// SOTT, setting or clearing a session, redirecting, or enforcing a demo-side
// rule. Everything else gets its wrapper generated and appears further down as
// a call<Name>.
// ---------------------------------------------------------------------------

// handleRegister wraps client.Registration.UserRegistrationBySottEmailPhoneUserName.
// Mints the SOTT via loginradius.GenerateSOTT — server-side only; the secret
// never leaves the demo process.
func (s *server) handleRegister(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	if err := decodeJSON(r, &in); err != nil {
		badRequest(w, "expected a JSON object body")
		return
	}
	if in.Email == "" {
		missingField(w, "email")
		return
	}
	if in.Password == "" {
		missingField(w, "password")
		return
	}

	sott, err := loginradius.GenerateSOTT(s.apiKey, s.apiSecret)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]any{
			"error": "generate SOTT: " + err.Error(),
		})
		return
	}

	emailType := "Primary"
	body := loginradius.ProfileRequestModel{
		Email: []loginradius.ProfileRequestModelEmailInner{
			{Type: &emailType, Value: &in.Email},
		},
		Password:  &in.Password,
		FirstName: optStr(in.FirstName),
		LastName:  optStr(in.LastName),
	}

	resp, _, err := s.client.Registration.
		UserRegistrationBySottEmailPhoneUserName(r.Context()).
		Sott(sott).
		ProfileRequestModel(body).
		Verificationurl(verificationURL()).
		Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}

	// The manifest marks this route setsSession. Tenants configured to require
	// email verification return a profile with no access token, so a session is
	// minted only when one actually came back — and the response says which
	// happened, because "registered but not signed in" is otherwise a confusing
	// state to land in.
	signedIn := false
	if access, refresh := registrationTokens(resp); access != "" {
		id, err := s.sessions.Create(access, refresh)
		if err != nil {
			http.Error(w, "could not mint session", http.StatusInternalServerError)
			return
		}
		setSessionCookie(w, id)
		signedIn = true
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"signed_in": signedIn,
		"profile":   resp,
	})
}

// handleLogin wraps client.Login.EmailByLoginUserNamePhone (LoginByEmailRequest
// branch). On success, mints a demo session. When the tenant requires a second
// factor, the challenge token goes into its own short-lived cookie and the
// caller is told which factors are available — see the /api/mfa/login/* routes.
func (s *server) handleLogin(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := decodeJSON(r, &in); err != nil {
		badRequest(w, "expected a JSON object body")
		return
	}
	if in.Email == "" {
		missingField(w, "email")
		return
	}
	if in.Password == "" {
		missingField(w, "password")
		return
	}

	body := loginradius.EmailByLoginUserNamePhoneRequest{
		LoginByEmailRequest: &loginradius.LoginByEmailRequest{
			Email:    &in.Email,
			Password: &in.Password,
		},
	}

	resp, _, err := s.client.Login.
		EmailByLoginUserNamePhone(r.Context()).
		EmailByLoginUserNamePhoneRequest(body).
		Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}

	// Login response is a oneOf — extract from whichever branch matched.
	var accessToken, refreshToken string
	switch {
	case resp.AuthResponseOptionalMfa != nil && resp.AuthResponseOptionalMfa.AccessToken != nil:
		accessToken = *resp.AuthResponseOptionalMfa.AccessToken
		refreshToken = valueOr(resp.AuthResponseOptionalMfa.RefreshToken)
	case resp.AuthResponseRequiredMfa != nil:
		// A second factor is required. Park the challenge token in its own
		// cookie so the follow-up call can complete it, and tell the UI which
		// factors this account has so it can offer the right ones.
		mfa := resp.AuthResponseRequiredMfa
		if mfa.SecondFactorAuthenticationToken == nil || *mfa.SecondFactorAuthenticationToken == "" {
			writeJSON(w, http.StatusBadGateway, map[string]any{
				"error": "MFA required but no SecondFactorAuthenticationToken returned",
			})
			return
		}
		setMFACookie(w, *mfa.SecondFactorAuthenticationToken)
		writeJSON(w, http.StatusOK, map[string]any{
			"mfa_required":      true,
			"totp_enrolled":     boolOr(mfa.IsGoogleAuthenticatorVerified) || boolOr(mfa.IsAuthenticatorVerified),
			"manual_entry_code": mfa.ManualEntryCode.Get(),
			"qr_code":           mfa.QRCode.Get(),
			"response":          mfa,
		})
		return
	}
	if accessToken == "" {
		writeJSON(w, http.StatusBadGateway, map[string]any{
			"error": "login succeeded but no access_token returned",
		})
		return
	}

	id, err := s.sessions.Create(accessToken, refreshToken)
	if err != nil {
		http.Error(w, "could not mint session", http.StatusInternalServerError)
		return
	}
	setSessionCookie(w, id)
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// handlePasswordlessLoginByEmailOtp completes an email passwordless login.
// Same oneOf response shape as handleLogin, including the MFA-challenge
// branch — a tenant with MFA enabled still enforces its second factor after
// the emailed code.
func (s *server) handlePasswordlessLoginByEmailOtp(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email string `json:"email"`
		Otp   string `json:"otp"`
	}
	if err := decodeJSON(r, &in); err != nil {
		badRequest(w, "expected a JSON object body")
		return
	}
	if in.Email == "" {
		missingField(w, "email")
		return
	}
	if in.Otp == "" {
		missingField(w, "otp")
		return
	}

	resp, _, err := s.client.Login.
		PasswordlessLoginByEmailAndOTP(r.Context()).
		PasswordLessEmailOTPModel(loginradius.PasswordLessEmailOTPModel{Otp: in.Otp, Email: in.Email}).
		Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}

	switch {
	case resp.AuthResponseOptionalMfa != nil && resp.AuthResponseOptionalMfa.AccessToken != nil:
		id, err := s.sessions.Create(*resp.AuthResponseOptionalMfa.AccessToken, valueOr(resp.AuthResponseOptionalMfa.RefreshToken))
		if err != nil {
			http.Error(w, "could not mint session", http.StatusInternalServerError)
			return
		}
		setSessionCookie(w, id)
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	case resp.AuthResponseRequiredMfa != nil:
		mfa := resp.AuthResponseRequiredMfa
		if mfa.SecondFactorAuthenticationToken == nil || *mfa.SecondFactorAuthenticationToken == "" {
			writeJSON(w, http.StatusBadGateway, map[string]any{
				"error": "MFA required but no SecondFactorAuthenticationToken returned",
			})
			return
		}
		setMFACookie(w, *mfa.SecondFactorAuthenticationToken)
		writeJSON(w, http.StatusOK, map[string]any{
			"mfa_required":      true,
			"totp_enrolled":     boolOr(mfa.IsGoogleAuthenticatorVerified) || boolOr(mfa.IsAuthenticatorVerified),
			"manual_entry_code": mfa.ManualEntryCode.Get(),
			"qr_code":           mfa.QRCode.Get(),
			"response":          mfa,
		})
	default:
		writeJSON(w, http.StatusBadGateway, map[string]any{
			"error": "login succeeded but no access_token returned",
		})
	}
}

// handlePasswordlessLoginByPhoneOtp completes a phone passwordless login.
// Same oneOf response shape as handleLogin, including the MFA-challenge branch.
func (s *server) handlePasswordlessLoginByPhoneOtp(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Phone string `json:"phone"`
		Otp   string `json:"otp"`
	}
	if err := decodeJSON(r, &in); err != nil {
		badRequest(w, "expected a JSON object body")
		return
	}
	if in.Phone == "" {
		missingField(w, "phone")
		return
	}
	if in.Otp == "" {
		missingField(w, "otp")
		return
	}

	resp, _, err := s.client.Login.
		PasswordlessLoginPhoneVerification(r.Context()).
		PhoneOTPModel(loginradius.PhoneOTPModel{OTP: in.Otp, Phone: in.Phone}).
		Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}

	switch {
	case resp.AuthResponseOptionalMfa != nil && resp.AuthResponseOptionalMfa.AccessToken != nil:
		id, err := s.sessions.Create(*resp.AuthResponseOptionalMfa.AccessToken, valueOr(resp.AuthResponseOptionalMfa.RefreshToken))
		if err != nil {
			http.Error(w, "could not mint session", http.StatusInternalServerError)
			return
		}
		setSessionCookie(w, id)
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	case resp.AuthResponseRequiredMfa != nil:
		mfa := resp.AuthResponseRequiredMfa
		if mfa.SecondFactorAuthenticationToken == nil || *mfa.SecondFactorAuthenticationToken == "" {
			writeJSON(w, http.StatusBadGateway, map[string]any{
				"error": "MFA required but no SecondFactorAuthenticationToken returned",
			})
			return
		}
		setMFACookie(w, *mfa.SecondFactorAuthenticationToken)
		writeJSON(w, http.StatusOK, map[string]any{
			"mfa_required":      true,
			"totp_enrolled":     boolOr(mfa.IsGoogleAuthenticatorVerified) || boolOr(mfa.IsAuthenticatorVerified),
			"manual_entry_code": mfa.ManualEntryCode.Get(),
			"qr_code":           mfa.QRCode.Get(),
			"response":          mfa,
		})
	default:
		writeJSON(w, http.StatusBadGateway, map[string]any{
			"error": "login succeeded but no access_token returned",
		})
	}
}

// handleLogout invalidates the access token upstream, then clears the demo
// session. The local session is cleared even if the upstream call fails —
// otherwise a transient API error would leave the user unable to sign out.
func (s *server) handleLogout(w http.ResponseWriter, r *http.Request) {
	var err error
	if uc, cerr := s.userClient(tokenFrom(r.Context())); cerr != nil {
		err = cerr
	} else {
		_, _, err = uc.Session.InvalidateAccessToken(r.Context()).Execute()
	}

	if id := getSessionCookie(r); id != "" {
		s.sessions.Delete(id)
	}
	clearSessionCookie(w)
	// Any half-finished MFA challenge goes too — leaving it behind would keep
	// a credential alive that outlives the session it belonged to.
	clearMFACookie(w)

	if err != nil {
		writeJSON(w, http.StatusOK, map[string]any{
			"ok":      true,
			"warning": "signed out locally, but the token could not be invalidated upstream",
		})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// handleVerifyEmail confirms an account using the verification token from the
// email-verification link, then redirects to the index page with a status query
// param. LoginRadius's verification email contains a link to whatever URL was
// passed as `verificationurl` on the register call (with `?vtoken=...`
// appended); LR_VERIFICATION_URL should point at this route.
//
// Wraps client.User.CheckEmailAvailability — the spec overloads that operation:
// with a verificationtoken query param it verifies an account, without one it
// just checks availability. The demo only uses the verification form.
func (s *server) handleVerifyEmail(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	token := firstNonEmpty(q.Get("vtoken"), q.Get("verificationToken"), q.Get("verificationtoken"))
	if token == "" {
		http.Redirect(w, r, "/?verify=missing", http.StatusFound)
		return
	}

	_, _, err := s.client.User.
		CheckEmailAvailability(r.Context()).
		Verificationtoken(token).
		Execute()
	if err != nil {
		msg := "verification failed"
		if lrErr, ok := loginradius.AsError(err); ok && lrErr != nil {
			switch {
			case lrErr.Description != "":
				msg = lrErr.Description
			case lrErr.Message != "":
				msg = lrErr.Message
			}
		}
		http.Redirect(w, r, "/?verify=error&message="+url.QueryEscape(msg), http.StatusFound)
		return
	}
	http.Redirect(w, r, "/?verify=success", http.StatusFound)
}

// handleDeleteAccount deletes the signed-in user's own account.
//
// Accounts.DeleteAccountByEmail is an ADMIN operation: it authenticates with
// the API secret and will delete any account in the tenant by email address.
// Exposing that straight through would let anyone with a demo session delete
// anyone else, so this handler first reads the signed-in profile and refuses
// unless the address matches one the session actually owns. That guard is demo
// policy, not an SDK limitation — a real integration would authorise this in
// whatever way its own model demands.
func (s *server) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email string `json:"email"`
	}
	if err := decodeJSON(r, &in); err != nil {
		badRequest(w, "expected a JSON object body")
		return
	}
	if in.Email == "" {
		missingField(w, "email")
		return
	}

	token := tokenFrom(r.Context())
	profile, _, err := s.client.User.GetAccountDetails(r.Context()).AccessToken(token).Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}

	if !ownsEmail(profile.Email, in.Email) {
		writeJSON(w, http.StatusForbidden, map[string]any{
			"error": "refusing to delete an account you are not signed in as",
			"hint":  "the demo only deletes the signed-in account; the underlying API would delete any address",
		})
		return
	}

	resp, _, err := s.client.Accounts.
		DeleteAccountByEmail(r.Context()).
		Email(in.Email).
		Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}

	// The account is gone; the session that pointed at it must go too.
	if id := getSessionCookie(r); id != "" {
		s.sessions.Delete(id)
	}
	clearSessionCookie(w)
	clearMFACookie(w)
	writeJSON(w, http.StatusOK, map[string]any{"deleted": resp})
}

// ownsEmail reports whether addr is one of the addresses on the profile,
// compared case-insensitively as email addresses are.
func ownsEmail(emails []loginradius.ProfileEmailInner, addr string) bool {
	for _, e := range emails {
		if e.Value != nil && strings.EqualFold(strings.TrimSpace(*e.Value), strings.TrimSpace(addr)) {
			return true
		}
	}
	return false
}

// handleFinishPasskeyRegistration completes WebAuthn enrolment with the
// attestation the browser produced, then signs the new user in.
func (s *server) handleFinishPasskeyRegistration(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Credential map[string]any `json:"credential"`
		Email      string         `json:"email"`
	}
	if err := decodeJSON(r, &in); err != nil {
		badRequest(w, "expected a JSON object body")
		return
	}
	if len(in.Credential) == 0 {
		missingField(w, "credential")
		return
	}
	// The API requires an address here: enrolment creates the account, so
	// there is no existing profile to take one from.
	if in.Email == "" {
		missingField(w, "email")
		return
	}

	var cred loginradius.PasskeyCredentialCreationResponse
	if err := remapJSON(in.Credential, &cred); err != nil {
		badRequest(w, "credential is not a WebAuthn attestation response")
		return
	}

	// Note the shape difference between the two passkey finish models: this one
	// takes the profile's array of {Type, Value} addresses, exactly as ordinary
	// registration does, while PasskeyLoginFinish takes a plain string. Sending
	// a bare string here silently leaves Email unset and the API replies that
	// email is required.
	emailType := "Primary"
	body := loginradius.PasskeyRegisterFinish{
		PasskeyCredential: &cred,
		Email: []loginradius.ProfileRequestModelEmailInner{
			{Type: &emailType, Value: &in.Email},
		},
	}
	resp, _, err := s.client.Registration.
		FinishPasskeyRegistration(r.Context()).
		PasskeyRegisterFinish(body).
		Verificationurl(verificationURL()).
		Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}

	signedIn := false
	if access, refresh := registrationTokens(resp); access != "" {
		id, cerr := s.sessions.Create(access, refresh)
		if cerr != nil {
			http.Error(w, "could not mint session", http.StatusInternalServerError)
			return
		}
		setSessionCookie(w, id)
		signedIn = true
	}
	writeJSON(w, http.StatusOK, map[string]any{"signed_in": signedIn, "profile": resp})
}

// handleFinishPasskeyLogin completes WebAuthn login with the assertion the
// browser produced and mints a session from the returned token.
func (s *server) handleFinishPasskeyLogin(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Credential map[string]any `json:"credential"`
		Email      string         `json:"email"`
	}
	if err := decodeJSON(r, &in); err != nil {
		badRequest(w, "expected a JSON object body")
		return
	}
	if len(in.Credential) == 0 {
		missingField(w, "credential")
		return
	}

	var cred loginradius.PasskeyCredentialAssertionResponse
	if err := remapJSON(in.Credential, &cred); err != nil {
		badRequest(w, "credential is not a WebAuthn assertion response")
		return
	}

	body := loginradius.PasskeyLoginFinish{
		PasskeyCredential: &cred,
		Email:             optStr(in.Email),
	}
	resp, _, err := s.client.Login.
		FinishPasskeyLogin(r.Context()).
		PasskeyLoginFinish(body).
		Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}
	s.finishWithSession(w, resp)
}

// handleMfaVerifyEmailOtp completes an in-progress MFA login challenge with the
// OTP delivered by email.
func (s *server) handleMfaVerifyEmailOtp(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email string `json:"email"`
		Otp   string `json:"otp"`
	}
	if err := decodeJSON(r, &in); err != nil {
		badRequest(w, "expected a JSON object body")
		return
	}
	if in.Email == "" {
		missingField(w, "email")
		return
	}
	if in.Otp == "" {
		missingField(w, "otp")
		return
	}

	body := loginradius.ReAuthModelByEmailOtp{Emailid: in.Email, Otp: in.Otp}
	resp, _, err := s.client.Security.
		ValidateMfaOTPByEmail(r.Context()).
		Secondfactorauthenticationtoken(tokenFrom(r.Context())).
		ReAuthModelByEmailOtp(body).
		Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}
	clearMFACookie(w)
	s.finishWithSession(w, resp)
}

// handleMfaVerifyTotp completes an in-progress MFA login challenge with a code
// from the user's authenticator app.
func (s *server) handleMfaVerifyTotp(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Totp string `json:"totp"`
	}
	if err := decodeJSON(r, &in); err != nil {
		badRequest(w, "expected a JSON object body")
		return
	}
	if in.Totp == "" {
		missingField(w, "totp")
		return
	}

	resp, _, err := s.client.Security.
		VerifyTotpByMfaToken(r.Context()).
		Secondfactorauthenticationtoken(tokenFrom(r.Context())).
		AuthenticatorCodeRequest(totpBody(in.Totp)).
		Execute()
	if err != nil {
		writeSDKError(w, err)
		return
	}
	clearMFACookie(w)
	s.finishWithSession(w, resp)
}

// finishWithSession mints a demo session from an AuthResponse and clears the
// MFA challenge. Shared by every route that completes an authentication —
// passkey login and both MFA verifications — so they cannot drift apart in how
// they establish a session.
func (s *server) finishWithSession(w http.ResponseWriter, resp *loginradius.AuthResponse) {
	if resp == nil || resp.AccessToken == nil || *resp.AccessToken == "" {
		writeJSON(w, http.StatusBadGateway, map[string]any{
			"error": "authentication succeeded but no access_token returned",
		})
		return
	}
	id, err := s.sessions.Create(*resp.AccessToken, valueOr(resp.RefreshToken))
	if err != nil {
		http.Error(w, "could not mint session", http.StatusInternalServerError)
		return
	}
	setSessionCookie(w, id)
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "profile": resp.Profile})
}

// ---------------------------------------------------------------------------
// call<Name> — the SDK call for every generated route.
//
// routes_gen.go has already decoded the body, filled path and query values,
// checked required fields and resolved the token. What is left is the part a
// customer actually copies, which is why it stays here in hand-written form
// rather than being generated out of sight.
// ---------------------------------------------------------------------------

func (s *server) callPasswordlessLoginByEmail(ctx context.Context, in passwordlessLoginByEmailRequest) (any, error) {
	resp, _, err := s.client.Login.
		PasswordlessLoginByEmail(ctx).
		Email(in.Email).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callPasswordlessLoginByPhone(ctx context.Context, in passwordlessLoginByPhoneRequest) (any, error) {
	resp, _, err := s.client.Login.
		PasswordlessLoginByPhone(ctx).
		Phone(in.Phone).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callForgotPassword(ctx context.Context, in forgotPasswordRequest) (any, error) {
	body := loginradius.ForgotPasswordRequest{Email: &in.Email}
	resp, _, err := s.client.Password.
		ForgotPassword(ctx).
		ForgotPasswordRequest(body).
		Resetpasswordurl(resetPasswordURL()).
		Execute()
	if err != nil {
		return nil, err
	}
	return map[string]any{"is_posted": resp.GetIsPosted()}, nil
}

func (s *server) callResetPassword(ctx context.Context, in resetPasswordRequest) (any, error) {
	body := loginradius.ResetPassword{
		ResetPasswordOneOf: &loginradius.ResetPasswordOneOf{
			ResetToken: in.ResetToken,
			Password:   in.Password,
		},
	}
	resp, _, err := s.client.Password.
		ResetPasswordByResetToken(ctx).
		ResetPassword(body).
		Execute()
	if err != nil {
		return nil, err
	}
	return map[string]any{"is_posted": resp.GetIsPosted()}, nil
}

func (s *server) callChangePassword(ctx context.Context, in changePasswordRequest, token string) (any, error) {
	body := loginradius.ChangePassword{
		OldPassword: in.OldPassword,
		NewPassword: in.NewPassword,
	}
	resp, _, err := s.client.Password.
		ChangePassword(ctx).
		AccessToken(token).
		ChangePassword(body).
		Execute()
	if err != nil {
		return nil, err
	}
	return map[string]any{"is_posted": resp.GetIsPosted()}, nil
}

func (s *server) callGetProfile(ctx context.Context, _ getProfileRequest, token string) (any, error) {
	profile, _, err := s.client.User.
		GetAccountDetails(ctx).
		AccessToken(token).
		Execute()
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *server) callUpdateProfile(ctx context.Context, in updateProfileRequest, token string) (any, error) {
	body := loginradius.UpdateAccountByAccessTokenRequest{
		FirstName: optStr(in.FirstName),
		LastName:  optStr(in.LastName),
		About:     optStr(in.About),
	}
	resp, _, err := s.client.User.
		UpdateAccountByAccessToken(ctx).
		AccessToken(token).
		UpdateAccountByAccessTokenRequest(body).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callAddEmail(ctx context.Context, in addEmailRequest, token string) (any, error) {
	body := loginradius.AddEmailModel{Email: in.Email, Type: optStr(in.Type)}
	resp, _, err := s.client.User.
		AddEmail(ctx).
		AccessToken(token).
		AddEmailModel(body).
		Verificationurl(verificationURL()).
		Execute()
	if err != nil {
		return nil, err
	}
	return map[string]any{"is_posted": resp.GetIsPosted()}, nil
}

func (s *server) callDeleteEmail(ctx context.Context, in deleteEmailRequest, token string) (any, error) {
	uc, err := s.userClient(token)
	if err != nil {
		return nil, err
	}
	body := loginradius.DeleteemailbyaccesstokenRequest{Email: in.Email}
	resp, _, err := uc.User.
		Deleteemailbyaccesstoken(ctx).
		DeleteemailbyaccesstokenRequest(body).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callUpdatePhone(ctx context.Context, in updatePhoneRequest, token string) (any, error) {
	body := loginradius.PhoneIdModel{Phone: in.Phone}
	resp, _, err := s.client.User.
		ChangePhoneNumber(ctx).
		AccessToken(token).
		PhoneIdModel(body).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callRequestResetOtp(ctx context.Context, in requestResetOtpRequest) (any, error) {
	body := loginradius.ForgotPasswordPhoneModel{Phone: in.Phone}
	resp, _, err := s.client.Password.
		RequestOTPForPasswordReset(ctx).
		ForgotPasswordPhoneModel(body).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callResetPasswordWithOtp(ctx context.Context, in resetPasswordWithOtpRequest) (any, error) {
	body := loginradius.ResetPasswordWithOTP{
		Otp:      in.Otp,
		Phone:    in.Phone,
		Password: in.Password,
	}
	resp, _, err := s.client.Password.
		ResetPasswordWithOTP(ctx).
		ResetPasswordWithOTP(body).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callCreateCustomObject(ctx context.Context, in createCustomObjectRequest, token string) (any, error) {
	resp, _, err := s.client.CustomObject.
		CreateCustomObjectByToken(ctx).
		AccessToken(token).
		Objectname(in.Objectname).
		RequestBody(in.Data).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callListCustomObjects(ctx context.Context, in listCustomObjectsRequest, token string) (any, error) {
	resp, _, err := s.client.CustomObject.
		GetCustomObjectByToken(ctx).
		AccessToken(token).
		Objectname(in.Objectname).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callUpdateCustomObject(ctx context.Context, in updateCustomObjectRequest, token string) (any, error) {
	// The record id is a path parameter of the operation itself, so it is a
	// positional argument rather than a builder call.
	resp, _, err := s.client.CustomObject.
		UpdateCustomObjectByTokenAndRecordId(ctx, in.ObjectRecordID).
		AccessToken(token).
		Objectname(in.Objectname).
		RequestBody(in.Data).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callDeleteCustomObject(ctx context.Context, in deleteCustomObjectRequest, token string) (any, error) {
	resp, _, err := s.client.CustomObject.
		DeleteCustomObjectByTokenAndRecordId(ctx, in.ObjectRecordID).
		AccessToken(token).
		Objectname(in.Objectname).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// callRefreshToken exchanges the session's refresh token for a new access
// token and rotates what the demo session holds.
//
// This uses the /manage/ refresh operation deliberately. The native variant
// takes the access token as a QUERY parameter, which would put a bearer
// credential into access logs, proxy logs and Referer headers.
func (s *server) callRefreshToken(ctx context.Context, _ refreshTokenRequest, _ string) (any, error) {
	refresh := refreshTokenFrom(ctx)
	if refresh == "" {
		return nil, errNoRefreshToken
	}
	resp, _, err := s.client.AccountSession.
		RefreshAccessToken(ctx).
		RefreshToken(refresh).
		Execute()
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.AccessToken == nil || *resp.AccessToken == "" {
		return nil, errRefreshReturnedNoToken
	}

	// Write the rotated tokens back to the session. Refreshing INVALIDATES the
	// previous access token upstream, so a session left holding the old one is
	// not merely stale — every subsequent authenticated call fails. The refresh
	// token is rotated too when the tenant returns a new one.
	s.sessions.Replace(sessionIDFrom(ctx), *resp.AccessToken, valueOr(resp.RefreshToken))

	return map[string]any{
		"refreshed":  true,
		"rotated":    resp.RefreshToken != nil && *resp.RefreshToken != "",
		"expires_in": resp.ExpiresIn,
	}, nil
}

func (s *server) callValidateToken(ctx context.Context, _ validateTokenRequest, token string) (any, error) {
	resp, _, err := s.client.Session.
		AuthValidateAccessToken(ctx).
		AccessToken(token).
		Execute()
	if err != nil {
		return nil, err
	}
	return map[string]any{"valid": true, "expires_in": resp.ExpiresIn}, nil
}

func (s *server) callActiveSession(ctx context.Context, _ activeSessionRequest, token string) (any, error) {
	resp, _, err := s.client.AccountSession.
		GetActiveSession(ctx).
		Token(token).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callBeginPasskeyRegistration(ctx context.Context, in beginPasskeyRegistrationRequest) (any, error) {
	resp, _, err := s.client.Registration.
		BeginPasskeyRegistration(ctx).
		Identifier(in.Identifier).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callBeginPasskeyLogin(ctx context.Context, in beginPasskeyLoginRequest) (any, error) {
	resp, _, err := s.client.Login.
		BeginPasskeyLogin(ctx).
		Identifier(in.Identifier).
		Verificationurl(verificationURL()).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callMfaSettings(ctx context.Context, in mfaSettingsRequest, token string) (any, error) {
	req := s.client.Security.
		GetMFASettings(ctx).
		AccessToken(token)
	// Only sent when supplied. Duo returns the user here after its challenge;
	// passing an empty string would put `duoredirecturi=` on the wire for every
	// tenant, including those with no Duo configured.
	if in.DuoRedirectUri != "" {
		req = req.Duoredirecturi(in.DuoRedirectUri)
	}
	resp, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callMfaEnrolTotp(ctx context.Context, in mfaEnrolTotpRequest, token string) (any, error) {
	resp, _, err := s.client.Security.
		Verify2faTOTPAuth(ctx).
		AccessToken(token).
		AuthenticatorCodeRequest(totpBody(in.Totp)).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callMfaBackupCodes(ctx context.Context, _ mfaBackupCodesRequest, token string) (any, error) {
	resp, _, err := s.client.Security.
		MfaGenerateBackupCodes(ctx).
		AccessToken(token).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *server) callMfaSendEmailOtp(ctx context.Context, in mfaSendEmailOtpRequest, token string) (any, error) {
	body := loginradius.EmailModel{Email: in.Email}
	resp, _, err := s.client.Security.
		ResendEmailOTPMFAToken(ctx).
		Secondfactorauthenticationtoken(token).
		EmailModel(body).
		Execute()
	if err != nil {
		return nil, err
	}
	return map[string]any{"is_posted": resp.GetIsPosted()}, nil
}

// ---------------------------------------------------------------------------
// Small shared helpers.
// ---------------------------------------------------------------------------

// errNoRefreshToken is returned when the refresh route is called on a session
// whose tenant never issued a refresh token. Reported as a plain 502 by
// writeSDKError, which is accurate: nothing is wrong with the request.
var errNoRefreshToken = &demoError{msg: "this session has no refresh token; the tenant did not return one at login"}

// errRefreshReturnedNoToken guards against replacing a live session with an
// empty token, which would sign the user out on a successful-looking call.
var errRefreshReturnedNoToken = &demoError{msg: "refresh succeeded but returned no access_token"}

type demoError struct{ msg string }

func (e *demoError) Error() string { return e.msg }

// firstNonEmpty returns the first non-empty value among the supplied strings.
func firstNonEmpty(vs ...string) string {
	for _, v := range vs {
		if v != "" {
			return v
		}
	}
	return ""
}

// registrationTokens pulls the access and refresh tokens out of a registration
// response. They sit under Data rather than at the top level, and a tenant that
// requires email verification returns none at all — so both registration paths
// (password and passkey) go through here rather than each walking the nesting
// themselves.
func registrationTokens(resp *loginradius.RegistrationResponse) (access, refresh string) {
	if resp == nil || resp.Data == nil {
		return "", ""
	}
	return valueOr(resp.Data.AccessToken), valueOr(resp.Data.RefreshToken)
}

func valueOr(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func boolOr(p *bool) bool { return p != nil && *p }

// writeJSON marshals body as JSON with the given status. Errors during
// marshaling fall through to a 500 — the demo never serialises anything
// exotic so this is a belt-and-braces handler.
func writeJSON(w http.ResponseWriter, status int, body any) {
	buf, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "demo: marshal failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write(buf)
}

// writeSDKError unwraps *loginradius.Error and surfaces its structured
// fields to the browser. Transport-level errors fall back to a 502.
//
// raw_body is included so the demo UI can show whatever the upstream
// actually returned — useful when extractEnvelope's parsers don't pick
// anything up (empty bodies, HTML error pages, gateway responses).
//
// SDK methods today return *openapi.GenericOpenAPIError directly, not the
// facade's *loginradius.Error — loginradius.AsError() converts at the
// call site. Customers should follow the same pattern.
func writeSDKError(w http.ResponseWriter, err error) {
	lrErr, ok := loginradius.AsError(err)
	if !ok {
		// Transport failure or unknown wrapper.
		writeJSON(w, http.StatusBadGateway, map[string]any{"error": err.Error()})
		return
	}
	status := lrErr.StatusCode
	if status < 400 || status >= 600 {
		status = http.StatusBadGateway
	}
	writeJSON(w, status, map[string]any{
		"error":       lrErr.Message,
		"code":        lrErr.Code,
		"description": lrErr.Description,
		"status_code": lrErr.StatusCode,
		"raw_body":    string(lrErr.RawBody),
	})
}
