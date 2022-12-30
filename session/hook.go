package session

import (
	"reflect"
	"sorm/log"
)

const (
	TABLE_NAME = "TableName"

	BEFORE_QUERY = "BeforeQuery"
	AFTER_QUERY  = "AfterQuery"

	BEFORE_UPDATE = "BeforeUpdate"
	AFTER_UPDATE  = "AfterUpdate"

	BEFORE_DELETE = "BeforeDelete"
	AFTER_DELETE  = "AfterDelete"

	BEFORE_INSERT = "BeforeInsert"
	AFTER_INSERT  = "AfterInsert"
)

func (s *Session) CallMethod(method string, value interface{}) {
	funcMethod := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil {
		funcMethod = reflect.ValueOf(value).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if funcMethod.IsValid() {
		v := funcMethod.Call(param)
		if len(v) > 0 {
			err, ok := v[0].Interface().(error)
			if ok {
				log.Error(err)
			}
		}
	}
}
