package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"seckillProject/db"
	"seckillProject/handler"
	"seckillProject/middleware"
	"seckillProject/model"
	"seckillProject/redis"
	"time"
)

func main() {
	// 初始化
	redis.InitRedis()
	defer func() {
		if err := redis.Rdb.Close(); err != nil {
			log.Printf("Redis 关闭失败: %v", err)
		}
	}()

	openDB, err := db.OpenDB()
	if err != nil {
		log.Fatalf("打开数据库失败: %v", err)
	}
	defer func(db *gorm.DB) {
		sqlDB, err := db.DB()
		if err != nil {
			return
		}
		err = sqlDB.Close()
		if err != nil {
			return
		}
	}(openDB)

	if err := openDB.AutoMigrate(&model.Order{}); err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	}

	// 中间件 & 路由
	middleware.RedisClient = redis.Rdb
	r := gin.Default()
	r.Use(middleware.RateLimit2(5, time.Second))
	r.GET("/seckill", handler.SeckillHandler)

	// 启动服务
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("Gin 启动失败: %v", err)
	}
}
