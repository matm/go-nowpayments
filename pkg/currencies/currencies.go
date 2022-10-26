package currencies

import (
	"encoding/json"
	"net/http"

	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/rotisserie/eris"
)

type curr struct {
	Currencies []string `json:"currencies"`
}

// All returns a list of all supported cryptocurrencies.
func All() ([]string, error) {
	client := &http.Client{}
	curs := make([]string, 0)

	method, path := core.Route("currencies")
	req, err := http.NewRequest(method, core.BaseURL()+path, nil)
	if err != nil {
		return curs, eris.Wrap(err, "status")
	}
	req.Header.Add("x-api-key", core.APIKey())
	res, err := client.Do(req)
	if err != nil {
		return curs, eris.Wrap(err, "status")
	}
	defer res.Body.Close()

	c := &curr{}
	d := json.NewDecoder(res.Body)
	err = d.Decode(&c)
	if err != nil {
		return curs, eris.Wrap(err, "status")
	}
	return c.Currencies, nil
}
