package lrauthentication

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// PostAuthAddEmail is used to add additional emails to a user's account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-add-email#

// Pass data in struct lrbody.AddEmail as body to help ensure parameters satisfy API requirements

// Required query parameters: apiKey; optional queries: verificationurl, emailtemplate

// Required body parameters: email -string, type -string
func (lr Loginradius) PostAuthAddEmail(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReqWithToken("/identity/v2/auth/email", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "emailtemplate": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PostAuthForgotPassword is used to send the reset password url to a specified account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-forgot-password

// Note: If you have the UserName workflow enabled, you may replace the 'email' parameter with 'username'

// Required query parameters: apikey, resetpasswordurl; optional query parameter: emailtemplate

// Required post parameter - email: string

// Pass data in struct lrbody.EmailStr as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PostAuthForgotPassword(body interface{}, queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"resetpasswordurl": true, "emailtemplate": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)

	if err != nil {
		return nil, err
	}

	request, err := lr.Client.NewPostReq("/identity/v2/auth/password", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PostAuthUserRegistrationByEmail creates a user in the database as well as sends a verification email to the user.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-user-registration-by-email

// Required post parameter: email - array(Check docs for more info); password: string
// Required  parameter: sott
// Pass data in struct lrbody.RegistrationUser as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PostAuthUserRegistrationByEmail(sott string,body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	
	queryParams := map[string]string{}
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "emailtemplate": true, "options": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			queryParams[k] = v
		}
	}
	queryParams["apiKey"] = lr.Client.Context.ApiKey
	request, err := lr.Client.NewPostReq("/identity/v2/auth/register", body, queryParams)

	request.Headers["X-LoginRadius-Sott"] = sott
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PostAuthLoginByEmail retrieves a copy of the user data based on the Email after verifying
// the validity of submitted credentials

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-login-by-email

// Pass data in struct lrbody.EmailLogin as body to help ensure parameters satisfy API requirements

// Required query parameters: apiKey; optional query parameters: verificationurl, loginurl, emailtemplate, g-recaptcha-response

// Required body parameters: email, password; optional body parameters: security answer
func (lr Loginradius) PostAuthLoginByEmail(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/identity/v2/auth/login", body)
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "loginurl": true, "emailtemplate": true, "g-recaptcha-response": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PostAuthLoginByUsername retrieves a copy of the user data based on the Username after verifying
// the validity of submitted credentials

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-login-by-username

// Required post parameters - username: string, password: string; optional post parameter - securityanswer: string

// Pass data in struct lrbody.UsernameLogin as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PostAuthLoginByUsername(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/identity/v2/auth/login", body)
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "loginurl": true, "emailtemplate": true, "g-recaptcha-response": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
