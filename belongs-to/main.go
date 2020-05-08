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

// Address model 属于 User
type Address struct {
	gorm.Model
	User   User
	UserID int // 默认使用 UserID 作为关联外键
	Addr   string
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	defer db.Close()

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	// address := Address{
	// 	User: User{
	// 		Name:     "小红",
	// 		Age:      18,
	// 		Birthday: "2000-10-10",
	// 	},
	// 	Addr: "安徽合肥市蜀山区110号",
	// }

	// if db.NewRecord(address) {
	// 	db.Create(&address) // User 会先被保存 然后保存 Address
	// }
	// fmt.Printf("%+v\n", address)
	// {Model:{ID:2 CreatedAt:2020-05-07 14:36:36.1454087 +0800 CST m=+0.079808201 UpdatedAt:2020-05-07 14:36:36.1454087 +0800 CST m=+0.079808201 DeletedAt:<nil>} User:{Model:{ID:8 CreatedAt:2020-05-07 14:36:36.0925289 +0800 CST m=+0.026928401 UpdatedAt:2020-05-07 14:36:36.0925289 +0800 CST m=+0.026928401 DeletedAt:<nil>} Name:小红 Age:18 Birthday:2000-10-10} UserID:8 Addr:安徽合肥市蜀山区110号}

	user := new(User)
	user.ID = 8
	db.First(user)
	var addr Address
	db.Model(user).Related(&addr) // 查找出来的 的 addr 的User 字段是未赋值的
	addr.User = *user             // 我这里手动给 User 赋值下
	fmt.Printf("%+v\n", addr)
	// {Model:{ID:2 CreatedAt:2020-05-07 14:36:36 +0800 CST UpdatedAt:2020-05-07 14:36:36 +0800 CST DeletedAt:<nil>} User:{Model:{ID:8 CreatedAt:2020-05-07 14:36:36 +0800 CST UpdatedAt:2020-05-07 14:36:36 +0800 CST DeletedAt:<nil>} Name:小红 Age:18 Birthday:2000-10-10T00:00:00+08:00} UserID:8 Addr:安徽合肥市蜀山区110号}

}
