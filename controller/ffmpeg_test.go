package controller

import (
	"testing"

	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/assert"
)

//测试样例：图片是否可以从视频中提取出来
func TestExampleStream(t *testing.T) {
	paramMp4 := "../public/videos/bear.mp4"
	paramPng := "../public/covers/aa.jpeg"
	reader := ExampleReadFrameAsJpeg(paramMp4, 1)
	img, err := imaging.Decode(reader)
	if err != nil {
		t.Fatal(err)
	}
	err = imaging.Save(img, paramPng)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, err)
}
