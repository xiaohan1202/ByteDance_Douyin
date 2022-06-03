package model

import "gorm.io/gorm"

// 评论信息表
// idx_video_id: 查找视频ID对应的所有评论
type Comment struct {
	gorm.Model
	UserID  int64  `gorm:"not null" json:"user_id"`
	VideoID int64  `gorm:"not null; index:idx_video_id" json:"video_id"`
	Content string `gorm:"not null" json:"content"`
}

type CommentDao struct {
}

func NewCommentDao() *CommentDao {
	return &CommentDao{}
}

func (c *CommentDao) AddComment(userId, videoId int64, content string) (*Comment, int) {
	comment := &Comment{
		UserID:  userId,
		VideoID: videoId,
		Content: content,
	}
	if err := DB.Create(comment).Error; err != nil {
		return nil, 1
	}
	if DB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + 1")).Error != nil {
		return nil, 1
	}
	return comment, 0
}

func (c *CommentDao) DeleteComment(commentId int64) int {
	var comment Comment
	if err := DB.Model(&Comment{}).Where("id = ? AND deleted_at IS NULL", commentId).First(&comment).Error; err != nil {
		return 1
	}
	if DB.Model(&Video{}).Where("id = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count - 1")).Error != nil {
		return 2
	}
	if err := DB.Model(&Comment{}).Where("id = ? AND deleted_at IS NULL", commentId).Delete(&Comment{}).Error; err != nil {
		return 2
	}

	return 0
}

func (c *CommentDao) QueryCommentListByVideoId(videoId int64) ([]Comment, int) {
	var commentList []Comment
	if err := DB.Model(&Comment{}).Where("video_id = ? AND deleted_at IS NULL", videoId).Order("created_at ASC").Find(&commentList).Error; err != nil {
		return nil, 1
	}
	return commentList, 0
}
