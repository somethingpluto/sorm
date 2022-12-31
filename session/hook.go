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

type IBeforeInsert interface {
	BeforeInsert(s *Session) error
}

type IAfterInsert interface {
	AfterInsert(s *Session) error
}

func (s *Session) CallMethod(method string, values ...interface{}) {
	switch method {
	case BEFORE_INSERT:
		handleBeforeInsert(s, values...)
	case AFTER_INSERT:
		handleAfterInsert(s, values...)
	}
}

func handleBeforeInsert(s *Session, values ...interface{}) {
	dest := s.RefTable().Model
	if len(values) == 0 {
		i, ok := dest.(IBeforeInsert)
		if ok {
			err := i.BeforeInsert(s)
			if err != nil {
				log.Error(err)
			}
		}
	} else {
		for _, value := range values {
			i, ok := value.(IBeforeInsert)
			if ok {
				err := i.BeforeInsert(s)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}

}

func handleAfterInsert(s *Session, values ...interface{}) {
	dest := s.RefTable().Model
	if len(values) == 0 {
		i, ok := dest.(IAfterInsert)
		if ok {
			err := i.AfterInsert(s)
			if err != nil {
				log.Error(err)
			}
		}
	} else {
		for _, value := range values {
			i, ok := value.(IAfterInsert)
			if ok {
				err := i.AfterInsert(s)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
}
