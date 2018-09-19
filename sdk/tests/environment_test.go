package loginradius

import (
	"os"
)

// Candidate token can be retrieved by logging in with a social provider here https://docs.com/api/v2/social-login/access-token
// and retrieving the token from the query parameters

// Api Key and Api secret can be obtained from the LoginRadius dashboard

// PresetLoginRadiusTestEnv is used if you would like to set your environment variables outside of your written code
// Fill out the information below and call PresetLoginRadiusTestEnv
func PresetLoginRadiusTestEnv() {
	os.Setenv("CANDIDATETOKEN", "")
	os.Setenv("FACEBOOKTOKEN", "")
	os.Setenv("TWITTERTOKEN", "-")
	os.Setenv("TWITTERSECRET", "")
	os.Setenv("VKONTAKTETOKEN", "")
	os.Setenv("APIKEY", "")
	os.Setenv("APISECRET", "")
	os.Setenv("DOMAIN", "")
	os.Setenv("SOCIALPAGEID", "")
	os.Setenv("SOCIALMESSAGE", "")
	os.Setenv("SECURITYQUESTION", "")
	os.Setenv("CUSTOMOBJECTNAME", "")
	os.Setenv("PHONENUM", "")
}

// Put the question ID under test in json
type SecurityQuestion struct {
	QuestionID string `json:"34f4ee507d3548788a8d0b22102fa21f"`
}
