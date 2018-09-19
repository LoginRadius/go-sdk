package loginradius

import (
	"log"
	"os"
)

// Candidate token can be retrieved by logging in with a
// social provider here https://docs.loginradius.com/api/v2/social-login/access-token
// and retrieving the token from the query parameters

// Api Key and Api secret can be obtained from the LoginRadius dashboard

// SetLoginRadiusEnv sets environment variables for use by the LoginRadius API.
func SetLoginRadiusEnv(apiKey, apiSecret, domain string) {
	setAPIKeyErr := os.Setenv("APIKEY", apiKey)
	if setAPIKeyErr != nil {
		log.Printf("Error setting Api Key using SetLoginRadiusEnv")
		log.Fatal(setAPIKeyErr)
	}
	setAPISecretErr := os.Setenv("APISECRET", apiSecret)
	if setAPISecretErr != nil {
		log.Printf("Error setting Api Secret using SetLoginRadiusEnv")
		log.Fatal(setAPISecretErr)
	}
	setDomainErr := os.Setenv("DOMAIN", domain)
	if setDomainErr != nil {
		log.Printf("Error setting domain using SetLoginRadiusEnv")
		log.Fatal(setDomainErr)
	}
}

// PresetLoginRadiusEnv is used if you would like to set your environment variables outside of your written code
// Fill out the information below and call PresetLoginRadiusEnv
func PresetLoginRadiusEnv() {
	setAPIKeyErr := os.Setenv("APIKEY", "")
	if setAPIKeyErr != nil {
		log.Printf("Error setting Api Key using PresetLoginRadiusEnv")
		log.Fatal(setAPIKeyErr)
	}
	setAPISecretErr := os.Setenv("APISECRET", "")
	if setAPISecretErr != nil {
		log.Printf("Error setting Api Secret using PresetLoginRadiusEnv")
		log.Fatal(setAPISecretErr)
	}
	setDomainErr := os.Setenv("DOMAIN", "")
	if setDomainErr != nil {
		log.Printf("Error setting domain using PresetLoginRadiusEnv")
		log.Fatal(setDomainErr)
	}
}
