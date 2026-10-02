package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return client
}
