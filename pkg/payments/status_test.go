package payments

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/matm/go-nowpayments/mocks"
	"github.com/matm/go-nowpayments/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name      string
		paymentID string
		init      func(*mocks.HTTPClient)
		after     func(*mocks.HTTPClient, *PaymentStatus, error)
	}{
		{"empty payment ID", "", nil,
			func(c *mocks.HTTPClient, s *PaymentStatus, err error) {
				require.Error(t, err)
				assert.Nil(s)
			},
		},
		{"authentication ok", "ID",
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Call.Return(
					func(req *http.Request) *http.Response {
						switch req.URL.Path {
						case "/v1/auth":
							return newResponseOK(`{"token":"tok"}`)
						case "/v1/payment":
							return newResponseOK(`{"payment_status":"done","pay_amount":10.0}`)
						default:
							t.Fatalf("unexpected route call %q", req.URL.Path)
						}
						return nil
					}, nil)
			},
			func(c *mocks.HTTPClient, s *PaymentStatus, err error) {
				assert.NoError(err)
				assert.NotNil(s)
				assert.Equal(10.0, s.PayAmount)
				assert.Equal("done", s.Status)
				c.AssertNumberOfCalls(t, "Do", 2)
			},
		},
		{"authentication call failed", "ID",
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Call.Return(
					func(req *http.Request) *http.Response {
						switch req.URL.Path {
						case "/v1/auth":
							return newResponse(http.StatusForbidden, "")
						default:
							t.Fatalf("unexpected route call %q", req.URL.Path)
						}
						return nil
					}, errors.New("bad credentials"))
			},
			func(c *mocks.HTTPClient, s *PaymentStatus, err error) {
				assert.Error(err)
				assert.Nil(s)
				assert.Equal("status: auth: bad credentials", err.Error())
				c.AssertNumberOfCalls(t, "Do", 1)
			},
		},
		{"status call failed", "ID",
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Call.Return(
					func(req *http.Request) *http.Response {
						switch req.URL.Path {
						case "/v1/auth":
							return newResponseOK(`{"token":"tok"}`)
						case "/v1/payment":
							return newResponse(http.StatusInternalServerError, "")
						default:
							return nil
						}
					},
					func(req *http.Request) error {
						switch req.URL.Path {
						case "/v1/auth":
							return nil
						case "/v1/payment":
							return errors.New("network error")
						default:
							return fmt.Errorf("unexpected route call %q", req.URL.Path)
						}
					},
				)
			},
			func(c *mocks.HTTPClient, s *PaymentStatus, err error) {
				assert.Error(err)
				assert.Nil(s)
				assert.Equal("payment-status: network error", err.Error())
				c.AssertNumberOfCalls(t, "Do", 2)
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
			got, err := Status(tt.paymentID)
			if tt.after != nil {
				tt.after(c, got, err)
			}
		})
	}
}
