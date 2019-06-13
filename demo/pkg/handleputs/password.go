package handleputs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	lr "github.com/LoginRadius/go-sdk"
	account "github.com/LoginRadius/go-sdk/api/account"
	lrauthentication "github.com/LoginRadius/go-sdk/api/authentication"
	"github.com/LoginRadius/go-sdk/lrerror"
)

func ResetPassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	data := struct {
		ResetToken string
		Password   string
	}{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &data)

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).
		PutAuthResetPasswordByResetToken(data)

	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	if errors != "" {
		log.Printf(errors)
		w.Write([]byte(errors))
		return
	}
	w.Write([]byte(res.Body))
}

func ChangePassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var errors string
	respCode := 200

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	data := struct {
		OldPassword string
		NewPassword string
		Token       string
	}{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &data)

	lrclient, err := lr.NewLoginradius(
		&cfg,
		map[string]string{"token": data.Token},
	)

	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).
		PutAuthChangePassword(data)

	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	if errors != "" {
		log.Printf(errors)
		w.Write([]byte(errors))
		return
	}
	w.Write([]byte(res.Body))
}

func SetPassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	data := struct {
		Password string
	}{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &data)

	res, err := account.Loginradius(account.Loginradius{lrclient}).
		PutManageAccountSetPassword(
			r.URL.Query().Get("uid"),
			data,
		)

	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	if errors != "" {
		log.Printf(errors)
		w.Write([]byte(errors))
		return
	}
	w.Write([]byte(res.Body))
}
