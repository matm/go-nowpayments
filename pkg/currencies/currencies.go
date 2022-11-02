package currencies

import (
	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/matm/go-nowpayments/pkg/types"
	"github.com/rotisserie/eris"
)

type curr struct {
	Currencies []string `json:"currencies"`
}

// All returns a list of all supported cryptocurrencies.
func All() ([]string, error) {
	c := &curr{}
	par := &types.SendParams{
		RouteName: "currencies",
		Into:      &c,
	}
	err := core.HTTPSend(par)
	return c.Currencies, eris.Wrap(err, "all")
}
