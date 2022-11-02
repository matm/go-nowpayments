package types

type Estimate struct {
	CurrencyFrom    string  `json:"currency_from"`
	CurrencyTo      string  `json:"currency_to"`
	AmountFrom      float64 `json:"amount_from"`
	EstimatedAmount string  `json:"estimated_amount"`
}

type MinimumAmount struct {
	CurrencyFrom string  `json:"currency_from"`
	CurrencyTo   string  `json:"currency_to"`
	Amount       float64 `json:"min_amount"`
}

type PaymentStatus struct {
	ActuallyPaid float64 `json:"actually_paid"`
	// CreatedAt looks like 2019-04-18T13:39:27.982Z.
	CreatedAt       string  `json:"created_at"`
	OutcomeAmount   float64 `json:"outcome_amount"`
	OutcomeCurrency string  `json:"outcome_currency"`
	PayAddress      string  `json:"pay_address"`
	PayAmount       float64 `json:"pay_amount"`
	PayCurrency     float64 `json:"pay_currency"`
	PriceAmount     float64 `json:"price_amount"`
	PriceCurrency   string  `json:"price_currency"`
	PurchaseID      string  `json:"purchase_id"`
	Status          string  `json:"payment_status"`
	UpdatedAt       string  `json:"updated_at"`
}
