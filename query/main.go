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

	// 查询全部
	var users, users1, users2 []User
	db.Find(&users, "age>?", 13) // 带条件的内联查询
	fmt.Println(users)

	// 根据ID查询(仅当主键为整型时可用)
	var user User
	db.First(&user, 1)
	fmt.Printf("%+v\n", user)

	// 普通 where条件查询
	db.Where("name = ?", "jinzhu").First(&user)
	db.Where("name in (?)", []string{"Tim", "Jack"}).Find(&users)
	fmt.Println(users)
	db.Where("age between ? and ?", 10, 12).Find(&users)
	fmt.Println(users)

	// Struct 当通过结构体进行查询时，GORM将会只通过非零值字段查询，这意味着如果你的字段值为0，''， false 或者其他 零值时，将不会被用于构建查询条件
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键切片
	db.Where([]int64{20, 21, 22}).Find(&users)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);

	// Not 条件，作用与 Where 类似
	db.Not("name", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE name <> "jinzhu" ORDER BY id LIMIT 1;
	// Not In
	db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
	// 不在主键切片中
	db.Not([]int64{1, 2, 3}).First(&user)
	//// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
	db.Not([]int64{}).First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	// 普通 SQL
	db.Not("name = ?", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE NOT(name = "jinzhu") ORDER BY id LIMIT 1;
	// Struct
	db.Not(User{Name: "jinzhu"}).First(&user)
	//// SELECT * FROM users WHERE name <> "jinzhu" ORDER BY id LIMIT 1;

	// Or 条件
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	//// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
	// Struct
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';
	// Map
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';

	// 	选择字段
	// Select，指定你想从数据库中检索出的字段，默认会选择全部字段。
	db.Select("name, age").Find(&users)
	//// SELECT name, age FROM users;
	db.Select([]string{"name", "age"}).Find(&users)
	//// SELECT name, age FROM users;
	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	//// SELECT COALESCE(age,'42') FROM users;

	// 排序
	// Order，指定从数据库中检索出记录的顺序。设置第二个参数 reorder 为 true ，可以覆盖前面定义的排序条件。
	db.Order("age desc, name").Find(&users)
	//// SELECT * FROM users ORDER BY age desc, name;
	// 多字段排序
	db.Order("age desc").Order("name").Find(&users)
	//// SELECT * FROM users ORDER BY age desc, name;
	// 覆盖排序
	db.Order("age desc").Find(&users1).Order("age", true).Find(&users2) // 执行了多个查询
	//// SELECT * FROM users ORDER BY age desc; (users1)
	//// SELECT * FROM users ORDER BY age; (users2)

	// 数量
	// Limit，指定从数据库检索出的最大记录数。
	db.Limit(3).Find(&users)
	//// SELECT * FROM users LIMIT 3;
	// -1 取消 Limit 条件
	db.Limit(10).Find(&users).Limit(-1).Find(&users) // 执行了多个查询
	//// SELECT * FROM users LIMIT 10; (users1)
	//// SELECT * FROM users; (users2)

	// 偏移
	// Offset，指定开始返回记录前要跳过的记录数。
	db.Offset(3).Find(&users)
	//// SELECT * FROM users OFFSET 3;
	// -1 取消 Offset 条件
	db.Offset(10).Find(&users1).Offset(-1).Find(&users2) // 执行了多个查询
	//// SELECT * FROM users OFFSET 10; (users1)
	//// SELECT * FROM users; (users2)

	// 总数
	// Count，该 model 能获取的记录总数。
	var count int
	db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Find(&users).Count(&count) // 执行了多个查询
	//// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
	//// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)
	db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
	//// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)
	db.Table("deleted_users").Count(&count)
	//// SELECT count(*) FROM deleted_users;
	db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
	//// SELECT count( distinct(name) ) FROM deleted_users; (count)

	// Pluck
	// Pluck，查询 model 中的一个列作为切片，如果您想要查询多个列，您应该使用 Scan
	var ages []int64
	db.Find(&users).Pluck("age", &ages)

	var names []string
	db.Model(&User{}).Pluck("name", &names)

	db.Table("deleted_users").Pluck("name", &names)

	// 想查询多个字段？ 这样做：
	db.Select("name, age").Find(&users)

	// 扫描
	// Scan，扫描结果至一个 struct.
	type Result struct {
		Name string
		Age  int
	}
	var result Result
	db.Table("users").Select("name, age").Where("name = ?", "Antonio").Scan(&result)

	// 原生 SQL
	db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
}
