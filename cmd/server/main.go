package main

import (
	"fmt"
	"net/http"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/config"
	"github.com/PragyaTandi/distributed-rate-limiter/internal/limiter"
	"github.com/PragyaTandi/distributed-rate-limiter/internal/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Distributed Rate Limiter is Running!")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {

	cfg := config.LoadConfig()

	rl := limiter.NewMemoryRateLimiter(
		cfg.Capacity,
		cfg.RefillRate,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	handler := middleware.RateLimit(rl, mux)
	mux.HandleFunc("/health", healthHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		fmt.Println(err)
	}
}
