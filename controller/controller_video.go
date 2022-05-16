package controller

import (
	"github.com/yhmain/simple-tiktok/dao"
	"github.com/yhmain/simple-tiktok/model"
)

func SelectAllVideos() []model.Video {
	var videos = dao.SelectAllVideos()
	for i := range videos {
		//获得拼接后的字符串
		videos[i].PlayUrl = ConcatByBuilder(PREFIX_VIDEOS, videos[i].PlayUrl)
		videos[i].CoverUrl = ConcatByBuilder(PREFIX_COVERS, videos[i].CoverUrl)
	}
	return videos
}
