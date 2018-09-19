package loginradius

import (
	"fmt"
	"testing"
	"time"
)

func setupRole(t *testing.T) (string, func(t *testing.T)) {
	t.Log("Setting up role")
	PresetLoginRadiusTestEnv()
	time := time.Now()
	roleName := time.Format("20060102150405")
	permissions := Permissions{true, true, true}
	roles := Roles{{roleName, permissions}}
	testRole := TestRole{roles}
	_, err := PostRolesCreate(testRole)
	if err != nil {
		t.Errorf("Error creating role")
		fmt.Println(err)
	}
	return roleName, func(t *testing.T) {
		t.Log("Tearing down test case")
		_, err2 := DeleteAccountRole(roleName)
		if err2 != nil {
			t.Errorf("Error deleting role")
			fmt.Println(err)
		}
	}
}

func TestPostRolesCreate(t *testing.T) {
	fmt.Println("Starting test TestPostRolesCreate")
	PresetLoginRadiusTestEnv()
	permissions := Permissions{true, true, true}
	roles := Roles{{"RoleName", permissions}}
	testRole := TestRole{roles}
	_, err := PostRolesCreate(testRole)
	if err != nil {
		t.Errorf("Error creating role")
		fmt.Println(err)
	}
	_, err2 := DeleteAccountRole("RoleName")
	if err2 != nil {
		t.Errorf("Error deleting role")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPostRolesCreateInvalid(t *testing.T) {
	fmt.Println("Starting test TestPostRolesCreateInvalid")
	PresetLoginRadiusTestEnv()
	invalid := InvalidBody{"invalid"}
	_, err := PostRolesCreate(invalid)
	if err == nil {
		t.Errorf("Should be error")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetContextRolesPermissions(t *testing.T) {
	fmt.Println("Starting test TestGetContextRolesPermissions")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	_, err := GetContextRolesPermissions(testuid)
	if err != nil {
		t.Errorf("Error getting context roles permissions")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetRolesList(t *testing.T) {
	fmt.Println("Starting test TestGetRolesList")
	_, teardownTestCase := setupRole(t)
	defer teardownTestCase(t)
	_, err := GetRolesList()
	if err != nil {
		t.Errorf("Error getting roles list")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestGetRolesByUID(t *testing.T) {
	fmt.Println("Starting test TestGetRolesByUID")
	_, _, testuid, _, teardownTestCase := setupAccount(t)
	defer teardownTestCase(t)
	_, err := GetRolesByUID(testuid)
	if err != nil {
		t.Errorf("Error getting roles for user")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutAccountAddPermissionsToRole(t *testing.T) {
	fmt.Println("Starting test TestPutAccountAddPermissionsToRole")
	roleName, teardownTestCase := setupRole(t)
	defer teardownTestCase(t)
	permissions := PermissionList{[]string{"permission1", "permission2"}}
	_, err := PutAccountAddPermissionsToRole(roleName, permissions)
	if err != nil {
		t.Errorf("Error getting roles for user")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutAccountAddPermissionsToRoleInvalid(t *testing.T) {
	fmt.Println("Starting test TestPutAccountAddPermissionsToRoleInvalid")
	roleName, teardownTestCase := setupRole(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"invalid"}
	_, err := PutAccountAddPermissionsToRole(roleName, invalid)
	if err == nil {
		t.Errorf("Should be error")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutRolesAssignToUser(t *testing.T) {
	fmt.Println("Starting test TestPutRolesAssignToUser")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	roleName, teardownRole := setupRole(t)
	defer teardownRole(t)
	roles := RoleList{[]string{roleName}}
	_, err := PutRolesAssignToUser(testuid, roles)
	if err != nil {
		t.Errorf("Error setting role for user")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutRolesAssignToUserInvalid(t *testing.T) {
	fmt.Println("Starting test TestPutRolesAssignToUserInvalid")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	invalid := InvalidBody{"invalid"}
	_, err := PutRolesAssignToUser(testuid, invalid)
	if err == nil {
		t.Errorf("Should be error")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutRolesUpsertContext(t *testing.T) {
	fmt.Println("Starting test TestPutRolesUpsertContext")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	roleName, teardownRole := setupRole(t)
	defer teardownRole(t)
	roleContext := RoleContext{"contextTest", []string{roleName}, []string{"permission1"}, ""}
	roleContextContainer := RoleContextContainer{[]RoleContext{roleContext}}
	_, err := PutRolesUpsertContext(testuid, roleContextContainer)
	if err != nil {
		t.Errorf("Error setting role context for user")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestPutRolesUpsertContextInvalid(t *testing.T) {
	fmt.Println("Starting test TestPutRolesUpsertContextInvalid")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	invalid := InvalidBody{"invalid"}
	_, err := PutRolesUpsertContext(testuid, invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}

func TestDeleteAccountRole(t *testing.T) {
	fmt.Println("Starting test TestDeleteAccountRole")
	roleName, _ := setupRole(t)
	_, err := DeleteAccountRole(roleName)
	if err != nil {
		t.Errorf("Error deleting role")
		fmt.Println(err)
	}
	fmt.Println("Test complete")
}

func TestDeleteRolesAssignedToUser(t *testing.T) {
	fmt.Println("Starting test TestDeleteRolesAssignedToUser")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	roleName, teardownRole := setupRole(t)
	defer teardownRole(t)
	roles := RoleList{[]string{roleName}}
	_, err := PutRolesAssignToUser(testuid, roles)
	if err != nil {
		t.Errorf("Error setting role for user")
		fmt.Println(err)
	}
	_, err2 := DeleteRolesAssignedToUser(testuid, roles)
	if err2 != nil {
		t.Errorf("Error deleting role for user")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestDeleteRolesAssignedToUserInvalid(t *testing.T) {
	fmt.Println("Starting test TestDeleteRolesAssignedToUserInvalid")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	invalid := InvalidBody{"invalid"}
	_, err := PutRolesAssignToUser(testuid, invalid)
	if err == nil {
		t.Errorf("Error should be error")
	}
	fmt.Println("Test complete")
}

func TestDeleteRolesAccountRemovePermissions(t *testing.T) {
	fmt.Println("Starting test TestDeleteRolesAccountRemovePermissions")
	roleName, teardownTestCase := setupRole(t)
	defer teardownTestCase(t)
	permissions := PermissionList{[]string{"permission1", "permission2"}}
	_, err := PutAccountAddPermissionsToRole(roleName, permissions)
	if err != nil {
		t.Errorf("Error adding permissions to role")
		fmt.Println(err)
	}
	_, err2 := DeleteRolesAccountRemovePermissions(roleName, permissions)
	if err2 != nil {
		t.Errorf("Error deleting permissions from role")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestDeleteRolesAccountRemovePermissionsInvalid(t *testing.T) {
	fmt.Println("Starting test TestDeleteRolesAccountRemovePermissions")
	roleName, teardownTestCase := setupRole(t)
	defer teardownTestCase(t)
	invalid := InvalidBody{"invalid"}
	_, err := PutAccountAddPermissionsToRole(roleName, invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}

func TestDeleteContextFromRole(t *testing.T) {
	fmt.Println("Starting test TestDeleteContextFromRole")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	roleName, teardownRole := setupRole(t)
	defer teardownRole(t)
	roleContext := RoleContext{"contextTest", []string{roleName}, []string{"permission1"}, ""}
	roleContextContainer := RoleContextContainer{[]RoleContext{roleContext}}
	_, err := PutRolesUpsertContext(testuid, roleContextContainer)
	if err != nil {
		t.Errorf("Error adding contexts and roles to user")
		fmt.Println(err)
	}
	_, err2 := DeleteContextFromRole(testuid, "contextTest")
	if err2 != nil {
		t.Errorf("Error deleting role context")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestDeleteContextFromRoleInvalid(t *testing.T) {
	fmt.Println("Starting test TestDeleteContextFromRoleInvalid")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	invalid := InvalidBody{"invalid"}
	_, err := PutRolesUpsertContext(testuid, invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}

func TestDeleteRoleFromContext(t *testing.T) {
	fmt.Println("Starting test TestDeleteRoleFromContext")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	roleName, teardownRole := setupRole(t)
	defer teardownRole(t)
	roleContext := RoleContext{"contextTest", []string{roleName}, []string{"permission1"}, ""}
	roleContextContainer := RoleContextContainer{[]RoleContext{roleContext}}
	roles := RoleList{[]string{roleName}}
	_, err := PutRolesUpsertContext(testuid, roleContextContainer)
	if err != nil {
		t.Errorf("Error adding contexts and roles to user")
		fmt.Println(err)
	}
	_, err2 := DeleteRoleFromContext(testuid, "contextTest", roles)
	if err2 != nil {
		t.Errorf("Error deleting role context")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestDeleteRoleFromContextInvalid(t *testing.T) {
	fmt.Println("Starting test TestDeleteRoleFromContextInvalid")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	invalid := InvalidBody{"invalid"}
	_, err := PutRolesUpsertContext(testuid, invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}

func TestDeleteAdditionalPermissionFromContext(t *testing.T) {
	fmt.Println("Starting test TestDeleteAdditionalPermissionFromContext")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	roleName, teardownRole := setupRole(t)
	defer teardownRole(t)
	roleContext := RoleContext{"contextTest", []string{roleName}, []string{"permission1"}, ""}
	roleContextContainer := RoleContextContainer{[]RoleContext{roleContext}}
	permissions := DeletePermissionList{[]string{"permission1"}}
	_, err := PutRolesUpsertContext(testuid, roleContextContainer)
	if err != nil {
		t.Errorf("Error adding contexts and roles to user")
		fmt.Println(err)
	}
	_, err2 := DeleteAdditionalPermissionFromContext(testuid, "contextTest", permissions)
	if err2 != nil {
		t.Errorf("Error deleting role context")
		fmt.Println(err2)
	}
	fmt.Println("Test complete")
}

func TestDeleteAdditionalPermissionFromContextInvalid(t *testing.T) {
	fmt.Println("Starting test TestDeleteAdditionalPermissionFromContextInvalid")
	_, _, testuid, _, teardownAccount := setupAccount(t)
	defer teardownAccount(t)
	invalid := InvalidBody{"invalid"}
	_, err := PutRolesUpsertContext(testuid, invalid)
	if err == nil {
		t.Errorf("Should be error")
	}
	fmt.Println("Test complete")
}
