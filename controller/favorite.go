package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yhmain/simple-tiktok/service"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	videoID := c.Query("video_id")
	actionType := c.Query("action_type") //1-点赞，2-取消点赞

	vid, err1 := strconv.ParseInt(videoID, 10, 64) //转化为int64
	atype, err2 := strconv.Atoi(actionType)        //转化为int
	if err1 != nil || err2 != nil {                //两个字符串解析是否出错
		c.JSON(http.StatusOK, InvalidVideoIDErr)
		return
	}
	//1代表指1，2代表值0
	updValue := map[int]int{1: 1, 2: 0}
	if err := service.UpdateVideoIsFavorite(vid, updValue[atype]); err != nil {
		c.JSON(http.StatusOK, UpdateSQLErr)
		return
	}
	c.JSON(http.StatusOK, Success)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {

	c.JSON(http.StatusOK, VideoListResponse{
		Response:  Success,
		VideoList: DemoVideos,
	})
}
