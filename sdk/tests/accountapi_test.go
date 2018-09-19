package loginradius

import (
	"fmt"
	"testing"
	"time"
)

func TestPostManageAccountCreate(t *testing.T) {
	fmt.Println("Starting test TestPostManageAccountCreate")
	PresetLoginRadiusTestEnv()
	time := time.Now()
	timestamp := time.Format("20060102150405")
	timestampEmail := "testemail" + timestamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", timestampEmail}}
	testAccount := TestAccount{true, testEmails, timestampEmail}
	user, err := PostManageAccountCreate(testAccount)
	uid := user.UID
	if err != nil || uid == "" {
		t.Errorf("Error creating account")
		fmt.Println(err)
	}
	_, err2 := DeleteManageAccount(uid)
	if err2 != nil {
		t.Errorf("Error cleaning up account")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestPostManageAccountCreateInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostManageAccountCreateInvalid")
	PresetLoginRadiusTestEnv()
	invalid := InvalidBody{"invalid"}
	_, err := PostManageAccountCreate(invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}
func TestPostManageEmailVerificationToken(t *testing.T) {
	fmt.Println("Starting test TestPostManageEmailVerificationToken")
	_, testEmail, _, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	emailObj := TestEmail{testEmail}
	session, err := PostManageEmailVerificationToken(emailObj)
	if err != nil || session.VerificationToken == "" {
		t.Errorf("Error creating verification token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostManageEmailVerificationTokenInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostManageEmailVerificationTokenInvalid")
	_, _, _, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"invalid"}
	_, err := PostManageEmailVerificationToken(invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}

func TestPostManageForgotPasswordToken(t *testing.T) {
	fmt.Println("Starting test TestPostManageForgotPasswordToken")
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	emailObj := TestEmail{testEmail}
	session, err := PostManageForgotPasswordToken(emailObj)
	if err != nil || session.ForgotToken == "" {
		t.Errorf("Error creating forgot password token")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostManageForgotPasswordTokenInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostManageForgotPasswordTokenInvalid")
	_, _, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"failure"}
	_, err := PostManageForgotPasswordToken(invalid)
	if err == nil {
		t.Errorf("Error should appear.")
	}
	fmt.Println("Test complete")
}

func TestGetManageAccountIdentitiesByEmail(t *testing.T) {
	fmt.Println("Starting test TestGetManageAccountIdentitiesByEmail")
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	_, err := GetManageAccountIdentitiesByEmail(testEmail)
	if err != nil {
		t.Errorf("Error retrieving profiles")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetManageAccessTokenUID(t *testing.T) {
	fmt.Println("Starting test TestGetManageAccessTokenUID")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	session, err := GetManageAccessTokenUID(testuid)
	if err != nil || session.AccessToken == "" {
		t.Errorf("Error retrieving access token associated with uid")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetManageAccountPassword(t *testing.T) {
	fmt.Println("Starting test TestGetManageAccountPassword")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	session, err := GetManageAccountPassword(testuid)
	if err != nil || session.PasswordHash == "" {
		t.Errorf("Error retrieving access token associated with uid")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetManageAccountProfilesByEmail(t *testing.T) {
	fmt.Println("Starting test TestGetManageAccountProfilesByEmail")
	_, _, testuid, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	session, err := GetManageAccountProfilesByEmail(testEmail)
	if err != nil || session.UID != testuid {
		t.Errorf("Error retrieving profile associated with email")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetManageAccountProfilesByUsername(t *testing.T) {
	fmt.Println("Starting test TestGetManageAccountProfilesByUsername")
	_, username, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	session, err := GetManageAccountProfilesByUsername(username)
	if err != nil || session.UID != testuid {
		t.Errorf("Error retrieving profile associated with username")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetManageAccountProfilesByPhoneID(t *testing.T) {
	fmt.Println("Starting test TestGetManageAccountProfilesByPhoneID")
	testphone, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	session, err := GetManageAccountProfilesByPhoneID(testphone)
	if err != nil || session.UID != testuid {
		t.Errorf("Error retrieving profile associated with phone id")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetManageAccountProfilesByUID(t *testing.T) {
	fmt.Println("Starting test TestGetManageAccountProfilesByUID")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	session, err := GetManageAccountProfilesByUID(testuid)
	if err != nil || session.UID != testuid {
		t.Errorf("Error retrieving profile associated with uid")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutManageAccountSetPassword(t *testing.T) {
	fmt.Println("Starting test TestPutManageAccountSetPassword")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	passwordObj := TestPassword{"password"}
	session, err := PutManageAccountSetPassword(testuid, passwordObj)
	if err != nil || session.PasswordHash == "" {
		t.Errorf("Error changing account password")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutManageAccountSetPasswordInvalid(t *testing.T) {
	fmt.Println("Starting test TestPutManageAccountSetPasswordInvalid")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"Invalid"}
	_, err := PutManageAccountSetPassword(testuid, invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}

func TestPutManageAccountUpdate(t *testing.T) {
	fmt.Println("Starting test TestPutManageAccountUpdate")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	firstNameObj := TestFirstName{"First"}
	session, err := PutManageAccountUpdate(testuid, firstNameObj)
	if err != nil || session.UID != testuid || session.FirstName != "First" {
		t.Errorf("Error updating account")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutManageAccountUpdateInvalid(t *testing.T) {
	fmt.Println("Starting test TestPutManageAccountUpdateInvalid")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"Invalid"}
	_, err := PutManageAccountUpdate(testuid, invalid)
	if err == nil {
		t.Errorf("Error changing account password")
	}
	fmt.Println("Test complete")
}

func TestPutManageAccountUpdateSecurityQuestionConfig(t *testing.T) {
	fmt.Println("Starting test TestPutManageAccountUpdateSecurityQuestionConfig")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	securityQuestion := SecurityQuestion{"Answer"}
	securityTest := SecurityQuestionTest{securityQuestion}
	session, err := PutManageAccountUpdateSecurityQuestionConfig(testuid, securityTest)
	if err != nil || session.UID != testuid {
		t.Errorf("Error changing account password")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutManageAccountUpdateSecurityQuestionConfigInvalid(t *testing.T) {
	fmt.Println("Starting test TestPutManageAccountUpdateSecurityQuestionConfigInvalid")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"Invalid"}
	_, err := PutManageAccountUpdateSecurityQuestionConfig(testuid, invalid)
	if err == nil {
		t.Errorf("Error changing account password")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutManageAccountInvalidateVerificationEmail(t *testing.T) {
	fmt.Println("Starting test TestPutManageAccountInvalidateVerificationEmail")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	_, err := PutManageAccountInvalidateVerificationEmail("", "", testuid)
	if err != nil {
		t.Errorf("Error invalidating verification email")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteManageAccountEmailDelete(t *testing.T) {
	fmt.Println("Starting test TestDeleteManageAccountEmail")
	_, _, testuid, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	testEmails := ProfileEmail{TestEmailArr{{"Primary", testEmail}, {"Secondary", "1" + testEmail}}}
	_, err := PutManageAccountUpdate(testuid, testEmails)
	if err != nil {
		t.Errorf("Error adding email")
		fmt.Println(err)
	}
	testEmailObj := TestEmail{testEmail}
	session2, err2 := DeleteManageAccountEmail(testuid, testEmailObj)
	if err2 != nil || session2.UID != testuid {
		t.Errorf("Error deleting email")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestDeleteManageAccountEmailDeleteInvalid(t *testing.T) {
	fmt.Println("Starting test TestDeleteManageAccountEmailDeleteInvalid")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"Invalid"}
	_, err := DeleteManageAccountEmail(testuid, invalid)
	if err == nil {
		t.Errorf("Should be error")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteManageAccount(t *testing.T) {
	fmt.Println("Starting test TestDeleteManageAccount")
	_, _, testuid, _, _ := setupAccount(t)
	_, err := DeleteManageAccount(testuid)
	if err != nil {
		t.Errorf("Error deleting account")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
