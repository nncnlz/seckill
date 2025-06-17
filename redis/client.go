package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-12436.c1.asia-northeast1-1.gce.redns.redis-cloud.com:12436",
		Username: "default",
		Password: "z8MWA0o7QTzseGuNtY5NIW1SaEF1bvcn",
		DB:       0,
	})
}
