package lrintegrationtest

import (
	"reflect"
	"testing"

	"github.com/LoginRadius/go-sdk/api/role"
	lrbody "github.com/LoginRadius/go-sdk/lrbody"
	lrjson "github.com/LoginRadius/go-sdk/lrjson"
)

func TestPostRolesCreate(t *testing.T) {
	_, _, _, tearDown := setupRole(t)
	defer tearDown(t)
}

func TestDeleteAccountRole(t *testing.T) {
	_, _, _, tearDown := setupRole(t)
	defer tearDown(t)
}

func TestGetContextRolesPermissions(t *testing.T) {
	_, _, uid, _, _, lrclient, tearDown := setupLogin(t)
	defer tearDown(t)
	_, err := role.Loginradius(role.Loginradius{lrclient}).GetContextRolesPermissions(uid)
	if err != nil {
		t.Errorf("Error calling GetContextRolesPermissions: %v", err)
	}
}

func TestGetRolesList(t *testing.T) {
	_, rolename, lrclient, tearDown := setupRole(t)
	defer tearDown(t)
	res, err := role.Loginradius(role.Loginradius{lrclient}).GetRolesList()
	if err != nil {
		t.Errorf("Error calling GetRolesList: %v", err)
	}
	roles, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil {
		t.Errorf("Error returned from GetRolesList: %v", err)
	}
	exists := false
	for _, r := range roles["data"].([]interface{}) {
		if r.(map[string]interface{})["Name"].(string) == rolename {
			exists = true
		}
	}

	if !exists {
		t.Errorf("Error returning created role %v from GetRolesList: %v", rolename, roles)
	}
}

func TestGetRolesByUID(t *testing.T) {
	_, _, uid, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)

	_, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesAssignToUser(
		uid,
		lrbody.RoleList{[]string{rolename}},
	)

	if err != nil {
		t.Errorf("Error calling PutRolesAssignToUser for GetRolesByUID%v", err)
	}

	res, err := role.Loginradius(role.Loginradius{lrclient}).GetRolesByUID(uid)
	if err != nil {
		t.Errorf("Error calling GetRolesByUID%v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reflect.DeepEqual(data["Roles"].([]interface{})[0].(string), rolename) {
		t.Errorf("Error returned from GetRolesByUID %v, %v", err, data)
	}
}

func TestPutAccountAddPermissionsToRole(t *testing.T) {
	_, _, _, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)

	permissionName := "example_permission_name"

	res, err := role.Loginradius(role.Loginradius{lrclient}).PutAccountAddPermissionsToRole(
		rolename,
		lrbody.PermissionList{[]string{permissionName}},
	)
	if err != nil {
		t.Errorf("Error calling PutAccountAddPermissionsToRole: %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	included := false
	for k, _ := range data["Permissions"].(map[string]interface{}) {
		if k == permissionName {
			included = true
		}
	}
	if err != nil || !included {
		t.Errorf("Error returned from PutAccountAddPermissionsToRole %v, %v", err, data)
	}
}

func TestPutRolesAssignToUser(t *testing.T) {
	_, _, uid, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)

	res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesAssignToUser(
		uid,
		[]byte(`{"roles": ["`+rolename+`"]}`),
	)

	if err != nil {
		t.Errorf("Error calling PutRolesAssignToUser %v", err)
	}
	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reflect.DeepEqual(data["Roles"].([]interface{})[0].(string), rolename) {
		t.Errorf("Error returned from PutRolesAssignToUser %v, %v, %v", err, data["Roles"], []string{rolename})
	}
}

func TestPutRolesUpsertContext(t *testing.T) {
	_, _, uid, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)
	body := []byte(`{"rolecontext":[{"context":"example_context", "roles":["` + rolename + `"], "additionalpermissions":["permissionx", "permissiony"]}]}`)

	res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesUpsertContext(
		uid,
		body,
	)

	if err != nil {
		t.Errorf("Error calling PutRolesUpsertContext %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reflect.DeepEqual(data["Data"].([]interface{})[0].(map[string]interface{})["Context"].(string), "example_context") {
		t.Errorf("Error returned from PutRolesUpsertContext %v, %v, %v", err, data["Roles"], []string{rolename})
	}
}

func TestDeleteRolesAssignedToUser(t *testing.T) {
	_, _, uid, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)

	res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesAssignToUser(
		uid,
		[]byte(`{"roles": ["example_role_name"]}`),
	)

	if err != nil {
		t.Errorf("Error calling PutRolesAssignToUser for DeleteRolesAssignedToUser %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reflect.DeepEqual(data["Roles"].([]interface{})[0].(string), rolename) {
		t.Errorf("Error returned from PutRolesAssignToUser for DeleteRolesAssignedToUser: %v, %v, %v", err, data["Roles"], []string{rolename})
	}

	res, err = role.Loginradius(role.Loginradius{lrclient}).DeleteRolesAssignedToUser(
		uid,
		[]byte(`{"roles": ["example_role_name"]}`),
	)

	if err != nil {
		t.Errorf("Error calling DeleteRolesAssignedToUser %v", err)
	}

	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteRolesAssignedToUser: %v, %v", err, data)
	}

}

