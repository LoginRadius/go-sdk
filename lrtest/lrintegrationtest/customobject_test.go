package lrintegrationtest

import (
	"os"
	"testing"
	"time"

	"github.com/LoginRadius/go-sdk/api/customobject"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

func TestPostCustomObjectCreateByUID(t *testing.T) {
	_, _, uid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	objName := os.Getenv("CUSTOMOBJECTNAME")
	timestamp := time.Now().Format("20060102150405")
	customObj := map[string]string{
		"custom1": timestamp + "0",
		"custom2": timestamp + "1",
	}
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByUID(
		uid,
		map[string]string{"objectname": objName},
		customObj,
	)
	if err != nil {
		t.Errorf("Error calling PostCustomObjectCreateByUID: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	id := obj["Id"].(string)
	if err != nil || id == "" {
		t.Errorf("Error returned from PostCustomObjectCreateByUID: %v", err)
	}

	resp, err = customobject.Loginradius(customobject.Loginradius{lrclient}).DeleteCustomObjectByObjectRecordIDAndUID(
		uid,
		id,
		map[string]string{"objectname": objName},
	)

	if err != nil {
		t.Errorf("Error calling DeleteCustomObjectByObjectRecordIDAndUID while cleaning up for PostCustomObjectCreateByUID: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteCustomObjectByObjectRecordIDAndUID while cleaning up for PostCustomObjectCreateByUID: %v", err)
	}
}

func TestPostCustomObjectCreateByToken(t *testing.T) {
	_, _, uid, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	objName := os.Getenv("CUSTOMOBJECTNAME")
	timestamp := time.Now().Format("20060102150405")
	customObj := map[string]string{
		"custom1": timestamp + "0",
		"custom2": timestamp + "1",
	}
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByToken(
		map[string]string{"objectname": objName},
		customObj,
	)
	if err != nil {
		t.Errorf("Error calling PostCustomObjectCreateByToken: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	id := obj["Id"].(string)
	if err != nil || id == "" {
		t.Errorf("Error returned from PostCustomObjectCreateByToken: %v", err)
	}

	resp, err = customobject.Loginradius(customobject.Loginradius{lrclient}).DeleteCustomObjectByObjectRecordIDAndUID(
		uid,
		id,
		map[string]string{"objectname": objName},
	)

	if err != nil {
		t.Errorf("Error calling DeleteCustomObjectByObjectRecordIDAndUID while cleaning up for PostCustomObjectCreateByToken: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteCustomObjectByObjectRecordIDAndUID while cleaning up for PostCustomObjectCreateByToken: %v", err)
	}
}

func TestGetCustomObjectByObjectRecordIDAndUID(t *testing.T) {
	_, uid, objectID, objName, lrclient, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).GetCustomObjectByObjectRecordIDAndUID(
		uid,
		objectID,
		map[string]string{"objectname": objName},
	)
	if err != nil {
		t.Errorf("Error calling GetCustomObjectByObjectRecordIDAndUID: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || obj["Id"].(string) != objectID {
		t.Errorf("Error returned from GetCustomObjectByObjectRecordIDAndUID: %v", err)
	}
}

func TestGetCustomObjectByObjectRecordIDAndToken(t *testing.T) {
	_, _, objectID, objName, lrclient, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).GetCustomObjectByObjectRecordIDAndToken(
		objectID,
		map[string]string{"objectname": objName},
	)
	if err != nil {
		t.Errorf("Error calling GetCustomObjectByObjectRecordIDAndToken: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil || obj["Id"].(string) != objectID {
		t.Errorf("Error returned from GetCustomObjectByObjectRecordIDAndToken: %v", err)
	}
}
func TestGetCustomObjectByToken(t *testing.T) {
	_, _, objectID, objName, lrclient, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).GetCustomObjectByToken(
		map[string]string{"objectname": objName},
	)
	if err != nil {
		t.Errorf("Error calling GetCustomObjectByToken: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	data := obj["data"].([]interface{})
	if err != nil || data[0].(map[string]interface{})["Id"].(string) != objectID {
		t.Errorf("Error returned from GetCustomObjectByToken: %v", err)
	}
}

func TestGetCustomObjectByUID(t *testing.T) {
	_, uid, objectID, objName, lrclient, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).GetCustomObjectByUID(
		uid,
		map[string]string{"objectname": objName},
	)
	if err != nil {
		t.Errorf("Error calling GetCustomObjectByUID: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	data := obj["data"].([]interface{})
	if err != nil || data[0].(map[string]interface{})["Id"].(string) != objectID {
		t.Errorf("Error returned from GetCustomObjectByUID: %v", err)
	}
}

func TestPutCustomObjectUpdateByUID(t *testing.T) {
	_, uid, objectId, objName, lrclient, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	customObj := map[string]string{
		"custom1": "value1",
		"custom2": "value2",
	}
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PutCustomObjectUpdateByUID(
		uid,
		objectId,
		map[string]string{"objectname": objName, "updatetype": "replace"},
		customObj,
	)
	if err != nil {
		t.Errorf("Error calling PutCustomObjectUpdateByUID: %v", err)
	}
	unmarshalled, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil {
		t.Errorf("Error returned from PutCustomObjectUpdateByUID: %v", err)
	}
	returnedObj := unmarshalled["CustomObject"].(map[string]interface{})
	for k, v := range returnedObj {
		if v != customObj[k] {
			t.Errorf("PutCustomObjectUpdateByUID was supposed to update custom object to %v, got %v instead", customObj, returnedObj)
		}
	}
}

func TestPutCustomObjectUpdateByToken(t *testing.T) {
	_, _, objectId, objName, lrclient, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	customObj := map[string]string{
		"custom1": "value1",
		"custom2": "value2",
	}
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PutCustomObjectUpdateByToken(
		objectId,
		map[string]string{"objectname": objName, "updatetype": "replace"},
		customObj,
	)
	if err != nil {
		t.Errorf("Error calling PutCustomObjectUpdateByToken: %v", err)
	}
	unmarshalled, err := lrjson.DynamicUnmarshal(resp.Body)
	if err != nil {
		t.Errorf("Error returned from PutCustomObjectUpdateByToken: %v", err)
	}
	returnedObj := unmarshalled["CustomObject"].(map[string]interface{})
	for k, v := range returnedObj {
		if v != customObj[k] {
			t.Errorf("PutCustomObjectUpdateByToken was supposed to update custom object to %v, got %v instead", customObj, returnedObj)
		}
	}
}

func TestDeleteCustomObjectByObjectRecordIDAndToken(t *testing.T) {
	_, _, uid, _, _, lrclient, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	objName := os.Getenv("CUSTOMOBJECTNAME")
	timestamp := time.Now().Format("20060102150405")
	customObj := map[string]string{
		"custom1": timestamp + "0",
		"custom2": timestamp + "1",
	}
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByUID(
		uid,
		map[string]string{"objectname": objName},
		customObj,
	)
	if err != nil {
		t.Errorf("Error calling PostCustomObjectCreateByUID for DeleteCustomObjectByObjectRecordIDAndToken: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	id := obj["Id"].(string)
	if err != nil || obj["Id"].(string) == "" {
		t.Errorf("Error returned from PostCustomObjectCreateByUID for DeleteCustomObjectByObjectRecordIDAndToken: %v", err)
	}

	resp, err = customobject.Loginradius(customobject.Loginradius{lrclient}).DeleteCustomObjectByObjectRecordIDAndToken(
		id,
		map[string]string{"objectname": objName},
	)

	if err != nil {
		t.Errorf("Error calling DeleteCustomObjectByObjectRecordIDAndToken: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteCustomObjectByObjectRecordIDAndToken: %v", err)
	}
}

func TestDeleteCustomObjectByObjectRecordIDAndUID(t *testing.T) {
	_, _, uid, _, lrclient, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	objName := os.Getenv("CUSTOMOBJECTNAME")
	timestamp := time.Now().Format("20060102150405")
	customObj := map[string]string{
		"custom1": timestamp + "0",
		"custom2": timestamp + "1",
	}
	resp, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByUID(
		uid,
		map[string]string{"objectname": objName},
		customObj,
	)
	if err != nil {
		t.Errorf("Error calling PostCustomObjectCreateByUID for DeleteCustomObjectByObjectRecordIDAndUID: %v", err)
	}
	obj, err := lrjson.DynamicUnmarshal(resp.Body)
	id := obj["Id"].(string)
	if err != nil || obj["Id"].(string) == "" {
		t.Errorf("Error returned from PostCustomObjectCreateByUID for DeleteCustomObjectByObjectRecordIDAndUID: %v", err)
	}

	resp, err = customobject.Loginradius(customobject.Loginradius{lrclient}).DeleteCustomObjectByObjectRecordIDAndUID(uid, id, map[string]string{"objectname": objName})

	if err != nil {
		t.Errorf("Error calling DeleteCustomObjectByObjectRecordIDAndUID: %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(resp.Body)

	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteCustomObjectByObjectRecordIDAndUID: %v", err)
	}
}
