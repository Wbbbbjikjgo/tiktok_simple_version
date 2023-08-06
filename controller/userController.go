package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goTouch/TicTok_SimpleVersion/domain"
	"log"
	"net/http"
	"sync/atomic"
)

// 注：用户id的名字统一为：userIdInt64,在token验证，将用户id写入gin上下文的时候，还有其他时候，用户id的key都设置为：userIdInt64。如果不是int64类型另当别论
//该部分其他代码是测试demo，编辑的时候可以删除
func getUserIdByGinContext(c *gin.Context) (userIdInt64 int64, ok bool) {

	//	//将id写入gin上下文中 相对应
	//	c.Set("id", id)
	// c.Get(key) 方法来获取之前存储的值时，返回的数据类型将始终是 interface{} 类型，要转换一下
	userIdObject, exists := c.Get("userIdInt64")
	if !exists {
		return
	}
	userIdInt64, ok = userIdObject.(int64) //这个userIdInt64与 ok在返回值上直接对应，因此无需声明。
	if !ok {
		log.Println("出现无法解析成64位整数的token")
		return
	}
	return
}

//下面的代码是demo的，可以选择删除或者保留
//*********************************************************************************************//

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]domain.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	domain.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	domain.Response
	User domain.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: domain.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := domain.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: domain.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: domain.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: domain.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: domain.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: domain.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
