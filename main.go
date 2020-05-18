package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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

	// 创建数据
	user := User{Name: "Jinzhu", Age: 18, Birthday: "1990-10-10"}
	if db.NewRecord(user) {
		db.Create(&user)
	}
	fmt.Printf("%+v\n", user)
	// os.Exit(0)

	// 修改数据
	user2 := new(User)
	user2.ID = 1
	db.First(user2) // 先查询出数据
	// RowsAffected 更新影响的行数
	c := db.Model(user2).Update("name", "Zoom").RowsAffected // 更新一个字段，如果不想gorm自动运行钩子 hooks ，使用 UpdateColumn()
	fmt.Printf("%d, %+v\n", c, *user2)

	c = db.Model(user2).Updates(map[string]interface{}{"age": 1, "birthday": "2001-10-10"}).RowsAffected // 更新多个字段，如果不想gorm自动运行钩子 hooks ，使用 UpdateColumns()
	fmt.Printf("%d, %+v\n", c, *user2)

	// 批量更新时 Hooks 不会运行（即更新多条记录）
	// db.Table("user").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
	//// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);

	// 使用 struct 更新时，只会更新非零值字段，若想更新所有字段，请使用map[string]interface{}
	// db.Model(User{}).Updates(User{Name: "hello", Age: 18})
	//// UPDATE users SET name='hello', age=18;

	// 表达式更新
	c = db.Model(user2).Update("age", gorm.Expr("age+?", 10)).RowsAffected
	fmt.Printf("%d, %+v\n", c, *user2) // 这里有个问题是users 的 age 还是更新前的值

	// 删除
	user3 := new(User)
	user3.ID = 2
	db.First(user3)             // 先查询出数据
	db.Delete(user3)            // 如果一个 model 有 DeletedAt 字段，他将自动获得软删除的功能！ 当调用 Delete 方法时， 记录不会真正的从数据库中被删除， 只会将DeletedAt 字段的值会被设置为当前时间
	db.Unscoped().Delete(user3) // 物理删除，不管有没有  DeletedAt 字段
	// 批量删除
	db.Where("name=?", "Jinzhu").Delete(User{}) // 这里也是有 DeletedAt 字段就软删除
	// db.Unscoped().Where("name=?", "Jinzhu").Delete(User{}) // 物理删除

}
