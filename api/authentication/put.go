package lrauthentication

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// PutAuthVerifyEmailByOtp will send the welcome email.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-verify-email-by-otp

// Required post parameters - otp: string; email: string

// Optional post parameters - qq_captcha_randstr: string;g-recaptcha-response:string; securityanswer; string; qq_captcha_ticket: string;

// Required query parameter: apiKey; Optional query parameters: url, welcometemplate
func (lr Loginradius) PutAuthVerifyEmailByOtp(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/email", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"url": true, "welcomeemailtemplate": true,
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

// PutAuthChangePassword is used to change the accounts password based on the previous password.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-change-password

// Post parameters- oldpassword: string;newpassword: string

// Required query paramter: apikey

// Pass data in struct lrbody.ChangePassword as body to help ensure parameters satisfy API requirements; alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) PutAuthChangePassword(body interface{}) (*httprutils.Response, error) {

	request, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/password/change", body)

	if err != nil {
		return nil, err
	}

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutAuthLinkSocialIdentities is used to link up a social provider account with the specified
// account based on the access token and the social providers user access token.

// Document: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-link-social-identities

// Post parameters- candidatetoken: string

// Pass data in struct lrbody.LinkSocialIds as body to help ensure parameters satisfy API requirements
func (lr Loginradius) PutAuthLinkSocialIdentities(token string, body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/socialidentity", body)
	if err != nil {
		return nil, err
	}

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutResendEmailVerification resends the verification email to the user.

// Post parameter- email: string

// Required query parameter: apikey ; optional query parameters: emailtemplate, verificationurl

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-resend-email-verification

// Pass data in struct lrbody.EmailStr as body to help ensure parameters satisfy API requirements; alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) PutResendEmailVerification(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/register", body)
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

// PutAuthResetPasswordByResetToken is used to set a new password for the specified account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-reset-token

// Required post parameters- resettoken: string; password: string

// optional post parameters- welcomeemailtemplate: string; resetpasswordemailtemplate: string

// Required query parameter: apiKey

// Pass data in struct lrbody.ResetPw as body to help ensure parameters satisfy API requirement;alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) PutAuthResetPasswordByResetToken(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/password/reset", body)
	if err != nil {
		return nil, err
	}

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutAuthResetPasswordByOTP is used to set a new password for the specified account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-otp

// Required query parameter: apiKey

// Required post parameters -the password: string; otp: string; email: string

// Optional post parameters: welcomeemailtemplate: string; resetpasswordemailtemplate: string

// Pass data in struct lrbody.ResetPwOtp as body to help ensure parameters satisfy API requirements;alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) PutAuthResetPasswordByOTP(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/password/reset", body)

	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"welcomeemailtemplate": true, "resetpasswordemailtemplate": true, "otp": true, "email": true, "password": true,
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

// PutAuthResetPasswordBySecurityAnswerAndEmail is used to reset password for the specified account by security question.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-email

// For additional information on security questions, please see https://www.loginradius.com/docs/api/v2/dashboard/platform-security/password-policy#securityquestion4

// Required post parameters - password: string; email: string; securityanswer: string

// Optional post parameters: resetpasswordemailtemplate: string

// Required query parameter: apiKey

// Pass data in struct lrbody.ResetPwSecurityQuestionEmail as body to help ensure parameters satisfy API requirements;alternatively,
// []byte could also be passed as body
func (lr Loginradius) PutAuthResetPasswordBySecurityAnswerAndEmail(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/password/securityanswer", body)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutAuthResetPasswordBySecurityAnswerAndPhone is used to reset password for the specified account by security question.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-phone

// For additional information on security questions, please see https://www.loginradius.com/docs/api/v2/dashboard/platform-security/password-policy#securityquestion4

// Required post parameters - password: string; email: string; securityanswer: string

// Optional post parameters: resetpasswordemailtemplate: string

// Required query parameter: apiKey

// Pass data in struct lrbody.ResetPwSecurityQuestionPhone as body to help ensure parameters satisfy API requirements; alternatively,
// []byte could also be passed as body
func (lr Loginradius) PutAuthResetPasswordBySecurityAnswerAndPhone(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/password/securityanswer", body)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutAuthResetPasswordBySecurityAnswerAndUsername is used to reset password for the specified account by security question.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-reset-password-by-username

// For additional information on security questions, please see https://www.loginradius.com/docs/api/v2/dashboard/platform-security/password-policy#securityquestion4

// Required post parameters - password: string; email: string; securityanswer: string

// Optional post parameters: resetpasswordemailtemplate: string

// Required query parameter: apiKey

// Pass data in struct lrbody.ResetPwSecurityQuestionusername as body to help ensure parameters satisfy API requirements; alternatively,
// []byte could also be passed as body
func (lr Loginradius) PutAuthResetPasswordBySecurityAnswerAndUsername(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/password/securityanswer", body)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutAuthSetOrChangeUsername is used to set or change UserName by access token.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-set-or-change-user-name

// Post parameter is username: string

// Pass data in struct lrbody.AuthUsername as body to help ensure parameters satisfy API requirements; alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) PutAuthSetOrChangeUsername(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/username", body)

	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutAuthUpdateProfileByToken is used to update the user's profile by passing the access_token.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-update-profile-by-token

// Post parameters are fields in the profile that need to be updated -- modify struct fields based on need

// Required query parameter: apiKey; optional query parameters: smstemplate, emailtemplate, verificationurl

// Pass data in struct lrbody.UpdateProfile as body to help ensure parameters satisfy API requirements; alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) PutAuthUpdateProfileByToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"verificationurl": true, "emailtemplate": true, "smstemplate": true,
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

// PutAuthUpdateSecurityQuestionByAccessToken is used to update security questions by the access token.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-update-security-question-by-access-token

// Required body parameter: securityquestionanswer: string

// Pass data in struct lrbody.SecurityQuestionAnswer as body to help ensure parameters satisfy API requirements; alternatively,
// []byte or map[string]string{} could also be passed as body

// For more information on this parameter, please see: https://www.loginradius.com/docs/api/v2/dashboard/platform-security/password-policy#securityquestion4
func (lr Loginradius) PutAuthUpdateSecurityQuestionByAccessToken(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account", body)
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
