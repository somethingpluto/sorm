package session

import "sorm/log"

func (s *Session) Begin() (err error) {
	log.Info("TRANSACTION BEGIN")
	s.tx, err = s.db.Begin()
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("TRANSACTION COMMIT")
	err = s.tx.Commit()
	if err != nil {
		log.Error(err)
	}
	return err
}

func (s *Session) Rollback() (err error) {
	log.Info("TRANSACTION COMMIT")
	err = s.tx.Rollback()
	if err != nil {
		log.Error(err)
	}
	return err
}
