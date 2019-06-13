package lrintegrationtest

import (
	"os"
	"testing"

	lr "github.com/LoginRadius/go-sdk"
	"github.com/LoginRadius/go-sdk/api/tokenmanagement"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

func TestGetAccessTokenViaFacebook(t *testing.T) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, err := lr.NewLoginradius(&cfg)

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetAccessTokenViaFacebook(
		map[string]string{"fb_access_token": "abcd1234abcd"},
	)

	if err != nil {
		t.Errorf("Error calling GetAccessTokenViaFacebook: %v", err)
	}

	tokens, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || tokens["access_token"].(string) == "" {
		t.Errorf("Error returned from GetAccessTokenViaFacebook: %v, %v", err, tokens)
	}
}

func TestGetAccessTokenViaTwitter(t *testing.T) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, err := lr.NewLoginradius(&cfg)

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetAccessTokenViaTwitter(
		map[string]string{"tw_access_token": "abcd1234abcd", "tw_token_secret": "abcd1234"},
	)

	if err != nil {
		t.Errorf("Error calling GetAccessTokenViaTwitter: %v", err)
	}

	tokens, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || tokens["access_token"].(string) == "" {
		t.Errorf("Error returned from GetAccessTokenViaTwitter: %v, %v", err, tokens)
	}
}

func TestGetAccessTokenViaVkontakte(t *testing.T) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, err := lr.NewLoginradius(&cfg)

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetAccessTokenViaVkontakte(
		map[string]string{"vk_access_token": "abcd1234abcd"},
	)

	if err != nil {
		t.Errorf("Error calling GetAccessTokenViaVkontakte: %v", err)
	}

	tokens, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || tokens["access_token"].(string) == "" {
		t.Errorf("Error returned from GetAccessTokenViaVkontakte: %v, %v", err, tokens)
	}
}

func TestGetRefreshUserProfile(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetRefreshUserProfile()
	if err != nil {
		t.Errorf("Error making call to GetRefreshUserProfile: %+v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || profile["Uid"] == "" {
		t.Errorf("Error returned from GetRefreshUserProfile: %+v", err)
	}
}

func TestGetRefreshToken(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := tokenmanagement.Loginradius(tokenmanagement.Loginradius{lrclient}).GetRefreshToken()
	if err != nil {
		t.Errorf("Error making call to GetRefreshToken: %+v", err)
	}
	profile, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || profile["access_token"] == "" {
		t.Errorf("Error returned from GetRefreshToken: %+v", err)
	}
}
