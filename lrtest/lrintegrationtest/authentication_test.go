package lrintegrationtest

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	lr "github.com/LoginRadius/go-sdk"
	lraccount "github.com/LoginRadius/go-sdk/api/account"
	lrauthentication "github.com/LoginRadius/go-sdk/api/authentication"
	lrbody "github.com/LoginRadius/go-sdk/lrbody"
	"github.com/LoginRadius/go-sdk/lrerror"
	"github.com/LoginRadius/go-sdk/lrjson"
	"github.com/LoginRadius/go-sdk/lrstruct"
)

type Email struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

type User struct {
	Email    []Email `json:"Email"`
	Password string  `json:"Password"`
}

func TestPostAuthUserRegistrationByEmail(t *testing.T) {

	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	loginradius := lrauthentication.Loginradius{lrclient}

	testEmail := "lrtest" + strconv.FormatInt(time.Now().Unix(), 10) + "@mailinator.com"
	user := lrbody.RegistrationUser{}

	res, err := lrauthentication.Loginradius(loginradius).PostAuthUserRegistrationByEmail(user)
	if err == nil || err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostAuthUserRegistrationByEmail Fail: Expected Error %v, instead received res: %+v, received error: %+v", "LoginradiusRespondedWithError", res, err)
	}

	user = lrbody.RegistrationUser{
		Email: []lrbody.AuthEmail{
			lrbody.AuthEmail{
				Type:  "Primary",
				Value: testEmail,
			},
		},
		Password: "password",
	}

	res, err = lrauthentication.Loginradius(loginradius).PostAuthUserRegistrationByEmail(user)
	if res.StatusCode != 200 {
		t.Errorf("PostAuthUserRegistrationByEmail Success: Expected StatusCode %v, received %v", 200, res)
	}

	res, err = lrauthentication.Loginradius(loginradius).PostAuthUserRegistrationByEmail(user)
	if err == nil || err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostAuthUserRegistrationByEmail Fail: Expected Error %v, instead received res: %+v, received error: %+v", "LoginradiusRespondedWithError", res, err)
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

func TestPostAuthAddEmail(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)

	testEmail := "lrtest" + strconv.FormatInt(time.Now().Unix(), 10) + "@mailinator.com"
	testAddEmail := TestEmailCreator{testEmail, "secondary"}

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{Client: lrclient}).PostAuthAddEmail(testAddEmail)

	if err != nil {
		t.Errorf("Error making PostAuthAddEmail call: %v", err)
	}
	success, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !success["IsPosted"].(bool) {
		t.Errorf("Error returned from PostAuthAddEmail call: %v", err)
	}
}

func TestPostAuthAddEmailInvalidBody(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthAddEmail(invalid)

	if err == nil {
		t.Errorf("PostAuthAddEmail should fail but did not :%v", res.Body)
	}
}

func TestPostAuthAddEmailInvalidQueries(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	generated := "lrtest" + strconv.FormatInt(time.Now().Unix(), 10) + "@mailinator.com"
	email := TestEmailCreator{generated, "secondary"}

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthAddEmail(email, map[string]string{"wrongquery": "value"})

	if err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostAuthAddEmail should fail with ValidationError but did not :%v, %+v", res.Body, err)
	}
}

func TestPostAuthForgotPassword(t *testing.T) {
	_, _, _, testEmail, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	email := TestEmail{testEmail}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthForgotPassword(email, map[string]string{"resetpasswordurl": "resetpassword.com"})
	if err != nil {
		t.Errorf("Error making PostAuthForgotPassword call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from PostAuthForgotPassword call: %v", err)
	}
}

func TestPostAuthForgotPasswordInvalidQuery(t *testing.T) {
	_, _, _, testEmail, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	email := TestEmail{testEmail}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthForgotPassword(email, map[string]string{"wrongqueryname": "www.example.com"})
	if err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostAuthForgotPassword should fail with ValidationError but did not :%v, %+v", res.Body, err)
	}
}

func TestPostAuthForgotPasswordInvalid(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthForgotPassword(invalid, map[string]string{"resetpasswordurl": "www.example.com"})
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostAuthForgotPassword should fail with LoginradiusRespondedWithError but did not :%v, %+v", res.Body, err)
	}
}

