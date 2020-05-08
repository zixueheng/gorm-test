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

	// User 有多个 Language
	Languages []Language `gorm:"many2many:user_languages;"` // 关联表 user_languages

	// 跳过自动创建及更新，意思保存User时不会会主动创建或更新 Language和关联
	// Languages []Language `gorm:"many2many:user_languages;association_autoupdate:false;association_autocreate:false"`
}

// Language model
type Language struct {
	gorm.Model
	Name  string `gorm:"unique_index"`              // 名称
	Users []User `gorm:"many2many:user_languages;"` // 关联表 user_languages
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	defer db.Close()

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	// user := User{
	// 	Name:     "小欢",
	// 	Age:      28,
	// 	Birthday: "1992-12-12",
	// 	Languages: []Language{ // 注意这里 的language 名称不带ID会新建一个 ，带上ID就使用表中已有的，所以真实开发要先查询或者创建 language 然后再使用
	// 		{Model: gorm.Model{ID: 1}, Name: "中文"},
	// 		// {Model: gorm.Model{ID: 2}, Name: "英文"},
	// 	},
	// }
	// if db.NewRecord(user) {
	// 	db.Create(&user)
	// }
	// fmt.Printf("%+v\n", user)

	var users []User
	var language = Language{}
	db.First(&language, "id=?", 2)
	db.Model(&language).Related(&users, "Users")
	fmt.Printf("%+v\n", users) // 只有 User 信息，没有 Languages
	// [{Model:{ID:13 CreatedAt:2020-05-07 16:17:11 +0800 CST UpdatedAt:2020-05-07 16:17:11 +0800 CST DeletedAt:<nil>} Name:小天 Age:32 Birthday:1988-12-12T00:00:00+08:00 Languages:[]} {Model:{ID:19 CreatedAt:2020-05-07 16:34:21 +0800 CST UpdatedAt:2020-05-07 16:34:21 +0800 CST DeletedAt:<nil>} Name:小喜 Age:32 Birthday:1988-12-12T00:00:00+08:00 Languages:[]}]
}
