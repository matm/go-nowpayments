package payments

import (
	"fmt"
	"net/url"

	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/rotisserie/eris"
)

// Estimate holds the estimated amount of crypto for a given Fiat value.
type Estimate struct {
	CurrencyFrom    string  `json:"currency_from"`
	CurrencyTo      string  `json:"currency_to"`
	AmountFrom      float64 `json:"amount_from"`
	EstimatedAmount string  `json:"estimated_amount"`
}

// EstimatedPrice calculates the approximate price in cryptocurrency for a given value in Fiat currency.
// Need to provide the initial cost in the Fiat currency (amount, currency_from) and the necessary cryptocurrency (currency_to).
// Currently following fiat currencies are available: usd, eur, nzd, brl.
func EstimatedPrice(amount float64, currencyFrom, currencyTo string) (*Estimate, error) {
	if amount == 0 {
		return nil, eris.New("use a price greater than zero")
	}
	u := url.Values{}
	u.Set("amount", fmt.Sprintf("%f", amount))
	u.Set("currency_from", currencyFrom)
	u.Set("currency_to", currencyTo)
	e := &Estimate{}
	par := &core.SendParams{
		RouteName: "estimate",
		Into:      &e,
		Values:    u,
	}
	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}
	return e, nil
}
