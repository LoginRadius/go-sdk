// +build !mfa

package loginradius

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type TestCustomObject struct {
	Custom1 string
	Custom2 string
}

func setupCustomObject(t *testing.T) (string, string, string, string, func(t *testing.T)) {
	t.Log("Setting up test case")
	_, _, testuid, _, accessToken, teardownTestCase := setupLogin(t)

	objName := os.Getenv("CUSTOMOBJECTNAME")
	time := time.Now()
	timestamp := time.Format("20060102150405")
	customObj := TestCustomObject{timestamp + "0", timestamp + "1"}
	object, err := PostCustomObjectCreateByUID(objName, testuid, customObj)
	if err != nil {
		t.Errorf("Error creating custom object")
		fmt.Println(err)
	}
	return accessToken, testuid, object.ID, objName, func(t *testing.T) {
		t.Log("Tearing down test case")
		defer teardownTestCase(t)
		_, err2 := DeleteCustomObjectByObjectRecordIDAndUID(objName, testuid, object.ID)
		if err2 != nil {
			t.Errorf("Error deleting custom object")
			fmt.Println(err2)
		}
	}
}

func TestPostCustomObjectCreateByUID(t *testing.T) {
	fmt.Println("Starting test TestPostCustomObjectCreateByUID")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	objName := os.Getenv("CUSTOMOBJECTNAME")
	time := time.Now()
	timestamp := time.Format("20060102150405")
	customObj := TestCustomObject{timestamp + "0", timestamp + "1"}
	session, err := PostCustomObjectCreateByUID(objName, testuid, customObj)
	if err != nil || session.ID == "" {
		t.Errorf("Error creating custom object")
		fmt.Println(err)
	}
	_, err2 := DeleteCustomObjectByObjectRecordIDAndUID(objName, testuid, session.ID)
	if err2 != nil {
		t.Errorf("Error deleting custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostCustomObjectCreateByToken(t *testing.T) {
	fmt.Println("Starting test TestPostCustomObjectCreateByUID")
	_, _, _, _, accessToken, teardownTestCase := setupLogin(t)
	defer teardownTestCase(t)
	objName := os.Getenv("CUSTOMOBJECTNAME")
	time := time.Now()
	timestamp := time.Format("20060102150405")
	customObj := TestCustomObject{timestamp + "0", timestamp + "1"}
	session, err := PostCustomObjectCreateByToken(objName, accessToken, customObj)
	if err != nil || session.ID == "" {
		t.Errorf("Error creating custom object")
		fmt.Println(err)
	}
	_, err2 := DeleteCustomObjectByObjectRecordIDAndToken(objName, accessToken, session.ID)
	if err2 != nil {
		t.Errorf("Error deleting custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetCustomObjectByObjectRecordIDAndUID(t *testing.T) {
	fmt.Println("Starting test TestGetCustomObjectByObjectRecordIDAndUID")
	_, testuid, objectID, objName, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	session, err := GetCustomObjectByObjectRecordIDAndUID(objName, testuid, objectID)
	if err != nil || session.ID != objectID {
		t.Errorf("Error getting custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetCustomObjectByObjectRecordIDAndToken(t *testing.T) {
	fmt.Println("Starting test TestGetCustomObjectByObjectRecordIDAndToken")
	accessToken, _, objectID, objName, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	session, err := GetCustomObjectByObjectRecordIDAndToken(objName, accessToken, objectID)
	if err != nil || session.ID != objectID {
		t.Errorf("Error getting custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetCustomObjectByToken(t *testing.T) {
	fmt.Println("Starting test TestGetCustomObjectByToken")
	accessToken, _, _, objName, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	session, err := GetCustomObjectByToken(objName, accessToken)
	if err != nil || session.Count != 1 {
		t.Errorf("Error getting custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetCustomObjectByUID(t *testing.T) {
	fmt.Println("Starting test TestGetCustomObjectByUID")
	_, testuid, _, objName, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	session, err := GetCustomObjectByUID(objName, testuid)
	if err != nil || session.Count != 1 {
		t.Errorf("Error getting custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutCustomObjectUpdateByUID(t *testing.T) {
	fmt.Println("Starting test TestPutCustomObjectUpdateByUID")
	_, testuid, objectID, objName, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	customObj := TestCustomObject{objName + "3", objName + "4"}
	_, err := PutCustomObjectUpdateByUID(objName, "replace", testuid, objectID, customObj)
	if err != nil {
		t.Errorf("Error updating custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutCustomObjectUpdateByToken(t *testing.T) {
	fmt.Println("Starting test TestPutCustomObjectUpdateByToken")
	accessToken, _, objectID, objName, teardownTestCase := setupCustomObject(t)
	defer teardownTestCase(t)
	customObj := TestCustomObject{objName + "3", objName + "4"}
	_, err := PutCustomObjectUpdateByToken(objName, "replace", accessToken, objectID, customObj)
	if err != nil {
		t.Errorf("Error updating custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteCustomObjectByObjectRecordIDAndUID(t *testing.T) {
	fmt.Println("Starting test TestDeleteCustomObjectByObjectRecordIDAndUID")
	_, testuid, objectID, objName, _ := setupCustomObject(t)
	session, err := DeleteCustomObjectByObjectRecordIDAndUID(objName, testuid, objectID)
	if err != nil || session.IsDeleted == false {
		t.Errorf("Error updating custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteCustomObjectByObjectRecordIDAndToken(t *testing.T) {
	fmt.Println("Starting test TestDeleteCustomObjectByObjectRecordIDAndToken")
	accessToken, _, objectID, objName, _ := setupCustomObject(t)
	session, err := DeleteCustomObjectByObjectRecordIDAndToken(objName, accessToken, objectID)
	if err != nil || session.IsDeleted == false {
		t.Errorf("Error updating custom object")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}
