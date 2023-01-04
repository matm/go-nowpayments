package payments

import (
	"net/url"

	"github.com/matm/go-nowpayments/core"
)

// CurrencyAmount has info about minimum payment amount for a specific pair.
type CurrencyAmount struct {
	CurrencyFrom string  `json:"currency_from"`
	CurrencyTo   string  `json:"currency_to"`
	Amount       float64 `json:"min_amount"`
}

// MinimumAmount returns the minimum payment amount for a specific pair.
func MinimumAmount(currencyFrom, currencyTo string) (*CurrencyAmount, error) {
	u := url.Values{}
	u.Set("currency_from", currencyFrom)
	u.Set("currency_to", currencyTo)
	e := &CurrencyAmount{}
	par := &core.SendParams{
		RouteName: "min-amount",
		Into:      &e,
		Values:    u,
	}
	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}
	return e, nil
}
