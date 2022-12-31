package test

import (
	"fmt"
	"testing"
)

func TestInsertSQLBuild(t *testing.T) {
	user1 := User{
		Name: "张三1",
		Age:  18,
	}
	user2 := User{
		Name: "李四2",
		Age:  19,
	}
	session := Engine.NewSession().Model(&User{})
	insert, err := session.Insert(&user1, &user2)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(insert)
}

func TestFindSQLBuild(t *testing.T) {
	session := Engine.NewSession().Model(&User{})
	var users []User
	err := session.Where("Age =?", 18).Order("Name desc").Limit(2).Offset(0).Find(&users)
	if err != nil {
		t.Error(err)
	}
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
}

func TestUpdateSQLBuild(t *testing.T) {
	session := Engine.NewSession().Model(&User{})
	result, err := session.Where("Age =?", 20).Update("Name", "ccccc")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

func TestDeleteSQLBuild(t *testing.T) {
	session := Engine.NewSession().Model(&User{})
	result, err := session.Where("Age = ?", 20).Delete()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
