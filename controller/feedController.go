package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goTouch/TicTok_SimpleVersion/domain"
	"github.com/goTouch/TicTok_SimpleVersion/service"
	"net/http"
	"strconv"
)

type FeedResponse struct {
	domain.Response
	VideoList []domain.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	//根据接口文档，前端传来的request中有token和latest_time， 这里一个用于存当前用户id，一个存下次视频时间戳
	//tokenReq := c.Query("token")
	latestTimeReq := c.Query("latest_time")                         //字符串类型
	latestTimeInt64, err := strconv.ParseInt(latestTimeReq, 10, 64) //转为时间戳
	if err != nil {
		c.JSON(http.StatusOK, domain.Response{StatusCode: 1, StatusMsg: "时间戳格式错误"}) //定义1为时间戳格式错误
		return
	}
	id := int64(1) // TODO 用户id暂时这样
	videoList, nextTimeInt64 := service.FeedService(id, latestTimeInt64)
	//to be continue....
}
