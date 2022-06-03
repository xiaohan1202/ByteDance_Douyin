package controller

import (
	"ByteDance_Douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

//发布视频响应
type PublishResponse struct {
	Response
}

//获取登录用户的视频发布列表响应
type PublishListResponse struct {
	Response
	Data *service.PublishListData
}

func Publish(c *gin.Context) {
	// 从POST请求 读取参数
	title := c.PostForm("title")
	////从token取出当前用户ID
	token := c.PostForm("token")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, PublishResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	// 存放到本地
	fileName := filepath.Base(data.Filename) // filename 应该 与 title 对应
	saveFile := filepath.Join("./public/video", fileName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, PublishResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//调取service层
	status := service.NewPublishService(token, title, data).Publish()
	switch status {
	case 0:
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "上传成功"})
	case 1:
		c.JSON(http.StatusOK, PublishResponse{
			Response: Response{StatusCode: 1, StatusMsg: "解析token失败"},
		})
	case 2:
		c.JSON(http.StatusOK, PublishResponse{
			Response: Response{StatusCode: 1, StatusMsg: "视频上传失败"},
		})
	}
}

func PublishList(c *gin.Context) {
	//获取用户id
	userId := c.Query("user_id")
	token := c.Query("token")
	publishListData, status := service.NewPublishListService(userId, token).PublishList()
	switch status {
	case 0:
		c.JSON(http.StatusOK, PublishListResponse{
			Response: Response{StatusCode: 0},
			Data:     publishListData,
		})
	case 1:
		c.JSON(http.StatusOK, PublishListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "token解析失败"},
		})
	case 2:
		c.JSON(http.StatusOK, PublishListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户获取失败"},
		})
	case 3:
		c.JSON(http.StatusOK, PublishListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "视频获取失败"},
		})
	}

}
