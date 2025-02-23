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

//插入新发布的视频，整个流程涉及如下操作
//1. videos表里面新增记录
//2. 该用户的视频发布数+1
func InsertNewVideo(video *model.Video) error {
	result := MyDB.Create(video)
	return result.Error
}

//优化：存在则更新，不存在则插入
//赞操作
func UpdateVideoIsFavorite(Vid int64, ActionType int) error {
	//更新单列
	//以video的主键更新
	result := MyDB.Model(&model.Video{}).Where("Vid=?", Vid).Update("IsFavorite", ActionType)
	return result.Error
}

//获取登录用户的所有点赞视频
func SelectUserLikeVideos(Uid int64) []model.Video {
	return nil
}
