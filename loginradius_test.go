package loginradius

import (
	"errors"
	"reflect"
	"testing"

	"github.com/LoginRadius/go-sdk/lrerror"
)

const errMsg = "Must initialize Loginradius client with ApiKey and ApiSecret"

var config = Config{"abc123", "abc123"}
var lr, _ = NewLoginradius(&config)

type testCase struct {
	config      Config
	errResp     error
	successResp *Loginradius
}

var testCases = []testCase{
	testCase{config, nil, lr},
	testCase{Config{}, lrerror.New("IntializationError", errMsg, errors.New(errMsg)), nil},
}

func TestNewLoginradius(t *testing.T) {
	for _, c := range testCases {
		loginradius, err := NewLoginradius(&c.config)
		if !reflect.DeepEqual(err, c.errResp) || !reflect.DeepEqual(loginradius, c.successResp) {
			t.Errorf("Expected NewLoginRadius to return %v, but instead got: %v", c.errResp, err)
		}
	}
}
