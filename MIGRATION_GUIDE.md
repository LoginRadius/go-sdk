# Migrating from v11 to v12

This guide walks you through upgrading from `github.com/LoginRadius/go-sdk`
(v11) to `github.com/LoginRadius/go-sdk/v12`. v11 stays supported on its
current tag line — there is no forced cut-over.

## Why migrate

- **Full spec coverage.** v11 hand-wraps ~90 endpoints across 13 packages.
  v12 is generated from the canonical OpenAPI 3.0.1 spec — every operation
  the LoginRadius platform exposes is reachable through a typed service.
- **Typed request and response models.** v11 hands you `*httprutils.Response`
  with a `Body string` field; you unmarshal yourself. v12 returns typed Go
  structs for ~1,500 schemas — IDE autocomplete works, schema mistakes are
  compile errors.
- **Centralised authentication.** v11 set credentials per-package; v12
  configures them once at `NewClient` and a single `http.RoundTripper`
  injects them on every request.
- **Regenerable.** v12 is generated from the OpenAPI spec, so a new spec
  version is a regeneration rather than hand-edited code. v11 changes were
  coded line-by-line.
- **Idiomatic Go.** Functional options, typed errors, `context.Context`
  threading on every call, and a clean `Client` struct rather than a global.

## Decide whether to migrate

Stay on **v11** if any of the following apply today:

- You depend on the exact `*httprutils.Response`/`Body string` shape from
  v11 — e.g. you persist raw response bodies and don't want to re-shape
  your storage.
- You ship a stable downstream SDK that hand-wraps v11 — wait until v12
  reaches a tagged GA before re-pinning.

For all other use cases, v12 is the recommended path forward. v11 continues
to receive security fixes; new endpoints land on v12 only.

## What's different at a glance

| Concern | v11 | v12 |
|---|---|---|
| Import path | `github.com/LoginRadius/go-sdk` | `github.com/LoginRadius/go-sdk/v12` |
| Construction | `lr.NewLoginradius(cfg, opts)` | `loginradius.NewClient(WithAPIKey(...), ...)` |
| Per-package handle | `account.Loginradius{...}` (re-wrapped each call) | `client.Accounts` on the `*Client` |
| Request bodies | `body interface{}` + `queries ...interface{}` | Typed builders + typed structs |
| Response | `*httprutils.Response` + `Body string` | Typed struct + `*http.Response` |
| Errors | `lrerror.Error` interface | `*loginradius.Error` (concrete type) |
| Context | implicit | explicit `context.Context` on every call |
| Auth header injection | per-package | single `http.RoundTripper`, configured once |

## Module path change

v12 is published at the repository root. The `/v12` in the import path is
Go's major-version suffix, not a directory — module `github.com/LoginRadius/go-sdk/v12`
with `go.mod` at the root is exactly how a v2+ Go module is released.

```go
// v11
import "github.com/LoginRadius/go-sdk/api/authentication"

// v12
import loginradius "github.com/LoginRadius/go-sdk/v12"
```

Both modules can live in the same `go.mod` during a phased rollout if you
need to migrate one call site at a time. `go get` them independently:

```bash
# 12.0.0-rc.1 is a prerelease. `go get` never selects one via @latest,
# so pin the tag explicitly; @latest works once a GA version is tagged.
go get github.com/LoginRadius/go-sdk/v12@v12.0.0-rc.1
```

## Construction diff

```go
// v11
cfg := lr.Config{
    APIKey:    os.Getenv("LR_API_KEY"),
    APISecret: os.Getenv("LR_API_SECRET"),
    Domain:    "acme",
}
client, err := lr.NewLoginradius(&cfg, map[string]string{"token": userAT})
authAPI := authentication.Loginradius{Client: client}

// v12
client, err := loginradius.NewClient(
    loginradius.WithAPIKey(os.Getenv("LR_API_KEY")),
    loginradius.WithAPISecret(os.Getenv("LR_API_SECRET")),
    loginradius.WithDomain("acme"),
    loginradius.WithAccessToken(userAT),
)
// No per-package wrapper — call directly via client.Authentication, client.Accounts, etc.
```

`WithDomain("acme")` resolves to `https://acme.hub.loginradius.com` — the
domain is your **tenant name** from the LoginRadius admin console, not an
arbitrary host.

