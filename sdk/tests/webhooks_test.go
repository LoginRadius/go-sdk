package loginradius

import (
	"fmt"
	"testing"
)

func TestPostWebhookSubscribe(t *testing.T) {
	fmt.Println("Starting test TestPostWebhookSubscribe")
	PresetLoginRadiusTestEnv()
	webhook := WebhookTest{"https://www.google.ca", "Register"}
	resp, err := PostWebhookSubscribe(webhook)
	if err != nil || resp.IsPosted != true {
		t.Errorf("Error setting webhook")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetWebhookTest(t *testing.T) {
	fmt.Println("Starting test TestGetWebhookTest")
	PresetLoginRadiusTestEnv()
	_, err := GetWebhookTest()
	if err != nil {
		t.Errorf("Error getting webhook")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetWebhookSubscribedURLs(t *testing.T) {
	fmt.Println("Starting test TestGetWebhookSubscribedURLs")
	PresetLoginRadiusTestEnv()
	_, err := GetWebhookSubscribedURLs("Register")
	if err != nil {
		t.Errorf("Error getting webhook")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteWebhookUnsubscribe(t *testing.T) {
	fmt.Println("Starting test TestDeleteWebhookUnsubscribe")
	PresetLoginRadiusTestEnv()
	webhook := WebhookTest{"https://www.google.ca", "Register"}
	_, err := DeleteWebhookUnsubscribe(webhook)
	if err != nil {
		t.Errorf("Error getting webhook")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
