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

	FeePaidByUser  bool   `json:"is_fee_paid_by_user,omitempty"`
	FixedRate      bool   `json:"fixed_rate,omitempty"`
	PayoutAddress  string `json:"payout_address,omitempty"`
	PayoutCurrency string `json:"payout_currency,omitempty"`
	PayoutExtraID  string `json:"payout_extra_id,omitempty"`
	PurchaseID     string `json:"purchase_id,omitempty"`
}

// Invoice describes an invoice. InvoiceURL is the URL to follow to
// make the payment.
type Invoice struct {
	PaymentAmount

	ID         string `json:"id"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	SuccessURL string `json:"success_url,omitempty"`
	CancelURL  string `json:"cancel_url,omitempty"`
	InvoiceURL string `json:"invoice_url,omitempty"`
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
