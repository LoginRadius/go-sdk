package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/LoginRadius/go-sdk/demo/pkg/handledeletes"
	"github.com/LoginRadius/go-sdk/demo/pkg/handlegets"
	"github.com/LoginRadius/go-sdk/demo/pkg/handleposts"
	"github.com/LoginRadius/go-sdk/demo/pkg/handleputs"
)

func main() {
	cwd, _ := os.Getwd()

	err := godotenv.Load(
		filepath.Join(cwd, "../../config/secret.env"),
		filepath.Join(cwd, "../../config/public.env"),
	)

	if err != nil {
		log.Fatal("Error loading env files, please configure your secret.env and public.env.")
	}

	router := httprouter.New()

	router.POST("/index", handleposts.Index)
	router.GET("/api/register/verify/email", handleposts.Verify)
	router.POST("/api/register", handleposts.Signup)
	router.POST("/api/login/email", handleposts.Login)
	router.POST("/api/profile", handleposts.Profile)
	router.GET("/api/roles/get", handlegets.UserRoles)
	router.GET("/api/roles", handlegets.Roles)
	router.POST("/api/forgotpassword", handleposts.ForgotPassword)
	router.PUT("/api/login/resetpassword", handleputs.ResetPassword)
	router.POST("/api/mfa/login/email", handleposts.MfaLogin)
	router.PUT("/api/mfa/google/auth", handleputs.MfaGoogleAuth)
	router.PUT("/api/profile/changepassword", handleputs.ChangePassword)
	router.PUT("/api/profile/setpassword", handleputs.SetPassword)
	router.PUT("/api/profile/update", handleputs.UpdateProfile)
	router.POST("/api/customobj", handleposts.CustomObject)
	router.GET("/api/customobj", handlegets.CustomObject)
	router.PUT("/api/customobj", handleputs.CustomObject)
	router.DELETE("/api/customobj", handledeletes.CustomObject)
	router.GET("/api/mfa/validate", handlegets.Mfa)
	router.PUT("/api/mfa/google/enable", handleputs.MfaGoogleEnable)
	router.DELETE("/api/mfa/google", handledeletes.MfaGoogleReset)
	router.POST("/api/roles", handleposts.Role)
	router.DELETE("/api/roles", handledeletes.Role)
	router.PUT("/api/roles", handleputs.Role)
	router.GET("/api/login/passwordless", handlegets.Passwordless)
	router.GET("/api/login/passwordless/auth", handlegets.PasswordlessAuth)

	// if not found look for a static file
	static := httprouter.New()
	static.ServeFiles("/*filepath", http.Dir(filepath.Join(cwd, "../../ui/assets")))
	router.NotFound = static

	http.ListenAndServe(":3000", router)
}
