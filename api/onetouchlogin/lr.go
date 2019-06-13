// Package onetouchlogin contain methods for calling LoginRadius One Touch Login APIs,
// which use email and phone verification to create links that allow the user to login.
package onetouchlogin

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
