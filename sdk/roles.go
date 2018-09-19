package loginradius

import (
	"os"
	"time"
)

// Role is the struct used to contain information from the roles management module
// Data from Roles Create, Roles List, Account Add Permissions to Role, Account Delete
// Role, Roles UnAssign to User, Account Remove Permissions, Delete Context,
// Delete Role from Context, Delete Additional Permission from Context
type Role struct {
	IsDeleted bool
	Data      []struct {
		Name                  string          `json:"Name"`
		Permissions           map[string]bool `json:"Permissions"`
		AdditionalPermissions []string        `json:"AdditionalPermissions"`
		Expiration            time.Time       `json:"Expiration"`
	} `json:"data"`
	Count       int    `json:"Count"`
	Name        string `json:"Name"`
	Permissions map[string]bool
}

// ContextRole is the struct used for APIs involving context
// Data from Get Context, Upsert Context use this struct
type ContextRole struct {
	Data []struct {
		Context               string    `json:"Context"`
		Roles                 []string  `json:"Roles"`
		AdditionalPermissions []string  `json:"AdditionalPermissions"`
		Expiration            time.Time `json:"Expiration"`
	} `json:"Data"`
}

// RoleArray contains an array of roles
// Roles by UID and Roles Assign to User use this to contain data
type RoleArray struct {
	Roles []string `json:"Roles"`
}

// PostRolesCreate creates a role with permissions.
// Post Parameter is an array of roles
func PostRolesCreate(body interface{}) (Role, error) {
	data := new(Role)
	req, reqErr := CreateRequest("POST", os.Getenv("DOMAIN") + "/identity/v2/manage/role", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// GetContextRolesPermissions gets the contexts that have been configured and the associated roles and permissions.
func GetContextRolesPermissions(uid string) (ContextRole, error) {
	data := new(ContextRole)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/rolecontext", "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetRolesList retrieves the complete list of created roles with permissions of your app.
func GetRolesList() (Role, error) {
	data := new(Role)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/manage/role", "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetRolesByUID is used to retrieve all the assigned roles of a particular User.
func GetRolesByUID(uid string) (RoleArray, error) {
	data := new(RoleArray)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/role", "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// PutAccountAddPermissionsToRole is used to add permissions to the role.
// Post Parameters are permissions: string
func PutAccountAddPermissionsToRole(role string, body interface{}) (Role, error) {
	data := new(Role)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/manage/role/"+role+"/permission", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutRolesAssignToUser is used to assign created roles to the user.
// Post Parameters is an array of roles
func PutRolesAssignToUser(uid string, body interface{}) (RoleArray, error) {
	data := new(RoleArray)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/role", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// PutRolesUpsertContext creates a Context with a set of Roles.
// Post Parameters are rolecontext: string, context: string, roles: string, additionalpermissions: string and an
// optional expiration: time.Time
func PutRolesUpsertContext(uid string, body interface{}) (ContextRole, error) {
	data := new(ContextRole)
	req, reqErr := CreateRequest("PUT", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/rolecontext", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// DeleteAccountRole is used to delete the role.
func DeleteAccountRole(role string) (Role, error) {
	data := new(Role)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/role/"+role, "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// DeleteRolesAssignedToUser is used to unassign roles to the user.
// Post Parameter is an array of roles
func DeleteRolesAssignedToUser(uid string, body interface{}) (Role, error) {
	data := new(Role)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/role", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// DeleteRolesAccountRemovePermissions is used to remove permissions to the role.
// Post Parameter is the permissions from which you want to remove the role
func DeleteRolesAccountRemovePermissions(roleName string, body interface{}) (Role, error) {
	data := new(Role)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/role/"+roleName+"/permission", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// DeleteContextFromRole deletes the specified Role Context
func DeleteContextFromRole(uid, rolecontextname string) (Role, error) {
	data := new(Role)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+"/rolecontext/"+rolecontextname, "")
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// DeleteRoleFromContext deletes the specified Role from a Context.
// Post Parameters is an array of roles
func DeleteRoleFromContext(uid, rolecontextname string, body interface{}) (Role, error) {
	data := new(Role)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+
		"/rolecontext/"+rolecontextname+"/role", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}

// DeleteAdditionalPermissionFromContext deletes Additional Permissions from Context.
// Post Parameter is the array of strings which represent additional permissions
func DeleteAdditionalPermissionFromContext(uid, rolecontextname string, body interface{}) (Role, error) {
	data := new(Role)
	req, reqErr := CreateRequest("DELETE", os.Getenv("DOMAIN") + "/identity/v2/manage/account/"+uid+
		"/rolecontext/"+rolecontextname+"/additionalpermission", body)
	if reqErr != nil {
		return *data, reqErr
	}

	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/json")

	err := RunRequest(req, data)
	return *data, err
}
