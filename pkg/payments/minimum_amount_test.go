package payments

import (
	"errors"
	"net/http"
	"testing"

	"github.com/matm/go-nowpayments/mocks"
	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMinimumAmount(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		currencyFrom string
		currencyTo   string
	}
	tests := []struct {
		name  string
		args  args
		init  func(*mocks.HTTPClient)
		after func(*CurrencyAmount, error)
	}{
		{"query parameters and route", args{"eur", "btc"},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"currency_from":"eur","currency_to":"btc","min_amount":1.0}`)
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.Equal("currency_from=eur&currency_to=btc",
						req.URL.Query().Encode(), "check query parameters")
					assert.Equal("/v1/min-amount", req.URL.Path, "bad endpoint")
				}).Return(resp, nil)
			}, func(c *CurrencyAmount, err error) {
				assert.NotNil(c)
				assert.NoError(err)
				assert.Equal("eur", c.CurrencyFrom)
				assert.Equal("btc", c.CurrencyTo)
				assert.Equal(1.0, c.Amount)
			},
		},
		{"api error", args{"eur", "btc"},
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(c *CurrencyAmount, err error) {
				assert.Nil(c)
				require.Error(t, err)
				assert.Equal("min-amount: network error", err.Error())
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
			got, err := MinimumAmount(tt.args.currencyFrom, tt.args.currencyTo)
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}