## Auth posture diff

v12 sends the `X-LoginRadius-ApiKey` / `X-LoginRadius-ApiSecret` headers by
preference — this keeps credentials out of access logs, browser history, and
proxy URL caches — **and** populates the `apikey=` / `apisecret=` query
parameters as a compatibility fallback, because a large share of endpoints in
the spec accept only the query-param scheme.

> If a gateway or proxy in front of the API keys off the credential's location,
> check it against both: v12 sends the headers and the query parameters on the
> same request.

v12 supports every scheme the spec defines:

| v12 option | Effect |
|---|---|
| `WithAPIKey` | `X-LoginRadius-ApiKey` header + `apikey=` query fallback |
| `WithAPISecret` | `X-LoginRadius-ApiSecret` header + `apisecret=` query fallback |
| `WithXLoginRadiusAPIKey` | Header override (rare — different header value than query) |
| `WithXLoginRadiusAPISecret` | Header override (rare — same) |
| `WithClientID` / `WithClientSecret` | `client_id=` / `client_secret=` query |
| `WithAccessToken` | `access_token=` query (user-context operations) |
| `WithBearerToken` | `Authorization: Bearer …` |
| `WithM2MBearerToken` | `Authorization: Bearer …` (machine-to-machine JWT) |

> **Note:** `apiRequestSigning` (`Digest` + `X-Request-Expires`) is available
> via `WithAPIRequestSigning(true)`. It is opt-in, off by default, and applies
> only to `/manage/` paths (excluding `/account/access_token`). Validate it
> against your own tenant before relying on it.

### Cross-cutting request options

These are client-wide settings applied to every outgoing request, replacing the
per-call plumbing the legacy SDK required.

| v11 config | v12 option | Effect |
|---|---|---|
| `originIp` | `WithOriginIP` | `X-Origin-IP` header. Forwards the end user's IP for risk-based auth and audit trails. Note this header is **not** in the OpenAPI spec — the facade sends it to match v11. |
| `serverRegion` | `WithServerRegion` | `region=` query parameter. `domain`/`customDomain` select a *tenant*, not a region. |
| `fieldsParam` / `fieldsValue` | `WithFields` | `fields=` query parameter on every request. v11 made you supply the separator (`?fields=` vs `&fields=`); pass only the value now. |
| `preventWebhook` | `WithPreventWebhook` | `X-PreventWebhook: true` on every request. v12 also has a per-operation parameter. |
| *(none)* | `WithDefaultHeaders` | Headers merged into every request at the **lowest** precedence — the SDK's own credential and User-Agent headers always win. |
| *(none)* | `WithDebug` | Request/response logging with credentials redacted. |
| *(none)* | `WithServerIndex` | Escape hatch for the 42 operations the spec pins to their own host. An explicit `WithBaseURL` already overrides those pins. |

## Request and response shape diff

```go
// v11 — register a new user
body := map[string]interface{}{
    "email": []map[string]string{
        {"type": "primary", "value": "alice@example.com"},
    },
    "password": "<...>",
}
resp, err := authAPI.PostAuthRegister(body, map[string]string{"verificationurl": "..."})
if err != nil { return err }
var profile UserProfile
if err := json.Unmarshal([]byte(resp.Body), &profile); err != nil { return err }

// v12 — same flow, typed
profile, _, err := client.Registration.UserRegistration(ctx).
    SOTT(sott).
    UserProfile(loginradius.UserProfile{
        Email: []loginradius.UserProfileEmail{
            {Type: "primary", Value: "alice@example.com"},
        },
        Password: "<...>",
    }).
    Verificationurl("...").
    Execute()
if err != nil { return err }
// profile is a *loginradius.UserProfilePostResponse with typed fields.
```

The chained builder is generated by openapi-generator. Customers never
import the internal `openapi` package — all request and response types are
re-exported at the top level of `loginradius`.

## Error handling diff

