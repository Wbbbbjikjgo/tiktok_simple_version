package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goTouch/TicTok_SimpleVersion/domain"
	"net/http"
	"path/filepath"
)

// 既然是发布视频，首先需要校验token，登入的问题

//下面代码为demo示例。可删除
//*******************************************************************************************************//
// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, domain.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, domain.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// filepath.Base(filename) 是一个用于提取文件路径中的文件名部分的函数。它会返回指定文件路径 filename 的最后一个元素，即文件名部分。
	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, domain.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish videos list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, domain.VideoListResponse{
		Response: domain.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
