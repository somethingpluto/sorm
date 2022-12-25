package test

import (
	_ "github.com/go-sql-driver/mysql"
	"sorm/v0"
	"testing"
)

func TestNewEngine(t *testing.T) {
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/sorm"
	_, err := sorm.NewEngine("mysql", dsn)
	if err != nil {
		t.Log(err)
	}
	t.Log("数据库连接成功")
}
