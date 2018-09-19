package loginradius

import (
	"os"
	"time"
)

// CustomRegistrationBool holds data methods that return a boolean
type CustomRegistrationBool struct {
	IsPosted  bool `json:"IsPosted"`
	IsDeleted bool `json:"IsDeleted"`
	IsValid   bool `json:"IsValid"`
}

// CustomRegistrationData holds data custom registration data
type CustomRegistrationData struct {
	Code         string    `json:"Code"`
	ID           string    `json:"Id"`
	DateCreated  time.Time `json:"DateCreated"`
	DateModified time.Time `json:"DateModified"`
	IsActive     bool      `json:"IsActive"`
	Type         string    `json:"Type"`
	Key          string    `json:"Key"`
	Value        string    `json:"Value"`
	ParentID     string    `json:"ParentId"`
	ParentType   string    `json:"ParentType"`
}

// CustomRegistrationDataList holds custom registration data
type CustomRegistrationDataList []CustomRegistrationData

// CustomRegistrationUpdateData holds data returned after updating registration data
type CustomRegistrationUpdateData struct {
	IsPosted bool                   `json:"IsPosted"`
	Data     CustomRegistrationData `json:"Data"`
}

// PostCustomRegistrationAddData allows you to fill data in dropDownList which you have created for user Registration.
// Required Post Parameters are type: string, key: string, value: string
// Optional Post Parameters are isactive: boolean, parentid: string, code: string
func PostCustomRegistrationAddData(body interface{}) (CustomRegistrationBool, error) {
	data := new(CustomRegistrationBool)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/manage/registrationdata", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PostCustomRegistrationValidateSecretCode allows you to validate code for a particular dropdown member.
// Required Post Parameters are recordid: string, code: string
func PostCustomRegistrationValidateSecretCode(body interface{}) (CustomRegistrationBool, error) {
	data := new(CustomRegistrationBool)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/registrationdata/validatecode", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// GetCustomRegistrationData is used to retrieve dropdown data through management APIs.
func GetCustomRegistrationData(dataType, parentid, skip, limit string) (CustomRegistrationDataList, error) {
	data := new(CustomRegistrationDataList)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/registrationdata/"+dataType, "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("parentid", parentid)
	q.Add("skip", skip)
	q.Add("limit", limit)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// GetCustomRegistrationDataServer is used to retrieve dropdown data through authentication APIs.
func GetCustomRegistrationDataServer(dataType, parentid, skip, limit string) (CustomRegistrationDataList, error) {
	data := new(CustomRegistrationDataList)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/registrationdata/"+dataType, "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("parentid", parentid)
	q.Add("skip", skip)
	q.Add("limit", limit)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutCustomRegistrationDataUpdate allows you to update member of dropDown.
// Required Post Parameters are type: string, key: string, value: string
// Optional Post Parameters are isactive: boolean, parentid: string, code: string
func PutCustomRegistrationDataUpdate(recordID string, body interface{}) (CustomRegistrationUpdateData, error) {
	data := new(CustomRegistrationUpdateData)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/manage/registrationdata/"+recordID, body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// DeleteCustomRegistrationData allows you to delete a member from dropDownList.
func DeleteCustomRegistrationData(recordID string) (CustomRegistrationBool, error) {
	data := new(CustomRegistrationBool)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/manage/registrationdata/"+recordID, "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}
