package mfa

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// PutMFAValidateGoogleAuthCode is used to login via Multi-factor-authentication by passing the google authenticator code.

// Documentation:https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/google-authenticator/mfa-validate-google-auth-code

// Required query parameters: apikey, secondfactorauthenticationtoken

// secondfactorauthenticationtoken can be obtained by successful logins through MFA login routes

// Optional query parameters: smstemplate2fa

// Required post parameter: googleauthenticatorcode: string
func (lr Loginradius) PutMFAValidateGoogleAuthCode(queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"secondfactorauthenticationtoken": true, "smstemplate2fa": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPutReq("/identity/v2/auth/login/2fa/verification/googleauthenticatorcode", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAValidateOTP is used to login via Multi-factor authentication by passing the One Time Password received via SMS.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-validate-otp

// Required query parameters: apikey, secondfactorauthenticationtoken

// Optional query parameter: smstemplate2fa

// Required post parameter: otp - string

// Optional query parameters: securityanswer, g-recaptcha-response, qq_captcha_ticket, qq_captcha_randstr
func (lr Loginradius) PutMFAValidateOTP(queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"secondfactorauthenticationtoken": true, "smstemplate2fa": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPutReq("/identity/v2/auth/login/2fa/verification/otp", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

//PutMFAUpdateByToken  used to Enable Multi-factor authentication by access token on user login.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/google-authenticator/update-mfa-by-access-token

// Required query parameters: apikey

// Optional query parameter: smstemplate

// Rerquired post parameters: googleauthenticatorcode - string
func (lr Loginradius) PutMFAUpdateByToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account/2FA/Verification/GoogleAuthenticatorCode", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"smstemplate": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAUpdatePhoneNumber is used to update (if configured) the phone number used for Multi-factor authentication by sending the verification OTP to the provided phone number.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-update-phone-number

// Required query parameters: apikey, secondfactorauthenticationtoken

// Optional query parameter: smstemplate2fa

// Required post parameter: phoneno2fa - string
func (lr Loginradius) PutMFAUpdatePhoneNumber(queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"secondfactorauthenticationtoken": true, "smstemplate2fa": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPutReq("/identity/v2/auth/login/2fa", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAUpdatePhoneNumberByToken is used to update (if configured) the phone number used for Multi-factor authentication by sending the verification OTP to the provided phone number.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-update-phone-number-by-token

// Required query parameters: apikey

// Optional query parameter: smstemplate2fa

// Required post parameter: phoneno2fa - string
func (lr Loginradius) PutMFAUpdatePhoneNumberByToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account/2fa", body)
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"smstemplate2fa": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}

	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// This API is used to validate the backup code provided by the user and if valid, we return an access_token allowing the user to login incases where Multi-factor authentication (MFA) is enabled and the secondary factor is unavailable. When a user intially downloads the Backup codes, We generate 10 codes, each code can only be consumed once. if any user attempts to go over the number of invalid login attempts configured in the Dashboard then the account gets blocked automatically

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-validate-backup-code

// Required query parameters: apikey, secondfactorauthenticationtoken

// Required body parameter: backupcode
func (lr Loginradius) PutMFAValidateBackupCode(queries interface{}, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"secondfactorauthenticationtoken": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPutReq("/identity/v2/auth/login/2fa/verification/backupcode", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAReauthenticateByGoogleAuthenticator is used to re-authenticate via Multi-factor-authentication by passing the google authenticator code.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/re-authentication/validate-mfa-by-google-authenticator-code

// Required query parameter: apikey

// Required body parameters: googleauthenticatorcode
func (lr Loginradius) PutMFAReauthenticateByGoogleAuthenticator(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account/reauth/2fa/GoogleAuthenticatorCode", body)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAReauthenticateByBackupCode is used to re-authenticate via Multi-factor-authentication by passing the backup code.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/re-authentication/validate-mfa-by-backup-code
// Required query parameter: apikey
// Required body parameters: backupcode
func (lr Loginradius) PutMFAReauthenticateByBackupCode(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account/reauth/2fa/BackupCode", body)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAReauthenticateByOTP is used to re-authenticate via Multi-factor-authentication by passing the sms otp.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/re-authentication/validate-mfa-by-otp

// Required query parameter: apikey

// Required body parameters: otp

// Optional bodys parameters: securityanswer, qq_captcha_ticket, qq_captcha_Randstr, g-recaptcha-response
func (lr Loginradius) PutMFAReauthenticateByOTP(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account/reauth/2fa/otp", body)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAReauthenticateByPassword is used to re-authenticate via Multi-factor-authentication by passing the guser's password.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/re-authentication/validate-mfa-by-password

// Required query parameter: apikey

// Required body parameters: password

// Optional body parameters: securityanswer - object, qq_captcha_ticket, qq_captcha_Randstr, g-recaptcha-response
func (lr Loginradius) PutMFAReauthenticateByPassword(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account/reauth/password", body)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutMFAUpdateSettingsis used to trigger the Multi-factor authentication settings after login for secure actions.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/update-mfa-setting

// Required query parameter: apikey

// Required body parameter: otp - string

// optional body parameters: securityanswer - object; g-recaptcha-response - string; qq_captcha_ticket - string; qq-captcha-randstr - string
func (lr Loginradius) PutMFAUpdateSettings(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/account/2FA/Verification/otp", body)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
