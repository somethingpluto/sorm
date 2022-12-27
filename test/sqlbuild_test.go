package test

import (
	"fmt"
	"testing"
)

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
