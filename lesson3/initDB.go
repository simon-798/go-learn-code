package lesson3

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

const (
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
	DB_USER     = "root"
	DB_PASSWORD = "20250423qwER"
	DB_NAME     = "gorm_test"
)

var db *gorm.DB

func init() {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	var err error
	db, err = gorm.Open(mysql.Open(connect), &gorm.Config{
		// 配置日志级别（生产环境建议关闭）
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	//自动迁移表结构（如果不存在则创建）
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("自动迁移失败:%v", err)
	}

	fmt.Println("数据库连接成功")

}
