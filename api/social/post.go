package lrsocial

import (
	"errors"

	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
	"github.com/LoginRadius/go-sdk/lrerror"
)

// PostSocialMessageAPI is used to post messages to the user’s contacts.
// Supported Providers: Twitter, LinkedIn
// This is one of the APIs that makes up the LoginRadius Friend Invite System. After using the Contact API,
// you can send messages to the retrieved contacts. This API requires setting permissions in your LoginRadius Dashboard.
// Please ensure the access to Post Messages on behalf of your user is enabled through your admin console:
// Platform Configuration >> Social Login >> Profile Access Permissions

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/post-message-api

// Required query parameters: to - string; subject - string; message - string; access_token - string
func (lr Loginradius) PostSocialMessageAPI(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"to": true, "subject": true, "message": true,
	}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	validatedQueries["access_token"] = lr.Client.Context.Token

	request, err := lr.Client.NewPostReq("/api/v2/message", "", validatedQueries)
	if err != nil {
		return nil, err
	}

	request.Headers = httprutils.URLEncodedHeader

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}

// PostSocialStatusPost is used to update the status on the user’s wall.
// Supported Providers: Facebook, Twitter, LinkedIn

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/trackable-status-posting

// Required query parameters: url - string; title - string; imageurl-string; status-string; caption - string; description - string;
func (lr Loginradius) PostSocialStatusPost(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{
		"url": true, "title": true, "imageurl": true, "status": true, "caption": true, "description": true,
	}

	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}

	if lr.Client.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	validatedQueries["access_token"] = lr.Client.Context.Token

	request, err := lr.Client.NewPostReq("/api/v2/status", "", validatedQueries)
	if err != nil {
		return nil, err
	}

	request.Headers = httprutils.URLEncodedHeader

	response, err := httprutils.TimeoutClient.Send(*request)
	return response, err
}
