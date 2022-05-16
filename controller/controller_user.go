package controller

import (
	"github.com/yhmain/simple-tiktok/dao"
	"github.com/yhmain/simple-tiktok/model"
)

func SelectAllUsers() []model.User {
	return dao.SelectAllUsers()
}
