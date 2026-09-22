module github.com/LoginRadius/go-sdk/v12

go 1.24

require (
	// v5.2.1 carried GHSA-mh63-6h87-95cp / GO-2025-3553 (excessive memory
	// allocation parsing a JWT header) — directly relevant here, since this
	// package parses untrusted JWT headers. Fixed at 5.2.2; pinned to the
	// latest 5.x (5.3.1).
	github.com/golang-jwt/jwt/v5 v5.3.1
	gopkg.in/validator.v2 v2.0.1
)
