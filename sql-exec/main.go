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

	// 1、执行原生SQL
	// 执行原生 SQL 时，不支持与其它方法的链式操作
	// db.Exec("DROP TABLE users;")
	// db.Exec("UPDATE orders SET shipped_at=? WHERE id IN (?)", time.Now(), []int64{11, 22, 33})

	// 2、Scan
	type Result struct {
		Name string
		Age  int
	}
	var result Result
	db.Raw("SELECT name, age FROM user WHERE id = ?", 1).Scan(&result) // Raw()构建原生sql 不会自己执行，需要配合其他方法获取执行结果
	fmt.Printf("%+v\n\n", result)
	// {Name:Jinzhu Age:10}

	// 3、通过 *sql.Row 或 *sql.Rows 获取查询结果
	var (
		name string
		age  int
	)
	row := db.Table("user").Where("id = ?", 3).Select("name, age").Row() // (*sql.Row)
	row.Scan(&name, &age)
	fmt.Println(name, age)
	// Tim 12

	// rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
	// defer rows.Close()
	// for rows.Next() {
	// 	// ...
	// 	rows.Scan(&name, &age, &email)
	// 	// ...
	// }

	// 原生 SQL
	// rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows() // (*sql.Rows, error)
	// defer rows.Close()
	// for rows.Next() {
	// 	// ...
	// 	rows.Scan(&name, &age, &email)
	// 	// ...
	// }

	// 4、将 sql.Rows 扫描至 model
	// rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
	// defer rows.Close()
	// for rows.Next() {
	// 	var user User
	// 	// ScanRows 扫描一行记录到 user
	// 	db.ScanRows(rows, &user)
	// 	// do something
	// }

}
