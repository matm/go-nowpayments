package core

type st struct {
	Message string `json:"message"`
}

// Status returns current state of the API. "OK" is returned if everything is
// fine, otherwise an error message is returned.
func Status() (string, error) {
	s := &st{}
	par := &SendParams{
		RouteName: "status",
		Into:      &s,
	}
	err := HTTPSend(par)
	if err != nil {
		return "", err
	}
	return s.Message, nil
}
