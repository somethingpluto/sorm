package test

import (
	"fmt"
	"testing"
)

func TestInsertSQLBuild(t *testing.T) {
	user1 := User{
		Name: "张三",
		Age:  18,
	}
	user2 := User{
		Name: "李四",
		Age:  19,
	}
	session := Engine.NewSession().Model(&User{})
	insert, err := session.Insert(user1, user2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(insert)
}

func TestUpdateSQLBuild(t *testing.T) {
	session := Engine.NewSession().Model(&User{})
	err := session.DropTable()
	if err != nil {
		t.Error(err)
	}
	err = session.CreateTable()
	if err != nil {
		t.Error(err)
	}
	result, err := session.Where("Name = ?", "tom").Update("Age", "30")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
