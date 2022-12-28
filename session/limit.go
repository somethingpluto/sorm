package session

import "sorm/clause"

func (s *Session) Limit(values ...interface{}) *Session {
	s.clause.Set(clause.LIMIT, values...)
	return s
}

func (s *Session) Offset(values ...interface{}) *Session {
	s.clause.Set(clause.OFFSET, values...)
	return s
}
