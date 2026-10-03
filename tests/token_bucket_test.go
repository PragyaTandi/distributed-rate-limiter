package tests

import (
	"testing"
	"time"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
)

func TestTokenBucket(t *testing.T) {

	tb := limiter.NewTokenBucket(5, 1)

	for i := 0; i < 5; i++ {

		allowed, _ := tb.Allow("user1")

		if !allowed {
			t.Fatalf("expected request %d to be allowed", i+1)
		}
	}

	allowed, _ := tb.Allow("user1")

	if allowed {
		t.Fatal("expected sixth request to be rejected")
	}
}

func TestTokenBucketRefill(t *testing.T) {

	tb := limiter.NewTokenBucket(2, 1)

	// Consume both tokens.
	tb.Allow("user1")
	tb.Allow("user1")

	// Third request should fail.
	allowed, _ := tb.Allow("user1")

	if allowed {
		t.Fatal("expected request to be rejected")
	}

	// Wait for one token to refill.
	time.Sleep(1100 * time.Millisecond)

	// Now one request should succeed.
	allowed, _ = tb.Allow("user1")

	if !allowed {
		t.Fatal("expected request after refill to be allowed")
	}
}