```go
// v11
resp, err := authAPI.PostAuthLoginByEmail(body)
if err != nil {
    if lrErr, ok := err.(lrerror.Error); ok {
        log.Printf("%s: %s", lrErr.Code(), lrErr.Message())
    }
    return err
}

// v12
_, _, err := client.Authentication.LoginByEmail(ctx).LoginRequest(req).Execute()
if err != nil {
    var lrErr *loginradius.Error
    if errors.As(err, &lrErr) {
        switch {
        case lrErr.IsAuth():      // 401 — re-authenticate
        case lrErr.IsForbidden(): // 403 — permissions / plan limit, NOT a credential issue
        case lrErr.IsRateLimit(): // 429 — back off
        case lrErr.IsServer():    // 5xx — retry
        }
        log.Printf("status=%d code=%s msg=%s", lrErr.StatusCode, lrErr.Code, lrErr.Message)
    }
    return err
}
```

v12's `*loginradius.Error` parses the standard LoginRadius envelope
(`{ErrorCode, Message, Description}`) **and** the `/api/v2/access_token/*`
lowercase variant (`{errorCode, message, description}`) **and** OAuth-style
envelopes (`{error, error_description}`) — case-insensitive, one extractor
for all of them.

Notably, **`IsAuth()` returns true only for 401**. 403 is a separate
`IsForbidden()` check because in LoginRadius a 403 typically means IP/domain
access restrictions or a plan-level feature gate, not a credential failure —
re-prompting the user for credentials would be wrong.

## Server selection

```go
// Production (default)
loginradius.WithBaseURL("https://api.loginradius.com")     // explicit, or omit for default

// Multi-tenant hosted-page
loginradius.WithDomain("acme")                              // → https://acme.hub.loginradius.com

// Custom domain
loginradius.WithCustomDomain("auth.acme.com")               // → https://auth.acme.com

// Staging / proxy
loginradius.WithBaseURL("https://staging.api.example")      // exact override
```

Precedence: `WithBaseURL` > `WithCustomDomain` > `WithDomain` > default.

## Custom HTTP client

```go
httpClient := &http.Client{Timeout: 10 * time.Second}
client, _ := loginradius.NewClient(
    loginradius.WithAPIKey(key),
    loginradius.WithHTTPClient(httpClient),
)
```

The SDK installs its auth `RoundTripper` on top of your client's
`Transport`, preserving any custom dialer / proxy / TLS config you set up.

## Endpoint mapping (representative)

The full v11 surface (~90 methods) maps to v12 services as follows. Use
`pkg.go.dev/github.com/LoginRadius/go-sdk/v12` for the exhaustive method
list per service.

| v11 method | v12 service.method |
|---|---|
| `account.PostManageAccountCreate` | `client.Accounts.CreateAccount` |
| `account.PostManageForgotPasswordToken` | `client.Accounts.ForgotPasswordToken` |
| `account.PostManageEmailVerificationToken` | `client.Accounts.EmailVerificationToken` |
| `account.GetManageAccountProfilesByEmail` | `client.Accounts.GetAccountProfilesByEmail` |
| `authentication.PostAuthAddEmail` | `client.Login.AddEmail` (or `client.User` depending on context) |
| `authentication.DeleteAuthRemoveEmail` | `client.User.RemoveEmail` |
| `authentication.DeleteAuthUnlinkSocialIdentities` | `client.SocialProviders.UnlinkSocialIdentities` |
| `configuration.GetConfiguration` | `client.CaptchaConfiguration.GetConfiguration` |
| `configuration.GetServerTime` | `client.Insights.GetServerTime` |
| `configuration.GetGenerateSottAPI` | `client.SOTT.GenerateSOTT` |
| `configuration.GetActiveSessionDetails` | `client.AccountSession.GetActiveSessionDetails` |
| `customobject.PostCustomObjectCreateByUID` | `client.CustomObject.CreateCustomObjectByUID` |
| `customobject.PostCustomObjectCreateByToken` | `client.CustomObject.CreateCustomObjectByToken` |
| `customobject.GetCustomObjectByObjectRecordIDAndUID` | `client.CustomObject.GetCustomObjectByRecordIDAndUID` |
| `mfa.PostMFAEmailLogin` | `client.SecondFactorConfiguration.MFAEmailLogin` |
| `mfa.PostMFAUsernameLogin` | `client.SecondFactorConfiguration.MFAUsernameLogin` |
| `mfa.PostMFAPhoneLogin` | `client.SecondFactorConfiguration.MFAPhoneLogin` |
| `mfa.GetMFAValidateAccessToken` | `client.SecondFactorConfiguration.MFAValidateAccessToken` |
| `onetouchlogin.PostOneTouchLoginByEmail` | `client.Login.OneTouchLoginByEmail` |
| `onetouchlogin.PostOneTouchLoginByPhone` | `client.Login.OneTouchLoginByPhone` |
| `phoneauthentication.PostPhoneLogin` | `client.Login.PhoneLogin` |
| `phoneauthentication.PostPhoneForgotPasswordByOTP` | `client.Password.PhoneForgotPasswordByOTP` |
| `role.PostRolesCreate` | `client.Roles.CreateRoles` |
| `role.DeleteAccountRole` | `client.Roles.DeleteAccountRole` |
| `role.GetContextRolesPermissions` | `client.Roles.GetContextRolesPermissions` |
| `role.GetRolesList` | `client.Roles.GetRolesList` |
| `smartlogin.GetSmartLoginByEmail` | `client.Login.GetSmartLoginByEmail` |
| `smartlogin.GetSmartLoginPing` | `client.Login.PingSmartLogin` |
| `social.GetSocialAccessToken` | `client.SocialProviders.GetAccessToken` |
| `tokenmanagement.GetAccessTokenViaFacebook` | `client.AccountSession.GetAccessTokenViaFacebook` |
| `tokenmanagement.GetRefreshUserProfile` | `client.Session.GetRefreshedAccessToken` |
| `webhook.PostWebhookSubscribe` | `client.Webhooks.SubscribeWebhook` |
| `webhook.GetWebhookSubscribedURLs` | `client.Webhooks.GetSubscribedWebhooks` |
| `webhook.DeleteWebhookUnsubscribe` | `client.Webhooks.UnsubscribeWebhook` |

