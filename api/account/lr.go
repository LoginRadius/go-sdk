// the lraccount package contains
// The Account Management APIs are used to manage a user's account.
// These calls require the API Key and API Secret and often the User's Account UID(Unified Identifier) to perform an operation.
// For this reason these APIs are considered to be for back-end purposes.
package lraccount

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
