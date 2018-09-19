package loginradius

import (
	"os"
	"time"
)

// SmartLoginBool contains data from responses that return a single boolean attribute
type SmartLoginBool struct {
	IsPosted   bool `json:"IsPosted"`
	IsVerified bool `json:"IsVerified"`
}

// SmartLogin contains the login information received by Smart Login Ping
type SmartLogin struct {
	Profile     AuthProfile `json:"Profile"`
	AccessToken string      `json:"access_token"`
	ExpiresIn   time.Time   `json:"expires_in"`
}

// GetSmartLoginByEmail sends a Smart Login link to the user's Email Id.
func GetSmartLoginByEmail(email, clientGUID, smartLoginEmailTemplate,
	welcomeEmailTemplate, redirectURL string) (SmartLoginBool, error) {
	data := new(SmartLoginBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/auth/login/smartlogin", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("email", email)
	q.Add("clientguid", clientGUID)
	q.Add("smartloginemailtemplate", smartLoginEmailTemplate)
	q.Add("welcomeemailtemplate", welcomeEmailTemplate)
	q.Add("redirecturl", redirectURL)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSmartLoginByUsername sends a Smart Login link to the user's Email Id.
func GetSmartLoginByUsername(username, clientGUID, smartLoginEmailTemplate,
	welcomeEmailTemplate, redirectURL string) (SmartLoginBool, error) {
	data := new(SmartLoginBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/auth/login/smartlogin", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("username", username)
	q.Add("clientguid", clientGUID)
	q.Add("smartloginemailtemplate", smartLoginEmailTemplate)
	q.Add("welcomeemailtemplate", welcomeEmailTemplate)
	q.Add("redirecturl", redirectURL)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSmartLoginPing is used to check if the Smart Login link has been clicked or not.
func GetSmartLoginPing(clientGUID string) (SmartLogin, error) {
	data := new(SmartLogin)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/auth/login/smartlogin/ping", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("clientguid", clientGUID)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSmartLoginVerifyToken verifies the provided token for Smart Login.
func GetSmartLoginVerifyToken(verificationToken, welcomeEmailTemplate string) (SmartLoginBool, error) {
	data := new(SmartLoginBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/auth/email/smartlogin", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("verificationtoken", verificationToken)
	q.Add("welcomeemailtemplate", welcomeEmailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}
