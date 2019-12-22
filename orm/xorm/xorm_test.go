package xorm

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	Host     = "127.0.0.1"
	Port     = 3306
	User     = "root"
	Password = "yyb3911965"
	DB       = "world"
)

func Example() {
	// 定义 engine，相当于 db
	driver := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		User, Password, Host, Port, DB)
	engine, err := xorm.NewEngine(driver, dsn)
	if err != nil {
		panic(err)
	}
	// 定义表结构体
	type User struct {
		ID      int64     `xorm:"id pk autoincr"`
		Name    string    `xorm:"varchar(200) not null unique"`
		Created time.Time `xorm:"created not null"`
		Updated time.Time `xorm:"updated not null"`
	}
	// 同步表到数据库中
	if err := engine.Sync2(new(User)); err != nil {
		panic(err)
	}
	// 插入数据
	user := User{
		ID:   1,
		Name: "admin",
	}
	affected, err := engine.Insert(user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("affected: %d\n", affected)
	// 查询数据库
	name := "admin"
	has, err := engine.Where("name = ?", name).Desc("id").Get(&user)
	fmt.Printf("has: %t\n", has)
	fmt.Printf("user: %s\n", user.Name)
	// Output:
	// affected: 1
	// has: true
	// user: admin
}
