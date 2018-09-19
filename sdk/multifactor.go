package loginradius

import (
	"os"
	"time"
)

// MFALogin is a struct to contain data returned from MFA Validate Backup code
type MFALogin struct {
	Profile     AuthProfile `json:"Profile"`
	AccessToken string      `json:"access_token"`
	ExpiresIn   time.Time   `json:"expires_in"`
}

// MFAPost is a struct used to contain the data from MFA Post requests
type MFAPost struct {
	SecondFactorAuthentication struct {
		SecondFactorAuthenticationToken string    `json:"SecondFactorAuthenticationToken"`
		ExpireIn                        time.Time `json:"ExpireIn"`
		QRCode                          string    `json:"QRCode"`
		ManualEntryCode                 string    `json:"ManualEntryCode"`
		IsGoogleAuthenticatorVerified   bool      `json:"IsGoogleAuthenticatorVerified"`
		IsOTPAuthenticatorVerified      bool      `json:"IsOTPAuthenticatorVerified"`
		OTPPhoneNo                      string    `json:"OTPPhoneNo"`
		OTPStatus                       MFAPhone  `json:"OTPStatus"`
	} `json:"SecondFactorAuthentication"`
	Profile     AuthProfile `json:"Profile"`
	AccessToken string      `json:"access_token"`
	ExpiresIn   string      `json:"expires_in"`
}

// MFAValidate is a struct used to contain data returned from MFA Validate Access Token
type MFAValidate struct {
	SecondFactorAuthenticationToken string    `json:"SecondFactorAuthenticationToken"`
	QRCode                          string    `json:"QRCode"`
	ManualEntryCode                 string    `json:"ManualEntryCode"`
	IsGoogleAuthenticatorVerified   bool      `json:"IsGoogleAuthenticatorVerified"`
	IsOTPAuthenticatorVerified      bool      `json:"IsOTPAuthenticatorVerified"`
	OTPPhoneNo                      string    `json:"OTPPhoneNo"`
	OTPStatus                       string    `json:"OTPStatus"`
	ExpireIn                        time.Time `json:"ExpireIn"`
}

// MFABackUpCodes is a struct used to contain data from responses that return backup codes
type MFABackUpCodes struct {
	BackUpCodes []string `json:"BackUpCodes"`
}

// MFAPhone is a struct used to hold the data from MFA Update Phone Number
type MFAPhone struct {
	AccountSid string `json:"AccountSid"`
	Sid        string `json:"Sid"`
}

// MFAIsDeleted holds the boolean for the DELETE responses
type MFAIsDeleted struct {
	IsDeleted bool `json:"IsDeleted"`
}

