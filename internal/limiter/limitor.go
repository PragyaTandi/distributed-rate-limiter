package limiter

// RateLimiter defines the interface for rate limiting.
type RateLimiter interface {
	Allow(key string) (bool, error)
}
