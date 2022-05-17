package model

//结构体名称+ID 即可设置外键
type User struct {
	Id            int64  `json:"id,omitempty" gorm:"column:Uid;primary_key"`           //用户ID
	NickName      string `json:"name,omitempty" gorm:"column:NickName"`                //昵称
	Password      string `json:"pwd,omitempty" gorm:"column:UserPwd"`                  //密码
	FollowCount   int    `json:"follow_count,omitempty" gorm:"column:FollowCount"`     //关注数
	FollowerCount int    `json:"follower_count,omitempty" gorm:"column:FollowerCount"` //粉丝数
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"column:IsFollow"`           //true:已关注，false:未关注
	CommentCount  int    `json:"comment_count,omitempty" gorm:"column:CommentCount"`   //评论数目
	IsFocused     bool   `json:"is_favorite,omitempty" gorm:"column:IsFavorite"`       //是否喜欢
	LikeCount     bool   `json:"favorite_count,omitempty" gorm:"column:FavoriteCount"` //喜欢的人数
}

//结构体User对应数据库中user表
func (e *User) TableName() string {
	return "users"
}
