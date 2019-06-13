// Package tokenmanagement contain methods for calling the LoginRadius token management APIs, which allow management of access tokens and generation tokens usable by the social APIs.
package tokenmanagement

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
