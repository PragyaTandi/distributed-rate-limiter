package redis

import (
	"context"
	_ "embed"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

//go:embed token_bucket.lua
var tokenBucketScript string

type RedisRateLimiter struct {
	client *redis.Client

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

	redisKey := "rate_limit:" + key

	now := time.Now().Unix()

	result, err := r.client.Eval(
		ctx,
		tokenBucketScript,
		[]string{redisKey},
		r.capacity,
		r.refillRate,
		now,
	).Result()

	if err != nil {
		return false, err
	}

	return result.(int64) == 1, nil
}
