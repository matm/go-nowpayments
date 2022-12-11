package core

import (
	"errors"
	"testing"

	"github.com/matm/go-nowpayments/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name  string
		init  func(*mocks.HTTPClient)
		after func(string, error)
	}{
		{"api error",
			func(c *mocks.HTTPClient) {
				c.EXPECT().Do(mock.Anything).Return(nil, errors.New("network error"))
			}, func(s string, err error) {
				assert.Empty(s)
				require.Error(t, err)
				assert.Equal("status: network error", err.Error())
			},
		},
		{"status OK",
			func(c *mocks.HTTPClient) {
				resp := newResponseOK(`{"message":"OK"}`)
				c.EXPECT().Do(mock.Anything).Return(resp, nil)
			}, func(s string, err error) {
				assert.Equal("OK", s)
				assert.NoError(err)
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
			got, err := Status()
			if tt.after != nil {
				tt.after(got, err)
			}
		})
	}
}
