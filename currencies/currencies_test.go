package currencies

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/matm/go-nowpayments/mocks"
	"github.com/matm/go-nowpayments/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type rc struct {
	*strings.Reader
}

func (*rc) Close() error {
	return nil
}

func newResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Body:       &rc{strings.NewReader(body)},
	}
}

func newResponseOK(body string) *http.Response {
	return newResponse(http.StatusOK, body)
}

func TestAll(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		init  func(*mocks.HTTPClient)
		after func([]string, error)
	}{
		{"route name",
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"currencies":["xmr","xno"]}`)
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.NotNil(req)
					assert.Equal("/v1/currencies", req.URL.Path, "bad endpoint")
				}).Return(resp, nil)
			}, func(c []string, err error) {
				require.NoError(t, err)
				assert.EqualValues([]string{"xmr", "xno"}, c)
			},
		},
		{"api error",
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(c []string, err error) {
				assert.Nil(c)
				require.Error(t, err)
				assert.Equal("currencies: network error", err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mocks.NewHTTPClient(t)
			core.UseClient(c)
			if tt.init != nil {
				tt.init(c)
			}
			got, err := All()
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}

func TestSelected(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		init  func(*mocks.HTTPClient)
		after func([]string, error)
	}{
		{"route name",
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"selectedCurrencies":["xmr","xno"]}`)
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.Equal("/v1/merchant/coins", req.URL.Path, "bad endpoint")
				}).Return(resp, nil)
			}, func(c []string, err error) {
				assert.NoError(err)
				assert.EqualValues([]string{"xmr", "xno"}, c)
			},
		},
		{"api error",
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(c []string, err error) {
				assert.Nil(c)
				require.Error(t, err)
				assert.Equal("selected-currencies: network error", err.Error())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mocks.NewHTTPClient(t)
			core.UseClient(c)
			if tt.init != nil {
				tt.init(c)
			}
			got, err := Selected()
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}
