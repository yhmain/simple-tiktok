package dao

import (
	"github.com/yhmain/simple-tiktok/model"
)

//MyDB是dao包下的全局变量
//加载全部的 User 到 users
func SelectAllUsers() []model.User {
	var users []model.User
	MyDB.Find(&users)
	return users
}
