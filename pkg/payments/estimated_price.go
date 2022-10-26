package payments

import (
	"fmt"
	"net/url"

	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/matm/go-nowpayments/pkg/types"
	"github.com/rotisserie/eris"
)

// EstimatedPrice calculates the approximate price in cryptocurrency for a given value in Fiat currency.
// Need to provide the initial cost in the Fiat currency (amount, currency_from) and the necessary cryptocurrency (currency_to).
// Currently following fiat currencies are available: usd, eur, nzd, brl.
func EstimatedPrice(amount float64, currencyFrom, currencyTo string) (*types.Estimate, error) {
	u := url.Values{}
	u.Set("amount", fmt.Sprintf("%f", amount))
	u.Set("currency_from", currencyFrom)
	u.Set("currency_to", currencyTo)
	e := &types.Estimate{}
	err := core.HTTPSend("estimate", nil, u, &e)
	return e, eris.Wrap(err, "estimate")
}
