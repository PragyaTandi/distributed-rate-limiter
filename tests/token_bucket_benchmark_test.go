package tests

import (
	"testing"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
)

func BenchmarkTokenBucket(b *testing.B) {
	tb := limiter.NewTokenBucket(1000000, 1000000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tb.Allow()
	}
}
