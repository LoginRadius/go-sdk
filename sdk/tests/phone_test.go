package loginradius

import (
	"fmt"
	"testing"
	"time"
)

func unverifyPhone(uid string) {
	falsePhoneIDVerified := UndoPhoneVerify{false}
	PutManageAccountUpdate(uid, falsePhoneIDVerified)
}

func TestPostPhoneLogin(t *testing.T) {
	fmt.Println("Starting test TestPostPhoneLogin")
	phoneID, _, _, password, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	phoneLogin := TestPhoneLogin{phoneID, password}
	_, err := PostPhoneLogin("", "", "", phoneLogin)
	if err != nil {
		t.Errorf("Error logging in with phone")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostPhoneForgotPasswordByOTP(t *testing.T) {
	fmt.Println("Starting test TestPostPhoneForgotPasswordByOTP")
	phoneID, _, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	phone := TestPhone{phoneID}
	_, err := PostPhoneForgotPasswordByOTP("SMS TEMPLATE", phone)
	if err != nil {
		t.Errorf("Error sending forgot password OTP")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostPhoneResendVerificationOTP(t *testing.T) {
	fmt.Println("Starting test TestPostPhoneResendVerificationOTP")
	phoneID, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	unverifyPhone(testuid)
	phone := TestPhone{phoneID}
	_, err := PostPhoneResendVerificationOTP("SMS TEMPLATE", phone)
	if err != nil {
		t.Errorf("Error sending verification OTP")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostPhoneResendVerificationOTPByToken(t *testing.T) {
	fmt.Println("Starting test TestPostPhoneResendVerificationOTPByToken")
	phoneID, _, testuid, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	unverifyPhone(testuid)
	phone := TestPhone{phoneID}
	_, err := PostPhoneResendVerificationOTPByToken("SMS TEMPLATE", accessToken, phone)
	if err != nil {
		t.Errorf("Error sending verification by OTP")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostPhoneUserRegistrationBySMS(t *testing.T) {
	PresetLoginRadiusTestEnv()
	fmt.Println("Starting test TestPostPhoneUserRegistrationBySMS")
	time := time.Now()
	timestamp := time.Format("20060102150405")
	timestampEmail := "testemail" + timestamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", timestampEmail}}
	phoneAccount := PhoneRegister{testEmails, "+12016768872", "password"}
	session, err := PostPhoneUserRegistrationBySMS("", "", "", phoneAccount)
	if err != nil && session.IsPosted != true {
		t.Errorf("Error registering phone number")
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

func TestGetPhoneNumberAvailability(t *testing.T) {
	fmt.Println("Starting test TestGetPhoneNumberAvailability")
	phoneID, _, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	_, err := GetPhoneNumberAvailability(phoneID)
	if err != nil {
		t.Errorf("Error checking phone number availability")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutPhoneNumberUpdate(t *testing.T) {
	fmt.Println("Starting test TestPutPhoneNumberUpdate")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	phone := TestPhone{"+12016768874"}
	session, err := PutPhoneNumberUpdate("", accessToken, phone)
	if err != nil && session.IsPosted != true {
		t.Errorf("Error updating phone number")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutResetPhoneIDVerification(t *testing.T) {
	fmt.Println("Starting test TestPutResetPhoneIDVerification")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	_, err := PutResetPhoneIDVerification(testuid)
	if err != nil {
		t.Errorf("Error resetting verification")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteRemovePhoneIDByAccessToken(t *testing.T) {
	fmt.Println("Starting test TestDeleteRemovePhoneIDByAccessToken")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	_, err := DeleteRemovePhoneIDByAccessToken(accessToken)
	if err != nil {
		t.Errorf("Error removing phone ID")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
