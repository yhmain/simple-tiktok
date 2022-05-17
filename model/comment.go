package model

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"column:Cid"`                 //评论id
	Content    string `json:"content,omitempty" gorm:"column:Content"`        //评论内容
	CreateDate string `json:"create_date,omitempty" gorm:"column:CreateDate"` //评论时间
	User       User   `json:"user"`                                           //创建该评论的用户id
	UserID     int64  `json:"user_id" gorm:"column:Uid"`                      //外键：发布评论的用户ID
	VideoID    int64  `json:"video_id" gorm:"column:Vid"`                     //外键：发布评论的视频ID
}

//结构体Comment对应数据库中comments表
func (e *Comment) TableName() string {
	return "comments"
}
