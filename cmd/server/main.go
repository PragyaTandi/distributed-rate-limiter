package main

import (
	"fmt"
	"net/http"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/config"
	"github.com/PragyaTandi/distributed-rate-limiter/internal/middleware"
	"github.com/PragyaTandi/distributed-rate-limiter/internal/redis"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// App handlers wrapped with rate limit middleware
	appMux := http.NewServeMux()
	appMux.HandleFunc("GET /{$}", homeHandler)
	appMux.HandleFunc("GET /health", healthHandler)

	// Main server mux
	mainMux := http.NewServeMux()
	mainMux.Handle("GET /metrics", promhttp.Handler())
	mainMux.Handle("/", middleware.RateLimit(rl, appMux))

	port := "8080"
	if cfg.Port != "" {
		port = cfg.Port
	}

	fmt.Println("Server running on http://localhost:" + port)

	err := http.ListenAndServe(":"+port, mainMux)
	if err != nil {
		fmt.Println(err)
	}
}
