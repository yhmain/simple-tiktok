package dao

import (
	"github.com/yhmain/simple-tiktok/model"
)

//MyDB是dao包下的全局变量
//加载全部的 Video 到videos
func SelectAllVideos() []model.Video {
	// 坑：preload里不是对应的表的名字，而是结构体中的字段名字！！！
	var videos []model.Video
	MyDB.Preload("User").Find(&videos)
	return videos
}
