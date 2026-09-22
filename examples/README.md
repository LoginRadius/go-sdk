# Examples

Each subdirectory is a self-contained `main` package. Pick the one that
matches the LoginRadius auth scheme(s) the operation you're calling uses.

All examples read credentials from environment variables. Copy `v12/.env.example`
to `v12/.env`, fill in only the keys for the examples you plan to run, and
load them however you prefer (`source .env`, `direnv`, etc.).

To run any example:

```bash
export $(grep -v '^#' .env | xargs)   # or use direnv / your loader of choice
go run ./examples/<name>
```

## By concept

| Directory       | Demonstrates                                                          |
| --------------- | --------------------------------------------------------------------- |
| `quickstart/`   | Minimum setup — `loginradius.NewClient(WithAPIKey(...))` + one call   |
| `login/`        | Passwordless email login + typed `*loginradius.Error` branching       |
| `custom-http/`  | Caller-supplied `*http.Client` (proxy, timeout, custom transport)     |

## By auth scheme

Match the example to the security scheme(s) listed for the LoginRadius
operation you're calling in `LoginRadius-Public-APIs.yaml`.

| Directory                 | Auth scheme(s)                                                    | Sample operation                                       |
| ------------------------- | ----------------------------------------------------------------- | ------------------------------------------------------ |
| `quickstart/`             | `APIKey`                                                          | `Login.CheckUserNameAvailability`                      |
| `api-key-secret/`         | `APIKey` + `APISecret`                                            | `AccountSession.GetAccessToken`                        |
| `access-token/`           | `AccessToken` (query)                                             | `OIDC.GetOIDCUserinfo`                                 |
| `bearer-token/`           | `BearerToken` (`Authorization: Bearer …`)                         | `User.GetAccountDetails`                               |
| `m2m-bearer-token/`       | `M2MBearerToken` (JWT bearer)                                     | `SOTT.GetAllSOTT`                                      |
| `client-id-secret/`       | `ClientId` + `ClientSecret` (query)                               | `MultipurposeTokens.MultipurposeEmailTokenAPI`         |
| `x-loginradius-headers/`  | `XLoginRadiusAPIKey` + `XLoginRadiusAPISecret` (header overrides) | `AccountSession.GetAccessToken`                        |

## Combining schemes

The SDK accepts every credential at once — supply whichever schemes the
operation you're calling lists in the spec:

```go
client, _ := loginradius.NewClient(
    loginradius.WithAPIKey(os.Getenv("LR_API_KEY")),
    loginradius.WithAPISecret(os.Getenv("LR_API_SECRET")),
    loginradius.WithAccessToken(os.Getenv("LR_ACCESS_TOKEN")),
    loginradius.WithBearerToken(os.Getenv("LR_BEARER_TOKEN")),
)
```

The auth `http.RoundTripper` only sets headers/query for the credentials
that are non-empty. If you're worried about cross-tenant credential bleeding,
construct one `*loginradius.Client` per scheme instead.

## Cross-cutting request options

These six examples cover the client-wide features the legacy v11 SDK had and
v12 was missing, plus two defects that silently ignored your configuration.

**They run offline.** Each installs a recording transport instead of sending
anything, so you can inspect exactly what the SDK put on the wire without a
tenant, credentials, or network access:

```bash
go run ./examples/request-signing
```

| Directory               | Verifies                                                                                   |
| ----------------------- | ------------------------------------------------------------------------------------------ |
| `request-signing/`      | `WithAPIRequestSigning` — `digest` + `x-Request-Expires` on `/manage/` paths only, `apisecret` stripped before signing |
| `request-options/`      | `WithOriginIP`, `WithServerRegion`, `WithFields`, `WithPreventWebhook` applied to every request |
| `default-headers/`      | `WithDefaultHeaders` merges, but can never mask a credential or the User-Agent              |
| `debug-logging/`        | `WithDebug` redacts every credential value; exits non-zero if one leaks                     |
| `operation-servers/`    | The 42 spec-pinned operations now honour `WithBaseURL` / `WithDomain` instead of the spec's placeholder tenant |
| `timeout-http-client/`  | `WithTimeout` applies to an injected client only when set explicitly                        |
