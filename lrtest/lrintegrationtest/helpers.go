// Package lrintegrationtest contains integration tests for all API calls.
package lrintegrationtest

import (
	"encoding/base64"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	lr "github.com/LoginRadius/go-sdk"
	lraccount "github.com/LoginRadius/go-sdk/api/account"
	lrauthentication "github.com/LoginRadius/go-sdk/api/authentication"
	"github.com/LoginRadius/go-sdk/api/customobject"
	"github.com/LoginRadius/go-sdk/api/role"
	"github.com/LoginRadius/go-sdk/api/webhook"
	"github.com/LoginRadius/go-sdk/lrbody"
	"github.com/LoginRadius/go-sdk/lrjson"
)

func setupAccount(t *testing.T) (string, string, string, string, *lr.Loginradius, func(t *testing.T)) {
	t.Log("Setting up test case")

	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	loginradius, _ := lr.NewLoginradius(&cfg)
	authlr := lraccount.Loginradius{loginradius}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	testEmail := "lrtest" + timeStamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", testEmail}, {"Secondary", "1" + testEmail}}
	username := "lrtest" + timeStamp
	phoneID := "+1" + timeStamp
	testAccount := AccountSetup{true, true, testEmails, testEmail, username, phoneID}

	response, err := lraccount.Loginradius(authlr).PostManageAccountCreate(testAccount)
	if err != nil {
		t.Errorf("Error calling PostManageAccountCreate from setupAccount: %v", err)
	}
	user, err := lrjson.DynamicUnmarshal(response.Body)
	uid := user["Uid"].(string)
	if err != nil || uid == "" {
		t.Errorf("Error creating account: %+v", err)
	}

	return phoneID, username, uid, testEmail, loginradius, func(t *testing.T) {
		t.Log("Tearing down test case")
		_, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).DeleteManageAccount(uid)
		if err != nil {
			t.Errorf("Error cleaning up account: %+v", err)
		}
	}
}

func setupEmailVerificationAccount(t *testing.T) (string, string, string, *lr.Loginradius, func(t *testing.T)) {
	t.Log("Setting up test case")

	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	loginradius, _ := lr.NewLoginradius(&cfg)
	authlr := lrauthentication.Loginradius{Client: loginradius}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	testEmail := "lrtest" + timeStamp + "@mailinator.com"
	testEmails := TestEmailArr{{"Primary", testEmail}}
	username := "lrtest" + timeStamp

	phoneID := "+" + timeStamp
	testAccount := AccountSetup{false, false, testEmails, testEmail, username, phoneID}
	response, err := lraccount.Loginradius(authlr).PostManageAccountCreate(testAccount)
	user, _ := lrjson.DynamicUnmarshal(response.Body)
	uid := user["Uid"].(string)
	if err != nil || uid == "" {
		t.Errorf("Error creating account: %+v", err)
	}

	tokenGen := TestEmail{testEmail}
	response, err = lraccount.Loginradius(authlr).PostManageEmailVerificationToken(tokenGen)
	data, _ := lrjson.DynamicUnmarshal(response.Body)
	token := data["VerificationToken"].(string)
	if err != nil {
		t.Errorf("Error generating token: %+v", err)
	}

	return phoneID, testEmail, token, loginradius, func(t *testing.T) {
		t.Log("Tearing down test case")
		_, err := lraccount.Loginradius(lraccount.Loginradius{loginradius}).DeleteManageAccount(uid)
		if err != nil {
			t.Errorf("Error cleaning up account: %+v", err)
		}
	}
}

func setupLogin(t *testing.T) (string, string, string, string, string, *lr.Loginradius, func(t *testing.T)) {
	phoneID, username, uid, email, loginradius, teardownTestCase := setupAccount(t)
	authlr := lrauthentication.Loginradius{loginradius}
	// testLogin := TestEmailLogin{email, email}
	// response, err := lrauthentication.Loginradius(authlr).PostAuthLoginByEmail(testLogin)
	response, err := lraccount.Loginradius(authlr).GetManageAccessTokenUID(map[string]string{"uid": uid})
	session, _ := lrjson.DynamicUnmarshal(response.Body)
	accessToken := session["access_token"].(string)
	if err != nil || accessToken == "" {
		t.Errorf("Error logging in: %+v", err)
	}
	loginradius.Context.Token = accessToken
	return phoneID, username, uid, email, accessToken, loginradius, func(t *testing.T) {
		defer teardownTestCase(t)
	}
}

