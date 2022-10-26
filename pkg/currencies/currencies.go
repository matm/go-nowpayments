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
	err := core.HTTPSend("currencies", nil, nil, &c)
	return c.Currencies, eris.Wrap(err, "all")
}
