package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/matm/go-nowpayments/mocks"
	"github.com/matm/go-nowpayments/pkg/config"
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

func conf() *strings.Reader {
	return strings.NewReader(`
{
	"login":"l","password":"p","apiKey":"key","server":"http://some.tld"
}
`)
}

func init() {
	err := config.Load(conf())
	if err != nil {
		panic(err)
	}
}

func TestHTTPSend(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	defaultURL = "host"

	tests := []struct {
		name    string
		p       *SendParams
		wantErr bool
		init    func(*mocks.HTTPClient)
		after   func(*SendParams, error)
	}{
		{"nil params", nil, true, nil, nil},
		{"unknown route", &SendParams{RouteName: "bad"}, true, nil, nil},
		{"no url values", &SendParams{RouteName: "status"}, false,
			func(c *mocks.HTTPClient) {
				resp := newResponseOK("{}")
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.NotNil(req)
					assert.Equal("host/status", req.URL.Path)
					assert.Equal("key", req.Header.Get("X-API-KEY"))
				}).Return(resp, nil)
			},
			nil,
		},
		{"with url values", &SendParams{
			RouteName: "status",
			Values:    url.Values{"a": []string{"1"}, "b": []string{"2", "3"}},
		}, false,
			func(c *mocks.HTTPClient) {
				resp := newResponseOK("{}")
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.Equal("host/status", req.URL.Path)
					assert.Equal("a=1&b=2&b=3", req.URL.RawQuery)
				}).Return(resp, nil)
			},
			nil,
		},
		{"with auth token", &SendParams{
			RouteName: "status",
			Token:     "token",
		}, false,
			func(c *mocks.HTTPClient) {
				resp := newResponseOK("{}")
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.Equal("Bearer token", req.Header.Get("Authorization"))
				}).Return(resp, nil)
			},
			nil,
		},
		{"req exec error", &SendParams{RouteName: "status"}, true,
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			},
			nil,
		},
		{"with a request body", &SendParams{
			RouteName: "status",
			Body:      strings.NewReader("body"),
		}, false,
			func(c *mocks.HTTPClient) {
				resp := newResponseOK("{}")
				c.EXPECT().Do(mock.Anything).Run(func(req *http.Request) {
					assert.Equal("application/json", req.Header.Get("Content-Type"))
				}).Return(resp, nil)
			},
			nil,
		},
		{"response body", &SendParams{RouteName: "status"}, false,
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"some":"value"}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			},
			func(p *SendParams, err error) {
				require.NotEmpty(p.Into)
				data, err := json.Marshal(p.Into)
				require.NoError(err)
				assert.Equal(`{"some":"value"}`, string(data))
			},
		},
		{"error status code", &SendParams{RouteName: "status"}, true,
			func(c *mocks.HTTPClient) {
				resp := newResponse(http.StatusInternalServerError, `
				{
					"statusCode": 500,
					"code": "server error",
					"message": "damn"
				}
				`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			},
			func(p *SendParams, err error) {
				assert.Equal("code 500 (server error): damn", err.Error())
			},
		},
		{"decode error status", &SendParams{RouteName: "status"}, true,
			func(c *mocks.HTTPClient) {
				resp := newResponse(http.StatusInternalServerError, "bad")
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			},
			func(p *SendParams, err error) {
				assert.Equal("status: JSON decode error: invalid character 'b' looking for beginning of value", err.Error())
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
			err := HTTPSend(tt.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPSend() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.after != nil {
				tt.after(tt.p, err)
			}
		})
	}
}

func TestWithDebug(t *testing.T) {
	type args struct {
		d bool
	}
	tests := []struct {
		name  string
		args  args
		after func()
	}{
		{"debug on", args{true}, func() {
			assert.Equal(t, true, debug)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WithDebug(tt.args.d)
			if tt.after != nil {
				tt.after()
			}
		})
	}
}

func TestUseBaseURL(t *testing.T) {
	type args struct {
		b BaseURL
	}
	tests := []struct {
		name  string
		args  args
		after func()
	}{
		{"set url", args{ProductionBaseURL}, func() {
			assert.Equal(t, ProductionBaseURL, defaultURL)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UseBaseURL(tt.args.b)
			if tt.after != nil {
				tt.after()
			}
		})
	}
}
