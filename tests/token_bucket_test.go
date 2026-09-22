package tests

import (
	"testing"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
)

func TestTokenBucket(t *testing.T) {

	tb := limiter.NewTokenBucket(5, 1)

	for i := 0; i < 5; i++ {
		if !tb.Allow() {
			t.Fatalf("expected request %d to be allowed", i+1)
		}
	}

	if tb.Allow() {
		t.Fatal("expected sixth request to be rejected")
	}
}
