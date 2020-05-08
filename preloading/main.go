package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User model
type User struct {
	gorm.Model
	Name        string
	Age         int
	Birthday    string
	CreditCards []CreditCard // User 有多个 Creditcard
}

// CreditCard model
type CreditCard struct {
	gorm.Model
	No     string // 卡号
	Issue  string // 发卡行
	UserID int    // 外键，外键名通常使用 has many 拥有者 User model 的类型加 主键 生成，所以外键名为 UserID.
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	defer db.Close()

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	user := new(User)
	user.ID = 11
	db.First(user)
	fmt.Printf("%+v\n\n", *user)
	// Preload 方法的参数应该是主结构体的字段名
	db.Preload("CreditCards").First(user) // 意思是 会将 user 的 credit_card 查询出来填充到结构体字段 CreditCards，实际gorm执行了先查询user 再查询 user 的 credit_card
	fmt.Printf("%+v\n\n", *user)

	db.Preload("CreditCards", "issue=?", "中国工商银行").First(user) // 只差 user 的 工行卡
	fmt.Printf("%+v\n\n", *user)

	// 其他示例
	// 1、嵌套预加载
	// db.Preload("Orders.OrderItems").First(&user)
	// db.Preload("Orders", "state = ?", "paid").Preload("Orders.OrderItems").First(&user)

	// 2、自定义预加载 SQL
	// 你可以通过传入 func(db *gorm.DB) *gorm.DB 来自定义预加载，比如：
	// db.Preload("Orders", func(db *gorm.DB) *gorm.DB {
	// return db.Order("orders.amount DESC")
	// }).First(&user)
	//// SELECT * FROM users;
	//// SELECT * FROM orders WHERE user_id IN (1,2,3,4) order by orders.amount DESC;
}
