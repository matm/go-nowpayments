package payments

import (
	"errors"
	"net/http"
	"testing"

	"github.com/matm/go-nowpayments/mocks"
	"github.com/matm/go-nowpayments/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewInvoice(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		ia    *InvoiceArgs
		init  func(*mocks.HTTPClient)
		after func(*Invoice, error)
	}{
		{"nil args", nil, nil,
			func(e *Invoice, err error) {
				assert.Nil(e)
				assert.Error(err)
				assert.Equal("nil invoice args", err.Error())
			},
		},
		{"api error", &InvoiceArgs{},
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(e *Invoice, err error) {
				assert.Nil(e)
				assert.Error(err)
				assert.Equal("invoice-create: network error", err.Error())
			},
		},
		{"route name", &InvoiceArgs{},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK("{}")
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.Equal("/v1/invoice", req.URL.Path, "bad endpoint")
				}).Return(resp, nil)
			},
			nil,
		},
		{"some data", &InvoiceArgs{},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"id":"ID","price_amount":"5.00"}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			},
			func(e *Invoice, err error) {
				assert.NotNil(e)
				assert.NoError(err)
				assert.Equal("ID", e.ID)
				assert.Equal("5.00", e.PriceAmount)
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
			got, err := NewInvoice(tt.ia)
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}
