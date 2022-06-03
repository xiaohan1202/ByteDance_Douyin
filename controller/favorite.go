package controller

import (
	"ByteDance_Douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FavoriteResponse struct {
	Response
}

type FavoriteListResponse struct {
	Response
	Data *service.FavoriteListData
}

func FavoriteAction(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	status := service.NewFavoriteService(userId, token, videoId, actionType).Favorite()
	switch actionType {
	case 1:
		switch status {
		case 0:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "点赞成功"})
		case 1:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "点赞已经存在"})
		case 2:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "点赞失败"})
		case 3:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "token解析失败"})

		}
	case 2:
		switch status {
		case 0:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "取消点赞成功"})
		case 1:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "点赞不存在"})
		case 2:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "取消点赞失败"})
		case 3:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "token解析失败"})
		}
	}
}

func FavoriteList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	favoriteListData, status := service.NewFavoriteListService(userId, token).FavoriteList()
	switch status {
	case 0:
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "查询成功"},
			Data:     favoriteListData,
		})
	case 1:
		c.JSON(http.StatusOK,
			Response{StatusCode: int32(status), StatusMsg: "查询失败"})
	case 2:
		c.JSON(http.StatusOK,
			Response{StatusCode: int32(status), StatusMsg: "token解析失败"})
	}
}
