package lrintegrationtest

import (
	"encoding/json"
	"os"
	"testing"

	lr "github.com/LoginRadius/go-sdk"
	lrsocial "github.com/LoginRadius/go-sdk/api/social"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

// Comment out t.SkipNow() and set LoginRadius access token in the test to run this test
// User must be manually created and account linked to either Twitter or LinkedIn
func TestPostSocialMessageAPI(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}
	// Set LoginRadius access token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "e902ad18-a237-4162-b0cd-74aba81ac4ab"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	queries := map[string]string{
		"to":      "740309126501146624",
		"subject": "Testing sdk",
		"message": "test message ignore please ignore",
	}

	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).PostSocialMessageAPI(queries)

	if err != nil {
		t.Errorf("Error calling PostSocialMessageAPI: %+v", err)
	}

	posted, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !posted["isPosted"].(bool) {
		t.Errorf("Error returned from PostSocialMessageAPI: %+v", err)
	}
}

// To run this test, comment out t.SkipNow() and manually fill out a social provider token
// for a LoginRadius user with a linked account
func TestPostSocialStatusPost(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	//Fill out social provider token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "3d00f57d-3215-4895-bd40-d46f5f693de1"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	queries := map[string]string{
		"url":         "https://www.cbc.ca/news/canada/toronto/ontario-school-classroom-cellphone-ban-1.5052564",
		"title":       "Ontario bans cellphones",
		"imageurl":    "https://i.cbc.ca/1.2692036.1404142390!/fileImage/httpImage/image.jpg_gen/derivatives/16x9_780/cellphone.jpg",
		"status":      "news article of ontario banning cellphoneeee",
		"caption":     "no more phones in classrooms next year",
		"description": "ontario bans cellphones in classrooms",
	}

	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).PostSocialStatusPost(queries)

	if err != nil {
		t.Errorf("Error calling PostSocialStatusPost: %+v", err)
	}

	posted, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !posted["isPosted"].(bool) {
		t.Errorf("Error returned from PostSocialStatusPost: %+v", err)
	}
}

// To run this test, comment out t.SkipNow() and manually set token
func TestGetSocialAccessToken(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}
	lrclient, err := lr.NewLoginradius(&cfg)

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	// Manually fill LoginRadius request token here
	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialAccessToken("")

	if err != nil {
		t.Errorf("Error calling GetSocialAccessToken: %+v", err)
	}

	data, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || data["access_token"].(string) == "" {
		t.Errorf("Error returned from GetSocialAccessToken: %+v", err)
	}
}

func TestGetSocialTokenValidate(t *testing.T) {
	_, _, _, _, token, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)

	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialTokenValidate()
	if err != nil {
		t.Errorf("Error calling GetSocialTokenValidate: %+v", err)
	}

	data, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || token != data["access_token"].(string) {
		t.Errorf("Error returned from GetSocialTokenValidate: %+v", err)
	}
}

func TestGetSocialTokenInvalidate(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)

	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialTokenInvalidate()
	if err != nil {
		t.Errorf("Error calling GetSocialTokenInvalidate: %+v", err)
	}

	posted, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !posted["isPosted"].(bool) {
		t.Errorf("Error returned from GetSocialTokenInvalidate: %+v", err)
	}

}

