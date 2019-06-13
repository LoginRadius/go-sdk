package lrintegrationtest

import (
	"os"
	"strconv"
	"testing"
	"time"

	lr "github.com/LoginRadius/go-sdk"
	lraccount "github.com/LoginRadius/go-sdk/api/account"
	lrauthentication "github.com/LoginRadius/go-sdk/api/authentication"
	lrbody "github.com/LoginRadius/go-sdk/lrbody"
	"github.com/LoginRadius/go-sdk/lrerror"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

func TestGetManageAccountProfilesByEmail(t *testing.T) {
	_, _, testuid, testEmail, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByEmail(
		map[string]string{"email": testEmail},
	)
	if err != nil {
		t.Errorf("Error making call to GetManageAccountProfilesByEmail: %+v", err)
	}
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	uid := session["Uid"].(string)
	if err != nil || uid != testuid {
		t.Errorf("Error returned from GetManageAccountProfilesByEmail: %v", err)
	}
}

func TestGetManageAccountProfilesByUsername(t *testing.T) {
	_, username, testuid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByUsername(map[string]string{"username": username})
	if err != nil {
		t.Errorf("Error making call to GetManageAccountProfilesByUsername: %+v", err)
	}
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	uid := session["Uid"].(string)
	if err != nil || uid != testuid {
		t.Errorf("Error returned from GetManageAccountProfilesByUsername: %v", err)
	}
}

func TestGetManageAccountProfilesByPhoneID(t *testing.T) {
	phoneid, _, testuid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByPhoneID(map[string]string{"phone": phoneid})
	if err != nil {
		t.Errorf("Error making call to GetManageAccountProfilesByPhoneID: %+v", err)
	}
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	uid := session["Uid"].(string)
	if err != nil || uid != testuid {
		t.Errorf("Error returned from GetManageAccountProfilesByPhoneID: %v", err)
	}
}

func TestGetManageAccountIdentitiesByEmail(t *testing.T) {
	_, _, testuid, testEmail, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountIdentitiesByEmail(map[string]string{"email": testEmail})
	if err != nil {
		t.Errorf("Error making call to GetManageAccountIdentitiesByEmail: %v", err)
	}
	body, _ := lrjson.DynamicUnmarshal(response.Body)
	profiles := body["Data"].([]interface{})
	profile := profiles[0].(map[string]interface{})
	uid := profile["Uid"].(string)
	if err != nil || uid != testuid {
		t.Errorf("Error returned from GetManageAccountIdentitiesByEmail: %v", err)
	}
}

func TestGetManageAccountIdentitiesByUid(t *testing.T) {
	_, _, testuid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByUid(testuid)
	if err != nil {
		t.Errorf("Error making call to GetManageAccountProfilesByUid: %+v", err)
	}
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	uid := session["Uid"].(string)
	if err != nil || uid != testuid {
		t.Errorf("Error returned from GetManageAccountProfilesByUid: %v", err)
	}
}

func TestDeleteManageAccount(t *testing.T) {
	_, _, testuid, _, lrclient, _ := setupAccount(t)
	_, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).DeleteManageAccount(testuid)
	if err != nil {
		t.Errorf("Error deleting account: %v", err)
	}
}

func TestPostManageAccountCreate(t *testing.T) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	loginradius := lrauthentication.Loginradius{lrclient}

	testEmail := "lrtest" + strconv.FormatInt(time.Now().Unix(), 10) + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", testEmail}}
	testAccount := TestAccount{true, testEmails, testEmail}

	response, err := lraccount.Loginradius(loginradius).PostManageAccountCreate(testAccount)
	if err != nil {
		t.Errorf("Error calling PostManageAccountCreate: %v", err)
	}
	user, err := lrjson.DynamicUnmarshal(response.Body)
	uid := user["Uid"].(string)
	if err != nil || uid == "" {
		t.Errorf("Error returned from PostManageAccountCreate: %v", err)
	}
	_, err = lraccount.Loginradius(lraccount.Loginradius{lrclient}).DeleteManageAccount(uid)
	if err != nil {
		t.Errorf("Error cleaning up account: %v", err)
	}
}

func TestPostManageEmailVerificationToken(t *testing.T) {
	_, testEmail, _, loginradius, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	emailObj := TestEmail{testEmail}
	response, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).PostManageEmailVerificationToken(emailObj)
	if err != nil {
		t.Errorf(" Error making call to PostManageEmailVerificationToken: %v", err)
	}
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || session["VerificationToken"].(string) == "" {
		t.Errorf("Error returned from PostManageEmailVerificationToken: %v", err)
	}
}

