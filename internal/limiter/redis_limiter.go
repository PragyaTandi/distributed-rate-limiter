package limiter

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisRateLimiter struct {
	client     *redis.Client
	capacity   int
	refillRate float64
}

func NewRedisRateLimiter(
	client *redis.Client,
	capacity int,
	refillRate float64,
) *RedisRateLimiter {

	return &RedisRateLimiter{
		client:     client,
		capacity:   capacity,
		refillRate: refillRate,
	}
}

func (r *RedisRateLimiter) Allow(key string) (bool, error) {

	ctx := context.Background()

	redisKey := "rate_limit:" + key

	tokens, err := r.client.Get(ctx, redisKey).Int()

	if err == redis.Nil {

		err = r.client.Set(
			ctx,
			redisKey,
			r.capacity-1,
			0,
		).Err()

		if err != nil {
			return false, err
		}

		return true, nil
	}

	if err != nil {
		return false, err
	}

	if tokens <= 0 {
		return false, nil
	}

	err = r.client.Set(
		ctx,
		redisKey,
		tokens-1,
		0,
	).Err()

	if err != nil {
		return false, err
	}

	return true, nil
}
