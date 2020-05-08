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

	// user := User{
	// 	Name:     "小明",
	// 	Age:      22,
	// 	Birthday: "1998-12-12",
	// 	CreditCard: CreditCard{
	// 		No:    "62234324234324",
	// 		Issue: "中国工商银行",
	// 	},
	// }

	// if db.NewRecord(user) {
	// 	db.Create(&user)
	// }
	// fmt.Printf("%+v\n", user)
	// {Model:{ID:10 CreatedAt:2020-05-07 15:08:24.9832538 +0800 CST m=+0.030917801 UpdatedAt:2020-05-07 15:08:24.9832538 +0800 CST m=+0.030917801 DeletedAt:<nil>} Name:小明 Age:22 Birthday:1998-12-12 CreditCard:{Model:{ID:1 CreatedAt:2020-05-07 15:08:25.0281333 +0800 CST m=+0.075797301 UpdatedAt:2020-05-07 15:08:25.0281333 +0800 CST m=+0.075797301 DeletedAt:<nil>} No:62234324234324 Issue:中国工商银行 UserID:10}}

	user2 := new(User)
	user2.ID = 10
	db.First(user2)
	fmt.Printf("%+v\n\n", *user2)

	var card CreditCard
	db.Model(user2).Related(&card, "CreditCard") // 获取关联
	//// SELECT * FROM credit_card WHERE user_id = 8; // 8 is user's primary key
	// CreditCard 是 user 的字段，其含义是，获取 user 的 CreditCard 并填充至 card 变量
	// 如果字段名与 model 名相同，比如上面的例子，此时字段名可以省略不写，像这样：
	// db.Model(user2).Related(&card)
	user2.CreditCard = card // 注意这里的 user2 的 CreditCard 是未赋值的 可以手动赋值下
	fmt.Printf("%+v\n\n", *user2)
	fmt.Printf("%+v\n\n", card)
}
