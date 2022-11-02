package config

import (
	"encoding/json"
	"os"

	"github.com/rotisserie/eris"
)

type credentials struct {
	Login, Password string
	APIKey          string
}

var conf credentials

// Load parses a JSON file to get the required credentials to operate NOWPayment's API.
func Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return eris.Wrap(err, "load config")
	}
	d := json.NewDecoder(f)
	err = d.Decode(&conf)
	return eris.Wrap(err, "decode config")
}

// Login returns the email address to use with the API.
func Login() string {
	return conf.Login
}

// Password returns the related password to use.
func Password() string {
	return conf.Password
}

// APIKey is the API key to use.
func APIKey() string {
	return conf.APIKey
}
