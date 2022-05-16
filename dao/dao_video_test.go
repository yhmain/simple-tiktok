package dao

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//测试是否按照时间戳 准确获取对应的视频
func TestSelectAllVideos(t *testing.T) {
	vid := SelectAllVideos(1652675588)
	for i := range vid {
		fmt.Printf("%v\n", vid[i])
	}
	assert.Equal(t, 2, len(vid))
}
