package ratelimiter

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens          int
	bucketSize      int
	tokensPerSecond int

	mux sync.Mutex
}

func NewTokenBucket(bucketSize int, tokensPerSecond int, initialNumTokens int) *TokenBucket {
	if initialNumTokens > bucketSize {
		initialNumTokens = bucketSize
	}

	tb := &TokenBucket{
		tokens:          initialNumTokens,
		bucketSize:      bucketSize,
		tokensPerSecond: tokensPerSecond,
	}

	go tb.tokenRefiller()

	return tb
}

func (tb *TokenBucket) tokenRefiller() {
	for range time.Tick(1 * time.Second) {
		numTokens := tb.AddTokens(tb.tokensPerSecond)
		fmt.Printf("Added more tokens, now have %v\n", numTokens)
	}
}

func (tb *TokenBucket) AddTokens(numTokens int) int {
	tb.mux.Lock()
	defer tb.mux.Unlock()

	tb.tokens += numTokens
	if tb.tokens > tb.bucketSize {
		tb.tokens = tb.bucketSize
	}

	return tb.tokens
}

func (tb *TokenBucket) RequestToken() bool {
	tb.mux.Lock()
	defer tb.mux.Unlock()

	if tb.tokens == 0 {
		return false
	}

	tb.tokens -= 1
	return true
}
