package main

import (
	"ByteDance_Douyin/model"
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	model.InitMySQL()
	userList, _ := model.NewFollowDao().QueryFollowById(4)
	fmt.Printf("%+v", userList)
}
