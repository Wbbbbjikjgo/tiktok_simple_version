package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goTouch/TicTok_SimpleVersion/domain"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
//前端接口文档中，有token，videoId，actionType(是否点赞)三个参数，利用好
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, domain.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, domain.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite videos list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: domain.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
