package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//数据库连接的常量
const (
	USER          = "root"
	PASSWORD      = "123456"
	SERVERIP      = "127.0.0.1"
	PORT          = "3306"
	DATABASE_NAME = "tiktok"
)

var MyDB *gorm.DB //全局变量，数据库连接

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		USER, PASSWORD, SERVERIP, PORT, DATABASE_NAME)
	var err error
	MyDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	fmt.Println("连接数据库成功")
}

// func GetDB() *gorm.DB {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		USER, PASSWORD, SERVERIP, PORT, DATABASE_NAME)
// 	db, err := gorm.Open(mysql.New(mysql.Config{
// 		DSN:                       dsn,   // DSN data source name
// 		DefaultStringSize:         256,   // string 类型字段的默认长度
// 		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
// 		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
// 		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
// 		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
// 	}), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("数据库连接失败", err)
// 		return nil
// 	}
// 	fmt.Println("连接数据库成功")
// 	return db
// }
