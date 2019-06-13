// Package smartlogin contains methods for calling the LoginRadius Smart Login APIs. Smart Logins are
//  logins that allow a user to login through a unique client GUID. This is focused towards creating an easy method for Smart devices to
// access the LoginRadius authentication features. The client GUID is a unique identifier that can only be used once per login.
package smartlogin

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
