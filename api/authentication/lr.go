// lrauthentication contains APIs for Loginradius authentication end points
// The Authentication (Auth) APIs allow changes to the account once some form of authentication has been performed.
// For this reason, they are considered to be user facing client-side/front-end API calls.
package lrauthentication

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
