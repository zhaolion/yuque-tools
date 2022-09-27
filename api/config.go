package api

import (
	"time"
)

const (
	// URIBase https://www.yuque.com/yuque/developer/api#c8c7d76f
	URIBase = "https://www.yuque.com/api/v2"
	// UserAgentDefault default client headers - UserAgent
	// To ensure that the server knows who the visitor is.
	UserAgentDefault = "client-go-yuque-tools"
	// TimeoutDefault default timeout for all requests.
	TimeoutDefault = 5 * time.Second
	// RetryCountDefault default number of retries to retry
	RetryCountDefault = 5
)

type Config struct {
	AuthToken   string
	APIBase     string
	UserAgent   string
	Timeout     time.Duration
	EnableDebug bool
	RetryCount  int
}

func NewConfig(authToken string) *Config {
	cfg := &Config{AuthToken: authToken}
	cfg.initDefault()
	return cfg
}

func (cfg *Config) initDefault() *Config {
	if cfg.APIBase == "" {
		cfg.APIBase = URIBase
	}
	if cfg.UserAgent == "" {
		cfg.UserAgent = UserAgentDefault
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = TimeoutDefault
	}
	if cfg.RetryCount == 0 {
		cfg.RetryCount = RetryCountDefault
	}

	return cfg
}
