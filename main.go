package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	// 命令行运行：go build && ./simple-demo
	r.Run("192.168.1.106:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
