package test

import (
	"sorm/session"
	"testing"
)

func TestTransaction(t *testing.T) {
	s := Engine.NewSession()
	s = s.Model(&User{})
	user := User{
		Name: "aaa2",
		Age:  1,
	}
	_, err := s.Transaction(func(s *session.Session) (result interface{}, err error) {
		_, err = s.Insert(&user)
		return
	})

	if err != nil {
		t.Error(err)
	}
}
