package currencies

import (
	"github.com/matm/go-nowpayments/core"
)

// All returns a list of all supported cryptocurrencies.
func All() ([]string, error) {
	type curr struct {
		All []string `json:"currencies"`
	}
	c := &curr{}
	par := &core.SendParams{
		RouteName: "currencies",
		Into:      &c,
	}
	return c.All, core.HTTPSend(par)
}

// Selected returns information about the cryptocurrencies available for payments.
// Shows the coins set as available for payments in the "coins settings" tab
// on personal account page.
func Selected() ([]string, error) {
	type selCur struct {
		All []string `json:"selectedCurrencies"`
	}
	c := &selCur{}
	par := &core.SendParams{
		RouteName: "selected-currencies",
		Into:      &c,
	}
	return c.All, core.HTTPSend(par)
}
