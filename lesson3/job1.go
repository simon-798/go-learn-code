package lesson3

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "gorm.io/gorm/logger"
	"log"
)

/*
*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、
age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
func Job1Method() {

	var err error
	db, err = gorm.Open(mysql.Open("root:20250423qwER@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		// 配置日志级别（生产环境建议关闭）
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("数据库连接失败:%v", err)
	}

	err = db.AutoMigrate(&Student{})
	if err != nil {
		log.Fatalf("数据自动迁移失败:%v", err)
	}

	fmt.Println("数据库连接成功")

	//创建一条记录
	student := &Student{Name: "test", Age: 13, Grade: "一年级"}
	db.Debug().Create(&student)
	//fmt.Printf("返回插入数据的主键:%d,result.Error:%d,返回插入记录的条数：%d", student.ID, result.Error, result.RowsAffected)

	//创建多条记录
	studentList := []*Student{
		{Name: "张三", Age: 20, Grade: "三年级"},
		{Name: "李四", Age: 21, Grade: "四年级"},
		{Name: "王五", Age: 22, Grade: "五年级"},
	}
	db.Debug().Create(studentList)

	db.Debug().Where("age > ?", 18).Find(&studentList)

	db.Debug().Model(&Student{}).Where("name=?", "张三").Update("grade", "四年级")

	db.Debug().Where("age < ?", 15).Delete(&Student{})

}

type Student struct {
	ID    uint   `gorm:"primary_key"` //默认主键自增
	Name  string `gorm:"type:varchar(255)"`
	Age   int    `gorm:"type:int"`
	Grade string `gorm:"type:varchar(255)"`
}
