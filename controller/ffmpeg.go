package controller

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

//从视频中抽取一帧作为封面
func GetVideoFrame(paramMp4, paramPng string) error {
	reader := ExampleReadFrameAsJpeg(paramMp4, 1)
	img, err := imaging.Decode(reader)
	if err != nil {
		return err
	}
	err = imaging.Save(img, paramPng)
	if err != nil {
		return err
	}
	return nil
}
