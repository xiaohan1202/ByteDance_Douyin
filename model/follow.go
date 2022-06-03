package model

import (
	"gorm.io/gorm"
)

// 关注用户信息表
// idx_author_id: 搜索用户所有投稿视频
type Follow struct {
	gorm.Model
	FollowerID int64 `gorm:"not null; index:idx_follower" json:"follower_id"`
	FolloweeID int64 `gorm:"not null; index:idx_followee" json:"followee_id"`
}

type FollowDao struct {
}

func NewFollowDao() *FollowDao {
	return new(FollowDao)
}

func (f *FollowDao) AddFollow(followerId, followeeId int64) int {
	follow := Follow{
		FollowerID: followerId,
		FolloweeID: followeeId,
	}
	if f.IsFollow(followerId, followeeId) {
		return 1
	}
	if err := DB.Create(&follow).Error; err != nil {
		return 2
	}
	if DB.Model(&User{}).Where("id = ?", followerId).Update("follow_count", gorm.Expr("follow_count + 1")).Error != nil {
		return 2
	}
	if DB.Model(&User{}).Where("id = ?", followeeId).Update("follower_count", gorm.Expr("follower_count + 1")).Error != nil {
		return 2
	}
	return 0
}

func (f *FollowDao) DeleteFollow(followerId, followeeId int64) int {
	if !f.IsFollow(followerId, followeeId) {
		return 1
	}
	if err := DB.Model(&Follow{}).Where("follower_id = ? AND followee_id = ? AND deleted_at IS NULL", followerId, followeeId).Delete(&Follow{}).Error; err != nil {
		return 2
	}
	if DB.Model(&User{}).Where("id = ?", followerId).Update("follow_count", gorm.Expr("follow_count - 1")).Error != nil {
		return 2
	}
	if DB.Model(&User{}).Where("id = ?", followeeId).Update("follower_count", gorm.Expr("follower_count - 1")).Error != nil {
		return 2
	}
	return 0
}

func (f *FollowDao) QueryFollowById(userId int64) ([]int64, int) {
	var userList []int64
	if err := DB.Model(&Follow{}).Select("followee_id").Where("follower_id = ? AND deleted_at IS NULL", userId).Find(&userList).Error; err != nil {
		return nil, 1
	}
	return userList, 0
}

func (f *FollowDao) QueryFollowerById(userId int64) ([]int64, int) {
	var userList []int64
	if err := DB.Model(&Follow{}).Select("follower_id").Where("followee_id = ? AND deleted_at IS NULL", userId).Find(&userList).Error; err != nil {
		return nil, 1
	}
	return userList, 0
}

func (f *FollowDao) IsFollow(followerId, followeeId int64) bool {
	return DB.Model(Follow{}).Where("follower_id = ? AND followee_id = ? AND deleted_at IS NULL", followerId, followeeId).First(&Follow{}).Error == nil
}