func TestPostAuthLoginByEmail(t *testing.T) {
	_, _, _, testEmail, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	testLogin := TestEmailLogin{testEmail, testEmail}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthLoginByEmail(testLogin)
	if err != nil {
		t.Errorf("Error making PostAuthLoginByEmail call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostAuthLoginByEmail call: %v", err)
	}

	res, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthLoginByEmail(testLogin, map[string]string{"emailtemplate": "hello"})

	if err != nil {
		t.Errorf("Error making PostAuthLoginByEmail call with optional queries: %v", err)
	}
	session, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostAuthLoginByEmail call with optional queries: %v", err)
	}
}

func TestPostAuthLoginByEmailInvalidBody(t *testing.T) {
	_, _, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthLoginByEmail(invalid)
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("PostAuthLoginByEmail should fail with LoginradiusRespondedWithError but did not: %v", res.Body)
	}
}

func TestPostAuthLoginByEmailInvalidQuery(t *testing.T) {
	_, _, _, email, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	user := TestEmailLogin{email, email}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthLoginByEmail(user, map[string]string{"invalidparam": "value"})
	if err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("PostAuthLoginByEmail should fail with ValidationError but did not :%v, %+v", res.Body, err)
	}
}

func TestPostAuthLoginByUsername(t *testing.T) {
	_, userName, _, email, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)

	body := struct {
		Username string
		Password string
	}{
		userName,
		email, // uses generated email as password
	}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthLoginByUsername(body)
	if err != nil {
		t.Errorf("Error making PostAuthLoginByUsername call: %v", err)
	}
	session, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || session["access_token"].(string) == "" {
		t.Errorf("Error returned from PostAuthLoginByUsername call: %v", err)
	}
}

func TestPostAuthLoginByUsernameInvalid(t *testing.T) {
	_, _, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := struct{ foo string }{"bar"}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthLoginByUsername(invalid)
	if err == nil {
		t.Errorf("PostAuthLoginByUsername should return error but did not: %v", res.Body)
	}
}

func TestGetAuthCheckEmailAvailability(t *testing.T) {
	_, _, _, testEmail, loginradius, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{loginradius}).
		GetAuthCheckEmailAvailability(map[string]string{"email": testEmail})
	if err != nil {
		t.Errorf("Error making GetAuthCheckEmailAvailability call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsExist"].(bool) {
		t.Errorf("Error returned from GetAuthCheckEmailAvailability call: %v", err)
	}
}

func TestGetAuthCheckUsernameAvailability(t *testing.T) {
	_, username, _, _, loginradius, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{loginradius}).GetAuthCheckUsernameAvailability(map[string]string{"username": username})

	if err != nil {
		t.Errorf("Error making GetAuthCheckUsernameAvailability call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsExist"].(bool) {
		t.Errorf("Error returned from GetAuthCheckUsernameAvailability call: %v", err)
	}
}

