package customobject

import (
	"fmt"

	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"

	"github.com/LoginRadius/go-sdk/httprutils"
)

// PostCustomObjectCreateByUID is used to write information in JSON format to the custom object for the specified account.

// Post parameters: custom data to be created in the object.

// Required query parameter: objectname - string

// Required template parameter: string representing uid

// Please ensure this feature is enabled for your LoginRadius account

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/create-custom-object-by-uid
func (lr Loginradius) PostCustomObjectCreateByUID(uid string, queries interface{}, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPostReq("/identity/v2/manage/account/"+uid+"/customobject", body, validatedQueries)
	if err != nil {
		return nil, err
	}

	lr.Client.AddApiCredentialsToReqHeader(req)
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// PostCustomObjectCreateByToken is used to write information in JSON format to the custom object for the specified account.

// Post parameters: custom data to be created in the object.

// Required query parameter: objectname - string

// Please ensure this feature is enabled for your LoginRadius account

// Documentation - https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/create-custom-object-by-token
func (lr Loginradius) PostCustomObjectCreateByToken(queries interface{}, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewPostReqWithToken("/identity/v2/auth/customobject", body, validatedQueries)
	if err != nil {
		return nil, err
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

//GetCustomObjectByObjectRecordIDAndUID is used to retrieve the Custom Object data for the specified account.

// Required query parameter: objectname - string

// Required template parameter: string representing uid, string representing objectrecordid

// Please ensure this feature is enabled for your LoginRadius account

// Documentation - https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-by-objectrecordid-and-uid
func (lr Loginradius) GetCustomObjectByObjectRecordIDAndUID(uid, objectRecordID string, queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req := lr.Client.NewGetReq("/identity/v2/manage/account/"+uid+"/customobject/"+objectRecordID, validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(req)
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetCustomObjectByObjectRecordIDAndToken is used to retrieve the Custom Object data for the specified account.

// Required query parameter: objectname - string; apikey - string

// Required template parameter: string representing objectrecordid

// Please ensure this feature is enabled for your LoginRadius account

// Documentation - https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-by-objectrecordid-and-token
func (lr Loginradius) GetCustomObjectByObjectRecordIDAndToken(objectRecordID string, queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/customobject/"+objectRecordID, validatedQueries)
	if err != nil {
		return nil, err
	}
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetCustomObjectByToken is used to retrieve the specified Custom Object data for the specified account.

// Required parameters: objectname - string; apikey - string

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-by-token
func (lr Loginradius) GetCustomObjectByToken(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req, err := lr.Client.NewGetReqWithToken("/identity/v2/auth/customobject", validatedQueries)
	if err != nil {
		return nil, err
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// GetCustomObjectByUID is used to retrieve all the custom objects by UID from cloud storage.

// Required parameters: objectname - string

// Required template parameter: string representing uid

// Documentation - https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-by-uid
func (lr Loginradius) GetCustomObjectByUID(uid string, queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	req := lr.Client.NewGetReq("/identity/v2/manage/account/"+uid+"/customobject/", validatedQueries)
	lr.Client.AddApiCredentialsToReqHeader(req)
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// PutCustomObjectUpdateByUID is used to update the specified custom object data of a specified account.

// Post parameters:the fields that need to be changed.

// Required query parameters: objectname - string; updatetype - string

// If the value of updatetype is 'replace' then it will fully replace custom object with new custom object and
// if the value of updatetype is partialreplace then it will perform an upsert type operation.

// Required template parameters: string representing uid, string repesenting objectrecordid

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-update-by-objectrecordid-and-uid
func (lr Loginradius) PutCustomObjectUpdateByUID(uid, objectrecordid string, queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true, "updatetype": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid+"/customobject/"+objectrecordid, body, validatedQueries)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(req)

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// PutCustomObjectUpdateByToken is used to update the specified custom object data of a specified account.

// Post parameters:the fields that need to be changed.

// Required query parameters: objectname - string; updatetype - string

// If the value of updatetype is 'replace' then it will fully replace custom object with new custom object and
// if the value of updatetype is partialreplace then it will perform an upsert type operation.

// Required template parameters: string repesenting objectrecordid

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-update-by-objectrecordid-and-token
func (lr Loginradius) PutCustomObjectUpdateByToken(objectrecordid string, queries, body interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"objectname": true, "updatetype": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req, err := lr.Client.NewPutReqWithToken("/identity/v2/auth/customobject/"+objectrecordid, body, validatedQueries)
	if err != nil {
		return nil, err
	}

	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// DeleteCustomObjectByObjectRecordIDAndUID is used to remove the
// specified Custom Object data using ObjectRecordId of specified account.

// Required template parameters: string representing uid, string representing objectrecordid

// Required query parameters: objectname

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-delete-by-objectrecordid-and-uid
func (lr Loginradius) DeleteCustomObjectByObjectRecordIDAndUID(uid, objectRecordId string, queries interface{}) (*httprutils.Response, error) {

	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req := lr.Client.NewDeleteReq("/identity/v2/manage/account/" + uid + "/customobject/" + objectRecordId)
	req.QueryParams = validatedQueries
	lr.Client.AddApiCredentialsToReqHeader(req)
	req.Headers["content-Type"] = "application/json"
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}

// DeleteCustomObjectByObjectRecordIDAndUID is used to remove the

// specified Custom Object data using ObjectRecordId of specified account.

// Required template parameters: string representing objectrecordid

// Required query parameters: objectname

// Documentation - https://www.loginradius.com/docs/api/v2/customer-identity-api/custom-object/custom-object-delete-by-objectrecordid-and-token
func (lr Loginradius) DeleteCustomObjectByObjectRecordIDAndToken(objectRecordId string, queries interface{}) (*httprutils.Response, error) {

	allowedQueries := map[string]bool{"objectname": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	req, err := lr.Client.NewDeleteReqWithToken("/identity/v2/auth/customobject/"+objectRecordId, "", validatedQueries)
	if err != nil {
		return nil, err
	}
	fmt.Println(req)
	req.QueryParams = validatedQueries
	req.Headers["content-Type"] = "application/json"
	lr.Client.NormalizeApiKey(req)
	resp, err := httprutils.TimeoutClient.Send(*req)
	return resp, err
}
