package config

import (
	"encoding/json"
	"io"

	"github.com/rotisserie/eris"
)

type credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	APIKey   string `json:"apiKey"`
}

var conf credentials

// Load parses a JSON file to get the required credentials to operate NOWPayment's API.
func Load(r io.Reader) error {
	d := json.NewDecoder(r)
	err := d.Decode(&conf)
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
