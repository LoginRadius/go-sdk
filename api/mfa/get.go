package mfa

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// GetMFAValidateAccessToken is used to configure the Multi-factor authentication

// after login by using the access_token when MFA is set as optional on the LoginRadius site.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-validate-access-token

// Required query parameter: apikey

// Optional query parameter: smstemplate2fa

// Needs Authorization Bearer token header
func (lr Loginradius) GetMFAValidateAccessToken(queries ...interface{}) (*httprutils.Response, error) {
	queryParams := map[string]string{}

	for _, arg := range queries {
		allowedQueries := map[string]bool{"smstemplate2fa": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)
		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			queryParams[k] = v
		}
	}
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account/2fa", queryParams)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetMFABackUpCodeByAccessToken is used to get a set of backup codes via access_token to allow the user login on a site that has Multi-factor Authentication enabled in the event that the user does not have a secondary factor available. We generate 10 codes, each code can only be consumed once. If any user attempts to go over the number of invalid login attempts configured in the Dashboard then the account gets blocked automatically

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-backup-code-by-access-token

// Required query parameter: apikey
func (lr Loginradius) GetMFABackUpCodeByAccessToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account/2fa/backupcode")
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

//GetMFABackUpCodeByAccessToken is used to reset the backup codes on a given account via the access_token. This API call will generate 10 new codes, each code can only be consumed once.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-reset-backup-code-by-access-token

// Required query parameter: apikey
func (lr Loginradius) GetMFAResetBackUpCodeByAccessToken() (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account/2fa/backupcode/reset")
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetMFABackUpCodeByUID is used to get a set of backup codes to allow the user login on a site that has Multi-factor
// authentication enabled in the event that the user does not have a secondary factor available.
// We generate 10 codes, each code can only be consumed once.
// If any user attempts to go over the number of invalid login attempts configured in the
// admin console then the account gets blocked automatically

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-backup-code-by-uid

// Required query parameter: apikey, apisecret, uid
func (lr Loginradius) GetMFABackUpCodeByUID(queries interface{}) (*httprutils.Response, error) {
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

	req := lr.Client.NewGetReq("/identity/v2/manage/account/2fa/backupcode", queryParams)
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

//GetMFAResetBackUpCodeByUID is used to reset the backup codes on a given account via the UID.
//This API call will generate 10 new codes, each code can only be consumed once.

//Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/mfa-reset-backup-code-by-uid

// Required query parameter: apikey, apisecret, uid
func (lr Loginradius) GetMFAResetBackUpCodeByUID(queries interface{}) (*httprutils.Response, error) {
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

	req := lr.Client.NewGetReq("/identity/v2/manage/account/2fa/backupcode/reset", queryParams)
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// This API is used to trigger the Multi-Factor Autentication workflow for the provided access_token

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/multi-factor-authentication/re-authentication/mfa-re-authenticate

// Required query parameters: apikey

// Optional query parameter: smstemplate2fa
func (lr Loginradius) GetMFAReAuthenticate(queries ...interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/account/reauth/2fa")
	for _, arg := range queries {
		allowedQueries := map[string]bool{"smstemplate2fa": true}
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
