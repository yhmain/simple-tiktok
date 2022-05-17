package dao

import (
	"github.com/yhmain/simple-tiktok/model"
)

//MyDB是dao包下的全局变量
//加载 <=latestTime 的 Video 到videos
func SelectVideosByTime(latestTime int64) []model.Video {
	// 坑：preload里不是对应的表的名字，而是结构体中的字段名字！！！
	var videos []model.Video
	MyDB.Where("CreatedTime<=?", latestTime).Order("CreatedTime desc").Limit(FEED_VIDEOS_NUM).Preload("User").Find(&videos)
	return videos
}

//获取某用户发布的所有视频
func SelectVideosByUserID(UserID int64) []model.Video {
	var videos []model.Video
	MyDB.Where("Uid=?", UserID).Preload("User").Find(&videos)
	return videos
}

//插入新发布的视频
func InsertNewVideo(video model.Video) error {
	result := MyDB.Create(&video)
	return result.Error
}
