package controller

import (
	"ByteDance_Douyin/model"
	"ByteDance_Douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

//发布视频响应
type PublishResponse struct {
	Response
}

//获取登录用户的视频发布列表响应
type PublishListResponse struct {
	Response
	VideoList []service.VideoDisplay `json:"video_list,omitempty"`
}

func Publish(c *gin.Context) {
	// 从POST请求 读取参数
	title := c.PostForm("title")
	//从token取出当前用户ID
	token := c.PostForm("token")
	rawUserId := parseToken(token).Id
	userId, _ := strconv.ParseInt(rawUserId, 10, 64)
	// 读取data数据
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
	// 生成视频 封面 url链接
	playURL := "http://10.116.89:8080/static/video/" + fileName
	coverURL := "http://10.116.89:8080/static/cover/" + fileName
	video := &model.Video{
		AuthorID: userId,
		Title:    title,
		PlayURL:  playURL,
		CoverURL: coverURL,
	}
	//调取Service层函数
	publishService := service.PublishService{
		Video: video,
	}
	if err := publishService.Publish(); err != nil {
		c.JSON(http.StatusOK, PublishResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

func PublishList(c *gin.Context) {
	//获取用户id
	rawId := c.Query("user_id")
	userId, _ := strconv.ParseInt(rawId, 10, 64)
	fmt.Println(userId)
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, PublishListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "token错误"},
		})
		return
	}
	publishService := service.PublishService{
		Id: userId,
	}
	videoList, err := publishService.PublishList()
	if err != nil {
		c.JSON(http.StatusOK, PublishListResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, PublishListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
	})
}
