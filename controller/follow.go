package controller

import (
	"ByteDance_Douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RelationActionResponse struct {
	Response
}

type RelationListResponse struct {
	Response
	Data *service.RelationListData
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	switch actionType {
	case 1:
		status := service.NewRelationActionService(toUserId, token).Follow()
		switch status {
		case 0:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "关注成功"})
		case 1:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "关注已存在"})
		case 2:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "关注失败"})
		case 3:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "token解析失败"})
		}
	case 2:
		status := service.NewRelationActionService(toUserId, token).UnFollow()
		switch status {
		case 0:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "取消关注成功"})
		case 1:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "关注关系不存在"})
		case 2:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "取消关注失败"})
		case 3:
			c.JSON(http.StatusOK,
				Response{StatusCode: int32(status), StatusMsg: "token解析失败"})
		}
	}
}

func FollowList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	relationListData, status := service.NewRelationListService(token, userId).FollowList()
	switch status {
	case 0:
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "查询成功"},
			Data:     relationListData,
		})
	case 4:
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "token解析失败"},
		})
	default:
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "查询失败"},
		})
	}
}

func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	relationListData, status := service.NewRelationListService(token, userId).FollowerList()
	switch status {
	case 0:
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "查询成功"},
			Data:     relationListData,
		})
	case 4:
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "token解析失败"},
		})
	default:
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "查询失败"},
		})
	}
}
