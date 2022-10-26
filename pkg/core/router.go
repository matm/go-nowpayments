package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/matm/go-nowpayments/pkg/types"
	"github.com/rotisserie/eris"
)

type routeAttr struct {
	method string
	path   string
}

var routes map[string]routeAttr = map[string]routeAttr{
	"status":     {http.MethodGet, "/status"},
	"currencies": {http.MethodGet, "/currencies"},
	"estimate":   {http.MethodGet, "/estimate"},
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

// UseAPIKey sets the API key to use for all requests.
func UseAPIKey(key string) {
	apiKey = key
}

// APIKey returns the current API key set or the default URL to sandbox.
func APIKey() string {
	return apiKey
}

// HTTPSend sends to endpoint with an optional request body and get the HTTP
// response result in into.
func HTTPSend(endpoint string, body io.Reader, values url.Values, into interface{}) error {
	client := &http.Client{}
	method, path := routes[endpoint].method, routes[endpoint].path
	u := string(defaultURL) + path
	if values != nil {
		u += "?" + values.Encode()
	}
	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return eris.Wrap(err, endpoint)
	}
	req.Header.Add("X-API-KEY", apiKey)
	res, err := client.Do(req)
	if err != nil {
		return eris.Wrap(err, endpoint)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		type errResp struct {
			StatusCode int    `json:"statusCode"`
			Code       string `json:"code"`
			Message    string `json:"message"`
		}
		z := &errResp{}
		d := json.NewDecoder(res.Body)
		err = d.Decode(&z)
		if err != nil {
			return eris.Wrap(err, endpoint)
		}
		return eris.New(fmt.Sprintf("code %d (%s): %s", z.StatusCode, z.Code, z.Message))
	}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&into)
	return eris.Wrap(err, endpoint)
}
