package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goTouch/TicTok_SimpleVersion/domain"
	"github.com/goTouch/TicTok_SimpleVersion/service"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	domain.Response
	VideoList []domain.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo videos list for every request
func Feed(c *gin.Context) {

	//根据接口文档，前端传来的request中有token和latest_time， 这里一个用于存当前用户id，一个存下次视频时间戳
	//tokenReq := c.Query("token")
	latestTimeReq := c.Query("latest_time")                         //字符串类型
	latestTimeInt64, err := strconv.ParseInt(latestTimeReq, 10, 64) //转为时间戳
	if err != nil {
		c.JSON(http.StatusOK, domain.Response{StatusCode: 1, StatusMsg: "时间戳格式错误"}) //定义1为错误的返回
		return
	}
	id := time.Now().UnixMilli() // TODO 暂时先生成一个id,后续和用户模块配合。用现在的时间戳，直接就是int64类型
	videoList, nextTimeInt64 := service.FeedService(id, latestTimeInt64)
	if len(videoList) != 0 { //说明查到了视频
		c.JSON(http.StatusOK, FeedResponse{
			domain.Response{StatusCode: 0, StatusMsg: "成功查询视频并返回"},
			videoList,
			nextTimeInt64,
		})
	} else {
		//注意feedResponse和response不一样，继承关系
		c.JSON(http.StatusOK,
			domain.Response{
				StatusCode: 1,
				StatusMsg:  "请求成功，但是查到0条视频！",
			},
		)
	}
}
