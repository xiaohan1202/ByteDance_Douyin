package controller

import (
	"ByteDance_Douyin/service"
	"ByteDance_Douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//用户注册返回的响应
type UserRegisterResponse struct {
	Response
	Data *service.UserRegisterData
}

//用户登录返回的响应
type UserLoginResponse struct {
	Response
	Data *service.UserLoginData
}

//用户查询返回的响应
type UserInfoResponse struct {
	Response
	User *service.UserInfo `json:"user"`
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !utils.CheckNameValid(username) {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "用户名不符合规范",
		})
		return
	}
	if !utils.CheckPasswordValid(password) {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "密码不符合规范",
		})
		return
	}
	userRegisterData, status := service.NewRegisterService(username, password).Register()
	switch status {
	case 0:
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{
				StatusCode: int32(status),
				StatusMsg:  "注册成功",
			},
			Data: userRegisterData,
		})
	case 1:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "系统内部错误",
		})
	case 2:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "用户名已存在",
		})
	case 3:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "注册失败",
		})
	case 4:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "token颁发失败",
		})
	}

}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !utils.CheckNameValid(username) {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "用户名不符合规范",
		})
		return
	}
	if !utils.CheckPasswordValid(password) {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "密码不符合规范",
		})
		return
	}
	userLoginData, status := service.NewLoginService(username, password).Login()
	switch status {
	case 0:
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: int32(status),
				StatusMsg:  "登录成功",
			},
			Data: userLoginData,
		})
	case 1:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "用户名不存在",
		})
	case 2:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "密码错误",
		})
	case 3:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "token颁发失败",
		})
	}

}

func UserInfo(c *gin.Context) {
	//取出待查询用户ID
	rawId := c.Query("user_id")
	token := c.Query("token")
	userId, _ := strconv.ParseInt(rawId, 10, 64)
	user, status := service.NewUserInfoService(token, userId).QueryUserInfo()
	switch status {
	case 0:
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{
				StatusCode: int32(status),
				StatusMsg:  "查询成功",
			},
			User: user,
		})
	case 1:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "token解析失败",
		})
	case 2:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "查询失败",
		})
	}
}
