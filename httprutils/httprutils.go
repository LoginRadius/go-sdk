// The httputils package holds functions and structs for making RESTful API calls
package httprutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/LoginRadius/go-sdk/lrerror"
)

type Method string

const (
	Get    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Delete Method = "DELETE"
)

// Request holds the request to an API Call.
type Request struct {
	Method      Method
	URL         string
	Headers     map[string]string
	QueryParams map[string]string
	Body        *bytes.Buffer
}

// DefaultClient is used if no custom HTTP client is defined
// http.Client configures a timeout that short-circuits long-running connections.
// The default for this value is 0, which is interpreted as “no timeout”, meaning if no custom client
// with a TimeOut value specified is defined, in case of API outage the GO program will continue to hang
// A custom TimeoutClient is defined below.
var DefaultClient = &Client{HTTPClient: http.DefaultClient}

// Client allows modification of client headers, redirect policy
// and other settings
// See https://golang.org/pkg/net/http
type Client struct {
	HTTPClient *http.Client
}

// Response holds the response from an API call.
type Response struct {
	StatusCode int
	Body       string
	Headers    map[string][]string
	OrigBody   []byte
}

// BuildRequestObject creates the HTTP request object.
func BuildRequestObject(request Request) (*http.Request, error) {
	// Add any query parameters to the URL.
	if len(request.QueryParams) != 0 {
		request.URL = AddQueryParams(request.URL, request.QueryParams)
	}

	if request.Body == nil {
		encodedBody, _ := EncodeBody("")
		request.Body = encodedBody
	}

	req, err := http.NewRequest(string(request.Method), request.URL, request.Body)
	if err != nil {
		err = lrerror.New("EncodingError", "Error constructing http request", err)
		return req, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// MakeRequest makes the API call.
func MakeRequest(req *http.Request) (*http.Response, error) {
	return DefaultClient.HTTPClient.Do(req)
}

// BuildResponse builds the response struct.
func BuildResponse(res *http.Response) (*Response, error) {
	response := Response{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err := lrerror.New("EncodingError", "Error reading the response body", err)
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		err = lrerror.New("LoginradiusRespondedWithError", "Received error response from Loginradius", errors.New(string(body)))
	} else {
		response = Response{
			StatusCode: res.StatusCode,
			Body:       string(body),
			Headers:    res.Header,
			OrigBody:   body,
		}
	}

	res.Body.Close()
	return &response, err
}

func Send(request Request) (*Response, error) {
	return DefaultClient.Send(request)
}

// The following functions enable the ability to define a
// custom HTTP Client

// MakeRequest makes the API call.
func (c *Client) MakeRequest(req *http.Request) (*http.Response, error) {
	return c.HTTPClient.Do(req)
}

// Send will build your request, make the request, and build your response.
func (c *Client) Send(request Request) (*Response, error) {
	// Build the HTTP request object.
	req, err := BuildRequestObject(request)
	if err != nil {
		return nil, err
	}

	// Build the HTTP client and make the request.
	res, err := c.MakeRequest(req)
	if err != nil {
		err := lrerror.New("MakeRequestError", "Error making the request", err)
		return nil, err
	}

	// Build Response object.
	return BuildResponse(res)
}

// Custom client with time out after 8 seconds
var NetClient = &http.Client{
	Timeout: time.Second * 8,
}

var TimeoutClient = &Client{HTTPClient: NetClient}

// EncodeBody takes an interface and returns a *bytes.Buffer suitable for making RESTful requests with
func EncodeBody(body interface{}) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	asserted, ok := body.([]byte)
	if ok {
		var raw map[string]interface{}
		json.Unmarshal(asserted, &raw)
		encodeErr := json.NewEncoder(buffer).Encode(&raw)
		if encodeErr != nil {
			err := lrerror.New("EncodingError", "Error encoding the request body", encodeErr)
			return nil, err
		}
	} else {

		encodeErr := json.NewEncoder(buffer).Encode(body)
		if encodeErr != nil {
			err := lrerror.New("EncodingError", "Error encoding the request body", encodeErr)
			return nil, err
		}
	}
	return buffer, nil
}

// AddQueryParams takes a map of query params and appends each to the URL
func AddQueryParams(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}
