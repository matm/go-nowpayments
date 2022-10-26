package core

import (
	"net/http"

	"github.com/matm/go-nowpayments/pkg/types"
)

type routeAttr struct {
	method string
	path   string
}

var routes map[string]routeAttr = map[string]routeAttr{
	"status":     {http.MethodGet, "/status"},
	"currencies": {http.MethodGet, "/currencies"},
}

func Route(name string) (string, string) {
	return routes[name].method, routes[name].path
}

var (
	defaultURL types.BaseURL = types.SandBoxBaseURL
	apiKey     string
)

// UseBaseURL sets the base URL to use to connect to NOWPayment's API.
func UseBaseURL(b types.BaseURL) {
	defaultURL = b
}

// BaseURL returns the base URL used to connect to NOWPayment's API.
func BaseURL() string {
	return string(defaultURL)
}

func UseAPIKey(key string) {
	apiKey = key
}

func APIKey() string {
	return apiKey
}
