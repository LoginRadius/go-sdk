package lrintegrationtest

import (
	"os"
	"testing"

	lr "github.com/LoginRadius/go-sdk"
	"github.com/LoginRadius/go-sdk/api/mfa"
	"github.com/LoginRadius/go-sdk/lrerror"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

// Tests in this file are skipped by default; they will only run with LoginRadius sites with MFA turned on
// If you enable MFA for your site, tests in authentication_test.go, social_test.go and phoneauthentication_test.go will
// no longer run
func TestPostMFAEmailLogin(t *testing.T) {
	t.SkipNow()

	_, _, _, testEmail, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	testLogin := TestEmailLogin{testEmail, testEmail}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(testLogin)
	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAEmailLogin call: %v", err)
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(testLogin, map[string]string{"emailtemplate": "hello"})

	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call with optional queries: %v", err)
	}
	session, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAEmailLogin call with optional queries: %v", err)
	}
}

func TestPostMFAEmailLoginInvalidBody(t *testing.T) {
	t.SkipNow()

	_, _, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(invalid)
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostMFAEmailLogin should fail with LoginradiusRespondedWithError but did not: %v", res.Body)
	}
}

func TestPostMFAEmailLoginInvalidQuery(t *testing.T) {
	t.SkipNow()

	_, _, _, email, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	user := TestEmailLogin{email, email}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(user, map[string]string{"invalidparam": "value"})
	if err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostMFAEmailLogin should fail with ValidationError but did not :%v, %+v", res.Body, err)
	}
}

func TestPostMFAUsernameLogin(t *testing.T) {
	t.SkipNow()

	_, username, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(
		map[string]string{"username": username, "password": password},
	)
	if err != nil {
		t.Errorf("Error making PostMFAUsernameLogin call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAUsernameLogin call: %v", err)
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(
		map[string]string{"username": username, "password": password},
		map[string]string{"emailtemplate": "hello"},
	)

	if err != nil {
		t.Errorf("Error making PostMFAUsernameLogin call with optional queries: %v", err)
	}
	session, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAUsernameLogin call with optional queries: %v", err)
	}
}

func TestPostMFAUsernameLoginInvalidBody(t *testing.T) {
	t.SkipNow()

	_, _, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(invalid)
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostMFAUsernameLogin should fail with LoginradiusRespondedWithError but did not: %v", res.Body)
	}
}

func TestPostMFAUsernameLoginInvalidQuery(t *testing.T) {
	t.SkipNow()

	_, username, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAUsernameLogin(
		map[string]string{"username": username, "password": password},
		map[string]string{"invalidparam": "value"},
	)
	if err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostMFAUsernameLogin should fail with ValidationError but did not :%v, %+v", res.Body, err)
	}
}

func TestPostMFAPhoneLogin(t *testing.T) {
	t.SkipNow()

	phone, _, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAPhoneLogin(
		map[string]string{"phone": phone, "password": password},
	)
	if err != nil {
		t.Errorf("Error making PostMFAPhoneLogin call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAPhoneLogin call: %v", err)
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAPhoneLogin(
		map[string]string{"phone": phone, "password": password},
		map[string]string{"emailtemplate": "hello"},
	)

	if err != nil {
		t.Errorf("Error making PostMFAPhoneLogin call with optional queries: %v", err)
	}
	session, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostMFAPhoneLogin call with optional queries: %v", err)
	}
}

func TestPostMFAPhoneLoginInvalidBody(t *testing.T) {
	t.SkipNow()

	_, _, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAPhoneLogin(invalid)
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostMFAPhoneLogin should fail with LoginradiusRespondedWithError but did not: %v", res.Body)
	}
}

func TestPostMFAPhoneLoginInvalidQuery(t *testing.T) {
	t.SkipNow()
	phone, _, _, password, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAPhoneLogin(
		map[string]string{"phone": phone, "password": password},
		map[string]string{"invalidparam": "value"},
	)
	if err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostMFAPhoneLogin should fail with ValidationError but did not :%v, %+v", res.Body, err)
	}
}

