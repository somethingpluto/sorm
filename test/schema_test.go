package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sorm"
	"sorm/dialect"
	"sorm/log"
	"sorm/schema"
	"testing"
	"time"
)

type User struct {
	Name string `sorm:"PRIMARY KEY"`
	Age  int
}

type Student struct {
	Name     string `sorm:"PRIMARY KEY"`
	Age      int
	worked   bool
	Birthday time.Time
}

var TestDial, _ = dialect.GetDialect("mysql")

func TestParseOne(t *testing.T) {
	sch := schema.Parse(&User{}, TestDial)
	fmt.Println("表名: ", sch.Name)
	for _, field := range sch.Fields {
		fmt.Printf("字段名: %s  类型: %s 标签: %s\n", field.Name, field.Type, field.Tag)
	}
	sch = schema.Parse(&Student{}, TestDial)
	fmt.Println("表名: ", sch.Name)
	for _, field := range sch.Fields {
		fmt.Printf("字段名: %s  类型: %s 标签: %s\n", field.Name, field.Type, field.Tag)
	}
}

func TestTableOperation(t *testing.T) {
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/sorm"
	engine, _ := sorm.NewEngine("mysql", dsn)
	defer engine.Close()
	fmt.Println("user表创建")
	s := engine.NewSession()
	table := s.Model(&User{})
	err := table.DropTable()
	if err != nil {
		log.Error(err)
	}
	err = table.CreateTable()
	if err != nil {
		log.Error(err)
	}
	fmt.Println("student表创建")
	table = s.Model(&Student{})
	err = table.DropTable()
	if err != nil {
		log.Error(err)
	}

	err = table.CreateTable()
	if err != nil {
		log.Error(err)
	}
}