func setupCustomObject(t *testing.T) (string, string, string, string, *lr.Loginradius, func(t *testing.T)) {
	t.Log("Setting up test case")
	_, _, uid, _, accessToken, lrclient, teardownTestCase := setupLogin(t)

	objName := os.Getenv("CUSTOMOBJECTNAME")
	time := time.Now()
	timestamp := time.Format("20060102150405")
	customObj := map[string]string{
		"custom1": timestamp + "0",
		"custom2": timestamp + "1",
	}
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByUID(uid, map[string]string{"objectname": objName}, customObj)
	if err != nil {
		t.Errorf("Error creating custom object: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	objId := obj["Id"].(string)
	return accessToken, uid, objId, objName, lrclient, func(t *testing.T) {
		t.Log("Tearing down test case")
		defer teardownTestCase(t)
		_, err = customobject.Loginradius(customobject.Loginradius{lrclient}).DeleteCustomObjectByObjectRecordIDAndUID(uid, objId, map[string]string{"objectname": objName})
		if err != nil {
			t.Errorf("Error deleting custom object")
		}
	}
}

func genGUID() string {
	rand.Seed(time.Now().UnixNano())
	buff := make([]byte, 64)
	rand.Read(buff)
	return base64.StdEncoding.EncodeToString(buff)
}

func setupWebhook(t *testing.T) (string, string, *lr.Loginradius, func(t *testing.T)) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, err := lr.NewLoginradius(&cfg)

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	targeturl := "http://requestbin.fullcontact.com/u1imzuu1"
	event := "Register"

	res, err := webhook.Loginradius(webhook.Loginradius{lrclient}).PostWebhookSubscribe(
		map[string]string{
			"TargetUrl": targeturl,
			"Event":     event,
		},
	)

	if err != nil {
		t.Errorf("Error calling PostWebHookSubscribe: %v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !posted["IsPosted"].(bool) {
		t.Errorf("Error returned from PostWebHookSubscribe: %v", err)
	}

	return targeturl, event, lrclient, func(t *testing.T) {
		res, err := webhook.Loginradius(webhook.Loginradius{lrclient}).DeleteWebhookUnsubscribe(
			map[string]string{
				"targeturl": targeturl,
				"event":     "Register",
			},
		)

		if err != nil {
			t.Errorf("Error calling DeleteWebhookUnsubscribe to tear down test case: %v", err)
		}

		deleted, err := lrjson.DynamicUnmarshal(res.Body)
		if err != nil || !deleted["IsDeleted"].(bool) {
			t.Errorf("Error returned from DeleteWebhookUnsubscribe to tear down test case: %v", err)
		}
	}
}

func setupRole(t *testing.T) (string, string, *lr.Loginradius, func(t *testing.T)) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, err := lr.NewLoginradius(&cfg)
	permissionName := "example_permission_1"

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	rolename := "example_role_name"

	testrole := lrbody.Role{
		Name: rolename,
		Permissions: map[string]bool{
			permissionName:         true,
			"example_permission_2": true,
		},
	}
	roles := lrbody.Roles{[]lrbody.Role{testrole}}
	_, err = role.Loginradius(role.Loginradius{lrclient}).PostRolesCreate(roles)
	if err != nil {
		t.Errorf("Error calling PostRolesCreate: %v", err)
	}

	return permissionName, rolename, lrclient, func(t *testing.T) {
		res, err := role.Loginradius(role.Loginradius{lrclient}).DeleteAccountRole(rolename)
		if err != nil {
			t.Errorf("Error calling DeleteAccountRole: %v", err)
		}
		deleted, err := lrjson.DynamicUnmarshal(res.Body)
		if err != nil || !deleted["IsDeleted"].(bool) {
			t.Errorf("Error returned from DeleteAccountRole: %v, %v", deleted, err)
		}
	}
}
