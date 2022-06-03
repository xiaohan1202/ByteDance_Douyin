package service

import (
	"ByteDance_Douyin/model"
	"ByteDance_Douyin/utils"
	"strconv"
)

type FavoriteService struct {
	UserId     string
	Token      string
	VideoId    string
	ActionType int64
}

type FavoriteListService struct {
	UserId string
	Token  string
}

type FavoriteListData struct {
	VideoList []VideoDisplay `json:"video_list,omitempty"`
}

func NewFavoriteService(userId, token, videoId string, actionType int64) *FavoriteService {
	return &FavoriteService{UserId: userId, Token: token, VideoId: videoId, ActionType: actionType}
}

func NewFavoriteListService(userId, token string) *FavoriteListService {
	return &FavoriteListService{UserId: userId, Token: token}
}

func (f *FavoriteService) Favorite() int {
	status := 0
	userId, err := utils.ParseToken(f.Token)
	if err != nil {
		status = 3
		return status
	}
	videoId, _ := strconv.ParseInt(f.VideoId, 10, 64)
	switch f.ActionType {
	case 1:
		status = model.NewFavoriteDao().AddFavorite(userId, videoId)
	case 2:
		status = model.NewFavoriteDao().DeleteFavorite(userId, videoId)
	}
	return status
}

func (f *FavoriteListService) FavoriteList() (*FavoriteListData, int) {
	curUser, err := utils.ParseToken(f.Token)
	if err != nil {
		return nil, 2
	}
	queryUser, _ := strconv.ParseInt(f.UserId, 10, 64)
	videoIdList, status := model.NewFavoriteDao().QueryFavoriteById(queryUser)
	if status != 0 {
		return nil, status
	}
	videoDisplayList := make([]VideoDisplay, 0, 30)
	videoDao := model.NewVideoDao()
	favoriteDao := model.NewFavoriteDao()
	for i := range videoIdList {
		var videoDisplay VideoDisplay
		video, err := videoDao.QueryVideosById(videoIdList[i])
		if err != nil {
			status = 1
			return nil, status
		}
		videoDisplay = VideoDisplay{
			Id:            int64(video.ID),
			Title:         video.Title,
			CreatedAt:     video.CreatedAt,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
		}
		videoDisplay.Author, _ = NewUserInfoService(f.Token, video.AuthorID).QueryUserInfo()
		videoDisplay.IsFavorite = favoriteDao.IsFavorite(curUser, videoIdList[i])
		videoDisplayList = append(videoDisplayList, videoDisplay)
	}
	return &FavoriteListData{VideoList: videoDisplayList}, status
}
