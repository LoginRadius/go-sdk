package lrauthentication

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// GetAuthVerifyEmail is used to verify the email of user.

// Note: This API will only return the full profile if you have'Enable auto login after email verification' set in your
// LoginRadius Dashboard's Email Workflow settings under 'Verification Email'

// Documentation:https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-verify-email

// Required query parameters: apiKey, verificationtoken;  Optional query parameter: url
func (lr Loginradius) GetAuthVerifyEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"url": true, "verificationtoken": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey

	req := lr.Client.NewGetReq("/identity/v2/auth/email", validatedQueries)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthCheckEmailAvailability is used to check whether an email exists or not on your site.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-verify-email

// Required query parameters: apiKey, email
func (lr Loginradius) GetAuthCheckEmailAvailability(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"email": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey

	req := lr.Client.NewGetReq("/identity/v2/auth/email", validatedQueries)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthCheckUsernameAvailability is used to check the UserName exists or not on your site.

// Documentation:
// https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-username-availability

// Required query parameters: apiKey, username
func (lr Loginradius) GetAuthCheckUsernameAvailability(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"username": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey

	req := lr.Client.NewGetReq("/identity/v2/auth/username", validatedQueries)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthReadProfilesByToken retrieves a copy of the user data based on the access token.

// Required query parameters: apiKey

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-read-profiles-by-token
func (lr Loginradius) GetAuthReadProfilesByToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account")
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthPrivatePolicyAccept is used update the privacy policy stored in the user's profile based on user's access token

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-privacy-policy-accept
func (lr Loginradius) GetAuthPrivatePolicyAccept() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/privacypolicy/accept")
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthSendWelcomeEmail sends the welcome email.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-send-welcome-email
func (lr Loginradius) GetAuthSendWelcomeEmail(queries ...interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account/sendwelcomeemail")
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{"welcomeemailtemplate": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}

	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthSocialIdentity is called just before account linking API and it prevents
// the profile of the second account from getting created.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-social-identity

// Required query parameters: apiKey
func (lr Loginradius) GetAuthSocialIdentity() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/socialidentity")
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthValidateAccessToken returns an expiry date for the access token if it is valid
// and an error if it is invalid

//Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-validate-access-token

// Required query parameters: apiKey
func (lr Loginradius) GetAuthValidateAccessToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/access_token/validate")
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthDeleteAccount is used to delete an account by passing it a delete token.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-delete-account

// Required query parameters: apiKey, deletetoken
func (lr Loginradius) GetAuthDeleteAccount(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"deletetoken": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/identity/v2/auth/account/delete", validatedQueries)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthInvalidateAccessToken invalidates the active access_token or expires an access token's validity.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-invalidate-access-token

// Required query parameter: apiKey; optional query parameter: preventRefresh
func (lr Loginradius) GetAuthInvalidateAccessToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/access_token/invalidate")
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthSecurityQuestionByAccessToken is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question is enabled

// Documentation: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy

// Required query parameters: apiKey
func (lr Loginradius) GetAuthSecurityQuestionByAccessToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/securityquestion/accesstoken")
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthSecurityQuestionByEmail is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question feature is enabled

// Documentation: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy

// Required query parameters: apiKey, email
func (lr Loginradius) GetAuthSecurityQuestionByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"email": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/securityquestion/email", validatedQueries)
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthSecurityQuestionByUsername is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question feature is enabled.

// Documentation: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy

// Required query parameters: apikey
func (lr Loginradius) GetAuthSecurityQuestionByUsername(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"username": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/securityquestion/username", validatedQueries)
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAuthSecurityQuestionByPhone is used to retrieve the
// list of questions that are configured on the respective LoginRadius site for the user.
// Will return error unless security question feature is enabled

// Documentation: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy

// Required query parameters: phone
func (lr Loginradius) GetAuthSecurityQuestionByPhone(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"phone": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/securityquestion/phone", validatedQueries)
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetPasswordlessLoginByEmail is used to send a Passwordless Login verification link to the provided Email ID.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-by-email

// Required query parameters: email, apiKey; optional queries: passwordlesslogintemplate, verificationurl
func (lr Loginradius) GetPasswordlessLoginByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"email": true, "passwordlesslogintemplate": true, "verificationurl": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey

	req := lr.Client.NewGetReq("/identity/v2/auth/login/passwordlesslogin/email", validatedQueries)

	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetPasswordlessLoginByUsername is used to send a Passwordless Login verification link to the provided Username.

// Required query parameters: username, apiKey; optional queries: passwordlesslogintemplate, verificationurl

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-by-username
func (lr Loginradius) GetPasswordlessLoginByUsername(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"username": true, "passwordlesslogintemplate": true, "verificationurl": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/identity/v2/auth/login/passwordlesslogin/email", validatedQueries)

	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetPasswordlessLoginVerification is used to verify the Passwordless Login verification link.

// Required query parameters: verificationtoken; optional queries: welcomeemailtemplate

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-verification
func (lr Loginradius) GetPasswordlessLoginVerification(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"verificationtoken": true, "welcomeemailtemplate": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apiKey"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/identity/v2/auth/login/passwordlesslogin/email/verify", validatedQueries)

	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
