package api

import (
	"strconv"
	"time"

	"github.com/imroc/req/v3"
)

// YuqueAPI yuque developer API documentation - v2
// api doc: https://www.yuque.com/yuque/developer/api
type YuqueAPI interface {
}

type Client struct {
	cfg *Config
	*req.Client
}

func NewClient(authToken string) *Client {
	cfg := NewConfig(authToken)
	client := &Client{
		cfg:    cfg,
		Client: initClient(cfg),
	}
	return client
}

func initClient(cfg *Config) *req.Client {
	client := req.C()
	options := []clientOption{
		clientOptionInit(cfg),
		clientOptionDebug(cfg),
		clientOptionRetry(cfg),
	}
	for _, option := range options {
		option(client)
	}

	return client
}

type clientOption func(c *req.Client)

func clientOptionInit(cfg *Config) clientOption {
	return func(client *req.Client) {
		client.SetBaseURL(cfg.APIBase).
			SetTimeout(cfg.Timeout).
			SetUserAgent(cfg.UserAgent).
			SetCommonHeader("X-Auth-Token", cfg.AuthToken).
			SetCommonContentType("application/json")
	}
}

func clientOptionDebug(cfg *Config) clientOption {

	return func(client *req.Client) {
		if cfg.EnableDebug {
			// DevMode enables:
			// 1. Dump content of all requests and responses to see details.
			// 2. Output debug level log for deeper insights.
			// 3. Trace all requests, so you can get trace info to analyze performance.
			client.DevMode()
		}
	}
}

func clientOptionRetry(cfg *Config) clientOption {
	return func(client *req.Client) {
		if cfg.RetryCount >= 1 {
			client.SetCommonRetryCount(cfg.RetryCount).
				// Set the retry sleep interval with a commonly used algorithm: capped exponential backoff with
				// jitter (https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/).
				SetCommonRetryBackoffInterval(1*time.Second, 5*time.Second).
				SetCommonRetryCondition(func(resp *req.Response, err error) bool {
					return resp.StatusCode >= 500
				}).
				// Set the retry to use a custom retry interval algorithm.
				SetCommonRetryInterval(func(resp *req.Response, attempt int) time.Duration {
					// Sleep seconds from "Retry-After" response header if it is present and correct.
					// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
					if resp.Response != nil {
						if ra := resp.Header.Get("Retry-After"); ra != "" {
							after, err := strconv.Atoi(ra)
							if err == nil {
								return time.Duration(after) * time.Second
							}
						}
					}
					return 2 * time.Second // Otherwise, sleep 2 seconds
				})
		}

	}
}
