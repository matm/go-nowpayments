package types

type BaseURL string

const (
	// ProductionBaseURL is the URL to the production service.
	ProductionBaseURL BaseURL = "https://api.nowpayments.io/v1"
	// SandBoxBaseURL is the URL to the sandbox service.
	SandBoxBaseURL = "https://api-sandbox.nowpayments.io/v1"
)
