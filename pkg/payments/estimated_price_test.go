package payments

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/matm/go-nowpayments/mocks"
	"github.com/matm/go-nowpayments/pkg/core"
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

func TestEstimatedPrice(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		amount       float64
		currencyFrom string
		currencyTo   string
	}
	tests := []struct {
		name  string
		args  args
		init  func(*mocks.HTTPClient)
		after func(*Estimate, error)
	}{
		{"zero amount", args{0.0, "a", "b"}, nil,
			func(e *Estimate, err error) {
				assert.Nil(e, "should return no estimate for 0.0 amount")
				assert.Error(err, "should prevent useless call to API server")
			},
		},
		{"query parameters", args{1.0, "eur", "btc"},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK("{}")
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.NotNil(req)
					assert.Equal("amount=1.000000&currency_from=eur&currency_to=btc",
						req.URL.Query().Encode(), "check query parameters")
				}).Return(resp, nil)
			}, func(e *Estimate, err error) {
				assert.NotNil(e)
				assert.NoError(err)
			},
		},
		{"route name", args{1.0, "eur", "btc"},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK("{}")
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.NotNil(req)
					assert.Equal("/v1/estimate", req.URL.Path, "bad endpoint")
				}).Return(resp, nil)

			}, func(e *Estimate, err error) {
				assert.NotNil(e)
				assert.NoError(err)
			},
		},
		{"api error", args{1.0, "eur", "btc"},
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(e *Estimate, err error) {
				assert.Nil(e)
				require.Error(t, err)
				assert.Equal("estimate: network error", err.Error())
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
			got, err := EstimatedPrice(tt.args.amount, tt.args.currencyFrom, tt.args.currencyTo)
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}

func TestRefreshEstimatedPrice(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		paymentID string
		init      func(*mocks.HTTPClient)
		after     func(*LatestEstimate, error)
	}{
		{"route name", "PID",
			func(c *mocks.HTTPClient) {
				resp := newResponseOK("{}")
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.Equal("/v1/payment/PID/update-merchant-estimate", req.URL.Path, "bad endpoint")
				}).Return(resp, nil)

			},
			nil,
		},
		{"api error", "PID",
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(e *LatestEstimate, err error) {
				assert.Nil(e)
				assert.Error(err)
				assert.Equal("last-estimate: network error", err.Error())
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
			got, err := RefreshEstimatedPrice(tt.paymentID)
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}
