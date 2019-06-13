// The role package contains methods for making calls to LoginRadius Role Management API endpoints
// The Roles APIs are used to manage the creation and assignment of user roles for a customer account.
package role

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