func TestGetAuthReadProfilesByToken(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthReadProfilesByToken()
	if err != nil {
		t.Errorf("Error making GetAuthReadProfilesByToken call: %v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || profile["Uid"].(string) == "" {
		t.Errorf("Error returned from GetAuthReadProfilesByToken call: %v", err)
	}
}

// // Test will fail if the feature Privacy Policy Versioning is not enabled through the dashboard
// // To run test, comment out t.SkipNow()
func TestGetAuthPrivatePolicyAccept(t *testing.T) {
	t.SkipNow()
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthPrivatePolicyAccept()
	if err != nil {
		t.Errorf("Error making GetAuthPrivatePolicyAccept call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["Uid"].(string) == "" {
		t.Errorf("Error returned from GetAuthPrivatePolicyAccept call: %v", err)
	}
}

func TestGetAuthSendWelcomeEmail(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSendWelcomeEmail(map[string]string{"welcomeemailtemplate": "hello"})

	if err != nil {
		t.Errorf("Error making GetAuthSendWelcomeEmail call with optional argument: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from GetAuthSendWelcomeEmail call: %v", err)
	}

	res, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSendWelcomeEmail(map[string]string{"wrong argument": "hello"})

	if err == nil {
		t.Errorf("Optional argument validation was supposed to return error, did not return error")
	}

	res, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSendWelcomeEmail()
	if err != nil {
		t.Errorf("Error making GetAuthSendWelcomeEmail call with no optional argument: %v", err)
	}
}

func TestGetAuthSocialIdentity(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSocialIdentity()

	if err != nil {
		t.Errorf("Error making GetAuthSocialIdentity call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["Uid"].(string) == "" {
		t.Errorf("Error returned from GetAuthSocialIdentity call: %v", err)
	}
}

func TestGetAuthSocialIdentityFail(t *testing.T) {

	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	//initialize lrclient without access token
	lrclient, _ := lr.NewLoginradius(&cfg)

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSocialIdentity()
	if err.(lrerror.Error).Code() != "MissingTokenErr" {
		t.Errorf("TestGetAuthSocialIdentityFail Should fail with MissingTokenErr but instead got: %v, %v", res, err)
	}
}

func TestGetAuthValidateAccessToken(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthValidateAccessToken()
	if err != nil {
		t.Errorf("Error making GetAuthValidateAccessToken call, %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || data["access_token"].(string) == "" {
		t.Errorf("Error returned from GetAuthValidateAccessToken call %v", err)
	}
}

func TestGetAuthVerifyEmail(t *testing.T) {
	_, _, verificationToken, loginradius, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{loginradius}).GetAuthVerifyEmail(map[string]string{"verificationtoken": verificationToken})
	if err != nil {
		t.Errorf("Error making TestAuthVerifyEmail call, %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from TestAuthVerifyEmail call, %v", err)
	}
}

func TestGetAuthInvalidateAccessToken(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthInvalidateAccessToken()
	if err != nil {
		t.Errorf("Error making GetAuthInvalidateAccessToken call, %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from GetAuthInvalidateAccessToken call, %v", err)
	}
}

// Comment out skipnow and manually set a delete token to run test
// Delete token must be retrieved from email inbox after calling DeleteAuthDeleteAccountEmailConfirmation with
// an account that was manually set up
func TestGetAuthDeleteAccount(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}
	lrclient, err := lr.NewLoginradius(&cfg)

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthDeleteAccount(map[string]string{"deletetoken": "064102295d22491aae48aaddb0e818c0"})

	if err != nil {
		t.Errorf("Error making GetAuthDeleteAccount call: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from GetAuthDeleteAccount call: %v", err)
	}
}

// // Will return error unless security question feature is enabled
// // Follow instructions in this document: https://docs.lrauthentication.com/api/v2/dashboard/platform-security/password-policy
func TestGetAuthSecurityQuestionByAccessToken(t *testing.T) {
	_, _, uid, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSecurityQuestionByAccessToken()
	if err != nil {
		t.Errorf("Error making GetAuthSecurityQuestionByAccessToken call: %v", err)
	}
	question := lrstruct.AuthSecurityQuestion{}
	err = json.Unmarshal([]byte(res.Body), &question)
	if err != nil || (question[0].QuestionID == "") {
		t.Errorf("Error returned from GetAuthSecurityQuestionByUsername call: %v", err)
	}
}

func TestGetAuthSecurityQuestionByEmail(t *testing.T) {
	_, _, uid, email, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSecurityQuestionByEmail(map[string]string{"email": email})

	if err != nil {
		t.Errorf("Error making GetAuthSecurityQuestionByUsername call: %v", err)
	}
	question := lrstruct.AuthSecurityQuestion{}
	err = json.Unmarshal([]byte(res.Body), &question)
	if err != nil || (question[0].QuestionID == "") {
		t.Errorf("Error returned from GetAuthSecurityQuestionByUsername call: %v", err)
	}
}

func TestGetAuthSecurityQuestionByUsername(t *testing.T) {
	_, username, uid, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSecurityQuestionByUsername(map[string]string{"username": username})
	if err != nil {
		t.Errorf("Error making GetAuthSecurityQuestionByUsername call: %v", err)
	}
	question := lrstruct.AuthSecurityQuestion{}
	err = json.Unmarshal([]byte(res.Body), &question)
	if err != nil || (question[0].QuestionID == "") {
		t.Errorf("Error returned from GetAuthSecurityQuestionByUsername call: %v", err)
	}
}

func TestGetAuthSecurityQuestionByPhone(t *testing.T) {
	phone, _, uid, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthSecurityQuestionByPhone(map[string]string{"phone": phone})
	if err != nil {
		t.Errorf("Error making GetAuthSecurityQuestionByPhone call: %v", err)
	}
	question := lrstruct.AuthSecurityQuestion{}
	err = json.Unmarshal([]byte(res.Body), &question)
	if err != nil || (question[0].QuestionID == "") {
		t.Errorf("Error returned from GetAuthSecurityQuestionByPhone call: %v", err)
	}
}

func TestPutAuthChangePassword(t *testing.T) {
	_, _, _, email, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	passwords := PassChange{email, email + "1"}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthChangePassword(passwords)
	if err != nil {
		t.Errorf("Error calling PutAuthChangePassword: %+v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned from PutAuthChangePassword: %+v", err)
	}
}

func TestPutResendEmailVerification(t *testing.T) {
	_, retEmail, _, lrclient, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	emailRef := TestEmail{retEmail}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutResendEmailVerification(emailRef)
	if err != nil {
		t.Errorf("Error calling PutResendEmailVerification: %v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned for PutResendEmailVerification: %v", err)
	}

	res, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutResendEmailVerification(emailRef, map[string]string{"emailtemplate": "hello"})
	if err != nil {
		t.Errorf("Error calling PutResendEmailVerification: %v", err)
	}
	posted, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned for PutResendEmailVerification: %v", err)
	}
}

func TestPutResendEmailVerificationInvalid(t *testing.T) {
	_, retEmail, _, lrclient, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	emailRef := TestEmail{retEmail}
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutResendEmailVerification(map[string]string{"invalidquery": "hello"}, emailRef)
	if err == nil || err.(lrerror.Error).Code() != "ValidationError" {
		t.Errorf("Should fail with ValidationError, but got instead:%+v, %v", res, err)
	}
}

func TestPutAuthResetPasswordByResetToken(t *testing.T) {
	_, _, _, email, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)

	resetEmail := TestEmail{email}
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PostManageForgotPasswordToken(resetEmail)
	if err != nil {
		t.Errorf(
			"Error calling PostManageForgotPasswordToken for PutAuthResetPasswordByResetToken: %v",
			err,
		)
	}
	data, _ := lrjson.DynamicUnmarshal(response.Body)
	req := PasswordReset{data["ForgotToken"].(string), email + "1"}
	response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthResetPasswordByResetToken(req)
	if err != nil {
		t.Errorf("Error calling PutAuthResetPasswordByResetToken: %v", err)
	}
	data, err = lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from PutAuthResetPasswordByResetToken: %+v", err)
	}
}

