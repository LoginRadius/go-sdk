package phoneauthentication

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// PostPhoneLogin retrieves a copy of the user data based on the Phone.

// Required post parameters: phone - string; password - string;

// Optional post parameters: securityanswer - object - required when account locked and unlock strategy is securityanswer

// For more information on this parameter, please see: https://www.loginradius.com/docs/api/v2/dashboard/platform-security/password-policy#securityquestion4

// Required query parameter: apikey

// Optional query parameters: loginurl - string; smstemplate-string; g-recaptcha-response - string

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-login
func (lr Loginradius) PostPhoneLogin(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPostReq("/identity/v2/auth/login", body)
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"loginurl": true, "smstemplate": true, "g-recaptcha-response": true,
		}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// PostPhoneForgotPasswordByOTP is used to send the OTP to reset the account password.

// Required query parameter: apikey - string

// Optional query parameter: smstemplate

// Required post parameter: phone - string

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-forgot-password-by-otp
func (lr Loginradius) PostPhoneForgotPasswordByOTP(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPostReq("/identity/v2/auth/password/otp", body)
	if err != nil {
		return nil, err
	}
	for _, arg := range queries {
		allowedQueries := map[string]bool{"smstemplate": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// PostPhoneResendVerificationOTP is used to resend a verification OTP to verify a user's Phone Number.

// The user will receive a verification code that they will need to input.

// Required query parameter: apikey - string

// Optional query parameter: smstemplate

// Required post parameter: phone - string

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-resend-otp
func (lr Loginradius) PostPhoneResendVerificationOTP(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPostReq("/identity/v2/auth/phone/otp", body)
	if err != nil {
		return nil, err
	}
	for _, arg := range queries {
		allowedQueries := map[string]bool{"smstemplate": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// PostPhoneResendVerificationOTPByToken is used to resend a verification OTP to verify a user's Phone Number in cases in which an active token already exists.

// Required query parameter: apikey - string

// Optional query parameter: smstemplate

// Required post parameter: phone - string

// Requires user access token to be submited in Authorization Bearer header

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-resend-otp-by-token
func (lr Loginradius) PostPhoneResendVerificationOTPByToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPostReqWithToken("/identity/v2/auth/phone/otp", body)
	if err != nil {
		return nil, err
	}
	for _, arg := range queries {
		allowedQueries := map[string]bool{"smstemplate": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// PostPhoneUserRegistrationBySMS registers the new users into your Cloud Storage and triggers the phone verification process.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-user-registration-by-sms

// Required query parameter: apikey

// Optional query parameters: verificationurl, smstemplate, options (takes value PreventVerificationEmail)

// Required body parameters: email, password, and other form fields configured for your LoginRadius app

// Optional body parameters: other optional profile fields for your user
// Required  parameter: sott

func (lr Loginradius) PostPhoneUserRegistrationBySMS(sott string,body interface{}, queries ...interface{}) (*httprutils.Response, error) {
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

// GetPhoneSendOTP is used to send your phone an OTP.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-by-phone

// Required query parameters: apikey, phone

// Optional query parameter: smstemplate
func (lr Loginradius) GetPhoneSendOTP(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"phone": true, "smstemplate": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apikey"] = lr.Client.Context.ApiKey
	request := lr.Client.NewGetReq("/identity/v2/auth/login/passwordlesslogin/otp", validatedQueries)
	delete(request.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*request)
	return res, err
}

// GetPhoneNumberAvailability is used to check the whether the phone number exists or not on your site.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-number-availability

// Required query parameter: apikey, phone
func (lr Loginradius) GetPhoneNumberAvailability(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"phone": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req := lr.Client.NewGetReq("/identity/v2/auth/phone", validatedQueries)
	lr.Client.NormalizeApiKey(req)
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// PutPhoneLoginUsingOTP is used to login using OTP flow.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/passwordless-login/passwordless-login-phone-verification

// Required query parameter: apikey

// Optional query parameter: smstemplate

// Required post parameters: phone - string; otp - string

// Optional post parameters: securityanswer - string; g-recaptcha-response - string; qq_captcha_ticket - string; qq_captcha_randstr - string
func (lr Loginradius) PutPhoneLoginUsingOTP(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/auth/login/passwordlesslogin/otp/verify", body)
	for _, arg := range queries {
		allowedQueries := map[string]bool{"smstemplate": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			request.QueryParams[k] = v
		}
	}
	lr.Client.NormalizeApiKey(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutPhoneNumberUpdate is used to update the phone number of a user.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-number-update

// Required query parameter: apikey

// Optional query parameter: smstemplate

// Required post parameter: phone - string (the new number to be updated for the account)
func (lr Loginradius) PutPhoneNumberUpdate(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	queryParams := map[string]string{}
	for _, arg := range queries {
		allowedQueries := map[string]bool{"smstemplate": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			queryParams[k] = v
		}
	}
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/phone", body, queryParams)
	lr.Client.NormalizeApiKey(req)
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutPhoneResetPasswordByOTP is used to reset the password.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-reset-password-by-otp

// Required query parameter: apikey

// Optional post parameters: smstemplate - string; resetpasswordmailtemplate - string

// Required post parameters: phone - string; otp - string; password-string
func (lr Loginradius) PutPhoneResetPasswordByOTP(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReq("/identity/v2/auth/password/otp", body)
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutPhoneVerificationByOTP is used to validate the verification code sent to verify a user's phone number.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-verify-otp

// Required query parameters: apikey, otp

// Optional query parameter: smstemplate

// Required post parameter: phone - string
func (lr Loginradius) PutPhoneVerificationByOTP(queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"otp": true, "smstemplate": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req, err := lr.Client.NewPutReq("/identity/v2/auth/phone/otp", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutPhoneVerificationByOTPByToken is used to consume the verification code sent to verify a user's phone number.

// Use this call for front-end purposes in cases where the user is already logged in by passing the user's access token.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/phone-verify-otp-by-token

// Required query parameters: apikey, otp

// Optional query parameter: smstemplate

// Requires Authorization Bearer token
func (lr Loginradius) PutPhoneVerificationByOTPByToken(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"otp": true, "smstemplate": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/phone/otp", "", validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)

	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutResetPhoneIDVerification allows you to reset the phone number verification of an end user’s account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/reset-phone-id-verification

// Required template parameter: string representing uid of the user profile
func (lr Loginradius) PutResetPhoneIDVerification(uid string) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid+"/invalidatephone", "")
	if err != nil {
		return nil, err
	}
	req.QueryParams = map[string]string{
		"apikey":    lr.Client.Context.ApiKey,
		"apisecret": lr.Client.Context.ApiSecret,
	}
	req.Headers = httprutils.URLEncodedHeader
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteRemovePhoneIDByAccessToken is used to delete the Phone ID on a user's account via the access_token.

// Required query parameter: apikey

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/phone-authentication/remove-phone-id-by-access-token
func (lr Loginradius) DeleteRemovePhoneIDByAccessToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewDeleteReqWithToken("/identity/v2/auth/phone", "")
	if err != nil {
		return nil, err
	}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