func TestGetSocialTokenInvalidateWithOptionalParam(t *testing.T) {
	_, _, _, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)

	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialTokenInvalidate(map[string]string{"preventRefresh": "true"})
	if err != nil {
		t.Errorf("Error calling GetSocialTokenInvalidate: %+v", err)
	}

	posted, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !posted["isPosted"].(bool) {
		t.Errorf("Error returned from GetSocialTokenInvalidate: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
// Returns an array
func TestGetSocialAlbum(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "8c8dc688-7aeb-4025-abd2-97d1ff4e52a3"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialAlbum()
	data := []interface{}{}
	err = json.Unmarshal([]byte(resp.Body), &data)
	if err != nil {
		t.Errorf("Error calling GetSocialAlbum: %+v", err)
	}

	firstAlbum := data[0].(map[string]interface{})
	if err != nil || firstAlbum["ID"].(string) == "" {
		t.Errorf("Error returned from GetSocialAlbum: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialAudio(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "2f0680e6-aa28-4972-afe1-d29dfba7f57a"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialAudio()
	data := []interface{}{}
	err = json.Unmarshal([]byte(resp.Body), &data)
	if err != nil {
		t.Errorf("Error returned from GetSocialAudio: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialCheckin(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "8dcc6e32-c3a3-4ad0-ac3d-2dc500c1f60b"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialCheckin()
	if err != nil {
		t.Errorf("Error calling GetSocialCheckin: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialCompany(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "bfba9704-b78a-4b69-b30f-f92c93278127"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialCompany()
	if err != nil {
		t.Errorf("Error calling GetSocialCompany: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialContact(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "84641575-1de4-4dc9-a112-c855bd1df374"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialContact()
	if err != nil {
		t.Errorf("Error calling GetSocialContact: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialEvent(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "84641575-1de4-4dc9-a112-c855bd1df374"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialEvent()
	if err != nil {
		t.Errorf("Error calling GetSocialEvent: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialFollowing(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "be1a87fa-2c08-48aa-9e55-3e80487a9589"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialFollowing()
	if err != nil {
		t.Errorf("Error calling GetSocialFollowing: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialGroup(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "0e00e5fa-6e0f-41a4-912a-9acf332bda44"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialGroup()
	if err != nil {
		t.Errorf("Error calling GetSocialGroup: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialLike(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "e891a42e-6f40-4016-bc16-530c9e8c907b"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialLike()
	if err != nil {
		t.Errorf("Error calling GetSocialLike: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialMention(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "a63d7ff0-770e-4c67-93a0-6d39fd1db48d"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialMention()
	if err != nil {
		t.Errorf("Error calling GetSocialMention: %+v", err)
	}
}

// To run this test, comment out t.SkipNow() and manually fill out a social provider token
// for a LoginRadius user with a linked account
func TestGetSocialStatusPost(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	//Fill out social provider token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "74693104-56e0-44e3-b473-bdb73da77051"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	queries := map[string]string{
		"url":         "https://www.cbc.ca/news/canada/toronto/ontario-school-classroom-cellphone-ban-1.5052564",
		"title":       "Ontario bans cellphones",
		"imageurl":    "https://i.cbc.ca/1.2692036.1404142390!/fileImage/httpImage/image.jpg_gen/derivatives/16x9_780/cellphone.jpg",
		"status":      "bbbb",
		"caption":     "no more phones in classrooms next year",
		"description": "ontario bans cellphones in classroomssss",
	}

	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialStatusPost(queries)

	if err != nil {
		t.Errorf("Error calling GetSocialStatusPost: %+v", err)
	}
	posted, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || !posted["isPosted"].(bool) {
		t.Errorf("Error returned from GetSocialStatusPost: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialUserProfile(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "74693104-56e0-44e3-b473-bdb73da77051"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	resp, err := lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialUserProfile()
	if err != nil {
		t.Errorf("Error calling GetSocialUserProfile: %+v", err)
	}

	profile, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || profile["Uid"].(string) == "" {
		t.Errorf("Error returned from GetSocialUserProfile: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialPage(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "3b4d9b8e-f9c6-4a2c-b0ef-a75f0aa49720"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialPage("pagename")
	if err != nil {
		t.Errorf("Error calling GetSocialPage: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialPhoto(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "c6395c46-b5d6-4a25-adfa-c9b140f3c8b9"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialPhoto("album535656757_260363817")
	if err != nil {
		t.Errorf("Error calling GetSocialPhoto: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialPost(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "fd04789b-1b49-43b7-b6a8-239128c1bc84"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialPost()
	if err != nil {
		t.Errorf("Error calling GetSocialPost: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialStatus(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "05853ee8-e348-4478-9cd3-c883c665cb8e"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialStatus()
	if err != nil {
		t.Errorf("Error calling GetSocialStatus: %+v", err)
	}
}

// To run this test, comment out t.SkipNow(), manually create user, link social account and set token
func TestGetSocialVideo(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "2fd4ba91-b2ff-4345-a6b9-0bbd7b6c2d0d"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialVideo()
	if err != nil {
		t.Errorf("Error calling GetSocialVideo: %+v", err)
	}
}

func TestGetSocialVideoWithQueryParam(t *testing.T) {
	t.SkipNow()
	SetTestEnv()
	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	// Manually fill LoginRadius token here
	lrclient, err := lr.NewLoginradius(&cfg, map[string]string{"token": "2fd4ba91-b2ff-4345-a6b9-0bbd7b6c2d0d"})

	if err != nil {
		t.Errorf("Error initiating lrclient")
	}

	_, err = lrsocial.Loginradius(lrsocial.Loginradius{lrclient}).GetSocialVideo(map[string]string{"nextcursor": "cursor"})
	if err != nil {
		t.Errorf("Error calling GetSocialVideo: %+v", err)
	}
}
