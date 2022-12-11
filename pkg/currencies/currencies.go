package currencies

import (
	"github.com/matm/go-nowpayments/pkg/core"
)

type curr struct {
	Currencies []string `json:"currencies"`
}

// All returns a list of all supported cryptocurrencies.
func All() ([]string, error) {
	c := &curr{}
	par := &core.SendParams{
		RouteName: "currencies",
		Into:      &c,
	}
	return c.Currencies, core.HTTPSend(par)
}
