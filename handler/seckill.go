package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seckillProject/redis"
)

const (
	stockKey = "seckill:stock"
	userKey  = "seckill:users"
)

var luaScript = `
if redis.call("get", KEYS[1]) and tonumber(redis.call("get", KEYS[1])) > 0 then
    redis.call("decr", KEYS[1])
    redis.call("sadd", KEYS[2], ARGV[1])
    return 1
else
    return 0
end
`

func SeckillHandler(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	res, err := redis.Rdb.Eval(redis.Ctx, luaScript, []string{stockKey, userKey}, userID).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res.(int64) == 1 {
		c.JSON(http.StatusOK, gin.H{"message": "秒杀成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "秒杀失败，可能已抢完或重复"})
	}
}
