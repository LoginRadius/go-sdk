package lrunittest

import (
	"fmt"
	"net/http/httptest"
	"testing"

	lr "github.com/LoginRadius/go-sdk"
	lrauth "github.com/LoginRadius/go-sdk/api/authentication"
	"github.com/LoginRadius/go-sdk/httprutils"
)

const body = "test body"

func initTest(path string) (lr.Loginradius, *httptest.Server) {
	response := httprutils.Response{
		StatusCode: 200,
		Body:       body,
		Headers:    map[string][]string{},
	}

	stub := initTestServer(path, response)

	lrclient := initLr()
	lrclient.Domain = stub.URL
	return lrclient, stub
}

func TestPostAuthUserRegistrationByEmail(t *testing.T) {
	lr, stub := initTest("/identity/v2/auth/register?apiKey=&emailtemplate=&options=&verificationurl=")
	defer stub.Close()
	user := "user struct"
	res, _ := lrauth.Loginradius(lrauth.Loginradius{&lr}).PostAuthUserRegistrationByEmail(user)
	if res.StatusCode != 200 || res.Body != body {
		t.Errorf("Unit TestPostAuthUserRegistrationByEmail: received %v", res)
	}
}

func TestGetAuthVerifyEmail(t *testing.T) {
	lr, stub := initTest("/identity/v2/auth/email")
	defer stub.Close()
	res, _ := lrauth.Loginradius(lrauth.Loginradius{&lr}).GetAuthVerifyEmail(map[string]string{"verificationtoken": "abcd"})

	fmt.Println(res)
	if res.StatusCode != 200 || res.Body != body {
		t.Errorf("Unit TestPostAuthUserRegistrationByEmail: received %v", res)
	}
}
