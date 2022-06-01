package service

import (
	"ByteDance_Douyin/model"
	"ByteDance_Douyin/utils"
	"math/rand"
	"reflect"
)

type UserRegisterService struct {
	Username string
	Password string
}

type UserRegisterData struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserLoginService struct {
	Username string
	Password string
}

type UserLoginData struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfoService struct {
	token  string
	userId int64
}

//返回的用户信息表
type UserInfo struct {
	Id            int64  `json:"id,omitempty"`
	Username      string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func NewRegisterService(username, password string) *UserRegisterService {
	return &UserRegisterService{Username: username, Password: password}
}

func NewLoginService(username, password string) *UserLoginService {
	return &UserLoginService{Username: username, Password: password}
}

func NewUserInfoService(token string, id int64) *UserInfoService {
	return &UserInfoService{token: token, userId: id}
}

func (u *UserRegisterService) Register() (*UserRegisterData, int) {
	status := 0
	user := model.User{}
	salt := make([]byte, 32)
	for i := range salt {
		salt[i] = byte(rand.Intn(256))
	}
	user.Salt = salt
	user.Password = utils.Encrypt(u.Password, salt)
	user.Username = u.Username
	status = model.NewUserDao().AddUser(&user)
	if status != 0 {
		return nil, status
	}
	token, err := utils.SignToken(user)
	if err != nil {
		status = 4
		return nil, status
	}
	return &UserRegisterData{UserId: user.Id, Token: token}, status
}

func (u *UserLoginService) Login() (*UserLoginData, int) {
	status := 0
	user, _ := model.NewUserDao().QueryUserByName(u.Username)
	if user == nil {
		status = 1
		return nil, status
	}
	upassword := utils.Encrypt(u.Password, user.Salt)
	if !reflect.DeepEqual(upassword, user.Password) {
		status = 2
		return nil, status
	}
	token, err := utils.SignToken(*user)
	if err != nil {
		status = 3
		return nil, status
	}
	return &UserLoginData{UserId: user.Id, Token: token}, status
}

func (u *UserInfoService) QueryUserInfo() (*UserInfo, int) {
	status := 0
	//解析token
	//claims, err := utils.ParseToken(u.token)
	//if err != nil {
	//	status = 1
	//	return nil, status
	//}
	//currentId, _ := strconv.ParseInt(claims.Id, 10, 64)
	user, err := model.NewUserDao().QueryUserById(u.userId)
	if err != nil {
		status = 2
		return nil, status
	}
	userInfo := &UserInfo{
		Id:            user.Id,
		Username:      user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
	}
	//查询是否有关注关系
	//user.IsFollow = model.isFollow(u.CurrentUser, u.QueryUser)
	return userInfo, status
}
