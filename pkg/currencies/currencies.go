package currencies

import (
	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/rotisserie/eris"
)

type curr struct {
	Currencies []string `json:"currencies"`
}

// All returns a list of all supported cryptocurrencies.
func All() ([]string, error) {
	c := &curr{}
	err := core.HTTPSend("currencies", nil, &c)
	if err != nil {
		return c.Currencies, eris.Wrap(err, "all")
	}
	return c.Currencies, nil
}
