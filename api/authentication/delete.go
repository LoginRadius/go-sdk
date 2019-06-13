package lrauthentication

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// DeleteAuthDeleteAccountEmailConfirmation sends a confirmation email for account deletion to the customer's email when passed the access token

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-delete-account-with-email-confirmation

// Required query param: apiKey

// Optional query params: deleteurl, emailtemplate
func (lr Loginradius) DeleteAuthDeleteAccountEmailConfirmation(queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewDeleteReqWithToken("/identity/v2/auth/account", "")
	if err != nil {
		return nil, err
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{
			"deleteurl": true, "emailtemplate": true,
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

// DeleteAuthRemoveEmail is used to remove additional emails from a user's account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-remove-email

// Post parameter - email: string.

// Pass data in struct lrbody.EmailStr as body to help ensure parameters satisfy API requirements; alternatively,
// []byte or map[string]string{} could also be passed as body
func (lr Loginradius) DeleteAuthRemoveEmail(body interface{}) (*httprutils.Response, error) {

	request, err := lr.Client.NewDeleteReqWithToken("/identity/v2/auth/email", body)
	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// DeleteAuthUnlinkSocialIdentities is used to unlink up a social provider account with the specified account
// based on the access token and the social providers user id, the latter is returned with any API call that returns the full
// user profile. The unlinked account will automatically get removed from your database.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/authentication/auth-unlink-social-identities

// Required body parameters: provider, providerid

// Required query parameter: apiKey
func (lr Loginradius) DeleteAuthUnlinkSocialIdentities(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewDeleteReqWithToken("/identity/v2/auth/socialidentity", body)

	if err != nil {
		return nil, err
	}
	response, err := httprutils.TimeoutClient.Send(*req)
	return response, err
}
