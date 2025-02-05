package app

import (
	"stress-test/config"
	"time"

	"github.com/go-resty/resty/v2"
)

func InitDependencies(envs *config.Env) *resty.Client {
	return buildApiClient("", envs.ApiTimeout)
}

func buildApiClient(url string, timeout time.Duration) *resty.Client {
	c := resty.New()
	c.SetBaseURL(url).SetTimeout(timeout)

	return c
}
