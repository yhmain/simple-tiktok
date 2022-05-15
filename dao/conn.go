package main

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dsn := "root:123456@tcp(localhost:3306)/tiktok"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	fmt.Println("连接数据库成功")
	defer db.Close()
}
