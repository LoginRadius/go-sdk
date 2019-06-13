package lraccount

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// PutManageAccountUpdateSecurityQuestionConfig is used to update security questions configuration on an existing account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-update-security-question

// Required post parameter: security question answer - object.

// Required template parameter: uid

// For more information regarding security questions, refer to this document: https://docs.loginradius.com/api/v2/dashboard/platform-security/password-policy

// Pass data in struct lrbody.AccountSecurityQuestion as body to help ensure parameters satisfy API requirements; alternatively,
// []byte could also be passed as body
func (lr Loginradius) PutManageAccountUpdateSecurityQuestionConfig(uid string, body interface{}) (*httprutils.Response, error) {

	request, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid, body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(request)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutManageAccountSetPassword is used to set the password of an account in Cloud Storage.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-set-password

// Required post parameter: password - string

// Required template parameter: uid

// Pass data in struct lrbody.AccountPassword as body to help ensure parameters satisfy API requirements; alternatively,
// []byte could also be passed as body
func (lr Loginradius) PutManageAccountSetPassword(uid string, body interface{}) (*httprutils.Response, error) {

	request, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid+"/password", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(request)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutManageAccountUpdate is used to update the information of existing accounts in your Cloud Storage.
// See our Advanced API Usage section for more capabilities.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-update

// Optional query parameter: nullsupport (takes true/false)

// Post parameters: profile data that needs to be updated. - modify struct fields based on need

// Pass data in struct lrbody.UpdateProfile as body to help ensure parameters satisfy API requirements; alternatively,
// []byte could also be passed as body
func (lr Loginradius) PutManageAccountUpdate(uid string, body interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid, body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(request)

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PutManageAccountInvalidateVerificationEmail is used to invalidate the Email Verification status on an account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-invalidate-verification-email

// Optional query params: verificationurl,  emailtemplate
func (lr Loginradius) PutManageAccountInvalidateVerificationEmail(uid string, queries ...interface{}) (*httprutils.Response, error) {
	request, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid+"/invalidateemail", "")
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

	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
