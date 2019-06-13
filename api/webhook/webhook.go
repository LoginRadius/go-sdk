package webhook

import (
	"github.com/LoginRadius/go-sdk/httprutils"
	lrvalidate "github.com/LoginRadius/go-sdk/internal/validate"
)

// PostWebhookSubscribe can be used to configure a WebHook on your LoginRadius site.
// Webhooks also works on subscribe and notification models by subscribing your hook and getting a notification.
// Equivalent to RESThook but these provide security on basis of signature and RESThooks work on unique URLs.
// Documentation: https://www.loginradius.com/docs/api/v2/integrations/webhooks/webhook-subscribe
// Required query parameters: apikey, apisecret
// Required post parameters: TargetUrl - string; Event - string
// For a list of all supported values for the Event parameter see documentation
func (lr Loginradius) PostWebhookSubscribe(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPostReq("/api/v2/webhook", body)
	if err != nil {
		return nil, err
	}
	lr.Client.NormalizeApiKey(req)
	req.QueryParams["apisecret"] = lr.Client.Context.ApiSecret
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetWebhookTest can be used to test a subscribed WebHook.
// Documentation https://www.loginradius.com/docs/api/v2/integrations/webhooks/webhook-test
// Required query parameters: apikey, apisecret
func (lr Loginradius) GetWebhookTest() (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/api/v2/webhook/test")
	lr.Client.NormalizeApiKey(req)
	req.QueryParams["apisecret"] = lr.Client.Context.ApiSecret
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetWebhookSubscribedURLs is used to fetch all the subscribed URLs,
// Documentation: https://www.loginradius.com/docs/api/v2/integrations/webhooks/webhook-subscribed-urls
// Required query parameters: apikey, apisecret, event
// For a list of all supported values for the Event parameter see documentation
func (lr Loginradius) GetWebhookSubscribedURLs(queries interface{}) (*httprutils.Response, error) {
	allowedQueries := map[string]bool{"event": true}
	validatedQueries, err := lrvalidate.Validate(allowedQueries, queries)
	if err != nil {
		return nil, err
	}
	validatedQueries["apisecret"] = lr.Client.Context.ApiSecret
	req := lr.Client.NewGetReq("/api/v2/webhook", validatedQueries)
	lr.Client.NormalizeApiKey(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteWebhookUnsubscribe can be used to unsubscribe a WebHook configured on your LoginRadius site.
// Documentation: https://www.loginradius.com/docs/api/v2/integrations/webhooks/webhook-unsubscribe
// For a list of all supported values for the Event parameter see documentation
// Required query parameters: apikey, apisecret
// Required post parameters: targeturl - string, event - string
func (lr Loginradius) DeleteWebhookUnsubscribe(body interface{}) (*httprutils.Response, error) {
	req := lr.Client.NewDeleteReq("/api/v2/webhook", body)
	req.QueryParams = map[string]string{
		"apisecret": lr.Client.Context.ApiSecret,
		"apikey":    lr.Client.Context.ApiKey,
	}
	req.Headers = httprutils.JSONHeader
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
