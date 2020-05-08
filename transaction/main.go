package main

import (
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

	CreateUsers(db)
}

// CreateUsers 使用 Transaction 方法执行回调函数
func CreateUsers(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 在事务中做一些数据库操作 (这里应该使用 'tx' ，而不是 'db')
		if err := tx.Create(&User{Name: "Giraffe"}).Error; err != nil {
			// 返回任意 err ，整个事务都会 rollback
			return err
		}

		if err := tx.Create(&User{Name: "Lion"}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

// CreateUsers2 手动执行事务
func CreateUsers2(db *gorm.DB) error {
	// 请注意，事务一旦开始，你就应该使用 tx 作为数据库句柄
	tx := db.Begin()

	// 当下面的程序发生 panic(), defer 的函数会执行  recover()捕获panic 的值
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&User{Name: "Giraffe"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&User{Name: "Lion"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
