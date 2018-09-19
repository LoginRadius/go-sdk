package loginradius

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestRequestReturns200(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		profile := Success{"Alex", []string{"snowboarding", "programming"}}
		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}))

	defer ts.Close()

	fmt.Println("Starting TestRequestReturns200 Test")
	req, reqErr := CreateRequest("GET", ts.URL, "")
	if reqErr != nil {
		fmt.Println(reqErr)
		t.Errorf("Shouldn't be error.")
	}
	data := new(Success)
	err := RunRequest(req, data)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Shouldn't be error.")
	}
	fmt.Println("Test complete")
}

func TestRequestReturns500(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		resp := Error{"Error 500", "TestRequestReturns500 result"}
		js, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}))
	defer ts.Close()

	fmt.Println("Starting TestRequestReturns500 Test")
	req, reqErr := CreateRequest("GET", ts.URL, "")
	if reqErr != nil {
		fmt.Println(reqErr)
		t.Errorf("Shouldn't be error.")
	}
	data := new(Error)
	err := RunRequest(req, data)
	if err == nil {
		t.Errorf("Should be error.")
	}
	fmt.Println("Test complete")
}

func TestRequestReturns404(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		resp := Error{"Error 404", "TestRequestReturns404 result"}
		js, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}))
	fmt.Println("Starting TestRequestReturns404 Test")
	req, reqErr := CreateRequest("GET", ts.URL, "")
	if reqErr != nil {
		fmt.Println(reqErr)
		t.Errorf("Shouldn't be error.")
	}
	data := new(Error)
	err := RunRequest(req, data)
	if err == nil {
		t.Errorf("Should be error.")
	}
	fmt.Println("Test complete")
}

func setupLogin(t *testing.T) (string, string, string, string, string, func(t *testing.T)) {
	phoneID, username, testuid, testEmail, teardownTestCase := setupAccount(t)
	testLogin := TestEmailLogin{testEmail, testEmail}
	session, err := PostAuthLoginByEmail("", "", "", "", "", testLogin)
	accessToken := session.AccessToken
	if err != nil || accessToken == "" {
		t.Errorf("Error logging in")
		fmt.Println(err)
	}
	return phoneID, username, testuid, testEmail, accessToken, func(t *testing.T) {
		defer teardownTestCase(t)
	}
}

func setupAccount(t *testing.T) (string, string, string, string, func(t *testing.T)) {
	t.Log("Setting up test case")
	PresetLoginRadiusTestEnv()
	time := time.Now()
	timestamp := time.Format("20060102150405")
	retEmail := "testemail" + timestamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", retEmail}, {"Secondary", "1" + retEmail}}
	username := "testemail" + timestamp
	phoneID := os.Getenv("PHONENUM")
	testAccount := AccountSetup{true, true, testEmails, retEmail, username, phoneID}
	user, err := PostManageAccountCreate(testAccount)
	uid := user.UID
	if err != nil || uid == "" {
		t.Errorf("Error creating account")
		fmt.Println(err)
	}
	return phoneID, username, uid, retEmail, func(t *testing.T) {
		t.Log("Tearing down test case")
		_, err2 := DeleteManageAccount(uid)
		if err2 != nil {
			t.Errorf("Error cleaning up account")
			fmt.Println(err2)
		}
	}
}

func setupEmailVerificationAccount(t *testing.T) (string, string, string, func(t *testing.T)) {
	t.Log("Setting up test case")
	PresetLoginRadiusTestEnv()
	time := time.Now()
	timestamp := time.Format("20060102150405")
	retEmail := "testemail" + timestamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", retEmail}}
	username := "testemail" + timestamp
	phoneID := "+" + timestamp
	testAccount := AccountSetup{false, false, testEmails, retEmail, username, phoneID}
	user, err := PostManageAccountCreate(testAccount)
	uid := user.UID
	if err != nil || uid == "" {
		t.Errorf("Error creating account")
		fmt.Println(err)
	}
	tokenGen := TestEmail{retEmail}
	token, err3 := PostManageEmailVerificationToken(tokenGen)
	if err3 != nil {
		t.Errorf("Error generating token")
		fmt.Println(err3)
	}
	return phoneID, retEmail, token.VerificationToken, func(t *testing.T) {
		t.Log("Tearing down test case")
		_, err2 := DeleteManageAccount(uid)
		if err2 != nil {
			t.Errorf("Error cleaning up account")
			fmt.Println(err2)
		}
	}
}

func getPasswordResetToken(email string) string {
	resetEmail := TestEmail{email}
	passToken, err := PostManageForgotPasswordToken(resetEmail)
	if err != nil {
		return ""
	}
	return passToken.ForgotToken
}

func createClientGUID() string {
	rand.Seed(time.Now().UnixNano())
	buff := make([]byte, 64)
	rand.Read(buff)
	return base64.StdEncoding.EncodeToString(buff)
}
