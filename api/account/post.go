package lraccount

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// PostManageAccountCreate is used to create an account in Cloud Storage.
// This API bypasses the normal email verification process and manually creates the user.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-create

// In order to use this API, you need to format a JSON request body with all of the mandatory fields

// Required post parameters: email - object; Password - string. Rest are optional profile parameters.

// Required query parameters: apiKey, apiSecret

// Pass data in struct lrbody.AccountCreate as body to help ensure parameters satisfy API requirements; alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) PostManageAccountCreate(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/identity/v2/manage/account", body)
	if err != nil {
		return nil, err
	}

	lr.Client.AddApiCredentialsToReqHeader(request)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PostManageForgotPasswordToken returns a forgot password token. Note: If you have the
// UserName workflow enabled, you may replace the 'email' parameter with 'username'.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/get-forgot-password-token

// Required post parameters: email - string OR username - string

// Optional query parameters: sendemail; emailTemplate; resetPasswordUrl

// Pass data in struct lrbody.Username or lrbody.Email as body to help ensure parameters satisfy API requirements; alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) PostManageForgotPasswordToken(body interface{}, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/identity/v2/manage/account/forgot/token", body)
	if err != nil {
		return nil, err
	}

	lr.Client.AddApiCredentialsToReqHeader(request)

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"sendemail": true, "emailTemplate": true, "resetPasswordUrl": true,
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

// PostManageEmailVerificationToken Returns an Email Verification token.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/get-email-verification-token

// Post parameter - email: string

// Pass data in struct lrbody.EmailForVToken as body to help ensure parameters satisfy API requirements; alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) PostManageEmailVerificationToken(body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPostReq("/identity/v2/manage/account/verify/token", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
