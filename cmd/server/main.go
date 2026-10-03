package main

import (
	"fmt"
	"net/http"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/config"
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

	rl := redis.NewRedisRateLimiter(
		redisClient,
		cfg.Capacity,
		cfg.RefillRate,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)

	handler := middleware.RateLimit(rl, mux)

	port := "8080"

	if cfg.Port != "" {
		port = cfg.Port
	}

	fmt.Println("Server running on http://localhost:" + port)

	err := http.ListenAndServe(":"+port, handler)

	if err != nil {
		fmt.Println(err)
	}
}
