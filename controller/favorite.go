package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	videoID := c.Query("video_id")
	actionType := c.Query("action_type") //1-点赞，2-取消点赞

	vid, err1 := strconv.ParseInt(videoID, 10, 64)
	atype, err2 := strconv.Atoi(actionType)
	if err1 != nil || err2 != nil { //两个字符串解析是否出错
		c.JSON(http.StatusOK, InvalidVideoIDErr)
		return
	}
	fmt.Printf("Video ID: %v\n Actio Type: %v\n", vid, atype)
	c.JSON(http.StatusOK, Success)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  Success,
		VideoList: DemoVideos,
	})
}
