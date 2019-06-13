package lraccount

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// GetManageAccountProfilesByEmail is used to retrieve all of the profile data,
// associated with the specified account by email in Cloud Storage.
// This end point returns a single profile

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-profiles-by-email

// Required query parameter: email
func (lr Loginradius) GetManageAccountProfilesByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"email": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountProfilesByUsername is used to retrieve all of the profile data,
// associated with the specified account by username in Cloud Storage.
// This end point returns a single profile

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-profiles-by-user-name

// Required query parameter: username
func (lr Loginradius) GetManageAccountProfilesByUsername(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"username": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountProfilesByPhoneID is used to retrieve all of the profile data,
// associated with the specified account by PhoneID in Cloud Storage.
// This end point returns a single profile

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-profiles-by-phone-id

// Required query param: phone
func (lr Loginradius) GetManageAccountProfilesByPhoneID(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"phone": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountProfilesByUid is used to retrieve all of the profile data,
// associated with the specified account by uid in Cloud Storage.
// This end point returns a single profile

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-profiles-by-uid

// Required template param: uid - string representing uid
func (lr Loginradius) GetManageAccountProfilesByUid(uid string) (*httprutils.Response, error) {
	request := lr.Client.NewGetReq("/identity/v2/manage/account/" + uid)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountIdentitiesByEmail is used to retrieve all of the identities (UID and Profiles),
// associated with a specified email in Cloud Storage.
// Note: This is intended for specific workflows where an email may be associated to multiple UIDs.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-identities-by-email

// This end point returns data in an array, the response needs to be handled like so:
// 						body, _ := lrjson.DynamicUnmarshal(response.Body) // unmarshals body
// 						profiles := body["Data"].([]interface{}) // type assertion
// 						profile := profiles[0].(map[string]interface{}) // get first profile
// 						uid := profile["Uid"].(string) // get id of first profile

// Required query param: email - string
func (lr Loginradius) GetManageAccountIdentitiesByEmail(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"email": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account/identities", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccessTokenUID is used to get LoginRadius access token based on UID.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-impersonation-api

// Required query params: uid
func (lr Loginradius) GetManageAccessTokenUID(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"uid": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	request := lr.Client.NewGetReq("/identity/v2/manage/account/access_token", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// GetManageAccountPassword is used to retrieve the hashed password of a specified account in Cloud Storage.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-password

// Required template parameter: string representing uid
func (lr Loginradius) GetManageAccountPassword(uid string) (*httprutils.Response, error) {
	request := lr.Client.NewGetReq("/identity/v2/manage/account/" + uid + "/password")
	lr.Client.AddApiCredentialsToReqHeader(request)
	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
