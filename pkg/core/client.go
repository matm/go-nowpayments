package core

import "net/http"

// HTTPClient defines methods of an HTTP client.
type HTTPClient interface {
	// Do executes the HTTP request to the API server.
	Do(*http.Request) (*http.Response, error)
}

var client HTTPClient

// UseClient specifies which API server to use.
func UseClient(s HTTPClient) {
	client = s
}

type httpclient struct{}

func (*httpclient) Do(r *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(r)
}

// NewHTTPClient uses the default http.Client.
func NewHTTPClient() HTTPClient {
	return &httpclient{}
}
