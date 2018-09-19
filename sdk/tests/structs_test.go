package loginradius

// Helper Function Test Structs
type Success struct {
	Name    string
	Hobbies []string
}

type Error struct {
	Name    string
	Message string
}

type AccountSetup struct {
	PhoneIDVerified bool
	EmailVerified   bool
	Email           TestEmailArr
	Password        string
	Username        string
	PhoneID         string
}

type InvalidBody struct {
	invalidType string
}

// Multi-module Test Structs
type TestEmailLogin struct {
	Email    string
	Password string
}

type TestUsernameLogin struct {
	Username string
	Password string
}

type TestPhoneLogin struct {
	Phone    string
	Password string
}

type TestEmail struct {
	Email string
}

type TestUsername struct {
	Username string
}

type TestEmailArr []struct {
	Type  string
	Value string
}

// Authentication Test Structs
type TestEmailCreator struct {
	Email string
	Type  string
}

type OTPVerify struct {
	Otp   string
	Email string
}

type PassChange struct {
	Oldpassword string
	Newpassword string
}

type PasswordReset struct {
	ResetToken string
	Password   string
}

type TestAccountRegister struct {
	Email    TestEmailArr
	Password string
}

type ResetWithEmailSecurity struct {
	SecurityAnswer             SecurityQuestion
	Email                      string
	Password                   string
	ResetPasswordEmailTemplate string
}

type ResetWithPhoneSecurity struct {
	SecurityAnswer             SecurityQuestion
	Phone                      string
	Password                   string
	ResetPasswordEmailTemplate string
}

type ResetWithUsernameSecurity struct {
	SecurityAnswer             SecurityQuestion
	Username                   string
	Password                   string
	ResetPasswordEmailTemplate string
}

type SecurityQuestionTest struct {
	SecurityQuestionAnswer SecurityQuestion `json:"securityquestionanswer"`
}

type CandidateToken struct {
	CandidateToken string `json:"candidatetoken"`
}

type Provider struct {
	Provider   string `json:"provider"`
	ProviderID string `json:"providerid"`
}

// Account API Test Structs
type TestAccount struct {
	EmailVerified bool
	Email         TestEmailArr
	Password      string
}

type TestPassword struct {
	Password string
}

type TestFirstName struct {
	FirstName string
}

type ProfileEmail struct {
	Email TestEmailArr
}

// Role Test Structs
type TestRole struct {
	Roles Roles
}

type Roles []struct {
	Name        string `json:"name"`
	Permissions Permissions
}

type TestUpdateRole struct {
	Name        string `json:"name"`
	Permissions Permissions
}

type Permissions struct {
	PermissionName1 bool `json:"permission_name1"`
	PermissionName2 bool `json:"permission_name2"`
	PermissionName3 bool `json:"permission_name3"`
}

type PermissionList struct {
	Permissions []string
}

type DeletePermissionList struct {
	AdditionalPermissions []string
}

type RoleList struct {
	Roles []string `json:"roles"`
}

type RoleContext struct {
	Context               string   `json:"context"`
	Roles                 []string `json:"roles"`
	Additionalpermissions []string `json:"additionalpermissions"`
	Expiration            string   `json:"expiration"`
}

type RoleContextContainer struct {
	RoleContext []RoleContext
}

// Phone Test Structs

type PhoneNumberArray []struct {
	PhoneType   string `json:"PhoneType"`
	PhoneNumber string `json:"PhoneNumber"`
}

type PhoneRegister struct {
	Email    TestEmailArr
	PhoneID  string
	Password string
}

type TestPhone struct {
	Phone string
}

type ResetPhonePassword struct {
	Phone    string
	OTP      string
	Password string
}

type UndoPhoneVerify struct {
	PhoneIDVerified bool
}

// Webhook test structs

type WebhookTest struct {
	TargetURL string
	Event     string
}