func TestPutAuthResetPasswordByOTP(t *testing.T) {
	t.SkipNow()
}

func TestPutAuthResetPasswordBySecurityAnswerAndEmail(t *testing.T) {
	_, _, uid, email, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)

	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}

	request := ResetWithEmailSecurity{securityQuestion, email, email + "1", ""}
	response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthResetPasswordBySecurityAnswerAndEmail(request)
	if err != nil {
		t.Errorf("Error making call to PutAuthResetPasswordBySecurityAnswerAndEmail: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from call to PutAuthResetPasswordBySecurityAnswerAndEmail: %+v", err)
	}
}

func TestPutAuthResetPasswordBySecurityAnswerAndUsername(t *testing.T) {
	_, username, uid, email, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)

	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	response, err := lraccount.Loginradius(lraccount.Loginradius{lrclient}).PutManageAccountUpdateSecurityQuestionConfig(uid, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question: %v", err)
	}

	request := ResetWithUsernameSecurity{securityQuestion, username, email + "1", ""}
	response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthResetPasswordBySecurityAnswerAndUsername(request)
	if err != nil {
		t.Errorf("Error making call to PutAuthResetPasswordBySecurityAnswerAndUsername: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !data["IsPosted"].(bool) {
		t.Errorf("Error returned from PutAuthResetPasswordBySecurityAnswerAndUsername: %+v", err)
	}
}

