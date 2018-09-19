package loginradius

import "os"

// WebHookBool is a container for all the single attribute boolean responses from a webhook API
type WebHookBool struct {
	IsPosted  bool `json:"IsPosted"`
	IsAllowed bool `json:"IsAllowed"`
	IsDeleted bool `json:"IsDeleted"`
}

// WebHook is the response for subscribed webhook URLs
type WebHook struct {
	IsAllowed bool `json:"IsAllowed"`
	Data      []struct {
		TargetURL string `json:"TargetUrl"`
		Event     string `json:"Event"`
	} `json:"data"`
	Count int `json:"Count"`
}

// PostWebhookSubscribe can be used to configure a WebHook on your LoginRadius site.
// Webhooks also works on subscribe and notification models by subscribing your hook and getting a notification.
// Equivalent to RESThook but these provide security on basis of signature and RESThooks work on unique URLs.
// Body parameters include a TargetUrl:string and Event:string
// Allowed events can be found on the online API documentation
func PostWebhookSubscribe(body interface{}) (WebHookBool, error) {
	data := new(WebHookBool)
	req, reqErr := CreateRequest("POST", "http://api.loginradius.com/api/v2/webhook", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("apisecret", os.Getenv("APISECRET"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// GetWebhookTest can be used to test a subscribed WebHook.
func GetWebhookTest() (WebHookBool, error) {
	data := new(WebHookBool)
	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/webhook/test", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("apisecret", os.Getenv("APISECRET"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetWebhookSubscribedURLs is used to fetch all the subscribed URLs,
// for particular event. For list of allowed events, check API docs.
func GetWebhookSubscribedURLs(event string) (WebHook, error) {
	data := new(WebHook)
	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/webhook/test", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("event", event)
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("apisecret", os.Getenv("APISECRET"))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// DeleteWebhookUnsubscribe can be used to unsubscribe a WebHook configured on your LoginRadius site.
// Body parameters include a TargetUrl:string and Event:string
// Allowed events can be found on the online API documentation
func DeleteWebhookUnsubscribe(body interface{}) (WebHookBool, error) {
	data := new(WebHookBool)
	req, reqErr := CreateRequest("DELETE", "http://api.loginradius.com/api/v2/webhook", body)
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("apisecret", os.Getenv("APISECRET"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}
