package controller

// omitempty: 如果信息不存在，则转化成json时不包含默认值

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

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

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"column:Vid;primary_key"`           //视频主键
	Detail        string `json:"video_detail,omitempty" gorm:"column:VideoDetail"`     //视频简介
	PlayUrl       string `json:"play_url,omitempty" gorm:"column:PlayUrl"`             //视频播放地址
	CoverUrl      string `json:"cover_url,omitempty" gorm:"column:CoverUrl"`           //视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty" gorm:"column:FavoriteCount"` //视频点赞总数
	CommentCount  int64  `json:"comment_count,omitempty" gorm:"column:CommentCount"`   //视频评论总数
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"column:IsFavorite"`       //true:已点赞，false:未点赞
	CreatedTime   int64  `json:"created_time,omitempty" gorm:"column:CreatedTime"`     //视频创建时间，时间戳形式
	User          User   `json:"user" gorm:"ForeignKey:UserID"`                        //视频作者
	UserID        int64  `json:"user_id" gorm:"column:Uid"`                            //外键：视频作者的ID
}

//结构体Video对应数据库中videos表
func (e *Video) TableName() string {
	return "videos"
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`          //评论id
	User       User   `json:"user"`                  //创建该评论的用户id
	Content    string `json:"content,omitempty"`     //评论内容
	CreateDate string `json:"create_date,omitempty"` //评论时间
}

//结构体Comment对应数据库中comments表
func (e *Comment) TableName() string {
	return "comments"
}
