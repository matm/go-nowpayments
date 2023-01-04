package payments

import (
	"fmt"
	"net/url"

	"github.com/matm/go-nowpayments/config"
	"github.com/matm/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

// PaymentHack has all the fields from Payment except the ID field which has
// a different type due to an API inconsistency on their side.
type PaymentHack struct {
	Payment
	// Using ID as an int instead a string as a quick work around.
	ID int `json:"payment_id"`
}

// ListOption are options applying to the list of transactions.
type ListOption struct {
	DateFrom string
	DateTo   string
	Limit    int
	OrderBy  string
	Page     int
	SortBy   string
}

// List returns a list of all transactions for a given API key, depending
// on the supplied options (which can be nil).
func List(o *ListOption) ([]*PaymentHack, error) {
	u := url.Values{}
	if o != nil {
		if o.Limit != 0 {
			u.Set("limit", fmt.Sprintf("%d", o.Limit))
		}
		if o.DateFrom != "" {
			u.Set("dateFrom", o.DateFrom)
		}
		if o.DateTo != "" {
			u.Set("dateTo", o.DateTo)
		}
		u.Set("page", fmt.Sprintf("%d", o.Page))
		if o.SortBy != "" {
			u.Set("sortBy", o.SortBy)
		}
		if o.OrderBy != "" {
			u.Set("orderBy", o.OrderBy)
		}
	}
	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "list")
	}
	type plist struct {
		Data []*PaymentHack `json:"data"`
	}
	pl := &plist{Data: make([]*PaymentHack, 0)}
	par := &core.SendParams{
		RouteName: "payments-list",
		Into:      pl,
		Values:    u,
		Token:     tok,
	}
	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}
	return pl.Data, nil
}
