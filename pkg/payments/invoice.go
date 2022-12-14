package payments

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/rotisserie/eris"
)

// InvoiceArgs are the arguments used to make an invoice.
type InvoiceArgs struct {
	PaymentAmount

	CancelURL  string `json:"cancel_url,omitempty"`
	SuccessURL string `json:"success_url,omitempty"`
}

// Invoice describes an invoice. InvoiceURL is the URL to follow to
// make the payment.
type Invoice struct {
	InvoiceArgs

	// FIXME: inconsistency on their side: should be a float64, like
	// the field used for a payment.
	PriceAmount string `json:"price_amount"`
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at,omitempty"`
	InvoiceURL  string `json:"invoice_url,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// NewInvoice creates an invoice.
func NewInvoice(ia *InvoiceArgs) (*Invoice, error) {
	if ia == nil {
		return nil, errors.New("nil invoice args")
	}
	d, err := json.Marshal(ia)
	if err != nil {
		return nil, eris.Wrap(err, "invoice args")
	}
	p := &Invoice{}
	par := &core.SendParams{
		RouteName: "invoice-create",
		Into:      &p,
		Body:      strings.NewReader(string(d)),
	}
	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}
	return p, nil
}
