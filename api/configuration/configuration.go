package lrconfiguration

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// GetConfiguration is used to get the configurations which are set in the
// LoginRadius Dashboard for a particular LoginRadius site/environment.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/configuration/get-configurations
// Required query parameter: apikey
func (lr Loginradius) GetConfiguration() (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("")
	req.URL = "https://config.lrcontent.com/ciam/appinfo"
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetServerTime allows you to query your LoginRadius account for basic server information
// and server time information which is useful when generating an SOTT token.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/configuration/get-server-time
// Required query parameter: apikey
// Optional query parameter: timedifference
func (lr Loginradius) GetServerTime(queries ...interface{}) (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/identity/v2/serverinfo")
	for _, arg := range queries {
		allowedQueries := map[string]bool{"timedifference": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetGenerateSottAPI allows you to generate SOTT with a given expiration time.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/session/generate-sott-token
// Optional query parameter: timedifference
func (lr Loginradius) GetGenerateSottAPI(queries ...interface{}) (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/identity/v2/manage/account/sott")
	for _, arg := range queries {
		allowedQueries := map[string]bool{"timedifference": true}
		validatedQueries, err := lrvalidate.Validate(allowedQueries, arg)

		if err != nil {
			return nil, err
		}
		for k, v := range validatedQueries {
			req.QueryParams[k] = v
		}
	}
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetActiveSessionDetails is used to get all active sessions by Access Token.
// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/social-login/advanced-social-api/get-active-session-details
// Required query parameters: key, secret, access_token
func (lr Loginradius) GetActiveSessionDetails() (*httprutils.Response, error) {
	req := lr.Client.NewGetReq(
		"",
		map[string]string{
			"key":    lr.Client.Context.ApiKey,
			"token":  lr.Client.Context.Token,
			"secret": lr.Client.Context.ApiSecret,
		},
	)
	req.URL = "http://api.loginradius.com/api/v2/access_token/activesession"
	delete(req.QueryParams, "apiKey")
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
	// data := new(ActiveSession)
	// req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/access_token/activesession", "")
	// if reqErr != nil {
	// 	return *data, reqErr
	// }

	// q := req.URL.Query()
	// q.Add("token", accessToken)
	// q.Add("key", os.Getenv("APIKEY"))
	// q.Add("secret", os.Getenv("APISECRET"))
	// req.URL.RawQuery = q.Encode()
	// req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	// err := RunRequest(req, data)
	// return *data, err
}
