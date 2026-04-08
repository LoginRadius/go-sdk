package handlegets

import (
	"log"
	"net/http"
	"os"

	lr "github.com/LoginRadius/go-sdk"
	"github.com/LoginRadius/go-sdk/api/role"
	"github.com/LoginRadius/go-sdk/lrerror"
	"github.com/julienschmidt/httprouter"
)

func UserRoles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	uid := r.URL.Query().Get("uid")
	res, err := role.Loginradius(role.Loginradius{lrclient}).GetRolesByUID(uid)
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

func Roles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	res, err := role.Loginradius(role.Loginradius{lrclient}).GetRolesList()
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
