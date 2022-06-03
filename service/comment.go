package service

import (
	"ByteDance_Douyin/model"
)

type CommentActionService struct {
	UserId      int64
	Token       string
	VideoId     int64
	ActionType  int64
	CommentText string
	CommentId   int64
}

type CommentListService struct {
	Token   string
	VideoId int64
}

type CommentWithAuthor struct {
	Id         int64     `json:"id"`
	User       *UserInfo `json:"user"`
	Content    string    `json:"content"`
	CreateDate string    `json:"create_date"`
}

type CommentActionData struct {
	CommentWithAuthor `json:"comment,omitempty"`
}

type CommentListData struct {
	CommentList []CommentWithAuthor `json:"comment_list,omitempty"`
}

func NewCommentActionService(userId, videoId, commentId, actionType int64, commentText, token string) *CommentActionService {
	return &CommentActionService{UserId: userId, Token: token, VideoId: videoId, ActionType: actionType, CommentText: commentText, CommentId: commentId}
}

func NewCommentListService(token string, videoId int64) *CommentListService {
	return &CommentListService{Token: token, VideoId: videoId}
}

func (c *CommentActionService) CommentAciton() (*CommentActionData, int) {
	status := 0
	//userId, _ := strconv.ParseInt(c.UserId, 10, 64)

	var commentActionData *CommentActionData
	commentActionData = nil
	switch c.ActionType {
	case 1:
		var userInfo *UserInfo
		var comment *model.Comment
		comment, status = model.NewCommentDao().AddComment(c.UserId, c.VideoId, c.CommentText)
		userInfo, status = NewUserInfoService(c.Token, c.UserId).QueryUserInfo()
		commentActionData = &CommentActionData{CommentWithAuthor: CommentWithAuthor{
			Id:         int64(comment.ID),
			User:       userInfo,
			Content:    c.CommentText,
			CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05")[5:10],
		}}
	case 2:
		status = model.NewCommentDao().DeleteComment(c.CommentId)
	}
	return commentActionData, status
}

func (c *CommentListService) CommentList() (*CommentListData, int) {
	commentList, status := model.NewCommentDao().QueryCommentListByVideoId(c.VideoId)
	var commentListData []CommentWithAuthor
	for i := range commentList {
		var commentWithAuthor CommentWithAuthor
		var userInfo *UserInfo
		userInfo, status = NewUserInfoService(c.Token, commentList[i].UserID).QueryUserInfo()
		commentWithAuthor = CommentWithAuthor{
			Id:         int64(commentList[i].ID),
			User:       userInfo,
			Content:    commentList[i].Content,
			CreateDate: commentList[i].CreatedAt.Format("2006-01-02 15:04:05")[5:10],
		}
		commentListData = append(commentListData, commentWithAuthor)
	}
	return &CommentListData{CommentList: commentListData}, status
}
