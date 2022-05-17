package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yhmain/simple-tiktok/model"
	"github.com/yhmain/simple-tiktok/service"
)

//Feed响应体
type FeedResponse struct {
	Response                //标准响应体
	VideoList []model.Video `json:"video_list,omitempty"` //视频列表
	NextTime  int64         `json:"next_time,omitempty"`  //本次返回的视频中，发布最早的时间，作为下次请求的Latest_time
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	t := c.Query("latest_time")
	latest_time, err := strconv.ParseInt(t, 10, 64) //string转化为int64
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: -1, StatusMsg: "非法的时间戳！"}, //失败
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	} else {
		if latest_time == 0 { //空字符串会转化为0，则表示取当前时间的时间戳
			latest_time = time.Now().Unix()
		}
		//调用service层，获取数据
		var videos = service.SelectVideosByTime(latest_time)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0}, //成功
			VideoList: videos,
			NextTime:  time.Now().Unix(),
		})
	}
}
