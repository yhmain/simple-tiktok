package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yhmain/simple-tiktok/model"
)

//该测试用例期待结果为空
func TestSelectUserByName100(t *testing.T) {
	user := SelectUserByName("hahaha")
	expectedResult := model.User{Id: 0, NickName: "", Password: "", FollowCount: 0, FollowerCount: 0, IsFollow: false}
	assert.Equal(t, expectedResult, user)
}

//该测试用例期待结果为 id为2的admin用户
func TestSelectUserByName200(t *testing.T) {
	user := SelectUserByName("admin")
	expectedResult := model.User{Id: 2, NickName: "admin", Password: "123456", FollowCount: 20, FollowerCount: 20, IsFollow: false}
	assert.Equal(t, expectedResult, user)
}

//测试按照用户名和密码搜索用户
//正确查找到
func TestSelectUserByNamePwd100(t *testing.T) {
	user := SelectUserByNamePwd("admin", "123456")
	expectedResult := model.User{Id: 2, NickName: "admin", Password: "123456", FollowCount: 20, FollowerCount: 20, IsFollow: false}
	assert.Equal(t, expectedResult, user)
}

//测试按照用户名和密码搜索用户
//查找失败！
func TestSelectUserByNamePwd200(t *testing.T) {
	user := SelectUserByNamePwd("zhanglei", "123456")
	expectedResult := model.User{Id: 0, NickName: "", Password: "", FollowCount: 0, FollowerCount: 0, IsFollow: false}
	assert.Equal(t, expectedResult, user)
}

//插入新用户
func TestInsertNewUser(t *testing.T) {
	err := InsertNewUser(model.User{Id: 1000, NickName: "user1000", Password: "7777909"})
	var expectedResult error
	assert.Equal(t, expectedResult, err)
}
