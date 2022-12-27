package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sorm/v3"
	"sorm/v3/dialect"
	"sorm/v3/log"
	"sorm/v3/schema"
	"testing"
)

type User struct {
	Name string `sorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("mysql")

func TestParse(t *testing.T) {
	schema := schema.Parse(&User{}, TestDial)
	fmt.Println(schema.Name)
	for _, name := range schema.FieldNames {
		fmt.Println(name)
	}
	for _, field := range schema.Fields {
		fmt.Printf("%s %s %s\n", field.Name, field.Tag, field.Type)
	}
}

func TestTableOperation(t *testing.T) {
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/sorm"
	engine, _ := sorm.NewEngine("mysql", dsn)
	defer engine.Close()
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
}
