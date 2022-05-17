package controller

import "github.com/yhmain/simple-tiktok/model"

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
// 准备删除的文件
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		NickName:      "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
	"admin123456": {
		Id:            2,
		NickName:      "admin",
		FollowCount:   13,
		FollowerCount: 51,
		IsFollow:      true,
	},
}
