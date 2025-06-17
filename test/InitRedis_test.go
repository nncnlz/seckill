package test

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestInitRedis(t *testing.T) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-12436.c1.asia-northeast1-1.gce.redns.redis-cloud.com:12436",
		Username: "default",
		Password: "z8MWA0o7QTzseGuNtY5NIW1SaEF1bvcn",
		DB:       0,
	})

	// 测试连接
	if err := rdb.Ping(ctx).Err(); err != nil {
		t.Fatalf("Redis 连接失败: %v", err)
	}

	// 测试 SET 和 GET
	key := "foo"
	value := "bar"
	if err := rdb.Set(ctx, key, value, 0).Err(); err != nil {
		t.Fatalf("Redis SET 失败: %v", err)
	}

	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		t.Fatalf("Redis GET 失败: %v", err)
	}

	if result != value {
		t.Errorf("GET 结果错误：期望 %s，实际 %s", value, result)
	}
}
