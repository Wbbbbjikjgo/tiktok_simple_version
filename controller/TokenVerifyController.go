package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goTouch/TicTok_SimpleVersion/dao"
	"github.com/goTouch/TicTok_SimpleVersion/domain"
	"log"
	"net/http"
	"strconv"
	"time"
)

// VerifyTokenReturnUserIdInt64 token验证成功，就把其中存的id取出来，转为int64。验证失败或者没取成功，那就返回-1
func VerifyTokenReturnUserIdInt64(c *gin.Context) (int64, error) {
	// token的key是随机生成的，val是userIdInt64
	//	由于存的时候，是使用redis的set，所以存到redis之后再取出来是string类型
	token := c.Query("token")

	if token == "" {
		return -1, fmt.Errorf("token不存在")
	}

	userIdFromRedis, err := dao.RedisClient.Get(context.Background(), token).Result()
	if err != nil {
		//token在缓存中不存在，记录下本次ip访问
		clientIPAddress := c.ClientIP()
		//redis里面查一下，这个ip一分钟捏访问了多少次了
		visitTimes, _ := dao.RedisClient.Get(context.Background(), clientIPAddress).Int64()
		if visitTimes > 10 {
			c.JSON(http.StatusOK, domain.Response{
				StatusCode: 1,
				StatusMsg:  "携带非法token进行访问，次数一分钟内大于十次！",
			})
			c.Abort()
			return -1, fmt.Errorf("携带非法token进行访问，次数一分钟内大于十次")
		}
		//执行到这里说明非法的token访问次数不大于十次，使用redis记录
		dao.RedisClient.Set(context.Background(), clientIPAddress, visitTimes+1, time.Minute)
		return -1, fmt.Errorf("携带非法token进行访问")
	} else {
		//token在redis中得到验证，刷新token在redis中的时间。
		dao.RedisClient.Expire(context.Background(), token, 12*time.Hour)
		//注意，int64也是十进制整数，表示64位有符号整数！！
		userIdInt64, err := strconv.ParseInt(userIdFromRedis, 10, 64)
		if err != nil {
			log.Println("无法将token中的id解析为int64")
			return -1, fmt.Errorf("无法将token中的id解析为int64")
		}
		return userIdInt64, nil
	}

}
