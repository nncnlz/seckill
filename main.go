package main

import (
	"github.com/gin-gonic/gin"
	"seckillProject/handler"
	"seckillProject/middleware"
	"seckillProject/redis"
	"time"
)

func main() {
	redis.InitRedis()
	middleware.RedisClient = redis.Rdb
	r := gin.Default()
	r.Use(middleware.RateLimit2(5, time.Second))
	r.GET("/seckill", handler.SeckillHandler)

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
