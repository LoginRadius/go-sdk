package handleposts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	lr "github.com/LoginRadius/go-sdk"
	"github.com/LoginRadius/go-sdk/api/customobject"
	"github.com/LoginRadius/go-sdk/lrerror"
)

func CustomObject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var errors string
	respCode := 200

	cfg := lr.Config{
		ApiKey:    os.Getenv("APIKEY"),
		ApiSecret: os.Getenv("APISECRET"),
	}

	token := r.Header.Get("Authorization")[7:]
	lrclient, err := lr.NewLoginradius(
		&cfg,
		map[string]string{"token": token},
	)

	if err != nil {
		errors = errors + err.(lrerror.Error).OrigErr().Error()
		respCode = 500
	}

	var customObj json.RawMessage
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &customObj)

	res, err := customobject.Loginradius(customobject.Loginradius{lrclient}).PostCustomObjectCreateByToken(
		map[string]string{"objectname": r.URL.Query().Get("object_name")},
		customObj,
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
