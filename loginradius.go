// Package loginradius contains struct and initialization methods for loginradius api client
package loginradius

import (
	"errors"

	"github.com/LoginRadius/go-sdk/lrerror"
)

// domain is the default domain for API calls to Loginradius
const domain = "https://api.loginradius.com"

// Loginradius struct holds context for intializing the Loginradius client and the domain for API calls
// Domain can be changed after intialization
type Loginradius struct {
	Context *Context
	Domain  string
}

// Config struct contains Loginradius credentials and is used when initalizing the Loginradius API client struct
type Config struct {
	ApiKey    string
	ApiSecret string
}

// Context struct is a field in the Loginradius struct
type Context struct {
	ApiKey    string
	ApiSecret string
	Token     string
}

// NewLoginradius initializes a new Loginradius struct with a Config struct
// Config struct must contain the ApiKey and ApiSecret of your Loginradius site
// Example:
// 			cfg := lr.Config{
// 				ApiKey:    os.Getenv("APIKEY"),
// 				ApiSecret: os.Getenv("APISECRET"),
// 			}
// 			lrclient, _ := lr.NewLoginradius(&cfg)
// It also takes optional variadic arguments
// Some APIs require for an access token to be passed with Authorization Bearer header
// Initialize Loginradius struct with a token passed in the variadic argument like so:
// 			lrclient, _ := lr.NewLoginradius(&cfg, map[string]string{"token": "9c3208ae-2848-4ac5-baef-41dd4103e263"})
func NewLoginradius(cfg *Config, optionalArgs ...map[string]string) (*Loginradius, error) {

	if cfg.ApiKey == "" || cfg.ApiSecret == "" {
		errMsg := "Must initialize Loginradius client with ApiKey and ApiSecret"
		err := lrerror.New("IntializationError", errMsg, errors.New(errMsg))
		return nil, err
	}

	ctx := Context{
		ApiKey:    cfg.ApiKey,
		ApiSecret: cfg.ApiSecret,
	}

	// If an access token is passed on initiation, set it in Context
	for _, arg := range optionalArgs {
		if arg["token"] != "" {
			ctx.Token = arg["token"]
		} else {
			ctx.Token = ""
		}
	}

	return &Loginradius{
		Context: &ctx,
		Domain:  domain,
	}, nil
}
