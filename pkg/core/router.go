package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/matm/go-nowpayments/pkg/config"
	"github.com/rotisserie/eris"
)

// BaseURL is the URL to NOWPayment's service.
type BaseURL string

const (
	// ProductionBaseURL is the URL to the production service.
	ProductionBaseURL BaseURL = "https://api.nowpayments.io/v1"
	// SandBoxBaseURL is the URL to the sandbox service.
	SandBoxBaseURL = "https://api-sandbox.nowpayments.io/v1"
)

// SendParams are parameters needed to build and send an HTTP request to the service.
type SendParams struct {
	Body      io.Reader
	Into      interface{}
	Path      string
	RouteName string
	Values    url.Values
	// JWT token obtained after authentication.
	Token string
}

type routeAttr struct {
	method string
	path   string
}

var routes map[string]routeAttr = map[string]routeAttr{
	"auth":                {http.MethodPost, "/auth"},
	"currencies":          {http.MethodGet, "/currencies"},
	"estimate":            {http.MethodGet, "/estimate"},
	"invoice-create":      {http.MethodPost, "/invoice"},
	"invoice-payment":     {http.MethodPost, "/invoice-payment"},
	"last-estimate":       {http.MethodPost, "/payment"},
	"min-amount":          {http.MethodGet, "/min-amount"},
	"payment-create":      {http.MethodPost, "/payment"},
	"payment-status":      {http.MethodGet, "/payment"},
	"payments-list":       {http.MethodGet, "/payment/"},
	"selected-currencies": {http.MethodGet, "/merchant/coins"},
	"status":              {http.MethodGet, "/status"},
}

var (
	defaultURL BaseURL = SandBoxBaseURL
)

var debug = false

// WithDebug prints out debugging info about HTTP traffic.
func WithDebug(d bool) {
	debug = d
}

// UseBaseURL sets the base URL to use to connect to NOWPayment's API.
func UseBaseURL(b BaseURL) {
	defaultURL = b
}

// HTTPSend sends to endpoint with an optional request body and get the HTTP
// response result in into.
func HTTPSend(p *SendParams) error {
	if p == nil {
		return eris.New("nil params")
	}
	method, path := routes[p.RouteName].method, routes[p.RouteName].path
	if path == "" {
		return eris.New(fmt.Sprintf("bad route name: empty path for endpoint %q", p.RouteName))
	}
	u := string(defaultURL) + path
	if p.Path != "" {
		u += "/" + p.Path
	}
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
		fmt.Println(">>> DEBUG REQUEST")
		fmt.Printf("X-API-KEY: %s\n", req.Header.Get("X-API-KEY"))
		fmt.Printf("Authorization: %s\n", req.Header.Get("Authorization"))
		fmt.Println(req.Method, req.URL.String())
		fmt.Println("<<< END DEBUG REQUEST")
	}
	res, err := client.Do(req)
	if err != nil {
		return eris.Wrap(err, p.RouteName)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		if debug {
			fmt.Printf(">>> DEBUG HTTP error %d: %s\n", res.StatusCode, res.Status)
		}
		type errResp struct {
			StatusCode int    `json:"statusCode"`
			Code       string `json:"code"`
			Message    string `json:"message"`
		}
		z := &errResp{}
		d := json.NewDecoder(res.Body)
		err = d.Decode(&z)
		if err != nil {
			return eris.Wrapf(err, "%s: JSON decode error", p.RouteName)
		}
		return eris.New(fmt.Sprintf("code %d (%s): %s", z.StatusCode, z.Code, z.Message))
	}
	if debug {
		fmt.Println(">>> DEBUG RAW RESPONSE BODY")
		all, err := io.ReadAll(res.Body)
		if err != nil {
			return eris.Wrap(err, "debug response")
		}
		fmt.Println(string(all))
		fmt.Println("<<< END DEBUG RAW RESPONSE BODY")
		return eris.Wrap(json.Unmarshal(all, &p.Into), p.RouteName)
	}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&p.Into)
	return eris.Wrap(err, p.RouteName)
}
