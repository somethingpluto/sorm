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

	s := engine.NewSession()
	err := s.Model(&User{}).CreateTable()
	if err != nil {
		panic(err)
	}
	fmt.Println(s.RefTable().Name, "创建成功")
	fmt.Printf("字段: %s", s.RefTable().FieldNames)
}
