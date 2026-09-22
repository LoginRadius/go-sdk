# v12 changelog

## v12.0.0-rc.1

Initial v12 release. v12 is a new module path (`github.com/LoginRadius/go-sdk/v12`) — v11 customers are not affected.

### What's new

- **Generated from the OpenAPI spec.** Every operation defined in
  `LoginRadius-Public-APIs.yaml` is reachable through a typed service on the
  `Client` (392 operations across 57 services).
- **Typed request and response models.** ~1,500 schemas are accessible as Go
  structs through top-level aliases — no need to construct request bodies via
  `interface{}` and JSON marshaling.
- **Centralized authentication.** Nine credentials are injected via a single
  `http.RoundTripper`; configure them once at `NewClient`. The spec declares 12
  security schemes: `Digest` and `XRequestExpiresTime` are produced by
  `WithAPIRequestSigning` rather than configured as static credentials, and
  `ApiSecret` (the `secret=` query parameter, distinct from `apisecret=`) is
  not sent. API key and secret are sent via the `X-LoginRadius-ApiKey` /
  `X-LoginRadius-ApiSecret` headers by preference (keeping them out of access
  logs and URL caches); the query-string scheme is also populated to support
  the ~56 endpoints in the spec that only accept it.
- **Typed errors.** `*loginradius.Error` exposes HTTP status, LoginRadius
  error code, message, and raw body, plus helpers (`IsAuth`, `IsRateLimit`,
  `IsServer`).
- **Custom HTTP client support.** `WithHTTPClient` preserves caller's
  Transport, dialer, and TLS config under the SDK's auth/UA RoundTripper.
- **Server selection.** `WithDomain`, `WithCustomDomain`, `WithBaseURL`.

### Migrating from v11

See [README — Migrating from v11](./README.md#migrating-from-v11).

### Spec changes that landed with this release

Three operationId case-collisions in the source spec were corrected:

- `validateAccessToken` (Session tag) → `ValidateSessionAccessToken`
- `getAccessToken` (Session tag) → `GetSessionAccessToken`
- `createTenantRole` (Organization tag) → `CreateOrgTenantRole`

PascalCase variants under other tags kept their original operationIds.
