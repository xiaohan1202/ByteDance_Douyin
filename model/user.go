package model

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

func (u *UserDao) QueryUserById(id int64) (*User, error) {
	user := &User{}
	err := DB.First(user, "Id = ?", id).Error
	return user, err
}

func (u *UserDao) AddUser(user *User) int {
	if u.userExists(user) {
		return 1
	}
	if err := DB.Create(user).Error; err != nil {
		return 2
	}
	return 0
}

func (u *UserDao) QueryUserByName(name string) (*User, int) {
	var user User
	if err := DB.Where(&User{Username: name}, "username").First(&user).Error; err != nil {
		return nil, 1
	}
	return &user, 0
}

func (u *UserDao) userExists(user *User) bool {
	return DB.Model(&User{}).Where("username =?", user.Username).First(&User{}).Error == nil
}
