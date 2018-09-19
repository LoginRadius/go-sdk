package loginradius

import (
	"fmt"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	fmt.Println("Starting test TestGetConfiguration")
	PresetLoginRadiusTestEnv()
	_, err := GetConfiguration()
	if err != nil {
		t.Errorf("Error getting configuration")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
func TestGetServerTime(t *testing.T) {
	fmt.Println("Starting test TestGetServerTime")
	PresetLoginRadiusTestEnv()
	_, err := GetServerTime("")
	if err != nil {
		t.Errorf("Error getting server time")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
func TestGetGenerateSottAPI(t *testing.T) {
	fmt.Println("Starting test TestGetGenerateSottAPI")
	PresetLoginRadiusTestEnv()
	_, err := GetGenerateSottAPI("")
	if err != nil {
		t.Errorf("Error generating SOTT")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
func TestGetActiveSessionDetails(t *testing.T) {
	fmt.Println("Starting test TestGetActiveSessionDetails")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	_, err := GetActiveSessionDetails(accessToken)
	if err != nil {
		t.Errorf("Error getting active session details")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
