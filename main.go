package main

import "github.com/gin-gonic/gin"

const (
	SERVER_IP = "192.168.1.108:8080"
)

func main() {
	r := gin.Default()
	initRouter(r)
	r.Run(SERVER_IP)
}
