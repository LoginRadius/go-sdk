package onetouchlogin

import (
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"

	"github.com/LoginRadius/go-sdk/httprutils"
)

// PostOneTouchLoginByEmail is used to send a link to a specified email for a frictionless login/registration
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/one-touch-login/one-touch-login-by-email-captcha
// Required query parameter: apikey
// Optional query parameters: redirecturl, OneTouchLoginEmailTemplate, welcomeemailtemplate
// Required post parameters: clientguid - string; email - string; g-recaptcha-response - string;
// Optional post parameters: qq_captcha_ticket - string; qq_captcha_randstr - string;
func (lr Loginradius) PostOneTouchLoginByEmail(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	validatedQueries := map[string]string{}
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"redirecturl": true, "OneTouchLoginEmailTemplate": true, "welcomeemailtemplate": true,
		}
		validated, err := lrvalidate.Validate(allowedQueries, arg)
		if err != nil {
			return nil, err
		}
		for k, v := range validated {
			validatedQueries[k] = v
		}
	}

	req, err := lr.Client.NewPostReq("/identity/v2/auth/onetouchlogin/email", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PostOneTouchLoginByPhone is used to send a link to a specified email for a frictionless login/registration
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/one-touch-login/one-touch-login-by-email-captcha
// Required query parameter: apikey
// Optional query parameters: redirecturl, OneTouchLoginEmailTemplate, welcomeemailtemplate
// Required post parameters: clientguid - string; phone - string; g-recaptcha-response - string;
// Optional post parameters: qq_captcha_ticket - string; qq_captcha_randstr - string;
func (lr Loginradius) PostOneTouchLoginByPhone(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	validatedQueries := map[string]string{}
	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"redirecturl": true, "OneTouchLoginEmailTemplate": true, "welcomeemailtemplate": true,
		}
		validated, err := lrvalidate.Validate(allowedQueries, arg)
		if err != nil {
			return nil, err
		}
		for k, v := range validated {
			validatedQueries[k] = v
		}
	}

	req, err := lr.Client.NewPostReq("/identity/v2/auth/onetouchlogin/phone", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutOneTouchOTPVerification is used to verify the otp for One Touch Login.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/one-touch-login/one-touch-otp-verification
// Required query parameters: apikey, otp
// Optional query parameter: smstemplate
// Required post parameter: phone - string;
func (lr Loginradius) PutOneTouchOTPVerification(queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"otp": true, "smstemplate": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req, err := lr.Client.NewPutReq("/identity/v2/auth/onetouchlogin/phone/verify", body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
