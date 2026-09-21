package middleware

import (
	"net/http"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
)

func RateLimit(l limiter.RateLimiter, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.RemoteAddr

		allowed, err := l.Allow(key)

		if err != nil {

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)

			return
		}

		if !allowed {

			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)

			return
		}

		next.ServeHTTP(w, r)
	})
}
