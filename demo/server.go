package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/loginradius/go-sdk"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type LoginDetails struct {
	Email    string
	Password string
}

type MFALoginDetails struct {
	MultiFactorAuthToken    string
	GoogleAuthenticatorCode string
}

type GoogleAuthCode struct {
	GoogleAuthenticatorCode string
}

type SignupDetails struct {
	Email    SignupEmails
	Password string
}

type SignupEmails []struct {
	Type  string
	Value string
}

type Verification struct {
	VerificationToken string
}

type ForgotPasswordEmail struct {
	Email string
}

type ResetPasswordRequest struct {
	ResetToken string
	Password   string
	Email      string
}

type ResetPasswordEmail struct {
	ResetToken                 string
	Password                   string
	WelcomeEmailTemplate       string
	ResetPasswordEmailTemplate string
}

type ResetPasswordOTP struct {
	OTP                        string
	Password                   string
	WelcomeEmailTemplate       string
	ResetPasswordEmailTemplate string
	Email                      string
}

type ChangePasswordDetails struct {
	OldPassword string
	NewPassword string
}

type SetPasswordDetails struct {
	Password string
}

type UpdateAccount struct {
	FirstName string
	LastName  string
	About     string
}

type ResetGoogleAuth struct {
	GoogleAuthenticator bool
}

type Role struct {
	Name string
}

type CreateRoles struct {
	Roles []Role
}

type AssignRoles struct {
	Roles []string
}

type ProfileObject struct {
	FullName      string
	Provider      string
	Email         SignupEmails
	LastLoginDate time.Time
}

