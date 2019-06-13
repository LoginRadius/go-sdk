// The MFA package contains methods for calling LoginRadius multifactor authentication APIs

//  Multifactor authentication should be enabled on the admin console before utilizing these APIs. Please note that the access tokens used for this section are different from the access tokens obtained from standard login and is obtained from calling the validate Backup Code/OTP/Google Auth Code API; this token is often referred to as secondfactorauthenticationtoken
package mfa

import (
	lr "github.com/LoginRadius/go-sdk"
)

// Embed Loginradius struct
type Loginradius struct {
	Client *lr.Loginradius
}
