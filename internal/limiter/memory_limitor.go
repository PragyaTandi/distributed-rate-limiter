package limiter

import "sync"

type MemoryRateLimiter struct {
	buckets map[string]*TokenBucket
	mutex   sync.Mutex

	capacity   int
	refillRate float64
}

func NewMemoryRateLimiter(capacity int, refillRate float64) *MemoryRateLimiter {
	return &MemoryRateLimiter{
		buckets:    make(map[string]*TokenBucket),
		capacity:   capacity,
		refillRate: refillRate,
	}
}

func (m *MemoryRateLimiter) Allow(key string) (bool, error) {

	m.mutex.Lock()

	bucket, exists := m.buckets[key]

	if !exists {
		bucket = NewTokenBucket(
			m.capacity,
			m.refillRate,
		)

		m.buckets[key] = bucket
	}

	m.mutex.Unlock()

	return bucket.Allow(), nil
}
