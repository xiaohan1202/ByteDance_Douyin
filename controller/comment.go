package controller

import (
	"ByteDance_Douyin/service"
	"ByteDance_Douyin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentActionResponse struct {
	Response
	Data *service.CommentActionData
}

type CommentListResponse struct {
	Response
	Data *service.CommentListData
}

func CommentAction(c *gin.Context) {
	//userId := c.Query("user_id")
	token := c.Query("token")
	commentText := c.Query("comment_text")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	userId, _ := utils.ParseToken(c.Query("token"))
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	commentActionData, status := service.NewCommentActionService(userId, videoId, commentId, actionType, commentText, token).CommentAciton()
	switch actionType {
	case 1:
		switch status {
		case 0:
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: int32(status), StatusMsg: "评论成功"},
				Data:     commentActionData,
			})
		default:
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: int32(status), StatusMsg: "评论失败"},
			})
		}
	case 2:
		switch status {
		case 0:
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: int32(status), StatusMsg: "删除评论成功"},
			})
		case 1:
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: int32(status), StatusMsg: "该评论不存在"},
			})
		default:
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: int32(status), StatusMsg: "删除评论失败"},
			})
		}
	}
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	commentListData, status := service.NewCommentListService(token, videoId).CommentList()
	switch status {
	case 0:
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "获取评论列表成功"},
			Data:     commentListData,
		})
	default:
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{StatusCode: int32(status), StatusMsg: "获取评论列表失败"},
		})
	}
}
