package handleposts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	lr "github.com/LoginRadius/go-sdk"
	"github.com/LoginRadius/go-sdk/api/mfa"
	"github.com/LoginRadius/go-sdk/lrbody"
	"github.com/LoginRadius/go-sdk/lrerror"
)

func MfaLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	var credentials lrbody.EmailLogin
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &credentials)

	res, err := mfa.Loginradius(mfa.Loginradius{lrclient}).PostMFAEmailLogin(
		credentials,
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
