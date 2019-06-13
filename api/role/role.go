package role

import (
	"fmt"

	"github.com/LoginRadius/go-sdk/httprutils"
)

// PostRolesCreate creates a role with permissions.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/roles-create

// Required query parameters: apikey, apisecret

// Required post parameter: roles - array

// Pass data in struct lrbody.Roles as body to help ensure parameters satisfy API requirements; alternatively,
// []byte could also be passed as body
func (lr Loginradius) PostRolesCreate(body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPostReq("/identity/v2/manage/role", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteAccountRole is used to delete the role.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/delete-role

// Required template parameter: role - string representing the rolename of the role to be deleted
func (lr Loginradius) DeleteAccountRole(role string) (*httprutils.Response, error) {
	req := lr.Client.NewDeleteReq("/identity/v2/manage/role/" + role)
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetContextRolesPermissions gets the contexts that have been configured and the associated roles and permissions.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/get-context

// Required template parameter: uid - string representing uid of the user
func (lr Loginradius) GetContextRolesPermissions(uid string) (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/identity/v2/manage/account/" + uid + "/rolecontext")
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetRolesList retrieves the complete list of created roles with permissions of your app.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/roles-list
func (lr Loginradius) GetRolesList() (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/identity/v2/manage/role")
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// GetRolesByUID is used to retrieve all the assigned roles of a particular User.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/get-roles-by-uid

// Required template parameter: uid - string representing user's uid
func (lr Loginradius) GetRolesByUID(uid string) (*httprutils.Response, error) {
	req := lr.Client.NewGetReq("/identity/v2/manage/account/" + uid + "/role")
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutAccountAddPermissionsToRole is used to add permissions to the role.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/add-permissions-to-role

// Required template parameter: role - string representing role name

// Post parameters: permissions - array of permission names to be added to the role

// Pass data in struct lrbody.PermissionList as body to help ensure parameters satisfy API requirements; alternatively,
// []byte could also be passed as body
func (lr Loginradius) PutAccountAddPermissionsToRole(role string, body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReq("/identity/v2/manage/role/"+role+"/permission", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutRolesAssignToUser is used to assign created roles to the user.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/assign-roles-by-uid

// Required template parameter: uid - string representing uid of the user

// Required post parameter: roles - array of string(s) representing role name(s)

// Pass data in struct lrbody.RoleList as body to help ensure parameters satisfy API requirements; alternatively,
// []byte could also be passed as body
func (lr Loginradius) PutRolesAssignToUser(uid string, body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid+"/role", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// PutRolesUpsertContext creates a Context with a set of Roles.

// Documentation https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/upsert-context

// Required template parameter: uid - string representing user's UID

// Required post parameters: rolecontext - array of object containing role contexts

// Rolecontext object must contain: context - string; roles: array of strings representing role names; additionalpermissions: array of strings
// representing additional permissions; expiration: date of expiration of role context, format mm/dd/yyyy h:m:s
func (lr Loginradius) PutRolesUpsertContext(uid string, body interface{}) (*httprutils.Response, error) {
	req, err := lr.Client.NewPutReq("/identity/v2/manage/account/"+uid+"/rolecontext", body)
	if err != nil {
		return nil, err
	}
	lr.Client.AddApiCredentialsToReqHeader(req)
	fmt.Println(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteRolesAssignedToUser is used to unassign roles to the user.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/unassign-roles-by-uid

// Required template parameter: uid - string representing user's uid

// Required body parameters: roles - array of role name(strings)

// Sample body parameter: map[string][]string{"roles":[]string{"role1", "role2"}} or []byte{`{"roles":["role1", "role2"]}`}
func (lr Loginradius) DeleteRolesAssignedToUser(uid string, body interface{}) (*httprutils.Response, error) {
	req := lr.Client.NewDeleteReq("/identity/v2/manage/account/"+uid+"/role", body)
	req.Headers = httprutils.JSONHeader
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteRolesAccountRemovePermissions is used to remove permissions to the role.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/remove-permissions

// Required template parameter: role name - string representing role name for which the permissions are to be removed

// Required body parameter: permissions - array of permission names to be removed

// Pass data in struct lrbody.PermissionList as body to help ensure parameters satisfy API requirements; alternatively,
// []byte could also be passed as body
func (lr Loginradius) DeleteRolesAccountRemovePermissions(roleName string, body interface{}) (*httprutils.Response, error) {
	req := lr.Client.NewDeleteReq("/identity/v2/manage/role/"+roleName+"/permission", body)
	req.Headers = httprutils.JSONHeader
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteContextFromRole deletes the specified Role Context

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/delete-context

// Required template parameter: uid - string preresenting user's uid; rolecontextname - string representing the name of the role context
// to be deleted
func (lr Loginradius) DeleteContextFromRole(uid, rolecontextname string) (*httprutils.Response, error) {
	req := lr.Client.NewDeleteReq("/identity/v2/manage/account/" + uid + "/rolecontext/" + rolecontextname)
	req.Headers = httprutils.JSONHeader
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteRoleFromContext deletes the specified Role Context

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/delete-context

// Required template parameter: uid - string preresenting user's uid; rolecontextname - string representing the name of the context
// from which the role(s) is(are) to be deleted

// Required body parameters: roles - array of strings representing the role name(s) to be deleted
func (lr Loginradius) DeleteRoleFromContext(uid, rolecontextname string, body interface{}) (*httprutils.Response, error) {
	req := lr.Client.NewDeleteReq("/identity/v2/manage/account/"+uid+"/rolecontext/"+rolecontextname+"/role", body)
	req.Headers = httprutils.JSONHeader
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}

// DeleteAdditionalPermissionFromContext deletes Additional Permissions from Context.

// Documentation: https://www.loginradius.com/docs/api/v2/customer-identity-api/roles-management/delete-permissions-from-context

// Required template parameters: uid - string representing user's uid; rolecontextname - name of the context from which the additional
// permission(s) is(are) to be delted

// Required post parameters: additionalpermissions - array of strings representing names of additional permissions
func (lr Loginradius) DeleteAdditionalPermissionFromContext(uid, rolecontextname string, body interface{}) (*httprutils.Response, error) {
	req := lr.Client.NewDeleteReq("/identity/v2/manage/account/"+uid+"/rolecontext/"+rolecontextname+"/additionalpermission", body)
	req.Headers = httprutils.JSONHeader
	lr.Client.AddApiCredentialsToReqHeader(req)
	res, err := httprutils.TimeoutClient.Send(*req)
	return res, err
}
