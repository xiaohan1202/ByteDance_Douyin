package model

import (
	"errors"
)

var (
	ErrUserExists    = errors.New("error user exists")
	ErrUserNotExists = errors.New("error user not exists")
)

type UserDao struct {
}

type UserInfoDao struct {
}

// 用户表
type User struct {
	Id            int64  `json:"id,omitempty"`
	Username      string `json:"name,omitempty"`
	Password      []byte `json:"password"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	Salt          []byte `gorm:"not null; type:varbinary(32)" json:"-"`
}

func NewUserDao() *UserDao {
	return new(UserDao)
}

func NewUserInfoDao() *UserInfoDao {
	return new(UserInfoDao)
}

func (u *UserDao) QueryUserById(id int64) (*User, error) {
	user := &User{}
	err := DB.First(user, "Id = ?", id).Error
	return user, err
}

func (u *UserDao) AddUser(user *User) int {
	var count int64
	status := 0
	if err := DB.Table("users").Where("username =?", user.Username).Count(&count).Error; err != nil {
		status = 1
	}
	if count != 0 {
		status = 2
	}
	if err := DB.Create(user).Error; err != nil {
		status = 3
	}
	return status
}

func (u *UserDao) QueryUserByName(name string) (*User, error) {
	var user User
	err := DB.Where(&User{Username: name}, "username").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
