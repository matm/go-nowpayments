package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matm/go-nowpayments/pkg/config"
	"github.com/matm/go-nowpayments/pkg/types"
	"github.com/rotisserie/eris"
)

type routeAttr struct {
	method string
	path   string
}

var routes map[string]routeAttr = map[string]routeAttr{
	"status":         {http.MethodGet, "/status"},
	"currencies":     {http.MethodGet, "/currencies"},
	"estimate":       {http.MethodGet, "/estimate"},
	"min-amount":     {http.MethodGet, "/min-amount"},
	"payment-status": {http.MethodGet, "/payment"},
	"auth":           {http.MethodPost, "/auth"},
}

var (
	defaultURL types.BaseURL = types.SandBoxBaseURL
)

var debug = false

// WithDebug prints out debugging info about HTTP traffic.
func WithDebug(d bool) {
	debug = d
}

// UseBaseURL sets the base URL to use to connect to NOWPayment's API.
func UseBaseURL(b types.BaseURL) {
	defaultURL = b
}

// BaseURL returns the base URL used to connect to NOWPayment's API.
func BaseURL() string {
	return string(defaultURL)
}

// HTTPSend sends to endpoint with an optional request body and get the HTTP
// response result in into.
func HTTPSend(p *types.SendParams) error {
	if p == nil {
		return eris.New("nil params")
	}
	client := &http.Client{}
	method, path := routes[p.RouteName].method, routes[p.RouteName].path
	if path == "" {
		return eris.New(fmt.Sprintf("empty path for endpoint %q", p.RouteName))
	}
	u := string(defaultURL) + path
	if p.Values != nil {
		u += "?" + p.Values.Encode()
	}
	req, err := http.NewRequest(method, u, p.Body)
	if err != nil {
		return eris.Wrap(err, p.RouteName)
	}
	// Extra headers.
	req.Header.Add("X-API-KEY", config.APIKey())
	if p.Body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	if p.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.Token))
	}

	if debug {
		fmt.Println(">>> DEBUG")
		fmt.Println(req.Method, req.URL.String())
		fmt.Println("<<< DEBUG")
	}
	res, err := client.Do(req)
	if err != nil {
		return eris.Wrap(err, p.RouteName)
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
			return eris.Wrap(err, p.RouteName)
		}
		return eris.New(fmt.Sprintf("code %d (%s): %s", z.StatusCode, z.Code, z.Message))
	}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&p.Into)
	return eris.Wrap(err, p.RouteName)
}