func TestPutAuthSetOrChangeUsername(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	newName := TestUsername{"newusername"}
	_, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthSetOrChangeUsername(newName)
	if err != nil {
		t.Errorf("Error making call to PutAuthSetOrChangeUsername: %+v", err)
	}
	response, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthReadProfilesByToken()
	if err != nil {
		t.Errorf("Error making call to GetAuthReadProfilesByToken for PutAuthSetOrChangeUsername: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil {
		t.Errorf("Error returned from GetAuthReadProfilesByToken for PutAuthSetOrChangeUsername: %+v", err)
	}
	if data["UserName"].(string) != "newusername" {
		t.Errorf("PutAuthSetOrChangeUsername failed, expected username NewUserName, but instead got: %v", data["UserName"].(string))
	}
}

func TestPutAuthUpdateProfileByToken(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	request := TestUsername{"newname"}
	_, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthUpdateProfileByToken(request)
	if err != nil {
		t.Errorf("Error making call to PutAuthUpdateProfileByToken: %+v", err)
	}
	response, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthReadProfilesByToken()
	if err != nil {
		t.Errorf("Error making call to GetAuthReadProfilesByToken for PutAuthUpdateProfileByToken: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil {
		t.Errorf("Error returned from GetAuthReadProfilesByToken for PutAuthUpdateProfileByToken: %+v", err)
	}
	if data["UserName"].(string) != "newname" {
		t.Errorf("PutAuthSetOrChangeUsername failed, expected username NewUserName, but instead got: %v", data["UserName"].(string))
	}
}

func TestPutAuthUpdateSecurityQuestionByAccessToken(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PutAuthUpdateSecurityQuestionByAccessToken(securityTest)
	if err != nil {
		t.Errorf("Error making PutAuthUpdateSecurityQuestionByAccessToken call: %v", err)
	}
}

func TestDeleteAuthDeleteAccountEmailConfirmation(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	resp, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).DeleteAuthDeleteAccountEmailConfirmation()
	if err != nil {
		t.Errorf("Error making call to DeleteAuthDeleteAccountEmailConfirmation: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !data["IsDeleteRequestAccepted"].(bool) {
		t.Errorf("Error returned from DeleteAuthDeleteAccountEmailConfirmation: %+v", err)
	}
}

func TestDeleteAuthRemoveEmail(t *testing.T) {
	_, _, _, testEmail, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	removeEmail := TestEmail{testEmail}
	resp, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).DeleteAuthRemoveEmail(removeEmail)
	if err != nil {
		t.Errorf("Error making call to DeleteAuthRemoveEmail: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from call to DeleteAuthRemoveEmail: %+v", err)
	}
}

// // To run this test, comment out t.SkipNow(), and configure secret.env with valid user access token
// // Pre-create the user used for this test and link an account of a social provider; configure the
// // string of this social provider in the secret.env with lower case names
// // e.g.PROVIDER=google, PROVIDER=facebook
func TestDeleteAuthUnlinkSocialIdentities(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	accessToken := os.Getenv("USERTOKEN")

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg, map[string]string{"token": accessToken})

	response, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthReadProfilesByToken()
	if err != nil {
		t.Errorf("Error making call to GetAuthReadProfilesByToken: %+v", err)
	}

	data, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil {
		t.Errorf("Error parsing response from GetAuthReadProfilesByToken: %+v", err)
	}
	identities, ok := data["Identities"].([]interface{})
	if !ok {
		fmt.Println("Identities returned is null, not array")
		return
	}

	var id string
	providerstr := os.Getenv("PROVIDER")
	for _, v := range identities {
		asserted := v.(map[string]interface{})
		if asserted["Provider"] == providerstr {
			id = asserted["ID"].(string)
		}
	}

	provider := Provider{providerstr, id}

	response, err = lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).DeleteAuthUnlinkSocialIdentities(provider)
	if err != nil {
		t.Errorf("Error making call to DeleteAuthUnlinkSocialIdentities: %+v", err)
	}

	deleted, err := lrjson.DynamicUnmarshal(response.Body)
	if err != nil || !deleted["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteAuthUnlinkSocialIdentities: %+v", err)
	}
}
func TestGetPasswordlessLoginByEmail(t *testing.T) {
	_, _, _, email, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetPasswordlessLoginByEmail(
		map[string]string{"email": email},
	)
	if err != nil {
		t.Errorf("Error making call to GetPasswordlessLoginByEmail: %+v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned from GetPasswordlessLoginByEmail: %+v", err)
	}
}

func TestGetPasswordlessLoginByUsername(t *testing.T) {
	_, username, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetPasswordlessLoginByUsername(
		map[string]string{"username": username},
	)
	if err != nil {
		t.Errorf("Error making call to GetPasswordlessLoginByUsername: %+v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned from GetPasswordlessLoginByUsername: %+v", err)
	}
}

//Comment out t.SkipNow() and manually set verificationtoken to run test
//verificationtoken needs to be retrieved from email inbox after
//calling GetPasswordlessLoginByEmail or ByUsername APIs
func TestGetPasswordlessLoginVerification(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)
	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetPasswordlessLoginVerification(
		map[string]string{"verificationtoken": "7108eccb667940dcbcf6a6c31685f96a"},
	)
	if err != nil {
		t.Errorf("Error making call to GetPasswordlessLoginVerification: %+v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	profile := data["Profile"].(map[string]interface{})
	if err != nil || profile["Uid"].(string) == "" {
		t.Errorf("Error returned from GetPasswordlessLoginVerification call: %v", err)
	}
}
