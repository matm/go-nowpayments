package payments

import (
	"github.com/matm/go-nowpayments/pkg/core"
)

type PaymentAmount struct {
	PriceAmount      float64 `json:"price_amount"`
	PriceCurrency    string  `json:"price_currency"`
	PayAmount        float64 `json:"pay_amount"`
	PayCurrency      string  `json:"pay_currency"`
	CallbackURL      string  `json:"ipn_callback_url"`
	OrderID          string  `json:"order_id"`
	OrderDescription string  `json:"order_derscription"`
}

type PaymentArgs struct {
	PaymentAmount

	FeePaidByUser  bool   `json:"is_fee_paid_by_user"`
	FixedRate      bool   `json:"fixed_rate"`
	PayoutAddress  string `json:"payout_address"`
	PayoutCurrency string `json:"payout_currency"`
	PayoutExtraID  string `json:"payout_extra_id"`
	PurchaseID     string `json:"purchase_id"`
}

// Payment holds payment related information.
// FIXME: the API doc misses information about returned fields.
// https://documenter.getpostman.com/view/7907941/S1a32n38?version=latest#5e37f3ad-0fa1-4292-af51-5c7f95730486
type Payment struct {
	PaymentAmount

	ID                     string  `json:"payment_id"`
	AmountReceived         float64 `json:"amount_received"`
	BurningPercent         int     `json:"burning_percent"`
	CreatedAt              string  `json:"created_at"`
	ExpirationEstimateDate string  `json:"expiration_estimate_date"`
	Network                string  `json:"network"`
	NetworkPrecision       int     `json:"network_precision"`
	PayAddress             string  `json:"pay_address"`
	PayinExtraID           string  `json:"payin_extra_id"`
	PurchaseID             string  `json:"purchase_id"`
	SmartContract          string  `json:"smart_contract"`
	Status                 string  `json:"payment_status"`
	TimeLimit              string  `json:"time_limit"`
	UpdatedAt              string  `json:"updated_at"`
}

// Create creates a payment.
func Create(a *PaymentArgs) (*Payment, error) {
	p := &Payment{}
	par := &core.SendParams{
		RouteName: "payment-create",
		Into:      &p,
	}
	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}
	return p, nil
}