func TestGetMFAValidateAccessToken(t *testing.T) {
	t.SkipNow()
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAValidateAccessToken()
	if err != nil {
		t.Errorf("Error making call to MFAValidateAccessToken: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["QRCode"].(string) == "" {
		t.Errorf("Error returned from MFAValidateAccessToken: %v", err)
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAValidateAccessToken(map[string]string{"smstemplate2fa": "hello"})
	if err != nil {
		t.Errorf("Error making call to MFAValidateAccessToken with optional query params: %v", err)
	}
	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["QRCode"].(string) == "" {
		t.Errorf("Error returned from MFAValidateAccessToken with optional query params: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// and a Google authenticator added, enter the google authenticator code in this test.
func TestPutMFAValidateGoogleAuthCode(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
		// Set user credentials here
		map[string]string{"email": "", "password": ""},
	)
	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call for PutMFAValidateGoogleAuthCode: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PostMFAEmailLogin call for PutMFAValidateGoogleAuthCode: %v", err)
	}

	code, ok := data["SecondFactorAuthentication"].(map[string]interface{})["SecondFactorAuthenticationToken"].(string)
	if !ok {
		t.Errorf("Returned response from SecondFactorAuthentication does not contain SecondFactorAuthenticationToken")
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAValidateGoogleAuthCode(
		map[string]string{"secondfactorauthenticationtoken": code},
		// Set otp from Google Authenticator here
		map[string]string{"googleauthenticatorcode": "246803"},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAValidateGoogleAuthCode: %v", err)
	}
	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["access_token"].(string) == "" {
		t.Errorf("Error returned from PutMFAValidateGoogleAuthCode: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// this test tests for the ability to submit a valid request to the LoginRadius end point
// and will pass if a ""The OTP code is invalid, please request for a new OTP" error is returned
// from Loginradius
func TestPutMFAValidateOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
		// Set user credentials here
		map[string]string{"email": "blueberries@mailinator.com", "password": "password"},
	)
	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call for PutMFAValidateOTP: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PostMFAEmailLogin call for PutMFAValidateOTP: %v", err)
	}

	code, ok := data["SecondFactorAuthentication"].(map[string]interface{})["SecondFactorAuthenticationToken"].(string)
	if !ok {
		t.Errorf("Returned response from PutMFAValidateOTP does not contain SecondFactorAuthenticationToken")
	}

	_, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAValidateOTP(
		map[string]string{"secondfactorauthenticationtoken": code},
		map[string]string{"otp": "123456"},
	)

	errMsg, _ := lrjson.DynamicUnmarshal(err.(lrerror.Error).OrigErr().Error())

	if errMsg["Description"].(string) != "The OTP code is invalid, please request for a new OTP." {
		t.Errorf("PutMFAValidateOTP was supposed to return invalid OTP error, but did not: %v", errMsg)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid secondfactorauthenticationtoken through completing a mfa login attempt
// set the secondfactorauthenticationtoken and a phone number here
func TestPutMFAUpdatePhoneNumber(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
		// Set user credentials here
		map[string]string{"email": "blueberries@mailinator.com", "password": "password"},
	)
	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call for PutMFAUpdatePhoneNumber: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PostMFAEmailLogin call for PutMFAUpdatePhoneNumber: %v", err)
	}

	code, ok := data["SecondFactorAuthentication"].(map[string]interface{})["SecondFactorAuthenticationToken"].(string)
	if !ok {
		t.Errorf("Returned response from SecondFactorAuthentication does not contain SecondFactorAuthenticationToken")
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAUpdatePhoneNumber(
		// Set user here
		map[string]string{"secondfactorauthenticationtoken": code},
		map[string]string{"phoneno2fa": ""},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAUpdatePhoneNumber: %v", err)
	}
	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PutMFAUpdatePhoneNumber: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token through completing a mfa login attempt
// set the access_token and a phone number here
func TestPutMFAUpdatePhoneNumberByToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "7f875c92-b7fe-4f55-8658-58b24387ed64"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAUpdatePhoneNumberByToken(
		// Set user here
		map[string]string{"phoneno2fa": "16047711536"},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAUpdatePhoneNumber: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["Sid"] == "" {
		t.Errorf("Error returned from PutMFAUpdatePhoneNumber: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token through completing a mfa login attempt
// This test must be run with a user that has not called this end point previously

func TestGetMFABackUpCodeByAccessToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "77aa9464-815c-4dbe-8eec-c6c9e28e43b2"
	_, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFABackUpCodeByAccessToken()
	if err != nil {
		t.Errorf("Error making call to GetMFABackUpCodeByAccessToken: %v", err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token through completing a mfa login attempt
func TestGetMFAResetBackUpCodeByAccessToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "77aa9464-815c-4dbe-8eec-c6c9e28e43b2"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAResetBackUpCodeByAccessToken()
	if err != nil {
		t.Errorf("Error making call to GetMFAResetBackUpCodeByAccessToken: %v", err)
	}

	codes, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := codes["BackUpCodes"].([]interface{})
	if err != nil || !ok {
		t.Errorf("Error returned from GetMFAResetBackUpCodeByAccessToken:%v, %v", err, codes)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token through completing a mfa login attempt
func TestPutMFAValidateBackupCode(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "c3b8130e-e92d-40cc-8153-83da3744aa4b"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAResetBackUpCodeByAccessToken()
	if err != nil {
		t.Errorf("Error making call to GetMFAResetBackUpCodeByAccessToken for PutMFAValidateBackupCode: %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	codes, ok := data["BackUpCodes"].([]interface{})
	if err != nil || !ok {
		t.Errorf("Error returned from GetMFAResetBackUpCodeByAccessToken for PutMFAValidateBackupCode:%v, %v", err, codes)
	}

	// Get secondfactorauthenticationtoken
	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
		// Set user credentials here
		map[string]string{"email": "blueberries@mailinator.com", "password": "password"},
	)
	if err != nil {
		t.Errorf("Error making PostMFAEmailLogin call for PutMFAValidateBackupCode: %v", err)
	}
	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PostMFAEmailLogin call for PutMFAValidateBackupCode: %v", err)
	}

	token, ok := data["SecondFactorAuthentication"].(map[string]interface{})["SecondFactorAuthenticationToken"].(string)
	if !ok {
		t.Errorf("Returned response from PostMFAEmailLogin does not contain SecondFactorAuthenticationToken")
	}

	res, err = mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAValidateBackupCode(
		map[string]string{"secondfactorauthenticationtoken": token},
		map[string]string{"backupcode": codes[0].(string)},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAValidateBackupCode: %v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PutMFAValidateBackupCode: %v, %v", profile, err)
	}

	_, ok = profile["access_token"].(string)
	if !ok {
		t.Errorf("Error returned from PutMFAValidateBackupCode: %v, %v", profile, err)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token through completing a mfa login attempt
// This test must be run with a user that has not called this end point previously
func TestGetMFABackUpCodeByUID(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set uid here
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFABackUpCodeByUID("3ca313699dc8423b9f7c8af9dff9d7f2")
	if err != nil {
		t.Errorf("Error making call to GetMFABackUpCodeByUID: %v", err)
	}

	codes, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := codes["BackUpCodes"].([]interface{})
	if err != nil || !ok {
		t.Errorf("Error returned from GetMFABackUpCodeByUID:%v, %v", err, codes)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
func TestGetMFAResetBackUpCodeByUID(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// Set uid here
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAResetBackUpCodeByUID("3ca313699dc8423b9f7c8af9dff9d7f2")
	if err != nil {
		t.Errorf("Error making call to GetMFAResetBackUpCodeByUID: %v", err)
	}

	codes, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := codes["BackUpCodes"].([]interface{})
	if err != nil || !ok {
		t.Errorf("Error returned from GetMFAResetBackUpCodeByUID:%v, %v", err, codes)
	}
}

// To run this test, comment out t.SkipNow() and set a manually created user with mfa turned on
// and google authenticator configured
func TestDeleteMFAResetGoogleAuthenticatorByUid(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// Set uid here
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).DeleteMFAResetGoogleAuthenticatorByUid("3ca313699dc8423b9f7c8af9dff9d7f2")
	if err != nil {
		t.Errorf("Error making call to DeleteMFAResetGoogleAuthenticatorByUid: %v", err)
	}

	body, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := body["IsDeleted"].(bool)
	if err != nil || !ok {
		t.Errorf("Error returned from DeleteMFAResetGoogleAuthenticatorByUid :%v, %v", err, body)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// and sms authentication configured
func TestDeleteMFAResetSMSAuthenticatorByUid(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// Set uid here
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).DeleteMFAResetSMSAuthenticatorByUid("3ca313699dc8423b9f7c8af9dff9d7f2")
	if err != nil {
		t.Errorf("Error making call to DeleteMFAResetSMSAuthenticatorByUid: %v", err)
	}

	body, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := body["IsDeleted"].(bool)
	if err != nil || !ok {
		t.Errorf("Error returned from DeleteMFAResetSMSAuthenticatorByUid:%v, %v", err, body)
	}
}

// To run this test, comment out t.SkipNow() and set a manually created user with mfa turned on
// and google authenticator configured
func TestDeleteMFAResetGoogleAuthenticatorByToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg, map[string]string{"token": "01a67f99-8ab5-4176-a12b-a5c3d00859b5"})

	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).DeleteMFAResetGoogleAuthenticatorByToken()
	if err != nil {
		t.Errorf("Error making call to DeleteMFAResetGoogleAuthenticatorByToken: %v", err)
	}

	body, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := body["IsDeleted"].(bool)
	if err != nil || !ok {
		t.Errorf("Error returned from DeleteMFAResetGoogleAuthenticatorByToken :%v, %v", err, body)
	}
}

// To run this test, comment out t.SkipNow() and set a manually created user with mfa turned on
// and sms authentication configured
func TestDeleteMFAResetSMSAuthenticatorByToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg, map[string]string{"token": "01a67f99-8ab5-4176-a12b-a5c3d00859b5"})

	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).DeleteMFAResetSMSAuthenticatorByToken()
	if err != nil {
		t.Errorf("Error making call to DeleteMFAResetSMSAuthenticatorByToken: %v", err)
	}

	body, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := body["IsDeleted"].(bool)
	if err != nil || !ok {
		t.Errorf("Error returned from DeleteMFAResetSMSAuthenticatorByToken :%v, %v", err, body)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token and a google authenticator code
func TestPutMFAReauthenticateByGoogleAuthenticator(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "aebbf205-c9b6-458d-9e70-c3dfdabdb2ef"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAReauthenticateByGoogleAuthenticator(
		// set google authenticator code here
		map[string]string{"googleauthenticatorcode": ""},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAReauthenticateByGoogleAuthenticator: %v", err)
	}

	result, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := result["SecondFactorValidationToken"].(string)
	if err != nil || !ok {
		t.Errorf("Error returned from PutMFAReauthenticateByGoogleAuthenticator:%v, %v", err, result)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token and a valid back up code
func TestPutMFAReauthenticateByBackupCode(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "aebbf205-c9b6-458d-9e70-c3dfdabdb2ef"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAReauthenticateByBackupCode(
		// set backup code here
		map[string]string{"backupcode": "53141-b07fb"},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAReauthenticateByBackupCode: %v", err)
	}

	result, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := result["SecondFactorValidationToken"].(string)
	if err != nil || !ok {
		t.Errorf("Error returned from PutMFAReauthenticateByBackupCode:%v, %v", err, result)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token and a valid sms otp
func TestPutMFAReauthenticateByOTP(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "d3d95123-b14c-43d6-99ef-51528051b3bd"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAReauthenticateByOTP(
		// set otp here
		map[string]string{"otp": "53141-b07fb"},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAReauthenticateByOTP: %v", err)
	}

	result, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := result["SecondFactorValidationToken"].(string)
	if err != nil || !ok {
		t.Errorf("Error returned from PutMFAReauthenticateByOTP:%v, %v", err, result)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token and a valid password
func TestPutMFAReauthenticateByPassword(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "d3d95123-b14c-43d6-99ef-51528051b3bd"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAReauthenticateByPassword(
		// set Password here
		map[string]string{"password": "password"},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAReauthenticateByPassword: %v", err)
	}

	result, err := lrjson.DynamicUnmarshal(res.Body)
	_, ok := result["SecondFactorValidationToken"].(string)
	if err != nil || !ok {
		t.Errorf("Error returned from PutMFAReauthenticateByPassword:%v, %v", err, result)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token
func TestGetMFAReAuthenticate(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "96688431-0945-4ed5-9115-733521a13a53"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).GetMFAReAuthenticate()
	if err != nil {
		t.Errorf("Error making call to GetMFAReAuthenticate: %v", err)
	}

	result, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from GetMFAReAuthenticate:%v, %v", err, result)
	}
}

// To run this test, uncomment t.SkipNow() and set a manually created user with mfa turned on
// then obtain a valid access_token
func TestPutMFAUpdateSettings(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// set valid access_token here
	lrclient.Context.Token = "96688431-0945-4ed5-9115-733521a13a53"
	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PutMFAUpdateSettings(
		// manually set otp obtained from sms authenticator here
		map[string]string{"otp": "245212"},
	)
	if err != nil {
		t.Errorf("Error making call to PutMFAUpdateSettings: %v", err)
	}

	result, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from PutMFAUpdateSettings:%v, %v", err, result)
	}
}
