package controller

// omitempty: 如果信息不存在，则转化成json时不包含默认值

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`       //视频播放地址
	CoverUrl      string `json:"cover_url,omitempty"`      //视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"` //视频点赞总数
	CommentCount  int64  `json:"comment_count,omitempty"`  //视频评论总数
	IsFavorite    bool   `json:"is_favorite,omitempty"`    //true:已点赞，false:未点赞
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`          //评论id
	User       User   `json:"user"`                  //创建该评论的用户id
	Content    string `json:"content,omitempty"`     //评论内容
	CreateDate string `json:"create_date,omitempty"` //评论时间
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`           //昵称
	FollowCount   int64  `json:"follow_count,omitempty"`   //关注数
	FollowerCount int64  `json:"follower_count,omitempty"` //粉丝数
	IsFollow      bool   `json:"is_follow,omitempty"`      //true:已关注，false:未关注
}
