package loginradius

import "os"

// PostSocialMessageAPI is used to post messages to the user’s contacts.
// Supported Providers: Twitter, LinkedIn
// This is one of the APIs that makes up the LoginRadius Friend Invite System. After using the Contact API,
// you can send messages to the retrieved contacts. This API requires setting permissions in your LoginRadius Dashboard.
// GET & POST Message API work the same way except the API method is different
func PostSocialMessageAPI(accessToken, to, subject, message string) (SocialMessageAPI, error) {
	data := new(SocialMessageAPI)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/api/v2/message", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	q.Add("to", to)
	q.Add("subject", subject)
	q.Add("message", message)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// PostSocialStatusPost is used to update the status on the user’s wall.
// Supported Providers: Facebook, Twitter, LinkedIn
func PostSocialStatusPost(accessToken, title, url, imageurl, status, caption,
	description string) (SocialMessageAPI, error) {
	data := new(SocialMessageAPI)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN")+"/api/v2/status", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	q.Add("url", url)
	q.Add("title", title)
	q.Add("imageurl", imageurl)
	q.Add("status", status)
	q.Add("caption", caption)
	q.Add("description", description)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialAccessToken Is used to translate the Request Token
// returned during authentication into an Access Token that can be used with other API calls.
func GetSocialAccessToken(token string) (SocialAccessToken, error) {
	data := new(SocialAccessToken)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/access_token", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("secret", os.Getenv("APISECRET"))
	q.Add("token", token)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialTokenValidate validates access_token, if valid then returns a response with its expiry otherwise error.
func GetSocialTokenValidate(accessToken string) (SocialAccessToken, error) {
	data := new(SocialAccessToken)
	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/access_token/validate", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("APIKEY"))
	q.Add("secret", os.Getenv("APISECRET"))
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialAccessTokenInvalidate invalidates the active access token or expires an access token validity.
func GetSocialAccessTokenInvalidate(accessToken string) (SocialAccessToken, error) {
	data := new(SocialAccessToken)
	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/access_token/invalidate", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("APIKEY"))
	q.Add("secret", os.Getenv("APISECRET"))
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialAlbum returns the photo albums associated with the passed in access tokens Social Profile.
// Supported Providers: Facebook, Google, Live, Vkontakte.
func GetSocialAlbum(accessToken string) (SocialAlbum, error) {
	data := new(SocialAlbum)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/album", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialAudio is used to get audio files data from the user’s social account.
// Supported Providers: Live, Vkontakte
func GetSocialAudio(accessToken string) (SocialAudio, error) {
	data := new(SocialAudio)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/audio", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialCheckin is used to get check Ins data from the user’s social account.
// Supported Providers: Facebook, Foursquare, Vkontakte
func GetSocialCheckin(accessToken string) (SocialCheckin, error) {
	data := new(SocialCheckin)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/checkin", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialCompany is used to get the followed companies data from the user’s social account.
// Supported Providers: Facebook, LinkedIn
func GetSocialCompany(accessToken string) (SocialCompany, error) {
	data := new(SocialCompany)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/company", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialContact is used to get contacts/friends/connections data from the user’s social account.
// This is one of the APIs that makes up the LoginRadius Friend Invite System.
// The data will normalized into LoginRadius’ standard data format.
// This API requires setting permissions in your LoginRadius Dashboard.
// Note: Facebook restricts access to the list of friends that is returned.
// When using the Contacts API with Facebook you will only receive friends that have accepted some permissions with your app.
// Supported Providers: Facebook, Foursquare, Google, LinkedIn, Live, Twitter, Vkontakte, Yahoo
func GetSocialContact(accessToken string) (SocialContact, error) {
	data := new(SocialContact)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/contact", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialEvent is used to get the event data from the user’s social account.
// Supported Providers: Facebook, Live
func GetSocialEvent(accessToken string) (SocialEvent, error) {
	data := new(SocialEvent)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/event", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialFollowing is used to get the following user list from the user’s social account.
// Supported Providers: Twitter
func GetSocialFollowing(accessToken string) (SocialFollowing, error) {
	data := new(SocialFollowing)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/following", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialGroup is used to get group data from the user’s social account.
// Supported Providers: Facebook, Vkontakte
func GetSocialGroup(accessToken string) (SocialGroup, error) {
	data := new(SocialGroup)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/group", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialLike is used to get likes data from the user’s social account.
// Supported Providers: Facebook
func GetSocialLike(accessToken string) (SocialLike, error) {
	data := new(SocialLike)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/like", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialMention is used to get mention data from the user’s social account.
// Supported Providers: Facebook
func GetSocialMention(accessToken string) (SocialMention, error) {
	data := new(SocialMention)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/mention", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialMessageAPI is used to post messages to the user’s contacts.
// Supported Providers: Twitter, LinkedIn
// This is one of the APIs that makes up the LoginRadius Friend Invite System. After using the Contact API,
// you can send messages to the retrieved contacts. This API requires setting permissions in your LoginRadius Dashboard.
// GET & POST Message API work the same way except the API method is different
func GetSocialMessageAPI(accessToken, to, subject, message string) (SocialMessageAPI, error) {
	data := new(SocialMessageAPI)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/message", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	q.Add("to", to)
	q.Add("subject", subject)
	q.Add("message", message)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialPage is used to get the page data from the user’s social account.
// Supported Providers: Facebook, LinkedIn
func GetSocialPage(accessToken, pagename string) (SocialPage, error) {
	data := new(SocialPage)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/page", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	q.Add("pagename", pagename)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialPhoto is used to get the photo data from the user’s social account.
// Supported Providers: Facebook, LinkedIn
func GetSocialPhoto(accessToken, albumid string) (SocialPhoto, error) {
	data := new(SocialPhoto)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/photo", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	q.Add("albumid", albumid)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialPost is used to get post message data from the user’s social account.
// Supported Providers: Facebook
func GetSocialPost(accessToken string) (SocialPost, error) {
	data := new(SocialPost)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/post", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialStatusFetch is used to get the status messages from the user’s social account.
// Supported Providers: Facebook, LinkedIn, Twitter, Vkontakte
func GetSocialStatusFetch(accessToken string) (SocialStatus, error) {
	data := new(SocialStatus)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/status", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialStatusPost is used to update the status on the user’s wall.
// Supported Providers: Facebook, Twitter, LinkedIn
func GetSocialStatusPost(accessToken, title, url, imageurl, status,
	caption, description string) (SocialMessageAPI, error) {
	data := new(SocialMessageAPI)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/status/js", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	q.Add("title", title)
	q.Add("url", url)
	q.Add("imageurl", imageurl)
	q.Add("status", status)
	q.Add("caption", caption)
	q.Add("description", description)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialUserProfile is used to get social profile data from the user’s social account after authentication.
// Supported Providers: All
func GetSocialUserProfile(accessToken string) (AuthProfile, error) {
	data := new(AuthProfile)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/userprofile", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetSocialVideo is used to get video files data from the user’s social account.
// Supported Providers: All
func GetSocialVideo(accessToken, nextcursor string) (SocialVideo, error) {
	data := new(SocialVideo)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN")+"/api/v2/video", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("access_token", accessToken)
	q.Add("nextcursor", nextcursor)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}
