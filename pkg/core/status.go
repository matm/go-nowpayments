package core

import (
	"github.com/rotisserie/eris"
)

type st struct {
	Message string `json:"message"`
}

// Status returns current state of the API. "OK" is returned if everything is
// fine, otherwise an error message is returned.
func Status() (string, error) {
	s := &st{}
	err := HTTPSend("status", nil, &s)
	if err != nil {
		return "", eris.Wrap(err, "status")
	}
	return s.Message, nil
}
