package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User model
type User struct {
	gorm.Model
	Name       string
	Age        int
	Birthday   string
	CreditCard CreditCard // User 有一个 Creditcard
}

// CreditCard model
type CreditCard struct {
	gorm.Model
	No     string // 卡号
	Issue  string // 发卡行
	UserID int    // 外键，外键名通常使用 has one 拥有者 User model 的类型加 主键 生成，所以外键名为 UserID.
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	defer db.Close()

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	var users []User
	demo1(db).Find(&users)
	fmt.Printf("%+v\n\n", users)

	db.Scopes(AgeGreaterThan12, IDIn([]int{1, 2, 3})).Find(&users)
}

// 1、链式操作
func demo1(db *gorm.DB) *gorm.DB {
	// 创建一个查询，在调用立即执行方法前不会生成 Query 语句
	tx := db.Where("name=?", "Jinzhu")

	if true {
		tx = tx.Where("age < ?", 12)
	} else {
		tx = tx.Where("age >= ?", 12)
	}
	return tx
}

// 2、范围
// Scopes，Scope 是建立在链式操作的基础之上的。
// 基于它，你可以抽取一些通用逻辑，写出更多可重用的函数库。

// AgeGreaterThan12 固定条件函数
func AgeGreaterThan12(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 12)
}

// IDIn 带查询参数的闭包函数
func IDIn(ids []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (?)", ids)
	}
}
