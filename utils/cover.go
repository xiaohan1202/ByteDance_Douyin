package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GetCover(videoPath, coverPath string, frameNum int) (string, error) {
	coverName := ""
	videoPath = "./public/video/b0e820c4addea2517b7cf402c82f4378.mp4"
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return coverName, err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		//log.Fatal("生成缩略图失败：", err)
		return coverName, err
	}

	err = imaging.Save(img, coverPath+".jpeg")
	if err != nil {
		//log.Fatal("生成缩略图失败：", err)
		return coverName, err
	}

	// 成功则返回生成的缩略图名
	names := strings.Split(coverPath, "/")
	coverName = names[len(names)-1] + ".jpeg"
	return coverName, nil
}
