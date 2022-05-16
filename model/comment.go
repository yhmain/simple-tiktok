package model

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
