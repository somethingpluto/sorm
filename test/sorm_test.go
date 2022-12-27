package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sorm"
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

func TestExec(t *testing.T) {
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/sorm"
	engine, err := sorm.NewEngine("mysql", dsn)
	if err != nil {
		t.Log(err)
	}
	t.Log("数据库连接成功")
	session := engine.NewSession()
	// 创建数据表
	createTableSql := "CREATE TABLE user (name varchar(255),age int)"
	_, err = session.Raw(createTableSql).Exec()
	if err != nil {
		t.Error(err)
	}
	insertSql := "insert into user (name,age) values (\"张三\",18),(\"李四\",20)"
	_, err = session.Raw(insertSql).Exec()
	if err != nil {
		t.Error(err)
	}
	t.Log("数据表创建成功")
}

func TestQuery(t *testing.T) {
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/sorm"
	engine, err := sorm.NewEngine("mysql", dsn)
	if err != nil {
		t.Log(err)
	}
	t.Log("数据库连接成功")
	session := engine.NewSession()
	queryRowSql := "select * from user where name = ?"
	row := session.Raw(queryRowSql, "张三").QueryRow()
	var user User
	err = row.Scan(&user.Name, &user.Age)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v", user)
}
