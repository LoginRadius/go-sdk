package lrintegrationtest

import (
	"os"
	"testing"

	lr "github.com/LoginRadius/go-sdk"
	"github.com/LoginRadius/go-sdk/api/smartlogin"
	"github.com/LoginRadius/go-sdk/lrerror"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

func TestGetSmartLoginByEmail(t *testing.T) {
	_, _, _, email, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := smartlogin.Loginradius(smartlogin.Loginradius{lrclient}).GetSmartLoginByEmail(
		map[string]string{"email": email, "clientguid": genGUID()},
	)
	if err != nil {
		t.Errorf("Error calling GetSmartLoginByEmail: %v", err)
	}
	result, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !result["IsPosted"].(bool) {
		t.Errorf("Error returned from GetSmartLoginByEmail: %v", err)
	}
}

func TestGetSmartLoginByUsername(t *testing.T) {
	_, username, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	res, err := smartlogin.Loginradius(smartlogin.Loginradius{lrclient}).GetSmartLoginByUsername(
		map[string]string{"username": username, "clientguid": genGUID()},
	)
	if err != nil {
		t.Errorf("Error calling GetSmartLoginByUsername: %v", err)
	}
	result, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !result["IsPosted"].(bool) {
		t.Errorf("Error returned from GetSmartLoginByUsername: %v", err)
	}
}

func TestGetSmartLoginPing(t *testing.T) {
	_, username, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	guid := genGUID()
	_, err := smartlogin.Loginradius(smartlogin.Loginradius{lrclient}).GetSmartLoginByUsername(
		map[string]string{"username": username, "clientguid": guid},
	)
	if err != nil {
		t.Errorf("Error calling GetSmartLoginByUsername for GetSmartLoginPing: %v", err)
	}
	_, err = smartlogin.Loginradius(smartlogin.Loginradius{lrclient}).GetSmartLoginPing(
		map[string]string{"clientguid": guid},
	)
	// Making a request to SmartLoginByUsername without clicking on the link sent to the generated email
	// will result in GetSmartLoginPing returning an error from Loginradius
	if err == nil || err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("Error calling GetSmartLoginPing: %v", err)
	}
}

// To run this test, comment out t.SkipNow() and manually set the verification token
// after making a smart login request with a manually created user, and retrieving the verification
// token from the link received by the registraiton email
func TestGetSmartLoginVerifyToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	res, err := smartlogin.Loginradius(smartlogin.Loginradius{lrclient}).GetSmartLoginVerifyToken(
		// manually set verification token here
		map[string]string{"verificationtoken": "28cfbc5f6378490ea0ac4479faa64e3a"},
	)
	if err != nil {
		t.Errorf("Error calling GetSmartLoginVerifyToken: %v", err)
	}
	result, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !result["IsPosted"].(bool) {
		t.Errorf("Error returned from GetSmartLoginVerifyToken: %v", err)
	}
}
