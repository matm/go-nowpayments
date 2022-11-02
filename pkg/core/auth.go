package core

import (
	"fmt"
	"strings"

	"github.com/matm/go-nowpayments/pkg/types"
	"github.com/rotisserie/eris"
)

type token struct {
	Token string `json:"token"`
}

// Authenticate is used for obtaining a JWT token. Such a token is required for some API calls
// like payment status or create payment.
func Authenticate(email, password string) (string, error) {
	r := strings.NewReader(fmt.Sprintf(`{
			"email": "%s",
			"password": "%s"
		}`, email, password))
	t := &token{}
	par := &types.SendParams{
		RouteName: "auth",
		Body:      r,
		Into:      &t,
	}
	err := HTTPSend(par)
	return t.Token, eris.Wrap(err, "auth")
}
