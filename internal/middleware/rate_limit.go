package middleware

import (
	"net/http"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
	"github.com/PragyaTandi/distributed-rate-limiter/internal/metrics"
)

func RateLimit(rl limiter.RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		if key == "" {
			key = r.RemoteAddr
		}

		allowed, err := rl.Allow(key)
		if err != nil || !allowed {
			metrics.BlockedRequests.Inc()
			http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
			return
		}

		metrics.AllowedRequests.Inc()
		next.ServeHTTP(w, r)
	})
}
