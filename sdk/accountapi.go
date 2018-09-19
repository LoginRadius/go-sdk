package loginradius

import (
	"os"
	"time"
)

// AccountProfilesData is the struct used to hold the response from the GetProfiles API
type AccountProfilesData struct {
	Data []AuthProfile
}

// AccountAccessToken holds data about the access token retrieved from Get AccessToken from UID
type AccountAccessToken struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   time.Time `json:"expires_in"`
}

// AccountPassword holds data about the Password Hash
type AccountPassword struct {
	PasswordHash string `json:"PasswordHash"`
}

// AccountTokens holds data about the token generation POST methods
type AccountTokens struct {
	VerificationToken string   `json:"VerificationToken"`
	ForgotToken       string   `json:"ForgotToken"`
	IdentityProviders []string `json:"IdentityProviders"`
}

// AccountBool holds data methods that return a boolean
type AccountBool struct {
	IsPosted  bool `json:"IsPosted"`
	IsDeleted bool `json:"IsDeleted"`
}

// PostManageAccountCreate is used to create an account in Cloud Storage.
// This API bypasses the normal email verification process and manually creates the user.
// In order to use this API, you need to format a JSON request body with all of the mandatory fields
// Required post parameters are email object and password:string. Rest are optional profile parameters.
func PostManageAccountCreate(body interface{}) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/manage/account", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PostManageEmailVerificationToken Returns an Email Verification token.
// Post parameter is the email: string
func PostManageEmailVerificationToken(body interface{}) (AccountTokens, error) {
	data := new(AccountTokens)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/manage/account/verify/token", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PostManageForgotPasswordToken returns a forgot password token. Note: If you have the
// UserName workflow enabled, you may replace the 'email' parameter with 'username'.
// Post parameter is either the username: string or the email: string
func PostManageForgotPasswordToken(body interface{}) (AccountTokens, error) {
	data := new(AccountTokens)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/identity/v2/manage/account/forgot/token", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// GetManageAccountIdentitiesByEmail is used to retrieve all of the identities (UID and Profiles),
// associated with a specified email in Cloud Storage.
// Note: This is intended for specific workflows where an email may be associated to multiple UIDs.
func GetManageAccountIdentitiesByEmail(email string) (AccountProfilesData, error) {
	data := new(AccountProfilesData)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account/identities", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// GetManageAccessTokenUID is used to get LoginRadius access token based on UID.
func GetManageAccessTokenUID(uid string) (AccountAccessToken, error) {
	data := new(AccountAccessToken)
	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/identity/v2/manage/account/access_token", "")
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

// GetManageAccountPassword is used to retrieve the hashed password of a specified account in Cloud Storage.
func GetManageAccountPassword(uid string) (AccountPassword, error) {
	data := new(AccountPassword)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid+"/password", "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// GetManageAccountProfilesByEmail is used to retrieve all of the profile data,
// associated with the specified account by email in Cloud Storage.
func GetManageAccountProfilesByEmail(email string) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// GetManageAccountProfilesByUsername is used to retrieve all of the profile data
// associated with the specified account by user name in Cloud Storage.
func GetManageAccountProfilesByUsername(username string) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("username", username)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// GetManageAccountProfilesByPhoneID is used to retrieve all of the profile data,
// associated with the account by phone number in Cloud Storage.
func GetManageAccountProfilesByPhoneID(phone string) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("phone", phone)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// GetManageAccountProfilesByUID is used to retrieve all of the profile data,
// associated with the account by UID in Cloud Storage.
func GetManageAccountProfilesByUID(uid string) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid, "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PutManageAccountSetPassword is used to set the password of an account in Cloud Storage.
// Post parameter is the new password: string
func PutManageAccountSetPassword(uid string, body interface{}) (AccountPassword, error) {
	data := new(AccountPassword)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid+"/password", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PutManageAccountUpdate is used to update the information of existing accounts in your Cloud Storage.
// See our Advanced API Usage section  for more capabilities.
// Post parameters is the profile data that needs to be updated.
func PutManageAccountUpdate(uid string, body interface{}) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid, body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PutManageAccountUpdateSecurityQuestionConfig is used to update security questions configuration on an existing account.
// Post parameter is the security question answer object.
func PutManageAccountUpdateSecurityQuestionConfig(uid string, body interface{}) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid, body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// PutManageAccountInvalidateVerificationEmail is used to invalidate the Email Verification status on an account.
func PutManageAccountInvalidateVerificationEmail(verificationURL, emailTemplate, uid string) (AccountBool, error) {
	data := new(AccountBool)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid+"/invalidateemail", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("verificationurl", verificationURL)
	q.Add("emailtemplate", emailTemplate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// DeleteManageAccountEmail is used to remove emails from a user Account.
// Post parameters are the emails being removed.
func DeleteManageAccountEmail(uid string, body interface{}) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid+"/email", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/json")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}

// DeleteManageAccount is used to delete the Users account and allows them to re-register for a new account.
func DeleteManageAccount(uid string) (AccountBool, error) {
	data := new(AccountBool)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN")+"/identity/v2/manage/account/"+uid, "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))

	err := RunRequest(req, data)
	return *data, err
}
