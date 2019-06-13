package lrintegrationtest

import (
	"os"
	"strconv"
	"testing"
	"time"

	lr "github.com/LoginRadius/go-sdk"
	lraccount "github.com/LoginRadius/go-sdk/api/account"
	"github.com/LoginRadius/go-sdk/api/phoneauthentication"
	"github.com/LoginRadius/go-sdk/lrerror"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

func TestPostPhoneLogin(t *testing.T) {
	phoneID, _, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneLogin(
		map[string]string{"phone": phoneID, "password": password},
	)
	if err != nil {
		t.Errorf("Error calling PostPhoneLogin: %v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || profile["Uid"] == "" {
		t.Errorf("Error returned from PostPhoneLogin: %v", err)
	}
}

// To run this test comment out t.SkipNow() and set PHONENUMBER in secret.env
// with a valid phone number of a manually created user profile
func TestPostPhoneForgotPasswordByOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneForgotPasswordByOTP(
		map[string]string{"phone": os.Getenv("PHONENUMBER")},
	)
	if err != nil {
		t.Errorf("Error calling PostPhoneForgotPasswordByOTP: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from PostPhoneForgotPasswordByOTP: %v", err)
	}
}

// To run this test comment out t.SkipNow() and set PHONENUMBER in secret.env
// with a valid phone number of a manually created user profile
func TestPostPhoneResendVerificationOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneResendVerificationOTP(
		map[string]string{"phone": os.Getenv("PHONENUMBER")},
	)
	if err != nil {
		t.Errorf("Error calling ostPhoneResendVerificationOTP: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from ostPhoneResendVerificationOTP: %v", err)
	}
}

// To run this test comment out t.SkipNow() and set PHONENUMBER and USERTOKEN in secret.env
// with a valid phone number of a manually created user profile
func TestPostPhoneResendVerificationOTPByToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg, map[string]string{"token": os.Getenv("USERTOKEN")})
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PostPhoneResendVerificationOTPByToken(
		map[string]string{"phone": os.Getenv("PHONENUMBER")},
	)
	if err != nil {
		t.Errorf("Error calling ostPhoneResendVerificationOTP: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from ostPhoneResendVerificationOTP: %v", err)
	}
}

func TestPostPhoneUserRegistrationBySMS(t *testing.T) {

	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	loginradius := phoneauthentication.Loginradius{lrclient}

	testEmail := "lrtest" + strconv.FormatInt(time.Now().Unix(), 10) + "@mailinator.com"
	user := User{}

	res, err := phoneauthentication.Loginradius(loginradius).PostPhoneUserRegistrationBySMS(user)
	if err == nil || err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostPhoneUserRegistrationBySMS Fail: Expected Error %v, instead received res: %+v, received error: %+v", "LoginradiusRespondedWithError", res, err)
	}

	user = User{
		Email: []Email{
			Email{
				Type:  "Primary",
				Value: testEmail,
			},
		},
		Password: "password",
	}

	res, err = phoneauthentication.Loginradius(loginradius).PostPhoneUserRegistrationBySMS(user)
	if res.StatusCode != 200 {
		t.Errorf("PostPhoneUserRegistrationBySMS Success: Expected StatusCode %v, received %v", 200, res)
	}

	res, err = phoneauthentication.Loginradius(loginradius).PostPhoneUserRegistrationBySMS(user)
	if err == nil || err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostPhoneUserRegistrationBySMS Fail: Expected Error %v, instead received res: %+v, received error: %+v", "LoginradiusRespondedWithError", res, err)
	}

	res, err = lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByEmail(map[string]string{"email": testEmail})
	if err != nil {
		t.Errorf("Error retrieving uid of account to clean up: %v.", err)
	}

	profile, _ := lrjson.DynamicUnmarshal(res.Body)
	uid := profile["Uid"].(string)
	_, err = lraccount.Loginradius(lraccount.Loginradius{lrclient}).DeleteManageAccount(uid)
	if err != nil {
		t.Errorf("Error cleaning up account: %v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually attempt a log in with a registered user account
// and fill out the otp received as well as the phone number in this test
func TestPutPhoneLoginUsingOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneLoginUsingOTP(
		// Manually fill phone number of a registered user account and otp received here
		map[string]string{"phone": "", "otp": "871962"},
	)
	if err != nil {
		t.Errorf("Error making PutPhoneLoginUsingOTP call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PutPhoneLoginUsingOTP call: %v", err)
	}
}

// To run this test please comment out t.SkipNow() and manually set
// phone number with a phone number already registered for a user profile
func TestGetPhoneNumberAvailability(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	// Fill phone number here
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).GetPhoneNumberAvailability(
		map[string]string{"phone": ""},
	)
	if err != nil {
		t.Errorf("Error calling GetPhoneNumberAvailability: %v", err)
	}
	available, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !available["IsExist"].(bool) {
		t.Errorf("Error returned from GetPhoneNumberAvailability: %v", err)
	}
}

