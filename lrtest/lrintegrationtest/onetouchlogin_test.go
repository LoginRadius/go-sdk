package lrintegrationtest

import (
	"os"
	"testing"

	lr "github.com/LoginRadius/go-sdk"
	"github.com/LoginRadius/go-sdk/api/onetouchlogin"
	"github.com/LoginRadius/go-sdk/lrerror"
)

func TestPostOneTouchLoginByEmail(t *testing.T) {
	_, _, _, email, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	// Should return {"Description":"CAPTCHA is invalid, please enter the correct CAPTCHA value.","ErrorCode":1218,"Message":"CAPTCHA is invalid","IsProviderError":false,"ProviderErrorResponse":null}
	_, err := onetouchlogin.Loginradius(onetouchlogin.Loginradius{lrclient}).PostOneTouchLoginByEmail(
		map[string]string{"clientguid": genGUID(), "email": email, "g-recaptcha-response": "abcd"},
	)
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("Error calling PostOneTouchLoginByEmail: %v", err)
	}
}

func TestPostOneTouchLoginByPhone(t *testing.T) {
	phoneID, _, _, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	// Should return {"Description":"CAPTCHA is invalid, please enter the correct CAPTCHA value.","ErrorCode":1218,"Message":"CAPTCHA is invalid","IsProviderError":false,"ProviderErrorResponse":null}
	_, err := onetouchlogin.Loginradius(onetouchlogin.Loginradius{lrclient}).PostOneTouchLoginByPhone(
		map[string]string{"clientguid": genGUID(), "phone": phoneID, "g-recaptcha-response": "abcd"},
	)
	if err.(lrerror.Error).Code() != "LoginradiusRespondedWithError" {
		t.Errorf("Error calling PostOneTouchLoginByPhone: %v", err)
	}
}

// To run this test, comment out t.SkipNow() and manually
// fill out phone number of a valid user profile and the otp
// received by that phone number after making a one touch login request
// in the test
func TestPutOneTouchOTPVerification(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	// Manually set otp and phone here
	_, err := onetouchlogin.Loginradius(onetouchlogin.Loginradius{lrclient}).PutOneTouchOTPVerification(
		map[string]string{"otp": ""},
		map[string]string{"phone": ""},
	)
	if err != nil {
		t.Errorf("Error calling PutOneTouchOTPVerification: %v", err)
	}
}
