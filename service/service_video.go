package service

//service 层处理核心业务逻辑输出

import (
	"github.com/yhmain/simple-tiktok/dao"
	"github.com/yhmain/simple-tiktok/model"
)

func SelectAllVideos(latestTime int64) []model.Video {
	//调用dao层获取数据
	var videos = dao.SelectAllVideos(latestTime)
	for i := range videos {
		//获得拼接后的字符串
		videos[i].PlayUrl = ConcatByBuilder(PREFIX_VIDEOS, videos[i].PlayUrl)
		videos[i].CoverUrl = ConcatByBuilder(PREFIX_COVERS, videos[i].CoverUrl)
	}
	return videos
}
