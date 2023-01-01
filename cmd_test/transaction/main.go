package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sorm"
	"sorm/session"
)

type User struct {
	Name string `sorm:"PRIMARY KEY"`
	Age  int
}

func main() {
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/sorm"
	engine, _ := sorm.NewEngine("mysql", dsn)
	defer engine.Close()
	user := User{
		Name: "ccc",
		Age:  1,
	}
	s := engine.NewSession().Model(&User{})
	fmt.Println("事务第一次调用")
	// 1.第一次插入user
	_, err := s.Transaction(func(s *session.Session) (result interface{}, err error) {
		_, err = s.Insert(&user)
		return
	})
	fmt.Println("事务第二次调用")
	// 2.第二次插入user
	_, err = s.Transaction(func(s *session.Session) (result interface{}, err error) {
		_, err = s.Insert(&user)
		return
	})
	if err != nil {
		fmt.Println(err)
	}
}
