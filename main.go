package main

import (
	"github.com/gin-gonic/gin"
)

const (
	SERVER_IP = "192.168.1.108"
)

func main() {
	r := gin.Default()
	initRouter(r)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// 命令行运行：go build && ./simple-tiktok
	r.Run(SERVER_IP + ":8080")
}

// package main

// import (
// 	"fmt"

// 	"github.com/yhmain/simple-tiktok/controller"
// )

// func main() {
// 	//查询所有用户
// 	for _, v := range controller.SelectAllUsers() {
// 		fmt.Printf("%v\n", v)
// 	}

// 	//查询所有视频
// 	for _, v := range controller.SelectAllVideos() {
// 		fmt.Printf("%v\n", v)
// 	}

// }
