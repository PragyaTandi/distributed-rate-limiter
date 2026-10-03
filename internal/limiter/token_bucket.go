package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity       int
	tokens         float64
	refillRate     float64
	lastRefillTime time.Time
	mutex          sync.Mutex
}

func NewTokenBucket(capacity int, refillRate float64) *TokenBucket {

	return &TokenBucket{
		capacity:       capacity,
		tokens:         float64(capacity),
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) Allow(key string) (bool, error) {

	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	now := time.Now()

	elapsed := now.Sub(tb.lastRefillTime).Seconds()

	tb.tokens += elapsed * tb.refillRate

	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}

	tb.lastRefillTime = now

	if tb.tokens >= 1 {

		tb.tokens--

		return true, nil
	}

	return false, nil
}
