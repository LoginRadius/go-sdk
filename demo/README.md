# v12 Go SDK demo

A small `net/http` service that exercises the LoginRadius CIAM surface through
the v12 Go SDK: registration and login, profile and identifier management,
password reset by token or OTP, custom objects, the access-token lifecycle,
MFA enrolment and challenge, and passkey (WebAuthn). A single static HTML page
with tabs drives every flow from a browser. No external Go dependencies; no
JavaScript build step.

> **Demo only.** Cookies are unsigned, there is no CSRF protection, and the
> session map lives in process memory. Do **not** copy this into production.
> The goal is to show how to wire the v12 SDK into a Go HTTP server, not how to
> build a secure CIAM app.

## Run

```bash
cd demo
export LR_API_KEY=<your_api_key>
export LR_API_SECRET=<your_api_secret>       # optional, but most flows need it
export LR_DOMAIN=<your_tenant_name>          # optional
go run .
```

Browse <http://localhost:8080/> and walk through the tabs.

See [`.env.example`](./.env.example) for the full list of supported env vars.
On startup the demo loads `demo/.env`, falling back to `.env` at the module
root. Variables already exported in your shell win over the file. Never commit
a real `.env` — it is gitignored for a reason.

## Endpoints

The route table is **generated** from the shared SDK manifest, so every
LoginRadius SDK's demo exposes the same contract.
It is not hand-maintained here, and it cannot drift from the handlers: a route
with no handler fails to compile.

| Method | Path | SDK call |
|---|---|---|
| GET | `/` | (serves `static/index.html`) |
| POST | `/api/auth/register` | `Registration.UserRegistrationBySottEmailPhoneUserName` |
| POST | `/api/auth/login` | `Login.EmailByLoginUserNamePhone` |
| POST | `/api/auth/logout` | `Session.InvalidateAccessToken` |
| GET | `/api/auth/verify` | `User.CheckEmailAvailability` (with `Verificationtoken`) |
| POST | `/api/password/forgot` | `Password.ForgotPassword` |
| POST | `/api/password/reset` | `Password.ResetPasswordByResetToken` |
| POST | `/api/password/otp` | `Password.RequestOTPForPasswordReset` |
| PUT | `/api/password/otp` | `Password.ResetPasswordWithOTP` |
| POST | `/api/password/change` | `Password.ChangePassword` |
| GET | `/api/profile` | `User.GetAccountDetails` |
| POST | `/api/profile/update` | `User.UpdateAccountByAccessToken` |
| POST | `/api/email/add` | `User.AddEmail` |
| DELETE | `/api/email` | `User.Deleteemailbyaccesstoken` |
| PUT | `/api/phone` | `User.ChangePhoneNumber` |
| DELETE | `/api/account` | `Accounts.DeleteAccountByEmail` |
| POST | `/api/token/refresh` | `AccountSession.RefreshAccessToken` |
| GET | `/api/token/validate` | `Session.AuthValidateAccessToken` |
| GET | `/api/token/session` | `AccountSession.GetActiveSession` |
| GET | `/api/passkey/register/begin` | `Registration.BeginPasskeyRegistration` |
| POST | `/api/passkey/register/finish` | `Registration.FinishPasskeyRegistration` |
| GET | `/api/passkey/login/begin` | `Login.BeginPasskeyLogin` |
| POST | `/api/passkey/login/finish` | `Login.FinishPasskeyLogin` |
| GET | `/api/mfa/settings` | `Security.GetMFASettings` |
| PUT | `/api/mfa/totp` | `Security.Verify2faTOTPAuth` |
| GET | `/api/mfa/backupcodes` | `Security.MfaGenerateBackupCodes` |
| POST | `/api/mfa/login/email` | `Security.ResendEmailOTPMFAToken` |
| PUT | `/api/mfa/login/email` | `Security.ValidateMfaOTPByEmail` |
| PUT | `/api/mfa/login/totp` | `Security.VerifyTotpByMfaToken` |

### Two kinds of credential

Most authenticated routes read the demo session (`lr_session`) and pass its
access token to the SDK. The three `/api/mfa/login/*` routes do **not**: they
are authenticated by the second-factor token a challenge login returned, held
in a separate `lr_mfa` cookie. The two are deliberately different cookies, so a
half-authenticated user never holds anything the session middleware accepts.

