package test

import (
	"fmt"
	"sorm/log"
	"sorm/session"
	"testing"
)

type Teacher struct {
	Name string `sorm:"PRIMARY KEY"`
	Age  int
}

func (teacher *Teacher) BeforeInsert(s *session.Session) error {
	log.Info("BEFORE INSERT")
	return nil
}

func (teacher *Teacher) AfterInsert(s *session.Session) error {
	log.Info("AFTER INSERT")
	return nil
}

func TestInsertHooks(t *testing.T) {
	s := Engine.NewSession().Model(&Teacher{})
	err := s.DropTable()
	if err != nil {
		t.Error(err)
	}
	err = s.CreateTable()
	if err != nil {
		t.Error(err)
	}
	result, err := s.Insert(&Teacher{
		Name: "张三",
		Age:  10,
	}, &Teacher{
		Name: "李四",
		Age:  20,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
