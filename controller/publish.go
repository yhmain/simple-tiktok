package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yhmain/simple-tiktok/model"
	"github.com/yhmain/simple-tiktok/service"
)

//视频列表的响应体
type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
// 投稿接口，用户发布视频
// 2022/05/21  MuMu模拟器上测试成功，感觉与安卓机型有关
func Publish(c *gin.Context) {
	usertoken := c.MustGet("usertoken").(UserToken) //经过jwt鉴权后解析出的usertoekn

	//读取视频流数据
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, UploadFileFailedErr) //文件上传失败
		return
	}

	//类型
	typeOfA := reflect.TypeOf(data)
	fmt.Println("Data Type: ", typeOfA.Name(), typeOfA.Kind())

	fileName := filepath.Base(data.Filename) //视频名称
	//构建Video结构体所需要的参数
	newVideoID := service.GetNewVideoID()                   //获取新视频的ID
	paramTitle := c.PostForm("title")                       //视频标题
	playUrl := strconv.FormatInt(newVideoID, 10) + fileName //视频播放路径，因为视频都是放在一个目录下，所以视频名也要确保唯一

	//上传的视频保存到本地服务器
	saveFile := filepath.Join("./public/videos/", playUrl)
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, UploadFileFailedErr) //上传文件失败
		return
	}
	//从上传到本地的视频中抽取一帧作为封面
	coverUrl := strconv.FormatInt(newVideoID, 10) + ".jpeg" //视频封面路径
	saveImg := filepath.Join("./public/covers/", coverUrl)  //调用ffmpeg对应的 图片生成路径
	if err := GetVideoFrame(saveFile, saveImg); err != nil {
		c.JSON(http.StatusOK, ExtractVideoFailedErr) //生成封面失败
		return
	}

	//向数据库里面插入记录
	createdTime := time.Now().Unix() //获取当前时间戳
	userId := usertoken.UserID       //发布视频的用户ID
	newVideo := model.Video{Id: newVideoID, Title: paramTitle, PlayUrl: playUrl,
		CoverUrl: coverUrl, CreatedTime: createdTime, UserID: userId}
	if err := service.InsertNewVideo(&newVideo); err != nil {
		c.JSON(http.StatusOK, UploadFileFailedErr) //上传文件失败
		return

	}

	//上传成功！
	c.JSON(http.StatusOK, Success.WithMessage(playUrl+" upload successfully."))
}

// PublishList all users have same publish video list
//获取某用户发布的视频列表
func PublishList(c *gin.Context) {
	usertoken := c.MustGet("usertoken").(UserToken)

	//调用service 获取该用户发布的视频列表
	var videos = service.SelectVideosByUserID(usertoken.UserID)

	//返回视频列表和状态码
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  Success,
		VideoList: videos,
	})
}
