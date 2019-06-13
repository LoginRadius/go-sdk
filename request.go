package loginradius

import (
	"errors"

	"github.com/LoginRadius/go-sdk/httprutils"
	"github.com/LoginRadius/go-sdk/lrerror"
)

// NewGetRequest takes a uri and query parameters, then constructs a GET request for LoginRadius API endpoints requiring access tokens
// being passed in Authorization Bearer header
func (lr Loginradius) NewGetReqWithToken(path string, queries ...map[string]string) (*httprutils.Request, error) {
	if lr.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	request := &httprutils.Request{
		Method: httprutils.Get,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"Authorization": "Bearer " + lr.Context.Token,
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}

	return request, nil
}

// NewGetRequest takes a uri and query parameters, then constructs a GET request for a LoginRadius API endpoint
func (lr Loginradius) NewGetReq(path string, queries ...map[string]string) *httprutils.Request {
	request := &httprutils.Request{
		Method:      httprutils.Get,
		URL:         lr.Domain + path,
		Headers:     httprutils.URLEncodedHeader,
		QueryParams: map[string]string{},
	}
	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request
}

// NewPostReqWithToken takes a uri, body, and query parameters, then constructs the request for LoginRadius PUT API end points requiring access tokens being passed in Authorization Bearer header
func (lr Loginradius) NewPostReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {

	if lr.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Post,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": "Bearer " + lr.Context.Token,
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}

	return request, nil
}

// NewPostReq takes a uri, body, and optional queries to construct a POST request for a LoginRadius POST API endpoint
func (lr Loginradius) NewPostReq(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {
	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Post,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request, nil
}

// NewPutReq takes a uri, body, and optional queries to construct a PUT request for a LoginRadius API endpoint
func (lr Loginradius) NewPutReq(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {
	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Put,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request, nil
}

// NewPutReqWithToken takes a uri and query parameters, then constructs a PUT request for LoginRadius API endpoints requiring access tokens
// being passed in Authorization Bearer header
func (lr Loginradius) NewPutReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {
	if lr.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Put,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": "Bearer " + lr.Context.Token,
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request, nil
}

// NewDeleteReq takes a uri, body, and optional queries to construct a DELETE request for a LoginRadius POST API endpoint
func (lr Loginradius) NewDeleteReq(path string, body ...interface{}) *httprutils.Request {
	if len(body) != 0 {
		encoded, err := httprutils.EncodeBody(body[0])
		if err != nil {
			return nil
		}
		return &httprutils.Request{
			Method:  httprutils.Delete,
			URL:     lr.Domain + path,
			Headers: httprutils.URLEncodedHeader,
			Body:    encoded,
		}
	} else {
		return &httprutils.Request{
			Method:  httprutils.Delete,
			URL:     lr.Domain + path,
			Headers: httprutils.URLEncodedHeader,
		}
	}
}

// NewDeleteReqWithToken takes a uri and query parameters, then constructs a PUT request for LoginRadius API endpoints requiring access tokens
// being passed in Authorization Bearer header
func (lr Loginradius) NewDeleteReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {
	if lr.Context.Token == "" {
		errMsg := "Must initialize Loginradius with access token for this API call."
		err := lrerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Delete,
		URL:    lr.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": "Bearer " + lr.Context.Token,
		},
		QueryParams: map[string]string{
			"apiKey": lr.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}

	return request, nil
}

// AddApiCredentialsToReqHeader removes the apiKey query parameter from a constructed request
// and add LoginRadius app credentials in the request headers
func (lr Loginradius) AddApiCredentialsToReqHeader(req *httprutils.Request) {
	delete(req.QueryParams, "apiKey")
	req.Headers["X-LoginRadius-ApiKey"] = lr.Context.ApiKey
	req.Headers["X-LoginRadius-ApiSecret"] = lr.Context.ApiSecret
}

// NormalizeApiKey normalizes the apikey parameter in queries for requests to be sent to
// LoginRadius endpoints that only accept "apikey"
func (lr Loginradius) NormalizeApiKey(req *httprutils.Request) {
	delete(req.QueryParams, "apiKey")
	req.QueryParams["apikey"] = lr.Context.ApiKey
}
