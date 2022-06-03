package model

import (
	"gorm.io/gorm"
	"log"
)

// 用户点赞表
// idx_userid_videoid: 查找用户点赞列表，查找用户是否给某视频点了赞
type FavoriteDao struct {
}

type Favorite struct {
	gorm.Model
	UserID  int64 `gorm:"not null; index:idx_userid_videoid" json:"user_id"`
	VideoID int64 `gorm:"not null; index:idx_userid_videoid" json:"video_id"`
}

func NewFavoriteDao() *FavoriteDao {
	return new(FavoriteDao)
}

func (f *FavoriteDao) AddFavorite(userId, videoId int64) int {
	if f.IsFavorite(userId, videoId) {
		return 1
	}
	if err := DB.Create(&Favorite{UserID: userId, VideoID: videoId}).Error; err != nil {
		return 2
	}
	if DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error != nil {
		return 2
	}
	return 0
}

func (f *FavoriteDao) DeleteFavorite(userId, videoId int64) int {
	if !f.IsFavorite(userId, videoId) {
		return 1
	}
	if err := DB.Where("user_id = ? AND video_id = ? AND deleted_at IS NULL", userId, videoId).Delete(&Favorite{}).Error; err != nil {
		return 2
	}
	if DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error != nil {
		return 2
	}
	return 0
}

func (f *FavoriteDao) QueryFavoriteById(userId int64) ([]int64, int) {
	var videoIdList []int64
	if err := DB.Model(&Favorite{}).Select("video_id").Where("user_id = ? AND deleted_at IS NULL", userId).Find(&videoIdList).Error; err != nil {
		log.Fatal("获取视频列表失败：", err)
		return nil, 1
	}
	return videoIdList, 0
}

func (f *FavoriteDao) IsFavorite(userId, videoId int64) bool {
	return DB.Model(&Favorite{}).Where("user_id = ? AND video_id = ? AND deleted_at IS NULL", userId, videoId).First(&Favorite{}).Error == nil
}
