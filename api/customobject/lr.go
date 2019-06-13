// customobject package contains methods for calling the LoginRadius Custom Object APIs,
// which are used to interact with custom objects that are set on the customer account.
package customobject

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
