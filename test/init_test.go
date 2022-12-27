package test

import (
	_ "github.com/go-sql-driver/mysql"
	"sorm"
)

var Engine *sorm.Engine

func init() {
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/sorm"
	Engine, _ = sorm.NewEngine("mysql", dsn)
}
