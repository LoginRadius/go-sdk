// +build !mfa

package loginradius

import (
	"fmt"
	"testing"
)

func setupMFALogin(t *testing.T) (string, string, string, string, func(t *testing.T)) {
	_, _, testuid, testEmail, teardownTestCase := setupAccount(t)
	testLogin := TestEmailLogin{testEmail, testEmail}
	session, err := PostMFAEmailLogin("", "", "", "", testLogin)
	accessToken := session.AccessToken
	multiToken := session.SecondFactorAuthentication.SecondFactorAuthenticationToken
	if err != nil || accessToken == "" {
		t.Errorf("Error logging in")
		fmt.Println(err)
	}
	return multiToken, testuid, testEmail, accessToken, func(t *testing.T) {
		defer teardownTestCase(t)
	}
}

func TestPostMFAEmailLogin(t *testing.T) {
	fmt.Println("Starting test TestPostMFAEmailLogin")
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	testLogin := TestEmailLogin{testEmail, testEmail}
	session, err := PostMFAEmailLogin("", "", "", "", testLogin)
	if err != nil || session.AccessToken == "" {
		t.Errorf("Error logging in")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
func TestPostMFAEmailLoginInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostMFAEmailLogin")
	_, _, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"Invalid"}
	_, err := PostMFAEmailLogin("", "", "", "", invalid)
	if err == nil {
		t.Errorf("Should be error")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostMFAUsernameLogin(t *testing.T) {
	fmt.Println("Starting test TestPostMFAUsernameLogin")
	_, username, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	testLogin := TestUsernameLogin{username, testEmail}
	session, err := PostMFAUsernameLogin("", "", "", "", "", testLogin)
	if err != nil || session.AccessToken == "" {
		t.Errorf("Error logging in")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostMFAPhoneLogin(t *testing.T) {
	fmt.Println("Starting test TestPostMFAPhoneLogin")
	testPhone, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	testLogin := TestPhoneLogin{testPhone, testEmail}
	session, err := PostMFAPhoneLogin("", "", "", "", "", testLogin)
	if err != nil || session.AccessToken == "" {
		t.Errorf("Error logging in")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
