package tests

import (
	"sync"
	"testing"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
)

func TestConcurrentRequests(t *testing.T) {

	rl := limiter.NewMemoryRateLimiter(10, 0)

	var wg sync.WaitGroup

	allowedCount := 0
	var mutex sync.Mutex

	totalRequests := 100

	for i := 0; i < totalRequests; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			allowed, _ := rl.Allow("user1")

			if allowed {
				mutex.Lock()
				allowedCount++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()

	if allowedCount != 10 {
		t.Errorf(
			"expected 10 allowed requests, got %d",
			allowedCount,
		)
	}
}
