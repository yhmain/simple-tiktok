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

//查找用户名是否已存在（按照规定，用户名是唯一的）
//nil空则表示不存在
func SelectUserByName(name string) model.User {
	var user model.User
	MyDB.Where("NickName=?", name).Find(&user)
	return user
}

//查找用户名、密码是否已存在（按照规定，用户名是唯一的）
//nil空则表示不存在
func SelectUserByNamePwd(name, pwd string) model.User {
	var user model.User
	MyDB.Where("NickName=? AND UserPwd=?", name, pwd).Find(&user)
	return user
}

//插入新用户
//返回是否出错，为nil则表示插入成功
func InsertNewUser(user *model.User) error {
	result := MyDB.Create(user)
	return result.Error
}
