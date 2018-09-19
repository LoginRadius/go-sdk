package loginradius

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestPostAuthAddEmail(t *testing.T) {
	fmt.Println("Starting test TestPostAuthAddEmail")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	time := time.Now()
	timestamp := time.Format("20060102150405")
	timestampEmail := "testemail" + timestamp + "@mailinator.com"
	testAddEmail := TestEmailCreator{timestampEmail, timestamp}
	success, err := PostAuthAddEmail("", "", accessToken, testAddEmail)
	if err != nil || success.IsPosted != true {
		t.Errorf("Error adding email in account")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostAuthAddEmailInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostAuthAddEmailInvalid")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"failure"}
	_, err := PostAuthAddEmail("", "", accessToken, invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}

func TestPostAuthForgotPassword(t *testing.T) {
	fmt.Println("Starting test TestPostAuthForgotPassword")
	_, _, _, testEmail, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	testForgotPass := TestEmail{testEmail}
	success, err := PostAuthForgotPassword("resetpassword.com", "", testForgotPass)
	if err != nil || success.IsPosted != true {
		t.Errorf("Error sending password reset")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostAuthForgotPasswordInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostAuthForgotPasswordInvalid")
	_, _, _, _, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"failure"}
	_, err := PostAuthForgotPassword("resetpassword.com", "", invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}

func TestPostAuthUserRegistrationByEmail(t *testing.T) {
	fmt.Println("Starting test TestPostAuthUserRegistrationByEmail")
	PresetLoginRadiusTestEnv()
	time := time.Now()
	timestamp := time.Format("20060102150405")
	timestampEmail := "testemail" + timestamp + "@mailinator.com"
	testEmails := TestEmailArr{{"P2rimary", timestampEmail}}
	testAccount := TestAccountRegister{testEmails, timestampEmail}
	success, err := PostAuthUserRegistrationByEmail("", "", "", testAccount)
	if err != nil || success.IsPosted != true {
		t.Errorf("Error creating account")
		fmt.Println(err)
	}
	user, err2 := GetManageAccountProfilesByEmail(timestampEmail)
	if err2 != nil {
		t.Errorf("Error cleaning up account")
		fmt.Println(err2)
	}
	uid := user.UID
	_, err3 := DeleteManageAccount(uid)
	if err3 != nil {
		t.Errorf("Error cleaning up account")
		fmt.Println(err3)
	}
	fmt.Println("Test complete")
}

func TestPostAuthUserRegistrationByEmailInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostAuthUserRegistrationByEmailInvalid")
	_, _, _, _, _, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"failure"}
	_, err := PostAuthUserRegistrationByEmail("", "", "", invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}

func TestPostAuthLoginByEmail(t *testing.T) {
	fmt.Println("Starting test TestPostAuthLoginByEmail")
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	testLogin := TestEmailLogin{testEmail, testEmail}
	session, err := PostAuthLoginByEmail("", "", "", "", "", testLogin)
	accessToken := session.AccessToken
	if err != nil || accessToken == "" {
		t.Errorf("Error logging in")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostAuthLoginByEmailInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostAuthLoginByEmailInvalid")
	_, _, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"failure"}
	_, err := PostAuthLoginByEmail("", "", "", "", "", invalid)
	if err == nil {
		t.Errorf("Should be error")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostAuthLoginByUsername(t *testing.T) {
	fmt.Println("Starting test TestPostAuthLoginByUsername")
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	usernameArr := strings.Split(testEmail, "@")
	testLogin := TestUsernameLogin{usernameArr[0], testEmail}
	session, err := PostAuthLoginByUsername("", "", "", "", "", testLogin)
	accessToken := session.AccessToken
	if err != nil || accessToken == "" {
		t.Errorf("Error logging in")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostAuthLoginByUsernameInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostAuthLoginByUsernameInvalid")
	_, _, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"failure"}
	_, err := PostAuthLoginByUsername("", "", "", "", "", invalid)
	if err == nil {
		t.Errorf("Should be error")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAuthCheckEmailAvailability(t *testing.T) {
	fmt.Println("Starting test TestGetAuthCheckEmailAvailability")
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	success, err := GetAuthCheckEmailAvailability(testEmail)
	if (err != nil) || (success.IsExist == false) {
		t.Errorf("Error checking e-mail's availability")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAuthCheckUsernameAvailability(t *testing.T) {
	fmt.Println("Starting test TestGetAuthCheckUsernameAvailability")
	_, username, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	success, err := GetAuthCheckUsernameAvailability(username)
	if (err != nil) || (success.IsExist == false) {
		t.Errorf("Error find username's availability")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAuthReadProfilesByToken(t *testing.T) {
	fmt.Println("Starting test TestGetAuthReadProfilesByToken")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	resp, err := GetAuthReadProfilesByToken(accessToken)
	if err != nil || resp.UID == "" {
		t.Errorf("Error getting data from tokens")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

/*
func TestGetAuthPrivatePolicyAccept(t *testing.T) {
	fmt.Println("Starting test TestGetAuthPrivatePolicyAccept")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	resp, err := GetAuthPrivatePolicyAccept(accessToken)
	if err != nil || resp.UID != "" {
		t.Errorf("Error with privacy policy acceptance")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
*/
func TestGetAuthSendWelcomeEmail(t *testing.T) {
	fmt.Println("Starting test TestGetAuthPrivatePolicyAccept")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	resp, err := GetAuthSendWelcomeEmail("", accessToken)
	if err != nil || resp.IsPosted != true {
		t.Errorf("Error sending welcome email")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAuthSocialIdentity(t *testing.T) {
	fmt.Println("Starting test TestGetAuthSocialIdentity")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	resp, err := GetAuthSocialIdentity(accessToken)
	if err != nil || resp.UID == "" {
		t.Errorf("Error fetching data from provided social identity")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAuthValidateAccessToken(t *testing.T) {
	fmt.Println("Starting test TestGetAuthValidateAccessToken")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	resp, err := GetAuthValidateAccessToken(accessToken)
	if err != nil || resp.AccessToken == "" {
		t.Errorf("Error validating access token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAuthVerifyEmail(t *testing.T) {
	fmt.Println("Starting test TestGetAuthVerifyEmail")
	_, _, verificationToken, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	resp, err := GetAuthVerifyEmail(verificationToken, "", "")
	if err != nil || resp.IsPosted != true {
		t.Errorf("Error verifying the email")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAuthInvalidateAccessToken(t *testing.T) {
	fmt.Println("Starting test TestGetAuthInvalidateAccessToken")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	success, err := GetAuthInvalidateAccessToken(accessToken)
	if err != nil || success.IsPosted != true {
		t.Errorf("Error invalidating access token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetAuthSecurityQuestionByAccessToken(t *testing.T) {
	fmt.Println("Starting test TestGetAuthSecurityQuestionByAccessToken")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := PutAuthUpdateSecurityQuestionByAccessToken(accessToken, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question")
	}
	_, err2 := GetAuthSecurityQuestionByAccessToken(accessToken)
	if err2 != nil {
		t.Errorf("Error getting security questions with email with access token.")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestGetAuthSecurityQuestionByEmail(t *testing.T) {
	fmt.Println("Starting test TestGetAuthSecurityQuestionByEmail")
	_, _, _, testEmail, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := PutAuthUpdateSecurityQuestionByAccessToken(accessToken, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question")
	}
	_, err2 := GetAuthSecurityQuestionByEmail(testEmail)
	if err2 != nil {
		t.Errorf("Error getting security questions with email.")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestGetAuthSecurityQuestionByUsername(t *testing.T) {
	fmt.Println("Starting test TestGetAuthSecurityQuestionByUsername")
	_, username, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := PutAuthUpdateSecurityQuestionByAccessToken(accessToken, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question")
	}
	_, err2 := GetAuthSecurityQuestionByUsername(username)
	if err2 != nil {
		t.Errorf("Error getting security questions with username.")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestGetAuthSecurityQuestionByPhone(t *testing.T) {
	fmt.Println("Starting test TestGetAuthSecurityQuestionByPhone")
	phoneID, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := PutAuthUpdateSecurityQuestionByAccessToken(accessToken, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question")
	}
	_, err2 := GetAuthSecurityQuestionByPhone(phoneID)
	if err2 != nil {
		t.Errorf("Error getting security questions with phone.")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}
func TestPutAuthChangePassword(t *testing.T) {
	fmt.Println("Starting test TestPutAuthChangePassword")
	_, _, _, retEmail, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	passwords := PassChange{retEmail, retEmail + "1"}
	resp, err := PutAuthChangePassword(accessToken, passwords)
	if err != nil || resp.IsPosted != true {
		t.Errorf("Error changing password")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutResendEmailVerification(t *testing.T) {
	fmt.Println("Starting test TestPutResendEmailVerification")
	_, retEmail, _, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	emailRef := TestEmail{retEmail}
	resp, err := PutResendEmailVerification("", "", emailRef)
	if err != nil || resp.IsPosted != true {
		t.Errorf("Error re-sending email verification")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutAuthResetPasswordByResetToken(t *testing.T) {
	fmt.Println("Starting test TestPutAuthResetPasswordByResetToken")
	_, _, _, retEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	resetPass := PasswordReset{getPasswordResetToken(retEmail), retEmail + "1"}
	resp, err := PutAuthResetPasswordByResetToken(resetPass)
	if err != nil || resp.IsPosted != true {
		t.Errorf("Error resetting password")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

/*
func TestPutAuthResetPasswordByOTP(t *testing.T) {

}
*/
func TestPutAuthResetPasswordBySecurityAnswerAndEmail(t *testing.T) {
	fmt.Println("Starting test TestPutAuthResetPasswordBySecurityAnswerAndEmail")
	_, _, _, retEmail, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := PutAuthUpdateSecurityQuestionByAccessToken(accessToken, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question")
	}
	resetInfo := ResetWithEmailSecurity{securityQuestion, retEmail, retEmail + "1", ""}
	resp, err2 := PutAuthResetPasswordBySecurityAnswerAndEmail(resetInfo)
	if err2 != nil || resp.IsPosted != true {
		t.Errorf("Error resetting password")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestPutAuthResetPasswordBySecurityAnswerAndPhone(t *testing.T) {
	fmt.Println("Starting test TestPutAuthResetPasswordBySecurityAnswerAndPhone")
	phoneID, _, _, retEmail, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := PutAuthUpdateSecurityQuestionByAccessToken(accessToken, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question")
	}
	resetInfo := ResetWithPhoneSecurity{securityQuestion, phoneID, retEmail + "1", ""}
	resp, err2 := PutAuthResetPasswordBySecurityAnswerAndPhone(resetInfo)
	if err2 != nil || resp.IsPosted != true {
		t.Errorf("Error resetting password")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestPutAuthResetPasswordBySecurityAnswerAndUsername(t *testing.T) {
	fmt.Println("Starting test TestPutAuthResetPasswordBySecurityAnswerAndUsername")
	_, username, _, retEmail, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := PutAuthUpdateSecurityQuestionByAccessToken(accessToken, securityTest)
	if err != nil {
		t.Errorf("Error setting up security question")
	}
	resetInfo := ResetWithUsernameSecurity{securityQuestion, username, retEmail + "1", ""}
	resp, err2 := PutAuthResetPasswordBySecurityAnswerAndUsername(resetInfo)
	if err2 != nil || resp.IsPosted != true {
		t.Errorf("Error resetting password")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestPutAuthSetOrChangeUsername(t *testing.T) {
	fmt.Println("Starting test TestPutAuthSetOrChangeUsername")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	newName := TestUsername{"NewUsername5667567"}
	resp, err := PutAuthSetOrChangeUsername(accessToken, newName)
	if err != nil || resp.IsPosted != true {
		t.Errorf("Error changing username")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutAuthUpdateProfileByToken(t *testing.T) {
	fmt.Println("Starting test TestPutAuthUpdateProfileByToken")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	newName := TestUsername{"NewUsername5667567"}
	resp, err := PutAuthUpdateProfileByToken("", "", "", accessToken, newName)
	if err != nil || resp.IsPosted != true {
		t.Errorf("Error changing profile information.")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutAuthUpdateSecurityQuestionByAccessToken(t *testing.T) {
	fmt.Println("Starting test TestPutAuthUpdateSecurityQuestionByAccessToken")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	_, err := PutAuthUpdateSecurityQuestionByAccessToken(accessToken, securityTest)
	if err != nil {
		t.Errorf("Error sending account deletion email")
		fmt.Println(err)
	}
	fmt.Println("Test complete")

}
func TestDeleteAuthDeleteAccountEmailConfirmation(t *testing.T) {
	fmt.Println("Starting test TestDeleteAuthDeleteAccountEmailConfirmation")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	resp, err := DeleteAuthDeleteAccountEmailConfirmation("", "", accessToken)
	if err != nil || resp.IsDeleteRequestAccepted != true {
		t.Errorf("Error sending account deletion email")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteAuthRemoveEmail(t *testing.T) {
	fmt.Println("Starting test DeleteAuthRemoveEmail")
	_, _, _, testEmail, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	removeEmail := TestEmail{testEmail}
	resp, err := DeleteAuthRemoveEmail(accessToken, removeEmail)
	if err != nil || resp.IsDeleted != true {
		t.Errorf("Error removing email.")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteAuthUnlinkSocialIdentities(t *testing.T) {
	// Comment out SkipNow if Candidate Token is set
	t.SkipNow()
	fmt.Println("Starting test TestDeleteAuthUnlinkSocialIdentities")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	candidateToken := CandidateToken{os.Getenv("CANDIDATETOKEN")}
	session, err := PutAuthLinkSocialIdentities(accessToken, candidateToken)
	if err != nil || session.IsPosted == false {
		t.Errorf("Error linking account")
		fmt.Println(err)
	}
	id, err2 := GetAuthReadProfilesByToken(accessToken)
	if err2 != nil {
		t.Errorf("Account is not linked")
		fmt.Println(err2)
	}
	// Generate map object to store data
	m := make(map[string]interface{})
	// If Identity exists after social link, parse it, otherwise throw error
	if len(id.Identities) > 0 {
		err4 := json.Unmarshal(id.Identities[0], &m)
		if err4 != nil {
			t.Errorf("Error parsing the Identities JSON")
			fmt.Println(err4)
		}
		// Access map objects and use type assertion to assign them into the Provider struct
		provider := Provider{m["Provider"].(string), m["ID"].(string)}
		resp, err3 := DeleteAuthUnlinkSocialIdentities(accessToken, provider)
		if err3 != nil || resp.IsDeleted != true {
			t.Errorf("Error unlinking account")
			fmt.Println(err3)
		}
	} else {
		t.Errorf("Candidate Token invalid or expired, please replace the candidate token and try again.")
	}
	fmt.Println("Test complete")
}
