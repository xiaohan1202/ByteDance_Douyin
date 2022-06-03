package model

import (
	"gorm.io/gorm"
	"time"
)

type VideoDao struct {
}

//数据表Video结构
type Video struct {
	gorm.Model
	AuthorID      int64  `gorm:"not null; index:idx_author_id" json:"author_id"`
	Title         string `gorm:"not null" json:"title"`
	PlayURL       string `gorm:"not null" json:"play_url"`
	CoverURL      string `gorm:"not null" json:"coverurl"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
}

func NewVideoDao() *VideoDao {
	return new(VideoDao)
}

func (v *VideoDao) AddVideo(video *Video) error {
	return DB.Create(video).Error
}

func (vd *VideoDao) QueryVideosByUserId(id int64) ([]Video, error) {
	videoList := make([]Video, 0, 30)
	err := DB.Model(&Video{}).Where("author_id = ?", id).Order("created_at ASC").Limit(30).Find(&videoList).Error
	return videoList, err
}

func (vd *VideoDao) QueryVideosById(id int64) (*Video, error) {
	var video Video
	err := DB.Model(&Video{}).Where("id = ?", id).Find(&video).Error
	return &video, err
}

func (vd *VideoDao) QueryVideoByLatestTime(latestTime time.Time) ([]Video, error) {
	var videoList []Video
	err := DB.Model(&Video{}).Where("created_at<=?", latestTime).Order("created_at ASC").Limit(30).Find(&videoList).Error
	return videoList, err
}
