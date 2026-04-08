package handleposts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	lr "github.com/LoginRadius/go-sdk"
	lrauthentication "github.com/LoginRadius/go-sdk/api/authentication"
	"github.com/LoginRadius/go-sdk/internal/sott"
	"github.com/LoginRadius/go-sdk/lrerror"
	"github.com/julienschmidt/httprouter"
)

type Email struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

type User struct {
	Email                []Email `json:"Email"`
	Password             string  `json:"Password"`
	PasswordConfirmation string  `json:"Password_confirmation"`
}

func Signup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var errors string
	respCode := 200

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, err := lr.NewLoginradius(&cfg)
	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	var user User
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &user)
	verificationURL := r.URL.Query().Get("verification_url")
	startTime := ""
	endTime := ""
	timeDifference := ""
	sott := sott.Generate(lrclient.Context.ApiKey, lrclient.Context.ApiSecret, timeDifference, startTime, endTime)

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).PostAuthUserRegistrationByEmail(
		sott, user,
		map[string]string{"verificationurl": verificationURL},
	)
	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	if errors != "" {
		log.Print(errors)
		w.Write([]byte(errors))
		return
	}
	w.Write([]byte(res.Body))
}