Names follow a consistent pattern: drop the HTTP verb prefix (`Post`/`Get`/`Delete`)
and the package prefix (`MFA`/`Auth`); the rest matches the spec's `operationId`.

## Common-flow snippets

### Login by email

```go
// v11
resp, err := authAPI.PostAuthLoginByEmail(map[string]interface{}{
    "email":    "alice@example.com",
    "password": "<...>",
}, map[string]string{"loginurl": "..."})

// v12
profile, _, err := client.Login.PasswordlessLoginByEmail(ctx).
    Email("alice@example.com").
    Execute()
// Or for password-based flows, see client.Login.{LoginBy...} variants.
```

### Forgot password

```go
// v11
resp, err := authAPI.PostAuthForgotPassword(map[string]interface{}{"email": "alice@example.com"},
    map[string]string{"resetpasswordurl": "..."})

// v12
posted, _, err := client.Password.ForgotPassword(ctx).
    EmailOrUsername("alice@example.com").
    Resetpasswordurl("...").
    Execute()
```

### Get user profile

```go
// v11
client.Context.Token = userAccessToken
resp, err := authAPI.GetAuthReadAllProfileByToken()

// v12
client, _ := loginradius.NewClient(
    loginradius.WithAPIKey(key),
    loginradius.WithAccessToken(userAccessToken),
)
profile, _, err := client.User.GetAccountDetails(ctx).Execute()
```

### MFA enroll

```go
// v11
resp, err := mfaAPI.PutMFAUpdateAuthenticatorByAccessToken(body, queries)

// v12
result, _, err := client.SecondFactorConfiguration.
    MFAUpdateAuthenticatorByAccessToken(ctx).
    AuthenticatorVerificationModel(req).
    Execute()
```

## What we don't migrate for you

- **Persistent storage of response bodies.** If you serialise v11's
  `Body string` JSON to a database, those rows do not translate to v12's
  typed structs automatically. Treat the schema upgrade as a separate
  migration if it matters to your data layer.
- **Per-package wrappers in your own code.** If you wrap each `lrclient`
  in your own facade today, you'll want to redesign that around `*Client`
  rather than translate it line-for-line.
- **Inline base URLs.** v11 sometimes hard-codes `https://api.loginradius.com`
  into call sites. In v12, configure the server once via `WithBaseURL` /
  `WithDomain` / `WithCustomDomain`.

## Need help?

- Examples covering each auth scheme: [`v12/examples/`](./examples/README.md)
- Codegen and contribution flow: [`CONTRIBUTING.md`](./CONTRIBUTING.md)
- File an issue: <https://github.com/LoginRadius/go-sdk/issues>
