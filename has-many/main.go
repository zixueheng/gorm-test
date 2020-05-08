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

	// user := User{
	// 	Name:     "小花",
	// 	Age:      22,
	// 	Birthday: "1998-12-12",
	// 	CreditCards: []CreditCard{
	// 		{
	// 			No:    "62234324234324",
	// 			Issue: "中国工商银行",
	// 		},
	// 		{
	// 			No:    "46867876867867",
	// 			Issue: "中国建设银行",
	// 		},
	// 	},
	// }

	// if db.NewRecord(user) {
	// 	db.Create(&user)
	// }
	// fmt.Printf("%+v\n", user)
	// {Model:{ID:11 CreatedAt:2020-05-07 15:26:04.7375678 +0800 CST m=+0.026928601 UpdatedAt:2020-05-07 15:26:04.7375678 +0800 CST m=+0.026928601 DeletedAt:<nil>} Name:小花 Age:22 Birthday:1998-12-12 CreditCards:[{Model:{ID:2 CreatedAt:2020-05-07 15:26:04.8083785 +0800 CST m=+0.097739301 UpdatedAt:2020-05-07 15:26:04.8083785 +0800 CST m=+0.097739301 DeletedAt:<nil>} No:62234324234324 Issue:中国工商银行 UserID:11} {Model:{ID:3 CreatedAt:2020-05-07 15:26:04.9739347 +0800 CST m=+0.263295501 UpdatedAt:2020-05-07 15:26:04.9739347 +0800 CST m=+0.263295501 DeletedAt:<nil>} No:46867876867867 Issue:中国建设银行 UserID:11}]}

	user2 := new(User)
	user2.ID = 11
	db.First(user2)
	var cards []CreditCard
	db.Model(user2).Related(&cards, "CreditCards") // 获取关联
	user2.CreditCards = cards                      // 注意这里的 user2 的 CreditCard 是未赋值的 可以手动赋值下
	fmt.Printf("%+v\n", *user2)
	fmt.Printf("%+v\n", cards)
}
