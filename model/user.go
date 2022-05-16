package model

//结构体名称+ID 即可设置外键
type User struct {
	Id            int64  `json:"id,omitempty" gorm:"column:Uid;primary_key"`           //用户ID
	Name          string `json:"name,omitempty" gorm:"column:UserName"`                //昵称
	Password      string `json:"pwd,omitempty" gorm:"column:UserPwd"`                  //密码
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"column:FollowCount"`     //关注数
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:FollowerCount"` //粉丝数
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"column:IsFollow"`           //true:已关注，false:未关注
}

//结构体User对应数据库中user表
func (e *User) TableName() string {
	return "users"
}
