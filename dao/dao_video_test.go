package dao

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yhmain/simple-tiktok/model"
)

//测试是否按照时间戳 准确获取对应的视频
func TestSelectVideosByTime(t *testing.T) {
	vid := SelectVideosByTime(1652675588)
	assert.Equal(t, 30, len(vid))
}

//测试 用户ID：3发布的所有视频
func TestSelectVideosByUserID(t *testing.T) {
	videos := SelectVideosByUserID(3)
	for _, v := range videos {
		fmt.Printf("%v--%v\n", v.Id, v.CoverUrl)
	}
	assert.Equal(t, 7, len(videos))
}

//测试 插入新发布的视频
func TestInsertNewVideo(t *testing.T) {
	err := InsertNewVideo(&model.Video{Id: 1000, Title: "这是第一条视频的内容介绍：XXXX真可爱！", PlayUrl: "bear.mp4", CoverUrl: "bear.jpg", FavoriteCount: 10, CommentCount: 10, IsFavorite: false, CreatedTime: 1652597777, User: model.User{Id: 1, NickName: "zhanglei", Password: "douyin", FollowCount: 10, FollowerCount: 10, IsFollow: false}, UserID: 1})
	var expectedResult error
	assert.Equal(t, expectedResult, err)
}

//测试 点赞功能
func TestUpdateVideoIsFavorite100(t *testing.T) {
	err := UpdateVideoIsFavorite(1000, 1)
	var expectedResult error
	assert.Equal(t, expectedResult, err)
}

//测试 取消点赞功能
func TestUpdateVideoIsFavorite200(t *testing.T) {
	err := UpdateVideoIsFavorite(1000, 0)
	var expectedResult error
	assert.Equal(t, expectedResult, err)
}
