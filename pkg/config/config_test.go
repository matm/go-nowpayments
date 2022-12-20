package config

import (
	"io"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	emptyAPIKeyCfg := `{"server":"http://some.tld","login":"mylogin","password":"mypass"}`
	emptyLoginCfg := `{"server":"http://some.tld","apiKey":"key","password":"mypass"}`
	emptyPasswordCfg := `{"server":"http://some.tld","login":"mylogin","apiKey":"key"}`
	emptyServerCfg := `{"apiKey":"key","login":"mylogin","password":"mypass"}`
	tests := []struct {
		name    string
		r       io.Reader
		wantErr bool
	}{
		{"nil reader", nil, true},
		{"bad config", strings.NewReader("nojson"), true},
		{"valid config", strings.NewReader(validCfg), false},
		{"empty API key", strings.NewReader(emptyAPIKeyCfg), true},
		{"empty login", strings.NewReader(emptyLoginCfg), true},
		{"empty password", strings.NewReader(emptyPasswordCfg), true},
		{"empty server", strings.NewReader(emptyServerCfg), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Load(tt.r); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var validCfg = `
{
	"server": "http://some.tld",
	"login": "mylogin",
	"password": "mypass",
	"apiKey": "key"
}
`

func TestLogin(t *testing.T) {
	Load(strings.NewReader(validCfg))
	tests := []struct {
		name string
		want string
	}{
		{"login value", "mylogin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Login(); got != tt.want {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword(t *testing.T) {
	Load(strings.NewReader(validCfg))
	tests := []struct {
		name string
		want string
	}{
		{"password value", "mypass"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Password(); got != tt.want {
				t.Errorf("Password() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIKey(t *testing.T) {
	Load(strings.NewReader(validCfg))
	tests := []struct {
		name string
		want string
	}{
		{"key value", "key"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := APIKey(); got != tt.want {
				t.Errorf("APIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer(t *testing.T) {
	Load(strings.NewReader(validCfg))
	tests := []struct {
		name string
		want string
	}{
		{"server url", "http://some.tld"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Server(); got != tt.want {
				t.Errorf("Server() = %v, want %v", got, tt.want)
			}
		})
	}
}
