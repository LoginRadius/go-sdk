package loginradius

import (
	"encoding/json"
	"os"
	"time"
)

// CustomObject holds data retrieved from the custom objects API
type CustomObject struct {
	ID           string          `json:"Id"`
	IsActive     bool            `json:"IsActive"`
	DateCreated  time.Time       `json:"DateCreated"`
	DateModified time.Time       `json:"DateModified"`
	IsDeleted    bool            `json:"IsDeleted"`
	UID          string          `json:"Uid"`
	CustomObject json.RawMessage `json:"CustomObject"`
}

// CustomObjectMulti holds data retrieved from the custom objects API when multiples are returned
type CustomObjectMulti struct {
	Data []struct {
		ID           string          `json:"Id"`
		IsActive     bool            `json:"IsActive"`
		DateCreated  time.Time       `json:"DateCreated"`
		DateModified time.Time       `json:"DateModified"`
		IsDeleted    bool            `json:"IsDeleted"`
		UID          string          `json:"Uid"`
		CustomObject json.RawMessage `json:"CustomObject"`
	} `json:"data"`
	Count int `json:"Count"`
}

// PostCustomObjectCreateByUID is used to write information in JSON format to the custom object for the specified account.
// Post parameters are the custom data to be created in the object.
func PostCustomObjectCreateByUID(objectName, uid string, body interface{}) (CustomObject, error) {
	data := new(CustomObject)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/customobject", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("objectname", objectName)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PostCustomObjectCreateByToken is used to write information in JSON format to the custom object for the specified account.
// Post parameters are the custom data to be created in the object.
func PostCustomObjectCreateByToken(objectName, authorization string, body interface{}) (CustomObject, error) {
	data := new(CustomObject)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN") + "/identity/v2/auth/customobject", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("objectname", objectName)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetCustomObjectByObjectRecordIDAndUID is used to retrieve the Custom Object data for the specified account.
func GetCustomObjectByObjectRecordIDAndUID(objectName, uid, objectRecordID string) (CustomObject, error) {
	data := new(CustomObject)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/customobject/"+objectRecordID, "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("objectname", objectName)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// GetCustomObjectByObjectRecordIDAndToken is used to retrieve the Custom Object data for the specified account.
func GetCustomObjectByObjectRecordIDAndToken(objectName, authorization, objectRecordID string) (CustomObject, error) {
	data := new(CustomObject)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/auth/customobject/"+objectRecordID, "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("objectname", objectName)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetCustomObjectByToken is used to retrieve the specified Custom Object data for the specified account.
func GetCustomObjectByToken(objectName, authorization string) (CustomObjectMulti, error) {
	data := new(CustomObjectMulti)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/auth/customobject", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("objectname", objectName)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetCustomObjectByUID is used to retrieve all the custom objects by UID from cloud storage.
func GetCustomObjectByUID(objectName, uid string) (CustomObjectMulti, error) {
	data := new(CustomObjectMulti)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/customobject/", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("objectname", objectName)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PutCustomObjectUpdateByUID is used to update the specified custom object data of a specified account.
// If the value of updatetype is 'replace' then it will fully replace custom object with new custom object and
// if the value of updatetype is partialreplace then it will perform an upsert type operation.
// Post parameters are the fields that need to be changed.
func PutCustomObjectUpdateByUID(objectName, updateType, uid, objectRecordID string, body interface{}) (CustomObject, error) {
	data := new(CustomObject)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/customobject/"+
		objectRecordID, body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("objectname", objectName)
	q.Add("updatetype", updateType)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PutCustomObjectUpdateByToken is used to update the specified custom object data of a specified account.
// If the value of updatetype is 'replace' then it will fully replace custom object with new custom object and
// if the value of updatetype is partialreplace then it will perform an upsert type operation.
// Post parameters are the fields that need to be changed.
func PutCustomObjectUpdateByToken(objectName, updateType, authorization,
	objectRecordID string, body interface{}) (CustomObject, error) {
	data := new(CustomObject)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/auth/customobject/"+objectRecordID, body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("objectname", objectName)
	q.Add("updatetype", updateType)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// DeleteCustomObjectByObjectRecordIDAndUID is used to remove the
// specified Custom Object data using ObjectRecordId of specified account.
func DeleteCustomObjectByObjectRecordIDAndUID(objectName, uid, objectRecordID string) (CustomObject, error) {
	data := new(CustomObject)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/customobject/"+
		objectRecordID, "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("objectname", objectName)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// DeleteCustomObjectByObjectRecordIDAndToken is used to remove the
// specified Custom Object data using ObjectRecordId of specified account.
func DeleteCustomObjectByObjectRecordIDAndToken(objectName, authorization, objectRecordID string) (CustomObject, error) {
	data := new(CustomObject)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/auth/customobject/"+objectRecordID, "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("objectname", objectName)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}
