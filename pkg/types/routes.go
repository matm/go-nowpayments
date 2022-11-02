package types

import (
	"io"
	"net/url"
)

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
