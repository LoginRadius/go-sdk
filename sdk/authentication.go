/*
	This package is an SDK to interact with LoginRadius APIs.
	Documentation for the API can find in the following link:
	https://docs.loginradius.com/api/
*/

package loginradius

import (
	"os"
)

// PostAuthAddEmail is used to add additional emails to a user's account.
// Post Parameters are email: string and type: string
func PostAuthAddEmail(verificationURL, emailTemplate, authorization string, body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/email", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PostAuthForgotPassword is used to send the reset password url to a specified account.
// Note: If you have the UserName workflow enabled, you may replace the 'email' parameter with 'username'
// Post parameter is email: string
func PostAuthForgotPassword(resetPasswordURL, emailTemplate string, body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/password", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("resetpasswordurl", resetPasswordURL)
	q.Add("emailTemplate", emailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PostAuthUserRegistrationByEmail creates a user in the database as well as sends a verification email to the user.
// Post parameters are an array of email objects (Check docs for more info) and password: string
// To register an account with this function, make sure API secret is set in environment
func PostAuthUserRegistrationByEmail(verificationURL, emailTemplate, options string, body interface{}) (AuthRegister, error) {
	data := new(AuthRegister)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/register", body)
	if reqErr != nil {
		return *data, reqErr
	}

	sott := GenerateSOTT()
	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("verificationurl", verificationURL)
	q.Add("emailtemplate", emailTemplate)
	q.Add("options", options)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-Sott", sott)

	err := RunRequest(req, data)
	return *data, err
}

// PostAuthLoginByEmail retrieves a copy of the user data based on the Email.
func PostAuthLoginByEmail(verificationURL, loginURL, emailTemplate,
	gRecaptchaResponse, options string, body interface{}) (AuthLogin, error) {
	data := new(AuthLogin)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/login", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("verificationurl", verificationURL)
	q.Add("loginurl", loginURL)
	q.Add("emailtemplate", emailTemplate)
	q.Add("g-recaptcha-response", gRecaptchaResponse)
	q.Add("options", options)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PostAuthLoginByUsername retrieves a copy of the user data based on the Username.
// Post parameters are email: string, password: string and optional securityanswer: string
func PostAuthLoginByUsername(verificationURL, loginURL, emailTemplate,
	gRecaptchaResponse, options string, body interface{}) (AuthLogin, error) {
	data := new(AuthLogin)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/login", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("verificationurl", verificationURL)
	q.Add("loginurl", loginURL)
	q.Add("emailtemplate", emailTemplate)
	q.Add("g-recaptcha-response", gRecaptchaResponse)
	q.Add("options", options)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthCheckEmailAvailability is used to check whether an email exists or not on your site.
// Post parameters are email: string, password: string and optional securityanswer: string
func GetAuthCheckEmailAvailability(email string) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/email", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthCheckUsernameAvailability is used to check the UserName exists or not on your site.
func GetAuthCheckUsernameAvailability(username string) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/username", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("username", username)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthReadProfilesByToken retrieves a copy of the user data based on the access_token.
func GetAuthReadProfilesByToken(authorization string) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/account", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthPrivatePolicyAccept is used update the privacy policy stored in the user's profile
// by providing the access_token of the user accepting the privacy policy.
func GetAuthPrivatePolicyAccept(authorization string) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/privacypolicy/accept", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthSendWelcomeEmail will send the welcome email.
func GetAuthSendWelcomeEmail(welcomeEmailTemplate, authorization string) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/account/sendwelcomeemail", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("welcomeemailtemplate", welcomeEmailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthSocialIdentity is called just before account linking API and it prevents
// the raas profile of the second account from getting created.
func GetAuthSocialIdentity(authorization string) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/socialidentity", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthValidateAccessToken validates access token, if valid then returns a response with its expiry otherwise error.
func GetAuthValidateAccessToken(authorization string) (AuthAccessToken, error) {
	data := new(AuthAccessToken)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/access_token/validate", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthVerifyEmail is used to verify the email of user.
// Note: This API will only return the full profile if you have
// 'Enable auto login after email verification' set in your
// LoginRadius Dashboard's Email Workflow settings under 'Verification Email'.
func GetAuthVerifyEmail(verificationToken, url, welcomeEmailTemplate string) (AuthEmail, error) {
	data := new(AuthEmail)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/email", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("verificationtoken", verificationToken)
	q.Add("url", url)
	q.Add("welcomeemailtemplate", welcomeEmailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthDeleteAccount is used to delete an account by passing it a delete token.
func GetAuthDeleteAccount(deleteToken string) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/account/delete", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("deletetoken", deleteToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthInvalidateAccessToken invalidates the active access_token or expires an access token's validity.
func GetAuthInvalidateAccessToken(authorization string) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/access_token/invalidate", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthSecurityQuestionByAccessToken is used to retrieve the
// list of questions that are configured on the respective LoginRadius site.
func GetAuthSecurityQuestionByAccessToken(authorization string) (AuthSecurityQuestion, error) {
	data := new(AuthSecurityQuestion)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/securityquestion/accesstoken", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthSecurityQuestionByEmail is used to retrieve the
// list of questions that are configured on the respective LoginRadius site.
func GetAuthSecurityQuestionByEmail(email string) (AuthSecurityQuestion, error) {
	data := new(AuthSecurityQuestion)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/securityquestion/email", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthSecurityQuestionByUsername is used to retrieve the
// list of questions that are configured on the respective LoginRadius site.
func GetAuthSecurityQuestionByUsername(username string) (AuthSecurityQuestion, error) {
	data := new(AuthSecurityQuestion)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/securityquestion/username", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("username", username)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetAuthSecurityQuestionByPhone is used to retrieve the
// list of questions that are configured on the respective LoginRadius site.
func GetAuthSecurityQuestionByPhone(phone string) (AuthSecurityQuestion, error) {
	data := new(AuthSecurityQuestion)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/securityquestion/phone", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("phone", phone)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthVerifyEmailByOtp will send the welcome email.
// Post parameters include otp: string, email: string, optional securityanswer: string, optional qq_captcha_ticket: string,
// optional qq_captcha_randstr: string and optional g-recaptcha-response:string
func PutAuthVerifyEmailByOtp(url, welcomeEmailTemplate string, body interface{}) (AuthEmail, error) {
	data := new(AuthEmail)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/email", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("url", url)
	q.Add("welcomeemailTemplate", welcomeEmailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthChangePassword is used to change the accounts password based on the previous password.
// Post parameters include oldpassword: string and newpassword: string
func PutAuthChangePassword(authorization string, body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/password/change", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthLinkSocialIdentities is used to link up a social provider account with the specified
// account based on the access token and the social providers user access token.
// Post parameter is the candidatetoken: string
func PutAuthLinkSocialIdentities(authorization string, body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/socialidentity", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PutResendEmailVerification resends the verification email to the user.
// Post parameter is the email: string
func PutResendEmailVerification(verificationURL, emailTemplate string, body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/register", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("verificationurl", verificationURL)
	q.Add("emailtemplate", emailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthResetPasswordByResetToken is used to set a new password for the specified account.
// Post parameters are the resettoken: string, password: string, optional welcomeemailtemplate: string
// and optional resetpasswordemailtemplate: string
func PutAuthResetPasswordByResetToken(body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/password/reset", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthResetPasswordByOTP is used to set a new password for the specified account.
// Post parameters are the password: string, otp: string, email: string,
// optional welcomeemailtemplate: string and optional resetpasswordemailtemplate: string
func PutAuthResetPasswordByOTP(body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/password/reset", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthResetPasswordBySecurityAnswerAndEmail is used to reset password for the specified account by security question.
// Post parameters are the password: string, email: string, securityanswer: string
// and optional resetpasswordemailtemplate: string
func PutAuthResetPasswordBySecurityAnswerAndEmail(body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/password/securityanswer", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthResetPasswordBySecurityAnswerAndPhone is used to reset password for the specified account by security question.
// Post parameters are the password: string, phone: string, securityanswer: string
// and optional resetpasswordemailtemplate: string
func PutAuthResetPasswordBySecurityAnswerAndPhone(body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/password/securityanswer", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthResetPasswordBySecurityAnswerAndUsername is used to reset password for the specified account by security question.
// Post parameters are the password: string, username: string, securityanswer: string
// and optional resetpasswordemailtemplate: string
func PutAuthResetPasswordBySecurityAnswerAndUsername(body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/password/securityanswer", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthSetOrChangeUsername is used to set or change UserName by access token.
// Post parameter is username: string
func PutAuthSetOrChangeUsername(authorization string, body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/username", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthUpdateProfileByToken is used to update the user's profile by passing the access_token.
// Post parameters are fields in the profile that need to be updated
func PutAuthUpdateProfileByToken(verificationURL, emailTemplate,
	smsTemplate, authorization string, body interface{}) (AuthUpdate, error) {
	data := new(AuthUpdate)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/account", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("verificationurl", verificationURL)
	q.Add("emailtemplate", emailTemplate)
	q.Add("smstemplate", smsTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PutAuthUpdateSecurityQuestionByAccessToken is used to update security questions by the access token.
// Post parameter is the securityquestionanswer: string
func PutAuthUpdateSecurityQuestionByAccessToken(authorization string, body interface{}) (AuthUpdate, error) {
	data := new(AuthUpdate)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/account", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// DeleteAuthDeleteAccountEmailConfirmation deletes a user account by passing the user's access token.
func DeleteAuthDeleteAccountEmailConfirmation(deleteURL, emailTemplate, authorization string) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/auth/account", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("deleteurl", deleteURL)
	q.Add("emailtemplate", emailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// DeleteAuthRemoveEmail is used to remove additional emails from a user's account.
// Post parameter is the e-mail that is to be removed.
func DeleteAuthRemoveEmail(authorization string, body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/auth/email", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// DeleteAuthUnlinkSocialIdentities is used to unlink up a social provider account with the specified account
// based on the access token and the social providers user access token.
// The unlinked account will automatically get removed from your database.
func DeleteAuthUnlinkSocialIdentities(authorization string, body interface{}) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/auth/socialidentity", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetPasswordlessLoginByEmail is used to send a Passwordless Login verification link to the provided Email ID.
func GetPasswordlessLoginByEmail(email, passwordlessLoginTemplate, verificationURL string) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/login/passwordlesslogin/email", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("email", email)
	q.Add("passwordlesslogintemplate", passwordlessLoginTemplate)
	q.Add("verificationurl", verificationURL)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetPasswordlessLoginByUsername is used to send a Passwordless Login verification link to the provided Username.
func GetPasswordlessLoginByUsername(username, passwordlessLoginTemplate, verificationURL string) (AuthBool, error) {
	data := new(AuthBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/login/passwordlesslogin/email", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("username", username)
	q.Add("passwordlesslogintemplate", passwordlessLoginTemplate)
	q.Add("verificationurl", verificationURL)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetPasswordlessLoginVerification is used to verify the Passwordless Login verification link.
func GetPasswordlessLoginVerification(verificationToken, welcomeEmailTemplate string) (AuthLogin, error) {
	data := new(AuthLogin)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/login/passwordlesslogin/email/verify", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("verificationtoken", verificationToken)
	q.Add("welcomeemailtemplate", welcomeEmailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}
