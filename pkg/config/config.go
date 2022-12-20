package config

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"

	"github.com/rotisserie/eris"
)

type credentials struct {
	APIKey   string `json:"apiKey"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Server   string `json:"server"`
}

var conf credentials

func configErr(err error) error {
	return eris.Wrap(err, "config")
}

// Load parses a JSON file to get the required credentials to operate NOWPayment's API.
func Load(r io.Reader) error {
	if r == nil {
		return configErr(errors.New("nil reader"))
	}
	conf = credentials{}
	d := json.NewDecoder(r)
	err := d.Decode(&conf)
	if err != nil {
		return configErr(err)
	}
	// Sanity checks.
	if conf.APIKey == "" {
		return configErr(errors.New("API key is missing"))
	}
	if conf.Login == "" {
		return configErr(errors.New("login info missing"))
	}
	if conf.Password == "" {
		return configErr(errors.New("password info missing"))
	}
	if conf.Server == "" {
		return configErr(errors.New("server URL missing"))
	} else {
		_, err := url.Parse(conf.Server)
		if err != nil {
			return configErr(errors.New("server URL parsing"))
		}
	}
	return nil
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

// Server returns URL to connect to the API service.
func Server() string {
	return conf.Server
}
