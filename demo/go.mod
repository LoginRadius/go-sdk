// The demo is its OWN module, deliberately.
//
// As part of the SDK module it went into the module zip every consumer
// downloads, and any dependency it grew would have joined their graph. A nested
// module is excluded from the parent's zip and resolves its own requirements.
//
// The replace directive points at the checkout above so the demo always builds
// against the adjacent source rather than a published version — which is what
// you want when the demo exists to exercise unreleased changes.
module github.com/LoginRadius/go-sdk/v12/demo

go 1.24

// The version is ignored — `replace` below wins — but it must still parse,
// and a /v12 module path only accepts a v12.x.y version here.
require github.com/LoginRadius/go-sdk/v12 v12.0.0-rc.1

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	gopkg.in/validator.v2 v2.0.1 // indirect
)

replace github.com/LoginRadius/go-sdk/v12 => ../
