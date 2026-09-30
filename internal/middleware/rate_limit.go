package middleware

import (
	"net/http"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
)

func RateLimit(l limiter.RateLimiter, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := GetKey(r)

		allowed, err := l.Allow(key)

		if err != nil {

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)

			return
		}

		if !allowed {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)

			w.Write([]byte(`{
            	"error": "rate limit exceeded",
            	"message": "too many requests"
                 }`))

			return

		}

		next.ServeHTTP(w, r)
	})
}
