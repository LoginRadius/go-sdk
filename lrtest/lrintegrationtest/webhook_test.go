package lrintegrationtest

import (
	"testing"

	webhook "github.com/LoginRadius/go-sdk/api/webhook"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

func TestPostWebhookSubscribe(t *testing.T) {
	_, _, _, tearDown := setupWebhook(t)
	defer tearDown(t)
}

func TestGetWebhookSubscribedURLs(t *testing.T) {
	targeturl, event, lrclient, tearDown := setupWebhook(t)
	defer tearDown(t)
	res, err := webhook.Loginradius(webhook.Loginradius{lrclient}).GetWebhookSubscribedURLs(
		map[string]string{"event": event},
	)

	if err != nil {
		t.Errorf("Error calling GetWebhookSubscribedURLs: %v", err)
	}
	webhooks, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from GetWebhookSubscribedURLs: %v", err)
	}

	exists := false
	for _, hook := range webhooks["data"].([]interface{}) {
		asserted := hook.(map[string]interface{})
		if asserted["TargetUrl"] == targeturl {
			exists = true
		}
	}
	if !exists {
		t.Errorf("Target url was supposed to be returned from GetWebhookSubscribedURLs,%v but instead got: %v", targeturl, webhooks)
	}
}

func TestGetWebhookTest(t *testing.T) {
	_, _, lrclient, tearDown := setupWebhook(t)
	defer tearDown(t)
	res, err := webhook.Loginradius(webhook.Loginradius{lrclient}).GetWebhookTest()

	if err != nil {
		t.Errorf("Error calling GetWebhookTest: %v", err)
	}
	result, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !result["IsAllowed"].(bool) {
		t.Errorf("Error returned from GetWebhookTest: %v", err)
	}
}

func TestDeleteWebhookUnsubscribe(t *testing.T) {
	_, _, _, tearDown := setupWebhook(t)
	defer tearDown(t)
}
