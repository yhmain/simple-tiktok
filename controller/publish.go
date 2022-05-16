package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/yhmain/simple-tiktok/model"
	"github.com/yhmain/simple-tiktok/service"
)

// const (
// 	PREFIX_VIDEOS string = "http://192.168.1.108:8081/videos/"
// 	PREFIX_COVERS string = "http://192.168.1.108:8081/covers/"
// )

//视频列表的响应体
type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
// 投稿接口，用户发布视频
// 出现multipart: nextpart: EOF 是因为表单为空
// 暂未完成
func Publish(c *gin.Context) {
	fmt.Printf("%#v\n", c)
	// fmt.Printf("%#v\n", c.Request.GetBody())
	fmt.Println("点击了发布！")
	c.MultipartForm()
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("%#v\n", body)
	// paramToken := c.PostForm("token")
	// fmt.Printf("Token:%s \n", paramToken)

	// //用户Token 反序列化，获取UserToken结构体
	// var token UserToken
	// json.Unmarshal([]byte(paramToken), &token)

	// //调用 service层获取结果
	// //user代表 token对应的用户
	// user, exist := service.SelectUserByNamePwd(token.Name, token.Password)
	// if !exist {
	// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// 	return
	// }

	//读取视频流数据
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		fmt.Printf("读取视频流出现错误！！！Error: %v\n", err.Error())
		return
	}

	fmt.Println("准备存到服务器")
	//数据库里面存入相应记录
	// newVideo := model.Video{}
	//准备存到服务器
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", 7889, filename)
	saveFile := filepath.Join("./public/videos/", finalName)
	fmt.Printf("%s--%s--%s", filename, finalName, saveFile)
	// gin 简单做了封装,拷贝了文件流
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//上传成功！
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
//获取某用户发布的视频列表
func PublishList(c *gin.Context) {
	paramToken := c.Query("token")

	//用户Token 反序列化，获取UserToken结构体
	var token UserToken
	json.Unmarshal([]byte(paramToken), &token)
	user, exist := service.SelectUserByNamePwd(token.Name, token.Password)
	if !exist {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg:  "Token对应的用户不存在！",
			},
			VideoList: nil,
		})
		return
	}

	//调用service 获取该用户发布的视频列表
	var videos = service.SelectVideosByUserID(user.Id)

	//返回视频列表和状态码
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})

}
