package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sorm"
)

type User struct {
	Name string `sorm:"PRIMARY KEY"`
	Age  int
}

func main() {
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/sorm"
	engine, _ := sorm.NewEngine("mysql", dsn)
	defer engine.Close()

	s := engine.NewSession().Model(&User{})
	user1 := User{
		Name: "张三",
		Age:  18,
	}
	user2 := User{
		Name: "李四",
		Age:  19,
	}
	user3 := User{
		Name: "王五",
		Age:  18,
	}
	result, err := s.Insert(user1, user2, user3)
	if err != nil {
		panic(err)
	}
	fmt.Println("插入数据条数:", result)
}
