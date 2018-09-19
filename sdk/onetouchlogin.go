package loginradius

import (
	"os"
	"time"
)

// OneTouchLoginData struct is used to contain data from the login by email and phone APIs
type OneTouchLoginData struct {
	IsPosted bool `json:"IsPosted"`
	Data     struct {
		AccountSid string `json:"AccountSid"`
		Sid        string `json:"Sid"`
	} `json:"Data"`
}

// OneTouchLogin contains the response from the OTP verification API
type OneTouchLogin struct {
	Profile     AuthProfile `json:"Profile"`
	AccessToken string      `json:"access_token"`
	ExpiresIn   time.Time   `json:"expires_in"`
}

// GetOneTouchLoginByEmail is used to send a link to a specified email for a frictionless login/registration
func GetOneTouchLoginByEmail(email, name, clientGUID, redirectURL, oneTouchLoginEmailTemplate,
	welcomeEmailTemplate string) (OneTouchLoginData, error) {
	data := new(OneTouchLoginData)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/auth/onetouchlogin/email", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("email", email)
	q.Add("name", name)
	q.Add("clientguid", clientGUID)
	q.Add("onetouchloginemailtemplate", oneTouchLoginEmailTemplate)
	q.Add("redirecturl", redirectURL)
	q.Add("welcomeemailtemplate", welcomeEmailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetOneTouchLoginByPhone is used to send one time password to a given phone number for a frictionless login/registration.
func GetOneTouchLoginByPhone(phone, name, smsTemplate string) (OneTouchLoginData, error) {
	data := new(OneTouchLoginData)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/auth/onetouchlogin/phone", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("phone", phone)
	q.Add("name", name)
	q.Add("smstemplate", smsTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// PutOneTouchOTPVerification is used to verify the otp for One Touch Login.
// Body parameters include the phone: string.
func PutOneTouchOTPVerification(otp, smsTemplate string, body interface{}) (OneTouchLogin, error) {
	data := new(OneTouchLogin)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/auth/onetouchlogin/phone/verify", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("otp", otp)
	q.Add("smstemplate", smsTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}
