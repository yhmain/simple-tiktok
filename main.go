package main

import (
	"github.com/gin-gonic/gin"
)

const (
	SERVER_IP = "192.168.1.108:8080"
)

func main() {
	r := gin.Default()
	initRouter(r)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// 命令行运行：go build && ./simple-tiktok
	r.Run(SERVER_IP)
}
