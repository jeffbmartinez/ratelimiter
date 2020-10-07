package ratelimiter

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RateLimiterConfig struct {
	Skipper middleware.Skipper

	BucketSize       int
	TokensPerSecond  int
	InitialNumTokens int
}

func RateLimiterWithConfig(config RateLimiterConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}

	tokenBucket := NewTokenBucket(config.BucketSize, config.TokensPerSecond, config.InitialNumTokens)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			if tokenBucket.RequestToken() {
				return next(c)
			}

			return echo.ErrTooManyRequests
		}
	}
}
