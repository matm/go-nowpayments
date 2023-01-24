package payments

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/matm/go-nowpayments/core"
	"github.com/matm/go-nowpayments/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestList(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		o     *ListOption
		init  func(*mocks.HTTPClient)
		after func([]*Payment, error)
	}{
		{"route and response", nil,
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Call.Return(
					func(req *http.Request) *http.Response {
						switch req.URL.Path {
						case "/v1/auth":
							return newResponseOK(`{"token":"tok"}`)
						case "/v1/payment/":
							return newResponseOK(`[{"payment_id":1}]`)
						default:
							t.Fatalf("unexpected route call %q", req.URL.Path)
						}
						return nil
					}, nil)
			},
			func(ps []*Payment, err error) {
				assert.NoError(err)
				if assert.Len(ps, 1) {
					assert.Equal("1", ps[0].ID)
				}
			}},
		{"payment_id as a string", nil,
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Call.Return(
					func(req *http.Request) *http.Response {
						switch req.URL.Path {
						case "/v1/auth":
							return newResponseOK(`{"token":"tok"}`)
						case "/v1/payment/":
							return newResponseOK(`[{"payment_id":"54321"}]`)
						default:
							t.Fatalf("unexpected route call %q", req.URL.Path)
						}
						return nil
					}, nil)
			},
			func(ps []*Payment, err error) {
				assert.NoError(err)
				if assert.Len(ps, 1) {
					assert.Equal("54321", ps[0].ID)
				}
			}},
		{"api error", nil,
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Call.Return(
					func(req *http.Request) *http.Response {
						switch req.URL.Path {
						case "/v1/auth":
							return newResponseOK(`{"token":"tok"}`)
						}
						return nil
					},
					func(req *http.Request) error {
						switch req.URL.Path {
						case "/v1/payment/":
							return errors.New("network error")
						}
						return nil
					},
				)
			}, nil},
		{"auth fail", nil,
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("bad credentials"))
			},
			func(ps []*Payment, err error) {
				assert.Nil(ps)
				assert.Error(err)
				assert.Equal("list: auth: bad credentials", err.Error())
			},
		},
		{"with all options", &ListOption{
			Limit:    2,
			DateFrom: "2020-01-01",
			DateTo:   "2022-01-01",
			OrderBy:  "asc",
			SortBy:   "created_at",
			Page:     3,
		},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"data":[{"payment_id":1}]}`)
				c.EXPECT().Do(mock.Anything).Run(func(r *http.Request) {
					if strings.HasPrefix(r.URL.Path, "/v1/payment") {
						assert.Equal("dateFrom=2020-01-01&dateTo=2022-01-01&limit=2&orderBy=asc&page=3&sortBy=created_at", r.URL.RawQuery)
					}
				}).Return(resp, nil)
			},
			nil,
		},
		{"with empty options", &ListOption{},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"data":[{"payment_id":1}]}`)
				c.EXPECT().Do(mock.Anything).Run(func(r *http.Request) {
					if strings.HasPrefix(r.URL.Path, "/v1/payment") {
						assert.Equal("page=0", r.URL.RawQuery)
					}
				}).Return(resp, nil)
			},
			nil,
		},
		{"with some options", &ListOption{Limit: 5, OrderBy: "desc"},
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"data":[{"payment_id":1}]}`)
				c.EXPECT().Do(mock.Anything).Run(func(r *http.Request) {
					if strings.HasPrefix(r.URL.Path, "/v1/payment") {
						assert.Equal("limit=5&orderBy=desc&page=0", r.URL.RawQuery)
					}
				}).Return(resp, nil)
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mocks.NewHTTPClient(t)
			core.UseClient(c)
			if tt.init != nil {
				tt.init(c)
			}
			got, err := List(tt.o)
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}
