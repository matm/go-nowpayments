package types

type Estimate struct {
	CurrencyFrom    string  `json:"currency_from"`
	CurrencyTo      string  `json:"currency_to"`
	AmountFrom      float64 `json:"amount_from"`
	EstimatedAmount string  `json:"estimated_amount"`
}
