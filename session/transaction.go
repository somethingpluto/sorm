package session

import "sorm/log"

type TxFunc func(*Session) (interface{}, error)

func (s *Session) Begin() (err error) {
	log.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Rollback() (err error) {
	log.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Transaction(f TxFunc) (result interface{}, err error) {
	if err := s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		p := recover()
		if p != nil { // 运行触发了panic
			_ = s.Rollback()
		} else if err != nil { // 运行过程有err
			_ = s.Rollback()
		} else { // 运行过程没有err
			err = s.Commit()
		}
	}()

	return f(s)
}
