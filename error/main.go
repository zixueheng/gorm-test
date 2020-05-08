package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User model
type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday string
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	defer db.Close()

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	var user User
	if err := db.Where("aaaa=?", 100).First(&user).Error; err != nil {
		fmt.Println(err)
		// Error 1054: Unknown column 'aaaa' in 'where clause'
	}

	// 未找到记录会报一个 RecordNotFound 错误
	if db.Where("id=?", 100).First(&user).RecordNotFound() {
		fmt.Println("记录未找到")
	}

	if err := db.Where("id=?", 100).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		fmt.Println("记录未找到")
	}
}