func main() {
	loginradius.SetLoginRadiusEnv("", "", "https://api.loginradius.com")
	router := mux.NewRouter()
	router.HandleFunc("/login/email", loginHandler).Methods("POST")
	router.HandleFunc("/mfa/login/email", mfaLoginHandler).Methods("POST")
	router.HandleFunc("/mfa/google/auth", mfaLoginAuthHandler).Methods("PUT")
	router.HandleFunc("/login/passwordless", pwlessHandler).Methods("GET")
	router.HandleFunc("/login/passwordless/auth", verifyLoginHandler).Methods("GET")
	router.HandleFunc("/register", signupHandler).Methods("POST")
	router.HandleFunc("/register/verify/email", verifyEmailHandler).Methods("GET")
	router.HandleFunc("/forgotpassword", forgotPasswordHandler).Methods("POST")
	router.HandleFunc("/login/resetpassword", resetPasswordByEmailHandler).Methods("PUT")
	router.HandleFunc("/resetpassword", resetPasswordHandler).Methods("PUT")
	router.HandleFunc("/profile/changepassword", changePasswordHandler).Methods("PUT")
	router.HandleFunc("/profile/setpassword", setPasswordHandler).Methods("PUT")
	router.HandleFunc("/profile/update", updateAccountHandler).Methods("PUT")
	router.HandleFunc("/customobj", createCustomObjHandler).Methods("POST")
	router.HandleFunc("/customobj", updateCustomObjHandler).Methods("PUT")
	router.HandleFunc("/customobj", deleteCustomObjHandler).Methods("DELETE")
	router.HandleFunc("/customobj", getCustomObjHandler).Methods("GET")
	router.HandleFunc("/mfa/google", mfaResetGoogleHandler).Methods("DELETE")
	router.HandleFunc("/mfa/validate", mfaAccessTokenHandler).Methods("GET")
	router.HandleFunc("/mfa/google/enable", mfaAccessTokenAuthHandler).Methods("PUT")
	router.HandleFunc("/roles", getRolesHandler).Methods("GET")
	router.HandleFunc("/roles/get", getRoleHandler).Methods("GET")
	router.HandleFunc("/roles", createRoleHandler).Methods("POST")
	router.HandleFunc("/roles", deleteRoleHandler).Methods("DELETE")
	router.HandleFunc("/roles", assignRoleHandler).Methods("PUT")
	router.HandleFunc("/profile", profileHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials LoginDetails
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &credentials)
	profile, err := loginradius.PostAuthLoginByEmail("", "", "", "", "", credentials)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func mfaLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials LoginDetails
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &credentials)
	profile, err := loginradius.PostMFAEmailLogin("", "", "", "", credentials)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func mfaLoginAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials MFALoginDetails
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &credentials)
	googleCode := GoogleAuthCode{credentials.GoogleAuthenticatorCode}
	multiFactorAuthToken := r.URL.Query().Get("multi_factor_auth_token")
	profile, err := loginradius.PutMFAValidateGoogleAuthCode(multiFactorAuthToken, "", googleCode)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func pwlessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := r.URL.Query().Get("email")
	verificationURL := r.URL.Query().Get("verification_url")
	fmt.Println(verificationURL)
	profile, err := loginradius.GetPasswordlessLoginByEmail(email, "", verificationURL)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func verifyLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	verificationToken := r.URL.Query().Get("verification_token")
	profile, err := loginradius.GetPasswordlessLoginVerification(verificationToken, "")
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	verificationURL := r.URL.Query().Get("verification_url")
	var credentials SignupDetails
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &credentials)
	profile, err := loginradius.PostAuthUserRegistrationByEmail(verificationURL, "", "", credentials)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func verifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	verificationToken := r.URL.Query().Get("verification_token")
	profile, err := loginradius.GetAuthVerifyEmail(verificationToken, "", "")
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var email ForgotPasswordEmail
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &email)
	resetPassURL := r.URL.Query().Get("reset_password_url")
	profile, err := loginradius.PostAuthForgotPassword(resetPassURL, "", email)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func resetPasswordByEmailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resetPassInfo ResetPasswordRequest
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &resetPassInfo)
	resetPass := ResetPasswordEmail{resetPassInfo.ResetToken, resetPassInfo.Password, "", ""}
	profile, err := loginradius.PutAuthResetPasswordByResetToken(resetPass)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessToken := r.URL.Query().Get("auth")
	var passwords ChangePasswordDetails
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &passwords)
	profile, err := loginradius.PutAuthChangePassword(accessToken, passwords)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessToken := r.URL.Query().Get("auth")
	var passwords ChangePasswordDetails
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &passwords)
	profile, err := loginradius.PutAuthChangePassword(accessToken, passwords)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func setPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uid := r.URL.Query().Get("uid")
	var password SetPasswordDetails
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &password)
	profile, err := loginradius.PutManageAccountSetPassword(uid, password)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func updateAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uid := r.URL.Query().Get("uid")
	var updates UpdateAccount
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &updates)
	profile, err := loginradius.PutManageAccountUpdate(uid, updates)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func createCustomObjHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessToken := r.URL.Query().Get("auth")
	objectName := r.URL.Query().Get("object_name")
	var customObj json.RawMessage
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &customObj)
	profile, err := loginradius.PostCustomObjectCreateByToken(objectName, accessToken, customObj)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func updateCustomObjHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessToken := r.URL.Query().Get("auth")
	objectName := r.URL.Query().Get("object_name")
	objectID := r.URL.Query().Get("object_id")
	var customObj json.RawMessage
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &customObj)
	profile, err := loginradius.PutCustomObjectUpdateByToken(objectName, "replace", accessToken, objectID, customObj)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func deleteCustomObjHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessToken := r.URL.Query().Get("auth")
	objectName := r.URL.Query().Get("object_name")
	objectID := r.URL.Query().Get("object_id")
	profile, err := loginradius.DeleteCustomObjectByObjectRecordIDAndToken(objectName, accessToken, objectID)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func getCustomObjHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessToken := r.URL.Query().Get("auth")
	objectName := r.URL.Query().Get("object_name")
	profile, err := loginradius.GetCustomObjectByToken(objectName, accessToken)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func mfaResetGoogleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	auth := r.URL.Query().Get("auth")
	var resetAuth ResetGoogleAuth
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &resetAuth)
	profile, err := loginradius.DeleteMFAResetGoogleAuthenticatorByToken(auth, resetAuth)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func mfaAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	auth := r.URL.Query().Get("auth")
	profile, err := loginradius.GetMFAValidateAccessToken("", auth)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func mfaAccessTokenAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	auth := r.URL.Query().Get("auth")
	var googleCode GoogleAuthCode
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &googleCode)
	profile, err := loginradius.PutMFAUpdateByToken("", auth, googleCode)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func getRolesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	profile, err := loginradius.GetRolesList()
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func getRoleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uid := r.URL.Query().Get("uid")
	profile, err := loginradius.GetRolesByUID(uid)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func createRoleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var roles CreateRoles
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &roles)
	profile, err := loginradius.PostRolesCreate(roles)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func deleteRoleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	role := r.URL.Query().Get("role")
	profile, err := loginradius.DeleteAccountRole(role)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func assignRoleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uid := r.URL.Query().Get("uid")
	var roles AssignRoles
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &roles)
	profile, err := loginradius.PutRolesAssignToUser(uid, roles)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessToken := r.URL.Query().Get("auth")
	profile, err := loginradius.GetAuthReadProfilesByToken(accessToken)
	if err != nil {
		w.WriteHeader(err.(*loginradius.HTTPError).StatusCode)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(profile)
	w.Write(data)
}
