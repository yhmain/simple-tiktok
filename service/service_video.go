package service

//service 层处理核心业务逻辑输出

import (
	"github.com/yhmain/simple-tiktok/dao"
	"github.com/yhmain/simple-tiktok/model"

	idworker "github.com/gitstliu/go-id-worker"
)

//分布式id生成器
var videoIDGen *idworker.IdWorker

//ID生成器的初始化
func init() {
	videoIDGen = &idworker.IdWorker{}
	videoIDGen.InitIdWorker(1, 1) //WORKERID位数 (用于对工作进程进行编码), 数据中心ID位数 (用于对数据中心进行编码)
}

//生成新的视频ID
//失败则返回-1
func GetNewVideoID() int64 {
	id, err := videoIDGen.NextId() //生成新ID
	if err != nil {
		return -1
	}
	return id
}

//按照时间戳获取视频列表
func SelectVideosByTime(latestTime int64) []model.Video {
	//调用dao层获取数据
	var videos = dao.SelectVideosByTime(latestTime)
	for i := range videos {
		//获得拼接后的字符串
		videos[i].PlayUrl = ConcatByBuilder(PREFIX_VIDEOS, videos[i].PlayUrl)
		videos[i].CoverUrl = ConcatByBuilder(PREFIX_COVERS, videos[i].CoverUrl)
	}
	return videos
}

//按照用户ID获取视频列表
func SelectVideosByUserID(UserID int64) []model.Video {
	//调用dao层获取数据
	var videos = dao.SelectVideosByUserID(UserID)
	for i := range videos {
		//获得拼接后的字符串
		videos[i].PlayUrl = ConcatByBuilder(PREFIX_VIDEOS, videos[i].PlayUrl)
		videos[i].CoverUrl = ConcatByBuilder(PREFIX_COVERS, videos[i].CoverUrl)
	}
	return videos
}

//插入新发布的视频
func InsertNewVideo(video model.Video) error {
	return dao.InsertNewVideo(video)
}
