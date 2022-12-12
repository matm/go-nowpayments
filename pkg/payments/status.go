package payments

import (
	"github.com/matm/go-nowpayments/pkg/config"
	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/rotisserie/eris"
)

// PaymentStatus is the actual information about a payment.
type PaymentStatus struct {
	ActuallyPaid float64 `json:"actually_paid"`
	// CreatedAt looks like 2019-04-18T13:39:27.982Z.
	CreatedAt       string  `json:"created_at"`
	OutcomeAmount   float64 `json:"outcome_amount"`
	OutcomeCurrency string  `json:"outcome_currency"`
	PayAddress      string  `json:"pay_address"`
	PayAmount       float64 `json:"pay_amount"`
	PayCurrency     string  `json:"pay_currency"`
	PriceAmount     float64 `json:"price_amount"`
	PriceCurrency   string  `json:"price_currency"`
	PurchaseID      int     `json:"purchase_id"`
	Status          string  `json:"payment_status"`
	UpdatedAt       string  `json:"updated_at"`
}

// Status gets the actual information about the payment. You need to provide the ID of the payment in the request.
// Note that unlike what the official doc says, a Bearer JWTtoken is required for this endpoint
// to work.
func Status(paymentID string) (*PaymentStatus, error) {
	// payment status: code 401 (AUTH_REQUIRED): Authorization header is empty (Bearer JWTtoken is required)
	if paymentID == "" {
		return nil, eris.New("empty payment ID")
	}
	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "status")
	}
	st := &PaymentStatus{}
	par := &core.SendParams{
		RouteName: "payment-status",
		Path:      paymentID,
		Into:      &st,
		Token:     tok,
	}
	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}
	return st, nil
}
