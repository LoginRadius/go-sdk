package smartlogin

import (
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"

	"github.com/LoginRadius/go-sdk/httprutils"
)

// GetSmartLoginByEmail sends a Smart Login link to the user's Email Id.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-by-email
// Required query parameters: apikey, email, clientguid
// Optional query parameters: smartloginemailtemplate, welcomeemailtemplate, redirecturl
func (lr Loginradius) GetSmartLoginByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"email": true, "clientguid": true, "smartloginemailtemplate": true, "welcomeemailtemplate": true, "redirecturl": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req := lr.Client.NewGetReq("/identity/v2/auth/login/smartlogin", validatedQueries)
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetSmartLoginByUsername sends a Smart Login link to the user's Email Id.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-by-username
// Required query parameters: apikey, username, clientguid
// Optional query parameters: smartloginemailtemplate, welcomeemailtemplate, redirecturl
func (lr Loginradius) GetSmartLoginByUsername(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"username": true, "clientguid": true, "smartloginemailtemplate": true, "welcomeemailtemplate": true, "redirecturl": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req := lr.Client.NewGetReq("/identity/v2/auth/login/smartlogin", validatedQueries)
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetSmartLoginPing is used to check if the Smart Login link has been clicked or not.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-ping
// Required query parameters: apikey, clientguid
func (lr Loginradius) GetSmartLoginPing(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"clientguid": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req := lr.Client.NewGetReq("/identity/v2/auth/login/smartlogin/ping", validatedQueries)
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetSmartLoginVerifyToken verifies the provided token for Smart Login.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/smart-login/smart-login-verify-token
// Required query parameterS: apikey, verificationtoken,
// Optional query parameters: welcommeemailtemplate
func (lr Loginradius) GetSmartLoginVerifyToken(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"clientguid": true, "verificationtoken": true, "welcomeemailtemplate": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req := lr.Client.NewGetReq("/identity/v2/auth/email/smartlogin", validatedQueries)
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