// To run this test, comment out t.SkipNow() and set up SecurityQuestion struct
// in structs.go with querstion id
func TestPutManageAccountUpdateSecurityQuestionConfig(t *testing.T) {
	t.SkipNow()
	_, _, testuid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdateSecurityQuestionConfig(testuid, securityTest)
	if err != nil {
		t.Errorf("Error making PutManageAccountUpdateSecurityQuestionConfig call")
	}
	profile, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || profile["Uid"].(string) != testuid {
		t.Errorf("Error returned from PutManageAccountUpdateSecurityQuestionConfig: %v", err)
	}
}
func TestPostManageForgotPasswordToken(t *testing.T) {
	_, _, _, testEmail, loginradius, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	email := TestEmail{testEmail}
	response, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).PostManageForgotPasswordToken(email)
	if err != nil {
		t.Errorf("Error making call to PostManageForgotPasswordToken: %+v", err)
	}
	session, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || session["ForgotToken"].(string) == "" {
		t.Errorf("Error creating forgot password token: %v", err)
	}

	response, err = lraccount.Loginradius(lraccount.Loginradius{loginradius}).PostManageForgotPasswordToken(email, map[string]string{"sendemail": "true"})
	if err != nil {
		t.Errorf("Error making call to PostManageForgotPasswordToken: %+v", err)
	}
	session, err = lrjson.DynamicUnmarshal(response.Body)
	if err != nil || session["ForgotToken"].(string) == "" {
		t.Errorf("Error creating forgot password token: %v", err)
	}
}

func TestPostManageForgotPasswordTokenInvalid(t *testing.T) {
	_, _, _, testEmail, loginradius, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	email := TestEmail{testEmail}
	response, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).PostManageForgotPasswordToken(email, map[string]string{"invalidparam": "value"})
	if err == nil || err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostManageForgotPasswordToken with invalid param was supposed to return ValidationError, but instead got: %+v, %+v", response, err)
	}
}

func TestGetManageAccessTokenUID(t *testing.T) {
	_, _, uid, _, loginradius, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).GetManageAccessTokenUID(map[string]string{"uid": uid})
	if err != nil {
		t.Errorf("Error making call to GetManageAccessTokenUID: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || data["access_token"].(string) == "" {
		t.Errorf("Error returned from GetManageAccessTokenUID: %v", err)
	}
}

func TestGetManageAccountPassword(t *testing.T) {
	_, _, uid, _, loginradius, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).GetManageAccountPassword(uid)
	if err != nil {
		t.Errorf("Error making call to GetManageAccountPassword: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || data["PasswordHash"].(string) == "" {
		t.Errorf("Error returned from GetManageAccountPassword: %v", err)
	}
}

func TestPutManageAccountSetPassword(t *testing.T) {
	_, _, testuid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	body := lrbody.AccountPassword{"password"}
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountSetPassword(testuid, body)
	if err != nil {
		t.Errorf("Error making PutManageAccountbetPassword call: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || data["PasswordHash"].(string) == "" {
		t.Errorf("Error returned from PutManageAccountSetPassword: %v", err)
	}
}

func TestPutManageAccountUpdate(t *testing.T) {
	_, _, uid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	_, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdate(
		uid,
		[]byte(`{"Username":"newname"}`),
	)
	if err != nil {
		t.Errorf("Error making call to PutManageAccountUpdate: %+v", err)
	}

	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).GetManageAccountProfilesByUid(uid)
	if err != nil {
		t.Errorf("Error making call to GetManageAccountProfilesByUid for PutManageAccountUpdate: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil {
		t.Errorf("Error returned from GetManageAccountProfilesByUid for PutManageAccountUpdate: %+v", err)
	}
	if data["UserName"].(string) != "newname" {
		t.Errorf("PutAuthSetOrChangeUsername failed, expected username newname, but instead got: %v", data["UserName"].(string))
	}
}

func TestPutManageAccountInvalidateVerificationEmail(t *testing.T) {
	_, _, testuid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountInvalidateVerificationEmail(testuid)
	if err != nil {
		t.Errorf("Error making PutManageAccountInvalidateVerificationEmail call: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from PutManageAccountInvalidateVerificationEmail: %v", err)
	}
}

func TestDeleteManageAccountEmailDelete(t *testing.T) {
	_, _, uid, email, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	emails := ProfileEmail{TestEmailArr{{"Primary", email}, {"Secondary", "1" + email}}}
	_, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdate(uid, emails)
	if err != nil {
		t.Errorf("Error calling PutManageAccountUpdate for DeleteManageAccountEmailDelete: %+v", err)
	}
	toDelete := TestEmail{email}
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).DeleteManageAccountEmail(uid, toDelete)
	if err != nil {
		t.Errorf("Error calling DeleteManageAccountEmailDelete: %+v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || profile["Uid"].(string) != uid {
		t.Errorf("Error returned from DeleteManageAccountEmailDelete: %v", err)
	}
}
