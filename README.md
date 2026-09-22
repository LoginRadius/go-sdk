# LoginRadius Go SDK — v12

[![CI](https://github.com/LoginRadius/go-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/LoginRadius/go-sdk/actions/workflows/ci.yml)

Official Go SDK for the [LoginRadius](https://www.loginradius.com) Customer Identity and Access Management (CIAM) platform.

> **v12 is a major version with a new module path.** v11 customers can stay on `github.com/LoginRadius/go-sdk` indefinitely — see [Migrating from v11](#migrating-from-v11) when ready.

## What changed in v12

v11 was hand-written. v12 is **generated from the canonical [LoginRadius OpenAPI spec](./LoginRadius-Public-APIs.yaml)** and wrapped with a thin facade for idiomatic Go ergonomics. Consequences:

- **Full coverage** of every endpoint in the spec (392 operations across 57 services) on day one.
- **Typed request and response models** for all 1,500+ schemas — no more dynamic JSON unmarshaling for the common path.
- **Centralized auth**: every supported credential goes through a single `http.RoundTripper`, configured once at client construction.
- **Regenerable**: the SDK ships in sync with the spec via a single Docker-based codegen step.

## Install

```bash
# 12.0.0-rc.1 is a prerelease. `go get` never selects one via @latest,
# so pin the tag explicitly; @latest works once a GA version is tagged.
go get github.com/LoginRadius/go-sdk/v12@v12.0.0-rc.1
```

Requires Go 1.24 or later.

## Quickstart

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    loginradius "github.com/LoginRadius/go-sdk/v12"
)

func main() {
    client, err := loginradius.NewClient(
        loginradius.WithAPIKey(os.Getenv("LR_API_KEY")),
        loginradius.WithAPISecret(os.Getenv("LR_API_SECRET")),
    )
    if err != nil {
        log.Fatal(err)
    }

    resp, _, err := client.Login.
        CheckUserNameAvailability(context.Background()).
        Username("alice").
        Execute()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("response: %+v\n", resp)
}
```

More examples: [`examples/README.md`](./examples/README.md) — 9 runnable programs covering each auth scheme (apikey, apisecret, access_token, bearer, M2M, client_id/secret, header-only) plus the basic quickstart, login, and custom-HTTP flows. Copy [`v12/.env.example`](./.env.example) to `v12/.env` before running.

For a full server-side integration walkthrough — registration and login, profile and identifier management, password reset by token or OTP, custom objects, the access-token lifecycle, MFA, and passkey (WebAuthn) — see [`demo/`](./demo/README.md). `cd demo && go run .` starts an HTTP service on `:8080` (the demo is its own module).

## Authentication

Credentials are supplied as options to `NewClient`. Set whichever ones the endpoints you call require — the SDK selects the right scheme per request automatically.

| Option | How it's sent | Used for |
|---|---|---|
| `WithAPIKey` | `X-LoginRadius-ApiKey` header (preferred) + `?apikey=` fallback | Most tenant-level endpoints |
| `WithAPISecret` | `X-LoginRadius-ApiSecret` header (preferred) + `?apisecret=` fallback | Server-side management endpoints |
| `WithXLoginRadiusAPIKey` | `X-LoginRadius-ApiKey` header (override) | Rare — when the header value must differ from the query value |
| `WithXLoginRadiusAPISecret` | `X-LoginRadius-ApiSecret` header (override) | Rare — see above |
| `WithClientID` / `WithClientSecret` | `?client_id=` / `?client_secret=` | OAuth-style endpoints |
| `WithAccessToken` | `?access_token=` | User-context operations |
| `WithBearerToken` | `Authorization: Bearer` | Generic bearer-secured endpoints |
| `WithM2MBearerToken` | `Authorization: Bearer` | Machine-to-machine endpoints |

The API key and secret are sent as **headers by preference** — this keeps them out of access logs, browser history, and proxy URL caches. The query-string fallback is still populated because ~56 endpoints in the LoginRadius spec accept only the query-param scheme; servers that support the header scheme will use it.

**Never** hard-code these in source — load them from environment variables, a secrets manager, or your platform's credential store. The SDK never logs credentials but your application must not either.

## Server selection

| Option | Resolved base URL |
|---|---|
| (none) | `https://api.loginradius.com` (production) |
| `WithDomain("acme")` | `https://acme.hub.loginradius.com` |
| `WithCustomDomain("auth.acme.com")` | `https://auth.acme.com` |
| `WithBaseURL("https://staging.api.example")` | exact override (staging, proxies) |

Precedence: `WithBaseURL` > `WithCustomDomain` > `WithDomain` > default.

## Custom HTTP client

```go
httpClient := &http.Client{
    Timeout: 10 * time.Second,
    Transport: &http.Transport{ /* proxy, TLS, pooling */ },
}
client, _ := loginradius.NewClient(
    loginradius.WithAPIKey(key),
    loginradius.WithHTTPClient(httpClient),
)
```

The SDK installs its auth/UA `RoundTripper` on top of your client's `Transport`, preserving your dialer/proxy/TLS configuration.

## Error handling

Every API call returns `(response, *http.Response, error)`. On non-2xx responses, `error` is `*loginradius.Error` carrying status, LoginRadius error code, message, and raw body:

```go
resp, httpResp, err := client.Login.CheckUserNameAvailability(ctx).Username("alice").Execute()
if err != nil {
    var lrErr *loginradius.Error
    if errors.As(err, &lrErr) {
        switch {
        case lrErr.IsAuth():
            // 401/403 — re-auth or escalate
        case lrErr.IsRateLimit():
            // 429 — back off
        case lrErr.IsServer():
            // 5xx — retry with backoff
        }
        log.Printf("status=%d code=%s msg=%s", lrErr.StatusCode, lrErr.Code, lrErr.Message)
    }
    return err
}
```

## Services

The `Client` exposes one field per OpenAPI tag. Each field is the typed service interface — call its methods to reach the underlying operations. Field naming drops the trailing `API` for brevity (e.g. `client.Login`, not `client.LoginAPI`).

```text
AccountCustomObject       Login                  PasskeyConfiguration         SecondFactorConfiguration
AccountSecurity           MultipurposeTokens     Password                     Security
AccountSession            OAuth                  PasswordPolicy               SecurityQuestions
Accounts                  OAuthClients           PerfectMindSSO               Session
BigCommerceSSO            OAuthCustomProviders   Permissions                  ShopifySSO
CaptchaConfiguration      OIDC                   PushNotificationConfig…      SMSTemplates
Consent                   Organization           Registration                 SOTT
CrossDeviceSSO            OrganizationConn…      Roles                        SocialProviders
CustomFields              OrganizationConn…      RolesManagement              User
CustomObject              OrganizationDomains    SAML                         UserMigration
CustomObjects             OrganizationInvit…     SAMLClients                  Webhooks
DomainAccessRestrictions  OrganizationUser…      SAMLCustomProviders          Workflows
EmailTemplates            Identity / Insights / JWT / JWTClients / JWTCustomProviders / IPAccessRestrictions
```

Use [godoc](https://pkg.go.dev/github.com/LoginRadius/go-sdk/v12) for the full method list per service.

## Migrating from v11

See [`MIGRATION_GUIDE.md`](./MIGRATION_GUIDE.md) — step-by-step diffs for
construction, auth, requests, responses, errors, and a representative
endpoint mapping table. Highlights:

- New module path: `github.com/LoginRadius/go-sdk/v12`.
- `lr.NewLoginradius(cfg, opts)` → `loginradius.NewClient(WithAPIKey(...), ...)`.
- Typed request/response models replace `body interface{}` and
  `*httprutils.Response`.
- `*loginradius.Error` with `IsAuth` / `IsForbidden` / `IsRateLimit` /
  `IsServer` predicates replaces the `lrerror.Error` interface.
- `apiRequestSigning` (`Digest` / `X-Request-Expires`) is supported via
  `WithAPIRequestSigning(true)`. It is opt-in and off by default. Validate it
  against your own tenant before relying on it.

## API reference

The API itself is documented at
[https://www.loginradius.com/docs/api/openapi/customer-identity-api](https://www.loginradius.com/docs/api/openapi/customer-identity-api) — endpoint behaviour, request and response fields, and what
each operation does. This SDK is generated from the same specification, so the
two stay in step.

[`docs/API.md`](./docs/API.md) lists every one of the 392
operations with its method name, HTTP verb and path, grouped across the
57 services.

Per-operation reference for all 392 operations is on
[pkg.go.dev](https://pkg.go.dev/github.com/LoginRadius/go-sdk/v12) — every service, method,
request builder and model, rendered from the source.

Both packages are documented there: the root package is the facade you use, and
`github.com/LoginRadius/go-sdk/v12/openapi` is the generated client it wraps. Start with the
facade — it adds credential handling, base-URL resolution, request signing and
the error envelope, none of which you get by constructing a generated client
directly.

## Validating a LoginRadius JWT

`ValidateJWT` verifies a token issued by one of your JWT apps. It is entirely
local — no network call, no credentials, no client.

```go
claims, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
    Algorithm: loginradius.JWTHS256,           // the algorithm YOUR app is configured for
    Key:       []byte(os.Getenv("LR_JWT_SECRET")),
    Issuer:    "LoginRadius",                  // optional
})
```

Signature, `exp` and `nbf` are always checked; issuer and audience are checked
when supplied. HS256/384/512 take the shared secret; RS*/ES* take the
PEM-encoded **public** key.

> **The algorithm is yours to state, and is never read from the token.** A
> validator that trusts the token's own `alg` header can be attacked: against an
> RS256 app, an attacker signs with HS256 using the public key as the HMAC
> secret. Passing the algorithm your app is configured for is what prevents it.

## Generated code

This SDK is generated from the LoginRadius OpenAPI specification, so the client,
the facade, the tests and the docs all stay in step with the API and with the
other LoginRadius SDKs.

Files under `v12/` are not edited by hand — a change made here would be lost on
the next release. If something is wrong, please open an issue rather than a pull
request against generated files; see [`GENERATED.md`](./GENERATED.md).

## Support

- Bugs and feature requests: <https://github.com/LoginRadius/go-sdk/issues>
- Account or integration questions: <support@loginradius.com>
- API documentation: <https://www.loginradius.com/docs/api/openapi/customer-identity-api>

## License

MIT — see the repo root `LICENSE`.
