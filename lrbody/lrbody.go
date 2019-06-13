// The lrbody package holds the structs to be encoded as the body in POST and PUT calls
// These structs are meant to serve as convenient measures assisting API calls provided by the Loginradius
// Go SDK
// All functions in this SDK takes interface{} as the body, but initiating your
// data in the appropriate struct and passing in place of the body when calling the SDK functions
// will ensure the parameters submitted are correctly formatted and named for the LoginRadius APIs
// The usage of the structs in this package is optional and provided for convenience only
// Majority of methods take map[string]string as body parameter as well.
// These structs provide reference only, and do not include optional parameters
package lrbody

// Used by PostAuthUserRegistrationByEmail
type AuthEmail struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

// Used by PostAuthUserRegistrationByEmail
type RegistrationUser struct {
	Email    []AuthEmail `json:"Email"`
	Password string      `json:"Password"`
}

// Used by PostAuthLoginByEmail
type EmailLogin struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// Used by PostAuthLoginByUsername
type UsernameLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Used by PutAuthUpdateProfileByToken, PutManageAccountUpdate
// This serves as an example of struct that can be passed as the body for these endpoints
// Please adjust accordingly based on the field that needs to be updated
// For more information refer to https://www.loginradius.com/docs/api/v2/customer-identity-api/account/account-update
type UpdateProfile struct {
	Firstname string `json:"FirstName"`
	Lastname  string `json:"LastName"`
	Gender    string `json:"Gender"`
	Birthdate string `json:"BirthDate"`
	Country   string `json:"Country"`
	City      string `json:"City"`
	State     string `json:"State"`
}

// Used by PostAuthAddEmail
type AddEmail struct {
	Type  string `json:"type"`
	Email string `json:"email"`
}

// Used by PostAuthForgotPassword, PutResendEmailVerification, DeleteAuthRemoveEmail
type EmailStr struct {
	Email string
}

// Used by PutAuthUpdateSecurityQuestionByEmail
type SecurityQuestionAnswer struct {
	SecurityAnswer string `json:"2acec20722394dc3bd6362ef27df824e"`
}

// used by PutAuthChangePassword
type ChangePassword struct {
	OldPw string `json:"oldpassword"`
	NewPw string `json:"newpassword"`
}

// used by PutAuthLinkSocialIdentities
type LinkeSocialIds struct {
	CandidateToken string `json:"candidatetoken"`
}

// used by PutAuthResetPasswordByResetToken
type ResetPw struct {
	ResetToken string `json:"resettoken"`
	Password   string `json:"password"`
}

// used by PutAuthResetPasswordByOTP
type ResetPwOtp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
}

// used by PutAuthResetPasswordBySecurityAnswerAndEmail
type ResetPwSecurityQuestionEmail struct {
	Email          string `json:"email"`
	SecurityAnswer string `json:"securityanswer"`
	Password       string `json:"Password"`
}

// used by PutAuthResetPasswordBySecurityAnswerAndPhone
type ResetPwSecurityQuestionPhone struct {
	Phone          string `json:"phone"`
	SecurityAnswer string `json:"securityanswer"`
	Password       string `json:"password"`
}

// used by PutAuthResetPasswordBySecurityAnswerAndUsername
type ResetPwSecurityQuestionUsername struct {
	Username       string `json:"username"`
	SecurityAnswer string `json:"securityanswer"`
	Password       string `json:"password"`
}

type AuthUsername struct {
	Username string `json:"username"`
}

// Used by PostManageAccountCreate
type EmailArray []struct {
	Type  string
	Value string
}

// Used by PostManageAccountCreate
type AccountCreate struct {
	Email    EmailArray
	Password string
}

// Used by PostManageEmailVerificationToken, PostManageForgotPasswordToken
type Email struct {
	Email string `json:"Email"`
}

// Used by PostManageForgotPasswordToken
type Username struct {
	Username string `json:"Username"`
}

// Used by PutManageAccountUpdateSecurityQuestionConfig
type AccountSecurityQuestion struct {
	Securityquestionanswer accountSecurityQA `json:"securityquestionanswer"`
}

// The security question is identified by a random string key in the LoginRadius database
// You can retrieve this key with a call to GetConfiguration, and replace the
// json tag value with your question string
type accountSecurityQA struct {
	QuestionID string `json:"2acec20722394dc3bd6362ef27df824e"`
}

// Used by PutManageAccountSetPassword
type AccountPassword struct {
	Password string `json:"Password"`
}

// Used by PostRolesCreate
type Roles struct {
	Roles []Role `json:"roles"`
}

type Role struct {
	Name        string          `json:"name"`
	Permissions map[string]bool `json:"permissions"`
}

// Used by PutRolesAssignToUser
type RoleList struct {
	Roles []string `json:"roles"`
}

// Used by PutAccountAddPermissionsToRole, DeleteRolesAccountRemovePermissions
type PermissionList struct {
	Permissions []string `json:"permissions"`
}
