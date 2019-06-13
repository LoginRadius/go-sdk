package lrintegrationtest

import (
	"os"
	"testing"

	lr "github.com/LoginRadius/go-sdk"
	lrconfiguration "github.com/LoginRadius/go-sdk/api/configuration"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

func TestGetConfiguration(t *testing.T) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	_, err := lrconfiguration.Loginradius(lrconfiguration.Loginradius{lrclient}).GetConfiguration()
	if err != nil {
		t.Errorf("Error calling GetConfiguration: %v", err)
	}
}

func TestGetServerTime(t *testing.T) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	res, err := lrconfiguration.Loginradius(lrconfiguration.Loginradius{lrclient}).GetServerTime()
	if err != nil {
		t.Errorf("Error calling GetServerTime: %v", err)
	}
	time, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || time["CurrentTime"].(string) == "" {
		t.Errorf("Error returned from GetServerTime: %v", err)
	}
}

func TestGetGenerateSottAPI(t *testing.T) {
	SetTestEnv()

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, _ := lr.NewLoginradius(&cfg)

	res, err := lrconfiguration.Loginradius(lrconfiguration.Loginradius{lrclient}).GetGenerateSottAPI()
	if err != nil {
		t.Errorf("Error calling GetGenerateSottAPI: %v", err)
	}
	time, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || time["Sott"].(string) == "" {
		t.Errorf("Error returned from GetGenerateSottAPI: %v", err)
	}
}

func TestGetActiveSessionDetails(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	res, err := lrconfiguration.Loginradius(lrconfiguration.Loginradius{lrclient}).GetActiveSessionDetails()
	if err != nil {
		t.Errorf("Error making GetActiveSessionDetails call, %v", err)
	}
	unmarshalled, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error unmarshalling GetActiveSessionDetails return, %v", err)
	}
	data := unmarshalled["data"].([]interface{})
	if err != nil || data[0].(map[string]interface{})["AccessToken"].(string) == "" {
		t.Errorf("Error returned from GetActiveSessionDetails call %v", err)
	}
}
