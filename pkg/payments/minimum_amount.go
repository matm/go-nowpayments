package payments

import (
	"net/url"

	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/matm/go-nowpayments/pkg/types"
	"github.com/rotisserie/eris"
)

// MinimumAmount returns the minimum payment amount for a specific pair.
func MinimumAmount(currencyFrom, currencyTo string) (*types.MinimumAmount, error) {
	u := url.Values{}
	u.Set("currency_from", currencyFrom)
	u.Set("currency_to", currencyTo)
	e := &types.MinimumAmount{}
	err := core.HTTPSend("min-amount", nil, u, &e)
	return e, eris.Wrap(err, "minamount")
}
