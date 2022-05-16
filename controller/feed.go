package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yhmain/simple-tiktok/model"
	"github.com/yhmain/simple-tiktok/service"
)

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var videos = service.SelectAllVideos()
	fmt.Printf("%v\n", videos)
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{StatusCode: 0},
		// VideoList: DemoVideos,
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
