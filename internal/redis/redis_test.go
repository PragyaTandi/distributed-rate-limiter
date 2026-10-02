package redis_test

import (
	"testing"

	"github.com/PragyaTandi/distributed-rate-limiter/internal/redis"
)

func TestRedisConnection(t *testing.T) {

	client := redis.NewClient()

	err := client.Ping(redis.Ctx).Err()

	if err != nil {
		t.Fatal(err)
	}
}