// To run this test please comment out t.SkipNow() and manually set
// phone number with a phone number already registered for a user profile
func TestGetPhoneSendOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	// Fill phone number here
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).GetPhoneSendOTP(
		map[string]string{"phone": ""},
	)
	if err != nil {
		t.Errorf("Error calling GetPhoneSendOTP: %v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(resp.Body)
	data := posted["Data"].(map[string]interface{})
	if err != nil || data["Sid"].(string) == "" {
		t.Errorf("Error returned from GetPhoneSendOTP: %v", err)
	}
}

// To run this test, comment out t.SkipNow() and manually fill with a valid phone number
func TestPutPhoneNumberUpdate(t *testing.T) {
	t.SkipNow()
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneNumberUpdate(
		// Fill valid phone number here
		map[string]string{"phone": ""},
	)
	if err != nil {
		t.Errorf("Error calling PutPhoneNumberUpdate: %v", err)
	}
	updated, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil && !updated["IsPosted"].(bool) {
		t.Errorf("Error returned from PutPhoneNumberUpdate: %v", err)
	}
}

// To run this test please comment out t.SkipNow() and manually set
// phone number, password, o,tp with a phone number already registered for a user profile
// after making a request to receive a valid otp on this phone number
func TestPutPhoneResetPasswordByOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	// Fill phone number, password, and otp here
	resp, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneResetPasswordByOTP(
		map[string]string{"phone": "", "password": "", "otp": ""},
	)
	if err != nil {
		t.Errorf("Error calling PutPhoneResetPasswordByOTP: %v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned from PutPhoneResetPasswordByOTP: %v", err)
	}
}

// To run this test please comment out t.SkipNow() and manually set
// the phone number of a valid user profile in the test
// This phone number must be unverified.
// after making a request to receive a valid otp on this phone number
func TestPutPhoneVerificationByOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	// Set phone number and otp here
	res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneVerificationByOTP(
		map[string]string{"otp": ""}, map[string]string{"phone": ""},
	)
	if err != nil {
		t.Errorf("Error calling PutPhoneVerificationByOTP: %v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || profile["Uid"] == "" {
		t.Errorf("Error returned from PutPhoneVerificationByOTP: %v", err)
	}
}

// To run this test please comment out t.SkipNow() and manually set
// the access token and phone number of a valid user profile in the test
// This phone number must be unverified.
// after making a request to receive a valid otp on this phone number
func TestPutPhoneVerificationByOTPByToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Set user token here
	lrclient, _ := lr.NewLoginradius(&cfg, map[string]string{"token": ""})
	// Set phone number and otp here
	res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutPhoneVerificationByOTPByToken(
		map[string]string{"otp": ""},
	)
	if err != nil {
		t.Errorf("Error calling PutPhoneVerificationByOTPByToken: %v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || profile["Uid"] == "" {
		t.Errorf("Error returned from PutPhoneVerificationByOTPByToken: %v", err)
	}
}

// To run this test please comment out t.SkipNow() and manually set
// phone number with a phone number already registered for a user profile
func TestPutResetPhoneIDVerification(t *testing.T) {
	_, _, uid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).PutResetPhoneIDVerification(uid)
	if err != nil {
		t.Errorf("Error calling PutResetPhoneIDVerification: %v", err)
	}
	reset, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reset["IsPosted"].(bool) {
		t.Errorf("Error returned from PutResetPhoneIDVerification: %v", err)
	}
}

func TestDeleteRemovePhoneIDByAccessToken(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := phoneauthentication.Loginradius(phoneauthentication.Loginradius{lrclient}).DeleteRemovePhoneIDByAccessToken()
	if err != nil {
		t.Errorf("Error calling DeleteRemovePhoneIDByAccessToken: %v", err)
	}
	deleted, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !deleted["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteRemovePhoneIDByAccessToken: %v", err)
	}
}
