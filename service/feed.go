package service

import (
	"ByteDance_Douyin/model"
	"ByteDance_Douyin/utils"
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
	curUser, _ := utils.ParseToken(f.Token)
	videoList, err := model.NewVideoDao().QueryVideoByLatestTime(latestTime)
	if err != nil {
		status = 1
		return nil, status
	}
	//获得作者信息
	videoDisplayList := make([]VideoDisplay, 0, 30)
	favoriteDao := model.NewFavoriteDao()
	for i := range videoList {
		var videoDisplay VideoDisplay
		videoDisplay = VideoDisplay{
			Id:            int64(videoList[i].ID),
			Title:         videoList[i].Title,
			CreatedAt:     videoList[i].CreatedAt,
			PlayUrl:       videoList[i].PlayURL,
			CoverUrl:      videoList[i].CoverURL,
			FavoriteCount: videoList[i].FavoriteCount,
			CommentCount:  videoList[i].CommentCount,
		}
		videoDisplay.Author, _ = NewUserInfoService(f.Token, videoList[i].AuthorID).QueryUserInfo()
		videoDisplay.IsFavorite = favoriteDao.IsFavorite(curUser, videoDisplay.Id)
		videoDisplayList = append(videoDisplayList, videoDisplay)
	}
	nextTime := videoList[len(videoList)-1].CreatedAt.Unix()
	return &FeedData{NextTime: nextTime, VideoList: videoDisplayList}, status
}
