package main

import (
	"fmt"
	"net/http"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/config"
	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
	"github.com/PragyaTandi/distributed-rate-limiter/internal/middleware"
	"github.com/PragyaTandi/distributed-rate-limiter/internal/redis"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Distributed Rate Limiter is Running!")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {

	cfg := config.LoadConfig()

	redisClient := redis.NewClient()

	rl := limiter.NewRedisRateLimiter(
		redisClient,
		cfg.Capacity,
		cfg.RefillRate,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)

	handler := middleware.RateLimit(rl, mux)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		fmt.Println(err)
	}
}
