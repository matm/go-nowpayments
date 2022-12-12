package payments

import (
	"errors"
	"net/http"
	"testing"

	"github.com/matm/go-nowpayments/mocks"
	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		pa    *PaymentArgs
		init  func(*mocks.HTTPClient)
		after func(*Payment, error)
	}{
		{"nil args", nil, nil,
			func(p *Payment, err error) {
				assert.Nil(p)
				assert.Error(err)
			},
		},
		{"api error", &PaymentArgs{PurchaseID: "1234"},
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(p *Payment, err error) {
				assert.Nil(p)
				assert.Error(err)
				assert.Equal("payment-create: network error", err.Error())
			},
		},
		{"valid args", &PaymentArgs{
			PurchaseID:    "1234",
			PaymentAmount: PaymentAmount{PriceAmount: 10.0},
		},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234"}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			}, func(p *Payment, err error) {
				assert.NoError(err)
				assert.NotNil(p)
				assert.Equal("1234", p.ID)
			},
		},
		{"route check", &PaymentArgs{},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"payment_id":"1234"}`)
				c.EXPECT().Do(mock.Anything).Run(func(r *http.Request) {
					assert.Equal("/v1/payment", r.URL.Path, "bad endpoint")
				}).Return(resp, nil)
			}, func(p *Payment, err error) {
				assert.NoError(err)
				assert.NotNil(p)
				assert.Equal("1234", p.ID)
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
			got, err := New(tt.pa)
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}
