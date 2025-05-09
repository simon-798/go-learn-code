package lesson3

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"uniqueIndex;size:20;not null"`
	Age       int    `gorm:"default:18"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(user *User) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("创建用户成功，ID:%d\n", user.ID)
	return nil
}

func GetUser(id uint) (*User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
func UpdateUser(id uint, userMap map[string]interface{}) error {
	//Updates：更新多个，Update：更新单个
	result := db.Model(&User{}).Where("id = ?", id).Updates(userMap)
	if result.Error != nil {
		return result.Error
	}

	fmt.Printf("更新影响行数：%d\n", result.RowsAffected)
	return nil
}

func DeleteUser(id uint) error {
	result := db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func selectAllUser() ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GormTest() {

	//init函数自动执行，初始化数据连接

	//
	newUser := &User{
		Name:  "simon",
		Email: "simon@gmail.com",
		Age:   18,
	}

	if err := CreateUser(newUser); err != nil {
		log.Fatal(err)
	}

	fetchedUser, err := GetUser(newUser.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("查询用户:%d\n", fetchedUser)

	updateUserMap := map[string]interface{}{
		"Name":  "tom",
		"Email": "tom@gmail.com",
		"Age":   20,
	}

	if err := UpdateUser(1, updateUserMap); err != nil {
		log.Fatal(err)
	}

	//查询所有用户
	userList, err := selectAllUser()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("查询所有用户：")
	for _, v := range userList {
		fmt.Printf("%+v\n", v)
	}

	if err := DeleteUser(2); err != nil {
		log.Fatal(err)
	}

	//验证是否已经删除
	_, err = GetUser(2)
	if err == gorm.ErrRecordNotFound {
		fmt.Println("用户已经删除")
	} else if err != nil {
		log.Fatal(err)
	}

}