// PostMFAEmailLogin can be used to login by emailid on a Multi-factor authentication enabled LoginRadius site.
// The post parameters are the email: string and password: string
func PostMFAEmailLogin(loginURL, verificationURL, emailTemplate, smstemplate2fa string, body interface{}) (MFAPost, error) {
	data := new(MFAPost)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("loginurl", loginURL)
	q.Add("verificationurl", verificationURL)
	q.Add("emailtemplate", emailTemplate)
	q.Add("smstemplate2fa", smstemplate2fa)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PostMFAUsernameLogin can be used to login by username on a Multi factor authentication enabled LoginRadius site.
// The post parameters are the username: string and password: string
func PostMFAUsernameLogin(loginURL, verificationURL, emailTemplate,
	smsTemplate, smstemplate2fa string, body interface{}) (MFAPost, error) {
	data := new(MFAPost)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("loginurl", loginURL)
	q.Add("verificationurl", verificationURL)
	q.Add("emailtemplate", emailTemplate)
	q.Add("smsTemplate", smsTemplate)
	q.Add("smstemplate2fa", smstemplate2fa)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PostMFAPhoneLogin is used to log in by phone on a Multi-factor authentication enabled LoginRadius site.
// Multi-Factor Authentication: can be enabled in two ways:
// Required case: Multi-factor authentication is enabled forcefully for all users
// Optional case: the user can enable Multi-factor authentication on the profile.
// The post parameters are the phone: string and password: string
func PostMFAPhoneLogin(loginURL, verificationURL, emailTemplate,
	smsTemplate, smstemplate2fa string, body interface{}) (MFAPost, error) {
	data := new(MFAPost)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("loginurl", loginURL)
	q.Add("verificationurl", verificationURL)
	q.Add("emailtemplate", emailTemplate)
	q.Add("smsTemplate", smsTemplate)
	q.Add("smstemplate2fa", smstemplate2fa)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// GetMFAValidateAccessToken is used to configure the Multi-factor authentication
// after login by using the access_token when MFA is set as optional on the LoginRadius site.
func GetMFAValidateAccessToken(smstemplate2fa, authorization string) (MFAValidate, error) {
	data := new(MFAValidate)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/account/2fa", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("smstemplate2fa", smstemplate2fa)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetMFABackUpCodeByAccessToken is used to get a set of backup codes via access_token to
// allow the user login on a site that has Multi-factor Authentication enabled in the event
// that the user does not have a secondary factor available.
// We generate 10 codes, each code can only be consumed once.
// If any user attempts to go over the number of invalid login attempts
// configured in the Dashboard then the account gets blocked automatically
func GetMFABackUpCodeByAccessToken(authorization string) (MFABackUpCodes, error) {
	data := new(MFABackUpCodes)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/account/2fa/backupcode", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetMFAResetBackUpCodeByAccessToken is used to reset the backup codes on a given account via the access_token.
// This API call will generate 10 new codes, each code can only be consumed once.
func GetMFAResetBackUpCodeByAccessToken(authorization string) (MFABackUpCodes, error) {
	data := new(MFABackUpCodes)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/auth/account/2fa/backupcode/reset", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// GetMFABackUpCodeByUID is used to get a set of backup codes to allow the user login on a site
// that has Multi-factor authentication enabled in the event that the user does not have a secondary factor available.
// We generate 10 codes, each code can only be consumed once. If any user attempts to go over the
// number of invalid login attempts configured in the Dashboard then the account gets blocked automatically
func GetMFABackUpCodeByUID(uid string) (MFABackUpCodes, error) {
	data := new(MFABackUpCodes)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account/2fa/backupcode", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("uid", uid)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// GetMFAResetBackUpCodeByUID is used to reset the backup codes on a given account via the UID.
// This API call will generate 10 new codes, each code can only be consumed once.
func GetMFAResetBackUpCodeByUID(uid string) (MFABackUpCodes, error) {
	data := new(MFABackUpCodes)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account/2fa/backupcode/reset", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("uid", uid)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PutMFAValidateBackupCode  is used to validate the backup code provided by the user and if valid, we return an access_token
// allowing the user to login in cases where Multi-factor authentication (MFA) is enabled and the secondary factor is unavailable.
// When a user initially downloads the Backup codes, We generate 10 codes, each code can only be consumed once. if any user
// attempts to go over the number of invalid login attempts configured in the Dashboard then the account gets blocked automatically
// Post parameter is the backupcode: string for logging in
func PutMFAValidateBackupCode(secondFactorAuthenticationToken string, body interface{}) (MFALogin, error) {
	data := new(MFALogin)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa/verification/backupcode", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("secondfactorauthenticationtoken", secondFactorAuthenticationToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutMFAValidateOTP is used to login via Multi-factor authentication by passing the One Time Password received via SMS.
// Post parameters are otp: string and optional securityanswer:string, optional qq_captcha_ticket: string, optional
// qq_captcha_randstr: string and optional g-recaptcha-response: string
func PutMFAValidateOTP(secondFactorAuthenticationToken, smstemplate2fa string, body interface{}) (MFALogin, error) {
	data := new(MFALogin)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa/verification/otp", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("secondfactorauthenticationtoken", secondFactorAuthenticationToken)
	q.Add("smstemplate2fa", smstemplate2fa)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutMFAValidateGoogleAuthCode is used to login via Multi-factor-authentication by passing the google authenticator code.
// Post parameter is googleauthenticatorcode: string
func PutMFAValidateGoogleAuthCode(secondFactorAuthenticationToken, smstemplate2fa string, body interface{}) (MFALogin, error) {
	data := new(MFALogin)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa/verification/googleauthenticatorcode", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("secondfactorauthenticationtoken", secondFactorAuthenticationToken)
	q.Add("smstemplate2fa", smstemplate2fa)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutMFAUpdatePhoneNumber is update the user's phone number associated with MFA.
// Post parameter is phone number needing upgrade, phoneno2fa : string
func PutMFAUpdatePhoneNumber(secondFactorAuthenticationToken, smstemplate2fa string, body interface{}) (MFAPhone, error) {
	data := new(MFAPhone)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/login/2fa", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("secondfactorauthenticationtoken", secondFactorAuthenticationToken)
	q.Add("smstemplate2fa", smstemplate2fa)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutMFAUpdatePhoneNumberByToken is used to update the Multi-factor authentication phone number
// by sending the verification OTP to the provided phone number
// Post parameter is phone number needing upgrade, phoneno2fa : string
func PutMFAUpdatePhoneNumberByToken(smstemplate2fa, authorization string, body interface{}) (MFAPhone, error) {
	data := new(MFAPhone)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/account/2fa", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("smstemplate2fa", smstemplate2fa)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PutMFAUpdateByToken is used to trigger the Multi-factor authentication settings after login for secure actions.
// Post parameter is the google authenticator code.
func PutMFAUpdateByToken(smsTemplate, authorization string, body interface{}) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/account/2fa/verification/googleauthenticatorcode", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("smstemplate", smsTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	err := RunRequest(req, data)
	return *data, err
}

// PutMFAUpdateSettings is used to trigger the Multi-factor authentication settings after login for secure actions.
// Post parameter is the OTP.
func PutMFAUpdateSettings(authorization string, body interface{}) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/auth/account/2fa/verification/otp", body)
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

// DeleteMFAResetGoogleAuthenticatorByToken Resets the Google Authenticator configurations on a given account via the access_token.
// Post parameter is authenticator: string; pass 'googleauthenticator' to remove the Authenticator.
func DeleteMFAResetGoogleAuthenticatorByToken(authorization string, body interface{}) (MFAIsDeleted, error) {
	data := new(MFAIsDeleted)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/auth/account/2fa/authenticator", body)
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

// DeleteMFAResetSMSAuthenticatorByToken Resets the SMS Authenticator configurations on a given account via the access_token.
// Post parameter is authenticator: string; pass 'otpauthenticator' to remove the Authenticator.
func DeleteMFAResetSMSAuthenticatorByToken(authorization string, body interface{}) (MFAIsDeleted, error) {
	data := new(MFAIsDeleted)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/auth/account/2fa/authenticator", body)
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

// DeleteMFAResetGoogleAuthenticatorByUID resets the Google Authenticator configurations on a given account via the UID.
// Post parameter is authenticator: string; pass 'googleauthenticator' to remove the Authenticator.
func DeleteMFAResetGoogleAuthenticatorByUID(uid string, body interface{}) (MFAIsDeleted, error) {
	data := new(MFAIsDeleted)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/manage/account/2fa/authenticator", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("uid", uid)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// DeleteMFAResetSMSAuthenticatorByUID resets the Google Authenticator configurations on a given account via the UID.
// Post parameter is authenticator: string; pass 'otpauthenticator' to remove the Authenticator.
func DeleteMFAResetSMSAuthenticatorByUID(uid string, body interface{}) (MFAIsDeleted, error) {
	data := new(MFAIsDeleted)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/manage/account/2fa/authenticator", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("uid", uid)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}
