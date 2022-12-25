package test

import (
	"fmt"
	"sorm/v0"
)

var engine *sorm.Engine

func init() {
	var err error
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/sorm"
	engine, err = sorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库连接成功")
}
