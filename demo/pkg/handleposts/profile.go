package handleposts

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	lr "github.com/LoginRadius/go-sdk"
	lrauthentication "github.com/LoginRadius/go-sdk/api/authentication"
	"github.com/LoginRadius/go-sdk/lrerror"
)

func Profile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var errors string
	respCode := 200

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	lrclient, err := lr.NewLoginradius(
		&cfg,
		map[string]string{"token": r.URL.Query().Get("auth")},
	)
	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	res, err := lrauthentication.Loginradius(lrauthentication.Loginradius{lrclient}).GetAuthReadProfilesByToken()
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
