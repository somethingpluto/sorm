package session

import "sorm/log"

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

type ITableName interface {
	TableName(s *Session) error
}

type IBeforeInsert interface {
	BeforeInsert(s *Session) error
}

type IAfterInsert interface {
	AfterInsert(s *Session) error
}

func (s *Session) CallMethod(method string) {
	model := s.RefTable().Model
	switch method {
	case TABLE_NAME:
		i, ok := model.(ITableName)
		if ok {
			err := i.TableName(s)
			if err != nil {
				log.Error(err)
			}
		}

	case BEFORE_INSERT:
		i, ok := model.(IBeforeInsert)
		if ok {
			err := i.BeforeInsert(s)
			if err != nil {
				log.Error(err)
			}
		}
	case AFTER_INSERT:
		i, ok := model.(IAfterInsert)
		if ok {
			err := i.AfterInsert(s)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
