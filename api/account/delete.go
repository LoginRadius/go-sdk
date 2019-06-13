package lraccount

import (
	"github.com/LoginRadius/go-sdk/httprutils"
)

// DeleteManageAccount is used to delete the Users account and allows them to re-register for a new account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-email-delete

// Required template variable: uid
func (lr Loginradius) DeleteManageAccount(uid string) (*httprutils.Response, error) {
	request := lr.Client.NewDeleteReq("/identity/v2/manage/account/")
	lr.Client.AddApiCredentialsToReqHeader(request)
	request.URL = request.URL + uid

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// DeleteManageAccount is used to delete the Users account and allows them to re-register for a new account.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-delete

// Required template variable: uid

// Required body parameter: email
func (lr Loginradius) DeleteManageAccountEmail(uid string, body interface{}) (*httprutils.Response, error) {
	encoded, err := httprutils.EncodeBody(body)
	if err != nil {
		return nil, err
	}
	request := httprutils.Request{
		Method: httprutils.Delete,
		URL:    lr.Client.Domain + "/identity/v2/manage/account/" + uid + "/email",
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		Body: encoded,
	}
	lr.Client.AddApiCredentialsToReqHeader(&request)
	response, err := httprutils.TimeoutClient.Send(request)
	return response, err
}
