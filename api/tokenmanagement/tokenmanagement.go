package tokenmanagement

import (
	"errors"

	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
	"github.com/LoginRadius/go-sdk/lrerror"
)

// GetAccessTokenViaFacebook is used to get a LoginRadius access token by sending Facebook’s access token.
// It will be valid for the specific duration of time specified in the response.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/native-social-login-api/access-token-via-facebook-token
// Required query parameter: key, fb_access_token
func (lr Loginradius) GetAccessTokenViaFacebook(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"fb_access_token": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["key"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/api/v2/access_token/facebook", validatedQueries)

	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAccessTokenViaTwitter is used to get a LoginRadius access token by sending Twitter’s access token.
// It will be valid for the specific duration of time specified in the response.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/native-social-login-api/access-token-via-twitter-token
// Required query parameter: key, tw_access_token, tw_token_secret
func (lr Loginradius) GetAccessTokenViaTwitter(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"tw_access_token": true, "tw_token_secret": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["key"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/api/v2/access_token/twitter", validatedQueries)

	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetAccessTokenViaVkontakte is used to get a LoginRadius access token by sending Vkontakte’s access token.
// It will be valid for the specific duration of time specified in the response.
// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/native-social-login-api/access-token-via-vkontakte-token
// Required query parameter: key, vk_access_token
func (lr Loginradius) GetAccessTokenViaVkontakte(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"vk_access_token": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["key"] = lr.Client.Context.ApiKey
	req := lr.Client.NewGetReq("/api/v2/access_token/vkontakte", validatedQueries)

	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetRefreshUserProfile is used to get the latest updated
// social profile data from the user’s social account after authentication.
// The social profile will be retrieved via oAuth and OpenID protocols.
// The data is normalized into LoginRadius’ standard data format.
// This API should be called using the access token retrieved from the refresh access token API.
// Documentation:https://www.loginradius.com/docs/api/v2/customer-identity-api/refresh-token/refresh-user-profile
// Required query parameter: access_token
func (lr Loginradius) GetRefreshUserProfile() (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	req := lr.Client.NewGetReq("/api/v2/userprofile/refresh")
	req.QueryParams = map[string]string{"access_token": lr.Client.Context.Token}
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetRefreshToken is used to refresh the provider access token after authentication.
// It will be valid for up to 60 days on LoginRadius depending on the provider. In order
// to use the access token in other APIs, always refresh the token using this API.
// Supported Providers : Facebook,Yahoo,Google,Twitter, Linkedin.
// Contact LoginRadius support team to enable this API.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/refresh-token/refresh-token
// Required query parameters: access_token, secret
// Optional query parameter: expiresin (Allows you to specify a desired expiration time in minutes for the newly issued access_token.)
func (lr Loginradius) GetRefreshToken(queries ...interface{}) (*httprutils.Response, error) {
	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	queryParams := map[string]string{
		"secret":       lr.Client.Context.ApiSecret,
		"access_token": lr.Client.Context.Token,
	}

	for _, arg := range queries {
		allowedQueries := map[string]bool{"expiresin": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)
		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			queryParams[k] = v
		}
	}

	req := lr.Client.NewGetReq("/api/v2/access_token/refresh", queryParams)
	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
