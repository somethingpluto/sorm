package session

import "sorm/clause"

func (s *Session) Order(values ...interface{}) *Session {
	s.clause.Set(clause.ORDERBY, values...)
	return s
}
