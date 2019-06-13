package mfa

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// DeleteMFAResetGoogleAuthenticatorByToken resets the Google Authenticator configurations on a given account via the access_token.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/google-authenticator/mfa-reset-google-authenticator-by-token

// Required query parameter: apikey

// Required body parameter: googleauthenticator - pass true as value
func (lr Loginradius) DeleteMFAResetGoogleAuthenticatorByToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewDeleteReqWithToken(
		"/identity/v2/auth/account/2fa/authenticator",
		map[string]bool{"googleauthenticator": true},
	)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	req.Headers["content-Type"] = "application/json"
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteMFAResetSMSAuthenticatorByToken resets the SMS Authenticator configurations on a given account via the access_token.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-reset-sms-authenticator-by-token

// Required query parameter: apikey

// Required body parameter: otpauthenticator - pass true as value
func (lr Loginradius) DeleteMFAResetSMSAuthenticatorByToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewDeleteReqWithToken(
		"/identity/v2/auth/account/2fa/authenticator",
		map[string]bool{"otpauthenticator": true},
	)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	req.Headers["content-Type"] = "application/json"
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteMFAResetSMSAuthenticatorByUid resets the SMS Authenticator configurations on a given account via the access_token.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/sms-authenticator/mfa-reset-sms-authenticator-by-uid
// Required query parameter: apikey, apisecret, uid

// Required body parameter: otpauthenticator - pass true as value
func (lr Loginradius) DeleteMFAResetSMSAuthenticatorByUid(queries interface{}) (*httprutils.Response, error) {
	queryParams := map[string]string{}
	uid, ok := queries.(string)
	if ok {
		queryParams["uid"] = uid
	} else {
		allowedQueries := map[string]bool{"uid": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
		if err != nil {
			return nil, err
		}
		queryParams = validatedQueries
	}
	queryParams["apikey"] = lr.Client.Context.ApiKey
	queryParams["apisecret"] = lr.Client.Context.ApiSecret

	req := lr.Client.NewDeleteReq(
		"/identity/v2/manage/account/2fa/authenticator",
		map[string]bool{"otpauthenticator": true},
		queryParams,
	)
	req.QueryParams = queryParams
	req.Headers = httprutils.JSONHeader
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteMFAResetGoogleAuthenticatorByUid resets the SMS Authenticator configurations on a given account via the access_token.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/google-authenticator/mfa-reset-google-authenticator-by-uid

// Required query parameter: apikey, apisecret, uid

// Required body parameter: googleauthenticator - pass true as value
func (lr Loginradius) DeleteMFAResetGoogleAuthenticatorByUid(queries interface{}) (*httprutils.Response, error) {
	queryParams := map[string]string{}
	uid, ok := queries.(string)
	if ok {
		queryParams["uid"] = uid
	} else {
		allowedQueries := map[string]bool{"uid": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
		if err != nil {
			return nil, err
		}
		queryParams = validatedQueries
	}
	queryParams["apikey"] = lr.Client.Context.ApiKey
	queryParams["apisecret"] = lr.Client.Context.ApiSecret

	req := lr.Client.NewDeleteReq(
		"/identity/v2/manage/account/2fa/authenticator",
		map[string]bool{"googleauthenticator": true},
	)
	req.Headers = httprutils.JSONHeader
	req.QueryParams = queryParams
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
