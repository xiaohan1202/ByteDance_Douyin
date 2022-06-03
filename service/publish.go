package service

import (
	"ByteDance_Douyin/config"
	"ByteDance_Douyin/model"
	"ByteDance_Douyin/utils"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
)

type PublishService struct {
	Token      string
	Title      string
	Fileheader *multipart.FileHeader
}

type PublishListService struct {
	UserId string
	Token  string
}

type PublishListData struct {
	VideoList []VideoDisplay `json:"video_list,omitempty"`
}

func NewPublishService(token, title string, fileHeader *multipart.FileHeader) *PublishService {
	return &PublishService{Token: token, Title: title, Fileheader: fileHeader}
}

func NewPublishListService(userId, token string) *PublishListService {
	return &PublishListService{UserId: userId, Token: token}
}

func (p *PublishService) Publish() int {
	status := 0
	userId, err := utils.ParseToken(p.Token)
	if err != nil {
		status = 1
		return status
	}
	fmt.Println("收到的token:", p.Token)
	//生成封面
	videoName := filepath.Base(p.Fileheader.Filename) // filename 应该 与 title 对应
	videoPath := filepath.Join("./public/video", videoName)
	coverPath := "./public/cover/" + videoName
	coverName, _ := utils.GetCover(videoPath, coverPath, 1)

	playURL := "http://" + config.C.LocalIp.Ip + ":" + config.C.LocalIp.Port + "/static/video/" + videoName
	coverURL := "http://" + config.C.LocalIp.Ip + ":" + config.C.LocalIp.Port + "/static/cover/" + coverName
	video := &model.Video{
		AuthorID: userId,
		Title:    p.Title,
		PlayURL:  playURL,
		CoverURL: coverURL,
	}
	if err := model.NewVideoDao().AddVideo(video); err != nil {
		status = 2
		return status
	}
	return status
}

func (p *PublishListService) PublishList() (*PublishListData, int) {
	status := 0
	userId, _ := strconv.ParseInt(p.UserId, 10, 64)
	videoList, err := model.NewVideoDao().QueryVideosByUserId(userId)
	if err != nil {
		status = 3
		return nil, status
	}
	//获得作者信息
	var userInfo *UserInfo
	videoDisplayList := make([]VideoDisplay, 0, 30)
	userInfo, status = NewUserInfoService(p.Token, userId).QueryUserInfo()
	for i := range videoList {
		var videoDisplay VideoDisplay
		videoDisplay = VideoDisplay{
			Id:            int64(videoList[i].ID),
			Title:         videoList[i].Title,
			Author:        userInfo,
			CreatedAt:     videoList[i].CreatedAt,
			PlayUrl:       videoList[i].PlayURL,
			CoverUrl:      videoList[i].CoverURL,
			FavoriteCount: videoList[i].FavoriteCount,
			CommentCount:  videoList[i].CommentCount,
		}
		videoDisplayList = append(videoDisplayList, videoDisplay)
	}
	return &PublishListData{VideoList: videoDisplayList}, status
}
