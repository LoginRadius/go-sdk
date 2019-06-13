// Package lrsocial contains methods for calling the LoginRadius social apis, which are used to fetch user profile and other data from providers linked to the user accounts.
package lrsocial

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
