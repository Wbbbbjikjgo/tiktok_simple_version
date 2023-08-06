package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goTouch/TicTok_SimpleVersion/domain"
	"github.com/goTouch/TicTok_SimpleVersion/service"
	"github.com/goTouch/TicTok_SimpleVersion/util"
	"log"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
//前端接口文档中，前端有带有token，videoId，actionType(1表示点赞，2表示取消点赞)三个参数，利用好
func FavoriteAction(c *gin.Context) {
	//实现点赞三个数据是必要的，点赞用户的id（从token中拿），点赞视频的id，是否点赞

	//验证token，合法的话返回userId
	userIdInt64, err := util.VerifyTokenReturnUserIdInt64(c)
	if err != nil {
		c.JSON(http.StatusOK, domain.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		log.Println(err)
	}

	videoIdInt64, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, domain.Response{StatusCode: 1, StatusMsg: "获取视频id失败！"})
		log.Println("出现无法解析成64位整数的视频id")
		return
	}

	//actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32) //这个函数返回的就是64位的！
	actionType, err := strconv.Atoi(c.Query("action_type"))

	if err != nil {
		return
	}

	service.Favorite(videoIdInt64, userIdInt64, int32(actionType))
}

// FavoriteList all users have same favorite videos list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, domain.VideoListResponse{
		Response: domain.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
