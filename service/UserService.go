package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/goTouch/TicTok_SimpleVersion/dao"
	"log"
	"strconv"
	"time"
)

// SaveTokenToRedis 使用雪花算法生成int64的id作为token.并以token为key，value是userId ，存到redis中去。返回token
// 注：CreateToken使用这个替代
func SaveTokenToRedis(userIdInt64 int64) (token int64) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Println("雪花算法生成id错误")
		log.Println(err)
	}
	token = node.Generate().Int64()
	//存入redis
	dao.RedisClient.Set(context.Background(), strconv.FormatInt(token, 10), userIdInt64, 12*time.Hour)
	return
}
