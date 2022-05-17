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
	for i := range vid {
		fmt.Printf("%v\n", vid[i])
	}
	assert.Equal(t, 2, len(vid))
}

//测试 用户ID：1发布的所有视频
func TestSelectVideosByUserID(t *testing.T) {
	videos := SelectVideosByUserID(1)
	assert.Equal(t, 2, len(videos))
}

//测试 插入新发布的视频
func TestInsertNewVideo(t *testing.T) {
	err := InsertNewVideo(model.Video{Id: 1000, Title: "这是第一条视频的内容介绍：XXXX真可爱！", PlayUrl: "bear.mp4", CoverUrl: "bear.jpg", FavoriteCount: 10, CommentCount: 10, IsFavorite: false, CreatedTime: 1652597777, User: model.User{Id: 1, NickName: "zhanglei", Password: "douyin", FollowCount: 10, FollowerCount: 10, IsFollow: false}, UserID: 1})
	var expectedResult error
	assert.Equal(t, expectedResult, err)
}
