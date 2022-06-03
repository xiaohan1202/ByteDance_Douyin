package service

import (
	"ByteDance_Douyin/model"
	"ByteDance_Douyin/utils"
)

type RelationActionService struct {
	Token    string
	ToUserId int64
}

type RelationListService struct {
	UserId int64
	Token  string
}

type RelationListData struct {
	UserList []UserInfo
}

func NewRelationActionService(toUserId int64, token string) *RelationActionService {
	return &RelationActionService{Token: token, ToUserId: toUserId}
}

func NewRelationListService(token string, userId int64) *RelationListService {
	return &RelationListService{Token: token, UserId: userId}
}

func (r *RelationActionService) Follow() int {
	userId, err := utils.ParseToken(r.Token)
	if err != nil {
		return 3
	}
	return model.NewFollowDao().AddFollow(userId, r.ToUserId)
}

func (r *RelationActionService) UnFollow() int {
	userId, err := utils.ParseToken(r.Token)
	if err != nil {
		return 3
	}
	return model.NewFollowDao().DeleteFollow(userId, r.ToUserId)
}

func (r *RelationListService) FollowList() (*RelationListData, int) {
	userList, status := model.NewFollowDao().QueryFollowById(r.UserId)
	//fmt.Printf("关注列表：%+v", userList)
	if status != 0 {
		return nil, status
	}
	var userInfoList []UserInfo
	for i := range userList {
		var userInfo *UserInfo
		userInfo, status = NewUserInfoService(r.Token, userList[i]).QueryUserInfo()
		userInfoList = append(userInfoList, *userInfo)
	}
	return &RelationListData{UserList: userInfoList}, status
}

func (r *RelationListService) FollowerList() (*RelationListData, int) {
	userList, status := model.NewFollowDao().QueryFollowerById(r.UserId)
	if status != 0 {
		return nil, status
	}
	var userInfoList []UserInfo
	for i := range userList {
		var userInfo *UserInfo
		userInfo, status = NewUserInfoService(r.Token, userList[i]).QueryUserInfo()
		userInfoList = append(userInfoList, *userInfo)
	}
	return &RelationListData{UserList: userInfoList}, status
}
