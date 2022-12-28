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
	var users []User
	session := engine.NewSession().Model(&User{})
	fmt.Println("-----where限制")
	session.Where("Age = ?", "18").Find(&users)
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
	users = []User{}
	fmt.Println("-----order限制")
	session.Where("Age = ?", "18").Order("Name desc").Find(&users)
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
	users = []User{}

	fmt.Println("-----limit限制")
	session.Where("Age = ?", "18").Order("Name desc").Limit(1).Find(&users)
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
	users = []User{}

	fmt.Println("-----offset限制")
	session.Where("Age = ?", "18").Order("Name desc").Limit(1).Offset(1).Find(&users)
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}

}
