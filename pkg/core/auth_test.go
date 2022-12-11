package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/matm/go-nowpayments/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAuthenticate(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name  string
		args  args
		init  func(*mocks.HTTPClient)
		after func(string, error)
	}{
		{"request body and route", args{"a", "b"}, func(c *mocks.HTTPClient) {
			resp := newResponseOK(`{"token":"tok"}`)
			c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
				assert.NotNil(req)
				assert.Equal("/v1/auth", req.URL.Path, "bad endpoint")
				// Check request body.
				d, err := ioutil.ReadAll(req.Body)
				require.NoError(err)
				type auth struct {
					Email    string
					Password string
				}
				var body auth
				err = json.Unmarshal(d, &body)
				require.NoError(err)
				assert.Equal("a", body.Email)
				assert.Equal("b", body.Password)
			}).Return(resp, nil)
		}, nil,
		},
		{"api error", args{"a", "b"},
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			},
			func(token string, err error) {
				assert.Empty(token)
				require.Error(err)
				assert.Equal("auth: network error", err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mocks.NewHTTPClient(t)
			UseClient(c)
			if tt.init != nil {
				tt.init(c)
			}
			got, err := Authenticate(tt.args.email, tt.args.password)
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}
