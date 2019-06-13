package httprutils

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
	"testing"
	"time"
)

func TestBuildURL(t *testing.T) {
	t.Parallel()
	host := "https://api.loginradius.com"
	queryParams := make(map[string]string)
	queryParams["test"] = "1"
	queryParams["test2"] = "2"
	testURL := AddQueryParams(host, queryParams)
	if testURL != "https://api.loginradius.com?test=1&test2=2" {
		t.Error("Bad BuildURL result")
	}
}

func TestBuildRequest(t *testing.T) {
	t.Parallel()
	method := Get
	baseURL := "https://api.loginradius.com"
	key := "API_KEY"
	Headers := make(map[string]string)
	Headers["Content-Type"] = "application/json"
	Headers["Authorization"] = "Bearer " + key
	queryParams := make(map[string]string)
	queryParams["test"] = "1"
	queryParams["test2"] = "2"
	body, err := EncodeBody("")
	request := Request{
		Method:      method,
		URL:         baseURL,
		Headers:     Headers,
		QueryParams: queryParams,
		Body:        body,
	}
	req, e := BuildRequestObject(request)
	if e != nil {
		t.Errorf("Rest failed to BuildRequest. Returned error: %v", e)
	}
	if req == nil {
		t.Errorf("Failed to BuildRequest.")
	}

	//Start PrintRequest
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	fmt.Println("Request : ", string(requestDump))
	//End Print Request
}

func TestBuildBadRequest(t *testing.T) {
	t.Parallel()
	request := Request{
		Method: Method("@"),
	}
	req, e := BuildRequestObject(request)
	if e == nil {
		t.Errorf("Expected an error for a bad HTTP Method")
	}
	if req != nil {
		t.Errorf("If there's an error there shouldn't be a Request.")
	}
}

func TestBuildResponse(t *testing.T) {
	t.Parallel()
	stub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"message\": \"success\"}")
	}))
	defer stub.Close()
	baseURL := stub.URL
	method := Get
	body, _ := EncodeBody("")
	request := Request{
		Method: method,
		URL:    baseURL,
		Body:   body,
	}
	req, e := BuildRequestObject(request)
	if e != nil {
		t.Error("Failed to BuildRequestObject", e)
	}
	res, e := MakeRequest(req)
	if e != nil {
		t.Error("Failed to MakeRequest", e)
	}
	response, e := BuildResponse(res)
	if response.StatusCode != 200 {
		t.Error("Invalid status code in BuildResponse")
	}
	if len(response.Body) == 0 {
		t.Error("Invalid response body in BuildResponse")
	}
	if len(response.Headers) == 0 {
		t.Error("Invalid response headers in BuildResponse")
	}
	if e != nil {
		t.Errorf("Rest failed to make a valid API request. Returned error: %v", e)
	}

	//Start Print Request
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	fmt.Println("Request :", string(requestDump))
	//End Print Request

}

type panicResponse struct{}

func (*panicResponse) Read(p []byte) (n int, err error) {
	panic(bytes.ErrTooLarge)
}

func (*panicResponse) Close() error {
	return nil
}

func TestBuildBadResponse(t *testing.T) {
	t.Parallel()
	res := &http.Response{
		Body: new(panicResponse),
	}
	_, e := BuildResponse(res)
	if e == nil {
		t.Errorf("This was a bad response and error should be returned")
	}
}

func testingAPI(t *testing.T, fn func(request Request) (*Response, error)) {
	stub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"message\": \"success\"}")
	}))
	defer stub.Close()

	host := stub.URL
	endpoint := "/test_endpoint"
	baseURL := host + endpoint
	key := "API_KEY"
	Headers := make(map[string]string)
	Headers["Content-Type"] = "application/json"
	Headers["Authorization"] = "Bearer " + key
	method := Get
	queryParams := make(map[string]string)
	queryParams["test"] = "1"
	queryParams["test2"] = "2"
	request := Request{
		Method:      method,
		URL:         baseURL,
		Headers:     Headers,
		QueryParams: queryParams,
	}

	//Start Print Request
	req, e := BuildRequestObject(request)
	if e != nil {
		t.Errorf("Error during BuildRequestObject: %v", e)
	}
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	fmt.Println("Request :", string(requestDump))
	//End Print Request

	response, e := fn(request)

	if response.StatusCode != 200 {
		t.Error("Invalid status code")
	}
	if len(response.Body) == 0 {
		t.Error("Invalid response body")
	}
	if len(response.Headers) == 0 {
		t.Error("Invalid response headers")
	}
	if e != nil {
		t.Errorf("Rest failed to make a valid API request. Returned error: %v", e)
	}
}

func TestCustomContentType(t *testing.T) {
	t.Parallel()
	host := "http://localhost"
	Headers := make(map[string]string)
	Headers["Content-Type"] = "custom"

	method := Get
	encoded, _ := EncodeBody("Hello World")

	request := Request{
		Method:  method,
		URL:     host,
		Headers: Headers,
		Body:    encoded,
	}
	response, _ := BuildRequestObject(request)
	if response.Header.Get("Content-Type") != "custom" {
		t.Error("Content-Type not modified correctly")
	}

	//Start Print Request
	requestDump, err := httputil.DumpRequest(response, true)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	fmt.Println("Request :", string(requestDump))
	//End Print Request
}

func TestCustomHTTPClient(t *testing.T) {
	t.Parallel()
	stub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 20)
		fmt.Fprintln(w, "{\"message\": \"success\"}")
	}))
	defer stub.Close()
	host := stub.URL
	endpoint := "/test_endpoint"
	baseURL := host + endpoint
	method := Get
	body, _ := EncodeBody("")
	request := Request{
		Method: method,
		URL:    baseURL,
		Body:   body,
	}

	customClient := &Client{&http.Client{Timeout: time.Millisecond * 10}}
	_, err := customClient.Send(request)
	if err == nil {
		t.Error("A timeout did not trigger as expected")
	}
	if !strings.Contains(err.Error(), "Client.Timeout exceeded while awaiting headers") {
		t.Error("We did not receive the Timeout error")
	}
}
