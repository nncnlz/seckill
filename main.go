package main

import (
	"github.com/gin-gonic/gin"
	"seckillProject/handler"
	"seckillProject/redis"
)

func main() {
	redis.InitRedis()
	r := gin.Default()
	r.GET("/seckill", handler.SeckillHandler)

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
