package service

import (
	"ByteDance_Douyin/model"
	"strconv"
	"time"
)

type FeedService struct {
	TimeStamp string
	Token     string
}

//带Author的Video结构
type VideoDisplay struct {
	Id            int64     `json:"id,omitempty"`
	UserInfoId    int64     `json:"-"`
	Author        *UserInfo `json:"author" gorm:"-"`
	PlayUrl       string    `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

type FeedData struct {
	NextTime  int64          `json:"next_time"`
	VideoList []VideoDisplay `json:"video_list,omitempty"`
}

func NewFeedService(timeStamp, token string) *FeedService {
	return &FeedService{TimeStamp: timeStamp, Token: token}
}

func (f *FeedService) Feed() (*FeedData, int) {
	status := 0
	var latestTime time.Time
	rawTime, _ := strconv.ParseInt(f.TimeStamp, 10, 64)
	if rawTime == 0 {
		latestTime = time.Now()
	} else {
		latestTime = time.UnixMilli(rawTime)
	}
	videoList, err := model.NewVideoDao().QueryVideoByLatestTime(latestTime)
	if err != nil {
		status = 1
		return nil, status
	}
	//获得作者信息
	videoDisplayList := make([]VideoDisplay, 0, 30)
	for i := range videoList {
		var videoDisplay VideoDisplay
		videoDisplay.Title = videoList[i].Title
		videoDisplay.Id = int64(videoList[i].ID)
		videoDisplay.CreatedAt = videoList[i].CreatedAt
		videoDisplay.PlayUrl = videoList[i].PlayURL
		videoDisplay.CoverUrl = videoList[i].CoverURL
		videoDisplay.Author, _ = NewUserInfoService(f.Token, videoList[i].AuthorID).QueryUserInfo()
		//videoDisplayList[i].IsFavorite
		videoDisplayList = append(videoDisplayList, videoDisplay)
	}
	nextTime := videoList[len(videoList)-1].CreatedAt.Unix()
	return &FeedData{NextTime: nextTime, VideoList: videoDisplayList}, status
}
