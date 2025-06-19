package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

// RateLimit 固定窗口限流
func RateLimit(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user_id is required"})
			return
		}
		key := fmt.Sprintf("rl:%s", userID)
		count, err := RedisClient.Incr(Ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "rate limit redis error"})
			return
		}
		if count == 1 {
			RedisClient.Expire(Ctx, key, window)
		}
		if count > int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit"})
			return
		}
		c.Next()
	}
}

// RateLimit2 滑动窗口限流
func RateLimit2(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user_id is required"})
			return
		}
		key := fmt.Sprintf("rl:sliding:%s", userID)
		now := time.Now().UnixNano() / int64(time.Millisecond)
		windowSize := int64(window / time.Millisecond)
		//这是一秒前的时间搓
		minTime := now - windowSize

		pipe := RedisClient.TxPipeline()
		// 1. 把当前时间戳加入 ZSet
		pipe.ZAdd(Ctx, key, redis.Z{Score: float64(now), Member: now})
		// 2. 删除所有“早于窗口起点”的请求记录
		pipe.ZRemRangeByScore(Ctx, key, "-inf", fmt.Sprintf("%d", minTime))
		// 3. 获取当前窗口内请求数量
		countCmd := pipe.ZCard(Ctx, key)

		// 4. 设置 ZSet 过期时间（避免数据堆积）
		pipe.Expire(Ctx, key, window*2)

		// 5. 执行所有命令
		_, err := pipe.Exec(Ctx)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
			return
		}
		if countCmd.Val() > int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "请求过于频繁，请稍后再试"})
			return
		}
		c.Next()
	}
}
