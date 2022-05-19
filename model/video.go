package model

// 附带omitempty表示: 如果该字段不存在，则转化成json时不包含默认值

//json其实对应前端的 键
//结构体名称+ID 即可设置外键
type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"column:Vid;primary_key"`           //视频主键
	Title         string `json:"title,omitempty" gorm:"column:VideoTitle"`             //视频简介、标题
	PlayUrl       string `json:"play_url,omitempty" gorm:"column:PlayUrl"`             //视频播放地址
	CoverUrl      string `json:"cover_url,omitempty" gorm:"column:CoverUrl"`           //视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty" gorm:"column:FavoriteCount"` //视频点赞总数
	CommentCount  int64  `json:"comment_count,omitempty" gorm:"column:CommentCount"`   //视频评论总数
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"column:IsFavorite"`       //true:已点赞，false:未点赞
	CreatedTime   int64  `json:"created_time,omitempty" gorm:"column:CreatedTime"`     //视频创建时间，时间戳形式
	User          User   `json:"author" gorm:"ForeignKey:UserID"`                      //视频作者
	UserID        int64  `json:"user_id" gorm:"column:Uid"`                            //外键：视频作者的ID
}

//结构体Video对应数据库中videos表
func (e *Video) TableName() string {
	return "videos"
}
