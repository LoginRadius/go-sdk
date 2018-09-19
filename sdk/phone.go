package loginradius

import (
	"os"
	"time"
)

// PhoneLogin is a struct used to contain the login information retrieved when
// the login and verification APIs for phone authentications are called.
type PhoneLogin struct {
	Profile     AuthProfile
	AccessToken string    `json:"access_token"`
	ExpiresIn   time.Time `json:"expires_in"`
	Password    string
}

// PhoneOTP is the struct used to contain various OTP related responses when
// dealing with phone authentication responses
type PhoneOTP struct {
	IsPosted bool `json:"IsPosted"`
	Data     struct {
		AccountSid string `json:"AccountSid"`
		Sid        string `json:"Sid"`
	} `json:"Data"`
}

// PhoneBool is a struct that contains data from responses that contain a
// single boolean JSON attribute from the phone authentication API
type PhoneBool struct {
	IsPosted  bool `json:"IsPosted"`
	IsDeleted bool `json:"IsDeleted"`
	IsExist   bool `json:"IsExist"`
}

// PostPhoneLogin retrieves a copy of the user data based on the Phone.
// The post parameters are the phoneID: string, password: string
// and an optional securityanswer: string
func PostPhoneLogin(loginurl, smstemplate, gRecaptchaResponse string, body interface{}) (PhoneLogin, error) {
	data := new(PhoneLogin)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/login", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("loginurl", loginurl)
	q.Add("smstemplate", smstemplate)
	q.Add("g-recaptcha-response", gRecaptchaResponse)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PostPhoneForgotPasswordByOTP is used to send the OTP to reset the account password.
// The post parameter is phone: string
func PostPhoneForgotPasswordByOTP(smstemplate string, body interface{}) (PhoneOTP, error) {
	data := new(PhoneOTP)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/password/otp", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("smstemplate", smstemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PostPhoneResendVerificationOTP is used to resend a verification OTP to verify a user's Phone Number.
// The user will receive a verification code that they will need to input. The post parameter is phone: string.
func PostPhoneResendVerificationOTP(smstemplate string, body interface{}) (PhoneOTP, error) {
	data := new(PhoneOTP)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/phone/otp", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("smstemplate", smstemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PostPhoneResendVerificationOTPByToken is used to resend a verification OTP to verify a
// user's Phone Number in cases in which an active token already exists. The post parameter is phone: string.
func PostPhoneResendVerificationOTPByToken(smstemplate, authorization string, body interface{}) (PhoneOTP, error) {
	data := new(PhoneOTP)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/phone/otp", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("smstemplate", smstemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PostPhoneUserRegistrationBySMS registers the new users into your Cloud Storage and triggers the phone verification process.
// The post parameter is a user profile object, userprofileobject:string.
func PostPhoneUserRegistrationBySMS(verificationURL, smstemplate, options string, body interface{}) (PhoneBool, error) {
	data := new(PhoneBool)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/register", body)
	if reqErr != nil {
		return *data, reqErr
	}

	sott := GenerateSOTT()
	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("verificationURL", verificationURL)
	q.Add("smstemplate", smstemplate)
	q.Add("options", options)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-Sott", sott)

	err := RunRequest(req, data)
	return *data, err
}

// GetPhoneSendOTP is used to send your phone an OTP.
func GetPhoneSendOTP(phone, smsTemplate string) (PhoneOTP, error) {
	data := new(PhoneOTP)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/login/passwordlesslogin/otp", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("phone", phone)
	q.Add("smstemplate", smsTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// GetPhoneNumberAvailability is used to check the whether the phone number exists or not on your site.
func GetPhoneNumberAvailability(phone string) (PhoneBool, error) {
	data := new(PhoneBool)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/phone", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("phone", phone)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// PutPhoneLoginUsingOTP is used to login using OTP flow.
// The post parameters are phone:string, otp: string, optional smstemplate: string,
// optional securityanswer: string, optional g-recaptcha-response: string,
// optional qq_captcha_ticket: string, optional qq_captcha_randstr: string
func PutPhoneLoginUsingOTP(smsTemplate string, body interface{}) (PhoneLogin, error) {
	data := new(PhoneLogin)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/login/passwordlesslogin/otp/verify", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("smstemplate", smsTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutPhoneNumberUpdate is used to update the phone number of a user.
// The post parameter is a phoneID, phone:string.
func PutPhoneNumberUpdate(smstemplate, authorization string, body interface{}) (PhoneOTP, error) {
	data := new(PhoneOTP)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/phone", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("smstemplate", smstemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PutPhoneResetPasswordByOTP is used to reset the password.
// The post parameters are phone:string, otp: string, password: string and
// optional smstemplate: string and optional resetpasswordemailtemplate: string
func PutPhoneResetPasswordByOTP(body interface{}) (PhoneBool, error) {
	data := new(PhoneBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/password/otp", body)
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

// PutPhoneVerificationByOTP is used to validate the verification code sent to verify a user's phone number.
// The post parameter is the phoneID, phone:string
func PutPhoneVerificationByOTP(otp, smstemplate string, body interface{}) (PhoneLogin, error) {
	data := new(PhoneLogin)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/phone/otp", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("otp", otp)
	q.Add("smstemplate", smstemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutPhoneVerificationByOTPByToken is used to consume the verification code sent to verify a user's phone number.
// Use this call for front-end purposes in cases where the user is already logged in by passing the user's access token.
// The post parameter is the phoneID, phone:string
func PutPhoneVerificationByOTPByToken(otp, smstemplate, authorization string) (PhoneBool, error) {
	data := new(PhoneBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/phone/otp", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("otp", otp)
	q.Add("smstemplate", smstemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PutResetPhoneIDVerification allows you to reset the phone number verification of an end user’s account.
func PutResetPhoneIDVerification(uid string) (PhoneBool, error) {
	data := new(PhoneBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid+"/invalidatephone", "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// DeleteRemovePhoneIDByAccessToken is used to delete the Phone ID on a user's account via the access_token.
func DeleteRemovePhoneIDByAccessToken(authorization string) (PhoneBool, error) {
	data := new(PhoneBool)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/auth/phone", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}
