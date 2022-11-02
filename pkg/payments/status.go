package payments

import (
	"github.com/matm/go-nowpayments/pkg/config"
	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/matm/go-nowpayments/pkg/types"
	"github.com/rotisserie/eris"
)

// Status gets the actual information about the payment. You need to provide the ID of the payment in the request.
// Note that unlike what the official doc says, a Bearer JWTtoken is required for this endpoint
// to work.
func Status(paymentID string) (*types.PaymentStatus, error) {
	// payment status: code 401 (AUTH_REQUIRED): Authorization header is empty (Bearer JWTtoken is required)
	if paymentID == "" {
		return nil, eris.New("empty payment ID")
	}
	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "status")
	}

	st := &types.PaymentStatus{}
	par := &types.SendParams{
		RouteName: "payment-status",
		Path:      paymentID,
		Into:      &st,
		Token:     tok,
	}
	err = core.HTTPSend(par)
	return st, eris.Wrap(err, "payment status")
}
