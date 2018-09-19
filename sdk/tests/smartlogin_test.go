package loginradius

import (
	"fmt"
	"testing"
)

func TestGetSmartLoginByEmail(t *testing.T) {
	fmt.Println("Starting test TestGetSmartLoginByEmail")
	_, _, _, testEmail, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	_, err := GetSmartLoginByEmail(testEmail, createClientGUID(), "", "", "")
	if err != nil {
		t.Errorf("Smart login failed")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetSmartLoginByUsername(t *testing.T) {
	fmt.Println("Starting test TestGetSmartLoginByUsername")
	_, username, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	_, err := GetSmartLoginByUsername(username, createClientGUID(), "", "", "")
	if err != nil {
		t.Errorf("Smart login failed")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

/*
func TestGetSmartLoginPing(t *testing.T) {
	fmt.Println("Starting test TestGetSmartLoginPing")
	_, username, _, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	clientGUID := createClientGUID()
	_, err := GetSmartLoginByUsername(username, clientGUID, "", "", "")
	if err != nil {
		t.Errorf("Smart login failed")
		fmt.Println(err)
	}
	_, err2 := GetSmartLoginPing(clientGUID)
	if err2 != nil {
		t.Errorf("Smart login ping failed")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestGetSmartLoginVerifyToken(t *testing.T) {
	_, _, verificationToken, teardownTestCase := setupEmailVerificationAccount(t)
	defer teardownTestCase(t)
	fmt.Println("Starting test TestGetSmartLoginVerifyToken")
	_, err := GetSmartLoginVerifyToken(verificationToken, "")
	if err != nil {
		t.Errorf("Smart login email verify token check failed")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}*/
