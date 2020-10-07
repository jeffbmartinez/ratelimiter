# Rate Limiter for [Echo web framework](https://echo.labstack.com/)

This rate limiter middleware was built and tested against Echo v4.

I believe the middleware interface didn't change between v3 and v4, so it should also work for v3.

This middleware uses a token bucket algorithm for rate limiting.

## Token Bucket Algorithm Description

The token bucket is an analogy which helps explain when a client's request will be denied. A client starts with a bucket of tokens. Every time a client makes a request it costs exactly one token. Requests will be served as long as the client has tokens to spend. If the client uses up all the tokens it will receive an `HTTP 429 Too Many Requests` response. Tokens are put back in a client's bucket at a steady rate, based on some configureable timer.

For example, a client with who has saved 100 tokens can make up to 100 requests as fast as they send them. When they run out of tokens they will receive `HTTP 429` responses. At that point, if the client's token refill rate is 2 per second, they are effectively limited to making no more than 2 requests per second.

## Usage

```
import (
  "github.com/jeffbmartinez/ratelimiter"
  "github.com/labstack/echo"
)

// ...

e := echo.New()
e.Use(ratelimiter.RateLimiterWithConfig(ratelimiter.RateLimiterConfig{
  BucketSize:      5,
  TokensPerSecond: 1,
  InitialNumTokens: 0,
}))
```

## Configuration

* BucketSize - Maximum tokens a client can store at a time.
* TokensPerSecond - The token refresh rate. This many tokens are added back to the client's bucket each second.
* InitialNumTokens - The initial number of tokens in a client's bucket. If this is set to a number greater than `BucketSize`, it gets set to `BucketSize`.
* Skipper - Standard Echo middleware Skipper function (If this function returns false the middleware is skipped)
