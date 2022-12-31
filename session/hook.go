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

func (s *Session) CallMethod(method string, values ...interface{}) {
	param := []reflect.Value{reflect.ValueOf(s)}
	switch method {
	case BEFORE_INSERT:
		for _, value := range values {
			hookFunc := reflect.ValueOf(value).MethodByName(BEFORE_INSERT)
			result := hookFunc.Call(param)
			handleCallError(result)
		}
	case AFTER_INSERT:
		hookFunc := reflect.ValueOf(s.RefTable().Model).MethodByName(AFTER_INSERT)
		result := hookFunc.Call(param)
		handleCallError(result)
	}
}

func handleCallError(result []reflect.Value) {
	if len(result) > 0 {
		err, ok := result[0].Interface().(error)
		if ok {
			log.Error(err)
		}
	}
}
