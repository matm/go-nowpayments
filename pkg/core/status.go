package core

import (
	"io/ioutil"
	"net/http"

	"github.com/rotisserie/eris"
)

// Status returns current state of the API. "OK" is returned if everything is
// fine, otherwise an error message is returned.
func Status() (string, error) {
	client := &http.Client{}

	ra := routes["status"]
	req, err := http.NewRequest(ra.method, string(defaultURL)+ra.path, nil)
	if err != nil {
		return "", eris.Wrap(err, "status")
	}
	res, err := client.Do(req)
	if err != nil {
		return "", eris.Wrap(err, "status")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", eris.Wrap(err, "status")
	}
	return string(body), nil
}
