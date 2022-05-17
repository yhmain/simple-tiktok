package service

import (
	"github.com/yhmain/simple-tiktok/dao"
	"github.com/yhmain/simple-tiktok/model"

	idworker "github.com/gitstliu/go-id-worker"
)

//分布式id生成器
var userIDGen *idworker.IdWorker

//ID生成器的初始化
func init() {
	userIDGen = &idworker.IdWorker{}
	userIDGen.InitIdWorker(1, 1) //WORKERID位数 (用于对工作进程进行编码), 数据中心ID位数 (用于对数据中心进行编码)
}

//生成新的用户ID
//失败则返回-1
func GetNewUserID() int64 {
	id, err := userIDGen.NextId() //生成新ID
	if err != nil {
		return -1
	}
	return id
}

//按照用户名查找用户 true表示存在，false表示不存在
func SelectUserByName(name string) (model.User, bool) {
	//调用dao层获取数据
	var user = dao.SelectUserByName(name)
	if user.NickName == "" { //此时表示该用户名不存在，返回空nil
		return user, false
	}
	return user, true
}

//按照 用户名和密码查找用户
func SelectUserByNamePwd(name, pwd string) (model.User, bool) {
	//调用dao层获取数据
	var user = dao.SelectUserByNamePwd(name, pwd)
	if user.NickName == "" { //此时表示该用户名不存在，返回空nil
		return user, false
	}
	return user, true
}

//插入新用户
func InsertNewUser(user model.User) error {
	return dao.InsertNewUser(user)
}
