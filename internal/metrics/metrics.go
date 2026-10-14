package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	AllowedRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rate_limiter_allowed_total",
		Help: "Total number of requests allowed by the rate limiter.",
	})

	BlockedRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rate_limiter_blocked_total",
		Help: "Total number of requests blocked by the rate limiter.",
	})
)
