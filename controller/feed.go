package controller

import (
	"ByteDance_Douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedResponse struct {
	Response
	Data *service.FeedData
}

func Feed(c *gin.Context) {
	timeStamp := c.Query("latest_time")
	token := c.Query("token")
	feedData, status := service.NewFeedService(timeStamp, token).Feed()
	switch status {
	case 0:
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "查询成功"},
			Data:     feedData,
		})
	case 1:
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(status),
			StatusMsg:  "视频流拉取失败",
		})
	}
}