func TestDeleteRolesAccountRemovePermissions(t *testing.T) {
	_, _, _, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	permissionName, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)

	res, err := role.Loginradius(role.Loginradius{lrclient}).DeleteRolesAccountRemovePermissions(
		rolename,
		lrbody.PermissionList{[]string{permissionName}},
	)

	if err != nil {
		t.Errorf("Error calling DeleteRolesAccountRemovePermissions: %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	included := false
	for k, _ := range data["Permissions"].(map[string]interface{}) {
		if k == permissionName {
			included = true
		}
	}
	if err != nil || included {
		t.Errorf("Error returned from DeleteRolesAccountRemovePermission %v, %v", err, data)
	}
}

func TestDeleteContextFromRole(t *testing.T) {
	_, _, uid, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)
	body := []byte(`{"rolecontext":[{"context":"example_context", "roles":["` + rolename + `"], "additionalpermissions":["permissionx", "permissiony"]}]}`)

	res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesUpsertContext(
		uid,
		body,
	)

	if err != nil {
		t.Errorf("Error calling PutRolesUpsertContext for DeleteContextFromRole %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reflect.DeepEqual(data["Data"].([]interface{})[0].(map[string]interface{})["Context"].(string), "example_context") {
		t.Errorf("Error returned from PutRolesUpsertContext for DeleteContextFromRole: %v, %v, %v", err, data["Roles"], []string{rolename})
	}

	res, err = role.Loginradius(role.Loginradius{lrclient}).DeleteContextFromRole(
		uid,
		"example_context",
	)

	if err != nil {
		t.Errorf("Error calling DeleteContextFromRole %v", err)
	}

	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteContextFromRole: %v, %v, %v", err, data["Roles"], []string{rolename})
	}
}

func TestDeleteRoleFromContext(t *testing.T) {
	_, _, uid, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)
	body := []byte(`{"rolecontext":[{"context":"example_context", "roles":["` + rolename + `"], "additionalpermissions":["permissionx", "permissiony"]}]}`)

	res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesUpsertContext(
		uid,
		body,
	)

	if err != nil {
		t.Errorf("Error calling PutRolesUpsertContext for DeleteRoleFromContext %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reflect.DeepEqual(data["Data"].([]interface{})[0].(map[string]interface{})["Context"].(string), "example_context") {
		t.Errorf("Error returned from PutRolesUpsertContext for DeleteRoleFromContext: %v, %v, %v", err, data["Roles"], []string{rolename})
	}

	res, err = role.Loginradius(role.Loginradius{lrclient}).DeleteRoleFromContext(
		uid,
		"example_context",
		[]byte(`{"roles":["`+rolename+`"]}`),
	)

	if err != nil {
		t.Errorf("Error calling DeleteRoleFromContext %v", err)
	}

	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteRoleFromContext: %v, %v, %v", err, data["Roles"], []string{rolename})
	}
}

func TestDeleteAdditionalPermissionFromContext(t *testing.T) {
	_, _, uid, _, lrclient, tearDownAccount := setupAccount(t)
	defer tearDownAccount(t)

	_, rolename, lrclient, tearDownRole := setupRole(t)
	defer tearDownRole(t)
	body := []byte(`{"rolecontext":[{"context":"example_context", "roles":["` + rolename + `"], "additionalpermissions":["permissionx", "permissiony"]}]}`)

	res, err := role.Loginradius(role.Loginradius{lrclient}).PutRolesUpsertContext(
		uid,
		body,
	)

	if err != nil {
		t.Errorf("Error calling PutRolesUpsertContext for DeleteAdditionalPermissionFromContext %v", err)
	}

	data, err := lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !reflect.DeepEqual(data["Data"].([]interface{})[0].(map[string]interface{})["Context"].(string), "example_context") {
		t.Errorf("Error returned from PutRolesUpsertContext for DeleteAdditionalPermissionFromContext: %v, %v, %v", err, data["Roles"], []string{rolename})
	}

	res, err = role.Loginradius(role.Loginradius{lrclient}).DeleteAdditionalPermissionFromContext(
		uid,
		"example_context",
		[]byte(`{"additionalpermissions":["permissionx"]}`),
	)

	if err != nil {
		t.Errorf("Error calling DeleteAdditionalPermissionFromContext %v", err)
	}

	data, err = lrjson.DynamicUnmarshal(res.Body)
	if err != nil || !data["IsDeleted"].(bool) {
		t.Errorf("Error returned from DeleteAdditionalPermissionFromContext: %v, %v, %v", err, data["Roles"], []string{rolename})
	}
}

// func TestDeleteAdditionalPermissionFromContextInvalid(t *testing.T) {
// 	fmt.Println("Starting test TestDeleteAdditionalPermissionFromContextInvalid")
// 	_, _, testuid, _, teardownAccount := setupAccount(t)
// 	defer teardownAccount(t)
// 	invalid := InvalidBody{"invalid"}
// 	_, err := PutRolesUpsertContext(testuid, invalid)
// 	if err == nil {
// 		t.Errorf("Should be error")
// 	}
// 	fmt.Println("Test complete")
// }