### Routes that need extra configuration

- **Custom objects are currently disabled.** The four `/api/customobject`
  routes are commented out in the shared SDK manifest pending a tenant
  with a custom-object schema configured. The handlers and UI panel are
  commented out alongside them and carry a `CUSTOM-OBJECTS-DISABLED` marker;
  restore all three together.
- **Passkey** needs a secure context. `http://localhost` qualifies, so the demo
  works as shipped; over plain HTTP on any other host the browser refuses.
- **`DELETE /api/account`** deletes the signed-in account only. The underlying
  operation is admin-scoped and would delete *any* address in the tenant, so
  the handler reads the signed-in profile first and refuses a mismatch. That
  guard is demo policy, not an SDK limitation.

## Exercise with curl

```bash
# Register
curl -i -X POST http://localhost:8080/api/auth/register \
  -H 'Content-Type: application/json' \
  -d '{"email":"alice@example.com","password":"correct horse battery staple"}'

# Login (cookie jar captures lr_session)
curl -i -X POST http://localhost:8080/api/auth/login \
  -c /tmp/lr-demo-cookies.txt \
  -H 'Content-Type: application/json' \
  -d '{"email":"alice@example.com","password":"correct horse battery staple"}'

# Profile (uses the cookie)
curl -i http://localhost:8080/api/profile -b /tmp/lr-demo-cookies.txt

# Verify email — this is the exact URL the LoginRadius verification email
# links to; -L follows the 302 back to / and you'll see ?verify=success or
# ?verify=error&message=… in the final URL.
curl -i -L 'http://localhost:8080/api/auth/verify?vtoken=<token-from-email>'
```

## Files

- `main.go` — server bootstrap, env-var checks, route registration, and the
  `wrap` middleware that enforces one verb and the route's auth mode.
- `routes_gen.go` — **generated.** The route table, one request type per
  generated route, and the wrapper that decodes, validates, and reports errors.
- `handlers.go` — the hand-written half: `call<Name>` holds the SDK call for
  each generated route, and `handle<Name>` implements the routes with real
  logic (SOTT, sessions, redirects, the delete-account guard, WebAuthn).
- `session.go` — in-memory session store and cookie helpers.
- `static/index.html` — single-page UI (vanilla JS, no build step).
- `.env.example` — env vars the demo consumes.

### Why the split

The wrapper around every SDK call is identical — decode a body, check required
fields, resolve a token, map an error to a status. Generating it from the
manifest keeps four languages' demos from drifting apart on any of that. The
SDK call itself stays hand-written in `handlers.go`, because that is the part a
customer actually copies and it should be readable, not generated out of sight.

## SDK note — lenient `oneOf` decoder

The generated SDK models several responses (including the login one) as
`oneOf{…}` and the generator's default `UnmarshalJSON` uses
`json.Decoder.DisallowUnknownFields()` to make branches mutually exclusive.
Real LoginRadius responses carry tenant-specific extras (`Uid`, `IsActive`,
`IsDeleted`, …) that neither branch declares, so strict decoding rejects
the whole payload with `data failed to match schemas in oneOf(…)`.

The SDK generator patches `newStrictDecoder` so
`DisallowUnknownFields` is disabled. The re-marshal `{}` check that each branch
performs (drop the branch if the decoded struct round-trips to an empty object)
remains the discriminator — it's the actual signal of which branch matched. The
patch applies to all 70+ oneOf decoders the spec produces, not just the login
one.

## What's intentionally absent

If you copy this code, you must add:

- **Signed cookies.** Today the cookie is the raw session id, unsigned. An
  attacker who can read or guess it gets a session.
- **`Secure`, `SameSite`, `MaxAge` on the cookies.** A real cookie sets all
  three; the MFA cookie should also expire with the challenge itself.
- **CSRF protection.** All mutating endpoints accept JSON; in a real app
  you'd add a double-submit token or a `SameSite=Strict` cookie + Origin
  check.
- **Rate limiting**, particularly on login, the password routes, and the MFA
  verification routes — all of which are guessable-secret endpoints.
- **Automatic token refresh.** `POST /api/token/refresh` is wired up and
  rotates the stored tokens, but nothing calls it on expiry for you.
- **Persistent session storage.** The map is wiped on restart.
- **Real logging / observability.** The demo logs to stdout only.
