package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
	"github.com/PragyaTandi/distributed-rate-limiter/internal/middleware"
)

func TestRateLimitMiddleware(t *testing.T) {

	// small limit for testing
	rl := limiter.NewTokenBucket(2, 0)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	wrappedHandler := middleware.RateLimit(rl, handler)

	req1 := httptest.NewRequest(
		"GET",
		"/",
		nil,
	)

	req1.Header.Set("X-API-Key", "user1")

	rec1 := httptest.NewRecorder()

	wrappedHandler.ServeHTTP(rec1, req1)

	if rec1.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec1.Code)
	}

	req2 := httptest.NewRequest(
		"GET",
		"/",
		nil,
	)

	req2.Header.Set("X-API-Key", "user1")

	rec2 := httptest.NewRecorder()

	wrappedHandler.ServeHTTP(rec2, req2)

	if rec2.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec2.Code)
	}

	req3 := httptest.NewRequest(
		"GET",
		"/",
		nil,
	)

	req3.Header.Set("X-API-Key", "user1")

	rec3 := httptest.NewRecorder()

	wrappedHandler.ServeHTTP(rec3, req3)

	if rec3.Code != http.StatusTooManyRequests {
		t.Errorf("expected 429, got %d", rec3.Code)
	}
}
